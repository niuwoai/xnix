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
        "([{'id': <'org.xnix.sample.notepad'>, 'name': <'Sample Notepad'>, 'icon': <'accessories-text-editor'>, 'mode': <'automatic'>, 'supported_extensions': <['.txt', '.log']>}],)\n",
        "",
        Status.new(true)
      ]
    when "org.xnix.Compatibility1.GetDiagnostics"
      [
        "({'application_id': <'org.xnix.sample.notepad'>, 'status': <'known'>, 'runtime_mode': <'automatic'>},)\n",
        "",
        Status.new(true)
      ]
    when "org.xnix.Compatibility1.GetEngineCatalog"
      [
        "({'catalog_type': <'compatibility-engine'>, 'runtime_policy_owner': <true>, 'backend_details_exposed': <false>},)\n",
        "",
        Status.new(true)
      ]
    when "org.xnix.Compatibility1.GetRunPlan"
      [
        "({'plan_type': <'compatibility-run'>, 'strategy': <'automatic-managed'>, 'backend_details_exposed': <false>},)\n",
        "",
        Status.new(true)
      ]
    when "org.xnix.Compatibility1.GetRepairPlan"
      [
        "({'plan_type': <'compatibility-repair'>, 'issue': <'engine-binding-pending'>, 'snapshot_required': <true>},)\n",
        "",
        Status.new(true)
      ]
    when "org.xnix.Compatibility1.GetSnapshotPlan"
      [
        "({'plan_type': <'compatibility-snapshot'>, 'reason': <'before-repair'>, 'enabled_by_default': <true>},)\n",
        "",
        Status.new(true)
      ]
    when "org.xnix.Compatibility1.GetPortalAccessPolicy"
      [
        "({'policy_type': <'portal-access'>, 'operation': <'file-open'>, 'portal_required': <true>},)\n",
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
assert(applications.first["supported_extensions"] == [".txt", ".log"], "D-Bus client must parse string array fields")

diagnostics = client.diagnostics("org.xnix.sample.notepad")
assert(diagnostics["status"] == "known", "D-Bus client must parse diagnostics status")

engine_catalog = client.engine_catalog
assert(engine_catalog["catalog_type"] == "compatibility-engine", "D-Bus client must parse engine catalog")
assert(engine_catalog["runtime_policy_owner"], "D-Bus client must parse boolean true values")
assert(!engine_catalog["backend_details_exposed"], "D-Bus client must parse boolean false values")

run_plan = client.run_plan("org.xnix.sample.notepad")
assert(run_plan["plan_type"] == "compatibility-run", "D-Bus client must parse run plans")
assert(!run_plan["backend_details_exposed"], "D-Bus client must parse run plan booleans")

repair_plan = client.repair_plan("org.xnix.sample.notepad", "engine-binding-pending")
assert(repair_plan["plan_type"] == "compatibility-repair", "D-Bus client must parse repair plans")
assert(repair_plan["snapshot_required"], "D-Bus client must parse repair plan booleans")

snapshot_plan = client.snapshot_plan("org.xnix.sample.notepad", "before-repair")
assert(snapshot_plan["plan_type"] == "compatibility-snapshot", "D-Bus client must parse snapshot plans")
assert(snapshot_plan["enabled_by_default"], "D-Bus client must parse snapshot plan booleans")

portal_policy = client.portal_access_policy("org.xnix.sample.notepad", "file-open")
assert(portal_policy["policy_type"] == "portal-access", "D-Bus client must parse Portal access policies")
assert(portal_policy["portal_required"], "D-Bus client must parse Portal policy booleans")

assert(
  capture.commands.all? { |command| command.include?("--session") },
  "D-Bus client must use the session bus for KDE-facing reads"
)

puts "PASS: compatibility runtime D-Bus client unit tests"
