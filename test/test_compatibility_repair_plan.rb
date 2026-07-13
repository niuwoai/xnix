#!/usr/bin/env ruby
# frozen_string_literal: true

require "json"
require "open3"
require "pathname"
require_relative "../lib/xnix/compatibility/compatibility_repair_plan"

def assert(condition, message)
  return if condition

  warn "FAIL: #{message}"
  exit 1
end

project_root = Pathname.new(__dir__).join("..").realpath
issues = Xnix::Compatibility::CompatibilityRepairPlan::ISSUES.keys

issues.each do |issue|
  plan = Xnix::Compatibility::CompatibilityRepairPlan.new(
    application_id: "org.xnix.sample.notepad",
    issue: issue
  ).to_h

  assert(plan["version"] == "0.2.89", "compatibility repair plan must expose the current version")
  assert(plan["plan_type"] == "compatibility-repair", "compatibility repair plan must identify the model type")
  assert(plan["application_id"] == "org.xnix.sample.notepad", "compatibility repair plan must preserve the application id")
  assert(plan["issue"] == issue, "compatibility repair plan must preserve the issue")
  assert(%w[info warning critical].include?(plan["severity"]), "compatibility repair plan must expose severity")
  assert([true, false].include?(plan["automatic_allowed"]), "compatibility repair plan must expose automatic allowance")
  assert([true, false].include?(plan["user_approval_required"]), "compatibility repair plan must expose approval requirements")
  assert([true, false].include?(plan["snapshot_required"]), "compatibility repair plan must expose snapshot requirement")
  if plan["snapshot_required"]
    assert(plan["snapshot_plan"]["plan_type"] == "compatibility-snapshot", "snapshot-required repairs must include a snapshot plan")
    assert(plan["snapshot_plan"]["restore_available"], "snapshot-required repairs must expose restore availability")
    assert(plan["snapshot_plan"]["preserve_user_documents"], "snapshot-required repairs must preserve user documents")
  else
    assert(plan["snapshot_plan"].nil?, "repairs without snapshot requirements must not include a snapshot plan")
  end
  assert(plan["rollback_available"], "compatibility repair plan must keep rollback available")
  assert(!plan["actions"].empty?, "compatibility repair plan must include actions")
  assert(%w[approval-required repair-applied install-failed].include?(plan["notification_event"]), "compatibility repair plan must map to notification events")
end

pending_plan = Xnix::Compatibility::CompatibilityRepairPlan.new(
  application_id: "org.xnix.sample.notepad",
  issue: "engine-binding-pending"
).to_h
assert(pending_plan["user_approval_required"], "pending engine setup must require approval")
assert(pending_plan["snapshot_required"], "pending engine setup must require a snapshot")
assert(pending_plan["snapshot_plan"]["reason"] == "before-repair", "pending engine setup must plan a repair snapshot")
assert(pending_plan["notification_event"] == "approval-required", "pending engine setup must request approval notification")

applied_plan = Xnix::Compatibility::CompatibilityRepairPlan.new(
  application_id: "org.xnix.sample.notepad",
  issue: "runtime-repair-applied"
).to_h
assert(applied_plan["automatic_allowed"], "safe applied repair must allow automation")
assert(!applied_plan["user_approval_required"], "safe applied repair must not require approval")
assert(applied_plan["notification_event"] == "repair-applied", "safe applied repair must map to repair notification")

json = JSON.pretty_generate(pending_plan)
assert(!json.match?(/wine|prefix|\.wine|proton|virtual machine/i), "compatibility repair plan must not expose backend implementation terms")

stdout, stderr, status = Open3.capture3(
  "ruby",
  project_root.join("bin/xnix-compat-repair-plan").to_s,
  "--app",
  "org.xnix.sample.notepad",
  "--issue",
  "engine-binding-pending"
)
assert(status.success?, "compatibility repair plan CLI must exit successfully: #{stderr}")
assert(JSON.parse(stdout) == pending_plan, "compatibility repair plan CLI must emit the repair plan")

_stdout, stderr, status = Open3.capture3("ruby", project_root.join("bin/xnix-compat-repair-plan").to_s)
assert(!status.success?, "compatibility repair plan CLI must require an application id")
assert(stderr.include?("--app is required"), "compatibility repair plan CLI must explain missing application ids")

puts "PASS: compatibility repair plan unit tests"
