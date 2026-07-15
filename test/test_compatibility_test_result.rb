#!/usr/bin/env ruby
# frozen_string_literal: true

require "json"
require "open3"
require "pathname"
require_relative "../lib/xnix/compatibility/compatibility_test_result"
require_relative "../lib/xnix/compatibility/recipe_store"

def assert(condition, message)
  return if condition

  warn "FAIL: #{message}"
  exit 1
end

project_root = Pathname.new(__dir__).join("..").realpath
recipe = Xnix::Compatibility::RecipeStore.new(path: project_root.join("runtime/recipes")).find("org.xnix.sample.notepad")
result = Xnix::Compatibility::CompatibilityTestResult.new(recipe: recipe).to_h
step_ids = result.fetch("step_results").map { |step| step.fetch("id") }

assert(result["version"] == "0.2.237", "compatibility test result must expose the current version")
assert(result["result_type"] == "compatibility-test-result", "compatibility test result must identify result type")
assert(result["plan_type"] == "compatibility-test", "compatibility test result must reference the test plan")
assert(result["test_type"] == "preflight", "compatibility test result must default to preflight")
assert(result["application_id"] == "org.xnix.sample.notepad", "compatibility test result must preserve the application id")
assert(result["runtime_method"] == "GetTestResult", "compatibility test result must expose the Runtime method")
assert(result["runtime_owned"], "Runtime must own compatibility test results")
assert(result["c_runtime_backed"], "compatibility test result must be C Runtime-backed")
assert(!result["kde_policy_owner"], "KDE must not own compatibility test results")
assert(result["execution_state"] == "waiting-for-runtime", "test result must not claim execution completion before backend binding")
assert(result["overall_status"] == "pending", "test result must stay pending while Runtime work remains")
assert(result["counts"] == { "total" => 4, "passed" => 1, "pending" => 3, "blocked" => 0 }, "test result must count step outcomes")
assert(step_ids == %w[recipe-validation portal-preflight snapshot-preflight runtime-launch-binding], "test result must mirror planned steps")
assert(result["step_results"].any? { |step| step["id"] == "recipe-validation" && step["result"] == "passed" }, "recipe validation must be marked passed")
assert(result["step_results"].any? { |step| step["id"] == "runtime-launch-binding" && step["result"] == "not-run" }, "launch binding must remain not-run")
assert(result["step_results"].all? { |step| step["evidence"].is_a?(Array) && !step["evidence"].empty? }, "test result must include diagnostic evidence")
assert(result["artifacts"]["compatibility_center_card"], "test result must feed the Compatibility Center")
assert(result["artifacts"]["diagnostics_record"], "test result must feed diagnostics")
assert(result["artifacts"]["repair_plan_issue"] == "engine-binding-pending", "test result must map pending execution to repair planning")
assert(result["safe_for_ai_diagnostics"], "test result must be safe for AI diagnostics")
assert(!result["test_executed"], "compatibility test result must not execute tests")
assert(!result["host_root_modified"], "compatibility test result must not mutate the host root")
assert(!result["backend_details_exposed"], "compatibility test result must hide backend details")

smoke_result = Xnix::Compatibility::CompatibilityTestResult.new(recipe: recipe, test_type: "smoke").to_h
assert(smoke_result["test_type"] == "smoke", "compatibility test result must preserve explicit test type")

json = JSON.pretty_generate(result)
assert(!json.match?(/wine|prefix|\.wine|proton|virtual machine/i), "compatibility test result must not expose backend implementation terms")

stdout, stderr, status = Open3.capture3(
  "ruby",
  project_root.join("bin/xnix-compat-test-result").to_s,
  "--app",
  "org.xnix.sample.notepad"
)
assert(status.success?, "compatibility test result CLI must exit successfully: #{stderr}")
assert(JSON.parse(stdout) == result, "compatibility test result CLI must emit the result")

_stdout, stderr, status = Open3.capture3(
  "ruby",
  project_root.join("bin/xnix-compat-test-result").to_s,
  "--app",
  "org.xnix.sample.notepad",
  "--test-type",
  "unknown"
)
assert(!status.success?, "compatibility test result CLI must reject unknown test types")
assert(stderr.include?("invalid argument"), "compatibility test result CLI must explain invalid test types")

puts "PASS: compatibility test result unit tests"
