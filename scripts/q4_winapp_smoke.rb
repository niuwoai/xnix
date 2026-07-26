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
DEFAULT_REMOTE_TIMEOUT_SECONDS = Integer(ENV.fetch("XNIX_Q4_WINAPP_TIMEOUT_SECONDS", "420"), 10)

options = {
  execute: false,
  sync_source: true,
  remote_host: DEFAULT_REMOTE_HOST,
  known_app_id: ENV.fetch("XNIX_Q4_WINAPP_KNOWN_APP_ID", ""),
  remote_executable: ENV.fetch("XNIX_Q4_WINAPP_REMOTE_EXECUTABLE", ""),
  app_id: ENV.fetch("XNIX_Q4_WINAPP_APP_ID", ""),
  display_name: ENV.fetch("XNIX_Q4_WINAPP_DISPLAY_NAME", ""),
  remote_file_argument: ENV.fetch("XNIX_Q4_WINAPP_REMOTE_FILE_ARGUMENT", ""),
  sample_file_argument: ENV.fetch("XNIX_Q4_WINAPP_SAMPLE_FILE_ARGUMENT", ""),
  window_match: ENV.fetch("XNIX_Q4_WINAPP_WINDOW_MATCH", ""),
  owner_file_open: ENV.fetch("XNIX_Q4_WINAPP_OWNER_FILE_OPEN", "0") == "1",
  require_real_run_acceptance: ENV.fetch("XNIX_Q4_WINAPP_REQUIRE_REAL_RUN_ACCEPTANCE", "0") == "1",
  remote_timeout_seconds: DEFAULT_REMOTE_TIMEOUT_SECONDS,
  output: ENV.fetch("XNIX_Q4_WINAPP_OUTPUT", "")
}

OptionParser.new do |parser|
  parser.banner = "Usage: ruby scripts/q4_winapp_smoke.rb [--known-app-id ID | --remote-executable PATH] --window-match TEXT [--execute]"
  parser.on("--execute", "Run the q4 Windows GUI smoke.") { options[:execute] = true }
  parser.on("--no-sync-source", "Forward --no-sync-source to the delegated q4 GUI smoke.") { options[:sync_source] = false }
  parser.on("--remote HOST", "Remote SSH target, default: #{DEFAULT_REMOTE_HOST}") { |value| options[:remote_host] = value }
  parser.on("--known-app-id ID", "Known Windows GUI app id resolved by the Go Runtime.") { |value| options[:known_app_id] = value }
  parser.on("--remote-executable PATH", "Remote Windows GUI .exe under /home/xnix-* or /tmp/xnix-*.") { |value| options[:remote_executable] = value }
  parser.on("--app-id ID", "Evidence application id; defaults to --known-app-id or org.xnix.external.q4-winapp.") { |value| options[:app_id] = value }
  parser.on("--display-name NAME", "Evidence display name; defaults to the app id or remote executable basename.") { |value| options[:display_name] = value }
  parser.on("--remote-file-argument PATH", "Remote file under /home/xnix-* or /tmp/xnix-* copied into the Wine guest and passed to the app.") { |value| options[:remote_file_argument] = value }
  parser.on("--sample-file-argument NAME", "Create a remote sample file under the smoke state root and pass it to the app.") { |value| options[:sample_file_argument] = value }
  parser.on("--window-match TEXT", "Case-insensitive X window title/text required for GUI observation.") { |value| options[:window_match] = value }
  parser.on("--owner-file-open", "Route known-app or remote-executable runs through the owner-controlled xnix-compat-open file-open entrypoint.") { options[:owner_file_open] = true }
  parser.on("--require-real-run-acceptance", "Require the Go real-run acceptance summary to be ready.") { options[:require_real_run_acceptance] = true }
  parser.on("--remote-timeout-seconds SECONDS", Integer, "Timeout for delegated q4 smoke operations.") { |value| options[:remote_timeout_seconds] = value }
  parser.on("--output PATH", "Write the plan or passed result JSON under this checkout or /tmp/xnix-*.") { |value| options[:output] = value }
end.parse!

abort "q4 Windows app smoke does not accept positional arguments" unless ARGV.empty?

def ensure_local_output_path!(path)
  return nil if path.to_s.strip.empty?

  clean = Pathname.new(path).expand_path(PROJECT_ROOT).cleanpath
  return clean if clean.to_s.start_with?(PROJECT_ROOT.to_s)
  return clean if clean.to_s.start_with?("/tmp/xnix-")

  abort "output path must stay under this checkout or /tmp/xnix-*"
end

def ensure_optional_remote_xnix_path!(label, path)
  value = path.to_s.strip
  return "" if value.empty?

  clean = Pathname.new(value).cleanpath.to_s
  return clean if clean.start_with?("/home/xnix-") || clean.start_with?("/tmp/xnix-")

  abort "#{label} must stay under /home/xnix-* or /tmp/xnix-* on q4"
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
  remote_input = "/tmp/xnix-q4-winapp-acceptance-#{VERSION}.json"
  writer = shell_join(["ruby", "-e", "File.write(ARGV.fetch(0), STDIN.read)", remote_input])
  _write_stdout, write_stderr, write_status = Open3.capture3(*ssh_command(remote_host, writer), stdin_data: payload_json, chdir: PROJECT_ROOT.to_s)
  unless write_status.success?
    warn write_stderr unless write_stderr.empty?
    abort "FAIL: q4 Windows app Go acceptance input write failed"
  end

  acceptance_args = [
    remote_runtime_bin,
    "q4-winapp-acceptance-preview",
    "--q4-winapp-smoke", remote_input
  ]
  stdout, stderr, status = Open3.capture3(*ssh_command(remote_host, shell_join(acceptance_args)), chdir: PROJECT_ROOT.to_s)
  unless status.success?
    warn stdout unless stdout.empty?
    warn stderr unless stderr.empty?
    abort "FAIL: q4 Windows app Go acceptance preview failed"
  end
  JSON.parse(stdout)
end

def bool(payload, key)
  payload[key] == true
end

known_app_id = options.fetch(:known_app_id).strip
remote_executable = ensure_optional_remote_xnix_path!("remote executable", options.fetch(:remote_executable))
remote_file_argument = ensure_optional_remote_xnix_path!("remote file argument", options.fetch(:remote_file_argument))
sample_file_argument = options.fetch(:sample_file_argument).strip
window_match = options.fetch(:window_match).strip

abort "use exactly one of --known-app-id or --remote-executable" if known_app_id.empty? == remote_executable.empty?
abort "window match must not be empty" if window_match.empty?
abort "sample file argument must be a simple file name" if !sample_file_argument.empty? && File.basename(sample_file_argument) != sample_file_argument
abort "use either --remote-file-argument or --sample-file-argument, not both" if !remote_file_argument.empty? && !sample_file_argument.empty?
abort "--require-real-run-acceptance requires --owner-file-open" if options.fetch(:require_real_run_acceptance) && !options.fetch(:owner_file_open)

app_id = options.fetch(:app_id).strip
app_id = known_app_id unless known_app_id.empty? || !app_id.empty?
app_id = "org.xnix.external.q4-winapp" if app_id.empty?

display_name = options.fetch(:display_name).strip
display_name = app_id if display_name.empty? && !known_app_id.empty?
display_name = File.basename(remote_executable) if display_name.empty?

output_path = ensure_local_output_path!(options.fetch(:output))

launch_mode = options.fetch(:owner_file_open) ? "owner-controlled-launch" : "direct"
delegated_command = [
  "ruby",
  "scripts/remote_wine_guest_gui_smoke.rb",
  "--launch-mode", launch_mode,
  "--window-match", window_match,
  "--evidence-app-id", app_id,
  "--evidence-display-name", display_name,
  "--remote", options.fetch(:remote_host),
  "--remote-timeout-seconds", options.fetch(:remote_timeout_seconds).to_s
]
delegated_command.concat(["--known-app-id", known_app_id]) unless known_app_id.empty?
delegated_command.concat(["--remote-executable", remote_executable]) unless remote_executable.empty?
delegated_command.concat(["--remote-file-argument", remote_file_argument]) unless remote_file_argument.empty?
delegated_command.concat(["--sample-file-argument", sample_file_argument]) unless sample_file_argument.empty?
delegated_command << "--file-open-entrypoint" if options.fetch(:owner_file_open)
delegated_command << "--no-sync-source" unless options.fetch(:sync_source)
delegated_command << "--execute" if options.fetch(:execute)
delegated_args = delegated_command.dup
delegated_args[1] = REMOTE_GUI_SMOKE.to_s

plan = {
  "schema_version" => "xnix.scripts.q4_winapp_smoke.v1",
  "request_type" => "q4-winapp-smoke",
  "version" => VERSION,
  "status" => options.fetch(:execute) ? "running" : "planned",
  "execute" => options.fetch(:execute),
  "remote_host" => options.fetch(:remote_host),
  "delegated_script" => "scripts/remote_wine_guest_gui_smoke.rb",
  "delegated_command" => delegated_command,
  "output_path" => output_path ? output_path.to_s : "",
  "source_sync_planned" => options.fetch(:sync_source),
  "known_app_id" => known_app_id,
  "remote_executable_configured" => !remote_executable.empty?,
  "remote_executable_path_exposed" => false,
  "app_id" => app_id,
  "display_name" => display_name,
  "remote_file_argument_configured" => !remote_file_argument.empty?,
  "remote_file_argument_path_exposed" => false,
  "sample_file_argument" => sample_file_argument,
  "window_match" => window_match,
  "launch_mode" => launch_mode,
  "file_open_entrypoint_requested" => options.fetch(:owner_file_open),
  "real_run_acceptance_required" => options.fetch(:require_real_run_acceptance),
  "go_owned_q4_winapp_acceptance_planned" => true,
  "go_owned_q4_winapp_acceptance_ready" => false,
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
  warn "FAIL: q4 Windows app GUI smoke failed"
  exit 1
end

delegated = JSON.parse(stdout)
window_observed = bool(delegated, "window_match_observed") || bool(delegated, "x_window_observed") || bool(delegated, "window_observed")
unless delegated.fetch("status") == "passed" && window_observed
  warn stdout
  warn "FAIL: q4 Windows app smoke did not prove observed GUI execution"
  exit 1
end

if options.fetch(:require_real_run_acceptance)
  acceptance_ready = bool(delegated, "real_run_acceptance_ready")
  center_consumed = bool(delegated, "real_run_acceptance_center_projection_consumed")
  kde_consumed = bool(delegated, "real_run_acceptance_kde_page_projection_consumed")
  unless acceptance_ready && center_consumed && kde_consumed
    warn stdout
    warn "FAIL: q4 Windows app smoke did not satisfy real-run acceptance"
    exit 1
  end
end

result = plan.merge(
  "status" => "passed",
  "delegated_execute_result_schema" => delegated.fetch("schema_version"),
  "delegated_execute_result_output_written" => bool(delegated, "execute_result_output_written"),
  "remote_build_completed" => bool(delegated, "remote_build_completed"),
  "evidence_output_written" => bool(delegated, "evidence_output_written"),
  "kde_page_output_written" => bool(delegated, "kde_page_output_written"),
  "kde_action_output_written" => bool(delegated, "kde_action_output_written"),
  "window_observed" => window_observed,
  "window_match_observed" => bool(delegated, "window_match_observed"),
  "document_content_marker_observation_required" => bool(delegated, "document_content_marker_observation_required"),
  "document_content_marker_observed" => bool(delegated, "document_content_marker_observed"),
  "real_run_receipt_summary_ready" => bool(delegated, "real_run_receipt_summary_ready"),
  "real_run_receipt_summary_file_open_verified" => bool(delegated, "real_run_receipt_summary_file_open_verified"),
  "real_run_receipt_summary_document_content_marker_observation_required" => bool(delegated, "real_run_receipt_summary_document_content_marker_observation_required"),
  "real_run_receipt_summary_document_content_marker_observed" => bool(delegated, "real_run_receipt_summary_document_content_marker_observed"),
  "real_run_acceptance_output_written" => bool(delegated, "real_run_acceptance_output_written"),
  "real_run_acceptance_ready" => bool(delegated, "real_run_acceptance_ready"),
  "real_run_acceptance_document_content_marker_observation_required" => bool(delegated, "real_run_acceptance_document_content_marker_observation_required"),
  "real_run_acceptance_document_content_marker_observed" => bool(delegated, "real_run_acceptance_document_content_marker_observed"),
  "real_run_acceptance_center_projection_consumed" => bool(delegated, "real_run_acceptance_center_projection_consumed"),
  "real_run_acceptance_kde_page_projection_consumed" => bool(delegated, "real_run_acceptance_kde_page_projection_consumed"),
  "owner_file_open_entrypoint_invoked" => bool(delegated, "owner_file_open_entrypoint_invoked"),
  "runtime_evidence_owner_file_open_entrypoint_invoked" => bool(delegated, "runtime_evidence_owner_file_open_entrypoint_invoked"),
  "kde_page_known_app_owner_file_open_entrypoint_count" => delegated.fetch("kde_page_known_app_owner_file_open_entrypoint_count", 0),
  "host_root_modified" => bool(delegated, "host_root_modified"),
  "privileged_container_required" => bool(delegated, "privileged_container_required"),
  "host_networking_required" => bool(delegated, "host_networking_required"),
  "docker_socket_mounted" => bool(delegated, "docker_socket_mounted"),
  "broad_host_mount_required" => bool(delegated, "broad_host_mount_required")
)

go_acceptance = remote_q4_acceptance(
  options.fetch(:remote_host),
  delegated.fetch("remote_runtime_bin"),
  JSON.pretty_generate(result) + "\n"
)
unless go_acceptance.fetch("acceptance_ready")
  warn JSON.pretty_generate(go_acceptance)
  warn "FAIL: q4 Windows app Go-owned acceptance did not pass"
  exit 1
end

result.merge!(
  "go_owned_q4_winapp_acceptance_schema" => go_acceptance.fetch("schema_version"),
  "go_owned_q4_winapp_acceptance_request_type" => go_acceptance.fetch("request_type"),
  "go_owned_q4_winapp_acceptance_ready" => go_acceptance.fetch("acceptance_ready"),
  "go_owned_q4_winapp_acceptance_consumed" => go_acceptance.fetch("smoke_report_consumed"),
  "go_owned_q4_winapp_acceptance_document_content_marker_observation_required" => go_acceptance.fetch("document_content_marker_observation_required"),
  "go_owned_q4_winapp_acceptance_document_content_marker_observed" => go_acceptance.fetch("document_content_marker_observed"),
  "go_owned_q4_winapp_acceptance_path_exposed" => go_acceptance.fetch("smoke_report_path_exposed"),
  "go_owned_q4_winapp_acceptance_remote_host_exposed" => go_acceptance.fetch("remote_host_exposed"),
  "go_owned_q4_winapp_acceptance_delegated_command_exposed" => go_acceptance.fetch("delegated_command_exposed"),
  "go_owned_q4_winapp_acceptance_remote_executable_path_exposed" => go_acceptance.fetch("remote_executable_path_exposed"),
  "go_owned_q4_winapp_acceptance_remote_file_argument_path_exposed" => go_acceptance.fetch("remote_file_argument_path_exposed")
)

emit_json(result, output_path)
