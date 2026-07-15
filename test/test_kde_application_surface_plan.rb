# frozen_string_literal: true

require "json"
require "open3"
require "pathname"
require_relative "../lib/xnix/compatibility/kde_application_surface_plan"
require_relative "../lib/xnix/compatibility/recipe_store"

PROJECT_ROOT = Pathname.new(__dir__).join("..").realpath

def assert(condition, message)
  raise "FAIL: #{message}" unless condition
end

store = Xnix::Compatibility::RecipeStore.new(path: PROJECT_ROOT.join("runtime/recipes"))
recipe = store.find("org.xnix.sample.notepad")
plan = Xnix::Compatibility::KDEApplicationSurfacePlan.new(recipe: recipe).to_h

assert(plan["version"] == "0.2.181", "KDE application surface plan must expose the current version")
assert(plan["plan_type"] == "kde-application-surface-plan", "KDE application surface plan must identify the plan type")
assert(plan["runtime_method"] == "GetKDEApplicationSurfacePlan", "KDE application surface plan must identify the Runtime method")
assert(plan["application"]["id"] == "org.xnix.sample.notepad", "KDE application surface plan must preserve the application id")
assert(plan["desktop_shell"] == "KDE Plasma", "KDE application surface plan must target KDE Plasma")
assert(plan["entry_point_count"] == 7, "KDE application surface plan must cover the seven first-release entry points")
assert(plan["entry_points"].map { |entry| entry.fetch("id") } == %w[launcher task-manager file-manager system-tray notifications compatibility-center settings], "KDE application surface plan must preserve entry point order")
assert(plan["normal_linux_application_surface"], "KDE application surface plan must present compatibility apps as normal Linux apps")
assert(plan["standard_launcher_visible"], "KDE application surface plan must keep standard launcher visibility")
assert(plan["portal_review_required"], "KDE application surface plan must require Portal review")
assert(!plan["launch_enabled"], "KDE application surface plan must not enable launch")
assert(!plan["backend_process_started"], "KDE application surface plan must not start backend processes")
assert(!plan["desktop_files_written"], "KDE application surface plan must not write desktop files")
assert(!plan["mimeapps_written"], "KDE application surface plan must not write MIME defaults")
assert(!plan["host_root_modified"], "KDE application surface plan must not mutate the host root")
assert(!plan["backend_command_exposed"], "KDE application surface plan must not expose backend commands")
assert(!plan["raw_windows_executable_exposed"], "KDE application surface plan must not expose raw Windows executable paths")
assert(!plan["backend_details_exposed"], "KDE application surface plan must not expose backend details")

stdout, stderr, status = Open3.capture3(
  "ruby",
  PROJECT_ROOT.join("bin/xnix-kde-application-surface-plan").to_s,
  "--app",
  "org.xnix.sample.notepad",
  "--recipe-dir",
  PROJECT_ROOT.join("runtime/recipes").to_s
)
assert(status.success?, "KDE application surface plan CLI must exit successfully: #{stderr}")
cli_plan = JSON.parse(stdout)
assert(cli_plan["plan_type"] == "kde-application-surface-plan", "KDE application surface plan CLI must emit the plan")
assert(!cli_plan["backend_details_exposed"], "KDE application surface plan CLI must preserve backend detail gates")

puts "PASS: KDE application surface plan unit tests"
