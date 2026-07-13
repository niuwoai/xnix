bool
xnix_runtime_task_manager_identity_plan(
  const char *application_id,
  XnixRuntimeTaskManagerIdentityPlan *plan
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
  plan->launcher_url = "applications:xnix-org.xnix.sample.notepad.desktop";
  plan->window_kind = "compatibility-application";
  plan->class_group = "xnix-compatibility";
  plan->resource_name = application->id;
  plan->title_hint = application->name;
  plan->grouping_key = application->id;
  plan->restore_key = application->id;
  plan->kwin_script_role = "identity-only";
  plan->placement = "normal-window";
  plan->pinning_allowed = true;
  plan->restore_allowed = true;
  plan->skip_taskbar = false;
  plan->show_in_switcher = true;
  plan->prefer_existing_window = true;
  plan->window_manager_policy_only = true;
  plan->runtime_owns_backend_policy = true;
  plan->runtime_owned = true;
  plan->kde_policy_owner = false;
  plan->host_root_modified = false;
  plan->backend_details_exposed = false;
  plan->summary = "task-manager-identity-plan binds compatibility windows to a normal desktop entry for grouping, pinning, switching, and restore.";

  return true;
}
