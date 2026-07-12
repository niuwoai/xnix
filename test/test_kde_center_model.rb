#!/usr/bin/env ruby
# frozen_string_literal: true

require "json"
require "open3"
require "pathname"
require_relative "../lib/xnix/compatibility/kde_center_model"
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
model = Xnix::Compatibility::KdeCenterModel.new(runtime: runtime).to_h

assert(model["version"] == "0.2.8", "KDE center model must expose the current version")
assert(model["source"]["kind"] == "runtime-local-read-model", "KDE center model must describe the local fallback read model")
assert(model["source"]["bus_name"] == "org.xnix.Compatibility1", "KDE center model must keep the Runtime bus boundary visible")
assert(model["summary"]["application_count"] == 1, "KDE center model must summarize bundled applications")
assert(model["summary"]["known_application_count"] == 1, "KDE center model must summarize known applications")
assert(model["summary"]["pending_action_count"] == 1, "KDE center model must summarize pending compatibility work")

application = model.fetch("applications").first
assert(application["id"] == "org.xnix.sample.notepad", "KDE center model must include the sample application")
assert(application["mode_label"] == "Automatic", "KDE center model must present user-facing mode labels")
assert(application["compatibility_label"] == "Known", "KDE center model must present user-facing status labels")
assert(application["summary"] == "1 compatibility task pending", "KDE center model must present a concise task summary")
assert(application["supported_extensions"].include?(".txt"), "KDE center model must include supported file extensions")

json = JSON.pretty_generate(model)
assert(!json.match?(/prefix|\.wine|proton|virtual machine/i), "KDE center model must not expose backend storage or implementation terms")

stdout, stderr, status = Open3.capture3("ruby", project_root.join("bin/xnix-kde-center-model").to_s, "--source", "local")
assert(status.success?, "KDE center model CLI must exit successfully: #{stderr}")
cli_model = JSON.parse(stdout)
assert(cli_model == model, "KDE center model CLI must emit the same model")

puts "PASS: KDE compatibility center model unit tests"
