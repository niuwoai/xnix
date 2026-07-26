package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"xnix.local/xnix/internal/runtime/appidentity"
	"xnix.local/xnix/internal/runtime/artifact"
)

func main() {
	if err := run(os.Args[1:], os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func run(args []string, stdout io.Writer) error {
	if len(args) == 0 {
	}

	switch args[0] {
	case "acquisition-preflight-preview":
		return runAcquisitionPreflightPreview(args[1:], stdout)
	case "ai-diagnostic-input-preview":
		return runAIDiagnosticInputPreview(args[1:], stdout)
	case "ai-diagnostic-recommendation-preview":
		return runAIDiagnosticRecommendationPreview(args[1:], stdout)
	case "ai-repair-approval-gate-preview":
		return runAIRepairApprovalGatePreview(args[1:], stdout)
	case "application-preview":
		return runApplicationPreview(args[1:], stdout)
	case "application-readiness-preview":
		return runApplicationReadinessPreview(args[1:], stdout)
	case "application-upgrade-impact-preview":
		return runApplicationUpgradeImpactPreview(args[1:], stdout)
	case "applications-preview":
		return runApplicationsPreview(args[1:], stdout)
	case "artifact-manifest-preview":
		return runArtifactManifestPreview(args[1:], stdout)
	case "artifact-stage-record":
		return runArtifactStageRecord(args[1:], stdout)
	case "backend-adapter-contract-preview":
		return runBackendAdapterContractPreview(args[1:], stdout)
	case "backend-adapter-contract-owner-route-audit-preview":
		return runBackendAdapterContractOwnerRouteAuditPreview(args[1:], stdout)
	case "backend-adapter-redacted-profile-audit-preview":
		return runBackendAdapterRedactedProfileAuditPreview(args[1:], stdout)
	case "backend-binding-preview":
		return runBackendBindingPreview(args[1:], stdout)
	case "backend-capability-matrix-preview":
		return runBackendCapabilityMatrixPreview(args[1:], stdout)
	case "backend-environment-preview":
		return runBackendEnvironmentPreview(args[1:], stdout)
	case "backend-lifecycle-preview":
		return runBackendLifecyclePreview(args[1:], stdout)
	case "backend-lifecycle-record":
		return runBackendLifecycleRecord(args[1:], stdout)
	case "backend-manager-preview":
		return runBackendManagerPreview(args[1:], stdout)
	case "backend-manager-record":
		return runBackendManagerRecord(args[1:], stdout)
	case "backend-selection-preview":
		return runBackendSelectionPreview(args[1:], stdout)
	case "compatibility-center-preview":
		return runCompatibilityCenterPreview(args[1:], stdout)
	case "compatibility-install-preview":
		return runCompatibilityInstallPreview(args[1:], stdout)
	case "compatibility-backend-fallback-preview":
		return runCompatibilityBackendFallbackPreview(args[1:], stdout)
	case "compatibility-onboarding-checklist-preview":
		return runCompatibilityOnboardingChecklistPreview(args[1:], stdout)
	case "crash-hang-signal-summary-preview":
		return runCrashHangSignalSummaryPreview(args[1:], stdout)
	case "current-mainline-ownership-audit-preview":
		return runCurrentMainlineOwnershipAuditPreview(args[1:], stdout)
	case "desktop-activation-bundle-preview":
		return runDesktopActivationBundlePreview(args[1:], stdout)
	case "desktop-activation-manifest-preview":
		return runDesktopActivationManifestPreview(args[1:], stdout)
	case "desktop-activation-preflight-preview":
		return runDesktopActivationPreflightPreview(args[1:], stdout)
	case "desktop-activation-stage":
		return runDesktopActivationStage(args[1:], stdout)
	case "desktop-activation-staging-preview":
		return runDesktopActivationStagingPreview(args[1:], stdout)
	case "desktop-activation-status-preview":
		return runDesktopActivationStatusPreview(args[1:], stdout)
	case "desktop-deactivation-dry-run-preview":
		return runDesktopDeactivationDryRunPreview(args[1:], stdout)
	case "desktop-activation-transaction-preview":
		return runDesktopActivationTransactionPreview(args[1:], stdout)
	case "desktop-entry-preview":
		return runDesktopEntryPreview(args[1:], stdout)
	case "desktop-external-winapp-launch-packet-preview":
		return runDesktopExternalWinAppLaunchPacketPreview(args[1:], stdout)
	case "desktop-icon-preview":
		return runDesktopIconPreview(args[1:], stdout)
	case "desktop-identity-plan":
		return runDesktopIdentityPlan(args[1:], stdout)
	case "desktop-resource-bridge-preview":
		return runDesktopResourceBridgePreview(args[1:], stdout)
	case "desktop-safety-policy-preview":
		return runDesktopSafetyPolicyPreview(args[1:], stdout)
	case "desktop-trigger-staged-invocation-readiness-preview":
		return runDesktopTriggerStagedInvocationReadinessPreview(args[1:], stdout)
	case "diagnostic-history-preview":
		return runDiagnosticHistoryPreview(args[1:], stdout)
	case "diagnostic-run-history":
		return runDiagnosticRunHistory(args[1:], stdout)
	case "diagnostic-run-record":
		return runDiagnosticRunRecord(args[1:], stdout)
	case "diagnostics-preview":
		return runDiagnosticsPreview(args[1:], stdout)
	case "dolphin-ai-analysis-preview":
		return runDolphinAIAnalysisPreview(args[1:], stdout)
	case "dolphin-drop-preview":
		return runDolphinDropPreview(args[1:], stdout)
	case "engine-catalog-preview":
		return runEngineCatalogPreview(args[1:], stdout)
	case "execution-decision-preview":
		return runExecutionDecisionPreview(args[1:], stdout)
	case "external-winapp-import-record":
		return runExternalWinAppImportRecord(args[1:], stdout)
	case "windows-external-app-run":
		return runExternalWinAppRun(args[1:], stdout)
	case "execution-ledger-record":
		return runExecutionLedgerRecord(args[1:], stdout)
	case "execution-preflight-preview":
		return runExecutionPreflightPreview(args[1:], stdout)
	case "execution-readiness-preview":
		return runExecutionReadinessPreview(args[1:], stdout)
	case "execution-request-preview":
		return runExecutionRequestPreview(args[1:], stdout)
	case "execution-resource-grant-preview":
		return runExecutionResourceGrantPreview(args[1:], stdout)
	case "execution-review-preview":
		return runExecutionReviewPreview(args[1:], stdout)
	case "execution-session-preview":
		return runExecutionSessionPreview(args[1:], stdout)
	case "execution-session-record":
		return runExecutionSessionRecord(args[1:], stdout)
	case "execution-session-status-preview":
		return runExecutionSessionStatusPreview(args[1:], stdout)
	case "execution-transaction-preview":
		return runExecutionTransactionPreview(args[1:], stdout)
	case "file-open-preview":
		return runFileOpenPreview(args[1:], stdout)
	case "kde-action-card-deck-preview":
		return runKDEActionCardDeckPreview(args[1:], stdout)
	case "kde-action-card-preview":
		return runKDEActionCardPreview(args[1:], stdout)
	case "kde-action-dependency-graph-preview":
		return runKDEActionDependencyGraphPreview(args[1:], stdout)
	case "kde-action-queue-preview":
		return runKDEActionQueuePreview(args[1:], stdout)
	case "kde-action-preflight-preview":
		return runKDEActionPreflightPreview(args[1:], stdout)
	case "kde-action-receipt-preview":
		return runKDEActionReceiptPreview(args[1:], stdout)
	case "kde-action-review-preview":
		return runKDEActionReviewPreview(args[1:], stdout)
	case "kde-action-status-preview":
		return runKDEActionStatusPreview(args[1:], stdout)
	case "kde-application-surface-preview":
		return runKDEApplicationSurfacePreview(args[1:], stdout)
	case "kde-backend-lifecycle-evidence-record":
		return runKDEBackendLifecycleEvidenceRecord(args[1:], stdout)
	case "kde-center-page-preview":
		return runKDECenterPagePreview(args[1:], stdout)
	case "kde-center-page-section-detail-preview":
		return runKDECenterPageSectionDetailPreview(args[1:], stdout)
	case "kde-center-page-sections-preview":
		return runKDECenterPageSectionsPreview(args[1:], stdout)
	case "kde-entrypoint-action-preview":
		return runKDEEntryPointActionPreview(args[1:], stdout)
	case "kde-entrypoints-preview":
		return runKDEEntryPointsPreview(args[1:], stdout)
	case "kde-fake-execution-evidence-record":
		return runKDEFakeExecutionEvidenceRecord(args[1:], stdout)
	case "kde-fake-portal-evidence-record":
		return runKDEFakePortalEvidenceRecord(args[1:], stdout)
	case "kde-snapshot-diagnostics-evidence-record":
		return runKDESnapshotDiagnosticsEvidenceRecord(args[1:], stdout)
	case "kde-test-launch-materialization-fanout-consume-preview":
		return runKDETestLaunchMaterializationFanOutConsumePreview(args[1:], stdout)
	case "kde-test-launch-materialization-fanout-owner-route-preview":
		return runKDETestLaunchMaterializationFanOutOwnerRoutePreview(args[1:], stdout)
	case "kde-test-launch-materialization-fanout-preview":
		return runKDETestLaunchMaterializationFanOutPreview(args[1:], stdout)
	case "kde-test-launch-materialization-owner-route-audit-preview":
		return runKDETestLaunchMaterializationOwnerRouteAuditPreview(args[1:], stdout)
	case "kde-test-launch-materialization-receipt-lookup-preview":
		return runKDETestLaunchMaterializationReceiptLookupPreview(args[1:], stdout)
	case "kde-test-launch-materialization-record":
		return runKDETestLaunchMaterializationRecord(args[1:], stdout)
	case "kde-integration-status-preview":
		return runKDEIntegrationStatusPreview(args[1:], stdout)
	case "kde-journey-evidence-preview":
		return runKDEJourneyEvidencePreview(args[1:], stdout)
	case "kde-notification-digest-preview":
		return runKDENotificationDigestPreview(args[1:], stdout)
	case "kde-offline-application-identity-preview":
		return runKDEOfflineApplicationIdentityPreview(args[1:], stdout)
	case "kde-restricted-launch-authorization-record":
		return runKDERestrictedLaunchAuthorizationRecord(args[1:], stdout)
	case "kde-restricted-launch-preflight-record":
		return runKDERestrictedLaunchPreflightRecord(args[1:], stdout)
	case "kde-restricted-product-smoke-checkpoint-record":
		return runKDERestrictedProductSmokeCheckpointRecord(args[1:], stdout)
	case "kde-search-visibility-plan-preview":
		return runKDESearchVisibilityPlanPreview(args[1:], stdout)
	case "kde-shell-integration-preview":
		return runKDEShellIntegrationPreview(args[1:], stdout)
	case "krunner-query-preview":
		return runKRunnerQueryPreview(args[1:], stdout)
	case "kwin-window-rule-preview":
		return runKWinWindowRulePreview(args[1:], stdout)
	case "known-app-launch-gate-preview":
		return runKnownAppLaunchGatePreview(args[1:], stdout)
	case "known-app-launch-authorization-receipt-preview":
		return runKnownAppLaunchAuthorizationReceiptPreview(args[1:], stdout)
	case "known-app-controlled-dispatch-request-preview":
		return runKnownAppControlledDispatchRequestPreview(args[1:], stdout)
	case "known-app-controlled-execution-session-preview":
		return runKnownAppControlledExecutionSessionPreview(args[1:], stdout)
	case "known-app-controlled-execution-session-record":
		return runKnownAppControlledExecutionSessionRecord(args[1:], stdout)
	case "known-app-controlled-execution-session-consume-preview":
		return runKnownAppControlledExecutionSessionConsumePreview(args[1:], stdout)
	case "known-app-session-gated-launch-review-preview":
		return runKnownAppSessionGatedLaunchReviewPreview(args[1:], stdout)
	case "known-app-session-gated-launch-review-receipt-record":
		return runKnownAppSessionGatedLaunchReviewReceiptRecord(args[1:], stdout)
	case "known-app-session-gated-launch-review-gate-preview":
		return runKnownAppSessionGatedLaunchReviewGatePreview(args[1:], stdout)
	case "known-app-session-gated-controlled-dispatch-request-preview":
		return runKnownAppSessionGatedControlledDispatchRequestPreview(args[1:], stdout)
	case "known-app-kde-runtime-status-launch-request-preview":
		return runKnownAppKDERuntimeStatusLaunchRequestPreview(args[1:], stdout)
	case "known-app-kde-runtime-status-launch-execution":
		return runKnownAppKDERuntimeStatusLaunchExecution(args[1:], stdout)
	case "known-app-kde-runtime-status-launch-evidence-record":
		return runKnownAppKDERuntimeStatusLaunchEvidenceRecord(args[1:], stdout)
	case "known-app-kde-runtime-status-launch-evidence-preview":
		return runKnownAppKDERuntimeStatusLaunchEvidencePreview(args[1:], stdout)
	case "known-app-kde-runtime-status-launch-action-trigger-preview":
		return runKnownAppKDERuntimeStatusLaunchActionTriggerPreview(args[1:], stdout)
	case "known-app-matrix-evidence-preview":
		return runKnownAppMatrixEvidencePreview(args[1:], stdout)
	case "known-app-verified-catalog-preview":
		return runKnownAppVerifiedCatalogPreview(args[1:], stdout)
	case "known-app-verified-catalog-run-plan-preview":
		return runKnownAppVerifiedCatalogRunPlanPreview(args[1:], stdout)
	case "gui-smoke-evidence-preview":
		return runGUISmokeEvidencePreview(args[1:], stdout)
	case "real-winapp-gui-evidence-packet-preview":
		return runRealWinAppGUIEvidencePacketPreview(args[1:], stdout)
	case "real-winapp-run-receipt-summary-preview":
		return runRealWinAppRunReceiptSummaryPreview(args[1:], stdout)
	case "real-winapp-run-acceptance-preview":
		return runRealWinAppRunAcceptancePreview(args[1:], stdout)
	case "q4-sample-notepad-acceptance-preview":
		return runQ4SampleNotepadAcceptancePreview(args[1:], stdout)
	case "q4-winapp-acceptance-preview":
		return runQ4WinAppAcceptancePreview(args[1:], stdout)
	case "known-existing-winapp-acceptance-preview":
		return runKnownExistingWinAppAcceptancePreview(args[1:], stdout)
	case "known-app-runtime-status-launch-owner-fixture-record":
		return runKnownAppRuntimeStatusLaunchOwnerFixtureRecord(args[1:], stdout)
	case "known-app-runtime-status-launch-owner-trigger-preview":
		return runKnownAppRuntimeStatusLaunchOwnerTriggerPreview(args[1:], stdout)
	case "managed-launcher-acceptance-report-preview":
		return runManagedLauncherAcceptanceReportPreview(args[1:], stdout)
	case "desktop-trigger-dry-run-request-review-preview":
		return runDesktopTriggerDryRunRequestReviewPreview(args[1:], stdout)
	case "desktop-trigger-service-call-materialization-preview":
		return runDesktopTriggerServiceCallMaterializationPreview(args[1:], stdout)
	case "desktop-trigger-request-preflight-preview":
		return runDesktopTriggerRequestPreflightPreview(args[1:], stdout)
	case "kde-controlled-launch-action-preview":
		return runKDEControlledLaunchActionPreview(args[1:], stdout)
	case "kde-controlled-launch-action-surface-audit-preview":
		return runKDEControlledLaunchActionSurfaceAuditPreview(args[1:], stdout)
	case "kde-controlled-launch-session-bus-smoke-plan-preview":
		return runKDEControlledLaunchSessionBusSmokePlanPreview(args[1:], stdout)
	case appidentity.KnownAppKDERuntimeStatusLaunchAction:
		return runKnownAppKDEShowRuntimeControlledLaunch(args[1:], stdout)
	case "launch-intent-preview":
		return runLaunchIntentPreview(args[1:], stdout)
	case "mimeapps-preview":
		return runMIMEAppsPreview(args[1:], stdout)
	case "mode-switch-preview":
		return runModeSwitchPreview(args[1:], stdout)
	case "multi-application-install-queue-preview":
		return runMultiApplicationInstallQueuePreview(args[1:], stdout)
	case "notification-preview":
		return runNotificationPreview(args[1:], stdout)
	case "offline-application-fixture-matrix-preview":
		return runOfflineApplicationFixtureMatrixPreview(args[1:], stdout)
	case "owner-service-launch-envelope-guard-preview":
		return runOwnerServiceLaunchEnvelopeGuardPreview(args[1:], stdout)
	case "package-source-preview":
		return runPackageSourcePreview(args[1:], stdout)
	case "permission-evidence-audit-preview":
		return runPermissionEvidenceAuditPreview(args[1:], stdout)
	case "permission-review-preview":
		return runPermissionReviewPreview(args[1:], stdout)
	case "portal-access-policy-preview":
		return runPortalAccessPolicyPreview(args[1:], stdout)
	case "portal-permission-renewal-preview":
		return runPortalPermissionRenewalPreview(args[1:], stdout)
	case "portal-request-preview":
		return runPortalRequestPreview(args[1:], stdout)
	case "portal-request-record":
		return runPortalRequestRecord(args[1:], stdout)
	case "production-authorization-consumption-audit-preview":
		return runProductionAuthorizationConsumptionAuditPreview(args[1:], stdout)
	case "production-dbus-gate-review-preview":
		return runProductionDBusGateReviewPreview(args[1:], stdout)
	case "production-dbus-human-authorization-preflight-preview":
		return runProductionDBusHumanAuthorizationPreflightPreview(args[1:], stdout)
	case "production-dbus-method-review-preview":
		return runProductionDBusMethodReviewPreview(args[1:], stdout)
	case "production-desktop-side-effect-review-preview":
		return runProductionDesktopSideEffectReviewPreview(args[1:], stdout)
	case "production-human-authorization-receipt-consolidation-preview":
		return runProductionHumanAuthorizationReceiptConsolidationPreview(args[1:], stdout)
	case "production-receipt-acceptance-propagation-preflight-preview":
		return runProductionReceiptAcceptancePropagationPreflightPreview(args[1:], stdout)
	case "production-receipt-persistence-threat-review-preview":
		return runProductionReceiptPersistenceThreatReviewPreview(args[1:], stdout)
	case "production-receipt-notification-action-safety-audit-preview":
		return runProductionReceiptNotificationActionSafetyAuditPreview(args[1:], stdout)
	case "production-receipt-notification-action-request-object-audit-preview":
		return runProductionReceiptNotificationActionRequestObjectAuditPreview(args[1:], stdout)
	case "production-receipt-notification-action-dispatch-authorization-audit-preview":
		return runProductionReceiptNotificationActionDispatchAuthorizationAuditPreview(args[1:], stdout)
	case "production-receipt-notification-action-dispatch-dry-run-audit-preview":
		return runProductionReceiptNotificationActionDispatchDryRunAuditPreview(args[1:], stdout)
	case "production-receipt-notification-action-dry-run-result-visibility-audit-preview":
		return runProductionReceiptNotificationActionDryRunResultVisibilityAuditPreview(args[1:], stdout)
	case "production-receipt-notification-action-dry-run-result-persistence-authorization-audit-preview":
		return runProductionReceiptNotificationActionDryRunResultPersistenceAuthorizationAuditPreview(args[1:], stdout)
	case "production-receipt-notification-action-dry-run-result-retention-redaction-policy-audit-preview":
		return runProductionReceiptNotificationActionDryRunResultRetentionRedactionPolicyAuditPreview(args[1:], stdout)
	case "production-receipt-notification-action-dry-run-result-opaque-lookup-audit-preview":
		return runProductionReceiptNotificationActionDryRunResultOpaqueLookupAuditPreview(args[1:], stdout)
	case "production-receipt-notification-action-dry-run-result-lookup-route-authorization-audit-preview":
		return runProductionReceiptNotificationActionDryRunResultLookupRouteAuthorizationAuditPreview(args[1:], stdout)
	case "production-receipt-notification-action-dry-run-result-lookup-consumer-redaction-audit-preview":
		return runProductionReceiptNotificationActionDryRunResultLookupConsumerRedactionAuditPreview(args[1:], stdout)
	case "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-gate-audit-preview":
		return runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementGateAuditPreview(args[1:], stdout)
	case "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-authorization-receipt-audit-preview":
		return runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementAuthorizationReceiptAuditPreview(args[1:], stdout)
	case "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-receipt-writer-authorization-audit-preview":
		return runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptWriterAuthorizationAuditPreview(args[1:], stdout)
	case "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-receipt-persistence-authorization-audit-preview":
		return runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptPersistenceAuthorizationAuditPreview(args[1:], stdout)
	case "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-receipt-acceptance-authorization-audit-preview":
		return runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptAcceptanceAuthorizationAuditPreview(args[1:], stdout)
	case "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-receipt-consumption-gate-audit-preview":
		return runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumptionGateAuditPreview(args[1:], stdout)
	case "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-receipt-consumer-authorization-audit-preview":
		return runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerAuthorizationAuditPreview(args[1:], stdout)
	case "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-receipt-consumer-enablement-gate-audit-preview":
		return runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerEnablementGateAuditPreview(args[1:], stdout)
	case "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-status-fanout-audit-preview":
		return runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusFanOutAuditPreview(args[1:], stdout)
	case "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-status-persistence-authorization-audit-preview":
		return runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusPersistenceAuthorizationAuditPreview(args[1:], stdout)
	case "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-persistence-write-model-audit-preview":
		return runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusPersistenceWriteModelAuditPreview(args[1:], stdout)
	case "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-writer-authorization-gate-audit-preview":
		return runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAuthorizationGateAuditPreview(args[1:], stdout)
	case "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-writer-authorization-receipt-audit-preview":
		return runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAuthorizationReceiptAuditPreview(args[1:], stdout)
	case "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-writer-authorization-receipt-acceptance-audit-preview":
		return runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAuthorizationReceiptAcceptanceAuditPreview(args[1:], stdout)
	case "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-writer-accepted-receipt-gate-audit-preview":
		return runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAcceptedReceiptGateAuditPreview(args[1:], stdout)
	case "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-writer-grant-audit-preview":
		return runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterGrantAuditPreview(args[1:], stdout)
	case "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-writer-enablement-audit-preview":
		return runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterEnablementAuditPreview(args[1:], stdout)
	case "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-writer-implementation-preview":
		return runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedWriterImplementationPreview(args[1:], stdout)
	case "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-writer-persistence-authorization-preview":
		return runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterPersistenceAuthorizationPreview(args[1:], stdout)
	case "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-persistence-writer-preview":
		return runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedPersistenceWriterPreview(args[1:], stdout)
	case "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-storage-persistence-gate-audit-preview":
		return runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusStoragePersistenceGateAuditPreview(args[1:], stdout)
	case "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-persistence-record-contract-preview":
		return runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusPersistenceRecordContractPreview(args[1:], stdout)
	case "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-implementation-preview":
		return runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterImplementationPreview(args[1:], stdout)
	case "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-policy-audit-preview":
		return runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootPolicyAuditPreview(args[1:], stdout)
	case "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-consumption-audit-preview":
		return runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptConsumptionAuditPreview(args[1:], stdout)
	case "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-evidence-preview":
		return runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupEvidencePreview(args[1:], stdout)
	case "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-boundary-preview":
		return runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultBoundaryPreview(args[1:], stdout)
	case "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-preview":
		return runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionPreview(args[1:], stdout)
	case "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-evidence-audit-preview":
		return runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionEvidenceAuditPreview(args[1:], stdout)
	case "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-authorization-evidence-gate-preview":
		return runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteAuthorizationEvidenceGatePreview(args[1:], stdout)
	case "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-evidence-gate-preview":
		return runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementEvidenceGatePreview(args[1:], stdout)
	case "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-receipt-gate-preview":
		return runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementReceiptGatePreview(args[1:], stdout)
	case "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-receipt-acceptance-audit-preview":
		return runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementReceiptAcceptanceAuditPreview(args[1:], stdout)
	case "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-accepted-receipt-gate-audit-preview":
		return runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementAcceptedReceiptGateAuditPreview(args[1:], stdout)
	case "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-grant-audit-preview":
		return runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteGrantAuditPreview(args[1:], stdout)
	case "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-enablement-audit-preview":
		return runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteEnablementAuditPreview(args[1:], stdout)
	case "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-gate-audit-preview":
		return runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchGateAuditPreview(args[1:], stdout)
	case "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-authorization-audit-preview":
		return runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchAuthorizationAuditPreview(args[1:], stdout)
	case "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-execution-gate-audit-preview":
		return runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunExecutionGateAuditPreview(args[1:], stdout)
	case "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-capture-audit-preview":
		return runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultCaptureAuditPreview(args[1:], stdout)
	case "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-redaction-boundary-audit-preview":
		return runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultRedactionBoundaryAuditPreview(args[1:], stdout)
	case "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-redacted-presentation-eligibility-audit-preview":
		return runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultRedactedPresentationEligibilityAuditPreview(args[1:], stdout)
	case "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-presentation-consent-audit-preview":
		return runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultPresentationConsentAuditPreview(args[1:], stdout)
	case "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-presentation-audit-preview":
		return runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationPresentationAuditPreview(args[1:], stdout)
	case "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-delivery-authorization-audit-preview":
		return runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationDeliveryAuthorizationAuditPreview(args[1:], stdout)
	case "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-delivery-grant-audit-preview":
		return runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationDeliveryGrantAuditPreview(args[1:], stdout)
	case "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-enablement-audit-preview":
		return runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionEnablementAuditPreview(args[1:], stdout)
	case "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-audit-preview":
		return runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectAuditPreview(args[1:], stdout)
	case "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-authorization-audit-preview":
		return runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchAuthorizationAuditPreview(args[1:], stdout)
	case "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-audit-preview":
		return runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunAuditPreview(args[1:], stdout)
	case "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-visibility-audit-preview":
		return runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultVisibilityAuditPreview(args[1:], stdout)
	case "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-authorization-audit-preview":
		return runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceAuthorizationAuditPreview(args[1:], stdout)
	case "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-audit-preview":
		return runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptAuditPreview(args[1:], stdout)
	case "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-acceptance-audit-preview":
		return runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptAcceptanceAuditPreview(args[1:], stdout)
	case "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumption-gate-audit-preview":
		return runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumptionGateAuditPreview(args[1:], stdout)
	case "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-authorization-audit-preview":
		return runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerAuthorizationAuditPreview(args[1:], stdout)
	case "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-gate-audit-preview":
		return runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementGateAuditPreview(args[1:], stdout)
	case "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-audit-preview":
		return runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptAuditPreview(args[1:], stdout)
	case "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-writer-authorization-audit-preview":
		return runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptWriterAuthorizationAuditPreview(args[1:], stdout)
	case "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-persistence-authorization-audit-preview":
		return runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptPersistenceAuthorizationAuditPreview(args[1:], stdout)
	case "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-acceptance-authorization-audit-preview":
		return runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptAcceptanceAuthorizationAuditPreview(args[1:], stdout)
	case "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-acceptance-gate-audit-preview":
		return runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptAcceptanceGateAuditPreview(args[1:], stdout)
	case "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-accepted-receipt-gate-audit-preview":
		return runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptAcceptedReceiptGateAuditPreview(args[1:], stdout)
	case "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-grant-audit-preview":
		return runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptGrantAuditPreview(args[1:], stdout)
	case "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-enablement-audit-preview":
		return runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptEnablementAuditPreview(args[1:], stdout)
	case "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-closed-implementation-preview":
		return runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptClosedImplementationPreview(args[1:], stdout)
	case "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-persistence-authorization-preview":
		return runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptPersistenceAuthorizationPreview(args[1:], stdout)
	case "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-closed-persistence-implementation-preview":
		return runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptClosedPersistenceImplementationPreview(args[1:], stdout)
	case "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-persistence-gate-audit-preview":
		return runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptStoragePersistenceGateAuditPreview(args[1:], stdout)
	case "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-contract-preview":
		return runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptStorageRecordContractPreview(args[1:], stdout)
	case "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-closed-storage-record-writer-implementation-preview":
		return runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptClosedStorageRecordWriterImplementationPreview(args[1:], stdout)
	case "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-authorization-preview":
		return runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptStorageRecordWriterAuthorizationPreview(args[1:], stdout)
	case "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-authorization-receipt-preview":
		return runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptStorageRecordWriterAuthorizationReceiptPreview(args[1:], stdout)
	case "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-authorization-receipt-acceptance-preview":
		return runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptStorageRecordWriterAuthorizationReceiptAcceptancePreview(args[1:], stdout)
	case "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-accepted-receipt-gate-preview":
		return runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptStorageRecordWriterAcceptedReceiptGatePreview(args[1:], stdout)
	case "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-grant-preview":
		return runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptStorageRecordWriterGrantPreview(args[1:], stdout)
	case "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-enablement-preview":
		return runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptStorageRecordWriterEnablementPreview(args[1:], stdout)
	case "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-gate-preview":
		return runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptStorageRecordWriterCallGatePreview(args[1:], stdout)
	case "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-preview":
		return runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptStorageRecordWriterCallAuthorizationPreview(args[1:], stdout)
	case "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-receipt-preview":
		return runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptStorageRecordWriterCallAuthorizationReceiptPreview(args[1:], stdout)
	case "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-receipt-acceptance-preview":
		return runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptStorageRecordWriterCallAuthorizationReceiptAcceptancePreview(args[1:], stdout)
	case "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-receipt-accepted-receipt-gate-preview":
		return runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptStorageRecordWriterCallAuthorizationReceiptAcceptedReceiptGatePreview(args[1:], stdout)
	case "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-receipt-accepted-receipt-gate-authorization-preview":
		return runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptStorageRecordWriterCallAuthorizationReceiptAcceptedReceiptGateAuthorizationPreview(args[1:], stdout)
	case "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-receipt-accepted-receipt-gate-grant-preview":
		return runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptStorageRecordWriterCallAuthorizationReceiptAcceptedReceiptGateGrantPreview(args[1:], stdout)
	case "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-receipt-accepted-receipt-gate-enablement-preview":
		return runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptStorageRecordWriterCallAuthorizationReceiptAcceptedReceiptGateEnablementPreview(args[1:], stdout)
	case "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-receipt-accepted-receipt-gate-call-gate-preview":
		return runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptStorageRecordWriterCallAuthorizationReceiptAcceptedReceiptGateCallGatePreview(args[1:], stdout)
	case "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-receipt-accepted-receipt-gate-call-authorization-preview":
		return runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptStorageRecordWriterCallAuthorizationReceiptAcceptedReceiptGateCallAuthorizationPreview(args[1:], stdout)
	case "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-receipt-accepted-receipt-gate-call-authorization-receipt-preview":
		return runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptStorageRecordWriterCallAuthorizationReceiptAcceptedReceiptGateCallAuthorizationReceiptPreview(args[1:], stdout)
	case "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-receipt-accepted-receipt-gate-call-authorization-receipt-acceptance-preview":
		return runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptStorageRecordWriterCallAuthorizationReceiptAcceptedReceiptGateCallAuthorizationReceiptAcceptancePreview(args[1:], stdout)
	case "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-receipt-accepted-receipt-gate-call-authorization-receipt-accepted-receipt-gate-preview":
		return runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptStorageRecordWriterCallAuthorizationReceiptAcceptedReceiptGateCallAuthorizationReceiptAcceptedReceiptGatePreview(args[1:], stdout)
	case "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-receipt-accepted-receipt-gate-call-authorization-receipt-accepted-receipt-gate-authorization-preview":
		return runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptStorageRecordWriterCallAuthorizationReceiptAcceptedReceiptGateCallAuthorizationReceiptAcceptedReceiptGateAuthorizationPreview(args[1:], stdout)
	case "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-receipt-accepted-receipt-gate-call-authorization-receipt-accepted-receipt-gate-grant-preview":
		return runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptStorageRecordWriterCallAuthorizationReceiptAcceptedReceiptGateCallAuthorizationReceiptAcceptedReceiptGateGrantPreview(args[1:], stdout)
	case "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-receipt-accepted-receipt-gate-call-authorization-receipt-accepted-receipt-gate-enablement-preview":
		return runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptStorageRecordWriterCallAuthorizationReceiptAcceptedReceiptGateCallAuthorizationReceiptAcceptedReceiptGateEnablementPreview(args[1:], stdout)
	case "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-receipt-accepted-receipt-gate-call-authorization-receipt-accepted-receipt-gate-call-gate-preview":
		return runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptStorageRecordWriterCallAuthorizationReceiptAcceptedReceiptGateCallAuthorizationReceiptAcceptedReceiptGateCallGatePreview(args[1:], stdout)
	case "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-receipt-accepted-receipt-gate-call-authorization-receipt-accepted-receipt-gate-call-preview":
		return runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptStorageRecordWriterCallAuthorizationReceiptAcceptedReceiptGateCallAuthorizationReceiptAcceptedReceiptGateCallPreview(args[1:], stdout)
	case "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-receipt-accepted-receipt-gate-call-authorization-receipt-accepted-receipt-gate-call-receipt-preview":
		return runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptStorageRecordWriterCallAuthorizationReceiptAcceptedReceiptGateCallAuthorizationReceiptAcceptedReceiptGateCallReceiptPreview(args[1:], stdout)
	case "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-receipt-accepted-receipt-gate-call-authorization-receipt-accepted-receipt-gate-call-receipt-acceptance-preview":
		return runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptStorageRecordWriterCallAuthorizationReceiptAcceptedReceiptGateCallAuthorizationReceiptAcceptedReceiptGateCallReceiptAcceptancePreview(args[1:], stdout)
	case "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-receipt-accepted-receipt-gate-call-authorization-receipt-accepted-receipt-gate-call-receipt-consumption-gate-preview":
		return runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptStorageRecordWriterCallAuthorizationReceiptAcceptedReceiptGateCallAuthorizationReceiptAcceptedReceiptGateCallReceiptConsumptionGatePreview(args[1:], stdout)
	case "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-receipt-accepted-receipt-gate-call-authorization-receipt-accepted-receipt-gate-call-receipt-consumer-authorization-preview":
		return runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptStorageRecordWriterCallAuthorizationReceiptAcceptedReceiptGateCallAuthorizationReceiptAcceptedReceiptGateCallReceiptConsumerAuthorizationPreview(args[1:], stdout)
	case "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-receipt-accepted-receipt-gate-call-authorization-receipt-accepted-receipt-gate-call-receipt-consumer-enablement-preview":
		return runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptStorageRecordWriterCallAuthorizationReceiptAcceptedReceiptGateCallAuthorizationReceiptAcceptedReceiptGateCallReceiptConsumerEnablementPreview(args[1:], stdout)
	case "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-receipt-accepted-receipt-gate-call-authorization-receipt-accepted-receipt-gate-call-receipt-consumer-enablement-gate-preview":
		return runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptStorageRecordWriterCallAuthorizationReceiptAcceptedReceiptGateCallAuthorizationReceiptAcceptedReceiptGateCallReceiptConsumerEnablementGatePreview(args[1:], stdout)
	case "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-receipt-accepted-receipt-gate-call-authorization-receipt-accepted-receipt-gate-call-receipt-consumer-enablement-receipt-preview":
		return runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptStorageRecordWriterCallAuthorizationReceiptAcceptedReceiptGateCallAuthorizationReceiptAcceptedReceiptGateCallReceiptConsumerEnablementReceiptPreview(args[1:], stdout)
	case "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-receipt-accepted-receipt-gate-call-authorization-receipt-accepted-receipt-gate-call-receipt-consumer-enablement-receipt-persistence-authorization-preview":
		return runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptStorageRecordWriterCallAuthorizationReceiptAcceptedReceiptGateCallAuthorizationReceiptAcceptedReceiptGateCallReceiptConsumerEnablementReceiptPersistenceAuthorizationPreview(args[1:], stdout)
	case "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-receipt-accepted-receipt-gate-call-authorization-receipt-accepted-receipt-gate-call-receipt-consumer-enablement-receipt-closed-persistence-implementation-preview":
		return runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptStorageRecordWriterCallAuthorizationReceiptAcceptedReceiptGateCallAuthorizationReceiptAcceptedReceiptGateCallReceiptConsumerEnablementReceiptClosedPersistenceImplementationPreview(args[1:], stdout)
	case "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-receipt-accepted-receipt-gate-call-authorization-receipt-accepted-receipt-gate-call-receipt-consumer-enablement-receipt-storage-persistence-gate-audit-preview":
		return runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptStorageRecordWriterCallAuthorizationReceiptAcceptedReceiptGateCallAuthorizationReceiptAcceptedReceiptGateCallReceiptConsumerEnablementReceiptStoragePersistenceGateAuditPreview(args[1:], stdout)
	case "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-receipt-accepted-receipt-gate-call-authorization-receipt-accepted-receipt-gate-call-receipt-consumer-enablement-receipt-storage-record-contract-preview":
		return runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptStorageRecordWriterCallAuthorizationReceiptAcceptedReceiptGateCallAuthorizationReceiptAcceptedReceiptGateCallReceiptConsumerEnablementReceiptStorageRecordContractPreview(args[1:], stdout)
	case "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-receipt-accepted-receipt-gate-call-authorization-receipt-accepted-receipt-gate-call-receipt-consumer-enablement-receipt-closed-storage-record-writer-implementation-preview":
		return runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptStorageRecordWriterCallAuthorizationReceiptAcceptedReceiptGateCallAuthorizationReceiptAcceptedReceiptGateCallReceiptConsumerEnablementReceiptClosedStorageRecordWriterImplementationPreview(args[1:], stdout)
	case "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-receipt-accepted-receipt-gate-call-authorization-receipt-accepted-receipt-gate-call-receipt-consumer-enablement-receipt-storage-record-writer-authorization-preview":
		return runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptStorageRecordWriterCallAuthorizationReceiptAcceptedReceiptGateCallAuthorizationReceiptAcceptedReceiptGateCallReceiptConsumerEnablementReceiptStorageRecordWriterAuthorizationPreview(args[1:], stdout)
	case "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-receipt-accepted-receipt-gate-call-authorization-receipt-accepted-receipt-gate-call-receipt-consumer-enablement-receipt-storage-record-writer-authorization-receipt-preview":
		return runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptStorageRecordWriterCallAuthorizationReceiptAcceptedReceiptGateCallAuthorizationReceiptAcceptedReceiptGateCallReceiptConsumerEnablementReceiptStorageRecordWriterAuthorizationReceiptPreview(args[1:], stdout)
	case "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-receipt-accepted-receipt-gate-call-authorization-receipt-accepted-receipt-gate-call-receipt-consumer-enablement-receipt-storage-record-writer-authorization-receipt-acceptance-preview":
		return runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptStorageRecordWriterCallAuthorizationReceiptAcceptedReceiptGateCallAuthorizationReceiptAcceptedReceiptGateCallReceiptConsumerEnablementReceiptStorageRecordWriterAuthorizationReceiptAcceptancePreview(args[1:], stdout)
	case "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-receipt-accepted-receipt-gate-call-authorization-receipt-accepted-receipt-gate-call-receipt-consumer-enablement-receipt-storage-record-writer-accepted-receipt-gate-preview":
		return runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptStorageRecordWriterCallAuthorizationReceiptAcceptedReceiptGateCallAuthorizationReceiptAcceptedReceiptGateCallReceiptConsumerEnablementReceiptStorageRecordWriterAcceptedReceiptGatePreview(args[1:], stdout)
	case "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-receipt-accepted-receipt-gate-call-authorization-receipt-accepted-receipt-gate-call-receipt-consumer-enablement-receipt-storage-record-writer-grant-preview":
		return runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptStorageRecordWriterCallAuthorizationReceiptAcceptedReceiptGateCallAuthorizationReceiptAcceptedReceiptGateCallReceiptConsumerEnablementReceiptStorageRecordWriterGrantPreview(args[1:], stdout)
	case "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-receipt-accepted-receipt-gate-call-authorization-receipt-accepted-receipt-gate-call-receipt-consumer-enablement-receipt-storage-record-writer-enablement-preview":
		return runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptStorageRecordWriterCallAuthorizationReceiptAcceptedReceiptGateCallAuthorizationReceiptAcceptedReceiptGateCallReceiptConsumerEnablementReceiptStorageRecordWriterEnablementPreview(args[1:], stdout)
	case "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-receipt-accepted-receipt-gate-call-authorization-receipt-accepted-receipt-gate-call-receipt-consumer-enablement-receipt-storage-record-writer-call-gate-preview":
		return runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptStorageRecordWriterCallAuthorizationReceiptAcceptedReceiptGateCallAuthorizationReceiptAcceptedReceiptGateCallReceiptConsumerEnablementReceiptStorageRecordWriterCallGatePreview(args[1:], stdout)
	case "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-delivery-explicit-consent-audit-preview":
		return runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationDeliveryExplicitConsentAuditPreview(args[1:], stdout)
	case "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-delivery-consent-collection-gate-audit-preview":
		return runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationDeliveryConsentCollectionGateAuditPreview(args[1:], stdout)
	case "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-delivery-consent-collection-receipt-audit-preview":
		return runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationDeliveryConsentCollectionReceiptAuditPreview(args[1:], stdout)
	case "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-delivery-consent-collection-receipt-acceptance-audit-preview":
		return runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationDeliveryConsentCollectionReceiptAcceptanceAuditPreview(args[1:], stdout)
	case "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-delivery-consent-collection-receipt-consumption-gate-audit-preview":
		return runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationDeliveryConsentCollectionReceiptConsumptionGateAuditPreview(args[1:], stdout)
	case "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-delivery-consent-collection-receipt-consumer-authorization-audit-preview":
		return runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationDeliveryConsentCollectionReceiptConsumerAuthorizationAuditPreview(args[1:], stdout)
	case "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-delivery-consent-collection-receipt-consumer-enablement-gate-audit-preview":
		return runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationDeliveryConsentCollectionReceiptConsumerEnablementGateAuditPreview(args[1:], stdout)
	case "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-delivery-execution-authorization-audit-preview":
		return runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationDeliveryExecutionAuthorizationAuditPreview(args[1:], stdout)
	case "production-receipt-notification-delivery-gate-audit-preview":
		return runProductionReceiptNotificationDeliveryGateAuditPreview(args[1:], stdout)
	case "production-receipt-revocation-visibility-audit-preview":
		return runProductionReceiptRevocationVisibilityAuditPreview(args[1:], stdout)
	case "production-receipt-writer-authorization-review-preview":
		return runProductionReceiptWriterAuthorizationReviewPreview(args[1:], stdout)
	case "production-rollback-diagnostics-review-preview":
		return runProductionRollbackDiagnosticsReviewPreview(args[1:], stdout)
	case "recipe-conflict-audit-preview":
		return runRecipeConflictAuditPreview(args[1:], stdout)
	case "repair-plan-preview":
		return runRepairPlanPreview(args[1:], stdout)
	case "restricted-owner-smoke-receipt-fanout-preview":
		return runRestrictedOwnerSmokeReceiptFanOutPreview(args[1:], stdout)
	case "restricted-owner-smoke-receipt-fanout-owner-route-audit-preview":
		return runRestrictedOwnerSmokeReceiptFanOutOwnerRouteAuditPreview(args[1:], stdout)
	case "restricted-owner-smoke-receipt-fanout-owner-route-preview":
		return runRestrictedOwnerSmokeReceiptFanOutOwnerRoutePreview(args[1:], stdout)
	case "restricted-owner-smoke-receipt-lookup-preview":
		return runRestrictedOwnerSmokeReceiptLookupPreview(args[1:], stdout)
	case "restricted-owner-smoke-receipt-record":
		return runRestrictedOwnerSmokeReceiptRecord(args[1:], stdout)
	case "restricted-product-smoke-packet-preview":
		return runRestrictedProductSmokePacketPreview(args[1:], stdout)
	case "review-flow-preview":
		return runReviewFlowPreview(args[1:], stdout)
	case "run-plan-preview":
		return runRunPlanPreview(args[1:], stdout)
	case "runtime-live-owner-gate-preview":
		return runRuntimeLiveOwnerGatePreview(args[1:], stdout)
	case "runtime-method-parity-manifest-preview":
		return runRuntimeMethodParityManifestPreview(args[1:], stdout)
	case "runtime-owner-process-preview":
		return runRuntimeOwnerProcessPreview(args[1:], stdout)
	case "runtime-owner-readiness-preview":
		return runRuntimeOwnerReadinessPreview(args[1:], stdout)
	case "runtime-owner-recipe-trust-preview":
		return runRuntimeOwnerRecipeTrustPreview(args[1:], stdout)
	case "runtime-owner-route-manifest-preview":
		return runRuntimeOwnerRouteManifestPreview(args[1:], stdout)
	case "runtime-owner-smoke-plan-preview":
		return runRuntimeOwnerSmokePlanPreview(args[1:], stdout)
	case "runtime-policy-explanation-cards-preview":
		return runRuntimePolicyExplanationCardsPreview(args[1:], stdout)
	case "runtime-route-convergence-preview":
		return runRuntimeRouteConvergencePreview(args[1:], stdout)
	case "runtime-service-activation-preflight-preview":
		return runRuntimeServiceActivationPreflightPreview(args[1:], stdout)
	case "runtime-service-binding-preview":
		return runRuntimeServiceBindingPreview(args[1:], stdout)
	case "runtime-write-gate-preview":
		return runRuntimeWriteGatePreview(args[1:], stdout)
	case "settings-change-preview":
		return runSettingsChangePreview(args[1:], stdout)
	case "settings-preview":
		return runSettingsPreview(args[1:], stdout)
	case "settings-profile-migration-preview":
		return runSettingsProfileMigrationPreview(args[1:], stdout)
	case "signed-recipe-verifier-preview":
		return runSignedRecipeVerifierPreview(args[1:], stdout)
	case "snapshot-restore-candidates-preview":
		return runSnapshotRestoreCandidatesPreview(args[1:], stdout)
	case "snapshot-plan-preview":
		return runSnapshotPlanPreview(args[1:], stdout)
	case "state-root-preview":
		return runStateRootPreview(args[1:], stdout)
	case "state-root-quota-retention-preview":
		return runStateRootQuotaRetentionPreview(args[1:], stdout)
	case "support-bundle-manifest-preview":
		return runSupportBundleManifestPreview(args[1:], stdout)
	case "support-case-timeline-preview":
		return runSupportCaseTimelinePreview(args[1:], stdout)
	case "task-manager-identity-preview":
		return runTaskManagerIdentityPreview(args[1:], stdout)
	case "test-plan-preview":
		return runTestPlanPreview(args[1:], stdout)
	case "test-result-preview":
		return runTestResultPreview(args[1:], stdout)
	case "tray-status-preview":
		return runTrayStatusPreview(args[1:], stdout)
	case "window-identity-preview":
		return runWindowIdentityPreview(args[1:], stdout)
	case "windows-app-run-smoke":
		return runWindowsAppRunSmoke(args[1:], stdout)
	case "windows-app-launch-profile":
		return runWindowsAppLaunchProfile(args[1:], stdout)
	case "windows-app-runner-diagnostics":
		return runWindowsAppRunnerDiagnostics(args[1:], stdout)
	case "windows-app-smoke-profile-preflight":
		return runWindowsAppSmokeProfilePreflight(args[1:], stdout)
	case "windows-app-smoke-profile-render":
		return runWindowsAppSmokeProfileRender(args[1:], stdout)
	case "windows-app-launcher-bundle-record":
		return runWindowsAppLauncherBundleRecord(args[1:], stdout)
	case "windows-app-container-run-smoke":
		return runWindowsAppContainerRunSmoke(args[1:], stdout)
	case "windows-app-container-x-gui-smoke":
		return runWindowsAppContainerXGUISmoke(args[1:], stdout)
	case "windows-app-guest-wine-smoke":
		return runWindowsAppGuestWineSmoke(args[1:], stdout)
	case "windows-app-guest-wine-gui-smoke":
		return runWindowsAppGuestWineGUISmoke(args[1:], stdout)
	case "windows-known-app-fetch":
		return runWindowsKnownAppFetch(args[1:], stdout)
	case "windows-known-app-guest-wine-smoke":
		return runWindowsKnownAppGuestWineSmoke(args[1:], stdout)
	case "windows-known-app-launch-profile-materialize":
		return runWindowsKnownAppLaunchProfileMaterialize(args[1:], stdout)
	case "windows-known-app-prepare-launch-profile":
		return runWindowsKnownAppPrepareLaunchProfile(args[1:], stdout)
	case "windows-known-app-prepare-and-launch-profile":
		return runWindowsKnownAppPrepareAndLaunchProfile(args[1:], stdout)
	case "windows-known-app-run":
		return runWindowsKnownAppRun(args[1:], stdout)
	case "windows-known-app-managed-launch-preview":
		return runWindowsKnownAppManagedLaunchPreview(args[1:], stdout)
	case "windows-known-app-kde-launcher-preview":
		return runWindowsKnownAppKDELauncherPreview(args[1:], stdout)
	case "windows-known-app-launch-request-preview":
		return runWindowsKnownAppLaunchRequestPreview(args[1:], stdout)
	case "windows-known-app-dispatch-preview":
		return runWindowsKnownAppDispatchPreview(args[1:], stdout)
	case "windows-known-app-dispatch-smoke":
		return runWindowsKnownAppDispatchSmoke(args[1:], stdout)
	case "windows-known-app-launch-bridge-preview":
		return runWindowsKnownAppLaunchBridgePreview(args[1:], stdout)
	case "windows-compatibility-workstreams-preview":
		return runWindowsCompatibilityWorkstreamsPreview(args[1:], stdout)
	default:
		return fmt.Errorf("unknown command: %s", args[0])
	}
}

func runApplicationsPreview(args []string, stdout io.Writer) error {
	recipes, provenance, err := parseRegistryPreviewSource("applications-preview", args)
	if err != nil {
		return err
	}
	preview, err := appidentity.NewApplicationsPreview(recipes, provenance)
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(preview)
}

func runApplicationPreview(args []string, stdout io.Writer) error {
	recipe, provenance, err := parseRecipeSource("application-preview", args)
	if err != nil {
		return err
	}
	preview, err := appidentity.NewApplicationPreview(recipe, provenance)
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(preview)
}

func runEngineCatalogPreview(args []string, stdout io.Writer) error {
	if len(args) != 0 {
		return errors.New("engine-catalog-preview does not accept arguments")
	}
	preview, err := appidentity.NewEngineCatalogPreview()
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(preview)
}

func runCompatibilityCenterPreview(args []string, stdout io.Writer) error {
	recipes, provenance, options, err := parseCompatibilityCenterPreviewSource(args)
	if err != nil {
		return err
	}
	preview, err := appidentity.NewCompatibilityCenterPreviewWithOptions(recipes, provenance, options)
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(preview)
}

func runBackendSelectionPreview(args []string, stdout io.Writer) error {
	recipe, provenance, err := parseRecipeSource("backend-selection-preview", args)
	if err != nil {
		return err
	}
	plan, err := appidentity.NewPlanWithProvenance(recipe, provenance)
	if err != nil {
		return err
	}
	preview, err := plan.BackendSelectionPreview()
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(preview)
}

func runBackendEnvironmentPreview(args []string, stdout io.Writer) error {
	recipe, provenance, err := parseRecipeSource("backend-environment-preview", args)
	if err != nil {
		return err
	}
	plan, err := appidentity.NewPlanWithProvenance(recipe, provenance)
	if err != nil {
		return err
	}
	preview, err := plan.BackendEnvironmentPreview()
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(preview)
}

func runBackendBindingPreview(args []string, stdout io.Writer) error {
	recipe, provenance, err := parseRecipeSource("backend-binding-preview", args)
	if err != nil {
		return err
	}
	plan, err := appidentity.NewPlanWithProvenance(recipe, provenance)
	if err != nil {
		return err
	}
	preview, err := plan.BackendBindingPreview()
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(preview)
}

func runRepairPlanPreview(args []string, stdout io.Writer) error {
	recipe, provenance, issue, err := parseRepairPlanPreviewSource(args)
	if err != nil {
		return err
	}
	plan, err := appidentity.NewPlanWithProvenance(recipe, provenance)
	if err != nil {
		return err
	}
	preview, err := plan.RepairPlanPreview(issue)
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(preview)
}

func runTestPlanPreview(args []string, stdout io.Writer) error {
	recipe, provenance, testType, err := parseTestPreviewSource("test-plan-preview", args)
	if err != nil {
		return err
	}
	plan, err := appidentity.NewPlanWithProvenance(recipe, provenance)
	if err != nil {
		return err
	}
	preview, err := plan.TestPlanPreview(testType)
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(preview)
}

func runTestResultPreview(args []string, stdout io.Writer) error {
	recipe, provenance, testType, err := parseTestPreviewSource("test-result-preview", args)
	if err != nil {
		return err
	}
	plan, err := appidentity.NewPlanWithProvenance(recipe, provenance)
	if err != nil {
		return err
	}
	preview, err := plan.TestResultPreview(testType)
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(preview)
}

func parseRepairPlanPreviewSource(args []string) (appidentity.Recipe, appidentity.Provenance, string, error) {
	flags := flag.NewFlagSet("repair-plan-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	applicationID := flags.String("app", "", "application id to load from the recipe registry")
	registryPath := flags.String("registry", "", "path to an Xnix recipe registry JSON file")
	recipeRoot := flags.String("recipe-root", "", "directory containing registered recipe files; defaults to the registry directory")
	recipePath := flags.String("recipe", "", "path to an Xnix application recipe JSON file")
	issue := flags.String("issue", "", "repair issue type")
	if err := flags.Parse(args); err != nil {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", err
	}
	if *issue == "" {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", errors.New("repair-plan-preview requires --issue")
	}
	if (*recipePath == "") == (*registryPath == "") {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", errors.New("repair-plan-preview requires exactly one source: --recipe or --registry")
	}
	if *registryPath != "" && *applicationID == "" {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", errors.New("repair-plan-preview requires --app when --registry is used")
	}
	if *recipePath != "" && (*applicationID != "" || *recipeRoot != "") {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", errors.New("repair-plan-preview --recipe cannot be combined with --app or --recipe-root")
	}
	if flags.NArg() != 0 {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", errors.New("repair-plan-preview does not accept positional arguments")
	}

	recipe, provenance, err := loadRecipe(*recipePath, *registryPath, *recipeRoot, *applicationID)
	return recipe, provenance, *issue, err
}

func parseTestPreviewSource(command string, args []string) (appidentity.Recipe, appidentity.Provenance, string, error) {
	flags := flag.NewFlagSet(command, flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	applicationID := flags.String("app", "", "application id to load from the recipe registry")
	registryPath := flags.String("registry", "", "path to an Xnix recipe registry JSON file")
	recipeRoot := flags.String("recipe-root", "", "directory containing registered recipe files; defaults to the registry directory")
	recipePath := flags.String("recipe", "", "path to an Xnix application recipe JSON file")
	testType := flags.String("test-type", "preflight", "test type: preflight, smoke, or repair-readiness")
	if err := flags.Parse(args); err != nil {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", err
	}
	if (*recipePath == "") == (*registryPath == "") {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", errors.New(command + " requires exactly one source: --recipe or --registry")
	}
	if *registryPath != "" && *applicationID == "" {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", errors.New(command + " requires --app when --registry is used")
	}
	if *recipePath != "" && (*applicationID != "" || *recipeRoot != "") {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", errors.New(command + " --recipe cannot be combined with --app or --recipe-root")
	}
	if flags.NArg() != 0 {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", errors.New(command + " does not accept positional arguments")
	}

	recipe, provenance, err := loadRecipe(*recipePath, *registryPath, *recipeRoot, *applicationID)
	return recipe, provenance, *testType, err
}

func runRunPlanPreview(args []string, stdout io.Writer) error {
	recipe, provenance, err := parseRecipeSource("run-plan-preview", args)
	if err != nil {
		return err
	}
	plan, err := appidentity.NewPlanWithProvenance(recipe, provenance)
	if err != nil {
		return err
	}
	preview, err := plan.RunPlanPreview()
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(preview)
}

func runPackageSourcePreview(args []string, stdout io.Writer) error {
	recipe, provenance, err := parseRecipeSource("package-source-preview", args)
	if err != nil {
		return err
	}
	plan, err := appidentity.NewPlanWithProvenance(recipe, provenance)
	if err != nil {
		return err
	}
	preview, err := plan.PackageSourcePreview()
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(preview)
}

func runAcquisitionPreflightPreview(args []string, stdout io.Writer) error {
	recipe, provenance, err := parseRecipeSource("acquisition-preflight-preview", args)
	if err != nil {
		return err
	}
	plan, err := appidentity.NewPlanWithProvenance(recipe, provenance)
	if err != nil {
		return err
	}
	preview, err := plan.AcquisitionPreflightPreview()
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(preview)
}

func runBackendCapabilityMatrixPreview(args []string, stdout io.Writer) error {
	if len(args) != 0 {
		return errors.New("backend-capability-matrix-preview does not accept arguments")
	}
	preview, err := appidentity.NewBackendCapabilityMatrixPreview()
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(preview)
}

func runBackendManagerPreview(args []string, stdout io.Writer) error {
	if len(args) != 0 {
		return errors.New("backend-manager-preview does not accept arguments")
	}
	preview := appidentity.NewBackendManagerPreview()

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(preview)
}

func runBackendManagerRecord(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("backend-manager-record", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	stateRoot := flags.String("state-root", "", "Runtime state root for backend manager inventory")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if *stateRoot == "" {
		return errors.New("backend-manager-record requires --state-root")
	}
	if flags.NArg() != 0 {
		return errors.New("backend-manager-record does not accept positional arguments")
	}
	record, err := appidentity.RecordBackendManagerPreview(*stateRoot)
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(record)
}

func runBackendLifecyclePreview(args []string, stdout io.Writer) error {
	recipe, provenance, stateRoot, err := parseBackendLifecyclePreviewSource(args)
	if err != nil {
		return err
	}
	plan, err := appidentity.NewPlanWithProvenance(recipe, provenance)
	if err != nil {
		return err
	}
	var preview appidentity.BackendLifecyclePreview
	if stateRoot == "" {
		preview, err = plan.BackendLifecyclePreview()
	} else {
		preview, err = plan.BackendLifecyclePreviewWithStateRoot(stateRoot)
	}
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(preview)
}

func runBackendLifecycleRecord(args []string, stdout io.Writer) error {
	recipe, provenance, stateRoot, action, gate, repairHint, blockReason, err := parseBackendLifecycleRecordSource(args)
	if err != nil {
		return err
	}
	plan, err := appidentity.NewPlanWithProvenance(recipe, provenance)
	if err != nil {
		return err
	}
	record, err := plan.RecordBackendLifecycleState(stateRoot, action, gate, repairHint, blockReason)
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(record)
}

func parseBackendLifecyclePreviewSource(args []string) (appidentity.Recipe, appidentity.Provenance, string, error) {
	flags := flag.NewFlagSet("backend-lifecycle-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	applicationID := flags.String("app", "", "application id to load from the recipe registry")
	registryPath := flags.String("registry", "", "path to an Xnix recipe registry JSON file")
	recipeRoot := flags.String("recipe-root", "", "directory containing registered recipe files; defaults to the registry directory")
	recipePath := flags.String("recipe", "", "path to an Xnix application recipe JSON file")
	stateRoot := flags.String("state-root", "", "optional Runtime state root containing environment lifecycle records")
	if err := flags.Parse(args); err != nil {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", err
	}
	if (*recipePath == "") == (*registryPath == "") {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", errors.New("backend-lifecycle-preview requires exactly one source: --recipe or --registry")
	}
	if *registryPath != "" && *applicationID == "" {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", errors.New("backend-lifecycle-preview requires --app when --registry is used")
	}
	if *recipePath != "" && (*applicationID != "" || *recipeRoot != "") {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", errors.New("backend-lifecycle-preview --recipe cannot be combined with --app or --recipe-root")
	}
	if flags.NArg() != 0 {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", errors.New("backend-lifecycle-preview does not accept positional arguments")
	}

	recipe, provenance, err := loadRecipe(*recipePath, *registryPath, *recipeRoot, *applicationID)
	return recipe, provenance, *stateRoot, err
}

func parseBackendLifecycleRecordSource(args []string) (appidentity.Recipe, appidentity.Provenance, string, string, string, string, string, error) {
	flags := flag.NewFlagSet("backend-lifecycle-record", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	applicationID := flags.String("app", "", "application id to load from the recipe registry")
	registryPath := flags.String("registry", "", "path to an Xnix recipe registry JSON file")
	recipeRoot := flags.String("recipe-root", "", "directory containing registered recipe files; defaults to the registry directory")
	recipePath := flags.String("recipe", "", "path to an Xnix application recipe JSON file")
	stateRoot := flags.String("state-root", "", "Runtime state root for backend lifecycle records")
	action := flags.String("action", "", "backend lifecycle action: inspect, plan, stage, satisfy-gate, mark-ready, flag-repair, block, or retire")
	gate := flags.String("gate", "", "required gate to satisfy when --action=satisfy-gate")
	repairHint := flags.String("repair-hint", "", "repair hint when --action=flag-repair")
	blockReason := flags.String("block-reason", "", "block reason when --action=block")
	if err := flags.Parse(args); err != nil {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", "", "", "", "", err
	}
	if (*recipePath == "") == (*registryPath == "") {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", "", "", "", "", errors.New("backend-lifecycle-record requires exactly one source: --recipe or --registry")
	}
	if *registryPath != "" && *applicationID == "" {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", "", "", "", "", errors.New("backend-lifecycle-record requires --app when --registry is used")
	}
	if *recipePath != "" && (*applicationID != "" || *recipeRoot != "") {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", "", "", "", "", errors.New("backend-lifecycle-record --recipe cannot be combined with --app or --recipe-root")
	}
	if *stateRoot == "" {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", "", "", "", "", errors.New("backend-lifecycle-record requires --state-root")
	}
	if *action == "" {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", "", "", "", "", errors.New("backend-lifecycle-record requires --action")
	}
	if flags.NArg() != 0 {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", "", "", "", "", errors.New("backend-lifecycle-record does not accept positional arguments")
	}

	recipe, provenance, err := loadRecipe(*recipePath, *registryPath, *recipeRoot, *applicationID)
	return recipe, provenance, *stateRoot, *action, *gate, *repairHint, *blockReason, err
}

func runArtifactManifestPreview(args []string, stdout io.Writer) error {
	recipe, provenance, err := parseRecipeSource("artifact-manifest-preview", args)
	if err != nil {
		return err
	}
	plan, err := appidentity.NewPlanWithProvenance(recipe, provenance)
	if err != nil {
		return err
	}
	preview, err := plan.ArtifactManifestPreview()
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(preview)
}

func runDiagnosticsPreview(args []string, stdout io.Writer) error {
	recipe, provenance, err := parseRecipeSource("diagnostics-preview", args)
	if err != nil {
		return err
	}
	preview, err := appidentity.NewDiagnosticsPreview(recipe, provenance)
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(preview)
}

func runExecutionReadinessPreview(args []string, stdout io.Writer) error {
	recipe, provenance, err := parseRecipeSource("execution-readiness-preview", args)
	if err != nil {
		return err
	}
	plan, err := appidentity.NewPlanWithProvenance(recipe, provenance)
	if err != nil {
		return err
	}
	preview, err := plan.ExecutionReadinessPreview()
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(preview)
}

func runLaunchIntentPreview(args []string, stdout io.Writer) error {
	recipe, provenance, fileURIs, err := parseLaunchIntentPreviewSource(args)
	if err != nil {
		return err
	}
	plan, err := appidentity.NewPlanWithProvenance(recipe, provenance)
	if err != nil {
		return err
	}
	preview, err := plan.LaunchIntentPreview(fileURIs)
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(preview)
}

func runExecutionRequestPreview(args []string, stdout io.Writer) error {
	recipe, provenance, fileURIs, err := parseExecutionRequestPreviewSource(args)
	if err != nil {
		return err
	}
	plan, err := appidentity.NewPlanWithProvenance(recipe, provenance)
	if err != nil {
		return err
	}
	preview, err := plan.ExecutionRequestPreview(fileURIs)
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(preview)
}

func runExecutionReviewPreview(args []string, stdout io.Writer) error {
	recipe, provenance, fileURIs, err := parseExecutionReviewPreviewSource(args)
	if err != nil {
		return err
	}
	plan, err := appidentity.NewPlanWithProvenance(recipe, provenance)
	if err != nil {
		return err
	}
	preview, err := plan.ExecutionReviewPreview(fileURIs)
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(preview)
}

func runExecutionDecisionPreview(args []string, stdout io.Writer) error {
	recipe, provenance, decision, fileURIs, err := parseExecutionDecisionPreviewSource(args)
	if err != nil {
		return err
	}
	plan, err := appidentity.NewPlanWithProvenance(recipe, provenance)
	if err != nil {
		return err
	}
	preview, err := plan.ExecutionDecisionPreview(decision, fileURIs)
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(preview)
}

func runExecutionPreflightPreview(args []string, stdout io.Writer) error {
	recipe, provenance, decision, fileURIs, err := parseExecutionPreflightPreviewSource(args)
	if err != nil {
		return err
	}
	plan, err := appidentity.NewPlanWithProvenance(recipe, provenance)
	if err != nil {
		return err
	}
	preview, err := plan.ExecutionPreflightPreview(decision, fileURIs)
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(preview)
}

func runExecutionResourceGrantPreview(args []string, stdout io.Writer) error {
	recipe, provenance, decision, fileURIs, err := parseExecutionResourceGrantPreviewSource(args)
	if err != nil {
		return err
	}
	plan, err := appidentity.NewPlanWithProvenance(recipe, provenance)
	if err != nil {
		return err
	}
	preview, err := plan.ExecutionResourceGrantPreview(decision, fileURIs)
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(preview)
}

func runExecutionTransactionPreview(args []string, stdout io.Writer) error {
	recipe, provenance, decision, fileURIs, err := parseExecutionTransactionPreviewSource(args)
	if err != nil {
		return err
	}
	plan, err := appidentity.NewPlanWithProvenance(recipe, provenance)
	if err != nil {
		return err
	}
	preview, err := plan.ExecutionTransactionPreview(decision, fileURIs)
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(preview)
}

func runExecutionSessionPreview(args []string, stdout io.Writer) error {
	recipe, provenance, decision, fileURIs, err := parseExecutionSessionPreviewSource(args)
	if err != nil {
		return err
	}
	plan, err := appidentity.NewPlanWithProvenance(recipe, provenance)
	if err != nil {
		return err
	}
	preview, err := plan.ExecutionSessionPreview(decision, fileURIs)
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(preview)
}

func runExecutionSessionStatusPreview(args []string, stdout io.Writer) error {
	recipe, provenance, decision, fileURIs, err := parseExecutionSessionStatusPreviewSource(args)
	if err != nil {
		return err
	}
	plan, err := appidentity.NewPlanWithProvenance(recipe, provenance)
	if err != nil {
		return err
	}
	preview, err := plan.ExecutionSessionStatusPreview(decision, fileURIs)
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(preview)
}

func runFileOpenPreview(args []string, stdout io.Writer) error {
	recipes, provenance, applicationID, fileURIs, options, err := parseFileOpenPreviewSource(args)
	if err != nil {
		return err
	}
	preview, err := appidentity.NewFileOpenPreviewWithOptions(recipes, provenance, fileURIs, applicationID, options)
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(preview)
}

func runDolphinDropPreview(args []string, stdout io.Writer) error {
	recipes, provenance, applicationID, fileURIs, options, err := parseDolphinDropPreviewSource(args)
	if err != nil {
		return err
	}
	preview, err := appidentity.NewDolphinDropPreviewWithOptions(recipes, provenance, fileURIs, applicationID, options)
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(preview)
}

func runDolphinAIAnalysisPreview(args []string, stdout io.Writer) error {
	recipes, provenance, applicationID, fileURIs, err := parseDolphinAIAnalysisPreviewSource(args)
	if err != nil {
		return err
	}
	preview, err := appidentity.NewDolphinAIAnalysisPreview(recipes, provenance, fileURIs, applicationID)
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(preview)
}

func runKDEEntryPointsPreview(args []string, stdout io.Writer) error {
	recipe, provenance, decision, fileURIs, err := parseKDEEntryPointsPreviewSource(args)
	if err != nil {
		return err
	}
	plan, err := appidentity.NewPlanWithProvenance(recipe, provenance)
	if err != nil {
		return err
	}
	preview, err := plan.KDEEntryPointsPreview(decision, fileURIs)
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(preview)
}

func runKDEEntryPointActionPreview(args []string, stdout io.Writer) error {
	recipe, provenance, entryPointID, decision, fileURIs, err := parseKDEEntryPointActionPreviewSource(args)
	if err != nil {
		return err
	}
	plan, err := appidentity.NewPlanWithProvenance(recipe, provenance)
	if err != nil {
		return err
	}
	preview, err := plan.KDEEntryPointActionPreview(entryPointID, decision, fileURIs)
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(preview)
}

func runKDEActionQueuePreview(args []string, stdout io.Writer) error {
	recipe, provenance, decision, fileURIs, err := parseKDEActionQueuePreviewSource(args)
	if err != nil {
		return err
	}
	plan, err := appidentity.NewPlanWithProvenance(recipe, provenance)
	if err != nil {
		return err
	}
	preview, err := plan.KDEActionQueuePreview(decision, fileURIs)
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(preview)
}

func runKDEActionCardDeckPreview(args []string, stdout io.Writer) error {
	recipe, provenance, decision, fileURIs, err := parseKDEActionCardDeckPreviewSource(args)
	if err != nil {
		return err
	}
	plan, err := appidentity.NewPlanWithProvenance(recipe, provenance)
	if err != nil {
		return err
	}
	preview, err := plan.KDEActionCardDeckPreview(decision, fileURIs)
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(preview)
}

func runKDEActionDependencyGraphPreview(args []string, stdout io.Writer) error {
	recipe, provenance, decision, fileURIs, err := parseKDEActionDependencyGraphPreviewSource(args)
	if err != nil {
		return err
	}
	plan, err := appidentity.NewPlanWithProvenance(recipe, provenance)
	if err != nil {
		return err
	}
	preview, err := plan.KDEActionDependencyGraphPreview(decision, fileURIs)
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(preview)
}

func runKDEJourneyEvidencePreview(args []string, stdout io.Writer) error {
	recipe, provenance, decision, fileURIs, options, err := parseKDEJourneyEvidencePreviewSource(args)
	if err != nil {
		return err
	}
	plan, err := appidentity.NewPlanWithProvenance(recipe, provenance)
	if err != nil {
		return err
	}
	preview, err := plan.KDEJourneyEvidencePreviewWithOptions(decision, fileURIs, options)
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(preview)
}

func runKDECenterPagePreview(args []string, stdout io.Writer) error {
	recipe, provenance, decision, fileURIs, options, err := parseKDECenterPagePreviewSource(args)
	if err != nil {
		return err
	}
	preview, err := appidentity.NewKDECenterPagePreviewWithOptions(recipe, provenance, decision, fileURIs, options)
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(preview)
}

func runKDECenterPageSectionsPreview(args []string, stdout io.Writer) error {
	recipe, provenance, decision, fileURIs, err := parseKDECenterPageSectionsPreviewSource(args)
	if err != nil {
		return err
	}
	preview, err := appidentity.NewKDECenterPageSectionsPreview(recipe, provenance, decision, fileURIs)
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(preview)
}

func runKDECenterPageSectionDetailPreview(args []string, stdout io.Writer) error {
	recipe, provenance, sectionID, decision, fileURIs, err := parseKDECenterPageSectionDetailPreviewSource(args)
	if err != nil {
		return err
	}
	preview, err := appidentity.NewKDECenterPageSectionDetailPreview(recipe, provenance, sectionID, decision, fileURIs)
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(preview)
}

func runKDEActionReviewPreview(args []string, stdout io.Writer) error {
	recipe, provenance, actionID, decision, fileURIs, err := parseKDEActionReviewPreviewSource(args)
	if err != nil {
		return err
	}
	plan, err := appidentity.NewPlanWithProvenance(recipe, provenance)
	if err != nil {
		return err
	}
	preview, err := plan.KDEActionReviewPreview(actionID, decision, fileURIs)
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(preview)
}

func runKDEActionPreflightPreview(args []string, stdout io.Writer) error {
	recipe, provenance, actionID, decision, fileURIs, err := parseKDEActionPreflightPreviewSource(args)
	if err != nil {
		return err
	}
	plan, err := appidentity.NewPlanWithProvenance(recipe, provenance)
	if err != nil {
		return err
	}
	preview, err := plan.KDEActionPreflightPreview(actionID, decision, fileURIs)
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(preview)
}

func runKDEActionReceiptPreview(args []string, stdout io.Writer) error {
	recipe, provenance, actionID, decision, fileURIs, err := parseKDEActionReceiptPreviewSource(args)
	if err != nil {
		return err
	}
	plan, err := appidentity.NewPlanWithProvenance(recipe, provenance)
	if err != nil {
		return err
	}
	preview, err := plan.KDEActionReceiptPreview(actionID, decision, fileURIs)
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(preview)
}

func runKDEActionStatusPreview(args []string, stdout io.Writer) error {
	recipe, provenance, actionID, decision, fileURIs, err := parseKDEActionStatusPreviewSource(args)
	if err != nil {
		return err
	}
	plan, err := appidentity.NewPlanWithProvenance(recipe, provenance)
	if err != nil {
		return err
	}
	preview, err := plan.KDEActionStatusPreview(actionID, decision, fileURIs)
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(preview)
}

func runKDEActionCardPreview(args []string, stdout io.Writer) error {
	recipe, provenance, actionID, decision, fileURIs, err := parseKDEActionCardPreviewSource(args)
	if err != nil {
		return err
	}
	plan, err := appidentity.NewPlanWithProvenance(recipe, provenance)
	if err != nil {
		return err
	}
	preview, err := plan.KDEActionCardPreview(actionID, decision, fileURIs)
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(preview)
}

func runKRunnerQueryPreview(args []string, stdout io.Writer) error {
	recipes, provenance, query, options, err := parseKRunnerQueryPreviewSource(args)
	if err != nil {
		return err
	}
	preview, err := appidentity.NewKRunnerQueryPreviewWithOptions(recipes, provenance, query, options)
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(preview)
}

func runDesktopActivationBundlePreview(args []string, stdout io.Writer) error {
	recipe, provenance, err := parseRecipeSource("desktop-activation-bundle-preview", args)
	if err != nil {
		return err
	}
	plan, err := appidentity.NewPlanWithProvenance(recipe, provenance)
	if err != nil {
		return err
	}
	preview, err := plan.DesktopActivationBundlePreview()
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(preview)
}

func runDesktopActivationPreflightPreview(args []string, stdout io.Writer) error {
	recipe, provenance, mode, err := parseDesktopActivationPreflightPreviewSource(args)
	if err != nil {
		return err
	}
	plan, err := appidentity.NewPlanWithProvenance(recipe, provenance)
	if err != nil {
		return err
	}
	preview, err := plan.DesktopActivationPreflightPreview(mode)
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(preview)
}

func runDesktopActivationStagingPreview(args []string, stdout io.Writer) error {
	recipe, provenance, mode, err := parseDesktopActivationStagingPreviewSource(args)
	if err != nil {
		return err
	}
	plan, err := appidentity.NewPlanWithProvenance(recipe, provenance)
	if err != nil {
		return err
	}
	preview, err := plan.DesktopActivationStagingPreview(mode)
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(preview)
}

func runDesktopActivationTransactionPreview(args []string, stdout io.Writer) error {
	recipe, provenance, mode, err := parseDesktopActivationTransactionPreviewSource(args)
	if err != nil {
		return err
	}
	plan, err := appidentity.NewPlanWithProvenance(recipe, provenance)
	if err != nil {
		return err
	}
	preview, err := plan.DesktopActivationTransactionPreview(mode)
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(preview)
}

func runDesktopActivationStatusPreview(args []string, stdout io.Writer) error {
	recipe, provenance, mode, activationRoot, err := parseDesktopActivationStatusPreviewSource(args)
	if err != nil {
		return err
	}
	plan, err := appidentity.NewPlanWithProvenance(recipe, provenance)
	if err != nil {
		return err
	}
	preview, err := plan.DesktopActivationStatusPreviewWithReceipt(activationRoot, mode)
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(preview)
}

func runDesktopIdentityPlan(args []string, stdout io.Writer) error {
	recipe, provenance, err := parseRecipeSource("desktop-identity-plan", args)
	if err != nil {
		return err
	}
	plan, err := appidentity.NewPlanWithProvenance(recipe, provenance)
	if err != nil {
		return err
	}
	if err := plan.ValidateSafeForDesktop(); err != nil {
		return err
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(plan)
}

func runDesktopEntryPreview(args []string, stdout io.Writer) error {
	recipe, provenance, options, err := parseDesktopEntryPreviewSource(args)
	if err != nil {
		return err
	}
	plan, err := appidentity.NewPlanWithProvenance(recipe, provenance)
	if err != nil {
		return err
	}
	entry, err := plan.RenderDesktopEntryWithOptions(options)
	if err != nil {
		return err
	}
	_, err = io.WriteString(stdout, entry)
	return err
}

func parseDesktopEntryPreviewSource(args []string) (appidentity.Recipe, appidentity.Provenance, appidentity.DesktopEntryOptions, error) {
	flags := flag.NewFlagSet("desktop-entry-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	applicationID := flags.String("app", "", "application id to load from the recipe registry")
	registryPath := flags.String("registry", "", "path to an Xnix recipe registry JSON file")
	recipeRoot := flags.String("recipe-root", "", "directory containing registered recipe files; defaults to the registry directory")
	recipePath := flags.String("recipe", "", "path to an Xnix application recipe JSON file")
	externalAppImportRecord := flags.String("external-app-import-record", "", "Runtime external Windows app import record")
	activationRoot := flags.String("activation-root", "", "read staged desktop activation receipt evidence from this explicit root")
	if err := flags.Parse(args); err != nil {
		return appidentity.Recipe{}, appidentity.Provenance{}, appidentity.DesktopEntryOptions{}, err
	}
	sourceCount := 0
	for _, source := range []string{*recipePath, *registryPath, *externalAppImportRecord} {
		if source != "" {
			sourceCount++
		}
	}
	if sourceCount != 1 {
		return appidentity.Recipe{}, appidentity.Provenance{}, appidentity.DesktopEntryOptions{}, errors.New("desktop-entry-preview requires exactly one source: --recipe, --registry, or --external-app-import-record")
	}
	if *registryPath != "" && *applicationID == "" {
		return appidentity.Recipe{}, appidentity.Provenance{}, appidentity.DesktopEntryOptions{}, errors.New("desktop-entry-preview requires --app when --registry is used")
	}
	if *recipePath != "" && (*applicationID != "" || *recipeRoot != "") {
		return appidentity.Recipe{}, appidentity.Provenance{}, appidentity.DesktopEntryOptions{}, errors.New("desktop-entry-preview --recipe cannot be combined with --app or --recipe-root")
	}
	if *externalAppImportRecord != "" && (*applicationID != "" || *recipeRoot != "") {
		return appidentity.Recipe{}, appidentity.Provenance{}, appidentity.DesktopEntryOptions{}, errors.New("desktop-entry-preview --external-app-import-record cannot be combined with --app or --recipe-root")
	}
	if flags.NArg() != 0 {
		return appidentity.Recipe{}, appidentity.Provenance{}, appidentity.DesktopEntryOptions{}, errors.New("desktop-entry-preview does not accept positional arguments")
	}

	if *externalAppImportRecord != "" {
		record, err := appidentity.LoadExternalWinAppImportRecord(*externalAppImportRecord)
		if err != nil {
			return appidentity.Recipe{}, appidentity.Provenance{}, appidentity.DesktopEntryOptions{}, err
		}
		recipe, provenance, err := appidentity.ExternalAppRecipeFromImportRecord(record)
		return recipe, provenance, appidentity.DesktopEntryOptions{ActivationRoot: *activationRoot}, err
	}
	recipe, provenance, err := loadRecipe(*recipePath, *registryPath, *recipeRoot, *applicationID)
	return recipe, provenance, appidentity.DesktopEntryOptions{ActivationRoot: *activationRoot}, err
}

func runDesktopIconPreview(args []string, stdout io.Writer) error {
	recipe, provenance, options, err := parseDesktopIconPreviewSource(args)
	if err != nil {
		return err
	}
	plan, err := appidentity.NewPlanWithProvenance(recipe, provenance)
	if err != nil {
		return err
	}
	preview, err := plan.DesktopIconPreviewWithOptions(options)
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(preview)
}

func parseDesktopIconPreviewSource(args []string) (appidentity.Recipe, appidentity.Provenance, appidentity.DesktopIconOptions, error) {
	flags := flag.NewFlagSet("desktop-icon-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	applicationID := flags.String("app", "", "application id to load from the recipe registry")
	registryPath := flags.String("registry", "", "path to an Xnix recipe registry JSON file")
	recipeRoot := flags.String("recipe-root", "", "directory containing registered recipe files; defaults to the registry directory")
	recipePath := flags.String("recipe", "", "path to an Xnix application recipe JSON file")
	activationRoot := flags.String("activation-root", "", "read staged desktop activation receipt evidence from this explicit root")
	if err := flags.Parse(args); err != nil {
		return appidentity.Recipe{}, appidentity.Provenance{}, appidentity.DesktopIconOptions{}, err
	}
	if (*recipePath == "") == (*registryPath == "") {
		return appidentity.Recipe{}, appidentity.Provenance{}, appidentity.DesktopIconOptions{}, errors.New("desktop-icon-preview requires exactly one source: --recipe or --registry")
	}
	if *registryPath != "" && *applicationID == "" {
		return appidentity.Recipe{}, appidentity.Provenance{}, appidentity.DesktopIconOptions{}, errors.New("desktop-icon-preview requires --app when --registry is used")
	}
	if *recipePath != "" && (*applicationID != "" || *recipeRoot != "") {
		return appidentity.Recipe{}, appidentity.Provenance{}, appidentity.DesktopIconOptions{}, errors.New("desktop-icon-preview --recipe cannot be combined with --app or --recipe-root")
	}
	if flags.NArg() != 0 {
		return appidentity.Recipe{}, appidentity.Provenance{}, appidentity.DesktopIconOptions{}, errors.New("desktop-icon-preview does not accept positional arguments")
	}

	recipe, provenance, err := loadRecipe(*recipePath, *registryPath, *recipeRoot, *applicationID)
	return recipe, provenance, appidentity.DesktopIconOptions{ActivationRoot: *activationRoot}, err
}

func runMIMEAppsPreview(args []string, stdout io.Writer) error {
	recipe, provenance, options, err := parseMIMEAppsPreviewSource(args)
	if err != nil {
		return err
	}
	plan, err := appidentity.NewPlanWithProvenance(recipe, provenance)
	if err != nil {
		return err
	}
	mimeapps, err := plan.RenderMIMEAppsWithOptions(options)
	if err != nil {
		return err
	}
	_, err = io.WriteString(stdout, mimeapps)
	return err
}

func parseMIMEAppsPreviewSource(args []string) (appidentity.Recipe, appidentity.Provenance, appidentity.MIMEAppsOptions, error) {
	flags := flag.NewFlagSet("mimeapps-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	applicationID := flags.String("app", "", "application id to load from the recipe registry")
	registryPath := flags.String("registry", "", "path to an Xnix recipe registry JSON file")
	recipeRoot := flags.String("recipe-root", "", "directory containing registered recipe files; defaults to the registry directory")
	recipePath := flags.String("recipe", "", "path to an Xnix application recipe JSON file")
	activationRoot := flags.String("activation-root", "", "read staged desktop activation receipt evidence from this explicit root")
	if err := flags.Parse(args); err != nil {
		return appidentity.Recipe{}, appidentity.Provenance{}, appidentity.MIMEAppsOptions{}, err
	}
	if (*recipePath == "") == (*registryPath == "") {
		return appidentity.Recipe{}, appidentity.Provenance{}, appidentity.MIMEAppsOptions{}, errors.New("mimeapps-preview requires exactly one source: --recipe or --registry")
	}
	if *registryPath != "" && *applicationID == "" {
		return appidentity.Recipe{}, appidentity.Provenance{}, appidentity.MIMEAppsOptions{}, errors.New("mimeapps-preview requires --app when --registry is used")
	}
	if *recipePath != "" && (*applicationID != "" || *recipeRoot != "") {
		return appidentity.Recipe{}, appidentity.Provenance{}, appidentity.MIMEAppsOptions{}, errors.New("mimeapps-preview --recipe cannot be combined with --app or --recipe-root")
	}
	if flags.NArg() != 0 {
		return appidentity.Recipe{}, appidentity.Provenance{}, appidentity.MIMEAppsOptions{}, errors.New("mimeapps-preview does not accept positional arguments")
	}

	recipe, provenance, err := loadRecipe(*recipePath, *registryPath, *recipeRoot, *applicationID)
	return recipe, provenance, appidentity.MIMEAppsOptions{ActivationRoot: *activationRoot}, err
}

func runDesktopResourceBridgePreview(args []string, stdout io.Writer) error {
	recipe, provenance, err := parseRecipeSource("desktop-resource-bridge-preview", args)
	if err != nil {
		return err
	}
	plan, err := appidentity.NewPlanWithProvenance(recipe, provenance)
	if err != nil {
		return err
	}
	preview, err := plan.DesktopResourceBridgePreview()
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(preview)
}

func runNotificationPreview(args []string, stdout io.Writer) error {
	recipe, provenance, eventType, options, err := parseNotificationPreviewSource(args)
	if err != nil {
		return err
	}
	plan, err := appidentity.NewPlanWithProvenance(recipe, provenance)
	if err != nil {
		return err
	}
	preview, err := plan.NotificationPreviewWithOptions(eventType, options)
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(preview)
}

func runSettingsPreview(args []string, stdout io.Writer) error {
	recipe, provenance, options, err := parseSettingsPreviewSource(args)
	if err != nil {
		return err
	}
	plan, err := appidentity.NewPlanWithProvenance(recipe, provenance)
	if err != nil {
		return err
	}
	preview, err := plan.SettingsPreviewWithOptions(options)
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(preview)
}

func runModeSwitchPreview(args []string, stdout io.Writer) error {
	recipe, provenance, requestedMode, err := parseModeSwitchPreviewSource(args)
	if err != nil {
		return err
	}
	plan, err := appidentity.NewPlanWithProvenance(recipe, provenance)
	if err != nil {
		return err
	}
	preview, err := plan.ModeSwitchPreview(requestedMode)
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(preview)
}

func runSettingsChangePreview(args []string, stdout io.Writer) error {
	recipe, provenance, sectionID, fieldID, value, err := parseSettingsChangePreviewSource(args)
	if err != nil {
		return err
	}
	plan, err := appidentity.NewPlanWithProvenance(recipe, provenance)
	if err != nil {
		return err
	}
	preview, err := plan.SettingsChangePreview(sectionID, fieldID, value)
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(preview)
}

func runReviewFlowPreview(args []string, stdout io.Writer) error {
	recipe, provenance, sectionID, fieldID, value, operation, err := parseReviewFlowPreviewSource(args)
	if err != nil {
		return err
	}
	plan, err := appidentity.NewPlanWithProvenance(recipe, provenance)
	if err != nil {
		return err
	}
	preview, err := plan.ReviewFlowPreview(sectionID, fieldID, value, operation)
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(preview)
}

func runPermissionReviewPreview(args []string, stdout io.Writer) error {
	recipe, provenance, err := parseRecipeSource("permission-review-preview", args)
	if err != nil {
		return err
	}
	plan, err := appidentity.NewPlanWithProvenance(recipe, provenance)
	if err != nil {
		return err
	}
	preview, err := plan.PermissionReviewPreview()
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(preview)
}

func runPortalRequestPreview(args []string, stdout io.Writer) error {
	recipe, provenance, operation, reason, err := parsePortalRequestPreviewSource(args)
	if err != nil {
		return err
	}
	plan, err := appidentity.NewPlanWithProvenance(recipe, provenance)
	if err != nil {
		return err
	}
	preview, err := plan.PortalRequestPreview(operation, reason)
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(preview)
}

func runTrayStatusPreview(args []string, stdout io.Writer) error {
	recipe, provenance, options, err := parseTrayStatusPreviewSource(args)
	if err != nil {
		return err
	}
	plan, err := appidentity.NewPlanWithProvenance(recipe, provenance)
	if err != nil {
		return err
	}
	preview, err := plan.TrayStatusPreviewWithOptions(options)
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(preview)
}

func runWindowIdentityPreview(args []string, stdout io.Writer) error {
	recipe, provenance, err := parseRecipeSource("window-identity-preview", args)
	if err != nil {
		return err
	}
	plan, err := appidentity.NewPlanWithProvenance(recipe, provenance)
	if err != nil {
		return err
	}
	preview, err := plan.WindowIdentityPreview()
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(preview)
}

func parseSettingsPreviewSource(args []string) (appidentity.Recipe, appidentity.Provenance, appidentity.SettingsOptions, error) {
	flags := flag.NewFlagSet("settings-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	applicationID := flags.String("app", "", "application id to load from the recipe registry")
	registryPath := flags.String("registry", "", "path to an Xnix recipe registry JSON file")
	recipeRoot := flags.String("recipe-root", "", "directory containing registered recipe files; defaults to the registry directory")
	recipePath := flags.String("recipe", "", "path to an Xnix application recipe JSON file")
	activationRoot := flags.String("activation-root", "", "read staged desktop activation receipt evidence from this explicit root")
	if err := flags.Parse(args); err != nil {
		return appidentity.Recipe{}, appidentity.Provenance{}, appidentity.SettingsOptions{}, err
	}
	if (*recipePath == "") == (*registryPath == "") {
		return appidentity.Recipe{}, appidentity.Provenance{}, appidentity.SettingsOptions{}, errors.New("settings-preview requires exactly one source: --recipe or --registry")
	}
	if *registryPath != "" && *applicationID == "" {
		return appidentity.Recipe{}, appidentity.Provenance{}, appidentity.SettingsOptions{}, errors.New("settings-preview requires --app when --registry is used")
	}
	if *recipePath != "" && (*applicationID != "" || *recipeRoot != "") {
		return appidentity.Recipe{}, appidentity.Provenance{}, appidentity.SettingsOptions{}, errors.New("settings-preview --recipe cannot be combined with --app or --recipe-root")
	}
	if flags.NArg() != 0 {
		return appidentity.Recipe{}, appidentity.Provenance{}, appidentity.SettingsOptions{}, errors.New("settings-preview does not accept positional arguments")
	}

	recipe, provenance, err := loadRecipe(*recipePath, *registryPath, *recipeRoot, *applicationID)
	return recipe, provenance, appidentity.SettingsOptions{ActivationRoot: *activationRoot}, err
}

func parseNotificationPreviewSource(args []string) (appidentity.Recipe, appidentity.Provenance, string, appidentity.NotificationOptions, error) {
	flags := flag.NewFlagSet("notification-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	applicationID := flags.String("app", "", "application id to load from the recipe registry")
	registryPath := flags.String("registry", "", "path to an Xnix recipe registry JSON file")
	recipeRoot := flags.String("recipe-root", "", "directory containing registered recipe files; defaults to the registry directory")
	recipePath := flags.String("recipe", "", "path to an Xnix application recipe JSON file")
	eventType := flags.String("event", "", "notification event type")
	activationRoot := flags.String("activation-root", "", "read staged desktop activation receipt evidence from this explicit root")
	if err := flags.Parse(args); err != nil {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", appidentity.NotificationOptions{}, err
	}
	if *eventType == "" {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", appidentity.NotificationOptions{}, errors.New("notification-preview requires --event")
	}
	if (*recipePath == "") == (*registryPath == "") {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", appidentity.NotificationOptions{}, errors.New("notification-preview requires exactly one source: --recipe or --registry")
	}
	if *registryPath != "" && *applicationID == "" {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", appidentity.NotificationOptions{}, errors.New("notification-preview requires --app when --registry is used")
	}
	if *recipePath != "" && (*applicationID != "" || *recipeRoot != "") {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", appidentity.NotificationOptions{}, errors.New("notification-preview --recipe cannot be combined with --app or --recipe-root")
	}
	if flags.NArg() != 0 {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", appidentity.NotificationOptions{}, errors.New("notification-preview does not accept positional arguments")
	}

	recipe, provenance, err := loadRecipe(*recipePath, *registryPath, *recipeRoot, *applicationID)
	return recipe, provenance, *eventType, appidentity.NotificationOptions{ActivationRoot: *activationRoot}, err
}

func parsePortalRequestPreviewSource(args []string) (appidentity.Recipe, appidentity.Provenance, string, string, error) {
	flags := flag.NewFlagSet("portal-request-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	applicationID := flags.String("app", "", "application id to load from the recipe registry")
	registryPath := flags.String("registry", "", "path to an Xnix recipe registry JSON file")
	recipeRoot := flags.String("recipe-root", "", "directory containing registered recipe files; defaults to the registry directory")
	recipePath := flags.String("recipe", "", "path to an Xnix application recipe JSON file")
	operation := flags.String("operation", "", "sensitive desktop operation")
	reason := flags.String("reason", "", "user-facing Portal request reason")
	if err := flags.Parse(args); err != nil {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", "", err
	}
	if *operation == "" {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", "", errors.New("portal-request-preview requires --operation")
	}
	if (*recipePath == "") == (*registryPath == "") {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", "", errors.New("portal-request-preview requires exactly one source: --recipe or --registry")
	}
	if *registryPath != "" && *applicationID == "" {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", "", errors.New("portal-request-preview requires --app when --registry is used")
	}
	if *recipePath != "" && (*applicationID != "" || *recipeRoot != "") {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", "", errors.New("portal-request-preview --recipe cannot be combined with --app or --recipe-root")
	}
	if flags.NArg() != 0 {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", "", errors.New("portal-request-preview does not accept positional arguments")
	}

	recipe, provenance, err := loadRecipe(*recipePath, *registryPath, *recipeRoot, *applicationID)
	return recipe, provenance, *operation, *reason, err
}

func parseModeSwitchPreviewSource(args []string) (appidentity.Recipe, appidentity.Provenance, string, error) {
	flags := flag.NewFlagSet("mode-switch-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	applicationID := flags.String("app", "", "application id to load from the recipe registry")
	registryPath := flags.String("registry", "", "path to an Xnix recipe registry JSON file")
	recipeRoot := flags.String("recipe-root", "", "directory containing registered recipe files; defaults to the registry directory")
	recipePath := flags.String("recipe", "", "path to an Xnix application recipe JSON file")
	requestedMode := flags.String("mode", "", "requested compatibility mode")
	if err := flags.Parse(args); err != nil {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", err
	}
	if *requestedMode == "" {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", errors.New("mode-switch-preview requires --mode")
	}
	if (*recipePath == "") == (*registryPath == "") {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", errors.New("mode-switch-preview requires exactly one source: --recipe or --registry")
	}
	if *registryPath != "" && *applicationID == "" {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", errors.New("mode-switch-preview requires --app when --registry is used")
	}
	if *recipePath != "" && (*applicationID != "" || *recipeRoot != "") {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", errors.New("mode-switch-preview --recipe cannot be combined with --app or --recipe-root")
	}
	if flags.NArg() != 0 {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", errors.New("mode-switch-preview does not accept positional arguments")
	}

	recipe, provenance, err := loadRecipe(*recipePath, *registryPath, *recipeRoot, *applicationID)
	return recipe, provenance, *requestedMode, err
}

func parseReviewFlowPreviewSource(args []string) (appidentity.Recipe, appidentity.Provenance, string, string, string, string, error) {
	flags := flag.NewFlagSet("review-flow-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	applicationID := flags.String("app", "", "application id to load from the recipe registry")
	registryPath := flags.String("registry", "", "path to an Xnix recipe registry JSON file")
	recipeRoot := flags.String("recipe-root", "", "directory containing registered recipe files; defaults to the registry directory")
	recipePath := flags.String("recipe", "", "path to an Xnix application recipe JSON file")
	sectionID := flags.String("section", "resource-access", "settings section identifier")
	fieldID := flags.String("field", "documents", "settings field identifier")
	value := flags.String("value", "ask", "requested settings value")
	operation := flags.String("operation", "file-open", "sensitive desktop operation")
	if err := flags.Parse(args); err != nil {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", "", "", "", err
	}
	if (*recipePath == "") == (*registryPath == "") {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", "", "", "", errors.New("review-flow-preview requires exactly one source: --recipe or --registry")
	}
	if *registryPath != "" && *applicationID == "" {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", "", "", "", errors.New("review-flow-preview requires --app when --registry is used")
	}
	if *recipePath != "" && (*applicationID != "" || *recipeRoot != "") {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", "", "", "", errors.New("review-flow-preview --recipe cannot be combined with --app or --recipe-root")
	}
	if flags.NArg() != 0 {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", "", "", "", errors.New("review-flow-preview does not accept positional arguments")
	}

	recipe, provenance, err := loadRecipe(*recipePath, *registryPath, *recipeRoot, *applicationID)
	return recipe, provenance, *sectionID, *fieldID, *value, *operation, err
}

func parseTrayStatusPreviewSource(args []string) (appidentity.Recipe, appidentity.Provenance, appidentity.TrayStatusOptions, error) {
	flags := flag.NewFlagSet("tray-status-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	applicationID := flags.String("app", "", "application id to load from the recipe registry")
	registryPath := flags.String("registry", "", "path to an Xnix recipe registry JSON file")
	recipeRoot := flags.String("recipe-root", "", "directory containing registered recipe files; defaults to the registry directory")
	recipePath := flags.String("recipe", "", "path to an Xnix application recipe JSON file")
	activationRoot := flags.String("activation-root", "", "read staged desktop activation receipt evidence from this explicit root")
	sessionRoot := flags.String("session-root", "", "read execution session status record evidence from this explicit root")
	sessionRequestID := flags.String("session-request-id", "", "execution session request id to read from --session-root")
	if err := flags.Parse(args); err != nil {
		return appidentity.Recipe{}, appidentity.Provenance{}, appidentity.TrayStatusOptions{}, err
	}
	if (*recipePath == "") == (*registryPath == "") {
		return appidentity.Recipe{}, appidentity.Provenance{}, appidentity.TrayStatusOptions{}, errors.New("tray-status-preview requires exactly one source: --recipe or --registry")
	}
	if *registryPath != "" && *applicationID == "" {
		return appidentity.Recipe{}, appidentity.Provenance{}, appidentity.TrayStatusOptions{}, errors.New("tray-status-preview requires --app when --registry is used")
	}
	if *recipePath != "" && (*applicationID != "" || *recipeRoot != "") {
		return appidentity.Recipe{}, appidentity.Provenance{}, appidentity.TrayStatusOptions{}, errors.New("tray-status-preview --recipe cannot be combined with --app or --recipe-root")
	}
	if flags.NArg() != 0 {
		return appidentity.Recipe{}, appidentity.Provenance{}, appidentity.TrayStatusOptions{}, errors.New("tray-status-preview does not accept positional arguments")
	}

	recipe, provenance, err := loadRecipe(*recipePath, *registryPath, *recipeRoot, *applicationID)
	return recipe, provenance, appidentity.TrayStatusOptions{ActivationRoot: *activationRoot, ExecutionSessionRoot: *sessionRoot, ExecutionSessionRequestID: *sessionRequestID}, err
}

func parseSettingsChangePreviewSource(args []string) (appidentity.Recipe, appidentity.Provenance, string, string, string, error) {
	flags := flag.NewFlagSet("settings-change-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	applicationID := flags.String("app", "", "application id to load from the recipe registry")
	registryPath := flags.String("registry", "", "path to an Xnix recipe registry JSON file")
	recipeRoot := flags.String("recipe-root", "", "directory containing registered recipe files; defaults to the registry directory")
	recipePath := flags.String("recipe", "", "path to an Xnix application recipe JSON file")
	sectionID := flags.String("section", "", "settings section identifier")
	fieldID := flags.String("field", "", "settings field identifier")
	value := flags.String("value", "", "requested settings value")
	if err := flags.Parse(args); err != nil {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", "", "", err
	}
	if *sectionID == "" {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", "", "", errors.New("settings-change-preview requires --section")
	}
	if *fieldID == "" {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", "", "", errors.New("settings-change-preview requires --field")
	}
	if *value == "" {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", "", "", errors.New("settings-change-preview requires --value")
	}
	if (*recipePath == "") == (*registryPath == "") {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", "", "", errors.New("settings-change-preview requires exactly one source: --recipe or --registry")
	}
	if *registryPath != "" && *applicationID == "" {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", "", "", errors.New("settings-change-preview requires --app when --registry is used")
	}
	if *recipePath != "" && (*applicationID != "" || *recipeRoot != "") {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", "", "", errors.New("settings-change-preview --recipe cannot be combined with --app or --recipe-root")
	}
	if flags.NArg() != 0 {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", "", "", errors.New("settings-change-preview does not accept positional arguments")
	}

	recipe, provenance, err := loadRecipe(*recipePath, *registryPath, *recipeRoot, *applicationID)
	return recipe, provenance, *sectionID, *fieldID, *value, err
}

func parseKRunnerQueryPreviewSource(args []string) ([]appidentity.Recipe, appidentity.Provenance, string, appidentity.KRunnerQueryOptions, error) {
	flags := flag.NewFlagSet("krunner-query-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	registryPath := flags.String("registry", "", "path to an Xnix recipe registry JSON file")
	recipeRoot := flags.String("recipe-root", "", "directory containing registered recipe files; defaults to the registry directory")
	query := flags.String("query", "", "KRunner query text")
	activationRoot := flags.String("activation-root", "", "read staged desktop activation receipt evidence from this explicit root")
	if err := flags.Parse(args); err != nil {
		return nil, appidentity.Provenance{}, "", appidentity.KRunnerQueryOptions{}, err
	}
	if *registryPath == "" {
		return nil, appidentity.Provenance{}, "", appidentity.KRunnerQueryOptions{}, errors.New("krunner-query-preview requires --registry")
	}
	if flags.NArg() != 0 {
		return nil, appidentity.Provenance{}, "", appidentity.KRunnerQueryOptions{}, errors.New("krunner-query-preview does not accept positional arguments")
	}

	recipes, provenance, err := appidentity.LoadRecipesFromRegistry(*registryPath, *recipeRoot)
	return recipes, provenance, *query, appidentity.KRunnerQueryOptions{ActivationRoot: *activationRoot}, err
}

func parseRegistryPreviewSource(commandName string, args []string) ([]appidentity.Recipe, appidentity.Provenance, error) {
	flags := flag.NewFlagSet(commandName, flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	registryPath := flags.String("registry", "", "path to an Xnix recipe registry JSON file")
	recipeRoot := flags.String("recipe-root", "", "directory containing registered recipe files; defaults to the registry directory")
	if err := flags.Parse(args); err != nil {
		return nil, appidentity.Provenance{}, err
	}
	if *registryPath == "" {
		return nil, appidentity.Provenance{}, fmt.Errorf("%s requires --registry", commandName)
	}
	if flags.NArg() != 0 {
		return nil, appidentity.Provenance{}, fmt.Errorf("%s does not accept positional arguments", commandName)
	}

	return appidentity.LoadRecipesFromRegistry(*registryPath, *recipeRoot)
}

func parseCompatibilityCenterPreviewSource(args []string) ([]appidentity.Recipe, appidentity.Provenance, appidentity.CompatibilityCenterOptions, error) {
	const commandName = "compatibility-center-preview"
	flags := flag.NewFlagSet(commandName, flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	registryPath := flags.String("registry", "", "path to an Xnix recipe registry JSON file")
	recipeRoot := flags.String("recipe-root", "", "directory containing registered recipe files; defaults to the registry directory")
	knownAppSmokeApp := flags.String("known-app-smoke-app", "", "known Windows app id with redacted smoke evidence")
	knownAppSmokeName := flags.String("known-app-smoke-name", "7-Zip Console", "known Windows app display name")
	knownAppSmokeVersion := flags.String("known-app-smoke-version", "26.02", "known Windows app version")
	knownAppSmokeSource := flags.String("known-app-smoke-source", "known-app-guest-smoke", "redacted known Windows app smoke evidence source")
	knownAppSmokeStatus := flags.String("known-app-smoke-status", "", "known Windows app smoke status")
	knownAppSmokeMarkerObserved := flags.Bool("known-app-smoke-marker-observed", false, "known Windows app marker observation result")
	knownAppSmokeChecksumVerified := flags.Bool("known-app-smoke-checksum-verified", false, "known Windows app checksum verification result")
	knownAppEvidenceJSON := flags.String("known-app-evidence-json", "", "Runtime-projected known Windows app evidence JSON")
	knownAppEvidenceFile := flags.String("known-app-evidence-file", "", "file containing Runtime-projected known Windows app evidence JSON")
	knownAppMatrixReport := flags.String("known-app-matrix-report", "", "aggregate known Windows app matrix evidence report JSON")
	knownAppVerifiedCatalog := flags.String("known-app-verified-catalog", "", "Go-owned verified known Windows app catalog JSON")
	knownAppGUISmokeReport := flags.String("known-app-gui-smoke-report", "", "executed Wine guest GUI smoke evidence report JSON")
	knownAppGUISmokeApp := flags.String("known-app-gui-smoke-app", "", "known Windows GUI app id for a GUI smoke report")
	knownAppGUISmokeName := flags.String("known-app-gui-smoke-name", "", "known Windows GUI app display name for a GUI smoke report")
	knownAppGUISmokeVersion := flags.String("known-app-gui-smoke-version", "", "known Windows GUI app version for a GUI smoke report")
	knownAppLaunchAuthorizationReceiptState := flags.String("known-app-launch-authorization-receipt-state", "", "known Windows app launch authorization receipt state")
	knownAppLaunchAuthorizationReceiptID := flags.String("known-app-launch-authorization-receipt-id", "", "opaque known Windows app launch authorization receipt id")
	knownAppLaunchGateState := flags.String("known-app-launch-gate-state", "", "known Windows app launch gate state")
	knownAppLaunchGateConsumed := flags.Bool("known-app-launch-gate-consumed", false, "known Windows app launch gate consumption result")
	knownAppLaunchGateReceiptAccepted := flags.Bool("known-app-launch-gate-receipt-accepted", false, "known Windows app launch gate receipt acceptance result")
	knownAppLaunchGateGuestBoundaryAccepted := flags.Bool("known-app-launch-gate-guest-boundary-accepted", false, "known Windows app launch gate guest boundary acceptance result")
	knownAppControlledDispatchReady := flags.Bool("known-app-controlled-dispatch-ready", false, "known Windows app controlled dispatch readiness result")
	knownAppControlledExecutionSessionID := flags.String("known-app-controlled-execution-session-id", "", "opaque known Windows app controlled execution session id")
	knownAppLauncherSessionGateConsumed := flags.Bool("known-app-launcher-session-gate-consumed", false, "known Windows app launcher-side session gate consumption result")
	knownAppLauncherSessionDigestVerified := flags.Bool("known-app-launcher-session-digest-verified", false, "known Windows app launcher-side session digest verification result")
	knownAppLauncherSessionRelativePath := flags.String("known-app-launcher-session-relative-path", "", "relative known Windows app launcher-side session evidence path")
	knownAppLauncherSessionRuntimeOwnerConsumable := flags.Bool("known-app-launcher-session-runtime-owner-consumable", false, "known Windows app launcher-side Runtime-owner session consumption readiness")
	knownAppLauncherSessionKDEReadModelConsumable := flags.Bool("known-app-launcher-session-kde-read-model-consumable", false, "known Windows app launcher-side KDE read-model session consumption readiness")
	knownAppPostReviewDispatchConsumed := flags.Bool("known-app-post-review-dispatch-consumed", false, "known Windows app launcher-side post-review dispatch consumption result")
	knownAppPostReviewDispatchState := flags.String("known-app-post-review-dispatch-state", "", "known Windows app launcher-side post-review dispatch state")
	knownAppSessionGatedReviewReceiptID := flags.String("known-app-session-gated-review-receipt-id", "", "opaque known Windows app session-gated review receipt id")
	knownAppLaunchGateBlockedReason := flags.String("known-app-launch-gate-blocked-reason", "", "redacted known Windows app launch gate blocked reason")
	if err := flags.Parse(args); err != nil {
		return nil, appidentity.Provenance{}, appidentity.CompatibilityCenterOptions{}, err
	}
	if *registryPath == "" {
		return nil, appidentity.Provenance{}, appidentity.CompatibilityCenterOptions{}, fmt.Errorf("%s requires --registry", commandName)
	}
	if flags.NArg() != 0 {
		return nil, appidentity.Provenance{}, appidentity.CompatibilityCenterOptions{}, fmt.Errorf("%s does not accept positional arguments", commandName)
	}

	recipes, provenance, err := appidentity.LoadRecipesFromRegistry(*registryPath, *recipeRoot)
	if err != nil {
		return nil, appidentity.Provenance{}, appidentity.CompatibilityCenterOptions{}, err
	}

	options := appidentity.CompatibilityCenterOptions{}
	verifiedCatalog, err := loadKnownAppVerifiedCatalogForCenter(commandName, *knownAppVerifiedCatalog)
	if err != nil {
		return nil, appidentity.Provenance{}, appidentity.CompatibilityCenterOptions{}, err
	}
	options.KnownAppVerifiedCatalog = verifiedCatalog
	projectedEvidence, err := loadKnownAppSmokeEvidenceForCenter(commandName, *knownAppEvidenceJSON, *knownAppEvidenceFile, *knownAppMatrixReport, *knownAppGUISmokeReport, *knownAppGUISmokeApp, *knownAppGUISmokeName, *knownAppGUISmokeVersion)
	if err != nil {
		return nil, appidentity.Provenance{}, appidentity.CompatibilityCenterOptions{}, err
	}
	legacyKnownAppSmoke := *knownAppSmokeApp != "" || *knownAppSmokeStatus != ""
	if len(projectedEvidence) > 0 {
		if legacyKnownAppSmoke {
			return nil, appidentity.Provenance{}, appidentity.CompatibilityCenterOptions{}, fmt.Errorf("%s accepts either Runtime-projected known app evidence or legacy known app smoke flags, not both", commandName)
		}
		options.KnownAppSmokeEvidence = projectedEvidence
	} else if legacyKnownAppSmoke {
		options.KnownAppSmokeEvidence = []appidentity.KnownAppSmokeEvidenceSummary{{
			AppID:                                 *knownAppSmokeApp,
			DisplayName:                           *knownAppSmokeName,
			AppVersion:                            *knownAppSmokeVersion,
			EvidenceSource:                        *knownAppSmokeSource,
			SmokeStatus:                           *knownAppSmokeStatus,
			LaunchAuthorizationReceiptState:       *knownAppLaunchAuthorizationReceiptState,
			LaunchAuthorizationReceiptID:          *knownAppLaunchAuthorizationReceiptID,
			LaunchGateState:                       *knownAppLaunchGateState,
			LaunchGateConsumed:                    *knownAppLaunchGateConsumed,
			LaunchGateReceiptAccepted:             *knownAppLaunchGateReceiptAccepted,
			LaunchGateGuestBoundaryAccepted:       *knownAppLaunchGateGuestBoundaryAccepted,
			ControlledDispatchReady:               *knownAppControlledDispatchReady,
			ControlledExecutionSessionID:          *knownAppControlledExecutionSessionID,
			LauncherSessionGateConsumed:           *knownAppLauncherSessionGateConsumed,
			LauncherSessionDigestVerified:         *knownAppLauncherSessionDigestVerified,
			LauncherSessionRelativePath:           *knownAppLauncherSessionRelativePath,
			LauncherSessionRuntimeOwnerConsumable: *knownAppLauncherSessionRuntimeOwnerConsumable,
			LauncherSessionKDEReadModelConsumable: *knownAppLauncherSessionKDEReadModelConsumable,
			PostReviewDispatchConsumed:            *knownAppPostReviewDispatchConsumed,
			PostReviewDispatchState:               *knownAppPostReviewDispatchState,
			SessionGatedReviewReceiptID:           *knownAppSessionGatedReviewReceiptID,
			LaunchGateBlockedReason:               *knownAppLaunchGateBlockedReason,
			MarkerObserved:                        *knownAppSmokeMarkerObserved,
			ChecksumVerified:                      *knownAppSmokeChecksumVerified,
		}}
	}
	return recipes, provenance, options, nil
}

func loadKnownAppSmokeEvidenceForCenter(commandName string, projectionJSON string, projectionFile string, matrixReport string, guiSmokeReport string, guiSmokeApp string, guiSmokeName string, guiSmokeVersion string) ([]appidentity.KnownAppSmokeEvidenceSummary, error) {
	projectionJSON = strings.TrimSpace(projectionJSON)
	projectionFile = strings.TrimSpace(projectionFile)
	matrixReport = strings.TrimSpace(matrixReport)
	guiSmokeReport = strings.TrimSpace(guiSmokeReport)
	sourceCount := 0
	if projectionJSON != "" {
		sourceCount++
	}
	if projectionFile != "" {
		sourceCount++
	}
	if matrixReport != "" {
		sourceCount++
	}
	if guiSmokeReport != "" {
		sourceCount++
	}
	if sourceCount == 0 {
		return nil, nil
	}
	if sourceCount > 1 {
		return nil, fmt.Errorf("%s accepts only one known app evidence source", commandName)
	}
	if matrixReport != "" {
		preview, err := appidentity.PreviewKnownAppMatrixEvidence(appidentity.KnownAppMatrixEvidencePreviewRequest{MatrixReportPath: matrixReport})
		if err != nil {
			return nil, fmt.Errorf("consume known app matrix evidence: %w", err)
		}
		return preview.KnownAppSmokeEvidence, nil
	}
	if guiSmokeReport != "" {
		preview, err := appidentity.PreviewGUISmokeEvidence(appidentity.GUISmokeEvidencePreviewRequest{
			ReportPath:  guiSmokeReport,
			AppID:       guiSmokeApp,
			DisplayName: guiSmokeName,
			AppVersion:  guiSmokeVersion,
		})
		if err != nil {
			return nil, fmt.Errorf("consume GUI smoke evidence: %w", err)
		}
		return []appidentity.KnownAppSmokeEvidenceSummary{preview.KnownAppSmokeEvidence}, nil
	}
	payload := []byte(projectionJSON)
	if projectionFile != "" {
		loaded, err := os.ReadFile(projectionFile)
		if err != nil {
			return nil, fmt.Errorf("read Runtime-projected known app evidence: %w", err)
		}
		payload = loaded
	}
	var envelope struct {
		SchemaVersion string `json:"schema_version"`
		RequestType   string `json:"request_type"`
	}
	if err := json.Unmarshal(payload, &envelope); err != nil {
		return nil, fmt.Errorf("parse Runtime-projected known app evidence: %w", err)
	}
	if envelope.RequestType == appidentity.GUISmokeEvidencePreviewRequestType {
		return loadKnownAppSmokeEvidenceFromGUISmokeProjection(payload)
	}
	if envelope.RequestType == appidentity.RealWinAppGUIEvidencePacketRequestType {
		evidence, err := appidentity.KnownAppSmokeEvidenceFromRealWinAppGUIEvidencePacket(payload)
		if err != nil {
			return nil, err
		}
		return []appidentity.KnownAppSmokeEvidenceSummary{evidence}, nil
	}
	if envelope.RequestType == appidentity.RealWinAppRunReceiptSummaryRequestType {
		evidence, err := appidentity.KnownAppSmokeEvidenceFromRealWinAppRunReceiptSummary(payload)
		if err != nil {
			return nil, err
		}
		return []appidentity.KnownAppSmokeEvidenceSummary{evidence}, nil
	}
	var projection appidentity.KnownAppKDERuntimeStatusLaunchDelegatedEvidence
	if err := json.Unmarshal(payload, &projection); err != nil {
		return nil, fmt.Errorf("parse Runtime-projected known app evidence: %w", err)
	}
	evidence, err := appidentity.KnownAppSmokeEvidenceFromKDERuntimeStatusLaunchDelegatedEvidence(projection)
	if err != nil {
		return nil, fmt.Errorf("consume Runtime-projected known app evidence: %w", err)
	}
	return []appidentity.KnownAppSmokeEvidenceSummary{evidence}, nil
}

func loadKnownAppSmokeEvidenceFromGUISmokeProjection(payload []byte) ([]appidentity.KnownAppSmokeEvidenceSummary, error) {
	evidence, err := appidentity.KnownAppSmokeEvidenceFromGUISmokeProjection(payload)
	if err != nil {
		return nil, err
	}
	return []appidentity.KnownAppSmokeEvidenceSummary{evidence}, nil
}

func loadKnownAppSmokeEvidenceFromRuntimeProjection(commandName string, projectionJSON string, projectionFile string) ([]appidentity.KnownAppSmokeEvidenceSummary, error) {
	return loadKnownAppSmokeEvidenceForCenter(commandName, projectionJSON, projectionFile, "", "", "", "", "")
}

func loadKnownAppVerifiedCatalogForCenter(commandName string, verifiedCatalogPath string) (*appidentity.KnownAppVerifiedCatalogPreview, error) {
	verifiedCatalogPath = strings.TrimSpace(verifiedCatalogPath)
	if verifiedCatalogPath == "" {
		return nil, nil
	}
	payload, err := os.ReadFile(verifiedCatalogPath)
	if err != nil {
		return nil, fmt.Errorf("%s read known app verified catalog: %w", commandName, err)
	}
	var preview appidentity.KnownAppVerifiedCatalogPreview
	if err := json.Unmarshal(payload, &preview); err != nil {
		return nil, fmt.Errorf("%s parse known app verified catalog: %w", commandName, err)
	}
	return &preview, nil
}

func parseFileOpenPreviewSource(args []string) ([]appidentity.Recipe, appidentity.Provenance, string, []string, appidentity.FileOpenOptions, error) {
	flags := flag.NewFlagSet("file-open-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	applicationID := flags.String("app", "", "application id to use for the file-open preview")
	registryPath := flags.String("registry", "", "path to an Xnix recipe registry JSON file")
	recipeRoot := flags.String("recipe-root", "", "directory containing registered recipe files; defaults to the registry directory")
	activationRoot := flags.String("activation-root", "", "read staged desktop activation receipt evidence from this explicit root")
	if err := flags.Parse(args); err != nil {
		return nil, appidentity.Provenance{}, "", nil, appidentity.FileOpenOptions{}, err
	}
	if *registryPath == "" {
		return nil, appidentity.Provenance{}, "", nil, appidentity.FileOpenOptions{}, errors.New("file-open-preview requires --registry")
	}
	fileURIs := flags.Args()
	if len(fileURIs) == 0 {
		return nil, appidentity.Provenance{}, "", nil, appidentity.FileOpenOptions{}, errors.New("file-open-preview requires at least one file URI")
	}

	recipes, provenance, err := appidentity.LoadRecipesFromRegistry(*registryPath, *recipeRoot)
	return recipes, provenance, *applicationID, fileURIs, appidentity.FileOpenOptions{ActivationRoot: *activationRoot}, err
}

func parseDolphinDropPreviewSource(args []string) ([]appidentity.Recipe, appidentity.Provenance, string, []string, appidentity.FileOpenOptions, error) {
	flags := flag.NewFlagSet("dolphin-drop-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	applicationID := flags.String("app", "", "application id to use for the Dolphin drop preview")
	registryPath := flags.String("registry", "", "path to an Xnix recipe registry JSON file")
	recipeRoot := flags.String("recipe-root", "", "directory containing registered recipe files; defaults to the registry directory")
	activationRoot := flags.String("activation-root", "", "read staged desktop activation receipt evidence from this explicit root")
	if err := flags.Parse(args); err != nil {
		return nil, appidentity.Provenance{}, "", nil, appidentity.FileOpenOptions{}, err
	}
	if *registryPath == "" {
		return nil, appidentity.Provenance{}, "", nil, appidentity.FileOpenOptions{}, errors.New("dolphin-drop-preview requires --registry")
	}
	fileURIs := flags.Args()
	if len(fileURIs) == 0 {
		return nil, appidentity.Provenance{}, "", nil, appidentity.FileOpenOptions{}, errors.New("dolphin-drop-preview requires at least one file URI")
	}

	recipes, provenance, err := appidentity.LoadRecipesFromRegistry(*registryPath, *recipeRoot)
	return recipes, provenance, *applicationID, fileURIs, appidentity.FileOpenOptions{ActivationRoot: *activationRoot}, err
}

func parseDolphinAIAnalysisPreviewSource(args []string) ([]appidentity.Recipe, appidentity.Provenance, string, []string, error) {
	flags := flag.NewFlagSet("dolphin-ai-analysis-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	applicationID := flags.String("app", "", "application id to use for the Dolphin AI analysis preview")
	registryPath := flags.String("registry", "", "path to an Xnix recipe registry JSON file")
	recipeRoot := flags.String("recipe-root", "", "directory containing registered recipe files; defaults to the registry directory")
	if err := flags.Parse(args); err != nil {
		return nil, appidentity.Provenance{}, "", nil, err
	}
	if *registryPath == "" {
		return nil, appidentity.Provenance{}, "", nil, errors.New("dolphin-ai-analysis-preview requires --registry")
	}
	fileURIs := flags.Args()
	if len(fileURIs) == 0 {
		return nil, appidentity.Provenance{}, "", nil, errors.New("dolphin-ai-analysis-preview requires at least one file URI")
	}

	recipes, provenance, err := appidentity.LoadRecipesFromRegistry(*registryPath, *recipeRoot)
	return recipes, provenance, *applicationID, fileURIs, err
}

func parseLaunchIntentPreviewSource(args []string) (appidentity.Recipe, appidentity.Provenance, []string, error) {
	return parseLaunchActionPreviewSource("launch-intent-preview", args)
}

func parseExecutionRequestPreviewSource(args []string) (appidentity.Recipe, appidentity.Provenance, []string, error) {
	return parseLaunchActionPreviewSource("execution-request-preview", args)
}

func parseExecutionReviewPreviewSource(args []string) (appidentity.Recipe, appidentity.Provenance, []string, error) {
	return parseLaunchActionPreviewSource("execution-review-preview", args)
}

func parseExecutionDecisionPreviewSource(args []string) (appidentity.Recipe, appidentity.Provenance, string, []string, error) {
	return parseLaunchDecisionPreviewSource("execution-decision-preview", args)
}

func parseExecutionPreflightPreviewSource(args []string) (appidentity.Recipe, appidentity.Provenance, string, []string, error) {
	return parseLaunchDecisionPreviewSource("execution-preflight-preview", args)
}

func parseExecutionResourceGrantPreviewSource(args []string) (appidentity.Recipe, appidentity.Provenance, string, []string, error) {
	return parseLaunchDecisionPreviewSource("execution-resource-grant-preview", args)
}

func parseExecutionTransactionPreviewSource(args []string) (appidentity.Recipe, appidentity.Provenance, string, []string, error) {
	return parseLaunchDecisionPreviewSource("execution-transaction-preview", args)
}

func parseExecutionSessionPreviewSource(args []string) (appidentity.Recipe, appidentity.Provenance, string, []string, error) {
	return parseLaunchDecisionPreviewSource("execution-session-preview", args)
}

func parseExecutionSessionStatusPreviewSource(args []string) (appidentity.Recipe, appidentity.Provenance, string, []string, error) {
	return parseLaunchDecisionPreviewSource("execution-session-status-preview", args)
}

func parseKDEEntryPointsPreviewSource(args []string) (appidentity.Recipe, appidentity.Provenance, string, []string, error) {
	return parseLaunchDecisionPreviewSource("kde-entrypoints-preview", args)
}

func parseKDEActionQueuePreviewSource(args []string) (appidentity.Recipe, appidentity.Provenance, string, []string, error) {
	return parseLaunchDecisionPreviewSource("kde-action-queue-preview", args)
}

func parseKDEActionCardDeckPreviewSource(args []string) (appidentity.Recipe, appidentity.Provenance, string, []string, error) {
	return parseLaunchDecisionPreviewSource("kde-action-card-deck-preview", args)
}

func parseKDEActionDependencyGraphPreviewSource(args []string) (appidentity.Recipe, appidentity.Provenance, string, []string, error) {
	return parseLaunchDecisionPreviewSource("kde-action-dependency-graph-preview", args)
}

func parseKDEJourneyEvidencePreviewSource(args []string) (appidentity.Recipe, appidentity.Provenance, string, []string, appidentity.KDEJourneyEvidenceOptions, error) {
	flags := flag.NewFlagSet("kde-journey-evidence-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	applicationID := flags.String("app", "", "application id to load from the recipe registry")
	decision := flags.String("decision", "", "review decision: reviewed, approved, deferred, or rejected")
	registryPath := flags.String("registry", "", "path to an Xnix recipe registry JSON file")
	recipeRoot := flags.String("recipe-root", "", "directory containing registered recipe files; defaults to the registry directory")
	recipePath := flags.String("recipe", "", "path to an Xnix application recipe JSON file")
	runtimeRoot := flags.String("runtime-root", ".", "Runtime repository root for journey evidence reads")
	if err := flags.Parse(args); err != nil {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", nil, appidentity.KDEJourneyEvidenceOptions{}, err
	}
	if (*recipePath == "") == (*registryPath == "") {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", nil, appidentity.KDEJourneyEvidenceOptions{}, errors.New("kde-journey-evidence-preview requires exactly one source: --recipe or --registry")
	}
	if *registryPath != "" && *applicationID == "" {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", nil, appidentity.KDEJourneyEvidenceOptions{}, errors.New("kde-journey-evidence-preview requires --app when --registry is used")
	}
	if *decision == "" {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", nil, appidentity.KDEJourneyEvidenceOptions{}, errors.New("kde-journey-evidence-preview requires --decision")
	}
	if *recipePath != "" && (*applicationID != "" || *recipeRoot != "") {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", nil, appidentity.KDEJourneyEvidenceOptions{}, errors.New("kde-journey-evidence-preview --recipe cannot be combined with --app or --recipe-root")
	}

	recipe, provenance, err := loadRecipe(*recipePath, *registryPath, *recipeRoot, *applicationID)
	return recipe, provenance, *decision, flags.Args(), appidentity.KDEJourneyEvidenceOptions{RuntimeRoot: *runtimeRoot}, err
}

func parseKDECenterPagePreviewSource(args []string) (appidentity.Recipe, appidentity.Provenance, string, []string, appidentity.KDECenterPageOptions, error) {
	flags := flag.NewFlagSet("kde-center-page-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	applicationID := flags.String("app", "", "application id to load from the recipe registry")
	decision := flags.String("decision", "", "review decision: reviewed, approved, deferred, or rejected")
	registryPath := flags.String("registry", "", "path to an Xnix recipe registry JSON file")
	recipeRoot := flags.String("recipe-root", "", "directory containing registered recipe files; defaults to the registry directory")
	recipePath := flags.String("recipe", "", "path to an Xnix application recipe JSON file")
	externalAppEvidenceFile := flags.String("external-app-evidence-file", "", "real GUI evidence packet for a non-recipe external Windows application")
	externalAppImportRecord := flags.String("external-app-import-record", "", "Runtime external Windows app import record")
	activationRoot := flags.String("activation-root", "", "read staged desktop activation receipt evidence from this explicit root")
	sessionRoot := flags.String("session-root", "", "read execution session status record evidence from this explicit root")
	sessionRequestID := flags.String("session-request-id", "", "execution session request id to read from --session-root")
	readinessRoot := flags.String("readiness-root", "", "Runtime repository root for application readiness evidence reads")
	readinessStateRoot := flags.String("readiness-state-root", "", "optional Runtime state root for application readiness evidence")
	readinessArtifactReceipt := flags.String("readiness-artifact-receipt", "", "optional artifact stage receipt JSON file for application readiness evidence")
	readinessPortalOperation := flags.String("readiness-portal-operation", "file-open", "Portal operation to evaluate for application readiness evidence")
	readinessSnapshotReason := flags.String("readiness-snapshot-reason", "before-repair", "snapshot reason to evaluate for application readiness evidence")
	knownAppSmokeApp := flags.String("known-app-smoke-app", "", "known Windows app id with redacted smoke evidence")
	knownAppSmokeName := flags.String("known-app-smoke-name", "7-Zip Console", "known Windows app display name")
	knownAppSmokeVersion := flags.String("known-app-smoke-version", "26.02", "known Windows app version")
	knownAppSmokeSource := flags.String("known-app-smoke-source", "known-app-guest-smoke", "redacted known Windows app smoke evidence source")
	knownAppSmokeStatus := flags.String("known-app-smoke-status", "", "known Windows app smoke status")
	knownAppSmokeMarkerObserved := flags.Bool("known-app-smoke-marker-observed", false, "known Windows app marker observation result")
	knownAppSmokeChecksumVerified := flags.Bool("known-app-smoke-checksum-verified", false, "known Windows app checksum verification result")
	knownAppEvidenceJSON := flags.String("known-app-evidence-json", "", "Runtime-projected known Windows app evidence JSON")
	knownAppEvidenceFile := flags.String("known-app-evidence-file", "", "file containing Runtime-projected known Windows app evidence JSON")
	knownAppMatrixReport := flags.String("known-app-matrix-report", "", "aggregate known Windows app matrix evidence report JSON")
	knownAppVerifiedCatalog := flags.String("known-app-verified-catalog", "", "Go-owned verified known Windows app catalog JSON")
	knownAppGUISmokeReport := flags.String("known-app-gui-smoke-report", "", "executed Wine guest GUI smoke evidence report JSON")
	knownAppGUISmokeApp := flags.String("known-app-gui-smoke-app", "", "known Windows GUI app id for a GUI smoke report")
	knownAppGUISmokeName := flags.String("known-app-gui-smoke-name", "", "known Windows GUI app display name for a GUI smoke report")
	knownAppGUISmokeVersion := flags.String("known-app-gui-smoke-version", "", "known Windows GUI app version for a GUI smoke report")
	knownAppLaunchAuthorizationReceiptState := flags.String("known-app-launch-authorization-receipt-state", "", "known Windows app launch authorization receipt state")
	knownAppLaunchAuthorizationReceiptID := flags.String("known-app-launch-authorization-receipt-id", "", "opaque known Windows app launch authorization receipt id")
	knownAppLaunchGateState := flags.String("known-app-launch-gate-state", "", "known Windows app launch gate state")
	knownAppLaunchGateConsumed := flags.Bool("known-app-launch-gate-consumed", false, "known Windows app launch gate consumption result")
	knownAppLaunchGateReceiptAccepted := flags.Bool("known-app-launch-gate-receipt-accepted", false, "known Windows app launch gate receipt acceptance result")
	knownAppLaunchGateGuestBoundaryAccepted := flags.Bool("known-app-launch-gate-guest-boundary-accepted", false, "known Windows app launch gate guest boundary acceptance result")
	knownAppControlledDispatchReady := flags.Bool("known-app-controlled-dispatch-ready", false, "known Windows app controlled dispatch readiness result")
	knownAppControlledExecutionSessionID := flags.String("known-app-controlled-execution-session-id", "", "opaque known Windows app controlled execution session id")
	knownAppLauncherSessionGateConsumed := flags.Bool("known-app-launcher-session-gate-consumed", false, "known Windows app launcher-side session gate consumption result")
	knownAppLauncherSessionDigestVerified := flags.Bool("known-app-launcher-session-digest-verified", false, "known Windows app launcher-side session digest verification result")
	knownAppLauncherSessionRelativePath := flags.String("known-app-launcher-session-relative-path", "", "relative known Windows app launcher-side session evidence path")
	knownAppLauncherSessionRuntimeOwnerConsumable := flags.Bool("known-app-launcher-session-runtime-owner-consumable", false, "known Windows app launcher-side Runtime-owner session consumption readiness")
	knownAppLauncherSessionKDEReadModelConsumable := flags.Bool("known-app-launcher-session-kde-read-model-consumable", false, "known Windows app launcher-side KDE read-model session consumption readiness")
	knownAppPostReviewDispatchConsumed := flags.Bool("known-app-post-review-dispatch-consumed", false, "known Windows app launcher-side post-review dispatch consumption result")
	knownAppPostReviewDispatchState := flags.String("known-app-post-review-dispatch-state", "", "known Windows app launcher-side post-review dispatch state")
	knownAppSessionGatedReviewReceiptID := flags.String("known-app-session-gated-review-receipt-id", "", "opaque known Windows app session-gated review receipt id")
	knownAppLaunchGateBlockedReason := flags.String("known-app-launch-gate-blocked-reason", "", "redacted known Windows app launch gate blocked reason")
	if err := flags.Parse(args); err != nil {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", nil, appidentity.KDECenterPageOptions{}, err
	}
	sourceCount := 0
	if *recipePath != "" {
		sourceCount++
	}
	if *registryPath != "" {
		sourceCount++
	}
	if *externalAppEvidenceFile != "" {
		sourceCount++
	}
	if *externalAppImportRecord != "" {
		sourceCount++
	}
	if sourceCount != 1 {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", nil, appidentity.KDECenterPageOptions{}, errors.New("kde-center-page-preview requires exactly one source: --recipe, --registry, --external-app-evidence-file, or --external-app-import-record")
	}
	if *registryPath != "" && *applicationID == "" {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", nil, appidentity.KDECenterPageOptions{}, errors.New("kde-center-page-preview requires --app when --registry is used")
	}
	if *recipePath != "" && (*applicationID != "" || *recipeRoot != "") {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", nil, appidentity.KDECenterPageOptions{}, errors.New("kde-center-page-preview --recipe cannot be combined with --app or --recipe-root")
	}
	if *externalAppEvidenceFile != "" && (*applicationID != "" || *recipeRoot != "") {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", nil, appidentity.KDECenterPageOptions{}, errors.New("kde-center-page-preview --external-app-evidence-file cannot be combined with --app or --recipe-root")
	}
	if *externalAppImportRecord != "" && (*applicationID != "" || *recipeRoot != "") {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", nil, appidentity.KDECenterPageOptions{}, errors.New("kde-center-page-preview --external-app-import-record cannot be combined with --app or --recipe-root")
	}
	fileURIs := flags.Args()

	var receipt *artifact.StageReceipt
	if *readinessArtifactReceipt != "" {
		loaded, err := loadArtifactStageReceipt(*readinessArtifactReceipt)
		if err != nil {
			return appidentity.Recipe{}, appidentity.Provenance{}, "", nil, appidentity.KDECenterPageOptions{}, err
		}
		receipt = &loaded
	}
	if *externalAppEvidenceFile != "" {
		payload, err := os.ReadFile(*externalAppEvidenceFile)
		if err != nil {
			return appidentity.Recipe{}, appidentity.Provenance{}, "", nil, appidentity.KDECenterPageOptions{}, fmt.Errorf("read external app evidence file: %w", err)
		}
		recipe, provenance, evidence, err := appidentity.ExternalAppRecipeFromRealWinAppGUIEvidencePacket(payload)
		if err != nil {
			return appidentity.Recipe{}, appidentity.Provenance{}, "", nil, appidentity.KDECenterPageOptions{}, err
		}
		return recipe, provenance, *decision, fileURIs, appidentity.KDECenterPageOptions{
			ActivationRoot:                      *activationRoot,
			ExecutionSessionRoot:                *sessionRoot,
			ExecutionSessionRequestID:           *sessionRequestID,
			ApplicationReadinessRoot:            *readinessRoot,
			ApplicationReadinessStateRoot:       *readinessStateRoot,
			ApplicationReadinessArtifactReceipt: receipt,
			ApplicationReadinessPortalOperation: *readinessPortalOperation,
			ApplicationReadinessSnapshotReason:  *readinessSnapshotReason,
			KnownAppSmokeEvidence:               []appidentity.KnownAppSmokeEvidenceSummary{evidence},
		}, nil
	}
	if *externalAppImportRecord != "" {
		record, err := appidentity.LoadExternalWinAppImportRecord(*externalAppImportRecord)
		if err != nil {
			return appidentity.Recipe{}, appidentity.Provenance{}, "", nil, appidentity.KDECenterPageOptions{}, err
		}
		recipe, provenance, err := appidentity.ExternalAppRecipeFromImportRecord(record)
		if err != nil {
			return appidentity.Recipe{}, appidentity.Provenance{}, "", nil, appidentity.KDECenterPageOptions{}, err
		}
		return recipe, provenance, *decision, fileURIs, appidentity.KDECenterPageOptions{
			ActivationRoot:                      *activationRoot,
			ExecutionSessionRoot:                *sessionRoot,
			ExecutionSessionRequestID:           *sessionRequestID,
			ApplicationReadinessRoot:            *readinessRoot,
			ApplicationReadinessStateRoot:       *readinessStateRoot,
			ApplicationReadinessArtifactReceipt: receipt,
			ApplicationReadinessPortalOperation: *readinessPortalOperation,
			ApplicationReadinessSnapshotReason:  *readinessSnapshotReason,
		}, nil
	}

	recipe, provenance, err := loadRecipe(*recipePath, *registryPath, *recipeRoot, *applicationID)
	if err != nil {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", nil, appidentity.KDECenterPageOptions{}, err
	}
	knownAppSmokeEvidence := []appidentity.KnownAppSmokeEvidenceSummary(nil)
	verifiedCatalog, err := loadKnownAppVerifiedCatalogForCenter("kde-center-page-preview", *knownAppVerifiedCatalog)
	if err != nil {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", nil, appidentity.KDECenterPageOptions{}, err
	}
	projectedEvidence, err := loadKnownAppSmokeEvidenceForCenter("kde-center-page-preview", *knownAppEvidenceJSON, *knownAppEvidenceFile, *knownAppMatrixReport, *knownAppGUISmokeReport, *knownAppGUISmokeApp, *knownAppGUISmokeName, *knownAppGUISmokeVersion)
	if err != nil {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", nil, appidentity.KDECenterPageOptions{}, err
	}
	legacyKnownAppSmoke := *knownAppSmokeApp != "" || *knownAppSmokeStatus != ""
	if len(projectedEvidence) > 0 {
		if legacyKnownAppSmoke {
			return appidentity.Recipe{}, appidentity.Provenance{}, "", nil, appidentity.KDECenterPageOptions{}, errors.New("kde-center-page-preview accepts either Runtime-projected known app evidence or legacy known app smoke flags, not both")
		}
		knownAppSmokeEvidence = projectedEvidence
	} else if legacyKnownAppSmoke {
		knownAppSmokeEvidence = []appidentity.KnownAppSmokeEvidenceSummary{{
			AppID:                                 *knownAppSmokeApp,
			DisplayName:                           *knownAppSmokeName,
			AppVersion:                            *knownAppSmokeVersion,
			EvidenceSource:                        *knownAppSmokeSource,
			SmokeStatus:                           *knownAppSmokeStatus,
			LaunchAuthorizationReceiptState:       *knownAppLaunchAuthorizationReceiptState,
			LaunchAuthorizationReceiptID:          *knownAppLaunchAuthorizationReceiptID,
			LaunchGateState:                       *knownAppLaunchGateState,
			LaunchGateConsumed:                    *knownAppLaunchGateConsumed,
			LaunchGateReceiptAccepted:             *knownAppLaunchGateReceiptAccepted,
			LaunchGateGuestBoundaryAccepted:       *knownAppLaunchGateGuestBoundaryAccepted,
			ControlledDispatchReady:               *knownAppControlledDispatchReady,
			ControlledExecutionSessionID:          *knownAppControlledExecutionSessionID,
			LauncherSessionGateConsumed:           *knownAppLauncherSessionGateConsumed,
			LauncherSessionDigestVerified:         *knownAppLauncherSessionDigestVerified,
			LauncherSessionRelativePath:           *knownAppLauncherSessionRelativePath,
			LauncherSessionRuntimeOwnerConsumable: *knownAppLauncherSessionRuntimeOwnerConsumable,
			LauncherSessionKDEReadModelConsumable: *knownAppLauncherSessionKDEReadModelConsumable,
			PostReviewDispatchConsumed:            *knownAppPostReviewDispatchConsumed,
			PostReviewDispatchState:               *knownAppPostReviewDispatchState,
			SessionGatedReviewReceiptID:           *knownAppSessionGatedReviewReceiptID,
			LaunchGateBlockedReason:               *knownAppLaunchGateBlockedReason,
			MarkerObserved:                        *knownAppSmokeMarkerObserved,
			ChecksumVerified:                      *knownAppSmokeChecksumVerified,
		}}
	}
	return recipe, provenance, *decision, fileURIs, appidentity.KDECenterPageOptions{
		ActivationRoot:                      *activationRoot,
		ExecutionSessionRoot:                *sessionRoot,
		ExecutionSessionRequestID:           *sessionRequestID,
		ApplicationReadinessRoot:            *readinessRoot,
		ApplicationReadinessStateRoot:       *readinessStateRoot,
		ApplicationReadinessArtifactReceipt: receipt,
		ApplicationReadinessPortalOperation: *readinessPortalOperation,
		ApplicationReadinessSnapshotReason:  *readinessSnapshotReason,
		KnownAppSmokeEvidence:               knownAppSmokeEvidence,
		KnownAppVerifiedCatalog:             verifiedCatalog,
	}, nil
}

func parseKDECenterPageSectionsPreviewSource(args []string) (appidentity.Recipe, appidentity.Provenance, string, []string, error) {
	return parseLaunchDecisionPreviewSource("kde-center-page-sections-preview", args)
}

func parseKDECenterPageSectionDetailPreviewSource(args []string) (appidentity.Recipe, appidentity.Provenance, string, string, []string, error) {
	flags := flag.NewFlagSet("kde-center-page-section-detail-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	applicationID := flags.String("app", "", "application id to load from the recipe registry")
	sectionID := flags.String("section", "", "KDE Compatibility Center section id")
	decision := flags.String("decision", "", "review decision: reviewed, approved, deferred, or rejected")
	registryPath := flags.String("registry", "", "path to an Xnix recipe registry JSON file")
	recipeRoot := flags.String("recipe-root", "", "directory containing registered recipe files; defaults to the registry directory")
	recipePath := flags.String("recipe", "", "path to an Xnix application recipe JSON file")
	if err := flags.Parse(args); err != nil {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", "", nil, err
	}
	if (*recipePath == "") == (*registryPath == "") {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", "", nil, errors.New("kde-center-page-section-detail-preview requires exactly one source: --recipe or --registry")
	}
	if *registryPath != "" && *applicationID == "" {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", "", nil, errors.New("kde-center-page-section-detail-preview requires --app when --registry is used")
	}
	if *sectionID == "" {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", "", nil, errors.New("kde-center-page-section-detail-preview requires --section")
	}
	if *decision == "" {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", "", nil, errors.New("kde-center-page-section-detail-preview requires --decision")
	}
	if *recipePath != "" && (*applicationID != "" || *recipeRoot != "") {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", "", nil, errors.New("kde-center-page-section-detail-preview --recipe cannot be combined with --app or --recipe-root")
	}

	recipe, provenance, err := loadRecipe(*recipePath, *registryPath, *recipeRoot, *applicationID)
	return recipe, provenance, *sectionID, *decision, flags.Args(), err
}

func parseKDEActionReviewPreviewSource(args []string) (appidentity.Recipe, appidentity.Provenance, string, string, []string, error) {
	flags := flag.NewFlagSet("kde-action-review-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	applicationID := flags.String("app", "", "application id to load from the recipe registry")
	actionID := flags.String("action", "", "KDE action queue item id")
	decision := flags.String("decision", "", "review decision: reviewed, approved, deferred, or rejected")
	registryPath := flags.String("registry", "", "path to an Xnix recipe registry JSON file")
	recipeRoot := flags.String("recipe-root", "", "directory containing registered recipe files; defaults to the registry directory")
	recipePath := flags.String("recipe", "", "path to an Xnix application recipe JSON file")
	if err := flags.Parse(args); err != nil {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", "", nil, err
	}
	if (*recipePath == "") == (*registryPath == "") {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", "", nil, errors.New("kde-action-review-preview requires exactly one source: --recipe or --registry")
	}
	if *registryPath != "" && *applicationID == "" {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", "", nil, errors.New("kde-action-review-preview requires --app when --registry is used")
	}
	if *actionID == "" {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", "", nil, errors.New("kde-action-review-preview requires --action")
	}
	if *decision == "" {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", "", nil, errors.New("kde-action-review-preview requires --decision")
	}
	if *recipePath != "" && (*applicationID != "" || *recipeRoot != "") {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", "", nil, errors.New("kde-action-review-preview --recipe cannot be combined with --app or --recipe-root")
	}

	recipe, provenance, err := loadRecipe(*recipePath, *registryPath, *recipeRoot, *applicationID)
	return recipe, provenance, *actionID, *decision, flags.Args(), err
}

func parseKDEActionPreflightPreviewSource(args []string) (appidentity.Recipe, appidentity.Provenance, string, string, []string, error) {
	flags := flag.NewFlagSet("kde-action-preflight-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	applicationID := flags.String("app", "", "application id to load from the recipe registry")
	actionID := flags.String("action", "", "KDE action queue item id")
	decision := flags.String("decision", "", "review decision: reviewed, approved, deferred, or rejected")
	registryPath := flags.String("registry", "", "path to an Xnix recipe registry JSON file")
	recipeRoot := flags.String("recipe-root", "", "directory containing registered recipe files; defaults to the registry directory")
	recipePath := flags.String("recipe", "", "path to an Xnix application recipe JSON file")
	if err := flags.Parse(args); err != nil {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", "", nil, err
	}
	if (*recipePath == "") == (*registryPath == "") {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", "", nil, errors.New("kde-action-preflight-preview requires exactly one source: --recipe or --registry")
	}
	if *registryPath != "" && *applicationID == "" {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", "", nil, errors.New("kde-action-preflight-preview requires --app when --registry is used")
	}
	if *actionID == "" {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", "", nil, errors.New("kde-action-preflight-preview requires --action")
	}
	if *decision == "" {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", "", nil, errors.New("kde-action-preflight-preview requires --decision")
	}
	if *recipePath != "" && (*applicationID != "" || *recipeRoot != "") {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", "", nil, errors.New("kde-action-preflight-preview --recipe cannot be combined with --app or --recipe-root")
	}

	recipe, provenance, err := loadRecipe(*recipePath, *registryPath, *recipeRoot, *applicationID)
	return recipe, provenance, *actionID, *decision, flags.Args(), err
}

func parseKDEActionReceiptPreviewSource(args []string) (appidentity.Recipe, appidentity.Provenance, string, string, []string, error) {
	flags := flag.NewFlagSet("kde-action-receipt-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	applicationID := flags.String("app", "", "application id to load from the recipe registry")
	actionID := flags.String("action", "", "KDE action queue item id")
	decision := flags.String("decision", "", "review decision: reviewed, approved, deferred, or rejected")
	registryPath := flags.String("registry", "", "path to an Xnix recipe registry JSON file")
	recipeRoot := flags.String("recipe-root", "", "directory containing registered recipe files; defaults to the registry directory")
	recipePath := flags.String("recipe", "", "path to an Xnix application recipe JSON file")
	if err := flags.Parse(args); err != nil {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", "", nil, err
	}
	if (*recipePath == "") == (*registryPath == "") {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", "", nil, errors.New("kde-action-receipt-preview requires exactly one source: --recipe or --registry")
	}
	if *registryPath != "" && *applicationID == "" {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", "", nil, errors.New("kde-action-receipt-preview requires --app when --registry is used")
	}
	if *actionID == "" {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", "", nil, errors.New("kde-action-receipt-preview requires --action")
	}
	if *decision == "" {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", "", nil, errors.New("kde-action-receipt-preview requires --decision")
	}
	if *recipePath != "" && (*applicationID != "" || *recipeRoot != "") {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", "", nil, errors.New("kde-action-receipt-preview --recipe cannot be combined with --app or --recipe-root")
	}

	recipe, provenance, err := loadRecipe(*recipePath, *registryPath, *recipeRoot, *applicationID)
	return recipe, provenance, *actionID, *decision, flags.Args(), err
}

func parseKDEActionStatusPreviewSource(args []string) (appidentity.Recipe, appidentity.Provenance, string, string, []string, error) {
	flags := flag.NewFlagSet("kde-action-status-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	applicationID := flags.String("app", "", "application id to load from the recipe registry")
	actionID := flags.String("action", "", "KDE action queue item id")
	decision := flags.String("decision", "", "review decision: reviewed, approved, deferred, or rejected")
	registryPath := flags.String("registry", "", "path to an Xnix recipe registry JSON file")
	recipeRoot := flags.String("recipe-root", "", "directory containing registered recipe files; defaults to the registry directory")
	recipePath := flags.String("recipe", "", "path to an Xnix application recipe JSON file")
	if err := flags.Parse(args); err != nil {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", "", nil, err
	}
	if (*recipePath == "") == (*registryPath == "") {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", "", nil, errors.New("kde-action-status-preview requires exactly one source: --recipe or --registry")
	}
	if *registryPath != "" && *applicationID == "" {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", "", nil, errors.New("kde-action-status-preview requires --app when --registry is used")
	}
	if *actionID == "" {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", "", nil, errors.New("kde-action-status-preview requires --action")
	}
	if *decision == "" {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", "", nil, errors.New("kde-action-status-preview requires --decision")
	}
	if *recipePath != "" && (*applicationID != "" || *recipeRoot != "") {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", "", nil, errors.New("kde-action-status-preview --recipe cannot be combined with --app or --recipe-root")
	}

	recipe, provenance, err := loadRecipe(*recipePath, *registryPath, *recipeRoot, *applicationID)
	return recipe, provenance, *actionID, *decision, flags.Args(), err
}

func parseKDEActionCardPreviewSource(args []string) (appidentity.Recipe, appidentity.Provenance, string, string, []string, error) {
	flags := flag.NewFlagSet("kde-action-card-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	applicationID := flags.String("app", "", "application id to load from the recipe registry")
	actionID := flags.String("action", "", "KDE action queue item id")
	decision := flags.String("decision", "", "review decision: reviewed, approved, deferred, or rejected")
	registryPath := flags.String("registry", "", "path to an Xnix recipe registry JSON file")
	recipeRoot := flags.String("recipe-root", "", "directory containing registered recipe files; defaults to the registry directory")
	recipePath := flags.String("recipe", "", "path to an Xnix application recipe JSON file")
	if err := flags.Parse(args); err != nil {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", "", nil, err
	}
	if (*recipePath == "") == (*registryPath == "") {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", "", nil, errors.New("kde-action-card-preview requires exactly one source: --recipe or --registry")
	}
	if *registryPath != "" && *applicationID == "" {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", "", nil, errors.New("kde-action-card-preview requires --app when --registry is used")
	}
	if *actionID == "" {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", "", nil, errors.New("kde-action-card-preview requires --action")
	}
	if *decision == "" {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", "", nil, errors.New("kde-action-card-preview requires --decision")
	}
	if *recipePath != "" && (*applicationID != "" || *recipeRoot != "") {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", "", nil, errors.New("kde-action-card-preview --recipe cannot be combined with --app or --recipe-root")
	}

	recipe, provenance, err := loadRecipe(*recipePath, *registryPath, *recipeRoot, *applicationID)
	return recipe, provenance, *actionID, *decision, flags.Args(), err
}

func parseKDEEntryPointActionPreviewSource(args []string) (appidentity.Recipe, appidentity.Provenance, string, string, []string, error) {
	flags := flag.NewFlagSet("kde-entrypoint-action-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	applicationID := flags.String("app", "", "application id to load from the recipe registry")
	entryPointID := flags.String("entrypoint", "", "KDE entrypoint id")
	decision := flags.String("decision", "", "review decision: reviewed, approved, deferred, or rejected")
	registryPath := flags.String("registry", "", "path to an Xnix recipe registry JSON file")
	recipeRoot := flags.String("recipe-root", "", "directory containing registered recipe files; defaults to the registry directory")
	recipePath := flags.String("recipe", "", "path to an Xnix application recipe JSON file")
	if err := flags.Parse(args); err != nil {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", "", nil, err
	}
	if (*recipePath == "") == (*registryPath == "") {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", "", nil, errors.New("kde-entrypoint-action-preview requires exactly one source: --recipe or --registry")
	}
	if *registryPath != "" && *applicationID == "" {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", "", nil, errors.New("kde-entrypoint-action-preview requires --app when --registry is used")
	}
	if *entryPointID == "" {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", "", nil, errors.New("kde-entrypoint-action-preview requires --entrypoint")
	}
	if *decision == "" {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", "", nil, errors.New("kde-entrypoint-action-preview requires --decision")
	}
	if *recipePath != "" && (*applicationID != "" || *recipeRoot != "") {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", "", nil, errors.New("kde-entrypoint-action-preview --recipe cannot be combined with --app or --recipe-root")
	}

	recipe, provenance, err := loadRecipe(*recipePath, *registryPath, *recipeRoot, *applicationID)
	return recipe, provenance, *entryPointID, *decision, flags.Args(), err
}

func parseLaunchDecisionPreviewSource(commandName string, args []string) (appidentity.Recipe, appidentity.Provenance, string, []string, error) {
	flags := flag.NewFlagSet(commandName, flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	applicationID := flags.String("app", "", "application id to load from the recipe registry")
	decision := flags.String("decision", "", "review decision: reviewed, approved, deferred, or rejected")
	registryPath := flags.String("registry", "", "path to an Xnix recipe registry JSON file")
	recipeRoot := flags.String("recipe-root", "", "directory containing registered recipe files; defaults to the registry directory")
	recipePath := flags.String("recipe", "", "path to an Xnix application recipe JSON file")
	if err := flags.Parse(args); err != nil {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", nil, err
	}
	if (*recipePath == "") == (*registryPath == "") {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", nil, fmt.Errorf("%s requires exactly one source: --recipe or --registry", commandName)
	}
	if *registryPath != "" && *applicationID == "" {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", nil, fmt.Errorf("%s requires --app when --registry is used", commandName)
	}
	if *decision == "" {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", nil, fmt.Errorf("%s requires --decision", commandName)
	}
	if *recipePath != "" && (*applicationID != "" || *recipeRoot != "") {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", nil, fmt.Errorf("%s --recipe cannot be combined with --app or --recipe-root", commandName)
	}

	recipe, provenance, err := loadRecipe(*recipePath, *registryPath, *recipeRoot, *applicationID)
	return recipe, provenance, *decision, flags.Args(), err
}

func parseLaunchActionPreviewSource(commandName string, args []string) (appidentity.Recipe, appidentity.Provenance, []string, error) {
	flags := flag.NewFlagSet(commandName, flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	applicationID := flags.String("app", "", "application id to load from the recipe registry")
	registryPath := flags.String("registry", "", "path to an Xnix recipe registry JSON file")
	recipeRoot := flags.String("recipe-root", "", "directory containing registered recipe files; defaults to the registry directory")
	recipePath := flags.String("recipe", "", "path to an Xnix application recipe JSON file")
	if err := flags.Parse(args); err != nil {
		return appidentity.Recipe{}, appidentity.Provenance{}, nil, err
	}
	if (*recipePath == "") == (*registryPath == "") {
		return appidentity.Recipe{}, appidentity.Provenance{}, nil, fmt.Errorf("%s requires exactly one source: --recipe or --registry", commandName)
	}
	if *registryPath != "" && *applicationID == "" {
		return appidentity.Recipe{}, appidentity.Provenance{}, nil, fmt.Errorf("%s requires --app when --registry is used", commandName)
	}
	if *recipePath != "" && (*applicationID != "" || *recipeRoot != "") {
		return appidentity.Recipe{}, appidentity.Provenance{}, nil, fmt.Errorf("%s --recipe cannot be combined with --app or --recipe-root", commandName)
	}

	recipe, provenance, err := loadRecipe(*recipePath, *registryPath, *recipeRoot, *applicationID)
	return recipe, provenance, flags.Args(), err
}

func parseRecipeSource(commandName string, args []string) (appidentity.Recipe, appidentity.Provenance, error) {
	flags := flag.NewFlagSet(commandName, flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	applicationID := flags.String("app", "", "application id to load from the recipe registry")
	registryPath := flags.String("registry", "", "path to an Xnix recipe registry JSON file")
	recipeRoot := flags.String("recipe-root", "", "directory containing registered recipe files; defaults to the registry directory")
	recipePath := flags.String("recipe", "", "path to an Xnix application recipe JSON file")
	if err := flags.Parse(args); err != nil {
		return appidentity.Recipe{}, appidentity.Provenance{}, err
	}
	if (*recipePath == "") == (*registryPath == "") {
		return appidentity.Recipe{}, appidentity.Provenance{}, fmt.Errorf("%s requires exactly one source: --recipe or --registry", commandName)
	}
	if *registryPath != "" && *applicationID == "" {
		return appidentity.Recipe{}, appidentity.Provenance{}, fmt.Errorf("%s requires --app when --registry is used", commandName)
	}
	if *recipePath != "" && (*applicationID != "" || *recipeRoot != "") {
		return appidentity.Recipe{}, appidentity.Provenance{}, fmt.Errorf("%s --recipe cannot be combined with --app or --recipe-root", commandName)
	}
	if flags.NArg() != 0 {
		return appidentity.Recipe{}, appidentity.Provenance{}, fmt.Errorf("%s does not accept positional arguments", commandName)
	}

	return loadRecipe(*recipePath, *registryPath, *recipeRoot, *applicationID)
}

func parseDesktopActivationPreflightPreviewSource(args []string) (appidentity.Recipe, appidentity.Provenance, string, error) {
	flags := flag.NewFlagSet("desktop-activation-preflight-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	applicationID := flags.String("app", "", "application id to load from the recipe registry")
	registryPath := flags.String("registry", "", "path to an Xnix recipe registry JSON file")
	recipeRoot := flags.String("recipe-root", "", "directory containing registered recipe files; defaults to the registry directory")
	recipePath := flags.String("recipe", "", "path to an Xnix application recipe JSON file")
	mode := flags.String("mode", "production", "activation preflight mode: production or development")
	if err := flags.Parse(args); err != nil {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", err
	}
	if (*recipePath == "") == (*registryPath == "") {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", errors.New("desktop-activation-preflight-preview requires exactly one source: --recipe or --registry")
	}
	if *registryPath != "" && *applicationID == "" {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", errors.New("desktop-activation-preflight-preview requires --app when --registry is used")
	}
	if *recipePath != "" && (*applicationID != "" || *recipeRoot != "") {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", errors.New("desktop-activation-preflight-preview --recipe cannot be combined with --app or --recipe-root")
	}
	if flags.NArg() != 0 {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", errors.New("desktop-activation-preflight-preview does not accept positional arguments")
	}

	recipe, provenance, err := loadRecipe(*recipePath, *registryPath, *recipeRoot, *applicationID)
	return recipe, provenance, *mode, err
}

func parseDesktopActivationStagingPreviewSource(args []string) (appidentity.Recipe, appidentity.Provenance, string, error) {
	flags := flag.NewFlagSet("desktop-activation-staging-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	applicationID := flags.String("app", "", "application id to load from the recipe registry")
	registryPath := flags.String("registry", "", "path to an Xnix recipe registry JSON file")
	recipeRoot := flags.String("recipe-root", "", "directory containing registered recipe files; defaults to the registry directory")
	recipePath := flags.String("recipe", "", "path to an Xnix application recipe JSON file")
	mode := flags.String("mode", "production", "activation staging mode: production or development")
	if err := flags.Parse(args); err != nil {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", err
	}
	if (*recipePath == "") == (*registryPath == "") {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", errors.New("desktop-activation-staging-preview requires exactly one source: --recipe or --registry")
	}
	if *registryPath != "" && *applicationID == "" {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", errors.New("desktop-activation-staging-preview requires --app when --registry is used")
	}
	if *recipePath != "" && (*applicationID != "" || *recipeRoot != "") {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", errors.New("desktop-activation-staging-preview --recipe cannot be combined with --app or --recipe-root")
	}
	if flags.NArg() != 0 {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", errors.New("desktop-activation-staging-preview does not accept positional arguments")
	}

	recipe, provenance, err := loadRecipe(*recipePath, *registryPath, *recipeRoot, *applicationID)
	return recipe, provenance, *mode, err
}

func parseDesktopActivationTransactionPreviewSource(args []string) (appidentity.Recipe, appidentity.Provenance, string, error) {
	flags := flag.NewFlagSet("desktop-activation-transaction-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	applicationID := flags.String("app", "", "application id to load from the recipe registry")
	registryPath := flags.String("registry", "", "path to an Xnix recipe registry JSON file")
	recipeRoot := flags.String("recipe-root", "", "directory containing registered recipe files; defaults to the registry directory")
	recipePath := flags.String("recipe", "", "path to an Xnix application recipe JSON file")
	mode := flags.String("mode", "production", "activation transaction mode: production or development")
	if err := flags.Parse(args); err != nil {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", err
	}
	if (*recipePath == "") == (*registryPath == "") {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", errors.New("desktop-activation-transaction-preview requires exactly one source: --recipe or --registry")
	}
	if *registryPath != "" && *applicationID == "" {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", errors.New("desktop-activation-transaction-preview requires --app when --registry is used")
	}
	if *recipePath != "" && (*applicationID != "" || *recipeRoot != "") {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", errors.New("desktop-activation-transaction-preview --recipe cannot be combined with --app or --recipe-root")
	}
	if flags.NArg() != 0 {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", errors.New("desktop-activation-transaction-preview does not accept positional arguments")
	}

	recipe, provenance, err := loadRecipe(*recipePath, *registryPath, *recipeRoot, *applicationID)
	return recipe, provenance, *mode, err
}

func parseDesktopActivationStatusPreviewSource(args []string) (appidentity.Recipe, appidentity.Provenance, string, string, error) {
	flags := flag.NewFlagSet("desktop-activation-status-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	applicationID := flags.String("app", "", "application id to load from the recipe registry")
	registryPath := flags.String("registry", "", "path to an Xnix recipe registry JSON file")
	recipeRoot := flags.String("recipe-root", "", "directory containing registered recipe files; defaults to the registry directory")
	recipePath := flags.String("recipe", "", "path to an Xnix application recipe JSON file")
	externalAppImportRecord := flags.String("external-app-import-record", "", "Runtime external Windows app import record")
	mode := flags.String("mode", "production", "activation status mode: production or development")
	activationRoot := flags.String("activation-root", "", "read staged desktop activation receipt evidence from this explicit root")
	if err := flags.Parse(args); err != nil {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", "", err
	}
	sourceCount := 0
	for _, source := range []string{*recipePath, *registryPath, *externalAppImportRecord} {
		if source != "" {
			sourceCount++
		}
	}
	if sourceCount != 1 {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", "", errors.New("desktop-activation-status-preview requires exactly one source: --recipe, --registry, or --external-app-import-record")
	}
	if *registryPath != "" && *applicationID == "" {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", "", errors.New("desktop-activation-status-preview requires --app when --registry is used")
	}
	if *recipePath != "" && (*applicationID != "" || *recipeRoot != "") {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", "", errors.New("desktop-activation-status-preview --recipe cannot be combined with --app or --recipe-root")
	}
	if *externalAppImportRecord != "" && (*applicationID != "" || *recipeRoot != "") {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", "", errors.New("desktop-activation-status-preview --external-app-import-record cannot be combined with --app or --recipe-root")
	}
	if flags.NArg() != 0 {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", "", errors.New("desktop-activation-status-preview does not accept positional arguments")
	}

	if *externalAppImportRecord != "" {
		record, err := appidentity.LoadExternalWinAppImportRecord(*externalAppImportRecord)
		if err != nil {
			return appidentity.Recipe{}, appidentity.Provenance{}, "", "", err
		}
		recipe, provenance, err := appidentity.ExternalAppRecipeFromImportRecord(record)
		return recipe, provenance, *mode, *activationRoot, err
	}
	recipe, provenance, err := loadRecipe(*recipePath, *registryPath, *recipeRoot, *applicationID)
	return recipe, provenance, *mode, *activationRoot, err
}

func loadRecipe(recipePath string, registryPath string, recipeRoot string, applicationID string) (appidentity.Recipe, appidentity.Provenance, error) {
	if registryPath != "" {
		return appidentity.LoadRecipeFromRegistry(registryPath, recipeRoot, applicationID)
	}

	data, err := os.ReadFile(recipePath)
	if err != nil {
		return appidentity.Recipe{}, appidentity.Provenance{}, fmt.Errorf("read recipe: %w", err)
	}
	recipe, err := appidentity.ParseRecipe(data)
	if err != nil {
		return appidentity.Recipe{}, appidentity.Provenance{}, err
	}
	return recipe, appidentity.Provenance{Source: "direct-file"}, nil
}
