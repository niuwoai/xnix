#!/usr/bin/env ruby
# frozen_string_literal: true

require "json"
require "open3"
require "pathname"
require_relative "../lib/xnix/compatibility/portal_request_model"

def assert(condition, message)
  return if condition

  warn "FAIL: #{message}"
  exit 1
end

project_root = Pathname.new(__dir__).join("..").realpath
operations = Xnix::Compatibility::PortalAccessPolicy::OPERATIONS.keys

operations.each do |operation|
  request = Xnix::Compatibility::PortalRequestModel.new(
    application_id: "org.xnix.sample.notepad",
    operation: operation,
    reason: "Open a selected document."
  ).to_h

  assert(request["version"] == "0.2.122", "Portal request model must expose the current version")
  assert(request["request_type"] == "portal-request", "Portal request model must identify the request type")
  assert(request["desktop"] == "KDE Plasma", "Portal request model must target KDE Plasma")
  assert(request["application_id"] == "org.xnix.sample.notepad", "Portal request model must preserve the application id")
  assert(request["operation"] == operation, "Portal request model must preserve the operation")
  assert(request["portal"]["destination"] == "org.freedesktop.portal.Desktop", "Portal request model must target the portal service")
  assert(request["portal"]["interface"].start_with?("org.freedesktop.portal."), "Portal request model must expose the portal interface")
  assert(!request["portal"]["method"].empty?, "Portal request model must expose the portal method")
  assert(request["request"]["object_path_required"], "Portal request model must require request object paths")
  assert(request["request"]["user_mediation_required"], "Portal request model must require user mediation")
  assert(request["request"]["runtime_policy_owner"], "Runtime must own Portal request policy")
  assert(!request["request"]["desktop_shell_policy_owner"], "KDE must not own Runtime Portal request policy")
  assert(request["completion"]["signal"] == "Response", "Portal request model must wait for portal Response signals")
  assert(request["completion"]["result_owner"] == "Runtime", "Portal request results must return to the Runtime")
  assert(!request["safety"]["direct_access_allowed"], "Portal request model must deny direct access")
  assert(request["safety"]["portal_required"], "Portal request model must require portals")
  assert(!request["safety"]["backend_details_exposed"], "Portal request model must hide backend details")
  assert(!request["safety"]["host_permission_changed"], "Portal request model must not mutate host permissions")
end

file_request = Xnix::Compatibility::PortalRequestModel.new(
  application_id: "org.xnix.sample.notepad",
  operation: "file-open",
  reason: "Open a selected document."
).to_h
assert(file_request["decision"] == "ask", "file-open Portal requests must ask by default")
assert(file_request["request_allowed"], "ask decisions must allow a mediated request")
assert(file_request["portal"]["method"] == "OpenFile", "file-open Portal requests must call OpenFile")
assert(file_request["request"]["handle_token"] == "xnix_org_xnix_sample_notepad_file_open", "Portal request handle token must be deterministic")

camera_request = Xnix::Compatibility::PortalRequestModel.new(
  application_id: "org.xnix.sample.notepad",
  operation: "camera"
).to_h
assert(camera_request["decision"] == "deny", "camera Portal requests must default to deny")
assert(!camera_request["request_allowed"], "deny decisions must not allow a mediated request")
assert(camera_request["denied"]["next_action"] == "open-compatibility-settings", "denied Portal requests must guide users to settings")

json = JSON.pretty_generate(file_request)
assert(!json.match?(/wine|prefix|\.wine|proton|virtual machine/i), "Portal request model must not expose backend implementation terms")

stdout, stderr, result = Open3.capture3(
  "ruby",
  project_root.join("bin/xnix-portal-request-model").to_s,
  "--app",
  "org.xnix.sample.notepad",
  "--operation",
  "file-open",
  "--reason",
  "Open a selected document."
)
assert(result.success?, "Portal request model CLI must exit successfully: #{stderr}")
assert(JSON.parse(stdout) == file_request, "Portal request model CLI must emit the request model")

_stdout, stderr, result = Open3.capture3(
  "ruby",
  project_root.join("bin/xnix-portal-request-model").to_s,
  "--app",
  "invalid",
  "--operation",
  "file-open"
)
assert(!result.success?, "Portal request model CLI must reject invalid application ids")
assert(stderr.include?("reverse-DNS"), "Portal request model CLI must explain application id validation")

puts "PASS: compatibility Portal request model unit tests"
