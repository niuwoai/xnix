#!/usr/bin/env ruby
# frozen_string_literal: true

require "json"
require "open3"

BUS_NAME = "org.xnix.Compatibility1"

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

server_log = "/tmp/xnix-kde-center-dbus-smoke.log"
server_pid = spawn("xnix-dbus-smoke", out: server_log, err: [:child, :out])

begin
  _stdout, stderr, status = Open3.capture3("gdbus", "wait", "--session", "--timeout", "5", BUS_NAME)
  assert(status.success?, "runtime smoke adapter must own #{BUS_NAME}: #{stderr}")

  stdout, stderr, status = Open3.capture3("ruby", "bin/xnix-kde-center-model", "--source", "dbus")
  assert(status.success?, "KDE center model must read from D-Bus: #{stderr}")

  model = JSON.parse(stdout)
  assert(model["source"]["kind"] == "runtime-dbus-session", "KDE center model must report its D-Bus source")
  assert(model["summary"]["application_count"] == 1, "KDE center model must summarize D-Bus applications")
  assert(model["applications"].first["id"] == "org.xnix.sample.notepad", "KDE center model must expose D-Bus applications")
  assert(model["applications"].first["compatibility_label"] == "Known", "KDE center model must expose D-Bus diagnostics")

  puts "PASS: KDE compatibility center D-Bus smoke"
ensure
  begin
    Process.kill("TERM", server_pid)
    Process.wait(server_pid)
  rescue Errno::ESRCH, Errno::ECHILD
    nil
  end
end
