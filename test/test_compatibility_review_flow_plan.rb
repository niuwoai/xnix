# frozen_string_literal: true

require "json"
require "open3"
require "pathname"
require "fileutils"
require_relative "../lib/xnix/compatibility/compatibility_review_flow_plan"
require_relative "../lib/xnix/compatibility/recipe_store"

PROJECT_ROOT = Pathname.new(__dir__).join("..").realpath

def assert(condition, message)
  raise "FAIL: #{message}" unless condition
end

def assert_no_backend_terms(value, message)
  serialized = JSON.generate(value)
  assert(!serialized.match?(/wine|prefix|\.wine|proton|virtual machine/i), message)
end

store = Xnix::Compatibility::RecipeStore.new(path: PROJECT_ROOT.join("runtime/recipes"))
recipe = store.find("org.xnix.sample.notepad")
plan = Xnix::Compatibility::CompatibilityReviewFlowPlan.new(recipe: recipe).to_h

assert(plan["version"] == "0.2.118", "compatibility review flow plan must expose the current version")
assert(plan["plan_type"] == "compatibility-review-flow-plan", "compatibility review flow plan must identify the plan type")
assert(plan["runtime_method"] == "GetCompatibilityReviewFlowPlan", "compatibility review flow plan must identify the Runtime method")
assert(plan["application"]["id"] == "org.xnix.sample.notepad", "compatibility review flow plan must preserve the application id")
assert(plan["review_state"] == "planned", "compatibility review flow plan must stay planned")
assert(plan["section_id"] == "resource-access", "compatibility review flow plan must expose the settings section")
assert(plan["field_id"] == "documents", "compatibility review flow plan must expose the settings field")
assert(plan["requested_value"] == "ask", "compatibility review flow plan must expose the requested value")
assert(plan["operation"] == "file-open", "compatibility review flow plan must expose the Portal operation")
assert(plan["step_count"] == 5, "compatibility review flow plan must expose five review steps")
assert(plan["required_review_count"] == 3, "compatibility review flow plan must count required review steps")
assert(plan["blocked_step_count"] == 1, "compatibility review flow plan must count blocked steps")
assert(plan["pending_step_count"] == 1, "compatibility review flow plan must count pending steps")
assert(plan["steps"].map { |step| step.fetch("id") } == %w[settings-change-review permission-review portal-request-review runtime-write-gate review-receipt], "compatibility review flow plan must preserve step order")
assert(plan["steps"].map { |step| step.fetch("runtime_method") }.include?("GetCompatibilitySettingsChangePlan"), "compatibility review flow plan must include settings change review")
assert(plan["steps"].map { |step| step.fetch("runtime_method") }.include?("GetCompatibilityPermissionReviewPlan"), "compatibility review flow plan must include permission review")
assert(plan["steps"].map { |step| step.fetch("runtime_method") }.include?("GetPortalRequestPlan"), "compatibility review flow plan must include Portal request review")
assert(plan["runtime_owned"], "compatibility review flow plan must be Runtime-owned")
assert(plan["c_runtime_backed"], "compatibility review flow plan must be C Runtime-backed")
assert(!plan["kde_policy_owner"], "compatibility review flow plan must not make KDE the policy owner")
assert(plan["user_confirmation_required"], "compatibility review flow plan must require user confirmation")
assert(plan["portal_policy_review_required"], "compatibility review flow plan must require Portal policy review")
assert(plan["settings_change_planned"], "compatibility review flow plan must include settings planning")
assert(plan["permission_review_planned"], "compatibility review flow plan must include permission planning")
assert(plan["portal_request_planned"], "compatibility review flow plan must include Portal request planning")
assert(plan["review_receipt_required"], "compatibility review flow plan must include review receipts")
assert(!plan["apply_enabled"], "compatibility review flow plan must not enable applying changes")
assert(!plan["request_object_created"], "compatibility review flow plan must not create request objects")
assert(!plan["permission_granted"], "compatibility review flow plan must not grant permissions")
assert(!plan["settings_persisted"], "compatibility review flow plan must not persist settings")
assert(!plan["execution_started"], "compatibility review flow plan must not start execution")
assert(!plan["host_root_modified"], "compatibility review flow plan must not mutate the host root")
assert(!plan["backend_details_exposed"], "compatibility review flow plan must not expose backend details")
assert_no_backend_terms(plan, "compatibility review flow plan must avoid backend implementation terms")

stdout, stderr, status = Open3.capture3(
  "ruby",
  PROJECT_ROOT.join("bin/xnix-compat-review-flow-plan").to_s,
  "--app",
  "org.xnix.sample.notepad",
  "--recipe-dir",
  PROJECT_ROOT.join("runtime/recipes").to_s
)
assert(status.success?, "compatibility review flow CLI must exit successfully: #{stderr}")
cli_plan = JSON.parse(stdout)
assert(cli_plan["plan_type"] == "compatibility-review-flow-plan", "compatibility review flow CLI must emit the plan")
assert(cli_plan["step_count"] == 5, "compatibility review flow CLI must preserve step count")
assert(!cli_plan["apply_enabled"], "compatibility review flow CLI must preserve apply gates")
assert_no_backend_terms(cli_plan, "compatibility review flow CLI must avoid backend implementation terms")

core_source = PROJECT_ROOT.join("runtime/core/xnix_runtime_core.c")
core_cli_source = PROJECT_ROOT.join("runtime/core/xnix_runtime_core_cli.c")
core_binary = PROJECT_ROOT.join("tmp/xnix-runtime-core-review-flow-test")
FileUtils.mkdir_p(core_binary.dirname)

_stdout, stderr, status = Open3.capture3(
  "cc",
  "-std=c99",
  "-Wall",
  "-Wextra",
  "-I",
  PROJECT_ROOT.join("runtime/core").to_s,
  core_source.to_s,
  core_cli_source.to_s,
  "-o",
  core_binary.to_s
)
assert(status.success?, "C Runtime core CLI must compile for compatibility review flow plans: #{stderr}")

stdout, stderr, status = Open3.capture3(
  core_binary.to_s,
  "compatibility-review-flow-plan",
  "org.xnix.sample.notepad",
  "resource-access",
  "documents",
  "ask",
  "file-open"
)
assert(status.success?, "C Runtime core CLI must emit compatibility review flow plans: #{stderr}")
c_plan = JSON.parse(stdout)
assert(c_plan["runtime_method"] == "GetCompatibilityReviewFlowPlan", "C Runtime core CLI must expose the Runtime method")
assert(c_plan["step_count"] == 5, "C Runtime core CLI must expose five steps")
assert(c_plan["required_review_count"] == 3, "C Runtime core CLI must expose required review counts")
assert(c_plan["blocked_step_count"] == 1, "C Runtime core CLI must expose blocked step counts")
assert(c_plan["pending_step_count"] == 1, "C Runtime core CLI must expose pending step counts")
assert(!c_plan["apply_enabled"], "C Runtime core CLI must not apply review flows")
assert(!c_plan["request_object_created"], "C Runtime core CLI must not create request objects")
assert(!c_plan["permission_granted"], "C Runtime core CLI must not grant permissions")
assert(!c_plan["settings_persisted"], "C Runtime core CLI must not persist settings")
assert(!c_plan["execution_started"], "C Runtime core CLI must not start execution")
assert(!c_plan["host_root_modified"], "C Runtime core CLI must not mutate the host root")
assert(!c_plan["backend_details_exposed"], "C Runtime core CLI must not expose backend details")
assert_no_backend_terms(c_plan, "C Runtime core CLI must avoid backend implementation terms")

FileUtils.rm_f(core_binary)

puts "PASS: compatibility review flow plan unit tests"
