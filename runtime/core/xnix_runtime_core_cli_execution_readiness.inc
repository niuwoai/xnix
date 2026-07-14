static void
print_execution_readiness_application(const XnixRuntimeApplication *application)
{
  fputs("{", stdout);
  print_string_field("id", application->id);
  fputs(",", stdout);
  print_string_field("name", application->name);
  fputs(",", stdout);
  print_string_field("icon", application->icon);
  fputs(",", stdout);
  print_string_field("runtime_mode", application->runtime_mode);
  fputs(",", stdout);
  print_string_field("launcher_command", application->launcher_command);
  fputs("}", stdout);
}

static void
print_execution_readiness_engine(const XnixRuntimeEngine *engine)
{
  fputs("{", stdout);
  print_string_field("id", engine->id);
  fputs(",", stdout);
  print_string_field("label", engine->label);
  fputs(",", stdout);
  print_string_field("kind", engine->kind);
  fputs(",", stdout);
  fputs("\"ready\":", stdout);
  print_bool(engine->ready);
  fputs(",", stdout);
  fputs("\"launch_enabled\":", stdout);
  print_bool(engine->launch_enabled);
  fputs(",", stdout);
  fputs("\"backend_details_exposed\":", stdout);
  print_bool(engine->backend_details_exposed);
  fputs("}", stdout);
}

static void
print_execution_readiness_gates(const XnixRuntimeExecutionReadiness *readiness)
{
  fputc('[', stdout);
  for (size_t index = 0; index < readiness->gate_count; index++) {
    const XnixRuntimeExecutionReadinessGate *gate = &readiness->gates[index];

    if (index > 0) {
      fputs(",", stdout);
    }

    fputs("{", stdout);
    print_string_field("id", gate->id);
    fputs(",", stdout);
    print_string_field("status", gate->status);
    fputs(",", stdout);
    print_string_field("summary", gate->summary);
    fputs("}", stdout);
  }
  fputc(']', stdout);
}

static void
print_execution_readiness_record(const XnixRuntimeExecutionReadiness *readiness)
{
  fputs("{", stdout);
  printf("\"version\":\"%s\",", xnix_runtime_version());
  print_string_field("readiness_type", "compatibility-execution-readiness");
  fputs(",", stdout);
  print_string_field("runtime_method", "GetExecutionReadiness");
  fputs(",", stdout);
  fputs("\"application\":", stdout);
  print_execution_readiness_application(readiness->application);
  fputs(",", stdout);
  fputs("\"engine\":", stdout);
  print_execution_readiness_engine(readiness->engine);
  fputs(",", stdout);
  print_string_field("execution_state", readiness->execution_state);
  fputs(",", stdout);
  print_string_field("overall_status", readiness->overall_status);
  fputs(",", stdout);
  print_string_field("recommended_action", readiness->recommended_action);
  fputs(",", stdout);
  fputs("\"runtime_owned\":", stdout);
  print_bool(readiness->runtime_owned);
  fputs(",", stdout);
  fputs("\"c_runtime_backed\":", stdout);
  print_bool(readiness->c_runtime_backed);
  fputs(",", stdout);
  fputs("\"kde_policy_owner\":", stdout);
  print_bool(readiness->kde_policy_owner);
  fputs(",", stdout);
  fputs("\"compatibility_center_card\":", stdout);
  print_bool(readiness->compatibility_center_card);
  fputs(",", stdout);
  fputs("\"safe_for_ai_diagnostics\":", stdout);
  print_bool(readiness->safe_for_ai_diagnostics);
  fputs(",", stdout);
  fputs("\"desktop_entry_launch_visible\":", stdout);
  print_bool(readiness->desktop_entry_launch_visible);
  fputs(",", stdout);
  fputs("\"launch_allowed\":", stdout);
  print_bool(readiness->launch_allowed);
  fputs(",", stdout);
  fputs("\"launch_enabled\":", stdout);
  print_bool(readiness->launch_enabled);
  fputs(",", stdout);
  fputs("\"execution_request_created\":", stdout);
  print_bool(readiness->execution_request_created);
  fputs(",", stdout);
  fputs("\"backend_binding_ready\":", stdout);
  print_bool(readiness->backend_binding_ready);
  fputs(",", stdout);
  fputs("\"portal_policy_required\":", stdout);
  print_bool(readiness->portal_policy_required);
  fputs(",", stdout);
  fputs("\"snapshot_required\":", stdout);
  print_bool(readiness->snapshot_required);
  fputs(",", stdout);
  fputs("\"user_action_required\":", stdout);
  print_bool(readiness->user_action_required);
  fputs(",", stdout);
  fputs("\"gates\":", stdout);
  print_execution_readiness_gates(readiness);
  fputs(",", stdout);
  fputs("\"blocked_actions\":", stdout);
  print_string_array(readiness->blocked_actions, readiness->blocked_action_count);
  fputs(",", stdout);
  fputs("\"host_root_modified\":", stdout);
  print_bool(readiness->host_root_modified);
  fputs(",", stdout);
  fputs("\"network_required\":", stdout);
  print_bool(readiness->network_required);
  fputs(",", stdout);
  fputs("\"backend_details_exposed\":", stdout);
  print_bool(readiness->backend_details_exposed);
  fputs(",", stdout);
  print_string_field("desktop_safe_summary", readiness->summary);
  fputs("}", stdout);
}

static int
print_execution_readiness_for_application(const char *application_id)
{
  XnixRuntimeExecutionReadiness readiness;

  if (!xnix_runtime_execution_readiness(application_id, &readiness)) {
    fprintf(stderr, "xnix-runtime-core: unknown application or execution readiness catalog\n");
    return 64;
  }

  print_execution_readiness_record(&readiness);
  fputs("\n", stdout);
  return 0;
}
