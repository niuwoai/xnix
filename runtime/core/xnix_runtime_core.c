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

static const XnixRuntimePortalPolicy portal_policies[] = {
  {
    .operation = "file-open",
    .portal_interface = "org.freedesktop.portal.FileChooser",
    .decision = "ask",
    .summary = "File access requires a user-approved desktop portal request.",
    .resources = {"documents", "downloads", "selected-files"},
    .resource_count = 3,
    .portal_required = true,
    .user_mediation_required = true,
    .direct_access_allowed = false,
    .runtime_policy_owner = true,
    .desktop_shell_policy_owner = false,
    .backend_details_exposed = false,
  },
  {
    .operation = "uri-open",
    .portal_interface = "org.freedesktop.portal.OpenURI",
    .decision = "ask",
    .summary = "URI handling requires a user-approved desktop portal request.",
    .resources = {"external-uri"},
    .resource_count = 1,
    .portal_required = true,
    .user_mediation_required = true,
    .direct_access_allowed = false,
    .runtime_policy_owner = true,
    .desktop_shell_policy_owner = false,
    .backend_details_exposed = false,
  },
  {
    .operation = "print",
    .portal_interface = "org.freedesktop.portal.Print",
    .decision = "ask",
    .summary = "Printing requires a user-approved desktop portal request.",
    .resources = {"printer"},
    .resource_count = 1,
    .portal_required = true,
    .user_mediation_required = true,
    .direct_access_allowed = false,
    .runtime_policy_owner = true,
    .desktop_shell_policy_owner = false,
    .backend_details_exposed = false,
  },
  {
    .operation = "screenshot",
    .portal_interface = "org.freedesktop.portal.Screenshot",
    .decision = "ask",
    .summary = "Screenshots require a user-approved desktop portal request.",
    .resources = {"screen"},
    .resource_count = 1,
    .portal_required = true,
    .user_mediation_required = true,
    .direct_access_allowed = false,
    .runtime_policy_owner = true,
    .desktop_shell_policy_owner = false,
    .backend_details_exposed = false,
  },
  {
    .operation = "clipboard",
    .portal_interface = "org.freedesktop.portal.Clipboard",
    .decision = "ask",
    .summary = "Clipboard access requires a user-approved desktop portal request.",
    .resources = {"clipboard"},
    .resource_count = 1,
    .portal_required = true,
    .user_mediation_required = true,
    .direct_access_allowed = false,
    .runtime_policy_owner = true,
    .desktop_shell_policy_owner = false,
    .backend_details_exposed = false,
  },
  {
    .operation = "camera",
    .portal_interface = "org.freedesktop.portal.Camera",
    .decision = "deny",
    .summary = "Camera access is denied until the user changes the application policy.",
    .resources = {"camera"},
    .resource_count = 1,
    .portal_required = true,
    .user_mediation_required = true,
    .direct_access_allowed = false,
    .runtime_policy_owner = true,
    .desktop_shell_policy_owner = false,
    .backend_details_exposed = false,
  },
  {
    .operation = "remote-desktop",
    .portal_interface = "org.freedesktop.portal.RemoteDesktop",
    .decision = "deny",
    .summary = "Remote desktop access is denied until the user changes the application policy.",
    .resources = {"screen", "input-devices"},
    .resource_count = 2,
    .portal_required = true,
    .user_mediation_required = true,
    .direct_access_allowed = false,
    .runtime_policy_owner = true,
    .desktop_shell_policy_owner = false,
    .backend_details_exposed = false,
  },
};

static const XnixRuntimeSnapshotPolicy snapshot_policies[] = {
  {
    .reason = "before-repair",
    .summary = "Create a restore point before a compatibility repair.",
    .enabled_by_default = true,
    .application_state = true,
    .runtime_metadata = true,
    .desktop_activation_receipts = true,
    .user_documents = false,
    .host_system = false,
    .restore_available = true,
    .restore_requires_user_confirmation = true,
    .restore_preserve_user_documents = true,
    .retention_policy = "bounded",
    .keep_latest = 5,
    .prune_automatically = true,
    .runtime_policy_owner = true,
    .desktop_shell_policy_owner = false,
    .backend_details_exposed = false,
    .host_root_modified = false,
  },
  {
    .reason = "before-engine-change",
    .summary = "Create a restore point before changing the compatibility engine.",
    .enabled_by_default = true,
    .application_state = true,
    .runtime_metadata = true,
    .desktop_activation_receipts = true,
    .user_documents = false,
    .host_system = false,
    .restore_available = true,
    .restore_requires_user_confirmation = true,
    .restore_preserve_user_documents = true,
    .retention_policy = "bounded",
    .keep_latest = 5,
    .prune_automatically = true,
    .runtime_policy_owner = true,
    .desktop_shell_policy_owner = false,
    .backend_details_exposed = false,
    .host_root_modified = false,
  },
  {
    .reason = "manual",
    .summary = "Create a user-requested restore point.",
    .enabled_by_default = true,
    .application_state = true,
    .runtime_metadata = true,
    .desktop_activation_receipts = true,
    .user_documents = false,
    .host_system = false,
    .restore_available = true,
    .restore_requires_user_confirmation = true,
    .restore_preserve_user_documents = true,
    .retention_policy = "bounded",
    .keep_latest = 5,
    .prune_automatically = true,
    .runtime_policy_owner = true,
    .desktop_shell_policy_owner = false,
    .backend_details_exposed = false,
    .host_root_modified = false,
  },
};

static const XnixRuntimeStateRootPolicy state_root_policies[] = {
  {
    .application_id = "org.xnix.sample.notepad",
    .state_namespace = "org.xnix.sample.notepad",
    .storage_scope = "per-application",
    .allocation_state = "planned",
    .managed_scopes = {
      "application-data",
      "runtime-metadata",
      "diagnostic-cache",
      "desktop-activation-receipts",
    },
    .managed_scope_summaries = {
      "Application-managed state is isolated under Runtime ownership.",
      "Runtime metadata tracks compatibility state without exposing host paths.",
      "Diagnostic cache is Runtime-owned and can be regenerated.",
      "Desktop activation receipts remain part of rollback planning.",
    },
    .managed_scope_snapshot_included = {
      true,
      true,
      false,
      true,
    },
    .managed_scope_count = 4,
    .blocked_actions = {
      "write application state outside Runtime ownership",
      "include user documents in state snapshots",
      "expose host storage paths to KDE",
      "restore state without user confirmation",
    },
    .blocked_action_count = 4,
    .retention_policy = "bounded",
    .automatic_restore_points = 5,
    .manual_restore_points = 10,
    .runtime_owned = true,
    .kde_policy_owner = false,
    .directories_created = false,
    .host_root_modified = false,
    .user_documents_included = false,
    .portal_required_for_user_files = true,
    .snapshot_eligible = true,
    .restore_requires_confirmation = true,
    .user_documents_excluded = true,
    .backend_details_exposed = false,
    .summary = "Runtime application state root is planned and isolated from user documents.",
  },
};

static const XnixRuntimeInstallReadinessPolicy install_readiness_policies[] = {
  {
    .application_id = "org.xnix.sample.notepad",
    .install_state = "planned",
    .selected_strategy = "automatic-managed",
    .blocked_actions = {
      "download artifacts before signed manifest verification",
      "install packages before Runtime install plan readiness",
      "stage desktop integration before recipe install gate approval",
      "launch backend before managed binding readiness",
      "expose backend commands or storage paths to KDE",
      "mutate the host root during install planning",
    },
    .blocked_action_count = 6,
    .phases = {
      {
        .id = "resolve-artifact-manifest",
        .status = "blocked",
        .summary = "Wait for a signed compatibility artifact manifest.",
      },
      {
        .id = "verify-artifact-digests",
        .status = "blocked",
        .summary = "Wait for digest verification before artifact activation.",
      },
      {
        .id = "prepare-package-source",
        .status = "blocked",
        .summary = "Wait for Runtime-owned package source readiness.",
      },
      {
        .id = "allocate-application-state",
        .status = "blocked",
        .summary = "Wait for Runtime-owned application state allocation.",
      },
      {
        .id = "stage-desktop-integration",
        .status = "pending",
        .summary = "Stage desktop artifacts only after install gates pass.",
      },
      {
        .id = "enable-launch-binding",
        .status = "blocked",
        .summary = "Backend launch binding stays disabled until install preflight completes.",
      },
    },
    .phase_count = 6,
    .artifact_manifest_ready = false,
    .artifact_signature_verified = false,
    .acquisition_ready = false,
    .package_source_ready = false,
    .state_root_allocated = false,
    .development_recipe_install_allowed = true,
    .production_recipe_install_allowed = false,
    .install_ready = false,
    .desktop_activation_ready = false,
    .download_enabled = false,
    .install_enabled = false,
    .network_request_created = false,
    .artifacts_downloaded = false,
    .host_root_modified = false,
    .privileged_container_required = false,
    .desktop_shell_command_exposed = false,
    .runtime_owned = true,
    .kde_policy_owner = false,
    .backend_details_exposed = false,
    .summary = "Compatibility install is planned and waiting for Runtime-owned readiness gates.",
  },
};

static const XnixRuntimeArtifactManifestPolicy artifact_manifest_policies[] = {
  {
    .application_id = "org.xnix.sample.notepad",
    .manifest_state = "planned",
    .selected_strategy = "automatic-managed",
    .artifact_groups = {
      {
        .id = "runtime-launch-metadata",
        .kind = "metadata",
        .cache_namespace = "org.xnix.sample.notepad.launch-metadata",
        .summary = "Runtime launch metadata must be resolved before launch binding.",
        .required = true,
        .resolved = false,
        .downloaded = false,
      },
      {
        .id = "local-execution-artifacts",
        .kind = "execution-artifacts",
        .cache_namespace = "org.xnix.sample.notepad.local-execution",
        .summary = "Local execution artifacts remain planned until signed manifest verification passes.",
        .required = true,
        .resolved = false,
        .downloaded = false,
      },
      {
        .id = "isolated-environment-artifacts",
        .kind = "environment-artifacts",
        .cache_namespace = "org.xnix.sample.notepad.isolated-environment",
        .summary = "Isolated environment artifacts remain optional until Runtime policy selects them.",
        .required = false,
        .resolved = false,
        .downloaded = false,
      },
    },
    .artifact_group_count = 3,
    .required_preflight = {
      {
        .id = "acquisition-preflight-ready",
        .status = "pending",
        .summary = "Runtime acquisition preflight must be ready before artifact manifest resolution.",
      },
      {
        .id = "manifest-signature-verification",
        .status = "required",
        .summary = "Runtime must verify the signed artifact manifest before artifact use.",
      },
      {
        .id = "artifact-digest-verification",
        .status = "required",
        .summary = "Runtime must verify artifact digests before cache activation.",
      },
      {
        .id = "cache-namespace-allocation",
        .status = "pending",
        .summary = "Runtime must allocate cache namespaces before artifact acquisition.",
      },
      {
        .id = "rollback-reference",
        .status = "required",
        .summary = "Runtime must record rollback references before artifacts can affect state.",
      },
    },
    .required_preflight_count = 5,
    .blocked_actions = {
      "download artifacts before signed manifest verification",
      "activate artifacts before digest verification",
      "expose artifact cache paths to KDE",
      "mutate host root during artifact manifest planning",
    },
    .blocked_action_count = 4,
    .runtime_owned = true,
    .kde_policy_owner = false,
    .manifest_ready = false,
    .signature_verified = false,
    .acquisition_preflight_ready = false,
    .download_enabled = false,
    .install_enabled = false,
    .network_request_created = false,
    .artifacts_downloaded = false,
    .host_root_modified = false,
    .privileged_container_required = false,
    .desktop_shell_command_exposed = false,
    .backend_details_exposed = false,
    .summary = "Compatibility artifact manifest is planned and waiting for acquisition preflight.",
  },
};

static const XnixRuntimeAcquisitionPreflightPolicy acquisition_preflight_policies[] = {
  {
    .application_id = "org.xnix.sample.notepad",
    .preflight_state = "planned",
    .selected_strategy = "automatic-managed",
    .checks = {
      {
        .id = "package-source-ready",
        .status = "pending",
        .summary = "Runtime package source selection must be ready before acquisition.",
      },
      {
        .id = "signed-artifact-manifest",
        .status = "required",
        .summary = "Runtime must verify a signed artifact manifest before acquisition.",
      },
      {
        .id = "runtime-cache-space",
        .status = "pending",
        .summary = "Runtime cache capacity must be checked before artifact acquisition.",
      },
      {
        .id = "network-policy-review",
        .status = "required",
        .summary = "Runtime must approve network policy before any acquisition request.",
      },
      {
        .id = "rollback-marker",
        .status = "required",
        .summary = "Runtime must define rollback markers before acquisition can change state.",
      },
    },
    .check_count = 5,
    .blocked_actions = {
      "download compatibility artifacts before acquisition preflight",
      "install compatibility artifacts before signed manifest verification",
      "invoke network access from KDE",
      "mutate host root during acquisition preflight",
    },
    .blocked_action_count = 4,
    .runtime_owned = true,
    .kde_policy_owner = false,
    .acquisition_ready = false,
    .download_enabled = false,
    .install_enabled = false,
    .network_required_for_planning = false,
    .network_request_created = false,
    .artifacts_downloaded = false,
    .host_root_modified = false,
    .privileged_container_required = false,
    .desktop_shell_command_exposed = false,
    .package_source_ready = false,
    .backend_details_exposed = false,
    .summary = "Compatibility acquisition preflight is planned and waiting for Runtime source readiness.",
  },
};

static const XnixRuntimePackageSourcePolicy package_source_policies[] = {
  {
    .application_id = "org.xnix.sample.notepad",
    .source_selection_state = "planned",
    .selected_strategy = "automatic-managed",
    .source_channels = {
      {
        .id = "os-managed-compatibility-packages",
        .kind = "distribution-packages",
        .selection_state = "planned",
        .supported_strategies = {
          "local-compatibility-engine",
        },
        .supported_strategy_count = 1,
        .runtime_owned = true,
        .summary = "Distribution-provided compatibility packages can be selected only through Runtime policy.",
      },
      {
        .id = "runtime-managed-toolcache",
        .kind = "runtime-cache",
        .selection_state = "planned",
        .supported_strategies = {
          "automatic-managed",
          "local-compatibility-engine",
          "isolated-compatibility-engine",
        },
        .supported_strategy_count = 3,
        .runtime_owned = true,
        .summary = "Runtime-managed tool cache keeps package selection outside desktop shell code.",
      },
      {
        .id = "isolated-environment-template-catalog",
        .kind = "template-catalog",
        .selection_state = "planned",
        .supported_strategies = {
          "isolated-compatibility-engine",
        },
        .supported_strategy_count = 1,
        .runtime_owned = true,
        .summary = "Isolated environment templates remain Runtime-owned and are not launched during planning.",
      },
    },
    .source_channel_count = 3,
    .required_preflight = {
      {
        .id = "signed-source-verification",
        .status = "required",
        .summary = "Runtime must verify a signed source before package installation is enabled.",
      },
      {
        .id = "source-policy-review",
        .status = "pending",
        .summary = "Runtime policy must select the package source before launch binding.",
      },
      {
        .id = "runtime-cache-quota",
        .status = "pending",
        .summary = "Runtime cache quota must be checked before package acquisition.",
      },
      {
        .id = "offline-fallback",
        .status = "pending",
        .summary = "Runtime must define the offline behavior before package acquisition.",
      },
    },
    .required_preflight_count = 4,
    .blocked_actions = {
      "install compatibility packages without Runtime source selection",
      "expose package manager commands to KDE",
      "use unsigned package sources",
      "mutate the host root during package-source planning",
    },
    .blocked_action_count = 4,
    .signed_source_required = true,
    .runtime_cache_required = true,
    .direct_desktop_install_allowed = false,
    .user_visible_backend_names = false,
    .host_package_manager_invoked = false,
    .runtime_owned = true,
    .kde_policy_owner = false,
    .package_source_ready = false,
    .install_enabled = false,
    .network_required_for_planning = false,
    .host_root_modified = false,
    .privileged_container_required = false,
    .desktop_shell_command_exposed = false,
    .backend_details_exposed = false,
    .summary = "Compatibility package source selection is planned and Runtime-owned.",
  },
};

static const XnixRuntimeBackendBindingPolicy backend_binding_policies[] = {
  {
    .application_id = "org.xnix.sample.notepad",
    .selected_strategy = "automatic-managed",
    .required_preflight = {
      {
        .id = "engine-package-source",
        .status = "pending",
        .summary = "A managed compatibility engine package source must be selected by the Runtime.",
      },
      {
        .id = "application-state-root",
        .status = "pending",
        .summary = "A Runtime-owned application state root must be allocated before launch binding.",
      },
      {
        .id = "portal-policy-review",
        .status = "required",
        .summary = "Portal access policy must be reviewed before file and desktop resource bridging.",
      },
      {
        .id = "snapshot-baseline",
        .status = "required",
        .summary = "A baseline restore point must exist before managed compatibility execution.",
      },
    },
    .required_preflight_count = 4,
    .blocked_actions = {
      "launch compatibility engine without managed binding",
      "expose backend command to desktop shell",
      "write application state outside Runtime ownership",
      "grant desktop resources without Portal policy",
    },
    .blocked_action_count = 4,
    .runtime_owned = true,
    .kde_policy_owner = false,
    .managed_binding_ready = false,
    .launch_enabled = false,
    .execution_request_created = false,
    .host_root_modified = false,
    .network_required = false,
    .privileged_container_required = false,
    .backend_details_exposed = false,
    .summary = "Managed compatibility backend binding is pending Runtime preflight.",
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

size_t
xnix_runtime_portal_policy_count(void)
{
  return sizeof(portal_policies) / sizeof(portal_policies[0]);
}

const XnixRuntimePortalPolicy *
xnix_runtime_portal_policy_at(size_t index)
{
  if (index >= xnix_runtime_portal_policy_count()) {
    return NULL;
  }

  return &portal_policies[index];
}

const XnixRuntimePortalPolicy *
xnix_runtime_find_portal_policy(const char *operation)
{
  if (operation == NULL) {
    return NULL;
  }

  for (size_t index = 0; index < xnix_runtime_portal_policy_count(); index++) {
    const XnixRuntimePortalPolicy *policy = xnix_runtime_portal_policy_at(index);

    if (policy != NULL && strcmp(operation, policy->operation) == 0) {
      return policy;
    }
  }

  return NULL;
}

size_t
xnix_runtime_snapshot_policy_count(void)
{
  return sizeof(snapshot_policies) / sizeof(snapshot_policies[0]);
}

const XnixRuntimeSnapshotPolicy *
xnix_runtime_snapshot_policy_at(size_t index)
{
  if (index >= xnix_runtime_snapshot_policy_count()) {
    return NULL;
  }

  return &snapshot_policies[index];
}

const XnixRuntimeSnapshotPolicy *
xnix_runtime_find_snapshot_policy(const char *reason)
{
  if (reason == NULL) {
    return NULL;
  }

  for (size_t index = 0; index < xnix_runtime_snapshot_policy_count(); index++) {
    const XnixRuntimeSnapshotPolicy *policy = xnix_runtime_snapshot_policy_at(index);

    if (policy != NULL && strcmp(reason, policy->reason) == 0) {
      return policy;
    }
  }

  return NULL;
}

size_t
xnix_runtime_state_root_policy_count(void)
{
  return sizeof(state_root_policies) / sizeof(state_root_policies[0]);
}

const XnixRuntimeStateRootPolicy *
xnix_runtime_state_root_policy_at(size_t index)
{
  if (index >= xnix_runtime_state_root_policy_count()) {
    return NULL;
  }

  return &state_root_policies[index];
}

const XnixRuntimeStateRootPolicy *
xnix_runtime_find_state_root_policy(const char *application_id)
{
  if (application_id == NULL) {
    return NULL;
  }

  for (size_t index = 0; index < xnix_runtime_state_root_policy_count(); index++) {
    const XnixRuntimeStateRootPolicy *policy = xnix_runtime_state_root_policy_at(index);

    if (policy != NULL && strcmp(application_id, policy->application_id) == 0) {
      return policy;
    }
  }

  return NULL;
}

size_t
xnix_runtime_install_readiness_policy_count(void)
{
  return sizeof(install_readiness_policies) / sizeof(install_readiness_policies[0]);
}

const XnixRuntimeInstallReadinessPolicy *
xnix_runtime_install_readiness_policy_at(size_t index)
{
  if (index >= xnix_runtime_install_readiness_policy_count()) {
    return NULL;
  }

  return &install_readiness_policies[index];
}

const XnixRuntimeInstallReadinessPolicy *
xnix_runtime_find_install_readiness_policy(const char *application_id)
{
  if (application_id == NULL) {
    return NULL;
  }

  for (size_t index = 0; index < xnix_runtime_install_readiness_policy_count(); index++) {
    const XnixRuntimeInstallReadinessPolicy *policy =
      xnix_runtime_install_readiness_policy_at(index);

    if (policy != NULL && strcmp(application_id, policy->application_id) == 0) {
      return policy;
    }
  }

  return NULL;
}

bool
xnix_runtime_install_readiness_allows_recipe_install(
  const XnixRuntimeInstallReadinessPolicy *policy,
  const char *environment
)
{
  if (policy == NULL || environment == NULL) {
    return false;
  }

  if (strcmp(environment, "development") == 0) {
    return policy->development_recipe_install_allowed;
  }
  if (strcmp(environment, "production") == 0) {
    return policy->production_recipe_install_allowed;
  }

  return false;
}

size_t
xnix_runtime_artifact_manifest_policy_count(void)
{
  return sizeof(artifact_manifest_policies) / sizeof(artifact_manifest_policies[0]);
}

const XnixRuntimeArtifactManifestPolicy *
xnix_runtime_artifact_manifest_policy_at(size_t index)
{
  if (index >= xnix_runtime_artifact_manifest_policy_count()) {
    return NULL;
  }

  return &artifact_manifest_policies[index];
}

const XnixRuntimeArtifactManifestPolicy *
xnix_runtime_find_artifact_manifest_policy(const char *application_id)
{
  if (application_id == NULL) {
    return NULL;
  }

  for (size_t index = 0; index < xnix_runtime_artifact_manifest_policy_count(); index++) {
    const XnixRuntimeArtifactManifestPolicy *policy =
      xnix_runtime_artifact_manifest_policy_at(index);

    if (policy != NULL && strcmp(application_id, policy->application_id) == 0) {
      return policy;
    }
  }

  return NULL;
}

size_t
xnix_runtime_acquisition_preflight_policy_count(void)
{
  return sizeof(acquisition_preflight_policies) / sizeof(acquisition_preflight_policies[0]);
}

const XnixRuntimeAcquisitionPreflightPolicy *
xnix_runtime_acquisition_preflight_policy_at(size_t index)
{
  if (index >= xnix_runtime_acquisition_preflight_policy_count()) {
    return NULL;
  }

  return &acquisition_preflight_policies[index];
}

const XnixRuntimeAcquisitionPreflightPolicy *
xnix_runtime_find_acquisition_preflight_policy(const char *application_id)
{
  if (application_id == NULL) {
    return NULL;
  }

  for (size_t index = 0; index < xnix_runtime_acquisition_preflight_policy_count(); index++) {
    const XnixRuntimeAcquisitionPreflightPolicy *policy =
      xnix_runtime_acquisition_preflight_policy_at(index);

    if (policy != NULL && strcmp(application_id, policy->application_id) == 0) {
      return policy;
    }
  }

  return NULL;
}

size_t
xnix_runtime_package_source_policy_count(void)
{
  return sizeof(package_source_policies) / sizeof(package_source_policies[0]);
}

const XnixRuntimePackageSourcePolicy *
xnix_runtime_package_source_policy_at(size_t index)
{
  if (index >= xnix_runtime_package_source_policy_count()) {
    return NULL;
  }

  return &package_source_policies[index];
}

const XnixRuntimePackageSourcePolicy *
xnix_runtime_find_package_source_policy(const char *application_id)
{
  if (application_id == NULL) {
    return NULL;
  }

  for (size_t index = 0; index < xnix_runtime_package_source_policy_count(); index++) {
    const XnixRuntimePackageSourcePolicy *policy =
      xnix_runtime_package_source_policy_at(index);

    if (policy != NULL && strcmp(application_id, policy->application_id) == 0) {
      return policy;
    }
  }

  return NULL;
}

size_t
xnix_runtime_backend_binding_policy_count(void)
{
  return sizeof(backend_binding_policies) / sizeof(backend_binding_policies[0]);
}

const XnixRuntimeBackendBindingPolicy *
xnix_runtime_backend_binding_policy_at(size_t index)
{
  if (index >= xnix_runtime_backend_binding_policy_count()) {
    return NULL;
  }

  return &backend_binding_policies[index];
}

const XnixRuntimeBackendBindingPolicy *
xnix_runtime_find_backend_binding_policy(const char *application_id)
{
  if (application_id == NULL) {
    return NULL;
  }

  for (size_t index = 0; index < xnix_runtime_backend_binding_policy_count(); index++) {
    const XnixRuntimeBackendBindingPolicy *policy =
      xnix_runtime_backend_binding_policy_at(index);

    if (policy != NULL && strcmp(application_id, policy->application_id) == 0) {
      return policy;
    }
  }

  return NULL;
}
