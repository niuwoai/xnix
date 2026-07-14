static void
backend_environment_add_profile(
  XnixRuntimeBackendEnvironmentPlan *plan,
  const char *id,
  const char *kind,
  const char *status,
  const char *summary,
  bool selected
)
{
  XnixRuntimeBackendEnvironmentProfile *profile = &plan->profiles[plan->profile_count];

  profile->id = id;
  profile->kind = kind;
  profile->status = status;
  profile->summary = summary;
  profile->selected = selected;
  profile->environment_created = false;
  profile->process_started = false;
  profile->host_storage_exposed = false;
  profile->backend_details_exposed = false;
  plan->profile_count++;
}

static void
backend_environment_add_required_review(
  XnixRuntimeBackendEnvironmentPlan *plan,
  const char *review
)
{
  plan->required_reviews[plan->required_review_count] = review;
  plan->required_review_count++;
}

static void
backend_environment_add_blocked_action(
  XnixRuntimeBackendEnvironmentPlan *plan,
  const char *action
)
{
  plan->blocked_actions[plan->blocked_action_count] = action;
  plan->blocked_action_count++;
}

bool
xnix_runtime_backend_environment_plan(
  const char *application_id,
  XnixRuntimeBackendEnvironmentPlan *plan
)
{
  XnixRuntimeCompatibilityRunPlan run_plan;
  XnixRuntimeWriteGate launch_gate;

  if (application_id == NULL || plan == NULL) {
    return false;
  }

  if (!xnix_runtime_compatibility_run_plan(application_id, &run_plan)) {
    return false;
  }

  if (!xnix_runtime_write_gate("Launch", &launch_gate)) {
    return false;
  }

  memset(plan, 0, sizeof(*plan));
  plan->application = run_plan.application;
  plan->engine = run_plan.engine;
  plan->plan_type = "compatibility-backend-environment-plan";
  plan->runtime_method = "GetBackendEnvironmentPlan";
  plan->selected_strategy = run_plan.engine->id;
  plan->environment_state = "planned";
  plan->runtime_owned = true;
  plan->c_runtime_backed = true;
  plan->kde_policy_owner = false;
  plan->local_environment_ready = false;
  plan->isolated_environment_ready = false;
  plan->environment_created = false;
  plan->backend_process_started = false;
  plan->host_storage_exposed = false;
  plan->clipboard_bridge_enabled = false;
  plan->print_bridge_enabled = false;
  plan->portal_review_required = run_plan.portal_policy_required;
  plan->snapshot_required = run_plan.snapshot_before_risky_change;
  plan->launch_enabled = run_plan.launch_enabled && launch_gate.dispatch_enabled;
  plan->host_root_modified = false;
  plan->network_required = run_plan.network_required;
  plan->backend_details_exposed = false;

  backend_environment_add_profile(
    plan,
    "local-compatibility-environment",
    "local",
    "blocked",
    "Local compatibility environment preparation waits for Runtime package, state, Portal, snapshot, and Launch gates.",
    strcmp(run_plan.engine->id, "local-compatibility-engine") == 0
  );
  backend_environment_add_profile(
    plan,
    "isolated-compatibility-environment",
    "isolated",
    "blocked",
    "Isolated compatibility environment preparation waits for Runtime package, state, Portal, snapshot, and Launch gates.",
    strcmp(run_plan.engine->id, "isolated-compatibility-engine") == 0
  );

  backend_environment_add_required_review(plan, "package-source-review");
  backend_environment_add_required_review(plan, "application-state-root-review");
  backend_environment_add_required_review(plan, "portal-policy-review");
  backend_environment_add_required_review(plan, "snapshot-baseline-review");

  backend_environment_add_blocked_action(plan, "create local compatibility environment from KDE");
  backend_environment_add_blocked_action(plan, "create isolated compatibility environment from KDE");
  backend_environment_add_blocked_action(plan, "start backend process before environment gates pass");
  backend_environment_add_blocked_action(plan, "bridge clipboard or print devices without Portal review");
  backend_environment_add_blocked_action(plan, "expose backend environment commands to desktop shell");

  plan->summary = "Backend environments are planned by the Runtime but not created; KDE may display readiness only.";

  return true;
}
