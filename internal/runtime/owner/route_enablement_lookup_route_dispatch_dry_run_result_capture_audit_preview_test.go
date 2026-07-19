package owner

import (
	"os"
	"path/filepath"
	"testing"
)

func TestProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultCaptureAuditPreviewModelsDispatchDryRunResultCapture(t *testing.T) {
	preview, err := NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultCaptureAuditPreview(projectRootForOwnerTest(t))
	if err != nil {
		t.Fatalf("NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultCaptureAuditPreview returned error: %v", err)
	}

	if preview.Version != currentProjectVersion(t) ||
		preview.SchemaVersion != "xnix.runtime.production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_kde_safe_redacted_status_closed_record_writer_storage_root_authorization_receipt_dry_run_lookup_result_consumer_projection_route_enablement_lookup_route_dispatch_dry_run_result_capture_audit.v1" ||
		preview.RequestType != "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-capture-audit-preview" ||
		preview.AuditType != "receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-capture-audit" ||
		preview.AuditDecision != "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-capture-audit-ready-route-disabled-dispatch-disabled" {
		t.Fatalf("unexpected KDE-safe redacted status dry-run lookup result consumer projection route enablement lookup route dispatch dry-run result capture schema: %#v", preview)
	}
	if !preview.CurrentMainlineConsumed ||
		!preview.RouteEnablementLookupRouteDispatchDryRunExecutionGateConsumed ||
		!preview.RouteEnablementLookupRouteDispatchDryRunExecutionGateReady ||
		!preview.LookupRouteDispatchDryRunResultCaptureRequired ||
		!preview.LookupRouteDispatchDryRunResultCaptureModeled ||
		!preview.LookupRouteDispatchDryRunResultCaptureReady ||
		!preview.RouteEnablementLookupRouteDispatchDryRunResultCaptureBoundaryReady ||
		!preview.OpaqueRouteEnablementReceipt ||
		!preview.KDESafeRedactedStatusOnly ||
		!preview.CompatibilityCenterLookupRouteDispatchDryRunResultCaptureModeled ||
		!preview.RuntimeDiagnosticsLookupRouteDispatchDryRunResultCaptureModeled ||
		preview.ReceiptPresent ||
		preview.ReceiptAccepted ||
		preview.ReceiptConsumed ||
		preview.AcceptanceAuthorized ||
		preview.RouteEnablementAccepted ||
		preview.LookupRouteEnablementGranted ||
		preview.LookupRouteGrantAuthorized ||
		preview.LookupRouteEnabled ||
		preview.LookupRouteDispatchAuthorized ||
		preview.LookupRouteDispatchDryRunResultCapturePassed ||
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
		t.Fatalf("unsafe KDE-safe redacted status dry-run lookup result consumer projection route enablement lookup route dispatch dry-run result capture decision: %#v", preview)
	}
	if preview.DispatchDryRunResultCaptureItemCount != 10 ||
		preview.RequiredDispatchDryRunResultCaptureItemCount != 10 ||
		preview.ReadyDispatchDryRunResultCaptureItemCount != 10 ||
		preview.MissingDispatchDryRunResultCaptureItemCount != 0 ||
		preview.PassedDispatchDryRunResultCaptureItemCount != 0 ||
		preview.EnabledRouteItemCount != 0 ||
		preview.PersistedDispatchDryRunResultCaptureItemCount != 0 ||
		preview.RawExposedDispatchDryRunResultCaptureItemCount != 0 ||
		preview.SideEffectDispatchDryRunResultCaptureItemCount != 0 ||
		preview.CompatibilityCenterDispatchDryRunResultCaptureItemCount != 5 ||
		preview.RuntimeDiagnosticsDispatchDryRunResultCaptureItemCount != 5 {
		t.Fatalf("unexpected KDE-safe redacted status dry-run lookup result consumer projection route enablement lookup route dispatch dry-run result capture counts: %#v", preview)
	}
	if !sameStrings(preview.DispatchDryRunResultCaptureItemIDs, []string{
		"review-route-enablement-lookup-route-dispatch-dry-run-result-capture-compatibility-center",
		"review-route-enablement-lookup-route-dispatch-dry-run-result-capture-runtime-diagnostics",
		"renew-route-enablement-lookup-route-dispatch-dry-run-result-capture-compatibility-center",
		"renew-route-enablement-lookup-route-dispatch-dry-run-result-capture-runtime-diagnostics",
		"open-compatibility-center-route-enablement-lookup-route-dispatch-dry-run-result-capture-compatibility-center",
		"open-compatibility-center-route-enablement-lookup-route-dispatch-dry-run-result-capture-runtime-diagnostics",
		"dismiss-route-enablement-lookup-route-dispatch-dry-run-result-capture-compatibility-center",
		"dismiss-route-enablement-lookup-route-dispatch-dry-run-result-capture-runtime-diagnostics",
		"support-info-route-enablement-lookup-route-dispatch-dry-run-result-capture-compatibility-center",
		"support-info-route-enablement-lookup-route-dispatch-dry-run-result-capture-runtime-diagnostics",
	}) {
		t.Fatalf("unexpected KDE-safe redacted status dry-run lookup result consumer projection route enablement lookup route dispatch dry-run result capture items: %#v", preview.DispatchDryRunResultCaptureItemIDs)
	}
	for _, item := range preview.DispatchDryRunResultCaptureItems {
		if !item.EvidencePresent ||
			!item.CurrentMainlineConsumed ||
			!item.RouteEnablementLookupRouteDispatchDryRunExecutionGateConsumed ||
			!item.RouteEnablementLookupRouteDispatchDryRunExecutionGateReady ||
			!item.LookupRouteDispatchDryRunResultCaptureModeled ||
			!item.RouteEnablementLookupRouteDispatchDryRunResultCaptureBoundaryReady ||
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
			item.LookupRouteDispatchDryRunResultCapturePassed ||
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
			item.LookupRouteDispatchDryRunResultCaptureStatus != "redacted-status-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-capture-modeled-dispatch-disabled-route-disabled" {
			t.Fatalf("unsafe KDE-safe redacted status dry-run lookup result consumer projection route enablement lookup route dispatch dry-run result capture item: %#v", item)
		}
	}
	if preview.Counts.Total != 8 ||
		preview.Counts.Passed != 8 ||
		preview.Counts.Pending != 0 ||
		preview.Counts.Blocked != 0 ||
		!sameStrings(preview.CheckIDs, []string{"current-mainline-consumed", "route-enablement-lookup-route-dispatch-dry-run-execution-gate-consumed", "lookup-route-dispatch-dry-run-result-capture-modeled", "compatibility-center-and-runtime-lookup-route-dispatch-dry-run-result-captures-modeled", "ten-dispatch-dry-run-result-capture-items-ready-dispatch-disabled", "lookup-route-dispatch-dry-run-result-capture-and-persistence-disabled", "consumer-routes-dispatch-and-support-disabled", "production-and-host-boundary-closed"}) {
		t.Fatalf("unexpected KDE-safe redacted status dry-run lookup result consumer projection route enablement lookup route dispatch dry-run result capture checks: counts=%#v ids=%#v", preview.Counts, preview.CheckIDs)
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
		t.Fatalf("unsafe KDE-safe redacted status dry-run lookup result consumer projection route enablement lookup route dispatch dry-run result capture gates: %#v", preview)
	}
	if err := validateNoBackendTerms(preview, "KDE-safe redacted status dry-run lookup result consumer projection route enablement lookup route dispatch dry-run result capture test"); err != nil {
		t.Fatalf("validateNoBackendTerms returned error: %v", err)
	}
}

func TestProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultCaptureAuditPreviewFailsClosedWithoutSources(t *testing.T) {
	root, err := os.MkdirTemp("", "xnix-lookup-dispatch-dry-run-result-capture-*")
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
	preview, err := NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultCaptureAuditPreview(root)
	if err != nil {
		t.Fatalf("NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupRouteDispatchDryRunResultCaptureAuditPreview returned error: %v", err)
	}
	if preview.CurrentMainlineConsumed ||
		preview.RouteEnablementLookupRouteDispatchDryRunExecutionGateConsumed ||
		preview.RouteEnablementLookupRouteDispatchDryRunExecutionGateReady ||
		!preview.LookupRouteDispatchDryRunResultCaptureRequired ||
		!preview.LookupRouteDispatchDryRunResultCaptureModeled ||
		preview.LookupRouteDispatchDryRunResultCaptureReady ||
		preview.RouteEnablementLookupRouteDispatchDryRunResultCaptureBoundaryReady ||
		preview.OpaqueRouteEnablementReceipt ||
		preview.KDESafeRedactedStatusOnly ||
		preview.CompatibilityCenterLookupRouteDispatchDryRunResultCaptureModeled ||
		preview.RuntimeDiagnosticsLookupRouteDispatchDryRunResultCaptureModeled ||
		preview.ReceiptPresent ||
		preview.ReceiptAccepted ||
		preview.ReceiptConsumed ||
		preview.AcceptanceAuthorized ||
		preview.RouteEnablementAccepted ||
		preview.LookupRouteEnablementGranted ||
		preview.LookupRouteGrantAuthorized ||
		preview.LookupRouteEnabled ||
		preview.LookupRouteDispatchAuthorized ||
		preview.LookupRouteDispatchDryRunResultCapturePassed ||
		preview.LookupRouteDispatchCallable ||
		preview.StorageWriteEnabled ||
		preview.RawResultExposed ||
		preview.DispatchDryRunResultCaptureItemCount != 10 ||
		preview.RequiredDispatchDryRunResultCaptureItemCount != 10 ||
		preview.ReadyDispatchDryRunResultCaptureItemCount != 0 ||
		preview.MissingDispatchDryRunResultCaptureItemCount != 10 ||
		preview.Counts.Total != 8 ||
		preview.Counts.Passed != 4 ||
		preview.Counts.Blocked != 4 ||
		preview.AuditDecision != "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-capture-audit-blocked" {
		t.Fatalf("missing KDE-safe redacted status dry-run lookup result consumer projection route enablement lookup route dispatch dry-run result capture sources must fail closed: %#v", preview)
	}
	for _, item := range preview.DispatchDryRunResultCaptureItems {
		if item.EvidencePresent ||
			item.CurrentMainlineConsumed ||
			item.RouteEnablementLookupRouteDispatchDryRunExecutionGateConsumed ||
			item.RouteEnablementLookupRouteDispatchDryRunExecutionGateReady ||
			item.LookupRouteDispatchDryRunResultCaptureModeled ||
			item.RouteEnablementLookupRouteDispatchDryRunResultCaptureBoundaryReady ||
			item.OpaqueRouteEnablementReceipt ||
			item.ReceiptAccepted ||
			item.ReceiptConsumed ||
			item.AcceptanceAuthorized ||
			item.RouteEnablementAccepted ||
			item.LookupRouteEnablementGranted ||
			item.LookupRouteGrantAuthorized ||
			item.LookupRouteEnabled ||
			item.LookupRouteDispatchAuthorized ||
			item.LookupRouteDispatchDryRunResultCapturePassed ||
			item.LookupRouteDispatchCallable ||
			item.StorageWriteEnabled ||
			item.RawResultExposed ||
			item.UserVisible ||
			item.LookupRouteDispatchDryRunResultCaptureStatus != "missing-redacted-status-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-capture-evidence" {
			t.Fatalf("missing KDE-safe redacted status dry-run lookup result consumer projection route enablement lookup route dispatch dry-run result capture item must remain closed: %#v", item)
		}
	}
}
