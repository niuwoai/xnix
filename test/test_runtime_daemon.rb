#!/usr/bin/env ruby
# frozen_string_literal: true

require "json"
require "open3"
require "pathname"

def assert(condition, message)
  return if condition

  warn "FAIL: #{message}"
  exit 1
end

project_root = Pathname.new(__dir__).join("..").realpath
command = ["ruby", project_root.join("bin/xnix-compatd").to_s]

stdout, stderr, status = Open3.capture3(*command, "probe")
assert(status.success?, "runtime daemon probe must exit successfully: #{stderr}")
probe = JSON.parse(stdout)
assert(probe["version"] == "0.2.32", "runtime daemon probe must report the current version")
assert(probe["bus_name"] == "org.xnix.Compatibility1", "runtime daemon probe must keep the stable bus name")
assert(probe["capabilities"]["recipe_store"], "runtime daemon probe must expose recipe store capability")
assert(probe["capabilities"]["registry_backed_recipe_store"], "runtime daemon probe must expose registry-backed recipe loading")
assert(probe["capabilities"]["compatibility_engine_catalog"], "runtime daemon probe must expose engine catalog capability")
assert(probe["capabilities"]["compatibility_run_planning"], "runtime daemon probe must expose run planning capability")
assert(probe["capabilities"]["compatibility_repair_planning"], "runtime daemon probe must expose repair planning capability")
assert(probe["capabilities"]["dbus_method_dispatch"], "runtime daemon probe must expose method dispatch capability")
assert(!probe["capabilities"]["dbus_binding"], "runtime daemon must not claim a D-Bus binding before it exists")
assert(probe["recipe_trust"]["registry_backed"], "runtime daemon probe must report registry-backed recipe loading")
assert(probe["recipe_trust"]["digest_verified"], "runtime daemon probe must report digest-verified recipes")
assert(!probe["recipe_trust"]["signed_recipe_validation"], "runtime daemon probe must not claim production recipe signatures yet")

stdout, stderr, status = Open3.capture3(*command, "list-applications")
assert(status.success?, "runtime daemon list must exit successfully: #{stderr}")
applications = JSON.parse(stdout)
assert(applications.length == 1, "runtime daemon must list the bundled sample application")
assert(applications.first["id"] == "org.xnix.sample.notepad", "runtime daemon must expose sample recipe id")
assert(applications.first["mime_types"].include?("application/x-xnix-txt"), "runtime daemon must expose MIME types")

stdout, stderr, status = Open3.capture3(*command, "diagnostics", "org.xnix.sample.notepad")
assert(status.success?, "runtime daemon diagnostics must exit successfully: #{stderr}")
diagnostics = JSON.parse(stdout)
assert(diagnostics["status"] == "known", "runtime daemon diagnostics must know bundled recipes")
assert(diagnostics["checks"].any? { |check| check["status"] == "pending" }, "runtime daemon must report backend work as pending")
assert(diagnostics["repair_plan"]["plan_type"] == "compatibility-repair", "runtime daemon diagnostics must include repair planning")
assert(diagnostics["repair_plan"]["issue"] == "engine-binding-pending", "runtime daemon diagnostics must identify the pending repair issue")
assert(diagnostics["repair_plan"]["snapshot_required"], "runtime daemon diagnostics repair plan must require snapshots")
assert(diagnostics["repair_plan"]["snapshot_plan"]["plan_type"] == "compatibility-snapshot", "runtime daemon diagnostics must include snapshot planning")
assert(diagnostics["repair_plan"]["snapshot_plan"]["restore_available"], "runtime daemon diagnostics snapshot plan must expose restore availability")

diagnostics_json = JSON.pretty_generate(diagnostics)
assert(!diagnostics_json.match?(/wine|prefix|\.wine|proton|virtual machine/i), "runtime diagnostics must not expose backend implementation terms")

_stdout, stderr, status = Open3.capture3(*command, "diagnostics", "org.xnix.missing")
assert(!status.success?, "runtime daemon must reject unknown applications")
assert(stderr.include?("unknown application"), "runtime daemon must explain unknown applications")

puts "PASS: compatibility runtime daemon unit tests"
