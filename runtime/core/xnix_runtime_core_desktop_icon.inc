bool
xnix_runtime_desktop_icon_plan(
  const char *application_id,
  XnixRuntimeDesktopIconPlan *plan
)
{
  const XnixRuntimeApplication *application = NULL;
  XnixRuntimeDesktopEntryPlan desktop_entry_plan;

  if (application_id == NULL || plan == NULL) {
    return false;
  }

  application = xnix_runtime_find_application(application_id);
  if (application == NULL) {
    return false;
  }

  if (!xnix_runtime_desktop_entry_plan(application_id, &desktop_entry_plan)) {
    return false;
  }

  plan->application = application;
  plan->plan_type = "desktop-icon-plan";
  plan->runtime_method = "GetDesktopIconPlan";
  plan->desktop_file = desktop_entry_plan.desktop_file;
  plan->launcher_url = "applications:xnix-org.xnix.sample.notepad.desktop";
  plan->target_directory = "xdg-desktop-dir";
  plan->placement = "user-desktop";
  plan->exec = desktop_entry_plan.exec;
  plan->icon = desktop_entry_plan.icon;
  plan->blocked_actions[0] = "copy desktop files to user directories";
  plan->blocked_actions[1] = "write desktop files from desktop icon previews";
  plan->blocked_actions[2] = "persist desktop icon placement";
  plan->blocked_actions[3] = "launch compatibility backends from desktop icon previews";
  plan->blocked_actions[4] = "mutate the host root from desktop icon previews";
  plan->blocked_actions[5] = "expose backend implementation details in desktop icons";
  plan->blocked_action_count = 6;
  plan->standard_desktop_entry = true;
  plan->desktop_icon_visible = true;
  plan->user_visible = true;
  plan->runtime_owned = true;
  plan->c_runtime_backed = true;
  plan->kde_policy_owner = false;
  plan->official_desktop_only = true;
  plan->desktop_file_copy_enabled = false;
  plan->desktop_file_write_enabled = false;
  plan->icon_placement_persisted = false;
  plan->launch_enabled = false;
  plan->backend_launch_enabled = false;
  plan->host_root_modified = false;
  plan->raw_executable_exposed = false;
  plan->backend_details_exposed = false;
  plan->summary = "desktop-icon-plan previews a normal KDE desktop shortcut without copying files, persisting placement, launching backends, or exposing backend details.";

  return true;
}
