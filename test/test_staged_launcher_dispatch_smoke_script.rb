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
script = project_root.join("scripts/staged_launcher_dispatch_smoke.rb")
container_script = project_root.join("scripts/container.rb")
container_model = project_root.join("lib/xnix/container.rb")
contents = script.read

assert(script.executable?, "Staged launcher dispatch smoke script must be executable")
assert(contents.include?("windows-known-app-dispatch-preview"), "Staged launcher dispatch smoke must preflight the managed artifact")
assert(contents.include?("desktop-activation-stage"), "Staged launcher dispatch smoke must stage desktop activation files")
assert(contents.include?("--managed-launcher-bin"), "Staged launcher dispatch smoke must copy the Go launcher binary")
assert(contents.include?("usr/local/bin/xnix-compat-launch"), "Staged launcher dispatch smoke must execute the staged launcher")
assert(contents.include?("--guest-boundary"), "Staged launcher dispatch smoke must supply the controlled guest boundary")
assert(contents.include?("managed-known-app-guest-smoke"), "Staged launcher dispatch smoke must use the known app guest smoke boundary")
assert(contents.include?("windows-known-app-dispatch-smoke"), "Staged launcher dispatch smoke must enter the dispatch smoke path")
assert(contents.include?("SKIP:"), "Staged launcher dispatch smoke must skip when external smoke prerequisites are missing")
assert(contents.include?("PASS:"), "Staged launcher dispatch smoke must print a pass marker")
assert(container_script.read.include?("staged-launcher-dispatch-smoke"), "Container CLI must expose staged-launcher-dispatch-smoke")
assert(container_model.read.include?("staged_launcher_dispatch_smoke_command"), "Container model must expose staged launcher dispatch smoke command")

stdout, stderr, status = Open3.capture3("ruby", script.to_s)
assert(status.success?, "Staged launcher dispatch smoke script must pass or skip cleanly: #{stderr}")
assert(stdout.include?("PASS: staged managed launcher dispatch smoke") || stdout.include?("SKIP: staged managed launcher dispatch smoke"),
       "Staged launcher dispatch smoke script must print a pass or skip marker")

puts "PASS: staged launcher dispatch smoke script unit tests"
