package owner

import (
	"os"
	"path/filepath"
	"testing"
)

func TestProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultBoundaryPreviewModelsDryRunLookupResultBoundary(t *testing.T) {
	preview, err := NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultBoundaryPreview(projectRootForOwnerTest(t))
	if err != nil {
		t.Fatalf("NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultBoundaryPreview returned error: %v", err)
	}

	if preview.Version != currentProjectVersion(t) ||
		preview.SchemaVersion != "xnix.runtime.production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_kde_safe_redacted_status_closed_record_writer_storage_root_authorization_receipt_dry_run_lookup_result_boundary.v1" ||
		preview.RequestType != "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-boundary-preview" ||
		preview.PreviewType != "receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-boundary" ||
		preview.PreviewDecision != "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-boundary-ready-writes-disabled" {
		t.Fatalf("unexpected KDE-safe redacted status storage-root authorization receipt dry-run lookup result boundary schema: %#v", preview)
	}
	if !preview.CurrentMainlineConsumed ||
		!preview.AuthorizationReceiptDryRunLookupEvidenceConsumed ||
		!preview.AuthorizationReceiptDryRunLookupEvidenceReady ||
		!preview.AuthorizationReceiptDryRunLookupResultBoundaryRequired ||
		!preview.AuthorizationReceiptDryRunLookupResultBoundaryModeled ||
		!preview.AuthorizationReceiptDryRunLookupResultBoundaryReady ||
		!preview.RedactedLookupResultIdentityModeled ||
		!preview.RedactedLookupResultReadinessModeled ||
		!preview.RedactedLookupResultFailureBoundaryModeled ||
		!preview.AcceptedReceiptBoundaryModeled ||
		!preview.ReceiptScopeBoundaryModeled ||
		!preview.StorageRootPolicyGrantBoundaryModeled ||
		!preview.RecordWriterCallGrantBoundaryModeled ||
		!preview.KDESafeRedactedStatusOnly ||
		!preview.CompatibilityCenterDryRunLookupResultBoundaryModeled ||
		!preview.RuntimeDiagnosticsDryRunLookupResultBoundaryModeled ||
		preview.ReceiptPresent ||
		preview.ReceiptAccepted ||
		preview.ReceiptConsumed ||
		preview.StorageRootPolicyGrantAuthorized ||
		preview.RecordWriterCallAuthorized ||
		preview.WriterCallable ||
		preview.ClosedRecordWriterEnabled ||
		preview.StorageRootResolved ||
		preview.StorageRootCreated ||
		preview.StorageRootMounted ||
		preview.StorageRootOwnershipGranted ||
		preview.StorageRootNamespaceGranted ||
		preview.StorageRetentionEnforced ||
		preview.StorageRedactionEnforced ||
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
		preview.ConsumerDryRunLookupResultBoundaryAuthorized ||
		preview.ConsumerEnablementAuthorized ||
		preview.KDEConsumerEnabled ||
		preview.RuntimeConsumerEnabled ||
		preview.LookupRouteEnabled ||
		preview.OpaqueLookupEnabled ||
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
		t.Fatalf("unsafe KDE-safe redacted status storage-root authorization receipt dry-run lookup result boundary decision: %#v", preview)
	}
	if preview.DryRunLookupResultBoundaryItemCount != 10 ||
		preview.RequiredDryRunLookupResultBoundaryItemCount != 10 ||
		preview.ReadyDryRunLookupResultBoundaryItemCount != 10 ||
		preview.MissingDryRunLookupResultBoundaryItemCount != 0 ||
		preview.AcceptedReceiptItemCount != 0 ||
		preview.ConsumedReceiptItemCount != 0 ||
		preview.AuthorizedStorageRootGrantItemCount != 0 ||
		preview.AuthorizedRecordWriterCallItemCount != 0 ||
		preview.CallableWriterItemCount != 0 ||
		preview.EnabledRecordWriterItemCount != 0 ||
		preview.DurableWrittenRecordItemCount != 0 ||
		preview.RawExposedDryRunLookupResultBoundaryItemCount != 0 ||
		preview.SideEffectDryRunLookupResultBoundaryItemCount != 0 ||
		preview.CompatibilityCenterItemCount != 5 ||
		preview.RuntimeDiagnosticsItemCount != 5 {
		t.Fatalf("unexpected KDE-safe redacted status storage-root authorization receipt dry-run lookup result boundary counts: %#v", preview)
	}
	if len(preview.DryRunLookupResultBoundaryItems) != 10 {
		t.Fatalf("unexpected KDE-safe redacted status storage-root authorization receipt dry-run lookup result boundary items: %#v", preview.DryRunLookupResultBoundaryItemIDs)
	}
	for _, item := range preview.DryRunLookupResultBoundaryItems {
		if !item.EvidencePresent ||
			!item.CurrentMainlineConsumed ||
			!item.AuthorizationReceiptDryRunLookupEvidenceConsumed ||
			!item.AuthorizationReceiptDryRunLookupEvidenceReady ||
			!item.AuthorizationReceiptDryRunLookupResultBoundaryModeled ||
			!item.RedactedLookupResultIdentityModeled ||
			!item.RedactedLookupResultReadinessModeled ||
			!item.RedactedLookupResultFailureBoundaryModeled ||
			!item.AcceptedReceiptBoundaryModeled ||
			!item.ReceiptScopeBoundaryModeled ||
			!item.StorageRootPolicyGrantBoundaryModeled ||
			!item.RecordWriterCallGrantBoundaryModeled ||
			!item.KDESafeRedactedStatusOnly ||
			item.ReceiptPresent ||
			item.ReceiptAccepted ||
			item.ReceiptConsumed ||
			item.StorageRootPolicyGrantAuthorized ||
			item.RecordWriterCallAuthorized ||
			item.WriterCallable ||
			item.ClosedRecordWriterEnabled ||
			item.StorageRootResolved ||
			item.StorageRootOwnershipGranted ||
			item.StorageRootNamespaceGranted ||
			item.StorageRetentionEnforced ||
			item.StorageRedactionEnforced ||
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
			item.AuthorizationReceiptDryRunLookupResultBoundaryStatus != "redacted-status-storage-root-authorization-receipt-dry-run-lookup-result-boundary-modeled-writes-disabled" {
			t.Fatalf("unsafe KDE-safe redacted status storage-root authorization receipt dry-run lookup result boundary item: %#v", item)
		}
	}
	if preview.Counts.Total != 8 ||
		preview.Counts.Passed != 8 ||
		preview.Counts.Pending != 0 ||
		preview.Counts.Blocked != 0 ||
		!sameStrings(preview.CheckIDs, []string{"current-mainline-consumed", "authorization-receipt-dry-run-lookup-evidence-consumed", "authorization-receipt-dry-run-lookup-result-boundary-modeled", "compatibility-center-and-runtime-dry-run-lookup-result-boundary-modeled", "ten-dry-run-lookup-result-boundary-items-ready-writes-disabled", "authorization-dry-run-lookup-result-boundary-and-writes-disabled", "consumer-lookup-request-and-support-disabled", "production-and-host-boundary-closed"}) {
		t.Fatalf("unexpected KDE-safe redacted status storage-root authorization receipt dry-run lookup result boundary checks: counts=%#v ids=%#v", preview.Counts, preview.CheckIDs)
	}
	if err := validateNoBackendTerms(preview, "KDE-safe redacted status storage-root authorization receipt dry-run lookup result boundary test"); err != nil {
		t.Fatalf("validateNoBackendTerms returned error: %v", err)
	}
}

func TestProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultBoundaryPreviewFailsClosedWithoutSources(t *testing.T) {
	root := t.TempDir()
	if err := os.WriteFile(filepath.Join(root, "VERSION"), []byte(currentProjectVersion(t)+"\n"), 0o600); err != nil {
		t.Fatalf("WriteFile VERSION returned error: %v", err)
	}
	preview, err := NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultBoundaryPreview(root)
	if err != nil {
		t.Fatalf("NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultBoundaryPreview returned error: %v", err)
	}
	if preview.CurrentMainlineConsumed ||
		preview.AuthorizationReceiptDryRunLookupEvidenceConsumed ||
		preview.AuthorizationReceiptDryRunLookupEvidenceReady ||
		!preview.AuthorizationReceiptDryRunLookupResultBoundaryRequired ||
		preview.AuthorizationReceiptDryRunLookupResultBoundaryModeled ||
		preview.AuthorizationReceiptDryRunLookupResultBoundaryReady ||
		preview.RedactedLookupResultIdentityModeled ||
		preview.RedactedLookupResultReadinessModeled ||
		preview.RedactedLookupResultFailureBoundaryModeled ||
		preview.AcceptedReceiptBoundaryModeled ||
		preview.ReceiptScopeBoundaryModeled ||
		preview.StorageRootPolicyGrantBoundaryModeled ||
		preview.RecordWriterCallGrantBoundaryModeled ||
		preview.KDESafeRedactedStatusOnly ||
		preview.ReceiptPresent ||
		preview.ReceiptAccepted ||
		preview.ReceiptConsumed ||
		preview.StorageRootPolicyGrantAuthorized ||
		preview.RecordWriterCallAuthorized ||
		preview.WriterCallable ||
		preview.StorageRootResolved ||
		preview.StorageWriteEnabled ||
		preview.RawResultExposed ||
		preview.DryRunLookupResultBoundaryItemCount != 10 ||
		preview.RequiredDryRunLookupResultBoundaryItemCount != 10 ||
		preview.ReadyDryRunLookupResultBoundaryItemCount != 0 ||
		preview.MissingDryRunLookupResultBoundaryItemCount != 10 ||
		preview.Counts.Total != 8 ||
		preview.Counts.Passed != 3 ||
		preview.Counts.Blocked != 5 ||
		preview.PreviewDecision != "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-boundary-blocked" {
		t.Fatalf("missing KDE-safe redacted status storage-root authorization receipt dry-run lookup result boundary sources must fail closed: %#v", preview)
	}
	for _, item := range preview.DryRunLookupResultBoundaryItems {
		if item.EvidencePresent ||
			item.CurrentMainlineConsumed ||
			item.AuthorizationReceiptDryRunLookupEvidenceConsumed ||
			item.AuthorizationReceiptDryRunLookupEvidenceReady ||
			item.AuthorizationReceiptDryRunLookupResultBoundaryModeled ||
			item.RedactedLookupResultIdentityModeled ||
			item.RedactedLookupResultReadinessModeled ||
			item.RedactedLookupResultFailureBoundaryModeled ||
			item.AcceptedReceiptBoundaryModeled ||
			item.ReceiptScopeBoundaryModeled ||
			item.StorageRootPolicyGrantBoundaryModeled ||
			item.RecordWriterCallGrantBoundaryModeled ||
			item.ReceiptPresent ||
			item.ReceiptAccepted ||
			item.ReceiptConsumed ||
			item.StorageRootPolicyGrantAuthorized ||
			item.RecordWriterCallAuthorized ||
			item.WriterCallable ||
			item.StorageRootResolved ||
			item.StorageWriteEnabled ||
			item.RawResultExposed ||
			item.UserVisible ||
			item.AuthorizationReceiptDryRunLookupResultBoundaryStatus != "missing-redacted-status-storage-root-authorization-receipt-dry-run-lookup-result-boundary-missing" {
			t.Fatalf("missing KDE-safe redacted status storage-root authorization receipt dry-run lookup result boundary item must remain closed: %#v", item)
		}
	}
}
