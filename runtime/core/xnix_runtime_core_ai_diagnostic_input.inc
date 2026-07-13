static bool
ai_diagnostic_test_type_allowed(const char *test_type)
{
  return strcmp(test_type, "preflight") == 0 ||
         strcmp(test_type, "smoke") == 0 ||
         strcmp(test_type, "repair-readiness") == 0;
}

static XnixRuntimeAIDiagnosticContextSection
ai_diagnostic_context_section(
  const char *id,
  const char *source,
  const char *summary,
  const char *primary_fact_key,
  const char *primary_fact_value
)
{
  XnixRuntimeAIDiagnosticContextSection section = {
    .id = id,
    .source = source,
    .summary = summary,
    .primary_fact_key = primary_fact_key,
    .primary_fact_value = primary_fact_value,
  };
  return section;
}

static XnixRuntimeAIDiagnosticSignal
ai_diagnostic_signal(
  const char *id,
  const char *severity,
  const char *source,
  const char *summary,
  const char *recommended_next_action
)
{
  XnixRuntimeAIDiagnosticSignal signal = {
    .id = id,
    .severity = severity,
    .source = source,
    .summary = summary,
    .recommended_next_action = recommended_next_action,
  };
  return signal;
}

bool
xnix_runtime_ai_diagnostic_input(
  const char *application_id,
  const char *issue,
  const char *test_type,
  XnixRuntimeAIDiagnosticInput *input
)
{
  if (input == NULL || issue == NULL || test_type == NULL) {
    return false;
  }

  if (!ai_diagnostic_test_type_allowed(test_type)) {
    return false;
  }

  const XnixRuntimeApplication *application = xnix_runtime_find_application(application_id);
  if (application == NULL) {
    return false;
  }

  input->application = application;
  input->input_type = "ai-diagnostic-input";
  input->issue = issue;
  input->test_type = test_type;
  input->context_sections[0] = ai_diagnostic_context_section(
    "recipe",
    "application-recipe",
    "Application recipe metadata is available.",
    "application_id",
    application->id
  );
  input->context_sections[1] = ai_diagnostic_context_section(
    "run-plan",
    "compatibility-run",
    "Managed compatibility launch binding is not ready yet.",
    "backend_binding_ready",
    "false"
  );
  input->context_sections[2] = ai_diagnostic_context_section(
    "test-result",
    "compatibility-test-result",
    "Compatibility test work is pending and safe for AI summarization.",
    "test_type",
    test_type
  );
  input->context_sections[3] = ai_diagnostic_context_section(
    "repair-plan",
    "compatibility-repair",
    "Repair execution remains blocked until Runtime approval and restore-point gates pass.",
    "issue",
    issue
  );
  input->context_section_count = 4;
  input->diagnostic_signals[0] = ai_diagnostic_signal(
    "pending-runtime-launch-binding",
    "medium",
    "run-plan",
    "Managed compatibility launch binding is not ready yet.",
    "Keep automated smoke execution pending until a launch backend is bound."
  );
  input->diagnostic_signals[1] = ai_diagnostic_signal(
    "pending-test-work",
    "low",
    "test-result",
    "Compatibility test steps are pending.",
    "Show pending Runtime-controlled preflight work in the Compatibility Center."
  );
  input->diagnostic_signals[2] = ai_diagnostic_signal(
    "snapshot-before-risky-change",
    "low",
    "repair-plan",
    "A Runtime restore point is required before risky compatibility changes.",
    "Prepare a restore point before repair execution."
  );
  input->diagnostic_signal_count = 3;
  input->privacy_boundaries.user_documents_included = false;
  input->privacy_boundaries.host_paths_included = false;
  input->privacy_boundaries.raw_backend_logs_included = false;
  input->privacy_boundaries.secrets_included = false;
  input->privacy_boundaries.network_calls_allowed = false;
  input->privacy_boundaries.requires_user_approval_for_sensitive_actions = true;
  input->allowed_ai_tasks[0] = "summarize compatibility status";
  input->allowed_ai_tasks[1] = "explain pending Runtime work";
  input->allowed_ai_tasks[2] = "suggest safe next diagnostic steps";
  input->allowed_ai_tasks[3] = "prepare user-facing Compatibility Center text";
  input->allowed_ai_task_count = 4;
  input->blocked_ai_tasks[0] = "start a compatibility backend";
  input->blocked_ai_tasks[1] = "read user documents";
  input->blocked_ai_tasks[2] = "change desktop permissions";
  input->blocked_ai_tasks[3] = "execute repair actions without Runtime approval";
  input->blocked_ai_task_count = 4;
  input->runtime_owned = true;
  input->c_runtime_backed = true;
  input->kde_policy_owner = false;
  input->ai_provider_called = false;
  input->network_required = false;
  input->safe_for_ai_diagnostics = true;
  input->backend_details_exposed = false;
  input->host_root_modified = false;
  input->summary = "AI diagnostics can explain current compatibility status using Runtime-safe metadata only.";

  return true;
}
