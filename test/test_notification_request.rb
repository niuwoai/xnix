#!/usr/bin/env ruby
# frozen_string_literal: true

require "json"
require "open3"
require "pathname"
require_relative "../lib/xnix/compatibility/notification_request"

def assert(condition, message)
  return if condition

  warn "FAIL: #{message}"
  exit 1
end

project_root = Pathname.new(__dir__).join("..").realpath
request = Xnix::Compatibility::NotificationRequest.new(
  application_id: "org.xnix.sample.notepad",
  event_type: "install-failed"
).to_h

assert(request["request_type"] == "desktop-notification", "notification request must identify the request type")
assert(request["source"] == "runtime-event", "notification request must identify Runtime events as the source")
assert(request["desktop"] == "KDE Plasma", "notification request must target KDE Plasma")
assert(request["application_id"] == "org.xnix.sample.notepad", "notification request must preserve the Runtime application id")
assert(request["event_type"] == "install-failed", "notification request must preserve the event type")
assert(request["urgency"] == "critical", "install failures must be critical")
assert(request["actions"].include?("show-diagnostics"), "install failures must offer diagnostics")

approval = Xnix::Compatibility::NotificationRequest.new(
  application_id: "org.xnix.sample.notepad",
  event_type: "approval-required",
  body: "File access needs approval."
).to_h
assert(approval["actions"].include?("review-request"), "approval notifications must offer request review")
assert(approval["body"] == "File access needs approval.", "notification request must allow a safe body override")

json = JSON.pretty_generate(approval)
assert(!json.match?(/prefix|\.wine|proton|virtual machine/i), "notification request must not expose backend implementation terms")

stdout, stderr, status = Open3.capture3(
  "ruby",
  project_root.join("bin/xnix-compat-notify").to_s,
  "--app",
  "org.xnix.sample.notepad",
  "--event",
  "mode-changed"
)
assert(status.success?, "notification CLI must exit successfully: #{stderr}")
cli_request = JSON.parse(stdout)
assert(cli_request["event_type"] == "mode-changed", "notification CLI must emit the requested event")
assert(cli_request["urgency"] == "low", "mode-change notifications must be low urgency")

_stdout, stderr, status = Open3.capture3(
  "ruby",
  project_root.join("bin/xnix-compat-notify").to_s,
  "--app",
  "org.xnix.sample.notepad",
  "--event",
  "unknown-event"
)
assert(!status.success?, "notification CLI must reject unknown events")
assert(stderr.include?("event type must be one of"), "notification CLI must explain valid event types")

puts "PASS: compatibility notification request unit tests"
