#!/usr/bin/env ruby
# frozen_string_literal: true

require "json"
require "open3"
require "optparse"
require "pathname"
require "shellwords"

PROJECT_ROOT = Pathname.new(__dir__).join("..").realpath
VERSION = PROJECT_ROOT.join("VERSION").read.strip
REMOTE_GUI_SMOKE = PROJECT_ROOT.join("scripts/remote_wine_guest_gui_smoke.rb")

DEFAULT_REMOTE_HOST = ENV.fetch("XNIX_REMOTE_HOST", "root@q4")
DEFAULT_SAMPLE_FILE = ENV.fetch("XNIX_Q4_SAMPLE_NOTEPAD_FILE", "sample-document.txt")
DEFAULT_WINDOW_MATCH = ENV.fetch("XNIX_Q4_SAMPLE_NOTEPAD_WINDOW_MATCH", DEFAULT_SAMPLE_FILE)
DEFAULT_REMOTE_TIMEOUT_SECONDS = Integer(ENV.fetch("XNIX_Q4_SAMPLE_NOTEPAD_TIMEOUT_SECONDS", "420"), 10)

options = {
  execute: false,
  sync_source: true,
  remote_host: DEFAULT_REMOTE_HOST,
  sample_file: DEFAULT_SAMPLE_FILE,
  window_match: DEFAULT_WINDOW_MATCH,
  remote_timeout_seconds: DEFAULT_REMOTE_TIMEOUT_SECONDS,
  output: ENV.fetch("XNIX_Q4_SAMPLE_NOTEPAD_OUTPUT", "")
}

OptionParser.new do |parser|
  parser.banner = "Usage: ruby scripts/q4_sample_notepad_smoke.rb [--execute]"
  parser.on("--execute", "Run Sample Notepad on q4 through the owner-controlled file-open acceptance lane.") { options[:execute] = true }
  parser.on("--no-sync-source", "Forward --no-sync-source to the delegated q4 GUI smoke.") { options[:sync_source] = false }
  parser.on("--remote HOST", "Remote SSH target, default: #{DEFAULT_REMOTE_HOST}") { |value| options[:remote_host] = value }
  parser.on("--sample-file NAME", "Remote sample file name created under the delegated smoke state root.") { |value| options[:sample_file] = value }
  parser.on("--window-match TEXT", "X window title/text required for the Sample Notepad run.") { |value| options[:window_match] = value }
  parser.on("--remote-timeout-seconds SECONDS", Integer, "Timeout for delegated q4 smoke operations.") { |value| options[:remote_timeout_seconds] = value }
  parser.on("--output PATH", "Write the plan or passed result JSON under this checkout or /tmp/xnix-*.") { |value| options[:output] = value }
end.parse!

abort "q4 Sample Notepad smoke does not accept positional arguments" unless ARGV.empty?
abort "sample file must be a simple file name" unless File.basename(options.fetch(:sample_file)) == options.fetch(:sample_file)
abort "window match must not be empty" if options.fetch(:window_match).strip.empty?

def ensure_local_output_path!(path)
  clean = Pathname.new(path).expand_path(PROJECT_ROOT).cleanpath
  return nil if path.to_s.strip.empty?
  return clean if clean.to_s.start_with?(PROJECT_ROOT.to_s)
  return clean if clean.to_s.start_with?("/tmp/xnix-")

  abort "output path must stay under this checkout or /tmp/xnix-*"
end

def emit_json(payload, output_path)
  text = JSON.pretty_generate(payload) + "\n"
  File.write(output_path, text) if output_path
  puts text
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

def remote_q4_acceptance(remote_host, remote_runtime_bin, payload_json)
  remote_input = "/tmp/xnix-q4-sample-notepad-acceptance-#{VERSION}.json"
  writer = shell_join(["ruby", "-e", "File.write(ARGV.fetch(0), STDIN.read)", remote_input])
  _write_stdout, write_stderr, write_status = Open3.capture3(*ssh_command(remote_host, writer), stdin_data: payload_json, chdir: PROJECT_ROOT.to_s)
  unless write_status.success?
    warn write_stderr unless write_stderr.empty?
    abort "FAIL: q4 Sample Notepad Go acceptance input write failed"
  end

  acceptance_args = [
    remote_runtime_bin,
    "q4-sample-notepad-acceptance-preview",
    "--q4-sample-notepad-smoke", remote_input
  ]
  stdout, stderr, status = Open3.capture3(*ssh_command(remote_host, shell_join(acceptance_args)), chdir: PROJECT_ROOT.to_s)
  unless status.success?
    warn stdout unless stdout.empty?
    warn stderr unless stderr.empty?
    abort "FAIL: q4 Sample Notepad Go acceptance preview failed"
  end
  JSON.parse(stdout)
end

output_path = ensure_local_output_path!(options.fetch(:output))

delegated_command = [
  "ruby",
  "scripts/remote_wine_guest_gui_smoke.rb",
  "--launch-mode", "owner-controlled-launch",
  "--file-open-entrypoint",
  "--known-app-id", "org.xnix.sample.notepad",
  "--sample-file-argument", options.fetch(:sample_file),
  "--window-match", options.fetch(:window_match),
  "--evidence-app-id", "org.xnix.sample.notepad",
  "--evidence-display-name", "Sample Notepad",
  "--remote", options.fetch(:remote_host),
  "--remote-timeout-seconds", options.fetch(:remote_timeout_seconds).to_s
]
delegated_command << "--no-sync-source" unless options.fetch(:sync_source)
delegated_command << "--execute" if options.fetch(:execute)
delegated_args = delegated_command.dup
delegated_args[1] = REMOTE_GUI_SMOKE.to_s

plan = {
  "schema_version" => "xnix.scripts.q4_sample_notepad_smoke.v1",
  "request_type" => "q4-sample-notepad-smoke",
  "version" => VERSION,
  "status" => options.fetch(:execute) ? "running" : "planned",
  "execute" => options.fetch(:execute),
  "remote_host" => options.fetch(:remote_host),
  "delegated_script" => "scripts/remote_wine_guest_gui_smoke.rb",
  "delegated_command" => delegated_command,
  "output_path" => output_path ? output_path.to_s : "",
  "source_sync_planned" => options.fetch(:sync_source),
  "app_id" => "org.xnix.sample.notepad",
  "display_name" => "Sample Notepad",
  "sample_file_argument" => options.fetch(:sample_file),
  "window_match" => options.fetch(:window_match),
  "launch_mode" => "owner-controlled-launch",
  "file_open_entrypoint_requested" => true,
  "real_run_acceptance_required" => true,
  "go_owned_q4_sample_notepad_acceptance_planned" => true,
  "go_owned_q4_sample_notepad_acceptance_ready" => false,
  "q4_compile_required" => true,
  "host_compilation_avoided" => true,
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

stdout, stderr, status = Open3.capture3(*delegated_args, chdir: PROJECT_ROOT.to_s)
unless status.success?
  warn stdout unless stdout.empty?
  warn stderr unless stderr.empty?
  warn "FAIL: q4 Sample Notepad owner-controlled file-open smoke failed"
  exit 1
end

delegated = JSON.parse(stdout)
acceptance_ready = delegated.fetch("real_run_acceptance_ready")
center_consumed = delegated.fetch("real_run_acceptance_center_projection_consumed")
kde_consumed = delegated.fetch("real_run_acceptance_kde_page_projection_consumed")
unless delegated.fetch("status") == "passed" && acceptance_ready && center_consumed && kde_consumed
  warn stdout
  warn "FAIL: q4 Sample Notepad acceptance did not pass"
  exit 1
end

result = plan.merge(
  "status" => "passed",
  "delegated_execute_result_schema" => delegated.fetch("schema_version"),
  "delegated_execute_result_output_written" => delegated.fetch("execute_result_output_written"),
  "real_run_receipt_summary_ready" => delegated.fetch("real_run_receipt_summary_ready"),
  "real_run_receipt_summary_file_open_verified" => delegated.fetch("real_run_receipt_summary_file_open_verified"),
  "real_run_acceptance_output_written" => delegated.fetch("real_run_acceptance_output_written"),
  "real_run_acceptance_ready" => acceptance_ready,
  "real_run_acceptance_center_projection_consumed" => center_consumed,
  "real_run_acceptance_kde_page_projection_consumed" => kde_consumed,
  "window_match_observed" => delegated.fetch("window_match_observed"),
  "owner_file_open_entrypoint_invoked" => delegated.fetch("owner_file_open_entrypoint_invoked"),
  "runtime_evidence_owner_file_open_entrypoint_invoked" => delegated.fetch("runtime_evidence_owner_file_open_entrypoint_invoked"),
  "kde_page_known_app_owner_file_open_entrypoint_count" => delegated.fetch("kde_page_known_app_owner_file_open_entrypoint_count"),
  "host_root_modified" => delegated.fetch("host_root_modified"),
  "privileged_container_required" => delegated.fetch("privileged_container_required"),
  "host_networking_required" => delegated.fetch("host_networking_required"),
  "docker_socket_mounted" => delegated.fetch("docker_socket_mounted"),
  "broad_host_mount_required" => delegated.fetch("broad_host_mount_required")
)

go_acceptance = remote_q4_acceptance(
  options.fetch(:remote_host),
  delegated.fetch("remote_runtime_bin"),
  JSON.pretty_generate(result) + "\n"
)
unless go_acceptance.fetch("acceptance_ready")
  warn JSON.pretty_generate(go_acceptance)
  warn "FAIL: q4 Sample Notepad Go-owned acceptance did not pass"
  exit 1
end

result.merge!(
  "go_owned_q4_sample_notepad_acceptance_schema" => go_acceptance.fetch("schema_version"),
  "go_owned_q4_sample_notepad_acceptance_request_type" => go_acceptance.fetch("request_type"),
  "go_owned_q4_sample_notepad_acceptance_ready" => go_acceptance.fetch("acceptance_ready"),
  "go_owned_q4_sample_notepad_acceptance_consumed" => go_acceptance.fetch("smoke_report_consumed"),
  "go_owned_q4_sample_notepad_acceptance_path_exposed" => go_acceptance.fetch("smoke_report_path_exposed"),
  "go_owned_q4_sample_notepad_acceptance_remote_host_exposed" => go_acceptance.fetch("remote_host_exposed"),
  "go_owned_q4_sample_notepad_acceptance_delegated_command_exposed" => go_acceptance.fetch("delegated_command_exposed")
)

emit_json(result, output_path)
