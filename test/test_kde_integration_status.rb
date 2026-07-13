#!/usr/bin/env ruby
# frozen_string_literal: true

require "json"
require "open3"
require "pathname"
require_relative "../lib/xnix/compatibility/kde_integration_status"

def assert(condition, message)
  return if condition

  warn "FAIL: #{message}"
  exit 1
end

project_root = Pathname.new(__dir__).join("..").realpath
status = Xnix::Compatibility::KdeIntegrationStatus.new.to_h
entry_points = status.fetch("entry_points")
ids = entry_points.map { |entry| entry.fetch("id") }

expected_ids = %w[
  launcher
  task-manager
  file-manager
  system-tray
  notifications
  compatibility-center
  settings
]

assert(status["version"] == "0.2.78", "KDE integration status must expose the current version")
assert(status["desktop"] == "KDE Plasma", "KDE integration status must keep KDE as the official desktop")
assert(status["official_desktop_only"], "KDE integration status must reject first-release multi-desktop scope")
assert(ids == expected_ids, "KDE integration status must track the seven first-release entry points")
assert(status["summary"]["total"] == 7, "KDE integration status must summarize all entry points")
assert(status["summary"]["initial"] == 7, "KDE integration status must count initial entry points")
assert(status["summary"]["planned"] == 0, "KDE integration status must count planned entry points")

launcher = entry_points.find { |entry| entry["id"] == "launcher" }
assert(launcher["evidence"].any? { |item| item.include?("xnix-compat-launch") }, "launcher status must reference the managed launcher")

file_manager = entry_points.find { |entry| entry["id"] == "file-manager" }
assert(file_manager["evidence"].any? { |item| item.include?("xnix-compat-open") }, "file manager status must reference the managed file-open entry point")

compatibility_center = entry_points.find { |entry| entry["id"] == "compatibility-center" }
assert(compatibility_center["evidence"].any? { |item| item.include?("session bus") }, "Compatibility Center status must reference Runtime D-Bus consumption")

notifications = entry_points.find { |entry| entry["id"] == "notifications" }
assert(notifications["state"] == "initial", "Notifications must have an initial integration")
assert(notifications["evidence"].any? { |item| item.include?("xnix-compat-notify") }, "Notifications status must reference the Runtime notification entry point")

settings = entry_points.find { |entry| entry["id"] == "settings" }
assert(settings["state"] == "initial", "Settings must have an initial integration")
assert(settings["evidence"].any? { |item| item.include?("xnix-compat-settings") }, "Settings status must reference the settings entry point")

system_tray = entry_points.find { |entry| entry["id"] == "system-tray" }
assert(system_tray["state"] == "initial", "System Tray must have an initial integration")
assert(system_tray["evidence"].any? { |item| item.include?("xnix-compat-tray-status") }, "System Tray status must reference the tray status entry point")

task_manager = entry_points.find { |entry| entry["id"] == "task-manager" }
assert(task_manager["state"] == "initial", "Task Manager must have an initial integration")
assert(task_manager["evidence"].any? { |item| item.include?("xnix-compat-window-identity") }, "Task Manager status must reference the window identity entry point")

stdout, stderr, result = Open3.capture3("ruby", project_root.join("bin/xnix-kde-integration-status").to_s)
assert(result.success?, "KDE integration status CLI must exit successfully: #{stderr}")
assert(JSON.parse(stdout) == status, "KDE integration status CLI must emit the model")

puts "PASS: KDE integration status unit tests"
