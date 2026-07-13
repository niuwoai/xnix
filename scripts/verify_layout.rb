#!/usr/bin/env ruby
# frozen_string_literal: true

require "pathname"

PROJECT_ROOT = Pathname.new(__dir__).join("..").realpath
EXPECTED_VERSION = "0.2.75"
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
  lib/xnix/compatibility/action_review_receipt.rb
  lib/xnix/compatibility/ai_diagnostic_input.rb
  lib/xnix/compatibility/ai_diagnostic_recommendation.rb
  lib/xnix/compatibility/ai_repair_approval_gate.rb
  lib/xnix/compatibility/application_state_root.rb
  lib/xnix/compatibility/compatibility_acquisition_preflight.rb
  lib/xnix/compatibility/compatibility_action_queue.rb
  lib/xnix/compatibility/compatibility_artifact_manifest.rb
  lib/xnix/compatibility/compatibility_backend_binding.rb
  lib/xnix/compatibility/compatibility_engine_catalog.rb
  lib/xnix/compatibility/compatibility_install_plan.rb
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
  bin/xnix-compat-state-root
  bin/xnix-compat-backend-binding
  bin/xnix-compatd
  bin/xnix-compat-engine-catalog
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
  bin/xnix-runtime-live-owner-gate
  bin/xnix-runtime-method-parity-manifest
  bin/xnix-runtime-owner-smoke-plan
  bin/xnix-runtime-service-binding
  bin/xnix-runtime-write-gate
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
  runtime/core/xnix_runtime_core_cli.c
  runtime/recipes/registry.json
  runtime/recipes/org.xnix.sample.notepad.json
  runtime/systemd/xnix-compatd.service
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
  test/test_compatibility_install_plan.rb
  test/test_compatibility_backend_binding.rb
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

ai_recommendation_source = read_project_file("lib/xnix/compatibility/ai_diagnostic_recommendation.rb")
assert(ai_recommendation_source.include?("xnix-ai-diagnostic-recommendation"), "AI diagnostic recommendation must expose a CLI command")
%w[ai-diagnostic-recommendation recommendations approval_required_actions blocked_actions].each do |token|
  assert(ai_recommendation_source.include?(token), "AI diagnostic recommendation must include #{token}")
end
assert(ai_recommendation_source.include?("\"ai_provider_called\" => false"), "AI diagnostic recommendation must not call an AI provider")
assert(ai_recommendation_source.include?("\"network_required\" => false"), "AI diagnostic recommendation must not require network access")
assert(ai_recommendation_source.include?("\"auto_execute\" => false"), "AI diagnostic recommendation must not auto-execute recommendations")
assert(ai_recommendation_source.include?("\"backend_details_exposed\" => false"), "AI diagnostic recommendation must hide backend details")

ai_repair_gate_source = read_project_file("lib/xnix/compatibility/ai_repair_approval_gate.rb")
assert(ai_repair_gate_source.include?("xnix-ai-repair-approval-gate"), "AI repair approval gate must expose a CLI command")
%w[ai-repair-approval-gate blocked-until-approval required_gates approval_required_actions].each do |token|
  assert(ai_repair_gate_source.include?(token), "AI repair approval gate must include #{token}")
end
assert(ai_repair_gate_source.include?("\"repair_executed\" => false"), "AI repair approval gate must not execute repairs")
assert(ai_repair_gate_source.include?("\"auto_execution_allowed\" => false"), "AI repair approval gate must not permit automatic repair execution")
assert(ai_repair_gate_source.include?("\"backend_details_exposed\" => false"), "AI repair approval gate must hide backend details")

c_runtime_core_header = read_project_file("runtime/core/xnix_runtime_core.h")
c_runtime_core_source = read_project_file("runtime/core/xnix_runtime_core.c")
c_runtime_core_cli = read_project_file("runtime/core/xnix_runtime_core_cli.c") +
                     read_project_file("runtime/core/xnix_runtime_core_cli_dispatch.inc")
assert(c_runtime_core_header.include?("#define XNIX_RUNTIME_VERSION \"#{EXPECTED_VERSION}\""), "C Runtime core must expose #{EXPECTED_VERSION}")
%w[XNIX_RUNTIME_BUS_NAME XNIX_RUNTIME_OBJECT_PATH XNIX_RUNTIME_INTERFACE XNIX_RUNTIME_WRITE_ERROR].each do |token|
  assert(c_runtime_core_header.include?(token), "C Runtime core header must include #{token}")
end
%w[XnixRuntimeApplication xnix_runtime_application_count xnix_runtime_find_application XnixRuntimeEngine xnix_runtime_engine_count xnix_runtime_select_engine_for_mode XnixRuntimePortalPolicy xnix_runtime_portal_policy_count xnix_runtime_find_portal_policy XnixRuntimeSnapshotPolicy xnix_runtime_snapshot_policy_count xnix_runtime_find_snapshot_policy XnixRuntimeStateRootPolicy xnix_runtime_state_root_policy_count xnix_runtime_find_state_root_policy XnixRuntimeInstallReadinessPolicy xnix_runtime_install_readiness_policy_count xnix_runtime_find_install_readiness_policy XnixRuntimeArtifactManifestPolicy xnix_runtime_artifact_manifest_policy_count xnix_runtime_find_artifact_manifest_policy XnixRuntimeAcquisitionPreflightPolicy xnix_runtime_acquisition_preflight_policy_count xnix_runtime_find_acquisition_preflight_policy XnixRuntimePackageSourcePolicy xnix_runtime_package_source_policy_count xnix_runtime_find_package_source_policy XnixRuntimeBackendBindingPolicy xnix_runtime_backend_binding_policy_count xnix_runtime_find_backend_binding_policy XnixRuntimeSettingsPolicy xnix_runtime_settings_policy_count xnix_runtime_find_settings_policy XnixRuntimeSettingsChangePolicy xnix_runtime_settings_change_policy_count xnix_runtime_find_settings_change_policy XnixRuntimeServiceBindingPolicy xnix_runtime_service_binding_policy XnixRuntimeLiveOwnerGatePolicy xnix_runtime_live_owner_gate_policy XnixRuntimeOwnerSmokePlanPolicy xnix_runtime_owner_smoke_plan_policy XnixRuntimeMethodParityManifestPolicy xnix_runtime_method_parity_manifest_policy XnixRuntimeRecipeTrustPolicy xnix_runtime_recipe_trust_policy XnixRuntimeRecipeInstallGate xnix_runtime_recipe_install_gate XnixRuntimeCompatibilityInstallPlan xnix_runtime_compatibility_install_plan].each do |token|
  assert(c_runtime_core_header.include?(token), "C Runtime core header must include #{token}")
end
%w[xnix_runtime_write_gate blocked-until-production-backend InstallRecipe Launch CreateSnapshot RestoreSnapshot xnix_runtime_find_application org.xnix.sample.notepad application/x-xnix-txt xnix_runtime_select_engine_for_mode automatic-managed local-compatibility-engine isolated-compatibility-engine xnix_runtime_find_portal_policy org.freedesktop.portal.FileChooser org.freedesktop.portal.Camera remote-desktop selected-files xnix_runtime_find_snapshot_policy before-repair before-engine-change desktop_activation_receipts xnix_runtime_find_state_root_policy application-data runtime-metadata diagnostic-cache xnix_runtime_find_install_readiness_policy resolve-artifact-manifest verify-artifact-digests prepare-package-source allocate-application-state xnix_runtime_find_artifact_manifest_policy runtime-launch-metadata local-execution-artifacts isolated-environment-artifacts manifest-signature-verification artifact-digest-verification xnix_runtime_find_acquisition_preflight_policy package-source-ready signed-artifact-manifest runtime-cache-space network-policy-review rollback-marker xnix_runtime_find_package_source_policy os-managed-compatibility-packages runtime-managed-toolcache isolated-environment-template-catalog signed-source-verification source-policy-review runtime-cache-quota offline-fallback xnix_runtime_find_backend_binding_policy engine-package-source application-state-root portal-policy-review snapshot-baseline xnix_runtime_find_settings_policy run-mode resource-access devices snapshots settings_change_policies xnix_runtime_find_settings_change_policy validate-setting persist-runtime-setting service_binding_policy xnix_runtime_service_binding_policy dbus-service-activation live-dbus-owner live_owner_gate_policy xnix_runtime_live_owner_gate_policy production-recipe-trust owner_smoke_plan_policy xnix_runtime_owner_smoke_plan_policy validate-activation-files assert-stable-bus-name method_parity_manifest_policy xnix_runtime_method_parity_manifest_policy recipe_trust_policy xnix_runtime_recipe_trust_policy production_recipe_install_requirements xnix_runtime_compatibility_install_plan signed recipe validation is not enabled registry.digest registry.signature registry.development GetRuntimeMethodParityManifest GetRuntimeWriteGate].each do |token|
  assert(c_runtime_core_source.include?(token), "C Runtime core source must include #{token}")
end
%w[business_logic_runtime tests-and-development-tools application_catalog_owner engine_catalog_owner portal_policy_owner snapshot_policy_owner state_root_policy_owner install_readiness_policy_owner artifact_manifest_policy_owner acquisition_preflight_policy_owner package_source_policy_owner backend_binding_policy_owner settings_policy_owner settings_change_policy_owner service_binding_policy_owner live_owner_gate_policy_owner owner_smoke_plan_policy_owner method_parity_manifest_policy_owner recipe_trust_policy_owner recipe_install_gate_owner compatibility_install_plan_owner list-applications get-application list-engines select-engine list-portal-policies portal-policy list-snapshot-policies snapshot-policy list-state-root-policies state-root-policy list-install-readiness-policies install-readiness-policy compatibility-install-plan list-artifact-manifest-policies artifact-manifest-policy list-acquisition-preflight-policies acquisition-preflight-policy list-package-source-policies package-source-policy list-backend-binding-policies backend-binding-policy list-settings-policies settings-policy list-settings-change-policies settings-change-policy runtime-service-binding runtime-live-owner-gate runtime-owner-smoke-plan runtime-method-parity-manifest recipe-trust-policy recipe-install-gate write-gate].each do |token|
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
%w[runtime-method-parity-manifest READ_ONLY_METHODS dbus-contract runtime-dispatch dbus-client smoke-adapter session-smoke GetRuntimeMethodParityManifest GetRuntimeWriteGate].each do |token|
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
assert(runtime_daemon_source.include?("\"compatibility_acquisition_preflight\""), "Runtime daemon must expose compatibility acquisition preflight capability")
assert(runtime_daemon_source.include?("\"compatibility_action_queues\""), "Runtime daemon must expose Compatibility Center action queue capability")
assert(runtime_daemon_source.include?("\"compatibility_action_review_receipts\""), "Runtime daemon must expose Compatibility Center action review receipt capability")
assert(runtime_daemon_source.include?("\"compatibility_artifact_manifests\""), "Runtime daemon must expose compatibility artifact manifest capability")
assert(runtime_daemon_source.include?("\"compatibility_install_planning\""), "Runtime daemon must expose compatibility install planning capability")
assert(runtime_daemon_source.include?("\"compatibility_package_sources\""), "Runtime daemon must expose compatibility package source capability")
assert(runtime_daemon_source.include?("\"compatibility_backend_binding\""), "Runtime daemon must expose compatibility backend binding capability")
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
%w[GetEngineCatalog GetRunPlan GetApplicationStateRoot GetCompatibilityPackageSource GetCompatibilityAcquisitionPreflight GetCompatibilityActionQueue GetCompatibilityActionReviewReceipt GetCompatibilityArtifactManifest GetCompatibilityInstallPlan GetBackendBinding GetRepairPlan GetTestPlan GetTestResult GetAIDiagnosticInput GetAIDiagnosticRecommendation GetAIRepairApprovalGate GetSnapshotPlan GetPortalAccessPolicy GetRuntimeServiceBinding GetRuntimeLiveOwnerGate GetRuntimeOwnerSmokePlan GetRuntimeMethodParityManifest GetRuntimeWriteGate GetCompatibilitySettings GetCompatibilitySettingsChangePlan].each do |method_name|
  assert(runtime_daemon_source.include?("\"#{method_name}\""), "Runtime daemon dispatch must include #{method_name}")
  assert(read_project_file("runtime/dbus/org.xnix.Compatibility1.xml").include?("name=\"#{method_name}\""), "D-Bus contract must include #{method_name}")
  assert(read_project_file("runtime/dbus/xnix_compatd_smoke.c").include?(method_name), "D-Bus smoke adapter must include #{method_name}")
  assert(read_project_file("scripts/dbus_session_smoke.rb").include?(method_name), "D-Bus session smoke must call #{method_name}")
end

dbus_client_source = read_project_file("lib/xnix/compatibility/dbus_runtime_client.rb")
%w[engine_catalog run_plan state_root package_source acquisition_preflight action_queue action_review_receipt artifact_manifest install_plan backend_binding repair_plan test_plan test_result ai_diagnostic_input ai_diagnostic_recommendation ai_repair_approval_gate snapshot_plan portal_access_policy runtime_service_binding runtime_live_owner_gate runtime_owner_smoke_plan runtime_method_parity_manifest runtime_write_gate settings settings_change_plan].each do |method_name|
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
