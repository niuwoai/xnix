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
script = project_root.join("scripts/dbus_controlled_launch_owner_fixture_smoke.rb")
container_script = project_root.join("scripts/container.rb")
container_model = project_root.join("lib/xnix/container.rb")
contents = script.read

assert(script.executable?, "D-Bus controlled launch owner fixture smoke script must be executable")
assert(contents.include?("D-Bus controlled launch owner fixture smoke"), "D-Bus controlled launch owner fixture smoke must name its pass marker")
assert(contents.include?("dbus-run-session"), "D-Bus controlled launch owner fixture smoke must run inside a private session bus")
assert(contents.include?("DBUS_SESSION_BUS_ADDRESS"), "D-Bus controlled launch owner fixture smoke must require a session bus address")
assert(contents.include?("xnix-dbus-smoke"), "D-Bus controlled launch owner fixture smoke must launch the C D-Bus adapter")
assert(contents.include?("gdbus"), "D-Bus controlled launch owner fixture smoke must call the D-Bus method through gdbus")
assert(contents.include?("ShowRuntimeControlledLaunch"), "D-Bus controlled launch owner fixture smoke must call ShowRuntimeControlledLaunch")
assert(contents.include?("org.xnix.Compatibility1.ShowRuntimeControlledLaunch"), "D-Bus controlled launch owner fixture smoke must use the public D-Bus method name")
assert(contents.include?("known-app-runtime-status-launch-owner-fixture-record"), "D-Bus controlled launch owner fixture smoke must delegate fixture setup to the Go Runtime owner fixture helper")
assert(!contents.include?("known-app-launch-authorization-receipt-preview"), "D-Bus controlled launch owner fixture smoke must not manually create launch authorization receipts in Ruby")
assert(!contents.include?("known-app-controlled-execution-session-record"), "D-Bus controlled launch owner fixture smoke must not manually persist controlled sessions in Ruby")
assert(!contents.include?("known-app-session-gated-launch-review-receipt-record"), "D-Bus controlled launch owner fixture smoke must not manually record review receipts in Ruby")
assert(!contents.include?("known-app-kde-runtime-status-launch-evidence-record"), "D-Bus controlled launch owner fixture smoke must not manually persist evidence handoffs in Ruby")
assert(contents.include?("XNIX_RUNTIME_OWNER_STATE_ROOT"), "D-Bus controlled launch owner fixture smoke must inject state root through the Runtime owner environment")
assert(contents.include?("XNIX_RUNTIME_OWNER_KNOWN_APP_CACHE_ROOT"), "D-Bus controlled launch owner fixture smoke must inject cache root through the Runtime owner environment")
assert(contents.include?("XNIX_RUNTIME_OWNER_MANAGED_LAUNCHER"), "D-Bus controlled launch owner fixture smoke must inject the managed launcher through the Runtime owner environment")
assert(contents.include?("XNIX_DBUS_CONTROLLED_LAUNCH_ARGS_FILE"), "D-Bus controlled launch owner fixture smoke must capture delegated launcher argv")
assert(contents.include?(".xnix-dbus-controlled-launch-scratch"), "D-Bus controlled launch owner fixture smoke must prefer the managed container scratch volume")
assert(contents.include?("go_owner_service_call_json"), "D-Bus controlled launch owner fixture smoke must parse the nested owner service-call payload")
assert(contents.include?("runtime-owner-service-call"), "D-Bus controlled launch owner fixture smoke must verify the owner service-call envelope")
assert(contents.include?("desktop-action-dispatch"), "D-Bus controlled launch owner fixture smoke must verify desktop action dispatch")
assert(contents.include?("kde-dbus-runtime-status-action"), "D-Bus controlled launch owner fixture smoke must verify the KDE Runtime-status action route")
assert(contents.include?("kde_forwards_only_evidence_handle"), "D-Bus controlled launch owner fixture smoke must prove KDE forwards only evidence handles")
assert(contents.include?("desktop_receipt_fields_reconstructed"), "D-Bus controlled launch owner fixture smoke must prove KDE does not reconstruct receipt fields")
assert(contents.include?("desktop_kde_state_root_access"), "D-Bus controlled launch owner fixture smoke must prove KDE does not receive state-root access")
assert(contents.include?("docker.sock"), "D-Bus controlled launch owner fixture smoke must guard Docker socket exposure")
assert(contents.include?("--privileged"), "D-Bus controlled launch owner fixture smoke must guard privileged container exposure")
assert(contents.include?("--network host"), "D-Bus controlled launch owner fixture smoke must guard host networking exposure")
assert(contents.include?("type=bind"), "D-Bus controlled launch owner fixture smoke must guard bind-mount exposure")

assert(container_script.read.include?("dbus-controlled-launch-owner-fixture-smoke"), "Container CLI must expose dbus-controlled-launch-owner-fixture-smoke")
assert(container_model.read.include?("dbus_controlled_launch_owner_fixture_smoke_command"), "Container model must expose D-Bus controlled launch owner fixture smoke command")
assert(container_model.read.include?("CONTROLLED_LAUNCH_SCRATCH_SIZE_BYTES"), "Container model must expose a managed scratch tmpfs for D-Bus controlled launch owner fixture smoke")

stdout, stderr, status = Open3.capture3("ruby", script.to_s)
assert(status.success?, "D-Bus controlled launch owner fixture smoke must pass or skip cleanly: #{stderr}")
assert(stdout.include?("PASS: D-Bus controlled launch owner fixture smoke") || stdout.include?("SKIP: D-Bus controlled launch owner fixture smoke"),
       "D-Bus controlled launch owner fixture smoke must print a pass or skip marker")

puts "PASS: D-Bus controlled launch owner fixture smoke script unit tests"
