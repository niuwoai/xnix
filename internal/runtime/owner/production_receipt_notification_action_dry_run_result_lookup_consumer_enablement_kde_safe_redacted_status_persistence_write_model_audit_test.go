package owner

import (
	"os"
	"path/filepath"
	"testing"
)

func TestProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusPersistenceWriteModelAuditPreviewModelsWriteShape(t *testing.T) {
	preview, err := NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusPersistenceWriteModelAuditPreview(projectRootForOwnerTest(t))
	if err != nil {
		t.Fatalf("NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusPersistenceWriteModelAuditPreview returned error: %v", err)
	}

	if preview.Version != currentProjectVersion(t) ||
		preview.SchemaVersion != "xnix.runtime.production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_kde_safe_redacted_status_persistence_write_model_audit.v1" ||
		preview.RequestType != "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-persistence-write-model-audit-preview" ||
		preview.AuditType != "receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-persistence-write-model-audit" ||
		preview.AuditDecision != "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-persistence-write-model-audit-ready-writes-disabled" {
		t.Fatalf("unexpected KDE-safe redacted status write-model schema: %#v", preview)
	}
	if !preview.WriteModelAuditRequired ||
		!preview.WriteModelModeled ||
		!preview.PersistenceAuthorizationAuditConsumed ||
		!preview.RedactedWriteModelGuidanceConsumed ||
		!preview.CompatibilityCenterWriteModelModeled ||
		!preview.RuntimeDiagnosticsWriteModelModeled ||
		!preview.WriteModelBoundaryReady ||
		!preview.PersistenceAuthorizationBoundaryReady ||
		!preview.StatusFanOutReady ||
		!preview.KDESafeStatusOnly ||
		!preview.OpaqueResultIDSupported ||
		!preview.RedactedSummaryShapeModeled ||
		preview.StatusPersistenceAuthorized ||
		preview.StatusPersistenceWriteEnabled ||
		preview.KDEStatusWriteEnabled ||
		preview.RuntimeDiagnosticsWriteEnabled ||
		preview.RedactedSummaryPersisted ||
		preview.KDEStatusPersisted ||
		preview.RuntimeDiagnosticsPersisted ||
		preview.ConsumerConsumptionAuthorized ||
		preview.ConsumerEnablementAuthorized ||
		preview.KDEConsumerEnabled ||
		preview.RuntimeConsumerEnabled ||
		preview.LookupRouteEnabled ||
		preview.OpaqueLookupEnabled ||
		preview.RawResultExposed ||
		preview.DispatchDryRunExecuted ||
		preview.ProductionReadiness ||
		preview.ProductionOwnershipReady {
		t.Fatalf("unexpected KDE-safe redacted status write-model decision: %#v", preview)
	}
	if preview.WriteModelItemCount != 10 ||
		preview.RequiredWriteModelItemCount != 10 ||
		preview.ReadyWriteModelItemCount != 10 ||
		preview.MissingWriteModelItemCount != 0 ||
		preview.CompatibilityCenterWriteModelItemCount != 5 ||
		preview.RuntimeDiagnosticsWriteModelItemCount != 5 ||
		preview.WriteEnabledItemCount != 0 ||
		preview.PersistedStatusItemCount != 0 ||
		preview.ConsumerEnabledWriteModelItemCount != 0 ||
		preview.RawExposedWriteModelItemCount != 0 ||
		preview.SideEffectWriteModelItemCount != 0 {
		t.Fatalf("unexpected KDE-safe redacted status write-model inventory: %#v", preview)
	}
	for _, item := range preview.WriteModelItems {
		if !item.EvidencePresent ||
			!item.PersistenceAuthorizationConsumed ||
			!item.WriteModelModeled ||
			!item.RedactedSummaryShapeModeled ||
			!item.KDESafeStatusOnly ||
			len(item.RedactedFields) != 7 ||
			item.StatusPersistenceAuthorized ||
			item.StatusPersistenceWriteEnabled ||
			item.KDEStatusWriteEnabled ||
			item.RuntimeDiagnosticsWriteEnabled ||
			item.RedactedSummaryPersisted ||
			item.KDEStatusPersisted ||
			item.RuntimeDiagnosticsPersisted ||
			item.DryRunResultPersisted ||
			item.ConsumerConsumptionAuthorized ||
			item.ConsumerEnablementAuthorized ||
			item.KDEConsumerEnabled ||
			item.RuntimeConsumerEnabled ||
			item.LookupRouteEnabled ||
			item.OpaqueLookupEnabled ||
			item.RawResultExposed ||
			item.DispatchDryRunExecuted ||
			!item.UserVisible ||
			!item.ReviewOnly ||
			!item.RuntimeOwned ||
			!item.GoRuntimeBacked ||
			item.KDEPolicyOwner ||
			item.RequestObjectCreated ||
			item.RequestObjectDispatched ||
			item.PortalRequestCreated ||
			item.NotificationActionEnabled ||
			item.CompatibilityCenterOpened ||
			item.SupportBundleExported ||
			item.SupportCaseCreated ||
			item.CallerStateRootRequired ||
			item.StateRootPathExposed ||
			item.FilePathsExposed ||
			item.FileContentRead ||
			!item.SideEffectsDisabled ||
			item.HostRootModified ||
			item.InternalDetailsExposed ||
			item.WriteModelStatus != "kde-safe-redacted-status-write-model-modeled-writes-disabled" {
			t.Fatalf("unsafe KDE-safe redacted status write-model item: %#v", item)
		}
		if item.SurfaceKind == "compatibility-center" && (!item.CompatibilityCenterWriteModelCandidate || item.RuntimeDiagnosticsWriteModelCandidate) {
			t.Fatalf("Compatibility Center write-model item must stay surface-specific: %#v", item)
		}
		if item.SurfaceKind == "runtime-diagnostics" && (!item.RuntimeDiagnosticsWriteModelCandidate || item.CompatibilityCenterWriteModelCandidate) {
			t.Fatalf("Runtime diagnostics write-model item must stay surface-specific: %#v", item)
		}
	}
	if preview.Counts.Total != 10 ||
		preview.Counts.Passed != 10 ||
		preview.Counts.Pending != 0 ||
		preview.Counts.Blocked != 0 ||
		!sameStrings(preview.CheckIDs, []string{"kde-safe-redacted-status-write-model-audit-required", "persistence-authorization-audit-consumed", "redacted-write-model-guidance-consumed", "compatibility-center-write-model-modeled", "runtime-diagnostics-write-model-modeled", "ten-kde-safe-redacted-status-write-model-items-present", "redacted-shape-modeled-only", "writes-consumers-and-lookup-disabled", "request-notification-navigation-and-support-disabled", "production-and-host-boundary-closed"}) {
		t.Fatalf("unexpected KDE-safe redacted status write-model checks: counts=%#v ids=%#v", preview.Counts, preview.CheckIDs)
	}
	if preview.StatusPersistenceWriteEnabled ||
		preview.KDEStatusWriteEnabled ||
		preview.RuntimeDiagnosticsWriteEnabled ||
		preview.RedactedSummaryPersisted ||
		preview.KDEStatusPersisted ||
		preview.RuntimeDiagnosticsPersisted ||
		preview.ConsumerEnablementAuthorized ||
		preview.KDEConsumerEnabled ||
		preview.RuntimeConsumerEnabled ||
		preview.LookupRouteEnabled ||
		preview.OpaqueLookupEnabled ||
		preview.DispatchDryRunExecuted ||
		preview.RequestObjectCreationEnabled ||
		preview.RequestObjectDispatchEnabled ||
		preview.PortalRequestCreated ||
		preview.NotificationActionEnabled ||
		preview.CompatibilityCenterOpened ||
		preview.SupportBundleExported ||
		preview.BackendLaunchEnabled ||
		preview.HostRootModified ||
		preview.BackendDetailsExposed {
		t.Fatalf("unsafe KDE-safe redacted status write-model gates: %#v", preview)
	}
	if err := validateNoBackendTerms(preview, "KDE-safe redacted status write-model test"); err != nil {
		t.Fatalf("validateNoBackendTerms returned error: %v", err)
	}
}

func TestProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusPersistenceWriteModelAuditPreviewFailsClosedWithoutSources(t *testing.T) {
	root := t.TempDir()
	if err := os.WriteFile(filepath.Join(root, "VERSION"), []byte(currentProjectVersion(t)+"\n"), 0o600); err != nil {
		t.Fatalf("WriteFile VERSION returned error: %v", err)
	}
	preview, err := NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusPersistenceWriteModelAuditPreview(root)
	if err != nil {
		t.Fatalf("NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusPersistenceWriteModelAuditPreview returned error: %v", err)
	}
	if !preview.WriteModelAuditRequired ||
		!preview.WriteModelModeled ||
		preview.PersistenceAuthorizationAuditConsumed ||
		preview.RedactedWriteModelGuidanceConsumed ||
		preview.CompatibilityCenterWriteModelModeled ||
		preview.RuntimeDiagnosticsWriteModelModeled ||
		preview.WriteModelBoundaryReady ||
		preview.RedactedSummaryShapeModeled ||
		preview.KDESafeStatusOnly ||
		preview.OpaqueResultIDSupported ||
		preview.StatusPersistenceWriteEnabled ||
		preview.KDEStatusPersisted ||
		preview.RuntimeDiagnosticsPersisted ||
		preview.RawResultExposed ||
		preview.WriteModelItemCount != 10 ||
		preview.RequiredWriteModelItemCount != 10 ||
		preview.ReadyWriteModelItemCount != 0 ||
		preview.MissingWriteModelItemCount != 10 ||
		preview.CompatibilityCenterWriteModelItemCount != 5 ||
		preview.RuntimeDiagnosticsWriteModelItemCount != 5 ||
		preview.Counts.Total != 10 ||
		preview.Counts.Passed != 4 ||
		preview.Counts.Blocked != 6 ||
		preview.AuditDecision != "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-persistence-write-model-audit-blocked" {
		t.Fatalf("missing KDE-safe redacted status write-model sources must fail closed: %#v", preview)
	}
	for _, item := range preview.WriteModelItems {
		if item.EvidencePresent ||
			item.PersistenceAuthorizationConsumed ||
			item.WriteModelModeled ||
			item.RedactedSummaryShapeModeled ||
			item.KDESafeStatusOnly ||
			item.StatusPersistenceWriteEnabled ||
			item.KDEStatusPersisted ||
			item.RuntimeDiagnosticsPersisted ||
			item.UserVisible ||
			item.RawResultExposed ||
			item.WriteModelStatus != "missing-kde-safe-redacted-status-write-model-evidence" {
			t.Fatalf("missing KDE-safe redacted status write-model item must remain closed: %#v", item)
		}
	}
}
