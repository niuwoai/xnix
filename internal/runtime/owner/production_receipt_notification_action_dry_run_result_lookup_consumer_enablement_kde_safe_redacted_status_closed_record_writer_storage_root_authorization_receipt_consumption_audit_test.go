package owner

import (
	"os"
	"path/filepath"
	"testing"
)

func TestProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptConsumptionAuditPreviewModelsConsumption(t *testing.T) {
	preview, err := NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptConsumptionAuditPreview(projectRootForOwnerTest(t))
	if err != nil {
		t.Fatalf("NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptConsumptionAuditPreview returned error: %v", err)
	}

	if preview.Version != currentProjectVersion(t) ||
		preview.SchemaVersion != "xnix.runtime.production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_kde_safe_redacted_status_closed_record_writer_storage_root_authorization_receipt_consumption_audit.v1" ||
		preview.RequestType != "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-consumption-audit-preview" ||
		preview.PreviewType != "receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-consumption-audit" ||
		preview.PreviewDecision != "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-consumption-audit-ready-writes-disabled" {
		t.Fatalf("unexpected KDE-safe redacted status storage-root authorization receipt consumption schema: %#v", preview)
	}
	if !preview.CurrentMainlineConsumed ||
		!preview.StorageRootPolicyConsumed ||
		!preview.StorageRootPolicyReady ||
		!preview.AuthorizationReceiptConsumptionRequired ||
		!preview.AuthorizationReceiptConsumptionModeled ||
		!preview.AuthorizationReceiptConsumptionReady ||
		!preview.AcceptedReceiptBoundaryModeled ||
		!preview.ReceiptScopeBoundaryModeled ||
		!preview.StorageRootPolicyGrantBoundaryModeled ||
		!preview.RecordWriterCallGrantBoundaryModeled ||
		!preview.KDESafeRedactedStatusOnly ||
		!preview.CompatibilityCenterConsumptionModeled ||
		!preview.RuntimeDiagnosticsConsumptionModeled ||
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
		preview.ConsumerConsumptionAuthorized ||
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
		t.Fatalf("unsafe KDE-safe redacted status storage-root authorization receipt consumption decision: %#v", preview)
	}
	if preview.ConsumptionItemCount != 10 ||
		preview.RequiredConsumptionItemCount != 10 ||
		preview.ReadyConsumptionItemCount != 10 ||
		preview.MissingConsumptionItemCount != 0 ||
		preview.AcceptedReceiptItemCount != 0 ||
		preview.ConsumedReceiptItemCount != 0 ||
		preview.AuthorizedStorageRootGrantItemCount != 0 ||
		preview.AuthorizedRecordWriterCallItemCount != 0 ||
		preview.CallableWriterItemCount != 0 ||
		preview.EnabledRecordWriterItemCount != 0 ||
		preview.DurableWrittenRecordItemCount != 0 ||
		preview.RawExposedConsumptionItemCount != 0 ||
		preview.SideEffectConsumptionItemCount != 0 ||
		preview.CompatibilityCenterItemCount != 5 ||
		preview.RuntimeDiagnosticsItemCount != 5 {
		t.Fatalf("unexpected KDE-safe redacted status storage-root authorization receipt consumption counts: %#v", preview)
	}
	if len(preview.ConsumptionItems) != 10 {
		t.Fatalf("unexpected KDE-safe redacted status storage-root authorization receipt consumption items: %#v", preview.ConsumptionItemIDs)
	}
	for _, item := range preview.ConsumptionItems {
		if !item.EvidencePresent ||
			!item.CurrentMainlineConsumed ||
			!item.StorageRootPolicyConsumed ||
			!item.StorageRootPolicyReady ||
			!item.AuthorizationReceiptConsumptionModeled ||
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
			item.AuthorizationReceiptConsumptionStatus != "redacted-status-storage-root-authorization-receipt-consumption-modeled-writes-disabled" {
			t.Fatalf("unsafe KDE-safe redacted status storage-root authorization receipt consumption item: %#v", item)
		}
	}
	if preview.Counts.Total != 8 ||
		preview.Counts.Passed != 8 ||
		preview.Counts.Pending != 0 ||
		preview.Counts.Blocked != 0 ||
		!sameStrings(preview.CheckIDs, []string{"current-mainline-consumed", "storage-root-policy-consumed", "authorization-receipt-consumption-modeled", "compatibility-center-and-runtime-consumption-modeled", "ten-consumption-items-ready-writes-disabled", "authorization-consumption-and-writes-disabled", "consumer-lookup-request-and-support-disabled", "production-and-host-boundary-closed"}) {
		t.Fatalf("unexpected KDE-safe redacted status storage-root authorization receipt consumption checks: counts=%#v ids=%#v", preview.Counts, preview.CheckIDs)
	}
	if err := validateNoBackendTerms(preview, "KDE-safe redacted status storage-root authorization receipt consumption test"); err != nil {
		t.Fatalf("validateNoBackendTerms returned error: %v", err)
	}
}

func TestProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptConsumptionAuditPreviewFailsClosedWithoutSources(t *testing.T) {
	root := t.TempDir()
	if err := os.WriteFile(filepath.Join(root, "VERSION"), []byte(currentProjectVersion(t)+"\n"), 0o600); err != nil {
		t.Fatalf("WriteFile VERSION returned error: %v", err)
	}
	preview, err := NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptConsumptionAuditPreview(root)
	if err != nil {
		t.Fatalf("NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptConsumptionAuditPreview returned error: %v", err)
	}
	if preview.CurrentMainlineConsumed ||
		preview.StorageRootPolicyConsumed ||
		preview.StorageRootPolicyReady ||
		!preview.AuthorizationReceiptConsumptionRequired ||
		preview.AuthorizationReceiptConsumptionModeled ||
		preview.AuthorizationReceiptConsumptionReady ||
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
		preview.ConsumptionItemCount != 10 ||
		preview.RequiredConsumptionItemCount != 10 ||
		preview.ReadyConsumptionItemCount != 0 ||
		preview.MissingConsumptionItemCount != 10 ||
		preview.Counts.Total != 8 ||
		preview.Counts.Passed != 3 ||
		preview.Counts.Blocked != 5 ||
		preview.PreviewDecision != "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-consumption-audit-blocked" {
		t.Fatalf("missing KDE-safe redacted status storage-root authorization receipt consumption sources must fail closed: %#v", preview)
	}
	for _, item := range preview.ConsumptionItems {
		if item.EvidencePresent ||
			item.CurrentMainlineConsumed ||
			item.StorageRootPolicyConsumed ||
			item.StorageRootPolicyReady ||
			item.AuthorizationReceiptConsumptionModeled ||
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
			item.AuthorizationReceiptConsumptionStatus != "missing-redacted-status-storage-root-authorization-receipt-consumption-evidence" {
			t.Fatalf("missing KDE-safe redacted status storage-root authorization receipt consumption item must remain closed: %#v", item)
		}
	}
}
