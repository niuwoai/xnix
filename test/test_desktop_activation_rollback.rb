#!/usr/bin/env ruby
# frozen_string_literal: true

require "json"
require "open3"
require "pathname"
require "tmpdir"
require_relative "../lib/xnix/compatibility/desktop_activation_installer"
require_relative "../lib/xnix/compatibility/desktop_activation_rollback"

def assert(condition, message)
  return if condition

  warn "FAIL: #{message}"
  exit 1
end

project_root = Pathname.new(__dir__).join("..").realpath
recipe_store = Xnix::Compatibility::RecipeStore.new(path: project_root.join("runtime/recipes"))
recipe = recipe_store.find("org.xnix.sample.notepad")

Dir.mktmpdir("xnix-desktop-root") do |root|
  install_result = Xnix::Compatibility::DesktopActivationInstaller.new(root: root, recipe: recipe).install
  receipt_path = Pathname.new(root).join(install_result.fetch("receipt").fetch("path"))
  desktop_entry = Pathname.new(root).join("usr/share/applications/xnix-org.xnix.sample.notepad.desktop")
  mimeapps_list = Pathname.new(root).join("usr/share/applications/mimeapps.list")
  service_menu = Pathname.new(root).join("usr/share/kio/servicemenus/xnix-open-with-compatibility.desktop")
  manifest_path = Pathname.new(root).join("usr/share/xnix/compatibility/manifests/org.xnix.sample.notepad.json")

  assert(receipt_path.file?, "desktop activation installer must write a rollback receipt")
  receipt = JSON.parse(receipt_path.read)
  assert(receipt["rollback"]["command"] == "xnix-rollback-desktop-integration", "receipt must name the rollback command")
  assert(receipt["rollback"]["requires_matching_sha256"], "receipt must require checksum verification")
  assert(receipt["installed"].all? { |entry| entry["sha256"].match?(/\A[0-9a-f]{64}\z/) }, "receipt entries must include SHA-256 digests")

  result = Xnix::Compatibility::DesktopActivationRollback.new(
    root: root,
    application_id: "org.xnix.sample.notepad"
  ).rollback

  assert(result["version"] == "0.2.70", "desktop activation rollback must expose the current version")
  assert(result["application_id"] == "org.xnix.sample.notepad", "desktop activation rollback must identify the application")
  assert(result["removed"].length == 4, "desktop activation rollback must remove installed files")
  assert(result["removed"].all? { |entry| entry["status"] == "removed" }, "desktop activation rollback must report removed files")
  assert(result["receipt_removed"]["status"] == "removed", "desktop activation rollback must remove the receipt")
  assert(result["safety"]["sha256_verified_before_remove"], "desktop activation rollback must verify checksums before removal")
  assert(!desktop_entry.exist?, "desktop activation rollback must remove the desktop entry")
  assert(!mimeapps_list.exist?, "desktop activation rollback must remove the MIME association list")
  assert(!service_menu.exist?, "desktop activation rollback must remove the Dolphin service menu")
  assert(!manifest_path.exist?, "desktop activation rollback must remove the manifest")
  assert(!receipt_path.exist?, "desktop activation rollback must remove the receipt")
end

Dir.mktmpdir("xnix-desktop-root") do |root|
  Xnix::Compatibility::DesktopActivationInstaller.new(root: root, recipe: recipe).install
  desktop_entry = Pathname.new(root).join("usr/share/applications/xnix-org.xnix.sample.notepad.desktop")
  File.write(desktop_entry, "#{desktop_entry.read}# changed after install\n")

  _stdout, stderr, status = Open3.capture3(
    "ruby",
    project_root.join("bin/xnix-rollback-desktop-integration").to_s,
    "--root",
    root,
    "--app",
    "org.xnix.sample.notepad"
  )

  assert(!status.success?, "desktop activation rollback CLI must reject changed files")
  assert(stderr.include?("refusing to remove changed file"), "desktop activation rollback CLI must explain checksum failures")
  assert(desktop_entry.file?, "desktop activation rollback must preserve changed files")
end

Dir.mktmpdir("xnix-desktop-root") do |root|
  Xnix::Compatibility::DesktopActivationInstaller.new(root: root, recipe: recipe).install
  stdout, stderr, status = Open3.capture3(
    "ruby",
    project_root.join("bin/xnix-rollback-desktop-integration").to_s,
    "--root",
    root,
    "--app",
    "org.xnix.sample.notepad"
  )

  assert(status.success?, "desktop activation rollback CLI must exit successfully: #{stderr}")
  result = JSON.parse(stdout)
  assert(result["removed"].length == 4, "desktop activation rollback CLI must emit removed files")
end

_stdout, stderr, status = Open3.capture3(
  "ruby",
  project_root.join("bin/xnix-rollback-desktop-integration").to_s,
  "--root",
  "/",
  "--app",
  "org.xnix.sample.notepad"
)
assert(!status.success?, "desktop activation rollback CLI must reject host root rollback")
assert(stderr.include?("root must not be /"), "desktop activation rollback CLI must explain host root rejection")

puts "PASS: desktop activation rollback unit tests"
