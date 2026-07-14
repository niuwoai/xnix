static void
print_kde_shell_components(const XnixRuntimeKDEShellIntegrationPlan *plan)
{
  fputc('[', stdout);
  for (size_t index = 0; index < plan->component_count; index++) {
    const XnixRuntimeKDEShellComponent *component = &plan->components[index];

    if (index > 0) {
      fputs(",", stdout);
    }

    fputs("{", stdout);
    print_string_field("id", component->id);
    fputs(",", stdout);
    print_string_field("label", component->label);
    fputs(",", stdout);
    print_string_field("state", component->state);
    fputs(",", stdout);
    print_string_field("runtime_method", component->runtime_method);
    fputs(",", stdout);
    fputs("\"runtime_owned\":", stdout);
    print_bool(component->runtime_owned);
    fputs(",", stdout);
    fputs("\"kde_policy_owner\":", stdout);
    print_bool(component->kde_policy_owner);
    fputs(",", stdout);
    fputs("\"shell_writes_enabled\":", stdout);
    print_bool(component->shell_writes_enabled);
    fputs(",", stdout);
    fputs("\"backend_launch_enabled\":", stdout);
    print_bool(component->backend_launch_enabled);
    fputs(",", stdout);
    fputs("\"backend_details_exposed\":", stdout);
    print_bool(component->backend_details_exposed);
    fputs(",", stdout);
    print_string_field("summary", component->summary);
    fputs("}", stdout);
  }
  fputc(']', stdout);
}

static void
print_kde_shell_integration_plan_record(const XnixRuntimeKDEShellIntegrationPlan *plan)
{
  fputs("{", stdout);
  printf("\"version\":\"%s\",", xnix_runtime_version());
  print_string_field("plan_type", "kde-shell-integration-plan");
  fputs(",", stdout);
  print_string_field("runtime_method", "GetKDEShellIntegrationPlan");
  fputs(",", stdout);
  print_string_field("desktop_shell", plan->desktop_shell);
  fputs(",", stdout);
  fputs("\"official_desktop_only\":", stdout);
  print_bool(plan->official_desktop_only);
  fputs(",", stdout);
  fputs("\"fallback_desktops_supported\":", stdout);
  print_bool(plan->fallback_desktops_supported);
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
  fputs("\"plasma_fork_required\":", stdout);
  print_bool(plan->plasma_fork_required);
  fputs(",", stdout);
  fputs("\"plasma_source_modified\":", stdout);
  print_bool(plan->plasma_source_modified);
  fputs(",", stdout);
  fputs("\"shell_configuration_written\":", stdout);
  print_bool(plan->shell_configuration_written);
  fputs(",", stdout);
  fputs("\"component_activation_enabled\":", stdout);
  print_bool(plan->component_activation_enabled);
  fputs(",", stdout);
  fputs("\"backend_launch_enabled\":", stdout);
  print_bool(plan->backend_launch_enabled);
  fputs(",", stdout);
  fputs("\"backend_details_exposed\":", stdout);
  print_bool(plan->backend_details_exposed);
  fputs(",", stdout);
  fputs("\"host_root_modified\":", stdout);
  print_bool(plan->host_root_modified);
  fputs(",", stdout);
  fputs("\"privileged_container_required\":", stdout);
  print_bool(plan->privileged_container_required);
  fputs(",", stdout);
  printf("\"component_count\":%zu,", plan->component_count);
  printf("\"initial_component_count\":%zu,", plan->initial_component_count);
  printf("\"planned_component_count\":%zu,", plan->planned_component_count);
  fputs("\"components\":", stdout);
  print_kde_shell_components(plan);
  fputs(",", stdout);
  print_string_field("desktop_safe_summary", plan->summary);
  fputs("}", stdout);
}

static int
print_kde_shell_integration_plan(void)
{
  XnixRuntimeKDEShellIntegrationPlan plan;

  if (!xnix_runtime_kde_shell_integration_plan(&plan)) {
    fprintf(stderr, "xnix-runtime-core: KDE shell integration plan is unavailable\n");
    return 64;
  }

  print_kde_shell_integration_plan_record(&plan);
  fputs("\n", stdout);
  return 0;
}
