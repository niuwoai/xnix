#!/usr/bin/env ruby
# frozen_string_literal: true

require "json"
require "open3"
require "pathname"
require_relative "../lib/xnix/compatibility/portal_access_policy"

def assert(condition, message)
  return if condition

  warn "FAIL: #{message}"
  exit 1
end

project_root = Pathname.new(__dir__).join("..").realpath
operations = Xnix::Compatibility::PortalAccessPolicy::OPERATIONS.keys

operations.each do |operation|
  policy = Xnix::Compatibility::PortalAccessPolicy.new(
    application_id: "org.xnix.sample.notepad",
    operation: operation
  ).to_h

  assert(policy["version"] == "0.2.79", "portal access policy must expose the current version")
  assert(policy["policy_type"] == "portal-access", "portal access policy must identify the model type")
  assert(policy["desktop"] == "KDE Plasma", "portal access policy must target KDE Plasma")
  assert(policy["application_id"] == "org.xnix.sample.notepad", "portal access policy must preserve the application id")
  assert(policy["operation"] == operation, "portal access policy must preserve the operation")
  assert(policy["portal_required"], "portal access policy must require portals")
  assert(policy["portal_interface"].start_with?("org.freedesktop.portal."), "portal access policy must identify the portal interface")
  assert(policy["user_mediation_required"], "portal access policy must require user mediation")
  assert(!policy["direct_access_allowed"], "portal access policy must deny direct desktop access")
  assert(policy["request_flow"]["dbus_api"] == "XDG Desktop Portal", "portal access policy must use the XDG Desktop Portal D-Bus API")
  assert(policy["request_flow"]["request_object_required"], "portal access policy must require request objects")
  assert(policy["request_flow"]["runtime_policy_owner"], "Runtime must own portal access policy")
  assert(!policy["request_flow"]["desktop_shell_policy_owner"], "KDE must not own Runtime access policy")
end

file_policy = Xnix::Compatibility::PortalAccessPolicy.new(
  application_id: "org.xnix.sample.notepad",
  operation: "file-open"
).to_h
assert(file_policy["decision"] == "ask", "file access must default to ask")
assert(file_policy["resources"].include?("selected-files"), "file access must scope selected files")

camera_policy = Xnix::Compatibility::PortalAccessPolicy.new(
  application_id: "org.xnix.sample.notepad",
  operation: "camera"
).to_h
assert(camera_policy["decision"] == "deny", "camera access must default to deny")

remote_desktop_policy = Xnix::Compatibility::PortalAccessPolicy.new(
  application_id: "org.xnix.sample.notepad",
  operation: "remote-desktop"
).to_h
assert(remote_desktop_policy["decision"] == "deny", "remote desktop access must default to deny")

json = JSON.pretty_generate(file_policy)
assert(!json.match?(/wine|prefix|\.wine|proton|virtual machine/i), "portal access policy must not expose backend implementation terms")

stdout, stderr, result = Open3.capture3(
  "ruby",
  project_root.join("bin/xnix-portal-access-policy").to_s,
  "--app",
  "org.xnix.sample.notepad",
  "--operation",
  "file-open"
)
assert(result.success?, "portal access policy CLI must exit successfully: #{stderr}")
assert(JSON.parse(stdout) == file_policy, "portal access policy CLI must emit the policy model")

_stdout, stderr, result = Open3.capture3(
  "ruby",
  project_root.join("bin/xnix-portal-access-policy").to_s,
  "--app",
  "invalid",
  "--operation",
  "file-open"
)
assert(!result.success?, "portal access policy CLI must reject invalid application ids")
assert(stderr.include?("reverse-DNS"), "portal access policy CLI must explain application id validation")

_stdout, stderr, result = Open3.capture3(
  "ruby",
  project_root.join("bin/xnix-portal-access-policy").to_s,
  "--app",
  "org.xnix.sample.notepad",
  "--operation",
  "unknown"
)
assert(!result.success?, "portal access policy CLI must reject unknown operations")
assert(stderr.include?("invalid argument"), "portal access policy CLI must explain operation validation")

puts "PASS: compatibility portal access policy unit tests"
