static void
print_file_association_application(const XnixRuntimeFileAssociationPlan *plan)
{
  fputs("{", stdout);
  print_string_field("id", plan->application->id);
  fputs(",", stdout);
  print_string_field("name", plan->application->name);
  fputs(",", stdout);
  print_string_field("desktop_file", plan->desktop_file);
  fputs(",", stdout);
  fputs("\"supported_extensions\":", stdout);
  print_string_array(plan->supported_extensions, plan->supported_extension_count);
  fputs(",", stdout);
  fputs("\"mime_types\":", stdout);
  print_string_array(plan->mime_types, plan->mime_type_count);
  fputs("}", stdout);
}

static void
print_file_association_mimeapps(const XnixRuntimeFileAssociationPlan *plan)
{
  fputs("{", stdout);
  print_string_field("path", plan->mimeapps_path);
  fputs(",", stdout);
  print_string_field("default_section", plan->default_section);
  fputs(",", stdout);
  print_string_field("added_section", plan->added_section);
  fputs("}", stdout);
}

static void
print_file_association_file_open(const XnixRuntimeFileAssociationPlan *plan)
{
  fputs("{", stdout);
  fputs("\"argv\":[", stdout);
  print_string(plan->file_open_command);
  fputs(",", stdout);
  print_string(plan->file_open_argument);
  fputs("],", stdout);
  fputs("\"portal_required\":", stdout);
  print_bool(plan->portal_required_for_file_open);
  fputs("}", stdout);
}

static void
print_file_association_entries(const XnixRuntimeFileAssociationPlan *plan)
{
  fputs("[", stdout);
  for (size_t index = 0; index < plan->mime_type_count; index++) {
    if (index > 0) {
      fputs(",", stdout);
    }

    fputs("{", stdout);
    print_string_field("mime_type", plan->mime_types[index]);
    fputs(",", stdout);
    print_string_field("desktop_file", plan->desktop_file);
    fputs(",", stdout);
    fputs("\"default_application\":", stdout);
    print_bool(plan->default_application);
    fputs(",", stdout);
    fputs("\"file_open\":", stdout);
    print_file_association_file_open(plan);
    fputs("}", stdout);
  }
  fputs("]", stdout);
}

static void
print_file_association_safety(const XnixRuntimeFileAssociationPlan *plan)
{
  fputs("{", stdout);
  fputs("\"standard_mimeapps_list\":", stdout);
  print_bool(plan->standard_mimeapps_list);
  fputs(",", stdout);
  fputs("\"staged_root_only\":", stdout);
  print_bool(plan->staged_root_only);
  fputs(",", stdout);
  fputs("\"overwrite_existing_mimeapps\":", stdout);
  print_bool(plan->overwrite_existing_mimeapps);
  fputs(",", stdout);
  fputs("\"portal_required_for_file_open\":", stdout);
  print_bool(plan->portal_required_for_file_open);
  fputs(",", stdout);
  fputs("\"files_written\":", stdout);
  print_bool(plan->files_written);
  fputs(",", stdout);
  fputs("\"host_root_modified\":", stdout);
  print_bool(plan->host_root_modified);
  fputs(",", stdout);
  fputs("\"backend_details_exposed\":", stdout);
  print_bool(plan->backend_details_exposed);
  fputs("}", stdout);
}

static void
print_file_association_plan_record(const XnixRuntimeFileAssociationPlan *plan)
{
  fputs("{", stdout);
  printf("\"version\":\"%s\",", xnix_runtime_version());
  print_string_field("plan_type", "file-association-plan");
  fputs(",", stdout);
  print_string_field("association_type", "desktop-file-association");
  fputs(",", stdout);
  print_string_field("desktop", "KDE Plasma");
  fputs(",", stdout);
  fputs("\"runtime_owned\":", stdout);
  print_bool(plan->runtime_owned);
  fputs(",", stdout);
  fputs("\"kde_policy_owner\":", stdout);
  print_bool(plan->kde_policy_owner);
  fputs(",", stdout);
  fputs("\"application\":", stdout);
  print_file_association_application(plan);
  fputs(",", stdout);
  fputs("\"mimeapps\":", stdout);
  print_file_association_mimeapps(plan);
  fputs(",", stdout);
  fputs("\"associations\":", stdout);
  print_file_association_entries(plan);
  fputs(",", stdout);
  fputs("\"safety\":", stdout);
  print_file_association_safety(plan);
  fputs(",", stdout);
  print_string_field("desktop_safe_summary", plan->summary);
  fputs("}", stdout);
}

static int
print_file_association_plan_for_application(const char *application_id)
{
  XnixRuntimeFileAssociationPlan plan;

  if (!xnix_runtime_file_association_plan(application_id, &plan)) {
    fprintf(stderr, "xnix-runtime-core: unknown application or file association plan catalog\n");
    return 64;
  }

  print_file_association_plan_record(&plan);
  fputs("\n", stdout);
  return 0;
}
