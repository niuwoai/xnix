#!/usr/bin/env ruby
# frozen_string_literal: true

require "json"
require "open3"
require "pathname"
require_relative "../lib/xnix/compatibility/krunner_model"
require_relative "../lib/xnix/compatibility/recipe_store"
require_relative "../lib/xnix/compatibility/runtime_daemon"

def assert(condition, message)
  return if condition

  warn "FAIL: #{message}"
  exit 1
end

project_root = Pathname.new(__dir__).join("..").realpath
runtime = Xnix::Compatibility::RuntimeDaemon.new(
  recipe_store: Xnix::Compatibility::RecipeStore.new(path: project_root.join("runtime/recipes"))
)
model = Xnix::Compatibility::KRunnerModel.new(runtime: runtime, query: "notepad").to_h

assert(model["version"] == "0.2.87", "KRunner model must expose the current version")
assert(model["entry_point"] == "krunner", "KRunner model must identify the KDE entry point")
assert(model["desktop"] == "KDE Plasma", "KRunner model must stay scoped to KDE Plasma")
assert(model["source"]["kind"] == "runtime-local-read-model", "KRunner model must describe local fallback reads")
assert(model["summary"]["runtime_owned_launch"], "KRunner model must keep launch ownership in the Runtime")
assert(!model["summary"]["backend_details_exposed"], "KRunner model must hide backend details")

matches = model.fetch("matches")
assert(matches.length == 1, "KRunner model must return the sample application")
match = matches.first
assert(match["application_id"] == "org.xnix.sample.notepad", "KRunner model must resolve the Runtime application id")
assert(match["name"] == "Sample Notepad", "KRunner model must preserve the user-facing application name")
assert(match["subtitle"] == "Open as a normal Linux application", "KRunner model must present normal desktop wording")
assert(match["mode_label"] == "Automatic", "KRunner model must use user-facing mode labels")
assert(match["supported_extensions"].include?(".txt"), "KRunner model must include file extension hints")
assert(match["action"]["type"] == "runtime-launch", "KRunner model must emit a Runtime launch action")
assert(match["action"]["argv"] == ["xnix-compat-launch", "--app", "org.xnix.sample.notepad"], "KRunner model must delegate launch to the managed launcher")
assert(match["action"]["desktop_entry_id"] == "org.xnix.sample.notepad.desktop", "KRunner model must map matches to desktop entries")

extension_model = Xnix::Compatibility::KRunnerModel.new(runtime: runtime, query: "open txt").to_h
assert(extension_model["matches"].first["application_id"] == "org.xnix.sample.notepad", "KRunner model must resolve file-oriented natural queries")

blank_model = Xnix::Compatibility::KRunnerModel.new(runtime: runtime, query: " ").to_h
assert(blank_model["matches"].empty?, "KRunner model must not flood KRunner on blank queries")

json = JSON.pretty_generate(model)
assert(!json.match?(/prefix|\.wine|proton|virtual machine|wine\b|vm\b/i), "KRunner model must not expose backend storage or implementation terms")

stdout, stderr, status = Open3.capture3(
  "ruby",
  project_root.join("bin/xnix-krunner-model").to_s,
  "--source",
  "local",
  "--query",
  "notepad"
)
assert(status.success?, "KRunner model CLI must exit successfully: #{stderr}")
assert(JSON.parse(stdout) == model, "KRunner model CLI must emit the same model")

puts "PASS: KDE KRunner model unit tests"
