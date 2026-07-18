package owner

import (
	"os"
	"path/filepath"
	"testing"
)

func TestProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementGateAuditPreviewModelsGate(t *testing.T) {
	preview, err := NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementGateAuditPreview(projectRootForOwnerTest(t))
	if err != nil {
		t.Fatalf("NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementGateAuditPreview returned error: %v", err)
	}

	if preview.Version != currentProjectVersion(t) ||
		preview.SchemaVersion != "xnix.runtime.production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_gate_audit.v1" ||
		preview.RequestType != "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-gate-audit-preview" ||
		preview.AuditType != "receipt-notification-action-dry-run-result-lookup-consumer-enablement-gate-audit" ||
		preview.AuditDecision != "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-gate-audit-ready-consumers-disabled" {
		t.Fatalf("unexpected lookup consumer enablement gate schema: %#v", preview)
	}
	if !preview.ConsumerEnablementGateRequired ||
		!preview.ConsumerEnablementGateModeled ||
		!preview.LookupRouteAuthorizationAuditConsumed ||
		!preview.ConsumerRedactionAuditConsumed ||
		!preview.ConsumerEnablementGuidanceConsumed ||
		!preview.ConsumerEnablementGateReady ||
		!preview.RouteAuthorizationPrerequisiteModeled ||
		!preview.ConsumerRedactionPrerequisiteModeled ||
		!preview.OpaqueResultIDSupported ||
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
		preview.ProductionReadiness ||
		preview.ProductionOwnershipReady {
		t.Fatalf("unexpected lookup consumer enablement gate decision: %#v", preview)
	}
	if preview.GateItemCount != 5 ||
		preview.RequiredGateItemCount != 5 ||
		preview.ReadyGateItemCount != 5 ||
		preview.MissingGateItemCount != 0 ||
		preview.RouteReadyGateItemCount != 5 ||
		preview.RedactionReadyGateItemCount != 5 ||
		preview.AuthorizedGateItemCount != 0 ||
		preview.EnabledGateItemCount != 0 ||
		preview.PersistedGateItemCount != 0 ||
		preview.RawExposedGateItemCount != 0 ||
		preview.SideEffectGateItemCount != 0 ||
		!sameStrings(preview.GateItemIDs, []string{"review-receipt-dry-run-result-lookup-consumer-enablement-gate", "renew-receipt-dry-run-result-lookup-consumer-enablement-gate", "open-compatibility-center-dry-run-result-lookup-consumer-enablement-gate", "dismiss-receipt-dry-run-result-lookup-consumer-enablement-gate", "support-info-dry-run-result-lookup-consumer-enablement-gate"}) {
		t.Fatalf("unexpected lookup consumer enablement gate inventory: %#v", preview)
	}
	for _, item := range preview.GateItems {
		if !item.EvidencePresent ||
			!item.RouteAuthorizationPrerequisiteModeled ||
			!item.ConsumerRedactionPrerequisiteModeled ||
			!item.ConsumerEnablementGateModeled ||
			!item.OpaqueResultIDSupported ||
			item.ConsumerConsumptionAuthorized ||
			item.ConsumerEnablementAuthorized ||
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
			item.GateStatus != "lookup-consumer-enablement-gate-modeled-consumers-disabled" {
			t.Fatalf("unsafe lookup consumer enablement gate item: %#v", item)
		}
	}
	if preview.Counts.Total != 9 ||
		preview.Counts.Passed != 9 ||
		preview.Counts.Pending != 0 ||
		preview.Counts.Blocked != 0 ||
		!sameStrings(preview.CheckIDs, []string{"consumer-enablement-gate-audit-required", "lookup-route-authorization-audit-consumed", "consumer-redaction-audit-consumed", "consumer-enablement-guidance-consumed", "consumer-enablement-gate-modeled-only", "five-consumer-enablement-gate-items-present", "lookup-result-persistence-and-execution-disabled", "request-notification-navigation-and-support-disabled", "production-and-host-boundary-closed"}) {
		t.Fatalf("unexpected lookup consumer enablement gate checks: counts=%#v ids=%#v", preview.Counts, preview.CheckIDs)
	}
	if preview.ConsumerEnablementAuthorized ||
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
		preview.RequestObjectDispatchEnabled ||
		preview.PortalRequestCreated ||
		preview.NotificationActionEnabled ||
		preview.CompatibilityCenterOpened ||
		preview.SupportBundleExported ||
		preview.BackendLaunchEnabled ||
		preview.HostRootModified ||
		preview.BackendDetailsExposed {
		t.Fatalf("unsafe lookup consumer enablement gate gates: %#v", preview)
	}
	if err := validateNoBackendTerms(preview, "lookup consumer enablement gate test"); err != nil {
		t.Fatalf("validateNoBackendTerms returned error: %v", err)
	}
}

func TestProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementGateAuditPreviewFailsClosedWithoutSources(t *testing.T) {
	root := t.TempDir()
	if err := os.WriteFile(filepath.Join(root, "VERSION"), []byte(currentProjectVersion(t)+"\n"), 0o600); err != nil {
		t.Fatalf("WriteFile VERSION returned error: %v", err)
	}
	preview, err := NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementGateAuditPreview(root)
	if err != nil {
		t.Fatalf("NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementGateAuditPreview returned error: %v", err)
	}
	if !preview.ConsumerEnablementGateRequired ||
		!preview.ConsumerEnablementGateModeled ||
		preview.LookupRouteAuthorizationAuditConsumed ||
		preview.ConsumerRedactionAuditConsumed ||
		preview.ConsumerEnablementGuidanceConsumed ||
		preview.ConsumerEnablementGateReady ||
		preview.RouteAuthorizationPrerequisiteModeled ||
		preview.ConsumerRedactionPrerequisiteModeled ||
		preview.OpaqueResultIDSupported ||
		preview.ConsumerConsumptionAuthorized ||
		preview.ConsumerEnablementAuthorized ||
		preview.KDEConsumerEnabled ||
		preview.RuntimeConsumerEnabled ||
		preview.RawResultExposed ||
		preview.GateItemCount != 5 ||
		preview.RequiredGateItemCount != 5 ||
		preview.ReadyGateItemCount != 0 ||
		preview.MissingGateItemCount != 5 ||
		preview.RouteReadyGateItemCount != 0 ||
		preview.RedactionReadyGateItemCount != 0 ||
		preview.Counts.Total != 9 ||
		preview.Counts.Passed != 4 ||
		preview.Counts.Blocked != 5 ||
		preview.AuditDecision != "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-gate-audit-blocked" {
		t.Fatalf("missing lookup consumer enablement gate sources must fail closed: %#v", preview)
	}
	for _, item := range preview.GateItems {
		if item.EvidencePresent ||
			item.RouteAuthorizationPrerequisiteModeled ||
			item.ConsumerRedactionPrerequisiteModeled ||
			item.ConsumerEnablementGateModeled ||
			item.OpaqueResultIDSupported ||
			item.ConsumerConsumptionAuthorized ||
			item.ConsumerEnablementAuthorized ||
			item.KDEConsumerEnabled ||
			item.RuntimeConsumerEnabled ||
			item.RawResultExposed ||
			item.UserVisible ||
			item.GateStatus != "missing-lookup-consumer-enablement-gate-evidence" {
			t.Fatalf("missing lookup consumer enablement gate item must remain closed: %#v", item)
		}
	}
}
