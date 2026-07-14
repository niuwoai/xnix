static const XnixRuntimeCompatibilityReviewFlowStep compatibility_review_flow_steps[] = {
  {
    .id = "settings-change-review",
    .label = "Review settings change",
    .runtime_method = "GetCompatibilitySettingsChangePlan",
    .status = "required",
    .summary = "KDE presents the requested user-facing settings change before Runtime persistence.",
    .user_visible = true,
    .user_action_required = true,
    .runtime_gate_required = true,
    .blocks_apply = true,
  },
  {
    .id = "permission-review",
    .label = "Review permissions",
    .runtime_method = "GetCompatibilityPermissionReviewPlan",
    .status = "required",
    .summary = "KDE shows the grouped Runtime permission review before any permission grant.",
    .user_visible = true,
    .user_action_required = true,
    .runtime_gate_required = true,
    .blocks_apply = true,
  },
  {
    .id = "portal-request-review",
    .label = "Review Portal request",
    .runtime_method = "GetPortalRequestPlan",
    .status = "required",
    .summary = "Runtime maps sensitive desktop access to an XDG Desktop Portal request plan.",
    .user_visible = true,
    .user_action_required = true,
    .runtime_gate_required = true,
    .blocks_apply = true,
  },
  {
    .id = "runtime-write-gate",
    .label = "Check Runtime write gate",
    .runtime_method = "GetRuntimeWriteGate",
    .status = "blocked",
    .summary = "Runtime write gates still block persistence, launch, snapshots, and restore.",
    .user_visible = true,
    .user_action_required = false,
    .runtime_gate_required = true,
    .blocks_apply = true,
  },
  {
    .id = "review-receipt",
    .label = "Record review intent",
    .runtime_method = "GetCompatibilityActionReviewReceipt",
    .status = "pending",
    .summary = "Review intent can be recorded, but it cannot approve execution by itself.",
    .user_visible = true,
    .user_action_required = true,
    .runtime_gate_required = true,
    .blocks_apply = true,
  },
};

bool
xnix_runtime_compatibility_review_flow_plan(
  const char *application_id,
  const char *section_id,
  const char *field_id,
  const char *requested_value,
  const char *operation,
  XnixRuntimeCompatibilityReviewFlowPlan *plan
)
{
  const XnixRuntimeApplication *application;
  const XnixRuntimePortalPolicy *policy;

  if (!application_id || !section_id || !field_id || !requested_value || !operation || !plan) {
    return false;
  }

  application = xnix_runtime_find_application(application_id);
  if (!application) {
    return false;
  }

  policy = xnix_runtime_find_portal_policy(operation);
  if (!policy) {
    return false;
  }

  memset(plan, 0, sizeof(*plan));
  plan->application = application;
  plan->plan_type = "compatibility-review-flow-plan";
  plan->runtime_method = "GetCompatibilityReviewFlowPlan";
  plan->review_state = "planned";
  plan->section_id = section_id;
  plan->field_id = field_id;
  plan->requested_value = requested_value;
  plan->operation = operation;
  plan->step_count = sizeof(compatibility_review_flow_steps) / sizeof(compatibility_review_flow_steps[0]);
  for (size_t index = 0; index < plan->step_count; index++) {
    plan->steps[index] = compatibility_review_flow_steps[index];
    if (strcmp(plan->steps[index].status, "required") == 0) {
      plan->required_review_count++;
    } else if (strcmp(plan->steps[index].status, "blocked") == 0) {
      plan->blocked_step_count++;
    } else if (strcmp(plan->steps[index].status, "pending") == 0) {
      plan->pending_step_count++;
    }
  }
  plan->runtime_owned = true;
  plan->c_runtime_backed = true;
  plan->kde_policy_owner = false;
  plan->user_confirmation_required = true;
  plan->portal_policy_review_required = policy->portal_required;
  plan->settings_change_planned = true;
  plan->permission_review_planned = true;
  plan->portal_request_planned = true;
  plan->review_receipt_required = true;
  plan->apply_enabled = false;
  plan->request_object_created = false;
  plan->permission_granted = false;
  plan->settings_persisted = false;
  plan->execution_started = false;
  plan->host_root_modified = false;
  plan->backend_details_exposed = false;
  plan->summary = "KDE can display the full Runtime-owned review flow, but no permission, setting, or execution change is applied.";

  return true;
}
