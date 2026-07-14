static void
kde_shell_add_component(
  XnixRuntimeKDEShellIntegrationPlan *plan,
  const char *id,
  const char *label,
  const char *state,
  const char *runtime_method,
  const char *summary
)
{
  XnixRuntimeKDEShellComponent *component = &plan->components[plan->component_count];

  component->id = id;
  component->label = label;
  component->state = state;
  component->runtime_method = runtime_method;
  component->runtime_owned = true;
  component->kde_policy_owner = false;
  component->shell_writes_enabled = false;
  component->backend_launch_enabled = false;
  component->backend_details_exposed = false;
  component->summary = summary;
  plan->component_count++;

  if (strcmp(state, "initial") == 0) {
    plan->initial_component_count++;
  } else if (strcmp(state, "planned") == 0) {
    plan->planned_component_count++;
  }
}

bool
xnix_runtime_kde_shell_integration_plan(XnixRuntimeKDEShellIntegrationPlan *plan)
{
  if (!plan) {
    return false;
  }

  memset(plan, 0, sizeof(*plan));
  plan->plan_type = "kde-shell-integration-plan";
  plan->runtime_method = "GetKDEShellIntegrationPlan";
  plan->desktop_shell = "KDE Plasma";
  plan->official_desktop_only = true;
  plan->fallback_desktops_supported = false;
  plan->runtime_owned = true;
  plan->c_runtime_backed = true;
  plan->kde_policy_owner = false;
  plan->plasma_fork_required = false;
  plan->plasma_source_modified = false;
  plan->shell_configuration_written = false;
  plan->component_activation_enabled = false;
  plan->backend_launch_enabled = false;
  plan->backend_details_exposed = false;
  plan->host_root_modified = false;
  plan->privileged_container_required = false;
  plan->summary = "KDE Plasma is the first supported shell; Runtime owns policy and KDE components remain replaceable adapters.";

  kde_shell_add_component(plan, "start-menu", "Start menu", "initial", "GetDesktopEntryPlan", "Generated desktop entries make compatibility applications visible beside native Linux applications.");
  kde_shell_add_component(plan, "task-manager", "Task manager", "initial", "GetTaskManagerIdentityPlan", "Window identity plans let compatibility windows group and restore like normal applications.");
  kde_shell_add_component(plan, "file-manager", "Dolphin file manager", "initial", "GetFileAssociationPlan", "Dolphin actions and MIME associations route file opens through Runtime planning.");
  kde_shell_add_component(plan, "system-tray", "System tray", "planned", "GetTrayStatus", "Tray status is modeled while live tray bridging remains disabled.");
  kde_shell_add_component(plan, "notification-center", "Notification center", "initial", "GetNotificationPlan", "Runtime notification plans describe install, repair, approval, and mode-change events.");
  kde_shell_add_component(plan, "compatibility-center", "Compatibility Center", "initial", "GetCompatibilityCenterSummary", "Compatibility Center reads Runtime diagnostics, review tasks, and repair state.");
  kde_shell_add_component(plan, "unified-settings", "Unified settings", "initial", "GetCompatibilitySettings", "Settings expose user-facing Runtime policy without backend terminology.");
  kde_shell_add_component(plan, "krunner-search", "KRunner search", "initial", "GetKRunnerQueryPlan", "KRunner query plans map natural-language queries to Runtime application identities.");
  kde_shell_add_component(plan, "kwin-window-management", "KWin window management", "initial", "GetKWinWindowRulePlan", "KWin rule plans provide identity and layout hints without KDE owning backend policy.");

  return true;
}
