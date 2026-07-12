#!/usr/bin/env ruby
# frozen_string_literal: true

require "pathname"
require_relative "../lib/xnix/compatibility/dbus_runtime_client"

def assert(condition, message)
  return if condition

  warn "FAIL: #{message}"
  exit 1
end

Status = Struct.new(:success?)

class FakeCapture
  attr_reader :commands

  def initialize
    @commands = []
  end

  def call(*command)
    @commands << command
    method = command.fetch(command.index("--method") + 1)

    case method
    when "org.xnix.Compatibility1.ListApplications"
      [
        "([{'id': <'org.xnix.sample.notepad'>, 'name': <'Sample Notepad'>, 'icon': <'accessories-text-editor'>, 'mode': <'automatic'>}],)\n",
        "",
        Status.new(true)
      ]
    when "org.xnix.Compatibility1.GetDiagnostics"
      [
        "({'application_id': <'org.xnix.sample.notepad'>, 'status': <'known'>, 'runtime_mode': <'automatic'>},)\n",
        "",
        Status.new(true)
      ]
    else
      ["", "unexpected method", Status.new(false)]
    end
  end
end

capture = FakeCapture.new
client = Xnix::Compatibility::DBusRuntimeClient.new(capture: capture)

source = client.source_metadata
assert(source["kind"] == "runtime-dbus-session", "D-Bus client must declare a session bus source")
assert(source["bus_name"] == "org.xnix.Compatibility1", "D-Bus client must use the Runtime bus name")

applications = client.list_applications
assert(applications.length == 1, "D-Bus client must parse the application list")
assert(applications.first["id"] == "org.xnix.sample.notepad", "D-Bus client must parse string fields")
assert(applications.first["mode"] == "automatic", "D-Bus client must parse application mode")

diagnostics = client.diagnostics("org.xnix.sample.notepad")
assert(diagnostics["status"] == "known", "D-Bus client must parse diagnostics status")
assert(
  capture.commands.all? { |command| command.include?("--session") },
  "D-Bus client must use the session bus for KDE-facing reads"
)

puts "PASS: compatibility runtime D-Bus client unit tests"
