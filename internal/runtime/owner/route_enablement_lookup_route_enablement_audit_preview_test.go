package owner

import (
	"os"
	"path/filepath"
	"testing"
)

func TestProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteEnablementAuditPreviewModelsEnablement(t *testing.T) {
	preview, err := NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteEnablementAuditPreview(projectRootForOwnerTest(t))
	if err != nil {
		t.Fatalf("NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteEnablementAuditPreview returned error: %v", err)
	}

	if preview.Version != currentProjectVersion(t) ||
		preview.SchemaVersion != "xnix.runtime.production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_kde_safe_redacted_status_closed_record_writer_storage_root_authorization_receipt_dry_run_lookup_result_consumer_projection_route_enablement_lookup_route_enablement_audit.v1" ||
		preview.RequestType != "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-enablement-audit-preview" ||
		preview.AuditType != "receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-enablement-audit" ||
		preview.AuditDecision != "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-enablement-audit-ready-route-disabled-dispatch-disabled" {
		t.Fatalf("unexpected KDE-safe redacted status dry-run lookup result consumer projection route enablement lookup route enablement schema: %#v", preview)
	}
	if !preview.CurrentMainlineConsumed ||
		!preview.RouteEnablementLookupRouteGrantConsumed ||
		!preview.RouteEnablementLookupRouteGrantReady ||
		!preview.LookupRouteEnablementRequired ||
		!preview.LookupRouteEnablementModeled ||
		!preview.LookupRouteEnablementReady ||
		!preview.RouteEnablementLookupRouteEnablementBoundaryReady ||
		!preview.OpaqueRouteEnablementReceipt ||
		!preview.KDESafeRedactedStatusOnly ||
		!preview.CompatibilityCenterLookupRouteEnablementModeled ||
		!preview.RuntimeDiagnosticsLookupRouteEnablementModeled ||
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
		t.Fatalf("unsafe KDE-safe redacted status dry-run lookup result consumer projection route enablement lookup route enablement decision: %#v", preview)
	}
	if preview.EnablementItemCount != 10 ||
		preview.RequiredEnablementItemCount != 10 ||
		preview.ReadyEnablementItemCount != 10 ||
		preview.MissingEnablementItemCount != 0 ||
		preview.GrantedRouteItemCount != 0 ||
		preview.EnabledRouteItemCount != 0 ||
		preview.PersistedEnablementItemCount != 0 ||
		preview.RawExposedEnablementItemCount != 0 ||
		preview.SideEffectEnablementItemCount != 0 ||
		preview.CompatibilityCenterEnablementItemCount != 5 ||
		preview.RuntimeDiagnosticsEnablementItemCount != 5 {
		t.Fatalf("unexpected KDE-safe redacted status dry-run lookup result consumer projection route enablement lookup route enablement counts: %#v", preview)
	}
	if !sameStrings(preview.EnablementItemIDs, []string{
		"review-route-enablement-lookup-route-enablement-compatibility-center",
		"review-route-enablement-lookup-route-enablement-runtime-diagnostics",
		"renew-route-enablement-lookup-route-enablement-compatibility-center",
		"renew-route-enablement-lookup-route-enablement-runtime-diagnostics",
		"open-compatibility-center-route-enablement-lookup-route-enablement-compatibility-center",
		"open-compatibility-center-route-enablement-lookup-route-enablement-runtime-diagnostics",
		"dismiss-route-enablement-lookup-route-enablement-compatibility-center",
		"dismiss-route-enablement-lookup-route-enablement-runtime-diagnostics",
		"support-info-route-enablement-lookup-route-enablement-compatibility-center",
		"support-info-route-enablement-lookup-route-enablement-runtime-diagnostics",
	}) {
		t.Fatalf("unexpected KDE-safe redacted status dry-run lookup result consumer projection route enablement lookup route enablement items: %#v", preview.EnablementItemIDs)
	}
	for _, item := range preview.EnablementItems {
		if !item.EvidencePresent ||
			!item.CurrentMainlineConsumed ||
			!item.RouteEnablementLookupRouteGrantConsumed ||
			!item.RouteEnablementLookupRouteGrantReady ||
			!item.LookupRouteEnablementModeled ||
			!item.RouteEnablementLookupRouteEnablementBoundaryReady ||
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
			item.LookupRouteEnablementStatus != "redacted-status-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-enablement-modeled-route-disabled-dispatch-disabled" {
			t.Fatalf("unsafe KDE-safe redacted status dry-run lookup result consumer projection route enablement lookup route enablement item: %#v", item)
		}
	}
	if preview.Counts.Total != 8 ||
		preview.Counts.Passed != 8 ||
		preview.Counts.Pending != 0 ||
		preview.Counts.Blocked != 0 ||
		!sameStrings(preview.CheckIDs, []string{"current-mainline-consumed", "route-enablement-lookup-route-grant-consumed", "lookup-route-enablement-modeled", "compatibility-center-and-runtime-lookup-route-enablements-modeled", "ten-enablement-items-ready-route-disabled", "lookup-route-enablement-and-persistence-disabled", "consumer-routes-dispatch-and-support-disabled", "production-and-host-boundary-closed"}) {
		t.Fatalf("unexpected KDE-safe redacted status dry-run lookup result consumer projection route enablement lookup route enablement checks: counts=%#v ids=%#v", preview.Counts, preview.CheckIDs)
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
		t.Fatalf("unsafe KDE-safe redacted status dry-run lookup result consumer projection route enablement lookup route enablement gates: %#v", preview)
	}
	if err := validateNoBackendTerms(preview, "KDE-safe redacted status dry-run lookup result consumer projection route enablement lookup route enablement test"); err != nil {
		t.Fatalf("validateNoBackendTerms returned error: %v", err)
	}
}

func TestProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteEnablementAuditPreviewFailsClosedWithoutSources(t *testing.T) {
	root, err := os.MkdirTemp("", "xnix-lookup-enablement-*")
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
	preview, err := NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteEnablementAuditPreview(root)
	if err != nil {
		t.Fatalf("NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteEnablementAuditPreview returned error: %v", err)
	}
	if preview.CurrentMainlineConsumed ||
		preview.RouteEnablementLookupRouteGrantConsumed ||
		preview.RouteEnablementLookupRouteGrantReady ||
		!preview.LookupRouteEnablementRequired ||
		!preview.LookupRouteEnablementModeled ||
		preview.LookupRouteEnablementReady ||
		preview.RouteEnablementLookupRouteEnablementBoundaryReady ||
		preview.OpaqueRouteEnablementReceipt ||
		preview.KDESafeRedactedStatusOnly ||
		preview.CompatibilityCenterLookupRouteEnablementModeled ||
		preview.RuntimeDiagnosticsLookupRouteEnablementModeled ||
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
		preview.EnablementItemCount != 10 ||
		preview.RequiredEnablementItemCount != 10 ||
		preview.ReadyEnablementItemCount != 0 ||
		preview.MissingEnablementItemCount != 10 ||
		preview.Counts.Total != 8 ||
		preview.Counts.Passed != 4 ||
		preview.Counts.Blocked != 4 ||
		preview.AuditDecision != "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-enablement-audit-blocked" {
		t.Fatalf("missing KDE-safe redacted status dry-run lookup result consumer projection route enablement lookup route enablement sources must fail closed: %#v", preview)
	}
	for _, item := range preview.EnablementItems {
		if item.EvidencePresent ||
			item.CurrentMainlineConsumed ||
			item.RouteEnablementLookupRouteGrantConsumed ||
			item.RouteEnablementLookupRouteGrantReady ||
			item.LookupRouteEnablementModeled ||
			item.RouteEnablementLookupRouteEnablementBoundaryReady ||
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
			item.LookupRouteEnablementStatus != "missing-redacted-status-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-enablement-evidence" {
			t.Fatalf("missing KDE-safe redacted status dry-run lookup result consumer projection route enablement lookup route enablement item must remain closed: %#v", item)
		}
	}
}
