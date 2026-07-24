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

def run_container_step(arguments)
  captured_lines = []
  status = nil
  Open3.popen2e("ruby", "scripts/container.rb", *arguments) do |_stdin, output, wait_thread|
    output.each_line do |line|
      print line
      captured_lines << line
      captured_lines.shift if captured_lines.length > 400
    end
    status = wait_thread.value
  end
  [captured_lines.join, status]
end

def write_report(completed_steps:, serial_output:, failure: nil)
  FileUtils.mkdir_p(LOG_PATH.dirname)
  LOG_PATH.write(serial_output)
  report = Xnix::FullSmokeReport.new(
    version: VERSION,
    steps: completed_steps,
    serial_log_path: LOG_PATH.relative_path_from(PROJECT_ROOT).to_s,
    serial_log: Xnix::SerialLog.new(serial_output),
    failure: failure
  )
  REPORT_JSON_PATH.write(JSON.pretty_generate(report.to_h))
  REPORT_MARKDOWN_PATH.write(report.to_markdown)
end

def fail_full_smoke(step:, completed_steps:, error_text:, message:)
  failure = Xnix::FullSmokeReport.classify_failure(step: step, error_text: error_text)
  write_report(completed_steps: completed_steps + [step], serial_output: error_text, failure: failure)
  abort message
end

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

completed_steps = []
steps.each do |arguments|
  output, status = run_container_step(arguments)
  fail_full_smoke(
    step: arguments.first,
    completed_steps: completed_steps,
    error_text: output,
    message: "Full smoke step failed: #{arguments.first}"
  ) unless status.success?
  completed_steps << arguments.first
end

def require_pass_step(arguments, expected_pass, completed_steps)
  output, status = run_container_step(arguments)
  return output if status.success? && output.include?(expected_pass)

  message = status.success? ? "Full smoke step did not prove PASS: #{arguments.first}" : "Full smoke step failed: #{arguments.first}"
  fail_full_smoke(
    step: arguments.first,
    completed_steps: completed_steps,
    error_text: output,
    message: message
  )
end

serial_output, status = Open3.capture2e("ruby", "scripts/container.rb", "boot-system")
completed_steps << "boot-system"
fail_full_smoke(
  step: "boot-system",
  completed_steps: completed_steps[0...-1],
  error_text: serial_output,
  message: "QEMU boot process failed"
) unless status.success? || status.exitstatus == 124

serial_log = Xnix::SerialLog.new(serial_output)
fail_full_smoke(
  step: "boot-system",
  completed_steps: completed_steps[0...-1],
  error_text: "Boot markers missing: #{serial_log.missing_markers.join(", ")}\n#{serial_output}",
  message: "Boot markers missing: #{serial_log.missing_markers.join(", ")}"
) unless serial_log.booted?

real_app_steps = [
  [["staged-launcher-dispatch-smoke"], "PASS: staged managed launcher dispatch smoke"],
  [["winapp-guest-wine-smoke"], "PASS: QEMU guest real Windows app Wine smoke"],
  [["kde-controlled-launch-action-dbus-fixture-smoke"], "PASS: KDE controlled launch action smoke"]
]
real_app_steps.each do |arguments, expected_pass|
  require_pass_step(arguments, expected_pass, completed_steps)
  completed_steps << arguments.first
end

write_report(completed_steps: completed_steps, serial_output: serial_output)

puts "PASS: full build and QEMU serial smoke test"
