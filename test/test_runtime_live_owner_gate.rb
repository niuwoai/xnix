#!/usr/bin/env ruby
# frozen_string_literal: true

require "json"
require "open3"
require "pathname"
require_relative "../lib/xnix/compatibility/runtime_live_owner_gate"

def assert(condition, message)
  return if condition

  warn "FAIL: #{message}"
  exit 1
end

project_root = Pathname.new(__dir__).join("..").realpath
gate = Xnix::Compatibility::RuntimeLiveOwnerGate.new.to_h
gate_ids = gate.fetch("required_gates").map { |item| item.fetch("id") }
gate_statuses = gate.fetch("required_gates").map { |item| item.fetch("status") }

assert(gate["version"] == "0.2.231", "Runtime live owner gate must expose the current version")
assert(gate["gate_type"] == "runtime-live-owner-gate", "Runtime live owner gate must identify the gate type")
assert(gate["runtime_owned"], "Runtime must own live owner readiness")
assert(!gate["kde_policy_owner"], "KDE must not own live owner readiness")
assert(gate["bus_name"] == "org.xnix.Compatibility1", "Runtime live owner gate must expose the stable bus name")
assert(gate["object_path"] == "/org/xnix/Compatibility1", "Runtime live owner gate must expose the stable object path")
assert(gate["interface"] == "org.xnix.Compatibility1", "Runtime live owner gate must expose the stable interface")
assert(gate["activation_binding_ready"], "Runtime live owner gate must require aligned activation files")
assert(!gate["live_dbus_owner_ready"], "Runtime live owner gate must not claim live ownership yet")
assert(!gate["production_owner_enabled"], "Runtime live owner gate must not enable production ownership yet")
assert(!gate["owner_transition_ready"], "Runtime live owner gate must not allow transition while gates are pending")
assert(gate["smoke_adapter_available"], "Runtime live owner gate must report smoke adapter availability")
assert(!gate["smoke_adapter_is_production_owner"], "Smoke adapter must not be treated as the production owner")
assert(!gate["kde_may_claim_runtime_ownership"], "KDE must not claim Runtime ownership")
assert(gate_ids == %w[activation-binding long-running-runtime-owner bus-name-acquisition read-only-method-parity production-recipe-trust], "Runtime live owner gate must report expected gates")
assert(gate_statuses == %w[pass pending pending pending pending], "Runtime live owner gate must keep production gates pending")
assert(gate["blocked_reasons"].length == 5, "Runtime live owner gate must expose blocked reasons")
assert(!gate["network_required"], "Runtime live owner gate must not require network access")
assert(!gate["host_root_modified"], "Runtime live owner gate must not mutate the host root")
assert(!gate["privileged_container_required"], "Runtime live owner gate must not require privileged containers")
assert(!gate["backend_details_exposed"], "Runtime live owner gate must hide backend details")

json = JSON.pretty_generate(gate)
assert(!json.match?(/wine|prefix|\.wine|proton|virtual machine/i), "Runtime live owner gate must not expose backend implementation terms")

stdout, stderr, status = Open3.capture3("ruby", project_root.join("bin/xnix-runtime-live-owner-gate").to_s)
assert(status.success?, "Runtime live owner gate CLI must exit successfully: #{stderr}")
assert(JSON.parse(stdout) == gate, "Runtime live owner gate CLI must emit the gate model")

puts "PASS: Runtime live owner gate unit tests"
