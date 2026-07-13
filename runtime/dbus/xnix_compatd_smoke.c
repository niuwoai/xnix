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
  g_variant_builder_add(&recommendation, "{sv}", "safe_for_ai_diagnostics", g_variant_new_boolean(TRUE));
  g_variant_builder_add(&recommendation, "{sv}", "ai_provider_called", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&recommendation, "{sv}", "network_required", g_variant_new_boolean(FALSE));
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
