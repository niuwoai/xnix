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
    when "org.xnix.Compatibility1.GetKDEIntegrationStatus"
      [
        "({'status_type': <'kde-integration-status'>, 'desktop': <'KDE Plasma'>, 'entry_point_count': <7>, 'initial_count': <7>, 'planned_count': <0>, 'complete_count': <0>, 'entry_point_ids': <['launcher', 'task-manager', 'file-manager', 'system-tray', 'notifications', 'compatibility-center', 'settings']>, 'entry_point_names': <['Launcher', 'Task Manager', 'File Manager', 'System Tray', 'Notifications', 'Compatibility Center', 'Settings']>, 'entry_point_states': <['initial', 'initial', 'initial', 'initial', 'initial', 'initial', 'initial']>, 'runtime_methods': <['GetDesktopEntryPlan', 'GetTaskManagerIdentityPlan', 'GetFileAssociationPlan', 'GetTrayStatus', 'GetNotificationPlan', 'GetCompatibilityCenterSummary', 'GetCompatibilitySettings']>, 'runtime_owned': <true>, 'kde_policy_owner': <false>, 'official_desktop_only': <true>, 'stable_desktop_contract': <true>, 'host_root_modified': <false>, 'backend_details_exposed': <false>},)\n",
        "",
        Status.new(true)
      ]
    when "org.xnix.Compatibility1.GetKDEApplicationSurfacePlan"
      [
        "({'plan_type': <'kde-application-surface-plan'>, 'runtime_method': <'GetKDEApplicationSurfacePlan'>, 'application_id': <'org.xnix.sample.notepad'>, 'surface_state': <'planned'>, 'desktop_shell': <'KDE Plasma'>, 'entry_point_count': <7>, 'entry_point_ids': <['launcher', 'task-manager', 'file-manager', 'system-tray', 'notifications', 'compatibility-center', 'settings']>, 'required_runtime_gates': <['recipe-install-gate', 'portal-policy-review', 'snapshot-baseline', 'backend-environment-plan', 'backend-lifecycle-plan', 'runtime-write-gate']>, 'runtime_owned': <true>, 'c_runtime_backed': <true>, 'kde_policy_owner': <false>, 'official_desktop_only': <true>, 'normal_linux_application_surface': <true>, 'standard_launcher_visible': <true>, 'launch_enabled': <false>, 'backend_process_started': <false>, 'desktop_files_written': <false>, 'mimeapps_written': <false>, 'host_root_modified': <false>, 'backend_command_exposed': <false>, 'raw_windows_executable_exposed': <false>, 'backend_details_exposed': <false>},)\n",
        "",
        Status.new(true)
      ]
    when "org.xnix.Compatibility1.GetDesktopResourceBridgePlan"
      [
        "({'plan_type': <'desktop-resource-bridge-plan'>, 'runtime_method': <'GetDesktopResourceBridgePlan'>, 'application_id': <'org.xnix.sample.notepad'>, 'bridge_state': <'planned'>, 'resource_count': <5>, 'resource_ids': <['file-open', 'uri-open', 'print', 'clipboard', 'screenshot']>, 'portal_mediated': <true>, 'file_bridge_planned': <true>, 'print_bridge_planned': <true>, 'clipboard_bridge_planned': <true>, 'bridges_enabled': <false>, 'requests_created': <false>, 'backend_process_started': <false>, 'direct_host_file_access': <false>, 'direct_clipboard_access': <false>, 'direct_print_access': <false>, 'host_root_modified': <false>, 'backend_details_exposed': <false>},)\n",
        "",
        Status.new(true)
      ]
    when "org.xnix.Compatibility1.GetCompatibilityModeSwitchPlan"
      [
        "({'plan_type': <'compatibility-mode-switch-plan'>, 'runtime_method': <'GetCompatibilityModeSwitchPlan'>, 'application_id': <'org.xnix.sample.notepad'>, 'current_mode': <'automatic'>, 'requested_mode': <'prefer-compatibility'>, 'mode_state': <'planned'>, 'mode_count': <4>, 'mode_ids': <['automatic', 'prefer-performance', 'prefer-compatibility', 'isolated-execution']>, 'runtime_owned': <true>, 'c_runtime_backed': <true>, 'kde_policy_owner': <false>, 'valid_mode': <true>, 'requires_user_confirmation': <true>, 'portal_review_required': <true>, 'snapshot_required': <true>, 'settings_persistence_enabled': <false>, 'backend_reconfiguration_enabled': <false>, 'backend_process_started': <false>, 'launch_enabled': <false>, 'host_root_modified': <false>, 'backend_details_exposed': <false>},)\n",
        "",
        Status.new(true)
      ]
    when "org.xnix.Compatibility1.GetCompatibilityPermissionReviewPlan"
      [
        "({'plan_type': <'compatibility-permission-review-plan'>, 'runtime_method': <'GetCompatibilityPermissionReviewPlan'>, 'application_id': <'org.xnix.sample.notepad'>, 'review_state': <'planned'>, 'permission_count': <7>, 'allow_count': <1>, 'ask_count': <5>, 'deny_count': <1>, 'permission_ids': <['documents', 'downloads', 'camera', 'network', 'clipboard', 'print', 'screenshot']>, 'runtime_owned': <true>, 'c_runtime_backed': <true>, 'kde_policy_owner': <false>, 'user_review_required': <true>, 'portal_review_required': <true>, 'permission_changes_applied': <false>, 'request_objects_created': <false>, 'permissions_granted': <false>, 'settings_persisted': <false>, 'host_permission_changed': <false>, 'host_root_modified': <false>, 'backend_details_exposed': <false>},)\n",
        "",
        Status.new(true)
      ]
    when "org.xnix.Compatibility1.GetCompatibilityReviewFlowPlan"
      [
        "({'plan_type': <'compatibility-review-flow-plan'>, 'runtime_method': <'GetCompatibilityReviewFlowPlan'>, 'application_id': <'org.xnix.sample.notepad'>, 'review_state': <'planned'>, 'section_id': <'resource-access'>, 'field_id': <'documents'>, 'requested_value': <'ask'>, 'operation': <'file-open'>, 'step_count': <5>, 'required_review_count': <3>, 'blocked_step_count': <1>, 'pending_step_count': <1>, 'step_ids': <['settings-change-review', 'permission-review', 'portal-request-review', 'runtime-write-gate', 'review-receipt']>, 'runtime_owned': <true>, 'c_runtime_backed': <true>, 'kde_policy_owner': <false>, 'user_confirmation_required': <true>, 'portal_policy_review_required': <true>, 'settings_change_planned': <true>, 'permission_review_planned': <true>, 'portal_request_planned': <true>, 'review_receipt_required': <true>, 'apply_enabled': <false>, 'request_object_created': <false>, 'permission_granted': <false>, 'settings_persisted': <false>, 'execution_started': <false>, 'host_root_modified': <false>, 'backend_details_exposed': <false>},)\n",
        "",
        Status.new(true)
      ]
    when "org.xnix.Compatibility1.GetFileAssociationPlan"
      [
        "({'plan_type': <'file-association-plan'>, 'association_type': <'desktop-file-association'>, 'desktop_file': <'xnix-org.xnix.sample.notepad.desktop'>, 'mimeapps_path': <'usr/share/applications/mimeapps.list'>, 'file_open_command': <'xnix-compat-open'>, 'file_open_argument': <'%U'>, 'mime_type_count': <2>, 'standard_mimeapps_list': <true>, 'staged_root_only': <true>, 'overwrite_existing_mimeapps': <false>, 'portal_required_for_file_open': <true>, 'files_written': <false>, 'host_root_modified': <false>, 'backend_details_exposed': <false>},)\n",
        "",
        Status.new(true)
      ]
    when "org.xnix.Compatibility1.GetNotificationPlan"
      [
        "({'plan_type': <'notification-plan'>, 'event_type': <'approval-required'>, 'notification_id': <'org.xnix.sample.notepad.approval-required'>, 'title': <'Sample Notepad needs approval'>, 'urgency': <'critical'>, 'category': <'compatibility.approval'>, 'desktop_entry': <'xnix-org.xnix.sample.notepad.desktop'>, 'action_count': <2>, 'runtime_owned': <true>, 'kde_policy_owner': <false>, 'user_visible': <true>, 'requires_user_review': <true>, 'action_execution_enabled': <false>, 'repair_execution_enabled': <false>, 'settings_persistence_enabled': <false>, 'host_root_modified': <false>, 'backend_details_exposed': <false>},)\n",
        "",
        Status.new(true)
      ]
    when "org.xnix.Compatibility1.GetTrayStatus"
      [
        "({'status_type': <'tray-status-plan'>, 'desktop': <'KDE Plasma'>, 'active_application_count': <1>, 'attention_required_count': <1>, 'bridged_tray_application_count': <0>, 'compatibility_state': <'attention-required'>, 'tray_bridge_state': <'planned'>, 'runtime_owned': <true>, 'kde_policy_owner': <false>, 'user_visible': <true>, 'live_backend_bridge_enabled': <false>, 'bridge_configuration_persisted': <false>, 'host_root_modified': <false>, 'backend_details_exposed': <false>},)\n",
        "",
        Status.new(true)
      ]
    when "org.xnix.Compatibility1.GetKRunnerQueryPlan"
      [
        "({'query_type': <'krunner-query-plan'>, 'entry_point': <'krunner'>, 'desktop': <'KDE Plasma'>, 'query': <'notepad'>, 'match_count': <1>, 'top_application_id': <'org.xnix.sample.notepad'>, 'top_name': <'Sample Notepad'>, 'top_relevance_percent': <100>, 'action_type': <'runtime-launch'>, 'desktop_entry_id': <'org.xnix.sample.notepad.desktop'>, 'runtime_owned': <true>, 'kde_policy_owner': <false>, 'runtime_owned_launch': <true>, 'query_execution_enabled': <false>, 'backend_launch_enabled': <false>, 'host_root_modified': <false>, 'backend_details_exposed': <false>},)\n",
        "",
        Status.new(true)
      ]
    when "org.xnix.Compatibility1.GetKWinWindowRulePlan"
      [
        "({'request_type': <'kwin-window-rule'>, 'desktop': <'KDE Plasma'>, 'application_id': <'org.xnix.sample.notepad'>, 'name': <'Sample Notepad'>, 'script_role': <'identity-and-layout'>, 'resource_name': <'org.xnix.sample.notepad'>, 'class_group': <'xnix-compatibility'>, 'title_hint': <'Sample Notepad'>, 'desktop_file': <'xnix-org.xnix.sample.notepad.desktop'>, 'task_manager_grouping_key': <'org.xnix.sample.notepad'>, 'launcher_url': <'applications:xnix-org.xnix.sample.notepad.desktop'>, 'skip_taskbar': <false>, 'show_in_switcher': <true>, 'placement': <'normal-window'>, 'pinning_allowed': <true>, 'restore_allowed': <true>, 'restore_key': <'org.xnix.sample.notepad'>, 'prefer_existing_window': <true>, 'window_manager_policy_only': <true>, 'runtime_owns_backend_policy': <true>, 'runtime_owned': <true>, 'kde_policy_owner': <false>, 'host_root_modified': <false>, 'backend_details_exposed': <false>},)\n",
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
    when "org.xnix.Compatibility1.GetBackendLifecycle"
      [
        "({'lifecycle_type': <'compatibility-backend-lifecycle'>, 'runtime_method': <'GetBackendLifecycle'>, 'lifecycle_state': <'blocked'>, 'overall_status': <'not-ready'>, 'runtime_owned': <true>, 'c_runtime_backed': <true>, 'kde_policy_owner': <false>, 'backend_binding_ready': <false>, 'launch_enabled': <false>, 'execution_request_created': <false>, 'backend_process_started': <false>, 'local_backend_started': <false>, 'isolated_backend_started': <false>, 'host_root_modified': <false>, 'network_required': <false>, 'backend_details_exposed': <false>},)\n",
        "",
        Status.new(true)
      ]
    when "org.xnix.Compatibility1.GetBackendEnvironmentPlan"
      [
        "({'plan_type': <'compatibility-backend-environment-plan'>, 'runtime_method': <'GetBackendEnvironmentPlan'>, 'environment_state': <'planned'>, 'runtime_owned': <true>, 'c_runtime_backed': <true>, 'kde_policy_owner': <false>, 'local_environment_ready': <false>, 'isolated_environment_ready': <false>, 'environment_created': <false>, 'backend_process_started': <false>, 'host_storage_exposed': <false>, 'clipboard_bridge_enabled': <false>, 'print_bridge_enabled': <false>, 'launch_enabled': <false>, 'host_root_modified': <false>, 'network_required': <false>, 'backend_details_exposed': <false>},)\n",
        "",
        Status.new(true)
      ]
    when "org.xnix.Compatibility1.GetRepairPlan"
      [
        "({'plan_type': <'compatibility-repair'>, 'issue': <'engine-binding-pending'>, 'runtime_method': <'GetRepairPlan'>, 'runtime_owned': <true>, 'c_runtime_backed': <true>, 'kde_policy_owner': <false>, 'severity': <'warning'>, 'user_approval_required': <true>, 'snapshot_required': <true>, 'rollback_available': <true>, 'repair_execution_requested': <false>, 'repair_executed': <false>, 'backend_launch_enabled': <false>, 'network_required': <false>, 'host_root_modified': <false>, 'backend_details_exposed': <false>},)\n",
        "",
        Status.new(true)
      ]
    when "org.xnix.Compatibility1.GetTestPlan"
      [
        "({'plan_type': <'compatibility-test'>, 'test_type': <'preflight'>, 'runtime_method': <'GetTestPlan'>, 'runtime_owned': <true>, 'c_runtime_backed': <true>, 'kde_policy_owner': <false>, 'execution_request_created': <false>, 'test_executed': <false>, 'host_root_modified': <false>, 'backend_details_exposed': <false>},)\n",
        "",
        Status.new(true)
      ]
    when "org.xnix.Compatibility1.GetTestResult"
      [
        "({'result_type': <'compatibility-test-result'>, 'test_type': <'preflight'>, 'runtime_method': <'GetTestResult'>, 'runtime_owned': <true>, 'c_runtime_backed': <true>, 'kde_policy_owner': <false>, 'result_source': <'runtime-model'>, 'execution_state': <'waiting-for-runtime'>, 'overall_status': <'pending'>, 'safe_for_ai_diagnostics': <true>, 'test_executed': <false>, 'host_root_modified': <false>, 'backend_details_exposed': <false>},)\n",
        "",
        Status.new(true)
      ]
    when "org.xnix.Compatibility1.GetExecutionReadiness"
      [
        "({'readiness_type': <'compatibility-execution-readiness'>, 'runtime_method': <'GetExecutionReadiness'>, 'runtime_owned': <true>, 'c_runtime_backed': <true>, 'kde_policy_owner': <false>, 'execution_state': <'blocked'>, 'overall_status': <'not-ready'>, 'launch_allowed': <false>, 'launch_enabled': <false>, 'execution_request_created': <false>, 'backend_binding_ready': <false>, 'desktop_entry_launch_visible': <true>, 'safe_for_ai_diagnostics': <true>, 'host_root_modified': <false>, 'backend_details_exposed': <false>},)\n",
        "",
        Status.new(true)
      ]
    when "org.xnix.Compatibility1.GetLaunchIntent"
      [
        "({'intent_type': <'runtime-launch-intent'>, 'source': <'desktop-launcher'>, 'runtime_method': <'Launch'>, 'read_method': <'GetLaunchIntent'>, 'runtime_owned': <true>, 'c_runtime_backed': <true>, 'kde_policy_owner': <false>, 'standard_desktop_entry': <true>, 'launch_uses_runtime': <true>, 'desktop_entry_launch_visible': <true>, 'launch_allowed': <false>, 'launch_enabled': <false>, 'execution_request_created': <false>, 'execution_started': <false>, 'write_gate_decision': <'blocked-until-production-backend'>, 'host_root_modified': <false>, 'network_required': <false>, 'backend_details_exposed': <false>},)\n",
        "",
        Status.new(true)
      ]
    when "org.xnix.Compatibility1.GetAIDiagnosticInput"
      [
        "({'input_type': <'ai-diagnostic-input'>, 'runtime_method': <'GetAIDiagnosticInput'>, 'c_runtime_backed': <true>, 'diagnostic_signal_count': <3>, 'safe_for_ai_diagnostics': <true>, 'ai_provider_called': <false>, 'network_required': <false>, 'backend_details_exposed': <false>},)\n",
        "",
        Status.new(true)
      ]
    when "org.xnix.Compatibility1.GetAIDiagnosticRecommendation"
      [
        "({'recommendation_type': <'ai-diagnostic-recommendation'>, 'runtime_method': <'GetAIDiagnosticRecommendation'>, 'c_runtime_backed': <true>, 'recommendation_count': <3>, 'approval_required_count': <1>, 'safe_for_ai_diagnostics': <true>, 'ai_provider_called': <false>, 'network_required': <false>, 'auto_execution_allowed': <false>, 'backend_details_exposed': <false>},)\n",
        "",
        Status.new(true)
      ]
    when "org.xnix.Compatibility1.GetAIRepairApprovalGate"
      [
        "({'gate_type': <'ai-repair-approval-gate'>, 'runtime_method': <'GetAIRepairApprovalGate'>, 'c_runtime_backed': <true>, 'gate_decision': <'blocked-until-approval'>, 'required_gate_count': <3>, 'approval_required_count': <1>, 'safe_for_ai_diagnostics': <true>, 'ai_provider_called': <false>, 'network_required': <false>, 'auto_execution_allowed': <false>, 'repair_executed': <false>, 'backend_details_exposed': <false>},)\n",
        "",
        Status.new(true)
      ]
    when "org.xnix.Compatibility1.GetSnapshotPlan"
      [
        "({'plan_type': <'compatibility-snapshot'>, 'reason': <'before-repair'>, 'runtime_method': <'GetSnapshotPlan'>, 'runtime_owned': <true>, 'c_runtime_backed': <true>, 'kde_policy_owner': <false>, 'enabled_by_default': <true>, 'snapshot_request_created': <false>, 'snapshot_created': <false>, 'restore_requested': <false>, 'restore_executed': <false>, 'user_documents_included': <false>, 'host_system_included': <false>, 'host_root_modified': <false>, 'backend_details_exposed': <false>},)\n",
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
        "({'manifest_type': <'runtime-method-parity-manifest'>, 'method_count': <48>, 'read_only_method_parity_ready': <true>, 'passed_check_count': <5>, 'blocked_check_count': <0>, 'write_methods_supported': <false>, 'write_method_dispatch_enabled': <false>, 'backend_details_exposed': <false>},)\n",
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
    when "org.xnix.Compatibility1.GetCompatibilityCenterSummary"
      [
        "({'summary_type': <'compatibility-center-summary'>, 'compatibility_state': <'review-required'>, 'runtime_mode': <'automatic'>, 'known_issue_count': <1>, 'repair_record_state': <'pending-review'>, 'last_repair_event': <'approval-required'>, 'action_count': <4>, 'runtime_owned': <true>, 'kde_policy_owner': <false>, 'user_visible': <true>, 'action_execution_enabled': <false>, 'repair_execution_enabled': <false>, 'backend_launch_enabled': <false>, 'settings_persistence_enabled': <false>, 'host_root_modified': <false>, 'backend_details_exposed': <false>},)\n",
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

kde_status = client.kde_integration_status
assert(kde_status["status_type"] == "kde-integration-status", "D-Bus client must parse KDE integration status")
assert(kde_status["desktop"] == "KDE Plasma", "D-Bus client must parse KDE integration desktops")
assert(kde_status["entry_point_count"] == 7, "D-Bus client must parse KDE entry point counts")
assert(kde_status["entry_point_ids"].include?("launcher"), "D-Bus client must parse KDE entry point ids")
assert(kde_status["runtime_methods"].include?("GetDesktopEntryPlan"), "D-Bus client must parse KDE Runtime methods")
assert(kde_status["runtime_owned"], "D-Bus client must parse KDE integration Runtime ownership")
assert(!kde_status["kde_policy_owner"], "D-Bus client must parse KDE integration policy ownership")
assert(kde_status["official_desktop_only"], "D-Bus client must parse official desktop scope")
assert(kde_status["stable_desktop_contract"], "D-Bus client must parse stable desktop contract status")
assert(!kde_status["backend_details_exposed"], "D-Bus client must parse KDE integration backend detail status")

file_association_plan = client.file_association_plan("org.xnix.sample.notepad")
assert(file_association_plan["plan_type"] == "file-association-plan", "D-Bus client must parse file association plans")
assert(file_association_plan["association_type"] == "desktop-file-association", "D-Bus client must parse file association types")
assert(file_association_plan["desktop_file"] == "xnix-org.xnix.sample.notepad.desktop", "D-Bus client must parse file association desktop files")
assert(file_association_plan["mimeapps_path"] == "usr/share/applications/mimeapps.list", "D-Bus client must parse mimeapps paths")
assert(file_association_plan["file_open_command"] == "xnix-compat-open", "D-Bus client must parse file open commands")
assert(file_association_plan["file_open_argument"] == "%U", "D-Bus client must parse file open arguments")
assert(file_association_plan["mime_type_count"] == 2, "D-Bus client must parse file association counts")
assert(file_association_plan["standard_mimeapps_list"], "D-Bus client must parse standard mimeapps status")
assert(file_association_plan["staged_root_only"], "D-Bus client must parse staged-root status")
assert(!file_association_plan["overwrite_existing_mimeapps"], "D-Bus client must parse overwrite safety")
assert(file_association_plan["portal_required_for_file_open"], "D-Bus client must parse file-open Portal requirement")
assert(!file_association_plan["files_written"], "D-Bus client must parse file association write safety")
assert(!file_association_plan["host_root_modified"], "D-Bus client must parse file association host-root safety")
assert(!file_association_plan["backend_details_exposed"], "D-Bus client must parse file association backend detail status")

notification_plan = client.notification_plan("org.xnix.sample.notepad", "approval-required")
assert(notification_plan["plan_type"] == "notification-plan", "D-Bus client must parse notification plans")
assert(notification_plan["event_type"] == "approval-required", "D-Bus client must parse notification event types")
assert(notification_plan["notification_id"] == "org.xnix.sample.notepad.approval-required", "D-Bus client must parse notification ids")
assert(notification_plan["urgency"] == "critical", "D-Bus client must parse notification urgency")
assert(notification_plan["category"] == "compatibility.approval", "D-Bus client must parse notification categories")
assert(notification_plan["desktop_entry"] == "xnix-org.xnix.sample.notepad.desktop", "D-Bus client must parse notification desktop entries")
assert(notification_plan["action_count"] == 2, "D-Bus client must parse notification action counts")
assert(notification_plan["runtime_owned"], "D-Bus client must parse notification Runtime ownership")
assert(!notification_plan["kde_policy_owner"], "D-Bus client must parse notification KDE policy ownership")
assert(notification_plan["user_visible"], "D-Bus client must parse notification visibility")
assert(notification_plan["requires_user_review"], "D-Bus client must parse notification review requirements")
assert(!notification_plan["action_execution_enabled"], "D-Bus client must parse notification action gate status")
assert(!notification_plan["repair_execution_enabled"], "D-Bus client must parse notification repair gate status")
assert(!notification_plan["settings_persistence_enabled"], "D-Bus client must parse notification settings gate status")
assert(!notification_plan["host_root_modified"], "D-Bus client must parse notification host-root safety")
assert(!notification_plan["backend_details_exposed"], "D-Bus client must parse notification backend detail status")

tray_status = client.tray_status
assert(tray_status["status_type"] == "tray-status-plan", "D-Bus client must parse tray status plans")
assert(tray_status["desktop"] == "KDE Plasma", "D-Bus client must parse tray desktops")
assert(tray_status["active_application_count"] == 1, "D-Bus client must parse tray active application counts")
assert(tray_status["attention_required_count"] == 1, "D-Bus client must parse tray attention counts")
assert(tray_status["bridged_tray_application_count"] == 0, "D-Bus client must parse bridged tray counts")
assert(tray_status["compatibility_state"] == "attention-required", "D-Bus client must parse compatibility tray states")
assert(tray_status["tray_bridge_state"] == "planned", "D-Bus client must parse tray bridge states")
assert(tray_status["runtime_owned"], "D-Bus client must parse tray Runtime ownership")
assert(!tray_status["kde_policy_owner"], "D-Bus client must parse tray KDE policy ownership")
assert(tray_status["user_visible"], "D-Bus client must parse tray visibility")
assert(!tray_status["live_backend_bridge_enabled"], "D-Bus client must parse live tray bridge gates")
assert(!tray_status["bridge_configuration_persisted"], "D-Bus client must parse tray persistence gates")
assert(!tray_status["host_root_modified"], "D-Bus client must parse tray host-root safety")
assert(!tray_status["backend_details_exposed"], "D-Bus client must parse tray backend detail status")

krunner_query = client.krunner_query_plan("notepad")
assert(krunner_query["query_type"] == "krunner-query-plan", "D-Bus client must parse KRunner query plans")
assert(krunner_query["entry_point"] == "krunner", "D-Bus client must parse KRunner entry points")
assert(krunner_query["match_count"] == 1, "D-Bus client must parse KRunner match counts")
assert(krunner_query["top_application_id"] == "org.xnix.sample.notepad", "D-Bus client must parse KRunner top matches")
assert(krunner_query["action_type"] == "runtime-launch", "D-Bus client must parse KRunner action types")
assert(krunner_query["runtime_owned_launch"], "D-Bus client must parse KRunner Runtime launch ownership")
assert(!krunner_query["query_execution_enabled"], "D-Bus client must parse KRunner execution gates")
assert(!krunner_query["backend_launch_enabled"], "D-Bus client must parse KRunner backend launch gates")
assert(!krunner_query["backend_details_exposed"], "D-Bus client must parse KRunner backend detail status")

kwin_rule = client.kwin_window_rule_plan("org.xnix.sample.notepad")
assert(kwin_rule["request_type"] == "kwin-window-rule", "D-Bus client must parse KWin window rule plans")
assert(kwin_rule["desktop"] == "KDE Plasma", "D-Bus client must parse KWin desktops")
assert(kwin_rule["application_id"] == "org.xnix.sample.notepad", "D-Bus client must parse KWin application ids")
assert(kwin_rule["script_role"] == "identity-and-layout", "D-Bus client must parse KWin script roles")
assert(kwin_rule["resource_name"] == "org.xnix.sample.notepad", "D-Bus client must parse KWin resource names")
assert(kwin_rule["class_group"] == "xnix-compatibility", "D-Bus client must parse KWin class groups")
assert(kwin_rule["desktop_file"] == "xnix-org.xnix.sample.notepad.desktop", "D-Bus client must parse KWin desktop files")
assert(!kwin_rule["skip_taskbar"], "D-Bus client must parse KWin taskbar visibility")
assert(kwin_rule["show_in_switcher"], "D-Bus client must parse KWin switcher visibility")
assert(kwin_rule["pinning_allowed"], "D-Bus client must parse KWin pinning")
assert(kwin_rule["restore_allowed"], "D-Bus client must parse KWin restore")
assert(kwin_rule["window_manager_policy_only"], "D-Bus client must parse KWin policy scope")
assert(kwin_rule["runtime_owns_backend_policy"], "D-Bus client must parse KWin Runtime ownership")
assert(!kwin_rule["backend_details_exposed"], "D-Bus client must parse KWin backend detail status")

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

backend_lifecycle = client.backend_lifecycle("org.xnix.sample.notepad")
assert(backend_lifecycle["lifecycle_type"] == "compatibility-backend-lifecycle", "D-Bus client must parse backend lifecycle")
assert(backend_lifecycle["runtime_method"] == "GetBackendLifecycle", "D-Bus client must parse backend lifecycle Runtime methods")
assert(!backend_lifecycle["backend_process_started"], "D-Bus client must parse backend process start state")
assert(!backend_lifecycle["launch_enabled"], "D-Bus client must parse backend lifecycle launch status")

backend_environment_plan = client.backend_environment_plan("org.xnix.sample.notepad")
assert(backend_environment_plan["plan_type"] == "compatibility-backend-environment-plan", "D-Bus client must parse backend environment plans")
assert(backend_environment_plan["runtime_method"] == "GetBackendEnvironmentPlan", "D-Bus client must parse backend environment Runtime methods")
assert(!backend_environment_plan["environment_created"], "D-Bus client must parse backend environment creation state")
assert(!backend_environment_plan["backend_process_started"], "D-Bus client must parse backend environment process state")

kde_application_surface_plan = client.kde_application_surface_plan("org.xnix.sample.notepad")
assert(kde_application_surface_plan["plan_type"] == "kde-application-surface-plan", "D-Bus client must parse KDE application surface plans")
assert(kde_application_surface_plan["runtime_method"] == "GetKDEApplicationSurfacePlan", "D-Bus client must parse KDE application surface Runtime methods")
assert(kde_application_surface_plan["normal_linux_application_surface"], "D-Bus client must parse normal Linux app surface state")
assert(kde_application_surface_plan["standard_launcher_visible"], "D-Bus client must parse standard launcher visibility")
assert(!kde_application_surface_plan["backend_process_started"], "D-Bus client must parse KDE surface backend process state")
assert(!kde_application_surface_plan["backend_details_exposed"], "D-Bus client must parse KDE surface backend detail gates")

desktop_resource_bridge_plan = client.desktop_resource_bridge_plan("org.xnix.sample.notepad")
assert(desktop_resource_bridge_plan["plan_type"] == "desktop-resource-bridge-plan", "D-Bus client must parse desktop resource bridge plans")
assert(desktop_resource_bridge_plan["runtime_method"] == "GetDesktopResourceBridgePlan", "D-Bus client must parse desktop bridge Runtime methods")
assert(desktop_resource_bridge_plan["portal_mediated"], "D-Bus client must parse Portal mediation")
assert(desktop_resource_bridge_plan["file_bridge_planned"], "D-Bus client must parse file bridge planning")
assert(desktop_resource_bridge_plan["print_bridge_planned"], "D-Bus client must parse print bridge planning")
assert(desktop_resource_bridge_plan["clipboard_bridge_planned"], "D-Bus client must parse clipboard bridge planning")
assert(!desktop_resource_bridge_plan["bridges_enabled"], "D-Bus client must parse disabled bridges")
assert(!desktop_resource_bridge_plan["direct_host_file_access"], "D-Bus client must parse host file access gates")
assert(!desktop_resource_bridge_plan["backend_details_exposed"], "D-Bus client must parse desktop bridge backend detail gates")

mode_switch_plan = client.compatibility_mode_switch_plan("org.xnix.sample.notepad", "prefer-compatibility")
assert(mode_switch_plan["plan_type"] == "compatibility-mode-switch-plan", "D-Bus client must parse compatibility mode switch plans")
assert(mode_switch_plan["runtime_method"] == "GetCompatibilityModeSwitchPlan", "D-Bus client must parse compatibility mode switch Runtime methods")
assert(mode_switch_plan["requested_mode"] == "prefer-compatibility", "D-Bus client must parse compatibility mode switch requested modes")
assert(mode_switch_plan["mode_count"] == 4, "D-Bus client must parse compatibility mode counts")
assert(mode_switch_plan["requires_user_confirmation"], "D-Bus client must parse compatibility mode confirmation gates")
assert(!mode_switch_plan["settings_persistence_enabled"], "D-Bus client must parse compatibility mode persistence gates")
assert(!mode_switch_plan["backend_reconfiguration_enabled"], "D-Bus client must parse compatibility mode reconfiguration gates")
assert(!mode_switch_plan["backend_process_started"], "D-Bus client must parse compatibility mode backend process gates")
assert(!mode_switch_plan["backend_details_exposed"], "D-Bus client must parse compatibility mode backend detail gates")

permission_review_plan = client.compatibility_permission_review_plan("org.xnix.sample.notepad")
assert(permission_review_plan["plan_type"] == "compatibility-permission-review-plan", "D-Bus client must parse compatibility permission review plans")
assert(permission_review_plan["runtime_method"] == "GetCompatibilityPermissionReviewPlan", "D-Bus client must parse compatibility permission review Runtime methods")
assert(permission_review_plan["permission_count"] == 7, "D-Bus client must parse compatibility permission counts")
assert(permission_review_plan["allow_count"] == 1, "D-Bus client must parse compatibility permission allow counts")
assert(permission_review_plan["ask_count"] == 5, "D-Bus client must parse compatibility permission ask counts")
assert(permission_review_plan["deny_count"] == 1, "D-Bus client must parse compatibility permission deny counts")
assert(!permission_review_plan["permissions_granted"], "D-Bus client must parse compatibility permission grant gates")
assert(!permission_review_plan["settings_persisted"], "D-Bus client must parse compatibility permission persistence gates")
assert(!permission_review_plan["host_permission_changed"], "D-Bus client must parse compatibility permission host gates")
assert(!permission_review_plan["backend_details_exposed"], "D-Bus client must parse compatibility permission backend detail gates")

review_flow_plan = client.compatibility_review_flow_plan("org.xnix.sample.notepad", "resource-access", "documents", "ask", "file-open")
assert(review_flow_plan["plan_type"] == "compatibility-review-flow-plan", "D-Bus client must parse compatibility review flow plans")
assert(review_flow_plan["runtime_method"] == "GetCompatibilityReviewFlowPlan", "D-Bus client must parse compatibility review flow Runtime methods")
assert(review_flow_plan["step_count"] == 5, "D-Bus client must parse compatibility review flow step counts")
assert(review_flow_plan["required_review_count"] == 3, "D-Bus client must parse compatibility review required counts")
assert(review_flow_plan["blocked_step_count"] == 1, "D-Bus client must parse compatibility review blocked counts")
assert(review_flow_plan["pending_step_count"] == 1, "D-Bus client must parse compatibility review pending counts")
assert(review_flow_plan["user_confirmation_required"], "D-Bus client must parse compatibility review confirmation gates")
assert(review_flow_plan["portal_policy_review_required"], "D-Bus client must parse compatibility review Portal gates")
assert(!review_flow_plan["apply_enabled"], "D-Bus client must parse compatibility review apply gates")
assert(!review_flow_plan["request_object_created"], "D-Bus client must parse compatibility review request gates")
assert(!review_flow_plan["permission_granted"], "D-Bus client must parse compatibility review permission gates")
assert(!review_flow_plan["settings_persisted"], "D-Bus client must parse compatibility review persistence gates")
assert(!review_flow_plan["execution_started"], "D-Bus client must parse compatibility review execution gates")
assert(!review_flow_plan["host_root_modified"], "D-Bus client must parse compatibility review host gates")
assert(!review_flow_plan["backend_details_exposed"], "D-Bus client must parse compatibility review backend detail gates")

repair_plan = client.repair_plan("org.xnix.sample.notepad", "engine-binding-pending")
assert(repair_plan["plan_type"] == "compatibility-repair", "D-Bus client must parse repair plans")
assert(repair_plan["runtime_method"] == "GetRepairPlan", "D-Bus client must parse repair plan Runtime methods")
assert(repair_plan["runtime_owned"], "D-Bus client must parse repair plan Runtime ownership")
assert(repair_plan["c_runtime_backed"], "D-Bus client must parse repair plan C backing")
assert(!repair_plan["kde_policy_owner"], "D-Bus client must parse repair plan KDE policy false booleans")
assert(repair_plan["snapshot_required"], "D-Bus client must parse repair plan booleans")
assert(repair_plan["rollback_available"], "D-Bus client must parse repair rollback availability")
assert(!repair_plan["repair_execution_requested"], "D-Bus client must parse repair request false booleans")
assert(!repair_plan["repair_executed"], "D-Bus client must parse repair execution false booleans")
assert(!repair_plan["backend_launch_enabled"], "D-Bus client must parse repair backend launch false booleans")
assert(!repair_plan["network_required"], "D-Bus client must parse repair network false booleans")
assert(!repair_plan["host_root_modified"], "D-Bus client must parse repair host-root false booleans")
assert(!repair_plan["backend_details_exposed"], "D-Bus client must parse repair backend detail false booleans")

test_plan = client.test_plan("org.xnix.sample.notepad")
assert(test_plan["plan_type"] == "compatibility-test", "D-Bus client must parse test plans")
assert(test_plan["runtime_method"] == "GetTestPlan", "D-Bus client must parse test plan Runtime methods")
assert(test_plan["runtime_owned"], "D-Bus client must parse test plan booleans")
assert(test_plan["c_runtime_backed"], "D-Bus client must parse test plan C backing")
assert(!test_plan["kde_policy_owner"], "D-Bus client must parse test plan KDE false booleans")
assert(!test_plan["execution_request_created"], "D-Bus client must parse test plan execution request false booleans")
assert(!test_plan["test_executed"], "D-Bus client must parse test plan execution false booleans")
assert(!test_plan["host_root_modified"], "D-Bus client must parse test plan host-root false booleans")
assert(!test_plan["backend_details_exposed"], "D-Bus client must parse test plan false booleans")

test_result = client.test_result("org.xnix.sample.notepad")
assert(test_result["result_type"] == "compatibility-test-result", "D-Bus client must parse test results")
assert(test_result["runtime_method"] == "GetTestResult", "D-Bus client must parse test result Runtime methods")
assert(test_result["runtime_owned"], "D-Bus client must parse test result Runtime ownership")
assert(test_result["c_runtime_backed"], "D-Bus client must parse test result C backing")
assert(!test_result["kde_policy_owner"], "D-Bus client must parse test result KDE false booleans")
assert(test_result["result_source"] == "runtime-model", "D-Bus client must parse test result sources")
assert(test_result["execution_state"] == "waiting-for-runtime", "D-Bus client must parse test result execution state")
assert(test_result["safe_for_ai_diagnostics"], "D-Bus client must parse test result booleans")
assert(!test_result["test_executed"], "D-Bus client must parse test result execution false booleans")
assert(!test_result["host_root_modified"], "D-Bus client must parse test result host-root false booleans")
assert(!test_result["backend_details_exposed"], "D-Bus client must parse test result false booleans")

execution_readiness = client.execution_readiness("org.xnix.sample.notepad")
assert(execution_readiness["readiness_type"] == "compatibility-execution-readiness", "D-Bus client must parse execution readiness")
assert(execution_readiness["runtime_method"] == "GetExecutionReadiness", "D-Bus client must parse execution readiness Runtime methods")
assert(execution_readiness["runtime_owned"], "D-Bus client must parse execution readiness Runtime ownership")
assert(execution_readiness["c_runtime_backed"], "D-Bus client must parse execution readiness C backing")
assert(!execution_readiness["kde_policy_owner"], "D-Bus client must parse execution readiness KDE false booleans")
assert(execution_readiness["execution_state"] == "blocked", "D-Bus client must parse execution readiness state")
assert(!execution_readiness["launch_allowed"], "D-Bus client must parse execution readiness launch allowance")
assert(!execution_readiness["launch_enabled"], "D-Bus client must parse execution readiness launch enablement")
assert(!execution_readiness["execution_request_created"], "D-Bus client must parse execution readiness request creation")
assert(!execution_readiness["backend_binding_ready"], "D-Bus client must parse execution readiness backend binding state")
assert(execution_readiness["desktop_entry_launch_visible"], "D-Bus client must parse execution readiness desktop visibility")
assert(execution_readiness["safe_for_ai_diagnostics"], "D-Bus client must parse execution readiness AI safety")
assert(!execution_readiness["host_root_modified"], "D-Bus client must parse execution readiness host-root false booleans")
assert(!execution_readiness["backend_details_exposed"], "D-Bus client must parse execution readiness false booleans")

launch_intent = client.launch_intent("org.xnix.sample.notepad")
assert(launch_intent["intent_type"] == "runtime-launch-intent", "D-Bus client must parse launch intent")
assert(launch_intent["source"] == "desktop-launcher", "D-Bus client must parse launch intent source")
assert(launch_intent["runtime_method"] == "Launch", "D-Bus client must parse launch intent Runtime method")
assert(launch_intent["read_method"] == "GetLaunchIntent", "D-Bus client must parse launch intent read method")
assert(launch_intent["runtime_owned"], "D-Bus client must parse launch intent Runtime ownership")
assert(launch_intent["c_runtime_backed"], "D-Bus client must parse launch intent C backing")
assert(!launch_intent["kde_policy_owner"], "D-Bus client must parse launch intent KDE false booleans")
assert(launch_intent["standard_desktop_entry"], "D-Bus client must parse launch intent desktop-entry status")
assert(launch_intent["launch_uses_runtime"], "D-Bus client must parse launch intent Runtime launcher usage")
assert(launch_intent["desktop_entry_launch_visible"], "D-Bus client must parse launch intent desktop visibility")
assert(!launch_intent["launch_allowed"], "D-Bus client must parse launch intent allowance")
assert(!launch_intent["launch_enabled"], "D-Bus client must parse launch intent enablement")
assert(!launch_intent["execution_request_created"], "D-Bus client must parse launch intent request creation")
assert(!launch_intent["execution_started"], "D-Bus client must parse launch intent execution state")
assert(launch_intent["write_gate_decision"] == "blocked-until-production-backend", "D-Bus client must parse launch intent write gate")
assert(!launch_intent["host_root_modified"], "D-Bus client must parse launch intent host-root false booleans")
assert(!launch_intent["network_required"], "D-Bus client must parse launch intent network status")
assert(!launch_intent["backend_details_exposed"], "D-Bus client must parse launch intent backend detail status")

ai_input = client.ai_diagnostic_input("org.xnix.sample.notepad")
assert(ai_input["input_type"] == "ai-diagnostic-input", "D-Bus client must parse AI diagnostic inputs")
assert(ai_input["runtime_method"] == "GetAIDiagnosticInput", "D-Bus client must parse AI diagnostic Runtime method")
assert(ai_input["c_runtime_backed"], "D-Bus client must parse AI diagnostic C ownership")
assert(ai_input["diagnostic_signal_count"] == 3, "D-Bus client must parse AI diagnostic signal count")
assert(ai_input["safe_for_ai_diagnostics"], "D-Bus client must parse AI diagnostic input booleans")
assert(!ai_input["ai_provider_called"], "D-Bus client must parse AI provider false booleans")
assert(!ai_input["network_required"], "D-Bus client must parse network false booleans")

ai_recommendation = client.ai_diagnostic_recommendation("org.xnix.sample.notepad")
assert(ai_recommendation["recommendation_type"] == "ai-diagnostic-recommendation", "D-Bus client must parse AI diagnostic recommendations")
assert(ai_recommendation["runtime_method"] == "GetAIDiagnosticRecommendation", "D-Bus client must parse AI recommendation Runtime method")
assert(ai_recommendation["c_runtime_backed"], "D-Bus client must parse AI recommendation C ownership")
assert(ai_recommendation["recommendation_count"] == 3, "D-Bus client must parse AI recommendation count")
assert(ai_recommendation["approval_required_count"] == 1, "D-Bus client must parse AI recommendation approval count")
assert(ai_recommendation["safe_for_ai_diagnostics"], "D-Bus client must parse AI diagnostic recommendation booleans")
assert(!ai_recommendation["ai_provider_called"], "D-Bus client must parse recommendation AI provider false booleans")
assert(!ai_recommendation["network_required"], "D-Bus client must parse recommendation network false booleans")
assert(!ai_recommendation["auto_execution_allowed"], "D-Bus client must parse recommendation execution gates")

ai_gate = client.ai_repair_approval_gate("org.xnix.sample.notepad")
assert(ai_gate["gate_type"] == "ai-repair-approval-gate", "D-Bus client must parse AI repair approval gates")
assert(ai_gate["runtime_method"] == "GetAIRepairApprovalGate", "D-Bus client must parse AI repair approval Runtime methods")
assert(ai_gate["c_runtime_backed"], "D-Bus client must parse AI repair approval C backing")
assert(ai_gate["gate_decision"] == "blocked-until-approval", "D-Bus client must parse AI repair approval gate decisions")
assert(ai_gate["required_gate_count"] == 3, "D-Bus client must parse AI repair required gate count")
assert(ai_gate["approval_required_count"] == 1, "D-Bus client must parse AI repair approval action count")
assert(ai_gate["safe_for_ai_diagnostics"], "D-Bus client must parse AI repair diagnostic safety")
assert(!ai_gate["ai_provider_called"], "D-Bus client must parse AI repair provider false booleans")
assert(!ai_gate["network_required"], "D-Bus client must parse AI repair network false booleans")
assert(!ai_gate["auto_execution_allowed"], "D-Bus client must parse AI repair gate execution booleans")
assert(!ai_gate["repair_executed"], "D-Bus client must parse AI repair execution false booleans")

snapshot_plan = client.snapshot_plan("org.xnix.sample.notepad", "before-repair")
assert(snapshot_plan["plan_type"] == "compatibility-snapshot", "D-Bus client must parse snapshot plans")
assert(snapshot_plan["runtime_method"] == "GetSnapshotPlan", "D-Bus client must parse snapshot plan Runtime methods")
assert(snapshot_plan["runtime_owned"], "D-Bus client must parse snapshot plan Runtime ownership")
assert(snapshot_plan["c_runtime_backed"], "D-Bus client must parse snapshot plan C backing")
assert(!snapshot_plan["kde_policy_owner"], "D-Bus client must parse snapshot plan KDE false booleans")
assert(snapshot_plan["enabled_by_default"], "D-Bus client must parse snapshot plan booleans")
assert(!snapshot_plan["snapshot_request_created"], "D-Bus client must parse snapshot request false booleans")
assert(!snapshot_plan["snapshot_created"], "D-Bus client must parse snapshot created false booleans")
assert(!snapshot_plan["restore_requested"], "D-Bus client must parse restore request false booleans")
assert(!snapshot_plan["restore_executed"], "D-Bus client must parse restore execution false booleans")
assert(!snapshot_plan["user_documents_included"], "D-Bus client must parse user document exclusion")
assert(!snapshot_plan["host_system_included"], "D-Bus client must parse host system exclusion")
assert(!snapshot_plan["host_root_modified"], "D-Bus client must parse snapshot host-root false booleans")
assert(!snapshot_plan["backend_details_exposed"], "D-Bus client must parse snapshot backend detail false booleans")

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
assert(method_parity["method_count"] == 48, "D-Bus client must parse Runtime method counts")
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

compatibility_center_summary = client.compatibility_center_summary("org.xnix.sample.notepad")
assert(compatibility_center_summary["summary_type"] == "compatibility-center-summary", "D-Bus client must parse Compatibility Center summaries")
assert(compatibility_center_summary["compatibility_state"] == "review-required", "D-Bus client must parse Compatibility Center state")
assert(compatibility_center_summary["known_issue_count"] == 1, "D-Bus client must parse Compatibility Center issue counts")
assert(compatibility_center_summary["repair_record_state"] == "pending-review", "D-Bus client must parse Compatibility Center repair state")
assert(compatibility_center_summary["action_count"] == 4, "D-Bus client must parse Compatibility Center action counts")
assert(compatibility_center_summary["runtime_owned"], "D-Bus client must parse Compatibility Center Runtime ownership")
assert(!compatibility_center_summary["kde_policy_owner"], "D-Bus client must parse Compatibility Center KDE policy ownership")
assert(compatibility_center_summary["user_visible"], "D-Bus client must parse Compatibility Center visibility")
assert(!compatibility_center_summary["action_execution_enabled"], "D-Bus client must parse Compatibility Center action gates")
assert(!compatibility_center_summary["repair_execution_enabled"], "D-Bus client must parse Compatibility Center repair gates")
assert(!compatibility_center_summary["backend_launch_enabled"], "D-Bus client must parse Compatibility Center launch gates")
assert(!compatibility_center_summary["backend_details_exposed"], "D-Bus client must parse Compatibility Center backend detail status")

assert(
  capture.commands.all? { |command| command.include?("--session") },
  "D-Bus client must use the session bus for KDE-facing reads"
)

puts "PASS: compatibility runtime D-Bus client unit tests"
