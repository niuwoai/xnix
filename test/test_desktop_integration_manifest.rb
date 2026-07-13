#!/usr/bin/env ruby
# frozen_string_literal: true

require "json"
require "open3"
require "pathname"
require_relative "../lib/xnix/compatibility/desktop_integration_manifest"

def assert(condition, message)
  return if condition

  warn "FAIL: #{message}"
  exit 1
end

project_root = Pathname.new(__dir__).join("..").realpath
recipe_store = Xnix::Compatibility::RecipeStore.new(path: project_root.join("runtime/recipes"))
recipe = recipe_store.find("org.xnix.sample.notepad")
manifest = Xnix::Compatibility::DesktopIntegrationManifest.new(recipe: recipe).to_h
artifacts = manifest.fetch("artifacts")
entry_points = manifest.fetch("entry_points")

expected_entry_points = %w[
  launcher
  task-manager
  file-manager
  system-tray
  notifications
  compatibility-center
  settings
]

assert(manifest["version"] == "0.2.44", "desktop integration manifest must expose the current version")
assert(manifest["manifest_type"] == "desktop-integration", "manifest must identify its type")
assert(manifest["desktop"] == "KDE Plasma", "manifest must target the official KDE desktop")
assert(manifest["official_desktop_only"], "manifest must not expand first-release desktop scope")
assert(manifest["application"]["id"] == "org.xnix.sample.notepad", "manifest must identify the recipe application")
assert(entry_points == expected_entry_points, "manifest must include the seven first-release entry points")
assert(artifacts.map { |artifact| artifact.fetch("entry_point") } == expected_entry_points, "manifest artifacts must follow the entry point order")

launcher = artifacts.find { |artifact| artifact["entry_point"] == "launcher" }
assert(launcher["kind"] == "desktop-entry", "launcher artifact must be a desktop entry")
assert(launcher["path"] == "applications/xnix-org.xnix.sample.notepad.desktop", "launcher artifact must name the generated desktop file")
assert(launcher["argv"] == ["xnix-compat-launch", "--app", "org.xnix.sample.notepad", "%U"], "launcher artifact must delegate to the managed launcher")

task_manager = artifacts.find { |artifact| artifact["entry_point"] == "task-manager" }
assert(task_manager["argv"].include?("xnix-compat-window-identity"), "task manager artifact must use the window identity model")
assert(task_manager["kwin_rule"]["argv"].include?("xnix-kwin-window-rule"), "task manager artifact must expose the KWin window rule model")
assert(task_manager["kwin_rule"]["script_role"] == "identity-and-layout", "task manager artifact must keep KWin rule scope bounded")
assert(task_manager["desktop_file"] == "xnix-org.xnix.sample.notepad.desktop", "task manager artifact must share the generated desktop file")

file_manager = artifacts.find { |artifact| artifact["entry_point"] == "file-manager" }
assert(file_manager["path"] == "servicemenus/xnix-open-with-compatibility.desktop", "file manager artifact must name the Dolphin service menu")
assert(file_manager["argv"] == ["xnix-compat-open", "%U"], "file manager artifact must delegate to the managed file-open command")
assert(file_manager["file_association"]["argv"] == ["xnix-file-association-model", "--app", "org.xnix.sample.notepad"], "file manager artifact must expose file association generation")
assert(file_manager["file_association"]["path"] == "usr/share/applications/mimeapps.list", "file manager artifact must target the standard mimeapps list")
assert(file_manager["file_association"]["mime_types"].include?("application/x-xnix-txt"), "file manager artifact must include recipe MIME types")
assert(file_manager["portal_required"], "file manager artifact must require portal-mediated file access")

system_tray = artifacts.find { |artifact| artifact["entry_point"] == "system-tray" }
assert(system_tray["argv"] == ["xnix-compat-tray-status"], "system tray artifact must expose the tray status model")

notifications = artifacts.find { |artifact| artifact["entry_point"] == "notifications" }
assert(notifications["argv"] == ["xnix-compat-notify", "--app", "org.xnix.sample.notepad"], "notifications artifact must expose notification requests")
assert(notifications["event_types"].include?("approval-required"), "notifications artifact must include approval events")

compatibility_center = artifacts.find { |artifact| artifact["entry_point"] == "compatibility-center" }
assert(compatibility_center["argv"] == ["xnix-kde-center-model"], "Compatibility Center artifact must expose the KDE read model")

settings = artifacts.find { |artifact| artifact["entry_point"] == "settings" }
assert(settings["argv"] == ["xnix-compat-settings", "org.xnix.sample.notepad"], "settings artifact must expose user-facing settings")
assert(settings["portal_policy"]["argv"] == ["xnix-portal-access-policy", "--app", "org.xnix.sample.notepad", "--operation", "file-open"], "settings artifact must expose the portal policy command")
assert(settings["portal_request"]["argv"] == ["xnix-portal-request-model", "--app", "org.xnix.sample.notepad", "--operation", "file-open"], "settings artifact must expose the portal request model")
assert(settings["portal_request"]["destination"] == "org.freedesktop.portal.Desktop", "settings artifact must expose the portal D-Bus destination")
assert(settings["portal_policy"]["operations"].include?("camera"), "settings artifact must include sensitive device portal operations")
assert(settings["portal_policy"]["operations"].include?("remote-desktop"), "settings artifact must include remote desktop portal operations")
assert(settings["sections"].include?("snapshots"), "settings artifact must include snapshot policy")

safety = manifest.fetch("safety")
assert(!safety["backend_commands_exposed"], "manifest must not expose backend commands")
assert(safety["portal_required_for_file_access"], "manifest must require portal-mediated file access")
assert(!safety["host_privilege_required"], "manifest must not require host privilege")

json = JSON.pretty_generate(manifest)
assert(!json.match?(/wine|prefix|\.wine|proton|virtual machine/i), "manifest must not expose backend implementation terms")

stdout, stderr, result = Open3.capture3(
  "ruby",
  project_root.join("bin/xnix-desktop-integration-manifest").to_s,
  "--app",
  "org.xnix.sample.notepad"
)
assert(result.success?, "desktop integration manifest CLI must exit successfully: #{stderr}")
assert(JSON.parse(stdout) == manifest, "desktop integration manifest CLI must emit the manifest")

_stdout, stderr, result = Open3.capture3(
  "ruby",
  project_root.join("bin/xnix-desktop-integration-manifest").to_s,
  "--app",
  "org.xnix.missing"
)
assert(!result.success?, "desktop integration manifest CLI must reject unknown applications")
assert(stderr.include?("unknown application"), "desktop integration manifest CLI must explain unknown applications")

puts "PASS: desktop integration manifest unit tests"
