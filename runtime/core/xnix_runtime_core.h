#ifndef XNIX_RUNTIME_CORE_H
#define XNIX_RUNTIME_CORE_H

#include <stdbool.h>
#include <stddef.h>

#define XNIX_RUNTIME_VERSION "0.2.56"
#define XNIX_RUNTIME_BUS_NAME "org.xnix.Compatibility1"
#define XNIX_RUNTIME_OBJECT_PATH "/org/xnix/Compatibility1"
#define XNIX_RUNTIME_INTERFACE "org.xnix.Compatibility1"
#define XNIX_RUNTIME_WRITE_ERROR "org.xnix.Compatibility1.Error.WriteMethodDisabled"

typedef struct {
  const char *method_name;
  const char *decision;
  bool write_method_enabled;
  bool dispatch_enabled;
  bool request_object_created;
  bool execution_started;
  bool host_root_modified;
  bool backend_details_exposed;
} XnixRuntimeWriteGate;

const char *xnix_runtime_version(void);
const char *xnix_runtime_bus_name(void);
const char *xnix_runtime_object_path(void);
const char *xnix_runtime_interface(void);
const char *xnix_runtime_write_error(void);
const char *xnix_runtime_core_language(void);

bool xnix_runtime_is_write_method(const char *method_name);
bool xnix_runtime_write_gate(const char *method_name, XnixRuntimeWriteGate *gate);
size_t xnix_runtime_write_method_count(void);
const char *xnix_runtime_write_method_at(size_t index);

#endif
