static bool
notification_event_allowed(const char *event_type)
{
  return strcmp(event_type, "install-failed") == 0 ||
         strcmp(event_type, "repair-applied") == 0 ||
         strcmp(event_type, "mode-changed") == 0 ||
         strcmp(event_type, "approval-required") == 0;
}

bool
xnix_runtime_notification_plan(
  const char *application_id,
  const char *event_type,
  XnixRuntimeNotificationPlan *plan
)
{
  const XnixRuntimeApplication *application = NULL;

  if (application_id == NULL || event_type == NULL || plan == NULL) {
    return false;
  }

  application = xnix_runtime_find_application(application_id);
  if (application == NULL || !notification_event_allowed(event_type)) {
    return false;
  }

  plan->application = application;
  plan->event_type = event_type;
  plan->desktop_entry = "xnix-org.xnix.sample.notepad.desktop";
  plan->runtime_owned = true;
  plan->kde_policy_owner = false;
  plan->user_visible = true;
  plan->action_execution_enabled = false;
  plan->repair_execution_enabled = false;
  plan->settings_persistence_enabled = false;
  plan->host_root_modified = false;
  plan->backend_details_exposed = false;

  if (strcmp(event_type, "install-failed") == 0) {
    plan->notification_id = "org.xnix.sample.notepad.install-failed";
    plan->title = "Sample Notepad installation needs attention";
    plan->body = "The Runtime could not complete installation planning and needs Compatibility Center review.";
    plan->urgency = "critical";
    plan->category = "compatibility.install";
    plan->actions[0] = "open-compatibility-center";
    plan->actions[1] = "show-diagnostics";
    plan->action_count = 2;
    plan->requires_user_review = true;
    plan->summary = "notification-plan surfaces an installation failure for KDE review without starting repair or install execution.";
    return true;
  }

  if (strcmp(event_type, "repair-applied") == 0) {
    plan->notification_id = "org.xnix.sample.notepad.repair-applied";
    plan->title = "Sample Notepad repair receipt is ready";
    plan->body = "A Runtime repair receipt is available for review before any further action is enabled.";
    plan->urgency = "normal";
    plan->category = "compatibility.repair";
    plan->actions[0] = "open-compatibility-center";
    plan->action_count = 1;
    plan->requires_user_review = false;
    plan->summary = "notification-plan reports repair receipts while repair execution remains disabled in this milestone.";
    return true;
  }

  if (strcmp(event_type, "mode-changed") == 0) {
    plan->notification_id = "org.xnix.sample.notepad.mode-changed";
    plan->title = "Sample Notepad compatibility mode changed";
    plan->body = "The Runtime recorded a compatibility mode change plan for KDE settings review.";
    plan->urgency = "low";
    plan->category = "compatibility.settings";
    plan->actions[0] = "open-settings";
    plan->action_count = 1;
    plan->requires_user_review = false;
    plan->summary = "notification-plan reports mode changes while settings persistence remains disabled.";
    return true;
  }

  plan->notification_id = "org.xnix.sample.notepad.approval-required";
  plan->title = "Sample Notepad needs approval";
  plan->body = "A compatibility action needs explicit review before the Runtime can proceed.";
  plan->urgency = "critical";
  plan->category = "compatibility.approval";
  plan->actions[0] = "open-compatibility-center";
  plan->actions[1] = "review-request";
  plan->action_count = 2;
  plan->requires_user_review = true;
  plan->summary = "notification-plan asks KDE to show a user-review notification while Runtime execution gates stay closed.";
  return true;
}
