#ifndef XNIX_RUNTIME_CORE_H
#define XNIX_RUNTIME_CORE_H

#include <stdbool.h>
#include <stddef.h>

#define XNIX_RUNTIME_VERSION "0.2.514"
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
  char handle_token[160];
  const XnixRuntimeApplication *application;
  const XnixRuntimePortalPolicy *policy;
  const char *operation;
  const char *reason;
  const char *portal_destination;
  const char *portal_method;
  const char *portal_object_path;
  const char *completion_signal;
  const char *response_field;
  const char *result_owner;
  bool request_allowed;
  bool request_object_required;
  bool request_object_created;
  bool user_mediation_required;
  bool direct_access_allowed;
  bool portal_required;
  bool permission_granted;
  bool runtime_owned;
  bool kde_policy_owner;
  bool host_permission_changed;
  bool host_root_modified;
  bool backend_details_exposed;
  const char *summary;
} XnixRuntimePortalRequestPlan;

typedef struct {
  const char *id;
  const char *label;
  const char *operation;
  const char *decision;
  const char *portal_interface;
  bool portal_required;
  bool user_mediation_required;
  bool change_pending;
  bool request_object_created;
  bool permission_granted;
  bool direct_access_allowed;
  bool backend_details_exposed;
} XnixRuntimeCompatibilityPermission;

typedef struct {
  const XnixRuntimeApplication *application;
  const char *plan_type;
  const char *runtime_method;
  const char *review_state;
  XnixRuntimeCompatibilityPermission permissions[7];
  size_t permission_count;
  size_t allow_count;
  size_t ask_count;
  size_t deny_count;
  const char *required_runtime_gates[5];
  size_t required_runtime_gate_count;
  bool runtime_owned;
  bool c_runtime_backed;
  bool kde_policy_owner;
  bool user_review_required;
  bool portal_review_required;
  bool permission_changes_applied;
  bool request_objects_created;
  bool permissions_granted;
  bool settings_persisted;
  bool host_permission_changed;
  bool host_root_modified;
  bool backend_details_exposed;
  const char *summary;
} XnixRuntimeCompatibilityPermissionReviewPlan;

typedef struct {
  const char *id;
  const char *label;
  const char *runtime_method;
  const char *status;
  const char *summary;
  bool user_visible;
  bool user_action_required;
  bool runtime_gate_required;
  bool blocks_apply;
} XnixRuntimeCompatibilityReviewFlowStep;

typedef struct {
  const XnixRuntimeApplication *application;
  const char *plan_type;
  const char *runtime_method;
  const char *review_state;
  const char *section_id;
  const char *field_id;
  const char *requested_value;
  const char *operation;
  XnixRuntimeCompatibilityReviewFlowStep steps[5];
  size_t step_count;
  size_t required_review_count;
  size_t blocked_step_count;
  size_t pending_step_count;
  bool runtime_owned;
  bool c_runtime_backed;
  bool kde_policy_owner;
  bool user_confirmation_required;
  bool portal_policy_review_required;
  bool settings_change_planned;
  bool permission_review_planned;
  bool portal_request_planned;
  bool review_receipt_required;
  bool apply_enabled;
  bool request_object_created;
  bool permission_granted;
  bool settings_persisted;
  bool execution_started;
  bool host_root_modified;
  bool backend_details_exposed;
  const char *summary;
} XnixRuntimeCompatibilityReviewFlowPlan;

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
  const XnixRuntimeSnapshotPolicy *policy;
  bool runtime_owned;
  bool c_runtime_backed;
  bool kde_policy_owner;
  bool snapshot_request_created;
  bool snapshot_created;
  bool restore_requested;
  bool restore_executed;
  bool user_documents_included;
  bool host_system_included;
  bool host_root_modified;
  bool backend_details_exposed;
} XnixRuntimeCompatibilitySnapshotPlan;

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
  const char *status;
  const char *summary;
  bool user_visible;
  bool requires_portal_review;
  bool requires_state_root;
  bool requires_snapshot;
  bool activation_enabled;
  bool request_object_created;
  bool backend_process_started;
  bool backend_details_exposed;
} XnixRuntimeBackendCapability;

typedef struct {
  const char *id;
  const char *label;
  const char *kind;
  const char *selection_state;
  XnixRuntimeBackendCapability capabilities[7];
  size_t capability_count;
  size_t ready_count;
  size_t pending_count;
  size_t blocked_count;
  bool runtime_owned;
  bool kde_policy_owner;
  bool profile_ready;
  bool selection_enabled;
  bool backend_process_started;
  bool network_required_for_planning;
  bool host_root_modified;
  bool privileged_container_required;
  bool backend_details_exposed;
  const char *summary;
} XnixRuntimeBackendCapabilityProfile;

typedef struct {
  const char *matrix_type;
  const char *runtime_method;
  XnixRuntimeBackendCapabilityProfile profiles[2];
  size_t profile_count;
  size_t capability_count;
  size_t ready_capability_count;
  size_t pending_capability_count;
  size_t blocked_capability_count;
  bool runtime_owned;
  bool c_runtime_backed;
  bool kde_policy_owner;
  bool selection_enabled;
  bool backend_launch_enabled;
  bool capability_activation_enabled;
  bool request_objects_created;
  bool state_root_created;
  bool snapshots_created;
  bool host_root_modified;
  bool privileged_container_required;
  bool backend_details_exposed;
  const char *summary;
} XnixRuntimeBackendCapabilityMatrix;

typedef struct {
  const char *id;
  const char *label;
  const char *kind;
  const char *selection_state;
  const char *required_preflight[4];
  size_t required_preflight_count;
  size_t ready_capability_count;
  size_t pending_capability_count;
  bool recommended;
  bool ready;
  bool blocked;
  bool selection_committed;
  bool environment_created;
  bool backend_process_started;
  bool backend_details_exposed;
  const char *summary;
} XnixRuntimeBackendSelectionCandidate;

typedef struct {
  const XnixRuntimeApplication *application;
  const XnixRuntimeEngine *engine;
  const char *plan_type;
  const char *runtime_method;
  const char *selected_strategy;
  const char *recommended_profile_id;
  XnixRuntimeBackendSelectionCandidate candidates[2];
  size_t candidate_count;
  size_t ready_candidate_count;
  size_t blocked_candidate_count;
  const char *required_reviews[5];
  size_t required_review_count;
  const char *blocked_actions[4];
  size_t blocked_action_count;
  bool runtime_owned;
  bool c_runtime_backed;
  bool kde_policy_owner;
  bool selection_committed;
  bool selection_change_enabled;
  bool backend_launch_enabled;
  bool capability_activation_enabled;
  bool environment_created;
  bool request_object_created;
  bool state_root_created;
  bool snapshot_created;
  bool host_root_modified;
  bool privileged_container_required;
  bool backend_details_exposed;
  const char *summary;
} XnixRuntimeBackendSelectionPlan;

typedef struct {
  const char *id;
  const char *status;
  const char *summary;
} XnixRuntimeBackendLifecycleStage;

typedef struct {
  const XnixRuntimeApplication *application;
  const XnixRuntimeEngine *engine;
  const char *lifecycle_type;
  const char *runtime_method;
  const char *selected_strategy;
  const char *lifecycle_state;
  const char *overall_status;
  XnixRuntimeBackendLifecycleStage stages[5];
  size_t stage_count;
  const char *blocked_actions[5];
  size_t blocked_action_count;
  bool runtime_owned;
  bool c_runtime_backed;
  bool kde_policy_owner;
  bool backend_binding_ready;
  bool launch_enabled;
  bool execution_request_created;
  bool backend_process_started;
  bool local_backend_started;
  bool isolated_backend_started;
  bool state_root_ready;
  bool portal_review_required;
  bool snapshot_required;
  bool host_root_modified;
  bool network_required;
  bool backend_details_exposed;
  const char *summary;
} XnixRuntimeBackendLifecycle;

typedef struct {
  const char *id;
  const char *kind;
  const char *status;
  const char *summary;
  bool selected;
  bool environment_created;
  bool process_started;
  bool host_storage_exposed;
  bool backend_details_exposed;
} XnixRuntimeBackendEnvironmentProfile;

typedef struct {
  const XnixRuntimeApplication *application;
  const XnixRuntimeEngine *engine;
  const char *plan_type;
  const char *runtime_method;
  const char *selected_strategy;
  const char *environment_state;
  XnixRuntimeBackendEnvironmentProfile profiles[2];
  size_t profile_count;
  const char *required_reviews[4];
  size_t required_review_count;
  const char *blocked_actions[5];
  size_t blocked_action_count;
  bool runtime_owned;
  bool c_runtime_backed;
  bool kde_policy_owner;
  bool local_environment_ready;
  bool isolated_environment_ready;
  bool environment_created;
  bool backend_process_started;
  bool host_storage_exposed;
  bool clipboard_bridge_enabled;
  bool print_bridge_enabled;
  bool portal_review_required;
  bool snapshot_required;
  bool launch_enabled;
  bool host_root_modified;
  bool network_required;
  bool backend_details_exposed;
  const char *summary;
} XnixRuntimeBackendEnvironmentPlan;

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
  const char *id;
  const char *label;
  const char *intent;
  bool selected;
  bool requested;
  bool backend_details_exposed;
} XnixRuntimeCompatibilityModeOption;

typedef struct {
  const XnixRuntimeApplication *application;
  const char *plan_type;
  const char *runtime_method;
  const char *current_mode;
  const char *requested_mode;
  const char *mode_state;
  XnixRuntimeCompatibilityModeOption modes[4];
  size_t mode_count;
  const char *required_runtime_gates[5];
  size_t required_runtime_gate_count;
  bool runtime_owned;
  bool c_runtime_backed;
  bool kde_policy_owner;
  bool valid_mode;
  bool requires_user_confirmation;
  bool portal_review_required;
  bool snapshot_required;
  bool settings_persistence_enabled;
  bool backend_reconfiguration_enabled;
  bool backend_process_started;
  bool launch_enabled;
  bool host_root_modified;
  bool backend_details_exposed;
  const char *summary;
} XnixRuntimeCompatibilityModeSwitchPlan;

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
  const char *read_only_methods[61];
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
  const char *summary_type;
  const char *desktop;
  const char *compatibility_state;
  const char *runtime_mode;
  const char *known_issue_id;
  const char *known_issue_severity;
  const char *known_issue_summary;
  const char *repair_record_id;
  const char *repair_record_state;
  const char *last_repair_event;
  const char *actions[4];
  size_t action_count;
  size_t known_issue_count;
  size_t repair_record_count;
  bool runtime_owned;
  bool kde_policy_owner;
  bool user_visible;
  bool action_execution_enabled;
  bool repair_execution_enabled;
  bool backend_launch_enabled;
  bool settings_persistence_enabled;
  bool host_root_modified;
  bool backend_details_exposed;
  const char *summary;
} XnixRuntimeCompatibilityCenterSummary;

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
  const char *id;
  const char *status;
  const char *title;
  const char *summary;
  const char *required_for[3];
  size_t required_for_count;
  const char *portal_destination;
  const char *portal_interface;
  const char *portal_method;
  const char *portal_handle_token;
  const char *snapshot_reason;
  bool snapshot_enabled_by_default;
  bool snapshot_restore_available;
  const char *run_strategy;
  bool backend_ready;
  bool run_plan_backend_details_exposed;
} XnixRuntimeCompatibilityTestStep;

typedef struct {
  const XnixRuntimeApplication *application;
  const char *test_type;
  const char *desktop;
  XnixRuntimeCompatibilityTestStep steps[4];
  size_t step_count;
  const char *blocking_reasons[4];
  size_t blocking_reason_count;
  const char *notification_event;
  const char *repair_plan_issue;
  bool runtime_owned;
  bool c_runtime_backed;
  bool kde_policy_owner;
  bool blocked;
  bool compatibility_center_card;
  bool execution_request_created;
  bool test_executed;
  bool host_root_modified;
  bool backend_details_exposed;
  const char *summary;
} XnixRuntimeCompatibilityTestPlan;

typedef struct {
  size_t total;
  size_t passed;
  size_t pending;
  size_t blocked;
} XnixRuntimeCompatibilityTestResultCounts;

typedef struct {
  const char *id;
  const char *title;
  const char *status;
  const char *result;
  const char *evidence;
  const char *next_action;
} XnixRuntimeCompatibilityTestStepResult;

typedef struct {
  const XnixRuntimeApplication *application;
  const char *test_type;
  const char *result_source;
  const char *execution_state;
  const char *overall_status;
  XnixRuntimeCompatibilityTestResultCounts counts;
  XnixRuntimeCompatibilityTestStepResult step_results[4];
  size_t step_result_count;
  const char *notification_event;
  const char *repair_plan_issue;
  bool runtime_owned;
  bool c_runtime_backed;
  bool kde_policy_owner;
  bool compatibility_center_card;
  bool diagnostics_record;
  bool safe_for_ai_diagnostics;
  bool test_executed;
  bool host_root_modified;
  bool backend_details_exposed;
  const char *summary;
} XnixRuntimeCompatibilityTestResult;

typedef struct {
  const char *id;
  const char *status;
  const char *summary;
} XnixRuntimeExecutionReadinessGate;

typedef struct {
  const XnixRuntimeApplication *application;
  const XnixRuntimeEngine *engine;
  const char *readiness_type;
  const char *runtime_method;
  const char *execution_state;
  const char *overall_status;
  const char *recommended_action;
  XnixRuntimeExecutionReadinessGate gates[5];
  size_t gate_count;
  const char *blocked_actions[5];
  size_t blocked_action_count;
  bool runtime_owned;
  bool c_runtime_backed;
  bool kde_policy_owner;
  bool compatibility_center_card;
  bool safe_for_ai_diagnostics;
  bool desktop_entry_launch_visible;
  bool launch_allowed;
  bool launch_enabled;
  bool execution_request_created;
  bool backend_binding_ready;
  bool portal_policy_required;
  bool snapshot_required;
  bool user_action_required;
  bool host_root_modified;
  bool network_required;
  bool backend_details_exposed;
  const char *summary;
} XnixRuntimeExecutionReadiness;

typedef struct {
  const XnixRuntimeApplication *application;
  const XnixRuntimeEngine *engine;
  const char *intent_type;
  const char *source;
  const char *runtime_method;
  const char *read_method;
  const char *desktop_file;
  const char *launcher_command;
  const char *execution_state;
  const char *overall_status;
  const char *write_gate_decision;
  const char *denial_error_name;
  const char *blocked_actions[5];
  size_t blocked_action_count;
  bool runtime_owned;
  bool c_runtime_backed;
  bool kde_policy_owner;
  bool standard_desktop_entry;
  bool launch_uses_runtime;
  bool desktop_entry_launch_visible;
  bool portal_required;
  bool snapshot_required;
  bool backend_binding_ready;
  bool launch_allowed;
  bool launch_enabled;
  bool execution_request_created;
  bool execution_started;
  bool host_root_modified;
  bool network_required;
  bool backend_details_exposed;
  const char *summary;
} XnixRuntimeLaunchIntent;

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

typedef struct {
  const char *id;
  const char *name;
  const char *state;
  const char *model_command;
  const char *runtime_method;
  const char *adapter_role;
  bool runtime_backed;
  bool c_runtime_backed;
  bool dbus_read_available;
  bool kde_policy_owner;
  const char *summary;
  const char *next_step;
} XnixRuntimeKDEIntegrationEntryPoint;

typedef struct {
  XnixRuntimeKDEIntegrationEntryPoint entry_points[7];
  size_t entry_point_count;
  size_t initial_count;
  size_t planned_count;
  size_t complete_count;
  bool runtime_owned;
  bool kde_policy_owner;
  bool official_desktop_only;
  bool stable_desktop_contract;
  bool host_root_modified;
  bool backend_details_exposed;
  const char *desktop;
  const char *summary;
} XnixRuntimeKDEIntegrationStatus;

typedef struct {
  const char *id;
  const char *label;
  const char *state;
  const char *runtime_method;
  bool runtime_owned;
  bool kde_policy_owner;
  bool shell_writes_enabled;
  bool backend_launch_enabled;
  bool backend_details_exposed;
  const char *summary;
} XnixRuntimeKDEShellComponent;

typedef struct {
  const char *plan_type;
  const char *runtime_method;
  const char *desktop_shell;
  XnixRuntimeKDEShellComponent components[9];
  size_t component_count;
  size_t initial_component_count;
  size_t planned_component_count;
  bool official_desktop_only;
  bool fallback_desktops_supported;
  bool runtime_owned;
  bool c_runtime_backed;
  bool kde_policy_owner;
  bool plasma_fork_required;
  bool plasma_source_modified;
  bool shell_configuration_written;
  bool component_activation_enabled;
  bool backend_launch_enabled;
  bool backend_details_exposed;
  bool host_root_modified;
  bool privileged_container_required;
  const char *summary;
} XnixRuntimeKDEShellIntegrationPlan;

typedef struct {
  const char *id;
  const char *name;
  const char *runtime_method;
  const char *state;
  bool runtime_backed;
  bool c_runtime_backed;
  bool kde_writes_policy;
  bool backend_details_exposed;
  const char *summary;
} XnixRuntimeKDEApplicationSurfaceEntryPoint;

typedef struct {
  const XnixRuntimeApplication *application;
  const char *plan_type;
  const char *runtime_method;
  const char *surface_state;
  const char *desktop_shell;
  XnixRuntimeKDEApplicationSurfaceEntryPoint entry_points[7];
  size_t entry_point_count;
  const char *required_runtime_gates[6];
  size_t required_runtime_gate_count;
  bool runtime_owned;
  bool c_runtime_backed;
  bool kde_policy_owner;
  bool official_desktop_only;
  bool normal_linux_application_surface;
  bool standard_launcher_visible;
  bool task_manager_identity_ready;
  bool file_associations_planned;
  bool dolphin_action_planned;
  bool krunner_query_planned;
  bool tray_status_planned;
  bool notification_route_planned;
  bool settings_surface_planned;
  bool portal_review_required;
  bool execution_ready;
  bool launch_enabled;
  bool backend_process_started;
  bool desktop_files_written;
  bool mimeapps_written;
  bool host_root_modified;
  bool backend_command_exposed;
  bool raw_windows_executable_exposed;
  bool backend_details_exposed;
  const char *summary;
} XnixRuntimeKDEApplicationSurfacePlan;

typedef struct {
  const char *id;
  const char *name;
  const char *operation;
  const char *portal_interface;
  const char *runtime_method;
  const char *state;
  bool portal_required;
  bool user_approval_required;
  bool bridge_enabled;
  bool request_created;
  bool direct_backend_access_allowed;
  bool backend_details_exposed;
  const char *summary;
} XnixRuntimeDesktopResourceBridge;

typedef struct {
  const XnixRuntimeApplication *application;
  const char *plan_type;
  const char *runtime_method;
  const char *bridge_state;
  XnixRuntimeDesktopResourceBridge resources[5];
  size_t resource_count;
  const char *required_runtime_gates[5];
  size_t required_runtime_gate_count;
  bool runtime_owned;
  bool c_runtime_backed;
  bool kde_policy_owner;
  bool portal_mediated;
  bool file_bridge_planned;
  bool uri_bridge_planned;
  bool print_bridge_planned;
  bool clipboard_bridge_planned;
  bool screenshot_bridge_planned;
  bool bridges_enabled;
  bool requests_created;
  bool backend_process_started;
  bool direct_host_file_access;
  bool direct_clipboard_access;
  bool direct_print_access;
  bool host_root_modified;
  bool backend_details_exposed;
  const char *summary;
} XnixRuntimeDesktopResourceBridgePlan;

typedef struct {
  const XnixRuntimeApplication *application;
  const char *desktop_file;
  const char *relative_path;
  const char *name;
  const char *comment;
  const char *exec;
  const char *icon;
  const char *categories[2];
  size_t category_count;
  const char *mime_types[2];
  size_t mime_type_count;
  const char *startup_wm_class;
  const char *file_argument_mode;
  bool standard_desktop_entry;
  bool launch_uses_runtime;
  bool accepts_file_uris;
  bool user_visible;
  bool startup_notify;
  bool terminal;
  bool no_display;
  bool runtime_owned;
  bool kde_policy_owner;
  bool files_written;
  bool host_root_modified;
  bool backend_command_exposed;
  bool raw_windows_executable_exposed;
  bool compatibility_storage_path_exposed;
  bool backend_details_exposed;
  const char *summary;
} XnixRuntimeDesktopEntryPlan;

typedef struct {
  const XnixRuntimeApplication *application;
  const char *plan_type;
  const char *runtime_method;
  const char *desktop_file;
  const char *launcher_url;
  const char *target_directory;
  const char *placement;
  const char *exec;
  const char *icon;
  const char *blocked_actions[6];
  size_t blocked_action_count;
  bool standard_desktop_entry;
  bool desktop_icon_visible;
  bool user_visible;
  bool runtime_owned;
  bool c_runtime_backed;
  bool kde_policy_owner;
  bool official_desktop_only;
  bool desktop_file_copy_enabled;
  bool desktop_file_write_enabled;
  bool icon_placement_persisted;
  bool launch_enabled;
  bool backend_launch_enabled;
  bool host_root_modified;
  bool raw_executable_exposed;
  bool backend_details_exposed;
  const char *summary;
} XnixRuntimeDesktopIconPlan;

typedef struct {
  const XnixRuntimeApplication *application;
  const char *desktop_file;
  const char *launcher_url;
  const char *window_kind;
  const char *class_group;
  const char *resource_name;
  const char *title_hint;
  const char *grouping_key;
  const char *restore_key;
  const char *kwin_script_role;
  const char *placement;
  bool pinning_allowed;
  bool restore_allowed;
  bool skip_taskbar;
  bool show_in_switcher;
  bool prefer_existing_window;
  bool window_manager_policy_only;
  bool runtime_owns_backend_policy;
  bool runtime_owned;
  bool kde_policy_owner;
  bool host_root_modified;
  bool backend_details_exposed;
  const char *summary;
} XnixRuntimeTaskManagerIdentityPlan;

typedef struct {
  const XnixRuntimeApplication *application;
  const char *request_type;
  const char *desktop;
  const char *script_role;
  const char *resource_name;
  const char *class_group;
  const char *title_hint;
  const char *desktop_file;
  const char *application_id;
  const char *task_manager_grouping_key;
  const char *launcher_url;
  const char *placement;
  bool skip_taskbar;
  bool show_in_switcher;
  bool pinning_allowed;
  bool restore_allowed;
  const char *restore_key;
  bool prefer_existing_window;
  bool window_manager_policy_only;
  bool runtime_owns_backend_policy;
  bool runtime_owned;
  bool kde_policy_owner;
  bool host_root_modified;
  bool backend_details_exposed;
  const char *summary;
} XnixRuntimeKWinWindowRulePlan;

typedef struct {
  const XnixRuntimeApplication *application;
  const char *desktop_file;
  const char *mimeapps_path;
  const char *default_section;
  const char *added_section;
  const char *file_open_command;
  const char *file_open_argument;
  const char *supported_extensions[2];
  size_t supported_extension_count;
  const char *mime_types[2];
  size_t mime_type_count;
  bool runtime_owned;
  bool kde_policy_owner;
  bool standard_mimeapps_list;
  bool staged_root_only;
  bool overwrite_existing_mimeapps;
  bool portal_required_for_file_open;
  bool default_application;
  bool files_written;
  bool host_root_modified;
  bool backend_details_exposed;
  const char *summary;
} XnixRuntimeFileAssociationPlan;

typedef struct {
  const XnixRuntimeApplication *application;
  const char *event_type;
  const char *notification_id;
  const char *title;
  const char *body;
  const char *urgency;
  const char *category;
  const char *desktop_entry;
  const char *actions[3];
  size_t action_count;
  bool runtime_owned;
  bool kde_policy_owner;
  bool user_visible;
  bool requires_user_review;
  bool action_execution_enabled;
  bool repair_execution_enabled;
  bool settings_persistence_enabled;
  bool host_root_modified;
  bool backend_details_exposed;
  const char *summary;
} XnixRuntimeNotificationPlan;

typedef struct {
  const char *status_type;
  const char *desktop;
  size_t active_application_count;
  size_t attention_required_count;
  size_t bridged_tray_application_count;
  const char *runtime_activity_summary;
  const char *compatibility_state;
  const char *compatibility_label;
  const char *tray_bridge_state;
  const char *tray_bridge_label;
  const char *actions[2];
  size_t action_count;
  bool runtime_owned;
  bool kde_policy_owner;
  bool user_visible;
  bool live_backend_bridge_enabled;
  bool bridge_configuration_persisted;
  bool host_root_modified;
  bool backend_details_exposed;
  const char *summary;
} XnixRuntimeTrayStatusPlan;

typedef struct {
  const XnixRuntimeApplication *application;
  const char *runner_id;
  const char *subtitle;
  const char *mode_label;
  const char *action_type;
  const char *desktop_entry_id;
  const char *action_argv[3];
  size_t action_argc;
  const char *supported_extensions[2];
  size_t supported_extension_count;
  size_t relevance_percent;
  bool runtime_owned_launch;
  bool backend_details_exposed;
} XnixRuntimeKRunnerMatch;

typedef struct {
  const char *query_type;
  const char *entry_point;
  const char *desktop;
  const char *query;
  XnixRuntimeKRunnerMatch matches[1];
  size_t match_count;
  bool runtime_owned;
  bool kde_policy_owner;
  bool runtime_owned_launch;
  bool query_execution_enabled;
  bool backend_launch_enabled;
  bool host_root_modified;
  bool backend_details_exposed;
  const char *summary;
} XnixRuntimeKRunnerQueryPlan;

typedef struct {
  const char *id;
  const char *source;
  const char *summary;
  const char *primary_fact_key;
  const char *primary_fact_value;
} XnixRuntimeAIDiagnosticContextSection;

typedef struct {
  const char *id;
  const char *severity;
  const char *source;
  const char *summary;
  const char *recommended_next_action;
} XnixRuntimeAIDiagnosticSignal;

typedef struct {
  bool user_documents_included;
  bool host_paths_included;
  bool raw_backend_logs_included;
  bool secrets_included;
  bool network_calls_allowed;
  bool requires_user_approval_for_sensitive_actions;
} XnixRuntimeAIDiagnosticPrivacyBoundaries;

typedef struct {
  const XnixRuntimeApplication *application;
  const char *input_type;
  const char *issue;
  const char *test_type;
  XnixRuntimeAIDiagnosticContextSection context_sections[4];
  size_t context_section_count;
  XnixRuntimeAIDiagnosticSignal diagnostic_signals[3];
  size_t diagnostic_signal_count;
  XnixRuntimeAIDiagnosticPrivacyBoundaries privacy_boundaries;
  const char *allowed_ai_tasks[4];
  size_t allowed_ai_task_count;
  const char *blocked_ai_tasks[4];
  size_t blocked_ai_task_count;
  bool runtime_owned;
  bool c_runtime_backed;
  bool kde_policy_owner;
  bool ai_provider_called;
  bool network_required;
  bool safe_for_ai_diagnostics;
  bool backend_details_exposed;
  bool host_root_modified;
  const char *summary;
} XnixRuntimeAIDiagnosticInput;

typedef struct {
  const char *id;
  const char *priority;
  const char *source_signal;
  const char *title;
  const char *summary;
  const char *next_runtime_action;
  bool user_visible;
  bool auto_execute;
  bool requires_user_approval;
} XnixRuntimeAIDiagnosticRecommendationItem;

typedef struct {
  const char *id;
  const char *title;
  const char *reason;
  const char *approval_surface;
} XnixRuntimeAIApprovalRequiredAction;

typedef struct {
  XnixRuntimeAIDiagnosticInput diagnostic_input;
  const char *recommendation_type;
  XnixRuntimeAIDiagnosticRecommendationItem recommendations[3];
  size_t recommendation_count;
  XnixRuntimeAIApprovalRequiredAction approval_required_actions[1];
  size_t approval_required_action_count;
  const char *blocked_actions[4];
  size_t blocked_action_count;
  bool runtime_owned;
  bool c_runtime_backed;
  bool kde_policy_owner;
  bool ai_provider_called;
  bool network_required;
  bool safe_for_ai_diagnostics;
  bool auto_execution_allowed;
  bool backend_details_exposed;
  bool host_root_modified;
  const char *summary;
} XnixRuntimeAIDiagnosticRecommendation;

typedef struct {
  const char *id;
  const char *status;
  const char *summary;
} XnixRuntimeAIRepairRequiredGate;

typedef struct {
  XnixRuntimeAIDiagnosticRecommendation recommendation;
  const char *gate_type;
  const char *gate_decision;
  const char *approval_surface;
  XnixRuntimeAIRepairRequiredGate required_gates[3];
  size_t required_gate_count;
  const char *blocked_actions[7];
  size_t blocked_action_count;
  bool runtime_owned;
  bool c_runtime_backed;
  bool kde_policy_owner;
  bool ai_provider_called;
  bool network_required;
  bool safe_for_ai_diagnostics;
  bool repair_execution_requested;
  bool repair_executed;
  bool auto_execution_allowed;
  bool backend_details_exposed;
  bool host_root_modified;
  const char *summary;
} XnixRuntimeAIRepairApprovalGate;

typedef struct {
  const char *id;
  const char *kind;
  const char *label;
} XnixRuntimeCompatibilityRepairAction;

typedef struct {
  const char *application_id;
  const char *issue;
  const char *severity;
  XnixRuntimeCompatibilityRepairAction actions[2];
  size_t action_count;
  const XnixRuntimeSnapshotPolicy *snapshot_policy;
  const char *notification_event;
  bool automatic_allowed;
  bool user_approval_required;
  bool snapshot_required;
  bool rollback_available;
  bool runtime_owned;
  bool c_runtime_backed;
  bool kde_policy_owner;
  bool repair_execution_requested;
  bool repair_executed;
  bool backend_launch_enabled;
  bool network_required;
  bool host_root_modified;
  bool backend_details_exposed;
  const char *summary;
} XnixRuntimeCompatibilityRepairPlan;

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
bool xnix_runtime_compatibility_snapshot_plan(
  const char *application_id,
  const char *reason,
  XnixRuntimeCompatibilitySnapshotPlan *plan
);

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
bool xnix_runtime_backend_capability_matrix(XnixRuntimeBackendCapabilityMatrix *matrix);
bool xnix_runtime_backend_selection_plan(const char *application_id, XnixRuntimeBackendSelectionPlan *plan);
bool xnix_runtime_backend_lifecycle(const char *application_id, XnixRuntimeBackendLifecycle *lifecycle);
bool xnix_runtime_backend_environment_plan(const char *application_id, XnixRuntimeBackendEnvironmentPlan *plan);

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
bool xnix_runtime_compatibility_center_summary(
  const char *application_id,
  XnixRuntimeCompatibilityCenterSummary *summary
);
bool xnix_runtime_compatibility_run_plan(
  const char *application_id,
  XnixRuntimeCompatibilityRunPlan *plan
);
bool xnix_runtime_compatibility_test_plan(
  const char *application_id,
  const char *test_type,
  XnixRuntimeCompatibilityTestPlan *plan
);
bool xnix_runtime_compatibility_test_result(
  const char *application_id,
  const char *test_type,
  XnixRuntimeCompatibilityTestResult *result
);
bool xnix_runtime_execution_readiness(
  const char *application_id,
  XnixRuntimeExecutionReadiness *readiness
);
bool xnix_runtime_launch_intent(
  const char *application_id,
  XnixRuntimeLaunchIntent *intent
);
bool xnix_runtime_desktop_activation_manifest(
  const char *application_id,
  XnixRuntimeDesktopActivationManifest *manifest
);
bool xnix_runtime_kde_integration_status(
  XnixRuntimeKDEIntegrationStatus *status
);
bool xnix_runtime_kde_shell_integration_plan(
  XnixRuntimeKDEShellIntegrationPlan *plan
);
bool xnix_runtime_kde_application_surface_plan(
  const char *application_id,
  XnixRuntimeKDEApplicationSurfacePlan *plan
);
bool xnix_runtime_desktop_resource_bridge_plan(
  const char *application_id,
  XnixRuntimeDesktopResourceBridgePlan *plan
);
bool xnix_runtime_compatibility_mode_switch_plan(
  const char *application_id,
  const char *requested_mode,
  XnixRuntimeCompatibilityModeSwitchPlan *plan
);
bool xnix_runtime_compatibility_permission_review_plan(
  const char *application_id,
  XnixRuntimeCompatibilityPermissionReviewPlan *plan
);
bool xnix_runtime_compatibility_review_flow_plan(
  const char *application_id,
  const char *section_id,
  const char *field_id,
  const char *requested_value,
  const char *operation,
  XnixRuntimeCompatibilityReviewFlowPlan *plan
);
bool xnix_runtime_desktop_entry_plan(
  const char *application_id,
  XnixRuntimeDesktopEntryPlan *plan
);
bool xnix_runtime_desktop_icon_plan(
  const char *application_id,
  XnixRuntimeDesktopIconPlan *plan
);
bool xnix_runtime_task_manager_identity_plan(
  const char *application_id,
  XnixRuntimeTaskManagerIdentityPlan *plan
);
bool xnix_runtime_kwin_window_rule_plan(
  const char *application_id,
  XnixRuntimeKWinWindowRulePlan *plan
);
bool xnix_runtime_file_association_plan(
  const char *application_id,
  XnixRuntimeFileAssociationPlan *plan
);
bool xnix_runtime_notification_plan(
  const char *application_id,
  const char *event_type,
  XnixRuntimeNotificationPlan *plan
);
bool xnix_runtime_tray_status_plan(XnixRuntimeTrayStatusPlan *plan);
bool xnix_runtime_krunner_query_plan(
  const char *query,
  XnixRuntimeKRunnerQueryPlan *plan
);
bool xnix_runtime_portal_request_plan(
  const char *application_id,
  const char *operation,
  XnixRuntimePortalRequestPlan *plan
);
bool xnix_runtime_ai_diagnostic_input(
  const char *application_id,
  const char *issue,
  const char *test_type,
  XnixRuntimeAIDiagnosticInput *input
);
bool xnix_runtime_ai_diagnostic_recommendation(
  const char *application_id,
  const char *issue,
  const char *test_type,
  XnixRuntimeAIDiagnosticRecommendation *recommendation
);
bool xnix_runtime_ai_repair_approval_gate(
  const char *application_id,
  const char *issue,
  const char *test_type,
  XnixRuntimeAIRepairApprovalGate *gate
);
bool xnix_runtime_compatibility_repair_plan(
  const char *application_id,
  const char *issue,
  XnixRuntimeCompatibilityRepairPlan *plan
);

#endif
