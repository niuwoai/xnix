# frozen_string_literal: true

require "json"
require "open3"
require "pathname"
require_relative "../lib/xnix/compatibility/compatibility_mode_switch_plan"
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
plan = Xnix::Compatibility::CompatibilityModeSwitchPlan.new(
  recipe: recipe,
  requested_mode: "prefer-compatibility"
).to_h

assert(plan["version"] == "0.2.223", "compatibility mode switch plan must expose the current version")
assert(plan["plan_type"] == "compatibility-mode-switch-plan", "compatibility mode switch plan must identify the plan type")
assert(plan["runtime_method"] == "GetCompatibilityModeSwitchPlan", "compatibility mode switch plan must identify the Runtime method")
assert(plan["application"]["id"] == "org.xnix.sample.notepad", "compatibility mode switch plan must preserve the application id")
assert(plan["current_mode"] == "automatic", "compatibility mode switch plan must expose the current mode")
assert(plan["requested_mode"] == "prefer-compatibility", "compatibility mode switch plan must expose the requested mode")
assert(plan["mode_state"] == "planned", "compatibility mode switch plan must stay planned")
assert(plan["mode_count"] == 4, "compatibility mode switch plan must expose four user-facing modes")
assert(plan["modes"].map { |mode| mode.fetch("id") } == %w[automatic prefer-performance prefer-compatibility isolated-execution], "compatibility mode switch plan must preserve mode order")
assert(plan["modes"].one? { |mode| mode.fetch("selected") }, "compatibility mode switch plan must mark one selected mode")
assert(plan["modes"].one? { |mode| mode.fetch("requested") }, "compatibility mode switch plan must mark one requested mode")
assert(plan["valid_mode"], "compatibility mode switch plan must accept supported modes")
assert(plan["requires_user_confirmation"], "compatibility mode switch plan must require user confirmation")
assert(plan["portal_review_required"], "compatibility mode switch plan must require Portal review")
assert(plan["snapshot_required"], "compatibility mode switch plan must require snapshots")
assert(!plan["settings_persistence_enabled"], "compatibility mode switch plan must not persist settings")
assert(!plan["backend_reconfiguration_enabled"], "compatibility mode switch plan must not reconfigure backends")
assert(!plan["backend_process_started"], "compatibility mode switch plan must not start backend processes")
assert(!plan["launch_enabled"], "compatibility mode switch plan must not enable launch")
assert(!plan["host_root_modified"], "compatibility mode switch plan must not mutate the host root")
assert(!plan["backend_details_exposed"], "compatibility mode switch plan must not expose backend details")
assert_no_backend_terms(plan, "compatibility mode switch plan must avoid backend implementation terms")

stdout, stderr, status = Open3.capture3(
  "ruby",
  PROJECT_ROOT.join("bin/xnix-compat-mode-switch-plan").to_s,
  "--app",
  "org.xnix.sample.notepad",
  "--mode",
  "isolated-execution",
  "--recipe-dir",
  PROJECT_ROOT.join("runtime/recipes").to_s
)
assert(status.success?, "compatibility mode switch CLI must exit successfully: #{stderr}")
cli_plan = JSON.parse(stdout)
assert(cli_plan["requested_mode"] == "isolated-execution", "compatibility mode switch CLI must preserve requested mode")
assert(!cli_plan["settings_persistence_enabled"], "compatibility mode switch CLI must preserve persistence gate")
assert_no_backend_terms(cli_plan, "compatibility mode switch CLI must avoid backend implementation terms")

_stdout, stderr, status = Open3.capture3(
  "ruby",
  PROJECT_ROOT.join("bin/xnix-compat-mode-switch-plan").to_s,
  "--app",
  "org.xnix.sample.notepad",
  "--mode",
  "unsupported-mode",
  "--recipe-dir",
  PROJECT_ROOT.join("runtime/recipes").to_s
)
assert(!status.success?, "compatibility mode switch CLI must reject unsupported modes")
assert(stderr.include?("unsupported compatibility mode"), "compatibility mode switch CLI must explain unsupported modes")

core_source = PROJECT_ROOT.join("runtime/core/xnix_runtime_core.c")
core_cli_source = PROJECT_ROOT.join("runtime/core/xnix_runtime_core_cli.c")
core_binary = PROJECT_ROOT.join("tmp/xnix-runtime-core-mode-switch-test")
core_binary.dirname.mkpath

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
assert(status.success?, "C Runtime core CLI must compile for compatibility mode switch plans: #{stderr}")

stdout, stderr, status = Open3.capture3(
  core_binary.to_s,
  "compatibility-mode-switch-plan",
  "org.xnix.sample.notepad",
  "prefer-performance"
)
assert(status.success?, "C Runtime core CLI must emit compatibility mode switch plans: #{stderr}")
c_plan = JSON.parse(stdout)
assert(c_plan["runtime_method"] == "GetCompatibilityModeSwitchPlan", "C Runtime core CLI must expose the Runtime method")
assert(c_plan["mode_count"] == 4, "C Runtime core CLI must expose four modes")
assert(c_plan["requested_mode"] == "prefer-performance", "C Runtime core CLI must preserve requested mode")
assert(!c_plan["backend_process_started"], "C Runtime core CLI must not start backend processes")
assert(!c_plan["backend_details_exposed"], "C Runtime core CLI must not expose backend details")
assert_no_backend_terms(c_plan, "C Runtime core CLI must avoid backend implementation terms")

puts "PASS: compatibility mode switch plan unit tests"
