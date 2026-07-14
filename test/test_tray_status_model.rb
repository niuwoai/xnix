#!/usr/bin/env ruby
# frozen_string_literal: true

require "json"
require "open3"
require "pathname"
require_relative "../lib/xnix/compatibility/tray_status_model"

def assert(condition, message)
  return if condition

  warn "FAIL: #{message}"
  exit 1
end

project_root = Pathname.new(__dir__).join("..").realpath
model = Xnix::Compatibility::TrayStatusModel.new(
  active_count: 2,
  attention_count: 1,
  bridged_tray_count: 1
).to_h

assert(model["version"] == "0.2.126", "tray status model must expose the current version")
assert(model["request_type"] == "tray-status-model", "tray status model must identify the model type")
assert(model["desktop"] == "KDE Plasma", "tray status model must target KDE Plasma")
assert(model["runtime_activity"]["active_application_count"] == 2, "tray status model must expose active applications")
assert(model["runtime_activity"]["attention_required_count"] == 1, "tray status model must expose attention count")
assert(model["compatibility_status"]["state"] == "attention-required", "tray status model must flag attention state")
assert(model["tray_bridge"]["state"] == "active", "tray status model must expose tray bridge state")
assert(model["actions"].include?("open-compatibility-center"), "tray status model must offer Compatibility Center navigation")
assert(model["actions"].include?("open-settings"), "tray status model must offer Settings navigation")

json = JSON.pretty_generate(model)
assert(!json.match?(/wine|prefix|\.wine|proton|virtual machine/i), "tray status model must not expose backend implementation terms")

stdout, stderr, result = Open3.capture3(
  "ruby",
  project_root.join("bin/xnix-compat-tray-status").to_s,
  "--active",
  "1",
  "--attention",
  "0",
  "--bridged-tray",
  "0"
)
assert(result.success?, "tray status CLI must exit successfully: #{stderr}")
cli_model = JSON.parse(stdout)
assert(cli_model["compatibility_status"]["state"] == "ready", "tray status CLI must emit ready state")
assert(cli_model["runtime_activity"]["summary"] == "1 compatibility application active", "tray status CLI must summarize one active app")

_stdout, stderr, result = Open3.capture3(
  "ruby",
  project_root.join("bin/xnix-compat-tray-status").to_s,
  "--active",
  "-1"
)
assert(!result.success?, "tray status CLI must reject negative counts")
assert(stderr.include?("non-negative integer"), "tray status CLI must explain count validation")

puts "PASS: compatibility tray status model unit tests"
