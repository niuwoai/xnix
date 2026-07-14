static const XnixRuntimeCompatibilityPermission compatibility_permission_review_permissions[] = {
  {
    .id = "documents",
    .label = "Documents",
    .operation = "file-open",
    .decision = "ask",
    .portal_interface = "org.freedesktop.portal.FileChooser",
    .portal_required = true,
    .user_mediation_required = true,
    .change_pending = false,
    .request_object_created = false,
    .permission_granted = false,
    .direct_access_allowed = false,
    .backend_details_exposed = false,
  },
  {
    .id = "downloads",
    .label = "Downloads",
    .operation = "file-open",
    .decision = "ask",
    .portal_interface = "org.freedesktop.portal.FileChooser",
    .portal_required = true,
    .user_mediation_required = true,
    .change_pending = false,
    .request_object_created = false,
    .permission_granted = false,
    .direct_access_allowed = false,
    .backend_details_exposed = false,
  },
  {
    .id = "camera",
    .label = "Camera",
    .operation = "camera",
    .decision = "deny",
    .portal_interface = "org.freedesktop.portal.Camera",
    .portal_required = true,
    .user_mediation_required = true,
    .change_pending = false,
    .request_object_created = false,
    .permission_granted = false,
    .direct_access_allowed = false,
    .backend_details_exposed = false,
  },
  {
    .id = "network",
    .label = "Network",
    .operation = "network",
    .decision = "allow",
    .portal_interface = "none",
    .portal_required = false,
    .user_mediation_required = true,
    .change_pending = false,
    .request_object_created = false,
    .permission_granted = false,
    .direct_access_allowed = false,
    .backend_details_exposed = false,
  },
  {
    .id = "clipboard",
    .label = "Clipboard",
    .operation = "clipboard",
    .decision = "ask",
    .portal_interface = "org.freedesktop.portal.Clipboard",
    .portal_required = true,
    .user_mediation_required = true,
    .change_pending = false,
    .request_object_created = false,
    .permission_granted = false,
    .direct_access_allowed = false,
    .backend_details_exposed = false,
  },
  {
    .id = "print",
    .label = "Print",
    .operation = "print",
    .decision = "ask",
    .portal_interface = "org.freedesktop.portal.Print",
    .portal_required = true,
    .user_mediation_required = true,
    .change_pending = false,
    .request_object_created = false,
    .permission_granted = false,
    .direct_access_allowed = false,
    .backend_details_exposed = false,
  },
  {
    .id = "screenshot",
    .label = "Screenshot",
    .operation = "screenshot",
    .decision = "ask",
    .portal_interface = "org.freedesktop.portal.Screenshot",
    .portal_required = true,
    .user_mediation_required = true,
    .change_pending = false,
    .request_object_created = false,
    .permission_granted = false,
    .direct_access_allowed = false,
    .backend_details_exposed = false,
  },
};

static const char *compatibility_permission_review_required_runtime_gates[] = {
  "user-review",
  "portal-policy-review",
  "runtime-write-gate",
  "settings-persistence",
  "audit-log",
};

bool
xnix_runtime_compatibility_permission_review_plan(
  const char *application_id,
  XnixRuntimeCompatibilityPermissionReviewPlan *plan
)
{
  const XnixRuntimeApplication *application;

  if (!application_id || !plan) {
    return false;
  }

  application = xnix_runtime_find_application(application_id);
  if (!application) {
    return false;
  }

  memset(plan, 0, sizeof(*plan));
  plan->application = application;
  plan->plan_type = "compatibility-permission-review-plan";
  plan->runtime_method = "GetCompatibilityPermissionReviewPlan";
  plan->review_state = "planned";
  plan->permission_count = sizeof(compatibility_permission_review_permissions) / sizeof(compatibility_permission_review_permissions[0]);
  for (size_t index = 0; index < plan->permission_count; index++) {
    plan->permissions[index] = compatibility_permission_review_permissions[index];
    if (strcmp(plan->permissions[index].decision, "allow") == 0) {
      plan->allow_count++;
    } else if (strcmp(plan->permissions[index].decision, "ask") == 0) {
      plan->ask_count++;
    } else if (strcmp(plan->permissions[index].decision, "deny") == 0) {
      plan->deny_count++;
    }
  }
  plan->required_runtime_gate_count = sizeof(compatibility_permission_review_required_runtime_gates) / sizeof(compatibility_permission_review_required_runtime_gates[0]);
  for (size_t index = 0; index < plan->required_runtime_gate_count; index++) {
    plan->required_runtime_gates[index] = compatibility_permission_review_required_runtime_gates[index];
  }
  plan->runtime_owned = true;
  plan->c_runtime_backed = true;
  plan->kde_policy_owner = false;
  plan->user_review_required = true;
  plan->portal_review_required = true;
  plan->permission_changes_applied = false;
  plan->request_objects_created = false;
  plan->permissions_granted = false;
  plan->settings_persisted = false;
  plan->host_permission_changed = false;
  plan->host_root_modified = false;
  plan->backend_details_exposed = false;
  plan->summary = "Compatibility permissions are grouped for KDE review and remain unchanged until Runtime gates pass.";

  return true;
}
