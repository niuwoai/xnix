package owner

import (
	"os"
	"path/filepath"
	"testing"
)

func TestProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterEnablementAuditPreviewModelsEnablement(t *testing.T) {
	preview, err := NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterEnablementAuditPreview(projectRootForOwnerTest(t))
	if err != nil {
		t.Fatalf("NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterEnablementAuditPreview returned error: %v", err)
	}

	if preview.Version != currentProjectVersion(t) ||
		preview.SchemaVersion != "xnix.runtime.production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_kde_safe_redacted_status_writer_enablement_audit.v1" ||
		preview.RequestType != "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-writer-enablement-audit-preview" ||
		preview.AuditType != "receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-writer-enablement-audit" ||
		preview.AuditDecision != "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-writer-enablement-audit-ready-writer-disabled" {
		t.Fatalf("unexpected KDE-safe redacted status writer enablement schema: %#v", preview)
	}
	if !preview.CurrentMainlineConsumed ||
		!preview.WriterGrantAuditConsumed ||
		!preview.WriterGrantReady ||
		!preview.WriterEnablementRequired ||
		!preview.WriterEnablementModeled ||
		!preview.WriterEnablementBoundaryReady ||
		!preview.GrantConsumptionModeled ||
		!preview.KDESafeRedactedStatusOnly ||
		!preview.CompatibilityCenterModeled ||
		!preview.RuntimeDiagnosticsModeled ||
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
		t.Fatalf("unsafe KDE-safe redacted status writer enablement decision: %#v", preview)
	}
	if preview.EnablementItemCount != 10 ||
		preview.RequiredEnablementItemCount != 10 ||
		preview.ReadyEnablementItemCount != 10 ||
		preview.MissingEnablementItemCount != 0 ||
		preview.GrantConsumedItemCount != 0 ||
		preview.EnabledWriterItemCount != 0 ||
		preview.PersistedWriterItemCount != 0 ||
		preview.RawExposedEnablementItemCount != 0 ||
		preview.SideEffectEnablementItemCount != 0 ||
		preview.CompatibilityCenterItemCount != 5 ||
		preview.RuntimeDiagnosticsItemCount != 5 {
		t.Fatalf("unexpected KDE-safe redacted status writer enablement counts: %#v", preview)
	}
	if len(preview.EnablementItems) != 10 ||
		!sameStrings(preview.EnablementItemIDs, []string{
			"review-receipt-dry-run-result-lookup-consumer-enablement-kde-safe-compatibility-center-redacted-status-writer-enablement",
			"review-receipt-dry-run-result-lookup-consumer-enablement-kde-safe-runtime-diagnostics-redacted-status-writer-enablement",
			"renew-receipt-dry-run-result-lookup-consumer-enablement-kde-safe-compatibility-center-redacted-status-writer-enablement",
			"renew-receipt-dry-run-result-lookup-consumer-enablement-kde-safe-runtime-diagnostics-redacted-status-writer-enablement",
			"open-compatibility-center-dry-run-result-lookup-consumer-enablement-kde-safe-compatibility-center-redacted-status-writer-enablement",
			"open-compatibility-center-dry-run-result-lookup-consumer-enablement-kde-safe-runtime-diagnostics-redacted-status-writer-enablement",
			"dismiss-receipt-dry-run-result-lookup-consumer-enablement-kde-safe-compatibility-center-redacted-status-writer-enablement",
			"dismiss-receipt-dry-run-result-lookup-consumer-enablement-kde-safe-runtime-diagnostics-redacted-status-writer-enablement",
			"support-info-dry-run-result-lookup-consumer-enablement-kde-safe-compatibility-center-redacted-status-writer-enablement",
			"support-info-dry-run-result-lookup-consumer-enablement-kde-safe-runtime-diagnostics-redacted-status-writer-enablement",
		}) {
		t.Fatalf("unexpected KDE-safe redacted status writer enablement items: %#v", preview.EnablementItemIDs)
	}
	for _, item := range preview.EnablementItems {
		if !item.EvidencePresent ||
			!item.CurrentMainlineConsumed ||
			!item.WriterGrantAuditConsumed ||
			!item.WriterGrantReady ||
			!item.WriterEnablementModeled ||
			!item.WriterEnablementBoundaryReady ||
			!item.GrantConsumptionModeled ||
			!item.KDESafeRedactedStatusOnly ||
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
			item.WriterEnablementStatus != "redacted-status-writer-enablement-modeled-writer-disabled" {
			t.Fatalf("unsafe KDE-safe redacted status writer enablement item: %#v", item)
		}
	}
	if preview.Counts.Total != 8 ||
		preview.Counts.Passed != 8 ||
		preview.Counts.Pending != 0 ||
		preview.Counts.Blocked != 0 ||
		!sameStrings(preview.CheckIDs, []string{"current-mainline-consumed", "writer-grant-audit-consumed", "writer-enablement-boundary-modeled", "compatibility-center-and-runtime-enablement-modeled", "ten-enablement-items-ready-writer-disabled", "writer-enablement-writes-and-persistence-disabled", "consumer-lookup-request-and-support-disabled", "production-and-host-boundary-closed"}) {
		t.Fatalf("unexpected KDE-safe redacted status writer enablement checks: counts=%#v ids=%#v", preview.Counts, preview.CheckIDs)
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
		t.Fatalf("unsafe KDE-safe redacted status writer enablement gates: %#v", preview)
	}
	if err := validateNoBackendTerms(preview, "KDE-safe redacted status writer enablement test"); err != nil {
		t.Fatalf("validateNoBackendTerms returned error: %v", err)
	}
}

func TestProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterEnablementAuditPreviewFailsClosedWithoutSources(t *testing.T) {
	root := t.TempDir()
	if err := os.WriteFile(filepath.Join(root, "VERSION"), []byte(currentProjectVersion(t)+"\n"), 0o600); err != nil {
		t.Fatalf("WriteFile VERSION returned error: %v", err)
	}
	preview, err := NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterEnablementAuditPreview(root)
	if err != nil {
		t.Fatalf("NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterEnablementAuditPreview returned error: %v", err)
	}
	if preview.CurrentMainlineConsumed ||
		preview.WriterGrantAuditConsumed ||
		preview.WriterGrantReady ||
		!preview.WriterEnablementRequired ||
		!preview.WriterEnablementModeled ||
		preview.WriterEnablementBoundaryReady ||
		preview.GrantConsumptionModeled ||
		preview.KDESafeRedactedStatusOnly ||
		preview.CompatibilityCenterModeled ||
		preview.RuntimeDiagnosticsModeled ||
		preview.StatusWriterEnabled ||
		preview.RawResultExposed ||
		preview.EnablementItemCount != 10 ||
		preview.RequiredEnablementItemCount != 10 ||
		preview.ReadyEnablementItemCount != 0 ||
		preview.MissingEnablementItemCount != 10 ||
		preview.Counts.Total != 8 ||
		preview.Counts.Passed != 3 ||
		preview.Counts.Blocked != 5 ||
		preview.AuditDecision != "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-writer-enablement-audit-blocked" {
		t.Fatalf("missing KDE-safe redacted status writer enablement sources must fail closed: %#v", preview)
	}
	for _, item := range preview.EnablementItems {
		if item.EvidencePresent ||
			item.CurrentMainlineConsumed ||
			item.WriterGrantAuditConsumed ||
			item.WriterGrantReady ||
			item.WriterEnablementModeled ||
			item.WriterEnablementBoundaryReady ||
			item.GrantConsumptionModeled ||
			item.StatusWriterEnabled ||
			item.RawResultExposed ||
			item.UserVisible ||
			item.WriterEnablementStatus != "missing-redacted-status-writer-enablement-evidence" {
			t.Fatalf("missing KDE-safe redacted status writer enablement item must remain closed: %#v", item)
		}
	}
}
