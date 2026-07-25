#include <gio/gio.h>
#include <string.h>

#include "xnix_compatd_introspection.inc"

static GDBusNodeInfo *introspection_data = NULL;

static gboolean safe_evidence_relative_path(const gchar *evidence_relative_path);

static gchar *
go_owner_read_dispatch5(const gchar *method_name,
                        const gchar *dispatch_arg,
                        const gchar *second_dispatch_arg,
                        const gchar *third_dispatch_arg,
                        const gchar *fourth_dispatch_arg,
                        const gchar *fifth_dispatch_arg)
{
  gchar *stdout_data = NULL;
  gchar *stderr_data = NULL;
  GError *error = NULL;
  gint wait_status = 0;
  gchar *argv[11] = {0};

  argv[0] = "xnix-runtime-owner";
  argv[1] = "--root";
  argv[2] = ".";
  argv[3] = "--dispatch-read";
  argv[4] = (gchar *)method_name;
  argv[5] = (gchar *)dispatch_arg;
  argv[6] = (gchar *)second_dispatch_arg;
  argv[7] = (gchar *)third_dispatch_arg;
  argv[8] = (gchar *)fourth_dispatch_arg;
  argv[9] = (gchar *)fifth_dispatch_arg;
  argv[10] = NULL;

  if (!g_spawn_sync(NULL,
                    argv,
                    NULL,
                    G_SPAWN_SEARCH_PATH,
                    NULL,
                    NULL,
                    &stdout_data,
                    &stderr_data,
                    &wait_status,
                    &error)) {
    g_clear_error(&error);
    g_free(stdout_data);
    g_free(stderr_data);
    return NULL;
  }
  if (!g_spawn_check_wait_status(wait_status, &error)) {
    g_clear_error(&error);
    g_free(stdout_data);
    g_free(stderr_data);
    return NULL;
  }
  g_free(stderr_data);
  if (stdout_data == NULL || stdout_data[0] == '\0') {
    g_free(stdout_data);
    return NULL;
  }
  g_strchomp(stdout_data);
  return stdout_data;
}

static gchar *
go_owner_service_call5(const gchar *method_name,
                       const gchar *dispatch_arg,
                       const gchar *second_dispatch_arg,
                       const gchar *third_dispatch_arg,
                       const gchar *fourth_dispatch_arg,
                       const gchar *fifth_dispatch_arg)
{
  gchar *stdout_data = NULL;
  gchar *stderr_data = NULL;
  GError *error = NULL;
  gint wait_status = 0;
  gchar *argv[13] = {0};

  argv[0] = "xnix-runtime-owner";
  argv[1] = "--root";
  argv[2] = ".";
  argv[3] = "--mode";
  argv[4] = "smoke-owner";
  argv[5] = "--service-call";
  argv[6] = (gchar *)method_name;
  argv[7] = (gchar *)dispatch_arg;
  argv[8] = (gchar *)second_dispatch_arg;
  argv[9] = (gchar *)third_dispatch_arg;
  argv[10] = (gchar *)fourth_dispatch_arg;
  argv[11] = (gchar *)fifth_dispatch_arg;
  argv[12] = NULL;

  if (!g_spawn_sync(NULL,
                    argv,
                    NULL,
                    G_SPAWN_SEARCH_PATH,
                    NULL,
                    NULL,
                    &stdout_data,
                    &stderr_data,
                    &wait_status,
                    &error)) {
    g_clear_error(&error);
    g_free(stdout_data);
    g_free(stderr_data);
    return NULL;
  }
  if (!g_spawn_check_wait_status(wait_status, &error)) {
    g_clear_error(&error);
    g_free(stdout_data);
    g_free(stderr_data);
    return NULL;
  }
  g_free(stderr_data);
  if (stdout_data == NULL || stdout_data[0] == '\0') {
    g_free(stdout_data);
    return NULL;
  }
  g_strchomp(stdout_data);
  return stdout_data;
}

static gchar *
go_service_call_materialization_preview(const gchar *evidence_relative_path)
{
  const gchar *state_root = g_getenv("XNIX_RUNTIME_OWNER_STATE_ROOT");
  gchar *stdout_data = NULL;
  gchar *stderr_data = NULL;
  GError *error = NULL;
  gint wait_status = 0;
  gchar *argv[10] = {0};

  if (state_root == NULL || state_root[0] == '\0' || !safe_evidence_relative_path(evidence_relative_path)) {
    return NULL;
  }

  argv[0] = "xnix-runtime-go";
  argv[1] = "desktop-trigger-service-call-materialization-preview";
  argv[2] = "--state-root";
  argv[3] = (gchar *)state_root;
  argv[4] = "--desktop-entry-file";
  argv[5] = "kde/actions/xnix-runtime-status-controlled-launch.desktop";
  argv[6] = "--human-authorized-smoke";
  argv[7] = "--evidence-relative-path";
  argv[8] = (gchar *)evidence_relative_path;
  argv[9] = NULL;

  if (!g_spawn_sync(NULL,
                    argv,
                    NULL,
                    G_SPAWN_SEARCH_PATH,
                    NULL,
                    NULL,
                    &stdout_data,
                    &stderr_data,
                    &wait_status,
                    &error)) {
    g_clear_error(&error);
    g_free(stdout_data);
    g_free(stderr_data);
    return NULL;
  }
  if (!g_spawn_check_wait_status(wait_status, &error)) {
    g_clear_error(&error);
    g_free(stdout_data);
    g_free(stderr_data);
    return NULL;
  }
  g_free(stderr_data);
  if (stdout_data == NULL || stdout_data[0] == '\0') {
    g_free(stdout_data);
    return NULL;
  }
  g_strchomp(stdout_data);
  return stdout_data;
}

static const gchar *
json_skip_string(const gchar *cursor)
{
  if (cursor == NULL || *cursor != '"') {
    return NULL;
  }
  cursor++;
  while (*cursor != '\0') {
    if (*cursor == '\\' && cursor[1] != '\0') {
      cursor += 2;
      continue;
    }
    if (*cursor == '"') {
      return cursor + 1;
    }
    cursor++;
  }
  return NULL;
}

static gchar *
json_read_string(const gchar *cursor)
{
  GString *value = NULL;

  if (cursor == NULL || *cursor != '"') {
    return NULL;
  }
  cursor++;
  value = g_string_new("");
  while (*cursor != '\0') {
    if (*cursor == '\\' && cursor[1] != '\0') {
      cursor++;
      g_string_append_c(value, *cursor);
      cursor++;
      continue;
    }
    if (*cursor == '"') {
      return g_string_free(value, FALSE);
    }
    g_string_append_c(value, *cursor);
    cursor++;
  }
  g_string_free(value, TRUE);
  return NULL;
}

static const gchar *
json_find_key_value(const gchar *json, const gchar *key)
{
  gchar *quoted_key = NULL;
  const gchar *cursor = NULL;

  if (json == NULL || key == NULL || key[0] == '\0') {
    return NULL;
  }

  quoted_key = g_strdup_printf("\"%s\"", key);
  cursor = strstr(json, quoted_key);
  g_free(quoted_key);
  if (cursor == NULL) {
    return NULL;
  }
  cursor = strchr(cursor, ':');
  if (cursor == NULL) {
    return NULL;
  }
  cursor++;
  while (*cursor == ' ' || *cursor == '\n' || *cursor == '\r' || *cursor == '\t') {
    cursor++;
  }
  return cursor;
}

static gchar *
json_string_value(const gchar *json, const gchar *key)
{
  return json_read_string(json_find_key_value(json, key));
}

static gchar *
json_string_array_value(const gchar *json, const gchar *key, guint index)
{
  const gchar *cursor = json_find_key_value(json, key);
  guint current_index = 0;

  if (cursor == NULL || *cursor != '[') {
    return NULL;
  }
  cursor++;
  while (*cursor != '\0') {
    while (*cursor == ' ' || *cursor == '\n' || *cursor == '\r' || *cursor == '\t' || *cursor == ',') {
      cursor++;
    }
    if (*cursor == ']') {
      return NULL;
    }
    if (*cursor != '"') {
      return NULL;
    }
    if (current_index == index) {
      return json_read_string(cursor);
    }
    cursor = json_skip_string(cursor);
    if (cursor == NULL) {
      return NULL;
    }
    current_index++;
  }
  return NULL;
}

static gchar *
go_owner_read_dispatch3(const gchar *method_name,
                        const gchar *dispatch_arg,
                        const gchar *second_dispatch_arg,
                        const gchar *third_dispatch_arg)
{
  return go_owner_read_dispatch5(method_name, dispatch_arg, second_dispatch_arg, third_dispatch_arg, NULL, NULL);
}

static gchar *
go_owner_read_dispatch(const gchar *method_name,
                       const gchar *dispatch_arg,
                       const gchar *second_dispatch_arg)
{
  return go_owner_read_dispatch3(method_name, dispatch_arg, second_dispatch_arg, NULL);
}

static gchar *
go_owner_write_gate_dispatch(const gchar *method_name)
{
  return go_owner_read_dispatch("GetRuntimeWriteGate", method_name, NULL);
}

static void
add_go_owner_dispatch_bridge_fields5(GVariantBuilder *builder,
                                     const gchar *method_name,
                                     const gchar *dispatch_arg,
                                     const gchar *second_dispatch_arg,
                                     const gchar *third_dispatch_arg,
                                     const gchar *fourth_dispatch_arg,
                                     const gchar *fifth_dispatch_arg)
{
  gchar *go_owner_dispatch_json = NULL;
  gchar *go_owner_service_call_json = NULL;
  gboolean go_owner_dispatch_available = FALSE;
  gboolean go_owner_service_call_available = FALSE;

  if (g_strcmp0(method_name, "GetRuntimeWriteGate") == 0) {
    go_owner_dispatch_json = go_owner_write_gate_dispatch(dispatch_arg);
  } else {
    go_owner_dispatch_json = go_owner_read_dispatch5(method_name,
                                                    dispatch_arg,
                                                    second_dispatch_arg,
                                                    third_dispatch_arg,
                                                    fourth_dispatch_arg,
                                                    fifth_dispatch_arg);
  }
  go_owner_service_call_json = go_owner_service_call5(method_name,
                                                      dispatch_arg,
                                                      second_dispatch_arg,
                                                      third_dispatch_arg,
                                                      fourth_dispatch_arg,
                                                      fifth_dispatch_arg);
  go_owner_dispatch_available = go_owner_dispatch_json != NULL && go_owner_dispatch_json[0] != '\0';
  go_owner_service_call_available = go_owner_service_call_json != NULL && go_owner_service_call_json[0] != '\0';

  g_variant_builder_add(builder, "{sv}", "read_model_source", g_variant_new_string("go-runtime-owner-dispatch+c-smoke-bridge"));
  g_variant_builder_add(builder, "{sv}", "go_owner_dispatch_available", g_variant_new_boolean(go_owner_dispatch_available));
  g_variant_builder_add(builder, "{sv}", "go_owner_dispatch_schema", g_variant_new_string(go_owner_dispatch_available ? "xnix.runtime.owner_read_dispatch.v1" : ""));
  g_variant_builder_add(builder, "{sv}", "go_owner_dispatch_json", g_variant_new_string(go_owner_dispatch_available ? go_owner_dispatch_json : ""));
  g_variant_builder_add(builder, "{sv}", "go_owner_service_call_available", g_variant_new_boolean(go_owner_service_call_available));
  g_variant_builder_add(builder, "{sv}", "go_owner_service_call_schema", g_variant_new_string(go_owner_service_call_available ? "xnix.runtime.owner_service_call.v1" : ""));
  g_variant_builder_add(builder, "{sv}", "go_owner_service_call_request_type", g_variant_new_string(go_owner_service_call_available ? "runtime-owner-service-call" : ""));
  g_variant_builder_add(builder, "{sv}", "go_owner_service_call_json", g_variant_new_string(go_owner_service_call_available ? go_owner_service_call_json : ""));
  g_free(go_owner_dispatch_json);
  g_free(go_owner_service_call_json);
}

static void
add_go_owner_dispatch_payload_fields5(GVariantBuilder *builder,
                                      const gchar *method_name,
                                      const gchar *dispatch_arg,
                                      const gchar *second_dispatch_arg,
                                      const gchar *third_dispatch_arg,
                                      const gchar *fourth_dispatch_arg,
                                      const gchar *fifth_dispatch_arg)
{
  gchar *go_owner_dispatch_json = NULL;
  gchar *go_owner_service_call_json = NULL;
  gboolean go_owner_dispatch_available = FALSE;
  gboolean go_owner_service_call_available = FALSE;

  if (g_strcmp0(method_name, "GetRuntimeWriteGate") == 0) {
    go_owner_dispatch_json = go_owner_write_gate_dispatch(dispatch_arg);
  } else {
    go_owner_dispatch_json = go_owner_read_dispatch5(method_name,
                                                    dispatch_arg,
                                                    second_dispatch_arg,
                                                    third_dispatch_arg,
                                                    fourth_dispatch_arg,
                                                    fifth_dispatch_arg);
  }
  go_owner_service_call_json = go_owner_service_call5(method_name,
                                                      dispatch_arg,
                                                      second_dispatch_arg,
                                                      third_dispatch_arg,
                                                      fourth_dispatch_arg,
                                                      fifth_dispatch_arg);
  go_owner_dispatch_available = go_owner_dispatch_json != NULL && go_owner_dispatch_json[0] != '\0';
  go_owner_service_call_available = go_owner_service_call_json != NULL && go_owner_service_call_json[0] != '\0';

  g_variant_builder_add(builder, "{sv}", "go_owner_dispatch_available", g_variant_new_boolean(go_owner_dispatch_available));
  g_variant_builder_add(builder, "{sv}", "go_owner_dispatch_schema", g_variant_new_string(go_owner_dispatch_available ? "xnix.runtime.owner_read_dispatch.v1" : ""));
  g_variant_builder_add(builder, "{sv}", "go_owner_dispatch_json", g_variant_new_string(go_owner_dispatch_available ? go_owner_dispatch_json : ""));
  g_variant_builder_add(builder, "{sv}", "go_owner_service_call_available", g_variant_new_boolean(go_owner_service_call_available));
  g_variant_builder_add(builder, "{sv}", "go_owner_service_call_schema", g_variant_new_string(go_owner_service_call_available ? "xnix.runtime.owner_service_call.v1" : ""));
  g_variant_builder_add(builder, "{sv}", "go_owner_service_call_request_type", g_variant_new_string(go_owner_service_call_available ? "runtime-owner-service-call" : ""));
  g_variant_builder_add(builder, "{sv}", "go_owner_service_call_json", g_variant_new_string(go_owner_service_call_available ? go_owner_service_call_json : ""));
  g_free(go_owner_dispatch_json);
  g_free(go_owner_service_call_json);
}

static void
add_go_owner_dispatch_payload_fields2(GVariantBuilder *builder,
                                      const gchar *method_name,
                                      const gchar *dispatch_arg,
                                      const gchar *second_dispatch_arg)
{
  add_go_owner_dispatch_payload_fields5(builder, method_name, dispatch_arg, second_dispatch_arg, NULL, NULL, NULL);
}

static void
add_go_owner_dispatch_payload_fields(GVariantBuilder *builder,
                                     const gchar *method_name,
                                     const gchar *dispatch_arg)
{
  add_go_owner_dispatch_payload_fields2(builder, method_name, dispatch_arg, NULL);
}

static void
add_go_owner_dispatch_payload_fields3(GVariantBuilder *builder,
                                      const gchar *method_name,
                                      const gchar *dispatch_arg,
                                      const gchar *second_dispatch_arg,
                                      const gchar *third_dispatch_arg)
{
  add_go_owner_dispatch_payload_fields5(builder, method_name, dispatch_arg, second_dispatch_arg, third_dispatch_arg, NULL, NULL);
}

static void
add_go_owner_dispatch_bridge_fields2(GVariantBuilder *builder,
                                     const gchar *method_name,
                                     const gchar *dispatch_arg,
                                     const gchar *second_dispatch_arg)
{
  add_go_owner_dispatch_bridge_fields5(builder, method_name, dispatch_arg, second_dispatch_arg, NULL, NULL, NULL);
}

static void
add_go_owner_dispatch_bridge_fields3(GVariantBuilder *builder,
                                     const gchar *method_name,
                                     const gchar *dispatch_arg,
                                     const gchar *second_dispatch_arg,
                                     const gchar *third_dispatch_arg)
{
  add_go_owner_dispatch_bridge_fields5(builder, method_name, dispatch_arg, second_dispatch_arg, third_dispatch_arg, NULL, NULL);
}

static void
add_go_owner_dispatch_bridge_fields(GVariantBuilder *builder,
                                    const gchar *method_name,
                                    const gchar *dispatch_arg)
{
  add_go_owner_dispatch_bridge_fields2(builder, method_name, dispatch_arg, NULL);
}

static GVariant *
build_application(const gchar *method_name)
{
  GVariantBuilder builder;

  g_variant_builder_init(&builder, G_VARIANT_TYPE("a{sv}"));
  add_go_owner_dispatch_payload_fields(&builder,
                                       method_name,
                                       g_strcmp0(method_name, "ListApplications") == 0 ? NULL : "org.xnix.sample.notepad");
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
  add_go_owner_dispatch_bridge_fields(&catalog, "GetEngineCatalog", NULL);
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
  add_go_owner_dispatch_bridge_fields(&plan, "GetRunPlan", application_id);
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
  add_go_owner_dispatch_bridge_fields(&manifest, "GetDesktopActivationManifest", application_id);
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
build_desktop_activation_transaction_preview(const gchar *application_id, const gchar *mode)
{
  GVariantBuilder transaction;
  gboolean transaction_ready = g_strcmp0(mode, "development") == 0;

  g_variant_builder_init(&transaction, G_VARIANT_TYPE("a{sv}"));
  g_variant_builder_add(&transaction, "{sv}", "schema_version", g_variant_new_string("xnix.runtime.desktop_activation_transaction.v1"));
  g_variant_builder_add(&transaction, "{sv}", "request_type", g_variant_new_string("desktop-activation-transaction-preview"));
  add_go_owner_dispatch_bridge_fields2(&transaction, "GetDesktopActivationTransactionPreview", application_id, mode);
  g_variant_builder_add(&transaction, "{sv}", "transaction_type", g_variant_new_string("kde-desktop-activation-transaction"));
  g_variant_builder_add(&transaction, "{sv}", "transaction_state", g_variant_new_string(transaction_ready ? "transaction-ready" : "transaction-blocked"));
  g_variant_builder_add(&transaction, "{sv}", "desktop", g_variant_new_string("KDE Plasma"));
  g_variant_builder_add(&transaction, "{sv}", "runtime_method", g_variant_new_string("GetDesktopActivationTransactionPreview"));
  g_variant_builder_add(&transaction, "{sv}", "read_model_source", g_variant_new_string("xnix-runtime-go desktop-activation-transaction-preview"));
  g_variant_builder_add(&transaction, "{sv}", "write_method", g_variant_new_string("ActivateDesktopIntegration"));
  g_variant_builder_add(&transaction, "{sv}", "application_id", g_variant_new_string(application_id));
  g_variant_builder_add(&transaction, "{sv}", "install_mode", g_variant_new_string(mode));
  g_variant_builder_add(&transaction, "{sv}", "planned_file_count", g_variant_new_int32(5));
  g_variant_builder_add(&transaction, "{sv}", "transaction_step_count", g_variant_new_int32(9));
  g_variant_builder_add(&transaction, "{sv}", "rollback_step_count", g_variant_new_int32(7));
  g_variant_builder_add(&transaction, "{sv}", "staged_file_digests_required", g_variant_new_boolean(TRUE));
  g_variant_builder_add(&transaction, "{sv}", "staged_file_digests_verified", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&transaction, "{sv}", "runtime_owned", g_variant_new_boolean(TRUE));
  g_variant_builder_add(&transaction, "{sv}", "go_runtime_backed", g_variant_new_boolean(TRUE));
  g_variant_builder_add(&transaction, "{sv}", "kde_policy_owner", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&transaction, "{sv}", "transaction_plan_created", g_variant_new_boolean(TRUE));
  g_variant_builder_add(&transaction, "{sv}", "transaction_ready", g_variant_new_boolean(transaction_ready));
  g_variant_builder_add(&transaction, "{sv}", "transaction_committed", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&transaction, "{sv}", "write_method_enabled", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&transaction, "{sv}", "dispatch_enabled", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&transaction, "{sv}", "file_writes_performed", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&transaction, "{sv}", "desktop_files_written", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&transaction, "{sv}", "mimeapps_written", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&transaction, "{sv}", "receipt_written", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&transaction, "{sv}", "rollback_receipt_required", g_variant_new_boolean(TRUE));
  g_variant_builder_add(&transaction, "{sv}", "rollback_receipt_planned", g_variant_new_boolean(TRUE));
  g_variant_builder_add(&transaction, "{sv}", "rollback_available", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&transaction, "{sv}", "kde_service_cache_refreshed", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&transaction, "{sv}", "launch_enabled", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&transaction, "{sv}", "backend_launch_enabled", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&transaction, "{sv}", "execution_started", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&transaction, "{sv}", "host_root_modified", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&transaction, "{sv}", "network_required", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&transaction, "{sv}", "privileged_container_required", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&transaction, "{sv}", "backend_details_exposed", g_variant_new_boolean(FALSE));

  return g_variant_builder_end(&transaction);
}

static GVariant *
build_desktop_activation_status(const gchar *application_id, const gchar *mode)
{
  GVariantBuilder status;
  gboolean activation_ready = g_strcmp0(mode, "development") == 0;

  g_variant_builder_init(&status, G_VARIANT_TYPE("a{sv}"));
  g_variant_builder_add(&status, "{sv}", "schema_version", g_variant_new_string("xnix.runtime.desktop_activation_status.v1"));
  g_variant_builder_add(&status, "{sv}", "request_type", g_variant_new_string("desktop-activation-status-preview"));
  add_go_owner_dispatch_bridge_fields2(&status, "GetDesktopActivationStatus", application_id, mode);
  g_variant_builder_add(&status, "{sv}", "status_type", g_variant_new_string("kde-desktop-activation-status"));
  g_variant_builder_add(&status, "{sv}", "activation_state", g_variant_new_string(activation_ready ? "ready-for-runtime-commit" : "blocked-before-runtime-commit"));
  g_variant_builder_add(&status, "{sv}", "source", g_variant_new_string("desktop-activation-transaction-preview"));
  g_variant_builder_add(&status, "{sv}", "desktop", g_variant_new_string("KDE Plasma"));
  g_variant_builder_add(&status, "{sv}", "runtime_method", g_variant_new_string("GetDesktopActivationStatus"));
  g_variant_builder_add(&status, "{sv}", "read_model_source", g_variant_new_string("xnix-runtime-go desktop-activation-status-preview"));
  g_variant_builder_add(&status, "{sv}", "write_method", g_variant_new_string("ActivateDesktopIntegration"));
  g_variant_builder_add(&status, "{sv}", "application_id", g_variant_new_string(application_id));
  g_variant_builder_add(&status, "{sv}", "install_mode", g_variant_new_string(mode));
  g_variant_builder_add(&status, "{sv}", "preflight_decision", g_variant_new_string(activation_ready ? "development-staging-ready" : "activation-blocked"));
  g_variant_builder_add(&status, "{sv}", "transaction_state", g_variant_new_string(activation_ready ? "transaction-ready" : "transaction-blocked"));
  g_variant_builder_add(&status, "{sv}", "planned_file_count", g_variant_new_int32(5));
  g_variant_builder_add(&status, "{sv}", "transaction_step_count", g_variant_new_int32(9));
  g_variant_builder_add(&status, "{sv}", "rollback_step_count", g_variant_new_int32(7));
  g_variant_builder_add(&status, "{sv}", "status_signal_count", g_variant_new_int32(5));
  g_variant_builder_add(&status, "{sv}", "blocked_reason_count", g_variant_new_int32(activation_ready ? 4 : 5));
  g_variant_builder_add(&status, "{sv}", "next_safe_action_count", g_variant_new_int32(5));
  g_variant_builder_add(&status, "{sv}", "runtime_owned", g_variant_new_boolean(TRUE));
  g_variant_builder_add(&status, "{sv}", "go_runtime_backed", g_variant_new_boolean(TRUE));
  g_variant_builder_add(&status, "{sv}", "kde_policy_owner", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&status, "{sv}", "user_visible", g_variant_new_boolean(TRUE));
  g_variant_builder_add(&status, "{sv}", "activation_ready", g_variant_new_boolean(activation_ready));
  g_variant_builder_add(&status, "{sv}", "activation_committed", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&status, "{sv}", "commit_enabled", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&status, "{sv}", "installer_may_proceed", g_variant_new_boolean(activation_ready));
  g_variant_builder_add(&status, "{sv}", "staging_plan_ready", g_variant_new_boolean(activation_ready));
  g_variant_builder_add(&status, "{sv}", "rollback_planned", g_variant_new_boolean(TRUE));
  g_variant_builder_add(&status, "{sv}", "rollback_available", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&status, "{sv}", "host_root_modified", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&status, "{sv}", "file_writes_performed", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&status, "{sv}", "desktop_files_written", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&status, "{sv}", "mimeapps_written", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&status, "{sv}", "kde_service_cache_refreshed", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&status, "{sv}", "launch_enabled", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&status, "{sv}", "backend_launch_enabled", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&status, "{sv}", "execution_started", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&status, "{sv}", "network_required", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&status, "{sv}", "privileged_container_required", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&status, "{sv}", "backend_details_exposed", g_variant_new_boolean(FALSE));

  return g_variant_builder_end(&status);
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
  add_go_owner_dispatch_bridge_fields(&status, "GetKDEIntegrationStatus", NULL);
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
build_kde_shell_integration_plan(void)
{
  static const gchar *component_ids[] = {
    "start-menu",
    "task-manager",
    "file-manager",
    "system-tray",
    "notification-center",
    "compatibility-center",
    "unified-settings",
    "krunner-search",
    "kwin-window-management"
  };
  GVariantBuilder plan;

  g_variant_builder_init(&plan, G_VARIANT_TYPE("a{sv}"));
  g_variant_builder_add(&plan, "{sv}", "plan_type", g_variant_new_string("kde-shell-integration-plan"));
  g_variant_builder_add(&plan, "{sv}", "runtime_method", g_variant_new_string("GetKDEShellIntegrationPlan"));
  add_go_owner_dispatch_bridge_fields(&plan, "GetKDEShellIntegrationPlan", NULL);
  g_variant_builder_add(&plan, "{sv}", "desktop_shell", g_variant_new_string("KDE Plasma"));
  g_variant_builder_add(&plan, "{sv}", "component_ids", g_variant_new_strv(component_ids, 9));
  g_variant_builder_add(&plan, "{sv}", "component_count", g_variant_new_int32(9));
  g_variant_builder_add(&plan, "{sv}", "initial_component_count", g_variant_new_int32(8));
  g_variant_builder_add(&plan, "{sv}", "planned_component_count", g_variant_new_int32(1));
  g_variant_builder_add(&plan, "{sv}", "official_desktop_only", g_variant_new_boolean(TRUE));
  g_variant_builder_add(&plan, "{sv}", "fallback_desktops_supported", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&plan, "{sv}", "runtime_owned", g_variant_new_boolean(TRUE));
  g_variant_builder_add(&plan, "{sv}", "c_runtime_backed", g_variant_new_boolean(TRUE));
  g_variant_builder_add(&plan, "{sv}", "kde_policy_owner", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&plan, "{sv}", "plasma_fork_required", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&plan, "{sv}", "plasma_source_modified", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&plan, "{sv}", "shell_configuration_written", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&plan, "{sv}", "component_activation_enabled", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&plan, "{sv}", "backend_launch_enabled", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&plan, "{sv}", "backend_details_exposed", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&plan, "{sv}", "host_root_modified", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&plan, "{sv}", "privileged_container_required", g_variant_new_boolean(FALSE));

  return g_variant_builder_end(&plan);
}

static GVariant *
build_kde_application_surface_plan(const gchar *application_id)
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
  static const gchar *runtime_methods[] = {
    "GetDesktopEntryPlan",
    "GetTaskManagerIdentityPlan",
    "GetFileAssociationPlan",
    "GetTrayStatus",
    "GetNotificationPlan",
    "GetCompatibilityCenterSummary",
    "GetCompatibilitySettings"
  };
  static const gchar *required_runtime_gates[] = {
    "recipe-install-gate",
    "portal-policy-review",
    "snapshot-baseline",
    "backend-environment-plan",
    "backend-lifecycle-plan",
    "runtime-write-gate"
  };
  GVariantBuilder plan;

  g_variant_builder_init(&plan, G_VARIANT_TYPE("a{sv}"));
  g_variant_builder_add(&plan, "{sv}", "plan_type", g_variant_new_string("kde-application-surface-plan"));
  g_variant_builder_add(&plan, "{sv}", "runtime_method", g_variant_new_string("GetKDEApplicationSurfacePlan"));
  add_go_owner_dispatch_bridge_fields(&plan, "GetKDEApplicationSurfacePlan", application_id);
  g_variant_builder_add(&plan, "{sv}", "application_id", g_variant_new_string(application_id));
  g_variant_builder_add(&plan, "{sv}", "surface_state", g_variant_new_string("planned"));
  g_variant_builder_add(&plan, "{sv}", "desktop_shell", g_variant_new_string("KDE Plasma"));
  g_variant_builder_add(&plan, "{sv}", "entry_point_count", g_variant_new_int32(7));
  g_variant_builder_add(&plan, "{sv}", "entry_point_ids", g_variant_new_strv(entry_point_ids, 7));
  g_variant_builder_add(&plan, "{sv}", "runtime_methods", g_variant_new_strv(runtime_methods, 7));
  g_variant_builder_add(&plan, "{sv}", "required_runtime_gates", g_variant_new_strv(required_runtime_gates, 6));
  g_variant_builder_add(&plan, "{sv}", "runtime_owned", g_variant_new_boolean(TRUE));
  g_variant_builder_add(&plan, "{sv}", "c_runtime_backed", g_variant_new_boolean(TRUE));
  g_variant_builder_add(&plan, "{sv}", "kde_policy_owner", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&plan, "{sv}", "official_desktop_only", g_variant_new_boolean(TRUE));
  g_variant_builder_add(&plan, "{sv}", "normal_linux_application_surface", g_variant_new_boolean(TRUE));
  g_variant_builder_add(&plan, "{sv}", "standard_launcher_visible", g_variant_new_boolean(TRUE));
  g_variant_builder_add(&plan, "{sv}", "launch_enabled", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&plan, "{sv}", "backend_process_started", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&plan, "{sv}", "desktop_files_written", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&plan, "{sv}", "mimeapps_written", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&plan, "{sv}", "host_root_modified", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&plan, "{sv}", "backend_command_exposed", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&plan, "{sv}", "raw_windows_executable_exposed", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&plan, "{sv}", "backend_details_exposed", g_variant_new_boolean(FALSE));

  return g_variant_builder_end(&plan);
}

static GVariant *
build_desktop_resource_bridge_plan(const gchar *application_id)
{
  static const gchar *resource_ids[] = {
    "file-open",
    "uri-open",
    "print",
    "clipboard",
    "screenshot"
  };
  static const gchar *portal_interfaces[] = {
    "org.freedesktop.portal.FileChooser",
    "org.freedesktop.portal.OpenURI",
    "org.freedesktop.portal.Print",
    "org.freedesktop.portal.Clipboard",
    "org.freedesktop.portal.Screenshot"
  };
  static const gchar *required_runtime_gates[] = {
    "portal-policy-review",
    "portal-request-plan",
    "snapshot-baseline",
    "backend-environment-plan",
    "runtime-write-gate"
  };
  GVariantBuilder plan;

  g_variant_builder_init(&plan, G_VARIANT_TYPE("a{sv}"));
  g_variant_builder_add(&plan, "{sv}", "plan_type", g_variant_new_string("desktop-resource-bridge-plan"));
  add_go_owner_dispatch_bridge_fields(&plan, "GetDesktopResourceBridgePlan", application_id);
  g_variant_builder_add(&plan, "{sv}", "runtime_method", g_variant_new_string("GetDesktopResourceBridgePlan"));
  g_variant_builder_add(&plan, "{sv}", "application_id", g_variant_new_string(application_id));
  g_variant_builder_add(&plan, "{sv}", "bridge_state", g_variant_new_string("planned"));
  g_variant_builder_add(&plan, "{sv}", "resource_count", g_variant_new_int32(5));
  g_variant_builder_add(&plan, "{sv}", "resource_ids", g_variant_new_strv(resource_ids, 5));
  g_variant_builder_add(&plan, "{sv}", "portal_interfaces", g_variant_new_strv(portal_interfaces, 5));
  g_variant_builder_add(&plan, "{sv}", "required_runtime_gates", g_variant_new_strv(required_runtime_gates, 5));
  g_variant_builder_add(&plan, "{sv}", "runtime_owned", g_variant_new_boolean(TRUE));
  g_variant_builder_add(&plan, "{sv}", "c_runtime_backed", g_variant_new_boolean(TRUE));
  g_variant_builder_add(&plan, "{sv}", "kde_policy_owner", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&plan, "{sv}", "portal_mediated", g_variant_new_boolean(TRUE));
  g_variant_builder_add(&plan, "{sv}", "file_bridge_planned", g_variant_new_boolean(TRUE));
  g_variant_builder_add(&plan, "{sv}", "print_bridge_planned", g_variant_new_boolean(TRUE));
  g_variant_builder_add(&plan, "{sv}", "clipboard_bridge_planned", g_variant_new_boolean(TRUE));
  g_variant_builder_add(&plan, "{sv}", "bridges_enabled", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&plan, "{sv}", "requests_created", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&plan, "{sv}", "backend_process_started", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&plan, "{sv}", "direct_host_file_access", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&plan, "{sv}", "direct_clipboard_access", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&plan, "{sv}", "direct_print_access", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&plan, "{sv}", "host_root_modified", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&plan, "{sv}", "backend_details_exposed", g_variant_new_boolean(FALSE));

  return g_variant_builder_end(&plan);
}

static GVariant *
build_desktop_entry_plan(const gchar *application_id)
{
  GVariantBuilder plan;

  g_variant_builder_init(&plan, G_VARIANT_TYPE("a{sv}"));
  g_variant_builder_add(&plan, "{sv}", "plan_type", g_variant_new_string("desktop-entry-plan"));
  add_go_owner_dispatch_bridge_fields(&plan, "GetDesktopEntryPlan", application_id);
  g_variant_builder_add(&plan, "{sv}", "application_id", g_variant_new_string(application_id));
  g_variant_builder_add(&plan, "{sv}", "desktop", g_variant_new_string("KDE Plasma"));
  g_variant_builder_add(&plan, "{sv}", "desktop_file", g_variant_new_string("xnix-org.xnix.sample.notepad.desktop"));
  g_variant_builder_add(&plan, "{sv}", "name", g_variant_new_string("Sample Notepad"));
  g_variant_builder_add(&plan, "{sv}", "exec", g_variant_new_string("xnix-compat-launch --app org.xnix.sample.notepad --registry /usr/share/xnix/compatibility/recipes/registry.json %U"));
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
build_desktop_icon_plan(const gchar *application_id)
{
  GVariantBuilder plan;

  g_variant_builder_init(&plan, G_VARIANT_TYPE("a{sv}"));
  g_variant_builder_add(&plan, "{sv}", "request_type", g_variant_new_string("desktop-icon-plan"));
  g_variant_builder_add(&plan, "{sv}", "plan_type", g_variant_new_string("desktop-icon-plan"));
  add_go_owner_dispatch_bridge_fields(&plan, "GetDesktopIconPlan", application_id);
  g_variant_builder_add(&plan, "{sv}", "application_id", g_variant_new_string(application_id));
  g_variant_builder_add(&plan, "{sv}", "desktop", g_variant_new_string("KDE Plasma"));
  g_variant_builder_add(&plan, "{sv}", "runtime_method", g_variant_new_string("GetDesktopIconPlan"));
  g_variant_builder_add(&plan, "{sv}", "read_model_source", g_variant_new_string("go-desktop-icon-preview"));
  g_variant_builder_add(&plan, "{sv}", "desktop_file", g_variant_new_string("xnix-org.xnix.sample.notepad.desktop"));
  g_variant_builder_add(&plan, "{sv}", "launcher_url", g_variant_new_string("applications:xnix-org.xnix.sample.notepad.desktop"));
  g_variant_builder_add(&plan, "{sv}", "target_directory", g_variant_new_string("xdg-desktop-dir"));
  g_variant_builder_add(&plan, "{sv}", "placement", g_variant_new_string("user-desktop"));
  g_variant_builder_add(&plan, "{sv}", "exec", g_variant_new_string("xnix-compat-launch --app org.xnix.sample.notepad --registry /usr/share/xnix/compatibility/recipes/registry.json %U"));
  g_variant_builder_add(&plan, "{sv}", "icon", g_variant_new_string("accessories-text-editor"));
  g_variant_builder_add(&plan, "{sv}", "standard_desktop_entry", g_variant_new_boolean(TRUE));
  g_variant_builder_add(&plan, "{sv}", "desktop_icon_visible", g_variant_new_boolean(TRUE));
  g_variant_builder_add(&plan, "{sv}", "runtime_owned", g_variant_new_boolean(TRUE));
  g_variant_builder_add(&plan, "{sv}", "go_runtime_backed", g_variant_new_boolean(TRUE));
  g_variant_builder_add(&plan, "{sv}", "kde_policy_owner", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&plan, "{sv}", "desktop_file_copy_enabled", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&plan, "{sv}", "desktop_file_write_enabled", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&plan, "{sv}", "icon_placement_persisted", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&plan, "{sv}", "launch_enabled", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&plan, "{sv}", "backend_launch_enabled", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&plan, "{sv}", "host_root_modified", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&plan, "{sv}", "raw_executable_exposed", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&plan, "{sv}", "backend_details_exposed", g_variant_new_boolean(FALSE));

  return g_variant_builder_end(&plan);
}

static GVariant *
build_task_manager_identity_plan(const gchar *application_id)
{
  GVariantBuilder plan;

  g_variant_builder_init(&plan, G_VARIANT_TYPE("a{sv}"));
  g_variant_builder_add(&plan, "{sv}", "plan_type", g_variant_new_string("task-manager-identity-plan"));
  add_go_owner_dispatch_bridge_fields(&plan, "GetTaskManagerIdentityPlan", application_id);
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
  add_go_owner_dispatch_bridge_fields(&plan, "GetKWinWindowRulePlan", application_id);
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
  add_go_owner_dispatch_bridge_fields(&plan, "GetFileAssociationPlan", application_id);
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
  add_go_owner_dispatch_bridge_fields2(&plan, "GetNotificationPlan", application_id, event_type);
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
  add_go_owner_dispatch_bridge_fields(&status, "GetTrayStatus", NULL);
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
  add_go_owner_dispatch_bridge_fields(&plan, "GetKRunnerQueryPlan", query == NULL ? "" : query);
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

static gboolean
safe_evidence_relative_path(const gchar *evidence_relative_path)
{
  return evidence_relative_path != NULL &&
         evidence_relative_path[0] != '\0' &&
         !g_str_has_prefix(evidence_relative_path, "/") &&
         strstr(evidence_relative_path, "..") == NULL &&
         strchr(evidence_relative_path, '\n') == NULL &&
         strchr(evidence_relative_path, '\r') == NULL;
}

static GVariant *
build_runtime_controlled_launch_action(const gchar *evidence_relative_path)
{
  GVariantBuilder action;
  gchar *go_service_call_materialization_json = NULL;
  gchar *go_owner_service_call_json = NULL;
  gchar *trigger_runtime_method = NULL;
  gchar *trigger_action_type = NULL;
  gchar *trigger_call_type = NULL;
  gchar *trigger_handoff_kind = NULL;
  gchar *trigger_handoff_value = NULL;
  gchar *trigger_owner_service_method = NULL;
  gboolean go_service_call_materialization_available = FALSE;
  gboolean go_owner_service_call_available = FALSE;

  go_service_call_materialization_json = go_service_call_materialization_preview(evidence_relative_path);
  go_service_call_materialization_available = go_service_call_materialization_json != NULL && go_service_call_materialization_json[0] != '\0';
  if (go_service_call_materialization_available) {
    trigger_runtime_method = json_string_value(go_service_call_materialization_json, "desktop_callable_runtime_method");
    trigger_action_type = json_string_value(go_service_call_materialization_json, "desktop_callable_route");
    trigger_call_type = json_string_value(go_service_call_materialization_json, "owner_service_call_type");
    trigger_owner_service_method = json_string_array_value(go_service_call_materialization_json, "owner_service_call_args", 0);
    trigger_handoff_kind = json_string_array_value(go_service_call_materialization_json, "owner_service_call_args", 1);
    trigger_handoff_value = json_string_array_value(go_service_call_materialization_json, "owner_service_call_args", 2);
  }
  if (trigger_owner_service_method != NULL && trigger_owner_service_method[0] != '\0' &&
      trigger_handoff_kind != NULL && trigger_handoff_kind[0] != '\0' &&
      trigger_handoff_value != NULL && safe_evidence_relative_path(trigger_handoff_value)) {
    go_owner_service_call_json = go_owner_service_call5(trigger_owner_service_method,
                                                        trigger_handoff_kind,
                                                        trigger_handoff_value,
                                                        NULL,
                                                        NULL,
                                                        NULL);
  }
  go_owner_service_call_available = go_owner_service_call_json != NULL && go_owner_service_call_json[0] != '\0';

  g_variant_builder_init(&action, G_VARIANT_TYPE("a{sv}"));
  g_variant_builder_add(&action, "{sv}", "schema_version", g_variant_new_string("xnix.runtime.dbus_runtime_controlled_launch_action.v1"));
  g_variant_builder_add(&action, "{sv}", "request_type", g_variant_new_string("runtime-controlled-launch-dbus-action"));
  g_variant_builder_add(&action, "{sv}", "runtime_method", g_variant_new_string(trigger_runtime_method == NULL ? "" : trigger_runtime_method));
  g_variant_builder_add(&action, "{sv}", "owner_runtime_method", g_variant_new_string(trigger_owner_service_method == NULL ? "" : trigger_owner_service_method));
  g_variant_builder_add(&action, "{sv}", "desktop_callable_runtime_method", g_variant_new_string(trigger_runtime_method == NULL ? "" : trigger_runtime_method));
  g_variant_builder_add(&action, "{sv}", "action_type", g_variant_new_string(trigger_action_type == NULL ? "" : trigger_action_type));
  g_variant_builder_add(&action, "{sv}", "action_trigger_handoff_type", g_variant_new_string(trigger_handoff_kind == NULL ? "" : trigger_handoff_kind));
  g_variant_builder_add(&action, "{sv}", "evidence_relative_path", g_variant_new_string(evidence_relative_path));
  g_variant_builder_add(&action, "{sv}", "call_type", g_variant_new_string(trigger_call_type == NULL ? "" : trigger_call_type));
  g_variant_builder_add(&action, "{sv}", "go_service_call_materialization_available", g_variant_new_boolean(go_service_call_materialization_available));
  g_variant_builder_add(&action, "{sv}", "go_service_call_materialization_schema", g_variant_new_string(go_service_call_materialization_available ? "xnix.runtime.desktop_trigger_service_call_materialization.v1" : ""));
  g_variant_builder_add(&action, "{sv}", "go_service_call_materialization_request_type", g_variant_new_string(go_service_call_materialization_available ? "desktop-trigger-service-call-materialization-preview" : ""));
  g_variant_builder_add(&action, "{sv}", "go_service_call_materialization_json", g_variant_new_string(go_service_call_materialization_available ? go_service_call_materialization_json : ""));
  g_variant_builder_add(&action, "{sv}", "go_owner_service_call_available", g_variant_new_boolean(go_owner_service_call_available));
  g_variant_builder_add(&action, "{sv}", "go_owner_service_call_schema", g_variant_new_string(go_owner_service_call_available ? "xnix.runtime.owner_service_call.v1" : ""));
  g_variant_builder_add(&action, "{sv}", "go_owner_service_call_request_type", g_variant_new_string(go_owner_service_call_available ? "runtime-owner-service-call" : ""));
  g_variant_builder_add(&action, "{sv}", "go_owner_service_call_json", g_variant_new_string(go_owner_service_call_available ? go_owner_service_call_json : ""));
  g_variant_builder_add(&action, "{sv}", "runtime_owner_service_action_dispatch", g_variant_new_boolean(go_owner_service_call_available));
  g_variant_builder_add(&action, "{sv}", "kde_forwards_only_evidence_handle", g_variant_new_boolean(TRUE));
  g_variant_builder_add(&action, "{sv}", "desktop_evidence_handle_forwarded", g_variant_new_boolean(TRUE));
  g_variant_builder_add(&action, "{sv}", "desktop_receipt_fields_reconstructed", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&action, "{sv}", "desktop_kde_state_root_access", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&action, "{sv}", "runtime_owned", g_variant_new_boolean(TRUE));
  g_variant_builder_add(&action, "{sv}", "go_runtime_backed", g_variant_new_boolean(TRUE));
  g_variant_builder_add(&action, "{sv}", "kde_policy_owner", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&action, "{sv}", "production_bus_claimed", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&action, "{sv}", "write_method", g_variant_new_boolean(TRUE));
  g_variant_builder_add(&action, "{sv}", "write_methods_enabled", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&action, "{sv}", "read_only_dispatch", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&action, "{sv}", "host_root_modified", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&action, "{sv}", "network_required", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&action, "{sv}", "privileged_container_required", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&action, "{sv}", "backend_details_exposed", g_variant_new_boolean(FALSE));
  g_variant_builder_add(&action, "{sv}", "desktop_safe_summary", g_variant_new_string("D-Bus forwards a Runtime-status evidence handoff to the Go Runtime Owner controlled launch boundary."));

  g_free(go_service_call_materialization_json);
  g_free(go_owner_service_call_json);
  g_free(trigger_runtime_method);
  g_free(trigger_action_type);
  g_free(trigger_call_type);
  g_free(trigger_handoff_kind);
  g_free(trigger_handoff_value);
  g_free(trigger_owner_service_method);
  return g_variant_builder_end(&action);
}

#include "xnix_compatd_kde_center.inc"
#include "xnix_compatd_runtime_models.inc"

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
    g_variant_builder_add_value(&applications, build_application("ListApplications"));
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

    g_dbus_method_invocation_return_value(invocation, g_variant_new("(@a{sv})", build_application("GetApplication")));
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
    add_go_owner_dispatch_bridge_fields(&diagnostics, "GetDiagnostics", application_id);
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

  if (g_strcmp0(method_name, "GetDesktopActivationTransactionPreview") == 0) {
    const gchar *application_id = NULL;
    const gchar *mode = NULL;

    g_variant_get(parameters, "(&s&s)", &application_id, &mode);
    if (!known_application(application_id)) {
      return_unknown_application(invocation, application_id);
      return;
    }

    g_dbus_method_invocation_return_value(
      invocation,
      g_variant_new("(@a{sv})", build_desktop_activation_transaction_preview(application_id, mode))
    );
    return;
  }

  if (g_strcmp0(method_name, "GetDesktopActivationStatus") == 0) {
    const gchar *application_id = NULL;
    const gchar *mode = NULL;

    g_variant_get(parameters, "(&s&s)", &application_id, &mode);
    if (!known_application(application_id)) {
      return_unknown_application(invocation, application_id);
      return;
    }

    g_dbus_method_invocation_return_value(
      invocation,
      g_variant_new("(@a{sv})", build_desktop_activation_status(application_id, mode))
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

  if (g_strcmp0(method_name, "GetDesktopIconPlan") == 0) {
    const gchar *application_id = NULL;

    g_variant_get(parameters, "(&s)", &application_id);
    if (!known_application(application_id)) {
      return_unknown_application(invocation, application_id);
      return;
    }

    g_dbus_method_invocation_return_value(
      invocation,
      g_variant_new("(@a{sv})", build_desktop_icon_plan(application_id))
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

  if (g_strcmp0(method_name, "GetKDEShellIntegrationPlan") == 0) {
    g_dbus_method_invocation_return_value(
      invocation,
      g_variant_new("(@a{sv})", build_kde_shell_integration_plan())
    );
    return;
  }

  if (g_strcmp0(method_name, "GetKDEApplicationSurfacePlan") == 0) {
    const gchar *application_id = NULL;

    g_variant_get(parameters, "(&s)", &application_id);
    if (!known_application(application_id)) {
      return_unknown_application(invocation, application_id);
      return;
    }

    g_dbus_method_invocation_return_value(
      invocation,
      g_variant_new("(@a{sv})", build_kde_application_surface_plan(application_id))
    );
    return;
  }

  if (g_strcmp0(method_name, "GetDesktopResourceBridgePlan") == 0) {
    const gchar *application_id = NULL;

    g_variant_get(parameters, "(&s)", &application_id);
    if (!known_application(application_id)) {
      return_unknown_application(invocation, application_id);
      return;
    }

    g_dbus_method_invocation_return_value(
      invocation,
      g_variant_new("(@a{sv})", build_desktop_resource_bridge_plan(application_id))
    );
    return;
  }

  if (g_strcmp0(method_name, "GetCompatibilityModeSwitchPlan") == 0) {
    const gchar *application_id = NULL;
    const gchar *requested_mode = NULL;

    g_variant_get(parameters, "(&s&s)", &application_id, &requested_mode);
    if (!known_application(application_id)) {
      return_unknown_application(invocation, application_id);
      return;
    }
    if (!supported_compatibility_mode(requested_mode)) {
      g_dbus_method_invocation_return_error(
        invocation,
        G_IO_ERROR,
        G_IO_ERROR_INVALID_ARGUMENT,
        "Unsupported compatibility mode: %s",
        requested_mode == NULL ? "" : requested_mode
      );
      return;
    }

    g_dbus_method_invocation_return_value(
      invocation,
      g_variant_new("(@a{sv})", build_compatibility_mode_switch_plan(application_id, requested_mode))
    );
    return;
  }

  if (g_strcmp0(method_name, "GetCompatibilityPermissionReviewPlan") == 0) {
    const gchar *application_id = NULL;

    g_variant_get(parameters, "(&s)", &application_id);
    if (!known_application(application_id)) {
      return_unknown_application(invocation, application_id);
      return;
    }

    g_dbus_method_invocation_return_value(
      invocation,
      g_variant_new("(@a{sv})", build_compatibility_permission_review_plan(application_id))
    );
    return;
  }

  if (g_strcmp0(method_name, "GetCompatibilityReviewFlowPlan") == 0) {
    const gchar *application_id = NULL;
    const gchar *section_id = NULL;
    const gchar *field_id = NULL;
    const gchar *value = NULL;
    const gchar *operation = NULL;

    g_variant_get(parameters, "(&s&s&s&s&s)", &application_id, &section_id, &field_id, &value, &operation);
    if (!known_application(application_id)) {
      return_unknown_application(invocation, application_id);
      return;
    }

    g_dbus_method_invocation_return_value(
      invocation,
      g_variant_new("(@a{sv})", build_compatibility_review_flow_plan(application_id, section_id, field_id, value, operation))
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

  if (g_strcmp0(method_name, "GetBackendCapabilityMatrix") == 0) {
    g_dbus_method_invocation_return_value(invocation, g_variant_new("(@a{sv})", build_backend_capability_matrix()));
    return;
  }

  if (g_strcmp0(method_name, "GetBackendSelectionPlan") == 0) {
    const gchar *application_id = NULL;

    g_variant_get(parameters, "(&s)", &application_id);
    if (!known_application(application_id)) {
      return_unknown_application(invocation, application_id);
      return;
    }

    g_dbus_method_invocation_return_value(invocation, g_variant_new("(@a{sv})", build_backend_selection_plan(application_id)));
    return;
  }

  if (g_strcmp0(method_name, "GetBackendLifecycle") == 0) {
    const gchar *application_id = NULL;

    g_variant_get(parameters, "(&s)", &application_id);
    if (!known_application(application_id)) {
      return_unknown_application(invocation, application_id);
      return;
    }

    g_dbus_method_invocation_return_value(invocation, g_variant_new("(@a{sv})", build_backend_lifecycle(application_id)));
    return;
  }

  if (g_strcmp0(method_name, "GetBackendEnvironmentPlan") == 0) {
    const gchar *application_id = NULL;

    g_variant_get(parameters, "(&s)", &application_id);
    if (!known_application(application_id)) {
      return_unknown_application(invocation, application_id);
      return;
    }

    g_dbus_method_invocation_return_value(invocation, g_variant_new("(@a{sv})", build_backend_environment_plan(application_id)));
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

  if (g_strcmp0(method_name, "GetExecutionReadiness") == 0) {
    const gchar *application_id = NULL;

    g_variant_get(parameters, "(&s)", &application_id);
    if (!known_application(application_id)) {
      return_unknown_application(invocation, application_id);
      return;
    }

    g_dbus_method_invocation_return_value(invocation, g_variant_new("(@a{sv})", build_execution_readiness(application_id)));
    return;
  }

  if (g_strcmp0(method_name, "GetLaunchIntent") == 0) {
    const gchar *application_id = NULL;

    g_variant_get(parameters, "(&s)", &application_id);
    if (!known_application(application_id)) {
      return_unknown_application(invocation, application_id);
      return;
    }

    g_dbus_method_invocation_return_value(invocation, g_variant_new("(@a{sv})", build_launch_intent(application_id)));
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

  if (g_strcmp0(method_name, "GetRuntimeOwnerProcess") == 0) {
    g_dbus_method_invocation_return_value(invocation, g_variant_new("(@a{sv})", build_runtime_owner_process()));
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

  if (g_strcmp0(method_name, "GetRuntimeOwnerRouteManifest") == 0) {
    g_dbus_method_invocation_return_value(invocation, g_variant_new("(@a{sv})", build_runtime_owner_route_manifest()));
    return;
  }

  if (g_strcmp0(method_name, "GetRuntimeOwnerRecipeTrust") == 0) {
    g_dbus_method_invocation_return_value(invocation, g_variant_new("(@a{sv})", build_runtime_owner_recipe_trust()));
    return;
  }

  if (g_strcmp0(method_name, "GetRuntimeOwnerReadiness") == 0) {
    g_dbus_method_invocation_return_value(invocation, g_variant_new("(@a{sv})", build_runtime_owner_readiness()));
    return;
  }

  if (g_strcmp0(method_name, "GetRuntimeWriteGate") == 0) {
    const gchar *requested_method = NULL;

    g_variant_get(parameters, "(&s)", &requested_method);
    g_dbus_method_invocation_return_value(invocation, g_variant_new("(@a{sv})", build_runtime_write_gate(requested_method)));
    return;
  }

  if (g_strcmp0(method_name, "ShowRuntimeControlledLaunch") == 0) {
    const gchar *evidence_relative_path = NULL;

    g_variant_get(parameters, "(&s)", &evidence_relative_path);
    if (!safe_evidence_relative_path(evidence_relative_path)) {
      g_dbus_method_invocation_return_error(invocation,
                                            G_IO_ERROR,
                                            G_IO_ERROR_INVALID_ARGUMENT,
                                            "ShowRuntimeControlledLaunch requires a safe relative evidence path");
      return;
    }

    g_dbus_method_invocation_return_value(invocation, g_variant_new("(@a{sv})", build_runtime_controlled_launch_action(evidence_relative_path)));
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

  if (g_strcmp0(method_name, "GetKDECenterPage") == 0) {
    const gchar *application_id = NULL;
    const gchar *decision = NULL;

    g_variant_get(parameters, "(&s&s)", &application_id, &decision);
    if (!known_application(application_id)) {
      return_unknown_application(invocation, application_id);
      return;
    }

    g_dbus_method_invocation_return_value(
      invocation,
      g_variant_new("(@a{sv})", build_kde_center_page(application_id, decision))
    );
    return;
  }

  if (g_strcmp0(method_name, "GetKDECenterPageSections") == 0) {
    const gchar *application_id = NULL;
    const gchar *decision = NULL;

    g_variant_get(parameters, "(&s&s)", &application_id, &decision);
    if (!known_application(application_id)) {
      return_unknown_application(invocation, application_id);
      return;
    }

    g_dbus_method_invocation_return_value(
      invocation,
      g_variant_new("(@a{sv})", build_kde_center_page_sections(application_id, decision))
    );
    return;
  }

  if (g_strcmp0(method_name, "GetKDECenterPageSectionDetail") == 0) {
    const gchar *application_id = NULL;
    const gchar *section_id = NULL;
    const gchar *decision = NULL;
    GVariant *detail = NULL;

    g_variant_get(parameters, "(&s&s&s)", &application_id, &section_id, &decision);
    if (!known_application(application_id)) {
      return_unknown_application(invocation, application_id);
      return;
    }

    detail = build_kde_center_page_section_detail(application_id, section_id, decision);
    if (detail == NULL) {
      g_dbus_method_invocation_return_error(invocation,
                                            G_IO_ERROR,
                                            G_IO_ERROR_INVALID_ARGUMENT,
                                            "Unknown KDE center page section: %s",
                                            section_id);
      return;
    }

    g_dbus_method_invocation_return_value(
      invocation,
      g_variant_new("(@a{sv})", detail)
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
  .method_call = handle_method_call,
  .get_property = NULL,
  .set_property = NULL
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
