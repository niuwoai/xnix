package owner

import (
	"os"
	"path/filepath"
	"testing"
)

func TestProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerAuthorizationAuditPreviewModelsAuthorization(t *testing.T) {
	preview, err := NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerAuthorizationAuditPreview(projectRootForOwnerTest(t))
	if err != nil {
		t.Fatalf("NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerAuthorizationAuditPreview returned error: %v", err)
	}

	if preview.Version != currentProjectVersion(t) ||
		preview.SchemaVersion != "xnix.runtime.production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_receipt_consumer_authorization_audit.v1" ||
		preview.RequestType != "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-receipt-consumer-authorization-audit-preview" ||
		preview.AuditType != "receipt-notification-action-dry-run-result-lookup-consumer-enablement-receipt-consumer-authorization-audit" ||
		preview.AuditDecision != "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-receipt-consumer-authorization-audit-ready-authorization-disabled" ||
		preview.ReceiptSchema != "xnix.runtime.production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_authorization_receipt.v1" ||
		preview.OpaqueReceiptID != ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementAuthorizationReceiptID {
		t.Fatalf("unexpected lookup consumer enablement receipt consumer authorization schema: %#v", preview)
	}
	if !preview.ConsumerAuthorizationAuditRequired ||
		!preview.ConsumerAuthorizationAuditModeled ||
		!preview.ConsumptionGateAuditConsumed ||
		!preview.ConsumerRedactionAuditConsumed ||
		!preview.LookupRouteAuthorizationAuditConsumed ||
		!preview.ConsumerAuthorizationGuidanceConsumed ||
		!preview.ConsumerAuthorizationBoundaryReady ||
		!preview.ReceiptConsumptionGateReady ||
		!preview.ConsumerRedactionBoundaryReady ||
		!preview.LookupRouteAuthorizationBoundaryReady ||
		!preview.ConsumerAuthorizationReady ||
		preview.ReceiptAccepted ||
		preview.ReceiptConsumed ||
		preview.ConsumerAuthorizationGranted ||
		preview.ConsumerConsumptionAuthorized ||
		preview.ConsumerEnablementAuthorized ||
		preview.KDEConsumerAuthorized ||
		preview.RuntimeConsumerAuthorized ||
		preview.KDEConsumerEnabled ||
		preview.RuntimeConsumerEnabled ||
		preview.LookupRouteEnabled ||
		preview.OpaqueLookupEnabled ||
		preview.RawResultExposed ||
		preview.DispatchDryRunExecuted ||
		preview.RequestObjectCreationEnabled ||
		preview.NotificationActionEnabled ||
		preview.CompatibilityCenterOpened ||
		preview.ProductionReadiness ||
		preview.ProductionOwnershipReady {
		t.Fatalf("unexpected lookup consumer enablement receipt consumer authorization decision: %#v", preview)
	}
	if preview.AuthorizationItemCount != 5 ||
		preview.RequiredAuthorizationItemCount != 5 ||
		preview.ReadyAuthorizationItemCount != 5 ||
		preview.MissingAuthorizationItemCount != 0 ||
		preview.GrantedAuthorizationItemCount != 0 ||
		preview.AcceptedReceiptItemCount != 0 ||
		preview.ConsumedReceiptItemCount != 0 ||
		preview.AuthorizedConsumerItemCount != 0 ||
		preview.EnabledConsumerItemCount != 0 ||
		preview.LookupEnabledAuthorizationItemCount != 0 ||
		preview.RawExposedAuthorizationItemCount != 0 ||
		preview.SideEffectAuthorizationItemCount != 0 ||
		!sameStrings(preview.AuthorizationItemIDs, []string{"review-receipt-dry-run-result-lookup-consumer-enablement-consumer-authorization", "renew-receipt-dry-run-result-lookup-consumer-enablement-consumer-authorization", "open-compatibility-center-dry-run-result-lookup-consumer-enablement-consumer-authorization", "dismiss-receipt-dry-run-result-lookup-consumer-enablement-consumer-authorization", "support-info-dry-run-result-lookup-consumer-enablement-consumer-authorization"}) {
		t.Fatalf("unexpected lookup consumer enablement receipt consumer authorization inventory: %#v", preview)
	}
	for _, item := range preview.AuthorizationItems {
		if !item.EvidencePresent ||
			!item.ConsumptionGateAuditConsumed ||
			!item.ConsumerRedactionAuditConsumed ||
			!item.LookupRouteAuthorizationAuditConsumed ||
			!item.ConsumerAuthorizationModeled ||
			!item.ConsumerAuthorizationReady ||
			item.ReceiptAccepted ||
			item.ReceiptConsumed ||
			item.ConsumerAuthorizationGranted ||
			item.ConsumerConsumptionAuthorized ||
			item.ConsumerEnablementAuthorized ||
			item.KDEConsumerAuthorized ||
			item.RuntimeConsumerAuthorized ||
			item.KDEConsumerEnabled ||
			item.RuntimeConsumerEnabled ||
			item.LookupRouteEnabled ||
			item.OpaqueLookupEnabled ||
			item.RawResultExposed ||
			!item.UserVisible ||
			!item.ReviewOnly ||
			!item.RuntimeOwned ||
			!item.GoRuntimeBacked ||
			item.KDEPolicyOwner ||
			item.StateRootPathExposed ||
			item.FilePathsExposed ||
			item.FileContentRead ||
			item.RequestObjectCreated ||
			item.NotificationActionEnabled ||
			item.CompatibilityCenterOpened ||
			item.ProductionReadiness ||
			!item.SideEffectsDisabled ||
			item.HostRootModified ||
			item.InternalDetailsExposed ||
			item.AuthorizationStatus != "consumer-enablement-receipt-consumer-authorization-modeled-authorization-disabled" {
			t.Fatalf("unsafe lookup consumer enablement receipt consumer authorization item: %#v", item)
		}
	}
	if preview.Counts.Total != 10 ||
		preview.Counts.Passed != 10 ||
		preview.Counts.Pending != 0 ||
		preview.Counts.Blocked != 0 ||
		!sameStrings(preview.CheckIDs, []string{"consumption-gate-audit-consumed", "consumer-redaction-audit-consumed", "lookup-route-authorization-audit-consumed", "consumer-authorization-guidance-consumed", "consumer-authorization-boundary-modeled-only", "five-consumer-authorization-items-present", "consumer-authorization-items-ready-authorization-disabled", "receipt-consumption-and-consumer-authorization-disabled", "consumer-lookup-request-and-support-disabled", "production-and-host-boundary-closed"}) {
		t.Fatalf("unexpected lookup consumer enablement receipt consumer authorization checks: counts=%#v ids=%#v", preview.Counts, preview.CheckIDs)
	}
	if preview.SystemServiceStarted ||
		preview.SessionBusClaimed ||
		preview.ProductionBusClaimed ||
		preview.ProductionOwnerEnabled ||
		preview.WriteMethodsEnabled ||
		preview.RuntimeWritesEnabled ||
		preview.BackendLaunchEnabled ||
		preview.HostRootModified ||
		preview.BackendDetailsExposed {
		t.Fatalf("unsafe lookup consumer enablement receipt consumer authorization gates: %#v", preview)
	}
	if err := validateNoBackendTerms(preview, "lookup consumer enablement receipt consumer authorization test"); err != nil {
		t.Fatalf("validateNoBackendTerms returned error: %v", err)
	}
}

func TestProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerAuthorizationAuditPreviewFailsClosedWithoutSources(t *testing.T) {
	root := t.TempDir()
	if err := os.WriteFile(filepath.Join(root, "VERSION"), []byte(currentProjectVersion(t)+"\n"), 0o600); err != nil {
		t.Fatalf("WriteFile VERSION returned error: %v", err)
	}
	preview, err := NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerAuthorizationAuditPreview(root)
	if err != nil {
		t.Fatalf("NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerAuthorizationAuditPreview returned error: %v", err)
	}
	if !preview.ConsumerAuthorizationAuditRequired ||
		!preview.ConsumerAuthorizationAuditModeled ||
		preview.ConsumptionGateAuditConsumed ||
		preview.ConsumerRedactionAuditConsumed ||
		preview.LookupRouteAuthorizationAuditConsumed ||
		preview.ConsumerAuthorizationGuidanceConsumed ||
		preview.ConsumerAuthorizationBoundaryReady ||
		preview.ConsumerAuthorizationReady ||
		preview.ConsumerAuthorizationGranted ||
		preview.ConsumerConsumptionAuthorized ||
		preview.ConsumerEnablementAuthorized ||
		preview.KDEConsumerEnabled ||
		preview.RuntimeConsumerEnabled ||
		preview.AuthorizationItemCount != 5 ||
		preview.RequiredAuthorizationItemCount != 5 ||
		preview.ReadyAuthorizationItemCount != 0 ||
		preview.MissingAuthorizationItemCount != 5 ||
		preview.Counts.Total != 10 ||
		preview.Counts.Passed != 3 ||
		preview.Counts.Blocked != 7 ||
		preview.AuditDecision != "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-receipt-consumer-authorization-audit-blocked" {
		t.Fatalf("missing lookup consumer enablement receipt consumer authorization sources must fail closed: %#v", preview)
	}
	for _, item := range preview.AuthorizationItems {
		if item.EvidencePresent ||
			item.ConsumptionGateAuditConsumed ||
			item.ConsumerRedactionAuditConsumed ||
			item.LookupRouteAuthorizationAuditConsumed ||
			item.ConsumerAuthorizationModeled ||
			item.ConsumerAuthorizationReady ||
			item.ConsumerAuthorizationGranted ||
			item.ConsumerConsumptionAuthorized ||
			item.ConsumerEnablementAuthorized ||
			item.KDEConsumerEnabled ||
			item.RuntimeConsumerEnabled ||
			item.RawResultExposed ||
			item.UserVisible ||
			item.AuthorizationStatus != "missing-consumer-enablement-receipt-consumer-authorization-evidence" {
			t.Fatalf("missing lookup consumer enablement receipt consumer authorization item must remain closed: %#v", item)
		}
	}
}
