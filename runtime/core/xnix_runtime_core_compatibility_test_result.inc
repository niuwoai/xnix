static const char *
compatibility_test_step_result_for_status(const char *status)
{
  if (strcmp(status, "pass") == 0) {
    return "passed";
  }

  if (strcmp(status, "blocked") == 0) {
    return "blocked";
  }

  return "not-run";
}

static const char *
compatibility_test_step_evidence(const char *step_id, const char *summary)
{
  if (strcmp(step_id, "recipe-validation") == 0) {
    return "Recipe metadata is loaded and valid.";
  }

  if (strcmp(step_id, "portal-preflight") == 0) {
    return "XDG Desktop Portal request model is available.";
  }

  if (strcmp(step_id, "snapshot-preflight") == 0) {
    return "Runtime snapshot plan is available and restore-capable.";
  }

  if (strcmp(step_id, "runtime-launch-binding") == 0) {
    return "Runtime launch backend binding is still pending.";
  }

  return summary;
}

static const char *
compatibility_test_step_next_action(const char *step_id, const char *status)
{
  if (strcmp(status, "pass") == 0) {
    return "No action required.";
  }

  if (strcmp(status, "blocked") == 0) {
    return "Request user approval before continuing.";
  }

  if (strcmp(step_id, "portal-preflight") == 0) {
    return "Wait for a user-mediated Portal grant during execution.";
  }

  if (strcmp(step_id, "snapshot-preflight") == 0) {
    return "Create a Runtime restore point before risky compatibility changes.";
  }

  if (strcmp(step_id, "runtime-launch-binding") == 0) {
    return "Bind a managed compatibility launch backend before automated smoke execution.";
  }

  return "Run the pending compatibility test step.";
}

bool
xnix_runtime_compatibility_test_result(
  const char *application_id,
  const char *test_type,
  XnixRuntimeCompatibilityTestResult *result
)
{
  XnixRuntimeCompatibilityTestPlan plan;

  if (application_id == NULL || test_type == NULL || result == NULL) {
    return false;
  }

  if (!xnix_runtime_compatibility_test_plan(application_id, test_type, &plan)) {
    return false;
  }

  memset(result, 0, sizeof(*result));
  result->application = plan.application;
  result->test_type = plan.test_type;
  result->result_source = "runtime-model";
  result->runtime_owned = true;
  result->c_runtime_backed = true;
  result->kde_policy_owner = false;
  result->compatibility_center_card = true;
  result->diagnostics_record = true;
  result->safe_for_ai_diagnostics = true;
  result->test_executed = false;
  result->host_root_modified = false;
  result->backend_details_exposed = false;
  result->notification_event = "mode-changed";
  result->repair_plan_issue = "engine-binding-pending";

  for (size_t index = 0; index < plan.step_count; index++) {
    const XnixRuntimeCompatibilityTestStep *step = &plan.steps[index];
    XnixRuntimeCompatibilityTestStepResult *step_result = &result->step_results[index];

    step_result->id = step->id;
    step_result->title = step->title;
    step_result->status = step->status;
    step_result->result = compatibility_test_step_result_for_status(step->status);
    step_result->evidence = compatibility_test_step_evidence(step->id, step->summary);
    step_result->next_action = compatibility_test_step_next_action(step->id, step->status);

    result->counts.total++;
    if (strcmp(step->status, "pass") == 0) {
      result->counts.passed++;
    } else if (strcmp(step->status, "blocked") == 0) {
      result->counts.blocked++;
    } else {
      result->counts.pending++;
    }
  }
  result->step_result_count = plan.step_count;

  if (result->counts.blocked > 0) {
    result->execution_state = "blocked";
    result->overall_status = "blocked";
    result->notification_event = "approval-required";
    result->summary = "Compatibility test execution is blocked and needs user action.";
  } else if (result->counts.pending > 0) {
    result->execution_state = "waiting-for-runtime";
    result->overall_status = "pending";
    result->summary = "Compatibility test execution is waiting for Runtime-controlled preflight work.";
  } else {
    result->execution_state = "complete";
    result->overall_status = "pass";
    result->repair_plan_issue = NULL;
    result->summary = "Compatibility test execution completed successfully.";
  }

  return true;
}
