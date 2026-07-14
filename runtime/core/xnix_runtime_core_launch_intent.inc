static void
launch_intent_add_blocked_action(
  XnixRuntimeLaunchIntent *intent,
  const char *action
)
{
  intent->blocked_actions[intent->blocked_action_count] = action;
  intent->blocked_action_count++;
}

bool
xnix_runtime_launch_intent(
  const char *application_id,
  XnixRuntimeLaunchIntent *intent
)
{
  XnixRuntimeExecutionReadiness readiness;
  XnixRuntimeDesktopEntryPlan desktop_entry;
  XnixRuntimeWriteGate launch_gate;

  if (application_id == NULL || intent == NULL) {
    return false;
  }

  if (!xnix_runtime_execution_readiness(application_id, &readiness)) {
    return false;
  }

  if (!xnix_runtime_desktop_entry_plan(application_id, &desktop_entry)) {
    return false;
  }

  if (!xnix_runtime_write_gate("Launch", &launch_gate)) {
    return false;
  }

  memset(intent, 0, sizeof(*intent));
  intent->application = readiness.application;
  intent->engine = readiness.engine;
  intent->intent_type = "runtime-launch-intent";
  intent->source = "desktop-launcher";
  intent->runtime_method = "Launch";
  intent->read_method = "GetLaunchIntent";
  intent->desktop_file = desktop_entry.desktop_file;
  intent->launcher_command = desktop_entry.exec;
  intent->execution_state = readiness.execution_state;
  intent->overall_status = readiness.overall_status;
  intent->write_gate_decision = launch_gate.decision;
  intent->denial_error_name = XNIX_RUNTIME_WRITE_ERROR;
  intent->runtime_owned = true;
  intent->c_runtime_backed = true;
  intent->kde_policy_owner = false;
  intent->standard_desktop_entry = desktop_entry.standard_desktop_entry;
  intent->launch_uses_runtime = desktop_entry.launch_uses_runtime;
  intent->desktop_entry_launch_visible = readiness.desktop_entry_launch_visible;
  intent->portal_required = false;
  intent->snapshot_required = readiness.snapshot_required;
  intent->backend_binding_ready = readiness.backend_binding_ready;
  intent->launch_allowed = readiness.launch_allowed;
  intent->launch_enabled = readiness.launch_enabled && launch_gate.dispatch_enabled;
  intent->execution_request_created = false;
  intent->execution_started = false;
  intent->host_root_modified = false;
  intent->network_required = false;
  intent->backend_details_exposed = false;

  launch_intent_add_blocked_action(intent, "create launch request object before Runtime gates pass");
  launch_intent_add_blocked_action(intent, "start Wine or VM backend from KDE");
  launch_intent_add_blocked_action(intent, "expose raw backend command to desktop shell");
  launch_intent_add_blocked_action(intent, "grant desktop resources without Portal review");
  launch_intent_add_blocked_action(intent, "mutate host root during launch intent planning");

  if (intent->launch_allowed && intent->launch_enabled) {
    intent->summary = "KDE launch intent is ready to become a Runtime-owned Launch request.";
    intent->blocked_action_count = 0;
  } else {
    intent->summary = "KDE launch intent is captured without executing; Launch remains gated by the Runtime.";
  }

  return true;
}
