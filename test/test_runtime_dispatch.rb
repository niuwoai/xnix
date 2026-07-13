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
  "GetDesktopEntryPlan",
  JSON.generate(["org.xnix.sample.notepad"])
)
assert(status.success?, "dispatch GetDesktopEntryPlan must exit successfully: #{stderr}")
desktop_entry_plan = JSON.parse(stdout)
assert(desktop_entry_plan["plan_type"] == "desktop-entry-plan", "dispatch must route GetDesktopEntryPlan")
assert(desktop_entry_plan["standard_desktop_entry"], "dispatch must expose standard desktop entry status")
assert(desktop_entry_plan["launch_uses_runtime"], "dispatch must keep desktop entry launch Runtime-managed")
assert(!desktop_entry_plan["safety"]["backend_command_exposed"], "dispatch must keep backend commands hidden")

stdout, stderr, status = Open3.capture3(
  *command,
  "dispatch",
  "GetTaskManagerIdentityPlan",
  JSON.generate(["org.xnix.sample.notepad"])
)
assert(status.success?, "dispatch GetTaskManagerIdentityPlan must exit successfully: #{stderr}")
task_manager_identity_plan = JSON.parse(stdout)
assert(task_manager_identity_plan["plan_type"] == "task-manager-identity-plan", "dispatch must route GetTaskManagerIdentityPlan")
assert(task_manager_identity_plan["desktop_file"] == "xnix-org.xnix.sample.notepad.desktop", "dispatch must expose task manager desktop files")
assert(task_manager_identity_plan["task_manager"]["grouping_key"] == "org.xnix.sample.notepad", "dispatch must expose task manager grouping keys")
assert(task_manager_identity_plan["task_manager"]["pinning_allowed"], "dispatch must keep task manager pinning allowed")
assert(!task_manager_identity_plan["safety"]["backend_details_exposed"], "dispatch must keep task manager backend details hidden")

stdout, stderr, status = Open3.capture3(
  *command,
  "dispatch",
  "GetFileAssociationPlan",
  JSON.generate(["org.xnix.sample.notepad"])
)
assert(status.success?, "dispatch GetFileAssociationPlan must exit successfully: #{stderr}")
file_association_plan = JSON.parse(stdout)
assert(file_association_plan["plan_type"] == "file-association-plan", "dispatch must route GetFileAssociationPlan")
assert(file_association_plan["application"]["desktop_file"] == "xnix-org.xnix.sample.notepad.desktop", "dispatch must expose file association desktop files")
assert(file_association_plan["mimeapps"]["path"] == "usr/share/applications/mimeapps.list", "dispatch must expose mimeapps paths")
assert(file_association_plan["associations"].length == 2, "dispatch must expose file association records")
assert(file_association_plan["associations"].all? { |entry| entry["file_open"]["portal_required"] }, "dispatch must keep file opens Portal-mediated")
assert(!file_association_plan["safety"]["backend_details_exposed"], "dispatch must keep file association backend details hidden")

stdout, stderr, status = Open3.capture3(
  *command,
  "dispatch",
  "GetNotificationPlan",
  JSON.generate(["org.xnix.sample.notepad", "approval-required"])
)
assert(status.success?, "dispatch GetNotificationPlan must exit successfully: #{stderr}")
notification_plan = JSON.parse(stdout)
assert(notification_plan["plan_type"] == "notification-plan", "dispatch must route GetNotificationPlan")
assert(notification_plan["event_type"] == "approval-required", "dispatch must preserve notification event types")
assert(notification_plan["notification"]["urgency"] == "critical", "dispatch must expose notification urgency")
assert(notification_plan["notification"]["actions"].include?("review-request"), "dispatch must expose notification review actions")
assert(notification_plan["safety"]["requires_user_review"], "dispatch must expose notification review requirements")
assert(!notification_plan["safety"]["action_execution_enabled"], "dispatch must keep notification actions gated")
assert(!notification_plan["safety"]["backend_details_exposed"], "dispatch must keep notification backend details hidden")

stdout, stderr, status = Open3.capture3(
  *command,
  "dispatch",
  "GetTrayStatus"
)
assert(status.success?, "dispatch GetTrayStatus must exit successfully: #{stderr}")
tray_status = JSON.parse(stdout)
assert(tray_status["status_type"] == "tray-status-plan", "dispatch must route GetTrayStatus")
assert(tray_status["runtime_activity"]["active_application_count"] == 1, "dispatch must expose tray active counts")
assert(tray_status["compatibility_status"]["state"] == "attention-required", "dispatch must expose tray attention state")
assert(tray_status["tray_bridge"]["state"] == "idle", "dispatch must expose Ruby tray bridge state")
assert(!tray_status["safety"]["live_backend_bridge_enabled"], "dispatch must keep live tray bridge gated")
assert(!tray_status["safety"]["backend_details_exposed"], "dispatch must keep tray backend details hidden")

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
  "GetCompatibilityActionQueue",
  JSON.generate(["org.xnix.sample.notepad"])
)
assert(status.success?, "dispatch GetCompatibilityActionQueue must exit successfully: #{stderr}")
action_queue = JSON.parse(stdout)
assert(action_queue["queue_type"] == "compatibility-center-action-queue", "dispatch must route GetCompatibilityActionQueue")

stdout, stderr, status = Open3.capture3(
  *command,
  "dispatch",
  "GetCompatibilityActionReviewReceipt",
  JSON.generate(["org.xnix.sample.notepad", "review-ai-repair", "approved"])
)
assert(status.success?, "dispatch GetCompatibilityActionReviewReceipt must exit successfully: #{stderr}")
action_review = JSON.parse(stdout)
assert(action_review["receipt_type"] == "compatibility-center-action-review-receipt", "dispatch must route GetCompatibilityActionReviewReceipt")

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

stdout, stderr, status = Open3.capture3(*command, "dispatch", "GetRuntimeLiveOwnerGate")
assert(status.success?, "dispatch GetRuntimeLiveOwnerGate must exit successfully: #{stderr}")
live_owner_gate = JSON.parse(stdout)
assert(live_owner_gate["gate_type"] == "runtime-live-owner-gate", "dispatch must route GetRuntimeLiveOwnerGate")
assert(!live_owner_gate["owner_transition_ready"], "dispatch must keep live owner transition gated")

stdout, stderr, status = Open3.capture3(*command, "dispatch", "GetRuntimeOwnerSmokePlan")
assert(status.success?, "dispatch GetRuntimeOwnerSmokePlan must exit successfully: #{stderr}")
owner_smoke_plan = JSON.parse(stdout)
assert(owner_smoke_plan["plan_type"] == "runtime-owner-smoke-plan", "dispatch must route GetRuntimeOwnerSmokePlan")
assert(!owner_smoke_plan["production_bus_claimed"], "dispatch must keep production bus ownership unclaimed")

stdout, stderr, status = Open3.capture3(*command, "dispatch", "GetRuntimeMethodParityManifest")
assert(status.success?, "dispatch GetRuntimeMethodParityManifest must exit successfully: #{stderr}")
method_parity = JSON.parse(stdout)
assert(method_parity["manifest_type"] == "runtime-method-parity-manifest", "dispatch must route GetRuntimeMethodParityManifest")
assert(method_parity["read_only_method_parity_ready"], "dispatch must report method parity readiness")
assert(!method_parity["write_methods_supported"], "dispatch must not claim write method support")

stdout, stderr, status = Open3.capture3(
  *command,
  "dispatch",
  "GetRuntimeWriteGate",
  JSON.generate(["Launch"])
)
assert(status.success?, "dispatch GetRuntimeWriteGate must exit successfully: #{stderr}")
write_gate = JSON.parse(stdout)
assert(write_gate["gate_type"] == "runtime-write-gate", "dispatch must route GetRuntimeWriteGate")
assert(write_gate["method_name"] == "Launch", "dispatch must preserve Runtime write gate method names")
assert(write_gate["gate_decision"] == "blocked-until-production-backend", "dispatch must keep write methods gated")
assert(!write_gate["write_method_enabled"], "dispatch must not enable Runtime write methods")
assert(!write_gate["dispatch_enabled"], "dispatch must not enable Runtime write dispatch")

stdout, stderr, status = Open3.capture3(
  *command,
  "dispatch",
  "GetCompatibilitySettings",
  JSON.generate(["org.xnix.sample.notepad"])
)
assert(status.success?, "dispatch GetCompatibilitySettings must exit successfully: #{stderr}")
settings = JSON.parse(stdout)
assert(settings["request_type"] == "settings-model", "dispatch must route GetCompatibilitySettings")

stdout, stderr, status = Open3.capture3(
  *command,
  "dispatch",
  "GetCompatibilitySettingsChangePlan",
  JSON.generate(["org.xnix.sample.notepad", "resource-access", "documents", "ask"])
)
assert(status.success?, "dispatch GetCompatibilitySettingsChangePlan must exit successfully: #{stderr}")
settings_change = JSON.parse(stdout)
assert(settings_change["plan_type"] == "settings-change-plan", "dispatch must route GetCompatibilitySettingsChangePlan")
assert(!settings_change["apply_enabled"], "dispatch must keep settings changes planned")

_stdout, stderr, status = Open3.capture3(*command, "dispatch", "Launch", JSON.generate(["org.xnix.sample.notepad", {}]))
assert(!status.success?, "dispatch must reject unsupported write methods until a backend exists")
assert(stderr.include?("WriteMethodDisabled"), "dispatch must explain gated write methods")

puts "PASS: compatibility runtime dispatch unit tests"
