static void
print_backend_selection_application(const XnixRuntimeApplication *application)
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
print_backend_selection_candidates(const XnixRuntimeBackendSelectionPlan *plan)
{
  fputc('[', stdout);
  for (size_t index = 0; index < plan->candidate_count; index++) {
    const XnixRuntimeBackendSelectionCandidate *candidate = &plan->candidates[index];

    if (index > 0) {
      fputs(",", stdout);
    }

    fputs("{", stdout);
    print_string_field("id", candidate->id);
    fputs(",", stdout);
    print_string_field("label", candidate->label);
    fputs(",", stdout);
    print_string_field("kind", candidate->kind);
    fputs(",", stdout);
    print_string_field("selection_state", candidate->selection_state);
    fputs(",", stdout);
    fputs("\"recommended\":", stdout);
    print_bool(candidate->recommended);
    fputs(",", stdout);
    fputs("\"ready\":", stdout);
    print_bool(candidate->ready);
    fputs(",", stdout);
    fputs("\"blocked\":", stdout);
    print_bool(candidate->blocked);
    fputs(",", stdout);
    printf("\"ready_capability_count\":%zu,", candidate->ready_capability_count);
    printf("\"pending_capability_count\":%zu,", candidate->pending_capability_count);
    fputs("\"required_preflight\":", stdout);
    print_string_array(candidate->required_preflight, candidate->required_preflight_count);
    fputs(",", stdout);
    fputs("\"selection_committed\":", stdout);
    print_bool(candidate->selection_committed);
    fputs(",", stdout);
    fputs("\"environment_created\":", stdout);
    print_bool(candidate->environment_created);
    fputs(",", stdout);
    fputs("\"backend_process_started\":", stdout);
    print_bool(candidate->backend_process_started);
    fputs(",", stdout);
    fputs("\"backend_details_exposed\":", stdout);
    print_bool(candidate->backend_details_exposed);
    fputs(",", stdout);
    print_string_field("summary", candidate->summary);
    fputs("}", stdout);
  }
  fputc(']', stdout);
}

static void
print_backend_selection_plan_record(const XnixRuntimeBackendSelectionPlan *plan)
{
  fputs("{", stdout);
  printf("\"version\":\"%s\",", xnix_runtime_version());
  print_string_field("plan_type", plan->plan_type);
  fputs(",", stdout);
  print_string_field("runtime_method", "GetBackendSelectionPlan");
  fputs(",", stdout);
  fputs("\"application\":", stdout);
  print_backend_selection_application(plan->application);
  fputs(",", stdout);
  print_string_field("selected_strategy", plan->selected_strategy);
  fputs(",", stdout);
  print_string_field("recommended_profile_id", plan->recommended_profile_id);
  fputs(",", stdout);
  fputs("\"candidate_profiles\":", stdout);
  print_backend_selection_candidates(plan);
  fputs(",", stdout);
  printf("\"candidate_count\":%zu,", plan->candidate_count);
  printf("\"ready_candidate_count\":%zu,", plan->ready_candidate_count);
  printf("\"blocked_candidate_count\":%zu,", plan->blocked_candidate_count);
  fputs("\"required_reviews\":", stdout);
  print_string_array(plan->required_reviews, plan->required_review_count);
  fputs(",", stdout);
  fputs("\"blocked_actions\":", stdout);
  print_string_array(plan->blocked_actions, plan->blocked_action_count);
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
  fputs("\"selection_committed\":", stdout);
  print_bool(plan->selection_committed);
  fputs(",", stdout);
  fputs("\"selection_change_enabled\":", stdout);
  print_bool(plan->selection_change_enabled);
  fputs(",", stdout);
  fputs("\"backend_launch_enabled\":", stdout);
  print_bool(plan->backend_launch_enabled);
  fputs(",", stdout);
  fputs("\"capability_activation_enabled\":", stdout);
  print_bool(plan->capability_activation_enabled);
  fputs(",", stdout);
  fputs("\"environment_created\":", stdout);
  print_bool(plan->environment_created);
  fputs(",", stdout);
  fputs("\"request_object_created\":", stdout);
  print_bool(plan->request_object_created);
  fputs(",", stdout);
  fputs("\"state_root_created\":", stdout);
  print_bool(plan->state_root_created);
  fputs(",", stdout);
  fputs("\"snapshot_created\":", stdout);
  print_bool(plan->snapshot_created);
  fputs(",", stdout);
  fputs("\"host_root_modified\":", stdout);
  print_bool(plan->host_root_modified);
  fputs(",", stdout);
  fputs("\"privileged_container_required\":", stdout);
  print_bool(plan->privileged_container_required);
  fputs(",", stdout);
  fputs("\"backend_details_exposed\":", stdout);
  print_bool(plan->backend_details_exposed);
  fputs(",", stdout);
  print_string_field("desktop_safe_summary", plan->summary);
  fputs("}", stdout);
}

static int
print_backend_selection_plan_for_application(const char *application_id)
{
  XnixRuntimeBackendSelectionPlan plan;

  if (!xnix_runtime_backend_selection_plan(application_id, &plan)) {
    fprintf(stderr, "xnix-runtime-core: unknown application or backend selection plan catalog\n");
    return 64;
  }

  print_backend_selection_plan_record(&plan);
  fputs("\n", stdout);
  return 0;
}
