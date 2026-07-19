package owner

import (
	"os"
	"path/filepath"
	"testing"
)

func TestProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAuthorizationReceiptAcceptanceAuditPreviewModelsAcceptance(t *testing.T) {
	preview, err := NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAuthorizationReceiptAcceptanceAuditPreview(projectRootForOwnerTest(t))
	if err != nil {
		t.Fatalf("NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAuthorizationReceiptAcceptanceAuditPreview returned error: %v", err)
	}

	if preview.Version != currentProjectVersion(t) ||
		preview.SchemaVersion != "xnix.runtime.production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_kde_safe_redacted_status_writer_authorization_receipt_acceptance_audit.v1" ||
		preview.RequestType != "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-writer-authorization-receipt-acceptance-audit-preview" ||
		preview.AuditType != "receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-writer-authorization-receipt-acceptance-audit" ||
		preview.AuditDecision != "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-writer-authorization-receipt-acceptance-audit-ready-acceptance-disabled" {
		t.Fatalf("unexpected KDE-safe redacted status writer authorization receipt acceptance schema: %#v", preview)
	}
	if !preview.CurrentMainlineConsumed ||
		!preview.WriterAuthorizationReceiptAuditConsumed ||
		!preview.WriterAuthorizationReceiptReady ||
		!preview.AcceptanceAuthorizationRequired ||
		!preview.AcceptanceAuthorizationModeled ||
		!preview.AcceptanceAuthorizationReady ||
		!preview.ReceiptAcceptanceBoundaryReady ||
		!preview.OpaqueWriterAuthorizationReceipt ||
		!preview.KDESafeRedactedStatusOnly ||
		!preview.CompatibilityCenterAcceptanceModeled ||
		!preview.RuntimeDiagnosticsAcceptanceModeled ||
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
		t.Fatalf("unsafe KDE-safe redacted status writer authorization receipt acceptance decision: %#v", preview)
	}
	if preview.AcceptanceItemCount != 10 ||
		preview.RequiredAcceptanceItemCount != 10 ||
		preview.ReadyAcceptanceItemCount != 10 ||
		preview.MissingAcceptanceItemCount != 0 ||
		preview.AcceptedReceiptItemCount != 0 ||
		preview.GrantedWriterItemCount != 0 ||
		preview.EnabledWriterItemCount != 0 ||
		preview.PersistedWriterItemCount != 0 ||
		preview.RawExposedAcceptanceItemCount != 0 ||
		preview.SideEffectAcceptanceItemCount != 0 ||
		preview.CompatibilityCenterAcceptanceItemCount != 5 ||
		preview.RuntimeDiagnosticsAcceptanceItemCount != 5 {
		t.Fatalf("unexpected KDE-safe redacted status writer authorization receipt acceptance counts: %#v", preview)
	}
	if len(preview.AcceptanceItems) != 10 ||
		!sameStrings(preview.AcceptanceItemIDs, []string{
			"review-receipt-dry-run-result-lookup-consumer-enablement-kde-safe-compatibility-center-redacted-status-writer-authorization-receipt-acceptance",
			"review-receipt-dry-run-result-lookup-consumer-enablement-kde-safe-runtime-diagnostics-redacted-status-writer-authorization-receipt-acceptance",
			"renew-receipt-dry-run-result-lookup-consumer-enablement-kde-safe-compatibility-center-redacted-status-writer-authorization-receipt-acceptance",
			"renew-receipt-dry-run-result-lookup-consumer-enablement-kde-safe-runtime-diagnostics-redacted-status-writer-authorization-receipt-acceptance",
			"open-compatibility-center-dry-run-result-lookup-consumer-enablement-kde-safe-compatibility-center-redacted-status-writer-authorization-receipt-acceptance",
			"open-compatibility-center-dry-run-result-lookup-consumer-enablement-kde-safe-runtime-diagnostics-redacted-status-writer-authorization-receipt-acceptance",
			"dismiss-receipt-dry-run-result-lookup-consumer-enablement-kde-safe-compatibility-center-redacted-status-writer-authorization-receipt-acceptance",
			"dismiss-receipt-dry-run-result-lookup-consumer-enablement-kde-safe-runtime-diagnostics-redacted-status-writer-authorization-receipt-acceptance",
			"support-info-dry-run-result-lookup-consumer-enablement-kde-safe-compatibility-center-redacted-status-writer-authorization-receipt-acceptance",
			"support-info-dry-run-result-lookup-consumer-enablement-kde-safe-runtime-diagnostics-redacted-status-writer-authorization-receipt-acceptance",
		}) {
		t.Fatalf("unexpected KDE-safe redacted status writer authorization receipt acceptance items: %#v", preview.AcceptanceItemIDs)
	}
	for _, item := range preview.AcceptanceItems {
		if !item.EvidencePresent ||
			!item.CurrentMainlineConsumed ||
			!item.WriterAuthorizationReceiptConsumed ||
			!item.AcceptanceAuthorizationModeled ||
			!item.ReceiptAcceptanceBoundaryReady ||
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
			item.AcceptanceAuthorizationStatus != "redacted-status-writer-authorization-receipt-acceptance-modeled-acceptance-disabled" {
			t.Fatalf("unsafe KDE-safe redacted status writer authorization receipt acceptance item: %#v", item)
		}
	}
	if preview.Counts.Total != 8 ||
		preview.Counts.Passed != 8 ||
		preview.Counts.Pending != 0 ||
		preview.Counts.Blocked != 0 ||
		!sameStrings(preview.CheckIDs, []string{"current-mainline-consumed", "writer-authorization-receipt-audit-consumed", "acceptance-authorization-boundary-modeled", "compatibility-center-and-runtime-acceptance-modeled", "ten-acceptance-items-ready-acceptance-disabled", "receipt-acceptance-writers-and-persistence-disabled", "consumer-lookup-request-and-support-disabled", "production-and-host-boundary-closed"}) {
		t.Fatalf("unexpected KDE-safe redacted status writer authorization receipt acceptance checks: counts=%#v ids=%#v", preview.Counts, preview.CheckIDs)
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
		t.Fatalf("unsafe KDE-safe redacted status writer authorization receipt acceptance gates: %#v", preview)
	}
	if err := validateNoBackendTerms(preview, "KDE-safe redacted status writer authorization receipt acceptance test"); err != nil {
		t.Fatalf("validateNoBackendTerms returned error: %v", err)
	}
}

func TestProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAuthorizationReceiptAcceptanceAuditPreviewFailsClosedWithoutSources(t *testing.T) {
	root := t.TempDir()
	if err := os.WriteFile(filepath.Join(root, "VERSION"), []byte(currentProjectVersion(t)+"\n"), 0o600); err != nil {
		t.Fatalf("WriteFile VERSION returned error: %v", err)
	}
	preview, err := NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAuthorizationReceiptAcceptanceAuditPreview(root)
	if err != nil {
		t.Fatalf("NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAuthorizationReceiptAcceptanceAuditPreview returned error: %v", err)
	}
	if preview.CurrentMainlineConsumed ||
		preview.WriterAuthorizationReceiptAuditConsumed ||
		preview.WriterAuthorizationReceiptReady ||
		!preview.AcceptanceAuthorizationRequired ||
		!preview.AcceptanceAuthorizationModeled ||
		preview.AcceptanceAuthorizationReady ||
		preview.ReceiptAcceptanceBoundaryReady ||
		preview.OpaqueWriterAuthorizationReceipt ||
		preview.KDESafeRedactedStatusOnly ||
		preview.CompatibilityCenterAcceptanceModeled ||
		preview.RuntimeDiagnosticsAcceptanceModeled ||
		preview.ReceiptPresent ||
		preview.ReceiptAccepted ||
		preview.AuthorizationAccepted ||
		preview.WriterAuthorizationGranted ||
		preview.StatusWriterEnabled ||
		preview.RawResultExposed ||
		preview.AcceptanceItemCount != 10 ||
		preview.RequiredAcceptanceItemCount != 10 ||
		preview.ReadyAcceptanceItemCount != 0 ||
		preview.MissingAcceptanceItemCount != 10 ||
		preview.Counts.Total != 8 ||
		preview.Counts.Passed != 3 ||
		preview.Counts.Blocked != 5 ||
		preview.AuditDecision != "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-writer-authorization-receipt-acceptance-audit-blocked" {
		t.Fatalf("missing KDE-safe redacted status writer authorization receipt acceptance sources must fail closed: %#v", preview)
	}
	for _, item := range preview.AcceptanceItems {
		if item.EvidencePresent ||
			item.CurrentMainlineConsumed ||
			item.WriterAuthorizationReceiptConsumed ||
			item.AcceptanceAuthorizationModeled ||
			item.ReceiptAcceptanceBoundaryReady ||
			item.OpaqueWriterAuthorizationReceipt ||
			item.ReceiptAccepted ||
			item.AuthorizationAccepted ||
			item.WriterAuthorizationGranted ||
			item.StatusWriterEnabled ||
			item.RawResultExposed ||
			item.UserVisible ||
			item.AcceptanceAuthorizationStatus != "missing-redacted-status-writer-authorization-receipt-acceptance-evidence" {
			t.Fatalf("missing KDE-safe redacted status writer authorization receipt acceptance item must remain closed: %#v", item)
		}
	}
}
