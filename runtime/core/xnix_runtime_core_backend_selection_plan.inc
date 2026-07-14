static void
backend_selection_add_required_review(
  XnixRuntimeBackendSelectionPlan *plan,
  const char *review
)
{
  plan->required_reviews[plan->required_review_count] = review;
  plan->required_review_count++;
}

static void
backend_selection_add_blocked_action(
  XnixRuntimeBackendSelectionPlan *plan,
  const char *action
)
{
  plan->blocked_actions[plan->blocked_action_count] = action;
  plan->blocked_action_count++;
}

static void
backend_selection_populate_candidate(
  XnixRuntimeBackendSelectionCandidate *candidate,
  const XnixRuntimeBackendCapabilityProfile *profile,
  const XnixRuntimeBackendBindingPolicy *binding_policy,
  const char *recommended_profile_id
)
{
  memset(candidate, 0, sizeof(*candidate));
  candidate->id = profile->id;
  candidate->label = profile->label;
  candidate->kind = profile->kind;
  candidate->selection_state = strcmp(profile->id, recommended_profile_id) == 0 ?
    "recommended" :
    "available-after-review";
  candidate->recommended = strcmp(profile->id, recommended_profile_id) == 0;
  candidate->ready = false;
  candidate->blocked = true;
  candidate->ready_capability_count = profile->ready_count;
  candidate->pending_capability_count = profile->pending_count;
  for (size_t index = 0; index < binding_policy->required_preflight_count; index++) {
    candidate->required_preflight[index] = binding_policy->required_preflight[index].id;
    candidate->required_preflight_count++;
  }
  candidate->selection_committed = false;
  candidate->environment_created = false;
  candidate->backend_process_started = false;
  candidate->backend_details_exposed = false;
  candidate->summary = candidate->recommended ?
    "This compatibility profile is recommended by the Runtime but not committed." :
    "This compatibility profile is available only after Runtime review.";
}

bool
xnix_runtime_backend_selection_plan(
  const char *application_id,
  XnixRuntimeBackendSelectionPlan *plan
)
{
  XnixRuntimeCompatibilityRunPlan run_plan;
  XnixRuntimeBackendCapabilityMatrix matrix;
  const XnixRuntimeBackendBindingPolicy *binding_policy;
  const char *recommended_profile_id;

  if (application_id == NULL || plan == NULL) {
    return false;
  }

  if (!xnix_runtime_compatibility_run_plan(application_id, &run_plan)) {
    return false;
  }

  if (!xnix_runtime_backend_capability_matrix(&matrix)) {
    return false;
  }

  binding_policy = xnix_runtime_find_backend_binding_policy(application_id);
  if (!binding_policy) {
    return false;
  }

  recommended_profile_id = strcmp(run_plan.engine->id, "isolated-compatibility-engine") == 0 ?
    "isolated-compatibility" :
    "local-compatibility";

  memset(plan, 0, sizeof(*plan));
  plan->application = run_plan.application;
  plan->engine = run_plan.engine;
  plan->plan_type = "compatibility-backend-selection-plan";
  plan->runtime_method = "GetBackendSelectionPlan";
  plan->selected_strategy = run_plan.engine->id;
  plan->recommended_profile_id = recommended_profile_id;
  for (size_t index = 0; index < matrix.profile_count; index++) {
    backend_selection_populate_candidate(
      &plan->candidates[index],
      &matrix.profiles[index],
      binding_policy,
      recommended_profile_id
    );
    plan->candidate_count++;
    if (plan->candidates[index].ready) {
      plan->ready_candidate_count++;
    }
    if (plan->candidates[index].blocked) {
      plan->blocked_candidate_count++;
    }
  }
  backend_selection_add_required_review(plan, "backend-capability-review");
  backend_selection_add_required_review(plan, "backend-binding-review");
  backend_selection_add_required_review(plan, "application-state-root-review");
  backend_selection_add_required_review(plan, "portal-policy-review");
  backend_selection_add_required_review(plan, "snapshot-baseline-review");
  backend_selection_add_blocked_action(plan, "commit backend selection from KDE");
  backend_selection_add_blocked_action(plan, "start compatibility backend after selecting a profile");
  backend_selection_add_blocked_action(plan, "create profile environment before Runtime preflight");
  backend_selection_add_blocked_action(plan, "expose backend implementation details to the desktop shell");
  plan->runtime_owned = true;
  plan->c_runtime_backed = true;
  plan->kde_policy_owner = false;
  plan->selection_committed = false;
  plan->selection_change_enabled = false;
  plan->backend_launch_enabled = false;
  plan->capability_activation_enabled = false;
  plan->environment_created = false;
  plan->request_object_created = false;
  plan->state_root_created = false;
  plan->snapshot_created = false;
  plan->host_root_modified = false;
  plan->privileged_container_required = false;
  plan->backend_details_exposed = false;
  plan->summary = "Runtime can explain the recommended compatibility profile, but no backend selection is committed yet.";

  return true;
}
