static void
print_compatibility_center_application(const XnixRuntimeCompatibilityCenterSummary *summary)
{
  fputs("{", stdout);
  print_string_field("id", summary->application->id);
  fputs(",", stdout);
  print_string_field("name", summary->application->name);
  fputs(",", stdout);
  print_string_field("icon", summary->application->icon);
  fputs(",", stdout);
  print_string_field("runtime_mode", summary->runtime_mode);
  fputs("}", stdout);
}

static void
print_compatibility_center_known_issue(const XnixRuntimeCompatibilityCenterSummary *summary)
{
  fputs("{", stdout);
  print_string_field("id", summary->known_issue_id);
  fputs(",", stdout);
  print_string_field("severity", summary->known_issue_severity);
  fputs(",", stdout);
  print_string_field("summary", summary->known_issue_summary);
  fputs("}", stdout);
}

static void
print_compatibility_center_compatibility(const XnixRuntimeCompatibilityCenterSummary *summary)
{
  fputs("{", stdout);
  print_string_field("state", summary->compatibility_state);
  fputs(",", stdout);
  print_string_field("runtime_mode", summary->runtime_mode);
  fputs(",", stdout);
  fputs("\"known_issue_count\":", stdout);
  printf("%zu,", summary->known_issue_count);
  fputs("\"known_issues\":[", stdout);
  print_compatibility_center_known_issue(summary);
  fputs("]}", stdout);
}

static void
print_compatibility_center_repair_record(const XnixRuntimeCompatibilityCenterSummary *summary)
{
  fputs("{", stdout);
  print_string_field("id", summary->repair_record_id);
  fputs(",", stdout);
  print_string_field("state", summary->repair_record_state);
  fputs(",", stdout);
  print_string_field("last_event", summary->last_repair_event);
  fputs("}", stdout);
}

static void
print_compatibility_center_repair_records(const XnixRuntimeCompatibilityCenterSummary *summary)
{
  fputs("{", stdout);
  print_string_field("state", summary->repair_record_state);
  fputs(",", stdout);
  print_string_field("last_event", summary->last_repair_event);
  fputs(",", stdout);
  fputs("\"record_count\":", stdout);
  printf("%zu,", summary->repair_record_count);
  fputs("\"records\":[", stdout);
  print_compatibility_center_repair_record(summary);
  fputs("]}", stdout);
}

static void
print_compatibility_center_safety(const XnixRuntimeCompatibilityCenterSummary *summary)
{
  fputs("{", stdout);
  fputs("\"user_visible\":", stdout);
  print_bool(summary->user_visible);
  fputs(",", stdout);
  fputs("\"action_execution_enabled\":", stdout);
  print_bool(summary->action_execution_enabled);
  fputs(",", stdout);
  fputs("\"repair_execution_enabled\":", stdout);
  print_bool(summary->repair_execution_enabled);
  fputs(",", stdout);
  fputs("\"backend_launch_enabled\":", stdout);
  print_bool(summary->backend_launch_enabled);
  fputs(",", stdout);
  fputs("\"settings_persistence_enabled\":", stdout);
  print_bool(summary->settings_persistence_enabled);
  fputs(",", stdout);
  fputs("\"host_root_modified\":", stdout);
  print_bool(summary->host_root_modified);
  fputs(",", stdout);
  fputs("\"backend_details_exposed\":", stdout);
  print_bool(summary->backend_details_exposed);
  fputs("}", stdout);
}

static void
print_compatibility_center_summary_record(const XnixRuntimeCompatibilityCenterSummary *summary)
{
  fputs("{", stdout);
  printf("\"version\":\"%s\",", xnix_runtime_version());
  print_string_field("summary_type", summary->summary_type);
  fputs(",", stdout);
  print_string_field("desktop", summary->desktop);
  fputs(",", stdout);
  fputs("\"runtime_owned\":", stdout);
  print_bool(summary->runtime_owned);
  fputs(",", stdout);
  fputs("\"kde_policy_owner\":", stdout);
  print_bool(summary->kde_policy_owner);
  fputs(",", stdout);
  fputs("\"application\":", stdout);
  print_compatibility_center_application(summary);
  fputs(",", stdout);
  fputs("\"compatibility\":", stdout);
  print_compatibility_center_compatibility(summary);
  fputs(",", stdout);
  fputs("\"repair_records\":", stdout);
  print_compatibility_center_repair_records(summary);
  fputs(",", stdout);
  fputs("\"actions\":", stdout);
  print_string_array(summary->actions, summary->action_count);
  fputs(",", stdout);
  fputs("\"safety\":", stdout);
  print_compatibility_center_safety(summary);
  fputs(",", stdout);
  print_string_field("desktop_safe_summary", summary->summary);
  fputs("}", stdout);
}

static int
print_compatibility_center_summary_for_application(const char *application_id)
{
  XnixRuntimeCompatibilityCenterSummary summary;

  if (!xnix_runtime_compatibility_center_summary(application_id, &summary)) {
    fprintf(stderr, "xnix-runtime-core: application is not registered in the C Runtime compatibility center catalog\n");
    return 64;
  }

  print_compatibility_center_summary_record(&summary);
  fputs("\n", stdout);
  return 0;
}
