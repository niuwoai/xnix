#!/usr/bin/env ruby
# frozen_string_literal: true

require "json"
require "open3"
require "pathname"
require_relative "../lib/xnix/compatibility/runtime_write_gate"

def assert(condition, message)
  return if condition

  warn "FAIL: #{message}"
  exit 1
end

project_root = Pathname.new(__dir__).join("..").realpath

Xnix::Compatibility::RuntimeWriteGate::WRITE_METHODS.each do |method_name|
  gate = Xnix::Compatibility::RuntimeWriteGate.new(method_name: method_name)
  model = gate.to_h

  assert(model["version"] == "0.2.90", "Runtime write gate must expose the current version")
  assert(model["gate_type"] == "runtime-write-gate", "Runtime write gate must identify the gate type")
  assert(model["method_name"] == method_name, "Runtime write gate must preserve the method name")
  assert(model["runtime_owned"], "Runtime must own write gates")
  assert(!model["kde_policy_owner"], "KDE must not own write gate policy")
  assert(model["gate_decision"] == "blocked-until-production-backend", "Runtime write gate must block production writes")
  assert(!model["write_method_enabled"], "Runtime write gate must not enable write methods")
  assert(!model["dispatch_enabled"], "Runtime write gate must not enable dispatch")
  assert(!model["request_object_created"], "Runtime write gate must not create request objects")
  assert(!model["execution_started"], "Runtime write gate must not start execution")
  assert(model["required_gates"].length == 6, "Runtime write gate must list required production gates")
  assert(model["required_gates"].all? { |item| item["status"] == "pending" }, "Runtime write gate must keep required gates pending")
  assert(model["denial_error_name"] == Xnix::Compatibility::RuntimeWriteGate::ERROR_NAME, "Runtime write gate must expose the D-Bus denial error")
  assert(!model["network_required"], "Runtime write gate must not require network access")
  assert(!model["host_root_modified"], "Runtime write gate must not mutate the host root")
  assert(!model["privileged_container_required"], "Runtime write gate must not require privileged containers")
  assert(!model["backend_details_exposed"], "Runtime write gate must hide backend details")
  assert(gate.failure_message.include?(Xnix::Compatibility::RuntimeWriteGate::ERROR_NAME), "Runtime write gate failure message must include the D-Bus error name")

  json = JSON.pretty_generate(model)
  assert(!json.match?(/wine|prefix|\.wine|proton|virtual machine/i), "Runtime write gate must not expose backend implementation terms")
end

stdout, stderr, status = Open3.capture3("ruby", project_root.join("bin/xnix-runtime-write-gate").to_s, "--method", "Launch")
assert(status.success?, "Runtime write gate CLI must exit successfully: #{stderr}")
assert(JSON.parse(stdout)["method_name"] == "Launch", "Runtime write gate CLI must emit the selected method")

_stdout, stderr, status = Open3.capture3("ruby", project_root.join("bin/xnix-runtime-write-gate").to_s, "--method", "Unknown")
assert(!status.success?, "Runtime write gate CLI must reject unknown methods")
assert(stderr.include?("invalid argument"), "Runtime write gate CLI must explain invalid methods")

puts "PASS: Runtime write gate unit tests"
