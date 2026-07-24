#!/usr/bin/env ruby
# frozen_string_literal: true

require "json"
require "open3"
require "optparse"
require "pathname"
require "shellwords"
require "tempfile"

PROJECT_ROOT = Pathname.new(__dir__).join("..").realpath
VERSION = PROJECT_ROOT.join("VERSION").read.strip

DEFAULT_REMOTE_HOST = ENV.fetch("XNIX_REMOTE_HOST", "root@q4")
DEFAULT_SOURCE_SYNC_MODE = ENV.fetch("XNIX_SOURCE_SYNC_MODE", "runtime")
DEFAULT_REMOTE_SOURCE_ROOT = ENV.fetch("XNIX_REMOTE_SOURCE_ROOT", "/home/xnix-build/xnix-runtime-source-matrix-#{DEFAULT_SOURCE_SYNC_MODE}-#{VERSION}")
DEFAULT_REMOTE_BUILD_ROOT = ENV.fetch("XNIX_REMOTE_BUILD_ROOT", "/home/xnix-build-cache")
DEFAULT_REMOTE_MATERIALS_ROOT = ENV.fetch("XNIX_REMOTE_MATERIALS_ROOT", "/home/xnix-run-materials")
DEFAULT_MATRIX_REPORT_OUTPUT = ENV.fetch("XNIX_KNOWN_WINAPP_MATRIX_REPORT_OUTPUT", "#{DEFAULT_REMOTE_MATERIALS_ROOT}/state/known-run-matrix-#{VERSION}.json")
DEFAULT_REMOTE_GO = ENV.fetch("XNIX_REMOTE_GO", "/home/xnix-toolchains/go1.24.4-linux-amd64/bin/go")
DEFAULT_LOCAL_SHELL = ENV.fetch("XNIX_LOCAL_SHELL", "/bin/zsh")
DEFAULT_APP_IDS = ENV.fetch("XNIX_KNOWN_WINAPP_MATRIX", "7zr,busybox-w32").split(",").map(&:strip).reject(&:empty?)

options = {
  execute: false,
  sync_source: true,
  source_sync_mode: DEFAULT_SOURCE_SYNC_MODE,
  local_shell: DEFAULT_LOCAL_SHELL,
  remote_host: DEFAULT_REMOTE_HOST,
  remote_source_root: DEFAULT_REMOTE_SOURCE_ROOT,
  remote_build_root: DEFAULT_REMOTE_BUILD_ROOT,
  remote_materials_root: DEFAULT_REMOTE_MATERIALS_ROOT,
  matrix_report_output: DEFAULT_MATRIX_REPORT_OUTPUT,
  remote_go: DEFAULT_REMOTE_GO,
  app_ids: DEFAULT_APP_IDS,
  timeout: ENV.fetch("XNIX_KNOWN_WINAPP_GUEST_TIMEOUT", "90s"),
  boot_timeout: ENV.fetch("XNIX_KNOWN_WINAPP_QEMU_BOOT_TIMEOUT", "180s")
}

OptionParser.new do |parser|
  parser.banner = "Usage: ruby scripts/remote_known_winapp_matrix_smoke.rb [--execute] [--app APP_ID ...]"
  parser.on("--execute", "Run the remote q4 build and known-app QEMU/Wine matrix smoke.") { options[:execute] = true }
  parser.on("--no-sync-source", "Use the existing remote source tree without syncing this checkout.") { options[:sync_source] = false }
  parser.on("--source-sync-mode MODE", "Source sync mode: runtime or full.") { |value| options[:source_sync_mode] = value }
  parser.on("--local-shell PATH", "Local login shell used for ssh/rsync alias resolution.") { |value| options[:local_shell] = value }
  parser.on("--remote HOST", "Remote SSH target, default: #{DEFAULT_REMOTE_HOST}") { |value| options[:remote_host] = value }
  parser.on("--remote-source-root PATH", "Remote source root under /home/xnix*.") { |value| options[:remote_source_root] = value }
  parser.on("--remote-build-root PATH", "Remote build cache root under /home/xnix*.") { |value| options[:remote_build_root] = value }
  parser.on("--remote-materials-root PATH", "Remote run materials root under /home/xnix*.") { |value| options[:remote_materials_root] = value }
  parser.on("--matrix-report-output PATH", "Remote aggregate matrix JSON report path under /home/xnix*.") { |value| options[:matrix_report_output] = value }
  parser.on("--remote-go PATH", "Remote Go binary path.") { |value| options[:remote_go] = value }
  parser.on("--app APP_ID", "Known Windows app id to include; repeatable.") do |value|
    options[:app_ids] = [] if options[:app_ids] == DEFAULT_APP_IDS
    options[:app_ids] << value
  end
  parser.on("--timeout DURATION", "Guest Wine execution timeout for each app.") { |value| options[:timeout] = value }
  parser.on("--qemu-boot-timeout DURATION", "QEMU SSH boot timeout for each app.") { |value| options[:boot_timeout] = value }
end.parse!

abort "remote known Windows app matrix smoke does not accept positional arguments" unless ARGV.empty?
abort "remote known Windows app matrix smoke requires at least one app" if options.fetch(:app_ids).empty?

if !ENV.key?("XNIX_REMOTE_SOURCE_ROOT") && options.fetch(:remote_source_root) == DEFAULT_REMOTE_SOURCE_ROOT
  options[:remote_source_root] = "/home/xnix-build/xnix-runtime-source-matrix-#{options.fetch(:source_sync_mode)}-#{VERSION}"
end

if !ENV.key?("XNIX_KNOWN_WINAPP_MATRIX_REPORT_OUTPUT") && options.fetch(:matrix_report_output) == DEFAULT_MATRIX_REPORT_OUTPUT
  options[:matrix_report_output] = "#{options.fetch(:remote_materials_root)}/state/known-run-matrix-#{VERSION}.json"
end

def ensure_remote_xnix_path!(label, path)
  clean = Pathname.new(path).cleanpath.to_s
  return clean if clean.start_with?("/home/xnix-")

  abort "#{label} must stay under /home/xnix-* on the remote build host"
end

def run_shell(shell, command)
  stdout, stderr, status = Open3.capture3(shell, "-lc", command, chdir: PROJECT_ROOT.to_s)
  [stdout, stderr, status.exitstatus]
end

def shell_join(argv)
  Shellwords.join(argv)
end

def source_sync_entries(mode)
  case mode
  when "runtime"
    %w[go.mod cmd internal runtime]
  when "full"
    ["."]
  else
    abort "source sync mode must be runtime or full"
  end
end

def safe_app_id(app_id)
  clean = app_id.gsub(/[^a-zA-Z0-9_.-]+/, "-")
  abort "known app id produced an empty safe id" if clean.empty?

  clean
end

remote_host = options.fetch(:remote_host)
source_sync_mode = options.fetch(:source_sync_mode)
source_entries = source_sync_entries(source_sync_mode)
remote_source_root = ensure_remote_xnix_path!("remote source root", options.fetch(:remote_source_root))
remote_build_root = ensure_remote_xnix_path!("remote build root", options.fetch(:remote_build_root))
remote_materials_root = ensure_remote_xnix_path!("remote materials root", options.fetch(:remote_materials_root))
matrix_report_output = ensure_remote_xnix_path!("matrix report output", options.fetch(:matrix_report_output))
remote_bin = "#{remote_build_root}/bin/xnix-runtime-go"
remote_key = "#{remote_materials_root}/ssh/id_ed25519"
remote_kernel = "#{remote_materials_root}/wine-guest/bzImage"
remote_cache_root = "#{remote_materials_root}/known-winapps"
remote_go = options.fetch(:remote_go)
remote_go_dir = Pathname.new(remote_go).dirname.to_s

app_plans = options.fetch(:app_ids).map do |app_id|
  safe_id = safe_app_id(app_id)
  {
    "app_id" => app_id,
    "report_output" => "#{remote_materials_root}/state/known-run-#{VERSION}-#{safe_id}.json",
    "serial_log_output" => "#{remote_materials_root}/state/qemu-serial-#{VERSION}-#{safe_id}.log"
  }
end

base_plan = {
  "schema_version" => "xnix.scripts.remote_known_windows_app_matrix_smoke.v1",
  "request_type" => "remote-known-winapp-matrix-smoke",
  "status" => options.fetch(:execute) ? "running" : "planned",
  "execute" => options.fetch(:execute),
  "local_shell" => options.fetch(:local_shell),
  "remote_host" => remote_host,
  "source_sync_planned" => options.fetch(:sync_source),
  "source_sync_mode" => source_sync_mode,
  "source_sync_entry_count" => source_entries.length,
  "source_sync_entries" => source_entries,
  "remote_source_root" => remote_source_root,
  "remote_build_root" => remote_build_root,
  "remote_materials_root" => remote_materials_root,
  "matrix_report_output" => matrix_report_output,
  "matrix_report_output_written" => false,
  "app_count" => app_plans.length,
  "app_ids" => app_plans.map { |entry| entry.fetch("app_id") },
  "backend" => "guest-wine",
  "start_qemu" => true,
  "guest_port" => "auto",
  "redact_output" => true,
  "apps" => app_plans,
  "host_root_modified" => false,
  "privileged_container_required" => false,
  "host_networking_required" => false,
  "docker_socket_mounted" => false,
  "broad_host_mount_required" => false
}

unless options.fetch(:execute)
  puts JSON.pretty_generate(base_plan)
  exit 0
end

if options.fetch(:sync_source)
  mkdir_stdout, mkdir_stderr, mkdir_status = run_shell(options.fetch(:local_shell), shell_join(["ssh", remote_host, shell_join(["mkdir", "-p", remote_source_root])]))
  unless mkdir_status.zero?
    warn mkdir_stdout unless mkdir_stdout.empty?
    warn mkdir_stderr unless mkdir_stderr.empty?
    warn "FAIL: remote known Windows app matrix source root preparation failed"
    exit 1
  end

  rsync_sources = if source_sync_mode == "full"
                    ["#{PROJECT_ROOT}/"]
                  else
                    source_entries.map { |entry| "#{PROJECT_ROOT}/#{entry}" }
                  end
  rsync_args = [
    "rsync",
    "-az",
    "--exclude", ".git",
    "--exclude", ".cache",
    "--exclude", "buildroot/output",
    "--exclude", "docs/claude-code-implementation-packages.md",
    *rsync_sources,
    "#{remote_host}:#{remote_source_root}/"
  ]
  rsync_stdout, rsync_stderr, rsync_status = run_shell(options.fetch(:local_shell), shell_join(rsync_args))
  unless rsync_status.zero?
    warn rsync_stdout unless rsync_stdout.empty?
    warn rsync_stderr unless rsync_stderr.empty?
    warn "FAIL: remote known Windows app matrix source sync failed"
    exit 1
  end
end

build_commands = [
  "set -eu",
  shell_join(["mkdir", "-p", "#{remote_build_root}/bin", "#{remote_build_root}/go-build", "#{remote_build_root}/go-mod", "#{remote_build_root}/tmp", "#{remote_materials_root}/state"]),
  "cd #{Shellwords.escape(remote_source_root)}",
  "PATH=#{Shellwords.escape(remote_go_dir)}:$PATH GOCACHE=#{Shellwords.escape("#{remote_build_root}/go-build")} GOMODCACHE=#{Shellwords.escape("#{remote_build_root}/go-mod")} GOTMPDIR=#{Shellwords.escape("#{remote_build_root}/tmp")} #{shell_join([remote_go, "build", "-o", remote_bin, "./cmd/xnix-runtime-go"])}"
]
build_stdout, build_stderr, build_status = run_shell(options.fetch(:local_shell), shell_join(["ssh", remote_host, build_commands.join("\n")]))
unless build_status.zero?
  warn build_stdout unless build_stdout.empty?
  warn build_stderr unless build_stderr.empty?
  warn "FAIL: remote known Windows app matrix build failed"
  exit 1
end

results = []
app_plans.each do |app_plan|
  runtime_args = [
    remote_bin,
    "windows-known-app-run",
    "--backend", "guest-wine",
    "--start-qemu",
    "--port", "auto",
    "--redact-output",
    "--app", app_plan.fetch("app_id"),
    "--cache-root", remote_cache_root,
    "--key", remote_key,
    "--qemu-binary", "qemu-system-i386",
    "--qemu-kernel", remote_kernel,
    "--qemu-memory", "1024M",
    "--qemu-cpus", "2",
    "--qemu-cpu", "qemu32",
    "--qemu-serial-log", app_plan.fetch("serial_log_output"),
    "--qemu-boot-timeout", options.fetch(:boot_timeout),
    "--timeout", options.fetch(:timeout),
    "--report-output", app_plan.fetch("report_output")
  ]
  stdout, stderr, status = run_shell(options.fetch(:local_shell), shell_join(["ssh", remote_host, shell_join(runtime_args)]))
  if status.zero?
    payload = JSON.parse(stdout)
    guest = payload.fetch("guest_payload", {}).fetch("guest", {})
    results << app_plan.merge(
      "status" => payload.fetch("status"),
      "app_version" => payload.fetch("app_version"),
      "executable_name" => payload.fetch("executable_name"),
      "backend" => payload.fetch("backend"),
      "guest_start_mode" => payload.fetch("guest_start_mode", ""),
      "guest_started" => payload.fetch("guest_started", false),
      "guest_port_auto" => payload.fetch("guest_port_auto", false),
      "checksum_verified" => payload.fetch("checksum_verified", false),
      "raw_output_redacted" => payload.fetch("raw_output_redacted", false),
      "marker_observed" => payload.fetch("marker_observed", false),
      "qemu_serial_log_written" => payload.fetch("qemu_serial_log_written", false),
      "qemu_executed" => payload.fetch("qemu_executed", false),
      "wine_executed" => payload.fetch("wine_executed", false),
      "stdout_bytes" => guest.fetch("stdout_bytes", 0),
      "stdout_line_count" => guest.fetch("stdout_line_count", 0)
    )
  else
    results << app_plan.merge(
      "status" => "failed",
      "failure_reason" => "remote known app command failed",
      "stderr_bytes" => stderr.bytesize,
      "stdout_bytes" => stdout.bytesize
    )
  end
end

passed_count = results.count { |result| result.fetch("status") == "passed" }
matrix = base_plan.merge(
  "status" => passed_count == results.length ? "passed" : "failed",
  "execute" => true,
  "passed_count" => passed_count,
  "failed_count" => results.length - passed_count,
  "apps" => results,
  "matrix_report_output_written" => true
)

Tempfile.create(["xnix-known-winapp-matrix-", ".json"]) do |file|
  file.write(JSON.pretty_generate(matrix))
  file.write("\n")
  file.flush

  remote_report_dir = Pathname.new(matrix_report_output).dirname.to_s
  mkdir_stdout, mkdir_stderr, mkdir_status = run_shell(options.fetch(:local_shell), shell_join(["ssh", remote_host, shell_join(["mkdir", "-p", remote_report_dir])]))
  unless mkdir_status.zero?
    warn mkdir_stdout unless mkdir_stdout.empty?
    warn mkdir_stderr unless mkdir_stderr.empty?
    warn "FAIL: remote known Windows app matrix report directory preparation failed"
    exit 1
  end

  scp_stdout, scp_stderr, scp_status = run_shell(options.fetch(:local_shell), shell_join(["scp", file.path, "#{remote_host}:#{matrix_report_output}"]))
  unless scp_status.zero?
    warn scp_stdout unless scp_stdout.empty?
    warn scp_stderr unless scp_stderr.empty?
    warn "FAIL: remote known Windows app matrix report upload failed"
    exit 1
  end
end

puts JSON.pretty_generate(matrix)
exit(matrix.fetch("status") == "passed" ? 0 : 1)
