static void
print_ai_diagnostic_recommendation_item(const XnixRuntimeAIDiagnosticRecommendationItem *item)
{
  fputs("{", stdout);
  print_string_field("id", item->id);
  fputs(",", stdout);
  print_string_field("priority", item->priority);
  fputs(",", stdout);
  print_string_field("source_signal", item->source_signal);
  fputs(",", stdout);
  print_string_field("title", item->title);
  fputs(",", stdout);
  print_string_field("summary", item->summary);
  fputs(",", stdout);
  fputs("\"user_visible\":", stdout);
  print_bool(item->user_visible);
  fputs(",", stdout);
  fputs("\"auto_execute\":", stdout);
  print_bool(item->auto_execute);
  fputs(",", stdout);
  fputs("\"requires_user_approval\":", stdout);
  print_bool(item->requires_user_approval);
  fputs(",", stdout);
  print_string_field("next_runtime_action", item->next_runtime_action);
  fputs("}", stdout);
}

static void
print_ai_diagnostic_recommendations(const XnixRuntimeAIDiagnosticRecommendation *recommendation)
{
  fputc('[', stdout);
  for (size_t index = 0; index < recommendation->recommendation_count; index++) {
    if (index > 0) {
      fputs(",", stdout);
    }
    print_ai_diagnostic_recommendation_item(&recommendation->recommendations[index]);
  }
  fputc(']', stdout);
}

static void
print_ai_approval_action(const XnixRuntimeAIApprovalRequiredAction *action)
{
  fputs("{", stdout);
  print_string_field("id", action->id);
  fputs(",", stdout);
  print_string_field("title", action->title);
  fputs(",", stdout);
  print_string_field("reason", action->reason);
  fputs(",", stdout);
  print_string_field("approval_surface", action->approval_surface);
  fputs("}", stdout);
}

static void
print_ai_approval_actions(const XnixRuntimeAIDiagnosticRecommendation *recommendation)
{
  fputc('[', stdout);
  for (size_t index = 0; index < recommendation->approval_required_action_count; index++) {
    if (index > 0) {
      fputs(",", stdout);
    }
    print_ai_approval_action(&recommendation->approval_required_actions[index]);
  }
  fputc(']', stdout);
}

static void
print_ai_blocked_actions(const XnixRuntimeAIDiagnosticRecommendation *recommendation)
{
  print_ai_task_list(
    "blocked_actions",
    recommendation->blocked_actions,
    recommendation->blocked_action_count
  );
}

static void
print_ai_diagnostic_recommendation_record(const XnixRuntimeAIDiagnosticRecommendation *recommendation)
{
  const XnixRuntimeAIDiagnosticInput *input = &recommendation->diagnostic_input;

  fputs("{", stdout);
  printf("\"version\":\"%s\",", xnix_runtime_version());
  print_string_field("recommendation_type", recommendation->recommendation_type);
  fputs(",", stdout);
  print_string_field("input_type", input->input_type);
  fputs(",", stdout);
  fputs("\"application\":", stdout);
  print_ai_diagnostic_application(input->application);
  fputs(",", stdout);
  print_string_field("issue", input->issue);
  fputs(",", stdout);
  print_string_field("test_type", input->test_type);
  fputs(",", stdout);
  print_string_field("runtime_method", "GetAIDiagnosticRecommendation");
  fputs(",", stdout);
  fputs("\"runtime_owned\":", stdout);
  print_bool(recommendation->runtime_owned);
  fputs(",", stdout);
  fputs("\"c_runtime_backed\":", stdout);
  print_bool(recommendation->c_runtime_backed);
  fputs(",", stdout);
  fputs("\"kde_policy_owner\":", stdout);
  print_bool(recommendation->kde_policy_owner);
  fputs(",", stdout);
  fputs("\"ai_provider_called\":", stdout);
  print_bool(recommendation->ai_provider_called);
  fputs(",", stdout);
  fputs("\"network_required\":", stdout);
  print_bool(recommendation->network_required);
  fputs(",", stdout);
  fputs("\"safe_for_ai_diagnostics\":", stdout);
  print_bool(recommendation->safe_for_ai_diagnostics);
  fputs(",", stdout);
  fputs("\"auto_execution_allowed\":", stdout);
  print_bool(recommendation->auto_execution_allowed);
  fputs(",", stdout);
  fputs("\"recommendations\":", stdout);
  print_ai_diagnostic_recommendations(recommendation);
  fputs(",", stdout);
  fputs("\"approval_required_actions\":", stdout);
  print_ai_approval_actions(recommendation);
  fputs(",", stdout);
  print_ai_blocked_actions(recommendation);
  fputs(",", stdout);
  fputs("\"privacy_boundaries\":", stdout);
  print_ai_diagnostic_privacy_boundaries(&input->privacy_boundaries);
  fputs(",", stdout);
  fputs("\"host_root_modified\":", stdout);
  print_bool(recommendation->host_root_modified);
  fputs(",", stdout);
  fputs("\"backend_details_exposed\":", stdout);
  print_bool(recommendation->backend_details_exposed);
  fputs(",", stdout);
  print_string_field("desktop_safe_summary", recommendation->summary);
  fputs("}", stdout);
}

static int
print_ai_diagnostic_recommendation_for_application(const char *application_id, const char *issue, const char *test_type)
{
  XnixRuntimeAIDiagnosticRecommendation recommendation;

  if (!xnix_runtime_ai_diagnostic_recommendation(application_id, issue, test_type, &recommendation)) {
    fprintf(stderr, "xnix-runtime-core: unknown application or AI diagnostic recommendation catalog\n");
    return 64;
  }

  print_ai_diagnostic_recommendation_record(&recommendation);
  fputs("\n", stdout);
  return 0;
}
