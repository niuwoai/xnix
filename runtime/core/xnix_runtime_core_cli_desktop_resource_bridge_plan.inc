static void
print_desktop_resource_bridge_application(const XnixRuntimeApplication *application)
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
print_desktop_resource_bridges(const XnixRuntimeDesktopResourceBridgePlan *plan)
{
  fputc('[', stdout);
  for (size_t index = 0; index < plan->resource_count; index++) {
    const XnixRuntimeDesktopResourceBridge *bridge = &plan->resources[index];

    if (index > 0) {
      fputs(",", stdout);
    }

    fputs("{", stdout);
    print_string_field("id", bridge->id);
    fputs(",", stdout);
    print_string_field("name", bridge->name);
    fputs(",", stdout);
    print_string_field("operation", bridge->operation);
    fputs(",", stdout);
    print_string_field("portal_interface", bridge->portal_interface);
    fputs(",", stdout);
    print_string_field("runtime_method", bridge->runtime_method);
    fputs(",", stdout);
    print_string_field("state", bridge->state);
    fputs(",", stdout);
    fputs("\"portal_required\":", stdout);
    print_bool(bridge->portal_required);
    fputs(",", stdout);
    fputs("\"user_approval_required\":", stdout);
    print_bool(bridge->user_approval_required);
    fputs(",", stdout);
    fputs("\"bridge_enabled\":", stdout);
    print_bool(bridge->bridge_enabled);
    fputs(",", stdout);
    fputs("\"request_created\":", stdout);
    print_bool(bridge->request_created);
    fputs(",", stdout);
    fputs("\"direct_backend_access_allowed\":", stdout);
    print_bool(bridge->direct_backend_access_allowed);
    fputs(",", stdout);
    fputs("\"backend_details_exposed\":", stdout);
    print_bool(bridge->backend_details_exposed);
    fputs(",", stdout);
    print_string_field("summary", bridge->summary);
    fputs("}", stdout);
  }
  fputc(']', stdout);
}

static void
print_desktop_resource_bridge_plan_record(const XnixRuntimeDesktopResourceBridgePlan *plan)
{
  fputs("{", stdout);
  printf("\"version\":\"%s\",", xnix_runtime_version());
  print_string_field("plan_type", "desktop-resource-bridge-plan");
  fputs(",", stdout);
  print_string_field("runtime_method", "GetDesktopResourceBridgePlan");
  fputs(",", stdout);
  fputs("\"application\":", stdout);
  print_desktop_resource_bridge_application(plan->application);
  fputs(",", stdout);
  print_string_field("bridge_state", plan->bridge_state);
  fputs(",", stdout);
  fputs("\"resources\":", stdout);
  print_desktop_resource_bridges(plan);
  fputs(",", stdout);
  fputs("\"resource_count\":", stdout);
  printf("%zu,", plan->resource_count);
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
  fputs("\"portal_mediated\":", stdout);
  print_bool(plan->portal_mediated);
  fputs(",", stdout);
  fputs("\"file_bridge_planned\":", stdout);
  print_bool(plan->file_bridge_planned);
  fputs(",", stdout);
  fputs("\"uri_bridge_planned\":", stdout);
  print_bool(plan->uri_bridge_planned);
  fputs(",", stdout);
  fputs("\"print_bridge_planned\":", stdout);
  print_bool(plan->print_bridge_planned);
  fputs(",", stdout);
  fputs("\"clipboard_bridge_planned\":", stdout);
  print_bool(plan->clipboard_bridge_planned);
  fputs(",", stdout);
  fputs("\"screenshot_bridge_planned\":", stdout);
  print_bool(plan->screenshot_bridge_planned);
  fputs(",", stdout);
  fputs("\"bridges_enabled\":", stdout);
  print_bool(plan->bridges_enabled);
  fputs(",", stdout);
  fputs("\"requests_created\":", stdout);
  print_bool(plan->requests_created);
  fputs(",", stdout);
  fputs("\"backend_process_started\":", stdout);
  print_bool(plan->backend_process_started);
  fputs(",", stdout);
  fputs("\"direct_host_file_access\":", stdout);
  print_bool(plan->direct_host_file_access);
  fputs(",", stdout);
  fputs("\"direct_clipboard_access\":", stdout);
  print_bool(plan->direct_clipboard_access);
  fputs(",", stdout);
  fputs("\"direct_print_access\":", stdout);
  print_bool(plan->direct_print_access);
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
print_desktop_resource_bridge_plan_for_application(const char *application_id)
{
  XnixRuntimeDesktopResourceBridgePlan plan;

  if (!xnix_runtime_desktop_resource_bridge_plan(application_id, &plan)) {
    fprintf(stderr, "xnix-runtime-core: unknown application or desktop resource bridge plan catalog\n");
    return 64;
  }

  print_desktop_resource_bridge_plan_record(&plan);
  fputs("\n", stdout);
  return 0;
}
