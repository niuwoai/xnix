#ifndef XNIX_RUNTIME_CORE_H
#define XNIX_RUNTIME_CORE_H

#include <stdbool.h>
#include <stddef.h>

#define XNIX_RUNTIME_VERSION "0.2.61"
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

typedef struct {
  const char *id;
  const char *name;
  const char *icon;
  const char *runtime_mode;
  const char *desktop_category;
  const char *launcher_command;
  const char *primary_extension;
  const char *primary_mime_type;
  bool runtime_owned;
  bool kde_policy_owner;
  bool backend_details_exposed;
} XnixRuntimeApplication;

typedef struct {
  const char *id;
  const char *label;
  const char *kind;
  const char *summary;
  bool ready;
  bool launch_enabled;
  bool user_visible;
  bool backend_details_exposed;
} XnixRuntimeEngine;

typedef struct {
  const char *operation;
  const char *portal_interface;
  const char *decision;
  const char *summary;
  const char *resources[3];
  size_t resource_count;
  bool portal_required;
  bool user_mediation_required;
  bool direct_access_allowed;
  bool runtime_policy_owner;
  bool desktop_shell_policy_owner;
  bool backend_details_exposed;
} XnixRuntimePortalPolicy;

typedef struct {
  const char *reason;
  const char *summary;
  bool enabled_by_default;
  bool application_state;
  bool runtime_metadata;
  bool desktop_activation_receipts;
  bool user_documents;
  bool host_system;
  bool restore_available;
  bool restore_requires_user_confirmation;
  bool restore_preserve_user_documents;
  const char *retention_policy;
  size_t keep_latest;
  bool prune_automatically;
  bool runtime_policy_owner;
  bool desktop_shell_policy_owner;
  bool backend_details_exposed;
  bool host_root_modified;
} XnixRuntimeSnapshotPolicy;

typedef struct {
  const char *application_id;
  const char *state_namespace;
  const char *storage_scope;
  const char *allocation_state;
  const char *managed_scopes[4];
  const char *managed_scope_summaries[4];
  bool managed_scope_snapshot_included[4];
  size_t managed_scope_count;
  const char *blocked_actions[4];
  size_t blocked_action_count;
  const char *retention_policy;
  size_t automatic_restore_points;
  size_t manual_restore_points;
  bool runtime_owned;
  bool kde_policy_owner;
  bool directories_created;
  bool host_root_modified;
  bool user_documents_included;
  bool portal_required_for_user_files;
  bool snapshot_eligible;
  bool restore_requires_confirmation;
  bool user_documents_excluded;
  bool backend_details_exposed;
  const char *summary;
} XnixRuntimeStateRootPolicy;

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

size_t xnix_runtime_application_count(void);
const XnixRuntimeApplication *xnix_runtime_application_at(size_t index);
const XnixRuntimeApplication *xnix_runtime_find_application(const char *application_id);

size_t xnix_runtime_engine_count(void);
const XnixRuntimeEngine *xnix_runtime_engine_at(size_t index);
const XnixRuntimeEngine *xnix_runtime_find_engine(const char *engine_id);
const XnixRuntimeEngine *xnix_runtime_select_engine_for_mode(const char *mode);

size_t xnix_runtime_portal_policy_count(void);
const XnixRuntimePortalPolicy *xnix_runtime_portal_policy_at(size_t index);
const XnixRuntimePortalPolicy *xnix_runtime_find_portal_policy(const char *operation);

size_t xnix_runtime_snapshot_policy_count(void);
const XnixRuntimeSnapshotPolicy *xnix_runtime_snapshot_policy_at(size_t index);
const XnixRuntimeSnapshotPolicy *xnix_runtime_find_snapshot_policy(const char *reason);

size_t xnix_runtime_state_root_policy_count(void);
const XnixRuntimeStateRootPolicy *xnix_runtime_state_root_policy_at(size_t index);
const XnixRuntimeStateRootPolicy *xnix_runtime_find_state_root_policy(const char *application_id);

#endif
