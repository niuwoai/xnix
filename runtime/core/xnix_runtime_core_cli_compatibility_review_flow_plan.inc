static void
print_compatibility_review_flow_application(const XnixRuntimeApplication *application)
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
print_compatibility_review_flow_steps(const XnixRuntimeCompatibilityReviewFlowPlan *plan)
{
  fputc('[', stdout);
  for (size_t index = 0; index < plan->step_count; index++) {
    const XnixRuntimeCompatibilityReviewFlowStep *step = &plan->steps[index];

    if (index > 0) {
      fputs(",", stdout);
    }

    fputs("{", stdout);
    print_string_field("id", step->id);
    fputs(",", stdout);
    print_string_field("label", step->label);
    fputs(",", stdout);
    print_string_field("runtime_method", step->runtime_method);
    fputs(",", stdout);
    print_string_field("status", step->status);
    fputs(",", stdout);
    fputs("\"user_visible\":", stdout);
    print_bool(step->user_visible);
    fputs(",", stdout);
    fputs("\"user_action_required\":", stdout);
    print_bool(step->user_action_required);
    fputs(",", stdout);
    fputs("\"runtime_gate_required\":", stdout);
    print_bool(step->runtime_gate_required);
    fputs(",", stdout);
    fputs("\"blocks_apply\":", stdout);
    print_bool(step->blocks_apply);
    fputs(",", stdout);
    print_string_field("summary", step->summary);
    fputs("}", stdout);
  }
  fputc(']', stdout);
}

static void
print_compatibility_review_flow_plan_record(const XnixRuntimeCompatibilityReviewFlowPlan *plan)
{
  fputs("{", stdout);
  printf("\"version\":\"%s\",", xnix_runtime_version());
  print_string_field("plan_type", plan->plan_type);
  fputs(",", stdout);
  print_string_field("runtime_method", "GetCompatibilityReviewFlowPlan");
  fputs(",", stdout);
  fputs("\"application\":", stdout);
  print_compatibility_review_flow_application(plan->application);
  fputs(",", stdout);
  print_string_field("review_state", plan->review_state);
  fputs(",", stdout);
  print_string_field("section_id", plan->section_id);
  fputs(",", stdout);
  print_string_field("field_id", plan->field_id);
  fputs(",", stdout);
  print_string_field("requested_value", plan->requested_value);
  fputs(",", stdout);
  print_string_field("operation", plan->operation);
  fputs(",", stdout);
  fputs("\"steps\":", stdout);
  print_compatibility_review_flow_steps(plan);
  fputs(",", stdout);
  fputs("\"step_count\":", stdout);
  printf("%zu,", plan->step_count);
  fputs("\"required_review_count\":", stdout);
  printf("%zu,", plan->required_review_count);
  fputs("\"blocked_step_count\":", stdout);
  printf("%zu,", plan->blocked_step_count);
  fputs("\"pending_step_count\":", stdout);
  printf("%zu,", plan->pending_step_count);
  fputs("\"runtime_owned\":", stdout);
  print_bool(plan->runtime_owned);
  fputs(",", stdout);
  fputs("\"c_runtime_backed\":", stdout);
  print_bool(plan->c_runtime_backed);
  fputs(",", stdout);
  fputs("\"kde_policy_owner\":", stdout);
  print_bool(plan->kde_policy_owner);
  fputs(",", stdout);
  fputs("\"user_confirmation_required\":", stdout);
  print_bool(plan->user_confirmation_required);
  fputs(",", stdout);
  fputs("\"portal_policy_review_required\":", stdout);
  print_bool(plan->portal_policy_review_required);
  fputs(",", stdout);
  fputs("\"settings_change_planned\":", stdout);
  print_bool(plan->settings_change_planned);
  fputs(",", stdout);
  fputs("\"permission_review_planned\":", stdout);
  print_bool(plan->permission_review_planned);
  fputs(",", stdout);
  fputs("\"portal_request_planned\":", stdout);
  print_bool(plan->portal_request_planned);
  fputs(",", stdout);
  fputs("\"review_receipt_required\":", stdout);
  print_bool(plan->review_receipt_required);
  fputs(",", stdout);
  fputs("\"apply_enabled\":", stdout);
  print_bool(plan->apply_enabled);
  fputs(",", stdout);
  fputs("\"request_object_created\":", stdout);
  print_bool(plan->request_object_created);
  fputs(",", stdout);
  fputs("\"permission_granted\":", stdout);
  print_bool(plan->permission_granted);
  fputs(",", stdout);
  fputs("\"settings_persisted\":", stdout);
  print_bool(plan->settings_persisted);
  fputs(",", stdout);
  fputs("\"execution_started\":", stdout);
  print_bool(plan->execution_started);
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
print_compatibility_review_flow_plan_for_application(
  const char *application_id,
  const char *section_id,
  const char *field_id,
  const char *requested_value,
  const char *operation
)
{
  XnixRuntimeCompatibilityReviewFlowPlan plan;

  if (!xnix_runtime_compatibility_review_flow_plan(application_id, section_id, field_id, requested_value, operation, &plan)) {
    fprintf(stderr, "xnix-runtime-core: unknown application, Portal operation, or compatibility review flow plan catalog\n");
    return 64;
  }

  print_compatibility_review_flow_plan_record(&plan);
  fputs("\n", stdout);
  return 0;
}
