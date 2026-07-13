static const char *
portal_request_method_for_operation(const char *operation)
{
  if (strcmp(operation, "file-open") == 0) {
    return "OpenFile";
  }
  if (strcmp(operation, "uri-open") == 0) {
    return "OpenURI";
  }
  if (strcmp(operation, "print") == 0) {
    return "Print";
  }
  if (strcmp(operation, "screenshot") == 0) {
    return "Screenshot";
  }
  if (strcmp(operation, "clipboard") == 0) {
    return "RequestClipboard";
  }
  if (strcmp(operation, "camera") == 0) {
    return "AccessCamera";
  }
  if (strcmp(operation, "remote-desktop") == 0) {
    return "CreateSession";
  }

  return NULL;
}

static void
fill_portal_request_handle_token(
  char *target,
  size_t target_size,
  const char *application_id,
  const char *operation
)
{
  size_t offset = 0;

  if (target == NULL || target_size == 0) {
    return;
  }

  offset = (size_t)snprintf(target, target_size, "xnix_");
  if (offset >= target_size) {
    target[target_size - 1] = '\0';
    return;
  }

  for (const char *cursor = application_id; cursor != NULL && *cursor != '\0'; cursor++) {
    target[offset++] = ((*cursor >= 'a' && *cursor <= 'z') ||
                       (*cursor >= 'A' && *cursor <= 'Z') ||
                       (*cursor >= '0' && *cursor <= '9')) ? *cursor : '_';
    if (offset + 1 >= target_size) {
      target[offset] = '\0';
      return;
    }
  }

  target[offset++] = '_';
  if (offset + 1 >= target_size) {
    target[offset] = '\0';
    return;
  }

  for (const char *cursor = operation; cursor != NULL && *cursor != '\0'; cursor++) {
    target[offset++] = ((*cursor >= 'a' && *cursor <= 'z') ||
                       (*cursor >= 'A' && *cursor <= 'Z') ||
                       (*cursor >= '0' && *cursor <= '9')) ? *cursor : '_';
    if (offset + 1 >= target_size) {
      target[offset] = '\0';
      return;
    }
  }

  target[offset] = '\0';
}

bool
xnix_runtime_portal_request_plan(
  const char *application_id,
  const char *operation,
  XnixRuntimePortalRequestPlan *plan
)
{
  const XnixRuntimeApplication *application = NULL;
  const XnixRuntimePortalPolicy *policy = NULL;
  const char *portal_method = NULL;

  if (application_id == NULL || operation == NULL || plan == NULL) {
    return false;
  }

  application = xnix_runtime_find_application(application_id);
  policy = xnix_runtime_find_portal_policy(operation);
  portal_method = portal_request_method_for_operation(operation);
  if (application == NULL || policy == NULL || portal_method == NULL) {
    return false;
  }

  fill_portal_request_handle_token(
    plan->handle_token,
    sizeof(plan->handle_token),
    application_id,
    operation
  );
  plan->application = application;
  plan->policy = policy;
  plan->operation = operation;
  plan->reason = "Compatibility application requested desktop resource access.";
  plan->portal_destination = "org.freedesktop.portal.Desktop";
  plan->portal_method = portal_method;
  plan->portal_object_path = "/org/freedesktop/portal/desktop";
  plan->completion_signal = "Response";
  plan->response_field = "response";
  plan->result_owner = "Runtime";
  plan->request_allowed = strcmp(policy->decision, "deny") != 0;
  plan->request_object_required = true;
  plan->request_object_created = false;
  plan->user_mediation_required = policy->user_mediation_required;
  plan->direct_access_allowed = policy->direct_access_allowed;
  plan->portal_required = policy->portal_required;
  plan->permission_granted = false;
  plan->runtime_owned = true;
  plan->kde_policy_owner = false;
  plan->host_permission_changed = false;
  plan->host_root_modified = false;
  plan->backend_details_exposed = false;
  plan->summary = "Portal request planning is Runtime-owned and does not grant desktop resources.";

  return true;
}
