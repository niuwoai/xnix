package owner

import (
	"os"
	"path/filepath"
	"testing"
)

func TestProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptPersistenceAuthorizationAuditPreviewModelsPersistence(t *testing.T) {
	preview, err := NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptPersistenceAuthorizationAuditPreview(projectRootForOwnerTest(t))
	if err != nil {
		t.Fatalf("NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptPersistenceAuthorizationAuditPreview returned error: %v", err)
	}

	if preview.Version != currentProjectVersion(t) ||
		preview.SchemaVersion != "xnix.runtime.production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_receipt_persistence_authorization_audit.v1" ||
		preview.RequestType != "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-receipt-persistence-authorization-audit-preview" ||
		preview.AuditType != "receipt-notification-action-dry-run-result-lookup-consumer-enablement-receipt-persistence-authorization-audit" ||
		preview.AuditDecision != "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-receipt-persistence-authorization-audit-ready-persistence-disabled" ||
		preview.ReceiptSchema != "xnix.runtime.production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_authorization_receipt.v1" ||
		preview.OpaqueReceiptID != ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementAuthorizationReceiptID {
		t.Fatalf("unexpected lookup consumer enablement receipt persistence authorization schema: %#v", preview)
	}
	if !preview.PersistenceAuthorizationAuditRequired ||
		!preview.PersistenceAuthorizationAuditModeled ||
		!preview.WriterAuthorizationAuditConsumed ||
		!preview.BasePersistenceThreatReviewConsumed ||
		!preview.PersistenceAuthorizationGuidanceConsumed ||
		!preview.PersistenceAuthorizationBoundaryReady ||
		!preview.ReceiptWriterAuthorizationReady ||
		!preview.ReceiptPersistenceThreatBoundaryReady ||
		!preview.ConsumerEnablementReceiptPersistenceReady ||
		!preview.ReceiptPersistenceAuthorizationReady ||
		preview.ReceiptPresent ||
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
		preview.ProductionReadiness ||
		preview.ProductionOwnershipReady {
		t.Fatalf("unexpected lookup consumer enablement receipt persistence authorization decision: %#v", preview)
	}
	if preview.AuthorizationItemCount != 5 ||
		preview.RequiredAuthorizationItemCount != 5 ||
		preview.ReadyAuthorizationItemCount != 5 ||
		preview.MissingAuthorizationItemCount != 0 ||
		preview.PersistenceEnabledAuthorizationItemCount != 0 ||
		preview.ReplayEnabledAuthorizationItemCount != 0 ||
		preview.ExpiryWriteEnabledAuthorizationItemCount != 0 ||
		preview.RevocationWriteEnabledAuthorizationItemCount != 0 ||
		preview.AcceptedAuthorizationItemCount != 0 ||
		preview.EnabledConsumerItemCount != 0 ||
		preview.RawExposedAuthorizationItemCount != 0 ||
		preview.SideEffectAuthorizationItemCount != 0 ||
		!sameStrings(preview.AuthorizationItemIDs, []string{"review-receipt-dry-run-result-lookup-consumer-enablement-persistence-authorization", "renew-receipt-dry-run-result-lookup-consumer-enablement-persistence-authorization", "open-compatibility-center-dry-run-result-lookup-consumer-enablement-persistence-authorization", "dismiss-receipt-dry-run-result-lookup-consumer-enablement-persistence-authorization", "support-info-dry-run-result-lookup-consumer-enablement-persistence-authorization"}) {
		t.Fatalf("unexpected lookup consumer enablement receipt persistence authorization inventory: %#v", preview)
	}
	for _, item := range preview.AuthorizationItems {
		if !item.EvidencePresent ||
			!item.WriterAuthorizationAuditConsumed ||
			!item.PersistenceThreatReviewConsumed ||
			!item.PersistenceAuthorizationModeled ||
			!item.ReceiptPersistenceAuthorizationReady ||
			item.ReceiptPresent ||
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
			item.AuthorizationStatus != "consumer-enablement-receipt-persistence-authorization-modeled-persistence-disabled" {
			t.Fatalf("unsafe lookup consumer enablement receipt persistence authorization item: %#v", item)
		}
	}
	if preview.Counts.Total != 9 ||
		preview.Counts.Passed != 9 ||
		preview.Counts.Pending != 0 ||
		preview.Counts.Blocked != 0 ||
		!sameStrings(preview.CheckIDs, []string{"writer-authorization-audit-consumed", "base-persistence-threat-review-consumed", "persistence-authorization-guidance-consumed", "persistence-authorization-boundary-modeled-only", "five-persistence-authorization-items-present", "persistence-authorization-items-ready-persistence-disabled", "receipt-persistence-replay-expiry-revocation-disabled", "consumer-lookup-request-and-support-disabled", "production-and-host-boundary-closed"}) {
		t.Fatalf("unexpected lookup consumer enablement receipt persistence authorization checks: counts=%#v ids=%#v", preview.Counts, preview.CheckIDs)
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
		t.Fatalf("unsafe lookup consumer enablement receipt persistence authorization gates: %#v", preview)
	}
	if err := validateNoBackendTerms(preview, "lookup consumer enablement receipt persistence authorization test"); err != nil {
		t.Fatalf("validateNoBackendTerms returned error: %v", err)
	}
}

func TestProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptPersistenceAuthorizationAuditPreviewFailsClosedWithoutSources(t *testing.T) {
	root := t.TempDir()
	if err := os.WriteFile(filepath.Join(root, "VERSION"), []byte(currentProjectVersion(t)+"\n"), 0o600); err != nil {
		t.Fatalf("WriteFile VERSION returned error: %v", err)
	}
	preview, err := NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptPersistenceAuthorizationAuditPreview(root)
	if err != nil {
		t.Fatalf("NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptPersistenceAuthorizationAuditPreview returned error: %v", err)
	}
	if !preview.PersistenceAuthorizationAuditRequired ||
		!preview.PersistenceAuthorizationAuditModeled ||
		preview.WriterAuthorizationAuditConsumed ||
		preview.BasePersistenceThreatReviewConsumed ||
		preview.PersistenceAuthorizationGuidanceConsumed ||
		preview.PersistenceAuthorizationBoundaryReady ||
		preview.ReceiptWriterAuthorizationReady ||
		preview.ReceiptPersistenceThreatBoundaryReady ||
		preview.ConsumerEnablementReceiptPersistenceReady ||
		preview.ReceiptPersistenceAuthorizationReady ||
		preview.ReceiptPersistenceEnabled ||
		preview.ReceiptReplayEnabled ||
		preview.ReceiptExpiryWriteEnabled ||
		preview.ReceiptRevocationWriteEnabled ||
		preview.ReceiptAccepted ||
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
		preview.AuditDecision != "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-receipt-persistence-authorization-audit-blocked" {
		t.Fatalf("missing lookup consumer enablement receipt persistence authorization sources must fail closed: %#v", preview)
	}
	for _, item := range preview.AuthorizationItems {
		if item.EvidencePresent ||
			item.WriterAuthorizationAuditConsumed ||
			item.PersistenceThreatReviewConsumed ||
			item.PersistenceAuthorizationModeled ||
			item.ReceiptPersistenceAuthorizationReady ||
			item.ReceiptPersistenceEnabled ||
			item.ReceiptReplayEnabled ||
			item.ReceiptExpiryWriteEnabled ||
			item.ReceiptRevocationWriteEnabled ||
			item.ReceiptAccepted ||
			item.ConsumerEnablementAuthorized ||
			item.KDEConsumerEnabled ||
			item.RuntimeConsumerEnabled ||
			item.RawResultExposed ||
			item.UserVisible ||
			item.AuthorizationStatus != "missing-consumer-enablement-receipt-persistence-authorization-evidence" {
			t.Fatalf("missing lookup consumer enablement receipt persistence authorization item must remain closed: %#v", item)
		}
	}
}
