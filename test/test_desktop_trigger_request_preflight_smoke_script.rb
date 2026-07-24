#!/usr/bin/env ruby
# frozen_string_literal: true

require "open3"
require "json"
require "pathname"

def assert(condition, message)
  return if condition

  warn "FAIL: #{message}"
  exit 1
end

project_root = Pathname.new(__dir__).join("..").realpath
script = project_root.join("scripts/desktop_trigger_request_preflight_smoke.rb")
contents = script.read
container_script = project_root.join("scripts/container.rb")
container_model = project_root.join("lib/xnix/container.rb")

assert(script.executable?, "Desktop-trigger request preflight smoke script must be executable")
assert(contents.include?("known-app-kde-runtime-status-launch-evidence-record"), "preflight smoke must record Runtime-owned fixture evidence")
assert(contents.include?("desktop-trigger-request-preflight-preview"), "preflight smoke must invoke the Go-owned request preflight")
assert(contents.include?("blocked-missing-promotion"), "preflight smoke must verify the missing-promotion blocked state")
assert(contents.include?("ready-for-operator-request"), "preflight smoke must verify the operator-ready state")
assert(contents.include?("owner_service_call_shape_verified"), "preflight smoke must verify owner service call shape readiness")
assert(contents.include?("operator_request_ready"), "preflight smoke must verify operator request readiness")
assert(contents.include?("formal_promotion_observed"), "preflight smoke must verify explicit promotion observation")
assert(contents.include?("formal_release_ready"), "preflight smoke must keep formal release readiness separate")
assert(contents.include?("kde_receives_materialized_owner_args"), "preflight smoke must verify KDE does not receive owner args")
assert(contents.include?("service_call_dispatch_enabled"), "preflight smoke must keep service dispatch disabled")
assert(contents.include?("service_call_dispatched"), "preflight smoke must keep service calls undispatched")
assert(contents.include?("dbus_called"), "preflight smoke must keep D-Bus calls disabled")
assert(contents.include?("desktop_launch_enabled"), "preflight smoke must keep desktop launch disabled")
assert(contents.include?("backend_launch_enabled"), "preflight smoke must keep backend launch disabled")
assert(contents.include?("runtime_state_written"), "preflight smoke must keep Runtime writes disabled")
assert(contents.include?("kde_configuration_written"), "preflight smoke must keep KDE writes disabled")
assert(contents.include?("host_root_modified"), "preflight smoke must keep host-root mutation disabled")
assert(contents.include?("assert_no_forbidden"), "preflight smoke must verify redaction")
assert(contents.include?("owner_service_call_args"), "preflight smoke must guard owner service call argument exposure")
assert(contents.include?("owner_service_cli_args"), "preflight smoke must guard owner service CLI argument exposure")
assert(contents.include?("qemu-system"), "preflight smoke must guard QEMU backend exposure")
assert(contents.include?("wine "), "preflight smoke must guard Wine backend exposure")
assert(contents.include?("PASS: desktop-trigger request preflight smoke"), "preflight smoke must print the PASS marker")
assert(container_script.read.include?("desktop-trigger-request-preflight-smoke"), "Container CLI must expose desktop-trigger-request-preflight-smoke")
assert(container_model.read.include?("desktop_trigger_request_preflight_smoke_command"), "Container model must expose desktop trigger request preflight smoke command")

stdout, stderr, status = Open3.capture3("ruby", script.to_s, chdir: project_root.to_s)
assert(status.success?, "Desktop-trigger request preflight smoke must pass cleanly: #{stderr}\n#{stdout}")
assert(stdout.include?("PASS: desktop-trigger request preflight smoke"), "Desktop-trigger request preflight smoke must print the PASS marker")

json_stdout, json_stderr, json_status = Open3.capture3("ruby", script.to_s, "--format", "json", chdir: project_root.to_s)
assert(json_status.success?, "Desktop-trigger request preflight smoke JSON must pass cleanly: #{json_stderr}\n#{json_stdout}")
json_payload = JSON.parse(json_stdout)
assert(json_payload.fetch("schema_version") == "xnix.runtime.desktop_trigger_request_preflight_smoke.v1", "preflight smoke JSON must expose its schema")
assert(json_payload.fetch("smoke_passed") == true, "preflight smoke JSON must report pass state")
assert(json_payload.fetch("blocked_preflight_state") == "blocked-missing-promotion", "preflight smoke JSON must report blocked state")
assert(json_payload.fetch("ready_preflight_state") == "ready-for-operator-request", "preflight smoke JSON must report ready state")
assert(json_payload.fetch("service_call_dispatched") == false, "preflight smoke JSON must keep service dispatch disabled")
assert(!json_stdout.include?("owner_service_cli_args"), "preflight smoke JSON must not expose owner service CLI args")
assert(!json_stdout.include?("owner_service_call_args"), "preflight smoke JSON must not expose owner service call args")

puts "PASS: desktop-trigger request preflight smoke script unit tests"
