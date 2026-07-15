#!/usr/bin/env ruby
# frozen_string_literal: true

require "fileutils"
require "json"
require "open3"
require "pathname"
require_relative "../lib/xnix/full_smoke_report"
require_relative "../lib/xnix/milestone"
require_relative "../lib/xnix/serial_log"

PROJECT_ROOT = Pathname.new(__dir__).join("..").realpath
VERSION = PROJECT_ROOT.join("VERSION").read.strip
LOG_PATH = PROJECT_ROOT.join("output", "serial.log")
REPORT_JSON_PATH = PROJECT_ROOT.join("output", "full-smoke-report.json")
REPORT_MARKDOWN_PATH = PROJECT_ROOT.join("output", "full-smoke-report.md")

abort "Full smoke tests require a tenth formal version" unless Xnix::Milestone.full_build_required?(VERSION)

steps = [
  ["build"],
  ["fetch-sources"],
  ["configure-system"],
  ["download-system"],
  ["build-system"]
]

steps.each do |arguments|
  success = system("ruby", "scripts/container.rb", *arguments)
  abort "Full smoke step failed: #{arguments.first}" unless success
end

serial_output, status = Open3.capture2e("ruby", "scripts/container.rb", "boot-system")
FileUtils.mkdir_p(LOG_PATH.dirname)
LOG_PATH.write(serial_output)
abort "QEMU boot process failed" unless status.success? || status.exitstatus == 124

serial_log = Xnix::SerialLog.new(serial_output)
report = Xnix::FullSmokeReport.new(
  version: VERSION,
  steps: steps.flatten + ["boot-system"],
  serial_log_path: LOG_PATH.relative_path_from(PROJECT_ROOT).to_s,
  serial_log: serial_log
)
REPORT_JSON_PATH.write(JSON.pretty_generate(report.to_h))
REPORT_MARKDOWN_PATH.write(report.to_markdown)
abort "Boot markers missing: #{serial_log.missing_markers.join(", ")}" unless serial_log.booted?

puts "PASS: full build and QEMU serial smoke test"
