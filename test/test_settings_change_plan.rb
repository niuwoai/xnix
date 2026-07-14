#!/usr/bin/env ruby
# frozen_string_literal: true

require "json"
require "open3"
require "pathname"
require_relative "../lib/xnix/compatibility/settings_change_plan"

def assert(condition, message)
  return if condition

  warn "FAIL: #{message}"
  exit 1
end

project_root = Pathname.new(__dir__).join("..").realpath
plan = Xnix::Compatibility::SettingsChangePlan.new(
  application_id: "org.xnix.sample.notepad",
  section_id: "resource-access",
  field_id: "documents",
  value: "allow"
).to_h

assert(plan["version"] == "0.2.144", "settings change plan must expose the current version")
assert(plan["plan_type"] == "settings-change-plan", "settings change plan must identify plan type")
assert(plan["application_id"] == "org.xnix.sample.notepad", "settings change plan must preserve application id")
assert(plan["runtime_owned"], "Runtime must own settings change plans")
assert(!plan["kde_policy_owner"], "KDE must not own settings change policy")
assert(plan["section_id"] == "resource-access", "settings change plan must preserve section id")
assert(plan["field_id"] == "documents", "settings change plan must preserve field id")
assert(plan["requested_value"] == "allow", "settings change plan must preserve requested value")
assert(plan["change_state"] == "planned", "settings change plan must not claim persistence")
assert(!plan["apply_enabled"], "settings change plan must not enable apply before Runtime persistence exists")
assert(!plan["settings_persisted"], "settings change plan must not claim persistence")
assert(!plan["host_root_modified"], "settings change plan must not mutate host root")
assert(!plan["backend_details_exposed"], "settings change plan must hide backend details")
assert(plan["user_confirmation_required"], "settings change plan must require user confirmation")
assert(!plan["snapshot_recommended"], "resource access changes should not require snapshots by default")
assert(plan["portal_policy_review_required"], "resource access changes must require Portal policy review")
assert(!plan["runtime_restart_required"], "settings change planning must not require a Runtime restart")
assert(plan["affected_policy"]["options"].include?("ask"), "settings change plan must expose safe user options")
assert(plan["steps"].map { |step| step.fetch("id") } == %w[validate-setting review-user-confirmation review-portal-policy prepare-restore-point persist-runtime-setting], "settings change plan must expose expected steps")
assert(plan["blocked_actions"].include?("grant desktop resources without Portal policy review"), "settings change plan must block direct resource grants")

mode_plan = Xnix::Compatibility::SettingsChangePlan.new(
  application_id: "org.xnix.sample.notepad",
  section_id: "run-mode",
  field_id: "preference",
  value: "performance"
).to_h
assert(mode_plan["snapshot_recommended"], "run mode preference changes should recommend snapshots")
assert(!mode_plan["portal_policy_review_required"], "run mode preference changes should not require Portal policy review")

json = JSON.pretty_generate(plan)
assert(!json.match?(/wine|prefix|\.wine|proton|virtual machine/i), "settings change plan must not expose backend implementation terms")
assert(!json.match?(%r{/Users|/home|/var|/opt|/tmp}), "settings change plan must not expose host storage paths")

stdout, stderr, status = Open3.capture3(
  "ruby",
  project_root.join("bin/xnix-compat-settings-change").to_s,
  "--app",
  "org.xnix.sample.notepad",
  "--section",
  "resource-access",
  "--field",
  "documents",
  "--value",
  "allow"
)
assert(status.success?, "settings change CLI must exit successfully: #{stderr}")
assert(JSON.parse(stdout) == plan, "settings change CLI must emit the plan")

_stdout, stderr, status = Open3.capture3(
  "ruby",
  project_root.join("bin/xnix-compat-settings-change").to_s,
  "--app",
  "org.xnix.sample.notepad",
  "--section",
  "resource-access",
  "--field",
  "documents",
  "--value",
  "unsafe"
)
assert(!status.success?, "settings change CLI must reject unsupported values")
assert(stderr.include?("unsupported settings value"), "settings change CLI must explain unsupported values")

puts "PASS: compatibility settings change plan unit tests"
