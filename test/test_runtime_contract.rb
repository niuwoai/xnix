#!/usr/bin/env ruby
# frozen_string_literal: true

require "pathname"

def assert(condition, message)
  return if condition

  warn "FAIL: #{message}"
  exit 1
end

project_root = Pathname.new(__dir__).join("..").realpath
contract = project_root.join("runtime/dbus/org.xnix.Compatibility1.xml").read
service = project_root.join("runtime/dbus/org.xnix.Compatibility1.service").read
systemd_unit = project_root.join("runtime/systemd/xnix-compatd.service").read
plasmoid = project_root.join("kde/plasmoids/org.xnix.compatibilitycenter/metadata.json").read

%w[
  ListApplications
  GetApplication
  InstallRecipe
  Launch
  CreateSnapshot
  RestoreSnapshot
  GetDiagnostics
  GetEngineCatalog
  GetRunPlan
  GetDesktopActivationManifest
  GetKDEIntegrationStatus
  GetKDEApplicationSurfacePlan
  GetDesktopResourceBridgePlan
  GetDesktopEntryPlan
  GetTaskManagerIdentityPlan
  GetKWinWindowRulePlan
  GetFileAssociationPlan
  GetNotificationPlan
  GetTrayStatus
  GetKRunnerQueryPlan
  GetApplicationStateRoot
  GetCompatibilityPackageSource
  GetCompatibilityAcquisitionPreflight
  GetCompatibilityArtifactManifest
  GetCompatibilityInstallPlan
  GetBackendBinding
  GetBackendLifecycle
  GetBackendEnvironmentPlan
  GetRepairPlan
  GetTestPlan
  GetTestResult
  GetExecutionReadiness
  GetLaunchIntent
  GetAIDiagnosticInput
  GetAIDiagnosticRecommendation
  GetAIRepairApprovalGate
  GetSnapshotPlan
  GetPortalAccessPolicy
  GetRuntimeServiceBinding
  GetRuntimeLiveOwnerGate
  GetRuntimeOwnerSmokePlan
  GetRuntimeMethodParityManifest
  GetRuntimeWriteGate
  GetCompatibilitySettings
  GetCompatibilitySettingsChangePlan
  GetCompatibilityModeSwitchPlan
  GetCompatibilityPermissionReviewPlan
  GetCompatibilityActionQueue
  GetCompatibilityActionReviewReceipt
  GetCompatibilityCenterSummary
].each do |method|
  assert(contract.include?("name=\"#{method}\""), "D-Bus contract must expose #{method}")
end
assert(contract.include?("name=\"RequestCompleted\""), "D-Bus contract must emit async completion signals")
assert(service.include?("Name=org.xnix.Compatibility1"), "D-Bus service must use the runtime bus name")
assert(systemd_unit.include?("ProtectSystem=strict"), "runtime service must protect the system filesystem")
assert(systemd_unit.include?("NoNewPrivileges=yes"), "runtime service must forbid privilege escalation")
assert(plasmoid.include?("org.xnix.compatibilitycenter"), "KDE package must keep a stable plugin id")
assert(plasmoid.include?("X-Plasma-API-Minimum-Version"), "KDE package must target Plasma 6")

puts "PASS: compatibility runtime contract unit tests"
