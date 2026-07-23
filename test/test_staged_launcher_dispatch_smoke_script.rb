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
script = project_root.join("scripts/staged_launcher_dispatch_smoke.rb")
container_script = project_root.join("scripts/container.rb")
container_model = project_root.join("lib/xnix/container.rb")
contents = script.read

assert(script.executable?, "Staged launcher dispatch smoke script must be executable")
assert(contents.include?("windows-known-app-dispatch-preview"), "Staged launcher dispatch smoke must preflight the managed artifact")
assert(contents.include?("desktop-activation-stage"), "Staged launcher dispatch smoke must stage desktop activation files")
assert(contents.include?("--managed-launcher-bin"), "Staged launcher dispatch smoke must copy the Go launcher binary")
assert(contents.include?("usr/local/bin/xnix-compat-launch"), "Staged launcher dispatch smoke must execute the staged launcher")
assert(contents.include?("--guest-boundary"), "Staged launcher dispatch smoke must supply the controlled guest boundary")
assert(contents.include?("--state-root"), "Staged launcher dispatch smoke must supply the Runtime state root to the launcher")
assert(contents.include?("--receipt-id"), "Staged launcher dispatch smoke must supply the opaque receipt id to the launcher")
assert(contents.include?("managed-known-app-guest-smoke"), "Staged launcher dispatch smoke must use the known app guest smoke boundary")
assert(contents.include?("windows-known-app-dispatch-smoke"), "Staged launcher dispatch smoke must enter the dispatch smoke path")
assert(contents.include?("compatibility-center-preview"), "Staged launcher dispatch smoke must feed Compatibility Center preview evidence")
assert(contents.include?("known-app-launch-authorization-receipt-preview"), "Staged launcher dispatch smoke must record a Runtime launch authorization receipt")
assert(contents.include?("known-app-launch-gate-preview"), "Staged launcher dispatch smoke must consume the Runtime launch gate")
assert(contents.include?("known-app-controlled-dispatch-request-preview"), "Staged launcher dispatch smoke must materialize a controlled dispatch request")
assert(contents.include?("known-app-controlled-execution-session-preview"), "Staged launcher dispatch smoke must materialize a controlled execution session handoff")
assert(contents.include?("known_app_staged_launcher_passed_count"), "Staged launcher dispatch smoke must verify Compatibility Center staged launcher evidence")
assert(contents.include?("known_app_launch_authorization_required_count"), "Staged launcher dispatch smoke must verify launch authorization requirement count")
assert(contents.include?("known_app_launch_authorization_recorded_count"), "Staged launcher dispatch smoke must verify launch authorization receipt count")
assert(contents.include?("known_app_launch_gate_consumed_count"), "Staged launcher dispatch smoke must verify launch gate consumption count")
assert(contents.include?("known_app_controlled_dispatch_ready_count"), "Staged launcher dispatch smoke must verify controlled dispatch readiness count")
assert(contents.include?("validated-launch-gate-consumed"), "Staged launcher dispatch smoke must verify the launch-gate-consumed card state")
assert(contents.include?("review-launch-authorization"), "Staged launcher dispatch smoke must verify the launch authorization review action")
assert(contents.include?("review-controlled-dispatch"), "Staged launcher dispatch smoke must verify the controlled dispatch review action")
assert(contents.include?("direct_launch_enabled"), "Staged launcher dispatch smoke must keep direct launch disabled")
assert(contents.include?("launch_authorization_receipt_state"), "Staged launcher dispatch smoke must expose the launch authorization receipt state")
assert(contents.include?("launch_authorization_receipt_id"), "Staged launcher dispatch smoke must expose the opaque launch authorization receipt id")
assert(contents.include?("launch_gate_state"), "Staged launcher dispatch smoke must expose the launch gate state")
assert(contents.include?("launch_gate_consumed"), "Staged launcher dispatch smoke must expose launch gate consumption")
assert(contents.include?("launch_gate_receipt_accepted"), "Staged launcher dispatch smoke must expose launch gate receipt acceptance")
assert(contents.include?("launch_gate_guest_boundary_accepted"), "Staged launcher dispatch smoke must expose launch gate boundary acceptance")
assert(contents.include?("controlled_dispatch_ready"), "Staged launcher dispatch smoke must expose controlled dispatch readiness")
assert(contents.include?("controlled_dispatch_request_created"), "Staged launcher dispatch smoke must verify controlled dispatch request creation")
assert(contents.include?("runtime_owned_dispatch_request"), "Staged launcher dispatch smoke must verify Runtime ownership of controlled dispatch requests")
assert(contents.include?("runtime_owned_execution_session"), "Staged launcher dispatch smoke must verify Runtime ownership of controlled execution sessions")
assert(contents.include?("execution_session_handoff_created"), "Staged launcher dispatch smoke must verify controlled execution session handoff creation")
assert(contents.include?("session_handoff_ready"), "Staged launcher dispatch smoke must verify controlled execution session handoff readiness")
assert(contents.include?("dispatch_smoke_request_type"), "Staged launcher dispatch smoke must verify the controlled dispatch smoke request type")
assert(contents.include?("staged_launcher_verified"), "Staged launcher dispatch smoke must verify staged launcher evidence visibility")
assert(contents.include?("runtime_dispatch_verified"), "Staged launcher dispatch smoke must verify Runtime dispatch evidence visibility")
assert(contents.include?("launch_authorization_required"), "Staged launcher dispatch smoke must keep launch authorization visible")
assert(contents.include?("desktop_launch_enabled"), "Staged launcher dispatch smoke must keep desktop launch disabled in Center evidence")
assert(contents.include?("SKIP:"), "Staged launcher dispatch smoke must skip when external smoke prerequisites are missing")
assert(contents.include?("PASS:"), "Staged launcher dispatch smoke must print a pass marker")
assert(container_script.read.include?("staged-launcher-dispatch-smoke"), "Container CLI must expose staged-launcher-dispatch-smoke")
assert(container_model.read.include?("staged_launcher_dispatch_smoke_command"), "Container model must expose staged launcher dispatch smoke command")

stdout, stderr, status = Open3.capture3("ruby", script.to_s)
assert(status.success?, "Staged launcher dispatch smoke script must pass or skip cleanly: #{stderr}")
assert(stdout.include?("PASS: staged managed launcher dispatch smoke") || stdout.include?("SKIP: staged managed launcher dispatch smoke"),
       "Staged launcher dispatch smoke script must print a pass or skip marker")

puts "PASS: staged launcher dispatch smoke script unit tests"
