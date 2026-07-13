#ifndef XNIX_RUNTIME_CORE_H
#define XNIX_RUNTIME_CORE_H

#include <stdbool.h>
#include <stddef.h>

#define XNIX_RUNTIME_VERSION "0.2.79"
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

typedef struct {
  const char *id;
  const char *status;
  const char *summary;
} XnixRuntimeInstallPhase;

typedef struct {
  const char *application_id;
  const char *install_state;
  const char *selected_strategy;
  const char *blocked_actions[6];
  size_t blocked_action_count;
  XnixRuntimeInstallPhase phases[6];
  size_t phase_count;
  bool artifact_manifest_ready;
  bool artifact_signature_verified;
  bool acquisition_ready;
  bool package_source_ready;
  bool state_root_allocated;
  bool development_recipe_install_allowed;
  bool production_recipe_install_allowed;
  bool install_ready;
  bool desktop_activation_ready;
  bool download_enabled;
  bool install_enabled;
  bool network_request_created;
  bool artifacts_downloaded;
  bool host_root_modified;
  bool privileged_container_required;
  bool desktop_shell_command_exposed;
  bool runtime_owned;
  bool kde_policy_owner;
  bool backend_details_exposed;
  const char *summary;
} XnixRuntimeInstallReadinessPolicy;

typedef struct {
  const char *id;
  const char *kind;
  const char *cache_namespace;
  const char *summary;
  bool required;
  bool resolved;
  bool downloaded;
} XnixRuntimeArtifactGroup;

typedef struct {
  const char *id;
  const char *status;
  const char *summary;
} XnixRuntimeArtifactPreflight;

typedef struct {
  const char *application_id;
  const char *manifest_state;
  const char *selected_strategy;
  XnixRuntimeArtifactGroup artifact_groups[3];
  size_t artifact_group_count;
  XnixRuntimeArtifactPreflight required_preflight[5];
  size_t required_preflight_count;
  const char *blocked_actions[4];
  size_t blocked_action_count;
  bool runtime_owned;
  bool kde_policy_owner;
  bool manifest_ready;
  bool signature_verified;
  bool acquisition_preflight_ready;
  bool download_enabled;
  bool install_enabled;
  bool network_request_created;
  bool artifacts_downloaded;
  bool host_root_modified;
  bool privileged_container_required;
  bool desktop_shell_command_exposed;
  bool backend_details_exposed;
  const char *summary;
} XnixRuntimeArtifactManifestPolicy;

typedef struct {
  const char *id;
  const char *status;
  const char *summary;
} XnixRuntimeAcquisitionCheck;

typedef struct {
  const char *application_id;
  const char *preflight_state;
  const char *selected_strategy;
  XnixRuntimeAcquisitionCheck checks[5];
  size_t check_count;
  const char *blocked_actions[4];
  size_t blocked_action_count;
  bool runtime_owned;
  bool kde_policy_owner;
  bool acquisition_ready;
  bool download_enabled;
  bool install_enabled;
  bool network_required_for_planning;
  bool network_request_created;
  bool artifacts_downloaded;
  bool host_root_modified;
  bool privileged_container_required;
  bool desktop_shell_command_exposed;
  bool package_source_ready;
  bool backend_details_exposed;
  const char *summary;
} XnixRuntimeAcquisitionPreflightPolicy;

typedef struct {
  const char *id;
  const char *kind;
  const char *selection_state;
  const char *supported_strategies[3];
  size_t supported_strategy_count;
  bool runtime_owned;
  const char *summary;
} XnixRuntimePackageSourceChannel;

typedef struct {
  const char *id;
  const char *status;
  const char *summary;
} XnixRuntimePackageSourcePreflight;

typedef struct {
  const char *application_id;
  const char *source_selection_state;
  const char *selected_strategy;
  XnixRuntimePackageSourceChannel source_channels[3];
  size_t source_channel_count;
  XnixRuntimePackageSourcePreflight required_preflight[4];
  size_t required_preflight_count;
  const char *blocked_actions[4];
  size_t blocked_action_count;
  bool signed_source_required;
  bool runtime_cache_required;
  bool direct_desktop_install_allowed;
  bool user_visible_backend_names;
  bool host_package_manager_invoked;
  bool runtime_owned;
  bool kde_policy_owner;
  bool package_source_ready;
  bool install_enabled;
  bool network_required_for_planning;
  bool host_root_modified;
  bool privileged_container_required;
  bool desktop_shell_command_exposed;
  bool backend_details_exposed;
  const char *summary;
} XnixRuntimePackageSourcePolicy;

typedef struct {
  const char *id;
  const char *status;
  const char *summary;
} XnixRuntimeBackendBindingPreflight;

typedef struct {
  const char *application_id;
  const char *selected_strategy;
  XnixRuntimeBackendBindingPreflight required_preflight[4];
  size_t required_preflight_count;
  const char *blocked_actions[4];
  size_t blocked_action_count;
  bool runtime_owned;
  bool kde_policy_owner;
  bool managed_binding_ready;
  bool launch_enabled;
  bool execution_request_created;
  bool host_root_modified;
  bool network_required;
  bool privileged_container_required;
  bool backend_details_exposed;
  const char *summary;
} XnixRuntimeBackendBindingPolicy;

typedef struct {
  const char *id;
  const char *label;
  const char *value;
  const char *options[3];
  size_t option_count;
} XnixRuntimeSettingsField;

typedef struct {
  const char *id;
  const char *title;
  const char *description;
  XnixRuntimeSettingsField fields[2];
  size_t field_count;
} XnixRuntimeSettingsSection;

typedef struct {
  const char *application_id;
  const char *settings_state;
  XnixRuntimeSettingsSection sections[5];
  size_t section_count;
  bool runtime_owned;
  bool kde_policy_owner;
  bool settings_persisted;
  bool host_root_modified;
  bool backend_details_exposed;
  const char *summary;
} XnixRuntimeSettingsPolicy;

typedef struct {
  const char *id;
  const char *status;
  const char *summary;
} XnixRuntimeSettingsChangeStep;

typedef struct {
  const char *section;
  const char *field;
  const char *value;
  const char *options[3];
  size_t option_count;
} XnixRuntimeAffectedSettingsPolicy;

typedef struct {
  const char *application_id;
  const char *section_id;
  const char *field_id;
  const char *requested_value;
  const char *change_state;
  XnixRuntimeAffectedSettingsPolicy affected_policy;
  XnixRuntimeSettingsChangeStep steps[5];
  size_t step_count;
  const char *blocked_actions[4];
  size_t blocked_action_count;
  bool runtime_owned;
  bool kde_policy_owner;
  bool apply_enabled;
  bool settings_persisted;
  bool host_root_modified;
  bool backend_details_exposed;
  bool user_confirmation_required;
  bool snapshot_recommended;
  bool portal_policy_review_required;
  bool runtime_restart_required;
  const char *summary;
} XnixRuntimeSettingsChangePolicy;

typedef struct {
  const char *dbus_service_file;
  const char *systemd_unit;
  const char *libexec_wrapper;
  const char *dbus_contract;
  const char *packaged_wrapper;
} XnixRuntimeServiceActivation;

typedef struct {
  const char *id;
  const char *status;
  const char *summary;
} XnixRuntimeServiceBindingCheck;

typedef struct {
  size_t total;
  size_t passed;
  size_t pending;
  size_t blocked;
} XnixRuntimeServiceBindingCounts;

typedef struct {
  XnixRuntimeServiceActivation activation;
  XnixRuntimeServiceBindingCheck checks[5];
  size_t check_count;
  XnixRuntimeServiceBindingCounts counts;
  bool runtime_owned;
  bool kde_policy_owner;
  bool activation_binding_ready;
  bool live_dbus_owner_ready;
  bool smoke_adapter_available;
  bool network_required;
  bool host_root_modified;
  bool privileged_container_required;
  bool backend_details_exposed;
  const char *summary;
} XnixRuntimeServiceBindingPolicy;

typedef struct {
  const char *id;
  const char *status;
  const char *summary;
} XnixRuntimeLiveOwnerRequiredGate;

typedef struct {
  XnixRuntimeLiveOwnerRequiredGate required_gates[5];
  size_t required_gate_count;
  const char *blocked_reasons[5];
  size_t blocked_reason_count;
  bool runtime_owned;
  bool kde_policy_owner;
  bool activation_binding_ready;
  bool live_dbus_owner_ready;
  bool production_owner_enabled;
  bool owner_transition_ready;
  bool smoke_adapter_available;
  bool smoke_adapter_is_production_owner;
  bool kde_may_claim_runtime_ownership;
  bool network_required;
  bool host_root_modified;
  bool privileged_container_required;
  bool backend_details_exposed;
  const char *summary;
} XnixRuntimeLiveOwnerGatePolicy;

typedef struct {
  const char *id;
  const char *status;
  const char *summary;
} XnixRuntimeOwnerSmokeStep;

typedef struct {
  size_t total;
  size_t passed;
  size_t pending;
  size_t blocked;
} XnixRuntimeOwnerSmokeCounts;

typedef struct {
  XnixRuntimeOwnerSmokeStep steps[7];
  size_t step_count;
  XnixRuntimeOwnerSmokeCounts counts;
  const char *blocked_actions[5];
  size_t blocked_action_count;
  bool runtime_owned;
  bool kde_policy_owner;
  bool activation_binding_ready;
  bool live_dbus_owner_ready;
  bool production_owner_enabled;
  bool owner_transition_ready;
  const char *smoke_state;
  const char *smoke_environment;
  bool network_required;
  bool host_root_modified;
  bool privileged_container_required;
  bool system_service_started;
  bool production_bus_claimed;
  bool backend_details_exposed;
  const char *summary;
} XnixRuntimeOwnerSmokePlanPolicy;

typedef struct {
  const char *id;
  const char *status;
  size_t method_count;
  const char *missing_methods[1];
  size_t missing_method_count;
  const char *summary;
} XnixRuntimeMethodParityCheck;

typedef struct {
  size_t total;
  size_t passed;
  size_t blocked;
  size_t pending;
} XnixRuntimeMethodParityCounts;

typedef struct {
  const char *read_only_methods[29];
  size_t read_only_method_count;
  XnixRuntimeMethodParityCheck parity_checks[5];
  size_t parity_check_count;
  XnixRuntimeMethodParityCounts counts;
  const char *write_methods[4];
  size_t write_method_count;
  bool runtime_owned;
  bool kde_policy_owner;
  bool read_only_method_parity_ready;
  bool write_methods_supported;
  bool write_method_dispatch_enabled;
  bool network_required;
  bool host_root_modified;
  bool privileged_container_required;
  bool backend_details_exposed;
  const char *summary;
} XnixRuntimeMethodParityManifestPolicy;

typedef struct {
  const char *id;
  const char *status;
  const char *message;
} XnixRuntimeRecipeTrustCheck;

typedef struct {
  const char *decision;
  size_t recipe_count;
  XnixRuntimeRecipeTrustCheck checks[3];
  size_t check_count;
  const char *blocking_reasons[3];
  size_t blocking_reason_count;
  const char *next_requirements[3];
  size_t next_requirement_count;
  bool runtime_owned;
  bool kde_policy_owner;
  bool digest_verified;
  bool signed_recipe_validation;
  bool development_registry;
  bool production_trusted;
  bool development_staging_allowed;
  bool production_install_allowed;
  bool host_root_modified;
  bool backend_details_exposed;
  const char *summary;
} XnixRuntimeRecipeTrustPolicy;

typedef struct {
  const char *id;
  const char *signature_status;
  bool digest_verified;
} XnixRuntimeRecipeInstallMatch;

typedef struct {
  const char *application_id;
  const char *mode;
  const char *decision;
  const char *policy_decision;
  XnixRuntimeRecipeInstallMatch matched_recipe;
  bool matched_recipe_present;
  const char *blocking_reasons[4];
  size_t blocking_reason_count;
  const char *requirements[3];
  size_t requirement_count;
  bool runtime_owned;
  bool kde_policy_owner;
  bool development_staging_allowed;
  bool production_install_allowed;
  bool request_object_created;
  bool install_started;
  bool host_root_modified;
  bool backend_details_exposed;
  const char *summary;
} XnixRuntimeRecipeInstallGate;

typedef struct {
  const XnixRuntimeApplication *application;
  const XnixRuntimeInstallReadinessPolicy *install_readiness;
  XnixRuntimeRecipeInstallGate recipe_install_gate;
  const char *environment;
  bool runtime_owned;
  bool kde_policy_owner;
} XnixRuntimeCompatibilityInstallPlan;

typedef struct {
  const char *id;
  const char *source_type;
  const char *status;
  const char *priority;
  const char *title;
  const char *summary;
  const char *runtime_gate;
  const char *next_step;
  bool user_review_required;
  bool execution_enabled;
  bool backend_details_exposed;
} XnixRuntimeCompatibilityAction;

typedef struct {
  const XnixRuntimeApplication *application;
  XnixRuntimeCompatibilityAction actions[5];
  size_t action_count;
  size_t pending_action_count;
  size_t user_review_required_count;
  const char *blocked_actions[6];
  size_t blocked_action_count;
  bool runtime_owned;
  bool kde_policy_owner;
  bool execution_enabled;
  bool repair_execution_enabled;
  bool settings_persistence_enabled;
  bool host_root_modified;
  bool network_required;
  bool backend_details_exposed;
  const char *summary;
} XnixRuntimeCompatibilityActionQueue;

typedef struct {
  char receipt_id[192];
  const XnixRuntimeApplication *application;
  XnixRuntimeCompatibilityAction action;
  const char *decision;
  const char *next_step;
  const char *blocked_actions[6];
  size_t blocked_action_count;
  bool decision_recorded;
  bool runtime_owned;
  bool kde_policy_owner;
  bool execution_enabled;
  bool repair_execution_enabled;
  bool settings_persistence_enabled;
  bool resource_grant_created;
  bool host_root_modified;
  bool network_required;
  bool backend_details_exposed;
  const char *summary;
} XnixRuntimeCompatibilityActionReviewReceipt;

typedef struct {
  const XnixRuntimeApplication *application;
  const XnixRuntimeEngine *engine;
  bool runtime_owned;
  bool kde_policy_owner;
  bool portal_policy_required;
  bool snapshot_before_risky_change;
  bool diagnostics_required;
  bool backend_binding_ready;
  bool launch_enabled;
  bool execution_request_created;
  bool host_root_modified;
  bool network_required;
  bool backend_details_exposed;
  const char *summary;
} XnixRuntimeCompatibilityRunPlan;

typedef struct {
  const char *entry_point;
  const char *kind;
  const char *relative_path;
  const char *desktop_file;
  const char *command;
  bool user_visible;
  bool portal_required;
  bool activation_ready;
  bool runtime_owned;
  bool kde_policy_owner;
  bool host_root_modified;
  bool backend_details_exposed;
  const char *summary;
} XnixRuntimeDesktopActivationArtifact;

typedef struct {
  const XnixRuntimeApplication *application;
  XnixRuntimeDesktopActivationArtifact artifacts[7];
  size_t artifact_count;
  bool runtime_owned;
  bool kde_policy_owner;
  bool official_desktop_only;
  bool stable_desktop_contract;
  bool activation_ready;
  bool files_written;
  bool host_root_modified;
  bool backend_details_exposed;
  const char *desktop;
  const char *summary;
} XnixRuntimeDesktopActivationManifest;

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

size_t xnix_runtime_install_readiness_policy_count(void);
const XnixRuntimeInstallReadinessPolicy *xnix_runtime_install_readiness_policy_at(size_t index);
const XnixRuntimeInstallReadinessPolicy *xnix_runtime_find_install_readiness_policy(const char *application_id);
bool xnix_runtime_install_readiness_allows_recipe_install(
  const XnixRuntimeInstallReadinessPolicy *policy,
  const char *environment
);

size_t xnix_runtime_artifact_manifest_policy_count(void);
const XnixRuntimeArtifactManifestPolicy *xnix_runtime_artifact_manifest_policy_at(size_t index);
const XnixRuntimeArtifactManifestPolicy *xnix_runtime_find_artifact_manifest_policy(const char *application_id);

size_t xnix_runtime_acquisition_preflight_policy_count(void);
const XnixRuntimeAcquisitionPreflightPolicy *xnix_runtime_acquisition_preflight_policy_at(size_t index);
const XnixRuntimeAcquisitionPreflightPolicy *xnix_runtime_find_acquisition_preflight_policy(const char *application_id);

size_t xnix_runtime_package_source_policy_count(void);
const XnixRuntimePackageSourcePolicy *xnix_runtime_package_source_policy_at(size_t index);
const XnixRuntimePackageSourcePolicy *xnix_runtime_find_package_source_policy(const char *application_id);

size_t xnix_runtime_backend_binding_policy_count(void);
const XnixRuntimeBackendBindingPolicy *xnix_runtime_backend_binding_policy_at(size_t index);
const XnixRuntimeBackendBindingPolicy *xnix_runtime_find_backend_binding_policy(const char *application_id);

size_t xnix_runtime_settings_policy_count(void);
const XnixRuntimeSettingsPolicy *xnix_runtime_settings_policy_at(size_t index);
const XnixRuntimeSettingsPolicy *xnix_runtime_find_settings_policy(const char *application_id);

size_t xnix_runtime_settings_change_policy_count(void);
const XnixRuntimeSettingsChangePolicy *xnix_runtime_settings_change_policy_at(size_t index);
const XnixRuntimeSettingsChangePolicy *xnix_runtime_find_settings_change_policy(const char *application_id);

const XnixRuntimeServiceBindingPolicy *xnix_runtime_service_binding_policy(void);
const XnixRuntimeLiveOwnerGatePolicy *xnix_runtime_live_owner_gate_policy(void);
const XnixRuntimeOwnerSmokePlanPolicy *xnix_runtime_owner_smoke_plan_policy(void);
const XnixRuntimeMethodParityManifestPolicy *xnix_runtime_method_parity_manifest_policy(void);
const XnixRuntimeRecipeTrustPolicy *xnix_runtime_recipe_trust_policy(void);
bool xnix_runtime_recipe_install_gate(
  const char *application_id,
  const char *mode,
  XnixRuntimeRecipeInstallGate *gate
);
bool xnix_runtime_compatibility_install_plan(
  const char *application_id,
  const char *environment,
  XnixRuntimeCompatibilityInstallPlan *plan
);
bool xnix_runtime_compatibility_action_queue(
  const char *application_id,
  XnixRuntimeCompatibilityActionQueue *queue
);
bool xnix_runtime_compatibility_action_review_receipt(
  const char *application_id,
  const char *action_id,
  const char *decision,
  XnixRuntimeCompatibilityActionReviewReceipt *receipt
);
bool xnix_runtime_compatibility_run_plan(
  const char *application_id,
  XnixRuntimeCompatibilityRunPlan *plan
);
bool xnix_runtime_desktop_activation_manifest(
  const char *application_id,
  XnixRuntimeDesktopActivationManifest *manifest
);

#endif
