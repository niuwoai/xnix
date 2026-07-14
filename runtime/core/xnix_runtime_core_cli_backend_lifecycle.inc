static void
print_backend_lifecycle_application(const XnixRuntimeApplication *application)
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
print_backend_lifecycle_engine(const XnixRuntimeEngine *engine)
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
print_backend_lifecycle_stages(const XnixRuntimeBackendLifecycle *lifecycle)
{
  fputc('[', stdout);
  for (size_t index = 0; index < lifecycle->stage_count; index++) {
    const XnixRuntimeBackendLifecycleStage *stage = &lifecycle->stages[index];

    if (index > 0) {
      fputs(",", stdout);
    }

    fputs("{", stdout);
    print_string_field("id", stage->id);
    fputs(",", stdout);
    print_string_field("status", stage->status);
    fputs(",", stdout);
    print_string_field("summary", stage->summary);
    fputs("}", stdout);
  }
  fputc(']', stdout);
}

static void
print_backend_lifecycle_record(const XnixRuntimeBackendLifecycle *lifecycle)
{
  fputs("{", stdout);
  printf("\"version\":\"%s\",", xnix_runtime_version());
  print_string_field("lifecycle_type", "compatibility-backend-lifecycle");
  fputs(",", stdout);
  print_string_field("runtime_method", "GetBackendLifecycle");
  fputs(",", stdout);
  fputs("\"application\":", stdout);
  print_backend_lifecycle_application(lifecycle->application);
  fputs(",", stdout);
  fputs("\"engine\":", stdout);
  print_backend_lifecycle_engine(lifecycle->engine);
  fputs(",", stdout);
  print_string_field("selected_strategy", lifecycle->selected_strategy);
  fputs(",", stdout);
  print_string_field("lifecycle_state", lifecycle->lifecycle_state);
  fputs(",", stdout);
  print_string_field("overall_status", lifecycle->overall_status);
  fputs(",", stdout);
  fputs("\"runtime_owned\":", stdout);
  print_bool(lifecycle->runtime_owned);
  fputs(",", stdout);
  fputs("\"c_runtime_backed\":", stdout);
  print_bool(lifecycle->c_runtime_backed);
  fputs(",", stdout);
  fputs("\"kde_policy_owner\":", stdout);
  print_bool(lifecycle->kde_policy_owner);
  fputs(",", stdout);
  fputs("\"backend_binding_ready\":", stdout);
  print_bool(lifecycle->backend_binding_ready);
  fputs(",", stdout);
  fputs("\"launch_enabled\":", stdout);
  print_bool(lifecycle->launch_enabled);
  fputs(",", stdout);
  fputs("\"execution_request_created\":", stdout);
  print_bool(lifecycle->execution_request_created);
  fputs(",", stdout);
  fputs("\"backend_process_started\":", stdout);
  print_bool(lifecycle->backend_process_started);
  fputs(",", stdout);
  fputs("\"local_backend_started\":", stdout);
  print_bool(lifecycle->local_backend_started);
  fputs(",", stdout);
  fputs("\"isolated_backend_started\":", stdout);
  print_bool(lifecycle->isolated_backend_started);
  fputs(",", stdout);
  fputs("\"state_root_ready\":", stdout);
  print_bool(lifecycle->state_root_ready);
  fputs(",", stdout);
  fputs("\"portal_review_required\":", stdout);
  print_bool(lifecycle->portal_review_required);
  fputs(",", stdout);
  fputs("\"snapshot_required\":", stdout);
  print_bool(lifecycle->snapshot_required);
  fputs(",", stdout);
  fputs("\"stages\":", stdout);
  print_backend_lifecycle_stages(lifecycle);
  fputs(",", stdout);
  fputs("\"blocked_actions\":", stdout);
  print_string_array(lifecycle->blocked_actions, lifecycle->blocked_action_count);
  fputs(",", stdout);
  fputs("\"host_root_modified\":", stdout);
  print_bool(lifecycle->host_root_modified);
  fputs(",", stdout);
  fputs("\"network_required\":", stdout);
  print_bool(lifecycle->network_required);
  fputs(",", stdout);
  fputs("\"backend_details_exposed\":", stdout);
  print_bool(lifecycle->backend_details_exposed);
  fputs(",", stdout);
  print_string_field("desktop_safe_summary", lifecycle->summary);
  fputs("}", stdout);
}

static int
print_backend_lifecycle_for_application(const char *application_id)
{
  XnixRuntimeBackendLifecycle lifecycle;

  if (!xnix_runtime_backend_lifecycle(application_id, &lifecycle)) {
    fprintf(stderr, "xnix-runtime-core: unknown application or backend lifecycle catalog\n");
    return 64;
  }

  print_backend_lifecycle_record(&lifecycle);
  fputs("\n", stdout);
  return 0;
}
