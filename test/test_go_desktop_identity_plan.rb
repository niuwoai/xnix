#!/usr/bin/env ruby
# frozen_string_literal: true

require "json"
require "open3"
require "rbconfig"
require_relative "../lib/xnix/compatibility/application_recipe"

def assert(condition, message)
  return if condition

  warn "FAIL: #{message}"
  exit 1
end

project_root = File.expand_path("..", __dir__)
go_binary = ENV.fetch("XNIX_GO_BIN", "go")
runtime_go_binary = ENV.fetch("XNIX_RUNTIME_GO_BIN", "")

def capture_runtime_go(project_root, runtime_go_binary, go_binary, *arguments)
  if !runtime_go_binary.empty?
    Open3.capture2(runtime_go_binary, *arguments, chdir: project_root)
  else
    Open3.capture2(go_binary, "run", "./cmd/xnix-runtime-go", *arguments, chdir: project_root)
  end
end

go_available = if !runtime_go_binary.empty?
                 system(
                   runtime_go_binary,
                   "desktop-identity-plan",
                   "--registry",
                   "runtime/recipes/registry.json",
                   "--app",
                   "org.xnix.sample.notepad",
                   out: File::NULL,
                   err: File::NULL,
                   chdir: project_root
                 )
               else
                 begin
                   system(go_binary, "version", out: File::NULL, err: File::NULL)
                 rescue SystemCallError
                   false
                 end
               end

source = File.read(File.join(project_root, "internal/runtime/appidentity/identity.go"))
dockerfile = File.read(File.join(project_root, "Dockerfile"))

assert(source.include?("package appidentity"), "Go application identity package must exist")
assert(source.include?("BackendTerminologyHidden"), "Go plan must explicitly hide backend terminology")
assert(source.include?("ValidateSafeForDesktop"), "Go plan must include desktop safety validation")
assert(File.read(File.join(project_root, "internal/runtime/appidentity/registry.go")).include?("LoadRecipeFromRegistry"), "Go Runtime must load recipes from the registry")
assert(File.read(File.join(project_root, "cmd/xnix-runtime-go/main.go")).include?("--registry"), "Go Runtime CLI must support registry-backed recipe lookup")
assert(File.read(File.join(project_root, "cmd/xnix-runtime-go/main.go")).include?("backend-selection-preview"), "Go Runtime CLI must render KDE backend selection previews")
assert(File.read(File.join(project_root, "cmd/xnix-runtime-go/main.go")).include?("desktop-entry-preview"), "Go Runtime CLI must render desktop entry previews")
assert(File.read(File.join(project_root, "cmd/xnix-runtime-go/main.go")).include?("desktop-resource-bridge-preview"), "Go Runtime CLI must render KDE desktop resource bridge previews")
assert(File.read(File.join(project_root, "cmd/xnix-runtime-go/main.go")).include?("file-open-preview"), "Go Runtime CLI must render Dolphin file-open previews")
assert(File.read(File.join(project_root, "cmd/xnix-runtime-go/main.go")).include?("krunner-query-preview"), "Go Runtime CLI must render KDE KRunner query previews")
assert(File.read(File.join(project_root, "cmd/xnix-runtime-go/main.go")).include?("mimeapps-preview"), "Go Runtime CLI must render MIME association previews")
assert(File.read(File.join(project_root, "cmd/xnix-runtime-go/main.go")).include?("mode-switch-preview"), "Go Runtime CLI must render KDE mode switch previews")
assert(File.read(File.join(project_root, "cmd/xnix-runtime-go/main.go")).include?("notification-preview"), "Go Runtime CLI must render KDE notification previews")
assert(File.read(File.join(project_root, "cmd/xnix-runtime-go/main.go")).include?("permission-review-preview"), "Go Runtime CLI must render KDE permission review previews")
assert(File.read(File.join(project_root, "cmd/xnix-runtime-go/main.go")).include?("portal-request-preview"), "Go Runtime CLI must render KDE Portal request previews")
assert(File.read(File.join(project_root, "cmd/xnix-runtime-go/main.go")).include?("review-flow-preview"), "Go Runtime CLI must render KDE review flow previews")
assert(File.read(File.join(project_root, "cmd/xnix-runtime-go/main.go")).include?("settings-change-preview"), "Go Runtime CLI must render KDE settings change previews")
assert(File.read(File.join(project_root, "cmd/xnix-runtime-go/main.go")).include?("settings-preview"), "Go Runtime CLI must render KDE settings previews")
assert(File.read(File.join(project_root, "cmd/xnix-runtime-go/main.go")).include?("tray-status-preview"), "Go Runtime CLI must render KDE tray status previews")
assert(File.read(File.join(project_root, "cmd/xnix-runtime-go/main.go")).include?("window-identity-preview"), "Go Runtime CLI must render KDE window identity previews")
assert(dockerfile.include?("golang-go"), "Docker image must install Go for Runtime core validation")
assert(dockerfile.include?("go test ./..."), "Docker image must run Go tests")
assert(dockerfile.include?("go build -o /usr/local/bin/xnix-runtime-go ./cmd/xnix-runtime-go"), "Docker image must build the Go Runtime CLI")
assert(File.read(File.join(project_root, "cmd/xnix-runtime-go/main.go")).include?("compatibility-center-preview"), "Go Runtime CLI must render Compatibility Center previews")

if go_available
  output, status = capture_runtime_go(
    project_root,
    runtime_go_binary,
    go_binary,
    "desktop-identity-plan",
    "--registry",
    "runtime/recipes/registry.json",
    "--app",
    "org.xnix.sample.notepad"
  )
  assert(status.success?, "Go desktop identity CLI must run successfully")

  plan = JSON.parse(output)
  recipe = Xnix::Compatibility::ApplicationRecipe.from_hash(JSON.parse(File.read(File.join(project_root, "runtime/recipes/org.xnix.sample.notepad.json"))))

  assert(plan.fetch("schema_version") == "xnix.runtime.desktop_identity.v1", "plan schema version must be stable")
  assert(plan.fetch("application_id") == recipe.id, "plan must preserve the recipe application id")
  assert(plan.fetch("display_name") == recipe.name, "plan must preserve the user-facing application name")
  assert(plan.fetch("desktop_file") == "xnix-#{recipe.id}.desktop", "plan must expose a standard desktop file name")
  assert(plan.fetch("launch_command") == ["xnix-compat-launch", "--app", recipe.id, "%U"], "plan must use the managed Runtime launcher")
  assert(plan.fetch("mime_types") == recipe.mime_types.sort, "plan must expose normalized MIME types")
  assert(plan.fetch("runtime_owned") == true, "plan must be Runtime-owned")
  assert(plan.fetch("backend_terminology_hidden") == true, "plan must hide backend terminology")
  assert(plan.fetch("desktop_file_write_enabled") == false, "plan must not write desktop files")
  assert(plan.fetch("backend_launch_enabled") == false, "plan must not launch a backend")
  assert(plan.fetch("backend_details_exposed") == false, "plan must not expose backend details")
  assert(plan.fetch("recipe_source") == "registry", "plan must prefer registry-backed recipe lookup")
  assert(plan.fetch("registry_name") == "xnix-local-development", "plan must expose the registry name")
  assert(plan.fetch("recipe_digest_verified") == true, "plan must verify recipe digest")
  assert(plan.fetch("recipe_signature_status") == "development-only", "plan must expose development signature status")

  lower = output.downcase
  %w[prefix .exe qemu-system].each do |term|
    assert(!lower.include?(term), "plan JSON must not expose #{term}")
  end

  desktop_entry, entry_status = capture_runtime_go(
    project_root,
    runtime_go_binary,
    go_binary,
    "desktop-entry-preview",
    "--registry",
    "runtime/recipes/registry.json",
    "--app",
    "org.xnix.sample.notepad"
  )
  assert(entry_status.success?, "Go desktop entry preview CLI must run successfully")
  assert(desktop_entry.include?("[Desktop Entry]\n"), "desktop entry preview must use the desktop entry header")
  assert(desktop_entry.include?("Exec=xnix-compat-launch --app #{recipe.id} %U\n"), "desktop entry preview must use the managed Runtime launcher")
  assert(desktop_entry.include?("X-Xnix-ApplicationId=#{recipe.id}\n"), "desktop entry preview must include the Runtime application id")
  assert(!desktop_entry.downcase.include?("prefix"), "desktop entry preview must not expose implementation storage")
  assert(!desktop_entry.include?(".exe"), "desktop entry preview must not expose a Windows executable")

  center, center_status = capture_runtime_go(
    project_root,
    runtime_go_binary,
    go_binary,
    "compatibility-center-preview",
    "--registry",
    "runtime/recipes/registry.json"
  )
  assert(center_status.success?, "Go Compatibility Center preview CLI must run successfully")
  center_payload = JSON.parse(center)
  assert(center_payload.fetch("schema_version") == "xnix.runtime.compatibility_center.v1", "Compatibility Center preview schema version must be stable")
  assert(center_payload.fetch("summary_type") == "compatibility-center-preview", "Compatibility Center preview must identify its summary type")
  assert(center_payload.fetch("desktop") == "KDE Plasma", "Compatibility Center preview must target KDE Plasma")
  assert(center_payload.fetch("source").fetch("kind") == "runtime-go-registry", "Compatibility Center preview must use the Go registry source")
  assert(center_payload.fetch("source").fetch("registry_name") == "xnix-local-development", "Compatibility Center preview must expose the registry name")
  assert(center_payload.fetch("source").fetch("recipe_digest_verified") == true, "Compatibility Center preview must verify registry digests")
  assert(center_payload.fetch("application_count") == 1, "Compatibility Center preview must summarize registered applications")
  assert(center_payload.fetch("known_issue_count") == 0, "Compatibility Center preview must not invent known issues before diagnostics run")
  assert(center_payload.fetch("repair_record_count") == 0, "Compatibility Center preview must not invent repair records")
  assert(center_payload.fetch("pending_review_count") == 0, "Compatibility Center preview must not invent review requests")
  center_app = center_payload.fetch("applications").first
  assert(center_app.fetch("application_id") == recipe.id, "Compatibility Center preview must preserve application identity")
  assert(center_app.fetch("display_name") == recipe.name, "Compatibility Center preview must preserve display names")
  assert(center_app.fetch("desktop_file") == "xnix-#{recipe.id}.desktop", "Compatibility Center preview must bind generated desktop files")
  assert(center_app.fetch("compatibility_state") == "registered", "Compatibility Center preview must expose registered compatibility state")
  assert(center_app.fetch("diagnostics_state") == "not-run", "Compatibility Center preview must not claim diagnostics execution")
  assert(center_app.fetch("runtime_mode") == "Automatic", "Compatibility Center preview must expose user-facing runtime mode")
  assert(center_app.fetch("known_issue_count") == 0, "Compatibility Center preview must not invent per-app issues")
  assert(center_app.fetch("repair_record_state") == "none", "Compatibility Center preview must not invent repair records")
  assert(center_app.fetch("actions") == ["open-settings", "show-diagnostics", "review-application"], "Compatibility Center preview must expose safe navigation actions")
  assert(center_payload.fetch("runtime_owned") == true, "Compatibility Center preview must remain Runtime-owned")
  assert(center_payload.fetch("kde_policy_owner") == false, "Compatibility Center preview must not make KDE own backend policy")
  assert(center_payload.fetch("action_execution_enabled") == false, "Compatibility Center preview must not enable actions")
  assert(center_payload.fetch("repair_execution_enabled") == false, "Compatibility Center preview must not enable repairs")
  assert(center_payload.fetch("backend_launch_enabled") == false, "Compatibility Center preview must not launch backends")
  assert(center_payload.fetch("settings_persistence_enabled") == false, "Compatibility Center preview must not persist settings")
  assert(center_payload.fetch("host_root_modified") == false, "Compatibility Center preview must not mutate the host root")
  assert(center_payload.fetch("backend_details_exposed") == false, "Compatibility Center preview must not expose backend details")
  assert(!center.downcase.include?("prefix"), "Compatibility Center preview must not expose implementation storage")
  assert(!center.include?(".exe"), "Compatibility Center preview must not expose a Windows executable")
  assert(!center.downcase.include?("proton"), "Compatibility Center preview must not expose backend implementation names")

  backend_selection, backend_selection_status = capture_runtime_go(
    project_root,
    runtime_go_binary,
    go_binary,
    "backend-selection-preview",
    "--registry",
    "runtime/recipes/registry.json",
    "--app",
    "org.xnix.sample.notepad"
  )
  assert(backend_selection_status.success?, "Go backend selection preview CLI must run successfully")
  backend_selection_payload = JSON.parse(backend_selection)
  assert(backend_selection_payload.fetch("schema_version") == "xnix.runtime.backend_selection.v1", "backend selection preview schema version must be stable")
  assert(backend_selection_payload.fetch("request_type") == "backend-selection-preview", "backend selection preview must identify its request type")
  assert(backend_selection_payload.fetch("plan_type") == "compatibility-backend-selection-plan", "backend selection preview must identify the Runtime plan type")
  assert(backend_selection_payload.fetch("source") == "compatibility-center", "backend selection preview must identify the Compatibility Center source")
  assert(backend_selection_payload.fetch("desktop") == "KDE Plasma", "backend selection preview must target KDE Plasma")
  assert(backend_selection_payload.fetch("runtime_method") == "GetBackendSelectionPlan", "backend selection preview must expose the Runtime method")
  assert(backend_selection_payload.fetch("application_id") == recipe.id, "backend selection preview must preserve application identity")
  assert(backend_selection_payload.fetch("display_name") == recipe.name, "backend selection preview must preserve display names")
  assert(backend_selection_payload.fetch("desktop_file") == "xnix-#{recipe.id}.desktop", "backend selection preview must bind generated desktop files")
  assert(backend_selection_payload.fetch("recommended_profile_id") == "local-compatibility", "backend selection preview must recommend local compatibility for automatic recipes")
  assert(backend_selection_payload.fetch("candidate_count") == 2, "backend selection preview must expose two compatibility profiles")
  assert(backend_selection_payload.fetch("ready_candidate_count") == 0, "backend selection preview must keep profiles blocked before review")
  assert(backend_selection_payload.fetch("blocked_candidate_count") == 2, "backend selection preview must count blocked profiles")
  backend_profiles = backend_selection_payload.fetch("candidate_profiles")
  assert(backend_profiles.map { |profile| profile.fetch("id") } == %w[local-compatibility isolated-compatibility], "backend selection preview must preserve profile order")
  assert(backend_profiles.first.fetch("recommended") == true, "backend selection preview must mark the recommended profile")
  assert(backend_profiles.first.fetch("selection_state") == "recommended", "backend selection preview must expose the recommended profile state")
  assert(backend_profiles.all? { |profile| !profile.fetch("ready") }, "backend selection preview must not mark profiles ready before review")
  assert(backend_profiles.all? { |profile| profile.fetch("blocked") }, "backend selection preview must keep profiles blocked before review")
  assert(backend_profiles.all? { |profile| !profile.fetch("selection_committed") }, "backend selection preview must not commit profile selection")
  assert(backend_profiles.all? { |profile| !profile.fetch("environment_created") }, "backend selection preview must not create profile environments")
  assert(backend_profiles.all? { |profile| !profile.fetch("backend_process_started") }, "backend selection preview must not start backend processes")
  assert(backend_profiles.all? { |profile| !profile.fetch("backend_details_exposed") }, "backend selection preview must not expose backend details")
  assert(backend_selection_payload.fetch("required_reviews") == %w[backend-capability-review backend-binding-review application-state-root-review portal-policy-review snapshot-baseline-review], "backend selection preview must expose required Runtime reviews")
  assert(backend_selection_payload.fetch("runtime_owned") == true, "backend selection preview must remain Runtime-owned")
  assert(backend_selection_payload.fetch("go_runtime_backed") == true, "backend selection preview must be Go Runtime-backed")
  assert(backend_selection_payload.fetch("kde_policy_owner") == false, "backend selection preview must not make KDE own backend policy")
  assert(backend_selection_payload.fetch("user_visible") == true, "backend selection preview must remain user visible")
  assert(backend_selection_payload.fetch("selection_committed") == false, "backend selection preview must not commit selection")
  assert(backend_selection_payload.fetch("selection_change_enabled") == false, "backend selection preview must not enable selection changes")
  assert(backend_selection_payload.fetch("backend_launch_enabled") == false, "backend selection preview must not launch backends")
  assert(backend_selection_payload.fetch("capability_activation_enabled") == false, "backend selection preview must not activate capabilities")
  assert(backend_selection_payload.fetch("environment_created") == false, "backend selection preview must not create environments")
  assert(backend_selection_payload.fetch("request_object_created") == false, "backend selection preview must not create request objects")
  assert(backend_selection_payload.fetch("state_root_created") == false, "backend selection preview must not create state roots")
  assert(backend_selection_payload.fetch("snapshot_created") == false, "backend selection preview must not create snapshots")
  assert(backend_selection_payload.fetch("host_root_modified") == false, "backend selection preview must not mutate the host root")
  assert(backend_selection_payload.fetch("privileged_container_required") == false, "backend selection preview must not require privileged containers")
  assert(backend_selection_payload.fetch("backend_details_exposed") == false, "backend selection preview must not expose backend details")
  assert(!backend_selection.downcase.include?("prefix"), "backend selection preview must not expose implementation storage")
  assert(!backend_selection.include?(".exe"), "backend selection preview must not expose a Windows executable")
  assert(!backend_selection.downcase.include?("program files"), "backend selection preview must not expose Windows paths")
  assert(!backend_selection.downcase.include?("qemu-system"), "backend selection preview must not expose VM implementation commands")
  assert(!backend_selection.downcase.include?("proton"), "backend selection preview must not expose backend implementation names")
  assert(!backend_selection.downcase.include?("wine "), "backend selection preview must not expose backend implementation names")
  assert(!backend_selection.downcase.include?("virtual machine"), "backend selection preview must not expose implementation labels")

  file_open, file_open_status = capture_runtime_go(
    project_root,
    runtime_go_binary,
    go_binary,
    "file-open-preview",
    "--registry",
    "runtime/recipes/registry.json",
    "file:///home/test/Documents/example.txt"
  )
  assert(file_open_status.success?, "Go file-open preview CLI must run successfully")
  file_open_payload = JSON.parse(file_open)
  assert(file_open_payload.fetch("schema_version") == "xnix.runtime.file_open.v1", "file-open preview schema version must be stable")
  assert(file_open_payload.fetch("request_type") == "file-open-preview", "file-open preview must identify its request type")
  assert(file_open_payload.fetch("source") == "dolphin-service-menu", "file-open preview must identify the Dolphin source")
  assert(file_open_payload.fetch("desktop") == "KDE Plasma", "file-open preview must target KDE Plasma")
  assert(file_open_payload.fetch("application_id") == recipe.id, "file-open preview must resolve applications by extension")
  assert(file_open_payload.fetch("desktop_file") == "xnix-#{recipe.id}.desktop", "file-open preview must bind generated desktop files")
  assert(file_open_payload.fetch("runtime_method") == "Launch", "file-open preview must target the Runtime launch method")
  assert(file_open_payload.fetch("portal_required") == true, "file-open preview must require Portal mediation")
  assert(file_open_payload.fetch("portal_interface") == "org.freedesktop.portal.FileChooser", "file-open preview must target the FileChooser Portal")
  assert(file_open_payload.fetch("portal_method") == "OpenFile", "file-open preview must target the OpenFile Portal method")
  assert(file_open_payload.fetch("file_count") == 1, "file-open preview must count selected files")
  assert(file_open_payload.fetch("file_uris") == ["file:///home/test/Documents/example.txt"], "file-open preview must preserve file URIs")
  assert(file_open_payload.fetch("selected_extension") == ".txt", "file-open preview must expose the selected extension")
  assert(file_open_payload.fetch("selection_mode") == "extension-match", "file-open preview must select by extension when no app is provided")
  assert(file_open_payload.fetch("action").fetch("type") == "runtime-file-open", "file-open preview must expose a Runtime file-open action")
  assert(file_open_payload.fetch("action").fetch("argv") == ["xnix-compat-open", "--app", recipe.id, "%U"], "file-open preview must delegate to the managed file-open command")
  assert(file_open_payload.fetch("runtime_owned") == true, "file-open preview must remain Runtime-owned")
  assert(file_open_payload.fetch("kde_policy_owner") == false, "file-open preview must not make KDE own backend policy")
  assert(file_open_payload.fetch("request_object_created") == false, "file-open preview must not create request objects")
  assert(file_open_payload.fetch("permission_granted") == false, "file-open preview must not grant permissions")
  assert(file_open_payload.fetch("backend_launch_enabled") == false, "file-open preview must not launch backends")
  assert(file_open_payload.fetch("direct_host_file_access") == false, "file-open preview must not directly access host files")
  assert(file_open_payload.fetch("host_root_modified") == false, "file-open preview must not mutate the host root")
  assert(file_open_payload.fetch("backend_details_exposed") == false, "file-open preview must not expose backend details")
  assert(!file_open.downcase.include?("prefix"), "file-open preview must not expose implementation storage")
  assert(!file_open.include?(".exe"), "file-open preview must not expose a Windows executable")
  assert(!file_open.downcase.include?("proton"), "file-open preview must not expose backend implementation names")

  krunner, krunner_status = capture_runtime_go(
    project_root,
    runtime_go_binary,
    go_binary,
    "krunner-query-preview",
    "--registry",
    "runtime/recipes/registry.json",
    "--query",
    "notepad"
  )
  assert(krunner_status.success?, "Go KRunner query preview CLI must run successfully")
  krunner_payload = JSON.parse(krunner)
  assert(krunner_payload.fetch("schema_version") == "xnix.runtime.krunner_query.v1", "KRunner preview schema version must be stable")
  assert(krunner_payload.fetch("query_type") == "krunner-query-plan", "KRunner preview must identify Runtime query plans")
  assert(krunner_payload.fetch("entry_point") == "krunner", "KRunner preview must identify the KDE runner entry point")
  assert(krunner_payload.fetch("desktop") == "KDE Plasma", "KRunner preview must target KDE Plasma")
  assert(krunner_payload.fetch("source").fetch("kind") == "runtime-go-registry", "KRunner preview must use the Go registry source")
  assert(krunner_payload.fetch("source").fetch("registry_name") == "xnix-local-development", "KRunner preview must expose the registry name")
  assert(krunner_payload.fetch("source").fetch("recipe_digest_verified") == true, "KRunner preview must verify registry digests")
  assert(krunner_payload.fetch("runtime_owned") == true, "KRunner preview must remain Runtime-owned")
  assert(krunner_payload.fetch("kde_policy_owner") == false, "KRunner preview must not make KDE own backend policy")
  assert(krunner_payload.fetch("matches").length == 1, "KRunner preview must resolve the sample application")
  krunner_match = krunner_payload.fetch("matches").first
  assert(krunner_match.fetch("application_id") == recipe.id, "KRunner preview must preserve application identity")
  assert(krunner_match.fetch("name") == recipe.name, "KRunner preview must preserve display names")
  assert(krunner_match.fetch("subtitle") == "Open as a normal Linux application", "KRunner preview must use normal desktop wording")
  assert(krunner_match.fetch("action").fetch("type") == "runtime-launch", "KRunner preview must expose a Runtime launch action")
  assert(krunner_match.fetch("action").fetch("desktop_entry_id") == "xnix-#{recipe.id}.desktop", "KRunner preview must bind generated desktop files")
  assert(krunner_match.fetch("action").fetch("argv") == ["xnix-compat-launch", "--app", recipe.id], "KRunner preview must delegate to the managed launcher")
  assert(krunner_payload.fetch("summary").fetch("query_execution_enabled") == false, "KRunner preview must not execute queries")
  assert(krunner_payload.fetch("summary").fetch("backend_launch_enabled") == false, "KRunner preview must not launch backends")
  assert(krunner_payload.fetch("summary").fetch("backend_details_exposed") == false, "KRunner preview must not expose backend details")
  assert(krunner_payload.fetch("host_root_modified") == false, "KRunner preview must not mutate the host root")
  assert(krunner_payload.fetch("backend_details_exposed") == false, "KRunner preview must keep backend details hidden")
  assert(!krunner.downcase.include?("prefix"), "KRunner preview must not expose implementation storage")
  assert(!krunner.include?(".exe"), "KRunner preview must not expose a Windows executable")
  assert(!krunner.downcase.include?("proton"), "KRunner preview must not expose backend implementation names")

  mimeapps, mimeapps_status = capture_runtime_go(
    project_root,
    runtime_go_binary,
    go_binary,
    "mimeapps-preview",
    "--registry",
    "runtime/recipes/registry.json",
    "--app",
    "org.xnix.sample.notepad"
  )
  assert(mimeapps_status.success?, "Go MIME apps preview CLI must run successfully")
  assert(mimeapps.include?("[Default Applications]\n"), "MIME apps preview must include default associations")
  assert(mimeapps.include?("[Added Associations]\n"), "MIME apps preview must include added associations")
  assert(mimeapps.include?("application/x-xnix-txt=xnix-#{recipe.id}.desktop\n"), "MIME apps preview must map text files to the generated desktop file")
  assert(mimeapps.include?("application/x-xnix-log=xnix-#{recipe.id}.desktop;\n"), "MIME apps preview must add log file associations")
  assert(!mimeapps.downcase.include?("prefix"), "MIME apps preview must not expose implementation storage")
  assert(!mimeapps.include?(".exe"), "MIME apps preview must not expose a Windows executable")

  notification, notification_status = capture_runtime_go(
    project_root,
    runtime_go_binary,
    go_binary,
    "notification-preview",
    "--registry",
    "runtime/recipes/registry.json",
    "--app",
    "org.xnix.sample.notepad",
    "--event",
    "install-failed"
  )
  assert(notification_status.success?, "Go notification preview CLI must run successfully")
  notification_payload = JSON.parse(notification)
  assert(notification_payload.fetch("schema_version") == "xnix.runtime.notification.v1", "notification preview schema version must be stable")
  assert(notification_payload.fetch("request_type") == "desktop-notification-preview", "notification preview must identify its request type")
  assert(notification_payload.fetch("source") == "runtime-event", "notification preview must identify Runtime events as the source")
  assert(notification_payload.fetch("desktop") == "KDE Plasma", "notification preview must target KDE Plasma")
  assert(notification_payload.fetch("application_id") == recipe.id, "notification preview must preserve application identity")
  assert(notification_payload.fetch("desktop_file") == "xnix-#{recipe.id}.desktop", "notification preview must bind generated desktop files")
  assert(notification_payload.fetch("event_type") == "install-failed", "notification preview must preserve event type")
  assert(notification_payload.fetch("urgency") == "critical", "install-failed notifications must be critical")
  assert(notification_payload.fetch("category") == "compatibility.install", "install-failed notifications must use the install category")
  assert(notification_payload.fetch("actions") == ["open-compatibility-center", "show-diagnostics"], "install-failed notifications must expose review actions")
  assert(notification_payload.fetch("requires_user_review") == true, "install-failed notifications must require review")
  assert(notification_payload.fetch("runtime_owned") == true, "notification preview must remain Runtime-owned")
  assert(notification_payload.fetch("kde_policy_owner") == false, "notification preview must not make KDE own backend policy")
  assert(notification_payload.fetch("user_visible") == true, "notification preview must be user visible")
  assert(notification_payload.fetch("action_execution_enabled") == false, "notification preview must not enable action execution")
  assert(notification_payload.fetch("repair_execution_enabled") == false, "notification preview must not enable repair execution")
  assert(notification_payload.fetch("settings_persistence_enabled") == false, "notification preview must not persist settings")
  assert(notification_payload.fetch("host_root_modified") == false, "notification preview must not mutate the host root")
  assert(notification_payload.fetch("backend_details_exposed") == false, "notification preview must not expose backend details")
  assert(!notification.downcase.include?("prefix"), "notification preview must not expose implementation storage")
  assert(!notification.include?(".exe"), "notification preview must not expose a Windows executable")

  settings, settings_status = capture_runtime_go(
    project_root,
    runtime_go_binary,
    go_binary,
    "settings-preview",
    "--registry",
    "runtime/recipes/registry.json",
    "--app",
    "org.xnix.sample.notepad"
  )
  assert(settings_status.success?, "Go settings preview CLI must run successfully")
  settings_payload = JSON.parse(settings)
  assert(settings_payload.fetch("schema_version") == "xnix.runtime.settings.v1", "settings preview schema version must be stable")
  assert(settings_payload.fetch("request_type") == "settings-preview", "settings preview must identify its request type")
  assert(settings_payload.fetch("desktop") == "KDE Plasma", "settings preview must target KDE Plasma")
  assert(settings_payload.fetch("application_id") == recipe.id, "settings preview must preserve application identity")
  assert(settings_payload.fetch("display_name") == recipe.name, "settings preview must preserve display names")
  assert(settings_payload.fetch("desktop_file") == "xnix-#{recipe.id}.desktop", "settings preview must bind generated desktop files")
  assert(settings_payload.fetch("section_count") == 5, "settings preview must expose all user-facing settings sections")
  assert(settings_payload.fetch("sections").map { |section| section.fetch("id") } == %w[run-mode resource-access devices network snapshots], "settings preview must expose the required section order")
  run_mode = settings_payload.fetch("sections").find { |section| section.fetch("id") == "run-mode" }
  assert(run_mode.fetch("fields").find { |field| field.fetch("id") == "mode" }.fetch("value") == "automatic", "settings preview must default to automatic mode")
  access = settings_payload.fetch("sections").find { |section| section.fetch("id") == "resource-access" }
  assert(access.fetch("fields").find { |field| field.fetch("id") == "documents" }.fetch("value") == "ask", "settings preview must default documents access to review")
  assert(access.fetch("fields").find { |field| field.fetch("id") == "downloads" }.fetch("value") == "ask", "settings preview must default downloads access to review")
  devices = settings_payload.fetch("sections").find { |section| section.fetch("id") == "devices" }
  assert(devices.fetch("fields").find { |field| field.fetch("id") == "camera" }.fetch("value") == "deny", "settings preview must default camera access to deny")
  assert(settings_payload.fetch("runtime_owned") == true, "settings preview must remain Runtime-owned")
  assert(settings_payload.fetch("kde_policy_owner") == false, "settings preview must not make KDE own backend policy")
  assert(settings_payload.fetch("user_visible") == true, "settings preview must be user visible")
  assert(settings_payload.fetch("settings_persisted") == false, "settings preview must not claim persisted settings")
  assert(settings_payload.fetch("settings_persistence_enabled") == false, "settings preview must keep settings persistence disabled")
  assert(settings_payload.fetch("host_root_modified") == false, "settings preview must not mutate the host root")
  assert(settings_payload.fetch("backend_details_exposed") == false, "settings preview must not expose backend details")
  assert(!settings.downcase.include?("prefix"), "settings preview must not expose implementation storage")
  assert(!settings.include?(".exe"), "settings preview must not expose a Windows executable")
  assert(!settings.downcase.include?("proton"), "settings preview must not expose backend implementation names")

  mode_switch, mode_switch_status = capture_runtime_go(
    project_root,
    runtime_go_binary,
    go_binary,
    "mode-switch-preview",
    "--registry",
    "runtime/recipes/registry.json",
    "--app",
    "org.xnix.sample.notepad",
    "--mode",
    "prefer-compatibility"
  )
  assert(mode_switch_status.success?, "Go mode switch preview CLI must run successfully")
  mode_switch_payload = JSON.parse(mode_switch)
  assert(mode_switch_payload.fetch("schema_version") == "xnix.runtime.mode_switch.v1", "mode switch preview schema version must be stable")
  assert(mode_switch_payload.fetch("request_type") == "mode-switch-preview", "mode switch preview must identify its request type")
  assert(mode_switch_payload.fetch("plan_type") == "compatibility-mode-switch-plan", "mode switch preview must identify the Runtime plan type")
  assert(mode_switch_payload.fetch("source") == "unified-settings", "mode switch preview must identify the KDE settings source")
  assert(mode_switch_payload.fetch("desktop") == "KDE Plasma", "mode switch preview must target KDE Plasma")
  assert(mode_switch_payload.fetch("runtime_method") == "GetCompatibilityModeSwitchPlan", "mode switch preview must expose the Runtime method")
  assert(mode_switch_payload.fetch("application_id") == recipe.id, "mode switch preview must preserve application identity")
  assert(mode_switch_payload.fetch("display_name") == recipe.name, "mode switch preview must preserve display names")
  assert(mode_switch_payload.fetch("desktop_file") == "xnix-#{recipe.id}.desktop", "mode switch preview must bind generated desktop files")
  assert(mode_switch_payload.fetch("current_mode") == "automatic", "mode switch preview must expose the current mode")
  assert(mode_switch_payload.fetch("requested_mode") == "prefer-compatibility", "mode switch preview must expose the requested mode")
  assert(mode_switch_payload.fetch("mode_state") == "planned", "mode switch preview must stay planned")
  assert(mode_switch_payload.fetch("mode_count") == 4, "mode switch preview must expose four user-facing modes")
  assert(mode_switch_payload.fetch("modes").map { |mode| mode.fetch("id") } == %w[automatic prefer-performance prefer-compatibility isolated-execution], "mode switch preview must preserve mode order")
  assert(mode_switch_payload.fetch("modes").one? { |mode| mode.fetch("selected") }, "mode switch preview must mark one selected mode")
  assert(mode_switch_payload.fetch("modes").one? { |mode| mode.fetch("requested") }, "mode switch preview must mark one requested mode")
  assert(mode_switch_payload.fetch("required_runtime_gates") == %w[settings-review portal-policy-review snapshot-baseline backend-environment-plan runtime-write-gate], "mode switch preview must expose required Runtime gates")
  assert(mode_switch_payload.fetch("runtime_owned") == true, "mode switch preview must remain Runtime-owned")
  assert(mode_switch_payload.fetch("go_runtime_backed") == true, "mode switch preview must be Go Runtime-backed")
  assert(mode_switch_payload.fetch("kde_policy_owner") == false, "mode switch preview must not make KDE own backend policy")
  assert(mode_switch_payload.fetch("user_visible") == true, "mode switch preview must be user visible")
  assert(mode_switch_payload.fetch("valid_mode") == true, "mode switch preview must validate requested modes")
  assert(mode_switch_payload.fetch("requires_user_confirmation") == true, "mode switch preview must require user confirmation")
  assert(mode_switch_payload.fetch("portal_review_required") == true, "mode switch preview must require Portal review")
  assert(mode_switch_payload.fetch("snapshot_required") == true, "mode switch preview must require snapshot review")
  assert(mode_switch_payload.fetch("settings_persistence_enabled") == false, "mode switch preview must keep settings persistence disabled")
  assert(mode_switch_payload.fetch("backend_reconfiguration_enabled") == false, "mode switch preview must not reconfigure backends")
  assert(mode_switch_payload.fetch("backend_process_started") == false, "mode switch preview must not start backend processes")
  assert(mode_switch_payload.fetch("launch_enabled") == false, "mode switch preview must not enable launch")
  assert(mode_switch_payload.fetch("host_root_modified") == false, "mode switch preview must not mutate the host root")
  assert(mode_switch_payload.fetch("backend_details_exposed") == false, "mode switch preview must not expose backend details")
  assert(!mode_switch.downcase.include?("prefix"), "mode switch preview must not expose implementation storage")
  assert(!mode_switch.include?(".exe"), "mode switch preview must not expose a Windows executable")
  assert(!mode_switch.downcase.include?("proton"), "mode switch preview must not expose backend implementation names")

  settings_change, settings_change_status = capture_runtime_go(
    project_root,
    runtime_go_binary,
    go_binary,
    "settings-change-preview",
    "--registry",
    "runtime/recipes/registry.json",
    "--app",
    "org.xnix.sample.notepad",
    "--section",
    "resource-access",
    "--field",
    "documents",
    "--value",
    "allow"
  )
  assert(settings_change_status.success?, "Go settings change preview CLI must run successfully")
  settings_change_payload = JSON.parse(settings_change)
  assert(settings_change_payload.fetch("schema_version") == "xnix.runtime.settings_change.v1", "settings change preview schema version must be stable")
  assert(settings_change_payload.fetch("request_type") == "settings-change-preview", "settings change preview must identify its request type")
  assert(settings_change_payload.fetch("plan_type") == "settings-change-plan", "settings change preview must identify the Runtime plan type")
  assert(settings_change_payload.fetch("source") == "unified-settings", "settings change preview must identify the KDE settings source")
  assert(settings_change_payload.fetch("desktop") == "KDE Plasma", "settings change preview must target KDE Plasma")
  assert(settings_change_payload.fetch("runtime_method") == "GetCompatibilitySettingsChangePlan", "settings change preview must expose the Runtime method")
  assert(settings_change_payload.fetch("application_id") == recipe.id, "settings change preview must preserve application identity")
  assert(settings_change_payload.fetch("display_name") == recipe.name, "settings change preview must preserve display names")
  assert(settings_change_payload.fetch("desktop_file") == "xnix-#{recipe.id}.desktop", "settings change preview must bind generated desktop files")
  assert(settings_change_payload.fetch("section_id") == "resource-access", "settings change preview must expose the requested section")
  assert(settings_change_payload.fetch("field_id") == "documents", "settings change preview must expose the requested field")
  assert(settings_change_payload.fetch("requested_value") == "allow", "settings change preview must expose the requested value")
  assert(settings_change_payload.fetch("change_state") == "planned", "settings change preview must stay planned")
  assert(settings_change_payload.fetch("user_confirmation_required") == true, "settings change preview must require user confirmation")
  assert(settings_change_payload.fetch("portal_policy_review_required") == true, "settings change preview must require Portal policy review for file access")
  assert(settings_change_payload.fetch("snapshot_recommended") == false, "settings change preview must not recommend snapshots for file access only")
  assert(settings_change_payload.fetch("runtime_restart_required") == false, "settings change preview must not require a Runtime restart")
  assert(settings_change_payload.fetch("affected_policy").fetch("options") == %w[allow ask deny], "settings change preview must expose allowed policy options")
  assert(settings_change_payload.fetch("steps").map { |step| step.fetch("id") } == %w[validate-setting review-user-confirmation review-portal-policy prepare-restore-point persist-runtime-setting], "settings change preview must expose review steps")
  assert(settings_change_payload.fetch("steps").map { |step| step.fetch("status") } == %w[pass required required pass pending], "settings change preview must gate persistence behind required reviews")
  assert(settings_change_payload.fetch("blocked_actions").include?("persist compatibility settings before Runtime confirmation"), "settings change preview must block premature persistence")
  assert(settings_change_payload.fetch("runtime_owned") == true, "settings change preview must remain Runtime-owned")
  assert(settings_change_payload.fetch("kde_policy_owner") == false, "settings change preview must not make KDE own backend policy")
  assert(settings_change_payload.fetch("user_visible") == true, "settings change preview must be user visible")
  assert(settings_change_payload.fetch("apply_enabled") == false, "settings change preview must not enable apply")
  assert(settings_change_payload.fetch("settings_persisted") == false, "settings change preview must not persist settings")
  assert(settings_change_payload.fetch("settings_persistence_enabled") == false, "settings change preview must keep settings persistence disabled")
  assert(settings_change_payload.fetch("host_root_modified") == false, "settings change preview must not mutate the host root")
  assert(settings_change_payload.fetch("backend_details_exposed") == false, "settings change preview must not expose backend details")
  assert(!settings_change.downcase.include?("prefix"), "settings change preview must not expose implementation storage")
  assert(!settings_change.include?(".exe"), "settings change preview must not expose a Windows executable")
  assert(!settings_change.downcase.include?("proton"), "settings change preview must not expose backend implementation names")

  review_flow, review_flow_status = capture_runtime_go(
    project_root,
    runtime_go_binary,
    go_binary,
    "review-flow-preview",
    "--registry",
    "runtime/recipes/registry.json",
    "--app",
    "org.xnix.sample.notepad",
    "--section",
    "resource-access",
    "--field",
    "documents",
    "--value",
    "ask",
    "--operation",
    "file-open"
  )
  assert(review_flow_status.success?, "Go review flow preview CLI must run successfully")
  review_flow_payload = JSON.parse(review_flow)
  assert(review_flow_payload.fetch("schema_version") == "xnix.runtime.review_flow.v1", "review flow preview schema version must be stable")
  assert(review_flow_payload.fetch("request_type") == "review-flow-preview", "review flow preview must identify its request type")
  assert(review_flow_payload.fetch("plan_type") == "compatibility-review-flow-plan", "review flow preview must identify the Runtime plan type")
  assert(review_flow_payload.fetch("source") == "compatibility-center-review", "review flow preview must identify the Compatibility Center source")
  assert(review_flow_payload.fetch("desktop") == "KDE Plasma", "review flow preview must target KDE Plasma")
  assert(review_flow_payload.fetch("runtime_method") == "GetCompatibilityReviewFlowPlan", "review flow preview must expose the Runtime method")
  assert(review_flow_payload.fetch("application_id") == recipe.id, "review flow preview must preserve application identity")
  assert(review_flow_payload.fetch("display_name") == recipe.name, "review flow preview must preserve display names")
  assert(review_flow_payload.fetch("desktop_file") == "xnix-#{recipe.id}.desktop", "review flow preview must bind generated desktop files")
  assert(review_flow_payload.fetch("review_state") == "planned", "review flow preview must stay planned")
  assert(review_flow_payload.fetch("section_id") == "resource-access", "review flow preview must expose the requested section")
  assert(review_flow_payload.fetch("field_id") == "documents", "review flow preview must expose the requested field")
  assert(review_flow_payload.fetch("requested_value") == "ask", "review flow preview must expose the requested value")
  assert(review_flow_payload.fetch("operation") == "file-open", "review flow preview must expose the Portal operation")
  assert(review_flow_payload.fetch("step_count") == 5, "review flow preview must expose five review steps")
  assert(review_flow_payload.fetch("required_review_count") == 3, "review flow preview must require three reviews")
  assert(review_flow_payload.fetch("blocked_step_count") == 1, "review flow preview must expose one blocked Runtime write gate")
  assert(review_flow_payload.fetch("pending_step_count") == 1, "review flow preview must expose one pending receipt step")
  assert(review_flow_payload.fetch("steps").map { |step| step.fetch("id") } == %w[settings-change-review permission-review portal-request-review runtime-write-gate review-receipt], "review flow preview must preserve review step order")
  assert(review_flow_payload.fetch("settings_change_plan").fetch("apply_enabled") == false, "review flow preview must not apply settings changes")
  assert(review_flow_payload.fetch("settings_change_plan").fetch("settings_persisted") == false, "review flow preview must not persist settings")
  assert(review_flow_payload.fetch("permission_review_plan").fetch("permissions_granted") == false, "review flow preview must not grant permissions")
  assert(review_flow_payload.fetch("portal_request_plan").fetch("request_object_created") == false, "review flow preview must not create Portal request objects")
  assert(review_flow_payload.fetch("portal_request_plan").fetch("permission_granted") == false, "review flow preview must not grant Portal permissions")
  assert(review_flow_payload.fetch("runtime_write_gate").fetch("gate_type") == "runtime-write-gate", "review flow preview must include Runtime write gate summary")
  assert(review_flow_payload.fetch("runtime_write_gate").fetch("write_method_enabled") == false, "review flow preview must keep write methods disabled")
  assert(review_flow_payload.fetch("review_receipt").fetch("receipt_type") == "compatibility-center-action-review-receipt", "review flow preview must include review receipt summary")
  assert(review_flow_payload.fetch("review_receipt").fetch("decision_recorded") == false, "review flow preview must not record review intent in preview mode")
  assert(review_flow_payload.fetch("runtime_owned") == true, "review flow preview must remain Runtime-owned")
  assert(review_flow_payload.fetch("go_runtime_backed") == true, "review flow preview must be Go Runtime-backed")
  assert(review_flow_payload.fetch("kde_policy_owner") == false, "review flow preview must not make KDE own backend policy")
  assert(review_flow_payload.fetch("user_visible") == true, "review flow preview must be user visible")
  assert(review_flow_payload.fetch("user_confirmation_required") == true, "review flow preview must require user confirmation")
  assert(review_flow_payload.fetch("portal_policy_review_required") == true, "review flow preview must require Portal policy review")
  assert(review_flow_payload.fetch("apply_enabled") == false, "review flow preview must not enable apply")
  assert(review_flow_payload.fetch("request_object_created") == false, "review flow preview must not create request objects")
  assert(review_flow_payload.fetch("permission_granted") == false, "review flow preview must not grant permissions")
  assert(review_flow_payload.fetch("settings_persisted") == false, "review flow preview must not persist settings")
  assert(review_flow_payload.fetch("execution_started") == false, "review flow preview must not start execution")
  assert(review_flow_payload.fetch("host_root_modified") == false, "review flow preview must not mutate the host root")
  assert(review_flow_payload.fetch("backend_details_exposed") == false, "review flow preview must not expose backend details")
  assert(!review_flow.downcase.include?("prefix"), "review flow preview must not expose implementation storage")
  assert(!review_flow.include?(".exe"), "review flow preview must not expose a Windows executable")
  assert(!review_flow.downcase.include?("proton"), "review flow preview must not expose backend implementation names")

  permission_review, permission_review_status = capture_runtime_go(
    project_root,
    runtime_go_binary,
    go_binary,
    "permission-review-preview",
    "--registry",
    "runtime/recipes/registry.json",
    "--app",
    "org.xnix.sample.notepad"
  )
  assert(permission_review_status.success?, "Go permission review preview CLI must run successfully")
  permission_payload = JSON.parse(permission_review)
  assert(permission_payload.fetch("schema_version") == "xnix.runtime.permission_review.v1", "permission review preview schema version must be stable")
  assert(permission_payload.fetch("request_type") == "permission-review-preview", "permission review preview must identify its request type")
  assert(permission_payload.fetch("plan_type") == "compatibility-permission-review-plan", "permission review preview must identify the Runtime plan type")
  assert(permission_payload.fetch("source") == "unified-settings", "permission review preview must identify the KDE settings source")
  assert(permission_payload.fetch("desktop") == "KDE Plasma", "permission review preview must target KDE Plasma")
  assert(permission_payload.fetch("runtime_method") == "GetCompatibilityPermissionReviewPlan", "permission review preview must expose the Runtime method")
  assert(permission_payload.fetch("application_id") == recipe.id, "permission review preview must preserve application identity")
  assert(permission_payload.fetch("display_name") == recipe.name, "permission review preview must preserve display names")
  assert(permission_payload.fetch("desktop_file") == "xnix-#{recipe.id}.desktop", "permission review preview must bind generated desktop files")
  assert(permission_payload.fetch("review_state") == "planned", "permission review preview must stay planned")
  assert(permission_payload.fetch("permission_count") == 7, "permission review preview must expose seven permissions")
  assert(permission_payload.fetch("allow_count") == 1, "permission review preview must count allowed defaults")
  assert(permission_payload.fetch("ask_count") == 5, "permission review preview must count ask defaults")
  assert(permission_payload.fetch("deny_count") == 1, "permission review preview must count denied defaults")
  assert(permission_payload.fetch("permissions").map { |permission| permission.fetch("id") } == %w[documents downloads camera network clipboard print screenshot], "permission review preview must preserve permission order")
  assert(permission_payload.fetch("permissions").find { |permission| permission.fetch("id") == "network" }.fetch("decision") == "allow", "permission review preview must expose network defaults")
  assert(permission_payload.fetch("permissions").find { |permission| permission.fetch("id") == "camera" }.fetch("decision") == "deny", "permission review preview must expose camera defaults")
  assert(permission_payload.fetch("permissions").all? { |permission| !permission.fetch("change_pending") }, "permission review preview must not mark pending changes")
  assert(permission_payload.fetch("permissions").all? { |permission| !permission.fetch("request_object_created") }, "permission review preview must not create request objects")
  assert(permission_payload.fetch("permissions").all? { |permission| !permission.fetch("permission_granted") }, "permission review preview must not grant permissions")
  assert(permission_payload.fetch("permissions").all? { |permission| !permission.fetch("direct_access_allowed") }, "permission review preview must not allow direct access")
  assert(permission_payload.fetch("user_review_required") == true, "permission review preview must require user review")
  assert(permission_payload.fetch("portal_review_required") == true, "permission review preview must require Portal review")
  assert(permission_payload.fetch("permission_changes_applied") == false, "permission review preview must not apply permission changes")
  assert(permission_payload.fetch("request_objects_created") == false, "permission review preview must not create request objects")
  assert(permission_payload.fetch("permissions_granted") == false, "permission review preview must not grant permissions")
  assert(permission_payload.fetch("settings_persisted") == false, "permission review preview must not persist settings")
  assert(permission_payload.fetch("settings_persistence_enabled") == false, "permission review preview must keep settings persistence disabled")
  assert(permission_payload.fetch("host_permission_changed") == false, "permission review preview must not change host permissions")
  assert(permission_payload.fetch("host_root_modified") == false, "permission review preview must not mutate the host root")
  assert(permission_payload.fetch("backend_details_exposed") == false, "permission review preview must not expose backend details")
  assert(!permission_review.downcase.include?("prefix"), "permission review preview must not expose implementation storage")
  assert(!permission_review.include?(".exe"), "permission review preview must not expose a Windows executable")
  assert(!permission_review.downcase.include?("proton"), "permission review preview must not expose backend implementation names")

  desktop_resource_bridge, desktop_resource_bridge_status = capture_runtime_go(
    project_root,
    runtime_go_binary,
    go_binary,
    "desktop-resource-bridge-preview",
    "--registry",
    "runtime/recipes/registry.json",
    "--app",
    "org.xnix.sample.notepad"
  )
  assert(desktop_resource_bridge_status.success?, "Go desktop resource bridge preview CLI must run successfully")
  bridge_payload = JSON.parse(desktop_resource_bridge)
  assert(bridge_payload.fetch("schema_version") == "xnix.runtime.desktop_resource_bridge.v1", "desktop resource bridge preview schema version must be stable")
  assert(bridge_payload.fetch("request_type") == "desktop-resource-bridge-preview", "desktop resource bridge preview must identify its request type")
  assert(bridge_payload.fetch("plan_type") == "desktop-resource-bridge-plan", "desktop resource bridge preview must identify the Runtime plan type")
  assert(bridge_payload.fetch("source") == "runtime-resource-boundary", "desktop resource bridge preview must identify Runtime resource boundaries")
  assert(bridge_payload.fetch("desktop") == "KDE Plasma", "desktop resource bridge preview must target KDE Plasma")
  assert(bridge_payload.fetch("runtime_method") == "GetDesktopResourceBridgePlan", "desktop resource bridge preview must expose the Runtime method")
  assert(bridge_payload.fetch("application_id") == recipe.id, "desktop resource bridge preview must preserve application identity")
  assert(bridge_payload.fetch("display_name") == recipe.name, "desktop resource bridge preview must preserve display names")
  assert(bridge_payload.fetch("desktop_file") == "xnix-#{recipe.id}.desktop", "desktop resource bridge preview must bind generated desktop files")
  assert(bridge_payload.fetch("bridge_state") == "planned", "desktop resource bridge preview must stay planned")
  assert(bridge_payload.fetch("resource_count") == 5, "desktop resource bridge preview must expose five resource bridges")
  assert(bridge_payload.fetch("resources").map { |resource| resource.fetch("id") } == %w[file-open uri-open print clipboard screenshot], "desktop resource bridge preview must preserve resource order")
  assert(bridge_payload.fetch("resources").all? { |resource| resource.fetch("runtime_method") == "GetPortalRequestPlan" }, "desktop resource bridge preview resources must point at Portal request plans")
  assert(bridge_payload.fetch("resources").all? { |resource| resource.fetch("portal_required") }, "desktop resource bridge preview resources must require Portals")
  assert(bridge_payload.fetch("resources").all? { |resource| resource.fetch("user_approval_required") }, "desktop resource bridge preview resources must require user approval")
  assert(bridge_payload.fetch("resources").all? { |resource| !resource.fetch("bridge_enabled") }, "desktop resource bridge preview resources must keep bridges disabled")
  assert(bridge_payload.fetch("resources").all? { |resource| !resource.fetch("request_created") }, "desktop resource bridge preview resources must not create Portal requests")
  assert(bridge_payload.fetch("resources").all? { |resource| !resource.fetch("direct_backend_access_allowed") }, "desktop resource bridge preview resources must block direct access")
  assert(bridge_payload.fetch("portal_mediated") == true, "desktop resource bridge preview must require Portal mediation")
  assert(bridge_payload.fetch("file_bridge_planned") == true, "desktop resource bridge preview must plan file bridges")
  assert(bridge_payload.fetch("uri_bridge_planned") == true, "desktop resource bridge preview must plan URI bridges")
  assert(bridge_payload.fetch("print_bridge_planned") == true, "desktop resource bridge preview must plan print bridges")
  assert(bridge_payload.fetch("clipboard_bridge_planned") == true, "desktop resource bridge preview must plan clipboard bridges")
  assert(bridge_payload.fetch("screenshot_bridge_planned") == true, "desktop resource bridge preview must plan screenshot bridges")
  assert(bridge_payload.fetch("bridges_enabled") == false, "desktop resource bridge preview must not enable bridges")
  assert(bridge_payload.fetch("requests_created") == false, "desktop resource bridge preview must not create requests")
  assert(bridge_payload.fetch("backend_process_started") == false, "desktop resource bridge preview must not start backend processes")
  assert(bridge_payload.fetch("direct_host_file_access") == false, "desktop resource bridge preview must not grant direct host file access")
  assert(bridge_payload.fetch("direct_clipboard_access") == false, "desktop resource bridge preview must not grant direct clipboard access")
  assert(bridge_payload.fetch("direct_print_access") == false, "desktop resource bridge preview must not grant direct print access")
  assert(bridge_payload.fetch("host_root_modified") == false, "desktop resource bridge preview must not mutate the host root")
  assert(bridge_payload.fetch("backend_details_exposed") == false, "desktop resource bridge preview must not expose backend details")
  assert(!desktop_resource_bridge.downcase.include?("prefix"), "desktop resource bridge preview must not expose implementation storage")
  assert(!desktop_resource_bridge.include?(".exe"), "desktop resource bridge preview must not expose a Windows executable")
  assert(!desktop_resource_bridge.downcase.include?("proton"), "desktop resource bridge preview must not expose backend implementation names")
  assert(!desktop_resource_bridge.downcase.match?(%r{/users|/home|/var|/opt|/tmp}), "desktop resource bridge preview must not expose host paths")

  portal_request, portal_request_status = capture_runtime_go(
    project_root,
    runtime_go_binary,
    go_binary,
    "portal-request-preview",
    "--registry",
    "runtime/recipes/registry.json",
    "--app",
    "org.xnix.sample.notepad",
    "--operation",
    "file-open",
    "--reason",
    "Open a selected document."
  )
  assert(portal_request_status.success?, "Go Portal request preview CLI must run successfully")
  portal_payload = JSON.parse(portal_request)
  assert(portal_payload.fetch("schema_version") == "xnix.runtime.portal_request.v1", "Portal request preview schema version must be stable")
  assert(portal_payload.fetch("request_type") == "portal-request-preview", "Portal request preview must identify its request type")
  assert(portal_payload.fetch("source") == "runtime-portal-request-plan", "Portal request preview must identify the Runtime Portal plan source")
  assert(portal_payload.fetch("desktop") == "KDE Plasma", "Portal request preview must target KDE Plasma")
  assert(portal_payload.fetch("runtime_method") == "GetPortalRequestPlan", "Portal request preview must expose the Runtime method")
  assert(portal_payload.fetch("application_id") == recipe.id, "Portal request preview must preserve application identity")
  assert(portal_payload.fetch("display_name") == recipe.name, "Portal request preview must preserve display names")
  assert(portal_payload.fetch("desktop_file") == "xnix-#{recipe.id}.desktop", "Portal request preview must bind generated desktop files")
  assert(portal_payload.fetch("operation") == "file-open", "Portal request preview must preserve operation")
  assert(portal_payload.fetch("reason") == "Open a selected document.", "Portal request preview must preserve the user-facing reason")
  assert(portal_payload.fetch("decision") == "ask", "file-open Portal request preview must ask by default")
  assert(portal_payload.fetch("request_allowed") == true, "file-open Portal request preview must allow mediated requests")
  assert(portal_payload.fetch("runtime_owned") == true, "Portal request preview must remain Runtime-owned")
  assert(portal_payload.fetch("kde_policy_owner") == false, "Portal request preview must not make KDE own policy")
  assert(portal_payload.fetch("portal").fetch("destination") == "org.freedesktop.portal.Desktop", "Portal request preview must target the Portal service")
  assert(portal_payload.fetch("portal").fetch("interface") == "org.freedesktop.portal.FileChooser", "file-open Portal request preview must target FileChooser")
  assert(portal_payload.fetch("portal").fetch("method") == "OpenFile", "file-open Portal request preview must call OpenFile")
  assert(portal_payload.fetch("portal").fetch("object_path") == "/org/freedesktop/portal/desktop", "Portal request preview must expose the standard Portal object path")
  assert(portal_payload.fetch("request").fetch("object_path_required") == true, "Portal request preview must require request object paths")
  assert(portal_payload.fetch("request").fetch("request_object_created") == false, "Portal request preview must not create request objects")
  assert(portal_payload.fetch("request").fetch("handle_token") == "xnix_org_xnix_sample_notepad_file_open", "Portal request preview must produce deterministic handle tokens")
  assert(portal_payload.fetch("request").fetch("user_mediation_required") == true, "Portal request preview must require user mediation")
  assert(portal_payload.fetch("request").fetch("resources") == %w[documents downloads selected-files], "Portal request preview must expose selected file resources")
  assert(portal_payload.fetch("request").fetch("runtime_policy_owner") == true, "Portal request preview policy must be Runtime-owned")
  assert(portal_payload.fetch("request").fetch("desktop_shell_policy_owner") == false, "Portal request preview policy must not be KDE-owned")
  assert(portal_payload.fetch("completion").fetch("signal") == "Response", "Portal request preview must wait for Response")
  assert(portal_payload.fetch("completion").fetch("result_owner") == "Runtime", "Portal request preview results must return to Runtime")
  assert(portal_payload.fetch("denied").nil?, "allowed Portal request preview must not include denial guidance")
  assert(portal_payload.fetch("safety").fetch("direct_access_allowed") == false, "Portal request preview must deny direct access")
  assert(portal_payload.fetch("safety").fetch("portal_required") == true, "Portal request preview must require Portals")
  assert(portal_payload.fetch("safety").fetch("permission_granted") == false, "Portal request preview must not grant permissions")
  assert(portal_payload.fetch("safety").fetch("host_permission_changed") == false, "Portal request preview must not change host permissions")
  assert(portal_payload.fetch("safety").fetch("host_root_modified") == false, "Portal request preview must not mutate the host root")
  assert(portal_payload.fetch("safety").fetch("backend_details_exposed") == false, "Portal request preview must not expose backend details")
  assert(!portal_request.downcase.include?("prefix"), "Portal request preview must not expose implementation storage")
  assert(!portal_request.include?(".exe"), "Portal request preview must not expose a Windows executable")
  assert(!portal_request.downcase.include?("proton"), "Portal request preview must not expose backend implementation names")
  assert(!portal_request.downcase.match?(%r{/users|/home|/var|/opt|/tmp}), "Portal request preview must not expose host paths")

  denied_portal_request, denied_portal_request_status = capture_runtime_go(
    project_root,
    runtime_go_binary,
    go_binary,
    "portal-request-preview",
    "--registry",
    "runtime/recipes/registry.json",
    "--app",
    "org.xnix.sample.notepad",
    "--operation",
    "camera"
  )
  assert(denied_portal_request_status.success?, "Go denied Portal request preview CLI must run successfully")
  denied_portal_payload = JSON.parse(denied_portal_request)
  assert(denied_portal_payload.fetch("decision") == "deny", "camera Portal request preview must deny by default")
  assert(denied_portal_payload.fetch("request_allowed") == false, "denied Portal request preview must not allow requests")
  assert(denied_portal_payload.fetch("denied").fetch("next_action") == "open-compatibility-settings", "denied Portal request preview must guide users to settings")
  assert(denied_portal_payload.fetch("safety").fetch("permission_granted") == false, "denied Portal request preview must not grant permissions")

  tray_status, tray_status_result = capture_runtime_go(
    project_root,
    runtime_go_binary,
    go_binary,
    "tray-status-preview",
    "--registry",
    "runtime/recipes/registry.json",
    "--app",
    "org.xnix.sample.notepad"
  )
  assert(tray_status_result.success?, "Go tray status preview CLI must run successfully")
  tray_payload = JSON.parse(tray_status)
  assert(tray_payload.fetch("schema_version") == "xnix.runtime.tray_status.v1", "tray status preview schema version must be stable")
  assert(tray_payload.fetch("status_type") == "tray-status-preview", "tray status preview must identify its status type")
  assert(tray_payload.fetch("desktop") == "KDE Plasma", "tray status preview must target KDE Plasma")
  assert(tray_payload.fetch("application_id") == recipe.id, "tray status preview must preserve application identity")
  assert(tray_payload.fetch("display_name") == recipe.name, "tray status preview must preserve display names")
  assert(tray_payload.fetch("desktop_file") == "xnix-#{recipe.id}.desktop", "tray status preview must bind generated desktop files")
  assert(tray_payload.fetch("runtime_activity").fetch("registered_application_count") == 1, "tray status preview must register one application")
  assert(tray_payload.fetch("runtime_activity").fetch("active_application_count") == 0, "tray status preview must not claim live application activity")
  assert(tray_payload.fetch("runtime_activity").fetch("attention_required_count") == 0, "tray status preview must not invent attention requests")
  assert(tray_payload.fetch("compatibility_status").fetch("state") == "ready", "tray status preview must expose a ready compatibility state")
  assert(tray_payload.fetch("tray_bridge").fetch("state") == "planned", "tray status preview must keep live tray bridging planned")
  assert(tray_payload.fetch("tray_bridge").fetch("bridged_tray_application_count") == 0, "tray status preview must not claim bridged tray applications")
  assert(tray_payload.fetch("actions") == ["open-compatibility-center", "open-settings"], "tray status preview must expose KDE navigation actions")
  assert(tray_payload.fetch("runtime_owned") == true, "tray status preview must remain Runtime-owned")
  assert(tray_payload.fetch("kde_policy_owner") == false, "tray status preview must not make KDE own backend policy")
  assert(tray_payload.fetch("user_visible") == true, "tray status preview must be user visible")
  assert(tray_payload.fetch("live_backend_bridge_enabled") == false, "tray status preview must keep live tray bridges disabled")
  assert(tray_payload.fetch("bridge_configuration_persisted") == false, "tray status preview must keep bridge persistence disabled")
  assert(tray_payload.fetch("host_root_modified") == false, "tray status preview must not mutate the host root")
  assert(tray_payload.fetch("backend_details_exposed") == false, "tray status preview must not expose backend details")
  assert(!tray_status.downcase.include?("prefix"), "tray status preview must not expose implementation storage")
  assert(!tray_status.include?(".exe"), "tray status preview must not expose a Windows executable")

  window_identity, window_status = capture_runtime_go(
    project_root,
    runtime_go_binary,
    go_binary,
    "window-identity-preview",
    "--registry",
    "runtime/recipes/registry.json",
    "--app",
    "org.xnix.sample.notepad"
  )
  assert(window_status.success?, "Go window identity preview CLI must run successfully")
  window_payload = JSON.parse(window_identity)
  assert(window_payload.fetch("schema_version") == "xnix.runtime.window_identity.v1", "window identity preview schema version must be stable")
  assert(window_payload.fetch("desktop") == "KDE Plasma", "window identity preview must target KDE Plasma")
  assert(window_payload.fetch("desktop_file") == "xnix-#{recipe.id}.desktop", "window identity preview must bind generated desktop files")
  assert(window_payload.fetch("launcher_url") == "applications:xnix-#{recipe.id}.desktop", "window identity preview must expose a launcher URL")
  assert(window_payload.fetch("task_manager").fetch("grouping_key") == recipe.id, "window identity preview must group taskbar windows by Runtime application id")
  assert(window_payload.fetch("task_manager").fetch("pinning_allowed") == true, "window identity preview must allow taskbar pinning")
  assert(window_payload.fetch("task_manager").fetch("restore_allowed") == true, "window identity preview must allow restore")
  assert(window_payload.fetch("task_manager").fetch("skip_taskbar") == false, "window identity preview must keep windows visible in the taskbar")
  assert(window_payload.fetch("task_manager").fetch("show_in_switcher") == true, "window identity preview must keep windows visible in the switcher")
  assert(window_payload.fetch("kwin").fetch("script_role") == "identity-and-layout", "window identity preview must provide bounded KWin identity hints")
  assert(window_payload.fetch("kwin").fetch("window_manager_policy_only") == true, "window identity preview must keep KWin policy scoped to window identity")
  assert(window_payload.fetch("kwin").fetch("runtime_owns_backend_policy") == true, "window identity preview must keep backend policy in the Runtime")
  assert(window_payload.fetch("runtime_owned") == true, "window identity preview must remain Runtime-owned")
  assert(window_payload.fetch("kde_policy_owner") == false, "window identity preview must not make KDE own backend policy")
  assert(window_payload.fetch("host_root_modified") == false, "window identity preview must not mutate the host root")
  assert(window_payload.fetch("backend_details_exposed") == false, "window identity preview must not expose backend details")
  assert(!window_identity.downcase.include?("prefix"), "window identity preview must not expose implementation storage")
  assert(!window_identity.include?(".exe"), "window identity preview must not expose a Windows executable")
end

puts "PASS: Go Runtime desktop identity plan unit tests"
