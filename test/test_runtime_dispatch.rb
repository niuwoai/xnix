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
  "GetDesktopIconPlan",
  JSON.generate(["org.xnix.sample.notepad"])
)
assert(status.success?, "dispatch GetDesktopIconPlan must exit successfully: #{stderr}")
desktop_icon_plan = JSON.parse(stdout)
assert(desktop_icon_plan["plan_type"] == "desktop-icon-plan", "dispatch must route GetDesktopIconPlan")
assert(desktop_icon_plan["runtime_method"] == "GetDesktopIconPlan", "dispatch must expose desktop icon Runtime methods")
assert(desktop_icon_plan["desktop_file"] == "xnix-org.xnix.sample.notepad.desktop", "dispatch must expose desktop icon desktop files")
assert(desktop_icon_plan["launcher_url"] == "applications:xnix-org.xnix.sample.notepad.desktop", "dispatch must expose desktop icon launcher URLs")
assert(desktop_icon_plan["target_directory"] == "xdg-desktop-dir", "dispatch must keep desktop icon target abstract")
assert(desktop_icon_plan["desktop_icon_visible"], "dispatch must expose desktop icon visibility")
assert(!desktop_icon_plan["desktop_file_copy_enabled"], "dispatch must not copy desktop icon files")
assert(!desktop_icon_plan["desktop_file_write_enabled"], "dispatch must not write desktop icon files")
assert(!desktop_icon_plan["launch_enabled"], "dispatch must not enable launch from desktop icon previews")
assert(!desktop_icon_plan["backend_details_exposed"], "dispatch must keep desktop icon backend details hidden")

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
  "GetKRunnerQueryPlan",
  JSON.generate(["notepad"])
)
assert(status.success?, "dispatch GetKRunnerQueryPlan must exit successfully: #{stderr}")
krunner_query = JSON.parse(stdout)
assert(krunner_query["query_type"] == "krunner-query-plan", "dispatch must route GetKRunnerQueryPlan")
assert(krunner_query["matches"].first["application_id"] == "org.xnix.sample.notepad", "dispatch must expose KRunner matches")
assert(!krunner_query["summary"]["backend_launch_enabled"], "dispatch must keep KRunner backend launch disabled")

stdout, stderr, status = Open3.capture3(
  *command,
  "dispatch",
  "GetKDEIntegrationStatus",
  JSON.generate([])
)
assert(status.success?, "dispatch GetKDEIntegrationStatus must exit successfully: #{stderr}")
kde_status = JSON.parse(stdout)
assert(kde_status["status_type"] == "kde-integration-status", "dispatch must route GetKDEIntegrationStatus")
assert(kde_status["entry_point_count"] == 7, "dispatch must expose seven KDE entry points")
assert(kde_status["entry_points"].all? { |entry| entry.fetch("c_runtime_backed") }, "dispatch must expose C Runtime-backed entry points")
assert(!kde_status["backend_details_exposed"], "dispatch must keep KDE integration backend details hidden")

stdout, stderr, status = Open3.capture3(
  *command,
  "dispatch",
  "GetKDEShellIntegrationPlan",
  JSON.generate([])
)
assert(status.success?, "dispatch GetKDEShellIntegrationPlan must exit successfully: #{stderr}")
kde_shell_plan = JSON.parse(stdout)
assert(kde_shell_plan["plan_type"] == "kde-shell-integration-plan", "dispatch must route GetKDEShellIntegrationPlan")
assert(kde_shell_plan["component_count"] == 9, "dispatch must expose KDE shell component counts")
assert(!kde_shell_plan["shell_configuration_written"], "dispatch must keep KDE shell writes disabled")
assert(!kde_shell_plan["component_activation_enabled"], "dispatch must keep KDE shell component activation disabled")
assert(!kde_shell_plan["backend_launch_enabled"], "dispatch must keep KDE shell backend launch disabled")

stdout, stderr, status = Open3.capture3(
  *command,
  "dispatch",
  "GetKWinWindowRulePlan",
  JSON.generate(["org.xnix.sample.notepad"])
)
assert(status.success?, "dispatch GetKWinWindowRulePlan must exit successfully: #{stderr}")
kwin_rule = JSON.parse(stdout)
assert(kwin_rule["request_type"] == "kwin-window-rule", "dispatch must route GetKWinWindowRulePlan")
assert(kwin_rule["match"]["class_group"] == "xnix-compatibility", "dispatch must expose KWin class groups")
assert(kwin_rule["set"]["desktop_file"] == "xnix-org.xnix.sample.notepad.desktop", "dispatch must expose KWin desktop files")
assert(!kwin_rule["safety"]["backend_details_exposed"], "dispatch must keep KWin backend details hidden")

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
  "GetCompatibilityCenterSummary",
  JSON.generate(["org.xnix.sample.notepad"])
)
assert(status.success?, "dispatch GetCompatibilityCenterSummary must exit successfully: #{stderr}")
compatibility_center_summary = JSON.parse(stdout)
assert(compatibility_center_summary["summary_type"] == "compatibility-center-summary", "dispatch must route GetCompatibilityCenterSummary")
assert(compatibility_center_summary["compatibility"]["state"] == "review-required", "dispatch must expose Compatibility Center state")
assert(!compatibility_center_summary["safety"]["backend_launch_enabled"], "dispatch must keep Compatibility Center backend launch disabled")

stdout, stderr, status = Open3.capture3(
  *command,
  "dispatch",
  "GetKDECenterPage",
  JSON.generate(["org.xnix.sample.notepad", "approved"])
)
assert(status.success?, "dispatch GetKDECenterPage must exit successfully: #{stderr}")
kde_center_page = JSON.parse(stdout)
assert(kde_center_page["request_type"] == "kde-center-page", "dispatch must route GetKDECenterPage")
assert(kde_center_page["runtime_method"] == "GetKDECenterPage", "dispatch must expose KDE center page Runtime methods")
assert(kde_center_page["activation_status_runtime_method"] == "GetDesktopActivationStatus", "dispatch must expose KDE center activation status Runtime methods")
assert(kde_center_page["activation_state"] == "ready-for-runtime-commit", "dispatch must expose KDE center activation state")
assert(!kde_center_page["activation_commit_enabled"], "dispatch must keep KDE center activation commit disabled")
assert(!kde_center_page["activation_launch_enabled"], "dispatch must keep KDE center activation launch disabled")
assert(kde_center_page["card_count"] == 7, "dispatch must expose KDE center page card counts")
assert(kde_center_page["settings_section_count"] == 5, "dispatch must expose KDE center page settings sections")
assert(!kde_center_page["launch_enabled"], "dispatch must not enable launch from KDE center pages")
assert(!kde_center_page["execution_started"], "dispatch must not start execution from KDE center pages")

stdout, stderr, status = Open3.capture3(
  *command,
  "dispatch",
  "GetKDECenterPageSections",
  JSON.generate(["org.xnix.sample.notepad", "approved"])
)
assert(status.success?, "dispatch GetKDECenterPageSections must exit successfully: #{stderr}")
kde_center_page_sections = JSON.parse(stdout)
assert(kde_center_page_sections["request_type"] == "kde-center-page-sections", "dispatch must route GetKDECenterPageSections")
assert(kde_center_page_sections["runtime_method"] == "GetKDECenterPageSections", "dispatch must expose KDE center page section Runtime methods")
assert(kde_center_page_sections["section_count"] == 5, "dispatch must expose KDE center page section counts")
assert(kde_center_page_sections["read_only_section_count"] == 5, "dispatch must keep KDE center page sections read-only")
assert(kde_center_page_sections["executable_section_count"].zero?, "dispatch must not expose executable KDE center page sections")
assert(kde_center_page_sections["runtime_methods"] == %w[GetCompatibilityCenterSummary GetDesktopActivationStatus GetCompatibilityActionQueue GetCompatibilitySettings GetDiagnostics], "dispatch must expose KDE center page section read methods")
assert(!kde_center_page_sections["section_actions_enabled"], "dispatch must not enable KDE center page section actions")
assert(!kde_center_page_sections["execution_started"], "dispatch must not start execution from KDE center page sections")

stdout, stderr, status = Open3.capture3(
  *command,
  "dispatch",
  "GetKDECenterPageSectionDetail",
  JSON.generate(["org.xnix.sample.notepad", "settings", "approved"])
)
assert(status.success?, "dispatch GetKDECenterPageSectionDetail must exit successfully: #{stderr}")
kde_center_page_section_detail = JSON.parse(stdout)
assert(kde_center_page_section_detail["request_type"] == "kde-center-page-section-detail", "dispatch must route GetKDECenterPageSectionDetail")
assert(kde_center_page_section_detail["runtime_method"] == "GetKDECenterPageSectionDetail", "dispatch must expose KDE center page section detail Runtime methods")
assert(kde_center_page_section_detail["section_id"] == "settings", "dispatch must expose selected KDE center page sections")
assert(kde_center_page_section_detail["section_runtime_method"] == "GetCompatibilitySettings", "dispatch must expose selected section read methods")
assert(kde_center_page_section_detail["section_read_model"] == "settings-model", "dispatch must expose selected section read models")
assert(kde_center_page_section_detail["available_section_ids"] == %w[overview activation actions settings diagnostics], "dispatch must expose selected section navigation")
assert(kde_center_page_section_detail["read_only_navigation"], "dispatch must keep section details read-only")
assert(!kde_center_page_section_detail["section_actions_enabled"], "dispatch must not enable KDE center page section detail actions")
assert(!kde_center_page_section_detail["execution_started"], "dispatch must not start execution from KDE center page section details")

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
  "GetBackendCapabilityMatrix",
  JSON.generate([])
)
assert(status.success?, "dispatch GetBackendCapabilityMatrix must exit successfully: #{stderr}")
backend_capability_matrix = JSON.parse(stdout)
assert(backend_capability_matrix["matrix_type"] == "compatibility-backend-capability-matrix", "dispatch must route GetBackendCapabilityMatrix")
assert(!backend_capability_matrix["backend_launch_enabled"], "dispatch must keep backend capability launch gated")

stdout, stderr, status = Open3.capture3(
  *command,
  "dispatch",
  "GetBackendSelectionPlan",
  JSON.generate(["org.xnix.sample.notepad"])
)
assert(status.success?, "dispatch GetBackendSelectionPlan must exit successfully: #{stderr}")
backend_selection_plan = JSON.parse(stdout)
assert(backend_selection_plan["plan_type"] == "compatibility-backend-selection-plan", "dispatch must route GetBackendSelectionPlan")
assert(!backend_selection_plan["selection_committed"], "dispatch must keep backend selection uncommitted")

stdout, stderr, status = Open3.capture3(
  *command,
  "dispatch",
  "GetBackendLifecycle",
  JSON.generate(["org.xnix.sample.notepad"])
)
assert(status.success?, "dispatch GetBackendLifecycle must exit successfully: #{stderr}")
backend_lifecycle = JSON.parse(stdout)
assert(backend_lifecycle["lifecycle_type"] == "compatibility-backend-lifecycle", "dispatch must route GetBackendLifecycle")

stdout, stderr, status = Open3.capture3(
  *command,
  "dispatch",
  "GetBackendEnvironmentPlan",
  JSON.generate(["org.xnix.sample.notepad"])
)
assert(status.success?, "dispatch GetBackendEnvironmentPlan must exit successfully: #{stderr}")
backend_environment_plan = JSON.parse(stdout)
assert(backend_environment_plan["plan_type"] == "compatibility-backend-environment-plan", "dispatch must route GetBackendEnvironmentPlan")

stdout, stderr, status = Open3.capture3(
  *command,
  "dispatch",
  "GetKDEApplicationSurfacePlan",
  JSON.generate(["org.xnix.sample.notepad"])
)
assert(status.success?, "dispatch GetKDEApplicationSurfacePlan must exit successfully: #{stderr}")
kde_application_surface_plan = JSON.parse(stdout)
assert(kde_application_surface_plan["plan_type"] == "kde-application-surface-plan", "dispatch must route GetKDEApplicationSurfacePlan")

stdout, stderr, status = Open3.capture3(
  *command,
  "dispatch",
  "GetDesktopResourceBridgePlan",
  JSON.generate(["org.xnix.sample.notepad"])
)
assert(status.success?, "dispatch GetDesktopResourceBridgePlan must exit successfully: #{stderr}")
desktop_resource_bridge_plan = JSON.parse(stdout)
assert(desktop_resource_bridge_plan["plan_type"] == "desktop-resource-bridge-plan", "dispatch must route GetDesktopResourceBridgePlan")

stdout, stderr, status = Open3.capture3(
  *command,
  "dispatch",
  "GetCompatibilityModeSwitchPlan",
  JSON.generate(["org.xnix.sample.notepad", "prefer-compatibility"])
)
assert(status.success?, "dispatch GetCompatibilityModeSwitchPlan must exit successfully: #{stderr}")
mode_switch_plan = JSON.parse(stdout)
assert(mode_switch_plan["plan_type"] == "compatibility-mode-switch-plan", "dispatch must route GetCompatibilityModeSwitchPlan")
assert(!mode_switch_plan["settings_persistence_enabled"], "dispatch must keep compatibility mode switches planned")

stdout, stderr, status = Open3.capture3(
  *command,
  "dispatch",
  "GetCompatibilityPermissionReviewPlan",
  JSON.generate(["org.xnix.sample.notepad"])
)
assert(status.success?, "dispatch GetCompatibilityPermissionReviewPlan must exit successfully: #{stderr}")
permission_review_plan = JSON.parse(stdout)
assert(permission_review_plan["plan_type"] == "compatibility-permission-review-plan", "dispatch must route GetCompatibilityPermissionReviewPlan")
assert(!permission_review_plan["permissions_granted"], "dispatch must keep compatibility permissions ungranted")

stdout, stderr, status = Open3.capture3(
  *command,
  "dispatch",
  "GetCompatibilityReviewFlowPlan",
  JSON.generate(["org.xnix.sample.notepad", "resource-access", "documents", "ask", "file-open"])
)
assert(status.success?, "dispatch GetCompatibilityReviewFlowPlan must exit successfully: #{stderr}")
review_flow_plan = JSON.parse(stdout)
assert(review_flow_plan["plan_type"] == "compatibility-review-flow-plan", "dispatch must route GetCompatibilityReviewFlowPlan")
assert(!review_flow_plan["apply_enabled"], "dispatch must keep compatibility review flows non-applying")
assert(!review_flow_plan["request_object_created"], "dispatch must keep review flow request creation gated")

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
  "GetExecutionReadiness",
  JSON.generate(["org.xnix.sample.notepad"])
)
assert(status.success?, "dispatch GetExecutionReadiness must exit successfully: #{stderr}")
execution_readiness = JSON.parse(stdout)
assert(execution_readiness["readiness_type"] == "compatibility-execution-readiness", "dispatch must route GetExecutionReadiness")

stdout, stderr, status = Open3.capture3(
  *command,
  "dispatch",
  "GetLaunchIntent",
  JSON.generate(["org.xnix.sample.notepad"])
)
assert(status.success?, "dispatch GetLaunchIntent must exit successfully: #{stderr}")
launch_intent = JSON.parse(stdout)
assert(launch_intent["intent_type"] == "runtime-launch-intent", "dispatch must route GetLaunchIntent")
assert(!launch_intent["execution_started"], "dispatch GetLaunchIntent must not start execution")

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
