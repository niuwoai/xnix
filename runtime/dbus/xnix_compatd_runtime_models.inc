static GVariant *
build_portal_request_plan(const gchar *application_id, const gchar *operation)
{
  GVariantBuilder plan;
  gboolean denied = g_strcmp0(operation, "camera") == 0 ||
                    g_strcmp0(operation, "remote-desktop") == 0;

  g_variant_builder_init(&plan, G_VARIANT_TYPE("a{sv}"));
  g_variant_builder_add(&plan, "{sv}", "request_type", g_variant_new_string("portal-request-plan"));
  add_go_owner_dispatch_bridge_fields2(&plan, "GetPortalRequestPlan", application_id, operation);
  g_variant_builder_add(&plan, "{sv}", "application_id", g_variant_new_string(application_id));
  g_variant_builder_add(&plan, "{sv}", "operation", g_variant_new_string(operation));
  g_variant_builder_add(&plan, "{sv}", "decision", g_variant_new_string(denied ? "deny" : "ask"));
  g_variant_builder_add(&plan, "{sv}", "request_allowed", g_variant_new_boolean(!denied));
  g_variant_builder_add(&plan, "{sv}", "portal_required", g_variant_new_boolean(TRUE));
  g_variant_builder_add(&plan, "{sv}", "request_object_created", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&plan, "{sv}", "permission_granted", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&plan, "{sv}", "host_permission_changed", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&plan, "{sv}", "host_root_modified", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&plan, "{sv}", "backend_details_exposed", g_variant_new_boolean(FALSE));

  return g_variant_builder_end(&plan);
}

static GVariant *
build_application_state_root(const gchar *application_id)
{
  GVariantBuilder state_root;

  g_variant_builder_init(&state_root, G_VARIANT_TYPE("a{sv}"));
  g_variant_builder_add(&state_root, "{sv}", "root_type", g_variant_new_string("compatibility-application-state-root"));
  g_variant_builder_add(&state_root, "{sv}", "application_id", g_variant_new_string(application_id));
  g_variant_builder_add(&state_root, "{sv}", "state_namespace", g_variant_new_string(application_id));
  g_variant_builder_add(&state_root, "{sv}", "allocation_state", g_variant_new_string("planned"));
  g_variant_builder_add(&state_root, "{sv}", "snapshot_eligible", g_variant_new_boolean(TRUE));
  g_variant_builder_add(&state_root, "{sv}", "user_documents_included", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&state_root, "{sv}", "backend_details_exposed", g_variant_new_boolean(FALSE));

  return g_variant_builder_end(&state_root);
}

static GVariant *
build_compatibility_package_source(const gchar *application_id)
{
  GVariantBuilder source;

  g_variant_builder_init(&source, G_VARIANT_TYPE("a{sv}"));
  g_variant_builder_add(&source, "{sv}", "source_type", g_variant_new_string("compatibility-package-source"));
  add_go_owner_dispatch_bridge_fields(&source, "GetCompatibilityPackageSource", application_id);
  g_variant_builder_add(&source, "{sv}", "application_id", g_variant_new_string(application_id));
  g_variant_builder_add(&source, "{sv}", "selected_strategy", g_variant_new_string("automatic-managed"));
  g_variant_builder_add(&source, "{sv}", "source_selection_state", g_variant_new_string("planned"));
  g_variant_builder_add(&source, "{sv}", "package_source_ready", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&source, "{sv}", "install_enabled", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&source, "{sv}", "backend_details_exposed", g_variant_new_boolean(FALSE));

  return g_variant_builder_end(&source);
}

static GVariant *
build_compatibility_acquisition_preflight(const gchar *application_id)
{
  GVariantBuilder preflight;

  g_variant_builder_init(&preflight, G_VARIANT_TYPE("a{sv}"));
  g_variant_builder_add(&preflight, "{sv}", "preflight_type", g_variant_new_string("compatibility-acquisition-preflight"));
  add_go_owner_dispatch_bridge_fields(&preflight, "GetCompatibilityAcquisitionPreflight", application_id);
  g_variant_builder_add(&preflight, "{sv}", "application_id", g_variant_new_string(application_id));
  g_variant_builder_add(&preflight, "{sv}", "selected_strategy", g_variant_new_string("automatic-managed"));
  g_variant_builder_add(&preflight, "{sv}", "preflight_state", g_variant_new_string("planned"));
  g_variant_builder_add(&preflight, "{sv}", "acquisition_ready", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&preflight, "{sv}", "download_enabled", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&preflight, "{sv}", "install_enabled", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&preflight, "{sv}", "backend_details_exposed", g_variant_new_boolean(FALSE));

  return g_variant_builder_end(&preflight);
}

static GVariant *
build_compatibility_artifact_manifest(const gchar *application_id)
{
  GVariantBuilder manifest;

  g_variant_builder_init(&manifest, G_VARIANT_TYPE("a{sv}"));
  g_variant_builder_add(&manifest, "{sv}", "manifest_type", g_variant_new_string("compatibility-artifact-manifest"));
  add_go_owner_dispatch_bridge_fields(&manifest, "GetCompatibilityArtifactManifest", application_id);
  g_variant_builder_add(&manifest, "{sv}", "application_id", g_variant_new_string(application_id));
  g_variant_builder_add(&manifest, "{sv}", "selected_strategy", g_variant_new_string("automatic-managed"));
  g_variant_builder_add(&manifest, "{sv}", "manifest_state", g_variant_new_string("planned"));
  g_variant_builder_add(&manifest, "{sv}", "manifest_ready", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&manifest, "{sv}", "signature_verified", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&manifest, "{sv}", "download_enabled", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&manifest, "{sv}", "backend_details_exposed", g_variant_new_boolean(FALSE));

  return g_variant_builder_end(&manifest);
}

static GVariant *
build_compatibility_install_plan(const gchar *application_id, const gchar *environment)
{
  GVariantBuilder plan;

  g_variant_builder_init(&plan, G_VARIANT_TYPE("a{sv}"));
  g_variant_builder_add(&plan, "{sv}", "plan_type", g_variant_new_string("compatibility-install-plan"));
  add_go_owner_dispatch_bridge_fields2(&plan, "GetCompatibilityInstallPlan", application_id, environment);
  g_variant_builder_add(&plan, "{sv}", "application_id", g_variant_new_string(application_id));
  g_variant_builder_add(&plan, "{sv}", "environment", g_variant_new_string(environment));
  g_variant_builder_add(&plan, "{sv}", "selected_strategy", g_variant_new_string("automatic-managed"));
  g_variant_builder_add(&plan, "{sv}", "install_state", g_variant_new_string("planned"));
  g_variant_builder_add(&plan, "{sv}", "install_ready", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&plan, "{sv}", "desktop_activation_ready", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&plan, "{sv}", "download_enabled", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&plan, "{sv}", "install_enabled", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&plan, "{sv}", "backend_details_exposed", g_variant_new_boolean(FALSE));

  return g_variant_builder_end(&plan);
}

static GVariant *
build_backend_binding(const gchar *application_id)
{
  GVariantBuilder binding;

  g_variant_builder_init(&binding, G_VARIANT_TYPE("a{sv}"));
  g_variant_builder_add(&binding, "{sv}", "binding_type", g_variant_new_string("compatibility-backend-binding"));
  add_go_owner_dispatch_bridge_fields(&binding, "GetBackendBinding", application_id);
  g_variant_builder_add(&binding, "{sv}", "application_id", g_variant_new_string(application_id));
  g_variant_builder_add(&binding, "{sv}", "selected_strategy", g_variant_new_string("automatic-managed"));
  g_variant_builder_add(&binding, "{sv}", "managed_binding_ready", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&binding, "{sv}", "launch_enabled", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&binding, "{sv}", "backend_details_exposed", g_variant_new_boolean(FALSE));

  return g_variant_builder_end(&binding);
}

static GVariant *
build_backend_capability_matrix(void)
{
  const gchar *profile_ids[] = {
    "local-compatibility",
    "isolated-compatibility",
  };
  const gchar *capability_ids[] = {
    "application-launch",
    "package-management",
    "file-bridge",
    "clipboard-bridge",
    "print-bridge",
    "snapshot-restore",
    "diagnostics",
  };
  GVariantBuilder matrix;

  g_variant_builder_init(&matrix, G_VARIANT_TYPE("a{sv}"));
  g_variant_builder_add(&matrix, "{sv}", "matrix_type", g_variant_new_string("compatibility-backend-capability-matrix"));
  add_go_owner_dispatch_bridge_fields(&matrix, "GetBackendCapabilityMatrix", NULL);
  g_variant_builder_add(&matrix, "{sv}", "runtime_method", g_variant_new_string("GetBackendCapabilityMatrix"));
  g_variant_builder_add(&matrix, "{sv}", "profile_ids", g_variant_new_strv(profile_ids, 2));
  g_variant_builder_add(&matrix, "{sv}", "capability_ids", g_variant_new_strv(capability_ids, 7));
  g_variant_builder_add(&matrix, "{sv}", "profile_count", g_variant_new_int32(2));
  g_variant_builder_add(&matrix, "{sv}", "capability_count", g_variant_new_int32(7));
  g_variant_builder_add(&matrix, "{sv}", "ready_capability_count", g_variant_new_int32(2));
  g_variant_builder_add(&matrix, "{sv}", "pending_capability_count", g_variant_new_int32(12));
  g_variant_builder_add(&matrix, "{sv}", "blocked_capability_count", g_variant_new_int32(0));
  g_variant_builder_add(&matrix, "{sv}", "runtime_owned", g_variant_new_boolean(TRUE));
  g_variant_builder_add(&matrix, "{sv}", "c_runtime_backed", g_variant_new_boolean(TRUE));
  g_variant_builder_add(&matrix, "{sv}", "kde_policy_owner", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&matrix, "{sv}", "selection_enabled", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&matrix, "{sv}", "backend_launch_enabled", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&matrix, "{sv}", "capability_activation_enabled", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&matrix, "{sv}", "request_objects_created", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&matrix, "{sv}", "state_root_created", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&matrix, "{sv}", "snapshots_created", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&matrix, "{sv}", "host_root_modified", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&matrix, "{sv}", "privileged_container_required", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&matrix, "{sv}", "backend_details_exposed", g_variant_new_boolean(FALSE));

  return g_variant_builder_end(&matrix);
}

static GVariant *
build_backend_selection_plan(const gchar *application_id)
{
  const gchar *candidate_ids[] = {
    "local-compatibility",
    "isolated-compatibility",
  };
  const gchar *required_reviews[] = {
    "backend-capability-review",
    "backend-binding-review",
    "application-state-root-review",
    "portal-policy-review",
    "snapshot-baseline-review",
  };
  GVariantBuilder plan;

  g_variant_builder_init(&plan, G_VARIANT_TYPE("a{sv}"));
  g_variant_builder_add(&plan, "{sv}", "plan_type", g_variant_new_string("compatibility-backend-selection-plan"));
  add_go_owner_dispatch_bridge_fields(&plan, "GetBackendSelectionPlan", application_id);
  g_variant_builder_add(&plan, "{sv}", "runtime_method", g_variant_new_string("GetBackendSelectionPlan"));
  g_variant_builder_add(&plan, "{sv}", "application_id", g_variant_new_string(application_id));
  g_variant_builder_add(&plan, "{sv}", "selected_strategy", g_variant_new_string("automatic-managed"));
  g_variant_builder_add(&plan, "{sv}", "recommended_profile_id", g_variant_new_string("local-compatibility"));
  g_variant_builder_add(&plan, "{sv}", "candidate_profile_ids", g_variant_new_strv(candidate_ids, 2));
  g_variant_builder_add(&plan, "{sv}", "required_reviews", g_variant_new_strv(required_reviews, 5));
  g_variant_builder_add(&plan, "{sv}", "candidate_count", g_variant_new_int32(2));
  g_variant_builder_add(&plan, "{sv}", "ready_candidate_count", g_variant_new_int32(0));
  g_variant_builder_add(&plan, "{sv}", "blocked_candidate_count", g_variant_new_int32(2));
  g_variant_builder_add(&plan, "{sv}", "runtime_owned", g_variant_new_boolean(TRUE));
  g_variant_builder_add(&plan, "{sv}", "c_runtime_backed", g_variant_new_boolean(TRUE));
  g_variant_builder_add(&plan, "{sv}", "kde_policy_owner", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&plan, "{sv}", "selection_committed", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&plan, "{sv}", "selection_change_enabled", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&plan, "{sv}", "backend_launch_enabled", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&plan, "{sv}", "capability_activation_enabled", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&plan, "{sv}", "environment_created", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&plan, "{sv}", "request_object_created", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&plan, "{sv}", "state_root_created", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&plan, "{sv}", "snapshot_created", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&plan, "{sv}", "host_root_modified", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&plan, "{sv}", "privileged_container_required", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&plan, "{sv}", "backend_details_exposed", g_variant_new_boolean(FALSE));

  return g_variant_builder_end(&plan);
}

static GVariant *
build_backend_lifecycle(const gchar *application_id)
{
  GVariantBuilder lifecycle;

  g_variant_builder_init(&lifecycle, G_VARIANT_TYPE("a{sv}"));
  g_variant_builder_add(&lifecycle, "{sv}", "lifecycle_type", g_variant_new_string("compatibility-backend-lifecycle"));
  add_go_owner_dispatch_bridge_fields(&lifecycle, "GetBackendLifecycle", application_id);
  g_variant_builder_add(&lifecycle, "{sv}", "runtime_method", g_variant_new_string("GetBackendLifecycle"));
  g_variant_builder_add(&lifecycle, "{sv}", "application_id", g_variant_new_string(application_id));
  g_variant_builder_add(&lifecycle, "{sv}", "selected_strategy", g_variant_new_string("automatic-managed"));
  g_variant_builder_add(&lifecycle, "{sv}", "lifecycle_state", g_variant_new_string("blocked"));
  g_variant_builder_add(&lifecycle, "{sv}", "overall_status", g_variant_new_string("not-ready"));
  g_variant_builder_add(&lifecycle, "{sv}", "runtime_owned", g_variant_new_boolean(TRUE));
  g_variant_builder_add(&lifecycle, "{sv}", "c_runtime_backed", g_variant_new_boolean(TRUE));
  g_variant_builder_add(&lifecycle, "{sv}", "kde_policy_owner", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&lifecycle, "{sv}", "backend_binding_ready", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&lifecycle, "{sv}", "launch_enabled", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&lifecycle, "{sv}", "execution_request_created", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&lifecycle, "{sv}", "backend_process_started", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&lifecycle, "{sv}", "local_backend_started", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&lifecycle, "{sv}", "isolated_backend_started", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&lifecycle, "{sv}", "host_root_modified", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&lifecycle, "{sv}", "network_required", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&lifecycle, "{sv}", "backend_details_exposed", g_variant_new_boolean(FALSE));

  return g_variant_builder_end(&lifecycle);
}

static GVariant *
build_backend_environment_plan(const gchar *application_id)
{
  GVariantBuilder plan;

  g_variant_builder_init(&plan, G_VARIANT_TYPE("a{sv}"));
  g_variant_builder_add(&plan, "{sv}", "plan_type", g_variant_new_string("compatibility-backend-environment-plan"));
  add_go_owner_dispatch_bridge_fields(&plan, "GetBackendEnvironmentPlan", application_id);
  g_variant_builder_add(&plan, "{sv}", "runtime_method", g_variant_new_string("GetBackendEnvironmentPlan"));
  g_variant_builder_add(&plan, "{sv}", "application_id", g_variant_new_string(application_id));
  g_variant_builder_add(&plan, "{sv}", "selected_strategy", g_variant_new_string("automatic-managed"));
  g_variant_builder_add(&plan, "{sv}", "environment_state", g_variant_new_string("planned"));
  g_variant_builder_add(&plan, "{sv}", "runtime_owned", g_variant_new_boolean(TRUE));
  g_variant_builder_add(&plan, "{sv}", "c_runtime_backed", g_variant_new_boolean(TRUE));
  g_variant_builder_add(&plan, "{sv}", "kde_policy_owner", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&plan, "{sv}", "local_environment_ready", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&plan, "{sv}", "isolated_environment_ready", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&plan, "{sv}", "environment_created", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&plan, "{sv}", "backend_process_started", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&plan, "{sv}", "host_storage_exposed", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&plan, "{sv}", "clipboard_bridge_enabled", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&plan, "{sv}", "print_bridge_enabled", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&plan, "{sv}", "launch_enabled", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&plan, "{sv}", "host_root_modified", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&plan, "{sv}", "network_required", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&plan, "{sv}", "backend_details_exposed", g_variant_new_boolean(FALSE));

  return g_variant_builder_end(&plan);
}

static GVariant *
build_repair_plan(const gchar *application_id, const gchar *issue)
{
  GVariantBuilder plan;

  g_variant_builder_init(&plan, G_VARIANT_TYPE("a{sv}"));
  g_variant_builder_add(&plan, "{sv}", "plan_type", g_variant_new_string("compatibility-repair"));
  g_variant_builder_add(&plan, "{sv}", "application_id", g_variant_new_string(application_id));
  g_variant_builder_add(&plan, "{sv}", "issue", g_variant_new_string(issue));
  g_variant_builder_add(&plan, "{sv}", "runtime_method", g_variant_new_string("GetRepairPlan"));
  g_variant_builder_add(&plan, "{sv}", "runtime_owned", g_variant_new_boolean(TRUE));
  g_variant_builder_add(&plan, "{sv}", "c_runtime_backed", g_variant_new_boolean(TRUE));
  g_variant_builder_add(&plan, "{sv}", "kde_policy_owner", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&plan, "{sv}", "severity", g_variant_new_string("warning"));
  g_variant_builder_add(&plan, "{sv}", "user_approval_required", g_variant_new_boolean(TRUE));
  g_variant_builder_add(&plan, "{sv}", "snapshot_required", g_variant_new_boolean(TRUE));
  g_variant_builder_add(&plan, "{sv}", "rollback_available", g_variant_new_boolean(TRUE));
  g_variant_builder_add(&plan, "{sv}", "repair_execution_requested", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&plan, "{sv}", "repair_executed", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&plan, "{sv}", "backend_launch_enabled", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&plan, "{sv}", "network_required", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&plan, "{sv}", "host_root_modified", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&plan, "{sv}", "backend_details_exposed", g_variant_new_boolean(FALSE));

  return g_variant_builder_end(&plan);
}

static GVariant *
build_test_plan(const gchar *application_id, const gchar *test_type)
{
  GVariantBuilder plan;

  g_variant_builder_init(&plan, G_VARIANT_TYPE("a{sv}"));
  g_variant_builder_add(&plan, "{sv}", "plan_type", g_variant_new_string("compatibility-test"));
  g_variant_builder_add(&plan, "{sv}", "application_id", g_variant_new_string(application_id));
  g_variant_builder_add(&plan, "{sv}", "test_type", g_variant_new_string(test_type));
  g_variant_builder_add(&plan, "{sv}", "runtime_method", g_variant_new_string("GetTestPlan"));
  g_variant_builder_add(&plan, "{sv}", "runtime_owned", g_variant_new_boolean(TRUE));
  g_variant_builder_add(&plan, "{sv}", "c_runtime_backed", g_variant_new_boolean(TRUE));
  g_variant_builder_add(&plan, "{sv}", "kde_policy_owner", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&plan, "{sv}", "execution_request_created", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&plan, "{sv}", "test_executed", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&plan, "{sv}", "host_root_modified", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&plan, "{sv}", "backend_details_exposed", g_variant_new_boolean(FALSE));

  return g_variant_builder_end(&plan);
}

static GVariant *
build_test_result(const gchar *application_id, const gchar *test_type)
{
  GVariantBuilder result;

  g_variant_builder_init(&result, G_VARIANT_TYPE("a{sv}"));
  g_variant_builder_add(&result, "{sv}", "result_type", g_variant_new_string("compatibility-test-result"));
  g_variant_builder_add(&result, "{sv}", "application_id", g_variant_new_string(application_id));
  g_variant_builder_add(&result, "{sv}", "test_type", g_variant_new_string(test_type));
  g_variant_builder_add(&result, "{sv}", "runtime_method", g_variant_new_string("GetTestResult"));
  g_variant_builder_add(&result, "{sv}", "runtime_owned", g_variant_new_boolean(TRUE));
  g_variant_builder_add(&result, "{sv}", "c_runtime_backed", g_variant_new_boolean(TRUE));
  g_variant_builder_add(&result, "{sv}", "kde_policy_owner", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&result, "{sv}", "result_source", g_variant_new_string("runtime-model"));
  g_variant_builder_add(&result, "{sv}", "execution_state", g_variant_new_string("waiting-for-runtime"));
  g_variant_builder_add(&result, "{sv}", "overall_status", g_variant_new_string("pending"));
  g_variant_builder_add(&result, "{sv}", "safe_for_ai_diagnostics", g_variant_new_boolean(TRUE));
  g_variant_builder_add(&result, "{sv}", "test_executed", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&result, "{sv}", "host_root_modified", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&result, "{sv}", "backend_details_exposed", g_variant_new_boolean(FALSE));

  return g_variant_builder_end(&result);
}

static GVariant *
build_execution_readiness(const gchar *application_id)
{
  GVariantBuilder readiness;

  g_variant_builder_init(&readiness, G_VARIANT_TYPE("a{sv}"));
  g_variant_builder_add(&readiness, "{sv}", "readiness_type", g_variant_new_string("compatibility-execution-readiness"));
  g_variant_builder_add(&readiness, "{sv}", "application_id", g_variant_new_string(application_id));
  g_variant_builder_add(&readiness, "{sv}", "runtime_method", g_variant_new_string("GetExecutionReadiness"));
  g_variant_builder_add(&readiness, "{sv}", "runtime_owned", g_variant_new_boolean(TRUE));
  g_variant_builder_add(&readiness, "{sv}", "c_runtime_backed", g_variant_new_boolean(TRUE));
  g_variant_builder_add(&readiness, "{sv}", "kde_policy_owner", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&readiness, "{sv}", "execution_state", g_variant_new_string("blocked"));
  g_variant_builder_add(&readiness, "{sv}", "overall_status", g_variant_new_string("not-ready"));
  g_variant_builder_add(&readiness, "{sv}", "launch_allowed", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&readiness, "{sv}", "launch_enabled", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&readiness, "{sv}", "execution_request_created", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&readiness, "{sv}", "backend_binding_ready", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&readiness, "{sv}", "desktop_entry_launch_visible", g_variant_new_boolean(TRUE));
  g_variant_builder_add(&readiness, "{sv}", "safe_for_ai_diagnostics", g_variant_new_boolean(TRUE));
  g_variant_builder_add(&readiness, "{sv}", "host_root_modified", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&readiness, "{sv}", "backend_details_exposed", g_variant_new_boolean(FALSE));

  return g_variant_builder_end(&readiness);
}

static GVariant *
build_launch_intent(const gchar *application_id)
{
  GVariantBuilder intent;

  g_variant_builder_init(&intent, G_VARIANT_TYPE("a{sv}"));
  g_variant_builder_add(&intent, "{sv}", "intent_type", g_variant_new_string("runtime-launch-intent"));
  g_variant_builder_add(&intent, "{sv}", "source", g_variant_new_string("desktop-launcher"));
  g_variant_builder_add(&intent, "{sv}", "application_id", g_variant_new_string(application_id));
  g_variant_builder_add(&intent, "{sv}", "runtime_method", g_variant_new_string("Launch"));
  g_variant_builder_add(&intent, "{sv}", "read_method", g_variant_new_string("GetLaunchIntent"));
  g_variant_builder_add(&intent, "{sv}", "runtime_owned", g_variant_new_boolean(TRUE));
  g_variant_builder_add(&intent, "{sv}", "c_runtime_backed", g_variant_new_boolean(TRUE));
  g_variant_builder_add(&intent, "{sv}", "kde_policy_owner", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&intent, "{sv}", "standard_desktop_entry", g_variant_new_boolean(TRUE));
  g_variant_builder_add(&intent, "{sv}", "launch_uses_runtime", g_variant_new_boolean(TRUE));
  g_variant_builder_add(&intent, "{sv}", "desktop_entry_launch_visible", g_variant_new_boolean(TRUE));
  g_variant_builder_add(&intent, "{sv}", "launch_allowed", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&intent, "{sv}", "launch_enabled", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&intent, "{sv}", "execution_request_created", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&intent, "{sv}", "execution_started", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&intent, "{sv}", "write_gate_decision", g_variant_new_string("blocked-until-production-backend"));
  g_variant_builder_add(&intent, "{sv}", "denial_error_name", g_variant_new_string("org.xnix.Compatibility1.Error.WriteMethodDisabled"));
  g_variant_builder_add(&intent, "{sv}", "host_root_modified", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&intent, "{sv}", "network_required", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&intent, "{sv}", "backend_details_exposed", g_variant_new_boolean(FALSE));

  return g_variant_builder_end(&intent);
}

static GVariant *
build_ai_diagnostic_input(const gchar *application_id, const gchar *issue, const gchar *test_type)
{
  GVariantBuilder input;

  g_variant_builder_init(&input, G_VARIANT_TYPE("a{sv}"));
  g_variant_builder_add(&input, "{sv}", "input_type", g_variant_new_string("ai-diagnostic-input"));
  g_variant_builder_add(&input, "{sv}", "application_id", g_variant_new_string(application_id));
  g_variant_builder_add(&input, "{sv}", "issue", g_variant_new_string(issue));
  g_variant_builder_add(&input, "{sv}", "test_type", g_variant_new_string(test_type));
  g_variant_builder_add(&input, "{sv}", "runtime_method", g_variant_new_string("GetAIDiagnosticInput"));
  g_variant_builder_add(&input, "{sv}", "c_runtime_backed", g_variant_new_boolean(TRUE));
  g_variant_builder_add(&input, "{sv}", "diagnostic_signal_count", g_variant_new_int32(3));
  g_variant_builder_add(&input, "{sv}", "safe_for_ai_diagnostics", g_variant_new_boolean(TRUE));
  g_variant_builder_add(&input, "{sv}", "ai_provider_called", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&input, "{sv}", "network_required", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&input, "{sv}", "backend_details_exposed", g_variant_new_boolean(FALSE));

  return g_variant_builder_end(&input);
}

static GVariant *
build_ai_diagnostic_recommendation(const gchar *application_id, const gchar *issue, const gchar *test_type)
{
  GVariantBuilder recommendation;

  g_variant_builder_init(&recommendation, G_VARIANT_TYPE("a{sv}"));
  g_variant_builder_add(&recommendation, "{sv}", "recommendation_type", g_variant_new_string("ai-diagnostic-recommendation"));
  g_variant_builder_add(&recommendation, "{sv}", "application_id", g_variant_new_string(application_id));
  g_variant_builder_add(&recommendation, "{sv}", "issue", g_variant_new_string(issue));
  g_variant_builder_add(&recommendation, "{sv}", "test_type", g_variant_new_string(test_type));
  g_variant_builder_add(&recommendation, "{sv}", "runtime_method", g_variant_new_string("GetAIDiagnosticRecommendation"));
  g_variant_builder_add(&recommendation, "{sv}", "c_runtime_backed", g_variant_new_boolean(TRUE));
  g_variant_builder_add(&recommendation, "{sv}", "recommendation_count", g_variant_new_int32(3));
  g_variant_builder_add(&recommendation, "{sv}", "approval_required_count", g_variant_new_int32(1));
  g_variant_builder_add(&recommendation, "{sv}", "safe_for_ai_diagnostics", g_variant_new_boolean(TRUE));
  g_variant_builder_add(&recommendation, "{sv}", "ai_provider_called", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&recommendation, "{sv}", "network_required", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&recommendation, "{sv}", "auto_execution_allowed", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&recommendation, "{sv}", "backend_details_exposed", g_variant_new_boolean(FALSE));

  return g_variant_builder_end(&recommendation);
}

static GVariant *
build_ai_repair_approval_gate(const gchar *application_id, const gchar *issue, const gchar *test_type)
{
  GVariantBuilder gate;

  g_variant_builder_init(&gate, G_VARIANT_TYPE("a{sv}"));
  g_variant_builder_add(&gate, "{sv}", "gate_type", g_variant_new_string("ai-repair-approval-gate"));
  g_variant_builder_add(&gate, "{sv}", "application_id", g_variant_new_string(application_id));
  g_variant_builder_add(&gate, "{sv}", "issue", g_variant_new_string(issue));
  g_variant_builder_add(&gate, "{sv}", "test_type", g_variant_new_string(test_type));
  g_variant_builder_add(&gate, "{sv}", "runtime_method", g_variant_new_string("GetAIRepairApprovalGate"));
  g_variant_builder_add(&gate, "{sv}", "c_runtime_backed", g_variant_new_boolean(TRUE));
  g_variant_builder_add(&gate, "{sv}", "gate_decision", g_variant_new_string("blocked-until-approval"));
  g_variant_builder_add(&gate, "{sv}", "required_gate_count", g_variant_new_int32(3));
  g_variant_builder_add(&gate, "{sv}", "approval_required_count", g_variant_new_int32(1));
  g_variant_builder_add(&gate, "{sv}", "safe_for_ai_diagnostics", g_variant_new_boolean(TRUE));
  g_variant_builder_add(&gate, "{sv}", "ai_provider_called", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&gate, "{sv}", "network_required", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&gate, "{sv}", "auto_execution_allowed", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&gate, "{sv}", "repair_executed", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&gate, "{sv}", "backend_details_exposed", g_variant_new_boolean(FALSE));

  return g_variant_builder_end(&gate);
}

static GVariant *
build_snapshot_plan(const gchar *application_id, const gchar *reason)
{
  GVariantBuilder plan;

  g_variant_builder_init(&plan, G_VARIANT_TYPE("a{sv}"));
  g_variant_builder_add(&plan, "{sv}", "plan_type", g_variant_new_string("compatibility-snapshot"));
  g_variant_builder_add(&plan, "{sv}", "application_id", g_variant_new_string(application_id));
  g_variant_builder_add(&plan, "{sv}", "reason", g_variant_new_string(reason));
  g_variant_builder_add(&plan, "{sv}", "runtime_method", g_variant_new_string("GetSnapshotPlan"));
  g_variant_builder_add(&plan, "{sv}", "runtime_owned", g_variant_new_boolean(TRUE));
  g_variant_builder_add(&plan, "{sv}", "c_runtime_backed", g_variant_new_boolean(TRUE));
  g_variant_builder_add(&plan, "{sv}", "kde_policy_owner", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&plan, "{sv}", "enabled_by_default", g_variant_new_boolean(TRUE));
  g_variant_builder_add(&plan, "{sv}", "snapshot_request_created", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&plan, "{sv}", "snapshot_created", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&plan, "{sv}", "restore_requested", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&plan, "{sv}", "restore_executed", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&plan, "{sv}", "user_documents_included", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&plan, "{sv}", "host_system_included", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&plan, "{sv}", "host_root_modified", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&plan, "{sv}", "backend_details_exposed", g_variant_new_boolean(FALSE));

  return g_variant_builder_end(&plan);
}

static GVariant *
build_portal_policy(const gchar *application_id, const gchar *operation)
{
  GVariantBuilder policy;

  g_variant_builder_init(&policy, G_VARIANT_TYPE("a{sv}"));
  g_variant_builder_add(&policy, "{sv}", "policy_type", g_variant_new_string("portal-access"));
  g_variant_builder_add(&policy, "{sv}", "application_id", g_variant_new_string(application_id));
  g_variant_builder_add(&policy, "{sv}", "operation", g_variant_new_string(operation));
  g_variant_builder_add(&policy, "{sv}", "portal_required", g_variant_new_boolean(TRUE));

  return g_variant_builder_end(&policy);
}

static GVariant *
build_settings_model(const gchar *application_id)
{
  GVariantBuilder settings;

  g_variant_builder_init(&settings, G_VARIANT_TYPE("a{sv}"));
  g_variant_builder_add(&settings, "{sv}", "request_type", g_variant_new_string("settings-model"));
  add_go_owner_dispatch_bridge_fields(&settings, "GetCompatibilitySettings", application_id);
  g_variant_builder_add(&settings, "{sv}", "application_id", g_variant_new_string(application_id));
  g_variant_builder_add(&settings, "{sv}", "desktop", g_variant_new_string("KDE Plasma"));
  g_variant_builder_add(&settings, "{sv}", "runtime_owned", g_variant_new_boolean(TRUE));
  g_variant_builder_add(&settings, "{sv}", "kde_policy_owner", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&settings, "{sv}", "settings_state", g_variant_new_string("planned"));
  g_variant_builder_add(&settings, "{sv}", "settings_persisted", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&settings, "{sv}", "host_root_modified", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&settings, "{sv}", "backend_details_exposed", g_variant_new_boolean(FALSE));

  return g_variant_builder_end(&settings);
}

static GVariant *
build_settings_change_plan(const gchar *application_id,
                           const gchar *section_id,
                           const gchar *field_id,
                           const gchar *value)
{
  GVariantBuilder plan;

  g_variant_builder_init(&plan, G_VARIANT_TYPE("a{sv}"));
  g_variant_builder_add(&plan, "{sv}", "plan_type", g_variant_new_string("settings-change-plan"));
  g_variant_builder_add(&plan, "{sv}", "application_id", g_variant_new_string(application_id));
  g_variant_builder_add(&plan, "{sv}", "section_id", g_variant_new_string(section_id));
  g_variant_builder_add(&plan, "{sv}", "field_id", g_variant_new_string(field_id));
  g_variant_builder_add(&plan, "{sv}", "requested_value", g_variant_new_string(value));
  g_variant_builder_add(&plan, "{sv}", "change_state", g_variant_new_string("planned"));
  g_variant_builder_add(&plan, "{sv}", "apply_enabled", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&plan, "{sv}", "settings_persisted", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&plan, "{sv}", "portal_policy_review_required", g_variant_new_boolean(TRUE));
  g_variant_builder_add(&plan, "{sv}", "snapshot_recommended", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&plan, "{sv}", "backend_details_exposed", g_variant_new_boolean(FALSE));

  return g_variant_builder_end(&plan);
}

static gboolean
supported_compatibility_mode(const gchar *requested_mode)
{
  return g_strcmp0(requested_mode, "automatic") == 0 ||
         g_strcmp0(requested_mode, "prefer-performance") == 0 ||
         g_strcmp0(requested_mode, "prefer-compatibility") == 0 ||
         g_strcmp0(requested_mode, "isolated-execution") == 0;
}

static GVariant *
build_compatibility_mode_switch_plan(const gchar *application_id,
                                     const gchar *requested_mode)
{
  const gchar *mode_ids[] = {
    "automatic",
    "prefer-performance",
    "prefer-compatibility",
    "isolated-execution",
  };
  const gchar *required_runtime_gates[] = {
    "settings-review",
    "portal-policy-review",
    "snapshot-baseline",
    "backend-environment-plan",
    "runtime-write-gate",
  };
  GVariantBuilder plan;

  g_variant_builder_init(&plan, G_VARIANT_TYPE("a{sv}"));
  g_variant_builder_add(&plan, "{sv}", "plan_type", g_variant_new_string("compatibility-mode-switch-plan"));
  g_variant_builder_add(&plan, "{sv}", "runtime_method", g_variant_new_string("GetCompatibilityModeSwitchPlan"));
  g_variant_builder_add(&plan, "{sv}", "application_id", g_variant_new_string(application_id));
  g_variant_builder_add(&plan, "{sv}", "current_mode", g_variant_new_string("automatic"));
  g_variant_builder_add(&plan, "{sv}", "requested_mode", g_variant_new_string(requested_mode));
  g_variant_builder_add(&plan, "{sv}", "mode_state", g_variant_new_string("planned"));
  g_variant_builder_add(&plan, "{sv}", "mode_count", g_variant_new_int32(4));
  g_variant_builder_add(&plan, "{sv}", "mode_ids", g_variant_new_strv(mode_ids, 4));
  g_variant_builder_add(&plan, "{sv}", "required_runtime_gates", g_variant_new_strv(required_runtime_gates, 5));
  g_variant_builder_add(&plan, "{sv}", "runtime_owned", g_variant_new_boolean(TRUE));
  g_variant_builder_add(&plan, "{sv}", "c_runtime_backed", g_variant_new_boolean(TRUE));
  g_variant_builder_add(&plan, "{sv}", "kde_policy_owner", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&plan, "{sv}", "valid_mode", g_variant_new_boolean(TRUE));
  g_variant_builder_add(&plan, "{sv}", "requires_user_confirmation", g_variant_new_boolean(TRUE));
  g_variant_builder_add(&plan, "{sv}", "portal_review_required", g_variant_new_boolean(TRUE));
  g_variant_builder_add(&plan, "{sv}", "snapshot_required", g_variant_new_boolean(TRUE));
  g_variant_builder_add(&plan, "{sv}", "settings_persistence_enabled", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&plan, "{sv}", "backend_reconfiguration_enabled", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&plan, "{sv}", "backend_process_started", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&plan, "{sv}", "launch_enabled", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&plan, "{sv}", "host_root_modified", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&plan, "{sv}", "backend_details_exposed", g_variant_new_boolean(FALSE));

  return g_variant_builder_end(&plan);
}

static GVariant *
build_compatibility_permission_review_plan(const gchar *application_id)
{
  const gchar *permission_ids[] = {
    "documents",
    "downloads",
    "camera",
    "network",
    "clipboard",
    "print",
    "screenshot",
  };
  const gchar *required_runtime_gates[] = {
    "user-review",
    "portal-policy-review",
    "runtime-write-gate",
    "settings-persistence",
    "audit-log",
  };
  GVariantBuilder plan;

  g_variant_builder_init(&plan, G_VARIANT_TYPE("a{sv}"));
  g_variant_builder_add(&plan, "{sv}", "plan_type", g_variant_new_string("compatibility-permission-review-plan"));
  g_variant_builder_add(&plan, "{sv}", "runtime_method", g_variant_new_string("GetCompatibilityPermissionReviewPlan"));
  g_variant_builder_add(&plan, "{sv}", "application_id", g_variant_new_string(application_id));
  g_variant_builder_add(&plan, "{sv}", "review_state", g_variant_new_string("planned"));
  g_variant_builder_add(&plan, "{sv}", "permission_count", g_variant_new_int32(7));
  g_variant_builder_add(&plan, "{sv}", "allow_count", g_variant_new_int32(1));
  g_variant_builder_add(&plan, "{sv}", "ask_count", g_variant_new_int32(5));
  g_variant_builder_add(&plan, "{sv}", "deny_count", g_variant_new_int32(1));
  g_variant_builder_add(&plan, "{sv}", "permission_ids", g_variant_new_strv(permission_ids, 7));
  g_variant_builder_add(&plan, "{sv}", "required_runtime_gates", g_variant_new_strv(required_runtime_gates, 5));
  g_variant_builder_add(&plan, "{sv}", "runtime_owned", g_variant_new_boolean(TRUE));
  g_variant_builder_add(&plan, "{sv}", "c_runtime_backed", g_variant_new_boolean(TRUE));
  g_variant_builder_add(&plan, "{sv}", "kde_policy_owner", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&plan, "{sv}", "user_review_required", g_variant_new_boolean(TRUE));
  g_variant_builder_add(&plan, "{sv}", "portal_review_required", g_variant_new_boolean(TRUE));
  g_variant_builder_add(&plan, "{sv}", "permission_changes_applied", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&plan, "{sv}", "request_objects_created", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&plan, "{sv}", "permissions_granted", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&plan, "{sv}", "settings_persisted", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&plan, "{sv}", "host_permission_changed", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&plan, "{sv}", "host_root_modified", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&plan, "{sv}", "backend_details_exposed", g_variant_new_boolean(FALSE));

  return g_variant_builder_end(&plan);
}

static GVariant *
build_compatibility_review_flow_plan(
  const gchar *application_id,
  const gchar *section_id,
  const gchar *field_id,
  const gchar *value,
  const gchar *operation
)
{
  const gchar *step_ids[] = {
    "settings-change-review",
    "permission-review",
    "portal-request-review",
    "runtime-write-gate",
    "review-receipt",
  };
  const gchar *runtime_methods[] = {
    "GetCompatibilitySettingsChangePlan",
    "GetCompatibilityPermissionReviewPlan",
    "GetPortalRequestPlan",
    "GetRuntimeWriteGate",
    "GetCompatibilityActionReviewReceipt",
  };
  GVariantBuilder plan;

  g_variant_builder_init(&plan, G_VARIANT_TYPE("a{sv}"));
  g_variant_builder_add(&plan, "{sv}", "plan_type", g_variant_new_string("compatibility-review-flow-plan"));
  g_variant_builder_add(&plan, "{sv}", "runtime_method", g_variant_new_string("GetCompatibilityReviewFlowPlan"));
  g_variant_builder_add(&plan, "{sv}", "application_id", g_variant_new_string(application_id));
  g_variant_builder_add(&plan, "{sv}", "review_state", g_variant_new_string("planned"));
  g_variant_builder_add(&plan, "{sv}", "section_id", g_variant_new_string(section_id));
  g_variant_builder_add(&plan, "{sv}", "field_id", g_variant_new_string(field_id));
  g_variant_builder_add(&plan, "{sv}", "requested_value", g_variant_new_string(value));
  g_variant_builder_add(&plan, "{sv}", "operation", g_variant_new_string(operation));
  g_variant_builder_add(&plan, "{sv}", "step_ids", g_variant_new_strv(step_ids, 5));
  g_variant_builder_add(&plan, "{sv}", "runtime_methods", g_variant_new_strv(runtime_methods, 5));
  g_variant_builder_add(&plan, "{sv}", "step_count", g_variant_new_int32(5));
  g_variant_builder_add(&plan, "{sv}", "required_review_count", g_variant_new_int32(3));
  g_variant_builder_add(&plan, "{sv}", "blocked_step_count", g_variant_new_int32(1));
  g_variant_builder_add(&plan, "{sv}", "pending_step_count", g_variant_new_int32(1));
  g_variant_builder_add(&plan, "{sv}", "runtime_owned", g_variant_new_boolean(TRUE));
  g_variant_builder_add(&plan, "{sv}", "c_runtime_backed", g_variant_new_boolean(TRUE));
  g_variant_builder_add(&plan, "{sv}", "kde_policy_owner", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&plan, "{sv}", "user_confirmation_required", g_variant_new_boolean(TRUE));
  g_variant_builder_add(&plan, "{sv}", "portal_policy_review_required", g_variant_new_boolean(TRUE));
  g_variant_builder_add(&plan, "{sv}", "settings_change_planned", g_variant_new_boolean(TRUE));
  g_variant_builder_add(&plan, "{sv}", "permission_review_planned", g_variant_new_boolean(TRUE));
  g_variant_builder_add(&plan, "{sv}", "portal_request_planned", g_variant_new_boolean(TRUE));
  g_variant_builder_add(&plan, "{sv}", "review_receipt_required", g_variant_new_boolean(TRUE));
  g_variant_builder_add(&plan, "{sv}", "apply_enabled", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&plan, "{sv}", "request_object_created", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&plan, "{sv}", "permission_granted", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&plan, "{sv}", "settings_persisted", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&plan, "{sv}", "execution_started", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&plan, "{sv}", "host_root_modified", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&plan, "{sv}", "backend_details_exposed", g_variant_new_boolean(FALSE));

  return g_variant_builder_end(&plan);
}

static GVariant *
build_action_queue(const gchar *application_id)
{
  GVariantBuilder queue;

  g_variant_builder_init(&queue, G_VARIANT_TYPE("a{sv}"));
  g_variant_builder_add(&queue, "{sv}", "queue_type", g_variant_new_string("compatibility-center-action-queue"));
  g_variant_builder_add(&queue, "{sv}", "application_id", g_variant_new_string(application_id));
  g_variant_builder_add(&queue, "{sv}", "surface", g_variant_new_string("Compatibility Center"));
  g_variant_builder_add(&queue, "{sv}", "action_count", g_variant_new_int32(5));
  g_variant_builder_add(&queue, "{sv}", "pending_action_count", g_variant_new_int32(5));
  g_variant_builder_add(&queue, "{sv}", "user_review_required_count", g_variant_new_int32(3));
  g_variant_builder_add(&queue, "{sv}", "execution_enabled", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&queue, "{sv}", "repair_execution_enabled", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&queue, "{sv}", "settings_persistence_enabled", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&queue, "{sv}", "host_root_modified", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&queue, "{sv}", "backend_details_exposed", g_variant_new_boolean(FALSE));

  return g_variant_builder_end(&queue);
}

static GVariant *
build_action_review_receipt(const gchar *application_id, const gchar *action_id, const gchar *decision)
{
  GVariantBuilder receipt;

  g_variant_builder_init(&receipt, G_VARIANT_TYPE("a{sv}"));
  g_variant_builder_add(&receipt, "{sv}", "receipt_type", g_variant_new_string("compatibility-center-action-review-receipt"));
  g_variant_builder_add(&receipt, "{sv}", "application_id", g_variant_new_string(application_id));
  g_variant_builder_add(&receipt, "{sv}", "action_id", g_variant_new_string(action_id));
  g_variant_builder_add(&receipt, "{sv}", "decision", g_variant_new_string(decision));
  g_variant_builder_add(&receipt, "{sv}", "decision_recorded", g_variant_new_boolean(TRUE));
  g_variant_builder_add(&receipt, "{sv}", "execution_enabled", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&receipt, "{sv}", "repair_execution_enabled", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&receipt, "{sv}", "settings_persistence_enabled", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&receipt, "{sv}", "resource_grant_created", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&receipt, "{sv}", "host_root_modified", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&receipt, "{sv}", "backend_details_exposed", g_variant_new_boolean(FALSE));

  return g_variant_builder_end(&receipt);
}

static GVariant *
build_compatibility_center_summary(const gchar *application_id)
{
  GVariantBuilder summary;

  g_variant_builder_init(&summary, G_VARIANT_TYPE("a{sv}"));
  g_variant_builder_add(&summary, "{sv}", "summary_type", g_variant_new_string("compatibility-center-summary"));
  add_go_owner_dispatch_bridge_fields(&summary, "GetCompatibilityCenterSummary", application_id);
  g_variant_builder_add(&summary, "{sv}", "application_id", g_variant_new_string(application_id));
  g_variant_builder_add(&summary, "{sv}", "desktop", g_variant_new_string("KDE Plasma"));
  g_variant_builder_add(&summary, "{sv}", "compatibility_state", g_variant_new_string("review-required"));
  g_variant_builder_add(&summary, "{sv}", "runtime_mode", g_variant_new_string("automatic"));
  g_variant_builder_add(&summary, "{sv}", "known_issue_count", g_variant_new_int32(1));
  g_variant_builder_add(&summary, "{sv}", "repair_record_state", g_variant_new_string("pending-review"));
  g_variant_builder_add(&summary, "{sv}", "last_repair_event", g_variant_new_string("approval-required"));
  g_variant_builder_add(&summary, "{sv}", "action_count", g_variant_new_int32(4));
  g_variant_builder_add(&summary, "{sv}", "runtime_owned", g_variant_new_boolean(TRUE));
  g_variant_builder_add(&summary, "{sv}", "kde_policy_owner", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&summary, "{sv}", "user_visible", g_variant_new_boolean(TRUE));
  g_variant_builder_add(&summary, "{sv}", "action_execution_enabled", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&summary, "{sv}", "repair_execution_enabled", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&summary, "{sv}", "backend_launch_enabled", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&summary, "{sv}", "settings_persistence_enabled", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&summary, "{sv}", "host_root_modified", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&summary, "{sv}", "backend_details_exposed", g_variant_new_boolean(FALSE));

  return g_variant_builder_end(&summary);
}

static GVariant *
build_runtime_service_binding(void)
{
  GVariantBuilder binding;

  g_variant_builder_init(&binding, G_VARIANT_TYPE("a{sv}"));
  g_variant_builder_add(&binding, "{sv}", "binding_type", g_variant_new_string("runtime-service-binding"));
  add_go_owner_dispatch_bridge_fields(&binding, "GetRuntimeServiceBinding", NULL);
  g_variant_builder_add(&binding, "{sv}", "activation_binding_ready", g_variant_new_boolean(TRUE));
  g_variant_builder_add(&binding, "{sv}", "live_dbus_owner_ready", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&binding, "{sv}", "network_required", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&binding, "{sv}", "host_root_modified", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&binding, "{sv}", "backend_details_exposed", g_variant_new_boolean(FALSE));

  return g_variant_builder_end(&binding);
}

static GVariant *
build_runtime_live_owner_gate(void)
{
  GVariantBuilder gate;

  g_variant_builder_init(&gate, G_VARIANT_TYPE("a{sv}"));
  g_variant_builder_add(&gate, "{sv}", "gate_type", g_variant_new_string("runtime-live-owner-gate"));
  add_go_owner_dispatch_bridge_fields(&gate, "GetRuntimeLiveOwnerGate", NULL);
  g_variant_builder_add(&gate, "{sv}", "activation_binding_ready", g_variant_new_boolean(TRUE));
  g_variant_builder_add(&gate, "{sv}", "live_dbus_owner_ready", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&gate, "{sv}", "production_owner_enabled", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&gate, "{sv}", "owner_transition_ready", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&gate, "{sv}", "smoke_adapter_is_production_owner", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&gate, "{sv}", "kde_may_claim_runtime_ownership", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&gate, "{sv}", "host_root_modified", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&gate, "{sv}", "backend_details_exposed", g_variant_new_boolean(FALSE));

  return g_variant_builder_end(&gate);
}

static GVariant *
build_runtime_owner_smoke_plan(void)
{
  GVariantBuilder plan;

  g_variant_builder_init(&plan, G_VARIANT_TYPE("a{sv}"));
  g_variant_builder_add(&plan, "{sv}", "plan_type", g_variant_new_string("runtime-owner-smoke-plan"));
  add_go_owner_dispatch_bridge_fields(&plan, "GetRuntimeOwnerSmokePlan", NULL);
  g_variant_builder_add(&plan, "{sv}", "smoke_state", g_variant_new_string("planned"));
  g_variant_builder_add(&plan, "{sv}", "smoke_environment", g_variant_new_string("restricted-session"));
  g_variant_builder_add(&plan, "{sv}", "activation_binding_ready", g_variant_new_boolean(TRUE));
  g_variant_builder_add(&plan, "{sv}", "live_dbus_owner_ready", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&plan, "{sv}", "production_owner_enabled", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&plan, "{sv}", "owner_transition_ready", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&plan, "{sv}", "pending_step_count", g_variant_new_int32(6));
  g_variant_builder_add(&plan, "{sv}", "system_service_started", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&plan, "{sv}", "production_bus_claimed", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&plan, "{sv}", "host_root_modified", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&plan, "{sv}", "backend_details_exposed", g_variant_new_boolean(FALSE));

  return g_variant_builder_end(&plan);
}

static GVariant *
build_runtime_method_parity_manifest(void)
{
  GVariantBuilder manifest;

  g_variant_builder_init(&manifest, G_VARIANT_TYPE("a{sv}"));
  g_variant_builder_add(&manifest, "{sv}", "manifest_type", g_variant_new_string("runtime-method-parity-manifest"));
  add_go_owner_dispatch_bridge_fields(&manifest, "GetRuntimeMethodParityManifest", NULL);
  g_variant_builder_add(&manifest, "{sv}", "method_count", g_variant_new_int32(57));
  g_variant_builder_add(&manifest, "{sv}", "read_only_method_parity_ready", g_variant_new_boolean(TRUE));
  g_variant_builder_add(&manifest, "{sv}", "passed_check_count", g_variant_new_int32(5));
  g_variant_builder_add(&manifest, "{sv}", "blocked_check_count", g_variant_new_int32(0));
  g_variant_builder_add(&manifest, "{sv}", "write_methods_supported", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&manifest, "{sv}", "write_method_dispatch_enabled", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&manifest, "{sv}", "host_root_modified", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&manifest, "{sv}", "backend_details_exposed", g_variant_new_boolean(FALSE));

  return g_variant_builder_end(&manifest);
}

static GVariant *
build_runtime_write_gate(const gchar *method_name)
{
  GVariantBuilder gate;

  g_variant_builder_init(&gate, G_VARIANT_TYPE("a{sv}"));
  g_variant_builder_add(&gate, "{sv}", "gate_type", g_variant_new_string("runtime-write-gate"));
  g_variant_builder_add(&gate, "{sv}", "method_name", g_variant_new_string(method_name));
  g_variant_builder_add(&gate, "{sv}", "gate_decision", g_variant_new_string("blocked-until-production-backend"));
  add_go_owner_dispatch_bridge_fields(&gate, "GetRuntimeWriteGate", method_name);
  g_variant_builder_add(&gate, "{sv}", "write_method_enabled", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&gate, "{sv}", "dispatch_enabled", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&gate, "{sv}", "request_object_created", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&gate, "{sv}", "required_gate_count", g_variant_new_int32(6));
  g_variant_builder_add(&gate, "{sv}", "denial_error_name", g_variant_new_string("org.xnix.Compatibility1.Error.WriteMethodDisabled"));
  g_variant_builder_add(&gate, "{sv}", "host_root_modified", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&gate, "{sv}", "backend_details_exposed", g_variant_new_boolean(FALSE));

  return g_variant_builder_end(&gate);
}

static gboolean
is_write_method(const gchar *method_name)
{
  return g_strcmp0(method_name, "InstallRecipe") == 0 ||
         g_strcmp0(method_name, "Launch") == 0 ||
         g_strcmp0(method_name, "CreateSnapshot") == 0 ||
         g_strcmp0(method_name, "RestoreSnapshot") == 0;
}
