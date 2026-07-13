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
print_snapshot_scope(const XnixRuntimeSnapshotPolicy *policy)
{
  fputs("{", stdout);
  fputs("\"application_state\":", stdout);
  print_bool(policy->application_state);
  fputs(",", stdout);
  fputs("\"runtime_metadata\":", stdout);
  print_bool(policy->runtime_metadata);
  fputs(",", stdout);
  fputs("\"desktop_activation_receipts\":", stdout);
  print_bool(policy->desktop_activation_receipts);
  fputs(",", stdout);
  fputs("\"user_documents\":", stdout);
  print_bool(policy->user_documents);
  fputs(",", stdout);
  fputs("\"host_system\":", stdout);
  print_bool(policy->host_system);
  fputs("}", stdout);
}

static void
print_snapshot_restore(const XnixRuntimeSnapshotPolicy *policy)
{
  fputs("{", stdout);
  fputs("\"available\":", stdout);
  print_bool(policy->restore_available);
  fputs(",", stdout);
  print_string_field("method", "Runtime.RestoreSnapshot");
  fputs(",", stdout);
  fputs("\"requires_user_confirmation\":", stdout);
  print_bool(policy->restore_requires_user_confirmation);
  fputs(",", stdout);
  fputs("\"preserve_user_documents\":", stdout);
  print_bool(policy->restore_preserve_user_documents);
  fputs("}", stdout);
}

static void
print_snapshot_retention(const XnixRuntimeSnapshotPolicy *policy)
{
  fputs("{", stdout);
  print_string_field("policy", policy->retention_policy);
  fputs(",", stdout);
  fputs("\"keep_latest\":", stdout);
  printf("%zu,", policy->keep_latest);
  fputs("\"prune_automatically\":", stdout);
  print_bool(policy->prune_automatically);
  fputs("}", stdout);
}

static void
print_snapshot_safety(const XnixRuntimeSnapshotPolicy *policy)
{
  fputs("{", stdout);
  fputs("\"runtime_policy_owner\":", stdout);
  print_bool(policy->runtime_policy_owner);
  fputs(",", stdout);
  fputs("\"desktop_shell_policy_owner\":", stdout);
  print_bool(policy->desktop_shell_policy_owner);
  fputs(",", stdout);
  fputs("\"backend_details_exposed\":", stdout);
  print_bool(policy->backend_details_exposed);
  fputs(",", stdout);
  fputs("\"host_root_modified\":", stdout);
  print_bool(policy->host_root_modified);
  fputs("}", stdout);
}

static void
print_state_root_application(const XnixRuntimeApplication *application)
{
  fputs("{", stdout);
  print_string_field("id", application->id);
  fputs(",", stdout);
  print_string_field("name", application->name);
  fputs(",", stdout);
  print_string_field("requested_mode", application->runtime_mode);
  fputs("}", stdout);
}

static void
print_state_root_retention(const XnixRuntimeStateRootPolicy *policy)
{
  fputs("{", stdout);
  fputs("\"automatic_restore_points\":", stdout);
  printf("%zu,", policy->automatic_restore_points);
  fputs("\"manual_restore_points\":", stdout);
  printf("%zu,", policy->manual_restore_points);
  fputs("\"user_documents_excluded\":", stdout);
  print_bool(policy->user_documents_excluded);
  fputs("}", stdout);
}

static void
print_state_root_managed_scopes(const XnixRuntimeStateRootPolicy *policy)
{
  fputc('[', stdout);
  for (size_t index = 0; index < policy->managed_scope_count; index++) {
    if (index > 0) {
      fputs(",", stdout);
    }
    fputs("{", stdout);
    print_string_field("id", policy->managed_scopes[index]);
    fputs(",", stdout);
    fputs("\"runtime_owned\":true,", stdout);
    fputs("\"snapshot_included\":", stdout);
    print_bool(policy->managed_scope_snapshot_included[index]);
    fputs(",", stdout);
    print_string_field("summary", policy->managed_scope_summaries[index]);
    fputs("}", stdout);
  }
  fputc(']', stdout);
}

static void
print_state_root_policy_record(
  const XnixRuntimeApplication *application,
  const XnixRuntimeStateRootPolicy *policy
)
{
  fputs("{", stdout);
  print_string_field("application_id", application->id);
  fputs(",", stdout);
  print_string_field("state_namespace", policy->state_namespace);
  fputs(",", stdout);
  print_string_field("storage_scope", policy->storage_scope);
  fputs(",", stdout);
  print_string_field("allocation_state", policy->allocation_state);
  fputs(",", stdout);
  fputs("\"runtime_owned\":", stdout);
  print_bool(policy->runtime_owned);
  fputs(",", stdout);
  fputs("\"kde_policy_owner\":", stdout);
  print_bool(policy->kde_policy_owner);
  fputs(",", stdout);
  fputs("\"directories_created\":", stdout);
  print_bool(policy->directories_created);
  fputs(",", stdout);
  fputs("\"host_root_modified\":", stdout);
  print_bool(policy->host_root_modified);
  fputs(",", stdout);
  fputs("\"user_documents_included\":", stdout);
  print_bool(policy->user_documents_included);
  fputs(",", stdout);
  fputs("\"portal_required_for_user_files\":", stdout);
  print_bool(policy->portal_required_for_user_files);
  fputs(",", stdout);
  fputs("\"snapshot_eligible\":", stdout);
  print_bool(policy->snapshot_eligible);
  fputs(",", stdout);
  fputs("\"restore_requires_confirmation\":", stdout);
  print_bool(policy->restore_requires_confirmation);
  fputs(",", stdout);
  fputs("\"managed_scopes\":", stdout);
  print_string_array(policy->managed_scopes, policy->managed_scope_count);
  fputs(",", stdout);
  print_string_field("desktop_safe_summary", policy->summary);
  fputs(",", stdout);
  fputs("\"backend_details_exposed\":", stdout);
  print_bool(policy->backend_details_exposed);
  fputs("}", stdout);
}

static void
print_state_root_policy_for_application(
  const XnixRuntimeApplication *application,
  const XnixRuntimeStateRootPolicy *policy
)
{
  fputs("{", stdout);
  printf("\"version\":\"%s\",", xnix_runtime_version());
  print_string_field("root_type", "compatibility-application-state-root");
  fputs(",", stdout);
  fputs("\"application\":", stdout);
  print_state_root_application(application);
  fputs(",", stdout);
  fputs("\"runtime_owned\":", stdout);
  print_bool(policy->runtime_owned);
  fputs(",", stdout);
  fputs("\"kde_policy_owner\":", stdout);
  print_bool(policy->kde_policy_owner);
  fputs(",", stdout);
  print_string_field("state_namespace", policy->state_namespace);
  fputs(",", stdout);
  print_string_field("storage_scope", policy->storage_scope);
  fputs(",", stdout);
  print_string_field("allocation_state", policy->allocation_state);
  fputs(",", stdout);
  fputs("\"directories_created\":", stdout);
  print_bool(policy->directories_created);
  fputs(",", stdout);
  fputs("\"host_root_modified\":", stdout);
  print_bool(policy->host_root_modified);
  fputs(",", stdout);
  fputs("\"user_documents_included\":", stdout);
  print_bool(policy->user_documents_included);
  fputs(",", stdout);
  fputs("\"portal_required_for_user_files\":", stdout);
  print_bool(policy->portal_required_for_user_files);
  fputs(",", stdout);
  fputs("\"snapshot_eligible\":", stdout);
  print_bool(policy->snapshot_eligible);
  fputs(",", stdout);
  fputs("\"restore_requires_confirmation\":", stdout);
  print_bool(policy->restore_requires_confirmation);
  fputs(",", stdout);
  fputs("\"retention_policy\":", stdout);
  print_state_root_retention(policy);
  fputs(",", stdout);
  fputs("\"managed_scopes\":", stdout);
  print_state_root_managed_scopes(policy);
  fputs(",", stdout);
  fputs("\"blocked_actions\":", stdout);
  print_string_array(policy->blocked_actions, policy->blocked_action_count);
  fputs(",", stdout);
  print_string_field("desktop_safe_summary", policy->summary);
  fputs(",", stdout);
  fputs("\"backend_details_exposed\":", stdout);
  print_bool(policy->backend_details_exposed);
  fputs("}", stdout);
}

static void
print_artifact_application(const XnixRuntimeApplication *application)
{
  fputs("{", stdout);
  print_string_field("id", application->id);
  fputs(",", stdout);
  print_string_field("name", application->name);
  fputs(",", stdout);
  print_string_field("requested_mode", application->runtime_mode);
  fputs("}", stdout);
}

static void
print_artifact_groups(const XnixRuntimeArtifactManifestPolicy *policy)
{
  fputc('[', stdout);
  for (size_t index = 0; index < policy->artifact_group_count; index++) {
    const XnixRuntimeArtifactGroup *group = &policy->artifact_groups[index];

    if (index > 0) {
      fputs(",", stdout);
    }
    fputs("{", stdout);
    print_string_field("id", group->id);
    fputs(",", stdout);
    print_string_field("kind", group->kind);
    fputs(",", stdout);
    fputs("\"required\":", stdout);
    print_bool(group->required);
    fputs(",", stdout);
    fputs("\"resolved\":", stdout);
    print_bool(group->resolved);
    fputs(",", stdout);
    fputs("\"downloaded\":", stdout);
    print_bool(group->downloaded);
    fputs(",", stdout);
    print_string_field("cache_namespace", group->cache_namespace);
    fputs(",", stdout);
    print_string_field("summary", group->summary);
    fputs("}", stdout);
  }
  fputc(']', stdout);
}

static void
print_artifact_required_preflight(const XnixRuntimeArtifactManifestPolicy *policy)
{
  fputc('[', stdout);
  for (size_t index = 0; index < policy->required_preflight_count; index++) {
    const XnixRuntimeArtifactPreflight *preflight = &policy->required_preflight[index];

    if (index > 0) {
      fputs(",", stdout);
    }
    fputs("{", stdout);
    print_string_field("id", preflight->id);
    fputs(",", stdout);
    print_string_field("status", preflight->status);
    fputs(",", stdout);
    print_string_field("summary", preflight->summary);
    fputs("}", stdout);
  }
  fputc(']', stdout);
}

static void
print_artifact_manifest_policy_record(const XnixRuntimeArtifactManifestPolicy *policy)
{
  fputs("{", stdout);
  print_string_field("application_id", policy->application_id);
  fputs(",", stdout);
  print_string_field("manifest_state", policy->manifest_state);
  fputs(",", stdout);
  fputs("\"manifest_ready\":", stdout);
  print_bool(policy->manifest_ready);
  fputs(",", stdout);
  fputs("\"signature_verified\":", stdout);
  print_bool(policy->signature_verified);
  fputs(",", stdout);
  fputs("\"download_enabled\":", stdout);
  print_bool(policy->download_enabled);
  fputs(",", stdout);
  fputs("\"artifacts_downloaded\":", stdout);
  print_bool(policy->artifacts_downloaded);
  fputs(",", stdout);
  fputs("\"host_root_modified\":", stdout);
  print_bool(policy->host_root_modified);
  fputs(",", stdout);
  print_string_field("selected_strategy", policy->selected_strategy);
  fputs(",", stdout);
  print_string_field("desktop_safe_summary", policy->summary);
  fputs(",", stdout);
  fputs("\"backend_details_exposed\":", stdout);
  print_bool(policy->backend_details_exposed);
  fputs("}", stdout);
}

static void
print_artifact_manifest_policy_for_application(
  const XnixRuntimeApplication *application,
  const XnixRuntimeArtifactManifestPolicy *policy
)
{
  fputs("{", stdout);
  printf("\"version\":\"%s\",", xnix_runtime_version());
  print_string_field("manifest_type", "compatibility-artifact-manifest");
  fputs(",", stdout);
  fputs("\"application\":", stdout);
  print_artifact_application(application);
  fputs(",", stdout);
  fputs("\"runtime_owned\":", stdout);
  print_bool(policy->runtime_owned);
  fputs(",", stdout);
  fputs("\"kde_policy_owner\":", stdout);
  print_bool(policy->kde_policy_owner);
  fputs(",", stdout);
  print_string_field("selected_strategy", policy->selected_strategy);
  fputs(",", stdout);
  print_string_field("manifest_state", policy->manifest_state);
  fputs(",", stdout);
  fputs("\"manifest_ready\":", stdout);
  print_bool(policy->manifest_ready);
  fputs(",", stdout);
  fputs("\"signature_verified\":", stdout);
  print_bool(policy->signature_verified);
  fputs(",", stdout);
  fputs("\"acquisition_preflight_ready\":", stdout);
  print_bool(policy->acquisition_preflight_ready);
  fputs(",", stdout);
  fputs("\"download_enabled\":", stdout);
  print_bool(policy->download_enabled);
  fputs(",", stdout);
  fputs("\"install_enabled\":", stdout);
  print_bool(policy->install_enabled);
  fputs(",", stdout);
  fputs("\"network_request_created\":", stdout);
  print_bool(policy->network_request_created);
  fputs(",", stdout);
  fputs("\"artifacts_downloaded\":", stdout);
  print_bool(policy->artifacts_downloaded);
  fputs(",", stdout);
  fputs("\"host_root_modified\":", stdout);
  print_bool(policy->host_root_modified);
  fputs(",", stdout);
  fputs("\"privileged_container_required\":", stdout);
  print_bool(policy->privileged_container_required);
  fputs(",", stdout);
  fputs("\"desktop_shell_command_exposed\":", stdout);
  print_bool(policy->desktop_shell_command_exposed);
  fputs(",", stdout);
  fputs("\"artifact_groups\":", stdout);
  print_artifact_groups(policy);
  fputs(",", stdout);
  fputs("\"required_preflight\":", stdout);
  print_artifact_required_preflight(policy);
  fputs(",", stdout);
  fputs("\"blocked_actions\":", stdout);
  print_string_array(policy->blocked_actions, policy->blocked_action_count);
  fputs(",", stdout);
  print_string_field("desktop_safe_summary", policy->summary);
  fputs(",", stdout);
  fputs("\"backend_details_exposed\":", stdout);
  print_bool(policy->backend_details_exposed);
  fputs("}", stdout);
}

static void
print_acquisition_application(const XnixRuntimeApplication *application)
{
  fputs("{", stdout);
  print_string_field("id", application->id);
  fputs(",", stdout);
  print_string_field("name", application->name);
  fputs(",", stdout);
  print_string_field("requested_mode", application->runtime_mode);
  fputs("}", stdout);
}

static void
print_acquisition_checks(const XnixRuntimeAcquisitionPreflightPolicy *policy)
{
  fputc('[', stdout);
  for (size_t index = 0; index < policy->check_count; index++) {
    const XnixRuntimeAcquisitionCheck *check = &policy->checks[index];

    if (index > 0) {
      fputs(",", stdout);
    }
    fputs("{", stdout);
    print_string_field("id", check->id);
    fputs(",", stdout);
    print_string_field("status", check->status);
    fputs(",", stdout);
    print_string_field("summary", check->summary);
    fputs("}", stdout);
  }
  fputc(']', stdout);
}

static void
print_acquisition_preflight_policy_record(const XnixRuntimeAcquisitionPreflightPolicy *policy)
{
  fputs("{", stdout);
  print_string_field("application_id", policy->application_id);
  fputs(",", stdout);
  print_string_field("preflight_state", policy->preflight_state);
  fputs(",", stdout);
  fputs("\"acquisition_ready\":", stdout);
  print_bool(policy->acquisition_ready);
  fputs(",", stdout);
  fputs("\"download_enabled\":", stdout);
  print_bool(policy->download_enabled);
  fputs(",", stdout);
  fputs("\"network_request_created\":", stdout);
  print_bool(policy->network_request_created);
  fputs(",", stdout);
  fputs("\"artifacts_downloaded\":", stdout);
  print_bool(policy->artifacts_downloaded);
  fputs(",", stdout);
  fputs("\"host_root_modified\":", stdout);
  print_bool(policy->host_root_modified);
  fputs(",", stdout);
  print_string_field("selected_strategy", policy->selected_strategy);
  fputs(",", stdout);
  print_string_field("desktop_safe_summary", policy->summary);
  fputs(",", stdout);
  fputs("\"backend_details_exposed\":", stdout);
  print_bool(policy->backend_details_exposed);
  fputs("}", stdout);
}

static void
print_acquisition_preflight_policy_for_application(
  const XnixRuntimeApplication *application,
  const XnixRuntimeAcquisitionPreflightPolicy *policy
)
{
  fputs("{", stdout);
  printf("\"version\":\"%s\",", xnix_runtime_version());
  print_string_field("preflight_type", "compatibility-acquisition-preflight");
  fputs(",", stdout);
  fputs("\"application\":", stdout);
  print_acquisition_application(application);
  fputs(",", stdout);
  fputs("\"runtime_owned\":", stdout);
  print_bool(policy->runtime_owned);
  fputs(",", stdout);
  fputs("\"kde_policy_owner\":", stdout);
  print_bool(policy->kde_policy_owner);
  fputs(",", stdout);
  print_string_field("selected_strategy", policy->selected_strategy);
  fputs(",", stdout);
  print_string_field("preflight_state", policy->preflight_state);
  fputs(",", stdout);
  fputs("\"acquisition_ready\":", stdout);
  print_bool(policy->acquisition_ready);
  fputs(",", stdout);
  fputs("\"download_enabled\":", stdout);
  print_bool(policy->download_enabled);
  fputs(",", stdout);
  fputs("\"install_enabled\":", stdout);
  print_bool(policy->install_enabled);
  fputs(",", stdout);
  fputs("\"network_required_for_planning\":", stdout);
  print_bool(policy->network_required_for_planning);
  fputs(",", stdout);
  fputs("\"network_request_created\":", stdout);
  print_bool(policy->network_request_created);
  fputs(",", stdout);
  fputs("\"artifacts_downloaded\":", stdout);
  print_bool(policy->artifacts_downloaded);
  fputs(",", stdout);
  fputs("\"host_root_modified\":", stdout);
  print_bool(policy->host_root_modified);
  fputs(",", stdout);
  fputs("\"privileged_container_required\":", stdout);
  print_bool(policy->privileged_container_required);
  fputs(",", stdout);
  fputs("\"desktop_shell_command_exposed\":", stdout);
  print_bool(policy->desktop_shell_command_exposed);
  fputs(",", stdout);
  fputs("\"package_source_ready\":", stdout);
  print_bool(policy->package_source_ready);
  fputs(",", stdout);
  fputs("\"checks\":", stdout);
  print_acquisition_checks(policy);
  fputs(",", stdout);
  fputs("\"blocked_actions\":", stdout);
  print_string_array(policy->blocked_actions, policy->blocked_action_count);
  fputs(",", stdout);
  print_string_field("desktop_safe_summary", policy->summary);
  fputs(",", stdout);
  fputs("\"backend_details_exposed\":", stdout);
  print_bool(policy->backend_details_exposed);
  fputs("}", stdout);
}

static void
print_install_application(const XnixRuntimeApplication *application)
{
  fputs("{", stdout);
  print_string_field("id", application->id);
  fputs(",", stdout);
  print_string_field("name", application->name);
  fputs(",", stdout);
  print_string_field("requested_mode", application->runtime_mode);
  fputs(",", stdout);
  fputs("\"supported_extensions\":[", stdout);
  print_string(application->primary_extension);
  fputs("],", stdout);
  fputs("\"mime_types\":[", stdout);
  print_string(application->primary_mime_type);
  fputs("]}", stdout);
}

static void
print_install_readiness(
  const XnixRuntimeInstallReadinessPolicy *policy,
  const char *recipe_install_decision
)
{
  fputs("{", stdout);
  fputs("\"artifact_manifest_ready\":", stdout);
  print_bool(policy->artifact_manifest_ready);
  fputs(",", stdout);
  fputs("\"artifact_signature_verified\":", stdout);
  print_bool(policy->artifact_signature_verified);
  fputs(",", stdout);
  fputs("\"acquisition_ready\":", stdout);
  print_bool(policy->acquisition_ready);
  fputs(",", stdout);
  fputs("\"package_source_ready\":", stdout);
  print_bool(policy->package_source_ready);
  fputs(",", stdout);
  fputs("\"state_root_allocated\":", stdout);
  print_bool(policy->state_root_allocated);
  fputs(",", stdout);
  fputs("\"recipe_install_allowed\":", stdout);
  print_bool(strcmp(recipe_install_decision, "allow") == 0);
  fputs(",", stdout);
  print_string_field("recipe_install_decision", recipe_install_decision);
  fputs("}", stdout);
}

static void
print_install_phases(const XnixRuntimeInstallReadinessPolicy *policy)
{
  fputc('[', stdout);
  for (size_t index = 0; index < policy->phase_count; index++) {
    const XnixRuntimeInstallPhase *phase = &policy->phases[index];

    if (index > 0) {
      fputs(",", stdout);
    }
    fputs("{", stdout);
    print_string_field("id", phase->id);
    fputs(",", stdout);
    print_string_field("status", phase->status);
    fputs(",", stdout);
    print_string_field("summary", phase->summary);
    fputs("}", stdout);
  }
  fputc(']', stdout);
}

static void
print_install_policy_record(const XnixRuntimeInstallReadinessPolicy *policy)
{
  fputs("{", stdout);
  print_string_field("application_id", policy->application_id);
  fputs(",", stdout);
  print_string_field("install_state", policy->install_state);
  fputs(",", stdout);
  fputs("\"install_ready\":", stdout);
  print_bool(policy->install_ready);
  fputs(",", stdout);
  fputs("\"desktop_activation_ready\":", stdout);
  print_bool(policy->desktop_activation_ready);
  fputs(",", stdout);
  fputs("\"download_enabled\":", stdout);
  print_bool(policy->download_enabled);
  fputs(",", stdout);
  fputs("\"install_enabled\":", stdout);
  print_bool(policy->install_enabled);
  fputs(",", stdout);
  fputs("\"host_root_modified\":", stdout);
  print_bool(policy->host_root_modified);
  fputs(",", stdout);
  print_string_field("selected_strategy", policy->selected_strategy);
  fputs(",", stdout);
  print_string_field("desktop_safe_summary", policy->summary);
  fputs(",", stdout);
  fputs("\"backend_details_exposed\":", stdout);
  print_bool(policy->backend_details_exposed);
  fputs("}", stdout);
}

static void
print_install_policy_for_application(
  const XnixRuntimeApplication *application,
  const XnixRuntimeInstallReadinessPolicy *policy,
  const char *environment,
  const char *recipe_install_decision
)
{
  fputs("{", stdout);
  printf("\"version\":\"%s\",", xnix_runtime_version());
  print_string_field("plan_type", "compatibility-install-readiness");
  fputs(",", stdout);
  fputs("\"application\":", stdout);
  print_install_application(application);
  fputs(",", stdout);
  fputs("\"runtime_owned\":", stdout);
  print_bool(policy->runtime_owned);
  fputs(",", stdout);
  fputs("\"kde_policy_owner\":", stdout);
  print_bool(policy->kde_policy_owner);
  fputs(",", stdout);
  print_string_field("environment", environment);
  fputs(",", stdout);
  print_string_field("install_state", policy->install_state);
  fputs(",", stdout);
  fputs("\"install_ready\":", stdout);
  print_bool(policy->install_ready);
  fputs(",", stdout);
  fputs("\"desktop_activation_ready\":", stdout);
  print_bool(policy->desktop_activation_ready);
  fputs(",", stdout);
  fputs("\"download_enabled\":", stdout);
  print_bool(policy->download_enabled);
  fputs(",", stdout);
  fputs("\"install_enabled\":", stdout);
  print_bool(policy->install_enabled);
  fputs(",", stdout);
  fputs("\"network_request_created\":", stdout);
  print_bool(policy->network_request_created);
  fputs(",", stdout);
  fputs("\"artifacts_downloaded\":", stdout);
  print_bool(policy->artifacts_downloaded);
  fputs(",", stdout);
  fputs("\"host_root_modified\":", stdout);
  print_bool(policy->host_root_modified);
  fputs(",", stdout);
  fputs("\"privileged_container_required\":", stdout);
  print_bool(policy->privileged_container_required);
  fputs(",", stdout);
  fputs("\"desktop_shell_command_exposed\":", stdout);
  print_bool(policy->desktop_shell_command_exposed);
  fputs(",", stdout);
  print_string_field("selected_strategy", policy->selected_strategy);
  fputs(",", stdout);
  fputs("\"readiness\":", stdout);
  print_install_readiness(policy, recipe_install_decision);
  fputs(",", stdout);
  fputs("\"phases\":", stdout);
  print_install_phases(policy);
  fputs(",", stdout);
  fputs("\"blocked_actions\":", stdout);
  print_string_array(policy->blocked_actions, policy->blocked_action_count);
  fputs(",", stdout);
  print_string_field("desktop_safe_summary", policy->summary);
  fputs(",", stdout);
  fputs("\"backend_details_exposed\":", stdout);
  print_bool(policy->backend_details_exposed);
  fputs("}", stdout);
}

static void
print_snapshot_policy_record(const XnixRuntimeSnapshotPolicy *policy)
{
  fputs("{", stdout);
  print_string_field("reason", policy->reason);
  fputs(",", stdout);
  fputs("\"enabled_by_default\":", stdout);
  print_bool(policy->enabled_by_default);
  fputs(",", stdout);
  fputs("\"snapshot_scope\":", stdout);
  print_snapshot_scope(policy);
  fputs(",", stdout);
  fputs("\"restore\":", stdout);
  print_snapshot_restore(policy);
  fputs(",", stdout);
  fputs("\"retention\":", stdout);
  print_snapshot_retention(policy);
  fputs(",", stdout);
  print_string_field("desktop_safe_summary", policy->summary);
  fputs(",", stdout);
  fputs("\"safety\":", stdout);
  print_snapshot_safety(policy);
  fputs("}", stdout);
}

static void
print_snapshot_policy_for_application(const char *application_id, const XnixRuntimeSnapshotPolicy *policy)
{
  fputs("{", stdout);
  printf("\"version\":\"%s\",", xnix_runtime_version());
  print_string_field("plan_type", "compatibility-snapshot");
  fputs(",", stdout);
  print_string_field("application_id", application_id);
  fputs(",", stdout);
  print_string_field("reason", policy->reason);
  fputs(",", stdout);
  fputs("\"enabled_by_default\":", stdout);
  print_bool(policy->enabled_by_default);
  fputs(",", stdout);
  fputs("\"snapshot_scope\":", stdout);
  print_snapshot_scope(policy);
  fputs(",", stdout);
  fputs("\"restore\":", stdout);
  print_snapshot_restore(policy);
  fputs(",", stdout);
  fputs("\"retention\":", stdout);
  print_snapshot_retention(policy);
  fputs(",", stdout);
  print_string_field("desktop_safe_summary", policy->summary);
  fputs(",", stdout);
  fputs("\"safety\":", stdout);
  print_snapshot_safety(policy);
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
print_snapshot_policy_catalog(void)
{
  fputs("{", stdout);
  printf("\"version\":\"%s\",", xnix_runtime_version());
  print_string_field("catalog_type", "snapshot-policy");
  fputs(",", stdout);
  fputs("\"runtime_policy_owner\":true,", stdout);
  fputs("\"desktop_shell_policy_owner\":false,", stdout);
  print_string_field("catalog_owner", "c");
  fputs(",", stdout);
  fputs("\"backend_details_exposed\":false,", stdout);
  fputs("\"host_root_modified\":false,", stdout);
  fputs("\"policy_count\":", stdout);
  printf("%zu,", xnix_runtime_snapshot_policy_count());
  fputs("\"policies\":[", stdout);
  for (size_t index = 0; index < xnix_runtime_snapshot_policy_count(); index++) {
    const XnixRuntimeSnapshotPolicy *policy = xnix_runtime_snapshot_policy_at(index);

    if (index > 0) {
      fputs(",", stdout);
    }
    print_snapshot_policy_record(policy);
  }
  fputs("]}\n", stdout);
  return 0;
}

static int
print_snapshot_policy(const char *application_id, const char *reason)
{
  const XnixRuntimeApplication *application = xnix_runtime_find_application(application_id);
  const XnixRuntimeSnapshotPolicy *policy = xnix_runtime_find_snapshot_policy(reason);

  if (application == NULL) {
    fprintf(stderr, "xnix-runtime-core: application is not registered in the C Runtime catalog\n");
    return 64;
  }
  if (policy == NULL) {
    fprintf(stderr, "xnix-runtime-core: reason must be a registered snapshot reason\n");
    return 64;
  }

  print_snapshot_policy_for_application(application->id, policy);
  fputs("\n", stdout);
  return 0;
}

static int
print_state_root_policy_catalog(void)
{
  fputs("{", stdout);
  printf("\"version\":\"%s\",", xnix_runtime_version());
  print_string_field("catalog_type", "application-state-root-policy");
  fputs(",", stdout);
  fputs("\"runtime_policy_owner\":true,", stdout);
  fputs("\"desktop_shell_policy_owner\":false,", stdout);
  print_string_field("catalog_owner", "c");
  fputs(",", stdout);
  fputs("\"backend_details_exposed\":false,", stdout);
  fputs("\"host_root_modified\":false,", stdout);
  fputs("\"policy_count\":", stdout);
  printf("%zu,", xnix_runtime_state_root_policy_count());
  fputs("\"policies\":[", stdout);
  for (size_t index = 0; index < xnix_runtime_state_root_policy_count(); index++) {
    const XnixRuntimeStateRootPolicy *policy = xnix_runtime_state_root_policy_at(index);
    const XnixRuntimeApplication *application = xnix_runtime_find_application(policy->application_id);

    if (index > 0) {
      fputs(",", stdout);
    }
    print_state_root_policy_record(application, policy);
  }
  fputs("]}\n", stdout);
  return 0;
}

static int
print_state_root_policy(const char *application_id)
{
  const XnixRuntimeApplication *application = xnix_runtime_find_application(application_id);
  const XnixRuntimeStateRootPolicy *policy = xnix_runtime_find_state_root_policy(application_id);

  if (application == NULL || policy == NULL) {
    fprintf(stderr, "xnix-runtime-core: application is not registered in the C Runtime state root policy catalog\n");
    return 64;
  }

  print_state_root_policy_for_application(application, policy);
  fputs("\n", stdout);
  return 0;
}

static int
print_install_readiness_policy_catalog(void)
{
  fputs("{", stdout);
  printf("\"version\":\"%s\",", xnix_runtime_version());
  print_string_field("catalog_type", "install-readiness-policy");
  fputs(",", stdout);
  fputs("\"runtime_policy_owner\":true,", stdout);
  fputs("\"desktop_shell_policy_owner\":false,", stdout);
  print_string_field("catalog_owner", "c");
  fputs(",", stdout);
  fputs("\"backend_details_exposed\":false,", stdout);
  fputs("\"host_root_modified\":false,", stdout);
  fputs("\"policy_count\":", stdout);
  printf("%zu,", xnix_runtime_install_readiness_policy_count());
  fputs("\"policies\":[", stdout);
  for (size_t index = 0; index < xnix_runtime_install_readiness_policy_count(); index++) {
    const XnixRuntimeInstallReadinessPolicy *policy =
      xnix_runtime_install_readiness_policy_at(index);

    if (index > 0) {
      fputs(",", stdout);
    }
    print_install_policy_record(policy);
  }
  fputs("]}\n", stdout);
  return 0;
}

static int
print_install_readiness_policy(const char *application_id, const char *environment)
{
  const XnixRuntimeApplication *application = xnix_runtime_find_application(application_id);
  const XnixRuntimeInstallReadinessPolicy *policy =
    xnix_runtime_find_install_readiness_policy(application_id);
  const char *recipe_install_decision = "block";

  if (strcmp(environment, "development") != 0 && strcmp(environment, "production") != 0) {
    fprintf(stderr, "xnix-runtime-core: environment must be development or production\n");
    return 64;
  }
  if (application == NULL || policy == NULL) {
    fprintf(stderr, "xnix-runtime-core: application is not registered in the C Runtime install readiness policy catalog\n");
    return 64;
  }
  if (xnix_runtime_install_readiness_allows_recipe_install(policy, environment)) {
    recipe_install_decision = "allow";
  }

  print_install_policy_for_application(application, policy, environment, recipe_install_decision);
  fputs("\n", stdout);
  return 0;
}

static int
print_artifact_manifest_policy_catalog(void)
{
  fputs("{", stdout);
  printf("\"version\":\"%s\",", xnix_runtime_version());
  print_string_field("catalog_type", "artifact-manifest-policy");
  fputs(",", stdout);
  fputs("\"runtime_policy_owner\":true,", stdout);
  fputs("\"desktop_shell_policy_owner\":false,", stdout);
  print_string_field("catalog_owner", "c");
  fputs(",", stdout);
  fputs("\"backend_details_exposed\":false,", stdout);
  fputs("\"host_root_modified\":false,", stdout);
  fputs("\"policy_count\":", stdout);
  printf("%zu,", xnix_runtime_artifact_manifest_policy_count());
  fputs("\"policies\":[", stdout);
  for (size_t index = 0; index < xnix_runtime_artifact_manifest_policy_count(); index++) {
    const XnixRuntimeArtifactManifestPolicy *policy =
      xnix_runtime_artifact_manifest_policy_at(index);

    if (index > 0) {
      fputs(",", stdout);
    }
    print_artifact_manifest_policy_record(policy);
  }
  fputs("]}\n", stdout);
  return 0;
}

static int
print_artifact_manifest_policy(const char *application_id)
{
  const XnixRuntimeApplication *application = xnix_runtime_find_application(application_id);
  const XnixRuntimeArtifactManifestPolicy *policy =
    xnix_runtime_find_artifact_manifest_policy(application_id);

  if (application == NULL || policy == NULL) {
    fprintf(stderr, "xnix-runtime-core: application is not registered in the C Runtime artifact manifest policy catalog\n");
    return 64;
  }

  print_artifact_manifest_policy_for_application(application, policy);
  fputs("\n", stdout);
  return 0;
}

static int
print_acquisition_preflight_policy_catalog(void)
{
  fputs("{", stdout);
  printf("\"version\":\"%s\",", xnix_runtime_version());
  print_string_field("catalog_type", "acquisition-preflight-policy");
  fputs(",", stdout);
  fputs("\"runtime_policy_owner\":true,", stdout);
  fputs("\"desktop_shell_policy_owner\":false,", stdout);
  print_string_field("catalog_owner", "c");
  fputs(",", stdout);
  fputs("\"backend_details_exposed\":false,", stdout);
  fputs("\"host_root_modified\":false,", stdout);
  fputs("\"policy_count\":", stdout);
  printf("%zu,", xnix_runtime_acquisition_preflight_policy_count());
  fputs("\"policies\":[", stdout);
  for (size_t index = 0; index < xnix_runtime_acquisition_preflight_policy_count(); index++) {
    const XnixRuntimeAcquisitionPreflightPolicy *policy =
      xnix_runtime_acquisition_preflight_policy_at(index);

    if (index > 0) {
      fputs(",", stdout);
    }
    print_acquisition_preflight_policy_record(policy);
  }
  fputs("]}\n", stdout);
  return 0;
}

static int
print_acquisition_preflight_policy(const char *application_id)
{
  const XnixRuntimeApplication *application = xnix_runtime_find_application(application_id);
  const XnixRuntimeAcquisitionPreflightPolicy *policy =
    xnix_runtime_find_acquisition_preflight_policy(application_id);

  if (application == NULL || policy == NULL) {
    fprintf(stderr, "xnix-runtime-core: application is not registered in the C Runtime acquisition preflight policy catalog\n");
    return 64;
  }

  print_acquisition_preflight_policy_for_application(application, policy);
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
  fputs("\"snapshot_policy_owner\":\"c\",", stdout);
  fputs("\"snapshot_policy_count\":", stdout);
  printf("%zu,", xnix_runtime_snapshot_policy_count());
  fputs("\"state_root_policy_owner\":\"c\",", stdout);
  fputs("\"state_root_policy_count\":", stdout);
  printf("%zu,", xnix_runtime_state_root_policy_count());
  fputs("\"install_readiness_policy_owner\":\"c\",", stdout);
  fputs("\"install_readiness_policy_count\":", stdout);
  printf("%zu,", xnix_runtime_install_readiness_policy_count());
  fputs("\"artifact_manifest_policy_owner\":\"c\",", stdout);
  fputs("\"artifact_manifest_policy_count\":", stdout);
  printf("%zu,", xnix_runtime_artifact_manifest_policy_count());
  fputs("\"acquisition_preflight_policy_owner\":\"c\",", stdout);
  fputs("\"acquisition_preflight_policy_count\":", stdout);
  printf("%zu,", xnix_runtime_acquisition_preflight_policy_count());
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
  fputs("Usage: xnix-runtime-core {probe|list-applications|get-application APP_ID|list-engines|select-engine MODE|list-portal-policies|portal-policy APP_ID OPERATION|list-snapshot-policies|snapshot-policy APP_ID REASON|list-state-root-policies|state-root-policy APP_ID|list-install-readiness-policies|install-readiness-policy APP_ID ENVIRONMENT|list-artifact-manifest-policies|artifact-manifest-policy APP_ID|list-acquisition-preflight-policies|acquisition-preflight-policy APP_ID|write-gate METHOD}\n", stderr);
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

  if (strcmp(argv[1], "list-snapshot-policies") == 0 && argc == 2) {
    return print_snapshot_policy_catalog();
  }

  if (strcmp(argv[1], "snapshot-policy") == 0 && argc == 4) {
    return print_snapshot_policy(argv[2], argv[3]);
  }

  if (strcmp(argv[1], "list-state-root-policies") == 0 && argc == 2) {
    return print_state_root_policy_catalog();
  }

  if (strcmp(argv[1], "state-root-policy") == 0 && argc == 3) {
    return print_state_root_policy(argv[2]);
  }

  if (strcmp(argv[1], "list-install-readiness-policies") == 0 && argc == 2) {
    return print_install_readiness_policy_catalog();
  }

  if (strcmp(argv[1], "install-readiness-policy") == 0 && argc == 4) {
    return print_install_readiness_policy(argv[2], argv[3]);
  }

  if (strcmp(argv[1], "list-artifact-manifest-policies") == 0 && argc == 2) {
    return print_artifact_manifest_policy_catalog();
  }

  if (strcmp(argv[1], "artifact-manifest-policy") == 0 && argc == 3) {
    return print_artifact_manifest_policy(argv[2]);
  }

  if (strcmp(argv[1], "list-acquisition-preflight-policies") == 0 && argc == 2) {
    return print_acquisition_preflight_policy_catalog();
  }

  if (strcmp(argv[1], "acquisition-preflight-policy") == 0 && argc == 3) {
    return print_acquisition_preflight_policy(argv[2]);
  }

  if (strcmp(argv[1], "write-gate") == 0 && argc == 3) {
    return print_write_gate(argv[2]);
  }

  return usage();
}
