static XnixRuntimeCompatibilityRepairAction
compatibility_repair_action(const char *id, const char *kind, const char *label)
{
  XnixRuntimeCompatibilityRepairAction action = {
    .id = id,
    .kind = kind,
    .label = label,
  };
  return action;
}

static bool
set_compatibility_repair_issue(
  const char *issue,
  XnixRuntimeCompatibilityRepairPlan *plan
)
{
  if (strcmp(issue, "engine-binding-pending") == 0) {
    plan->severity = "warning";
    plan->automatic_allowed = false;
    plan->user_approval_required = true;
    plan->snapshot_required = true;
    plan->notification_event = "approval-required";
    plan->actions[0] = compatibility_repair_action(
      "open-diagnostics",
      "user-visible",
      "Open diagnostics"
    );
    plan->actions[1] = compatibility_repair_action(
      "prepare-engine-binding",
      "runtime-task",
      "Prepare compatibility engine binding"
    );
    plan->action_count = 2;
    plan->summary = "Compatibility engine setup is not ready yet.";
    return true;
  }

  if (strcmp(issue, "portal-approval-required") == 0) {
    plan->severity = "info";
    plan->automatic_allowed = false;
    plan->user_approval_required = true;
    plan->snapshot_required = false;
    plan->notification_event = "approval-required";
    plan->actions[0] = compatibility_repair_action(
      "request-portal-grant",
      "portal-request",
      "Request desktop permission"
    );
    plan->action_count = 1;
    plan->summary = "A desktop permission request needs user approval.";
    return true;
  }

  if (strcmp(issue, "recipe-trust-blocked") == 0) {
    plan->severity = "critical";
    plan->automatic_allowed = false;
    plan->user_approval_required = true;
    plan->snapshot_required = false;
    plan->notification_event = "approval-required";
    plan->actions[0] = compatibility_repair_action(
      "review-recipe-source",
      "user-visible",
      "Review recipe source"
    );
    plan->action_count = 1;
    plan->summary = "Recipe trust requirements are not satisfied.";
    return true;
  }

  if (strcmp(issue, "runtime-repair-applied") == 0) {
    plan->severity = "info";
    plan->automatic_allowed = true;
    plan->user_approval_required = false;
    plan->snapshot_required = true;
    plan->notification_event = "repair-applied";
    plan->actions[0] = compatibility_repair_action(
      "show-repair-record",
      "user-visible",
      "Show repair record"
    );
    plan->action_count = 1;
    plan->summary = "A safe compatibility repair was applied.";
    return true;
  }

  return false;
}

bool
xnix_runtime_compatibility_repair_plan(
  const char *application_id,
  const char *issue,
  XnixRuntimeCompatibilityRepairPlan *plan
)
{
  if (application_id == NULL || issue == NULL || plan == NULL) {
    return false;
  }

  if (xnix_runtime_find_application(application_id) == NULL) {
    return false;
  }

  memset(plan, 0, sizeof(*plan));
  plan->application_id = application_id;
  plan->issue = issue;
  if (!set_compatibility_repair_issue(issue, plan)) {
    return false;
  }

  if (plan->snapshot_required) {
    plan->snapshot_policy = xnix_runtime_find_snapshot_policy("before-repair");
    if (plan->snapshot_policy == NULL) {
      return false;
    }
  }

  plan->rollback_available = true;
  plan->runtime_owned = true;
  plan->c_runtime_backed = true;
  plan->kde_policy_owner = false;
  plan->repair_execution_requested = false;
  plan->repair_executed = false;
  plan->backend_launch_enabled = false;
  plan->network_required = false;
  plan->host_root_modified = false;
  plan->backend_details_exposed = false;

  return true;
}
