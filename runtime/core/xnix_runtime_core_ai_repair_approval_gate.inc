static XnixRuntimeAIRepairRequiredGate
ai_repair_required_gate(const char *id, const char *status, const char *summary)
{
  XnixRuntimeAIRepairRequiredGate gate = {
    .id = id,
    .status = status,
    .summary = summary,
  };
  return gate;
}

bool
xnix_runtime_ai_repair_approval_gate(
  const char *application_id,
  const char *issue,
  const char *test_type,
  XnixRuntimeAIRepairApprovalGate *gate
)
{
  const XnixRuntimeAIDiagnosticRecommendation *recommendation = NULL;

  if (gate == NULL) {
    return false;
  }

  if (!xnix_runtime_ai_diagnostic_recommendation(
        application_id,
        issue,
        test_type,
        &gate->recommendation
      )) {
    return false;
  }

  recommendation = &gate->recommendation;
  gate->gate_type = "ai-repair-approval-gate";
  gate->gate_decision =
    recommendation->approval_required_action_count == 0 ? "review-only" : "blocked-until-approval";
  gate->approval_surface = "Compatibility Center";
  gate->required_gates[0] = ai_repair_required_gate(
    "compatibility-center-review",
    "required",
    "A user-visible Compatibility Center review is required before any repair action."
  );
  gate->required_gates[1] = ai_repair_required_gate(
    "runtime-approval-token",
    "missing",
    "No Runtime approval token is present for repair execution."
  );
  gate->required_gates[2] = ai_repair_required_gate(
    "restore-point-preflight",
    "required",
    "A Runtime restore point must be prepared before risky compatibility repair."
  );
  gate->required_gate_count = 3;
  gate->blocked_actions[0] = "start a compatibility backend";
  gate->blocked_actions[1] = "read user documents";
  gate->blocked_actions[2] = "change desktop permissions";
  gate->blocked_actions[3] = "execute repair actions without Runtime approval";
  gate->blocked_actions[4] = "execute repair without approval";
  gate->blocked_actions[5] = "create restore point without approval";
  gate->blocked_actions[6] = "change compatibility mode without Runtime gate";
  gate->blocked_action_count = 7;
  gate->runtime_owned = true;
  gate->c_runtime_backed = true;
  gate->kde_policy_owner = false;
  gate->ai_provider_called = false;
  gate->network_required = false;
  gate->safe_for_ai_diagnostics = true;
  gate->repair_execution_requested = false;
  gate->repair_executed = false;
  gate->auto_execution_allowed = false;
  gate->backend_details_exposed = false;
  gate->host_root_modified = false;
  gate->summary =
    "AI repair recommendations are blocked from execution until Runtime approval gates pass.";

  return true;
}
