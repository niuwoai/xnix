static void
print_compatibility_permission_review_application(const XnixRuntimeApplication *application)
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
print_compatibility_permission_review_permissions(const XnixRuntimeCompatibilityPermissionReviewPlan *plan)
{
  fputc('[', stdout);
  for (size_t index = 0; index < plan->permission_count; index++) {
    const XnixRuntimeCompatibilityPermission *permission = &plan->permissions[index];

    if (index > 0) {
      fputs(",", stdout);
    }

    fputs("{", stdout);
    print_string_field("id", permission->id);
    fputs(",", stdout);
    print_string_field("label", permission->label);
    fputs(",", stdout);
    print_string_field("operation", permission->operation);
    fputs(",", stdout);
    print_string_field("decision", permission->decision);
    fputs(",", stdout);
    print_string_field("portal_interface", permission->portal_interface);
    fputs(",", stdout);
    fputs("\"portal_required\":", stdout);
    print_bool(permission->portal_required);
    fputs(",", stdout);
    fputs("\"user_mediation_required\":", stdout);
    print_bool(permission->user_mediation_required);
    fputs(",", stdout);
    print_string_field("current_value", permission->decision);
    fputs(",", stdout);
    print_string_field("requested_value", permission->decision);
    fputs(",", stdout);
    fputs("\"change_pending\":", stdout);
    print_bool(permission->change_pending);
    fputs(",", stdout);
    fputs("\"request_object_created\":", stdout);
    print_bool(permission->request_object_created);
    fputs(",", stdout);
    fputs("\"permission_granted\":", stdout);
    print_bool(permission->permission_granted);
    fputs(",", stdout);
    fputs("\"direct_access_allowed\":", stdout);
    print_bool(permission->direct_access_allowed);
    fputs(",", stdout);
    fputs("\"backend_details_exposed\":", stdout);
    print_bool(permission->backend_details_exposed);
    fputs("}", stdout);
  }
  fputc(']', stdout);
}

static void
print_compatibility_permission_review_plan_record(const XnixRuntimeCompatibilityPermissionReviewPlan *plan)
{
  fputs("{", stdout);
  printf("\"version\":\"%s\",", xnix_runtime_version());
  print_string_field("plan_type", "compatibility-permission-review-plan");
  fputs(",", stdout);
  print_string_field("runtime_method", "GetCompatibilityPermissionReviewPlan");
  fputs(",", stdout);
  fputs("\"application\":", stdout);
  print_compatibility_permission_review_application(plan->application);
  fputs(",", stdout);
  print_string_field("review_state", plan->review_state);
  fputs(",", stdout);
  fputs("\"permissions\":", stdout);
  print_compatibility_permission_review_permissions(plan);
  fputs(",", stdout);
  fputs("\"permission_count\":", stdout);
  printf("%zu,", plan->permission_count);
  fputs("\"allow_count\":", stdout);
  printf("%zu,", plan->allow_count);
  fputs("\"ask_count\":", stdout);
  printf("%zu,", plan->ask_count);
  fputs("\"deny_count\":", stdout);
  printf("%zu,", plan->deny_count);
  fputs("\"required_runtime_gates\":", stdout);
  print_string_array(plan->required_runtime_gates, plan->required_runtime_gate_count);
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
  fputs("\"user_review_required\":", stdout);
  print_bool(plan->user_review_required);
  fputs(",", stdout);
  fputs("\"portal_review_required\":", stdout);
  print_bool(plan->portal_review_required);
  fputs(",", stdout);
  fputs("\"permission_changes_applied\":", stdout);
  print_bool(plan->permission_changes_applied);
  fputs(",", stdout);
  fputs("\"request_objects_created\":", stdout);
  print_bool(plan->request_objects_created);
  fputs(",", stdout);
  fputs("\"permissions_granted\":", stdout);
  print_bool(plan->permissions_granted);
  fputs(",", stdout);
  fputs("\"settings_persisted\":", stdout);
  print_bool(plan->settings_persisted);
  fputs(",", stdout);
  fputs("\"host_permission_changed\":", stdout);
  print_bool(plan->host_permission_changed);
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
print_compatibility_permission_review_plan_for_application(const char *application_id)
{
  XnixRuntimeCompatibilityPermissionReviewPlan plan;

  if (!xnix_runtime_compatibility_permission_review_plan(application_id, &plan)) {
    fprintf(stderr, "xnix-runtime-core: unknown application or compatibility permission review plan catalog\n");
    return 64;
  }

  print_compatibility_permission_review_plan_record(&plan);
  fputs("\n", stdout);
  return 0;
}
