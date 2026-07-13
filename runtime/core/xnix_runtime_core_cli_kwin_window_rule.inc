static void
print_kwin_window_rule_match(const XnixRuntimeKWinWindowRulePlan *plan)
{
  fputs("{", stdout);
  print_string_field("resource_name", plan->resource_name);
  fputs(",", stdout);
  print_string_field("class_group", plan->class_group);
  fputs(",", stdout);
  print_string_field("title_hint", plan->title_hint);
  fputs("}", stdout);
}

static void
print_kwin_window_rule_set(const XnixRuntimeKWinWindowRulePlan *plan)
{
  fputs("{", stdout);
  print_string_field("desktop_file", plan->desktop_file);
  fputs(",", stdout);
  print_string_field("application_id", plan->application_id);
  fputs(",", stdout);
  print_string_field("task_manager_grouping_key", plan->task_manager_grouping_key);
  fputs(",", stdout);
  print_string_field("launcher_url", plan->launcher_url);
  fputs(",", stdout);
  fputs("\"skip_taskbar\":", stdout);
  print_bool(plan->skip_taskbar);
  fputs(",", stdout);
  fputs("\"show_in_switcher\":", stdout);
  print_bool(plan->show_in_switcher);
  fputs(",", stdout);
  print_string_field("placement", plan->placement);
  fputs("}", stdout);
}

static void
print_kwin_window_rule_restore(const XnixRuntimeKWinWindowRulePlan *plan)
{
  fputs("{", stdout);
  fputs("\"pinning_allowed\":", stdout);
  print_bool(plan->pinning_allowed);
  fputs(",", stdout);
  fputs("\"restore_allowed\":", stdout);
  print_bool(plan->restore_allowed);
  fputs(",", stdout);
  print_string_field("restore_key", plan->restore_key);
  fputs(",", stdout);
  fputs("\"prefer_existing_window\":", stdout);
  print_bool(plan->prefer_existing_window);
  fputs("}", stdout);
}

static void
print_kwin_window_rule_safety(const XnixRuntimeKWinWindowRulePlan *plan)
{
  fputs("{", stdout);
  fputs("\"window_manager_policy_only\":", stdout);
  print_bool(plan->window_manager_policy_only);
  fputs(",", stdout);
  fputs("\"runtime_owns_backend_policy\":", stdout);
  print_bool(plan->runtime_owns_backend_policy);
  fputs(",", stdout);
  fputs("\"host_root_modified\":", stdout);
  print_bool(plan->host_root_modified);
  fputs(",", stdout);
  fputs("\"backend_details_exposed\":", stdout);
  print_bool(plan->backend_details_exposed);
  fputs("}", stdout);
}

static void
print_kwin_window_rule_record(const XnixRuntimeKWinWindowRulePlan *plan)
{
  fputs("{", stdout);
  printf("\"version\":\"%s\",", xnix_runtime_version());
  print_string_field("request_type", plan->request_type);
  fputs(",", stdout);
  print_string_field("desktop", plan->desktop);
  fputs(",", stdout);
  fputs("\"runtime_owned\":", stdout);
  print_bool(plan->runtime_owned);
  fputs(",", stdout);
  fputs("\"kde_policy_owner\":", stdout);
  print_bool(plan->kde_policy_owner);
  fputs(",", stdout);
  print_string_field("application_id", plan->application_id);
  fputs(",", stdout);
  print_string_field("name", plan->application->name);
  fputs(",", stdout);
  print_string_field("script_role", plan->script_role);
  fputs(",", stdout);
  fputs("\"match\":", stdout);
  print_kwin_window_rule_match(plan);
  fputs(",", stdout);
  fputs("\"set\":", stdout);
  print_kwin_window_rule_set(plan);
  fputs(",", stdout);
  fputs("\"restore\":", stdout);
  print_kwin_window_rule_restore(plan);
  fputs(",", stdout);
  fputs("\"safety\":", stdout);
  print_kwin_window_rule_safety(plan);
  fputs(",", stdout);
  fputs("\"host_root_modified\":", stdout);
  print_bool(plan->host_root_modified);
  fputs(",", stdout);
  fputs("\"backend_details_exposed\":", stdout);
  print_bool(plan->backend_details_exposed);
  fputs(",", stdout);
  print_string_field("desktop_safe_summary", plan->summary);
  fputs("}", stdout);
}

static int
print_kwin_window_rule_plan_for_application(const char *application_id)
{
  XnixRuntimeKWinWindowRulePlan plan;

  if (!xnix_runtime_kwin_window_rule_plan(application_id, &plan)) {
    fprintf(stderr, "xnix-runtime-core: unknown application or KWin window rule plan catalog\n");
    return 64;
  }

  print_kwin_window_rule_record(&plan);
  fputs("\n", stdout);
  return 0;
}
