#!/usr/bin/env ruby
# frozen_string_literal: true

require "json"
require "open3"
require "pathname"
require_relative "../lib/xnix/compatibility/task_manager_identity"

def assert(condition, message)
  return if condition

  warn "FAIL: #{message}"
  exit 1
end

project_root = Pathname.new(__dir__).join("..").realpath
identity = Xnix::Compatibility::TaskManagerIdentity.new(
  application_id: "org.xnix.sample.notepad",
  name: "Sample Notepad"
).to_h

assert(identity["version"] == "0.2.48", "task manager identity must expose the current version")
assert(identity["request_type"] == "task-manager-identity", "task manager identity must identify the model type")
assert(identity["desktop"] == "KDE Plasma", "task manager identity must target KDE Plasma")
assert(identity["desktop_file"] == "xnix-org.xnix.sample.notepad.desktop", "task manager identity must use generated desktop files")
assert(identity["window"]["class_group"] == "xnix-compatibility", "task manager identity must provide a stable window class group")
assert(identity["task_manager"]["pinning_allowed"], "task manager identity must allow pinning")
assert(identity["task_manager"]["restore_allowed"], "task manager identity must allow restore")
assert(identity["task_manager"]["grouping_key"] == "org.xnix.sample.notepad", "task manager identity must group by Runtime application id")
assert(identity["kwin"]["script_role"] == "identity-only", "KWin model must stay identity-only")

json = JSON.pretty_generate(identity)
assert(!json.match?(/wine|prefix|\.wine|proton|virtual machine/i), "task manager identity must not expose backend implementation terms")

stdout, stderr, result = Open3.capture3(
  "ruby",
  project_root.join("bin/xnix-compat-window-identity").to_s,
  "--app",
  "org.xnix.sample.notepad",
  "--name",
  "Sample Notepad"
)
assert(result.success?, "window identity CLI must exit successfully: #{stderr}")
assert(JSON.parse(stdout) == identity, "window identity CLI must emit the model")

_stdout, stderr, result = Open3.capture3(
  "ruby",
  project_root.join("bin/xnix-compat-window-identity").to_s,
  "--app",
  "invalid",
  "--name",
  "Sample Notepad"
)
assert(!result.success?, "window identity CLI must reject invalid application ids")
assert(stderr.include?("reverse-DNS"), "window identity CLI must explain application id validation")

puts "PASS: compatibility task manager identity unit tests"
