package owner

import (
	"os"
	"path/filepath"
	"testing"
)

func TestProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteGrantAuditPreviewModelsGrant(t *testing.T) {
	preview, err := NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteGrantAuditPreview(projectRootForOwnerTest(t))
	if err != nil {
		t.Fatalf("NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteGrantAuditPreview returned error: %v", err)
	}

	if preview.Version != currentProjectVersion(t) ||
		preview.SchemaVersion != "xnix.runtime.production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_kde_safe_redacted_status_closed_record_writer_storage_root_authorization_receipt_dry_run_lookup_result_consumer_projection_route_enablement_lookup_route_grant_audit.v1" ||
		preview.RequestType != "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-grant-audit-preview" ||
		preview.AuditType != "receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-grant-audit" ||
		preview.AuditDecision != "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-grant-audit-ready-grant-disabled-routes-disabled" {
		t.Fatalf("unexpected KDE-safe redacted status dry-run lookup result consumer projection route enablement lookup route grant schema: %#v", preview)
	}
	if !preview.CurrentMainlineConsumed ||
		!preview.RouteEnablementAcceptedReceiptGateConsumed ||
		!preview.RouteEnablementAcceptedReceiptGateReady ||
		!preview.LookupRouteGrantRequired ||
		!preview.LookupRouteGrantModeled ||
		!preview.LookupRouteGrantReady ||
		!preview.RouteEnablementLookupRouteGrantBoundaryReady ||
		!preview.OpaqueRouteEnablementReceipt ||
		!preview.KDESafeRedactedStatusOnly ||
		!preview.CompatibilityCenterLookupRouteGrantModeled ||
		!preview.RuntimeDiagnosticsLookupRouteGrantModeled ||
		preview.ReceiptPresent ||
		preview.ReceiptAccepted ||
		preview.ReceiptConsumed ||
		preview.AcceptanceAuthorized ||
		preview.RouteEnablementAccepted ||
		preview.LookupRouteEnablementGranted ||
		preview.LookupRouteGrantAuthorized ||
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
		t.Fatalf("unsafe KDE-safe redacted status dry-run lookup result consumer projection route enablement lookup route grant decision: %#v", preview)
	}
	if preview.GrantItemCount != 10 ||
		preview.RequiredGrantItemCount != 10 ||
		preview.ReadyGrantItemCount != 10 ||
		preview.MissingGrantItemCount != 0 ||
		preview.GrantedRouteItemCount != 0 ||
		preview.EnabledRouteItemCount != 0 ||
		preview.PersistedGrantItemCount != 0 ||
		preview.RawExposedGrantItemCount != 0 ||
		preview.SideEffectGrantItemCount != 0 ||
		preview.CompatibilityCenterGrantItemCount != 5 ||
		preview.RuntimeDiagnosticsGrantItemCount != 5 {
		t.Fatalf("unexpected KDE-safe redacted status dry-run lookup result consumer projection route enablement lookup route grant counts: %#v", preview)
	}
	if !sameStrings(preview.GrantItemIDs, []string{
		"review-route-enablement-lookup-route-grant-compatibility-center",
		"review-route-enablement-lookup-route-grant-runtime-diagnostics",
		"renew-route-enablement-lookup-route-grant-compatibility-center",
		"renew-route-enablement-lookup-route-grant-runtime-diagnostics",
		"open-compatibility-center-route-enablement-lookup-route-grant-compatibility-center",
		"open-compatibility-center-route-enablement-lookup-route-grant-runtime-diagnostics",
		"dismiss-route-enablement-lookup-route-grant-compatibility-center",
		"dismiss-route-enablement-lookup-route-grant-runtime-diagnostics",
		"support-info-route-enablement-lookup-route-grant-compatibility-center",
		"support-info-route-enablement-lookup-route-grant-runtime-diagnostics",
	}) {
		t.Fatalf("unexpected KDE-safe redacted status dry-run lookup result consumer projection route enablement lookup route grant items: %#v", preview.GrantItemIDs)
	}
	for _, item := range preview.GrantItems {
		if !item.EvidencePresent ||
			!item.CurrentMainlineConsumed ||
			!item.RouteEnablementAcceptedReceiptGateConsumed ||
			!item.RouteEnablementAcceptedReceiptGateReady ||
			!item.LookupRouteGrantModeled ||
			!item.RouteEnablementLookupRouteGrantBoundaryReady ||
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
			item.LookupRouteGrantStatus != "redacted-status-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-grant-modeled-grant-disabled-routes-disabled" {
			t.Fatalf("unsafe KDE-safe redacted status dry-run lookup result consumer projection route enablement lookup route grant item: %#v", item)
		}
	}
	if preview.Counts.Total != 8 ||
		preview.Counts.Passed != 8 ||
		preview.Counts.Pending != 0 ||
		preview.Counts.Blocked != 0 ||
		!sameStrings(preview.CheckIDs, []string{"current-mainline-consumed", "route-enablement-accepted-receipt-gate-consumed", "lookup-route-grant-modeled", "compatibility-center-and-runtime-lookup-route-grants-modeled", "ten-grant-items-ready-grant-disabled", "lookup-route-grants-and-persistence-disabled", "consumer-routes-dispatch-and-support-disabled", "production-and-host-boundary-closed"}) {
		t.Fatalf("unexpected KDE-safe redacted status dry-run lookup result consumer projection route enablement lookup route grant checks: counts=%#v ids=%#v", preview.Counts, preview.CheckIDs)
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
		t.Fatalf("unsafe KDE-safe redacted status dry-run lookup result consumer projection route enablement lookup route grant gates: %#v", preview)
	}
	if err := validateNoBackendTerms(preview, "KDE-safe redacted status dry-run lookup result consumer projection route enablement lookup route grant test"); err != nil {
		t.Fatalf("validateNoBackendTerms returned error: %v", err)
	}
}

func TestProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteGrantAuditPreviewFailsClosedWithoutSources(t *testing.T) {
	root, err := os.MkdirTemp("", "xnix-lookup-grant-*")
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
	preview, err := NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteGrantAuditPreview(root)
	if err != nil {
		t.Fatalf("NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteGrantAuditPreview returned error: %v", err)
	}
	if preview.CurrentMainlineConsumed ||
		preview.RouteEnablementAcceptedReceiptGateConsumed ||
		preview.RouteEnablementAcceptedReceiptGateReady ||
		!preview.LookupRouteGrantRequired ||
		!preview.LookupRouteGrantModeled ||
		preview.LookupRouteGrantReady ||
		preview.RouteEnablementLookupRouteGrantBoundaryReady ||
		preview.OpaqueRouteEnablementReceipt ||
		preview.KDESafeRedactedStatusOnly ||
		preview.CompatibilityCenterLookupRouteGrantModeled ||
		preview.RuntimeDiagnosticsLookupRouteGrantModeled ||
		preview.ReceiptPresent ||
		preview.ReceiptAccepted ||
		preview.ReceiptConsumed ||
		preview.AcceptanceAuthorized ||
		preview.RouteEnablementAccepted ||
		preview.LookupRouteEnablementGranted ||
		preview.LookupRouteGrantAuthorized ||
		preview.LookupRouteEnabled ||
		preview.StorageWriteEnabled ||
		preview.RawResultExposed ||
		preview.GrantItemCount != 10 ||
		preview.RequiredGrantItemCount != 10 ||
		preview.ReadyGrantItemCount != 0 ||
		preview.MissingGrantItemCount != 10 ||
		preview.Counts.Total != 8 ||
		preview.Counts.Passed != 4 ||
		preview.Counts.Blocked != 4 ||
		preview.AuditDecision != "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-grant-audit-blocked" {
		t.Fatalf("missing KDE-safe redacted status dry-run lookup result consumer projection route enablement lookup route grant sources must fail closed: %#v", preview)
	}
	for _, item := range preview.GrantItems {
		if item.EvidencePresent ||
			item.CurrentMainlineConsumed ||
			item.RouteEnablementAcceptedReceiptGateConsumed ||
			item.RouteEnablementAcceptedReceiptGateReady ||
			item.LookupRouteGrantModeled ||
			item.RouteEnablementLookupRouteGrantBoundaryReady ||
			item.OpaqueRouteEnablementReceipt ||
			item.ReceiptAccepted ||
			item.ReceiptConsumed ||
			item.AcceptanceAuthorized ||
			item.RouteEnablementAccepted ||
			item.LookupRouteEnablementGranted ||
			item.LookupRouteGrantAuthorized ||
			item.LookupRouteEnabled ||
			item.StorageWriteEnabled ||
			item.RawResultExposed ||
			item.UserVisible ||
			item.LookupRouteGrantStatus != "missing-redacted-status-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-grant-evidence" {
			t.Fatalf("missing KDE-safe redacted status dry-run lookup result consumer projection route enablement lookup route grant item must remain closed: %#v", item)
		}
	}
}
