#!/usr/bin/env ruby
# frozen_string_literal: true

require "pathname"

PROJECT_ROOT = Pathname.new(__dir__).join("..").realpath
EXPECTED_VERSION = "0.2.133"
REQUIRED_FILES = %w[
  Dockerfile
  VERSION
  go.mod
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
  lib/xnix/compatibility/action_review_receipt.rb
  lib/xnix/compatibility/ai_diagnostic_input.rb
  lib/xnix/compatibility/ai_diagnostic_recommendation.rb
  lib/xnix/compatibility/ai_repair_approval_gate.rb
  lib/xnix/compatibility/application_state_root.rb
  lib/xnix/compatibility/compatibility_acquisition_preflight.rb
  lib/xnix/compatibility/compatibility_action_queue.rb
  lib/xnix/compatibility/compatibility_artifact_manifest.rb
  lib/xnix/compatibility/compatibility_backend_binding.rb
  lib/xnix/compatibility/compatibility_backend_capability_matrix.rb
  lib/xnix/compatibility/compatibility_backend_selection_plan.rb
  lib/xnix/compatibility/compatibility_backend_environment_plan.rb
  lib/xnix/compatibility/compatibility_backend_lifecycle.rb
  lib/xnix/compatibility/compatibility_engine_catalog.rb
  lib/xnix/compatibility/compatibility_execution_readiness.rb
  lib/xnix/compatibility/compatibility_install_plan.rb
  lib/xnix/compatibility/compatibility_mode_switch_plan.rb
  lib/xnix/compatibility/compatibility_permission_review_plan.rb
  lib/xnix/compatibility/compatibility_review_flow_plan.rb
  lib/xnix/compatibility/compatibility_package_source.rb
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
  lib/xnix/compatibility/desktop_resource_bridge_plan.rb
  lib/xnix/compatibility/dolphin_service_menu.rb
  lib/xnix/compatibility/file_association_model.rb
  lib/xnix/compatibility/kde_application_surface_plan.rb
  lib/xnix/compatibility/kde_shell_integration_plan.rb
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
  lib/xnix/compatibility/runtime_live_owner_gate.rb
  lib/xnix/compatibility/runtime_method_parity_manifest.rb
  lib/xnix/compatibility/runtime_owner_smoke_plan.rb
  lib/xnix/compatibility/settings_change_plan.rb
  lib/xnix/compatibility/runtime_service_binding.rb
  lib/xnix/compatibility/runtime_write_gate.rb
  lib/xnix/compatibility/settings_model.rb
  lib/xnix/compatibility/task_manager_identity.rb
  lib/xnix/compatibility/tray_status_model.rb
  lib/xnix/milestone.rb
  bin/xnix-ai-diagnostic-input
  bin/xnix-ai-diagnostic-recommendation
  bin/xnix-ai-repair-approval-gate
  bin/xnix-compat-acquisition-preflight
  bin/xnix-compat-action-review
  bin/xnix-compat-action-queue
  bin/xnix-compat-artifact-manifest
  bin/xnix-compat-install-plan
  bin/xnix-compat-backend-capability-matrix
  bin/xnix-compat-mode-switch-plan
  bin/xnix-compat-permission-review-plan
  bin/xnix-compat-review-flow-plan
  bin/xnix-compat-state-root
  bin/xnix-compat-backend-binding
  bin/xnix-compat-backend-environment-plan
  bin/xnix-compat-backend-selection-plan
  bin/xnix-compat-backend-lifecycle
  bin/xnix-compatd
  bin/xnix-compat-engine-catalog
  bin/xnix-compat-execution-readiness
  bin/xnix-compat-package-source
  bin/xnix-compat-launch
  bin/xnix-compat-notify
  bin/xnix-compat-open
  bin/xnix-compat-repair-plan
  bin/xnix-compat-run-plan
  bin/xnix-compat-snapshot-plan
  bin/xnix-compat-test-plan
  bin/xnix-compat-test-result
  bin/xnix-compat-settings
  bin/xnix-compat-settings-change
  bin/xnix-compat-tray-status
  bin/xnix-compat-window-identity
  bin/xnix-desktop-integration-manifest
  bin/xnix-desktop-resource-bridge-plan
  bin/xnix-file-association-model
  bin/xnix-install-desktop-integration
  bin/xnix-kde-center-model
  bin/xnix-kde-application-surface-plan
  bin/xnix-kde-integration-status
  bin/xnix-kde-shell-integration-plan
  bin/xnix-kwin-window-rule
  bin/xnix-krunner-model
  bin/xnix-portal-access-policy
  bin/xnix-portal-request-model
  bin/xnix-recipe-install-gate
  bin/xnix-recipe-registry
  bin/xnix-recipe-trust-policy
  bin/xnix-rollback-desktop-integration
  bin/xnix-runtime-live-owner-gate
  bin/xnix-runtime-method-parity-manifest
  bin/xnix-runtime-owner-smoke-plan
  bin/xnix-runtime-service-binding
  bin/xnix-runtime-write-gate
  cmd/xnix-runtime-go/main.go
  cmd/xnix-runtime-go/main_test.go
  internal/runtime/appidentity/identity.go
  internal/runtime/appidentity/identity_test.go
  internal/runtime/appidentity/registry.go
  internal/runtime/appidentity/registry_test.go
  libexec/xnix/compatd
  scripts/container.rb
  scripts/dbus_session_smoke.rb
  scripts/kde_center_dbus_smoke.rb
  scripts/fetch_buildroot.rb
  scripts/full_smoke.rb
  scripts/install_runtime_activation.rb
  scripts/prepare_ssh_test_key.rb
  scripts/runtime_activation_smoke.rb
  scripts/ssh_smoke.rb
  runtime/dbus/org.xnix.Compatibility1.xml
  runtime/dbus/org.xnix.Compatibility1.service
  runtime/dbus/xnix_compatd_smoke.c
  runtime/core/xnix_runtime_core.h
  runtime/core/xnix_runtime_core.c
  runtime/core/xnix_runtime_core_action_review.inc
  runtime/core/xnix_runtime_core_compatibility_center.inc
  runtime/core/xnix_runtime_core_run_plan.inc
  runtime/core/xnix_runtime_core_desktop_activation.inc
  runtime/core/xnix_runtime_core_kde_integration_status.inc
  runtime/core/xnix_runtime_core_kde_shell_integration_plan.inc
  runtime/core/xnix_runtime_core_kde_application_surface_plan.inc
  runtime/core/xnix_runtime_core_desktop_resource_bridge_plan.inc
  runtime/core/xnix_runtime_core_compatibility_mode_switch_plan.inc
  runtime/core/xnix_runtime_core_compatibility_permission_review_plan.inc
  runtime/core/xnix_runtime_core_compatibility_review_flow_plan.inc
  runtime/core/xnix_runtime_core_desktop_entry.inc
  runtime/core/xnix_runtime_core_task_manager_identity.inc
  runtime/core/xnix_runtime_core_kwin_window_rule.inc
  runtime/core/xnix_runtime_core_file_association.inc
  runtime/core/xnix_runtime_core_notification.inc
  runtime/core/xnix_runtime_core_tray_status.inc
  runtime/core/xnix_runtime_core_krunner_query.inc
  runtime/core/xnix_runtime_core_portal_request.inc
  runtime/core/xnix_runtime_core_compatibility_snapshot_plan.inc
  runtime/core/xnix_runtime_core_compatibility_test_plan.inc
  runtime/core/xnix_runtime_core_compatibility_test_result.inc
  runtime/core/xnix_runtime_core_execution_readiness.inc
  runtime/core/xnix_runtime_core_launch_intent.inc
  runtime/core/xnix_runtime_core_backend_capability_matrix.inc
  runtime/core/xnix_runtime_core_backend_selection_plan.inc
  runtime/core/xnix_runtime_core_compatibility_repair_plan.inc
  runtime/core/xnix_runtime_core_ai_diagnostic_input.inc
  runtime/core/xnix_runtime_core_ai_diagnostic_recommendation.inc
  runtime/core/xnix_runtime_core_ai_repair_approval_gate.inc
  runtime/core/xnix_runtime_core_cli.c
  runtime/core/xnix_runtime_core_cli_probe.inc
  runtime/core/xnix_runtime_core_cli_action_queue.inc
  runtime/core/xnix_runtime_core_cli_action_review.inc
  runtime/core/xnix_runtime_core_cli_compatibility_center.inc
  runtime/core/xnix_runtime_core_cli_run_plan.inc
  runtime/core/xnix_runtime_core_cli_desktop_activation.inc
  runtime/core/xnix_runtime_core_cli_kde_integration_status.inc
  runtime/core/xnix_runtime_core_cli_kde_shell_integration_plan.inc
  runtime/core/xnix_runtime_core_cli_kde_application_surface_plan.inc
  runtime/core/xnix_runtime_core_cli_desktop_resource_bridge_plan.inc
  runtime/core/xnix_runtime_core_cli_compatibility_mode_switch_plan.inc
  runtime/core/xnix_runtime_core_cli_compatibility_permission_review_plan.inc
  runtime/core/xnix_runtime_core_cli_compatibility_review_flow_plan.inc
  runtime/core/xnix_runtime_core_cli_backend_capability_matrix.inc
  runtime/core/xnix_runtime_core_cli_backend_selection_plan.inc
  runtime/core/xnix_runtime_core_cli_desktop_entry.inc
  runtime/core/xnix_runtime_core_cli_task_manager_identity.inc
  runtime/core/xnix_runtime_core_cli_kwin_window_rule.inc
  runtime/core/xnix_runtime_core_cli_file_association.inc
  runtime/core/xnix_runtime_core_cli_notification.inc
  runtime/core/xnix_runtime_core_cli_tray_status.inc
  runtime/core/xnix_runtime_core_cli_krunner_query.inc
  runtime/core/xnix_runtime_core_cli_portal_request.inc
  runtime/core/xnix_runtime_core_cli_compatibility_snapshot_plan.inc
  runtime/core/xnix_runtime_core_cli_compatibility_test_plan.inc
  runtime/core/xnix_runtime_core_cli_compatibility_test_result.inc
  runtime/core/xnix_runtime_core_cli_execution_readiness.inc
  runtime/core/xnix_runtime_core_cli_launch_intent.inc
  runtime/core/xnix_runtime_core_cli_compatibility_repair_plan.inc
  runtime/core/xnix_runtime_core_cli_ai_diagnostic_input.inc
  runtime/core/xnix_runtime_core_cli_ai_diagnostic_recommendation.inc
  runtime/core/xnix_runtime_core_cli_ai_repair_approval_gate.inc
  runtime/core/xnix_runtime_core_cli_dispatch.inc
  runtime/recipes/registry.json
  runtime/recipes/org.xnix.sample.notepad.json
  runtime/systemd/xnix-compatd.service
  test/fixtures/xnix-runtime-go-fake
  kde/plasmoids/org.xnix.compatibilitycenter/metadata.json
  kde/plasmoids/org.xnix.compatibilitycenter/contents/ui/main.qml
  kde/dolphin/servicemenus/xnix-open-with-compatibility.desktop
  docs/compatibility-runtime.md
  test/test_action_review_receipt.rb
  test/test_ai_diagnostic_input.rb
  test/test_ai_diagnostic_recommendation.rb
  test/test_ai_repair_approval_gate.rb
  test/test_application_state_root.rb
  test/test_compatibility_acquisition_preflight.rb
  test/test_compatibility_action_queue.rb
  test/test_compatibility_artifact_manifest.rb
  test/test_compatibility_backend_capability_matrix.rb
  test/test_compatibility_backend_selection_plan.rb
  test/test_compatibility_install_plan.rb
  test/test_compatibility_mode_switch_plan.rb
  test/test_compatibility_permission_review_plan.rb
  test/test_compatibility_review_flow_plan.rb
  test/test_compatibility_backend_binding.rb
  test/test_compatibility_backend_environment_plan.rb
  test/test_compatibility_backend_lifecycle.rb
  test/test_kde_application_surface_plan.rb
  test/test_container.rb
  test/test_buildroot.rb
  test/test_qemu.rb
  test/test_serial_log.rb
  test/test_sshd.rb
  test/test_ssh_probe.rb
  test/test_milestone.rb
  test/test_compatibility_engine_catalog.rb
  test/test_compatibility_package_source.rb
  test/test_application_recipe.rb
  test/test_compatibility_repair_plan.rb
  test/test_compatibility_snapshot_plan.rb
  test/test_compatibility_run_plan.rb
  test/test_compatibility_test_plan.rb
  test/test_compatibility_test_result.rb
  test/test_compatibility_execution_readiness.rb
  test/test_recipe_store.rb
  test/test_recipe_registry.rb
  test/test_recipe_install_gate.rb
  test/test_recipe_trust_policy.rb
  test/test_registry_backed_recipe_store.rb
  test/test_desktop_entry.rb
  test/test_desktop_activation_installer.rb
  test/test_desktop_activation_rollback.rb
  test/test_desktop_integration_manifest.rb
  test/test_desktop_resource_bridge_plan.rb
  test/test_dbus_runtime_client.rb
  test/test_dolphin_service_menu.rb
  test/test_file_association_model.rb
  test/test_go_desktop_identity_plan.rb
  test/test_file_open_request.rb
  test/test_full_smoke_script.rb
  test/test_launch_request.rb
  test/test_notification_request.rb
  test/test_portal_access_policy.rb
  test/test_portal_request_model.rb
  test/test_runtime_contract.rb
  test/test_runtime_core.rb
  test/test_runtime_daemon.rb
  test/test_runtime_dispatch.rb
  test/test_runtime_live_owner_gate.rb
  test/test_runtime_method_parity_manifest.rb
  test/test_runtime_owner_smoke_plan.rb
  test/test_runtime_service_binding.rb
  test/test_runtime_write_gate.rb
  test/test_runtime_activation.rb
  test/test_runtime_activation_install.rb
  test/test_runtime_activation_smoke_script.rb
  test/test_runtime_dbus_smoke_script.rb
  test/test_kde_center_model.rb
  test/test_kde_integration_status.rb
  test/test_kde_shell_integration_plan.rb
  test/test_kwin_window_rule.rb
  test/test_krunner_model.rb
  test/test_settings_model.rb
  test/test_settings_change_plan.rb
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
assert(dockerfile.include?("xnix_runtime_core.c"), "Dockerfile must compile the C Runtime core")
assert(dockerfile.include?("xnix-runtime-core"), "Dockerfile must install the C Runtime core smoke binary")
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
%w[runtime-launch-intent GetLaunchIntent RuntimeWriteGate execution_request_created launch_enabled].each do |token|
  assert(launch_request_source.include?(token), "Launch requests must include #{token}")
end

repair_plan_source = read_project_file("lib/xnix/compatibility/compatibility_repair_plan.rb")
assert(repair_plan_source.include?("xnix-compat-repair-plan"), "Compatibility repair plan must expose a CLI command")
%w[engine-binding-pending portal-approval-required recipe-trust-blocked runtime-repair-applied].each do |issue|
  assert(repair_plan_source.include?("\"#{issue}\""), "Compatibility repair plan must include #{issue}")
end
assert(repair_plan_source.include?("\"snapshot_required\""), "Compatibility repair plan must expose snapshot requirements")
assert(repair_plan_source.include?("\"rollback_available\" => true"), "Compatibility repair plan must keep rollback available")
assert(repair_plan_source.include?("\"c_runtime_backed\" => true"), "Compatibility repair plan must expose C Runtime backing")
assert(repair_plan_source.include?("GetRepairPlan"), "Compatibility repair plan must expose the Runtime method")
assert(repair_plan_source.include?("\"repair_executed\" => false"), "Compatibility repair plan must not execute repairs")
assert(repair_plan_source.include?("\"backend_launch_enabled\" => false"), "Compatibility repair plan must not enable backend launch")
assert(repair_plan_source.include?("\"host_root_modified\" => false"), "Compatibility repair plan must not mutate the host root")
assert(repair_plan_source.include?("\"backend_details_exposed\" => false"), "Compatibility repair plan must hide backend details")
assert(repair_plan_source.include?("CompatibilitySnapshotPlan"), "Compatibility repair plan must include snapshot planning")

snapshot_plan_source = read_project_file("lib/xnix/compatibility/compatibility_snapshot_plan.rb")
assert(snapshot_plan_source.include?("xnix-compat-snapshot-plan"), "Compatibility snapshot plan must expose a CLI command")
%w[before-repair before-engine-change manual].each do |reason|
  assert(snapshot_plan_source.include?("\"#{reason}\""), "Compatibility snapshot plan must include #{reason}")
end
assert(snapshot_plan_source.include?("\"host_system\" => false"), "Compatibility snapshot plan must not snapshot the host system")
assert(snapshot_plan_source.include?("\"user_documents\" => false"), "Compatibility snapshot plan must not snapshot user documents")
assert(snapshot_plan_source.include?("\"preserve_user_documents\" => true"), "Compatibility snapshot restore must preserve user documents")
assert(snapshot_plan_source.include?("\"c_runtime_backed\" => true"), "Compatibility snapshot plan must expose C Runtime backing")
assert(snapshot_plan_source.include?("GetSnapshotPlan"), "Compatibility snapshot plan must expose the Runtime method")
assert(snapshot_plan_source.include?("\"snapshot_created\" => false"), "Compatibility snapshot plan must not create snapshots")
assert(snapshot_plan_source.include?("\"restore_executed\" => false"), "Compatibility snapshot plan must not execute restore")
assert(snapshot_plan_source.include?("\"host_root_modified\" => false"), "Compatibility snapshot plan must not mutate the host root")
assert(snapshot_plan_source.include?("\"backend_details_exposed\" => false"), "Compatibility snapshot plan must hide backend details")

test_plan_source = read_project_file("lib/xnix/compatibility/compatibility_test_plan.rb")
assert(test_plan_source.include?("xnix-compat-test-plan"), "Compatibility test plan must expose a CLI command")
%w[compatibility-test portal-preflight snapshot-preflight runtime-launch-binding].each do |token|
  assert(test_plan_source.include?(token), "Compatibility test plan must include #{token}")
end
assert(test_plan_source.include?("\"runtime_method\" => \"GetTestPlan\""), "Compatibility test plan must expose the Runtime method")
assert(test_plan_source.include?("\"c_runtime_backed\" => true"), "Compatibility test plan must expose C Runtime backing")
assert(test_plan_source.include?("\"execution_request_created\" => false"), "Compatibility test plan must not create execution requests")
assert(test_plan_source.include?("\"test_executed\" => false"), "Compatibility test plan must not execute tests")
assert(test_plan_source.include?("\"host_root_modified\" => false"), "Compatibility test plan must not mutate the host root")
assert(test_plan_source.include?("\"backend_details_exposed\" => false"), "Compatibility test plan must hide backend details")

test_result_source = read_project_file("lib/xnix/compatibility/compatibility_test_result.rb")
assert(test_result_source.include?("xnix-compat-test-result"), "Compatibility test result must expose a CLI command")
%w[compatibility-test-result waiting-for-runtime safe_for_ai_diagnostics runtime-launch-binding].each do |token|
  assert(test_result_source.include?(token), "Compatibility test result must include #{token}")
end
assert(test_result_source.include?("\"runtime_method\" => \"GetTestResult\""), "Compatibility test result must expose the Runtime method")
assert(test_result_source.include?("\"c_runtime_backed\" => true"), "Compatibility test result must expose C Runtime backing")
assert(test_result_source.include?("\"test_executed\" => false"), "Compatibility test result must not execute tests")
assert(test_result_source.include?("\"host_root_modified\" => false"), "Compatibility test result must not mutate the host root")
assert(test_result_source.include?("\"backend_details_exposed\" => false"), "Compatibility test result must hide backend details")

execution_readiness_source = read_project_file("lib/xnix/compatibility/compatibility_execution_readiness.rb")
assert(execution_readiness_source.include?("xnix-compat-execution-readiness"), "Compatibility execution readiness must expose a CLI command")
%w[compatibility-execution-readiness GetExecutionReadiness runtime-launch-write-gate desktop_entry_launch_visible].each do |token|
  assert(execution_readiness_source.include?(token), "Compatibility execution readiness must include #{token}")
end
assert(execution_readiness_source.include?("\"c_runtime_backed\" => true"), "Compatibility execution readiness must expose C Runtime backing")
assert(execution_readiness_source.include?("\"launch_allowed\" => launch_allowed?"), "Compatibility execution readiness must expose launch allowance")
assert(execution_readiness_source.include?("\"execution_request_created\" => false"), "Compatibility execution readiness must not create execution requests")
assert(execution_readiness_source.include?("\"host_root_modified\" => false"), "Compatibility execution readiness must not mutate the host root")
assert(execution_readiness_source.include?("\"backend_details_exposed\" => false"), "Compatibility execution readiness must hide backend details")

ai_diagnostic_source = read_project_file("lib/xnix/compatibility/ai_diagnostic_input.rb")
assert(ai_diagnostic_source.include?("xnix-ai-diagnostic-input"), "AI diagnostic input must expose a CLI command")
%w[ai-diagnostic-input diagnostic_signals privacy_boundaries allowed_ai_tasks blocked_ai_tasks].each do |token|
  assert(ai_diagnostic_source.include?(token), "AI diagnostic input must include #{token}")
end
assert(ai_diagnostic_source.include?("\"ai_provider_called\" => false"), "AI diagnostic input must not call an AI provider")
assert(ai_diagnostic_source.include?("\"network_required\" => false"), "AI diagnostic input must not require network access")
assert(ai_diagnostic_source.include?("\"c_runtime_backed\" => true"), "AI diagnostic input must expose C Runtime backing")
assert(ai_diagnostic_source.include?("\"backend_details_exposed\" => false"), "AI diagnostic input must hide backend details")

ai_recommendation_source = read_project_file("lib/xnix/compatibility/ai_diagnostic_recommendation.rb")
assert(ai_recommendation_source.include?("xnix-ai-diagnostic-recommendation"), "AI diagnostic recommendation must expose a CLI command")
%w[ai-diagnostic-recommendation recommendations approval_required_actions blocked_actions].each do |token|
  assert(ai_recommendation_source.include?(token), "AI diagnostic recommendation must include #{token}")
end
assert(ai_recommendation_source.include?("\"ai_provider_called\" => false"), "AI diagnostic recommendation must not call an AI provider")
assert(ai_recommendation_source.include?("\"network_required\" => false"), "AI diagnostic recommendation must not require network access")
assert(ai_recommendation_source.include?("\"auto_execute\" => false"), "AI diagnostic recommendation must not auto-execute recommendations")
assert(ai_recommendation_source.include?("\"auto_execution_allowed\" => false"), "AI diagnostic recommendation must not allow automatic execution")
assert(ai_recommendation_source.include?("\"c_runtime_backed\" => true"), "AI diagnostic recommendation must expose C Runtime backing")
assert(ai_recommendation_source.include?("\"backend_details_exposed\" => false"), "AI diagnostic recommendation must hide backend details")

ai_repair_gate_source = read_project_file("lib/xnix/compatibility/ai_repair_approval_gate.rb")
assert(ai_repair_gate_source.include?("xnix-ai-repair-approval-gate"), "AI repair approval gate must expose a CLI command")
%w[ai-repair-approval-gate blocked-until-approval required_gates approval_required_actions].each do |token|
  assert(ai_repair_gate_source.include?(token), "AI repair approval gate must include #{token}")
end
assert(ai_repair_gate_source.include?("\"repair_executed\" => false"), "AI repair approval gate must not execute repairs")
assert(ai_repair_gate_source.include?("\"auto_execution_allowed\" => false"), "AI repair approval gate must not permit automatic repair execution")
assert(ai_repair_gate_source.include?("\"c_runtime_backed\" => true"), "AI repair approval gate must expose C Runtime backing")
assert(ai_repair_gate_source.include?("GetAIRepairApprovalGate"), "AI repair approval gate must expose the Runtime method")
assert(ai_repair_gate_source.include?("\"backend_details_exposed\" => false"), "AI repair approval gate must hide backend details")

c_runtime_core_header = read_project_file("runtime/core/xnix_runtime_core.h")
c_runtime_core_source = read_project_file("runtime/core/xnix_runtime_core.c") +
                        read_project_file("runtime/core/xnix_runtime_core_action_review.inc") +
                        read_project_file("runtime/core/xnix_runtime_core_compatibility_center.inc") +
                        read_project_file("runtime/core/xnix_runtime_core_run_plan.inc") +
                        read_project_file("runtime/core/xnix_runtime_core_desktop_activation.inc") +
                        read_project_file("runtime/core/xnix_runtime_core_kde_integration_status.inc") +
                        read_project_file("runtime/core/xnix_runtime_core_kde_shell_integration_plan.inc") +
                        read_project_file("runtime/core/xnix_runtime_core_kde_application_surface_plan.inc") +
                        read_project_file("runtime/core/xnix_runtime_core_desktop_resource_bridge_plan.inc") +
                        read_project_file("runtime/core/xnix_runtime_core_compatibility_mode_switch_plan.inc") +
                        read_project_file("runtime/core/xnix_runtime_core_compatibility_permission_review_plan.inc") +
                        read_project_file("runtime/core/xnix_runtime_core_compatibility_review_flow_plan.inc") +
                        read_project_file("runtime/core/xnix_runtime_core_desktop_entry.inc") +
                        read_project_file("runtime/core/xnix_runtime_core_task_manager_identity.inc") +
                        read_project_file("runtime/core/xnix_runtime_core_kwin_window_rule.inc") +
                        read_project_file("runtime/core/xnix_runtime_core_file_association.inc") +
                        read_project_file("runtime/core/xnix_runtime_core_notification.inc") +
                        read_project_file("runtime/core/xnix_runtime_core_tray_status.inc") +
                        read_project_file("runtime/core/xnix_runtime_core_krunner_query.inc") +
                        read_project_file("runtime/core/xnix_runtime_core_portal_request.inc") +
                        read_project_file("runtime/core/xnix_runtime_core_compatibility_snapshot_plan.inc") +
                        read_project_file("runtime/core/xnix_runtime_core_compatibility_test_plan.inc") +
                     read_project_file("runtime/core/xnix_runtime_core_compatibility_test_result.inc") +
                     read_project_file("runtime/core/xnix_runtime_core_execution_readiness.inc") +
                     read_project_file("runtime/core/xnix_runtime_core_launch_intent.inc") +
                     read_project_file("runtime/core/xnix_runtime_core_backend_capability_matrix.inc") +
                     read_project_file("runtime/core/xnix_runtime_core_backend_selection_plan.inc") +
                     read_project_file("runtime/core/xnix_runtime_core_backend_lifecycle.inc") +
                     read_project_file("runtime/core/xnix_runtime_core_backend_environment_plan.inc") +
                     read_project_file("runtime/core/xnix_runtime_core_compatibility_repair_plan.inc") +
                        read_project_file("runtime/core/xnix_runtime_core_ai_diagnostic_input.inc") +
                        read_project_file("runtime/core/xnix_runtime_core_ai_diagnostic_recommendation.inc") +
                        read_project_file("runtime/core/xnix_runtime_core_ai_repair_approval_gate.inc")
c_runtime_core_cli = read_project_file("runtime/core/xnix_runtime_core_cli.c") +
                     read_project_file("runtime/core/xnix_runtime_core_cli_probe.inc") +
                     read_project_file("runtime/core/xnix_runtime_core_cli_action_queue.inc") +
                     read_project_file("runtime/core/xnix_runtime_core_cli_action_review.inc") +
                     read_project_file("runtime/core/xnix_runtime_core_cli_compatibility_center.inc") +
                     read_project_file("runtime/core/xnix_runtime_core_cli_run_plan.inc") +
                     read_project_file("runtime/core/xnix_runtime_core_cli_desktop_activation.inc") +
                     read_project_file("runtime/core/xnix_runtime_core_cli_kde_integration_status.inc") +
                     read_project_file("runtime/core/xnix_runtime_core_cli_kde_shell_integration_plan.inc") +
                     read_project_file("runtime/core/xnix_runtime_core_cli_kde_application_surface_plan.inc") +
                     read_project_file("runtime/core/xnix_runtime_core_cli_desktop_resource_bridge_plan.inc") +
                     read_project_file("runtime/core/xnix_runtime_core_cli_compatibility_mode_switch_plan.inc") +
                     read_project_file("runtime/core/xnix_runtime_core_cli_compatibility_permission_review_plan.inc") +
                     read_project_file("runtime/core/xnix_runtime_core_cli_compatibility_review_flow_plan.inc") +
                     read_project_file("runtime/core/xnix_runtime_core_cli_desktop_entry.inc") +
                     read_project_file("runtime/core/xnix_runtime_core_cli_task_manager_identity.inc") +
                     read_project_file("runtime/core/xnix_runtime_core_cli_kwin_window_rule.inc") +
                     read_project_file("runtime/core/xnix_runtime_core_cli_file_association.inc") +
                     read_project_file("runtime/core/xnix_runtime_core_cli_notification.inc") +
                     read_project_file("runtime/core/xnix_runtime_core_cli_tray_status.inc") +
                     read_project_file("runtime/core/xnix_runtime_core_cli_krunner_query.inc") +
                     read_project_file("runtime/core/xnix_runtime_core_cli_portal_request.inc") +
                     read_project_file("runtime/core/xnix_runtime_core_cli_compatibility_snapshot_plan.inc") +
                     read_project_file("runtime/core/xnix_runtime_core_cli_compatibility_test_plan.inc") +
                     read_project_file("runtime/core/xnix_runtime_core_cli_compatibility_test_result.inc") +
                     read_project_file("runtime/core/xnix_runtime_core_cli_execution_readiness.inc") +
                     read_project_file("runtime/core/xnix_runtime_core_cli_launch_intent.inc") +
                     read_project_file("runtime/core/xnix_runtime_core_cli_backend_capability_matrix.inc") +
                     read_project_file("runtime/core/xnix_runtime_core_cli_backend_selection_plan.inc") +
                     read_project_file("runtime/core/xnix_runtime_core_cli_backend_lifecycle.inc") +
                     read_project_file("runtime/core/xnix_runtime_core_cli_backend_environment_plan.inc") +
                     read_project_file("runtime/core/xnix_runtime_core_cli_compatibility_repair_plan.inc") +
                     read_project_file("runtime/core/xnix_runtime_core_cli_ai_diagnostic_input.inc") +
                     read_project_file("runtime/core/xnix_runtime_core_cli_ai_diagnostic_recommendation.inc") +
                     read_project_file("runtime/core/xnix_runtime_core_cli_ai_repair_approval_gate.inc") +
                     read_project_file("runtime/core/xnix_runtime_core_cli_dispatch.inc")
assert(c_runtime_core_header.include?("#define XNIX_RUNTIME_VERSION \"#{EXPECTED_VERSION}\""), "C Runtime core must expose #{EXPECTED_VERSION}")
%w[XNIX_RUNTIME_BUS_NAME XNIX_RUNTIME_OBJECT_PATH XNIX_RUNTIME_INTERFACE XNIX_RUNTIME_WRITE_ERROR].each do |token|
  assert(c_runtime_core_header.include?(token), "C Runtime core header must include #{token}")
end
%w[XnixRuntimeApplication xnix_runtime_application_count xnix_runtime_find_application XnixRuntimeEngine xnix_runtime_engine_count xnix_runtime_select_engine_for_mode XnixRuntimePortalPolicy xnix_runtime_portal_policy_count xnix_runtime_find_portal_policy XnixRuntimePortalRequestPlan xnix_runtime_portal_request_plan XnixRuntimeSnapshotPolicy xnix_runtime_snapshot_policy_count xnix_runtime_find_snapshot_policy XnixRuntimeStateRootPolicy xnix_runtime_state_root_policy_count xnix_runtime_find_state_root_policy XnixRuntimeInstallReadinessPolicy xnix_runtime_install_readiness_policy_count xnix_runtime_find_install_readiness_policy XnixRuntimeArtifactManifestPolicy xnix_runtime_artifact_manifest_policy_count xnix_runtime_find_artifact_manifest_policy XnixRuntimeAcquisitionPreflightPolicy xnix_runtime_acquisition_preflight_policy_count xnix_runtime_find_acquisition_preflight_policy XnixRuntimePackageSourcePolicy xnix_runtime_package_source_policy_count xnix_runtime_find_package_source_policy XnixRuntimeBackendBindingPolicy xnix_runtime_backend_binding_policy_count xnix_runtime_find_backend_binding_policy XnixRuntimeSettingsPolicy xnix_runtime_settings_policy_count xnix_runtime_find_settings_policy XnixRuntimeSettingsChangePolicy xnix_runtime_settings_change_policy_count xnix_runtime_find_settings_change_policy XnixRuntimeServiceBindingPolicy xnix_runtime_service_binding_policy XnixRuntimeLiveOwnerGatePolicy xnix_runtime_live_owner_gate_policy XnixRuntimeOwnerSmokePlanPolicy xnix_runtime_owner_smoke_plan_policy XnixRuntimeMethodParityManifestPolicy xnix_runtime_method_parity_manifest_policy XnixRuntimeRecipeTrustPolicy xnix_runtime_recipe_trust_policy XnixRuntimeRecipeInstallGate xnix_runtime_recipe_install_gate XnixRuntimeCompatibilityInstallPlan xnix_runtime_compatibility_install_plan XnixRuntimeCompatibilityActionQueue xnix_runtime_compatibility_action_queue XnixRuntimeCompatibilityActionReviewReceipt xnix_runtime_compatibility_action_review_receipt XnixRuntimeCompatibilityCenterSummary xnix_runtime_compatibility_center_summary XnixRuntimeCompatibilityRunPlan xnix_runtime_compatibility_run_plan XnixRuntimeDesktopActivationManifest xnix_runtime_desktop_activation_manifest XnixRuntimeKDEIntegrationStatus xnix_runtime_kde_integration_status XnixRuntimeKDEApplicationSurfacePlan XnixRuntimeKDEApplicationSurfaceEntryPoint xnix_runtime_kde_application_surface_plan XnixRuntimeDesktopEntryPlan xnix_runtime_desktop_entry_plan XnixRuntimeTaskManagerIdentityPlan xnix_runtime_task_manager_identity_plan XnixRuntimeKWinWindowRulePlan xnix_runtime_kwin_window_rule_plan XnixRuntimeFileAssociationPlan xnix_runtime_file_association_plan XnixRuntimeNotificationPlan xnix_runtime_notification_plan XnixRuntimeTrayStatusPlan xnix_runtime_tray_status_plan XnixRuntimeKRunnerMatch XnixRuntimeKRunnerQueryPlan xnix_runtime_krunner_query_plan XnixRuntimeAIDiagnosticInput xnix_runtime_ai_diagnostic_input].each do |token|
  assert(c_runtime_core_header.include?(token), "C Runtime core header must include #{token}")
end
%w[XnixRuntimeAIDiagnosticRecommendation xnix_runtime_ai_diagnostic_recommendation].each do |token|
  assert(c_runtime_core_header.include?(token), "C Runtime core header must include #{token}")
end
%w[XnixRuntimeAIRepairApprovalGate XnixRuntimeAIRepairRequiredGate xnix_runtime_ai_repair_approval_gate].each do |token|
  assert(c_runtime_core_header.include?(token), "C Runtime core header must include #{token}")
end
%w[XnixRuntimeCompatibilityRepairPlan XnixRuntimeCompatibilityRepairAction xnix_runtime_compatibility_repair_plan].each do |token|
  assert(c_runtime_core_header.include?(token), "C Runtime core header must include #{token}")
end
%w[XnixRuntimeCompatibilitySnapshotPlan xnix_runtime_compatibility_snapshot_plan].each do |token|
  assert(c_runtime_core_header.include?(token), "C Runtime core header must include #{token}")
end
%w[XnixRuntimeCompatibilityTestPlan XnixRuntimeCompatibilityTestStep xnix_runtime_compatibility_test_plan].each do |token|
  assert(c_runtime_core_header.include?(token), "C Runtime core header must include #{token}")
end
%w[XnixRuntimeCompatibilityTestResult XnixRuntimeCompatibilityTestStepResult xnix_runtime_compatibility_test_result].each do |token|
  assert(c_runtime_core_header.include?(token), "C Runtime core header must include #{token}")
end
%w[XnixRuntimeExecutionReadiness XnixRuntimeExecutionReadinessGate xnix_runtime_execution_readiness].each do |token|
  assert(c_runtime_core_header.include?(token), "C Runtime core header must include #{token}")
end
%w[XnixRuntimeBackendLifecycle XnixRuntimeBackendLifecycleStage xnix_runtime_backend_lifecycle].each do |token|
  assert(c_runtime_core_header.include?(token), "C Runtime core header must include #{token}")
end
%w[XnixRuntimeBackendEnvironmentPlan XnixRuntimeBackendEnvironmentProfile xnix_runtime_backend_environment_plan].each do |token|
  assert(c_runtime_core_header.include?(token), "C Runtime core header must include #{token}")
end
%w[XnixRuntimeBackendCapabilityMatrix XnixRuntimeBackendCapabilityProfile XnixRuntimeBackendCapability xnix_runtime_backend_capability_matrix].each do |token|
  assert(c_runtime_core_header.include?(token), "C Runtime core header must include #{token}")
end
%w[XnixRuntimeBackendSelectionPlan XnixRuntimeBackendSelectionCandidate xnix_runtime_backend_selection_plan].each do |token|
  assert(c_runtime_core_header.include?(token), "C Runtime core header must include #{token}")
end
%w[XnixRuntimeKDEShellIntegrationPlan XnixRuntimeKDEShellComponent xnix_runtime_kde_shell_integration_plan].each do |token|
  assert(c_runtime_core_header.include?(token), "C Runtime core header must include #{token}")
end
%w[XnixRuntimeCompatibilityModeSwitchPlan XnixRuntimeCompatibilityModeOption xnix_runtime_compatibility_mode_switch_plan].each do |token|
  assert(c_runtime_core_header.include?(token), "C Runtime core header must include #{token}")
end
%w[XnixRuntimeCompatibilityPermissionReviewPlan XnixRuntimeCompatibilityPermission xnix_runtime_compatibility_permission_review_plan].each do |token|
  assert(c_runtime_core_header.include?(token), "C Runtime core header must include #{token}")
end
%w[XnixRuntimeCompatibilityReviewFlowPlan XnixRuntimeCompatibilityReviewFlowStep xnix_runtime_compatibility_review_flow_plan].each do |token|
  assert(c_runtime_core_header.include?(token), "C Runtime core header must include #{token}")
end
%w[xnix_runtime_write_gate blocked-until-production-backend InstallRecipe Launch CreateSnapshot RestoreSnapshot xnix_runtime_find_application org.xnix.sample.notepad application/x-xnix-txt xnix_runtime_select_engine_for_mode automatic-managed local-compatibility-engine isolated-compatibility-engine xnix_runtime_find_portal_policy org.freedesktop.portal.FileChooser org.freedesktop.portal.Camera remote-desktop selected-files xnix_runtime_portal_request_plan portal_request_method_for_operation handle_token permission_granted host_permission_changed OpenFile RequestClipboard AccessCamera xnix_runtime_find_snapshot_policy before-repair before-engine-change desktop_activation_receipts xnix_runtime_find_state_root_policy application-data runtime-metadata diagnostic-cache xnix_runtime_find_install_readiness_policy resolve-artifact-manifest verify-artifact-digests prepare-package-source allocate-application-state xnix_runtime_find_artifact_manifest_policy runtime-launch-metadata local-execution-artifacts isolated-environment-artifacts manifest-signature-verification artifact-digest-verification xnix_runtime_find_acquisition_preflight_policy package-source-ready signed-artifact-manifest runtime-cache-space network-policy-review rollback-marker xnix_runtime_find_package_source_policy os-managed-compatibility-packages runtime-managed-toolcache isolated-environment-template-catalog signed-source-verification source-policy-review runtime-cache-quota offline-fallback xnix_runtime_find_backend_binding_policy engine-package-source application-state-root portal-policy-review snapshot-baseline xnix_runtime_find_settings_policy run-mode resource-access devices snapshots settings_change_policies xnix_runtime_find_settings_change_policy validate-setting persist-runtime-setting service_binding_policy xnix_runtime_service_binding_policy dbus-service-activation live-dbus-owner live_owner_gate_policy xnix_runtime_live_owner_gate_policy production-recipe-trust owner_smoke_plan_policy xnix_runtime_owner_smoke_plan_policy validate-activation-files assert-stable-bus-name method_parity_manifest_policy xnix_runtime_method_parity_manifest_policy recipe_trust_policy xnix_runtime_recipe_trust_policy production_recipe_install_requirements xnix_runtime_compatibility_install_plan xnix_runtime_compatibility_action_queue xnix_runtime_compatibility_action_review_receipt xnix_runtime_compatibility_center_summary compatibility-center-summary known_issue_count repair_record_state backend_launch_enabled xnix_runtime_compatibility_run_plan xnix_runtime_desktop_activation_manifest xnix_runtime_kde_integration_status kde-integration-status standard-desktop-entry window-identity-and-restore xnix_runtime_desktop_entry_plan xnix_runtime_task_manager_identity_plan xnix_runtime_file_association_plan xnix_runtime_notification_plan xnix_runtime_tray_status_plan xnix_runtime_krunner_query_plan xnix_runtime_ai_diagnostic_input ai-diagnostic-input safe_for_ai_diagnostics desktop-activation desktop-entry-plan task-manager-identity-plan file-association-plan notification-plan tray-status-plan krunner-query-plan query_execution_enabled approval-required install-failed repair-applied mode-changed action_execution_enabled repair_execution_enabled settings_persistence_enabled live_backend_bridge_enabled bridge_configuration_persisted mimeapps.list applications:xnix-org.xnix.sample.notepad.desktop xnix-compat-open standard_desktop_entry launch_uses_runtime raw_windows_executable_exposed compatibility_storage_path_exposed pinning_allowed restore_allowed skip_taskbar show_in_switcher identity-only window_manager_policy_only portal_required_for_file_open overwrite_existing_mimeapps portal_policy_required snapshot_before_risky_change diagnostics_required execution_request_created review-install-readiness review-settings-change review-ai-repair verify-runtime-service review-portal-policy decision_recorded reviewed approved deferred rejected signed recipe validation is not enabled registry.digest registry.signature registry.development GetDesktopActivationManifest GetKDEIntegrationStatus GetDesktopEntryPlan GetTaskManagerIdentityPlan GetFileAssociationPlan GetNotificationPlan GetTrayStatus GetKRunnerQueryPlan GetCompatibilityCenterSummary GetPortalRequestPlan GetAIDiagnosticInput GetRuntimeMethodParityManifest GetRuntimeWriteGate].each do |token|
  assert(c_runtime_core_source.include?(token), "C Runtime core source must include #{token}")
end
assert(c_runtime_core_source.include?("read user documents"), "C Runtime core source must block AI user-document reads")
%w[xnix_runtime_ai_diagnostic_recommendation ai-diagnostic-recommendation auto_execute approval_surface GetAIDiagnosticRecommendation].each do |token|
  assert(c_runtime_core_source.include?(token), "C Runtime core source must include #{token}")
end
%w[xnix_runtime_ai_repair_approval_gate ai-repair-approval-gate blocked-until-approval runtime-approval-token restore-point-preflight GetAIRepairApprovalGate].each do |token|
  assert(c_runtime_core_source.include?(token), "C Runtime core source must include #{token}")
end
%w[xnix_runtime_compatibility_repair_plan compatibility-repair engine-binding-pending portal-approval-required recipe-trust-blocked runtime-repair-applied repair_execution_requested backend_launch_enabled GetRepairPlan].each do |token|
  assert(c_runtime_core_source.include?(token), "C Runtime core source must include #{token}")
end
%w[xnix_runtime_compatibility_snapshot_plan snapshot_request_created snapshot_created restore_requested restore_executed].each do |token|
  assert(c_runtime_core_source.include?(token), "C Runtime core source must include #{token}")
end
%w[xnix_runtime_compatibility_test_plan compatibility-test repair-readiness portal-preflight snapshot-preflight runtime-launch-binding test_executed GetTestPlan].each do |token|
  assert(c_runtime_core_source.include?(token), "C Runtime core source must include #{token}")
end
%w[xnix_runtime_compatibility_test_result compatibility-test-result waiting-for-runtime safe_for_ai_diagnostics GetTestResult].each do |token|
  assert(c_runtime_core_source.include?(token), "C Runtime core source must include #{token}")
end
%w[xnix_runtime_execution_readiness compatibility-execution-readiness GetExecutionReadiness runtime-launch-write-gate desktop_entry_launch_visible launch_allowed].each do |token|
  assert(c_runtime_core_source.include?(token), "C Runtime core source must include #{token}")
end
%w[xnix_runtime_launch_intent runtime-launch-intent GetLaunchIntent desktop-launcher write_gate_decision execution_started].each do |token|
  assert(c_runtime_core_source.include?(token), "C Runtime core source must include #{token}")
end
%w[xnix_runtime_backend_lifecycle compatibility-backend-lifecycle GetBackendLifecycle backend_process_started local_backend_started isolated_backend_started].each do |token|
  assert(c_runtime_core_source.include?(token), "C Runtime core source must include #{token}")
end
%w[xnix_runtime_backend_environment_plan compatibility-backend-environment-plan GetBackendEnvironmentPlan local-compatibility-environment isolated-compatibility-environment clipboard_bridge_enabled print_bridge_enabled].each do |token|
  assert(c_runtime_core_source.include?(token), "C Runtime core source must include #{token}")
end
%w[xnix_runtime_backend_capability_matrix compatibility-backend-capability-matrix GetBackendCapabilityMatrix local-compatibility isolated-compatibility application-launch file-bridge clipboard-bridge print-bridge snapshot-restore capability_activation_enabled].each do |token|
  assert(c_runtime_core_source.include?(token), "C Runtime core source must include #{token}")
end
%w[xnix_runtime_backend_selection_plan compatibility-backend-selection-plan GetBackendSelectionPlan recommended_profile_id selection_committed selection_change_enabled backend-capability-review backend-binding-review].each do |token|
  assert(c_runtime_core_source.include?(token), "C Runtime core source must include #{token}")
end
%w[xnix_runtime_kde_application_surface_plan kde-application-surface-plan GetKDEApplicationSurfacePlan normal_linux_application_surface standard_launcher_visible runtime-write-gate].each do |token|
  assert(c_runtime_core_source.include?(token), "C Runtime core source must include #{token}")
end
%w[xnix_runtime_kde_shell_integration_plan kde-shell-integration-plan GetKDEShellIntegrationPlan start-menu task-manager system-tray notification-center compatibility-center unified-settings plasma_fork_required shell_configuration_written component_activation_enabled].each do |token|
  assert(c_runtime_core_source.include?(token), "C Runtime core source must include #{token}")
end
%w[XnixRuntimeDesktopResourceBridgePlan XnixRuntimeDesktopResourceBridge xnix_runtime_desktop_resource_bridge_plan].each do |token|
  assert(c_runtime_core_header.include?(token), "C Runtime core header must include #{token}")
end
%w[xnix_runtime_desktop_resource_bridge_plan desktop-resource-bridge-plan GetDesktopResourceBridgePlan portal_mediated file-open clipboard screenshot direct_host_file_access].each do |token|
  assert(c_runtime_core_source.include?(token), "C Runtime core source must include #{token}")
end
%w[xnix_runtime_compatibility_mode_switch_plan compatibility-mode-switch-plan GetCompatibilityModeSwitchPlan prefer-performance prefer-compatibility isolated-execution settings-review backend_reconfiguration_enabled].each do |token|
  assert(c_runtime_core_source.include?(token), "C Runtime core source must include #{token}")
end
%w[xnix_runtime_compatibility_permission_review_plan compatibility-permission-review-plan GetCompatibilityPermissionReviewPlan documents downloads camera network clipboard permission_changes_applied].each do |token|
  assert(c_runtime_core_source.include?(token), "C Runtime core source must include #{token}")
end
%w[xnix_runtime_compatibility_review_flow_plan compatibility-review-flow-plan GetCompatibilityReviewFlowPlan settings-change-review permission-review portal-request-review runtime-write-gate review-receipt apply_enabled].each do |token|
  assert(c_runtime_core_source.include?(token), "C Runtime core source must include #{token}")
end
%w[xnix_runtime_kwin_window_rule_plan kwin-window-rule identity-and-layout GetKWinWindowRulePlan].each do |token|
  assert(c_runtime_core_source.include?(token), "C Runtime core source must include #{token}")
end
%w[business_logic_runtime tests-and-development-tools application_catalog_owner engine_catalog_owner portal_policy_owner portal_request_plan_owner ai_diagnostic_input_owner ai_diagnostic_signal_count desktop_entry_plan_owner task_manager_identity_plan_owner kde_integration_status_owner kde_integration_entry_point_count file_association_plan_owner notification_plan_owner tray_status_plan_owner krunner_query_plan_owner krunner_query_plan_count compatibility_center_summary_owner snapshot_policy_owner state_root_policy_owner install_readiness_policy_owner artifact_manifest_policy_owner acquisition_preflight_policy_owner package_source_policy_owner backend_binding_policy_owner backend_lifecycle_owner backend_environment_plan_owner settings_policy_owner settings_change_policy_owner service_binding_policy_owner live_owner_gate_policy_owner owner_smoke_plan_policy_owner method_parity_manifest_policy_owner recipe_trust_policy_owner recipe_install_gate_owner compatibility_install_plan_owner compatibility_action_queue_owner compatibility_action_review_receipt_owner compatibility_center_summary_count compatibility_run_plan_owner desktop_activation_manifest_owner compatibility-center-action-review-receipt compatibility-center-summary receipt_id required_runtime_gate compatibility-run desktop-activation kde-integration-status portal-request-plan ai-diagnostic-input desktop-entry-plan task-manager-identity-plan file-association-plan notification-plan tray-status-plan krunner-query-plan list-applications get-application list-engines select-engine list-portal-policies portal-policy portal-request-plan list-snapshot-policies snapshot-policy list-state-root-policies state-root-policy list-install-readiness-policies install-readiness-policy compatibility-install-plan compatibility-action-queue compatibility-action-review compatibility-center-summary compatibility-run-plan desktop-activation-manifest kde-integration-status desktop-entry-plan task-manager-identity-plan file-association-plan notification-plan tray-status-plan krunner-query-plan ai-diagnostic-input list-artifact-manifest-policies artifact-manifest-policy list-acquisition-preflight-policies acquisition-preflight-policy list-package-source-policies package-source-policy list-backend-binding-policies backend-binding-policy backend-lifecycle backend-environment-plan list-settings-policies settings-policy list-settings-change-policies settings-change-policy runtime-service-binding runtime-live-owner-gate runtime-owner-smoke-plan runtime-method-parity-manifest recipe-trust-policy recipe-install-gate write-gate].each do |token|
  assert(c_runtime_core_cli.include?(token), "C Runtime core CLI must include #{token}")
end
%w[compatibility_mode_switch_plan_owner compatibility_mode_switch_mode_count compatibility-mode-switch-plan GetCompatibilityModeSwitchPlan backend_reconfiguration_enabled].each do |token|
  assert(c_runtime_core_cli.include?(token), "C Runtime core CLI must include #{token}")
end
%w[compatibility_permission_review_plan_owner compatibility_permission_review_count compatibility-permission-review-plan GetCompatibilityPermissionReviewPlan permissions_granted host_permission_changed].each do |token|
  assert(c_runtime_core_cli.include?(token), "C Runtime core CLI must include #{token}")
end
%w[compatibility_review_flow_plan_owner compatibility_review_flow_step_count compatibility-review-flow-plan GetCompatibilityReviewFlowPlan request_object_created execution_started].each do |token|
  assert(c_runtime_core_cli.include?(token), "C Runtime core CLI must include #{token}")
end
%w[kde_shell_integration_plan_owner kde_shell_component_count kde-shell-integration-plan GetKDEShellIntegrationPlan plasma_fork_required shell_configuration_written component_activation_enabled].each do |token|
  assert(c_runtime_core_cli.include?(token), "C Runtime core CLI must include #{token}")
end
%w[ai_diagnostic_recommendation_owner ai_diagnostic_recommendation_count ai-diagnostic-recommendation].each do |token|
  assert(c_runtime_core_cli.include?(token), "C Runtime core CLI must include #{token}")
end
%w[ai_repair_approval_gate_owner ai_repair_required_gate_count ai-repair-approval-gate].each do |token|
  assert(c_runtime_core_cli.include?(token), "C Runtime core CLI must include #{token}")
end
%w[compatibility_repair_plan_owner compatibility_repair_issue_count compatibility-repair-plan].each do |token|
  assert(c_runtime_core_cli.include?(token), "C Runtime core CLI must include #{token}")
end
%w[compatibility_snapshot_plan_owner compatibility_snapshot_reason_count compatibility-snapshot-plan compatibility-snapshot GetSnapshotPlan].each do |token|
  assert(c_runtime_core_cli.include?(token), "C Runtime core CLI must include #{token}")
end
%w[compatibility_test_plan_owner compatibility_test_type_count compatibility-test-plan compatibility-test GetTestPlan].each do |token|
  assert(c_runtime_core_cli.include?(token), "C Runtime core CLI must include #{token}")
end
%w[compatibility_test_result_owner compatibility_test_result_type_count compatibility-test-result GetTestResult].each do |token|
  assert(c_runtime_core_cli.include?(token), "C Runtime core CLI must include #{token}")
end
%w[execution_readiness_owner execution_readiness_application_count execution-readiness compatibility-execution-readiness GetExecutionReadiness].each do |token|
  assert(c_runtime_core_cli.include?(token), "C Runtime core CLI must include #{token}")
end
%w[launch_intent_owner launch_intent_application_count launch-intent runtime-launch-intent GetLaunchIntent].each do |token|
  assert(c_runtime_core_cli.include?(token), "C Runtime core CLI must include #{token}")
end
%w[backend_lifecycle_owner backend_lifecycle_application_count backend-lifecycle compatibility-backend-lifecycle GetBackendLifecycle].each do |token|
  assert(c_runtime_core_cli.include?(token), "C Runtime core CLI must include #{token}")
end
%w[backend_environment_plan_owner backend_environment_plan_count backend-environment-plan compatibility-backend-environment-plan GetBackendEnvironmentPlan].each do |token|
  assert(c_runtime_core_cli.include?(token), "C Runtime core CLI must include #{token}")
end
%w[backend_capability_matrix_owner backend_capability_profile_count backend_capability_count backend-capability-matrix GetBackendCapabilityMatrix capability_activation_enabled state_root_created snapshots_created].each do |token|
  assert(c_runtime_core_cli.include?(token), "C Runtime core CLI must include #{token}")
end
%w[backend_selection_plan_owner backend_selection_candidate_count backend-selection-plan GetBackendSelectionPlan recommended_profile_id selection_committed selection_change_enabled request_object_created].each do |token|
  assert(c_runtime_core_cli.include?(token), "C Runtime core CLI must include #{token}")
end
%w[kde_application_surface_plan_owner kde_application_surface_entry_point_count kde-application-surface-plan GetKDEApplicationSurfacePlan normal_linux_application_surface standard_launcher_visible].each do |token|
  assert(c_runtime_core_cli.include?(token), "C Runtime core CLI must include #{token}")
end
%w[desktop_resource_bridge_plan_owner desktop_resource_bridge_count desktop-resource-bridge-plan GetDesktopResourceBridgePlan portal_mediated direct_host_file_access].each do |token|
  assert(c_runtime_core_cli.include?(token), "C Runtime core CLI must include #{token}")
end
%w[kwin_window_rule_plan_owner kwin_window_rule_plan_count kwin-window-rule-plan].each do |token|
  assert(c_runtime_core_cli.include?(token), "C Runtime core CLI must include #{token}")
end

application_state_root_source = read_project_file("lib/xnix/compatibility/application_state_root.rb")
assert(application_state_root_source.include?("xnix-compat-state-root"), "Application state root must expose a CLI command")
%w[compatibility-application-state-root managed_scopes portal_required_for_user_files snapshot_eligible].each do |token|
  assert(application_state_root_source.include?(token), "Application state root must include #{token}")
end
assert(application_state_root_source.include?("\"directories_created\" => false"), "Application state root must not create directories during planning")
assert(application_state_root_source.include?("\"host_root_modified\" => false"), "Application state root must not mutate the host root")
assert(application_state_root_source.include?("\"user_documents_included\" => false"), "Application state root must exclude user documents")
assert(application_state_root_source.include?("\"backend_details_exposed\" => false"), "Application state root must hide backend details")

acquisition_preflight_source = read_project_file("lib/xnix/compatibility/compatibility_acquisition_preflight.rb")
assert(acquisition_preflight_source.include?("xnix-compat-acquisition-preflight"), "Compatibility acquisition preflight must expose a CLI command")
%w[compatibility-acquisition-preflight package-source-ready signed-artifact-manifest runtime-cache-space network-policy-review rollback-marker].each do |token|
  assert(acquisition_preflight_source.include?(token), "Compatibility acquisition preflight must include #{token}")
end
assert(acquisition_preflight_source.include?("\"acquisition_ready\" => false"), "Compatibility acquisition preflight must not claim acquisition readiness")
assert(acquisition_preflight_source.include?("\"download_enabled\" => false"), "Compatibility acquisition preflight must not enable downloads")
assert(acquisition_preflight_source.include?("\"install_enabled\" => false"), "Compatibility acquisition preflight must not enable installation")
assert(acquisition_preflight_source.include?("\"network_request_created\" => false"), "Compatibility acquisition preflight must not create network requests")
assert(acquisition_preflight_source.include?("\"artifacts_downloaded\" => false"), "Compatibility acquisition preflight must not download artifacts")
assert(acquisition_preflight_source.include?("\"host_root_modified\" => false"), "Compatibility acquisition preflight must not mutate the host root")
assert(acquisition_preflight_source.include?("\"backend_details_exposed\" => false"), "Compatibility acquisition preflight must hide backend details")

artifact_manifest_source = read_project_file("lib/xnix/compatibility/compatibility_artifact_manifest.rb")
assert(artifact_manifest_source.include?("xnix-compat-artifact-manifest"), "Compatibility artifact manifest must expose a CLI command")
%w[compatibility-artifact-manifest runtime-launch-metadata local-execution-artifacts isolated-environment-artifacts manifest-signature-verification artifact-digest-verification cache-namespace-allocation rollback-reference].each do |token|
  assert(artifact_manifest_source.include?(token), "Compatibility artifact manifest must include #{token}")
end
assert(artifact_manifest_source.include?("\"manifest_ready\" => false"), "Compatibility artifact manifest must not claim manifest readiness")
assert(artifact_manifest_source.include?("\"signature_verified\" => false"), "Compatibility artifact manifest must not claim signature verification")
assert(artifact_manifest_source.include?("\"download_enabled\" => false"), "Compatibility artifact manifest must not enable downloads")
assert(artifact_manifest_source.include?("\"artifacts_downloaded\" => false"), "Compatibility artifact manifest must not download artifacts")
assert(artifact_manifest_source.include?("\"host_root_modified\" => false"), "Compatibility artifact manifest must not mutate the host root")
assert(artifact_manifest_source.include?("\"backend_details_exposed\" => false"), "Compatibility artifact manifest must hide backend details")

install_plan_source = read_project_file("lib/xnix/compatibility/compatibility_install_plan.rb")
assert(install_plan_source.include?("xnix-compat-install-plan"), "Compatibility install plan must expose a CLI command")
%w[compatibility-install-plan resolve-artifact-manifest verify-artifact-digests prepare-package-source allocate-application-state stage-desktop-integration enable-launch-binding].each do |token|
  assert(install_plan_source.include?(token), "Compatibility install plan must include #{token}")
end
assert(install_plan_source.include?("\"install_ready\" => false"), "Compatibility install plan must not claim install readiness")
assert(install_plan_source.include?("\"desktop_activation_ready\" => false"), "Compatibility install plan must not claim desktop activation readiness")
assert(install_plan_source.include?("\"download_enabled\" => false"), "Compatibility install plan must not enable downloads")
assert(install_plan_source.include?("\"install_enabled\" => false"), "Compatibility install plan must not enable installation")
assert(install_plan_source.include?("\"network_request_created\" => false"), "Compatibility install plan must not create network requests")
assert(install_plan_source.include?("\"artifacts_downloaded\" => false"), "Compatibility install plan must not download artifacts")
assert(install_plan_source.include?("\"host_root_modified\" => false"), "Compatibility install plan must not mutate the host root")
assert(install_plan_source.include?("\"backend_details_exposed\" => false"), "Compatibility install plan must hide backend details")

package_source_source = read_project_file("lib/xnix/compatibility/compatibility_package_source.rb")
assert(package_source_source.include?("xnix-compat-package-source"), "Compatibility package source must expose a CLI command")
%w[compatibility-package-source source_channels required_preflight signed_source_required runtime-cache offline-fallback].each do |token|
  assert(package_source_source.include?(token), "Compatibility package source must include #{token}")
end
assert(package_source_source.include?("\"package_source_ready\" => false"), "Compatibility package source must not claim source readiness")
assert(package_source_source.include?("\"install_enabled\" => false"), "Compatibility package source must not enable installation")
assert(package_source_source.include?("\"host_root_modified\" => false"), "Compatibility package source must not mutate the host root")
assert(package_source_source.include?("\"privileged_container_required\" => false"), "Compatibility package source must not require privileged containers")
assert(package_source_source.include?("\"desktop_shell_command_exposed\" => false"), "Compatibility package source must not expose desktop commands")
assert(package_source_source.include?("\"backend_details_exposed\" => false"), "Compatibility package source must hide backend details")

backend_binding_source = read_project_file("lib/xnix/compatibility/compatibility_backend_binding.rb")
assert(backend_binding_source.include?("xnix-compat-backend-binding"), "Compatibility backend binding must expose a CLI command")
%w[compatibility-backend-binding required_preflight managed_binding_ready launch_enabled].each do |token|
  assert(backend_binding_source.include?(token), "Compatibility backend binding must include #{token}")
end
assert(backend_binding_source.include?("\"execution_request_created\" => false"), "Compatibility backend binding must not create execution requests")
assert(backend_binding_source.include?("\"host_root_modified\" => false"), "Compatibility backend binding must not mutate the host root")
assert(backend_binding_source.include?("\"privileged_container_required\" => false"), "Compatibility backend binding must not require privileged containers")
assert(backend_binding_source.include?("\"backend_details_exposed\" => false"), "Compatibility backend binding must hide backend details")

backend_capability_matrix_source = read_project_file("lib/xnix/compatibility/compatibility_backend_capability_matrix.rb")
assert(backend_capability_matrix_source.include?("xnix-compat-backend-capability-matrix"), "Compatibility backend capability matrix must expose a CLI command")
%w[compatibility-backend-capability-matrix GetBackendCapabilityMatrix local-compatibility isolated-compatibility application-launch package-management file-bridge clipboard-bridge print-bridge snapshot-restore diagnostics].each do |token|
  assert(backend_capability_matrix_source.include?(token), "Compatibility backend capability matrix must include #{token}")
end
%w[selection_enabled backend_launch_enabled capability_activation_enabled request_objects_created state_root_created snapshots_created host_root_modified privileged_container_required backend_details_exposed].each do |token|
  assert(backend_capability_matrix_source.include?("\"#{token}\" => false"), "Compatibility backend capability matrix must keep #{token} false while planning")
end

backend_selection_plan_source = read_project_file("lib/xnix/compatibility/compatibility_backend_selection_plan.rb")
assert(backend_selection_plan_source.include?("xnix-compat-backend-selection-plan"), "Compatibility backend selection plan must expose a CLI command")
%w[compatibility-backend-selection-plan GetBackendSelectionPlan local-compatibility isolated-compatibility backend-capability-review backend-binding-review application-state-root-review portal-policy-review snapshot-baseline-review].each do |token|
  assert(backend_selection_plan_source.include?(token), "Compatibility backend selection plan must include #{token}")
end
%w[selection_committed selection_change_enabled backend_launch_enabled capability_activation_enabled environment_created request_object_created state_root_created snapshot_created host_root_modified privileged_container_required backend_details_exposed].each do |token|
  assert(backend_selection_plan_source.include?("\"#{token}\" => false"), "Compatibility backend selection plan must keep #{token} false while planning")
end

runtime_service_binding_source = read_project_file("lib/xnix/compatibility/runtime_service_binding.rb")
assert(runtime_service_binding_source.include?("xnix-runtime-service-binding"), "Runtime service binding must expose a CLI command")
%w[runtime-service-binding activation_binding_ready live_dbus_owner_ready dbus-service-activation systemd-service-hardening].each do |token|
  assert(runtime_service_binding_source.include?(token), "Runtime service binding must include #{token}")
end
assert(runtime_service_binding_source.include?("\"network_required\" => false"), "Runtime service binding must not require network access")
assert(runtime_service_binding_source.include?("\"host_root_modified\" => false"), "Runtime service binding must not mutate the host root")
assert(runtime_service_binding_source.include?("\"privileged_container_required\" => false"), "Runtime service binding must not require privileged containers")
assert(runtime_service_binding_source.include?("\"backend_details_exposed\" => false"), "Runtime service binding must hide backend details")

runtime_activation_installer_source = read_project_file("scripts/install_runtime_activation.rb")
%w[SOURCE_RUNTIME_LIB SOURCE_RECIPE_DIR SOURCE_DBUS_CONTRACT SOURCE_DBUS_SMOKE SOURCE_SESSION_SMOKE usr/lib/xnix usr/runtime/recipes].each do |token|
  assert(runtime_activation_installer_source.include?(token), "Runtime activation installer must include #{token}")
end

runtime_activation_smoke_source = read_project_file("scripts/runtime_activation_smoke.rb")
%w[usr/libexec/xnix/compatd usr/lib/xnix/compatibility/runtime_daemon.rb GetRuntimeMethodParityManifest GetRuntimeWriteGate WriteMethodDisabled].each do |token|
  assert(runtime_activation_smoke_source.include?(token), "Runtime activation smoke must include #{token}")
end

runtime_live_owner_gate_source = read_project_file("lib/xnix/compatibility/runtime_live_owner_gate.rb")
assert(runtime_live_owner_gate_source.include?("xnix-runtime-live-owner-gate"), "Runtime live owner gate must expose a CLI command")
%w[runtime-live-owner-gate required_gates long-running-runtime-owner bus-name-acquisition read-only-method-parity production-recipe-trust].each do |token|
  assert(runtime_live_owner_gate_source.include?(token), "Runtime live owner gate must include #{token}")
end
assert(runtime_live_owner_gate_source.include?("\"production_owner_enabled\" => false"), "Runtime live owner gate must not enable production ownership")
assert(runtime_live_owner_gate_source.include?("\"owner_transition_ready\" => false"), "Runtime live owner gate must not enable owner transition")
assert(runtime_live_owner_gate_source.include?("\"smoke_adapter_is_production_owner\" => false"), "Runtime live owner gate must keep the smoke adapter non-production")
assert(runtime_live_owner_gate_source.include?("\"kde_may_claim_runtime_ownership\" => false"), "Runtime live owner gate must prevent KDE Runtime ownership")
assert(runtime_live_owner_gate_source.include?("\"host_root_modified\" => false"), "Runtime live owner gate must not mutate the host root")
assert(runtime_live_owner_gate_source.include?("\"backend_details_exposed\" => false"), "Runtime live owner gate must hide backend details")

runtime_owner_smoke_plan_source = read_project_file("lib/xnix/compatibility/runtime_owner_smoke_plan.rb")
assert(runtime_owner_smoke_plan_source.include?("xnix-runtime-owner-smoke-plan"), "Runtime owner smoke plan must expose a CLI command")
%w[runtime-owner-smoke-plan validate-activation-files start-packaged-runtime-owner assert-stable-bus-name check-read-only-method-parity reject-write-methods verify-non-production-smoke-adapter-boundary].each do |token|
  assert(runtime_owner_smoke_plan_source.include?(token), "Runtime owner smoke plan must include #{token}")
end
assert(runtime_owner_smoke_plan_source.include?("\"smoke_state\" => \"planned\""), "Runtime owner smoke plan must stay planned")
assert(runtime_owner_smoke_plan_source.include?("\"system_service_started\" => false"), "Runtime owner smoke plan must not start system services")
assert(runtime_owner_smoke_plan_source.include?("\"production_bus_claimed\" => false"), "Runtime owner smoke plan must not claim production bus ownership")
assert(runtime_owner_smoke_plan_source.include?("\"host_root_modified\" => false"), "Runtime owner smoke plan must not mutate the host root")
assert(runtime_owner_smoke_plan_source.include?("\"backend_details_exposed\" => false"), "Runtime owner smoke plan must hide backend details")

runtime_method_parity_source = read_project_file("lib/xnix/compatibility/runtime_method_parity_manifest.rb")
assert(runtime_method_parity_source.include?("xnix-runtime-method-parity-manifest"), "Runtime method parity manifest must expose a CLI command")
%w[runtime-method-parity-manifest READ_ONLY_METHODS dbus-contract runtime-dispatch dbus-client smoke-adapter session-smoke GetRuntimeMethodParityManifest GetRuntimeWriteGate GetCompatibilityModeSwitchPlan GetCompatibilityPermissionReviewPlan GetCompatibilityReviewFlowPlan].each do |token|
  assert(runtime_method_parity_source.include?(token), "Runtime method parity manifest must include #{token}")
end
assert(runtime_method_parity_source.include?("\"write_methods_supported\" => false"), "Runtime method parity manifest must not claim write method support")
assert(runtime_method_parity_source.include?("\"write_method_dispatch_enabled\" => false"), "Runtime method parity manifest must not enable write method dispatch")
assert(runtime_method_parity_source.include?("\"host_root_modified\" => false"), "Runtime method parity manifest must not mutate the host root")
assert(runtime_method_parity_source.include?("\"backend_details_exposed\" => false"), "Runtime method parity manifest must hide backend details")

runtime_write_gate_source = read_project_file("lib/xnix/compatibility/runtime_write_gate.rb")
assert(runtime_write_gate_source.include?("xnix-runtime-write-gate"), "Runtime write gate must expose a CLI command")
%w[runtime-write-gate InstallRecipe Launch CreateSnapshot RestoreSnapshot blocked-until-production-backend WriteMethodDisabled production-runtime-owner backend-binding-ready user-action-review].each do |token|
  assert(runtime_write_gate_source.include?(token), "Runtime write gate must include #{token}")
end
assert(runtime_write_gate_source.include?("\"write_method_enabled\" => false"), "Runtime write gate must not enable write methods")
assert(runtime_write_gate_source.include?("\"dispatch_enabled\" => false"), "Runtime write gate must not enable dispatch")
assert(runtime_write_gate_source.include?("\"request_object_created\" => false"), "Runtime write gate must not create request objects")
assert(runtime_write_gate_source.include?("\"execution_started\" => false"), "Runtime write gate must not start execution")
assert(runtime_write_gate_source.include?("\"host_root_modified\" => false"), "Runtime write gate must not mutate the host root")
assert(runtime_write_gate_source.include?("\"backend_details_exposed\" => false"), "Runtime write gate must hide backend details")

notification_source = read_project_file("lib/xnix/compatibility/notification_request.rb")
%w[install-failed repair-applied mode-changed approval-required].each do |event_type|
  assert(notification_source.include?("\"#{event_type}\""), "Notification requests must include #{event_type}")
end

settings_source = read_project_file("lib/xnix/compatibility/settings_model.rb")
%w[run-mode resource-access devices network snapshots].each do |section_id|
  assert(settings_source.include?("\"id\" => \"#{section_id}\""), "Settings model must include #{section_id}")
end
assert(settings_source.include?("\"runtime_owned\" => true"), "Settings model must be Runtime-owned")
assert(settings_source.include?("\"settings_state\" => \"planned\""), "Settings model must expose planned settings state")
assert(settings_source.include?("\"host_root_modified\" => false"), "Settings model must not mutate the host root")
assert(settings_source.include?("\"backend_details_exposed\" => false"), "Settings model must hide backend details")

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
assert(kwin_rule_source.include?("runtime.kwin_window_rule_plan"), "KWin window rule must consume Runtime-owned plans")
assert(kwin_rule_source.include?("normalized_set"), "KWin window rule must normalize flat D-Bus plans")
assert(kwin_rule_source.include?("\"backend_details_exposed\" => false"), "KWin window rule must hide backend details")

kde_status_source = read_project_file("lib/xnix/compatibility/kde_integration_status.rb")
%w[launcher task-manager file-manager system-tray notifications compatibility-center settings].each do |entry_point|
  assert(kde_status_source.include?("\"#{entry_point}\""), "KDE integration status must include #{entry_point}")
end
assert(kde_status_source.include?("runtime.kde_integration_status"), "KDE integration status must consume Runtime-owned status")
assert(kde_status_source.include?("normalized_entry_points"), "KDE integration status must normalize flat D-Bus status")

kde_shell_plan_source = read_project_file("lib/xnix/compatibility/kde_shell_integration_plan.rb")
%w[xnix-kde-shell-integration-plan kde-shell-integration-plan GetKDEShellIntegrationPlan start-menu task-manager file-manager system-tray notification-center compatibility-center unified-settings krunner-search kwin-window-management].each do |token|
  assert(kde_shell_plan_source.include?(token), "KDE shell integration plan must include #{token}")
end
%w[plasma_fork_required plasma_source_modified shell_configuration_written component_activation_enabled backend_launch_enabled backend_details_exposed host_root_modified privileged_container_required].each do |token|
  assert(kde_shell_plan_source.include?("\"#{token}\" => false"), "KDE shell integration plan must keep #{token} false")
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
assert(runtime_daemon_source.include?("\"compatibility_acquisition_preflight\""), "Runtime daemon must expose compatibility acquisition preflight capability")
assert(runtime_daemon_source.include?("\"compatibility_action_queues\""), "Runtime daemon must expose Compatibility Center action queue capability")
assert(runtime_daemon_source.include?("\"compatibility_action_review_receipts\""), "Runtime daemon must expose Compatibility Center action review receipt capability")
assert(runtime_daemon_source.include?("\"compatibility_center_summaries\""), "Runtime daemon must expose Compatibility Center summary capability")
assert(runtime_daemon_source.include?("\"compatibility_artifact_manifests\""), "Runtime daemon must expose compatibility artifact manifest capability")
assert(runtime_daemon_source.include?("\"compatibility_install_planning\""), "Runtime daemon must expose compatibility install planning capability")
assert(runtime_daemon_source.include?("\"compatibility_mode_switch_planning\""), "Runtime daemon must expose compatibility mode switch planning capability")
assert(runtime_daemon_source.include?("\"compatibility_permission_review_planning\""), "Runtime daemon must expose compatibility permission review planning capability")
assert(runtime_daemon_source.include?("\"compatibility_package_sources\""), "Runtime daemon must expose compatibility package source capability")
assert(runtime_daemon_source.include?("\"compatibility_backend_binding\""), "Runtime daemon must expose compatibility backend binding capability")
assert(runtime_daemon_source.include?("\"compatibility_backend_capability_matrix\""), "Runtime daemon must expose compatibility backend capability matrix capability")
assert(runtime_daemon_source.include?("\"compatibility_backend_selection_plans\""), "Runtime daemon must expose compatibility backend selection planning capability")
assert(runtime_daemon_source.include?("\"compatibility_test_planning\""), "Runtime daemon must expose compatibility test planning capability")
assert(runtime_daemon_source.include?("\"compatibility_test_results\""), "Runtime daemon must expose compatibility test result capability")
assert(runtime_daemon_source.include?("\"ai_diagnostic_inputs\""), "Runtime daemon must expose AI diagnostic input capability")
assert(runtime_daemon_source.include?("\"ai_diagnostic_recommendations\""), "Runtime daemon must expose AI diagnostic recommendation capability")
assert(runtime_daemon_source.include?("\"ai_repair_approval_gates\""), "Runtime daemon must expose AI repair approval gate capability")
assert(runtime_daemon_source.include?("\"application_state_roots\""), "Runtime daemon must expose application state root capability")
assert(runtime_daemon_source.include?("\"runtime_live_owner_gates\""), "Runtime daemon must expose Runtime live owner gate capability")
assert(runtime_daemon_source.include?("\"runtime_method_parity_manifests\""), "Runtime daemon must expose Runtime method parity manifest capability")
assert(runtime_daemon_source.include?("\"runtime_owner_smoke_plans\""), "Runtime daemon must expose Runtime owner smoke plan capability")
assert(runtime_daemon_source.include?("\"runtime_service_binding\""), "Runtime daemon must expose Runtime service binding capability")
assert(runtime_daemon_source.include?("\"runtime_write_gates\""), "Runtime daemon must expose Runtime write gate capability")
assert(runtime_daemon_source.include?("\"compatibility_settings\""), "Runtime daemon must expose compatibility settings capability")
assert(runtime_daemon_source.include?("\"compatibility_settings_change_planning\""), "Runtime daemon must expose settings change planning capability")
assert(runtime_daemon_source.include?("\"compatibility_review_flow_planning\""), "Runtime daemon must expose compatibility review flow planning capability")
assert(runtime_daemon_source.include?("\"portal_request_planning\""), "Runtime daemon must expose Portal request planning capability")
assert(runtime_daemon_source.include?("\"kde_integration_status\""), "Runtime daemon must expose KDE integration status capability")
assert(runtime_daemon_source.include?("\"kde_shell_integration_plans\""), "Runtime daemon must expose KDE shell integration planning capability")
assert(runtime_daemon_source.include?("\"kde_application_surface_plans\""), "Runtime daemon must expose KDE application surface planning capability")
assert(runtime_daemon_source.include?("\"desktop_resource_bridge_plans\""), "Runtime daemon must expose desktop resource bridge planning capability")
assert(runtime_daemon_source.include?("\"desktop_entry_planning\""), "Runtime daemon must expose desktop entry planning capability")
assert(runtime_daemon_source.include?("\"task_manager_identity_planning\""), "Runtime daemon must expose task manager identity planning capability")
assert(runtime_daemon_source.include?("\"kwin_window_rule_planning\""), "Runtime daemon must expose KWin window rule planning capability")
assert(runtime_daemon_source.include?("\"file_association_planning\""), "Runtime daemon must expose file association planning capability")
assert(runtime_daemon_source.include?("\"notification_planning\""), "Runtime daemon must expose notification planning capability")
assert(runtime_daemon_source.include?("\"tray_status_planning\""), "Runtime daemon must expose tray status planning capability")
assert(runtime_daemon_source.include?("\"krunner_query_planning\""), "Runtime daemon must expose KRunner query planning capability")
assert(runtime_daemon_source.include?("\"compatibility_execution_readiness\""), "Runtime daemon must expose execution readiness capability")
assert(runtime_daemon_source.include?("\"compatibility_launch_intents\""), "Runtime daemon must expose launch intent capability")
%w[GetEngineCatalog GetRunPlan GetDesktopActivationManifest GetKDEIntegrationStatus GetKDEShellIntegrationPlan GetKDEApplicationSurfacePlan GetDesktopResourceBridgePlan GetCompatibilityModeSwitchPlan GetCompatibilityPermissionReviewPlan GetCompatibilityReviewFlowPlan GetDesktopEntryPlan GetTaskManagerIdentityPlan GetKWinWindowRulePlan GetFileAssociationPlan GetNotificationPlan GetTrayStatus GetKRunnerQueryPlan GetPortalRequestPlan GetApplicationStateRoot GetCompatibilityPackageSource GetCompatibilityAcquisitionPreflight GetCompatibilityActionQueue GetCompatibilityActionReviewReceipt GetCompatibilityCenterSummary GetCompatibilityArtifactManifest GetCompatibilityInstallPlan GetBackendBinding GetBackendCapabilityMatrix GetBackendSelectionPlan GetBackendLifecycle GetBackendEnvironmentPlan GetRepairPlan GetTestPlan GetTestResult GetExecutionReadiness GetLaunchIntent GetAIDiagnosticInput GetAIDiagnosticRecommendation GetAIRepairApprovalGate GetSnapshotPlan GetPortalAccessPolicy GetRuntimeServiceBinding GetRuntimeLiveOwnerGate GetRuntimeOwnerSmokePlan GetRuntimeMethodParityManifest GetRuntimeWriteGate GetCompatibilitySettings GetCompatibilitySettingsChangePlan].each do |method_name|
  assert(runtime_daemon_source.include?("\"#{method_name}\""), "Runtime daemon dispatch must include #{method_name}")
  assert(read_project_file("runtime/dbus/org.xnix.Compatibility1.xml").include?("name=\"#{method_name}\""), "D-Bus contract must include #{method_name}")
  assert(read_project_file("runtime/dbus/xnix_compatd_smoke.c").include?(method_name), "D-Bus smoke adapter must include #{method_name}")
  assert(read_project_file("scripts/dbus_session_smoke.rb").include?(method_name), "D-Bus session smoke must call #{method_name}")
end

dbus_client_source = read_project_file("lib/xnix/compatibility/dbus_runtime_client.rb")
%w[engine_catalog run_plan kde_integration_status kde_shell_integration_plan kde_application_surface_plan desktop_resource_bridge_plan compatibility_mode_switch_plan compatibility_permission_review_plan compatibility_review_flow_plan desktop_entry_plan task_manager_identity_plan kwin_window_rule_plan file_association_plan notification_plan tray_status krunner_query_plan state_root package_source acquisition_preflight action_queue action_review_receipt compatibility_center_summary artifact_manifest install_plan backend_binding backend_capability_matrix backend_selection_plan repair_plan test_plan test_result execution_readiness ai_diagnostic_input ai_diagnostic_recommendation ai_repair_approval_gate snapshot_plan portal_access_policy runtime_service_binding runtime_live_owner_gate runtime_owner_smoke_plan runtime_method_parity_manifest runtime_write_gate settings settings_change_plan].each do |method_name|
  assert(dbus_client_source.include?("def #{method_name}"), "D-Bus Runtime client must expose #{method_name}")
end
assert(dbus_client_source.include?("return true if value == \"true\""), "D-Bus Runtime client must parse boolean true values")
assert(dbus_client_source.include?("return false if value == \"false\""), "D-Bus Runtime client must parse boolean false values")
assert(dbus_client_source.include?("value.to_i"), "D-Bus Runtime client must parse integer values")
assert(dbus_client_source.include?("value.start_with?(\"[\")"), "D-Bus Runtime client must parse string arrays")

settings_change_source = read_project_file("lib/xnix/compatibility/settings_change_plan.rb")
assert(settings_change_source.include?("xnix-compat-settings-change"), "Settings change plan must expose a CLI command")
%w[settings-change-plan validate-setting review-user-confirmation review-portal-policy prepare-restore-point persist-runtime-setting].each do |token|
  assert(settings_change_source.include?(token), "Settings change plan must include #{token}")
end
%w[apply_enabled settings_persisted host_root_modified backend_details_exposed].each do |token|
  assert(settings_change_source.include?("\"#{token}\" => false"), "Settings change plan must keep #{token} false while planning")
end

mode_switch_source = read_project_file("lib/xnix/compatibility/compatibility_mode_switch_plan.rb")
assert(mode_switch_source.include?("xnix-compat-mode-switch-plan"), "Compatibility mode switch plan must expose a CLI command")
%w[compatibility-mode-switch-plan GetCompatibilityModeSwitchPlan automatic prefer-performance prefer-compatibility isolated-execution settings-review portal-policy-review snapshot-baseline backend-environment-plan runtime-write-gate].each do |token|
  assert(mode_switch_source.include?(token), "Compatibility mode switch plan must include #{token}")
end
%w[settings_persistence_enabled backend_reconfiguration_enabled backend_process_started launch_enabled host_root_modified backend_details_exposed].each do |token|
  assert(mode_switch_source.include?("\"#{token}\" => false"), "Compatibility mode switch plan must keep #{token} false while planning")
end

permission_review_source = read_project_file("lib/xnix/compatibility/compatibility_permission_review_plan.rb")
assert(permission_review_source.include?("xnix-compat-permission-review-plan"), "Compatibility permission review plan must expose a CLI command")
%w[compatibility-permission-review-plan GetCompatibilityPermissionReviewPlan documents downloads camera network clipboard print screenshot user-review portal-policy-review settings-persistence audit-log].each do |token|
  assert(permission_review_source.include?(token), "Compatibility permission review plan must include #{token}")
end
%w[permission_changes_applied request_objects_created permissions_granted settings_persisted host_permission_changed host_root_modified backend_details_exposed].each do |token|
  assert(permission_review_source.include?("\"#{token}\" => false"), "Compatibility permission review plan must keep #{token} false while planning")
end

review_flow_source = read_project_file("lib/xnix/compatibility/compatibility_review_flow_plan.rb")
assert(review_flow_source.include?("xnix-compat-review-flow-plan"), "Compatibility review flow plan must expose a CLI command")
%w[compatibility-review-flow-plan GetCompatibilityReviewFlowPlan settings-change-review permission-review portal-request-review runtime-write-gate review-receipt GetCompatibilitySettingsChangePlan GetCompatibilityPermissionReviewPlan GetPortalRequestPlan GetRuntimeWriteGate GetCompatibilityActionReviewReceipt].each do |token|
  assert(review_flow_source.include?(token), "Compatibility review flow plan must include #{token}")
end
%w[apply_enabled request_object_created permission_granted settings_persisted execution_started host_root_modified backend_details_exposed].each do |token|
  assert(review_flow_source.include?("\"#{token}\" => false"), "Compatibility review flow plan must keep #{token} false while planning")
end

action_queue_source = read_project_file("lib/xnix/compatibility/compatibility_action_queue.rb")
assert(action_queue_source.include?("xnix-compat-action-queue"), "Compatibility action queue must expose a CLI command")
%w[compatibility-center-action-queue review-install-readiness review-settings-change review-ai-repair verify-runtime-service review-portal-policy].each do |token|
  assert(action_queue_source.include?(token), "Compatibility action queue must include #{token}")
end
%w[execution_enabled repair_execution_enabled settings_persistence_enabled host_root_modified network_required backend_details_exposed].each do |token|
  assert(action_queue_source.include?("\"#{token}\" => false"), "Compatibility action queue must keep #{token} false")
end

action_review_source = read_project_file("lib/xnix/compatibility/action_review_receipt.rb")
assert(action_review_source.include?("xnix-compat-action-review"), "Compatibility action review receipt must expose a CLI command")
%w[compatibility-center-action-review-receipt decision_recorded reviewed approved deferred rejected].each do |token|
  assert(action_review_source.include?(token), "Compatibility action review receipt must include #{token}")
end
%w[execution_enabled repair_execution_enabled settings_persistence_enabled resource_grant_created host_root_modified network_required backend_details_exposed].each do |token|
  assert(action_review_source.include?("\"#{token}\" => false"), "Compatibility action review receipt must keep #{token} false")
end

krunner_source = read_project_file("lib/xnix/compatibility/krunner_model.rb")
assert(krunner_source.include?("xnix-compat-launch"), "KRunner model must delegate launches to the managed launcher")
assert(krunner_source.include?("\"backend_details_exposed\" => false"), "KRunner model must hide backend details")
assert(krunner_source.include?("runtime.krunner_query_plan"), "KRunner model must consume Runtime-owned query plans")
assert(krunner_source.include?("normalized_flat_plan_match"), "KRunner model must normalize flat D-Bus query plans")
assert(krunner_source.include?("\"query_execution_enabled\""), "KRunner model must preserve query execution gates")
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
