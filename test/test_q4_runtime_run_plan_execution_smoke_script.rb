#!/usr/bin/env ruby
# frozen_string_literal: true

require "json"
require "open3"
require "pathname"

def assert(condition, message)
  return if condition

  warn "FAIL: #{message}"
  exit 1
end

project_root = Pathname.new(__dir__).join("..").realpath
script = project_root.join("scripts/q4_runtime_run_plan_execution_smoke.rb")
source = script.read

assert(source.include?("xnix.scripts.q4_runtime_run_plan_execution_smoke.v1"), "q4 Runtime run-plan execution smoke must expose a stable schema")
assert(source.include?("q4-runtime-run-plan-execution-smoke"), "q4 Runtime run-plan execution smoke must expose a request type")
assert(source.include?("scripts/remote_go_build.rb"), "q4 Runtime run-plan execution smoke must build the Runtime binary on q4")
assert(source.include?("--goos"), "q4 Runtime run-plan execution smoke must pass GOOS to q4 build")
assert(source.include?("darwin"), "q4 Runtime run-plan execution smoke must support q4 cross-compiled macOS Runtime binaries")
assert(source.include?("arm64"), "q4 Runtime run-plan execution smoke must default to Apple Silicon Runtime binaries")
assert(source.include?("known-app-verified-catalog-preview"), "q4 Runtime run-plan execution smoke must generate the verified catalog through Go Runtime")
assert(source.include?("known-app-verified-catalog-app-execution"), "q4 Runtime run-plan execution smoke must execute the app from the verified catalog through Go Runtime")
assert(source.include?("compatibility-center-preview"), "q4 Runtime run-plan execution smoke must consume app execution evidence through Compatibility Center")
assert(source.include?("kde-center-page-preview"), "q4 Runtime run-plan execution smoke must consume app execution evidence through KDE Center")
assert(source.include?("verified_catalog_to_app_execution_planned"), "q4 Runtime run-plan execution smoke must plan verified-catalog-to-app execution")
assert(source.include?("desktop_consumption_planned"), "q4 Runtime run-plan execution smoke must plan desktop read-model consumption")
assert(source.include?("compatibility_center_consumed"), "q4 Runtime run-plan execution smoke must report Compatibility Center consumption")
assert(source.include?("kde_center_page_consumed"), "q4 Runtime run-plan execution smoke must report KDE Center page consumption")
assert(source.include?("known_app_gui_controlled_action_present"), "q4 Runtime run-plan execution smoke must report Runtime-status controlled action evidence")
assert(source.include?("kde_controlled_launch_action_preview_ready"), "q4 Runtime run-plan execution smoke must require KDE controlled action preview readiness")
assert(source.include?("known-app-runtime-status-launch-owner-fixture-record"), "q4 Runtime run-plan execution smoke must prepare owner fixture state from app execution evidence")
assert(source.include?("kde-controlled-launch-session-bus-smoke-plan-preview"), "q4 Runtime run-plan execution smoke must derive the restricted session-bus plan from owner fixture evidence")
assert(source.include?("q4_dbus_controlled_launch_owner_fixture_smoke.rb"), "q4 Runtime run-plan execution smoke must execute the restricted q4 D-Bus owner fixture")
assert(source.include?("q4_dbus_fixture_passed"), "q4 Runtime run-plan execution smoke must report q4 D-Bus fixture readiness")
assert(source.include?("q4_dbus_fixture_desktop_action_metadata_consumed"), "q4 Runtime run-plan execution smoke must report q4 desktop action metadata consumption")
assert(source.include?("q4_dbus_fixture_desktop_trigger_preflight_ready_for_operator_request"), "q4 Runtime run-plan execution smoke must report Go-owned desktop-trigger preflight readiness")
assert(source.include?("owner_fixture_from_app_execution_ready"), "q4 Runtime run-plan execution smoke must report owner fixture readiness")
assert(source.include?("session_bus_plan_from_app_execution_ready"), "q4 Runtime run-plan execution smoke must report session-bus plan readiness")
assert(source.include?("run_plan_generated_by_runtime"), "q4 Runtime run-plan execution smoke must report Runtime-generated run plans")
assert(source.include?("direct_run_plan_input"), "q4 Runtime run-plan execution smoke must prove no caller-supplied run plan is needed")
assert(source.include?("--execute"), "q4 Runtime run-plan execution smoke must support explicit execution")
assert(source.include?("q4_messagebox_smoke_planned"), "q4 Runtime run-plan execution smoke must target the real q4 MessageBox smoke")
assert(source.include?("actual_windows_app_run_observed"), "q4 Runtime run-plan execution smoke must require actual Windows app evidence")
assert(source.include?("compatibility_center_projection_ready"), "q4 Runtime run-plan execution smoke must require Compatibility Center projection readiness")
assert(source.include?("kde_center_projection_ready"), "q4 Runtime run-plan execution smoke must require KDE Center projection readiness")
assert(source.include?("owner_file_open_entrypoint_invoked"), "q4 Runtime run-plan execution smoke must require owner file-open evidence")
assert(source.include?("document_content_marker_observed"), "q4 Runtime run-plan execution smoke must require document marker evidence")
assert(source.include?("go_owned_q4_winapp_acceptance_ready"), "q4 Runtime run-plan execution smoke must require Go-owned q4 acceptance")
assert(source.include?("runtime_binary_path_exposed"), "q4 Runtime run-plan execution smoke must expose the binary path gate")
assert(source.include?("runtime_smoke_command_arguments_exposed"), "q4 Runtime run-plan execution smoke must expose the command argument gate")
assert(source.include?("host_compilation_avoided"), "q4 Runtime run-plan execution smoke must report host compilation avoidance")
assert(source.include?("docs/claude-code-implementation-packages.md") == false, "q4 Runtime run-plan execution smoke must not touch Claude's package document")

stdout, stderr, status = Open3.capture3(
  "ruby",
  script.to_s,
  "--output", "/tmp/xnix-q4-runtime-run-plan-execution-plan.json"
)
assert(status.success?, "q4 Runtime run-plan execution plan must exit successfully: #{stderr}")
payload = JSON.parse(stdout)
assert(payload.fetch("schema_version") == "xnix.scripts.q4_runtime_run_plan_execution_smoke.v1", "plan must expose schema")
assert(payload.fetch("request_type") == "q4-runtime-run-plan-execution-smoke", "plan must expose request type")
assert(payload.fetch("status") == "planned", "plan must not execute by default")
assert(payload.fetch("execute") == false, "plan must keep execute disabled by default")
assert(payload.fetch("remote_host") == "root@q4", "plan must default to q4")
assert(payload.fetch("remote_runtime_build_planned") == true, "plan must build Runtime on q4")
assert(payload.fetch("runtime_binary_goos") == "darwin", "plan must default Runtime host GOOS to darwin")
assert(payload.fetch("runtime_binary_goarch") == "arm64", "plan must default Runtime host GOARCH to arm64")
assert(payload.fetch("runtime_binary_built_on_q4") == false, "plan must not claim a built binary before execute")
assert(payload.fetch("runtime_binary_fetched") == false, "plan must not fetch a binary before execute")
assert(payload.fetch("runtime_binary_path_exposed") == false, "plan must not expose fetched binary paths")
assert(payload.fetch("runtime_command") == "known-app-verified-catalog-app-execution", "plan must use the Runtime app execution command")
assert(payload.fetch("runtime_command_execute_planned") == true, "plan must plan Runtime command execution")
assert(payload.fetch("verified_catalog_to_app_execution_planned") == true, "plan must execute from verified catalog and app id")
assert(payload.fetch("desktop_consumption_planned") == true, "plan must feed execution evidence to desktop read models after execute")
assert(payload.fetch("run_plan_generated_by_runtime") == false, "plan must not claim Runtime-generated run plans before execute")
assert(payload.fetch("direct_run_plan_input") == false, "plan must not rely on a caller-supplied run plan")
assert(payload.fetch("go_runtime_entrypoint_invoked") == false, "plan must not invoke Runtime before execute")
assert(payload.fetch("app_id") == "org.xnix.apps.messagebox", "plan must target MessageBox")
assert(payload.fetch("q4_messagebox_smoke_planned") == true, "plan must target q4 MessageBox smoke")
assert(payload.fetch("q4_messagebox_smoke_passed") == false, "plan must not claim q4 smoke before execute")
assert(payload.fetch("actual_windows_app_run_observed") == false, "plan must not claim actual run before execute")
assert(payload.fetch("compatibility_center_projection_ready") == false, "plan must not claim Compatibility Center projection before execute")
assert(payload.fetch("kde_center_projection_ready") == false, "plan must not claim KDE Center projection before execute")
assert(payload.fetch("compatibility_center_consumed") == false, "plan must not claim Compatibility Center consumption before execute")
assert(payload.fetch("kde_center_page_consumed") == false, "plan must not claim KDE Center page consumption before execute")
assert(payload.fetch("known_app_smoke_evidence_count") == 0, "plan must not claim smoke evidence counts before execute")
assert(payload.fetch("known_app_gui_evidence_count") == 0, "plan must not claim KDE GUI evidence counts before execute")
assert(payload.fetch("known_app_gui_controlled_action_present") == false, "plan must not claim controlled action evidence before execute")
assert(payload.fetch("kde_controlled_launch_action_preview_ready") == false, "plan must not claim KDE controlled action readiness before execute")
assert(payload.fetch("owner_fixture_from_app_execution_planned") == true, "plan must derive owner fixture readiness after execute")
assert(payload.fetch("owner_fixture_from_app_execution_ready") == false, "plan must not claim owner fixture readiness before execute")
assert(payload.fetch("owner_fixture_desktop_trigger_ready") == false, "plan must not claim owner desktop trigger before execute")
assert(payload.fetch("owner_fixture_service_call_ready") == false, "plan must not claim owner service bridge before execute")
assert(payload.fetch("owner_fixture_kde_forwards_only_evidence_handle") == false, "plan must not claim owner fixture evidence-only forwarding before execute")
assert(payload.fetch("owner_fixture_evidence_path_exposed") == false, "plan must not expose owner fixture evidence paths")
assert(payload.fetch("owner_fixture_backend_launch_enabled") == false, "plan must keep owner fixture backend launch disabled")
assert(payload.fetch("owner_fixture_backend_process_started") == false, "plan must not start owner fixture backend processes")
assert(payload.fetch("session_bus_plan_from_app_execution_planned") == true, "plan must derive a session-bus plan from app execution after execute")
assert(payload.fetch("session_bus_plan_from_app_execution_ready") == false, "plan must not claim session-bus plan readiness before execute")
assert(payload.fetch("session_bus_plan_private_bus_required") == false, "plan must not claim private bus before execute")
assert(payload.fetch("session_bus_plan_dbus_fixture_ready") == false, "plan must not claim D-Bus fixture readiness before execute")
assert(payload.fetch("session_bus_plan_kde_forwards_only_evidence_handle") == false, "plan must not claim session-bus evidence-only forwarding before execute")
assert(payload.fetch("session_bus_plan_backend_process_started") == false, "plan must not start session-bus backend processes")
assert(payload.fetch("session_bus_plan_host_root_modified") == false, "plan must not mutate host root through session-bus plan")
assert(payload.fetch("q4_dbus_fixture_execute_planned") == true, "plan must run q4 D-Bus fixture after execute")
assert(payload.fetch("q4_dbus_fixture_passed") == false, "plan must not claim q4 D-Bus fixture pass before execute")
assert(payload.fetch("q4_dbus_fixture_linux_runtime_built_on_q4") == false, "plan must not claim q4 Linux Runtime build before execute")
assert(payload.fetch("q4_dbus_fixture_adapter_compiled_in_q4_container") == false, "plan must not claim q4 C adapter compilation before execute")
assert(payload.fetch("q4_dbus_fixture_desktop_action_metadata_consumed") == false, "plan must not claim q4 desktop action metadata consumption before execute")
assert(payload.fetch("q4_dbus_fixture_desktop_action_metadata_path_exposed") == false, "plan must not expose q4 desktop action metadata paths")
assert(payload.fetch("q4_dbus_fixture_desktop_trigger_preflight_checked") == false, "plan must not claim q4 desktop-trigger preflight before execute")
assert(payload.fetch("q4_dbus_fixture_desktop_trigger_preflight_blocked_missing_promotion") == false, "plan must not claim q4 fail-closed preflight before execute")
assert(payload.fetch("q4_dbus_fixture_desktop_trigger_preflight_ready_for_operator_request") == false, "plan must not claim q4 operator request readiness before execute")
assert(payload.fetch("q4_dbus_fixture_desktop_trigger_preflight_owner_service_shape_verified") == false, "plan must not claim q4 owner service shape verification before execute")
assert(payload.fetch("q4_dbus_fixture_desktop_trigger_preflight_path_exposed") == false, "plan must not expose q4 desktop-trigger preflight paths")
assert(payload.fetch("q4_dbus_fixture_app_execution_evidence_copied_to_q4") == false, "plan must not copy q4 D-Bus evidence before execute")
assert(payload.fetch("q4_dbus_fixture_app_execution_evidence_path_exposed") == false, "plan must not expose q4 D-Bus evidence paths")
assert(payload.fetch("q4_dbus_fixture_docker_socket_mounted") == false, "plan must not mount Docker socket for q4 D-Bus fixture")
assert(payload.fetch("q4_dbus_fixture_broad_host_mount_required") == false, "plan must not broad-mount host paths for q4 D-Bus fixture")
assert(payload.fetch("q4_dbus_fixture_host_networking_required") == false, "plan must not require host networking for q4 D-Bus fixture")
assert(payload.fetch("q4_dbus_fixture_privileged_container_required") == false, "plan must not require privileged containers for q4 D-Bus fixture")
assert(payload.fetch("desktop_consumption_evidence_path_exposed") == false, "plan must not expose desktop evidence paths")
assert(payload.fetch("host_compilation_avoided") == true, "plan must avoid host compilation")
assert(payload.fetch("host_root_modified") == false, "plan must not mutate host root")
assert(payload.fetch("privileged_container_required") == false, "plan must not require privileged containers")
assert(payload.fetch("host_networking_required") == false, "plan must not require host networking")
assert(payload.fetch("docker_socket_mounted") == false, "plan must not mount Docker socket")
assert(payload.fetch("broad_host_mount_required") == false, "plan must not require broad host mounts")

bad_stdout, bad_stderr, bad_status = Open3.capture3(
  "ruby",
  script.to_s,
  "--fetch-root", "/Users/rocky/not-xnix"
)
assert(!bad_status.success?, "q4 Runtime run-plan execution smoke must reject unsafe fetch roots")
assert((bad_stdout + bad_stderr).include?("fetch root must stay under /tmp/xnix-*"), "q4 Runtime run-plan execution smoke must explain unsafe fetch roots")

puts "PASS: q4 Runtime run-plan execution smoke script is q4-built and Go-entrypoint-backed"
