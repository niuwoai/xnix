static void
print_ai_repair_required_gate(const XnixRuntimeAIRepairRequiredGate *gate)
{
  fputs("{", stdout);
  print_string_field("id", gate->id);
  fputs(",", stdout);
  print_string_field("status", gate->status);
  fputs(",", stdout);
  print_string_field("summary", gate->summary);
  fputs("}", stdout);
}

static void
print_ai_repair_required_gates(const XnixRuntimeAIRepairApprovalGate *gate)
{
  fputc('[', stdout);
  for (size_t index = 0; index < gate->required_gate_count; index++) {
    if (index > 0) {
      fputs(",", stdout);
    }
    print_ai_repair_required_gate(&gate->required_gates[index]);
  }
  fputc(']', stdout);
}

static void
print_ai_repair_approval_actions(const XnixRuntimeAIRepairApprovalGate *gate)
{
  fputc('[', stdout);
  for (size_t index = 0; index < gate->recommendation.approval_required_action_count; index++) {
    if (index > 0) {
      fputs(",", stdout);
    }
    print_ai_approval_action(&gate->recommendation.approval_required_actions[index]);
  }
  fputc(']', stdout);
}

static void
print_ai_repair_approval_gate_record(const XnixRuntimeAIRepairApprovalGate *gate)
{
  const XnixRuntimeAIDiagnosticRecommendation *recommendation = &gate->recommendation;
  const XnixRuntimeAIDiagnosticInput *input = &recommendation->diagnostic_input;

  fputs("{", stdout);
  printf("\"version\":\"%s\",", xnix_runtime_version());
  print_string_field("gate_type", gate->gate_type);
  fputs(",", stdout);
  print_string_field("recommendation_type", recommendation->recommendation_type);
  fputs(",", stdout);
  fputs("\"application\":", stdout);
  print_ai_diagnostic_application(input->application);
  fputs(",", stdout);
  print_string_field("issue", input->issue);
  fputs(",", stdout);
  print_string_field("test_type", input->test_type);
  fputs(",", stdout);
  print_string_field("runtime_method", "GetAIRepairApprovalGate");
  fputs(",", stdout);
  fputs("\"runtime_owned\":", stdout);
  print_bool(gate->runtime_owned);
  fputs(",", stdout);
  fputs("\"c_runtime_backed\":", stdout);
  print_bool(gate->c_runtime_backed);
  fputs(",", stdout);
  fputs("\"kde_policy_owner\":", stdout);
  print_bool(gate->kde_policy_owner);
  fputs(",", stdout);
  fputs("\"ai_provider_called\":", stdout);
  print_bool(gate->ai_provider_called);
  fputs(",", stdout);
  fputs("\"network_required\":", stdout);
  print_bool(gate->network_required);
  fputs(",", stdout);
  fputs("\"safe_for_ai_diagnostics\":", stdout);
  print_bool(gate->safe_for_ai_diagnostics);
  fputs(",", stdout);
  fputs("\"repair_execution_requested\":", stdout);
  print_bool(gate->repair_execution_requested);
  fputs(",", stdout);
  fputs("\"repair_executed\":", stdout);
  print_bool(gate->repair_executed);
  fputs(",", stdout);
  fputs("\"auto_execution_allowed\":", stdout);
  print_bool(gate->auto_execution_allowed);
  fputs(",", stdout);
  print_string_field("gate_decision", gate->gate_decision);
  fputs(",", stdout);
  print_string_field("approval_surface", gate->approval_surface);
  fputs(",", stdout);
  fputs("\"required_gates\":", stdout);
  print_ai_repair_required_gates(gate);
  fputs(",", stdout);
  fputs("\"approval_required_actions\":", stdout);
  print_ai_repair_approval_actions(gate);
  fputs(",", stdout);
  print_ai_task_list("blocked_actions", gate->blocked_actions, gate->blocked_action_count);
  fputs(",", stdout);
  fputs("\"host_root_modified\":", stdout);
  print_bool(gate->host_root_modified);
  fputs(",", stdout);
  fputs("\"backend_details_exposed\":", stdout);
  print_bool(gate->backend_details_exposed);
  fputs(",", stdout);
  print_string_field("desktop_safe_summary", gate->summary);
  fputs("}", stdout);
}

static int
print_ai_repair_approval_gate_for_application(
  const char *application_id,
  const char *issue,
  const char *test_type
)
{
  XnixRuntimeAIRepairApprovalGate gate;

  if (!xnix_runtime_ai_repair_approval_gate(application_id, issue, test_type, &gate)) {
    fprintf(stderr, "xnix-runtime-core: unknown application or AI repair approval gate catalog\n");
    return 64;
  }

  print_ai_repair_approval_gate_record(&gate);
  fputs("\n", stdout);
  return 0;
}
