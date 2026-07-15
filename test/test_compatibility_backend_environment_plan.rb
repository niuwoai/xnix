#!/usr/bin/env ruby
# frozen_string_literal: true

require "json"
require "open3"
require_relative "../lib/xnix/compatibility/compatibility_backend_environment_plan"
require_relative "../lib/xnix/compatibility/recipe_store"
require_relative "../lib/xnix/compatibility/runtime_daemon"

def assert(condition, message)
  return if condition

  warn "FAIL: #{message}"
  exit 1
end

recipe_store = Xnix::Compatibility::RecipeStore.new(path: Xnix::Compatibility::RuntimeDaemon::DEFAULT_RECIPE_DIR)
recipe = recipe_store.find("org.xnix.sample.notepad")
plan = Xnix::Compatibility::CompatibilityBackendEnvironmentPlan.new(recipe: recipe).to_h
profile_ids = plan.fetch("profiles").map { |profile| profile.fetch("id") }

assert(plan["version"] == "0.2.208", "backend environment plan must expose the current version")
assert(plan["plan_type"] == "compatibility-backend-environment-plan", "backend environment plan must identify the plan type")
assert(plan["runtime_method"] == "GetBackendEnvironmentPlan", "backend environment plan must identify the Runtime method")
assert(plan["application"]["id"] == "org.xnix.sample.notepad", "backend environment plan must preserve the application id")
assert(plan["runtime_owned"], "Runtime must own backend environments")
assert(plan["c_runtime_backed"], "backend environment plan must be marked as C Runtime-backed")
assert(!plan["kde_policy_owner"], "KDE must not own backend environments")
assert(plan["environment_state"] == "planned", "backend environment plan must stay planned")
assert(profile_ids == %w[local-compatibility-environment isolated-compatibility-environment], "backend environment plan must expose expected profiles")
assert(plan["profiles"].all? { |profile| profile["status"] == "blocked" }, "backend environment profiles must remain blocked")
assert(!plan["local_environment_ready"], "backend environment plan must not claim local readiness")
assert(!plan["isolated_environment_ready"], "backend environment plan must not claim isolated readiness")
assert(!plan["environment_created"], "backend environment plan must not create environments")
assert(!plan["backend_process_started"], "backend environment plan must not start backend processes")
assert(!plan["host_storage_exposed"], "backend environment plan must not expose host storage")
assert(!plan["clipboard_bridge_enabled"], "backend environment plan must not enable clipboard bridges")
assert(!plan["print_bridge_enabled"], "backend environment plan must not enable print bridges")
assert(plan["portal_review_required"], "backend environment plan must require Portal review")
assert(plan["snapshot_required"], "backend environment plan must require snapshot review")
assert(!plan["launch_enabled"], "backend environment plan must not enable launch")
assert(plan["required_reviews"] == %w[package-source-review application-state-root-review portal-policy-review snapshot-baseline-review], "backend environment plan must expose required reviews")
assert(plan["blocked_actions"].include?("create local compatibility environment from KDE"), "backend environment plan must block local environment creation from KDE")
assert(plan["blocked_actions"].include?("create isolated compatibility environment from KDE"), "backend environment plan must block isolated environment creation from KDE")
assert(!plan["host_root_modified"], "backend environment plan must not mutate the host root")
assert(!plan["network_required"], "backend environment plan must not require network access")
assert(!plan["backend_details_exposed"], "backend environment plan must hide backend details")

json = JSON.pretty_generate(plan)
assert(!json.match?(/wine|prefix|\.wine|proton|virtual machine/i), "backend environment plan must not expose backend implementation terms")
assert(!json.match?(%r{/Users|/home|/var|/opt|/tmp}), "backend environment plan must not expose host storage paths")

stdout, stderr, status = Open3.capture3(
  "ruby",
  File.expand_path("../bin/xnix-compat-backend-environment-plan", __dir__),
  "--app",
  "org.xnix.sample.notepad"
)
assert(status.success?, "backend environment plan CLI must exit successfully: #{stderr}")
assert(JSON.parse(stdout) == plan, "backend environment plan CLI must emit the plan model")

puts "PASS: compatibility backend environment plan unit tests"
