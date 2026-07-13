bool
xnix_runtime_file_association_plan(
  const char *application_id,
  XnixRuntimeFileAssociationPlan *plan
)
{
  const XnixRuntimeApplication *application = NULL;

  if (application_id == NULL || plan == NULL) {
    return false;
  }

  application = xnix_runtime_find_application(application_id);
  if (application == NULL) {
    return false;
  }

  plan->application = application;
  plan->desktop_file = "xnix-org.xnix.sample.notepad.desktop";
  plan->mimeapps_path = "usr/share/applications/mimeapps.list";
  plan->default_section = "Default Applications";
  plan->added_section = "Added Associations";
  plan->file_open_command = "xnix-compat-open";
  plan->file_open_argument = "%U";
  plan->supported_extensions[0] = ".txt";
  plan->supported_extensions[1] = ".log";
  plan->supported_extension_count = 2;
  plan->mime_types[0] = "application/x-xnix-txt";
  plan->mime_types[1] = "application/x-xnix-log";
  plan->mime_type_count = 2;
  plan->runtime_owned = true;
  plan->kde_policy_owner = false;
  plan->standard_mimeapps_list = true;
  plan->staged_root_only = true;
  plan->overwrite_existing_mimeapps = false;
  plan->portal_required_for_file_open = true;
  plan->default_application = true;
  plan->files_written = false;
  plan->host_root_modified = false;
  plan->backend_details_exposed = false;
  plan->summary = "file-association-plan maps recipe MIME types to a generated desktop entry and portal-mediated file opens.";

  return true;
}
