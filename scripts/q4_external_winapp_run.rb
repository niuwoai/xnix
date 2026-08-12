#!/usr/bin/env ruby
# frozen_string_literal: true

require "digest"
require "fileutils"
require "json"
require "open3"
require "optparse"
require "pathname"
require "shellwords"
require "timeout"

PROJECT_ROOT = Pathname.new(__dir__).join("..").realpath
VERSION = PROJECT_ROOT.join("VERSION").read.strip
Q4_STAGED_EXTERNAL_SMOKE = PROJECT_ROOT.join("scripts/q4_staged_desktop_external_winapp_smoke.rb")

SCHEMA_VERSION = "xnix.scripts.q4_external_winapp_run.v1"
REQUEST_TYPE = "q4-external-winapp-run"

DEFAULT_REMOTE_HOST = ENV.fetch("XNIX_REMOTE_HOST", "root@q4")
DEFAULT_LOCAL_SHELL = ENV.fetch("XNIX_LOCAL_SHELL", "/bin/zsh")
DEFAULT_REMOTE_MATERIALS_ROOT = ENV.fetch("XNIX_REMOTE_MATERIALS_ROOT", "/home/xnix-run-materials")
DEFAULT_REMOTE_TIMEOUT_SECONDS = Integer(ENV.fetch("XNIX_Q4_EXTERNAL_WINAPP_RUN_TIMEOUT_SECONDS", "900"), 10)
DEFAULT_OUTPUT = PROJECT_ROOT.join("output", "q4-external-winapp-run-#{VERSION}.json").to_s
DEFAULT_MARKDOWN_OUTPUT = PROJECT_ROOT.join("output", "q4-external-winapp-run-#{VERSION}.md").to_s
DEFAULT_APP_ID = "org.xnix.external.uploaded"
DEFAULT_DISPLAY_NAME = "Uploaded Windows App"

options = {
  execute: false,
  local_shell: DEFAULT_LOCAL_SHELL,
  remote_host: DEFAULT_REMOTE_HOST,
  remote_materials_root: DEFAULT_REMOTE_MATERIALS_ROOT,
  executable: "",
  app_id: ENV.fetch("XNIX_Q4_EXTERNAL_WINAPP_APP_ID", DEFAULT_APP_ID),
  display_name: ENV.fetch("XNIX_Q4_EXTERNAL_WINAPP_DISPLAY_NAME", DEFAULT_DISPLAY_NAME),
  window_match: ENV.fetch("XNIX_Q4_EXTERNAL_WINAPP_WINDOW_MATCH", ""),
  remote_timeout_seconds: DEFAULT_REMOTE_TIMEOUT_SECONDS,
  output: ENV.fetch("XNIX_Q4_EXTERNAL_WINAPP_RUN_OUTPUT", DEFAULT_OUTPUT),
  markdown_output: ENV.fetch("XNIX_Q4_EXTERNAL_WINAPP_RUN_MARKDOWN_OUTPUT", DEFAULT_MARKDOWN_OUTPUT)
}

OptionParser.new do |parser|
  parser.banner = "Usage: ruby scripts/q4_external_winapp_run.rb --executable /tmp/xnix-app/app.exe --window-match TEXT [--execute]"
  parser.on("--execute", "Upload the local executable to q4 and run it through the staged external Runtime/KDE path.") { options[:execute] = true }
  parser.on("--local-shell PATH", "Local shell used for SSH/SCP alias resolution.") { |value| options[:local_shell] = value }
  parser.on("--remote HOST", "Remote SSH target, default: #{DEFAULT_REMOTE_HOST}.") { |value| options[:remote_host] = value }
  parser.on("--remote-materials-root PATH", "Remote materials root under /home/xnix-* or /tmp/xnix-*.") { |value| options[:remote_materials_root] = value }
  parser.on("--executable PATH", "Local Windows .exe under this checkout or /tmp/xnix-*.") { |value| options[:executable] = value }
  parser.on("--app-id ID", "Application id for the uploaded executable.") { |value| options[:app_id] = value }
  parser.on("--display-name NAME", "Display name for the uploaded executable.") { |value| options[:display_name] = value }
  parser.on("--window-match TEXT", "Observed-window text proving the Windows process consumed the file-open argument.") { |value| options[:window_match] = value }
  parser.on("--remote-timeout-seconds SECONDS", Integer, "Timeout for q4 upload and staged external execution.") { |value| options[:remote_timeout_seconds] = value }
  parser.on("--output PATH", "Write JSON result under this checkout or /tmp/xnix-*.") { |value| options[:output] = value }
  parser.on("--markdown-output PATH", "Write Markdown result under this checkout or /tmp/xnix-*.") { |value| options[:markdown_output] = value }
end.parse!

abort "q4 external Windows app run does not accept positional arguments" unless ARGV.empty?

def ensure_local_output_path!(label, path)
  clean = Pathname.new(path).expand_path(PROJECT_ROOT).cleanpath
  return clean if clean.to_s.start_with?(PROJECT_ROOT.to_s)
  return clean if clean.to_s.start_with?("/tmp/xnix-")

  abort "#{label} must stay under this checkout or /tmp/xnix-*"
end

def ensure_local_executable_path!(path, execute:)
  value = path.to_s.strip
  abort "--executable is required" if value.empty?

  clean = Pathname.new(value).expand_path(PROJECT_ROOT).cleanpath
  allowed = clean.to_s.start_with?(PROJECT_ROOT.to_s) || clean.to_s.start_with?("/tmp/xnix-")
  abort "local executable must stay under this checkout or /tmp/xnix-*" unless allowed
  return clean unless execute

  abort "local executable must exist" unless clean.file?
  abort "local executable must be a regular file" unless clean.file?
  File.open(clean, "rb") do |file|
    abort "local executable must have an MZ header" unless file.read(2) == "MZ"
  end
  clean
end

def ensure_remote_xnix_path!(label, path)
  clean = Pathname.new(path).cleanpath.to_s
  return clean if clean.start_with?("/home/xnix-") || clean.start_with?("/tmp/xnix-")

  abort "#{label} must stay under /home/xnix-* or /tmp/xnix-* on q4"
end

def ensure_non_empty_single_line!(label, value)
  clean = value.to_s.strip
  abort "#{label} must be non-empty" if clean.empty?
  abort "#{label} must be single-line" if clean.include?("\n") || clean.include?("\r")
  clean
end

def safe_remote_basename(path)
  base = File.basename(path.to_s)
  base = "uploaded.exe" if base.strip.empty?
  base = "#{base}.exe" unless base.downcase.end_with?(".exe")
  base.gsub(/[^A-Za-z0-9._-]/, "_")
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

def scp_command(local_path, remote_host, remote_path)
  [
    "scp",
    "-q",
    "-o", "BatchMode=yes",
    "-o", "ConnectTimeout=15",
    "-o", "ServerAliveInterval=15",
    "-o", "ServerAliveCountMax=4",
    local_path.to_s,
    "#{remote_host}:#{remote_path}"
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
  stderr = [stderr, "q4 external Windows app run operation timed out after #{timeout_seconds}s"].reject(&:empty?).join("\n") if timed_out
  [stdout, stderr, timed_out ? 124 : status.exitstatus]
end

def bool(payload, key)
  payload[key] == true
end

def emit_json(payload, output_path)
  text = JSON.pretty_generate(payload) + "\n"
  FileUtils.mkdir_p(output_path.dirname)
  File.write(output_path, text)
  puts text
end

execute = options.fetch(:execute)
local_executable = ensure_local_executable_path!(options.fetch(:executable), execute: execute)
remote_host = options.fetch(:remote_host)
remote_materials_root = ensure_remote_xnix_path!("remote materials root", options.fetch(:remote_materials_root))
remote_external_root = "#{remote_materials_root}/external"
app_id = ensure_non_empty_single_line!("app id", options.fetch(:app_id))
display_name = ensure_non_empty_single_line!("display name", options.fetch(:display_name))
window_match = ensure_non_empty_single_line!("window match", options.fetch(:window_match))
output_path = ensure_local_output_path!("output path", options.fetch(:output))
markdown_output_path = ensure_local_output_path!("markdown output path", options.fetch(:markdown_output))
local_sha256 = execute ? Digest::SHA256.file(local_executable).hexdigest : ""
remote_name_prefix = execute ? local_sha256[0, 16] : "planned"
remote_executable = "#{remote_external_root}/#{remote_name_prefix}-#{safe_remote_basename(local_executable)}"

delegated_command = [
  "ruby",
  "scripts/q4_staged_desktop_external_winapp_smoke.rb",
  "--execute",
  "--fixture", "external",
  "--remote-executable", remote_executable,
  "--app-id", app_id,
  "--display-name", display_name,
  "--window-match", window_match,
  "--remote", remote_host,
  "--remote-timeout-seconds", options.fetch(:remote_timeout_seconds).to_s,
  "--output", output_path.to_s,
  "--markdown-output", markdown_output_path.to_s
]

plan = {
  "schema_version" => SCHEMA_VERSION,
  "request_type" => REQUEST_TYPE,
  "version" => VERSION,
  "status" => execute ? "running" : "planned",
  "execute" => execute,
  "remote_host" => remote_host,
  "local_executable_configured" => true,
  "local_executable_path_exposed" => false,
  "local_executable_sha256" => local_sha256,
  "local_executable_mz_header_checked" => execute,
  "remote_materials_root" => remote_materials_root,
  "remote_external_root" => remote_external_root,
  "remote_executable_configured" => true,
  "remote_executable_path_exposed_only_for_operator" => true,
  "remote_upload_planned" => true,
  "remote_upload_completed" => false,
  "app_id" => app_id,
  "display_name" => display_name,
  "window_match" => window_match,
  "delegated_script" => "scripts/q4_staged_desktop_external_winapp_smoke.rb",
  "delegated_command" => delegated_command,
  "output_path" => output_path.to_s,
  "markdown_output_path" => markdown_output_path.to_s,
  "runtime_owned" => true,
  "go_runtime_backed" => true,
  "kde_policy_owner" => false,
  "q4_compile_required" => true,
  "host_compilation_avoided" => true,
  "full_smoke_required" => false,
  "host_root_modified" => false,
  "privileged_container_required" => false,
  "host_networking_required" => false,
  "docker_socket_mounted" => false,
  "broad_host_mount_required" => false
}

unless execute
  emit_json(plan, output_path)
  exit 0
end

mkdir_stdout, mkdir_stderr, mkdir_status = run_shell(
  options.fetch(:local_shell),
  shell_join(ssh_command(remote_host, shell_join(["mkdir", "-p", remote_external_root]))),
  timeout_seconds: options.fetch(:remote_timeout_seconds)
)
unless mkdir_status.zero?
  warn mkdir_stdout unless mkdir_stdout.empty?
  warn mkdir_stderr unless mkdir_stderr.empty?
  abort "q4 external Windows app remote directory preparation failed"
end

scp_stdout, scp_stderr, scp_status = run_shell(
  options.fetch(:local_shell),
  shell_join(scp_command(local_executable, remote_host, remote_executable)),
  timeout_seconds: options.fetch(:remote_timeout_seconds)
)
unless scp_status.zero?
  warn scp_stdout unless scp_stdout.empty?
  warn scp_stderr unless scp_stderr.empty?
  abort "q4 external Windows app executable upload failed"
end

delegated_args = delegated_command.dup
delegated_args[1] = Q4_STAGED_EXTERNAL_SMOKE.to_s
stdout, stderr, status = Open3.capture3(*delegated_args, chdir: PROJECT_ROOT.to_s)
unless status.success?
  warn stdout unless stdout.empty?
  warn stderr unless stderr.empty?
  abort "q4 external Windows app staged external run failed"
end
delegated = JSON.parse(stdout)
unless delegated.fetch("status") == "passed" &&
       bool(delegated, "remote_executable_supplied") &&
       bool(delegated, "one_shot_runtime_launch_executed") &&
       bool(delegated, "windows_process_file_argument_window_observed") &&
       bool(delegated, "go_owned_q4_staged_external_winapp_acceptance_ready") &&
       delegated.fetch("accepted_application_detail_compatibility_state") == "runtime-accepted-real-app-run"
  warn stdout
  abort "q4 external Windows app staged external run did not reach accepted state"
end

result = plan.merge(
  "status" => "passed",
  "remote_upload_completed" => true,
  "delegated_status" => delegated.fetch("status"),
  "delegated_remote_build_completed" => bool(delegated, "remote_build_completed"),
  "delegated_remote_executable_supplied" => bool(delegated, "remote_executable_supplied"),
  "one_shot_status" => delegated.fetch("one_shot_status"),
  "one_shot_runtime_launch_executed" => bool(delegated, "one_shot_runtime_launch_executed"),
  "external_file_bridge_ready" => bool(delegated, "external_file_bridge_ready"),
  "windows_process_file_argument_window_observed" => bool(delegated, "windows_process_file_argument_window_observed"),
  "artifact_fetch_count" => delegated.fetch("artifact_fetch_count"),
  "acceptance_ready" => bool(delegated, "go_owned_q4_staged_external_winapp_acceptance_ready"),
  "accepted_application_detail_state" => delegated.fetch("accepted_application_detail_compatibility_state"),
  "accepted_application_detail_label" => delegated.fetch("accepted_application_detail_compatibility_label"),
  "kde_accepted_page_state" => delegated.fetch("kde_page_from_accepted_application_detail_compatibility_state"),
  "kde_accepted_page_label" => delegated.fetch("kde_page_from_accepted_application_detail_compatibility_label")
)
emit_json(result, output_path)
