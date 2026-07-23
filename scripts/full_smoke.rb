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

abort "Full smoke tests require a twentieth formal version" unless Xnix::Milestone.full_build_required?(VERSION)

steps = [
  ["build-tools"],
  ["build"],
  ["fetch-sources"],
  ["configure-system"],
  ["download-system"],
  ["build-system"],
  ["prepare-ssh-test-key"],
  ["configure-wine-guest"],
  ["download-wine-guest"],
  ["build-ssh-wine-guest"],
  ["fetch-known-winapp"]
]

steps.each do |arguments|
  success = system("ruby", "scripts/container.rb", *arguments)
  abort "Full smoke step failed: #{arguments.first}" unless success
end

def require_pass_step(arguments, expected_pass)
  output, status = Open3.capture2e("ruby", "scripts/container.rb", *arguments)
  puts output
  abort "Full smoke step failed: #{arguments.first}" unless status.success?
  abort "Full smoke step did not prove PASS: #{arguments.first}" unless output.include?(expected_pass)
end

serial_output, status = Open3.capture2e("ruby", "scripts/container.rb", "boot-system")
FileUtils.mkdir_p(LOG_PATH.dirname)
LOG_PATH.write(serial_output)
abort "QEMU boot process failed" unless status.success? || status.exitstatus == 124

real_app_steps = [
  [["known-winapp-guest-wine-smoke"], "PASS: known Windows app QEMU guest Wine smoke"],
  [["winapp-guest-wine-smoke"], "PASS: QEMU guest real Windows app Wine smoke"]
]
real_app_steps.each do |arguments, expected_pass|
  require_pass_step(arguments, expected_pass)
end

serial_log = Xnix::SerialLog.new(serial_output)
report = Xnix::FullSmokeReport.new(
  version: VERSION,
  steps: steps.flatten + ["boot-system"] + real_app_steps.map { |arguments, _expected_pass| arguments.first },
  serial_log_path: LOG_PATH.relative_path_from(PROJECT_ROOT).to_s,
  serial_log: serial_log
)
REPORT_JSON_PATH.write(JSON.pretty_generate(report.to_h))
REPORT_MARKDOWN_PATH.write(report.to_markdown)
abort "Boot markers missing: #{serial_log.missing_markers.join(", ")}" unless serial_log.booted?

puts "PASS: full build and QEMU serial smoke test"
