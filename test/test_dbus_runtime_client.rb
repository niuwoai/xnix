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
    when "org.xnix.Compatibility1.GetDesktopActivationManifest"
      [
        "({'manifest_type': <'desktop-activation'>, 'desktop': <'KDE Plasma'>, 'artifact_count': <7>, 'activation_ready': <true>, 'files_written': <false>, 'host_root_modified': <false>, 'backend_details_exposed': <false>},)\n",
        "",
        Status.new(true)
      ]
    when "org.xnix.Compatibility1.GetDesktopEntryPlan"
      [
        "({'plan_type': <'desktop-entry-plan'>, 'desktop_file': <'xnix-org.xnix.sample.notepad.desktop'>, 'name': <'Sample Notepad'>, 'exec': <'xnix-compat-launch --app org.xnix.sample.notepad %U'>, 'standard_desktop_entry': <true>, 'launch_uses_runtime': <true>, 'accepts_file_uris': <true>, 'files_written': <false>, 'host_root_modified': <false>, 'backend_command_exposed': <false>, 'raw_windows_executable_exposed': <false>, 'compatibility_storage_path_exposed': <false>, 'backend_details_exposed': <false>},)\n",
        "",
        Status.new(true)
      ]
    when "org.xnix.Compatibility1.GetTaskManagerIdentityPlan"
      [
        "({'plan_type': <'task-manager-identity-plan'>, 'desktop_file': <'xnix-org.xnix.sample.notepad.desktop'>, 'launcher_url': <'applications:xnix-org.xnix.sample.notepad.desktop'>, 'window_kind': <'compatibility-application'>, 'class_group': <'xnix-compatibility'>, 'grouping_key': <'org.xnix.sample.notepad'>, 'pinning_allowed': <true>, 'restore_allowed': <true>, 'skip_taskbar': <false>, 'show_in_switcher': <true>, 'prefer_existing_window': <true>, 'window_manager_policy_only': <true>, 'runtime_owns_backend_policy': <true>, 'host_root_modified': <false>, 'backend_details_exposed': <false>},)\n",
        "",
        Status.new(true)
      ]
    when "org.xnix.Compatibility1.GetPortalRequestPlan"
      [
        "({'request_type': <'portal-request-plan'>, 'operation': <'file-open'>, 'decision': <'ask'>, 'request_allowed': <true>, 'portal_required': <true>, 'request_object_created': <false>, 'permission_granted': <false>, 'host_permission_changed': <false>, 'backend_details_exposed': <false>},)\n",
        "",
        Status.new(true)
      ]
    when "org.xnix.Compatibility1.GetApplicationStateRoot"
      [
        "({'root_type': <'compatibility-application-state-root'>, 'allocation_state': <'planned'>, 'snapshot_eligible': <true>, 'user_documents_included': <false>, 'backend_details_exposed': <false>},)\n",
        "",
        Status.new(true)
      ]
    when "org.xnix.Compatibility1.GetCompatibilityPackageSource"
      [
        "({'source_type': <'compatibility-package-source'>, 'source_selection_state': <'planned'>, 'package_source_ready': <false>, 'install_enabled': <false>, 'backend_details_exposed': <false>},)\n",
        "",
        Status.new(true)
      ]
    when "org.xnix.Compatibility1.GetCompatibilityAcquisitionPreflight"
      [
        "({'preflight_type': <'compatibility-acquisition-preflight'>, 'preflight_state': <'planned'>, 'acquisition_ready': <false>, 'download_enabled': <false>, 'install_enabled': <false>, 'backend_details_exposed': <false>},)\n",
        "",
        Status.new(true)
      ]
    when "org.xnix.Compatibility1.GetCompatibilityArtifactManifest"
      [
        "({'manifest_type': <'compatibility-artifact-manifest'>, 'manifest_state': <'planned'>, 'manifest_ready': <false>, 'signature_verified': <false>, 'download_enabled': <false>, 'backend_details_exposed': <false>},)\n",
        "",
        Status.new(true)
      ]
    when "org.xnix.Compatibility1.GetCompatibilityInstallPlan"
      [
        "({'plan_type': <'compatibility-install-plan'>, 'install_state': <'planned'>, 'install_ready': <false>, 'desktop_activation_ready': <false>, 'download_enabled': <false>, 'install_enabled': <false>, 'backend_details_exposed': <false>},)\n",
        "",
        Status.new(true)
      ]
    when "org.xnix.Compatibility1.GetBackendBinding"
      [
        "({'binding_type': <'compatibility-backend-binding'>, 'selected_strategy': <'automatic-managed'>, 'managed_binding_ready': <false>, 'launch_enabled': <false>, 'backend_details_exposed': <false>},)\n",
        "",
        Status.new(true)
      ]
    when "org.xnix.Compatibility1.GetRepairPlan"
      [
        "({'plan_type': <'compatibility-repair'>, 'issue': <'engine-binding-pending'>, 'snapshot_required': <true>},)\n",
        "",
        Status.new(true)
      ]
    when "org.xnix.Compatibility1.GetTestPlan"
      [
        "({'plan_type': <'compatibility-test'>, 'test_type': <'preflight'>, 'runtime_owned': <true>, 'backend_details_exposed': <false>},)\n",
        "",
        Status.new(true)
      ]
    when "org.xnix.Compatibility1.GetTestResult"
      [
        "({'result_type': <'compatibility-test-result'>, 'test_type': <'preflight'>, 'overall_status': <'pending'>, 'safe_for_ai_diagnostics': <true>, 'backend_details_exposed': <false>},)\n",
        "",
        Status.new(true)
      ]
    when "org.xnix.Compatibility1.GetAIDiagnosticInput"
      [
        "({'input_type': <'ai-diagnostic-input'>, 'safe_for_ai_diagnostics': <true>, 'ai_provider_called': <false>, 'network_required': <false>, 'backend_details_exposed': <false>},)\n",
        "",
        Status.new(true)
      ]
    when "org.xnix.Compatibility1.GetAIDiagnosticRecommendation"
      [
        "({'recommendation_type': <'ai-diagnostic-recommendation'>, 'safe_for_ai_diagnostics': <true>, 'ai_provider_called': <false>, 'network_required': <false>, 'backend_details_exposed': <false>},)\n",
        "",
        Status.new(true)
      ]
    when "org.xnix.Compatibility1.GetAIRepairApprovalGate"
      [
        "({'gate_type': <'ai-repair-approval-gate'>, 'gate_decision': <'blocked-until-approval'>, 'auto_execution_allowed': <false>, 'repair_executed': <false>, 'backend_details_exposed': <false>},)\n",
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
    when "org.xnix.Compatibility1.GetRuntimeServiceBinding"
      [
        "({'binding_type': <'runtime-service-binding'>, 'activation_binding_ready': <true>, 'live_dbus_owner_ready': <false>, 'backend_details_exposed': <false>},)\n",
        "",
        Status.new(true)
      ]
    when "org.xnix.Compatibility1.GetRuntimeLiveOwnerGate"
      [
        "({'gate_type': <'runtime-live-owner-gate'>, 'activation_binding_ready': <true>, 'live_dbus_owner_ready': <false>, 'production_owner_enabled': <false>, 'owner_transition_ready': <false>, 'smoke_adapter_is_production_owner': <false>, 'kde_may_claim_runtime_ownership': <false>, 'backend_details_exposed': <false>},)\n",
        "",
        Status.new(true)
      ]
    when "org.xnix.Compatibility1.GetRuntimeOwnerSmokePlan"
      [
        "({'plan_type': <'runtime-owner-smoke-plan'>, 'smoke_state': <'planned'>, 'smoke_environment': <'restricted-session'>, 'activation_binding_ready': <true>, 'live_dbus_owner_ready': <false>, 'production_owner_enabled': <false>, 'owner_transition_ready': <false>, 'pending_step_count': <6>, 'system_service_started': <false>, 'production_bus_claimed': <false>, 'backend_details_exposed': <false>},)\n",
        "",
        Status.new(true)
      ]
    when "org.xnix.Compatibility1.GetRuntimeMethodParityManifest"
      [
        "({'manifest_type': <'runtime-method-parity-manifest'>, 'method_count': <32>, 'read_only_method_parity_ready': <true>, 'passed_check_count': <5>, 'blocked_check_count': <0>, 'write_methods_supported': <false>, 'write_method_dispatch_enabled': <false>, 'backend_details_exposed': <false>},)\n",
        "",
        Status.new(true)
      ]
    when "org.xnix.Compatibility1.GetRuntimeWriteGate"
      [
        "({'gate_type': <'runtime-write-gate'>, 'method_name': <'Launch'>, 'gate_decision': <'blocked-until-production-backend'>, 'write_method_enabled': <false>, 'dispatch_enabled': <false>, 'request_object_created': <false>, 'required_gate_count': <6>, 'denial_error_name': <'org.xnix.Compatibility1.Error.WriteMethodDisabled'>, 'backend_details_exposed': <false>},)\n",
        "",
        Status.new(true)
      ]
    when "org.xnix.Compatibility1.GetCompatibilitySettings"
      [
        "({'request_type': <'settings-model'>, 'settings_state': <'planned'>, 'settings_persisted': <false>, 'host_root_modified': <false>, 'backend_details_exposed': <false>},)\n",
        "",
        Status.new(true)
      ]
    when "org.xnix.Compatibility1.GetCompatibilitySettingsChangePlan"
      [
        "({'plan_type': <'settings-change-plan'>, 'change_state': <'planned'>, 'apply_enabled': <false>, 'settings_persisted': <false>, 'portal_policy_review_required': <true>, 'snapshot_recommended': <false>, 'backend_details_exposed': <false>},)\n",
        "",
        Status.new(true)
      ]
    when "org.xnix.Compatibility1.GetCompatibilityActionQueue"
      [
        "({'queue_type': <'compatibility-center-action-queue'>, 'action_count': <5>, 'pending_action_count': <5>, 'user_review_required_count': <3>, 'execution_enabled': <false>, 'repair_execution_enabled': <false>, 'settings_persistence_enabled': <false>, 'backend_details_exposed': <false>},)\n",
        "",
        Status.new(true)
      ]
    when "org.xnix.Compatibility1.GetCompatibilityActionReviewReceipt"
      [
        "({'receipt_type': <'compatibility-center-action-review-receipt'>, 'action_id': <'review-ai-repair'>, 'decision': <'approved'>, 'decision_recorded': <true>, 'execution_enabled': <false>, 'repair_execution_enabled': <false>, 'settings_persistence_enabled': <false>, 'resource_grant_created': <false>, 'backend_details_exposed': <false>},)\n",
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

desktop_activation = client.desktop_activation_manifest("org.xnix.sample.notepad")
assert(desktop_activation["manifest_type"] == "desktop-activation", "D-Bus client must parse desktop activation manifests")
assert(desktop_activation["desktop"] == "KDE Plasma", "D-Bus client must parse desktop activation target")
assert(desktop_activation["artifact_count"] == 7, "D-Bus client must parse desktop activation artifact counts")
assert(desktop_activation["activation_ready"], "D-Bus client must parse desktop activation readiness")
assert(!desktop_activation["files_written"], "D-Bus client must parse desktop activation write safety")
assert(!desktop_activation["host_root_modified"], "D-Bus client must parse desktop activation host-root safety")
assert(!desktop_activation["backend_details_exposed"], "D-Bus client must parse desktop activation backend detail status")

desktop_entry_plan = client.desktop_entry_plan("org.xnix.sample.notepad")
assert(desktop_entry_plan["plan_type"] == "desktop-entry-plan", "D-Bus client must parse desktop entry plans")
assert(desktop_entry_plan["desktop_file"] == "xnix-org.xnix.sample.notepad.desktop", "D-Bus client must parse desktop file names")
assert(desktop_entry_plan["exec"] == "xnix-compat-launch --app org.xnix.sample.notepad %U", "D-Bus client must parse managed launcher commands")
assert(desktop_entry_plan["standard_desktop_entry"], "D-Bus client must parse standard desktop entry status")
assert(desktop_entry_plan["launch_uses_runtime"], "D-Bus client must parse Runtime launcher status")
assert(desktop_entry_plan["accepts_file_uris"], "D-Bus client must parse file URI status")
assert(!desktop_entry_plan["files_written"], "D-Bus client must parse desktop entry write safety")
assert(!desktop_entry_plan["host_root_modified"], "D-Bus client must parse desktop entry host-root safety")
assert(!desktop_entry_plan["backend_command_exposed"], "D-Bus client must parse backend command hiding")
assert(!desktop_entry_plan["raw_windows_executable_exposed"], "D-Bus client must parse Windows executable hiding")
assert(!desktop_entry_plan["compatibility_storage_path_exposed"], "D-Bus client must parse prefix hiding")
assert(!desktop_entry_plan["backend_details_exposed"], "D-Bus client must parse desktop entry backend detail status")

task_manager_identity_plan = client.task_manager_identity_plan("org.xnix.sample.notepad")
assert(task_manager_identity_plan["plan_type"] == "task-manager-identity-plan", "D-Bus client must parse task manager identity plans")
assert(task_manager_identity_plan["desktop_file"] == "xnix-org.xnix.sample.notepad.desktop", "D-Bus client must parse task manager desktop files")
assert(task_manager_identity_plan["launcher_url"] == "applications:xnix-org.xnix.sample.notepad.desktop", "D-Bus client must parse task manager launcher URLs")
assert(task_manager_identity_plan["window_kind"] == "compatibility-application", "D-Bus client must parse task manager window kinds")
assert(task_manager_identity_plan["class_group"] == "xnix-compatibility", "D-Bus client must parse task manager class groups")
assert(task_manager_identity_plan["grouping_key"] == "org.xnix.sample.notepad", "D-Bus client must parse task manager grouping keys")
assert(task_manager_identity_plan["pinning_allowed"], "D-Bus client must parse task manager pinning status")
assert(task_manager_identity_plan["restore_allowed"], "D-Bus client must parse task manager restore status")
assert(!task_manager_identity_plan["skip_taskbar"], "D-Bus client must parse taskbar visibility")
assert(task_manager_identity_plan["show_in_switcher"], "D-Bus client must parse switcher visibility")
assert(task_manager_identity_plan["prefer_existing_window"], "D-Bus client must parse restore preferences")
assert(task_manager_identity_plan["window_manager_policy_only"], "D-Bus client must parse window-manager policy scope")
assert(task_manager_identity_plan["runtime_owns_backend_policy"], "D-Bus client must parse Runtime backend policy ownership")
assert(!task_manager_identity_plan["host_root_modified"], "D-Bus client must parse task manager host-root safety")
assert(!task_manager_identity_plan["backend_details_exposed"], "D-Bus client must parse task manager backend detail status")

portal_request_plan = client.portal_request_plan("org.xnix.sample.notepad", "file-open")
assert(portal_request_plan["request_type"] == "portal-request-plan", "D-Bus client must parse Portal request plans")
assert(portal_request_plan["operation"] == "file-open", "D-Bus client must parse Portal request operations")
assert(portal_request_plan["request_allowed"], "D-Bus client must parse Portal request allowance")
assert(portal_request_plan["portal_required"], "D-Bus client must parse Portal request requirement")
assert(!portal_request_plan["request_object_created"], "D-Bus client must parse Portal request object state")
assert(!portal_request_plan["permission_granted"], "D-Bus client must parse Portal permission grant state")
assert(!portal_request_plan["host_permission_changed"], "D-Bus client must parse Portal host permission safety")
assert(!portal_request_plan["backend_details_exposed"], "D-Bus client must parse Portal request backend detail status")

state_root = client.state_root("org.xnix.sample.notepad")
assert(state_root["root_type"] == "compatibility-application-state-root", "D-Bus client must parse application state roots")
assert(state_root["snapshot_eligible"], "D-Bus client must parse state root booleans")
assert(!state_root["user_documents_included"], "D-Bus client must parse state root exclusions")

package_source = client.package_source("org.xnix.sample.notepad")
assert(package_source["source_type"] == "compatibility-package-source", "D-Bus client must parse compatibility package sources")
assert(package_source["source_selection_state"] == "planned", "D-Bus client must parse package source state")
assert(!package_source["package_source_ready"], "D-Bus client must parse package source readiness")
assert(!package_source["install_enabled"], "D-Bus client must parse package install status")

acquisition_preflight = client.acquisition_preflight("org.xnix.sample.notepad")
assert(acquisition_preflight["preflight_type"] == "compatibility-acquisition-preflight", "D-Bus client must parse compatibility acquisition preflight")
assert(acquisition_preflight["preflight_state"] == "planned", "D-Bus client must parse acquisition preflight state")
assert(!acquisition_preflight["acquisition_ready"], "D-Bus client must parse acquisition readiness")
assert(!acquisition_preflight["download_enabled"], "D-Bus client must parse download status")

artifact_manifest = client.artifact_manifest("org.xnix.sample.notepad")
assert(artifact_manifest["manifest_type"] == "compatibility-artifact-manifest", "D-Bus client must parse compatibility artifact manifests")
assert(artifact_manifest["manifest_state"] == "planned", "D-Bus client must parse artifact manifest state")
assert(!artifact_manifest["manifest_ready"], "D-Bus client must parse artifact manifest readiness")
assert(!artifact_manifest["signature_verified"], "D-Bus client must parse artifact manifest signature status")
assert(!artifact_manifest["download_enabled"], "D-Bus client must parse artifact download status")

install_plan = client.install_plan("org.xnix.sample.notepad")
assert(install_plan["plan_type"] == "compatibility-install-plan", "D-Bus client must parse compatibility install plans")
assert(install_plan["install_state"] == "planned", "D-Bus client must parse install plan state")
assert(!install_plan["install_ready"], "D-Bus client must parse install readiness")
assert(!install_plan["desktop_activation_ready"], "D-Bus client must parse desktop activation readiness")
assert(!install_plan["download_enabled"], "D-Bus client must parse install plan download status")
assert(!install_plan["install_enabled"], "D-Bus client must parse install plan install status")

backend_binding = client.backend_binding("org.xnix.sample.notepad")
assert(backend_binding["binding_type"] == "compatibility-backend-binding", "D-Bus client must parse backend bindings")
assert(!backend_binding["managed_binding_ready"], "D-Bus client must parse backend binding readiness")
assert(!backend_binding["launch_enabled"], "D-Bus client must parse backend launch status")

repair_plan = client.repair_plan("org.xnix.sample.notepad", "engine-binding-pending")
assert(repair_plan["plan_type"] == "compatibility-repair", "D-Bus client must parse repair plans")
assert(repair_plan["snapshot_required"], "D-Bus client must parse repair plan booleans")

test_plan = client.test_plan("org.xnix.sample.notepad")
assert(test_plan["plan_type"] == "compatibility-test", "D-Bus client must parse test plans")
assert(test_plan["runtime_owned"], "D-Bus client must parse test plan booleans")
assert(!test_plan["backend_details_exposed"], "D-Bus client must parse test plan false booleans")

test_result = client.test_result("org.xnix.sample.notepad")
assert(test_result["result_type"] == "compatibility-test-result", "D-Bus client must parse test results")
assert(test_result["safe_for_ai_diagnostics"], "D-Bus client must parse test result booleans")
assert(!test_result["backend_details_exposed"], "D-Bus client must parse test result false booleans")

ai_input = client.ai_diagnostic_input("org.xnix.sample.notepad")
assert(ai_input["input_type"] == "ai-diagnostic-input", "D-Bus client must parse AI diagnostic inputs")
assert(ai_input["safe_for_ai_diagnostics"], "D-Bus client must parse AI diagnostic input booleans")
assert(!ai_input["ai_provider_called"], "D-Bus client must parse AI provider false booleans")
assert(!ai_input["network_required"], "D-Bus client must parse network false booleans")

ai_recommendation = client.ai_diagnostic_recommendation("org.xnix.sample.notepad")
assert(ai_recommendation["recommendation_type"] == "ai-diagnostic-recommendation", "D-Bus client must parse AI diagnostic recommendations")
assert(ai_recommendation["safe_for_ai_diagnostics"], "D-Bus client must parse AI diagnostic recommendation booleans")
assert(!ai_recommendation["ai_provider_called"], "D-Bus client must parse recommendation AI provider false booleans")
assert(!ai_recommendation["network_required"], "D-Bus client must parse recommendation network false booleans")

ai_gate = client.ai_repair_approval_gate("org.xnix.sample.notepad")
assert(ai_gate["gate_type"] == "ai-repair-approval-gate", "D-Bus client must parse AI repair approval gates")
assert(ai_gate["gate_decision"] == "blocked-until-approval", "D-Bus client must parse AI repair approval gate decisions")
assert(!ai_gate["auto_execution_allowed"], "D-Bus client must parse AI repair gate execution booleans")
assert(!ai_gate["repair_executed"], "D-Bus client must parse AI repair execution false booleans")

snapshot_plan = client.snapshot_plan("org.xnix.sample.notepad", "before-repair")
assert(snapshot_plan["plan_type"] == "compatibility-snapshot", "D-Bus client must parse snapshot plans")
assert(snapshot_plan["enabled_by_default"], "D-Bus client must parse snapshot plan booleans")

portal_policy = client.portal_access_policy("org.xnix.sample.notepad", "file-open")
assert(portal_policy["policy_type"] == "portal-access", "D-Bus client must parse Portal access policies")
assert(portal_policy["portal_required"], "D-Bus client must parse Portal policy booleans")

service_binding = client.runtime_service_binding
assert(service_binding["binding_type"] == "runtime-service-binding", "D-Bus client must parse Runtime service binding")
assert(service_binding["activation_binding_ready"], "D-Bus client must parse Runtime service binding readiness")
assert(!service_binding["live_dbus_owner_ready"], "D-Bus client must parse live D-Bus owner readiness")

live_owner_gate = client.runtime_live_owner_gate
assert(live_owner_gate["gate_type"] == "runtime-live-owner-gate", "D-Bus client must parse Runtime live owner gates")
assert(live_owner_gate["activation_binding_ready"], "D-Bus client must parse Runtime live owner activation readiness")
assert(!live_owner_gate["live_dbus_owner_ready"], "D-Bus client must parse Runtime live owner readiness")
assert(!live_owner_gate["production_owner_enabled"], "D-Bus client must parse Runtime production owner status")
assert(!live_owner_gate["owner_transition_ready"], "D-Bus client must parse Runtime owner transition status")
assert(!live_owner_gate["smoke_adapter_is_production_owner"], "D-Bus client must parse smoke adapter owner status")
assert(!live_owner_gate["kde_may_claim_runtime_ownership"], "D-Bus client must parse KDE ownership status")
assert(!live_owner_gate["backend_details_exposed"], "D-Bus client must parse Runtime live owner backend detail status")

owner_smoke_plan = client.runtime_owner_smoke_plan
assert(owner_smoke_plan["plan_type"] == "runtime-owner-smoke-plan", "D-Bus client must parse Runtime owner smoke plans")
assert(owner_smoke_plan["smoke_state"] == "planned", "D-Bus client must parse Runtime owner smoke state")
assert(owner_smoke_plan["smoke_environment"] == "restricted-session", "D-Bus client must parse Runtime owner smoke environment")
assert(owner_smoke_plan["activation_binding_ready"], "D-Bus client must parse owner smoke activation readiness")
assert(!owner_smoke_plan["live_dbus_owner_ready"], "D-Bus client must parse owner smoke live readiness")
assert(!owner_smoke_plan["production_owner_enabled"], "D-Bus client must parse owner smoke production owner status")
assert(!owner_smoke_plan["owner_transition_ready"], "D-Bus client must parse owner smoke transition status")
assert(owner_smoke_plan["pending_step_count"] == 6, "D-Bus client must parse owner smoke step counts")
assert(!owner_smoke_plan["system_service_started"], "D-Bus client must parse owner smoke service state")
assert(!owner_smoke_plan["production_bus_claimed"], "D-Bus client must parse owner smoke bus claim state")
assert(!owner_smoke_plan["backend_details_exposed"], "D-Bus client must parse owner smoke backend detail status")

method_parity = client.runtime_method_parity_manifest
assert(method_parity["manifest_type"] == "runtime-method-parity-manifest", "D-Bus client must parse Runtime method parity manifests")
assert(method_parity["method_count"] == 32, "D-Bus client must parse Runtime method counts")
assert(method_parity["read_only_method_parity_ready"], "D-Bus client must parse read-only method parity readiness")
assert(method_parity["passed_check_count"] == 5, "D-Bus client must parse Runtime parity pass counts")
assert(method_parity["blocked_check_count"].zero?, "D-Bus client must parse Runtime parity blocked counts")
assert(!method_parity["write_methods_supported"], "D-Bus client must parse write method support status")
assert(!method_parity["write_method_dispatch_enabled"], "D-Bus client must parse write method dispatch status")
assert(!method_parity["backend_details_exposed"], "D-Bus client must parse Runtime parity backend detail status")

write_gate = client.runtime_write_gate("Launch")
assert(write_gate["gate_type"] == "runtime-write-gate", "D-Bus client must parse Runtime write gates")
assert(write_gate["method_name"] == "Launch", "D-Bus client must preserve write gate method names")
assert(write_gate["gate_decision"] == "blocked-until-production-backend", "D-Bus client must parse write gate decisions")
assert(!write_gate["write_method_enabled"], "D-Bus client must parse disabled write method status")
assert(!write_gate["dispatch_enabled"], "D-Bus client must parse disabled dispatch status")
assert(!write_gate["request_object_created"], "D-Bus client must parse request object state")
assert(write_gate["required_gate_count"] == 6, "D-Bus client must parse write gate counts")
assert(write_gate["denial_error_name"] == "org.xnix.Compatibility1.Error.WriteMethodDisabled", "D-Bus client must parse write gate error names")
assert(!write_gate["backend_details_exposed"], "D-Bus client must parse write gate backend detail status")

settings = client.settings("org.xnix.sample.notepad")
assert(settings["request_type"] == "settings-model", "D-Bus client must parse compatibility settings")
assert(settings["settings_state"] == "planned", "D-Bus client must parse settings state")
assert(!settings["settings_persisted"], "D-Bus client must parse settings persistence status")
assert(!settings["host_root_modified"], "D-Bus client must parse settings host mutation status")
assert(!settings["backend_details_exposed"], "D-Bus client must parse settings backend detail status")

settings_change = client.settings_change_plan("org.xnix.sample.notepad", "resource-access", "documents", "ask")
assert(settings_change["plan_type"] == "settings-change-plan", "D-Bus client must parse settings change plans")
assert(settings_change["change_state"] == "planned", "D-Bus client must parse settings change state")
assert(!settings_change["apply_enabled"], "D-Bus client must parse settings change execution status")
assert(!settings_change["settings_persisted"], "D-Bus client must parse settings change persistence status")
assert(settings_change["portal_policy_review_required"], "D-Bus client must parse settings change Portal review status")
assert(!settings_change["snapshot_recommended"], "D-Bus client must parse settings change snapshot status")
assert(!settings_change["backend_details_exposed"], "D-Bus client must parse settings change backend detail status")

action_queue = client.action_queue("org.xnix.sample.notepad")
assert(action_queue["queue_type"] == "compatibility-center-action-queue", "D-Bus client must parse Compatibility Center action queues")
assert(action_queue["action_count"] == 5, "D-Bus client must parse action queue counts")
assert(action_queue["pending_action_count"] == 5, "D-Bus client must parse pending action counts")
assert(action_queue["user_review_required_count"] == 3, "D-Bus client must parse review action counts")
assert(!action_queue["execution_enabled"], "D-Bus client must parse action queue execution status")
assert(!action_queue["repair_execution_enabled"], "D-Bus client must parse repair execution status")
assert(!action_queue["settings_persistence_enabled"], "D-Bus client must parse settings persistence status")
assert(!action_queue["backend_details_exposed"], "D-Bus client must parse action queue backend detail status")

action_review = client.action_review_receipt("org.xnix.sample.notepad", "review-ai-repair", "approved")
assert(action_review["receipt_type"] == "compatibility-center-action-review-receipt", "D-Bus client must parse action review receipts")
assert(action_review["action_id"] == "review-ai-repair", "D-Bus client must parse action review ids")
assert(action_review["decision"] == "approved", "D-Bus client must parse action review decisions")
assert(action_review["decision_recorded"], "D-Bus client must parse action review recording status")
assert(!action_review["execution_enabled"], "D-Bus client must parse action review execution status")
assert(!action_review["repair_execution_enabled"], "D-Bus client must parse action review repair execution status")
assert(!action_review["settings_persistence_enabled"], "D-Bus client must parse action review settings persistence status")
assert(!action_review["resource_grant_created"], "D-Bus client must parse action review resource grant status")
assert(!action_review["backend_details_exposed"], "D-Bus client must parse action review backend detail status")

assert(
  capture.commands.all? { |command| command.include?("--session") },
  "D-Bus client must use the session bus for KDE-facing reads"
)

puts "PASS: compatibility runtime D-Bus client unit tests"
