package owner

import (
	"os"
	"path/filepath"
	"testing"
)

func TestProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptAcceptanceAuthorizationAuditPreviewModelsAcceptance(t *testing.T) {
	preview, err := NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptAcceptanceAuthorizationAuditPreview(projectRootForOwnerTest(t))
	if err != nil {
		t.Fatalf("NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptAcceptanceAuthorizationAuditPreview returned error: %v", err)
	}

	if preview.Version != currentProjectVersion(t) ||
		preview.SchemaVersion != "xnix.runtime.production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_receipt_acceptance_authorization_audit.v1" ||
		preview.RequestType != "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-receipt-acceptance-authorization-audit-preview" ||
		preview.AuditType != "receipt-notification-action-dry-run-result-lookup-consumer-enablement-receipt-acceptance-authorization-audit" ||
		preview.AuditDecision != "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-receipt-acceptance-authorization-audit-ready-acceptance-disabled" ||
		preview.ReceiptSchema != "xnix.runtime.production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_authorization_receipt.v1" ||
		preview.OpaqueReceiptID != ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementAuthorizationReceiptID {
		t.Fatalf("unexpected lookup consumer enablement receipt acceptance authorization schema: %#v", preview)
	}
	if !preview.AcceptanceAuthorizationAuditRequired ||
		!preview.AcceptanceAuthorizationAuditModeled ||
		!preview.PersistenceAuthorizationAuditConsumed ||
		!preview.BaseAcceptancePropagationPreflightConsumed ||
		!preview.AcceptanceAuthorizationGuidanceConsumed ||
		!preview.AcceptanceAuthorizationBoundaryReady ||
		!preview.ReceiptPersistenceAuthorizationReady ||
		!preview.ReceiptAcceptancePropagationReady ||
		!preview.ConsumerEnablementReceiptAcceptanceReady ||
		!preview.ReceiptAcceptanceAuthorizationReady ||
		preview.ReceiptPresent ||
		preview.ReceiptPersisted ||
		preview.ReceiptAccepted ||
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
		preview.OpaqueLookupEnabled ||
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
		t.Fatalf("unexpected lookup consumer enablement receipt acceptance authorization decision: %#v", preview)
	}
	if preview.AuthorizationItemCount != 5 ||
		preview.RequiredAuthorizationItemCount != 5 ||
		preview.ReadyAuthorizationItemCount != 5 ||
		preview.MissingAuthorizationItemCount != 0 ||
		preview.AcceptedAuthorizationItemCount != 0 ||
		preview.PersistedAuthorizationItemCount != 0 ||
		preview.ConsumerAuthorizedItemCount != 0 ||
		preview.EnabledConsumerItemCount != 0 ||
		preview.LookupEnabledAuthorizationItemCount != 0 ||
		preview.RawExposedAuthorizationItemCount != 0 ||
		preview.SideEffectAuthorizationItemCount != 0 ||
		!sameStrings(preview.AuthorizationItemIDs, []string{"review-receipt-dry-run-result-lookup-consumer-enablement-acceptance-authorization", "renew-receipt-dry-run-result-lookup-consumer-enablement-acceptance-authorization", "open-compatibility-center-dry-run-result-lookup-consumer-enablement-acceptance-authorization", "dismiss-receipt-dry-run-result-lookup-consumer-enablement-acceptance-authorization", "support-info-dry-run-result-lookup-consumer-enablement-acceptance-authorization"}) {
		t.Fatalf("unexpected lookup consumer enablement receipt acceptance authorization inventory: %#v", preview)
	}
	for _, item := range preview.AuthorizationItems {
		if !item.EvidencePresent ||
			!item.PersistenceAuthorizationAuditConsumed ||
			!item.AcceptancePropagationPreflightConsumed ||
			!item.AcceptanceAuthorizationModeled ||
			!item.ReceiptAcceptanceAuthorizationReady ||
			item.ReceiptPresent ||
			item.ReceiptPersisted ||
			item.ReceiptAccepted ||
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
			item.AuthorizationStatus != "consumer-enablement-receipt-acceptance-authorization-modeled-acceptance-disabled" {
			t.Fatalf("unsafe lookup consumer enablement receipt acceptance authorization item: %#v", item)
		}
	}
	if preview.Counts.Total != 9 ||
		preview.Counts.Passed != 9 ||
		preview.Counts.Pending != 0 ||
		preview.Counts.Blocked != 0 ||
		!sameStrings(preview.CheckIDs, []string{"persistence-authorization-audit-consumed", "base-acceptance-propagation-preflight-consumed", "acceptance-authorization-guidance-consumed", "acceptance-authorization-boundary-modeled-only", "five-acceptance-authorization-items-present", "acceptance-authorization-items-ready-acceptance-disabled", "receipt-acceptance-and-consumption-disabled", "consumer-lookup-request-and-support-disabled", "production-and-host-boundary-closed"}) {
		t.Fatalf("unexpected lookup consumer enablement receipt acceptance authorization checks: counts=%#v ids=%#v", preview.Counts, preview.CheckIDs)
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
		t.Fatalf("unsafe lookup consumer enablement receipt acceptance authorization gates: %#v", preview)
	}
	if err := validateNoBackendTerms(preview, "lookup consumer enablement receipt acceptance authorization test"); err != nil {
		t.Fatalf("validateNoBackendTerms returned error: %v", err)
	}
}

func TestProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptAcceptanceAuthorizationAuditPreviewFailsClosedWithoutSources(t *testing.T) {
	root := t.TempDir()
	if err := os.WriteFile(filepath.Join(root, "VERSION"), []byte(currentProjectVersion(t)+"\n"), 0o600); err != nil {
		t.Fatalf("WriteFile VERSION returned error: %v", err)
	}
	preview, err := NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptAcceptanceAuthorizationAuditPreview(root)
	if err != nil {
		t.Fatalf("NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptAcceptanceAuthorizationAuditPreview returned error: %v", err)
	}
	if !preview.AcceptanceAuthorizationAuditRequired ||
		!preview.AcceptanceAuthorizationAuditModeled ||
		preview.PersistenceAuthorizationAuditConsumed ||
		preview.BaseAcceptancePropagationPreflightConsumed ||
		preview.AcceptanceAuthorizationGuidanceConsumed ||
		preview.AcceptanceAuthorizationBoundaryReady ||
		preview.ReceiptPersistenceAuthorizationReady ||
		preview.ReceiptAcceptancePropagationReady ||
		preview.ConsumerEnablementReceiptAcceptanceReady ||
		preview.ReceiptAcceptanceAuthorizationReady ||
		preview.ReceiptPersisted ||
		preview.ReceiptAccepted ||
		preview.AuthorizationAccepted ||
		preview.ConsumerConsumptionAuthorized ||
		preview.ConsumerEnablementAuthorized ||
		preview.KDEConsumerEnabled ||
		preview.RuntimeConsumerEnabled ||
		preview.RawResultExposed ||
		preview.AuthorizationItemCount != 5 ||
		preview.RequiredAuthorizationItemCount != 5 ||
		preview.ReadyAuthorizationItemCount != 0 ||
		preview.MissingAuthorizationItemCount != 5 ||
		preview.Counts.Total != 9 ||
		preview.Counts.Passed != 3 ||
		preview.Counts.Blocked != 6 ||
		preview.AuditDecision != "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-receipt-acceptance-authorization-audit-blocked" {
		t.Fatalf("missing lookup consumer enablement receipt acceptance authorization sources must fail closed: %#v", preview)
	}
	for _, item := range preview.AuthorizationItems {
		if item.EvidencePresent ||
			item.PersistenceAuthorizationAuditConsumed ||
			item.AcceptancePropagationPreflightConsumed ||
			item.AcceptanceAuthorizationModeled ||
			item.ReceiptAcceptanceAuthorizationReady ||
			item.ReceiptPersisted ||
			item.ReceiptAccepted ||
			item.AuthorizationAccepted ||
			item.ConsumerConsumptionAuthorized ||
			item.ConsumerEnablementAuthorized ||
			item.KDEConsumerEnabled ||
			item.RuntimeConsumerEnabled ||
			item.RawResultExposed ||
			item.UserVisible ||
			item.AuthorizationStatus != "missing-consumer-enablement-receipt-acceptance-authorization-evidence" {
			t.Fatalf("missing lookup consumer enablement receipt acceptance authorization item must remain closed: %#v", item)
		}
	}
}
