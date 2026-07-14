static const XnixRuntimeDesktopResourceBridge desktop_resource_bridges[] = {
  {
    .id = "file-open",
    .name = "File Open",
    .operation = "file-open",
    .portal_interface = "org.freedesktop.portal.FileChooser",
    .runtime_method = "GetPortalRequestPlan",
    .state = "planned",
    .portal_required = true,
    .user_approval_required = true,
    .bridge_enabled = false,
    .request_created = false,
    .direct_backend_access_allowed = false,
    .backend_details_exposed = false,
    .summary = "File opens must be mediated by XDG Desktop Portal before a compatibility backend can read selected files.",
  },
  {
    .id = "uri-open",
    .name = "URI Open",
    .operation = "uri-open",
    .portal_interface = "org.freedesktop.portal.OpenURI",
    .runtime_method = "GetPortalRequestPlan",
    .state = "planned",
    .portal_required = true,
    .user_approval_required = true,
    .bridge_enabled = false,
    .request_created = false,
    .direct_backend_access_allowed = false,
    .backend_details_exposed = false,
    .summary = "URI opens must be routed through a Runtime-owned Portal request plan.",
  },
  {
    .id = "print",
    .name = "Print",
    .operation = "print",
    .portal_interface = "org.freedesktop.portal.Print",
    .runtime_method = "GetPortalRequestPlan",
    .state = "planned",
    .portal_required = true,
    .user_approval_required = true,
    .bridge_enabled = false,
    .request_created = false,
    .direct_backend_access_allowed = false,
    .backend_details_exposed = false,
    .summary = "Printing must use a user-approved desktop Portal request before a compatibility backend can reach printers.",
  },
  {
    .id = "clipboard",
    .name = "Clipboard",
    .operation = "clipboard",
    .portal_interface = "org.freedesktop.portal.Clipboard",
    .runtime_method = "GetPortalRequestPlan",
    .state = "planned",
    .portal_required = true,
    .user_approval_required = true,
    .bridge_enabled = false,
    .request_created = false,
    .direct_backend_access_allowed = false,
    .backend_details_exposed = false,
    .summary = "Clipboard bridging stays disabled until Runtime records a user-approved Portal path.",
  },
  {
    .id = "screenshot",
    .name = "Screenshot",
    .operation = "screenshot",
    .portal_interface = "org.freedesktop.portal.Screenshot",
    .runtime_method = "GetPortalRequestPlan",
    .state = "planned",
    .portal_required = true,
    .user_approval_required = true,
    .bridge_enabled = false,
    .request_created = false,
    .direct_backend_access_allowed = false,
    .backend_details_exposed = false,
    .summary = "Screenshot access stays Portal-mediated and disabled until user review exists.",
  },
};

static const char *desktop_resource_bridge_required_runtime_gates[] = {
  "portal-policy-review",
  "portal-request-plan",
  "snapshot-baseline",
  "backend-environment-plan",
  "runtime-write-gate",
};

bool
xnix_runtime_desktop_resource_bridge_plan(
  const char *application_id,
  XnixRuntimeDesktopResourceBridgePlan *plan
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
  plan->plan_type = "desktop-resource-bridge-plan";
  plan->runtime_method = "GetDesktopResourceBridgePlan";
  plan->bridge_state = "planned";
  plan->resource_count = sizeof(desktop_resource_bridges) / sizeof(desktop_resource_bridges[0]);
  for (size_t index = 0; index < plan->resource_count; index++) {
    plan->resources[index] = desktop_resource_bridges[index];
  }
  plan->required_runtime_gate_count = sizeof(desktop_resource_bridge_required_runtime_gates) / sizeof(desktop_resource_bridge_required_runtime_gates[0]);
  for (size_t index = 0; index < plan->required_runtime_gate_count; index++) {
    plan->required_runtime_gates[index] = desktop_resource_bridge_required_runtime_gates[index];
  }
  plan->runtime_owned = true;
  plan->c_runtime_backed = true;
  plan->kde_policy_owner = false;
  plan->portal_mediated = true;
  plan->file_bridge_planned = true;
  plan->uri_bridge_planned = true;
  plan->print_bridge_planned = true;
  plan->clipboard_bridge_planned = true;
  plan->screenshot_bridge_planned = true;
  plan->bridges_enabled = false;
  plan->requests_created = false;
  plan->backend_process_started = false;
  plan->direct_host_file_access = false;
  plan->direct_clipboard_access = false;
  plan->direct_print_access = false;
  plan->host_root_modified = false;
  plan->backend_details_exposed = false;
  plan->summary = "Desktop resource bridges are planned through Runtime-owned XDG Desktop Portal requests and remain disabled until review gates pass.";

  return true;
}
