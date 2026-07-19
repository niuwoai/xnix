package owner

import (
	"os"
	"path/filepath"
	"testing"
)

func TestProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterGrantAuditPreviewModelsGrant(t *testing.T) {
	preview, err := NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterGrantAuditPreview(projectRootForOwnerTest(t))
	if err != nil {
		t.Fatalf("NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterGrantAuditPreview returned error: %v", err)
	}

	if preview.Version != currentProjectVersion(t) ||
		preview.SchemaVersion != "xnix.runtime.production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_kde_safe_redacted_status_writer_grant_audit.v1" ||
		preview.RequestType != "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-writer-grant-audit-preview" ||
		preview.AuditType != "receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-writer-grant-audit" ||
		preview.AuditDecision != "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-writer-grant-audit-ready-writer-disabled" {
		t.Fatalf("unexpected KDE-safe redacted status writer grant schema: %#v", preview)
	}
	if !preview.CurrentMainlineConsumed ||
		!preview.AcceptedReceiptGateAuditConsumed ||
		!preview.AcceptedReceiptGateReady ||
		!preview.WriterGrantRequired ||
		!preview.WriterGrantModeled ||
		!preview.WriterGrantReady ||
		!preview.WriterGrantBoundaryReady ||
		!preview.OpaqueWriterAuthorizationReceipt ||
		!preview.KDESafeRedactedStatusOnly ||
		!preview.CompatibilityCenterGrantModeled ||
		!preview.RuntimeDiagnosticsGrantModeled ||
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
		t.Fatalf("unsafe KDE-safe redacted status writer grant decision: %#v", preview)
	}
	if preview.GrantItemCount != 10 ||
		preview.RequiredGrantItemCount != 10 ||
		preview.ReadyGrantItemCount != 10 ||
		preview.MissingGrantItemCount != 0 ||
		preview.AcceptedReceiptItemCount != 0 ||
		preview.GrantedWriterItemCount != 0 ||
		preview.EnabledWriterItemCount != 0 ||
		preview.PersistedWriterItemCount != 0 ||
		preview.RawExposedGrantItemCount != 0 ||
		preview.SideEffectGrantItemCount != 0 ||
		preview.CompatibilityCenterGrantItemCount != 5 ||
		preview.RuntimeDiagnosticsGrantItemCount != 5 {
		t.Fatalf("unexpected KDE-safe redacted status writer grant counts: %#v", preview)
	}
	if len(preview.GrantItems) != 10 ||
		!sameStrings(preview.GrantItemIDs, []string{
			"review-receipt-dry-run-result-lookup-consumer-enablement-kde-safe-compatibility-center-redacted-status-writer-grant",
			"review-receipt-dry-run-result-lookup-consumer-enablement-kde-safe-runtime-diagnostics-redacted-status-writer-grant",
			"renew-receipt-dry-run-result-lookup-consumer-enablement-kde-safe-compatibility-center-redacted-status-writer-grant",
			"renew-receipt-dry-run-result-lookup-consumer-enablement-kde-safe-runtime-diagnostics-redacted-status-writer-grant",
			"open-compatibility-center-dry-run-result-lookup-consumer-enablement-kde-safe-compatibility-center-redacted-status-writer-grant",
			"open-compatibility-center-dry-run-result-lookup-consumer-enablement-kde-safe-runtime-diagnostics-redacted-status-writer-grant",
			"dismiss-receipt-dry-run-result-lookup-consumer-enablement-kde-safe-compatibility-center-redacted-status-writer-grant",
			"dismiss-receipt-dry-run-result-lookup-consumer-enablement-kde-safe-runtime-diagnostics-redacted-status-writer-grant",
			"support-info-dry-run-result-lookup-consumer-enablement-kde-safe-compatibility-center-redacted-status-writer-grant",
			"support-info-dry-run-result-lookup-consumer-enablement-kde-safe-runtime-diagnostics-redacted-status-writer-grant",
		}) {
		t.Fatalf("unexpected KDE-safe redacted status writer grant items: %#v", preview.GrantItemIDs)
	}
	for _, item := range preview.GrantItems {
		if !item.EvidencePresent ||
			!item.CurrentMainlineConsumed ||
			!item.AcceptedReceiptGateAuditConsumed ||
			!item.WriterGrantModeled ||
			!item.WriterGrantBoundaryReady ||
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
			item.WriterGrantStatus != "redacted-status-writer-grant-modeled-writer-disabled" {
			t.Fatalf("unsafe KDE-safe redacted status writer grant item: %#v", item)
		}
	}
	if preview.Counts.Total != 8 ||
		preview.Counts.Passed != 8 ||
		preview.Counts.Pending != 0 ||
		preview.Counts.Blocked != 0 ||
		!sameStrings(preview.CheckIDs, []string{"current-mainline-consumed", "accepted-receipt-gate-audit-consumed", "writer-grant-boundary-modeled", "compatibility-center-and-runtime-grants-modeled", "ten-grant-items-ready-writer-disabled", "writer-grants-writers-and-persistence-disabled", "consumer-lookup-request-and-support-disabled", "production-and-host-boundary-closed"}) {
		t.Fatalf("unexpected KDE-safe redacted status writer grant checks: counts=%#v ids=%#v", preview.Counts, preview.CheckIDs)
	}
	if preview.SystemServiceStarted ||
		preview.SessionBusClaimed ||
		preview.ProductionBusClaimed ||
		preview.ProductionOwnerEnabled ||
		preview.WriteMethodsEnabled ||
		preview.RuntimeWritesEnabled ||
		preview.BackendLaunchEnabled ||
		preview.HostRootModified ||
		preview.BackendDetailsExposed {
		t.Fatalf("unsafe KDE-safe redacted status writer grant gates: %#v", preview)
	}
	if err := validateNoBackendTerms(preview, "KDE-safe redacted status writer grant test"); err != nil {
		t.Fatalf("validateNoBackendTerms returned error: %v", err)
	}
}

func TestProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterGrantAuditPreviewFailsClosedWithoutSources(t *testing.T) {
	root := t.TempDir()
	if err := os.WriteFile(filepath.Join(root, "VERSION"), []byte(currentProjectVersion(t)+"\n"), 0o600); err != nil {
		t.Fatalf("WriteFile VERSION returned error: %v", err)
	}
	preview, err := NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterGrantAuditPreview(root)
	if err != nil {
		t.Fatalf("NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterGrantAuditPreview returned error: %v", err)
	}
	if preview.CurrentMainlineConsumed ||
		preview.AcceptedReceiptGateAuditConsumed ||
		preview.AcceptedReceiptGateReady ||
		!preview.WriterGrantRequired ||
		!preview.WriterGrantModeled ||
		preview.WriterGrantReady ||
		preview.WriterGrantBoundaryReady ||
		preview.OpaqueWriterAuthorizationReceipt ||
		preview.KDESafeRedactedStatusOnly ||
		preview.CompatibilityCenterGrantModeled ||
		preview.RuntimeDiagnosticsGrantModeled ||
		preview.ReceiptPresent ||
		preview.ReceiptAccepted ||
		preview.AuthorizationAccepted ||
		preview.WriterAuthorizationGranted ||
		preview.StatusWriterEnabled ||
		preview.RawResultExposed ||
		preview.GrantItemCount != 10 ||
		preview.RequiredGrantItemCount != 10 ||
		preview.ReadyGrantItemCount != 0 ||
		preview.MissingGrantItemCount != 10 ||
		preview.Counts.Total != 8 ||
		preview.Counts.Passed != 3 ||
		preview.Counts.Blocked != 5 ||
		preview.AuditDecision != "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-writer-grant-audit-blocked" {
		t.Fatalf("missing KDE-safe redacted status writer grant sources must fail closed: %#v", preview)
	}
	for _, item := range preview.GrantItems {
		if item.EvidencePresent ||
			item.CurrentMainlineConsumed ||
			item.AcceptedReceiptGateAuditConsumed ||
			item.WriterGrantModeled ||
			item.WriterGrantBoundaryReady ||
			item.OpaqueWriterAuthorizationReceipt ||
			item.ReceiptAccepted ||
			item.AuthorizationAccepted ||
			item.WriterAuthorizationGranted ||
			item.StatusWriterEnabled ||
			item.RawResultExposed ||
			item.UserVisible ||
			item.WriterGrantStatus != "missing-redacted-status-writer-grant-evidence" {
			t.Fatalf("missing KDE-safe redacted status writer grant item must remain closed: %#v", item)
		}
	}
}
