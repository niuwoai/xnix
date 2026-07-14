static void
print_backend_environment_application(const XnixRuntimeApplication *application)
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
print_backend_environment_profiles(const XnixRuntimeBackendEnvironmentPlan *plan)
{
  fputc('[', stdout);
  for (size_t index = 0; index < plan->profile_count; index++) {
    const XnixRuntimeBackendEnvironmentProfile *profile = &plan->profiles[index];

    if (index > 0) {
      fputs(",", stdout);
    }

    fputs("{", stdout);
    print_string_field("id", profile->id);
    fputs(",", stdout);
    print_string_field("kind", profile->kind);
    fputs(",", stdout);
    print_string_field("status", profile->status);
    fputs(",", stdout);
    fputs("\"selected\":", stdout);
    print_bool(profile->selected);
    fputs(",", stdout);
    fputs("\"environment_created\":", stdout);
    print_bool(profile->environment_created);
    fputs(",", stdout);
    fputs("\"process_started\":", stdout);
    print_bool(profile->process_started);
    fputs(",", stdout);
    fputs("\"host_storage_exposed\":", stdout);
    print_bool(profile->host_storage_exposed);
    fputs(",", stdout);
    fputs("\"backend_details_exposed\":", stdout);
    print_bool(profile->backend_details_exposed);
    fputs(",", stdout);
    print_string_field("summary", profile->summary);
    fputs("}", stdout);
  }
  fputc(']', stdout);
}

static void
print_backend_environment_record(const XnixRuntimeBackendEnvironmentPlan *plan)
{
  fputs("{", stdout);
  printf("\"version\":\"%s\",", xnix_runtime_version());
  print_string_field("plan_type", "compatibility-backend-environment-plan");
  fputs(",", stdout);
  print_string_field("runtime_method", "GetBackendEnvironmentPlan");
  fputs(",", stdout);
  fputs("\"application\":", stdout);
  print_backend_environment_application(plan->application);
  fputs(",", stdout);
  print_string_field("selected_strategy", plan->selected_strategy);
  fputs(",", stdout);
  print_string_field("environment_state", plan->environment_state);
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
  fputs("\"local_environment_ready\":", stdout);
  print_bool(plan->local_environment_ready);
  fputs(",", stdout);
  fputs("\"isolated_environment_ready\":", stdout);
  print_bool(plan->isolated_environment_ready);
  fputs(",", stdout);
  fputs("\"environment_created\":", stdout);
  print_bool(plan->environment_created);
  fputs(",", stdout);
  fputs("\"backend_process_started\":", stdout);
  print_bool(plan->backend_process_started);
  fputs(",", stdout);
  fputs("\"host_storage_exposed\":", stdout);
  print_bool(plan->host_storage_exposed);
  fputs(",", stdout);
  fputs("\"clipboard_bridge_enabled\":", stdout);
  print_bool(plan->clipboard_bridge_enabled);
  fputs(",", stdout);
  fputs("\"print_bridge_enabled\":", stdout);
  print_bool(plan->print_bridge_enabled);
  fputs(",", stdout);
  fputs("\"portal_review_required\":", stdout);
  print_bool(plan->portal_review_required);
  fputs(",", stdout);
  fputs("\"snapshot_required\":", stdout);
  print_bool(plan->snapshot_required);
  fputs(",", stdout);
  fputs("\"launch_enabled\":", stdout);
  print_bool(plan->launch_enabled);
  fputs(",", stdout);
  fputs("\"profiles\":", stdout);
  print_backend_environment_profiles(plan);
  fputs(",", stdout);
  fputs("\"required_reviews\":", stdout);
  print_string_array(plan->required_reviews, plan->required_review_count);
  fputs(",", stdout);
  fputs("\"blocked_actions\":", stdout);
  print_string_array(plan->blocked_actions, plan->blocked_action_count);
  fputs(",", stdout);
  fputs("\"host_root_modified\":", stdout);
  print_bool(plan->host_root_modified);
  fputs(",", stdout);
  fputs("\"network_required\":", stdout);
  print_bool(plan->network_required);
  fputs(",", stdout);
  fputs("\"backend_details_exposed\":", stdout);
  print_bool(plan->backend_details_exposed);
  fputs(",", stdout);
  print_string_field("desktop_safe_summary", plan->summary);
  fputs("}", stdout);
}

static int
print_backend_environment_plan_for_application(const char *application_id)
{
  XnixRuntimeBackendEnvironmentPlan plan;

  if (!xnix_runtime_backend_environment_plan(application_id, &plan)) {
    fprintf(stderr, "xnix-runtime-core: unknown application or backend environment plan catalog\n");
    return 64;
  }

  print_backend_environment_record(&plan);
  fputs("\n", stdout);
  return 0;
}
