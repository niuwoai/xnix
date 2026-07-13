static void
fill_desktop_activation_artifact(
  XnixRuntimeDesktopActivationArtifact *artifact,
  const char *entry_point,
  const char *kind,
  const char *relative_path,
  const char *desktop_file,
  const char *command,
  bool user_visible,
  bool portal_required,
  bool activation_ready,
  const char *summary
)
{
  artifact->entry_point = entry_point;
  artifact->kind = kind;
  artifact->relative_path = relative_path;
  artifact->desktop_file = desktop_file;
  artifact->command = command;
  artifact->user_visible = user_visible;
  artifact->portal_required = portal_required;
  artifact->activation_ready = activation_ready;
  artifact->runtime_owned = true;
  artifact->kde_policy_owner = false;
  artifact->host_root_modified = false;
  artifact->backend_details_exposed = false;
  artifact->summary = summary;
}

bool
xnix_runtime_desktop_activation_manifest(
  const char *application_id,
  XnixRuntimeDesktopActivationManifest *manifest
)
{
  const XnixRuntimeApplication *application = NULL;
  const char *desktop_file = "xnix-org.xnix.sample.notepad.desktop";

  if (application_id == NULL || manifest == NULL) {
    return false;
  }

  application = xnix_runtime_find_application(application_id);
  if (application == NULL) {
    return false;
  }

  manifest->application = application;
  fill_desktop_activation_artifact(
    &manifest->artifacts[0],
    "launcher",
    "desktop-entry",
    "applications/xnix-org.xnix.sample.notepad.desktop",
    desktop_file,
    application->launcher_command,
    true,
    false,
    true,
    "Expose the compatibility application as a normal KDE launcher entry."
  );
  fill_desktop_activation_artifact(
    &manifest->artifacts[1],
    "task-manager",
    "window-identity",
    "kwin/rules/xnix-org.xnix.sample.notepad.json",
    desktop_file,
    "xnix-compat-window-identity --app org.xnix.sample.notepad --name Sample Notepad",
    true,
    false,
    true,
    "Bind windows to the desktop entry so they can be pinned and restored like native applications."
  );
  fill_desktop_activation_artifact(
    &manifest->artifacts[2],
    "file-manager",
    "dolphin-service-menu",
    "servicemenus/xnix-open-with-compatibility.desktop",
    desktop_file,
    "xnix-compat-open %U",
    true,
    true,
    true,
    "Route file opening through a Portal-mediated Runtime request."
  );
  fill_desktop_activation_artifact(
    &manifest->artifacts[3],
    "system-tray",
    "tray-status-model",
    "plasma/tray/xnix-compatibility-status.json",
    desktop_file,
    "xnix-compat-tray-status",
    true,
    false,
    true,
    "Show Runtime compatibility state without exposing backend implementation details."
  );
  fill_desktop_activation_artifact(
    &manifest->artifacts[4],
    "notifications",
    "notification-request-model",
    "notifications/xnix-compatibility.json",
    desktop_file,
    "xnix-compat-notify --app org.xnix.sample.notepad",
    true,
    false,
    true,
    "Surface install, repair, mode change, and approval events through KDE notifications."
  );
  fill_desktop_activation_artifact(
    &manifest->artifacts[5],
    "compatibility-center",
    "read-model",
    "plasma/plasmoids/org.xnix.compatibilitycenter",
    desktop_file,
    "xnix-kde-center-model",
    true,
    false,
    true,
    "Expose compatibility status and action cards through the KDE Compatibility Center."
  );
  fill_desktop_activation_artifact(
    &manifest->artifacts[6],
    "settings",
    "settings-model",
    "settings/xnix-compatibility.json",
    desktop_file,
    "xnix-compat-settings org.xnix.sample.notepad",
    true,
    true,
    true,
    "Expose user-facing settings while keeping Runtime policy ownership."
  );
  manifest->artifact_count = 7;
  manifest->runtime_owned = true;
  manifest->kde_policy_owner = false;
  manifest->official_desktop_only = true;
  manifest->stable_desktop_contract = true;
  manifest->activation_ready = true;
  manifest->files_written = false;
  manifest->host_root_modified = false;
  manifest->backend_details_exposed = false;
  manifest->desktop = "KDE Plasma";
  manifest->summary = "Desktop activation is planned as a safe KDE contract without mutating the host root.";

  return true;
}
