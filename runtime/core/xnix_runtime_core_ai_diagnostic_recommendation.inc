static XnixRuntimeAIDiagnosticRecommendationItem
ai_diagnostic_recommendation_item(
  const char *id,
  const char *priority,
  const char *source_signal,
  const char *title,
  const char *summary,
  const char *next_runtime_action,
  bool requires_user_approval
)
{
  XnixRuntimeAIDiagnosticRecommendationItem item = {
    .id = id,
    .priority = priority,
    .source_signal = source_signal,
    .title = title,
    .summary = summary,
    .next_runtime_action = next_runtime_action,
    .user_visible = true,
    .auto_execute = false,
    .requires_user_approval = requires_user_approval,
  };
  return item;
}

bool
xnix_runtime_ai_diagnostic_recommendation(
  const char *application_id,
  const char *issue,
  const char *test_type,
  XnixRuntimeAIDiagnosticRecommendation *recommendation
)
{
  if (recommendation == NULL) {
    return false;
  }

  if (!xnix_runtime_ai_diagnostic_input(
        application_id,
        issue,
        test_type,
        &recommendation->diagnostic_input
      )) {
    return false;
  }

  recommendation->recommendation_type = "ai-diagnostic-recommendation";
  recommendation->recommendations[0] = ai_diagnostic_recommendation_item(
    "explain-pending-runtime-work",
    "high",
    "pending-runtime-launch-binding",
    "Explain pending Runtime launch binding",
    "Tell the user that automated compatibility execution is waiting for a managed Runtime launch backend.",
    "Keep smoke execution pending until backend binding is available.",
    false
  );
  recommendation->recommendations[1] = ai_diagnostic_recommendation_item(
    "prepare-safe-restore-point",
    "medium",
    "snapshot-before-risky-change",
    "Prepare a restore point before repair",
    "Ask the Runtime to create an application-scoped restore point before risky compatibility changes.",
    "Create a Runtime restore point only after approval.",
    true
  );
  recommendation->recommendations[2] = ai_diagnostic_recommendation_item(
    "surface-test-progress",
    "medium",
    "pending-test-work",
    "Show compatibility test progress",
    "Display pending preflight work in the Compatibility Center without exposing backend details.",
    "Update the Compatibility Center card with pending test status.",
    false
  );
  recommendation->recommendation_count = 3;
  recommendation->approval_required_actions[0].id = "prepare-safe-restore-point";
  recommendation->approval_required_actions[0].title = "Prepare a restore point before repair";
  recommendation->approval_required_actions[0].reason = "Ask the Runtime to create an application-scoped restore point before risky compatibility changes.";
  recommendation->approval_required_actions[0].approval_surface = "Compatibility Center";
  recommendation->approval_required_action_count = 1;
  recommendation->blocked_actions[0] = "start a compatibility backend";
  recommendation->blocked_actions[1] = "read user documents";
  recommendation->blocked_actions[2] = "change desktop permissions";
  recommendation->blocked_actions[3] = "execute repair actions without Runtime approval";
  recommendation->blocked_action_count = 4;
  recommendation->runtime_owned = true;
  recommendation->c_runtime_backed = true;
  recommendation->kde_policy_owner = false;
  recommendation->ai_provider_called = false;
  recommendation->network_required = false;
  recommendation->safe_for_ai_diagnostics = true;
  recommendation->auto_execution_allowed = false;
  recommendation->backend_details_exposed = false;
  recommendation->host_root_modified = false;
  recommendation->summary = "AI diagnostic recommendations are ready for review without executing compatibility changes.";

  return true;
}
