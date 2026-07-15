# frozen_string_literal: true

require "json"
require "open3"
require "pathname"
require_relative "../lib/xnix/compatibility/desktop_resource_bridge_plan"
require_relative "../lib/xnix/compatibility/recipe_store"

PROJECT_ROOT = Pathname.new(__dir__).join("..").realpath

def assert(condition, message)
  raise "FAIL: #{message}" unless condition
end

store = Xnix::Compatibility::RecipeStore.new(path: PROJECT_ROOT.join("runtime/recipes"))
recipe = store.find("org.xnix.sample.notepad")
plan = Xnix::Compatibility::DesktopResourceBridgePlan.new(recipe: recipe).to_h

assert(plan["version"] == "0.2.175", "desktop resource bridge plan must expose the current version")
assert(plan["plan_type"] == "desktop-resource-bridge-plan", "desktop resource bridge plan must identify the plan type")
assert(plan["runtime_method"] == "GetDesktopResourceBridgePlan", "desktop resource bridge plan must identify the Runtime method")
assert(plan["application"]["id"] == "org.xnix.sample.notepad", "desktop resource bridge plan must preserve the application id")
assert(plan["bridge_state"] == "planned", "desktop resource bridge plan must stay planned")
assert(plan["resource_count"] == 5, "desktop resource bridge plan must expose five resource bridges")
assert(plan["resources"].map { |resource| resource.fetch("id") } == %w[file-open uri-open print clipboard screenshot], "desktop resource bridge plan must preserve resource order")
assert(plan["portal_mediated"], "desktop resource bridge plan must require Portal mediation")
assert(plan["file_bridge_planned"], "desktop resource bridge plan must plan file bridges")
assert(plan["print_bridge_planned"], "desktop resource bridge plan must plan print bridges")
assert(plan["clipboard_bridge_planned"], "desktop resource bridge plan must plan clipboard bridges")
assert(!plan["bridges_enabled"], "desktop resource bridge plan must not enable bridges")
assert(!plan["requests_created"], "desktop resource bridge plan must not create Portal requests")
assert(!plan["backend_process_started"], "desktop resource bridge plan must not start backend processes")
assert(!plan["direct_host_file_access"], "desktop resource bridge plan must not grant direct host file access")
assert(!plan["direct_clipboard_access"], "desktop resource bridge plan must not grant direct clipboard access")
assert(!plan["direct_print_access"], "desktop resource bridge plan must not grant direct print access")
assert(!plan["host_root_modified"], "desktop resource bridge plan must not mutate the host root")
assert(!plan["backend_details_exposed"], "desktop resource bridge plan must not expose backend details")

stdout, stderr, status = Open3.capture3(
  "ruby",
  PROJECT_ROOT.join("bin/xnix-desktop-resource-bridge-plan").to_s,
  "--app",
  "org.xnix.sample.notepad",
  "--recipe-dir",
  PROJECT_ROOT.join("runtime/recipes").to_s
)
assert(status.success?, "desktop resource bridge plan CLI must exit successfully: #{stderr}")
cli_plan = JSON.parse(stdout)
assert(cli_plan["plan_type"] == "desktop-resource-bridge-plan", "desktop resource bridge plan CLI must emit the plan")
assert(cli_plan["portal_mediated"], "desktop resource bridge plan CLI must preserve Portal mediation")
assert(!cli_plan["backend_details_exposed"], "desktop resource bridge plan CLI must preserve backend detail gates")

puts "PASS: desktop resource bridge plan unit tests"
