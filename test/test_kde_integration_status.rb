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

assert(status["version"] == "0.2.10", "KDE integration status must expose the current version")
assert(status["desktop"] == "KDE Plasma", "KDE integration status must keep KDE as the official desktop")
assert(status["official_desktop_only"], "KDE integration status must reject first-release multi-desktop scope")
assert(ids == expected_ids, "KDE integration status must track the seven first-release entry points")
assert(status["summary"]["total"] == 7, "KDE integration status must summarize all entry points")
assert(status["summary"]["initial"] == 3, "KDE integration status must count initial entry points")
assert(status["summary"]["planned"] == 4, "KDE integration status must count planned entry points")

launcher = entry_points.find { |entry| entry["id"] == "launcher" }
assert(launcher["evidence"].any? { |item| item.include?("xnix-compat-launch") }, "launcher status must reference the managed launcher")

file_manager = entry_points.find { |entry| entry["id"] == "file-manager" }
assert(file_manager["evidence"].any? { |item| item.include?("xnix-compat-open") }, "file manager status must reference the managed file-open entry point")

compatibility_center = entry_points.find { |entry| entry["id"] == "compatibility-center" }
assert(compatibility_center["evidence"].any? { |item| item.include?("session bus") }, "Compatibility Center status must reference Runtime D-Bus consumption")

stdout, stderr, result = Open3.capture3("ruby", project_root.join("bin/xnix-kde-integration-status").to_s)
assert(result.success?, "KDE integration status CLI must exit successfully: #{stderr}")
assert(JSON.parse(stdout) == status, "KDE integration status CLI must emit the model")

puts "PASS: KDE integration status unit tests"
