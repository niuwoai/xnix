#!/usr/bin/env ruby
# frozen_string_literal: true

require "json"
require "open3"
require "pathname"
require_relative "../lib/xnix/compatibility/kwin_window_rule"
require_relative "../lib/xnix/compatibility/runtime_daemon"

def assert(condition, message)
  return if condition

  warn "FAIL: #{message}"
  exit 1
end

project_root = Pathname.new(__dir__).join("..").realpath

class KWinPlanRuntime
  attr_reader :requested_application_id

  def source_metadata
    {
      "kind" => "runtime-kwin-plan-test",
      "bus_name" => Xnix::Compatibility::RuntimeDaemon::BUS_NAME,
      "object_path" => Xnix::Compatibility::RuntimeDaemon::OBJECT_PATH,
      "interface" => Xnix::Compatibility::RuntimeDaemon::INTERFACE
    }
  end

  def kwin_window_rule_plan(application_id)
    @requested_application_id = application_id
    {
      "request_type" => "kwin-window-rule",
      "desktop" => "KDE Plasma",
      "runtime_owned" => true,
      "kde_policy_owner" => false,
      "application_id" => application_id,
      "name" => "Sample Notepad",
      "script_role" => "identity-and-layout",
      "match" => {
        "resource_name" => application_id,
        "class_group" => "xnix-compatibility",
        "title_hint" => "Sample Notepad"
      },
      "set" => {
        "desktop_file" => "xnix-org.xnix.sample.notepad.desktop",
        "application_id" => application_id,
        "task_manager_grouping_key" => application_id,
        "launcher_url" => "applications:xnix-org.xnix.sample.notepad.desktop",
        "skip_taskbar" => false,
        "show_in_switcher" => true,
        "placement" => "normal-window"
      },
      "restore" => {
        "pinning_allowed" => true,
        "restore_allowed" => true,
        "restore_key" => application_id,
        "prefer_existing_window" => true
      },
      "safety" => {
        "window_manager_policy_only" => true,
        "runtime_owns_backend_policy" => true,
        "host_root_modified" => false,
        "backend_details_exposed" => false
      },
      "host_root_modified" => false,
      "backend_details_exposed" => false,
      "desktop_safe_summary" => "KWin window rules are Runtime-owned identity and layout hints for normal desktop behavior."
    }
  end
end

class FlatKWinPlanRuntime
  def kwin_window_rule_plan(application_id)
    {
      "request_type" => "kwin-window-rule",
      "desktop" => "KDE Plasma",
      "runtime_owned" => true,
      "kde_policy_owner" => false,
      "application_id" => application_id,
      "name" => "Sample Notepad",
      "script_role" => "identity-and-layout",
      "resource_name" => application_id,
      "class_group" => "xnix-compatibility",
      "title_hint" => "Sample Notepad",
      "desktop_file" => "xnix-org.xnix.sample.notepad.desktop",
      "task_manager_grouping_key" => application_id,
      "launcher_url" => "applications:xnix-org.xnix.sample.notepad.desktop",
      "skip_taskbar" => false,
      "show_in_switcher" => true,
      "placement" => "normal-window",
      "pinning_allowed" => true,
      "restore_allowed" => true,
      "restore_key" => application_id,
      "prefer_existing_window" => true,
      "window_manager_policy_only" => true,
      "runtime_owns_backend_policy" => true,
      "host_root_modified" => false,
      "backend_details_exposed" => false
    }
  end
end

rule = Xnix::Compatibility::KWinWindowRule.new(
  application_id: "org.xnix.sample.notepad",
  name: "Sample Notepad"
).to_h

assert(rule["version"] == "0.2.115", "KWin window rule must expose the current version")
assert(rule["request_type"] == "kwin-window-rule", "KWin window rule must identify the model type")
assert(rule["desktop"] == "KDE Plasma", "KWin window rule must target KDE Plasma")
assert(rule["source"]["kind"] == "runtime-local-read-model", "KWin window rule must describe local Runtime reads")
assert(rule["runtime_owned"], "KWin window rule must preserve Runtime ownership")
assert(!rule["kde_policy_owner"], "KWin window rule must not claim KDE policy ownership")
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
assert(!rule["host_root_modified"], "KWin window rule must not mutate the host root")
assert(!rule["backend_details_exposed"], "KWin window rule must hide top-level backend details")

plan_runtime = KWinPlanRuntime.new
plan_rule = Xnix::Compatibility::KWinWindowRule.new(
  application_id: "org.xnix.sample.notepad",
  runtime: plan_runtime
).to_h
assert(plan_runtime.requested_application_id == "org.xnix.sample.notepad", "KWin window rule must query the Runtime plan")
assert(plan_rule["source"]["kind"] == "runtime-kwin-plan-test", "KWin window rule must preserve Runtime source metadata")
assert(plan_rule["match"]["resource_name"] == "org.xnix.sample.notepad", "KWin window rule must use Runtime plan matches")

flat_rule = Xnix::Compatibility::KWinWindowRule.new(
  application_id: "org.xnix.sample.notepad",
  runtime: FlatKWinPlanRuntime.new
).to_h
assert(flat_rule["match"]["class_group"] == "xnix-compatibility", "KWin window rule must normalize flat D-Bus plan matches")
assert(flat_rule["set"]["desktop_file"] == "xnix-org.xnix.sample.notepad.desktop", "KWin window rule must normalize flat D-Bus plan desktop files")
assert(flat_rule["restore"]["restore_key"] == "org.xnix.sample.notepad", "KWin window rule must normalize flat D-Bus restore keys")

json = JSON.pretty_generate(rule)
assert(!json.match?(/wine|prefix|\.wine|proton|virtual machine/i), "KWin window rule must not expose backend implementation terms")

stdout, stderr, result = Open3.capture3(
  "ruby",
  project_root.join("bin/xnix-kwin-window-rule").to_s,
  "--source",
  "local",
  "--app",
  "org.xnix.sample.notepad"
)
assert(result.success?, "KWin window rule CLI must exit successfully: #{stderr}")
assert(JSON.parse(stdout) == rule, "KWin window rule CLI must emit the model")

_stdout, stderr, result = Open3.capture3(
  "ruby",
  project_root.join("bin/xnix-kwin-window-rule").to_s,
  "--source",
  "local",
  "--app",
  "invalid"
)
assert(!result.success?, "KWin window rule CLI must reject invalid application ids")
assert(stderr.include?("reverse-DNS"), "KWin window rule CLI must explain application id validation")

puts "PASS: KDE KWin window rule unit tests"
