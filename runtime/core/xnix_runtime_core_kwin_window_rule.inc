bool
xnix_runtime_kwin_window_rule_plan(
  const char *application_id,
  XnixRuntimeKWinWindowRulePlan *plan
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
  plan->request_type = "kwin-window-rule";
  plan->desktop = "KDE Plasma";
  plan->script_role = "identity-and-layout";
  plan->resource_name = application->id;
  plan->class_group = "xnix-compatibility";
  plan->title_hint = application->name;
  plan->desktop_file = "xnix-org.xnix.sample.notepad.desktop";
  plan->application_id = application->id;
  plan->task_manager_grouping_key = application->id;
  plan->launcher_url = "applications:xnix-org.xnix.sample.notepad.desktop";
  plan->placement = "normal-window";
  plan->skip_taskbar = false;
  plan->show_in_switcher = true;
  plan->pinning_allowed = true;
  plan->restore_allowed = true;
  plan->restore_key = application->id;
  plan->prefer_existing_window = true;
  plan->window_manager_policy_only = true;
  plan->runtime_owns_backend_policy = true;
  plan->runtime_owned = true;
  plan->kde_policy_owner = false;
  plan->host_root_modified = false;
  plan->backend_details_exposed = false;
  plan->summary = "KWin window rules are Runtime-owned identity and layout hints for normal desktop behavior.";

  return true;
}
