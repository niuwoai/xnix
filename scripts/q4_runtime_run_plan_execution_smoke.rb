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
REMOTE_GO_BUILD = PROJECT_ROOT.join("scripts/remote_go_build.rb")
Q4_DBUS_FIXTURE = PROJECT_ROOT.join("scripts/q4_dbus_controlled_launch_owner_fixture_smoke.rb")
DEFAULT_REMOTE_HOST = ENV.fetch("XNIX_REMOTE_HOST", "root@q4")
DEFAULT_LOCAL_SHELL = ENV.fetch("XNIX_LOCAL_SHELL", "/bin/zsh")
DEFAULT_REMOTE_SOURCE_ROOT = ENV.fetch("XNIX_Q4_RUNTIME_SOURCE_ROOT", "/home/xnix-build/xnix-remote-go-build-runtime-#{VERSION}")
DEFAULT_REMOTE_BUILD_ROOT = ENV.fetch("XNIX_REMOTE_BUILD_ROOT", "/home/xnix-build-cache")
DEFAULT_HOST_GOOS = ENV.fetch("XNIX_Q4_RUNTIME_HOST_GOOS", "darwin")
DEFAULT_HOST_GOARCH = ENV.fetch("XNIX_Q4_RUNTIME_HOST_GOARCH", "arm64")
DEFAULT_FETCH_ROOT = ENV.fetch("XNIX_Q4_RUNTIME_FETCH_ROOT", "/tmp/xnix-runtime-go-host-#{VERSION}")
DEFAULT_TIMEOUT_SECONDS = Integer(ENV.fetch("XNIX_Q4_RUNTIME_RUN_PLAN_TIMEOUT_SECONDS", "900"), 10)
MESSAGEBOX_APP_ID = "org.xnix.apps.messagebox"

options = {
  execute: false,
  local_shell: DEFAULT_LOCAL_SHELL,
  remote_host: DEFAULT_REMOTE_HOST,
  remote_source_root: DEFAULT_REMOTE_SOURCE_ROOT,
  remote_build_root: DEFAULT_REMOTE_BUILD_ROOT,
  host_goos: DEFAULT_HOST_GOOS,
  host_goarch: DEFAULT_HOST_GOARCH,
  fetch_root: DEFAULT_FETCH_ROOT,
  timeout_seconds: DEFAULT_TIMEOUT_SECONDS,
  output: ENV.fetch("XNIX_Q4_RUNTIME_RUN_PLAN_OUTPUT", "")
}

OptionParser.new do |parser|
  parser.banner = "Usage: ruby scripts/q4_runtime_run_plan_execution_smoke.rb [--execute]"
  parser.on("--execute", "Build a host Runtime binary on q4, fetch it, and execute MessageBox from the verified catalog through Go Runtime.") { options[:execute] = true }
  parser.on("--local-shell PATH", "Local shell used for ssh/rsync alias resolution.") { |value| options[:local_shell] = value }
  parser.on("--remote HOST", "Remote SSH target, default: #{DEFAULT_REMOTE_HOST}.") { |value| options[:remote_host] = value }
  parser.on("--remote-source-root PATH", "Remote source root under /home/xnix-* or /tmp/xnix-*.") { |value| options[:remote_source_root] = value }
  parser.on("--remote-build-root PATH", "Remote build cache root under /home/xnix-* or /tmp/xnix-*.") { |value| options[:remote_build_root] = value }
  parser.on("--host-goos GOOS", "Host binary GOOS compiled on q4, default: #{DEFAULT_HOST_GOOS}.") { |value| options[:host_goos] = value }
  parser.on("--host-goarch GOARCH", "Host binary GOARCH compiled on q4, default: #{DEFAULT_HOST_GOARCH}.") { |value| options[:host_goarch] = value }
  parser.on("--fetch-root PATH", "Local fetch root under /tmp/xnix-* for the q4-built Runtime binary.") { |value| options[:fetch_root] = value }
  parser.on("--timeout-seconds SECONDS", Integer, "Timeout for q4 build, fetch, and Go Runtime execution.") { |value| options[:timeout_seconds] = value }
  parser.on("--output PATH", "Write the planned or passed summary JSON under this checkout or /tmp/xnix-*.") { |value| options[:output] = value }
end.parse!

abort "q4 Runtime run-plan execution smoke does not accept positional arguments" unless ARGV.empty?

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

def ensure_fetch_root!(path)
  clean = Pathname.new(path).cleanpath
  return clean if clean.to_s.start_with?("/tmp/xnix-")

  abort "fetch root must stay under /tmp/xnix-*"
end

def emit_json(payload, output_path)
  text = JSON.pretty_generate(payload) + "\n"
  File.write(output_path, text) if output_path
  puts text
end

def run_command(argv, timeout_seconds:, chdir: PROJECT_ROOT)
  stdout = +""
  stderr = +""
  status = nil
  timed_out = false
  Open3.popen3(*argv, chdir: chdir.to_s, pgroup: true) do |_stdin, out, err, wait_thread|
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
  stderr = [stderr, "command timed out after #{timeout_seconds}s"].reject(&:empty?).join("\n") if timed_out
  [stdout, stderr, timed_out ? 124 : status.exitstatus]
end

def run_json_command(argv, timeout_seconds:, chdir: PROJECT_ROOT)
  stdout, stderr, status = run_command(argv, timeout_seconds: timeout_seconds, chdir: chdir)
  unless status.zero?
    warn stdout unless stdout.empty?
    warn stderr unless stderr.empty?
    abort "command failed: #{argv.first}"
  end
  JSON.parse(stdout)
end

def write_json(path, payload)
  File.write(path, JSON.pretty_generate(payload) + "\n")
end

def matrix_evidence_fixture
  app = lambda do |app_id, display_name, app_version|
    {
      "app_id" => app_id,
      "display_name" => display_name,
      "app_version" => app_version,
      "smoke_status" => "passed",
      "compatibility_state" => "real-qemu-wine-verified",
      "marker_observed" => true,
      "checksum_verified" => true,
      "qemu_executed" => true,
      "wine_executed" => true,
      "guest_started" => true,
      "guest_port_auto" => true,
      "raw_output_redacted" => true,
      "serial_log_evidence" => true,
      "report_evidence" => true,
      "runtime_owned" => true,
      "go_runtime_backed" => true,
      "kde_policy_owner" => false,
      "desktop_launch_enabled" => false,
      "backend_launch_enabled" => false,
      "backend_details_exposed" => false,
      "raw_output_exposed" => false,
      "remote_path_exposed" => false,
      "host_root_modified" => false
    }
  end
  {
    "schema_version" => "xnix.runtime.known_app_matrix_evidence_preview.v1",
    "request_type" => "known-app-matrix-evidence-preview",
    "source" => "remote-known-winapp-matrix-smoke+runtime-evidence-consumer",
    "runtime_method" => "PreviewKnownAppMatrixEvidence",
    "read_method" => "GetKnownAppMatrixEvidence",
    "matrix_status" => "passed",
    "matrix_report_consumed" => true,
    "matrix_report_path_exposed" => false,
    "matrix_report_output_written" => true,
    "app_count" => 2,
    "passed_count" => 2,
    "failed_count" => 0,
    "evidence_count" => 2,
    "passed_evidence_count" => 2,
    "failed_evidence_count" => 0,
    "qemu_executed_count" => 2,
    "wine_executed_count" => 2,
    "marker_observed_count" => 2,
    "checksum_verified_count" => 2,
    "raw_output_redacted_count" => 2,
    "serial_log_evidence_count" => 2,
    "guest_started_count" => 2,
    "guest_port_auto_count" => 2,
    "compatibility_center_projection_ready" => true,
    "kde_center_projection_ready" => true,
    "apps" => [
      app.call("7zr", "7-Zip standalone console executable", "26.02"),
      app.call("busybox-w32", "BusyBox-w32 standalone console executable", "current-2026-07-24")
    ],
    "runtime_owned" => true,
    "go_runtime_backed" => true,
    "kde_policy_owner" => false,
    "desktop_launch_enabled" => false,
    "backend_launch_enabled" => false,
    "action_execution_enabled" => false,
    "backend_details_exposed" => false,
    "raw_output_exposed" => false,
    "remote_path_exposed" => false,
    "host_root_modified" => false,
    "privileged_container_required" => false,
    "host_networking_required" => false,
    "docker_socket_mounted" => false,
    "broad_host_mount_required" => false
  }
end

def gui_evidence_packet_fixture
  {
    "version" => VERSION,
    "schema_version" => "xnix.runtime.real_winapp_gui_evidence_packet.v1",
    "request_type" => "real-winapp-gui-evidence-packet-preview",
    "packet_type" => "real-windows-app-gui-evidence",
    "source" => "wine-guest-gui-smoke+runtime-evidence-consumer+real-winapp-desktop-packet",
    "runtime_method" => "PreviewRealWinAppGUIEvidencePacket",
    "read_method" => "GetRealWinAppGUIEvidencePacket",
    "report_status" => "passed",
    "report_consumed" => true,
    "report_path_exposed" => false,
    "app_id" => MESSAGEBOX_APP_ID,
    "display_name" => "Xnix MessageBox",
    "app_version" => VERSION,
    "gui_app_name" => "xnix-messagebox-smoke.exe",
    "evidence_source" => "wine-guest-gui-smoke",
    "compatibility_state" => "owner-controlled-gui-qemu-wine-verified",
    "center_card_state" => "validated-owner-controlled-gui-runtime-run",
    "known_app_gui_evidence_count" => 1,
    "known_app_gui_evidence_verified_count" => 1,
    "known_app_smoke_evidence" => {
      "app_id" => MESSAGEBOX_APP_ID,
      "display_name" => "Xnix MessageBox",
      "app_version" => VERSION,
      "evidence_kind" => "known-application-gui-smoke",
      "evidence_source" => "wine-guest-gui-smoke",
      "smoke_status" => "passed",
      "x_window_observed" => true,
      "window_observed" => true,
      "compatibility_state" => "owner-controlled-gui-qemu-wine-verified",
      "center_card_state" => "validated-owner-controlled-gui-runtime-run",
      "owner_file_open_verified" => true,
      "owner_file_open_entrypoint_invoked" => true,
      "owner_delegated_file_argument_count" => 1,
      "owner_delegated_file_argument_copied_count" => 1,
      "owner_delegated_file_arguments_passed" => true,
      "owner_delegated_file_argument_winepath_translated" => true,
      "owner_delegated_file_argument_winepath_translated_count" => 1,
      "owner_delegated_raw_file_argument_path_exposed" => false,
      "owner_delegated_window_match_observed" => true,
      "marker_observed" => true,
      "execution_evidence_recorded" => true,
      "runtime_dispatch_verified" => true,
      "launch_authorization_required" => true,
      "desktop_launch_enabled" => false,
      "runtime_owned" => true,
      "kde_policy_owner" => false,
      "action_execution_enabled" => false,
      "backend_launch_enabled" => false,
      "host_root_modified" => false,
      "backend_details_exposed" => false,
      "raw_artifact_path_exposed" => false
    },
    "wineboot_invoked" => true,
    "x_window_observed" => true,
    "window_observed" => true,
    "x_window_child_count" => 1,
    "compatibility_center_projection_ready" => true,
    "kde_center_projection_ready" => true,
    "container_runtime_used" => false,
    "container_host_mount_count" => 0,
    "runtime_owned" => true,
    "go_runtime_backed" => true,
    "kde_policy_owner" => false,
    "desktop_launch_enabled" => false,
    "backend_launch_enabled" => false,
    "action_execution_enabled" => false,
    "backend_details_exposed" => false,
    "raw_output_exposed" => false,
    "host_root_modified" => false,
    "privileged_container_required" => false,
    "host_networking_required" => false,
    "docker_socket_mounted" => false,
    "broad_host_mount_required" => false,
    "desktop_safe_summary" => "Xnix MessageBox produced owner-controlled GUI evidence."
  }
end

def require_passed_runtime_execution!(result)
  required_true = %w[
    run_plan_consumed
    runtime_owned_action_ready
    desktop_forwards_only_app_id
    gui_evidence_required
    gui_evidence_consumed
    window_observation_required
    window_observed
    window_match_observed
    owner_file_open_required
    owner_file_open_verified
    owner_file_open_entrypoint_invoked
    document_content_marker_observed
    execution_requested
    execution_started
    execution_completed
    smoke_report_consumed
    smoke_passed
    verified_catalog_consumed
    run_plan_generated
    actual_windows_app_run_observed
    compatibility_center_projection_ready
    kde_center_projection_ready
    go_owned_q4_winapp_acceptance_ready
    go_owned_q4_winapp_acceptance_consumed
    kde_controlled_launch_action_preview_ready
    host_compilation_avoided
    targeted_remote_verification_ready
    runtime_owned
    go_runtime_backed
    review_only
  ]
  required_true.each do |key|
    abort "Runtime run-plan execution missing #{key}" unless result.fetch(key) == true
  end
  required_false = %w[
    desktop_receipt_fields_reconstructed
    desktop_kde_state_root_access
    desktop_owner_inputs_exposed
    direct_run_plan_input
    go_owned_q4_winapp_acceptance_path_exposed
    host_compilation_required
    kde_policy_owner
    direct_launch_enabled
    launch_enabled
    desktop_files_written
    host_root_modified
    backend_launch_enabled
    backend_details_exposed
    raw_output_exposed
    remote_path_exposed
    smoke_command_arguments_exposed
    privileged_container_required
    host_networking_required
    docker_socket_mounted
    broad_host_mount_required
  ]
  required_false.each do |key|
    abort "Runtime run-plan execution opened unsafe gate #{key}" unless result.fetch(key) == false
  end
end

def require_desktop_consumption!(center, kde)
  required_center = {
    "known_app_smoke_evidence_count" => 1,
    "known_app_smoke_passed_count" => 1,
    "action_execution_enabled" => false,
    "backend_launch_enabled" => false,
    "host_root_modified" => false,
    "backend_details_exposed" => false
  }
  required_center.each do |key, expected|
    abort "Compatibility Center consumption mismatch for #{key}" unless center.fetch(key) == expected
  end

  required_kde = {
    "known_app_gui_evidence_count" => 1,
    "known_app_owner_controlled_gui_evidence_count" => 1,
    "launch_enabled" => false,
    "execution_started" => false,
    "backend_process_started" => false,
    "host_root_modified" => false,
    "backend_details_exposed" => false
  }
  required_kde.each do |key, expected|
    abort "KDE Center consumption mismatch for #{key}" unless kde.fetch(key) == expected
  end

  cards = kde.fetch("known_app_gui_evidence_cards")
  abort "KDE Center did not expose exactly one GUI evidence card" unless cards.is_a?(Array) && cards.length == 1

  card = cards.first
  required_card = {
    "app_id" => MESSAGEBOX_APP_ID,
    "primary_action_id" => "show-runtime-controlled-launch",
    "primary_action_kind" => "runtime-status",
    "desktop_launch_enabled" => false,
    "backend_launch_enabled" => false,
    "host_root_modified" => false,
    "backend_details_exposed" => false,
    "raw_artifact_path_exposed" => false
  }
  required_card.each do |key, expected|
    abort "KDE Center GUI evidence card mismatch for #{key}" unless card.fetch(key) == expected
  end
end

def require_owner_fixture_from_app_execution!(fixture)
  required = {
    "request_type" => "known-app-runtime-status-launch-owner-fixture-record",
    "app_id" => MESSAGEBOX_APP_ID,
    "display_name" => "Xnix MessageBox",
    "fixture_state" => "ready",
    "fixture_ready" => true,
    "desktop_trigger_ready" => true,
    "desktop_callable_runtime_method" => "ShowRuntimeControlledLaunch",
    "desktop_dbus_method" => "org.xnix.Compatibility1.ShowRuntimeControlledLaunch",
    "owner_service_call_ready" => true,
    "runtime_owner_service_supplies_inputs" => true,
    "runtime_owned" => true,
    "go_runtime_backed" => true,
    "kde_policy_owner" => false,
    "kde_forwards_only_evidence_handle" => true,
    "desktop_receipt_fields_reconstructed" => false,
    "desktop_kde_state_root_access" => false,
    "state_root_path_exposed" => false,
    "evidence_path_exposed" => false,
    "managed_launcher_path_exposed" => false,
    "raw_launcher_output_exposed" => false,
    "backend_details_exposed" => false,
    "host_root_modified" => false,
    "network_required" => false,
    "privileged_container_required" => false,
    "docker_socket_mounted" => false,
    "broad_host_mount_required" => false,
    "desktop_launch_enabled" => false,
    "backend_launch_enabled" => false,
    "execution_started" => false,
    "backend_process_started" => false
  }
  required.each do |key, expected|
    abort "owner fixture mismatch for #{key}" unless fixture.fetch(key) == expected
  end
  args = fixture.fetch("owner_service_call_args")
  unless args == ["ShowRuntimeControlledLaunch", "evidence-relative-path", fixture.fetch("evidence_relative_path")]
    abort "owner fixture must forward only the evidence-relative-path handle"
  end
end

def require_session_bus_plan_from_app_execution!(plan, fixture)
  required = {
    "request_type" => "kde-controlled-launch-session-bus-smoke-plan-preview",
    "kde_action_id" => "xnix.runtime-status.controlled-launch",
    "kde_action_preview_request_type" => "kde-controlled-launch-action-preview",
    "public_dbus_method" => "org.xnix.Compatibility1.ShowRuntimeControlledLaunch",
    "evidence_relative_path" => fixture.fetch("evidence_relative_path"),
    "evidence_handoff_consumed" => true,
    "evidence_digest_verified" => true,
    "kde_forwarded_argument_kind" => "evidence-relative-path",
    "kde_forwards_only_evidence_handle" => true,
    "restricted_session_bus_plan_ready" => true,
    "private_session_bus_required" => true,
    "dbus_session_bus_address_required" => true,
    "dbus_controlled_launch_fixture_plan_ready" => true,
    "runtime_owned" => true,
    "go_runtime_backed" => true,
    "kde_policy_owner" => false,
    "owner_service_args_exposed_to_kde" => false,
    "desktop_kde_state_root_access" => false,
    "desktop_receipt_fields_reconstructed" => false,
    "state_root_path_exposed" => false,
    "raw_launcher_output_exposed" => false,
    "backend_details_exposed" => false,
    "host_root_modified" => false,
    "docker_socket_mounted" => false,
    "broad_host_mount_required" => false,
    "privileged_container_required" => false,
    "host_network_required" => false,
    "execution_started" => false,
    "backend_process_started" => false,
    "smoke_executed_by_preview" => false
  }
  required.each do |key, expected|
    abort "session-bus plan mismatch for #{key}" unless plan.fetch(key) == expected
  end
  unless plan.fetch("kde_forwarded_arguments") == [fixture.fetch("evidence_relative_path")]
    abort "session-bus plan must forward only the owner fixture evidence handle"
  end
end

def require_q4_dbus_fixture!(fixture)
  required = {
    "request_type" => "q4-dbus-controlled-launch-owner-fixture-smoke",
    "status" => "passed",
    "app_id" => MESSAGEBOX_APP_ID,
    "app_execution_evidence_required" => true,
    "app_execution_evidence_copied_to_q4" => true,
    "app_execution_evidence_path_exposed" => false,
    "desktop_action_metadata_consumed" => true,
    "desktop_action_metadata_path_exposed" => false,
    "desktop_trigger_preflight_checked" => true,
    "desktop_trigger_preflight_blocked_missing_promotion" => true,
    "desktop_trigger_preflight_ready_for_operator_request" => true,
    "desktop_trigger_preflight_owner_service_shape_verified" => true,
    "desktop_trigger_preflight_path_exposed" => false,
    "linux_runtime_build_planned" => true,
    "linux_runtime_built_on_q4" => true,
    "dbus_adapter_compiled_in_q4_container" => true,
    "dbus_fixture_executed" => true,
    "dbus_fixture_passed" => true,
    "container_runtime" => "docker",
    "container_network_none" => true,
    "container_read_only" => true,
    "container_cap_drop_all" => true,
    "docker_socket_mounted" => false,
    "broad_host_mount_required" => false,
    "host_networking_required" => false,
    "privileged_container_required" => false,
    "host_compilation_avoided" => true,
    "host_root_modified" => false
  }
  required.each do |key, expected|
    abort "q4 D-Bus fixture mismatch for #{key}" unless fixture.fetch(key) == expected
  end
end

remote_source_root = ensure_remote_xnix_path!("remote source root", options.fetch(:remote_source_root))
remote_build_root = ensure_remote_xnix_path!("remote build root", options.fetch(:remote_build_root))
fetch_root = ensure_fetch_root!(options.fetch(:fetch_root))
output_path = options.fetch(:output).empty? ? nil : ensure_local_xnix_path!("output path", options.fetch(:output))
remote_binary = "#{remote_build_root}/bin/#{options.fetch(:host_goos)}-#{options.fetch(:host_goarch)}/xnix-runtime-go"

plan = {
  "schema_version" => "xnix.scripts.q4_runtime_run_plan_execution_smoke.v1",
  "request_type" => "q4-runtime-run-plan-execution-smoke",
  "version" => VERSION,
  "status" => options.fetch(:execute) ? "running" : "planned",
  "execute" => options.fetch(:execute),
  "remote_host" => options.fetch(:remote_host),
  "remote_runtime_build_planned" => true,
  "runtime_binary_goos" => options.fetch(:host_goos),
  "runtime_binary_goarch" => options.fetch(:host_goarch),
  "runtime_binary_built_on_q4" => false,
  "runtime_binary_fetched" => false,
  "runtime_binary_path_exposed" => false,
  "runtime_command" => "known-app-verified-catalog-app-execution",
  "runtime_command_execute_planned" => true,
  "verified_catalog_to_app_execution_planned" => true,
  "desktop_consumption_planned" => true,
  "run_plan_generated_by_runtime" => false,
  "direct_run_plan_input" => false,
  "go_runtime_entrypoint_invoked" => false,
  "app_id" => MESSAGEBOX_APP_ID,
  "q4_messagebox_smoke_planned" => true,
  "q4_messagebox_smoke_passed" => false,
  "actual_windows_app_run_observed" => false,
  "compatibility_center_projection_ready" => false,
  "kde_center_projection_ready" => false,
  "owner_file_open_entrypoint_invoked" => false,
  "document_content_marker_observed" => false,
  "go_owned_q4_winapp_acceptance_ready" => false,
  "compatibility_center_consumed" => false,
  "kde_center_page_consumed" => false,
  "known_app_smoke_evidence_count" => 0,
  "known_app_smoke_passed_count" => 0,
  "known_app_gui_evidence_count" => 0,
  "known_app_owner_controlled_gui_evidence_count" => 0,
  "known_app_gui_controlled_action_present" => false,
  "kde_controlled_launch_action_preview_ready" => false,
  "owner_fixture_from_app_execution_planned" => true,
  "owner_fixture_from_app_execution_ready" => false,
  "owner_fixture_desktop_trigger_ready" => false,
  "owner_fixture_service_call_ready" => false,
  "owner_fixture_kde_forwards_only_evidence_handle" => false,
  "owner_fixture_evidence_path_exposed" => false,
  "owner_fixture_backend_launch_enabled" => false,
  "owner_fixture_backend_process_started" => false,
  "session_bus_plan_from_app_execution_planned" => true,
  "session_bus_plan_from_app_execution_ready" => false,
  "session_bus_plan_private_bus_required" => false,
  "session_bus_plan_dbus_fixture_ready" => false,
  "session_bus_plan_kde_forwards_only_evidence_handle" => false,
  "session_bus_plan_backend_process_started" => false,
  "session_bus_plan_host_root_modified" => false,
  "q4_dbus_fixture_execute_planned" => true,
  "q4_dbus_fixture_passed" => false,
  "q4_dbus_fixture_linux_runtime_built_on_q4" => false,
  "q4_dbus_fixture_adapter_compiled_in_q4_container" => false,
  "q4_dbus_fixture_desktop_action_metadata_consumed" => false,
  "q4_dbus_fixture_desktop_action_metadata_path_exposed" => false,
  "q4_dbus_fixture_desktop_trigger_preflight_checked" => false,
  "q4_dbus_fixture_desktop_trigger_preflight_blocked_missing_promotion" => false,
  "q4_dbus_fixture_desktop_trigger_preflight_ready_for_operator_request" => false,
  "q4_dbus_fixture_desktop_trigger_preflight_owner_service_shape_verified" => false,
  "q4_dbus_fixture_desktop_trigger_preflight_path_exposed" => false,
  "q4_dbus_fixture_app_execution_evidence_copied_to_q4" => false,
  "q4_dbus_fixture_app_execution_evidence_path_exposed" => false,
  "q4_dbus_fixture_docker_socket_mounted" => false,
  "q4_dbus_fixture_broad_host_mount_required" => false,
  "q4_dbus_fixture_host_networking_required" => false,
  "q4_dbus_fixture_privileged_container_required" => false,
  "desktop_consumption_evidence_path_exposed" => false,
  "desktop_consumption_backend_launch_enabled" => false,
  "desktop_consumption_host_root_modified" => false,
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

build = run_json_command(
  [
    "ruby", REMOTE_GO_BUILD.to_s,
    "--execute",
    "--package", "./cmd/xnix-runtime-go",
    "--goos", options.fetch(:host_goos),
    "--goarch", options.fetch(:host_goarch),
    "--remote", options.fetch(:remote_host),
    "--remote-source-root", remote_source_root,
    "--remote-build-root", remote_build_root,
    "--remote-timeout-seconds", options.fetch(:timeout_seconds).to_s
  ],
  timeout_seconds: options.fetch(:timeout_seconds)
)
abort "q4 Runtime binary build did not pass" unless build.fetch("status") == "passed" && build.fetch("host_compilation_avoided") == true

FileUtils.mkdir_p(fetch_root)
local_runtime = fetch_root.join("xnix-runtime-go")
fetch_stdout, fetch_stderr, fetch_status = run_command(
  ["rsync", "-az", "--timeout", "60", "--contimeout", "15", "#{options.fetch(:remote_host)}:#{remote_binary}", local_runtime.to_s],
  timeout_seconds: options.fetch(:timeout_seconds)
)
unless fetch_status.zero?
  warn fetch_stdout unless fetch_stdout.empty?
  warn fetch_stderr unless fetch_stderr.empty?
  abort "q4 Runtime binary fetch failed"
end
File.chmod(0o700, local_runtime)

work_root = fetch_root.join("runtime-run-plan-execution")
FileUtils.mkdir_p(work_root)
matrix_path = work_root.join("known-app-matrix-evidence.json")
gui_packet_path = work_root.join("real-winapp-gui-evidence-packet.json")
catalog_path = work_root.join("known-app-verified-catalog.json")
app_execution_path = work_root.join("known-app-verified-catalog-app-execution-messagebox.json")
write_json(matrix_path, matrix_evidence_fixture)
write_json(gui_packet_path, gui_evidence_packet_fixture)

catalog = run_json_command(
  [local_runtime.to_s, "known-app-verified-catalog-preview", "--matrix-evidence", matrix_path.to_s, "--gui-evidence-packet", gui_packet_path.to_s],
  timeout_seconds: options.fetch(:timeout_seconds)
)
write_json(catalog_path, catalog)
result = run_json_command(
  [local_runtime.to_s, "known-app-verified-catalog-app-execution", "--verified-catalog", catalog_path.to_s, "--app", MESSAGEBOX_APP_ID, "--execute", "--workdir", PROJECT_ROOT.to_s, "--timeout-seconds", options.fetch(:timeout_seconds).to_s],
  timeout_seconds: options.fetch(:timeout_seconds)
)
require_passed_runtime_execution!(result)
write_json(app_execution_path, result)

center = run_json_command(
  [local_runtime.to_s, "compatibility-center-preview", "--registry", PROJECT_ROOT.join("runtime/recipes/registry.json").to_s, "--known-app-evidence-file", app_execution_path.to_s],
  timeout_seconds: options.fetch(:timeout_seconds)
)
kde = run_json_command(
  [local_runtime.to_s, "kde-center-page-preview", "--registry", PROJECT_ROOT.join("runtime/recipes/registry.json").to_s, "--app", MESSAGEBOX_APP_ID, "--decision", "approved", "--known-app-evidence-file", app_execution_path.to_s],
  timeout_seconds: options.fetch(:timeout_seconds)
)
require_desktop_consumption!(center, kde)

owner_state_root = work_root.join("owner-action-state")
FileUtils.mkdir_p(owner_state_root)
owner_fixture = run_json_command(
  [
    local_runtime.to_s,
    "known-app-runtime-status-launch-owner-fixture-record",
    "--app", MESSAGEBOX_APP_ID,
    "--state-root", owner_state_root.to_s,
    "--cache-root", work_root.join("known-winapps").to_s,
    "--gui-smoke-evidence-file", app_execution_path.to_s
  ],
  timeout_seconds: options.fetch(:timeout_seconds)
)
require_owner_fixture_from_app_execution!(owner_fixture)

session_bus_plan = run_json_command(
  [
    local_runtime.to_s,
    "kde-controlled-launch-session-bus-smoke-plan-preview",
    "--state-root", owner_state_root.to_s,
    "--evidence-relative-path", owner_fixture.fetch("evidence_relative_path")
  ],
  timeout_seconds: options.fetch(:timeout_seconds)
)
require_session_bus_plan_from_app_execution!(session_bus_plan, owner_fixture)

q4_dbus_fixture = run_json_command(
  [
    "ruby", Q4_DBUS_FIXTURE.to_s,
    "--execute",
    "--no-sync-source",
    "--remote", options.fetch(:remote_host),
    "--remote-source-root", remote_source_root,
    "--remote-build-root", remote_build_root,
    "--app", MESSAGEBOX_APP_ID,
    "--app-execution-file", app_execution_path.to_s,
    "--timeout-seconds", options.fetch(:timeout_seconds).to_s
  ],
  timeout_seconds: options.fetch(:timeout_seconds)
)
require_q4_dbus_fixture!(q4_dbus_fixture)

summary = plan.merge(
  "status" => "passed",
  "runtime_binary_built_on_q4" => true,
  "runtime_binary_fetched" => true,
  "runtime_command_request_type" => result.fetch("request_type"),
  "runtime_command_schema_version" => result.fetch("schema_version"),
  "verified_catalog_consumed" => result.fetch("verified_catalog_consumed"),
  "requested_app_id" => result.fetch("requested_app_id"),
  "run_plan_generated_by_runtime" => result.fetch("run_plan_generated"),
  "direct_run_plan_input" => result.fetch("direct_run_plan_input"),
  "go_runtime_entrypoint_invoked" => true,
  "runtime_execution_requested" => result.fetch("execution_requested"),
  "runtime_execution_completed" => result.fetch("execution_completed"),
  "runtime_smoke_request_type" => result.fetch("smoke_request_type"),
  "q4_messagebox_smoke_passed" => result.fetch("smoke_passed"),
  "actual_windows_app_run_observed" => result.fetch("actual_windows_app_run_observed"),
  "compatibility_center_projection_ready" => result.fetch("compatibility_center_projection_ready"),
  "kde_center_projection_ready" => result.fetch("kde_center_projection_ready"),
  "window_observed" => result.fetch("window_observed"),
  "window_match_observed" => result.fetch("window_match_observed"),
  "owner_file_open_entrypoint_invoked" => result.fetch("owner_file_open_entrypoint_invoked"),
  "document_content_marker_observed" => result.fetch("document_content_marker_observed"),
  "go_owned_q4_winapp_acceptance_ready" => result.fetch("go_owned_q4_winapp_acceptance_ready"),
  "go_owned_q4_winapp_acceptance_consumed" => result.fetch("go_owned_q4_winapp_acceptance_consumed"),
  "go_owned_q4_winapp_acceptance_path_exposed" => result.fetch("go_owned_q4_winapp_acceptance_path_exposed"),
  "kde_controlled_launch_action_preview_ready" => result.fetch("kde_controlled_launch_action_preview_ready"),
  "kde_controlled_launch_action_output_written" => result.fetch("kde_controlled_launch_action_output_written"),
  "kde_action_owner_file_open_verified" => result.fetch("kde_action_owner_file_open_verified"),
  "kde_action_owner_file_open_entrypoint_invoked" => result.fetch("kde_action_owner_file_open_entrypoint_invoked"),
  "kde_action_owner_file_open_environment_ready" => result.fetch("kde_action_owner_file_open_environment_ready"),
  "kde_action_owner_file_open_environment_values_exposed" => result.fetch("kde_action_owner_file_open_environment_values_exposed"),
  "owner_service_call_ready" => result.fetch("owner_service_call_ready"),
  "owner_evidence_handoff_ready" => result.fetch("owner_evidence_handoff_ready"),
  "compatibility_center_consumed" => true,
  "kde_center_page_consumed" => true,
  "known_app_smoke_evidence_count" => center.fetch("known_app_smoke_evidence_count"),
  "known_app_smoke_passed_count" => center.fetch("known_app_smoke_passed_count"),
  "known_app_gui_evidence_count" => kde.fetch("known_app_gui_evidence_count"),
  "known_app_owner_controlled_gui_evidence_count" => kde.fetch("known_app_owner_controlled_gui_evidence_count"),
  "known_app_gui_controlled_action_present" => kde.fetch("known_app_gui_evidence_cards").first.fetch("primary_action_id") == "show-runtime-controlled-launch",
  "owner_fixture_from_app_execution_ready" => owner_fixture.fetch("fixture_ready"),
  "owner_fixture_desktop_trigger_ready" => owner_fixture.fetch("desktop_trigger_ready"),
  "owner_fixture_service_call_ready" => owner_fixture.fetch("owner_service_call_ready"),
  "owner_fixture_kde_forwards_only_evidence_handle" => owner_fixture.fetch("kde_forwards_only_evidence_handle"),
  "owner_fixture_evidence_path_exposed" => owner_fixture.fetch("evidence_path_exposed"),
  "owner_fixture_backend_launch_enabled" => owner_fixture.fetch("backend_launch_enabled"),
  "owner_fixture_backend_process_started" => owner_fixture.fetch("backend_process_started"),
  "session_bus_plan_from_app_execution_ready" => session_bus_plan.fetch("restricted_session_bus_plan_ready"),
  "session_bus_plan_private_bus_required" => session_bus_plan.fetch("private_session_bus_required"),
  "session_bus_plan_dbus_fixture_ready" => session_bus_plan.fetch("dbus_controlled_launch_fixture_plan_ready"),
  "session_bus_plan_kde_forwards_only_evidence_handle" => session_bus_plan.fetch("kde_forwards_only_evidence_handle"),
  "session_bus_plan_backend_process_started" => session_bus_plan.fetch("backend_process_started"),
  "session_bus_plan_host_root_modified" => session_bus_plan.fetch("host_root_modified"),
  "q4_dbus_fixture_passed" => q4_dbus_fixture.fetch("dbus_fixture_passed"),
  "q4_dbus_fixture_linux_runtime_built_on_q4" => q4_dbus_fixture.fetch("linux_runtime_built_on_q4"),
  "q4_dbus_fixture_adapter_compiled_in_q4_container" => q4_dbus_fixture.fetch("dbus_adapter_compiled_in_q4_container"),
  "q4_dbus_fixture_desktop_action_metadata_consumed" => q4_dbus_fixture.fetch("desktop_action_metadata_consumed"),
  "q4_dbus_fixture_desktop_action_metadata_path_exposed" => q4_dbus_fixture.fetch("desktop_action_metadata_path_exposed"),
  "q4_dbus_fixture_desktop_trigger_preflight_checked" => q4_dbus_fixture.fetch("desktop_trigger_preflight_checked"),
  "q4_dbus_fixture_desktop_trigger_preflight_blocked_missing_promotion" => q4_dbus_fixture.fetch("desktop_trigger_preflight_blocked_missing_promotion"),
  "q4_dbus_fixture_desktop_trigger_preflight_ready_for_operator_request" => q4_dbus_fixture.fetch("desktop_trigger_preflight_ready_for_operator_request"),
  "q4_dbus_fixture_desktop_trigger_preflight_owner_service_shape_verified" => q4_dbus_fixture.fetch("desktop_trigger_preflight_owner_service_shape_verified"),
  "q4_dbus_fixture_desktop_trigger_preflight_path_exposed" => q4_dbus_fixture.fetch("desktop_trigger_preflight_path_exposed"),
  "q4_dbus_fixture_app_execution_evidence_copied_to_q4" => q4_dbus_fixture.fetch("app_execution_evidence_copied_to_q4"),
  "q4_dbus_fixture_app_execution_evidence_path_exposed" => q4_dbus_fixture.fetch("app_execution_evidence_path_exposed"),
  "q4_dbus_fixture_docker_socket_mounted" => q4_dbus_fixture.fetch("docker_socket_mounted"),
  "q4_dbus_fixture_broad_host_mount_required" => q4_dbus_fixture.fetch("broad_host_mount_required"),
  "q4_dbus_fixture_host_networking_required" => q4_dbus_fixture.fetch("host_networking_required"),
  "q4_dbus_fixture_privileged_container_required" => q4_dbus_fixture.fetch("privileged_container_required"),
  "desktop_consumption_evidence_path_exposed" => false,
  "desktop_consumption_backend_launch_enabled" => center.fetch("backend_launch_enabled") || kde.fetch("backend_process_started"),
  "desktop_consumption_host_root_modified" => center.fetch("host_root_modified") || kde.fetch("host_root_modified"),
  "runtime_remote_path_exposed" => result.fetch("remote_path_exposed"),
  "runtime_raw_output_exposed" => result.fetch("raw_output_exposed"),
  "runtime_smoke_command_arguments_exposed" => result.fetch("smoke_command_arguments_exposed"),
  "host_compilation_avoided" => build.fetch("host_compilation_avoided") && result.fetch("host_compilation_avoided"),
  "host_root_modified" => result.fetch("host_root_modified"),
  "privileged_container_required" => result.fetch("privileged_container_required"),
  "host_networking_required" => result.fetch("host_networking_required"),
  "docker_socket_mounted" => result.fetch("docker_socket_mounted"),
  "broad_host_mount_required" => result.fetch("broad_host_mount_required")
)
emit_json(summary, output_path)
