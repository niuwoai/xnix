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

static const XnixRuntimeEngine engines[] = {
  {
    .id = "automatic-managed",
    .label = "Automatic",
    .kind = "orchestrator",
    .summary = "Xnix will choose the best available compatibility path.",
    .ready = false,
    .launch_enabled = false,
    .user_visible = true,
    .backend_details_exposed = false,
  },
  {
    .id = "local-compatibility-engine",
    .label = "Local compatibility engine",
    .kind = "local",
    .summary = "Local compatibility engine binding is pending.",
    .ready = false,
    .launch_enabled = false,
    .user_visible = true,
    .backend_details_exposed = false,
  },
  {
    .id = "isolated-compatibility-engine",
    .label = "Isolated compatibility engine",
    .kind = "isolated",
    .summary = "Isolated compatibility engine binding is pending.",
    .ready = false,
    .launch_enabled = false,
    .user_visible = true,
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

size_t
xnix_runtime_engine_count(void)
{
  return sizeof(engines) / sizeof(engines[0]);
}

const XnixRuntimeEngine *
xnix_runtime_engine_at(size_t index)
{
  if (index >= xnix_runtime_engine_count()) {
    return NULL;
  }

  return &engines[index];
}

const XnixRuntimeEngine *
xnix_runtime_find_engine(const char *engine_id)
{
  if (engine_id == NULL) {
    return NULL;
  }

  for (size_t index = 0; index < xnix_runtime_engine_count(); index++) {
    const XnixRuntimeEngine *engine = xnix_runtime_engine_at(index);

    if (engine != NULL && strcmp(engine_id, engine->id) == 0) {
      return engine;
    }
  }

  return NULL;
}

const XnixRuntimeEngine *
xnix_runtime_select_engine_for_mode(const char *mode)
{
  if (mode == NULL) {
    return NULL;
  }

  if (strcmp(mode, "automatic") == 0) {
    return xnix_runtime_find_engine("automatic-managed");
  }
  if (strcmp(mode, "wine") == 0) {
    return xnix_runtime_find_engine("local-compatibility-engine");
  }
  if (strcmp(mode, "vm") == 0) {
    return xnix_runtime_find_engine("isolated-compatibility-engine");
  }

  return NULL;
}
