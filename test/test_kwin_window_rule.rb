#!/usr/bin/env ruby
# frozen_string_literal: true

require "json"
require "open3"
require "pathname"
require_relative "../lib/xnix/compatibility/kwin_window_rule"

def assert(condition, message)
  return if condition

  warn "FAIL: #{message}"
  exit 1
end

project_root = Pathname.new(__dir__).join("..").realpath
rule = Xnix::Compatibility::KWinWindowRule.new(
  application_id: "org.xnix.sample.notepad",
  name: "Sample Notepad"
).to_h

assert(rule["version"] == "0.2.73", "KWin window rule must expose the current version")
assert(rule["request_type"] == "kwin-window-rule", "KWin window rule must identify the model type")
assert(rule["desktop"] == "KDE Plasma", "KWin window rule must target KDE Plasma")
assert(rule["script_role"] == "identity-and-layout", "KWin window rule must keep a bounded script role")
assert(rule["match"]["resource_name"] == "org.xnix.sample.notepad", "KWin window rule must match Runtime application identity")
assert(rule["match"]["class_group"] == "xnix-compatibility", "KWin window rule must match the compatibility class group")
assert(rule["set"]["desktop_file"] == "xnix-org.xnix.sample.notepad.desktop", "KWin window rule must bind generated desktop files")
assert(rule["set"]["task_manager_grouping_key"] == "org.xnix.sample.notepad", "KWin window rule must preserve task grouping")
assert(!rule["set"]["skip_taskbar"], "KWin window rule must keep compatibility windows visible in the taskbar")
assert(rule["set"]["show_in_switcher"], "KWin window rule must keep compatibility windows visible in the switcher")
assert(rule["restore"]["pinning_allowed"], "KWin window rule must preserve pinning")
assert(rule["restore"]["restore_allowed"], "KWin window rule must preserve restore")
assert(rule["restore"]["prefer_existing_window"], "KWin window rule must prefer restoring existing windows")
assert(rule["safety"]["window_manager_policy_only"], "KWin window rule must stay limited to window-manager policy")
assert(rule["safety"]["runtime_owns_backend_policy"], "KWin window rule must leave backend policy in the Runtime")
assert(!rule["safety"]["backend_details_exposed"], "KWin window rule must hide backend details")

json = JSON.pretty_generate(rule)
assert(!json.match?(/wine|prefix|\.wine|proton|virtual machine/i), "KWin window rule must not expose backend implementation terms")

stdout, stderr, result = Open3.capture3(
  "ruby",
  project_root.join("bin/xnix-kwin-window-rule").to_s,
  "--app",
  "org.xnix.sample.notepad",
  "--name",
  "Sample Notepad"
)
assert(result.success?, "KWin window rule CLI must exit successfully: #{stderr}")
assert(JSON.parse(stdout) == rule, "KWin window rule CLI must emit the model")

_stdout, stderr, result = Open3.capture3(
  "ruby",
  project_root.join("bin/xnix-kwin-window-rule").to_s,
  "--app",
  "invalid",
  "--name",
  "Sample Notepad"
)
assert(!result.success?, "KWin window rule CLI must reject invalid application ids")
assert(stderr.include?("reverse-DNS"), "KWin window rule CLI must explain application id validation")

puts "PASS: KDE KWin window rule unit tests"
