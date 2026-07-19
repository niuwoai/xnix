package owner

import (
	"os"
	"path/filepath"
	"testing"
)

func TestProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionPreviewModelsConsumerProjection(t *testing.T) {
	preview, err := NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionPreview(projectRootForOwnerTest(t))
	if err != nil {
		t.Fatalf("NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionPreview returned error: %v", err)
	}

	if preview.Version != currentProjectVersion(t) ||
		preview.SchemaVersion != "xnix.runtime.production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_kde_safe_redacted_status_closed_record_writer_storage_root_authorization_receipt_dry_run_lookup_result_consumer_projection.v1" ||
		preview.RequestType != "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-preview" ||
		preview.PreviewType != "receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection" ||
		preview.PreviewDecision != "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-ready-writes-disabled" {
		t.Fatalf("unexpected KDE-safe redacted status dry-run lookup result consumer projection schema: %#v", preview)
	}
	if !preview.CurrentMainlineConsumed ||
		!preview.AuthorizationReceiptDryRunLookupResultBoundaryConsumed ||
		!preview.AuthorizationReceiptDryRunLookupResultBoundaryReady ||
		!preview.ResultConsumerProjectionRequired ||
		!preview.ResultConsumerProjectionModeled ||
		!preview.ResultConsumerProjectionReady ||
		!preview.CompatibilityCenterProjectionModeled ||
		!preview.RuntimeDiagnosticsProjectionModeled ||
		!preview.RedactedLookupResultProjectionModeled ||
		!preview.UserVisibleProjectionModeled ||
		!preview.FailureProjectionModeled ||
		!preview.KDESafeRedactedStatusOnly ||
		preview.ReceiptPresent ||
		preview.ReceiptAccepted ||
		preview.ReceiptConsumed ||
		preview.StorageRootPolicyGrantAuthorized ||
		preview.RecordWriterCallAuthorized ||
		preview.WriterCallable ||
		preview.StorageRootResolved ||
		preview.StorageWriteEnabled ||
		preview.StatusPersistenceWriteEnabled ||
		preview.ConsumerDryRunLookupResultBoundaryAuthorized ||
		preview.ConsumerEnablementAuthorized ||
		preview.KDEConsumerEnabled ||
		preview.RuntimeConsumerEnabled ||
		preview.LookupRouteAuthorized ||
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
		!preview.RuntimeOwned ||
		!preview.GoRuntimeBacked ||
		preview.KDEPolicyOwner ||
		preview.ProductionReadiness ||
		preview.ProductionOwnershipReady {
		t.Fatalf("unsafe KDE-safe redacted status dry-run lookup result consumer projection decision: %#v", preview)
	}
	if preview.ConsumerProjectionItemCount != 10 ||
		preview.ReadyConsumerProjectionItemCount != 10 ||
		preview.MissingConsumerProjectionItemCount != 0 ||
		preview.CompatibilityCenterItemCount != 5 ||
		preview.RuntimeDiagnosticsItemCount != 5 ||
		preview.RawExposedConsumerProjectionItemCount != 0 ||
		preview.SideEffectConsumerProjectionItemCount != 0 {
		t.Fatalf("unexpected KDE-safe redacted status dry-run lookup result consumer projection counts: %#v", preview)
	}
	if len(preview.ConsumerProjectionItems) != 10 {
		t.Fatalf("unexpected KDE-safe redacted status dry-run lookup result consumer projection items: %#v", preview.ConsumerProjectionItemIDs)
	}
	for _, item := range preview.ConsumerProjectionItems {
		if !item.EvidencePresent ||
			!item.CurrentMainlineConsumed ||
			!item.ResultBoundaryConsumed ||
			!item.ResultBoundaryReady ||
			!item.ResultConsumerProjectionModeled ||
			!item.RedactedLookupResultProjectionModeled ||
			!item.UserVisibleProjectionModeled ||
			!item.FailureProjectionModeled ||
			!item.KDESafeRedactedStatusOnly ||
			item.ReceiptPresent ||
			item.ReceiptAccepted ||
			item.ReceiptConsumed ||
			item.StorageRootPolicyGrantAuthorized ||
			item.RecordWriterCallAuthorized ||
			item.WriterCallable ||
			item.StorageWriteEnabled ||
			item.ConsumerEnablementAuthorized ||
			item.LookupRouteEnabled ||
			item.RedactedSummaryPersisted ||
			item.DryRunResultPersisted ||
			item.RawResultExposed ||
			!item.UserVisible ||
			!item.ReviewOnly ||
			!item.RuntimeOwned ||
			!item.GoRuntimeBacked ||
			item.KDEPolicyOwner ||
			!item.SideEffectsDisabled ||
			item.HostRootModified ||
			item.InternalDetailsExposed ||
			item.ResultConsumerProjectionStatus != "redacted-status-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-modeled-writes-disabled" {
			t.Fatalf("unsafe KDE-safe redacted status dry-run lookup result consumer projection item: %#v", item)
		}
	}
	if preview.Counts.Total != 8 ||
		preview.Counts.Passed != 8 ||
		preview.Counts.Pending != 0 ||
		preview.Counts.Blocked != 0 ||
		!sameStrings(preview.CheckIDs, []string{"current-mainline-consumed", "authorization-receipt-dry-run-lookup-result-boundary-consumed", "result-consumer-projection-modeled", "compatibility-center-and-runtime-result-consumer-projections-modeled", "ten-result-consumer-projection-items-ready-writes-disabled", "result-consumer-projection-and-writes-disabled", "consumer-routes-dispatch-and-support-disabled", "production-and-host-boundary-closed"}) {
		t.Fatalf("unexpected KDE-safe redacted status dry-run lookup result consumer projection checks: counts=%#v ids=%#v", preview.Counts, preview.CheckIDs)
	}
	if err := validateNoBackendTerms(preview, "KDE-safe redacted status dry-run lookup result consumer projection test"); err != nil {
		t.Fatalf("validateNoBackendTerms returned error: %v", err)
	}
}

func TestProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionPreviewFailsClosedWithoutSources(t *testing.T) {
	root := t.TempDir()
	if err := os.WriteFile(filepath.Join(root, "VERSION"), []byte(currentProjectVersion(t)+"\n"), 0o600); err != nil {
		t.Fatalf("WriteFile VERSION returned error: %v", err)
	}
	preview, err := NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionPreview(root)
	if err != nil {
		t.Fatalf("NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionPreview returned error: %v", err)
	}
	if preview.CurrentMainlineConsumed ||
		preview.AuthorizationReceiptDryRunLookupResultBoundaryConsumed ||
		preview.AuthorizationReceiptDryRunLookupResultBoundaryReady ||
		!preview.ResultConsumerProjectionRequired ||
		preview.ResultConsumerProjectionModeled ||
		preview.ResultConsumerProjectionReady ||
		preview.CompatibilityCenterProjectionModeled ||
		preview.RuntimeDiagnosticsProjectionModeled ||
		preview.RedactedLookupResultProjectionModeled ||
		preview.UserVisibleProjectionModeled ||
		preview.FailureProjectionModeled ||
		preview.KDESafeRedactedStatusOnly ||
		preview.ReceiptPresent ||
		preview.ReceiptAccepted ||
		preview.ReceiptConsumed ||
		preview.StorageRootPolicyGrantAuthorized ||
		preview.RecordWriterCallAuthorized ||
		preview.WriterCallable ||
		preview.StorageWriteEnabled ||
		preview.RawResultExposed ||
		preview.ConsumerProjectionItemCount != 10 ||
		preview.ReadyConsumerProjectionItemCount != 0 ||
		preview.MissingConsumerProjectionItemCount != 10 ||
		preview.Counts.Total != 8 ||
		preview.Counts.Passed != 3 ||
		preview.Counts.Blocked != 5 ||
		preview.PreviewDecision != "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-blocked" {
		t.Fatalf("missing KDE-safe redacted status dry-run lookup result consumer projection sources must fail closed: %#v", preview)
	}
	for _, item := range preview.ConsumerProjectionItems {
		if item.EvidencePresent ||
			item.CurrentMainlineConsumed ||
			item.ResultBoundaryConsumed ||
			item.ResultBoundaryReady ||
			item.ResultConsumerProjectionModeled ||
			item.RedactedLookupResultProjectionModeled ||
			item.UserVisibleProjectionModeled ||
			item.FailureProjectionModeled ||
			item.KDESafeRedactedStatusOnly ||
			item.ReceiptPresent ||
			item.ReceiptAccepted ||
			item.ReceiptConsumed ||
			item.StorageRootPolicyGrantAuthorized ||
			item.RecordWriterCallAuthorized ||
			item.WriterCallable ||
			item.StorageWriteEnabled ||
			item.RawResultExposed ||
			item.UserVisible ||
			item.ResultConsumerProjectionStatus != "missing-redacted-status-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection" {
			t.Fatalf("missing KDE-safe redacted status dry-run lookup result consumer projection item must remain closed: %#v", item)
		}
	}
}
