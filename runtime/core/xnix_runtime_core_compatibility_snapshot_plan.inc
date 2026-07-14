bool
xnix_runtime_compatibility_snapshot_plan(
  const char *application_id,
  const char *reason,
  XnixRuntimeCompatibilitySnapshotPlan *plan
)
{
  const XnixRuntimeSnapshotPolicy *policy = NULL;

  if (application_id == NULL || reason == NULL || plan == NULL) {
    return false;
  }

  if (xnix_runtime_find_application(application_id) == NULL) {
    return false;
  }

  policy = xnix_runtime_find_snapshot_policy(reason);
  if (policy == NULL) {
    return false;
  }

  plan->application_id = application_id;
  plan->policy = policy;
  plan->runtime_owned = true;
  plan->c_runtime_backed = true;
  plan->kde_policy_owner = false;
  plan->snapshot_request_created = false;
  plan->snapshot_created = false;
  plan->restore_requested = false;
  plan->restore_executed = false;
  plan->user_documents_included = policy->user_documents;
  plan->host_system_included = policy->host_system;
  plan->host_root_modified = false;
  plan->backend_details_exposed = false;

  return true;
}
