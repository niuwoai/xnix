static void
print_compatibility_review_application(const XnixRuntimeApplication *application)
{
  fputs("{", stdout);
  print_string_field("id", application->id);
  fputs(",", stdout);
  print_string_field("name", application->name);
  fputs(",", stdout);
  print_string_field("icon", application->icon);
  fputs("}", stdout);
}

static void
print_compatibility_review_action_summary(
  const XnixRuntimeCompatibilityAction *action
)
{
  fputs("{", stdout);
  print_string_field("id", action->id);
  fputs(",", stdout);
  print_string_field("source_type", action->source_type);
  fputs(",", stdout);
  print_string_field("status", action->status);
  fputs(",", stdout);
  print_string_field("title", action->title);
  fputs(",", stdout);
  fputs("\"user_review_required\":", stdout);
  print_bool(action->user_review_required);
  fputs(",", stdout);
  fputs("\"execution_enabled\":", stdout);
  print_bool(action->execution_enabled);
  fputs(",", stdout);
  fputs("\"backend_details_exposed\":", stdout);
  print_bool(action->backend_details_exposed);
  fputs("}", stdout);
}

static void
print_compatibility_action_review_receipt_record(
  const XnixRuntimeCompatibilityActionReviewReceipt *receipt
)
{
  fputs("{", stdout);
  printf("\"version\":\"%s\",", xnix_runtime_version());
  print_string_field("receipt_type", "compatibility-center-action-review-receipt");
  fputs(",", stdout);
  print_string_field("receipt_id", receipt->receipt_id);
  fputs(",", stdout);
  print_string_field("queue_type", "compatibility-center-action-queue");
  fputs(",", stdout);
  fputs("\"application\":", stdout);
  print_compatibility_review_application(receipt->application);
  fputs(",", stdout);
  fputs("\"action\":", stdout);
  print_compatibility_review_action_summary(&receipt->action);
  fputs(",", stdout);
  print_string_field("decision", receipt->decision);
  fputs(",", stdout);
  fputs("\"decision_recorded\":", stdout);
  print_bool(receipt->decision_recorded);
  fputs(",", stdout);
  fputs("\"runtime_owned\":", stdout);
  print_bool(receipt->runtime_owned);
  fputs(",", stdout);
  fputs("\"kde_policy_owner\":", stdout);
  print_bool(receipt->kde_policy_owner);
  fputs(",", stdout);
  print_string_field("surface", "Compatibility Center");
  fputs(",", stdout);
  fputs("\"execution_enabled\":", stdout);
  print_bool(receipt->execution_enabled);
  fputs(",", stdout);
  fputs("\"repair_execution_enabled\":", stdout);
  print_bool(receipt->repair_execution_enabled);
  fputs(",", stdout);
  fputs("\"settings_persistence_enabled\":", stdout);
  print_bool(receipt->settings_persistence_enabled);
  fputs(",", stdout);
  fputs("\"resource_grant_created\":", stdout);
  print_bool(receipt->resource_grant_created);
  fputs(",", stdout);
  fputs("\"host_root_modified\":", stdout);
  print_bool(receipt->host_root_modified);
  fputs(",", stdout);
  fputs("\"network_required\":", stdout);
  print_bool(receipt->network_required);
  fputs(",", stdout);
  fputs("\"backend_details_exposed\":", stdout);
  print_bool(receipt->backend_details_exposed);
  fputs(",", stdout);
  print_string_field("required_runtime_gate", receipt->action.runtime_gate);
  fputs(",", stdout);
  print_string_field("next_step", receipt->next_step);
  fputs(",", stdout);
  fputs("\"blocked_actions\":", stdout);
  print_string_array(receipt->blocked_actions, receipt->blocked_action_count);
  fputs(",", stdout);
  print_string_field("desktop_safe_summary", receipt->summary);
  fputs("}", stdout);
}

static int
print_compatibility_action_review_receipt_for_application(
  const char *application_id,
  const char *action_id,
  const char *decision
)
{
  XnixRuntimeCompatibilityActionReviewReceipt receipt;

  if (!xnix_runtime_compatibility_action_review_receipt(
        application_id,
        action_id,
        decision,
        &receipt
      )) {
    fprintf(stderr, "xnix-runtime-core: unknown application, action, or review decision\n");
    return 64;
  }

  print_compatibility_action_review_receipt_record(&receipt);
  fputs("\n", stdout);
  return 0;
}
