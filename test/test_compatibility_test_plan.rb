#!/usr/bin/env ruby
# frozen_string_literal: true

require "json"
require "open3"
require "pathname"
require_relative "../lib/xnix/compatibility/compatibility_test_plan"
require_relative "../lib/xnix/compatibility/recipe_store"

def assert(condition, message)
  return if condition

  warn "FAIL: #{message}"
  exit 1
end

project_root = Pathname.new(__dir__).join("..").realpath
recipe = Xnix::Compatibility::RecipeStore.new(path: project_root.join("runtime/recipes")).find("org.xnix.sample.notepad")
plan = Xnix::Compatibility::CompatibilityTestPlan.new(recipe: recipe).to_h
step_ids = plan.fetch("steps").map { |step| step.fetch("id") }

assert(plan["version"] == "0.2.98", "compatibility test plan must expose the current version")
assert(plan["plan_type"] == "compatibility-test", "compatibility test plan must identify the plan type")
assert(plan["test_type"] == "preflight", "compatibility test plan must default to preflight tests")
assert(plan["desktop"] == "KDE Plasma", "compatibility test plan must target KDE Plasma")
assert(plan["application_id"] == "org.xnix.sample.notepad", "compatibility test plan must preserve the application id")
assert(plan["runtime_method"] == "GetTestPlan", "compatibility test plan must expose the Runtime method")
assert(plan["runtime_owned"], "Runtime must own compatibility tests")
assert(plan["c_runtime_backed"], "compatibility test plan must be C Runtime-backed")
assert(!plan["kde_policy_owner"], "KDE must not own compatibility test policy")
assert(step_ids == %w[recipe-validation portal-preflight snapshot-preflight runtime-launch-binding], "compatibility test plan must include the expected steps")
assert(plan["steps"].first["status"] == "pass", "recipe validation must pass for bundled recipes")
assert(plan["steps"].any? { |step| step["id"] == "portal-preflight" && step["portal_request"]["method"] == "OpenFile" }, "test plan must include Portal request preflight")
assert(plan["steps"].any? { |step| step["id"] == "snapshot-preflight" && step["snapshot"]["restore_available"] }, "test plan must include snapshot preflight")
assert(plan["steps"].any? { |step| step["id"] == "runtime-launch-binding" && step["status"] == "pending" }, "test plan must keep launch backend binding pending")
assert(!plan["blocked"], "preflight test plan must not be blocked when file access can ask")
assert(plan["blocking_reasons"].empty?, "preflight test plan must not report blocking reasons")
assert(plan["artifacts"]["compatibility_center_card"], "test plan must feed the Compatibility Center")
assert(plan["artifacts"]["repair_plan_issue"] == "engine-binding-pending", "test plan must map pending launch binding to repair planning")
assert(!plan["execution_request_created"], "compatibility test plan must not create execution requests")
assert(!plan["test_executed"], "compatibility test plan must not execute tests")
assert(!plan["host_root_modified"], "compatibility test plan must not mutate the host root")
assert(!plan["backend_details_exposed"], "compatibility test plan must hide backend details")

smoke_plan = Xnix::Compatibility::CompatibilityTestPlan.new(recipe: recipe, test_type: "smoke").to_h
assert(smoke_plan["test_type"] == "smoke", "compatibility test plan must preserve explicit test type")

json = JSON.pretty_generate(plan)
assert(!json.match?(/wine|prefix|\.wine|proton|virtual machine/i), "compatibility test plan must not expose backend implementation terms")

stdout, stderr, result = Open3.capture3(
  "ruby",
  project_root.join("bin/xnix-compat-test-plan").to_s,
  "--app",
  "org.xnix.sample.notepad"
)
assert(result.success?, "compatibility test plan CLI must exit successfully: #{stderr}")
assert(JSON.parse(stdout) == plan, "compatibility test plan CLI must emit the plan")

_stdout, stderr, result = Open3.capture3(
  "ruby",
  project_root.join("bin/xnix-compat-test-plan").to_s,
  "--app",
  "org.xnix.sample.notepad",
  "--test-type",
  "unknown"
)
assert(!result.success?, "compatibility test plan CLI must reject unknown test types")
assert(stderr.include?("invalid argument"), "compatibility test plan CLI must explain invalid test types")

puts "PASS: compatibility test plan unit tests"
