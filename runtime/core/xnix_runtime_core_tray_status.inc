bool
xnix_runtime_tray_status_plan(XnixRuntimeTrayStatusPlan *plan)
{
  if (plan == NULL) {
    return false;
  }

  plan->status_type = "tray-status-plan";
  plan->desktop = "KDE Plasma";
  plan->active_application_count = 1;
  plan->attention_required_count = 1;
  plan->bridged_tray_application_count = 0;
  plan->runtime_activity_summary = "1 compatibility application active";
  plan->compatibility_state = "attention-required";
  plan->compatibility_label = "Attention required";
  plan->tray_bridge_state = "planned";
  plan->tray_bridge_label = "Tray bridge is planned";
  plan->actions[0] = "open-compatibility-center";
  plan->actions[1] = "open-settings";
  plan->action_count = 2;
  plan->runtime_owned = true;
  plan->kde_policy_owner = false;
  plan->user_visible = true;
  plan->live_backend_bridge_enabled = false;
  plan->bridge_configuration_persisted = false;
  plan->host_root_modified = false;
  plan->backend_details_exposed = false;
  plan->summary = "tray-status-plan reports Runtime activity and tray bridge readiness without enabling backend tray bridging.";

  return true;
}
