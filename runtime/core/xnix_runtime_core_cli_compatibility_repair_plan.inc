static void
print_repair_action(const XnixRuntimeCompatibilityRepairAction *action)
{
  fputs("{", stdout);
  print_string_field("id", action->id);
  fputs(",", stdout);
  print_string_field("kind", action->kind);
  fputs(",", stdout);
  print_string_field("label", action->label);
  fputs("}", stdout);
}

static void
print_repair_actions(const XnixRuntimeCompatibilityRepairPlan *plan)
{
  fputc('[', stdout);
  for (size_t index = 0; index < plan->action_count; index++) {
    if (index > 0) {
      fputs(",", stdout);
    }
    print_repair_action(&plan->actions[index]);
  }
  fputc(']', stdout);
}

static void
print_repair_snapshot_plan_summary(const XnixRuntimeCompatibilityRepairPlan *plan)
{
  const XnixRuntimeSnapshotPolicy *policy = plan->snapshot_policy;

  if (policy == NULL) {
    fputs("null", stdout);
    return;
  }

  fputs("{", stdout);
  print_string_field("plan_type", "compatibility-snapshot");
  fputs(",", stdout);
  print_string_field("reason", policy->reason);
  fputs(",", stdout);
  fputs("\"enabled_by_default\":", stdout);
  print_bool(policy->enabled_by_default);
  fputs(",", stdout);
  fputs("\"restore_available\":", stdout);
  print_bool(policy->restore_available);
  fputs(",", stdout);
  fputs("\"preserve_user_documents\":", stdout);
  print_bool(policy->restore_preserve_user_documents);
  fputs(",", stdout);
  fputs("\"retention\":", stdout);
  print_snapshot_retention(policy);
  fputs("}", stdout);
}

static void
print_compatibility_repair_plan_record(const XnixRuntimeCompatibilityRepairPlan *plan)
{
  fputs("{", stdout);
  printf("\"version\":\"%s\",", xnix_runtime_version());
  print_string_field("plan_type", "compatibility-repair");
  fputs(",", stdout);
  print_string_field("application_id", plan->application_id);
  fputs(",", stdout);
  print_string_field("issue", plan->issue);
  fputs(",", stdout);
  print_string_field("runtime_method", "GetRepairPlan");
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
  print_string_field("severity", plan->severity);
  fputs(",", stdout);
  fputs("\"automatic_allowed\":", stdout);
  print_bool(plan->automatic_allowed);
  fputs(",", stdout);
  fputs("\"user_approval_required\":", stdout);
  print_bool(plan->user_approval_required);
  fputs(",", stdout);
  fputs("\"snapshot_required\":", stdout);
  print_bool(plan->snapshot_required);
  fputs(",", stdout);
  fputs("\"snapshot_plan\":", stdout);
  print_repair_snapshot_plan_summary(plan);
  fputs(",", stdout);
  fputs("\"rollback_available\":", stdout);
  print_bool(plan->rollback_available);
  fputs(",", stdout);
  fputs("\"actions\":", stdout);
  print_repair_actions(plan);
  fputs(",", stdout);
  print_string_field("notification_event", plan->notification_event);
  fputs(",", stdout);
  fputs("\"repair_execution_requested\":", stdout);
  print_bool(plan->repair_execution_requested);
  fputs(",", stdout);
  fputs("\"repair_executed\":", stdout);
  print_bool(plan->repair_executed);
  fputs(",", stdout);
  fputs("\"backend_launch_enabled\":", stdout);
  print_bool(plan->backend_launch_enabled);
  fputs(",", stdout);
  fputs("\"network_required\":", stdout);
  print_bool(plan->network_required);
  fputs(",", stdout);
  fputs("\"host_root_modified\":", stdout);
  print_bool(plan->host_root_modified);
  fputs(",", stdout);
  fputs("\"backend_details_exposed\":", stdout);
  print_bool(plan->backend_details_exposed);
  fputs(",", stdout);
  print_string_field("desktop_safe_summary", plan->summary);
  fputs("}", stdout);
}

static int
print_compatibility_repair_plan_for_application(const char *application_id, const char *issue)
{
  XnixRuntimeCompatibilityRepairPlan plan;

  if (!xnix_runtime_compatibility_repair_plan(application_id, issue, &plan)) {
    fprintf(stderr, "xnix-runtime-core: unknown application or compatibility repair issue catalog\n");
    return 64;
  }

  print_compatibility_repair_plan_record(&plan);
  fputs("\n", stdout);
  return 0;
}
