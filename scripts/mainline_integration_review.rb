#!/usr/bin/env ruby
# frozen_string_literal: true

require "json"
require "open3"
require "optparse"
require "pathname"

PROJECT_ROOT = Pathname.new(__dir__).join("..").realpath
VERSION = PROJECT_ROOT.join("VERSION").read.strip
PROTECTED_CLAUDE_FILE = "docs/claude-code-implementation-packages.md"
EXCLUDED_PREFIXES = %w[
  .gocache/
  tmp/
].freeze

RISKY_PATH_RULES = [
  { reason: "build-artifact", pattern: /\.(o|a|so|dylib|qcow2|img|iso)\z/ },
  { reason: "local-log", pattern: /\.log\z/ },
  { reason: "private-config", pattern: %r{(\A|/)\.env(\.|\z)|\.(pem|key|p12)\z|(\A|/)id_(rsa|ed25519|ecdsa)(\.pub)?\z} }
].freeze

CODE_CHANGE_PREFIXES = %w[
  bin/
  cmd/
  internal/
  kde/
  lib/
  runtime/
  scripts/
  test/
].freeze

VERSION_SYNC_FILES = %w[
  CLAUDE.md
  PRODUCT_OVERVIEW.md
  README.md
  runtime/core/xnix_runtime_core.h
  kde/plasmoids/org.xnix.compatibilitycenter/metadata.json
].freeze

LANES = [
  {
    id: "planning-documents",
    workstream: "planning",
    title: "Planning and handoff documents",
    description: "Roadmaps, assignment boards, Xnix current mainline notes, and mainline convergence documents that do not by themselves enable product behavior.",
    review_order: 1,
    patterns: [
      %r{\Adocs/claude-code-},
      %r{\Adocs/windows-app-compatibility-implementation-brief\.md\z},
      %r{\Adocs/mainline-integration-checkpoint\.md\z},
      %r{\Adocs/xnix-current-mainline\.md\z}
    ]
  },
  {
    id: "version-and-product-metadata",
    workstream: "release",
    title: "Version and product metadata",
    description: "Version, changelog, product overview, README, C header, and KDE metadata that must stay aligned with code changes.",
    review_order: 2,
    patterns: [
      %r{\ACHANGELOG\.md\z},
      %r{\ACLAUDE\.md\z},
      %r{\APRODUCT_OVERVIEW\.md\z},
      %r{\AREADME\.md\z},
      %r{\AVERSION\z},
      %r{\Aruntime/core/xnix_runtime_core\.h\z},
      %r{\Akde/plasmoids/org\.xnix\.compatibilitycenter/metadata\.json\z}
    ]
  },
  {
    id: "cw1-runtime-owner-read-boundary",
    workstream: "CW1 / A1",
    title: "Runtime owner read boundary",
    description: "Go Runtime owner read dispatch, private smoke owner evidence, route parity, and disabled write behavior.",
    review_order: 3,
    patterns: [
      %r{\Acmd/xnix-runtime-owner/},
      %r{\Ainternal/runtime/owner/},
      %r{\Ainternal/runtime/appidentity/runtime_owner_},
      %r{\Ainternal/runtime/appidentity/runtime_live_owner_gate_},
      %r{\Ainternal/runtime/appidentity/runtime_method_parity_manifest_},
      %r{\Ainternal/runtime/appidentity/runtime_route_convergence},
      %r{\Ainternal/runtime/appidentity/runtime_service_activation_preflight},
      %r{\Ainternal/runtime/appidentity/runtime_service_binding_},
      %r{\Ainternal/runtime/appidentity/runtime_write_gate\.go\z},
      %r{\Ainternal/runtime/appidentity/runtime_write_gate_},
      %r{\Ascripts/runtime_contract_drift_report\.rb\z},
      %r{\Ascripts/runtime_owner_candidate_smoke\.rb\z},
      %r{\Ascripts/staged_launcher_dispatch_smoke\.rb\z},
      %r{\Ascripts/desktop_trigger_request_preflight_smoke\.rb\z},
      %r{\Ascripts/dbus_controlled_launch_owner_fixture_smoke\.rb\z},
      %r{\Aruntime/dbus/xnix_compatd_smoke\.c\z},
      %r{\Atest/test_runtime_contract_drift_report\.rb\z},
      %r{\Atest/test_runtime_owner_candidate_smoke_script\.rb\z},
      %r{\Atest/test_runtime_status_owner_service_session_bus_smoke_script\.rb\z},
      %r{\Atest/test_staged_launcher_dispatch_smoke_script\.rb\z},
      %r{\Atest/test_desktop_trigger_request_preflight_smoke_script\.rb\z},
      %r{\Atest/test_dbus_controlled_launch_owner_fixture_smoke_script\.rb\z},
      %r{\Atest/test_runtime_dbus_smoke_script\.rb\z},
      %r{\Atest/test_runtime_owner_},
      %r{\Acmd/xnix-runtime-go/production_dbus_gate_},
      %r{\Acmd/xnix-runtime-go/restricted_owner_smoke_receipt_},
      %r{\Acmd/xnix-runtime-go/current_mainline_},
      %r{\Acmd/xnix-runtime-go/runtime_owner_},
      %r{\Acmd/xnix-runtime-go/runtime_route_convergence_},
      %r{\Acmd/xnix-runtime-go/owner_service_launch_envelope_guard_},
      %r{\Acmd/xnix-runtime-go/desktop_trigger_dry_run_request_review_},
      %r{\Acmd/xnix-runtime-go/desktop_trigger_service_call_materialization_},
      %r{\Acmd/xnix-runtime-go/desktop_trigger_request_preflight_},
      %r{\Acmd/xnix-runtime-go/runtime_live_owner_gate_cli_test\.go\z},
      %r{\Acmd/xnix-runtime-go/runtime_method_parity_manifest_cli_test\.go\z},
      %r{\Acmd/xnix-runtime-go/runtime_service_binding_cli_test\.go\z},
      %r{\Acmd/xnix-runtime-go/runtime_write_gate_cli_test\.go\z}
    ]
  },
  {
    id: "shared-runtime-cli-plumbing",
    workstream: "shared",
    title: "Shared Runtime CLI and identity plumbing",
    description: "Cross-lane command registration, identity metadata, C core version assertions, and generic Runtime read-model tests that must be reviewed with the lane that introduced the behavior.",
    review_order: 3.5,
    patterns: [
      %r{\Acmd/xnix-runtime-go/main\.go\z},
      %r{\Acmd/xnix-runtime-go/main_test\.go\z},
      %r{\Acmd/xnix-runtime-go/version_test\.go\z},
      %r{\Ainternal/runtime/appidentity/identity\.go\z},
      %r{\Ainternal/runtime/appidentity/identity_test\.go\z},
      %r{\Ainternal/runtime/appidentity/version_test\.go\z},
      %r{\Ainternal/testversion/},
      %r{\Atest/test_runtime_core\.rb\z},
      %r{\Atest/test_runtime_daemon\.rb\z},
      %r{\Atest/test_runtime_live_owner_gate\.rb\z},
      %r{\Atest/test_runtime_method_parity_manifest\.rb\z},
      %r{\Atest/test_runtime_service_binding\.rb\z},
      %r{\Atest/test_runtime_write_gate\.rb\z},
      %r{\Atest/test_compatibility_engine_catalog\.rb\z},
      %r{\Atest/test_compatibility_run_plan\.rb\z},
      %r{\Atest/test_compatibility_test_plan\.rb\z},
      %r{\Atest/test_compatibility_test_result\.rb\z}
    ]
  },
  {
    id: "cw2-recipe-artifact-trust",
    workstream: "CW2 / A2",
    title: "Recipe and artifact trust pipeline",
    description: "Recipe trust, artifact staging receipts, install-readiness consumption, and package-source gates.",
    review_order: 4,
    patterns: [
      %r{\Ainternal/runtime/recipe/},
      %r{\Ainternal/runtime/artifact/},
      %r{\Aruntime/recipes/},
      %r{\Ainternal/runtime/appidentity/install_plan},
      %r{\Ainternal/runtime/appidentity/multi_application_install_queue},
      %r{\Ainternal/runtime/appidentity/application_upgrade_impact},
      %r{\Ainternal/runtime/appidentity/runtime_owner_recipe_trust},
      %r{\Ainternal/runtime/appidentity/recipe_conflict_audit},
      %r{\Acmd/xnix-runtime-go/recipe_conflict_audit_},
      %r{\Acmd/xnix-runtime-go/signed_recipe_verifier_},
      %r{\Acmd/xnix-runtime-go/install_plan_},
      %r{\Acmd/xnix-runtime-go/multi_application_install_queue_},
      %r{\Acmd/xnix-runtime-go/application_upgrade_impact_},
      %r{\Acmd/xnix-runtime-go/artifact_stage_},
      %r{\Atest/test_recipe_},
      %r{\Atest/test_compatibility_artifact_manifest\.rb\z},
      %r{\Atest/test_compatibility_acquisition_preflight\.rb\z},
      %r{\Atest/test_compatibility_install_plan\.rb\z},
      %r{\Atest/test_compatibility_package_source\.rb\z}
    ]
  },
  {
    id: "cw3-runtime-state-backend-lifecycle",
    workstream: "CW3 / A3",
    title: "Runtime state root and backend lifecycle",
    description: "Runtime-owned backend inventory, backend lifecycle records, environment state, and launch-disabled readiness.",
    review_order: 5,
    patterns: [
      %r{\Ainternal/runtime/environment/},
      %r{\Ainternal/runtime/appidentity/backend_},
      %r{\Ainternal/runtime/appidentity/compatibility_backend_fallback},
      %r{\Ainternal/runtime/appidentity/state_root_quota_retention},
      %r{\Acmd/xnix-runtime-go/backend_adapter_contract_},
      %r{\Acmd/xnix-runtime-go/state_root_quota_retention_},
      %r{\Acmd/xnix-runtime-go/compatibility_backend_fallback_},
      %r{\Acmd/xnix-runtime-go/backend_group_cli_test\.go\z},
      %r{\Atest/test_compatibility_backend_},
      %r{\Atest/test_application_state_root\.rb\z},
      %r{\Atest/test_compatibility_execution_readiness\.rb\z}
    ]
  },
  {
    id: "cw5-portal-permission-safety",
    workstream: "CW5 / A4",
    title: "Portal permission safety plane",
    description: "Fake-mode Portal request records, permission review states, and safety command coverage without real Portal calls.",
    review_order: 6,
    patterns: [
      %r{\Ainternal/runtime/portal/},
      %r{\Ainternal/runtime/appidentity/permission_evidence_audit},
      %r{\Ainternal/runtime/appidentity/portal_permission_renewal},
      %r{\Acmd/xnix-runtime-go/runtime_safety_},
      %r{\Acmd/xnix-runtime-go/runtime_safety_commands\.go\z},
      %r{\Acmd/xnix-runtime-go/permission_evidence_audit_},
      %r{\Acmd/xnix-runtime-go/portal_permission_renewal_},
      %r{\Atest/test_portal_},
      %r{\Atest/test_compatibility_permission_review_plan\.rb\z}
    ]
  },
  {
    id: "cw8-execution-ledger-session-evidence",
    workstream: "CW8 / A6",
    title: "Execution ledger and session evidence",
    description: "Reviewed execution transaction records and KDE-safe session evidence while real launch remains disabled.",
    review_order: 7,
    patterns: [
      %r{\Ainternal/runtime/execution/},
      %r{\Ainternal/runtime/appidentity/application_readiness},
      %r{\Ainternal/runtime/appidentity/kde_test_launch_materialization},
      %r{\Acmd/xnix-runtime-go/execution_ledger_},
      %r{\Acmd/xnix-runtime-go/kde_test_launch_materialization_},
      %r{\Acmd/xnix-runtime-go/application_readiness_},
      %r{\Ainternal/runtime/appidentity/execution_session_record_evidence\.go\z},
      %r{\Atest/test_launch_request\.rb\z},
      %r{\Atest/test_compatibility_execution_readiness\.rb\z}
    ]
  },
  {
    id: "cw4-kde-entrypoint-consumers",
    workstream: "CW4 / A5",
    title: "KDE seven entry-point consumers",
    description: "KDE shell, Compatibility Center, task manager, KWin, tray, Dolphin, settings, and desktop-facing safe read models.",
    review_order: 8,
    patterns: [
      %r{\Akde/},
      %r{\Ainternal/runtime/appidentity/compatibility_onboarding},
      %r{\Ainternal/runtime/appidentity/desktop_safety_policy},
      %r{\Ainternal/runtime/appidentity/desktop_activation_transaction},
      %r{\Ainternal/runtime/appidentity/desktop_deactivation_dry_run},
      %r{\Ainternal/runtime/appidentity/kde_},
      %r{\Ainternal/runtime/appidentity/managed_launcher_acceptance_report},
      %r{\Ainternal/runtime/appidentity/runtime_policy_explanation_cards},
      %r{\Ainternal/runtime/appidentity/settings_profile_migration},
      %r{\Ainternal/runtime/appidentity/window_identity_},
      %r{\Acmd/xnix-runtime-go/compatibility_onboarding_},
      %r{\Acmd/xnix-runtime-go/desktop_activation_cli_test\.go\z},
      %r{\Acmd/xnix-runtime-go/desktop_deactivation_dry_run_},
      %r{\Acmd/xnix-runtime-go/desktop_safety_policy_},
      %r{\Acmd/xnix-runtime-go/kde_},
      %r{\Acmd/xnix-runtime-go/managed_launcher_acceptance_report_},
      %r{\Acmd/xnix-runtime-go/runtime_policy_explanation_cards_},
      %r{\Acmd/xnix-runtime-go/settings_profile_migration_},
      %r{\Acmd/xnix-runtime-go/window_identity_},
      %r{\Ascripts/kde_first_presence_smoke\.rb\z},
      %r{\Atest/test_kde_},
      %r{\Atest/test_kwin_},
      %r{\Atest/test_krunner_},
      %r{\Atest/test_tray_},
      %r{\Atest/test_task_manager_},
      %r{\Atest/test_file_association_},
      %r{\Atest/test_settings_},
      %r{\Atest/test_action_review_receipt\.rb\z},
      %r{\Atest/test_compatibility_action_queue\.rb\z},
      %r{\Atest/test_compatibility_mode_switch_plan\.rb\z},
      %r{\Atest/test_compatibility_review_flow_plan\.rb\z},
      %r{\Atest/test_desktop_activation_installer\.rb\z},
      %r{\Atest/test_desktop_activation_rollback\.rb\z},
      %r{\Atest/test_desktop_integration_manifest\.rb\z},
      %r{\Atest/test_desktop_resource_bridge_plan\.rb\z}
    ]
  },
  {
    id: "cw6-snapshot-rollback-store",
    workstream: "CW6 / A4",
    title: "Snapshot and rollback receipt store",
    description: "Content-addressed snapshot and rollback receipt evidence that stays disabled until safe receipt records exist.",
    review_order: 8.5,
    patterns: [
      %r{\Ainternal/runtime/snapshot/},
      %r{\Ainternal/runtime/appidentity/snapshot_},
      %r{\Acmd/xnix-runtime-go/snapshot_restore_candidates_},
      %r{\Atest/test_compatibility_snapshot_plan\.rb\z}
    ]
  },
  {
    id: "cw7-ai-diagnostic-privacy",
    workstream: "CW7 / A7",
    title: "AI diagnostic privacy boundary",
    description: "Fixture-backed diagnostic input, recommendation, repair, and privacy gates with real providers disabled.",
    review_order: 8.6,
    patterns: [
      %r{\Ainternal/runtime/diagnostics/},
      %r{\Ainternal/runtime/appidentity/ai_},
      %r{\Ainternal/runtime/appidentity/diagnostic_},
      %r{\Ainternal/runtime/appidentity/crash_hang_signal_summary},
      %r{\Ainternal/runtime/appidentity/support_bundle_},
      %r{\Ainternal/runtime/appidentity/support_case_timeline},
      %r{\Acmd/xnix-runtime-go/diagnostic_},
      %r{\Acmd/xnix-runtime-go/crash_hang_signal_summary_},
      %r{\Acmd/xnix-runtime-go/support_bundle_},
      %r{\Acmd/xnix-runtime-go/support_case_timeline_},
      %r{\Atest/test_ai_},
      %r{\Atest/test_compatibility_repair_plan\.rb\z}
    ]
  },
  {
    id: "cw10-evidence-drift-harness",
    workstream: "CW10 / A8",
    title: "Evidence and drift harness",
    description: "Reports and tests that prevent orphan contracts, preview-only regressions, and unsafe merge drift.",
    review_order: 9,
    patterns: [
      %r{\Ainternal/runtime/appidentity/offline_application_fixture_matrix},
      %r{\Acmd/xnix-runtime-go/offline_application_fixture_matrix_},
      %r{\Ascripts/implementation_evidence_report\.rb\z},
      %r{\Ascripts/verify_layout\.rb\z},
      %r{\Ascripts/mainline_integration_review\.rb\z},
      %r{\Ascripts/release_evidence_index\.rb\z},
      %r{\Ascripts/offline_application_fixture_matrix\.rb\z},
      %r{\Ascripts/merge_readiness_packet\.rb\z},
      %r{\Ascripts/remote_go_test\.rb\z},
      %r{\Atest/test_implementation_evidence_report\.rb\z},
      %r{\Atest/test_mainline_integration_review\.rb\z},
      %r{\Atest/test_release_evidence_index\.rb\z},
      %r{\Atest/test_offline_application_fixture_matrix\.rb\z},
      %r{\Atest/test_merge_readiness_packet\.rb\z},
      %r{\Ainternal/runtime/appidentity/windows_compatibility_}
    ]
  },
  {
    id: "cw9-windows-app-runner-path",
    workstream: "CW9 / A10",
    title: "Windows app runner path",
    description: "Real Windows executable smoke runner, known-app runner helpers, and product-safe runner evidence without broad backend launch.",
    review_order: 8.8,
    patterns: [
      %r{\Ainternal/runtime/winapp/},
      %r{\Ainternal/runtime/appidentity/q4_sample_notepad_acceptance},
      %r{\Acmd/xnix-runtime-go/windows_compatibility_},
      %r{\Acmd/xnix-runtime-go/q4_sample_notepad_acceptance_},
      %r{\Adocs/windows-app-smoke-profile-runbook\.md\z},
      %r{\Ascripts/winapp_smoke\.rb\z},
      %r{\Ascripts/winapp_container_smoke\.rb\z},
      %r{\Ascripts/winapp_guest_wine_smoke\.rb\z},
      %r{\Ascripts/q4_sample_notepad_smoke\.rb\z},
      %r{\Ascripts/known_winapp_},
      %r{\Atest/test_q4_sample_notepad_smoke_script\.rb\z},
      %r{\Atest/test_known_winapp_guest_wine_smoke_script\.rb\z},
      %r{\Atest/test_winapp_smoke_script\.rb\z},
      %r{\Atest/fixtures/winapp/}
    ]
  },
  {
    id: "cw11-product-image-acceptance",
    workstream: "CW11 / A9",
    title: "Product image and QEMU acceptance",
    description: "Image, full-smoke, restricted Docker, and QEMU acceptance assets. This lane should not run during light review.",
    review_order: 10,
    patterns: [
      %r{\A\.dockerignore\z},
      %r{\Aboot/},
      %r{\Abuildroot/},
      %r{\Adocs/kde-image-pipeline\.md\z},
      %r{\Adocs/release-evidence/},
      %r{\Aimage/},
      %r{\Ainternal/runtime/image/},
      %r{\Alib/xnix/container\.rb\z},
      %r{\Alib/xnix/full_smoke_report\.rb\z},
      %r{\Acmd/xnix-runtime-go/restricted_product_smoke_packet_},
      %r{\Ascripts/container\.rb\z},
      %r{\Ascripts/full_smoke\.rb\z},
      %r{\Ascripts/full_checkpoint_promotion_packet\.rb\z},
      %r{\Ascripts/restricted_product_smoke_packet\.rb\z},
      %r{\Atest/test_full_smoke},
      %r{\Atest/test_full_checkpoint_promotion_packet\.rb\z},
      %r{\Atest/test_restricted_product_smoke_packet\.rb\z},
      %r{\Atest/test_container\.rb\z}
    ]
  }
].freeze

LANE_REVIEW_REQUIREMENTS = {
  "planning-documents" => {
    required_verification: [
      "ruby scripts/verify_layout.rb",
      "git diff --check"
    ],
    required_evidence: [
      "Documents preserve the KDE-first Runtime-owned product goal.",
      "Claude Code handoff text names the protected implementation package boundary.",
      "No task asks for production D-Bus ownership, backend launch, real Portal calls, privileged containers, or host-root mutation."
    ],
    safety_guards: [
      "planning-only",
      "do-not-edit-protected-claude-package",
      "no-runtime-write-enable"
    ]
  },
  "version-and-product-metadata" => {
    required_verification: [
      "ruby scripts/verify_layout.rb",
      "git diff --check"
    ],
    required_evidence: [
      "VERSION, CHANGELOG.md, PRODUCT_OVERVIEW.md, README.md, KDE metadata, and C Runtime version are synchronized.",
      "Release notes describe only the lane being staged."
    ],
    safety_guards: [
      "version-sync-required",
      "no-secrets",
      "no-build-artifacts"
    ]
  },
  "cw1-runtime-owner-read-boundary" => {
    required_verification: [
      "go test ./cmd/xnix-runtime-owner ./internal/runtime/owner ./internal/runtime/appidentity",
      "ruby scripts/runtime_contract_drift_report.rb --format json",
      "ruby -Ilib test/test_runtime_contract_drift_report.rb",
      "ruby -Ilib test/test_runtime_owner_candidate_smoke_script.rb",
      "ruby scripts/verify_layout.rb"
    ],
    required_evidence: [
      "Private session-bus smoke proves read dispatch, disabled writes, unsupported-read rejection, and shutdown.",
      "Production D-Bus ownership remains blocked.",
      "Owner routes cover the Runtime read contract."
    ],
    safety_guards: [
      "production-dbus-owner-disabled",
      "runtime-writes-disabled",
      "backend-launch-disabled"
    ]
  },
  "shared-runtime-cli-plumbing" => {
    required_verification: [
      "go test ./cmd/xnix-runtime-go ./internal/runtime/appidentity",
      "ruby scripts/runtime_contract_drift_report.rb --format json",
      "ruby scripts/verify_layout.rb"
    ],
    required_evidence: [
      "CLI registration exposes only read or preview commands unless a lane explicitly owns a gated writer.",
      "Shared identity output hides raw executable paths, backend commands, and Runtime state-root paths.",
      "C Runtime version assertions still match VERSION."
    ],
    safety_guards: [
      "no-raw-backend-command-output",
      "no-state-root-path-exposure",
      "no-runtime-write-enable"
    ]
  },
  "cw2-recipe-artifact-trust" => {
    required_verification: [
      "go test ./internal/runtime/recipe ./internal/runtime/artifact ./internal/runtime/appidentity ./cmd/xnix-runtime-go",
      "ruby -Ilib test/test_recipe_registry.rb",
      "ruby -Ilib test/test_recipe_trust_policy.rb",
      "ruby -Ilib test/test_recipe_install_gate.rb",
      "ruby -Ilib test/test_compatibility_artifact_manifest.rb",
      "ruby -Ilib test/test_compatibility_install_plan.rb",
      "ruby scripts/implementation_evidence_report.rb --format json",
      "ruby scripts/verify_layout.rb"
    ],
    required_evidence: [
      "Digest-verified local recipe and artifact receipts can unlock install readiness.",
      "Tampered receipts fail closed for digest mismatch, path escape, app-id mismatch, missing artifacts, and unsafe side-effect flags.",
      "Desktop-facing output hides cache roots, fixture roots, host paths, backend commands, and raw executable paths."
    ],
    safety_guards: [
      "network-fetch-disabled",
      "package-manager-disabled",
      "host-root-mutation-disabled"
    ]
  },
  "cw3-runtime-state-backend-lifecycle" => {
    required_verification: [
      "go test ./internal/runtime/environment ./internal/runtime/appidentity ./cmd/xnix-runtime-go",
      "ruby -Ilib test/test_compatibility_backend_lifecycle.rb",
      "ruby -Ilib test/test_compatibility_backend_selection_plan.rb",
      "ruby -Ilib test/test_compatibility_backend_environment_plan.rb",
      "ruby scripts/implementation_evidence_report.rb --format json",
      "ruby scripts/verify_layout.rb"
    ],
    required_evidence: [
      "Backend inventory and lifecycle records persist under an explicit state root.",
      "Repair-required and blocked states expose stable user-safe reasons.",
      "KDE-facing output hides profile paths, state-root paths, raw backend commands, and backend implementation details."
    ],
    safety_guards: [
      "wine-proton-vm-launch-disabled",
      "backend-download-disabled",
      "host-root-mutation-disabled"
    ]
  },
  "cw5-portal-permission-safety" => {
    required_verification: [
      "go test ./internal/runtime/portal ./internal/runtime/appidentity ./cmd/xnix-runtime-go",
      "ruby -Ilib test/test_portal_access_policy.rb",
      "ruby -Ilib test/test_portal_request_model.rb",
      "ruby -Ilib test/test_compatibility_permission_review_plan.rb",
      "ruby scripts/implementation_evidence_report.rb --format json",
      "ruby scripts/verify_layout.rb"
    ],
    required_evidence: [
      "Fake-mode Portal records cover granted, denied, expired, cancelled, malformed, and blocked flows.",
      "Execution preflight can consume receipts without granting real host permissions.",
      "KDE summaries use user terms such as Documents allowed, Camera blocked, and Network allowed."
    ],
    safety_guards: [
      "real-portal-transport-disabled",
      "host-permission-change-disabled",
      "execution-approval-disabled"
    ]
  },
  "cw8-execution-ledger-session-evidence" => {
    required_verification: [
      "go test ./internal/runtime/execution ./internal/runtime/appidentity ./cmd/xnix-runtime-go",
      "ruby -Ilib test/test_launch_request.rb",
      "ruby -Ilib test/test_compatibility_execution_readiness.rb",
      "ruby scripts/implementation_evidence_report.rb --format json",
      "ruby scripts/verify_layout.rb"
    ],
    required_evidence: [
      "Reviewed execution transactions persist as records while launch remains disabled.",
      "Session evidence can feed KDE task-manager, KWin, tray, and Compatibility Center read models.",
      "Denied Portal receipts block execution evidence."
    ],
    safety_guards: [
      "launch-disabled",
      "backend-process-start-disabled",
      "window-observation-disabled"
    ]
  },
  "cw4-kde-entrypoint-consumers" => {
    required_verification: [
      "go test ./internal/runtime/appidentity ./cmd/xnix-runtime-go",
      "ruby -Ilib test/test_kde_integration_status.rb",
      "ruby -Ilib test/test_kde_shell_integration_plan.rb",
      "ruby -Ilib test/test_kde_application_surface_plan.rb",
      "ruby -Ilib test/test_task_manager_identity.rb",
      "ruby -Ilib test/test_kwin_window_rule.rb",
      "ruby -Ilib test/test_tray_status_model.rb",
      "ruby -Ilib test/test_file_association_model.rb",
      "ruby -Ilib test/test_settings_model.rb",
      "ruby scripts/verify_layout.rb"
    ],
    required_evidence: [
      "The seven official KDE entry points consume Runtime evidence rather than owning backend policy.",
      "User-facing labels hide Wine prefix, raw executable, backend command, state-root, and host path terms.",
      "KDE output remains presentation-only."
    ],
    safety_guards: [
      "kde-policy-ownership-disabled",
      "raw-executable-hidden",
      "backend-details-hidden"
    ]
  },
  "cw6-snapshot-rollback-store" => {
    required_verification: [
      "go test ./internal/runtime/snapshot ./internal/runtime/appidentity",
      "ruby -Ilib test/test_compatibility_snapshot_plan.rb",
      "ruby scripts/implementation_evidence_report.rb --format json",
      "ruby scripts/verify_layout.rb"
    ],
    required_evidence: [
      "Snapshot and rollback records are content-addressed and receipt-verified.",
      "Rollback plans expose user-safe state without mutating host files.",
      "Unsafe restore behavior remains gated."
    ],
    safety_guards: [
      "restore-disabled",
      "host-root-mutation-disabled",
      "state-root-path-hidden"
    ]
  },
  "cw7-ai-diagnostic-privacy" => {
    required_verification: [
      "go test ./internal/runtime/diagnostics ./internal/runtime/appidentity ./cmd/xnix-runtime-go",
      "ruby -Ilib test/test_ai_diagnostic_input.rb",
      "ruby -Ilib test/test_ai_diagnostic_recommendation.rb",
      "ruby -Ilib test/test_ai_repair_approval_gate.rb",
      "ruby -Ilib test/test_compatibility_repair_plan.rb",
      "ruby scripts/implementation_evidence_report.rb --format json",
      "ruby scripts/verify_layout.rb"
    ],
    required_evidence: [
      "AI diagnostic input is privacy-filtered and fixture-backed.",
      "Recommendations are review-first and cannot execute repairs automatically.",
      "Provider calls remain disabled unless a later explicit task enables a reviewed provider boundary."
    ],
    safety_guards: [
      "real-ai-provider-disabled",
      "auto-repair-disabled",
      "file-content-hidden"
    ]
  },
  "cw10-evidence-drift-harness" => {
    required_verification: [
      "ruby -Ilib test/test_implementation_evidence_report.rb",
      "ruby -Ilib test/test_mainline_integration_review.rb",
      "ruby -Ilib test/test_offline_application_fixture_matrix.rb",
      "ruby -Ilib test/test_merge_readiness_packet.rb",
      "ruby scripts/implementation_evidence_report.rb --format json",
      "ruby scripts/mainline_integration_review.rb --format json",
      "ruby scripts/merge_readiness_packet.rb --format json",
      "ruby scripts/verify_layout.rb",
      "git diff --check"
    ],
    required_evidence: [
      "Reports flag orphan contracts, unclassified paths, protected Claude file changes, and unsafe staging posture.",
      "Windows compatibility workstreams still name KDE-first, Go-owned Runtime lanes.",
      "Review reports remain read-only and never stage files, run Docker, run QEMU, launch backends, or mutate the host root."
    ],
    safety_guards: [
      "read-only-report",
      "never-stage-all",
      "docker-qemu-not-required"
    ]
  },
  "cw9-windows-app-runner-path" => {
    required_verification: [
      "go test ./internal/runtime/winapp -run 'TestRunSmoke' -count=1",
      "go test ./cmd/xnix-runtime-go -run 'TestWindowsAppRunSmokeCommand' -count=1",
      "ruby scripts/verify_layout.rb",
      "git diff --check"
    ],
    required_evidence: [
      "The local runner can execute a Windows .exe through an explicit compatibility runner or skip safely when unavailable.",
      "Product-facing output can redact raw stdout and stderr while preserving marker and size evidence.",
      "Runner evidence hides host paths, raw backend commands, state roots, secrets, tokens, usernames, and environment variables."
    ],
    safety_guards: [
      "host-root-mutation-disabled",
      "raw-output-redaction-available",
      "broad-backend-launch-disabled"
    ]
  },
  "cw11-product-image-acceptance" => {
    required_verification: [
      "ruby scripts/full_smoke.rb",
      "ruby scripts/verify_layout.rb",
      "git diff --check"
    ],
    required_evidence: [
      "Restricted product smoke uses loopback-only SSH and persisted serial logs.",
      "Docker or QEMU execution is explicitly authorized before running.",
      "Smoke validates real Runtime evidence rather than static contracts."
    ],
    safety_guards: [
      "requires-explicit-docker-or-qemu-authorization",
      "loopback-only-network",
      "no-privileged-container"
    ]
  }
}.freeze

def parse_options
  options = { format: "json", status_fixture: nil, repo_fixture: nil }
  OptionParser.new do |parser|
    parser.banner = "Usage: mainline_integration_review.rb [--format json|markdown] [--status-fixture PATH] [--repo-fixture DIR]"
    parser.on("--format FORMAT", "Output format: json or markdown") { |value| options[:format] = value }
    parser.on("--status-fixture PATH", "Read porcelain status lines from PATH instead of git") { |value| options[:status_fixture] = value }
    parser.on("--repo-fixture DIR", "Read VERSION and release metadata from DIR instead of the project root") { |value| options[:repo_fixture] = value }
  end.parse!
  options
end

def git_status_lines
  stdout, stderr, status = Open3.capture3("git", "-C", PROJECT_ROOT.to_s, "status", "--porcelain=v1", "-uall")
  abort "git status failed: #{stderr}" unless status.success?

  stdout.lines
end

def status_lines(options)
  return Pathname.new(options[:status_fixture]).read.lines if options[:status_fixture]

  git_status_lines
end

def normalize_status_path(line)
  raw = line.chomp
  return nil if raw.empty?

  status = raw[0, 2]
  path = raw[3..] || ""
  path = path.split(" -> ", 2).last if path.include?(" -> ")
  { status: status, path: path }
end

def excluded_path?(path)
  EXCLUDED_PREFIXES.any? { |prefix| path.start_with?(prefix) }
end

def classify_path(path)
  return "blocked-protected-claude-owned-file" if path == PROTECTED_CLAUDE_FILE

  lane = LANES.find { |candidate| candidate[:patterns].any? { |pattern| path.match?(pattern) } }
  lane ? lane[:id] : "unclassified"
end

def risky_path_reason(path)
  rule = RISKY_PATH_RULES.find { |candidate| path.match?(candidate[:pattern]) }
  rule && rule[:reason]
end

def code_change_path?(path)
  CODE_CHANGE_PREFIXES.any? { |prefix| path.start_with?(prefix) }
end

def repo_metadata_root(options)
  options[:repo_fixture] ? Pathname.new(options[:repo_fixture]) : PROJECT_ROOT
end

def current_repo_version(root)
  version_file = root.join("VERSION")
  return nil unless version_file.file?

  version_file.read.strip
end

def stale_version_strings(root, version)
  return [{ file: "VERSION", expected: nil, reason: "missing-version-file" }] unless version

  stale = []
  changelog = root.join("CHANGELOG.md")
  if changelog.file? && !changelog.read.include?("## [#{version}]")
    stale << { file: "CHANGELOG.md", expected: version, reason: "missing-current-version-entry" }
  end
  VERSION_SYNC_FILES.each do |relative|
    file = root.join(relative)
    next unless file.file?

    stale << { file: relative, expected: version, reason: "missing-current-version-string" } unless file.read.include?(version)
  end
  stale
end

def build_intake_guardrails(included, root)
  risky_paths = included.filter_map do |entry|
    reason = risky_path_reason(entry[:path])
    reason && { path: entry[:path], reason: reason }
  end
  code_change_detected = included.any? { |entry| code_change_path?(entry[:path]) }
  version_bump_included = included.any? { |entry| entry[:path] == "VERSION" }
  changelog_update_included = included.any? { |entry| entry[:path] == "CHANGELOG.md" }
  version = current_repo_version(root)
  stale_versions = stale_version_strings(root, version)

  {
    code_change_detected: code_change_detected,
    version_bump_included: version_bump_included,
    changelog_update_included: changelog_update_included,
    missing_version_bump: code_change_detected && !(version_bump_included && changelog_update_included),
    current_version: version,
    stale_version_detected: stale_versions.any?,
    stale_version_strings: stale_versions,
    risky_path_count: risky_paths.length,
    risky_paths: risky_paths
  }
end

def build_staging_recommendation(lanes, excluded, risky_paths)
  risky_by_path = risky_paths.to_h { |entry| [entry[:path], entry[:reason]] }
  protected_lane = lanes.find { |lane| lane[:id] == "blocked-protected-claude-owned-file" }
  unclassified_lane = lanes.find { |lane| lane[:id] == "unclassified" }

  exclude = excluded.map { |entry| { path: entry[:path], reason: "excluded-prefix" } }
  exclude += risky_paths
  exclude += (protected_lane ? protected_lane[:paths] : []).map do |path|
    { path: path, reason: "protected-claude-owned-file" }
  end
  exclude += (unclassified_lane ? unclassified_lane[:paths] : []).reject { |path| risky_by_path.key?(path) }.map do |path|
    { path: path, reason: "manual-classification-required" }
  end

  stage_by_lane = lanes.select { |lane| lane[:file_count].positive? }
                       .reject { |lane| %w[blocked-protected-claude-owned-file unclassified].include?(lane[:id]) }
                       .map do |lane|
    { id: lane[:id], workstream: lane[:workstream], paths: lane[:paths].reject { |path| risky_by_path.key?(path) } }
  end

  {
    never_stage_all: true,
    stage_by_lane: stage_by_lane,
    exclude: exclude
  }
end

def lane_definitions
  LANES.map do |lane|
    lane.slice(:id, :workstream, :title, :description, :review_order)
        .merge(LANE_REVIEW_REQUIREMENTS.fetch(lane[:id]))
  end + [
    {
      id: "blocked-protected-claude-owned-file",
      workstream: "blocked",
      title: "Protected Claude-owned file",
      description: "Changes to docs/claude-code-implementation-packages.md must be reviewed outside this checkpoint.",
      review_order: 98,
      required_verification: [],
      required_evidence: [
        "Do not stage this file from Codex-owned integration batches.",
        "Ask the user or Claude owner to reconcile the protected document separately."
      ],
      safety_guards: [
        "protected-claude-file",
        "manual-review-required"
      ]
    },
    {
      id: "unclassified",
      workstream: "unknown",
      title: "Unclassified",
      description: "Files that do not match a mainline lane and need manual review before staging.",
      review_order: 99,
      required_verification: [],
      required_evidence: [
        "Classify or remove these files before staging.",
        "Do not merge side quests into the KDE-first Runtime mainline."
      ],
      safety_guards: [
        "manual-classification-required",
        "never-stage-all"
      ]
    }
  ]
end

def build_report(options)
  parsed = status_lines(options).filter_map { |line| normalize_status_path(line) }
  excluded = parsed.select { |entry| excluded_path?(entry[:path]) }
  included = parsed.reject { |entry| excluded_path?(entry[:path]) }

  lane_map = lane_definitions.to_h { |lane| [lane[:id], lane.merge(files: [])] }
  included.each do |entry|
    lane_id = classify_path(entry[:path])
    lane_map.fetch(lane_id)[:files] << entry
  end

  lanes = lane_map.values
                  .select { |lane| lane[:files].any? || %w[blocked-protected-claude-owned-file unclassified].include?(lane[:id]) }
                  .sort_by { |lane| lane[:review_order] }
                  .map do |lane|
    files = lane.fetch(:files)
    lane.merge(
      file_count: files.length,
      paths: files.map { |entry| entry[:path] },
      statuses: files.map { |entry| entry[:status] }.uniq.sort
    )
  end

  protected_lane = lanes.find { |lane| lane[:id] == "blocked-protected-claude-owned-file" }
  unclassified_lane = lanes.find { |lane| lane[:id] == "unclassified" }
  intake_guardrails = build_intake_guardrails(included, repo_metadata_root(options))
  staging_recommendation = build_staging_recommendation(lanes, excluded, intake_guardrails.fetch(:risky_paths))

  {
    version: VERSION,
    schema_version: "xnix.runtime.mainline_integration_review.v1",
    report_type: "mainline-integration-review",
    source: options[:status_fixture] ? "status-fixture+mainline-integration-checkpoint" : "git-status+mainline-integration-checkpoint",
    checkpoint_document: "docs/mainline-integration-checkpoint.md",
    protected_claude_file: PROTECTED_CLAUDE_FILE,
    protected_claude_file_modified: protected_lane && protected_lane[:file_count].positive?,
    excluded_prefixes: EXCLUDED_PREFIXES,
    excluded_file_count: excluded.length,
    changed_file_count: included.length,
    unclassified_file_count: unclassified_lane ? unclassified_lane[:file_count] : 0,
    safe_to_stage_all: false,
    docker_or_qemu_required: false,
    host_root_modified: false,
    privileged_container_required: false,
    backend_launch_enabled: false,
    lanes: lanes,
    review_matrix: lanes.reject { |lane| lane[:file_count].zero? }.map do |lane|
      lane.slice(:id, :workstream, :title, :required_verification, :required_evidence, :safety_guards)
    end,
    review_order: lanes.reject { |lane| lane[:file_count].zero? }.map { |lane| lane[:id] },
    intake_guardrails: intake_guardrails,
    staging_recommendation: staging_recommendation,
    next_action: "Review and stage one lane at a time; never stage excluded paths or the protected Claude-owned file."
  }
end

def render_markdown(report)
  lines = []
  lines << "# Mainline Integration Review"
  lines << ""
  lines << "- Version: #{report.fetch(:version)}"
  lines << "- Source: #{report.fetch(:source)}"
  lines << "- Changed files reviewed: #{report.fetch(:changed_file_count)}"
  lines << "- Excluded files: #{report.fetch(:excluded_file_count)}"
  lines << "- Protected Claude file modified: #{report.fetch(:protected_claude_file_modified)}"
  lines << "- Safe to stage all: #{report.fetch(:safe_to_stage_all)}"
  guardrails = report.fetch(:intake_guardrails)
  lines << "- Missing version bump: #{guardrails.fetch(:missing_version_bump)}"
  lines << "- Stale version strings: #{guardrails.fetch(:stale_version_strings).length}"
  lines << "- Risky paths: #{guardrails.fetch(:risky_path_count)}"
  lines << ""
  lines << "## Lanes"
  lines << ""
  report.fetch(:lanes).each do |lane|
    next if lane.fetch(:file_count).zero?

    lines << "### #{lane.fetch(:id)}"
    lines << ""
    lines << "- Workstream: #{lane.fetch(:workstream)}"
    lines << "- Files: #{lane.fetch(:file_count)}"
    lines << ""
    lane.fetch(:files).each do |entry|
      status = entry.fetch(:status).strip.empty? ? "modified" : entry.fetch(:status).strip
      lines << "- `#{status}` `#{entry.fetch(:path)}`"
    end
    unless lane.fetch(:required_verification).empty?
      lines << ""
      lines << "Required verification:"
      lane.fetch(:required_verification).each { |command| lines << "- `#{command}`" }
    end
    unless lane.fetch(:required_evidence).empty?
      lines << ""
      lines << "Required evidence:"
      lane.fetch(:required_evidence).each { |evidence| lines << "- #{evidence}" }
    end
    unless lane.fetch(:safety_guards).empty?
      lines << ""
      lines << "Safety guards:"
      lane.fetch(:safety_guards).each { |guard| lines << "- `#{guard}`" }
    end
    lines << ""
  end
  lines << "## Intake Guardrails"
  lines << ""
  lines << "- Code change detected: #{guardrails.fetch(:code_change_detected)}"
  lines << "- Version bump included: #{guardrails.fetch(:version_bump_included)}"
  lines << "- Changelog update included: #{guardrails.fetch(:changelog_update_included)}"
  lines << "- Missing version bump: #{guardrails.fetch(:missing_version_bump)}"
  lines << "- Current version: #{guardrails.fetch(:current_version)}"
  lines << "- Stale version detected: #{guardrails.fetch(:stale_version_detected)}"
  guardrails.fetch(:stale_version_strings).each do |entry|
    lines << "- Stale version string: `#{entry.fetch(:file)}` (#{entry.fetch(:reason)})"
  end
  guardrails.fetch(:risky_paths).each do |entry|
    lines << "- Risky path: `#{entry.fetch(:path)}` (#{entry.fetch(:reason)})"
  end
  lines << ""
  lines << "## Staging Recommendation"
  lines << ""
  recommendation = report.fetch(:staging_recommendation)
  lines << "Never stage all: #{recommendation.fetch(:never_stage_all)}. Stage one lane at a time."
  lines << ""
  recommendation.fetch(:stage_by_lane).each do |lane|
    lines << "### Stage `#{lane.fetch(:id)}`"
    lines << ""
    lane.fetch(:paths).each { |path| lines << "- `#{path}`" }
    lines << ""
  end
  unless recommendation.fetch(:exclude).empty?
    lines << "### Exclude"
    lines << ""
    recommendation.fetch(:exclude).each do |entry|
      lines << "- `#{entry.fetch(:path)}` (#{entry.fetch(:reason)})"
    end
    lines << ""
  end
  lines << "## Next Action"
  lines << ""
  lines << report.fetch(:next_action)
  lines.join("\n")
end

options = parse_options
report = build_report(options)

case options[:format]
when "json"
  puts JSON.pretty_generate(report)
when "markdown"
  puts render_markdown(report)
else
  abort "unsupported format: #{options[:format]}"
end
