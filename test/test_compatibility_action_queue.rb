#!/usr/bin/env ruby
# frozen_string_literal: true

require "json"
require "open3"
require "pathname"
require_relative "../lib/xnix/compatibility/compatibility_action_queue"
require_relative "../lib/xnix/compatibility/recipe_store"

def assert(condition, message)
  return if condition

  warn "FAIL: #{message}"
  exit 1
end

project_root = Pathname.new(__dir__).join("..").realpath
recipe = Xnix::Compatibility::RecipeStore.new(path: project_root.join("runtime/recipes")).find("org.xnix.sample.notepad")
queue = Xnix::Compatibility::CompatibilityActionQueue.new(recipe: recipe).to_h

assert(queue["version"] == "0.2.247", "Compatibility action queue must expose the current version")
assert(queue["queue_type"] == "compatibility-center-action-queue", "Compatibility action queue must identify its type")
assert(queue["application"]["id"] == "org.xnix.sample.notepad", "Compatibility action queue must preserve the application id")
assert(queue["runtime_owned"], "Runtime must own Compatibility Center action queues")
assert(!queue["kde_policy_owner"], "KDE must not own action queue policy")
assert(queue["surface"] == "Compatibility Center", "Compatibility action queue must target the Compatibility Center")
assert(queue["action_count"] == 5, "Compatibility action queue must include expected actions")
assert(queue["pending_action_count"] == 5, "Compatibility action queue must count pending actions")
assert(queue["user_review_required_count"] == 3, "Compatibility action queue must count user review actions")
assert(!queue["execution_enabled"], "Compatibility action queue must not enable execution")
assert(!queue["repair_execution_enabled"], "Compatibility action queue must not enable repair execution")
assert(!queue["settings_persistence_enabled"], "Compatibility action queue must not enable settings persistence")
assert(!queue["host_root_modified"], "Compatibility action queue must not mutate the host root")
assert(!queue["network_required"], "Compatibility action queue must not require network access")
assert(!queue["backend_details_exposed"], "Compatibility action queue must hide backend details")

action_ids = queue["actions"].map { |action| action.fetch("id") }
assert(action_ids == %w[review-install-readiness review-settings-change review-ai-repair verify-runtime-service review-portal-policy], "Compatibility action queue must expose expected action order")
assert(queue["actions"].all? { |action| action["execution_enabled"] == false }, "Compatibility action queue actions must not be executable")
assert(queue["actions"].all? { |action| action["backend_details_exposed"] == false }, "Compatibility action queue actions must hide backend details")
assert(queue["actions"].any? { |action| action["status"] == "approval-required" }, "Compatibility action queue must include approval-required actions")
assert(queue["blocked_actions"].include?("execute queued actions from KDE without Runtime approval"), "Compatibility action queue must block direct KDE execution")
assert(queue["blocked_actions"].include?("grant desktop resources without XDG Desktop Portal review"), "Compatibility action queue must block direct resource grants")

json = JSON.pretty_generate(queue)
assert(!json.match?(/wine|prefix|\.wine|proton|virtual machine/i), "Compatibility action queue must not expose backend implementation terms")
assert(!json.match?(%r{/Users|/home|/var|/opt|/tmp}), "Compatibility action queue must not expose host storage paths")

stdout, stderr, status = Open3.capture3(
  "ruby",
  project_root.join("bin/xnix-compat-action-queue").to_s,
  "--app",
  "org.xnix.sample.notepad"
)
assert(status.success?, "Compatibility action queue CLI must exit successfully: #{stderr}")
assert(JSON.parse(stdout) == queue, "Compatibility action queue CLI must emit the queue")

_stdout, stderr, status = Open3.capture3(
  "ruby",
  project_root.join("bin/xnix-compat-action-queue").to_s,
  "--app",
  "org.xnix.missing"
)
assert(!status.success?, "Compatibility action queue CLI must reject unknown applications")
assert(stderr.include?("unknown application"), "Compatibility action queue CLI must explain unknown applications")

puts "PASS: Compatibility Center action queue unit tests"
