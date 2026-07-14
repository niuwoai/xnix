static int
usage(void)
{
  fputs("Usage: xnix-runtime-core {probe|list-applications|get-application APP_ID|list-engines|select-engine MODE|list-portal-policies|portal-policy APP_ID OPERATION|portal-request-plan APP_ID OPERATION|compatibility-review-flow-plan APP_ID SECTION FIELD VALUE OPERATION|list-snapshot-policies|snapshot-policy APP_ID REASON|compatibility-snapshot-plan APP_ID REASON|list-state-root-policies|state-root-policy APP_ID|list-install-readiness-policies|install-readiness-policy APP_ID ENVIRONMENT|compatibility-install-plan APP_ID ENVIRONMENT|compatibility-action-queue APP_ID|compatibility-action-review APP_ID ACTION_ID DECISION|compatibility-center-summary APP_ID|compatibility-run-plan APP_ID|compatibility-repair-plan APP_ID ISSUE|compatibility-test-plan APP_ID TEST_TYPE|compatibility-test-result APP_ID TEST_TYPE|execution-readiness APP_ID|launch-intent APP_ID|backend-capability-matrix|backend-selection-plan APP_ID|desktop-activation-manifest APP_ID|kde-integration-status|kde-application-surface-plan APP_ID|desktop-resource-bridge-plan APP_ID|compatibility-mode-switch-plan APP_ID MODE|compatibility-permission-review-plan APP_ID|desktop-entry-plan APP_ID|task-manager-identity-plan APP_ID|kwin-window-rule-plan APP_ID|file-association-plan APP_ID|notification-plan APP_ID EVENT_TYPE|tray-status-plan|krunner-query-plan QUERY|ai-diagnostic-input APP_ID ISSUE TEST_TYPE|ai-diagnostic-recommendation APP_ID ISSUE TEST_TYPE|ai-repair-approval-gate APP_ID ISSUE TEST_TYPE|list-artifact-manifest-policies|artifact-manifest-policy APP_ID|list-acquisition-preflight-policies|acquisition-preflight-policy APP_ID|list-package-source-policies|package-source-policy APP_ID|list-backend-binding-policies|backend-binding-policy APP_ID|list-settings-policies|settings-policy APP_ID|list-settings-change-policies|settings-change-policy APP_ID|runtime-service-binding|runtime-live-owner-gate|runtime-owner-smoke-plan|runtime-method-parity-manifest|recipe-trust-policy|recipe-install-gate APP_ID MODE|write-gate METHOD}\n", stderr);
  return 64;
}

int
main(int argc, char **argv)
{
  if (argc < 2) {
    return usage();
  }

  if (strcmp(argv[1], "probe") == 0 && argc == 2) {
    return print_probe();
  }

  if (strcmp(argv[1], "list-applications") == 0 && argc == 2) {
    return print_applications();
  }

  if (strcmp(argv[1], "get-application") == 0 && argc == 3) {
    return print_application_by_id(argv[2]);
  }

  if (strcmp(argv[1], "list-engines") == 0 && argc == 2) {
    return print_engine_catalog();
  }

  if (strcmp(argv[1], "select-engine") == 0 && argc == 3) {
    return print_engine_for_mode(argv[2]);
  }

  if (strcmp(argv[1], "list-portal-policies") == 0 && argc == 2) {
    return print_portal_policy_catalog();
  }

  if (strcmp(argv[1], "portal-policy") == 0 && argc == 4) {
    return print_portal_policy(argv[2], argv[3]);
  }

  if (strcmp(argv[1], "portal-request-plan") == 0 && argc == 4) {
    return print_portal_request_plan_for_application(argv[2], argv[3]);
  }

  if (strcmp(argv[1], "compatibility-review-flow-plan") == 0 && argc == 7) {
    return print_compatibility_review_flow_plan_for_application(
      argv[2],
      argv[3],
      argv[4],
      argv[5],
      argv[6]
    );
  }

  if (strcmp(argv[1], "list-snapshot-policies") == 0 && argc == 2) {
    return print_snapshot_policy_catalog();
  }

  if (strcmp(argv[1], "snapshot-policy") == 0 && argc == 4) {
    return print_snapshot_policy(argv[2], argv[3]);
  }

  if (strcmp(argv[1], "compatibility-snapshot-plan") == 0 && argc == 4) {
    return print_compatibility_snapshot_plan_for_application(argv[2], argv[3]);
  }

  if (strcmp(argv[1], "list-state-root-policies") == 0 && argc == 2) {
    return print_state_root_policy_catalog();
  }

  if (strcmp(argv[1], "state-root-policy") == 0 && argc == 3) {
    return print_state_root_policy(argv[2]);
  }

  if (strcmp(argv[1], "list-install-readiness-policies") == 0 && argc == 2) {
    return print_install_readiness_policy_catalog();
  }

  if (strcmp(argv[1], "install-readiness-policy") == 0 && argc == 4) {
    return print_install_readiness_policy(argv[2], argv[3]);
  }

  if (strcmp(argv[1], "compatibility-install-plan") == 0 && argc == 4) {
    return print_compatibility_install_plan_for_application(argv[2], argv[3]);
  }

  if (strcmp(argv[1], "compatibility-action-queue") == 0 && argc == 3) {
    return print_compatibility_action_queue_for_application(argv[2]);
  }

  if (strcmp(argv[1], "compatibility-action-review") == 0 && argc == 5) {
    return print_compatibility_action_review_receipt_for_application(
      argv[2],
      argv[3],
      argv[4]
    );
  }

  if (strcmp(argv[1], "compatibility-center-summary") == 0 && argc == 3) {
    return print_compatibility_center_summary_for_application(argv[2]);
  }

  if (strcmp(argv[1], "compatibility-run-plan") == 0 && argc == 3) {
    return print_compatibility_run_plan_for_application(argv[2]);
  }

  if (strcmp(argv[1], "compatibility-repair-plan") == 0 && argc == 4) {
    return print_compatibility_repair_plan_for_application(argv[2], argv[3]);
  }

  if (strcmp(argv[1], "compatibility-test-plan") == 0 && argc == 4) {
    return print_compatibility_test_plan_for_application(argv[2], argv[3]);
  }

  if (strcmp(argv[1], "compatibility-test-result") == 0 && argc == 4) {
    return print_compatibility_test_result_for_application(argv[2], argv[3]);
  }

  if (strcmp(argv[1], "execution-readiness") == 0 && argc == 3) {
    return print_execution_readiness_for_application(argv[2]);
  }

  if (strcmp(argv[1], "launch-intent") == 0 && argc == 3) {
    return print_launch_intent_for_application(argv[2]);
  }

  if (strcmp(argv[1], "backend-capability-matrix") == 0 && argc == 2) {
    return print_backend_capability_matrix();
  }

  if (strcmp(argv[1], "backend-selection-plan") == 0 && argc == 3) {
    return print_backend_selection_plan_for_application(argv[2]);
  }

  if (strcmp(argv[1], "desktop-activation-manifest") == 0 && argc == 3) {
    return print_desktop_activation_manifest_for_application(argv[2]);
  }

  if (strcmp(argv[1], "kde-integration-status") == 0 && argc == 2) {
    return print_kde_integration_status();
  }

  if (strcmp(argv[1], "kde-application-surface-plan") == 0 && argc == 3) {
    return print_kde_application_surface_plan_for_application(argv[2]);
  }

  if (strcmp(argv[1], "desktop-resource-bridge-plan") == 0 && argc == 3) {
    return print_desktop_resource_bridge_plan_for_application(argv[2]);
  }

  if (strcmp(argv[1], "compatibility-mode-switch-plan") == 0 && argc == 4) {
    return print_compatibility_mode_switch_plan_for_application(argv[2], argv[3]);
  }

  if (strcmp(argv[1], "compatibility-permission-review-plan") == 0 && argc == 3) {
    return print_compatibility_permission_review_plan_for_application(argv[2]);
  }

  if (strcmp(argv[1], "desktop-entry-plan") == 0 && argc == 3) {
    return print_desktop_entry_plan_for_application(argv[2]);
  }

  if (strcmp(argv[1], "task-manager-identity-plan") == 0 && argc == 3) {
    return print_task_manager_identity_plan_for_application(argv[2]);
  }

  if (strcmp(argv[1], "kwin-window-rule-plan") == 0 && argc == 3) {
    return print_kwin_window_rule_plan_for_application(argv[2]);
  }

  if (strcmp(argv[1], "file-association-plan") == 0 && argc == 3) {
    return print_file_association_plan_for_application(argv[2]);
  }

  if (strcmp(argv[1], "notification-plan") == 0 && argc == 4) {
    return print_notification_plan_for_application(argv[2], argv[3]);
  }

  if (strcmp(argv[1], "tray-status-plan") == 0 && argc == 2) {
    return print_tray_status_plan();
  }

  if (strcmp(argv[1], "krunner-query-plan") == 0 && argc == 3) {
    return print_krunner_query_plan_for_query(argv[2]);
  }

  if (strcmp(argv[1], "ai-diagnostic-input") == 0 && argc == 5) {
    return print_ai_diagnostic_input_for_application(argv[2], argv[3], argv[4]);
  }

  if (strcmp(argv[1], "ai-diagnostic-recommendation") == 0 && argc == 5) {
    return print_ai_diagnostic_recommendation_for_application(argv[2], argv[3], argv[4]);
  }

  if (strcmp(argv[1], "ai-repair-approval-gate") == 0 && argc == 5) {
    return print_ai_repair_approval_gate_for_application(argv[2], argv[3], argv[4]);
  }

  if (strcmp(argv[1], "list-artifact-manifest-policies") == 0 && argc == 2) {
    return print_artifact_manifest_policy_catalog();
  }

  if (strcmp(argv[1], "artifact-manifest-policy") == 0 && argc == 3) {
    return print_artifact_manifest_policy(argv[2]);
  }

  if (strcmp(argv[1], "list-acquisition-preflight-policies") == 0 && argc == 2) {
    return print_acquisition_preflight_policy_catalog();
  }

  if (strcmp(argv[1], "acquisition-preflight-policy") == 0 && argc == 3) {
    return print_acquisition_preflight_policy(argv[2]);
  }

  if (strcmp(argv[1], "list-package-source-policies") == 0 && argc == 2) {
    return print_package_source_policy_catalog();
  }

  if (strcmp(argv[1], "package-source-policy") == 0 && argc == 3) {
    return print_package_source_policy(argv[2]);
  }

  if (strcmp(argv[1], "list-backend-binding-policies") == 0 && argc == 2) {
    return print_backend_binding_policy_catalog();
  }

  if (strcmp(argv[1], "backend-binding-policy") == 0 && argc == 3) {
    return print_backend_binding_policy(argv[2]);
  }

  if (strcmp(argv[1], "backend-lifecycle") == 0 && argc == 3) {
    return print_backend_lifecycle_for_application(argv[2]);
  }

  if (strcmp(argv[1], "backend-environment-plan") == 0 && argc == 3) {
    return print_backend_environment_plan_for_application(argv[2]);
  }

  if (strcmp(argv[1], "list-settings-policies") == 0 && argc == 2) {
    return print_settings_policy_catalog();
  }

  if (strcmp(argv[1], "settings-policy") == 0 && argc == 3) {
    return print_settings_policy(argv[2]);
  }

  if (strcmp(argv[1], "list-settings-change-policies") == 0 && argc == 2) {
    return print_settings_change_policy_catalog();
  }

  if (strcmp(argv[1], "settings-change-policy") == 0 && argc == 3) {
    return print_settings_change_policy(argv[2]);
  }

  if (strcmp(argv[1], "runtime-service-binding") == 0 && argc == 2) {
    return print_runtime_service_binding();
  }

  if (strcmp(argv[1], "runtime-live-owner-gate") == 0 && argc == 2) {
    return print_runtime_live_owner_gate();
  }

  if (strcmp(argv[1], "runtime-owner-smoke-plan") == 0 && argc == 2) {
    return print_runtime_owner_smoke_plan();
  }

  if (strcmp(argv[1], "runtime-method-parity-manifest") == 0 && argc == 2) {
    return print_runtime_method_parity_manifest();
  }

  if (strcmp(argv[1], "recipe-trust-policy") == 0 && argc == 2) {
    return print_recipe_trust();
  }

  if (strcmp(argv[1], "recipe-install-gate") == 0 && argc == 4) {
    return print_recipe_install_gate_for_application(argv[2], argv[3]);
  }

  if (strcmp(argv[1], "write-gate") == 0 && argc == 3) {
    return print_write_gate(argv[2]);
  }

  return usage();
}
