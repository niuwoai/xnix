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
assert(File.read(File.join(project_root, "cmd/xnix-runtime-go/main.go")).include?("desktop-icon-preview"), "Go Runtime CLI must render desktop icon previews")
assert(File.read(File.join(project_root, "cmd/xnix-runtime-go/main.go")).include?("desktop-activation-status-preview"), "Go Runtime CLI must render desktop activation status previews")
assert(File.read(File.join(project_root, "cmd/xnix-runtime-go/main.go")).include?("desktop-resource-bridge-preview"), "Go Runtime CLI must render KDE desktop resource bridge previews")
assert(File.read(File.join(project_root, "cmd/xnix-runtime-go/main.go")).include?("execution-readiness-preview"), "Go Runtime CLI must render KDE execution readiness previews")
assert(File.read(File.join(project_root, "cmd/xnix-runtime-go/main.go")).include?("file-open-preview"), "Go Runtime CLI must render Dolphin file-open previews")
assert(File.read(File.join(project_root, "cmd/xnix-runtime-go/main.go")).include?("krunner-query-preview"), "Go Runtime CLI must render KDE KRunner query previews")
assert(File.read(File.join(project_root, "cmd/xnix-runtime-go/main.go")).include?("launch-intent-preview"), "Go Runtime CLI must render KDE launch intent previews")
assert(File.read(File.join(project_root, "cmd/xnix-runtime-go/main.go")).include?("execution-request-preview"), "Go Runtime CLI must render KDE execution request previews")
assert(File.read(File.join(project_root, "cmd/xnix-runtime-go/main.go")).include?("execution-review-preview"), "Go Runtime CLI must render KDE execution review previews")
assert(File.read(File.join(project_root, "cmd/xnix-runtime-go/main.go")).include?("execution-decision-preview"), "Go Runtime CLI must render KDE execution decision previews")
assert(File.read(File.join(project_root, "cmd/xnix-runtime-go/main.go")).include?("execution-preflight-preview"), "Go Runtime CLI must render KDE execution preflight previews")
assert(File.read(File.join(project_root, "cmd/xnix-runtime-go/main.go")).include?("execution-resource-grant-preview"), "Go Runtime CLI must render KDE execution resource grant previews")
assert(File.read(File.join(project_root, "cmd/xnix-runtime-go/main.go")).include?("execution-transaction-preview"), "Go Runtime CLI must render KDE execution transaction previews")
assert(File.read(File.join(project_root, "cmd/xnix-runtime-go/main.go")).include?("execution-session-preview"), "Go Runtime CLI must render KDE execution session previews")
assert(File.read(File.join(project_root, "cmd/xnix-runtime-go/main.go")).include?("execution-session-status-preview"), "Go Runtime CLI must render KDE execution session status previews")
assert(File.read(File.join(project_root, "cmd/xnix-runtime-go/main.go")).include?("kde-action-card-deck-preview"), "Go Runtime CLI must render KDE action card deck previews")
assert(File.read(File.join(project_root, "cmd/xnix-runtime-go/main.go")).include?("kde-action-card-preview"), "Go Runtime CLI must render KDE action card previews")
assert(File.read(File.join(project_root, "cmd/xnix-runtime-go/main.go")).include?("kde-center-page-preview"), "Go Runtime CLI must render KDE Compatibility Center page previews")
assert(File.read(File.join(project_root, "cmd/xnix-runtime-go/main.go")).include?("kde-center-page-sections-preview"), "Go Runtime CLI must render KDE Compatibility Center page section previews")
assert(File.read(File.join(project_root, "cmd/xnix-runtime-go/main.go")).include?("kde-center-page-section-detail-preview"), "Go Runtime CLI must render KDE Compatibility Center page section detail previews")
assert(File.read(File.join(project_root, "cmd/xnix-runtime-go/main.go")).include?("kde-action-preflight-preview"), "Go Runtime CLI must render KDE action preflight previews")
assert(File.read(File.join(project_root, "cmd/xnix-runtime-go/main.go")).include?("kde-action-queue-preview"), "Go Runtime CLI must render KDE action queue previews")
assert(File.read(File.join(project_root, "cmd/xnix-runtime-go/main.go")).include?("kde-action-receipt-preview"), "Go Runtime CLI must render KDE action receipt previews")
assert(File.read(File.join(project_root, "cmd/xnix-runtime-go/main.go")).include?("kde-action-review-preview"), "Go Runtime CLI must render KDE action review previews")
assert(File.read(File.join(project_root, "cmd/xnix-runtime-go/main.go")).include?("kde-action-status-preview"), "Go Runtime CLI must render KDE action status previews")
assert(File.read(File.join(project_root, "cmd/xnix-runtime-go/main.go")).include?("kde-entrypoint-action-preview"), "Go Runtime CLI must render KDE entrypoint action previews")
assert(File.read(File.join(project_root, "cmd/xnix-runtime-go/main.go")).include?("kde-entrypoints-preview"), "Go Runtime CLI must render KDE entrypoint previews")
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

  desktop_icon, desktop_icon_status = capture_runtime_go(
    project_root,
    runtime_go_binary,
    go_binary,
    "desktop-icon-preview",
    "--registry",
    "runtime/recipes/registry.json",
    "--app",
    "org.xnix.sample.notepad"
  )
  assert(desktop_icon_status.success?, "Go desktop icon preview CLI must run successfully")
  desktop_icon_payload = JSON.parse(desktop_icon)
  assert(desktop_icon_payload.fetch("schema_version") == "xnix.runtime.desktop_icon.v1", "desktop icon preview schema version must be stable")
  assert(desktop_icon_payload.fetch("request_type") == "desktop-icon-preview", "desktop icon preview must identify its request type")
  assert(desktop_icon_payload.fetch("plan_type") == "desktop-icon-plan", "desktop icon preview must identify its plan type")
  assert(desktop_icon_payload.fetch("runtime_method") == "GetDesktopIconPlan", "desktop icon preview must expose the Runtime method")
  assert(desktop_icon_payload.fetch("desktop_file") == "xnix-#{recipe.id}.desktop", "desktop icon preview must bind generated desktop files")
  assert(desktop_icon_payload.fetch("launcher_url") == "applications:xnix-#{recipe.id}.desktop", "desktop icon preview must expose application launcher URLs")
  assert(desktop_icon_payload.fetch("target_directory") == "xdg-desktop-dir", "desktop icon preview must keep the target directory abstract")
  assert(desktop_icon_payload.fetch("placement") == "user-desktop", "desktop icon preview must target the user desktop")
  assert(desktop_icon_payload.fetch("standard_desktop_entry") == true, "desktop icon preview must reuse standard desktop entries")
  assert(desktop_icon_payload.fetch("desktop_icon_visible") == true, "desktop icon preview must be visible to KDE")
  assert(desktop_icon_payload.fetch("runtime_owned") == true, "desktop icon preview must remain Runtime-owned")
  assert(desktop_icon_payload.fetch("go_runtime_backed") == true, "desktop icon preview must be backed by Go Runtime product logic")
  assert(desktop_icon_payload.fetch("kde_policy_owner") == false, "desktop icon preview must not make KDE own policy")
  assert(desktop_icon_payload.fetch("desktop_file_copy_enabled") == false, "desktop icon preview must not copy desktop files")
  assert(desktop_icon_payload.fetch("desktop_file_write_enabled") == false, "desktop icon preview must not write desktop files")
  assert(desktop_icon_payload.fetch("icon_placement_persisted") == false, "desktop icon preview must not persist placement")
  assert(desktop_icon_payload.fetch("launch_enabled") == false, "desktop icon preview must not enable launch")
  assert(desktop_icon_payload.fetch("host_root_modified") == false, "desktop icon preview must not mutate the host root")
  assert(desktop_icon_payload.fetch("backend_details_exposed") == false, "desktop icon preview must not expose backend details")
  assert(!desktop_icon.downcase.include?("prefix"), "desktop icon preview must not expose implementation storage")
  assert(!desktop_icon.include?(".exe"), "desktop icon preview must not expose a Windows executable")
  assert(!desktop_icon.downcase.include?("virtual machine"), "desktop icon preview must not expose implementation labels")

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

  execution_readiness, execution_readiness_status = capture_runtime_go(
    project_root,
    runtime_go_binary,
    go_binary,
    "execution-readiness-preview",
    "--registry",
    "runtime/recipes/registry.json",
    "--app",
    "org.xnix.sample.notepad"
  )
  assert(execution_readiness_status.success?, "Go execution readiness preview CLI must run successfully")
  execution_readiness_payload = JSON.parse(execution_readiness)
  assert(execution_readiness_payload.fetch("schema_version") == "xnix.runtime.launch_readiness.v1", "execution readiness preview schema version must be stable")
  assert(execution_readiness_payload.fetch("request_type") == "execution-readiness-preview", "execution readiness preview must identify its request type")
  assert(execution_readiness_payload.fetch("readiness_type") == "compatibility-execution-readiness", "execution readiness preview must identify the readiness type")
  assert(execution_readiness_payload.fetch("source") == "compatibility-center", "execution readiness preview must identify the Compatibility Center source")
  assert(execution_readiness_payload.fetch("desktop") == "KDE Plasma", "execution readiness preview must target KDE Plasma")
  assert(execution_readiness_payload.fetch("runtime_method") == "GetExecutionReadiness", "execution readiness preview must expose the Runtime method")
  readiness_application = execution_readiness_payload.fetch("application")
  assert(readiness_application.fetch("id") == recipe.id, "execution readiness preview must preserve application identity")
  assert(readiness_application.fetch("name") == recipe.name, "execution readiness preview must preserve display names")
  assert(readiness_application.fetch("desktop_file") == "xnix-#{recipe.id}.desktop", "execution readiness preview must bind generated desktop files")
  assert(readiness_application.fetch("launcher_command") == ["xnix-compat-launch", "--app", recipe.id, "%U"], "execution readiness preview must expose the managed launcher command")
  readiness_profile = execution_readiness_payload.fetch("compatibility_profile")
  assert(readiness_profile.fetch("id") == "local-compatibility", "execution readiness preview must use the recommended profile")
  assert(readiness_profile.fetch("kind") == "local", "execution readiness preview must keep profile kind user-facing")
  assert(readiness_profile.fetch("ready") == false, "execution readiness preview must not mark the profile ready")
  assert(readiness_profile.fetch("launch_enabled") == false, "execution readiness preview must not enable profile launch")
  assert(readiness_profile.fetch("backend_details_exposed") == false, "execution readiness preview must hide backend details")
  assert(execution_readiness_payload.fetch("execution_state") == "blocked", "execution readiness preview must stay blocked before Runtime gates pass")
  assert(execution_readiness_payload.fetch("overall_status") == "not-ready", "execution readiness preview must not claim launch readiness")
  assert(execution_readiness_payload.fetch("gate_count") == 5, "execution readiness preview must expose five Runtime gates")
  assert(execution_readiness_payload.fetch("required_gate_count") == 2, "execution readiness preview must count required gates")
  assert(execution_readiness_payload.fetch("pending_gate_count") == 1, "execution readiness preview must count pending gates")
  assert(execution_readiness_payload.fetch("blocked_gate_count") == 1, "execution readiness preview must count blocked gates")
  assert(execution_readiness_payload.fetch("gates").map { |gate| gate.fetch("id") } == %w[recipe-validation portal-policy-review snapshot-baseline backend-binding runtime-launch-write-gate], "execution readiness preview must preserve gate order")
  assert(execution_readiness_payload.fetch("runtime_owned") == true, "execution readiness preview must remain Runtime-owned")
  assert(execution_readiness_payload.fetch("go_runtime_backed") == true, "execution readiness preview must be Go Runtime-backed")
  assert(execution_readiness_payload.fetch("kde_policy_owner") == false, "execution readiness preview must not make KDE own backend policy")
  assert(execution_readiness_payload.fetch("compatibility_center_card") == true, "execution readiness preview must be suitable for Compatibility Center cards")
  assert(execution_readiness_payload.fetch("safe_for_ai_diagnostics") == true, "execution readiness preview must be safe for AI diagnostics")
  assert(execution_readiness_payload.fetch("desktop_entry_launch_visible") == true, "execution readiness preview must allow KDE to show the desktop entry")
  assert(execution_readiness_payload.fetch("launch_allowed") == false, "execution readiness preview must not allow launch")
  assert(execution_readiness_payload.fetch("launch_enabled") == false, "execution readiness preview must not enable launch")
  assert(execution_readiness_payload.fetch("execution_request_created") == false, "execution readiness preview must not create execution requests")
  assert(execution_readiness_payload.fetch("backend_binding_ready") == false, "execution readiness preview must not claim backend binding readiness")
  assert(execution_readiness_payload.fetch("portal_policy_required") == true, "execution readiness preview must require Portal review")
  assert(execution_readiness_payload.fetch("snapshot_required") == true, "execution readiness preview must require snapshot review")
  assert(execution_readiness_payload.fetch("user_action_required") == true, "execution readiness preview must require user or Runtime action")
  assert(execution_readiness_payload.fetch("blocked_actions").include?("launch compatibility profile"), "execution readiness preview must block compatibility launch")
  assert(execution_readiness_payload.fetch("host_root_modified") == false, "execution readiness preview must not mutate the host root")
  assert(execution_readiness_payload.fetch("network_required") == false, "execution readiness preview must not require network access")
  assert(execution_readiness_payload.fetch("backend_details_exposed") == false, "execution readiness preview must not expose backend details")
  assert(!execution_readiness.downcase.include?("prefix"), "execution readiness preview must not expose implementation storage")
  assert(!execution_readiness.include?(".exe"), "execution readiness preview must not expose a Windows executable")
  assert(!execution_readiness.downcase.include?("program files"), "execution readiness preview must not expose Windows paths")
  assert(!execution_readiness.downcase.include?("qemu-system"), "execution readiness preview must not expose VM implementation commands")
  assert(!execution_readiness.downcase.include?("proton"), "execution readiness preview must not expose backend implementation names")
  assert(!execution_readiness.downcase.include?("wine "), "execution readiness preview must not expose backend implementation names")
  assert(!execution_readiness.downcase.include?("virtual machine"), "execution readiness preview must not expose implementation labels")

  launch_intent, launch_intent_status = capture_runtime_go(
    project_root,
    runtime_go_binary,
    go_binary,
    "launch-intent-preview",
    "--registry",
    "runtime/recipes/registry.json",
    "--app",
    "org.xnix.sample.notepad",
    "file:///home/test/Documents/example.txt"
  )
  assert(launch_intent_status.success?, "Go launch intent preview CLI must run successfully")
  launch_intent_payload = JSON.parse(launch_intent)
  assert(launch_intent_payload.fetch("schema_version") == "xnix.runtime.launch_intent.v1", "launch intent preview schema version must be stable")
  assert(launch_intent_payload.fetch("request_type") == "launch-intent-preview", "launch intent preview must identify its request type")
  assert(launch_intent_payload.fetch("intent_type") == "runtime-launch-intent", "launch intent preview must identify the Runtime launch intent")
  assert(launch_intent_payload.fetch("source") == "desktop-launcher", "launch intent preview must identify the desktop source")
  assert(launch_intent_payload.fetch("desktop") == "KDE Plasma", "launch intent preview must target KDE Plasma")
  assert(launch_intent_payload.fetch("runtime_method") == "Launch", "launch intent preview must target the Runtime Launch method")
  assert(launch_intent_payload.fetch("read_method") == "GetLaunchIntent", "launch intent preview must expose the Runtime read method")
  assert(launch_intent_payload.fetch("application_id") == recipe.id, "launch intent preview must preserve application identity")
  assert(launch_intent_payload.fetch("application_name") == recipe.name, "launch intent preview must preserve display names")
  assert(launch_intent_payload.fetch("desktop_file") == "xnix-#{recipe.id}.desktop", "launch intent preview must bind generated desktop files")
  assert(launch_intent_payload.fetch("launcher_command") == ["xnix-compat-launch", "--app", recipe.id, "%U"], "launch intent preview must expose the managed launcher command")
  assert(launch_intent_payload.fetch("execution_state") == "blocked", "launch intent preview must stay blocked before Runtime gates pass")
  assert(launch_intent_payload.fetch("overall_status") == "not-ready", "launch intent preview must not claim launch readiness")
  assert(launch_intent_payload.fetch("write_gate_decision") == "blocked-until-production-backend", "launch intent preview must expose the Launch write gate")
  assert(launch_intent_payload.fetch("denial_error_name") == "org.xnix.Compatibility1.Error.WriteMethodDisabled", "launch intent preview must expose the write-gate error")
  assert(launch_intent_payload.fetch("portal_required") == true, "launch intent preview with files must require Portal review")
  assert(launch_intent_payload.fetch("snapshot_required") == true, "launch intent preview must require snapshot review")
  assert(launch_intent_payload.fetch("file_count") == 1, "launch intent preview must count selected files")
  assert(launch_intent_payload.fetch("file_uris") == ["file:///home/test/Documents/example.txt"], "launch intent preview must preserve file URIs")
  launch_profile = launch_intent_payload.fetch("compatibility_profile")
  assert(launch_profile.fetch("id") == "local-compatibility", "launch intent preview must use the recommended profile")
  assert(launch_profile.fetch("ready") == false, "launch intent preview must not mark the profile ready")
  assert(launch_profile.fetch("launch_enabled") == false, "launch intent preview must not enable profile launch")
  assert(launch_profile.fetch("backend_details_exposed") == false, "launch intent preview must hide backend details")
  launch_run_plan = launch_intent_payload.fetch("run_plan")
  assert(launch_run_plan.fetch("plan_type") == "compatibility-run", "launch intent preview must include a compatibility run plan")
  assert(launch_run_plan.fetch("strategy") == "automatic-managed", "launch intent preview must expose a desktop-safe strategy")
  assert(launch_run_plan.fetch("backend_details_exposed") == false, "launch intent run plan must hide backend details")
  assert(launch_run_plan.fetch("backend_ready") == false, "launch intent run plan must not claim backend readiness")
  assert(launch_run_plan.fetch("portal_policy_required") == true, "launch intent run plan must require Portal policy review")
  assert(launch_run_plan.fetch("snapshot_before_risky_change") == true, "launch intent run plan must require snapshot preflight")
  assert(launch_run_plan.fetch("runtime_write_gate_required") == true, "launch intent run plan must require the Runtime write gate")
  assert(launch_run_plan.fetch("execution_request_created") == false, "launch intent run plan must not create execution requests")
  assert(launch_intent_payload.fetch("runtime_owned") == true, "launch intent preview must remain Runtime-owned")
  assert(launch_intent_payload.fetch("go_runtime_backed") == true, "launch intent preview must be Go Runtime-backed")
  assert(launch_intent_payload.fetch("kde_policy_owner") == false, "launch intent preview must not make KDE own backend policy")
  assert(launch_intent_payload.fetch("standard_desktop_entry") == true, "launch intent preview must preserve standard desktop entries")
  assert(launch_intent_payload.fetch("launch_uses_runtime") == true, "launch intent preview must route through the Runtime")
  assert(launch_intent_payload.fetch("desktop_entry_launch_visible") == true, "launch intent preview must allow KDE to show launcher entries")
  assert(launch_intent_payload.fetch("launch_allowed") == false, "launch intent preview must not allow launch")
  assert(launch_intent_payload.fetch("launch_enabled") == false, "launch intent preview must not enable launch")
  assert(launch_intent_payload.fetch("execution_request_created") == false, "launch intent preview must not create execution requests")
  assert(launch_intent_payload.fetch("execution_started") == false, "launch intent preview must not start execution")
  assert(launch_intent_payload.fetch("backend_binding_ready") == false, "launch intent preview must not claim backend binding readiness")
  assert(launch_intent_payload.fetch("request_object_created") == false, "launch intent preview must not create request objects")
  assert(launch_intent_payload.fetch("permission_granted") == false, "launch intent preview must not grant permissions")
  assert(launch_intent_payload.fetch("host_root_modified") == false, "launch intent preview must not mutate the host root")
  assert(launch_intent_payload.fetch("network_required") == false, "launch intent preview must not require network access")
  assert(launch_intent_payload.fetch("backend_details_exposed") == false, "launch intent preview must not expose backend details")
  assert(launch_intent_payload.fetch("blocked_actions").include?("start compatibility profile from KDE"), "launch intent preview must block direct KDE starts")
  assert(!launch_intent.downcase.include?("prefix"), "launch intent preview must not expose implementation storage")
  assert(!launch_intent.include?(".exe"), "launch intent preview must not expose a Windows executable")
  assert(!launch_intent.downcase.include?("program files"), "launch intent preview must not expose Windows paths")
  assert(!launch_intent.downcase.include?("qemu-system"), "launch intent preview must not expose VM implementation commands")
  assert(!launch_intent.downcase.include?("proton"), "launch intent preview must not expose backend implementation names")
  assert(!launch_intent.downcase.include?("wine "), "launch intent preview must not expose backend implementation names")
  assert(!launch_intent.downcase.include?("virtual machine"), "launch intent preview must not expose implementation labels")

  execution_request, execution_request_status = capture_runtime_go(
    project_root,
    runtime_go_binary,
    go_binary,
    "execution-request-preview",
    "--registry",
    "runtime/recipes/registry.json",
    "--app",
    "org.xnix.sample.notepad",
    "file:///home/test/Documents/example.txt"
  )
  assert(execution_request_status.success?, "Go execution request preview CLI must run successfully")
  execution_request_payload = JSON.parse(execution_request)
  assert(execution_request_payload.fetch("schema_version") == "xnix.runtime.request_intake.v1", "execution request preview schema version must be stable")
  assert(execution_request_payload.fetch("request_type") == "execution-request-preview", "execution request preview must identify its request type")
  assert(execution_request_payload.fetch("intent_type") == "runtime-launch-intent", "execution request preview must preserve Runtime launch intent identity")
  assert(execution_request_payload.fetch("request_state") == "blocked", "execution request preview must remain blocked")
  assert(execution_request_payload.fetch("source") == "runtime-launch-intent", "execution request preview must derive from launch intent")
  assert(execution_request_payload.fetch("desktop") == "KDE Plasma", "execution request preview must target KDE Plasma")
  assert(execution_request_payload.fetch("runtime_method") == "Launch", "execution request preview must target the Runtime Launch method")
  assert(execution_request_payload.fetch("read_method") == "GetExecutionRequestPreview", "execution request preview must expose the Runtime read method")
  assert(execution_request_payload.fetch("application_id") == recipe.id, "execution request preview must preserve application identity")
  assert(execution_request_payload.fetch("application_name") == recipe.name, "execution request preview must preserve display names")
  assert(execution_request_payload.fetch("desktop_file") == "xnix-#{recipe.id}.desktop", "execution request preview must bind generated desktop files")
  assert(execution_request_payload.fetch("launcher_command") == ["xnix-compat-launch", "--app", recipe.id, "%U"], "execution request preview must expose the managed launcher command")
  request_intent = execution_request_payload.fetch("launch_intent")
  assert(request_intent.fetch("source") == "desktop-launcher", "execution request preview must summarize desktop launch intent")
  assert(request_intent.fetch("intent_type") == "runtime-launch-intent", "execution request preview must keep launch intent type")
  assert(request_intent.fetch("runtime_method") == "Launch", "execution request preview must preserve launch method")
  assert(request_intent.fetch("read_method") == "GetLaunchIntent", "execution request preview must point back to the launch intent read method")
  assert(request_intent.fetch("launch_allowed") == false, "execution request preview must not allow launch through its launch intent")
  assert(request_intent.fetch("launch_enabled") == false, "execution request preview must not enable launch through its launch intent")
  assert(request_intent.fetch("request_object_created") == false, "execution request preview launch intent must not create request objects")
  assert(request_intent.fetch("execution_started") == false, "execution request preview launch intent must not start execution")
  assert(request_intent.fetch("backend_details_exposed") == false, "execution request preview launch intent must hide backend details")
  request_profile = execution_request_payload.fetch("compatibility_profile")
  assert(request_profile.fetch("id") == "local-compatibility", "execution request preview must use the recommended profile")
  assert(request_profile.fetch("kind") == "local", "execution request preview must keep profile kind user-facing")
  assert(request_profile.fetch("ready") == false, "execution request preview must not mark the profile ready")
  assert(request_profile.fetch("launch_enabled") == false, "execution request preview must not enable profile launch")
  assert(request_profile.fetch("backend_details_exposed") == false, "execution request preview must hide backend details")
  request_gates = execution_request_payload.fetch("gate_summary")
  assert(request_gates.fetch("gate_count") == 5, "execution request preview must summarize Runtime gates")
  assert(request_gates.fetch("required_gate_count") == 2, "execution request preview must count required gates")
  assert(request_gates.fetch("pending_gate_count") == 1, "execution request preview must count pending gates")
  assert(request_gates.fetch("blocked_gate_count") == 1, "execution request preview must count blocked gates")
  assert(execution_request_payload.fetch("execution_state") == "blocked", "execution request preview must stay blocked before Runtime gates pass")
  assert(execution_request_payload.fetch("overall_status") == "not-ready", "execution request preview must not claim launch readiness")
  assert(execution_request_payload.fetch("write_gate_decision") == "blocked-until-production-backend", "execution request preview must expose the Launch write gate")
  assert(execution_request_payload.fetch("denial_error_name") == "org.xnix.Compatibility1.Error.WriteMethodDisabled", "execution request preview must expose the write-gate error")
  assert(execution_request_payload.fetch("portal_required") == true, "execution request preview with files must require Portal review")
  assert(execution_request_payload.fetch("snapshot_required") == true, "execution request preview must require snapshot review")
  assert(execution_request_payload.fetch("file_count") == 1, "execution request preview must count selected files")
  assert(execution_request_payload.fetch("file_uris") == ["file:///home/test/Documents/example.txt"], "execution request preview must preserve file URIs")
  assert(request_intent.fetch("portal_required") == true, "execution request preview launch intent must require Portal review for files")
  assert(request_intent.fetch("file_count") == 1, "execution request preview launch intent must count selected files")
  assert(request_intent.fetch("file_uris") == ["file:///home/test/Documents/example.txt"], "execution request preview launch intent must preserve file URIs")
  assert(execution_request_payload.fetch("runtime_owned") == true, "execution request preview must remain Runtime-owned")
  assert(execution_request_payload.fetch("go_runtime_backed") == true, "execution request preview must be Go Runtime-backed")
  assert(execution_request_payload.fetch("kde_policy_owner") == false, "execution request preview must not make KDE own backend policy")
  assert(execution_request_payload.fetch("compatibility_center_card") == true, "execution request preview must be suitable for Compatibility Center cards")
  assert(execution_request_payload.fetch("safe_for_ai_diagnostics") == true, "execution request preview must be safe for AI diagnostics")
  assert(execution_request_payload.fetch("desktop_entry_launch_visible") == true, "execution request preview must allow KDE to show launcher entries")
  assert(execution_request_payload.fetch("launch_intent_captured") == true, "execution request preview must capture launch intent")
  assert(execution_request_payload.fetch("launch_allowed") == false, "execution request preview must not allow launch")
  assert(execution_request_payload.fetch("launch_enabled") == false, "execution request preview must not enable launch")
  assert(execution_request_payload.fetch("execution_request_created") == false, "execution request preview must not create execution requests")
  assert(execution_request_payload.fetch("execution_request_persisted") == false, "execution request preview must not persist execution requests")
  assert(execution_request_payload.fetch("execution_started") == false, "execution request preview must not start execution")
  assert(execution_request_payload.fetch("backend_binding_ready") == false, "execution request preview must not claim backend binding readiness")
  assert(execution_request_payload.fetch("request_object_created") == false, "execution request preview must not create request objects")
  assert(execution_request_payload.fetch("permission_granted") == false, "execution request preview must not grant permissions")
  assert(execution_request_payload.fetch("host_root_modified") == false, "execution request preview must not mutate the host root")
  assert(execution_request_payload.fetch("network_required") == false, "execution request preview must not require network access")
  assert(execution_request_payload.fetch("backend_details_exposed") == false, "execution request preview must not expose backend details")
  assert(execution_request_payload.fetch("blocked_actions").include?("persist execution request before Runtime gates pass"), "execution request preview must block request persistence")
  assert(execution_request_payload.fetch("blocked_actions").include?("start compatibility profile from execution request preview"), "execution request preview must block compatibility launch")
  assert(!execution_request.downcase.include?("prefix"), "execution request preview must not expose implementation storage")
  assert(!execution_request.include?(".exe"), "execution request preview must not expose a Windows executable")
  assert(!execution_request.downcase.include?("program files"), "execution request preview must not expose Windows paths")
  assert(!execution_request.downcase.include?("qemu-system"), "execution request preview must not expose VM implementation commands")
  assert(!execution_request.downcase.include?("proton"), "execution request preview must not expose backend implementation names")
  assert(!execution_request.downcase.include?("wine "), "execution request preview must not expose backend implementation names")
  assert(!execution_request.downcase.include?("virtual machine"), "execution request preview must not expose implementation labels")

  execution_review, execution_review_status = capture_runtime_go(
    project_root,
    runtime_go_binary,
    go_binary,
    "execution-review-preview",
    "--registry",
    "runtime/recipes/registry.json",
    "--app",
    "org.xnix.sample.notepad",
    "file:///home/test/Documents/example.txt"
  )
  assert(execution_review_status.success?, "Go execution review preview CLI must run successfully")
  execution_review_payload = JSON.parse(execution_review)
  assert(execution_review_payload.fetch("schema_version") == "xnix.runtime.request_review.v1", "execution review preview schema version must be stable")
  assert(execution_review_payload.fetch("request_type") == "execution-review-preview", "execution review preview must identify its request type")
  assert(execution_review_payload.fetch("review_type") == "compatibility-center-launch-review", "execution review preview must identify the review type")
  assert(execution_review_payload.fetch("request_state") == "blocked", "execution review preview must remain blocked")
  assert(execution_review_payload.fetch("source") == "execution-request-preview", "execution review preview must derive from request preview")
  assert(execution_review_payload.fetch("desktop") == "KDE Plasma", "execution review preview must target KDE Plasma")
  assert(execution_review_payload.fetch("runtime_method") == "Launch", "execution review preview must target the Runtime Launch method")
  assert(execution_review_payload.fetch("read_method") == "GetExecutionReviewPreview", "execution review preview must expose the Runtime read method")
  assert(execution_review_payload.fetch("application_id") == recipe.id, "execution review preview must preserve application identity")
  assert(execution_review_payload.fetch("application_name") == recipe.name, "execution review preview must preserve display names")
  assert(execution_review_payload.fetch("desktop_file") == "xnix-#{recipe.id}.desktop", "execution review preview must bind generated desktop files")
  review_request = execution_review_payload.fetch("execution_request")
  assert(review_request.fetch("schema_version") == "xnix.runtime.request_intake.v1", "execution review preview must summarize request intake")
  assert(review_request.fetch("request_type") == "execution-request-preview", "execution review preview must keep request intake type")
  assert(review_request.fetch("request_state") == "blocked", "execution review preview request summary must remain blocked")
  assert(review_request.fetch("source") == "runtime-launch-intent", "execution review preview request summary must derive from launch intent")
  assert(review_request.fetch("read_method") == "GetExecutionRequestPreview", "execution review preview request summary must point to request read method")
  assert(review_request.fetch("portal_required") == true, "execution review preview request summary must require Portal review for files")
  assert(review_request.fetch("snapshot_required") == true, "execution review preview request summary must require snapshot review")
  assert(review_request.fetch("file_count") == 1, "execution review preview request summary must count selected files")
  assert(review_request.fetch("launch_intent_captured") == true, "execution review preview request summary must capture launch intent")
  assert(review_request.fetch("execution_request_created") == false, "execution review preview request summary must not create execution requests")
  assert(review_request.fetch("execution_request_persisted") == false, "execution review preview request summary must not persist execution requests")
  assert(review_request.fetch("execution_started") == false, "execution review preview request summary must not start execution")
  assert(review_request.fetch("backend_details_exposed") == false, "execution review preview request summary must hide backend details")
  review_card = execution_review_payload.fetch("review_card")
  assert(review_card.fetch("id") == "#{recipe.id}:launch-review", "execution review preview must provide a stable review card id")
  assert(review_card.fetch("card_type") == "compatibility-center-launch-review", "execution review preview must provide a launch review card")
  assert(review_card.fetch("status") == "blocked", "execution review card must remain blocked")
  assert(review_card.fetch("severity") == "requires-runtime-gates", "execution review card must explain gate severity")
  assert(review_card.fetch("user_review_required") == true, "execution review card must require user review")
  assert(review_card.fetch("runtime_approval_needed") == true, "execution review card must require Runtime approval")
  assert(review_card.fetch("primary_action") == "Open Compatibility Center", "execution review card must navigate to the Compatibility Center")
  assert(review_card.fetch("secondary_actions").include?("Review resource access"), "execution review card must include resource review")
  review_queue = execution_review_payload.fetch("action_queue")
  assert(review_queue.fetch("queue_type") == "compatibility-center-request-review-queue", "execution review preview must expose request review queue semantics")
  assert(review_queue.fetch("action_count") == 1, "execution review queue must contain one review action")
  assert(review_queue.fetch("pending_action_count") == 1, "execution review queue must mark the action pending")
  assert(review_queue.fetch("user_review_required_count") == 1, "execution review queue must require user review")
  assert(review_queue.fetch("execution_enabled") == false, "execution review queue must not enable execution")
  assert(review_queue.fetch("queue_persisted") == false, "execution review queue must not be persisted in preview mode")
  assert(review_queue.fetch("review_receipt_recorded") == false, "execution review queue must not record review receipts in preview mode")
  assert(review_queue.fetch("actions") == ["review-launch-request"], "execution review queue must expose the launch review action")
  review_profile = execution_review_payload.fetch("compatibility_profile")
  assert(review_profile.fetch("id") == "local-compatibility", "execution review preview must use the recommended profile")
  assert(review_profile.fetch("ready") == false, "execution review preview must not mark the profile ready")
  assert(review_profile.fetch("launch_enabled") == false, "execution review preview must not enable profile launch")
  assert(review_profile.fetch("backend_details_exposed") == false, "execution review preview must hide backend details")
  review_gates = execution_review_payload.fetch("gate_summary")
  assert(review_gates.fetch("gate_count") == 5, "execution review preview must summarize Runtime gates")
  assert(review_gates.fetch("required_gate_count") == 2, "execution review preview must count required gates")
  assert(review_gates.fetch("pending_gate_count") == 1, "execution review preview must count pending gates")
  assert(review_gates.fetch("blocked_gate_count") == 1, "execution review preview must count blocked gates")
  assert(execution_review_payload.fetch("execution_state") == "blocked", "execution review preview must stay blocked before Runtime gates pass")
  assert(execution_review_payload.fetch("overall_status") == "not-ready", "execution review preview must not claim launch readiness")
  assert(execution_review_payload.fetch("write_gate_decision") == "blocked-until-production-backend", "execution review preview must expose the Launch write gate")
  assert(execution_review_payload.fetch("denial_error_name") == "org.xnix.Compatibility1.Error.WriteMethodDisabled", "execution review preview must expose the write-gate error")
  assert(execution_review_payload.fetch("portal_required") == true, "execution review preview with files must require Portal review")
  assert(execution_review_payload.fetch("snapshot_required") == true, "execution review preview must require snapshot review")
  assert(execution_review_payload.fetch("file_count") == 1, "execution review preview must count selected files")
  assert(execution_review_payload.fetch("file_uris") == ["file:///home/test/Documents/example.txt"], "execution review preview must preserve file URIs")
  assert(execution_review_payload.fetch("runtime_owned") == true, "execution review preview must remain Runtime-owned")
  assert(execution_review_payload.fetch("go_runtime_backed") == true, "execution review preview must be Go Runtime-backed")
  assert(execution_review_payload.fetch("kde_policy_owner") == false, "execution review preview must not make KDE own backend policy")
  assert(execution_review_payload.fetch("compatibility_center_card") == true, "execution review preview must be suitable for Compatibility Center cards")
  assert(execution_review_payload.fetch("safe_for_ai_diagnostics") == true, "execution review preview must be safe for AI diagnostics")
  assert(execution_review_payload.fetch("desktop_entry_launch_visible") == true, "execution review preview must allow KDE to show launcher entries")
  assert(execution_review_payload.fetch("launch_intent_captured") == true, "execution review preview must capture launch intent")
  assert(execution_review_payload.fetch("action_queue_candidate") == true, "execution review preview must be a queue candidate")
  assert(execution_review_payload.fetch("review_receipt_required") == true, "execution review preview must require a later review receipt")
  assert(execution_review_payload.fetch("review_receipt_recorded") == false, "execution review preview must not record review receipts")
  assert(execution_review_payload.fetch("launch_allowed") == false, "execution review preview must not allow launch")
  assert(execution_review_payload.fetch("launch_enabled") == false, "execution review preview must not enable launch")
  assert(execution_review_payload.fetch("execution_request_created") == false, "execution review preview must not create execution requests")
  assert(execution_review_payload.fetch("execution_request_persisted") == false, "execution review preview must not persist execution requests")
  assert(execution_review_payload.fetch("action_queue_persisted") == false, "execution review preview must not persist action queues")
  assert(execution_review_payload.fetch("execution_started") == false, "execution review preview must not start execution")
  assert(execution_review_payload.fetch("backend_binding_ready") == false, "execution review preview must not claim backend binding readiness")
  assert(execution_review_payload.fetch("request_object_created") == false, "execution review preview must not create request objects")
  assert(execution_review_payload.fetch("permission_granted") == false, "execution review preview must not grant permissions")
  assert(execution_review_payload.fetch("host_root_modified") == false, "execution review preview must not mutate the host root")
  assert(execution_review_payload.fetch("network_required") == false, "execution review preview must not require network access")
  assert(execution_review_payload.fetch("backend_details_exposed") == false, "execution review preview must not expose backend details")
  assert(execution_review_payload.fetch("blocked_actions").include?("persist launch review before Runtime gates pass"), "execution review preview must block review persistence")
  assert(execution_review_payload.fetch("blocked_actions").include?("record launch approval from preview state"), "execution review preview must block approval recording")
  assert(!execution_review.downcase.include?("prefix"), "execution review preview must not expose implementation storage")
  assert(!execution_review.include?(".exe"), "execution review preview must not expose a Windows executable")
  assert(!execution_review.downcase.include?("program files"), "execution review preview must not expose Windows paths")
  assert(!execution_review.downcase.include?("qemu-system"), "execution review preview must not expose VM implementation commands")
  assert(!execution_review.downcase.include?("proton"), "execution review preview must not expose backend implementation names")
  assert(!execution_review.downcase.include?("wine "), "execution review preview must not expose backend implementation names")
  assert(!execution_review.downcase.include?("virtual machine"), "execution review preview must not expose implementation labels")

  execution_decision, execution_decision_status = capture_runtime_go(
    project_root,
    runtime_go_binary,
    go_binary,
    "execution-decision-preview",
    "--registry",
    "runtime/recipes/registry.json",
    "--app",
    "org.xnix.sample.notepad",
    "--decision",
    "approved",
    "file:///home/test/Documents/example.txt"
  )
  assert(execution_decision_status.success?, "Go execution decision preview CLI must run successfully")
  execution_decision_payload = JSON.parse(execution_decision)
  assert(execution_decision_payload.fetch("schema_version") == "xnix.runtime.request_decision.v1", "execution decision preview schema version must be stable")
  assert(execution_decision_payload.fetch("request_type") == "execution-decision-preview", "execution decision preview must identify its request type")
  assert(execution_decision_payload.fetch("decision_type") == "compatibility-center-launch-decision", "execution decision preview must identify the decision type")
  assert(execution_decision_payload.fetch("request_state") == "blocked", "execution decision preview must remain blocked")
  assert(execution_decision_payload.fetch("source") == "execution-review-preview", "execution decision preview must derive from review preview")
  assert(execution_decision_payload.fetch("desktop") == "KDE Plasma", "execution decision preview must target KDE Plasma")
  assert(execution_decision_payload.fetch("runtime_method") == "Launch", "execution decision preview must target the Runtime Launch method")
  assert(execution_decision_payload.fetch("read_method") == "GetExecutionDecisionPreview", "execution decision preview must expose the Runtime read method")
  assert(execution_decision_payload.fetch("application_id") == recipe.id, "execution decision preview must preserve application identity")
  assert(execution_decision_payload.fetch("application_name") == recipe.name, "execution decision preview must preserve display names")
  assert(execution_decision_payload.fetch("desktop_file") == "xnix-#{recipe.id}.desktop", "execution decision preview must bind generated desktop files")
  decision_review = execution_decision_payload.fetch("execution_review")
  assert(decision_review.fetch("schema_version") == "xnix.runtime.request_review.v1", "execution decision preview must summarize review state")
  assert(decision_review.fetch("request_type") == "execution-review-preview", "execution decision preview must keep review type")
  assert(decision_review.fetch("review_type") == "compatibility-center-launch-review", "execution decision preview must keep launch review type")
  assert(decision_review.fetch("source") == "execution-request-preview", "execution decision preview must derive from request review")
  assert(decision_review.fetch("read_method") == "GetExecutionReviewPreview", "execution decision preview must point to review read method")
  assert(decision_review.fetch("action_queue_candidate") == true, "execution decision preview must preserve queue candidate state")
  assert(decision_review.fetch("review_receipt_required") == true, "execution decision preview must require later review receipt")
  assert(decision_review.fetch("review_receipt_recorded") == false, "execution decision preview must not record review receipts")
  assert(decision_review.fetch("action_queue_persisted") == false, "execution decision preview must not persist action queues")
  assert(decision_review.fetch("execution_started") == false, "execution decision preview must not start execution")
  assert(decision_review.fetch("backend_details_exposed") == false, "execution decision preview review summary must hide backend details")
  decision_payload = execution_decision_payload.fetch("decision")
  assert(decision_payload.fetch("decision") == "approved", "execution decision preview must preserve the requested decision")
  assert(decision_payload.fetch("decision_accepted") == true, "execution decision preview must accept supported decisions")
  assert(decision_payload.fetch("user_intent_captured") == true, "execution decision preview must capture user intent")
  assert(decision_payload.fetch("decision_recorded") == false, "execution decision preview must not record decisions")
  assert(decision_payload.fetch("runtime_approval_granted") == false, "execution decision preview must not grant Runtime approval")
  assert(decision_payload.fetch("execution_allowed") == false, "execution decision preview must not allow execution")
  assert(decision_payload.fetch("review_receipt_created") == false, "execution decision preview must not create review receipts")
  assert(decision_payload.fetch("queue_state_changed") == false, "execution decision preview must not mutate queue state")
  assert(decision_payload.fetch("permission_grant_created") == false, "execution decision preview must not grant permissions")
  assert(decision_payload.fetch("desktop_notification_intent") == "show-launch-approval-pending", "execution decision preview must expose a safe notification intent")
  decision_queue = execution_decision_payload.fetch("action_queue")
  assert(decision_queue.fetch("queue_type") == "compatibility-center-request-review-queue", "execution decision preview must keep queue semantics")
  assert(decision_queue.fetch("execution_enabled") == false, "execution decision preview queue must not enable execution")
  assert(decision_queue.fetch("queue_persisted") == false, "execution decision preview queue must not be persisted")
  assert(decision_queue.fetch("review_receipt_recorded") == false, "execution decision preview queue must not record receipts")
  assert(execution_decision_payload.fetch("portal_required") == true, "execution decision preview with files must require Portal review")
  assert(execution_decision_payload.fetch("snapshot_required") == true, "execution decision preview must require snapshot review")
  assert(execution_decision_payload.fetch("file_count") == 1, "execution decision preview must count selected files")
  assert(execution_decision_payload.fetch("file_uris") == ["file:///home/test/Documents/example.txt"], "execution decision preview must preserve file URIs")
  assert(execution_decision_payload.fetch("runtime_owned") == true, "execution decision preview must remain Runtime-owned")
  assert(execution_decision_payload.fetch("go_runtime_backed") == true, "execution decision preview must be Go Runtime-backed")
  assert(execution_decision_payload.fetch("kde_policy_owner") == false, "execution decision preview must not make KDE own backend policy")
  assert(execution_decision_payload.fetch("compatibility_center_card") == true, "execution decision preview must be suitable for Compatibility Center cards")
  assert(execution_decision_payload.fetch("safe_for_ai_diagnostics") == true, "execution decision preview must be safe for AI diagnostics")
  assert(execution_decision_payload.fetch("desktop_entry_launch_visible") == true, "execution decision preview must allow KDE to show launcher entries")
  assert(execution_decision_payload.fetch("launch_intent_captured") == true, "execution decision preview must capture launch intent")
  assert(execution_decision_payload.fetch("action_queue_candidate") == true, "execution decision preview must be a queue candidate")
  assert(execution_decision_payload.fetch("user_decision_captured") == true, "execution decision preview must capture the user decision")
  assert(execution_decision_payload.fetch("review_receipt_required") == true, "execution decision preview must require a later review receipt")
  assert(execution_decision_payload.fetch("review_receipt_recorded") == false, "execution decision preview must not record review receipts")
  assert(execution_decision_payload.fetch("runtime_launch_approval") == false, "execution decision preview must not grant Runtime launch approval")
  assert(execution_decision_payload.fetch("launch_allowed") == false, "execution decision preview must not allow launch")
  assert(execution_decision_payload.fetch("launch_enabled") == false, "execution decision preview must not enable launch")
  assert(execution_decision_payload.fetch("execution_request_created") == false, "execution decision preview must not create execution requests")
  assert(execution_decision_payload.fetch("execution_request_persisted") == false, "execution decision preview must not persist execution requests")
  assert(execution_decision_payload.fetch("action_queue_persisted") == false, "execution decision preview must not persist action queues")
  assert(execution_decision_payload.fetch("execution_started") == false, "execution decision preview must not start execution")
  assert(execution_decision_payload.fetch("backend_binding_ready") == false, "execution decision preview must not claim backend binding readiness")
  assert(execution_decision_payload.fetch("request_object_created") == false, "execution decision preview must not create request objects")
  assert(execution_decision_payload.fetch("permission_granted") == false, "execution decision preview must not grant permissions")
  assert(execution_decision_payload.fetch("host_root_modified") == false, "execution decision preview must not mutate the host root")
  assert(execution_decision_payload.fetch("network_required") == false, "execution decision preview must not require network access")
  assert(execution_decision_payload.fetch("backend_details_exposed") == false, "execution decision preview must not expose backend details")
  assert(execution_decision_payload.fetch("blocked_actions").include?("record launch decision from preview state"), "execution decision preview must block decision recording")
  assert(execution_decision_payload.fetch("blocked_actions").include?("treat user decision as Runtime launch approval"), "execution decision preview must block approval escalation")
  assert(!execution_decision.downcase.include?("prefix"), "execution decision preview must not expose implementation storage")
  assert(!execution_decision.include?(".exe"), "execution decision preview must not expose a Windows executable")
  assert(!execution_decision.downcase.include?("program files"), "execution decision preview must not expose Windows paths")
  assert(!execution_decision.downcase.include?("qemu-system"), "execution decision preview must not expose VM implementation commands")
  assert(!execution_decision.downcase.include?("proton"), "execution decision preview must not expose backend implementation names")
  assert(!execution_decision.downcase.include?("wine "), "execution decision preview must not expose backend implementation names")
  assert(!execution_decision.downcase.include?("virtual machine"), "execution decision preview must not expose implementation labels")

  execution_preflight, execution_preflight_status = capture_runtime_go(
    project_root,
    runtime_go_binary,
    go_binary,
    "execution-preflight-preview",
    "--registry",
    "runtime/recipes/registry.json",
    "--app",
    "org.xnix.sample.notepad",
    "--decision",
    "approved",
    "file:///home/test/Documents/example.txt"
  )
  assert(execution_preflight_status.success?, "Go execution preflight preview CLI must run successfully")
  execution_preflight_payload = JSON.parse(execution_preflight)
  assert(execution_preflight_payload.fetch("schema_version") == "xnix.runtime.launch_preflight.v1", "execution preflight preview schema version must be stable")
  assert(execution_preflight_payload.fetch("request_type") == "execution-preflight-preview", "execution preflight preview must identify its request type")
  assert(execution_preflight_payload.fetch("preflight_type") == "compatibility-launch-preflight", "execution preflight preview must identify launch preflight")
  assert(execution_preflight_payload.fetch("request_state") == "blocked", "execution preflight preview must remain blocked")
  assert(execution_preflight_payload.fetch("source") == "execution-decision-preview", "execution preflight preview must derive from decision preview")
  assert(execution_preflight_payload.fetch("desktop") == "KDE Plasma", "execution preflight preview must target KDE Plasma")
  assert(execution_preflight_payload.fetch("runtime_method") == "Launch", "execution preflight preview must target the Runtime Launch method")
  assert(execution_preflight_payload.fetch("read_method") == "GetExecutionPreflightPreview", "execution preflight preview must expose the Runtime read method")
  assert(execution_preflight_payload.fetch("application_id") == recipe.id, "execution preflight preview must preserve application identity")
  assert(execution_preflight_payload.fetch("application_name") == recipe.name, "execution preflight preview must preserve display names")
  assert(execution_preflight_payload.fetch("desktop_file") == "xnix-#{recipe.id}.desktop", "execution preflight preview must bind generated desktop files")
  preflight_decision = execution_preflight_payload.fetch("execution_decision")
  assert(preflight_decision.fetch("schema_version") == "xnix.runtime.request_decision.v1", "execution preflight preview must summarize decision state")
  assert(preflight_decision.fetch("request_type") == "execution-decision-preview", "execution preflight preview must keep decision type")
  assert(preflight_decision.fetch("decision") == "approved", "execution preflight preview must preserve requested decision")
  assert(preflight_decision.fetch("decision_accepted") == true, "execution preflight preview must accept supported decisions")
  assert(preflight_decision.fetch("user_intent_captured") == true, "execution preflight preview must capture user intent")
  assert(preflight_decision.fetch("decision_recorded") == false, "execution preflight preview must not record decisions")
  assert(preflight_decision.fetch("runtime_approval_granted") == false, "execution preflight preview must not grant Runtime approval")
  assert(preflight_decision.fetch("execution_allowed") == false, "execution preflight preview must not allow execution")
  assert(preflight_decision.fetch("review_receipt_created") == false, "execution preflight preview must not create receipts")
  assert(preflight_decision.fetch("queue_state_changed") == false, "execution preflight preview must not change queue state")
  assert(execution_preflight_payload.fetch("check_count") == 5, "execution preflight preview must expose five preflight checks")
  assert(execution_preflight_payload.fetch("passed_check_count") == 1, "execution preflight preview must count passed checks")
  assert(execution_preflight_payload.fetch("required_check_count") == 2, "execution preflight preview must count required checks")
  assert(execution_preflight_payload.fetch("pending_check_count") == 1, "execution preflight preview must count pending checks")
  assert(execution_preflight_payload.fetch("blocked_check_count") == 1, "execution preflight preview must count blocked checks")
  preflight_checks = execution_preflight_payload.fetch("preflight_checks")
  assert(preflight_checks.map { |check| check.fetch("id") } == %w[user-decision portal-policy-review snapshot-baseline backend-binding runtime-launch-write-gate], "execution preflight preview must preserve check order")
  assert(preflight_checks.first.fetch("status") == "pass", "execution preflight preview must pass approved user decision")
  assert(preflight_checks[1].fetch("status") == "required", "execution preflight preview must require Portal review")
  assert(preflight_checks[3].fetch("status") == "pending", "execution preflight preview must keep backend binding pending")
  assert(preflight_checks[4].fetch("status") == "blocked", "execution preflight preview must keep Runtime write gate blocked")
  assert(execution_preflight_payload.fetch("execution_state") == "blocked", "execution preflight preview must stay blocked before Runtime gates pass")
  assert(execution_preflight_payload.fetch("overall_status") == "not-ready", "execution preflight preview must not claim launch readiness")
  assert(execution_preflight_payload.fetch("write_gate_decision") == "blocked-until-production-backend", "execution preflight preview must expose the Launch write gate")
  assert(execution_preflight_payload.fetch("denial_error_name") == "org.xnix.Compatibility1.Error.WriteMethodDisabled", "execution preflight preview must expose the write-gate error")
  assert(execution_preflight_payload.fetch("portal_required") == true, "execution preflight preview with files must require Portal review")
  assert(execution_preflight_payload.fetch("snapshot_required") == true, "execution preflight preview must require snapshot review")
  assert(execution_preflight_payload.fetch("file_count") == 1, "execution preflight preview must count selected files")
  assert(execution_preflight_payload.fetch("file_uris") == ["file:///home/test/Documents/example.txt"], "execution preflight preview must preserve file URIs")
  assert(execution_preflight_payload.fetch("runtime_owned") == true, "execution preflight preview must remain Runtime-owned")
  assert(execution_preflight_payload.fetch("go_runtime_backed") == true, "execution preflight preview must be Go Runtime-backed")
  assert(execution_preflight_payload.fetch("kde_policy_owner") == false, "execution preflight preview must not make KDE own backend policy")
  assert(execution_preflight_payload.fetch("compatibility_center_card") == true, "execution preflight preview must be suitable for Compatibility Center cards")
  assert(execution_preflight_payload.fetch("safe_for_ai_diagnostics") == true, "execution preflight preview must be safe for AI diagnostics")
  assert(execution_preflight_payload.fetch("desktop_entry_launch_visible") == true, "execution preflight preview must allow KDE to show launcher entries")
  assert(execution_preflight_payload.fetch("launch_intent_captured") == true, "execution preflight preview must capture launch intent")
  assert(execution_preflight_payload.fetch("user_decision_captured") == true, "execution preflight preview must capture user decision")
  assert(execution_preflight_payload.fetch("user_decision_allows_launch") == true, "execution preflight preview must model approved user intent")
  assert(execution_preflight_payload.fetch("preflight_complete") == false, "execution preflight preview must not complete preflight")
  assert(execution_preflight_payload.fetch("preflight_passed") == false, "execution preflight preview must not pass preflight")
  assert(execution_preflight_payload.fetch("portal_preflight_ready") == false, "execution preflight preview must not mark Portal ready")
  assert(execution_preflight_payload.fetch("snapshot_preflight_ready") == false, "execution preflight preview must not mark snapshot ready")
  assert(execution_preflight_payload.fetch("backend_preflight_ready") == false, "execution preflight preview must not mark backend binding ready")
  assert(execution_preflight_payload.fetch("write_gate_open") == false, "execution preflight preview must not open the Runtime write gate")
  assert(execution_preflight_payload.fetch("runtime_launch_approval") == false, "execution preflight preview must not grant Runtime launch approval")
  assert(execution_preflight_payload.fetch("launch_allowed") == false, "execution preflight preview must not allow launch")
  assert(execution_preflight_payload.fetch("launch_enabled") == false, "execution preflight preview must not enable launch")
  assert(execution_preflight_payload.fetch("execution_request_created") == false, "execution preflight preview must not create execution requests")
  assert(execution_preflight_payload.fetch("execution_request_persisted") == false, "execution preflight preview must not persist execution requests")
  assert(execution_preflight_payload.fetch("action_queue_persisted") == false, "execution preflight preview must not persist action queues")
  assert(execution_preflight_payload.fetch("review_receipt_recorded") == false, "execution preflight preview must not record review receipts")
  assert(execution_preflight_payload.fetch("execution_started") == false, "execution preflight preview must not start execution")
  assert(execution_preflight_payload.fetch("backend_binding_ready") == false, "execution preflight preview must not claim backend binding readiness")
  assert(execution_preflight_payload.fetch("request_object_created") == false, "execution preflight preview must not create request objects")
  assert(execution_preflight_payload.fetch("permission_granted") == false, "execution preflight preview must not grant permissions")
  assert(execution_preflight_payload.fetch("host_root_modified") == false, "execution preflight preview must not mutate the host root")
  assert(execution_preflight_payload.fetch("network_required") == false, "execution preflight preview must not require network access")
  assert(execution_preflight_payload.fetch("backend_details_exposed") == false, "execution preflight preview must not expose backend details")
  assert(execution_preflight_payload.fetch("blocked_actions").include?("start compatibility profile before preflight passes"), "execution preflight preview must block launch before preflight")
  assert(execution_preflight_payload.fetch("blocked_actions").include?("grant desktop resources before Portal review"), "execution preflight preview must block desktop resource grants")
  assert(!execution_preflight.downcase.include?("prefix"), "execution preflight preview must not expose implementation storage")
  assert(!execution_preflight.include?(".exe"), "execution preflight preview must not expose a Windows executable")
  assert(!execution_preflight.downcase.include?("program files"), "execution preflight preview must not expose Windows paths")
  assert(!execution_preflight.downcase.include?("qemu-system"), "execution preflight preview must not expose VM implementation commands")
  assert(!execution_preflight.downcase.include?("proton"), "execution preflight preview must not expose backend implementation names")
  assert(!execution_preflight.downcase.include?("wine "), "execution preflight preview must not expose backend implementation names")
  assert(!execution_preflight.downcase.include?("virtual machine"), "execution preflight preview must not expose implementation labels")

  execution_resource_grant, execution_resource_grant_status = capture_runtime_go(
    project_root,
    runtime_go_binary,
    go_binary,
    "execution-resource-grant-preview",
    "--registry",
    "runtime/recipes/registry.json",
    "--app",
    "org.xnix.sample.notepad",
    "--decision",
    "approved",
    "file:///home/test/Documents/example.txt"
  )
  assert(execution_resource_grant_status.success?, "Go execution resource grant preview CLI must run successfully")
  execution_resource_grant_payload = JSON.parse(execution_resource_grant)
  assert(execution_resource_grant_payload.fetch("schema_version") == "xnix.runtime.resource_grant.v1", "execution resource grant preview schema version must be stable")
  assert(execution_resource_grant_payload.fetch("request_type") == "execution-resource-grant-preview", "execution resource grant preview must identify its request type")
  assert(execution_resource_grant_payload.fetch("grant_type") == "compatibility-launch-resource-grant", "execution resource grant preview must identify launch resource grants")
  assert(execution_resource_grant_payload.fetch("request_state") == "blocked", "execution resource grant preview must remain blocked")
  assert(execution_resource_grant_payload.fetch("source") == "execution-preflight-preview", "execution resource grant preview must derive from preflight")
  assert(execution_resource_grant_payload.fetch("desktop") == "KDE Plasma", "execution resource grant preview must target KDE Plasma")
  assert(execution_resource_grant_payload.fetch("runtime_method") == "Launch", "execution resource grant preview must target the Runtime Launch method")
  assert(execution_resource_grant_payload.fetch("read_method") == "GetExecutionResourceGrantPreview", "execution resource grant preview must expose the Runtime read method")
  assert(execution_resource_grant_payload.fetch("application_id") == recipe.id, "execution resource grant preview must preserve application identity")
  assert(execution_resource_grant_payload.fetch("application_name") == recipe.name, "execution resource grant preview must preserve display names")
  assert(execution_resource_grant_payload.fetch("desktop_file") == "xnix-#{recipe.id}.desktop", "execution resource grant preview must bind generated desktop files")
  resource_preflight = execution_resource_grant_payload.fetch("execution_preflight")
  assert(resource_preflight.fetch("schema_version") == "xnix.runtime.launch_preflight.v1", "execution resource grant preview must summarize preflight schema")
  assert(resource_preflight.fetch("request_type") == "execution-preflight-preview", "execution resource grant preview must derive from preflight")
  assert(resource_preflight.fetch("user_decision_allows_launch") == true, "execution resource grant preview must preserve approved user intent")
  assert(resource_preflight.fetch("preflight_passed") == false, "execution resource grant preview must not pass preflight")
  assert(resource_preflight.fetch("portal_required") == true, "execution resource grant preview must require Portal review")
  assert(resource_preflight.fetch("write_gate_open") == false, "execution resource grant preview must not open the Runtime write gate")
  assert(resource_preflight.fetch("runtime_launch_approval") == false, "execution resource grant preview must not grant Runtime approval")
  assert(resource_preflight.fetch("permission_granted") == false, "execution resource grant preview must not grant permissions through preflight")
  assert(execution_resource_grant_payload.fetch("grant_count") == 7, "execution resource grant preview must expose every user-facing resource")
  assert(execution_resource_grant_payload.fetch("allow_count") == 1, "execution resource grant preview must count planned allow resources")
  assert(execution_resource_grant_payload.fetch("ask_count") == 5, "execution resource grant preview must count review-required resources")
  assert(execution_resource_grant_payload.fetch("deny_count") == 1, "execution resource grant preview must count denied resources")
  assert(execution_resource_grant_payload.fetch("portal_required_count") == 6, "execution resource grant preview must count Portal-mediated resources")
  assert(execution_resource_grant_payload.fetch("pending_review_count") == 5, "execution resource grant preview must count pending user reviews")
  grants_by_id = execution_resource_grant_payload.fetch("resource_grants").to_h { |grant| [grant.fetch("id"), grant] }
  assert(grants_by_id.fetch("documents").fetch("grant_state") == "requires-review", "execution resource grant preview must require Documents review")
  assert(grants_by_id.fetch("downloads").fetch("grant_state") == "requires-review", "execution resource grant preview must require Downloads review")
  assert(grants_by_id.fetch("camera").fetch("grant_state") == "planned-deny", "execution resource grant preview must keep Camera denied")
  assert(grants_by_id.fetch("network").fetch("grant_state") == "planned-allow", "execution resource grant preview must model Network as planned allow")
  assert(grants_by_id.fetch("network").fetch("portal_required") == false, "execution resource grant preview must not invent a Portal for Network")
  assert(execution_resource_grant_payload.fetch("resource_grants").all? { |grant| grant.fetch("request_object_created") == false }, "execution resource grant preview must not create request objects")
  assert(execution_resource_grant_payload.fetch("resource_grants").all? { |grant| grant.fetch("permission_granted") == false }, "execution resource grant preview must not grant individual permissions")
  assert(execution_resource_grant_payload.fetch("resource_grants").all? { |grant| grant.fetch("direct_access_allowed") == false }, "execution resource grant preview must not allow direct access")
  resource_bridge = execution_resource_grant_payload.fetch("bridge_summary")
  assert(resource_bridge.fetch("request_type") == "desktop-resource-bridge-preview", "execution resource grant preview must summarize resource bridge planning")
  assert(resource_bridge.fetch("portal_mediated") == true, "execution resource grant preview must keep resource bridges Portal-mediated")
  assert(resource_bridge.fetch("resource_bridges_enabled") == false, "execution resource grant preview must not enable resource bridges")
  assert(resource_bridge.fetch("request_objects_created") == false, "execution resource grant preview must not create bridge request objects")
  assert(execution_resource_grant_payload.fetch("runtime_owned") == true, "execution resource grant preview must remain Runtime-owned")
  assert(execution_resource_grant_payload.fetch("go_runtime_backed") == true, "execution resource grant preview must be Go Runtime-backed")
  assert(execution_resource_grant_payload.fetch("kde_policy_owner") == false, "execution resource grant preview must not make KDE own backend policy")
  assert(execution_resource_grant_payload.fetch("compatibility_center_card") == true, "execution resource grant preview must be suitable for Compatibility Center cards")
  assert(execution_resource_grant_payload.fetch("safe_for_ai_diagnostics") == true, "execution resource grant preview must be safe for AI diagnostics")
  assert(execution_resource_grant_payload.fetch("desktop_entry_launch_visible") == true, "execution resource grant preview must allow KDE to show launcher entries")
  assert(execution_resource_grant_payload.fetch("launch_intent_captured") == true, "execution resource grant preview must capture launch intent")
  assert(execution_resource_grant_payload.fetch("user_decision_captured") == true, "execution resource grant preview must capture user decision")
  assert(execution_resource_grant_payload.fetch("user_decision_allows_launch") == true, "execution resource grant preview must model approved user intent")
  assert(execution_resource_grant_payload.fetch("preflight_passed") == false, "execution resource grant preview must not pass preflight")
  assert(execution_resource_grant_payload.fetch("portal_review_required") == true, "execution resource grant preview must require Portal review")
  assert(execution_resource_grant_payload.fetch("snapshot_required") == true, "execution resource grant preview must require snapshot review")
  assert(execution_resource_grant_payload.fetch("grant_plan_created") == true, "execution resource grant preview must create only a read model")
  assert(execution_resource_grant_payload.fetch("grant_objects_created") == false, "execution resource grant preview must not create grant objects")
  assert(execution_resource_grant_payload.fetch("request_objects_created") == false, "execution resource grant preview must not create request objects")
  assert(execution_resource_grant_payload.fetch("permission_granted") == false, "execution resource grant preview must not grant permissions")
  assert(execution_resource_grant_payload.fetch("resource_bridges_enabled") == false, "execution resource grant preview must not enable resource bridges")
  assert(execution_resource_grant_payload.fetch("settings_persisted") == false, "execution resource grant preview must not persist settings")
  assert(execution_resource_grant_payload.fetch("runtime_launch_approval") == false, "execution resource grant preview must not grant Runtime approval")
  assert(execution_resource_grant_payload.fetch("launch_allowed") == false, "execution resource grant preview must not allow launch")
  assert(execution_resource_grant_payload.fetch("launch_enabled") == false, "execution resource grant preview must not enable launch")
  assert(execution_resource_grant_payload.fetch("execution_started") == false, "execution resource grant preview must not start execution")
  assert(execution_resource_grant_payload.fetch("host_permission_changed") == false, "execution resource grant preview must not change host permissions")
  assert(execution_resource_grant_payload.fetch("host_root_modified") == false, "execution resource grant preview must not mutate the host root")
  assert(execution_resource_grant_payload.fetch("network_required") == false, "execution resource grant preview must not require network access")
  assert(execution_resource_grant_payload.fetch("backend_details_exposed") == false, "execution resource grant preview must not expose backend details")
  assert(execution_resource_grant_payload.fetch("blocked_actions").include?("create resource grant objects from preview"), "execution resource grant preview must block grant object creation")
  assert(execution_resource_grant_payload.fetch("blocked_actions").include?("enable desktop resource bridges before Portal review"), "execution resource grant preview must block resource bridge activation")
  assert(!execution_resource_grant.downcase.include?("prefix"), "execution resource grant preview must not expose implementation storage")
  assert(!execution_resource_grant.include?(".exe"), "execution resource grant preview must not expose a Windows executable")
  assert(!execution_resource_grant.downcase.include?("program files"), "execution resource grant preview must not expose Windows paths")
  assert(!execution_resource_grant.downcase.include?("qemu-system"), "execution resource grant preview must not expose VM implementation commands")
  assert(!execution_resource_grant.downcase.include?("proton"), "execution resource grant preview must not expose backend implementation names")
  assert(!execution_resource_grant.downcase.include?("wine "), "execution resource grant preview must not expose backend implementation names")
  assert(!execution_resource_grant.downcase.include?("virtual machine"), "execution resource grant preview must not expose implementation labels")

  execution_transaction, execution_transaction_status = capture_runtime_go(
    project_root,
    runtime_go_binary,
    go_binary,
    "execution-transaction-preview",
    "--registry",
    "runtime/recipes/registry.json",
    "--app",
    "org.xnix.sample.notepad",
    "--decision",
    "approved",
    "file:///home/test/Documents/example.txt"
  )
  assert(execution_transaction_status.success?, "Go execution transaction preview CLI must run successfully")
  execution_transaction_payload = JSON.parse(execution_transaction)
  assert(execution_transaction_payload.fetch("schema_version") == "xnix.runtime.launch_transaction.v1", "execution transaction preview schema version must be stable")
  assert(execution_transaction_payload.fetch("request_type") == "execution-transaction-preview", "execution transaction preview must identify its request type")
  assert(execution_transaction_payload.fetch("transaction_type") == "compatibility-launch-transaction", "execution transaction preview must identify launch transaction planning")
  assert(execution_transaction_payload.fetch("request_state") == "blocked", "execution transaction preview must remain blocked")
  assert(execution_transaction_payload.fetch("source") == "execution-resource-grant-preview", "execution transaction preview must derive from resource grant preview")
  assert(execution_transaction_payload.fetch("desktop") == "KDE Plasma", "execution transaction preview must target KDE Plasma")
  assert(execution_transaction_payload.fetch("runtime_method") == "Launch", "execution transaction preview must target the Runtime Launch method")
  assert(execution_transaction_payload.fetch("read_method") == "GetExecutionTransactionPreview", "execution transaction preview must expose the Runtime read method")
  assert(execution_transaction_payload.fetch("application_id") == recipe.id, "execution transaction preview must preserve application identity")
  assert(execution_transaction_payload.fetch("application_name") == recipe.name, "execution transaction preview must preserve display names")
  assert(execution_transaction_payload.fetch("desktop_file") == "xnix-#{recipe.id}.desktop", "execution transaction preview must bind generated desktop files")
  transaction_resource = execution_transaction_payload.fetch("resource_grant")
  assert(transaction_resource.fetch("request_type") == "execution-resource-grant-preview", "execution transaction preview must summarize resource grant planning")
  assert(transaction_resource.fetch("grant_count") == 7, "execution transaction preview must preserve resource grant count")
  assert(transaction_resource.fetch("portal_required_count") == 6, "execution transaction preview must preserve Portal resource count")
  assert(transaction_resource.fetch("grant_plan_created") == true, "execution transaction preview must depend on the read-only grant plan")
  assert(transaction_resource.fetch("grant_objects_created") == false, "execution transaction preview must not create grant objects")
  assert(transaction_resource.fetch("permission_granted") == false, "execution transaction preview must not grant permissions")
  transaction_readiness = execution_transaction_payload.fetch("readiness")
  assert(transaction_readiness.fetch("request_type") == "execution-readiness-preview", "execution transaction preview must summarize readiness")
  assert(transaction_readiness.fetch("execution_state") == "blocked", "execution transaction preview must keep execution blocked")
  assert(transaction_readiness.fetch("overall_status") == "not-ready", "execution transaction preview must not claim readiness")
  assert(transaction_readiness.fetch("launch_allowed") == false, "execution transaction preview readiness must not allow launch")
  assert(transaction_readiness.fetch("backend_binding_ready") == false, "execution transaction preview readiness must not claim backend binding")
  transaction_binding = execution_transaction_payload.fetch("backend_binding")
  assert(transaction_binding.fetch("recommended_profile_id") == "local-compatibility", "execution transaction preview must preserve recommended compatibility profile")
  assert(transaction_binding.fetch("candidate_count") == 2, "execution transaction preview must preserve compatibility profile count")
  assert(transaction_binding.fetch("binding_required") == true, "execution transaction preview must require managed compatibility binding")
  assert(transaction_binding.fetch("binding_committed") == false, "execution transaction preview must not commit compatibility binding")
  assert(transaction_binding.fetch("environment_created") == false, "execution transaction preview must not create environments")
  transaction_snapshot = execution_transaction_payload.fetch("snapshot_baseline")
  assert(transaction_snapshot.fetch("required") == true, "execution transaction preview must require a snapshot baseline")
  assert(transaction_snapshot.fetch("state") == "required", "execution transaction preview must keep snapshot baseline required")
  assert(transaction_snapshot.fetch("baseline_created") == false, "execution transaction preview must not create snapshot baselines")
  assert(transaction_snapshot.fetch("user_documents_included") == false, "execution transaction preview snapshot must exclude user documents")
  assert(transaction_snapshot.fetch("host_system_included") == false, "execution transaction preview snapshot must exclude the host system")
  transaction_write_gate = execution_transaction_payload.fetch("write_gate")
  assert(transaction_write_gate.fetch("method_name") == "Launch", "execution transaction preview must expose the Launch write gate")
  assert(transaction_write_gate.fetch("gate_decision") == "blocked-until-production-backend", "execution transaction preview must keep Launch write gate blocked")
  assert(transaction_write_gate.fetch("write_method_enabled") == false, "execution transaction preview must not enable write methods")
  assert(transaction_write_gate.fetch("dispatch_enabled") == false, "execution transaction preview must not enable dispatch")
  assert(execution_transaction_payload.fetch("step_count") == 7, "execution transaction preview must expose seven transaction steps")
  assert(execution_transaction_payload.fetch("passed_step_count") == 2, "execution transaction preview must count passed steps")
  assert(execution_transaction_payload.fetch("required_step_count") == 2, "execution transaction preview must count required steps")
  assert(execution_transaction_payload.fetch("pending_step_count") == 2, "execution transaction preview must count pending steps")
  assert(execution_transaction_payload.fetch("blocked_step_count") == 1, "execution transaction preview must count blocked steps")
  transaction_steps = execution_transaction_payload.fetch("transaction_steps")
  assert(transaction_steps.map { |step| step.fetch("id") } == %w[identity-validation user-decision portal-resource-review resource-grant-objects snapshot-baseline backend-binding runtime-launch-write-gate], "execution transaction preview must preserve step order")
  assert(transaction_steps.last.fetch("status") == "blocked", "execution transaction preview must keep the Launch write gate blocked")
  assert(execution_transaction_payload.fetch("runtime_owned") == true, "execution transaction preview must remain Runtime-owned")
  assert(execution_transaction_payload.fetch("go_runtime_backed") == true, "execution transaction preview must be Go Runtime-backed")
  assert(execution_transaction_payload.fetch("kde_policy_owner") == false, "execution transaction preview must not make KDE own backend policy")
  assert(execution_transaction_payload.fetch("compatibility_center_card") == true, "execution transaction preview must be suitable for Compatibility Center cards")
  assert(execution_transaction_payload.fetch("safe_for_ai_diagnostics") == true, "execution transaction preview must be safe for AI diagnostics")
  assert(execution_transaction_payload.fetch("desktop_entry_launch_visible") == true, "execution transaction preview must allow KDE to show launcher entries")
  assert(execution_transaction_payload.fetch("launch_intent_captured") == true, "execution transaction preview must capture launch intent")
  assert(execution_transaction_payload.fetch("user_decision_captured") == true, "execution transaction preview must capture user decision")
  assert(execution_transaction_payload.fetch("user_decision_allows_launch") == true, "execution transaction preview must model approved user intent")
  assert(execution_transaction_payload.fetch("transaction_plan_created") == true, "execution transaction preview must create only a read model")
  assert(execution_transaction_payload.fetch("transaction_committed") == false, "execution transaction preview must not commit transactions")
  assert(execution_transaction_payload.fetch("resource_grants_committed") == false, "execution transaction preview must not commit resource grants")
  assert(execution_transaction_payload.fetch("portal_approval_recorded") == false, "execution transaction preview must not record Portal approval")
  assert(execution_transaction_payload.fetch("snapshot_baseline_created") == false, "execution transaction preview must not create snapshot baselines")
  assert(execution_transaction_payload.fetch("backend_binding_committed") == false, "execution transaction preview must not commit compatibility binding")
  assert(execution_transaction_payload.fetch("runtime_launch_approval") == false, "execution transaction preview must not grant Runtime approval")
  assert(execution_transaction_payload.fetch("launch_allowed") == false, "execution transaction preview must not allow launch")
  assert(execution_transaction_payload.fetch("launch_enabled") == false, "execution transaction preview must not enable launch")
  assert(execution_transaction_payload.fetch("execution_started") == false, "execution transaction preview must not start execution")
  assert(execution_transaction_payload.fetch("request_objects_created") == false, "execution transaction preview must not create request objects")
  assert(execution_transaction_payload.fetch("host_permission_changed") == false, "execution transaction preview must not change host permissions")
  assert(execution_transaction_payload.fetch("host_root_modified") == false, "execution transaction preview must not mutate the host root")
  assert(execution_transaction_payload.fetch("network_required") == false, "execution transaction preview must not require network access")
  assert(execution_transaction_payload.fetch("privileged_container_required") == false, "execution transaction preview must not require privileged containers")
  assert(execution_transaction_payload.fetch("backend_details_exposed") == false, "execution transaction preview must not expose backend details")
  assert(execution_transaction_payload.fetch("blocked_actions").include?("commit launch transaction from preview"), "execution transaction preview must block transaction commits")
  assert(execution_transaction_payload.fetch("blocked_actions").include?("start compatibility profile from transaction preview"), "execution transaction preview must block compatibility profile starts")
  assert(!execution_transaction.downcase.include?("prefix"), "execution transaction preview must not expose implementation storage")
  assert(!execution_transaction.include?(".exe"), "execution transaction preview must not expose a Windows executable")
  assert(!execution_transaction.downcase.include?("program files"), "execution transaction preview must not expose Windows paths")
  assert(!execution_transaction.downcase.include?("qemu-system"), "execution transaction preview must not expose VM implementation commands")
  assert(!execution_transaction.downcase.include?("proton"), "execution transaction preview must not expose backend implementation names")
  assert(!execution_transaction.downcase.include?("wine "), "execution transaction preview must not expose backend implementation names")
  assert(!execution_transaction.downcase.include?("virtual machine"), "execution transaction preview must not expose implementation labels")

  execution_session, execution_session_status = capture_runtime_go(
    project_root,
    runtime_go_binary,
    go_binary,
    "execution-session-preview",
    "--registry",
    "runtime/recipes/registry.json",
    "--app",
    "org.xnix.sample.notepad",
    "--decision",
    "approved",
    "file:///home/test/Documents/example.txt"
  )
  assert(execution_session_status.success?, "Go execution session preview CLI must run successfully")
  execution_session_payload = JSON.parse(execution_session)
  assert(execution_session_payload.fetch("schema_version") == "xnix.runtime.session_identity.v1", "execution session preview schema version must be stable")
  assert(execution_session_payload.fetch("request_type") == "execution-session-preview", "execution session preview must identify its request type")
  assert(execution_session_payload.fetch("session_type") == "compatibility-execution-session", "execution session preview must identify session planning")
  assert(execution_session_payload.fetch("request_state") == "blocked", "execution session preview must remain blocked")
  assert(execution_session_payload.fetch("source") == "execution-transaction-preview", "execution session preview must derive from transaction preview")
  assert(execution_session_payload.fetch("desktop") == "KDE Plasma", "execution session preview must target KDE Plasma")
  assert(execution_session_payload.fetch("runtime_method") == "Launch", "execution session preview must target the Runtime Launch method")
  assert(execution_session_payload.fetch("read_method") == "GetExecutionSessionPreview", "execution session preview must expose the Runtime read method")
  assert(execution_session_payload.fetch("application_id") == recipe.id, "execution session preview must preserve application identity")
  assert(execution_session_payload.fetch("application_name") == recipe.name, "execution session preview must preserve display names")
  assert(execution_session_payload.fetch("desktop_file") == "xnix-#{recipe.id}.desktop", "execution session preview must bind generated desktop files")
  session_transaction = execution_session_payload.fetch("transaction")
  assert(session_transaction.fetch("request_type") == "execution-transaction-preview", "execution session preview must summarize transaction state")
  assert(session_transaction.fetch("step_count") == 7, "execution session preview must preserve transaction step count")
  assert(session_transaction.fetch("blocked_step_count") == 1, "execution session preview must preserve blocked transaction steps")
  assert(session_transaction.fetch("transaction_committed") == false, "execution session preview must not commit transactions")
  assert(session_transaction.fetch("execution_started") == false, "execution session preview transaction summary must not start execution")
  session_window = execution_session_payload.fetch("window_identity")
  assert(session_window.fetch("schema_version") == "xnix.runtime.window_identity.v1", "execution session preview must summarize window identity")
  assert(session_window.fetch("window_kind") == "compatibility-application", "execution session preview must preserve window kind")
  assert(session_window.fetch("class_group") == "xnix-compatibility", "execution session preview must preserve KDE class grouping")
  assert(session_window.fetch("resource_name") == recipe.id, "execution session preview must preserve window resource name")
  assert(session_window.fetch("launcher_url") == "applications:xnix-#{recipe.id}.desktop", "execution session preview must preserve launcher URL")
  assert(session_window.fetch("window_observed") == false, "execution session preview must not claim live windows")
  assert(session_window.fetch("window_registration") == "planned", "execution session preview must keep window registration planned")
  session_task_manager = execution_session_payload.fetch("task_manager")
  assert(session_task_manager.fetch("grouping_key") == recipe.id, "execution session preview must preserve task-manager grouping")
  assert(session_task_manager.fetch("pinning_allowed") == true, "execution session preview must allow normal launcher pinning")
  assert(session_task_manager.fetch("restore_allowed") == true, "execution session preview must preserve restore hints")
  assert(session_task_manager.fetch("entry_planned") == true, "execution session preview must plan task-manager entry identity")
  assert(session_task_manager.fetch("entry_active") == false, "execution session preview must not activate task-manager entries")
  session_kwin = execution_session_payload.fetch("kwin")
  assert(session_kwin.fetch("script_role") == "identity-and-layout", "execution session preview must preserve KWin role")
  assert(session_kwin.fetch("desktop_file") == "xnix-#{recipe.id}.desktop", "execution session preview must preserve KWin desktop file")
  assert(session_kwin.fetch("window_manager_policy_only") == true, "execution session preview must keep KWin policy display-only")
  assert(session_kwin.fetch("rule_planned") == true, "execution session preview must plan KWin rule identity")
  assert(session_kwin.fetch("rule_applied") == false, "execution session preview must not apply KWin rules")
  session_tray = execution_session_payload.fetch("tray")
  assert(session_tray.fetch("status_type") == "tray-status-preview", "execution session preview must summarize tray status")
  assert(session_tray.fetch("registered_application_count") == 1, "execution session preview must preserve registered tray application count")
  assert(session_tray.fetch("active_application_count") == 0, "execution session preview must not claim active tray applications")
  assert(session_tray.fetch("tray_bridge_state") == "planned", "execution session preview must keep tray bridge planned")
  assert(session_tray.fetch("entry_planned") == true, "execution session preview must plan tray entry identity")
  assert(session_tray.fetch("live_bridge_enabled") == false, "execution session preview must not enable live tray bridges")
  assert(execution_session_payload.fetch("runtime_owned") == true, "execution session preview must remain Runtime-owned")
  assert(execution_session_payload.fetch("go_runtime_backed") == true, "execution session preview must be Go Runtime-backed")
  assert(execution_session_payload.fetch("kde_policy_owner") == false, "execution session preview must not make KDE own backend policy")
  assert(execution_session_payload.fetch("compatibility_center_card") == true, "execution session preview must be suitable for Compatibility Center cards")
  assert(execution_session_payload.fetch("safe_for_ai_diagnostics") == true, "execution session preview must be safe for AI diagnostics")
  assert(execution_session_payload.fetch("desktop_entry_launch_visible") == true, "execution session preview must allow KDE to show launcher entries")
  assert(execution_session_payload.fetch("launch_intent_captured") == true, "execution session preview must capture launch intent")
  assert(execution_session_payload.fetch("user_decision_captured") == true, "execution session preview must capture user decision")
  assert(execution_session_payload.fetch("user_decision_allows_launch") == true, "execution session preview must model approved user intent")
  assert(execution_session_payload.fetch("session_plan_created") == true, "execution session preview must create only a read model")
  assert(execution_session_payload.fetch("session_created") == false, "execution session preview must not create live sessions")
  assert(execution_session_payload.fetch("session_registered") == false, "execution session preview must not register live sessions")
  assert(execution_session_payload.fetch("window_observed") == false, "execution session preview must not observe live windows")
  assert(execution_session_payload.fetch("task_manager_entry_planned") == true, "execution session preview must plan task-manager identity")
  assert(execution_session_payload.fetch("task_manager_entry_active") == false, "execution session preview must not activate task-manager entries")
  assert(execution_session_payload.fetch("kwin_rule_planned") == true, "execution session preview must plan KWin identity")
  assert(execution_session_payload.fetch("kwin_rule_applied") == false, "execution session preview must not apply KWin rules")
  assert(execution_session_payload.fetch("tray_entry_planned") == true, "execution session preview must plan tray identity")
  assert(execution_session_payload.fetch("live_tray_bridge_enabled") == false, "execution session preview must not enable live tray bridge")
  assert(execution_session_payload.fetch("transaction_committed") == false, "execution session preview must not commit transactions")
  assert(execution_session_payload.fetch("runtime_launch_approval") == false, "execution session preview must not grant Runtime approval")
  assert(execution_session_payload.fetch("launch_allowed") == false, "execution session preview must not allow launch")
  assert(execution_session_payload.fetch("launch_enabled") == false, "execution session preview must not enable launch")
  assert(execution_session_payload.fetch("execution_started") == false, "execution session preview must not start execution")
  assert(execution_session_payload.fetch("backend_process_started") == false, "execution session preview must not start backend processes")
  assert(execution_session_payload.fetch("request_objects_created") == false, "execution session preview must not create request objects")
  assert(execution_session_payload.fetch("host_root_modified") == false, "execution session preview must not mutate the host root")
  assert(execution_session_payload.fetch("network_required") == false, "execution session preview must not require network access")
  assert(execution_session_payload.fetch("backend_details_exposed") == false, "execution session preview must not expose backend details")
  assert(execution_session_payload.fetch("blocked_actions").include?("create execution session from preview"), "execution session preview must block live session creation")
  assert(execution_session_payload.fetch("blocked_actions").include?("enable live tray bridge before session exists"), "execution session preview must block live tray bridge activation")
  assert(!execution_session.downcase.include?("prefix"), "execution session preview must not expose implementation storage")
  assert(!execution_session.include?(".exe"), "execution session preview must not expose a Windows executable")
  assert(!execution_session.downcase.include?("program files"), "execution session preview must not expose Windows paths")
  assert(!execution_session.downcase.include?("qemu-system"), "execution session preview must not expose VM implementation commands")
  assert(!execution_session.downcase.include?("proton"), "execution session preview must not expose backend implementation names")
  assert(!execution_session.downcase.include?("wine "), "execution session preview must not expose backend implementation names")
  assert(!execution_session.downcase.include?("virtual machine"), "execution session preview must not expose implementation labels")

  execution_session_status, execution_session_status_result = capture_runtime_go(
    project_root,
    runtime_go_binary,
    go_binary,
    "execution-session-status-preview",
    "--registry",
    "runtime/recipes/registry.json",
    "--app",
    "org.xnix.sample.notepad",
    "--decision",
    "approved",
    "file:///home/test/Documents/example.txt"
  )
  assert(execution_session_status_result.success?, "Go execution session status preview CLI must run successfully")
  execution_session_status_payload = JSON.parse(execution_session_status)
  assert(execution_session_status_payload.fetch("schema_version") == "xnix.runtime.session_status.v1", "execution session status preview schema version must be stable")
  assert(execution_session_status_payload.fetch("request_type") == "execution-session-status-preview", "execution session status preview must identify its request type")
  assert(execution_session_status_payload.fetch("status_type") == "compatibility-session-status", "execution session status preview must identify its status type")
  assert(execution_session_status_payload.fetch("session_type") == "compatibility-execution-session", "execution session status preview must preserve session type")
  assert(execution_session_status_payload.fetch("request_state") == "blocked", "execution session status preview must remain blocked")
  assert(execution_session_status_payload.fetch("session_state") == "planned-blocked", "execution session status preview must keep sessions planned and blocked")
  assert(execution_session_status_payload.fetch("source") == "execution-session-preview", "execution session status preview must derive from session preview")
  assert(execution_session_status_payload.fetch("desktop") == "KDE Plasma", "execution session status preview must target KDE Plasma")
  assert(execution_session_status_payload.fetch("runtime_method") == "Launch", "execution session status preview must target the Runtime Launch method")
  assert(execution_session_status_payload.fetch("read_method") == "GetExecutionSessionStatusPreview", "execution session status preview must expose the Runtime read method")
  assert(execution_session_status_payload.fetch("application_id") == recipe.id, "execution session status preview must preserve application identity")
  assert(execution_session_status_payload.fetch("application_name") == recipe.name, "execution session status preview must preserve display names")
  assert(execution_session_status_payload.fetch("desktop_file") == "xnix-#{recipe.id}.desktop", "execution session status preview must bind generated desktop files")
  status_session = execution_session_status_payload.fetch("session")
  assert(status_session.fetch("session_state") == "planned-blocked", "execution session status preview must keep the session blocked")
  assert(status_session.fetch("transaction_state") == "blocked", "execution session status preview must preserve transaction state")
  assert(status_session.fetch("transaction_step_count") == 7, "execution session status preview must preserve transaction step count")
  assert(status_session.fetch("blocked_transaction_steps") == 1, "execution session status preview must preserve blocked transaction steps")
  assert(status_session.fetch("window_registration") == "planned", "execution session status preview must keep window registration planned")
  assert(status_session.fetch("desktop_surface_state") == "planned", "execution session status preview must keep desktop surfaces planned")
  assert(status_session.fetch("compatibility_center_state") == "waiting-for-runtime-gates", "execution session status preview must tell Compatibility Center to wait for Runtime gates")
  assert(status_session.fetch("task_manager_state") == "planned", "execution session status preview must keep task-manager state planned")
  assert(status_session.fetch("tray_state") == "planned", "execution session status preview must keep tray state planned")
  assert(status_session.fetch("session_created") == false, "execution session status preview must not create sessions")
  assert(status_session.fetch("session_registered") == false, "execution session status preview must not register live sessions")
  assert(status_session.fetch("session_active") == false, "execution session status preview must not mark sessions active")
  assert(status_session.fetch("live_state_observed") == false, "execution session status preview must not observe live state")
  assert(status_session.fetch("status_persisted") == false, "execution session status preview must not persist live status")
  assert(status_session.fetch("runtime_launch_approval") == false, "execution session status preview must not grant Runtime launch approval")
  assert(status_session.fetch("launch_allowed") == false, "execution session status preview must not allow launch")
  assert(status_session.fetch("execution_started") == false, "execution session status preview must not start execution")
  assert(status_session.fetch("backend_process_started") == false, "execution session status preview must not start backend processes")
  status_surface = execution_session_status_payload.fetch("desktop_surface")
  assert(status_surface.fetch("window_kind") == "compatibility-application", "execution session status preview must preserve normal window kind")
  assert(status_surface.fetch("class_group") == "xnix-compatibility", "execution session status preview must preserve KDE grouping")
  assert(status_surface.fetch("resource_name") == recipe.id, "execution session status preview must preserve resource name")
  assert(status_surface.fetch("launcher_url") == "applications:xnix-#{recipe.id}.desktop", "execution session status preview must preserve launcher URL")
  assert(status_surface.fetch("window_state") == "not-observed", "execution session status preview must not claim a live window")
  assert(status_surface.fetch("task_manager_grouping_key") == recipe.id, "execution session status preview must preserve task-manager grouping")
  assert(status_surface.fetch("task_manager_state") == "planned", "execution session status preview must keep task-manager status planned")
  assert(status_surface.fetch("kwin_state") == "planned", "execution session status preview must keep KWin status planned")
  assert(status_surface.fetch("tray_state") == "planned", "execution session status preview must keep tray status planned")
  assert(status_surface.fetch("compatibility_center_state") == "waiting-for-runtime-gates", "execution session status preview must preserve Compatibility Center state")
  assert(status_surface.fetch("window_observed") == false, "execution session status preview must not observe live windows")
  assert(status_surface.fetch("task_manager_entry_planned") == true, "execution session status preview must plan task-manager identity")
  assert(status_surface.fetch("task_manager_entry_active") == false, "execution session status preview must not activate task-manager entries")
  assert(status_surface.fetch("kwin_rule_planned") == true, "execution session status preview must plan KWin rules")
  assert(status_surface.fetch("kwin_rule_applied") == false, "execution session status preview must not apply KWin rules")
  assert(status_surface.fetch("tray_entry_planned") == true, "execution session status preview must plan tray identity")
  assert(status_surface.fetch("live_tray_bridge_enabled") == false, "execution session status preview must not enable live tray bridge")
  user_visible_status = execution_session_status_payload.fetch("user_visible_state")
  assert(user_visible_status.fetch("primary_label") == recipe.name, "execution session status preview must preserve the user-visible name")
  assert(user_visible_status.fetch("secondary_label") == "Waiting for Runtime gates", "execution session status preview must expose a user-safe secondary state")
  assert(user_visible_status.fetch("taskbar_badge") == "Planned", "execution session status preview must expose a taskbar badge")
  assert(user_visible_status.fetch("tray_label") == "Ready for review", "execution session status preview must expose a tray label")
  assert(user_visible_status.fetch("compatibility_center_status") == "Runtime gates required", "execution session status preview must expose Compatibility Center status")
  assert(user_visible_status.fetch("next_user_action") == "Open Compatibility Center", "execution session status preview must point users to Compatibility Center")
  assert(user_visible_status.fetch("user_facing_mode") == "Automatic", "execution session status preview must preserve user-facing mode")
  assert(user_visible_status.fetch("user_facing_access") == "Review required", "execution session status preview must preserve user-facing access")
  assert(user_visible_status.fetch("backend_details_exposed") == false, "execution session status preview must not expose backend details in user-visible state")
  status_gates = execution_session_status_payload.fetch("gates")
  assert(status_gates.map { |gate| gate.fetch("id") } == %w[session-identity user-decision runtime-launch-approval live-window-observation desktop-surface-activation], "execution session status preview must preserve gate order")
  assert(execution_session_status_payload.fetch("gate_count") == 5, "execution session status preview must count gates")
  assert(execution_session_status_payload.fetch("passed_gate_count") == 2, "execution session status preview must count passed gates")
  assert(execution_session_status_payload.fetch("required_gate_count") == 1, "execution session status preview must count required gates")
  assert(execution_session_status_payload.fetch("pending_gate_count") == 2, "execution session status preview must count pending gates")
  assert(execution_session_status_payload.fetch("blocked_gate_count") == 1, "execution session status preview must count blocked gates")
  assert(status_gates[2].fetch("status") == "blocked", "execution session status preview must keep Runtime launch approval blocked")
  assert(status_gates[2].fetch("blocks_live_session") == true, "execution session status preview must block live sessions on Runtime launch approval")
  assert(execution_session_status_payload.fetch("runtime_owned") == true, "execution session status preview must remain Runtime-owned")
  assert(execution_session_status_payload.fetch("go_runtime_backed") == true, "execution session status preview must be Go Runtime-backed")
  assert(execution_session_status_payload.fetch("kde_policy_owner") == false, "execution session status preview must not make KDE own backend policy")
  assert(execution_session_status_payload.fetch("compatibility_center_card") == true, "execution session status preview must be suitable for Compatibility Center cards")
  assert(execution_session_status_payload.fetch("safe_for_ai_diagnostics") == true, "execution session status preview must be safe for AI diagnostics")
  assert(execution_session_status_payload.fetch("desktop_entry_launch_visible") == true, "execution session status preview must allow KDE launcher entries")
  assert(execution_session_status_payload.fetch("launch_intent_captured") == true, "execution session status preview must capture launch intent")
  assert(execution_session_status_payload.fetch("user_decision_captured") == true, "execution session status preview must capture user decision")
  assert(execution_session_status_payload.fetch("user_decision_allows_launch") == true, "execution session status preview must model approved user intent")
  assert(execution_session_status_payload.fetch("status_read_model_created") == true, "execution session status preview must create a read model")
  assert(execution_session_status_payload.fetch("session_plan_created") == true, "execution session status preview must reuse session planning")
  assert(execution_session_status_payload.fetch("session_created") == false, "execution session status preview must not create live sessions")
  assert(execution_session_status_payload.fetch("session_registered") == false, "execution session status preview must not register live sessions")
  assert(execution_session_status_payload.fetch("session_active") == false, "execution session status preview must not mark sessions active")
  assert(execution_session_status_payload.fetch("live_state_observed") == false, "execution session status preview must not observe live state")
  assert(execution_session_status_payload.fetch("status_persisted") == false, "execution session status preview must not persist live status")
  assert(execution_session_status_payload.fetch("window_observed") == false, "execution session status preview must not observe live windows")
  assert(execution_session_status_payload.fetch("task_manager_entry_planned") == true, "execution session status preview must plan task-manager identity")
  assert(execution_session_status_payload.fetch("task_manager_entry_active") == false, "execution session status preview must not activate task-manager entries")
  assert(execution_session_status_payload.fetch("kwin_rule_planned") == true, "execution session status preview must plan KWin rules")
  assert(execution_session_status_payload.fetch("kwin_rule_applied") == false, "execution session status preview must not apply KWin rules")
  assert(execution_session_status_payload.fetch("tray_entry_planned") == true, "execution session status preview must plan tray identity")
  assert(execution_session_status_payload.fetch("live_tray_bridge_enabled") == false, "execution session status preview must not enable live tray bridge")
  assert(execution_session_status_payload.fetch("transaction_committed") == false, "execution session status preview must not commit transactions")
  assert(execution_session_status_payload.fetch("runtime_launch_approval") == false, "execution session status preview must not grant Runtime approval")
  assert(execution_session_status_payload.fetch("launch_allowed") == false, "execution session status preview must not allow launch")
  assert(execution_session_status_payload.fetch("launch_enabled") == false, "execution session status preview must not enable launch")
  assert(execution_session_status_payload.fetch("execution_started") == false, "execution session status preview must not start execution")
  assert(execution_session_status_payload.fetch("backend_process_started") == false, "execution session status preview must not start backend processes")
  assert(execution_session_status_payload.fetch("request_objects_created") == false, "execution session status preview must not create request objects")
  assert(execution_session_status_payload.fetch("host_root_modified") == false, "execution session status preview must not mutate the host root")
  assert(execution_session_status_payload.fetch("network_required") == false, "execution session status preview must not require network access")
  assert(execution_session_status_payload.fetch("backend_details_exposed") == false, "execution session status preview must not expose backend details")
  assert(execution_session_status_payload.fetch("blocked_actions").include?("persist live session status from preview"), "execution session status preview must block status persistence")
  assert(execution_session_status_payload.fetch("blocked_actions").include?("mark task manager entry active from status preview"), "execution session status preview must block task-manager activation")
  assert(!execution_session_status.downcase.include?("prefix"), "execution session status preview must not expose implementation storage")
  assert(!execution_session_status.include?(".exe"), "execution session status preview must not expose a Windows executable")
  assert(!execution_session_status.downcase.include?("program files"), "execution session status preview must not expose Windows paths")
  assert(!execution_session_status.downcase.include?("qemu-system"), "execution session status preview must not expose VM implementation commands")
  assert(!execution_session_status.downcase.include?("proton"), "execution session status preview must not expose backend implementation names")
  assert(!execution_session_status.downcase.include?("wine "), "execution session status preview must not expose backend implementation names")
  assert(!execution_session_status.downcase.include?("virtual machine"), "execution session status preview must not expose implementation labels")

  kde_entrypoints, kde_entrypoints_status = capture_runtime_go(
    project_root,
    runtime_go_binary,
    go_binary,
    "kde-entrypoints-preview",
    "--registry",
    "runtime/recipes/registry.json",
    "--app",
    "org.xnix.sample.notepad",
    "--decision",
    "approved",
    "file:///home/test/Documents/example.txt"
  )
  assert(kde_entrypoints_status.success?, "Go KDE entrypoints preview CLI must run successfully")
  kde_entrypoints_payload = JSON.parse(kde_entrypoints)
  assert(kde_entrypoints_payload.fetch("schema_version") == "xnix.runtime.kde_entrypoints.v1", "KDE entrypoints preview schema version must be stable")
  assert(kde_entrypoints_payload.fetch("request_type") == "kde-entrypoints-preview", "KDE entrypoints preview must identify its request type")
  assert(kde_entrypoints_payload.fetch("surface_type") == "kde-first-release-entrypoints", "KDE entrypoints preview must identify the first-release surface")
  assert(kde_entrypoints_payload.fetch("source") == "execution-session-status-preview", "KDE entrypoints preview must derive from execution session status")
  assert(kde_entrypoints_payload.fetch("desktop") == "KDE Plasma", "KDE entrypoints preview must target KDE Plasma")
  assert(kde_entrypoints_payload.fetch("runtime_method") == "Launch", "KDE entrypoints preview must target the Runtime Launch method")
  assert(kde_entrypoints_payload.fetch("read_method") == "GetKDEEntryPointsPreview", "KDE entrypoints preview must expose the Runtime read method")
  assert(kde_entrypoints_payload.fetch("application_id") == recipe.id, "KDE entrypoints preview must preserve application identity")
  assert(kde_entrypoints_payload.fetch("application_name") == recipe.name, "KDE entrypoints preview must preserve display names")
  assert(kde_entrypoints_payload.fetch("desktop_file") == "xnix-#{recipe.id}.desktop", "KDE entrypoints preview must bind generated desktop files")
  kde_session_status = kde_entrypoints_payload.fetch("session_status")
  assert(kde_session_status.fetch("request_type") == "execution-session-status-preview", "KDE entrypoints preview must summarize session status")
  assert(kde_session_status.fetch("session_state") == "planned-blocked", "KDE entrypoints preview must keep session state planned and blocked")
  assert(kde_session_status.fetch("gate_count") == 5, "KDE entrypoints preview must preserve session gate count")
  assert(kde_session_status.fetch("blocked_gate_count") == 1, "KDE entrypoints preview must preserve blocked session gates")
  assert(kde_session_status.fetch("desktop_surface_state") == "planned", "KDE entrypoints preview must keep desktop surface planned")
  assert(kde_session_status.fetch("user_visible_state") == "Runtime gates required", "KDE entrypoints preview must expose user-visible Runtime gate state")
  assert(kde_session_status.fetch("runtime_launch_approval") == false, "KDE entrypoints preview must not grant Runtime launch approval")
  assert(kde_session_status.fetch("launch_allowed") == false, "KDE entrypoints preview must not allow launch")
  assert(kde_session_status.fetch("execution_started") == false, "KDE entrypoints preview must not start execution")
  assert(kde_entrypoints_payload.fetch("entry_point_ids") == %w[launcher task-manager file-manager system-tray notifications compatibility-center settings], "KDE entrypoints preview must preserve the seven first-release entry points")
  assert(kde_entrypoints_payload.fetch("entry_point_count") == 7, "KDE entrypoints preview must count seven entry points")
  assert(kde_entrypoints_payload.fetch("visible_entry_point_count") == 7, "KDE entrypoints preview must make all first-release entry points visible")
  assert(kde_entrypoints_payload.fetch("planned_entry_point_count") == 7, "KDE entrypoints preview must keep all entry points planned")
  assert(kde_entrypoints_payload.fetch("active_entry_point_count") == 0, "KDE entrypoints preview must not mark entry points active")
  assert(kde_entrypoints_payload.fetch("portal_entry_point_count") == 1, "KDE entrypoints preview must count the Portal-mediated file-manager entry point")
  assert(kde_entrypoints_payload.fetch("runtime_gate_entry_point_count") == 7, "KDE entrypoints preview must keep all entry points behind Runtime gates")
  kde_entrypoint_items = kde_entrypoints_payload.fetch("entry_points")
  kde_entrypoint_items.each do |entry|
    assert(entry.fetch("visible") == true, "KDE entrypoint #{entry.fetch("id")} must be visible")
    assert(entry.fetch("planned") == true, "KDE entrypoint #{entry.fetch("id")} must be planned")
    assert(entry.fetch("active") == false, "KDE entrypoint #{entry.fetch("id")} must not be active")
    assert(entry.fetch("requires_runtime_gate") == true, "KDE entrypoint #{entry.fetch("id")} must require Runtime gates")
    assert(entry.fetch("blocked_by_runtime_gate") == true, "KDE entrypoint #{entry.fetch("id")} must remain blocked by Runtime gates")
    assert(entry.fetch("writes_host") == false, "KDE entrypoint #{entry.fetch("id")} must not write host files")
    assert(entry.fetch("starts_backend") == false, "KDE entrypoint #{entry.fetch("id")} must not start backends")
    assert(entry.fetch("backend_details_exposed") == false, "KDE entrypoint #{entry.fetch("id")} must not expose backend details")
  end
  kde_file_manager = kde_entrypoint_items.find { |entry| entry.fetch("id") == "file-manager" }
  assert(kde_file_manager.fetch("kde_component") == "Dolphin", "KDE entrypoints preview must bind file-manager entry to Dolphin")
  assert(kde_file_manager.fetch("runtime_source") == "file-open-preview", "KDE entrypoints preview must route file-manager through file-open previews")
  assert(kde_file_manager.fetch("requires_portal") == true, "KDE entrypoints preview must require Portal mediation for file-manager")
  kde_task_manager = kde_entrypoint_items.find { |entry| entry.fetch("id") == "task-manager" }
  assert(kde_task_manager.fetch("runtime_source") == "execution-session-status-preview", "KDE entrypoints preview must route task-manager through session status")
  assert(kde_entrypoints_payload.fetch("runtime_owned") == true, "KDE entrypoints preview must remain Runtime-owned")
  assert(kde_entrypoints_payload.fetch("go_runtime_backed") == true, "KDE entrypoints preview must be Go Runtime-backed")
  assert(kde_entrypoints_payload.fetch("kde_policy_owner") == false, "KDE entrypoints preview must not make KDE own backend policy")
  assert(kde_entrypoints_payload.fetch("official_desktop_only") == true, "KDE entrypoints preview must keep KDE as the official first desktop")
  assert(kde_entrypoints_payload.fetch("stable_desktop_contract") == true, "KDE entrypoints preview must expose a stable desktop contract")
  assert(kde_entrypoints_payload.fetch("normal_application_surface") == true, "KDE entrypoints preview must present a normal application surface")
  assert(kde_entrypoints_payload.fetch("compatibility_center_card") == true, "KDE entrypoints preview must be suitable for Compatibility Center cards")
  assert(kde_entrypoints_payload.fetch("safe_for_ai_diagnostics") == true, "KDE entrypoints preview must be safe for AI diagnostics")
  assert(kde_entrypoints_payload.fetch("desktop_entry_launch_visible") == true, "KDE entrypoints preview must allow KDE launcher visibility")
  assert(kde_entrypoints_payload.fetch("user_decision_captured") == true, "KDE entrypoints preview must capture user decision")
  assert(kde_entrypoints_payload.fetch("user_decision_allows_launch") == true, "KDE entrypoints preview must model approved user intent")
  assert(kde_entrypoints_payload.fetch("entry_point_plan_created") == true, "KDE entrypoints preview must create a read model")
  assert(kde_entrypoints_payload.fetch("desktop_files_written") == false, "KDE entrypoints preview must not write desktop files")
  assert(kde_entrypoints_payload.fetch("mimeapps_written") == false, "KDE entrypoints preview must not write MIME defaults")
  assert(kde_entrypoints_payload.fetch("settings_persisted") == false, "KDE entrypoints preview must not persist settings")
  assert(kde_entrypoints_payload.fetch("notifications_sent") == false, "KDE entrypoints preview must not send notifications")
  assert(kde_entrypoints_payload.fetch("task_manager_entry_active") == false, "KDE entrypoints preview must not activate task-manager entries")
  assert(kde_entrypoints_payload.fetch("kwin_rule_applied") == false, "KDE entrypoints preview must not apply KWin rules")
  assert(kde_entrypoints_payload.fetch("live_tray_bridge_enabled") == false, "KDE entrypoints preview must not enable live tray bridge")
  assert(kde_entrypoints_payload.fetch("runtime_launch_approval") == false, "KDE entrypoints preview must not grant Runtime approval")
  assert(kde_entrypoints_payload.fetch("launch_allowed") == false, "KDE entrypoints preview must not allow launch")
  assert(kde_entrypoints_payload.fetch("launch_enabled") == false, "KDE entrypoints preview must not enable launch")
  assert(kde_entrypoints_payload.fetch("execution_started") == false, "KDE entrypoints preview must not start execution")
  assert(kde_entrypoints_payload.fetch("backend_process_started") == false, "KDE entrypoints preview must not start backend processes")
  assert(kde_entrypoints_payload.fetch("request_objects_created") == false, "KDE entrypoints preview must not create request objects")
  assert(kde_entrypoints_payload.fetch("host_root_modified") == false, "KDE entrypoints preview must not mutate the host root")
  assert(kde_entrypoints_payload.fetch("network_required") == false, "KDE entrypoints preview must not require network access")
  assert(kde_entrypoints_payload.fetch("backend_details_exposed") == false, "KDE entrypoints preview must not expose backend details")
  assert(kde_entrypoints_payload.fetch("blocked_actions").include?("write KDE entrypoint files from preview"), "KDE entrypoints preview must block host entrypoint writes")
  assert(kde_entrypoints_payload.fetch("blocked_actions").include?("start compatibility profile from entrypoint preview"), "KDE entrypoints preview must block compatibility profile starts")
  assert(!kde_entrypoints.downcase.include?("prefix"), "KDE entrypoints preview must not expose implementation storage")
  assert(!kde_entrypoints.include?(".exe"), "KDE entrypoints preview must not expose a Windows executable")
  assert(!kde_entrypoints.downcase.include?("program files"), "KDE entrypoints preview must not expose Windows paths")
  assert(!kde_entrypoints.downcase.include?("qemu-system"), "KDE entrypoints preview must not expose VM implementation commands")
  assert(!kde_entrypoints.downcase.include?("proton"), "KDE entrypoints preview must not expose backend implementation names")
  assert(!kde_entrypoints.downcase.include?("wine "), "KDE entrypoints preview must not expose backend implementation names")
  assert(!kde_entrypoints.downcase.include?("virtual machine"), "KDE entrypoints preview must not expose implementation labels")

  kde_entrypoint_action, kde_entrypoint_action_status = capture_runtime_go(
    project_root,
    runtime_go_binary,
    go_binary,
    "kde-entrypoint-action-preview",
    "--registry",
    "runtime/recipes/registry.json",
    "--app",
    "org.xnix.sample.notepad",
    "--entrypoint",
    "file-manager",
    "--decision",
    "approved",
    "file:///home/test/Documents/example.txt"
  )
  assert(kde_entrypoint_action_status.success?, "Go KDE entrypoint action preview CLI must run successfully")
  kde_entrypoint_action_payload = JSON.parse(kde_entrypoint_action)
  assert(kde_entrypoint_action_payload.fetch("schema_version") == "xnix.runtime.kde_entrypoint_action.v1", "KDE entrypoint action preview schema version must be stable")
  assert(kde_entrypoint_action_payload.fetch("request_type") == "kde-entrypoint-action-preview", "KDE entrypoint action preview must identify its request type")
  assert(kde_entrypoint_action_payload.fetch("source") == "kde-entrypoints-preview", "KDE entrypoint action preview must derive from KDE entrypoint previews")
  assert(kde_entrypoint_action_payload.fetch("entry_point_id") == "file-manager", "KDE entrypoint action preview must preserve the selected entrypoint")
  assert(kde_entrypoint_action_payload.fetch("kde_component") == "Dolphin", "KDE entrypoint action preview must bind Dolphin actions")
  assert(kde_entrypoint_action_payload.fetch("runtime_source") == "file-open-preview", "KDE entrypoint action preview must route file-manager actions through file-open previews")
  action = kde_entrypoint_action_payload.fetch("action")
  assert(action.fetch("intent") == "open-files", "KDE entrypoint action preview must classify file-manager intent")
  assert(action.fetch("safe_result") == "show file access review", "KDE entrypoint action preview must keep file actions in review")
  assert(action.fetch("requires_portal") == true, "KDE entrypoint action preview must require Portal mediation for file actions")
  assert(action.fetch("starts_backend") == false, "KDE entrypoint action preview must not start backends")
  assert(action.fetch("creates_request_object") == false, "KDE entrypoint action preview must not create request objects")
  assert(kde_entrypoint_action_payload.fetch("entry_point_action_captured") == true, "KDE entrypoint action preview must capture entrypoint actions")
  assert(kde_entrypoint_action_payload.fetch("desktop_files_written") == false, "KDE entrypoint action preview must not write desktop files")
  assert(kde_entrypoint_action_payload.fetch("request_objects_created") == false, "KDE entrypoint action preview must not create Runtime request objects")
  assert(kde_entrypoint_action_payload.fetch("execution_started") == false, "KDE entrypoint action preview must not start execution")
  assert(kde_entrypoint_action_payload.fetch("host_root_modified") == false, "KDE entrypoint action preview must not mutate the host root")
  assert(!kde_entrypoint_action.downcase.include?("prefix"), "KDE entrypoint action preview must not expose implementation storage")
  assert(!kde_entrypoint_action.include?(".exe"), "KDE entrypoint action preview must not expose a Windows executable")
  assert(!kde_entrypoint_action.downcase.include?("virtual machine"), "KDE entrypoint action preview must not expose implementation labels")

  kde_action_queue, kde_action_queue_status = capture_runtime_go(
    project_root,
    runtime_go_binary,
    go_binary,
    "kde-action-queue-preview",
    "--registry",
    "runtime/recipes/registry.json",
    "--app",
    "org.xnix.sample.notepad",
    "--decision",
    "approved",
    "file:///home/test/Documents/example.txt"
  )
  assert(kde_action_queue_status.success?, "Go KDE action queue preview CLI must run successfully")
  kde_action_queue_payload = JSON.parse(kde_action_queue)
  assert(kde_action_queue_payload.fetch("schema_version") == "xnix.runtime.kde_action_queue.v1", "KDE action queue preview schema version must be stable")
  assert(kde_action_queue_payload.fetch("request_type") == "kde-action-queue-preview", "KDE action queue preview must identify its request type")
  assert(kde_action_queue_payload.fetch("queue_type") == "compatibility-center-kde-action-queue", "KDE action queue preview must identify its queue type")
  assert(kde_action_queue_payload.fetch("source") == "kde-entrypoint-action-preview", "KDE action queue preview must derive from entrypoint action previews")
  assert(kde_action_queue_payload.fetch("action_count") == 7, "KDE action queue preview must aggregate the seven first-release actions")
  assert(kde_action_queue_payload.fetch("pending_action_count") == 7, "KDE action queue preview must keep actions pending")
  assert(kde_action_queue_payload.fetch("portal_action_count") == 1, "KDE action queue preview must count the Portal-mediated file action")
  assert(kde_action_queue_payload.fetch("action_queue_created") == true, "KDE action queue preview must create a read model")
  assert(kde_action_queue_payload.fetch("action_queue_persisted") == false, "KDE action queue preview must not persist queues")
  assert(kde_action_queue_payload.fetch("request_objects_created") == false, "KDE action queue preview must not create Runtime request objects")
  assert(kde_action_queue_payload.fetch("execution_started") == false, "KDE action queue preview must not start execution")
  assert(kde_action_queue_payload.fetch("host_root_modified") == false, "KDE action queue preview must not mutate the host root")
  assert(!kde_action_queue.downcase.include?("prefix"), "KDE action queue preview must not expose implementation storage")
  assert(!kde_action_queue.include?(".exe"), "KDE action queue preview must not expose a Windows executable")
  assert(!kde_action_queue.downcase.include?("virtual machine"), "KDE action queue preview must not expose implementation labels")

  kde_action_review, kde_action_review_status = capture_runtime_go(
    project_root,
    runtime_go_binary,
    go_binary,
    "kde-action-review-preview",
    "--registry",
    "runtime/recipes/registry.json",
    "--app",
    "org.xnix.sample.notepad",
    "--action",
    "review-file-manager-action",
    "--decision",
    "approved",
    "file:///home/test/Documents/example.txt"
  )
  assert(kde_action_review_status.success?, "Go KDE action review preview CLI must run successfully")
  kde_action_review_payload = JSON.parse(kde_action_review)
  assert(kde_action_review_payload.fetch("schema_version") == "xnix.runtime.kde_action_review.v1", "KDE action review preview schema version must be stable")
  assert(kde_action_review_payload.fetch("request_type") == "kde-action-review-preview", "KDE action review preview must identify its request type")
  assert(kde_action_review_payload.fetch("source") == "kde-action-queue-preview", "KDE action review preview must derive from action queue previews")
  assert(kde_action_review_payload.fetch("action").fetch("id") == "review-file-manager-action", "KDE action review preview must preserve the queued action")
  assert(kde_action_review_payload.fetch("decision").fetch("decision") == "approved", "KDE action review preview must preserve review decision")
  assert(kde_action_review_payload.fetch("decision").fetch("decision_recorded") == false, "KDE action review preview must not record decisions")
  assert(kde_action_review_payload.fetch("review_receipt_recorded") == false, "KDE action review preview must not record review receipts")
  assert(kde_action_review_payload.fetch("queue_state_changed") == false, "KDE action review preview must not mutate queues")
  assert(kde_action_review_payload.fetch("execution_started") == false, "KDE action review preview must not start execution")
  assert(kde_action_review_payload.fetch("host_root_modified") == false, "KDE action review preview must not mutate the host root")
  assert(!kde_action_review.downcase.include?("prefix"), "KDE action review preview must not expose implementation storage")
  assert(!kde_action_review.include?(".exe"), "KDE action review preview must not expose a Windows executable")
  assert(!kde_action_review.downcase.include?("virtual machine"), "KDE action review preview must not expose implementation labels")

  kde_action_preflight, kde_action_preflight_status = capture_runtime_go(
    project_root,
    runtime_go_binary,
    go_binary,
    "kde-action-preflight-preview",
    "--registry",
    "runtime/recipes/registry.json",
    "--app",
    "org.xnix.sample.notepad",
    "--action",
    "review-file-manager-action",
    "--decision",
    "approved",
    "file:///home/test/Documents/example.txt"
  )
  assert(kde_action_preflight_status.success?, "Go KDE action preflight preview CLI must run successfully")
  kde_action_preflight_payload = JSON.parse(kde_action_preflight)
  assert(kde_action_preflight_payload.fetch("schema_version") == "xnix.runtime.kde_action_preflight.v1", "KDE action preflight preview schema version must be stable")
  assert(kde_action_preflight_payload.fetch("request_type") == "kde-action-preflight-preview", "KDE action preflight preview must identify its request type")
  assert(kde_action_preflight_payload.fetch("source") == "kde-action-review-preview", "KDE action preflight preview must derive from action review previews")
  assert(kde_action_preflight_payload.fetch("action").fetch("id") == "review-file-manager-action", "KDE action preflight preview must preserve the queued action")
  assert(kde_action_preflight_payload.fetch("review").fetch("decision") == "approved", "KDE action preflight preview must preserve review decision")
  assert(kde_action_preflight_payload.fetch("execution_preflight").fetch("request_type") == "execution-preflight-preview", "KDE action preflight preview must summarize execution preflight")
  assert(kde_action_preflight_payload.fetch("action_preflight_created") == true, "KDE action preflight preview must create a read model")
  assert(kde_action_preflight_payload.fetch("preflight_complete") == false, "KDE action preflight preview must not complete preflight")
  assert(kde_action_preflight_payload.fetch("review_receipt_recorded") == false, "KDE action preflight preview must not record review receipts")
  assert(kde_action_preflight_payload.fetch("request_objects_created") == false, "KDE action preflight preview must not create Runtime request objects")
  assert(kde_action_preflight_payload.fetch("execution_started") == false, "KDE action preflight preview must not start execution")
  assert(kde_action_preflight_payload.fetch("host_root_modified") == false, "KDE action preflight preview must not mutate the host root")
  assert(!kde_action_preflight.downcase.include?("prefix"), "KDE action preflight preview must not expose implementation storage")
  assert(!kde_action_preflight.include?(".exe"), "KDE action preflight preview must not expose a Windows executable")
  assert(!kde_action_preflight.downcase.include?("virtual machine"), "KDE action preflight preview must not expose implementation labels")

  kde_action_receipt, kde_action_receipt_status = capture_runtime_go(
    project_root,
    runtime_go_binary,
    go_binary,
    "kde-action-receipt-preview",
    "--registry",
    "runtime/recipes/registry.json",
    "--app",
    "org.xnix.sample.notepad",
    "--action",
    "review-file-manager-action",
    "--decision",
    "approved",
    "file:///home/test/Documents/example.txt"
  )
  assert(kde_action_receipt_status.success?, "Go KDE action receipt preview CLI must run successfully")
  kde_action_receipt_payload = JSON.parse(kde_action_receipt)
  assert(kde_action_receipt_payload.fetch("schema_version") == "xnix.runtime.kde_action_receipt.v1", "KDE action receipt preview schema version must be stable")
  assert(kde_action_receipt_payload.fetch("request_type") == "kde-action-receipt-preview", "KDE action receipt preview must identify its request type")
  assert(kde_action_receipt_payload.fetch("receipt_type") == "compatibility-center-kde-action-review-receipt", "KDE action receipt preview must identify receipt type")
  assert(kde_action_receipt_payload.fetch("source") == "kde-action-preflight-preview", "KDE action receipt preview must derive from action preflight previews")
  assert(kde_action_receipt_payload.fetch("action").fetch("id") == "review-file-manager-action", "KDE action receipt preview must preserve the queued action")
  assert(kde_action_receipt_payload.fetch("review").fetch("decision") == "approved", "KDE action receipt preview must preserve review decision")
  assert(kde_action_receipt_payload.fetch("preflight").fetch("receipt_gate_status") == "pending", "KDE action receipt preview must expose pending receipt gate status")
  assert(kde_action_receipt_payload.fetch("receipt_preview_created") == true, "KDE action receipt preview must create a read model")
  assert(kde_action_receipt_payload.fetch("receipt_recordable") == true, "KDE action receipt preview must describe recordable receipt shape")
  assert(kde_action_receipt_payload.fetch("review_receipt_recorded") == false, "KDE action receipt preview must not record review receipts")
  assert(kde_action_receipt_payload.fetch("request_objects_created") == false, "KDE action receipt preview must not create Runtime request objects")
  assert(kde_action_receipt_payload.fetch("execution_started") == false, "KDE action receipt preview must not start execution")
  assert(kde_action_receipt_payload.fetch("host_root_modified") == false, "KDE action receipt preview must not mutate the host root")
  assert(!kde_action_receipt.downcase.include?("prefix"), "KDE action receipt preview must not expose implementation storage")
  assert(!kde_action_receipt.include?(".exe"), "KDE action receipt preview must not expose a Windows executable")
  assert(!kde_action_receipt.downcase.include?("virtual machine"), "KDE action receipt preview must not expose implementation labels")

  kde_action_status, kde_action_status_status = capture_runtime_go(
    project_root,
    runtime_go_binary,
    go_binary,
    "kde-action-status-preview",
    "--registry",
    "runtime/recipes/registry.json",
    "--app",
    "org.xnix.sample.notepad",
    "--action",
    "review-file-manager-action",
    "--decision",
    "approved",
    "file:///home/test/Documents/example.txt"
  )
  assert(kde_action_status_status.success?, "Go KDE action status preview CLI must run successfully")
  kde_action_status_payload = JSON.parse(kde_action_status)
  assert(kde_action_status_payload.fetch("schema_version") == "xnix.runtime.kde_action_status.v1", "KDE action status preview schema version must be stable")
  assert(kde_action_status_payload.fetch("request_type") == "kde-action-status-preview", "KDE action status preview must identify its request type")
  assert(kde_action_status_payload.fetch("source") == "kde-action-receipt-preview", "KDE action status preview must derive from action receipt previews")
  assert(kde_action_status_payload.fetch("status_state") == "waiting-for-runtime-gates", "KDE action status preview must keep approved actions waiting for Runtime gates")
  assert(kde_action_status_payload.fetch("receipt").fetch("request_type") == "kde-action-receipt-preview", "KDE action status preview must summarize receipt previews")
  assert(kde_action_status_payload.fetch("user_visible_state").fetch("badge") == "Waiting", "KDE action status preview must expose a user-facing waiting badge")
  assert(kde_action_status_payload.fetch("compatibility_center_state") == "waiting-for-runtime-gates", "KDE action status preview must keep Compatibility Center waiting")
  assert(kde_action_status_payload.fetch("status_preview_created") == true, "KDE action status preview must create a read model")
  assert(kde_action_status_payload.fetch("status_persisted") == false, "KDE action status preview must not persist status")
  assert(kde_action_status_payload.fetch("review_receipt_recorded") == false, "KDE action status preview must not record review receipts")
  assert(kde_action_status_payload.fetch("request_objects_created") == false, "KDE action status preview must not create Runtime request objects")
  assert(kde_action_status_payload.fetch("execution_started") == false, "KDE action status preview must not start execution")
  assert(kde_action_status_payload.fetch("host_root_modified") == false, "KDE action status preview must not mutate the host root")
  assert(!kde_action_status.downcase.include?("prefix"), "KDE action status preview must not expose implementation storage")
  assert(!kde_action_status.include?(".exe"), "KDE action status preview must not expose a Windows executable")
  assert(!kde_action_status.downcase.include?("virtual machine"), "KDE action status preview must not expose implementation labels")

  kde_action_card, kde_action_card_status = capture_runtime_go(
    project_root,
    runtime_go_binary,
    go_binary,
    "kde-action-card-preview",
    "--registry",
    "runtime/recipes/registry.json",
    "--app",
    "org.xnix.sample.notepad",
    "--action",
    "review-file-manager-action",
    "--decision",
    "approved",
    "file:///home/test/Documents/example.txt"
  )
  assert(kde_action_card_status.success?, "Go KDE action card preview CLI must run successfully")
  kde_action_card_payload = JSON.parse(kde_action_card)
  assert(kde_action_card_payload.fetch("schema_version") == "xnix.runtime.kde_action_card.v1", "KDE action card preview schema version must be stable")
  assert(kde_action_card_payload.fetch("request_type") == "kde-action-card-preview", "KDE action card preview must identify its request type")
  assert(kde_action_card_payload.fetch("card_type") == "compatibility-center-kde-action-card", "KDE action card preview must identify its card type")
  assert(kde_action_card_payload.fetch("source") == "kde-action-status-preview", "KDE action card preview must derive from action status previews")
  assert(kde_action_card_payload.fetch("card_state") == "waiting-for-runtime-gates", "KDE action card preview must keep approved actions waiting for Runtime gates")
  assert(kde_action_card_payload.fetch("status").fetch("request_type") == "kde-action-status-preview", "KDE action card preview must summarize status previews")
  assert(kde_action_card_payload.fetch("card").fetch("title") == recipe.name, "KDE action card preview must expose the application title")
  assert(kde_action_card_payload.fetch("card").fetch("subtitle") == "Waiting for desktop access review", "KDE action card preview must expose user-facing waiting copy")
  assert(kde_action_card_payload.fetch("card").fetch("badge") == "Waiting", "KDE action card preview must expose a waiting badge")
  assert(kde_action_card_payload.fetch("card").fetch("badge_tone") == "warning", "KDE action card preview must expose a warning badge tone")
  assert(kde_action_card_payload.fetch("card").fetch("primary_action").fetch("id") == "review-required-gates", "KDE action card preview must expose the gate review primary action")
  assert(kde_action_card_payload.fetch("card").fetch("primary_action").fetch("navigation_only") == true, "KDE action card primary action must remain navigation-only")
  assert(kde_action_card_payload.fetch("card").fetch("primary_action").fetch("mutates_runtime") == false, "KDE action card primary action must not mutate Runtime state")
  assert(kde_action_card_payload.fetch("card").fetch("primary_action").fetch("starts_program") == false, "KDE action card primary action must not start programs")
  assert(kde_action_card_payload.fetch("card").fetch("disabled_actions").map { |action| action.fetch("id") } == %w[start-application record-review-receipt grant-resource-access], "KDE action card preview must disable execution and grant actions")
  assert(kde_action_card_payload.fetch("card_preview_created") == true, "KDE action card preview must create a read model")
  assert(kde_action_card_payload.fetch("card_persisted") == false, "KDE action card preview must not persist cards")
  assert(kde_action_card_payload.fetch("status_persisted") == false, "KDE action card preview must not persist status")
  assert(kde_action_card_payload.fetch("notifications_sent") == false, "KDE action card preview must not send notifications")
  assert(kde_action_card_payload.fetch("resource_grant_created") == false, "KDE action card preview must not grant resources")
  assert(kde_action_card_payload.fetch("request_objects_created") == false, "KDE action card preview must not create Runtime request objects")
  assert(kde_action_card_payload.fetch("execution_started") == false, "KDE action card preview must not start execution")
  assert(kde_action_card_payload.fetch("host_root_modified") == false, "KDE action card preview must not mutate the host root")
  assert(!kde_action_card.downcase.include?("prefix"), "KDE action card preview must not expose implementation storage")
  assert(!kde_action_card.include?(".exe"), "KDE action card preview must not expose a Windows executable")
  assert(!kde_action_card.downcase.include?("virtual machine"), "KDE action card preview must not expose implementation labels")

  kde_action_card_deck, kde_action_card_deck_status = capture_runtime_go(
    project_root,
    runtime_go_binary,
    go_binary,
    "kde-action-card-deck-preview",
    "--registry",
    "runtime/recipes/registry.json",
    "--app",
    "org.xnix.sample.notepad",
    "--decision",
    "approved",
    "file:///home/test/Documents/example.txt"
  )
  assert(kde_action_card_deck_status.success?, "Go KDE action card deck preview CLI must run successfully")
  kde_action_card_deck_payload = JSON.parse(kde_action_card_deck)
  assert(kde_action_card_deck_payload.fetch("schema_version") == "xnix.runtime.kde_action_card_deck.v1", "KDE action card deck preview schema version must be stable")
  assert(kde_action_card_deck_payload.fetch("request_type") == "kde-action-card-deck-preview", "KDE action card deck preview must identify its request type")
  assert(kde_action_card_deck_payload.fetch("deck_type") == "compatibility-center-kde-action-card-deck", "KDE action card deck preview must identify its deck type")
  assert(kde_action_card_deck_payload.fetch("source") == "kde-action-card-preview", "KDE action card deck preview must derive from action card previews")
  assert(kde_action_card_deck_payload.fetch("runtime_method") == "GetKDEActionCardDeck", "KDE action card deck preview must expose the Runtime method")
  assert(kde_action_card_deck_payload.fetch("queue").fetch("request_type") == "kde-action-queue-preview", "KDE action card deck preview must summarize action queue previews")
  assert(kde_action_card_deck_payload.fetch("card_count") == 7, "KDE action card deck preview must render the seven first-release cards")
  assert(kde_action_card_deck_payload.fetch("waiting_card_count") == 7, "KDE action card deck preview must keep approved cards waiting for Runtime gates")
  assert(kde_action_card_deck_payload.fetch("deferred_card_count") == 0, "KDE action card deck preview must not defer approved cards")
  assert(kde_action_card_deck_payload.fetch("rejected_card_count") == 0, "KDE action card deck preview must not reject approved cards")
  assert(kde_action_card_deck_payload.fetch("navigation_action_count") == 28, "KDE action card deck preview must expose navigation-only actions")
  assert(kde_action_card_deck_payload.fetch("disabled_action_count") == 21, "KDE action card deck preview must disable execution and grant actions on all cards")
  assert(kde_action_card_deck_payload.fetch("primary_card_id") == "#{recipe.id}:review-launcher-action:card", "KDE action card deck preview must prefer launcher as the primary card")
  deck_cards = kde_action_card_deck_payload.fetch("cards")
  assert(deck_cards.length == 7, "KDE action card deck preview must include seven card records")
  assert(deck_cards.first.fetch("action_id") == "review-launcher-action", "KDE action card deck preview must keep launcher first")
  assert(deck_cards.first.fetch("card").fetch("badge") == "Waiting", "KDE action card deck preview must expose waiting card copy")
  assert(deck_cards.first.fetch("primary_action").fetch("navigation_only") == true, "KDE action card deck primary actions must remain navigation-only")
  assert(deck_cards.first.fetch("primary_action").fetch("mutates_runtime") == false, "KDE action card deck primary actions must not mutate Runtime state")
  assert(deck_cards.first.fetch("primary_action").fetch("starts_program") == false, "KDE action card deck primary actions must not start programs")
  file_card = deck_cards.find { |card| card.fetch("action_id") == "review-file-manager-action" }
  assert(file_card.fetch("entry_point_id") == "file-manager", "KDE action card deck preview must include the Dolphin card")
  assert(file_card.fetch("requires_portal") == true, "KDE action card deck preview must keep Dolphin card Portal-aware")
  assert(file_card.fetch("required_runtime_gate") == "portal-file-open-review", "KDE action card deck preview must preserve the file-open Runtime gate")
  assert(kde_action_card_deck_payload.fetch("deck_preview_created") == true, "KDE action card deck preview must create a read model")
  assert(kde_action_card_deck_payload.fetch("deck_persisted") == false, "KDE action card deck preview must not persist decks")
  assert(kde_action_card_deck_payload.fetch("cards_persisted") == false, "KDE action card deck preview must not persist cards")
  assert(kde_action_card_deck_payload.fetch("card_actions_enabled") == false, "KDE action card deck preview must not enable card actions")
  assert(kde_action_card_deck_payload.fetch("card_actions_persisted") == false, "KDE action card deck preview must not persist card actions")
  assert(kde_action_card_deck_payload.fetch("status_persisted") == false, "KDE action card deck preview must not persist status")
  assert(kde_action_card_deck_payload.fetch("review_receipt_recorded") == false, "KDE action card deck preview must not record review receipts")
  assert(kde_action_card_deck_payload.fetch("request_objects_created") == false, "KDE action card deck preview must not create Runtime request objects")
  assert(kde_action_card_deck_payload.fetch("resource_grant_created") == false, "KDE action card deck preview must not grant resources")
  assert(kde_action_card_deck_payload.fetch("notifications_sent") == false, "KDE action card deck preview must not send notifications")
  assert(kde_action_card_deck_payload.fetch("execution_started") == false, "KDE action card deck preview must not start execution")
  assert(kde_action_card_deck_payload.fetch("host_root_modified") == false, "KDE action card deck preview must not mutate the host root")
  assert(!kde_action_card_deck.downcase.include?("prefix"), "KDE action card deck preview must not expose implementation storage")
  assert(!kde_action_card_deck.include?(".exe"), "KDE action card deck preview must not expose a Windows executable")
  assert(!kde_action_card_deck.downcase.include?("virtual machine"), "KDE action card deck preview must not expose implementation labels")

  kde_center_page, kde_center_page_status = capture_runtime_go(
    project_root,
    runtime_go_binary,
    go_binary,
    "kde-center-page-preview",
    "--registry",
    "runtime/recipes/registry.json",
    "--app",
    "org.xnix.sample.notepad",
    "--decision",
    "approved",
    "file:///home/test/Documents/example.txt"
  )
  assert(kde_center_page_status.success?, "Go KDE center page preview CLI must run successfully")
  kde_center_page_payload = JSON.parse(kde_center_page)
  assert(kde_center_page_payload.fetch("schema_version") == "xnix.runtime.kde_center_page.v1", "KDE center page preview schema version must be stable")
  assert(kde_center_page_payload.fetch("request_type") == "kde-center-page-preview", "KDE center page preview must identify its request type")
  assert(kde_center_page_payload.fetch("page_type") == "compatibility-center-application-page", "KDE center page preview must identify the application page")
  assert(kde_center_page_payload.fetch("source") == "compatibility-center-preview+backend-selection-preview+kde-action-card-deck-preview+settings-preview", "KDE center page preview must compose existing Runtime read models")
  assert(kde_center_page_payload.fetch("runtime_method") == "GetKDECenterPage", "KDE center page preview must expose the Runtime method")
  assert(kde_center_page_payload.fetch("read_method") == "GetKDECenterPagePreview", "KDE center page preview must expose the read method")
  assert(kde_center_page_payload.fetch("header").fetch("title") == recipe.name, "KDE center page header must expose the app name")
  assert(kde_center_page_payload.fetch("header").fetch("badge") == "Review ready", "KDE center page header must expose review-ready state")
  assert(kde_center_page_payload.fetch("application_summary").fetch("compatibility_state") == "registered", "KDE center page must include the Compatibility Center application summary")
  assert(kde_center_page_payload.fetch("action_deck").fetch("request_type") == "kde-action-card-deck-preview", "KDE center page must include the action deck summary")
  assert(kde_center_page_payload.fetch("action_deck").fetch("card_count") == 7, "KDE center page must include all seven first-release action cards")
  assert(kde_center_page_payload.fetch("action_deck").fetch("waiting_card_count") == 7, "KDE center page must keep cards waiting for Runtime gates")
  assert(kde_center_page_payload.fetch("settings_snapshot").fetch("request_type") == "settings-preview", "KDE center page must include the settings snapshot")
  assert(kde_center_page_payload.fetch("settings_snapshot").fetch("section_count") == 5, "KDE center page settings snapshot must include all user-facing sections")
  backend_selection_snapshot = kde_center_page_payload.fetch("backend_selection_snapshot")
  assert(backend_selection_snapshot.fetch("request_type") == "backend-selection-preview", "KDE center page must include backend selection snapshots")
  assert(backend_selection_snapshot.fetch("runtime_method") == "GetBackendSelectionPlan", "KDE center page backend selection must route to Runtime selection reads")
  assert(backend_selection_snapshot.fetch("recommended_profile_id") == "local-compatibility", "KDE center page backend selection must expose the Runtime recommendation")
  assert(backend_selection_snapshot.fetch("candidate_count") == 2, "KDE center page backend selection must expose candidate counts")
  assert(backend_selection_snapshot.fetch("selection_committed") == false, "KDE center page backend selection must keep commit disabled")
  assert(backend_selection_snapshot.fetch("selection_change_enabled") == false, "KDE center page backend selection must keep selection changes disabled")
  assert(backend_selection_snapshot.fetch("backend_launch_enabled") == false, "KDE center page backend selection must keep launch disabled")
  assert(backend_selection_snapshot.fetch("environment_created") == false, "KDE center page backend selection must not create environments")
  assert(backend_selection_snapshot.fetch("host_root_modified") == false, "KDE center page backend selection must not mutate the host root")
  assert(backend_selection_snapshot.fetch("backend_details_exposed") == false, "KDE center page backend selection must hide backend details")
  activation_status_snapshot = kde_center_page_payload.fetch("activation_status_snapshot")
  assert(activation_status_snapshot.fetch("request_type") == "desktop-activation-status-preview", "KDE center page must include desktop activation status snapshots")
  assert(activation_status_snapshot.fetch("runtime_method") == "GetDesktopActivationStatus", "KDE center page activation status must route to Runtime status reads")
  assert(activation_status_snapshot.fetch("renderer") == "xnix-runtime-go desktop-activation-status-preview", "KDE center page activation status must use the Go renderer")
  assert(activation_status_snapshot.fetch("activation_state") == "ready-for-runtime-commit", "KDE center page must expose activation state")
  assert(activation_status_snapshot.fetch("commit_enabled") == false, "KDE center page activation status must keep commit disabled")
  assert(activation_status_snapshot.fetch("launch_enabled") == false, "KDE center page activation status must keep launch disabled")
  assert(activation_status_snapshot.fetch("host_root_modified") == false, "KDE center page activation status must not mutate the host root")
  assert(activation_status_snapshot.fetch("backend_details_exposed") == false, "KDE center page activation status must hide backend details")
  assert(kde_center_page_payload.fetch("navigation_count") == 6, "KDE center page must expose navigation sections")
  assert(kde_center_page_payload.fetch("primary_navigation_target") == "compatibility-center-gates", "KDE center page must route primary review to gates")
  assert(kde_center_page_payload.fetch("runtime_owned") == true, "KDE center page must remain Runtime-owned")
  assert(kde_center_page_payload.fetch("go_runtime_backed") == true, "KDE center page must be backed by Go Runtime product logic")
  assert(kde_center_page_payload.fetch("kde_policy_owner") == false, "KDE center page must not make KDE own policy")
  assert(kde_center_page_payload.fetch("page_preview_created") == true, "KDE center page must create a read model")
  assert(kde_center_page_payload.fetch("page_persisted") == false, "KDE center page must not persist page state")
  assert(kde_center_page_payload.fetch("card_actions_enabled") == false, "KDE center page must not enable card actions")
  assert(kde_center_page_payload.fetch("settings_persisted") == false, "KDE center page must not persist settings")
  assert(kde_center_page_payload.fetch("request_objects_created") == false, "KDE center page must not create Runtime request objects")
  assert(kde_center_page_payload.fetch("resource_grant_created") == false, "KDE center page must not grant resources")
  assert(kde_center_page_payload.fetch("notifications_sent") == false, "KDE center page must not send notifications")
  assert(kde_center_page_payload.fetch("execution_started") == false, "KDE center page must not start execution")
  assert(kde_center_page_payload.fetch("host_root_modified") == false, "KDE center page must not mutate the host root")
  assert(!kde_center_page.downcase.include?("prefix"), "KDE center page preview must not expose implementation storage")
  assert(!kde_center_page.include?(".exe"), "KDE center page preview must not expose a Windows executable")
  assert(!kde_center_page.downcase.include?("virtual machine"), "KDE center page preview must not expose implementation labels")

  kde_center_page_sections, kde_center_page_sections_status = capture_runtime_go(
    project_root,
    runtime_go_binary,
    go_binary,
    "kde-center-page-sections-preview",
    "--registry",
    "runtime/recipes/registry.json",
    "--app",
    "org.xnix.sample.notepad",
    "--decision",
    "approved",
    "file:///home/test/Documents/example.txt"
  )
  assert(kde_center_page_sections_status.success?, "Go KDE center page sections preview CLI must run successfully")
  kde_center_page_sections_payload = JSON.parse(kde_center_page_sections)
  assert(kde_center_page_sections_payload.fetch("schema_version") == "xnix.runtime.kde_center_page_sections.v1", "KDE center page sections preview schema version must be stable")
  assert(kde_center_page_sections_payload.fetch("request_type") == "kde-center-page-sections-preview", "KDE center page sections preview must identify its request type")
  assert(kde_center_page_sections_payload.fetch("runtime_method") == "GetKDECenterPageSections", "KDE center page sections preview must expose the Runtime method")
  assert(kde_center_page_sections_payload.fetch("read_method") == "GetKDECenterPageSectionsPreview", "KDE center page sections preview must expose the read method")
  assert(kde_center_page_sections_payload.fetch("section_count") == 6, "KDE center page sections preview must expose six sections")
  assert(kde_center_page_sections_payload.fetch("read_only_section_count") == 6, "KDE center page sections preview must mark all sections read-only")
  assert(kde_center_page_sections_payload.fetch("navigation_only_section_count") == 6, "KDE center page sections preview must mark all sections navigation-only")
  assert(kde_center_page_sections_payload.fetch("executable_section_count") == 0, "KDE center page sections preview must not expose executable sections")
  assert(kde_center_page_sections_payload.fetch("sections").map { |section| section.fetch("id") } == %w[overview backend activation actions settings diagnostics], "KDE center page sections preview must preserve page section order")
  assert(kde_center_page_sections_payload.fetch("sections").map { |section| section.fetch("runtime_method") } == %w[GetCompatibilityCenterSummary GetBackendSelectionPlan GetDesktopActivationStatus GetCompatibilityActionQueue GetCompatibilitySettings GetDiagnostics], "KDE center page sections preview must bind sections to Runtime read methods")
  assert(kde_center_page_sections_payload.fetch("sections").all? { |section| section.fetch("read_only") && section.fetch("navigation_only") }, "KDE center page sections must remain read-only navigation")
  assert(kde_center_page_sections_payload.fetch("runtime_owned") == true, "KDE center page sections must remain Runtime-owned")
  assert(kde_center_page_sections_payload.fetch("go_runtime_backed") == true, "KDE center page sections must be backed by Go Runtime product logic")
  assert(kde_center_page_sections_payload.fetch("sections_preview_created") == true, "KDE center page sections must create a read model")
  assert(kde_center_page_sections_payload.fetch("sections_persisted") == false, "KDE center page sections must not persist section state")
  assert(kde_center_page_sections_payload.fetch("section_actions_enabled") == false, "KDE center page sections must not enable actions")
  assert(kde_center_page_sections_payload.fetch("settings_persisted") == false, "KDE center page sections must not persist settings")
  assert(kde_center_page_sections_payload.fetch("request_objects_created") == false, "KDE center page sections must not create Runtime request objects")
  assert(kde_center_page_sections_payload.fetch("resource_grant_created") == false, "KDE center page sections must not grant resources")
  assert(kde_center_page_sections_payload.fetch("notifications_sent") == false, "KDE center page sections must not send notifications")
  assert(kde_center_page_sections_payload.fetch("execution_started") == false, "KDE center page sections must not start execution")
  assert(kde_center_page_sections_payload.fetch("host_root_modified") == false, "KDE center page sections must not mutate the host root")
  assert(!kde_center_page_sections.downcase.include?("prefix"), "KDE center page sections preview must not expose implementation storage")
  assert(!kde_center_page_sections.include?(".exe"), "KDE center page sections preview must not expose a Windows executable")
  assert(!kde_center_page_sections.downcase.include?("virtual machine"), "KDE center page sections preview must not expose implementation labels")

  kde_center_page_section_detail, kde_center_page_section_detail_status = capture_runtime_go(
    project_root,
    runtime_go_binary,
    go_binary,
    "kde-center-page-section-detail-preview",
    "--registry",
    "runtime/recipes/registry.json",
    "--app",
    "org.xnix.sample.notepad",
    "--section",
    "settings",
    "--decision",
    "approved",
    "file:///home/test/Documents/example.txt"
  )
  assert(kde_center_page_section_detail_status.success?, "Go KDE center page section detail preview CLI must run successfully")
  kde_center_page_section_detail_payload = JSON.parse(kde_center_page_section_detail)
  assert(kde_center_page_section_detail_payload.fetch("schema_version") == "xnix.runtime.kde_center_page_section_detail.v1", "KDE center page section detail preview schema version must be stable")
  assert(kde_center_page_section_detail_payload.fetch("request_type") == "kde-center-page-section-detail-preview", "KDE center page section detail preview must identify its request type")
  assert(kde_center_page_section_detail_payload.fetch("runtime_method") == "GetKDECenterPageSectionDetail", "KDE center page section detail preview must expose the Runtime method")
  assert(kde_center_page_section_detail_payload.fetch("read_method") == "GetKDECenterPageSectionDetailPreview", "KDE center page section detail preview must expose the read method")
  assert(kde_center_page_section_detail_payload.fetch("section_id") == "settings", "KDE center page section detail preview must preserve the selected section")
  assert(kde_center_page_section_detail_payload.fetch("section_runtime_method") == "GetCompatibilitySettings", "KDE center page section detail preview must route settings to Runtime settings")
  assert(kde_center_page_section_detail_payload.fetch("section_read_model") == "settings-model", "KDE center page section detail preview must expose the backing read model")
  assert(kde_center_page_section_detail_payload.fetch("available_section_ids") == %w[overview backend activation actions settings diagnostics], "KDE center page section detail preview must expose available navigation")
  assert(kde_center_page_section_detail_payload.fetch("read_only_navigation") == true, "KDE center page section detail preview must remain read-only navigation")
  assert(kde_center_page_section_detail_payload.fetch("runtime_owned") == true, "KDE center page section detail must remain Runtime-owned")
  assert(kde_center_page_section_detail_payload.fetch("go_runtime_backed") == true, "KDE center page section detail must be backed by Go Runtime product logic")
  assert(kde_center_page_section_detail_payload.fetch("detail_preview_created") == true, "KDE center page section detail must create a read model")
  assert(kde_center_page_section_detail_payload.fetch("detail_persisted") == false, "KDE center page section detail must not persist detail state")
  assert(kde_center_page_section_detail_payload.fetch("section_actions_enabled") == false, "KDE center page section detail must not enable actions")
  assert(kde_center_page_section_detail_payload.fetch("settings_persisted") == false, "KDE center page section detail must not persist settings")
  assert(kde_center_page_section_detail_payload.fetch("request_objects_created") == false, "KDE center page section detail must not create Runtime request objects")
  assert(kde_center_page_section_detail_payload.fetch("resource_grant_created") == false, "KDE center page section detail must not grant resources")
  assert(kde_center_page_section_detail_payload.fetch("notifications_sent") == false, "KDE center page section detail must not send notifications")
  assert(kde_center_page_section_detail_payload.fetch("execution_started") == false, "KDE center page section detail must not start execution")
  assert(kde_center_page_section_detail_payload.fetch("host_root_modified") == false, "KDE center page section detail must not mutate the host root")
  assert(!kde_center_page_section_detail.downcase.include?("prefix"), "KDE center page section detail preview must not expose implementation storage")
  assert(!kde_center_page_section_detail.include?(".exe"), "KDE center page section detail preview must not expose a Windows executable")
  assert(!kde_center_page_section_detail.downcase.include?("virtual machine"), "KDE center page section detail preview must not expose implementation labels")

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
