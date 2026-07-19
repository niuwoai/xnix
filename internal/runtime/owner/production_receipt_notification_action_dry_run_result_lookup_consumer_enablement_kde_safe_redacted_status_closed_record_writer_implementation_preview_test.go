package owner

import (
	"os"
	"path/filepath"
	"testing"
)

func TestProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterImplementationPreviewModelsWriter(t *testing.T) {
	preview, err := NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterImplementationPreview(projectRootForOwnerTest(t))
	if err != nil {
		t.Fatalf("NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterImplementationPreview returned error: %v", err)
	}

	if preview.Version != currentProjectVersion(t) ||
		preview.SchemaVersion != "xnix.runtime.production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_kde_safe_redacted_status_closed_record_writer_implementation.v1" ||
		preview.RequestType != "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-implementation-preview" ||
		preview.PreviewType != "receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-implementation" ||
		preview.PreviewDecision != "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-implementation-ready-writes-disabled" {
		t.Fatalf("unexpected KDE-safe redacted status closed record writer implementation schema: %#v", preview)
	}
	if !preview.CurrentMainlineConsumed ||
		!preview.PersistenceRecordContractConsumed ||
		!preview.PersistenceRecordContractReady ||
		!preview.ClosedRecordWriterImplementationRequired ||
		!preview.ClosedRecordWriterImplementationModeled ||
		!preview.ClosedRecordWriterImplementationReady ||
		!preview.RecordWriterShapeModeled ||
		!preview.RecordWriterInputBoundaryModeled ||
		!preview.RecordWriterOutputBoundaryModeled ||
		!preview.DurableRecordContractConsumed ||
		!preview.KDESafeRedactedStatusOnly ||
		!preview.CompatibilityCenterWriterModeled ||
		!preview.RuntimeDiagnosticsWriterModeled ||
		preview.WriterCallable ||
		preview.ClosedRecordWriterEnabled ||
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
		t.Fatalf("unsafe KDE-safe redacted status closed record writer implementation decision: %#v", preview)
	}
	if preview.ImplementationItemCount != 10 ||
		preview.RequiredImplementationItemCount != 10 ||
		preview.ReadyImplementationItemCount != 10 ||
		preview.MissingImplementationItemCount != 0 ||
		preview.CallableRecordWriterItemCount != 0 ||
		preview.EnabledRecordWriterItemCount != 0 ||
		preview.DurableWrittenRecordItemCount != 0 ||
		preview.PersistedRecordItemCount != 0 ||
		preview.RawExposedImplementationItemCount != 0 ||
		preview.SideEffectImplementationItemCount != 0 ||
		preview.CompatibilityCenterItemCount != 5 ||
		preview.RuntimeDiagnosticsItemCount != 5 {
		t.Fatalf("unexpected KDE-safe redacted status closed record writer implementation counts: %#v", preview)
	}
	if len(preview.ImplementationItems) != 10 {
		t.Fatalf("unexpected KDE-safe redacted status closed record writer implementation item list: %#v", preview.ImplementationItemIDs)
	}
	for _, item := range preview.ImplementationItems {
		if !item.EvidencePresent ||
			!item.CurrentMainlineConsumed ||
			!item.PersistenceRecordContractConsumed ||
			!item.PersistenceRecordContractReady ||
			!item.ClosedRecordWriterImplementationModeled ||
			!item.RecordWriterShapeModeled ||
			!item.RecordWriterInputBoundaryModeled ||
			!item.RecordWriterOutputBoundaryModeled ||
			!item.DurableRecordContractConsumed ||
			!item.KDESafeRedactedStatusOnly ||
			item.WriterMethodName != "PreviewRedactedStatusRecordWriter" ||
			item.WriterCallable ||
			item.ClosedRecordWriterEnabled ||
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
			item.ClosedRecordWriterImplementationStatus != "redacted-status-closed-record-writer-implementation-modeled-writes-disabled" {
			t.Fatalf("unsafe KDE-safe redacted status closed record writer implementation item: %#v", item)
		}
	}
	if preview.Counts.Total != 8 ||
		preview.Counts.Passed != 8 ||
		preview.Counts.Pending != 0 ||
		preview.Counts.Blocked != 0 ||
		!sameStrings(preview.CheckIDs, []string{"current-mainline-consumed", "persistence-record-contract-consumed", "closed-record-writer-implementation-modeled", "compatibility-center-and-runtime-record-writers-modeled", "ten-implementation-items-ready-writes-disabled", "closed-record-writer-writes-disabled", "consumer-lookup-request-and-support-disabled", "production-and-host-boundary-closed"}) {
		t.Fatalf("unexpected KDE-safe redacted status closed record writer implementation checks: counts=%#v ids=%#v", preview.Counts, preview.CheckIDs)
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
		t.Fatalf("unsafe KDE-safe redacted status closed record writer implementation gates: %#v", preview)
	}
	if err := validateNoBackendTerms(preview, "KDE-safe redacted status closed record writer implementation test"); err != nil {
		t.Fatalf("validateNoBackendTerms returned error: %v", err)
	}
}

func TestProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterImplementationPreviewFailsClosedWithoutSources(t *testing.T) {
	root := t.TempDir()
	if err := os.WriteFile(filepath.Join(root, "VERSION"), []byte(currentProjectVersion(t)+"\n"), 0o600); err != nil {
		t.Fatalf("WriteFile VERSION returned error: %v", err)
	}
	preview, err := NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterImplementationPreview(root)
	if err != nil {
		t.Fatalf("NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterImplementationPreview returned error: %v", err)
	}
	if preview.CurrentMainlineConsumed ||
		preview.PersistenceRecordContractConsumed ||
		preview.PersistenceRecordContractReady ||
		!preview.ClosedRecordWriterImplementationRequired ||
		preview.ClosedRecordWriterImplementationModeled ||
		preview.ClosedRecordWriterImplementationReady ||
		preview.RecordWriterShapeModeled ||
		preview.RecordWriterInputBoundaryModeled ||
		preview.RecordWriterOutputBoundaryModeled ||
		preview.DurableRecordContractConsumed ||
		preview.KDESafeRedactedStatusOnly ||
		preview.CompatibilityCenterWriterModeled ||
		preview.RuntimeDiagnosticsWriterModeled ||
		preview.WriterCallable ||
		preview.ClosedRecordWriterEnabled ||
		preview.StorageGatePassed ||
		preview.StoragePersistenceAuthorized ||
		preview.DurableRecordWriteEnabled ||
		preview.StatusPersistenceAuthorized ||
		preview.StatusPersistenceWriteEnabled ||
		preview.StorageWriteEnabled ||
		preview.RawResultExposed ||
		preview.ImplementationItemCount != 10 ||
		preview.RequiredImplementationItemCount != 10 ||
		preview.ReadyImplementationItemCount != 0 ||
		preview.MissingImplementationItemCount != 10 ||
		preview.Counts.Total != 8 ||
		preview.Counts.Passed != 3 ||
		preview.Counts.Blocked != 5 ||
		preview.PreviewDecision != "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-implementation-blocked" {
		t.Fatalf("missing KDE-safe redacted status closed record writer implementation sources must fail closed: %#v", preview)
	}
	for _, item := range preview.ImplementationItems {
		if item.EvidencePresent ||
			item.CurrentMainlineConsumed ||
			item.PersistenceRecordContractConsumed ||
			item.PersistenceRecordContractReady ||
			item.ClosedRecordWriterImplementationModeled ||
			item.RecordWriterShapeModeled ||
			item.RecordWriterInputBoundaryModeled ||
			item.RecordWriterOutputBoundaryModeled ||
			item.DurableRecordContractConsumed ||
			item.WriterCallable ||
			item.ClosedRecordWriterEnabled ||
			item.StorageGatePassed ||
			item.StoragePersistenceAuthorized ||
			item.DurableRecordWriteEnabled ||
			item.StatusPersistenceAuthorized ||
			item.StatusPersistenceWriteEnabled ||
			item.StorageWriteEnabled ||
			item.RawResultExposed ||
			item.UserVisible ||
			item.ClosedRecordWriterImplementationStatus != "missing-redacted-status-closed-record-writer-implementation-evidence" {
			t.Fatalf("missing KDE-safe redacted status closed record writer implementation item must remain closed: %#v", item)
		}
	}
}
