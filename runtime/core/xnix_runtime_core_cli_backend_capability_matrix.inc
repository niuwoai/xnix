static void
print_backend_capability_records(const XnixRuntimeBackendCapabilityProfile *profile)
{
  fputc('[', stdout);
  for (size_t index = 0; index < profile->capability_count; index++) {
    const XnixRuntimeBackendCapability *capability = &profile->capabilities[index];

    if (index > 0) {
      fputs(",", stdout);
    }

    fputs("{", stdout);
    print_string_field("id", capability->id);
    fputs(",", stdout);
    print_string_field("label", capability->label);
    fputs(",", stdout);
    print_string_field("status", capability->status);
    fputs(",", stdout);
    fputs("\"user_visible\":", stdout);
    print_bool(capability->user_visible);
    fputs(",", stdout);
    fputs("\"requires_portal_review\":", stdout);
    print_bool(capability->requires_portal_review);
    fputs(",", stdout);
    fputs("\"requires_state_root\":", stdout);
    print_bool(capability->requires_state_root);
    fputs(",", stdout);
    fputs("\"requires_snapshot\":", stdout);
    print_bool(capability->requires_snapshot);
    fputs(",", stdout);
    fputs("\"activation_enabled\":", stdout);
    print_bool(capability->activation_enabled);
    fputs(",", stdout);
    fputs("\"request_object_created\":", stdout);
    print_bool(capability->request_object_created);
    fputs(",", stdout);
    fputs("\"backend_process_started\":", stdout);
    print_bool(capability->backend_process_started);
    fputs(",", stdout);
    fputs("\"backend_details_exposed\":", stdout);
    print_bool(capability->backend_details_exposed);
    fputs(",", stdout);
    print_string_field("summary", capability->summary);
    fputs("}", stdout);
  }
  fputc(']', stdout);
}

static void
print_backend_capability_profiles(const XnixRuntimeBackendCapabilityMatrix *matrix)
{
  fputc('[', stdout);
  for (size_t index = 0; index < matrix->profile_count; index++) {
    const XnixRuntimeBackendCapabilityProfile *profile = &matrix->profiles[index];

    if (index > 0) {
      fputs(",", stdout);
    }

    fputs("{", stdout);
    print_string_field("id", profile->id);
    fputs(",", stdout);
    print_string_field("label", profile->label);
    fputs(",", stdout);
    print_string_field("kind", profile->kind);
    fputs(",", stdout);
    print_string_field("selection_state", profile->selection_state);
    fputs(",", stdout);
    fputs("\"capabilities\":", stdout);
    print_backend_capability_records(profile);
    fputs(",", stdout);
    fputs("\"capability_count\":", stdout);
    printf("%zu,", profile->capability_count);
    fputs("\"ready_count\":", stdout);
    printf("%zu,", profile->ready_count);
    fputs("\"pending_count\":", stdout);
    printf("%zu,", profile->pending_count);
    fputs("\"blocked_count\":", stdout);
    printf("%zu,", profile->blocked_count);
    fputs("\"runtime_owned\":", stdout);
    print_bool(profile->runtime_owned);
    fputs(",", stdout);
    fputs("\"kde_policy_owner\":", stdout);
    print_bool(profile->kde_policy_owner);
    fputs(",", stdout);
    fputs("\"profile_ready\":", stdout);
    print_bool(profile->profile_ready);
    fputs(",", stdout);
    fputs("\"selection_enabled\":", stdout);
    print_bool(profile->selection_enabled);
    fputs(",", stdout);
    fputs("\"backend_process_started\":", stdout);
    print_bool(profile->backend_process_started);
    fputs(",", stdout);
    fputs("\"network_required_for_planning\":", stdout);
    print_bool(profile->network_required_for_planning);
    fputs(",", stdout);
    fputs("\"host_root_modified\":", stdout);
    print_bool(profile->host_root_modified);
    fputs(",", stdout);
    fputs("\"privileged_container_required\":", stdout);
    print_bool(profile->privileged_container_required);
    fputs(",", stdout);
    fputs("\"backend_details_exposed\":", stdout);
    print_bool(profile->backend_details_exposed);
    fputs(",", stdout);
    print_string_field("summary", profile->summary);
    fputs("}", stdout);
  }
  fputc(']', stdout);
}

static int
print_backend_capability_matrix(void)
{
  XnixRuntimeBackendCapabilityMatrix matrix;

  if (!xnix_runtime_backend_capability_matrix(&matrix)) {
    fprintf(stderr, "xnix-runtime-core: backend capability matrix unavailable\n");
    return 64;
  }

  fputs("{", stdout);
  printf("\"version\":\"%s\",", xnix_runtime_version());
  print_string_field("matrix_type", matrix.matrix_type);
  fputs(",", stdout);
  print_string_field("runtime_method", "GetBackendCapabilityMatrix");
  fputs(",", stdout);
  fputs("\"profiles\":", stdout);
  print_backend_capability_profiles(&matrix);
  fputs(",", stdout);
  fputs("\"profile_count\":", stdout);
  printf("%zu,", matrix.profile_count);
  fputs("\"capability_count\":", stdout);
  printf("%zu,", matrix.capability_count);
  fputs("\"ready_capability_count\":", stdout);
  printf("%zu,", matrix.ready_capability_count);
  fputs("\"pending_capability_count\":", stdout);
  printf("%zu,", matrix.pending_capability_count);
  fputs("\"blocked_capability_count\":", stdout);
  printf("%zu,", matrix.blocked_capability_count);
  fputs("\"runtime_owned\":", stdout);
  print_bool(matrix.runtime_owned);
  fputs(",", stdout);
  fputs("\"c_runtime_backed\":", stdout);
  print_bool(matrix.c_runtime_backed);
  fputs(",", stdout);
  fputs("\"kde_policy_owner\":", stdout);
  print_bool(matrix.kde_policy_owner);
  fputs(",", stdout);
  fputs("\"selection_enabled\":", stdout);
  print_bool(matrix.selection_enabled);
  fputs(",", stdout);
  fputs("\"backend_launch_enabled\":", stdout);
  print_bool(matrix.backend_launch_enabled);
  fputs(",", stdout);
  fputs("\"capability_activation_enabled\":", stdout);
  print_bool(matrix.capability_activation_enabled);
  fputs(",", stdout);
  fputs("\"request_objects_created\":", stdout);
  print_bool(matrix.request_objects_created);
  fputs(",", stdout);
  fputs("\"state_root_created\":", stdout);
  print_bool(matrix.state_root_created);
  fputs(",", stdout);
  fputs("\"snapshots_created\":", stdout);
  print_bool(matrix.snapshots_created);
  fputs(",", stdout);
  fputs("\"host_root_modified\":", stdout);
  print_bool(matrix.host_root_modified);
  fputs(",", stdout);
  fputs("\"privileged_container_required\":", stdout);
  print_bool(matrix.privileged_container_required);
  fputs(",", stdout);
  fputs("\"backend_details_exposed\":", stdout);
  print_bool(matrix.backend_details_exposed);
  fputs(",", stdout);
  print_string_field("desktop_safe_summary", matrix.summary);
  fputs("}\n", stdout);
  return 0;
}
