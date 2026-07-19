package owner

import (
	"os"
	"path/filepath"
	"testing"
)

func TestProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusFanOutAuditPreviewModelsStatusFanOut(t *testing.T) {
	preview, err := NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusFanOutAuditPreview(projectRootForOwnerTest(t))
	if err != nil {
		t.Fatalf("NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusFanOutAuditPreview returned error: %v", err)
	}

	if preview.Version != currentProjectVersion(t) ||
		preview.SchemaVersion != "xnix.runtime.production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_kde_safe_status_fanout_audit.v1" ||
		preview.RequestType != "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-status-fanout-audit-preview" ||
		preview.AuditType != "receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-status-fanout-audit" ||
		preview.AuditDecision != "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-status-fanout-audit-ready-status-only-consumers-disabled" {
		t.Fatalf("unexpected KDE-safe status fan-out schema: %#v", preview)
	}
	if !preview.StatusFanOutAuditRequired ||
		!preview.StatusFanOutModeled ||
		!preview.ConsumerEnablementGateAuditConsumed ||
		!preview.KDESafeStatusGuidanceConsumed ||
		!preview.CompatibilityCenterStatusModeled ||
		!preview.RuntimeDiagnosticsStatusModeled ||
		!preview.StatusFanOutReady ||
		!preview.ConsumerEnablementGateClosed ||
		!preview.ConsumerAuthorizationPrerequisiteSeen ||
		!preview.RouteAuthorizationPrerequisiteSeen ||
		!preview.ConsumerRedactionPrerequisiteSeen ||
		!preview.OpaqueResultIDSupported ||
		preview.ConsumerConsumptionAuthorized ||
		preview.ConsumerEnablementAuthorized ||
		preview.KDEConsumerEnabled ||
		preview.RuntimeConsumerEnabled ||
		preview.LookupRouteEnabled ||
		preview.OpaqueLookupEnabled ||
		preview.KDEStatusPersisted ||
		preview.RuntimeDiagnosticsPersisted ||
		preview.RawResultExposed ||
		preview.DispatchDryRunExecuted ||
		preview.ProductionReadiness ||
		preview.ProductionOwnershipReady {
		t.Fatalf("unexpected KDE-safe status fan-out decision: %#v", preview)
	}
	if preview.StatusItemCount != 10 ||
		preview.RequiredStatusItemCount != 10 ||
		preview.ReadyStatusItemCount != 10 ||
		preview.MissingStatusItemCount != 0 ||
		preview.CompatibilityCenterStatusItemCount != 5 ||
		preview.RuntimeDiagnosticsStatusItemCount != 5 ||
		preview.ConsumerEnabledStatusItemCount != 0 ||
		preview.PersistedStatusItemCount != 0 ||
		preview.RawExposedStatusItemCount != 0 ||
		preview.SideEffectStatusItemCount != 0 ||
		!sameStrings(preview.StatusItemIDs, []string{
			"review-receipt-dry-run-result-lookup-consumer-enablement-kde-safe-compatibility-center-status-fanout",
			"review-receipt-dry-run-result-lookup-consumer-enablement-kde-safe-runtime-diagnostics-status-fanout",
			"renew-receipt-dry-run-result-lookup-consumer-enablement-kde-safe-compatibility-center-status-fanout",
			"renew-receipt-dry-run-result-lookup-consumer-enablement-kde-safe-runtime-diagnostics-status-fanout",
			"open-compatibility-center-dry-run-result-lookup-consumer-enablement-kde-safe-compatibility-center-status-fanout",
			"open-compatibility-center-dry-run-result-lookup-consumer-enablement-kde-safe-runtime-diagnostics-status-fanout",
			"dismiss-receipt-dry-run-result-lookup-consumer-enablement-kde-safe-compatibility-center-status-fanout",
			"dismiss-receipt-dry-run-result-lookup-consumer-enablement-kde-safe-runtime-diagnostics-status-fanout",
			"support-info-dry-run-result-lookup-consumer-enablement-kde-safe-compatibility-center-status-fanout",
			"support-info-dry-run-result-lookup-consumer-enablement-kde-safe-runtime-diagnostics-status-fanout",
		}) {
		t.Fatalf("unexpected KDE-safe status fan-out inventory: %#v", preview)
	}
	for _, item := range preview.StatusItems {
		if !item.EvidencePresent ||
			!item.KDESafeStatus ||
			!item.ConsumerEnablementGateClosed ||
			!item.ConsumerEnablementGateAuditConsumed ||
			!item.OpaqueResultIDSupported ||
			!item.UserVisible ||
			!item.ReviewOnly ||
			!item.RuntimeOwned ||
			!item.GoRuntimeBacked ||
			item.KDEPolicyOwner ||
			item.ConsumerConsumptionAuthorized ||
			item.ConsumerEnablementAuthorized ||
			item.KDEConsumerEnabled ||
			item.RuntimeConsumerEnabled ||
			item.LookupRouteEnabled ||
			item.OpaqueLookupEnabled ||
			item.RedactedSummaryPersisted ||
			item.KDEStatusPersisted ||
			item.RuntimeDiagnosticsPersisted ||
			item.DryRunResultPersisted ||
			item.RawResultExposed ||
			item.DispatchDryRunExecuted ||
			item.RequestObjectCreated ||
			item.RequestObjectDispatched ||
			item.PortalRequestCreated ||
			item.NotificationActionEnabled ||
			item.CompatibilityCenterOpened ||
			item.SupportBundleExported ||
			item.SupportCaseCreated ||
			item.CallerStateRootRequired ||
			item.StateRootPathExposed ||
			item.FilePathsExposed ||
			item.FileContentRead ||
			!item.SideEffectsDisabled ||
			item.HostRootModified ||
			item.InternalDetailsExposed ||
			item.Status != "kde-safe-status-fanout-modeled-consumers-disabled" {
			t.Fatalf("unsafe KDE-safe status fan-out item: %#v", item)
		}
		if item.SurfaceKind == "compatibility-center" && (!item.CompatibilityCenterStatus || item.RuntimeDiagnosticsStatus) {
			t.Fatalf("Compatibility Center fan-out item must not masquerade as diagnostics: %#v", item)
		}
		if item.SurfaceKind == "runtime-diagnostics" && (!item.RuntimeDiagnosticsStatus || item.CompatibilityCenterStatus) {
			t.Fatalf("Runtime diagnostics fan-out item must not masquerade as Compatibility Center status: %#v", item)
		}
	}
	if preview.Counts.Total != 10 ||
		preview.Counts.Passed != 10 ||
		preview.Counts.Pending != 0 ||
		preview.Counts.Blocked != 0 ||
		!sameStrings(preview.CheckIDs, []string{"kde-safe-status-fanout-audit-required", "consumer-enablement-gate-audit-consumed", "kde-safe-status-guidance-consumed", "compatibility-center-status-modeled", "runtime-diagnostics-status-modeled", "ten-kde-safe-status-fanout-items-present", "status-fanout-modeled-only", "consumer-lookup-and-persistence-disabled", "request-notification-navigation-and-support-disabled", "production-and-host-boundary-closed"}) {
		t.Fatalf("unexpected KDE-safe status fan-out checks: counts=%#v ids=%#v", preview.Counts, preview.CheckIDs)
	}
	if preview.ConsumerEnablementAuthorized ||
		preview.KDEConsumerEnabled ||
		preview.RuntimeConsumerEnabled ||
		preview.LookupRouteEnabled ||
		preview.OpaqueLookupEnabled ||
		preview.KDEStatusPersisted ||
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
		t.Fatalf("unsafe KDE-safe status fan-out gates: %#v", preview)
	}
	if err := validateNoBackendTerms(preview, "KDE-safe status fan-out test"); err != nil {
		t.Fatalf("validateNoBackendTerms returned error: %v", err)
	}
}

func TestProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusFanOutAuditPreviewFailsClosedWithoutSources(t *testing.T) {
	root := t.TempDir()
	if err := os.WriteFile(filepath.Join(root, "VERSION"), []byte(currentProjectVersion(t)+"\n"), 0o600); err != nil {
		t.Fatalf("WriteFile VERSION returned error: %v", err)
	}
	preview, err := NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusFanOutAuditPreview(root)
	if err != nil {
		t.Fatalf("NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusFanOutAuditPreview returned error: %v", err)
	}
	if !preview.StatusFanOutAuditRequired ||
		!preview.StatusFanOutModeled ||
		preview.ConsumerEnablementGateAuditConsumed ||
		preview.KDESafeStatusGuidanceConsumed ||
		preview.CompatibilityCenterStatusModeled ||
		preview.RuntimeDiagnosticsStatusModeled ||
		preview.StatusFanOutReady ||
		preview.ConsumerEnablementGateClosed ||
		preview.OpaqueResultIDSupported ||
		preview.ConsumerConsumptionAuthorized ||
		preview.KDEConsumerEnabled ||
		preview.RuntimeConsumerEnabled ||
		preview.RawResultExposed ||
		preview.StatusItemCount != 10 ||
		preview.RequiredStatusItemCount != 10 ||
		preview.ReadyStatusItemCount != 0 ||
		preview.MissingStatusItemCount != 10 ||
		preview.CompatibilityCenterStatusItemCount != 5 ||
		preview.RuntimeDiagnosticsStatusItemCount != 5 ||
		preview.Counts.Total != 10 ||
		preview.Counts.Passed != 4 ||
		preview.Counts.Blocked != 6 ||
		preview.AuditDecision != "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-status-fanout-audit-blocked" {
		t.Fatalf("missing KDE-safe status fan-out sources must fail closed: %#v", preview)
	}
	for _, item := range preview.StatusItems {
		if item.EvidencePresent ||
			item.KDESafeStatus ||
			item.ConsumerEnablementGateClosed ||
			item.ConsumerEnablementGateAuditConsumed ||
			item.OpaqueResultIDSupported ||
			item.UserVisible ||
			item.KDEConsumerEnabled ||
			item.RuntimeConsumerEnabled ||
			item.RawResultExposed ||
			item.Status != "missing-kde-safe-status-fanout-evidence" {
			t.Fatalf("missing KDE-safe status fan-out item must remain closed: %#v", item)
		}
	}
}
