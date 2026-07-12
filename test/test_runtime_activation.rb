#!/usr/bin/env ruby
# frozen_string_literal: true

require "open3"
require "pathname"

def assert(condition, message)
  return if condition

  warn "FAIL: #{message}"
  exit 1
end

project_root = Pathname.new(__dir__).join("..").realpath
service = project_root.join("runtime/dbus/org.xnix.Compatibility1.service").read
unit = project_root.join("runtime/systemd/xnix-compatd.service").read
wrapper = project_root.join("libexec/xnix/compatd")

assert(service.include?("Exec=/usr/libexec/xnix/compatd"), "D-Bus service must point at the packaged libexec wrapper")
assert(unit.include?("ExecStart=/usr/libexec/xnix/compatd"), "systemd unit must point at the packaged libexec wrapper")
assert(wrapper.file?, "packaged libexec wrapper must exist")
assert(wrapper.executable?, "packaged libexec wrapper must be executable")

stdout, stderr, status = Open3.capture3("ruby", wrapper.to_s, "probe")
assert(status.success?, "packaged libexec wrapper must run the Runtime probe: #{stderr}")
assert(stdout.include?("\"bus_name\": \"org.xnix.Compatibility1\""), "packaged wrapper probe must expose the bus name")

puts "PASS: compatibility runtime activation unit tests"
