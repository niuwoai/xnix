package owner

import (
	"os"
	"path/filepath"
	"testing"
)

func TestProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultPresentationConsentAuditPreviewModelsConsent(t *testing.T) {
	preview, err := NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultPresentationConsentAuditPreview(projectRootForOwnerTest(t))
	if err != nil {
		t.Fatalf("NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultPresentationConsentAuditPreview returned error: %v", err)
	}

	if preview.Version != currentProjectVersion(t) ||
		preview.SchemaVersion != "xnix.runtime.production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_kde_safe_redacted_status_closed_record_writer_storage_root_authorization_receipt_dry_run_lookup_result_consumer_projection_route_enablement_lookup_route_dispatch_dry_run_result_presentation_consent_audit.v1" ||
		preview.RequestType != "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-presentation-consent-audit-preview" ||
		preview.AuditType != "receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-presentation-consent-audit" ||
		preview.AuditDecision != "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-presentation-consent-audit-ready-consent-review-only-notifications-disabled-action-cards-disabled" {
		t.Fatalf("unexpected presentation consent schema: %#v", preview)
	}
	if !preview.CurrentMainlineConsumed ||
		!preview.RouteEnablementLookupRouteDispatchDryRunResultRedactedPresentationEligibilityConsumed ||
		!preview.RouteEnablementLookupRouteDispatchDryRunResultRedactedPresentationEligibilityReady ||
		!preview.LookupRouteDispatchDryRunResultPresentationConsentRequired ||
		!preview.LookupRouteDispatchDryRunResultPresentationConsentModeled ||
		!preview.LookupRouteDispatchDryRunResultPresentationConsentReady ||
		!preview.RouteEnablementLookupRouteDispatchDryRunResultPresentationConsentReady ||
		!preview.EligibleRedactedPresentationConsumed ||
		!preview.KDEPresentationConsentModeled ||
		!preview.CompatibilityCenterPresentationConsentModeled ||
		!preview.NotificationCenterPresentationConsentModeled ||
		!preview.SettingsPresentationConsentModeled ||
		!preview.RuntimeDiagnosticsPresentationConsentModeled ||
		!preview.NotificationAndActionCardConsentBoundaryReady ||
		!preview.ExplicitUserConsentRequired ||
		preview.ExplicitUserConsentCollected ||
		!preview.ConsentReviewOnly ||
		!preview.NotificationsRemainDisabled ||
		!preview.ActionCardsRemainDisabled ||
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
		preview.NotificationCenterEventTriggered ||
		preview.ActionCardEnabled ||
		preview.CompatibilityCenterOpened ||
		preview.SupportBundleExported ||
		preview.SupportCaseCreated ||
		!preview.RuntimeOwned ||
		!preview.GoRuntimeBacked ||
		preview.KDEPolicyOwner ||
		preview.ProductionReadiness ||
		preview.ProductionOwnershipReady {
		t.Fatalf("unsafe presentation consent decision: %#v", preview)
	}
	if preview.PresentationConsentItemCount != 4 ||
		preview.RequiredPresentationConsentItemCount != 4 ||
		preview.ReadyPresentationConsentItemCount != 4 ||
		preview.MissingPresentationConsentItemCount != 0 ||
		preview.ExplicitConsentGrantedItemCount != 0 ||
		preview.NotificationTriggerEnabledItemCount != 0 ||
		preview.ActionCardEnabledItemCount != 0 ||
		preview.SideEffectPresentationConsentItemCount != 0 {
		t.Fatalf("unexpected presentation consent counts: %#v", preview)
	}
	if !sameStrings(preview.PresentationConsentItemIDs, []string{
		"route-enablement-lookup-route-dispatch-dry-run-result-presentation-consent-compatibility-center",
		"route-enablement-lookup-route-dispatch-dry-run-result-presentation-consent-notification-center",
		"route-enablement-lookup-route-dispatch-dry-run-result-presentation-consent-settings",
		"route-enablement-lookup-route-dispatch-dry-run-result-presentation-consent-runtime-diagnostics",
	}) {
		t.Fatalf("unexpected presentation consent item ids: %#v", preview.PresentationConsentItemIDs)
	}
	for _, item := range preview.PresentationConsentItems {
		if !item.EvidencePresent ||
			!item.CurrentMainlineConsumed ||
			!item.RouteEnablementLookupRouteDispatchDryRunResultRedactedPresentationEligibilityConsumed ||
			!item.RouteEnablementLookupRouteDispatchDryRunResultRedactedPresentationEligibilityReady ||
			!item.PresentationConsentRequired ||
			!item.PresentationConsentModeled ||
			!item.PresentationConsentReady ||
			!item.ExplicitUserConsentRequired ||
			item.ExplicitUserConsentCollected ||
			!item.ConsentReviewOnly ||
			!item.KDESafeRedactedResultOnly ||
			!item.RawResultHidden ||
			item.NotificationTriggerEnabled ||
			item.ActionCardEnabled ||
			!item.SideEffectsDisabled ||
			!item.RuntimeOwned ||
			!item.GoRuntimeBacked ||
			item.KDEPolicyOwner ||
			item.HostRootModified ||
			item.InternalDetailsExposed ||
			item.PresentationConsentStatus != "redacted-status-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-presentation-consent-ready-review-only-notifications-disabled-action-cards-disabled" {
			t.Fatalf("unsafe presentation consent item: %#v", item)
		}
	}
	if preview.Counts.Total != 7 ||
		preview.Counts.Passed != 7 ||
		preview.Counts.Pending != 0 ||
		preview.Counts.Blocked != 0 ||
		!sameStrings(preview.CheckIDs, []string{"current-mainline-consumed", "route-enablement-lookup-route-dispatch-dry-run-result-redacted-presentation-eligibility-consumed", "lookup-route-dispatch-dry-run-result-presentation-consent-modeled", "four-kde-presentation-surfaces-consent-review-only", "notifications-and-action-cards-disabled", "raw-result-exposure-and-persistence-disabled", "production-and-host-boundary-closed"}) {
		t.Fatalf("unexpected presentation consent checks: counts=%#v ids=%#v", preview.Counts, preview.CheckIDs)
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
		t.Fatalf("unsafe presentation consent gates: %#v", preview)
	}
	if err := validateNoBackendTerms(preview, "KDE-safe presentation consent test"); err != nil {
		t.Fatalf("validateNoBackendTerms returned error: %v", err)
	}
}

func TestProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultPresentationConsentAuditPreviewFailsClosedWithoutSources(t *testing.T) {
	root, err := os.MkdirTemp("", "xnix-lookup-dispatch-dry-run-result-presentation-consent-*")
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
	preview, err := NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultPresentationConsentAuditPreview(root)
	if err != nil {
		t.Fatalf("NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultPresentationConsentAuditPreview returned error: %v", err)
	}
	if preview.CurrentMainlineConsumed ||
		preview.RouteEnablementLookupRouteDispatchDryRunResultRedactedPresentationEligibilityConsumed ||
		preview.RouteEnablementLookupRouteDispatchDryRunResultRedactedPresentationEligibilityReady ||
		!preview.LookupRouteDispatchDryRunResultPresentationConsentRequired ||
		!preview.LookupRouteDispatchDryRunResultPresentationConsentModeled ||
		preview.LookupRouteDispatchDryRunResultPresentationConsentReady ||
		preview.RouteEnablementLookupRouteDispatchDryRunResultPresentationConsentReady ||
		preview.EligibleRedactedPresentationConsumed ||
		preview.KDEPresentationConsentModeled ||
		preview.CompatibilityCenterPresentationConsentModeled ||
		preview.NotificationCenterPresentationConsentModeled ||
		preview.SettingsPresentationConsentModeled ||
		preview.RuntimeDiagnosticsPresentationConsentModeled ||
		preview.NotificationAndActionCardConsentBoundaryReady ||
		!preview.ExplicitUserConsentRequired ||
		preview.ExplicitUserConsentCollected ||
		!preview.ConsentReviewOnly ||
		!preview.NotificationsRemainDisabled ||
		!preview.ActionCardsRemainDisabled ||
		preview.KDESafeRedactedResultOnly ||
		preview.RawResultHidden ||
		preview.PresentationConsentItemCount != 4 ||
		preview.ReadyPresentationConsentItemCount != 0 ||
		preview.MissingPresentationConsentItemCount != 4 ||
		preview.Counts.Total != 7 ||
		preview.Counts.Passed != 2 ||
		preview.Counts.Blocked != 5 ||
		preview.AuditDecision != "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-presentation-consent-audit-blocked" {
		t.Fatalf("missing presentation consent sources must fail closed: %#v", preview)
	}
	for _, item := range preview.PresentationConsentItems {
		if item.EvidencePresent ||
			item.CurrentMainlineConsumed ||
			item.RouteEnablementLookupRouteDispatchDryRunResultRedactedPresentationEligibilityConsumed ||
			item.RouteEnablementLookupRouteDispatchDryRunResultRedactedPresentationEligibilityReady ||
			!item.PresentationConsentRequired ||
			item.PresentationConsentModeled ||
			item.PresentationConsentReady ||
			!item.ExplicitUserConsentRequired ||
			item.ExplicitUserConsentCollected ||
			!item.ConsentReviewOnly ||
			item.KDESafeRedactedResultOnly ||
			item.RawResultHidden ||
			item.NotificationTriggerEnabled ||
			item.ActionCardEnabled ||
			!item.SideEffectsDisabled ||
			item.HostRootModified ||
			item.PresentationConsentStatus != "missing-redacted-status-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-presentation-consent-evidence" {
			t.Fatalf("missing presentation consent item must remain closed: %#v", item)
		}
	}
}
