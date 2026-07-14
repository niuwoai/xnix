static void
print_kde_application_surface_application(const XnixRuntimeApplication *application)
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
print_kde_application_surface_entry_points(const XnixRuntimeKDEApplicationSurfacePlan *plan)
{
  fputc('[', stdout);
  for (size_t index = 0; index < plan->entry_point_count; index++) {
    const XnixRuntimeKDEApplicationSurfaceEntryPoint *entry_point = &plan->entry_points[index];

    if (index > 0) {
      fputs(",", stdout);
    }

    fputs("{", stdout);
    print_string_field("id", entry_point->id);
    fputs(",", stdout);
    print_string_field("name", entry_point->name);
    fputs(",", stdout);
    print_string_field("runtime_method", entry_point->runtime_method);
    fputs(",", stdout);
    print_string_field("state", entry_point->state);
    fputs(",", stdout);
    fputs("\"runtime_backed\":", stdout);
    print_bool(entry_point->runtime_backed);
    fputs(",", stdout);
    fputs("\"c_runtime_backed\":", stdout);
    print_bool(entry_point->c_runtime_backed);
    fputs(",", stdout);
    fputs("\"kde_writes_policy\":", stdout);
    print_bool(entry_point->kde_writes_policy);
    fputs(",", stdout);
    fputs("\"backend_details_exposed\":", stdout);
    print_bool(entry_point->backend_details_exposed);
    fputs(",", stdout);
    print_string_field("summary", entry_point->summary);
    fputs("}", stdout);
  }
  fputc(']', stdout);
}

static void
print_kde_application_surface_record(const XnixRuntimeKDEApplicationSurfacePlan *plan)
{
  fputs("{", stdout);
  printf("\"version\":\"%s\",", xnix_runtime_version());
  print_string_field("plan_type", plan->plan_type);
  fputs(",", stdout);
  print_string_field("runtime_method", "GetKDEApplicationSurfacePlan");
  fputs(",", stdout);
  fputs("\"application\":", stdout);
  print_kde_application_surface_application(plan->application);
  fputs(",", stdout);
  print_string_field("surface_state", plan->surface_state);
  fputs(",", stdout);
  print_string_field("desktop_shell", plan->desktop_shell);
  fputs(",", stdout);
  fputs("\"entry_points\":", stdout);
  print_kde_application_surface_entry_points(plan);
  fputs(",", stdout);
  fputs("\"entry_point_count\":", stdout);
  printf("%zu,", plan->entry_point_count);
  fputs("\"required_runtime_gates\":", stdout);
  print_string_array(plan->required_runtime_gates, plan->required_runtime_gate_count);
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
  fputs("\"normal_linux_application_surface\":", stdout);
  print_bool(plan->normal_linux_application_surface);
  fputs(",", stdout);
  fputs("\"standard_launcher_visible\":", stdout);
  print_bool(plan->standard_launcher_visible);
  fputs(",", stdout);
  fputs("\"task_manager_identity_ready\":", stdout);
  print_bool(plan->task_manager_identity_ready);
  fputs(",", stdout);
  fputs("\"file_associations_planned\":", stdout);
  print_bool(plan->file_associations_planned);
  fputs(",", stdout);
  fputs("\"dolphin_action_planned\":", stdout);
  print_bool(plan->dolphin_action_planned);
  fputs(",", stdout);
  fputs("\"krunner_query_planned\":", stdout);
  print_bool(plan->krunner_query_planned);
  fputs(",", stdout);
  fputs("\"tray_status_planned\":", stdout);
  print_bool(plan->tray_status_planned);
  fputs(",", stdout);
  fputs("\"notification_route_planned\":", stdout);
  print_bool(plan->notification_route_planned);
  fputs(",", stdout);
  fputs("\"settings_surface_planned\":", stdout);
  print_bool(plan->settings_surface_planned);
  fputs(",", stdout);
  fputs("\"portal_review_required\":", stdout);
  print_bool(plan->portal_review_required);
  fputs(",", stdout);
  fputs("\"execution_ready\":", stdout);
  print_bool(plan->execution_ready);
  fputs(",", stdout);
  fputs("\"launch_enabled\":", stdout);
  print_bool(plan->launch_enabled);
  fputs(",", stdout);
  fputs("\"backend_process_started\":", stdout);
  print_bool(plan->backend_process_started);
  fputs(",", stdout);
  fputs("\"desktop_files_written\":", stdout);
  print_bool(plan->desktop_files_written);
  fputs(",", stdout);
  fputs("\"mimeapps_written\":", stdout);
  print_bool(plan->mimeapps_written);
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
  fputs("\"backend_details_exposed\":", stdout);
  print_bool(plan->backend_details_exposed);
  fputs(",", stdout);
  print_string_field("desktop_safe_summary", plan->summary);
  fputs("}", stdout);
}

static int
print_kde_application_surface_plan_for_application(const char *application_id)
{
  XnixRuntimeKDEApplicationSurfacePlan plan;

  if (!xnix_runtime_kde_application_surface_plan(application_id, &plan)) {
    fprintf(stderr, "xnix-runtime-core: unknown application or KDE application surface plan catalog\n");
    return 64;
  }

  print_kde_application_surface_record(&plan);
  fputs("\n", stdout);
  return 0;
}
