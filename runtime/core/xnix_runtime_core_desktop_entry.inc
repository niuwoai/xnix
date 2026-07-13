bool
xnix_runtime_desktop_entry_plan(
  const char *application_id,
  XnixRuntimeDesktopEntryPlan *plan
)
{
  const XnixRuntimeApplication *application = NULL;

  if (application_id == NULL || plan == NULL) {
    return false;
  }

  application = xnix_runtime_find_application(application_id);
  if (application == NULL) {
    return false;
  }

  plan->application = application;
  plan->desktop_file = "xnix-org.xnix.sample.notepad.desktop";
  plan->relative_path = "applications/xnix-org.xnix.sample.notepad.desktop";
  plan->name = application->name;
  plan->comment = "Run with Xnix Compatibility Runtime";
  plan->exec = application->launcher_command;
  plan->icon = application->icon;
  plan->categories[0] = application->desktop_category;
  plan->categories[1] = NULL;
  plan->category_count = 1;
  plan->mime_types[0] = application->primary_mime_type;
  plan->mime_types[1] = NULL;
  plan->mime_type_count = 1;
  plan->startup_wm_class = "xnix-org.xnix.sample.notepad";
  plan->file_argument_mode = "%U";
  plan->standard_desktop_entry = true;
  plan->launch_uses_runtime = true;
  plan->accepts_file_uris = true;
  plan->user_visible = true;
  plan->startup_notify = true;
  plan->terminal = false;
  plan->no_display = false;
  plan->runtime_owned = true;
  plan->kde_policy_owner = false;
  plan->files_written = false;
  plan->host_root_modified = false;
  plan->backend_command_exposed = false;
  plan->raw_windows_executable_exposed = false;
  plan->compatibility_storage_path_exposed = false;
  plan->backend_details_exposed = false;
  plan->summary = "desktop-entry-plan exposes the compatibility application as a normal Linux desktop entry while hiding backend commands and storage details.";

  return true;
}
