package owner

import (
	"os"
	"path/filepath"
	"testing"
)

func TestProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementReceiptAcceptanceAuditPreviewModelsAcceptance(t *testing.T) {
	preview, err := NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementReceiptAcceptanceAuditPreview(projectRootForOwnerTest(t))
	if err != nil {
		t.Fatalf("NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementReceiptAcceptanceAuditPreview returned error: %v", err)
	}

	if preview.Version != currentProjectVersion(t) ||
		preview.SchemaVersion != "xnix.runtime.production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_kde_safe_redacted_status_closed_record_writer_storage_root_authorization_receipt_dry_run_lookup_result_consumer_projection_route_enablement_receipt_acceptance_audit.v1" ||
		preview.RequestType != "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-receipt-acceptance-audit-preview" ||
		preview.AuditType != "receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-receipt-acceptance-audit" ||
		preview.AuditDecision != "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-receipt-acceptance-audit-ready-acceptance-disabled-routes-disabled" {
		t.Fatalf("unexpected KDE-safe redacted status dry-run lookup result consumer projection route enablement receipt acceptance schema: %#v", preview)
	}
	if !preview.CurrentMainlineConsumed ||
		!preview.RouteEnablementReceiptGateConsumed ||
		!preview.RouteEnablementReceiptGateReady ||
		!preview.AcceptanceAuthorizationRequired ||
		!preview.AcceptanceAuthorizationModeled ||
		!preview.AcceptanceAuthorizationReady ||
		!preview.RouteEnablementAcceptanceBoundaryReady ||
		!preview.OpaqueRouteEnablementReceipt ||
		!preview.KDESafeRedactedStatusOnly ||
		!preview.CompatibilityCenterAcceptanceModeled ||
		!preview.RuntimeDiagnosticsAcceptanceModeled ||
		preview.ReceiptPresent ||
		preview.ReceiptAccepted ||
		preview.ReceiptConsumed ||
		preview.AcceptanceAuthorized ||
		preview.RouteEnablementAccepted ||
		preview.LookupRouteEnablementGranted ||
		preview.LookupRouteEnabled ||
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
		t.Fatalf("unsafe KDE-safe redacted status dry-run lookup result consumer projection route enablement receipt acceptance decision: %#v", preview)
	}
	if preview.AcceptanceItemCount != 10 ||
		preview.RequiredAcceptanceItemCount != 10 ||
		preview.ReadyAcceptanceItemCount != 10 ||
		preview.MissingAcceptanceItemCount != 0 ||
		preview.AcceptedReceiptItemCount != 0 ||
		preview.EnabledRouteItemCount != 0 ||
		preview.PersistedAcceptanceItemCount != 0 ||
		preview.RawExposedAcceptanceItemCount != 0 ||
		preview.SideEffectAcceptanceItemCount != 0 ||
		preview.CompatibilityCenterAcceptanceItemCount != 5 ||
		preview.RuntimeDiagnosticsAcceptanceItemCount != 5 {
		t.Fatalf("unexpected KDE-safe redacted status dry-run lookup result consumer projection route enablement receipt acceptance counts: %#v", preview)
	}
	if !sameStrings(preview.AcceptanceItemIDs, []string{
		"review-route-enablement-receipt-acceptance-compatibility-center",
		"review-route-enablement-receipt-acceptance-runtime-diagnostics",
		"renew-route-enablement-receipt-acceptance-compatibility-center",
		"renew-route-enablement-receipt-acceptance-runtime-diagnostics",
		"open-compatibility-center-route-enablement-receipt-acceptance-compatibility-center",
		"open-compatibility-center-route-enablement-receipt-acceptance-runtime-diagnostics",
		"dismiss-route-enablement-receipt-acceptance-compatibility-center",
		"dismiss-route-enablement-receipt-acceptance-runtime-diagnostics",
		"support-info-route-enablement-receipt-acceptance-compatibility-center",
		"support-info-route-enablement-receipt-acceptance-runtime-diagnostics",
	}) {
		t.Fatalf("unexpected KDE-safe redacted status dry-run lookup result consumer projection route enablement receipt acceptance items: %#v", preview.AcceptanceItemIDs)
	}
	for _, item := range preview.AcceptanceItems {
		if !item.EvidencePresent ||
			!item.CurrentMainlineConsumed ||
			!item.RouteEnablementReceiptGateConsumed ||
			!item.RouteEnablementReceiptGateReady ||
			!item.AcceptanceAuthorizationModeled ||
			!item.RouteEnablementAcceptanceBoundaryReady ||
			!item.OpaqueRouteEnablementReceipt ||
			!item.KDESafeRedactedStatusOnly ||
			item.ReceiptPresent ||
			item.ReceiptAccepted ||
			item.ReceiptConsumed ||
			item.AcceptanceAuthorized ||
			item.RouteEnablementAccepted ||
			item.LookupRouteEnablementGranted ||
			item.LookupRouteEnabled ||
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
			item.AcceptanceAuthorizationStatus != "redacted-status-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-receipt-acceptance-modeled-acceptance-disabled-routes-disabled" {
			t.Fatalf("unsafe KDE-safe redacted status dry-run lookup result consumer projection route enablement receipt acceptance item: %#v", item)
		}
	}
	if preview.Counts.Total != 8 ||
		preview.Counts.Passed != 8 ||
		preview.Counts.Pending != 0 ||
		preview.Counts.Blocked != 0 ||
		!sameStrings(preview.CheckIDs, []string{"current-mainline-consumed", "route-enablement-receipt-gate-consumed", "acceptance-authorization-boundary-modeled", "compatibility-center-and-runtime-acceptance-modeled", "ten-acceptance-items-ready-acceptance-disabled", "receipt-acceptance-routes-and-persistence-disabled", "consumer-routes-dispatch-and-support-disabled", "production-and-host-boundary-closed"}) {
		t.Fatalf("unexpected KDE-safe redacted status dry-run lookup result consumer projection route enablement receipt acceptance checks: counts=%#v ids=%#v", preview.Counts, preview.CheckIDs)
	}
	if preview.SystemServiceStarted ||
		preview.SessionBusClaimed ||
		preview.ProductionBusClaimed ||
		preview.WriteMethodsEnabled ||
		preview.RuntimeWritesEnabled ||
		preview.DesktopFilesWritten ||
		preview.KDEConfigurationWritten ||
		preview.PortalCallExecuted ||
		preview.AdapterInvocationEnabled ||
		preview.BackendLaunchEnabled ||
		preview.BackendProcessStarted ||
		preview.NetworkRequired ||
		preview.HostRootModified ||
		preview.PrivilegedContainerRequired ||
		preview.StateRootPathExposed ||
		preview.FilePathsExposed ||
		preview.FileContentRead ||
		preview.RawCommandExposed ||
		preview.RawExecutableExposed ||
		preview.BackendDetailsExposed {
		t.Fatalf("unsafe KDE-safe redacted status dry-run lookup result consumer projection route enablement receipt acceptance gates: %#v", preview)
	}
	if err := validateNoBackendTerms(preview, "KDE-safe redacted status dry-run lookup result consumer projection route enablement receipt acceptance test"); err != nil {
		t.Fatalf("validateNoBackendTerms returned error: %v", err)
	}
}

func TestProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementReceiptAcceptanceAuditPreviewFailsClosedWithoutSources(t *testing.T) {
	root := t.TempDir()
	if err := os.WriteFile(filepath.Join(root, "VERSION"), []byte(currentProjectVersion(t)+"\n"), 0o600); err != nil {
		t.Fatalf("WriteFile VERSION returned error: %v", err)
	}
	preview, err := NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementReceiptAcceptanceAuditPreview(root)
	if err != nil {
		t.Fatalf("NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementReceiptAcceptanceAuditPreview returned error: %v", err)
	}
	if preview.CurrentMainlineConsumed ||
		preview.RouteEnablementReceiptGateConsumed ||
		preview.RouteEnablementReceiptGateReady ||
		!preview.AcceptanceAuthorizationRequired ||
		!preview.AcceptanceAuthorizationModeled ||
		preview.AcceptanceAuthorizationReady ||
		preview.RouteEnablementAcceptanceBoundaryReady ||
		preview.OpaqueRouteEnablementReceipt ||
		preview.KDESafeRedactedStatusOnly ||
		preview.CompatibilityCenterAcceptanceModeled ||
		preview.RuntimeDiagnosticsAcceptanceModeled ||
		preview.ReceiptPresent ||
		preview.ReceiptAccepted ||
		preview.ReceiptConsumed ||
		preview.AcceptanceAuthorized ||
		preview.RouteEnablementAccepted ||
		preview.LookupRouteEnablementGranted ||
		preview.LookupRouteEnabled ||
		preview.StorageWriteEnabled ||
		preview.RawResultExposed ||
		preview.AcceptanceItemCount != 10 ||
		preview.RequiredAcceptanceItemCount != 10 ||
		preview.ReadyAcceptanceItemCount != 0 ||
		preview.MissingAcceptanceItemCount != 10 ||
		preview.Counts.Total != 8 ||
		preview.Counts.Passed != 4 ||
		preview.Counts.Blocked != 4 ||
		preview.AuditDecision != "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-receipt-acceptance-audit-blocked" {
		t.Fatalf("missing KDE-safe redacted status dry-run lookup result consumer projection route enablement receipt acceptance sources must fail closed: %#v", preview)
	}
	for _, item := range preview.AcceptanceItems {
		if item.EvidencePresent ||
			item.CurrentMainlineConsumed ||
			item.RouteEnablementReceiptGateConsumed ||
			item.RouteEnablementReceiptGateReady ||
			item.AcceptanceAuthorizationModeled ||
			item.RouteEnablementAcceptanceBoundaryReady ||
			item.OpaqueRouteEnablementReceipt ||
			item.ReceiptAccepted ||
			item.ReceiptConsumed ||
			item.AcceptanceAuthorized ||
			item.RouteEnablementAccepted ||
			item.LookupRouteEnabled ||
			item.StorageWriteEnabled ||
			item.RawResultExposed ||
			item.UserVisible ||
			item.AcceptanceAuthorizationStatus != "missing-redacted-status-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-receipt-acceptance-evidence" {
			t.Fatalf("missing KDE-safe redacted status dry-run lookup result consumer projection route enablement receipt acceptance item must remain closed: %#v", item)
		}
	}
}
