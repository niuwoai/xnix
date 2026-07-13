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

stdout, stderr, status = Open3.capture3(*command, "dispatch", "ListApplications")
assert(status.success?, "dispatch ListApplications must exit successfully: #{stderr}")
applications = JSON.parse(stdout)
assert(applications.first["id"] == "org.xnix.sample.notepad", "dispatch must route ListApplications")

stdout, stderr, status = Open3.capture3(
  *command,
  "dispatch",
  "GetApplication",
  JSON.generate(["org.xnix.sample.notepad"])
)
assert(status.success?, "dispatch GetApplication must exit successfully: #{stderr}")
application = JSON.parse(stdout)
assert(application["name"] == "Sample Notepad", "dispatch must route GetApplication")

stdout, stderr, status = Open3.capture3(
  *command,
  "dispatch",
  "GetDiagnostics",
  JSON.generate(["org.xnix.sample.notepad"])
)
assert(status.success?, "dispatch GetDiagnostics must exit successfully: #{stderr}")
diagnostics = JSON.parse(stdout)
assert(diagnostics["application_id"] == "org.xnix.sample.notepad", "dispatch must route GetDiagnostics")

stdout, stderr, status = Open3.capture3(*command, "dispatch", "GetEngineCatalog")
assert(status.success?, "dispatch GetEngineCatalog must exit successfully: #{stderr}")
catalog = JSON.parse(stdout)
assert(catalog["catalog_type"] == "compatibility-engine", "dispatch must route GetEngineCatalog")

stdout, stderr, status = Open3.capture3(
  *command,
  "dispatch",
  "GetRunPlan",
  JSON.generate(["org.xnix.sample.notepad"])
)
assert(status.success?, "dispatch GetRunPlan must exit successfully: #{stderr}")
run_plan = JSON.parse(stdout)
assert(run_plan["plan_type"] == "compatibility-run", "dispatch must route GetRunPlan")

stdout, stderr, status = Open3.capture3(
  *command,
  "dispatch",
  "GetRepairPlan",
  JSON.generate(["org.xnix.sample.notepad", "engine-binding-pending"])
)
assert(status.success?, "dispatch GetRepairPlan must exit successfully: #{stderr}")
repair_plan = JSON.parse(stdout)
assert(repair_plan["plan_type"] == "compatibility-repair", "dispatch must route GetRepairPlan")

stdout, stderr, status = Open3.capture3(
  *command,
  "dispatch",
  "GetTestPlan",
  JSON.generate(["org.xnix.sample.notepad"])
)
assert(status.success?, "dispatch GetTestPlan must exit successfully: #{stderr}")
test_plan = JSON.parse(stdout)
assert(test_plan["plan_type"] == "compatibility-test", "dispatch must route GetTestPlan")

stdout, stderr, status = Open3.capture3(
  *command,
  "dispatch",
  "GetSnapshotPlan",
  JSON.generate(["org.xnix.sample.notepad", "before-repair"])
)
assert(status.success?, "dispatch GetSnapshotPlan must exit successfully: #{stderr}")
snapshot_plan = JSON.parse(stdout)
assert(snapshot_plan["plan_type"] == "compatibility-snapshot", "dispatch must route GetSnapshotPlan")

stdout, stderr, status = Open3.capture3(
  *command,
  "dispatch",
  "GetPortalAccessPolicy",
  JSON.generate(["org.xnix.sample.notepad", "file-open"])
)
assert(status.success?, "dispatch GetPortalAccessPolicy must exit successfully: #{stderr}")
portal_policy = JSON.parse(stdout)
assert(portal_policy["policy_type"] == "portal-access", "dispatch must route GetPortalAccessPolicy")

_stdout, stderr, status = Open3.capture3(*command, "dispatch", "Launch", JSON.generate(["org.xnix.sample.notepad", {}]))
assert(!status.success?, "dispatch must reject unsupported write methods until a backend exists")
assert(stderr.include?("unsupported runtime method"), "dispatch must explain unsupported methods")

puts "PASS: compatibility runtime dispatch unit tests"
