#!/usr/bin/env ruby
# frozen_string_literal: true

require "pathname"

def assert(condition, message)
  return if condition

  warn "FAIL: #{message}"
  exit 1
end

project_root = Pathname.new(__dir__).join("..").realpath
script = project_root.join("scripts/dbus_session_smoke.rb")
source = project_root.join("runtime/dbus/xnix_compatd_smoke.c")
introspection = project_root.join("runtime/dbus/xnix_compatd_introspection.inc")
dockerfile = project_root.join("Dockerfile").read
contents = script.read

assert(script.file?, "D-Bus smoke script must exist")
assert(script.executable?, "D-Bus smoke script must be executable")
assert(source.file?, "D-Bus smoke adapter source must exist")
assert(introspection.file?, "D-Bus smoke adapter introspection include must exist")
assert(source.read.include?("xnix_compatd_introspection.inc"), "D-Bus smoke adapter must include its introspection XML")
assert(introspection.read.include?("GetDesktopActivationTransactionPreview"), "D-Bus smoke adapter introspection must expose activation transaction previews")
assert(introspection.read.include?("GetDesktopActivationStatus"), "D-Bus smoke adapter introspection must expose activation status previews")
assert(dockerfile.include?("libglib2.0-dev"), "Dockerfile must install GIO headers for the smoke adapter")
assert(dockerfile.include?("pkg-config"), "Dockerfile must install pkg-config for the smoke adapter build")
assert(dockerfile.include?("xnix_compatd_smoke.c"), "Dockerfile must compile the smoke adapter")
assert(contents.include?("dbus-run-session"), "D-Bus smoke script must start a session bus")
assert(contents.include?("gdbus"), "D-Bus smoke script must use gdbus for runtime calls")
assert(contents.include?("ListApplications"), "D-Bus smoke script must call ListApplications")
assert(contents.include?("GetDiagnostics"), "D-Bus smoke script must call GetDiagnostics")
assert(contents.include?("GetEngineCatalog"), "D-Bus smoke script must call GetEngineCatalog")
assert(contents.include?("GetRunPlan"), "D-Bus smoke script must call GetRunPlan")
assert(contents.include?("GetApplicationStateRoot"), "D-Bus smoke script must call GetApplicationStateRoot")
assert(contents.include?("GetCompatibilityPackageSource"), "D-Bus smoke script must call GetCompatibilityPackageSource")
assert(contents.include?("GetCompatibilityAcquisitionPreflight"), "D-Bus smoke script must call GetCompatibilityAcquisitionPreflight")
assert(contents.include?("GetCompatibilityArtifactManifest"), "D-Bus smoke script must call GetCompatibilityArtifactManifest")
assert(contents.include?("GetCompatibilityInstallPlan"), "D-Bus smoke script must call GetCompatibilityInstallPlan")
assert(contents.include?("GetBackendBinding"), "D-Bus smoke script must call GetBackendBinding")
assert(contents.include?("GetBackendCapabilityMatrix"), "D-Bus smoke script must call GetBackendCapabilityMatrix")
assert(contents.include?("GetBackendSelectionPlan"), "D-Bus smoke script must call GetBackendSelectionPlan")
assert(contents.include?("GetBackendLifecycle"), "D-Bus smoke script must call GetBackendLifecycle")
assert(contents.include?("GetBackendEnvironmentPlan"), "D-Bus smoke script must call GetBackendEnvironmentPlan")
assert(contents.include?("GetRepairPlan"), "D-Bus smoke script must call GetRepairPlan")
assert(contents.include?("GetTestPlan"), "D-Bus smoke script must call GetTestPlan")
assert(contents.include?("GetTestResult"), "D-Bus smoke script must call GetTestResult")
assert(contents.include?("GetExecutionReadiness"), "D-Bus smoke script must call GetExecutionReadiness")
assert(contents.include?("GetLaunchIntent"), "D-Bus smoke script must call GetLaunchIntent")
assert(contents.include?("GetAIDiagnosticInput"), "D-Bus smoke script must call GetAIDiagnosticInput")
assert(contents.include?("GetAIDiagnosticRecommendation"), "D-Bus smoke script must call GetAIDiagnosticRecommendation")
assert(contents.include?("GetAIRepairApprovalGate"), "D-Bus smoke script must call GetAIRepairApprovalGate")
assert(contents.include?("GetSnapshotPlan"), "D-Bus smoke script must call GetSnapshotPlan")
assert(contents.include?("GetPortalAccessPolicy"), "D-Bus smoke script must call GetPortalAccessPolicy")
assert(contents.include?("GetRuntimeServiceBinding"), "D-Bus smoke script must call GetRuntimeServiceBinding")
assert(contents.include?("GetRuntimeLiveOwnerGate"), "D-Bus smoke script must call GetRuntimeLiveOwnerGate")
assert(contents.include?("GetRuntimeOwnerSmokePlan"), "D-Bus smoke script must call GetRuntimeOwnerSmokePlan")
assert(contents.include?("GetRuntimeMethodParityManifest"), "D-Bus smoke script must call GetRuntimeMethodParityManifest")
assert(contents.include?("GetRuntimeWriteGate"), "D-Bus smoke script must call GetRuntimeWriteGate")
assert(contents.include?("GetCompatibilitySettings"), "D-Bus smoke script must call GetCompatibilitySettings")
assert(contents.include?("GetCompatibilitySettingsChangePlan"), "D-Bus smoke script must call GetCompatibilitySettingsChangePlan")
assert(contents.include?("GetDesktopIconPlan"), "D-Bus smoke script must call GetDesktopIconPlan")
assert(contents.include?("GetCompatibilityModeSwitchPlan"), "D-Bus smoke script must call GetCompatibilityModeSwitchPlan")
assert(contents.include?("GetCompatibilityPermissionReviewPlan"), "D-Bus smoke script must call GetCompatibilityPermissionReviewPlan")
assert(contents.include?("GetCompatibilityReviewFlowPlan"), "D-Bus smoke script must call GetCompatibilityReviewFlowPlan")
assert(contents.include?("GetCompatibilityActionQueue"), "D-Bus smoke script must call GetCompatibilityActionQueue")
assert(contents.include?("GetCompatibilityActionReviewReceipt"), "D-Bus smoke script must call GetCompatibilityActionReviewReceipt")
assert(contents.include?("GetKDECenterPage"), "D-Bus smoke script must call GetKDECenterPage")
assert(contents.include?("GetKDECenterPageSections"), "D-Bus smoke script must call GetKDECenterPageSections")
assert(contents.include?("GetKDECenterPageSectionDetail"), "D-Bus smoke script must call GetKDECenterPageSectionDetail")
assert(contents.include?("GetKRunnerQueryPlan"), "D-Bus smoke script must call GetKRunnerQueryPlan")
assert(contents.include?("bin/xnix-krunner-model"), "D-Bus smoke script must verify the KRunner D-Bus read model")
assert(contents.include?("GetKDEIntegrationStatus"), "D-Bus smoke script must call GetKDEIntegrationStatus")
assert(contents.include?("bin/xnix-kde-integration-status"), "D-Bus smoke script must verify the KDE integration D-Bus read model")
assert(contents.include?("GetKDEShellIntegrationPlan"), "D-Bus smoke script must call GetKDEShellIntegrationPlan")
assert(contents.include?("GetKWinWindowRulePlan"), "D-Bus smoke script must call GetKWinWindowRulePlan")
assert(contents.include?("bin/xnix-kwin-window-rule"), "D-Bus smoke script must verify the KWin D-Bus read model")

puts "PASS: compatibility runtime D-Bus smoke script unit tests"
