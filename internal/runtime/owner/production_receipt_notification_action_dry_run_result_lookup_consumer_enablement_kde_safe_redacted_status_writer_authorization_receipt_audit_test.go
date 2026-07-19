package owner

import (
	"os"
	"path/filepath"
	"testing"
)

func TestProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAuthorizationReceiptAuditPreviewModelsReceipt(t *testing.T) {
	preview, err := NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAuthorizationReceiptAuditPreview(projectRootForOwnerTest(t))
	if err != nil {
		t.Fatalf("NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAuthorizationReceiptAuditPreview returned error: %v", err)
	}

	if preview.Version != currentProjectVersion(t) ||
		preview.SchemaVersion != "xnix.runtime.production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_kde_safe_redacted_status_writer_authorization_receipt_audit.v1" ||
		preview.RequestType != "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-writer-authorization-receipt-audit-preview" ||
		preview.AuditType != "receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-writer-authorization-receipt-audit" ||
		preview.AuditDecision != "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-writer-authorization-receipt-audit-ready-receipt-disabled" {
		t.Fatalf("unexpected KDE-safe redacted status writer authorization receipt schema: %#v", preview)
	}
	if !preview.CurrentMainlineConsumed ||
		!preview.WriterAuthorizationGateAuditConsumed ||
		!preview.WriterAuthorizationGateReady ||
		!preview.WriterAuthorizationReceiptRequired ||
		!preview.WriterAuthorizationReceiptModeled ||
		!preview.WriterAuthorizationReceiptReady ||
		!preview.OpaqueWriterAuthorizationReceipt ||
		!preview.KDESafeRedactedStatusOnly ||
		!preview.CompatibilityCenterReceiptModeled ||
		!preview.RuntimeDiagnosticsReceiptModeled {
		t.Fatalf("writer authorization receipt readiness was not modeled: %#v", preview)
	}
	if preview.ReceiptPresent ||
		preview.ReceiptAccepted ||
		preview.AuthorizationAccepted ||
		preview.WriterAuthorizationGranted ||
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
		t.Fatalf("unsafe writer authorization receipt side effects: %#v", preview)
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
	if preview.ReceiptItemCount != 10 ||
		preview.RequiredReceiptItemCount != 10 ||
		preview.ReadyReceiptItemCount != 10 ||
		preview.MissingReceiptItemCount != 0 ||
		preview.AcceptedReceiptItemCount != 0 ||
		preview.GrantedWriterItemCount != 0 ||
		preview.EnabledWriterItemCount != 0 ||
		preview.PersistedWriterItemCount != 0 ||
		preview.RawExposedReceiptItemCount != 0 ||
		preview.SideEffectReceiptItemCount != 0 ||
		preview.CompatibilityCenterReceiptItemCount != 5 ||
		preview.RuntimeDiagnosticsReceiptItemCount != 5 {
		t.Fatalf("unexpected writer authorization receipt counts: %#v", preview)
	}
	if len(preview.ReceiptItems) != 10 || len(preview.ReceiptItemIDs) != 10 {
		t.Fatalf("writer authorization receipt must expose ten items: %#v", preview)
	}
	compatibilityCenterCount := 0
	runtimeDiagnosticsCount := 0
	for _, item := range preview.ReceiptItems {
		if !item.EvidencePresent ||
			!item.CurrentMainlineConsumed ||
			!item.WriterAuthorizationGateAuditConsumed ||
			!item.WriterAuthorizationReceiptModeled ||
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
			item.WriterAuthorizationReceiptStatus != "redacted-status-writer-authorization-receipt-modeled-acceptance-disabled" {
			t.Fatalf("unsafe writer authorization receipt item: %#v", item)
		}
		if item.SurfaceKind == "compatibility-center" {
			compatibilityCenterCount++
		}
		if item.SurfaceKind == "runtime-diagnostics" {
			runtimeDiagnosticsCount++
		}
	}
	if compatibilityCenterCount != 5 || runtimeDiagnosticsCount != 5 {
		t.Fatalf("unexpected writer authorization receipt surface split: %#v", preview.ReceiptItems)
	}
	if preview.Counts.Total != 8 ||
		preview.Counts.Passed != 8 ||
		preview.Counts.Pending != 0 ||
		preview.Counts.Blocked != 0 ||
		!sameStrings(preview.CheckIDs, []string{"current-mainline-consumed", "writer-authorization-gate-consumed", "writer-authorization-receipt-modeled", "compatibility-center-and-runtime-receipts-modeled", "ten-receipt-items-ready-acceptance-disabled", "receipt-acceptance-and-writers-disabled", "consumer-lookup-request-and-support-disabled", "production-and-host-boundary-closed"}) {
		t.Fatalf("unexpected writer authorization receipt checks: counts=%#v ids=%#v", preview.Counts, preview.CheckIDs)
	}
	if err := validateNoBackendTerms(preview, "KDE-safe redacted status writer authorization receipt test"); err != nil {
		t.Fatalf("validateNoBackendTerms returned error: %v", err)
	}
}

func TestProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAuthorizationReceiptAuditPreviewFailsClosedWithoutSources(t *testing.T) {
	root := t.TempDir()
	if err := os.WriteFile(filepath.Join(root, "VERSION"), []byte(currentProjectVersion(t)+"\n"), 0o600); err != nil {
		t.Fatalf("WriteFile VERSION returned error: %v", err)
	}
	preview, err := NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAuthorizationReceiptAuditPreview(root)
	if err != nil {
		t.Fatalf("NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAuthorizationReceiptAuditPreview returned error: %v", err)
	}

	if preview.CurrentMainlineConsumed ||
		preview.WriterAuthorizationGateAuditConsumed ||
		preview.WriterAuthorizationGateReady ||
		preview.WriterAuthorizationReceiptModeled ||
		preview.WriterAuthorizationReceiptReady ||
		preview.OpaqueWriterAuthorizationReceipt ||
		preview.KDESafeRedactedStatusOnly ||
		preview.CompatibilityCenterReceiptModeled ||
		preview.RuntimeDiagnosticsReceiptModeled ||
		preview.ReadyReceiptItemCount != 0 ||
		preview.MissingReceiptItemCount != 10 ||
		preview.AuditDecision != "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-writer-authorization-receipt-audit-blocked" ||
		preview.Counts.Total != 8 ||
		preview.Counts.Passed != 3 ||
		preview.Counts.Blocked != 5 {
		t.Fatalf("missing writer authorization receipt sources must fail closed: %#v", preview)
	}
}
