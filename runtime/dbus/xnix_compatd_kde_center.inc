static GVariant *
build_kde_center_page(const gchar *application_id, const gchar *decision)
{
  gboolean decision_allowed = g_strcmp0(decision, "approved") == 0 ||
                              g_strcmp0(decision, "reviewed") == 0;
  GVariantBuilder page;

  g_variant_builder_init(&page, G_VARIANT_TYPE("a{sv}"));
  g_variant_builder_add(&page, "{sv}", "request_type", g_variant_new_string("kde-center-page"));
  g_variant_builder_add(&page, "{sv}", "page_type", g_variant_new_string("compatibility-center-application-page"));
  g_variant_builder_add(&page, "{sv}", "desktop", g_variant_new_string("KDE Plasma"));
  g_variant_builder_add(&page, "{sv}", "runtime_method", g_variant_new_string("GetKDECenterPage"));
  g_variant_builder_add(&page, "{sv}", "read_model_source", g_variant_new_string("go-kde-center-page-preview"));
  g_variant_builder_add(&page, "{sv}", "application_id", g_variant_new_string(application_id));
  g_variant_builder_add(&page, "{sv}", "application_name", g_variant_new_string("Sample Notepad"));
  g_variant_builder_add(&page, "{sv}", "compatibility_state", g_variant_new_string("review-required"));
  g_variant_builder_add(&page, "{sv}", "runtime_mode", g_variant_new_string("automatic"));
  g_variant_builder_add(&page, "{sv}", "known_issue_count", g_variant_new_int32(1));
  g_variant_builder_add(&page, "{sv}", "repair_record_state", g_variant_new_string("pending-review"));
  g_variant_builder_add(&page, "{sv}", "backend_selection_plan_type", g_variant_new_string("compatibility-backend-selection-plan"));
  g_variant_builder_add(&page, "{sv}", "backend_selection_runtime_method", g_variant_new_string("GetBackendSelectionPlan"));
  g_variant_builder_add(&page, "{sv}", "backend_recommended_profile_id", g_variant_new_string("local-compatibility"));
  g_variant_builder_add(&page, "{sv}", "backend_candidate_count", g_variant_new_int32(2));
  g_variant_builder_add(&page, "{sv}", "backend_selection_committed", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&page, "{sv}", "backend_selection_change_enabled", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&page, "{sv}", "backend_launch_enabled", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&page, "{sv}", "activation_status_request_type", g_variant_new_string("desktop-activation-status-preview"));
  g_variant_builder_add(&page, "{sv}", "activation_state", g_variant_new_string("ready-for-runtime-commit"));
  g_variant_builder_add(&page, "{sv}", "activation_status_runtime_method", g_variant_new_string("GetDesktopActivationStatus"));
  g_variant_builder_add(&page, "{sv}", "activation_status_renderer", g_variant_new_string("xnix-runtime-go desktop-activation-status-preview"));
  g_variant_builder_add(&page, "{sv}", "activation_status_signal_count", g_variant_new_int32(5));
  g_variant_builder_add(&page, "{sv}", "activation_status_blocked_reason_count", g_variant_new_int32(4));
  g_variant_builder_add(&page, "{sv}", "activation_status_next_action_count", g_variant_new_int32(5));
  g_variant_builder_add(&page, "{sv}", "activation_commit_enabled", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&page, "{sv}", "activation_launch_enabled", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&page, "{sv}", "execution_readiness_type", g_variant_new_string("compatibility-execution-readiness"));
  g_variant_builder_add(&page, "{sv}", "execution_readiness_runtime_method", g_variant_new_string("GetExecutionReadiness"));
  g_variant_builder_add(&page, "{sv}", "execution_state", g_variant_new_string("blocked"));
  g_variant_builder_add(&page, "{sv}", "execution_overall_status", g_variant_new_string("not-ready"));
  g_variant_builder_add(&page, "{sv}", "execution_launch_allowed", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&page, "{sv}", "execution_launch_enabled", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&page, "{sv}", "execution_request_created", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&page, "{sv}", "execution_backend_binding_ready", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&page, "{sv}", "desktop_entry_launch_visible", g_variant_new_boolean(TRUE));
  g_variant_builder_add(&page, "{sv}", "launch_intent_type", g_variant_new_string("runtime-launch-intent"));
  g_variant_builder_add(&page, "{sv}", "launch_intent_source", g_variant_new_string("desktop-launcher"));
  g_variant_builder_add(&page, "{sv}", "launch_intent_runtime_method", g_variant_new_string("Launch"));
  g_variant_builder_add(&page, "{sv}", "launch_intent_read_method", g_variant_new_string("GetLaunchIntent"));
  g_variant_builder_add(&page, "{sv}", "launch_intent_allowed", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&page, "{sv}", "launch_intent_enabled", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&page, "{sv}", "launch_intent_request_created", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&page, "{sv}", "window_identity_runtime_method", g_variant_new_string("GetTaskManagerIdentityPlan"));
  g_variant_builder_add(&page, "{sv}", "window_identity_desktop_file", g_variant_new_string("xnix-org.xnix.sample.notepad.desktop"));
  g_variant_builder_add(&page, "{sv}", "window_identity_grouping_key", g_variant_new_string(application_id));
  g_variant_builder_add(&page, "{sv}", "window_identity_launcher_url", g_variant_new_string("applications:xnix-org.xnix.sample.notepad.desktop"));
  g_variant_builder_add(&page, "{sv}", "window_identity_pinning_allowed", g_variant_new_boolean(TRUE));
  g_variant_builder_add(&page, "{sv}", "window_identity_restore_allowed", g_variant_new_boolean(TRUE));
  g_variant_builder_add(&page, "{sv}", "window_identity_skip_taskbar", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&page, "{sv}", "window_identity_show_in_switcher", g_variant_new_boolean(TRUE));
  g_variant_builder_add(&page, "{sv}", "window_identity_prefer_existing_window", g_variant_new_boolean(TRUE));
  g_variant_builder_add(&page, "{sv}", "window_identity_kwin_role", g_variant_new_string("identity-and-layout"));
  g_variant_builder_add(&page, "{sv}", "window_identity_kwin_placement", g_variant_new_string("normal-window"));
  g_variant_builder_add(&page, "{sv}", "window_identity_policy_only", g_variant_new_boolean(TRUE));
  g_variant_builder_add(&page, "{sv}", "window_identity_backend_policy_owned", g_variant_new_boolean(TRUE));
  g_variant_builder_add(&page, "{sv}", "window_identity_task_manager_active", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&page, "{sv}", "window_identity_kwin_rule_applied", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&page, "{sv}", "file_association_runtime_method", g_variant_new_string("GetFileAssociationPlan"));
  g_variant_builder_add(&page, "{sv}", "file_association_plan_type", g_variant_new_string("file-association-plan"));
  g_variant_builder_add(&page, "{sv}", "file_association_type", g_variant_new_string("desktop-file-association"));
  g_variant_builder_add(&page, "{sv}", "file_association_desktop_file", g_variant_new_string("xnix-org.xnix.sample.notepad.desktop"));
  g_variant_builder_add(&page, "{sv}", "file_association_mimeapps_path", g_variant_new_string("usr/share/applications/mimeapps.list"));
  g_variant_builder_add(&page, "{sv}", "file_association_count", g_variant_new_int32(2));
  g_variant_builder_add(&page, "{sv}", "file_association_standard_mimeapps", g_variant_new_boolean(TRUE));
  g_variant_builder_add(&page, "{sv}", "file_association_staged_root_only", g_variant_new_boolean(TRUE));
  g_variant_builder_add(&page, "{sv}", "file_association_overwrite_existing", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&page, "{sv}", "file_association_portal_required", g_variant_new_boolean(TRUE));
  g_variant_builder_add(&page, "{sv}", "file_association_direct_host_file_access", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&page, "{sv}", "file_association_request_object_created", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&page, "{sv}", "file_association_permission_granted", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&page, "{sv}", "file_association_files_written", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&page, "{sv}", "tray_status_runtime_method", g_variant_new_string("GetTrayStatus"));
  g_variant_builder_add(&page, "{sv}", "tray_status_type", g_variant_new_string("tray-status-plan"));
  g_variant_builder_add(&page, "{sv}", "tray_status_active_application_count", g_variant_new_int32(1));
  g_variant_builder_add(&page, "{sv}", "tray_status_attention_required_count", g_variant_new_int32(1));
  g_variant_builder_add(&page, "{sv}", "tray_status_registered_application_count", g_variant_new_int32(1));
  g_variant_builder_add(&page, "{sv}", "tray_status_compatibility_state", g_variant_new_string("attention-required"));
  g_variant_builder_add(&page, "{sv}", "tray_status_bridge_state", g_variant_new_string("planned"));
  g_variant_builder_add(&page, "{sv}", "tray_status_bridged_application_count", g_variant_new_int32(0));
  g_variant_builder_add(&page, "{sv}", "tray_status_live_bridge_enabled", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&page, "{sv}", "tray_status_bridge_configuration_persisted", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&page, "{sv}", "tray_status_ready", g_variant_new_boolean(TRUE));
  g_variant_builder_add(&page, "{sv}", "notification_runtime_method", g_variant_new_string("GetNotificationPlan"));
  g_variant_builder_add(&page, "{sv}", "notification_plan_type", g_variant_new_string("notification-plan"));
  g_variant_builder_add(&page, "{sv}", "notification_event_type", g_variant_new_string("approval-required"));
  g_variant_builder_add(&page, "{sv}", "notification_id", g_variant_new_string("org.xnix.sample.notepad.approval-required"));
  g_variant_builder_add(&page, "{sv}", "notification_urgency", g_variant_new_string("critical"));
  g_variant_builder_add(&page, "{sv}", "notification_category", g_variant_new_string("compatibility.approval.required"));
  g_variant_builder_add(&page, "{sv}", "notification_title", g_variant_new_string("Sample Notepad needs approval"));
  g_variant_builder_add(&page, "{sv}", "notification_action_count", g_variant_new_int32(2));
  g_variant_builder_add(&page, "{sv}", "notification_requires_user_review", g_variant_new_boolean(TRUE));
  g_variant_builder_add(&page, "{sv}", "notification_action_execution_enabled", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&page, "{sv}", "notification_repair_execution_enabled", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&page, "{sv}", "notification_settings_persistence_enabled", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&page, "{sv}", "notification_sent", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&page, "{sv}", "notification_ready", g_variant_new_boolean(TRUE));
  g_variant_builder_add(&page, "{sv}", "action_deck_request_type", g_variant_new_string("kde-action-card-deck-preview"));
  g_variant_builder_add(&page, "{sv}", "card_count", g_variant_new_int32(7));
  g_variant_builder_add(&page, "{sv}", "waiting_card_count", g_variant_new_int32(decision_allowed ? 7 : 0));
  g_variant_builder_add(&page, "{sv}", "deferred_card_count", g_variant_new_int32(g_strcmp0(decision, "deferred") == 0 ? 7 : 0));
  g_variant_builder_add(&page, "{sv}", "rejected_card_count", g_variant_new_int32(g_strcmp0(decision, "rejected") == 0 ? 7 : 0));
  g_variant_builder_add(&page, "{sv}", "navigation_action_count", g_variant_new_int32(28));
  g_variant_builder_add(&page, "{sv}", "disabled_action_count", g_variant_new_int32(21));
  g_variant_builder_add(&page, "{sv}", "settings_request_type", g_variant_new_string("settings-model"));
  g_variant_builder_add(&page, "{sv}", "settings_section_count", g_variant_new_int32(5));
  g_variant_builder_add(&page, "{sv}", "primary_navigation_target", g_variant_new_string("compatibility-center-gates"));
  g_variant_builder_add(&page, "{sv}", "runtime_owned", g_variant_new_boolean(TRUE));
  g_variant_builder_add(&page, "{sv}", "go_runtime_backed", g_variant_new_boolean(TRUE));
  g_variant_builder_add(&page, "{sv}", "kde_policy_owner", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&page, "{sv}", "official_desktop_only", g_variant_new_boolean(TRUE));
  g_variant_builder_add(&page, "{sv}", "user_visible", g_variant_new_boolean(TRUE));
  g_variant_builder_add(&page, "{sv}", "safe_for_ai_diagnostics", g_variant_new_boolean(TRUE));
  g_variant_builder_add(&page, "{sv}", "user_decision_captured", g_variant_new_boolean(TRUE));
  g_variant_builder_add(&page, "{sv}", "user_decision_allows_launch", g_variant_new_boolean(decision_allowed));
  g_variant_builder_add(&page, "{sv}", "page_preview_created", g_variant_new_boolean(TRUE));
  g_variant_builder_add(&page, "{sv}", "page_persisted", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&page, "{sv}", "deck_persisted", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&page, "{sv}", "cards_persisted", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&page, "{sv}", "card_actions_enabled", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&page, "{sv}", "settings_persisted", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&page, "{sv}", "settings_persistence_enabled", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&page, "{sv}", "notifications_sent", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&page, "{sv}", "resource_grant_created", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&page, "{sv}", "runtime_launch_approval", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&page, "{sv}", "launch_allowed", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&page, "{sv}", "launch_enabled", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&page, "{sv}", "execution_started", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&page, "{sv}", "backend_process_started", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&page, "{sv}", "request_objects_created", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&page, "{sv}", "permission_grant_created", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&page, "{sv}", "host_root_modified", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&page, "{sv}", "network_required", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&page, "{sv}", "backend_details_exposed", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&page, "{sv}", "desktop_safe_summary", g_variant_new_string("KDE can read a Runtime-owned Compatibility Center page over D-Bus, but the page cannot approve, persist, grant, notify, or start execution."));

  return g_variant_builder_end(&page);
}

static void
add_kde_center_page_section(GVariantBuilder *sections,
                            const gchar *id,
                            const gchar *runtime_method,
                            const gchar *read_model,
                            const gchar *state)
{
  GVariantBuilder section;

  g_variant_builder_init(&section, G_VARIANT_TYPE("a{sv}"));
  g_variant_builder_add(&section, "{sv}", "id", g_variant_new_string(id));
  g_variant_builder_add(&section, "{sv}", "runtime_method", g_variant_new_string(runtime_method));
  g_variant_builder_add(&section, "{sv}", "read_model", g_variant_new_string(read_model));
  g_variant_builder_add(&section, "{sv}", "state", g_variant_new_string(state));
  g_variant_builder_add(&section, "{sv}", "navigation_only", g_variant_new_boolean(TRUE));
  g_variant_builder_add(&section, "{sv}", "read_only", g_variant_new_boolean(TRUE));
  g_variant_builder_add(&section, "{sv}", "mutates_runtime", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&section, "{sv}", "starts_program", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&section, "{sv}", "settings_persisted", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&section, "{sv}", "backend_details_exposed", g_variant_new_boolean(FALSE));
  g_variant_builder_add(sections, "a{sv}", &section);
}

static GVariant *
build_kde_center_page_sections(const gchar *application_id, const gchar *decision)
{
  gboolean decision_allowed = g_strcmp0(decision, "approved") == 0 ||
                              g_strcmp0(decision, "reviewed") == 0;
  static const gchar *section_ids[] = {
    "overview",
    "backend",
    "activation",
    "execution",
    "launch",
    "window",
    "files",
    "tray",
    "notifications",
    "actions",
    "settings",
    "diagnostics"
  };
  static const gchar *runtime_methods[] = {
    "GetCompatibilityCenterSummary",
    "GetBackendSelectionPlan",
    "GetDesktopActivationStatus",
    "GetExecutionReadiness",
    "GetLaunchIntent",
    "GetTaskManagerIdentityPlan",
    "GetFileAssociationPlan",
    "GetTrayStatus",
    "GetNotificationPlan",
    "GetCompatibilityActionQueue",
    "GetCompatibilitySettings",
    "GetDiagnostics"
  };
  static const gchar *read_models[] = {
    "compatibility-center-summary",
    "backend-selection-preview",
    "desktop-activation-status-preview",
    "execution-readiness-preview",
    "launch-intent-preview",
    "window-identity-preview",
    "file-association-plan",
    "tray-status-preview",
    "notification-preview",
    "compatibility-center-action-queue",
    "settings-model",
    "runtime-diagnostics"
  };
  static const gchar *section_states[] = {
    "ready",
    "selection-pending",
    "ready-for-runtime-commit",
    "blocked",
    "blocked",
    "planned",
    "planned",
    "planned",
    "planned",
    "waiting-for-runtime-gates",
    "planned",
    "planned"
  };
  GVariantBuilder sections;
  GVariantBuilder model;

  g_variant_builder_init(&sections, G_VARIANT_TYPE("aa{sv}"));
  add_kde_center_page_section(&sections, "overview", "GetCompatibilityCenterSummary", "compatibility-center-summary", "ready");
  add_kde_center_page_section(&sections, "backend", "GetBackendSelectionPlan", "backend-selection-preview", "selection-pending");
  add_kde_center_page_section(&sections, "activation", "GetDesktopActivationStatus", "desktop-activation-status-preview", "ready-for-runtime-commit");
  add_kde_center_page_section(&sections, "execution", "GetExecutionReadiness", "execution-readiness-preview", "blocked");
  add_kde_center_page_section(&sections, "launch", "GetLaunchIntent", "launch-intent-preview", "blocked");
  add_kde_center_page_section(&sections, "window", "GetTaskManagerIdentityPlan", "window-identity-preview", "planned");
  add_kde_center_page_section(&sections, "files", "GetFileAssociationPlan", "file-association-plan", "planned");
  add_kde_center_page_section(&sections, "tray", "GetTrayStatus", "tray-status-preview", "planned");
  add_kde_center_page_section(&sections, "notifications", "GetNotificationPlan", "notification-preview", "planned");
  add_kde_center_page_section(&sections, "actions", "GetCompatibilityActionQueue", "compatibility-center-action-queue", "waiting-for-runtime-gates");
  add_kde_center_page_section(&sections, "settings", "GetCompatibilitySettings", "settings-model", "planned");
  add_kde_center_page_section(&sections, "diagnostics", "GetDiagnostics", "runtime-diagnostics", "planned");

  g_variant_builder_init(&model, G_VARIANT_TYPE("a{sv}"));
  g_variant_builder_add(&model, "{sv}", "request_type", g_variant_new_string("kde-center-page-sections"));
  g_variant_builder_add(&model, "{sv}", "page_type", g_variant_new_string("compatibility-center-application-page"));
  g_variant_builder_add(&model, "{sv}", "desktop", g_variant_new_string("KDE Plasma"));
  g_variant_builder_add(&model, "{sv}", "runtime_method", g_variant_new_string("GetKDECenterPageSections"));
  g_variant_builder_add(&model, "{sv}", "read_model_source", g_variant_new_string("go-kde-center-page-sections-preview"));
  g_variant_builder_add(&model, "{sv}", "application_id", g_variant_new_string(application_id));
  g_variant_builder_add(&model, "{sv}", "section_count", g_variant_new_int32(12));
  g_variant_builder_add(&model, "{sv}", "read_only_section_count", g_variant_new_int32(12));
  g_variant_builder_add(&model, "{sv}", "navigation_only_section_count", g_variant_new_int32(12));
  g_variant_builder_add(&model, "{sv}", "executable_section_count", g_variant_new_int32(0));
  g_variant_builder_add(&model, "{sv}", "primary_section_id", g_variant_new_string("overview"));
  g_variant_builder_add(&model, "{sv}", "section_ids", g_variant_new_strv(section_ids, 12));
  g_variant_builder_add(&model, "{sv}", "runtime_methods", g_variant_new_strv(runtime_methods, 12));
  g_variant_builder_add(&model, "{sv}", "read_models", g_variant_new_strv(read_models, 12));
  g_variant_builder_add(&model, "{sv}", "section_states", g_variant_new_strv(section_states, 12));
  g_variant_builder_add(&model, "{sv}", "sections", g_variant_builder_end(&sections));
  g_variant_builder_add(&model, "{sv}", "runtime_owned", g_variant_new_boolean(TRUE));
  g_variant_builder_add(&model, "{sv}", "go_runtime_backed", g_variant_new_boolean(TRUE));
  g_variant_builder_add(&model, "{sv}", "kde_policy_owner", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&model, "{sv}", "official_desktop_only", g_variant_new_boolean(TRUE));
  g_variant_builder_add(&model, "{sv}", "user_visible", g_variant_new_boolean(TRUE));
  g_variant_builder_add(&model, "{sv}", "safe_for_ai_diagnostics", g_variant_new_boolean(TRUE));
  g_variant_builder_add(&model, "{sv}", "user_decision_captured", g_variant_new_boolean(TRUE));
  g_variant_builder_add(&model, "{sv}", "user_decision_allows_launch", g_variant_new_boolean(decision_allowed));
  g_variant_builder_add(&model, "{sv}", "sections_preview_created", g_variant_new_boolean(TRUE));
  g_variant_builder_add(&model, "{sv}", "sections_persisted", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&model, "{sv}", "section_actions_enabled", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&model, "{sv}", "settings_persisted", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&model, "{sv}", "settings_persistence_enabled", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&model, "{sv}", "notifications_sent", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&model, "{sv}", "resource_grant_created", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&model, "{sv}", "runtime_launch_approval", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&model, "{sv}", "launch_enabled", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&model, "{sv}", "execution_started", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&model, "{sv}", "request_objects_created", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&model, "{sv}", "permission_grant_created", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&model, "{sv}", "host_root_modified", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&model, "{sv}", "network_required", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&model, "{sv}", "backend_details_exposed", g_variant_new_boolean(FALSE));

  return g_variant_builder_end(&model);
}

static gboolean
kde_center_page_section_detail(const gchar *section_id,
                               const gchar **label,
                               const gchar **target,
                               const gchar **runtime_method,
                               const gchar **read_model,
                               const gchar **state)
{
  if (g_strcmp0(section_id, "overview") == 0) {
    *label = "Overview";
    *target = "compatibility-center-overview";
    *runtime_method = "GetCompatibilityCenterSummary";
    *read_model = "compatibility-center-summary";
    *state = "ready";
    return TRUE;
  }
  if (g_strcmp0(section_id, "backend") == 0) {
    *label = "Backend";
    *target = "compatibility-backend-selection";
    *runtime_method = "GetBackendSelectionPlan";
    *read_model = "backend-selection-preview";
    *state = "selection-pending";
    return TRUE;
  }
  if (g_strcmp0(section_id, "activation") == 0) {
    *label = "Activation";
    *target = "compatibility-activation-status";
    *runtime_method = "GetDesktopActivationStatus";
    *read_model = "desktop-activation-status-preview";
    *state = "ready-for-runtime-commit";
    return TRUE;
  }
  if (g_strcmp0(section_id, "execution") == 0) {
    *label = "Execution";
    *target = "compatibility-execution-readiness";
    *runtime_method = "GetExecutionReadiness";
    *read_model = "execution-readiness-preview";
    *state = "blocked";
    return TRUE;
  }
  if (g_strcmp0(section_id, "launch") == 0) {
    *label = "Launch";
    *target = "compatibility-launch-intent";
    *runtime_method = "GetLaunchIntent";
    *read_model = "launch-intent-preview";
    *state = "blocked";
    return TRUE;
  }
  if (g_strcmp0(section_id, "window") == 0) {
    *label = "Window";
    *target = "compatibility-window-identity";
    *runtime_method = "GetTaskManagerIdentityPlan";
    *read_model = "window-identity-preview";
    *state = "planned";
    return TRUE;
  }
  if (g_strcmp0(section_id, "files") == 0) {
    *label = "Files";
    *target = "compatibility-file-association";
    *runtime_method = "GetFileAssociationPlan";
    *read_model = "file-association-plan";
    *state = "planned";
    return TRUE;
  }
  if (g_strcmp0(section_id, "tray") == 0) {
    *label = "Tray";
    *target = "compatibility-tray-status";
    *runtime_method = "GetTrayStatus";
    *read_model = "tray-status-preview";
    *state = "planned";
    return TRUE;
  }
  if (g_strcmp0(section_id, "notifications") == 0) {
    *label = "Notifications";
    *target = "compatibility-notification-plan";
    *runtime_method = "GetNotificationPlan";
    *read_model = "notification-preview";
    *state = "planned";
    return TRUE;
  }
  if (g_strcmp0(section_id, "actions") == 0) {
    *label = "Actions";
    *target = "compatibility-center-gates";
    *runtime_method = "GetCompatibilityActionQueue";
    *read_model = "compatibility-center-action-queue";
    *state = "waiting-for-runtime-gates";
    return TRUE;
  }
  if (g_strcmp0(section_id, "settings") == 0) {
    *label = "Settings";
    *target = "compatibility-settings";
    *runtime_method = "GetCompatibilitySettings";
    *read_model = "settings-model";
    *state = "planned";
    return TRUE;
  }
  if (g_strcmp0(section_id, "diagnostics") == 0) {
    *label = "Diagnostics";
    *target = "compatibility-diagnostics";
    *runtime_method = "GetDiagnostics";
    *read_model = "runtime-diagnostics";
    *state = "planned";
    return TRUE;
  }
  return FALSE;
}

static GVariant *
build_kde_center_page_section_detail(const gchar *application_id,
                                     const gchar *section_id,
                                     const gchar *decision)
{
  gboolean decision_allowed = g_strcmp0(decision, "approved") == 0 ||
                              g_strcmp0(decision, "reviewed") == 0;
  static const gchar *available_section_ids[] = {
    "overview",
    "backend",
    "activation",
    "execution",
    "launch",
    "window",
    "files",
    "tray",
    "notifications",
    "actions",
    "settings",
    "diagnostics"
  };
  const gchar *label = NULL;
  const gchar *target = NULL;
  const gchar *runtime_method = NULL;
  const gchar *read_model = NULL;
  const gchar *state = NULL;
  GVariantBuilder detail;

  if (!kde_center_page_section_detail(section_id, &label, &target, &runtime_method, &read_model, &state)) {
    return NULL;
  }

  g_variant_builder_init(&detail, G_VARIANT_TYPE("a{sv}"));
  g_variant_builder_add(&detail, "{sv}", "request_type", g_variant_new_string("kde-center-page-section-detail"));
  g_variant_builder_add(&detail, "{sv}", "page_type", g_variant_new_string("compatibility-center-application-page"));
  g_variant_builder_add(&detail, "{sv}", "desktop", g_variant_new_string("KDE Plasma"));
  g_variant_builder_add(&detail, "{sv}", "runtime_method", g_variant_new_string("GetKDECenterPageSectionDetail"));
  g_variant_builder_add(&detail, "{sv}", "read_model_source", g_variant_new_string("go-kde-center-page-section-detail-preview"));
  g_variant_builder_add(&detail, "{sv}", "application_id", g_variant_new_string(application_id));
  g_variant_builder_add(&detail, "{sv}", "section_id", g_variant_new_string(section_id));
  g_variant_builder_add(&detail, "{sv}", "section_label", g_variant_new_string(label));
  g_variant_builder_add(&detail, "{sv}", "section_target", g_variant_new_string(target));
  g_variant_builder_add(&detail, "{sv}", "section_state", g_variant_new_string(state));
  g_variant_builder_add(&detail, "{sv}", "section_runtime_method", g_variant_new_string(runtime_method));
  g_variant_builder_add(&detail, "{sv}", "section_read_model", g_variant_new_string(read_model));
  g_variant_builder_add(&detail, "{sv}", "available_section_ids", g_variant_new_strv(available_section_ids, 12));
  g_variant_builder_add(&detail, "{sv}", "read_only_navigation", g_variant_new_boolean(TRUE));
  g_variant_builder_add(&detail, "{sv}", "detail_preview_created", g_variant_new_boolean(TRUE));
  g_variant_builder_add(&detail, "{sv}", "detail_persisted", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&detail, "{sv}", "section_actions_enabled", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&detail, "{sv}", "settings_persisted", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&detail, "{sv}", "settings_persistence_enabled", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&detail, "{sv}", "notifications_sent", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&detail, "{sv}", "resource_grant_created", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&detail, "{sv}", "runtime_launch_approval", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&detail, "{sv}", "launch_enabled", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&detail, "{sv}", "execution_started", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&detail, "{sv}", "request_objects_created", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&detail, "{sv}", "permission_grant_created", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&detail, "{sv}", "host_root_modified", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&detail, "{sv}", "network_required", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&detail, "{sv}", "backend_details_exposed", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&detail, "{sv}", "runtime_owned", g_variant_new_boolean(TRUE));
  g_variant_builder_add(&detail, "{sv}", "go_runtime_backed", g_variant_new_boolean(TRUE));
  g_variant_builder_add(&detail, "{sv}", "kde_policy_owner", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&detail, "{sv}", "official_desktop_only", g_variant_new_boolean(TRUE));
  g_variant_builder_add(&detail, "{sv}", "user_visible", g_variant_new_boolean(TRUE));
  g_variant_builder_add(&detail, "{sv}", "safe_for_ai_diagnostics", g_variant_new_boolean(TRUE));
  g_variant_builder_add(&detail, "{sv}", "user_decision_captured", g_variant_new_boolean(TRUE));
  g_variant_builder_add(&detail, "{sv}", "user_decision_allows_launch", g_variant_new_boolean(decision_allowed));

  return g_variant_builder_end(&detail);
}
