static void
print_ai_diagnostic_application(const XnixRuntimeApplication *application)
{
  fputs("{", stdout);
  print_string_field("id", application->id);
  fputs(",", stdout);
  print_string_field("name", application->name);
  fputs(",", stdout);
  print_string_field("icon", application->icon);
  fputs(",", stdout);
  print_string_field("mode", application->runtime_mode);
  fputs(",", stdout);
  print_string_field("primary_extension", application->primary_extension);
  fputs(",", stdout);
  print_string_field("primary_mime_type", application->primary_mime_type);
  fputs("}", stdout);
}

static void
print_ai_diagnostic_context_section(const XnixRuntimeAIDiagnosticContextSection *section)
{
  fputs("{", stdout);
  print_string_field("id", section->id);
  fputs(",", stdout);
  print_string_field("source", section->source);
  fputs(",", stdout);
  print_string_field("summary", section->summary);
  fputs(",", stdout);
  fputs("\"facts\":{", stdout);
  print_string_field(section->primary_fact_key, section->primary_fact_value);
  fputs("}", stdout);
  fputs("}", stdout);
}

static void
print_ai_diagnostic_context_sections(const XnixRuntimeAIDiagnosticInput *input)
{
  fputc('[', stdout);
  for (size_t index = 0; index < input->context_section_count; index++) {
    if (index > 0) {
      fputs(",", stdout);
    }
    print_ai_diagnostic_context_section(&input->context_sections[index]);
  }
  fputc(']', stdout);
}

static void
print_ai_diagnostic_signal(const XnixRuntimeAIDiagnosticSignal *signal)
{
  fputs("{", stdout);
  print_string_field("id", signal->id);
  fputs(",", stdout);
  print_string_field("severity", signal->severity);
  fputs(",", stdout);
  print_string_field("source", signal->source);
  fputs(",", stdout);
  print_string_field("summary", signal->summary);
  fputs(",", stdout);
  print_string_field("recommended_next_action", signal->recommended_next_action);
  fputs("}", stdout);
}

static void
print_ai_diagnostic_signals(const XnixRuntimeAIDiagnosticInput *input)
{
  fputc('[', stdout);
  for (size_t index = 0; index < input->diagnostic_signal_count; index++) {
    if (index > 0) {
      fputs(",", stdout);
    }
    print_ai_diagnostic_signal(&input->diagnostic_signals[index]);
  }
  fputc(']', stdout);
}

static void
print_ai_diagnostic_privacy_boundaries(const XnixRuntimeAIDiagnosticPrivacyBoundaries *privacy)
{
  fputs("{", stdout);
  fputs("\"user_documents_included\":", stdout);
  print_bool(privacy->user_documents_included);
  fputs(",", stdout);
  fputs("\"host_paths_included\":", stdout);
  print_bool(privacy->host_paths_included);
  fputs(",", stdout);
  fputs("\"raw_backend_logs_included\":", stdout);
  print_bool(privacy->raw_backend_logs_included);
  fputs(",", stdout);
  fputs("\"secrets_included\":", stdout);
  print_bool(privacy->secrets_included);
  fputs(",", stdout);
  fputs("\"network_calls_allowed\":", stdout);
  print_bool(privacy->network_calls_allowed);
  fputs(",", stdout);
  fputs("\"requires_user_approval_for_sensitive_actions\":", stdout);
  print_bool(privacy->requires_user_approval_for_sensitive_actions);
  fputs("}", stdout);
}

static void
print_ai_task_list(const char *field_name, const char *const *tasks, size_t task_count)
{
  printf("\"%s\":[", field_name);
  for (size_t index = 0; index < task_count; index++) {
    if (index > 0) {
      fputs(",", stdout);
    }
    print_string(tasks[index]);
  }
  fputs("]", stdout);
}

static void
print_ai_diagnostic_input_record(const XnixRuntimeAIDiagnosticInput *input)
{
  fputs("{", stdout);
  printf("\"version\":\"%s\",", xnix_runtime_version());
  print_string_field("input_type", input->input_type);
  fputs(",", stdout);
  fputs("\"application\":", stdout);
  print_ai_diagnostic_application(input->application);
  fputs(",", stdout);
  print_string_field("issue", input->issue);
  fputs(",", stdout);
  print_string_field("test_type", input->test_type);
  fputs(",", stdout);
  fputs("\"runtime_owned\":", stdout);
  print_bool(input->runtime_owned);
  fputs(",", stdout);
  fputs("\"c_runtime_backed\":", stdout);
  print_bool(input->c_runtime_backed);
  fputs(",", stdout);
  fputs("\"kde_policy_owner\":", stdout);
  print_bool(input->kde_policy_owner);
  fputs(",", stdout);
  fputs("\"ai_provider_called\":", stdout);
  print_bool(input->ai_provider_called);
  fputs(",", stdout);
  fputs("\"network_required\":", stdout);
  print_bool(input->network_required);
  fputs(",", stdout);
  fputs("\"safe_for_ai_diagnostics\":", stdout);
  print_bool(input->safe_for_ai_diagnostics);
  fputs(",", stdout);
  fputs("\"context_sections\":", stdout);
  print_ai_diagnostic_context_sections(input);
  fputs(",", stdout);
  fputs("\"diagnostic_signals\":", stdout);
  print_ai_diagnostic_signals(input);
  fputs(",", stdout);
  fputs("\"privacy_boundaries\":", stdout);
  print_ai_diagnostic_privacy_boundaries(&input->privacy_boundaries);
  fputs(",", stdout);
  print_ai_task_list("allowed_ai_tasks", input->allowed_ai_tasks, input->allowed_ai_task_count);
  fputs(",", stdout);
  print_ai_task_list("blocked_ai_tasks", input->blocked_ai_tasks, input->blocked_ai_task_count);
  fputs(",", stdout);
  fputs("\"host_root_modified\":", stdout);
  print_bool(input->host_root_modified);
  fputs(",", stdout);
  fputs("\"backend_details_exposed\":", stdout);
  print_bool(input->backend_details_exposed);
  fputs(",", stdout);
  print_string_field("desktop_safe_summary", input->summary);
  fputs("}", stdout);
}

static int
print_ai_diagnostic_input_for_application(const char *application_id, const char *issue, const char *test_type)
{
  XnixRuntimeAIDiagnosticInput input;

  if (!xnix_runtime_ai_diagnostic_input(application_id, issue, test_type, &input)) {
    fprintf(stderr, "xnix-runtime-core: unknown application or AI diagnostic input catalog\n");
    return 64;
  }

  print_ai_diagnostic_input_record(&input);
  fputs("\n", stdout);
  return 0;
}
