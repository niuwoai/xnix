#include "xnix_runtime_core.h"

#include <string.h>

static const char *const write_methods[] = {
  "InstallRecipe",
  "Launch",
  "CreateSnapshot",
  "RestoreSnapshot",
};

static const XnixRuntimeApplication applications[] = {
  {
    .id = "org.xnix.sample.notepad",
    .name = "Sample Notepad",
    .icon = "accessories-text-editor",
    .runtime_mode = "automatic",
    .desktop_category = "Utility",
    .launcher_command = "xnix-compat-launch --app org.xnix.sample.notepad %U",
    .primary_extension = ".txt",
    .primary_mime_type = "application/x-xnix-txt",
    .runtime_owned = true,
    .kde_policy_owner = false,
    .backend_details_exposed = false,
  },
};

const char *
xnix_runtime_version(void)
{
  return XNIX_RUNTIME_VERSION;
}

const char *
xnix_runtime_bus_name(void)
{
  return XNIX_RUNTIME_BUS_NAME;
}

const char *
xnix_runtime_object_path(void)
{
  return XNIX_RUNTIME_OBJECT_PATH;
}

const char *
xnix_runtime_interface(void)
{
  return XNIX_RUNTIME_INTERFACE;
}

const char *
xnix_runtime_write_error(void)
{
  return XNIX_RUNTIME_WRITE_ERROR;
}

const char *
xnix_runtime_core_language(void)
{
  return "c";
}

bool
xnix_runtime_is_write_method(const char *method_name)
{
  if (method_name == NULL) {
    return false;
  }

  for (size_t index = 0; index < xnix_runtime_write_method_count(); index++) {
    if (strcmp(method_name, write_methods[index]) == 0) {
      return true;
    }
  }

  return false;
}

bool
xnix_runtime_write_gate(const char *method_name, XnixRuntimeWriteGate *gate)
{
  if (gate == NULL || !xnix_runtime_is_write_method(method_name)) {
    return false;
  }

  gate->method_name = method_name;
  gate->decision = "blocked-until-production-backend";
  gate->write_method_enabled = false;
  gate->dispatch_enabled = false;
  gate->request_object_created = false;
  gate->execution_started = false;
  gate->host_root_modified = false;
  gate->backend_details_exposed = false;
  return true;
}

size_t
xnix_runtime_write_method_count(void)
{
  return sizeof(write_methods) / sizeof(write_methods[0]);
}

const char *
xnix_runtime_write_method_at(size_t index)
{
  if (index >= xnix_runtime_write_method_count()) {
    return NULL;
  }

  return write_methods[index];
}

size_t
xnix_runtime_application_count(void)
{
  return sizeof(applications) / sizeof(applications[0]);
}

const XnixRuntimeApplication *
xnix_runtime_application_at(size_t index)
{
  if (index >= xnix_runtime_application_count()) {
    return NULL;
  }

  return &applications[index];
}

const XnixRuntimeApplication *
xnix_runtime_find_application(const char *application_id)
{
  if (application_id == NULL) {
    return NULL;
  }

  for (size_t index = 0; index < xnix_runtime_application_count(); index++) {
    const XnixRuntimeApplication *application = xnix_runtime_application_at(index);

    if (application != NULL && strcmp(application_id, application->id) == 0) {
      return application;
    }
  }

  return NULL;
}
