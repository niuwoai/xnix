#!/usr/bin/env ruby
# frozen_string_literal: true

require "open3"
require "pathname"

def assert(condition, message)
  return if condition

  warn "FAIL: #{message}"
  exit 1
end

project_root = Pathname.new(__dir__).join("..").realpath
script = project_root.join("scripts/kde_controlled_launch_action_smoke.rb")
contents = script.read
container_script = project_root.join("scripts/container.rb")
container_model = project_root.join("lib/xnix/container.rb")

assert(script.executable?, "KDE controlled launch action smoke script must be executable")
assert(contents.include?("kde-controlled-launch-session-bus-smoke-plan-preview"), "KDE action smoke must consume the Go-owned session-bus smoke plan")
assert(contents.include?("known-app-runtime-status-launch-owner-fixture-record"), "KDE action smoke must prepare Runtime-owned fixture evidence")
assert(contents.include?("runtime_status_owner_session_smoke_command"), "KDE action smoke must execute the plan-provided session smoke command")
assert(contents.include?("dbus_controlled_launch_fixture_command"), "KDE action smoke must consume the plan-provided D-Bus fixture command")
assert(contents.include?("dbus_controlled_launch_fixture_container_command"), "KDE action smoke must consume the plan-provided D-Bus fixture container command")
assert(contents.include?("gui-smoke-evidence-preview"), "KDE action smoke must project safe GUI smoke evidence")
assert(contents.include?("kde-center-page-preview"), "KDE action smoke must render the KDE center page")
assert(contents.include?("kde-controlled-launch-action-preview"), "KDE action smoke must consume the KDE controlled-launch action preview")
assert(contents.include?("--kde-center-page-file"), "KDE action smoke must feed the full KDE center page back to Runtime")
assert(contents.include?("known_app_owner_controlled_gui_evidence_count"), "KDE action smoke must verify owner-controlled GUI evidence count")
assert(contents.include?("known_app_owner_file_open_verified_count"), "KDE action smoke must verify owner file-open evidence count")
assert(contents.include?("owner_file_open_verified"), "KDE action smoke must verify owner file-open action evidence")
assert(contents.include?("owner_delegated_window_match"), "KDE action smoke must verify safe owner window-match evidence")
assert(contents.include?("GUI center-page handoff"), "KDE action smoke skip text must mention GUI center-page handoff validation")
assert(contents.include?("XNIX_KDE_CONTROLLED_LAUNCH_ACTION_SMOKE_EXECUTE"), "KDE action smoke must guard execution behind an explicit environment variable")
assert(contents.include?("XNIX_KDE_CONTROLLED_LAUNCH_ACTION_SMOKE_EXECUTE_DBUS_FIXTURE"), "KDE action smoke must guard D-Bus fixture execution behind an explicit environment variable")
assert(contents.include?("PASS: KDE controlled launch action smoke"), "KDE action smoke must print a pass marker")
assert(contents.include?("SKIP: KDE controlled launch action smoke"), "KDE action smoke must print a skip marker")
assert(contents.include?("kde_forwards_only_evidence_handle"), "KDE action smoke must verify evidence-only KDE forwarding")
assert(contents.include?("owner_service_args_exposed_to_kde"), "KDE action smoke must verify owner service args stay hidden")
assert(contents.include?("desktop_kde_state_root_access"), "KDE action smoke must verify KDE state-root access stays disabled")
assert(contents.include?("desktop_receipt_fields_reconstructed"), "KDE action smoke must verify KDE does not reconstruct receipts")
assert(contents.include?("smoke_executed_by_preview"), "KDE action smoke must verify the Go preview does not execute")
assert(contents.include?("docker.sock"), "KDE action smoke must guard Docker socket exposure")
assert(contents.include?("--privileged"), "KDE action smoke must guard privileged container exposure")
assert(contents.include?("--network host"), "KDE action smoke must guard host networking exposure")
assert(container_script.read.include?("kde-controlled-launch-action-smoke"), "Container CLI must expose kde-controlled-launch-action-smoke")
assert(container_model.read.include?("kde_controlled_launch_action_smoke_command"), "Container model must expose KDE controlled launch action smoke command")

stdout, stderr, status = Open3.capture3("ruby", script.to_s, chdir: project_root.to_s)
assert(status.success?, "KDE controlled launch action smoke must pass or skip cleanly: #{stderr}")
assert(stdout.include?("PASS: KDE controlled launch action smoke") || stdout.include?("SKIP: KDE controlled launch action smoke"),
       "KDE controlled launch action smoke must print a pass or skip marker")

puts "PASS: KDE controlled launch action smoke script unit tests"
