static const XnixRuntimeBackendCapability backend_capabilities[] = {
  {
    .id = "application-launch",
    .label = "Application launch",
    .status = "pending",
    .summary = "Runtime launch binding is still gated.",
    .user_visible = true,
    .requires_portal_review = false,
    .requires_state_root = true,
    .requires_snapshot = false,
    .activation_enabled = false,
    .request_object_created = false,
    .backend_process_started = false,
    .backend_details_exposed = false,
  },
  {
    .id = "package-management",
    .label = "Managed packages",
    .status = "pending",
    .summary = "Runtime package sources are planned before backend availability.",
    .user_visible = true,
    .requires_portal_review = false,
    .requires_state_root = false,
    .requires_snapshot = false,
    .activation_enabled = false,
    .request_object_created = false,
    .backend_process_started = false,
    .backend_details_exposed = false,
  },
  {
    .id = "file-bridge",
    .label = "File bridge",
    .status = "pending",
    .summary = "File access must pass Portal policy and state-root preflight.",
    .user_visible = true,
    .requires_portal_review = true,
    .requires_state_root = true,
    .requires_snapshot = false,
    .activation_enabled = false,
    .request_object_created = false,
    .backend_process_started = false,
    .backend_details_exposed = false,
  },
  {
    .id = "clipboard-bridge",
    .label = "Clipboard bridge",
    .status = "pending",
    .summary = "Clipboard access must pass Portal policy review.",
    .user_visible = true,
    .requires_portal_review = true,
    .requires_state_root = false,
    .requires_snapshot = false,
    .activation_enabled = false,
    .request_object_created = false,
    .backend_process_started = false,
    .backend_details_exposed = false,
  },
  {
    .id = "print-bridge",
    .label = "Print bridge",
    .status = "pending",
    .summary = "Print access must pass Portal policy review.",
    .user_visible = true,
    .requires_portal_review = true,
    .requires_state_root = false,
    .requires_snapshot = false,
    .activation_enabled = false,
    .request_object_created = false,
    .backend_process_started = false,
    .backend_details_exposed = false,
  },
  {
    .id = "snapshot-restore",
    .label = "Snapshot and restore",
    .status = "pending",
    .summary = "Restore points must exist before risky compatibility changes.",
    .user_visible = true,
    .requires_portal_review = false,
    .requires_state_root = true,
    .requires_snapshot = true,
    .activation_enabled = false,
    .request_object_created = false,
    .backend_process_started = false,
    .backend_details_exposed = false,
  },
  {
    .id = "diagnostics",
    .label = "Diagnostics",
    .status = "ready",
    .summary = "Runtime can describe diagnostics without starting a backend.",
    .user_visible = true,
    .requires_portal_review = false,
    .requires_state_root = false,
    .requires_snapshot = false,
    .activation_enabled = false,
    .request_object_created = false,
    .backend_process_started = false,
    .backend_details_exposed = false,
  },
};

static void
populate_backend_capability_profile(
  XnixRuntimeBackendCapabilityProfile *profile,
  const char *id,
  const char *label,
  const char *kind
)
{
  memset(profile, 0, sizeof(*profile));
  profile->id = id;
  profile->label = label;
  profile->kind = kind;
  profile->selection_state = "planned";
  profile->capability_count = sizeof(backend_capabilities) / sizeof(backend_capabilities[0]);
  for (size_t index = 0; index < profile->capability_count; index++) {
    profile->capabilities[index] = backend_capabilities[index];
    if (strcmp(profile->capabilities[index].status, "ready") == 0) {
      profile->ready_count++;
    } else if (strcmp(profile->capabilities[index].status, "pending") == 0) {
      profile->pending_count++;
    } else if (strcmp(profile->capabilities[index].status, "blocked") == 0) {
      profile->blocked_count++;
    }
  }
  profile->runtime_owned = true;
  profile->kde_policy_owner = false;
  profile->profile_ready = false;
  profile->selection_enabled = false;
  profile->backend_process_started = false;
  profile->network_required_for_planning = false;
  profile->host_root_modified = false;
  profile->privileged_container_required = false;
  profile->backend_details_exposed = false;
  profile->summary = "Compatibility backend profile is planned but not selectable until Runtime preflight passes.";
}

bool
xnix_runtime_backend_capability_matrix(XnixRuntimeBackendCapabilityMatrix *matrix)
{
  if (!matrix) {
    return false;
  }

  memset(matrix, 0, sizeof(*matrix));
  matrix->matrix_type = "compatibility-backend-capability-matrix";
  matrix->runtime_method = "GetBackendCapabilityMatrix";
  matrix->profile_count = 2;
  populate_backend_capability_profile(&matrix->profiles[0], "local-compatibility", "Local compatibility profile", "local");
  populate_backend_capability_profile(&matrix->profiles[1], "isolated-compatibility", "Isolated compatibility profile", "isolated");
  matrix->capability_count = sizeof(backend_capabilities) / sizeof(backend_capabilities[0]);
  for (size_t index = 0; index < matrix->profile_count; index++) {
    matrix->ready_capability_count += matrix->profiles[index].ready_count;
    matrix->pending_capability_count += matrix->profiles[index].pending_count;
    matrix->blocked_capability_count += matrix->profiles[index].blocked_count;
  }
  matrix->runtime_owned = true;
  matrix->c_runtime_backed = true;
  matrix->kde_policy_owner = false;
  matrix->selection_enabled = false;
  matrix->backend_launch_enabled = false;
  matrix->capability_activation_enabled = false;
  matrix->request_objects_created = false;
  matrix->state_root_created = false;
  matrix->snapshots_created = false;
  matrix->host_root_modified = false;
  matrix->privileged_container_required = false;
  matrix->backend_details_exposed = false;
  matrix->summary = "Runtime backend capability planning is visible to KDE, while backend selection and launch remain disabled.";

  return true;
}
