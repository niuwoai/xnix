#!/usr/bin/env ruby
# frozen_string_literal: true

require "pathname"

def assert(condition, message)
  return if condition

  warn "FAIL: #{message}"
  exit 1
end

project_root = Pathname.new(__dir__).join("..").realpath
script = project_root.join("scripts/dbus_session_smoke.rb")
source = project_root.join("runtime/dbus/xnix_compatd_smoke.c")
dockerfile = project_root.join("Dockerfile").read
contents = script.read

assert(script.file?, "D-Bus smoke script must exist")
assert(script.executable?, "D-Bus smoke script must be executable")
assert(source.file?, "D-Bus smoke adapter source must exist")
assert(dockerfile.include?("libglib2.0-dev"), "Dockerfile must install GIO headers for the smoke adapter")
assert(dockerfile.include?("pkg-config"), "Dockerfile must install pkg-config for the smoke adapter build")
assert(dockerfile.include?("xnix_compatd_smoke.c"), "Dockerfile must compile the smoke adapter")
assert(contents.include?("dbus-run-session"), "D-Bus smoke script must start a session bus")
assert(contents.include?("gdbus"), "D-Bus smoke script must use gdbus for runtime calls")
assert(contents.include?("ListApplications"), "D-Bus smoke script must call ListApplications")
assert(contents.include?("GetDiagnostics"), "D-Bus smoke script must call GetDiagnostics")
assert(contents.include?("GetEngineCatalog"), "D-Bus smoke script must call GetEngineCatalog")
assert(contents.include?("GetRunPlan"), "D-Bus smoke script must call GetRunPlan")
assert(contents.include?("GetRepairPlan"), "D-Bus smoke script must call GetRepairPlan")
assert(contents.include?("GetTestPlan"), "D-Bus smoke script must call GetTestPlan")
assert(contents.include?("GetSnapshotPlan"), "D-Bus smoke script must call GetSnapshotPlan")
assert(contents.include?("GetPortalAccessPolicy"), "D-Bus smoke script must call GetPortalAccessPolicy")

puts "PASS: compatibility runtime D-Bus smoke script unit tests"
