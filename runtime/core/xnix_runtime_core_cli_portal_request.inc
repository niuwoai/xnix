static void
print_portal_request_application(const XnixRuntimeApplication *application)
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
print_portal_request_portal(const XnixRuntimePortalRequestPlan *plan)
{
  fputs("{", stdout);
  print_string_field("destination", plan->portal_destination);
  fputs(",", stdout);
  print_string_field("interface", plan->policy->portal_interface);
  fputs(",", stdout);
  print_string_field("method", plan->portal_method);
  fputs(",", stdout);
  print_string_field("object_path", plan->portal_object_path);
  fputs(",", stdout);
  print_string_field("dbus_api", "XDG Desktop Portal");
  fputs("}", stdout);
}

static void
print_portal_request_request(const XnixRuntimePortalRequestPlan *plan)
{
  fputs("{", stdout);
  fputs("\"object_path_required\":", stdout);
  print_bool(plan->request_object_required);
  fputs(",", stdout);
  fputs("\"request_object_created\":", stdout);
  print_bool(plan->request_object_created);
  fputs(",", stdout);
  print_string_field("handle_token", plan->handle_token);
  fputs(",", stdout);
  fputs("\"user_mediation_required\":", stdout);
  print_bool(plan->user_mediation_required);
  fputs(",", stdout);
  fputs("\"resources\":", stdout);
  print_string_array(plan->policy->resources, plan->policy->resource_count);
  fputs(",", stdout);
  fputs("\"runtime_policy_owner\":", stdout);
  print_bool(plan->runtime_owned);
  fputs(",", stdout);
  fputs("\"desktop_shell_policy_owner\":", stdout);
  print_bool(plan->kde_policy_owner);
  fputs("}", stdout);
}

static void
print_portal_request_completion(const XnixRuntimePortalRequestPlan *plan)
{
  fputs("{", stdout);
  print_string_field("signal", plan->completion_signal);
  fputs(",", stdout);
  print_string_field("response_field", plan->response_field);
  fputs(",", stdout);
  fputs("\"success_code\":0,", stdout);
  fputs("\"cancelled_code\":1,", stdout);
  fputs("\"denied_code\":2,", stdout);
  print_string_field("result_owner", plan->result_owner);
  fputs("}", stdout);
}

static void
print_portal_request_denied(const XnixRuntimePortalRequestPlan *plan)
{
  if (plan->request_allowed) {
    fputs("null", stdout);
    return;
  }

  fputs("{", stdout);
  print_string_field("reason", plan->policy->summary);
  fputs(",", stdout);
  print_string_field("next_action", "open-compatibility-settings");
  fputs(",", stdout);
  print_string_field("notification_event", "approval-required");
  fputs("}", stdout);
}

static void
print_portal_request_safety(const XnixRuntimePortalRequestPlan *plan)
{
  fputs("{", stdout);
  fputs("\"direct_access_allowed\":", stdout);
  print_bool(plan->direct_access_allowed);
  fputs(",", stdout);
  fputs("\"portal_required\":", stdout);
  print_bool(plan->portal_required);
  fputs(",", stdout);
  fputs("\"permission_granted\":", stdout);
  print_bool(plan->permission_granted);
  fputs(",", stdout);
  fputs("\"host_permission_changed\":", stdout);
  print_bool(plan->host_permission_changed);
  fputs(",", stdout);
  fputs("\"host_root_modified\":", stdout);
  print_bool(plan->host_root_modified);
  fputs(",", stdout);
  fputs("\"backend_details_exposed\":", stdout);
  print_bool(plan->backend_details_exposed);
  fputs("}", stdout);
}

static void
print_portal_request_plan_record(const XnixRuntimePortalRequestPlan *plan)
{
  fputs("{", stdout);
  printf("\"version\":\"%s\",", xnix_runtime_version());
  print_string_field("request_type", "portal-request-plan");
  fputs(",", stdout);
  print_string_field("desktop", "KDE Plasma");
  fputs(",", stdout);
  fputs("\"application\":", stdout);
  print_portal_request_application(plan->application);
  fputs(",", stdout);
  print_string_field("operation", plan->operation);
  fputs(",", stdout);
  print_string_field("reason", plan->reason);
  fputs(",", stdout);
  print_string_field("decision", plan->policy->decision);
  fputs(",", stdout);
  fputs("\"request_allowed\":", stdout);
  print_bool(plan->request_allowed);
  fputs(",", stdout);
  fputs("\"runtime_owned\":", stdout);
  print_bool(plan->runtime_owned);
  fputs(",", stdout);
  fputs("\"kde_policy_owner\":", stdout);
  print_bool(plan->kde_policy_owner);
  fputs(",", stdout);
  fputs("\"portal\":", stdout);
  print_portal_request_portal(plan);
  fputs(",", stdout);
  fputs("\"request\":", stdout);
  print_portal_request_request(plan);
  fputs(",", stdout);
  fputs("\"completion\":", stdout);
  print_portal_request_completion(plan);
  fputs(",", stdout);
  fputs("\"denied\":", stdout);
  print_portal_request_denied(plan);
  fputs(",", stdout);
  fputs("\"safety\":", stdout);
  print_portal_request_safety(plan);
  fputs(",", stdout);
  print_string_field("desktop_safe_summary", plan->summary);
  fputs("}", stdout);
}

static int
print_portal_request_plan_for_application(const char *application_id, const char *operation)
{
  XnixRuntimePortalRequestPlan plan;

  if (!xnix_runtime_portal_request_plan(application_id, operation, &plan)) {
    fprintf(stderr, "xnix-runtime-core: unknown application or portal request plan catalog\n");
    return 64;
  }

  print_portal_request_plan_record(&plan);
  fputs("\n", stdout);
  return 0;
}
