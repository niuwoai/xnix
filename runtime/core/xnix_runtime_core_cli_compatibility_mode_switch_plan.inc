static void
print_compatibility_mode_switch_application(const XnixRuntimeApplication *application)
{
  fputs("{", stdout);
  print_string_field("id", application->id);
  fputs(",", stdout);
  print_string_field("name", application->name);
  fputs(",", stdout);
  print_string_field("icon", application->icon);
  fputs("}", stdout);
}

static void
print_compatibility_mode_switch_modes(const XnixRuntimeCompatibilityModeSwitchPlan *plan)
{
  fputc('[', stdout);
  for (size_t index = 0; index < plan->mode_count; index++) {
    const XnixRuntimeCompatibilityModeOption *mode = &plan->modes[index];

    if (index > 0) {
      fputs(",", stdout);
    }

    fputs("{", stdout);
    print_string_field("id", mode->id);
    fputs(",", stdout);
    print_string_field("label", mode->label);
    fputs(",", stdout);
    print_string_field("intent", mode->intent);
    fputs(",", stdout);
    fputs("\"selected\":", stdout);
    print_bool(mode->selected);
    fputs(",", stdout);
    fputs("\"requested\":", stdout);
    print_bool(mode->requested);
    fputs(",", stdout);
    fputs("\"backend_details_exposed\":", stdout);
    print_bool(mode->backend_details_exposed);
    fputs("}", stdout);
  }
  fputc(']', stdout);
}

static void
print_compatibility_mode_switch_plan_record(const XnixRuntimeCompatibilityModeSwitchPlan *plan)
{
  fputs("{", stdout);
  printf("\"version\":\"%s\",", xnix_runtime_version());
  print_string_field("plan_type", plan->plan_type);
  fputs(",", stdout);
  print_string_field("runtime_method", "GetCompatibilityModeSwitchPlan");
  fputs(",", stdout);
  fputs("\"application\":", stdout);
  print_compatibility_mode_switch_application(plan->application);
  fputs(",", stdout);
  print_string_field("current_mode", plan->current_mode);
  fputs(",", stdout);
  print_string_field("requested_mode", plan->requested_mode);
  fputs(",", stdout);
  print_string_field("mode_state", plan->mode_state);
  fputs(",", stdout);
  fputs("\"modes\":", stdout);
  print_compatibility_mode_switch_modes(plan);
  fputs(",", stdout);
  fputs("\"mode_count\":", stdout);
  printf("%zu,", plan->mode_count);
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
  fputs("\"valid_mode\":", stdout);
  print_bool(plan->valid_mode);
  fputs(",", stdout);
  fputs("\"requires_user_confirmation\":", stdout);
  print_bool(plan->requires_user_confirmation);
  fputs(",", stdout);
  fputs("\"portal_review_required\":", stdout);
  print_bool(plan->portal_review_required);
  fputs(",", stdout);
  fputs("\"snapshot_required\":", stdout);
  print_bool(plan->snapshot_required);
  fputs(",", stdout);
  fputs("\"settings_persistence_enabled\":", stdout);
  print_bool(plan->settings_persistence_enabled);
  fputs(",", stdout);
  fputs("\"backend_reconfiguration_enabled\":", stdout);
  print_bool(plan->backend_reconfiguration_enabled);
  fputs(",", stdout);
  fputs("\"backend_process_started\":", stdout);
  print_bool(plan->backend_process_started);
  fputs(",", stdout);
  fputs("\"launch_enabled\":", stdout);
  print_bool(plan->launch_enabled);
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
print_compatibility_mode_switch_plan_for_application(const char *application_id, const char *requested_mode)
{
  XnixRuntimeCompatibilityModeSwitchPlan plan;

  if (!xnix_runtime_compatibility_mode_switch_plan(application_id, requested_mode, &plan)) {
    fprintf(stderr, "xnix-runtime-core: unknown application or unsupported compatibility mode\n");
    return 64;
  }

  print_compatibility_mode_switch_plan_record(&plan);
  fputs("\n", stdout);
  return 0;
}
