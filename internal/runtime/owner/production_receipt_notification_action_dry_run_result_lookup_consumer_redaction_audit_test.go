package owner

import (
	"os"
	"path/filepath"
	"testing"
)

func TestProductionReceiptNotificationActionDryRunResultLookupConsumerRedactionAuditPreviewModelsConsumers(t *testing.T) {
	preview, err := NewProductionReceiptNotificationActionDryRunResultLookupConsumerRedactionAuditPreview(projectRootForOwnerTest(t))
	if err != nil {
		t.Fatalf("NewProductionReceiptNotificationActionDryRunResultLookupConsumerRedactionAuditPreview returned error: %v", err)
	}

	if preview.Version != currentProjectVersion(t) ||
		preview.SchemaVersion != "xnix.runtime.production_receipt_notification_action_dry_run_result_lookup_consumer_redaction_audit.v1" ||
		preview.RequestType != "production-receipt-notification-action-dry-run-result-lookup-consumer-redaction-audit-preview" ||
		preview.AuditType != "receipt-notification-action-dry-run-result-lookup-consumer-redaction-audit" ||
		preview.AuditDecision != "production-receipt-notification-action-dry-run-result-lookup-consumer-redaction-audit-ready-consumers-disabled" {
		t.Fatalf("unexpected lookup consumer redaction schema: %#v", preview)
	}
	if !preview.ConsumerRedactionAuditRequired ||
		!preview.ConsumerRedactionAuditModeled ||
		!preview.LookupRouteAuthorizationAuditConsumed ||
		!preview.ConsumerRedactionGuidanceConsumed ||
		!preview.ConsumerRedactionReady ||
		!preview.KDEConsumerRedactionModeled ||
		!preview.RuntimeConsumerRedactionModeled ||
		!preview.OpaqueResultIDSupported ||
		preview.ConsumerConsumptionAuthorized ||
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
		preview.ProductionReadiness ||
		preview.ProductionOwnershipReady {
		t.Fatalf("unexpected lookup consumer redaction decision: %#v", preview)
	}
	if preview.ConsumerItemCount != 5 ||
		preview.RequiredConsumerItemCount != 5 ||
		preview.ReadyConsumerItemCount != 5 ||
		preview.MissingConsumerItemCount != 0 ||
		preview.RedactedConsumerItemCount != 5 ||
		preview.EnabledConsumerItemCount != 0 ||
		preview.RawExposedConsumerItemCount != 0 ||
		preview.PersistedConsumerItemCount != 0 ||
		preview.SideEffectConsumerItemCount != 0 ||
		!sameStrings(preview.ConsumerItemIDs, []string{"review-receipt-dry-run-result-lookup-consumer-redaction", "renew-receipt-dry-run-result-lookup-consumer-redaction", "open-compatibility-center-dry-run-result-lookup-consumer-redaction", "dismiss-receipt-dry-run-result-lookup-consumer-redaction", "support-info-dry-run-result-lookup-consumer-redaction"}) {
		t.Fatalf("unexpected lookup consumer redaction inventory: %#v", preview)
	}
	for _, item := range preview.ConsumerItems {
		if !item.EvidencePresent ||
			!item.ConsumerRedactionModeled ||
			!item.KDEConsumerRedactionModeled ||
			!item.RuntimeConsumerRedactionModeled ||
			!item.OpaqueResultIDSupported ||
			item.ConsumerConsumptionAuthorized ||
			item.KDEConsumerEnabled ||
			item.RuntimeConsumerEnabled ||
			item.LookupRouteAuthorized ||
			item.LookupRouteEnabled ||
			item.LookupRoutePersisted ||
			item.OpaqueLookupEnabled ||
			item.OpaqueLookupPersisted ||
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
			item.ConsumerStatus != "lookup-consumer-redaction-modeled-consumers-disabled" {
			t.Fatalf("unsafe lookup consumer redaction item: %#v", item)
		}
	}
	if preview.Counts.Total != 9 ||
		preview.Counts.Passed != 9 ||
		preview.Counts.Pending != 0 ||
		preview.Counts.Blocked != 0 ||
		!sameStrings(preview.CheckIDs, []string{"lookup-route-authorization-audit-consumed", "consumer-redaction-guidance-consumed", "consumer-redaction-modeled-only", "five-consumer-redaction-items-present", "consumer-redaction-items-ready-consumers-disabled", "lookup-persistence-result-persistence-and-execution-disabled", "raw-result-and-path-exposure-disabled", "request-dispatch-notification-and-support-disabled", "production-and-host-boundary-closed"}) {
		t.Fatalf("unexpected lookup consumer redaction checks: counts=%#v ids=%#v", preview.Counts, preview.CheckIDs)
	}
	if preview.LookupRouteEnabled ||
		preview.KDEConsumerEnabled ||
		preview.RuntimeConsumerEnabled ||
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
		preview.BackendLaunchEnabled ||
		preview.HostRootModified ||
		preview.BackendDetailsExposed {
		t.Fatalf("unsafe lookup consumer redaction gates: %#v", preview)
	}
	if err := validateNoBackendTerms(preview, "lookup consumer redaction test"); err != nil {
		t.Fatalf("validateNoBackendTerms returned error: %v", err)
	}
}

func TestProductionReceiptNotificationActionDryRunResultLookupConsumerRedactionAuditPreviewFailsClosedWithoutSources(t *testing.T) {
	root := t.TempDir()
	if err := os.WriteFile(filepath.Join(root, "VERSION"), []byte(currentProjectVersion(t)+"\n"), 0o600); err != nil {
		t.Fatalf("WriteFile VERSION returned error: %v", err)
	}
	preview, err := NewProductionReceiptNotificationActionDryRunResultLookupConsumerRedactionAuditPreview(root)
	if err != nil {
		t.Fatalf("NewProductionReceiptNotificationActionDryRunResultLookupConsumerRedactionAuditPreview returned error: %v", err)
	}
	if preview.LookupRouteAuthorizationAuditConsumed ||
		preview.ConsumerRedactionGuidanceConsumed ||
		!preview.ConsumerRedactionAuditRequired ||
		!preview.ConsumerRedactionAuditModeled ||
		preview.ConsumerRedactionReady ||
		preview.KDEConsumerRedactionModeled ||
		preview.RuntimeConsumerRedactionModeled ||
		preview.OpaqueResultIDSupported ||
		preview.ConsumerConsumptionAuthorized ||
		preview.KDEConsumerEnabled ||
		preview.RuntimeConsumerEnabled ||
		preview.RawResultExposed ||
		preview.ConsumerItemCount != 5 ||
		preview.RequiredConsumerItemCount != 5 ||
		preview.ReadyConsumerItemCount != 0 ||
		preview.MissingConsumerItemCount != 5 ||
		preview.RedactedConsumerItemCount != 0 ||
		preview.Counts.Total != 9 ||
		preview.Counts.Passed != 4 ||
		preview.Counts.Blocked != 5 ||
		preview.AuditDecision != "production-receipt-notification-action-dry-run-result-lookup-consumer-redaction-audit-blocked" {
		t.Fatalf("missing lookup consumer redaction sources must fail closed: %#v", preview)
	}
	for _, item := range preview.ConsumerItems {
		if item.EvidencePresent ||
			item.ConsumerRedactionModeled ||
			item.KDEConsumerRedactionModeled ||
			item.RuntimeConsumerRedactionModeled ||
			item.OpaqueResultIDSupported ||
			item.ConsumerConsumptionAuthorized ||
			item.KDEConsumerEnabled ||
			item.RuntimeConsumerEnabled ||
			item.RawResultExposed ||
			item.UserVisible ||
			item.ConsumerStatus != "missing-lookup-consumer-redaction-evidence" {
			t.Fatalf("missing lookup consumer redaction item must remain closed: %#v", item)
		}
	}
}
