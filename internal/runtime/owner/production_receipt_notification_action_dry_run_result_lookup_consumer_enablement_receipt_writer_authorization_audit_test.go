package owner

import (
	"os"
	"path/filepath"
	"testing"
)

func TestProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptWriterAuthorizationAuditPreviewModelsWriter(t *testing.T) {
	preview, err := NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptWriterAuthorizationAuditPreview(projectRootForOwnerTest(t))
	if err != nil {
		t.Fatalf("NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptWriterAuthorizationAuditPreview returned error: %v", err)
	}

	if preview.Version != currentProjectVersion(t) ||
		preview.SchemaVersion != "xnix.runtime.production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_receipt_writer_authorization_audit.v1" ||
		preview.RequestType != "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-receipt-writer-authorization-audit-preview" ||
		preview.AuditType != "receipt-notification-action-dry-run-result-lookup-consumer-enablement-receipt-writer-authorization-audit" ||
		preview.AuditDecision != "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-receipt-writer-authorization-audit-ready-writes-disabled" ||
		preview.ReceiptSchema != "xnix.runtime.production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_authorization_receipt.v1" ||
		preview.OpaqueReceiptID != ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementAuthorizationReceiptID {
		t.Fatalf("unexpected lookup consumer enablement receipt writer authorization schema: %#v", preview)
	}
	if !preview.WriterAuthorizationAuditRequired ||
		!preview.WriterAuthorizationAuditModeled ||
		!preview.AuthorizationReceiptAuditConsumed ||
		!preview.BaseWriterAuthorizationReviewConsumed ||
		!preview.WriterAuthorizationGuidanceConsumed ||
		!preview.WriterAuthorizationBoundaryReady ||
		!preview.AuthorizationReceiptBoundaryReady ||
		!preview.ConsumerEnablementReceiptWriterReady ||
		!preview.ReceiptWriterAuthorizationReady ||
		preview.ReceiptPresent ||
		preview.ReceiptAccepted ||
		preview.AuthorizationAccepted ||
		preview.ReceiptWriterEnabled ||
		preview.ReceiptPersistenceEnabled ||
		preview.ReceiptLookupWritesEnabled ||
		preview.ReceiptReplayEnabled ||
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
		t.Fatalf("unexpected lookup consumer enablement receipt writer authorization decision: %#v", preview)
	}
	if preview.AuthorizationItemCount != 5 ||
		preview.RequiredAuthorizationItemCount != 5 ||
		preview.ReadyAuthorizationItemCount != 5 ||
		preview.MissingAuthorizationItemCount != 0 ||
		preview.WriteEnabledAuthorizationItemCount != 0 ||
		preview.PersistedAuthorizationItemCount != 0 ||
		preview.AcceptedAuthorizationItemCount != 0 ||
		preview.EnabledConsumerItemCount != 0 ||
		preview.RawExposedAuthorizationItemCount != 0 ||
		preview.SideEffectAuthorizationItemCount != 0 ||
		!sameStrings(preview.AuthorizationItemIDs, []string{"review-receipt-dry-run-result-lookup-consumer-enablement-writer-authorization", "renew-receipt-dry-run-result-lookup-consumer-enablement-writer-authorization", "open-compatibility-center-dry-run-result-lookup-consumer-enablement-writer-authorization", "dismiss-receipt-dry-run-result-lookup-consumer-enablement-writer-authorization", "support-info-dry-run-result-lookup-consumer-enablement-writer-authorization"}) {
		t.Fatalf("unexpected lookup consumer enablement receipt writer authorization inventory: %#v", preview)
	}
	for _, item := range preview.AuthorizationItems {
		if !item.EvidencePresent ||
			!item.AuthorizationReceiptAuditConsumed ||
			!item.WriterAuthorizationModeled ||
			!item.ReceiptWriterAuthorizationReady ||
			item.ReceiptPresent ||
			item.ReceiptAccepted ||
			item.AuthorizationAccepted ||
			item.ReceiptWriterEnabled ||
			item.ReceiptPersistenceEnabled ||
			item.ReceiptLookupWritesEnabled ||
			item.ReceiptReplayEnabled ||
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
			item.AuthorizationStatus != "consumer-enablement-receipt-writer-authorization-modeled-writes-disabled" {
			t.Fatalf("unsafe lookup consumer enablement receipt writer authorization item: %#v", item)
		}
	}
	if preview.Counts.Total != 9 ||
		preview.Counts.Passed != 9 ||
		preview.Counts.Pending != 0 ||
		preview.Counts.Blocked != 0 ||
		!sameStrings(preview.CheckIDs, []string{"authorization-receipt-audit-consumed", "base-writer-authorization-review-consumed", "writer-authorization-guidance-consumed", "writer-authorization-boundary-modeled-only", "five-writer-authorization-items-present", "writer-authorization-items-ready-writes-disabled", "receipt-writes-persistence-and-replay-disabled", "consumer-lookup-request-and-support-disabled", "production-and-host-boundary-closed"}) {
		t.Fatalf("unexpected lookup consumer enablement receipt writer authorization checks: counts=%#v ids=%#v", preview.Counts, preview.CheckIDs)
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
		preview.RequestObjectCreationEnabled ||
		preview.RequestObjectDispatchEnabled ||
		preview.PortalRequestCreated ||
		preview.NotificationActionEnabled ||
		preview.CompatibilityCenterOpened ||
		preview.SupportBundleExported ||
		preview.BackendLaunchEnabled ||
		preview.HostRootModified ||
		preview.BackendDetailsExposed {
		t.Fatalf("unsafe lookup consumer enablement receipt writer authorization gates: %#v", preview)
	}
	if err := validateNoBackendTerms(preview, "lookup consumer enablement receipt writer authorization test"); err != nil {
		t.Fatalf("validateNoBackendTerms returned error: %v", err)
	}
}

func TestProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptWriterAuthorizationAuditPreviewFailsClosedWithoutSources(t *testing.T) {
	root := t.TempDir()
	if err := os.WriteFile(filepath.Join(root, "VERSION"), []byte(currentProjectVersion(t)+"\n"), 0o600); err != nil {
		t.Fatalf("WriteFile VERSION returned error: %v", err)
	}
	preview, err := NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptWriterAuthorizationAuditPreview(root)
	if err != nil {
		t.Fatalf("NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptWriterAuthorizationAuditPreview returned error: %v", err)
	}
	if !preview.WriterAuthorizationAuditRequired ||
		!preview.WriterAuthorizationAuditModeled ||
		preview.AuthorizationReceiptAuditConsumed ||
		preview.BaseWriterAuthorizationReviewConsumed ||
		preview.WriterAuthorizationGuidanceConsumed ||
		preview.WriterAuthorizationBoundaryReady ||
		preview.AuthorizationReceiptBoundaryReady ||
		preview.ConsumerEnablementReceiptWriterReady ||
		preview.ReceiptWriterAuthorizationReady ||
		preview.ReceiptWriterEnabled ||
		preview.ReceiptPersistenceEnabled ||
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
		preview.AuditDecision != "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-receipt-writer-authorization-audit-blocked" {
		t.Fatalf("missing lookup consumer enablement receipt writer authorization sources must fail closed: %#v", preview)
	}
	for _, item := range preview.AuthorizationItems {
		if item.EvidencePresent ||
			item.AuthorizationReceiptAuditConsumed ||
			item.WriterAuthorizationModeled ||
			item.ReceiptWriterAuthorizationReady ||
			item.ReceiptWriterEnabled ||
			item.ReceiptPersistenceEnabled ||
			item.ReceiptReplayEnabled ||
			item.ReceiptAccepted ||
			item.ConsumerEnablementAuthorized ||
			item.KDEConsumerEnabled ||
			item.RuntimeConsumerEnabled ||
			item.RawResultExposed ||
			item.UserVisible ||
			item.AuthorizationStatus != "missing-consumer-enablement-receipt-writer-authorization-evidence" {
			t.Fatalf("missing lookup consumer enablement receipt writer authorization item must remain closed: %#v", item)
		}
	}
}
