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
script = project_root.join("scripts/runtime_status_owner_service_session_bus_smoke.rb")
inner_script = project_root.join("scripts/staged_launcher_dispatch_smoke.rb")
container_script = project_root.join("scripts/container.rb")
container_model = project_root.join("lib/xnix/container.rb")
contents = script.read
inner_contents = inner_script.read

assert(script.executable?, "Runtime-status owner service session-bus smoke script must be executable")
assert(contents.include?("dbus-run-session"), "Runtime-status owner service session-bus smoke must run inside a private session bus")
assert(contents.include?("DBUS_SESSION_BUS_ADDRESS"), "Runtime-status owner service session-bus smoke must require a session bus address")
assert(contents.include?("scripts/staged_launcher_dispatch_smoke.rb"), "Runtime-status owner service session-bus smoke must reuse the real staged launcher dispatch smoke")
assert(contents.include?("PASS: staged managed launcher dispatch smoke"), "Runtime-status owner service session-bus smoke must accept the staged launcher pass marker")
assert(contents.include?("SKIP: staged managed launcher dispatch smoke"), "Runtime-status owner service session-bus smoke must accept the staged launcher skip marker")
assert(contents.include?("PASS: Runtime-status owner service session-bus smoke"), "Runtime-status owner service session-bus smoke must print a pass marker")
assert(contents.include?("SKIP: Runtime-status owner service session-bus smoke"), "Runtime-status owner service session-bus smoke must print a skip marker")
assert(contents.include?("docker.sock"), "Runtime-status owner service session-bus smoke must guard Docker socket exposure")
assert(contents.include?("--privileged"), "Runtime-status owner service session-bus smoke must guard privileged container exposure")
assert(contents.include?("--network host"), "Runtime-status owner service session-bus smoke must guard host networking exposure")

assert(inner_contents.include?("xnix-runtime-owner"), "Inner staged smoke must route through the Runtime owner binary")
assert(inner_contents.include?("--service-call"), "Inner staged smoke must call the owner service boundary")
assert(inner_contents.include?("ShowRuntimeControlledLaunch"), "Inner staged smoke must call ShowRuntimeControlledLaunch")
assert(inner_contents.include?("evidence-relative-path"), "Inner staged smoke must forward only the Runtime-status evidence path")
assert(inner_contents.include?("desktop-trigger-service-call-materialization-preview"), "Inner staged smoke must consume the Go-owned service call materialization packet")
assert(inner_contents.include?("--human-authorized-smoke"), "Inner staged smoke must explicitly request a human-authorized smoke candidate")
assert(inner_contents.include?("owner_service_cli_args"), "Inner staged smoke must use Go-materialized owner service CLI arguments")
assert(inner_contents.include?("runtime-owner-service-call"), "Inner staged smoke must verify the owner service envelope")
assert(inner_contents.include?("desktop-action-dispatch"), "Inner staged smoke must verify desktop action dispatch")
assert(inner_contents.include?("kde-dbus-runtime-status-action"), "Inner staged smoke must expose the D-Bus callable owner route")
assert(inner_contents.include?("kde_forwards_only_evidence_handle"), "Inner staged smoke must verify KDE forwards only the evidence handle")
assert(inner_contents.include?("runtime_owner_service_supplies_owner_inputs"), "Inner staged smoke must verify owner-only inputs are supplied by Runtime")

assert(container_script.read.include?("runtime-status-owner-service-session-bus-smoke"), "Container CLI must expose runtime-status-owner-service-session-bus-smoke")
assert(container_model.read.include?("runtime_status_owner_service_session_bus_smoke_command"), "Container model must expose Runtime-status owner service session-bus smoke command")

stdout, stderr, status = Open3.capture3("ruby", script.to_s)
assert(status.success?, "Runtime-status owner service session-bus smoke must pass or skip cleanly: #{stderr}")
assert(stdout.include?("PASS: Runtime-status owner service session-bus smoke") || stdout.include?("SKIP: Runtime-status owner service session-bus smoke"),
       "Runtime-status owner service session-bus smoke must print a pass or skip marker")

puts "PASS: Runtime-status owner service session-bus smoke script unit tests"
