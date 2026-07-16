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
environment = {"GOCACHE" => ENV.fetch("GOCACHE", "/tmp/xnix-go-cache")}

json_output, json_error, json_status = Open3.capture3(
  environment,
  "ruby", "scripts/restricted_product_smoke_packet.rb", "--format", "json",
  chdir: project_root.to_s
)
assert(json_status.success?, "restricted product smoke JSON must run: #{json_error}")
packet = JSON.parse(json_output)
assert(packet.fetch("schema_version") == "xnix.runtime.restricted_product_smoke_packet.v1", "packet schema must be stable")
assert(packet.fetch("packet_prepared"), "packet must be prepared")
assert(packet.fetch("ready_for_authorized_smoke"), "packet must be ready for an authorized smoke")
assert(packet.fetch("human_authorization_required"), "packet must require human authorization")
assert(packet.fetch("evidence_count") == 5, "packet must cover five evidence groups")
%w[docker_executed qemu_executed product_smoke_executed serial_log_persisted docker_socket_mounted host_network_enabled broad_host_mount_enabled privileged_container_required backend_launch_enabled host_root_modified release_ready].each do |key|
  assert(!packet.fetch(key), "packet must keep #{key} disabled")
end
assert(packet.fetch("loopback_only_networking"), "packet must require loopback-only networking")
assert(packet.fetch("serial_log_persistence_required"), "packet must require persisted serial logs")

markdown_output, markdown_error, markdown_status = Open3.capture3(
  environment,
  "ruby", "scripts/restricted_product_smoke_packet.rb", "--format", "markdown",
  chdir: project_root.to_s
)
assert(markdown_status.success?, "restricted product smoke Markdown must run: #{markdown_error}")
assert(markdown_output.include?("# Restricted Product Smoke Packet"), "Markdown must include a title")
assert(markdown_output.include?("runtime-owner: ready"), "Markdown must include Runtime owner evidence")
assert(markdown_output.include?("QEMU executed: false"), "Markdown must expose skipped QEMU execution")

script = project_root.join("scripts/restricted_product_smoke_packet.rb").read
assert(!script.include?("scripts/container.rb"), "restricted packet must not run the container harness")
assert(!script.include?("boot-system"), "restricted packet must not boot QEMU")

puts "PASS: restricted product smoke packet tests"
