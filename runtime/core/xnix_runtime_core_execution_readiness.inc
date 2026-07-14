static const char *
execution_readiness_gate_status(bool ready, bool required)
{
  if (ready) {
    return "pass";
  }

  return required ? "required" : "pending";
}

static void
execution_readiness_add_gate(
  XnixRuntimeExecutionReadiness *readiness,
  const char *id,
  const char *status,
  const char *summary
)
{
  XnixRuntimeExecutionReadinessGate *gate = &readiness->gates[readiness->gate_count];

  gate->id = id;
  gate->status = status;
  gate->summary = summary;
  readiness->gate_count++;
}

static void
execution_readiness_add_blocked_action(
  XnixRuntimeExecutionReadiness *readiness,
  const char *action
)
{
  readiness->blocked_actions[readiness->blocked_action_count] = action;
  readiness->blocked_action_count++;
}

bool
xnix_runtime_execution_readiness(
  const char *application_id,
  XnixRuntimeExecutionReadiness *readiness
)
{
  XnixRuntimeCompatibilityRunPlan run_plan;
  XnixRuntimeCompatibilityTestResult test_result;
  XnixRuntimeWriteGate launch_gate;

  if (application_id == NULL || readiness == NULL) {
    return false;
  }

  if (!xnix_runtime_compatibility_run_plan(application_id, &run_plan)) {
    return false;
  }

  if (!xnix_runtime_compatibility_test_result(application_id, "preflight", &test_result)) {
    return false;
  }

  if (!xnix_runtime_write_gate("Launch", &launch_gate)) {
    return false;
  }

  memset(readiness, 0, sizeof(*readiness));
  readiness->application = run_plan.application;
  readiness->engine = run_plan.engine;
  readiness->readiness_type = "compatibility-execution-readiness";
  readiness->runtime_method = "GetExecutionReadiness";
  readiness->runtime_owned = true;
  readiness->c_runtime_backed = true;
  readiness->kde_policy_owner = false;
  readiness->compatibility_center_card = true;
  readiness->safe_for_ai_diagnostics = true;
  readiness->desktop_entry_launch_visible = true;
  readiness->launch_allowed = false;
  readiness->launch_enabled = run_plan.launch_enabled && launch_gate.dispatch_enabled;
  readiness->execution_request_created = false;
  readiness->backend_binding_ready = run_plan.backend_binding_ready;
  readiness->portal_policy_required = run_plan.portal_policy_required;
  readiness->snapshot_required = run_plan.snapshot_before_risky_change;
  readiness->user_action_required = true;
  readiness->host_root_modified = false;
  readiness->network_required = run_plan.network_required;
  readiness->backend_details_exposed = false;

  execution_readiness_add_gate(
    readiness,
    "recipe-validation",
    "pass",
    "Recipe metadata is loaded and can produce a desktop-safe execution plan."
  );
  execution_readiness_add_gate(
    readiness,
    "portal-policy-review",
    execution_readiness_gate_status(false, true),
    "Portal-mediated resource access must be reviewed before execution."
  );
  execution_readiness_add_gate(
    readiness,
    "snapshot-baseline",
    execution_readiness_gate_status(false, true),
    "A Runtime-managed restore point must exist before compatibility execution."
  );
  execution_readiness_add_gate(
    readiness,
    "backend-binding",
    execution_readiness_gate_status(run_plan.backend_binding_ready, false),
    "A managed compatibility backend binding is required before launch."
  );
  execution_readiness_add_gate(
    readiness,
    "runtime-launch-write-gate",
    launch_gate.dispatch_enabled ? "pass" : "blocked",
    "Runtime Launch remains disabled until production backend ownership is ready."
  );

  execution_readiness_add_blocked_action(readiness, "create execution request");
  execution_readiness_add_blocked_action(readiness, "launch compatibility backend");
  execution_readiness_add_blocked_action(readiness, "expose raw backend command to desktop shell");
  execution_readiness_add_blocked_action(readiness, "grant desktop resources without Portal review");
  execution_readiness_add_blocked_action(readiness, "mutate host root during execution readiness planning");

  if (run_plan.backend_binding_ready && launch_gate.dispatch_enabled && test_result.counts.blocked == 0) {
    readiness->execution_state = "ready";
    readiness->overall_status = "ready";
    readiness->recommended_action = "Create a Runtime-owned execution request.";
    readiness->launch_allowed = true;
    readiness->user_action_required = false;
    readiness->blocked_action_count = 0;
    readiness->summary = "Compatibility execution is ready for a Runtime-owned launch request.";
  } else {
    readiness->execution_state = "blocked";
    readiness->overall_status = "not-ready";
    readiness->recommended_action = "Complete Runtime-owned backend, Portal, snapshot, and Launch gates before execution.";
    readiness->summary = "Compatibility execution is not ready; KDE may show the desktop entry but must route launch intent through Runtime gates.";
  }

  return true;
}
