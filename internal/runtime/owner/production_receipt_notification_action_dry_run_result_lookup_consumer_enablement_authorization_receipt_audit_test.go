package owner

import (
	"os"
	"path/filepath"
	"testing"
)

func TestProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementAuthorizationReceiptAuditPreviewModelsReceipt(t *testing.T) {
	preview, err := NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementAuthorizationReceiptAuditPreview(projectRootForOwnerTest(t))
	if err != nil {
		t.Fatalf("NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementAuthorizationReceiptAuditPreview returned error: %v", err)
	}

	if preview.Version != currentProjectVersion(t) ||
		preview.SchemaVersion != "xnix.runtime.production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_authorization_receipt_audit.v1" ||
		preview.RequestType != "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-authorization-receipt-audit-preview" ||
		preview.AuditType != "receipt-notification-action-dry-run-result-lookup-consumer-enablement-authorization-receipt-audit" ||
		preview.AuditDecision != "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-authorization-receipt-audit-ready-receipt-disabled" ||
		preview.ReceiptSchema != "xnix.runtime.production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_authorization_receipt.v1" ||
		preview.OpaqueReceiptID != ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementAuthorizationReceiptID {
		t.Fatalf("unexpected lookup consumer enablement authorization receipt schema: %#v", preview)
	}
	if !preview.AuthorizationReceiptAuditRequired ||
		!preview.AuthorizationReceiptAuditModeled ||
		!preview.ConsumerEnablementGateAuditConsumed ||
		!preview.AuthorizationReceiptGuidanceConsumed ||
		!preview.AuthorizationReceiptBoundaryReady ||
		!preview.ConsumerEnablementGateReady ||
		!preview.OpaqueAuthorizationReceiptModeled ||
		!preview.ConsumerAuthorizationReady ||
		preview.ReceiptPresent ||
		preview.ReceiptAccepted ||
		preview.AuthorizationAccepted ||
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
		t.Fatalf("unexpected lookup consumer enablement authorization receipt decision: %#v", preview)
	}
	if preview.ReceiptItemCount != 5 ||
		preview.RequiredReceiptItemCount != 5 ||
		preview.ReadyReceiptItemCount != 5 ||
		preview.MissingReceiptItemCount != 0 ||
		preview.AcceptedReceiptItemCount != 0 ||
		preview.AuthorizedConsumerItemCount != 0 ||
		preview.EnabledConsumerItemCount != 0 ||
		preview.PersistedReceiptItemCount != 0 ||
		preview.RawExposedReceiptItemCount != 0 ||
		preview.SideEffectReceiptItemCount != 0 ||
		!sameStrings(preview.ReceiptItemIDs, []string{"review-receipt-dry-run-result-lookup-consumer-enablement-authorization-receipt", "renew-receipt-dry-run-result-lookup-consumer-enablement-authorization-receipt", "open-compatibility-center-dry-run-result-lookup-consumer-enablement-authorization-receipt", "dismiss-receipt-dry-run-result-lookup-consumer-enablement-authorization-receipt", "support-info-dry-run-result-lookup-consumer-enablement-authorization-receipt"}) {
		t.Fatalf("unexpected lookup consumer enablement authorization receipt inventory: %#v", preview)
	}
	for _, item := range preview.ReceiptItems {
		if !item.EvidencePresent ||
			!item.ConsumerEnablementGateConsumed ||
			!item.AuthorizationReceiptBoundaryModeled ||
			!item.OpaqueAuthorizationReceiptModeled ||
			item.ReceiptPresent ||
			item.ReceiptAccepted ||
			item.AuthorizationAccepted ||
			item.ConsumerConsumptionAuthorized ||
			item.ConsumerEnablementAuthorized ||
			item.KDEConsumerEnabled ||
			item.RuntimeConsumerEnabled ||
			item.LookupRouteAuthorized ||
			item.LookupRouteEnabled ||
			item.LookupRoutePersisted ||
			item.OpaqueLookupEnabled ||
			item.RedactedSummaryPersisted ||
			item.RawResultExposed ||
			item.RuntimeDiagnosticsPersisted ||
			item.DryRunResultPersisted ||
			item.DispatchDryRunExecuted ||
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
			item.ReceiptStatus != "consumer-enablement-authorization-receipt-modeled-receipt-disabled" {
			t.Fatalf("unsafe lookup consumer enablement authorization receipt item: %#v", item)
		}
	}
	if preview.Counts.Total != 9 ||
		preview.Counts.Passed != 9 ||
		preview.Counts.Pending != 0 ||
		preview.Counts.Blocked != 0 ||
		!sameStrings(preview.CheckIDs, []string{"consumer-enablement-gate-audit-consumed", "authorization-receipt-guidance-consumed", "authorization-receipt-boundary-modeled-only", "five-authorization-receipt-items-present", "authorization-receipt-items-ready-receipt-disabled", "receipt-acceptance-consumer-and-lookup-disabled", "request-notification-navigation-and-support-disabled", "raw-result-and-path-exposure-disabled", "production-and-host-boundary-closed"}) {
		t.Fatalf("unexpected lookup consumer enablement authorization receipt checks: counts=%#v ids=%#v", preview.Counts, preview.CheckIDs)
	}
	if preview.ReceiptPresent ||
		preview.ReceiptAccepted ||
		preview.AuthorizationAccepted ||
		preview.ConsumerEnablementAuthorized ||
		preview.KDEConsumerEnabled ||
		preview.RuntimeConsumerEnabled ||
		preview.LookupRouteEnabled ||
		preview.OpaqueLookupEnabled ||
		preview.RedactedSummaryPersisted ||
		preview.RawResultExposed ||
		preview.RuntimeDiagnosticsPersisted ||
		preview.DryRunResultPersisted ||
		preview.DispatchDryRunExecuted ||
		preview.RequestObjectCreationEnabled ||
		preview.PortalRequestCreated ||
		preview.NotificationActionEnabled ||
		preview.CompatibilityCenterOpened ||
		preview.SupportBundleExported ||
		preview.BackendLaunchEnabled ||
		preview.HostRootModified ||
		preview.BackendDetailsExposed {
		t.Fatalf("unsafe lookup consumer enablement authorization receipt gates: %#v", preview)
	}
	if err := validateNoBackendTerms(preview, "lookup consumer enablement authorization receipt test"); err != nil {
		t.Fatalf("validateNoBackendTerms returned error: %v", err)
	}
}

func TestProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementAuthorizationReceiptAuditPreviewFailsClosedWithoutSources(t *testing.T) {
	root := t.TempDir()
	if err := os.WriteFile(filepath.Join(root, "VERSION"), []byte(currentProjectVersion(t)+"\n"), 0o600); err != nil {
		t.Fatalf("WriteFile VERSION returned error: %v", err)
	}
	preview, err := NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementAuthorizationReceiptAuditPreview(root)
	if err != nil {
		t.Fatalf("NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementAuthorizationReceiptAuditPreview returned error: %v", err)
	}
	if !preview.AuthorizationReceiptAuditRequired ||
		!preview.AuthorizationReceiptAuditModeled ||
		preview.ConsumerEnablementGateAuditConsumed ||
		preview.AuthorizationReceiptGuidanceConsumed ||
		preview.AuthorizationReceiptBoundaryReady ||
		preview.ConsumerEnablementGateReady ||
		preview.OpaqueAuthorizationReceiptModeled ||
		preview.ConsumerAuthorizationReady ||
		preview.ReceiptPresent ||
		preview.ReceiptAccepted ||
		preview.ConsumerEnablementAuthorized ||
		preview.KDEConsumerEnabled ||
		preview.RuntimeConsumerEnabled ||
		preview.RawResultExposed ||
		preview.ReceiptItemCount != 5 ||
		preview.RequiredReceiptItemCount != 5 ||
		preview.ReadyReceiptItemCount != 0 ||
		preview.MissingReceiptItemCount != 5 ||
		preview.Counts.Total != 9 ||
		preview.Counts.Passed != 4 ||
		preview.Counts.Blocked != 5 ||
		preview.AuditDecision != "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-authorization-receipt-audit-blocked" {
		t.Fatalf("missing lookup consumer enablement authorization receipt sources must fail closed: %#v", preview)
	}
	for _, item := range preview.ReceiptItems {
		if item.EvidencePresent ||
			item.ConsumerEnablementGateConsumed ||
			item.AuthorizationReceiptBoundaryModeled ||
			item.OpaqueAuthorizationReceiptModeled ||
			item.ReceiptPresent ||
			item.ReceiptAccepted ||
			item.ConsumerEnablementAuthorized ||
			item.KDEConsumerEnabled ||
			item.RuntimeConsumerEnabled ||
			item.RawResultExposed ||
			item.UserVisible ||
			item.ReceiptStatus != "missing-consumer-enablement-authorization-receipt-evidence" {
			t.Fatalf("missing lookup consumer enablement authorization receipt item must remain closed: %#v", item)
		}
	}
}
