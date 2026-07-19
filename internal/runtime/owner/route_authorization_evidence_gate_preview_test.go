package owner

import (
	"os"
	"path/filepath"
	"testing"
)

func TestProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteAuthorizationEvidenceGatePreviewAuditsProjectionEvidence(t *testing.T) {
	preview, err := NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteAuthorizationEvidenceGatePreview(projectRootForOwnerTest(t))
	if err != nil {
		t.Fatalf("NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteAuthorizationEvidenceGatePreview returned error: %v", err)
	}

	if preview.Version != currentProjectVersion(t) ||
		preview.SchemaVersion != "xnix.runtime.production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_kde_safe_redacted_status_closed_record_writer_storage_root_authorization_receipt_dry_run_lookup_result_consumer_projection_route_authorization_evidence_gate.v1" ||
		preview.RequestType != "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-authorization-evidence-gate-preview" ||
		preview.GateType != "receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-authorization-evidence-gate" ||
		preview.AuditDecision != "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-authorization-evidence-gate-ready-routes-disabled" {
		t.Fatalf("unexpected KDE-safe redacted status dry-run lookup result consumer projection route authorization evidence gate schema: %#v", preview)
	}
	if !preview.CurrentMainlineConsumed ||
		!preview.ResultConsumerProjectionEvidenceAuditConsumed ||
		!preview.ResultConsumerProjectionEvidenceAuditReady ||
		!preview.RouteAuthorizationEvidenceGateRequired ||
		!preview.ResultConsumerProjectionRouteAuthorizationEvidenceGated ||
		!preview.RouteAuthorizationEvidenceGateReady ||
		!preview.CompatibilityCenterRouteAuthorizationEvidenceGated ||
		!preview.RuntimeDiagnosticsRouteAuthorizationEvidenceGated ||
		!preview.RedactedProjectionRouteAuthorizationEvidenceGated ||
		!preview.UserVisibleRouteAuthorizationEvidenceGated ||
		!preview.FailureProjectionRouteAuthorizationEvidenceGated ||
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
		t.Fatalf("unsafe KDE-safe redacted status dry-run lookup result consumer projection route authorization evidence gate decision: %#v", preview)
	}
	if preview.EvidenceItemCount != 10 ||
		preview.ReadyEvidenceItemCount != 10 ||
		preview.MissingEvidenceItemCount != 0 ||
		preview.CompatibilityCenterEvidenceItemCount != 5 ||
		preview.RuntimeDiagnosticsEvidenceItemCount != 5 ||
		preview.RawExposedEvidenceItemCount != 0 ||
		preview.SideEffectEvidenceItemCount != 0 {
		t.Fatalf("unexpected KDE-safe redacted status dry-run lookup result consumer projection route authorization evidence gate counts: %#v", preview)
	}
	if len(preview.EvidenceItems) != 10 {
		t.Fatalf("unexpected KDE-safe redacted status dry-run lookup result consumer projection route authorization evidence gate items: %#v", preview.EvidenceItemIDs)
	}
	for _, item := range preview.EvidenceItems {
		if !item.EvidencePresent ||
			!item.CurrentMainlineConsumed ||
			!item.ProjectionEvidenceAuditConsumed ||
			!item.ProjectionEvidenceAuditReady ||
			!item.ResultConsumerProjectionRouteAuthorizationEvidenceGated ||
			!item.RedactedProjectionRouteAuthorizationEvidenceGated ||
			!item.UserVisibleRouteAuthorizationEvidenceGated ||
			!item.FailureProjectionRouteAuthorizationEvidenceGated ||
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
			item.RouteAuthorizationGateStatus != "redacted-status-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-authorization-evidence-gated-routes-disabled" {
			t.Fatalf("unsafe KDE-safe redacted status dry-run lookup result consumer projection route authorization evidence gate item: %#v", item)
		}
	}
	if preview.Counts.Total != 8 ||
		preview.Counts.Passed != 8 ||
		preview.Counts.Pending != 0 ||
		preview.Counts.Blocked != 0 ||
		!sameStrings(preview.CheckIDs, []string{"current-mainline-consumed", "result-consumer-projection-evidence-audit-consumed", "result-consumer-projection-route-authorization-evidence-gated", "compatibility-center-and-runtime-route-authorization-evidence-gated", "ten-route-authorization-evidence-items-ready-routes-disabled", "route-authorization-evidence-and-writes-disabled", "consumer-routes-dispatch-and-support-disabled", "production-and-host-boundary-closed"}) {
		t.Fatalf("unexpected KDE-safe redacted status dry-run lookup result consumer projection route authorization evidence gate checks: counts=%#v ids=%#v", preview.Counts, preview.CheckIDs)
	}
	if err := validateNoBackendTerms(preview, "KDE-safe redacted status dry-run lookup result consumer projection route authorization evidence gate test"); err != nil {
		t.Fatalf("validateNoBackendTerms returned error: %v", err)
	}
}

func TestProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteAuthorizationEvidenceGatePreviewFailsClosedWithoutSources(t *testing.T) {
	root := t.TempDir()
	if err := os.WriteFile(filepath.Join(root, "VERSION"), []byte(currentProjectVersion(t)+"\n"), 0o600); err != nil {
		t.Fatalf("WriteFile VERSION returned error: %v", err)
	}
	preview, err := NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteAuthorizationEvidenceGatePreview(root)
	if err != nil {
		t.Fatalf("NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteAuthorizationEvidenceGatePreview returned error: %v", err)
	}
	if preview.CurrentMainlineConsumed ||
		preview.ResultConsumerProjectionEvidenceAuditConsumed ||
		preview.ResultConsumerProjectionEvidenceAuditReady ||
		!preview.RouteAuthorizationEvidenceGateRequired ||
		preview.ResultConsumerProjectionRouteAuthorizationEvidenceGated ||
		preview.RouteAuthorizationEvidenceGateReady ||
		preview.CompatibilityCenterRouteAuthorizationEvidenceGated ||
		preview.RuntimeDiagnosticsRouteAuthorizationEvidenceGated ||
		preview.RedactedProjectionRouteAuthorizationEvidenceGated ||
		preview.UserVisibleRouteAuthorizationEvidenceGated ||
		preview.FailureProjectionRouteAuthorizationEvidenceGated ||
		preview.KDESafeRedactedStatusOnly ||
		preview.ReceiptPresent ||
		preview.ReceiptAccepted ||
		preview.ReceiptConsumed ||
		preview.StorageRootPolicyGrantAuthorized ||
		preview.RecordWriterCallAuthorized ||
		preview.WriterCallable ||
		preview.StorageWriteEnabled ||
		preview.RawResultExposed ||
		preview.EvidenceItemCount != 10 ||
		preview.ReadyEvidenceItemCount != 0 ||
		preview.MissingEvidenceItemCount != 10 ||
		preview.Counts.Total != 8 ||
		preview.Counts.Passed != 3 ||
		preview.Counts.Blocked != 5 ||
		preview.AuditDecision != "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-authorization-evidence-gate-blocked" {
		t.Fatalf("missing KDE-safe redacted status dry-run lookup result consumer projection route authorization evidence gate sources must fail closed: %#v", preview)
	}
	for _, item := range preview.EvidenceItems {
		if item.EvidencePresent ||
			item.CurrentMainlineConsumed ||
			item.ProjectionEvidenceAuditConsumed ||
			item.ProjectionEvidenceAuditReady ||
			item.ResultConsumerProjectionRouteAuthorizationEvidenceGated ||
			item.RedactedProjectionRouteAuthorizationEvidenceGated ||
			item.UserVisibleRouteAuthorizationEvidenceGated ||
			item.FailureProjectionRouteAuthorizationEvidenceGated ||
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
			item.RouteAuthorizationGateStatus != "missing-redacted-status-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-authorization-evidence-gate" {
			t.Fatalf("missing KDE-safe redacted status dry-run lookup result consumer projection route authorization evidence gate item must remain closed: %#v", item)
		}
	}
}
