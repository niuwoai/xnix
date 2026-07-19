package owner

import (
	"os"
	"path/filepath"
	"testing"
)

func TestProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationDeliveryAuthorizationAuditPreviewModelsAuthorization(t *testing.T) {
	preview, err := NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationDeliveryAuthorizationAuditPreview(projectRootForOwnerTest(t))
	if err != nil {
		t.Fatalf("NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationDeliveryAuthorizationAuditPreview returned error: %v", err)
	}

	if preview.Version != currentProjectVersion(t) ||
		preview.SchemaVersion != "xnix.runtime.production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_kde_safe_redacted_status_closed_record_writer_storage_root_authorization_receipt_dry_run_lookup_result_consumer_projection_route_enablement_lookup_route_dispatch_dry_run_result_notification_delivery_authorization_audit.v1" ||
		preview.RequestType != "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-delivery-authorization-audit-preview" ||
		preview.AuditType != "receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-delivery-authorization-audit" ||
		preview.AuditDecision != "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-delivery-authorization-audit-ready-delivery-grant-disabled" {
		t.Fatalf("unexpected notification delivery authorization schema: %#v", preview)
	}
	if !preview.CurrentMainlineConsumed ||
		!preview.RouteEnablementLookupRouteDispatchDryRunResultNotificationPresentationConsumed ||
		!preview.RouteEnablementLookupRouteDispatchDryRunResultNotificationPresentationReady ||
		!preview.LookupRouteDispatchDryRunResultNotificationDeliveryAuthorizationRequired ||
		!preview.LookupRouteDispatchDryRunResultNotificationDeliveryAuthorizationModeled ||
		!preview.LookupRouteDispatchDryRunResultNotificationDeliveryAuthorizationReady ||
		!preview.RouteEnablementLookupRouteDispatchDryRunResultNotificationDeliveryReady ||
		!preview.NotificationPresentationReviewOnlyConsumed ||
		!preview.KDENotificationDeliveryAuthorizationModeled ||
		!preview.NotificationCenterDeliveryAuthorizationModeled ||
		!preview.InstallFailureDeliveryAuthorizationModeled ||
		!preview.RepairSuggestionDeliveryAuthorizationModeled ||
		!preview.EnvironmentSwitchDeliveryAuthorizationModeled ||
		!preview.ApprovalInfoDeliveryAuthorizationModeled ||
		!preview.NotificationDeliveryAuthorizationReviewOnly ||
		preview.NotificationDeliveryAuthorizationGranted ||
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
		t.Fatalf("unsafe notification delivery authorization decision: %#v", preview)
	}
	if preview.NotificationDeliveryAuthorizationItemCount != 4 ||
		preview.RequiredNotificationDeliveryAuthorizationItemCount != 4 ||
		preview.ReadyNotificationDeliveryAuthorizationItemCount != 4 ||
		preview.MissingNotificationDeliveryAuthorizationItemCount != 0 ||
		preview.AuthorizedNotificationDeliveryItemCount != 0 ||
		preview.DeliveredNotificationItemCount != 0 ||
		preview.NotificationActionEnabledItemCount != 0 ||
		preview.ActionCardEnabledItemCount != 0 ||
		preview.SideEffectNotificationDeliveryAuthorizationItemCount != 0 {
		t.Fatalf("unexpected notification delivery authorization counts: %#v", preview)
	}
	if !sameStrings(preview.NotificationDeliveryAuthorizationItemIDs, []string{
		"route-enablement-lookup-route-dispatch-dry-run-result-notification-delivery-authorization-install-failure",
		"route-enablement-lookup-route-dispatch-dry-run-result-notification-delivery-authorization-repair-suggestion",
		"route-enablement-lookup-route-dispatch-dry-run-result-notification-delivery-authorization-environment-switch",
		"route-enablement-lookup-route-dispatch-dry-run-result-notification-delivery-authorization-approval-info",
	}) {
		t.Fatalf("unexpected notification delivery authorization item ids: %#v", preview.NotificationDeliveryAuthorizationItemIDs)
	}
	for _, item := range preview.NotificationDeliveryAuthorizationItems {
		if !item.EvidencePresent ||
			!item.CurrentMainlineConsumed ||
			!item.NotificationPresentationConsumed ||
			!item.NotificationPresentationReady ||
			!item.DeliveryAuthorizationRequired ||
			!item.DeliveryAuthorizationModeled ||
			!item.DeliveryAuthorizationReady ||
			item.DeliveryAuthorizationGranted ||
			!item.ExplicitUserConsentRequired ||
			item.ExplicitUserConsentCollected ||
			!item.DeliveryReviewOnly ||
			!item.KDESafeRedactedResultOnly ||
			!item.RawResultHidden ||
			item.NotificationDeliveryEnabled ||
			item.NotificationSent ||
			item.NotificationCenterEventTriggered ||
			item.NotificationActionEnabled ||
			item.ActionCardEnabled ||
			!item.SideEffectsDisabled ||
			!item.RuntimeOwned ||
			!item.GoRuntimeBacked ||
			item.KDEPolicyOwner ||
			item.HostRootModified ||
			item.InternalDetailsExposed ||
			item.DeliveryAuthorizationStatus != "redacted-status-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-delivery-authorization-ready-delivery-grant-disabled" {
			t.Fatalf("unsafe notification delivery authorization item: %#v", item)
		}
	}
	if preview.Counts.Total != 7 ||
		preview.Counts.Passed != 7 ||
		preview.Counts.Pending != 0 ||
		preview.Counts.Blocked != 0 ||
		!sameStrings(preview.CheckIDs, []string{"current-mainline-consumed", "route-enablement-lookup-route-dispatch-dry-run-result-notification-presentation-consumed", "lookup-route-dispatch-dry-run-result-notification-delivery-authorization-modeled", "four-kde-notification-delivery-events-review-only", "delivery-grants-sending-and-events-disabled", "notification-actions-action-cards-and-raw-results-disabled", "production-and-host-boundary-closed"}) {
		t.Fatalf("unexpected notification delivery authorization checks: counts=%#v ids=%#v", preview.Counts, preview.CheckIDs)
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
		t.Fatalf("unsafe notification delivery authorization gates: %#v", preview)
	}
	if err := validateNoBackendTerms(preview, "KDE-safe notification delivery authorization test"); err != nil {
		t.Fatalf("validateNoBackendTerms returned error: %v", err)
	}
}

func TestProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationDeliveryAuthorizationAuditPreviewFailsClosedWithoutSources(t *testing.T) {
	root, err := os.MkdirTemp("", "xnix-lookup-dispatch-dry-run-result-notification-delivery-authorization-*")
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
	preview, err := NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationDeliveryAuthorizationAuditPreview(root)
	if err != nil {
		t.Fatalf("NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationDeliveryAuthorizationAuditPreview returned error: %v", err)
	}
	if preview.CurrentMainlineConsumed ||
		preview.RouteEnablementLookupRouteDispatchDryRunResultNotificationPresentationConsumed ||
		preview.RouteEnablementLookupRouteDispatchDryRunResultNotificationPresentationReady ||
		!preview.LookupRouteDispatchDryRunResultNotificationDeliveryAuthorizationRequired ||
		!preview.LookupRouteDispatchDryRunResultNotificationDeliveryAuthorizationModeled ||
		preview.LookupRouteDispatchDryRunResultNotificationDeliveryAuthorizationReady ||
		preview.RouteEnablementLookupRouteDispatchDryRunResultNotificationDeliveryReady ||
		preview.NotificationPresentationReviewOnlyConsumed ||
		preview.KDENotificationDeliveryAuthorizationModeled ||
		preview.NotificationCenterDeliveryAuthorizationModeled ||
		preview.InstallFailureDeliveryAuthorizationModeled ||
		preview.RepairSuggestionDeliveryAuthorizationModeled ||
		preview.EnvironmentSwitchDeliveryAuthorizationModeled ||
		preview.ApprovalInfoDeliveryAuthorizationModeled ||
		!preview.NotificationDeliveryAuthorizationReviewOnly ||
		preview.NotificationDeliveryAuthorizationGranted ||
		!preview.NotificationDeliveryDisabled ||
		!preview.NotificationActionDisabled ||
		!preview.ActionCardsRemainDisabled ||
		!preview.ExplicitUserConsentRequired ||
		preview.ExplicitUserConsentCollected ||
		preview.KDESafeRedactedResultOnly ||
		preview.RawResultHidden ||
		preview.NotificationDeliveryAuthorizationItemCount != 4 ||
		preview.ReadyNotificationDeliveryAuthorizationItemCount != 0 ||
		preview.MissingNotificationDeliveryAuthorizationItemCount != 4 ||
		preview.Counts.Total != 7 ||
		preview.Counts.Passed != 2 ||
		preview.Counts.Blocked != 5 ||
		preview.AuditDecision != "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-delivery-authorization-audit-blocked" {
		t.Fatalf("missing notification delivery authorization sources must fail closed: %#v", preview)
	}
	for _, item := range preview.NotificationDeliveryAuthorizationItems {
		if item.EvidencePresent ||
			item.CurrentMainlineConsumed ||
			item.NotificationPresentationConsumed ||
			item.NotificationPresentationReady ||
			!item.DeliveryAuthorizationRequired ||
			item.DeliveryAuthorizationModeled ||
			item.DeliveryAuthorizationReady ||
			item.DeliveryAuthorizationGranted ||
			!item.ExplicitUserConsentRequired ||
			item.ExplicitUserConsentCollected ||
			!item.DeliveryReviewOnly ||
			item.KDESafeRedactedResultOnly ||
			item.RawResultHidden ||
			item.NotificationDeliveryEnabled ||
			item.NotificationSent ||
			item.NotificationCenterEventTriggered ||
			item.NotificationActionEnabled ||
			item.ActionCardEnabled ||
			!item.SideEffectsDisabled ||
			item.HostRootModified ||
			item.DeliveryAuthorizationStatus != "missing-redacted-status-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-delivery-authorization-evidence" {
			t.Fatalf("missing notification delivery authorization item must remain closed: %#v", item)
		}
	}
}
