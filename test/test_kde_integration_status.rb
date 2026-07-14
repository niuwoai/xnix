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

class KDEStatusRuntime
  def source_metadata
    {
      "kind" => "runtime-kde-status-test",
      "bus_name" => Xnix::Compatibility::RuntimeDaemon::BUS_NAME,
      "object_path" => Xnix::Compatibility::RuntimeDaemon::OBJECT_PATH,
      "interface" => Xnix::Compatibility::RuntimeDaemon::INTERFACE
    }
  end

  def kde_integration_status
    {
      "status_type" => "kde-integration-status",
      "desktop" => "KDE Plasma",
      "runtime_owned" => true,
      "kde_policy_owner" => false,
      "official_desktop_only" => true,
      "stable_desktop_contract" => true,
      "entry_points" => [
        {
          "id" => "launcher",
          "name" => "Launcher",
          "state" => "initial",
          "runtime_method" => "GetDesktopEntryPlan",
          "adapter_role" => "standard-desktop-entry",
          "runtime_backed" => true,
          "c_runtime_backed" => true,
          "dbus_read_available" => true,
          "kde_policy_owner" => false
        }
      ],
      "entry_point_count" => 1,
      "initial_count" => 1,
      "planned_count" => 0,
      "complete_count" => 0,
      "host_root_modified" => false,
      "backend_details_exposed" => false
    }
  end
end

class FlatKDEStatusRuntime
  def kde_integration_status
    {
      "status_type" => "kde-integration-status",
      "desktop" => "KDE Plasma",
      "entry_point_ids" => %w[launcher task-manager file-manager system-tray notifications compatibility-center settings],
      "entry_point_names" => ["Launcher", "Task Manager", "File Manager", "System Tray", "Notifications", "Compatibility Center", "Settings"],
      "entry_point_states" => %w[initial initial initial initial initial initial initial],
      "runtime_methods" => %w[
        GetDesktopEntryPlan
        GetTaskManagerIdentityPlan
        GetFileAssociationPlan
        GetTrayStatus
        GetNotificationPlan
        GetCompatibilityCenterSummary
        GetCompatibilitySettings
      ],
      "entry_point_count" => 7,
      "initial_count" => 7,
      "planned_count" => 0,
      "complete_count" => 0,
      "runtime_owned" => true,
      "kde_policy_owner" => false,
      "official_desktop_only" => true,
      "stable_desktop_contract" => true,
      "host_root_modified" => false,
      "backend_details_exposed" => false
    }
  end
end

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

assert(status["version"] == "0.2.111", "KDE integration status must expose the current version")
assert(status["status_type"] == "kde-integration-status", "KDE integration status must identify the status type")
assert(status["desktop"] == "KDE Plasma", "KDE integration status must keep KDE as the official desktop")
assert(status["source"]["kind"] == "runtime-local-read-model", "KDE integration status must describe the local Runtime read model")
assert(status["runtime_owned"], "KDE integration status must be Runtime-owned")
assert(!status["kde_policy_owner"], "KDE integration status must not make KDE the policy owner")
assert(status["official_desktop_only"], "KDE integration status must reject first-release multi-desktop scope")
assert(status["stable_desktop_contract"], "KDE integration status must expose a stable desktop contract")
assert(ids == expected_ids, "KDE integration status must track the seven first-release entry points")
assert(status["summary"]["total"] == 7, "KDE integration status must summarize all entry points")
assert(status["summary"]["initial"] == 7, "KDE integration status must count initial entry points")
assert(status["summary"]["planned"] == 0, "KDE integration status must count planned entry points")
assert(!status["safety"]["host_root_modified"], "KDE integration status must not mutate the host root")
assert(!status["safety"]["backend_details_exposed"], "KDE integration status must hide backend details")

launcher = entry_points.find { |entry| entry["id"] == "launcher" }
assert(launcher["runtime_method"] == "GetDesktopEntryPlan", "launcher status must reference Runtime desktop entry plans")
assert(launcher["c_runtime_backed"], "launcher status must be C Runtime-backed")

file_manager = entry_points.find { |entry| entry["id"] == "file-manager" }
assert(file_manager["runtime_method"] == "GetFileAssociationPlan", "file manager status must reference Runtime file association plans")

compatibility_center = entry_points.find { |entry| entry["id"] == "compatibility-center" }
assert(compatibility_center["runtime_method"] == "GetCompatibilityCenterSummary", "Compatibility Center status must reference Runtime summaries")

notifications = entry_points.find { |entry| entry["id"] == "notifications" }
assert(notifications["state"] == "initial", "Notifications must have an initial integration")
assert(notifications["runtime_method"] == "GetNotificationPlan", "Notifications status must reference Runtime notification plans")

settings = entry_points.find { |entry| entry["id"] == "settings" }
assert(settings["state"] == "initial", "Settings must have an initial integration")
assert(settings["runtime_method"] == "GetCompatibilitySettings", "Settings status must reference Runtime settings")

system_tray = entry_points.find { |entry| entry["id"] == "system-tray" }
assert(system_tray["state"] == "initial", "System Tray must have an initial integration")
assert(system_tray["runtime_method"] == "GetTrayStatus", "System Tray status must reference Runtime tray status")

task_manager = entry_points.find { |entry| entry["id"] == "task-manager" }
assert(task_manager["state"] == "initial", "Task Manager must have an initial integration")
assert(task_manager["runtime_method"] == "GetTaskManagerIdentityPlan", "Task Manager status must reference Runtime window identity plans")

runtime_status = Xnix::Compatibility::KdeIntegrationStatus.new(runtime: KDEStatusRuntime.new).to_h
assert(runtime_status["source"]["kind"] == "runtime-kde-status-test", "KDE integration status must preserve Runtime source metadata")
assert(runtime_status["entry_points"].length == 1, "KDE integration status must consume nested Runtime entry points")
assert(runtime_status["entry_points"].first["id"] == "launcher", "KDE integration status must preserve Runtime entry ids")

flat_status = Xnix::Compatibility::KdeIntegrationStatus.new(runtime: FlatKDEStatusRuntime.new).to_h
assert(flat_status["entry_points"].length == 7, "KDE integration status must normalize flat D-Bus entry points")
assert(flat_status["entry_points"].last["runtime_method"] == "GetCompatibilitySettings", "KDE integration status must normalize flat Runtime methods")
assert(flat_status["entry_points"].all? { |entry| entry.fetch("dbus_read_available") }, "KDE integration status must preserve D-Bus read availability")

stdout, stderr, result = Open3.capture3("ruby", project_root.join("bin/xnix-kde-integration-status").to_s, "--source", "local")
assert(result.success?, "KDE integration status CLI must exit successfully: #{stderr}")
assert(JSON.parse(stdout) == status, "KDE integration status CLI must emit the model")

puts "PASS: KDE integration status unit tests"
