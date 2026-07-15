static void
print_desktop_icon_application(const XnixRuntimeApplication *application)
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
print_desktop_icon_safety(const XnixRuntimeDesktopIconPlan *plan)
{
  fputs("{", stdout);
  fputs("\"desktop_file_copy_enabled\":", stdout);
  print_bool(plan->desktop_file_copy_enabled);
  fputs(",", stdout);
  fputs("\"desktop_file_write_enabled\":", stdout);
  print_bool(plan->desktop_file_write_enabled);
  fputs(",", stdout);
  fputs("\"icon_placement_persisted\":", stdout);
  print_bool(plan->icon_placement_persisted);
  fputs(",", stdout);
  fputs("\"launch_enabled\":", stdout);
  print_bool(plan->launch_enabled);
  fputs(",", stdout);
  fputs("\"backend_launch_enabled\":", stdout);
  print_bool(plan->backend_launch_enabled);
  fputs(",", stdout);
  fputs("\"host_root_modified\":", stdout);
  print_bool(plan->host_root_modified);
  fputs(",", stdout);
  fputs("\"raw_executable_exposed\":", stdout);
  print_bool(plan->raw_executable_exposed);
  fputs(",", stdout);
  fputs("\"backend_details_exposed\":", stdout);
  print_bool(plan->backend_details_exposed);
  fputs("}", stdout);
}

static void
print_desktop_icon_plan_record(const XnixRuntimeDesktopIconPlan *plan)
{
  fputs("{", stdout);
  printf("\"version\":\"%s\",", xnix_runtime_version());
  print_string_field("plan_type", plan->plan_type);
  fputs(",", stdout);
  print_string_field("desktop", "KDE Plasma");
  fputs(",", stdout);
  print_string_field("runtime_method", plan->runtime_method);
  fputs(",", stdout);
  fputs("\"runtime_owned\":", stdout);
  print_bool(plan->runtime_owned);
  fputs(",", stdout);
  fputs("\"c_runtime_backed\":", stdout);
  print_bool(plan->c_runtime_backed);
  fputs(",", stdout);
  fputs("\"kde_policy_owner\":", stdout);
  print_bool(plan->kde_policy_owner);
  fputs(",", stdout);
  fputs("\"official_desktop_only\":", stdout);
  print_bool(plan->official_desktop_only);
  fputs(",", stdout);
  fputs("\"application\":", stdout);
  print_desktop_icon_application(plan->application);
  fputs(",", stdout);
  print_string_field("desktop_file", plan->desktop_file);
  fputs(",", stdout);
  print_string_field("launcher_url", plan->launcher_url);
  fputs(",", stdout);
  print_string_field("target_directory", plan->target_directory);
  fputs(",", stdout);
  print_string_field("placement", plan->placement);
  fputs(",", stdout);
  print_string_field("exec", plan->exec);
  fputs(",", stdout);
  print_string_field("icon", plan->icon);
  fputs(",", stdout);
  fputs("\"standard_desktop_entry\":", stdout);
  print_bool(plan->standard_desktop_entry);
  fputs(",", stdout);
  fputs("\"desktop_icon_visible\":", stdout);
  print_bool(plan->desktop_icon_visible);
  fputs(",", stdout);
  fputs("\"user_visible\":", stdout);
  print_bool(plan->user_visible);
  fputs(",", stdout);
  fputs("\"blocked_actions\":", stdout);
  print_string_array(plan->blocked_actions, plan->blocked_action_count);
  fputs(",", stdout);
  fputs("\"safety\":", stdout);
  print_desktop_icon_safety(plan);
  fputs(",", stdout);
  print_string_field("desktop_safe_summary", plan->summary);
  fputs("}", stdout);
}

static int
print_desktop_icon_plan_for_application(const char *application_id)
{
  XnixRuntimeDesktopIconPlan plan;

  if (!xnix_runtime_desktop_icon_plan(application_id, &plan)) {
    fprintf(stderr, "xnix-runtime-core: unknown application or desktop icon plan catalog\n");
    return 64;
  }

  print_desktop_icon_plan_record(&plan);
  fputs("\n", stdout);
  return 0;
}
