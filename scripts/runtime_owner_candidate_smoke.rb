#!/usr/bin/env ruby
# frozen_string_literal: true

require "json"
require "open3"

def assert(condition, message)
  return if condition

  warn "FAIL: #{message}"
  exit 1
end

def executable_in_path?(name)
  ENV.fetch("PATH", "").split(File::PATH_SEPARATOR).any? do |path|
    candidate = File.join(path, name)
    File.file?(candidate) && File.executable?(candidate)
  end
end

def runtime_owner_command
  override = ENV["XNIX_RUNTIME_OWNER_BIN"].to_s
  return [override] unless override.empty?
  return ["xnix-runtime-owner"] if executable_in_path?("xnix-runtime-owner")

  ["go", "run", "./cmd/xnix-runtime-owner"]
end

unless ENV["DBUS_SESSION_BUS_ADDRESS"]
  assert(executable_in_path?("dbus-run-session"), "Runtime owner candidate smoke requires dbus-run-session when no session bus is active")
  stdout, stderr, status = Open3.capture3("dbus-run-session", "--", "ruby", __FILE__)
  print stdout
  warn stderr unless stderr.empty?
  exit status.exitstatus
end

owner_command = runtime_owner_command
stdout, stderr, status = Open3.capture3(*owner_command, "--root", ".", "--mode", "smoke-owner")
assert(status.success?, "Runtime owner candidate smoke must render smoke-owner JSON: #{stderr}")

candidate = JSON.parse(stdout)
assert(candidate.fetch("version") == File.read("VERSION").strip, "Runtime owner candidate smoke must expose the current version")
assert(candidate.fetch("schema_version") == "xnix.runtime.owner_candidate.v1", "Runtime owner candidate smoke must expose the owner candidate schema")
assert(candidate.fetch("request_type") == "runtime-owner-candidate", "Runtime owner candidate smoke must expose the request type")
assert(candidate.fetch("owner_type") == "go-runtime-owner-candidate", "Runtime owner candidate smoke must expose the owner type")
assert(candidate.fetch("mode") == "smoke-owner", "Runtime owner candidate smoke must run in smoke-owner mode")
assert(candidate.fetch("bus_name") == "org.xnix.Compatibility1", "Runtime owner candidate smoke must preserve the Runtime bus name")
assert(candidate.fetch("service_activation_ready"), "Runtime owner candidate smoke must see activation readiness")
assert(candidate.fetch("read_only_route_table_ready"), "Runtime owner candidate smoke must see route table readiness")
assert(candidate.fetch("read_only_serve_ready"), "Runtime owner candidate smoke must see read-only serve readiness")
assert(candidate.fetch("route_count") == 57, "Runtime owner candidate smoke must see all Runtime read-only routes")
assert(candidate.fetch("go_route_count") == 57, "Runtime owner candidate smoke must see all routes as Go-routed")
assert(candidate.fetch("c_core_route_count").zero?, "Runtime owner candidate smoke must not require C owner adapters")
assert(candidate.fetch("ruby_legacy_route_count").zero?, "Runtime owner candidate smoke must not require Ruby legacy routes")
assert(candidate.fetch("write_method_count") == 4, "Runtime owner candidate smoke must see every reserved write method")
assert(!candidate.fetch("write_methods_enabled"), "Runtime owner candidate smoke must keep write methods disabled")
assert(candidate.fetch("runtime_owned"), "Runtime owner candidate smoke must keep Runtime ownership in the Runtime")
assert(candidate.fetch("go_runtime_backed"), "Runtime owner candidate smoke must report Go backing")
assert(!candidate.fetch("kde_policy_owner"), "Runtime owner candidate smoke must not make KDE the policy owner")
assert(!candidate.fetch("kde_may_claim_runtime_ownership"), "Runtime owner candidate smoke must not let KDE claim Runtime ownership")
assert(candidate.fetch("smoke_owner_mode"), "Runtime owner candidate smoke must report smoke owner mode")
assert(!candidate.fetch("production_owner_mode"), "Runtime owner candidate smoke must not report production owner mode")
assert(!candidate.fetch("event_loop_started"), "Runtime owner candidate smoke must not start an event loop")
assert(!candidate.fetch("session_bus_claimed"), "Runtime owner candidate smoke must not claim the session bus")
assert(!candidate.fetch("production_bus_claimed"), "Runtime owner candidate smoke must not claim the production bus")
assert(!candidate.fetch("system_service_started"), "Runtime owner candidate smoke must not start a system service")
assert(!candidate.fetch("network_required"), "Runtime owner candidate smoke must not require network")
assert(!candidate.fetch("host_root_modified"), "Runtime owner candidate smoke must not mutate the host root")
assert(!candidate.fetch("privileged_container_required"), "Runtime owner candidate smoke must not require privileged containers")
assert(!candidate.fetch("backend_details_exposed"), "Runtime owner candidate smoke must not expose backend details")
assert(candidate.fetch("counts").fetch("passed") == 5, "Runtime owner candidate smoke must pass five checks")
assert(candidate.fetch("counts").fetch("pending") == 1, "Runtime owner candidate smoke must keep production bus claim pending")

candidate.fetch("write_methods").each do |write_method|
  assert(write_method.fetch("error_name") == "org.xnix.Compatibility1.Error.WriteMethodDisabled", "Runtime owner candidate smoke must expose deterministic write errors")
  assert(!write_method.fetch("dispatch_enabled"), "Runtime owner candidate smoke must not enable write dispatch")
  assert(!write_method.fetch("request_created"), "Runtime owner candidate smoke must not create write requests")
end

stdout, stderr, status = Open3.capture3(*owner_command, "--deny-write", "Launch")
assert(status.success?, "Runtime owner candidate smoke must render disabled write responses: #{stderr}")
write_response = JSON.parse(stdout)
assert(write_response.fetch("method") == "Launch", "Runtime owner candidate smoke must echo the denied write method")
assert(write_response.fetch("error_name") == "org.xnix.Compatibility1.Error.WriteMethodDisabled", "Runtime owner candidate smoke must use the stable disabled-write error")
assert(!write_response.fetch("dispatch_enabled"), "Runtime owner candidate smoke must keep denied write dispatch disabled")
assert(!write_response.fetch("request_created"), "Runtime owner candidate smoke must keep denied write request creation disabled")

stdout, stderr, status = Open3.capture3(*owner_command, "--root", ".", "--dispatch-read", "GetRuntimeWriteGate", "Launch")
assert(status.success?, "Runtime owner candidate smoke must render read-only owner dispatch responses: #{stderr}")
read_dispatch = JSON.parse(stdout)
assert(read_dispatch.fetch("schema_version") == "xnix.runtime.owner_read_dispatch.v1", "Runtime owner candidate smoke must expose the read dispatch schema")
assert(read_dispatch.fetch("method") == "GetRuntimeWriteGate", "Runtime owner candidate smoke must dispatch the requested read method")
assert(read_dispatch.fetch("read_only_dispatch"), "Runtime owner candidate smoke must mark owner dispatch as read-only")
assert(!read_dispatch.fetch("write_methods_enabled"), "Runtime owner candidate smoke must keep write methods disabled during read dispatch")
assert(!read_dispatch.fetch("event_loop_started"), "Runtime owner candidate smoke must not start an event loop for read dispatch")
assert(!read_dispatch.fetch("session_bus_claimed"), "Runtime owner candidate smoke must not claim the session bus for read dispatch")
assert(!read_dispatch.fetch("production_bus_claimed"), "Runtime owner candidate smoke must not claim the production bus for read dispatch")
dispatch_payload = read_dispatch.fetch("payload")
assert(dispatch_payload.fetch("request_type") == "runtime-write-gate-preview", "Runtime owner candidate smoke must include the dispatched Runtime payload")
assert(dispatch_payload.fetch("method_name") == "Launch", "Runtime owner candidate smoke must pass read dispatch arguments")
assert(!dispatch_payload.fetch("dispatch_enabled"), "Runtime owner candidate smoke must preserve nested write dispatch gates")

puts "PASS: Runtime owner candidate restricted session smoke"
