bool
xnix_runtime_compatibility_run_plan(
  const char *application_id,
  XnixRuntimeCompatibilityRunPlan *plan
)
{
  const XnixRuntimeApplication *application = NULL;
  const XnixRuntimeEngine *engine = NULL;

  if (application_id == NULL || plan == NULL) {
    return false;
  }

  application = xnix_runtime_find_application(application_id);
  if (application == NULL) {
    return false;
  }

  engine = xnix_runtime_select_engine_for_mode(application->runtime_mode);
  if (engine == NULL) {
    return false;
  }

  plan->application = application;
  plan->engine = engine;
  plan->runtime_owned = true;
  plan->kde_policy_owner = false;
  plan->portal_policy_required = true;
  plan->snapshot_before_risky_change = true;
  plan->diagnostics_required = true;
  plan->backend_binding_ready = engine->ready;
  plan->launch_enabled = engine->launch_enabled;
  plan->execution_request_created = false;
  plan->host_root_modified = false;
  plan->network_required = false;
  plan->backend_details_exposed = false;
  plan->summary = engine->summary;

  return true;
}
