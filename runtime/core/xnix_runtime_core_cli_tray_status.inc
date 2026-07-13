static void
print_tray_runtime_activity(const XnixRuntimeTrayStatusPlan *plan)
{
  fputs("{", stdout);
  fputs("\"active_application_count\":", stdout);
  printf("%zu,", plan->active_application_count);
  fputs("\"attention_required_count\":", stdout);
  printf("%zu,", plan->attention_required_count);
  print_string_field("summary", plan->runtime_activity_summary);
  fputs("}", stdout);
}

static void
print_tray_compatibility_status(const XnixRuntimeTrayStatusPlan *plan)
{
  fputs("{", stdout);
  print_string_field("state", plan->compatibility_state);
  fputs(",", stdout);
  print_string_field("label", plan->compatibility_label);
  fputs("}", stdout);
}

static void
print_tray_bridge_status(const XnixRuntimeTrayStatusPlan *plan)
{
  fputs("{", stdout);
  fputs("\"bridged_tray_application_count\":", stdout);
  printf("%zu,", plan->bridged_tray_application_count);
  print_string_field("state", plan->tray_bridge_state);
  fputs(",", stdout);
  print_string_field("label", plan->tray_bridge_label);
  fputs("}", stdout);
}

static void
print_tray_status_safety(const XnixRuntimeTrayStatusPlan *plan)
{
  fputs("{", stdout);
  fputs("\"user_visible\":", stdout);
  print_bool(plan->user_visible);
  fputs(",", stdout);
  fputs("\"live_backend_bridge_enabled\":", stdout);
  print_bool(plan->live_backend_bridge_enabled);
  fputs(",", stdout);
  fputs("\"bridge_configuration_persisted\":", stdout);
  print_bool(plan->bridge_configuration_persisted);
  fputs(",", stdout);
  fputs("\"host_root_modified\":", stdout);
  print_bool(plan->host_root_modified);
  fputs(",", stdout);
  fputs("\"backend_details_exposed\":", stdout);
  print_bool(plan->backend_details_exposed);
  fputs("}", stdout);
}

static void
print_tray_status_plan_record(const XnixRuntimeTrayStatusPlan *plan)
{
  fputs("{", stdout);
  printf("\"version\":\"%s\",", xnix_runtime_version());
  print_string_field("status_type", plan->status_type);
  fputs(",", stdout);
  print_string_field("desktop", plan->desktop);
  fputs(",", stdout);
  fputs("\"runtime_owned\":", stdout);
  print_bool(plan->runtime_owned);
  fputs(",", stdout);
  fputs("\"kde_policy_owner\":", stdout);
  print_bool(plan->kde_policy_owner);
  fputs(",", stdout);
  fputs("\"runtime_activity\":", stdout);
  print_tray_runtime_activity(plan);
  fputs(",", stdout);
  fputs("\"compatibility_status\":", stdout);
  print_tray_compatibility_status(plan);
  fputs(",", stdout);
  fputs("\"tray_bridge\":", stdout);
  print_tray_bridge_status(plan);
  fputs(",", stdout);
  fputs("\"actions\":", stdout);
  print_string_array(plan->actions, plan->action_count);
  fputs(",", stdout);
  fputs("\"safety\":", stdout);
  print_tray_status_safety(plan);
  fputs(",", stdout);
  print_string_field("desktop_safe_summary", plan->summary);
  fputs("}", stdout);
}

static int
print_tray_status_plan(void)
{
  XnixRuntimeTrayStatusPlan plan;

  if (!xnix_runtime_tray_status_plan(&plan)) {
    fprintf(stderr, "xnix-runtime-core: tray status plan catalog is unavailable\n");
    return 64;
  }

  print_tray_status_plan_record(&plan);
  fputs("\n", stdout);
  return 0;
}
