package owner

import (
	"os"
	"path/filepath"
	"testing"
)

func TestProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultRedactedPresentationEligibilityAuditPreviewModelsPresentationEligibility(t *testing.T) {
	preview, err := NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultRedactedPresentationEligibilityAuditPreview(projectRootForOwnerTest(t))
	if err != nil {
		t.Fatalf("NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultRedactedPresentationEligibilityAuditPreview returned error: %v", err)
	}

	if preview.Version != currentProjectVersion(t) ||
		preview.SchemaVersion != "xnix.runtime.production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_kde_safe_redacted_status_closed_record_writer_storage_root_authorization_receipt_dry_run_lookup_result_consumer_projection_route_enablement_lookup_route_dispatch_dry_run_result_redacted_presentation_eligibility_audit.v1" ||
		preview.RequestType != "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-redacted-presentation-eligibility-audit-preview" ||
		preview.AuditType != "receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-redacted-presentation-eligibility-audit" ||
		preview.AuditDecision != "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-redacted-presentation-eligibility-audit-ready-presentation-review-only-raw-result-hidden" {
		t.Fatalf("unexpected presentation eligibility schema: %#v", preview)
	}
	if !preview.CurrentMainlineConsumed ||
		!preview.RouteEnablementLookupRouteDispatchDryRunResultRedactionBoundaryConsumed ||
		!preview.RouteEnablementLookupRouteDispatchDryRunResultRedactionBoundaryReady ||
		!preview.LookupRouteDispatchDryRunResultRedactedPresentationEligibilityRequired ||
		!preview.LookupRouteDispatchDryRunResultRedactedPresentationEligibilityModeled ||
		!preview.LookupRouteDispatchDryRunResultRedactedPresentationEligibilityReady ||
		!preview.RouteEnablementLookupRouteDispatchDryRunResultRedactedPresentationReady ||
		!preview.KDEPresentationSurfaceEligibilityModeled ||
		!preview.CompatibilityCenterPresentationEligible ||
		!preview.NotificationCenterPresentationEligible ||
		!preview.SettingsPresentationEligible ||
		!preview.RuntimeDiagnosticsPresentationEligible ||
		!preview.KDESafeRedactedResultOnly ||
		!preview.RawResultHidden ||
		preview.ReceiptPresent ||
		preview.ReceiptAccepted ||
		preview.ReceiptConsumed ||
		preview.RouteEnablementAccepted ||
		preview.LookupRouteEnabled ||
		preview.LookupRouteDispatchAuthorized ||
		preview.LookupRouteDispatchDryRunResultRedactionPassed ||
		preview.LookupRouteDispatchCallable ||
		preview.StorageWriteEnabled ||
		preview.StatusPersistenceWriteEnabled ||
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
		t.Fatalf("unsafe presentation eligibility decision: %#v", preview)
	}
	if preview.PresentationEligibilityItemCount != 4 ||
		preview.RequiredPresentationEligibilityItemCount != 4 ||
		preview.ReadyPresentationEligibilityItemCount != 4 ||
		preview.MissingPresentationEligibilityItemCount != 0 ||
		preview.UserVisiblePresentationItemCount != 4 ||
		preview.PersistedPresentationEligibilityItemCount != 0 ||
		preview.RawExposedPresentationEligibilityItemCount != 0 ||
		preview.SideEffectPresentationEligibilityItemCount != 0 {
		t.Fatalf("unexpected presentation eligibility counts: %#v", preview)
	}
	if !sameStrings(preview.PresentationEligibilityItemIDs, []string{
		"route-enablement-lookup-route-dispatch-dry-run-result-redacted-presentation-eligibility-compatibility-center",
		"route-enablement-lookup-route-dispatch-dry-run-result-redacted-presentation-eligibility-notification-center",
		"route-enablement-lookup-route-dispatch-dry-run-result-redacted-presentation-eligibility-settings",
		"route-enablement-lookup-route-dispatch-dry-run-result-redacted-presentation-eligibility-runtime-diagnostics",
	}) {
		t.Fatalf("unexpected presentation eligibility item ids: %#v", preview.PresentationEligibilityItemIDs)
	}
	for _, item := range preview.PresentationEligibilityItems {
		if !item.EvidencePresent ||
			!item.CurrentMainlineConsumed ||
			!item.RouteEnablementLookupRouteDispatchDryRunResultRedactionBoundaryConsumed ||
			!item.RouteEnablementLookupRouteDispatchDryRunResultRedactionBoundaryReady ||
			!item.LookupRouteDispatchDryRunResultRedactedPresentationEligibilityModeled ||
			!item.KDESafeRedactedResultOnly ||
			!item.RawResultHidden ||
			!item.UserVisible ||
			!item.ReviewOnly ||
			!item.RuntimeOwned ||
			!item.GoRuntimeBacked ||
			item.KDEPolicyOwner ||
			!item.SideEffectsDisabled ||
			item.RedactedSummaryPersisted ||
			item.RawResultExposed ||
			item.LookupRouteEnabled ||
			item.LookupRouteDispatchCallable ||
			item.HostRootModified ||
			item.InternalDetailsExposed ||
			item.LookupRouteDispatchDryRunResultRedactedPresentationEligibilityStatus != "redacted-status-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-redacted-presentation-eligible-review-only-raw-result-hidden" {
			t.Fatalf("unsafe presentation eligibility item: %#v", item)
		}
	}
	if preview.Counts.Total != 7 ||
		preview.Counts.Passed != 7 ||
		preview.Counts.Pending != 0 ||
		preview.Counts.Blocked != 0 ||
		!sameStrings(preview.CheckIDs, []string{"current-mainline-consumed", "route-enablement-lookup-route-dispatch-dry-run-result-redaction-boundary-consumed", "lookup-route-dispatch-dry-run-result-redacted-presentation-eligibility-modeled", "four-kde-presentation-surfaces-eligible-review-only", "raw-result-exposure-and-persistence-disabled", "routes-dispatch-notifications-and-support-disabled", "production-and-host-boundary-closed"}) {
		t.Fatalf("unexpected presentation eligibility checks: counts=%#v ids=%#v", preview.Counts, preview.CheckIDs)
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
		t.Fatalf("unsafe presentation eligibility gates: %#v", preview)
	}
	if err := validateNoBackendTerms(preview, "KDE-safe redacted presentation eligibility test"); err != nil {
		t.Fatalf("validateNoBackendTerms returned error: %v", err)
	}
}

func TestProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultRedactedPresentationEligibilityAuditPreviewFailsClosedWithoutSources(t *testing.T) {
	root, err := os.MkdirTemp("", "xnix-lookup-dispatch-dry-run-result-redacted-presentation-eligibility-*")
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
	preview, err := NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultRedactedPresentationEligibilityAuditPreview(root)
	if err != nil {
		t.Fatalf("NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultRedactedPresentationEligibilityAuditPreview returned error: %v", err)
	}
	if preview.CurrentMainlineConsumed ||
		preview.RouteEnablementLookupRouteDispatchDryRunResultRedactionBoundaryConsumed ||
		preview.RouteEnablementLookupRouteDispatchDryRunResultRedactionBoundaryReady ||
		!preview.LookupRouteDispatchDryRunResultRedactedPresentationEligibilityRequired ||
		!preview.LookupRouteDispatchDryRunResultRedactedPresentationEligibilityModeled ||
		preview.LookupRouteDispatchDryRunResultRedactedPresentationEligibilityReady ||
		preview.RouteEnablementLookupRouteDispatchDryRunResultRedactedPresentationReady ||
		preview.KDEPresentationSurfaceEligibilityModeled ||
		preview.CompatibilityCenterPresentationEligible ||
		preview.NotificationCenterPresentationEligible ||
		preview.SettingsPresentationEligible ||
		preview.RuntimeDiagnosticsPresentationEligible ||
		preview.KDESafeRedactedResultOnly ||
		preview.RawResultHidden ||
		preview.RawResultExposed ||
		preview.PresentationEligibilityItemCount != 4 ||
		preview.ReadyPresentationEligibilityItemCount != 0 ||
		preview.MissingPresentationEligibilityItemCount != 4 ||
		preview.UserVisiblePresentationItemCount != 0 ||
		preview.Counts.Total != 7 ||
		preview.Counts.Passed != 2 ||
		preview.Counts.Blocked != 5 ||
		preview.AuditDecision != "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-redacted-presentation-eligibility-audit-blocked" {
		t.Fatalf("missing presentation eligibility sources must fail closed: %#v", preview)
	}
	for _, item := range preview.PresentationEligibilityItems {
		if item.EvidencePresent ||
			item.CurrentMainlineConsumed ||
			item.RouteEnablementLookupRouteDispatchDryRunResultRedactionBoundaryConsumed ||
			item.RouteEnablementLookupRouteDispatchDryRunResultRedactionBoundaryReady ||
			item.LookupRouteDispatchDryRunResultRedactedPresentationEligibilityModeled ||
			item.KDESafeRedactedResultOnly ||
			item.RawResultHidden ||
			item.UserVisible ||
			item.RawResultExposed ||
			item.LookupRouteEnabled ||
			item.LookupRouteDispatchCallable ||
			item.HostRootModified ||
			item.LookupRouteDispatchDryRunResultRedactedPresentationEligibilityStatus != "missing-redacted-status-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-redacted-presentation-eligibility-evidence" {
			t.Fatalf("missing presentation eligibility item must remain closed: %#v", item)
		}
	}
}
