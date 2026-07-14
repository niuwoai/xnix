# frozen_string_literal: true

require "json"
require "open3"
require "pathname"
require "fileutils"
require_relative "../lib/xnix/compatibility/compatibility_permission_review_plan"
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
plan = Xnix::Compatibility::CompatibilityPermissionReviewPlan.new(recipe: recipe).to_h

assert(plan["version"] == "0.2.129", "compatibility permission review plan must expose the current version")
assert(plan["plan_type"] == "compatibility-permission-review-plan", "compatibility permission review plan must identify the plan type")
assert(plan["runtime_method"] == "GetCompatibilityPermissionReviewPlan", "compatibility permission review plan must identify the Runtime method")
assert(plan["application"]["id"] == "org.xnix.sample.notepad", "compatibility permission review plan must preserve the application id")
assert(plan["review_state"] == "planned", "compatibility permission review plan must stay planned")
assert(plan["permission_count"] == 7, "compatibility permission review plan must expose seven user-facing permissions")
assert(plan["allow_count"] == 1, "compatibility permission review plan must count allowed defaults")
assert(plan["ask_count"] == 5, "compatibility permission review plan must count ask defaults")
assert(plan["deny_count"] == 1, "compatibility permission review plan must count denied defaults")
assert(plan["permissions"].map { |permission| permission.fetch("id") } == %w[documents downloads camera network clipboard print screenshot], "compatibility permission review plan must preserve permission order")
assert(plan["permissions"].find { |permission| permission.fetch("id") == "network" }.fetch("decision") == "allow", "compatibility permission review plan must expose network defaults")
assert(plan["permissions"].find { |permission| permission.fetch("id") == "camera" }.fetch("decision") == "deny", "compatibility permission review plan must expose camera defaults")
assert(plan["permissions"].all? { |permission| !permission.fetch("change_pending") }, "compatibility permission review plan must not mark pending changes")
assert(plan["permissions"].all? { |permission| !permission.fetch("request_object_created") }, "compatibility permission review plan must not create request objects")
assert(plan["permissions"].all? { |permission| !permission.fetch("permission_granted") }, "compatibility permission review plan must not grant permissions")
assert(plan["user_review_required"], "compatibility permission review plan must require user review")
assert(plan["portal_review_required"], "compatibility permission review plan must require Portal review")
assert(!plan["permission_changes_applied"], "compatibility permission review plan must not apply permission changes")
assert(!plan["request_objects_created"], "compatibility permission review plan must not create request objects")
assert(!plan["permissions_granted"], "compatibility permission review plan must not grant permissions")
assert(!plan["settings_persisted"], "compatibility permission review plan must not persist settings")
assert(!plan["host_permission_changed"], "compatibility permission review plan must not change host permissions")
assert(!plan["host_root_modified"], "compatibility permission review plan must not mutate the host root")
assert(!plan["backend_details_exposed"], "compatibility permission review plan must not expose backend details")
assert_no_backend_terms(plan, "compatibility permission review plan must avoid backend implementation terms")

stdout, stderr, status = Open3.capture3(
  "ruby",
  PROJECT_ROOT.join("bin/xnix-compat-permission-review-plan").to_s,
  "--app",
  "org.xnix.sample.notepad",
  "--recipe-dir",
  PROJECT_ROOT.join("runtime/recipes").to_s
)
assert(status.success?, "compatibility permission review CLI must exit successfully: #{stderr}")
cli_plan = JSON.parse(stdout)
assert(cli_plan["plan_type"] == "compatibility-permission-review-plan", "compatibility permission review CLI must emit the plan")
assert(cli_plan["permission_count"] == 7, "compatibility permission review CLI must preserve permission count")
assert(!cli_plan["permissions_granted"], "compatibility permission review CLI must preserve permission grant gates")
assert_no_backend_terms(cli_plan, "compatibility permission review CLI must avoid backend implementation terms")

core_source = PROJECT_ROOT.join("runtime/core/xnix_runtime_core.c")
core_cli_source = PROJECT_ROOT.join("runtime/core/xnix_runtime_core_cli.c")
core_binary = PROJECT_ROOT.join("tmp/xnix-runtime-core-permission-review-test")
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
assert(status.success?, "C Runtime core CLI must compile for compatibility permission review plans: #{stderr}")

stdout, stderr, status = Open3.capture3(
  core_binary.to_s,
  "compatibility-permission-review-plan",
  "org.xnix.sample.notepad"
)
assert(status.success?, "C Runtime core CLI must emit compatibility permission review plans: #{stderr}")
c_plan = JSON.parse(stdout)
assert(c_plan["runtime_method"] == "GetCompatibilityPermissionReviewPlan", "C Runtime core CLI must expose the Runtime method")
assert(c_plan["permission_count"] == 7, "C Runtime core CLI must expose seven permissions")
assert(!c_plan["permission_changes_applied"], "C Runtime core CLI must not apply permission changes")
assert(!c_plan["permissions_granted"], "C Runtime core CLI must not grant permissions")
assert(!c_plan["host_permission_changed"], "C Runtime core CLI must not change host permissions")
assert(!c_plan["backend_details_exposed"], "C Runtime core CLI must not expose backend details")
assert_no_backend_terms(c_plan, "C Runtime core CLI must avoid backend implementation terms")

FileUtils.rm_f(core_binary)

puts "PASS: compatibility permission review plan unit tests"
