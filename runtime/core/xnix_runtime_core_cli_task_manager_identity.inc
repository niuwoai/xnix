static void
print_task_manager_application(const XnixRuntimeApplication *application)
{
  fputs("{", stdout);
  print_string_field("id", application->id);
  fputs(",", stdout);
  print_string_field("name", application->name);
  fputs(",", stdout);
  print_string_field("icon", application->icon);
  fputs(",", stdout);
  print_string_field("runtime_mode", application->runtime_mode);
  fputs("}", stdout);
}

static void
print_task_manager_window(const XnixRuntimeTaskManagerIdentityPlan *plan)
{
  fputs("{", stdout);
  print_string_field("kind", plan->window_kind);
  fputs(",", stdout);
  print_string_field("class_group", plan->class_group);
  fputs(",", stdout);
  print_string_field("resource_name", plan->resource_name);
  fputs(",", stdout);
  print_string_field("title_hint", plan->title_hint);
  fputs("}", stdout);
}

static void
print_task_manager_taskbar(const XnixRuntimeTaskManagerIdentityPlan *plan)
{
  fputs("{", stdout);
  fputs("\"pinning_allowed\":", stdout);
  print_bool(plan->pinning_allowed);
  fputs(",", stdout);
  fputs("\"restore_allowed\":", stdout);
  print_bool(plan->restore_allowed);
  fputs(",", stdout);
  print_string_field("grouping_key", plan->grouping_key);
  fputs(",", stdout);
  print_string_field("launcher_url", plan->launcher_url);
  fputs(",", stdout);
  fputs("\"skip_taskbar\":", stdout);
  print_bool(plan->skip_taskbar);
  fputs(",", stdout);
  fputs("\"show_in_switcher\":", stdout);
  print_bool(plan->show_in_switcher);
  fputs("}", stdout);
}

static void
print_task_manager_kwin(const XnixRuntimeTaskManagerIdentityPlan *plan)
{
  fputs("{", stdout);
  print_string_field("script_role", plan->kwin_script_role);
  fputs(",", stdout);
  print_string_field("placement", plan->placement);
  fputs(",", stdout);
  fputs("\"match\":{", stdout);
  print_string_field("resource_name", plan->resource_name);
  fputs(",", stdout);
  print_string_field("class_group", plan->class_group);
  fputs("},", stdout);
  fputs("\"set\":{", stdout);
  print_string_field("desktop_file", plan->desktop_file);
  fputs(",", stdout);
  print_string_field("application_id", plan->application->id);
  fputs("}", stdout);
  fputs("}", stdout);
}

static void
print_task_manager_restore(const XnixRuntimeTaskManagerIdentityPlan *plan)
{
  fputs("{", stdout);
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
print_task_manager_safety(const XnixRuntimeTaskManagerIdentityPlan *plan)
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
print_task_manager_identity_record(const XnixRuntimeTaskManagerIdentityPlan *plan)
{
  fputs("{", stdout);
  printf("\"version\":\"%s\",", xnix_runtime_version());
  print_string_field("plan_type", "task-manager-identity-plan");
  fputs(",", stdout);
  print_string_field("desktop", "KDE Plasma");
  fputs(",", stdout);
  fputs("\"runtime_owned\":", stdout);
  print_bool(plan->runtime_owned);
  fputs(",", stdout);
  fputs("\"kde_policy_owner\":", stdout);
  print_bool(plan->kde_policy_owner);
  fputs(",", stdout);
  fputs("\"application\":", stdout);
  print_task_manager_application(plan->application);
  fputs(",", stdout);
  print_string_field("desktop_file", plan->desktop_file);
  fputs(",", stdout);
  fputs("\"window\":", stdout);
  print_task_manager_window(plan);
  fputs(",", stdout);
  fputs("\"task_manager\":", stdout);
  print_task_manager_taskbar(plan);
  fputs(",", stdout);
  fputs("\"kwin\":", stdout);
  print_task_manager_kwin(plan);
  fputs(",", stdout);
  fputs("\"restore\":", stdout);
  print_task_manager_restore(plan);
  fputs(",", stdout);
  fputs("\"safety\":", stdout);
  print_task_manager_safety(plan);
  fputs(",", stdout);
  print_string_field("desktop_safe_summary", plan->summary);
  fputs("}", stdout);
}

static int
print_task_manager_identity_plan_for_application(const char *application_id)
{
  XnixRuntimeTaskManagerIdentityPlan plan;

  if (!xnix_runtime_task_manager_identity_plan(application_id, &plan)) {
    fprintf(stderr, "xnix-runtime-core: unknown application or task manager identity plan catalog\n");
    return 64;
  }

  print_task_manager_identity_record(&plan);
  fputs("\n", stdout);
  return 0;
}
