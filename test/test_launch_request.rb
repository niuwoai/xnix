#!/usr/bin/env ruby
# frozen_string_literal: true

require "json"
require "open3"
require "pathname"
require_relative "../lib/xnix/compatibility/launch_request"
require_relative "../lib/xnix/compatibility/recipe_store"

def assert(condition, message)
  return if condition

  warn "FAIL: #{message}"
  exit 1
end

project_root = Pathname.new(__dir__).join("..").realpath
store = Xnix::Compatibility::RecipeStore.new(path: project_root.join("runtime/recipes"))
request = Xnix::Compatibility::LaunchRequest.new(recipe_store: store).build(
  application_id: "org.xnix.sample.notepad"
)

assert(request["version"] == "0.2.224", "launch request must expose the current version")
assert(request["intent_type"] == "runtime-launch-intent", "launch request must identify the Runtime launch intent")
assert(request["request_type"] == "launch-application", "launch request must identify the request type")
assert(request["source"] == "desktop-launcher", "launch request must identify the desktop source")
assert(request["application_id"] == "org.xnix.sample.notepad", "launch request must keep the Runtime application id")
assert(request["runtime_method"] == "Launch", "launch request must target the Runtime launch method")
assert(request["read_method"] == "GetLaunchIntent", "launch request must expose the Runtime read method")
assert(!request["portal_required"], "plain launcher requests must not require file portal access")
assert(request["run_plan"]["plan_type"] == "compatibility-run", "launch request must include the compatibility run plan")
assert(request["run_plan"]["strategy"] == "automatic-managed", "launch request must expose a desktop-safe run strategy")
assert(!request["run_plan"]["backend_details_exposed"], "launch request must not expose backend details")
assert(!request["run_plan"]["backend_ready"], "launch request must not claim backend readiness")
assert(request["run_plan"]["portal_policy_required"], "launch request must require Portal policy preflight")
assert(request["run_plan"]["snapshot_before_risky_change"], "launch request must require snapshot preflight")
assert(request["file_count"] == 0, "plain launcher requests must not include files")
assert(request["runtime_owned"], "launch request must be Runtime-owned")
assert(request["c_runtime_backed"], "launch request must align with the C Runtime boundary")
assert(!request["kde_policy_owner"], "launch request must not be KDE-owned")
assert(request["standard_desktop_entry"], "launch request must preserve standard desktop entry semantics")
assert(request["launch_uses_runtime"], "launch request must route through the Runtime")
assert(request["desktop_entry_launch_visible"], "launch request must keep KDE launcher visibility")
assert(!request["launch_allowed"], "launch request must not allow launch before Runtime gates pass")
assert(!request["launch_enabled"], "launch request must not enable launch")
assert(!request["execution_request_created"], "launch request must not create Runtime request objects")
assert(!request["execution_started"], "launch request must not start execution")
assert(request["write_gate_decision"] == "blocked-until-production-backend", "launch request must expose the Launch write gate")
assert(request["denial_error_name"] == "org.xnix.Compatibility1.Error.WriteMethodDisabled", "launch request must expose the write-gate error")
assert(request["blocked_actions"].include?("start Wine or VM backend from KDE"), "launch request must block KDE backend starts")
assert(!request["host_root_modified"], "launch request must not mutate the host root")
assert(!request["network_required"], "launch request must not require network access")
assert(!request["backend_details_exposed"], "launch request must hide backend details")

file_request = Xnix::Compatibility::LaunchRequest.new(recipe_store: store).build(
  application_id: "org.xnix.sample.notepad",
  file_uris: ["file:///home/test/Documents/example.txt"]
)
assert(file_request["portal_required"], "launcher file requests must require portal-mediated access")
assert(file_request["file_count"] == 1, "launcher file requests must count selected files")
assert(file_request["file_uris"].first == "file:///home/test/Documents/example.txt", "launcher file requests must preserve file URIs")

json = JSON.pretty_generate(file_request)
assert(!json.match?(/prefix|\.wine|proton|virtual machine/i), "launch request must not expose backend implementation terms")

stdout, stderr, status = Open3.capture3(
  "ruby",
  project_root.join("bin/xnix-compat-launch").to_s,
  "--app",
  "org.xnix.sample.notepad",
  "file:///home/test/Documents/example.txt"
)
assert(status.success?, "launch CLI must exit successfully: #{stderr}")
cli_request = JSON.parse(stdout)
assert(cli_request == file_request, "launch CLI must emit the request model")

_stdout, stderr, status = Open3.capture3("ruby", project_root.join("bin/xnix-compat-launch").to_s)
assert(!status.success?, "launch CLI must require an application id")
assert(stderr.include?("--app is required"), "launch CLI must explain missing application ids")

puts "PASS: compatibility launch request unit tests"
