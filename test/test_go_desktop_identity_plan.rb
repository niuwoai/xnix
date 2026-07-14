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
assert(File.read(File.join(project_root, "cmd/xnix-runtime-go/main.go")).include?("desktop-entry-preview"), "Go Runtime CLI must render desktop entry previews")
assert(File.read(File.join(project_root, "cmd/xnix-runtime-go/main.go")).include?("krunner-query-preview"), "Go Runtime CLI must render KDE KRunner query previews")
assert(File.read(File.join(project_root, "cmd/xnix-runtime-go/main.go")).include?("mimeapps-preview"), "Go Runtime CLI must render MIME association previews")
assert(File.read(File.join(project_root, "cmd/xnix-runtime-go/main.go")).include?("notification-preview"), "Go Runtime CLI must render KDE notification previews")
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
