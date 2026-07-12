#!/usr/bin/env ruby
# frozen_string_literal: true

require "open3"

BUS_NAME = "org.xnix.Compatibility1"
OBJECT_PATH = "/org/xnix/Compatibility1"
INTERFACE = "org.xnix.Compatibility1"

def assert(condition, message)
  return if condition

  warn "FAIL: #{message}"
  exit 1
end

unless ENV["DBUS_SESSION_BUS_ADDRESS"]
  stdout, stderr, status = Open3.capture3("dbus-run-session", "--", "ruby", __FILE__)
  print stdout
  warn stderr unless stderr.empty?
  exit status.exitstatus
end

server_log = "/tmp/xnix-dbus-smoke.log"
server_pid = spawn("xnix-dbus-smoke", out: server_log, err: [:child, :out])

begin
  _stdout, stderr, status = Open3.capture3("gdbus", "wait", "--session", "--timeout", "5", BUS_NAME)
  assert(status.success?, "runtime smoke adapter must own #{BUS_NAME}: #{stderr}")

  stdout, stderr, status = Open3.capture3(
    "gdbus", "introspect",
    "--session",
    "--dest", BUS_NAME,
    "--object-path", OBJECT_PATH
  )
  assert(status.success?, "runtime smoke adapter must be introspectable: #{stderr}")
  assert(stdout.include?(INTERFACE), "runtime smoke adapter introspection must expose #{INTERFACE}")
  assert(stdout.include?("ListApplications"), "runtime smoke adapter introspection must expose ListApplications")

  stdout, stderr, status = Open3.capture3(
    "gdbus", "call",
    "--session",
    "--dest", BUS_NAME,
    "--object-path", OBJECT_PATH,
    "--method", "#{INTERFACE}.ListApplications"
  )
  assert(status.success?, "runtime smoke adapter must answer ListApplications: #{stderr}")
  assert(stdout.include?("org.xnix.sample.notepad"), "runtime smoke adapter must expose the sample recipe over D-Bus")

  stdout, stderr, status = Open3.capture3(
    "gdbus", "call",
    "--session",
    "--dest", BUS_NAME,
    "--object-path", OBJECT_PATH,
    "--method", "#{INTERFACE}.GetDiagnostics",
    "org.xnix.sample.notepad"
  )
  assert(status.success?, "runtime smoke adapter must answer GetDiagnostics: #{stderr}")
  assert(stdout.include?("known"), "runtime smoke adapter diagnostics must identify known applications")

  puts "PASS: compatibility runtime D-Bus session smoke"
ensure
  begin
    Process.kill("TERM", server_pid)
    Process.wait(server_pid)
  rescue Errno::ESRCH, Errno::ECHILD
    nil
  end
end
