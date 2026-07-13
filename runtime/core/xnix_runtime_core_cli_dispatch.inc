static int
usage(void)
{
  fputs("Usage: xnix-runtime-core {probe|list-applications|get-application APP_ID|list-engines|select-engine MODE|list-portal-policies|portal-policy APP_ID OPERATION|portal-request-plan APP_ID OPERATION|list-snapshot-policies|snapshot-policy APP_ID REASON|list-state-root-policies|state-root-policy APP_ID|list-install-readiness-policies|install-readiness-policy APP_ID ENVIRONMENT|compatibility-install-plan APP_ID ENVIRONMENT|compatibility-action-queue APP_ID|compatibility-action-review APP_ID ACTION_ID DECISION|compatibility-run-plan APP_ID|desktop-activation-manifest APP_ID|list-artifact-manifest-policies|artifact-manifest-policy APP_ID|list-acquisition-preflight-policies|acquisition-preflight-policy APP_ID|list-package-source-policies|package-source-policy APP_ID|list-backend-binding-policies|backend-binding-policy APP_ID|list-settings-policies|settings-policy APP_ID|list-settings-change-policies|settings-change-policy APP_ID|runtime-service-binding|runtime-live-owner-gate|runtime-owner-smoke-plan|runtime-method-parity-manifest|recipe-trust-policy|recipe-install-gate APP_ID MODE|write-gate METHOD}\n", stderr);
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

  if (strcmp(argv[1], "list-snapshot-policies") == 0 && argc == 2) {
    return print_snapshot_policy_catalog();
  }

  if (strcmp(argv[1], "snapshot-policy") == 0 && argc == 4) {
    return print_snapshot_policy(argv[2], argv[3]);
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

  if (strcmp(argv[1], "compatibility-run-plan") == 0 && argc == 3) {
    return print_compatibility_run_plan_for_application(argv[2]);
  }

  if (strcmp(argv[1], "desktop-activation-manifest") == 0 && argc == 3) {
    return print_desktop_activation_manifest_for_application(argv[2]);
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
