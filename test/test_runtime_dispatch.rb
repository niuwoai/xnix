#!/usr/bin/env ruby
# frozen_string_literal: true

require "json"
require "open3"
require "pathname"

def assert(condition, message)
  return if condition

  warn "FAIL: #{message}"
  exit 1
end

project_root = Pathname.new(__dir__).join("..").realpath
command = ["ruby", project_root.join("bin/xnix-compatd").to_s]

stdout, stderr, status = Open3.capture3(*command, "dispatch", "ListApplications")
assert(status.success?, "dispatch ListApplications must exit successfully: #{stderr}")
applications = JSON.parse(stdout)
assert(applications.first["id"] == "org.xnix.sample.notepad", "dispatch must route ListApplications")

stdout, stderr, status = Open3.capture3(
  *command,
  "dispatch",
  "GetApplication",
  JSON.generate(["org.xnix.sample.notepad"])
)
assert(status.success?, "dispatch GetApplication must exit successfully: #{stderr}")
application = JSON.parse(stdout)
assert(application["name"] == "Sample Notepad", "dispatch must route GetApplication")

stdout, stderr, status = Open3.capture3(
  *command,
  "dispatch",
  "GetDiagnostics",
  JSON.generate(["org.xnix.sample.notepad"])
)
assert(status.success?, "dispatch GetDiagnostics must exit successfully: #{stderr}")
diagnostics = JSON.parse(stdout)
assert(diagnostics["application_id"] == "org.xnix.sample.notepad", "dispatch must route GetDiagnostics")

_stdout, stderr, status = Open3.capture3(*command, "dispatch", "Launch", JSON.generate(["org.xnix.sample.notepad", {}]))
assert(!status.success?, "dispatch must reject unsupported write methods until a backend exists")
assert(stderr.include?("unsupported runtime method"), "dispatch must explain unsupported methods")

puts "PASS: compatibility runtime dispatch unit tests"
