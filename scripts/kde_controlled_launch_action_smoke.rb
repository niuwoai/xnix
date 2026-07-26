#!/usr/bin/env ruby
# frozen_string_literal: true

require "fileutils"
require "json"
require "open3"
require "pathname"
require "securerandom"

PROJECT_ROOT = Pathname.new(__dir__).join("..").realpath
VERSION = PROJECT_ROOT.join("VERSION").read.strip
SMOKE_NAME = "KDE controlled launch action smoke"
PASS_MARKER = "PASS: KDE controlled launch action smoke"
SKIP_MARKER = "SKIP: KDE controlled launch action smoke"
APP_ID = "7zr"
GUI_APP_ID = "org.xnix.apps.messagebox"
GUI_APP_NAME = "Xnix MessageBox"
GUEST_BOUNDARY = "managed-known-app-guest-smoke"
RUN_ID = "#{Time.now.utc.strftime("%Y%m%d%H%M%S")}-#{Process.pid}-#{SecureRandom.hex(4)}"
WORK_ROOT = PROJECT_ROOT.join(".cache", "xnix", "kde-controlled-launch-action-smoke", RUN_ID)
STATE_ROOT = WORK_ROOT.join("state")
CACHE_ROOT = PROJECT_ROOT.join(".cache", "xnix", "known-winapps")
GO_CACHE_ROOT = PROJECT_ROOT.join(".cache", "go")
GO_TMP_ROOT = GO_CACHE_ROOT.join("tmp")
EXECUTE_ENV = "XNIX_KDE_CONTROLLED_LAUNCH_ACTION_SMOKE_EXECUTE"
EXECUTE_DBUS_FIXTURE_ENV = "XNIX_KDE_CONTROLLED_LAUNCH_ACTION_SMOKE_EXECUTE_DBUS_FIXTURE"

def assert(condition, message)
  return if condition

  warn "FAIL: #{message}"
  exit 1
end

def command_available?(name)
  ENV.fetch("PATH", "").split(File::PATH_SEPARATOR).any? do |directory|
    path = File.join(directory, name)
    File.file?(path) && File.executable?(path)
  end
end

def runtime_go_command
  return ["xnix-runtime-go"] if command_available?("xnix-runtime-go")
  return ["go", "run", "./cmd/xnix-runtime-go"] if command_available?("go")

  nil
end

def go_env
  {
    "GOCACHE" => GO_CACHE_ROOT.join("build").to_s,
    "GOMODCACHE" => GO_CACHE_ROOT.join("mod").to_s,
    "GOTMPDIR" => GO_TMP_ROOT.to_s
  }
end

def run_command(env, *argv)
  stdout, stderr, status = Open3.capture3(env, *argv, chdir: PROJECT_ROOT.to_s)
  [stdout, stderr, status.exitstatus]
end

def run_json(env, *argv)
  stdout, stderr, status = run_command(env, *argv)
  assert(status.zero?, "#{argv.join(" ")} must exit successfully: #{stderr}\n#{stdout}")
  [JSON.parse(stdout), stdout]
rescue JSON::ParserError => e
  warn stdout
  warn "FAIL: #{argv.join(" ")} must emit JSON: #{e.message}"
  exit 1
end

def assert_no_forbidden(text, forbidden_terms, label)
  downcased = text.downcase
  forbidden_terms.each do |term|
    next if term.to_s.empty?

    assert(!downcased.include?(term.to_s.downcase), "#{label} must not expose #{term}")
  end
end

def assert_false_payload(payload, fields, label)
  fields.each do |field|
    assert(payload[field] == false, "#{label} must keep #{field}=false")
  end
end

def messagebox_gui_projection
  session_id = "xnix-messagebox-controlled-session-#{VERSION}"
  {
    "projection_type" => "known-app-kde-runtime-status-launch-delegated-evidence",
    "runtime_method" => "ProjectKnownAppKDERuntimeStatusLaunchDelegatedEvidence",
    "request_type" => "windows-known-app-dispatch-smoke",
    "status" => "passed",
    "app_id" => GUI_APP_ID,
    "display_name" => GUI_APP_NAME,
    "app_version" => VERSION,
    "evidence_source" => "wine-guest-gui-smoke",
    "guest_boundary" => GUEST_BOUNDARY,
    "runtime_owned_dispatch" => true,
    "artifact_verified" => true,
    "marker_observed" => true,
    "session_gated_controlled_dispatch_consumed" => true,
    "session_gated_controlled_dispatch_state" => "created-after-session-gated-review",
    "session_gated_review_receipt_id" => "xnix-messagebox-session-review-#{VERSION}",
    "launch_authorization_receipt_id" => "xnix-messagebox-launch-authorization-#{VERSION}",
    "launch_authorization_receipt_state" => "recorded",
    "launch_gate_state" => "controlled-dispatch-ready",
    "launch_gate_consumed" => true,
    "launch_gate_receipt_accepted" => true,
    "launch_gate_guest_boundary_accepted" => true,
    "controlled_dispatch_ready" => true,
    "controlled_execution_session_consumed" => true,
    "controlled_execution_session_id" => session_id,
    "controlled_session_digest_verified" => true,
    "controlled_session_relative_path" => "execution-ledger/sessions/#{session_id}.json",
    "runtime_owner_consumable_session" => true,
    "kde_read_model_consumable_session" => true,
    "controlled_session_window_observed" => true,
    "compatibility_center_projection_ready" => true,
    "kde_center_projection_ready" => true,
    "state_root_path_exposed" => false,
    "managed_launcher_path_exposed" => false,
    "raw_launcher_output_exposed" => false,
    "backend_details_exposed" => false,
    "host_root_modified" => false,
    "docker_socket_mounted" => false,
    "broad_host_mount_required" => false
  }
end

def messagebox_owner_gui_report(evidence_relative_path)
  {
    "schema_version" => "xnix.scripts.wine_guest_gui_smoke.v1",
    "request_type" => "wine-guest-gui-smoke",
    "status" => "passed",
    "execute" => true,
    "gui_app_name" => "xnix-messagebox-smoke.exe",
    "local_gui_executable_configured" => true,
    "executable_copied" => true,
    "owner_controlled_launch_requested" => true,
    "owner_external_gui_app_requested" => true,
    "owner_external_gui_app_delivery" => "owner-managed-copy",
    "owner_seed_evidence_projected" => true,
    "owner_service_call_ready" => true,
    "owner_managed_launcher_invoked" => true,
    "owner_delegated_managed_artifact_copied" => true,
    "owner_delegated_smoke_passed" => true,
    "owner_delegated_evidence_source" => "wine-guest-gui-smoke",
    "owner_delegated_execution_started" => true,
    "owner_delegated_controlled_session_window_observed" => true,
    "owner_delegated_file_argument_count" => 1,
    "owner_delegated_file_argument_copied_count" => 1,
    "owner_delegated_file_arguments_passed" => true,
    "owner_delegated_file_argument_winepath_translated" => true,
    "owner_delegated_file_argument_winepath_translated_count" => 1,
    "owner_delegated_raw_file_argument_path_exposed" => false,
    "owner_delegated_window_match" => "messagebox-document.txt",
    "owner_delegated_window_match_observed" => true,
    "owner_delegated_host_root_modified" => false,
    "owner_delegated_docker_socket_mounted" => false,
    "owner_delegated_broad_host_mount_required" => false,
    "owner_delegated_raw_command_exposed" => false,
    "owner_delegated_backend_details_exposed" => false,
    "owner_evidence_handoff_ready" => true,
    "owner_evidence_relative_path" => evidence_relative_path,
    "runtime_go_owned_gui_smoke" => true,
    "runtime_payload_schema_version" => "xnix.runtime.windows_app_guest_wine_gui_smoke.v1",
    "wineboot_invoked" => true,
    "x_window_observed" => true,
    "x_window_child_count" => 3,
    "xwininfo_bytes" => 512,
    "guest_stderr_bytes" => 0,
    "guest_graphics_driver_error_observed" => false,
    "host_root_modified" => false,
    "privileged_container_required" => false,
    "host_networking_required" => false,
    "docker_socket_mounted" => false,
    "broad_host_mount_required" => false,
    "kde_safe_output_summary" => "MessageBox GUI smoke evidence observed a Runtime-owned X window."
  }
end

FileUtils.mkdir_p(STATE_ROOT)
FileUtils.mkdir_p(GO_CACHE_ROOT.join("build"))
FileUtils.mkdir_p(GO_CACHE_ROOT.join("mod"))
FileUtils.mkdir_p(GO_TMP_ROOT)

runtime_command = runtime_go_command
unless runtime_command
  puts "#{SKIP_MARKER} (xnix-runtime-go and go are unavailable)"
  exit 0
end

fixture, fixture_stdout = run_json(
  go_env,
  *runtime_command,
  "known-app-runtime-status-launch-owner-fixture-record",
  "--app", APP_ID,
  "--state-root", STATE_ROOT.to_s,
  "--cache-root", CACHE_ROOT.to_s,
  "--guest-boundary", GUEST_BOUNDARY
)

unless fixture["fixture_ready"] == true
  reason = fixture.fetch("skip_reason", "known Windows app artifact unavailable")
  puts "#{SKIP_MARKER} (#{reason})"
  exit 0
end

evidence_relative_path = fixture.fetch("evidence_relative_path")
assert(evidence_relative_path.start_with?("runtime/kde-runtime-status-launch-evidence/"), "fixture must expose only a relative evidence path")
assert(fixture["desktop_trigger_ready"] == true, "fixture must expose a desktop trigger")
assert(fixture["desktop_dbus_method"] == "org.xnix.Compatibility1.ShowRuntimeControlledLaunch", "fixture must target the controlled-launch D-Bus method")
assert(fixture["kde_forwards_only_evidence_handle"] == true, "fixture must keep KDE evidence-only")
assert_false_payload(
  fixture,
  %w[state_root_path_exposed raw_launcher_output_exposed backend_details_exposed host_root_modified docker_socket_mounted broad_host_mount_required privileged_container_required network_required],
  "fixture"
)
assert_no_forbidden(fixture_stdout, [STATE_ROOT.to_s, "wine ", "wine/", ".wine", "qemu-system", "program files"], "fixture output")

plan, plan_stdout = run_json(
  go_env,
  *runtime_command,
  "kde-controlled-launch-session-bus-smoke-plan-preview",
  "--state-root", STATE_ROOT.to_s,
  "--evidence-relative-path", evidence_relative_path
)

assert(plan["request_type"] == "kde-controlled-launch-session-bus-smoke-plan-preview", "plan request type must match")
assert(plan["kde_action_id"] == "xnix.runtime-status.controlled-launch", "plan must target the KDE controlled-launch action")
assert(plan["kde_action_preview_request_type"] == "kde-controlled-launch-action-preview", "plan must consume the KDE action preview")
assert(plan["public_dbus_method"] == "org.xnix.Compatibility1.ShowRuntimeControlledLaunch", "plan must target the controlled-launch D-Bus method")
assert(plan["evidence_relative_path"] == evidence_relative_path, "plan must preserve the fixture evidence path")
assert(plan["evidence_sha256"] == fixture.fetch("evidence_sha256"), "plan must preserve the fixture evidence digest")
assert(plan["evidence_handoff_consumed"] == true, "plan must consume the evidence handoff")
assert(plan["evidence_digest_verified"] == true, "plan must verify the evidence digest")
assert(plan["kde_forwarded_arguments"] == [evidence_relative_path], "plan must forward only the evidence handle")
assert(plan["kde_forwarded_argument_kind"] == "evidence-relative-path", "plan must label the forwarded evidence argument")
assert(plan["kde_forwards_only_evidence_handle"] == true, "plan must keep KDE evidence-only")
assert(plan["restricted_session_bus_plan_ready"] == true, "plan must be ready for the restricted session-bus smoke")
assert(plan["private_session_bus_required"] == true, "plan must require a private session bus")
assert(plan["dbus_session_bus_address_required"] == true, "plan must require DBUS_SESSION_BUS_ADDRESS")
assert(plan["outer_private_session_bus_command"] == ["dbus-run-session", "--", "ruby", "scripts/runtime_status_owner_service_session_bus_smoke.rb"], "plan must expose the private session-bus runner")
assert(plan["runtime_status_owner_session_smoke_command"] == ["ruby", "scripts/runtime_status_owner_service_session_bus_smoke.rb"], "plan must expose the Runtime-status owner session smoke command")
assert(plan["inner_staged_launcher_dispatch_command"] == ["ruby", "scripts/staged_launcher_dispatch_smoke.rb"], "plan must expose the staged launcher dispatch smoke command")
assert(plan["container_smoke_command"] == ["ruby", "scripts/container.rb", "runtime-status-owner-service-session-bus-smoke"], "plan must expose the restricted container smoke command")
assert(plan["dbus_controlled_launch_fixture_plan_ready"] == true, "plan must expose the D-Bus controlled-launch fixture lane")
assert(plan["dbus_controlled_launch_fixture_command"] == ["ruby", "scripts/dbus_controlled_launch_owner_fixture_smoke.rb"], "plan must expose the D-Bus controlled-launch fixture command")
assert(plan["dbus_controlled_launch_fixture_container_command"] == ["ruby", "scripts/container.rb", "dbus-controlled-launch-owner-fixture-smoke"], "plan must expose the restricted D-Bus fixture container command")
assert(plan["expected_pass_marker"] == "PASS: Runtime-status owner service session-bus smoke", "plan must expose the pass marker")
assert(plan["expected_skip_marker"] == "SKIP: Runtime-status owner service session-bus smoke", "plan must expose the skip marker")
assert(plan["expected_dbus_fixture_pass_marker"] == "PASS: D-Bus controlled launch owner fixture smoke", "plan must expose the D-Bus fixture pass marker")
assert(plan["expected_dbus_fixture_skip_marker"] == "SKIP: D-Bus controlled launch owner fixture smoke", "plan must expose the D-Bus fixture skip marker")
assert(plan["runtime_owned"] == true, "plan must keep Runtime ownership")
assert(plan["go_runtime_backed"] == true, "plan must be Go backed")
assert_false_payload(
  plan,
  %w[kde_policy_owner owner_service_args_exposed_to_kde desktop_kde_state_root_access desktop_receipt_fields_reconstructed state_root_path_exposed raw_launcher_output_exposed backend_details_exposed host_root_modified docker_socket_mounted broad_host_mount_required privileged_container_required host_network_required execution_started backend_process_started smoke_executed_by_preview],
  "plan"
)
assert_no_forbidden(plan_stdout, [STATE_ROOT.to_s, "owner_service_call_args", "wine ", "wine/", ".wine", "qemu-system", "program files"], "plan output")

runtime_status_projection_path = WORK_ROOT.join("messagebox-runtime-status-projection.json")
gui_report_path = WORK_ROOT.join("messagebox-owner-gui-report.json")
gui_evidence_path = WORK_ROOT.join("messagebox-owner-gui-evidence.json")
gui_page_path = WORK_ROOT.join("messagebox-kde-center-page.json")
File.write(runtime_status_projection_path, JSON.pretty_generate(messagebox_gui_projection))

gui_record, gui_record_stdout = run_json(
  go_env,
  *runtime_command,
  "known-app-kde-runtime-status-launch-evidence-record",
  "--state-root", STATE_ROOT.to_s,
  "--evidence-file", runtime_status_projection_path.to_s
)
gui_evidence_relative_path = gui_record.fetch("evidence_relative_path")
assert(gui_record["app_id"] == GUI_APP_ID, "GUI record must preserve the MessageBox app id")
assert(gui_record["known_app_smoke_evidence_ready"] == true, "GUI record must be center-page consumable")
assert(gui_evidence_relative_path.start_with?("runtime/kde-runtime-status-launch-evidence/"), "GUI record must expose only relative Runtime evidence")
assert_false_payload(
  gui_record,
  %w[state_root_path_exposed evidence_path_exposed managed_launcher_path_exposed raw_launcher_output_exposed backend_details_exposed host_root_modified docker_socket_mounted broad_host_mount_required desktop_launch_enabled backend_launch_enabled execution_started backend_process_started],
  "GUI record"
)
assert_no_forbidden(gui_record_stdout, [STATE_ROOT.to_s, runtime_status_projection_path.to_s, "owner_service_call_args", "wine/", ".wine", "qemu-system", "program files"], "GUI record output")

File.write(gui_report_path, JSON.pretty_generate(messagebox_owner_gui_report(gui_evidence_relative_path)))
gui_evidence, gui_evidence_stdout = run_json(
  go_env,
  *runtime_command,
  "gui-smoke-evidence-preview",
  "--gui-smoke-report", gui_report_path.to_s,
  "--app-id", GUI_APP_ID,
  "--display-name", GUI_APP_NAME,
  "--app-version", VERSION,
  "--output", gui_evidence_path.to_s
)
assert(gui_evidence["request_type"] == "gui-smoke-evidence-preview", "GUI evidence projection request type must match")
assert(gui_evidence["owner_controlled_launch_verified"] == true, "GUI evidence must verify owner-controlled launch")
assert(gui_evidence["owner_managed_copy_verified"] == true, "GUI evidence must verify owner-managed copy")
assert(gui_evidence["owner_file_open_verified"] == true, "GUI evidence must verify owner file-open handoff")
assert(gui_evidence["owner_delegated_window_match"] == "messagebox-document.txt", "GUI evidence must preserve safe owner window-match evidence")
assert(gui_evidence["owner_evidence_handoff_ready"] == true, "GUI evidence must preserve owner handoff readiness")
assert(gui_evidence["owner_evidence_relative_path"] == gui_evidence_relative_path, "GUI evidence must preserve the Runtime-status handoff")
assert(gui_evidence.dig("known_app_smoke_evidence", "primary_action_id") == "show-runtime-controlled-launch", "GUI evidence card must expose the Runtime-controlled action")
assert_false_payload(
  gui_evidence,
  %w[desktop_launch_enabled backend_launch_enabled action_execution_enabled backend_details_exposed raw_output_exposed remote_path_exposed host_root_modified privileged_container_required host_networking_required docker_socket_mounted broad_host_mount_required],
  "GUI evidence"
)
assert_no_forbidden(gui_evidence_stdout, [STATE_ROOT.to_s, gui_report_path.to_s, gui_evidence_path.to_s, "owner_service_call_args", "wine/", ".wine", "qemu-system", "program files"], "GUI evidence output")

gui_page, gui_page_stdout = run_json(
  go_env,
  *runtime_command,
  "kde-center-page-preview",
  "--registry", PROJECT_ROOT.join("runtime/recipes/registry.json").to_s,
  "--app", GUI_APP_ID,
  "--decision", "approved",
  "--known-app-evidence-file", gui_evidence_path.to_s
)
File.write(gui_page_path, JSON.pretty_generate(gui_page))
assert(gui_page["request_type"] == "kde-center-page-preview", "GUI page request type must match")
assert(gui_page["application_id"] == GUI_APP_ID, "GUI page must target MessageBox")
assert(gui_page["known_app_gui_evidence_count"] == 1, "GUI page must include one GUI evidence card")
assert(gui_page["known_app_owner_controlled_gui_evidence_count"] == 1, "GUI page must include owner-controlled GUI evidence")
assert(gui_page["known_app_owner_managed_copy_verified_count"] == 1, "GUI page must include owner-managed copy evidence")
assert(gui_page["known_app_owner_file_open_verified_count"] == 1, "GUI page must include owner file-open evidence")
assert(gui_page["known_app_gui_evidence_cards"].length == 1, "GUI page must carry one safe GUI card")
owner_gui_card = gui_page.fetch("known_app_gui_evidence_cards").first
assert(owner_gui_card.fetch("primary_action_id") == "show-runtime-controlled-launch", "owner GUI card must expose the Runtime-controlled action")
assert(owner_gui_card.fetch("desktop_callable_route") == "kde-dbus-runtime-status-action", "owner GUI card must expose the controlled desktop route")
assert(owner_gui_card.fetch("kde_forwarded_arguments") == [gui_evidence_relative_path], "owner GUI card must forward only the Runtime evidence handle")
assert(owner_gui_card.fetch("owner_file_open_verified") == true, "owner GUI card must expose owner file-open evidence")
assert(owner_gui_card.fetch("owner_delegated_window_match") == "messagebox-document.txt", "owner GUI card must expose safe owner window-match evidence")
assert(owner_gui_card.fetch("owner_service_args_exposed_to_kde") == false, "owner GUI card must hide owner service arguments")
assert_false_payload(
  gui_page,
  %w[launch_enabled execution_started backend_process_started host_root_modified backend_details_exposed],
  "GUI page"
)
assert_no_forbidden(gui_page_stdout, [STATE_ROOT.to_s, gui_evidence_path.to_s, "owner_service_call_args", "wine/", ".wine", "qemu-system", "program files"], "GUI page output")

gui_action, gui_action_stdout = run_json(
  go_env,
  *runtime_command,
  "kde-controlled-launch-action-preview",
  "--state-root", STATE_ROOT.to_s,
  "--kde-center-page-file", gui_page_path.to_s,
  "--app", GUI_APP_ID
)
assert(gui_action["request_type"] == "kde-controlled-launch-action-preview", "GUI page action request type must match")
assert(gui_action["application_id"] == GUI_APP_ID, "GUI page action must target MessageBox")
assert(gui_action["public_dbus_method"] == "org.xnix.Compatibility1.ShowRuntimeControlledLaunch", "GUI page action must target the controlled D-Bus method")
assert(gui_action["kde_forwarded_arguments"] == [gui_evidence_relative_path], "GUI page action must forward only the Runtime evidence handle")
assert(gui_action["desktop_callable_route"] == "kde-dbus-runtime-status-action", "GUI page action must preserve the controlled route")
assert(gui_action["kde_forwards_only_evidence_handle"] == true, "GUI page action must keep KDE evidence-only")
assert(gui_action["owner_file_open_verified"] == true, "GUI page action must expose owner file-open evidence")
assert(gui_action["owner_delegated_window_match"] == "messagebox-document.txt", "GUI page action must expose safe owner window-match evidence")
assert_false_payload(
  gui_action,
  %w[kde_policy_owner owner_service_args_exposed_to_kde kde_owns_owner_service_args desktop_kde_state_root_access desktop_receipt_fields_reconstructed state_root_path_exposed evidence_path_exposed managed_launcher_path_exposed raw_launcher_output_exposed backend_details_exposed host_root_modified docker_socket_mounted broad_host_mount_required privileged_container_required host_network_required desktop_launch_enabled backend_launch_enabled execution_started backend_process_started request_object_created_by_kde],
  "GUI page action"
)
assert_no_forbidden(gui_action_stdout, [STATE_ROOT.to_s, gui_page_path.to_s, gui_evidence_path.to_s, runtime_status_projection_path.to_s, "owner_service_call_args", "wine/", ".wine", "qemu-system", "program files"], "GUI page action output")

unless ENV.fetch(EXECUTE_ENV, "") == "1" || ENV.fetch(EXECUTE_DBUS_FIXTURE_ENV, "") == "1"
  puts "#{SKIP_MARKER} (validated Go-owned plan, GUI center-page handoff, and D-Bus fixture lane; set #{EXECUTE_ENV}=1 or #{EXECUTE_DBUS_FIXTURE_ENV}=1 to execute a restricted smoke)"
  exit 0
end

command = if ENV.fetch(EXECUTE_DBUS_FIXTURE_ENV, "") == "1"
            plan.fetch("dbus_controlled_launch_fixture_command")
          else
            plan.fetch("runtime_status_owner_session_smoke_command")
          end
expected_pass_marker = if ENV.fetch(EXECUTE_DBUS_FIXTURE_ENV, "") == "1"
                         plan.fetch("expected_dbus_fixture_pass_marker")
                       else
                         plan.fetch("expected_pass_marker")
                       end
expected_skip_marker = if ENV.fetch(EXECUTE_DBUS_FIXTURE_ENV, "") == "1"
                         plan.fetch("expected_dbus_fixture_skip_marker")
                       else
                         plan.fetch("expected_skip_marker")
                       end
stdout, stderr, status = run_command({}, *command)
print stdout
warn stderr unless stderr.empty?
assert(status.zero?, "#{command.join(" ")} must pass or skip cleanly")
assert_no_forbidden(stdout, [STATE_ROOT.to_s, "docker.sock", "--privileged", "--network host", "type=bind", "program files", ".wine", "qemu-system"], "KDE action smoke stdout")
assert_no_forbidden(stderr, [STATE_ROOT.to_s, "docker.sock", "--privileged", "--network host", "type=bind", "program files", ".wine", "qemu-system"], "KDE action smoke stderr")

case stdout
when /#{Regexp.escape(expected_pass_marker)}/
  puts PASS_MARKER
when /#{Regexp.escape(expected_skip_marker)}/
  reason = stdout.lines.find { |line| line.include?(expected_skip_marker) }.to_s.sub(expected_skip_marker, "").strip
  puts "#{SKIP_MARKER} #{reason}".rstrip
else
  warn stdout
  warn stderr unless stderr.empty?
  warn "FAIL: #{SMOKE_NAME} must report a plan-defined pass or skip marker"
  exit 1
end
