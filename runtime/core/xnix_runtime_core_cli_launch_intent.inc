static void
print_launch_intent_application(const XnixRuntimeApplication *application)
{
  fputs("{", stdout);
  print_string_field("id", application->id);
  fputs(",", stdout);
  print_string_field("name", application->name);
  fputs(",", stdout);
  print_string_field("icon", application->icon);
  fputs(",", stdout);
  print_string_field("runtime_mode", application->runtime_mode);
  fputs("}", stdout);
}

static void
print_launch_intent_engine(const XnixRuntimeEngine *engine)
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
print_launch_intent_record(const XnixRuntimeLaunchIntent *intent)
{
  fputs("{", stdout);
  printf("\"version\":\"%s\",", xnix_runtime_version());
  print_string_field("intent_type", "runtime-launch-intent");
  fputs(",", stdout);
  print_string_field("source", intent->source);
  fputs(",", stdout);
  print_string_field("runtime_method", intent->runtime_method);
  fputs(",", stdout);
  print_string_field("read_method", "GetLaunchIntent");
  fputs(",", stdout);
  fputs("\"application\":", stdout);
  print_launch_intent_application(intent->application);
  fputs(",", stdout);
  fputs("\"engine\":", stdout);
  print_launch_intent_engine(intent->engine);
  fputs(",", stdout);
  print_string_field("desktop_file", intent->desktop_file);
  fputs(",", stdout);
  print_string_field("launcher_command", intent->launcher_command);
  fputs(",", stdout);
  print_string_field("execution_state", intent->execution_state);
  fputs(",", stdout);
  print_string_field("overall_status", intent->overall_status);
  fputs(",", stdout);
  print_string_field("write_gate_decision", intent->write_gate_decision);
  fputs(",", stdout);
  print_string_field("denial_error_name", intent->denial_error_name);
  fputs(",", stdout);
  fputs("\"runtime_owned\":", stdout);
  print_bool(intent->runtime_owned);
  fputs(",", stdout);
  fputs("\"c_runtime_backed\":", stdout);
  print_bool(intent->c_runtime_backed);
  fputs(",", stdout);
  fputs("\"kde_policy_owner\":", stdout);
  print_bool(intent->kde_policy_owner);
  fputs(",", stdout);
  fputs("\"standard_desktop_entry\":", stdout);
  print_bool(intent->standard_desktop_entry);
  fputs(",", stdout);
  fputs("\"launch_uses_runtime\":", stdout);
  print_bool(intent->launch_uses_runtime);
  fputs(",", stdout);
  fputs("\"desktop_entry_launch_visible\":", stdout);
  print_bool(intent->desktop_entry_launch_visible);
  fputs(",", stdout);
  fputs("\"portal_required\":", stdout);
  print_bool(intent->portal_required);
  fputs(",", stdout);
  fputs("\"snapshot_required\":", stdout);
  print_bool(intent->snapshot_required);
  fputs(",", stdout);
  fputs("\"backend_binding_ready\":", stdout);
  print_bool(intent->backend_binding_ready);
  fputs(",", stdout);
  fputs("\"launch_allowed\":", stdout);
  print_bool(intent->launch_allowed);
  fputs(",", stdout);
  fputs("\"launch_enabled\":", stdout);
  print_bool(intent->launch_enabled);
  fputs(",", stdout);
  fputs("\"execution_request_created\":", stdout);
  print_bool(intent->execution_request_created);
  fputs(",", stdout);
  fputs("\"execution_started\":", stdout);
  print_bool(intent->execution_started);
  fputs(",", stdout);
  fputs("\"blocked_actions\":", stdout);
  print_string_array(intent->blocked_actions, intent->blocked_action_count);
  fputs(",", stdout);
  fputs("\"host_root_modified\":", stdout);
  print_bool(intent->host_root_modified);
  fputs(",", stdout);
  fputs("\"network_required\":", stdout);
  print_bool(intent->network_required);
  fputs(",", stdout);
  fputs("\"backend_details_exposed\":", stdout);
  print_bool(intent->backend_details_exposed);
  fputs(",", stdout);
  print_string_field("desktop_safe_summary", intent->summary);
  fputs("}", stdout);
}

static int
print_launch_intent_for_application(const char *application_id)
{
  XnixRuntimeLaunchIntent intent;

  if (!xnix_runtime_launch_intent(application_id, &intent)) {
    fprintf(stderr, "xnix-runtime-core: unknown application or launch intent catalog\n");
    return 64;
  }

  print_launch_intent_record(&intent);
  fputs("\n", stdout);
  return 0;
}
