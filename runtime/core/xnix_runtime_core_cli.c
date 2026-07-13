#include "xnix_runtime_core.h"

#include <stdio.h>
#include <string.h>

static void
print_bool(bool value)
{
  fputs(value ? "true" : "false", stdout);
}

static void
print_write_methods(void)
{
  fputc('[', stdout);
  for (size_t index = 0; index < xnix_runtime_write_method_count(); index++) {
    if (index > 0) {
      fputs(",", stdout);
    }
    printf("\"%s\"", xnix_runtime_write_method_at(index));
  }
  fputc(']', stdout);
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
  fputs("Usage: xnix-runtime-core {probe|write-gate METHOD}\n", stderr);
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

  if (strcmp(argv[1], "write-gate") == 0 && argc == 3) {
    return print_write_gate(argv[2]);
  }

  return usage();
}
