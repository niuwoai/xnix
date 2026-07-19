package owner

import (
	"os"
	"path/filepath"
	"testing"
)

func TestProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementEvidenceGatePreviewAuditsRouteAuthorizationGate(t *testing.T) {
	preview, err := NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementEvidenceGatePreview(projectRootForOwnerTest(t))
	if err != nil {
		t.Fatalf("NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementEvidenceGatePreview returned error: %v", err)
	}

	if preview.Version != currentProjectVersion(t) ||
		preview.SchemaVersion != "xnix.runtime.production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_kde_safe_redacted_status_closed_record_writer_storage_root_authorization_receipt_dry_run_lookup_result_consumer_projection_route_enablement_evidence_gate.v1" ||
		preview.RequestType != "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-evidence-gate-preview" ||
		preview.GateType != "receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-evidence-gate" ||
		preview.AuditDecision != "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-evidence-gate-ready-routes-disabled" {
		t.Fatalf("unexpected KDE-safe redacted status dry-run lookup result consumer projection route enablement evidence gate schema: %#v", preview)
	}
	if !preview.CurrentMainlineConsumed ||
		!preview.RouteAuthorizationEvidenceGateConsumed ||
		!preview.RouteAuthorizationEvidenceGateReady ||
		!preview.RouteEnablementEvidenceGateRequired ||
		!preview.ResultConsumerProjectionRouteEnablementGated ||
		!preview.RouteEnablementEvidenceGateReady ||
		!preview.CompatibilityCenterRouteEnablementGated ||
		!preview.RuntimeDiagnosticsRouteEnablementGated ||
		!preview.RedactedProjectionRouteEnablementGated ||
		!preview.UserVisibleRouteEnablementGated ||
		!preview.FailureProjectionRouteEnablementGated ||
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
		preview.LookupRouteEnablementAuthorized ||
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
		t.Fatalf("unsafe KDE-safe redacted status dry-run lookup result consumer projection route enablement evidence gate decision: %#v", preview)
	}
	if preview.EvidenceItemCount != 10 ||
		preview.ReadyEvidenceItemCount != 10 ||
		preview.MissingEvidenceItemCount != 0 ||
		preview.CompatibilityCenterEvidenceItemCount != 5 ||
		preview.RuntimeDiagnosticsEvidenceItemCount != 5 ||
		preview.RawExposedEvidenceItemCount != 0 ||
		preview.SideEffectEvidenceItemCount != 0 {
		t.Fatalf("unexpected KDE-safe redacted status dry-run lookup result consumer projection route enablement evidence gate counts: %#v", preview)
	}
	for _, item := range preview.EvidenceItems {
		if !item.EvidencePresent ||
			!item.CurrentMainlineConsumed ||
			!item.RouteAuthorizationEvidenceGateConsumed ||
			!item.RouteAuthorizationEvidenceGateReady ||
			!item.ResultConsumerProjectionRouteEnablementGated ||
			!item.RedactedProjectionRouteEnablementGated ||
			!item.UserVisibleRouteEnablementGated ||
			!item.FailureProjectionRouteEnablementGated ||
			!item.KDESafeRedactedStatusOnly ||
			item.LookupRouteEnablementAuthorized ||
			item.LookupRouteEnabled ||
			item.ConsumerEnablementAuthorized ||
			item.StorageWriteEnabled ||
			item.RawResultExposed ||
			!item.UserVisible ||
			!item.ReviewOnly ||
			!item.RuntimeOwned ||
			!item.GoRuntimeBacked ||
			item.KDEPolicyOwner ||
			!item.SideEffectsDisabled ||
			item.HostRootModified ||
			item.InternalDetailsExposed ||
			item.RouteEnablementGateStatus != "redacted-status-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-evidence-gated-routes-disabled" {
			t.Fatalf("unsafe KDE-safe redacted status dry-run lookup result consumer projection route enablement evidence gate item: %#v", item)
		}
	}
	if preview.Counts.Total != 8 ||
		preview.Counts.Passed != 8 ||
		preview.Counts.Pending != 0 ||
		preview.Counts.Blocked != 0 ||
		!sameStrings(preview.CheckIDs, []string{"current-mainline-consumed", "route-authorization-evidence-gate-consumed", "result-consumer-projection-route-enablement-gated", "compatibility-center-and-runtime-route-enablement-gated", "ten-route-enablement-evidence-items-ready-routes-disabled", "route-enablement-evidence-and-writes-disabled", "consumer-routes-dispatch-and-support-disabled", "production-and-host-boundary-closed"}) {
		t.Fatalf("unexpected KDE-safe redacted status dry-run lookup result consumer projection route enablement evidence gate checks: counts=%#v ids=%#v", preview.Counts, preview.CheckIDs)
	}
	if err := validateNoBackendTerms(preview, "KDE-safe redacted status dry-run lookup result consumer projection route enablement evidence gate test"); err != nil {
		t.Fatalf("validateNoBackendTerms returned error: %v", err)
	}
}

func TestProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementEvidenceGatePreviewFailsClosedWithoutSources(t *testing.T) {
	root := t.TempDir()
	if err := os.WriteFile(filepath.Join(root, "VERSION"), []byte(currentProjectVersion(t)+"\n"), 0o600); err != nil {
		t.Fatalf("WriteFile VERSION returned error: %v", err)
	}
	preview, err := NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementEvidenceGatePreview(root)
	if err != nil {
		t.Fatalf("NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementEvidenceGatePreview returned error: %v", err)
	}
	if preview.CurrentMainlineConsumed ||
		preview.RouteAuthorizationEvidenceGateConsumed ||
		preview.RouteAuthorizationEvidenceGateReady ||
		!preview.RouteEnablementEvidenceGateRequired ||
		preview.ResultConsumerProjectionRouteEnablementGated ||
		preview.RouteEnablementEvidenceGateReady ||
		preview.CompatibilityCenterRouteEnablementGated ||
		preview.RuntimeDiagnosticsRouteEnablementGated ||
		preview.RedactedProjectionRouteEnablementGated ||
		preview.UserVisibleRouteEnablementGated ||
		preview.FailureProjectionRouteEnablementGated ||
		preview.KDESafeRedactedStatusOnly ||
		preview.StorageWriteEnabled ||
		preview.RawResultExposed ||
		preview.EvidenceItemCount != 10 ||
		preview.ReadyEvidenceItemCount != 0 ||
		preview.MissingEvidenceItemCount != 10 ||
		preview.Counts.Total != 8 ||
		preview.Counts.Passed != 3 ||
		preview.Counts.Blocked != 5 ||
		preview.AuditDecision != "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-evidence-gate-blocked" {
		t.Fatalf("missing KDE-safe redacted status dry-run lookup result consumer projection route enablement evidence gate sources must fail closed: %#v", preview)
	}
	for _, item := range preview.EvidenceItems {
		if item.EvidencePresent ||
			item.CurrentMainlineConsumed ||
			item.RouteAuthorizationEvidenceGateConsumed ||
			item.RouteAuthorizationEvidenceGateReady ||
			item.ResultConsumerProjectionRouteEnablementGated ||
			item.RedactedProjectionRouteEnablementGated ||
			item.UserVisibleRouteEnablementGated ||
			item.FailureProjectionRouteEnablementGated ||
			item.KDESafeRedactedStatusOnly ||
			item.LookupRouteEnablementAuthorized ||
			item.LookupRouteEnabled ||
			item.ConsumerEnablementAuthorized ||
			item.StorageWriteEnabled ||
			item.RawResultExposed ||
			item.UserVisible ||
			item.RouteEnablementGateStatus != "missing-redacted-status-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-evidence-gate" {
			t.Fatalf("missing KDE-safe redacted status dry-run lookup result consumer projection route enablement evidence gate item must remain closed: %#v", item)
		}
	}
}
