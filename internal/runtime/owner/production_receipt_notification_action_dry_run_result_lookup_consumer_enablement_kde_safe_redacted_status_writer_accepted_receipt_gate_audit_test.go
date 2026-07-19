package owner

import (
	"os"
	"path/filepath"
	"testing"
)

func TestProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAcceptedReceiptGateAuditPreviewModelsGate(t *testing.T) {
	preview, err := NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAcceptedReceiptGateAuditPreview(projectRootForOwnerTest(t))
	if err != nil {
		t.Fatalf("NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAcceptedReceiptGateAuditPreview returned error: %v", err)
	}

	if preview.Version != currentProjectVersion(t) ||
		preview.SchemaVersion != "xnix.runtime.production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_kde_safe_redacted_status_writer_accepted_receipt_gate_audit.v1" ||
		preview.RequestType != "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-writer-accepted-receipt-gate-audit-preview" ||
		preview.AuditType != "receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-writer-accepted-receipt-gate-audit" ||
		preview.AuditDecision != "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-writer-accepted-receipt-gate-audit-ready-writer-disabled" {
		t.Fatalf("unexpected KDE-safe redacted status writer accepted receipt gate schema: %#v", preview)
	}
	if !preview.CurrentMainlineConsumed ||
		!preview.WriterAuthorizationReceiptAcceptanceAuditConsumed ||
		!preview.WriterAuthorizationReceiptAcceptanceReady ||
		!preview.AcceptedReceiptGateRequired ||
		!preview.AcceptedReceiptGateModeled ||
		!preview.AcceptedReceiptGateReady ||
		!preview.AcceptedReceiptGateBoundaryReady ||
		!preview.OpaqueWriterAuthorizationReceipt ||
		!preview.KDESafeRedactedStatusOnly ||
		!preview.CompatibilityCenterGateModeled ||
		!preview.RuntimeDiagnosticsGateModeled ||
		preview.ReceiptPresent ||
		preview.ReceiptAccepted ||
		preview.AuthorizationAccepted ||
		preview.WriterAuthorizationGranted ||
		preview.StatusWriterEnabled ||
		preview.StatusPersistenceWriteEnabled ||
		preview.KDEStatusWriteEnabled ||
		preview.RuntimeDiagnosticsWriteEnabled ||
		preview.ConsumerConsumptionAuthorized ||
		preview.ConsumerEnablementAuthorized ||
		preview.KDEConsumerEnabled ||
		preview.RuntimeConsumerEnabled ||
		preview.LookupRouteAuthorized ||
		preview.LookupRouteEnabled ||
		preview.OpaqueLookupEnabled ||
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
		preview.CompatibilityCenterOpened ||
		preview.SupportBundleExported ||
		preview.SupportCaseCreated ||
		preview.ProductionReadiness ||
		preview.ProductionOwnershipReady {
		t.Fatalf("unsafe KDE-safe redacted status writer accepted receipt gate decision: %#v", preview)
	}
	if preview.GateItemCount != 10 ||
		preview.RequiredGateItemCount != 10 ||
		preview.ReadyGateItemCount != 10 ||
		preview.MissingGateItemCount != 0 ||
		preview.AcceptedReceiptItemCount != 0 ||
		preview.GrantedWriterItemCount != 0 ||
		preview.EnabledWriterItemCount != 0 ||
		preview.PersistedWriterItemCount != 0 ||
		preview.RawExposedGateItemCount != 0 ||
		preview.SideEffectGateItemCount != 0 ||
		preview.CompatibilityCenterGateItemCount != 5 ||
		preview.RuntimeDiagnosticsGateItemCount != 5 {
		t.Fatalf("unexpected KDE-safe redacted status writer accepted receipt gate counts: %#v", preview)
	}
	if len(preview.GateItems) != 10 ||
		!sameStrings(preview.GateItemIDs, []string{
			"review-receipt-dry-run-result-lookup-consumer-enablement-kde-safe-compatibility-center-redacted-status-writer-accepted-receipt-gate",
			"review-receipt-dry-run-result-lookup-consumer-enablement-kde-safe-runtime-diagnostics-redacted-status-writer-accepted-receipt-gate",
			"renew-receipt-dry-run-result-lookup-consumer-enablement-kde-safe-compatibility-center-redacted-status-writer-accepted-receipt-gate",
			"renew-receipt-dry-run-result-lookup-consumer-enablement-kde-safe-runtime-diagnostics-redacted-status-writer-accepted-receipt-gate",
			"open-compatibility-center-dry-run-result-lookup-consumer-enablement-kde-safe-compatibility-center-redacted-status-writer-accepted-receipt-gate",
			"open-compatibility-center-dry-run-result-lookup-consumer-enablement-kde-safe-runtime-diagnostics-redacted-status-writer-accepted-receipt-gate",
			"dismiss-receipt-dry-run-result-lookup-consumer-enablement-kde-safe-compatibility-center-redacted-status-writer-accepted-receipt-gate",
			"dismiss-receipt-dry-run-result-lookup-consumer-enablement-kde-safe-runtime-diagnostics-redacted-status-writer-accepted-receipt-gate",
			"support-info-dry-run-result-lookup-consumer-enablement-kde-safe-compatibility-center-redacted-status-writer-accepted-receipt-gate",
			"support-info-dry-run-result-lookup-consumer-enablement-kde-safe-runtime-diagnostics-redacted-status-writer-accepted-receipt-gate",
		}) {
		t.Fatalf("unexpected KDE-safe redacted status writer accepted receipt gate items: %#v", preview.GateItemIDs)
	}
	for _, item := range preview.GateItems {
		if !item.EvidencePresent ||
			!item.CurrentMainlineConsumed ||
			!item.WriterAuthorizationReceiptAcceptanceConsumed ||
			!item.AcceptedReceiptGateModeled ||
			!item.AcceptedReceiptGateBoundaryReady ||
			!item.OpaqueWriterAuthorizationReceipt ||
			!item.KDESafeRedactedStatusOnly ||
			item.ReceiptPresent ||
			item.ReceiptAccepted ||
			item.AuthorizationAccepted ||
			item.WriterAuthorizationGranted ||
			item.StatusWriterEnabled ||
			item.StatusPersistenceWriteEnabled ||
			item.KDEStatusWriteEnabled ||
			item.RuntimeDiagnosticsWriteEnabled ||
			item.RedactedSummaryPersisted ||
			item.KDEStatusPersisted ||
			item.RuntimeDiagnosticsPersisted ||
			item.RawResultExposed ||
			!item.UserVisible ||
			!item.ReviewOnly ||
			!item.RuntimeOwned ||
			!item.GoRuntimeBacked ||
			item.KDEPolicyOwner ||
			!item.SideEffectsDisabled ||
			item.HostRootModified ||
			item.InternalDetailsExposed ||
			item.AcceptedReceiptGateStatus != "redacted-status-writer-accepted-receipt-gate-modeled-writer-disabled" {
			t.Fatalf("unsafe KDE-safe redacted status writer accepted receipt gate item: %#v", item)
		}
	}
	if preview.Counts.Total != 8 ||
		preview.Counts.Passed != 8 ||
		preview.Counts.Pending != 0 ||
		preview.Counts.Blocked != 0 ||
		!sameStrings(preview.CheckIDs, []string{"current-mainline-consumed", "writer-authorization-receipt-acceptance-audit-consumed", "accepted-receipt-gate-modeled", "compatibility-center-and-runtime-gates-modeled", "ten-gate-items-ready-writer-disabled", "receipt-writers-and-persistence-disabled", "consumer-lookup-request-and-support-disabled", "production-and-host-boundary-closed"}) {
		t.Fatalf("unexpected KDE-safe redacted status writer accepted receipt gate checks: counts=%#v ids=%#v", preview.Counts, preview.CheckIDs)
	}
	if preview.SystemServiceStarted ||
		preview.SessionBusClaimed ||
		preview.ProductionBusClaimed ||
		preview.ProductionOwnerEnabled ||
		preview.WriteMethodsEnabled ||
		preview.RuntimeWritesEnabled ||
		preview.StatusWriterEnabled ||
		preview.StatusPersistenceWriteEnabled ||
		preview.KDEStatusWriteEnabled ||
		preview.RuntimeDiagnosticsWriteEnabled ||
		preview.RequestObjectCreationEnabled ||
		preview.RequestObjectDispatchEnabled ||
		preview.PortalRequestCreated ||
		preview.NotificationActionEnabled ||
		preview.CompatibilityCenterOpened ||
		preview.SupportBundleExported ||
		preview.BackendLaunchEnabled ||
		preview.HostRootModified ||
		preview.BackendDetailsExposed {
		t.Fatalf("unsafe KDE-safe redacted status writer accepted receipt gate gates: %#v", preview)
	}
	if err := validateNoBackendTerms(preview, "KDE-safe redacted status writer accepted receipt gate test"); err != nil {
		t.Fatalf("validateNoBackendTerms returned error: %v", err)
	}
}

func TestProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAcceptedReceiptGateAuditPreviewFailsClosedWithoutSources(t *testing.T) {
	root := t.TempDir()
	if err := os.WriteFile(filepath.Join(root, "VERSION"), []byte(currentProjectVersion(t)+"\n"), 0o600); err != nil {
		t.Fatalf("WriteFile VERSION returned error: %v", err)
	}
	preview, err := NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAcceptedReceiptGateAuditPreview(root)
	if err != nil {
		t.Fatalf("NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAcceptedReceiptGateAuditPreview returned error: %v", err)
	}
	if preview.CurrentMainlineConsumed ||
		preview.WriterAuthorizationReceiptAcceptanceAuditConsumed ||
		preview.WriterAuthorizationReceiptAcceptanceReady ||
		!preview.AcceptedReceiptGateRequired ||
		!preview.AcceptedReceiptGateModeled ||
		preview.AcceptedReceiptGateReady ||
		preview.AcceptedReceiptGateBoundaryReady ||
		preview.OpaqueWriterAuthorizationReceipt ||
		preview.KDESafeRedactedStatusOnly ||
		preview.CompatibilityCenterGateModeled ||
		preview.RuntimeDiagnosticsGateModeled ||
		preview.ReceiptPresent ||
		preview.ReceiptAccepted ||
		preview.AuthorizationAccepted ||
		preview.WriterAuthorizationGranted ||
		preview.StatusWriterEnabled ||
		preview.RawResultExposed ||
		preview.GateItemCount != 10 ||
		preview.RequiredGateItemCount != 10 ||
		preview.ReadyGateItemCount != 0 ||
		preview.MissingGateItemCount != 10 ||
		preview.Counts.Total != 8 ||
		preview.Counts.Passed != 3 ||
		preview.Counts.Blocked != 5 ||
		preview.AuditDecision != "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-writer-accepted-receipt-gate-audit-blocked" {
		t.Fatalf("missing KDE-safe redacted status writer accepted receipt gate sources must fail closed: %#v", preview)
	}
	for _, item := range preview.GateItems {
		if item.EvidencePresent ||
			item.CurrentMainlineConsumed ||
			item.WriterAuthorizationReceiptAcceptanceConsumed ||
			item.AcceptedReceiptGateModeled ||
			item.AcceptedReceiptGateBoundaryReady ||
			item.OpaqueWriterAuthorizationReceipt ||
			item.ReceiptAccepted ||
			item.AuthorizationAccepted ||
			item.WriterAuthorizationGranted ||
			item.StatusWriterEnabled ||
			item.RawResultExposed ||
			item.UserVisible ||
			item.AcceptedReceiptGateStatus != "missing-redacted-status-writer-accepted-receipt-gate-evidence" {
			t.Fatalf("missing KDE-safe redacted status writer accepted receipt gate item must remain closed: %#v", item)
		}
	}
}
