#!/usr/bin/env ruby
# frozen_string_literal: true

require "json"
require "open3"
require "pathname"
require_relative "../lib/xnix/compatibility/application_recipe"
require_relative "../lib/xnix/compatibility/compatibility_run_plan"
require_relative "../lib/xnix/compatibility/recipe_store"

def assert(condition, message)
  return if condition

  warn "FAIL: #{message}"
  exit 1
end

def recipe_for(mode)
  labels = {
    "automatic" => ["org.example.automatic", "Automatic Example"],
    "wine" => ["org.example.local", "Local Example"],
    "vm" => ["org.example.isolated", "Isolated Example"]
  }
  id, name = labels.fetch(mode)

  Xnix::Compatibility::ApplicationRecipe.new(
    id: id,
    name: name,
    icon: "application-x-executable",
    mode: mode,
    supported_extensions: [".abc"]
  )
end

project_root = Pathname.new(__dir__).join("..").realpath
store = Xnix::Compatibility::RecipeStore.new(path: project_root.join("runtime/recipes"))
sample_recipe = store.find("org.xnix.sample.notepad")
sample_plan = Xnix::Compatibility::CompatibilityRunPlan.new(recipe: sample_recipe).to_h

assert(sample_plan["version"] == "0.2.180", "compatibility run plan must expose the current version")
assert(sample_plan["plan_type"] == "compatibility-run", "compatibility run plan must identify the model type")
assert(sample_plan["application"]["id"] == "org.xnix.sample.notepad", "compatibility run plan must identify the application")
assert(sample_plan["execution"]["strategy"] == "automatic-managed", "automatic recipes must use managed automatic strategy")
assert(sample_plan["execution"]["engine"]["engine_id"] == "automatic-managed", "automatic recipes must include selected engine summary")
assert(!sample_plan["execution"]["engine"]["ready"], "selected engines must not claim readiness yet")
assert(!sample_plan["execution"]["backend_details_exposed"], "compatibility run plan must hide backend details")
assert(!sample_plan["execution"]["backend_binding"]["ready"], "compatibility run plan must not claim backend readiness yet")
assert(!sample_plan["execution"]["backend_binding"]["launch_enabled"], "compatibility run plan must not claim launch enablement yet")
assert(sample_plan["preflight"]["portal_policy_required"], "compatibility run plan must require Portal policy preflight")
assert(sample_plan["preflight"]["snapshot_before_risky_change"], "compatibility run plan must require snapshots before risky changes")
assert(sample_plan["preflight"]["diagnostics_required"], "compatibility run plan must require diagnostics")

local_plan = Xnix::Compatibility::CompatibilityRunPlan.new(recipe: recipe_for("wine")).to_h
isolated_plan = Xnix::Compatibility::CompatibilityRunPlan.new(recipe: recipe_for("vm")).to_h
assert(local_plan["execution"]["strategy"] == "local-compatibility-engine", "local recipes must map to local compatibility strategy")
assert(isolated_plan["execution"]["strategy"] == "isolated-compatibility-engine", "isolated recipes must map to isolated compatibility strategy")

[sample_plan, local_plan, isolated_plan].each do |plan|
  json = JSON.pretty_generate(plan)
  assert(!json.match?(/wine|prefix|\.wine|proton|virtual machine/i), "compatibility run plan must not expose backend implementation terms")
end

stdout, stderr, status = Open3.capture3(
  "ruby",
  project_root.join("bin/xnix-compat-run-plan").to_s,
  "--app",
  "org.xnix.sample.notepad"
)
assert(status.success?, "compatibility run plan CLI must exit successfully: #{stderr}")
assert(JSON.parse(stdout) == sample_plan, "compatibility run plan CLI must emit the plan")

_stdout, stderr, status = Open3.capture3("ruby", project_root.join("bin/xnix-compat-run-plan").to_s)
assert(!status.success?, "compatibility run plan CLI must require an application id")
assert(stderr.include?("--app is required"), "compatibility run plan CLI must explain missing application ids")

puts "PASS: compatibility run plan unit tests"
