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
script = project_root.join("scripts/runtime_activation_smoke.rb")
contents = script.read

assert(script.executable?, "Runtime activation smoke script must be executable")
assert(contents.include?("install_runtime_activation.rb"), "Runtime activation smoke must install activation files under a staging root")
assert(contents.include?("usr/libexec/xnix/compatd"), "Runtime activation smoke must execute the staged libexec wrapper")
assert(contents.include?("usr/lib/xnix/compatibility/runtime_daemon.rb"), "Runtime activation smoke must require installed Runtime libraries")
assert(contents.include?("GetRuntimeMethodParityManifest"), "Runtime activation smoke must verify staged read-only dispatch")
assert(contents.include?("GetRuntimeWriteGate"), "Runtime activation smoke must verify staged write gate dispatch")
assert(contents.include?("WriteMethodDisabled"), "Runtime activation smoke must verify gated write rejection")

stdout, stderr, status = Open3.capture3("ruby", script.to_s)
assert(status.success?, "Runtime activation smoke script must pass: #{stderr}")
assert(stdout.include?("PASS: Runtime activation smoke"), "Runtime activation smoke script must print a pass marker")

puts "PASS: Runtime activation smoke script unit tests"
