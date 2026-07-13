#!/usr/bin/env ruby
# frozen_string_literal: true

require "pathname"

PROJECT_ROOT = Pathname.new(__dir__).join("..").realpath
EXPECTED_VERSION = "0.2.37"
REQUIRED_FILES = %w[
  Dockerfile
  VERSION
  buildroot/Config.in
  buildroot/external.desc
  buildroot/external.mk
  buildroot/sources.lock
  buildroot/configs/xnix_x86_64_defconfig
  buildroot/board/xnix/rootfs-overlay/etc/ssh/sshd_config
  buildroot/board/xnix/post-build.sh
  lib/xnix/container.rb
  lib/xnix/buildroot.rb
  lib/xnix/qemu.rb
  lib/xnix/serial_log.rb
  lib/xnix/sshd.rb
  lib/xnix/ssh_probe.rb
  lib/xnix/ssh_test_key.rb
  lib/xnix/compatibility/ai_diagnostic_input.rb
  lib/xnix/compatibility/compatibility_engine_catalog.rb
  lib/xnix/compatibility/compatibility_repair_plan.rb
  lib/xnix/compatibility/compatibility_test_plan.rb
  lib/xnix/compatibility/compatibility_test_result.rb
  lib/xnix/compatibility/application_recipe.rb
  lib/xnix/compatibility/compatibility_snapshot_plan.rb
  lib/xnix/compatibility/compatibility_run_plan.rb
  lib/xnix/compatibility/dbus_runtime_client.rb
  lib/xnix/compatibility/desktop_activation_installer.rb
  lib/xnix/compatibility/desktop_activation_rollback.rb
  lib/xnix/compatibility/desktop_integration_manifest.rb
  lib/xnix/compatibility/desktop_entry.rb
  lib/xnix/compatibility/dolphin_service_menu.rb
  lib/xnix/compatibility/file_association_model.rb
  lib/xnix/compatibility/file_open_request.rb
  lib/xnix/compatibility/kde_center_model.rb
  lib/xnix/compatibility/kde_integration_status.rb
  lib/xnix/compatibility/kwin_window_rule.rb
  lib/xnix/compatibility/krunner_model.rb
  lib/xnix/compatibility/launch_request.rb
  lib/xnix/compatibility/notification_request.rb
  lib/xnix/compatibility/portal_access_policy.rb
  lib/xnix/compatibility/portal_request_model.rb
  lib/xnix/compatibility/registry_backed_recipe_store.rb
  lib/xnix/compatibility/recipe_install_gate.rb
  lib/xnix/compatibility/recipe_registry.rb
  lib/xnix/compatibility/recipe_store.rb
  lib/xnix/compatibility/recipe_trust_policy.rb
  lib/xnix/compatibility/runtime_daemon.rb
  lib/xnix/compatibility/settings_model.rb
  lib/xnix/compatibility/task_manager_identity.rb
  lib/xnix/compatibility/tray_status_model.rb
  lib/xnix/milestone.rb
  bin/xnix-ai-diagnostic-input
  bin/xnix-compatd
  bin/xnix-compat-engine-catalog
  bin/xnix-compat-launch
  bin/xnix-compat-notify
  bin/xnix-compat-open
  bin/xnix-compat-repair-plan
  bin/xnix-compat-run-plan
  bin/xnix-compat-snapshot-plan
  bin/xnix-compat-test-plan
  bin/xnix-compat-test-result
  bin/xnix-compat-settings
  bin/xnix-compat-tray-status
  bin/xnix-compat-window-identity
  bin/xnix-desktop-integration-manifest
  bin/xnix-file-association-model
  bin/xnix-install-desktop-integration
  bin/xnix-kde-center-model
  bin/xnix-kde-integration-status
  bin/xnix-kwin-window-rule
  bin/xnix-krunner-model
  bin/xnix-portal-access-policy
  bin/xnix-portal-request-model
  bin/xnix-recipe-install-gate
  bin/xnix-recipe-registry
  bin/xnix-recipe-trust-policy
  bin/xnix-rollback-desktop-integration
  libexec/xnix/compatd
  scripts/container.rb
  scripts/dbus_session_smoke.rb
  scripts/kde_center_dbus_smoke.rb
  scripts/fetch_buildroot.rb
  scripts/full_smoke.rb
  scripts/install_runtime_activation.rb
  scripts/prepare_ssh_test_key.rb
  scripts/ssh_smoke.rb
  runtime/dbus/org.xnix.Compatibility1.xml
  runtime/dbus/org.xnix.Compatibility1.service
  runtime/dbus/xnix_compatd_smoke.c
  runtime/recipes/registry.json
  runtime/recipes/org.xnix.sample.notepad.json
  runtime/systemd/xnix-compatd.service
  kde/plasmoids/org.xnix.compatibilitycenter/metadata.json
  kde/plasmoids/org.xnix.compatibilitycenter/contents/ui/main.qml
  kde/dolphin/servicemenus/xnix-open-with-compatibility.desktop
  docs/compatibility-runtime.md
  test/test_ai_diagnostic_input.rb
  test/test_container.rb
  test/test_buildroot.rb
  test/test_qemu.rb
  test/test_serial_log.rb
  test/test_sshd.rb
  test/test_ssh_probe.rb
  test/test_milestone.rb
  test/test_compatibility_engine_catalog.rb
  test/test_application_recipe.rb
  test/test_compatibility_repair_plan.rb
  test/test_compatibility_snapshot_plan.rb
  test/test_compatibility_run_plan.rb
  test/test_compatibility_test_plan.rb
  test/test_compatibility_test_result.rb
  test/test_recipe_store.rb
  test/test_recipe_registry.rb
  test/test_recipe_install_gate.rb
  test/test_recipe_trust_policy.rb
  test/test_registry_backed_recipe_store.rb
  test/test_desktop_entry.rb
  test/test_desktop_activation_installer.rb
  test/test_desktop_activation_rollback.rb
  test/test_desktop_integration_manifest.rb
  test/test_dbus_runtime_client.rb
  test/test_dolphin_service_menu.rb
  test/test_file_association_model.rb
  test/test_file_open_request.rb
  test/test_launch_request.rb
  test/test_notification_request.rb
  test/test_portal_access_policy.rb
  test/test_portal_request_model.rb
  test/test_runtime_contract.rb
  test/test_runtime_daemon.rb
  test/test_runtime_dispatch.rb
  test/test_runtime_activation.rb
  test/test_runtime_activation_install.rb
  test/test_runtime_dbus_smoke_script.rb
  test/test_kde_center_model.rb
  test/test_kde_integration_status.rb
  test/test_kwin_window_rule.rb
  test/test_krunner_model.rb
  test/test_settings_model.rb
  test/test_task_manager_identity.rb
  test/test_tray_status_model.rb
].freeze
FORBIDDEN_CONTAINER_TOKENS = ["--privileged", "--network host", "docker.sock"].freeze
REQUIRED_CONFIG_LINES = [
  "BR2_x86_64=y",
  "BR2_SYSTEM_DHCP=\"eth0\"",
  "BR2_TARGET_GENERIC_GETTY_PORT=\"ttyS0\"",
  "BR2_LINUX_KERNEL=y",
  "BR2_TARGET_ROOTFS_INITRAMFS=y",
  "BR2_PACKAGE_OPENSSH_SERVER=y"
].freeze
REQUIRED_VERSION_MARKERS = {
  "CLAUDE.md" => "Version: `#{EXPECTED_VERSION}`",
  "PRODUCT_OVERVIEW.md" => "Current version: v#{EXPECTED_VERSION}",
  "CHANGELOG.md" => "## [#{EXPECTED_VERSION}]",
  "README.md" => "`v#{EXPECTED_VERSION}`"
}.freeze

def read_project_file(relative_path)
  PROJECT_ROOT.join(relative_path).read
end

def assert(condition, message)
  return if condition

  warn "FAIL: #{message}"
  exit 1
end

REQUIRED_FILES.each do |relative_path|
  assert(PROJECT_ROOT.join(relative_path).file?, "missing required file: #{relative_path}")
end

assert(read_project_file("VERSION").strip == EXPECTED_VERSION, "VERSION must be #{EXPECTED_VERSION}")

REQUIRED_VERSION_MARKERS.each do |relative_path, marker|
  assert(read_project_file(relative_path).include?(marker), "#{relative_path} must reference #{EXPECTED_VERSION}")
end

source_lock = read_project_file("buildroot/sources.lock")
assert(source_lock.include?("buildroot.version=2025.02.15"), "source lock must pin Buildroot 2025.02.15")
assert(source_lock.match?(/^buildroot.sha256=[0-9a-f]{64}$/), "source lock must contain a SHA-256")

config = read_project_file("buildroot/configs/xnix_x86_64_defconfig")
REQUIRED_CONFIG_LINES.each do |line|
  assert(config.include?(line), "defconfig must include #{line}")
end

dockerfile = read_project_file("Dockerfile")
assert(dockerfile.include?("qemu-system-x86"), "Dockerfile must install QEMU system emulation")
assert(dockerfile.include?("ruby"), "Dockerfile must install Ruby for build utilities")
assert(dockerfile.include?("dbus"), "Dockerfile must install D-Bus tooling for runtime smoke tests")
assert(dockerfile.include?("libglib2.0-bin"), "Dockerfile must install gdbus for runtime smoke tests")
assert(dockerfile.include?("libglib2.0-dev"), "Dockerfile must install GIO headers for runtime smoke tests")
assert(dockerfile.include?("pkg-config"), "Dockerfile must install pkg-config for runtime smoke tests")
assert(dockerfile.include?("USER xnix"), "Dockerfile must run as the non-root xnix user")
FORBIDDEN_CONTAINER_TOKENS.each do |token|
  assert(!dockerfile.include?(token), "Dockerfile must not contain #{token}")
end

plasmoid_metadata = read_project_file("kde/plasmoids/org.xnix.compatibilitycenter/metadata.json")
assert(plasmoid_metadata.include?("\"Version\": \"#{EXPECTED_VERSION}\""), "Plasma metadata must reference #{EXPECTED_VERSION}")
assert(plasmoid_metadata.include?("xnix-kde-center-model"), "Plasma metadata must declare the Runtime read-model command")

dolphin_service_menu = read_project_file("kde/dolphin/servicemenus/xnix-open-with-compatibility.desktop")
assert(dolphin_service_menu.include?("X-KDE-ServiceTypes=KonqPopupMenu/Plugin"), "Dolphin service menu must target Dolphin popup integration")
assert(dolphin_service_menu.include?("Exec=xnix-compat-open %U"), "Dolphin service menu must delegate to the Runtime file-open command")

desktop_entry_source = read_project_file("lib/xnix/compatibility/desktop_entry.rb")
assert(desktop_entry_source.include?("xnix-compat-launch"), "Desktop entries must delegate to the Runtime launcher command")

file_association_source = read_project_file("lib/xnix/compatibility/file_association_model.rb")
assert(file_association_source.include?("mimeapps.list"), "File association model must write mimeapps.list")
assert(file_association_source.include?("xnix-compat-open"), "File association model must keep file opens Runtime-mediated")
assert(file_association_source.include?("\"overwrite_existing_mimeapps\" => false"), "File association model must not permit blind mimeapps overwrite")
assert(file_association_source.include?("\"backend_details_exposed\" => false"), "File association model must hide backend details")

engine_catalog_source = read_project_file("lib/xnix/compatibility/compatibility_engine_catalog.rb")
assert(engine_catalog_source.include?("xnix-compat-engine-catalog"), "Compatibility engine catalog must expose a CLI command")
%w[automatic-managed local-compatibility-engine isolated-compatibility-engine].each do |engine_id|
  assert(engine_catalog_source.include?("\"#{engine_id}\""), "Compatibility engine catalog must include #{engine_id}")
end
assert(engine_catalog_source.include?("\"backend_details_exposed\" => false"), "Compatibility engine catalog must hide backend details")
assert(engine_catalog_source.include?("\"launch_enabled\" => false"), "Compatibility engine catalog must not claim launch enablement yet")

run_plan_source = read_project_file("lib/xnix/compatibility/compatibility_run_plan.rb")
assert(run_plan_source.include?("xnix-compat-run-plan"), "Compatibility run plan must expose a CLI command")
assert(run_plan_source.include?("automatic-managed"), "Compatibility run plan must include automatic strategy")
assert(run_plan_source.include?("local-compatibility-engine"), "Compatibility run plan must include local compatibility strategy")
assert(run_plan_source.include?("isolated-compatibility-engine"), "Compatibility run plan must include isolated compatibility strategy")
assert(run_plan_source.include?("\"backend_details_exposed\" => false"), "Compatibility run plan must hide backend details")
assert(run_plan_source.include?("CompatibilityEngineCatalog"), "Compatibility run plan must select engines through the catalog")

launch_request_source = read_project_file("lib/xnix/compatibility/launch_request.rb")
assert(launch_request_source.include?("CompatibilityRunPlan"), "Launch requests must include compatibility run planning")
assert(launch_request_source.include?("\"run_plan\""), "Launch requests must expose a run plan summary")

repair_plan_source = read_project_file("lib/xnix/compatibility/compatibility_repair_plan.rb")
assert(repair_plan_source.include?("xnix-compat-repair-plan"), "Compatibility repair plan must expose a CLI command")
%w[engine-binding-pending portal-approval-required recipe-trust-blocked runtime-repair-applied].each do |issue|
  assert(repair_plan_source.include?("\"#{issue}\""), "Compatibility repair plan must include #{issue}")
end
assert(repair_plan_source.include?("\"snapshot_required\""), "Compatibility repair plan must expose snapshot requirements")
assert(repair_plan_source.include?("\"rollback_available\" => true"), "Compatibility repair plan must keep rollback available")
assert(repair_plan_source.include?("CompatibilitySnapshotPlan"), "Compatibility repair plan must include snapshot planning")

snapshot_plan_source = read_project_file("lib/xnix/compatibility/compatibility_snapshot_plan.rb")
assert(snapshot_plan_source.include?("xnix-compat-snapshot-plan"), "Compatibility snapshot plan must expose a CLI command")
%w[before-repair before-engine-change manual].each do |reason|
  assert(snapshot_plan_source.include?("\"#{reason}\""), "Compatibility snapshot plan must include #{reason}")
end
assert(snapshot_plan_source.include?("\"host_system\" => false"), "Compatibility snapshot plan must not snapshot the host system")
assert(snapshot_plan_source.include?("\"user_documents\" => false"), "Compatibility snapshot plan must not snapshot user documents")
assert(snapshot_plan_source.include?("\"preserve_user_documents\" => true"), "Compatibility snapshot restore must preserve user documents")

test_plan_source = read_project_file("lib/xnix/compatibility/compatibility_test_plan.rb")
assert(test_plan_source.include?("xnix-compat-test-plan"), "Compatibility test plan must expose a CLI command")
%w[compatibility-test portal-preflight snapshot-preflight runtime-launch-binding].each do |token|
  assert(test_plan_source.include?(token), "Compatibility test plan must include #{token}")
end
assert(test_plan_source.include?("\"backend_details_exposed\" => false"), "Compatibility test plan must hide backend details")

test_result_source = read_project_file("lib/xnix/compatibility/compatibility_test_result.rb")
assert(test_result_source.include?("xnix-compat-test-result"), "Compatibility test result must expose a CLI command")
%w[compatibility-test-result waiting-for-runtime safe_for_ai_diagnostics runtime-launch-binding].each do |token|
  assert(test_result_source.include?(token), "Compatibility test result must include #{token}")
end
assert(test_result_source.include?("\"backend_details_exposed\" => false"), "Compatibility test result must hide backend details")

ai_diagnostic_source = read_project_file("lib/xnix/compatibility/ai_diagnostic_input.rb")
assert(ai_diagnostic_source.include?("xnix-ai-diagnostic-input"), "AI diagnostic input must expose a CLI command")
%w[ai-diagnostic-input diagnostic_signals privacy_boundaries allowed_ai_tasks blocked_ai_tasks].each do |token|
  assert(ai_diagnostic_source.include?(token), "AI diagnostic input must include #{token}")
end
assert(ai_diagnostic_source.include?("\"ai_provider_called\" => false"), "AI diagnostic input must not call an AI provider")
assert(ai_diagnostic_source.include?("\"network_required\" => false"), "AI diagnostic input must not require network access")
assert(ai_diagnostic_source.include?("\"backend_details_exposed\" => false"), "AI diagnostic input must hide backend details")

notification_source = read_project_file("lib/xnix/compatibility/notification_request.rb")
%w[install-failed repair-applied mode-changed approval-required].each do |event_type|
  assert(notification_source.include?("\"#{event_type}\""), "Notification requests must include #{event_type}")
end

settings_source = read_project_file("lib/xnix/compatibility/settings_model.rb")
%w[run-mode resource-access devices network snapshots].each do |section_id|
  assert(settings_source.include?("\"id\" => \"#{section_id}\""), "Settings model must include #{section_id}")
end

portal_policy_source = read_project_file("lib/xnix/compatibility/portal_access_policy.rb")
%w[file-open uri-open print screenshot clipboard camera remote-desktop].each do |operation|
  assert(portal_policy_source.include?("\"#{operation}\""), "Portal access policy must include #{operation}")
end
assert(portal_policy_source.include?("org.freedesktop.portal"), "Portal access policy must use XDG Desktop Portal interfaces")
assert(portal_policy_source.include?("\"direct_access_allowed\" => false"), "Portal access policy must deny direct desktop access")

portal_request_source = read_project_file("lib/xnix/compatibility/portal_request_model.rb")
assert(portal_request_source.include?("org.freedesktop.portal.Desktop"), "Portal request model must target the portal desktop service")
assert(portal_request_source.include?("\"portal-request\""), "Portal request model must identify request type")
assert(portal_request_source.include?("\"host_permission_changed\" => false"), "Portal request model must not mutate host permissions")
assert(portal_request_source.include?("\"backend_details_exposed\" => false"), "Portal request model must hide backend details")

tray_status_source = read_project_file("lib/xnix/compatibility/tray_status_model.rb")
%w[runtime_activity compatibility_status tray_bridge].each do |method_name|
  assert(tray_status_source.include?(method_name), "Tray status model must include #{method_name}")
end

task_manager_source = read_project_file("lib/xnix/compatibility/task_manager_identity.rb")
%w[task_manager kwin grouping_key launcher_url].each do |token|
  assert(task_manager_source.include?(token), "Task manager identity must include #{token}")
end

kwin_rule_source = read_project_file("lib/xnix/compatibility/kwin_window_rule.rb")
%w[identity-and-layout skip_taskbar show_in_switcher runtime_owns_backend_policy].each do |token|
  assert(kwin_rule_source.include?(token), "KWin window rule must include #{token}")
end
assert(kwin_rule_source.include?("\"backend_details_exposed\" => false"), "KWin window rule must hide backend details")

kde_status_source = read_project_file("lib/xnix/compatibility/kde_integration_status.rb")
%w[launcher task-manager file-manager system-tray notifications compatibility-center settings].each do |entry_point|
  assert(kde_status_source.include?("\"id\" => \"#{entry_point}\""), "KDE integration status must include #{entry_point}")
end

desktop_manifest_source = read_project_file("lib/xnix/compatibility/desktop_integration_manifest.rb")
%w[xnix-compat-launch xnix-compat-window-identity xnix-kwin-window-rule xnix-file-association-model xnix-compat-open xnix-compat-tray-status xnix-compat-notify xnix-kde-center-model xnix-compat-settings xnix-portal-access-policy xnix-portal-request-model].each do |command|
  assert(desktop_manifest_source.include?(command), "Desktop integration manifest must include #{command}")
end
assert(desktop_manifest_source.include?("\"backend_commands_exposed\" => false"), "Desktop integration manifest must hide backend commands")

desktop_activation_source = read_project_file("lib/xnix/compatibility/desktop_activation_installer.rb")
%w[usr/share/applications usr/share/kio/servicemenus usr/share/xnix/compatibility/manifests usr/share/xnix/compatibility/activation-receipts].each do |target|
  assert(desktop_activation_source.include?(target), "Desktop activation installer must stage #{target}")
end
assert(desktop_activation_source.include?("refusing to overwrite existing mimeapps list"), "Desktop activation installer must preserve existing mimeapps lists")
assert(desktop_activation_source.include?("\"host_root_modified\" => false"), "Desktop activation installer must not claim host root changes")
assert(desktop_activation_source.include?("RecipeInstallGate"), "Desktop activation installer must enforce recipe install gate preflight")
assert(desktop_activation_source.include?("recipe_install_gate_enforced"), "Desktop activation installer must report install gate enforcement")

desktop_rollback_source = read_project_file("lib/xnix/compatibility/desktop_activation_rollback.rb")
assert(desktop_rollback_source.include?("xnix-rollback-desktop-integration"), "Desktop activation rollback must expose a CLI command")
assert(desktop_rollback_source.include?("sha256_verified_before_remove"), "Desktop activation rollback must verify checksums before removal")
assert(desktop_rollback_source.include?("refusing to remove changed file"), "Desktop activation rollback must preserve changed files")

recipe_registry_source = read_project_file("lib/xnix/compatibility/recipe_registry.rb")
assert(recipe_registry_source.include?("xnix-recipe-registry"), "Recipe registry must expose a CLI command")
assert(recipe_registry_source.include?("DIGEST_PATTERN"), "Recipe registry must validate SHA-256 digests")
assert(recipe_registry_source.include?("signed_recipe_validation"), "Recipe registry must report signed recipe validation status")

registry_backed_store_source = read_project_file("lib/xnix/compatibility/registry_backed_recipe_store.rb")
assert(registry_backed_store_source.include?("RecipeRegistry"), "Registry-backed recipe store must verify the registry")
assert(registry_backed_store_source.include?("RecipeStore.new"), "Registry-backed recipe store must keep a no-registry development fallback")

runtime_daemon_source = read_project_file("lib/xnix/compatibility/runtime_daemon.rb")
assert(runtime_daemon_source.include?("RegistryBackedRecipeStore.for_path"), "Runtime daemon must use registry-backed recipe loading")
assert(runtime_daemon_source.include?("\"recipe_trust\""), "Runtime daemon must expose recipe trust status")
assert(runtime_daemon_source.include?("\"registry_backed_recipe_store\""), "Runtime daemon must expose registry-backed store capability")
assert(runtime_daemon_source.include?("\"compatibility_test_planning\""), "Runtime daemon must expose compatibility test planning capability")
assert(runtime_daemon_source.include?("\"compatibility_test_results\""), "Runtime daemon must expose compatibility test result capability")
assert(runtime_daemon_source.include?("\"ai_diagnostic_inputs\""), "Runtime daemon must expose AI diagnostic input capability")
%w[GetEngineCatalog GetRunPlan GetRepairPlan GetTestPlan GetTestResult GetAIDiagnosticInput GetSnapshotPlan GetPortalAccessPolicy].each do |method_name|
  assert(runtime_daemon_source.include?("\"#{method_name}\""), "Runtime daemon dispatch must include #{method_name}")
  assert(read_project_file("runtime/dbus/org.xnix.Compatibility1.xml").include?("name=\"#{method_name}\""), "D-Bus contract must include #{method_name}")
  assert(read_project_file("runtime/dbus/xnix_compatd_smoke.c").include?(method_name), "D-Bus smoke adapter must include #{method_name}")
  assert(read_project_file("scripts/dbus_session_smoke.rb").include?(method_name), "D-Bus session smoke must call #{method_name}")
end

dbus_client_source = read_project_file("lib/xnix/compatibility/dbus_runtime_client.rb")
%w[engine_catalog run_plan repair_plan test_plan test_result ai_diagnostic_input snapshot_plan portal_access_policy].each do |method_name|
  assert(dbus_client_source.include?("def #{method_name}"), "D-Bus Runtime client must expose #{method_name}")
end
assert(dbus_client_source.include?("return true if value == \"true\""), "D-Bus Runtime client must parse boolean true values")
assert(dbus_client_source.include?("return false if value == \"false\""), "D-Bus Runtime client must parse boolean false values")
assert(dbus_client_source.include?("value.start_with?(\"[\")"), "D-Bus Runtime client must parse string arrays")

krunner_source = read_project_file("lib/xnix/compatibility/krunner_model.rb")
assert(krunner_source.include?("xnix-compat-launch"), "KRunner model must delegate launches to the managed launcher")
assert(krunner_source.include?("\"backend_details_exposed\" => false"), "KRunner model must hide backend details")
assert(krunner_source.include?("extension_launch_query?"), "KRunner model must support file-oriented natural queries")

recipe_trust_policy_source = read_project_file("lib/xnix/compatibility/recipe_trust_policy.rb")
assert(recipe_trust_policy_source.include?("xnix-recipe-trust-policy"), "Recipe trust policy must expose a CLI command")
assert(recipe_trust_policy_source.include?("production-trusted"), "Recipe trust policy must model production trust")
assert(recipe_trust_policy_source.include?("development-only"), "Recipe trust policy must model development-only trust")

recipe_install_gate_source = read_project_file("lib/xnix/compatibility/recipe_install_gate.rb")
assert(recipe_install_gate_source.include?("xnix-recipe-install-gate"), "Recipe install gate must expose a CLI command")
assert(recipe_install_gate_source.include?("\"recipe-install\""), "Recipe install gate must identify install decisions")
assert(recipe_install_gate_source.include?("production signed source"), "Recipe install gate must require production signed sources")
assert(recipe_install_gate_source.include?("development"), "Recipe install gate must preserve development staging")

puts "PASS: Xnix #{EXPECTED_VERSION} scaffold is consistent"
