#!/usr/bin/env ruby
# frozen_string_literal: true

require_relative "../lib/xnix/serial_log"

def assert(condition, message)
  return if condition

  warn "FAIL: #{message}"
  exit 1
end

successful_log = "Linux version 6.x\nWelcome to Buildroot\nxnix login: "
successful_boot = Xnix::SerialLog.new(successful_log)
assert(successful_boot.booted?, "complete serial log must prove the system booted")
assert(successful_boot.missing_markers.empty?, "complete serial log must have no missing boot markers")
assert(successful_boot.summary == "boot completed", "complete serial log must have a success summary")

incomplete_boot = Xnix::SerialLog.new("Linux version 6.x\nWelcome to Buildroot\n")
assert(!incomplete_boot.booted?, "missing login prompt must fail boot verification")
assert(incomplete_boot.missing_markers == ["xnix login:"], "missing marker report must identify the login prompt")
assert(incomplete_boot.summary == "missing boot markers: xnix login:", "incomplete log summary must identify the missing marker")

puts "PASS: serial log unit tests"
