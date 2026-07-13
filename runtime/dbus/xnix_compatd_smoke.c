#include <gio/gio.h>
#include <string.h>

static const gchar introspection_xml[] =
  "<node>"
  "  <interface name='org.xnix.Compatibility1'>"
  "    <method name='ListApplications'>"
  "      <arg name='applications' type='aa{sv}' direction='out'/>"
  "    </method>"
  "    <method name='GetApplication'>"
  "      <arg name='application_id' type='s' direction='in'/>"
  "      <arg name='application' type='a{sv}' direction='out'/>"
  "    </method>"
  "    <method name='InstallRecipe'>"
  "      <arg name='recipe_uri' type='s' direction='in'/>"
  "      <arg name='request' type='o' direction='out'/>"
  "    </method>"
  "    <method name='Launch'>"
  "      <arg name='application_id' type='s' direction='in'/>"
  "      <arg name='options' type='a{sv}' direction='in'/>"
  "      <arg name='request' type='o' direction='out'/>"
  "    </method>"
  "    <method name='CreateSnapshot'>"
  "      <arg name='application_id' type='s' direction='in'/>"
  "      <arg name='label' type='s' direction='in'/>"
  "      <arg name='request' type='o' direction='out'/>"
  "    </method>"
  "    <method name='RestoreSnapshot'>"
  "      <arg name='snapshot_id' type='s' direction='in'/>"
  "      <arg name='request' type='o' direction='out'/>"
  "    </method>"
  "    <method name='GetDiagnostics'>"
  "      <arg name='application_id' type='s' direction='in'/>"
  "      <arg name='diagnostics' type='a{sv}' direction='out'/>"
  "    </method>"
  "    <method name='GetEngineCatalog'>"
  "      <arg name='catalog' type='a{sv}' direction='out'/>"
  "    </method>"
  "    <method name='GetRunPlan'>"
  "      <arg name='application_id' type='s' direction='in'/>"
  "      <arg name='plan' type='a{sv}' direction='out'/>"
  "    </method>"
  "    <method name='GetDesktopActivationManifest'>"
  "      <arg name='application_id' type='s' direction='in'/>"
  "      <arg name='manifest' type='a{sv}' direction='out'/>"
  "    </method>"
  "    <method name='GetDesktopEntryPlan'>"
  "      <arg name='application_id' type='s' direction='in'/>"
  "      <arg name='plan' type='a{sv}' direction='out'/>"
  "    </method>"
  "    <method name='GetTaskManagerIdentityPlan'>"
  "      <arg name='application_id' type='s' direction='in'/>"
  "      <arg name='plan' type='a{sv}' direction='out'/>"
  "    </method>"
  "    <method name='GetKDEIntegrationStatus'>"
  "      <arg name='status' type='a{sv}' direction='out'/>"
  "    </method>"
  "    <method name='GetKWinWindowRulePlan'>"
  "      <arg name='application_id' type='s' direction='in'/>"
  "      <arg name='plan' type='a{sv}' direction='out'/>"
  "    </method>"
  "    <method name='GetFileAssociationPlan'>"
  "      <arg name='application_id' type='s' direction='in'/>"
  "      <arg name='plan' type='a{sv}' direction='out'/>"
  "    </method>"
  "    <method name='GetNotificationPlan'>"
  "      <arg name='application_id' type='s' direction='in'/>"
  "      <arg name='event_type' type='s' direction='in'/>"
  "      <arg name='plan' type='a{sv}' direction='out'/>"
  "    </method>"
  "    <method name='GetTrayStatus'>"
  "      <arg name='status' type='a{sv}' direction='out'/>"
  "    </method>"
  "    <method name='GetKRunnerQueryPlan'>"
  "      <arg name='query' type='s' direction='in'/>"
  "      <arg name='plan' type='a{sv}' direction='out'/>"
  "    </method>"
  "    <method name='GetPortalRequestPlan'>"
  "      <arg name='application_id' type='s' direction='in'/>"
  "      <arg name='operation' type='s' direction='in'/>"
  "      <arg name='plan' type='a{sv}' direction='out'/>"
  "    </method>"
  "    <method name='GetApplicationStateRoot'>"
  "      <arg name='application_id' type='s' direction='in'/>"
  "      <arg name='state_root' type='a{sv}' direction='out'/>"
  "    </method>"
  "    <method name='GetCompatibilityPackageSource'>"
  "      <arg name='application_id' type='s' direction='in'/>"
  "      <arg name='source' type='a{sv}' direction='out'/>"
  "    </method>"
  "    <method name='GetCompatibilityAcquisitionPreflight'>"
  "      <arg name='application_id' type='s' direction='in'/>"
  "      <arg name='preflight' type='a{sv}' direction='out'/>"
  "    </method>"
  "    <method name='GetCompatibilityArtifactManifest'>"
  "      <arg name='application_id' type='s' direction='in'/>"
  "      <arg name='manifest' type='a{sv}' direction='out'/>"
  "    </method>"
  "    <method name='GetCompatibilityInstallPlan'>"
  "      <arg name='application_id' type='s' direction='in'/>"
  "      <arg name='environment' type='s' direction='in'/>"
  "      <arg name='plan' type='a{sv}' direction='out'/>"
  "    </method>"
  "    <method name='GetBackendBinding'>"
  "      <arg name='application_id' type='s' direction='in'/>"
  "      <arg name='binding' type='a{sv}' direction='out'/>"
  "    </method>"
  "    <method name='GetRepairPlan'>"
  "      <arg name='application_id' type='s' direction='in'/>"
  "      <arg name='issue' type='s' direction='in'/>"
  "      <arg name='plan' type='a{sv}' direction='out'/>"
  "    </method>"
  "    <method name='GetTestPlan'>"
  "      <arg name='application_id' type='s' direction='in'/>"
  "      <arg name='test_type' type='s' direction='in'/>"
  "      <arg name='plan' type='a{sv}' direction='out'/>"
  "    </method>"
  "    <method name='GetTestResult'>"
  "      <arg name='application_id' type='s' direction='in'/>"
  "      <arg name='test_type' type='s' direction='in'/>"
  "      <arg name='result' type='a{sv}' direction='out'/>"
  "    </method>"
  "    <method name='GetAIDiagnosticInput'>"
  "      <arg name='application_id' type='s' direction='in'/>"
  "      <arg name='issue' type='s' direction='in'/>"
  "      <arg name='test_type' type='s' direction='in'/>"
  "      <arg name='input' type='a{sv}' direction='out'/>"
  "    </method>"
  "    <method name='GetAIDiagnosticRecommendation'>"
  "      <arg name='application_id' type='s' direction='in'/>"
  "      <arg name='issue' type='s' direction='in'/>"
  "      <arg name='test_type' type='s' direction='in'/>"
  "      <arg name='recommendation' type='a{sv}' direction='out'/>"
  "    </method>"
  "    <method name='GetAIRepairApprovalGate'>"
  "      <arg name='application_id' type='s' direction='in'/>"
  "      <arg name='issue' type='s' direction='in'/>"
  "      <arg name='test_type' type='s' direction='in'/>"
  "      <arg name='gate' type='a{sv}' direction='out'/>"
  "    </method>"
  "    <method name='GetSnapshotPlan'>"
  "      <arg name='application_id' type='s' direction='in'/>"
  "      <arg name='reason' type='s' direction='in'/>"
  "      <arg name='plan' type='a{sv}' direction='out'/>"
  "    </method>"
  "    <method name='GetPortalAccessPolicy'>"
  "      <arg name='application_id' type='s' direction='in'/>"
  "      <arg name='operation' type='s' direction='in'/>"
  "      <arg name='policy' type='a{sv}' direction='out'/>"
  "    </method>"
  "    <method name='GetRuntimeServiceBinding'>"
  "      <arg name='binding' type='a{sv}' direction='out'/>"
  "    </method>"
  "    <method name='GetRuntimeLiveOwnerGate'>"
  "      <arg name='gate' type='a{sv}' direction='out'/>"
  "    </method>"
  "    <method name='GetRuntimeOwnerSmokePlan'>"
  "      <arg name='plan' type='a{sv}' direction='out'/>"
  "    </method>"
  "    <method name='GetRuntimeMethodParityManifest'>"
  "      <arg name='manifest' type='a{sv}' direction='out'/>"
  "    </method>"
  "    <method name='GetRuntimeWriteGate'>"
  "      <arg name='method_name' type='s' direction='in'/>"
  "      <arg name='gate' type='a{sv}' direction='out'/>"
  "    </method>"
  "    <method name='GetCompatibilitySettings'>"
  "      <arg name='application_id' type='s' direction='in'/>"
  "      <arg name='settings' type='a{sv}' direction='out'/>"
  "    </method>"
  "    <method name='GetCompatibilitySettingsChangePlan'>"
  "      <arg name='application_id' type='s' direction='in'/>"
  "      <arg name='section_id' type='s' direction='in'/>"
  "      <arg name='field_id' type='s' direction='in'/>"
  "      <arg name='value' type='s' direction='in'/>"
  "      <arg name='plan' type='a{sv}' direction='out'/>"
  "    </method>"
  "    <method name='GetCompatibilityActionQueue'>"
  "      <arg name='application_id' type='s' direction='in'/>"
  "      <arg name='queue' type='a{sv}' direction='out'/>"
  "    </method>"
  "    <method name='GetCompatibilityActionReviewReceipt'>"
  "      <arg name='application_id' type='s' direction='in'/>"
  "      <arg name='action_id' type='s' direction='in'/>"
  "      <arg name='decision' type='s' direction='in'/>"
  "      <arg name='receipt' type='a{sv}' direction='out'/>"
  "    </method>"
  "    <method name='GetCompatibilityCenterSummary'>"
  "      <arg name='application_id' type='s' direction='in'/>"
  "      <arg name='summary' type='a{sv}' direction='out'/>"
  "    </method>"
  "    <signal name='ApplicationChanged'>"
  "      <arg name='application_id' type='s'/>"
  "    </signal>"
  "    <signal name='RequestCompleted'>"
  "      <arg name='request' type='o'/>"
  "      <arg name='result' type='a{sv}'/>"
  "    </signal>"
  "  </interface>"
  "</node>";

static GDBusNodeInfo *introspection_data = NULL;

static GVariant *
build_application(void)
{
  GVariantBuilder builder;

  g_variant_builder_init(&builder, G_VARIANT_TYPE("a{sv}"));
  g_variant_builder_add(&builder, "{sv}", "id", g_variant_new_string("org.xnix.sample.notepad"));
  g_variant_builder_add(&builder, "{sv}", "name", g_variant_new_string("Sample Notepad"));
  g_variant_builder_add(&builder, "{sv}", "icon", g_variant_new_string("accessories-text-editor"));
  g_variant_builder_add(&builder, "{sv}", "mode", g_variant_new_string("automatic"));

  return g_variant_builder_end(&builder);
}

static gboolean
known_application(const gchar *application_id)
{
  return g_strcmp0(application_id, "org.xnix.sample.notepad") == 0;
}

static void
return_unknown_application(GDBusMethodInvocation *invocation, const gchar *application_id)
{
  g_dbus_method_invocation_return_error(invocation,
                                        G_IO_ERROR,
                                        G_IO_ERROR_NOT_FOUND,
                                        "unknown application: %s",
                                        application_id);
}

static GVariant *
build_engine_catalog(void)
{
  GVariantBuilder catalog;

  g_variant_builder_init(&catalog, G_VARIANT_TYPE("a{sv}"));
  g_variant_builder_add(&catalog, "{sv}", "catalog_type", g_variant_new_string("compatibility-engine"));
  g_variant_builder_add(&catalog, "{sv}", "runtime_policy_owner", g_variant_new_boolean(TRUE));
  g_variant_builder_add(&catalog, "{sv}", "backend_details_exposed", g_variant_new_boolean(FALSE));

  return g_variant_builder_end(&catalog);
}

static GVariant *
build_run_plan(const gchar *application_id)
{
  GVariantBuilder plan;

  g_variant_builder_init(&plan, G_VARIANT_TYPE("a{sv}"));
  g_variant_builder_add(&plan, "{sv}", "plan_type", g_variant_new_string("compatibility-run"));
  g_variant_builder_add(&plan, "{sv}", "application_id", g_variant_new_string(application_id));
  g_variant_builder_add(&plan, "{sv}", "strategy", g_variant_new_string("automatic-managed"));
  g_variant_builder_add(&plan, "{sv}", "backend_details_exposed", g_variant_new_boolean(FALSE));

  return g_variant_builder_end(&plan);
}

static GVariant *
build_desktop_activation_manifest(const gchar *application_id)
{
  GVariantBuilder manifest;

  g_variant_builder_init(&manifest, G_VARIANT_TYPE("a{sv}"));
  g_variant_builder_add(&manifest, "{sv}", "manifest_type", g_variant_new_string("desktop-activation"));
  g_variant_builder_add(&manifest, "{sv}", "application_id", g_variant_new_string(application_id));
  g_variant_builder_add(&manifest, "{sv}", "desktop", g_variant_new_string("KDE Plasma"));
  g_variant_builder_add(&manifest, "{sv}", "artifact_count", g_variant_new_int32(7));
  g_variant_builder_add(&manifest, "{sv}", "runtime_owned", g_variant_new_boolean(TRUE));
  g_variant_builder_add(&manifest, "{sv}", "kde_policy_owner", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&manifest, "{sv}", "activation_ready", g_variant_new_boolean(TRUE));
  g_variant_builder_add(&manifest, "{sv}", "files_written", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&manifest, "{sv}", "host_root_modified", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&manifest, "{sv}", "backend_details_exposed", g_variant_new_boolean(FALSE));

  return g_variant_builder_end(&manifest);
}

static GVariant *
build_kde_integration_status(void)
{
  static const gchar *entry_point_ids[] = {
    "launcher",
    "task-manager",
    "file-manager",
    "system-tray",
    "notifications",
    "compatibility-center",
    "settings"
  };
  static const gchar *entry_point_names[] = {
    "Launcher",
    "Task Manager",
    "File Manager",
    "System Tray",
    "Notifications",
    "Compatibility Center",
    "Settings"
  };
  static const gchar *entry_point_states[] = {
    "initial",
    "initial",
    "initial",
    "initial",
    "initial",
    "initial",
    "initial"
  };
  static const gchar *runtime_methods[] = {
    "GetDesktopEntryPlan",
    "GetTaskManagerIdentityPlan",
    "GetFileAssociationPlan",
    "GetTrayStatus",
    "GetNotificationPlan",
    "GetCompatibilityCenterSummary",
    "GetCompatibilitySettings"
  };
  GVariantBuilder status;

  g_variant_builder_init(&status, G_VARIANT_TYPE("a{sv}"));
  g_variant_builder_add(&status, "{sv}", "status_type", g_variant_new_string("kde-integration-status"));
  g_variant_builder_add(&status, "{sv}", "desktop", g_variant_new_string("KDE Plasma"));
  g_variant_builder_add(&status, "{sv}", "entry_point_count", g_variant_new_int32(7));
  g_variant_builder_add(&status, "{sv}", "initial_count", g_variant_new_int32(7));
  g_variant_builder_add(&status, "{sv}", "planned_count", g_variant_new_int32(0));
  g_variant_builder_add(&status, "{sv}", "complete_count", g_variant_new_int32(0));
  g_variant_builder_add(&status, "{sv}", "entry_point_ids", g_variant_new_strv(entry_point_ids, 7));
  g_variant_builder_add(&status, "{sv}", "entry_point_names", g_variant_new_strv(entry_point_names, 7));
  g_variant_builder_add(&status, "{sv}", "entry_point_states", g_variant_new_strv(entry_point_states, 7));
  g_variant_builder_add(&status, "{sv}", "runtime_methods", g_variant_new_strv(runtime_methods, 7));
  g_variant_builder_add(&status, "{sv}", "runtime_owned", g_variant_new_boolean(TRUE));
  g_variant_builder_add(&status, "{sv}", "kde_policy_owner", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&status, "{sv}", "official_desktop_only", g_variant_new_boolean(TRUE));
  g_variant_builder_add(&status, "{sv}", "stable_desktop_contract", g_variant_new_boolean(TRUE));
  g_variant_builder_add(&status, "{sv}", "host_root_modified", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&status, "{sv}", "backend_details_exposed", g_variant_new_boolean(FALSE));

  return g_variant_builder_end(&status);
}

static GVariant *
build_desktop_entry_plan(const gchar *application_id)
{
  GVariantBuilder plan;

  g_variant_builder_init(&plan, G_VARIANT_TYPE("a{sv}"));
  g_variant_builder_add(&plan, "{sv}", "plan_type", g_variant_new_string("desktop-entry-plan"));
  g_variant_builder_add(&plan, "{sv}", "application_id", g_variant_new_string(application_id));
  g_variant_builder_add(&plan, "{sv}", "desktop", g_variant_new_string("KDE Plasma"));
  g_variant_builder_add(&plan, "{sv}", "desktop_file", g_variant_new_string("xnix-org.xnix.sample.notepad.desktop"));
  g_variant_builder_add(&plan, "{sv}", "name", g_variant_new_string("Sample Notepad"));
  g_variant_builder_add(&plan, "{sv}", "exec", g_variant_new_string("xnix-compat-launch --app org.xnix.sample.notepad %U"));
  g_variant_builder_add(&plan, "{sv}", "standard_desktop_entry", g_variant_new_boolean(TRUE));
  g_variant_builder_add(&plan, "{sv}", "launch_uses_runtime", g_variant_new_boolean(TRUE));
  g_variant_builder_add(&plan, "{sv}", "accepts_file_uris", g_variant_new_boolean(TRUE));
  g_variant_builder_add(&plan, "{sv}", "backend_command_exposed", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&plan, "{sv}", "raw_windows_executable_exposed", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&plan, "{sv}", "compatibility_storage_path_exposed", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&plan, "{sv}", "files_written", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&plan, "{sv}", "host_root_modified", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&plan, "{sv}", "backend_details_exposed", g_variant_new_boolean(FALSE));

  return g_variant_builder_end(&plan);
}

static GVariant *
build_task_manager_identity_plan(const gchar *application_id)
{
  GVariantBuilder plan;

  g_variant_builder_init(&plan, G_VARIANT_TYPE("a{sv}"));
  g_variant_builder_add(&plan, "{sv}", "plan_type", g_variant_new_string("task-manager-identity-plan"));
  g_variant_builder_add(&plan, "{sv}", "application_id", g_variant_new_string(application_id));
  g_variant_builder_add(&plan, "{sv}", "desktop", g_variant_new_string("KDE Plasma"));
  g_variant_builder_add(&plan, "{sv}", "desktop_file", g_variant_new_string("xnix-org.xnix.sample.notepad.desktop"));
  g_variant_builder_add(&plan, "{sv}", "launcher_url", g_variant_new_string("applications:xnix-org.xnix.sample.notepad.desktop"));
  g_variant_builder_add(&plan, "{sv}", "window_kind", g_variant_new_string("compatibility-application"));
  g_variant_builder_add(&plan, "{sv}", "class_group", g_variant_new_string("xnix-compatibility"));
  g_variant_builder_add(&plan, "{sv}", "resource_name", g_variant_new_string(application_id));
  g_variant_builder_add(&plan, "{sv}", "grouping_key", g_variant_new_string(application_id));
  g_variant_builder_add(&plan, "{sv}", "pinning_allowed", g_variant_new_boolean(TRUE));
  g_variant_builder_add(&plan, "{sv}", "restore_allowed", g_variant_new_boolean(TRUE));
  g_variant_builder_add(&plan, "{sv}", "skip_taskbar", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&plan, "{sv}", "show_in_switcher", g_variant_new_boolean(TRUE));
  g_variant_builder_add(&plan, "{sv}", "prefer_existing_window", g_variant_new_boolean(TRUE));
  g_variant_builder_add(&plan, "{sv}", "window_manager_policy_only", g_variant_new_boolean(TRUE));
  g_variant_builder_add(&plan, "{sv}", "runtime_owns_backend_policy", g_variant_new_boolean(TRUE));
  g_variant_builder_add(&plan, "{sv}", "host_root_modified", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&plan, "{sv}", "backend_details_exposed", g_variant_new_boolean(FALSE));

  return g_variant_builder_end(&plan);
}

static GVariant *
build_kwin_window_rule_plan(const gchar *application_id)
{
  GVariantBuilder plan;

  g_variant_builder_init(&plan, G_VARIANT_TYPE("a{sv}"));
  g_variant_builder_add(&plan, "{sv}", "request_type", g_variant_new_string("kwin-window-rule"));
  g_variant_builder_add(&plan, "{sv}", "desktop", g_variant_new_string("KDE Plasma"));
  g_variant_builder_add(&plan, "{sv}", "runtime_owned", g_variant_new_boolean(TRUE));
  g_variant_builder_add(&plan, "{sv}", "kde_policy_owner", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&plan, "{sv}", "application_id", g_variant_new_string(application_id));
  g_variant_builder_add(&plan, "{sv}", "name", g_variant_new_string("Sample Notepad"));
  g_variant_builder_add(&plan, "{sv}", "script_role", g_variant_new_string("identity-and-layout"));
  g_variant_builder_add(&plan, "{sv}", "resource_name", g_variant_new_string(application_id));
  g_variant_builder_add(&plan, "{sv}", "class_group", g_variant_new_string("xnix-compatibility"));
  g_variant_builder_add(&plan, "{sv}", "title_hint", g_variant_new_string("Sample Notepad"));
  g_variant_builder_add(&plan, "{sv}", "desktop_file", g_variant_new_string("xnix-org.xnix.sample.notepad.desktop"));
  g_variant_builder_add(&plan, "{sv}", "task_manager_grouping_key", g_variant_new_string(application_id));
  g_variant_builder_add(&plan, "{sv}", "launcher_url", g_variant_new_string("applications:xnix-org.xnix.sample.notepad.desktop"));
  g_variant_builder_add(&plan, "{sv}", "skip_taskbar", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&plan, "{sv}", "show_in_switcher", g_variant_new_boolean(TRUE));
  g_variant_builder_add(&plan, "{sv}", "placement", g_variant_new_string("normal-window"));
  g_variant_builder_add(&plan, "{sv}", "pinning_allowed", g_variant_new_boolean(TRUE));
  g_variant_builder_add(&plan, "{sv}", "restore_allowed", g_variant_new_boolean(TRUE));
  g_variant_builder_add(&plan, "{sv}", "restore_key", g_variant_new_string(application_id));
  g_variant_builder_add(&plan, "{sv}", "prefer_existing_window", g_variant_new_boolean(TRUE));
  g_variant_builder_add(&plan, "{sv}", "window_manager_policy_only", g_variant_new_boolean(TRUE));
  g_variant_builder_add(&plan, "{sv}", "runtime_owns_backend_policy", g_variant_new_boolean(TRUE));
  g_variant_builder_add(&plan, "{sv}", "host_root_modified", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&plan, "{sv}", "backend_details_exposed", g_variant_new_boolean(FALSE));

  return g_variant_builder_end(&plan);
}

static GVariant *
build_file_association_plan(const gchar *application_id)
{
  GVariantBuilder plan;

  g_variant_builder_init(&plan, G_VARIANT_TYPE("a{sv}"));
  g_variant_builder_add(&plan, "{sv}", "plan_type", g_variant_new_string("file-association-plan"));
  g_variant_builder_add(&plan, "{sv}", "association_type", g_variant_new_string("desktop-file-association"));
  g_variant_builder_add(&plan, "{sv}", "application_id", g_variant_new_string(application_id));
  g_variant_builder_add(&plan, "{sv}", "desktop", g_variant_new_string("KDE Plasma"));
  g_variant_builder_add(&plan, "{sv}", "desktop_file", g_variant_new_string("xnix-org.xnix.sample.notepad.desktop"));
  g_variant_builder_add(&plan, "{sv}", "mimeapps_path", g_variant_new_string("usr/share/applications/mimeapps.list"));
  g_variant_builder_add(&plan, "{sv}", "file_open_command", g_variant_new_string("xnix-compat-open"));
  g_variant_builder_add(&plan, "{sv}", "file_open_argument", g_variant_new_string("%U"));
  g_variant_builder_add(&plan, "{sv}", "mime_type_count", g_variant_new_int32(2));
  g_variant_builder_add(&plan, "{sv}", "standard_mimeapps_list", g_variant_new_boolean(TRUE));
  g_variant_builder_add(&plan, "{sv}", "staged_root_only", g_variant_new_boolean(TRUE));
  g_variant_builder_add(&plan, "{sv}", "overwrite_existing_mimeapps", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&plan, "{sv}", "portal_required_for_file_open", g_variant_new_boolean(TRUE));
  g_variant_builder_add(&plan, "{sv}", "files_written", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&plan, "{sv}", "host_root_modified", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&plan, "{sv}", "backend_details_exposed", g_variant_new_boolean(FALSE));

  return g_variant_builder_end(&plan);
}

static GVariant *
build_notification_plan(const gchar *application_id, const gchar *event_type)
{
  GVariantBuilder plan;
  const gboolean review_required =
    g_strcmp0(event_type, "approval-required") == 0 ||
    g_strcmp0(event_type, "install-failed") == 0;

  g_variant_builder_init(&plan, G_VARIANT_TYPE("a{sv}"));
  g_variant_builder_add(&plan, "{sv}", "plan_type", g_variant_new_string("notification-plan"));
  g_variant_builder_add(&plan, "{sv}", "application_id", g_variant_new_string(application_id));
  g_variant_builder_add(&plan, "{sv}", "desktop", g_variant_new_string("KDE Plasma"));
  g_variant_builder_add(&plan, "{sv}", "event_type", g_variant_new_string(event_type));
  g_variant_builder_add(&plan, "{sv}", "notification_id", g_variant_new_string("org.xnix.sample.notepad.approval-required"));
  g_variant_builder_add(&plan, "{sv}", "title", g_variant_new_string("Sample Notepad needs approval"));
  g_variant_builder_add(&plan, "{sv}", "urgency", g_variant_new_string(review_required ? "critical" : "normal"));
  g_variant_builder_add(&plan, "{sv}", "category", g_variant_new_string("compatibility.approval"));
  g_variant_builder_add(&plan, "{sv}", "desktop_entry", g_variant_new_string("xnix-org.xnix.sample.notepad.desktop"));
  g_variant_builder_add(&plan, "{sv}", "action_count", g_variant_new_int32(2));
  g_variant_builder_add(&plan, "{sv}", "runtime_owned", g_variant_new_boolean(TRUE));
  g_variant_builder_add(&plan, "{sv}", "kde_policy_owner", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&plan, "{sv}", "user_visible", g_variant_new_boolean(TRUE));
  g_variant_builder_add(&plan, "{sv}", "requires_user_review", g_variant_new_boolean(review_required));
  g_variant_builder_add(&plan, "{sv}", "action_execution_enabled", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&plan, "{sv}", "repair_execution_enabled", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&plan, "{sv}", "settings_persistence_enabled", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&plan, "{sv}", "host_root_modified", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&plan, "{sv}", "backend_details_exposed", g_variant_new_boolean(FALSE));

  return g_variant_builder_end(&plan);
}

static GVariant *
build_tray_status(void)
{
  GVariantBuilder status;

  g_variant_builder_init(&status, G_VARIANT_TYPE("a{sv}"));
  g_variant_builder_add(&status, "{sv}", "status_type", g_variant_new_string("tray-status-plan"));
  g_variant_builder_add(&status, "{sv}", "desktop", g_variant_new_string("KDE Plasma"));
  g_variant_builder_add(&status, "{sv}", "active_application_count", g_variant_new_int32(1));
  g_variant_builder_add(&status, "{sv}", "attention_required_count", g_variant_new_int32(1));
  g_variant_builder_add(&status, "{sv}", "bridged_tray_application_count", g_variant_new_int32(0));
  g_variant_builder_add(&status, "{sv}", "compatibility_state", g_variant_new_string("attention-required"));
  g_variant_builder_add(&status, "{sv}", "tray_bridge_state", g_variant_new_string("planned"));
  g_variant_builder_add(&status, "{sv}", "runtime_owned", g_variant_new_boolean(TRUE));
  g_variant_builder_add(&status, "{sv}", "kde_policy_owner", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&status, "{sv}", "user_visible", g_variant_new_boolean(TRUE));
  g_variant_builder_add(&status, "{sv}", "live_backend_bridge_enabled", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&status, "{sv}", "bridge_configuration_persisted", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&status, "{sv}", "host_root_modified", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&status, "{sv}", "backend_details_exposed", g_variant_new_boolean(FALSE));

  return g_variant_builder_end(&status);
}

static GVariant *
build_krunner_query_plan(const gchar *query)
{
  GVariantBuilder plan;
  gboolean has_match = FALSE;

  if (query != NULL) {
    gchar *normalized = g_ascii_strdown(query, -1);
    has_match = g_strrstr(normalized, "notepad") != NULL ||
                g_strcmp0(normalized, "open txt") == 0 ||
                g_strcmp0(normalized, "txt") == 0;
    g_free(normalized);
  }

  g_variant_builder_init(&plan, G_VARIANT_TYPE("a{sv}"));
  g_variant_builder_add(&plan, "{sv}", "query_type", g_variant_new_string("krunner-query-plan"));
  g_variant_builder_add(&plan, "{sv}", "entry_point", g_variant_new_string("krunner"));
  g_variant_builder_add(&plan, "{sv}", "desktop", g_variant_new_string("KDE Plasma"));
  g_variant_builder_add(&plan, "{sv}", "query", g_variant_new_string(query == NULL ? "" : query));
  g_variant_builder_add(&plan, "{sv}", "match_count", g_variant_new_int32(has_match ? 1 : 0));
  g_variant_builder_add(&plan, "{sv}", "top_application_id", g_variant_new_string(has_match ? "org.xnix.sample.notepad" : ""));
  g_variant_builder_add(&plan, "{sv}", "top_name", g_variant_new_string(has_match ? "Sample Notepad" : ""));
  g_variant_builder_add(&plan, "{sv}", "top_relevance_percent", g_variant_new_int32(has_match ? 100 : 0));
  g_variant_builder_add(&plan, "{sv}", "action_type", g_variant_new_string(has_match ? "runtime-launch" : ""));
  g_variant_builder_add(&plan, "{sv}", "desktop_entry_id", g_variant_new_string(has_match ? "org.xnix.sample.notepad.desktop" : ""));
  g_variant_builder_add(&plan, "{sv}", "runtime_owned", g_variant_new_boolean(TRUE));
  g_variant_builder_add(&plan, "{sv}", "kde_policy_owner", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&plan, "{sv}", "runtime_owned_launch", g_variant_new_boolean(TRUE));
  g_variant_builder_add(&plan, "{sv}", "query_execution_enabled", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&plan, "{sv}", "backend_launch_enabled", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&plan, "{sv}", "host_root_modified", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&plan, "{sv}", "backend_details_exposed", g_variant_new_boolean(FALSE));

  return g_variant_builder_end(&plan);
}

static GVariant *
build_portal_request_plan(const gchar *application_id, const gchar *operation)
{
  GVariantBuilder plan;
  gboolean denied = g_strcmp0(operation, "camera") == 0 ||
                    g_strcmp0(operation, "remote-desktop") == 0;

  g_variant_builder_init(&plan, G_VARIANT_TYPE("a{sv}"));
  g_variant_builder_add(&plan, "{sv}", "request_type", g_variant_new_string("portal-request-plan"));
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
  g_variant_builder_add(&binding, "{sv}", "application_id", g_variant_new_string(application_id));
  g_variant_builder_add(&binding, "{sv}", "selected_strategy", g_variant_new_string("automatic-managed"));
  g_variant_builder_add(&binding, "{sv}", "managed_binding_ready", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&binding, "{sv}", "launch_enabled", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&binding, "{sv}", "backend_details_exposed", g_variant_new_boolean(FALSE));

  return g_variant_builder_end(&binding);
}

static GVariant *
build_repair_plan(const gchar *application_id, const gchar *issue)
{
  GVariantBuilder plan;

  g_variant_builder_init(&plan, G_VARIANT_TYPE("a{sv}"));
  g_variant_builder_add(&plan, "{sv}", "plan_type", g_variant_new_string("compatibility-repair"));
  g_variant_builder_add(&plan, "{sv}", "application_id", g_variant_new_string(application_id));
  g_variant_builder_add(&plan, "{sv}", "issue", g_variant_new_string(issue));
  g_variant_builder_add(&plan, "{sv}", "snapshot_required", g_variant_new_boolean(TRUE));

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
  g_variant_builder_add(&plan, "{sv}", "runtime_owned", g_variant_new_boolean(TRUE));
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
  g_variant_builder_add(&result, "{sv}", "overall_status", g_variant_new_string("pending"));
  g_variant_builder_add(&result, "{sv}", "safe_for_ai_diagnostics", g_variant_new_boolean(TRUE));
  g_variant_builder_add(&result, "{sv}", "backend_details_exposed", g_variant_new_boolean(FALSE));

  return g_variant_builder_end(&result);
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
  g_variant_builder_add(&gate, "{sv}", "gate_decision", g_variant_new_string("blocked-until-approval"));
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
  g_variant_builder_add(&plan, "{sv}", "enabled_by_default", g_variant_new_boolean(TRUE));

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
  g_variant_builder_add(&manifest, "{sv}", "method_count", g_variant_new_int32(39));
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

static void
handle_method_call(GDBusConnection *connection,
                   const gchar *sender,
                   const gchar *object_path,
                   const gchar *interface_name,
                   const gchar *method_name,
                   GVariant *parameters,
                   GDBusMethodInvocation *invocation,
                   gpointer user_data)
{
  (void) connection;
  (void) sender;
  (void) object_path;
  (void) interface_name;

  if (is_write_method(method_name)) {
    g_dbus_method_invocation_return_dbus_error(
      invocation,
      "org.xnix.Compatibility1.Error.WriteMethodDisabled",
      "Runtime write methods are blocked until production backend gates pass"
    );
    return;
  }
  (void) user_data;

  if (g_strcmp0(method_name, "ListApplications") == 0) {
    GVariantBuilder applications;

    g_variant_builder_init(&applications, G_VARIANT_TYPE("aa{sv}"));
    g_variant_builder_add_value(&applications, build_application());
    g_dbus_method_invocation_return_value(invocation, g_variant_new("(aa{sv})", &applications));
    return;
  }

  if (g_strcmp0(method_name, "GetApplication") == 0) {
    const gchar *application_id = NULL;

    g_variant_get(parameters, "(&s)", &application_id);
    if (!known_application(application_id)) {
      return_unknown_application(invocation, application_id);
      return;
    }

    g_dbus_method_invocation_return_value(invocation, g_variant_new("(@a{sv})", build_application()));
    return;
  }

  if (g_strcmp0(method_name, "GetDiagnostics") == 0) {
    const gchar *application_id = NULL;
    GVariantBuilder diagnostics;

    g_variant_get(parameters, "(&s)", &application_id);
    if (!known_application(application_id)) {
      return_unknown_application(invocation, application_id);
      return;
    }

    g_variant_builder_init(&diagnostics, G_VARIANT_TYPE("a{sv}"));
    g_variant_builder_add(&diagnostics, "{sv}", "application_id", g_variant_new_string(application_id));
    g_variant_builder_add(&diagnostics, "{sv}", "status", g_variant_new_string("known"));
    g_variant_builder_add(&diagnostics, "{sv}", "runtime_mode", g_variant_new_string("automatic"));
    g_dbus_method_invocation_return_value(invocation, g_variant_new("(a{sv})", &diagnostics));
    return;
  }

  if (g_strcmp0(method_name, "GetEngineCatalog") == 0) {
    g_dbus_method_invocation_return_value(invocation, g_variant_new("(@a{sv})", build_engine_catalog()));
    return;
  }

  if (g_strcmp0(method_name, "GetRunPlan") == 0) {
    const gchar *application_id = NULL;

    g_variant_get(parameters, "(&s)", &application_id);
    if (!known_application(application_id)) {
      return_unknown_application(invocation, application_id);
      return;
    }

    g_dbus_method_invocation_return_value(invocation, g_variant_new("(@a{sv})", build_run_plan(application_id)));
    return;
  }

  if (g_strcmp0(method_name, "GetDesktopActivationManifest") == 0) {
    const gchar *application_id = NULL;

    g_variant_get(parameters, "(&s)", &application_id);
    if (!known_application(application_id)) {
      return_unknown_application(invocation, application_id);
      return;
    }

    g_dbus_method_invocation_return_value(
      invocation,
      g_variant_new("(@a{sv})", build_desktop_activation_manifest(application_id))
    );
    return;
  }

  if (g_strcmp0(method_name, "GetDesktopEntryPlan") == 0) {
    const gchar *application_id = NULL;

    g_variant_get(parameters, "(&s)", &application_id);
    if (!known_application(application_id)) {
      return_unknown_application(invocation, application_id);
      return;
    }

    g_dbus_method_invocation_return_value(
      invocation,
      g_variant_new("(@a{sv})", build_desktop_entry_plan(application_id))
    );
    return;
  }

  if (g_strcmp0(method_name, "GetTaskManagerIdentityPlan") == 0) {
    const gchar *application_id = NULL;

    g_variant_get(parameters, "(&s)", &application_id);
    if (!known_application(application_id)) {
      return_unknown_application(invocation, application_id);
      return;
    }

    g_dbus_method_invocation_return_value(
      invocation,
      g_variant_new("(@a{sv})", build_task_manager_identity_plan(application_id))
    );
    return;
  }

  if (g_strcmp0(method_name, "GetKDEIntegrationStatus") == 0) {
    g_dbus_method_invocation_return_value(
      invocation,
      g_variant_new("(@a{sv})", build_kde_integration_status())
    );
    return;
  }

  if (g_strcmp0(method_name, "GetKWinWindowRulePlan") == 0) {
    const gchar *application_id = NULL;

    g_variant_get(parameters, "(&s)", &application_id);
    if (!known_application(application_id)) {
      return_unknown_application(invocation, application_id);
      return;
    }

    g_dbus_method_invocation_return_value(
      invocation,
      g_variant_new("(@a{sv})", build_kwin_window_rule_plan(application_id))
    );
    return;
  }

  if (g_strcmp0(method_name, "GetFileAssociationPlan") == 0) {
    const gchar *application_id = NULL;

    g_variant_get(parameters, "(&s)", &application_id);
    if (!known_application(application_id)) {
      return_unknown_application(invocation, application_id);
      return;
    }

    g_dbus_method_invocation_return_value(
      invocation,
      g_variant_new("(@a{sv})", build_file_association_plan(application_id))
    );
    return;
  }

  if (g_strcmp0(method_name, "GetNotificationPlan") == 0) {
    const gchar *application_id = NULL;
    const gchar *event_type = NULL;

    g_variant_get(parameters, "(&s&s)", &application_id, &event_type);
    if (!known_application(application_id)) {
      return_unknown_application(invocation, application_id);
      return;
    }

    g_dbus_method_invocation_return_value(
      invocation,
      g_variant_new("(@a{sv})", build_notification_plan(application_id, event_type))
    );
    return;
  }

  if (g_strcmp0(method_name, "GetTrayStatus") == 0) {
    g_dbus_method_invocation_return_value(
      invocation,
      g_variant_new("(@a{sv})", build_tray_status())
    );
    return;
  }

  if (g_strcmp0(method_name, "GetKRunnerQueryPlan") == 0) {
    const gchar *query = NULL;

    g_variant_get(parameters, "(&s)", &query);
    g_dbus_method_invocation_return_value(
      invocation,
      g_variant_new("(@a{sv})", build_krunner_query_plan(query))
    );
    return;
  }

  if (g_strcmp0(method_name, "GetPortalRequestPlan") == 0) {
    const gchar *application_id = NULL;
    const gchar *operation = NULL;

    g_variant_get(parameters, "(&s&s)", &application_id, &operation);
    if (!known_application(application_id)) {
      return_unknown_application(invocation, application_id);
      return;
    }

    g_dbus_method_invocation_return_value(
      invocation,
      g_variant_new("(@a{sv})", build_portal_request_plan(application_id, operation))
    );
    return;
  }

  if (g_strcmp0(method_name, "GetApplicationStateRoot") == 0) {
    const gchar *application_id = NULL;

    g_variant_get(parameters, "(&s)", &application_id);
    if (!known_application(application_id)) {
      return_unknown_application(invocation, application_id);
      return;
    }

    g_dbus_method_invocation_return_value(invocation, g_variant_new("(@a{sv})", build_application_state_root(application_id)));
    return;
  }

  if (g_strcmp0(method_name, "GetCompatibilityPackageSource") == 0) {
    const gchar *application_id = NULL;

    g_variant_get(parameters, "(&s)", &application_id);
    if (!known_application(application_id)) {
      return_unknown_application(invocation, application_id);
      return;
    }

    g_dbus_method_invocation_return_value(invocation, g_variant_new("(@a{sv})", build_compatibility_package_source(application_id)));
    return;
  }

  if (g_strcmp0(method_name, "GetCompatibilityAcquisitionPreflight") == 0) {
    const gchar *application_id = NULL;

    g_variant_get(parameters, "(&s)", &application_id);
    if (!known_application(application_id)) {
      return_unknown_application(invocation, application_id);
      return;
    }

    g_dbus_method_invocation_return_value(invocation, g_variant_new("(@a{sv})", build_compatibility_acquisition_preflight(application_id)));
    return;
  }

  if (g_strcmp0(method_name, "GetCompatibilityArtifactManifest") == 0) {
    const gchar *application_id = NULL;

    g_variant_get(parameters, "(&s)", &application_id);
    if (!known_application(application_id)) {
      return_unknown_application(invocation, application_id);
      return;
    }

    g_dbus_method_invocation_return_value(invocation, g_variant_new("(@a{sv})", build_compatibility_artifact_manifest(application_id)));
    return;
  }

  if (g_strcmp0(method_name, "GetCompatibilityInstallPlan") == 0) {
    const gchar *application_id = NULL;
    const gchar *environment = NULL;

    g_variant_get(parameters, "(&s&s)", &application_id, &environment);
    if (!known_application(application_id)) {
      return_unknown_application(invocation, application_id);
      return;
    }

    g_dbus_method_invocation_return_value(invocation, g_variant_new("(@a{sv})", build_compatibility_install_plan(application_id, environment)));
    return;
  }

  if (g_strcmp0(method_name, "GetBackendBinding") == 0) {
    const gchar *application_id = NULL;

    g_variant_get(parameters, "(&s)", &application_id);
    if (!known_application(application_id)) {
      return_unknown_application(invocation, application_id);
      return;
    }

    g_dbus_method_invocation_return_value(invocation, g_variant_new("(@a{sv})", build_backend_binding(application_id)));
    return;
  }

  if (g_strcmp0(method_name, "GetRepairPlan") == 0) {
    const gchar *application_id = NULL;
    const gchar *issue = NULL;

    g_variant_get(parameters, "(&s&s)", &application_id, &issue);
    if (!known_application(application_id)) {
      return_unknown_application(invocation, application_id);
      return;
    }

    g_dbus_method_invocation_return_value(invocation, g_variant_new("(@a{sv})", build_repair_plan(application_id, issue)));
    return;
  }

  if (g_strcmp0(method_name, "GetTestPlan") == 0) {
    const gchar *application_id = NULL;
    const gchar *test_type = NULL;

    g_variant_get(parameters, "(&s&s)", &application_id, &test_type);
    if (!known_application(application_id)) {
      return_unknown_application(invocation, application_id);
      return;
    }

    g_dbus_method_invocation_return_value(invocation, g_variant_new("(@a{sv})", build_test_plan(application_id, test_type)));
    return;
  }

  if (g_strcmp0(method_name, "GetTestResult") == 0) {
    const gchar *application_id = NULL;
    const gchar *test_type = NULL;

    g_variant_get(parameters, "(&s&s)", &application_id, &test_type);
    if (!known_application(application_id)) {
      return_unknown_application(invocation, application_id);
      return;
    }

    g_dbus_method_invocation_return_value(invocation, g_variant_new("(@a{sv})", build_test_result(application_id, test_type)));
    return;
  }

  if (g_strcmp0(method_name, "GetAIDiagnosticInput") == 0) {
    const gchar *application_id = NULL;
    const gchar *issue = NULL;
    const gchar *test_type = NULL;

    g_variant_get(parameters, "(&s&s&s)", &application_id, &issue, &test_type);
    if (!known_application(application_id)) {
      return_unknown_application(invocation, application_id);
      return;
    }

    g_dbus_method_invocation_return_value(
      invocation,
      g_variant_new("(@a{sv})", build_ai_diagnostic_input(application_id, issue, test_type))
    );
    return;
  }

  if (g_strcmp0(method_name, "GetAIDiagnosticRecommendation") == 0) {
    const gchar *application_id = NULL;
    const gchar *issue = NULL;
    const gchar *test_type = NULL;

    g_variant_get(parameters, "(&s&s&s)", &application_id, &issue, &test_type);
    if (!known_application(application_id)) {
      return_unknown_application(invocation, application_id);
      return;
    }

    g_dbus_method_invocation_return_value(
      invocation,
      g_variant_new("(@a{sv})", build_ai_diagnostic_recommendation(application_id, issue, test_type))
    );
    return;
  }

  if (g_strcmp0(method_name, "GetAIRepairApprovalGate") == 0) {
    const gchar *application_id = NULL;
    const gchar *issue = NULL;
    const gchar *test_type = NULL;

    g_variant_get(parameters, "(&s&s&s)", &application_id, &issue, &test_type);
    if (!known_application(application_id)) {
      return_unknown_application(invocation, application_id);
      return;
    }

    g_dbus_method_invocation_return_value(
      invocation,
      g_variant_new("(@a{sv})", build_ai_repair_approval_gate(application_id, issue, test_type))
    );
    return;
  }

  if (g_strcmp0(method_name, "GetSnapshotPlan") == 0) {
    const gchar *application_id = NULL;
    const gchar *reason = NULL;

    g_variant_get(parameters, "(&s&s)", &application_id, &reason);
    if (!known_application(application_id)) {
      return_unknown_application(invocation, application_id);
      return;
    }

    g_dbus_method_invocation_return_value(invocation, g_variant_new("(@a{sv})", build_snapshot_plan(application_id, reason)));
    return;
  }

  if (g_strcmp0(method_name, "GetPortalAccessPolicy") == 0) {
    const gchar *application_id = NULL;
    const gchar *operation = NULL;

    g_variant_get(parameters, "(&s&s)", &application_id, &operation);
    if (!known_application(application_id)) {
      return_unknown_application(invocation, application_id);
      return;
    }

    g_dbus_method_invocation_return_value(invocation, g_variant_new("(@a{sv})", build_portal_policy(application_id, operation)));
    return;
  }

  if (g_strcmp0(method_name, "GetRuntimeServiceBinding") == 0) {
    g_dbus_method_invocation_return_value(invocation, g_variant_new("(@a{sv})", build_runtime_service_binding()));
    return;
  }

  if (g_strcmp0(method_name, "GetRuntimeLiveOwnerGate") == 0) {
    g_dbus_method_invocation_return_value(invocation, g_variant_new("(@a{sv})", build_runtime_live_owner_gate()));
    return;
  }

  if (g_strcmp0(method_name, "GetRuntimeOwnerSmokePlan") == 0) {
    g_dbus_method_invocation_return_value(invocation, g_variant_new("(@a{sv})", build_runtime_owner_smoke_plan()));
    return;
  }

  if (g_strcmp0(method_name, "GetRuntimeMethodParityManifest") == 0) {
    g_dbus_method_invocation_return_value(invocation, g_variant_new("(@a{sv})", build_runtime_method_parity_manifest()));
    return;
  }

  if (g_strcmp0(method_name, "GetRuntimeWriteGate") == 0) {
    const gchar *requested_method = NULL;

    g_variant_get(parameters, "(&s)", &requested_method);
    g_dbus_method_invocation_return_value(invocation, g_variant_new("(@a{sv})", build_runtime_write_gate(requested_method)));
    return;
  }

  if (g_strcmp0(method_name, "GetCompatibilitySettings") == 0) {
    const gchar *application_id = NULL;

    g_variant_get(parameters, "(&s)", &application_id);
    if (!known_application(application_id)) {
      return_unknown_application(invocation, application_id);
      return;
    }

    g_dbus_method_invocation_return_value(invocation, g_variant_new("(@a{sv})", build_settings_model(application_id)));
    return;
  }

  if (g_strcmp0(method_name, "GetCompatibilitySettingsChangePlan") == 0) {
    const gchar *application_id = NULL;
    const gchar *section_id = NULL;
    const gchar *field_id = NULL;
    const gchar *value = NULL;

    g_variant_get(parameters, "(&s&s&s&s)", &application_id, &section_id, &field_id, &value);
    if (!known_application(application_id)) {
      return_unknown_application(invocation, application_id);
      return;
    }

    g_dbus_method_invocation_return_value(
      invocation,
      g_variant_new("(@a{sv})", build_settings_change_plan(application_id, section_id, field_id, value))
    );
    return;
  }

  if (g_strcmp0(method_name, "GetCompatibilityActionQueue") == 0) {
    const gchar *application_id = NULL;

    g_variant_get(parameters, "(&s)", &application_id);
    if (!known_application(application_id)) {
      return_unknown_application(invocation, application_id);
      return;
    }

    g_dbus_method_invocation_return_value(invocation, g_variant_new("(@a{sv})", build_action_queue(application_id)));
    return;
  }

  if (g_strcmp0(method_name, "GetCompatibilityActionReviewReceipt") == 0) {
    const gchar *application_id = NULL;
    const gchar *action_id = NULL;
    const gchar *decision = NULL;

    g_variant_get(parameters, "(&s&s&s)", &application_id, &action_id, &decision);
    if (!known_application(application_id)) {
      return_unknown_application(invocation, application_id);
      return;
    }

    g_dbus_method_invocation_return_value(
      invocation,
      g_variant_new("(@a{sv})", build_action_review_receipt(application_id, action_id, decision))
    );
    return;
  }

  if (g_strcmp0(method_name, "GetCompatibilityCenterSummary") == 0) {
    const gchar *application_id = NULL;

    g_variant_get(parameters, "(&s)", &application_id);
    if (!known_application(application_id)) {
      return_unknown_application(invocation, application_id);
      return;
    }

    g_dbus_method_invocation_return_value(
      invocation,
      g_variant_new("(@a{sv})", build_compatibility_center_summary(application_id))
    );
    return;
  }

  g_dbus_method_invocation_return_error(invocation,
                                        G_IO_ERROR,
                                        G_IO_ERROR_NOT_SUPPORTED,
                                        "unsupported runtime method: %s",
                                        method_name);
}

static const GDBusInterfaceVTable interface_vtable = {
  handle_method_call,
  NULL,
  NULL
};

static void
on_bus_acquired(GDBusConnection *connection, const gchar *name, gpointer user_data)
{
  GError *error = NULL;
  guint registration_id;

  (void) name;
  (void) user_data;

  registration_id = g_dbus_connection_register_object(connection,
                                                      "/org/xnix/Compatibility1",
                                                      introspection_data->interfaces[0],
                                                      &interface_vtable,
                                                      NULL,
                                                      NULL,
                                                      &error);
  if (registration_id == 0) {
    g_printerr("failed to register object: %s\n", error->message);
    g_clear_error(&error);
    g_main_loop_quit((GMainLoop *) user_data);
  }
}

static void
on_name_lost(GDBusConnection *connection, const gchar *name, gpointer user_data)
{
  (void) connection;
  (void) name;

  g_main_loop_quit((GMainLoop *) user_data);
}

int
main(void)
{
  GError *error = NULL;
  GMainLoop *loop;
  guint owner_id;

  introspection_data = g_dbus_node_info_new_for_xml(introspection_xml, &error);
  if (introspection_data == NULL) {
    g_printerr("failed to parse introspection XML: %s\n", error->message);
    g_clear_error(&error);
    return 1;
  }

  loop = g_main_loop_new(NULL, FALSE);
  owner_id = g_bus_own_name(G_BUS_TYPE_SESSION,
                            "org.xnix.Compatibility1",
                            G_BUS_NAME_OWNER_FLAGS_NONE,
                            on_bus_acquired,
                            NULL,
                            on_name_lost,
                            loop,
                            NULL);

  g_main_loop_run(loop);

  g_bus_unown_name(owner_id);
  g_main_loop_unref(loop);
  g_dbus_node_info_unref(introspection_data);
  return 0;
}
