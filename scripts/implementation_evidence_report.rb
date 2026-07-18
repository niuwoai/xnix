#!/usr/bin/env ruby
# frozen_string_literal: true

require "json"
require "optparse"
require "pathname"

PROJECT_ROOT = Pathname.new(__dir__).join("..").realpath
MAINLINE_DOCUMENT = "docs/claude-code-mainline-implementation-plan.md"
WINDOWS_WORKSTREAM_DOCUMENT = "docs/claude-code-windows-compatibility-workstreams.md"
MAINLINE_FIRST_WAVE = [
  {
    mainline_package: "M1",
    suggested_branch: "codex/runtime-owner-read-service",
    reason: "Runtime-owned read paths should replace preview-only dispatch before broader product ownership advances.",
    minimal_mergeable_outcome: "A constrained session-bus owner serves the Go owner read dispatch table for every D-Bus read-only Runtime method and fails every write method closed."
  },
  {
    mainline_package: "M2",
    suggested_branch: "codex/recipe-artifact-trust-pipeline",
    reason: "Install, environment, activation, and execution work need trusted local inputs first.",
    minimal_mergeable_outcome: "Local recipes and fixture artifacts verify digests, stage under a controlled root, and produce blocked receipts for invalid inputs."
  },
  {
    mainline_package: "M3",
    suggested_branch: "codex/environment-lifecycle-state",
    reason: "Execution readiness needs durable environment state before launch paths become meaningful.",
    minimal_mergeable_outcome: "A state-root lifecycle store records missing, planned, staged, ready, repair-required, and blocked states."
  },
  {
    mainline_package: "M8",
    suggested_branch: "codex/implementation-evidence-harness",
    reason: "Contract-heavy work needs an evidence harness that keeps empty domains visible.",
    minimal_mergeable_outcome: "JSON and Markdown reports classify domain evidence and fail on orphan contract-only surfaces."
  }
].freeze
WINDOWS_COMPATIBILITY_FIRST_WAVE = [
  {
    workstream: "CW1",
    mainline_package: "M1",
    suggested_branch: "codex/cw-runtime-owner-read-boundary",
    reason: "Windows compatibility needs a Go-owned Runtime read boundary before KDE and execution surfaces depend on it.",
    minimal_mergeable_outcome: "A constrained Go Runtime owner serves read-only Runtime methods through owner routes and keeps writes disabled."
  },
  {
    workstream: "CW2",
    mainline_package: "M2",
    suggested_branch: "codex/cw-recipe-artifact-trust",
    reason: "KDE presentation and execution planning must consume digest-verified recipes and artifacts.",
    minimal_mergeable_outcome: "Local recipes and fixture artifacts verify digests, stage under explicit roots, and fail closed for invalid inputs."
  },
  {
    workstream: "CW3",
    mainline_package: "M3",
    suggested_branch: "codex/cw-runtime-state-lifecycle",
    reason: "Windows application status, diagnostics, and future launch gates need durable state-root lifecycle evidence.",
    minimal_mergeable_outcome: "Runtime state-root lifecycle records expose missing, planned, staged, ready, repair-required, and blocked states without backend launch."
  },
  {
    workstream: "CW10",
    mainline_package: "M8",
    suggested_branch: "codex/cw-evidence-drift-harness",
    reason: "The KDE-first Windows compatibility strategy needs CI-friendly evidence gates to prevent more contract-only expansion.",
    minimal_mergeable_outcome: "Reports fail on orphan Runtime contracts, KDE entry-point drift, preview-only regressions, and missing workstream ownership."
  }
].freeze
WRITE_METHODS = %w[
  InstallRecipe
  Launch
  CreateSnapshot
  RestoreSnapshot
].freeze

STATUS_ORDER = {
  "missing" => 0,
  "contract-only" => 1,
  "fixture-implemented" => 2,
  "state-root-implemented" => 3,
  "smoke-owned" => 4,
  "production-gated" => 5
}.freeze

DOMAIN_DEFINITIONS = [
  {
    id: "runtime-owner-service",
    name: "Runtime owner service",
    package: "P1",
    mainline_package: "M1",
    contract_files: %w[
      runtime/dbus/org.xnix.Compatibility1.xml
      internal/runtime/appidentity/runtime_owner_readiness.go
      internal/runtime/appidentity/runtime_owner_route_manifest.go
      internal/runtime/appidentity/runtime_service_activation_preflight.go
    ],
    fixture_files: %w[
      cmd/xnix-runtime-owner/main.go
      internal/runtime/owner/candidate.go
      internal/runtime/owner/dispatch.go
      internal/runtime/owner/dispatch_test.go
      internal/runtime/owner/service.go
      internal/runtime/owner/service_test.go
      internal/runtime/owner/route_checkpoint.go
      internal/runtime/owner/route_checkpoint_test.go
      internal/runtime/owner/kde_offline_identity_checkpoint.go
      internal/runtime/owner/kde_offline_identity_checkpoint_test.go
      internal/runtime/owner/lifecycle.go
      internal/runtime/owner/lifecycle_test.go
      internal/runtime/owner/smoke_batch.go
      internal/runtime/owner/smoke_batch_test.go
      internal/runtime/owner/kde_test_launch_materialization_fanout_owner_smoke_coverage.go
      internal/runtime/owner/kde_test_launch_materialization_fanout_owner_smoke_coverage_test.go
      internal/runtime/owner/restricted_smoke_receipt_fanout_owner_smoke_coverage.go
      internal/runtime/owner/restricted_smoke_receipt_fanout_owner_smoke_coverage_test.go
      internal/runtime/owner/backend_adapter_redacted_profile_owner_smoke_coverage.go
      internal/runtime/owner/backend_adapter_redacted_profile_owner_smoke_coverage_test.go
      internal/runtime/owner/restricted_smoke_receipt.go
      internal/runtime/owner/restricted_smoke_receipt_test.go
      internal/runtime/owner/restricted_smoke_receipt_fanout.go
      internal/runtime/owner/restricted_smoke_receipt_fanout_test.go
      internal/runtime/owner/restricted_smoke_receipt_fanout_owner_route_audit.go
      internal/runtime/owner/restricted_smoke_receipt_fanout_owner_route_audit_test.go
      internal/runtime/owner/session_bus.go
      internal/runtime/owner/session_bus_test.go
      cmd/xnix-runtime-go/restricted_owner_smoke_receipt_commands.go
      cmd/xnix-runtime-go/restricted_owner_smoke_receipt_cli_test.go
    ],
    state_files: [],
    smoke_files: %w[
      runtime/dbus/xnix_compatd_smoke.c
      runtime/dbus/xnix_compatd_runtime_models.inc
      runtime/dbus/xnix_compatd_kde_center.inc
      scripts/runtime_owner_candidate_smoke.rb
      scripts/dbus_session_smoke.rb
      test/test_runtime_owner_candidate_smoke_script.rb
    ],
    gate_tokens: {
      "cmd/xnix-runtime-owner/main.go" => %w[smoke-owner deny-write materialization-fanout-owner-smoke-coverage restricted-smoke-fanout-owner-smoke-coverage redacted-adapter-profile-owner-smoke-coverage NewKDETestLaunchMaterializationFanOutOwnerSmokeCoveragePreview NewRestrictedOwnerSmokeReceiptFanOutOwnerSmokeCoveragePreview NewBackendAdapterRedactedProfileOwnerSmokeCoveragePreview],
      "internal/runtime/owner/service.go" => %w[xnix.runtime.owner_service_call.v1 runtime-owner-service-call go-runtime-owner-in-process-service Service],
      "internal/runtime/owner/route_checkpoint.go" => %w[xnix.runtime.owner_route_checkpoint.v1 runtime-owner-route-checkpoint go-owner-read-route-band-checkpoint NewRouteCheckpoint FormalReadRouteCount OwnerReadMethodCount OwnerLocalReadMethodCount SmokeReadRecordCount SmokeWriteDenialCount MethodParityReady FormalRouteCoverageReady OwnerLocalRouteCoverageReady SmokeBatchCoverageReady DeterministicWriteDenialsReady RouteBandReady ProductionBusClaimed WriteMethodsEnabled HostRootModified],
      "internal/runtime/owner/kde_offline_identity_checkpoint.go" => %w[xnix.runtime.kde_offline_identity_checkpoint.v1 kde-offline-identity-checkpoint offline-kde-application-identity-band-checkpoint NewKDEOfflineIdentityCheckpoint recipe-trust surface-coverage identity-parity owner-route owner-service side-effects-disabled OfflineIdentityReady ProductionSignatureReady SurfaceCount ExpectedSurfaceCount FormalReadRouteCount OwnerReadMethodCount OwnerLocalReadMethodCount SmokeReadRecordCount SmokeWriteDenialCount DesktopFilesWritten MIMEDefaultsWritten KRunnerIndexPersisted LiveTrayBridgeEnabled NotificationSent SettingsPersisted CompatibilityCenterPersisted LaunchEnabled ExecutionStarted BackendProcessStarted ProductionBusClaimed WriteMethodsEnabled HostRootModified],
      "internal/runtime/owner/lifecycle.go" => %w[xnix.runtime.owner_lifecycle_event.v1 runtime-owner-lifecycle-event preview-complete],
      "internal/runtime/owner/smoke_batch.go" => %w[xnix.runtime.owner_smoke_batch.v1 runtime-owner-smoke-batch-record restricted-session-owner-call-batch NewService service.Call runtime-owner-service-call],
      "internal/runtime/owner/kde_test_launch_materialization_fanout_owner_smoke_coverage.go" => %w[xnix.runtime.kde_test_launch_materialization_fanout_owner_smoke_coverage.v1 kde-test-launch-materialization-fanout-owner-smoke-coverage-preview NewKDETestLaunchMaterializationFanOutOwnerSmokeCoveragePreview NewKDETestLaunchMaterializationFanOutOwnerSmokeCoveragePreviewFromRecords materialization-fanout-owner-smoke-coverage GetKDETestLaunchMaterializationFanOut owner-smoke-batch+runtime-owner-service-call+kde-test-launch-materialization-fanout-owner-route smoke-record-present service-call-read-dispatch owner-route-payload-present opaque-receipt-preserved missing-receipt-fail-closed desktop-side-effects-disabled production-dbus-blocked unsafe-gates-closed SmokeCoverageReady StateRootWritesEnabled RuntimeWritesEnabled FanOutWritesEnabled ProductionBusClaimed WriteMethodsEnabled BackendLaunchEnabled HostRootModified],
      "internal/runtime/owner/kde_test_launch_materialization_fanout_owner_smoke_coverage_test.go" => %w[TestKDETestLaunchMaterializationFanOutOwnerSmokeCoveragePreviewCoversOwnerRoute TestKDETestLaunchMaterializationFanOutOwnerSmokeCoveragePreviewFailsClosedWithoutRecord SmokeCoverageReady],
      "internal/runtime/owner/restricted_smoke_receipt_fanout_owner_smoke_coverage.go" => %w[xnix.runtime.restricted_owner_smoke_receipt_fanout_owner_smoke_coverage.v1 restricted-owner-smoke-receipt-fanout-owner-smoke-coverage-preview NewRestrictedOwnerSmokeReceiptFanOutOwnerSmokeCoveragePreview NewRestrictedOwnerSmokeReceiptFanOutOwnerSmokeCoveragePreviewFromRecords restricted-owner-smoke-fanout-owner-smoke-coverage GetRestrictedOwnerSmokeReceiptFanOut owner-smoke-batch+runtime-owner-service-call+restricted-owner-smoke-receipt-fanout-owner-route smoke-record-present service-call-read-dispatch owner-route-payload-present opaque-receipt-preserved missing-receipt-fail-closed support-side-effects-disabled production-dbus-blocked unsafe-gates-closed SmokeCoverageReady StateRootWritesEnabled RuntimeWritesEnabled FanOutWritesEnabled ProductionBusClaimed WriteMethodsEnabled SupportBundleExported SupportCaseCreated NotificationSent BackendLaunchEnabled HostRootModified],
      "internal/runtime/owner/restricted_smoke_receipt_fanout_owner_smoke_coverage_test.go" => %w[TestRestrictedOwnerSmokeReceiptFanOutOwnerSmokeCoveragePreviewCoversOwnerRoute TestRestrictedOwnerSmokeReceiptFanOutOwnerSmokeCoveragePreviewFailsClosedWithoutRecord SmokeCoverageReady],
      "internal/runtime/owner/backend_adapter_redacted_profile_owner_smoke_coverage.go" => %w[xnix.runtime.backend_adapter_redacted_profile_owner_smoke_coverage.v1 backend-adapter-redacted-profile-owner-smoke-coverage-preview NewBackendAdapterRedactedProfileOwnerSmokeCoveragePreview NewBackendAdapterRedactedProfileOwnerSmokeCoveragePreviewFromRecords redacted-adapter-profile-owner-smoke-coverage GetBackendAdapterProfileAudit owner-smoke-batch+runtime-owner-service-call+backend-adapter-redacted-profile-audit-owner-route smoke-record-present service-call-read-dispatch owner-route-payload-present redacted-profiles-preserved full-contract-fixture-local production-dbus-blocked caller-paths-hidden unsafe-gates-closed SmokeCoverageReady StateRootWritesEnabled RuntimeWritesEnabled AdapterInvocationEnabled BackendInstallEnabled BackendDownloadEnabled BackendLaunchEnabled HostRootModified],
      "internal/runtime/owner/backend_adapter_redacted_profile_owner_smoke_coverage_test.go" => %w[TestBackendAdapterRedactedProfileOwnerSmokeCoveragePreviewCoversOwnerRoute TestBackendAdapterRedactedProfileOwnerSmokeCoveragePreviewFailsClosedWithoutRecord SmokeCoverageReady],
      "internal/runtime/owner/restricted_smoke_receipt.go" => %w[xnix.runtime.restricted_owner_smoke_receipt.v1 restricted-owner-smoke-execution-receipt RecordRestrictedOwnerSmokeReceipt runtime-service-activation-preflight+runtime-owner-smoke-batch RestrictedOwnerSmokeMode RestrictedOwnerSmokeDirective authorize-restricted-owner-smoke restricted-owner-smoke-ready runtime-owner-smoke-batch-record StateRootWriteScope ReceiptPersisted ReceiptReadBack ProductionActivationReady ProductionOwnerEnabled SystemServiceStarted SessionBusClaimed ProductionBusClaimed WriteMethodsEnabled BackendLaunchEnabled HostRootModified validateRestrictedOwnerSmokeReceipt],
      "internal/runtime/owner/restricted_smoke_receipt_test.go" => %w[TestRecordRestrictedOwnerSmokeReceiptConsumesPreflightAndSmokeBatch TestRecordRestrictedOwnerSmokeReceiptRequiresExplicitAuthorization TestRecordRestrictedOwnerSmokeReceiptRejectsUnreadyPreflight TestRestrictedOwnerSmokeReceiptRejectsTampering],
      "internal/runtime/owner/restricted_smoke_receipt_fanout.go" => %w[xnix.runtime.restricted_owner_smoke_receipt_fanout.v1 restricted-owner-smoke-receipt-fanout-preview NewRestrictedOwnerSmokeReceiptFanOutPreview xnix.runtime.restricted_owner_smoke_receipt_fanout_owner_route.v1 restricted-owner-smoke-receipt-fanout-owner-route-preview NewRestrictedOwnerSmokeReceiptFanOutOwnerRoutePreview owner-local-restricted-smoke-receipt-fanout missing-receipt-fail-closed owner-route-ready-production-dbus-blocked restricted-smoke-receipt-ready runtime-owner-readiness service-activation-preflight compatibility-onboarding support-bundle-manifest support-case-timeline receipt-consumed readiness-surface-coverage support-surface-coverage surfaces-read-only ownership-boundary-closed support-side-effects-disabled unsafe-data-hidden host-boundary-closed StateRootWritesEnabled RuntimeWritesEnabled FanOutWritesEnabled SupportBundleExported SupportCaseCreated NotificationSent ProductionBusClaimed WriteMethodsEnabled BackendLaunchEnabled HostRootModified],
      "internal/runtime/owner/restricted_smoke_receipt_fanout_test.go" => %w[TestRestrictedOwnerSmokeReceiptFanOutPreviewCoversReadinessAndSupportSurfaces TestRestrictedOwnerSmokeReceiptFanOutPreviewRequiresExistingReceipt TestRestrictedOwnerSmokeReceiptFanOutPreviewRejectsTamperedReceipt],
      "internal/runtime/owner/restricted_smoke_receipt_fanout_owner_route_test.go" => %w[TestRestrictedOwnerSmokeReceiptFanOutOwnerRoutePreviewUsesOpaqueLookup TestRestrictedOwnerSmokeReceiptFanOutOwnerRouteRejectsUnknownOpaqueID missing-receipt-fail-closed caller-state-root-hidden unsafe-gates-closed],
      "internal/runtime/owner/restricted_smoke_receipt_lookup.go" => %w[xnix.runtime.restricted_owner_smoke_receipt_lookup.v1 restricted-owner-smoke-receipt-lookup-preview ResolveRestrictedOwnerSmokeReceipt owner-managed-receipt-lookup GetRestrictedOwnerSmokeReceiptLookup GetRestrictedOwnerSmokeReceiptLookupPreview opaque_receipt_id restricted-owner-smoke-receipt-id missing-receipt CallerStateRootRequired OwnerManagedLookup OpaqueReceiptIDSupported StateRootWritesEnabled FanOutWritesEnabled BackendLaunchEnabled HostRootModified],
      "internal/runtime/owner/restricted_smoke_receipt_lookup_test.go" => %w[TestResolveRestrictedOwnerSmokeReceiptReturnsOpaqueLookup TestResolveRestrictedOwnerSmokeReceiptRejectsUnknownOpaqueID missing-receipt caller-state-root-hidden unsafe-gates-closed],
      "internal/runtime/owner/restricted_smoke_receipt_fanout_owner_route_audit.go" => %w[xnix.runtime.restricted_owner_smoke_receipt_fanout_owner_route_audit.v1 restricted-owner-smoke-receipt-fanout-owner-route-audit-preview NewRestrictedOwnerSmokeReceiptFanOutOwnerRouteAuditPreview GetRestrictedOwnerSmokeReceiptFanOutOwnerRouteAudit owner-local-route-smoke-covered owner-local-read-route-smoke-covered-production-dbus-blocked restricted-owner-smoke-fanout-owner-smoke-coverage owner-smoke-coverage-present owner-route-present caller-state-root-boundary owner-managed-receipt-lookup route-decision ConsumesExistingReceipt RequiresCallerStateRoot ReceiptLookupOwnerManaged OpaqueReceiptIDSupported OwnerSmokeCoverageReady OwnerDispatchRoutePresent FanOutWritesEnabled StateRootWritesEnabled BackendLaunchEnabled HostRootModified],
      "internal/runtime/owner/restricted_smoke_receipt_fanout_owner_route_audit_test.go" => %w[TestRestrictedOwnerSmokeReceiptFanOutOwnerRouteAuditRecognizesOpaqueLookup TestRestrictedOwnerSmokeReceiptFanOutOwnerRouteAuditFailsClosedWithoutSources owner-local-route-smoke-covered caller-state-root-boundary owner-managed-receipt-lookup owner-smoke-coverage-present],
      "cmd/xnix-runtime-go/restricted_owner_smoke_receipt_commands.go" => %w[restricted-owner-smoke-receipt-record runRestrictedOwnerSmokeReceiptRecord authorize-restricted-owner-smoke restricted-owner-smoke-receipt-lookup-preview runRestrictedOwnerSmokeReceiptLookupPreview restricted-owner-smoke-receipt-fanout-owner-route-preview runRestrictedOwnerSmokeReceiptFanOutOwnerRoutePreview restricted-owner-smoke-receipt-fanout-owner-route-audit-preview runRestrictedOwnerSmokeReceiptFanOutOwnerRouteAuditPreview],
      "cmd/xnix-runtime-go/restricted_owner_smoke_receipt_cli_test.go" => %w[TestRestrictedOwnerSmokeReceiptRecordCommandPersistsReceipt TestRestrictedOwnerSmokeReceiptRecordCommandRequiresAuthorization TestRestrictedOwnerSmokeReceiptFanOutPreviewCommandCoversReadinessAndSupportSurfaces TestRestrictedOwnerSmokeReceiptLookupPreviewCommand TestRestrictedOwnerSmokeReceiptFanOutOwnerRoutePreviewCommand TestRestrictedOwnerSmokeReceiptFanOutOwnerRouteAuditPreviewCommand xnix.runtime.restricted_owner_smoke_receipt.v1 xnix.runtime.restricted_owner_smoke_receipt_fanout.v1 xnix.runtime.restricted_owner_smoke_receipt_fanout_owner_route.v1 xnix.runtime.restricted_owner_smoke_receipt_lookup.v1 xnix.runtime.restricted_owner_smoke_receipt_fanout_owner_route_audit.v1],
      "internal/runtime/owner/session_bus.go" => %w[xnix.runtime.owner_session_bus_smoke.v1 runtime-owner-session-bus-smoke-step restricted-private-session-bus-owner-smoke NewSessionBusSmokeTranscript private-session-bus-smoke],
      "runtime/dbus/xnix_compatd_smoke.c" => %w[go_owner_read_dispatch go_owner_service_call5 add_go_owner_dispatch_bridge_fields add_go_owner_dispatch_bridge_fields2 add_go_owner_dispatch_bridge_fields3 add_go_owner_dispatch_bridge_fields5 add_go_owner_dispatch_payload_fields add_go_owner_dispatch_payload_fields2 add_go_owner_dispatch_payload_fields3 go_owner_write_gate_dispatch go_owner_dispatch_available go_owner_dispatch_json go_owner_service_call_available go_owner_service_call_request_type go_owner_service_call_json xnix.runtime.owner_service_call.v1 runtime-owner-service-call go-runtime-owner-dispatch+c-smoke-bridge ListApplications GetApplication GetDiagnostics GetEngineCatalog GetRunPlan GetDesktopActivationManifest GetDesktopActivationTransactionPreview GetDesktopActivationStatus GetKDEIntegrationStatus GetKDEShellIntegrationPlan GetKDEApplicationSurfacePlan GetDesktopEntryPlan GetDesktopIconPlan GetTaskManagerIdentityPlan GetDesktopResourceBridgePlan GetKWinWindowRulePlan GetFileAssociationPlan GetNotificationPlan GetTrayStatus GetKRunnerQueryPlan GetCompatibilitySettingsChangePlan GetCompatibilityModeSwitchPlan GetCompatibilityPermissionReviewPlan GetCompatibilityReviewFlowPlan GetCompatibilityActionQueue GetCompatibilityActionReviewReceipt],
      "runtime/dbus/xnix_compatd_runtime_models.inc" => %w[GetRuntimeServiceBinding GetRuntimeLiveOwnerGate GetRuntimeOwnerProcess GetRuntimeOwnerSmokePlan GetRuntimeMethodParityManifest GetRuntimeOwnerRouteManifest GetRuntimeOwnerRecipeTrust GetRuntimeOwnerReadiness GetPortalRequestPlan GetApplicationStateRoot GetCompatibilityPackageSource GetCompatibilityAcquisitionPreflight GetCompatibilityArtifactManifest GetCompatibilityInstallPlan GetBackendBinding GetBackendCapabilityMatrix GetBackendSelectionPlan GetBackendLifecycle GetBackendEnvironmentPlan GetRepairPlan GetTestPlan GetTestResult GetExecutionReadiness GetLaunchIntent GetAIDiagnosticInput GetAIDiagnosticRecommendation GetAIRepairApprovalGate GetSnapshotPlan GetPortalAccessPolicy GetCompatibilitySettings GetCompatibilitySettingsChangePlan GetCompatibilityModeSwitchPlan GetCompatibilityPermissionReviewPlan GetCompatibilityReviewFlowPlan GetCompatibilityActionQueue GetCompatibilityActionReviewReceipt GetCompatibilityCenterSummary],
      "runtime/dbus/xnix_compatd_kde_center.inc" => %w[GetKDECenterPage GetKDECenterPageSections GetKDECenterPageSectionDetail add_go_owner_dispatch_payload_fields2 add_go_owner_dispatch_payload_fields3 go-kde-center-page-preview go-kde-center-page-sections-preview go-kde-center-page-section-detail-preview],
      "internal/runtime/appidentity/runtime_owner_readiness.go" => %w[ProductionOwnerEnabled ProductionBusClaimed WriteMethodsEnabled],
      "internal/runtime/appidentity/runtime_service_activation_preflight.go" => %w[xnix.runtime.service_activation_preflight.v1 runtime-service-activation-preflight-preview GetRuntimeServiceActivationPreflight production-runtime-service-activation-preflight restricted-owner-smoke-ready ProductionActivationReady RestrictedSmokeReady HumanAuthorizationRequired SystemServiceStarted ProductionBusClaimed WriteMethodsEnabled BackendLaunchEnabled HostRootModified]
    },
    summary: "The Go owner candidate has an in-process service-call boundary, private session-bus smoke transcript evidence, D-Bus service-call bridge evidence, full D-Bus read dispatch coverage, lifecycle JSONL, service-call smoke-batch read/write evidence, restricted owner smoke execution receipts under explicit state roots, read-only receipt fan-out to Runtime readiness and support surfaces, explicit materialization fan-out owner smoke coverage, explicit restricted owner smoke fan-out smoke coverage, explicit redacted adapter profile owner smoke coverage, and C D-Bus bridge evidence for owner self-description, Runtime owner readiness payloads, foundation catalog/status payloads, KDE shell, desktop activation payloads, seven-entry-point payloads, second-ring KDE resource payloads, install-input payloads, backend lifecycle payloads, execution readiness payloads, AI safety payloads, settings/review payloads, Compatibility Center action payloads, and KDE Compatibility Center page payloads; production ownership remains gated."
  },
  {
    id: "recipe-artifact-trust-pipeline",
    name: "Recipe and artifact trust pipeline",
    package: "P2",
    mainline_package: "M2",
    contract_files: %w[
      runtime/recipes/registry.json
      internal/runtime/appidentity/runtime_owner_recipe_trust.go
      internal/runtime/appidentity/artifact_manifest.go
      internal/runtime/appidentity/acquisition_preflight.go
      internal/runtime/appidentity/install_plan.go
      internal/runtime/appidentity/application_upgrade_impact.go
    ],
    fixture_files: %w[
      internal/runtime/recipe/store.go
      internal/runtime/recipe/trust.go
      internal/runtime/recipe/signature.go
      internal/runtime/recipe/signature_test.go
      internal/runtime/artifact/artifact.go
      internal/runtime/artifact/acquire.go
      internal/runtime/artifact/cache.go
      internal/runtime/artifact/artifact_test.go
      internal/runtime/appidentity/recipe_conflict_audit_test.go
      cmd/xnix-runtime-go/artifact_stage_commands.go
      cmd/xnix-runtime-go/artifact_stage_cli_test.go
      cmd/xnix-runtime-go/application_upgrade_impact_commands.go
      cmd/xnix-runtime-go/application_upgrade_impact_cli_test.go
      cmd/xnix-runtime-go/recipe_conflict_audit_commands.go
      cmd/xnix-runtime-go/recipe_conflict_audit_cli_test.go
      cmd/xnix-runtime-go/signed_recipe_verifier_commands.go
      cmd/xnix-runtime-go/signed_recipe_verifier_cli_test.go
    ],
    state_files: %w[
      internal/runtime/artifact/stage.go
      internal/runtime/appidentity/recipe_conflict_audit.go
    ],
    smoke_files: [],
    gate_tokens: {
      "internal/runtime/artifact/acquire.go" => %w[ErrNetworkDisabled NewNetworkSource],
      "internal/runtime/artifact/stage.go" => %w[xnix.runtime.artifact_stage_receipt.v1 go-runtime-local-fixture-artifact-staging artifact-ledger ValidateStageReceipt required artifacts are not staged cache_root_path_exposed fixture_root_path_exposed package_manager_invoked],
      "cmd/xnix-runtime-go/artifact_stage_commands.go" => %w[artifact-stage-record cache-root fixture-root],
      "internal/runtime/appidentity/install_plan.go" => %w[CompatibilityInstallPlanPreviewWithArtifactReceipt artifact_stage_receipt_ready artifact_stage_digest_verified required_artifacts_staged artifact_stage_blocking_reasons artifact-stage-receipt],
      "cmd/xnix-runtime-go/install_plan_commands.go" => %w[artifact-receipt loadArtifactStageReceipt CompatibilityInstallPlanPreviewWithArtifactReceipt],
      "internal/runtime/appidentity/application_upgrade_impact.go" => %w[xnix.runtime.application_upgrade_impact.v1 application-upgrade-impact-preview review-only-application-upgrade-impact candidate-newer same-version candidate-older RecipeWritten RegistryMigrated ArtifactsStaged HostRootModified],
      "cmd/xnix-runtime-go/application_upgrade_impact_commands.go" => %w[application-upgrade-impact-preview candidate-registry candidate-app LoadRecipeFromRegistry],
      "cmd/xnix-runtime-go/application_upgrade_impact_cli_test.go" => %w[TestApplicationUpgradeImpactPreviewCommandRendersCandidateRegistryImpact TestApplicationUpgradeImpactPreviewCommandRejectsDigestDrift recipe digest mismatch],
      "internal/runtime/appidentity/runtime_owner_recipe_trust.go" => %w[ProductionRecipeTrustReady SignedRecipeValidation],
      "internal/runtime/appidentity/recipe_conflict_audit.go" => %w[xnix.runtime.recipe_conflict_audit.v1 recipe-conflict-audit-preview duplicate-app-ids stale-recipe-versions unsupported-capability-claims mismatched-package-source-pins trust-policy-blockers artifact-digest-drift RecipeWritesEnabled RegistryMigrationEnabled ArtifactStagingEnabled PackageManagerInvoked HostRootModified],
      "cmd/xnix-runtime-go/recipe_conflict_audit_commands.go" => %w[recipe-conflict-audit-preview RecipeConflictAuditOptions NewRecipeConflictAuditPreview],
      "cmd/xnix-runtime-go/recipe_conflict_audit_cli_test.go" => %w[TestRecipeConflictAuditPreviewCLI xnix.runtime.recipe_conflict_audit.v1],
      "internal/runtime/recipe/signature.go" => %w[SignedMetadata SignedVerifier SigningPayload ed25519.Verify xnix.runtime.signed_recipe_verification.v1 signed-recipe-verifier-preview fixture-verified ProductionKeyConfigured ProductionTrustReady PrivateKeyLoaded NetworkRequired PackageManagerInvoked RecipeWritten RegistryMigrated BackendLaunchEnabled HostRootModified RecipePathExposed PublicKeyPathExposed SignatureMaterialExposed],
      "cmd/xnix-runtime-go/signed_recipe_verifier_commands.go" => %w[signed-recipe-verifier-preview NewSignedRecipeVerificationPreview public-key metadata],
      "cmd/xnix-runtime-go/signed_recipe_verifier_cli_test.go" => %w[TestSignedRecipeVerifierPreviewCLI TestSignedRecipeVerifierPreviewCLIReportsInvalidSignature TestSignedRecipeVerifierPreviewCLIRejectsMissingArguments]
    },
    summary: "Local recipe and artifact staging persists receipts under a controlled cache root, validates receipt integrity, lets install readiness consume safe receipt diagnostics, previews application upgrade impact before registry migration or activation, audits a registry for trust and digest conflicts, and now verifies offline Ed25519 signed recipe metadata through a fail-closed Go boundary; production key configuration, network fetch, backend launch, recipe writes, and host mutation remain gated."
  },
  {
    id: "environment-lifecycle-state",
    name: "Environment lifecycle state",
    package: "P3",
    mainline_package: "M3",
    contract_files: %w[
      internal/runtime/appidentity/backend_environment.go
      internal/runtime/appidentity/backend_binding.go
      internal/runtime/appidentity/backend_lifecycle.go
      internal/runtime/appidentity/execution_readiness.go
    ],
    fixture_files: %w[
      internal/runtime/environment/environment.go
      internal/runtime/environment/lifecycle.go
      internal/runtime/environment/lifecycle_test.go
      internal/runtime/appidentity/backend_manager.go
      internal/runtime/appidentity/backend_manager_test.go
      internal/runtime/appidentity/backend_adapter_contract.go
      internal/runtime/appidentity/backend_adapter_contract_test.go
      internal/runtime/appidentity/backend_adapter_contract_owner_route_audit.go
      internal/runtime/appidentity/backend_adapter_contract_owner_route_audit_test.go
      internal/runtime/appidentity/backend_adapter_contract_redacted_profile_audit.go
      internal/runtime/appidentity/backend_adapter_contract_redacted_profile_audit_test.go
      internal/runtime/appidentity/backend_lifecycle_test.go
      internal/runtime/appidentity/kde_backend_lifecycle_evidence.go
      internal/runtime/appidentity/kde_backend_lifecycle_evidence_test.go
      internal/runtime/appidentity/state_root_quota_retention_test.go
      internal/runtime/appidentity/compatibility_backend_fallback_test.go
      cmd/xnix-runtime-go/backend_group_cli_test.go
      cmd/xnix-runtime-go/backend_adapter_contract_commands.go
      cmd/xnix-runtime-go/backend_adapter_contract_cli_test.go
      cmd/xnix-runtime-go/backend_adapter_contract_owner_route_audit_cli_test.go
      cmd/xnix-runtime-go/backend_adapter_contract_redacted_profile_audit_cli_test.go
      cmd/xnix-runtime-go/kde_backend_lifecycle_evidence_commands.go
      cmd/xnix-runtime-go/kde_backend_lifecycle_evidence_cli_test.go
      cmd/xnix-runtime-go/state_root_quota_retention_cli_test.go
      cmd/xnix-runtime-go/compatibility_backend_fallback_commands.go
      cmd/xnix-runtime-go/compatibility_backend_fallback_cli_test.go
    ],
    state_files: %w[
      internal/runtime/appidentity/state_root.go
      internal/runtime/appidentity/state_root_quota_retention.go
      internal/runtime/appidentity/compatibility_backend_fallback.go
      internal/runtime/appidentity/kde_backend_lifecycle_evidence.go
    ],
    smoke_files: [],
    gate_tokens: {
      "internal/runtime/environment/lifecycle.go" => %w[repair-required StateBlocked],
      "internal/runtime/appidentity/backend_manager.go" => %w[xnix.runtime.backend_manager.v1 xnix.runtime.backend_manager_record.v1 backend-manager-preview backend-manager-inventory-record go-runtime-state-root-backend-manager wine proton windows-vm BackendDetailsExposedToKDE BackendLaunchEnabled BackendInstallEnabled BackendDownloadEnabled StateRootPathExposed LoadBackendManagerRecord validateBackendManagerRecord backendManagerRecordUnsafe digest\ mismatch unsafe\ enabled\ gates],
      "internal/runtime/appidentity/backend_manager_test.go" => %w[TestBackendManagerPreviewTracksRuntimeBackendsWithoutStartingThem TestBackendManagerPreviewKeepsUserFacingProfilesBackendSafe TestRecordBackendManagerPreviewPersistsStateRootInventory TestBackendManagerRecordRejectsTamperingAndManagedPathSymlink],
      "internal/runtime/appidentity/backend_adapter_contract.go" => %w[xnix.runtime.backend_adapter_contract.v1 backend-adapter-contract-preview compatibility-backend-adapter-noop-contract go-runtime-backend-manager+adapter-noop-boundary GetBackendAdapterContract GetBackendAdapterContractPreview wine proton windows-vm NoopImplementation AdapterInvocationEnabled BackendLaunchEnabled BackendInstallEnabled BackendDownloadEnabled BackendProcessStarted VMProcessStarted CommandMaterialized ExecutablePathResolved RawCommandExposed ProfilePathExposed StateRootPathExposed BackendDetailsExposedToKDE HostRootModified PrivilegedContainerRequired test-only-materialization],
      "internal/runtime/appidentity/backend_adapter_contract_test.go" => %w[TestBackendAdapterContractPreviewDefinesNoopBoundary TestBackendAdapterContractAdaptersStayNoop TestBackendAdapterContractKDEProjectionIsBackendSafe],
      "internal/runtime/appidentity/backend_adapter_contract_owner_route_audit.go" => %w[xnix.runtime.backend_adapter_contract_owner_route_audit.v1 backend-adapter-contract-owner-route-audit-preview NewBackendAdapterContractOwnerRouteAuditPreview GetBackendAdapterContractOwnerRouteAudit redacted-profile-route-smoke-covered redacted-profile-route-smoke-covered-full-contract-fixture-local-production-dbus-blocked redacted-adapter-profile-owner-smoke-coverage owner-smoke-coverage-present route-decision FixtureMatrixConsumesContract OwnerDispatchRoutePresent ProductionDBusMethodPresent FullContractContainsAdapterIDs RedactedProfileRoutePresent OwnerSmokeCoverageReady BackendLaunchEnabled HostRootModified],
      "internal/runtime/appidentity/backend_adapter_contract_owner_route_audit_test.go" => %w[TestBackendAdapterContractOwnerRouteAuditKeepsContractFixtureLocal TestBackendAdapterContractOwnerRouteAuditFailsClosedWithoutSources redacted-profile-route-smoke-covered owner-smoke-coverage-present internal-detail-boundary],
      "internal/runtime/appidentity/backend_adapter_contract_redacted_profile_audit.go" => %w[xnix.runtime.backend_adapter_redacted_profile_audit.v1 backend-adapter-redacted-profile-audit-preview owner-local-redacted-adapter-profile-audit NewBackendAdapterRedactedProfileAuditPreview GetBackendAdapterProfileAudit redacted-profile-route-ready contract-profile-projection-consumed internal-adapter-ids-redacted owner-local-route-candidate full-contract-fixture-local production-dbus-not-claimed no-caller-state-root unsafe-gates-closed OwnerLocalRouteCandidateReady ProductionDBusExposureReady BackendLaunchEnabled HostRootModified],
      "internal/runtime/appidentity/backend_adapter_contract_redacted_profile_audit_test.go" => %w[TestBackendAdapterRedactedProfileAuditPreviewDefinesOwnerRouteSplit TestBackendAdapterRedactedProfileAuditOutputIsRedacted TestBackendAdapterRedactedProfileAuditFailsClosedWithoutOwnerRouteSource automatic performance-priority compatibility-priority],
      "internal/runtime/appidentity/kde_backend_lifecycle_evidence.go" => %w[xnix.runtime.kde_backend_lifecycle_evidence.v1 kde-backend-lifecycle-evidence-record NewKDEBackendLifecycleEvidenceRecord prerequisite-convergence inventory-readback backend-state lifecycle-join execution-readback session-readback backend-boundary unsafe-gates-closed BackendStateJoined BackendKindsExposedToKDE BackendInstallEnabled BackendDownloadEnabled BackendLaunchEnabled BackendProcessStarted VMProcessStarted RawCommandExposed ProfilePathExposed HostRootModified SecretsExposed],
      "internal/runtime/appidentity/kde_backend_lifecycle_evidence_test.go" => %w[TestKDEBackendLifecycleEvidenceJoinsInventoryWithoutStartingProcess TestKDEBackendLifecycleEvidenceRejectsUnsafeBoundary explicit-test-root-only backend-manager],
      "cmd/xnix-runtime-go/kde_backend_lifecycle_evidence_commands.go" => %w[kde-backend-lifecycle-evidence-record runKDEBackendLifecycleEvidenceRecord test-only LoadRecipeFromRegistry NewKDEBackendLifecycleEvidenceRecord],
      "cmd/xnix-runtime-go/kde_backend_lifecycle_evidence_cli_test.go" => %w[TestKDEBackendLifecycleEvidenceRecordCommandJoinsInventory TestKDEBackendLifecycleEvidenceRecordCommandRequiresTestOnlyBoundary backend_kinds_exposed_to_kde backend_process_started host_root_modified secrets_exposed],
      "cmd/xnix-runtime-go/main.go" => %w[backend-adapter-contract-preview backend-adapter-contract-owner-route-audit-preview backend-adapter-redacted-profile-audit-preview backend-manager-preview backend-manager-record runBackendManagerPreview runBackendManagerRecord backend-lifecycle-record runBackendLifecycleRecord],
      "cmd/xnix-runtime-go/backend_adapter_contract_commands.go" => %w[backend-adapter-contract-preview runBackendAdapterContractPreview NewBackendAdapterContractPreview backend-adapter-contract-owner-route-audit-preview runBackendAdapterContractOwnerRouteAuditPreview backend-adapter-redacted-profile-audit-preview runBackendAdapterRedactedProfileAuditPreview],
      "cmd/xnix-runtime-go/backend_adapter_contract_cli_test.go" => %w[TestBackendAdapterContractPreviewCommandRendersNoopBoundary TestBackendAdapterContractPreviewCommandKeepsKDEProjectionBackendSafe xnix.runtime.backend_adapter_contract.v1],
      "cmd/xnix-runtime-go/backend_adapter_contract_owner_route_audit_cli_test.go" => %w[TestBackendAdapterContractOwnerRouteAuditPreviewCommand backend-adapter-contract-owner-route-audit-preview route_decision fixture_matrix_consumes_contract owner_dispatch_route_present production_dbus_method_present],
      "cmd/xnix-runtime-go/backend_adapter_contract_redacted_profile_audit_cli_test.go" => %w[TestBackendAdapterRedactedProfileAuditPreviewCommand backend-adapter-redacted-profile-audit-preview owner_local_route_candidate_ready full_contract_fixture_local production_dbus_exposure_ready],
      "internal/runtime/appidentity/backend_lifecycle.go" => %w[BackendLifecyclePreviewWithStateRoot RecordBackendLifecycleState BackendLifecycleRecord xnix.runtime.backend_lifecycle_record.v1 go-runtime-state-root-backend-lifecycle environment-state-root StateRootBacked StateRootPathExposed SatisfiedGates PendingGates RepairHints BlockReason],
      "internal/runtime/appidentity/backend_lifecycle_test.go" => %w[TestBackendLifecyclePreviewConsumesEnvironmentStateRoot TestRecordBackendLifecycleStatePersistsActionsAndPreview BackendLifecyclePreviewWithStateRoot state-root],
      "internal/runtime/appidentity/state_root_quota_retention.go" => %w[xnix.runtime.state_root_quota_retention.v1 state-root-quota-retention-preview go-runtime-state-root-quota-retention-preview snapshots diagnostics execution-receipts portal-receipts artifact-receipts activation-receipts unknown-records FileDeletionEnabled DirectoriesCreated LogTruncationEnabled ReceiptsRewritten SnapshotDeletionEnabled StateRootPathExposed HostRootModified],
      "internal/runtime/appidentity/state_root_quota_retention_test.go" => %w[TestStateRootQuotaRetentionPreviewMissingRootDoesNotCreateDirectories TestStateRootQuotaRetentionPreviewOverQuotaAndMalformedUnknownEvidence TestStateRootQuotaRetentionPreviewActiveSessionBlocksExecutionReceipts TestStateRootQuotaRetentionPreviewRetentionExemptEvidenceIsKept],
      "cmd/xnix-runtime-go/state_root_quota_retention_commands.go" => %w[state-root-quota-retention-preview quota-bytes active-session StateRootQuotaRetentionOptions],
      "cmd/xnix-runtime-go/state_root_quota_retention_cli_test.go" => %w[TestStateRootQuotaRetentionPreviewCLI TestStateRootQuotaRetentionPreviewCLIMissingRootIsReadOnly xnix.runtime.state_root_quota_retention.v1],
      "cmd/xnix-runtime-go/backend_group_cli_test.go" => %w[backend-lifecycle-record TestBackendLifecycleRecordCommandPersistsStateTransitions --state-root state_root_backed environment-state-root],
      "internal/runtime/appidentity/compatibility_backend_fallback.go" => %w[xnix.runtime.compatibility_backend_fallback.v1 compatibility-backend-fallback-preview local-compatibility isolated-compatibility recipe-requests-isolation blocked-missing-evidence SelectionPersisted EngineInstallEnabled BackendLaunchEnabled VMStartEnabled HostRootModified],
      "cmd/xnix-runtime-go/compatibility_backend_fallback_commands.go" => %w[compatibility-backend-fallback-preview CompatibilityBackendFallbackPreview],
      "cmd/xnix-runtime-go/compatibility_backend_fallback_cli_test.go" => %w[TestCompatibilityBackendFallbackPreviewCLI xnix.runtime.compatibility_backend_fallback.v1],
      "internal/runtime/appidentity/execution_readiness.go" => %w[LaunchAllowed BackendDetailsExposed]
    },
    summary: "Runtime environment lifecycle logic exists with state-root integration; backend lifecycle records can be persisted and advanced through Go Runtime CLI actions under an explicit state root; state-root quota and retention previews can inspect snapshots, diagnostics, execution receipts, Portal receipts, artifact receipts, activation receipts, and unknown records in dry-run mode without deleting or rewriting anything; the Go backend manager tracks Wine, Proton, and Windows VM backends as internal Runtime-managed inventory, can persist that inventory under an explicit state root while hiding root paths, and now defines no-op adapter contracts for those backends while exposing only safe compatibility profiles to KDE and keeping real backend starts disabled."
  },
  {
    id: "portal-snapshot-control-plane",
    name: "Portal and snapshot control plane",
    package: "P4",
    mainline_package: "M4",
    contract_files: %w[
      internal/runtime/appidentity/portal_access_policy.go
      internal/runtime/appidentity/snapshot_plan.go
    ],
    fixture_files: %w[
      internal/runtime/portal/broker.go
      internal/runtime/portal/request.go
      internal/runtime/portal/ledger.go
      internal/runtime/portal/read.go
      internal/runtime/portal/broker_test.go
      internal/runtime/portal/request_test.go
      internal/runtime/portal/ledger_test.go
      internal/runtime/portal/read_test.go
      internal/runtime/appidentity/permission_evidence_audit_test.go
      internal/runtime/appidentity/portal_permission_renewal_test.go
      internal/runtime/appidentity/snapshot_restore_candidates_test.go
      cmd/xnix-runtime-go/runtime_safety_cli_test.go
      cmd/xnix-runtime-go/permission_evidence_audit_commands.go
      cmd/xnix-runtime-go/permission_evidence_audit_cli_test.go
      cmd/xnix-runtime-go/portal_permission_renewal_commands.go
      cmd/xnix-runtime-go/portal_permission_renewal_cli_test.go
      cmd/xnix-runtime-go/snapshot_restore_candidates_commands.go
      cmd/xnix-runtime-go/snapshot_restore_candidates_cli_test.go
    ],
    state_files: %w[
      internal/runtime/portal/ledger.go
      internal/runtime/appidentity/permission_evidence_audit.go
      internal/runtime/appidentity/portal_permission_renewal.go
      internal/runtime/appidentity/snapshot_restore_candidates.go
      internal/runtime/snapshot/store.go
      internal/runtime/snapshot/store_test.go
    ],
    smoke_files: [],
    gate_tokens: {
      "internal/runtime/portal/broker.go" => %w[FakeBroker denied cancelled failed],
      "internal/runtime/portal/ledger.go" => %w[xnix.runtime.portal_request_record.v1 portal-permission-request-record go-runtime-state-root-portal-broker RealPortalCallEnabled ExecutionApproved StateRootPathExposed HostRootModified],
      "internal/runtime/portal/read.go" => %w[ReadRequests portal-requests malformed],
      "internal/runtime/appidentity/permission_evidence_audit.go" => %w[xnix.runtime.permission_evidence_audit.v1 permission-evidence-audit-preview GetPermissionEvidenceAuditPreview consistent setting-only receipt-only expired denied missing-review blocked-by-policy unsupported RealPortalCallEnabled PermissionGrantEnabled PermissionRevokeEnabled ReceiptWriteEnabled SettingsPersisted StateRootPathExposed HostRootModified],
      "cmd/xnix-runtime-go/permission_evidence_audit_commands.go" => %w[permission-evidence-audit-preview ReadRequests Summarize PermissionEvidenceAuditPreview],
      "cmd/xnix-runtime-go/permission_evidence_audit_cli_test.go" => %w[TestPermissionEvidenceAuditPreviewCLIWithReceipts TestPermissionEvidenceAuditPreviewCLIMalformedLedgerIsSurfaced state_root_path_exposed],
      "internal/runtime/appidentity/portal_permission_renewal.go" => %w[xnix.runtime.portal_permission_renewal.v1 portal-permission-renewal-preview GetPortalPermissionRenewalPreview current needs-review expiring-soon denied revoked missing-receipt blocked-by-policy RealPortalCallEnabled PermissionGrantEnabled PermissionRevokeEnabled ReceiptWriteEnabled SettingsPersisted ExecutionApproved HostRootModified],
      "cmd/xnix-runtime-go/portal_permission_renewal_commands.go" => %w[portal-permission-renewal-preview ReadRequests PortalPermissionRenewalPreview StateCancelled expiring],
      "cmd/xnix-runtime-go/portal_permission_renewal_cli_test.go" => %w[TestPortalPermissionRenewalPreviewCLI TestPortalPermissionRenewalPreviewCLIMalformedLedgerIsSurfaced xnix.runtime.portal_permission_renewal.v1],
      "cmd/xnix-runtime-go/main.go" => %w[portal-request-record runPortalRequestRecord permission-evidence-audit-preview runPermissionEvidenceAuditPreview portal-permission-renewal-preview runPortalPermissionRenewalPreview],
      "cmd/xnix-runtime-go/runtime_safety_cli_test.go" => %w[TestPortalRequestRecordCommandPersistsPermissionFlow portal-request-record permission_granted execution_approved],
      "internal/runtime/snapshot/store.go" => %w[Rollback Verify OpenReadOnly],
      "internal/runtime/appidentity/snapshot_restore_candidates.go" => %w[xnix.runtime.snapshot_restore_candidates.v1 snapshot-restore-candidates-preview latest-good latest-tested last-known-running digest-mismatch active-session RestoreExecuted SnapshotDeletionEnabled FileContentRead SessionTerminated StateRootPathExposed HostRootModified],
      "cmd/xnix-runtime-go/snapshot_restore_candidates_commands.go" => %w[snapshot-restore-candidates-preview OpenReadOnly NewSnapshotRestoreCandidatesPreview active-session],
      "cmd/xnix-runtime-go/snapshot_restore_candidates_cli_test.go" => %w[TestSnapshotRestoreCandidatesPreviewCLI TestSnapshotRestoreCandidatesPreviewCLIEmptyStateRoot state_root_path_exposed]
    },
    summary: "Fake Portal records can persist state-root permission request receipts with grant, deny, cancel, expire, and complete states; a read-only permission evidence audit joins Portal access policy, recorded receipts, permission review, resource bridge, and execution preflight into consistency states without changing any permission; a Portal permission renewal preview explains current, needs-review, expiring-soon, denied, revoked, missing-receipt, and policy-blocked cases without real Portal transport, grants, revocation, receipt writes, settings persistence, execution approval, or host mutation; content-addressed snapshots exist under controlled roots and can be ranked read-only as restore candidates with latest-good, latest-tested, last-known-running, blocked, and unknown labels; real Portal calls, execution approval, snapshot restore, snapshot deletion, and system snapshots remain gated."
  },
  {
    id: "kde-activation-shell-materialization",
    name: "KDE activation and shell materialization",
    package: "P5",
    mainline_package: "M5",
    contract_files: %w[
      internal/runtime/appidentity/desktop_activation_bundle.go
      internal/runtime/appidentity/desktop_activation_status.go
      kde/dolphin/servicemenus/xnix-open-with-compatibility.desktop
      kde/plasmoids/org.xnix.compatibilitycenter/metadata.json
    ],
    fixture_files: %w[
      internal/runtime/activation/stage.go
      internal/runtime/activation/stage_test.go
      internal/runtime/appidentity/execution_session_record_evidence.go
      internal/runtime/appidentity/desktop_safety_policy.go
      internal/runtime/appidentity/desktop_safety_policy_test.go
      internal/runtime/appidentity/runtime_policy_explanation_cards.go
      internal/runtime/appidentity/runtime_policy_explanation_cards_test.go
      internal/runtime/appidentity/settings_profile_migration_test.go
      internal/runtime/appidentity/window_identity_routes_test.go
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
      internal/runtime/appidentity/desktop_deactivation_dry_run_test.go
      cmd/xnix-runtime-go/desktop_safety_policy_commands.go
      cmd/xnix-runtime-go/desktop_safety_policy_cli_test.go
      cmd/xnix-runtime-go/runtime_policy_explanation_cards_commands.go
      cmd/xnix-runtime-go/runtime_policy_explanation_cards_cli_test.go
      cmd/xnix-runtime-go/settings_profile_migration_commands.go
      cmd/xnix-runtime-go/settings_profile_migration_cli_test.go
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
      cmd/xnix-runtime-go/desktop_deactivation_dry_run_commands.go
      cmd/xnix-runtime-go/desktop_deactivation_dry_run_cli_test.go
      scripts/install_runtime_activation.rb
      scripts/runtime_activation_smoke.rb
    ],
    state_files: %w[
      internal/runtime/activation/stage.go
      internal/runtime/appidentity/kde_search_visibility.go
      internal/runtime/appidentity/desktop_deactivation_dry_run.go
      internal/runtime/appidentity/settings_profile_migration.go
      internal/runtime/appidentity/kde_fake_execution_evidence.go
      internal/runtime/appidentity/kde_fake_portal_evidence.go
      internal/runtime/appidentity/kde_snapshot_diagnostics_evidence.go
    ],
    smoke_files: %w[
      scripts/runtime_activation_smoke.rb
      test/test_runtime_activation_smoke_script.rb
      scripts/kde_first_presence_smoke.rb
      test/test_kde_first_presence_smoke_script.rb
    ],
    gate_tokens: {
      "internal/runtime/activation/stage.go" => ["refusing to stage", "host_root_modified", "RollbackReceiptWritten"],
      "internal/runtime/appidentity/desktop_activation_transaction.go" => %w[TransactionCommitted WriteGate ReceiptEvidence CommitReceiptPlanned RollbackReceiptPlanned CommitAvailable RollbackAvailable],
      "internal/runtime/appidentity/kde_search_visibility.go" => %w[xnix.runtime.kde_search_visibility.v1 kde-search-visibility-plan-preview receipt-required no-mime-association application-hidden DesktopFilesWritten MIMEDefaultsWritten KDECacheRefreshed HostFilesIndexed SearchIndexPersisted HostRootModified],
      "cmd/xnix-runtime-go/kde_search_visibility_commands.go" => %w[kde-search-visibility-plan-preview KDESearchVisibilityPlanPreview activation-root],
      "cmd/xnix-runtime-go/kde_search_visibility_cli_test.go" => %w[TestKDESearchVisibilityPlanPreviewCLI xnix.runtime.kde_search_visibility.v1],
      "internal/runtime/appidentity/kde_notification_digest.go" => %w[xnix.runtime.kde_notification_digest.v1 kde-notification-digest-preview review-only-runtime-event-digest GetKDENotificationDigestPreview needs-review blocked-action permission-attention diagnostic-issue snapshot-warning readiness-change NotificationsSent LiveTrayBridgeEnabled RequestObjectsCreated PermissionGrantEnabled BackendLaunchEnabled HostRootModified],
      "cmd/xnix-runtime-go/kde_notification_digest_commands.go" => %w[kde-notification-digest-preview repeatedNotificationDigestEvents KDENotificationDigestPreview],
      "cmd/xnix-runtime-go/kde_notification_digest_cli_test.go" => %w[TestKDENotificationDigestPreviewCLI TestKDENotificationDigestPreviewCLIRejectsInvalidArguments xnix.runtime.kde_notification_digest.v1 notifications_sent live_tray_bridge_enabled backend_launch_enabled host_root_modified],
      "internal/runtime/appidentity/kde_offline_application_identity.go" => %w[xnix.runtime.kde_offline_application_identity.v1 kde-offline-application-identity-preview GetKDEOfflineApplicationIdentity GetKDEOfflineApplicationIdentityPreview DesktopEntry MIMEAssociations KRunner TaskManager KWin Tray Notification Settings CompatibilityCenter CrossSurfaceIdentityConsistent NotificationIDNamespace DesktopFilesWritten MIMEDefaultsWritten KRunnerIndexPersisted TaskManagerEntryActive KWinRuleApplied LiveTrayBridgeEnabled TrayBridgePersisted NotificationSent NotificationDeliveryEnabled NotificationActionsEnabled SettingsPersisted SettingsPersistenceEnabled CompatibilityCenterPersisted CompatibilityCenterActionsEnabled LaunchEnabled ExecutionStarted BackendProcessStarted HostRootModified],
      "cmd/xnix-runtime-go/kde_offline_application_identity_commands.go" => %w[kde-offline-application-identity-preview runKDEOfflineApplicationIdentityPreview parseKDEOfflineApplicationIdentityPreviewSource LoadRecipeFromRegistry],
      "cmd/xnix-runtime-go/kde_offline_application_identity_cli_test.go" => %w[TestKDEOfflineApplicationIdentityPreviewCommandUsesSampleFixture TestKDEOfflineApplicationIdentityPreviewCommandRejectsIncompleteArguments xnix.runtime.kde_offline_application_identity.v1 cross_surface_identity_consistent launch_enabled execution_started host_root_modified],
      "internal/runtime/appidentity/kde_fake_execution_evidence.go" => %w[xnix.runtime.kde_fake_execution_evidence.v1 kde-fake-execution-evidence-record test-only explicit-test-root-only NewKDEFakeExecutionEvidenceRecord prepareFakeExecutionLifecycle execution-ledger execution-session-record StateRootWritesEnabled StateRootPathExposed FakeExecutionRecorded LaunchAllowed LaunchEnabled ExecutionStarted BackendProcessStarted RealPortalCallEnabled ProductionBusOwnership HostRootModified],
      "internal/runtime/appidentity/kde_fake_execution_evidence_test.go" => %w[TestKDEFakeExecutionEvidencePersistsAndReadsBackControlledRecords TestKDEFakeExecutionEvidenceRejectsUnsafeInputs TestKDEFakeExecutionEvidenceRejectsManagedPathSymlink staged blocked explicit-test-root-only],
      "cmd/xnix-runtime-go/kde_fake_execution_evidence_commands.go" => %w[kde-fake-execution-evidence-record runKDEFakeExecutionEvidenceRecord test-only LoadRecipeFromRegistry NewKDEFakeExecutionEvidenceRecord],
      "cmd/xnix-runtime-go/kde_fake_execution_evidence_cli_test.go" => %w[TestKDEFakeExecutionEvidenceRecordCommandWritesControlledEvidence TestKDEFakeExecutionEvidenceRecordCommandRequiresExplicitTestBoundary all_checks_passed state_root_writes_enabled launch_enabled backend_process_started host_root_modified],
      "internal/runtime/appidentity/kde_fake_portal_evidence.go" => %w[xnix.runtime.kde_fake_portal_evidence.v1 kde-fake-portal-evidence-record NewKDEFakePortalEvidenceRecord ensureCompletedFakePortalReceipt reviewedFakeExecutionTransaction portal-gate-delta portal-policy-review snapshot-baseline PortalEvidenceRecorded PortalGateChangedOnly RealPortalCallEnabled HostPermissionChanged ExecutionApproved LaunchEnabled BackendProcessStarted HostRootModified],
      "internal/runtime/appidentity/kde_fake_portal_evidence_test.go" => %w[TestKDEFakePortalEvidenceChangesOnlyPortalGate TestKDEFakePortalEvidenceRejectsManagedPathSymlink pending-user-mediation granted completed explicit-test-root-only],
      "cmd/xnix-runtime-go/kde_fake_portal_evidence_commands.go" => %w[kde-fake-portal-evidence-record runKDEFakePortalEvidenceRecord test-only LoadRecipeFromRegistry NewKDEFakePortalEvidenceRecord],
      "cmd/xnix-runtime-go/kde_fake_portal_evidence_cli_test.go" => %w[TestKDEFakePortalEvidenceRecordCommandJoinsPortalReceipt TestKDEFakePortalEvidenceRecordCommandRequiresTestOnlyBoundary portal_gate_changed_only portal_evidence_recorded real_portal_call_enabled host_permission_changed execution_approved],
      "internal/runtime/appidentity/kde_snapshot_diagnostics_evidence.go" => %w[xnix.runtime.kde_snapshot_diagnostics_evidence.v1 kde-snapshot-diagnostics-evidence-record NewKDESnapshotDiagnosticsEvidenceRecord stability-diagnostics-v0.2.315 stability-baseline-v0.2.315 recordStabilityDiagnostic ensureStabilitySnapshot portal-evidence diagnostic-readback diagnostic-redaction snapshot-baseline lifecycle-ready execution-readiness session-readback unsafe-gates-closed SnapshotCreationEnabled SnapshotRestoreEnabled DiagnosticExecutionEnabled AIProviderCallEnabled RepairExecutionEnabled LaunchEnabled BackendProcessStarted HostRootModified FileContentsExposed],
      "internal/runtime/appidentity/kde_snapshot_diagnostics_evidence_test.go" => %w[TestKDESnapshotDiagnosticsEvidenceConvergesControlledState TestKDESnapshotDiagnosticsEvidenceRejectsUnsafeBoundaries ready blocked explicit-test-root-only .xnix-snapshots diagnostics-ledger test-results],
      "cmd/xnix-runtime-go/kde_snapshot_diagnostics_evidence_commands.go" => %w[kde-snapshot-diagnostics-evidence-record runKDESnapshotDiagnosticsEvidenceRecord test-only LoadRecipeFromRegistry NewKDESnapshotDiagnosticsEvidenceRecord],
      "cmd/xnix-runtime-go/kde_snapshot_diagnostics_evidence_cli_test.go" => %w[TestKDESnapshotDiagnosticsEvidenceRecordCommandConvergesEvidence TestKDESnapshotDiagnosticsEvidenceRecordCommandRequiresTestOnlyBoundary snapshot_restore_enabled diagnostic_execution_enabled ai_provider_call_enabled repair_execution_enabled launch_enabled backend_process_started host_root_modified file_contents_exposed],
      "internal/runtime/owner/dispatch.go" => %w[GetKDENotificationDigestPreview kde-notification-digest-preview ownerNotificationDigestEvents KDENotificationDigestPreview GetSignedRecipeVerificationPreview signed-recipe-verifier-preview NewRegistrySignedRecipeVerificationPreview GetRestrictedProductSmokePacketPreview restricted-product-smoke-packet-preview PrepareRestrictedProductSmokePacket GetKDEOfflineApplicationIdentityPreview kde-offline-application-identity-preview NewKDEOfflineApplicationIdentityPreview GetBackendAdapterProfileAudit backend-adapter-redacted-profile-audit-preview NewBackendAdapterRedactedProfileAuditPreview GetKDETestLaunchMaterializationReceiptLookupPreview kde-test-launch-materialization-receipt-lookup-preview ResolveKDETestLaunchMaterializationReceipt GetKDETestLaunchMaterializationFanOut kde-test-launch-materialization-fanout-owner-route-preview NewKDETestLaunchMaterializationFanOutOwnerRoutePreview],
      "internal/runtime/owner/service_test.go" => %w[TestServiceCallServesOwnerLocalNotificationDigest GetKDENotificationDigestPreview TestServiceCallServesOwnerLocalSignedRecipeVerification GetSignedRecipeVerificationPreview TestServiceCallServesOwnerLocalRestrictedSmokePacket GetRestrictedProductSmokePacketPreview TestServiceCallServesOwnerLocalOfflineKDEIdentity GetKDEOfflineApplicationIdentityPreview TestServiceCallServesOwnerLocalRedactedBackendAdapterProfileAudit GetBackendAdapterProfileAudit TestServiceCallServesOwnerLocalKDETestLaunchMaterializationReceiptLookup GetKDETestLaunchMaterializationReceiptLookupPreview TestServiceCallServesOwnerLocalKDETestLaunchMaterializationFanOut GetKDETestLaunchMaterializationFanOut read-dispatch],
      "cmd/xnix-runtime-owner/main_test.go" => %w[TestRuntimeOwnerCommandRendersNotificationDigestOwnerLocalReadDispatch GetKDENotificationDigestPreview kde-notification-digest-preview TestRuntimeOwnerCommandRendersSignedRecipeOwnerLocalReadDispatch GetSignedRecipeVerificationPreview signed-recipe-verifier-preview TestRuntimeOwnerCommandRendersRestrictedSmokeOwnerLocalReadDispatch GetRestrictedProductSmokePacketPreview restricted-product-smoke-packet-preview TestRuntimeOwnerCommandRendersOfflineKDEIdentityOwnerLocalReadDispatch GetKDEOfflineApplicationIdentityPreview kde-offline-application-identity-preview TestRuntimeOwnerCommandRendersRedactedAdapterProfileOwnerLocalReadDispatch GetBackendAdapterProfileAudit backend-adapter-redacted-profile-audit-preview TestRuntimeOwnerCommandRendersMaterializationReceiptLookupReadDispatch GetKDETestLaunchMaterializationReceiptLookupPreview kde-test-launch-materialization-receipt-lookup-preview TestRuntimeOwnerCommandRendersMaterializationFanOutReadDispatch GetKDETestLaunchMaterializationFanOut kde-test-launch-materialization-fanout-owner-route-preview],
      "internal/runtime/recipe/signature.go" => %w[NewRegistrySignedRecipeVerificationPreview owner-registry-signed-recipe-verifier-evidence production-signature-required production-key-configuration-missing],
      "internal/runtime/appidentity/desktop_deactivation_dry_run.go" => %w[xnix.runtime.desktop_deactivation_dry_run.v1 desktop-deactivation-dry-run-preview missing-activation-receipt installed-file-digest-mismatch unknown-file-owner shared-mime-association active-session-evidence FileDeletionEnabled MIMEDefaultsWritten KDECacheRefreshed ReceiptsRewritten SessionTerminated TargetPathExposed HostRootModified],
      "cmd/xnix-runtime-go/desktop_deactivation_dry_run_commands.go" => %w[desktop-deactivation-dry-run-preview DesktopDeactivationDryRunPreview active-session shared-mime],
      "cmd/xnix-runtime-go/desktop_deactivation_dry_run_cli_test.go" => %w[TestDesktopDeactivationDryRunPreviewCLIReceiptBackedRemovable TestDesktopDeactivationDryRunPreviewCLIActiveSessionAndSharedMIMEBlock xnix.runtime.desktop_deactivation_dry_run.v1],
      "internal/runtime/appidentity/desktop_activation_status.go" => %w[DesktopActivationStatusPreviewWithReceipt receipt_backed receipt_evidence RollbackAvailable],
      "internal/runtime/appidentity/kde_center_page.go" => %w[KDECenterPageOptions ActivationRoot ReceiptBacked ReceiptEvidenceState ExecutionSessionRoot ExecutionSessionBacked ExecutionSessionPath TaskManagerSessionState KWinSessionState],
      "internal/runtime/appidentity/execution_session_record_evidence.go" => %w[ExecutionSessionRecordEvidence ExecutionSessionFanOutEvidence xnix.runtime.execution_session_record.v1 xnix.runtime.session_fanout.v1 execution-session-fanout-evidence GetExecutionSessionFanOutEvidence SafeForKDE StateRootPathExposed SessionActive LiveTrayBridgeEnabled],
      "internal/runtime/appidentity/identity.go" => %w[DesktopEntryOptions DesktopIconOptions MIMEAppsOptions KRunnerQueryOptions ActivationReceiptBacked ReceiptBackedMatchCount FileOpenOptions TrayStatusOptions NotificationOptions SettingsOptions ActivationReceiptRoot ExecutionSessionBacked ExecutionSessionPath],
      "internal/runtime/appidentity/kde_shell_surface.go" => %w[KDEApplicationSurfaceOptions ActivationReceiptBacked ActivationReceiptPath],
      "internal/runtime/appidentity/window_identity_routes.go" => %w[TaskManagerIdentityOptions KWinWindowRuleOptions ActivationReceiptBacked ActivationReceiptPath ExecutionSessionBacked ExecutionSessionPath],
      "internal/runtime/appidentity/desktop_safety_policy.go" => %w[xnix.runtime.desktop_safety_policy.v1 kde-first-user-facing-safety-policy forbidden_user_terms settings_field_ids BackendTerminologyHidden WriteMethodsEnabled RealPortalTransport AIProviderCallEnabled],
      "cmd/xnix-runtime-go/desktop_safety_policy_commands.go" => %w[desktop-safety-policy-preview NewDesktopSafetyPolicyPreview],
      "cmd/xnix-runtime-go/desktop_safety_policy_cli_test.go" => %w[TestDesktopSafetyPolicyPreviewCommand xnix.runtime.desktop_safety_policy.v1 forbidden_user_terms settings_field_ids],
      "internal/runtime/appidentity/runtime_policy_explanation_cards.go" => %w[xnix.runtime.policy_explanation_cards.v1 runtime-policy-explanation-cards-preview kde-runtime-policy-explanation-card-deck GetRuntimePolicyExplanationCards GetRuntimePolicyExplanationCardsPreview install launch execution portal-permission snapshot diagnostics repair settings desktop-activation backend-readiness unsupported-production-route RuntimeOwned GoRuntimeBacked KDEPolicyOwner CardsPersisted ActionEnablementChanged RequestObjectsCreated PermissionGrantsCreated SettingsPersisted AIProviderCalled BackendProcessStarted HostRootModified RawCommandExposed BackendDetailsExposed],
      "cmd/xnix-runtime-go/runtime_policy_explanation_cards_commands.go" => %w[runtime-policy-explanation-cards-preview RuntimePolicyExplanationCardsPreview runtime-root portal-operation snapshot-reason],
      "cmd/xnix-runtime-go/runtime_policy_explanation_cards_cli_test.go" => %w[TestRuntimePolicyExplanationCardsPreviewCommand xnix.runtime.policy_explanation_cards.v1 kde-runtime-policy-explanation-card-deck cards_persisted action_enablement_changed request_objects_created backend_details_exposed],
      "internal/runtime/appidentity/settings_profile_migration.go" => %w[xnix.runtime.settings_profile_migration.v1 settings-profile-migration-preview GetCompatibilitySettingsProfileMigration GetCompatibilitySettingsProfileMigrationPreview old-schema current-schema future-schema run-mode.mode resource-access.documents diagnostics.privacy SettingsPersisted ResourceGrantEnabled RealPortalCallEnabled BackendLaunchEnabled HostRootModified],
      "cmd/xnix-runtime-go/settings_profile_migration_commands.go" => %w[settings-profile-migration-preview SettingsProfileMigrationPreview from-schema to-schema blocked-setting],
      "cmd/xnix-runtime-go/settings_profile_migration_cli_test.go" => %w[TestSettingsProfileMigrationPreviewCLI xnix.runtime.settings_profile_migration.v1 settings_persisted resource_grant_enabled real_portal_call_enabled backend_launch_enabled host_root_modified],
      "scripts/kde_first_presence_smoke.rb" => %w[xnix.kde_first_presence_smoke.v1 kde-first-presence-smoke --format preview_commands route_baseline entrypoint_count settings_field_ids forbidden_user_terms desktop-safety-policy-preview runtime-policy-explanation-cards-preview xnix.runtime.desktop_safety_policy.v1 xnix.runtime.policy_explanation_cards.v1 assert_desktop_safety_policy assert_runtime_policy_explanation_cards safety_false_keys prefix bottle backend_launch_enabled host_root_modified privileged_container_required real_portal_transport_enabled ai_provider_call_enabled],
      "test/test_kde_first_presence_smoke_script.rb" => %w[kde-first-presence-smoke xnix.kde_first_presence_smoke.v1 JSON.pretty_generate render_markdown settings_field_ids forbidden_user_terms desktop-safety-policy-preview xnix.runtime.desktop_safety_policy.v1 safety_false_keys]
    },
    summary: "KDE activation can stage desktop files, MIME data, service menus, manifests, and receipts under explicit roots; KDE task-manager, KWin, tray status, and Compatibility Center page previews can consume durable execution session evidence while staying blocked by Runtime gates; KDE status, KDE application surface previews, Compatibility Center pages, desktop entry previews, KRunner query previews, Dolphin file-manager previews, desktop icon previews, MIME association previews, notification previews, notification digest previews, settings previews, settings profile migration previews, task-manager identity previews, KWin window-rule previews, and Runtime policy explanation cards can consume shared Runtime evidence and present KDE-safe blocked, review-only, missing-evidence, and not-yet-implemented explanations; and the KDE-first presence smoke now emits text, JSON, and Markdown evidence that one digest-verified recipe appears across all seven KDE entry points while execution, backend launch, real Portal transport, network, privileged containers, AI provider calls, card persistence, request creation, permission grants, settings persistence, and host-root mutation remain disabled."
  },
  {
    id: "execution-transaction-ledger",
    name: "Execution transaction ledger",
    package: "P6",
    mainline_package: "M6",
    contract_files: %w[
      internal/runtime/appidentity/launch_intent.go
      internal/runtime/appidentity/execution_request.go
      internal/runtime/appidentity/execution_preflight.go
      internal/runtime/appidentity/execution_transaction.go
      internal/runtime/appidentity/execution_session_status.go
      internal/runtime/appidentity/application_readiness.go
    ],
    fixture_files: %w[
      internal/runtime/execution/execution.go
      internal/runtime/execution/execution_test.go
      internal/runtime/execution/ledger_test.go
      internal/runtime/execution/session_test.go
      internal/runtime/execution/authorization.go
      internal/runtime/execution/authorization_test.go
      internal/runtime/execution/restricted_preflight.go
      internal/runtime/execution/restricted_preflight_test.go
      internal/runtime/execution/restricted_materialization.go
      internal/runtime/execution/restricted_materialization_test.go
      internal/runtime/appidentity/kde_restricted_launch_authorization.go
      internal/runtime/appidentity/kde_restricted_launch_authorization_test.go
      internal/runtime/appidentity/kde_restricted_launch_preflight.go
      internal/runtime/appidentity/kde_restricted_launch_preflight_test.go
      internal/runtime/appidentity/kde_test_launch_materialization.go
      internal/runtime/appidentity/kde_test_launch_materialization_test.go
      internal/runtime/appidentity/kde_test_launch_materialization_fanout.go
      internal/runtime/appidentity/kde_test_launch_materialization_fanout_test.go
      internal/runtime/appidentity/kde_test_launch_materialization_owner_route_audit.go
      internal/runtime/appidentity/kde_test_launch_materialization_owner_route_audit_test.go
      internal/runtime/appidentity/application_readiness_test.go
      cmd/xnix-runtime-go/application_readiness_commands.go
      cmd/xnix-runtime-go/application_readiness_cli_test.go
      cmd/xnix-runtime-go/execution_ledger_commands.go
      cmd/xnix-runtime-go/execution_ledger_cli_test.go
      cmd/xnix-runtime-go/kde_restricted_launch_authorization_commands.go
      cmd/xnix-runtime-go/kde_restricted_launch_authorization_cli_test.go
      cmd/xnix-runtime-go/kde_restricted_launch_preflight_commands.go
      cmd/xnix-runtime-go/kde_restricted_launch_preflight_cli_test.go
      cmd/xnix-runtime-go/kde_test_launch_materialization_commands.go
      cmd/xnix-runtime-go/kde_test_launch_materialization_cli_test.go
      cmd/xnix-runtime-go/kde_test_launch_materialization_fanout_commands.go
      cmd/xnix-runtime-go/kde_test_launch_materialization_fanout_cli_test.go
      cmd/xnix-runtime-go/kde_test_launch_materialization_owner_route_audit_cli_test.go
    ],
    state_files: %w[
      internal/runtime/execution/ledger.go
      internal/runtime/execution/session.go
      internal/runtime/execution/authorization.go
      internal/runtime/execution/restricted_preflight.go
      internal/runtime/execution/restricted_materialization.go
      internal/runtime/appidentity/kde_restricted_launch_authorization.go
      internal/runtime/appidentity/kde_restricted_launch_preflight.go
      internal/runtime/appidentity/kde_test_launch_materialization.go
      internal/runtime/appidentity/kde_test_launch_materialization_fanout.go
      internal/runtime/appidentity/kde_test_launch_materialization_owner_route_audit.go
    ],
    smoke_files: [],
    gate_tokens: {
      "internal/runtime/execution/execution.go" => %w[blocked runtimeWriteGate],
      "internal/runtime/execution/ledger.go" => %w[execution-transaction-ledger-record go-runtime-state-root-execution-ledger state_root_path_exposed portal_permission_receipt_relative_paths portal_permission_receipt_consumed] + ["execution ledger record digest mismatch", "execution ledger record has unsafe enabled gates"],
      "internal/runtime/execution/ledger_test.go" => %w[TestLedgerLoadRejectsTamperedReceipt] + ["tampered execution receipt"],
      "internal/runtime/execution/session.go" => %w[xnix.runtime.execution_session_record.v1 execution-session-status-record go-runtime-state-root-execution-session LoadSession StatusPersisted SessionActive TaskManagerEntryActive LiveTrayBridgeEnabled] + ["execution session record digest mismatch", "execution session record has unsafe enabled gates"],
      "internal/runtime/execution/session_test.go" => %w[TestLoadSessionRejectsTamperedReceipt] + ["tampered session receipt"],
      "internal/runtime/execution/authorization.go" => %w[xnix.runtime.restricted_launch_authorization.v1 restricted-launch-authorization-receipt go-runtime-state-root-restricted-launch-authorization restricted-test-preparation authorize-restricted-test-preparation authorized-preparation-only PreparationAuthorized LaunchAuthorized ProcessStartAuthorized ProductionTrustSatisfied RuntimeWriteGateEnabled ArtifactAcquisitionEnabled BackendLaunchEnabled BackendProcessStarted HostRootModified],
      "internal/runtime/execution/authorization_test.go" => %w[TestRestrictedAuthorizationStorePersistsPreparationOnlyReceipt TestRestrictedAuthorizationStoreRejectsInvalidDirectiveTamperingAndSymlink],
      "internal/runtime/execution/restricted_preflight.go" => %w[xnix.runtime.restricted_launch_preflight.v1 restricted-launch-preflight-packet go-runtime-state-root-restricted-launch-preflight blocked recipe-trust runtime-write-gate ReadyForPacketAssembly ProductImageReady LaunchPreflightPassed LaunchAuthorized ProcessStartAuthorized CommandMaterialized ExecutablePathResolved BackendSelectedForLaunch BackendLaunchEnabled BackendProcessStarted HostRootModified],
      "internal/runtime/execution/restricted_preflight_test.go" => %w[TestRestrictedPreflightStorePersistsFailClosedPacket TestRestrictedPreflightStoreRejectsTamperedPacket],
      "internal/runtime/execution/restricted_materialization.go" => %w[xnix.runtime.restricted_launch_materialization.v1 restricted-launch-materialization-plan go-runtime-state-root-restricted-launch-materialization blocked-plan-materialized test-only-review-plan PlanMaterialized CommandMaterialized ExecutablePathResolved BackendSelectedForLaunch BackendLaunchEnabled BackendProcessStarted RawExecutableExposed HostRootModified],
      "internal/runtime/execution/restricted_materialization_test.go" => %w[TestRestrictedMaterializationStorePersistsPlanOnly TestRestrictedMaterializationStoreRejectsTamperedPlan],
      "internal/runtime/appidentity/kde_restricted_launch_authorization.go" => %w[xnix.runtime.kde_restricted_launch_authorization.v1 kde-restricted-launch-authorization-record NewKDERestrictedLaunchAuthorizationRecord prerequisite-convergence explicit-test-boundary authorization-readback trust-independent write-gate-independent execution-unchanged session-unchanged unsafe-gates-closed AuthorizationBoundaryJoined ProductionTrustSatisfied RuntimeWriteGateEnabled LaunchAuthorized ProcessStartAuthorized BackendProcessStarted HostRootModified],
      "internal/runtime/appidentity/kde_restricted_launch_authorization_test.go" => %w[TestKDERestrictedLaunchAuthorizationRecordsPreparationOnly TestKDERestrictedLaunchAuthorizationRequiresExactDirective authorized-preparation-only explicit-test-root-only],
      "internal/runtime/appidentity/kde_restricted_launch_preflight.go" => %w[xnix.runtime.kde_restricted_launch_preflight.v1 kde-restricted-launch-preflight-record NewKDERestrictedLaunchPreflightRecord authorization-boundary safe-inputs preflight-readback explicit-blockers packet-assembly-boundary execution-unchanged session-unchanged unsafe-gates-closed PreflightBoundaryJoined ProductImageReady LaunchPreflightPassed LaunchAuthorized ProcessStartAuthorized CommandMaterialized BackendSelectedForLaunch BackendProcessStarted HostRootModified],
      "internal/runtime/appidentity/kde_restricted_launch_preflight_test.go" => %w[TestKDERestrictedLaunchPreflightRemainsBlocked explicit-test-root-only recipe-trust runtime-write-gate],
      "internal/runtime/appidentity/kde_test_launch_materialization.go" => %w[xnix.runtime.kde_test_launch_materialization.v1 kde-test-launch-materialization-record NewKDETestLaunchMaterializationRecord preflight-boundary materialization-readback materialized-plan-only write-gate-remains-blocked command-boundary launch-boundary execution-unchanged session-unchanged unsafe-gates-closed MaterializationBoundary PlanMaterialized CommandMaterialized RawExecutableExposed HostRootModified],
      "internal/runtime/appidentity/kde_test_launch_materialization_test.go" => %w[TestKDETestLaunchMaterializationRecordMaterializesPlanOnly TestKDETestLaunchMaterializationRecordRequiresExactBoundary blocked-plan-materialized test-only-review-plan],
      "internal/runtime/appidentity/kde_test_launch_materialization_fanout.go" => %w[xnix.runtime.kde_test_launch_materialization_fanout.v1 kde-test-launch-materialization-fanout-preview NewKDETestLaunchMaterializationFanOutPreview NewKDETestLaunchMaterializationFanOutPreviewFromReceipt LoadKDETestLaunchMaterializationRecordForFanOut read-only-existing-materialization-receipt read-only-receipt-consumption materialization-consumed session-fanout-consumed materialization-receipt-consumed execution-receipt-consumed session-receipt-consumed read-only-consumption surface-coverage surfaces-read-only desktop-side-effects-disabled notification-not-delivered launch-boundary-closed unsafe-data-hidden host-boundary-closed MaterializationReceiptConsumed ExecutionSessionFanOutConsumed FanOutWritesEnabled StateRootWritesEnabled NotificationSent CompatibilityCenterActionsEnabled RequestObjectsCreated RuntimeWritesEnabled HostRootModified],
      "internal/runtime/appidentity/kde_test_launch_materialization_fanout_test.go" => %w[TestKDETestLaunchMaterializationFanOutPreviewCoversKDESurfaces TestKDETestLaunchMaterializationFanOutPreviewConsumesExistingReceiptReadOnly TestKDETestLaunchMaterializationFanOutPreviewFromReceiptRejectsMissingPlan TestKDETestLaunchMaterializationFanOutPreviewRequiresExactBoundary compatibility-center task-manager tray notification review-plan-available read-only-existing-materialization-receipt],
      "internal/runtime/appidentity/kde_test_launch_materialization_receipt_lookup.go" => %w[xnix.runtime.kde_test_launch_materialization_receipt_lookup.v1 kde-test-launch-materialization-receipt-lookup-preview ResolveKDETestLaunchMaterializationReceipt owner-managed-materialization-receipt-lookup GetKDETestLaunchMaterializationReceiptLookup GetKDETestLaunchMaterializationReceiptLookupPreview opaque_materialization_receipt_id kde-test-launch-materialization-receipt-id missing-receipt owner-managed-opaque-lookup caller-paths-hidden missing-receipt-fails-closed read-only-lookup StateRootWritesEnabled RuntimeWritesEnabled FanOutWritesEnabled BackendLaunchEnabled HostRootModified],
      "internal/runtime/appidentity/kde_test_launch_materialization_receipt_lookup_test.go" => %w[TestResolveKDETestLaunchMaterializationReceiptReturnsOpaqueLookup TestResolveKDETestLaunchMaterializationReceiptRejectsUnknownOpaqueID owner-managed-materialization-receipt-lookup kde-test-launch-materialization-receipt-id missing-receipt],
      "internal/runtime/appidentity/kde_test_launch_materialization_fanout_owner_route.go" => %w[xnix.runtime.kde_test_launch_materialization_fanout_owner_route.v1 kde-test-launch-materialization-fanout-owner-route-preview NewKDETestLaunchMaterializationFanOutOwnerRoutePreview GetKDETestLaunchMaterializationFanOut GetKDETestLaunchMaterializationFanOutPreview owner-local-kde-test-launch-materialization-fanout opaque-lookup-consumed caller-paths-hidden missing-receipt-fails-closed surface-fanout-deferred desktop-side-effects-disabled owner-route-ready-production-dbus-blocked StateRootWritesEnabled RuntimeWritesEnabled FanOutWritesEnabled BackendLaunchEnabled HostRootModified],
      "internal/runtime/appidentity/kde_test_launch_materialization_fanout_owner_route_test.go" => %w[TestKDETestLaunchMaterializationFanOutOwnerRoutePreviewUsesOpaqueLookup TestKDETestLaunchMaterializationFanOutOwnerRouteRejectsUnknownOpaqueID owner-local-kde-test-launch-materialization-fanout kde-test-launch-materialization-receipt-id missing-receipt-fail-closed surface-fanout-deferred],
      "internal/runtime/appidentity/kde_test_launch_materialization_owner_route_audit.go" => %w[xnix.runtime.kde_test_launch_materialization_owner_route_audit.v1 kde-test-launch-materialization-owner-route-audit-preview NewKDETestLaunchMaterializationOwnerRouteAuditPreview GetKDETestLaunchMaterializationOwnerRouteAudit owner-local-route-smoke-covered owner-local-read-route-smoke-covered-production-dbus-blocked materialization-fanout-owner-smoke-coverage owner-smoke-coverage-present owner-route-present production-dbus-absent read-only-consume-registered caller-path-boundary receipt-creation-split owner-managed-opaque-lookup route-decision OwnerDispatchRoutePresent ProductionDBusMethodPresent MaterializationWritesStateRoot ReadOnlyReceiptConsumptionReady OwnerManagedOpaqueReceiptLookupReady OpaqueMaterializationReceiptIDSupported OwnerSmokeCoverageReady OwnerLocalRouteCandidateReady ProductionDBusExposureReady BackendLaunchEnabled HostRootModified],
      "internal/runtime/appidentity/kde_test_launch_materialization_owner_route_audit_test.go" => %w[TestKDETestLaunchMaterializationOwnerRouteAuditSeesConsumeSplit TestKDETestLaunchMaterializationOwnerRouteAuditFailsClosedWithoutSources owner-local-route-smoke-covered materialization-owner-route-sources-missing caller-path-boundary owner-managed-opaque-lookup owner-smoke-coverage-present],
      "cmd/xnix-runtime-go/kde_restricted_launch_authorization_commands.go" => %w[kde-restricted-launch-authorization-record runKDERestrictedLaunchAuthorizationRecord authorize-restricted-test-preparation test-only],
      "cmd/xnix-runtime-go/kde_restricted_launch_authorization_cli_test.go" => %w[TestKDERestrictedLaunchAuthorizationRecordCommandRequiresExplicitDirective TestKDERestrictedLaunchAuthorizationRecordCommandRejectsImplicitAuthorization preparation_authorized launch_authorized process_start_authorized],
      "cmd/xnix-runtime-go/kde_restricted_launch_preflight_commands.go" => %w[kde-restricted-launch-preflight-record runKDERestrictedLaunchPreflightRecord authorize-restricted-test-preparation test-only],
      "cmd/xnix-runtime-go/kde_restricted_launch_preflight_cli_test.go" => %w[TestKDERestrictedLaunchPreflightRecordCommandIsFailClosed TestKDERestrictedLaunchPreflightRecordCommandRequiresAuthorization ready_for_packet_assembly product_image_ready launch_preflight_passed],
      "cmd/xnix-runtime-go/kde_test_launch_materialization_commands.go" => %w[kde-test-launch-materialization-record runKDETestLaunchMaterializationRecord authorize-restricted-test-preparation test-only],
      "cmd/xnix-runtime-go/kde_test_launch_materialization_cli_test.go" => %w[TestKDETestLaunchMaterializationRecordCommandMaterializesPlanOnly TestKDETestLaunchMaterializationRecordCommandRequiresAuthorization plan_materialized command_materialized raw_executable_exposed],
      "cmd/xnix-runtime-go/kde_test_launch_materialization_fanout_commands.go" => %w[kde-test-launch-materialization-fanout-preview kde-test-launch-materialization-fanout-consume-preview kde-test-launch-materialization-receipt-lookup-preview kde-test-launch-materialization-fanout-owner-route-preview runKDETestLaunchMaterializationFanOutPreview runKDETestLaunchMaterializationFanOutConsumePreview runKDETestLaunchMaterializationReceiptLookupPreview runKDETestLaunchMaterializationFanOutOwnerRoutePreview authorize-restricted-test-preparation test-only materialization-plan-id],
      "cmd/xnix-runtime-go/kde_test_launch_materialization_fanout_cli_test.go" => %w[TestKDETestLaunchMaterializationFanOutPreviewCommandCoversKDESurfaces TestKDETestLaunchMaterializationFanOutConsumePreviewCommandReadsExistingReceipt TestKDETestLaunchMaterializationFanOutConsumePreviewCommandRequiresExistingReceipt TestKDETestLaunchMaterializationReceiptLookupPreviewCommand TestKDETestLaunchMaterializationReceiptLookupPreviewCommandRejectsBadInputs TestKDETestLaunchMaterializationFanOutOwnerRoutePreviewCommand TestKDETestLaunchMaterializationFanOutOwnerRoutePreviewCommandRejectsBadInputs TestKDETestLaunchMaterializationFanOutPreviewCommandRequiresAuthorization materialization_receipt_consumed execution_session_fan_out_consumed notification_sent fan_out_writes_enabled state_root_writes_enabled read-only-existing-materialization-receipt owner_managed_opaque_receipt_lookup_ready],
      "cmd/xnix-runtime-go/kde_test_launch_materialization_owner_route_audit_cli_test.go" => %w[TestKDETestLaunchMaterializationOwnerRouteAuditPreviewCommand kde-test-launch-materialization-owner-route-audit-preview route_decision owner_dispatch_route_present production_dbus_method_present owner_local_route_candidate_ready],
      "cmd/xnix-runtime-go/execution_ledger_commands.go" => %w[execution-ledger-record execution-session-record state-root portal-request executionLedgerPortalReceipts RecordSession],
      "cmd/xnix-runtime-go/execution_ledger_cli_test.go" => %w[TestExecutionLedgerRecordCommandConsumesPortalReceipt TestExecutionLedgerRecordCommandBlocksDeniedPortalReceipt TestExecutionSessionRecordCommandPersistsStatusFromLedger portal_permission_receipt_count],
      "internal/runtime/appidentity/application_readiness.go" => %w[xnix.runtime.application_readiness.v1 application-readiness-preview runtime-application-readiness-evidence-graph RecipeTrustDecision WriteGateDecision RealPortalTransportEnabled BackendLaunchEnabled StateRootPathExposed RawCommandExposed],
      "cmd/xnix-runtime-go/application_readiness_commands.go" => %w[application-readiness-preview ApplicationReadinessPreview artifact-receipt portal-operation snapshot-reason],
      "cmd/xnix-runtime-go/application_readiness_cli_test.go" => %w[TestApplicationReadinessPreviewCommand xnix.runtime.application_readiness.v1 runtime-application-readiness-evidence-graph backend_launch_enabled real_portal_transport_enabled],
      "internal/runtime/appidentity/runtime_write_gate.go" => %w[Launch WriteMethodDisabled]
    },
    summary: "Execution transaction and session status records persist under a controlled state root, can consume state-root Portal permission receipts for preflight evidence, persist preparation-only authorization and fail-closed restricted launch preflight packets, expose a Go Runtime application readiness evidence graph across install, backend lifecycle, Portal, snapshot, execution, and write-gate evidence, and keep real Launch disabled."
  },
  {
    id: "diagnostics-repair-ai-boundary",
    name: "Diagnostics, repair, and AI boundary",
    package: "P7",
    mainline_package: "M7",
    contract_files: %w[
      internal/runtime/appidentity/test_plan.go
      internal/runtime/appidentity/test_result.go
      internal/runtime/appidentity/repair_plan.go
      internal/runtime/appidentity/ai_diagnostics.go
    ],
    fixture_files: %w[
      internal/runtime/diagnostics/runner.go
      internal/runtime/diagnostics/ai.go
      internal/runtime/diagnostics/diagnostics_test.go
      internal/runtime/diagnostics/ai_test.go
      cmd/xnix-runtime-go/diagnostic_record_commands.go
      cmd/xnix-runtime-go/diagnostic_record_cli_test.go
      internal/runtime/appidentity/diagnostic_history_test.go
      internal/runtime/appidentity/crash_hang_signal_summary_test.go
      cmd/xnix-runtime-go/crash_hang_signal_summary_commands.go
      cmd/xnix-runtime-go/crash_hang_signal_summary_cli_test.go
      internal/runtime/appidentity/support_case_timeline_test.go
      cmd/xnix-runtime-go/support_case_timeline_commands.go
      cmd/xnix-runtime-go/support_case_timeline_cli_test.go
    ],
    state_files: %w[
      internal/runtime/diagnostics/record.go
      internal/runtime/diagnostics/history.go
      internal/runtime/appidentity/diagnostic_history.go
      internal/runtime/appidentity/crash_hang_signal_summary.go
      internal/runtime/appidentity/support_case_timeline.go
    ],
    smoke_files: [],
    gate_tokens: {
      "internal/runtime/diagnostics/ai.go" => %w[DisabledProvider FakeProvider],
      "internal/runtime/appidentity/crash_hang_signal_summary.go" => %w[xnix.runtime.crash_hang_signal_summary.v1 crash-hang-signal-summary-preview GetCrashHangSignalSummaryPreview crash hang timeout missing-dependency permission-denial graphics-issue network-issue regression-after-repair private_log_read file_content_read ai_provider_call_enabled repair_executed backend_process_started state_root_path_exposed host_root_modified],
      "cmd/xnix-runtime-go/crash_hang_signal_summary_commands.go" => %w[crash-hang-signal-summary-preview LenientHistory OpenRunRecordStoreReadOnly],
      "cmd/xnix-runtime-go/crash_hang_signal_summary_cli_test.go" => %w[TestCrashHangSignalSummaryPreviewCLI TestCrashHangSignalSummaryPreviewCLIMalformedRecordIsBlocked blocked-malformed-history],
      "internal/runtime/appidentity/support_case_timeline.go" => %w[xnix.runtime.support_case_timeline.v1 support-case-timeline-preview GetSupportCaseTimelinePreview diagnostic-runs blocked-actions repair-recommendations onboarding-gaps kde-entrypoint-state TicketCreated BundleExported AIProviderCalled RepairExecuted ActionExecuted BackendProcessStarted StateRootPathExposed HostRootModified],
      "cmd/xnix-runtime-go/support_case_timeline_commands.go" => %w[support-case-timeline-preview LenientHistory OpenRunRecordStoreReadOnly],
      "cmd/xnix-runtime-go/support_case_timeline_cli_test.go" => %w[TestSupportCaseTimelinePreviewCLI TestSupportCaseTimelinePreviewCLIMalformedRecordIsBlocked TestSupportCaseTimelinePreviewCLIMissingStateRootDoesNotCreate],
      "internal/runtime/diagnostics/record.go" => %w[xnix.runtime.diagnostic_run_record.v1 diagnostic-run-record go-runtime-state-root-diagnostic-run-record diagnostics-ledger state_root_path_exposed fixture_path_exposed ai_provider_called real_ai_provider_enabled repair_executed unsupported\ schema identity\ or\ path\ mismatch digest\ mismatch unsafe\ enabled\ gates marshalDiagnosticRunRecord],
      "internal/runtime/diagnostics/history.go" => %w[xnix.runtime.diagnostic_run_history.v1 diagnostic-run-history go-runtime-state-root-diagnostic-run-history state_root_path_exposed ai_provider_called repair_executed],
      "internal/runtime/appidentity/diagnostic_history.go" => %w[xnix.runtime.diagnostic_history_preview.v1 diagnostic-history-preview go-runtime-state-root-diagnostic-run-history+kde-read-model compatibility_center_card safe_for_ai_diagnostics ai_provider_call_enabled repair_execution_enabled backend_details_exposed],
      "cmd/xnix-runtime-go/diagnostic_record_commands.go" => %w[diagnostic-run-record diagnostic-run-history diagnostic-history-preview state-root fixture run-id],
      "internal/runtime/appidentity/ai_diagnostics.go" => %w[AIProviderCalled AutoExecutionAllowed]
    },
    summary: "Fixture diagnostics can persist state-root run records, expose KDE-safe history summaries, project diagnostic history into a Compatibility Center read model, summarize crash and hang signals from receipt metadata, and join diagnostic history, blocked actions, repair recommendations, onboarding gaps, and KDE entry-point state into a redacted support case timeline; real ticket creation, bundle export, provider calls, private log reads, backend launch, host mutation, and auto-repair remain disabled."
  },
  {
    id: "atomic-kde-image-qemu-acceptance",
    name: "Atomic KDE image and QEMU acceptance",
    package: "P8",
    mainline_package: "M9",
    contract_files: %w[
      image/kinoite/manifest.json
      image/kinoite/Containerfile
      docs/kde-image-pipeline.md
    ],
    fixture_files: %w[
      internal/runtime/image/manifest.go
      internal/runtime/image/smoke.go
      internal/runtime/image/smoke_test.go
      internal/runtime/image/restricted_smoke_packet.go
      internal/runtime/image/restricted_smoke_packet_test.go
      internal/runtime/appidentity/kde_restricted_product_smoke_checkpoint.go
      internal/runtime/appidentity/kde_restricted_product_smoke_checkpoint_test.go
      cmd/xnix-runtime-go/restricted_product_smoke_packet_commands.go
      cmd/xnix-runtime-go/restricted_product_smoke_packet_cli_test.go
      cmd/xnix-runtime-go/kde_restricted_product_smoke_checkpoint_commands.go
      cmd/xnix-runtime-go/kde_restricted_product_smoke_checkpoint_cli_test.go
      lib/xnix/full_smoke_report.rb
      scripts/build_kde_image.rb
      scripts/boot_kde_image.rb
      docs/release-evidence/v0.2.320-rc7-kde-product-smoke.json
      scripts/full_smoke.rb
      scripts/restricted_product_smoke_packet.rb
      test/test_restricted_product_smoke_packet.rb
    ],
    state_files: [],
    smoke_files: %w[
      scripts/boot_kde_image.rb
      test/test_kde_image.rb
      test/test_qemu.rb
      test/test_full_smoke_report.rb
      test/test_full_smoke_script.rb
    ],
    gate_tokens: {
      "image/kinoite/Containerfile" => %w[org.xnix.image org.xnix.flagship],
      "internal/runtime/image/smoke.go" => %w[SerialMarkers ProductImageReady RuntimeReady],
      "internal/runtime/image/restricted_smoke_packet.go" => %w[xnix.runtime.restricted_product_smoke_packet.v1 restricted-product-smoke-packet-preview dry-run-product-image-smoke-readiness runtime-owner artifact-trust backend-lifecycle portal-safety kde-entrypoints ReadyForAuthorizedSmoke ProductionRuntimeReady HumanAuthorizationRequired DockerExecuted QEMUExecuted ProductSmokeExecuted SerialLogPersisted LoopbackOnlyNetworking DockerSocketMounted HostNetworkEnabled BroadHostMountEnabled PrivilegedContainerRequired BackendLaunchEnabled HostRootModified ReleaseReady],
      "internal/runtime/appidentity/kde_restricted_product_smoke_checkpoint.go" => %w[xnix.runtime.kde_restricted_product_smoke_checkpoint.v1 kde-restricted-product-smoke-checkpoint-record NewKDERestrictedProductSmokeCheckpointRecord preflight-boundary product-image-manifest repository-evidence authorization-boundary execution-unchanged session-unchanged smoke-not-executed host-boundary ProductImageMetadataReady ReadyForAuthorizedSmoke CheckpointReady ReadyForTrainGate HumanAuthorizationRequired DockerExecuted QEMUExecuted ProductSmokeExecuted SerialLogPersisted ReleaseReady BackendLaunchEnabled HostRootModified RepositoryRootPathExposed ManifestSourcePathExposed],
      "internal/runtime/appidentity/kde_restricted_product_smoke_checkpoint_test.go" => %w[TestKDERestrictedProductSmokeCheckpointJoinsMetadataWithoutExecution recipe-trust runtime-write-gate assertKDERestrictedProductSmokeCheckpointDisabled],
      "cmd/xnix-runtime-go/restricted_product_smoke_packet_commands.go" => %w[restricted-product-smoke-packet-preview PrepareRestrictedProductSmokePacket repo-root manifest],
      "cmd/xnix-runtime-go/kde_restricted_product_smoke_checkpoint_commands.go" => %w[kde-restricted-product-smoke-checkpoint-record runKDERestrictedProductSmokeCheckpointRecord authorize-restricted-test-preparation test-only repo-root manifest],
      "cmd/xnix-runtime-go/kde_restricted_product_smoke_checkpoint_cli_test.go" => %w[TestKDERestrictedProductSmokeCheckpointRecordCommandIsReviewOnly TestKDERestrictedProductSmokeCheckpointRecordCommandRequiresAuthorization checkpoint_ready ready_for_train_gate docker_executed qemu_executed release_ready],
      "scripts/restricted_product_smoke_packet.rb" => %w[restricted-product-smoke-packet-preview JSON.pretty_generate render_markdown GOCACHE human_authorization_required docker_executed qemu_executed serial_log_persisted loopback_only_networking docker_socket_mounted host_network_enabled broad_host_mount_enabled host_root_modified],
      "test/test_restricted_product_smoke_packet.rb" => %w[xnix.runtime.restricted_product_smoke_packet.v1 ready_for_authorized_smoke human_authorization_required docker_executed qemu_executed release_ready],
      "lib/xnix/full_smoke_report.rb" => %w[xnix.full_smoke_report.v1 full-build-qemu-smoke-report qemu_network_restricted host_root_modified privileged_container_required docker_socket_mounted host_network_enabled],
      "scripts/full_smoke.rb" => %w[FullSmokeReport full-smoke-report.json full-smoke-report.md serial.log boot-system]
    },
    summary: "KDE image validation, a Go-owned restricted product smoke packet, a KDE-safe restricted product-image checkpoint, offline JSON and Markdown packet rendering, constrained QEMU smoke tooling, and full milestone smoke reports exist. Authorized q4 evidence records the Fedora 44 container build, clean qcow2 integrity check, persisted serial log, and KVM graphical-login pass; production Runtime ownership and Windows application execution remain disabled."
  },
  {
    id: "developer-verification-harness",
    name: "Developer verification harness",
    package: "P9",
    mainline_package: "M8",
    contract_files: %w[
      scripts/verify_layout.rb
      scripts/runtime_contract_drift_report.rb
      scripts/kde_first_presence_smoke.rb
      scripts/offline_application_fixture_matrix.rb
      scripts/merge_readiness_packet.rb
      docs/claude-code-windows-compatibility-workstreams.md
    ],
    fixture_files: %w[
      internal/runtime/appidentity/windows_compatibility_workstreams.go
      internal/runtime/appidentity/windows_compatibility_workstreams_test.go
      internal/runtime/appidentity/offline_application_fixture_matrix.go
      internal/runtime/appidentity/offline_application_fixture_matrix_test.go
      cmd/xnix-runtime-go/windows_compatibility_commands.go
      cmd/xnix-runtime-go/windows_compatibility_cli_test.go
      cmd/xnix-runtime-go/offline_application_fixture_matrix_commands.go
      cmd/xnix-runtime-go/offline_application_fixture_matrix_cli_test.go
      scripts/implementation_evidence_report.rb
      test/test_implementation_evidence_report.rb
      scripts/release_evidence_index.rb
      test/test_release_evidence_index.rb
      test/test_offline_application_fixture_matrix.rb
      test/test_merge_readiness_packet.rb
    ],
    state_files: [],
    smoke_files: [],
    gate_tokens: {
      "scripts/implementation_evidence_report.rb" => %w[contract-only fixture-implemented state-root-implemented smoke-owned production-gated windows_compatibility_first_wave windows_workstream_document],
      "scripts/release_evidence_index.rb" => %w[xnix.runtime.release_evidence_index.v1 release-evidence-index implementation-evidence+contract-drift+mainline-review+kde-first-presence implemented fixture-only contract-only blocked skipped human-authorized docker_executed qemu_executed automatic_release_tagging_enabled],
      "scripts/merge_readiness_packet.rb" => %w[xnix.runtime.merge_readiness_packet.v1 merge-readiness-packet TOOL_DEFINITIONS offline_only tool_statuses lane_classification protected_file_status unsafe_operation_status release_blocking_reasons restricted-docker-or-qemu-smoke-requires-human-authorization],
      "internal/runtime/appidentity/windows_compatibility_workstreams.go" => %w[xnix.runtime.windows_compatibility_workstreams.v1 windows-compatibility-workstreams-preview CW1 CW2 CW3 CW10 RuntimeOwned GoRuntimeBacked RubyCoreLogicAllowed KDEPolicyOwner BackendLaunchEnabled HostRootModified],
      "internal/runtime/appidentity/offline_application_fixture_matrix.go" => %w[xnix.runtime.offline_application_fixture_matrix.v1 offline-application-fixture-matrix-preview GetOfflineApplicationFixtureMatrix GetOfflineApplicationFixtureMatrixPreview built-in-fixtures+runtime-read-models+backend-adapter-contract-preview BackendAdapterContractRead BackendAdapterAuditReady OfflineApplicationFixtureBackendAdapterContract NewBackendAdapterContractPreview backend-adapter-noop-contract noop-contract-ready AdapterInvocationEnabled CommandMaterialized document-editor game installer launcher network-heavy tray-heavy unsupported NetworkFetchEnabled PackageManagerInvoked ArtifactStagingEnabled BackendLaunchEnabled DockerRequired QEMURequired HostRootModified],
      "internal/runtime/appidentity/offline_application_fixture_matrix_test.go" => %w[TestOfflineApplicationFixtureMatrixPreviewAuditsBackendAdapterContractMappings noop-contract-ready adapter_invocation_enabled command_materialized],
      "cmd/xnix-runtime-go/offline_application_fixture_matrix_commands.go" => %w[offline-application-fixture-matrix-preview OfflineApplicationFixtureMatrixOptions repeatedFixtureShapeIDs],
      "cmd/xnix-runtime-go/offline_application_fixture_matrix_cli_test.go" => %w[TestOfflineApplicationFixtureMatrixPreviewCommandAuditsBackendAdapterContractMappings backend_adapter_contract_read backend_adapter_audit_ready noop-contract-ready adapter_invocation_enabled command_materialized],
      "scripts/offline_application_fixture_matrix.rb" => %w[offline-application-fixture-matrix-preview JSON.pretty_generate render_markdown GOCACHE document-editor unsupported backend_adapter_contract_read backend_adapter_audit_ready noop-contract-ready adapter_invocation_enabled command_materialized network_fetch_enabled package_manager_invoked artifact_staging_enabled backend_launch_enabled docker_required qemu_required host_root_modified],
      "test/test_offline_application_fixture_matrix.rb" => %w[offline-application-fixture-matrix-preview blocked-unsupported noop-contract-ready backend_adapter_contract_read backend_adapter_audit_ready JSON.parse render_markdown],
      "test/test_merge_readiness_packet.rb" => %w[merge-readiness-packet malformed-json missing-command protected-claude-file-modified unclassified-files-present unsafe-operation-detected markdown]
    },
    summary: "Local evidence reporting classifies implementation depth, exposes KDE-first Windows compatibility workstream dispatch from a Go Runtime read model, checks a seven-shape offline application fixture matrix across Runtime evidence surfaces and no-op adapter contract profiles, aggregates offline merge readiness packets, flags orphan Runtime read methods, and now indexes release-critical evidence claims without running Docker, QEMU, package managers, adapter invocation, backend launch, staging, committing, pushing, release tagging, or host-root mutation."
  }
].freeze

def parse_options
  options = {
    root: PROJECT_ROOT,
    format: "json"
  }

  OptionParser.new do |parser|
    parser.banner = "Usage: ruby scripts/implementation_evidence_report.rb [--root PATH] [--format json|markdown]"
    parser.on("--root PATH", "Project root to inspect") do |value|
      options[:root] = Pathname.new(value).expand_path
    end
    parser.on("--format FORMAT", "Output format: json or markdown") do |value|
      options[:format] = value
    end
  end.parse!

  unless %w[json markdown].include?(options[:format])
    abort "implementation_evidence_report supports --format json or --format markdown"
  end

  options
end

def read_project_file(root, relative_path)
  root.join(relative_path).read
rescue Errno::ENOENT
  ""
end

def file_present?(root, relative_path)
  root.join(relative_path).file?
end

def present_files(root, relative_paths)
  relative_paths.select { |path| file_present?(root, path) }
end

def missing_files(root, relative_paths)
  relative_paths.reject { |path| file_present?(root, path) }
end

def token_results(root, token_map)
  token_map.flat_map do |relative_path, tokens|
    source = read_project_file(root, relative_path)
    tokens.map do |token|
      {
        "file" => relative_path,
        "token" => token,
        "present" => source.include?(token)
      }
    end
  end
end

def choose_status(contract_missing:, fixture_missing:, state_missing:, smoke_missing:, gate_missing:, has_state_files:, has_smoke_files:)
  return "missing" unless contract_missing.empty?
  return "contract-only" unless fixture_missing.empty?

  status = "fixture-implemented"
  status = "state-root-implemented" if has_state_files && state_missing.empty?
  status = "smoke-owned" if has_smoke_files && smoke_missing.empty?
  status
end

def domain_report(root, definition)
  contract_missing = missing_files(root, definition.fetch(:contract_files))
  fixture_missing = missing_files(root, definition.fetch(:fixture_files))
  state_files = definition.fetch(:state_files)
  smoke_files = definition.fetch(:smoke_files)
  state_missing = missing_files(root, state_files)
  smoke_missing = missing_files(root, smoke_files)
  tokens = token_results(root, definition.fetch(:gate_tokens))
  gate_missing = tokens.reject { |item| item.fetch("present") }

  status = choose_status(
    contract_missing: contract_missing,
    fixture_missing: fixture_missing,
    state_missing: state_missing,
    smoke_missing: smoke_missing,
    gate_missing: gate_missing,
    has_state_files: !state_files.empty?,
    has_smoke_files: !smoke_files.empty?
  )

  {
    "id" => definition.fetch(:id),
    "name" => definition.fetch(:name),
    "package" => definition.fetch(:package),
    "mainline_package" => definition.fetch(:mainline_package),
    "mainline_document" => MAINLINE_DOCUMENT,
    "status" => status,
    "rank" => STATUS_ORDER.fetch(status),
    "contract_files_present" => present_files(root, definition.fetch(:contract_files)),
    "contract_files_missing" => contract_missing,
    "fixture_files_present" => present_files(root, definition.fetch(:fixture_files)),
    "fixture_files_missing" => fixture_missing,
    "state_files_present" => present_files(root, state_files),
    "state_files_missing" => state_missing,
    "smoke_files_present" => present_files(root, smoke_files),
    "smoke_files_missing" => smoke_missing,
    "gate_tokens" => tokens,
    "gate_tokens_missing" => gate_missing,
    "production_gate_evidence" => gate_missing.empty?,
    "host_root_modified" => false,
    "network_required" => false,
    "privileged_container_required" => false,
    "backend_launch_enabled" => false,
    "summary" => definition.fetch(:summary)
  }
end

def dbus_contract_methods(root)
  read_project_file(root, "runtime/dbus/org.xnix.Compatibility1.xml").scan(/<method name="([^"]+)"/).flatten.uniq
end

def go_string_map_keys(source, variable_name)
  body = source[/var #{Regexp.escape(variable_name)} = map\[string\]string\{(.*?)\n\}/m, 1].to_s
  body.scan(/"([^"]+)"\s*:/).flatten
end

def owner_route_methods(root)
  source = read_project_file(root, "internal/runtime/appidentity/runtime_owner_route_manifest.go")
  go_string_map_keys(source, "runtimeOwnerRouteGoCommands")
end

def counts_for(domains)
  statuses = STATUS_ORDER.keys.to_h { |status| [status, 0] }
  domains.each { |domain| statuses[domain.fetch("status")] += 1 }

  {
    "total" => domains.length,
    "missing" => statuses.fetch("missing"),
    "contract_only" => statuses.fetch("contract-only"),
    "fixture_implemented" => statuses.fetch("fixture-implemented"),
    "state_root_implemented" => statuses.fetch("state-root-implemented"),
    "smoke_owned" => statuses.fetch("smoke-owned"),
    "production_gated" => statuses.fetch("production-gated"),
    "production_gate_evidence" => domains.count { |domain| domain.fetch("production_gate_evidence") }
  }
end

def mainline_first_wave(domains)
  domains_by_mainline = domains.to_h { |domain| [domain.fetch("mainline_package"), domain] }

  MAINLINE_FIRST_WAVE.filter_map do |entry|
    domain = domains_by_mainline[entry.fetch(:mainline_package)]
    next unless domain

    {
      "mainline_package" => entry.fetch(:mainline_package),
      "package" => domain.fetch("package"),
      "domain_id" => domain.fetch("id"),
      "domain_name" => domain.fetch("name"),
      "status" => domain.fetch("status"),
      "rank" => domain.fetch("rank"),
      "suggested_branch" => entry.fetch(:suggested_branch),
      "reason" => entry.fetch(:reason),
      "minimal_mergeable_outcome" => entry.fetch(:minimal_mergeable_outcome),
      "host_root_modified" => false,
      "network_required" => false,
      "privileged_container_required" => false,
      "backend_launch_enabled" => false
    }
  end
end

def windows_compatibility_first_wave(domains)
  domains_by_mainline = domains.to_h { |domain| [domain.fetch("mainline_package"), domain] }

  WINDOWS_COMPATIBILITY_FIRST_WAVE.filter_map do |entry|
    domain = domains_by_mainline[entry.fetch(:mainline_package)]
    next unless domain

    {
      "workstream" => entry.fetch(:workstream),
      "mainline_package" => entry.fetch(:mainline_package),
      "package" => domain.fetch("package"),
      "domain_id" => domain.fetch("id"),
      "domain_name" => domain.fetch("name"),
      "status" => domain.fetch("status"),
      "rank" => domain.fetch("rank"),
      "suggested_branch" => entry.fetch(:suggested_branch),
      "reason" => entry.fetch(:reason),
      "minimal_mergeable_outcome" => entry.fetch(:minimal_mergeable_outcome),
      "host_root_modified" => false,
      "network_required" => false,
      "privileged_container_required" => false,
      "backend_launch_enabled" => false
    }
  end
end

def build_report(root)
  version = read_project_file(root, "VERSION").strip
  domains = DOMAIN_DEFINITIONS.map { |definition| domain_report(root, definition) }
  first_wave = mainline_first_wave(domains)
  windows_first_wave = windows_compatibility_first_wave(domains)
  read_methods = dbus_contract_methods(root) - WRITE_METHODS
  route_methods = owner_route_methods(root)
  orphan_read_methods = read_methods - route_methods
  counts = counts_for(domains)
  highest_rank = domains.map { |domain| domain.fetch("rank") }.max || 0

  {
    "version" => version,
    "schema_version" => "xnix.runtime.implementation_evidence_report.v1",
    "report_type" => "implementation-evidence-report",
    "source" => "filesystem+runtime-contract+owner-route-manifest+empty-domain-packages+mainline-implementation-plan+windows-compatibility-workstreams",
    "runtime_owned" => true,
    "go_runtime_backed" => true,
    "kde_policy_owner" => false,
    "mainline_document" => MAINLINE_DOCUMENT,
    "mainline_plan_present" => file_present?(root, MAINLINE_DOCUMENT),
    "windows_workstream_document" => WINDOWS_WORKSTREAM_DOCUMENT,
    "windows_workstream_board_present" => file_present?(root, WINDOWS_WORKSTREAM_DOCUMENT),
    "mainline_package_count" => domains.map { |domain| domain.fetch("mainline_package") }.uniq.length,
    "mainline_first_wave" => first_wave,
    "windows_compatibility_first_wave" => windows_first_wave,
    "windows_compatibility_next_dispatch_summary" => "KDE-first Windows compatibility dispatch recommends CW1, CW2, CW3, and CW10 before KDE execution or product image acceptance.",
    "next_dispatch_packages" => first_wave,
    "next_dispatch_summary" => "First-wave mainline dispatch recommends M1, M2, M3, and M8 before execution or image acceptance.",
    "domain_status_order" => STATUS_ORDER.keys,
    "domains" => domains,
    "counts" => counts,
    "read_only_method_count" => read_methods.length,
    "orphan_read_methods" => orphan_read_methods,
    "orphan_preview_methods_detected" => !orphan_read_methods.empty?,
    "write_methods" => WRITE_METHODS,
    "write_methods_supported" => false,
    "write_method_dispatch_enabled" => false,
    "network_required" => false,
    "host_root_modified" => false,
    "privileged_container_required" => false,
    "backend_launch_enabled" => false,
    "production_ready" => false,
    "highest_evidence_rank" => highest_rank,
    "highest_evidence_status" => STATUS_ORDER.key(highest_rank),
    "desktop_safe_summary" => "Implementation evidence spans #{counts.fetch("total")} domains; production readiness remains gated."
  }
end

def render_markdown(report)
  lines = [
    "# Implementation Evidence Report",
    "",
    "- Version: #{report.fetch("version")}",
    "- Schema: #{report.fetch("schema_version")}",
    "- Domains: #{report.fetch("counts").fetch("total")}",
    "- Highest evidence status: #{report.fetch("highest_evidence_status")}",
    "- Mainline document: #{report.fetch("mainline_document")}",
    "- Mainline plan present: #{report.fetch("mainline_plan_present")}",
    "- Windows workstream document: #{report.fetch("windows_workstream_document")}",
    "- Windows workstream board present: #{report.fetch("windows_workstream_board_present")}",
    "- Next dispatch: #{report.fetch("next_dispatch_summary")}",
    "- Windows compatibility dispatch: #{report.fetch("windows_compatibility_next_dispatch_summary")}",
    "- Orphan read methods detected: #{report.fetch("orphan_preview_methods_detected")}",
    "- Production ready: #{report.fetch("production_ready")}",
    "- Host root modified: #{report.fetch("host_root_modified")}",
    "",
    "| Mainline | Package | Domain | Status | Missing fixtures | Missing gates |",
    "| --- | --- | --- | --- | --- | --- |"
  ]

  report.fetch("domains").each do |domain|
    missing_fixtures = domain.fetch("fixture_files_missing").empty? ? "-" : domain.fetch("fixture_files_missing").join(", ")
    missing_gates = domain.fetch("gate_tokens_missing").empty? ? "-" : domain.fetch("gate_tokens_missing").map { |item| "#{item.fetch("file")}:#{item.fetch("token")}" }.join(", ")
    lines << "| #{domain.fetch("mainline_package")} | #{domain.fetch("package")} | #{domain.fetch("name")} | #{domain.fetch("status")} | #{missing_fixtures} | #{missing_gates} |"
  end

  lines << ""
  lines << "## First-Wave Dispatch"
  lines << ""
  lines << "| Mainline | Branch | Status | Minimal mergeable outcome |"
  lines << "| --- | --- | --- | --- |"
  report.fetch("mainline_first_wave").each do |entry|
    lines << "| #{entry.fetch("mainline_package")} | `#{entry.fetch("suggested_branch")}` | #{entry.fetch("status")} | #{entry.fetch("minimal_mergeable_outcome")} |"
  end

  lines << ""
  lines << "## Windows Compatibility First-Wave Workstreams"
  lines << ""
  lines << "| Workstream | Mainline | Branch | Status | Minimal mergeable outcome |"
  lines << "| --- | --- | --- | --- | --- |"
  report.fetch("windows_compatibility_first_wave").each do |entry|
    lines << "| #{entry.fetch("workstream")} | #{entry.fetch("mainline_package")} | `#{entry.fetch("suggested_branch")}` | #{entry.fetch("status")} | #{entry.fetch("minimal_mergeable_outcome")} |"
  end

  lines << ""
  lines << report.fetch("desktop_safe_summary")
  lines.join("\n")
end

options = parse_options
report = build_report(options.fetch(:root))

case options.fetch(:format)
when "json"
  puts JSON.pretty_generate(report)
when "markdown"
  puts render_markdown(report)
end

exit(report.fetch("orphan_preview_methods_detected") ? 1 : 0)
