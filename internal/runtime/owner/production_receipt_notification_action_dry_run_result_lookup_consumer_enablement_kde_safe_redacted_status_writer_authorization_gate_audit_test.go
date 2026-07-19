package owner

import (
	"os"
	"path/filepath"
	"testing"
)

func TestProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAuthorizationGateAuditPreviewModelsGate(t *testing.T) {
	preview, err := NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAuthorizationGateAuditPreview(projectRootForOwnerTest(t))
	if err != nil {
		t.Fatalf("NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAuthorizationGateAuditPreview returned error: %v", err)
	}

	if preview.Version != currentProjectVersion(t) ||
		preview.SchemaVersion != "xnix.runtime.production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_kde_safe_redacted_status_writer_authorization_gate_audit.v1" ||
		preview.RequestType != "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-writer-authorization-gate-audit-preview" ||
		preview.AuditType != "receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-writer-authorization-gate-audit" ||
		preview.AuditDecision != "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-writer-authorization-gate-audit-ready-writer-disabled" {
		t.Fatalf("unexpected KDE-safe redacted status writer authorization gate schema: %#v", preview)
	}
	if !preview.CurrentMainlineConsumed ||
		!preview.RedactedWriteModelAuditConsumed ||
		!preview.RedactedWriteModelBoundaryReady ||
		!preview.WriterAuthorizationGateRequired ||
		!preview.WriterAuthorizationGateModeled ||
		!preview.WriterAuthorizationBoundaryReady ||
		!preview.CompatibilityCenterWriterModeled ||
		!preview.RuntimeDiagnosticsWriterModeled ||
		!preview.KDESafeRedactedStatusOnly ||
		!preview.OpaqueResultIDSupported {
		t.Fatalf("writer authorization gate readiness was not modeled: %#v", preview)
	}
	if preview.WriterAuthorizationGranted ||
		preview.StatusWriterEnabled ||
		preview.StatusPersistenceWriteEnabled ||
		preview.KDEStatusWriteEnabled ||
		preview.RuntimeDiagnosticsWriteEnabled ||
		preview.RedactedSummaryPersisted ||
		preview.KDEStatusPersisted ||
		preview.RuntimeDiagnosticsPersisted ||
		preview.DryRunResultPersisted ||
		preview.RawResultExposed ||
		preview.ConsumerConsumptionAuthorized ||
		preview.ConsumerEnablementAuthorized ||
		preview.KDEConsumerEnabled ||
		preview.RuntimeConsumerEnabled ||
		preview.LookupRouteAuthorized ||
		preview.LookupRouteEnabled ||
		preview.OpaqueLookupEnabled ||
		preview.RequestObjectCreationEnabled ||
		preview.RequestObjectDispatchEnabled ||
		preview.PortalRequestCreated ||
		preview.NotificationActionEnabled ||
		preview.CompatibilityCenterOpened ||
		preview.SupportBundleExported ||
		preview.SupportCaseCreated {
		t.Fatalf("unsafe writer authorization gate side effects: %#v", preview)
	}
	if preview.SystemServiceStarted ||
		preview.SessionBusClaimed ||
		preview.ProductionBusClaimed ||
		preview.ProductionOwnerEnabled ||
		preview.ProductionReadiness ||
		preview.ProductionOwnershipReady ||
		preview.WriteMethodsEnabled ||
		preview.RuntimeWritesEnabled ||
		preview.DesktopFilesWritten ||
		preview.SettingsPersisted ||
		preview.AdapterInvocationEnabled ||
		preview.BackendLaunchEnabled ||
		preview.BackendProcessStarted ||
		preview.SnapshotRestoreExecuted ||
		preview.StateCleanupExecuted ||
		preview.NetworkRequired ||
		preview.HostRootModified ||
		preview.PrivilegedContainerRequired ||
		preview.CallerStateRootRequired ||
		preview.StateRootPathExposed ||
		preview.FilePathsExposed ||
		preview.FileContentRead ||
		preview.RawCommandExposed ||
		preview.RawExecutableExposed ||
		preview.BackendDetailsExposed {
		t.Fatalf("unsafe production or host gates: %#v", preview)
	}
	if preview.WriterItemCount != 10 ||
		preview.RequiredWriterItemCount != 10 ||
		preview.ReadyWriterItemCount != 10 ||
		preview.MissingWriterItemCount != 0 ||
		preview.CompatibilityCenterWriterItemCount != 5 ||
		preview.RuntimeDiagnosticsWriterItemCount != 5 ||
		preview.GrantedWriterItemCount != 0 ||
		preview.EnabledWriterItemCount != 0 ||
		preview.PersistedWriterItemCount != 0 ||
		preview.RawExposedWriterItemCount != 0 ||
		preview.SideEffectWriterItemCount != 0 {
		t.Fatalf("unexpected writer authorization gate counts: %#v", preview)
	}
	if len(preview.WriterItems) != 10 || len(preview.WriterItemIDs) != 10 {
		t.Fatalf("writer authorization gate must expose ten items: %#v", preview)
	}
	compatibilityCenterCount := 0
	runtimeDiagnosticsCount := 0
	for _, item := range preview.WriterItems {
		if !item.EvidencePresent ||
			!item.CurrentMainlineConsumed ||
			!item.RedactedWriteModelAuditConsumed ||
			!item.WriterAuthorizationModeled ||
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
			item.WriterAuthorizationStatus != "redacted-status-writer-authorization-modeled-writer-disabled" {
			t.Fatalf("unsafe writer authorization gate item: %#v", item)
		}
		if item.SurfaceKind == "compatibility-center" {
			compatibilityCenterCount++
		}
		if item.SurfaceKind == "runtime-diagnostics" {
			runtimeDiagnosticsCount++
		}
	}
	if compatibilityCenterCount != 5 || runtimeDiagnosticsCount != 5 {
		t.Fatalf("unexpected writer authorization gate surface split: %#v", preview.WriterItems)
	}
	if preview.Counts.Total != 8 ||
		preview.Counts.Passed != 8 ||
		preview.Counts.Pending != 0 ||
		preview.Counts.Blocked != 0 ||
		!sameStrings(preview.CheckIDs, []string{"current-mainline-consumed", "redacted-write-model-audit-consumed", "writer-authorization-gate-modeled", "compatibility-center-and-runtime-writers-modeled", "ten-writer-items-ready-writer-disabled", "writer-grants-and-persistence-disabled", "consumer-lookup-request-and-support-disabled", "production-and-host-boundary-closed"}) {
		t.Fatalf("unexpected writer authorization gate checks: counts=%#v ids=%#v", preview.Counts, preview.CheckIDs)
	}
	if err := validateNoBackendTerms(preview, "KDE-safe redacted status writer authorization gate test"); err != nil {
		t.Fatalf("validateNoBackendTerms returned error: %v", err)
	}
}

func TestProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAuthorizationGateAuditPreviewFailsClosedWithoutSources(t *testing.T) {
	root := t.TempDir()
	if err := os.WriteFile(filepath.Join(root, "VERSION"), []byte(currentProjectVersion(t)+"\n"), 0o600); err != nil {
		t.Fatalf("WriteFile VERSION returned error: %v", err)
	}
	preview, err := NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAuthorizationGateAuditPreview(root)
	if err != nil {
		t.Fatalf("NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAuthorizationGateAuditPreview returned error: %v", err)
	}

	if preview.CurrentMainlineConsumed ||
		preview.RedactedWriteModelAuditConsumed ||
		preview.RedactedWriteModelBoundaryReady ||
		preview.WriterAuthorizationGateModeled ||
		preview.WriterAuthorizationBoundaryReady ||
		preview.CompatibilityCenterWriterModeled ||
		preview.RuntimeDiagnosticsWriterModeled ||
		preview.KDESafeRedactedStatusOnly ||
		preview.OpaqueResultIDSupported ||
		preview.ReadyWriterItemCount != 0 ||
		preview.MissingWriterItemCount != 10 ||
		preview.AuditDecision != "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-writer-authorization-gate-audit-blocked" ||
		preview.Counts.Total != 8 ||
		preview.Counts.Passed != 3 ||
		preview.Counts.Blocked != 5 {
		t.Fatalf("missing writer authorization gate sources must fail closed: %#v", preview)
	}
}
