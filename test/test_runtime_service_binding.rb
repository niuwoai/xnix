#!/usr/bin/env ruby
# frozen_string_literal: true

require "json"
require "open3"
require "pathname"
require_relative "../lib/xnix/compatibility/runtime_service_binding"

def assert(condition, message)
  return if condition

  warn "FAIL: #{message}"
  exit 1
end

project_root = Pathname.new(__dir__).join("..").realpath
binding = Xnix::Compatibility::RuntimeServiceBinding.new.to_h
check_ids = binding.fetch("checks").map { |item| item.fetch("id") }

assert(binding["version"] == "0.2.68", "Runtime service binding must expose the current version")
assert(binding["binding_type"] == "runtime-service-binding", "Runtime service binding must identify the binding type")
assert(binding["runtime_owned"], "Runtime must own service binding status")
assert(!binding["kde_policy_owner"], "KDE must not own service binding policy")
assert(binding["bus_name"] == "org.xnix.Compatibility1", "Runtime service binding must expose the stable bus name")
assert(binding["object_path"] == "/org/xnix/Compatibility1", "Runtime service binding must expose the stable object path")
assert(binding["interface"] == "org.xnix.Compatibility1", "Runtime service binding must expose the stable interface")
assert(binding["activation"]["dbus_service_file"] == "runtime/dbus/org.xnix.Compatibility1.service", "Runtime service binding must expose the D-Bus service file")
assert(binding["activation"]["systemd_unit"] == "runtime/systemd/xnix-compatd.service", "Runtime service binding must expose the systemd unit")
assert(binding["activation"]["libexec_wrapper"] == "libexec/xnix/compatd", "Runtime service binding must expose the libexec wrapper")
assert(binding["activation"]["packaged_wrapper"] == "/usr/libexec/xnix/compatd", "Runtime service binding must expose the packaged wrapper path")
assert(check_ids == %w[dbus-service-activation systemd-service-hardening libexec-wrapper dbus-contract live-dbus-owner], "Runtime service binding must report expected checks")
assert(binding["counts"] == { "total" => 5, "passed" => 4, "pending" => 1, "blocked" => 0 }, "Runtime service binding must count activation checks")
assert(binding["activation_binding_ready"], "Runtime service binding activation files must be aligned")
assert(!binding["live_dbus_owner_ready"], "Runtime service binding must not claim live bus ownership yet")
assert(binding["smoke_adapter_available"], "Runtime service binding must report the smoke adapter")
assert(!binding["network_required"], "Runtime service binding must not require network access")
assert(!binding["host_root_modified"], "Runtime service binding must not mutate the host root")
assert(!binding["privileged_container_required"], "Runtime service binding must not require privileged containers")
assert(!binding["backend_details_exposed"], "Runtime service binding must hide backend details")

json = JSON.pretty_generate(binding)
assert(!json.match?(/wine|prefix|\.wine|proton|virtual machine/i), "Runtime service binding must not expose backend implementation terms")

stdout, stderr, status = Open3.capture3("ruby", project_root.join("bin/xnix-runtime-service-binding").to_s)
assert(status.success?, "Runtime service binding CLI must exit successfully: #{stderr}")
assert(JSON.parse(stdout) == binding, "Runtime service binding CLI must emit the binding model")

puts "PASS: Runtime service binding unit tests"
