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
  "GetApplicationStateRoot",
  JSON.generate(["org.xnix.sample.notepad"])
)
assert(status.success?, "dispatch GetApplicationStateRoot must exit successfully: #{stderr}")
state_root = JSON.parse(stdout)
assert(state_root["root_type"] == "compatibility-application-state-root", "dispatch must route GetApplicationStateRoot")

stdout, stderr, status = Open3.capture3(
  *command,
  "dispatch",
  "GetCompatibilityPackageSource",
  JSON.generate(["org.xnix.sample.notepad"])
)
assert(status.success?, "dispatch GetCompatibilityPackageSource must exit successfully: #{stderr}")
package_source = JSON.parse(stdout)
assert(package_source["source_type"] == "compatibility-package-source", "dispatch must route GetCompatibilityPackageSource")

stdout, stderr, status = Open3.capture3(
  *command,
  "dispatch",
  "GetCompatibilityAcquisitionPreflight",
  JSON.generate(["org.xnix.sample.notepad"])
)
assert(status.success?, "dispatch GetCompatibilityAcquisitionPreflight must exit successfully: #{stderr}")
acquisition_preflight = JSON.parse(stdout)
assert(acquisition_preflight["preflight_type"] == "compatibility-acquisition-preflight", "dispatch must route GetCompatibilityAcquisitionPreflight")

stdout, stderr, status = Open3.capture3(
  *command,
  "dispatch",
  "GetCompatibilityArtifactManifest",
  JSON.generate(["org.xnix.sample.notepad"])
)
assert(status.success?, "dispatch GetCompatibilityArtifactManifest must exit successfully: #{stderr}")
artifact_manifest = JSON.parse(stdout)
assert(artifact_manifest["manifest_type"] == "compatibility-artifact-manifest", "dispatch must route GetCompatibilityArtifactManifest")

stdout, stderr, status = Open3.capture3(
  *command,
  "dispatch",
  "GetCompatibilityInstallPlan",
  JSON.generate(["org.xnix.sample.notepad", "development"])
)
assert(status.success?, "dispatch GetCompatibilityInstallPlan must exit successfully: #{stderr}")
install_plan = JSON.parse(stdout)
assert(install_plan["plan_type"] == "compatibility-install-plan", "dispatch must route GetCompatibilityInstallPlan")

stdout, stderr, status = Open3.capture3(
  *command,
  "dispatch",
  "GetBackendBinding",
  JSON.generate(["org.xnix.sample.notepad"])
)
assert(status.success?, "dispatch GetBackendBinding must exit successfully: #{stderr}")
backend_binding = JSON.parse(stdout)
assert(backend_binding["binding_type"] == "compatibility-backend-binding", "dispatch must route GetBackendBinding")

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
  "GetTestResult",
  JSON.generate(["org.xnix.sample.notepad"])
)
assert(status.success?, "dispatch GetTestResult must exit successfully: #{stderr}")
test_result = JSON.parse(stdout)
assert(test_result["result_type"] == "compatibility-test-result", "dispatch must route GetTestResult")

stdout, stderr, status = Open3.capture3(
  *command,
  "dispatch",
  "GetAIDiagnosticInput",
  JSON.generate(["org.xnix.sample.notepad"])
)
assert(status.success?, "dispatch GetAIDiagnosticInput must exit successfully: #{stderr}")
ai_input = JSON.parse(stdout)
assert(ai_input["input_type"] == "ai-diagnostic-input", "dispatch must route GetAIDiagnosticInput")

stdout, stderr, status = Open3.capture3(
  *command,
  "dispatch",
  "GetAIDiagnosticRecommendation",
  JSON.generate(["org.xnix.sample.notepad"])
)
assert(status.success?, "dispatch GetAIDiagnosticRecommendation must exit successfully: #{stderr}")
ai_recommendation = JSON.parse(stdout)
assert(ai_recommendation["recommendation_type"] == "ai-diagnostic-recommendation", "dispatch must route GetAIDiagnosticRecommendation")

stdout, stderr, status = Open3.capture3(
  *command,
  "dispatch",
  "GetAIRepairApprovalGate",
  JSON.generate(["org.xnix.sample.notepad"])
)
assert(status.success?, "dispatch GetAIRepairApprovalGate must exit successfully: #{stderr}")
ai_gate = JSON.parse(stdout)
assert(ai_gate["gate_type"] == "ai-repair-approval-gate", "dispatch must route GetAIRepairApprovalGate")

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

stdout, stderr, status = Open3.capture3(*command, "dispatch", "GetRuntimeServiceBinding")
assert(status.success?, "dispatch GetRuntimeServiceBinding must exit successfully: #{stderr}")
service_binding = JSON.parse(stdout)
assert(service_binding["binding_type"] == "runtime-service-binding", "dispatch must route GetRuntimeServiceBinding")

stdout, stderr, status = Open3.capture3(
  *command,
  "dispatch",
  "GetCompatibilitySettings",
  JSON.generate(["org.xnix.sample.notepad"])
)
assert(status.success?, "dispatch GetCompatibilitySettings must exit successfully: #{stderr}")
settings = JSON.parse(stdout)
assert(settings["request_type"] == "settings-model", "dispatch must route GetCompatibilitySettings")

_stdout, stderr, status = Open3.capture3(*command, "dispatch", "Launch", JSON.generate(["org.xnix.sample.notepad", {}]))
assert(!status.success?, "dispatch must reject unsupported write methods until a backend exists")
assert(stderr.include?("unsupported runtime method"), "dispatch must explain unsupported methods")

puts "PASS: compatibility runtime dispatch unit tests"
