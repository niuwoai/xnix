package owner

import (
	"os"
	"path/filepath"
	"testing"
)

func TestProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumptionGateAuditPreviewModelsConsumption(t *testing.T) {
	preview, err := NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumptionGateAuditPreview(projectRootForOwnerTest(t))
	if err != nil {
		t.Fatalf("NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumptionGateAuditPreview returned error: %v", err)
	}

	if preview.Version != currentProjectVersion(t) ||
		preview.SchemaVersion != "xnix.runtime.production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_receipt_consumption_gate_audit.v1" ||
		preview.RequestType != "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-receipt-consumption-gate-audit-preview" ||
		preview.AuditType != "receipt-notification-action-dry-run-result-lookup-consumer-enablement-receipt-consumption-gate-audit" ||
		preview.AuditDecision != "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-receipt-consumption-gate-audit-ready-consumption-disabled" ||
		preview.ReceiptSchema != "xnix.runtime.production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_authorization_receipt.v1" ||
		preview.OpaqueReceiptID != ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementAuthorizationReceiptID {
		t.Fatalf("unexpected lookup consumer enablement receipt consumption gate schema: %#v", preview)
	}
	if !preview.ConsumptionGateRequired ||
		!preview.ConsumptionGateModeled ||
		!preview.AcceptanceAuthorizationAuditConsumed ||
		!preview.ConsumerRedactionAuditConsumed ||
		!preview.LookupRouteAuthorizationAuditConsumed ||
		!preview.ConsumptionGateGuidanceConsumed ||
		!preview.ConsumptionGateBoundaryReady ||
		!preview.ReceiptAcceptanceAuthorizationReady ||
		!preview.ConsumerRedactionBoundaryReady ||
		!preview.LookupRouteAuthorizationBoundaryReady ||
		!preview.ConsumerEnablementReceiptConsumptionReady ||
		!preview.ReceiptConsumptionGateReady ||
		preview.ReceiptPresent ||
		preview.ReceiptPersisted ||
		preview.ReceiptAccepted ||
		preview.ReceiptConsumed ||
		preview.AuthorizationAccepted ||
		preview.ReceiptWriterEnabled ||
		preview.ReceiptPersistenceEnabled ||
		preview.ReceiptLookupWritesEnabled ||
		preview.ReceiptReplayEnabled ||
		preview.ReceiptExpiryWriteEnabled ||
		preview.ReceiptRevocationWriteEnabled ||
		preview.ConsumerConsumptionAuthorized ||
		preview.ConsumerEnablementAuthorized ||
		preview.KDEConsumerEnabled ||
		preview.RuntimeConsumerEnabled ||
		preview.LookupRouteAuthorized ||
		preview.LookupRouteEnabled ||
		preview.LookupRoutePersisted ||
		preview.OpaqueLookupEnabled ||
		preview.OpaqueLookupPersisted ||
		preview.RedactedSummaryPersisted ||
		preview.RawResultExposed ||
		preview.RuntimeDiagnosticsPersisted ||
		preview.DryRunResultPersisted ||
		preview.DispatchDryRunExecuted ||
		preview.RequestObjectCreationEnabled ||
		preview.RequestObjectDispatchEnabled ||
		preview.PortalRequestCreated ||
		preview.NotificationActionEnabled ||
		preview.CompatibilityCenterOpened ||
		preview.SupportBundleExported ||
		preview.SupportCaseCreated ||
		preview.ProductionReadiness ||
		preview.ProductionOwnershipReady {
		t.Fatalf("unexpected lookup consumer enablement receipt consumption gate decision: %#v", preview)
	}
	if preview.GateItemCount != 5 ||
		preview.RequiredGateItemCount != 5 ||
		preview.ReadyGateItemCount != 5 ||
		preview.MissingGateItemCount != 0 ||
		preview.AcceptedReceiptItemCount != 0 ||
		preview.ConsumedReceiptItemCount != 0 ||
		preview.AuthorizedConsumerItemCount != 0 ||
		preview.EnabledConsumerItemCount != 0 ||
		preview.LookupEnabledGateItemCount != 0 ||
		preview.RawExposedGateItemCount != 0 ||
		preview.SideEffectGateItemCount != 0 ||
		!sameStrings(preview.GateItemIDs, []string{"review-receipt-dry-run-result-lookup-consumer-enablement-consumption-gate", "renew-receipt-dry-run-result-lookup-consumer-enablement-consumption-gate", "open-compatibility-center-dry-run-result-lookup-consumer-enablement-consumption-gate", "dismiss-receipt-dry-run-result-lookup-consumer-enablement-consumption-gate", "support-info-dry-run-result-lookup-consumer-enablement-consumption-gate"}) {
		t.Fatalf("unexpected lookup consumer enablement receipt consumption gate inventory: %#v", preview)
	}
	for _, item := range preview.GateItems {
		if !item.EvidencePresent ||
			!item.AcceptanceAuthorizationAuditConsumed ||
			!item.ConsumerRedactionAuditConsumed ||
			!item.LookupRouteAuthorizationAuditConsumed ||
			!item.ConsumptionGateModeled ||
			!item.ReceiptConsumptionGateReady ||
			item.ReceiptPresent ||
			item.ReceiptPersisted ||
			item.ReceiptAccepted ||
			item.ReceiptConsumed ||
			item.AuthorizationAccepted ||
			item.ReceiptWriterEnabled ||
			item.ReceiptPersistenceEnabled ||
			item.ReceiptLookupWritesEnabled ||
			item.ReceiptReplayEnabled ||
			item.ReceiptExpiryWriteEnabled ||
			item.ReceiptRevocationWriteEnabled ||
			item.ConsumerConsumptionAuthorized ||
			item.ConsumerEnablementAuthorized ||
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
			item.CallerStateRootRequired ||
			item.StateRootPathExposed ||
			item.FilePathsExposed ||
			item.FileContentRead ||
			item.RequestObjectCreated ||
			item.RequestObjectDispatched ||
			item.PortalRequestCreated ||
			item.NotificationActionEnabled ||
			item.CompatibilityCenterOpened ||
			item.SupportBundleExported ||
			item.SupportCaseCreated ||
			item.ProductionReadiness ||
			item.ProductionOwnershipReady ||
			!item.SideEffectsDisabled ||
			item.HostRootModified ||
			item.InternalDetailsExposed ||
			item.GateStatus != "consumer-enablement-receipt-consumption-gate-modeled-consumption-disabled" {
			t.Fatalf("unsafe lookup consumer enablement receipt consumption gate item: %#v", item)
		}
	}
	if preview.Counts.Total != 10 ||
		preview.Counts.Passed != 10 ||
		preview.Counts.Pending != 0 ||
		preview.Counts.Blocked != 0 ||
		!sameStrings(preview.CheckIDs, []string{"acceptance-authorization-audit-consumed", "consumer-redaction-audit-consumed", "lookup-route-authorization-audit-consumed", "consumption-gate-guidance-consumed", "consumption-gate-modeled-only", "five-consumption-gate-items-present", "consumption-gate-items-ready-consumption-disabled", "receipt-acceptance-and-consumption-disabled", "consumer-lookup-request-and-support-disabled", "production-and-host-boundary-closed"}) {
		t.Fatalf("unexpected lookup consumer enablement receipt consumption gate checks: counts=%#v ids=%#v", preview.Counts, preview.CheckIDs)
	}
	if preview.SystemServiceStarted ||
		preview.SessionBusClaimed ||
		preview.ProductionBusClaimed ||
		preview.ProductionOwnerEnabled ||
		preview.ProductionActivationReady ||
		preview.WriteMethodsEnabled ||
		preview.RuntimeWritesEnabled ||
		preview.ReceiptWriterEnabled ||
		preview.ReceiptPersistenceEnabled ||
		preview.ReceiptLookupWritesEnabled ||
		preview.ReceiptReplayEnabled ||
		preview.ReceiptExpiryWriteEnabled ||
		preview.ReceiptRevocationWriteEnabled ||
		preview.RequestObjectCreationEnabled ||
		preview.RequestObjectDispatchEnabled ||
		preview.PortalRequestCreated ||
		preview.NotificationActionEnabled ||
		preview.CompatibilityCenterOpened ||
		preview.SupportBundleExported ||
		preview.BackendLaunchEnabled ||
		preview.HostRootModified ||
		preview.BackendDetailsExposed {
		t.Fatalf("unsafe lookup consumer enablement receipt consumption gate gates: %#v", preview)
	}
	if err := validateNoBackendTerms(preview, "lookup consumer enablement receipt consumption gate test"); err != nil {
		t.Fatalf("validateNoBackendTerms returned error: %v", err)
	}
}

func TestProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumptionGateAuditPreviewFailsClosedWithoutSources(t *testing.T) {
	root := t.TempDir()
	if err := os.WriteFile(filepath.Join(root, "VERSION"), []byte(currentProjectVersion(t)+"\n"), 0o600); err != nil {
		t.Fatalf("WriteFile VERSION returned error: %v", err)
	}
	preview, err := NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumptionGateAuditPreview(root)
	if err != nil {
		t.Fatalf("NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumptionGateAuditPreview returned error: %v", err)
	}
	if !preview.ConsumptionGateRequired ||
		!preview.ConsumptionGateModeled ||
		preview.AcceptanceAuthorizationAuditConsumed ||
		preview.ConsumerRedactionAuditConsumed ||
		preview.LookupRouteAuthorizationAuditConsumed ||
		preview.ConsumptionGateGuidanceConsumed ||
		preview.ConsumptionGateBoundaryReady ||
		preview.ReceiptAcceptanceAuthorizationReady ||
		preview.ConsumerRedactionBoundaryReady ||
		preview.LookupRouteAuthorizationBoundaryReady ||
		preview.ConsumerEnablementReceiptConsumptionReady ||
		preview.ReceiptConsumptionGateReady ||
		preview.ReceiptAccepted ||
		preview.ReceiptConsumed ||
		preview.ConsumerConsumptionAuthorized ||
		preview.ConsumerEnablementAuthorized ||
		preview.KDEConsumerEnabled ||
		preview.RuntimeConsumerEnabled ||
		preview.RawResultExposed ||
		preview.GateItemCount != 5 ||
		preview.RequiredGateItemCount != 5 ||
		preview.ReadyGateItemCount != 0 ||
		preview.MissingGateItemCount != 5 ||
		preview.Counts.Total != 10 ||
		preview.Counts.Passed != 3 ||
		preview.Counts.Blocked != 7 ||
		preview.AuditDecision != "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-receipt-consumption-gate-audit-blocked" {
		t.Fatalf("missing lookup consumer enablement receipt consumption gate sources must fail closed: %#v", preview)
	}
	for _, item := range preview.GateItems {
		if item.EvidencePresent ||
			item.AcceptanceAuthorizationAuditConsumed ||
			item.ConsumerRedactionAuditConsumed ||
			item.LookupRouteAuthorizationAuditConsumed ||
			item.ConsumptionGateModeled ||
			item.ReceiptConsumptionGateReady ||
			item.ReceiptAccepted ||
			item.ReceiptConsumed ||
			item.ConsumerConsumptionAuthorized ||
			item.ConsumerEnablementAuthorized ||
			item.KDEConsumerEnabled ||
			item.RuntimeConsumerEnabled ||
			item.RawResultExposed ||
			item.UserVisible ||
			item.GateStatus != "missing-consumer-enablement-receipt-consumption-gate-evidence" {
			t.Fatalf("missing lookup consumer enablement receipt consumption gate item must remain closed: %#v", item)
		}
	}
}
