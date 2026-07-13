static void
fill_kde_integration_entry_point(
  XnixRuntimeKDEIntegrationEntryPoint *entry,
  const char *id,
  const char *name,
  const char *model_command,
  const char *runtime_method,
  const char *adapter_role,
  bool c_runtime_backed,
  bool dbus_read_available,
  const char *summary,
  const char *next_step
)
{
  entry->id = id;
  entry->name = name;
  entry->state = "initial";
  entry->model_command = model_command;
  entry->runtime_method = runtime_method;
  entry->adapter_role = adapter_role;
  entry->runtime_backed = true;
  entry->c_runtime_backed = c_runtime_backed;
  entry->dbus_read_available = dbus_read_available;
  entry->kde_policy_owner = false;
  entry->summary = summary;
  entry->next_step = next_step;
}

bool
xnix_runtime_kde_integration_status(XnixRuntimeKDEIntegrationStatus *status)
{
  if (status == NULL) {
    return false;
  }

  fill_kde_integration_entry_point(
    &status->entry_points[0],
    "launcher",
    "Launcher",
    "xnix-runtime-core desktop-entry-plan",
    "GetDesktopEntryPlan",
    "standard-desktop-entry",
    true,
    true,
    "Windows applications appear in the KDE launcher through generated desktop entries.",
    "Connect activation receipts to a production recipe installer."
  );
  fill_kde_integration_entry_point(
    &status->entry_points[1],
    "task-manager",
    "Task Manager",
    "xnix-runtime-core task-manager-identity-plan",
    "GetTaskManagerIdentityPlan",
    "window-identity-and-restore",
    true,
    true,
    "Compatibility windows expose grouping, pinning, switcher, and restore identity.",
    "Connect Runtime identity plans to a production KWin script and task manager bridge."
  );
  fill_kde_integration_entry_point(
    &status->entry_points[2],
    "file-manager",
    "File Manager",
    "xnix-runtime-core file-association-plan",
    "GetFileAssociationPlan",
    "dolphin-service-menu-and-mime",
    true,
    true,
    "Dolphin opens selected files through portal-mediated Runtime file-open planning.",
    "Connect file-open requests to production Runtime launch requests."
  );
  fill_kde_integration_entry_point(
    &status->entry_points[3],
    "system-tray",
    "System Tray",
    "xnix-runtime-core tray-status-plan",
    "GetTrayStatus",
    "runtime-status-surface",
    true,
    true,
    "The tray can show Runtime activity, attention state, and bridge readiness.",
    "Connect tray status plans to a production Plasma tray surface."
  );
  fill_kde_integration_entry_point(
    &status->entry_points[4],
    "notifications",
    "Notifications",
    "xnix-runtime-core notification-plan",
    "GetNotificationPlan",
    "runtime-event-notification",
    true,
    true,
    "Runtime events map to KDE notification payloads for install, repair, mode, and approval states.",
    "Connect notification plans to the production KDE notification path."
  );
  fill_kde_integration_entry_point(
    &status->entry_points[5],
    "compatibility-center",
    "Compatibility Center",
    "xnix-runtime-core compatibility-center-summary",
    "GetCompatibilityCenterSummary",
    "plasma-read-model",
    true,
    true,
    "The Compatibility Center can show Runtime-owned application state and safe action cards.",
    "Render live Runtime applications and diagnostics in the Plasmoid."
  );
  fill_kde_integration_entry_point(
    &status->entry_points[6],
    "settings",
    "Settings",
    "xnix-runtime-core settings-policy",
    "GetCompatibilitySettings",
    "user-facing-policy-controls",
    true,
    true,
    "Settings expose user-facing Runtime policy without backend terminology.",
    "Connect settings plans to a production KDE settings module and persisted Runtime policy."
  );

  status->entry_point_count = 7;
  status->initial_count = 7;
  status->planned_count = 0;
  status->complete_count = 0;
  status->runtime_owned = true;
  status->kde_policy_owner = false;
  status->official_desktop_only = true;
  status->stable_desktop_contract = true;
  status->host_root_modified = false;
  status->backend_details_exposed = false;
  status->desktop = "KDE Plasma";
  status->summary = "The kde-integration-status record confirms KDE Plasma is the only official first-release shell, and all seven entry points are Runtime-backed.";

  return true;
}
