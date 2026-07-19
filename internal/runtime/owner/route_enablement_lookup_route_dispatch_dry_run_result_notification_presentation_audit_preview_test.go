package owner

import (
	"os"
	"path/filepath"
	"testing"
)

func TestProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationPresentationAuditPreviewModelsNotifications(t *testing.T) {
	preview, err := NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationPresentationAuditPreview(projectRootForOwnerTest(t))
	if err != nil {
		t.Fatalf("NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationPresentationAuditPreview returned error: %v", err)
	}

	if preview.Version != currentProjectVersion(t) ||
		preview.SchemaVersion != "xnix.runtime.production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_kde_safe_redacted_status_closed_record_writer_storage_root_authorization_receipt_dry_run_lookup_result_consumer_projection_route_enablement_lookup_route_dispatch_dry_run_result_notification_presentation_audit.v1" ||
		preview.RequestType != "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-presentation-audit-preview" ||
		preview.AuditType != "receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-presentation-audit" ||
		preview.AuditDecision != "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-presentation-audit-ready-notification-review-only-delivery-disabled" {
		t.Fatalf("unexpected notification presentation schema: %#v", preview)
	}
	if !preview.CurrentMainlineConsumed ||
		!preview.RouteEnablementLookupRouteDispatchDryRunResultPresentationConsentConsumed ||
		!preview.RouteEnablementLookupRouteDispatchDryRunResultPresentationConsentReady ||
		!preview.LookupRouteDispatchDryRunResultNotificationPresentationRequired ||
		!preview.LookupRouteDispatchDryRunResultNotificationPresentationModeled ||
		!preview.LookupRouteDispatchDryRunResultNotificationPresentationReady ||
		!preview.RouteEnablementLookupRouteDispatchDryRunResultNotificationPresentationReady ||
		!preview.ConsentModeledRedactedPresentationConsumed ||
		!preview.KDENotificationPresentationModeled ||
		!preview.NotificationCenterSurfaceModeled ||
		!preview.InstallFailureNotificationModeled ||
		!preview.RepairSuggestionNotificationModeled ||
		!preview.EnvironmentSwitchNotificationModeled ||
		!preview.ApprovalInfoNotificationModeled ||
		!preview.NotificationPresentationReviewOnly ||
		!preview.NotificationDeliveryDisabled ||
		!preview.NotificationActionDisabled ||
		!preview.ActionCardsRemainDisabled ||
		!preview.ExplicitUserConsentRequired ||
		preview.ExplicitUserConsentCollected ||
		!preview.KDESafeRedactedResultOnly ||
		!preview.RawResultHidden ||
		preview.ReceiptPresent ||
		preview.ReceiptAccepted ||
		preview.ReceiptConsumed ||
		preview.RouteEnablementAccepted ||
		preview.LookupRouteEnabled ||
		preview.LookupRouteDispatchAuthorized ||
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
		preview.NotificationDeliveryEnabled ||
		preview.NotificationSent ||
		preview.NotificationCenterEventTriggered ||
		preview.NotificationActionEnabled ||
		preview.ActionCardEnabled ||
		preview.CompatibilityCenterOpened ||
		preview.SupportBundleExported ||
		preview.SupportCaseCreated ||
		!preview.RuntimeOwned ||
		!preview.GoRuntimeBacked ||
		preview.KDEPolicyOwner ||
		preview.ProductionReadiness ||
		preview.ProductionOwnershipReady {
		t.Fatalf("unsafe notification presentation decision: %#v", preview)
	}
	if preview.NotificationPresentationItemCount != 4 ||
		preview.RequiredNotificationPresentationItemCount != 4 ||
		preview.ReadyNotificationPresentationItemCount != 4 ||
		preview.MissingNotificationPresentationItemCount != 0 ||
		preview.DeliveredNotificationItemCount != 0 ||
		preview.NotificationActionEnabledItemCount != 0 ||
		preview.ActionCardEnabledItemCount != 0 ||
		preview.SideEffectNotificationPresentationItemCount != 0 {
		t.Fatalf("unexpected notification presentation counts: %#v", preview)
	}
	if !sameStrings(preview.NotificationPresentationItemIDs, []string{
		"route-enablement-lookup-route-dispatch-dry-run-result-notification-presentation-install-failure",
		"route-enablement-lookup-route-dispatch-dry-run-result-notification-presentation-repair-suggestion",
		"route-enablement-lookup-route-dispatch-dry-run-result-notification-presentation-environment-switch",
		"route-enablement-lookup-route-dispatch-dry-run-result-notification-presentation-approval-info",
	}) {
		t.Fatalf("unexpected notification presentation item ids: %#v", preview.NotificationPresentationItemIDs)
	}
	for _, item := range preview.NotificationPresentationItems {
		if !item.EvidencePresent ||
			!item.CurrentMainlineConsumed ||
			!item.RouteEnablementLookupRouteDispatchDryRunResultPresentationConsentConsumed ||
			!item.RouteEnablementLookupRouteDispatchDryRunResultPresentationConsentReady ||
			!item.NotificationPresentationRequired ||
			!item.NotificationPresentationModeled ||
			!item.NotificationPresentationReady ||
			!item.ExplicitUserConsentRequired ||
			item.ExplicitUserConsentCollected ||
			!item.NotificationReviewOnly ||
			!item.KDESafeRedactedResultOnly ||
			!item.RawResultHidden ||
			item.NotificationDeliveryEnabled ||
			item.NotificationSent ||
			item.NotificationActionEnabled ||
			item.ActionCardEnabled ||
			!item.SideEffectsDisabled ||
			!item.RuntimeOwned ||
			!item.GoRuntimeBacked ||
			item.KDEPolicyOwner ||
			item.HostRootModified ||
			item.InternalDetailsExposed ||
			item.NotificationPresentationStatus != "redacted-status-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-presentation-ready-review-only-delivery-disabled" {
			t.Fatalf("unsafe notification presentation item: %#v", item)
		}
	}
	if preview.Counts.Total != 7 ||
		preview.Counts.Passed != 7 ||
		preview.Counts.Pending != 0 ||
		preview.Counts.Blocked != 0 ||
		!sameStrings(preview.CheckIDs, []string{"current-mainline-consumed", "route-enablement-lookup-route-dispatch-dry-run-result-presentation-consent-consumed", "lookup-route-dispatch-dry-run-result-notification-presentation-modeled", "four-kde-notification-events-review-only", "notification-delivery-actions-and-action-cards-disabled", "raw-result-exposure-and-persistence-disabled", "production-and-host-boundary-closed"}) {
		t.Fatalf("unexpected notification presentation checks: counts=%#v ids=%#v", preview.Counts, preview.CheckIDs)
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
		t.Fatalf("unsafe notification presentation gates: %#v", preview)
	}
	if err := validateNoBackendTerms(preview, "KDE-safe notification presentation test"); err != nil {
		t.Fatalf("validateNoBackendTerms returned error: %v", err)
	}
}

func TestProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationPresentationAuditPreviewFailsClosedWithoutSources(t *testing.T) {
	root, err := os.MkdirTemp("", "xnix-lookup-dispatch-dry-run-result-notification-presentation-*")
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
	preview, err := NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationPresentationAuditPreview(root)
	if err != nil {
		t.Fatalf("NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationPresentationAuditPreview returned error: %v", err)
	}
	if preview.CurrentMainlineConsumed ||
		preview.RouteEnablementLookupRouteDispatchDryRunResultPresentationConsentConsumed ||
		preview.RouteEnablementLookupRouteDispatchDryRunResultPresentationConsentReady ||
		!preview.LookupRouteDispatchDryRunResultNotificationPresentationRequired ||
		!preview.LookupRouteDispatchDryRunResultNotificationPresentationModeled ||
		preview.LookupRouteDispatchDryRunResultNotificationPresentationReady ||
		preview.RouteEnablementLookupRouteDispatchDryRunResultNotificationPresentationReady ||
		preview.ConsentModeledRedactedPresentationConsumed ||
		preview.KDENotificationPresentationModeled ||
		preview.NotificationCenterSurfaceModeled ||
		preview.InstallFailureNotificationModeled ||
		preview.RepairSuggestionNotificationModeled ||
		preview.EnvironmentSwitchNotificationModeled ||
		preview.ApprovalInfoNotificationModeled ||
		!preview.NotificationPresentationReviewOnly ||
		!preview.NotificationDeliveryDisabled ||
		!preview.NotificationActionDisabled ||
		!preview.ActionCardsRemainDisabled ||
		!preview.ExplicitUserConsentRequired ||
		preview.ExplicitUserConsentCollected ||
		preview.KDESafeRedactedResultOnly ||
		preview.RawResultHidden ||
		preview.NotificationPresentationItemCount != 4 ||
		preview.ReadyNotificationPresentationItemCount != 0 ||
		preview.MissingNotificationPresentationItemCount != 4 ||
		preview.Counts.Total != 7 ||
		preview.Counts.Passed != 2 ||
		preview.Counts.Blocked != 5 ||
		preview.AuditDecision != "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-presentation-audit-blocked" {
		t.Fatalf("missing notification presentation sources must fail closed: %#v", preview)
	}
	for _, item := range preview.NotificationPresentationItems {
		if item.EvidencePresent ||
			item.CurrentMainlineConsumed ||
			item.RouteEnablementLookupRouteDispatchDryRunResultPresentationConsentConsumed ||
			item.RouteEnablementLookupRouteDispatchDryRunResultPresentationConsentReady ||
			!item.NotificationPresentationRequired ||
			item.NotificationPresentationModeled ||
			item.NotificationPresentationReady ||
			!item.ExplicitUserConsentRequired ||
			item.ExplicitUserConsentCollected ||
			!item.NotificationReviewOnly ||
			item.KDESafeRedactedResultOnly ||
			item.RawResultHidden ||
			item.NotificationDeliveryEnabled ||
			item.NotificationSent ||
			item.NotificationActionEnabled ||
			item.ActionCardEnabled ||
			!item.SideEffectsDisabled ||
			item.HostRootModified ||
			item.NotificationPresentationStatus != "missing-redacted-status-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-presentation-evidence" {
			t.Fatalf("missing notification presentation item must remain closed: %#v", item)
		}
	}
}
