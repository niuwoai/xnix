static void
print_desktop_activation_application(const XnixRuntimeApplication *application)
{
  fputs("{", stdout);
  print_string_field("id", application->id);
  fputs(",", stdout);
  print_string_field("name", application->name);
  fputs(",", stdout);
  print_string_field("icon", application->icon);
  fputs(",", stdout);
  print_string_field("desktop_category", application->desktop_category);
  fputs(",", stdout);
  fputs("\"supported_extensions\":[", stdout);
  print_string(application->primary_extension);
  fputs("],", stdout);
  fputs("\"mime_types\":[", stdout);
  print_string(application->primary_mime_type);
  fputs("]", stdout);
  fputs("}", stdout);
}

static void
print_desktop_activation_artifact(const XnixRuntimeDesktopActivationArtifact *artifact)
{
  fputs("{", stdout);
  print_string_field("entry_point", artifact->entry_point);
  fputs(",", stdout);
  print_string_field("kind", artifact->kind);
  fputs(",", stdout);
  print_string_field("relative_path", artifact->relative_path);
  fputs(",", stdout);
  print_string_field("desktop_file", artifact->desktop_file);
  fputs(",", stdout);
  print_string_field("command", artifact->command);
  fputs(",", stdout);
  fputs("\"user_visible\":", stdout);
  print_bool(artifact->user_visible);
  fputs(",", stdout);
  fputs("\"portal_required\":", stdout);
  print_bool(artifact->portal_required);
  fputs(",", stdout);
  fputs("\"activation_ready\":", stdout);
  print_bool(artifact->activation_ready);
  fputs(",", stdout);
  fputs("\"runtime_owned\":", stdout);
  print_bool(artifact->runtime_owned);
  fputs(",", stdout);
  fputs("\"kde_policy_owner\":", stdout);
  print_bool(artifact->kde_policy_owner);
  fputs(",", stdout);
  fputs("\"host_root_modified\":", stdout);
  print_bool(artifact->host_root_modified);
  fputs(",", stdout);
  fputs("\"backend_details_exposed\":", stdout);
  print_bool(artifact->backend_details_exposed);
  fputs(",", stdout);
  print_string_field("summary", artifact->summary);
  fputs("}", stdout);
}

static void
print_desktop_activation_artifacts(const XnixRuntimeDesktopActivationManifest *manifest)
{
  fputc('[', stdout);
  for (size_t index = 0; index < manifest->artifact_count; index++) {
    if (index > 0) {
      fputs(",", stdout);
    }
    print_desktop_activation_artifact(&manifest->artifacts[index]);
  }
  fputc(']', stdout);
}

static void
print_desktop_activation_safety(const XnixRuntimeDesktopActivationManifest *manifest)
{
  fputs("{", stdout);
  fputs("\"official_desktop_only\":", stdout);
  print_bool(manifest->official_desktop_only);
  fputs(",", stdout);
  fputs("\"stable_desktop_contract\":", stdout);
  print_bool(manifest->stable_desktop_contract);
  fputs(",", stdout);
  fputs("\"files_written\":", stdout);
  print_bool(manifest->files_written);
  fputs(",", stdout);
  fputs("\"host_root_modified\":", stdout);
  print_bool(manifest->host_root_modified);
  fputs(",", stdout);
  fputs("\"backend_details_exposed\":", stdout);
  print_bool(manifest->backend_details_exposed);
  fputs("}", stdout);
}

static void
print_desktop_activation_manifest_record(const XnixRuntimeDesktopActivationManifest *manifest)
{
  fputs("{", stdout);
  printf("\"version\":\"%s\",", xnix_runtime_version());
  print_string_field("manifest_type", "desktop-activation");
  fputs(",", stdout);
  print_string_field("desktop", manifest->desktop);
  fputs(",", stdout);
  fputs("\"runtime_owned\":", stdout);
  print_bool(manifest->runtime_owned);
  fputs(",", stdout);
  fputs("\"kde_policy_owner\":", stdout);
  print_bool(manifest->kde_policy_owner);
  fputs(",", stdout);
  fputs("\"application\":", stdout);
  print_desktop_activation_application(manifest->application);
  fputs(",", stdout);
  fputs("\"artifacts\":", stdout);
  print_desktop_activation_artifacts(manifest);
  fputs(",", stdout);
  fputs("\"artifact_count\":", stdout);
  printf("%zu", manifest->artifact_count);
  fputs(",", stdout);
  fputs("\"activation_ready\":", stdout);
  print_bool(manifest->activation_ready);
  fputs(",", stdout);
  fputs("\"safety\":", stdout);
  print_desktop_activation_safety(manifest);
  fputs(",", stdout);
  print_string_field("desktop_safe_summary", manifest->summary);
  fputs("}", stdout);
}

static int
print_desktop_activation_manifest_for_application(const char *application_id)
{
  XnixRuntimeDesktopActivationManifest manifest;

  if (!xnix_runtime_desktop_activation_manifest(application_id, &manifest)) {
    fprintf(stderr, "xnix-runtime-core: unknown application or desktop activation manifest catalog\n");
    return 64;
  }

  print_desktop_activation_manifest_record(&manifest);
  fputs("\n", stdout);
  return 0;
}
