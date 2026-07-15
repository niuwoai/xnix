# frozen_string_literal: true

require "json"
require "open3"
require "pathname"
require "fileutils"
require_relative "../lib/xnix/compatibility/compatibility_backend_selection_plan"
require_relative "../lib/xnix/compatibility/recipe_store"

PROJECT_ROOT = Pathname.new(__dir__).join("..").realpath

def assert(condition, message)
  raise "FAIL: #{message}" unless condition
end

def assert_no_backend_terms(value, message)
  serialized = JSON.generate(value)
  assert(!serialized.match?(/wine|prefix|\.wine|proton|virtual machine/i), message)
end

recipe = Xnix::Compatibility::RecipeStore.new(path: PROJECT_ROOT.join("runtime/recipes")).find("org.xnix.sample.notepad")
plan = Xnix::Compatibility::CompatibilityBackendSelectionPlan.new(recipe: recipe).to_h

assert(plan["version"] == "0.2.206", "backend selection plan must expose the current version")
assert(plan["plan_type"] == "compatibility-backend-selection-plan", "backend selection plan must identify the plan type")
assert(plan["runtime_method"] == "GetBackendSelectionPlan", "backend selection plan must expose the Runtime method")
assert(plan["application"]["id"] == "org.xnix.sample.notepad", "backend selection plan must identify the application")
assert(plan["recommended_profile_id"] == "local-compatibility", "backend selection plan must recommend a desktop-safe local profile for the sample recipe")
assert(plan["candidate_count"] == 2, "backend selection plan must expose two compatibility candidates")
assert(plan["ready_candidate_count"] == 0, "backend selection plan must not mark candidates ready before Runtime gates pass")
assert(plan["blocked_candidate_count"] == 2, "backend selection plan must keep both candidates blocked before Runtime gates pass")
assert(plan["candidate_profiles"].map { |candidate| candidate.fetch("id") } == %w[local-compatibility isolated-compatibility], "backend selection plan must preserve candidate order")
assert(plan["candidate_profiles"].one? { |candidate| candidate.fetch("recommended") }, "backend selection plan must recommend exactly one candidate")
assert(plan["candidate_profiles"].all? { |candidate| candidate.fetch("required_preflight").include?("portal-policy-review") }, "backend selection plan must require Portal preflight")
assert(plan["required_reviews"].include?("backend-capability-review"), "backend selection plan must require capability review")
assert(plan["required_reviews"].include?("snapshot-baseline-review"), "backend selection plan must require snapshot review")
assert(plan["runtime_owned"], "backend selection plan must be Runtime-owned")
assert(plan["c_runtime_backed"], "backend selection plan must be C Runtime-backed")
assert(!plan["kde_policy_owner"], "backend selection plan must not make KDE the policy owner")
assert(!plan["selection_committed"], "backend selection plan must not commit backend selection")
assert(!plan["selection_change_enabled"], "backend selection plan must not enable selection changes")
assert(!plan["backend_launch_enabled"], "backend selection plan must not enable backend launch")
assert(!plan["capability_activation_enabled"], "backend selection plan must not enable capability activation")
assert(!plan["environment_created"], "backend selection plan must not create environments")
assert(!plan["request_object_created"], "backend selection plan must not create request objects")
assert(!plan["state_root_created"], "backend selection plan must not create state roots")
assert(!plan["snapshot_created"], "backend selection plan must not create snapshots")
assert(!plan["host_root_modified"], "backend selection plan must not mutate the host root")
assert(!plan["privileged_container_required"], "backend selection plan must not require privileged containers")
assert(!plan["backend_details_exposed"], "backend selection plan must not expose backend details")
assert_no_backend_terms(plan, "backend selection plan must avoid backend implementation terms")

stdout, stderr, status = Open3.capture3("ruby", PROJECT_ROOT.join("bin/xnix-compat-backend-selection-plan").to_s, "--app", "org.xnix.sample.notepad")
assert(status.success?, "backend selection plan CLI must exit successfully: #{stderr}")
cli_plan = JSON.parse(stdout)
assert(cli_plan["plan_type"] == "compatibility-backend-selection-plan", "backend selection plan CLI must emit the plan")
assert(cli_plan["candidate_count"] == 2, "backend selection plan CLI must preserve candidate count")
assert(!cli_plan["selection_committed"], "backend selection plan CLI must keep selection uncommitted")
assert_no_backend_terms(cli_plan, "backend selection plan CLI must avoid backend implementation terms")

core_source = PROJECT_ROOT.join("runtime/core/xnix_runtime_core.c")
core_cli_source = PROJECT_ROOT.join("runtime/core/xnix_runtime_core_cli.c")
core_binary = PROJECT_ROOT.join("tmp/xnix-runtime-core-backend-selection-plan-test")
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
assert(status.success?, "C Runtime core CLI must compile for backend selection plan: #{stderr}")

stdout, stderr, status = Open3.capture3(core_binary.to_s, "backend-selection-plan", "org.xnix.sample.notepad")
assert(status.success?, "C Runtime core CLI must emit backend selection plan: #{stderr}")
c_plan = JSON.parse(stdout)
assert(c_plan["runtime_method"] == "GetBackendSelectionPlan", "C Runtime core CLI must expose the Runtime method")
assert(c_plan["candidate_count"] == 2, "C Runtime core CLI must expose two candidates")
assert(c_plan["blocked_candidate_count"] == 2, "C Runtime core CLI must keep candidates blocked")
assert(c_plan["recommended_profile_id"] == "local-compatibility", "C Runtime core CLI must recommend a profile")
assert(!c_plan["selection_committed"], "C Runtime core CLI must not commit backend selection")
assert(!c_plan["selection_change_enabled"], "C Runtime core CLI must not enable selection changes")
assert(!c_plan["backend_launch_enabled"], "C Runtime core CLI must not enable backend launch")
assert(!c_plan["capability_activation_enabled"], "C Runtime core CLI must not enable capability activation")
assert(!c_plan["environment_created"], "C Runtime core CLI must not create environments")
assert(!c_plan["request_object_created"], "C Runtime core CLI must not create request objects")
assert(!c_plan["state_root_created"], "C Runtime core CLI must not create state roots")
assert(!c_plan["snapshot_created"], "C Runtime core CLI must not create snapshots")
assert(!c_plan["host_root_modified"], "C Runtime core CLI must not mutate the host root")
assert(!c_plan["privileged_container_required"], "C Runtime core CLI must not require privileged containers")
assert(!c_plan["backend_details_exposed"], "C Runtime core CLI must not expose backend details")
assert_no_backend_terms(c_plan, "C Runtime core CLI must avoid backend implementation terms")

FileUtils.rm_f(core_binary)

puts "PASS: compatibility backend selection plan unit tests"
