#include "xnix_runtime_core.h"

#include <ctype.h>
#include <stdio.h>
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

static const XnixRuntimeSettingsPolicy settings_policies[] = {
  {
    .application_id = "org.xnix.sample.notepad",
    .settings_state = "planned",
    .sections = {
      {
        .id = "run-mode",
        .title = "Run mode",
        .description = "Choose how Xnix balances speed and compatibility.",
        .fields = {
          {
            .id = "mode",
            .label = "Run mode",
            .value = "automatic",
            .options = {"automatic", "performance", "compatibility"},
            .option_count = 3,
          },
          {
            .id = "preference",
            .label = "Priority",
            .value = "compatibility",
            .options = {"performance", "compatibility"},
            .option_count = 2,
          },
        },
        .field_count = 2,
      },
      {
        .id = "resource-access",
        .title = "File access",
        .description = "Control which user folders this application may request.",
        .fields = {
          {
            .id = "documents",
            .label = "Documents",
            .value = "ask",
            .options = {"allow", "ask", "deny"},
            .option_count = 3,
          },
          {
            .id = "downloads",
            .label = "Downloads",
            .value = "ask",
            .options = {"allow", "ask", "deny"},
            .option_count = 3,
          },
        },
        .field_count = 2,
      },
      {
        .id = "devices",
        .title = "Devices",
        .description = "Control sensitive device access.",
        .fields = {
          {
            .id = "camera",
            .label = "Camera",
            .value = "deny",
            .options = {"allow", "ask", "deny"},
            .option_count = 3,
          },
        },
        .field_count = 1,
      },
      {
        .id = "network",
        .title = "Network",
        .description = "Control network access for compatibility actions.",
        .fields = {
          {
            .id = "network",
            .label = "Network",
            .value = "allow",
            .options = {"allow", "ask", "deny"},
            .option_count = 3,
          },
        },
        .field_count = 1,
      },
      {
        .id = "snapshots",
        .title = "Snapshots",
        .description = "Keep restore points before risky compatibility changes.",
        .fields = {
          {
            .id = "snapshots",
            .label = "Environment snapshots",
            .value = "enabled",
            .options = {"enabled", "disabled"},
            .option_count = 2,
          },
        },
        .field_count = 1,
      },
    },
    .section_count = 5,
    .runtime_owned = true,
    .kde_policy_owner = false,
    .settings_persisted = false,
    .host_root_modified = false,
    .backend_details_exposed = false,
    .summary = "Compatibility settings are modeled by the Runtime and ready for KDE display.",
  },
};

static const XnixRuntimeSettingsChangePolicy settings_change_policies[] = {
  {
    .application_id = "org.xnix.sample.notepad",
    .section_id = "resource-access",
    .field_id = "documents",
    .requested_value = "allow",
    .change_state = "planned",
    .affected_policy = {
      .section = "resource-access",
      .field = "documents",
      .value = "allow",
      .options = {"allow", "ask", "deny"},
      .option_count = 3,
    },
    .steps = {
      {
        .id = "validate-setting",
        .status = "pass",
        .summary = "The requested settings value is valid for the Runtime settings schema.",
      },
      {
        .id = "review-user-confirmation",
        .status = "required",
        .summary = "KDE must present the change for user review before persistence.",
      },
      {
        .id = "review-portal-policy",
        .status = "required",
        .summary = "Runtime Portal policy must be reviewed before desktop resource access changes.",
      },
      {
        .id = "prepare-restore-point",
        .status = "pass",
        .summary = "Runtime should prepare a restore point before risky compatibility settings changes.",
      },
      {
        .id = "persist-runtime-setting",
        .status = "pending",
        .summary = "Runtime persistence is not enabled in this version.",
      },
    },
    .step_count = 5,
    .blocked_actions = {
      "persist compatibility settings before Runtime confirmation",
      "grant desktop resources without Portal policy review",
      "modify host root while planning settings changes",
      "expose backend implementation settings to KDE",
    },
    .blocked_action_count = 4,
    .runtime_owned = true,
    .kde_policy_owner = false,
    .apply_enabled = false,
    .settings_persisted = false,
    .host_root_modified = false,
    .backend_details_exposed = false,
    .user_confirmation_required = true,
    .snapshot_recommended = false,
    .portal_policy_review_required = true,
    .runtime_restart_required = false,
    .summary = "Compatibility settings change is planned and waiting for Runtime persistence support.",
  },
};

static const XnixRuntimeServiceBindingPolicy service_binding_policy = {
  .activation = {
    .dbus_service_file = "runtime/dbus/org.xnix.Compatibility1.service",
    .systemd_unit = "runtime/systemd/xnix-compatd.service",
    .libexec_wrapper = "libexec/xnix/compatd",
    .dbus_contract = "runtime/dbus/org.xnix.Compatibility1.xml",
    .packaged_wrapper = "/usr/libexec/xnix/compatd",
  },
  .checks = {
    {
      .id = "dbus-service-activation",
      .status = "pass",
      .summary = "D-Bus activation points to the packaged Runtime wrapper.",
    },
    {
      .id = "systemd-service-hardening",
      .status = "pass",
      .summary = "systemd activation uses the stable bus name and hardened service settings.",
    },
    {
      .id = "libexec-wrapper",
      .status = "pass",
      .summary = "Packaged Runtime wrapper delegates to the Runtime daemon entry point.",
    },
    {
      .id = "dbus-contract",
      .status = "pass",
      .summary = "D-Bus contract exposes read-only Runtime service binding status.",
    },
    {
      .id = "live-dbus-owner",
      .status = "pending",
      .summary = "A long-running production D-Bus owner is still pending.",
    },
  },
  .check_count = 5,
  .counts = {
    .total = 5,
    .passed = 4,
    .pending = 1,
    .blocked = 0,
  },
  .runtime_owned = true,
  .kde_policy_owner = false,
  .activation_binding_ready = true,
  .live_dbus_owner_ready = false,
  .smoke_adapter_available = true,
  .network_required = false,
  .host_root_modified = false,
  .privileged_container_required = false,
  .backend_details_exposed = false,
  .summary = "Runtime service activation files are aligned; live D-Bus ownership remains pending.",
};

static const XnixRuntimeLiveOwnerGatePolicy live_owner_gate_policy = {
  .required_gates = {
    {
      .id = "activation-binding",
      .status = "pass",
      .summary = "D-Bus activation files, systemd unit, libexec wrapper, and contract must stay aligned.",
    },
    {
      .id = "long-running-runtime-owner",
      .status = "pending",
      .summary = "The Runtime needs a packaged long-running process that owns the stable bus name.",
    },
    {
      .id = "bus-name-acquisition",
      .status = "pending",
      .summary = "Production smoke must prove the packaged Runtime owns org.xnix.Compatibility1.",
    },
    {
      .id = "read-only-method-parity",
      .status = "pending",
      .summary = "The live owner must answer the same read-only planning methods as the smoke adapter.",
    },
    {
      .id = "production-recipe-trust",
      .status = "pending",
      .summary = "Production ownership must be gated by signed recipe validation instead of development registry trust.",
    },
  },
  .required_gate_count = 5,
  .blocked_reasons = {
    "Do not treat the D-Bus smoke adapter as the production Runtime owner.",
    "Do not let KDE own Runtime policy or bus-name readiness decisions.",
    "Do not enable launch, install, repair, restore, or settings persistence from this gate.",
    "Do not mutate the host root while evaluating live-owner readiness.",
    "Do not expose compatibility backend implementation details in live-owner readiness.",
  },
  .blocked_reason_count = 5,
  .runtime_owned = true,
  .kde_policy_owner = false,
  .activation_binding_ready = true,
  .live_dbus_owner_ready = false,
  .production_owner_enabled = false,
  .owner_transition_ready = false,
  .smoke_adapter_available = true,
  .smoke_adapter_is_production_owner = false,
  .kde_may_claim_runtime_ownership = false,
  .network_required = false,
  .host_root_modified = false,
  .privileged_container_required = false,
  .backend_details_exposed = false,
  .summary = "Runtime activation files are aligned, but production D-Bus ownership remains gated.",
};

static const XnixRuntimeOwnerSmokePlanPolicy owner_smoke_plan_policy = {
  .steps = {
    {
      .id = "validate-activation-files",
      .status = "pass",
      .summary = "Verify D-Bus service activation, systemd hardening, libexec wrapper, and contract alignment.",
    },
    {
      .id = "start-packaged-runtime-owner",
      .status = "pending",
      .summary = "Start the packaged Runtime owner in an isolated session without modifying the host root.",
    },
    {
      .id = "assert-stable-bus-name",
      .status = "pending",
      .summary = "Prove the packaged Runtime owner owns org.xnix.Compatibility1 on the test bus.",
    },
    {
      .id = "check-read-only-method-parity",
      .status = "pending",
      .summary = "Call the read-only planning methods required by KDE and compare them with the contract.",
    },
    {
      .id = "reject-write-methods",
      .status = "pending",
      .summary = "Confirm launch, install, snapshot, restore, repair, and settings persistence remain gated.",
    },
    {
      .id = "verify-non-production-smoke-adapter-boundary",
      .status = "pending",
      .summary = "Confirm the smoke adapter is never accepted as a production Runtime owner.",
    },
    {
      .id = "report-kde-safe-summary",
      .status = "pending",
      .summary = "Return a Compatibility Center summary without backend details or host paths.",
    },
  },
  .step_count = 7,
  .counts = {
    .total = 7,
    .passed = 1,
    .pending = 6,
    .blocked = 0,
  },
  .blocked_actions = {
    "Do not start a host system service from the smoke plan.",
    "Do not claim the production Runtime bus name from the smoke adapter.",
    "Do not enable backend launch, install, repair, restore, or settings persistence.",
    "Do not mutate the host root while planning owner smoke.",
    "Do not expose backend implementation details or host paths to KDE.",
  },
  .blocked_action_count = 5,
  .runtime_owned = true,
  .kde_policy_owner = false,
  .activation_binding_ready = true,
  .live_dbus_owner_ready = false,
  .production_owner_enabled = false,
  .owner_transition_ready = false,
  .smoke_state = "planned",
  .smoke_environment = "restricted-session",
  .network_required = false,
  .host_root_modified = false,
  .privileged_container_required = false,
  .system_service_started = false,
  .production_bus_claimed = false,
  .backend_details_exposed = false,
  .summary = "Runtime owner smoke is planned; production bus ownership remains disabled until all gates pass.",
};

static const XnixRuntimeMethodParityManifestPolicy method_parity_manifest_policy = {
  .read_only_methods = {
    "ListApplications",
    "GetApplication",
    "GetDiagnostics",
    "GetEngineCatalog",
    "GetRunPlan",
    "GetDesktopActivationManifest",
    "GetDesktopEntryPlan",
    "GetTaskManagerIdentityPlan",
    "GetKDEIntegrationStatus",
    "GetKDEShellIntegrationPlan",
    "GetKDEApplicationSurfacePlan",
    "GetDesktopResourceBridgePlan",
    "GetFileAssociationPlan",
    "GetNotificationPlan",
    "GetTrayStatus",
    "GetPortalRequestPlan",
    "GetApplicationStateRoot",
    "GetCompatibilityPackageSource",
    "GetCompatibilityAcquisitionPreflight",
    "GetCompatibilityArtifactManifest",
    "GetCompatibilityInstallPlan",
    "GetBackendBinding",
    "GetBackendCapabilityMatrix",
    "GetBackendSelectionPlan",
    "GetBackendLifecycle",
    "GetBackendEnvironmentPlan",
    "GetRepairPlan",
    "GetTestPlan",
    "GetTestResult",
    "GetExecutionReadiness",
    "GetLaunchIntent",
    "GetAIDiagnosticInput",
    "GetAIDiagnosticRecommendation",
    "GetAIRepairApprovalGate",
    "GetSnapshotPlan",
    "GetPortalAccessPolicy",
    "GetRuntimeServiceBinding",
    "GetRuntimeLiveOwnerGate",
    "GetRuntimeOwnerSmokePlan",
    "GetRuntimeMethodParityManifest",
    "GetRuntimeWriteGate",
    "GetCompatibilitySettings",
    "GetCompatibilitySettingsChangePlan",
    "GetCompatibilityModeSwitchPlan",
    "GetCompatibilityPermissionReviewPlan",
    "GetCompatibilityReviewFlowPlan",
    "GetCompatibilityActionQueue",
    "GetCompatibilityActionReviewReceipt",
    "GetCompatibilityCenterSummary",
    "GetKDECenterPage",
    "GetKDECenterPageSections",
    "GetKDECenterPageSectionDetail",
    "GetKRunnerQueryPlan",
    "GetKWinWindowRulePlan",
  },
  .read_only_method_count = 54,
  .parity_checks = {
    {
      .id = "dbus-contract",
      .status = "pass",
      .method_count = 54,
      .missing_method_count = 0,
      .summary = "dbus-contract covers all Runtime read-only methods.",
    },
    {
      .id = "runtime-dispatch",
      .status = "pass",
      .method_count = 54,
      .missing_method_count = 0,
      .summary = "runtime-dispatch covers all Runtime read-only methods.",
    },
    {
      .id = "dbus-client",
      .status = "pass",
      .method_count = 54,
      .missing_method_count = 0,
      .summary = "dbus-client covers all Runtime read-only methods.",
    },
    {
      .id = "smoke-adapter",
      .status = "pass",
      .method_count = 54,
      .missing_method_count = 0,
      .summary = "smoke-adapter covers all Runtime read-only methods.",
    },
    {
      .id = "session-smoke",
      .status = "pass",
      .method_count = 54,
      .missing_method_count = 0,
      .summary = "session-smoke covers all Runtime read-only methods.",
    },
  },
  .parity_check_count = 5,
  .counts = {
    .total = 5,
    .passed = 5,
    .blocked = 0,
    .pending = 0,
  },
  .write_methods = {
    "InstallRecipe",
    "Launch",
    "CreateSnapshot",
    "RestoreSnapshot",
  },
  .write_method_count = 4,
  .runtime_owned = true,
  .kde_policy_owner = false,
  .read_only_method_parity_ready = true,
  .write_methods_supported = false,
  .write_method_dispatch_enabled = false,
  .network_required = false,
  .host_root_modified = false,
  .privileged_container_required = false,
  .backend_details_exposed = false,
  .summary = "Runtime read-only D-Bus method parity is ready for owner smoke.",
};

static const XnixRuntimeRecipeTrustPolicy recipe_trust_policy = {
  .decision = "development-only",
  .recipe_count = 1,
  .checks = {
    {
      .id = "registry.digest",
      .status = "pass",
      .message = "Recipe digests are verified against registry metadata.",
    },
    {
      .id = "registry.signature",
      .status = "pending",
      .message = "Production signed recipe validation is required before external recipes are trusted.",
    },
    {
      .id = "registry.development",
      .status = "warn",
      .message = "Development-only recipes are allowed for local testing but not for production trust.",
    },
  },
  .check_count = 3,
  .blocking_reasons = {
    "production signed recipe validation is not enabled",
    "registry contains development-only recipes",
  },
  .blocking_reason_count = 2,
  .next_requirements = {
    "Add production trust roots for recipe signatures.",
    "Require verified signatures before accepting external recipe registries.",
    "Bind recipe trust decisions to install and activation requests.",
  },
  .next_requirement_count = 3,
  .runtime_owned = true,
  .kde_policy_owner = false,
  .digest_verified = true,
  .signed_recipe_validation = false,
  .development_registry = true,
  .production_trusted = false,
  .development_staging_allowed = true,
  .production_install_allowed = false,
  .host_root_modified = false,
  .backend_details_exposed = false,
  .summary = "Recipes are verified for local development only.",
};

static const char *const production_recipe_install_requirements[] = {
  "Use a registry with a production signed source and verified recipe signatures.",
  "Keep SHA-256 digest verification enabled before activation.",
  "Do not promote development-only recipes into production installation.",
};

static const char *const development_recipe_install_requirements[] = {
  "Register the application recipe in the selected registry.",
  "Keep SHA-256 digest verification enabled for development staging.",
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

size_t
xnix_runtime_settings_policy_count(void)
{
  return sizeof(settings_policies) / sizeof(settings_policies[0]);
}

const XnixRuntimeSettingsPolicy *
xnix_runtime_settings_policy_at(size_t index)
{
  if (index >= xnix_runtime_settings_policy_count()) {
    return NULL;
  }

  return &settings_policies[index];
}

const XnixRuntimeSettingsPolicy *
xnix_runtime_find_settings_policy(const char *application_id)
{
  if (application_id == NULL) {
    return NULL;
  }

  for (size_t index = 0; index < xnix_runtime_settings_policy_count(); index++) {
    const XnixRuntimeSettingsPolicy *policy = xnix_runtime_settings_policy_at(index);

    if (policy != NULL && strcmp(application_id, policy->application_id) == 0) {
      return policy;
    }
  }

  return NULL;
}

size_t
xnix_runtime_settings_change_policy_count(void)
{
  return sizeof(settings_change_policies) / sizeof(settings_change_policies[0]);
}

const XnixRuntimeSettingsChangePolicy *
xnix_runtime_settings_change_policy_at(size_t index)
{
  if (index >= xnix_runtime_settings_change_policy_count()) {
    return NULL;
  }

  return &settings_change_policies[index];
}

const XnixRuntimeSettingsChangePolicy *
xnix_runtime_find_settings_change_policy(const char *application_id)
{
  if (application_id == NULL) {
    return NULL;
  }

  for (size_t index = 0; index < xnix_runtime_settings_change_policy_count(); index++) {
    const XnixRuntimeSettingsChangePolicy *policy =
      xnix_runtime_settings_change_policy_at(index);

    if (policy != NULL && strcmp(application_id, policy->application_id) == 0) {
      return policy;
    }
  }

  return NULL;
}

const XnixRuntimeServiceBindingPolicy *
xnix_runtime_service_binding_policy(void)
{
  return &service_binding_policy;
}

const XnixRuntimeLiveOwnerGatePolicy *
xnix_runtime_live_owner_gate_policy(void)
{
  return &live_owner_gate_policy;
}

const XnixRuntimeOwnerSmokePlanPolicy *
xnix_runtime_owner_smoke_plan_policy(void)
{
  return &owner_smoke_plan_policy;
}

const XnixRuntimeMethodParityManifestPolicy *
xnix_runtime_method_parity_manifest_policy(void)
{
  return &method_parity_manifest_policy;
}

const XnixRuntimeRecipeTrustPolicy *
xnix_runtime_recipe_trust_policy(void)
{
  return &recipe_trust_policy;
}

bool
xnix_runtime_recipe_install_gate(
  const char *application_id,
  const char *mode,
  XnixRuntimeRecipeInstallGate *gate
)
{
  bool development_mode = false;
  bool production_mode = false;
  bool matched = false;

  if (application_id == NULL || mode == NULL || gate == NULL) {
    return false;
  }

  development_mode = strcmp(mode, "development") == 0;
  production_mode = strcmp(mode, "production") == 0;

  if (!development_mode && !production_mode) {
    return false;
  }

  matched = strcmp(application_id, "org.xnix.sample.notepad") == 0;

  gate->application_id = application_id;
  gate->mode = mode;
  gate->policy_decision = recipe_trust_policy.decision;
  gate->matched_recipe.id = matched ? "org.xnix.sample.notepad" : NULL;
  gate->matched_recipe.signature_status = matched ? "development" : NULL;
  gate->matched_recipe.digest_verified = matched;
  gate->matched_recipe_present = matched;
  gate->runtime_owned = true;
  gate->kde_policy_owner = false;
  gate->request_object_created = false;
  gate->install_started = false;
  gate->host_root_modified = false;
  gate->backend_details_exposed = false;
  gate->development_staging_allowed = matched && development_mode;
  gate->production_install_allowed = false;
  gate->blocking_reason_count = 0;
  gate->requirement_count = 0;
  gate->summary = "Recipe installation is blocked until trust requirements are satisfied.";

  if (!matched) {
    gate->blocking_reasons[gate->blocking_reason_count++] = "recipe is not registered";
  }

  if (matched && !gate->matched_recipe.digest_verified) {
    gate->blocking_reasons[gate->blocking_reason_count++] = "recipe digest is not verified";
  }

  if (production_mode || !matched) {
    for (size_t index = 0; index < recipe_trust_policy.blocking_reason_count; index++) {
      gate->blocking_reasons[gate->blocking_reason_count++] =
        recipe_trust_policy.blocking_reasons[index];
    }
  }

  if (development_mode && matched && gate->matched_recipe.digest_verified) {
    gate->blocking_reason_count = 0;
  }

  if (gate->blocking_reason_count == 0) {
    gate->decision = "allow";
    gate->summary = "Recipe installation may continue.";
    return true;
  }

  gate->decision = "block";

  if (production_mode) {
    for (size_t index = 0;
         index < sizeof(production_recipe_install_requirements) /
                   sizeof(production_recipe_install_requirements[0]);
         index++) {
      gate->requirements[gate->requirement_count++] =
        production_recipe_install_requirements[index];
    }
  } else {
    for (size_t index = 0;
         index < sizeof(development_recipe_install_requirements) /
                   sizeof(development_recipe_install_requirements[0]);
         index++) {
      gate->requirements[gate->requirement_count++] =
        development_recipe_install_requirements[index];
    }
  }

  return true;
}

bool
xnix_runtime_compatibility_install_plan(
  const char *application_id,
  const char *environment,
  XnixRuntimeCompatibilityInstallPlan *plan
)
{
  const XnixRuntimeApplication *application = NULL;
  const XnixRuntimeInstallReadinessPolicy *install_readiness = NULL;

  if (application_id == NULL || environment == NULL || plan == NULL) {
    return false;
  }

  if (strcmp(environment, "development") != 0 &&
      strcmp(environment, "production") != 0) {
    return false;
  }

  application = xnix_runtime_find_application(application_id);
  install_readiness = xnix_runtime_find_install_readiness_policy(application_id);

  if (application == NULL || install_readiness == NULL) {
    return false;
  }

  if (!xnix_runtime_recipe_install_gate(
        application_id,
        environment,
        &plan->recipe_install_gate
      )) {
    return false;
  }

  plan->application = application;
  plan->install_readiness = install_readiness;
  plan->environment = environment;
  plan->runtime_owned = true;
  plan->kde_policy_owner = false;
  return true;
}

static XnixRuntimeCompatibilityAction
compatibility_action(
  const char *id,
  const char *source_type,
  const char *status,
  const char *priority,
  const char *title,
  const char *summary,
  bool user_review_required,
  const char *runtime_gate,
  const char *next_step
)
{
  XnixRuntimeCompatibilityAction action;

  action.id = id;
  action.source_type = source_type;
  action.status = status;
  action.priority = priority;
  action.title = title;
  action.summary = summary;
  action.runtime_gate = runtime_gate;
  action.next_step = next_step;
  action.user_review_required = user_review_required;
  action.execution_enabled = false;
  action.backend_details_exposed = false;

  return action;
}

bool
xnix_runtime_compatibility_action_queue(
  const char *application_id,
  XnixRuntimeCompatibilityActionQueue *queue
)
{
  const XnixRuntimeApplication *application = NULL;
  const XnixRuntimeInstallReadinessPolicy *install_readiness = NULL;
  const XnixRuntimeSettingsChangePolicy *settings_change = NULL;
  const XnixRuntimeServiceBindingPolicy *service_binding = NULL;
  const XnixRuntimePortalPolicy *portal_policy = NULL;

  if (application_id == NULL || queue == NULL) {
    return false;
  }

  application = xnix_runtime_find_application(application_id);
  install_readiness = xnix_runtime_find_install_readiness_policy(application_id);
  settings_change = xnix_runtime_find_settings_change_policy(application_id);
  service_binding = xnix_runtime_service_binding_policy();
  portal_policy = xnix_runtime_find_portal_policy("file-open");

  if (application == NULL || install_readiness == NULL ||
      settings_change == NULL || service_binding == NULL ||
      portal_policy == NULL) {
    return false;
  }

  queue->application = application;
  queue->actions[0] = compatibility_action(
    "review-install-readiness",
    "compatibility-install-plan",
    install_readiness->install_ready ? "pass" : "blocked",
    "high",
    "Review install readiness",
    install_readiness->summary,
    false,
    "install-plan-readiness",
    "Wait for signed artifacts, package source readiness, state allocation, and recipe install gate approval."
  );
  queue->actions[1] = compatibility_action(
    "review-settings-change",
    "compatibility-settings-change-plan",
    settings_change->apply_enabled ? "pass" : "review-required",
    "medium",
    "Review compatibility settings change",
    settings_change->summary,
    settings_change->user_confirmation_required,
    "runtime-settings-persistence",
    "Keep the change pending until Runtime persistence support and Portal policy review are available."
  );
  queue->actions[2] = compatibility_action(
    "review-ai-repair",
    "ai-repair-approval-gate",
    "approval-required",
    "high",
    "Review AI repair recommendation",
    "AI repair recommendations are blocked from execution until Runtime approval gates pass.",
    true,
    "ai-repair-approval",
    "Require Compatibility Center review, Runtime approval, and restore-point preflight before repair execution."
  );
  queue->actions[3] = compatibility_action(
    "verify-runtime-service",
    "runtime-service-binding",
    service_binding->live_dbus_owner_ready ? "pass" : "pending",
    "medium",
    "Verify Runtime service ownership",
    service_binding->summary,
    false,
    "live-dbus-owner",
    "Keep using the session-bus smoke adapter until production D-Bus ownership is enabled."
  );
  queue->actions[4] = compatibility_action(
    "review-portal-policy",
    "portal-access-policy",
    portal_policy->portal_required ? "review-required" : "pass",
    "medium",
    "Review desktop resource policy",
    portal_policy->summary,
    portal_policy->user_mediation_required,
    "portal-policy-review",
    "Use XDG Desktop Portal approval before granting file access to compatibility applications."
  );
  queue->action_count = 5;
  queue->pending_action_count = 0;
  queue->user_review_required_count = 0;
  for (size_t index = 0; index < queue->action_count; index++) {
    if (strcmp(queue->actions[index].status, "pass") != 0) {
      queue->pending_action_count++;
    }
    if (queue->actions[index].user_review_required) {
      queue->user_review_required_count++;
    }
  }
  queue->blocked_actions[0] = "execute queued actions from KDE without Runtime approval";
  queue->blocked_actions[1] = "persist settings while action queue execution is disabled";
  queue->blocked_actions[2] = "start compatibility backends from the Compatibility Center";
  queue->blocked_actions[3] = "grant desktop resources without XDG Desktop Portal review";
  queue->blocked_actions[4] = "mutate the host root from Compatibility Center actions";
  queue->blocked_actions[5] = "expose backend implementation details in action cards";
  queue->blocked_action_count = 6;
  queue->runtime_owned = true;
  queue->kde_policy_owner = false;
  queue->execution_enabled = false;
  queue->repair_execution_enabled = false;
  queue->settings_persistence_enabled = false;
  queue->host_root_modified = false;
  queue->network_required = false;
  queue->backend_details_exposed = false;
  queue->summary = "Compatibility Center actions are queued for review and cannot execute until Runtime gates are implemented.";

  return true;
}

#include "xnix_runtime_core_action_review.inc"
#include "xnix_runtime_core_compatibility_center.inc"
#include "xnix_runtime_core_run_plan.inc"
#include "xnix_runtime_core_desktop_activation.inc"
#include "xnix_runtime_core_kde_integration_status.inc"
#include "xnix_runtime_core_kde_application_surface_plan.inc"
#include "xnix_runtime_core_desktop_resource_bridge_plan.inc"
#include "xnix_runtime_core_compatibility_mode_switch_plan.inc"
#include "xnix_runtime_core_compatibility_permission_review_plan.inc"
#include "xnix_runtime_core_desktop_entry.inc"
#include "xnix_runtime_core_task_manager_identity.inc"
#include "xnix_runtime_core_kwin_window_rule.inc"
#include "xnix_runtime_core_file_association.inc"
#include "xnix_runtime_core_notification.inc"
#include "xnix_runtime_core_tray_status.inc"
#include "xnix_runtime_core_krunner_query.inc"
#include "xnix_runtime_core_portal_request.inc"
#include "xnix_runtime_core_kde_shell_integration_plan.inc"
#include "xnix_runtime_core_compatibility_review_flow_plan.inc"
#include "xnix_runtime_core_compatibility_snapshot_plan.inc"
#include "xnix_runtime_core_compatibility_test_plan.inc"
#include "xnix_runtime_core_compatibility_test_result.inc"
#include "xnix_runtime_core_execution_readiness.inc"
#include "xnix_runtime_core_launch_intent.inc"
#include "xnix_runtime_core_backend_capability_matrix.inc"
#include "xnix_runtime_core_backend_selection_plan.inc"
#include "xnix_runtime_core_backend_lifecycle.inc"
#include "xnix_runtime_core_backend_environment_plan.inc"
#include "xnix_runtime_core_compatibility_repair_plan.inc"
#include "xnix_runtime_core_ai_diagnostic_input.inc"
#include "xnix_runtime_core_ai_diagnostic_recommendation.inc"
#include "xnix_runtime_core_ai_repair_approval_gate.inc"
