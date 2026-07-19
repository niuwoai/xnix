package owner

import (
	"os"
	"path/filepath"
	"testing"
)

func TestProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupEvidencePreviewModelsDryRunLookupEvidence(t *testing.T) {
	preview, err := NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupEvidencePreview(projectRootForOwnerTest(t))
	if err != nil {
		t.Fatalf("NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupEvidencePreview returned error: %v", err)
	}

	if preview.Version != currentProjectVersion(t) ||
		preview.SchemaVersion != "xnix.runtime.production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_kde_safe_redacted_status_closed_record_writer_storage_root_authorization_receipt_dry_run_lookup_evidence.v1" ||
		preview.RequestType != "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-evidence-preview" ||
		preview.PreviewType != "receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-evidence" ||
		preview.PreviewDecision != "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-evidence-ready-writes-disabled" {
		t.Fatalf("unexpected KDE-safe redacted status storage-root authorization receipt dry-run lookup evidence schema: %#v", preview)
	}
	if !preview.CurrentMainlineConsumed ||
		!preview.AuthorizationReceiptConsumptionAuditConsumed ||
		!preview.AuthorizationReceiptConsumptionAuditReady ||
		!preview.AuthorizationReceiptDryRunLookupEvidenceRequired ||
		!preview.AuthorizationReceiptDryRunLookupEvidenceModeled ||
		!preview.AuthorizationReceiptDryRunLookupEvidenceReady ||
		!preview.AcceptedReceiptBoundaryModeled ||
		!preview.ReceiptScopeBoundaryModeled ||
		!preview.StorageRootPolicyGrantBoundaryModeled ||
		!preview.RecordWriterCallGrantBoundaryModeled ||
		!preview.KDESafeRedactedStatusOnly ||
		!preview.CompatibilityCenterDryRunLookupEvidenceModeled ||
		!preview.RuntimeDiagnosticsDryRunLookupEvidenceModeled ||
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
		preview.ConsumerDryRunLookupEvidenceAuthorized ||
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
		t.Fatalf("unsafe KDE-safe redacted status storage-root authorization receipt dry-run lookup evidence decision: %#v", preview)
	}
	if preview.DryRunLookupEvidenceItemCount != 10 ||
		preview.RequiredDryRunLookupEvidenceItemCount != 10 ||
		preview.ReadyDryRunLookupEvidenceItemCount != 10 ||
		preview.MissingDryRunLookupEvidenceItemCount != 0 ||
		preview.AcceptedReceiptItemCount != 0 ||
		preview.ConsumedReceiptItemCount != 0 ||
		preview.AuthorizedStorageRootGrantItemCount != 0 ||
		preview.AuthorizedRecordWriterCallItemCount != 0 ||
		preview.CallableWriterItemCount != 0 ||
		preview.EnabledRecordWriterItemCount != 0 ||
		preview.DurableWrittenRecordItemCount != 0 ||
		preview.RawExposedDryRunLookupEvidenceItemCount != 0 ||
		preview.SideEffectDryRunLookupEvidenceItemCount != 0 ||
		preview.CompatibilityCenterItemCount != 5 ||
		preview.RuntimeDiagnosticsItemCount != 5 {
		t.Fatalf("unexpected KDE-safe redacted status storage-root authorization receipt dry-run lookup evidence counts: %#v", preview)
	}
	if len(preview.DryRunLookupEvidenceItems) != 10 {
		t.Fatalf("unexpected KDE-safe redacted status storage-root authorization receipt dry-run lookup evidence items: %#v", preview.DryRunLookupEvidenceItemIDs)
	}
	for _, item := range preview.DryRunLookupEvidenceItems {
		if !item.EvidencePresent ||
			!item.CurrentMainlineConsumed ||
			!item.AuthorizationReceiptConsumptionAuditConsumed ||
			!item.AuthorizationReceiptConsumptionAuditReady ||
			!item.AuthorizationReceiptDryRunLookupEvidenceModeled ||
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
			item.AuthorizationReceiptDryRunLookupEvidenceStatus != "redacted-status-storage-root-authorization-receipt-dry-run-lookup-evidence-modeled-writes-disabled" {
			t.Fatalf("unsafe KDE-safe redacted status storage-root authorization receipt dry-run lookup evidence item: %#v", item)
		}
	}
	if preview.Counts.Total != 8 ||
		preview.Counts.Passed != 8 ||
		preview.Counts.Pending != 0 ||
		preview.Counts.Blocked != 0 ||
		!sameStrings(preview.CheckIDs, []string{"current-mainline-consumed", "authorization-receipt-consumption-audit-consumed", "authorization-receipt-dry-run-lookup-evidence-modeled", "compatibility-center-and-runtime-dry-run-lookup-evidence-modeled", "ten-dry-run-lookup-evidence-items-ready-writes-disabled", "authorization-dry-run-lookup-evidence-and-writes-disabled", "consumer-lookup-request-and-support-disabled", "production-and-host-boundary-closed"}) {
		t.Fatalf("unexpected KDE-safe redacted status storage-root authorization receipt dry-run lookup evidence checks: counts=%#v ids=%#v", preview.Counts, preview.CheckIDs)
	}
	if err := validateNoBackendTerms(preview, "KDE-safe redacted status storage-root authorization receipt dry-run lookup evidence test"); err != nil {
		t.Fatalf("validateNoBackendTerms returned error: %v", err)
	}
}

func TestProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupEvidencePreviewFailsClosedWithoutSources(t *testing.T) {
	root := t.TempDir()
	if err := os.WriteFile(filepath.Join(root, "VERSION"), []byte(currentProjectVersion(t)+"\n"), 0o600); err != nil {
		t.Fatalf("WriteFile VERSION returned error: %v", err)
	}
	preview, err := NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupEvidencePreview(root)
	if err != nil {
		t.Fatalf("NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupEvidencePreview returned error: %v", err)
	}
	if preview.CurrentMainlineConsumed ||
		preview.AuthorizationReceiptConsumptionAuditConsumed ||
		preview.AuthorizationReceiptConsumptionAuditReady ||
		!preview.AuthorizationReceiptDryRunLookupEvidenceRequired ||
		preview.AuthorizationReceiptDryRunLookupEvidenceModeled ||
		preview.AuthorizationReceiptDryRunLookupEvidenceReady ||
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
		preview.DryRunLookupEvidenceItemCount != 10 ||
		preview.RequiredDryRunLookupEvidenceItemCount != 10 ||
		preview.ReadyDryRunLookupEvidenceItemCount != 0 ||
		preview.MissingDryRunLookupEvidenceItemCount != 10 ||
		preview.Counts.Total != 8 ||
		preview.Counts.Passed != 3 ||
		preview.Counts.Blocked != 5 ||
		preview.PreviewDecision != "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-evidence-blocked" {
		t.Fatalf("missing KDE-safe redacted status storage-root authorization receipt dry-run lookup evidence sources must fail closed: %#v", preview)
	}
	for _, item := range preview.DryRunLookupEvidenceItems {
		if item.EvidencePresent ||
			item.CurrentMainlineConsumed ||
			item.AuthorizationReceiptConsumptionAuditConsumed ||
			item.AuthorizationReceiptConsumptionAuditReady ||
			item.AuthorizationReceiptDryRunLookupEvidenceModeled ||
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
			item.AuthorizationReceiptDryRunLookupEvidenceStatus != "missing-redacted-status-storage-root-authorization-receipt-dry-run-lookup-evidence-missing" {
			t.Fatalf("missing KDE-safe redacted status storage-root authorization receipt dry-run lookup evidence item must remain closed: %#v", item)
		}
	}
}
