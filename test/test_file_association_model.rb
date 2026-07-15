#!/usr/bin/env ruby
# frozen_string_literal: true

require "json"
require "open3"
require "pathname"
require_relative "../lib/xnix/compatibility/file_association_model"
require_relative "../lib/xnix/compatibility/recipe_store"

def assert(condition, message)
  return if condition

  warn "FAIL: #{message}"
  exit 1
end

project_root = Pathname.new(__dir__).join("..").realpath
recipe = Xnix::Compatibility::RecipeStore.new(path: project_root.join("runtime/recipes")).find("org.xnix.sample.notepad")
model = Xnix::Compatibility::FileAssociationModel.new(recipe: recipe)
payload = model.to_h
contents = model.render_mimeapps

assert(payload["version"] == "0.2.163", "file association model must expose the current version")
assert(payload["association_type"] == "desktop-file-association", "file association model must identify its type")
assert(payload["desktop"] == "KDE Plasma", "file association model must target KDE Plasma")
assert(payload["application"]["desktop_file"] == "xnix-org.xnix.sample.notepad.desktop", "file association model must bind the generated desktop file")
assert(payload["mimeapps"]["path"] == "usr/share/applications/mimeapps.list", "file association model must target the standard mimeapps list")
assert(payload["associations"].length == 2, "file association model must include recipe MIME types")
assert(payload["associations"].all? { |entry| entry["default_application"] }, "file association model must mark default applications")
assert(payload["associations"].all? { |entry| entry["file_open"]["portal_required"] }, "file association model must require portal-mediated file open")
assert(payload["safety"]["staged_root_only"], "file association model must stay scoped to a staging root")
assert(!payload["safety"]["overwrite_existing_mimeapps"], "file association model must not permit blind mimeapps overwrite")
assert(!payload["safety"]["backend_details_exposed"], "file association model must hide backend details")

assert(contents.include?("[Default Applications]"), "mimeapps output must include default applications")
assert(contents.include?("application/x-xnix-txt=xnix-org.xnix.sample.notepad.desktop"), "mimeapps output must set txt default")
assert(contents.include?("[Added Associations]"), "mimeapps output must include added associations")
assert(contents.include?("application/x-xnix-log=xnix-org.xnix.sample.notepad.desktop;"), "mimeapps output must set log association")

json = JSON.pretty_generate(payload)
assert(!json.match?(/wine|prefix|\.wine|proton|virtual machine/i), "file association model must not expose backend implementation terms")

stdout, stderr, result = Open3.capture3(
  "ruby",
  project_root.join("bin/xnix-file-association-model").to_s,
  "--app",
  "org.xnix.sample.notepad"
)
assert(result.success?, "file association model CLI must exit successfully: #{stderr}")
assert(JSON.parse(stdout) == payload, "file association model CLI must emit the model")

_stdout, stderr, result = Open3.capture3(
  "ruby",
  project_root.join("bin/xnix-file-association-model").to_s,
  "--app",
  "org.xnix.missing"
)
assert(!result.success?, "file association model CLI must reject unknown applications")
assert(stderr.include?("unknown application"), "file association model CLI must explain unknown applications")

puts "PASS: compatibility file association model unit tests"
