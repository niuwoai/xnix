bool
xnix_runtime_compatibility_center_summary(
  const char *application_id,
  XnixRuntimeCompatibilityCenterSummary *summary
)
{
  if (summary == NULL) {
    return false;
  }

  const XnixRuntimeApplication *application = xnix_runtime_find_application(application_id);
  if (application == NULL) {
    return false;
  }

  summary->application = application;
  summary->summary_type = "compatibility-center-summary";
  summary->desktop = "KDE Plasma";
  summary->compatibility_state = "review-required";
  summary->runtime_mode = application->runtime_mode;
  summary->known_issue_id = "engine-binding-pending";
  summary->known_issue_severity = "high";
  summary->known_issue_summary = "Runtime backend launch binding is pending and must stay disabled.";
  summary->repair_record_id = "review-ai-repair";
  summary->repair_record_state = "pending-review";
  summary->last_repair_event = "approval-required";
  summary->actions[0] = "open-compatibility-center";
  summary->actions[1] = "review-actions";
  summary->actions[2] = "open-settings";
  summary->actions[3] = "run-preflight-test";
  summary->action_count = 4;
  summary->known_issue_count = 1;
  summary->repair_record_count = 1;
  summary->runtime_owned = true;
  summary->kde_policy_owner = false;
  summary->user_visible = true;
  summary->action_execution_enabled = false;
  summary->repair_execution_enabled = false;
  summary->backend_launch_enabled = false;
  summary->settings_persistence_enabled = false;
  summary->host_root_modified = false;
  summary->backend_details_exposed = false;
  summary->summary = "Compatibility Center summary is Runtime-owned, user-visible, and non-executing.";

  return true;
}
