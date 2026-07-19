package owner

import (
	"os"
	"path/filepath"
	"testing"
)

func TestProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedPersistenceWriterPreviewModelsWriter(t *testing.T) {
	preview, err := NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedPersistenceWriterPreview(projectRootForOwnerTest(t))
	if err != nil {
		t.Fatalf("NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedPersistenceWriterPreview returned error: %v", err)
	}

	if preview.Version != currentProjectVersion(t) ||
		preview.SchemaVersion != "xnix.runtime.production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_kde_safe_redacted_status_closed_persistence_writer.v1" ||
		preview.RequestType != "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-persistence-writer-preview" ||
		preview.PreviewType != "receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-persistence-writer" ||
		preview.PreviewDecision != "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-persistence-writer-ready-storage-disabled" {
		t.Fatalf("unexpected KDE-safe redacted status closed persistence writer schema: %#v", preview)
	}
	if !preview.CurrentMainlineConsumed ||
		!preview.WriterPersistenceAuthorizationConsumed ||
		!preview.WriterPersistenceAuthorizationReady ||
		!preview.ClosedPersistenceWriterRequired ||
		!preview.ClosedPersistenceWriterModeled ||
		!preview.ClosedPersistenceWriterReady ||
		!preview.StorageWriterShapeModeled ||
		!preview.StorageInputBoundaryModeled ||
		!preview.StorageOutputBoundaryModeled ||
		!preview.KDESafeRedactedStatusOnly ||
		!preview.CompatibilityCenterStorageModeled ||
		!preview.RuntimeDiagnosticsStorageModeled ||
		preview.WriterCallable ||
		preview.ClosedPersistenceWriterEnabled ||
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
		t.Fatalf("unsafe KDE-safe redacted status closed persistence writer decision: %#v", preview)
	}
	if preview.WriterItemCount != 10 ||
		preview.RequiredWriterItemCount != 10 ||
		preview.ReadyWriterItemCount != 10 ||
		preview.MissingWriterItemCount != 0 ||
		preview.CallableWriterItemCount != 0 ||
		preview.EnabledPersistenceWriterItemCount != 0 ||
		preview.PersistedWriterItemCount != 0 ||
		preview.RawExposedWriterItemCount != 0 ||
		preview.SideEffectWriterItemCount != 0 ||
		preview.CompatibilityCenterItemCount != 5 ||
		preview.RuntimeDiagnosticsItemCount != 5 {
		t.Fatalf("unexpected KDE-safe redacted status closed persistence writer counts: %#v", preview)
	}
	if len(preview.WriterItems) != 10 ||
		!sameStrings(preview.WriterItemIDs, []string{
			"review-receipt-dry-run-result-lookup-consumer-enablement-kde-safe-compatibility-center-redacted-status-closed-persistence-writer",
			"review-receipt-dry-run-result-lookup-consumer-enablement-kde-safe-runtime-diagnostics-redacted-status-closed-persistence-writer",
			"renew-receipt-dry-run-result-lookup-consumer-enablement-kde-safe-compatibility-center-redacted-status-closed-persistence-writer",
			"renew-receipt-dry-run-result-lookup-consumer-enablement-kde-safe-runtime-diagnostics-redacted-status-closed-persistence-writer",
			"open-compatibility-center-dry-run-result-lookup-consumer-enablement-kde-safe-compatibility-center-redacted-status-closed-persistence-writer",
			"open-compatibility-center-dry-run-result-lookup-consumer-enablement-kde-safe-runtime-diagnostics-redacted-status-closed-persistence-writer",
			"dismiss-receipt-dry-run-result-lookup-consumer-enablement-kde-safe-compatibility-center-redacted-status-closed-persistence-writer",
			"dismiss-receipt-dry-run-result-lookup-consumer-enablement-kde-safe-runtime-diagnostics-redacted-status-closed-persistence-writer",
			"support-info-dry-run-result-lookup-consumer-enablement-kde-safe-compatibility-center-redacted-status-closed-persistence-writer",
			"support-info-dry-run-result-lookup-consumer-enablement-kde-safe-runtime-diagnostics-redacted-status-closed-persistence-writer",
		}) {
		t.Fatalf("unexpected KDE-safe redacted status closed persistence writer items: %#v", preview.WriterItemIDs)
	}
	for _, item := range preview.WriterItems {
		if !item.EvidencePresent ||
			!item.CurrentMainlineConsumed ||
			!item.WriterPersistenceAuthorizationReady ||
			!item.ClosedPersistenceWriterModeled ||
			!item.StorageWriterShapeModeled ||
			!item.StorageInputBoundaryModeled ||
			!item.StorageOutputBoundaryModeled ||
			!item.KDESafeRedactedStatusOnly ||
			item.WriterCallable ||
			item.ClosedPersistenceWriterEnabled ||
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
			item.ClosedPersistenceWriterStatus != "redacted-status-closed-persistence-writer-modeled-storage-disabled" {
			t.Fatalf("unsafe KDE-safe redacted status closed persistence writer item: %#v", item)
		}
	}
	if preview.Counts.Total != 8 ||
		preview.Counts.Passed != 8 ||
		preview.Counts.Pending != 0 ||
		preview.Counts.Blocked != 0 ||
		!sameStrings(preview.CheckIDs, []string{"current-mainline-consumed", "writer-persistence-authorization-consumed", "closed-persistence-writer-shape-modeled", "compatibility-center-and-runtime-storage-writers-modeled", "ten-writer-items-ready-storage-disabled", "closed-persistence-writes-disabled", "consumer-lookup-request-and-support-disabled", "production-and-host-boundary-closed"}) {
		t.Fatalf("unexpected KDE-safe redacted status closed persistence writer checks: counts=%#v ids=%#v", preview.Counts, preview.CheckIDs)
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
		t.Fatalf("unsafe KDE-safe redacted status closed persistence writer gates: %#v", preview)
	}
	if err := validateNoBackendTerms(preview, "KDE-safe redacted status closed persistence writer test"); err != nil {
		t.Fatalf("validateNoBackendTerms returned error: %v", err)
	}
}

func TestProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedPersistenceWriterPreviewFailsClosedWithoutSources(t *testing.T) {
	root := t.TempDir()
	if err := os.WriteFile(filepath.Join(root, "VERSION"), []byte(currentProjectVersion(t)+"\n"), 0o600); err != nil {
		t.Fatalf("WriteFile VERSION returned error: %v", err)
	}
	preview, err := NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedPersistenceWriterPreview(root)
	if err != nil {
		t.Fatalf("NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedPersistenceWriterPreview returned error: %v", err)
	}
	if preview.CurrentMainlineConsumed ||
		preview.WriterPersistenceAuthorizationConsumed ||
		preview.WriterPersistenceAuthorizationReady ||
		!preview.ClosedPersistenceWriterRequired ||
		preview.ClosedPersistenceWriterModeled ||
		preview.ClosedPersistenceWriterReady ||
		preview.StorageWriterShapeModeled ||
		preview.StorageInputBoundaryModeled ||
		preview.StorageOutputBoundaryModeled ||
		preview.KDESafeRedactedStatusOnly ||
		preview.CompatibilityCenterStorageModeled ||
		preview.RuntimeDiagnosticsStorageModeled ||
		preview.WriterCallable ||
		preview.ClosedPersistenceWriterEnabled ||
		preview.StatusPersistenceAuthorized ||
		preview.StatusPersistenceWriteEnabled ||
		preview.StorageWriteEnabled ||
		preview.RawResultExposed ||
		preview.WriterItemCount != 10 ||
		preview.RequiredWriterItemCount != 10 ||
		preview.ReadyWriterItemCount != 0 ||
		preview.MissingWriterItemCount != 10 ||
		preview.Counts.Total != 8 ||
		preview.Counts.Passed != 3 ||
		preview.Counts.Blocked != 5 ||
		preview.PreviewDecision != "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-persistence-writer-blocked" {
		t.Fatalf("missing KDE-safe redacted status closed persistence writer sources must fail closed: %#v", preview)
	}
	for _, item := range preview.WriterItems {
		if item.EvidencePresent ||
			item.CurrentMainlineConsumed ||
			item.WriterPersistenceAuthorizationReady ||
			item.ClosedPersistenceWriterModeled ||
			item.StorageWriterShapeModeled ||
			item.StorageInputBoundaryModeled ||
			item.StorageOutputBoundaryModeled ||
			item.WriterCallable ||
			item.ClosedPersistenceWriterEnabled ||
			item.StatusPersistenceAuthorized ||
			item.StatusPersistenceWriteEnabled ||
			item.StorageWriteEnabled ||
			item.RawResultExposed ||
			item.UserVisible ||
			item.ClosedPersistenceWriterStatus != "missing-redacted-status-closed-persistence-writer-evidence" {
			t.Fatalf("missing KDE-safe redacted status closed persistence writer item must remain closed: %#v", item)
		}
	}
}
