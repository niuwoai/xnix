static const XnixRuntimeCompatibilityModeOption compatibility_mode_options[] = {
  {
    .id = "automatic",
    .label = "Automatic",
    .intent = "Let Xnix choose the safest available compatibility path.",
    .backend_details_exposed = false,
  },
  {
    .id = "prefer-performance",
    .label = "Prefer performance",
    .intent = "Prefer a low-overhead compatibility path when policy gates allow it.",
    .backend_details_exposed = false,
  },
  {
    .id = "prefer-compatibility",
    .label = "Prefer compatibility",
    .intent = "Prefer a more conservative compatibility path for difficult applications.",
    .backend_details_exposed = false,
  },
  {
    .id = "isolated-execution",
    .label = "Isolated execution",
    .intent = "Prefer a stronger isolation boundary for higher-risk applications.",
    .backend_details_exposed = false,
  },
};

static const char *compatibility_mode_required_runtime_gates[] = {
  "settings-review",
  "portal-policy-review",
  "snapshot-baseline",
  "backend-environment-plan",
  "runtime-write-gate",
};

static bool
xnix_runtime_is_supported_compatibility_mode(const char *requested_mode)
{
  if (!requested_mode) {
    return false;
  }

  for (size_t index = 0; index < sizeof(compatibility_mode_options) / sizeof(compatibility_mode_options[0]); index++) {
    if (strcmp(compatibility_mode_options[index].id, requested_mode) == 0) {
      return true;
    }
  }

  return false;
}

bool
xnix_runtime_compatibility_mode_switch_plan(
  const char *application_id,
  const char *requested_mode,
  XnixRuntimeCompatibilityModeSwitchPlan *plan
)
{
  const XnixRuntimeApplication *application;

  if (!application_id || !requested_mode || !plan) {
    return false;
  }

  application = xnix_runtime_find_application(application_id);
  if (!application || !xnix_runtime_is_supported_compatibility_mode(requested_mode)) {
    return false;
  }

  memset(plan, 0, sizeof(*plan));
  plan->application = application;
  plan->plan_type = "compatibility-mode-switch-plan";
  plan->runtime_method = "GetCompatibilityModeSwitchPlan";
  plan->current_mode = application->runtime_mode;
  plan->requested_mode = requested_mode;
  plan->mode_state = "planned";
  plan->mode_count = sizeof(compatibility_mode_options) / sizeof(compatibility_mode_options[0]);
  for (size_t index = 0; index < plan->mode_count; index++) {
    plan->modes[index] = compatibility_mode_options[index];
    plan->modes[index].selected = strcmp(compatibility_mode_options[index].id, application->runtime_mode) == 0;
    plan->modes[index].requested = strcmp(compatibility_mode_options[index].id, requested_mode) == 0;
  }
  plan->required_runtime_gate_count = sizeof(compatibility_mode_required_runtime_gates) / sizeof(compatibility_mode_required_runtime_gates[0]);
  for (size_t index = 0; index < plan->required_runtime_gate_count; index++) {
    plan->required_runtime_gates[index] = compatibility_mode_required_runtime_gates[index];
  }
  plan->runtime_owned = true;
  plan->c_runtime_backed = true;
  plan->kde_policy_owner = false;
  plan->valid_mode = true;
  plan->requires_user_confirmation = true;
  plan->portal_review_required = true;
  plan->snapshot_required = true;
  plan->settings_persistence_enabled = false;
  plan->backend_reconfiguration_enabled = false;
  plan->backend_process_started = false;
  plan->launch_enabled = false;
  plan->host_root_modified = false;
  plan->backend_details_exposed = false;
  plan->summary = "Compatibility mode switch is planned for user review and cannot change Runtime state in this version.";

  return true;
}
