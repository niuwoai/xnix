#!/usr/bin/env ruby
# frozen_string_literal: true

require "json"
require "open3"
require "pathname"
require_relative "../lib/xnix/compatibility/runtime_owner_smoke_plan"

def assert(condition, message)
  return if condition

  warn "FAIL: #{message}"
  exit 1
end

project_root = Pathname.new(__dir__).join("..").realpath
plan = Xnix::Compatibility::RuntimeOwnerSmokePlan.new.to_h
step_ids = plan.fetch("steps").map { |item| item.fetch("id") }

assert(plan["version"] == "0.2.154", "Runtime owner smoke plan must expose the current version")
assert(plan["plan_type"] == "runtime-owner-smoke-plan", "Runtime owner smoke plan must identify the plan type")
assert(plan["runtime_owned"], "Runtime must own production owner smoke planning")
assert(!plan["kde_policy_owner"], "KDE must not own production owner smoke planning")
assert(plan["bus_name"] == "org.xnix.Compatibility1", "Runtime owner smoke plan must expose the stable bus name")
assert(plan["object_path"] == "/org/xnix/Compatibility1", "Runtime owner smoke plan must expose the stable object path")
assert(plan["interface"] == "org.xnix.Compatibility1", "Runtime owner smoke plan must expose the stable interface")
assert(plan["activation_binding_ready"], "Runtime owner smoke plan must depend on aligned activation binding")
assert(!plan["live_dbus_owner_ready"], "Runtime owner smoke plan must not claim live owner readiness")
assert(!plan["production_owner_enabled"], "Runtime owner smoke plan must not enable production ownership")
assert(!plan["owner_transition_ready"], "Runtime owner smoke plan must not enable owner transition")
assert(plan["smoke_state"] == "planned", "Runtime owner smoke plan must stay planned")
assert(plan["smoke_environment"] == "restricted-session", "Runtime owner smoke plan must stay in a restricted session")
assert(step_ids == %w[
  validate-activation-files
  start-packaged-runtime-owner
  assert-stable-bus-name
  check-read-only-method-parity
  reject-write-methods
  verify-non-production-smoke-adapter-boundary
  report-kde-safe-summary
], "Runtime owner smoke plan must report expected steps")
assert(plan["counts"] == { "total" => 7, "passed" => 1, "pending" => 6, "blocked" => 0 }, "Runtime owner smoke plan must count step states")
assert(plan["blocked_actions"].length == 5, "Runtime owner smoke plan must list blocked actions")
assert(!plan["network_required"], "Runtime owner smoke plan must not require network access")
assert(!plan["host_root_modified"], "Runtime owner smoke plan must not mutate the host root")
assert(!plan["privileged_container_required"], "Runtime owner smoke plan must not require privileged containers")
assert(!plan["system_service_started"], "Runtime owner smoke plan must not start a system service")
assert(!plan["production_bus_claimed"], "Runtime owner smoke plan must not claim the production bus")
assert(!plan["backend_details_exposed"], "Runtime owner smoke plan must hide backend details")

json = JSON.pretty_generate(plan)
assert(!json.match?(/wine|prefix|\.wine|proton|virtual machine/i), "Runtime owner smoke plan must not expose backend implementation terms")

stdout, stderr, status = Open3.capture3("ruby", project_root.join("bin/xnix-runtime-owner-smoke-plan").to_s)
assert(status.success?, "Runtime owner smoke plan CLI must exit successfully: #{stderr}")
assert(JSON.parse(stdout) == plan, "Runtime owner smoke plan CLI must emit the plan model")

puts "PASS: Runtime owner smoke plan unit tests"
