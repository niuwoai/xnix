static bool
xnix_runtime_known_test_type(const char *test_type)
{
  return strcmp(test_type, "preflight") == 0 ||
         strcmp(test_type, "smoke") == 0 ||
         strcmp(test_type, "repair-readiness") == 0;
}

bool
xnix_runtime_compatibility_test_plan(
  const char *application_id,
  const char *test_type,
  XnixRuntimeCompatibilityTestPlan *plan
)
{
  const XnixRuntimeApplication *application = NULL;
  XnixRuntimeCompatibilityRunPlan run_plan;
  XnixRuntimeCompatibilitySnapshotPlan snapshot_plan;

  if (application_id == NULL || test_type == NULL || plan == NULL) {
    return false;
  }

  application = xnix_runtime_find_application(application_id);
  if (application == NULL || !xnix_runtime_known_test_type(test_type)) {
    return false;
  }

  if (!xnix_runtime_compatibility_run_plan(application_id, &run_plan)) {
    return false;
  }

  if (!xnix_runtime_compatibility_snapshot_plan(application_id, "before-engine-change", &snapshot_plan)) {
    return false;
  }

  memset(plan, 0, sizeof(*plan));
  plan->application = application;
  plan->test_type = test_type;
  plan->desktop = "KDE Plasma";
  plan->runtime_owned = true;
  plan->c_runtime_backed = true;
  plan->kde_policy_owner = false;
  plan->blocked = false;
  plan->compatibility_center_card = true;
  plan->notification_event = "mode-changed";
  plan->repair_plan_issue = "engine-binding-pending";
  plan->execution_request_created = false;
  plan->test_executed = false;
  plan->host_root_modified = false;
  plan->backend_details_exposed = false;
  plan->summary = "Compatibility test plan is ready for Runtime-controlled preflight work.";

  plan->steps[0].id = "recipe-validation";
  plan->steps[0].status = "pass";
  plan->steps[0].title = "Validate application recipe";
  plan->steps[0].summary = "Recipe metadata can produce desktop integration artifacts.";
  plan->steps[0].required_for[0] = "launcher";
  plan->steps[0].required_for[1] = "file-association";
  plan->steps[0].required_for[2] = "diagnostics";
  plan->steps[0].required_for_count = 3;

  plan->steps[1].id = "portal-preflight";
  plan->steps[1].status = "pending";
  plan->steps[1].title = "Prepare user-mediated file access";
  plan->steps[1].summary = "File access must use an XDG Desktop Portal request before selected documents are opened.";
  plan->steps[1].portal_destination = "org.freedesktop.portal.Desktop";
  plan->steps[1].portal_interface = "org.freedesktop.portal.FileChooser";
  plan->steps[1].portal_method = "OpenFile";
  plan->steps[1].portal_handle_token = "xnix_org_xnix_sample_notepad_file_open";

  plan->steps[2].id = "snapshot-preflight";
  plan->steps[2].status = "pending";
  plan->steps[2].title = "Prepare restore point";
  plan->steps[2].summary = "Create an application restore point before risky compatibility changes.";
  plan->steps[2].snapshot_reason = snapshot_plan.policy->reason;
  plan->steps[2].snapshot_enabled_by_default = snapshot_plan.policy->enabled_by_default;
  plan->steps[2].snapshot_restore_available = snapshot_plan.policy->restore_available;

  plan->steps[3].id = "runtime-launch-binding";
  plan->steps[3].status = run_plan.backend_binding_ready ? "pass" : "pending";
  plan->steps[3].title = "Bind managed launch backend";
  plan->steps[3].summary = "Runtime launch backend binding is required before automated compatibility execution.";
  plan->steps[3].run_strategy = run_plan.engine->id;
  plan->steps[3].backend_ready = run_plan.backend_binding_ready;
  plan->steps[3].run_plan_backend_details_exposed = run_plan.backend_details_exposed;
  plan->step_count = 4;

  return true;
}
