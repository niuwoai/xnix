static void
print_kde_integration_entry_point(const XnixRuntimeKDEIntegrationEntryPoint *entry)
{
  fputs("{", stdout);
  print_string_field("id", entry->id);
  fputs(",", stdout);
  print_string_field("name", entry->name);
  fputs(",", stdout);
  print_string_field("state", entry->state);
  fputs(",", stdout);
  print_string_field("model_command", entry->model_command);
  fputs(",", stdout);
  print_string_field("runtime_method", entry->runtime_method);
  fputs(",", stdout);
  print_string_field("adapter_role", entry->adapter_role);
  fputs(",", stdout);
  fputs("\"runtime_backed\":", stdout);
  print_bool(entry->runtime_backed);
  fputs(",", stdout);
  fputs("\"c_runtime_backed\":", stdout);
  print_bool(entry->c_runtime_backed);
  fputs(",", stdout);
  fputs("\"dbus_read_available\":", stdout);
  print_bool(entry->dbus_read_available);
  fputs(",", stdout);
  fputs("\"kde_policy_owner\":", stdout);
  print_bool(entry->kde_policy_owner);
  fputs(",", stdout);
  print_string_field("summary", entry->summary);
  fputs(",", stdout);
  print_string_field("next_step", entry->next_step);
  fputs("}", stdout);
}

static void
print_kde_integration_entry_points(const XnixRuntimeKDEIntegrationStatus *status)
{
  fputc('[', stdout);
  for (size_t index = 0; index < status->entry_point_count; index++) {
    if (index > 0) {
      fputs(",", stdout);
    }
    print_kde_integration_entry_point(&status->entry_points[index]);
  }
  fputc(']', stdout);
}

static void
print_kde_integration_summary(const XnixRuntimeKDEIntegrationStatus *status)
{
  fputs("{", stdout);
  fputs("\"total\":", stdout);
  printf("%zu", status->entry_point_count);
  fputs(",", stdout);
  fputs("\"initial\":", stdout);
  printf("%zu", status->initial_count);
  fputs(",", stdout);
  fputs("\"planned\":", stdout);
  printf("%zu", status->planned_count);
  fputs(",", stdout);
  fputs("\"complete\":", stdout);
  printf("%zu", status->complete_count);
  fputs("}", stdout);
}

static void
print_kde_integration_safety(const XnixRuntimeKDEIntegrationStatus *status)
{
  fputs("{", stdout);
  fputs("\"official_desktop_only\":", stdout);
  print_bool(status->official_desktop_only);
  fputs(",", stdout);
  fputs("\"stable_desktop_contract\":", stdout);
  print_bool(status->stable_desktop_contract);
  fputs(",", stdout);
  fputs("\"host_root_modified\":", stdout);
  print_bool(status->host_root_modified);
  fputs(",", stdout);
  fputs("\"backend_details_exposed\":", stdout);
  print_bool(status->backend_details_exposed);
  fputs("}", stdout);
}

static void
print_kde_integration_status_record(const XnixRuntimeKDEIntegrationStatus *status)
{
  fputs("{", stdout);
  printf("\"version\":\"%s\",", xnix_runtime_version());
  print_string_field("status_type", "kde-integration-status");
  fputs(",", stdout);
  print_string_field("desktop", status->desktop);
  fputs(",", stdout);
  fputs("\"runtime_owned\":", stdout);
  print_bool(status->runtime_owned);
  fputs(",", stdout);
  fputs("\"kde_policy_owner\":", stdout);
  print_bool(status->kde_policy_owner);
  fputs(",", stdout);
  fputs("\"official_desktop_only\":", stdout);
  print_bool(status->official_desktop_only);
  fputs(",", stdout);
  fputs("\"entry_points\":", stdout);
  print_kde_integration_entry_points(status);
  fputs(",", stdout);
  fputs("\"summary\":", stdout);
  print_kde_integration_summary(status);
  fputs(",", stdout);
  fputs("\"safety\":", stdout);
  print_kde_integration_safety(status);
  fputs(",", stdout);
  print_string_field("desktop_safe_summary", status->summary);
  fputs("}", stdout);
}

static int
print_kde_integration_status(void)
{
  XnixRuntimeKDEIntegrationStatus status;

  if (!xnix_runtime_kde_integration_status(&status)) {
    fprintf(stderr, "xnix-runtime-core: KDE integration status catalog is unavailable\n");
    return 64;
  }

  print_kde_integration_status_record(&status);
  fputs("\n", stdout);
  return 0;
}
