#!/usr/bin/env ruby
# frozen_string_literal: true

require "pathname"

def assert(condition, message)
  return if condition

  warn "FAIL: #{message}"
  exit 1
end

project_root = Pathname.new(__dir__).join("..").realpath
script = project_root.join("scripts/runtime_owner_candidate_smoke.rb")
dockerfile = project_root.join("Dockerfile").read
contents = script.read

assert(script.file?, "Runtime owner candidate smoke script must exist")
assert(contents.include?("dbus-run-session"), "Runtime owner candidate smoke must run inside a restricted session bus")
assert(contents.include?("xnix-runtime-owner"), "Runtime owner candidate smoke must execute the Go owner candidate")
assert(contents.include?("smoke-owner"), "Runtime owner candidate smoke must request smoke-owner mode")
assert(contents.include?("dispatch-read"), "Runtime owner candidate smoke must verify owner read dispatch")
assert(contents.include?("xnix.runtime.owner_read_dispatch.v1"), "Runtime owner candidate smoke must validate the read dispatch schema")
assert(contents.include?("xnix.runtime.owner_candidate.v1"), "Runtime owner candidate smoke must validate the owner candidate schema")
assert(contents.include?("org.xnix.Compatibility1.Error.WriteMethodDisabled"), "Runtime owner candidate smoke must validate deterministic disabled write errors")
assert(contents.include?("session_bus_claimed"), "Runtime owner candidate smoke must verify that the candidate does not claim the session bus")
assert(contents.include?("production_bus_claimed"), "Runtime owner candidate smoke must verify that the candidate does not claim the production bus")
assert(contents.include?("host_root_modified"), "Runtime owner candidate smoke must verify the host-root mutation gate")
assert(contents.include?("privileged_container_required"), "Runtime owner candidate smoke must verify the privileged-container gate")
assert(dockerfile.include?("go build -o /usr/local/bin/xnix-runtime-owner"), "Dockerfile must install the Go owner candidate binary for container smoke")

puts "PASS: Runtime owner candidate smoke script unit tests"
