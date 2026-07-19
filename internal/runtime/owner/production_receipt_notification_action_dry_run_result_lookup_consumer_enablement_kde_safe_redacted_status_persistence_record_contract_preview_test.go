package owner

import (
	"os"
	"path/filepath"
	"testing"
)

func TestProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusPersistenceRecordContractPreviewModelsRecord(t *testing.T) {
	preview, err := NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusPersistenceRecordContractPreview(projectRootForOwnerTest(t))
	if err != nil {
		t.Fatalf("NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusPersistenceRecordContractPreview returned error: %v", err)
	}

	if preview.Version != currentProjectVersion(t) ||
		preview.SchemaVersion != "xnix.runtime.production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_kde_safe_redacted_status_persistence_record_contract.v1" ||
		preview.RequestType != "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-persistence-record-contract-preview" ||
		preview.PreviewType != "receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-persistence-record-contract" ||
		preview.PreviewDecision != "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-persistence-record-contract-ready-writes-disabled" {
		t.Fatalf("unexpected KDE-safe redacted status persistence record contract schema: %#v", preview)
	}
	if !preview.CurrentMainlineConsumed ||
		!preview.StoragePersistenceGateConsumed ||
		!preview.StoragePersistenceGateReady ||
		!preview.PersistenceRecordContractRequired ||
		!preview.PersistenceRecordContractModeled ||
		!preview.PersistenceRecordContractReady ||
		!preview.RecordShapeModeled ||
		!preview.RecordIdentityBoundaryModeled ||
		!preview.RecordRedactionBoundaryModeled ||
		!preview.KDESafeRedactedStatusOnly ||
		!preview.CompatibilityCenterRecordModeled ||
		!preview.RuntimeDiagnosticsRecordModeled ||
		preview.WriterCallable ||
		preview.ClosedPersistenceWriterEnabled ||
		preview.StorageGatePassed ||
		preview.StoragePersistenceAuthorized ||
		preview.DurableRecordWriteEnabled ||
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
		t.Fatalf("unsafe KDE-safe redacted status persistence record contract decision: %#v", preview)
	}
	if preview.RecordItemCount != 10 ||
		preview.RequiredRecordItemCount != 10 ||
		preview.ReadyRecordItemCount != 10 ||
		preview.MissingRecordItemCount != 0 ||
		preview.DurableRecordItemCount != 0 ||
		preview.WritableRecordItemCount != 0 ||
		preview.PersistedRecordItemCount != 0 ||
		preview.RawExposedRecordItemCount != 0 ||
		preview.SideEffectRecordItemCount != 0 ||
		preview.CompatibilityCenterItemCount != 5 ||
		preview.RuntimeDiagnosticsItemCount != 5 {
		t.Fatalf("unexpected KDE-safe redacted status persistence record contract counts: %#v", preview)
	}
	if len(preview.RecordItems) != 10 ||
		!sameStrings(preview.RecordItemIDs, []string{
			"review-receipt-dry-run-result-lookup-consumer-enablement-kde-safe-compatibility-center-redacted-status-persistence-record-contract",
			"review-receipt-dry-run-result-lookup-consumer-enablement-kde-safe-runtime-diagnostics-redacted-status-persistence-record-contract",
			"renew-receipt-dry-run-result-lookup-consumer-enablement-kde-safe-compatibility-center-redacted-status-persistence-record-contract",
			"renew-receipt-dry-run-result-lookup-consumer-enablement-kde-safe-runtime-diagnostics-redacted-status-persistence-record-contract",
			"open-compatibility-center-dry-run-result-lookup-consumer-enablement-kde-safe-compatibility-center-redacted-status-persistence-record-contract",
			"open-compatibility-center-dry-run-result-lookup-consumer-enablement-kde-safe-runtime-diagnostics-redacted-status-persistence-record-contract",
			"dismiss-receipt-dry-run-result-lookup-consumer-enablement-kde-safe-compatibility-center-redacted-status-persistence-record-contract",
			"dismiss-receipt-dry-run-result-lookup-consumer-enablement-kde-safe-runtime-diagnostics-redacted-status-persistence-record-contract",
			"support-info-dry-run-result-lookup-consumer-enablement-kde-safe-compatibility-center-redacted-status-persistence-record-contract",
			"support-info-dry-run-result-lookup-consumer-enablement-kde-safe-runtime-diagnostics-redacted-status-persistence-record-contract",
		}) {
		t.Fatalf("unexpected KDE-safe redacted status persistence record contract items: %#v", preview.RecordItemIDs)
	}
	for _, item := range preview.RecordItems {
		if !item.EvidencePresent ||
			!item.RecordIDModeled ||
			!item.OpaqueLookupIDModeled ||
			!item.RedactedSummaryModeled ||
			!item.ConsumerProjectionModeled ||
			!item.CurrentMainlineConsumed ||
			!item.StoragePersistenceGateConsumed ||
			!item.StoragePersistenceGateReady ||
			!item.PersistenceRecordContractModeled ||
			!item.RecordShapeModeled ||
			!item.RecordIdentityBoundaryModeled ||
			!item.RecordRedactionBoundaryModeled ||
			!item.KDESafeRedactedStatusOnly ||
			item.WriterCallable ||
			item.ClosedPersistenceWriterEnabled ||
			item.StorageGatePassed ||
			item.StoragePersistenceAuthorized ||
			item.DurableRecordWriteEnabled ||
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
			item.PersistenceRecordContractStatus != "redacted-status-persistence-record-contract-modeled-writes-disabled" {
			t.Fatalf("unsafe KDE-safe redacted status persistence record contract item: %#v", item)
		}
	}
	if preview.Counts.Total != 8 ||
		preview.Counts.Passed != 8 ||
		preview.Counts.Pending != 0 ||
		preview.Counts.Blocked != 0 ||
		!sameStrings(preview.CheckIDs, []string{"current-mainline-consumed", "storage-persistence-gate-consumed", "persistence-record-contract-modeled", "compatibility-center-and-runtime-record-contracts-modeled", "ten-record-items-ready-writes-disabled", "persistence-record-writes-disabled", "consumer-lookup-request-and-support-disabled", "production-and-host-boundary-closed"}) {
		t.Fatalf("unexpected KDE-safe redacted status persistence record contract checks: counts=%#v ids=%#v", preview.Counts, preview.CheckIDs)
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
		t.Fatalf("unsafe KDE-safe redacted status persistence record contract gates: %#v", preview)
	}
	if err := validateNoBackendTerms(preview, "KDE-safe redacted status persistence record contract test"); err != nil {
		t.Fatalf("validateNoBackendTerms returned error: %v", err)
	}
}

func TestProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusPersistenceRecordContractPreviewFailsClosedWithoutSources(t *testing.T) {
	root := t.TempDir()
	if err := os.WriteFile(filepath.Join(root, "VERSION"), []byte(currentProjectVersion(t)+"\n"), 0o600); err != nil {
		t.Fatalf("WriteFile VERSION returned error: %v", err)
	}
	preview, err := NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusPersistenceRecordContractPreview(root)
	if err != nil {
		t.Fatalf("NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusPersistenceRecordContractPreview returned error: %v", err)
	}
	if preview.CurrentMainlineConsumed ||
		preview.StoragePersistenceGateConsumed ||
		preview.StoragePersistenceGateReady ||
		!preview.PersistenceRecordContractRequired ||
		preview.PersistenceRecordContractModeled ||
		preview.PersistenceRecordContractReady ||
		preview.RecordShapeModeled ||
		preview.RecordIdentityBoundaryModeled ||
		preview.RecordRedactionBoundaryModeled ||
		preview.KDESafeRedactedStatusOnly ||
		preview.CompatibilityCenterRecordModeled ||
		preview.RuntimeDiagnosticsRecordModeled ||
		preview.WriterCallable ||
		preview.ClosedPersistenceWriterEnabled ||
		preview.StorageGatePassed ||
		preview.StoragePersistenceAuthorized ||
		preview.DurableRecordWriteEnabled ||
		preview.StatusPersistenceAuthorized ||
		preview.StatusPersistenceWriteEnabled ||
		preview.StorageWriteEnabled ||
		preview.RawResultExposed ||
		preview.RecordItemCount != 10 ||
		preview.RequiredRecordItemCount != 10 ||
		preview.ReadyRecordItemCount != 0 ||
		preview.MissingRecordItemCount != 10 ||
		preview.Counts.Total != 8 ||
		preview.Counts.Passed != 3 ||
		preview.Counts.Blocked != 5 ||
		preview.PreviewDecision != "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-persistence-record-contract-blocked" {
		t.Fatalf("missing KDE-safe redacted status persistence record contract sources must fail closed: %#v", preview)
	}
	for _, item := range preview.RecordItems {
		if item.EvidencePresent ||
			item.RecordIDModeled ||
			item.OpaqueLookupIDModeled ||
			item.RedactedSummaryModeled ||
			item.ConsumerProjectionModeled ||
			item.CurrentMainlineConsumed ||
			item.StoragePersistenceGateConsumed ||
			item.StoragePersistenceGateReady ||
			item.PersistenceRecordContractModeled ||
			item.RecordShapeModeled ||
			item.RecordIdentityBoundaryModeled ||
			item.RecordRedactionBoundaryModeled ||
			item.WriterCallable ||
			item.ClosedPersistenceWriterEnabled ||
			item.StorageGatePassed ||
			item.StoragePersistenceAuthorized ||
			item.DurableRecordWriteEnabled ||
			item.StatusPersistenceAuthorized ||
			item.StatusPersistenceWriteEnabled ||
			item.StorageWriteEnabled ||
			item.RawResultExposed ||
			item.UserVisible ||
			item.PersistenceRecordContractStatus != "missing-redacted-status-persistence-record-contract-evidence" {
			t.Fatalf("missing KDE-safe redacted status persistence record contract item must remain closed: %#v", item)
		}
	}
}
