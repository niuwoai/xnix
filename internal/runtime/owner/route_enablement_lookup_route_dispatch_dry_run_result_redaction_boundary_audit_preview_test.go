package owner

import (
	"os"
	"path/filepath"
	"testing"
)

func TestProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultRedactionBoundaryAuditPreviewModelsResultRedactionBoundary(t *testing.T) {
	preview, err := NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultRedactionBoundaryAuditPreview(projectRootForOwnerTest(t))
	if err != nil {
		t.Fatalf("NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultRedactionBoundaryAuditPreview returned error: %v", err)
	}

	if preview.Version != currentProjectVersion(t) ||
		preview.SchemaVersion != "xnix.runtime.production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_kde_safe_redacted_status_closed_record_writer_storage_root_authorization_receipt_dry_run_lookup_result_consumer_projection_route_enablement_lookup_route_dispatch_dry_run_result_redaction_boundary_audit.v1" ||
		preview.RequestType != "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-redaction-boundary-audit-preview" ||
		preview.AuditType != "receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-redaction-boundary-audit" ||
		preview.AuditDecision != "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-redaction-boundary-audit-ready-route-disabled-dispatch-disabled-raw-result-hidden" {
		t.Fatalf("unexpected KDE-safe redacted status dry-run lookup result consumer projection route enablement lookup route dispatch dry-run result redaction boundary schema: %#v", preview)
	}
	if !preview.CurrentMainlineConsumed ||
		!preview.RouteEnablementLookupRouteDispatchDryRunResultCaptureConsumed ||
		!preview.RouteEnablementLookupRouteDispatchDryRunResultCaptureReady ||
		!preview.LookupRouteDispatchDryRunResultRedactionBoundaryRequired ||
		!preview.LookupRouteDispatchDryRunResultRedactionBoundaryModeled ||
		!preview.LookupRouteDispatchDryRunResultRedactionBoundaryReady ||
		!preview.RouteEnablementLookupRouteDispatchDryRunResultRedactionBoundaryReady ||
		!preview.CapturedDryRunResultOpaque ||
		!preview.KDESafeRedactedResultOnly ||
		!preview.CompatibilityCenterLookupRouteDispatchDryRunResultRedactionModeled ||
		!preview.RuntimeDiagnosticsLookupRouteDispatchDryRunResultRedactionModeled ||
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
		preview.LookupRouteDispatchDryRunResultRedactionPassed ||
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
		t.Fatalf("unsafe KDE-safe redacted status dry-run lookup result consumer projection route enablement lookup route dispatch dry-run result redaction boundary decision: %#v", preview)
	}
	if preview.DispatchDryRunResultRedactionBoundaryItemCount != 10 ||
		preview.RequiredDispatchDryRunResultRedactionBoundaryItemCount != 10 ||
		preview.ReadyDispatchDryRunResultRedactionBoundaryItemCount != 10 ||
		preview.MissingDispatchDryRunResultRedactionBoundaryItemCount != 0 ||
		preview.PassedDispatchDryRunResultRedactionItemCount != 0 ||
		preview.EnabledRouteItemCount != 0 ||
		preview.PersistedDispatchDryRunResultRedactionItemCount != 0 ||
		preview.RawExposedDispatchDryRunResultItemCount != 0 ||
		preview.SideEffectDispatchDryRunResultRedactionItemCount != 0 ||
		preview.CompatibilityCenterDispatchDryRunResultRedactionItemCount != 5 ||
		preview.RuntimeDiagnosticsDispatchDryRunResultRedactionItemCount != 5 {
		t.Fatalf("unexpected KDE-safe redacted status dry-run lookup result consumer projection route enablement lookup route dispatch dry-run result redaction boundary counts: %#v", preview)
	}
	if !sameStrings(preview.DispatchDryRunResultRedactionBoundaryItemIDs, []string{
		"review-route-enablement-lookup-route-dispatch-dry-run-result-redaction-boundary-compatibility-center",
		"review-route-enablement-lookup-route-dispatch-dry-run-result-redaction-boundary-runtime-diagnostics",
		"renew-route-enablement-lookup-route-dispatch-dry-run-result-redaction-boundary-compatibility-center",
		"renew-route-enablement-lookup-route-dispatch-dry-run-result-redaction-boundary-runtime-diagnostics",
		"open-compatibility-center-route-enablement-lookup-route-dispatch-dry-run-result-redaction-boundary-compatibility-center",
		"open-compatibility-center-route-enablement-lookup-route-dispatch-dry-run-result-redaction-boundary-runtime-diagnostics",
		"dismiss-route-enablement-lookup-route-dispatch-dry-run-result-redaction-boundary-compatibility-center",
		"dismiss-route-enablement-lookup-route-dispatch-dry-run-result-redaction-boundary-runtime-diagnostics",
		"support-info-route-enablement-lookup-route-dispatch-dry-run-result-redaction-boundary-compatibility-center",
		"support-info-route-enablement-lookup-route-dispatch-dry-run-result-redaction-boundary-runtime-diagnostics",
	}) {
		t.Fatalf("unexpected KDE-safe redacted status dry-run lookup result consumer projection route enablement lookup route dispatch dry-run result redaction boundary items: %#v", preview.DispatchDryRunResultRedactionBoundaryItemIDs)
	}
	for _, item := range preview.DispatchDryRunResultRedactionBoundaryItems {
		if !item.EvidencePresent ||
			!item.CurrentMainlineConsumed ||
			!item.RouteEnablementLookupRouteDispatchDryRunResultCaptureConsumed ||
			!item.RouteEnablementLookupRouteDispatchDryRunResultCaptureReady ||
			!item.LookupRouteDispatchDryRunResultRedactionBoundaryModeled ||
			!item.RouteEnablementLookupRouteDispatchDryRunResultRedactionReady ||
			!item.CapturedDryRunResultOpaque ||
			!item.KDESafeRedactedResultOnly ||
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
			item.LookupRouteDispatchDryRunResultRedactionPassed ||
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
			item.LookupRouteDispatchDryRunResultRedactionBoundaryStatus != "redacted-status-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-redaction-boundary-modeled-raw-result-hidden-dispatch-disabled-route-disabled" {
			t.Fatalf("unsafe KDE-safe redacted status dry-run lookup result consumer projection route enablement lookup route dispatch dry-run result redaction boundary item: %#v", item)
		}
	}
	if preview.Counts.Total != 8 ||
		preview.Counts.Passed != 8 ||
		preview.Counts.Pending != 0 ||
		preview.Counts.Blocked != 0 ||
		!sameStrings(preview.CheckIDs, []string{"current-mainline-consumed", "route-enablement-lookup-route-dispatch-dry-run-result-capture-consumed", "lookup-route-dispatch-dry-run-result-redaction-boundary-modeled", "compatibility-center-and-runtime-lookup-route-dispatch-dry-run-result-redaction-boundaries-modeled", "ten-dispatch-dry-run-result-redaction-boundary-items-ready-dispatch-disabled", "raw-result-exposure-and-persistence-disabled", "consumer-routes-dispatch-and-support-disabled", "production-and-host-boundary-closed"}) {
		t.Fatalf("unexpected KDE-safe redacted status dry-run lookup result consumer projection route enablement lookup route dispatch dry-run result redaction boundary checks: counts=%#v ids=%#v", preview.Counts, preview.CheckIDs)
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
		t.Fatalf("unsafe KDE-safe redacted status dry-run lookup result consumer projection route enablement lookup route dispatch dry-run result redaction boundary gates: %#v", preview)
	}
	if err := validateNoBackendTerms(preview, "KDE-safe redacted status dry-run lookup result consumer projection route enablement lookup route dispatch dry-run result redaction boundary test"); err != nil {
		t.Fatalf("validateNoBackendTerms returned error: %v", err)
	}
}

func TestProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultRedactionBoundaryAuditPreviewFailsClosedWithoutSources(t *testing.T) {
	root, err := os.MkdirTemp("", "xnix-lookup-dispatch-dry-run-result-redaction-boundary-*")
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
	preview, err := NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultRedactionBoundaryAuditPreview(root)
	if err != nil {
		t.Fatalf("NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultRedactionBoundaryAuditPreview returned error: %v", err)
	}
	if preview.CurrentMainlineConsumed ||
		preview.RouteEnablementLookupRouteDispatchDryRunResultCaptureConsumed ||
		preview.RouteEnablementLookupRouteDispatchDryRunResultCaptureReady ||
		!preview.LookupRouteDispatchDryRunResultRedactionBoundaryRequired ||
		!preview.LookupRouteDispatchDryRunResultRedactionBoundaryModeled ||
		preview.LookupRouteDispatchDryRunResultRedactionBoundaryReady ||
		preview.RouteEnablementLookupRouteDispatchDryRunResultRedactionBoundaryReady ||
		preview.CapturedDryRunResultOpaque ||
		preview.KDESafeRedactedResultOnly ||
		preview.CompatibilityCenterLookupRouteDispatchDryRunResultRedactionModeled ||
		preview.RuntimeDiagnosticsLookupRouteDispatchDryRunResultRedactionModeled ||
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
		preview.LookupRouteDispatchDryRunResultRedactionPassed ||
		preview.LookupRouteDispatchCallable ||
		preview.StorageWriteEnabled ||
		preview.RawResultExposed ||
		preview.DispatchDryRunResultRedactionBoundaryItemCount != 10 ||
		preview.RequiredDispatchDryRunResultRedactionBoundaryItemCount != 10 ||
		preview.ReadyDispatchDryRunResultRedactionBoundaryItemCount != 0 ||
		preview.MissingDispatchDryRunResultRedactionBoundaryItemCount != 10 ||
		preview.Counts.Total != 8 ||
		preview.Counts.Passed != 3 ||
		preview.Counts.Blocked != 5 ||
		preview.AuditDecision != "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-redaction-boundary-audit-blocked" {
		t.Fatalf("missing KDE-safe redacted status dry-run lookup result consumer projection route enablement lookup route dispatch dry-run result redaction boundary sources must fail closed: %#v", preview)
	}
	for _, item := range preview.DispatchDryRunResultRedactionBoundaryItems {
		if item.EvidencePresent ||
			item.CurrentMainlineConsumed ||
			item.RouteEnablementLookupRouteDispatchDryRunResultCaptureConsumed ||
			item.RouteEnablementLookupRouteDispatchDryRunResultCaptureReady ||
			item.LookupRouteDispatchDryRunResultRedactionBoundaryModeled ||
			item.RouteEnablementLookupRouteDispatchDryRunResultRedactionReady ||
			item.CapturedDryRunResultOpaque ||
			item.ReceiptAccepted ||
			item.ReceiptConsumed ||
			item.AcceptanceAuthorized ||
			item.RouteEnablementAccepted ||
			item.LookupRouteEnablementGranted ||
			item.LookupRouteGrantAuthorized ||
			item.LookupRouteEnabled ||
			item.LookupRouteDispatchAuthorized ||
			item.LookupRouteDispatchDryRunResultCapturePassed ||
			item.LookupRouteDispatchDryRunResultRedactionPassed ||
			item.LookupRouteDispatchCallable ||
			item.StorageWriteEnabled ||
			item.RawResultExposed ||
			item.UserVisible ||
			item.LookupRouteDispatchDryRunResultRedactionBoundaryStatus != "missing-redacted-status-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-redaction-boundary-evidence" {
			t.Fatalf("missing KDE-safe redacted status dry-run lookup result consumer projection route enablement lookup route dispatch dry-run result redaction boundary item must remain closed: %#v", item)
		}
	}
}
