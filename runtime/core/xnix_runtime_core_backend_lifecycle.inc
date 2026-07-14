static void
backend_lifecycle_add_stage(
  XnixRuntimeBackendLifecycle *lifecycle,
  const char *id,
  const char *status,
  const char *summary
)
{
  XnixRuntimeBackendLifecycleStage *stage = &lifecycle->stages[lifecycle->stage_count];

  stage->id = id;
  stage->status = status;
  stage->summary = summary;
  lifecycle->stage_count++;
}

static void
backend_lifecycle_add_blocked_action(
  XnixRuntimeBackendLifecycle *lifecycle,
  const char *action
)
{
  lifecycle->blocked_actions[lifecycle->blocked_action_count] = action;
  lifecycle->blocked_action_count++;
}

bool
xnix_runtime_backend_lifecycle(
  const char *application_id,
  XnixRuntimeBackendLifecycle *lifecycle
)
{
  XnixRuntimeCompatibilityRunPlan run_plan;
  XnixRuntimeWriteGate launch_gate;

  if (application_id == NULL || lifecycle == NULL) {
    return false;
  }

  if (!xnix_runtime_compatibility_run_plan(application_id, &run_plan)) {
    return false;
  }

  if (!xnix_runtime_write_gate("Launch", &launch_gate)) {
    return false;
  }

  memset(lifecycle, 0, sizeof(*lifecycle));
  lifecycle->application = run_plan.application;
  lifecycle->engine = run_plan.engine;
  lifecycle->lifecycle_type = "compatibility-backend-lifecycle";
  lifecycle->runtime_method = "GetBackendLifecycle";
  lifecycle->selected_strategy = run_plan.engine->id;
  lifecycle->lifecycle_state = "blocked";
  lifecycle->overall_status = "not-ready";
  lifecycle->runtime_owned = true;
  lifecycle->c_runtime_backed = true;
  lifecycle->kde_policy_owner = false;
  lifecycle->backend_binding_ready = run_plan.backend_binding_ready;
  lifecycle->launch_enabled = run_plan.launch_enabled && launch_gate.dispatch_enabled;
  lifecycle->execution_request_created = false;
  lifecycle->backend_process_started = false;
  lifecycle->local_backend_started = false;
  lifecycle->isolated_backend_started = false;
  lifecycle->state_root_ready = false;
  lifecycle->portal_review_required = run_plan.portal_policy_required;
  lifecycle->snapshot_required = run_plan.snapshot_before_risky_change;
  lifecycle->host_root_modified = false;
  lifecycle->network_required = run_plan.network_required;
  lifecycle->backend_details_exposed = false;

  backend_lifecycle_add_stage(
    lifecycle,
    "recipe-loaded",
    "pass",
    "Recipe metadata is loaded and mapped to a Runtime-owned backend lifecycle."
  );
  backend_lifecycle_add_stage(
    lifecycle,
    "state-root-ready",
    "pending",
    "A Runtime-owned application state root must exist before backend lifecycle activation."
  );
  backend_lifecycle_add_stage(
    lifecycle,
    "backend-binding-ready",
    run_plan.backend_binding_ready ? "pass" : "pending",
    "A managed backend binding must be ready before any lifecycle start request."
  );
  backend_lifecycle_add_stage(
    lifecycle,
    "portal-and-snapshot-review",
    "required",
    "Portal access and restore-point policy must be reviewed before backend lifecycle activation."
  );
  backend_lifecycle_add_stage(
    lifecycle,
    "runtime-launch-write-gate",
    launch_gate.dispatch_enabled ? "pass" : "blocked",
    "Runtime Launch write dispatch remains disabled until production backend ownership is ready."
  );

  backend_lifecycle_add_blocked_action(lifecycle, "start local compatibility backend from KDE");
  backend_lifecycle_add_blocked_action(lifecycle, "start isolated compatibility backend from KDE");
  backend_lifecycle_add_blocked_action(lifecycle, "create backend process before Runtime lifecycle gates pass");
  backend_lifecycle_add_blocked_action(lifecycle, "expose backend command to desktop shell");
  backend_lifecycle_add_blocked_action(lifecycle, "mutate host root during lifecycle planning");

  lifecycle->summary = "Runtime backend lifecycle is modeled but blocked; KDE may display state and must not start compatibility backends.";

  return true;
}
