#!/usr/bin/env ruby
# frozen_string_literal: true

require_relative "../lib/xnix/full_smoke_report"
require_relative "../lib/xnix/serial_log"

def assert(condition, message)
  return if condition

  warn "FAIL: #{message}"
  exit 1
end

serial_log = Xnix::SerialLog.new("Welcome to Buildroot\nxnix login: ")
report = Xnix::FullSmokeReport.new(
  version: "0.2.300",
  steps: %w[build fetch-sources configure-system download-system build-system boot-system],
  serial_log_path: "output/serial.log",
  serial_log: serial_log
)

data = report.to_h
assert(data.fetch("version") == "0.2.300", "full smoke report must expose the version")
assert(data.fetch("schema_version") == "xnix.full_smoke_report.v1", "full smoke report must expose its schema")
assert(data.fetch("report_type") == "full-build-qemu-smoke-report", "full smoke report must identify its type")
assert(data.fetch("step_count") == 6, "full smoke report must count smoke steps")
assert(data.fetch("steps").last == "boot-system", "full smoke report must include the QEMU boot step")
assert(data.fetch("serial_log_path") == "output/serial.log", "full smoke report must expose the serial log path")
assert(data.fetch("serial_log_persisted"), "full smoke report must record serial log persistence")
assert(data.fetch("qemu_booted"), "full smoke report must record successful QEMU boot markers")
assert(data.fetch("missing_boot_markers").empty?, "full smoke report must expose missing boot markers")
assert(!data.fetch("host_root_modified"), "full smoke report must not claim host root mutation")
assert(!data.fetch("privileged_container_required"), "full smoke report must not require privileged containers")
assert(!data.fetch("docker_socket_mounted"), "full smoke report must not mount the Docker socket")
assert(!data.fetch("host_network_enabled"), "full smoke report must keep host networking disabled")
assert(data.fetch("qemu_network_restricted"), "full smoke report must keep QEMU networking restricted")
assert(data.fetch("network_required_for_source_download"), "full smoke report must identify source download network use")

markdown = report.to_markdown
assert(markdown.include?("# Full Build and QEMU Smoke Report"), "full smoke report markdown must include a title")
assert(markdown.include?("- QEMU booted: true"), "full smoke report markdown must include boot status")
assert(markdown.include?("6. boot-system"), "full smoke report markdown must list the boot step")

puts "PASS: full smoke report unit tests"
