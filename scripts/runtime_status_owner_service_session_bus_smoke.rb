#!/usr/bin/env ruby
# frozen_string_literal: true

require "open3"
require "pathname"

PROJECT_ROOT = Pathname.new(__dir__).join("..").realpath
SMOKE_NAME = "Runtime-status owner service session-bus smoke"
INNER_PASS = "PASS: staged managed launcher dispatch smoke"
INNER_SKIP = "SKIP: staged managed launcher dispatch smoke"
PASS_MARKER = "PASS: Runtime-status owner service session-bus smoke"
SKIP_MARKER = "SKIP: Runtime-status owner service session-bus smoke"

def assert(condition, message)
  return if condition

  warn "FAIL: #{message}"
  exit 1
end

def assert_no_forbidden(text, label)
  downcased = text.downcase
  [
    "docker.sock",
    "--privileged",
    "--network host",
    "type=bind",
    "program files",
    ".wine",
    "qemu-system"
  ].each do |term|
    assert(!downcased.include?(term), "#{label} must not expose #{term}")
  end
end

def command_available?(name)
  ENV.fetch("PATH", "").split(File::PATH_SEPARATOR).any? do |directory|
    path = File.join(directory, name)
    File.file?(path) && File.executable?(path)
  end
end

unless ENV["DBUS_SESSION_BUS_ADDRESS"]
  unless command_available?("dbus-run-session")
    puts "#{SKIP_MARKER} (dbus-run-session unavailable)"
    exit 0
  end
  stdout, stderr, status = Open3.capture3("dbus-run-session", "--", "ruby", __FILE__, chdir: PROJECT_ROOT.to_s)
  print stdout
  warn stderr unless stderr.empty?
  exit status.exitstatus
end

stdout, stderr, status = Open3.capture3("ruby", "scripts/staged_launcher_dispatch_smoke.rb", chdir: PROJECT_ROOT.to_s)
print stdout
warn stderr unless stderr.empty?

assert(status.success?, "#{SMOKE_NAME} inner staged launcher dispatch smoke must pass or skip cleanly")
assert_no_forbidden(stdout, "#{SMOKE_NAME} stdout")
assert_no_forbidden(stderr, "#{SMOKE_NAME} stderr")

case stdout
when /#{Regexp.escape(INNER_PASS)}/
  puts PASS_MARKER
when /#{Regexp.escape(INNER_SKIP)}/
  reason = stdout.lines.find { |line| line.include?(INNER_SKIP) }.to_s.sub(INNER_SKIP, "").strip
  puts "#{SKIP_MARKER} #{reason}".rstrip
else
  warn stdout
  warn stderr unless stderr.empty?
  warn "FAIL: #{SMOKE_NAME} must report an inner PASS or SKIP marker"
  exit 1
end
