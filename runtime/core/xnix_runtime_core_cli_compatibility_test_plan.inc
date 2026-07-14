static void
print_test_step_required_for(const XnixRuntimeCompatibilityTestStep *step)
{
  fputc('[', stdout);
  for (size_t index = 0; index < step->required_for_count; index++) {
    if (index > 0) {
      fputs(",", stdout);
    }
    print_string(step->required_for[index]);
  }
  fputc(']', stdout);
}

static void
print_compatibility_test_step(const XnixRuntimeCompatibilityTestStep *step)
{
  fputs("{", stdout);
  print_string_field("id", step->id);
  fputs(",", stdout);
  print_string_field("status", step->status);
  fputs(",", stdout);
  print_string_field("title", step->title);
  fputs(",", stdout);
  print_string_field("summary", step->summary);

  if (step->required_for_count > 0) {
    fputs(",", stdout);
    fputs("\"required_for\":", stdout);
    print_test_step_required_for(step);
  }

  if (step->portal_method != NULL) {
    fputs(",", stdout);
    fputs("\"portal_request\":{", stdout);
    print_string_field("destination", step->portal_destination);
    fputs(",", stdout);
    print_string_field("interface", step->portal_interface);
    fputs(",", stdout);
    print_string_field("method", step->portal_method);
    fputs(",", stdout);
    print_string_field("handle_token", step->portal_handle_token);
    fputs("}", stdout);
  }

  if (step->snapshot_reason != NULL) {
    fputs(",", stdout);
    fputs("\"snapshot\":{", stdout);
    print_string_field("reason", step->snapshot_reason);
    fputs(",", stdout);
    fputs("\"enabled_by_default\":", stdout);
    print_bool(step->snapshot_enabled_by_default);
    fputs(",", stdout);
    fputs("\"restore_available\":", stdout);
    print_bool(step->snapshot_restore_available);
    fputs("}", stdout);
  }

  if (step->run_strategy != NULL) {
    fputs(",", stdout);
    fputs("\"run_plan\":{", stdout);
    print_string_field("strategy", step->run_strategy);
    fputs(",", stdout);
    fputs("\"backend_ready\":", stdout);
    print_bool(step->backend_ready);
    fputs(",", stdout);
    fputs("\"backend_details_exposed\":", stdout);
    print_bool(step->run_plan_backend_details_exposed);
    fputs("}", stdout);
  }

  fputs("}", stdout);
}

static void
print_compatibility_test_steps(const XnixRuntimeCompatibilityTestPlan *plan)
{
  fputc('[', stdout);
  for (size_t index = 0; index < plan->step_count; index++) {
    if (index > 0) {
      fputs(",", stdout);
    }
    print_compatibility_test_step(&plan->steps[index]);
  }
  fputc(']', stdout);
}

static void
print_compatibility_test_blocking_reasons(const XnixRuntimeCompatibilityTestPlan *plan)
{
  fputc('[', stdout);
  for (size_t index = 0; index < plan->blocking_reason_count; index++) {
    if (index > 0) {
      fputs(",", stdout);
    }
    print_string(plan->blocking_reasons[index]);
  }
  fputc(']', stdout);
}

static void
print_compatibility_test_artifacts(const XnixRuntimeCompatibilityTestPlan *plan)
{
  fputs("{", stdout);
  fputs("\"compatibility_center_card\":", stdout);
  print_bool(plan->compatibility_center_card);
  fputs(",", stdout);
  print_string_field("notification_event", plan->notification_event);
  fputs(",", stdout);
  print_string_field("repair_plan_issue", plan->repair_plan_issue);
  fputs("}", stdout);
}

static void
print_compatibility_test_plan_record(const XnixRuntimeCompatibilityTestPlan *plan)
{
  fputs("{", stdout);
  printf("\"version\":\"%s\",", xnix_runtime_version());
  print_string_field("plan_type", "compatibility-test");
  fputs(",", stdout);
  print_string_field("test_type", plan->test_type);
  fputs(",", stdout);
  print_string_field("desktop", plan->desktop);
  fputs(",", stdout);
  print_string_field("application_id", plan->application->id);
  fputs(",", stdout);
  print_string_field("name", plan->application->name);
  fputs(",", stdout);
  print_string_field("runtime_method", "GetTestPlan");
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
  fputs("\"steps\":", stdout);
  print_compatibility_test_steps(plan);
  fputs(",", stdout);
  fputs("\"blocked\":", stdout);
  print_bool(plan->blocked);
  fputs(",", stdout);
  fputs("\"blocking_reasons\":", stdout);
  print_compatibility_test_blocking_reasons(plan);
  fputs(",", stdout);
  fputs("\"artifacts\":", stdout);
  print_compatibility_test_artifacts(plan);
  fputs(",", stdout);
  fputs("\"execution_request_created\":", stdout);
  print_bool(plan->execution_request_created);
  fputs(",", stdout);
  fputs("\"test_executed\":", stdout);
  print_bool(plan->test_executed);
  fputs(",", stdout);
  fputs("\"host_root_modified\":", stdout);
  print_bool(plan->host_root_modified);
  fputs(",", stdout);
  print_string_field("desktop_safe_summary", plan->summary);
  fputs(",", stdout);
  fputs("\"backend_details_exposed\":", stdout);
  print_bool(plan->backend_details_exposed);
  fputs("}", stdout);
}

static int
print_compatibility_test_plan_for_application(const char *application_id, const char *test_type)
{
  XnixRuntimeCompatibilityTestPlan plan;

  if (!xnix_runtime_compatibility_test_plan(application_id, test_type, &plan)) {
    fprintf(stderr, "xnix-runtime-core: unknown application or compatibility test type catalog\n");
    return 64;
  }

  print_compatibility_test_plan_record(&plan);
  fputs("\n", stdout);
  return 0;
}
