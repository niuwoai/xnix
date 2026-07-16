#!/usr/bin/env ruby
# frozen_string_literal: true

require "json"
require "open3"
require "pathname"
require_relative "../lib/xnix/compatibility/compatibility_snapshot_plan"

def assert(condition, message)
  return if condition

  warn "FAIL: #{message}"
  exit 1
end

project_root = Pathname.new(__dir__).join("..").realpath

Xnix::Compatibility::CompatibilitySnapshotPlan::REASONS.keys.each do |reason|
  plan = Xnix::Compatibility::CompatibilitySnapshotPlan.new(
    application_id: "org.xnix.sample.notepad",
    reason: reason
  ).to_h

  assert(plan["version"] == "0.2.248", "compatibility snapshot plan must expose the current version")
  assert(plan["plan_type"] == "compatibility-snapshot", "compatibility snapshot plan must identify the model type")
  assert(plan["application_id"] == "org.xnix.sample.notepad", "compatibility snapshot plan must preserve the application id")
  assert(plan["reason"] == reason, "compatibility snapshot plan must preserve the reason")
  assert(plan["runtime_method"] == "GetSnapshotPlan", "compatibility snapshot plan must expose the Runtime method")
  assert(plan["runtime_owned"], "compatibility snapshot plan must be Runtime-owned")
  assert(plan["c_runtime_backed"], "compatibility snapshot plan must expose C Runtime backing")
  assert(!plan["kde_policy_owner"], "compatibility snapshot plan must not be KDE-owned")
  assert(plan["enabled_by_default"], "compatibility snapshot plan must default to enabled")
  assert(plan["snapshot_scope"]["application_state"], "compatibility snapshot plan must include application state")
  assert(plan["snapshot_scope"]["runtime_metadata"], "compatibility snapshot plan must include Runtime metadata")
  assert(plan["snapshot_scope"]["desktop_activation_receipts"], "compatibility snapshot plan must include desktop activation receipts")
  assert(!plan["snapshot_scope"]["user_documents"], "compatibility snapshot plan must not snapshot user documents")
  assert(!plan["snapshot_scope"]["host_system"], "compatibility snapshot plan must not snapshot the host system")
  assert(plan["restore"]["available"], "compatibility snapshot plan must expose restore availability")
  assert(plan["restore"]["requires_user_confirmation"], "compatibility snapshot restore must require confirmation")
  assert(plan["restore"]["preserve_user_documents"], "compatibility snapshot restore must preserve user documents")
  assert(plan["retention"]["policy"] == "bounded", "compatibility snapshot plan must use bounded retention")
  assert(plan["retention"]["keep_latest"] == 5, "compatibility snapshot plan must keep a bounded latest set")
  assert(!plan["snapshot_request_created"], "compatibility snapshot plan must not create snapshot requests")
  assert(!plan["snapshot_created"], "compatibility snapshot plan must not create snapshots")
  assert(!plan["restore_requested"], "compatibility snapshot plan must not request restore")
  assert(!plan["restore_executed"], "compatibility snapshot plan must not execute restore")
  assert(!plan["user_documents_included"], "compatibility snapshot plan must exclude user documents")
  assert(!plan["host_system_included"], "compatibility snapshot plan must exclude the host system")
  assert(!plan["host_root_modified"], "compatibility snapshot plan must not mutate the host root")
  assert(!plan["backend_details_exposed"], "compatibility snapshot plan must hide backend details")
end

repair_plan = Xnix::Compatibility::CompatibilitySnapshotPlan.new(
  application_id: "org.xnix.sample.notepad",
  reason: "before-repair"
).to_h
json = JSON.pretty_generate(repair_plan)
assert(!json.match?(/wine|prefix|\.wine|proton|virtual machine/i), "compatibility snapshot plan must not expose backend implementation terms")

stdout, stderr, status = Open3.capture3(
  "ruby",
  project_root.join("bin/xnix-compat-snapshot-plan").to_s,
  "--app",
  "org.xnix.sample.notepad",
  "--reason",
  "before-repair"
)
assert(status.success?, "compatibility snapshot plan CLI must exit successfully: #{stderr}")
assert(JSON.parse(stdout) == repair_plan, "compatibility snapshot plan CLI must emit the plan")

_stdout, stderr, status = Open3.capture3("ruby", project_root.join("bin/xnix-compat-snapshot-plan").to_s)
assert(!status.success?, "compatibility snapshot plan CLI must require an application id")
assert(stderr.include?("--app is required"), "compatibility snapshot plan CLI must explain missing application ids")

puts "PASS: compatibility snapshot plan unit tests"
