#!/usr/bin/env ruby
# frozen_string_literal: true

require "json"
require "open3"
require "pathname"
require_relative "../lib/xnix/compatibility/compatibility_engine_catalog"

def assert(condition, message)
  return if condition

  warn "FAIL: #{message}"
  exit 1
end

project_root = Pathname.new(__dir__).join("..").realpath
catalog = Xnix::Compatibility::CompatibilityEngineCatalog.new
model = catalog.to_h

assert(model["version"] == "0.2.78", "compatibility engine catalog must expose the current version")
assert(model["catalog_type"] == "compatibility-engine", "compatibility engine catalog must identify the model type")
assert(model["runtime_policy_owner"], "Runtime must own engine selection policy")
assert(!model["desktop_shell_policy_owner"], "KDE must not own engine selection policy")
assert(!model["backend_details_exposed"], "compatibility engine catalog must hide backend details")
assert(model["engines"].length == 3, "compatibility engine catalog must expose three engine choices")
assert(model["engines"].map { |engine| engine.fetch("id") } == %w[automatic-managed local-compatibility-engine isolated-compatibility-engine], "compatibility engine catalog must keep engine order stable")
assert(model["engines"].all? { |engine| !engine.fetch("ready") }, "compatibility engine catalog must not claim engine readiness yet")
assert(model["engines"].all? { |engine| !engine.fetch("launch_enabled") }, "compatibility engine catalog must not claim launch enablement yet")

automatic = catalog.select_for_mode("automatic")
local = catalog.select_for_mode("wine")
isolated = catalog.select_for_mode("vm")
assert(automatic["engine_id"] == "automatic-managed", "automatic mode must select automatic engine")
assert(local["engine_id"] == "local-compatibility-engine", "local mode must select local engine")
assert(isolated["engine_id"] == "isolated-compatibility-engine", "isolated mode must select isolated engine")
assert(!local["backend_details_exposed"], "selected engines must hide backend details")

json = JSON.pretty_generate(model)
assert(!json.match?(/wine|prefix|\.wine|proton|virtual machine/i), "compatibility engine catalog must not expose backend implementation terms")

stdout, stderr, status = Open3.capture3("ruby", project_root.join("bin/xnix-compat-engine-catalog").to_s)
assert(status.success?, "compatibility engine catalog CLI must exit successfully: #{stderr}")
assert(JSON.parse(stdout) == model, "compatibility engine catalog CLI must emit the catalog")

stdout, stderr, status = Open3.capture3(
  "ruby",
  project_root.join("bin/xnix-compat-engine-catalog").to_s,
  "--mode",
  "vm"
)
assert(status.success?, "compatibility engine catalog CLI must select by mode: #{stderr}")
assert(JSON.parse(stdout) == isolated, "compatibility engine catalog CLI must emit selected engine")

puts "PASS: compatibility engine catalog unit tests"
