static const XnixRuntimeKDEApplicationSurfaceEntryPoint kde_application_surface_entry_points[] = {
  {
    .id = "launcher",
    .name = "Launcher",
    .runtime_method = "GetDesktopEntryPlan",
    .state = "planned",
    .runtime_backed = true,
    .c_runtime_backed = true,
    .kde_writes_policy = false,
    .backend_details_exposed = false,
    .summary = "KDE shows a standard desktop entry that routes activation through the Runtime.",
  },
  {
    .id = "task-manager",
    .name = "Task Manager",
    .runtime_method = "GetTaskManagerIdentityPlan",
    .state = "planned",
    .runtime_backed = true,
    .c_runtime_backed = true,
    .kde_writes_policy = false,
    .backend_details_exposed = false,
    .summary = "KDE groups compatibility windows under the generated desktop identity.",
  },
  {
    .id = "file-manager",
    .name = "File Manager",
    .runtime_method = "GetFileAssociationPlan",
    .state = "planned",
    .runtime_backed = true,
    .c_runtime_backed = true,
    .kde_writes_policy = false,
    .backend_details_exposed = false,
    .summary = "Dolphin actions and MIME defaults route file opens through Portal-gated Runtime requests.",
  },
  {
    .id = "system-tray",
    .name = "System Tray",
    .runtime_method = "GetTrayStatus",
    .state = "planned",
    .runtime_backed = true,
    .c_runtime_backed = true,
    .kde_writes_policy = false,
    .backend_details_exposed = false,
    .summary = "KDE displays Runtime attention and compatibility status without bridging live backend trays yet.",
  },
  {
    .id = "notifications",
    .name = "Notifications",
    .runtime_method = "GetNotificationPlan",
    .state = "planned",
    .runtime_backed = true,
    .c_runtime_backed = true,
    .kde_writes_policy = false,
    .backend_details_exposed = false,
    .summary = "KDE receives install, repair, mode-change, and approval notification plans from Runtime.",
  },
  {
    .id = "compatibility-center",
    .name = "Compatibility Center",
    .runtime_method = "GetCompatibilityCenterSummary",
    .state = "planned",
    .runtime_backed = true,
    .c_runtime_backed = true,
    .kde_writes_policy = false,
    .backend_details_exposed = false,
    .summary = "KDE renders Runtime-owned compatibility status, actions, diagnostics, and review records.",
  },
  {
    .id = "settings",
    .name = "Settings",
    .runtime_method = "GetCompatibilitySettings",
    .state = "planned",
    .runtime_backed = true,
    .c_runtime_backed = true,
    .kde_writes_policy = false,
    .backend_details_exposed = false,
    .summary = "KDE exposes user-facing compatibility settings without backend terminology.",
  },
};

static const char *kde_application_surface_required_runtime_gates[] = {
  "recipe-install-gate",
  "portal-policy-review",
  "snapshot-baseline",
  "backend-environment-plan",
  "backend-lifecycle-plan",
  "runtime-write-gate",
};

bool
xnix_runtime_kde_application_surface_plan(
  const char *application_id,
  XnixRuntimeKDEApplicationSurfacePlan *plan
)
{
  const XnixRuntimeApplication *application;

  if (!application_id || !plan) {
    return false;
  }

  application = xnix_runtime_find_application(application_id);
  if (!application) {
    return false;
  }

  memset(plan, 0, sizeof(*plan));
  plan->application = application;
  plan->plan_type = "kde-application-surface-plan";
  plan->runtime_method = "GetKDEApplicationSurfacePlan";
  plan->surface_state = "planned";
  plan->desktop_shell = "KDE Plasma";
  plan->entry_point_count = sizeof(kde_application_surface_entry_points) / sizeof(kde_application_surface_entry_points[0]);
  for (size_t index = 0; index < plan->entry_point_count; index++) {
    plan->entry_points[index] = kde_application_surface_entry_points[index];
  }
  plan->required_runtime_gate_count = sizeof(kde_application_surface_required_runtime_gates) / sizeof(kde_application_surface_required_runtime_gates[0]);
  for (size_t index = 0; index < plan->required_runtime_gate_count; index++) {
    plan->required_runtime_gates[index] = kde_application_surface_required_runtime_gates[index];
  }
  plan->runtime_owned = true;
  plan->c_runtime_backed = true;
  plan->kde_policy_owner = false;
  plan->official_desktop_only = true;
  plan->normal_linux_application_surface = true;
  plan->standard_launcher_visible = true;
  plan->task_manager_identity_ready = true;
  plan->file_associations_planned = true;
  plan->dolphin_action_planned = true;
  plan->krunner_query_planned = true;
  plan->tray_status_planned = true;
  plan->notification_route_planned = true;
  plan->settings_surface_planned = true;
  plan->portal_review_required = true;
  plan->execution_ready = false;
  plan->launch_enabled = false;
  plan->backend_process_started = false;
  plan->desktop_files_written = false;
  plan->mimeapps_written = false;
  plan->host_root_modified = false;
  plan->backend_command_exposed = false;
  plan->raw_windows_executable_exposed = false;
  plan->backend_details_exposed = false;
  plan->summary = "KDE can present this compatibility application as a normal Linux application while Runtime gates still block execution.";

  return true;
}
