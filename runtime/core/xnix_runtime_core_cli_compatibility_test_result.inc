static void
print_compatibility_test_result_counts(const XnixRuntimeCompatibilityTestResult *result)
{
  fputs("{", stdout);
  fputs("\"total\":", stdout);
  printf("%zu,", result->counts.total);
  fputs("\"passed\":", stdout);
  printf("%zu,", result->counts.passed);
  fputs("\"pending\":", stdout);
  printf("%zu,", result->counts.pending);
  fputs("\"blocked\":", stdout);
  printf("%zu", result->counts.blocked);
  fputs("}", stdout);
}

static void
print_compatibility_test_step_result(const XnixRuntimeCompatibilityTestStepResult *step)
{
  fputs("{", stdout);
  print_string_field("id", step->id);
  fputs(",", stdout);
  print_string_field("title", step->title);
  fputs(",", stdout);
  print_string_field("status", step->status);
  fputs(",", stdout);
  print_string_field("result", step->result);
  fputs(",", stdout);
  fputs("\"evidence\":[", stdout);
  print_string(step->evidence);
  fputs("],", stdout);
  print_string_field("next_action", step->next_action);
  fputs("}", stdout);
}

static void
print_compatibility_test_step_results(const XnixRuntimeCompatibilityTestResult *result)
{
  fputc('[', stdout);
  for (size_t index = 0; index < result->step_result_count; index++) {
    if (index > 0) {
      fputs(",", stdout);
    }
    print_compatibility_test_step_result(&result->step_results[index]);
  }
  fputc(']', stdout);
}

static void
print_compatibility_test_result_artifacts(const XnixRuntimeCompatibilityTestResult *result)
{
  fputs("{", stdout);
  fputs("\"compatibility_center_card\":", stdout);
  print_bool(result->compatibility_center_card);
  fputs(",", stdout);
  fputs("\"diagnostics_record\":", stdout);
  print_bool(result->diagnostics_record);
  fputs(",", stdout);
  print_string_field("notification_event", result->notification_event);
  fputs(",", stdout);
  fputs("\"repair_plan_issue\":", stdout);
  if (result->repair_plan_issue == NULL) {
    fputs("null", stdout);
  } else {
    print_string(result->repair_plan_issue);
  }
  fputs("}", stdout);
}

static void
print_compatibility_test_result_record(const XnixRuntimeCompatibilityTestResult *result)
{
  fputs("{", stdout);
  printf("\"version\":\"%s\",", xnix_runtime_version());
  print_string_field("result_type", "compatibility-test-result");
  fputs(",", stdout);
  print_string_field("plan_type", "compatibility-test");
  fputs(",", stdout);
  print_string_field("test_type", result->test_type);
  fputs(",", stdout);
  print_string_field("application_id", result->application->id);
  fputs(",", stdout);
  print_string_field("name", result->application->name);
  fputs(",", stdout);
  print_string_field("runtime_method", "GetTestResult");
  fputs(",", stdout);
  fputs("\"runtime_owned\":", stdout);
  print_bool(result->runtime_owned);
  fputs(",", stdout);
  fputs("\"c_runtime_backed\":", stdout);
  print_bool(result->c_runtime_backed);
  fputs(",", stdout);
  fputs("\"kde_policy_owner\":", stdout);
  print_bool(result->kde_policy_owner);
  fputs(",", stdout);
  print_string_field("result_source", result->result_source);
  fputs(",", stdout);
  print_string_field("execution_state", result->execution_state);
  fputs(",", stdout);
  print_string_field("overall_status", result->overall_status);
  fputs(",", stdout);
  fputs("\"counts\":", stdout);
  print_compatibility_test_result_counts(result);
  fputs(",", stdout);
  fputs("\"step_results\":", stdout);
  print_compatibility_test_step_results(result);
  fputs(",", stdout);
  fputs("\"artifacts\":", stdout);
  print_compatibility_test_result_artifacts(result);
  fputs(",", stdout);
  fputs("\"safe_for_ai_diagnostics\":", stdout);
  print_bool(result->safe_for_ai_diagnostics);
  fputs(",", stdout);
  fputs("\"test_executed\":", stdout);
  print_bool(result->test_executed);
  fputs(",", stdout);
  fputs("\"host_root_modified\":", stdout);
  print_bool(result->host_root_modified);
  fputs(",", stdout);
  print_string_field("desktop_safe_summary", result->summary);
  fputs(",", stdout);
  fputs("\"backend_details_exposed\":", stdout);
  print_bool(result->backend_details_exposed);
  fputs("}", stdout);
}

static int
print_compatibility_test_result_for_application(const char *application_id, const char *test_type)
{
  XnixRuntimeCompatibilityTestResult result;

  if (!xnix_runtime_compatibility_test_result(application_id, test_type, &result)) {
    fprintf(stderr, "xnix-runtime-core: unknown application or compatibility test result catalog\n");
    return 64;
  }

  print_compatibility_test_result_record(&result);
  fputs("\n", stdout);
  return 0;
}
