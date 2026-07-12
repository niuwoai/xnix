#!/usr/bin/env ruby
# frozen_string_literal: true

require "json"
require "open3"
require "pathname"
require "tmpdir"
require_relative "../lib/xnix/compatibility/desktop_activation_installer"

def assert(condition, message)
  return if condition

  warn "FAIL: #{message}"
  exit 1
end

project_root = Pathname.new(__dir__).join("..").realpath
recipe_store = Xnix::Compatibility::RecipeStore.new(path: project_root.join("runtime/recipes"))
recipe = recipe_store.find("org.xnix.sample.notepad")

Dir.mktmpdir("xnix-desktop-root") do |root|
  result = Xnix::Compatibility::DesktopActivationInstaller.new(root: root, recipe: recipe).install
  root_path = Pathname.new(root)
  desktop_entry = root_path.join("usr/share/applications/xnix-org.xnix.sample.notepad.desktop")
  service_menu = root_path.join("usr/share/kio/servicemenus/xnix-open-with-compatibility.desktop")
  manifest_path = root_path.join("usr/share/xnix/compatibility/manifests/org.xnix.sample.notepad.json")

  assert(result["version"] == "0.2.16", "desktop activation installer must expose the current version")
  assert(result["application_id"] == "org.xnix.sample.notepad", "desktop activation installer must identify the application")
  assert(result["root"] == root, "desktop activation installer must report the staging root")
  assert(result["installed"].length == 3, "desktop activation installer must install three staged files")
  assert(result["activated_entry_points"].length == 7, "desktop activation installer must activate all manifest entry points")
  assert(result["safety"]["staging_root_required"], "desktop activation installer must require a staging root")
  assert(!result["safety"]["host_root_modified"], "desktop activation installer must not modify the host root")
  assert(!result["safety"]["backend_commands_exposed"], "desktop activation installer must not expose backend commands")

  assert(desktop_entry.file?, "desktop activation installer must install the generated desktop entry")
  assert(desktop_entry.read.include?("Exec=xnix-compat-launch --app org.xnix.sample.notepad %U"), "installed desktop entry must call the managed launcher")
  assert((desktop_entry.stat.mode & 0o777) == 0o644, "installed desktop entry must be non-executable")

  assert(service_menu.file?, "desktop activation installer must install the Dolphin service menu")
  assert(service_menu.read.include?("Exec=xnix-compat-open %U"), "installed Dolphin service menu must call the managed file-open command")
  assert((service_menu.stat.mode & 0o777) == 0o644, "installed service menu must be non-executable")

  assert(manifest_path.file?, "desktop activation installer must persist the desktop integration manifest")
  manifest = JSON.parse(manifest_path.read)
  assert(manifest["manifest_type"] == "desktop-integration", "installed manifest must keep its manifest type")
  assert(manifest["entry_points"] == result["activated_entry_points"], "installed manifest must match activated entry points")
  assert(!manifest_path.read.match?(/wine|prefix|\.wine|proton|virtual machine/i), "installed manifest must not expose backend implementation terms")
end

Dir.mktmpdir("xnix-desktop-root") do |root|
  stdout, stderr, status = Open3.capture3(
    "ruby",
    project_root.join("bin/xnix-install-desktop-integration").to_s,
    "--root",
    root,
    "--app",
    "org.xnix.sample.notepad"
  )
  assert(status.success?, "desktop activation installer CLI must exit successfully: #{stderr}")
  result = JSON.parse(stdout)
  assert(result["application_id"] == "org.xnix.sample.notepad", "desktop activation installer CLI must emit install result")
  assert(Pathname.new(root).join("usr/share/applications/xnix-org.xnix.sample.notepad.desktop").file?, "desktop activation installer CLI must install the desktop entry")
end

_stdout, stderr, status = Open3.capture3(
  "ruby",
  project_root.join("bin/xnix-install-desktop-integration").to_s,
  "--root",
  "/",
  "--app",
  "org.xnix.sample.notepad"
)
assert(!status.success?, "desktop activation installer CLI must reject host root installation")
assert(stderr.include?("root must not be /"), "desktop activation installer CLI must explain host root rejection")

puts "PASS: desktop activation installer unit tests"
