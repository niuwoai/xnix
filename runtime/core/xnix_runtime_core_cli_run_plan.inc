static void
print_run_plan_application(const XnixRuntimeApplication *application)
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
print_run_plan_backend_binding(const XnixRuntimeCompatibilityRunPlan *plan)
{
  fputs("{", stdout);
  fputs("\"ready\":", stdout);
  print_bool(plan->backend_binding_ready);
  fputs(",", stdout);
  fputs("\"launch_enabled\":", stdout);
  print_bool(plan->launch_enabled);
  fputs(",", stdout);
  print_string_field("reason", "Compatibility backend binding is pending.");
  fputs("}", stdout);
}

static void
print_run_plan_execution(const XnixRuntimeCompatibilityRunPlan *plan)
{
  fputs("{", stdout);
  print_string_field("strategy", plan->engine->id);
  fputs(",", stdout);
  print_string_field("selection_source", "recipe");
  fputs(",", stdout);
  fputs("\"engine\":", stdout);
  print_selected_engine(plan->engine);
  fputs(",", stdout);
  fputs("\"backend_details_exposed\":", stdout);
  print_bool(plan->backend_details_exposed);
  fputs(",", stdout);
  fputs("\"backend_binding\":", stdout);
  print_run_plan_backend_binding(plan);
  fputs(",", stdout);
  fputs("\"execution_request_created\":", stdout);
  print_bool(plan->execution_request_created);
  fputs(",", stdout);
  fputs("\"host_root_modified\":", stdout);
  print_bool(plan->host_root_modified);
  fputs(",", stdout);
  fputs("\"network_required\":", stdout);
  print_bool(plan->network_required);
  fputs("}", stdout);
}

static void
print_run_plan_preflight(const XnixRuntimeCompatibilityRunPlan *plan)
{
  fputs("{", stdout);
  fputs("\"portal_policy_required\":", stdout);
  print_bool(plan->portal_policy_required);
  fputs(",", stdout);
  fputs("\"snapshot_before_risky_change\":", stdout);
  print_bool(plan->snapshot_before_risky_change);
  fputs(",", stdout);
  fputs("\"diagnostics_required\":", stdout);
  print_bool(plan->diagnostics_required);
  fputs("}", stdout);
}

static void
print_compatibility_run_plan_record(const XnixRuntimeCompatibilityRunPlan *plan)
{
  fputs("{", stdout);
  printf("\"version\":\"%s\",", xnix_runtime_version());
  print_string_field("plan_type", "compatibility-run");
  fputs(",", stdout);
  fputs("\"application\":", stdout);
  print_run_plan_application(plan->application);
  fputs(",", stdout);
  fputs("\"runtime_owned\":", stdout);
  print_bool(plan->runtime_owned);
  fputs(",", stdout);
  fputs("\"kde_policy_owner\":", stdout);
  print_bool(plan->kde_policy_owner);
  fputs(",", stdout);
  fputs("\"execution\":", stdout);
  print_run_plan_execution(plan);
  fputs(",", stdout);
  fputs("\"preflight\":", stdout);
  print_run_plan_preflight(plan);
  fputs(",", stdout);
  print_string_field("desktop_safe_summary", plan->summary);
  fputs("}", stdout);
}

static int
print_compatibility_run_plan_for_application(const char *application_id)
{
  XnixRuntimeCompatibilityRunPlan plan;

  if (!xnix_runtime_compatibility_run_plan(application_id, &plan)) {
    fprintf(stderr, "xnix-runtime-core: unknown application or run plan catalog\n");
    return 64;
  }

  print_compatibility_run_plan_record(&plan);
  fputs("\n", stdout);
  return 0;
}
