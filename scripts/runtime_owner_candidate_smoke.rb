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
assert(candidate.fetch("route_count") == 61, "Runtime owner candidate smoke must see all Runtime read-only routes")
assert(candidate.fetch("go_route_count") == 61, "Runtime owner candidate smoke must see all routes as Go-routed")
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

stdout, stderr, status = Open3.capture3(*owner_command, "--root", ".", "--mode", "smoke-owner", "--service-call", "GetRuntimeWriteGate", "Launch")
assert(status.success?, "Runtime owner candidate smoke must render in-process owner service calls: #{stderr}")
service_call = JSON.parse(stdout)
assert(service_call.fetch("schema_version") == "xnix.runtime.owner_service_call.v1", "Runtime owner service call must expose the service-call schema")
assert(service_call.fetch("request_type") == "runtime-owner-service-call", "Runtime owner service call must expose the request type")
assert(service_call.fetch("service_type") == "go-runtime-owner-in-process-service", "Runtime owner service call must expose the service type")
assert(service_call.fetch("method") == "GetRuntimeWriteGate", "Runtime owner service call must echo the read method")
assert(service_call.fetch("call_type") == "read-dispatch", "Runtime owner service call must route read methods through read dispatch")
assert(service_call.fetch("read_only_serve_ready"), "Runtime owner service call must report read-only serve readiness")
assert(service_call.fetch("read_only_dispatch"), "Runtime owner service call must mark read calls as read-only")
assert(!service_call.fetch("write_method"), "Runtime owner service call must not mark read calls as writes")
assert(!service_call.fetch("write_methods_enabled"), "Runtime owner service call must keep write methods disabled")
assert(service_call.fetch("dispatch_ready"), "Runtime owner service call must be dispatch-ready")
assert(service_call.fetch("in_process_service_ready"), "Runtime owner service call must report the in-process service boundary as ready")
assert(!service_call.fetch("event_loop_started"), "Runtime owner service call must not start an event loop yet")
assert(!service_call.fetch("session_bus_claimed"), "Runtime owner service call must not claim the session bus yet")
assert(!service_call.fetch("production_bus_claimed"), "Runtime owner service call must not claim the production bus")
assert(!service_call.fetch("system_service_started"), "Runtime owner service call must not start a system service")
assert(!service_call.fetch("network_required"), "Runtime owner service call must not require network")
assert(!service_call.fetch("host_root_modified"), "Runtime owner service call must not mutate the host root")
assert(!service_call.fetch("privileged_container_required"), "Runtime owner service call must not require privileged containers")
assert(!service_call.fetch("backend_details_exposed"), "Runtime owner service call must not expose backend details")
service_payload = service_call.fetch("payload")
assert(service_payload.fetch("schema_version") == "xnix.runtime.owner_read_dispatch.v1", "Runtime owner service call must include the nested read dispatch payload")
assert(service_payload.fetch("method") == "GetRuntimeWriteGate", "Runtime owner service call must preserve the nested read dispatch method")

stdout, stderr, status = Open3.capture3(*owner_command, "--root", ".", "--mode", "smoke-owner", "--service-call", "Launch")
assert(status.success?, "Runtime owner service call must render write denials: #{stderr}")
write_service_call = JSON.parse(stdout)
assert(write_service_call.fetch("schema_version") == "xnix.runtime.owner_service_call.v1", "Runtime owner write service call must expose the service-call schema")
assert(write_service_call.fetch("method") == "Launch", "Runtime owner write service call must echo the write method")
assert(write_service_call.fetch("call_type") == "write-denial", "Runtime owner write service call must route writes to denials")
assert(!write_service_call.fetch("read_only_dispatch"), "Runtime owner write service call must not mark writes as read dispatch")
assert(write_service_call.fetch("write_method"), "Runtime owner write service call must mark writes explicitly")
assert(!write_service_call.fetch("write_methods_enabled"), "Runtime owner write service call must keep write methods disabled")
assert(write_service_call.fetch("dispatch_ready"), "Runtime owner write service call must be dispatch-ready")
assert(write_service_call.fetch("error_name") == "org.xnix.Compatibility1.Error.WriteMethodDisabled", "Runtime owner write service call must use stable disabled-write errors")
assert(!write_service_call.fetch("event_loop_started"), "Runtime owner write service call must not start an event loop")
assert(!write_service_call.fetch("session_bus_claimed"), "Runtime owner write service call must not claim the session bus")
assert(!write_service_call.fetch("production_bus_claimed"), "Runtime owner write service call must not claim the production bus")
assert(!write_service_call.fetch("host_root_modified"), "Runtime owner write service call must not mutate the host root")
write_service_payload = write_service_call.fetch("payload")
assert(write_service_payload.fetch("method") == "Launch", "Runtime owner write service call must include the nested write denial method")
assert(!write_service_payload.fetch("dispatch_enabled"), "Runtime owner write service call must keep nested write dispatch disabled")
assert(!write_service_payload.fetch("request_created"), "Runtime owner write service call must not create nested write requests")

stdout, stderr, status = Open3.capture3(*owner_command, "--root", ".", "--mode", "smoke-owner", "--lifecycle-log")
assert(status.success?, "Runtime owner candidate smoke must render lifecycle JSONL: #{stderr}")
lifecycle_events = stdout.lines.map { |line| JSON.parse(line) }
assert(lifecycle_events.length == 4, "Runtime owner candidate smoke must emit four lifecycle events")
assert(lifecycle_events.map { |event| event.fetch("event_type") } == %w[startup route-table readiness shutdown], "Runtime owner candidate smoke must emit stable lifecycle event types")
lifecycle_events.each_with_index do |event, index|
  assert(event.fetch("schema_version") == "xnix.runtime.owner_lifecycle_event.v1", "Runtime owner candidate smoke must expose lifecycle event schema")
  assert(event.fetch("request_type") == "runtime-owner-lifecycle-event", "Runtime owner candidate smoke must expose lifecycle request type")
  assert(event.fetch("sequence") == index + 1, "Runtime owner candidate smoke must emit ordered lifecycle events")
  assert(event.fetch("mode") == "smoke-owner", "Runtime owner candidate smoke lifecycle must stay in smoke-owner mode")
  assert(event.fetch("route_count") == 61, "Runtime owner candidate smoke lifecycle must expose route count")
  assert(event.fetch("go_route_count") == 61, "Runtime owner candidate smoke lifecycle must expose Go route count")
  assert(event.fetch("write_method_count") == 4, "Runtime owner candidate smoke lifecycle must expose write method count")
  assert(event.fetch("read_only_serve_ready"), "Runtime owner candidate smoke lifecycle must show read-only serve readiness")
  assert(!event.fetch("write_methods_enabled"), "Runtime owner candidate smoke lifecycle must keep write methods disabled")
  assert(!event.fetch("event_loop_started"), "Runtime owner candidate smoke lifecycle must not start an event loop")
  assert(!event.fetch("session_bus_claimed"), "Runtime owner candidate smoke lifecycle must not claim the session bus")
  assert(!event.fetch("production_bus_claimed"), "Runtime owner candidate smoke lifecycle must not claim the production bus")
  assert(!event.fetch("system_service_started"), "Runtime owner candidate smoke lifecycle must not start a system service")
  assert(!event.fetch("network_required"), "Runtime owner candidate smoke lifecycle must not require network")
  assert(!event.fetch("host_root_modified"), "Runtime owner candidate smoke lifecycle must not mutate the host root")
  assert(!event.fetch("privileged_container_required"), "Runtime owner candidate smoke lifecycle must not require privileged containers")
  assert(!event.fetch("backend_details_exposed"), "Runtime owner candidate smoke lifecycle must not expose backend details")
end
assert(lifecycle_events.last.fetch("shutdown_reason") == "preview-complete", "Runtime owner candidate smoke lifecycle must expose a safe shutdown reason")

stdout, stderr, status = Open3.capture3(*owner_command, "--root", ".", "--smoke-batch")
assert(status.success?, "Runtime owner candidate smoke must render smoke batch JSONL: #{stderr}")
smoke_batch_records = stdout.lines.map { |line| JSON.parse(line) }
assert(smoke_batch_records.length == 66, "Runtime owner smoke batch must include 62 read records and 4 write records")
read_records = smoke_batch_records.select { |record| record.fetch("record_type") == "read-dispatch" }
write_records = smoke_batch_records.select { |record| record.fetch("record_type") == "write-denial" }
assert(read_records.length == 62, "Runtime owner smoke batch must cover every owner read dispatch method")
assert(write_records.length == 4, "Runtime owner smoke batch must cover every disabled write method")
smoke_batch_records.each_with_index do |record, index|
  assert(record.fetch("schema_version") == "xnix.runtime.owner_smoke_batch.v1", "Runtime owner smoke batch must expose its schema")
  assert(record.fetch("request_type") == "runtime-owner-smoke-batch-record", "Runtime owner smoke batch must expose its request type")
  assert(record.fetch("batch_type") == "restricted-session-owner-call-batch", "Runtime owner smoke batch must identify restricted session evidence")
  assert(record.fetch("sequence") == index + 1, "Runtime owner smoke batch must emit ordered records")
  assert(record.fetch("read_dispatch_method_count") == 62, "Runtime owner smoke batch must expose read dispatch coverage")
  assert(record.fetch("write_method_count") == 4, "Runtime owner smoke batch must expose write denial coverage")
  assert(record.fetch("runtime_owned"), "Runtime owner smoke batch must keep Runtime ownership in the Runtime")
  assert(record.fetch("go_runtime_backed"), "Runtime owner smoke batch must report Go backing")
  assert(!record.fetch("kde_policy_owner"), "Runtime owner smoke batch must not make KDE the policy owner")
  assert(!record.fetch("event_loop_started"), "Runtime owner smoke batch must not start an event loop")
  assert(!record.fetch("session_bus_claimed"), "Runtime owner smoke batch must not claim the session bus")
  assert(!record.fetch("production_bus_claimed"), "Runtime owner smoke batch must not claim the production bus")
  assert(!record.fetch("host_root_modified"), "Runtime owner smoke batch must not mutate the host root")
  assert(!record.fetch("backend_details_exposed"), "Runtime owner smoke batch must not expose backend details")
  assert(record.fetch("dispatch_ready"), "Runtime owner smoke batch records must be dispatch-ready")
  payload = record.fetch("payload")
  assert(payload.fetch("schema_version") == "xnix.runtime.owner_service_call.v1", "Runtime owner smoke batch payloads must use service-call schema")
  assert(payload.fetch("request_type") == "runtime-owner-service-call", "Runtime owner smoke batch payloads must use service-call request type")
  assert(!payload.fetch("session_bus_claimed"), "Runtime owner smoke batch service calls must not claim the session bus")
  assert(!payload.fetch("production_bus_claimed"), "Runtime owner smoke batch service calls must not claim the production bus")
end
assert(read_records.all? { |record| record.fetch("read_only_dispatch") && !record.fetch("write_method") }, "Runtime owner smoke batch read records must remain read-only")
assert(write_records.all? { |record| !record.fetch("read_only_dispatch") && record.fetch("write_method") }, "Runtime owner smoke batch write records must be explicit denials")
assert(write_records.all? { |record| record.fetch("error_name") == "org.xnix.Compatibility1.Error.WriteMethodDisabled" }, "Runtime owner smoke batch write records must use stable disabled-write errors")
assert(read_records.all? { |record| record.fetch("payload").fetch("call_type") == "read-dispatch" }, "Runtime owner smoke batch read records must use service-call read dispatch")
assert(write_records.all? { |record| record.fetch("payload").fetch("call_type") == "write-denial" }, "Runtime owner smoke batch write records must use service-call write denials")

stdout, stderr, status = Open3.capture3(*owner_command, "--root", ".", "--session-bus-smoke")
assert(status.success?, "Runtime owner candidate smoke must render private session-bus JSONL: #{stderr}")
session_bus_steps = stdout.lines.map { |line| JSON.parse(line) }
assert(session_bus_steps.length == 71, "Runtime owner private session-bus smoke must include startup, claim, route table, 62 reads, 4 writes, unsupported-read, and shutdown")
session_read_steps = session_bus_steps.select { |step| step.fetch("step_type") == "read-dispatch" }
session_write_steps = session_bus_steps.select { |step| step.fetch("step_type") == "write-denial" }
session_unsupported_steps = session_bus_steps.select { |step| step.fetch("step_type") == "reject-unsupported-read" }
assert(session_read_steps.length == 62, "Runtime owner private session-bus smoke must cover every owner read dispatch method")
assert(session_write_steps.length == 4, "Runtime owner private session-bus smoke must cover every disabled write method")
assert(session_unsupported_steps.length == 1, "Runtime owner private session-bus smoke must prove unsupported reads fail closed")
session_bus_steps.each_with_index do |step, index|
  assert(step.fetch("schema_version") == "xnix.runtime.owner_session_bus_smoke.v1", "Runtime owner private session-bus smoke must expose its schema")
  assert(step.fetch("request_type") == "runtime-owner-session-bus-smoke-step", "Runtime owner private session-bus smoke must expose its request type")
  assert(step.fetch("transcript_type") == "restricted-private-session-bus-owner-smoke", "Runtime owner private session-bus smoke must identify restricted evidence")
  assert(step.fetch("sequence") == index + 1, "Runtime owner private session-bus smoke must emit ordered steps")
  assert(step.fetch("read_dispatch_method_count") == 62, "Runtime owner private session-bus smoke must expose owner read dispatch coverage")
  assert(step.fetch("write_method_count") == 4, "Runtime owner private session-bus smoke must expose write denial coverage")
  assert(step.fetch("runtime_owned"), "Runtime owner private session-bus smoke must keep Runtime ownership in the Runtime")
  assert(step.fetch("go_runtime_backed"), "Runtime owner private session-bus smoke must report Go backing")
  assert(!step.fetch("kde_policy_owner"), "Runtime owner private session-bus smoke must not make KDE the policy owner")
  assert(step.fetch("private_session_bus"), "Runtime owner private session-bus smoke must be scoped to a private session bus")
  assert(step.fetch("event_loop_started"), "Runtime owner private session-bus smoke must model the restricted event loop")
  assert(step.fetch("session_bus_claimed"), "Runtime owner private session-bus smoke must claim only the private session bus")
  assert(!step.fetch("production_bus_claimed"), "Runtime owner private session-bus smoke must not claim production bus ownership")
  assert(!step.fetch("system_service_started"), "Runtime owner private session-bus smoke must not start a system service")
  assert(!step.fetch("write_methods_enabled"), "Runtime owner private session-bus smoke must keep write methods disabled")
  assert(!step.fetch("network_required"), "Runtime owner private session-bus smoke must not require network")
  assert(!step.fetch("host_root_modified"), "Runtime owner private session-bus smoke must not mutate the host root")
  assert(!step.fetch("privileged_container_required"), "Runtime owner private session-bus smoke must not require privileged containers")
  assert(!step.fetch("backend_details_exposed"), "Runtime owner private session-bus smoke must not expose backend details")
end
assert(session_read_steps.all? { |step| step.fetch("read_only_dispatch") && !step.fetch("write_method") && step.fetch("dispatch_ready") }, "Runtime owner private session-bus read steps must remain dispatch-ready reads")
assert(session_write_steps.all? { |step| !step.fetch("read_only_dispatch") && step.fetch("write_method") && step.fetch("error_name") == "org.xnix.Compatibility1.Error.WriteMethodDisabled" }, "Runtime owner private session-bus write steps must be stable denials")
assert(session_unsupported_steps.first.fetch("error_name") == "org.xnix.Compatibility1.Error.UnsupportedMethod", "Runtime owner private session-bus unsupported reads must use stable unsupported-method errors")

puts "PASS: Runtime owner candidate restricted session smoke"
