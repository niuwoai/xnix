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

class QueryPlanRuntime
  attr_reader :queried

  def source_metadata
    {
      "kind" => "runtime-query-plan-test",
      "bus_name" => Xnix::Compatibility::RuntimeDaemon::BUS_NAME,
      "object_path" => Xnix::Compatibility::RuntimeDaemon::OBJECT_PATH,
      "interface" => Xnix::Compatibility::RuntimeDaemon::INTERFACE
    }
  end

  def krunner_query_plan(query)
    @queried = query
    {
      "version" => Xnix::Compatibility::RuntimeDaemon::VERSION,
      "query_type" => "krunner-query-plan",
      "entry_point" => "krunner",
      "desktop" => "KDE Plasma",
      "query" => query,
      "runtime_owned" => true,
      "kde_policy_owner" => false,
      "matches" => [
        {
          "runner_id" => "xnix.compatibility.org.xnix.sample.notepad",
          "application_id" => "org.xnix.sample.notepad",
          "name" => "Sample Notepad",
          "icon" => "accessories-text-editor",
          "relevance_percent" => 100,
          "subtitle" => "Open as a normal Linux application",
          "mode_label" => "Automatic",
          "supported_extensions" => [".txt"],
          "runtime_owned_launch" => true,
          "backend_details_exposed" => false,
          "action" => {
            "type" => "runtime-launch",
            "desktop_entry_id" => "org.xnix.sample.notepad.desktop",
            "argv" => ["xnix-compat-launch", "--app", "org.xnix.sample.notepad"]
          }
        }
      ],
      "summary" => {
        "match_count" => 1,
        "official_desktop" => "KDE Plasma",
        "runtime_owned_launch" => true,
        "query_execution_enabled" => false,
        "backend_launch_enabled" => false,
        "backend_details_exposed" => false
      },
      "host_root_modified" => false,
      "backend_details_exposed" => false,
      "desktop_safe_summary" => "KRunner query planning is Runtime-owned and returns safe launcher actions only."
    }
  end

  def list_applications
    raise "KRunner model must use Runtime query plans before application-list fallback"
  end
end

class FlatQueryPlanRuntime
  def krunner_query_plan(query)
    {
      "query_type" => "krunner-query-plan",
      "entry_point" => "krunner",
      "desktop" => "KDE Plasma",
      "query" => query,
      "match_count" => 1,
      "top_application_id" => "org.xnix.sample.notepad",
      "top_name" => "Sample Notepad",
      "top_relevance_percent" => 100,
      "action_type" => "runtime-launch",
      "desktop_entry_id" => "org.xnix.sample.notepad.desktop",
      "runtime_owned" => true,
      "kde_policy_owner" => false,
      "runtime_owned_launch" => true,
      "query_execution_enabled" => false,
      "backend_launch_enabled" => false,
      "host_root_modified" => false,
      "backend_details_exposed" => false
    }
  end

  def list_applications
    raise "KRunner model must normalize flat Runtime query plans before application-list fallback"
  end
end

runtime = Xnix::Compatibility::RuntimeDaemon.new(
  recipe_store: Xnix::Compatibility::RecipeStore.new(path: project_root.join("runtime/recipes"))
)
model = Xnix::Compatibility::KRunnerModel.new(runtime: runtime, query: "notepad").to_h

assert(model["version"] == "0.2.89", "KRunner model must expose the current version")
assert(model["query_type"] == "krunner-query-plan", "KRunner model must expose the Runtime query plan type")
assert(model["entry_point"] == "krunner", "KRunner model must identify the KDE entry point")
assert(model["desktop"] == "KDE Plasma", "KRunner model must stay scoped to KDE Plasma")
assert(model["source"]["kind"] == "runtime-local-read-model", "KRunner model must describe local fallback reads")
assert(model["runtime_owned"], "KRunner model must preserve Runtime ownership")
assert(!model["kde_policy_owner"], "KRunner model must not claim KDE policy ownership")
assert(model["summary"]["runtime_owned_launch"], "KRunner model must keep launch ownership in the Runtime")
assert(!model["summary"]["query_execution_enabled"], "KRunner model must keep direct query execution disabled")
assert(!model["summary"]["backend_launch_enabled"], "KRunner model must keep backend launch disabled")
assert(!model["summary"]["backend_details_exposed"], "KRunner model must hide backend details")

matches = model.fetch("matches")
assert(matches.length == 1, "KRunner model must return the sample application")
match = matches.first
assert(match["application_id"] == "org.xnix.sample.notepad", "KRunner model must resolve the Runtime application id")
assert(match["name"] == "Sample Notepad", "KRunner model must preserve the user-facing application name")
assert(match["subtitle"] == "Open as a normal Linux application", "KRunner model must present normal desktop wording")
assert(match["mode_label"] == "Automatic", "KRunner model must use user-facing mode labels")
assert(match["supported_extensions"].include?(".txt"), "KRunner model must include file extension hints")
assert(match["relevance_percent"] == 100, "KRunner model must preserve Runtime relevance percentages")
assert(match["action"]["type"] == "runtime-launch", "KRunner model must emit a Runtime launch action")
assert(match["action"]["argv"] == ["xnix-compat-launch", "--app", "org.xnix.sample.notepad"], "KRunner model must delegate launch to the managed launcher")
assert(match["action"]["desktop_entry_id"] == "org.xnix.sample.notepad.desktop", "KRunner model must map matches to desktop entries")

plan_runtime = QueryPlanRuntime.new
plan_model = Xnix::Compatibility::KRunnerModel.new(runtime: plan_runtime, query: "runtime owned").to_h
assert(plan_runtime.queried == "runtime owned", "KRunner model must query the Runtime query plan")
assert(plan_model["source"]["kind"] == "runtime-query-plan-test", "KRunner model must preserve Runtime source metadata")
assert(plan_model["matches"].first["application_id"] == "org.xnix.sample.notepad", "KRunner model must use Runtime query plan matches")
assert(!plan_model["summary"]["query_execution_enabled"], "KRunner model must preserve Runtime query execution gates")

flat_model = Xnix::Compatibility::KRunnerModel.new(runtime: FlatQueryPlanRuntime.new, query: "notepad").to_h
assert(flat_model["matches"].first["application_id"] == "org.xnix.sample.notepad", "KRunner model must normalize flat D-Bus query plans")
assert(flat_model["matches"].first["action"]["argv"] == ["xnix-compat-launch", "--app", "org.xnix.sample.notepad"], "KRunner model must rebuild managed launcher actions from flat D-Bus query plans")

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
