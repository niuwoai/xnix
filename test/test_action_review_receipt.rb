#!/usr/bin/env ruby
# frozen_string_literal: true

require "json"
require "open3"
require "pathname"
require_relative "../lib/xnix/compatibility/action_review_receipt"
require_relative "../lib/xnix/compatibility/recipe_store"

def assert(condition, message)
  return if condition

  warn "FAIL: #{message}"
  exit 1
end

project_root = Pathname.new(__dir__).join("..").realpath
recipe = Xnix::Compatibility::RecipeStore.new(path: project_root.join("runtime/recipes")).find("org.xnix.sample.notepad")
receipt = Xnix::Compatibility::ActionReviewReceipt.new(
  recipe: recipe,
  action_id: "review-ai-repair",
  decision: "approved"
).to_h

assert(receipt["version"] == "0.2.99", "action review receipt must expose the current version")
assert(receipt["receipt_type"] == "compatibility-center-action-review-receipt", "action review receipt must identify receipt type")
assert(receipt["queue_type"] == "compatibility-center-action-queue", "action review receipt must reference the action queue")
assert(receipt["application"]["id"] == "org.xnix.sample.notepad", "action review receipt must preserve application id")
assert(receipt["action"]["id"] == "review-ai-repair", "action review receipt must preserve action id")
assert(receipt["decision"] == "approved", "action review receipt must preserve decision")
assert(receipt["decision_recorded"], "action review receipt must record review intent")
assert(receipt["runtime_owned"], "Runtime must own action review receipts")
assert(!receipt["kde_policy_owner"], "KDE must not own review receipt policy")
assert(!receipt["execution_enabled"], "action review receipt must not enable action execution")
assert(!receipt["repair_execution_enabled"], "action review receipt must not enable repair execution")
assert(!receipt["settings_persistence_enabled"], "action review receipt must not enable settings persistence")
assert(!receipt["resource_grant_created"], "action review receipt must not create resource grants")
assert(!receipt["host_root_modified"], "action review receipt must not mutate host root")
assert(!receipt["network_required"], "action review receipt must not require network")
assert(!receipt["backend_details_exposed"], "action review receipt must hide backend details")
assert(receipt["required_runtime_gate"] == "ai-repair-approval", "action review receipt must preserve the Runtime gate")
assert(receipt["blocked_actions"].include?("treat KDE review intent as Runtime execution approval"), "action review receipt must block treating review as execution")

deferred = Xnix::Compatibility::ActionReviewReceipt.new(
  recipe: recipe,
  action_id: "review-settings-change",
  decision: "deferred"
).to_h
assert(deferred["next_step"].include?("deferred"), "deferred review receipts must stay deferred")

json = JSON.pretty_generate(receipt)
assert(!json.match?(/wine|prefix|\.wine|proton|virtual machine/i), "action review receipt must not expose backend implementation terms")
assert(!json.match?(%r{/Users|/home|/var|/opt|/tmp}), "action review receipt must not expose host storage paths")

stdout, stderr, status = Open3.capture3(
  "ruby",
  project_root.join("bin/xnix-compat-action-review").to_s,
  "--app",
  "org.xnix.sample.notepad",
  "--action",
  "review-ai-repair",
  "--decision",
  "approved"
)
assert(status.success?, "action review receipt CLI must exit successfully: #{stderr}")
assert(JSON.parse(stdout) == receipt, "action review receipt CLI must emit the receipt")

_stdout, stderr, status = Open3.capture3(
  "ruby",
  project_root.join("bin/xnix-compat-action-review").to_s,
  "--app",
  "org.xnix.sample.notepad",
  "--action",
  "review-ai-repair",
  "--decision",
  "execute"
)
assert(!status.success?, "action review receipt CLI must reject unsupported decisions")
assert(stderr.include?("decision must be one of"), "action review receipt CLI must explain unsupported decisions")

puts "PASS: Compatibility Center action review receipt unit tests"
