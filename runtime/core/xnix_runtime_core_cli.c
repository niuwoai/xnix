#include "xnix_runtime_core.h"

#include <stdio.h>
#include <string.h>

static void
print_bool(bool value)
{
  fputs(value ? "true" : "false", stdout);
}

static void
print_string(const char *value)
{
  fputc('"', stdout);
  for (const char *cursor = value; cursor != NULL && *cursor != '\0'; cursor++) {
    if (*cursor == '"' || *cursor == '\\') {
      fputc('\\', stdout);
    }
    fputc(*cursor, stdout);
  }
  fputc('"', stdout);
}

static void
print_string_field(const char *name, const char *value)
{
  print_string(name);
  fputc(':', stdout);
  print_string(value);
}

static void
print_write_methods(void)
{
  fputc('[', stdout);
  for (size_t index = 0; index < xnix_runtime_write_method_count(); index++) {
    if (index > 0) {
      fputs(",", stdout);
    }
    print_string(xnix_runtime_write_method_at(index));
  }
  fputc(']', stdout);
}

static void
print_string_array(const char *const *values, size_t count)
{
  fputc('[', stdout);
  for (size_t index = 0; index < count; index++) {
    if (index > 0) {
      fputs(",", stdout);
    }
    print_string(values[index]);
  }
  fputc(']', stdout);
}

static void
print_application(const XnixRuntimeApplication *application)
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
  print_string_field("desktop_category", application->desktop_category);
  fputs(",", stdout);
  print_string_field("launcher_command", application->launcher_command);
  fputs(",", stdout);
  print_string_field("primary_extension", application->primary_extension);
  fputs(",", stdout);
  print_string_field("primary_mime_type", application->primary_mime_type);
  fputs(",", stdout);
  fputs("\"runtime_owned\":", stdout);
  print_bool(application->runtime_owned);
  fputs(",", stdout);
  fputs("\"kde_policy_owner\":", stdout);
  print_bool(application->kde_policy_owner);
  fputs(",", stdout);
  fputs("\"backend_details_exposed\":", stdout);
  print_bool(application->backend_details_exposed);
  fputs("}", stdout);
}

static void
print_portal_request_flow(const XnixRuntimePortalPolicy *policy)
{
  fputs("{", stdout);
  print_string_field("dbus_api", "XDG Desktop Portal");
  fputs(",", stdout);
  fputs("\"request_object_required\":true,", stdout);
  print_string_field("completion_signal", "portal-response");
  fputs(",", stdout);
  fputs("\"runtime_policy_owner\":", stdout);
  print_bool(policy->runtime_policy_owner);
  fputs(",", stdout);
  fputs("\"desktop_shell_policy_owner\":", stdout);
  print_bool(policy->desktop_shell_policy_owner);
  fputs("}", stdout);
}

static void
print_portal_policy_record(const XnixRuntimePortalPolicy *policy)
{
  fputs("{", stdout);
  print_string_field("operation", policy->operation);
  fputs(",", stdout);
  print_string_field("decision", policy->decision);
  fputs(",", stdout);
  fputs("\"portal_required\":", stdout);
  print_bool(policy->portal_required);
  fputs(",", stdout);
  print_string_field("portal_interface", policy->portal_interface);
  fputs(",", stdout);
  fputs("\"user_mediation_required\":", stdout);
  print_bool(policy->user_mediation_required);
  fputs(",", stdout);
  fputs("\"direct_access_allowed\":", stdout);
  print_bool(policy->direct_access_allowed);
  fputs(",", stdout);
  fputs("\"resources\":", stdout);
  print_string_array(policy->resources, policy->resource_count);
  fputs(",", stdout);
  fputs("\"runtime_policy_owner\":", stdout);
  print_bool(policy->runtime_policy_owner);
  fputs(",", stdout);
  fputs("\"desktop_shell_policy_owner\":", stdout);
  print_bool(policy->desktop_shell_policy_owner);
  fputs(",", stdout);
  print_string_field("desktop_safe_summary", policy->summary);
  fputs(",", stdout);
  fputs("\"backend_details_exposed\":", stdout);
  print_bool(policy->backend_details_exposed);
  fputs("}", stdout);
}

static void
print_portal_policy_for_application(const char *application_id, const XnixRuntimePortalPolicy *policy)
{
  fputs("{", stdout);
  printf("\"version\":\"%s\",", xnix_runtime_version());
  print_string_field("policy_type", "portal-access");
  fputs(",", stdout);
  print_string_field("desktop", "KDE Plasma");
  fputs(",", stdout);
  print_string_field("application_id", application_id);
  fputs(",", stdout);
  print_string_field("operation", policy->operation);
  fputs(",", stdout);
  print_string_field("decision", policy->decision);
  fputs(",", stdout);
  fputs("\"portal_required\":", stdout);
  print_bool(policy->portal_required);
  fputs(",", stdout);
  print_string_field("portal_interface", policy->portal_interface);
  fputs(",", stdout);
  fputs("\"user_mediation_required\":", stdout);
  print_bool(policy->user_mediation_required);
  fputs(",", stdout);
  fputs("\"direct_access_allowed\":", stdout);
  print_bool(policy->direct_access_allowed);
  fputs(",", stdout);
  fputs("\"resources\":", stdout);
  print_string_array(policy->resources, policy->resource_count);
  fputs(",", stdout);
  fputs("\"request_flow\":", stdout);
  print_portal_request_flow(policy);
  fputs(",", stdout);
  print_string_field("desktop_safe_summary", policy->summary);
  fputs(",", stdout);
  fputs("\"backend_details_exposed\":", stdout);
  print_bool(policy->backend_details_exposed);
  fputs("}", stdout);
}

static void
print_engine_record(const XnixRuntimeEngine *engine)
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
  fputs("\"user_visible\":", stdout);
  print_bool(engine->user_visible);
  fputs(",", stdout);
  print_string_field("summary", engine->summary);
  fputs(",", stdout);
  fputs("\"backend_details_exposed\":", stdout);
  print_bool(engine->backend_details_exposed);
  fputs("}", stdout);
}

static void
print_selected_engine(const XnixRuntimeEngine *engine)
{
  fputs("{", stdout);
  print_string_field("engine_id", engine->id);
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
  fputs(",", stdout);
  print_string_field("summary", engine->summary);
  fputs("}", stdout);
}

static int
print_applications(void)
{
  fputc('[', stdout);
  for (size_t index = 0; index < xnix_runtime_application_count(); index++) {
    const XnixRuntimeApplication *application = xnix_runtime_application_at(index);

    if (index > 0) {
      fputs(",", stdout);
    }
    print_application(application);
  }
  fputs("]\n", stdout);
  return 0;
}

static int
print_engine_catalog(void)
{
  fputs("{", stdout);
  printf("\"version\":\"%s\",", xnix_runtime_version());
  fputs("\"catalog_type\":\"compatibility-engine\",", stdout);
  fputs("\"runtime_policy_owner\":true,", stdout);
  fputs("\"desktop_shell_policy_owner\":false,", stdout);
  fputs("\"catalog_owner\":\"c\",", stdout);
  fputs("\"backend_details_exposed\":false,", stdout);
  fputs("\"engine_count\":", stdout);
  printf("%zu,", xnix_runtime_engine_count());
  fputs("\"engines\":[", stdout);
  for (size_t index = 0; index < xnix_runtime_engine_count(); index++) {
    const XnixRuntimeEngine *engine = xnix_runtime_engine_at(index);

    if (index > 0) {
      fputs(",", stdout);
    }
    print_engine_record(engine);
  }
  fputs("]}\n", stdout);
  return 0;
}

static int
print_portal_policy_catalog(void)
{
  fputs("{", stdout);
  printf("\"version\":\"%s\",", xnix_runtime_version());
  print_string_field("catalog_type", "portal-access-policy");
  fputs(",", stdout);
  fputs("\"runtime_policy_owner\":true,", stdout);
  fputs("\"desktop_shell_policy_owner\":false,", stdout);
  print_string_field("catalog_owner", "c");
  fputs(",", stdout);
  fputs("\"backend_details_exposed\":false,", stdout);
  fputs("\"policy_count\":", stdout);
  printf("%zu,", xnix_runtime_portal_policy_count());
  fputs("\"policies\":[", stdout);
  for (size_t index = 0; index < xnix_runtime_portal_policy_count(); index++) {
    const XnixRuntimePortalPolicy *policy = xnix_runtime_portal_policy_at(index);

    if (index > 0) {
      fputs(",", stdout);
    }
    print_portal_policy_record(policy);
  }
  fputs("]}\n", stdout);
  return 0;
}

static int
print_portal_policy(const char *application_id, const char *operation)
{
  const XnixRuntimeApplication *application = xnix_runtime_find_application(application_id);
  const XnixRuntimePortalPolicy *policy = xnix_runtime_find_portal_policy(operation);

  if (application == NULL) {
    fprintf(stderr, "xnix-runtime-core: application is not registered in the C Runtime catalog\n");
    return 64;
  }
  if (policy == NULL) {
    fprintf(stderr, "xnix-runtime-core: operation must be a registered sensitive desktop operation\n");
    return 64;
  }

  print_portal_policy_for_application(application->id, policy);
  fputs("\n", stdout);
  return 0;
}

static int
print_engine_for_mode(const char *mode)
{
  const XnixRuntimeEngine *engine = xnix_runtime_select_engine_for_mode(mode);

  if (engine == NULL) {
    fprintf(stderr, "xnix-runtime-core: mode must be one of automatic, wine, vm\n");
    return 64;
  }

  print_selected_engine(engine);
  fputs("\n", stdout);
  return 0;
}

static int
print_application_by_id(const char *application_id)
{
  const XnixRuntimeApplication *application = xnix_runtime_find_application(application_id);

  if (application == NULL) {
    fprintf(stderr, "xnix-runtime-core: application is not registered in the C Runtime catalog\n");
    return 64;
  }

  print_application(application);
  fputs("\n", stdout);
  return 0;
}

static int
print_probe(void)
{
  fputs("{", stdout);
  printf("\"version\":\"%s\",", xnix_runtime_version());
  printf("\"core_language\":\"%s\",", xnix_runtime_core_language());
  fputs("\"business_logic_runtime\":\"c\",", stdout);
  fputs("\"ruby_role\":\"tests-and-development-tools\",", stdout);
  printf("\"bus_name\":\"%s\",", xnix_runtime_bus_name());
  printf("\"object_path\":\"%s\",", xnix_runtime_object_path());
  printf("\"interface\":\"%s\",", xnix_runtime_interface());
  fputs("\"runtime_owned\":true,", stdout);
  fputs("\"kde_policy_owner\":false,", stdout);
  fputs("\"application_catalog_owner\":\"c\",", stdout);
  fputs("\"application_count\":", stdout);
  printf("%zu,", xnix_runtime_application_count());
  fputs("\"engine_catalog_owner\":\"c\",", stdout);
  fputs("\"engine_count\":", stdout);
  printf("%zu,", xnix_runtime_engine_count());
  fputs("\"portal_policy_owner\":\"c\",", stdout);
  fputs("\"portal_policy_count\":", stdout);
  printf("%zu,", xnix_runtime_portal_policy_count());
  fputs("\"write_methods\":", stdout);
  print_write_methods();
  fputs(",", stdout);
  fputs("\"write_method_count\":", stdout);
  printf("%zu,", xnix_runtime_write_method_count());
  fputs("\"write_methods_supported\":false,", stdout);
  fputs("\"write_method_dispatch_enabled\":false,", stdout);
  fputs("\"host_root_modified\":false,", stdout);
  fputs("\"backend_details_exposed\":false", stdout);
  fputs("}\n", stdout);
  return 0;
}

static int
print_write_gate(const char *method_name)
{
  XnixRuntimeWriteGate gate;

  if (!xnix_runtime_write_gate(method_name, &gate)) {
    fprintf(stderr, "xnix-runtime-core: method must be a reserved Runtime write method\n");
    return 64;
  }

  fputs("{", stdout);
  printf("\"version\":\"%s\",", xnix_runtime_version());
  fputs("\"gate_type\":\"runtime-write-gate\",", stdout);
  printf("\"method_name\":\"%s\",", gate.method_name);
  printf("\"gate_decision\":\"%s\",", gate.decision);
  fputs("\"write_method_enabled\":", stdout);
  print_bool(gate.write_method_enabled);
  fputs(",", stdout);
  fputs("\"dispatch_enabled\":", stdout);
  print_bool(gate.dispatch_enabled);
  fputs(",", stdout);
  fputs("\"request_object_created\":", stdout);
  print_bool(gate.request_object_created);
  fputs(",", stdout);
  fputs("\"execution_started\":", stdout);
  print_bool(gate.execution_started);
  fputs(",", stdout);
  printf("\"denial_error_name\":\"%s\",", xnix_runtime_write_error());
  fputs("\"host_root_modified\":", stdout);
  print_bool(gate.host_root_modified);
  fputs(",", stdout);
  fputs("\"backend_details_exposed\":", stdout);
  print_bool(gate.backend_details_exposed);
  fputs("}\n", stdout);
  return 0;
}

static int
usage(void)
{
  fputs("Usage: xnix-runtime-core {probe|list-applications|get-application APP_ID|list-engines|select-engine MODE|list-portal-policies|portal-policy APP_ID OPERATION|write-gate METHOD}\n", stderr);
  return 64;
}

int
main(int argc, char **argv)
{
  if (argc < 2) {
    return usage();
  }

  if (strcmp(argv[1], "probe") == 0 && argc == 2) {
    return print_probe();
  }

  if (strcmp(argv[1], "list-applications") == 0 && argc == 2) {
    return print_applications();
  }

  if (strcmp(argv[1], "get-application") == 0 && argc == 3) {
    return print_application_by_id(argv[2]);
  }

  if (strcmp(argv[1], "list-engines") == 0 && argc == 2) {
    return print_engine_catalog();
  }

  if (strcmp(argv[1], "select-engine") == 0 && argc == 3) {
    return print_engine_for_mode(argv[2]);
  }

  if (strcmp(argv[1], "list-portal-policies") == 0 && argc == 2) {
    return print_portal_policy_catalog();
  }

  if (strcmp(argv[1], "portal-policy") == 0 && argc == 4) {
    return print_portal_policy(argv[2], argv[3]);
  }

  if (strcmp(argv[1], "write-gate") == 0 && argc == 3) {
    return print_write_gate(argv[2]);
  }

  return usage();
}
