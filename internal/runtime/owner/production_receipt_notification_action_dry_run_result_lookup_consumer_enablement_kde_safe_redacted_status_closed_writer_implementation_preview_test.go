package owner

import (
	"os"
	"path/filepath"
	"testing"
)

func TestProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedWriterImplementationPreviewModelsClosedWriter(t *testing.T) {
	preview, err := NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedWriterImplementationPreview(projectRootForOwnerTest(t))
	if err != nil {
		t.Fatalf("NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedWriterImplementationPreview returned error: %v", err)
	}

	if preview.Version != currentProjectVersion(t) ||
		preview.SchemaVersion != "xnix.runtime.production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_kde_safe_redacted_status_closed_writer_implementation.v1" ||
		preview.RequestType != "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-writer-implementation-preview" ||
		preview.PreviewType != "receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-writer-implementation" ||
		preview.PreviewDecision != "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-writer-implementation-ready-writer-disabled" {
		t.Fatalf("unexpected KDE-safe redacted status closed writer implementation schema: %#v", preview)
	}
	if !preview.CurrentMainlineConsumed ||
		!preview.WriterEnablementAuditConsumed ||
		!preview.WriterEnablementReady ||
		!preview.WriterImplementationRequired ||
		!preview.WriterImplementationModeled ||
		!preview.WriterImplementationReady ||
		!preview.ClosedWriterShapeModeled ||
		!preview.WriterInputBoundaryModeled ||
		!preview.WriterOutputBoundaryModeled ||
		!preview.KDESafeRedactedStatusOnly ||
		!preview.CompatibilityCenterModeled ||
		!preview.RuntimeDiagnosticsModeled ||
		preview.WriterCallable ||
		preview.WriterImplementationEnabled ||
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
		t.Fatalf("unsafe KDE-safe redacted status closed writer implementation decision: %#v", preview)
	}
	if preview.ImplementationItemCount != 10 ||
		preview.RequiredImplementationItemCount != 10 ||
		preview.ReadyImplementationItemCount != 10 ||
		preview.MissingImplementationItemCount != 0 ||
		preview.CallableWriterItemCount != 0 ||
		preview.EnabledWriterItemCount != 0 ||
		preview.PersistedWriterItemCount != 0 ||
		preview.RawExposedImplementationItemCount != 0 ||
		preview.SideEffectImplementationItemCount != 0 ||
		preview.CompatibilityCenterItemCount != 5 ||
		preview.RuntimeDiagnosticsItemCount != 5 {
		t.Fatalf("unexpected KDE-safe redacted status closed writer implementation counts: %#v", preview)
	}
	if len(preview.ImplementationItems) != 10 ||
		!sameStrings(preview.ImplementationItemIDs, []string{
			"review-receipt-dry-run-result-lookup-consumer-enablement-kde-safe-compatibility-center-redacted-status-closed-writer",
			"review-receipt-dry-run-result-lookup-consumer-enablement-kde-safe-runtime-diagnostics-redacted-status-closed-writer",
			"renew-receipt-dry-run-result-lookup-consumer-enablement-kde-safe-compatibility-center-redacted-status-closed-writer",
			"renew-receipt-dry-run-result-lookup-consumer-enablement-kde-safe-runtime-diagnostics-redacted-status-closed-writer",
			"open-compatibility-center-dry-run-result-lookup-consumer-enablement-kde-safe-compatibility-center-redacted-status-closed-writer",
			"open-compatibility-center-dry-run-result-lookup-consumer-enablement-kde-safe-runtime-diagnostics-redacted-status-closed-writer",
			"dismiss-receipt-dry-run-result-lookup-consumer-enablement-kde-safe-compatibility-center-redacted-status-closed-writer",
			"dismiss-receipt-dry-run-result-lookup-consumer-enablement-kde-safe-runtime-diagnostics-redacted-status-closed-writer",
			"support-info-dry-run-result-lookup-consumer-enablement-kde-safe-compatibility-center-redacted-status-closed-writer",
			"support-info-dry-run-result-lookup-consumer-enablement-kde-safe-runtime-diagnostics-redacted-status-closed-writer",
		}) {
		t.Fatalf("unexpected KDE-safe redacted status closed writer implementation items: %#v", preview.ImplementationItemIDs)
	}
	for _, item := range preview.ImplementationItems {
		if !item.EvidencePresent ||
			!item.CurrentMainlineConsumed ||
			!item.WriterEnablementAuditConsumed ||
			!item.WriterEnablementReady ||
			!item.ClosedWriterShapeModeled ||
			!item.WriterInputBoundaryModeled ||
			!item.WriterOutputBoundaryModeled ||
			!item.KDESafeRedactedStatusOnly ||
			item.WriterMethodName != "PreviewRedactedStatusWriter" ||
			item.WriterCallable ||
			item.WriterImplementationEnabled ||
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
			item.WriterImplementationStatus != "redacted-status-closed-writer-shape-modeled-call-disabled" {
			t.Fatalf("unsafe KDE-safe redacted status closed writer implementation item: %#v", item)
		}
	}
	if preview.Counts.Total != 8 ||
		preview.Counts.Passed != 8 ||
		preview.Counts.Pending != 0 ||
		preview.Counts.Blocked != 0 ||
		!sameStrings(preview.CheckIDs, []string{"current-mainline-consumed", "writer-enablement-audit-consumed", "closed-writer-shape-modeled", "writer-input-and-output-boundaries-modeled", "ten-implementation-items-ready-call-disabled", "writer-calls-writes-and-persistence-disabled", "consumer-lookup-request-and-support-disabled", "production-and-host-boundary-closed"}) {
		t.Fatalf("unexpected KDE-safe redacted status closed writer implementation checks: counts=%#v ids=%#v", preview.Counts, preview.CheckIDs)
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
		t.Fatalf("unsafe KDE-safe redacted status closed writer implementation gates: %#v", preview)
	}
	if err := validateNoBackendTerms(preview, "KDE-safe redacted status closed writer implementation test"); err != nil {
		t.Fatalf("validateNoBackendTerms returned error: %v", err)
	}
}

func TestProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedWriterImplementationPreviewFailsClosedWithoutSources(t *testing.T) {
	root := t.TempDir()
	if err := os.WriteFile(filepath.Join(root, "VERSION"), []byte(currentProjectVersion(t)+"\n"), 0o600); err != nil {
		t.Fatalf("WriteFile VERSION returned error: %v", err)
	}
	preview, err := NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedWriterImplementationPreview(root)
	if err != nil {
		t.Fatalf("NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedWriterImplementationPreview returned error: %v", err)
	}
	if preview.CurrentMainlineConsumed ||
		preview.WriterEnablementAuditConsumed ||
		preview.WriterEnablementReady ||
		!preview.WriterImplementationRequired ||
		!preview.WriterImplementationModeled ||
		preview.WriterImplementationReady ||
		preview.ClosedWriterShapeModeled ||
		preview.WriterInputBoundaryModeled ||
		preview.WriterOutputBoundaryModeled ||
		preview.KDESafeRedactedStatusOnly ||
		preview.CompatibilityCenterModeled ||
		preview.RuntimeDiagnosticsModeled ||
		preview.WriterCallable ||
		preview.StatusWriterEnabled ||
		preview.RawResultExposed ||
		preview.ImplementationItemCount != 10 ||
		preview.RequiredImplementationItemCount != 10 ||
		preview.ReadyImplementationItemCount != 0 ||
		preview.MissingImplementationItemCount != 10 ||
		preview.Counts.Total != 8 ||
		preview.Counts.Passed != 3 ||
		preview.Counts.Blocked != 5 ||
		preview.PreviewDecision != "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-writer-implementation-blocked" {
		t.Fatalf("missing KDE-safe redacted status closed writer implementation sources must fail closed: %#v", preview)
	}
	for _, item := range preview.ImplementationItems {
		if item.EvidencePresent ||
			item.CurrentMainlineConsumed ||
			item.WriterEnablementAuditConsumed ||
			item.WriterEnablementReady ||
			item.ClosedWriterShapeModeled ||
			item.WriterInputBoundaryModeled ||
			item.WriterOutputBoundaryModeled ||
			item.WriterCallable ||
			item.StatusWriterEnabled ||
			item.RawResultExposed ||
			item.UserVisible ||
			item.WriterImplementationStatus != "missing-redacted-status-closed-writer-implementation-evidence" {
			t.Fatalf("missing KDE-safe redacted status closed writer implementation item must remain closed: %#v", item)
		}
	}
}
