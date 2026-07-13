static bool
compatibility_action_review_decision_valid(const char *decision)
{
  return strcmp(decision, "reviewed") == 0 ||
         strcmp(decision, "approved") == 0 ||
         strcmp(decision, "deferred") == 0 ||
         strcmp(decision, "rejected") == 0;
}

static const char *
compatibility_action_review_next_step(
  const XnixRuntimeCompatibilityAction *action,
  const char *decision
)
{
  if (strcmp(decision, "rejected") == 0) {
    return "Runtime records the rejection and keeps the queued action non-executing.";
  }

  if (strcmp(decision, "deferred") == 0) {
    return "Runtime keeps the queued action deferred until the user reviews it again.";
  }

  if (strcmp(action->runtime_gate, "ai-repair-approval") == 0) {
    return "Runtime records the review intent, then waits for ai-repair-approval before any execution.";
  }

  if (strcmp(action->runtime_gate, "runtime-settings-persistence") == 0) {
    return "Runtime records the review intent, then waits for runtime-settings-persistence before any execution.";
  }

  if (strcmp(action->runtime_gate, "portal-policy-review") == 0) {
    return "Runtime records the review intent, then waits for portal-policy-review before any execution.";
  }

  if (strcmp(action->runtime_gate, "install-plan-readiness") == 0) {
    return "Runtime records the review intent, then waits for install-plan-readiness before any execution.";
  }

  return "Runtime records the review intent, then waits for live-dbus-owner before any execution.";
}

bool
xnix_runtime_compatibility_action_review_receipt(
  const char *application_id,
  const char *action_id,
  const char *decision,
  XnixRuntimeCompatibilityActionReviewReceipt *receipt
)
{
  XnixRuntimeCompatibilityActionQueue queue;
  const XnixRuntimeCompatibilityAction *action = NULL;

  if (application_id == NULL || action_id == NULL || decision == NULL ||
      receipt == NULL) {
    return false;
  }

  if (!compatibility_action_review_decision_valid(decision)) {
    return false;
  }

  if (!xnix_runtime_compatibility_action_queue(application_id, &queue)) {
    return false;
  }

  for (size_t index = 0; index < queue.action_count; index++) {
    if (strcmp(queue.actions[index].id, action_id) == 0) {
      action = &queue.actions[index];
      break;
    }
  }

  if (action == NULL) {
    return false;
  }

  snprintf(
    receipt->receipt_id,
    sizeof(receipt->receipt_id),
    "compat-review-%s-%s-%s",
    application_id,
    action_id,
    decision
  );
  receipt->application = queue.application;
  receipt->action = *action;
  receipt->decision = decision;
  receipt->next_step = compatibility_action_review_next_step(action, decision);
  receipt->blocked_actions[0] = "execute queued action from a review receipt";
  receipt->blocked_actions[1] = "treat KDE review intent as Runtime execution approval";
  receipt->blocked_actions[2] = "persist compatibility settings from a review receipt";
  receipt->blocked_actions[3] = "grant desktop resources from a review receipt";
  receipt->blocked_actions[4] = "mutate the host root from a review receipt";
  receipt->blocked_actions[5] = "expose backend implementation details in review receipts";
  receipt->blocked_action_count = 6;
  receipt->decision_recorded = true;
  receipt->runtime_owned = true;
  receipt->kde_policy_owner = false;
  receipt->execution_enabled = false;
  receipt->repair_execution_enabled = false;
  receipt->settings_persistence_enabled = false;
  receipt->resource_grant_created = false;
  receipt->host_root_modified = false;
  receipt->network_required = false;
  receipt->backend_details_exposed = false;
  receipt->summary = "Compatibility Center review intent is recorded, but Runtime execution gates still block action execution.";

  return true;
}
