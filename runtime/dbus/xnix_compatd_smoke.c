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
