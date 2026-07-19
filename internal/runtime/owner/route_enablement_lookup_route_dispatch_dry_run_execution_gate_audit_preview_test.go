package owner

import (
	"os"
	"path/filepath"
	"testing"
)

func TestProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunExecutionGateAuditPreviewModelsDispatchDryRunExecutionGate(t *testing.T) {
	preview, err := NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunExecutionGateAuditPreview(projectRootForOwnerTest(t))
	if err != nil {
		t.Fatalf("NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunExecutionGateAuditPreview returned error: %v", err)
	}

	if preview.Version != currentProjectVersion(t) ||
		preview.SchemaVersion != "xnix.runtime.production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_kde_safe_redacted_status_closed_record_writer_storage_root_authorization_receipt_dry_run_lookup_result_consumer_projection_route_enablement_lookup_route_dispatch_dry_run_execution_gate_audit.v1" ||
		preview.RequestType != "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-execution-gate-audit-preview" ||
		preview.AuditType != "receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-execution-gate-audit" ||
		preview.AuditDecision != "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-execution-gate-audit-ready-route-disabled-dispatch-disabled" {
		t.Fatalf("unexpected KDE-safe redacted status dry-run lookup result consumer projection route enablement lookup route dispatch dry-run execution gate schema: %#v", preview)
	}
	if !preview.CurrentMainlineConsumed ||
		!preview.RouteEnablementLookupRouteDispatchAuthorizationConsumed ||
		!preview.RouteEnablementLookupRouteDispatchAuthorizationReady ||
		!preview.LookupRouteDispatchDryRunExecutionGateRequired ||
		!preview.LookupRouteDispatchDryRunExecutionGateModeled ||
		!preview.LookupRouteDispatchDryRunExecutionGateReady ||
		!preview.RouteEnablementLookupRouteDispatchDryRunExecutionGateBoundaryReady ||
		!preview.OpaqueRouteEnablementReceipt ||
		!preview.KDESafeRedactedStatusOnly ||
		!preview.CompatibilityCenterLookupRouteDispatchDryRunExecutionGateModeled ||
		!preview.RuntimeDiagnosticsLookupRouteDispatchDryRunExecutionGateModeled ||
		preview.ReceiptPresent ||
		preview.ReceiptAccepted ||
		preview.ReceiptConsumed ||
		preview.AcceptanceAuthorized ||
		preview.RouteEnablementAccepted ||
		preview.LookupRouteEnablementGranted ||
		preview.LookupRouteGrantAuthorized ||
		preview.LookupRouteEnabled ||
		preview.LookupRouteDispatchAuthorized ||
		preview.LookupRouteDispatchDryRunExecutionGatePassed ||
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
		t.Fatalf("unsafe KDE-safe redacted status dry-run lookup result consumer projection route enablement lookup route dispatch dry-run execution gate decision: %#v", preview)
	}
	if preview.DispatchDryRunExecutionGateItemCount != 10 ||
		preview.RequiredDispatchDryRunExecutionGateItemCount != 10 ||
		preview.ReadyDispatchDryRunExecutionGateItemCount != 10 ||
		preview.MissingDispatchDryRunExecutionGateItemCount != 0 ||
		preview.PassedDispatchDryRunExecutionGateItemCount != 0 ||
		preview.EnabledRouteItemCount != 0 ||
		preview.PersistedDispatchDryRunExecutionGateItemCount != 0 ||
		preview.RawExposedDispatchDryRunExecutionGateItemCount != 0 ||
		preview.SideEffectDispatchDryRunExecutionGateItemCount != 0 ||
		preview.CompatibilityCenterDispatchDryRunExecutionGateItemCount != 5 ||
		preview.RuntimeDiagnosticsDispatchDryRunExecutionGateItemCount != 5 {
		t.Fatalf("unexpected KDE-safe redacted status dry-run lookup result consumer projection route enablement lookup route dispatch dry-run execution gate counts: %#v", preview)
	}
	if !sameStrings(preview.DispatchDryRunExecutionGateItemIDs, []string{
		"review-route-enablement-lookup-route-dispatch-dry-run-execution-gate-compatibility-center",
		"review-route-enablement-lookup-route-dispatch-dry-run-execution-gate-runtime-diagnostics",
		"renew-route-enablement-lookup-route-dispatch-dry-run-execution-gate-compatibility-center",
		"renew-route-enablement-lookup-route-dispatch-dry-run-execution-gate-runtime-diagnostics",
		"open-compatibility-center-route-enablement-lookup-route-dispatch-dry-run-execution-gate-compatibility-center",
		"open-compatibility-center-route-enablement-lookup-route-dispatch-dry-run-execution-gate-runtime-diagnostics",
		"dismiss-route-enablement-lookup-route-dispatch-dry-run-execution-gate-compatibility-center",
		"dismiss-route-enablement-lookup-route-dispatch-dry-run-execution-gate-runtime-diagnostics",
		"support-info-route-enablement-lookup-route-dispatch-dry-run-execution-gate-compatibility-center",
		"support-info-route-enablement-lookup-route-dispatch-dry-run-execution-gate-runtime-diagnostics",
	}) {
		t.Fatalf("unexpected KDE-safe redacted status dry-run lookup result consumer projection route enablement lookup route dispatch dry-run execution gate items: %#v", preview.DispatchDryRunExecutionGateItemIDs)
	}
	for _, item := range preview.DispatchDryRunExecutionGateItems {
		if !item.EvidencePresent ||
			!item.CurrentMainlineConsumed ||
			!item.RouteEnablementLookupRouteDispatchAuthorizationConsumed ||
			!item.RouteEnablementLookupRouteDispatchAuthorizationReady ||
			!item.LookupRouteDispatchDryRunExecutionGateModeled ||
			!item.RouteEnablementLookupRouteDispatchDryRunExecutionGateBoundaryReady ||
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
			item.LookupRouteDispatchDryRunExecutionGatePassed ||
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
			item.LookupRouteDispatchDryRunExecutionGateStatus != "redacted-status-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-execution-gate-modeled-dispatch-disabled-route-disabled" {
			t.Fatalf("unsafe KDE-safe redacted status dry-run lookup result consumer projection route enablement lookup route dispatch dry-run execution gate item: %#v", item)
		}
	}
	if preview.Counts.Total != 8 ||
		preview.Counts.Passed != 8 ||
		preview.Counts.Pending != 0 ||
		preview.Counts.Blocked != 0 ||
		!sameStrings(preview.CheckIDs, []string{"current-mainline-consumed", "route-enablement-lookup-route-dispatch-authorization-consumed", "lookup-route-dispatch-dry-run-execution-gate-modeled", "compatibility-center-and-runtime-lookup-route-dispatch-dry-run-execution-gates-modeled", "ten-dispatch-dry-run-execution-gate-items-ready-dispatch-disabled", "lookup-route-dispatch-dry-run-execution-gate-and-persistence-disabled", "consumer-routes-dispatch-and-support-disabled", "production-and-host-boundary-closed"}) {
		t.Fatalf("unexpected KDE-safe redacted status dry-run lookup result consumer projection route enablement lookup route dispatch dry-run execution gate checks: counts=%#v ids=%#v", preview.Counts, preview.CheckIDs)
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
		t.Fatalf("unsafe KDE-safe redacted status dry-run lookup result consumer projection route enablement lookup route dispatch dry-run execution gate gates: %#v", preview)
	}
	if err := validateNoBackendTerms(preview, "KDE-safe redacted status dry-run lookup result consumer projection route enablement lookup route dispatch dry-run execution gate test"); err != nil {
		t.Fatalf("validateNoBackendTerms returned error: %v", err)
	}
}

func TestProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunExecutionGateAuditPreviewFailsClosedWithoutSources(t *testing.T) {
	root, err := os.MkdirTemp("", "xnix-lookup-dispatch-dry-run-execution-gate-*")
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
	preview, err := NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunExecutionGateAuditPreview(root)
	if err != nil {
		t.Fatalf("NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupRouteDispatchDryRunExecutionGateAuditPreview returned error: %v", err)
	}
	if preview.CurrentMainlineConsumed ||
		preview.RouteEnablementLookupRouteDispatchAuthorizationConsumed ||
		preview.RouteEnablementLookupRouteDispatchAuthorizationReady ||
		!preview.LookupRouteDispatchDryRunExecutionGateRequired ||
		!preview.LookupRouteDispatchDryRunExecutionGateModeled ||
		preview.LookupRouteDispatchDryRunExecutionGateReady ||
		preview.RouteEnablementLookupRouteDispatchDryRunExecutionGateBoundaryReady ||
		preview.OpaqueRouteEnablementReceipt ||
		preview.KDESafeRedactedStatusOnly ||
		preview.CompatibilityCenterLookupRouteDispatchDryRunExecutionGateModeled ||
		preview.RuntimeDiagnosticsLookupRouteDispatchDryRunExecutionGateModeled ||
		preview.ReceiptPresent ||
		preview.ReceiptAccepted ||
		preview.ReceiptConsumed ||
		preview.AcceptanceAuthorized ||
		preview.RouteEnablementAccepted ||
		preview.LookupRouteEnablementGranted ||
		preview.LookupRouteGrantAuthorized ||
		preview.LookupRouteEnabled ||
		preview.LookupRouteDispatchAuthorized ||
		preview.LookupRouteDispatchDryRunExecutionGatePassed ||
		preview.LookupRouteDispatchCallable ||
		preview.StorageWriteEnabled ||
		preview.RawResultExposed ||
		preview.DispatchDryRunExecutionGateItemCount != 10 ||
		preview.RequiredDispatchDryRunExecutionGateItemCount != 10 ||
		preview.ReadyDispatchDryRunExecutionGateItemCount != 0 ||
		preview.MissingDispatchDryRunExecutionGateItemCount != 10 ||
		preview.Counts.Total != 8 ||
		preview.Counts.Passed != 4 ||
		preview.Counts.Blocked != 4 ||
		preview.AuditDecision != "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-execution-gate-audit-blocked" {
		t.Fatalf("missing KDE-safe redacted status dry-run lookup result consumer projection route enablement lookup route dispatch dry-run execution gate sources must fail closed: %#v", preview)
	}
	for _, item := range preview.DispatchDryRunExecutionGateItems {
		if item.EvidencePresent ||
			item.CurrentMainlineConsumed ||
			item.RouteEnablementLookupRouteDispatchAuthorizationConsumed ||
			item.RouteEnablementLookupRouteDispatchAuthorizationReady ||
			item.LookupRouteDispatchDryRunExecutionGateModeled ||
			item.RouteEnablementLookupRouteDispatchDryRunExecutionGateBoundaryReady ||
			item.OpaqueRouteEnablementReceipt ||
			item.ReceiptAccepted ||
			item.ReceiptConsumed ||
			item.AcceptanceAuthorized ||
			item.RouteEnablementAccepted ||
			item.LookupRouteEnablementGranted ||
			item.LookupRouteGrantAuthorized ||
			item.LookupRouteEnabled ||
			item.LookupRouteDispatchAuthorized ||
			item.LookupRouteDispatchDryRunExecutionGatePassed ||
			item.LookupRouteDispatchCallable ||
			item.StorageWriteEnabled ||
			item.RawResultExposed ||
			item.UserVisible ||
			item.LookupRouteDispatchDryRunExecutionGateStatus != "missing-redacted-status-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-execution-gate-evidence" {
			t.Fatalf("missing KDE-safe redacted status dry-run lookup result consumer projection route enablement lookup route dispatch dry-run execution gate item must remain closed: %#v", item)
		}
	}
}
