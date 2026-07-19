package owner

import (
	"os"
	"path/filepath"
	"testing"
)

func TestProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterPersistenceAuthorizationPreviewModelsAuthorization(t *testing.T) {
	preview, err := NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterPersistenceAuthorizationPreview(projectRootForOwnerTest(t))
	if err != nil {
		t.Fatalf("NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterPersistenceAuthorizationPreview returned error: %v", err)
	}

	if preview.Version != currentProjectVersion(t) ||
		preview.SchemaVersion != "xnix.runtime.production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_kde_safe_redacted_status_writer_persistence_authorization.v1" ||
		preview.RequestType != "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-writer-persistence-authorization-preview" ||
		preview.PreviewType != "receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-writer-persistence-authorization" ||
		preview.PreviewDecision != "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-writer-persistence-authorization-ready-persistence-disabled" {
		t.Fatalf("unexpected KDE-safe redacted status writer persistence authorization schema: %#v", preview)
	}
	if !preview.CurrentMainlineConsumed ||
		!preview.ClosedWriterImplementationConsumed ||
		!preview.ClosedWriterImplementationReady ||
		!preview.PersistenceAuthorizationRequired ||
		!preview.PersistenceAuthorizationModeled ||
		!preview.PersistenceAuthorizationReady ||
		!preview.WriterPersistenceBoundaryReady ||
		!preview.KDESafeRedactedStatusOnly ||
		!preview.CompatibilityCenterAuthorizationModeled ||
		!preview.RuntimeDiagnosticsAuthorizationModeled ||
		preview.WriterCallable ||
		preview.WriterImplementationEnabled ||
		preview.WriterAuthorizationGranted ||
		preview.StatusWriterEnabled ||
		preview.StatusPersistenceAuthorized ||
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
		t.Fatalf("unsafe KDE-safe redacted status writer persistence authorization decision: %#v", preview)
	}
	if preview.AuthorizationItemCount != 10 ||
		preview.RequiredAuthorizationItemCount != 10 ||
		preview.ReadyAuthorizationItemCount != 10 ||
		preview.MissingAuthorizationItemCount != 0 ||
		preview.AuthorizedPersistenceItemCount != 0 ||
		preview.CallableWriterItemCount != 0 ||
		preview.PersistedWriterItemCount != 0 ||
		preview.RawExposedAuthorizationItemCount != 0 ||
		preview.SideEffectAuthorizationItemCount != 0 ||
		preview.CompatibilityCenterItemCount != 5 ||
		preview.RuntimeDiagnosticsItemCount != 5 {
		t.Fatalf("unexpected KDE-safe redacted status writer persistence authorization counts: %#v", preview)
	}
	if len(preview.AuthorizationItems) != 10 ||
		!sameStrings(preview.AuthorizationItemIDs, []string{
			"review-receipt-dry-run-result-lookup-consumer-enablement-kde-safe-compatibility-center-redacted-status-writer-persistence-authorization",
			"review-receipt-dry-run-result-lookup-consumer-enablement-kde-safe-runtime-diagnostics-redacted-status-writer-persistence-authorization",
			"renew-receipt-dry-run-result-lookup-consumer-enablement-kde-safe-compatibility-center-redacted-status-writer-persistence-authorization",
			"renew-receipt-dry-run-result-lookup-consumer-enablement-kde-safe-runtime-diagnostics-redacted-status-writer-persistence-authorization",
			"open-compatibility-center-dry-run-result-lookup-consumer-enablement-kde-safe-compatibility-center-redacted-status-writer-persistence-authorization",
			"open-compatibility-center-dry-run-result-lookup-consumer-enablement-kde-safe-runtime-diagnostics-redacted-status-writer-persistence-authorization",
			"dismiss-receipt-dry-run-result-lookup-consumer-enablement-kde-safe-compatibility-center-redacted-status-writer-persistence-authorization",
			"dismiss-receipt-dry-run-result-lookup-consumer-enablement-kde-safe-runtime-diagnostics-redacted-status-writer-persistence-authorization",
			"support-info-dry-run-result-lookup-consumer-enablement-kde-safe-compatibility-center-redacted-status-writer-persistence-authorization",
			"support-info-dry-run-result-lookup-consumer-enablement-kde-safe-runtime-diagnostics-redacted-status-writer-persistence-authorization",
		}) {
		t.Fatalf("unexpected KDE-safe redacted status writer persistence authorization items: %#v", preview.AuthorizationItemIDs)
	}
	for _, item := range preview.AuthorizationItems {
		if !item.EvidencePresent ||
			!item.CurrentMainlineConsumed ||
			!item.ClosedWriterImplementationConsumed ||
			!item.ClosedWriterImplementationReady ||
			!item.PersistenceAuthorizationModeled ||
			!item.WriterPersistenceBoundaryReady ||
			!item.KDESafeRedactedStatusOnly ||
			item.WriterCallable ||
			item.StatusPersistenceAuthorized ||
			item.StatusPersistenceWriteEnabled ||
			item.StatusWriterEnabled ||
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
			item.PersistenceAuthorizationStatus != "redacted-status-writer-persistence-authorization-modeled-writes-disabled" {
			t.Fatalf("unsafe KDE-safe redacted status writer persistence authorization item: %#v", item)
		}
	}
	if preview.Counts.Total != 8 ||
		preview.Counts.Passed != 8 ||
		preview.Counts.Pending != 0 ||
		preview.Counts.Blocked != 0 ||
		!sameStrings(preview.CheckIDs, []string{"current-mainline-consumed", "closed-writer-implementation-consumed", "writer-persistence-authorization-modeled", "compatibility-center-and-runtime-authorization-modeled", "ten-authorization-items-ready-persistence-disabled", "writer-persistence-writes-disabled", "consumer-lookup-request-and-support-disabled", "production-and-host-boundary-closed"}) {
		t.Fatalf("unexpected KDE-safe redacted status writer persistence authorization checks: counts=%#v ids=%#v", preview.Counts, preview.CheckIDs)
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
		t.Fatalf("unsafe KDE-safe redacted status writer persistence authorization gates: %#v", preview)
	}
	if err := validateNoBackendTerms(preview, "KDE-safe redacted status writer persistence authorization test"); err != nil {
		t.Fatalf("validateNoBackendTerms returned error: %v", err)
	}
}

func TestProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterPersistenceAuthorizationPreviewFailsClosedWithoutSources(t *testing.T) {
	root := t.TempDir()
	if err := os.WriteFile(filepath.Join(root, "VERSION"), []byte(currentProjectVersion(t)+"\n"), 0o600); err != nil {
		t.Fatalf("WriteFile VERSION returned error: %v", err)
	}
	preview, err := NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterPersistenceAuthorizationPreview(root)
	if err != nil {
		t.Fatalf("NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterPersistenceAuthorizationPreview returned error: %v", err)
	}
	if preview.CurrentMainlineConsumed ||
		preview.ClosedWriterImplementationConsumed ||
		preview.ClosedWriterImplementationReady ||
		!preview.PersistenceAuthorizationRequired ||
		!preview.PersistenceAuthorizationModeled ||
		preview.PersistenceAuthorizationReady ||
		preview.WriterPersistenceBoundaryReady ||
		preview.KDESafeRedactedStatusOnly ||
		preview.CompatibilityCenterAuthorizationModeled ||
		preview.RuntimeDiagnosticsAuthorizationModeled ||
		preview.WriterCallable ||
		preview.StatusPersistenceAuthorized ||
		preview.StatusPersistenceWriteEnabled ||
		preview.RawResultExposed ||
		preview.AuthorizationItemCount != 10 ||
		preview.RequiredAuthorizationItemCount != 10 ||
		preview.ReadyAuthorizationItemCount != 0 ||
		preview.MissingAuthorizationItemCount != 10 ||
		preview.Counts.Total != 8 ||
		preview.Counts.Passed != 3 ||
		preview.Counts.Blocked != 5 ||
		preview.PreviewDecision != "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-writer-persistence-authorization-blocked" {
		t.Fatalf("missing KDE-safe redacted status writer persistence authorization sources must fail closed: %#v", preview)
	}
	for _, item := range preview.AuthorizationItems {
		if item.EvidencePresent ||
			item.CurrentMainlineConsumed ||
			item.ClosedWriterImplementationConsumed ||
			item.ClosedWriterImplementationReady ||
			item.PersistenceAuthorizationModeled ||
			item.WriterPersistenceBoundaryReady ||
			item.WriterCallable ||
			item.StatusPersistenceAuthorized ||
			item.StatusPersistenceWriteEnabled ||
			item.RawResultExposed ||
			item.UserVisible ||
			item.PersistenceAuthorizationStatus != "missing-redacted-status-writer-persistence-authorization-evidence" {
			t.Fatalf("missing KDE-safe redacted status writer persistence authorization item must remain closed: %#v", item)
		}
	}
}
