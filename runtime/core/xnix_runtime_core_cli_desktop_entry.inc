static void
print_desktop_entry_application(const XnixRuntimeApplication *application)
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
print_desktop_entry_safety(const XnixRuntimeDesktopEntryPlan *plan)
{
  fputs("{", stdout);
  fputs("\"files_written\":", stdout);
  print_bool(plan->files_written);
  fputs(",", stdout);
  fputs("\"host_root_modified\":", stdout);
  print_bool(plan->host_root_modified);
  fputs(",", stdout);
  fputs("\"backend_command_exposed\":", stdout);
  print_bool(plan->backend_command_exposed);
  fputs(",", stdout);
  fputs("\"raw_windows_executable_exposed\":", stdout);
  print_bool(plan->raw_windows_executable_exposed);
  fputs(",", stdout);
  fputs("\"compatibility_storage_path_exposed\":", stdout);
  print_bool(plan->compatibility_storage_path_exposed);
  fputs(",", stdout);
  fputs("\"backend_details_exposed\":", stdout);
  print_bool(plan->backend_details_exposed);
  fputs("}", stdout);
}

static void
print_desktop_entry_plan_record(const XnixRuntimeDesktopEntryPlan *plan)
{
  fputs("{", stdout);
  printf("\"version\":\"%s\",", xnix_runtime_version());
  print_string_field("plan_type", "desktop-entry-plan");
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
  print_desktop_entry_application(plan->application);
  fputs(",", stdout);
  print_string_field("desktop_file", plan->desktop_file);
  fputs(",", stdout);
  print_string_field("relative_path", plan->relative_path);
  fputs(",", stdout);
  print_string_field("name", plan->name);
  fputs(",", stdout);
  print_string_field("comment", plan->comment);
  fputs(",", stdout);
  print_string_field("exec", plan->exec);
  fputs(",", stdout);
  print_string_field("icon", plan->icon);
  fputs(",", stdout);
  fputs("\"categories\":", stdout);
  print_string_array(plan->categories, plan->category_count);
  fputs(",", stdout);
  fputs("\"mime_types\":", stdout);
  print_string_array(plan->mime_types, plan->mime_type_count);
  fputs(",", stdout);
  print_string_field("startup_wm_class", plan->startup_wm_class);
  fputs(",", stdout);
  print_string_field("file_argument_mode", plan->file_argument_mode);
  fputs(",", stdout);
  fputs("\"standard_desktop_entry\":", stdout);
  print_bool(plan->standard_desktop_entry);
  fputs(",", stdout);
  fputs("\"launch_uses_runtime\":", stdout);
  print_bool(plan->launch_uses_runtime);
  fputs(",", stdout);
  fputs("\"accepts_file_uris\":", stdout);
  print_bool(plan->accepts_file_uris);
  fputs(",", stdout);
  fputs("\"user_visible\":", stdout);
  print_bool(plan->user_visible);
  fputs(",", stdout);
  fputs("\"startup_notify\":", stdout);
  print_bool(plan->startup_notify);
  fputs(",", stdout);
  fputs("\"terminal\":", stdout);
  print_bool(plan->terminal);
  fputs(",", stdout);
  fputs("\"no_display\":", stdout);
  print_bool(plan->no_display);
  fputs(",", stdout);
  fputs("\"safety\":", stdout);
  print_desktop_entry_safety(plan);
  fputs(",", stdout);
  print_string_field("desktop_safe_summary", plan->summary);
  fputs("}", stdout);
}

static int
print_desktop_entry_plan_for_application(const char *application_id)
{
  XnixRuntimeDesktopEntryPlan plan;

  if (!xnix_runtime_desktop_entry_plan(application_id, &plan)) {
    fprintf(stderr, "xnix-runtime-core: unknown application or desktop entry plan catalog\n");
    return 64;
  }

  print_desktop_entry_plan_record(&plan);
  fputs("\n", stdout);
  return 0;
}
