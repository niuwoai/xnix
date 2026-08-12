#!/usr/bin/env ruby
# frozen_string_literal: true

require "fileutils"
require "json"
require "open3"
require "optparse"
require "pathname"
require "shellwords"
require "timeout"
require_relative "../lib/xnix/milestone"

PROJECT_ROOT = Pathname.new(__dir__).join("..").realpath
VERSION = PROJECT_ROOT.join("VERSION").read.strip
DEFAULT_REMOTE_HOST = ENV.fetch("XNIX_REMOTE_HOST", "root@q4")
DEFAULT_LOCAL_SHELL = ENV.fetch("XNIX_LOCAL_SHELL", "/bin/zsh")
DEFAULT_REMOTE_SOURCE_ROOT = ENV.fetch("XNIX_Q4_FULL_SMOKE_SOURCE_ROOT", "/home/xnix-build/xnix-q4-full-smoke-#{VERSION}")
DEFAULT_REPORT_ROOT = ENV.fetch("XNIX_Q4_FULL_SMOKE_REPORT_ROOT", PROJECT_ROOT.join("output", "q4-full-smoke-#{VERSION}").to_s)
DEFAULT_TIMEOUT_SECONDS = Integer(ENV.fetch("XNIX_Q4_FULL_SMOKE_TIMEOUT_SECONDS", "7200"), 10)
DEFAULT_TOOLS_BASE_IMAGE = ENV.fetch("XNIX_Q4_FULL_SMOKE_TOOLS_BASE_IMAGE", "python:3.12-slim")
DEFAULT_WINE_BASE_IMAGE = ENV.fetch("XNIX_Q4_FULL_SMOKE_WINE_BASE_IMAGE", DEFAULT_TOOLS_BASE_IMAGE)
DEFAULT_CONTAINER_MEMORY_LIMIT = ENV.fetch("XNIX_Q4_FULL_SMOKE_CONTAINER_MEMORY_LIMIT", "12g")
DEFAULT_CONTAINER_CPU_LIMIT = ENV.fetch("XNIX_Q4_FULL_SMOKE_CONTAINER_CPU_LIMIT", "2.0")
DEFAULT_KNOWN_WINAPP_FETCH_TIMEOUT = ENV.fetch("XNIX_Q4_FULL_SMOKE_KNOWN_WINAPP_FETCH_TIMEOUT", "300s")
DEFAULT_REMOTE_PREFLIGHT_TIMEOUT_SECONDS = Integer(ENV.fetch("XNIX_Q4_FULL_SMOKE_PREFLIGHT_TIMEOUT_SECONDS", "45"), 10)
NON_COLIMA_DOCKER_ENV = "XNIX_ALLOW_NON_COLIMA_DOCKER"
NON_COLIMA_DOCKER_ASSIGNMENT = "XNIX_ALLOW_NON_COLIMA_DOCKER=1"
TOOLS_BASE_IMAGE_ENV = "XNIX_TOOLS_BASE_IMAGE"
WINE_BASE_IMAGE_ENV = "XNIX_WINE_BASE_IMAGE"
CONTAINER_MEMORY_LIMIT_ENV = "XNIX_CONTAINER_MEMORY_LIMIT"
CONTAINER_CPU_LIMIT_ENV = "XNIX_CONTAINER_CPU_LIMIT"
KNOWN_WINAPP_FETCH_TIMEOUT_ENV = "XNIX_KNOWN_WINAPP_FETCH_TIMEOUT"
SOURCE_SYNC_EXCLUDES = %w[
  .git
  .cache
  .gocache
  output
  buildroot/output
  docs/claude-code-implementation-packages.md
].freeze
REMOTE_REPORT_FILES = {
  "output/full-smoke-report.json" => "q4-full-smoke-report.json",
  "output/full-smoke-report.md" => "q4-full-smoke-report.md",
  "output/serial.log" => "q4-serial.log"
}.freeze

options = {
  execute: false,
  local_shell: DEFAULT_LOCAL_SHELL,
  remote_host: DEFAULT_REMOTE_HOST,
  remote_source_root: DEFAULT_REMOTE_SOURCE_ROOT,
  report_root: DEFAULT_REPORT_ROOT,
  timeout_seconds: DEFAULT_TIMEOUT_SECONDS,
  tools_base_image: DEFAULT_TOOLS_BASE_IMAGE,
  wine_base_image: DEFAULT_WINE_BASE_IMAGE,
  container_memory_limit: DEFAULT_CONTAINER_MEMORY_LIMIT,
  container_cpu_limit: DEFAULT_CONTAINER_CPU_LIMIT,
  known_winapp_fetch_timeout: DEFAULT_KNOWN_WINAPP_FETCH_TIMEOUT,
  remote_preflight_timeout_seconds: DEFAULT_REMOTE_PREFLIGHT_TIMEOUT_SECONDS,
  output: ENV.fetch("XNIX_Q4_FULL_SMOKE_OUTPUT", "")
}

OptionParser.new do |parser|
  parser.banner = "Usage: ruby scripts/q4_full_smoke.rb [--execute]"
  parser.on("--execute", "Sync the checkout and run the full smoke on q4 instead of the macOS host.") { options[:execute] = true }
  parser.on("--local-shell PATH", "Local shell used for ssh/rsync alias resolution.") { |value| options[:local_shell] = value }
  parser.on("--remote HOST", "Remote SSH target, default: #{DEFAULT_REMOTE_HOST}.") { |value| options[:remote_host] = value }
  parser.on("--remote-source-root PATH", "Remote source root under /home/xnix-* or /tmp/xnix-*.") { |value| options[:remote_source_root] = value }
  parser.on("--report-root PATH", "Local report directory under this checkout or /tmp/xnix-*.") { |value| options[:report_root] = value }
  parser.on("--timeout-seconds SECONDS", Integer, "Timeout for q4 sync, execution, and report fetch.") { |value| options[:timeout_seconds] = value }
  parser.on("--tools-base-image IMAGE", "q4-local Docker base image for the tools stage, default: #{DEFAULT_TOOLS_BASE_IMAGE}.") { |value| options[:tools_base_image] = value }
  parser.on("--wine-base-image IMAGE", "q4-local Docker base image for the Wine GUI smoke image, default: #{DEFAULT_WINE_BASE_IMAGE}.") { |value| options[:wine_base_image] = value }
  parser.on("--container-memory-limit LIMIT", "Remote q4 Docker run memory limit, default: #{DEFAULT_CONTAINER_MEMORY_LIMIT}.") { |value| options[:container_memory_limit] = value }
  parser.on("--container-cpu-limit LIMIT", "Remote q4 Docker run CPU limit, default: #{DEFAULT_CONTAINER_CPU_LIMIT}.") { |value| options[:container_cpu_limit] = value }
  parser.on("--known-winapp-fetch-timeout DURATION", "Remote q4 known Windows app fetch timeout, default: #{DEFAULT_KNOWN_WINAPP_FETCH_TIMEOUT}.") { |value| options[:known_winapp_fetch_timeout] = value }
  parser.on("--remote-preflight-timeout-seconds SECONDS", Integer, "Timeout for the q4 SSH/tooling preflight, default: #{DEFAULT_REMOTE_PREFLIGHT_TIMEOUT_SECONDS}.") { |value| options[:remote_preflight_timeout_seconds] = value }
  parser.on("--output PATH", "Write the planned, passed, or failed summary JSON under this checkout or /tmp/xnix-*.") { |value| options[:output] = value }
end.parse!

abort "q4 full smoke does not accept positional arguments" unless ARGV.empty?

def ensure_remote_xnix_path!(label, path)
  clean = Pathname.new(path).cleanpath.to_s
  return clean if clean.start_with?("/home/xnix-") || clean.start_with?("/tmp/xnix-")

  abort "#{label} must stay under /home/xnix-* or /tmp/xnix-* on q4"
end

def ensure_local_xnix_path!(label, path)
  clean = Pathname.new(path).expand_path(PROJECT_ROOT).cleanpath
  return clean if clean.to_s.start_with?(PROJECT_ROOT.to_s)
  return clean if clean.to_s.start_with?("/tmp/xnix-")

  abort "#{label} must stay under this checkout or /tmp/xnix-*"
end

def ensure_docker_image_ref!(label, value)
  clean = value.to_s.strip
  abort "#{label} must be non-empty" if clean.empty?
  abort "#{label} must be single-line" if clean.include?("\n") || clean.include?("\r")
  abort "#{label} must not contain shell whitespace" if clean.match?(/\s/)

  clean
end

def ensure_docker_limit!(label, value)
  clean = value.to_s.strip
  abort "#{label} must be non-empty" if clean.empty?
  abort "#{label} must be single-line" if clean.include?("\n") || clean.include?("\r")
  abort "#{label} must not contain shell whitespace" if clean.match?(/\s/)

  clean
end

def ensure_duration!(label, value)
  clean = value.to_s.strip
  abort "#{label} must be non-empty" if clean.empty?
  abort "#{label} must be single-line" if clean.include?("\n") || clean.include?("\r")
  abort "#{label} must not contain shell whitespace" if clean.match?(/\s/)
  abort "#{label} must look like a Go duration such as 300s or 5m" unless clean.match?(/\A\d+(?:ns|us|µs|ms|s|m|h)\z/)

  clean
end

def shell_join(argv)
  Shellwords.join(argv)
end

def ssh_command(remote_host, remote_command)
  [
    "ssh",
    "-o", "BatchMode=yes",
    "-o", "ConnectTimeout=15",
    "-o", "ServerAliveInterval=15",
    "-o", "ServerAliveCountMax=4",
    remote_host,
    remote_command
  ]
end

def rsync_ssh_transport
  [
    "ssh",
    "-o", "BatchMode=yes",
    "-o", "ConnectTimeout=15",
    "-o", "ServerAliveInterval=15",
    "-o", "ServerAliveCountMax=4"
  ].join(" ")
end

def run_shell(shell, command, timeout_seconds:)
  stdout = +""
  stderr = +""
  status = nil
  timed_out = false
  Open3.popen3(shell, "-lc", command, chdir: PROJECT_ROOT.to_s, pgroup: true) do |_stdin, out, err, wait_thread|
    out_reader = Thread.new { stdout = out.read }
    err_reader = Thread.new { stderr = err.read }
    begin
      Timeout.timeout(timeout_seconds) { status = wait_thread.value }
    rescue Timeout::Error
      timed_out = true
      terminate_process_group(wait_thread.pid)
      status = wait_thread.value
    ensure
      out_reader.join
      err_reader.join
    end
  end
  stderr = [stderr, "q4 full smoke shell operation timed out after #{timeout_seconds}s"].reject(&:empty?).join("\n") if timed_out
  [stdout, stderr, timed_out ? 124 : status.exitstatus]
end

def run_shell_streaming_to_stderr(shell, command, timeout_seconds:)
  tail = []
  status = nil
  timed_out = false
  Open3.popen2e(shell, "-lc", command, chdir: PROJECT_ROOT.to_s, pgroup: true) do |_stdin, output, wait_thread|
    reader = Thread.new do
      output.each_line do |line|
        $stderr.print line
        tail << line
        tail.shift if tail.length > 400
      end
    end
    begin
      Timeout.timeout(timeout_seconds) { status = wait_thread.value }
    rescue Timeout::Error
      timed_out = true
      terminate_process_group(wait_thread.pid)
      status = wait_thread.value
    ensure
      reader.join
    end
  end
  tail << "q4 full smoke timed out after #{timeout_seconds}s\n" if timed_out
  [tail.join, timed_out ? 124 : status.exitstatus]
end

def terminate_process_group(pid)
  begin
    Process.kill("TERM", -pid)
  rescue Errno::ESRCH
    nil
  end
  sleep 2
  begin
    Process.kill("KILL", -pid)
  rescue Errno::ESRCH
    nil
  end
end

def emit_json(payload, output_path)
  text = JSON.pretty_generate(payload) + "\n"
  File.write(output_path, text) if output_path
  puts text
end

def display_local_path(path)
  path.to_s.start_with?(PROJECT_ROOT.to_s) ? path.relative_path_from(PROJECT_ROOT).to_s : path.to_s
end

def fetch_remote_reports(options, remote_source_root, report_root)
  FileUtils.mkdir_p(report_root)
  fetched = []
  REMOTE_REPORT_FILES.each do |remote_relative_path, local_name|
    local_path = report_root.join(local_name)
    command = shell_join([
      "rsync",
      "-az",
      "-e", rsync_ssh_transport,
      "--timeout", "60",
      "--contimeout", "15",
      "#{options.fetch(:remote_host)}:#{remote_source_root}/#{remote_relative_path}",
      local_path.to_s
    ])
    _stdout, _stderr, status = run_shell(options.fetch(:local_shell), command, timeout_seconds: options.fetch(:timeout_seconds))
    fetched << local_name if status.zero?
  end
  fetched
end

def remote_preflight_script
  [
    "set -eu",
    "printf 'XNIX_Q4_FULL_SMOKE_REMOTE_READY=true\\n'",
    "command -v ruby >/dev/null",
    "command -v rsync >/dev/null",
    "command -v docker >/dev/null"
  ].join("\n")
end

remote_source_root = ensure_remote_xnix_path!("remote source root", options.fetch(:remote_source_root))
report_root = ensure_local_xnix_path!("report root", options.fetch(:report_root))
output_path = options.fetch(:output).empty? ? nil : ensure_local_xnix_path!("output", options.fetch(:output))
tools_base_image = ensure_docker_image_ref!("tools base image", options.fetch(:tools_base_image))
wine_base_image = ensure_docker_image_ref!("Wine GUI smoke base image", options.fetch(:wine_base_image))
container_memory_limit = ensure_docker_limit!("container memory limit", options.fetch(:container_memory_limit))
container_cpu_limit = ensure_docker_limit!("container CPU limit", options.fetch(:container_cpu_limit))
known_winapp_fetch_timeout = ensure_duration!("known Windows app fetch timeout", options.fetch(:known_winapp_fetch_timeout))
full_build_required = Xnix::Milestone.full_build_required?(VERSION)

plan = {
  "schema_version" => "xnix.scripts.q4_full_smoke.v1",
  "request_type" => "q4-full-smoke",
  "version" => VERSION,
  "status" => options.fetch(:execute) ? "running" : "planned",
  "execute" => options.fetch(:execute),
  "local_shell" => options.fetch(:local_shell),
  "remote_host" => options.fetch(:remote_host),
  "remote_source_root" => remote_source_root,
  "source_sync_mode" => "full",
  "source_sync_entries" => ["."],
  "source_sync_excludes" => SOURCE_SYNC_EXCLUDES,
  "source_sync_ssh_transport" => "BatchMode+ConnectTimeout+ServerAlive",
  "protected_claude_file_excluded" => SOURCE_SYNC_EXCLUDES.include?("docs/claude-code-implementation-packages.md"),
  "report_root" => display_local_path(report_root),
  "remote_report_files" => REMOTE_REPORT_FILES.keys,
  "remote_full_smoke_command" => "ruby scripts/full_smoke.rb",
  "remote_docker_context_override" => NON_COLIMA_DOCKER_ASSIGNMENT,
  "non_colima_docker_opt_in_env" => NON_COLIMA_DOCKER_ENV,
  "tools_base_image_env" => TOOLS_BASE_IMAGE_ENV,
  "tools_base_image" => tools_base_image,
  "q4_cached_tools_base_image_preferred" => true,
  "wine_base_image_env" => WINE_BASE_IMAGE_ENV,
  "wine_base_image" => wine_base_image,
  "q4_cached_wine_base_image_preferred" => true,
  "container_memory_limit_env" => CONTAINER_MEMORY_LIMIT_ENV,
  "container_memory_limit" => container_memory_limit,
  "container_cpu_limit_env" => CONTAINER_CPU_LIMIT_ENV,
  "container_cpu_limit" => container_cpu_limit,
  "q4_heavy_smoke_resource_override" => true,
  "known_winapp_fetch_timeout_env" => KNOWN_WINAPP_FETCH_TIMEOUT_ENV,
  "known_winapp_fetch_timeout" => known_winapp_fetch_timeout,
  "q4_known_winapp_fetch_timeout_extended" => true,
  "full_build_required" => full_build_required,
  "twentieth_version_full_build_required" => full_build_required,
  "remote_full_smoke_planned" => true,
  "remote_preflight_planned" => true,
  "remote_preflight_command" => "ruby+rsync+docker availability over SSH",
  "remote_preflight_timeout_seconds" => options.fetch(:remote_preflight_timeout_seconds),
  "q4_compile_required" => true,
  "host_compilation_avoided" => true,
  "host_root_modified" => false,
  "privileged_container_required" => false,
  "host_networking_required" => false,
  "docker_socket_mounted" => false,
  "broad_host_mount_required" => false,
  "remote_timeout_seconds" => options.fetch(:timeout_seconds)
}

unless options.fetch(:execute)
  emit_json(plan, output_path)
  exit 0
end

unless full_build_required
  failed = plan.merge(
    "status" => "failed",
    "remote_full_smoke_completed" => false,
    "failure" => "q4 full smoke requires a twentieth formal version"
  )
  emit_json(failed, output_path)
  exit 1
end

preflight_stdout, preflight_stderr, preflight_status = run_shell(
  options.fetch(:local_shell),
  shell_join(ssh_command(options.fetch(:remote_host), remote_preflight_script)),
  timeout_seconds: options.fetch(:remote_preflight_timeout_seconds)
)
unless preflight_status.zero?
  warn preflight_stdout unless preflight_stdout.empty?
  warn preflight_stderr unless preflight_stderr.empty?
  failed = plan.merge(
    "status" => "failed",
    "remote_preflight_completed" => false,
    "remote_preflight_exit_code" => preflight_status,
    "remote_full_smoke_completed" => false,
    "failure" => "q4 full smoke remote preflight failed"
  )
  emit_json(failed, output_path)
  exit 1
end

mkdir_stdout, mkdir_stderr, mkdir_status = run_shell(
  options.fetch(:local_shell),
  shell_join(ssh_command(options.fetch(:remote_host), shell_join(["mkdir", "-p", remote_source_root]))),
  timeout_seconds: options.fetch(:timeout_seconds)
)
unless mkdir_status.zero?
  warn mkdir_stdout unless mkdir_stdout.empty?
  warn mkdir_stderr unless mkdir_stderr.empty?
  failed = plan.merge("status" => "failed", "failure" => "remote source root preparation failed")
  emit_json(failed, output_path)
  exit 1
end

rsync_args = [
  "rsync",
  "-az",
  "-e", rsync_ssh_transport,
  "--delete",
  "--timeout", "60",
  "--contimeout", "15",
  *SOURCE_SYNC_EXCLUDES.flat_map { |entry| ["--exclude", entry] },
  "#{PROJECT_ROOT}/",
  "#{options.fetch(:remote_host)}:#{remote_source_root}/"
]
rsync_stdout, rsync_stderr, rsync_status = run_shell(
  options.fetch(:local_shell),
  shell_join(rsync_args),
  timeout_seconds: options.fetch(:timeout_seconds)
)
unless rsync_status.zero?
  warn rsync_stdout unless rsync_stdout.empty?
  warn rsync_stderr unless rsync_stderr.empty?
  failed = plan.merge(
    "status" => "failed",
    "source_sync_completed" => false,
    "source_sync_exit_code" => rsync_status,
    "failure" => "q4 full smoke source sync failed"
  )
  emit_json(failed, output_path)
  exit 1
end

remote_commands = [
  "set -eu",
  "cd #{Shellwords.escape(remote_source_root)}",
  "#{NON_COLIMA_DOCKER_ASSIGNMENT} #{TOOLS_BASE_IMAGE_ENV}=#{Shellwords.escape(tools_base_image)} #{WINE_BASE_IMAGE_ENV}=#{Shellwords.escape(wine_base_image)} #{CONTAINER_MEMORY_LIMIT_ENV}=#{Shellwords.escape(container_memory_limit)} #{CONTAINER_CPU_LIMIT_ENV}=#{Shellwords.escape(container_cpu_limit)} #{KNOWN_WINAPP_FETCH_TIMEOUT_ENV}=#{Shellwords.escape(known_winapp_fetch_timeout)} ruby scripts/full_smoke.rb"
]
remote_tail, remote_status = run_shell_streaming_to_stderr(
  options.fetch(:local_shell),
  shell_join(ssh_command(options.fetch(:remote_host), remote_commands.join("\n"))),
  timeout_seconds: options.fetch(:timeout_seconds)
)
fetched_reports = fetch_remote_reports(options, remote_source_root, report_root)

if remote_status.zero?
  finished = plan.merge(
    "status" => "passed",
    "remote_full_smoke_completed" => true,
    "report_fetch_completed" => true,
    "fetched_report_count" => fetched_reports.length,
    "fetched_reports" => fetched_reports
  )
  emit_json(finished, output_path)
  exit 0
end

failed = plan.merge(
  "status" => "failed",
  "remote_full_smoke_completed" => false,
  "report_fetch_completed" => fetched_reports.any?,
  "fetched_report_count" => fetched_reports.length,
  "fetched_reports" => fetched_reports,
  "failure" => "q4 full smoke failed",
  "remote_tail" => remote_tail.lines.last(40).join
)
emit_json(failed, output_path)
exit 1
