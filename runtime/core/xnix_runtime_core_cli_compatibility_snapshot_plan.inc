static void
print_compatibility_snapshot_plan_record(const XnixRuntimeCompatibilitySnapshotPlan *plan)
{
  const XnixRuntimeSnapshotPolicy *policy = plan->policy;

  fputs("{", stdout);
  printf("\"version\":\"%s\",", xnix_runtime_version());
  print_string_field("plan_type", "compatibility-snapshot");
  fputs(",", stdout);
  print_string_field("application_id", plan->application_id);
  fputs(",", stdout);
  print_string_field("reason", policy->reason);
  fputs(",", stdout);
  print_string_field("runtime_method", "GetSnapshotPlan");
  fputs(",", stdout);
  fputs("\"runtime_owned\":", stdout);
  print_bool(plan->runtime_owned);
  fputs(",", stdout);
  fputs("\"c_runtime_backed\":", stdout);
  print_bool(plan->c_runtime_backed);
  fputs(",", stdout);
  fputs("\"kde_policy_owner\":", stdout);
  print_bool(plan->kde_policy_owner);
  fputs(",", stdout);
  fputs("\"enabled_by_default\":", stdout);
  print_bool(policy->enabled_by_default);
  fputs(",", stdout);
  fputs("\"snapshot_scope\":", stdout);
  print_snapshot_scope(policy);
  fputs(",", stdout);
  fputs("\"restore\":", stdout);
  print_snapshot_restore(policy);
  fputs(",", stdout);
  fputs("\"retention\":", stdout);
  print_snapshot_retention(policy);
  fputs(",", stdout);
  fputs("\"snapshot_request_created\":", stdout);
  print_bool(plan->snapshot_request_created);
  fputs(",", stdout);
  fputs("\"snapshot_created\":", stdout);
  print_bool(plan->snapshot_created);
  fputs(",", stdout);
  fputs("\"restore_requested\":", stdout);
  print_bool(plan->restore_requested);
  fputs(",", stdout);
  fputs("\"restore_executed\":", stdout);
  print_bool(plan->restore_executed);
  fputs(",", stdout);
  fputs("\"user_documents_included\":", stdout);
  print_bool(plan->user_documents_included);
  fputs(",", stdout);
  fputs("\"host_system_included\":", stdout);
  print_bool(plan->host_system_included);
  fputs(",", stdout);
  fputs("\"host_root_modified\":", stdout);
  print_bool(plan->host_root_modified);
  fputs(",", stdout);
  fputs("\"backend_details_exposed\":", stdout);
  print_bool(plan->backend_details_exposed);
  fputs(",", stdout);
  print_string_field("desktop_safe_summary", policy->summary);
  fputs("}", stdout);
}

static int
print_compatibility_snapshot_plan_for_application(const char *application_id, const char *reason)
{
  XnixRuntimeCompatibilitySnapshotPlan plan;

  if (!xnix_runtime_compatibility_snapshot_plan(application_id, reason, &plan)) {
    fprintf(stderr, "xnix-runtime-core: unknown application or compatibility snapshot reason catalog\n");
    return 64;
  }

  print_compatibility_snapshot_plan_record(&plan);
  fputs("\n", stdout);
  return 0;
}
