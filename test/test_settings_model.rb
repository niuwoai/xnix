#!/usr/bin/env ruby
# frozen_string_literal: true

require "json"
require "open3"
require "pathname"
require_relative "../lib/xnix/compatibility/settings_model"

def assert(condition, message)
  return if condition

  warn "FAIL: #{message}"
  exit 1
end

project_root = Pathname.new(__dir__).join("..").realpath
model = Xnix::Compatibility::SettingsModel.new(application_id: "org.xnix.sample.notepad").to_h
sections = model.fetch("sections")
section_ids = sections.map { |section| section.fetch("id") }

assert(model["version"] == "0.2.208", "settings model must expose the current version")
assert(model["request_type"] == "settings-model", "settings model must identify the model type")
assert(model["desktop"] == "KDE Plasma", "settings model must target KDE Plasma")
assert(model["runtime_owned"], "Runtime must own compatibility settings")
assert(!model["kde_policy_owner"], "KDE must not own compatibility settings policy")
assert(model["settings_state"] == "planned", "settings model must expose planned settings state")
assert(!model["settings_persisted"], "settings model must not claim persisted settings yet")
assert(!model["host_root_modified"], "settings model must not mutate the host root")
assert(!model["backend_details_exposed"], "settings model must hide backend details")
assert(model["section_count"] == 5, "settings model must count user-facing sections")
assert(section_ids == %w[run-mode resource-access devices network snapshots], "settings model must expose the required sections")

run_mode = sections.find { |section| section["id"] == "run-mode" }
mode = run_mode.fetch("fields").find { |field| field["id"] == "mode" }
assert(mode["value"] == "automatic", "settings model must default to automatic mode")
assert(mode["options"].include?("performance"), "settings model must include performance priority")
assert(mode["options"].include?("compatibility"), "settings model must include compatibility priority")

documents = sections.find { |section| section["id"] == "resource-access" }.fetch("fields").find { |field| field["id"] == "documents" }
downloads = sections.find { |section| section["id"] == "resource-access" }.fetch("fields").find { |field| field["id"] == "downloads" }
assert(documents["value"] == "ask", "documents access must default to ask")
assert(downloads["value"] == "ask", "downloads access must default to ask")

camera = sections.find { |section| section["id"] == "devices" }.fetch("fields").find { |field| field["id"] == "camera" }
network = sections.find { |section| section["id"] == "network" }.fetch("fields").find { |field| field["id"] == "network" }
snapshots = sections.find { |section| section["id"] == "snapshots" }.fetch("fields").find { |field| field["id"] == "snapshots" }
assert(camera["value"] == "deny", "camera access must default to deny")
assert(network["value"] == "allow", "network access must default to allow")
assert(snapshots["value"] == "enabled", "snapshots must default to enabled")

json = JSON.pretty_generate(model)
assert(!json.match?(/wine|prefix|\.wine|proton|virtual machine/i), "settings model must not expose backend implementation terms")

stdout, stderr, result = Open3.capture3(
  "ruby",
  project_root.join("bin/xnix-compat-settings").to_s,
  "org.xnix.sample.notepad"
)
assert(result.success?, "settings CLI must exit successfully: #{stderr}")
assert(JSON.parse(stdout) == model, "settings CLI must emit the model")

_stdout, stderr, result = Open3.capture3("ruby", project_root.join("bin/xnix-compat-settings").to_s, "invalid")
assert(!result.success?, "settings CLI must reject invalid application ids")
assert(stderr.include?("reverse-DNS"), "settings CLI must explain application id validation")

puts "PASS: compatibility settings model unit tests"
