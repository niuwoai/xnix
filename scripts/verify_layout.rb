#!/usr/bin/env ruby
# frozen_string_literal: true

require "pathname"

PROJECT_ROOT = Pathname.new(__dir__).join("..").realpath
EXPECTED_VERSION = "0.2.382"
REQUIRED_FILES = %w[
  .dockerignore
  Dockerfile
  VERSION
  docs/claude-code-implementation-packages.md
  docs/claude-code-domain-dispatch.md
  docs/claude-code-next-implementation-assignments.md
  docs/claude-code-contract-implementation-handoff.md
  docs/claude-code-empty-domain-implementation-packages.md
  docs/claude-code-independent-implementation-briefs.md
  docs/claude-code-mainline-implementation-plan.md
  docs/claude-code-mainline-task-batch.md
  docs/claude-code-second-wave-task-batch.md
  docs/claude-code-third-wave-task-batch.md
  docs/claude-code-fourth-wave-task-batch.md
  docs/claude-code-fifth-wave-task-batch.md
  docs/claude-code-sixth-wave-task-batch.md
  docs/claude-code-seventh-wave-task-batch.md
  docs/claude-code-stability-release-train.md
  docs/claude-code-dispatch-runbook.md
  docs/claude-code-open-domain-work-packages.md
  docs/claude-code-windows-compatibility-workstreams.md
  docs/xnix-current-mainline.md
  docs/mainline-integration-checkpoint.md
  docs/kde-first-compatibility-acceptance.md
  docs/kde-first-current-gap-audit.md
  docs/kde-first-presence-smoke-spec.md
  docs/windows-app-compatibility-implementation-brief.md
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
  lib/xnix/full_smoke_report.rb
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
  cmd/xnix-runtime-go/version_test.go
  cmd/xnix-runtime-go/application_readiness_commands.go
  cmd/xnix-runtime-go/application_readiness_cli_test.go
  cmd/xnix-runtime-go/application_upgrade_impact_commands.go
  cmd/xnix-runtime-go/application_upgrade_impact_cli_test.go
  cmd/xnix-runtime-go/runtime_policy_explanation_cards_commands.go
  cmd/xnix-runtime-go/runtime_policy_explanation_cards_cli_test.go
  cmd/xnix-runtime-go/compatibility_onboarding_commands.go
  cmd/xnix-runtime-go/compatibility_onboarding_cli_test.go
  cmd/xnix-runtime-go/artifact_stage_commands.go
  cmd/xnix-runtime-go/artifact_stage_cli_test.go
  cmd/xnix-runtime-go/backend_adapter_contract_commands.go
  cmd/xnix-runtime-go/backend_adapter_contract_cli_test.go
  cmd/xnix-runtime-go/backend_adapter_contract_owner_route_audit_cli_test.go
  cmd/xnix-runtime-go/backend_adapter_contract_redacted_profile_audit_cli_test.go
  cmd/xnix-runtime-go/restricted_owner_smoke_receipt_commands.go
  cmd/xnix-runtime-go/restricted_owner_smoke_receipt_cli_test.go
  cmd/xnix-runtime-go/desktop_activation_manifest_commands.go
  cmd/xnix-runtime-go/desktop_activation_stage_commands.go
  cmd/xnix-runtime-go/desktop_activation_stage_cli_test.go
  cmd/xnix-runtime-go/install_plan_commands.go
  cmd/xnix-runtime-go/install_plan_cli_test.go
  cmd/xnix-runtime-go/execution_ledger_commands.go
  cmd/xnix-runtime-go/execution_ledger_cli_test.go
  cmd/xnix-runtime-go/diagnostic_record_commands.go
  cmd/xnix-runtime-go/diagnostic_record_cli_test.go
  cmd/xnix-runtime-go/support_bundle_manifest_commands.go
  cmd/xnix-runtime-go/support_bundle_manifest_cli_test.go
  cmd/xnix-runtime-go/support_case_timeline_commands.go
  cmd/xnix-runtime-go/support_case_timeline_cli_test.go
  cmd/xnix-runtime-go/multi_application_install_queue_commands.go
  cmd/xnix-runtime-go/multi_application_install_queue_cli_test.go
  cmd/xnix-runtime-go/offline_application_fixture_matrix_commands.go
  cmd/xnix-runtime-go/offline_application_fixture_matrix_cli_test.go
  cmd/xnix-runtime-go/ai_diagnostics_commands.go
  cmd/xnix-runtime-go/ai_diagnostics_cli_test.go
  cmd/xnix-runtime-go/desktop_safety_policy_commands.go
  cmd/xnix-runtime-go/desktop_safety_policy_cli_test.go
  cmd/xnix-runtime-go/kde_shell_commands.go
  cmd/xnix-runtime-go/kde_shell_cli_test.go
  cmd/xnix-runtime-go/window_identity_commands.go
  cmd/xnix-runtime-go/window_identity_cli_test.go
  cmd/xnix-runtime-go/windows_compatibility_commands.go
  cmd/xnix-runtime-go/windows_compatibility_cli_test.go
  cmd/xnix-runtime-go/runtime_safety_commands.go
  cmd/xnix-runtime-go/runtime_safety_cli_test.go
  cmd/xnix-runtime-go/state_root_quota_retention_commands.go
  cmd/xnix-runtime-go/state_root_quota_retention_cli_test.go
  cmd/xnix-runtime-go/crash_hang_signal_summary_commands.go
  cmd/xnix-runtime-go/crash_hang_signal_summary_cli_test.go
  cmd/xnix-runtime-go/permission_evidence_audit_commands.go
  cmd/xnix-runtime-go/permission_evidence_audit_cli_test.go
  cmd/xnix-runtime-go/portal_permission_renewal_commands.go
  cmd/xnix-runtime-go/portal_permission_renewal_cli_test.go
  cmd/xnix-runtime-go/compatibility_backend_fallback_commands.go
  cmd/xnix-runtime-go/compatibility_backend_fallback_cli_test.go
  cmd/xnix-runtime-go/kde_search_visibility_commands.go
  cmd/xnix-runtime-go/kde_search_visibility_cli_test.go
  cmd/xnix-runtime-go/kde_notification_digest_commands.go
  cmd/xnix-runtime-go/kde_notification_digest_cli_test.go
  cmd/xnix-runtime-go/kde_offline_application_identity_commands.go
  cmd/xnix-runtime-go/kde_offline_application_identity_cli_test.go
  cmd/xnix-runtime-go/kde_fake_execution_evidence_commands.go
  cmd/xnix-runtime-go/kde_fake_execution_evidence_cli_test.go
  cmd/xnix-runtime-go/kde_fake_portal_evidence_commands.go
  cmd/xnix-runtime-go/kde_fake_portal_evidence_cli_test.go
  cmd/xnix-runtime-go/kde_snapshot_diagnostics_evidence_commands.go
  cmd/xnix-runtime-go/kde_snapshot_diagnostics_evidence_cli_test.go
  cmd/xnix-runtime-go/kde_backend_lifecycle_evidence_commands.go
  cmd/xnix-runtime-go/kde_backend_lifecycle_evidence_cli_test.go
  cmd/xnix-runtime-go/kde_restricted_launch_authorization_commands.go
  cmd/xnix-runtime-go/kde_restricted_launch_authorization_cli_test.go
  cmd/xnix-runtime-go/kde_test_launch_materialization_commands.go
  cmd/xnix-runtime-go/kde_test_launch_materialization_cli_test.go
  cmd/xnix-runtime-go/kde_test_launch_materialization_fanout_commands.go
  cmd/xnix-runtime-go/kde_test_launch_materialization_fanout_cli_test.go
  cmd/xnix-runtime-go/kde_test_launch_materialization_owner_route_audit_cli_test.go
  cmd/xnix-runtime-go/desktop_deactivation_dry_run_commands.go
  cmd/xnix-runtime-go/desktop_deactivation_dry_run_cli_test.go
  cmd/xnix-runtime-go/recipe_conflict_audit_commands.go
  cmd/xnix-runtime-go/recipe_conflict_audit_cli_test.go
  cmd/xnix-runtime-go/signed_recipe_verifier_commands.go
  cmd/xnix-runtime-go/signed_recipe_verifier_cli_test.go
  cmd/xnix-runtime-go/restricted_product_smoke_packet_commands.go
  cmd/xnix-runtime-go/restricted_product_smoke_packet_cli_test.go
  cmd/xnix-runtime-go/snapshot_restore_candidates_commands.go
  cmd/xnix-runtime-go/snapshot_restore_candidates_cli_test.go
  cmd/xnix-runtime-go/runtime_write_gate_cli_test.go
  cmd/xnix-runtime-go/test_repair_group_cli_test.go
  cmd/xnix-runtime-owner/main.go
  cmd/xnix-runtime-owner/main_test.go
  cmd/xnix-runtime-owner/version_test.go
  internal/runtime/appidentity/engine_catalog.go
  internal/runtime/appidentity/version_test.go
  internal/runtime/appidentity/run_plan.go
  internal/runtime/appidentity/application_readiness.go
  internal/runtime/appidentity/application_readiness_test.go
  internal/runtime/appidentity/application_upgrade_impact.go
  internal/runtime/appidentity/application_upgrade_impact_test.go
  internal/runtime/appidentity/runtime_policy_explanation_cards.go
  internal/runtime/appidentity/runtime_policy_explanation_cards_test.go
  internal/runtime/appidentity/compatibility_onboarding_checklist.go
  internal/runtime/appidentity/compatibility_onboarding_checklist_test.go
  internal/runtime/appidentity/offline_application_fixture_matrix.go
  internal/runtime/appidentity/offline_application_fixture_matrix_test.go
  internal/runtime/appidentity/desktop_activation_manifest.go
  internal/runtime/appidentity/desktop_activation_manifest_test.go
  internal/runtime/appidentity/desktop_safety_policy.go
  internal/runtime/appidentity/desktop_safety_policy_test.go
  internal/runtime/activation/stage.go
  internal/runtime/activation/stage_test.go
  internal/runtime/appidentity/install_plan.go
  internal/runtime/appidentity/install_plan_test.go
  internal/runtime/appidentity/kde_shell_surface.go
  internal/runtime/appidentity/kde_shell_surface_test.go
  internal/runtime/appidentity/window_identity_routes.go
  internal/runtime/appidentity/window_identity_routes_test.go
  internal/runtime/appidentity/windows_compatibility_workstreams.go
  internal/runtime/appidentity/windows_compatibility_workstreams_test.go
  internal/runtime/appidentity/state_root.go
  internal/runtime/appidentity/state_root_quota_retention.go
  internal/runtime/appidentity/state_root_quota_retention_test.go
  internal/runtime/appidentity/crash_hang_signal_summary.go
  internal/runtime/appidentity/crash_hang_signal_summary_test.go
  internal/runtime/appidentity/permission_evidence_audit.go
  internal/runtime/appidentity/permission_evidence_audit_test.go
  internal/runtime/appidentity/portal_permission_renewal.go
  internal/runtime/appidentity/portal_permission_renewal_test.go
  internal/runtime/appidentity/compatibility_backend_fallback.go
  internal/runtime/appidentity/compatibility_backend_fallback_test.go
  internal/runtime/appidentity/backend_adapter_contract.go
  internal/runtime/appidentity/backend_adapter_contract_test.go
  internal/runtime/appidentity/backend_adapter_contract_owner_route_audit.go
  internal/runtime/appidentity/backend_adapter_contract_owner_route_audit_test.go
  internal/runtime/appidentity/backend_adapter_contract_redacted_profile_audit.go
  internal/runtime/appidentity/backend_adapter_contract_redacted_profile_audit_test.go
  internal/runtime/appidentity/kde_search_visibility.go
  internal/runtime/appidentity/kde_search_visibility_test.go
  internal/runtime/appidentity/kde_notification_digest.go
  internal/runtime/appidentity/kde_notification_digest_test.go
  internal/runtime/appidentity/kde_offline_application_identity.go
  internal/runtime/appidentity/kde_offline_application_identity_test.go
  internal/runtime/appidentity/kde_fake_execution_evidence.go
  internal/runtime/appidentity/kde_fake_execution_evidence_test.go
  internal/runtime/appidentity/kde_fake_portal_evidence.go
  internal/runtime/appidentity/kde_fake_portal_evidence_test.go
  internal/runtime/appidentity/kde_snapshot_diagnostics_evidence.go
  internal/runtime/appidentity/kde_snapshot_diagnostics_evidence_test.go
  internal/runtime/appidentity/kde_backend_lifecycle_evidence.go
  internal/runtime/appidentity/kde_backend_lifecycle_evidence_test.go
  internal/runtime/appidentity/kde_restricted_launch_authorization.go
  internal/runtime/appidentity/kde_restricted_launch_authorization_test.go
  internal/runtime/appidentity/kde_test_launch_materialization.go
  internal/runtime/appidentity/kde_test_launch_materialization_test.go
  internal/runtime/appidentity/kde_test_launch_materialization_fanout.go
  internal/runtime/appidentity/kde_test_launch_materialization_fanout_test.go
  internal/runtime/appidentity/kde_test_launch_materialization_owner_route_audit.go
  internal/runtime/appidentity/kde_test_launch_materialization_owner_route_audit_test.go
  internal/runtime/appidentity/desktop_deactivation_dry_run.go
  internal/runtime/appidentity/desktop_deactivation_dry_run_test.go
  internal/runtime/appidentity/recipe_conflict_audit.go
  internal/runtime/appidentity/recipe_conflict_audit_test.go
  internal/runtime/recipe/signature.go
  internal/runtime/recipe/signature_test.go
  internal/runtime/image/restricted_smoke_packet.go
  internal/runtime/image/restricted_smoke_packet_test.go
  internal/runtime/appidentity/snapshot_restore_candidates.go
  internal/runtime/appidentity/snapshot_restore_candidates_test.go
  internal/runtime/appidentity/snapshot_plan.go
  internal/runtime/appidentity/portal_access_policy.go
  internal/runtime/appidentity/runtime_safety_plans_test.go
  internal/runtime/appidentity/identity.go
  internal/runtime/appidentity/identity_test.go
  internal/runtime/appidentity/registry.go
  internal/runtime/appidentity/registry_test.go
  internal/runtime/appidentity/ai_diagnostics.go
  internal/runtime/appidentity/ai_diagnostics_test.go
  internal/runtime/appidentity/diagnostic_history.go
  internal/runtime/appidentity/diagnostic_history_test.go
  internal/runtime/appidentity/support_bundle_manifest.go
  internal/runtime/appidentity/support_bundle_manifest_test.go
  internal/runtime/appidentity/support_case_timeline.go
  internal/runtime/appidentity/support_case_timeline_test.go
  internal/runtime/appidentity/multi_application_install_queue.go
  internal/runtime/appidentity/multi_application_install_queue_test.go
  internal/runtime/appidentity/execution_session_record_evidence.go
  internal/runtime/appidentity/repair_plan.go
  internal/runtime/appidentity/repair_plan_test.go
  internal/runtime/appidentity/runtime_write_gate.go
  internal/runtime/appidentity/runtime_write_gate_test.go
  internal/runtime/appidentity/runtime_route_convergence.go
  internal/runtime/appidentity/runtime_route_convergence_test.go
  internal/runtime/appidentity/kde_action_dependency_graph.go
  internal/runtime/appidentity/kde_action_dependency_graph_test.go
  cmd/xnix-runtime-go/kde_action_dependency_graph_cli_test.go
  internal/runtime/appidentity/kde_journey_evidence.go
  internal/runtime/appidentity/kde_journey_evidence_test.go
  cmd/xnix-runtime-go/kde_journey_evidence_cli_test.go
  internal/runtime/appidentity/test_plan.go
  internal/runtime/appidentity/test_plan_test.go
  internal/runtime/appidentity/test_result.go
  internal/runtime/appidentity/test_result_test.go
  internal/runtime/execution/execution.go
  internal/runtime/execution/execution_test.go
  internal/runtime/execution/ledger.go
  internal/runtime/execution/ledger_test.go
  internal/runtime/execution/session.go
  internal/runtime/execution/session_test.go
  internal/runtime/execution/authorization.go
  internal/runtime/execution/authorization_test.go
  internal/runtime/execution/restricted_materialization.go
  internal/runtime/execution/restricted_materialization_test.go
  internal/runtime/diagnostics/record.go
  internal/runtime/diagnostics/history.go
  internal/runtime/artifact/stage.go
  internal/runtime/owner/candidate.go
  internal/runtime/owner/candidate_test.go
  internal/runtime/owner/version_test.go
  internal/runtime/owner/dispatch.go
  internal/runtime/owner/dispatch_test.go
  internal/runtime/owner/route_checkpoint.go
  internal/runtime/owner/route_checkpoint_test.go
  internal/runtime/owner/kde_offline_identity_checkpoint.go
  internal/runtime/owner/kde_offline_identity_checkpoint_test.go
  internal/runtime/owner/lifecycle.go
  internal/runtime/owner/lifecycle_test.go
  internal/runtime/owner/smoke_batch.go
  internal/runtime/owner/smoke_batch_test.go
  internal/runtime/owner/restricted_smoke_receipt.go
  internal/runtime/owner/restricted_smoke_receipt_test.go
  internal/runtime/owner/restricted_smoke_receipt_fanout.go
  internal/runtime/owner/restricted_smoke_receipt_fanout_test.go
  internal/runtime/owner/restricted_smoke_receipt_lookup.go
  internal/runtime/owner/restricted_smoke_receipt_lookup_test.go
  internal/runtime/owner/restricted_smoke_receipt_fanout_owner_route_audit.go
  internal/runtime/owner/restricted_smoke_receipt_fanout_owner_route_audit_test.go
  internal/runtime/portal/broker.go
  internal/runtime/portal/broker_test.go
  internal/runtime/portal/request.go
  internal/runtime/portal/request_test.go
  internal/runtime/portal/read.go
  internal/runtime/portal/read_test.go
  internal/runtime/snapshot/store.go
  internal/runtime/snapshot/store_test.go
  internal/testversion/version.go
  internal/testversion/version_test.go
  libexec/xnix/compatd
  scripts/container.rb
  scripts/dbus_session_smoke.rb
  scripts/kde_center_dbus_smoke.rb
  scripts/kde_first_presence_smoke.rb
  scripts/offline_application_fixture_matrix.rb
  scripts/fetch_buildroot.rb
  scripts/full_smoke.rb
  scripts/restricted_product_smoke_packet.rb
  scripts/install_runtime_activation.rb
  scripts/prepare_ssh_test_key.rb
  scripts/runtime_activation_smoke.rb
  scripts/runtime_contract_drift_report.rb
  scripts/implementation_evidence_report.rb
  scripts/mainline_integration_review.rb
  scripts/release_evidence_index.rb
  scripts/merge_readiness_packet.rb
  scripts/runtime_owner_candidate_smoke.rb
  scripts/ssh_smoke.rb
  runtime/dbus/org.xnix.Compatibility1.xml
  runtime/dbus/org.xnix.Compatibility1.service
  runtime/dbus/xnix_compatd_smoke.c
  runtime/dbus/xnix_compatd_introspection.inc
  runtime/dbus/xnix_compatd_kde_center.inc
  runtime/dbus/xnix_compatd_runtime_models.inc
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
  test/test_full_smoke_report.rb
  test/test_full_smoke_script.rb
  test/test_restricted_product_smoke_packet.rb
  test/test_launch_request.rb
  test/test_notification_request.rb
  test/test_portal_access_policy.rb
  test/test_portal_request_model.rb
  test/test_runtime_contract.rb
  test/test_runtime_core.rb
  test/test_runtime_contract_drift_report.rb
  test/test_implementation_evidence_report.rb
  test/test_mainline_integration_review.rb
  test/test_release_evidence_index.rb
  test/test_offline_application_fixture_matrix.rb
  test/test_merge_readiness_packet.rb
  test/test_runtime_daemon.rb
  test/test_runtime_dispatch.rb
  test/test_runtime_live_owner_gate.rb
  test/test_runtime_method_parity_manifest.rb
  test/test_runtime_owner_candidate_smoke_script.rb
  test/test_runtime_owner_smoke_plan.rb
  test/test_runtime_service_binding.rb
  test/test_runtime_write_gate.rb
  test/test_runtime_activation.rb
  test/test_runtime_activation_install.rb
  test/test_runtime_activation_smoke_script.rb
  test/test_runtime_dbus_smoke_script.rb
  test/test_kde_center_model.rb
  test/test_kde_first_presence_smoke_script.rb
  test/test_kde_integration_status.rb
  test/test_kde_shell_integration_plan.rb
  test/test_kwin_window_rule.rb
  test/test_krunner_model.rb
  test/test_settings_model.rb
  test/test_settings_change_plan.rb
  test/test_task_manager_identity.rb
  test/test_tray_status_model.rb
  lib/xnix/image/kde_image.rb
  lib/xnix/image/boot_smoke.rb
  lib/xnix/image/serial_probe.rb
  lib/xnix/image/disk_build.rb
  scripts/build_kde_image.rb
  scripts/build_kde_disk.rb
  scripts/boot_kde_image.rb
  test/test_kde_image.rb
  test/test_kde_disk.rb
  image/kinoite/manifest.json
  image/kinoite/Containerfile
  image/kinoite/disk-config.json
  image/kinoite/config/portal/xnix-portals.conf
  image/kinoite/config/os-release.d/xnix.conf
  image/kinoite/config/systemd-preset/80-xnix.preset
  docs/kde-image-pipeline.md
  docs/release-evidence/v0.2.320-rc7-kde-product-smoke.json
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

dockerignore = read_project_file(".dockerignore").lines.map(&:strip)
assert(dockerignore.include?(".cache/"), "Docker build context must exclude the managed Buildroot cache")
assert(dockerignore.include?(".gocache/"), "Docker build context must exclude the local Go build cache")

assert(read_project_file("VERSION").strip == EXPECTED_VERSION, "VERSION must be #{EXPECTED_VERSION}")

REQUIRED_VERSION_MARKERS.each do |relative_path, marker|
  assert(read_project_file(relative_path).include?(marker), "#{relative_path} must reference #{EXPECTED_VERSION}")
end

test_version_source = read_project_file("internal/testversion/version.go")
%w[Current ReadFrom VERSION os.ReadFile].each do |token|
  assert(test_version_source.include?(token), "Go test version helper must include #{token}")
end
hardcoded_go_test_versions = Dir.glob(PROJECT_ROOT.join("{cmd,internal}/**/*_test.go").to_s).select do |path|
  File.read(path).match?(/\b0\.2\.\d+\b/)
end
assert(hardcoded_go_test_versions.empty?,
       "Go tests must read the canonical VERSION file instead of hardcoding project versions: #{hardcoded_go_test_versions.join(', ')}")

hardcoded_ruby_test_versions = Dir.glob(PROJECT_ROOT.join("test/*.rb").to_s).reject do |path|
  File.basename(path) == "test_runtime_core.rb"
end.select do |path|
  File.read(path).match?(/\["version"\]\s*==\s*"0\.2\./)
end
assert(hardcoded_ruby_test_versions.empty?,
       "Ruby model tests must read the canonical VERSION file instead of hardcoding project versions: #{hardcoded_ruby_test_versions.join(', ')}")

windows_workstreams = read_project_file("docs/claude-code-windows-compatibility-workstreams.md")
%w[CW1 CW2 CW3 CW4 CW5 CW6 CW7 CW8 CW9 CW10 CW11].each do |workstream_id|
  assert(windows_workstreams.include?("`#{workstream_id}`") || windows_workstreams.include?("## #{workstream_id}:"),
         "Windows compatibility workstreams must include #{workstream_id}")
end
[
  "Product target: best Linux desktop for existing Windows applications",
  "The first official shell is KDE Plasma",
  "KDE is the replaceable user shell. The Runtime is the product core.",
  "Go owns durable Runtime product logic",
  "Ruby is for tests, smoke scripts, reports, and lightweight developer tooling.",
  "Do not modify `docs/claude-code-implementation-packages.md`",
  "Production D-Bus ownership",
  "Runtime write methods",
  "Real Wine, Proton, VM, or backend launch",
  "Host-root mutation",
  "KDE-owned compatibility policy",
  "Start menu.",
  "Task manager.",
  "File manager.",
  "System tray.",
  "Notification center.",
  "AI Compatibility Center.",
  "Unified settings."
].each do |token|
  assert(windows_workstreams.include?(token), "Windows compatibility workstreams must include #{token}")
end
assert(windows_workstreams.include?("## Recommended First Wave"),
       "Windows compatibility workstreams must include a recommended first wave")
%w[CW1 CW2 CW3 CW10].each do |workstream_id|
  assert(windows_workstreams.include?("`#{workstream_id}`"),
         "Windows compatibility first wave must include #{workstream_id}")
end

windows_workstreams_go = read_project_file("internal/runtime/appidentity/windows_compatibility_workstreams.go")
[
  "xnix.runtime.windows_compatibility_workstreams.v1",
  "windows-compatibility-workstreams-preview",
  "GetWindowsCompatibilityWorkstreamsPreview",
  "KDE Plasma",
  "CW1",
  "CW2",
  "CW3",
  "CW10",
  "RuntimeOwned",
  "GoRuntimeBacked",
  "RubyCoreLogicAllowed",
  "KDEPolicyOwner",
  "BackendLaunchEnabled",
  "HostRootModified",
  "docs/claude-code-implementation-packages.md"
].each do |token|
  assert(windows_workstreams_go.include?(token), "Go Windows compatibility workstreams preview must include #{token}")
end

windows_workstreams_cli = read_project_file("cmd/xnix-runtime-go/windows_compatibility_commands.go")
assert(windows_workstreams_cli.include?("windows-compatibility-workstreams-preview"),
       "Windows compatibility workstreams CLI must expose its preview command")
assert(windows_workstreams_cli.include?("NewWindowsCompatibilityWorkstreamsPreview"),
       "Windows compatibility workstreams CLI must use the Go Runtime read model")

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
assert(dockerfile.include?("go build -o /usr/local/bin/xnix-runtime-owner"), "Dockerfile must install the Go Runtime owner candidate")
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
c_runtime_core_source = c_runtime_core_source +
                        read_project_file("runtime/core/xnix_runtime_core_desktop_icon.inc")
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
                     read_project_file("runtime/core/xnix_runtime_core_cli_desktop_icon.inc") +
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
%w[XnixRuntimeApplication xnix_runtime_application_count xnix_runtime_find_application XnixRuntimeEngine xnix_runtime_engine_count xnix_runtime_select_engine_for_mode XnixRuntimePortalPolicy xnix_runtime_portal_policy_count xnix_runtime_find_portal_policy XnixRuntimePortalRequestPlan xnix_runtime_portal_request_plan XnixRuntimeSnapshotPolicy xnix_runtime_snapshot_policy_count xnix_runtime_find_snapshot_policy XnixRuntimeStateRootPolicy xnix_runtime_state_root_policy_count xnix_runtime_find_state_root_policy XnixRuntimeInstallReadinessPolicy xnix_runtime_install_readiness_policy_count xnix_runtime_find_install_readiness_policy XnixRuntimeArtifactManifestPolicy xnix_runtime_artifact_manifest_policy_count xnix_runtime_find_artifact_manifest_policy XnixRuntimeAcquisitionPreflightPolicy xnix_runtime_acquisition_preflight_policy_count xnix_runtime_find_acquisition_preflight_policy XnixRuntimePackageSourcePolicy xnix_runtime_package_source_policy_count xnix_runtime_find_package_source_policy XnixRuntimeBackendBindingPolicy xnix_runtime_backend_binding_policy_count xnix_runtime_find_backend_binding_policy XnixRuntimeSettingsPolicy xnix_runtime_settings_policy_count xnix_runtime_find_settings_policy XnixRuntimeSettingsChangePolicy xnix_runtime_settings_change_policy_count xnix_runtime_find_settings_change_policy XnixRuntimeServiceBindingPolicy xnix_runtime_service_binding_policy XnixRuntimeLiveOwnerGatePolicy xnix_runtime_live_owner_gate_policy XnixRuntimeOwnerSmokePlanPolicy xnix_runtime_owner_smoke_plan_policy XnixRuntimeMethodParityManifestPolicy xnix_runtime_method_parity_manifest_policy XnixRuntimeRecipeTrustPolicy xnix_runtime_recipe_trust_policy XnixRuntimeRecipeInstallGate xnix_runtime_recipe_install_gate XnixRuntimeCompatibilityInstallPlan xnix_runtime_compatibility_install_plan XnixRuntimeCompatibilityActionQueue xnix_runtime_compatibility_action_queue XnixRuntimeCompatibilityActionReviewReceipt xnix_runtime_compatibility_action_review_receipt XnixRuntimeCompatibilityCenterSummary xnix_runtime_compatibility_center_summary XnixRuntimeCompatibilityRunPlan xnix_runtime_compatibility_run_plan XnixRuntimeDesktopActivationManifest xnix_runtime_desktop_activation_manifest XnixRuntimeKDEIntegrationStatus xnix_runtime_kde_integration_status XnixRuntimeKDEApplicationSurfacePlan XnixRuntimeKDEApplicationSurfaceEntryPoint xnix_runtime_kde_application_surface_plan XnixRuntimeDesktopEntryPlan xnix_runtime_desktop_entry_plan XnixRuntimeDesktopIconPlan xnix_runtime_desktop_icon_plan XnixRuntimeTaskManagerIdentityPlan xnix_runtime_task_manager_identity_plan XnixRuntimeKWinWindowRulePlan xnix_runtime_kwin_window_rule_plan XnixRuntimeFileAssociationPlan xnix_runtime_file_association_plan XnixRuntimeNotificationPlan xnix_runtime_notification_plan XnixRuntimeTrayStatusPlan xnix_runtime_tray_status_plan XnixRuntimeKRunnerMatch XnixRuntimeKRunnerQueryPlan xnix_runtime_krunner_query_plan XnixRuntimeAIDiagnosticInput xnix_runtime_ai_diagnostic_input].each do |token|
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
%w[xnix_runtime_write_gate blocked-until-production-backend InstallRecipe Launch CreateSnapshot RestoreSnapshot xnix_runtime_find_application org.xnix.sample.notepad application/x-xnix-txt xnix_runtime_select_engine_for_mode automatic-managed local-compatibility-engine isolated-compatibility-engine xnix_runtime_find_portal_policy org.freedesktop.portal.FileChooser org.freedesktop.portal.Camera remote-desktop selected-files xnix_runtime_portal_request_plan portal_request_method_for_operation handle_token permission_granted host_permission_changed OpenFile RequestClipboard AccessCamera xnix_runtime_find_snapshot_policy before-repair before-engine-change desktop_activation_receipts xnix_runtime_find_state_root_policy application-data runtime-metadata diagnostic-cache xnix_runtime_find_install_readiness_policy resolve-artifact-manifest verify-artifact-digests prepare-package-source allocate-application-state xnix_runtime_find_artifact_manifest_policy runtime-launch-metadata local-execution-artifacts isolated-environment-artifacts manifest-signature-verification artifact-digest-verification xnix_runtime_find_acquisition_preflight_policy package-source-ready signed-artifact-manifest runtime-cache-space network-policy-review rollback-marker xnix_runtime_find_package_source_policy os-managed-compatibility-packages runtime-managed-toolcache isolated-environment-template-catalog signed-source-verification source-policy-review runtime-cache-quota offline-fallback xnix_runtime_find_backend_binding_policy engine-package-source application-state-root portal-policy-review snapshot-baseline xnix_runtime_find_settings_policy run-mode resource-access devices snapshots settings_change_policies xnix_runtime_find_settings_change_policy validate-setting persist-runtime-setting service_binding_policy xnix_runtime_service_binding_policy dbus-service-activation live-dbus-owner live_owner_gate_policy xnix_runtime_live_owner_gate_policy production-recipe-trust owner_smoke_plan_policy xnix_runtime_owner_smoke_plan_policy validate-activation-files assert-stable-bus-name method_parity_manifest_policy xnix_runtime_method_parity_manifest_policy recipe_trust_policy xnix_runtime_recipe_trust_policy production_recipe_install_requirements xnix_runtime_compatibility_install_plan xnix_runtime_compatibility_action_queue xnix_runtime_compatibility_action_review_receipt xnix_runtime_compatibility_center_summary compatibility-center-summary known_issue_count repair_record_state backend_launch_enabled xnix_runtime_compatibility_run_plan xnix_runtime_desktop_activation_manifest xnix_runtime_kde_integration_status kde-integration-status standard-desktop-entry window-identity-and-restore xnix_runtime_desktop_entry_plan xnix_runtime_task_manager_identity_plan xnix_runtime_file_association_plan xnix_runtime_notification_plan xnix_runtime_tray_status_plan xnix_runtime_krunner_query_plan xnix_runtime_ai_diagnostic_input ai-diagnostic-input safe_for_ai_diagnostics desktop-activation desktop-entry-plan desktop-icon-plan task-manager-identity-plan file-association-plan notification-plan tray-status-plan krunner-query-plan query_execution_enabled approval-required install-failed repair-applied mode-changed action_execution_enabled repair_execution_enabled settings_persistence_enabled live_backend_bridge_enabled bridge_configuration_persisted mimeapps.list applications:xnix-org.xnix.sample.notepad.desktop xnix-compat-open standard_desktop_entry launch_uses_runtime raw_windows_executable_exposed compatibility_storage_path_exposed pinning_allowed restore_allowed skip_taskbar show_in_switcher identity-only window_manager_policy_only portal_required_for_file_open overwrite_existing_mimeapps portal_policy_required snapshot_before_risky_change diagnostics_required execution_request_created review-install-readiness review-settings-change review-ai-repair verify-runtime-service review-portal-policy decision_recorded reviewed approved deferred rejected signed recipe validation is not enabled registry.digest registry.signature registry.development GetDesktopActivationManifest GetKDEIntegrationStatus GetDesktopEntryPlan GetDesktopIconPlan GetTaskManagerIdentityPlan GetFileAssociationPlan GetNotificationPlan GetTrayStatus GetKRunnerQueryPlan GetCompatibilityCenterSummary GetKDECenterPage GetKDECenterPageSections GetKDECenterPageSectionDetail GetPortalRequestPlan GetAIDiagnosticInput GetRuntimeMethodParityManifest GetRuntimeWriteGate].each do |token|
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
%w[business_logic_runtime tests-and-development-tools application_catalog_owner engine_catalog_owner portal_policy_owner portal_request_plan_owner ai_diagnostic_input_owner ai_diagnostic_signal_count desktop_entry_plan_owner task_manager_identity_plan_owner kde_integration_status_owner kde_integration_entry_point_count file_association_plan_owner notification_plan_owner tray_status_plan_owner krunner_query_plan_owner krunner_query_plan_count compatibility_center_summary_owner snapshot_policy_owner state_root_policy_owner install_readiness_policy_owner artifact_manifest_policy_owner acquisition_preflight_policy_owner package_source_policy_owner backend_binding_policy_owner backend_lifecycle_owner backend_environment_plan_owner settings_policy_owner settings_change_policy_owner service_binding_policy_owner live_owner_gate_policy_owner owner_smoke_plan_policy_owner method_parity_manifest_policy_owner recipe_trust_policy_owner recipe_install_gate_owner compatibility_install_plan_owner compatibility_action_queue_owner compatibility_action_review_receipt_owner compatibility_center_summary_count compatibility_run_plan_owner desktop_activation_manifest_owner compatibility-center-action-review-receipt compatibility-center-summary receipt_id required_runtime_gate compatibility-run desktop-activation kde-integration-status portal-request-plan ai-diagnostic-input desktop-entry-plan desktop-icon-plan task-manager-identity-plan file-association-plan notification-plan tray-status-plan krunner-query-plan list-applications get-application list-engines select-engine list-portal-policies portal-policy portal-request-plan list-snapshot-policies snapshot-policy list-state-root-policies state-root-policy list-install-readiness-policies install-readiness-policy compatibility-install-plan compatibility-action-queue compatibility-action-review compatibility-center-summary compatibility-run-plan desktop-activation-manifest kde-integration-status desktop-entry-plan desktop-icon-plan task-manager-identity-plan file-association-plan notification-plan tray-status-plan krunner-query-plan ai-diagnostic-input list-artifact-manifest-policies artifact-manifest-policy list-acquisition-preflight-policies acquisition-preflight-policy list-package-source-policies package-source-policy list-backend-binding-policies backend-binding-policy backend-lifecycle backend-environment-plan list-settings-policies settings-policy list-settings-change-policies settings-change-policy runtime-service-binding runtime-live-owner-gate runtime-owner-smoke-plan runtime-method-parity-manifest recipe-trust-policy recipe-install-gate write-gate].each do |token|
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

go_backend_lifecycle_source = read_project_file("internal/runtime/appidentity/backend_lifecycle.go")
%w[BackendLifecyclePreview BackendLifecyclePreviewWithStateRoot BackendLifecycleRecord RecordBackendLifecycleState backend-lifecycle-preview xnix.runtime.backend_lifecycle.v1 xnix.runtime.backend_lifecycle_record.v1 GetBackendLifecycle GetBackendLifecyclePreview go-runtime-state-root-backend-lifecycle environment-state-root StateRootBacked StateRootPathExposed EnvironmentProfile SatisfiedGates PendingGates RepairHints BlockReason RelativePath].each do |token|
  assert(go_backend_lifecycle_source.include?(token), "Go Runtime backend lifecycle preview must include #{token}")
end
%w[RuntimeOwned GoRuntimeBacked KDEPolicyOwner BackendBindingReady LaunchEnabled ExecutionRequestCreated BackendProcessStarted LocalBackendStarted IsolatedBackendStarted StateRootReady PortalReviewRequired SnapshotRequired HostRootModified NetworkRequired BackendDetailsExposed].each do |token|
  assert(go_backend_lifecycle_source.include?(token), "Go Runtime backend lifecycle preview must expose #{token}")
end
assert(go_backend_lifecycle_source.include?("environment.New"), "Go Runtime backend lifecycle preview must consume the environment state-root store")
assert(go_backend_lifecycle_source.include?("RequiredGates"), "Go Runtime backend lifecycle preview must expose pending environment gates")
assert(go_backend_lifecycle_source.include?("validateNoBackendTerms"), "Go Runtime backend lifecycle preview must hide backend terms")
%w[backend-lifecycle-record runBackendLifecycleRecord].each do |token|
  assert(read_project_file("cmd/xnix-runtime-go/main.go").include?(token), "Go Runtime backend lifecycle record CLI must include #{token}")
end
assert(read_project_file("cmd/xnix-runtime-go/backend_group_cli_test.go").include?("TestBackendLifecycleRecordCommandPersistsStateTransitions"), "Go Runtime backend lifecycle CLI tests must persist state-root transitions")

go_backend_manager_source = read_project_file("internal/runtime/appidentity/backend_manager.go")
%w[BackendManagerPreview BackendManagerRecord ManagedCompatibilityBackend UserFacingBackendProfile xnix.runtime.backend_manager.v1 xnix.runtime.backend_manager_record.v1 backend-manager-preview backend-manager-inventory-record go-runtime-backend-manager go-runtime-state-root-backend-manager wine proton windows-vm WineManaged ProtonManaged WindowsVMManaged BackendInstallEnabled BackendDownloadEnabled BackendLaunchEnabled BackendProcessStarted VMProcessStarted BackendDetailsExposedToKDE StateRootPathExposed HostRootModified PrivilegedContainerRequired LoadBackendManagerRecord validateBackendManagerRecord backendManagerRecordUnsafe].each do |token|
  assert(go_backend_manager_source.include?(token), "Go Runtime backend manager preview must include #{token}")
end
%w[unsupported\ schema identity,\ path,\ or\ count\ mismatch digest\ mismatch unsafe\ enabled\ gates record\ directory\ must\ be\ a\ real\ directory].each do |token|
  assert(go_backend_manager_source.include?(token), "Go Runtime backend manager readback must validate #{token}")
end
%w[backend-manager-preview backend-manager-record runBackendManagerPreview runBackendManagerRecord].each do |token|
  assert(read_project_file("cmd/xnix-runtime-go/main.go").include?(token), "Go Runtime backend manager CLI must include #{token}")
end
assert(read_project_file("internal/runtime/appidentity/backend_manager_test.go").include?("TestBackendManagerPreviewKeepsUserFacingProfilesBackendSafe"), "Go Runtime backend manager tests must keep user-facing profiles backend-safe")
assert(read_project_file("internal/runtime/appidentity/backend_manager_test.go").include?("TestRecordBackendManagerPreviewPersistsStateRootInventory"), "Go Runtime backend manager tests must persist state-root inventory")
assert(read_project_file("internal/runtime/appidentity/backend_manager_test.go").include?("TestBackendManagerRecordRejectsTamperingAndManagedPathSymlink"), "Go Runtime backend manager tests must reject tampering and managed-path symlinks")

go_backend_adapter_contract_source = read_project_file("internal/runtime/appidentity/backend_adapter_contract.go")
%w[BackendAdapterContractPreview BackendAdapterContract BackendAdapterProfile BackendAdapterCounts NewBackendAdapterContractPreview xnix.runtime.backend_adapter_contract.v1 backend-adapter-contract-preview compatibility-backend-adapter-noop-contract go-runtime-backend-manager+adapter-noop-boundary GetBackendAdapterContract GetBackendAdapterContractPreview wine proton windows-vm NoopImplementation AdapterInvocationEnabled BackendInstallEnabled BackendDownloadEnabled BackendLaunchEnabled BackendProcessStarted VMProcessStarted CommandMaterialized ExecutablePathResolved RawCommandExposed ProfilePathExposed StateRootPathExposed BackendDetailsExposedToKDE HostRootModified PrivilegedContainerRequired test-only-materialization].each do |token|
  assert(go_backend_adapter_contract_source.include?(token), "Go Runtime backend adapter contract preview must include #{token}")
end
%w[backend-adapter-contract-preview runBackendAdapterContractPreview].each do |token|
  assert(read_project_file("cmd/xnix-runtime-go/main.go").include?(token) || read_project_file("cmd/xnix-runtime-go/backend_adapter_contract_commands.go").include?(token), "Go Runtime backend adapter contract CLI must include #{token}")
end
assert(read_project_file("internal/runtime/appidentity/backend_adapter_contract_test.go").include?("TestBackendAdapterContractKDEProjectionIsBackendSafe"), "Go Runtime backend adapter contract tests must keep KDE projection backend-safe")
assert(read_project_file("cmd/xnix-runtime-go/backend_adapter_contract_cli_test.go").include?("TestBackendAdapterContractPreviewCommandRendersNoopBoundary"), "Go Runtime backend adapter contract CLI tests must render the no-op boundary")
go_backend_adapter_contract_owner_route_audit_source = read_project_file("internal/runtime/appidentity/backend_adapter_contract_owner_route_audit.go")
%w[BackendAdapterContractOwnerRouteAuditPreview NewBackendAdapterContractOwnerRouteAuditPreview xnix.runtime.backend_adapter_contract_owner_route_audit.v1 backend-adapter-contract-owner-route-audit-preview GetBackendAdapterContractOwnerRouteAudit GetBackendAdapterContractOwnerRouteAuditPreview adapter-contract-owner-route-audit GetBackendAdapterProfileAudit redacted-profile-route-ready redacted-profile-route-ready-full-contract-fixture-local redacted-profile-route-smoke-covered redacted-profile-route-smoke-covered-full-contract-fixture-local-production-dbus-blocked redacted-adapter-profile-owner-smoke-coverage redacted-adapter-profile-production-dbus-gate-review cli-preview-registered go-read-model-present fixture-consumption-present owner-route-absent production-dbus-absent redacted-route-missing internal-detail-boundary owner-smoke-coverage-present route-decision unsafe-gates-closed].each do |token|
  assert(go_backend_adapter_contract_owner_route_audit_source.include?(token), "Go Runtime backend adapter contract owner-route audit must include #{token}")
end
%w[FixtureMatrixConsumesContract OwnerDispatchRoutePresent ProductionDBusMethodPresent FullContractContainsAdapterIDs KDEFacingProjectionPresent RedactedProfileRoutePresent RequiresCallerRoot OwnerSmokeCoverageReady AdapterInvocationEnabled BackendInstallEnabled BackendDownloadEnabled BackendLaunchEnabled BackendProcessStarted CommandMaterialized ExecutablePathResolved OwnerLocalRouteCandidateReady ProductionDBusExposureReady SystemServiceStarted SessionBusClaimed ProductionBusClaimed WriteMethodsEnabled RuntimeWritesEnabled HostRootModified StateRootPathExposed BackendDetailsExposed].each do |token|
  assert(go_backend_adapter_contract_owner_route_audit_source.include?(token), "Go Runtime backend adapter contract owner-route audit must expose safety gate #{token}")
end
assert(go_backend_adapter_contract_owner_route_audit_source.include?("validateNoBackendTerms"), "Go Runtime backend adapter contract owner-route audit must hide backend terms")
go_backend_adapter_contract_owner_route_audit_test_source = read_project_file("internal/runtime/appidentity/backend_adapter_contract_owner_route_audit_test.go")
%w[TestBackendAdapterContractOwnerRouteAuditKeepsContractFixtureLocal TestBackendAdapterContractOwnerRouteAuditFailsClosedWithoutSources redacted-profile-route-smoke-covered owner-route-absent production-dbus-absent redacted-route-missing internal-detail-boundary owner-smoke-coverage-present].each do |token|
  assert(go_backend_adapter_contract_owner_route_audit_test_source.include?(token), "Go Runtime backend adapter contract owner-route audit tests must include #{token}")
end
go_backend_adapter_contract_owner_route_audit_cli_source = read_project_file("cmd/xnix-runtime-go/backend_adapter_contract_commands.go") +
                                                          read_project_file("cmd/xnix-runtime-go/backend_adapter_contract_owner_route_audit_cli_test.go") +
                                                          read_project_file("cmd/xnix-runtime-go/main.go")
%w[backend-adapter-contract-owner-route-audit-preview runBackendAdapterContractOwnerRouteAuditPreview TestBackendAdapterContractOwnerRouteAuditPreviewCommand route_decision fixture_matrix_consumes_contract owner_dispatch_route_present production_dbus_method_present owner_local_route_candidate_ready owner_smoke_coverage_ready].each do |token|
  assert(go_backend_adapter_contract_owner_route_audit_cli_source.include?(token), "Go Runtime backend adapter contract owner-route audit CLI must include #{token}")
end

go_backend_adapter_redacted_profile_audit_source = read_project_file("internal/runtime/appidentity/backend_adapter_contract_redacted_profile_audit.go")
%w[BackendAdapterRedactedProfileAuditPreview BackendAdapterRedactedProfile NewBackendAdapterRedactedProfileAuditPreview xnix.runtime.backend_adapter_redacted_profile_audit.v1 backend-adapter-redacted-profile-audit-preview owner-local-redacted-adapter-profile-audit GetBackendAdapterProfileAudit GetBackendAdapterProfileAuditPreview redacted-profile-route-ready contract-profile-projection-consumed internal-adapter-ids-redacted owner-local-route-candidate full-contract-fixture-local production-dbus-not-claimed no-caller-state-root unsafe-gates-closed].each do |token|
  assert(go_backend_adapter_redacted_profile_audit_source.include?(token), "Go Runtime backend adapter redacted profile audit must include #{token}")
end
%w[InternalAdapterIDsRedacted InternalProfilePathsRedacted OwnerLocalRouteCandidateReady FullContractFixtureLocal ProductionDBusExposureReady CallerStateRootRequired AdapterInvocationEnabled BackendInstallEnabled BackendDownloadEnabled BackendLaunchEnabled BackendProcessStarted CommandMaterialized ExecutablePathResolved HostRootModified StateRootPathExposed RawCommandExposed RawExecutableExposed BackendDetailsExposed].each do |token|
  assert(go_backend_adapter_redacted_profile_audit_source.include?(token), "Go Runtime backend adapter redacted profile audit must expose safety gate #{token}")
end
assert(go_backend_adapter_redacted_profile_audit_source.include?("validateNoBackendTerms"), "Go Runtime backend adapter redacted profile audit must hide backend terms")
go_backend_adapter_redacted_profile_audit_test_source = read_project_file("internal/runtime/appidentity/backend_adapter_contract_redacted_profile_audit_test.go")
%w[TestBackendAdapterRedactedProfileAuditPreviewDefinesOwnerRouteSplit TestBackendAdapterRedactedProfileAuditOutputIsRedacted TestBackendAdapterRedactedProfileAuditFailsClosedWithoutOwnerRouteSource automatic performance-priority compatibility-priority owner-local-route-candidate].each do |token|
  assert(go_backend_adapter_redacted_profile_audit_test_source.include?(token), "Go Runtime backend adapter redacted profile audit tests must include #{token}")
end
go_backend_adapter_redacted_profile_audit_cli_source = read_project_file("cmd/xnix-runtime-go/backend_adapter_contract_commands.go") +
                                                      read_project_file("cmd/xnix-runtime-go/backend_adapter_contract_redacted_profile_audit_cli_test.go") +
                                                      read_project_file("cmd/xnix-runtime-go/main.go")
%w[backend-adapter-redacted-profile-audit-preview runBackendAdapterRedactedProfileAuditPreview TestBackendAdapterRedactedProfileAuditPreviewCommand owner_local_route_candidate_ready full_contract_fixture_local production_dbus_exposure_ready].each do |token|
  assert(go_backend_adapter_redacted_profile_audit_cli_source.include?(token), "Go Runtime backend adapter redacted profile audit CLI must include #{token}")
end

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

go_runtime_service_binding_source = read_project_file("internal/runtime/appidentity/runtime_service_binding.go")
%w[RuntimeServiceBindingPreview runtime-service-binding-preview xnix.runtime.service_binding.v1 GetRuntimeServiceBinding GetRuntimeServiceBindingPreview dbus-service-activation systemd-service-hardening live-dbus-owner].each do |token|
  assert(go_runtime_service_binding_source.include?(token), "Go Runtime service binding preview must include #{token}")
end
%w[RuntimeOwned GoRuntimeBacked ActivationBindingReady LiveDBusOwnerReady SmokeAdapterAvailable ServiceStarted ProductionBusClaimed KDEMayClaimRuntimeOwnership NetworkRequired HostRootModified PrivilegedContainerRequired BackendDetailsExposed].each do |token|
  assert(go_runtime_service_binding_source.include?(token), "Go Runtime service binding preview must expose #{token}")
end
assert(go_runtime_service_binding_source.include?("validateNoBackendTerms"), "Go Runtime service binding preview must hide backend terms")
assert(go_runtime_service_binding_source.include?("runtimeServiceBindingPackagedWrapper"), "Go Runtime service binding preview must track the packaged wrapper")

go_runtime_service_binding_cli_source = read_project_file("cmd/xnix-runtime-go/main.go")
assert(go_runtime_service_binding_cli_source.include?("runtime-service-binding-preview"), "Go Runtime CLI must expose Runtime service binding preview")

go_runtime_service_activation_preflight_source = read_project_file("internal/runtime/appidentity/runtime_service_activation_preflight.go")
%w[RuntimeServiceActivationPreflightPreview RuntimeServiceActivationProductionDBusGate runtime-service-activation-preflight-preview xnix.runtime.service_activation_preflight.v1 GetRuntimeServiceActivationPreflight GetRuntimeServiceActivationPreflightPreview runtime-service-binding-preview runtime-owner-readiness-preview runtime-owner-smoke-plan-preview production-dbus-gate-review-preview production-dbus-human-authorization-preflight-preview production-runtime-service-activation-preflight activation-binding read-only-method-parity owner-smoke-plan production-dbus-gate-review human-authorization-preflight human-authorization-receipt write-method-gate kde-ownership-boundary host-safety-boundary long-running-runtime-owner restricted-owner-smoke production-bus-claim production-recipe-trust restricted-owner-smoke-ready production-activation-blocked production-dbus-gate-review-consumed-activation-still-blocked production-dbus-gate-review-missing].each do |token|
  assert(go_runtime_service_activation_preflight_source.include?(token), "Go Runtime service activation preflight preview must include #{token}")
end
%w[ProductionActivationReady RestrictedSmokeReady ProductionDBusGateReady HumanAuthorizationPreflightReady HumanAuthorizationRequired HumanAuthorizationGranted AuthorizationReceiptAccepted RuntimeOwned GoRuntimeBacked KDEPolicyOwner KDEMayClaimRuntimeOwnership SystemServiceStarted SessionBusClaimed ProductionBusClaimed WriteMethodsEnabled RuntimeWritesEnabled BackendLaunchEnabled NetworkRequired HostRootModified PrivilegedContainerRequired BackendDetailsExposed].each do |token|
  assert(go_runtime_service_activation_preflight_source.include?(token), "Go Runtime service activation preflight preview must expose #{token}")
end
assert(go_runtime_service_activation_preflight_source.include?("NewRuntimeServiceBindingPreview"), "Go Runtime service activation preflight must derive from the service binding preview")
assert(go_runtime_service_activation_preflight_source.include?("NewRuntimeOwnerReadinessPreview"), "Go Runtime service activation preflight must derive from owner readiness")
assert(go_runtime_service_activation_preflight_source.include?("NewRuntimeOwnerSmokePlanPreview"), "Go Runtime service activation preflight must derive from owner smoke planning")
assert(go_runtime_service_activation_preflight_source.include?("runtimeServiceActivationProductionDBusGate"), "Go Runtime service activation preflight must consume production D-Bus gate sources")
assert(go_runtime_service_activation_preflight_source.include?("validateNoBackendTerms"), "Go Runtime service activation preflight must hide backend terms")

go_runtime_service_activation_preflight_cli_source = read_project_file("cmd/xnix-runtime-go/main.go")
assert(go_runtime_service_activation_preflight_cli_source.include?("runtime-service-activation-preflight-preview"), "Go Runtime CLI must expose Runtime service activation preflight preview")

runtime_activation_installer_source = read_project_file("scripts/install_runtime_activation.rb")
%w[SOURCE_RUNTIME_LIB SOURCE_RECIPE_DIR SOURCE_DBUS_CONTRACT SOURCE_DBUS_SMOKE SOURCE_DBUS_SMOKE_INTROSPECTION SOURCE_DBUS_SMOKE_KDE_CENTER SOURCE_DBUS_SMOKE_RUNTIME_MODELS SOURCE_SESSION_SMOKE usr/lib/xnix usr/runtime/recipes].each do |token|
  assert(runtime_activation_installer_source.include?(token), "Runtime activation installer must include #{token}")
end

runtime_activation_smoke_source = read_project_file("scripts/runtime_activation_smoke.rb")
%w[usr/libexec/xnix/compatd usr/lib/xnix/compatibility/runtime_daemon.rb GetRuntimeMethodParityManifest GetRuntimeWriteGate WriteMethodDisabled].each do |token|
  assert(runtime_activation_smoke_source.include?(token), "Runtime activation smoke must include #{token}")
end

kde_first_presence_smoke_source = read_project_file("scripts/kde_first_presence_smoke.rb")
%w[desktop-safety-policy-preview desktop-identity-plan desktop-entry-preview mimeapps-preview desktop-activation-bundle-preview desktop-activation-staging-preview desktop-activation-transaction-preview desktop-activation-status-preview kde-entrypoints-preview kde-action-card-deck-preview kde-action-dependency-graph-preview kde-journey-evidence-preview compatibility-onboarding-checklist-preview support-bundle-manifest-preview multi-application-install-queue-preview runtime-policy-explanation-cards-preview kde-center-page-preview kde-center-page-sections-preview kde-center-page-section-detail-preview file-open-preview dolphin-drop-preview dolphin-ai-analysis-preview window-identity-preview tray-status-preview notification-preview settings-preview runtime-owner-route-manifest-preview runtime-route-convergence-preview runtime-method-parity-manifest-preview runtime-write-gate-preview ai-diagnostic-input-preview ai-diagnostic-recommendation-preview ai-repair-approval-gate-preview].each do |token|
  assert(kde_first_presence_smoke_source.include?(token), "KDE-first presence smoke must inspect #{token}")
end
%w[launcher task-manager file-manager system-tray notifications compatibility-center settings].each do |token|
  assert(kde_first_presence_smoke_source.include?(token), "KDE-first presence smoke must assert #{token}")
end
%w[host_root_modified network_required privileged_container_required backend_details_exposed raw_command_exposed raw_windows_executable_exposed execution_started backend_launch_enabled launch_enabled request_object_created permission_granted file_content_read file_paths_exposed ai_provider_call_enabled].each do |token|
  assert(kde_first_presence_smoke_source.include?(token), "KDE-first presence smoke must enforce #{token}")
end
%w[
  desktop-entry
  dolphin-service-menu
  mimeapps-list
  desktop-integration-manifest
  desktop-activation-receipt
  GOCACHE
  --format
  kde-first-presence-smoke
  xnix.kde_first_presence_smoke.v1
  report_type
  preview_commands
  route_baseline
  settings_field_ids
  forbidden_user_terms
  xnix.runtime.desktop_safety_policy.v1
  assert_desktop_safety_policy
  assert_onboarding_checklist
  safety_false_keys
  prefix
  bottle
  JSON.pretty_generate
  render_markdown
].each do |token|
  assert(kde_first_presence_smoke_source.include?(token), "KDE-first presence smoke must include #{token}")
end
assert(!kde_first_presence_smoke_source.include?("scripts/container.rb"), "KDE-first presence smoke must not run Docker")
assert(!kde_first_presence_smoke_source.include?("boot-system"), "KDE-first presence smoke must not boot QEMU")

desktop_safety_policy_source = read_project_file("internal/runtime/appidentity/desktop_safety_policy.go")
%w[xnix.runtime.desktop_safety_policy.v1 kde-first-user-facing-safety-policy forbidden_user_terms settings_field_ids BackendTerminologyHidden WriteMethodsEnabled RealPortalTransport AIProviderCallEnabled].each do |token|
  assert(desktop_safety_policy_source.include?(token), "Runtime desktop safety policy must include #{token}")
end
desktop_safety_policy_cli_source = read_project_file("cmd/xnix-runtime-go/main.go") + read_project_file("cmd/xnix-runtime-go/desktop_safety_policy_commands.go")
%w[desktop-safety-policy-preview NewDesktopSafetyPolicyPreview].each do |token|
  assert(desktop_safety_policy_cli_source.include?(token), "Runtime desktop safety policy CLI must include #{token}")
end

application_readiness_source = read_project_file("internal/runtime/appidentity/application_readiness.go")
%w[xnix.runtime.application_readiness.v1 application-readiness-preview runtime-application-readiness-evidence-graph RecipeTrustDecision WriteGateDecision RealPortalTransportEnabled BackendLaunchEnabled StateRootPathExposed RawCommandExposed].each do |token|
  assert(application_readiness_source.include?(token), "Runtime application readiness graph must include #{token}")
end
application_readiness_cli_source = read_project_file("cmd/xnix-runtime-go/main.go") + read_project_file("cmd/xnix-runtime-go/application_readiness_commands.go")
%w[application-readiness-preview ApplicationReadinessPreview artifact-receipt portal-operation snapshot-reason].each do |token|
  assert(application_readiness_cli_source.include?(token), "Runtime application readiness CLI must include #{token}")
end

compatibility_onboarding_source = read_project_file("internal/runtime/appidentity/compatibility_onboarding_checklist.go")
%w[CompatibilityOnboardingChecklistPreview compatibility-onboarding-checklist-preview xnix.runtime.compatibility_onboarding_checklist.v1 first-run-compatibility-onboarding GetCompatibilityOnboardingChecklist GetCompatibilityOnboardingChecklistPreview runtime-owner-readiness recipe-trust artifact-staging backend-lifecycle portal-review snapshot-baseline diagnostics-privacy kde-entry-points production-activation ready needs-review missing-evidence blocked not-yet-implemented RuntimeWriteMethodsEnabled RequestObjectsCreated PermissionGrantsCreated ArtifactStaged SettingsPersisted BackendProcessStarted LaunchEnabled NetworkRequired HostPackageManagerInvoked HostRootModified PrivilegedContainerRequired AIProviderCalled StateRootPathExposed RawExecutableExposed RawCommandExposed FileContentRead BackendDetailsExposed].each do |token|
  assert(compatibility_onboarding_source.include?(token), "Runtime compatibility onboarding checklist must include #{token}")
end
%w[NewRuntimeOwnerReadinessPreview ApplicationReadinessPreview NewPortalAccessPolicyPreview NewSnapshotPlanPreview AIDiagnosticInputPreview NewDesktopSafetyPolicyPreview validateNoBackendTerms].each do |token|
  assert(compatibility_onboarding_source.include?(token), "Runtime compatibility onboarding checklist must derive evidence from #{token}")
end
compatibility_onboarding_test_source = read_project_file("internal/runtime/appidentity/compatibility_onboarding_checklist_test.go")
%w[TestCompatibilityOnboardingChecklistPreviewAggregatesFirstRunEvidence TestCompatibilityOnboardingChecklistPreviewRejectsMalformedInput runtime-owner-readiness recipe-trust artifact-staging backend-lifecycle portal-review snapshot-baseline diagnostics-privacy kde-entry-points production-activation not-yet-implemented program files].each do |token|
  assert(compatibility_onboarding_test_source.include?(token), "Runtime compatibility onboarding checklist tests must include #{token}")
end
compatibility_onboarding_cli_source = read_project_file("cmd/xnix-runtime-go/main.go") + read_project_file("cmd/xnix-runtime-go/compatibility_onboarding_commands.go")
%w[compatibility-onboarding-checklist-preview runCompatibilityOnboardingChecklistPreview CompatibilityOnboardingChecklistPreview artifact-receipt portal-operation snapshot-reason].each do |token|
  assert(compatibility_onboarding_cli_source.include?(token), "Runtime compatibility onboarding CLI must include #{token}")
end
compatibility_onboarding_cli_test_source = read_project_file("cmd/xnix-runtime-go/compatibility_onboarding_cli_test.go")
%w[TestCompatibilityOnboardingChecklistPreviewCommand compatibility-onboarding-checklist-preview xnix.runtime.compatibility_onboarding_checklist.v1 GetCompatibilityOnboardingChecklist GetCompatibilityOnboardingChecklistPreview runtime_write_methods_enabled request_objects_created permission_grants_created artifact_staged settings_persisted backend_process_started launch_enabled host_package_manager_invoked state_root_path_exposed raw_command_exposed file_content_read].each do |token|
  assert(compatibility_onboarding_cli_test_source.include?(token), "Runtime compatibility onboarding CLI tests must include #{token}")
end

runtime_policy_explanation_cards_source = read_project_file("internal/runtime/appidentity/runtime_policy_explanation_cards.go")
%w[RuntimePolicyExplanationCardsPreview RuntimePolicyExplanationCard RuntimePolicyExplanationCounts runtime-policy-explanation-cards-preview kde-runtime-policy-explanation-card-deck xnix.runtime.policy_explanation_cards.v1 GetRuntimePolicyExplanationCards GetRuntimePolicyExplanationCardsPreview install launch execution portal-permission snapshot diagnostics repair settings desktop-activation backend-readiness unsupported-production-route].each do |token|
  assert(runtime_policy_explanation_cards_source.include?(token), "Runtime policy explanation cards must include #{token}")
end
%w[RuntimeOwned GoRuntimeBacked KDEPolicyOwner UserVisible ReviewOnly CardsPersisted ActionEnablementChanged RequestObjectsCreated PermissionGrantsCreated SettingsPersisted AIProviderCalled AIProviderCallEnabled BackendProcessStarted LaunchEnabled ExecutionStarted HostRootModified NetworkRequired PrivilegedContainerRequired StateRootPathExposed RawExecutableExposed RawCommandExposed BackendDetailsExposed].each do |token|
  assert(runtime_policy_explanation_cards_source.include?(token), "Runtime policy explanation cards must expose #{token}")
end
%w[ApplicationReadinessPreview RepairPlanPreview NewDesktopSafetyPolicyPreview validateNoBackendTerms uniqueSortedStrings runtimePolicyExplanationCard countRuntimePolicyExplanationCards].each do |token|
  assert(runtime_policy_explanation_cards_source.include?(token), "Runtime policy explanation cards must derive or validate evidence through #{token}")
end
runtime_policy_explanation_cards_test_source = read_project_file("internal/runtime/appidentity/runtime_policy_explanation_cards_test.go")
%w[TestRuntimePolicyExplanationCardsPreviewBuildsConsistentDeck TestRuntimePolicyExplanationCardsPreviewDeduplicatesBlockers TestRuntimePolicyExplanationCardsPreviewRejectsBadMode TestRuntimePolicyExplanationCardsPreviewRejectsBadRepairIssue runtime-policy-explanation-cards-preview kde-runtime-policy-explanation-card-deck cards_persisted action_enablement_changed request_objects_created permission_grants_created ai_provider_called backend_process_started host_root_modified raw_command_exposed].each do |token|
  assert(runtime_policy_explanation_cards_test_source.include?(token), "Runtime policy explanation card tests must include #{token}")
end
runtime_policy_explanation_cards_cli_source = read_project_file("cmd/xnix-runtime-go/main.go") + read_project_file("cmd/xnix-runtime-go/runtime_policy_explanation_cards_commands.go")
%w[runtime-policy-explanation-cards-preview runRuntimePolicyExplanationCardsPreview parseRuntimePolicyExplanationCardsSource RuntimePolicyExplanationCardsPreview runtime-root portal-operation snapshot-reason].each do |token|
  assert(runtime_policy_explanation_cards_cli_source.include?(token), "Runtime policy explanation card CLI must include #{token}")
end
runtime_policy_explanation_cards_cli_test_source = read_project_file("cmd/xnix-runtime-go/runtime_policy_explanation_cards_cli_test.go")
%w[TestRuntimePolicyExplanationCardsPreviewCommand TestRuntimePolicyExplanationCardsPreviewCommandRejectsAmbiguousSource TestRuntimePolicyExplanationCardsPreviewCommandRejectsPositionalArguments xnix.runtime.policy_explanation_cards.v1 GetRuntimePolicyExplanationCards GetRuntimePolicyExplanationCardsPreview kde-runtime-policy-explanation-card-deck cards_persisted action_enablement_changed request_objects_created host_root_modified backend_details_exposed].each do |token|
  assert(runtime_policy_explanation_cards_cli_test_source.include?(token), "Runtime policy explanation card CLI tests must include #{token}")
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

go_runtime_live_owner_gate_source = read_project_file("internal/runtime/appidentity/runtime_live_owner_gate.go")
%w[RuntimeLiveOwnerGatePreview runtime-live-owner-gate-preview xnix.runtime.live_owner_gate.v1 GetRuntimeLiveOwnerGate GetRuntimeLiveOwnerGatePreview runtime-service-binding-preview activation-binding long-running-runtime-owner bus-name-acquisition read-only-method-parity production-recipe-trust].each do |token|
  assert(go_runtime_live_owner_gate_source.include?(token), "Go Runtime live owner gate preview must include #{token}")
end
%w[RuntimeOwned GoRuntimeBacked ActivationBindingReady LiveDBusOwnerReady ProductionOwnerEnabled OwnerTransitionReady SmokeAdapterAvailable SmokeAdapterIsProduction KDEMayClaimRuntimeOwnership NetworkRequired HostRootModified PrivilegedContainerRequired BackendDetailsExposed].each do |token|
  assert(go_runtime_live_owner_gate_source.include?(token), "Go Runtime live owner gate preview must expose #{token}")
end
assert(go_runtime_live_owner_gate_source.include?("NewRuntimeServiceBindingPreview"), "Go Runtime live owner gate preview must derive from the service binding preview")
assert(go_runtime_live_owner_gate_source.include?("validateNoBackendTerms"), "Go Runtime live owner gate preview must hide backend terms")

go_runtime_live_owner_gate_cli_source = read_project_file("cmd/xnix-runtime-go/main.go")
assert(go_runtime_live_owner_gate_cli_source.include?("runtime-live-owner-gate-preview"), "Go Runtime CLI must expose Runtime live owner gate preview")

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

go_runtime_owner_smoke_plan_source = read_project_file("internal/runtime/appidentity/runtime_owner_smoke_plan.go")
%w[RuntimeOwnerSmokePlanPreview runtime-owner-smoke-plan-preview xnix.runtime.owner_smoke_plan.v1 GetRuntimeOwnerSmokePlan GetRuntimeOwnerSmokePlanPreview runtime-live-owner-gate-preview validate-activation-files start-packaged-runtime-owner assert-stable-bus-name check-read-only-method-parity reject-write-methods verify-non-production-smoke-adapter-boundary report-kde-safe-summary].each do |token|
  assert(go_runtime_owner_smoke_plan_source.include?(token), "Go Runtime owner smoke plan preview must include #{token}")
end
%w[RuntimeOwned GoRuntimeBacked ActivationBindingReady LiveDBusOwnerReady ProductionOwnerEnabled OwnerTransitionReady SmokeState SmokeEnvironment NetworkRequired HostRootModified PrivilegedContainerRequired SystemServiceStarted ProductionBusClaimed BackendDetailsExposed].each do |token|
  assert(go_runtime_owner_smoke_plan_source.include?(token), "Go Runtime owner smoke plan preview must expose #{token}")
end
assert(go_runtime_owner_smoke_plan_source.include?("NewRuntimeLiveOwnerGatePreview"), "Go Runtime owner smoke plan preview must derive from the live owner gate preview")
assert(go_runtime_owner_smoke_plan_source.include?("validateNoBackendTerms"), "Go Runtime owner smoke plan preview must hide backend terms")

go_runtime_owner_smoke_plan_cli_source = read_project_file("cmd/xnix-runtime-go/main.go")
assert(go_runtime_owner_smoke_plan_cli_source.include?("runtime-owner-smoke-plan-preview"), "Go Runtime CLI must expose Runtime owner smoke plan preview")

runtime_method_parity_source = read_project_file("lib/xnix/compatibility/runtime_method_parity_manifest.rb")
assert(runtime_method_parity_source.include?("xnix-runtime-method-parity-manifest"), "Runtime method parity manifest must expose a CLI command")
%w[runtime-method-parity-manifest READ_ONLY_METHODS dbus-contract runtime-dispatch dbus-client smoke-adapter session-smoke GetRuntimeMethodParityManifest GetRuntimeWriteGate GetCompatibilityModeSwitchPlan GetCompatibilityPermissionReviewPlan GetCompatibilityReviewFlowPlan GetKDECenterPage GetKDECenterPageSections GetKDECenterPageSectionDetail].each do |token|
  assert(runtime_method_parity_source.include?(token), "Runtime method parity manifest must include #{token}")
end
assert(runtime_method_parity_source.include?("\"write_methods_supported\" => false"), "Runtime method parity manifest must not claim write method support")
assert(runtime_method_parity_source.include?("\"write_method_dispatch_enabled\" => false"), "Runtime method parity manifest must not enable write method dispatch")
assert(runtime_method_parity_source.include?("\"host_root_modified\" => false"), "Runtime method parity manifest must not mutate the host root")
assert(runtime_method_parity_source.include?("\"backend_details_exposed\" => false"), "Runtime method parity manifest must hide backend details")

go_runtime_method_parity_source = read_project_file("internal/runtime/appidentity/runtime_method_parity_manifest.go")
%w[RuntimeMethodParityManifestPreview runtime-method-parity-manifest-preview xnix.runtime.method_parity_manifest.v1 GetRuntimeMethodParityManifest GetRuntimeMethodParityManifestPreview dbus-contract runtime-dispatch dbus-client smoke-adapter session-smoke ListApplications GetRuntimeOwnerProcess GetRuntimeOwnerSmokePlan GetRuntimeMethodParityManifest GetRuntimeOwnerRouteManifest GetRuntimeOwnerRecipeTrust GetRuntimeOwnerReadiness GetRuntimeWriteGate GetKDECenterPageSectionDetail].each do |token|
  assert(go_runtime_method_parity_source.include?(token), "Go Runtime method parity manifest preview must include #{token}")
end
%w[ReadOnlyMethodParityReady WriteMethodsSupported WriteMethodDispatchEnabled RuntimeOwned GoRuntimeBacked NetworkRequired HostRootModified PrivilegedContainerRequired BackendDetailsExposed].each do |token|
  assert(go_runtime_method_parity_source.include?(token), "Go Runtime method parity manifest preview must expose #{token}")
end
assert(go_runtime_method_parity_source.include?("InstallRecipe"), "Go Runtime method parity manifest preview must list gated write methods")
assert(go_runtime_method_parity_source.include?("validateNoBackendTerms"), "Go Runtime method parity manifest preview must hide backend terms")

go_runtime_method_parity_cli_source = read_project_file("cmd/xnix-runtime-go/main.go")
assert(go_runtime_method_parity_cli_source.include?("runtime-method-parity-manifest-preview"), "Go Runtime CLI must expose Runtime method parity manifest preview")

go_runtime_owner_recipe_trust_source = read_project_file("internal/runtime/appidentity/runtime_owner_recipe_trust.go")
%w[RuntimeOwnerRecipeTrustPreview runtime-owner-recipe-trust-preview xnix.runtime.owner_recipe_trust.v1 GetRuntimeOwnerRecipeTrust GetRuntimeOwnerRecipeTrustPreview go-recipe-store-verifier registry-digests+recipe-signature-status registry-present recipe-digests signed-recipe-validation development-registry unsigned-recipes runtime/recipes/registry.json].each do |token|
  assert(go_runtime_owner_recipe_trust_source.include?(token), "Go Runtime owner recipe trust preview must include #{token}")
end
%w[DigestVerified SignedRecipeValidation DevelopmentRegistry UnsignedRecipesPresent ProductionRecipeTrustReady RuntimeOwned GoRuntimeBacked KDEPolicyOwner NetworkRequired HostRootModified PrivilegedContainerRequired BackendDetailsExposed].each do |token|
  assert(go_runtime_owner_recipe_trust_source.include?(token), "Go Runtime owner recipe trust preview must expose #{token}")
end
assert(go_runtime_owner_recipe_trust_source.include?("NewLocalStore"), "Go Runtime owner recipe trust preview must consume the recipe Store boundary")
assert(go_runtime_owner_recipe_trust_source.include?("DigestVerifier"), "Go Runtime owner recipe trust preview must consume the recipe Verifier boundary")
assert(go_runtime_owner_recipe_trust_source.include?("validateNoBackendTerms"), "Go Runtime owner recipe trust preview must hide backend terms")

go_runtime_owner_recipe_trust_cli_source = read_project_file("cmd/xnix-runtime-go/main.go")
assert(go_runtime_owner_recipe_trust_cli_source.include?("runtime-owner-recipe-trust-preview"), "Go Runtime CLI must expose Runtime owner recipe trust preview")

go_runtime_owner_process_source = read_project_file("internal/runtime/appidentity/runtime_owner_process.go")
%w[RuntimeOwnerProcessPreview runtime-owner-process-preview xnix.runtime.owner_process.v1 GetRuntimeOwnerProcess GetRuntimeOwnerProcessPreview runtime-service-binding-preview libexec-wrapper go-owner-target service-activation packaged-entrypoint production-owner-loop host-safety-boundary ruby-wrapper].each do |token|
  assert(go_runtime_owner_process_source.include?(token), "Go Runtime owner process preview must include #{token}")
end
%w[ServiceActivationReady PackagedEntrypointReady GoOwnerProcessReady ProductionOwnerProcessReady RuntimeOwned GoRuntimeBacked KDEPolicyOwner KDEMayClaimRuntimeOwnership SystemServiceStarted ProductionBusClaimed NetworkRequired HostRootModified PrivilegedContainerRequired BackendDetailsExposed].each do |token|
  assert(go_runtime_owner_process_source.include?(token), "Go Runtime owner process preview must expose #{token}")
end
assert(go_runtime_owner_process_source.include?("NewRuntimeServiceBindingPreview"), "Go Runtime owner process preview must derive from service binding")
assert(go_runtime_owner_process_source.include?("runtimeOwnerProcessGoCandidateReady"), "Go Runtime owner process preview must detect the Go owner candidate")
assert(go_runtime_owner_process_source.include?("validateNoBackendTerms"), "Go Runtime owner process preview must hide backend terms")

go_runtime_owner_candidate_source = read_project_file("internal/runtime/owner/candidate.go")
%w[Candidate runtime-owner-candidate xnix.runtime.owner_candidate.v1 go-runtime-owner-candidate runtime-service-binding-preview+runtime-owner-route-manifest-preview+runtime-write-gate-preview ModeSmokeOwner DisabledWriteResponse org.xnix.Compatibility1.Error.WriteMethodDisabled ReadOnlyServeReady SmokeOwnerMode ProductionOwnerMode SessionBusClaimed ProductionBusClaimed WriteMethodsEnabled HostRootModified PrivilegedContainerRequired BackendDetailsExposed validateNoBackendTerms].each do |token|
  assert(go_runtime_owner_candidate_source.include?(token), "Go Runtime owner candidate must include #{token}")
end
%w[RuntimeOwned GoRuntimeBacked KDEPolicyOwner KDEMayClaimRuntimeOwnership ServiceActivationReady ReadOnlyRouteTableReady EventLoopStarted SystemServiceStarted NetworkRequired].each do |token|
  assert(go_runtime_owner_candidate_source.include?(token), "Go Runtime owner candidate must expose #{token}")
end

go_runtime_owner_candidate_cli_source = read_project_file("cmd/xnix-runtime-owner/main.go")
%w[xnix-runtime-owner mode deny-write dispatch-read lifecycle-log smoke-batch session-bus-smoke route-checkpoint materialization-fanout-owner-smoke-coverage restricted-smoke-fanout-owner-smoke-coverage redacted-adapter-profile-owner-smoke-coverage NewCandidate DisabledWriteResponse DispatchRead NewLifecycleEvents NewSmokeBatchRecords NewSessionBusSmokeTranscript NewRouteCheckpoint NewKDETestLaunchMaterializationFanOutOwnerSmokeCoveragePreview NewRestrictedOwnerSmokeReceiptFanOutOwnerSmokeCoveragePreview NewBackendAdapterRedactedProfileOwnerSmokeCoveragePreview smoke-owner].each do |token|
  assert(go_runtime_owner_candidate_cli_source.include?(token), "Go Runtime owner candidate CLI must include #{token}")
end

go_runtime_owner_dispatch_source = read_project_file("internal/runtime/owner/dispatch.go")
%w[ReadDispatch runtime-owner-read-dispatch xnix.runtime.owner_read_dispatch.v1 go-owner-read-dispatch go-runtime-owner-candidate+in-process-read-dispatch DispatchRead SupportedReadDispatchMethods ListApplications GetApplication GetDiagnostics GetRunPlan GetDesktopActivationManifest GetKDEApplicationSurfacePlan GetCompatibilityInstallPlan GetRuntimeServiceBinding GetRuntimeLiveOwnerGate GetRuntimeOwnerProcess GetRuntimeOwnerSmokePlan GetRuntimeMethodParityManifest GetRuntimeOwnerRouteManifest GetRuntimeOwnerRecipeTrust GetRuntimeOwnerReadiness GetWindowsCompatibilityWorkstreamsPreview windows-compatibility-workstreams-preview GetKDENotificationDigestPreview kde-notification-digest-preview ownerNotificationDigestEvents GetSignedRecipeVerificationPreview signed-recipe-verifier-preview NewRegistrySignedRecipeVerificationPreview GetRestrictedProductSmokePacketPreview restricted-product-smoke-packet-preview PrepareRestrictedProductSmokePacket GetBackendAdapterProfileAudit backend-adapter-redacted-profile-audit-preview NewBackendAdapterRedactedProfileAuditPreview GetKDETestLaunchMaterializationReceiptLookupPreview kde-test-launch-materialization-receipt-lookup-preview ResolveKDETestLaunchMaterializationReceipt GetKDETestLaunchMaterializationFanOut kde-test-launch-materialization-fanout-owner-route-preview NewKDETestLaunchMaterializationFanOutOwnerRoutePreview go-owner-local-preview GetRuntimeWriteGate GetKDECenterPageSectionDetail].each do |token|
  assert(go_runtime_owner_dispatch_source.include?(token), "Go Runtime owner read dispatch must include #{token}")
end
%w[ReadOnlyDispatch WriteMethod WriteMethodsEnabled RuntimeOwned GoRuntimeBacked KDEPolicyOwner KDEMayClaimRuntimeOwnership EventLoopStarted SessionBusClaimed ProductionBusClaimed SystemServiceStarted NetworkRequired HostRootModified PrivilegedContainerRequired BackendDetailsExposed].each do |token|
  assert(go_runtime_owner_dispatch_source.include?(token), "Go Runtime owner read dispatch must expose #{token}")
end
assert(go_runtime_owner_dispatch_source.include?("NewRuntimeOwnerReadinessPreview"), "Go Runtime owner read dispatch must call owner readiness preview")
assert(go_runtime_owner_dispatch_source.include?("LoadRecipeFromRegistry"), "Go Runtime owner read dispatch must load registry recipes directly")
assert(go_runtime_owner_dispatch_source.include?("NewApplicationsPreview"), "Go Runtime owner read dispatch must render application catalog payloads")
assert(go_runtime_owner_dispatch_source.include?("NewWindowsCompatibilityWorkstreamsPreview"), "Go Runtime owner read dispatch must render Windows compatibility workstream payloads")
assert(go_runtime_owner_dispatch_source.include?("NewRuntimeWriteGatePreview"), "Go Runtime owner read dispatch must call write gate preview")
assert(go_runtime_owner_dispatch_source.include?("validateNoBackendTerms"), "Go Runtime owner read dispatch must hide backend terms")

go_runtime_owner_route_checkpoint_source = read_project_file("internal/runtime/owner/route_checkpoint.go")
%w[RouteCheckpoint RouteCheckpointCheck xnix.runtime.owner_route_checkpoint.v1 runtime-owner-route-checkpoint go-owner-read-route-band-checkpoint NewRouteCheckpoint FormalReadRouteCount GoFormalReadRouteCount OwnerReadMethodCount OwnerLocalReadMethodCount SmokeReadRecordCount SmokeWriteDenialCount MethodParityReady FormalRouteCoverageReady OwnerLocalRouteCoverageReady SmokeBatchCoverageReady DeterministicWriteDenialsReady RouteBandReady ProductionBusClaimed WriteMethodsEnabled HostRootModified validateNoBackendTerms].each do |token|
  assert(go_runtime_owner_route_checkpoint_source.include?(token), "Go Runtime owner route checkpoint must include #{token}")
end
go_runtime_owner_route_checkpoint_test_source = read_project_file("internal/runtime/owner/route_checkpoint_test.go") +
                                                read_project_file("cmd/xnix-runtime-owner/main_test.go")
%w[TestRouteCheckpointClosesReadRouteBandWithWritesDisabled TestRuntimeOwnerCommandRendersRouteCheckpoint formal_read_route_count owner_read_method_count owner_local_read_method_count smoke_read_record_count smoke_write_denial_count route_band_ready write_methods_enabled production_bus_claimed host_root_modified].each do |token|
  assert(go_runtime_owner_route_checkpoint_test_source.include?(token), "Go Runtime owner route checkpoint tests must include #{token}")
end

go_runtime_kde_offline_identity_checkpoint_source = read_project_file("internal/runtime/owner/kde_offline_identity_checkpoint.go")
%w[KDEOfflineIdentityCheckpoint KDEOfflineIdentityCheckpointCheck xnix.runtime.kde_offline_identity_checkpoint.v1 kde-offline-identity-checkpoint offline-kde-application-identity-band-checkpoint NewKDEOfflineIdentityCheckpoint recipe-trust surface-coverage identity-parity owner-route owner-service side-effects-disabled OfflineIdentityReady ProductionSignatureReady FormalReadRouteCount OwnerReadMethodCount OwnerLocalReadMethodCount SmokeReadRecordCount SmokeWriteDenialCount].each do |token|
  assert(go_runtime_kde_offline_identity_checkpoint_source.include?(token), "Go Runtime offline KDE identity checkpoint must include #{token}")
end
%w[DesktopFilesWritten MIMEDefaultsWritten KRunnerIndexPersisted TaskManagerEntryActive KWinRuleApplied LiveTrayBridgeEnabled NotificationSent SettingsPersisted CompatibilityCenterPersisted LaunchEnabled ExecutionStarted BackendProcessStarted ProductionBusClaimed WriteMethodsEnabled NetworkRequired HostRootModified PrivilegedContainerRequired RawCommandExposed BackendDetailsExposed].each do |token|
  assert(go_runtime_kde_offline_identity_checkpoint_source.include?(token), "Go Runtime offline KDE identity checkpoint must expose disabled gate #{token}")
end

go_runtime_kde_offline_identity_checkpoint_test_source = read_project_file("internal/runtime/owner/kde_offline_identity_checkpoint_test.go") +
                                                       read_project_file("cmd/xnix-runtime-owner/main_test.go") +
                                                       read_project_file("cmd/xnix-runtime-owner/main.go")
%w[TestKDEOfflineIdentityCheckpointClosesNineSurfaceBand TestRuntimeOwnerCommandRendersKDEOfflineIdentityCheckpoint kde-identity-checkpoint xnix.runtime.kde_offline_identity_checkpoint.v1 offline_identity_ready production_signature_ready production_bus_claimed write_methods_enabled launch_enabled backend_process_started host_root_modified].each do |token|
  assert(go_runtime_kde_offline_identity_checkpoint_test_source.include?(token), "Go Runtime offline KDE identity checkpoint tests must include #{token}")
end

go_runtime_owner_lifecycle_source = read_project_file("internal/runtime/owner/lifecycle.go")
%w[LifecycleEvent runtime-owner-lifecycle-event xnix.runtime.owner_lifecycle_event.v1 NewLifecycleEvents startup route-table readiness shutdown preview-complete RouteTableVersion ShutdownReason DesktopSafeSummary].each do |token|
  assert(go_runtime_owner_lifecycle_source.include?(token), "Go Runtime owner lifecycle events must include #{token}")
end
%w[RuntimeOwned GoRuntimeBacked KDEPolicyOwner KDEMayClaimRuntimeOwnership EventLoopStarted SessionBusClaimed ProductionBusClaimed SystemServiceStarted NetworkRequired HostRootModified PrivilegedContainerRequired BackendDetailsExposed WriteMethodsEnabled].each do |token|
  assert(go_runtime_owner_lifecycle_source.include?(token), "Go Runtime owner lifecycle events must expose #{token}")
end

go_runtime_owner_smoke_batch_source = read_project_file("internal/runtime/owner/smoke_batch.go")
%w[SmokeBatchRecord runtime-owner-smoke-batch-record xnix.runtime.owner_smoke_batch.v1 restricted-session-owner-call-batch NewSmokeBatchRecords read-dispatch write-denial NewService service.Call runtime-owner-service-call SupportedReadDispatchMethods].each do |token|
  assert(go_runtime_owner_smoke_batch_source.include?(token), "Go Runtime owner smoke batch must include #{token}")
end
%w[ReadOnlyDispatch WriteMethod RouteReady DispatchReady ReadDispatchMethodCount WriteMethodCount RuntimeOwned GoRuntimeBacked KDEPolicyOwner KDEMayClaimRuntimeOwnership EventLoopStarted SessionBusClaimed ProductionBusClaimed SystemServiceStarted NetworkRequired HostRootModified PrivilegedContainerRequired BackendDetailsExposed].each do |token|
  assert(go_runtime_owner_smoke_batch_source.include?(token), "Go Runtime owner smoke batch must expose #{token}")
end
assert(go_runtime_owner_smoke_batch_source.include?("validateNoBackendTerms"), "Go Runtime owner smoke batch must hide backend terms")

go_runtime_materialization_owner_smoke_coverage_source = read_project_file("internal/runtime/owner/kde_test_launch_materialization_fanout_owner_smoke_coverage.go")
%w[KDETestLaunchMaterializationFanOutOwnerSmokeCoveragePreview KDETestLaunchMaterializationOwnerSmokeCheck NewKDETestLaunchMaterializationFanOutOwnerSmokeCoveragePreview NewKDETestLaunchMaterializationFanOutOwnerSmokeCoveragePreviewFromRecords xnix.runtime.kde_test_launch_materialization_fanout_owner_smoke_coverage.v1 kde-test-launch-materialization-fanout-owner-smoke-coverage-preview materialization-fanout-owner-smoke-coverage GetKDETestLaunchMaterializationFanOut GetKDETestLaunchMaterializationFanOutOwnerSmokeCoveragePreview kde-test-launch-materialization-fanout-owner-route-preview owner-smoke-batch+runtime-owner-service-call+kde-test-launch-materialization-fanout-owner-route smoke-record-present service-call-read-dispatch owner-route-payload-present opaque-receipt-preserved missing-receipt-fail-closed desktop-side-effects-disabled production-dbus-blocked unsafe-gates-closed SmokeCoverageReady].each do |token|
  assert(go_runtime_materialization_owner_smoke_coverage_source.include?(token), "Go Runtime materialization fan-out owner smoke coverage must include #{token}")
end
%w[StateRootPathExposed StateRootWritesEnabled RuntimeWritesEnabled FanOutWritesEnabled SystemServiceStarted SessionBusClaimed ProductionBusClaimed WriteMethodsEnabled BackendLaunchEnabled BackendProcessStarted TaskManagerEntryActive LiveTrayBridgeEnabled NotificationSent CompatibilityCenterActionsEnabled NetworkRequired HostRootModified PrivilegedContainerRequired RawCommandExposed RawExecutableExposed BackendDetailsExposed validateNoBackendTerms].each do |token|
  assert(go_runtime_materialization_owner_smoke_coverage_source.include?(token), "Go Runtime materialization fan-out owner smoke coverage must expose safety gate #{token}")
end
go_runtime_materialization_owner_smoke_coverage_test_source = read_project_file("internal/runtime/owner/kde_test_launch_materialization_fanout_owner_smoke_coverage_test.go") +
                                                             read_project_file("cmd/xnix-runtime-owner/main_test.go") +
                                                             read_project_file("cmd/xnix-runtime-owner/main.go")
%w[TestKDETestLaunchMaterializationFanOutOwnerSmokeCoveragePreviewCoversOwnerRoute TestKDETestLaunchMaterializationFanOutOwnerSmokeCoveragePreviewFailsClosedWithoutRecord TestRuntimeOwnerCommandRendersMaterializationFanOutOwnerSmokeCoverage materialization-fanout-owner-smoke-coverage kde-test-launch-materialization-fanout-owner-smoke-coverage-preview smoke_coverage_ready service_call_dispatch_ready dispatch_go_command].each do |token|
  assert(go_runtime_materialization_owner_smoke_coverage_test_source.include?(token), "Go Runtime materialization fan-out owner smoke coverage tests must include #{token}")
end

go_runtime_restricted_smoke_fanout_owner_smoke_coverage_source = read_project_file("internal/runtime/owner/restricted_smoke_receipt_fanout_owner_smoke_coverage.go")
%w[RestrictedOwnerSmokeReceiptFanOutOwnerSmokeCoveragePreview RestrictedOwnerSmokeFanOutCoverageCheck NewRestrictedOwnerSmokeReceiptFanOutOwnerSmokeCoveragePreview NewRestrictedOwnerSmokeReceiptFanOutOwnerSmokeCoveragePreviewFromRecords xnix.runtime.restricted_owner_smoke_receipt_fanout_owner_smoke_coverage.v1 restricted-owner-smoke-receipt-fanout-owner-smoke-coverage-preview restricted-owner-smoke-fanout-owner-smoke-coverage GetRestrictedOwnerSmokeReceiptFanOut GetRestrictedOwnerSmokeReceiptFanOutOwnerSmokeCoveragePreview restricted-owner-smoke-receipt-fanout-owner-route-preview owner-smoke-batch+runtime-owner-service-call+restricted-owner-smoke-receipt-fanout-owner-route smoke-record-present service-call-read-dispatch owner-route-payload-present opaque-receipt-preserved missing-receipt-fail-closed support-side-effects-disabled production-dbus-blocked unsafe-gates-closed SmokeCoverageReady].each do |token|
  assert(go_runtime_restricted_smoke_fanout_owner_smoke_coverage_source.include?(token), "Go Runtime restricted owner smoke fan-out owner smoke coverage must include #{token}")
end
%w[StateRootPathExposed StateRootWritesEnabled RuntimeWritesEnabled FanOutWritesEnabled SystemServiceStarted SessionBusClaimed ProductionBusClaimed WriteMethodsEnabled ProductionActivationReady ProductionOwnerEnabled SupportBundleExported SupportCaseCreated NotificationSent BackendLaunchEnabled BackendProcessStarted NetworkRequired HostRootModified PrivilegedContainerRequired BackendDetailsExposed validateNoBackendTerms].each do |token|
  assert(go_runtime_restricted_smoke_fanout_owner_smoke_coverage_source.include?(token), "Go Runtime restricted owner smoke fan-out owner smoke coverage must expose safety gate #{token}")
end
go_runtime_restricted_smoke_fanout_owner_smoke_coverage_test_source = read_project_file("internal/runtime/owner/restricted_smoke_receipt_fanout_owner_smoke_coverage_test.go") +
                                                                     read_project_file("cmd/xnix-runtime-owner/main_test.go") +
                                                                     read_project_file("cmd/xnix-runtime-owner/main.go")
%w[TestRestrictedOwnerSmokeReceiptFanOutOwnerSmokeCoveragePreviewCoversOwnerRoute TestRestrictedOwnerSmokeReceiptFanOutOwnerSmokeCoveragePreviewFailsClosedWithoutRecord TestRuntimeOwnerCommandRendersRestrictedSmokeFanOutOwnerSmokeCoverage restricted-smoke-fanout-owner-smoke-coverage restricted-owner-smoke-receipt-fanout-owner-smoke-coverage-preview smoke_coverage_ready service_call_dispatch_ready dispatch_go_command].each do |token|
  assert(go_runtime_restricted_smoke_fanout_owner_smoke_coverage_test_source.include?(token), "Go Runtime restricted owner smoke fan-out owner smoke coverage tests must include #{token}")
end

go_runtime_redacted_adapter_profile_owner_smoke_coverage_source = read_project_file("internal/runtime/owner/backend_adapter_redacted_profile_owner_smoke_coverage.go")
%w[BackendAdapterRedactedProfileOwnerSmokeCoveragePreview BackendAdapterRedactedProfileSmokeCoverageCheck NewBackendAdapterRedactedProfileOwnerSmokeCoveragePreview NewBackendAdapterRedactedProfileOwnerSmokeCoveragePreviewFromRecords xnix.runtime.backend_adapter_redacted_profile_owner_smoke_coverage.v1 backend-adapter-redacted-profile-owner-smoke-coverage-preview redacted-adapter-profile-owner-smoke-coverage GetBackendAdapterProfileAudit GetBackendAdapterRedactedProfileOwnerSmokeCoveragePreview backend-adapter-redacted-profile-audit-preview owner-smoke-batch+runtime-owner-service-call+backend-adapter-redacted-profile-audit-owner-route smoke-record-present service-call-read-dispatch owner-route-payload-present redacted-profiles-preserved full-contract-fixture-local production-dbus-blocked caller-paths-hidden unsafe-gates-closed SmokeCoverageReady].each do |token|
  assert(go_runtime_redacted_adapter_profile_owner_smoke_coverage_source.include?(token), "Go Runtime redacted adapter profile owner smoke coverage must include #{token}")
end
%w[StateRootPathExposed StateRootWritesEnabled RuntimeWritesEnabled AdapterInvocationEnabled BackendInstallEnabled BackendDownloadEnabled BackendLaunchEnabled BackendProcessStarted CommandMaterialized ExecutablePathResolved NetworkRequired HostRootModified PrivilegedContainerRequired RawCommandExposed RawExecutableExposed BackendDetailsExposed validateNoBackendTerms].each do |token|
  assert(go_runtime_redacted_adapter_profile_owner_smoke_coverage_source.include?(token), "Go Runtime redacted adapter profile owner smoke coverage must expose safety gate #{token}")
end
go_runtime_redacted_adapter_profile_owner_smoke_coverage_test_source = read_project_file("internal/runtime/owner/backend_adapter_redacted_profile_owner_smoke_coverage_test.go") +
                                                                      read_project_file("cmd/xnix-runtime-owner/main_test.go") +
                                                                      read_project_file("cmd/xnix-runtime-owner/main.go")
%w[TestBackendAdapterRedactedProfileOwnerSmokeCoveragePreviewCoversOwnerRoute TestBackendAdapterRedactedProfileOwnerSmokeCoveragePreviewFailsClosedWithoutRecord TestRuntimeOwnerCommandRendersRedactedAdapterProfileOwnerSmokeCoverage redacted-adapter-profile-owner-smoke-coverage backend-adapter-redacted-profile-owner-smoke-coverage-preview smoke_coverage_ready service_call_dispatch_ready dispatch_go_command internal_adapter_ids_redacted adapter_invocation_enabled].each do |token|
  assert(go_runtime_redacted_adapter_profile_owner_smoke_coverage_test_source.include?(token), "Go Runtime redacted adapter profile owner smoke coverage tests must include #{token}")
end

go_runtime_production_dbus_gate_review_source = read_project_file("internal/runtime/owner/production_dbus_gate_review.go")
%w[ProductionDBusGateReviewPreview ProductionDBusGateReviewRoute ProductionDBusGateReviewCheck NewProductionDBusGateReviewPreview xnix.runtime.production_dbus_gate_review.v1 production-dbus-gate-review-preview owner-local-smoke-covered-production-dbus-gate-review materialization-fanout-owner-route restricted-smoke-fanout-owner-route redacted-adapter-profile-owner-route production-dbus-gate-review-ready production-dbus-gate-review-blocked human-authorization production-service-activation-preflight runtime-write-gate-review route-by-route-production-dbus-method-review desktop-side-effect-review rollback-and-diagnostics-review tracked-routes-present smoke-coverage-present production-dbus-disabled route-production-exposure-disabled writes-and-launch-disabled desktop-side-effects-disabled host-boundary-closed human-authorization-required].each do |token|
  assert(go_runtime_production_dbus_gate_review_source.include?(token), "Go Runtime production D-Bus gate review must include #{token}")
end
%w[ProductionReadiness HumanAuthorizationRequired SystemServiceStarted SessionBusClaimed ProductionBusClaimed ProductionOwnerEnabled ProductionActivationReady WriteMethodsEnabled RuntimeWritesEnabled AdapterInvocationEnabled BackendLaunchEnabled BackendProcessStarted SupportBundleExported SupportCaseCreated NotificationSent NetworkRequired HostRootModified PrivilegedContainerRequired StateRootPathExposed RawCommandExposed RawExecutableExposed BackendDetailsExposed validateNoBackendTerms].each do |token|
  assert(go_runtime_production_dbus_gate_review_source.include?(token), "Go Runtime production D-Bus gate review must expose safety gate #{token}")
end
%w[HumanAuthorizationPreflightReady HumanAuthorizationGranted AuthorizationReceiptAccepted human-authorization-preflight-present production-dbus-human-authorization-preflight-preview].each do |token|
  assert(go_runtime_production_dbus_gate_review_source.include?(token), "Go Runtime production D-Bus gate review must consume human authorization preflight #{token}")
end
go_runtime_production_dbus_gate_review_test_source = read_project_file("internal/runtime/owner/production_dbus_gate_review_test.go") +
                                                    read_project_file("cmd/xnix-runtime-go/production_dbus_gate_cli_test.go") +
                                                    read_project_file("cmd/xnix-runtime-go/production_dbus_gate_commands.go") +
                                                    read_project_file("cmd/xnix-runtime-go/main.go")
%w[TestProductionDBusGateReviewPreviewInventoriesSmokeCoveredRoutes TestProductionDBusGateReviewPreviewFailsClosedWithoutSources TestProductionDBusGateReviewPreviewCommand production-dbus-gate-review-preview production-dbus-human-authorization-preflight-preview production-dbus-gate-review-ready smoke_covered_route_count human_authorization_required human_authorization_preflight_ready human_authorization_granted authorization_receipt_accepted production_readiness production_owner_enabled production_activation_ready].each do |token|
  assert(go_runtime_production_dbus_gate_review_test_source.include?(token), "Go Runtime production D-Bus gate review tests and CLI must include #{token}")
end

go_runtime_production_dbus_human_authorization_preflight_source = read_project_file("internal/runtime/owner/production_dbus_human_authorization_preflight.go")
%w[ProductionDBusHumanAuthorizationPreflightPreview NewProductionDBusHumanAuthorizationPreflightPreview xnix.runtime.production_dbus_human_authorization_preflight.v1 production-dbus-human-authorization-preflight-preview read-only-production-dbus-human-authorization-preflight xnix.runtime.production_dbus_human_authorization_receipt.v1 production-dbus-human-authorization-receipt-id gate-review-present route-inventory-present human-authorization-required receipt-shape-declared authorization-not-granted production-ownership-disabled unsafe-gates-closed PreflightReady AuthorizationReceiptRequired AuthorizationReceiptPresent AuthorizationGrantReady AuthorizationAccepted ProductionDBusGateReviewRequired].each do |token|
  assert(go_runtime_production_dbus_human_authorization_preflight_source.include?(token), "Go Runtime production D-Bus human authorization preflight must include #{token}")
end
%w[ProductionReadiness SystemServiceStarted SessionBusClaimed ProductionBusClaimed ProductionOwnerEnabled ProductionActivationReady WriteMethodsEnabled RuntimeWritesEnabled AdapterInvocationEnabled BackendLaunchEnabled BackendProcessStarted SupportBundleExported SupportCaseCreated NotificationSent NetworkRequired HostRootModified PrivilegedContainerRequired StateRootPathExposed RawCommandExposed RawExecutableExposed BackendDetailsExposed validateNoBackendTerms].each do |token|
  assert(go_runtime_production_dbus_human_authorization_preflight_source.include?(token), "Go Runtime production D-Bus human authorization preflight must expose safety gate #{token}")
end
go_runtime_production_dbus_human_authorization_preflight_test_source = read_project_file("internal/runtime/owner/production_dbus_human_authorization_preflight_test.go") +
                                                                      read_project_file("cmd/xnix-runtime-go/production_dbus_gate_cli_test.go") +
                                                                      read_project_file("cmd/xnix-runtime-go/production_dbus_gate_commands.go") +
                                                                      read_project_file("cmd/xnix-runtime-go/main.go")
%w[TestProductionDBusHumanAuthorizationPreflightPreviewDefinesReceiptShape TestProductionDBusHumanAuthorizationPreflightPreviewFailsClosedWithoutGateReview TestProductionDBusHumanAuthorizationPreflightPreviewCommand production-dbus-human-authorization-preflight-preview authorization_receipt_required authorization_receipt_present authorization_grant_ready authorization_accepted preflight_ready production_readiness].each do |token|
  assert(go_runtime_production_dbus_human_authorization_preflight_test_source.include?(token), "Go Runtime production D-Bus human authorization preflight tests and CLI must include #{token}")
end

go_runtime_production_human_authorization_receipt_consolidation_source = read_project_file("internal/runtime/owner/production_human_authorization_receipt_consolidation.go")
%w[ProductionHumanAuthorizationReceiptConsolidationPreview ProductionHumanAuthorizationReceiptGate ProductionHumanAuthorizationReceiptConsolidationCheck NewProductionHumanAuthorizationReceiptConsolidationPreview xnix.runtime.production_human_authorization_receipt_consolidation.v1 production-human-authorization-receipt-consolidation-preview owner-managed-opaque-human-authorization-receipt-boundary production-human-authorization-receipt-consolidation-ready-authorization-disabled production-human-authorization-receipt-consolidation-blocked production-dbus-human-authorization-preflight-preview production-dbus-gate-review-preview production-dbus-method-review-preview runtime-service-activation-preflight-preview runtime-write-gate-preview production-rollback-diagnostics-review-preview production-desktop-side-effect-review-preview preflight-shape-consumed production-gate-consumed method-review-consumed service-and-write-gates-consumed rollback-and-desktop-reviews-consumed opaque-receipt-boundary authorization-not-granted unsafe-gates-closed human-authorization-preflight production-dbus-gate-review production-dbus-method-review runtime-service-activation-preflight runtime-write-gate rollback-diagnostics-review desktop-side-effect-review].each do |token|
  assert(go_runtime_production_human_authorization_receipt_consolidation_source.include?(token), "Go Runtime production human authorization receipt consolidation must include #{token}")
end
%w[ReceiptBoundaryConsolidated OwnerManagedOpaqueReceiptLookupReady CallerStateRootRequired ExplicitOperatorActionRequired AuthorizationGrantReady AuthorizationAccepted ProductionReadiness ProductionOwnershipReady GateCount RequiredGateCount ConsumedGateCount MissingGateCount AuthorizationAcceptedGateCount ProductionReadyGateCount SystemServiceStarted SessionBusClaimed ProductionBusClaimed ProductionOwnerEnabled ProductionActivationReady WriteMethodsEnabled RuntimeWritesEnabled ReceiptWriterEnabled ReceiptPersistenceEnabled ReceiptLookupWritesEnabled DesktopFilesWritten MIMEAppsWritten ShellConfigurationWritten SettingsPersisted NotificationSent NotificationDeliveryEnabled PortalRequestCreated RequestObjectsCreated AdapterInvocationEnabled BackendLaunchEnabled BackendProcessStarted SupportBundleExported SupportCaseCreated SnapshotRestoreExecuted StateCleanupExecuted FileContentRead FilePathsExposed NetworkRequired HostRootModified PrivilegedContainerRequired StateRootPathExposed RawCommandExposed RawExecutableExposed BackendDetailsExposed validateNoBackendTerms].each do |token|
  assert(go_runtime_production_human_authorization_receipt_consolidation_source.include?(token), "Go Runtime production human authorization receipt consolidation must expose safety gate #{token}")
end
%w[production_dbus_human_authorization_preflight.go production_dbus_gate_review.go production_dbus_method_review.go runtime_service_activation_preflight.go runtime_write_gate.go production_rollback_diagnostics_review.go production_desktop_side_effect_review.go].each do |token|
  assert(go_runtime_production_human_authorization_receipt_consolidation_source.include?(token), "Go Runtime production human authorization receipt consolidation must consume #{token}")
end
go_runtime_production_human_authorization_receipt_consolidation_test_source = read_project_file("internal/runtime/owner/production_human_authorization_receipt_consolidation_test.go") +
                                                                              read_project_file("cmd/xnix-runtime-go/production_dbus_gate_cli_test.go") +
                                                                              read_project_file("cmd/xnix-runtime-go/production_dbus_gate_commands.go") +
                                                                              read_project_file("cmd/xnix-runtime-go/main.go")
%w[TestProductionHumanAuthorizationReceiptConsolidationPreviewConsumesProductionGates TestProductionHumanAuthorizationReceiptConsolidationPreviewFailsClosedWithoutSources TestProductionHumanAuthorizationReceiptConsolidationPreviewCommand production-human-authorization-receipt-consolidation-preview receipt_boundary_consolidated owner_managed_opaque_receipt_lookup_ready caller_state_root_required receipt_writer_enabled consumed_gate_count missing_gate_count authorization_accepted_gate_count production_ready_gate_count].each do |token|
  assert(go_runtime_production_human_authorization_receipt_consolidation_test_source.include?(token), "Go Runtime production human authorization receipt consolidation tests and CLI must include #{token}")
end

go_runtime_production_authorization_consumption_audit_source = read_project_file("internal/runtime/owner/production_authorization_consumption_audit.go")
%w[ProductionAuthorizationConsumptionAuditPreview ProductionAuthorizationConsumptionConsumer ProductionAuthorizationConsumptionAuditCheck NewProductionAuthorizationConsumptionAuditPreview xnix.runtime.production_authorization_consumption_audit.v1 production-authorization-consumption-audit-preview production-gate-consolidated-authorization-consumption-audit production-authorization-consumption-audit-ready-authorization-disabled production-authorization-consumption-audit-blocked production-human-authorization-receipt-consolidation-preview production-dbus-gate-review-preview production-dbus-method-review-preview runtime-service-activation-preflight-preview runtime-write-gate-preview production-rollback-diagnostics-review-preview production-desktop-side-effect-review-preview consolidation-preview-consumed six-production-consumers-present consumers-use-consolidated-boundary authorization-not-accepted production-ownership-disabled write-and-launch-disabled desktop-and-support-side-effects-disabled host-boundary-closed production-dbus-gate-review production-dbus-method-review runtime-service-activation-preflight runtime-write-gate rollback-diagnostics-review desktop-side-effect-review].each do |token|
  assert(go_runtime_production_authorization_consumption_audit_source.include?(token), "Go Runtime production authorization consumption audit must include #{token}")
end
%w[ReceiptBoundaryConsolidated ConsolidationPreviewConsumed OwnerManagedOpaqueBoundaryReady CallerStateRootRequired AuthorizationAccepted ProductionReadiness ProductionOwnershipReady ConsumerCount RequiredConsumerCount ConsumedConsumerCount MissingConsumerCount AuthorizationAcceptedConsumerCount ProductionReadyConsumerCount SideEffectConsumerCount SystemServiceStarted SessionBusClaimed ProductionBusClaimed ProductionOwnerEnabled ProductionActivationReady WriteMethodsEnabled RuntimeWritesEnabled ReceiptWriterEnabled ReceiptPersistenceEnabled ReceiptLookupWritesEnabled DesktopFilesWritten MIMEAppsWritten ShellConfigurationWritten SettingsPersisted NotificationSent NotificationDeliveryEnabled PortalRequestCreated RequestObjectsCreated AdapterInvocationEnabled BackendLaunchEnabled BackendProcessStarted SupportBundleExported SupportCaseCreated SnapshotRestoreExecuted StateCleanupExecuted FileContentRead FilePathsExposed NetworkRequired HostRootModified PrivilegedContainerRequired StateRootPathExposed RawCommandExposed RawExecutableExposed BackendDetailsExposed validateNoBackendTerms].each do |token|
  assert(go_runtime_production_authorization_consumption_audit_source.include?(token), "Go Runtime production authorization consumption audit must expose safety gate #{token}")
end
%w[production_human_authorization_receipt_consolidation.go production_dbus_gate_review.go production_dbus_method_review.go runtime_service_activation_preflight.go runtime_write_gate.go production_rollback_diagnostics_review.go production_desktop_side_effect_review.go].each do |token|
  assert(go_runtime_production_authorization_consumption_audit_source.include?(token), "Go Runtime production authorization consumption audit must consume #{token}")
end
go_runtime_production_authorization_consumption_audit_test_source = read_project_file("internal/runtime/owner/production_authorization_consumption_audit_test.go") +
                                                                   read_project_file("cmd/xnix-runtime-go/production_dbus_gate_cli_test.go") +
                                                                   read_project_file("cmd/xnix-runtime-go/production_dbus_gate_commands.go") +
                                                                   read_project_file("cmd/xnix-runtime-go/main.go")
%w[TestProductionAuthorizationConsumptionAuditPreviewVerifiesConsumers TestProductionAuthorizationConsumptionAuditPreviewFailsClosedWithoutSources TestProductionAuthorizationConsumptionAuditPreviewCommand production-authorization-consumption-audit-preview receipt_boundary_consolidated consolidation_preview_consumed owner_managed_opaque_boundary_ready consumed_consumer_count missing_consumer_count authorization_accepted_consumer_count production_ready_consumer_count side_effect_consumer_count].each do |token|
  assert(go_runtime_production_authorization_consumption_audit_test_source.include?(token), "Go Runtime production authorization consumption audit tests and CLI must include #{token}")
end

go_runtime_production_receipt_acceptance_propagation_preflight_source = read_project_file("internal/runtime/owner/production_receipt_acceptance_propagation_preflight.go")
%w[ProductionReceiptAcceptancePropagationPreflightPreview ProductionReceiptAcceptancePropagationTarget ProductionReceiptAcceptancePropagationPreflightCheck NewProductionReceiptAcceptancePropagationPreflightPreview xnix.runtime.production_receipt_acceptance_propagation_preflight.v1 production-receipt-acceptance-propagation-preflight-preview future-authorization-receipt-acceptance-propagation-preflight production-receipt-acceptance-propagation-ready-acceptance-disabled production-receipt-acceptance-propagation-preflight-blocked production-authorization-consumption-audit-preview production-dbus-gate-review-preview production-dbus-method-review-preview runtime-service-activation-preflight-preview runtime-write-gate-preview production-rollback-diagnostics-review-preview production-desktop-side-effect-review-preview consumption-audit-consumed future-acceptance-modeled-only six-propagation-targets-present targets-propagate-future-acceptance acceptance-and-production-disabled production-ownership-disabled write-and-launch-disabled desktop-support-and-host-boundary-closed production-dbus-gate-review production-dbus-method-review runtime-service-activation-preflight runtime-write-gate rollback-diagnostics-review desktop-side-effect-review].each do |token|
  assert(go_runtime_production_receipt_acceptance_propagation_preflight_source.include?(token), "Go Runtime production receipt acceptance propagation preflight must include #{token}")
end
%w[FutureAcceptanceModeled AcceptanceSimulationOnly ConsumptionAuditConsumed OwnerManagedOpaqueBoundaryReady CallerStateRootRequired AuthorizationAccepted PropagationPreflightReady ProductionReadiness ProductionOwnershipReady TargetCount RequiredTargetCount PropagationReadyTargetCount MissingTargetCount AcceptanceEnabledTargetCount ProductionReadyTargetCount SideEffectTargetCount SystemServiceStarted SessionBusClaimed ProductionBusClaimed ProductionOwnerEnabled ProductionActivationReady WriteMethodsEnabled RuntimeWritesEnabled ReceiptWriterEnabled ReceiptPersistenceEnabled ReceiptLookupWritesEnabled DesktopFilesWritten MIMEAppsWritten ShellConfigurationWritten SettingsPersisted NotificationSent NotificationDeliveryEnabled PortalRequestCreated RequestObjectsCreated AdapterInvocationEnabled BackendLaunchEnabled BackendProcessStarted SupportBundleExported SupportCaseCreated SnapshotRestoreExecuted StateCleanupExecuted FileContentRead FilePathsExposed NetworkRequired HostRootModified PrivilegedContainerRequired StateRootPathExposed RawCommandExposed RawExecutableExposed BackendDetailsExposed validateNoBackendTerms].each do |token|
  assert(go_runtime_production_receipt_acceptance_propagation_preflight_source.include?(token), "Go Runtime production receipt acceptance propagation preflight must expose safety gate #{token}")
end
%w[production_authorization_consumption_audit.go production_dbus_gate_review.go production_dbus_method_review.go runtime_service_activation_preflight.go runtime_write_gate.go production_rollback_diagnostics_review.go production_desktop_side_effect_review.go].each do |token|
  assert(go_runtime_production_receipt_acceptance_propagation_preflight_source.include?(token), "Go Runtime production receipt acceptance propagation preflight must consume #{token}")
end
go_runtime_production_receipt_acceptance_propagation_preflight_test_source = read_project_file("internal/runtime/owner/production_receipt_acceptance_propagation_preflight_test.go") +
                                                                            read_project_file("cmd/xnix-runtime-go/production_dbus_gate_cli_test.go") +
                                                                            read_project_file("cmd/xnix-runtime-go/production_dbus_gate_commands.go") +
                                                                            read_project_file("cmd/xnix-runtime-go/main.go")
%w[TestProductionReceiptAcceptancePropagationPreflightPreviewModelsFutureAcceptance TestProductionReceiptAcceptancePropagationPreflightPreviewFailsClosedWithoutSources TestProductionReceiptAcceptancePropagationPreflightPreviewCommand production-receipt-acceptance-propagation-preflight-preview future_acceptance_modeled acceptance_simulation_only consumption_audit_consumed propagation_preflight_ready propagation_ready_target_count missing_target_count acceptance_enabled_target_count].each do |token|
  assert(go_runtime_production_receipt_acceptance_propagation_preflight_test_source.include?(token), "Go Runtime production receipt acceptance propagation preflight tests and CLI must include #{token}")
end

go_runtime_production_receipt_writer_authorization_review_source = read_project_file("internal/runtime/owner/production_receipt_writer_authorization_review.go")
%w[ProductionReceiptWriterAuthorizationReviewPreview ProductionReceiptWriterAuthorizationReviewItem ProductionReceiptWriterAuthorizationReviewCheck NewProductionReceiptWriterAuthorizationReviewPreview xnix.runtime.production_receipt_writer_authorization_review.v1 production-receipt-writer-authorization-review-preview receipt-writer-operator-authorization-boundary-review production-receipt-writer-authorization-review-ready-writes-disabled production-receipt-writer-authorization-review-blocked production-receipt-acceptance-propagation-preflight-preview production-human-authorization-receipt-consolidation-preview production-authorization-consumption-audit-preview acceptance-propagation-consumed writer-authorization-modeled-only five-review-items-present review-items-ready-writes-disabled receipt-writes-disabled production-ownership-disabled runtime-side-effects-disabled host-boundary-closed operator-action-boundary opaque-boundary-consolidated production-gates-consume-boundary persistence-disabled].each do |token|
  assert(go_runtime_production_receipt_writer_authorization_review_source.include?(token), "Go Runtime production receipt writer authorization review must include #{token}")
end
%w[WriterAuthorizationRequired WriterAuthorizationModeled OperatorActionRequired AcceptancePropagationConsumed OwnerManagedOpaqueBoundaryReady CallerStateRootRequired AuthorizationAccepted WriterReviewReady ProductionReadiness ProductionOwnershipReady ReviewItemCount RequiredReviewItemCount ReadyReviewItemCount MissingReviewItemCount WriteEnabledReviewItemCount AcceptanceEnabledReviewItemCount SideEffectReviewItemCount SystemServiceStarted SessionBusClaimed ProductionBusClaimed ProductionOwnerEnabled ProductionActivationReady WriteMethodsEnabled RuntimeWritesEnabled ReceiptWriterEnabled ReceiptPersistenceEnabled ReceiptLookupWritesEnabled ReceiptReplayEnabled DesktopFilesWritten MIMEAppsWritten ShellConfigurationWritten SettingsPersisted NotificationSent NotificationDeliveryEnabled PortalRequestCreated RequestObjectsCreated AdapterInvocationEnabled BackendLaunchEnabled BackendProcessStarted SupportBundleExported SupportCaseCreated SnapshotRestoreExecuted StateCleanupExecuted FileContentRead FilePathsExposed NetworkRequired HostRootModified PrivilegedContainerRequired StateRootPathExposed RawCommandExposed RawExecutableExposed BackendDetailsExposed validateNoBackendTerms].each do |token|
  assert(go_runtime_production_receipt_writer_authorization_review_source.include?(token), "Go Runtime production receipt writer authorization review must expose safety gate #{token}")
end
%w[production_receipt_acceptance_propagation_preflight.go production_human_authorization_receipt_consolidation.go production_authorization_consumption_audit.go].each do |token|
  assert(go_runtime_production_receipt_writer_authorization_review_source.include?(token), "Go Runtime production receipt writer authorization review must consume #{token}")
end
go_runtime_production_receipt_writer_authorization_review_test_source = read_project_file("internal/runtime/owner/production_receipt_writer_authorization_review_test.go") +
                                                                       read_project_file("cmd/xnix-runtime-go/production_dbus_gate_cli_test.go") +
                                                                       read_project_file("cmd/xnix-runtime-go/production_dbus_gate_commands.go") +
                                                                       read_project_file("cmd/xnix-runtime-go/main.go")
%w[TestProductionReceiptWriterAuthorizationReviewPreviewModelsWriterBoundary TestProductionReceiptWriterAuthorizationReviewPreviewFailsClosedWithoutSources TestProductionReceiptWriterAuthorizationReviewPreviewCommand production-receipt-writer-authorization-review-preview writer_authorization_required writer_authorization_modeled acceptance_propagation_consumed writer_review_ready ready_review_item_count write_enabled_review_item_count receipt_replay_enabled].each do |token|
  assert(go_runtime_production_receipt_writer_authorization_review_test_source.include?(token), "Go Runtime production receipt writer authorization review tests and CLI must include #{token}")
end

go_runtime_production_receipt_persistence_threat_review_source = read_project_file("internal/runtime/owner/production_receipt_persistence_threat_review.go")
%w[ProductionReceiptPersistenceThreatReviewPreview ProductionReceiptPersistenceThreatReviewItem ProductionReceiptPersistenceThreatReviewCheck NewProductionReceiptPersistenceThreatReviewPreview xnix.runtime.production_receipt_persistence_threat_review.v1 production-receipt-persistence-threat-review-preview receipt-persistence-expiry-revocation-replay-threat-review production-receipt-persistence-threat-review-ready-persistence-disabled production-receipt-persistence-threat-review-blocked production-receipt-writer-authorization-review-preview production-receipt-acceptance-propagation-preflight-preview writer-authorization-review-consumed persistence-threat-modeled-only five-threat-items-present threat-items-ready-persistence-disabled receipt-persistence-disabled production-ownership-disabled runtime-side-effects-disabled host-boundary-closed storage-confidentiality expiry-policy revocation-policy replay-protection audit-visibility].each do |token|
  assert(go_runtime_production_receipt_persistence_threat_review_source.include?(token), "Go Runtime production receipt persistence threat review must include #{token}")
end
%w[PersistenceThreatReviewRequired PersistenceThreatReviewModeled WriterAuthorizationReviewConsumed OwnerManagedOpaqueBoundaryReady CallerStateRootRequired AuthorizationAccepted PersistenceThreatReviewReady ProductionReadiness ProductionOwnershipReady ThreatItemCount RequiredThreatItemCount ReadyThreatItemCount MissingThreatItemCount PersistenceEnabledThreatCount ReplayEnabledThreatCount AcceptanceEnabledThreatCount SideEffectThreatCount SystemServiceStarted SessionBusClaimed ProductionBusClaimed ProductionOwnerEnabled ProductionActivationReady WriteMethodsEnabled RuntimeWritesEnabled ReceiptWriterEnabled ReceiptPersistenceEnabled ReceiptLookupWritesEnabled ReceiptReplayEnabled ReceiptExpiryWriteEnabled ReceiptRevocationWriteEnabled DesktopFilesWritten MIMEAppsWritten ShellConfigurationWritten SettingsPersisted NotificationSent NotificationDeliveryEnabled PortalRequestCreated RequestObjectsCreated AdapterInvocationEnabled BackendLaunchEnabled BackendProcessStarted SupportBundleExported SupportCaseCreated SnapshotRestoreExecuted StateCleanupExecuted FileContentRead FilePathsExposed NetworkRequired HostRootModified PrivilegedContainerRequired StateRootPathExposed RawCommandExposed RawExecutableExposed BackendDetailsExposed validateNoBackendTerms].each do |token|
  assert(go_runtime_production_receipt_persistence_threat_review_source.include?(token), "Go Runtime production receipt persistence threat review must expose safety gate #{token}")
end
%w[production_receipt_writer_authorization_review.go production_receipt_acceptance_propagation_preflight.go docs/claude-code-current-dispatch-picks.md].each do |token|
  assert(go_runtime_production_receipt_persistence_threat_review_source.include?(token), "Go Runtime production receipt persistence threat review must consume #{token}")
end
go_runtime_production_receipt_persistence_threat_review_test_source = read_project_file("internal/runtime/owner/production_receipt_persistence_threat_review_test.go") +
                                                                     read_project_file("cmd/xnix-runtime-go/production_dbus_gate_cli_test.go") +
                                                                     read_project_file("cmd/xnix-runtime-go/production_dbus_gate_commands.go") +
                                                                     read_project_file("cmd/xnix-runtime-go/main.go")
%w[TestProductionReceiptPersistenceThreatReviewPreviewModelsThreats TestProductionReceiptPersistenceThreatReviewPreviewFailsClosedWithoutSources TestProductionReceiptPersistenceThreatReviewPreviewCommand production-receipt-persistence-threat-review-preview persistence_threat_review_required persistence_threat_review_modeled writer_authorization_review_consumed persistence_threat_review_ready ready_threat_item_count persistence_enabled_threat_count receipt_expiry_write_enabled receipt_revocation_write_enabled].each do |token|
  assert(go_runtime_production_receipt_persistence_threat_review_test_source.include?(token), "Go Runtime production receipt persistence threat review tests and CLI must include #{token}")
end

go_runtime_production_receipt_revocation_visibility_audit_source = read_project_file("internal/runtime/owner/production_receipt_revocation_visibility_audit.go")
%w[ProductionReceiptRevocationVisibilityAuditPreview ProductionReceiptRevocationVisibilityAuditItem ProductionReceiptRevocationVisibilityAuditCheck NewProductionReceiptRevocationVisibilityAuditPreview xnix.runtime.production_receipt_revocation_visibility_audit.v1 production-receipt-revocation-visibility-audit-preview receipt-expiry-revocation-production-kde-visibility-audit production-receipt-revocation-visibility-audit-ready-visibility-only production-receipt-revocation-visibility-audit-blocked production-receipt-persistence-threat-review-preview production-desktop-side-effect-review-preview persistence-threat-review-consumed desktop-side-effect-review-consumed revocation-visibility-modeled-only five-visibility-items-present visibility-items-ready-writes-disabled revocation-expiry-writes-disabled notifications-and-desktop-side-effects-disabled production-ownership-disabled host-boundary-closed receipt-current-visible receipt-expiring-visible receipt-expired-visible receipt-revoked-visible receipt-missing-review-visible].each do |token|
  assert(go_runtime_production_receipt_revocation_visibility_audit_source.include?(token), "Go Runtime production receipt revocation visibility audit must include #{token}")
end
%w[VisibilityAuditRequired VisibilityAuditModeled PersistenceThreatReviewConsumed DesktopSideEffectReviewConsumed ProductionGateVisibilityModeled KDEStatusVisibilityModeled CallerStateRootRequired AuthorizationAccepted RevocationVisibilityReady ProductionReadiness ProductionOwnershipReady VisibilityItemCount RequiredVisibilityItemCount ReadyVisibilityItemCount MissingVisibilityItemCount RevocationWriteEnabledItemCount ExpiryWriteEnabledItemCount NotificationEnabledItemCount SideEffectItemCount SystemServiceStarted SessionBusClaimed ProductionBusClaimed ProductionOwnerEnabled ProductionActivationReady WriteMethodsEnabled RuntimeWritesEnabled ReceiptWriterEnabled ReceiptPersistenceEnabled ReceiptLookupWritesEnabled ReceiptReplayEnabled ReceiptExpiryWriteEnabled ReceiptRevocationWriteEnabled ReceiptRevocationVisibilityPersisted DesktopFilesWritten MIMEAppsWritten ShellConfigurationWritten SettingsPersisted KRunnerIndexPersisted TaskManagerEntryActive KWinRuleApplied LiveTrayBridgeEnabled TrayBridgePersisted NotificationSent NotificationDeliveryEnabled CompatibilityCenterPersisted PortalRequestCreated RequestObjectsCreated AdapterInvocationEnabled BackendLaunchEnabled BackendProcessStarted SupportBundleExported SupportCaseCreated SnapshotRestoreExecuted StateCleanupExecuted FileContentRead FilePathsExposed NetworkRequired HostRootModified PrivilegedContainerRequired StateRootPathExposed RawCommandExposed RawExecutableExposed BackendDetailsExposed validateNoBackendTerms].each do |token|
  assert(go_runtime_production_receipt_revocation_visibility_audit_source.include?(token), "Go Runtime production receipt revocation visibility audit must expose safety gate #{token}")
end
%w[production_receipt_persistence_threat_review.go production_desktop_side_effect_review.go docs/claude-code-current-dispatch-picks.md].each do |token|
  assert(go_runtime_production_receipt_revocation_visibility_audit_source.include?(token), "Go Runtime production receipt revocation visibility audit must consume #{token}")
end
go_runtime_production_receipt_revocation_visibility_audit_test_source = read_project_file("internal/runtime/owner/production_receipt_revocation_visibility_audit_test.go") +
                                                                       read_project_file("cmd/xnix-runtime-go/production_dbus_gate_cli_test.go") +
                                                                       read_project_file("cmd/xnix-runtime-go/production_dbus_gate_commands.go") +
                                                                       read_project_file("cmd/xnix-runtime-go/main.go")
%w[TestProductionReceiptRevocationVisibilityAuditPreviewModelsVisibility TestProductionReceiptRevocationVisibilityAuditPreviewFailsClosedWithoutSources TestProductionReceiptRevocationVisibilityAuditPreviewCommand production-receipt-revocation-visibility-audit-preview visibility_audit_required visibility_audit_modeled persistence_threat_review_consumed desktop_side_effect_review_consumed revocation_visibility_ready ready_visibility_item_count revocation_write_enabled_item_count expiry_write_enabled_item_count notification_enabled_item_count].each do |token|
  assert(go_runtime_production_receipt_revocation_visibility_audit_test_source.include?(token), "Go Runtime production receipt revocation visibility audit tests and CLI must include #{token}")
end

go_runtime_production_receipt_notification_delivery_gate_audit_source = read_project_file("internal/runtime/owner/production_receipt_notification_delivery_gate_audit.go")
%w[ProductionReceiptNotificationDeliveryGateAuditPreview ProductionReceiptNotificationDeliveryGateAuditItem ProductionReceiptNotificationDeliveryGateAuditCheck NewProductionReceiptNotificationDeliveryGateAuditPreview xnix.runtime.production_receipt_notification_delivery_gate_audit.v1 production-receipt-notification-delivery-gate-audit-preview receipt-expiry-revocation-notification-delivery-gate-audit production-receipt-notification-delivery-gate-audit-ready-delivery-disabled production-receipt-notification-delivery-gate-audit-blocked production-receipt-revocation-visibility-audit-preview production-desktop-side-effect-review-preview revocation-visibility-audit-consumed desktop-side-effect-review-consumed notification-gate-modeled-only five-notification-gate-items-present gate-items-ready-delivery-disabled notification-delivery-disabled requests-and-persistence-disabled receipt-writes-disabled production-and-host-boundary-closed expiring-receipt-warning-gate expired-receipt-blocker-gate revoked-receipt-blocker-gate missing-review-reminder-gate operator-approval-gate].each do |token|
  assert(go_runtime_production_receipt_notification_delivery_gate_audit_source.include?(token), "Go Runtime production receipt notification delivery gate audit must include #{token}")
end
%w[NotificationGateRequired NotificationGateModeled RevocationVisibilityAuditConsumed DesktopSideEffectReviewConsumed OperatorNotificationApprovalRequired OperatorNotificationApprovalPresent NotificationDeliveryGateReady CallerStateRootRequired AuthorizationAccepted ProductionReadiness ProductionOwnershipReady GateItemCount RequiredGateItemCount ReadyGateItemCount MissingGateItemCount DeliveryEnabledItemCount NotificationSentItemCount RequestCreatedItemCount SideEffectItemCount SystemServiceStarted SessionBusClaimed ProductionBusClaimed ProductionOwnerEnabled ProductionActivationReady WriteMethodsEnabled RuntimeWritesEnabled ReceiptWriterEnabled ReceiptPersistenceEnabled ReceiptLookupWritesEnabled ReceiptReplayEnabled ReceiptExpiryWriteEnabled ReceiptRevocationWriteEnabled ReceiptRevocationVisibilityPersisted NotificationDeliveryGatePersisted DesktopFilesWritten MIMEAppsWritten ShellConfigurationWritten SettingsPersisted KRunnerIndexPersisted TaskManagerEntryActive KWinRuleApplied LiveTrayBridgeEnabled TrayBridgePersisted NotificationSent NotificationDeliveryEnabled NotificationActionEnabled CompatibilityCenterPersisted PortalRequestCreated RequestObjectsCreated AdapterInvocationEnabled BackendLaunchEnabled BackendProcessStarted SupportBundleExported SupportCaseCreated SnapshotRestoreExecuted StateCleanupExecuted FileContentRead FilePathsExposed NetworkRequired HostRootModified PrivilegedContainerRequired StateRootPathExposed RawCommandExposed RawExecutableExposed BackendDetailsExposed validateNoBackendTerms].each do |token|
  assert(go_runtime_production_receipt_notification_delivery_gate_audit_source.include?(token), "Go Runtime production receipt notification delivery gate audit must expose safety gate #{token}")
end
%w[production_receipt_revocation_visibility_audit.go production_desktop_side_effect_review.go docs/claude-code-current-dispatch-picks.md].each do |token|
  assert(go_runtime_production_receipt_notification_delivery_gate_audit_source.include?(token), "Go Runtime production receipt notification delivery gate audit must consume #{token}")
end
go_runtime_production_receipt_notification_delivery_gate_audit_test_source = read_project_file("internal/runtime/owner/production_receipt_notification_delivery_gate_audit_test.go") +
                                                                            read_project_file("cmd/xnix-runtime-go/production_dbus_gate_cli_test.go") +
                                                                            read_project_file("cmd/xnix-runtime-go/production_dbus_gate_commands.go") +
                                                                            read_project_file("cmd/xnix-runtime-go/main.go")
%w[TestProductionReceiptNotificationDeliveryGateAuditPreviewModelsGate TestProductionReceiptNotificationDeliveryGateAuditPreviewFailsClosedWithoutSources TestProductionReceiptNotificationDeliveryGateAuditPreviewCommand production-receipt-notification-delivery-gate-audit-preview notification_gate_required notification_gate_modeled revocation_visibility_audit_consumed notification_delivery_gate_ready ready_gate_item_count delivery_enabled_item_count notification_sent_item_count request_created_item_count].each do |token|
  assert(go_runtime_production_receipt_notification_delivery_gate_audit_test_source.include?(token), "Go Runtime production receipt notification delivery gate audit tests and CLI must include #{token}")
end

go_runtime_production_receipt_notification_action_safety_audit_source = read_project_file("internal/runtime/owner/production_receipt_notification_action_safety_audit.go")
%w[ProductionReceiptNotificationActionSafetyAuditPreview ProductionReceiptNotificationActionSafetyAuditItem ProductionReceiptNotificationActionSafetyAuditCheck NewProductionReceiptNotificationActionSafetyAuditPreview xnix.runtime.production_receipt_notification_action_safety_audit.v1 production-receipt-notification-action-safety-audit-preview receipt-notification-action-request-safety-audit production-receipt-notification-action-safety-audit-ready-actions-disabled production-receipt-notification-action-safety-audit-blocked production-receipt-notification-delivery-gate-audit-preview production-desktop-side-effect-review-preview notification-delivery-gate-consumed desktop-side-effect-review-consumed action-safety-modeled-only five-action-items-present action-items-ready-actions-disabled notification-actions-disabled requests-and-navigation-disabled receipt-and-notification-writes-disabled production-and-host-boundary-closed review-receipt-action renew-receipt-action open-compatibility-center-action dismiss-receipt-action support-info-action].each do |token|
  assert(go_runtime_production_receipt_notification_action_safety_audit_source.include?(token), "Go Runtime production receipt notification action safety audit must include #{token}")
end
%w[ActionSafetyAuditRequired ActionSafetyAuditModeled NotificationDeliveryGateConsumed DesktopSideEffectReviewConsumed OperatorActionApprovalRequired OperatorActionApprovalPresent NotificationActionSafetyReady CallerStateRootRequired AuthorizationAccepted ProductionReadiness ProductionOwnershipReady ActionItemCount RequiredActionItemCount ReadyActionItemCount MissingActionItemCount EnabledActionItemCount RequestCreatedItemCount NavigationEnabledItemCount SideEffectItemCount SystemServiceStarted SessionBusClaimed ProductionBusClaimed ProductionOwnerEnabled ProductionActivationReady WriteMethodsEnabled RuntimeWritesEnabled ReceiptWriterEnabled ReceiptPersistenceEnabled ReceiptLookupWritesEnabled ReceiptReplayEnabled ReceiptExpiryWriteEnabled ReceiptRevocationWriteEnabled NotificationSent NotificationDeliveryEnabled NotificationActionEnabled NotificationActionSafetyPersisted ReviewActionEnabled RenewActionEnabled OpenCompatibilityCenterEnabled DismissActionEnabled SupportInfoActionEnabled CompatibilityCenterOpened CompatibilityCenterPersisted PortalRequestCreated RequestObjectsCreated DesktopFilesWritten MIMEAppsWritten ShellConfigurationWritten SettingsPersisted KRunnerIndexPersisted TaskManagerEntryActive KWinRuleApplied LiveTrayBridgeEnabled TrayBridgePersisted AdapterInvocationEnabled BackendLaunchEnabled BackendProcessStarted SupportBundleExported SupportCaseCreated SnapshotRestoreExecuted StateCleanupExecuted FileContentRead FilePathsExposed NetworkRequired HostRootModified PrivilegedContainerRequired StateRootPathExposed RawCommandExposed RawExecutableExposed BackendDetailsExposed validateNoBackendTerms].each do |token|
  assert(go_runtime_production_receipt_notification_action_safety_audit_source.include?(token), "Go Runtime production receipt notification action safety audit must expose safety gate #{token}")
end
%w[production_receipt_notification_delivery_gate_audit.go production_desktop_side_effect_review.go docs/claude-code-current-dispatch-picks.md].each do |token|
  assert(go_runtime_production_receipt_notification_action_safety_audit_source.include?(token), "Go Runtime production receipt notification action safety audit must consume #{token}")
end
go_runtime_production_receipt_notification_action_safety_audit_test_source = read_project_file("internal/runtime/owner/production_receipt_notification_action_safety_audit_test.go") +
                                                                            read_project_file("cmd/xnix-runtime-go/production_dbus_gate_cli_test.go") +
                                                                            read_project_file("cmd/xnix-runtime-go/production_dbus_gate_commands.go") +
                                                                            read_project_file("cmd/xnix-runtime-go/main.go")
%w[TestProductionReceiptNotificationActionSafetyAuditPreviewModelsActions TestProductionReceiptNotificationActionSafetyAuditPreviewFailsClosedWithoutSources TestProductionReceiptNotificationActionSafetyAuditPreviewCommand production-receipt-notification-action-safety-audit-preview action_safety_audit_required action_safety_audit_modeled notification_delivery_gate_consumed notification_action_safety_ready ready_action_item_count enabled_action_item_count navigation_enabled_item_count review_action_enabled renew_action_enabled open_compatibility_center_enabled].each do |token|
  assert(go_runtime_production_receipt_notification_action_safety_audit_test_source.include?(token), "Go Runtime production receipt notification action safety audit tests and CLI must include #{token}")
end

go_runtime_production_receipt_notification_action_request_object_audit_source = read_project_file("internal/runtime/owner/production_receipt_notification_action_request_object_audit.go")
%w[ProductionReceiptNotificationActionRequestObjectAuditPreview ProductionReceiptNotificationActionRequestObjectAuditItem ProductionReceiptNotificationActionRequestObjectAuditCheck NewProductionReceiptNotificationActionRequestObjectAuditPreview xnix.runtime.production_receipt_notification_action_request_object_audit.v1 production-receipt-notification-action-request-object-audit-preview receipt-notification-action-runtime-request-object-audit production-receipt-notification-action-request-object-audit-ready-requests-disabled production-receipt-notification-action-request-object-audit-blocked production-receipt-notification-action-safety-audit-preview windows-app-compatibility-implementation-brief action-safety-audit-consumed dispatch-request-boundary-consumed request-object-modeled-only five-request-objects-present request-objects-ready-creation-disabled request-creation-disabled actions-navigation-and-portal-disabled receipt-notification-and-support-writes-disabled production-and-host-boundary-closed review-receipt-request-object renew-receipt-request-object open-compatibility-center-request-object dismiss-receipt-request-object support-info-request-object].each do |token|
  assert(go_runtime_production_receipt_notification_action_request_object_audit_source.include?(token), "Go Runtime production receipt notification action request-object audit must include #{token}")
end
%w[RequestObjectAuditRequired RequestObjectAuditModeled ActionSafetyAuditConsumed DispatchRequestBoundaryConsumed OperatorActionApprovalRequired OperatorActionApprovalPresent RequestObjectBoundaryReady CallerStateRootRequired AuthorizationAccepted ProductionReadiness ProductionOwnershipReady RequestObjectCount RequiredRequestObjectCount ReadyRequestObjectCount MissingRequestObjectCount CreatedRequestObjectCount DispatchedRequestObjectCount PortalRequestCreatedCount NavigationRequestedCount SideEffectRequestObjectCount SystemServiceStarted SessionBusClaimed ProductionBusClaimed ProductionOwnerEnabled ProductionActivationReady WriteMethodsEnabled RuntimeWritesEnabled RequestObjectCreationEnabled RequestObjectDispatchEnabled RequestObjectPersistenceEnabled PortalRequestCreated RequestObjectsCreated RequestObjectsDispatched NotificationSent NotificationDeliveryEnabled NotificationActionEnabled ReviewActionEnabled RenewActionEnabled OpenCompatibilityCenterEnabled DismissActionEnabled SupportInfoActionEnabled CompatibilityCenterOpened CompatibilityCenterPersisted ReceiptWriterEnabled ReceiptPersistenceEnabled ReceiptLookupWritesEnabled ReceiptReplayEnabled ReceiptExpiryWriteEnabled ReceiptRevocationWriteEnabled DesktopFilesWritten MIMEAppsWritten ShellConfigurationWritten SettingsPersisted KRunnerIndexPersisted TaskManagerEntryActive KWinRuleApplied LiveTrayBridgeEnabled TrayBridgePersisted AdapterInvocationEnabled BackendLaunchEnabled BackendProcessStarted SupportBundleExported SupportCaseCreated SnapshotRestoreExecuted StateCleanupExecuted FileContentRead FilePathsExposed NetworkRequired HostRootModified PrivilegedContainerRequired StateRootPathExposed RawCommandExposed RawExecutableExposed BackendDetailsExposed validateNoBackendTerms].each do |token|
  assert(go_runtime_production_receipt_notification_action_request_object_audit_source.include?(token), "Go Runtime production receipt notification action request-object audit must expose safety gate #{token}")
end
%w[production_receipt_notification_action_safety_audit.go docs/claude-code-current-dispatch-picks.md docs/windows-app-compatibility-implementation-brief.md].each do |token|
  assert(go_runtime_production_receipt_notification_action_request_object_audit_source.include?(token), "Go Runtime production receipt notification action request-object audit must consume #{token}")
end
go_runtime_production_receipt_notification_action_request_object_audit_test_source = read_project_file("internal/runtime/owner/production_receipt_notification_action_request_object_audit_test.go") +
                                                                                   read_project_file("cmd/xnix-runtime-go/production_dbus_gate_cli_test.go") +
                                                                                   read_project_file("cmd/xnix-runtime-go/production_dbus_gate_commands.go") +
                                                                                   read_project_file("cmd/xnix-runtime-go/main.go")
%w[TestProductionReceiptNotificationActionRequestObjectAuditPreviewModelsRequests TestProductionReceiptNotificationActionRequestObjectAuditPreviewFailsClosedWithoutSources TestProductionReceiptNotificationActionRequestObjectAuditPreviewCommand production-receipt-notification-action-request-object-audit-preview request_object_audit_required request_object_audit_modeled action_safety_audit_consumed dispatch_request_boundary_consumed request_object_boundary_ready ready_request_object_count created_request_object_count dispatched_request_object_count portal_request_created_count navigation_requested_count request_object_creation_enabled request_object_dispatch_enabled].each do |token|
  assert(go_runtime_production_receipt_notification_action_request_object_audit_test_source.include?(token), "Go Runtime production receipt notification action request-object audit tests and CLI must include #{token}")
end

go_runtime_production_receipt_notification_action_dispatch_authorization_audit_source = read_project_file("internal/runtime/owner/production_receipt_notification_action_dispatch_authorization_audit.go")
%w[ProductionReceiptNotificationActionDispatchAuthorizationAuditPreview ProductionReceiptNotificationActionDispatchAuthorizationAuditItem ProductionReceiptNotificationActionDispatchAuthorizationAuditCheck NewProductionReceiptNotificationActionDispatchAuthorizationAuditPreview xnix.runtime.production_receipt_notification_action_dispatch_authorization_audit.v1 production-receipt-notification-action-dispatch-authorization-audit-preview receipt-notification-action-request-dispatch-authorization-audit production-receipt-notification-action-dispatch-authorization-audit-ready-dispatch-disabled production-receipt-notification-action-dispatch-authorization-audit-blocked production-receipt-notification-action-request-object-audit-preview production-authorization-consumption-audit-preview request-object-audit-consumed receipt-authorization-boundary-consumed dispatch-authorization-modeled-only five-dispatch-authorizations-present dispatch-authorizations-ready-dispatch-disabled dispatch-disabled request-creation-actions-and-portal-disabled receipt-notification-navigation-and-support-writes-disabled production-and-host-boundary-closed review-receipt-dispatch-authorization renew-receipt-dispatch-authorization open-compatibility-center-dispatch-authorization dismiss-receipt-dispatch-authorization support-info-dispatch-authorization].each do |token|
  assert(go_runtime_production_receipt_notification_action_dispatch_authorization_audit_source.include?(token), "Go Runtime production receipt notification action dispatch authorization audit must include #{token}")
end
%w[DispatchAuthorizationAuditRequired DispatchAuthorizationAuditModeled RequestObjectAuditConsumed ReceiptAuthorizationBoundaryConsumed OperatorDispatchApprovalRequired OperatorDispatchApprovalPresent DispatchAuthorizationReady DispatchAuthorizationGranted CallerStateRootRequired AuthorizationAccepted ProductionReadiness ProductionOwnershipReady AuthorizationItemCount RequiredAuthorizationItemCount ReadyAuthorizationItemCount MissingAuthorizationItemCount GrantedAuthorizationItemCount DispatchedAuthorizationItemCount CreatedRequestObjectCount PortalRequestCreatedCount SideEffectAuthorizationItemCount SystemServiceStarted SessionBusClaimed ProductionBusClaimed ProductionOwnerEnabled ProductionActivationReady WriteMethodsEnabled RuntimeWritesEnabled RequestObjectCreationEnabled RequestObjectDispatchEnabled RequestObjectPersistenceEnabled DispatchAuthorizationPersisted PortalRequestCreated RequestObjectsCreated RequestObjectsDispatched NotificationSent NotificationDeliveryEnabled NotificationActionEnabled ReviewActionEnabled RenewActionEnabled OpenCompatibilityCenterEnabled DismissActionEnabled SupportInfoActionEnabled CompatibilityCenterOpened CompatibilityCenterPersisted ReceiptWriterEnabled ReceiptPersistenceEnabled ReceiptLookupWritesEnabled ReceiptReplayEnabled ReceiptExpiryWriteEnabled ReceiptRevocationWriteEnabled DesktopFilesWritten MIMEAppsWritten ShellConfigurationWritten SettingsPersisted KRunnerIndexPersisted TaskManagerEntryActive KWinRuleApplied LiveTrayBridgeEnabled TrayBridgePersisted AdapterInvocationEnabled BackendLaunchEnabled BackendProcessStarted SupportBundleExported SupportCaseCreated SnapshotRestoreExecuted StateCleanupExecuted FileContentRead FilePathsExposed NetworkRequired HostRootModified PrivilegedContainerRequired StateRootPathExposed RawCommandExposed RawExecutableExposed BackendDetailsExposed validateNoBackendTerms].each do |token|
  assert(go_runtime_production_receipt_notification_action_dispatch_authorization_audit_source.include?(token), "Go Runtime production receipt notification action dispatch authorization audit must expose safety gate #{token}")
end
%w[production_receipt_notification_action_request_object_audit.go production_authorization_consumption_audit.go docs/claude-code-current-dispatch-picks.md].each do |token|
  assert(go_runtime_production_receipt_notification_action_dispatch_authorization_audit_source.include?(token), "Go Runtime production receipt notification action dispatch authorization audit must consume #{token}")
end
go_runtime_production_receipt_notification_action_dispatch_authorization_audit_test_source = read_project_file("internal/runtime/owner/production_receipt_notification_action_dispatch_authorization_audit_test.go") +
                                                                                           read_project_file("cmd/xnix-runtime-go/production_dbus_gate_cli_test.go") +
                                                                                           read_project_file("cmd/xnix-runtime-go/production_dbus_gate_commands.go") +
                                                                                           read_project_file("cmd/xnix-runtime-go/main.go")
%w[TestProductionReceiptNotificationActionDispatchAuthorizationAuditPreviewModelsDispatch TestProductionReceiptNotificationActionDispatchAuthorizationAuditPreviewFailsClosedWithoutSources TestProductionReceiptNotificationActionDispatchAuthorizationAuditPreviewCommand production-receipt-notification-action-dispatch-authorization-audit-preview dispatch_authorization_audit_required dispatch_authorization_audit_modeled request_object_audit_consumed receipt_authorization_boundary_consumed dispatch_authorization_ready dispatch_authorization_granted ready_authorization_item_count granted_authorization_item_count dispatched_authorization_item_count request_object_dispatch_enabled dispatch_authorization_persisted].each do |token|
  assert(go_runtime_production_receipt_notification_action_dispatch_authorization_audit_test_source.include?(token), "Go Runtime production receipt notification action dispatch authorization audit tests and CLI must include #{token}")
end

go_runtime_production_receipt_notification_action_dispatch_dry_run_audit_source = read_project_file("internal/runtime/owner/production_receipt_notification_action_dispatch_dry_run_audit.go")
%w[ProductionReceiptNotificationActionDispatchDryRunAuditPreview ProductionReceiptNotificationActionDispatchDryRunAuditItem ProductionReceiptNotificationActionDispatchDryRunAuditCheck NewProductionReceiptNotificationActionDispatchDryRunAuditPreview xnix.runtime.production_receipt_notification_action_dispatch_dry_run_audit.v1 production-receipt-notification-action-dispatch-dry-run-audit-preview receipt-notification-action-request-dispatch-dry-run-audit production-receipt-notification-action-dispatch-dry-run-audit-ready-dry-run-only production-receipt-notification-action-dispatch-dry-run-audit-blocked production-receipt-notification-action-dispatch-authorization-audit-preview dispatch-authorization-audit-consumed dispatch-dry-run-guidance-consumed dispatch-dry-run-modeled-only five-dispatch-dry-run-items-present dispatch-dry-run-items-ready-execution-disabled dry-run-execution-disabled request-creation-dispatch-actions-and-portal-disabled receipt-notification-navigation-and-support-writes-disabled production-and-host-boundary-closed review-receipt-dispatch-dry-run renew-receipt-dispatch-dry-run open-compatibility-center-dispatch-dry-run dismiss-receipt-dispatch-dry-run support-info-dispatch-dry-run].each do |token|
  assert(go_runtime_production_receipt_notification_action_dispatch_dry_run_audit_source.include?(token), "Go Runtime production receipt notification action dispatch dry-run audit must include #{token}")
end
%w[DispatchDryRunAuditRequired DispatchDryRunAuditModeled DispatchAuthorizationAuditConsumed DispatchDryRunGuidanceConsumed DispatchDryRunPlanReady DispatchDryRunExecuted DispatchAuthorizationGranted OperatorDispatchApprovalRequired OperatorDispatchApprovalPresent CallerStateRootRequired AuthorizationAccepted ProductionReadiness ProductionOwnershipReady DryRunItemCount RequiredDryRunItemCount ReadyDryRunItemCount MissingDryRunItemCount ExecutedDryRunItemCount DispatchedDryRunItemCount CreatedRequestObjectCount PortalRequestCreatedCount SideEffectDryRunItemCount SystemServiceStarted SessionBusClaimed ProductionBusClaimed ProductionOwnerEnabled ProductionActivationReady WriteMethodsEnabled RuntimeWritesEnabled RequestObjectCreationEnabled RequestObjectDispatchEnabled RequestObjectPersistenceEnabled DispatchAuthorizationPersisted DispatchDryRunExecutionEnabled DryRunResultPersisted PortalRequestCreated RequestObjectsCreated RequestObjectsDispatched NotificationSent NotificationDeliveryEnabled NotificationActionEnabled ReviewActionEnabled RenewActionEnabled OpenCompatibilityCenterEnabled DismissActionEnabled SupportInfoActionEnabled CompatibilityCenterOpened CompatibilityCenterPersisted ReceiptWriterEnabled ReceiptPersistenceEnabled ReceiptLookupWritesEnabled ReceiptReplayEnabled ReceiptExpiryWriteEnabled ReceiptRevocationWriteEnabled DesktopFilesWritten MIMEAppsWritten ShellConfigurationWritten SettingsPersisted KRunnerIndexPersisted TaskManagerEntryActive KWinRuleApplied LiveTrayBridgeEnabled TrayBridgePersisted AdapterInvocationEnabled BackendLaunchEnabled BackendProcessStarted SupportBundleExported SupportCaseCreated SnapshotRestoreExecuted StateCleanupExecuted FileContentRead FilePathsExposed NetworkRequired HostRootModified PrivilegedContainerRequired StateRootPathExposed RawCommandExposed RawExecutableExposed BackendDetailsExposed validateNoBackendTerms].each do |token|
  assert(go_runtime_production_receipt_notification_action_dispatch_dry_run_audit_source.include?(token), "Go Runtime production receipt notification action dispatch dry-run audit must expose safety gate #{token}")
end
%w[production_receipt_notification_action_dispatch_authorization_audit.go docs/claude-code-current-dispatch-picks.md].each do |token|
  assert(go_runtime_production_receipt_notification_action_dispatch_dry_run_audit_source.include?(token), "Go Runtime production receipt notification action dispatch dry-run audit must consume #{token}")
end
go_runtime_production_receipt_notification_action_dispatch_dry_run_audit_test_source = read_project_file("internal/runtime/owner/production_receipt_notification_action_dispatch_dry_run_audit_test.go") +
                                                                                      read_project_file("cmd/xnix-runtime-go/production_dbus_gate_cli_test.go") +
                                                                                      read_project_file("cmd/xnix-runtime-go/production_dbus_gate_commands.go") +
                                                                                      read_project_file("cmd/xnix-runtime-go/main.go")
%w[TestProductionReceiptNotificationActionDispatchDryRunAuditPreviewModelsDryRun TestProductionReceiptNotificationActionDispatchDryRunAuditPreviewFailsClosedWithoutSources TestProductionReceiptNotificationActionDispatchDryRunAuditPreviewCommand production-receipt-notification-action-dispatch-dry-run-audit-preview dispatch_dry_run_audit_required dispatch_dry_run_audit_modeled dispatch_authorization_audit_consumed dispatch_dry_run_guidance_consumed dispatch_dry_run_plan_ready dispatch_dry_run_executed ready_dry_run_item_count executed_dry_run_item_count dispatched_dry_run_item_count dispatch_dry_run_execution_enabled dry_run_result_persisted].each do |token|
  assert(go_runtime_production_receipt_notification_action_dispatch_dry_run_audit_test_source.include?(token), "Go Runtime production receipt notification action dispatch dry-run audit tests and CLI must include #{token}")
end

go_runtime_production_receipt_notification_action_dry_run_result_visibility_audit_source = read_project_file("internal/runtime/owner/production_receipt_notification_action_dry_run_result_visibility_audit.go")
%w[ProductionReceiptNotificationActionDryRunResultVisibilityAuditPreview ProductionReceiptNotificationActionDryRunResultVisibilityAuditItem ProductionReceiptNotificationActionDryRunResultVisibilityAuditCheck NewProductionReceiptNotificationActionDryRunResultVisibilityAuditPreview xnix.runtime.production_receipt_notification_action_dry_run_result_visibility_audit.v1 production-receipt-notification-action-dry-run-result-visibility-audit-preview receipt-notification-action-dry-run-result-visibility-audit production-receipt-notification-action-dry-run-result-visibility-audit-ready-visibility-only production-receipt-notification-action-dry-run-result-visibility-audit-blocked production-receipt-notification-action-dispatch-dry-run-audit-preview dispatch-dry-run-audit-consumed dry-run-result-visibility-guidance-consumed dry-run-result-visibility-modeled-only five-dry-run-result-visibility-items-present visibility-items-ready-results-redacted dry-run-execution-and-result-persistence-disabled request-creation-dispatch-actions-and-portal-disabled receipt-notification-navigation-and-support-writes-disabled production-and-host-boundary-closed review-receipt-dry-run-result-visibility renew-receipt-dry-run-result-visibility open-compatibility-center-dry-run-result-visibility dismiss-receipt-dry-run-result-visibility support-info-dry-run-result-visibility].each do |token|
  assert(go_runtime_production_receipt_notification_action_dry_run_result_visibility_audit_source.include?(token), "Go Runtime production receipt notification action dry-run result visibility audit must include #{token}")
end
%w[ResultVisibilityAuditRequired ResultVisibilityAuditModeled DispatchDryRunAuditConsumed ResultVisibilityGuidanceConsumed ResultVisibilityPlanReady RedactedKDEVisibilityModeled RuntimeDiagnosticsVisibilityModeled DispatchDryRunExecuted DryRunResultPersisted DispatchAuthorizationGranted OperatorDispatchApprovalRequired OperatorDispatchApprovalPresent CallerStateRootRequired AuthorizationAccepted ProductionReadiness ProductionOwnershipReady VisibilityItemCount RequiredVisibilityItemCount ReadyVisibilityItemCount MissingVisibilityItemCount PersistedVisibilityItemCount ExecutedDryRunItemCount DispatchedDryRunItemCount CreatedRequestObjectCount PortalRequestCreatedCount SideEffectVisibilityItemCount SystemServiceStarted SessionBusClaimed ProductionBusClaimed ProductionOwnerEnabled ProductionActivationReady WriteMethodsEnabled RuntimeWritesEnabled RequestObjectCreationEnabled RequestObjectDispatchEnabled RequestObjectPersistenceEnabled DispatchAuthorizationPersisted DispatchDryRunExecutionEnabled DryRunResultPersistenceEnabled ResultVisibilityPersistenceEnabled RuntimeDiagnosticsPersisted PortalRequestCreated RequestObjectsCreated RequestObjectsDispatched NotificationSent NotificationDeliveryEnabled NotificationActionEnabled ReviewActionEnabled RenewActionEnabled OpenCompatibilityCenterEnabled DismissActionEnabled SupportInfoActionEnabled CompatibilityCenterOpened CompatibilityCenterPersisted ReceiptWriterEnabled ReceiptPersistenceEnabled ReceiptLookupWritesEnabled ReceiptReplayEnabled ReceiptExpiryWriteEnabled ReceiptRevocationWriteEnabled DesktopFilesWritten MIMEAppsWritten ShellConfigurationWritten SettingsPersisted KRunnerIndexPersisted TaskManagerEntryActive KWinRuleApplied LiveTrayBridgeEnabled TrayBridgePersisted AdapterInvocationEnabled BackendLaunchEnabled BackendProcessStarted SupportBundleExported SupportCaseCreated SnapshotRestoreExecuted StateCleanupExecuted FileContentRead FilePathsExposed NetworkRequired HostRootModified PrivilegedContainerRequired StateRootPathExposed RawCommandExposed RawExecutableExposed BackendDetailsExposed validateNoBackendTerms].each do |token|
  assert(go_runtime_production_receipt_notification_action_dry_run_result_visibility_audit_source.include?(token), "Go Runtime production receipt notification action dry-run result visibility audit must expose safety gate #{token}")
end
%w[production_receipt_notification_action_dispatch_dry_run_audit.go docs/claude-code-current-dispatch-picks.md].each do |token|
  assert(go_runtime_production_receipt_notification_action_dry_run_result_visibility_audit_source.include?(token), "Go Runtime production receipt notification action dry-run result visibility audit must consume #{token}")
end
go_runtime_production_receipt_notification_action_dry_run_result_visibility_audit_test_source = read_project_file("internal/runtime/owner/production_receipt_notification_action_dry_run_result_visibility_audit_test.go") +
                                                                                              read_project_file("cmd/xnix-runtime-go/production_dbus_gate_cli_test.go") +
                                                                                              read_project_file("cmd/xnix-runtime-go/production_dbus_gate_commands.go") +
                                                                                              read_project_file("cmd/xnix-runtime-go/main.go")
%w[TestProductionReceiptNotificationActionDryRunResultVisibilityAuditPreviewModelsVisibility TestProductionReceiptNotificationActionDryRunResultVisibilityAuditPreviewFailsClosedWithoutSources TestProductionReceiptNotificationActionDryRunResultVisibilityAuditPreviewCommand production-receipt-notification-action-dry-run-result-visibility-audit-preview result_visibility_audit_required result_visibility_audit_modeled dispatch_dry_run_audit_consumed result_visibility_guidance_consumed result_visibility_plan_ready redacted_kde_visibility_modeled runtime_diagnostics_visibility_modeled ready_visibility_item_count persisted_visibility_item_count dry_run_result_persistence_enabled result_visibility_persistence_enabled runtime_diagnostics_persisted].each do |token|
  assert(go_runtime_production_receipt_notification_action_dry_run_result_visibility_audit_test_source.include?(token), "Go Runtime production receipt notification action dry-run result visibility audit tests and CLI must include #{token}")
end

go_runtime_production_receipt_notification_action_dry_run_result_persistence_authorization_audit_source = read_project_file("internal/runtime/owner/production_receipt_notification_action_dry_run_result_persistence_authorization_audit.go")
%w[ProductionReceiptNotificationActionDryRunResultPersistenceAuthorizationAuditPreview ProductionReceiptNotificationActionDryRunResultPersistenceAuthorizationAuditItem ProductionReceiptNotificationActionDryRunResultPersistenceAuthorizationAuditCheck NewProductionReceiptNotificationActionDryRunResultPersistenceAuthorizationAuditPreview xnix.runtime.production_receipt_notification_action_dry_run_result_persistence_authorization_audit.v1 production-receipt-notification-action-dry-run-result-persistence-authorization-audit-preview receipt-notification-action-dry-run-result-persistence-authorization-audit production-receipt-notification-action-dry-run-result-persistence-authorization-audit-ready-persistence-disabled production-receipt-notification-action-dry-run-result-persistence-authorization-audit-blocked production-receipt-notification-action-dry-run-result-visibility-audit-preview result-visibility-audit-consumed persistence-authorization-guidance-consumed persistence-authorization-modeled-only five-persistence-authorizations-present persistence-authorizations-ready-persistence-disabled dry-run-execution-result-and-visibility-persistence-disabled request-creation-dispatch-actions-and-portal-disabled receipt-notification-navigation-and-support-writes-disabled production-and-host-boundary-closed review-receipt-dry-run-result-persistence-authorization renew-receipt-dry-run-result-persistence-authorization open-compatibility-center-dry-run-result-persistence-authorization dismiss-receipt-dry-run-result-persistence-authorization support-info-dry-run-result-persistence-authorization].each do |token|
  assert(go_runtime_production_receipt_notification_action_dry_run_result_persistence_authorization_audit_source.include?(token), "Go Runtime production receipt notification action dry-run result persistence authorization audit must include #{token}")
end
%w[PersistenceAuthorizationAuditRequired PersistenceAuthorizationAuditModeled ResultVisibilityAuditConsumed PersistenceAuthorizationGuidanceConsumed PersistenceAuthorizationReady ResultPersistenceAuthorized DispatchDryRunExecuted DryRunResultPersisted ResultVisibilityPersisted DispatchAuthorizationGranted OperatorPersistenceApprovalRequired OperatorPersistenceApprovalPresent CallerStateRootRequired AuthorizationAccepted ProductionReadiness ProductionOwnershipReady AuthorizationItemCount RequiredAuthorizationItemCount ReadyAuthorizationItemCount MissingAuthorizationItemCount GrantedPersistenceItemCount PersistedResultItemCount ExecutedDryRunItemCount CreatedRequestObjectCount PortalRequestCreatedCount SideEffectAuthorizationItemCount SystemServiceStarted SessionBusClaimed ProductionBusClaimed ProductionOwnerEnabled ProductionActivationReady WriteMethodsEnabled RuntimeWritesEnabled RequestObjectCreationEnabled RequestObjectDispatchEnabled RequestObjectPersistenceEnabled DispatchAuthorizationPersisted DispatchDryRunExecutionEnabled DryRunResultPersistenceEnabled ResultVisibilityPersistenceEnabled RuntimeDiagnosticsPersisted PortalRequestCreated RequestObjectsCreated RequestObjectsDispatched NotificationSent NotificationDeliveryEnabled NotificationActionEnabled ReviewActionEnabled RenewActionEnabled OpenCompatibilityCenterEnabled DismissActionEnabled SupportInfoActionEnabled CompatibilityCenterOpened CompatibilityCenterPersisted ReceiptWriterEnabled ReceiptPersistenceEnabled ReceiptLookupWritesEnabled ReceiptReplayEnabled ReceiptExpiryWriteEnabled ReceiptRevocationWriteEnabled DesktopFilesWritten MIMEAppsWritten ShellConfigurationWritten SettingsPersisted KRunnerIndexPersisted TaskManagerEntryActive KWinRuleApplied LiveTrayBridgeEnabled TrayBridgePersisted AdapterInvocationEnabled BackendLaunchEnabled BackendProcessStarted SupportBundleExported SupportCaseCreated SnapshotRestoreExecuted StateCleanupExecuted FileContentRead FilePathsExposed NetworkRequired HostRootModified PrivilegedContainerRequired StateRootPathExposed RawCommandExposed RawExecutableExposed BackendDetailsExposed validateNoBackendTerms].each do |token|
  assert(go_runtime_production_receipt_notification_action_dry_run_result_persistence_authorization_audit_source.include?(token), "Go Runtime production receipt notification action dry-run result persistence authorization audit must expose safety gate #{token}")
end
%w[production_receipt_notification_action_dry_run_result_visibility_audit.go docs/claude-code-current-dispatch-picks.md].each do |token|
  assert(go_runtime_production_receipt_notification_action_dry_run_result_persistence_authorization_audit_source.include?(token), "Go Runtime production receipt notification action dry-run result persistence authorization audit must consume #{token}")
end
go_runtime_production_receipt_notification_action_dry_run_result_persistence_authorization_audit_test_source = read_project_file("internal/runtime/owner/production_receipt_notification_action_dry_run_result_persistence_authorization_audit_test.go") +
                                                                                                               read_project_file("cmd/xnix-runtime-go/production_dbus_gate_cli_test.go") +
                                                                                                               read_project_file("cmd/xnix-runtime-go/production_dbus_gate_commands.go") +
                                                                                                               read_project_file("cmd/xnix-runtime-go/main.go")
%w[TestProductionReceiptNotificationActionDryRunResultPersistenceAuthorizationAuditPreviewModelsAuthorization TestProductionReceiptNotificationActionDryRunResultPersistenceAuthorizationAuditPreviewFailsClosedWithoutSources TestProductionReceiptNotificationActionDryRunResultPersistenceAuthorizationAuditPreviewCommand production-receipt-notification-action-dry-run-result-persistence-authorization-audit-preview persistence_authorization_audit_required persistence_authorization_audit_modeled result_visibility_audit_consumed persistence_authorization_guidance_consumed persistence_authorization_ready result_persistence_authorized result_visibility_persisted ready_authorization_item_count granted_persistence_item_count persisted_result_item_count].each do |token|
  assert(go_runtime_production_receipt_notification_action_dry_run_result_persistence_authorization_audit_test_source.include?(token), "Go Runtime production receipt notification action dry-run result persistence authorization audit tests and CLI must include #{token}")
end

go_runtime_production_receipt_notification_action_dry_run_result_retention_redaction_policy_audit_source = read_project_file("internal/runtime/owner/production_receipt_notification_action_dry_run_result_retention_redaction_policy_audit.go")
%w[ProductionReceiptNotificationActionDryRunResultRetentionRedactionPolicyAuditPreview ProductionReceiptNotificationActionDryRunResultRetentionRedactionPolicyAuditItem ProductionReceiptNotificationActionDryRunResultRetentionRedactionPolicyAuditCheck NewProductionReceiptNotificationActionDryRunResultRetentionRedactionPolicyAuditPreview xnix.runtime.production_receipt_notification_action_dry_run_result_retention_redaction_policy_audit.v1 production-receipt-notification-action-dry-run-result-retention-redaction-policy-audit-preview receipt-notification-action-dry-run-result-retention-redaction-policy-audit production-receipt-notification-action-dry-run-result-retention-redaction-policy-audit-ready-policy-only production-receipt-notification-action-dry-run-result-retention-redaction-policy-audit-blocked production-receipt-notification-action-dry-run-result-persistence-authorization-audit-preview persistence-authorization-audit-consumed retention-redaction-guidance-consumed retention-redaction-policy-modeled-only five-retention-redaction-policy-items-present policy-items-ready-enforcement-disabled enforcement-persistence-and-execution-disabled request-creation-dispatch-actions-and-portal-disabled receipt-notification-navigation-and-support-writes-disabled production-and-host-boundary-closed review-receipt-dry-run-result-retention-redaction-policy renew-receipt-dry-run-result-retention-redaction-policy open-compatibility-center-dry-run-result-retention-redaction-policy dismiss-receipt-dry-run-result-retention-redaction-policy support-info-dry-run-result-retention-redaction-policy].each do |token|
  assert(go_runtime_production_receipt_notification_action_dry_run_result_retention_redaction_policy_audit_source.include?(token), "Go Runtime production receipt notification action dry-run result retention redaction policy audit must include #{token}")
end
%w[RetentionRedactionPolicyAuditRequired RetentionRedactionPolicyAuditModeled PersistenceAuthorizationAuditConsumed RetentionRedactionGuidanceConsumed RetentionRedactionPolicyReady RetentionWindowModeled RedactionRulesModeled ResultPersistenceAuthorized RetentionEnforcementEnabled RedactionEnforcementEnabled DispatchDryRunExecuted DryRunResultPersisted ResultVisibilityPersisted ProductionReadiness ProductionOwnershipReady PolicyItemCount RequiredPolicyItemCount ReadyPolicyItemCount MissingPolicyItemCount RetentionEnabledItemCount RedactionWriteItemCount PersistedResultItemCount ExecutedDryRunItemCount CreatedRequestObjectCount PortalRequestCreatedCount SideEffectPolicyItemCount SystemServiceStarted SessionBusClaimed ProductionBusClaimed ProductionOwnerEnabled ProductionActivationReady WriteMethodsEnabled RuntimeWritesEnabled RequestObjectCreationEnabled RequestObjectDispatchEnabled RequestObjectPersistenceEnabled DispatchDryRunExecutionEnabled DryRunResultPersistenceEnabled ResultVisibilityPersistenceEnabled RetentionPolicyPersistenceEnabled RuntimeDiagnosticsPersisted PortalRequestCreated RequestObjectsCreated RequestObjectsDispatched NotificationSent NotificationDeliveryEnabled NotificationActionEnabled CompatibilityCenterOpened CompatibilityCenterPersisted ReceiptWriterEnabled ReceiptPersistenceEnabled DesktopFilesWritten SupportBundleExported SnapshotRestoreExecuted StateCleanupExecuted FileContentRead FilePathsExposed NetworkRequired HostRootModified PrivilegedContainerRequired StateRootPathExposed RawCommandExposed RawExecutableExposed BackendDetailsExposed validateNoBackendTerms].each do |token|
  assert(go_runtime_production_receipt_notification_action_dry_run_result_retention_redaction_policy_audit_source.include?(token), "Go Runtime production receipt notification action dry-run result retention redaction policy audit must expose safety gate #{token}")
end
%w[production_receipt_notification_action_dry_run_result_persistence_authorization_audit.go docs/claude-code-current-dispatch-picks.md].each do |token|
  assert(go_runtime_production_receipt_notification_action_dry_run_result_retention_redaction_policy_audit_source.include?(token), "Go Runtime production receipt notification action dry-run result retention redaction policy audit must consume #{token}")
end
go_runtime_production_receipt_notification_action_dry_run_result_retention_redaction_policy_audit_test_source = read_project_file("internal/runtime/owner/production_receipt_notification_action_dry_run_result_retention_redaction_policy_audit_test.go") +
                                                                                                               read_project_file("cmd/xnix-runtime-go/production_dbus_gate_cli_test.go") +
                                                                                                               read_project_file("cmd/xnix-runtime-go/production_dbus_gate_commands.go") +
                                                                                                               read_project_file("cmd/xnix-runtime-go/main.go")
%w[TestProductionReceiptNotificationActionDryRunResultRetentionRedactionPolicyAuditPreviewModelsPolicy TestProductionReceiptNotificationActionDryRunResultRetentionRedactionPolicyAuditPreviewFailsClosedWithoutSources TestProductionReceiptNotificationActionDryRunResultRetentionRedactionPolicyAuditPreviewCommand production-receipt-notification-action-dry-run-result-retention-redaction-policy-audit-preview retention_redaction_policy_audit_required retention_redaction_policy_audit_modeled persistence_authorization_audit_consumed retention_redaction_guidance_consumed retention_redaction_policy_ready retention_window_modeled redaction_rules_modeled ready_policy_item_count retention_enabled_item_count redaction_write_item_count retention_policy_persistence_enabled].each do |token|
  assert(go_runtime_production_receipt_notification_action_dry_run_result_retention_redaction_policy_audit_test_source.include?(token), "Go Runtime production receipt notification action dry-run result retention redaction policy audit tests and CLI must include #{token}")
end

go_runtime_production_receipt_notification_action_dry_run_result_opaque_lookup_audit_source = read_project_file("internal/runtime/owner/production_receipt_notification_action_dry_run_result_opaque_lookup_audit.go")
%w[ProductionReceiptNotificationActionDryRunResultOpaqueLookupAuditPreview ProductionReceiptNotificationActionDryRunResultOpaqueLookupAuditItem ProductionReceiptNotificationActionDryRunResultOpaqueLookupAuditCheck NewProductionReceiptNotificationActionDryRunResultOpaqueLookupAuditPreview xnix.runtime.production_receipt_notification_action_dry_run_result_opaque_lookup_audit.v1 production-receipt-notification-action-dry-run-result-opaque-lookup-audit-preview receipt-notification-action-dry-run-result-opaque-lookup-audit production-receipt-notification-action-dry-run-result-opaque-lookup-audit-ready-lookup-disabled production-receipt-notification-action-dry-run-result-opaque-lookup-audit-blocked production-receipt-notification-action-dry-run-result-retention-redaction-policy-audit-preview retention-redaction-policy-audit-consumed opaque-lookup-guidance-consumed opaque-lookup-modeled-only five-opaque-lookup-items-present opaque-lookup-items-ready-lookup-disabled lookup-persistence-result-persistence-and-execution-disabled request-creation-dispatch-actions-and-portal-disabled receipt-notification-navigation-and-support-writes-disabled production-and-host-boundary-closed review-receipt-dry-run-result-opaque-lookup renew-receipt-dry-run-result-opaque-lookup open-compatibility-center-dry-run-result-opaque-lookup dismiss-receipt-dry-run-result-opaque-lookup support-info-dry-run-result-opaque-lookup].each do |token|
  assert(go_runtime_production_receipt_notification_action_dry_run_result_opaque_lookup_audit_source.include?(token), "Go Runtime production receipt notification action dry-run result opaque lookup audit must include #{token}")
end
%w[OpaqueLookupAuditRequired OpaqueLookupAuditModeled RetentionRedactionPolicyAuditConsumed OpaqueLookupGuidanceConsumed OpaqueLookupBoundaryReady OwnerManagedOpaqueLookupModeled OpaqueLookupEnabled OpaqueLookupPersisted OpaqueResultIDSupported CallerStateRootRequired StateRootPathExposed FilePathsExposed FileContentRead ResultPersistenceAuthorized RetentionEnforcementEnabled RedactionEnforcementEnabled DispatchDryRunExecuted DryRunResultPersisted ResultVisibilityPersisted RuntimeDiagnosticsPersisted ProductionReadiness ProductionOwnershipReady LookupItemCount RequiredLookupItemCount ReadyLookupItemCount MissingLookupItemCount EnabledLookupItemCount PersistedLookupItemCount PathExposedLookupItemCount PersistedResultItemCount ExecutedDryRunItemCount CreatedRequestObjectCount PortalRequestCreatedCount SideEffectLookupItemCount SystemServiceStarted SessionBusClaimed ProductionBusClaimed ProductionOwnerEnabled ProductionActivationReady WriteMethodsEnabled RuntimeWritesEnabled RequestObjectCreationEnabled RequestObjectDispatchEnabled RequestObjectPersistenceEnabled DispatchDryRunExecutionEnabled DryRunResultPersistenceEnabled ResultVisibilityPersistenceEnabled RetentionPolicyPersistenceEnabled RuntimeDiagnosticsPersisted PortalRequestCreated RequestObjectsCreated RequestObjectsDispatched NotificationSent NotificationDeliveryEnabled NotificationActionEnabled CompatibilityCenterOpened CompatibilityCenterPersisted ReceiptWriterEnabled ReceiptPersistenceEnabled DesktopFilesWritten SupportBundleExported SnapshotRestoreExecuted StateCleanupExecuted FileContentRead FilePathsExposed NetworkRequired HostRootModified PrivilegedContainerRequired StateRootPathExposed RawCommandExposed RawExecutableExposed BackendDetailsExposed validateNoBackendTerms].each do |token|
  assert(go_runtime_production_receipt_notification_action_dry_run_result_opaque_lookup_audit_source.include?(token), "Go Runtime production receipt notification action dry-run result opaque lookup audit must expose safety gate #{token}")
end
%w[production_receipt_notification_action_dry_run_result_retention_redaction_policy_audit.go docs/claude-code-current-dispatch-picks.md].each do |token|
  assert(go_runtime_production_receipt_notification_action_dry_run_result_opaque_lookup_audit_source.include?(token), "Go Runtime production receipt notification action dry-run result opaque lookup audit must consume #{token}")
end
go_runtime_production_receipt_notification_action_dry_run_result_opaque_lookup_audit_test_source = read_project_file("internal/runtime/owner/production_receipt_notification_action_dry_run_result_opaque_lookup_audit_test.go") +
                                                                                                  read_project_file("cmd/xnix-runtime-go/production_dbus_gate_cli_test.go") +
                                                                                                  read_project_file("cmd/xnix-runtime-go/production_dbus_gate_commands.go") +
                                                                                                  read_project_file("cmd/xnix-runtime-go/main.go")
%w[TestProductionReceiptNotificationActionDryRunResultOpaqueLookupAuditPreviewModelsLookup TestProductionReceiptNotificationActionDryRunResultOpaqueLookupAuditPreviewFailsClosedWithoutSources TestProductionReceiptNotificationActionDryRunResultOpaqueLookupAuditPreviewCommand production-receipt-notification-action-dry-run-result-opaque-lookup-audit-preview opaque_lookup_audit_required opaque_lookup_audit_modeled retention_redaction_policy_audit_consumed opaque_lookup_guidance_consumed opaque_lookup_boundary_ready owner_managed_opaque_lookup_modeled opaque_lookup_enabled opaque_lookup_persisted ready_lookup_item_count enabled_lookup_item_count path_exposed_lookup_item_count state_root_path_exposed file_paths_exposed file_content_read].each do |token|
  assert(go_runtime_production_receipt_notification_action_dry_run_result_opaque_lookup_audit_test_source.include?(token), "Go Runtime production receipt notification action dry-run result opaque lookup audit tests and CLI must include #{token}")
end

go_runtime_production_receipt_notification_action_dry_run_result_lookup_route_authorization_audit_source = read_project_file("internal/runtime/owner/production_receipt_notification_action_dry_run_result_lookup_route_authorization_audit.go")
%w[ProductionReceiptNotificationActionDryRunResultLookupRouteAuthorizationAuditPreview ProductionReceiptNotificationActionDryRunResultLookupRouteAuthorizationAuditItem ProductionReceiptNotificationActionDryRunResultLookupRouteAuthorizationAuditCheck NewProductionReceiptNotificationActionDryRunResultLookupRouteAuthorizationAuditPreview xnix.runtime.production_receipt_notification_action_dry_run_result_lookup_route_authorization_audit.v1 production-receipt-notification-action-dry-run-result-lookup-route-authorization-audit-preview receipt-notification-action-dry-run-result-lookup-route-authorization-audit production-receipt-notification-action-dry-run-result-lookup-route-authorization-audit-ready-route-disabled production-receipt-notification-action-dry-run-result-lookup-route-authorization-audit-blocked production-receipt-notification-action-dry-run-result-opaque-lookup-audit-preview opaque-lookup-audit-consumed lookup-route-guidance-consumed lookup-route-authorization-modeled-only five-lookup-route-authorization-items-present route-authorization-items-ready-route-disabled route-lookup-persistence-result-persistence-and-execution-disabled request-creation-dispatch-actions-and-portal-disabled receipt-notification-navigation-and-support-writes-disabled production-and-host-boundary-closed review-receipt-dry-run-result-lookup-route-authorization renew-receipt-dry-run-result-lookup-route-authorization open-compatibility-center-dry-run-result-lookup-route-authorization dismiss-receipt-dry-run-result-lookup-route-authorization support-info-dry-run-result-lookup-route-authorization].each do |token|
  assert(go_runtime_production_receipt_notification_action_dry_run_result_lookup_route_authorization_audit_source.include?(token), "Go Runtime production receipt notification action dry-run result lookup route authorization audit must include #{token}")
end
%w[LookupRouteAuthorizationAuditRequired LookupRouteAuthorizationAuditModeled OpaqueLookupAuditConsumed LookupRouteGuidanceConsumed LookupRouteAuthorizationReady OwnerLocalLookupRouteModeled KDEConsumptionAuthorized RuntimeConsumptionAuthorized LookupRouteAuthorized LookupRouteEnabled LookupRoutePersisted OpaqueLookupEnabled OpaqueLookupPersisted OpaqueResultIDSupported CallerStateRootRequired StateRootPathExposed FilePathsExposed FileContentRead ResultPersistenceAuthorized DispatchDryRunExecuted DryRunResultPersisted ResultVisibilityPersisted RuntimeDiagnosticsPersisted ProductionReadiness ProductionOwnershipReady RouteItemCount RequiredRouteItemCount ReadyRouteItemCount MissingRouteItemCount AuthorizedRouteItemCount EnabledRouteItemCount PersistedRouteItemCount LookupEnabledItemCount PersistedLookupItemCount PathExposedRouteItemCount PersistedResultItemCount ExecutedDryRunItemCount CreatedRequestObjectCount PortalRequestCreatedCount SideEffectRouteItemCount SystemServiceStarted SessionBusClaimed ProductionBusClaimed ProductionOwnerEnabled ProductionActivationReady WriteMethodsEnabled RuntimeWritesEnabled RequestObjectCreationEnabled RequestObjectDispatchEnabled RequestObjectPersistenceEnabled DispatchDryRunExecutionEnabled DryRunResultPersistenceEnabled ResultVisibilityPersistenceEnabled RetentionPolicyPersistenceEnabled RuntimeDiagnosticsPersisted PortalRequestCreated RequestObjectsCreated RequestObjectsDispatched NotificationSent NotificationDeliveryEnabled NotificationActionEnabled CompatibilityCenterOpened CompatibilityCenterPersisted ReceiptWriterEnabled ReceiptPersistenceEnabled DesktopFilesWritten SupportBundleExported SnapshotRestoreExecuted StateCleanupExecuted FileContentRead FilePathsExposed NetworkRequired HostRootModified PrivilegedContainerRequired StateRootPathExposed RawCommandExposed RawExecutableExposed BackendDetailsExposed validateNoBackendTerms].each do |token|
  assert(go_runtime_production_receipt_notification_action_dry_run_result_lookup_route_authorization_audit_source.include?(token), "Go Runtime production receipt notification action dry-run result lookup route authorization audit must expose safety gate #{token}")
end
%w[production_receipt_notification_action_dry_run_result_opaque_lookup_audit.go docs/claude-code-current-dispatch-picks.md].each do |token|
  assert(go_runtime_production_receipt_notification_action_dry_run_result_lookup_route_authorization_audit_source.include?(token), "Go Runtime production receipt notification action dry-run result lookup route authorization audit must consume #{token}")
end
go_runtime_production_receipt_notification_action_dry_run_result_lookup_route_authorization_audit_test_source = read_project_file("internal/runtime/owner/production_receipt_notification_action_dry_run_result_lookup_route_authorization_audit_test.go") +
                                                                                                               read_project_file("cmd/xnix-runtime-go/production_dbus_gate_cli_test.go") +
                                                                                                               read_project_file("cmd/xnix-runtime-go/production_dbus_gate_commands.go") +
                                                                                                               read_project_file("cmd/xnix-runtime-go/main.go")
%w[TestProductionReceiptNotificationActionDryRunResultLookupRouteAuthorizationAuditPreviewModelsRoutes TestProductionReceiptNotificationActionDryRunResultLookupRouteAuthorizationAuditPreviewFailsClosedWithoutSources TestProductionReceiptNotificationActionDryRunResultLookupRouteAuthorizationAuditPreviewCommand production-receipt-notification-action-dry-run-result-lookup-route-authorization-audit-preview lookup_route_authorization_audit_required lookup_route_authorization_audit_modeled opaque_lookup_audit_consumed lookup_route_guidance_consumed lookup_route_authorization_ready owner_local_lookup_route_modeled lookup_route_authorized lookup_route_enabled lookup_route_persisted ready_route_item_count authorized_route_item_count enabled_route_item_count path_exposed_route_item_count state_root_path_exposed file_paths_exposed file_content_read].each do |token|
  assert(go_runtime_production_receipt_notification_action_dry_run_result_lookup_route_authorization_audit_test_source.include?(token), "Go Runtime production receipt notification action dry-run result lookup route authorization audit tests and CLI must include #{token}")
end

go_runtime_production_receipt_notification_action_dry_run_result_lookup_consumer_redaction_audit_source = read_project_file("internal/runtime/owner/production_receipt_notification_action_dry_run_result_lookup_consumer_redaction_audit.go")
%w[ProductionReceiptNotificationActionDryRunResultLookupConsumerRedactionAuditPreview ProductionReceiptNotificationActionDryRunResultLookupConsumerRedactionAuditItem ProductionReceiptNotificationActionDryRunResultLookupConsumerRedactionAuditCheck NewProductionReceiptNotificationActionDryRunResultLookupConsumerRedactionAuditPreview xnix.runtime.production_receipt_notification_action_dry_run_result_lookup_consumer_redaction_audit.v1 production-receipt-notification-action-dry-run-result-lookup-consumer-redaction-audit-preview receipt-notification-action-dry-run-result-lookup-consumer-redaction-audit production-receipt-notification-action-dry-run-result-lookup-consumer-redaction-audit-ready-consumers-disabled production-receipt-notification-action-dry-run-result-lookup-consumer-redaction-audit-blocked production-receipt-notification-action-dry-run-result-lookup-route-authorization-audit-preview lookup-route-authorization-audit-consumed consumer-redaction-guidance-consumed consumer-redaction-modeled-only five-consumer-redaction-items-present consumer-redaction-items-ready-consumers-disabled lookup-persistence-result-persistence-and-execution-disabled raw-result-and-path-exposure-disabled request-dispatch-notification-and-support-disabled production-and-host-boundary-closed review-receipt-dry-run-result-lookup-consumer-redaction renew-receipt-dry-run-result-lookup-consumer-redaction open-compatibility-center-dry-run-result-lookup-consumer-redaction dismiss-receipt-dry-run-result-lookup-consumer-redaction support-info-dry-run-result-lookup-consumer-redaction].each do |token|
  assert(go_runtime_production_receipt_notification_action_dry_run_result_lookup_consumer_redaction_audit_source.include?(token), "Go Runtime production receipt notification action dry-run result lookup consumer redaction audit must include #{token}")
end
%w[ConsumerRedactionAuditRequired ConsumerRedactionAuditModeled LookupRouteAuthorizationAuditConsumed ConsumerRedactionGuidanceConsumed ConsumerRedactionReady KDEConsumerRedactionModeled RuntimeConsumerRedactionModeled OpaqueResultIDSupported ConsumerConsumptionAuthorized KDEConsumerEnabled RuntimeConsumerEnabled LookupRouteAuthorized LookupRouteEnabled LookupRoutePersisted OpaqueLookupEnabled OpaqueLookupPersisted RedactedSummaryPersisted RawResultExposed RuntimeDiagnosticsPersisted DryRunResultPersisted DispatchDryRunExecuted ProductionReadiness ProductionOwnershipReady ConsumerItemCount RequiredConsumerItemCount ReadyConsumerItemCount MissingConsumerItemCount RedactedConsumerItemCount EnabledConsumerItemCount RawExposedConsumerItemCount PersistedConsumerItemCount SideEffectConsumerItemCount RequestObjectCreationEnabled RequestObjectDispatchEnabled PortalRequestCreated RequestObjectsCreated RequestObjectsDispatched NotificationSent NotificationActionEnabled CompatibilityCenterOpened ReceiptWriterEnabled ReceiptPersistenceEnabled SystemServiceStarted ProductionBusClaimed WriteMethodsEnabled RuntimeWritesEnabled DesktopFilesWritten SupportBundleExported SnapshotRestoreExecuted StateCleanupExecuted FileContentRead FilePathsExposed NetworkRequired HostRootModified PrivilegedContainerRequired StateRootPathExposed RawCommandExposed RawExecutableExposed BackendDetailsExposed validateNoBackendTerms].each do |token|
  assert(go_runtime_production_receipt_notification_action_dry_run_result_lookup_consumer_redaction_audit_source.include?(token), "Go Runtime production receipt notification action dry-run result lookup consumer redaction audit must expose safety gate #{token}")
end
%w[production_receipt_notification_action_dry_run_result_lookup_route_authorization_audit.go docs/claude-code-current-dispatch-picks.md].each do |token|
  assert(go_runtime_production_receipt_notification_action_dry_run_result_lookup_consumer_redaction_audit_source.include?(token), "Go Runtime production receipt notification action dry-run result lookup consumer redaction audit must consume #{token}")
end
go_runtime_production_receipt_notification_action_dry_run_result_lookup_consumer_redaction_audit_test_source = read_project_file("internal/runtime/owner/production_receipt_notification_action_dry_run_result_lookup_consumer_redaction_audit_test.go") +
                                                                                                             read_project_file("cmd/xnix-runtime-go/production_dbus_gate_cli_test.go") +
                                                                                                             read_project_file("cmd/xnix-runtime-go/production_dbus_gate_commands.go") +
                                                                                                             read_project_file("cmd/xnix-runtime-go/main.go")
%w[TestProductionReceiptNotificationActionDryRunResultLookupConsumerRedactionAuditPreviewModelsConsumers TestProductionReceiptNotificationActionDryRunResultLookupConsumerRedactionAuditPreviewFailsClosedWithoutSources TestProductionReceiptNotificationActionDryRunResultLookupConsumerRedactionAuditPreviewCommand production-receipt-notification-action-dry-run-result-lookup-consumer-redaction-audit-preview consumer_redaction_audit_required consumer_redaction_audit_modeled lookup_route_authorization_audit_consumed consumer_redaction_guidance_consumed consumer_redaction_ready kde_consumer_redaction_modeled runtime_consumer_redaction_modeled consumer_consumption_authorized kde_consumer_enabled runtime_consumer_enabled redacted_summary_persisted raw_result_exposed consumer_item_count ready_consumer_item_count redacted_consumer_item_count enabled_consumer_item_count raw_exposed_consumer_item_count state_root_path_exposed file_paths_exposed file_content_read].each do |token|
  assert(go_runtime_production_receipt_notification_action_dry_run_result_lookup_consumer_redaction_audit_test_source.include?(token), "Go Runtime production receipt notification action dry-run result lookup consumer redaction audit tests and CLI must include #{token}")
end

go_runtime_production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_gate_audit_source = read_project_file("internal/runtime/owner/production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_gate_audit.go")
%w[ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementGateAuditPreview ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementGateAuditItem ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementGateAuditCheck NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementGateAuditPreview xnix.runtime.production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_gate_audit.v1 production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-gate-audit-preview receipt-notification-action-dry-run-result-lookup-consumer-enablement-gate-audit production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-gate-audit-ready-consumers-disabled production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-gate-audit-blocked production-receipt-notification-action-dry-run-result-lookup-route-authorization-audit-preview production-receipt-notification-action-dry-run-result-lookup-consumer-redaction-audit-preview consumer-enablement-gate-audit-required lookup-route-authorization-audit-consumed consumer-redaction-audit-consumed consumer-enablement-guidance-consumed consumer-enablement-gate-modeled-only five-consumer-enablement-gate-items-present lookup-result-persistence-and-execution-disabled request-notification-navigation-and-support-disabled production-and-host-boundary-closed review-receipt-dry-run-result-lookup-consumer-enablement-gate renew-receipt-dry-run-result-lookup-consumer-enablement-gate open-compatibility-center-dry-run-result-lookup-consumer-enablement-gate dismiss-receipt-dry-run-result-lookup-consumer-enablement-gate support-info-dry-run-result-lookup-consumer-enablement-gate].each do |token|
  assert(go_runtime_production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_gate_audit_source.include?(token), "Go Runtime production receipt notification action dry-run result lookup consumer enablement gate audit must include #{token}")
end
%w[ConsumerEnablementGateRequired ConsumerEnablementGateModeled LookupRouteAuthorizationAuditConsumed ConsumerRedactionAuditConsumed ConsumerEnablementGuidanceConsumed ConsumerEnablementGateReady RouteAuthorizationPrerequisiteModeled ConsumerRedactionPrerequisiteModeled OpaqueResultIDSupported ConsumerConsumptionAuthorized ConsumerEnablementAuthorized KDEConsumerEnabled RuntimeConsumerEnabled LookupRouteAuthorized LookupRouteEnabled LookupRoutePersisted OpaqueLookupEnabled OpaqueLookupPersisted RedactedSummaryPersisted RawResultExposed RuntimeDiagnosticsPersisted DryRunResultPersisted DispatchDryRunExecuted ProductionReadiness ProductionOwnershipReady GateItemCount RequiredGateItemCount ReadyGateItemCount MissingGateItemCount AuthorizedGateItemCount EnabledGateItemCount RedactionReadyGateItemCount RouteReadyGateItemCount PersistedGateItemCount RawExposedGateItemCount SideEffectGateItemCount RequestObjectCreationEnabled RequestObjectDispatchEnabled PortalRequestCreated RequestObjectsCreated RequestObjectsDispatched NotificationSent NotificationActionEnabled CompatibilityCenterOpened ReceiptWriterEnabled ReceiptPersistenceEnabled SystemServiceStarted ProductionBusClaimed WriteMethodsEnabled RuntimeWritesEnabled DesktopFilesWritten SupportBundleExported SnapshotRestoreExecuted StateCleanupExecuted FileContentRead FilePathsExposed NetworkRequired HostRootModified PrivilegedContainerRequired StateRootPathExposed RawCommandExposed RawExecutableExposed BackendDetailsExposed validateNoBackendTerms].each do |token|
  assert(go_runtime_production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_gate_audit_source.include?(token), "Go Runtime production receipt notification action dry-run result lookup consumer enablement gate audit must expose safety gate #{token}")
end
%w[production_receipt_notification_action_dry_run_result_lookup_route_authorization_audit.go production_receipt_notification_action_dry_run_result_lookup_consumer_redaction_audit.go docs/claude-code-current-dispatch-picks.md].each do |token|
  assert(go_runtime_production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_gate_audit_source.include?(token), "Go Runtime production receipt notification action dry-run result lookup consumer enablement gate audit must consume #{token}")
end
go_runtime_production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_gate_audit_test_source = read_project_file("internal/runtime/owner/production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_gate_audit_test.go") +
                                                                                                                    read_project_file("cmd/xnix-runtime-go/production_dbus_gate_cli_test.go") +
                                                                                                                    read_project_file("cmd/xnix-runtime-go/production_dbus_gate_commands.go") +
                                                                                                                    read_project_file("cmd/xnix-runtime-go/main.go")
%w[TestProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementGateAuditPreviewModelsGate TestProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementGateAuditPreviewFailsClosedWithoutSources TestProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementGateAuditPreviewCommand production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-gate-audit-preview consumer_enablement_gate_required consumer_enablement_gate_modeled consumer_redaction_audit_consumed consumer_enablement_guidance_consumed consumer_enablement_gate_ready route_authorization_prerequisite_modeled consumer_redaction_prerequisite_modeled consumer_enablement_authorized gate_item_count ready_gate_item_count authorized_gate_item_count enabled_gate_item_count raw_exposed_gate_item_count state_root_path_exposed file_paths_exposed file_content_read].each do |token|
  assert(go_runtime_production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_gate_audit_test_source.include?(token), "Go Runtime production receipt notification action dry-run result lookup consumer enablement gate audit tests and CLI must include #{token}")
end

go_runtime_production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_authorization_receipt_audit_source = read_project_file("internal/runtime/owner/production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_authorization_receipt_audit.go")
%w[ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementAuthorizationReceiptAuditPreview ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementAuthorizationReceiptAuditItem ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementAuthorizationReceiptAuditCheck NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementAuthorizationReceiptAuditPreview xnix.runtime.production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_authorization_receipt_audit.v1 production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-authorization-receipt-audit-preview receipt-notification-action-dry-run-result-lookup-consumer-enablement-authorization-receipt-audit production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-authorization-receipt-audit-ready-receipt-disabled production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-authorization-receipt-audit-blocked production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-gate-audit-preview consumer-enablement-gate-audit-consumed authorization-receipt-guidance-consumed authorization-receipt-boundary-modeled-only five-authorization-receipt-items-present authorization-receipt-items-ready-receipt-disabled receipt-acceptance-consumer-and-lookup-disabled request-notification-navigation-and-support-disabled raw-result-and-path-exposure-disabled production-and-host-boundary-closed review-receipt-dry-run-result-lookup-consumer-enablement-authorization-receipt renew-receipt-dry-run-result-lookup-consumer-enablement-authorization-receipt open-compatibility-center-dry-run-result-lookup-consumer-enablement-authorization-receipt dismiss-receipt-dry-run-result-lookup-consumer-enablement-authorization-receipt support-info-dry-run-result-lookup-consumer-enablement-authorization-receipt].each do |token|
  assert(go_runtime_production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_authorization_receipt_audit_source.include?(token), "Go Runtime production receipt notification action dry-run result lookup consumer enablement authorization receipt audit must include #{token}")
end
%w[AuthorizationReceiptAuditRequired AuthorizationReceiptAuditModeled ConsumerEnablementGateAuditConsumed AuthorizationReceiptGuidanceConsumed AuthorizationReceiptBoundaryReady ConsumerEnablementGateReady OpaqueAuthorizationReceiptModeled ConsumerAuthorizationReady ReceiptPresent ReceiptAccepted AuthorizationAccepted ConsumerConsumptionAuthorized ConsumerEnablementAuthorized KDEConsumerEnabled RuntimeConsumerEnabled LookupRouteAuthorized LookupRouteEnabled LookupRoutePersisted OpaqueLookupEnabled OpaqueLookupPersisted RedactedSummaryPersisted RawResultExposed RuntimeDiagnosticsPersisted DryRunResultPersisted DispatchDryRunExecuted ProductionReadiness ProductionOwnershipReady ReceiptItemCount RequiredReceiptItemCount ReadyReceiptItemCount MissingReceiptItemCount AcceptedReceiptItemCount AuthorizedConsumerItemCount EnabledConsumerItemCount PersistedReceiptItemCount RawExposedReceiptItemCount SideEffectReceiptItemCount RequestObjectCreationEnabled RequestObjectDispatchEnabled PortalRequestCreated RequestObjectsCreated RequestObjectsDispatched NotificationSent NotificationActionEnabled CompatibilityCenterOpened ReceiptWriterEnabled ReceiptPersistenceEnabled SystemServiceStarted ProductionBusClaimed WriteMethodsEnabled RuntimeWritesEnabled DesktopFilesWritten SupportBundleExported SnapshotRestoreExecuted StateCleanupExecuted FileContentRead FilePathsExposed NetworkRequired HostRootModified PrivilegedContainerRequired StateRootPathExposed RawCommandExposed RawExecutableExposed BackendDetailsExposed validateNoBackendTerms].each do |token|
  assert(go_runtime_production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_authorization_receipt_audit_source.include?(token), "Go Runtime production receipt notification action dry-run result lookup consumer enablement authorization receipt audit must expose safety gate #{token}")
end
%w[production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_gate_audit.go docs/claude-code-current-dispatch-picks.md].each do |token|
  assert(go_runtime_production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_authorization_receipt_audit_source.include?(token), "Go Runtime production receipt notification action dry-run result lookup consumer enablement authorization receipt audit must consume #{token}")
end
go_runtime_production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_authorization_receipt_audit_test_source = read_project_file("internal/runtime/owner/production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_authorization_receipt_audit_test.go") +
                                                                                                                              read_project_file("cmd/xnix-runtime-go/production_dbus_gate_cli_test.go") +
                                                                                                                              read_project_file("cmd/xnix-runtime-go/production_dbus_gate_commands.go") +
                                                                                                                              read_project_file("cmd/xnix-runtime-go/main.go")
%w[TestProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementAuthorizationReceiptAuditPreviewModelsReceipt TestProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementAuthorizationReceiptAuditPreviewFailsClosedWithoutSources TestProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementAuthorizationReceiptAuditPreviewCommand production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-authorization-receipt-audit-preview authorization_receipt_audit_required authorization_receipt_audit_modeled consumer_enablement_gate_audit_consumed authorization_receipt_guidance_consumed authorization_receipt_boundary_ready consumer_enablement_gate_ready opaque_authorization_receipt_modeled consumer_authorization_ready receipt_present receipt_accepted authorization_accepted consumer_enablement_authorized receipt_item_count ready_receipt_item_count accepted_receipt_item_count authorized_consumer_item_count enabled_consumer_item_count raw_exposed_receipt_item_count state_root_path_exposed file_paths_exposed file_content_read].each do |token|
  assert(go_runtime_production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_authorization_receipt_audit_test_source.include?(token), "Go Runtime production receipt notification action dry-run result lookup consumer enablement authorization receipt audit tests and CLI must include #{token}")
end

go_runtime_production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_receipt_writer_authorization_audit_source = read_project_file("internal/runtime/owner/production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_receipt_writer_authorization_audit.go")
%w[ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptWriterAuthorizationAuditPreview ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptWriterAuthorizationAuditItem ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptWriterAuthorizationAuditCheck NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptWriterAuthorizationAuditPreview xnix.runtime.production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_receipt_writer_authorization_audit.v1 production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-receipt-writer-authorization-audit-preview receipt-notification-action-dry-run-result-lookup-consumer-enablement-receipt-writer-authorization-audit production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-receipt-writer-authorization-audit-ready-writes-disabled production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-receipt-writer-authorization-audit-blocked production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-authorization-receipt-audit-preview production-receipt-writer-authorization-review-preview authorization-receipt-audit-consumed base-writer-authorization-review-consumed writer-authorization-guidance-consumed writer-authorization-boundary-modeled-only five-writer-authorization-items-present writer-authorization-items-ready-writes-disabled receipt-writes-persistence-and-replay-disabled consumer-lookup-request-and-support-disabled production-and-host-boundary-closed review-receipt-dry-run-result-lookup-consumer-enablement-writer-authorization renew-receipt-dry-run-result-lookup-consumer-enablement-writer-authorization open-compatibility-center-dry-run-result-lookup-consumer-enablement-writer-authorization dismiss-receipt-dry-run-result-lookup-consumer-enablement-writer-authorization support-info-dry-run-result-lookup-consumer-enablement-writer-authorization].each do |token|
  assert(go_runtime_production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_receipt_writer_authorization_audit_source.include?(token), "Go Runtime production receipt notification action dry-run result lookup consumer enablement receipt writer authorization audit must include #{token}")
end
%w[WriterAuthorizationAuditRequired WriterAuthorizationAuditModeled AuthorizationReceiptAuditConsumed BaseWriterAuthorizationReviewConsumed WriterAuthorizationGuidanceConsumed WriterAuthorizationBoundaryReady AuthorizationReceiptBoundaryReady ConsumerEnablementReceiptWriterReady ReceiptWriterAuthorizationReady ReceiptPresent ReceiptAccepted AuthorizationAccepted ReceiptWriterEnabled ReceiptPersistenceEnabled ReceiptLookupWritesEnabled ReceiptReplayEnabled ConsumerConsumptionAuthorized ConsumerEnablementAuthorized KDEConsumerEnabled RuntimeConsumerEnabled LookupRouteAuthorized LookupRouteEnabled OpaqueLookupEnabled RedactedSummaryPersisted RawResultExposed RuntimeDiagnosticsPersisted DryRunResultPersisted DispatchDryRunExecuted AuthorizationItemCount RequiredAuthorizationItemCount ReadyAuthorizationItemCount MissingAuthorizationItemCount WriteEnabledAuthorizationItemCount PersistedAuthorizationItemCount AcceptedAuthorizationItemCount EnabledConsumerItemCount RawExposedAuthorizationItemCount SideEffectAuthorizationItemCount SystemServiceStarted ProductionBusClaimed WriteMethodsEnabled RuntimeWritesEnabled DesktopFilesWritten SupportBundleExported SnapshotRestoreExecuted StateCleanupExecuted FileContentRead FilePathsExposed NetworkRequired HostRootModified PrivilegedContainerRequired StateRootPathExposed RawCommandExposed RawExecutableExposed BackendDetailsExposed validateNoBackendTerms].each do |token|
  assert(go_runtime_production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_receipt_writer_authorization_audit_source.include?(token), "Go Runtime production receipt notification action dry-run result lookup consumer enablement receipt writer authorization audit must expose safety gate #{token}")
end
%w[production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_authorization_receipt_audit.go production_receipt_writer_authorization_review.go docs/claude-code-current-dispatch-picks.md].each do |token|
  assert(go_runtime_production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_receipt_writer_authorization_audit_source.include?(token), "Go Runtime production receipt notification action dry-run result lookup consumer enablement receipt writer authorization audit must consume #{token}")
end
go_runtime_production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_receipt_writer_authorization_audit_test_source = read_project_file("internal/runtime/owner/production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_receipt_writer_authorization_audit_test.go") +
                                                                                                                                     read_project_file("cmd/xnix-runtime-go/production_dbus_gate_cli_test.go") +
                                                                                                                                     read_project_file("cmd/xnix-runtime-go/production_dbus_gate_commands.go") +
                                                                                                                                     read_project_file("cmd/xnix-runtime-go/main.go")
%w[TestProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptWriterAuthorizationAuditPreviewModelsWriter TestProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptWriterAuthorizationAuditPreviewFailsClosedWithoutSources TestProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptWriterAuthorizationAuditPreviewCommand production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-receipt-writer-authorization-audit-preview writer_authorization_audit_required writer_authorization_audit_modeled authorization_receipt_audit_consumed base_writer_authorization_review_consumed writer_authorization_guidance_consumed writer_authorization_boundary_ready consumer_enablement_receipt_writer_ready receipt_writer_authorization_ready receipt_writer_enabled receipt_persistence_enabled receipt_lookup_writes_enabled receipt_replay_enabled authorization_item_count ready_authorization_item_count write_enabled_authorization_item_count persisted_authorization_item_count accepted_authorization_item_count raw_exposed_authorization_item_count state_root_path_exposed file_paths_exposed file_content_read].each do |token|
  assert(go_runtime_production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_receipt_writer_authorization_audit_test_source.include?(token), "Go Runtime production receipt notification action dry-run result lookup consumer enablement receipt writer authorization audit tests and CLI must include #{token}")
end

go_runtime_production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_receipt_persistence_authorization_audit_source = read_project_file("internal/runtime/owner/production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_receipt_persistence_authorization_audit.go")
%w[ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptPersistenceAuthorizationAuditPreview ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptPersistenceAuthorizationAuditItem ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptPersistenceAuthorizationAuditCheck NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptPersistenceAuthorizationAuditPreview xnix.runtime.production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_receipt_persistence_authorization_audit.v1 production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-receipt-persistence-authorization-audit-preview receipt-notification-action-dry-run-result-lookup-consumer-enablement-receipt-persistence-authorization-audit production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-receipt-persistence-authorization-audit-ready-persistence-disabled production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-receipt-persistence-authorization-audit-blocked production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-receipt-writer-authorization-audit-preview production-receipt-persistence-threat-review-preview writer-authorization-audit-consumed base-persistence-threat-review-consumed persistence-authorization-guidance-consumed persistence-authorization-boundary-modeled-only five-persistence-authorization-items-present persistence-authorization-items-ready-persistence-disabled receipt-persistence-replay-expiry-revocation-disabled consumer-lookup-request-and-support-disabled production-and-host-boundary-closed review-receipt-dry-run-result-lookup-consumer-enablement-persistence-authorization renew-receipt-dry-run-result-lookup-consumer-enablement-persistence-authorization open-compatibility-center-dry-run-result-lookup-consumer-enablement-persistence-authorization dismiss-receipt-dry-run-result-lookup-consumer-enablement-persistence-authorization support-info-dry-run-result-lookup-consumer-enablement-persistence-authorization].each do |token|
  assert(go_runtime_production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_receipt_persistence_authorization_audit_source.include?(token), "Go Runtime production receipt notification action dry-run result lookup consumer enablement receipt persistence authorization audit must include #{token}")
end
%w[PersistenceAuthorizationAuditRequired PersistenceAuthorizationAuditModeled WriterAuthorizationAuditConsumed BasePersistenceThreatReviewConsumed PersistenceAuthorizationGuidanceConsumed PersistenceAuthorizationBoundaryReady ReceiptWriterAuthorizationReady ReceiptPersistenceThreatBoundaryReady ConsumerEnablementReceiptPersistenceReady ReceiptPersistenceAuthorizationReady ReceiptPresent ReceiptAccepted AuthorizationAccepted ReceiptWriterEnabled ReceiptPersistenceEnabled ReceiptLookupWritesEnabled ReceiptReplayEnabled ReceiptExpiryWriteEnabled ReceiptRevocationWriteEnabled ConsumerConsumptionAuthorized ConsumerEnablementAuthorized KDEConsumerEnabled RuntimeConsumerEnabled LookupRouteAuthorized LookupRouteEnabled OpaqueLookupEnabled RedactedSummaryPersisted RawResultExposed RuntimeDiagnosticsPersisted DryRunResultPersisted DispatchDryRunExecuted AuthorizationItemCount RequiredAuthorizationItemCount ReadyAuthorizationItemCount MissingAuthorizationItemCount PersistenceEnabledAuthorizationItemCount ReplayEnabledAuthorizationItemCount ExpiryWriteEnabledAuthorizationItemCount RevocationWriteEnabledAuthorizationItemCount AcceptedAuthorizationItemCount EnabledConsumerItemCount RawExposedAuthorizationItemCount SideEffectAuthorizationItemCount SystemServiceStarted ProductionBusClaimed WriteMethodsEnabled RuntimeWritesEnabled DesktopFilesWritten SupportBundleExported SnapshotRestoreExecuted StateCleanupExecuted FileContentRead FilePathsExposed NetworkRequired HostRootModified PrivilegedContainerRequired StateRootPathExposed RawCommandExposed RawExecutableExposed BackendDetailsExposed validateNoBackendTerms].each do |token|
  assert(go_runtime_production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_receipt_persistence_authorization_audit_source.include?(token), "Go Runtime production receipt notification action dry-run result lookup consumer enablement receipt persistence authorization audit must expose safety gate #{token}")
end
%w[production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_receipt_writer_authorization_audit.go production_receipt_persistence_threat_review.go docs/claude-code-current-dispatch-picks.md].each do |token|
  assert(go_runtime_production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_receipt_persistence_authorization_audit_source.include?(token), "Go Runtime production receipt notification action dry-run result lookup consumer enablement receipt persistence authorization audit must consume #{token}")
end
go_runtime_production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_receipt_persistence_authorization_audit_test_source = read_project_file("internal/runtime/owner/production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_receipt_persistence_authorization_audit_test.go") +
                                                                                                                                         read_project_file("cmd/xnix-runtime-go/production_dbus_gate_cli_test.go") +
                                                                                                                                         read_project_file("cmd/xnix-runtime-go/production_dbus_gate_commands.go") +
                                                                                                                                         read_project_file("cmd/xnix-runtime-go/main.go")
%w[TestProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptPersistenceAuthorizationAuditPreviewModelsPersistence TestProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptPersistenceAuthorizationAuditPreviewFailsClosedWithoutSources TestProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptPersistenceAuthorizationAuditPreviewCommand production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-receipt-persistence-authorization-audit-preview persistence_authorization_audit_required persistence_authorization_audit_modeled writer_authorization_audit_consumed base_persistence_threat_review_consumed persistence_authorization_guidance_consumed persistence_authorization_boundary_ready consumer_enablement_receipt_persistence_ready receipt_persistence_authorization_ready receipt_writer_enabled receipt_persistence_enabled receipt_lookup_writes_enabled receipt_replay_enabled receipt_expiry_write_enabled receipt_revocation_write_enabled authorization_item_count ready_authorization_item_count persistence_enabled_authorization_item_count replay_enabled_authorization_item_count expiry_write_enabled_authorization_item_count revocation_write_enabled_authorization_item_count accepted_authorization_item_count raw_exposed_authorization_item_count state_root_path_exposed file_paths_exposed file_content_read].each do |token|
  assert(go_runtime_production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_receipt_persistence_authorization_audit_test_source.include?(token), "Go Runtime production receipt notification action dry-run result lookup consumer enablement receipt persistence authorization audit tests and CLI must include #{token}")
end

go_runtime_production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_receipt_acceptance_authorization_audit_source = read_project_file("internal/runtime/owner/production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_receipt_acceptance_authorization_audit.go")
%w[ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptAcceptanceAuthorizationAuditPreview ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptAcceptanceAuthorizationAuditItem ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptAcceptanceAuthorizationAuditCheck NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptAcceptanceAuthorizationAuditPreview xnix.runtime.production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_receipt_acceptance_authorization_audit.v1 production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-receipt-acceptance-authorization-audit-preview receipt-notification-action-dry-run-result-lookup-consumer-enablement-receipt-acceptance-authorization-audit production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-receipt-acceptance-authorization-audit-ready-acceptance-disabled production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-receipt-acceptance-authorization-audit-blocked production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-receipt-persistence-authorization-audit-preview production-receipt-acceptance-propagation-preflight-preview persistence-authorization-audit-consumed base-acceptance-propagation-preflight-consumed acceptance-authorization-guidance-consumed acceptance-authorization-boundary-modeled-only five-acceptance-authorization-items-present acceptance-authorization-items-ready-acceptance-disabled receipt-acceptance-and-consumption-disabled consumer-lookup-request-and-support-disabled production-and-host-boundary-closed review-receipt-dry-run-result-lookup-consumer-enablement-acceptance-authorization renew-receipt-dry-run-result-lookup-consumer-enablement-acceptance-authorization open-compatibility-center-dry-run-result-lookup-consumer-enablement-acceptance-authorization dismiss-receipt-dry-run-result-lookup-consumer-enablement-acceptance-authorization support-info-dry-run-result-lookup-consumer-enablement-acceptance-authorization].each do |token|
  assert(go_runtime_production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_receipt_acceptance_authorization_audit_source.include?(token), "Go Runtime production receipt notification action dry-run result lookup consumer enablement receipt acceptance authorization audit must include #{token}")
end
%w[AcceptanceAuthorizationAuditRequired AcceptanceAuthorizationAuditModeled PersistenceAuthorizationAuditConsumed BaseAcceptancePropagationPreflightConsumed AcceptanceAuthorizationGuidanceConsumed AcceptanceAuthorizationBoundaryReady ReceiptPersistenceAuthorizationReady ReceiptAcceptancePropagationReady ConsumerEnablementReceiptAcceptanceReady ReceiptAcceptanceAuthorizationReady ReceiptPresent ReceiptPersisted ReceiptAccepted AuthorizationAccepted ReceiptWriterEnabled ReceiptPersistenceEnabled ReceiptLookupWritesEnabled ReceiptReplayEnabled ReceiptExpiryWriteEnabled ReceiptRevocationWriteEnabled ConsumerConsumptionAuthorized ConsumerEnablementAuthorized KDEConsumerEnabled RuntimeConsumerEnabled LookupRouteAuthorized LookupRouteEnabled OpaqueLookupEnabled RedactedSummaryPersisted RawResultExposed RuntimeDiagnosticsPersisted DryRunResultPersisted DispatchDryRunExecuted AuthorizationItemCount RequiredAuthorizationItemCount ReadyAuthorizationItemCount MissingAuthorizationItemCount AcceptedAuthorizationItemCount PersistedAuthorizationItemCount ConsumerAuthorizedItemCount EnabledConsumerItemCount LookupEnabledAuthorizationItemCount RawExposedAuthorizationItemCount SideEffectAuthorizationItemCount SystemServiceStarted ProductionBusClaimed WriteMethodsEnabled RuntimeWritesEnabled DesktopFilesWritten SupportBundleExported SnapshotRestoreExecuted StateCleanupExecuted FileContentRead FilePathsExposed NetworkRequired HostRootModified PrivilegedContainerRequired StateRootPathExposed RawCommandExposed RawExecutableExposed BackendDetailsExposed validateNoBackendTerms].each do |token|
  assert(go_runtime_production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_receipt_acceptance_authorization_audit_source.include?(token), "Go Runtime production receipt notification action dry-run result lookup consumer enablement receipt acceptance authorization audit must expose safety gate #{token}")
end
%w[production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_receipt_persistence_authorization_audit.go production_receipt_acceptance_propagation_preflight.go docs/claude-code-current-dispatch-picks.md].each do |token|
  assert(go_runtime_production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_receipt_acceptance_authorization_audit_source.include?(token), "Go Runtime production receipt notification action dry-run result lookup consumer enablement receipt acceptance authorization audit must consume #{token}")
end
go_runtime_production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_receipt_acceptance_authorization_audit_test_source = read_project_file("internal/runtime/owner/production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_receipt_acceptance_authorization_audit_test.go") +
                                                                                                                                          read_project_file("cmd/xnix-runtime-go/production_dbus_gate_cli_test.go") +
                                                                                                                                          read_project_file("cmd/xnix-runtime-go/production_dbus_gate_commands.go") +
                                                                                                                                          read_project_file("cmd/xnix-runtime-go/main.go")
%w[TestProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptAcceptanceAuthorizationAuditPreviewModelsAcceptance TestProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptAcceptanceAuthorizationAuditPreviewFailsClosedWithoutSources TestProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptAcceptanceAuthorizationAuditPreviewCommand production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-receipt-acceptance-authorization-audit-preview acceptance_authorization_audit_required acceptance_authorization_audit_modeled persistence_authorization_audit_consumed base_acceptance_propagation_preflight_consumed acceptance_authorization_guidance_consumed acceptance_authorization_boundary_ready consumer_enablement_receipt_acceptance_ready receipt_acceptance_authorization_ready receipt_persisted receipt_accepted authorization_item_count ready_authorization_item_count accepted_authorization_item_count persisted_authorization_item_count consumer_authorized_item_count enabled_consumer_item_count lookup_enabled_authorization_item_count raw_exposed_authorization_item_count state_root_path_exposed file_paths_exposed file_content_read].each do |token|
  assert(go_runtime_production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_receipt_acceptance_authorization_audit_test_source.include?(token), "Go Runtime production receipt notification action dry-run result lookup consumer enablement receipt acceptance authorization audit tests and CLI must include #{token}")
end

go_runtime_production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_receipt_consumption_gate_audit_source = read_project_file("internal/runtime/owner/production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_receipt_consumption_gate_audit.go")
%w[ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumptionGateAuditPreview ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumptionGateAuditItem ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumptionGateAuditCheck NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumptionGateAuditPreview xnix.runtime.production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_receipt_consumption_gate_audit.v1 production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-receipt-consumption-gate-audit-preview receipt-notification-action-dry-run-result-lookup-consumer-enablement-receipt-consumption-gate-audit production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-receipt-consumption-gate-audit-ready-consumption-disabled production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-receipt-consumption-gate-audit-blocked production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-receipt-acceptance-authorization-audit-preview production-receipt-notification-action-dry-run-result-lookup-consumer-redaction-audit-preview production-receipt-notification-action-dry-run-result-lookup-route-authorization-audit-preview acceptance-authorization-audit-consumed consumer-redaction-audit-consumed lookup-route-authorization-audit-consumed consumption-gate-guidance-consumed consumption-gate-modeled-only five-consumption-gate-items-present consumption-gate-items-ready-consumption-disabled receipt-acceptance-and-consumption-disabled consumer-lookup-request-and-support-disabled production-and-host-boundary-closed review-receipt-dry-run-result-lookup-consumer-enablement-consumption-gate renew-receipt-dry-run-result-lookup-consumer-enablement-consumption-gate open-compatibility-center-dry-run-result-lookup-consumer-enablement-consumption-gate dismiss-receipt-dry-run-result-lookup-consumer-enablement-consumption-gate support-info-dry-run-result-lookup-consumer-enablement-consumption-gate].each do |token|
  assert(go_runtime_production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_receipt_consumption_gate_audit_source.include?(token), "Go Runtime production receipt notification action dry-run result lookup consumer enablement receipt consumption gate audit must include #{token}")
end
%w[ConsumptionGateRequired ConsumptionGateModeled AcceptanceAuthorizationAuditConsumed ConsumerRedactionAuditConsumed LookupRouteAuthorizationAuditConsumed ConsumptionGateGuidanceConsumed ConsumptionGateBoundaryReady ReceiptAcceptanceAuthorizationReady ConsumerRedactionBoundaryReady LookupRouteAuthorizationBoundaryReady ConsumerEnablementReceiptConsumptionReady ReceiptConsumptionGateReady ReceiptPresent ReceiptPersisted ReceiptAccepted ReceiptConsumed AuthorizationAccepted ReceiptWriterEnabled ReceiptPersistenceEnabled ReceiptLookupWritesEnabled ReceiptReplayEnabled ReceiptExpiryWriteEnabled ReceiptRevocationWriteEnabled ConsumerConsumptionAuthorized ConsumerEnablementAuthorized KDEConsumerEnabled RuntimeConsumerEnabled LookupRouteAuthorized LookupRouteEnabled OpaqueLookupEnabled RedactedSummaryPersisted RawResultExposed RuntimeDiagnosticsPersisted DryRunResultPersisted DispatchDryRunExecuted GateItemCount RequiredGateItemCount ReadyGateItemCount MissingGateItemCount AcceptedReceiptItemCount ConsumedReceiptItemCount AuthorizedConsumerItemCount EnabledConsumerItemCount LookupEnabledGateItemCount RawExposedGateItemCount SideEffectGateItemCount SystemServiceStarted ProductionBusClaimed WriteMethodsEnabled RuntimeWritesEnabled DesktopFilesWritten SupportBundleExported SnapshotRestoreExecuted StateCleanupExecuted FileContentRead FilePathsExposed NetworkRequired HostRootModified PrivilegedContainerRequired StateRootPathExposed RawCommandExposed RawExecutableExposed BackendDetailsExposed validateNoBackendTerms].each do |token|
  assert(go_runtime_production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_receipt_consumption_gate_audit_source.include?(token), "Go Runtime production receipt notification action dry-run result lookup consumer enablement receipt consumption gate audit must expose safety gate #{token}")
end
%w[production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_receipt_acceptance_authorization_audit.go production_receipt_notification_action_dry_run_result_lookup_consumer_redaction_audit.go production_receipt_notification_action_dry_run_result_lookup_route_authorization_audit.go docs/claude-code-current-dispatch-picks.md].each do |token|
  assert(go_runtime_production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_receipt_consumption_gate_audit_source.include?(token), "Go Runtime production receipt notification action dry-run result lookup consumer enablement receipt consumption gate audit must consume #{token}")
end
go_runtime_production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_receipt_consumption_gate_audit_test_source = read_project_file("internal/runtime/owner/production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_receipt_consumption_gate_audit_test.go") +
                                                                                                                                      read_project_file("cmd/xnix-runtime-go/production_dbus_gate_cli_test.go") +
                                                                                                                                      read_project_file("cmd/xnix-runtime-go/production_dbus_gate_commands.go") +
                                                                                                                                      read_project_file("cmd/xnix-runtime-go/main.go")
%w[TestProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumptionGateAuditPreviewModelsConsumption TestProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumptionGateAuditPreviewFailsClosedWithoutSources TestProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumptionGateAuditPreviewCommand production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-receipt-consumption-gate-audit-preview consumption_gate_required consumption_gate_modeled acceptance_authorization_audit_consumed consumer_redaction_audit_consumed lookup_route_authorization_audit_consumed consumption_gate_guidance_consumed consumption_gate_boundary_ready consumer_enablement_receipt_consumption_ready receipt_consumption_gate_ready receipt_accepted receipt_consumed gate_item_count ready_gate_item_count accepted_receipt_item_count consumed_receipt_item_count authorized_consumer_item_count enabled_consumer_item_count lookup_enabled_gate_item_count raw_exposed_gate_item_count state_root_path_exposed file_paths_exposed file_content_read].each do |token|
  assert(go_runtime_production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_receipt_consumption_gate_audit_test_source.include?(token), "Go Runtime production receipt notification action dry-run result lookup consumer enablement receipt consumption gate audit tests and CLI must include #{token}")
end

go_runtime_production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_receipt_consumer_authorization_audit_source = read_project_file("internal/runtime/owner/production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_receipt_consumer_authorization_audit.go")
%w[ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerAuthorizationAuditPreview ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerAuthorizationItem ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerAuthorizationCheck NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerAuthorizationAuditPreview xnix.runtime.production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_receipt_consumer_authorization_audit.v1 production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-receipt-consumer-authorization-audit-preview receipt-notification-action-dry-run-result-lookup-consumer-enablement-receipt-consumer-authorization-audit production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-receipt-consumer-authorization-audit-ready-authorization-disabled production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-receipt-consumer-authorization-audit-blocked production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-receipt-consumption-gate-audit-preview production-receipt-notification-action-dry-run-result-lookup-consumer-redaction-audit-preview production-receipt-notification-action-dry-run-result-lookup-route-authorization-audit-preview consumption-gate-audit-consumed consumer-redaction-audit-consumed lookup-route-authorization-audit-consumed consumer-authorization-guidance-consumed consumer-authorization-boundary-modeled-only five-consumer-authorization-items-present consumer-authorization-items-ready-authorization-disabled receipt-consumption-and-consumer-authorization-disabled consumer-lookup-request-and-support-disabled production-and-host-boundary-closed review-receipt-dry-run-result-lookup-consumer-enablement-consumer-authorization renew-receipt-dry-run-result-lookup-consumer-enablement-consumer-authorization open-compatibility-center-dry-run-result-lookup-consumer-enablement-consumer-authorization dismiss-receipt-dry-run-result-lookup-consumer-enablement-consumer-authorization support-info-dry-run-result-lookup-consumer-enablement-consumer-authorization].each do |token|
  assert(go_runtime_production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_receipt_consumer_authorization_audit_source.include?(token), "Go Runtime production receipt notification action dry-run result lookup consumer enablement receipt consumer authorization audit must include #{token}")
end
%w[ConsumerAuthorizationAuditRequired ConsumerAuthorizationAuditModeled ConsumptionGateAuditConsumed ConsumerRedactionAuditConsumed LookupRouteAuthorizationAuditConsumed ConsumerAuthorizationGuidanceConsumed ConsumerAuthorizationBoundaryReady ReceiptConsumptionGateReady ConsumerRedactionBoundaryReady LookupRouteAuthorizationBoundaryReady ConsumerAuthorizationReady ReceiptPresent ReceiptPersisted ReceiptAccepted ReceiptConsumed AuthorizationAccepted ConsumerAuthorizationGranted ConsumerConsumptionAuthorized ConsumerEnablementAuthorized KDEConsumerAuthorized RuntimeConsumerAuthorized KDEConsumerEnabled RuntimeConsumerEnabled LookupRouteAuthorized LookupRouteEnabled OpaqueLookupEnabled RedactedSummaryPersisted RawResultExposed RuntimeDiagnosticsPersisted DryRunResultPersisted DispatchDryRunExecuted AuthorizationItemCount RequiredAuthorizationItemCount ReadyAuthorizationItemCount MissingAuthorizationItemCount GrantedAuthorizationItemCount AcceptedReceiptItemCount ConsumedReceiptItemCount AuthorizedConsumerItemCount EnabledConsumerItemCount LookupEnabledAuthorizationItemCount RawExposedAuthorizationItemCount SideEffectAuthorizationItemCount SystemServiceStarted ProductionBusClaimed WriteMethodsEnabled RuntimeWritesEnabled DesktopFilesWritten SupportBundleExported SnapshotRestoreExecuted StateCleanupExecuted FileContentRead FilePathsExposed NetworkRequired HostRootModified PrivilegedContainerRequired StateRootPathExposed RawCommandExposed RawExecutableExposed BackendDetailsExposed validateNoBackendTerms].each do |token|
  assert(go_runtime_production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_receipt_consumer_authorization_audit_source.include?(token), "Go Runtime production receipt notification action dry-run result lookup consumer enablement receipt consumer authorization audit must expose safety gate #{token}")
end
%w[production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_receipt_consumption_gate_audit.go production_receipt_notification_action_dry_run_result_lookup_consumer_redaction_audit.go production_receipt_notification_action_dry_run_result_lookup_route_authorization_audit.go docs/claude-code-current-dispatch-picks.md].each do |token|
  assert(go_runtime_production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_receipt_consumer_authorization_audit_source.include?(token), "Go Runtime production receipt notification action dry-run result lookup consumer enablement receipt consumer authorization audit must consume #{token}")
end
go_runtime_production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_receipt_consumer_authorization_audit_test_source = read_project_file("internal/runtime/owner/production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_receipt_consumer_authorization_audit_test.go") +
                                                                                                                                                    read_project_file("cmd/xnix-runtime-go/production_dbus_gate_cli_test.go") +
                                                                                                                                                    read_project_file("cmd/xnix-runtime-go/production_dbus_gate_commands.go") +
                                                                                                                                                    read_project_file("cmd/xnix-runtime-go/main.go")
%w[TestProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerAuthorizationAuditPreviewModelsAuthorization TestProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerAuthorizationAuditPreviewFailsClosedWithoutSources TestProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerAuthorizationAuditPreviewCommand production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-receipt-consumer-authorization-audit-preview consumer_authorization_audit_required consumer_authorization_audit_modeled consumption_gate_audit_consumed consumer_redaction_audit_consumed lookup_route_authorization_audit_consumed consumer_authorization_guidance_consumed consumer_authorization_boundary_ready consumer_authorization_ready consumer_authorization_granted consumer_consumption_authorized consumer_enablement_authorized kde_consumer_authorized runtime_consumer_authorized kde_consumer_enabled runtime_consumer_enabled authorization_item_count ready_authorization_item_count granted_authorization_item_count authorized_consumer_item_count enabled_consumer_item_count state_root_path_exposed file_paths_exposed file_content_read].each do |token|
  assert(go_runtime_production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_receipt_consumer_authorization_audit_test_source.include?(token), "Go Runtime production receipt notification action dry-run result lookup consumer enablement receipt consumer authorization audit tests and CLI must include #{token}")
end

go_runtime_production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_receipt_consumer_enablement_gate_audit_source = read_project_file("internal/runtime/owner/production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_receipt_consumer_enablement_gate_audit.go")
%w[ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerEnablementGateAuditPreview ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerEnablementGateAuditItem ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerEnablementGateAuditCheck NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerEnablementGateAuditPreview xnix.runtime.production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_receipt_consumer_enablement_gate_audit.v1 production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-receipt-consumer-enablement-gate-audit-preview receipt-notification-action-dry-run-result-lookup-consumer-enablement-receipt-consumer-enablement-gate-audit production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-receipt-consumer-enablement-gate-audit-ready-consumers-disabled production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-receipt-consumer-enablement-gate-audit-blocked production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-receipt-consumer-authorization-audit-preview production-receipt-notification-action-dry-run-result-lookup-consumer-redaction-audit-preview production-receipt-notification-action-dry-run-result-lookup-route-authorization-audit-preview consumer-authorization-audit-consumed lookup-route-authorization-audit-consumed consumer-redaction-audit-consumed consumer-enablement-guidance-consumed consumer-enablement-receipt-consumer-enablement-gate-modeled-only five-consumer-enablement-receipt-consumer-enablement-gate-items-present lookup-result-persistence-and-execution-disabled request-notification-navigation-and-support-disabled production-and-host-boundary-closed review-receipt-dry-run-result-lookup-consumer-enablement-receipt-consumer-enablement-gate renew-receipt-dry-run-result-lookup-consumer-enablement-receipt-consumer-enablement-gate open-compatibility-center-dry-run-result-lookup-consumer-enablement-receipt-consumer-enablement-gate dismiss-receipt-dry-run-result-lookup-consumer-enablement-receipt-consumer-enablement-gate support-info-dry-run-result-lookup-consumer-enablement-receipt-consumer-enablement-gate].each do |token|
  assert(go_runtime_production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_receipt_consumer_enablement_gate_audit_source.include?(token), "Go Runtime production receipt notification action dry-run result lookup consumer enablement receipt consumer enablement gate audit must include #{token}")
end
%w[ConsumerEnablementGateRequired ConsumerEnablementGateModeled ConsumerAuthorizationAuditConsumed LookupRouteAuthorizationAuditConsumed ConsumerRedactionAuditConsumed ConsumerEnablementGuidanceConsumed ConsumerEnablementGateReady ConsumerAuthorizationPrerequisiteModeled RouteAuthorizationPrerequisiteModeled ConsumerRedactionPrerequisiteModeled OpaqueResultIDSupported ConsumerConsumptionAuthorized ConsumerEnablementAuthorized KDEConsumerEnabled RuntimeConsumerEnabled LookupRouteAuthorized LookupRouteEnabled OpaqueLookupEnabled RedactedSummaryPersisted RawResultExposed RuntimeDiagnosticsPersisted DryRunResultPersisted DispatchDryRunExecuted GateItemCount RequiredGateItemCount ReadyGateItemCount MissingGateItemCount AuthorizedGateItemCount EnabledGateItemCount PersistedGateItemCount RawExposedGateItemCount SideEffectGateItemCount SystemServiceStarted ProductionBusClaimed WriteMethodsEnabled RuntimeWritesEnabled DesktopFilesWritten SupportBundleExported SnapshotRestoreExecuted StateCleanupExecuted FileContentRead FilePathsExposed NetworkRequired HostRootModified PrivilegedContainerRequired StateRootPathExposed RawCommandExposed RawExecutableExposed BackendDetailsExposed validateNoBackendTerms].each do |token|
  assert(go_runtime_production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_receipt_consumer_enablement_gate_audit_source.include?(token), "Go Runtime production receipt notification action dry-run result lookup consumer enablement receipt consumer enablement gate audit must expose safety gate #{token}")
end
%w[production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_receipt_consumer_authorization_audit.go production_receipt_notification_action_dry_run_result_lookup_consumer_redaction_audit.go production_receipt_notification_action_dry_run_result_lookup_route_authorization_audit.go docs/claude-code-current-dispatch-picks.md].each do |token|
  assert(go_runtime_production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_receipt_consumer_enablement_gate_audit_source.include?(token), "Go Runtime production receipt notification action dry-run result lookup consumer enablement receipt consumer enablement gate audit must consume #{token}")
end
go_runtime_production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_receipt_consumer_enablement_gate_audit_test_source = read_project_file("internal/runtime/owner/production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_receipt_consumer_enablement_gate_audit_test.go") +
                                                                                                                                                    read_project_file("cmd/xnix-runtime-go/production_dbus_gate_cli_test.go") +
                                                                                                                                                    read_project_file("cmd/xnix-runtime-go/production_dbus_gate_commands.go") +
                                                                                                                                                    read_project_file("cmd/xnix-runtime-go/main.go")
%w[TestProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerEnablementGateAuditPreviewModelsGate TestProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerEnablementGateAuditPreviewFailsClosedWithoutSources TestProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerEnablementGateAuditPreviewCommand production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-receipt-consumer-enablement-gate-audit-preview consumer_enablement_gate_required consumer_enablement_gate_modeled consumer_authorization_audit_consumed lookup_route_authorization_audit_consumed consumer_redaction_audit_consumed consumer_enablement_guidance_consumed consumer_enablement_gate_ready consumer_authorization_prerequisite_modeled route_authorization_prerequisite_modeled consumer_redaction_prerequisite_modeled consumer_consumption_authorized consumer_enablement_authorized kde_consumer_enabled runtime_consumer_enabled gate_item_count ready_gate_item_count authorized_gate_item_count enabled_gate_item_count state_root_path_exposed file_paths_exposed file_content_read].each do |token|
  assert(go_runtime_production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_receipt_consumer_enablement_gate_audit_test_source.include?(token), "Go Runtime production receipt notification action dry-run result lookup consumer enablement receipt consumer enablement gate audit tests and CLI must include #{token}")
end

go_runtime_production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_kde_safe_status_fanout_audit_source = read_project_file("internal/runtime/owner/production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_kde_safe_status_fanout_audit.go")
%w[ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusFanOutAuditPreview ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusFanOutAuditItem ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusFanOutAuditCheck NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusFanOutAuditPreview xnix.runtime.production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_kde_safe_status_fanout_audit.v1 production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-status-fanout-audit-preview receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-status-fanout-audit production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-status-fanout-audit-ready-status-only-consumers-disabled production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-status-fanout-audit-blocked production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-receipt-consumer-enablement-gate-audit-preview kde-safe-status-fanout-audit-required consumer-enablement-gate-audit-consumed kde-safe-status-guidance-consumed compatibility-center-status-modeled runtime-diagnostics-status-modeled ten-kde-safe-status-fanout-items-present status-fanout-modeled-only consumer-lookup-and-persistence-disabled request-notification-navigation-and-support-disabled production-and-host-boundary-closed review-receipt-dry-run-result-lookup-consumer-enablement-kde-safe-compatibility-center-status-fanout review-receipt-dry-run-result-lookup-consumer-enablement-kde-safe-runtime-diagnostics-status-fanout renew-receipt-dry-run-result-lookup-consumer-enablement-kde-safe-compatibility-center-status-fanout renew-receipt-dry-run-result-lookup-consumer-enablement-kde-safe-runtime-diagnostics-status-fanout open-compatibility-center-dry-run-result-lookup-consumer-enablement-kde-safe-compatibility-center-status-fanout open-compatibility-center-dry-run-result-lookup-consumer-enablement-kde-safe-runtime-diagnostics-status-fanout dismiss-receipt-dry-run-result-lookup-consumer-enablement-kde-safe-compatibility-center-status-fanout dismiss-receipt-dry-run-result-lookup-consumer-enablement-kde-safe-runtime-diagnostics-status-fanout support-info-dry-run-result-lookup-consumer-enablement-kde-safe-compatibility-center-status-fanout support-info-dry-run-result-lookup-consumer-enablement-kde-safe-runtime-diagnostics-status-fanout].each do |token|
  assert(go_runtime_production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_kde_safe_status_fanout_audit_source.include?(token), "Go Runtime production receipt notification action dry-run result lookup consumer enablement KDE-safe status fan-out audit must include #{token}")
end
%w[StatusFanOutAuditRequired StatusFanOutModeled ConsumerEnablementGateAuditConsumed KDESafeStatusGuidanceConsumed CompatibilityCenterStatusModeled RuntimeDiagnosticsStatusModeled StatusFanOutReady ConsumerEnablementGateClosed ConsumerAuthorizationPrerequisiteSeen RouteAuthorizationPrerequisiteSeen ConsumerRedactionPrerequisiteSeen OpaqueResultIDSupported ConsumerConsumptionAuthorized ConsumerEnablementAuthorized KDEConsumerEnabled RuntimeConsumerEnabled LookupRouteAuthorized LookupRouteEnabled OpaqueLookupEnabled RedactedSummaryPersisted KDEStatusPersisted RuntimeDiagnosticsPersisted DryRunResultPersisted RawResultExposed DispatchDryRunExecuted StatusItemCount RequiredStatusItemCount ReadyStatusItemCount MissingStatusItemCount CompatibilityCenterStatusItemCount RuntimeDiagnosticsStatusItemCount ConsumerEnabledStatusItemCount PersistedStatusItemCount RawExposedStatusItemCount SideEffectStatusItemCount RequestObjectCreationEnabled RequestObjectDispatchEnabled PortalRequestCreated NotificationActionEnabled CompatibilityCenterOpened ProductionReadiness ProductionOwnershipReady SystemServiceStarted ProductionBusClaimed WriteMethodsEnabled RuntimeWritesEnabled DesktopFilesWritten SupportBundleExported SnapshotRestoreExecuted StateCleanupExecuted FileContentRead FilePathsExposed NetworkRequired HostRootModified PrivilegedContainerRequired StateRootPathExposed RawCommandExposed RawExecutableExposed BackendDetailsExposed validateNoBackendTerms].each do |token|
  assert(go_runtime_production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_kde_safe_status_fanout_audit_source.include?(token), "Go Runtime production receipt notification action dry-run result lookup consumer enablement KDE-safe status fan-out audit must expose safety gate #{token}")
end
%w[production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_receipt_consumer_enablement_gate_audit.go docs/claude-code-current-dispatch-picks.md].each do |token|
  assert(go_runtime_production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_kde_safe_status_fanout_audit_source.include?(token), "Go Runtime production receipt notification action dry-run result lookup consumer enablement KDE-safe status fan-out audit must consume #{token}")
end
go_runtime_production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_kde_safe_status_fanout_audit_test_source = read_project_file("internal/runtime/owner/production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_kde_safe_status_fanout_audit_test.go") +
                                                                                                                                                    read_project_file("cmd/xnix-runtime-go/production_dbus_gate_cli_test.go") +
                                                                                                                                                    read_project_file("cmd/xnix-runtime-go/production_dbus_gate_commands.go") +
                                                                                                                                                    read_project_file("cmd/xnix-runtime-go/main.go")
%w[TestProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusFanOutAuditPreviewModelsStatusFanOut TestProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusFanOutAuditPreviewFailsClosedWithoutSources TestProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusFanOutAuditPreviewCommand production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-status-fanout-audit-preview status_fanout_audit_required status_fanout_modeled consumer_enablement_gate_audit_consumed kde_safe_status_guidance_consumed compatibility_center_status_modeled runtime_diagnostics_status_modeled status_fanout_ready consumer_enablement_gate_closed compatibility_center_status_item_count runtime_diagnostics_status_item_count status_item_count ready_status_item_count missing_status_item_count consumer_enabled_status_item_count persisted_status_item_count raw_exposed_status_item_count side_effect_status_item_count kde_status_persisted runtime_diagnostics_persisted state_root_path_exposed file_paths_exposed file_content_read].each do |token|
  assert(go_runtime_production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_kde_safe_status_fanout_audit_test_source.include?(token), "Go Runtime production receipt notification action dry-run result lookup consumer enablement KDE-safe status fan-out audit tests and CLI must include #{token}")
end

go_runtime_production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_kde_safe_status_persistence_authorization_audit_source = read_project_file("internal/runtime/owner/production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_kde_safe_status_persistence_authorization_audit.go")
%w[ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusPersistenceAuthorizationAuditPreview ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusPersistenceAuthorizationAuditItem ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusPersistenceAuthorizationAuditCheck NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusPersistenceAuthorizationAuditPreview xnix.runtime.production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_kde_safe_status_persistence_authorization_audit.v1 production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-status-persistence-authorization-audit-preview receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-status-persistence-authorization-audit production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-status-persistence-authorization-audit-ready-persistence-disabled production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-status-persistence-authorization-audit-blocked production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-status-fanout-audit-preview kde-safe-status-persistence-authorization-audit-required status-fanout-audit-consumed kde-safe-status-persistence-guidance-consumed compatibility-center-persistence-authorization-modeled runtime-diagnostics-persistence-authorization-modeled ten-kde-safe-status-persistence-authorization-items-present persistence-authorization-boundary-modeled-only persistence-writes-consumers-and-lookup-disabled request-notification-navigation-and-support-disabled production-and-host-boundary-closed review-receipt-dry-run-result-lookup-consumer-enablement-kde-safe-compatibility-center-status-persistence-authorization review-receipt-dry-run-result-lookup-consumer-enablement-kde-safe-runtime-diagnostics-status-persistence-authorization renew-receipt-dry-run-result-lookup-consumer-enablement-kde-safe-compatibility-center-status-persistence-authorization renew-receipt-dry-run-result-lookup-consumer-enablement-kde-safe-runtime-diagnostics-status-persistence-authorization open-compatibility-center-dry-run-result-lookup-consumer-enablement-kde-safe-compatibility-center-status-persistence-authorization open-compatibility-center-dry-run-result-lookup-consumer-enablement-kde-safe-runtime-diagnostics-status-persistence-authorization dismiss-receipt-dry-run-result-lookup-consumer-enablement-kde-safe-compatibility-center-status-persistence-authorization dismiss-receipt-dry-run-result-lookup-consumer-enablement-kde-safe-runtime-diagnostics-status-persistence-authorization support-info-dry-run-result-lookup-consumer-enablement-kde-safe-compatibility-center-status-persistence-authorization support-info-dry-run-result-lookup-consumer-enablement-kde-safe-runtime-diagnostics-status-persistence-authorization].each do |token|
  assert(go_runtime_production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_kde_safe_status_persistence_authorization_audit_source.include?(token), "Go Runtime production receipt notification action dry-run result lookup consumer enablement KDE-safe status persistence authorization audit must include #{token}")
end
%w[PersistenceAuthorizationAuditRequired PersistenceAuthorizationModeled StatusFanOutAuditConsumed KDESafeStatusPersistenceGuidanceConsumed CompatibilityCenterPersistenceAuthorizationModeled RuntimeDiagnosticsPersistenceAuthorizationModeled PersistenceAuthorizationBoundaryReady StatusFanOutReady ConsumerEnablementGateClosed KDESafeStatusOnly OpaqueResultIDSupported StatusPersistenceAuthorized KDEStatusPersistenceAuthorized RuntimeDiagnosticsPersistenceAuthorized ConsumerConsumptionAuthorized ConsumerEnablementAuthorized KDEConsumerEnabled RuntimeConsumerEnabled LookupRouteAuthorized LookupRouteEnabled OpaqueLookupEnabled RedactedSummaryPersisted KDEStatusPersisted RuntimeDiagnosticsPersisted DryRunResultPersisted RawResultExposed DispatchDryRunExecuted AuthorizationItemCount RequiredAuthorizationItemCount ReadyAuthorizationItemCount MissingAuthorizationItemCount CompatibilityCenterAuthorizationItemCount RuntimeDiagnosticsAuthorizationItemCount AuthorizedPersistenceItemCount PersistedStatusItemCount ConsumerEnabledAuthorizationItemCount RawExposedAuthorizationItemCount SideEffectAuthorizationItemCount RequestObjectCreationEnabled RequestObjectDispatchEnabled PortalRequestCreated NotificationActionEnabled CompatibilityCenterOpened ProductionReadiness ProductionOwnershipReady SystemServiceStarted ProductionBusClaimed WriteMethodsEnabled RuntimeWritesEnabled DesktopFilesWritten SupportBundleExported SnapshotRestoreExecuted StateCleanupExecuted FileContentRead FilePathsExposed NetworkRequired HostRootModified PrivilegedContainerRequired StateRootPathExposed RawCommandExposed RawExecutableExposed BackendDetailsExposed validateNoBackendTerms].each do |token|
  assert(go_runtime_production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_kde_safe_status_persistence_authorization_audit_source.include?(token), "Go Runtime production receipt notification action dry-run result lookup consumer enablement KDE-safe status persistence authorization audit must expose safety gate #{token}")
end
%w[production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_kde_safe_status_fanout_audit.go docs/claude-code-current-dispatch-picks.md].each do |token|
  assert(go_runtime_production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_kde_safe_status_persistence_authorization_audit_source.include?(token), "Go Runtime production receipt notification action dry-run result lookup consumer enablement KDE-safe status persistence authorization audit must consume #{token}")
end
go_runtime_production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_kde_safe_status_persistence_authorization_audit_test_source = read_project_file("internal/runtime/owner/production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_kde_safe_status_persistence_authorization_audit_test.go") +
                                                                                                                                                    read_project_file("cmd/xnix-runtime-go/production_dbus_gate_cli_test.go") +
                                                                                                                                                    read_project_file("cmd/xnix-runtime-go/production_dbus_gate_commands.go") +
                                                                                                                                                    read_project_file("cmd/xnix-runtime-go/main.go")
%w[TestProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusPersistenceAuthorizationAuditPreviewModelsAuthorization TestProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusPersistenceAuthorizationAuditPreviewFailsClosedWithoutSources TestProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusPersistenceAuthorizationAuditPreviewCommand production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-status-persistence-authorization-audit-preview persistence_authorization_audit_required persistence_authorization_modeled status_fanout_audit_consumed kde_safe_status_persistence_guidance_consumed compatibility_center_persistence_authorization_modeled runtime_diagnostics_persistence_authorization_modeled persistence_authorization_boundary_ready status_persistence_authorized kde_status_persistence_authorized runtime_diagnostics_persistence_authorized authorization_item_count ready_authorization_item_count missing_authorization_item_count authorized_persistence_item_count persisted_status_item_count consumer_enabled_authorization_item_count raw_exposed_authorization_item_count side_effect_authorization_item_count kde_status_persisted runtime_diagnostics_persisted state_root_path_exposed file_paths_exposed file_content_read].each do |token|
  assert(go_runtime_production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_kde_safe_status_persistence_authorization_audit_test_source.include?(token), "Go Runtime production receipt notification action dry-run result lookup consumer enablement KDE-safe status persistence authorization audit tests and CLI must include #{token}")
end

go_runtime_production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_kde_safe_redacted_status_persistence_write_model_audit_source = read_project_file("internal/runtime/owner/production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_kde_safe_redacted_status_persistence_write_model_audit.go")
%w[ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusPersistenceWriteModelAuditPreview ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusPersistenceWriteModelAuditItem ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusPersistenceWriteModelAuditCheck NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusPersistenceWriteModelAuditPreview xnix.runtime.production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_kde_safe_redacted_status_persistence_write_model_audit.v1 production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-persistence-write-model-audit-preview receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-persistence-write-model-audit production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-persistence-write-model-audit-ready-writes-disabled production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-persistence-write-model-audit-blocked production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-status-persistence-authorization-audit-preview kde-safe-redacted-status-write-model-audit-required persistence-authorization-audit-consumed redacted-write-model-guidance-consumed compatibility-center-write-model-modeled runtime-diagnostics-write-model-modeled ten-kde-safe-redacted-status-write-model-items-present redacted-shape-modeled-only writes-consumers-and-lookup-disabled request-notification-navigation-and-support-disabled production-and-host-boundary-closed review-receipt-dry-run-result-lookup-consumer-enablement-kde-safe-compatibility-center-redacted-status-write-model review-receipt-dry-run-result-lookup-consumer-enablement-kde-safe-runtime-diagnostics-redacted-status-write-model renew-receipt-dry-run-result-lookup-consumer-enablement-kde-safe-compatibility-center-redacted-status-write-model renew-receipt-dry-run-result-lookup-consumer-enablement-kde-safe-runtime-diagnostics-redacted-status-write-model open-compatibility-center-dry-run-result-lookup-consumer-enablement-kde-safe-compatibility-center-redacted-status-write-model open-compatibility-center-dry-run-result-lookup-consumer-enablement-kde-safe-runtime-diagnostics-redacted-status-write-model dismiss-receipt-dry-run-result-lookup-consumer-enablement-kde-safe-compatibility-center-redacted-status-write-model dismiss-receipt-dry-run-result-lookup-consumer-enablement-kde-safe-runtime-diagnostics-redacted-status-write-model support-info-dry-run-result-lookup-consumer-enablement-kde-safe-compatibility-center-redacted-status-write-model support-info-dry-run-result-lookup-consumer-enablement-kde-safe-runtime-diagnostics-redacted-status-write-model].each do |token|
  assert(go_runtime_production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_kde_safe_redacted_status_persistence_write_model_audit_source.include?(token), "Go Runtime production receipt notification action dry-run result lookup consumer enablement KDE-safe redacted status persistence write-model audit must include #{token}")
end
%w[WriteModelAuditRequired WriteModelModeled PersistenceAuthorizationAuditConsumed RedactedWriteModelGuidanceConsumed CompatibilityCenterWriteModelModeled RuntimeDiagnosticsWriteModelModeled WriteModelBoundaryReady PersistenceAuthorizationBoundaryReady StatusFanOutReady KDESafeStatusOnly OpaqueResultIDSupported RedactedSummaryShapeModeled StatusPersistenceAuthorized StatusPersistenceWriteEnabled KDEStatusWriteEnabled RuntimeDiagnosticsWriteEnabled ConsumerConsumptionAuthorized ConsumerEnablementAuthorized KDEConsumerEnabled RuntimeConsumerEnabled LookupRouteAuthorized LookupRouteEnabled OpaqueLookupEnabled RedactedSummaryPersisted KDEStatusPersisted RuntimeDiagnosticsPersisted DryRunResultPersisted RawResultExposed DispatchDryRunExecuted WriteModelItemCount RequiredWriteModelItemCount ReadyWriteModelItemCount MissingWriteModelItemCount CompatibilityCenterWriteModelItemCount RuntimeDiagnosticsWriteModelItemCount WriteEnabledItemCount PersistedStatusItemCount ConsumerEnabledWriteModelItemCount RawExposedWriteModelItemCount SideEffectWriteModelItemCount RequestObjectCreationEnabled RequestObjectDispatchEnabled PortalRequestCreated NotificationActionEnabled CompatibilityCenterOpened ProductionReadiness ProductionOwnershipReady SystemServiceStarted ProductionBusClaimed WriteMethodsEnabled RuntimeWritesEnabled DesktopFilesWritten SupportBundleExported SnapshotRestoreExecuted StateCleanupExecuted FileContentRead FilePathsExposed NetworkRequired HostRootModified PrivilegedContainerRequired StateRootPathExposed RawCommandExposed RawExecutableExposed BackendDetailsExposed validateNoBackendTerms].each do |token|
  assert(go_runtime_production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_kde_safe_redacted_status_persistence_write_model_audit_source.include?(token), "Go Runtime production receipt notification action dry-run result lookup consumer enablement KDE-safe redacted status persistence write-model audit must expose safety gate #{token}")
end
%w[production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_kde_safe_status_persistence_authorization_audit.go docs/claude-code-current-dispatch-picks.md].each do |token|
  assert(go_runtime_production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_kde_safe_redacted_status_persistence_write_model_audit_source.include?(token), "Go Runtime production receipt notification action dry-run result lookup consumer enablement KDE-safe redacted status persistence write-model audit must consume #{token}")
end
go_runtime_production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_kde_safe_redacted_status_persistence_write_model_audit_test_source = read_project_file("internal/runtime/owner/production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_kde_safe_redacted_status_persistence_write_model_audit_test.go") +
                                                                                                                                                    read_project_file("cmd/xnix-runtime-go/production_dbus_gate_cli_test.go") +
                                                                                                                                                    read_project_file("cmd/xnix-runtime-go/production_dbus_gate_commands.go") +
                                                                                                                                                    read_project_file("cmd/xnix-runtime-go/main.go")
%w[TestProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusPersistenceWriteModelAuditPreviewModelsWriteShape TestProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusPersistenceWriteModelAuditPreviewFailsClosedWithoutSources TestProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusPersistenceWriteModelAuditPreviewCommand production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-persistence-write-model-audit-preview write_model_audit_required write_model_modeled persistence_authorization_audit_consumed redacted_write_model_guidance_consumed compatibility_center_write_model_modeled runtime_diagnostics_write_model_modeled write_model_boundary_ready redacted_summary_shape_modeled status_persistence_write_enabled kde_status_write_enabled runtime_diagnostics_write_enabled write_model_item_count ready_write_model_item_count missing_write_model_item_count write_enabled_item_count persisted_status_item_count consumer_enabled_write_model_item_count raw_exposed_write_model_item_count side_effect_write_model_item_count kde_status_persisted runtime_diagnostics_persisted state_root_path_exposed file_paths_exposed file_content_read].each do |token|
  assert(go_runtime_production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_kde_safe_redacted_status_persistence_write_model_audit_test_source.include?(token), "Go Runtime production receipt notification action dry-run result lookup consumer enablement KDE-safe redacted status persistence write-model audit tests and CLI must include #{token}")
end

go_runtime_production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_kde_safe_redacted_status_writer_authorization_gate_audit_source = read_project_file("internal/runtime/owner/production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_kde_safe_redacted_status_writer_authorization_gate_audit.go")
%w[ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAuthorizationGateAuditPreview ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAuthorizationGateAuditItem ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAuthorizationGateAuditCheck NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAuthorizationGateAuditPreview xnix.runtime.production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_kde_safe_redacted_status_writer_authorization_gate_audit.v1 production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-writer-authorization-gate-audit-preview receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-writer-authorization-gate-audit production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-writer-authorization-gate-audit-ready-writer-disabled production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-writer-authorization-gate-audit-blocked current-mainline-consumed redacted-write-model-audit-consumed writer-authorization-gate-modeled compatibility-center-and-runtime-writers-modeled ten-writer-items-ready-writer-disabled writer-grants-and-persistence-disabled consumer-lookup-request-and-support-disabled production-and-host-boundary-closed review-receipt-dry-run-result-lookup-consumer-enablement-kde-safe-compatibility-center-redacted-status-writer-authorization review-receipt-dry-run-result-lookup-consumer-enablement-kde-safe-runtime-diagnostics-redacted-status-writer-authorization renew-receipt-dry-run-result-lookup-consumer-enablement-kde-safe-compatibility-center-redacted-status-writer-authorization renew-receipt-dry-run-result-lookup-consumer-enablement-kde-safe-runtime-diagnostics-redacted-status-writer-authorization open-compatibility-center-dry-run-result-lookup-consumer-enablement-kde-safe-compatibility-center-redacted-status-writer-authorization open-compatibility-center-dry-run-result-lookup-consumer-enablement-kde-safe-runtime-diagnostics-redacted-status-writer-authorization dismiss-receipt-dry-run-result-lookup-consumer-enablement-kde-safe-compatibility-center-redacted-status-writer-authorization dismiss-receipt-dry-run-result-lookup-consumer-enablement-kde-safe-runtime-diagnostics-redacted-status-writer-authorization support-info-dry-run-result-lookup-consumer-enablement-kde-safe-compatibility-center-redacted-status-writer-authorization support-info-dry-run-result-lookup-consumer-enablement-kde-safe-runtime-diagnostics-redacted-status-writer-authorization].each do |token|
  assert(go_runtime_production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_kde_safe_redacted_status_writer_authorization_gate_audit_source.include?(token), "Go Runtime production receipt notification action dry-run result lookup consumer enablement KDE-safe redacted status writer authorization gate audit must include #{token}")
end
%w[CurrentMainlineConsumed RedactedWriteModelAuditConsumed RedactedWriteModelBoundaryReady WriterAuthorizationGateRequired WriterAuthorizationGateModeled WriterAuthorizationBoundaryReady CompatibilityCenterWriterModeled RuntimeDiagnosticsWriterModeled KDESafeRedactedStatusOnly OpaqueResultIDSupported WriterAuthorizationGranted StatusWriterEnabled StatusPersistenceWriteEnabled KDEStatusWriteEnabled RuntimeDiagnosticsWriteEnabled WriterItemCount RequiredWriterItemCount ReadyWriterItemCount MissingWriterItemCount CompatibilityCenterWriterItemCount RuntimeDiagnosticsWriterItemCount GrantedWriterItemCount EnabledWriterItemCount PersistedWriterItemCount RawExposedWriterItemCount SideEffectWriterItemCount ConsumerConsumptionAuthorized ConsumerEnablementAuthorized KDEConsumerEnabled RuntimeConsumerEnabled LookupRouteAuthorized LookupRouteEnabled OpaqueLookupEnabled RedactedSummaryPersisted KDEStatusPersisted RuntimeDiagnosticsPersisted DryRunResultPersisted RawResultExposed DispatchDryRunExecuted RequestObjectCreationEnabled RequestObjectDispatchEnabled PortalRequestCreated NotificationActionEnabled CompatibilityCenterOpened ProductionReadiness ProductionOwnershipReady SystemServiceStarted ProductionBusClaimed WriteMethodsEnabled RuntimeWritesEnabled DesktopFilesWritten SupportBundleExported SnapshotRestoreExecuted StateCleanupExecuted FileContentRead FilePathsExposed NetworkRequired HostRootModified PrivilegedContainerRequired StateRootPathExposed RawCommandExposed RawExecutableExposed BackendDetailsExposed validateNoBackendTerms].each do |token|
  assert(go_runtime_production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_kde_safe_redacted_status_writer_authorization_gate_audit_source.include?(token), "Go Runtime production receipt notification action dry-run result lookup consumer enablement KDE-safe redacted status writer authorization gate audit must expose safety gate #{token}")
end
%w[docs/xnix-current-mainline.md production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_kde_safe_redacted_status_persistence_write_model_audit.go].each do |token|
  assert(go_runtime_production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_kde_safe_redacted_status_writer_authorization_gate_audit_source.include?(token), "Go Runtime production receipt notification action dry-run result lookup consumer enablement KDE-safe redacted status writer authorization gate audit must consume #{token}")
end
go_runtime_production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_kde_safe_redacted_status_writer_authorization_gate_audit_test_source = read_project_file("internal/runtime/owner/production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_kde_safe_redacted_status_writer_authorization_gate_audit_test.go") +
                                                                                                                                                       read_project_file("cmd/xnix-runtime-go/production_dbus_gate_cli_test.go") +
                                                                                                                                                       read_project_file("cmd/xnix-runtime-go/production_dbus_gate_commands.go") +
                                                                                                                                                       read_project_file("cmd/xnix-runtime-go/main.go")
%w[TestProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAuthorizationGateAuditPreviewModelsGate TestProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAuthorizationGateAuditPreviewFailsClosedWithoutSources TestProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAuthorizationGateAuditPreviewCommand production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-writer-authorization-gate-audit-preview current_mainline_consumed redacted_write_model_audit_consumed redacted_write_model_boundary_ready writer_authorization_gate_required writer_authorization_gate_modeled writer_authorization_boundary_ready compatibility_center_writer_modeled runtime_diagnostics_writer_modeled writer_authorization_granted status_writer_enabled status_persistence_write_enabled kde_status_write_enabled runtime_diagnostics_write_enabled writer_item_count ready_writer_item_count missing_writer_item_count granted_writer_item_count enabled_writer_item_count persisted_writer_item_count raw_exposed_writer_item_count side_effect_writer_item_count kde_status_persisted runtime_diagnostics_persisted state_root_path_exposed file_paths_exposed file_content_read].each do |token|
  assert(go_runtime_production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_kde_safe_redacted_status_writer_authorization_gate_audit_test_source.include?(token), "Go Runtime production receipt notification action dry-run result lookup consumer enablement KDE-safe redacted status writer authorization gate audit tests and CLI must include #{token}")
end

xnix_current_mainline_source = read_project_file("docs/xnix-current-mainline.md")
%w[Xnix\ Current\ Mainline Codex-owned\ current\ mainline Product\ Goal Ownership\ Rules First-Release\ KDE\ Entrypoints Current\ Safe\ Next\ Task Safety\ Gates KDE\ Plasma\ is\ the\ only\ flagship\ desktop Runtime\ business\ logic\ is\ Go-first Ruby\ is\ for\ tests C\ remains\ for\ low-level redacted\ status\ writer\ authorization\ gate\ audit].each do |token|
  assert(xnix_current_mainline_source.include?(token.tr("\\", "")), "Xnix current mainline must include #{token}")
end
go_runtime_current_mainline_ownership_audit_source = read_project_file("internal/runtime/owner/current_mainline_ownership_audit.go")
%w[CurrentMainlineOwnershipAuditPreview CurrentMainlineOwnershipAuditCheck CurrentMainlineOwnershipAuditCounts NewCurrentMainlineOwnershipAuditPreview xnix.runtime.current_mainline_ownership_audit.v1 current-mainline-ownership-audit-preview current-mainline-ownership-audit current-mainline-ownership-audit-ready-autonomous-mainline current-mainline-ownership-audit-blocked xnix-current-mainline current-mainline-document-present external-agent-dependency-removed kde-flagship-only runtime-ownership-language-split seven-kde-entrypoints-present safe-next-task-present full-gate-authorization-required unsafe-runtime-and-host-gates-closed docs/xnix-current-mainline.md].each do |token|
  assert(go_runtime_current_mainline_ownership_audit_source.include?(token), "Go Runtime current mainline ownership audit must include #{token}")
end
%w[CurrentMainlineDocumentPresent CodexOwnedMainline ExternalAgentDependencyRequired HistoricalExternalAgentDocumentsAllowed KDEFlagshipOnly GNOMEFirstReleaseTarget XFCEFirstReleaseTarget GoRuntimeOwned RubyTestHarnessOnly CLowLevelOnly KDEPluginsPresentationOnly RuntimeOwnsCompatibilityDecisions SevenKDEEntrypointsPresent SafeNextTaskPresent FullGateRequiresAuthorization DockerExecuted QEMUExecuted NetworkFetchEnabled PackageManagerInvoked ProductionBusClaimed WriteMethodsEnabled RuntimeWritesEnabled KDEConfigurationWritten PortalCallExecuted BackendLaunchEnabled BackendProcessStarted HostRootModified PrivilegedContainerRequired StateRootPathExposed FilePathsExposed FileContentRead RawCommandExposed RawExecutableExposed BackendDetailsExposed validateNoBackendTerms].each do |token|
  assert(go_runtime_current_mainline_ownership_audit_source.include?(token), "Go Runtime current mainline ownership audit must expose safety gate #{token}")
end
go_runtime_current_mainline_ownership_audit_test_source = read_project_file("internal/runtime/owner/current_mainline_ownership_audit_test.go") +
                                                        read_project_file("cmd/xnix-runtime-go/current_mainline_cli_test.go") +
                                                        read_project_file("cmd/xnix-runtime-go/current_mainline_commands.go") +
                                                        read_project_file("cmd/xnix-runtime-go/main.go")
%w[TestCurrentMainlineOwnershipAuditPreviewModelsAutonomousMainline TestCurrentMainlineOwnershipAuditPreviewFailsClosedWithoutMainlineDocument TestCurrentMainlineOwnershipAuditPreviewCommand current-mainline-ownership-audit-preview current_mainline_document_present codex_owned_mainline external_agent_dependency_required kde_flagship_only go_runtime_owned seven_kde_entrypoints_present full_gate_requires_authorization docker_executed qemu_executed network_fetch_enabled package_manager_invoked production_bus_claimed write_methods_enabled runtime_writes_enabled kde_configuration_written portal_call_executed backend_launch_enabled host_root_modified state_root_path_exposed file_paths_exposed file_content_read].each do |token|
  assert(go_runtime_current_mainline_ownership_audit_test_source.include?(token), "Go Runtime current mainline ownership audit tests and CLI must include #{token}")
end

go_runtime_production_dbus_method_review_source = read_project_file("internal/runtime/owner/production_dbus_method_review.go")
%w[ProductionDBusMethodReviewPreview ProductionDBusMethodReviewRoute ProductionDBusMethodReviewCheck NewProductionDBusMethodReviewPreview xnix.runtime.production_dbus_method_review.v1 production-dbus-method-review-preview route-by-route-production-dbus-method-review runtime-method-parity-manifest-preview runtime-owner-route-manifest-preview production-dbus-gate-review-preview runtime-write-gate-preview production-dbus-method-review-ready-production-exposure-disabled production-dbus-method-review-blocked dbus-read-only-contract owner-local-candidate reserved-write-method reviewed-read-only-contract-production-owner-disabled owner-local-only-no-production-dbus-method write-method-disabled-no-production-dispatch read-only-contract-methods-reviewed owner-local-candidates-reviewed write-methods-reviewed-disabled no-new-production-methods-requested production-exposure-disabled route-manifest-consumed unsafe-gates-closed].each do |token|
  assert(go_runtime_production_dbus_method_review_source.include?(token), "Go Runtime production D-Bus method review must include #{token}")
end
%w[ReadOnlyContractMethodCount OwnerLocalCandidateCount WriteMethodCount ReviewedMethodCount ProductionExposureReadyCount NewProductionMethodRequestCount SystemServiceStarted SessionBusClaimed ProductionBusClaimed ProductionOwnerEnabled ProductionActivationReady WriteMethodsEnabled RuntimeWritesEnabled AdapterInvocationEnabled BackendLaunchEnabled BackendProcessStarted NotificationSent NetworkRequired HostRootModified PrivilegedContainerRequired StateRootPathExposed RawCommandExposed RawExecutableExposed BackendDetailsExposed validateNoBackendTerms].each do |token|
  assert(go_runtime_production_dbus_method_review_source.include?(token), "Go Runtime production D-Bus method review must expose safety gate #{token}")
end
%w[NewRuntimeOwnerRouteManifestPreview NewRuntimeMethodParityManifestPreview NewProductionDBusGateReviewPreview NewRuntimeWriteGatePreview].each do |token|
  assert(go_runtime_production_dbus_method_review_source.include?(token), "Go Runtime production D-Bus method review must consume #{token}")
end
go_runtime_production_dbus_method_review_test_source = read_project_file("internal/runtime/owner/production_dbus_method_review_test.go") +
                                                      read_project_file("cmd/xnix-runtime-go/production_dbus_gate_cli_test.go") +
                                                      read_project_file("cmd/xnix-runtime-go/production_dbus_gate_commands.go") +
                                                      read_project_file("cmd/xnix-runtime-go/main.go")
%w[TestProductionDBusMethodReviewPreviewInventoriesRoutes TestProductionDBusMethodReviewPreviewFailsClosedWithoutSources TestProductionDBusMethodReviewPreviewCommand production-dbus-method-review-preview read_only_contract_method_count owner_local_candidate_count write_method_count reviewed_method_count production_exposure_ready_count new_production_method_request_count].each do |token|
  assert(go_runtime_production_dbus_method_review_test_source.include?(token), "Go Runtime production D-Bus method review tests and CLI must include #{token}")
end

go_runtime_production_rollback_diagnostics_review_source = read_project_file("internal/runtime/owner/production_rollback_diagnostics_review.go")
%w[ProductionRollbackDiagnosticsReviewPreview ProductionRollbackDiagnosticsReviewItem ProductionRollbackDiagnosticsReviewCheck NewProductionRollbackDiagnosticsReviewPreview xnix.runtime.production_rollback_diagnostics_review.v1 production-rollback-diagnostics-review-preview production-dbus-rollback-diagnostics-review production-rollback-diagnostics-review-ready-side-effects-disabled production-rollback-diagnostics-review-blocked production-dbus-gate-review-preview production-dbus-method-review-preview runtime-service-activation-preflight-preview support-bundle-manifest-preview support-case-timeline-preview snapshot-restore-candidates-preview state-root-quota-retention-preview production-gate-consumed method-review-consumed rollback-controls-reviewed diagnostics-controls-reviewed support-side-effects-disabled restore-and-cleanup-disabled unsafe-data-hidden host-boundary-closed].each do |token|
  assert(go_runtime_production_rollback_diagnostics_review_source.include?(token), "Go Runtime production rollback diagnostics review must include #{token}")
end
%w[ReviewItemCount RollbackControlCount DiagnosticsControlCount ReadyControlCount SideEffectControlCount ProductionOwnershipReady SystemServiceStarted SessionBusClaimed ProductionBusClaimed ProductionOwnerEnabled ProductionActivationReady WriteMethodsEnabled RuntimeWritesEnabled AdapterInvocationEnabled BackendLaunchEnabled BackendProcessStarted SupportBundleExported SupportCaseCreated NotificationSent SnapshotRestoreExecuted StateCleanupExecuted FileContentRead FilePathsExposed NetworkRequired HostRootModified PrivilegedContainerRequired StateRootPathExposed RawCommandExposed RawExecutableExposed BackendDetailsExposed validateNoBackendTerms].each do |token|
  assert(go_runtime_production_rollback_diagnostics_review_source.include?(token), "Go Runtime production rollback diagnostics review must expose safety gate #{token}")
end
%w[production_dbus_gate_review.go production_dbus_method_review.go runtime_service_activation_preflight.go support_bundle_manifest.go support_case_timeline.go snapshot_restore_candidates.go state_root_quota_retention.go].each do |token|
  assert(go_runtime_production_rollback_diagnostics_review_source.include?(token), "Go Runtime production rollback diagnostics review must consume #{token}")
end
go_runtime_production_rollback_diagnostics_review_test_source = read_project_file("internal/runtime/owner/production_rollback_diagnostics_review_test.go") +
                                                               read_project_file("cmd/xnix-runtime-go/production_dbus_gate_cli_test.go") +
                                                               read_project_file("cmd/xnix-runtime-go/production_dbus_gate_commands.go") +
                                                               read_project_file("cmd/xnix-runtime-go/main.go")
%w[TestProductionRollbackDiagnosticsReviewPreviewConsumesSafetySurfaces TestProductionRollbackDiagnosticsReviewPreviewFailsClosedWithoutSources TestProductionRollbackDiagnosticsReviewPreviewCommand production-rollback-diagnostics-review-preview review_item_count rollback_control_count diagnostics_control_count ready_control_count side_effect_control_count production_ownership_ready snapshot_restore_executed state_cleanup_executed].each do |token|
  assert(go_runtime_production_rollback_diagnostics_review_test_source.include?(token), "Go Runtime production rollback diagnostics review tests and CLI must include #{token}")
end

go_runtime_production_desktop_side_effect_review_source = read_project_file("internal/runtime/owner/production_desktop_side_effect_review.go")
%w[ProductionDesktopSideEffectReviewPreview ProductionDesktopSideEffectSurface ProductionDesktopSideEffectReviewCheck NewProductionDesktopSideEffectReviewPreview xnix.runtime.production_desktop_side_effect_review.v1 production-desktop-side-effect-review-preview kde-production-desktop-side-effect-review production-desktop-side-effect-review-ready-side-effects-disabled production-desktop-side-effect-review-blocked production-dbus-gate-review-preview production-rollback-diagnostics-review-preview kde-entrypoints-preview kde-action-queue-preview desktop-activation-manifest-preview kde-shell-integration-preview kde-application-surface-preview production-gate-consumed rollback-diagnostics-consumed seven-kde-surfaces-reviewed runtime-policy-owner desktop-writes-disabled live-shell-activation-disabled notifications-and-requests-disabled host-boundary-closed launcher task-manager file-manager system-tray notifications compatibility-center unified-settings].each do |token|
  assert(go_runtime_production_desktop_side_effect_review_source.include?(token), "Go Runtime production desktop side-effect review must include #{token}")
end
%w[SurfaceCount RequiredSurfaceCount ReviewedSurfaceCount ActiveSurfaceCount SideEffectSurfaceCount ProductionOwnershipReady PlasmaForkRequired PlasmaSourceModified SystemServiceStarted SessionBusClaimed ProductionBusClaimed ProductionOwnerEnabled ProductionActivationReady WriteMethodsEnabled RuntimeWritesEnabled DesktopFilesWritten MIMEAppsWritten ShellConfigurationWritten SettingsPersisted KRunnerIndexPersisted TaskManagerEntryActive KWinRuleApplied LiveTrayBridgeEnabled TrayBridgePersisted NotificationSent NotificationDeliveryEnabled CompatibilityCenterPersisted PortalRequestCreated RequestObjectsCreated AdapterInvocationEnabled BackendLaunchEnabled BackendProcessStarted FileContentRead FilePathsExposed NetworkRequired HostRootModified PrivilegedContainerRequired StateRootPathExposed RawCommandExposed RawExecutableExposed BackendDetailsExposed validateNoBackendTerms].each do |token|
  assert(go_runtime_production_desktop_side_effect_review_source.include?(token), "Go Runtime production desktop side-effect review must expose safety gate #{token}")
end
%w[production_dbus_gate_review.go production_rollback_diagnostics_review.go kde_entrypoints.go kde_action_queue.go desktop_activation_manifest.go kde_shell_surface.go].each do |token|
  assert(go_runtime_production_desktop_side_effect_review_source.include?(token), "Go Runtime production desktop side-effect review must consume #{token}")
end
go_runtime_production_desktop_side_effect_review_test_source = read_project_file("internal/runtime/owner/production_desktop_side_effect_review_test.go") +
                                                              read_project_file("cmd/xnix-runtime-go/production_dbus_gate_cli_test.go") +
                                                              read_project_file("cmd/xnix-runtime-go/production_dbus_gate_commands.go") +
                                                              read_project_file("cmd/xnix-runtime-go/main.go")
%w[TestProductionDesktopSideEffectReviewPreviewInventoriesKDESurfaces TestProductionDesktopSideEffectReviewPreviewFailsClosedWithoutSources TestProductionDesktopSideEffectReviewPreviewCommand production-desktop-side-effect-review-preview surface_count required_surface_count reviewed_surface_count active_surface_count side_effect_surface_count production_ownership_ready desktop_files_written settings_persisted notification_sent live_tray_bridge_enabled].each do |token|
  assert(go_runtime_production_desktop_side_effect_review_test_source.include?(token), "Go Runtime production desktop side-effect review tests and CLI must include #{token}")
end

go_restricted_owner_smoke_receipt_source = read_project_file("internal/runtime/owner/restricted_smoke_receipt.go")
%w[RestrictedOwnerSmokeReceipt restricted-owner-smoke-execution-receipt xnix.runtime.restricted_owner_smoke_receipt.v1 RecordRestrictedOwnerSmokeReceipt LoadRestrictedOwnerSmokeReceipt restricted-owner-smoke-receipt.json runtime-service-activation-preflight+runtime-owner-smoke-batch RestrictedOwnerSmokeMode RestrictedOwnerSmokeDirective authorize-restricted-owner-smoke restricted-owner-smoke-ready runtime-owner-smoke-batch-record StateRootWriteScope explicit-test-root-only ReceiptPersisted ReceiptReadBack ProductionActivationReady ProductionOwnerEnabled SystemServiceStarted SessionBusClaimed ProductionBusClaimed WriteMethodsEnabled BackendLaunchEnabled HostRootModified validateRestrictedOwnerSmokeReceipt].each do |token|
  assert(go_restricted_owner_smoke_receipt_source.include?(token), "Go Runtime restricted owner smoke receipt must include #{token}")
end
assert(go_restricted_owner_smoke_receipt_source.include?("validateNoBackendTerms"), "Go Runtime restricted owner smoke receipt must hide backend terms")
go_restricted_owner_smoke_receipt_fanout_source = read_project_file("internal/runtime/owner/restricted_smoke_receipt_fanout.go")
%w[RestrictedOwnerSmokeReceiptFanOutPreview RestrictedOwnerSmokeReceiptFanSurface NewRestrictedOwnerSmokeReceiptFanOutPreview xnix.runtime.restricted_owner_smoke_receipt_fanout.v1 restricted-owner-smoke-receipt-fanout-preview GetRestrictedOwnerSmokeReceiptFanOut GetRestrictedOwnerSmokeReceiptFanOutPreview restricted-smoke-receipt-ready runtime-owner-readiness service-activation-preflight compatibility-onboarding support-bundle-manifest support-case-timeline receipt-consumed readiness-surface-coverage support-surface-coverage surfaces-read-only ownership-boundary-closed support-side-effects-disabled unsafe-data-hidden host-boundary-closed StateRootWritesEnabled FanOutWritesEnabled SupportBundleExported SupportCaseCreated NotificationSent ProductionBusClaimed WriteMethodsEnabled BackendLaunchEnabled HostRootModified].each do |token|
  assert(go_restricted_owner_smoke_receipt_fanout_source.include?(token), "Go Runtime restricted owner smoke receipt fan-out must include #{token}")
end
%w[RestrictedOwnerSmokeReceiptFanOutOwnerRoutePreview NewRestrictedOwnerSmokeReceiptFanOutOwnerRoutePreview xnix.runtime.restricted_owner_smoke_receipt_fanout_owner_route.v1 restricted-owner-smoke-receipt-fanout-owner-route-preview owner-local-restricted-smoke-receipt-fanout opaque-lookup-consumed surface-fanout-deferred owner-route-ready-production-dbus-blocked missing-receipt-fail-closed OwnerLocalRouteCandidateReady ProductionDBusExposureReady RuntimeWritesEnabled].each do |token|
  assert(go_restricted_owner_smoke_receipt_fanout_source.include?(token), "Go Runtime restricted owner smoke receipt fan-out owner route must include #{token}")
end
assert(go_restricted_owner_smoke_receipt_fanout_source.include?("validateNoBackendTerms"), "Go Runtime restricted owner smoke receipt fan-out must hide backend terms")
%w[restricted-owner-smoke-receipt-record runRestrictedOwnerSmokeReceiptRecord authorize-restricted-owner-smoke].each do |token|
  assert(read_project_file("cmd/xnix-runtime-go/restricted_owner_smoke_receipt_commands.go").include?(token) || read_project_file("cmd/xnix-runtime-go/main.go").include?(token), "Go Runtime restricted owner smoke receipt CLI must include #{token}")
end
%w[restricted-owner-smoke-receipt-fanout-preview runRestrictedOwnerSmokeReceiptFanOutPreview NewRestrictedOwnerSmokeReceiptFanOutPreview].each do |token|
  assert(read_project_file("cmd/xnix-runtime-go/restricted_owner_smoke_receipt_commands.go").include?(token) || read_project_file("cmd/xnix-runtime-go/main.go").include?(token), "Go Runtime restricted owner smoke receipt fan-out CLI must include #{token}")
end
go_restricted_owner_smoke_receipt_lookup_source = read_project_file("internal/runtime/owner/restricted_smoke_receipt_lookup.go")
%w[RestrictedOwnerSmokeReceiptLookupPreview ResolveRestrictedOwnerSmokeReceipt xnix.runtime.restricted_owner_smoke_receipt_lookup.v1 restricted-owner-smoke-receipt-lookup-preview owner-managed-receipt-lookup GetRestrictedOwnerSmokeReceiptLookup GetRestrictedOwnerSmokeReceiptLookupPreview opaque_receipt_id restricted-owner-smoke-receipt-id missing-receipt opaque-receipt-id-supported caller-state-root-hidden receipt-slot-redacted missing-receipt-fails-closed StateRootWritesEnabled RuntimeWritesEnabled FanOutWritesEnabled ProductionBusClaimed WriteMethodsEnabled BackendLaunchEnabled HostRootModified].each do |token|
  assert(go_restricted_owner_smoke_receipt_lookup_source.include?(token), "Go Runtime restricted owner smoke receipt lookup must include #{token}")
end
assert(go_restricted_owner_smoke_receipt_lookup_source.include?("validateNoBackendTerms"), "Go Runtime restricted owner smoke receipt lookup must hide backend terms")
assert(read_project_file("internal/runtime/owner/restricted_smoke_receipt_test.go").include?("TestRecordRestrictedOwnerSmokeReceiptConsumesPreflightAndSmokeBatch"), "Go Runtime restricted owner smoke receipt tests must consume preflight and smoke batch")
assert(read_project_file("cmd/xnix-runtime-go/restricted_owner_smoke_receipt_cli_test.go").include?("TestRestrictedOwnerSmokeReceiptRecordCommandPersistsReceipt"), "Go Runtime restricted owner smoke receipt CLI tests must persist the receipt")
assert(read_project_file("internal/runtime/owner/restricted_smoke_receipt_fanout_test.go").include?("TestRestrictedOwnerSmokeReceiptFanOutPreviewCoversReadinessAndSupportSurfaces"), "Go Runtime restricted owner smoke receipt fan-out tests must cover readiness and support surfaces")
assert(read_project_file("cmd/xnix-runtime-go/restricted_owner_smoke_receipt_cli_test.go").include?("TestRestrictedOwnerSmokeReceiptFanOutPreviewCommandCoversReadinessAndSupportSurfaces"), "Go Runtime restricted owner smoke receipt fan-out CLI tests must cover readiness and support surfaces")
assert(read_project_file("internal/runtime/owner/restricted_smoke_receipt_lookup_test.go").include?("TestResolveRestrictedOwnerSmokeReceiptReturnsOpaqueLookup"), "Go Runtime restricted owner smoke receipt lookup tests must cover opaque lookup")
assert(read_project_file("cmd/xnix-runtime-go/restricted_owner_smoke_receipt_cli_test.go").include?("TestRestrictedOwnerSmokeReceiptLookupPreviewCommand"), "Go Runtime restricted owner smoke receipt lookup CLI tests must cover the lookup command")
assert(read_project_file("internal/runtime/owner/restricted_smoke_receipt_fanout_owner_route_test.go").include?("TestRestrictedOwnerSmokeReceiptFanOutOwnerRoutePreviewUsesOpaqueLookup"), "Go Runtime restricted owner smoke receipt fan-out owner route tests must cover opaque lookup")
assert(read_project_file("cmd/xnix-runtime-go/restricted_owner_smoke_receipt_cli_test.go").include?("TestRestrictedOwnerSmokeReceiptFanOutOwnerRoutePreviewCommand"), "Go Runtime restricted owner smoke receipt fan-out owner route CLI tests must cover the owner route command")

go_restricted_owner_smoke_receipt_fanout_owner_route_audit_source = read_project_file("internal/runtime/owner/restricted_smoke_receipt_fanout_owner_route_audit.go")
%w[RestrictedOwnerSmokeReceiptFanOutOwnerRouteAuditPreview NewRestrictedOwnerSmokeReceiptFanOutOwnerRouteAuditPreview xnix.runtime.restricted_owner_smoke_receipt_fanout_owner_route_audit.v1 restricted-owner-smoke-receipt-fanout-owner-route-audit-preview GetRestrictedOwnerSmokeReceiptFanOutOwnerRouteAudit GetRestrictedOwnerSmokeReceiptFanOutOwnerRouteAuditPreview restricted-owner-smoke-receipt-fanout-owner-route-audit GetRestrictedOwnerSmokeReceiptFanOut owner-local-route-ready owner-local-read-route-ready-production-dbus-blocked owner-local-route-smoke-covered owner-local-read-route-smoke-covered-production-dbus-blocked restricted-owner-smoke-fanout-owner-smoke-coverage owner-smoke-coverage-present cli-preview-registered go-read-model-present owner-route-present production-dbus-absent receipt-consumption-present caller-state-root-boundary owner-managed-receipt-lookup route-decision unsafe-gates-closed].each do |token|
  assert(go_restricted_owner_smoke_receipt_fanout_owner_route_audit_source.include?(token), "Go Runtime restricted owner smoke receipt fan-out owner-route audit must include #{token}")
end
%w[ConsumesExistingReceipt RequiresCallerStateRoot ReceiptLookupOwnerManaged OpaqueReceiptIDSupported OwnerSmokeCoverageReady FanOutWritesEnabled StateRootWritesEnabled SupportBundleExported SupportCaseCreated NotificationSent OwnerLocalRouteCandidateReady ProductionDBusExposureReady SystemServiceStarted SessionBusClaimed ProductionBusClaimed WriteMethodsEnabled RuntimeWritesEnabled BackendLaunchEnabled BackendProcessStarted NetworkRequired HostRootModified PrivilegedContainerRequired StateRootPathExposed BackendDetailsExposed].each do |token|
  assert(go_restricted_owner_smoke_receipt_fanout_owner_route_audit_source.include?(token), "Go Runtime restricted owner smoke receipt fan-out owner-route audit must expose safety gate #{token}")
end
assert(go_restricted_owner_smoke_receipt_fanout_owner_route_audit_source.include?("validateNoBackendTerms"), "Go Runtime restricted owner smoke receipt fan-out owner-route audit must hide backend terms")
go_restricted_owner_smoke_receipt_fanout_owner_route_audit_test_source = read_project_file("internal/runtime/owner/restricted_smoke_receipt_fanout_owner_route_audit_test.go")
%w[TestRestrictedOwnerSmokeReceiptFanOutOwnerRouteAuditRecognizesOpaqueLookup TestRestrictedOwnerSmokeReceiptFanOutOwnerRouteAuditFailsClosedWithoutSources owner-local-route-smoke-covered caller-state-root-boundary owner-managed-receipt-lookup owner-smoke-coverage-present unsafe-gates-closed].each do |token|
  assert(go_restricted_owner_smoke_receipt_fanout_owner_route_audit_test_source.include?(token), "Go Runtime restricted owner smoke receipt fan-out owner-route audit tests must include #{token}")
end
go_restricted_owner_smoke_receipt_fanout_owner_route_audit_cli_source = read_project_file("cmd/xnix-runtime-go/restricted_owner_smoke_receipt_commands.go") +
                                                                 read_project_file("cmd/xnix-runtime-go/restricted_owner_smoke_receipt_cli_test.go") +
                                                                 read_project_file("cmd/xnix-runtime-go/main.go")
%w[restricted-owner-smoke-receipt-fanout-owner-route-audit-preview runRestrictedOwnerSmokeReceiptFanOutOwnerRouteAuditPreview TestRestrictedOwnerSmokeReceiptFanOutOwnerRouteAuditPreviewCommand route_decision requires_caller_state_root receipt_lookup_owner_managed opaque_receipt_id_supported owner_smoke_coverage_ready].each do |token|
  assert(go_restricted_owner_smoke_receipt_fanout_owner_route_audit_cli_source.include?(token), "Go Runtime restricted owner smoke receipt fan-out owner-route audit CLI must include #{token}")
end
%w[restricted-owner-smoke-receipt-lookup-preview runRestrictedOwnerSmokeReceiptLookupPreview TestRestrictedOwnerSmokeReceiptLookupPreviewCommand owner_managed_lookup caller_state_root_required opaque_receipt_id].each do |token|
  assert(go_restricted_owner_smoke_receipt_fanout_owner_route_audit_cli_source.include?(token), "Go Runtime restricted owner smoke receipt lookup CLI must include #{token}")
end
%w[restricted-owner-smoke-receipt-fanout-owner-route-preview runRestrictedOwnerSmokeReceiptFanOutOwnerRoutePreview TestRestrictedOwnerSmokeReceiptFanOutOwnerRoutePreviewCommand fan_out_result_state owner_local_route_candidate_ready].each do |token|
  assert(go_restricted_owner_smoke_receipt_fanout_owner_route_audit_cli_source.include?(token), "Go Runtime restricted owner smoke receipt fan-out owner route CLI must include #{token}")
end

go_runtime_owner_session_bus_source = read_project_file("internal/runtime/owner/session_bus.go")
%w[SessionBusSmokeStep runtime-owner-session-bus-smoke-step xnix.runtime.owner_session_bus_smoke.v1 restricted-private-session-bus-owner-smoke NewSessionBusSmokeTranscript NewSmokeBatchRecords record.RecordType reject-unsupported-read org.xnix.Compatibility1.Error.UnsupportedMethod private-session-bus-smoke].each do |token|
  assert(go_runtime_owner_session_bus_source.include?(token), "Go Runtime owner private session-bus smoke must include #{token}")
end
%w[ReadOnlyDispatch WriteMethod UnsupportedRead RouteReady DispatchReady ReadDispatchMethodCount WriteMethodCount RuntimeOwned GoRuntimeBacked KDEPolicyOwner KDEMayClaimRuntimeOwnership PrivateSessionBus EventLoopStarted SessionBusClaimed ProductionBusClaimed SystemServiceStarted WriteMethodsEnabled NetworkRequired HostRootModified PrivilegedContainerRequired BackendDetailsExposed].each do |token|
  assert(go_runtime_owner_session_bus_source.include?(token), "Go Runtime owner private session-bus smoke must expose #{token}")
end
assert(go_runtime_owner_session_bus_source.include?("validateNoBackendTerms"), "Go Runtime owner private session-bus smoke must hide backend terms")

go_runtime_owner_commands_source = read_project_file("cmd/xnix-runtime-go/runtime_owner_commands.go")
%w[runRuntimeServiceBindingPreview runRuntimeServiceActivationPreflightPreview runRuntimeLiveOwnerGatePreview runRuntimeOwnerProcessPreview runRuntimeOwnerSmokePlanPreview runRuntimeMethodParityManifestPreview runRuntimeOwnerReadinessPreview runRuntimeOwnerRouteManifestPreview runRuntimeRouteConvergencePreview runRuntimeOwnerRecipeTrustPreview encodeIndentedJSON].each do |token|
  assert(go_runtime_owner_commands_source.include?(token), "Go Runtime owner commands must include #{token}")
end
assert(go_runtime_owner_commands_source.include?("runtime-service-activation-preflight-preview"), "Go Runtime owner commands must expose Runtime service activation preflight preview")
assert(go_runtime_owner_commands_source.include?("runtime-owner-process-preview"), "Go Runtime owner commands must expose Runtime owner process preview")
assert(go_runtime_owner_commands_source.include?("runtime-owner-route-manifest-preview"), "Go Runtime owner commands must expose Runtime owner route manifest preview")
assert(go_runtime_owner_commands_source.include?("runtime-route-convergence-preview"), "Go Runtime owner commands must expose Runtime route convergence preview")
%w[NewRuntimeServiceBindingPreview NewRuntimeServiceActivationPreflightPreview NewRuntimeLiveOwnerGatePreview NewRuntimeOwnerProcessPreview NewRuntimeOwnerSmokePlanPreview NewRuntimeMethodParityManifestPreview NewRuntimeOwnerReadinessPreview NewRuntimeOwnerRouteManifestPreview NewRuntimeRouteConvergencePreview NewRuntimeWriteGatePreview NewRuntimeOwnerRecipeTrustPreview].each do |token|
  assert(go_runtime_owner_commands_source.include?(token), "Go Runtime owner commands must call #{token}")
end

go_runtime_owner_route_manifest_source = read_project_file("internal/runtime/appidentity/runtime_owner_route_manifest.go")
%w[RuntimeOwnerRouteManifestPreview runtime-owner-route-manifest-preview xnix.runtime.owner_route_manifest.v1 GetRuntimeOwnerRouteManifest GetRuntimeOwnerRouteManifestPreview runtime-method-parity-manifest-preview+go-runtime-cli+c-runtime-core+runtime-dispatch applications-preview application-preview engine-catalog-preview run-plan-preview desktop-activation-manifest-preview state-root-preview snapshot-plan-preview portal-access-policy-preview kde-integration-status-preview kde-shell-integration-preview kde-application-surface-preview task-manager-identity-preview kwin-window-rule-preview compatibility-install-preview runtime-write-gate-preview ai-diagnostic-input-preview ai-diagnostic-recommendation-preview ai-repair-approval-gate-preview diagnostics-preview go-runtime-cli c-runtime-core ruby-runtime-dispatch go-owner-native go-owner-c-adapter go-preview-ready c-adapter-pending legacy-dispatch-pending method-parity go-route-coverage c-core-adapter-boundary ruby-legacy-dispatch write-route-gate host-safety-boundary].each do |token|
  assert(go_runtime_owner_route_manifest_source.include?(token), "Go Runtime owner route manifest preview must include #{token}")
end
%w[RouteCounts MethodParityReady GoOwnerRouteCoverageReady CCoreAdapterRequired LegacyRuntimeRoutesPresent ProductionOwnerRoutesReady RuntimeOwned GoRuntimeBacked KDEPolicyOwner KDEMayClaimRuntimeOwnership SystemServiceStarted ProductionBusClaimed WriteMethodsEnabled NetworkRequired HostRootModified PrivilegedContainerRequired BackendDetailsExposed].each do |token|
  assert(go_runtime_owner_route_manifest_source.include?(token), "Go Runtime owner route manifest preview must expose #{token}")
end
assert(go_runtime_owner_route_manifest_source.include?("runtimeOwnerRouteGoCommands"), "Go Runtime owner route manifest must classify Go routes")
assert(go_runtime_owner_route_manifest_source.include?("runtimeOwnerRouteCCoreCommands"), "Go Runtime owner route manifest must classify C Runtime routes")
assert(go_runtime_owner_route_manifest_source.include?("NewRuntimeMethodParityManifestPreview"), "Go Runtime owner route manifest must derive from method parity")
assert(go_runtime_owner_route_manifest_source.include?("validateNoBackendTerms"), "Go Runtime owner route manifest must hide backend terms")

go_runtime_route_convergence_source = read_project_file("internal/runtime/appidentity/runtime_route_convergence.go")
%w[RuntimeRouteConvergencePreview runtime-route-convergence-preview xnix.runtime.route_convergence.v1 runtime-route-convergence GetRuntimeOwnerRouteManifest GetRuntimeRouteConvergencePreview runtime-owner-route-manifest-preview+runtime-method-parity-manifest-preview go-product-logic c-policy-bridge ruby-smoke-bridge fixture-only contract-only deprecated unsupported unclassified owner-route-manifest-ready all-routes-classified native-go-route-coverage write-gate-disabled host-safety-boundary].each do |token|
  assert(go_runtime_route_convergence_source.include?(token), "Go Runtime route convergence preview must include #{token}")
end
%w[ClassificationCounts MigrationGroups AllRoutesClassified NativeGoCoverageReady ProductionOwnerReady RuntimeOwned GoRuntimeBacked KDEPolicyOwner WriteMethodsEnabled NetworkRequired HostRootModified PrivilegedContainerRequired BackendLaunchEnabled BackendDetailsExposed UnclassifiedRoutes NextMigrationOrder].each do |token|
  assert(go_runtime_route_convergence_source.include?(token), "Go Runtime route convergence preview must expose #{token}")
end
assert(go_runtime_route_convergence_source.include?("NewRuntimeOwnerRouteManifestPreview"), "Go Runtime route convergence preview must derive from owner route manifest")
assert(go_runtime_route_convergence_source.include?("validateNoBackendTerms"), "Go Runtime route convergence preview must hide backend terms")

go_kde_action_dependency_graph_source = read_project_file("internal/runtime/appidentity/kde_action_dependency_graph.go")
%w[KDEActionDependencyGraphPreview kde-action-dependency-graph-preview xnix.runtime.kde_action_dependency_graph.v1 compatibility-center-action-dependency-graph GetKDEActionDependencyGraph GetKDEActionDependencyGraphPreview kde-action-queue-preview runtime-evidence-prerequisites runtime-write-gate-preview application-readiness action-review-receipt portal-file-access-receipt settings-change-review review-flow RejectsMismatchedAppID RejectsMalformedOperation RejectsPathEscapeEvidence RejectsUnsafeSideEffects DependencyGraphPersisted PermissionGrantCreated RequestObjectsCreated RuntimeLaunchApproval LaunchAllowed ExecutionStarted BackendProcessStarted StateRootPathExposed RawCommandExposed FileContentRead BackendDetailsExposed].each do |token|
  assert(go_kde_action_dependency_graph_source.include?(token), "Go Runtime KDE action dependency graph preview must include #{token}")
end
%w[SettingsPersisted HostRootModified NetworkRequired validateNoBackendTerms].each do |token|
  assert(go_kde_action_dependency_graph_source.include?(token), "Go Runtime KDE action dependency graph preview must expose safety token #{token}")
end

go_kde_action_dependency_graph_test_source = read_project_file("internal/runtime/appidentity/kde_action_dependency_graph_test.go")
%w[TestKDEActionDependencyGraphPreviewMapsActionEvidence TestKDEActionDependencyGraphPreviewValidationAndSafeText MissingEvidenceCount BlockedActionCount portal-file-access-receipt runtime-write-gate-preview RejectsMismatchedAppID RejectsMalformedOperation RejectsPathEscapeEvidence RejectsUnsafeSideEffects validateNoBackendTerms].each do |token|
  assert(go_kde_action_dependency_graph_test_source.include?(token), "Go Runtime KDE action dependency graph tests must include #{token}")
end

go_kde_action_dependency_graph_cli_test_source = read_project_file("cmd/xnix-runtime-go/kde_action_dependency_graph_cli_test.go")
%w[TestKDEActionDependencyGraphPreviewCommand kde-action-dependency-graph-preview xnix.runtime.kde_action_dependency_graph.v1 GetKDEActionDependencyGraph GetKDEActionDependencyGraphPreview permission_grant_created state_root_path_exposed raw_command_exposed file_content_read].each do |token|
  assert(go_kde_action_dependency_graph_cli_test_source.include?(token), "Go Runtime KDE action dependency graph CLI tests must include #{token}")
end

go_runtime_cli_source = read_project_file("cmd/xnix-runtime-go/main.go")
assert(go_runtime_cli_source.include?("kde-action-dependency-graph-preview"), "Go Runtime CLI must expose KDE action dependency graph preview")
assert(go_runtime_cli_source.include?("runKDEActionDependencyGraphPreview"), "Go Runtime CLI must call KDE action dependency graph preview")

go_runtime_kde_journey_source = read_project_file("internal/runtime/appidentity/kde_journey_evidence.go")
%w[KDEJourneyEvidencePreview kde-journey-evidence-preview xnix.runtime.kde_journey_evidence.v1 kde-seven-entrypoint-runtime-evidence GetKDEJourneyEvidence GetKDEJourneyEvidencePreview application-readiness-preview kde-action-dependency-graph-preview task-manager-identity-preview kwin-window-rule-preview tray-status-preview notification-preview kde-center-page-preview settings-preview JourneyEvidencePersisted RuntimeWriteMethodsEnabled RequestObjectsCreated PermissionGrantCreated SettingsPersisted LaunchEnabled ExecutionStarted KWinRuleApplied TrayBridgeActivated NotificationSent HostRootModified StateRootPathExposed RawExecutableExposed RawCommandExposed FileContentRead BackendDetailsExposed].each do |token|
  assert(go_runtime_kde_journey_source.include?(token), "Go Runtime KDE journey evidence preview must include #{token}")
end

go_runtime_kde_journey_test_source = read_project_file("internal/runtime/appidentity/kde_journey_evidence_test.go")
[
  "TestKDEJourneyEvidencePreviewStitchesSevenEntryPoints",
  "TestKDEJourneyEvidencePreviewRejectsMalformedInput",
  "launcher",
  "task-manager",
  "file-manager",
  "system-tray",
  "notification-center",
  "ai-compatibility-center",
  "unified-settings",
  "MissingEvidenceCount",
  "BlockedActionCount",
  "program files"
].each do |token|
  assert(go_runtime_kde_journey_test_source.include?(token), "Go Runtime KDE journey evidence tests must include #{token}")
end

go_runtime_kde_journey_cli_test_source = read_project_file("cmd/xnix-runtime-go/kde_journey_evidence_cli_test.go")
%w[TestKDEJourneyEvidencePreviewCommand kde-journey-evidence-preview xnix.runtime.kde_journey_evidence.v1 GetKDEJourneyEvidence GetKDEJourneyEvidencePreview entry_point_count missing_evidence_count runtime_write_methods_enabled request_objects_created permission_grant_created state_root_path_exposed raw_command_exposed].each do |token|
  assert(go_runtime_kde_journey_cli_test_source.include?(token), "Go Runtime KDE journey evidence CLI tests must include #{token}")
end

assert(go_runtime_cli_source.include?("kde-journey-evidence-preview"), "Go Runtime CLI must expose KDE journey evidence preview")
assert(go_runtime_cli_source.include?("runKDEJourneyEvidencePreview"), "Go Runtime CLI must call KDE journey evidence preview")

go_runtime_desktop_activation_manifest_source = read_project_file("internal/runtime/appidentity/desktop_activation_manifest.go")
%w[DesktopActivationManifestPreview DesktopActivationManifestBundle desktop-activation-manifest-preview xnix.runtime.desktop_activation_manifest.v1 GetDesktopActivationManifest GetDesktopActivationManifestPreview desktop-activation-bundle-preview kde-entrypoints-preview launcher task-manager file-manager system-tray notifications compatibility-center unified-settings].each do |token|
  assert(go_runtime_desktop_activation_manifest_source.include?(token), "Go Runtime desktop activation manifest preview must include #{token}")
end
%w[RuntimeOwned GoRuntimeBacked KDEPolicyOwner UserVisible OfficialDesktopOnly StableDesktopContract NormalApplicationSurface StandardDesktopEntry SevenEntryPointContract PortalMediatedFileAccess DesktopFilesWritten MIMEAppsWritten ManifestWritten SettingsPersisted NotificationsSent TaskManagerEntryActive KWinRuleApplied LiveTrayBridgeEnabled LaunchEnabled BackendLaunchEnabled ExecutionStarted HostRootModified NetworkRequired PrivilegedContainerRequired BackendDetailsExposed RawWindowsExecutableExposed CompatibilityStorageExposed].each do |token|
  assert(go_runtime_desktop_activation_manifest_source.include?(token), "Go Runtime desktop activation manifest preview must expose safety flag #{token}")
end
assert(go_runtime_desktop_activation_manifest_source.include?("validateNoBackendTerms"), "Go Runtime desktop activation manifest preview must hide backend terms")

go_runtime_desktop_activation_manifest_cli_source = read_project_file("cmd/xnix-runtime-go/desktop_activation_manifest_commands.go")
%w[runDesktopActivationManifestPreview desktop-activation-manifest-preview DesktopActivationManifestPreview].each do |token|
  assert(go_runtime_desktop_activation_manifest_cli_source.include?(token), "Go Runtime desktop activation manifest CLI must include #{token}")
end

go_runtime_activation_stage_source = read_project_file("internal/runtime/activation/stage.go")
%w[StageRequest StageResult StagedFile desktop-activation-stage xnix.runtime.desktop_activation_stage.v1 kde-desktop-activation-test-root-stage go-runtime-controlled-staging-writer usr/share/applications usr/share/kio/servicemenus usr/share/xnix/compatibility/manifests usr/share/xnix/compatibility/activation-receipts xnix-open-with-compatibility.desktop RenderDesktopEntry xnix-compat-open requires_matching_sha256 refusing to stage desktop activation into filesystem root refusing to overwrite existing staged file].each do |token|
  assert(go_runtime_activation_stage_source.include?(token), "Go Runtime activation stage writer must include #{token}")
end
%w[RuntimeOwned GoRuntimeBacked KDEPolicyOwner StagingRootRequired StagingRootPathExposed HostRootAllowed FileWritesPerformed DesktopFilesWritten MIMEAppsWritten ManifestWritten ReceiptWritten RollbackReceiptWritten SettingsPersisted NotificationsSent TaskManagerEntryActive KWinRuleApplied LiveTrayBridgeEnabled LaunchEnabled BackendLaunchEnabled ExecutionStarted HostRootModified NetworkRequired PrivilegedContainerRequired BackendDetailsExposed].each do |token|
  assert(go_runtime_activation_stage_source.include?(token), "Go Runtime activation stage writer must expose safety flag #{token}")
end

go_runtime_activation_transaction_source = read_project_file("internal/runtime/appidentity/desktop_activation_transaction.go")
%w[DesktopActivationTransactionPreview desktop-activation-transaction-preview xnix.runtime.desktop_activation_transaction.v1 GetDesktopActivationTransactionPreview ReceiptEvidence CommitReceiptPlanned RollbackReceiptPlanned CommitAvailable RollbackAvailable CommitReceiptWritten RollbackReceiptWritten DigestGateReady HostRootModified BackendDetailsExposed].each do |token|
  assert(go_runtime_activation_transaction_source.include?(token), "Go Runtime activation transaction preview must include #{token}")
end

go_runtime_activation_stage_cli_source = read_project_file("cmd/xnix-runtime-go/desktop_activation_stage_commands.go")
%w[runDesktopActivationStage desktop-activation-stage --staging-root activation.Stage NewPlanWithProvenance].each do |token|
  assert(go_runtime_activation_stage_cli_source.include?(token), "Go Runtime activation stage CLI must include #{token}")
end

go_runtime_activation_stage_cli_test_source = read_project_file("cmd/xnix-runtime-go/desktop_activation_stage_cli_test.go")
%w[desktop-activation-stage staging_root_path_exposed host_root_modified backend_details_exposed xnix-org.example.ledger.desktop requires_matching_sha256].each do |token|
  assert(go_runtime_activation_stage_cli_test_source.include?(token), "Go Runtime activation stage CLI test must include #{token}")
end

go_runtime_window_identity_source = read_project_file("internal/runtime/appidentity/window_identity_routes.go")
%w[TaskManagerIdentityPlanPreview KWinWindowRulePlanPreview task-manager-identity-preview kwin-window-rule-preview xnix.runtime.task_manager_identity.v1 xnix.runtime.kwin_window_rule.v1 GetTaskManagerIdentityPlan GetTaskManagerIdentityPlanPreview GetKWinWindowRulePlan GetKWinWindowRulePlanPreview window-identity-preview execution-session-record ExecutionSessionBacked ExecutionSessionPath].each do |token|
  assert(go_runtime_window_identity_source.include?(token), "Go Runtime window identity route previews must include #{token}")
end
%w[RuntimeOwned GoRuntimeBacked KDEPolicyOwner PinningAllowed RestoreAllowed PreferExistingWindow SkipTaskbar ShowInSwitcher TaskManagerEntryActive WindowObservationStarted LaunchEnabled ExecutionStarted HostRootModified BackendDetailsExposed WindowManagerPolicyOnly RuntimeOwnsBackendPolicy KWinRuleApplied ExecutionSessionRoot ExecutionSessionBacked].each do |token|
  assert(go_runtime_window_identity_source.include?(token), "Go Runtime window identity route previews must expose safety flag #{token}")
end
assert(go_runtime_window_identity_source.include?("validateNoBackendTerms"), "Go Runtime window identity route previews must hide backend terms")

go_runtime_execution_session_evidence_source = read_project_file("internal/runtime/appidentity/execution_session_record_evidence.go")
%w[ExecutionSessionRecordEvidence ExecutionSessionFanOutEvidence xnix.runtime.execution_session_record.v1 xnix.runtime.session_fanout.v1 execution-session-status-record execution-session-fanout-evidence GetExecutionSessionFanOutEvidence task-manager kwin tray compatibility-center SafeForKDE state_root_path_exposed status_persisted session_active live_state_observed window_observed task_manager_entry_active kwin_rule_applied live_tray_bridge_enabled launch_enabled execution_started backend_process_started host_root_modified backend_details_exposed].each do |token|
  assert(go_runtime_execution_session_evidence_source.include?(token), "Go Runtime execution session evidence must include #{token}")
end

go_runtime_window_identity_cli_source = read_project_file("cmd/xnix-runtime-go/window_identity_commands.go")
%w[runTaskManagerIdentityPreview runKWinWindowRulePreview task-manager-identity-preview kwin-window-rule-preview TaskManagerIdentityPlanPreview KWinWindowRulePlanPreview session-root session-request-id].each do |token|
  assert(go_runtime_window_identity_cli_source.include?(token), "Go Runtime window identity CLI must include #{token}")
end

go_runtime_window_identity_cli_test_source = read_project_file("cmd/xnix-runtime-go/window_identity_cli_test.go")
%w[TestTaskManagerIdentityPreviewCommandConsumesSessionRoot TestKWinWindowRulePreviewCommandConsumesSessionRoot execution_session_backed execution_session_path].each do |token|
  assert(go_runtime_window_identity_cli_test_source.include?(token), "Go Runtime window identity CLI tests must include #{token}")
end

go_runtime_tray_source = read_project_file("internal/runtime/appidentity/identity.go")
%w[TrayStatusOptions ExecutionSessionBacked ExecutionSessionPath ExecutionSessionFanOutEvidence CompatibilityCenter State compatibilityState].each do |token|
  assert(go_runtime_tray_source.include?(token), "Go Runtime tray status preview must include #{token}")
end

go_runtime_kde_center_page_source = read_project_file("internal/runtime/appidentity/kde_center_page.go")
%w[KDECenterPageOptions ExecutionSessionRoot ExecutionSessionBacked ExecutionSessionPath TaskManagerSessionState KWinSessionState TrayStatusPreviewWithOptions TaskManagerIdentityPlanPreviewWithOptions KWinWindowRulePlanPreviewWithOptions KDECenterPageActionGraph ActionDependencyGraph kde-action-dependency-graph-preview DependencyGraphPersisted MissingEvidenceCount BlockedActionCount ReceiptValidation execution-session-record].each do |token|
  assert(go_runtime_kde_center_page_source.include?(token), "Go Runtime KDE center page preview must include #{token}")
end

go_runtime_kde_center_page_cli_source = read_project_file("cmd/xnix-runtime-go/main.go")
%w[kde-center-page-preview session-root session-request-id KDECenterPageOptions].each do |token|
  assert(go_runtime_kde_center_page_cli_source.include?(token), "Go Runtime KDE center page CLI must include #{token}")
end

go_runtime_kde_center_page_cli_test_source = read_project_file("cmd/xnix-runtime-go/kde_center_page_cli_test.go")
%w[TestKDECenterPagePreviewCommandConsumesSessionRoot execution_session_backed execution_session_path waiting-for-runtime-gates].each do |token|
  assert(go_runtime_kde_center_page_cli_test_source.include?(token), "Go Runtime KDE center page CLI tests must include #{token}")
end

go_runtime_kde_center_cli_test_source = read_project_file("cmd/xnix-runtime-go/kde_center_cli_test.go")
%w[action_dependency_graph kde-action-dependency-graph-preview missing_evidence_count blocked_action_count dependency_graph_persisted permission_grant_created settings_persisted].each do |token|
  assert(go_runtime_kde_center_cli_test_source.include?(token), "Go Runtime KDE center CLI tests must include action graph token #{token}")
end

go_runtime_install_plan_source = read_project_file("internal/runtime/appidentity/install_plan.go")
%w[CompatibilityInstallPlanPreview CompatibilityInstallPlanPreviewWithArtifactReceipt CompatibilityInstallReadiness CompatibilityInstallArtifactStageReceipt CompatibilityInstallPhase compatibility-install-preview xnix.runtime.compatibility_install_plan.v1 GetCompatibilityInstallPlan GetCompatibilityInstallPlanPreview artifact-manifest-preview acquisition-preflight-preview package-source-preview state-root-preview recipe-install-gate artifact-stage-receipt RecipeTrustDiagnosticsReady RecipeTrustBlockingReasons ArtifactStageReceiptReady ArtifactStageDigestVerified RequiredArtifactsStaged ArtifactStageBlockingReasons recipe_trust_diagnostics_ready recipe_trust_blocking_reasons artifact_stage_receipt_ready artifact_stage_digest_verified required_artifacts_staged artifact_stage_blocking_reasons resolve-artifact-manifest verify-artifact-digests consume-artifact-stage-receipt prepare-package-source allocate-application-state review-recipe-install-gate stage-desktop-integration enable-launch-binding].each do |token|
  assert(go_runtime_install_plan_source.include?(token), "Go Runtime compatibility install preview must include #{token}")
end
%w[RuntimeOwned GoRuntimeBacked KDEPolicyOwner UserVisible InstallReady DesktopActivationReady DownloadEnabled InstallEnabled NetworkRequestCreated ArtifactsDownloaded HostRootModified PrivilegedContainerRequired DesktopShellCommandExposed BackendLaunchEnabled ExecutionStarted BackendDetailsExposed].each do |token|
  assert(go_runtime_install_plan_source.include?(token), "Go Runtime compatibility install preview must expose safety flag #{token}")
end
assert(go_runtime_install_plan_source.include?("validateNoBackendTerms"), "Go Runtime compatibility install preview must hide backend terms")

go_runtime_install_plan_cli_source = read_project_file("cmd/xnix-runtime-go/install_plan_commands.go")
%w[runCompatibilityInstallPreview parseCompatibilityInstallPreviewSource compatibility-install-preview artifact-receipt loadArtifactStageReceipt CompatibilityInstallPlanPreviewWithArtifactReceipt].each do |token|
  assert(go_runtime_install_plan_cli_source.include?(token), "Go Runtime compatibility install CLI must include #{token}")
end

go_runtime_engine_catalog_source = read_project_file("internal/runtime/appidentity/engine_catalog.go")
%w[EngineCatalogPreview CompatibilityEnginePreview engine-catalog-preview xnix.runtime.engine_catalog.v1 GetEngineCatalog GetEngineCatalogPreview go-runtime-engine-catalog automatic local-compatible isolated-compatible].each do |token|
  assert(go_runtime_engine_catalog_source.include?(token), "Go Runtime engine catalog preview must include #{token}")
end
%w[RuntimeOwned GoRuntimeBacked KDEPolicyOwner UserVisible BackendTerminologyHidden SelectionPersisted BackendInstalled BackendLaunchEnabled ExecutionStarted HostRootModified NetworkRequired BackendDetailsExposed RawCommandExposed].each do |token|
  assert(go_runtime_engine_catalog_source.include?(token), "Go Runtime engine catalog preview must expose #{token}")
end
assert(go_runtime_engine_catalog_source.include?("validateNoBackendTerms"), "Go Runtime engine catalog preview must hide backend terms")

go_runtime_engine_catalog_cli_source = read_project_file("cmd/xnix-runtime-go/main.go")
assert(go_runtime_engine_catalog_cli_source.include?("engine-catalog-preview"), "Go Runtime CLI must expose engine catalog preview")
assert(go_runtime_engine_catalog_cli_source.include?("NewEngineCatalogPreview"), "Go Runtime CLI must call engine catalog preview")

go_runtime_run_plan_source = read_project_file("internal/runtime/appidentity/run_plan.go")
%w[RunPlanPreview RunPlanApplication RunPlanExecution RunPlanEngineSummary RunPlanProfile RunPlanBinding RunPlanPreflight run-plan-preview xnix.runtime.run_plan.v1 GetRunPlan GetRunPlanPreview registry+go-runtime-run-plan compatibility-run automatic-managed local-compatible-managed isolated-compatible-managed].each do |token|
  assert(go_runtime_run_plan_source.include?(token), "Go Runtime run plan preview must include #{token}")
end
%w[RuntimeOwned GoRuntimeBacked KDEPolicyOwner UserVisible LaunchEnabled ExecutionRequestCreated ExecutionStarted BackendBindingReady RequestObjectCreated PermissionGranted HostRootModified NetworkRequired BackendDetailsExposed RawCommandExposed].each do |token|
  assert(go_runtime_run_plan_source.include?(token), "Go Runtime run plan preview must expose #{token}")
end
assert(go_runtime_run_plan_source.include?("ValidateSafeForDesktop"), "Go Runtime run plan preview must validate desktop safety")
assert(go_runtime_run_plan_source.include?("validateNoBackendTerms"), "Go Runtime run plan preview must hide backend terms")

go_runtime_run_plan_cli_source = read_project_file("cmd/xnix-runtime-go/main.go")
assert(go_runtime_run_plan_cli_source.include?("run-plan-preview"), "Go Runtime CLI must expose run plan preview")
assert(go_runtime_run_plan_cli_source.include?("RunPlanPreview"), "Go Runtime CLI must call run plan preview")

go_runtime_applications_source = read_project_file("internal/runtime/appidentity/applications.go")
%w[ApplicationsPreview ApplicationPreview RuntimeApplicationPreview applications-preview application-preview xnix.runtime.applications.v1 xnix.runtime.application.v1 ListApplications GetApplication registry+go-runtime-desktop-identity start-menu task-manager file-manager system-tray notification-center compatibility-center unified-settings].each do |token|
  assert(go_runtime_applications_source.include?(token), "Go Runtime applications preview must include #{token}")
end
%w[RuntimeOwned GoRuntimeBacked KDEPolicyOwner UserVisible StandardDesktopEntries StandardDesktopEntry RecipeRegistryVerified BackendTerminologyHidden LaunchEnabled ExecutionStarted DesktopFilesWritten DesktopFileWritten HostRootModified NetworkRequired BackendDetailsExposed RawWindowsExecutableExposed].each do |token|
  assert(go_runtime_applications_source.include?(token), "Go Runtime applications preview must expose #{token}")
end
%w[NewPlanWithProvenance ValidateSafeForDesktop modeLabel normalizedExtensions validateNoBackendTerms].each do |token|
  assert(go_runtime_applications_source.include?(token), "Go Runtime applications preview must call #{token}")
end

go_runtime_applications_cli_source = read_project_file("cmd/xnix-runtime-go/main.go")
assert(go_runtime_applications_cli_source.include?("applications-preview"), "Go Runtime CLI must expose applications preview")
assert(go_runtime_applications_cli_source.include?("application-preview"), "Go Runtime CLI must expose application preview")
assert(go_runtime_applications_cli_source.include?("NewApplicationsPreview"), "Go Runtime CLI must call applications preview")
assert(go_runtime_applications_cli_source.include?("NewApplicationPreview"), "Go Runtime CLI must call application preview")

go_runtime_diagnostics_source = read_project_file("internal/runtime/appidentity/diagnostics.go")
%w[DiagnosticsPreview diagnostics-preview xnix.runtime.diagnostics.v1 GetDiagnostics GetDiagnosticsPreview registry+go-runtime-previews GetAIDiagnosticInput compatibility-status-review runtime-metadata-only].each do |token|
  assert(go_runtime_diagnostics_source.include?(token), "Go Runtime diagnostics preview must include #{token}")
end
%w[RuntimeOwned GoRuntimeBacked KDEPolicyOwner UserVisible CompatibilityCenterCard SafeForAIDiagnostics AIProviderCallEnabled FileContentRead FilePathsExposed RequestObjectCreated PermissionGranted LaunchEnabled ExecutionStarted RepairExecutionEnabled SettingsPersisted HostRootModified NetworkRequired BackendDetailsExposed].each do |token|
  assert(go_runtime_diagnostics_source.include?(token), "Go Runtime diagnostics preview must expose #{token}")
end
%w[NewPlanWithProvenance ExecutionReadinessPreview LaunchIntentPreview KDEActionQueuePreview NewCompatibilityCenterPreview NewKDECenterPageSectionsPreview validateNoBackendTerms].each do |token|
  assert(go_runtime_diagnostics_source.include?(token), "Go Runtime diagnostics preview must call #{token}")
end

go_runtime_diagnostics_cli_source = read_project_file("cmd/xnix-runtime-go/main.go")
assert(go_runtime_diagnostics_cli_source.include?("diagnostics-preview"), "Go Runtime CLI must expose Runtime diagnostics preview")
assert(go_runtime_diagnostics_cli_source.include?("NewDiagnosticsPreview"), "Go Runtime CLI must call Runtime diagnostics preview")

go_runtime_kde_center_page_source = read_project_file("internal/runtime/appidentity/kde_center_page.go")
%w[KDECenterPageSectionsPreview KDECenterPageSectionDetailPreview KDEDiagnosticHistoryRoute diagnostic_history_route diagnostic-history-route compatibility-diagnostics GetDiagnostics GetDiagnosticHistoryPreview diagnostic-history-preview NewDolphinAIAnalysisPreview StateRootRequired StateRootPathExposed HistoryPreviewCreated].each do |token|
  assert(go_runtime_kde_center_page_source.include?(token), "Go Runtime KDE center page diagnostics section must include #{token}")
end

go_runtime_ai_diagnostics_source = read_project_file("internal/runtime/appidentity/ai_diagnostics.go")
%w[AIDiagnosticInputPreview AIDiagnosticRecommendationPreview AIRepairApprovalGatePreview ai-diagnostic-input-preview ai-diagnostic-recommendation-preview ai-repair-approval-gate-preview xnix.runtime.ai_diagnostic_input.v1 xnix.runtime.ai_diagnostic_recommendation.v1 xnix.runtime.ai_repair_approval_gate.v1 GetAIDiagnosticInput GetAIDiagnosticInputPreview GetAIDiagnosticRecommendation GetAIDiagnosticRecommendationPreview GetAIRepairApprovalGate GetAIRepairApprovalGatePreview registry+go-runtime-ai-diagnostics compatibility-center-review runtime-approval-token restore-point-preflight blocked-until-approval].each do |token|
  assert(go_runtime_ai_diagnostics_source.include?(token), "Go Runtime AI diagnostics previews must include #{token}")
end
%w[RuntimeOwned GoRuntimeBacked KDEPolicyOwner AIProviderCalled AIProviderCallEnabled NetworkRequired SafeForAIDiagnostics FileContentRead FilePathsExposed RequestObjectCreated PermissionGranted RepairExecutionRequested RepairExecuted AutoExecutionAllowed HostRootModified BackendDetailsExposed].each do |token|
  assert(go_runtime_ai_diagnostics_source.include?(token), "Go Runtime AI diagnostics previews must expose safety flag #{token}")
end
%w[RunPlanPreview TestResultPreview RepairPlanPreview validateNoBackendTerms].each do |token|
  assert(go_runtime_ai_diagnostics_source.include?(token), "Go Runtime AI diagnostics previews must call #{token}")
end

go_runtime_ai_diagnostics_cli_source = read_project_file("cmd/xnix-runtime-go/ai_diagnostics_commands.go")
%w[runAIDiagnosticInputPreview runAIDiagnosticRecommendationPreview runAIRepairApprovalGatePreview ai-diagnostic-input-preview ai-diagnostic-recommendation-preview ai-repair-approval-gate-preview].each do |token|
  assert(go_runtime_ai_diagnostics_cli_source.include?(token), "Go Runtime AI diagnostics CLI must include #{token}")
end

go_runtime_owner_readiness_source = read_project_file("internal/runtime/appidentity/runtime_owner_readiness.go")
%w[RuntimeOwnerReadinessPreview runtime-owner-readiness-preview xnix.runtime.owner_readiness.v1 GetRuntimeOwnerReadiness GetRuntimeOwnerReadinessPreview runtime-service-binding-preview runtime-live-owner-gate-preview runtime-owner-process-preview runtime-owner-smoke-plan-preview runtime-method-parity-manifest-preview runtime-owner-route-manifest-preview runtime-owner-recipe-trust-preview activation-binding read-only-method-parity owner-smoke-plan write-method-gate kde-ownership-boundary host-safety-boundary long-running-runtime-owner read-only-owner-routes production-bus-claim production-recipe-trust].each do |token|
  assert(go_runtime_owner_readiness_source.include?(token), "Go Runtime owner readiness preview must include #{token}")
end
%w[ActivationBindingReady ReadOnlyMethodParityReady OwnerSmokePlanned LiveDBusOwnerReady ProductionOwnerEnabled OwnerTransitionReady ProductionRecipeTrustReady ProductionOwnerRoutesReady RuntimeOwned GoRuntimeBacked KDEPolicyOwner KDEMayClaimRuntimeOwnership NetworkRequired HostRootModified PrivilegedContainerRequired SystemServiceStarted ProductionBusClaimed WriteMethodsEnabled BackendDetailsExposed].each do |token|
  assert(go_runtime_owner_readiness_source.include?(token), "Go Runtime owner readiness preview must expose #{token}")
end
assert(go_runtime_owner_readiness_source.include?("NewRuntimeServiceBindingPreview"), "Go Runtime owner readiness preview must derive from the service binding preview")
assert(go_runtime_owner_readiness_source.include?("NewRuntimeLiveOwnerGatePreview"), "Go Runtime owner readiness preview must derive from the live owner gate preview")
assert(go_runtime_owner_readiness_source.include?("NewRuntimeOwnerProcessPreview"), "Go Runtime owner readiness preview must derive from the owner process preview")
assert(go_runtime_owner_readiness_source.include?("NewRuntimeOwnerSmokePlanPreview"), "Go Runtime owner readiness preview must derive from the owner smoke plan preview")
assert(go_runtime_owner_readiness_source.include?("NewRuntimeMethodParityManifestPreview"), "Go Runtime owner readiness preview must derive from the method parity manifest preview")
assert(go_runtime_owner_readiness_source.include?("NewRuntimeOwnerRouteManifestPreview"), "Go Runtime owner readiness preview must derive from the owner route manifest preview")
assert(go_runtime_owner_readiness_source.include?("NewRuntimeOwnerRecipeTrustPreview"), "Go Runtime owner readiness preview must derive from the owner recipe trust preview")
assert(go_runtime_owner_readiness_source.include?("validateNoBackendTerms"), "Go Runtime owner readiness preview must hide backend terms")

go_runtime_owner_readiness_cli_source = read_project_file("cmd/xnix-runtime-go/main.go")
assert(go_runtime_owner_readiness_cli_source.include?("runtime-owner-readiness-preview"), "Go Runtime CLI must expose Runtime owner readiness preview")

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

go_runtime_write_gate_source = read_project_file("internal/runtime/appidentity/runtime_write_gate.go")
%w[RuntimeWriteGatePreview RuntimeWriteGateProductionGate RuntimeWriteGateCheck RuntimeWriteGateCounts runtime-write-gate-preview xnix.runtime.write_gate.v1 go-runtime-write-gate runtime-service-activation-preflight-preview production-dbus-gate-review-preview production-dbus-human-authorization-preflight-preview GetRuntimeWriteGate GetRuntimeWriteGatePreview InstallRecipe Launch CreateSnapshot RestoreSnapshot blocked-until-production-backend production-gates-consumed-write-gate-disabled production-gate-consumption-blocked WriteMethodDisabled production-dbus-gate-review production-service-activation-preflight human-authorization-receipt production-runtime-owner backend-binding-ready recipe-trust-production user-action-review portal-approval-if-sensitive snapshot-preflight-for-risky-change].each do |token|
  assert(go_runtime_write_gate_source.include?(token), "Go Runtime write gate preview must include #{token}")
end
%w[ProductionGateDecision ServiceActivationPreflightReady ProductionDBusGateReady HumanAuthorizationPreflightReady HumanAuthorizationRequired HumanAuthorizationGranted AuthorizationReceiptAccepted ProductionActivationReady RestrictedSmokeReady SystemServiceStarted ProductionBusClaimed WriteMethodsEnabled WriteMethodEnabled DispatchEnabled RequestObjectCreated ExecutionStarted NetworkRequired HostRootModified PrivilegedContainerRequired BackendDetailsExposed].each do |token|
  assert(go_runtime_write_gate_source.include?(token), "Go Runtime write gate preview must expose safety flag #{token}")
end
assert(go_runtime_write_gate_source.include?("NewRuntimeServiceActivationPreflightPreview"), "Go Runtime write gate preview must consume Runtime service activation preflight")
assert(go_runtime_write_gate_source.include?("validateNoBackendTerms"), "Go Runtime write gate preview must hide backend terms")

go_runtime_write_gate_cli_source = read_project_file("cmd/xnix-runtime-go/runtime_write_gate_cli_test.go")
%w[runtime-write-gate-preview GetRuntimeWriteGate GetRuntimeWriteGatePreview WriteMethodDisabled production_gate production_gate_decision production-gates-consumed-write-gate-disabled authorization_receipt_accepted].each do |token|
  assert(go_runtime_write_gate_cli_source.include?(token), "Go Runtime write gate CLI test must include #{token}")
end

go_runtime_safety_source = [
  read_project_file("internal/runtime/appidentity/state_root.go"),
  read_project_file("internal/runtime/appidentity/snapshot_plan.go"),
  read_project_file("internal/runtime/appidentity/portal_access_policy.go")
].join("\n")
%w[ApplicationStateRootPreview SnapshotPlanPreview PortalAccessPolicyPreview state-root-preview snapshot-plan-preview portal-access-policy-preview xnix.runtime.state_root.v1 xnix.runtime.snapshot_plan.v1 xnix.runtime.portal_access_policy.v1 GetApplicationStateRoot GetApplicationStateRootPreview GetSnapshotPlan GetSnapshotPlanPreview GetPortalAccessPolicy GetPortalAccessPolicyPreview go-runtime-state-root go-runtime-snapshot-plan go-runtime-portal-access-policy application-data runtime-metadata diagnostic-cache desktop-activation-receipts before-repair before-engine-change manual file-open camera remote-desktop].each do |token|
  assert(go_runtime_safety_source.include?(token), "Go Runtime safety previews must include #{token}")
end
%w[DirectoriesCreated HostRootModified UserDocumentsIncluded PortalRequiredForUserFiles SnapshotEligible RestoreRequiresConfirm SnapshotRequestCreated SnapshotCreated RestoreRequested RestoreExecuted DirectAccessAllowed RequestObjectCreated PermissionGranted HostPermissionChanged NetworkRequired BackendDetailsExposed].each do |token|
  assert(go_runtime_safety_source.include?(token), "Go Runtime safety previews must expose safety flag #{token}")
end
assert(go_runtime_safety_source.include?("validateNoBackendTerms"), "Go Runtime safety previews must hide backend terms")

go_runtime_safety_cli_source = read_project_file("cmd/xnix-runtime-go/runtime_safety_commands.go") +
                               read_project_file("cmd/xnix-runtime-go/main.go")
%w[runStateRootPreview runSnapshotPlanPreview runPortalAccessPolicyPreview runPortalRequestRecord state-root-preview snapshot-plan-preview portal-access-policy-preview portal-request-record].each do |token|
  assert(go_runtime_safety_cli_source.include?(token), "Go Runtime safety CLI must include #{token}")
end

go_runtime_state_root_quota_source = read_project_file("internal/runtime/appidentity/state_root_quota_retention.go")
%w[StateRootQuotaRetentionPreview StateRootQuotaRetentionSection StateRootQuotaRetentionCounts state-root-quota-retention-preview xnix.runtime.state_root_quota_retention.v1 GetStateRootQuotaRetention GetStateRootQuotaRetentionPreview go-runtime-state-root-quota-retention-preview snapshots diagnostics execution-receipts portal-receipts artifact-receipts activation-receipts unknown-records cleanup-candidate-over-quota retain-under-quota retain-malformed-record retain-unknown-record-review retain-active-session retain-retention-exempt].each do |token|
  assert(go_runtime_state_root_quota_source.include?(token), "Go Runtime state-root quota retention preview must include #{token}")
end
%w[FileDeletionEnabled DirectoriesCreated LogTruncationEnabled ReceiptsRewritten SnapshotDeletionEnabled StateRootPathExposed HostRootModified BackendDetailsExposed NetworkRequired PrivilegedContainerRequired].each do |token|
  assert(go_runtime_state_root_quota_source.include?(token), "Go Runtime state-root quota retention preview must expose #{token}")
end
%w[WalkDir ReadFile Stat validateNoBackendTerms refusing\ to\ inspect\ filesystem\ root].each do |token|
  assert(go_runtime_state_root_quota_source.include?(token), "Go Runtime state-root quota retention preview must implement #{token}")
end

go_runtime_state_root_quota_test_source = read_project_file("internal/runtime/appidentity/state_root_quota_retention_test.go")
%w[TestStateRootQuotaRetentionPreviewMissingRootDoesNotCreateDirectories TestStateRootQuotaRetentionPreviewUnderQuotaClassifiesKnownSections TestStateRootQuotaRetentionPreviewOverQuotaAndMalformedUnknownEvidence TestStateRootQuotaRetentionPreviewActiveSessionBlocksExecutionReceipts TestStateRootQuotaRetentionPreviewRetentionExemptEvidenceIsKept TestStateRootQuotaRetentionPreviewRejectsUnsafeInputs retain-missing-state-root cleanup-candidate-over-quota malformed-record-review-required unknown-record-review-required active-session-blocked retention-exempt].each do |token|
  assert(go_runtime_state_root_quota_test_source.include?(token), "Go Runtime state-root quota retention tests must include #{token}")
end

go_runtime_state_root_quota_cli_source = read_project_file("cmd/xnix-runtime-go/main.go") +
                                         read_project_file("cmd/xnix-runtime-go/state_root_quota_retention_commands.go")
%w[runStateRootQuotaRetentionPreview parseStateRootQuotaRetentionPreviewSource state-root-quota-retention-preview quota-bytes active-session StateRootQuotaRetentionOptions].each do |token|
  assert(go_runtime_state_root_quota_cli_source.include?(token), "Go Runtime state-root quota retention CLI must include #{token}")
end

go_runtime_state_root_quota_cli_test_source = read_project_file("cmd/xnix-runtime-go/state_root_quota_retention_cli_test.go")
%w[TestStateRootQuotaRetentionPreviewCLI TestStateRootQuotaRetentionPreviewCLIMissingRootIsReadOnly TestStateRootQuotaRetentionPreviewCLIRequiresFlags xnix.runtime.state_root_quota_retention.v1 retain-missing-state-root state_root_path_exposed host_root_modified].each do |token|
  assert(go_runtime_state_root_quota_cli_test_source.include?(token), "Go Runtime state-root quota retention CLI tests must include #{token}")
end

go_runtime_crash_hang_source = read_project_file("internal/runtime/appidentity/crash_hang_signal_summary.go")
%w[CrashHangSignalSummaryPreview CrashHangSignalGroup CrashHangSignalSummaryCounts CrashHangSignalPrivacyRedaction crash-hang-signal-summary-preview xnix.runtime.crash_hang_signal_summary.v1 GetCrashHangSignalSummary GetCrashHangSignalSummaryPreview crash hang timeout missing-dependency permission-denial graphics-issue network-issue regression-after-repair].each do |token|
  assert(go_runtime_crash_hang_source.include?(token), "Go Runtime crash-hang signal summary preview must include #{token}")
end
%w[PrivateLogRead FileContentRead AIProviderCallEnabled RepairExecuted BackendProcessStarted HostRootModified StateRootPathExposed BackendDetailsExposed crashHangGroupForSignal crashHangPrivacySensitive].each do |token|
  assert(go_runtime_crash_hang_source.include?(token), "Go Runtime crash-hang signal summary preview must expose #{token}")
end
%w[validateNoBackendTerms safety.ValidatePayload].each do |token|
  assert(go_runtime_crash_hang_source.include?(token), "Go Runtime crash-hang signal summary preview must validate through #{token}")
end

go_runtime_crash_hang_test_source = read_project_file("internal/runtime/appidentity/crash_hang_signal_summary_test.go")
%w[TestCrashHangSignalSummaryNoHistoryIsSafe TestCrashHangSignalSummaryGroupsAndRecurrence TestCrashHangSignalSummaryMalformedHistoryIsBlocked TestCrashHangSignalSummaryMixedApplicationIDsAreFiltered TestCrashHangSignalSummaryDuplicateSignalIDsAreCounted TestCrashHangSignalSummaryPrivacySensitiveFieldsAreOmitted TestCrashHangSignalSummaryBlockedRepairEvidenceIsCounted TestCrashHangSignalSummaryRegressionAfterRepair TestCrashHangSignalSummaryRejectsUnsafeInputs].each do |token|
  assert(go_runtime_crash_hang_test_source.include?(token), "Go Runtime crash-hang signal summary tests must include #{token}")
end

go_runtime_crash_hang_cli_source = read_project_file("cmd/xnix-runtime-go/main.go") +
                                   read_project_file("cmd/xnix-runtime-go/crash_hang_signal_summary_commands.go")
%w[runCrashHangSignalSummaryPreview crash-hang-signal-summary-preview LenientHistory NewCrashHangSignalSummaryPreview OpenRunRecordStoreReadOnly].each do |token|
  assert(go_runtime_crash_hang_cli_source.include?(token), "Go Runtime crash-hang signal summary CLI must include #{token}")
end

go_runtime_crash_hang_cli_test_source = read_project_file("cmd/xnix-runtime-go/crash_hang_signal_summary_cli_test.go")
%w[TestCrashHangSignalSummaryPreviewCLI TestCrashHangSignalSummaryPreviewCLIMalformedRecordIsBlocked TestCrashHangSignalSummaryPreviewCLIMissingRootIsReadOnly TestCrashHangSignalSummaryPreviewCLIRequiresFlags xnix.runtime.crash_hang_signal_summary.v1 blocked-malformed-history state_root_path_exposed].each do |token|
  assert(go_runtime_crash_hang_cli_test_source.include?(token), "Go Runtime crash-hang signal summary CLI tests must include #{token}")
end

go_runtime_diagnostic_history_source = read_project_file("internal/runtime/diagnostics/history.go")
%w[LenientHistory buildHistory].each do |token|
  assert(go_runtime_diagnostic_history_source.include?(token), "Go Runtime diagnostic history must implement #{token}")
end
assert(read_project_file("internal/runtime/diagnostics/record.go").include?("listLenient"), "Go Runtime diagnostic records must implement listLenient")

go_runtime_permission_audit_source = read_project_file("internal/runtime/appidentity/permission_evidence_audit.go")
%w[PermissionEvidenceAuditPreview PermissionEvidenceAuditRow PermissionEvidenceAuditCounts PermissionEvidenceRowInput permission-evidence-audit-preview xnix.runtime.permission_evidence_audit.v1 GetPermissionEvidenceAudit GetPermissionEvidenceAuditPreview documents downloads uris print clipboard screenshot camera remote-desktop network].each do |token|
  assert(go_runtime_permission_audit_source.include?(token), "Go Runtime permission evidence audit preview must include #{token}")
end
%w[consistent setting-only receipt-only expired denied missing-review blocked-by-policy unsupported].each do |token|
  assert(go_runtime_permission_audit_source.include?(token), "Go Runtime permission evidence audit preview must classify #{token}")
end
%w[RealPortalCallEnabled PermissionGrantEnabled PermissionRevokeEnabled ReceiptWriteEnabled SettingsPersisted ExecutionApproved StateRootPathExposed HostRootModified validateNoBackendTerms safety.ValidatePayload PermissionReviewPreview DesktopResourceBridgePreview ExecutionPreflightPreview].each do |token|
  assert(go_runtime_permission_audit_source.include?(token), "Go Runtime permission evidence audit preview must expose #{token}")
end

go_runtime_permission_audit_test_source = read_project_file("internal/runtime/appidentity/permission_evidence_audit_test.go")
%w[TestPermissionEvidenceAuditConsistencyStates TestPermissionEvidenceAuditAllConsistentIsClean TestPermissionEvidenceAuditMalformedReceiptsAreSurfaced TestPermissionEvidenceAuditUnsupportedResourceIsFlagged TestPermissionEvidenceAuditRejectsUnsafeInputs TestPlanPermissionEvidenceAuditPreviewJoinsSources].each do |token|
  assert(go_runtime_permission_audit_test_source.include?(token), "Go Runtime permission evidence audit tests must include #{token}")
end

go_runtime_permission_audit_cli_source = read_project_file("cmd/xnix-runtime-go/main.go") +
                                         read_project_file("cmd/xnix-runtime-go/permission_evidence_audit_commands.go")
%w[runPermissionEvidenceAuditPreview permission-evidence-audit-preview PermissionEvidenceAuditPreview ReadRequests Summarize loadRecipe].each do |token|
  assert(go_runtime_permission_audit_cli_source.include?(token), "Go Runtime permission evidence audit CLI must include #{token}")
end

go_runtime_permission_audit_cli_test_source = read_project_file("cmd/xnix-runtime-go/permission_evidence_audit_cli_test.go")
%w[TestPermissionEvidenceAuditPreviewCLINoReceipts TestPermissionEvidenceAuditPreviewCLIWithReceipts TestPermissionEvidenceAuditPreviewCLIMalformedLedgerIsSurfaced TestPermissionEvidenceAuditPreviewCLIRequiresSource xnix.runtime.permission_evidence_audit.v1 state_root_path_exposed].each do |token|
  assert(go_runtime_permission_audit_cli_test_source.include?(token), "Go Runtime permission evidence audit CLI tests must include #{token}")
end

go_runtime_portal_renewal_source = read_project_file("internal/runtime/appidentity/portal_permission_renewal.go")
%w[PortalPermissionRenewalPreview PortalPermissionRenewalRow PortalPermissionRenewalCounts PortalPermissionRenewalReceipt portal-permission-renewal-preview xnix.runtime.portal_permission_renewal.v1 GetPortalPermissionRenewal GetPortalPermissionRenewalPreview current needs-review expiring-soon denied revoked missing-receipt blocked-by-policy files uris print clipboard screen camera remote-desktop network].each do |token|
  assert(go_runtime_portal_renewal_source.include?(token), "Go Runtime Portal permission renewal preview must include #{token}")
end
%w[RealPortalCallEnabled PermissionGrantEnabled PermissionRevokeEnabled ReceiptWriteEnabled SettingsPersisted ExecutionApproved StateRootPathExposed FileContentRead FilePathsExposed RawCommandExposed BackendDetailsExposed HostRootModified validateNoBackendTerms safety.ValidatePayload].each do |token|
  assert(go_runtime_portal_renewal_source.include?(token), "Go Runtime Portal permission renewal preview must expose #{token}")
end

go_runtime_portal_renewal_test_source = read_project_file("internal/runtime/appidentity/portal_permission_renewal_test.go")
%w[TestPortalPermissionRenewalStates TestPortalPermissionRenewalMissingReceiptAndMalformedLedger TestPortalPermissionRenewalRejectsUnsafePlan current expiring-soon needs-review denied revoked missing-receipt blocked-by-policy RealPortalCallEnabled PermissionGrantEnabled PermissionRevokeEnabled ReceiptWriteEnabled SettingsPersisted ExecutionApproved HostRootModified].each do |token|
  assert(go_runtime_portal_renewal_test_source.include?(token), "Go Runtime Portal permission renewal tests must include #{token}")
end

go_runtime_portal_renewal_cli_source = read_project_file("cmd/xnix-runtime-go/main.go") +
                                       read_project_file("cmd/xnix-runtime-go/portal_permission_renewal_commands.go")
%w[runPortalPermissionRenewalPreview portal-permission-renewal-preview PortalPermissionRenewalPreview ReadRequests portalPermissionRenewalStateFromRequest StateCancelled expiring].each do |token|
  assert(go_runtime_portal_renewal_cli_source.include?(token), "Go Runtime Portal permission renewal CLI must include #{token}")
end

go_runtime_portal_renewal_cli_test_source = read_project_file("cmd/xnix-runtime-go/portal_permission_renewal_cli_test.go")
%w[TestPortalPermissionRenewalPreviewCLI TestPortalPermissionRenewalPreviewCLIMissingStateRootDoesNotCreate TestPortalPermissionRenewalPreviewCLIMalformedLedgerIsSurfaced TestPortalPermissionRenewalPreviewCLIRequiresRecipeSource xnix.runtime.portal_permission_renewal.v1 real_portal_call_enabled permission_grant_enabled permission_revoke_enabled receipt_write_enabled settings_persisted execution_approved host_root_modified].each do |token|
  assert(go_runtime_portal_renewal_cli_test_source.include?(token), "Go Runtime Portal permission renewal CLI tests must include #{token}")
end

go_runtime_settings_profile_migration_source = read_project_file("internal/runtime/appidentity/settings_profile_migration.go")
%w[SettingsProfileMigrationPreview SettingsProfileMigrationRow SettingsProfileMigrationCounts settings-profile-migration-preview xnix.runtime.settings_profile_migration.v1 GetCompatibilitySettingsProfileMigration GetCompatibilitySettingsProfileMigrationPreview old-schema current-schema future-schema run-mode.mode run-mode.preference resource-access.documents resource-access.downloads devices.camera network.network snapshots.snapshots diagnostics.privacy].each do |token|
  assert(go_runtime_settings_profile_migration_source.include?(token), "Go Runtime settings profile migration preview must include #{token}")
end
%w[SettingsPersisted SettingsPersistenceEnabled ResourceGrantEnabled RealPortalCallEnabled BackendLaunchEnabled BackendProcessStarted HostRootModified StateRootPathExposed RawCommandExposed RawExecutableExposed FileContentRead BackendDetailsExposed validateNoBackendTerms settingsProfileSchemaNumber].each do |token|
  assert(go_runtime_settings_profile_migration_source.include?(token), "Go Runtime settings profile migration preview must expose #{token}")
end

go_runtime_settings_profile_migration_test_source = read_project_file("internal/runtime/appidentity/settings_profile_migration_test.go")
%w[TestSettingsProfileMigrationOldSchemaRequiresReview TestSettingsProfileMigrationCurrentSchemaIsUnchanged TestSettingsProfileMigrationFutureSchemaBlocksMigration TestSettingsProfileMigrationBlockedAndInvalidSetting rollback notes SettingsPersisted].each do |token|
  assert(go_runtime_settings_profile_migration_test_source.include?(token), "Go Runtime settings profile migration tests must include #{token}")
end

go_runtime_settings_profile_migration_cli_source = read_project_file("cmd/xnix-runtime-go/main.go") +
                                                   read_project_file("cmd/xnix-runtime-go/settings_profile_migration_commands.go")
%w[runSettingsProfileMigrationPreview settings-profile-migration-preview SettingsProfileMigrationPreview from-schema to-schema blocked-setting].each do |token|
  assert(go_runtime_settings_profile_migration_cli_source.include?(token), "Go Runtime settings profile migration CLI must include #{token}")
end

go_runtime_settings_profile_migration_cli_test_source = read_project_file("cmd/xnix-runtime-go/settings_profile_migration_cli_test.go")
%w[TestSettingsProfileMigrationPreviewCLI TestSettingsProfileMigrationPreviewCLIRejectsUnsafeArguments xnix.runtime.settings_profile_migration.v1 settings_persisted resource_grant_enabled real_portal_call_enabled backend_launch_enabled host_root_modified].each do |token|
  assert(go_runtime_settings_profile_migration_cli_test_source.include?(token), "Go Runtime settings profile migration CLI tests must include #{token}")
end

go_runtime_backend_fallback_source = read_project_file("internal/runtime/appidentity/compatibility_backend_fallback.go")
%w[CompatibilityBackendFallbackPreview CompatibilityBackendFallbackCandidate compatibility-backend-fallback-preview xnix.runtime.compatibility_backend_fallback.v1 GetCompatibilityBackendFallback GetCompatibilityBackendFallbackPreview local-compatibility isolated-compatibility automatic recipe-requests-isolation local-selected-by-default capability-not-ready diagnostics-regressions-present blocked-missing-evidence blocked-by-policy].each do |token|
  assert(go_runtime_backend_fallback_source.include?(token), "Go Runtime compatibility backend fallback preview must include #{token}")
end
%w[SelectionPersisted EngineInstallEnabled BackendLaunchEnabled VMStartEnabled BackendDetailsExposed StateRootPathExposed HostRootModified validateNoBackendTerms safety.ValidatePayload recommendedCompatibilityProfileID BackendLifecyclePreview NewBackendCapabilityMatrixPreview].each do |token|
  assert(go_runtime_backend_fallback_source.include?(token), "Go Runtime compatibility backend fallback preview must expose #{token}")
end

go_runtime_backend_fallback_test_source = read_project_file("internal/runtime/appidentity/compatibility_backend_fallback_test.go")
%w[TestCompatibilityBackendFallbackLocalPrimaryStaysBlocked TestCompatibilityBackendFallbackIsolationRequiredForbidsLocal TestCompatibilityBackendFallbackDiagnosticsRiskIsSurfaced TestCompatibilityBackendFallbackAutomaticExplainsBlocked].each do |token|
  assert(go_runtime_backend_fallback_test_source.include?(token), "Go Runtime compatibility backend fallback tests must include #{token}")
end

go_runtime_backend_fallback_cli_source = read_project_file("cmd/xnix-runtime-go/main.go") +
                                         read_project_file("cmd/xnix-runtime-go/compatibility_backend_fallback_commands.go")
%w[runCompatibilityBackendFallbackPreview compatibility-backend-fallback-preview CompatibilityBackendFallbackPreview loadRecipe].each do |token|
  assert(go_runtime_backend_fallback_cli_source.include?(token), "Go Runtime compatibility backend fallback CLI must include #{token}")
end

go_runtime_backend_fallback_cli_test_source = read_project_file("cmd/xnix-runtime-go/compatibility_backend_fallback_cli_test.go")
%w[TestCompatibilityBackendFallbackPreviewCLI TestCompatibilityBackendFallbackPreviewCLIRequiresSource xnix.runtime.compatibility_backend_fallback.v1].each do |token|
  assert(go_runtime_backend_fallback_cli_test_source.include?(token), "Go Runtime compatibility backend fallback CLI tests must include #{token}")
end

go_runtime_search_visibility_source = read_project_file("internal/runtime/appidentity/kde_search_visibility.go")
%w[KDESearchVisibilityPlanPreview KDESearchVisibilityRow KDESearchVisibilityCounts kde-search-visibility-plan-preview xnix.runtime.kde_search_visibility.v1 GetKDESearchVisibilityPlan GetKDESearchVisibilityPlanPreview launcher krunner file-association dolphin-action compatibility-center settings task-manager receipt-required no-mime-association application-hidden].each do |token|
  assert(go_runtime_search_visibility_source.include?(token), "Go Runtime KDE search visibility plan preview must include #{token}")
end
%w[DesktopFilesWritten MIMEDefaultsWritten KDECacheRefreshed HostFilesIndexed SearchIndexPersisted BackendLaunchEnabled HostRootModified kdeSearchSafeSynonyms validateNoBackendTerms safety.ValidatePayload DesktopActivationReceiptEvidence].each do |token|
  assert(go_runtime_search_visibility_source.include?(token), "Go Runtime KDE search visibility plan preview must expose #{token}")
end

go_runtime_search_visibility_test_source = read_project_file("internal/runtime/appidentity/kde_search_visibility_test.go")
%w[TestKDESearchVisibilityReceiptRequiredWithoutActivation TestKDESearchVisibilityUnsupportedMIME TestKDESearchVisibilityHiddenApplication TestKDESearchSafeSynonymsDropUnsafeAndDuplicates TestKDESearchVisibilityRejectsUnsafePlan].each do |token|
  assert(go_runtime_search_visibility_test_source.include?(token), "Go Runtime KDE search visibility tests must include #{token}")
end

go_runtime_search_visibility_cli_source = read_project_file("cmd/xnix-runtime-go/main.go") +
                                          read_project_file("cmd/xnix-runtime-go/kde_search_visibility_commands.go")
%w[runKDESearchVisibilityPlanPreview kde-search-visibility-plan-preview KDESearchVisibilityPlanPreview loadRecipe].each do |token|
  assert(go_runtime_search_visibility_cli_source.include?(token), "Go Runtime KDE search visibility CLI must include #{token}")
end

go_runtime_search_visibility_cli_test_source = read_project_file("cmd/xnix-runtime-go/kde_search_visibility_cli_test.go")
%w[TestKDESearchVisibilityPlanPreviewCLI TestKDESearchVisibilityPlanPreviewCLIRejectsMissingApp xnix.runtime.kde_search_visibility.v1].each do |token|
  assert(go_runtime_search_visibility_cli_test_source.include?(token), "Go Runtime KDE search visibility CLI tests must include #{token}")
end

go_runtime_notification_digest_source = read_project_file("internal/runtime/appidentity/kde_notification_digest.go")
%w[KDENotificationDigestPreview KDENotificationDigestEntry KDENotificationDigestCounts kde-notification-digest-preview review-only-runtime-event-digest xnix.runtime.kde_notification_digest.v1 GetKDENotificationDigest GetKDENotificationDigestPreview needs-review blocked-action permission-attention diagnostic-issue snapshot-warning readiness-change].each do |token|
  assert(go_runtime_notification_digest_source.include?(token), "Go Runtime KDE notification digest must include #{token}")
end
%w[NotificationsSent LiveTrayBridgeEnabled RequestObjectsCreated PermissionGrantEnabled BackendLaunchEnabled BackendProcessStarted StateRootPathExposed RawCommandExposed BackendDetailsExposed HostRootModified].each do |token|
  assert(go_runtime_notification_digest_source.include?(token), "Go Runtime KDE notification digest must expose disabled gate #{token}")
end

go_runtime_notification_digest_test_source = read_project_file("internal/runtime/appidentity/kde_notification_digest_test.go")
%w[TestKDENotificationDigestCoversSixGroups TestKDENotificationDigestDeduplicatesEvents TestKDENotificationDigestRejectsUnsupportedAndUnsafeInput TestKDENotificationDigestRejectsUnsafeIdentity].each do |token|
  assert(go_runtime_notification_digest_test_source.include?(token), "Go Runtime KDE notification digest tests must include #{token}")
end

go_runtime_notification_digest_cli_source = read_project_file("cmd/xnix-runtime-go/kde_notification_digest_commands.go")
%w[kde-notification-digest-preview runKDENotificationDigestPreview parseKDENotificationDigestPreviewSource repeatedNotificationDigestEvents].each do |token|
  assert(go_runtime_notification_digest_cli_source.include?(token), "Go Runtime KDE notification digest CLI must include #{token}")
end

go_runtime_notification_digest_cli_test_source = read_project_file("cmd/xnix-runtime-go/kde_notification_digest_cli_test.go")
%w[TestKDENotificationDigestPreviewCLI TestKDENotificationDigestPreviewCLIRejectsInvalidArguments xnix.runtime.kde_notification_digest.v1 notifications_sent live_tray_bridge_enabled request_objects_created permission_grant_enabled backend_launch_enabled host_root_modified].each do |token|
  assert(go_runtime_notification_digest_cli_test_source.include?(token), "Go Runtime KDE notification digest CLI tests must include #{token}")
end

go_runtime_offline_kde_identity_source = read_project_file("internal/runtime/appidentity/kde_offline_application_identity.go")
%w[KDEOfflineApplicationIdentityPreview KDEOfflineDesktopEntryIdentity KDEOfflineMIMEIdentity KDEOfflineKRunnerIdentity KDEOfflineTaskManagerIdentity KDEOfflineKWinIdentity KDEOfflineTrayIdentity KDEOfflineNotificationIdentity KDEOfflineSettingsIdentity KDEOfflineCenterIdentity xnix.runtime.kde_offline_application_identity.v1 kde-offline-application-identity-preview GetKDEOfflineApplicationIdentity GetKDEOfflineApplicationIdentityPreview desktop-entry mime-associations krunner task-manager kwin system-tray notification-center unified-settings compatibility-center CrossSurfaceIdentityConsistent NotificationIDNamespace].each do |token|
  assert(go_runtime_offline_kde_identity_source.include?(token), "Go Runtime offline KDE application identity must include #{token}")
end
%w[DesktopFilesWritten MIMEDefaultsWritten KRunnerIndexPersisted TaskManagerEntryActive KWinRuleApplied LiveTrayBridgeEnabled TrayBridgePersisted NotificationSent NotificationDeliveryEnabled NotificationActionsEnabled SettingsPersisted SettingsPersistenceEnabled CompatibilityCenterPersisted CompatibilityCenterActionsEnabled LaunchEnabled ExecutionStarted BackendProcessStarted NetworkRequired HostRootModified RawCommandExposed BackendDetailsExposed].each do |token|
  assert(go_runtime_offline_kde_identity_source.include?(token), "Go Runtime offline KDE application identity must expose disabled gate #{token}")
end

go_runtime_offline_kde_identity_test_source = read_project_file("internal/runtime/appidentity/kde_offline_application_identity_test.go")
%w[TestKDEOfflineApplicationIdentityPreviewJoinsCanonicalKDESurfaces TestKDEOfflineApplicationIdentityPreviewUsesDeterministicAttentionIdentity TestKDEOfflineApplicationIdentityPreviewJoinsSettingsAndCenter TestKDEOfflineApplicationIdentityPreviewRequiresVerifiedRegistryRecipe CrossSurfaceIdentityConsistent].each do |token|
  assert(go_runtime_offline_kde_identity_test_source.include?(token), "Go Runtime offline KDE application identity tests must include #{token}")
end

runtime_owner_offline_kde_identity_test_source = read_project_file("internal/runtime/owner/dispatch.go") +
                                                  read_project_file("internal/runtime/owner/dispatch_test.go") +
                                                  read_project_file("internal/runtime/owner/service_test.go") +
                                                  read_project_file("cmd/xnix-runtime-owner/main_test.go")
%w[GetKDEOfflineApplicationIdentityPreview kde-offline-application-identity-preview NewKDEOfflineApplicationIdentityPreview TestDispatchReadRendersOfflineKDEIdentityAsOwnerLocalPayload TestServiceCallServesOwnerLocalOfflineKDEIdentity TestRuntimeOwnerCommandRendersOfflineKDEIdentityOwnerLocalReadDispatch surface_count cross_surface_identity_consistent settings_persisted compatibility_center_persisted launch_enabled host_root_modified].each do |token|
  assert(runtime_owner_offline_kde_identity_test_source.include?(token), "Go Runtime owner offline KDE identity evidence must include #{token}")
end

go_runtime_offline_kde_identity_cli_source = read_project_file("cmd/xnix-runtime-go/kde_offline_application_identity_commands.go") +
                                         read_project_file("cmd/xnix-runtime-go/kde_offline_application_identity_cli_test.go") +
                                         read_project_file("cmd/xnix-runtime-go/main.go")
%w[kde-offline-application-identity-preview runKDEOfflineApplicationIdentityPreview parseKDEOfflineApplicationIdentityPreviewSource LoadRecipeFromRegistry TestKDEOfflineApplicationIdentityPreviewCommandUsesSampleFixture TestKDEOfflineApplicationIdentityPreviewCommandRejectsIncompleteArguments xnix.runtime.kde_offline_application_identity.v1].each do |token|
  assert(go_runtime_offline_kde_identity_cli_source.include?(token), "Go Runtime offline KDE application identity CLI must include #{token}")
end

go_runtime_kde_fake_execution_source = read_project_file("internal/runtime/appidentity/kde_fake_execution_evidence.go")
%w[KDEFakeExecutionEvidenceRecord KDEFakeExecutionEvidenceOptions NewKDEFakeExecutionEvidenceRecord xnix.runtime.kde_fake_execution_evidence.v1 kde-fake-execution-evidence-record test-only explicit-test-root-only prepareFakeExecutionLifecycle execution-ledger execution-session-record lifecycle-staged StateRootWritesEnabled StateRootRecordCount FakeExecutionRecorded AllChecksPassed].each do |token|
  assert(go_runtime_kde_fake_execution_source.include?(token), "Go Runtime KDE fake execution evidence must include #{token}")
end
%w[StateRootPathExposed LaunchAllowed LaunchEnabled ExecutionStarted BackendProcessStarted RealPortalCallEnabled ProductionBusOwnership NetworkRequired HostRootModified PrivilegedContainerRequired BackendDetailsExposed].each do |token|
  assert(go_runtime_kde_fake_execution_source.include?(token), "Go Runtime KDE fake execution evidence must expose safety gate #{token}")
end

go_runtime_kde_fake_execution_test_source = read_project_file("internal/runtime/appidentity/kde_fake_execution_evidence_test.go")
%w[TestKDEFakeExecutionEvidencePersistsAndReadsBackControlledRecords TestKDEFakeExecutionEvidenceRejectsUnsafeInputs TestKDEFakeExecutionEvidenceRejectsManagedPathSymlink staged blocked explicit-test-root-only].each do |token|
  assert(go_runtime_kde_fake_execution_test_source.include?(token), "Go Runtime KDE fake execution tests must include #{token}")
end

go_runtime_kde_fake_execution_cli_source = read_project_file("cmd/xnix-runtime-go/kde_fake_execution_evidence_commands.go") +
                                             read_project_file("cmd/xnix-runtime-go/kde_fake_execution_evidence_cli_test.go") +
                                             read_project_file("cmd/xnix-runtime-go/main.go")
%w[kde-fake-execution-evidence-record runKDEFakeExecutionEvidenceRecord LoadRecipeFromRegistry NewKDEFakeExecutionEvidenceRecord test-only TestKDEFakeExecutionEvidenceRecordCommandWritesControlledEvidence TestKDEFakeExecutionEvidenceRecordCommandRequiresExplicitTestBoundary all_checks_passed state_root_writes_enabled launch_enabled backend_process_started host_root_modified].each do |token|
  assert(go_runtime_kde_fake_execution_cli_source.include?(token), "Go Runtime KDE fake execution CLI must include #{token}")
end

go_runtime_kde_fake_portal_source = read_project_file("internal/runtime/appidentity/kde_fake_portal_evidence.go")
%w[KDEFakePortalEvidenceRecord NewKDEFakePortalEvidenceRecord xnix.runtime.kde_fake_portal_evidence.v1 kde-fake-portal-evidence-record ensureCompletedFakePortalReceipt reviewedFakeExecutionTransaction portal-gate-delta portal-policy-review snapshot-baseline PortalEvidenceRecorded PortalGateChangedOnly StateRootRecordCount].each do |token|
  assert(go_runtime_kde_fake_portal_source.include?(token), "Go Runtime KDE fake Portal evidence must include #{token}")
end
%w[StateRootPathExposed RealPortalCallEnabled HostPermissionChanged ExecutionApproved LaunchAllowed LaunchEnabled ExecutionStarted BackendProcessStarted ProductionBusOwnership NetworkRequired HostRootModified PrivilegedContainerRequired BackendDetailsExposed].each do |token|
  assert(go_runtime_kde_fake_portal_source.include?(token), "Go Runtime KDE fake Portal evidence must expose safety gate #{token}")
end

go_runtime_kde_fake_portal_test_source = read_project_file("internal/runtime/appidentity/kde_fake_portal_evidence_test.go")
%w[TestKDEFakePortalEvidenceChangesOnlyPortalGate TestKDEFakePortalEvidenceRejectsManagedPathSymlink pending-user-mediation granted completed explicit-test-root-only].each do |token|
  assert(go_runtime_kde_fake_portal_test_source.include?(token), "Go Runtime KDE fake Portal tests must include #{token}")
end

go_runtime_kde_fake_portal_cli_source = read_project_file("cmd/xnix-runtime-go/kde_fake_portal_evidence_commands.go") +
                                          read_project_file("cmd/xnix-runtime-go/kde_fake_portal_evidence_cli_test.go") +
                                          read_project_file("cmd/xnix-runtime-go/main.go")
%w[kde-fake-portal-evidence-record runKDEFakePortalEvidenceRecord LoadRecipeFromRegistry NewKDEFakePortalEvidenceRecord test-only TestKDEFakePortalEvidenceRecordCommandJoinsPortalReceipt TestKDEFakePortalEvidenceRecordCommandRequiresTestOnlyBoundary portal_gate_changed_only portal_evidence_recorded real_portal_call_enabled host_permission_changed execution_approved].each do |token|
  assert(go_runtime_kde_fake_portal_cli_source.include?(token), "Go Runtime KDE fake Portal CLI must include #{token}")
end

go_runtime_kde_snapshot_diagnostics_source = read_project_file("internal/runtime/appidentity/kde_snapshot_diagnostics_evidence.go")
%w[KDESnapshotDiagnosticsEvidenceRecord NewKDESnapshotDiagnosticsEvidenceRecord xnix.runtime.kde_snapshot_diagnostics_evidence.v1 kde-snapshot-diagnostics-evidence-record stability-diagnostics-v0.2.315 stability-baseline-v0.2.315 recordStabilityDiagnostic ensureStabilitySnapshot portal-evidence diagnostic-readback diagnostic-redaction snapshot-baseline lifecycle-ready execution-readiness session-readback unsafe-gates-closed CoreReceiptCount AllChecksPassed].each do |token|
  assert(go_runtime_kde_snapshot_diagnostics_source.include?(token), "Go Runtime KDE snapshot diagnostics evidence must include #{token}")
end
%w[StateRootPathExposed SnapshotCreationEnabled SnapshotRestoreEnabled SnapshotDeletionEnabled DiagnosticExecutionEnabled AIProviderCallEnabled RepairExecutionEnabled RealPortalCallEnabled HostPermissionChanged ExecutionApproved LaunchAllowed LaunchEnabled ExecutionStarted BackendProcessStarted ProductionBusOwnership NetworkRequired HostRootModified PrivilegedContainerRequired BackendDetailsExposed FileContentsExposed].each do |token|
  assert(go_runtime_kde_snapshot_diagnostics_source.include?(token), "Go Runtime KDE snapshot diagnostics evidence must expose safety gate #{token}")
end

go_runtime_kde_snapshot_diagnostics_test_source = read_project_file("internal/runtime/appidentity/kde_snapshot_diagnostics_evidence_test.go")
%w[TestKDESnapshotDiagnosticsEvidenceConvergesControlledState TestKDESnapshotDiagnosticsEvidenceRejectsUnsafeBoundaries ready blocked explicit-test-root-only .xnix-snapshots diagnostics-ledger test-results].each do |token|
  assert(go_runtime_kde_snapshot_diagnostics_test_source.include?(token), "Go Runtime KDE snapshot diagnostics tests must include #{token}")
end

go_runtime_kde_snapshot_diagnostics_cli_source = read_project_file("cmd/xnix-runtime-go/kde_snapshot_diagnostics_evidence_commands.go") +
                                                  read_project_file("cmd/xnix-runtime-go/kde_snapshot_diagnostics_evidence_cli_test.go") +
                                                  read_project_file("cmd/xnix-runtime-go/main.go")
%w[kde-snapshot-diagnostics-evidence-record runKDESnapshotDiagnosticsEvidenceRecord LoadRecipeFromRegistry NewKDESnapshotDiagnosticsEvidenceRecord test-only TestKDESnapshotDiagnosticsEvidenceRecordCommandConvergesEvidence TestKDESnapshotDiagnosticsEvidenceRecordCommandRequiresTestOnlyBoundary snapshot_restore_enabled diagnostic_execution_enabled ai_provider_call_enabled repair_execution_enabled launch_enabled backend_process_started host_root_modified file_contents_exposed].each do |token|
  assert(go_runtime_kde_snapshot_diagnostics_cli_source.include?(token), "Go Runtime KDE snapshot diagnostics CLI must include #{token}")
end

go_runtime_kde_backend_lifecycle_source = read_project_file("internal/runtime/appidentity/kde_backend_lifecycle_evidence.go")
%w[KDEBackendLifecycleEvidenceRecord NewKDEBackendLifecycleEvidenceRecord xnix.runtime.kde_backend_lifecycle_evidence.v1 kde-backend-lifecycle-evidence-record prerequisite-convergence inventory-readback backend-state lifecycle-join execution-readback session-readback backend-boundary unsafe-gates-closed BackendStateJoined CoreReceiptCount AllChecksPassed].each do |token|
  assert(go_runtime_kde_backend_lifecycle_source.include?(token), "Go Runtime KDE backend lifecycle evidence must include #{token}")
end
%w[StateRootPathExposed BackendKindsExposedToKDE BackendInstallEnabled BackendDownloadEnabled BackendLaunchEnabled BackendProcessStarted VMProcessStarted RawCommandExposed ProfilePathExposed RealPortalCallEnabled SnapshotRestoreEnabled DiagnosticExecutionEnabled AIProviderCallEnabled RepairExecutionEnabled ExecutionApproved LaunchAllowed LaunchEnabled ExecutionStarted ProductionBusOwnership NetworkRequired HostRootModified PrivilegedContainerRequired BackendDetailsExposed SecretsExposed].each do |token|
  assert(go_runtime_kde_backend_lifecycle_source.include?(token), "Go Runtime KDE backend lifecycle evidence must expose safety gate #{token}")
end

go_runtime_kde_backend_lifecycle_test_source = read_project_file("internal/runtime/appidentity/kde_backend_lifecycle_evidence_test.go")
%w[TestKDEBackendLifecycleEvidenceJoinsInventoryWithoutStartingProcess TestKDEBackendLifecycleEvidenceRejectsUnsafeBoundary ready blocked explicit-test-root-only backend-manager].each do |token|
  assert(go_runtime_kde_backend_lifecycle_test_source.include?(token), "Go Runtime KDE backend lifecycle tests must include #{token}")
end

go_runtime_kde_backend_lifecycle_cli_source = read_project_file("cmd/xnix-runtime-go/kde_backend_lifecycle_evidence_commands.go") +
                                              read_project_file("cmd/xnix-runtime-go/kde_backend_lifecycle_evidence_cli_test.go") +
                                              read_project_file("cmd/xnix-runtime-go/main.go")
%w[kde-backend-lifecycle-evidence-record runKDEBackendLifecycleEvidenceRecord LoadRecipeFromRegistry NewKDEBackendLifecycleEvidenceRecord test-only TestKDEBackendLifecycleEvidenceRecordCommandJoinsInventory TestKDEBackendLifecycleEvidenceRecordCommandRequiresTestOnlyBoundary backend_kinds_exposed_to_kde backend_process_started host_root_modified secrets_exposed].each do |token|
  assert(go_runtime_kde_backend_lifecycle_cli_source.include?(token), "Go Runtime KDE backend lifecycle CLI must include #{token}")
end

go_runtime_kde_restricted_authorization_source = read_project_file("internal/runtime/appidentity/kde_restricted_launch_authorization.go")
%w[KDERestrictedLaunchAuthorizationRecord NewKDERestrictedLaunchAuthorizationRecord xnix.runtime.kde_restricted_launch_authorization.v1 kde-restricted-launch-authorization-record prerequisite-convergence explicit-test-boundary authorization-readback trust-independent write-gate-independent execution-unchanged session-unchanged unsafe-gates-closed AuthorizationBoundaryJoined CoreReceiptCount AllChecksPassed].each do |token|
  assert(go_runtime_kde_restricted_authorization_source.include?(token), "Go Runtime KDE restricted launch authorization must include #{token}")
end
%w[StateRootPathExposed ProductionTrustSatisfied RuntimeWriteGateEnabled ArtifactAcquisitionEnabled BackendInstallEnabled BackendLaunchEnabled BackendProcessStarted RealPortalCallEnabled ExecutionApproved LaunchAuthorized LaunchAllowed LaunchEnabled ExecutionStarted ProcessStartAuthorized ProductionBusOwnership NetworkRequired HostRootModified PrivilegedContainerRequired RawCommandExposed BackendDetailsExposed].each do |token|
  assert(go_runtime_kde_restricted_authorization_source.include?(token), "Go Runtime KDE restricted launch authorization must expose safety gate #{token}")
end

go_runtime_kde_restricted_authorization_test_source = read_project_file("internal/runtime/appidentity/kde_restricted_launch_authorization_test.go")
%w[TestKDERestrictedLaunchAuthorizationRecordsPreparationOnly TestKDERestrictedLaunchAuthorizationRequiresExactDirective authorized-preparation-only explicit-test-root-only].each do |token|
  assert(go_runtime_kde_restricted_authorization_test_source.include?(token), "Go Runtime KDE restricted launch authorization tests must include #{token}")
end

go_runtime_kde_restricted_authorization_cli_source = read_project_file("cmd/xnix-runtime-go/kde_restricted_launch_authorization_commands.go") +
                                                   read_project_file("cmd/xnix-runtime-go/kde_restricted_launch_authorization_cli_test.go") +
                                                   read_project_file("cmd/xnix-runtime-go/main.go")
%w[kde-restricted-launch-authorization-record runKDERestrictedLaunchAuthorizationRecord authorize-restricted-test-preparation test-only TestKDERestrictedLaunchAuthorizationRecordCommandRequiresExplicitDirective TestKDERestrictedLaunchAuthorizationRecordCommandRejectsImplicitAuthorization preparation_authorized launch_authorized process_start_authorized].each do |token|
  assert(go_runtime_kde_restricted_authorization_cli_source.include?(token), "Go Runtime KDE restricted launch authorization CLI must include #{token}")
end

go_runtime_restricted_preflight_source = read_project_file("internal/runtime/execution/restricted_preflight.go")
%w[RestrictedPreflightStore RestrictedPreflightPacket xnix.runtime.restricted_launch_preflight.v1 restricted-launch-preflight-packet go-runtime-state-root-restricted-launch-preflight blocked recipe-trust runtime-write-gate ReadyForPacketAssembly ProductImageReady LaunchPreflightPassed].each do |token|
  assert(go_runtime_restricted_preflight_source.include?(token), "Go Runtime restricted launch preflight store must include #{token}")
end
%w[LaunchAuthorized ExecutionApproved ProcessStartAuthorized CommandMaterialized ExecutablePathResolved BackendSelectedForLaunch BackendLaunchEnabled BackendProcessStarted NetworkRequired HostRootModified PrivilegedContainerRequired].each do |token|
  assert(go_runtime_restricted_preflight_source.include?(token), "Go Runtime restricted launch preflight store must expose safety gate #{token}")
end

go_runtime_restricted_preflight_test_source = read_project_file("internal/runtime/execution/restricted_preflight_test.go")
%w[TestRestrictedPreflightStorePersistsFailClosedPacket TestRestrictedPreflightStoreRejectsTamperedPacket recipe-trust runtime-write-gate].each do |token|
  assert(go_runtime_restricted_preflight_test_source.include?(token), "Go Runtime restricted launch preflight tests must include #{token}")
end

go_runtime_kde_restricted_preflight_source = read_project_file("internal/runtime/appidentity/kde_restricted_launch_preflight.go")
%w[KDERestrictedLaunchPreflightRecord NewKDERestrictedLaunchPreflightRecord xnix.runtime.kde_restricted_launch_preflight.v1 kde-restricted-launch-preflight-record authorization-boundary safe-inputs preflight-readback explicit-blockers packet-assembly-boundary execution-unchanged session-unchanged unsafe-gates-closed PreflightBoundaryJoined CoreReceiptCount AllChecksPassed].each do |token|
  assert(go_runtime_kde_restricted_preflight_source.include?(token), "Go Runtime KDE restricted launch preflight must include #{token}")
end
%w[ProductImageReady ProductionTrustSatisfied RuntimeWriteGateEnabled LaunchPreflightPassed LaunchAuthorized ExecutionApproved ProcessStartAuthorized CommandMaterialized ExecutablePathResolved BackendSelectedForLaunch BackendLaunchEnabled BackendProcessStarted ProductionBusOwnership NetworkRequired HostRootModified PrivilegedContainerRequired RawCommandExposed BackendDetailsExposed].each do |token|
  assert(go_runtime_kde_restricted_preflight_source.include?(token), "Go Runtime KDE restricted launch preflight must expose safety gate #{token}")
end

go_runtime_kde_restricted_preflight_test_source = read_project_file("internal/runtime/appidentity/kde_restricted_launch_preflight_test.go")
%w[TestKDERestrictedLaunchPreflightRemainsBlocked recipe-trust runtime-write-gate explicit-test-root-only].each do |token|
  assert(go_runtime_kde_restricted_preflight_test_source.include?(token), "Go Runtime KDE restricted launch preflight tests must include #{token}")
end

go_runtime_kde_restricted_preflight_cli_source = read_project_file("cmd/xnix-runtime-go/kde_restricted_launch_preflight_commands.go") +
                                                read_project_file("cmd/xnix-runtime-go/kde_restricted_launch_preflight_cli_test.go") +
                                                read_project_file("cmd/xnix-runtime-go/main.go")
%w[kde-restricted-launch-preflight-record runKDERestrictedLaunchPreflightRecord authorize-restricted-test-preparation test-only TestKDERestrictedLaunchPreflightRecordCommandIsFailClosed TestKDERestrictedLaunchPreflightRecordCommandRequiresAuthorization ready_for_packet_assembly product_image_ready launch_preflight_passed].each do |token|
  assert(go_runtime_kde_restricted_preflight_cli_source.include?(token), "Go Runtime KDE restricted launch preflight CLI must include #{token}")
end

go_runtime_restricted_materialization_source = read_project_file("internal/runtime/execution/restricted_materialization.go")
%w[RestrictedMaterializationStore RestrictedMaterializationPlan xnix.runtime.restricted_launch_materialization.v1 restricted-launch-materialization-plan go-runtime-state-root-restricted-launch-materialization blocked-plan-materialized test-only-review-plan PlanMaterialized CommandMaterialized ExecutablePathResolved BackendSelectedForLaunch BackendLaunchEnabled BackendProcessStarted RawExecutableExposed HostRootModified].each do |token|
  assert(go_runtime_restricted_materialization_source.include?(token), "Go Runtime restricted launch materialization store must include #{token}")
end

go_runtime_restricted_materialization_test_source = read_project_file("internal/runtime/execution/restricted_materialization_test.go")
%w[TestRestrictedMaterializationStorePersistsPlanOnly TestRestrictedMaterializationStoreRejectsTamperedPlan recipe-trust runtime-write-gate command_materialized].each do |token|
  assert(go_runtime_restricted_materialization_test_source.include?(token), "Go Runtime restricted launch materialization tests must include #{token}")
end

go_runtime_kde_test_launch_materialization_source = read_project_file("internal/runtime/appidentity/kde_test_launch_materialization.go")
%w[KDETestLaunchMaterializationRecord NewKDETestLaunchMaterializationRecord xnix.runtime.kde_test_launch_materialization.v1 kde-test-launch-materialization-record preflight-boundary materialization-readback materialized-plan-only write-gate-remains-blocked command-boundary launch-boundary execution-unchanged session-unchanged unsafe-gates-closed MaterializationBoundary PlanMaterialized CoreReceiptCount AllChecksPassed].each do |token|
  assert(go_runtime_kde_test_launch_materialization_source.include?(token), "Go Runtime KDE test launch materialization must include #{token}")
end
%w[ProductImageReady ProductionTrustSatisfied RuntimeWriteGateEnabled LaunchPreflightPassed LaunchAuthorized ExecutionApproved ProcessStartAuthorized CommandMaterialized ExecutablePathResolved BackendSelectedForLaunch BackendLaunchEnabled BackendProcessStarted ProductionBusOwnership NetworkRequired HostRootModified PrivilegedContainerRequired RawCommandExposed RawExecutableExposed BackendDetailsExposed].each do |token|
  assert(go_runtime_kde_test_launch_materialization_source.include?(token), "Go Runtime KDE test launch materialization must expose safety gate #{token}")
end

go_runtime_kde_test_launch_materialization_test_source = read_project_file("internal/runtime/appidentity/kde_test_launch_materialization_test.go")
%w[TestKDETestLaunchMaterializationRecordMaterializesPlanOnly TestKDETestLaunchMaterializationRecordRequiresExactBoundary blocked-plan-materialized test-only-review-plan explicit-test-root-only recipe-trust runtime-write-gate].each do |token|
  assert(go_runtime_kde_test_launch_materialization_test_source.include?(token), "Go Runtime KDE test launch materialization tests must include #{token}")
end

go_runtime_kde_test_launch_materialization_cli_source = read_project_file("cmd/xnix-runtime-go/kde_test_launch_materialization_commands.go") +
                                                       read_project_file("cmd/xnix-runtime-go/kde_test_launch_materialization_cli_test.go") +
                                                       read_project_file("cmd/xnix-runtime-go/main.go")
%w[kde-test-launch-materialization-record runKDETestLaunchMaterializationRecord authorize-restricted-test-preparation test-only TestKDETestLaunchMaterializationRecordCommandMaterializesPlanOnly TestKDETestLaunchMaterializationRecordCommandRequiresAuthorization plan_materialized command_materialized raw_executable_exposed].each do |token|
  assert(go_runtime_kde_test_launch_materialization_cli_source.include?(token), "Go Runtime KDE test launch materialization CLI must include #{token}")
end

go_runtime_kde_test_launch_materialization_fanout_source = read_project_file("internal/runtime/appidentity/kde_test_launch_materialization_fanout.go")
%w[KDETestLaunchMaterializationFanOutPreview KDETestLaunchMaterializationFanSurface KDETestLaunchMaterializationFanOutOptions NewKDETestLaunchMaterializationFanOutPreview NewKDETestLaunchMaterializationFanOutPreviewFromReceipt LoadKDETestLaunchMaterializationRecordForFanOut xnix.runtime.kde_test_launch_materialization_fanout.v1 kde-test-launch-materialization-fanout-preview GetKDETestLaunchMaterializationFanOut GetKDETestLaunchMaterializationFanOutPreview materialization-consumed session-fanout-consumed surface-coverage surfaces-read-only desktop-side-effects-disabled notification-not-delivered launch-boundary-closed unsafe-data-hidden host-boundary-closed compatibility-center task-manager tray notification review-plan-available read-only-existing-materialization-receipt read-only-receipt-consumption materialization-receipt-consumed execution-receipt-consumed session-receipt-consumed read-only-consumption].each do |token|
  assert(go_runtime_kde_test_launch_materialization_fanout_source.include?(token), "Go Runtime KDE test launch materialization fan-out must include #{token}")
end
%w[MaterializationReceiptConsumed ExecutionSessionFanOutConsumed FanOutWritesEnabled StateRootPathExposed ProductImageReady ProductionTrustSatisfied RuntimeWriteGateEnabled LaunchAuthorized ExecutionApproved CommandMaterialized ExecutablePathResolved BackendSelectedForLaunch BackendLaunchEnabled BackendProcessStarted TaskManagerEntryActive LiveTrayBridgeEnabled NotificationSent NotificationDeliveryEnabled CompatibilityCenterActionsEnabled RequestObjectsCreated RuntimeWritesEnabled ProductionBusOwnership NetworkRequired HostRootModified PrivilegedContainerRequired RawCommandExposed RawExecutableExposed BackendDetailsExposed].each do |token|
  assert(go_runtime_kde_test_launch_materialization_fanout_source.include?(token), "Go Runtime KDE test launch materialization fan-out must expose safety gate #{token}")
end
assert(go_runtime_kde_test_launch_materialization_fanout_source.include?("NewKDETestLaunchMaterializationRecord"), "Go Runtime KDE test launch materialization fan-out must consume the materialization record")
assert(go_runtime_kde_test_launch_materialization_fanout_source.include?("ExecutionSessionFanOutEvidence"), "Go Runtime KDE test launch materialization fan-out must consume execution session fan-out evidence")
assert(go_runtime_kde_test_launch_materialization_fanout_source.include?("validateNoBackendTerms"), "Go Runtime KDE test launch materialization fan-out must hide backend terms")

go_runtime_kde_test_launch_materialization_fanout_test_source = read_project_file("internal/runtime/appidentity/kde_test_launch_materialization_fanout_test.go")
%w[TestKDETestLaunchMaterializationFanOutPreviewCoversKDESurfaces TestKDETestLaunchMaterializationFanOutPreviewConsumesExistingReceiptReadOnly TestKDETestLaunchMaterializationFanOutPreviewFromReceiptRejectsMissingPlan TestKDETestLaunchMaterializationFanOutPreviewRequiresExactBoundary compatibility-center task-manager tray notification review-plan-available read-only-existing-materialization-receipt].each do |token|
  assert(go_runtime_kde_test_launch_materialization_fanout_test_source.include?(token), "Go Runtime KDE test launch materialization fan-out tests must include #{token}")
end

go_runtime_kde_test_launch_materialization_fanout_cli_source = read_project_file("cmd/xnix-runtime-go/kde_test_launch_materialization_fanout_commands.go") +
                                                              read_project_file("cmd/xnix-runtime-go/kde_test_launch_materialization_fanout_cli_test.go") +
                                                              read_project_file("cmd/xnix-runtime-go/kde_test_launch_materialization_owner_route_audit_cli_test.go") +
                                                              read_project_file("cmd/xnix-runtime-go/main.go")
%w[kde-test-launch-materialization-fanout-preview kde-test-launch-materialization-fanout-consume-preview kde-test-launch-materialization-receipt-lookup-preview kde-test-launch-materialization-fanout-owner-route-preview runKDETestLaunchMaterializationFanOutPreview runKDETestLaunchMaterializationFanOutConsumePreview runKDETestLaunchMaterializationReceiptLookupPreview runKDETestLaunchMaterializationFanOutOwnerRoutePreview authorize-restricted-test-preparation test-only materialization-plan-id TestKDETestLaunchMaterializationFanOutPreviewCommandCoversKDESurfaces TestKDETestLaunchMaterializationFanOutConsumePreviewCommandReadsExistingReceipt TestKDETestLaunchMaterializationFanOutConsumePreviewCommandRequiresExistingReceipt TestKDETestLaunchMaterializationReceiptLookupPreviewCommand TestKDETestLaunchMaterializationReceiptLookupPreviewCommandRejectsBadInputs TestKDETestLaunchMaterializationFanOutOwnerRoutePreviewCommand TestKDETestLaunchMaterializationFanOutOwnerRoutePreviewCommandRejectsBadInputs TestKDETestLaunchMaterializationFanOutPreviewCommandRequiresAuthorization materialization_receipt_consumed execution_session_fan_out_consumed notification_sent fan_out_writes_enabled state_root_writes_enabled read-only-existing-materialization-receipt owner_managed_opaque_receipt_lookup_ready].each do |token|
  assert(go_runtime_kde_test_launch_materialization_fanout_cli_source.include?(token), "Go Runtime KDE test launch materialization fan-out CLI must include #{token}")
end

go_runtime_kde_test_launch_materialization_receipt_lookup_source = read_project_file("internal/runtime/appidentity/kde_test_launch_materialization_receipt_lookup.go")
%w[KDETestLaunchMaterializationReceiptLookupPreview ResolveKDETestLaunchMaterializationReceipt KDETestLaunchMaterializationOpaqueReceiptID xnix.runtime.kde_test_launch_materialization_receipt_lookup.v1 kde-test-launch-materialization-receipt-lookup-preview owner-managed-materialization-receipt-lookup GetKDETestLaunchMaterializationReceiptLookup GetKDETestLaunchMaterializationReceiptLookupPreview opaque_materialization_receipt_id kde-test-launch-materialization-receipt-id missing-receipt opaque-materialization-receipt-id-supported owner-managed-opaque-lookup caller-paths-hidden missing-receipt-fails-closed read-only-lookup unsafe-gates-closed].each do |token|
  assert(go_runtime_kde_test_launch_materialization_receipt_lookup_source.include?(token), "Go Runtime KDE test launch materialization receipt lookup must include #{token}")
end
%w[RequiresCallerRegistryPath RequiresCallerApplicationID RequiresCallerStateRoot StateRootWritesEnabled RuntimeWritesEnabled MaterializationWritesEnabled FanOutWritesEnabled ProductionBusClaimed WriteMethodsEnabled LaunchAuthorized ExecutionApproved ProcessStartAuthorized CommandMaterialized ExecutablePathResolved BackendSelectedForLaunch BackendLaunchEnabled HostRootModified StateRootPathExposed BackendDetailsExposed].each do |token|
  assert(go_runtime_kde_test_launch_materialization_receipt_lookup_source.include?(token), "Go Runtime KDE test launch materialization receipt lookup must expose safety gate #{token}")
end
go_runtime_kde_test_launch_materialization_receipt_lookup_test_source = read_project_file("internal/runtime/appidentity/kde_test_launch_materialization_receipt_lookup_test.go")
%w[TestResolveKDETestLaunchMaterializationReceiptReturnsOpaqueLookup TestResolveKDETestLaunchMaterializationReceiptRejectsUnknownOpaqueID owner-managed-materialization-receipt-lookup kde-test-launch-materialization-receipt-id missing-receipt].each do |token|
  assert(go_runtime_kde_test_launch_materialization_receipt_lookup_test_source.include?(token), "Go Runtime KDE test launch materialization receipt lookup tests must include #{token}")
end

go_runtime_kde_test_launch_materialization_fanout_owner_route_source = read_project_file("internal/runtime/appidentity/kde_test_launch_materialization_fanout_owner_route.go")
%w[KDETestLaunchMaterializationFanOutOwnerRoutePreview NewKDETestLaunchMaterializationFanOutOwnerRoutePreview xnix.runtime.kde_test_launch_materialization_fanout_owner_route.v1 kde-test-launch-materialization-fanout-owner-route-preview owner-local-kde-test-launch-materialization-fanout GetKDETestLaunchMaterializationFanOut GetKDETestLaunchMaterializationFanOutPreview opaque_materialization_receipt_id kde-test-launch-materialization-receipt-id missing-receipt-fail-closed opaque-lookup-consumed caller-paths-hidden missing-receipt-fails-closed surface-fanout-deferred desktop-side-effects-disabled owner-route-ready-production-dbus-blocked unsafe-gates-closed].each do |token|
  assert(go_runtime_kde_test_launch_materialization_fanout_owner_route_source.include?(token), "Go Runtime KDE test launch materialization fan-out owner route must include #{token}")
end
%w[RequiresCallerRegistryPath RequiresCallerApplicationID RequiresCallerStateRoot MaterializationReceiptConsumed ExecutionSessionFanOutConsumed StateRootWritesEnabled RuntimeWritesEnabled FanOutWritesEnabled ProductImageReady ProductionTrustSatisfied RuntimeWriteGateEnabled LaunchAuthorized ExecutionApproved CommandMaterialized ExecutablePathResolved BackendSelectedForLaunch BackendLaunchEnabled BackendProcessStarted TaskManagerEntryActive LiveTrayBridgeEnabled NotificationSent NotificationDeliveryEnabled CompatibilityCenterActionsEnabled RequestObjectsCreated ProductionBusOwnership NetworkRequired HostRootModified PrivilegedContainerRequired RawCommandExposed RawExecutableExposed BackendDetailsExposed].each do |token|
  assert(go_runtime_kde_test_launch_materialization_fanout_owner_route_source.include?(token), "Go Runtime KDE test launch materialization fan-out owner route must expose safety gate #{token}")
end
go_runtime_kde_test_launch_materialization_fanout_owner_route_test_source = read_project_file("internal/runtime/appidentity/kde_test_launch_materialization_fanout_owner_route_test.go")
%w[TestKDETestLaunchMaterializationFanOutOwnerRoutePreviewUsesOpaqueLookup TestKDETestLaunchMaterializationFanOutOwnerRouteRejectsUnknownOpaqueID owner-local-kde-test-launch-materialization-fanout kde-test-launch-materialization-receipt-id missing-receipt-fail-closed surface-fanout-deferred].each do |token|
  assert(go_runtime_kde_test_launch_materialization_fanout_owner_route_test_source.include?(token), "Go Runtime KDE test launch materialization fan-out owner route tests must include #{token}")
end

go_runtime_kde_test_launch_materialization_owner_route_audit_source = read_project_file("internal/runtime/appidentity/kde_test_launch_materialization_owner_route_audit.go")
%w[KDETestLaunchMaterializationOwnerRouteAuditPreview NewKDETestLaunchMaterializationOwnerRouteAuditPreview xnix.runtime.kde_test_launch_materialization_owner_route_audit.v1 kde-test-launch-materialization-owner-route-audit-preview GetKDETestLaunchMaterializationOwnerRouteAudit GetKDETestLaunchMaterializationOwnerRouteAuditPreview materialization-fanout-owner-route-audit GetKDETestLaunchMaterializationFanOut consume-ready-opaque-lookup-missing read-only-consume-ready-owner-route-blocked opaque-lookup-ready-owner-route-blocked owner-managed-lookup-ready-fanout-cli-only owner-local-route-ready owner-local-read-route-ready-production-dbus-blocked owner-local-route-smoke-covered owner-local-read-route-smoke-covered-production-dbus-blocked materialization-fanout-owner-smoke-coverage owner-smoke-coverage-present owner-managed-materialization-receipt-lookup cli-preview-registered go-read-model-present read-only-consume-registered owner-route-present production-dbus-absent caller-path-boundary receipt-creation-split owner-managed-opaque-lookup route-decision unsafe-gates-closed].each do |token|
  assert(go_runtime_kde_test_launch_materialization_owner_route_audit_source.include?(token), "Go Runtime KDE test launch materialization owner-route audit must include #{token}")
end
%w[OwnerDispatchRoutePresent ProductionDBusMethodPresent RequiresCallerRegistryPath RequiresCallerStateRoot MaterializationWritesStateRoot ReadOnlyConsumeCommandRegistered ReadOnlyReceiptConsumptionReady ReadOnlyConsumeRequiresCallerStateRoot OwnerManagedOpaqueReceiptLookupReady OpaqueMaterializationReceiptIDSupported OwnerSmokeCoverageReady OwnerLocalRouteCandidateReady ProductionDBusExposureReady SystemServiceStarted SessionBusClaimed ProductionBusClaimed WriteMethodsEnabled RuntimeWritesEnabled BackendLaunchEnabled HostRootModified StateRootPathExposed BackendDetailsExposed].each do |token|
  assert(go_runtime_kde_test_launch_materialization_owner_route_audit_source.include?(token), "Go Runtime KDE test launch materialization owner-route audit must expose safety gate #{token}")
end
assert(go_runtime_kde_test_launch_materialization_owner_route_audit_source.include?("validateNoBackendTerms"), "Go Runtime KDE test launch materialization owner-route audit must hide backend terms")
go_runtime_kde_test_launch_materialization_owner_route_audit_test_source = read_project_file("internal/runtime/appidentity/kde_test_launch_materialization_owner_route_audit_test.go")
%w[TestKDETestLaunchMaterializationOwnerRouteAuditSeesConsumeSplit TestKDETestLaunchMaterializationOwnerRouteAuditFailsClosedWithoutSources owner-local-route-smoke-covered materialization-owner-route-sources-missing owner-route-present production-dbus-absent caller-path-boundary receipt-creation-split owner-managed-opaque-lookup owner-smoke-coverage-present].each do |token|
  assert(go_runtime_kde_test_launch_materialization_owner_route_audit_test_source.include?(token), "Go Runtime KDE test launch materialization owner-route audit tests must include #{token}")
end
%w[kde-test-launch-materialization-owner-route-audit-preview runKDETestLaunchMaterializationOwnerRouteAuditPreview TestKDETestLaunchMaterializationOwnerRouteAuditPreviewCommand route_decision read_only_receipt_consumption_ready owner_managed_opaque_receipt_lookup_ready owner_smoke_coverage_ready owner_dispatch_route_present production_dbus_method_present owner_local_route_candidate_ready].each do |token|
  assert(go_runtime_kde_test_launch_materialization_fanout_cli_source.include?(token), "Go Runtime KDE test launch materialization owner-route audit CLI must include #{token}")
end

runtime_owner_notification_digest_test_source = read_project_file("internal/runtime/owner/dispatch_test.go") +
                                                read_project_file("internal/runtime/owner/service_test.go") +
                                                read_project_file("cmd/xnix-runtime-owner/main_test.go")
%w[TestDispatchReadRendersKDENotificationDigestAsOwnerLocalPayload TestServiceCallServesOwnerLocalNotificationDigest TestRuntimeOwnerCommandRendersNotificationDigestOwnerLocalReadDispatch GetKDENotificationDigestPreview kde-notification-digest-preview notifications_sent backend_launch_enabled host_root_modified].each do |token|
  assert(runtime_owner_notification_digest_test_source.include?(token), "Go Runtime owner notification digest evidence must include #{token}")
end

runtime_owner_signed_recipe_test_source = read_project_file("internal/runtime/recipe/signature_test.go") +
                                          read_project_file("internal/runtime/owner/dispatch_test.go") +
                                          read_project_file("internal/runtime/owner/service_test.go") +
                                          read_project_file("cmd/xnix-runtime-owner/main_test.go")
%w[TestRegistrySignedRecipeVerificationPreviewFailsClosedWithoutProductionKey TestDispatchReadRendersSignedRecipeVerificationAsOwnerLocalPayload TestServiceCallServesOwnerLocalSignedRecipeVerification TestRuntimeOwnerCommandRendersSignedRecipeOwnerLocalReadDispatch GetSignedRecipeVerificationPreview signed-recipe-verifier-preview production-signature-required recipe_digest_verified signature_verified signature_material_exposed].each do |token|
  assert(runtime_owner_signed_recipe_test_source.include?(token), "Go Runtime owner signed recipe evidence must include #{token}")
end

runtime_owner_restricted_smoke_test_source = read_project_file("internal/runtime/owner/dispatch_test.go") +
                                             read_project_file("internal/runtime/owner/service_test.go") +
                                             read_project_file("cmd/xnix-runtime-owner/main_test.go")
%w[TestDispatchReadRendersRestrictedProductSmokePacketAsOwnerLocalPayload TestServiceCallServesOwnerLocalRestrictedSmokePacket TestRuntimeOwnerCommandRendersRestrictedSmokeOwnerLocalReadDispatch GetRestrictedProductSmokePacketPreview restricted-product-smoke-packet-preview human_authorization_required execution_authorized docker_executed qemu_executed release_ready].each do |token|
  assert(runtime_owner_restricted_smoke_test_source.include?(token), "Go Runtime owner restricted smoke evidence must include #{token}")
end

go_runtime_deactivation_source = read_project_file("internal/runtime/appidentity/desktop_deactivation_dry_run.go")
%w[DesktopDeactivationDryRunPreview DesktopDeactivationSurface DesktopDeactivationCounts desktop-deactivation-dry-run-preview xnix.runtime.desktop_deactivation_dry_run.v1 GetDesktopDeactivationDryRun GetDesktopDeactivationDryRunPreview launcher desktop-icon mime-association dolphin-service-menu kwin system-tray notification compatibility-center settings].each do |token|
  assert(go_runtime_deactivation_source.include?(token), "Go Runtime desktop deactivation dry run preview must include #{token}")
end
%w[missing-activation-receipt installed-file-digest-mismatch unknown-file-owner shared-mime-association active-session-evidence no-host-artifact removable].each do |token|
  assert(go_runtime_deactivation_source.include?(token), "Go Runtime desktop deactivation dry run preview must classify #{token}")
end
%w[FileDeletionEnabled MIMEDefaultsWritten KDECacheRefreshed ReceiptsRewritten SessionTerminated TargetPathExposed HostRootModified validateNoBackendTerms safety.ValidatePayload DesktopActivationReceiptEvidence DesktopActivationTransactionPreview].each do |token|
  assert(go_runtime_deactivation_source.include?(token), "Go Runtime desktop deactivation dry run preview must expose #{token}")
end

go_runtime_deactivation_test_source = read_project_file("internal/runtime/appidentity/desktop_deactivation_dry_run_test.go")
%w[TestDesktopDeactivationDryRunMissingReceiptIsBlocked TestDesktopDeactivationDryRunClassifiesReceiptErrors TestDesktopDeactivationDryRunRejectsUnsafePlan].each do |token|
  assert(go_runtime_deactivation_test_source.include?(token), "Go Runtime desktop deactivation dry run tests must include #{token}")
end

go_runtime_deactivation_cli_source = read_project_file("cmd/xnix-runtime-go/main.go") +
                                     read_project_file("cmd/xnix-runtime-go/desktop_deactivation_dry_run_commands.go")
%w[runDesktopDeactivationDryRunPreview desktop-deactivation-dry-run-preview DesktopDeactivationDryRunPreview active-session shared-mime loadRecipe].each do |token|
  assert(go_runtime_deactivation_cli_source.include?(token), "Go Runtime desktop deactivation dry run CLI must include #{token}")
end

go_runtime_deactivation_cli_test_source = read_project_file("cmd/xnix-runtime-go/desktop_deactivation_dry_run_cli_test.go")
%w[TestDesktopDeactivationDryRunPreviewCLIMissingReceipt TestDesktopDeactivationDryRunPreviewCLIReceiptBackedRemovable TestDesktopDeactivationDryRunPreviewCLIActiveSessionAndSharedMIMEBlock TestDesktopDeactivationDryRunPreviewCLIRequiresSource xnix.runtime.desktop_deactivation_dry_run.v1].each do |token|
  assert(go_runtime_deactivation_cli_test_source.include?(token), "Go Runtime desktop deactivation dry run CLI tests must include #{token}")
end

go_runtime_recipe_conflict_source = read_project_file("internal/runtime/appidentity/recipe_conflict_audit.go")
%w[RecipeConflictAuditPreview RecipeConflictGroup RecipeConflictFinding RecipeConflictCounts recipe-conflict-audit-preview xnix.runtime.recipe_conflict_audit.v1 GetRecipeConflictAudit GetRecipeConflictAuditPreview duplicate-app-ids stale-recipe-versions unsupported-capability-claims mismatched-package-source-pins trust-policy-blockers artifact-digest-drift].each do |token|
  assert(go_runtime_recipe_conflict_source.include?(token), "Go Runtime recipe conflict audit preview must include #{token}")
end
%w[RecipeWritesEnabled RegistryMigrationEnabled ArtifactStagingEnabled NetworkRequired PackageManagerInvoked BackendLaunchEnabled HostRootModified validateNoBackendTerms safety.ValidatePayload compareSemanticVersions ParseRecipe].each do |token|
  assert(go_runtime_recipe_conflict_source.include?(token), "Go Runtime recipe conflict audit preview must expose #{token}")
end

go_runtime_recipe_conflict_test_source = read_project_file("internal/runtime/appidentity/recipe_conflict_audit_test.go")
%w[TestRecipeConflictAuditCleanRegistry TestRecipeConflictAuditDetectsEveryConflictClass TestRecipeConflictAuditFlagsMissingRecipeFile TestRecipeConflictAuditRejectsMissingRegistry].each do |token|
  assert(go_runtime_recipe_conflict_test_source.include?(token), "Go Runtime recipe conflict audit tests must include #{token}")
end

go_runtime_recipe_conflict_cli_source = read_project_file("cmd/xnix-runtime-go/main.go") +
                                        read_project_file("cmd/xnix-runtime-go/recipe_conflict_audit_commands.go")
%w[runRecipeConflictAuditPreview recipe-conflict-audit-preview NewRecipeConflictAuditPreview RecipeConflictAuditOptions].each do |token|
  assert(go_runtime_recipe_conflict_cli_source.include?(token), "Go Runtime recipe conflict audit CLI must include #{token}")
end

go_runtime_recipe_conflict_cli_test_source = read_project_file("cmd/xnix-runtime-go/recipe_conflict_audit_cli_test.go")
%w[TestRecipeConflictAuditPreviewCLI TestRecipeConflictAuditPreviewCLIRequiresRegistry xnix.runtime.recipe_conflict_audit.v1].each do |token|
  assert(go_runtime_recipe_conflict_cli_test_source.include?(token), "Go Runtime recipe conflict audit CLI tests must include #{token}")
end

go_runtime_signed_recipe_verifier_source = read_project_file("internal/runtime/recipe/signature.go")
%w[SignedMetadata SignedVerifier SigningPayload ed25519.Verify xnix.recipe.signed_metadata.v1 xnix.runtime.signed_recipe_verification.v1 signed-recipe-verifier-preview offline-signed-recipe-verifier-evidence fixture-verified missing-signature-metadata application-id-mismatch digest-mismatch unsupported-schema unsupported-algorithm invalid-public-key invalid-signature].each do |token|
  assert(go_runtime_signed_recipe_verifier_source.include?(token), "Go Runtime signed recipe verifier must include #{token}")
end
%w[ProductionKeyConfigured ProductionTrustReady PrivateKeyLoaded NetworkRequired PackageManagerInvoked RecipeWritten RegistryMigrated BackendLaunchEnabled HostRootModified RecipePathExposed PublicKeyPathExposed SignatureMaterialExposed].each do |token|
  assert(go_runtime_signed_recipe_verifier_source.include?(token), "Go Runtime signed recipe verifier must expose disabled gate #{token}")
end

go_runtime_signed_recipe_verifier_test_source = read_project_file("internal/runtime/recipe/signature_test.go")
%w[TestSignedVerifierAcceptsValidFixture TestSignedVerifierFailsClosed TestSignedVerifierRequiresSignedStatusAndMetadata].each do |token|
  assert(go_runtime_signed_recipe_verifier_test_source.include?(token), "Go Runtime signed recipe verifier tests must include #{token}")
end

go_runtime_signed_recipe_verifier_cli_source = read_project_file("cmd/xnix-runtime-go/signed_recipe_verifier_commands.go")
%w[signed-recipe-verifier-preview runSignedRecipeVerifierPreview NewSignedRecipeVerificationPreview public-key metadata].each do |token|
  assert(go_runtime_signed_recipe_verifier_cli_source.include?(token), "Go Runtime signed recipe verifier CLI must include #{token}")
end

go_runtime_signed_recipe_verifier_cli_test_source = read_project_file("cmd/xnix-runtime-go/signed_recipe_verifier_cli_test.go")
%w[TestSignedRecipeVerifierPreviewCLI TestSignedRecipeVerifierPreviewCLIReportsInvalidSignature TestSignedRecipeVerifierPreviewCLIRejectsMissingArguments xnix.runtime.signed_recipe_verification.v1 production_key_configured production_trust_ready private_key_loaded backend_launch_enabled host_root_modified signature_material_exposed].each do |token|
  assert(go_runtime_signed_recipe_verifier_cli_test_source.include?(token), "Go Runtime signed recipe verifier CLI tests must include #{token}")
end

go_runtime_restricted_smoke_packet_source = read_project_file("internal/runtime/image/restricted_smoke_packet.go")
%w[RestrictedProductSmokePacket RestrictedSmokeEvidence PrepareRestrictedProductSmokePacket xnix.runtime.restricted_product_smoke_packet.v1 restricted-product-smoke-packet-preview dry-run-product-image-smoke-readiness runtime-owner artifact-trust backend-lifecycle portal-safety kde-entrypoints].each do |token|
  assert(go_runtime_restricted_smoke_packet_source.include?(token), "Go Runtime restricted product smoke packet must include #{token}")
end
%w[ReadyForAuthorizedSmoke ProductionRuntimeReady HumanAuthorizationRequired ExecutionAuthorized DockerExecuted QEMUExecuted ProductSmokeExecuted SerialLogPersistenceRequired SerialLogPersisted LoopbackOnlyNetworking DockerSocketMounted HostNetworkEnabled BroadHostMountEnabled PrivilegedContainerRequired BackendLaunchEnabled HostRootModified ReleaseReady].each do |token|
  assert(go_runtime_restricted_smoke_packet_source.include?(token), "Go Runtime restricted product smoke packet must expose gate #{token}")
end

go_runtime_restricted_smoke_packet_test_source = read_project_file("internal/runtime/image/restricted_smoke_packet_test.go")
%w[TestPrepareRestrictedProductSmokePacketRealRepo TestPrepareRestrictedProductSmokePacketReportsMissingEvidence TestPrepareRestrictedProductSmokePacketRejectsManifestEscape].each do |token|
  assert(go_runtime_restricted_smoke_packet_test_source.include?(token), "Go Runtime restricted product smoke packet tests must include #{token}")
end

restricted_smoke_packet_cli_source = read_project_file("cmd/xnix-runtime-go/restricted_product_smoke_packet_commands.go")
%w[restricted-product-smoke-packet-preview runRestrictedProductSmokePacketPreview PrepareRestrictedProductSmokePacket repo-root manifest].each do |token|
  assert(restricted_smoke_packet_cli_source.include?(token), "Restricted product smoke packet CLI must include #{token}")
end

restricted_smoke_packet_script_source = read_project_file("scripts/restricted_product_smoke_packet.rb")
%w[restricted-product-smoke-packet-preview JSON.pretty_generate render_markdown GOCACHE human_authorization_required docker_executed qemu_executed serial_log_persisted loopback_only_networking docker_socket_mounted host_network_enabled broad_host_mount_enabled host_root_modified].each do |token|
  assert(restricted_smoke_packet_script_source.include?(token), "Restricted product smoke packet report must include #{token}")
end
assert(!restricted_smoke_packet_script_source.include?("scripts/container.rb"), "Restricted product smoke packet must not run Docker")
assert(!restricted_smoke_packet_script_source.include?("boot-system"), "Restricted product smoke packet must not boot QEMU")

restricted_smoke_packet_test_source = read_project_file("test/test_restricted_product_smoke_packet.rb")
%w[xnix.runtime.restricted_product_smoke_packet.v1 ready_for_authorized_smoke human_authorization_required docker_executed qemu_executed release_ready].each do |token|
  assert(restricted_smoke_packet_test_source.include?(token), "Restricted product smoke packet tests must include #{token}")
end

go_runtime_kde_restricted_product_smoke_checkpoint_source = read_project_file("internal/runtime/appidentity/kde_restricted_product_smoke_checkpoint.go")
%w[KDERestrictedProductSmokeCheckpointRecord NewKDERestrictedProductSmokeCheckpointRecord xnix.runtime.kde_restricted_product_smoke_checkpoint.v1 kde-restricted-product-smoke-checkpoint-record preflight-boundary product-image-manifest repository-evidence authorization-boundary execution-unchanged session-unchanged smoke-not-executed host-boundary CheckpointReady ReadyForTrainGate ProductImageMetadataReady ReadyForAuthorizedSmoke].each do |token|
  assert(go_runtime_kde_restricted_product_smoke_checkpoint_source.include?(token), "Go Runtime KDE restricted product smoke checkpoint must include #{token}")
end
%w[HumanAuthorizationRequired ExecutionAuthorized DockerExecuted QEMUExecuted ProductSmokeExecuted SerialLogPersisted ReleaseReady LaunchPreflightPassed LaunchAuthorized CommandMaterialized ExecutablePathResolved BackendSelectedForLaunch BackendLaunchEnabled BackendProcessStarted ProductionBusOwnership DockerSocketMounted HostNetworkEnabled BroadHostMountEnabled PrivilegedContainerRequired HostRootModified StateRootPathExposed RepositoryRootPathExposed ManifestSourcePathExposed RawCommandExposed BackendDetailsExposed].each do |token|
  assert(go_runtime_kde_restricted_product_smoke_checkpoint_source.include?(token), "Go Runtime KDE restricted product smoke checkpoint must expose gate #{token}")
end

go_runtime_kde_restricted_product_smoke_checkpoint_test_source = read_project_file("internal/runtime/appidentity/kde_restricted_product_smoke_checkpoint_test.go")
%w[TestKDERestrictedProductSmokeCheckpointJoinsMetadataWithoutExecution recipe-trust runtime-write-gate assertKDERestrictedProductSmokeCheckpointDisabled].each do |token|
  assert(go_runtime_kde_restricted_product_smoke_checkpoint_test_source.include?(token), "Go Runtime KDE restricted product smoke checkpoint tests must include #{token}")
end

go_runtime_kde_restricted_product_smoke_checkpoint_cli_source = read_project_file("cmd/xnix-runtime-go/kde_restricted_product_smoke_checkpoint_commands.go") +
                                                          read_project_file("cmd/xnix-runtime-go/kde_restricted_product_smoke_checkpoint_cli_test.go") +
                                                          read_project_file("cmd/xnix-runtime-go/main.go")
%w[kde-restricted-product-smoke-checkpoint-record runKDERestrictedProductSmokeCheckpointRecord authorize-restricted-test-preparation test-only repo-root manifest TestKDERestrictedProductSmokeCheckpointRecordCommandIsReviewOnly TestKDERestrictedProductSmokeCheckpointRecordCommandRequiresAuthorization checkpoint_ready ready_for_train_gate docker_executed qemu_executed release_ready].each do |token|
  assert(go_runtime_kde_restricted_product_smoke_checkpoint_cli_source.include?(token), "Go Runtime KDE restricted product smoke checkpoint CLI must include #{token}")
end

go_runtime_snapshot_restore_source = read_project_file("internal/runtime/appidentity/snapshot_restore_candidates.go")
%w[SnapshotRestoreCandidatesPreview SnapshotRestoreCandidate SnapshotRestoreCounts SnapshotCandidateInput snapshot-restore-candidates-preview xnix.runtime.snapshot_restore_candidates.v1 GetSnapshotRestoreCandidates GetSnapshotRestoreCandidatesPreview latest-good latest-tested last-known-running blocked unknown active-session digest-mismatch].each do |token|
  assert(go_runtime_snapshot_restore_source.include?(token), "Go Runtime snapshot restore candidates preview must include #{token}")
end
%w[RestoreExecuted SnapshotDeletionEnabled FileContentRead SessionTerminated BackendLaunchEnabled StateRootPathExposed HostRootModified validateNoBackendTerms safety.ValidatePayload CompatibilityRisk].each do |token|
  assert(go_runtime_snapshot_restore_source.include?(token), "Go Runtime snapshot restore candidates preview must expose #{token}")
end

go_runtime_snapshot_restore_test_source = read_project_file("internal/runtime/appidentity/snapshot_restore_candidates_test.go")
%w[TestSnapshotRestoreCandidatesRankAndClassify TestSnapshotRestoreCandidatesActiveSessionBlocksAll TestSnapshotRestoreCandidatesNoCandidates TestSnapshotRestoreCandidatesRejectsUnsafeInputs].each do |token|
  assert(go_runtime_snapshot_restore_test_source.include?(token), "Go Runtime snapshot restore candidates tests must include #{token}")
end

go_runtime_snapshot_restore_cli_source = read_project_file("cmd/xnix-runtime-go/main.go") +
                                         read_project_file("cmd/xnix-runtime-go/snapshot_restore_candidates_commands.go")
%w[runSnapshotRestoreCandidatesPreview snapshot-restore-candidates-preview OpenReadOnly NewSnapshotRestoreCandidatesPreview active-session].each do |token|
  assert(go_runtime_snapshot_restore_cli_source.include?(token), "Go Runtime snapshot restore candidates CLI must include #{token}")
end

go_runtime_snapshot_restore_cli_test_source = read_project_file("cmd/xnix-runtime-go/snapshot_restore_candidates_cli_test.go")
%w[TestSnapshotRestoreCandidatesPreviewCLI TestSnapshotRestoreCandidatesPreviewCLIEmptyStateRoot TestSnapshotRestoreCandidatesPreviewCLIRequiresFlags xnix.runtime.snapshot_restore_candidates.v1].each do |token|
  assert(go_runtime_snapshot_restore_cli_test_source.include?(token), "Go Runtime snapshot restore candidates CLI tests must include #{token}")
end
assert(read_project_file("internal/runtime/snapshot/store.go").include?("OpenReadOnly"), "Go Runtime snapshot store must implement OpenReadOnly")

go_runtime_portal_read_source = read_project_file("internal/runtime/portal/read.go")
%w[ReadRequests portal-requests malformed].each do |token|
  assert(go_runtime_portal_read_source.include?(token), "Go Runtime portal read-only lister must include #{token}")
end

go_runtime_diagnostic_record_source = read_project_file("internal/runtime/diagnostics/record.go")
%w[RunRecordStore RunRecordRequest RunRecord xnix.runtime.diagnostic_run_record.v1 diagnostic-run-record go-runtime-state-root-diagnostic-run-record diagnostics-ledger runs result diagnostic_input repair_recommendation].each do |token|
  assert(go_runtime_diagnostic_record_source.include?(token), "Go Runtime diagnostic run records must include #{token}")
end
%w[RuntimeOwned GoRuntimeBacked KDEPolicyOwner StateRootPathExposed FixturePathExposed BackendStarted AIProviderCalled RealAIProviderEnabled AutoRepairAllowed RepairExecuted HostRootModified NetworkRequired PrivilegedContainerRequired BackendDetailsExposed FileContentsIncluded].each do |token|
  assert(go_runtime_diagnostic_record_source.include?(token), "Go Runtime diagnostic run records must expose #{token}")
end
%w[NewRunRecordStore OpenRunRecordStoreReadOnly safeDiagnosticRecordRoot safeDiagnosticReadOnlyRoot validateDiagnosticRecordID BuildDiagnosticInput Recommend NewResultStore].each do |token|
  assert(go_runtime_diagnostic_record_source.include?(token), "Go Runtime diagnostic run records must implement #{token}")
end
%w[unsupported\ schema identity\ or\ path\ mismatch digest\ mismatch unsafe\ enabled\ gates marshalDiagnosticRunRecord].each do |token|
  assert(go_runtime_diagnostic_record_source.include?(token), "Go Runtime diagnostic run record readback must validate #{token}")
end
go_runtime_diagnostic_record_test_source = read_project_file("internal/runtime/diagnostics/diagnostics_test.go")
%w[TestRunRecordStoreLoadRejectsTamperedReceipt tamper-check digest\ mismatch].each do |token|
  assert(go_runtime_diagnostic_record_test_source.include?(token), "Go Runtime diagnostic run record integrity tests must include #{token}")
end
assert(go_runtime_diagnostic_record_source.include?("refusing to use filesystem root"), "Go Runtime diagnostic run records must reject filesystem root state roots")

go_runtime_diagnostic_history_source = read_project_file("internal/runtime/diagnostics/history.go")
%w[RunHistory RunHistoryRecord RunHistoryCounts xnix.runtime.diagnostic_run_history.v1 diagnostic-run-history go-runtime-state-root-diagnostic-run-history latest failing_ids repair_issue snapshot_required].each do |token|
  assert(go_runtime_diagnostic_history_source.include?(token), "Go Runtime diagnostic run history must include #{token}")
end
%w[RuntimeOwned GoRuntimeBacked KDEPolicyOwner StateRootPathExposed BackendStarted AIProviderCalled RealAIProviderEnabled AutoRepairAllowed RepairExecuted HostRootModified NetworkRequired PrivilegedContainerRequired BackendDetailsExposed FileContentsIncluded].each do |token|
  assert(go_runtime_diagnostic_history_source.include?(token), "Go Runtime diagnostic run history must expose #{token}")
end
%w[History historyRecord runHistorySummary].each do |token|
  assert(go_runtime_diagnostic_history_source.include?(token), "Go Runtime diagnostic run history must implement #{token}")
end
assert(
  go_runtime_diagnostic_history_source.include?("applicationIDPattern") ||
    go_runtime_diagnostic_history_source.include?("appid.Valid"),
  "Go Runtime diagnostic run history must validate application ids"
)

go_runtime_diagnostic_history_preview_source = read_project_file("internal/runtime/appidentity/diagnostic_history.go")
%w[DiagnosticHistoryPreview DiagnosticHistoryCenterSummary xnix.runtime.diagnostic_history_preview.v1 diagnostic-history-preview go-runtime-state-root-diagnostic-run-history+kde-read-model KDE\ Plasma GetDiagnostics GetDiagnosticHistoryPreview compatibility_center not-run needs-review blocked healthy].each do |token|
  assert(go_runtime_diagnostic_history_preview_source.include?(token), "Go Runtime diagnostic history preview must include #{token}")
end
%w[RuntimeOwned GoRuntimeBacked KDEPolicyOwner UserVisible CompatibilityCenterCard SafeForAIDiagnostics StateRootPathExposed FileContentRead FilePathsExposed AIProviderCallEnabled RequestObjectCreated PermissionGranted LaunchEnabled ExecutionStarted RepairExecutionEnabled SettingsPersisted HostRootModified NetworkRequired PrivilegedContainerRequired BackendDetailsExposed].each do |token|
  assert(go_runtime_diagnostic_history_preview_source.include?(token), "Go Runtime diagnostic history preview must expose #{token}")
end
%w[NewDiagnosticHistoryPreview diagnosticHistoryState diagnosticHistorySummary validateNoBackendTerms].each do |token|
  assert(go_runtime_diagnostic_history_preview_source.include?(token), "Go Runtime diagnostic history preview must implement #{token}")
end

go_runtime_diagnostic_record_cli_source = [
  read_project_file("cmd/xnix-runtime-go/diagnostic_record_commands.go"),
  read_project_file("cmd/xnix-runtime-go/diagnostic_record_cli_test.go"),
  read_project_file("cmd/xnix-runtime-go/main.go")
].join("\n")
%w[diagnostic-run-record diagnostic-run-history diagnostic-history-preview state-root fixture run-id NewRunRecordStore RunRecordRequest NewDiagnosticHistoryPreview xnix.runtime.diagnostic_run_record.v1 xnix.runtime.diagnostic_run_history.v1 xnix.runtime.diagnostic_history_preview.v1 compatibility_center state_root_path_exposed fixture_path_exposed ai_provider_called real_ai_provider_enabled repair_executed].each do |token|
  assert(go_runtime_diagnostic_record_cli_source.include?(token), "Go Runtime diagnostic run record CLI must include #{token}")
end

go_runtime_support_bundle_source = read_project_file("internal/runtime/appidentity/support_bundle_manifest.go")
%w[SupportBundleManifestPreview SupportBundleManifestOptions SupportBundlePrivacyRedaction SupportBundleOmittedEvidenceCounts support-bundle-manifest-preview redacted-offline-support-bundle-manifest xnix.runtime.support_bundle_manifest.v1 GetSupportBundleManifest GetSupportBundleManifestPreview diagnostic-history-preview ai-diagnostic-input-preview ai-diagnostic-recommendation-preview runtime-version application-identity diagnostic-summaries failing-signals repair-recommendations ai-diagnostic-boundary omitted-evidence].each do |token|
  assert(go_runtime_support_bundle_source.include?(token), "Go Runtime support bundle manifest must include #{token}")
end
%w[RuntimeOwned GoRuntimeBacked KDEPolicyOwner UserVisible OfflineOnly ArchiveCreated FileContentRead FilePathsExposed AIProviderCalled AIProviderCallEnabled AutoRepairRequested AutoRepairExecuted BackendProcessStarted HostRootModified StateRootPathExposed RawExecutableExposed RawCommandExposed BackendDetailsExposed NetworkRequired PrivilegedContainerRequired].each do |token|
  assert(go_runtime_support_bundle_source.include?(token), "Go Runtime support bundle manifest must expose #{token}")
end
%w[SupportBundleManifestPreview NewDiagnosticHistoryPreview AIDiagnosticInputPreview AIDiagnosticRecommendationPreview validateNoBackendTerms supportBundleOmittedEvidence supportBundleFailingSignalIDs supportBundleRecommendationCategories].each do |token|
  assert(go_runtime_support_bundle_source.include?(token), "Go Runtime support bundle manifest must implement #{token}")
end

go_runtime_support_bundle_test_source = read_project_file("internal/runtime/appidentity/support_bundle_manifest_test.go")
%w[TestSupportBundleManifestPreviewSummarizesRedactedEvidence TestSupportBundleManifestPreviewHandlesEmptyHistory TestSupportBundleManifestPreviewRejectsMismatchedHistory support-bundle-manifest-preview redacted-offline-support-bundle-manifest archive_created file_content_read ai_provider_called host_root_modified state_root_path_exposed raw_command_exposed].each do |token|
  assert(go_runtime_support_bundle_test_source.include?(token), "Go Runtime support bundle manifest tests must include #{token}")
end

go_runtime_support_bundle_cli_source = [
  read_project_file("cmd/xnix-runtime-go/support_bundle_manifest_commands.go"),
  read_project_file("cmd/xnix-runtime-go/support_bundle_manifest_cli_test.go"),
  read_project_file("cmd/xnix-runtime-go/main.go")
].join("\n")
%w[support-bundle-manifest-preview runSupportBundleManifestPreview parseSupportBundleManifestPreviewSource OpenRunRecordStoreReadOnly runtime-root state-root xnix.runtime.support_bundle_manifest.v1 GetSupportBundleManifest GetSupportBundleManifestPreview redacted-offline-support-bundle-manifest archive_created file_content_read ai_provider_called host_root_modified].each do |token|
  assert(go_runtime_support_bundle_cli_source.include?(token), "Go Runtime support bundle manifest CLI must include #{token}")
end

go_runtime_support_case_timeline_source = read_project_file("internal/runtime/appidentity/support_case_timeline.go")
%w[SupportCaseTimelinePreview SupportCaseTimelineEvent SupportCaseTimelineCounts SupportCaseTimelineRedaction support-case-timeline-preview redacted-runtime-support-case-timeline xnix.runtime.support_case_timeline.v1 GetSupportCaseTimeline GetSupportCaseTimelinePreview diagnostic-runs blocked-actions repair-recommendations onboarding-gaps kde-entrypoint-state malformed-history].each do |token|
  assert(go_runtime_support_case_timeline_source.include?(token), "Go Runtime support case timeline preview must include #{token}")
end
%w[RuntimeOwned GoRuntimeBacked KDEPolicyOwner UserVisible ReviewOnly TicketCreated BundleExported AIProviderCalled AIProviderCallEnabled RepairApplied RepairExecuted ActionExecuted BackendProcessStarted FileContentRead FilePathsExposed StateRootPathExposed RawCommandExposed RawExecutableExposed BackendDetailsExposed HostRootModified NetworkRequired PrivilegedContainerRequired].each do |token|
  assert(go_runtime_support_case_timeline_source.include?(token), "Go Runtime support case timeline preview must expose #{token}")
end
%w[AIDiagnosticRecommendationPreview CompatibilityOnboardingChecklistPreview KDEActionDependencyGraphPreview KDEJourneyEvidencePreview validateNoBackendTerms safety.ValidatePayload supportCaseTimelineRedaction supportCaseTimelineEvents].each do |token|
  assert(go_runtime_support_case_timeline_source.include?(token), "Go Runtime support case timeline preview must implement #{token}")
end

go_runtime_support_case_timeline_test_source = read_project_file("internal/runtime/appidentity/support_case_timeline_test.go")
%w[TestSupportCaseTimelinePreviewNoHistory TestSupportCaseTimelinePreviewJoinsDiagnosticActionsRepairsOnboardingAndKDE TestSupportCaseTimelinePreviewMalformedMixedAndRedactedEvidence TestSupportCaseTimelineRejectsMismatchedHistoryApplication TicketCreated BundleExported AIProviderCalled RepairExecuted ActionExecuted BackendProcessStarted HostRootModified OmittedSensitiveEvidenceCount].each do |token|
  assert(go_runtime_support_case_timeline_test_source.include?(token), "Go Runtime support case timeline tests must include #{token}")
end

go_runtime_support_case_timeline_cli_source = [
  read_project_file("cmd/xnix-runtime-go/support_case_timeline_commands.go"),
  read_project_file("cmd/xnix-runtime-go/support_case_timeline_cli_test.go"),
  read_project_file("cmd/xnix-runtime-go/main.go")
].join("\n")
%w[support-case-timeline-preview runSupportCaseTimelinePreview parseSupportCaseTimelinePreviewSource LenientHistory OpenRunRecordStoreReadOnly runtime-root state-root xnix.runtime.support_case_timeline.v1 GetSupportCaseTimeline GetSupportCaseTimelinePreview ticket_created bundle_exported ai_provider_called repair_executed action_executed backend_process_started host_root_modified state_root_path_exposed].each do |token|
  assert(go_runtime_support_case_timeline_cli_source.include?(token), "Go Runtime support case timeline CLI must include #{token}")
end

go_runtime_support_case_timeline_cli_test_source = read_project_file("cmd/xnix-runtime-go/support_case_timeline_cli_test.go")
%w[TestSupportCaseTimelinePreviewCLI TestSupportCaseTimelinePreviewCLIMissingStateRootDoesNotCreate TestSupportCaseTimelinePreviewCLIMalformedRecordIsBlocked TestSupportCaseTimelinePreviewCLIRequiresRecipeSource xnix.runtime.support_case_timeline.v1].each do |token|
  assert(go_runtime_support_case_timeline_cli_test_source.include?(token), "Go Runtime support case timeline CLI tests must include #{token}")
end

go_runtime_multi_application_install_queue_source = read_project_file("internal/runtime/appidentity/multi_application_install_queue.go")
%w[MultiApplicationInstallQueuePreview MultiApplicationInstallItem MultiApplicationInstallCounts multi-application-install-queue-preview review-only-multi-application-install-queue xnix.runtime.multi_application_install_queue.v1 GetMultiApplicationInstallQueue GetMultiApplicationInstallQueuePreview registry+compatibility-install-preview].each do |token|
  assert(go_runtime_multi_application_install_queue_source.include?(token), "Go Runtime multi-application install queue must include #{token}")
end
%w[RuntimeOwned GoRuntimeBacked KDEPolicyOwner UserVisible ReviewOnly QueuePersisted RequestObjectsCreated ArtifactsStaged ArtifactsDownloaded NetworkRequestCreated HostPackageManagerInvoked DesktopActivationStarted InstallStarted BackendProcessStarted LaunchEnabled ExecutionStarted SettingsPersisted HostRootModified PrivilegedContainerRequired StateRootPathExposed RawExecutableExposed RawCommandExposed BackendDetailsExposed].each do |token|
  assert(go_runtime_multi_application_install_queue_source.include?(token), "Go Runtime multi-application install queue must expose #{token}")
end
%w[CompatibilityInstallPlanPreview validateNoBackendTerms multiApplicationInstallItemState multiApplicationInstallQueueStatus uniqueSortedStrings].each do |token|
  assert(go_runtime_multi_application_install_queue_source.include?(token), "Go Runtime multi-application install queue must implement #{token}")
end

go_runtime_multi_application_install_queue_test_source = read_project_file("internal/runtime/appidentity/multi_application_install_queue_test.go")
%w[TestMultiApplicationInstallQueuePreviewAggregatesInstallPlans TestMultiApplicationInstallQueuePreviewRejectsEmptyQueue TestMultiApplicationInstallQueuePreviewRejectsBadMode multi-application-install-queue-preview review-only-multi-application-install-queue request_objects_created artifacts_staged network_request_created host_package_manager_invoked backend_process_started host_root_modified raw_command_exposed].each do |token|
  assert(go_runtime_multi_application_install_queue_test_source.include?(token), "Go Runtime multi-application install queue tests must include #{token}")
end

go_runtime_multi_application_install_queue_cli_source = [
  read_project_file("cmd/xnix-runtime-go/multi_application_install_queue_commands.go"),
  read_project_file("cmd/xnix-runtime-go/multi_application_install_queue_cli_test.go"),
  read_project_file("cmd/xnix-runtime-go/main.go")
].join("\n")
%w[multi-application-install-queue-preview runMultiApplicationInstallQueuePreview parseMultiApplicationInstallQueueSource repeatedAppIDs LoadRecipeFromRegistry ParseRegistry xnix.runtime.multi_application_install_queue.v1 GetMultiApplicationInstallQueue GetMultiApplicationInstallQueuePreview review-only-multi-application-install-queue request_objects_created artifacts_staged host_root_modified backend_details_exposed].each do |token|
  assert(go_runtime_multi_application_install_queue_cli_source.include?(token), "Go Runtime multi-application install queue CLI must include #{token}")
end

go_runtime_offline_application_fixture_matrix_source = read_project_file("internal/runtime/appidentity/offline_application_fixture_matrix.go")
%w[OfflineApplicationFixtureMatrixPreview OfflineApplicationFixtureMatrixRow OfflineApplicationFixtureMatrixCounts OfflineApplicationFixtureBackendAdapterContract offline-application-fixture-matrix-preview offline-cross-application-fixture-matrix xnix.runtime.offline_application_fixture_matrix.v1 GetOfflineApplicationFixtureMatrix GetOfflineApplicationFixtureMatrixPreview built-in-fixtures+runtime-read-models+backend-adapter-contract-preview document-editor game installer launcher network-heavy tray-heavy unsupported backend-adapter-noop-contract noop-contract-ready].each do |token|
  assert(go_runtime_offline_application_fixture_matrix_source.include?(token), "Go Runtime offline application fixture matrix must include #{token}")
end
%w[RuntimeOwned GoRuntimeBacked KDEPolicyOwner BackendAdapterContractRead BackendAdapterProfileCount BackendAdapterNoopContracts BackendAdapterAuditReady UserVisible ReviewOnly OfflineDefault NetworkFetchEnabled PackageManagerInvoked ArtifactStagingEnabled BackendLaunchEnabled DockerRequired QEMURequired RequestObjectsCreated SettingsPersisted FileContentRead StateRootPathExposed RawExecutableExposed RawCommandExposed BackendDetailsExposed HostRootModified PrivilegedContainerRequired AdapterInvocationEnabled InstallEnabled DownloadEnabled CommandMaterialized ExecutablePathResolved ProfilePathExposed].each do |token|
  assert(go_runtime_offline_application_fixture_matrix_source.include?(token), "Go Runtime offline application fixture matrix must expose #{token}")
end
%w[CompatibilityInstallPlanPreview BackendSelectionPreview NewBackendAdapterContractPreview NewSnapshotPlanPreview AIDiagnosticInputPreview KDEJourneyEvidencePreviewWithOptions validateNoBackendTerms offlineApplicationFixtureBackendAdapterContract offlineApplicationFixtureRowState countOfflineApplicationFixtureMatrixRows].each do |token|
  assert(go_runtime_offline_application_fixture_matrix_source.include?(token), "Go Runtime offline application fixture matrix must implement #{token}")
end

go_runtime_offline_application_fixture_matrix_test_source = read_project_file("internal/runtime/appidentity/offline_application_fixture_matrix_test.go")
%w[TestOfflineApplicationFixtureMatrixPreviewCoversRepresentativeShapes TestOfflineApplicationFixtureMatrixPreviewAuditsBackendAdapterContractMappings TestOfflineApplicationFixtureMatrixPreviewReportsMissingFixtures TestOfflineApplicationFixtureMatrixPreviewRejectsUnknownShape TestOfflineApplicationFixtureMatrixPreviewUnsupportedShapeIsBlocked TestOfflineApplicationFixtureMatrixPreviewNeverEnablesSideEffects TestOfflineApplicationFixtureMatrixPreviewIsDesktopSafe offline-application-fixture-matrix-preview blocked-unsupported noop-contract-ready backend-adapter-noop-contract network_fetch_enabled package_manager_invoked artifact_staging_enabled backend_process_started adapter_invocation_enabled command_materialized host_root_modified raw_command_exposed].each do |token|
  assert(go_runtime_offline_application_fixture_matrix_test_source.include?(token), "Go Runtime offline application fixture matrix tests must include #{token}")
end

go_runtime_offline_application_fixture_matrix_cli_source = [
  read_project_file("cmd/xnix-runtime-go/offline_application_fixture_matrix_commands.go"),
  read_project_file("cmd/xnix-runtime-go/offline_application_fixture_matrix_cli_test.go"),
  read_project_file("cmd/xnix-runtime-go/main.go")
].join("\n")
%w[offline-application-fixture-matrix-preview runOfflineApplicationFixtureMatrixPreview parseOfflineApplicationFixtureMatrixOptions repeatedFixtureShapeIDs xnix.runtime.offline_application_fixture_matrix.v1 GetOfflineApplicationFixtureMatrix GetOfflineApplicationFixtureMatrixPreview offline-cross-application-fixture-matrix backend_adapter_contract_read backend_adapter_audit_ready noop-contract-ready adapter_invocation_enabled command_materialized network_fetch_enabled package_manager_invoked artifact_staging_enabled backend_launch_enabled docker_required qemu_required host_root_modified].each do |token|
  assert(go_runtime_offline_application_fixture_matrix_cli_source.include?(token), "Go Runtime offline application fixture matrix CLI must include #{token}")
end

offline_application_fixture_matrix_script_source = read_project_file("scripts/offline_application_fixture_matrix.rb")
%w[offline-application-fixture-matrix-preview --format --shape JSON.pretty_generate render_markdown GOCACHE .cache document-editor game installer launcher network-heavy tray-heavy unsupported backend_adapter_contract_read backend_adapter_audit_ready noop-contract-ready adapter_invocation_enabled command_materialized network_fetch_enabled package_manager_invoked artifact_staging_enabled backend_launch_enabled docker_required qemu_required host_root_modified].each do |token|
  assert(offline_application_fixture_matrix_script_source.include?(token), "Offline application fixture matrix script must include #{token}")
end
assert(!offline_application_fixture_matrix_script_source.include?("scripts/container.rb"), "Offline application fixture matrix must not run Docker")
assert(!offline_application_fixture_matrix_script_source.include?("boot-system"), "Offline application fixture matrix must not boot QEMU")

offline_application_fixture_matrix_script_test_source = read_project_file("test/test_offline_application_fixture_matrix.rb")
%w[offline-application-fixture-matrix-preview JSON.parse render_markdown document-editor unsupported blocked-unsupported noop-contract-ready backend_adapter_contract_read backend_adapter_audit_ready network_fetch_enabled package_manager_invoked artifact_staging_enabled backend_launch_enabled docker_required qemu_required host_root_modified].each do |token|
  assert(offline_application_fixture_matrix_script_test_source.include?(token), "Offline application fixture matrix script tests must include #{token}")
end

go_runtime_application_upgrade_impact_source = read_project_file("internal/runtime/appidentity/application_upgrade_impact.go")
%w[ApplicationUpgradeImpactPreview ApplicationUpgradeImpactSection ApplicationUpgradeImpactCounts application-upgrade-impact-preview review-only-application-upgrade-impact xnix.runtime.application_upgrade_impact.v1 GetApplicationUpgradeImpact GetApplicationUpgradeImpactPreview registry+compatibility-install-preview+desktop-activation-evidence recipe-metadata package-source artifact-digests backend-profile portal-permissions snapshot-requirements desktop-activation diagnostics candidate-newer same-version candidate-older changed-without-version].each do |token|
  assert(go_runtime_application_upgrade_impact_source.include?(token), "Go Runtime application upgrade impact must include #{token}")
end
%w[RuntimeOwned GoRuntimeBacked KDEPolicyOwner UserVisible ReviewOnly RecipeWritten RegistryMigrated RequestObjectsCreated ArtifactsStaged ArtifactsDownloaded NetworkRequestCreated HostPackageManagerInvoked SettingsPersisted DesktopActivationStarted DesktopFilesWritten BackendProcessStarted LaunchEnabled ExecutionStarted HostRootModified PrivilegedContainerRequired StateRootPathExposed RawExecutableExposed RawCommandExposed BackendDetailsExposed].each do |token|
  assert(go_runtime_application_upgrade_impact_source.include?(token), "Go Runtime application upgrade impact must expose #{token}")
end
%w[CompatibilityInstallPlanPreview validateNoBackendTerms recipeVersionRelation compareSemanticVersions applicationUpgradeImpactOverallState applicationUpgradeArtifactState].each do |token|
  assert(go_runtime_application_upgrade_impact_source.include?(token), "Go Runtime application upgrade impact must implement #{token}")
end

go_runtime_application_upgrade_impact_test_source = read_project_file("internal/runtime/appidentity/application_upgrade_impact_test.go")
%w[TestApplicationUpgradeImpactPreviewClassifiesNewerCandidate TestApplicationUpgradeImpactPreviewClassifiesSameAndOlderCandidates TestApplicationUpgradeImpactPreviewBlocksUntrustedProductionCandidate TestApplicationUpgradeImpactPreviewRejectsMalformedCandidate TestApplicationUpgradeImpactPreviewBlocksCandidateIdentityDrift application-upgrade-impact-preview review-only-application-upgrade-impact recipe_written registry_migrated artifacts_staged network_request_created host_package_manager_invoked backend_process_started host_root_modified raw_command_exposed].each do |token|
  assert(go_runtime_application_upgrade_impact_test_source.include?(token), "Go Runtime application upgrade impact tests must include #{token}")
end

go_runtime_application_upgrade_impact_cli_source = [
  read_project_file("cmd/xnix-runtime-go/application_upgrade_impact_commands.go"),
  read_project_file("cmd/xnix-runtime-go/application_upgrade_impact_cli_test.go"),
  read_project_file("cmd/xnix-runtime-go/main.go")
].join("\n")
%w[application-upgrade-impact-preview runApplicationUpgradeImpactPreview parseApplicationUpgradeImpactPreviewSource candidate-registry candidate-app LoadRecipeFromRegistry xnix.runtime.application_upgrade_impact.v1 GetApplicationUpgradeImpact GetApplicationUpgradeImpactPreview review-only-application-upgrade-impact recipe_written registry_migrated artifacts_staged host_root_modified backend_details_exposed recipe digest mismatch].each do |token|
  assert(go_runtime_application_upgrade_impact_cli_source.include?(token), "Go Runtime application upgrade impact CLI must include #{token}")
end

go_runtime_kde_shell_source = read_project_file("internal/runtime/appidentity/kde_shell_surface.go")
%w[KDEIntegrationStatusPreview KDEShellIntegrationPlanPreview KDEApplicationSurfacePlanPreview kde-integration-status-preview kde-shell-integration-preview kde-application-surface-preview xnix.runtime.kde_integration_status.v1 xnix.runtime.kde_shell_integration.v1 xnix.runtime.kde_application_surface.v1 GetKDEIntegrationStatus GetKDEIntegrationStatusPreview GetKDEShellIntegrationPlan GetKDEShellIntegrationPlanPreview GetKDEApplicationSurfacePlan GetKDEApplicationSurfacePlanPreview go-runtime-kde-integration-status go-runtime-kde-shell-integration go-runtime-kde-application-surface launcher task-manager file-manager system-tray notifications compatibility-center settings krunner-search kwin-window-management].each do |token|
  assert(go_runtime_kde_shell_source.include?(token), "Go Runtime KDE shell previews must include #{token}")
end
%w[RuntimeOwned GoRuntimeBacked KDEPolicyOwner OfficialDesktopOnly StableDesktopContract FallbackDesktopsSupported PlasmaForkRequired PlasmaSourceModified ShellConfigurationWritten ComponentActivationEnabled BackendLaunchEnabled HostRootModified PrivilegedContainerRequired BackendDetailsExposed StandardLauncherVisible TaskManagerIdentityReady FileAssociationsPlanned DolphinActionPlanned KRunnerQueryPlanned TrayStatusPlanned NotificationRoutePlanned SettingsSurfacePlanned PortalReviewRequired ExecutionReady LaunchEnabled BackendProcessStarted DesktopFilesWritten MIMEAppsWritten BackendCommandExposed RawWindowsExecutableExposed].each do |token|
  assert(go_runtime_kde_shell_source.include?(token), "Go Runtime KDE shell previews must expose safety flag #{token}")
end
assert(go_runtime_kde_shell_source.include?("validateNoBackendTerms"), "Go Runtime KDE shell previews must hide backend terms")

go_runtime_kde_shell_cli_source = read_project_file("cmd/xnix-runtime-go/kde_shell_commands.go")
%w[runKDEIntegrationStatusPreview runKDEShellIntegrationPreview runKDEApplicationSurfacePreview kde-integration-status-preview kde-shell-integration-preview kde-application-surface-preview NewKDEIntegrationStatusPreview NewKDEShellIntegrationPlanPreview KDEApplicationSurfacePlanPreview].each do |token|
  assert(go_runtime_kde_shell_cli_source.include?(token), "Go Runtime KDE shell CLI must include #{token}")
end

snapshot_store_source = read_project_file("internal/runtime/snapshot/store.go")
%w[Package snapshot Store Manifest RollbackReceipt New Create List Verify Restore .xnix-snapshots host_root_touched filepath.Rel sha256].each do |token|
  assert(snapshot_store_source.include?(token), "Snapshot store must include #{token}")
end
%w[validSnapshotID within manifestContentHash writeObject load objectPath manifestPath].each do |token|
  assert(snapshot_store_source.include?(token), "Snapshot store must implement #{token}")
end
assert(snapshot_store_source.include?("filepath.SkipDir"), "Snapshot store must skip its private metadata directory")
assert(snapshot_store_source.include?("d.Type().IsRegular()"), "Snapshot store must snapshot only regular files")
assert(snapshot_store_source.include?("HostRootTouched: false"), "Snapshot rollback receipts must keep host-root mutation false")

snapshot_store_test_source = read_project_file("internal/runtime/snapshot/store_test.go")
%w[TestCreateVerifyList TestRestoreRollsBackFixtureState TestVerifyDetectsCorruption TestStoreRefusesPathsOutsideRoot TestMissingStateRootFails].each do |token|
  assert(snapshot_store_test_source.include?(token), "Snapshot store tests must include #{token}")
end

portal_broker_source = [
  read_project_file("internal/runtime/portal/request.go"),
  read_project_file("internal/runtime/portal/broker.go"),
  read_project_file("internal/runtime/portal/ledger.go")
].join("\n")
%w[Package portal Broker FakeBroker NewFakeBroker CreateRequest Resolve Complete Cancel Get List RequestSpec Request RequestState PermissionState PortalDestination PortalObjectPath SupportedOperations pending-user-mediation granted denied cancelled failed completed org.freedesktop.portal.Desktop org.freedesktop.portal.FileChooser org.freedesktop.portal.OpenURI org.freedesktop.portal.Print org.freedesktop.portal.Screenshot org.freedesktop.portal.Clipboard org.freedesktop.portal.Camera org.freedesktop.portal.RemoteDesktop file-open uri-open print screenshot clipboard camera remote-desktop].each do |token|
  assert(portal_broker_source.include?(token), "Portal broker must include #{token}")
end
%w[Ledger NewLedger Record xnix.runtime.portal_request_record.v1 portal-permission-request-record go-runtime-state-root-portal-broker RequestRelativePath RealPortalCallEnabled RequestObjectCreated PermissionGranted ExecutionApproved StateRootPathExposed HostRootModified BackendDetailsExposed PrivilegedContainerRequired].each do |token|
  assert(portal_broker_source.include?(token), "Portal request ledger must include #{token}")
end
%w[UserMediationRequired RequestObjectRequired DirectAccessAllowed Recoverable HostPermissionChanged BackendDetailsExposed PermissionPending PermissionGranted PermissionDenied PermissionNotGranted OutcomeGranted OutcomeDenied OutcomeCancelled OutcomeFailed].each do |token|
  assert(portal_broker_source.include?(token), "Portal broker must expose #{token}")
end
assert(
  portal_broker_source.include?("applicationIDPattern") ||
    portal_broker_source.include?("appid.Valid"),
  "Portal broker must validate application ids"
)
assert(portal_broker_source.include?("handleToken"), "Portal broker must create deterministic request handles")

portal_broker_test_source = [
  read_project_file("internal/runtime/portal/request_test.go"),
  read_project_file("internal/runtime/portal/broker_test.go"),
  read_project_file("internal/runtime/portal/ledger_test.go"),
  read_project_file("cmd/xnix-runtime-go/runtime_safety_cli_test.go")
].join("\n")
%w[TestNewRequestAskOperationStartsPending TestNewRequestDenyOperationIsTerminal TestNewRequestRejectsBadInput TestSupportedOperationsSorted TestFakeBrokerGrantThenCompleteFlow TestFakeBrokerFailedIsRecoverableThenRetried TestFakeBrokerCompleteRequiresGrant TestFakeBrokerCancelPending TestFakeBrokerListAndGetTrackInCreationOrder TestLedgerPersistsGrantAndCompletionWithoutRealPortal TestPortalRequestRecordCommandPersistsPermissionFlow].each do |token|
  assert(portal_broker_test_source.include?(token), "Portal broker tests must include #{token}")
end

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
assert(runtime_daemon_source.include?("\"desktop_icon_planning\""), "Runtime daemon must expose desktop icon planning capability")
assert(runtime_daemon_source.include?("\"task_manager_identity_planning\""), "Runtime daemon must expose task manager identity planning capability")
assert(runtime_daemon_source.include?("\"kwin_window_rule_planning\""), "Runtime daemon must expose KWin window rule planning capability")
assert(runtime_daemon_source.include?("\"file_association_planning\""), "Runtime daemon must expose file association planning capability")
assert(runtime_daemon_source.include?("\"notification_planning\""), "Runtime daemon must expose notification planning capability")
assert(runtime_daemon_source.include?("\"tray_status_planning\""), "Runtime daemon must expose tray status planning capability")
assert(runtime_daemon_source.include?("\"krunner_query_planning\""), "Runtime daemon must expose KRunner query planning capability")
assert(runtime_daemon_source.include?("\"compatibility_execution_readiness\""), "Runtime daemon must expose execution readiness capability")
assert(runtime_daemon_source.include?("\"compatibility_launch_intents\""), "Runtime daemon must expose launch intent capability")
dbus_smoke_source = [
  read_project_file("runtime/dbus/xnix_compatd_smoke.c"),
  read_project_file("runtime/dbus/xnix_compatd_introspection.inc"),
  read_project_file("runtime/dbus/xnix_compatd_kde_center.inc"),
  read_project_file("runtime/dbus/xnix_compatd_runtime_models.inc")
].join("\n")
%w[GetEngineCatalog GetRunPlan GetDesktopActivationManifest GetKDEIntegrationStatus GetKDEShellIntegrationPlan GetKDEApplicationSurfacePlan GetDesktopResourceBridgePlan GetCompatibilityModeSwitchPlan GetCompatibilityPermissionReviewPlan GetCompatibilityReviewFlowPlan GetDesktopEntryPlan GetDesktopIconPlan GetTaskManagerIdentityPlan GetKWinWindowRulePlan GetFileAssociationPlan GetNotificationPlan GetTrayStatus GetKRunnerQueryPlan GetPortalRequestPlan GetApplicationStateRoot GetCompatibilityPackageSource GetCompatibilityAcquisitionPreflight GetCompatibilityActionQueue GetCompatibilityActionReviewReceipt GetCompatibilityCenterSummary GetKDECenterPage GetKDECenterPageSections GetKDECenterPageSectionDetail GetCompatibilityArtifactManifest GetCompatibilityInstallPlan GetBackendBinding GetBackendCapabilityMatrix GetBackendSelectionPlan GetBackendLifecycle GetBackendEnvironmentPlan GetRepairPlan GetTestPlan GetTestResult GetExecutionReadiness GetLaunchIntent GetAIDiagnosticInput GetAIDiagnosticRecommendation GetAIRepairApprovalGate GetSnapshotPlan GetPortalAccessPolicy GetRuntimeServiceBinding GetRuntimeLiveOwnerGate GetRuntimeOwnerProcess GetRuntimeOwnerSmokePlan GetRuntimeMethodParityManifest GetRuntimeOwnerRouteManifest GetRuntimeOwnerRecipeTrust GetRuntimeOwnerReadiness GetRuntimeWriteGate GetCompatibilitySettings GetCompatibilitySettingsChangePlan].each do |method_name|
  assert(runtime_daemon_source.include?("\"#{method_name}\""), "Runtime daemon dispatch must include #{method_name}")
  assert(read_project_file("runtime/dbus/org.xnix.Compatibility1.xml").include?("name=\"#{method_name}\""), "D-Bus contract must include #{method_name}")
  assert(dbus_smoke_source.include?(method_name), "D-Bus smoke adapter must include #{method_name}")
  assert(read_project_file("scripts/dbus_session_smoke.rb").include?(method_name), "D-Bus session smoke must call #{method_name}")
end
%w[g_spawn_sync go_owner_read_dispatch go_owner_service_call5 add_go_owner_dispatch_bridge_fields go_owner_write_gate_dispatch go_owner_dispatch_available go_owner_dispatch_schema go_owner_dispatch_json go_owner_service_call_available go_owner_service_call_schema go_owner_service_call_request_type go_owner_service_call_json go-runtime-owner-dispatch+c-smoke-bridge xnix.runtime.owner_read_dispatch.v1 runtime-owner-read-dispatch xnix.runtime.owner_service_call.v1 runtime-owner-service-call variant_string_field assert_go_owner_service_call_envelope go-runtime-owner-in-process-service dispatch_ready write_methods_enabled].each do |token|
  assert(dbus_smoke_source.include?(token) || read_project_file("scripts/dbus_session_smoke.rb").include?(token), "D-Bus smoke adapter must expose Go owner bridge token #{token}")
end

dbus_client_source = read_project_file("lib/xnix/compatibility/dbus_runtime_client.rb")
%w[engine_catalog run_plan kde_integration_status kde_shell_integration_plan kde_application_surface_plan desktop_resource_bridge_plan compatibility_mode_switch_plan compatibility_permission_review_plan compatibility_review_flow_plan desktop_entry_plan desktop_icon_plan task_manager_identity_plan kwin_window_rule_plan file_association_plan notification_plan tray_status krunner_query_plan state_root package_source acquisition_preflight action_queue action_review_receipt compatibility_center_summary artifact_manifest install_plan backend_binding backend_capability_matrix backend_selection_plan repair_plan test_plan test_result execution_readiness ai_diagnostic_input ai_diagnostic_recommendation ai_repair_approval_gate snapshot_plan portal_access_policy runtime_service_binding runtime_live_owner_gate runtime_owner_process runtime_owner_smoke_plan runtime_method_parity_manifest runtime_owner_route_manifest runtime_owner_recipe_trust runtime_owner_readiness runtime_write_gate settings settings_change_plan].each do |method_name|
  assert(dbus_client_source.include?("def #{method_name}"), "D-Bus Runtime client must expose #{method_name}")
end
assert(dbus_client_source.include?("return true if value == \"true\""), "D-Bus Runtime client must parse boolean true values")
assert(dbus_client_source.include?("return false if value == \"false\""), "D-Bus Runtime client must parse boolean false values")
assert(dbus_client_source.include?("value.to_i"), "D-Bus Runtime client must parse integer values")
assert(dbus_client_source.include?("value.start_with?(\"[\")"), "D-Bus Runtime client must parse string arrays")

runtime_owner_candidate_smoke_source = read_project_file("scripts/runtime_owner_candidate_smoke.rb")
%w[dbus-run-session xnix-runtime-owner smoke-owner dispatch-read service-call lifecycle-log smoke-batch session-bus-smoke xnix.runtime.owner_candidate.v1 xnix.runtime.owner_read_dispatch.v1 xnix.runtime.owner_service_call.v1 xnix.runtime.owner_lifecycle_event.v1 xnix.runtime.owner_smoke_batch.v1 xnix.runtime.owner_session_bus_smoke.v1 runtime-owner-lifecycle-event runtime-owner-service-call runtime-owner-smoke-batch-record runtime-owner-session-bus-smoke-step restricted-private-session-bus-owner-smoke org.xnix.Compatibility1.Error.WriteMethodDisabled org.xnix.Compatibility1.Error.UnsupportedMethod session_bus_claimed production_bus_claimed system_service_started network_required host_root_modified privileged_container_required backend_details_exposed preview-complete].each do |token|
  assert(runtime_owner_candidate_smoke_source.include?(token), "Runtime owner candidate smoke must include #{token}")
end

runtime_contract_drift_report_source = read_project_file("scripts/runtime_contract_drift_report.rb")
%w[
  runtime-contract-drift-report
  xnix.runtime.contract_drift_report.v1
  parity-read-methods
  owner-route-go-methods
  dbus-client-method-map
  owner-read-dispatch-local-methods
  owner-read-dispatch-all-read-methods
  owner-read-dispatch-contract-subset
  owner-smoke-batch-source
  owner-session-bus-smoke-source
  go-owner-smoke-bridge
  runtime-dispatch
  dbus-client-definitions
  smoke-adapter
  session-smoke
  go-cli-route-commands
  GetBackendAdapterProfileAudit
  GetRestrictedOwnerSmokeReceiptLookupPreview
  GetRestrictedOwnerSmokeReceiptFanOut
  GetKDETestLaunchMaterializationReceiptLookupPreview
  GetKDETestLaunchMaterializationFanOut
  json
  markdown
  drift_detected
  write_methods_supported
  write_method_dispatch_enabled
  network_required
  host_root_modified
  privileged_container_required
  backend_details_exposed
].each do |token|
  assert(runtime_contract_drift_report_source.include?(token), "Runtime contract drift report must include #{token}")
end

runtime_contract_drift_report_test_source = read_project_file("test/test_runtime_contract_drift_report.rb")
%w[--format json markdown parity-read-methods owner-read-dispatch-local-methods owner-read-dispatch-all-read-methods owner-session-bus-smoke-source GetBackendAdapterProfileAudit GetRestrictedOwnerSmokeReceiptLookupPreview GetKDETestLaunchMaterializationReceiptLookupPreview GetKDETestLaunchMaterializationFanOut owner_read_dispatch_method_count owner_session_bus_smoke_step_count drift_detected host_root_modified backend_details_exposed].each do |token|
  assert(runtime_contract_drift_report_test_source.include?(token), "Runtime contract drift report test must include #{token}")
end

implementation_evidence_report_source = read_project_file("scripts/implementation_evidence_report.rb")
%w[
  implementation-evidence-report
  xnix.runtime.implementation_evidence_report.v1
  runtime-owner-service
  xnix.runtime.owner_lifecycle_event.v1
  full D-Bus read dispatch coverage
  lifecycle JSONL
  recipe-artifact-trust-pipeline
  environment-lifecycle-state
  portal-snapshot-control-plane
  kde-activation-shell-materialization
  execution-transaction-ledger
  diagnostics-repair-ai-boundary
  atomic-kde-image-qemu-acceptance
  developer-verification-harness
  docs/claude-code-mainline-implementation-plan.md
  mainline-implementation-plan
  mainline_package
  mainline_document
  mainline_plan_present
  mainline_package_count
  mainline_first_wave
  next_dispatch_packages
  next_dispatch_summary
  codex/runtime-owner-read-service
  codex/recipe-artifact-trust-pipeline
  codex/environment-lifecycle-state
  codex/implementation-evidence-harness
  M1
  M5
  M9
  kde_first_presence_smoke.rb
  desktop_safety_policy.go
  desktop-safety-policy-preview
  xnix.runtime.desktop_safety_policy.v1
  xnix.kde_first_presence_smoke.v1
  kde-first-presence-smoke
  route_baseline
  entrypoint_count
  settings_field_ids
  forbidden_user_terms
  safety_false_keys
  contract-only
  fixture-implemented
  state-root-implemented
  smoke-owned
  production-gated
  production_gate_evidence
  orphan_read_methods
  orphan_preview_methods_detected
  write_methods_supported
  write_method_dispatch_enabled
  network_required
  host_root_modified
  privileged_container_required
  backend_launch_enabled
  production_ready
  json
  markdown
].each do |token|
  assert(implementation_evidence_report_source.include?(token), "Implementation evidence report must include #{token}")
end

implementation_evidence_report_test_source = read_project_file("test/test_implementation_evidence_report.rb")
%w[
  --format
  json
  markdown
  runtime-owner-service
  lifecycle.go
  full D-Bus read dispatch coverage
  mainline_document
  mainline_package
  mainline_first_wave
  next_dispatch_packages
  suggested_branch
  codex/runtime-owner-read-service
  M1
  M5
  M9
  kde-activation-shell-materialization
  kde_first_presence_smoke.rb
  xnix.kde_first_presence_smoke.v1
  all seven KDE entry points
  smoke-owned
  fixture-implemented
  state-root-implemented
  production_gate_evidence
  orphan_preview_methods_detected
  production_ready
  host_root_modified
  backend_launch_enabled
].each do |token|
  assert(implementation_evidence_report_test_source.include?(token), "Implementation evidence report test must include #{token}")
end

mainline_integration_review_source = read_project_file("scripts/mainline_integration_review.rb")
%w[
  mainline-integration-review
  xnix.runtime.mainline_integration_review.v1
  docs/mainline-integration-checkpoint.md
  docs/claude-code-implementation-packages.md
  planning-documents
  version-and-product-metadata
  cw1-runtime-owner-read-boundary
  runtime_route_convergence
  shared-runtime-cli-plumbing
  cw2-recipe-artifact-trust
  cw3-runtime-state-backend-lifecycle
  cw5-portal-permission-safety
  cw8-execution-ledger-session-evidence
  cw4-kde-entrypoint-consumers
  cw6-snapshot-rollback-store
  cw7-ai-diagnostic-privacy
  cw10-evidence-drift-harness
  blocked-protected-claude-owned-file
  unclassified
  .gocache/
  tmp/
  --status-fixture
  markdown
  safe_to_stage_all
  review_matrix
  required_verification
  required_evidence
  safety_guards
  never-stage-all
  docker_or_qemu_required
  host_root_modified
  backend_launch_enabled
].each do |token|
  assert(mainline_integration_review_source.include?(token), "Mainline integration review script must include #{token}")
end

mainline_integration_review_test_source = read_project_file("test/test_mainline_integration_review.rb")
%w[
  --status-fixture
  mainline-integration-review
  protected_claude_file_modified
  excluded_file_count
  safe_to_stage_all
  review_matrix
  required_verification
  required_evidence
  Safety guards
  planning-documents
  cw1-runtime-owner-read-boundary
  shared-runtime-cli-plumbing
  cw2-recipe-artifact-trust
  cw3-runtime-state-backend-lifecycle
  cw5-portal-permission-safety
  cw8-execution-ledger-session-evidence
  cw4-kde-entrypoint-consumers
  cw6-snapshot-rollback-store
  cw7-ai-diagnostic-privacy
  cw10-evidence-drift-harness
  blocked-protected-claude-owned-file
  unclassified
  markdown
].each do |token|
  assert(mainline_integration_review_test_source.include?(token), "Mainline integration review test must include #{token}")
end

release_evidence_index_source = read_project_file("scripts/release_evidence_index.rb")
%w[
  release-evidence-index
  xnix.runtime.release_evidence_index.v1
  implementation-evidence+contract-drift+mainline-review+kde-first-presence
  REPORT_COMMANDS
  CLAIM_DEFINITIONS
  runtime-owner-read-boundary
  recipe-artifact-trust
  runtime-state-backend-lifecycle
  portal-snapshot-safety
  kde-seven-entrypoints
  execution-session-evidence
  diagnostics-repair-ai-boundary
  evidence-drift-harness
  protected-claude-file
  unclassified-files
  product-image-qemu-acceptance
  restricted-heavy-smoke-skipped
  implemented
  fixture-only
  contract-only
  blocked
  skipped
  human-authorized
  malformed-report
  docker_executed
  qemu_executed
  network_checks_run
  package_manager_invoked
  backend_launch_enabled
  host_root_modified
  automatic_staging_enabled
  automatic_release_tagging_enabled
].each do |token|
  assert(release_evidence_index_source.include?(token), "Release evidence index script must include #{token}")
end

release_evidence_index_test_source = read_project_file("test/test_release_evidence_index.rb")
%w[
  release-evidence-index
  xnix.runtime.release_evidence_index.v1
  implemented
  fixture-only
  contract-only
  blocked
  skipped
  malformed_report_detected
  protected-file
  unclassified-file
  human-authorized
  product-image-qemu-acceptance
  restricted-heavy-smoke-skipped
  docker_executed
  qemu_executed
  network_checks_run
  package_manager_invoked
  backend_launch_enabled
  host_root_modified
  automatic_staging_enabled
  automatic_release_tagging_enabled
].each do |token|
  assert(release_evidence_index_test_source.include?(token), "Release evidence index test must include #{token}")
end

merge_readiness_packet_source = read_project_file("scripts/merge_readiness_packet.rb")
%w[
  merge-readiness-packet
  xnix.runtime.merge_readiness_packet.v1
  TOOL_DEFINITIONS
  layout
  implementation
  contract_drift
  kde_smoke
  mainline_review
  release_evidence
  offline_fixture_matrix
  --offline-only
  --skip-tool
  --tool-command
  tool_statuses
  lane_classification
  protected_file_status
  unsafe_operation_status
  changed_file_counts
  required_follow_up_commands
  merge_blocking_reasons
  release_blocking_reasons
  restricted-docker-or-qemu-smoke-requires-human-authorization
  staging
  committing
  tagging
  pushing
].each do |token|
  assert(merge_readiness_packet_source.include?(token), "Merge readiness packet script must include #{token}")
end
%w[
  docker_executed
  qemu_executed
  network_checks_run
  package_manager_invoked
  backend_launch_enabled
  host_root_modified
  automatic_staging_enabled
  automatic_commit_enabled
  automatic_release_tagging_enabled
  automatic_push_enabled
].each do |token|
  assert(merge_readiness_packet_source.include?("\"#{token}\" => false"), "Merge readiness packet script must keep #{token} false")
end
assert(!merge_readiness_packet_source.include?("scripts/container.rb"), "Merge readiness packet must not run Docker")
assert(!merge_readiness_packet_source.include?("boot-system"), "Merge readiness packet must not boot QEMU")

merge_readiness_packet_test_source = read_project_file("test/test_merge_readiness_packet.rb")
%w[
  merge-readiness-packet
  xnix.runtime.merge_readiness_packet.v1
  malformed-json
  missing-command
  implementation:skipped
  protected-claude-file-modified
  unclassified-files-present
  unsafe-operation-detected
  docker_executed
  qemu_executed
  network_checks_run
  package_manager_invoked
  host_root_modified
  markdown
].each do |token|
  assert(merge_readiness_packet_test_source.include?(token), "Merge readiness packet test must include #{token}")
end

claude_code_dispatch_runbook_source = read_project_file("docs/claude-code-dispatch-runbook.md")
[
  "Claude Code Dispatch Runbook",
  "docs/claude-code-mainline-task-batch.md",
  "docs/claude-code-third-wave-task-batch.md",
  "docs/claude-code-fourth-wave-task-batch.md",
  "docs/claude-code-fifth-wave-task-batch.md",
  "docs/claude-code-sixth-wave-task-batch.md",
  "docs/claude-code-seventh-wave-task-batch.md",
  "docs/mainline-integration-checkpoint.md",
  "docs/claude-code-implementation-packages.md",
  "CB1",
  "CB2",
  "CB3",
  "CB4",
  "CB5",
  "CB6",
  "CB7",
  "CB8",
  "CB9",
  "protected_claude_file_modified",
  "unclassified_file_count",
  "mainline_integration_review.rb",
  "runtime_contract_drift_report.rb",
  "implementation_evidence_report.rb",
  "git diff --check",
  "Never run",
  "git add ."
].each do |token|
  assert(claude_code_dispatch_runbook_source.include?(token), "Claude Code dispatch runbook must include #{token}")
end

claude_code_second_wave_source = read_project_file("docs/claude-code-second-wave-task-batch.md")
[
  "Claude Code Second-Wave Task Batch",
  "SW1",
  "SW2",
  "SW3",
  "SW4",
  "SW5",
  "SW6",
  "SW7",
  "SW8",
  "Do not modify `docs/claude-code-implementation-packages.md`",
  "KDE",
  "Go",
  "Runtime",
  "Docker",
  "QEMU",
  "host-root mutation",
  "Copyable Claude Code prompt"
].each do |token|
  assert(claude_code_second_wave_source.include?(token), "Claude Code second-wave task batch must include #{token}")
end

claude_code_third_wave_source = read_project_file("docs/claude-code-third-wave-task-batch.md")
[
  "Claude Code Third-Wave Task Batch",
  "TW1",
  "TW2",
  "TW3",
  "TW4",
  "TW5",
  "TW6",
  "TW7",
  "TW8",
  "Runtime route convergence plan",
  "Action queue receipt dependency graph",
  "Portal-to-execution preflight join",
  "Snapshot baseline readiness join",
  "AI diagnostics privacy corpus",
  "KDE action-card evidence deck",
  "Runtime owner unsupported-route hardening",
  "Restricted acceptance dashboard",
  "Do not modify `docs/claude-code-implementation-packages.md`",
  "Docker or QEMU execution unless a human explicitly authorizes a restricted smoke",
  "Copyable Claude Code prompt"
].each do |token|
  assert(claude_code_third_wave_source.include?(token), "Claude Code third-wave task batch must include #{token}")
end

claude_code_fourth_wave_source = read_project_file("docs/claude-code-fourth-wave-task-batch.md")
[
  "Claude Code Fourth-Wave Task Batch",
  "FW1",
  "FW2",
  "FW3",
  "FW4",
  "FW5",
  "FW6",
  "FW7",
  "FW8",
  "Runtime evidence audit timeline",
  "Portal permission lifecycle dashboard",
  "Snapshot restore rehearsal plan",
  "Backend capability fixture probe harness",
  "Settings change dependency review",
  "Diagnostics repair playbook review queue",
  "Desktop activation materialization audit",
  "Restricted release readiness packet",
  "Do not modify `docs/claude-code-implementation-packages.md`",
  "Docker or QEMU execution unless a human explicitly authorizes a restricted smoke",
  "Copyable Claude Code prompt"
].each do |token|
  assert(claude_code_fourth_wave_source.include?(token), "Claude Code fourth-wave task batch must include #{token}")
end

claude_code_fifth_wave_source = read_project_file("docs/claude-code-fifth-wave-task-batch.md")
[
  "Claude Code Fifth-Wave Task Batch",
  "F5W1",
  "F5W2",
  "F5W3",
  "F5W4",
  "F5W5",
  "F5W6",
  "F5W7",
  "F5W8",
  "KDE journey evidence stitcher",
  "Compatibility onboarding checklist",
  "Offline support bundle manifest",
  "Multi-application install queue preview",
  "Runtime state maintenance planner",
  "Recipe update migration preview",
  "KDE notification digest model",
  "Restricted acceptance fixture suite",
  "Do not modify `docs/claude-code-implementation-packages.md`",
  "Docker or QEMU execution unless a human explicitly authorizes a restricted smoke",
  "Copyable Claude Code prompt"
].each do |token|
  assert(claude_code_fifth_wave_source.include?(token), "Claude Code fifth-wave task batch must include #{token}")
end

claude_code_sixth_wave_source = read_project_file("docs/claude-code-sixth-wave-task-batch.md")
[
  "Claude Code Sixth-Wave Task Batch",
  "S6W1",
  "S6W2",
  "S6W3",
  "S6W4",
  "S6W5",
  "S6W6",
  "S6W7",
  "S6W8",
  "Support case timeline preview",
  "Desktop deactivation dry-run plan",
  "Recipe conflict and pin audit",
  "Portal permission renewal preview",
  "Snapshot restore candidate ranking",
  "Compatibility settings profile migration preview",
  "Offline application fixture matrix",
  "Merge readiness packet aggregator",
  "Do not modify `docs/claude-code-implementation-packages.md`",
  "Docker or QEMU execution unless a human explicitly authorizes a restricted smoke",
  "Copyable Claude Code prompt"
].each do |token|
  assert(claude_code_sixth_wave_source.include?(token), "Claude Code sixth-wave task batch must include #{token}")
end

claude_code_seventh_wave_source = read_project_file("docs/claude-code-seventh-wave-task-batch.md")
[
  "Claude Code Seventh-Wave Task Batch",
  "S7W1",
  "S7W2",
  "S7W3",
  "S7W4",
  "S7W5",
  "S7W6",
  "S7W7",
  "S7W8",
  "Application upgrade impact preview",
  "Runtime policy explanation cards",
  "State-root quota and retention preview",
  "Crash and hang signal summary preview",
  "Compatibility backend fallback preview",
  "KDE search visibility plan",
  "Permission evidence audit preview",
  "Release evidence index",
  "Do not modify `docs/claude-code-implementation-packages.md`",
  "Docker or QEMU execution unless a human explicitly authorizes a restricted smoke",
  "Copyable Claude Code prompt"
].each do |token|
  assert(claude_code_seventh_wave_source.include?(token), "Claude Code seventh-wave task batch must include #{token}")
end

execution_ledger_source = read_project_file("internal/runtime/execution/ledger.go")
%w[
  xnix.runtime.execution_ledger.v1
  execution-transaction-ledger-record
  go-runtime-state-root-execution-ledger
  execution-ledger
  transactions
  state_root_path_exposed
  LaunchAllowed
  LaunchEnabled
  BackendStarted
  PermissionGranted
  PortalPermissionReceiptCount
  portal_permission_receipt_relative_paths
  portal_permission_receipt_consumed
  HostRootModified
  NetworkRequired
  PrivilegedContainerRequired
  BackendDetailsExposed
  refusing\ to\ use\ filesystem\ root
  path\ traversal
  execution\ ledger\ record\ digest\ mismatch
  execution\ ledger\ record\ has\ unsafe\ enabled\ gates
].each do |token|
  assert(execution_ledger_source.include?(token.gsub("\\ ", " ")), "Execution ledger must include #{token}")
end

execution_session_source = read_project_file("internal/runtime/execution/session.go")
%w[
  xnix.runtime.execution_session_record.v1
  execution-session-status-record
  go-runtime-state-root-execution-session
  execution-ledger
  sessions
  transaction_relative_path
  StatusPersisted
  SessionActive
  LiveStateObserved
  WindowObserved
  TaskManagerEntryActive
  KWinRuleApplied
  LiveTrayBridgeEnabled
  LaunchEnabled
  ExecutionStarted
  BackendProcessStarted
  HostRootModified
  BackendDetailsExposed
  LoadSession
  execution\ session\ record\ digest\ mismatch
  execution\ session\ record\ has\ unsafe\ enabled\ gates
].each do |token|
  assert(execution_session_source.include?(token.gsub("\\ ", " ")), "Execution session record must include #{token}")
end

restricted_authorization_source = read_project_file("internal/runtime/execution/authorization.go")
%w[
  xnix.runtime.restricted_launch_authorization.v1
  restricted-launch-authorization-receipt
  go-runtime-state-root-restricted-launch-authorization
  restricted-test-preparation
  authorize-restricted-test-preparation
  authorized-preparation-only
  RestrictedAuthorizationStore
  RestrictedAuthorizationRequest
  PreparationAuthorized
  LaunchAuthorized
  ProcessStartAuthorized
  ProductionTrustSatisfied
  RuntimeWriteGateEnabled
  ArtifactAcquisitionEnabled
  BackendInstallEnabled
  BackendLaunchEnabled
  BackendProcessStarted
  StateRootPathExposed
  HostRootModified
  digest\ mismatch
  invalid\ or\ unsafe\ gates
  managed\ path\ must\ be\ a\ real\ directory
].each do |token|
  assert(restricted_authorization_source.include?(token.gsub("\\ ", " ")), "Restricted launch authorization store must include #{token}")
end
restricted_authorization_test_source = read_project_file("internal/runtime/execution/authorization_test.go")
%w[TestRestrictedAuthorizationStorePersistsPreparationOnlyReceipt TestRestrictedAuthorizationStoreRejectsInvalidDirectiveTamperingAndSymlink launch_authorized managed\ path\ symlink].each do |token|
  assert(restricted_authorization_test_source.include?(token.gsub("\\ ", " ")), "Restricted launch authorization tests must include #{token}")
end

execution_receipt_integrity_test_source = read_project_file("internal/runtime/execution/ledger_test.go") +
                                          read_project_file("internal/runtime/execution/session_test.go")
%w[TestLedgerLoadRejectsTamperedReceipt TestLoadSessionRejectsTamperedReceipt tampered\ execution\ receipt tampered\ session\ receipt].each do |token|
  assert(execution_receipt_integrity_test_source.include?(token.gsub("\\ ", " ")), "Execution receipt integrity tests must include #{token}")
end

execution_ledger_cli_source = read_project_file("cmd/xnix-runtime-go/execution_ledger_commands.go")
%w[
  execution-ledger-record
  execution-session-record
  state-root
  request-id
  recipe-trust
  snapshot-baseline
  portal-required
  portal-granted
  portal-request
  executionLedgerPortalReceipts
  NewLedger
  Record
  RecordSession
].each do |token|
  assert(execution_ledger_cli_source.include?(token), "Execution ledger CLI must include #{token}")
end

execution_ledger_cli_test_source = read_project_file("cmd/xnix-runtime-go/execution_ledger_cli_test.go")
%w[
  execution-ledger-record
  state-root
  xnix.runtime.execution_ledger.v1
  state_root_path_exposed
  host_root_modified
  backend_details_exposed
  TestExecutionLedgerRecordCommandConsumesPortalReceipt
  TestExecutionLedgerRecordCommandBlocksDeniedPortalReceipt
  TestExecutionSessionRecordCommandPersistsStatusFromLedger
  execution-session-record
  portal_permission_receipt_count
].each do |token|
  assert(execution_ledger_cli_test_source.include?(token), "Execution ledger CLI test must include #{token}")
end

artifact_stage_source = read_project_file("internal/runtime/artifact/stage.go")
%w[
  xnix.runtime.artifact_stage_receipt.v1
  compatibility-artifact-stage-receipt
  go-runtime-local-fixture-artifact-staging
  artifact-ledger
  receipts
  cache_root_path_exposed
  fixture_root_path_exposed
  network_fetch_enabled
  package_manager_invoked
  host_root_modified
  backend_launch_enabled
  backend_details_exposed
  ValidateStageReceipt
  RequiredStaged
  artifact stage receipt digest mismatch
  artifact stage receipt contains unsafe side effects
].each do |token|
  assert(artifact_stage_source.include?(token), "Artifact stage receipt must include #{token}")
end

artifact_stage_cli_source = read_project_file("cmd/xnix-runtime-go/artifact_stage_commands.go")
%w[
  artifact-stage-record
  manifest
  cache-root
  fixture-root
  StageFromFixture
].each do |token|
  assert(artifact_stage_cli_source.include?(token), "Artifact stage CLI must include #{token}")
end

artifact_stage_cli_test_source = read_project_file("cmd/xnix-runtime-go/artifact_stage_cli_test.go")
%w[
  artifact-stage-record
  xnix.runtime.artifact_stage_receipt.v1
  cache_root_path_exposed
  fixture_root_path_exposed
  network_fetch_enabled
  package_manager_invoked
  host_root_modified
  backend_launch_enabled
].each do |token|
  assert(artifact_stage_cli_test_source.include?(token), "Artifact stage CLI test must include #{token}")
end

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

kde_image_source = read_project_file("lib/xnix/image/kde_image.rb")
assert(kde_image_source.include?("xnix-kde-image"), "KDE image model must expose a CLI command")
assert(kde_image_source.include?("kde-plasma-atomic-image"), "KDE image model must identify the atomic image artifact")
assert(kde_image_source.include?("render_containerfile"), "KDE image model must render a Containerfile from the manifest")
kde_image_manifest = read_project_file("image/kinoite/manifest.json")
assert(kde_image_manifest.include?("fedora-kinoite"), "KDE image manifest must build on a Fedora Kinoite base")
assert(kde_image_manifest.include?("kde-plasma-6"), "KDE image manifest must target KDE Plasma 6")
assert(kde_image_manifest.include?("\"runtime_owned\": true"), "KDE image manifest must keep Runtime policy ownership")
assert(kde_image_manifest.include?("\"kde_policy_owner\": false"), "KDE image manifest must not hand policy ownership to KDE")
kde_containerfile = read_project_file("image/kinoite/Containerfile")
assert(kde_containerfile.include?("GENERATED by xnix-kde-image"), "KDE Containerfile must be a generated snapshot")
assert(kde_containerfile.include?("FROM quay.io/fedora/fedora-kinoite"), "KDE Containerfile must pin the Kinoite base")

kde_disk_source = read_project_file("lib/xnix/image/disk_build.rb")
assert(kde_disk_source.include?("xnix-kde-disk"), "KDE disk build must expose a CLI command")
assert(kde_disk_source.include?("kde-plasma-disk-image"), "KDE disk build must identify the disk image artifact")
assert(kde_disk_source.include?("bootc-image-builder") || kde_disk_source.include?("builder_image"),
       "KDE disk build must drive bootc-image-builder")
kde_disk_config = read_project_file("image/kinoite/disk-config.json")
assert(kde_disk_config.include?("bootc-image-builder"), "KDE disk config must reference bootc-image-builder")
assert(kde_disk_config.include?("\"qcow2\""), "KDE disk config must produce a qcow2 image")
assert(kde_disk_config.include?("ttyS0"), "KDE disk config must enable the serial console for boot smoke")

puts "PASS: Xnix #{EXPECTED_VERSION} scaffold is consistent"
