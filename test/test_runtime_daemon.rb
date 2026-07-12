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

stdout, stderr, status = Open3.capture3(*command, "probe")
assert(status.success?, "runtime daemon probe must exit successfully: #{stderr}")
probe = JSON.parse(stdout)
assert(probe["version"] == "0.2.4", "runtime daemon probe must report the current version")
assert(probe["bus_name"] == "org.xnix.Compatibility1", "runtime daemon probe must keep the stable bus name")
assert(probe["capabilities"]["recipe_store"], "runtime daemon probe must expose recipe store capability")
assert(probe["capabilities"]["dbus_method_dispatch"], "runtime daemon probe must expose method dispatch capability")
assert(!probe["capabilities"]["dbus_binding"], "runtime daemon must not claim a D-Bus binding before it exists")

stdout, stderr, status = Open3.capture3(*command, "list-applications")
assert(status.success?, "runtime daemon list must exit successfully: #{stderr}")
applications = JSON.parse(stdout)
assert(applications.length == 1, "runtime daemon must list the bundled sample application")
assert(applications.first["id"] == "org.xnix.sample.notepad", "runtime daemon must expose sample recipe id")
assert(applications.first["mime_types"].include?("application/x-xnix-txt"), "runtime daemon must expose MIME types")

stdout, stderr, status = Open3.capture3(*command, "diagnostics", "org.xnix.sample.notepad")
assert(status.success?, "runtime daemon diagnostics must exit successfully: #{stderr}")
diagnostics = JSON.parse(stdout)
assert(diagnostics["status"] == "known", "runtime daemon diagnostics must know bundled recipes")
assert(diagnostics["checks"].any? { |check| check["status"] == "pending" }, "runtime daemon must report backend work as pending")

_stdout, stderr, status = Open3.capture3(*command, "diagnostics", "org.xnix.missing")
assert(!status.success?, "runtime daemon must reject unknown applications")
assert(stderr.include?("unknown application"), "runtime daemon must explain unknown applications")

puts "PASS: compatibility runtime daemon unit tests"
