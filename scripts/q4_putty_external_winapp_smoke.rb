#!/usr/bin/env ruby
# frozen_string_literal: true

require "fileutils"
require "json"
require "open3"
require "optparse"
require "pathname"
require "shellwords"
require "timeout"

PROJECT_ROOT = Pathname.new(__dir__).join("..").realpath
VERSION = PROJECT_ROOT.join("VERSION").read.strip

SCHEMA_VERSION = "xnix.scripts.q4_putty_external_winapp_smoke.v1"
REQUEST_TYPE = "q4-putty-external-winapp-smoke"
PUTTY_VERSION = "0.84"
PUTTY_EXE = "putty.exe"
PUTTY_URL = "https://the.earth.li/~sgtatham/putty/0.84/w32/putty.exe"
PUTTY_SHA256 = "d5a83cd1233f6da38fa82b14d970dbb2c2705769b5ebabb464918b9b57180bc4"
PUTTY_WINDOW_MATCH = "PuTTY"

DEFAULT_REMOTE_HOST = ENV.fetch("XNIX_REMOTE_HOST", "root@q4")
DEFAULT_LOCAL_SHELL = ENV.fetch("XNIX_LOCAL_SHELL", "/bin/zsh")
DEFAULT_REMOTE_MATERIALS_ROOT = ENV.fetch("XNIX_REMOTE_MATERIALS_ROOT", "/home/xnix-run-materials")
DEFAULT_REMOTE_TIMEOUT_SECONDS = Integer(ENV.fetch("XNIX_Q4_PUTTY_TIMEOUT_SECONDS", "1200"), 10)
DEFAULT_OUTPUT = PROJECT_ROOT.join("output", "q4-putty-external-winapp-smoke-#{VERSION}.json").to_s
DEFAULT_MARKDOWN_OUTPUT = PROJECT_ROOT.join("output", "q4-putty-external-winapp-smoke-#{VERSION}.md").to_s

options = {
  execute: false,
  local_shell: DEFAULT_LOCAL_SHELL,
  remote_host: DEFAULT_REMOTE_HOST,
  remote_materials_root: DEFAULT_REMOTE_MATERIALS_ROOT,
  remote_timeout_seconds: DEFAULT_REMOTE_TIMEOUT_SECONDS,
  output: ENV.fetch("XNIX_Q4_PUTTY_OUTPUT", DEFAULT_OUTPUT),
  markdown_output: ENV.fetch("XNIX_Q4_PUTTY_MARKDOWN_OUTPUT", DEFAULT_MARKDOWN_OUTPUT)
}

OptionParser.new do |parser|
  parser.banner = "Usage: ruby scripts/q4_putty_external_winapp_smoke.rb [--execute]"
  parser.on("--execute", "Download, verify, and run the real PuTTY standalone Windows GUI app on q4.") { options[:execute] = true }
  parser.on("--local-shell PATH", "Local shell used for SSH alias resolution.") { |value| options[:local_shell] = value }
  parser.on("--remote HOST", "Remote SSH target, default: #{DEFAULT_REMOTE_HOST}.") { |value| options[:remote_host] = value }
  parser.on("--remote-materials-root PATH", "Remote materials root under /home/xnix-* or /tmp/xnix-*.") { |value| options[:remote_materials_root] = value }
  parser.on("--remote-timeout-seconds SECONDS", Integer, "Timeout for q4 download and staged execution.") { |value| options[:remote_timeout_seconds] = value }
  parser.on("--output PATH", "Write JSON result under this checkout or /tmp/xnix-*.") { |value| options[:output] = value }
  parser.on("--markdown-output PATH", "Write Markdown result under this checkout or /tmp/xnix-*.") { |value| options[:markdown_output] = value }
end.parse!

abort "q4 PuTTY external Windows app smoke does not accept positional arguments" unless ARGV.empty?

def ensure_local_output_path!(label, path)
  clean = Pathname.new(path).expand_path(PROJECT_ROOT).cleanpath
  return clean if clean.to_s.start_with?(PROJECT_ROOT.to_s)
  return clean if clean.to_s.start_with?("/tmp/xnix-")

  abort "#{label} must stay under this checkout or /tmp/xnix-*"
end

def ensure_remote_xnix_path!(label, path)
  clean = Pathname.new(path).cleanpath.to_s
  return clean if clean.start_with?("/home/xnix-") || clean.start_with?("/tmp/xnix-")

  abort "#{label} must stay under /home/xnix-* or /tmp/xnix-* on q4"
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
      begin
        Process.kill("TERM", -wait_thread.pid)
      rescue Errno::ESRCH
        nil
      end
      sleep 2
      begin
        Process.kill("KILL", -wait_thread.pid)
      rescue Errno::ESRCH
        nil
      end
      status = wait_thread.value
    ensure
      out_reader.join
      err_reader.join
    end
  end
  stderr = [stderr, "q4 PuTTY external Windows app smoke operation timed out after #{timeout_seconds}s"].reject(&:empty?).join("\n") if timed_out
  [stdout, stderr, timed_out ? 124 : status.exitstatus]
end

def emit_json(payload, output_path)
  text = JSON.pretty_generate(payload) + "\n"
  FileUtils.mkdir_p(output_path.dirname)
  File.write(output_path, text)
  puts text
end

def bool(payload, key)
  payload[key] == true
end

def remote_step!(shell, remote_host, remote_command, timeout_seconds, failure)
  stdout, stderr, status = run_shell(
    shell,
    shell_join(ssh_command(remote_host, remote_command)),
    timeout_seconds: timeout_seconds
  )
  unless status.zero?
    warn stdout unless stdout.empty?
    warn stderr unless stderr.empty?
    abort failure
  end
  stdout
end

remote_host = options.fetch(:remote_host)
remote_materials_root = ensure_remote_xnix_path!("remote materials root", options.fetch(:remote_materials_root))
remote_app_root = "#{remote_materials_root}/third-party/putty/#{PUTTY_VERSION}"
remote_executable = "#{remote_app_root}/#{PUTTY_EXE}"
output_path = ensure_local_output_path!("output path", options.fetch(:output))
markdown_output_path = ensure_local_output_path!("markdown output path", options.fetch(:markdown_output))

staged_command = [
  "ruby",
  "scripts/q4_staged_desktop_external_winapp_smoke.rb",
  "--execute",
  "--fixture", "external",
  "--remote-executable", remote_executable,
  "--app-id", "org.xnix.external.putty",
  "--display-name", "PuTTY",
  "--window-match", PUTTY_WINDOW_MATCH,
  "--remote", remote_host,
  "--remote-timeout-seconds", options.fetch(:remote_timeout_seconds).to_s,
  "--output", output_path.to_s,
  "--markdown-output", markdown_output_path.to_s
]

plan = {
  "schema_version" => SCHEMA_VERSION,
  "request_type" => REQUEST_TYPE,
  "version" => VERSION,
  "status" => options.fetch(:execute) ? "running" : "planned",
  "execute" => options.fetch(:execute),
  "remote_host" => remote_host,
  "source_kind" => "official-putty-release",
  "putty_version" => PUTTY_VERSION,
  "putty_executable" => PUTTY_EXE,
  "putty_source_page" => "https://www.chiark.greenend.org.uk/~sgtatham/putty/latest.html?textonly=1",
  "putty_download_url" => PUTTY_URL,
  "putty_sha256" => PUTTY_SHA256,
  "putty_download_planned" => true,
  "putty_downloaded" => false,
  "putty_sha256_verified" => false,
  "putty_executable_configured" => true,
  "putty_executable_path_exposed_only_for_operator" => true,
  "remote_materials_root" => remote_materials_root,
  "remote_app_root" => remote_app_root,
  "remote_executable" => remote_executable,
  "app_id" => "org.xnix.external.putty",
  "display_name" => "PuTTY",
  "window_match" => PUTTY_WINDOW_MATCH,
  "real_third_party_windows_app" => true,
  "single_file_windows_app" => true,
  "delegated_script" => "scripts/q4_staged_desktop_external_winapp_smoke.rb",
  "delegated_command" => staged_command,
  "output_path" => output_path.to_s,
  "markdown_output_path" => markdown_output_path.to_s,
  "runtime_owned" => true,
  "go_runtime_backed" => true,
  "kde_policy_owner" => false,
  "q4_download_required" => true,
  "q4_compile_required" => true,
  "host_compilation_avoided" => true,
  "host_download_avoided" => true,
  "full_smoke_required" => false,
  "host_root_modified" => false,
  "privileged_container_required" => false,
  "host_networking_required" => false,
  "docker_socket_mounted" => false,
  "broad_host_mount_required" => false
}

unless options.fetch(:execute)
  emit_json(plan, output_path)
  exit 0
end

remote_step!(
  options.fetch(:local_shell),
  remote_host,
  shell_join(["mkdir", "-p", remote_app_root]),
  options.fetch(:remote_timeout_seconds),
  "q4 PuTTY remote directory preparation failed"
)
remote_step!(
  options.fetch(:local_shell),
  remote_host,
  shell_join(["curl", "-L", "--fail", "--show-error", "--retry", "8", "--retry-delay", "2", "--retry-all-errors", "--connect-timeout", "20", "--continue-at", "-", "--output", remote_executable, PUTTY_URL]),
  options.fetch(:remote_timeout_seconds),
  "q4 PuTTY executable download failed"
)
plan["putty_downloaded"] = true
checksum_stdout = remote_step!(
  options.fetch(:local_shell),
  remote_host,
  shell_join(["sha256sum", remote_executable]),
  options.fetch(:remote_timeout_seconds),
  "q4 PuTTY executable checksum calculation failed"
)
actual_sha256 = checksum_stdout.split.first.to_s
abort "q4 PuTTY executable SHA256 mismatch" unless actual_sha256 == PUTTY_SHA256
plan["putty_sha256_verified"] = true

staged_stdout, staged_stderr, staged_status = run_shell(
  options.fetch(:local_shell),
  shell_join(staged_command),
  timeout_seconds: options.fetch(:remote_timeout_seconds)
)
unless staged_status.zero?
  warn staged_stdout unless staged_stdout.empty?
  warn staged_stderr unless staged_stderr.empty?
  abort "q4 PuTTY staged external run failed"
end

delegated = JSON.parse(staged_stdout)
passed = delegated.fetch("status") == "passed" &&
         bool(delegated, "external_file_bridge_ready") &&
         bool(delegated, "windows_process_file_argument_window_observed") &&
         bool(delegated, "go_owned_q4_staged_external_winapp_acceptance_ready") &&
         delegated.fetch("accepted_application_detail_compatibility_state", "") == "runtime-accepted-real-app-run" &&
         delegated.fetch("kde_page_from_accepted_application_detail_compatibility_state", "") == "runtime-accepted-real-app-run"
abort "q4 PuTTY staged external run did not reach accepted state" unless passed

result = plan.merge(
  "status" => "passed",
  "delegated_status" => delegated.fetch("status"),
  "remote_build_completed" => bool(delegated, "remote_build_completed"),
  "external_file_bridge_ready" => bool(delegated, "external_file_bridge_ready"),
  "windows_process_file_argument_window_observed" => bool(delegated, "windows_process_file_argument_window_observed"),
  "acceptance_ready" => bool(delegated, "go_owned_q4_staged_external_winapp_acceptance_ready"),
  "accepted_application_detail_state" => delegated.fetch("accepted_application_detail_compatibility_state", ""),
  "kde_accepted_page_state" => delegated.fetch("kde_page_from_accepted_application_detail_compatibility_state", ""),
  "artifact_fetch_count" => delegated.fetch("artifact_fetch_count", 0)
)
emit_json(result, output_path)
