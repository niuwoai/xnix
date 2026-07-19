package owner

import (
	"os"
	"path/filepath"
	"testing"
)

func TestProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusStoragePersistenceGateAuditPreviewModelsGate(t *testing.T) {
	preview, err := NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusStoragePersistenceGateAuditPreview(projectRootForOwnerTest(t))
	if err != nil {
		t.Fatalf("NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusStoragePersistenceGateAuditPreview returned error: %v", err)
	}

	if preview.Version != currentProjectVersion(t) ||
		preview.SchemaVersion != "xnix.runtime.production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_kde_safe_redacted_status_storage_persistence_gate_audit.v1" ||
		preview.RequestType != "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-storage-persistence-gate-audit-preview" ||
		preview.PreviewType != "receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-storage-persistence-gate-audit" ||
		preview.PreviewDecision != "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-storage-persistence-gate-audit-ready-storage-disabled" {
		t.Fatalf("unexpected KDE-safe redacted status storage persistence gate schema: %#v", preview)
	}
	if !preview.CurrentMainlineConsumed ||
		!preview.ClosedPersistenceWriterConsumed ||
		!preview.ClosedPersistenceWriterReady ||
		!preview.StoragePersistenceGateRequired ||
		!preview.StoragePersistenceGateModeled ||
		!preview.StoragePersistenceGateReady ||
		!preview.GateInputBoundaryModeled ||
		!preview.GateOutputBoundaryModeled ||
		!preview.KDESafeRedactedStatusOnly ||
		!preview.CompatibilityCenterGateModeled ||
		!preview.RuntimeDiagnosticsGateModeled ||
		preview.WriterCallable ||
		preview.ClosedPersistenceWriterEnabled ||
		preview.StorageGatePassed ||
		preview.StoragePersistenceAuthorized ||
		preview.WriterAuthorizationGranted ||
		preview.StatusWriterEnabled ||
		preview.StatusPersistenceAuthorized ||
		preview.StatusPersistenceWriteEnabled ||
		preview.KDEStatusWriteEnabled ||
		preview.RuntimeDiagnosticsWriteEnabled ||
		preview.StorageWriteEnabled ||
		preview.ConsumerConsumptionAuthorized ||
		preview.ConsumerEnablementAuthorized ||
		preview.KDEConsumerEnabled ||
		preview.RuntimeConsumerEnabled ||
		preview.LookupRouteEnabled ||
		preview.OpaqueLookupEnabled ||
		preview.RedactedSummaryPersisted ||
		preview.KDEStatusPersisted ||
		preview.RuntimeDiagnosticsPersisted ||
		preview.DryRunResultPersisted ||
		preview.RawResultExposed ||
		preview.DispatchDryRunExecuted ||
		preview.RequestObjectCreationEnabled ||
		preview.RequestObjectDispatchEnabled ||
		preview.PortalRequestCreated ||
		preview.NotificationActionEnabled ||
		preview.CompatibilityCenterOpened ||
		preview.SupportBundleExported ||
		preview.SupportCaseCreated ||
		preview.ProductionReadiness ||
		preview.ProductionOwnershipReady {
		t.Fatalf("unsafe KDE-safe redacted status storage persistence gate decision: %#v", preview)
	}
	if preview.GateItemCount != 10 ||
		preview.RequiredGateItemCount != 10 ||
		preview.ReadyGateItemCount != 10 ||
		preview.MissingGateItemCount != 0 ||
		preview.PassedStorageGateItemCount != 0 ||
		preview.AuthorizedStorageItemCount != 0 ||
		preview.CallableWriterItemCount != 0 ||
		preview.EnabledPersistenceWriterItemCount != 0 ||
		preview.PersistedWriterItemCount != 0 ||
		preview.RawExposedGateItemCount != 0 ||
		preview.SideEffectGateItemCount != 0 ||
		preview.CompatibilityCenterItemCount != 5 ||
		preview.RuntimeDiagnosticsItemCount != 5 {
		t.Fatalf("unexpected KDE-safe redacted status storage persistence gate counts: %#v", preview)
	}
	if len(preview.GateItems) != 10 ||
		!sameStrings(preview.GateItemIDs, []string{
			"review-receipt-dry-run-result-lookup-consumer-enablement-kde-safe-compatibility-center-redacted-status-storage-persistence-gate",
			"review-receipt-dry-run-result-lookup-consumer-enablement-kde-safe-runtime-diagnostics-redacted-status-storage-persistence-gate",
			"renew-receipt-dry-run-result-lookup-consumer-enablement-kde-safe-compatibility-center-redacted-status-storage-persistence-gate",
			"renew-receipt-dry-run-result-lookup-consumer-enablement-kde-safe-runtime-diagnostics-redacted-status-storage-persistence-gate",
			"open-compatibility-center-dry-run-result-lookup-consumer-enablement-kde-safe-compatibility-center-redacted-status-storage-persistence-gate",
			"open-compatibility-center-dry-run-result-lookup-consumer-enablement-kde-safe-runtime-diagnostics-redacted-status-storage-persistence-gate",
			"dismiss-receipt-dry-run-result-lookup-consumer-enablement-kde-safe-compatibility-center-redacted-status-storage-persistence-gate",
			"dismiss-receipt-dry-run-result-lookup-consumer-enablement-kde-safe-runtime-diagnostics-redacted-status-storage-persistence-gate",
			"support-info-dry-run-result-lookup-consumer-enablement-kde-safe-compatibility-center-redacted-status-storage-persistence-gate",
			"support-info-dry-run-result-lookup-consumer-enablement-kde-safe-runtime-diagnostics-redacted-status-storage-persistence-gate",
		}) {
		t.Fatalf("unexpected KDE-safe redacted status storage persistence gate items: %#v", preview.GateItemIDs)
	}
	for _, item := range preview.GateItems {
		if !item.EvidencePresent ||
			!item.CurrentMainlineConsumed ||
			!item.ClosedPersistenceWriterConsumed ||
			!item.ClosedPersistenceWriterReady ||
			!item.StoragePersistenceGateModeled ||
			!item.GateInputBoundaryModeled ||
			!item.GateOutputBoundaryModeled ||
			!item.KDESafeRedactedStatusOnly ||
			item.WriterCallable ||
			item.ClosedPersistenceWriterEnabled ||
			item.StorageGatePassed ||
			item.StoragePersistenceAuthorized ||
			item.StatusPersistenceAuthorized ||
			item.StatusPersistenceWriteEnabled ||
			item.StatusWriterEnabled ||
			item.KDEStatusWriteEnabled ||
			item.RuntimeDiagnosticsWriteEnabled ||
			item.StorageWriteEnabled ||
			item.RedactedSummaryPersisted ||
			item.KDEStatusPersisted ||
			item.RuntimeDiagnosticsPersisted ||
			item.RawResultExposed ||
			!item.UserVisible ||
			!item.ReviewOnly ||
			!item.RuntimeOwned ||
			!item.GoRuntimeBacked ||
			item.KDEPolicyOwner ||
			!item.SideEffectsDisabled ||
			item.HostRootModified ||
			item.InternalDetailsExposed ||
			item.StoragePersistenceGateStatus != "redacted-status-storage-persistence-gate-modeled-storage-disabled" {
			t.Fatalf("unsafe KDE-safe redacted status storage persistence gate item: %#v", item)
		}
	}
	if preview.Counts.Total != 8 ||
		preview.Counts.Passed != 8 ||
		preview.Counts.Pending != 0 ||
		preview.Counts.Blocked != 0 ||
		!sameStrings(preview.CheckIDs, []string{"current-mainline-consumed", "closed-persistence-writer-consumed", "storage-persistence-gate-modeled", "compatibility-center-and-runtime-storage-gates-modeled", "ten-gate-items-ready-storage-disabled", "storage-persistence-writes-disabled", "consumer-lookup-request-and-support-disabled", "production-and-host-boundary-closed"}) {
		t.Fatalf("unexpected KDE-safe redacted status storage persistence gate checks: counts=%#v ids=%#v", preview.Counts, preview.CheckIDs)
	}
	if preview.SystemServiceStarted ||
		preview.SessionBusClaimed ||
		preview.ProductionBusClaimed ||
		preview.ProductionOwnerEnabled ||
		preview.WriteMethodsEnabled ||
		preview.RuntimeWritesEnabled ||
		preview.BackendLaunchEnabled ||
		preview.HostRootModified ||
		preview.BackendDetailsExposed {
		t.Fatalf("unsafe KDE-safe redacted status storage persistence gate gates: %#v", preview)
	}
	if err := validateNoBackendTerms(preview, "KDE-safe redacted status storage persistence gate test"); err != nil {
		t.Fatalf("validateNoBackendTerms returned error: %v", err)
	}
}

func TestProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusStoragePersistenceGateAuditPreviewFailsClosedWithoutSources(t *testing.T) {
	root := t.TempDir()
	if err := os.WriteFile(filepath.Join(root, "VERSION"), []byte(currentProjectVersion(t)+"\n"), 0o600); err != nil {
		t.Fatalf("WriteFile VERSION returned error: %v", err)
	}
	preview, err := NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusStoragePersistenceGateAuditPreview(root)
	if err != nil {
		t.Fatalf("NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusStoragePersistenceGateAuditPreview returned error: %v", err)
	}
	if preview.CurrentMainlineConsumed ||
		preview.ClosedPersistenceWriterConsumed ||
		preview.ClosedPersistenceWriterReady ||
		!preview.StoragePersistenceGateRequired ||
		preview.StoragePersistenceGateModeled ||
		preview.StoragePersistenceGateReady ||
		preview.GateInputBoundaryModeled ||
		preview.GateOutputBoundaryModeled ||
		preview.KDESafeRedactedStatusOnly ||
		preview.CompatibilityCenterGateModeled ||
		preview.RuntimeDiagnosticsGateModeled ||
		preview.WriterCallable ||
		preview.ClosedPersistenceWriterEnabled ||
		preview.StorageGatePassed ||
		preview.StoragePersistenceAuthorized ||
		preview.StatusPersistenceAuthorized ||
		preview.StatusPersistenceWriteEnabled ||
		preview.StorageWriteEnabled ||
		preview.RawResultExposed ||
		preview.GateItemCount != 10 ||
		preview.RequiredGateItemCount != 10 ||
		preview.ReadyGateItemCount != 0 ||
		preview.MissingGateItemCount != 10 ||
		preview.Counts.Total != 8 ||
		preview.Counts.Passed != 3 ||
		preview.Counts.Blocked != 5 ||
		preview.PreviewDecision != "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-storage-persistence-gate-audit-blocked" {
		t.Fatalf("missing KDE-safe redacted status storage persistence gate sources must fail closed: %#v", preview)
	}
	for _, item := range preview.GateItems {
		if item.EvidencePresent ||
			item.CurrentMainlineConsumed ||
			item.ClosedPersistenceWriterConsumed ||
			item.ClosedPersistenceWriterReady ||
			item.StoragePersistenceGateModeled ||
			item.GateInputBoundaryModeled ||
			item.GateOutputBoundaryModeled ||
			item.WriterCallable ||
			item.ClosedPersistenceWriterEnabled ||
			item.StorageGatePassed ||
			item.StoragePersistenceAuthorized ||
			item.StatusPersistenceAuthorized ||
			item.StatusPersistenceWriteEnabled ||
			item.StorageWriteEnabled ||
			item.RawResultExposed ||
			item.UserVisible ||
			item.StoragePersistenceGateStatus != "missing-redacted-status-storage-persistence-gate-evidence" {
			t.Fatalf("missing KDE-safe redacted status storage persistence gate item must remain closed: %#v", item)
		}
	}
}
