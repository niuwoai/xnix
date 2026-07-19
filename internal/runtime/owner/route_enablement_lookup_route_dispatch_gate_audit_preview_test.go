package owner

import (
	"os"
	"path/filepath"
	"testing"
)

func TestProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchGateAuditPreviewModelsDispatchGate(t *testing.T) {
	preview, err := NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchGateAuditPreview(projectRootForOwnerTest(t))
	if err != nil {
		t.Fatalf("NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchGateAuditPreview returned error: %v", err)
	}

	if preview.Version != currentProjectVersion(t) ||
		preview.SchemaVersion != "xnix.runtime.production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_kde_safe_redacted_status_closed_record_writer_storage_root_authorization_receipt_dry_run_lookup_result_consumer_projection_route_enablement_lookup_route_dispatch_gate_audit.v1" ||
		preview.RequestType != "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-gate-audit-preview" ||
		preview.AuditType != "receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-gate-audit" ||
		preview.AuditDecision != "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-gate-audit-ready-route-disabled-dispatch-disabled" {
		t.Fatalf("unexpected KDE-safe redacted status dry-run lookup result consumer projection route enablement lookup route dispatch gate schema: %#v", preview)
	}
	if !preview.CurrentMainlineConsumed ||
		!preview.RouteEnablementLookupRouteEnablementConsumed ||
		!preview.RouteEnablementLookupRouteEnablementReady ||
		!preview.LookupRouteDispatchGateRequired ||
		!preview.LookupRouteDispatchGateModeled ||
		!preview.LookupRouteDispatchGateReady ||
		!preview.RouteEnablementLookupRouteDispatchGateBoundaryReady ||
		!preview.OpaqueRouteEnablementReceipt ||
		!preview.KDESafeRedactedStatusOnly ||
		!preview.CompatibilityCenterLookupRouteDispatchGateModeled ||
		!preview.RuntimeDiagnosticsLookupRouteDispatchGateModeled ||
		preview.ReceiptPresent ||
		preview.ReceiptAccepted ||
		preview.ReceiptConsumed ||
		preview.AcceptanceAuthorized ||
		preview.RouteEnablementAccepted ||
		preview.LookupRouteEnablementGranted ||
		preview.LookupRouteGrantAuthorized ||
		preview.LookupRouteEnabled ||
		preview.LookupRouteDispatchAuthorized ||
		preview.LookupRouteDispatchGatePassed ||
		preview.LookupRouteDispatchCallable ||
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
		t.Fatalf("unsafe KDE-safe redacted status dry-run lookup result consumer projection route enablement lookup route dispatch gate decision: %#v", preview)
	}
	if preview.DispatchGateItemCount != 10 ||
		preview.RequiredDispatchGateItemCount != 10 ||
		preview.ReadyDispatchGateItemCount != 10 ||
		preview.MissingDispatchGateItemCount != 0 ||
		preview.PassedDispatchGateItemCount != 0 ||
		preview.EnabledRouteItemCount != 0 ||
		preview.PersistedDispatchGateItemCount != 0 ||
		preview.RawExposedDispatchGateItemCount != 0 ||
		preview.SideEffectDispatchGateItemCount != 0 ||
		preview.CompatibilityCenterDispatchGateItemCount != 5 ||
		preview.RuntimeDiagnosticsDispatchGateItemCount != 5 {
		t.Fatalf("unexpected KDE-safe redacted status dry-run lookup result consumer projection route enablement lookup route dispatch gate counts: %#v", preview)
	}
	if !sameStrings(preview.DispatchGateItemIDs, []string{
		"review-route-enablement-lookup-route-dispatch-gate-compatibility-center",
		"review-route-enablement-lookup-route-dispatch-gate-runtime-diagnostics",
		"renew-route-enablement-lookup-route-dispatch-gate-compatibility-center",
		"renew-route-enablement-lookup-route-dispatch-gate-runtime-diagnostics",
		"open-compatibility-center-route-enablement-lookup-route-dispatch-gate-compatibility-center",
		"open-compatibility-center-route-enablement-lookup-route-dispatch-gate-runtime-diagnostics",
		"dismiss-route-enablement-lookup-route-dispatch-gate-compatibility-center",
		"dismiss-route-enablement-lookup-route-dispatch-gate-runtime-diagnostics",
		"support-info-route-enablement-lookup-route-dispatch-gate-compatibility-center",
		"support-info-route-enablement-lookup-route-dispatch-gate-runtime-diagnostics",
	}) {
		t.Fatalf("unexpected KDE-safe redacted status dry-run lookup result consumer projection route enablement lookup route dispatch gate items: %#v", preview.DispatchGateItemIDs)
	}
	for _, item := range preview.DispatchGateItems {
		if !item.EvidencePresent ||
			!item.CurrentMainlineConsumed ||
			!item.RouteEnablementLookupRouteEnablementConsumed ||
			!item.RouteEnablementLookupRouteEnablementReady ||
			!item.LookupRouteDispatchGateModeled ||
			!item.RouteEnablementLookupRouteDispatchGateBoundaryReady ||
			!item.OpaqueRouteEnablementReceipt ||
			!item.KDESafeRedactedStatusOnly ||
			item.ReceiptPresent ||
			item.ReceiptAccepted ||
			item.ReceiptConsumed ||
			item.AcceptanceAuthorized ||
			item.RouteEnablementAccepted ||
			item.LookupRouteEnablementGranted ||
			item.LookupRouteGrantAuthorized ||
			item.LookupRouteEnabled ||
			item.LookupRouteDispatchAuthorized ||
			item.LookupRouteDispatchGatePassed ||
			item.LookupRouteDispatchCallable ||
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
			item.LookupRouteDispatchGateStatus != "redacted-status-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-gate-modeled-dispatch-disabled-route-disabled" {
			t.Fatalf("unsafe KDE-safe redacted status dry-run lookup result consumer projection route enablement lookup route dispatch gate item: %#v", item)
		}
	}
	if preview.Counts.Total != 8 ||
		preview.Counts.Passed != 8 ||
		preview.Counts.Pending != 0 ||
		preview.Counts.Blocked != 0 ||
		!sameStrings(preview.CheckIDs, []string{"current-mainline-consumed", "route-enablement-lookup-route-enablement-consumed", "lookup-route-dispatch-gate-modeled", "compatibility-center-and-runtime-lookup-route-dispatch-gates-modeled", "ten-dispatch-gate-items-ready-dispatch-disabled", "lookup-route-dispatch-gate-and-persistence-disabled", "consumer-routes-dispatch-and-support-disabled", "production-and-host-boundary-closed"}) {
		t.Fatalf("unexpected KDE-safe redacted status dry-run lookup result consumer projection route enablement lookup route dispatch gate checks: counts=%#v ids=%#v", preview.Counts, preview.CheckIDs)
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
		t.Fatalf("unsafe KDE-safe redacted status dry-run lookup result consumer projection route enablement lookup route dispatch gate gates: %#v", preview)
	}
	if err := validateNoBackendTerms(preview, "KDE-safe redacted status dry-run lookup result consumer projection route enablement lookup route dispatch gate test"); err != nil {
		t.Fatalf("validateNoBackendTerms returned error: %v", err)
	}
}

func TestProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchGateAuditPreviewFailsClosedWithoutSources(t *testing.T) {
	root, err := os.MkdirTemp("", "xnix-lookup-dispatch-gate-*")
	if err != nil {
		t.Fatalf("MkdirTemp returned error: %v", err)
	}
	t.Cleanup(func() {
		if err := os.RemoveAll(root); err != nil {
			t.Fatalf("RemoveAll returned error: %v", err)
		}
	})
	if err := os.WriteFile(filepath.Join(root, "VERSION"), []byte(currentProjectVersion(t)+"\n"), 0o600); err != nil {
		t.Fatalf("WriteFile VERSION returned error: %v", err)
	}
	preview, err := NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchGateAuditPreview(root)
	if err != nil {
		t.Fatalf("NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupRouteDispatchGateAuditPreview returned error: %v", err)
	}
	if preview.CurrentMainlineConsumed ||
		preview.RouteEnablementLookupRouteEnablementConsumed ||
		preview.RouteEnablementLookupRouteEnablementReady ||
		!preview.LookupRouteDispatchGateRequired ||
		!preview.LookupRouteDispatchGateModeled ||
		preview.LookupRouteDispatchGateReady ||
		preview.RouteEnablementLookupRouteDispatchGateBoundaryReady ||
		preview.OpaqueRouteEnablementReceipt ||
		preview.KDESafeRedactedStatusOnly ||
		preview.CompatibilityCenterLookupRouteDispatchGateModeled ||
		preview.RuntimeDiagnosticsLookupRouteDispatchGateModeled ||
		preview.ReceiptPresent ||
		preview.ReceiptAccepted ||
		preview.ReceiptConsumed ||
		preview.AcceptanceAuthorized ||
		preview.RouteEnablementAccepted ||
		preview.LookupRouteEnablementGranted ||
		preview.LookupRouteGrantAuthorized ||
		preview.LookupRouteEnabled ||
		preview.LookupRouteDispatchAuthorized ||
		preview.LookupRouteDispatchGatePassed ||
		preview.LookupRouteDispatchCallable ||
		preview.StorageWriteEnabled ||
		preview.RawResultExposed ||
		preview.DispatchGateItemCount != 10 ||
		preview.RequiredDispatchGateItemCount != 10 ||
		preview.ReadyDispatchGateItemCount != 0 ||
		preview.MissingDispatchGateItemCount != 10 ||
		preview.Counts.Total != 8 ||
		preview.Counts.Passed != 4 ||
		preview.Counts.Blocked != 4 ||
		preview.AuditDecision != "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-gate-audit-blocked" {
		t.Fatalf("missing KDE-safe redacted status dry-run lookup result consumer projection route enablement lookup route dispatch gate sources must fail closed: %#v", preview)
	}
	for _, item := range preview.DispatchGateItems {
		if item.EvidencePresent ||
			item.CurrentMainlineConsumed ||
			item.RouteEnablementLookupRouteEnablementConsumed ||
			item.RouteEnablementLookupRouteEnablementReady ||
			item.LookupRouteDispatchGateModeled ||
			item.RouteEnablementLookupRouteDispatchGateBoundaryReady ||
			item.OpaqueRouteEnablementReceipt ||
			item.ReceiptAccepted ||
			item.ReceiptConsumed ||
			item.AcceptanceAuthorized ||
			item.RouteEnablementAccepted ||
			item.LookupRouteEnablementGranted ||
			item.LookupRouteGrantAuthorized ||
			item.LookupRouteEnabled ||
			item.LookupRouteDispatchAuthorized ||
			item.LookupRouteDispatchGatePassed ||
			item.LookupRouteDispatchCallable ||
			item.StorageWriteEnabled ||
			item.RawResultExposed ||
			item.UserVisible ||
			item.LookupRouteDispatchGateStatus != "missing-redacted-status-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-gate-evidence" {
			t.Fatalf("missing KDE-safe redacted status dry-run lookup result consumer projection route enablement lookup route dispatch gate item must remain closed: %#v", item)
		}
	}
}
