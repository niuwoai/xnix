package owner

import (
	"os"
	"path/filepath"
	"testing"
)

func TestProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusPersistenceAuthorizationAuditPreviewModelsAuthorization(t *testing.T) {
	preview, err := NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusPersistenceAuthorizationAuditPreview(projectRootForOwnerTest(t))
	if err != nil {
		t.Fatalf("NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusPersistenceAuthorizationAuditPreview returned error: %v", err)
	}

	if preview.Version != currentProjectVersion(t) ||
		preview.SchemaVersion != "xnix.runtime.production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_kde_safe_status_persistence_authorization_audit.v1" ||
		preview.RequestType != "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-status-persistence-authorization-audit-preview" ||
		preview.AuditType != "receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-status-persistence-authorization-audit" ||
		preview.AuditDecision != "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-status-persistence-authorization-audit-ready-persistence-disabled" {
		t.Fatalf("unexpected KDE-safe status persistence authorization schema: %#v", preview)
	}
	if !preview.PersistenceAuthorizationAuditRequired ||
		!preview.PersistenceAuthorizationModeled ||
		!preview.StatusFanOutAuditConsumed ||
		!preview.KDESafeStatusPersistenceGuidanceConsumed ||
		!preview.CompatibilityCenterPersistenceAuthorizationModeled ||
		!preview.RuntimeDiagnosticsPersistenceAuthorizationModeled ||
		!preview.PersistenceAuthorizationBoundaryReady ||
		!preview.StatusFanOutReady ||
		!preview.ConsumerEnablementGateClosed ||
		!preview.KDESafeStatusOnly ||
		!preview.OpaqueResultIDSupported ||
		preview.StatusPersistenceAuthorized ||
		preview.KDEStatusPersistenceAuthorized ||
		preview.RuntimeDiagnosticsPersistenceAuthorized ||
		preview.KDEStatusPersisted ||
		preview.RuntimeDiagnosticsPersisted ||
		preview.ConsumerConsumptionAuthorized ||
		preview.ConsumerEnablementAuthorized ||
		preview.KDEConsumerEnabled ||
		preview.RuntimeConsumerEnabled ||
		preview.LookupRouteEnabled ||
		preview.OpaqueLookupEnabled ||
		preview.RawResultExposed ||
		preview.DispatchDryRunExecuted ||
		preview.ProductionReadiness ||
		preview.ProductionOwnershipReady {
		t.Fatalf("unexpected KDE-safe status persistence authorization decision: %#v", preview)
	}
	if preview.AuthorizationItemCount != 10 ||
		preview.RequiredAuthorizationItemCount != 10 ||
		preview.ReadyAuthorizationItemCount != 10 ||
		preview.MissingAuthorizationItemCount != 0 ||
		preview.CompatibilityCenterAuthorizationItemCount != 5 ||
		preview.RuntimeDiagnosticsAuthorizationItemCount != 5 ||
		preview.AuthorizedPersistenceItemCount != 0 ||
		preview.PersistedStatusItemCount != 0 ||
		preview.ConsumerEnabledAuthorizationItemCount != 0 ||
		preview.RawExposedAuthorizationItemCount != 0 ||
		preview.SideEffectAuthorizationItemCount != 0 {
		t.Fatalf("unexpected KDE-safe status persistence authorization inventory: %#v", preview)
	}
	for _, item := range preview.AuthorizationItems {
		if !item.EvidencePresent ||
			!item.StatusFanOutConsumed ||
			!item.PersistenceAuthorizationModeled ||
			!item.KDESafeStatusOnly ||
			item.StatusPersistenceAuthorized ||
			item.KDEStatusPersistenceAuthorized ||
			item.RuntimeDiagnosticsPersistenceAuthorized ||
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
			!item.UserVisible ||
			!item.ReviewOnly ||
			!item.RuntimeOwned ||
			!item.GoRuntimeBacked ||
			item.KDEPolicyOwner ||
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
			item.AuthorizationStatus != "kde-safe-status-persistence-authorization-modeled-persistence-disabled" {
			t.Fatalf("unsafe KDE-safe status persistence authorization item: %#v", item)
		}
		if item.SurfaceKind == "compatibility-center" && (!item.CompatibilityCenterPersistenceCandidate || item.RuntimeDiagnosticsPersistenceCandidate) {
			t.Fatalf("Compatibility Center persistence authorization item must stay surface-specific: %#v", item)
		}
		if item.SurfaceKind == "runtime-diagnostics" && (!item.RuntimeDiagnosticsPersistenceCandidate || item.CompatibilityCenterPersistenceCandidate) {
			t.Fatalf("Runtime diagnostics persistence authorization item must stay surface-specific: %#v", item)
		}
	}
	if preview.Counts.Total != 10 ||
		preview.Counts.Passed != 10 ||
		preview.Counts.Pending != 0 ||
		preview.Counts.Blocked != 0 ||
		!sameStrings(preview.CheckIDs, []string{"kde-safe-status-persistence-authorization-audit-required", "status-fanout-audit-consumed", "kde-safe-status-persistence-guidance-consumed", "compatibility-center-persistence-authorization-modeled", "runtime-diagnostics-persistence-authorization-modeled", "ten-kde-safe-status-persistence-authorization-items-present", "persistence-authorization-boundary-modeled-only", "persistence-writes-consumers-and-lookup-disabled", "request-notification-navigation-and-support-disabled", "production-and-host-boundary-closed"}) {
		t.Fatalf("unexpected KDE-safe status persistence authorization checks: counts=%#v ids=%#v", preview.Counts, preview.CheckIDs)
	}
	if preview.StatusPersistenceAuthorized ||
		preview.KDEStatusPersistenceAuthorized ||
		preview.RuntimeDiagnosticsPersistenceAuthorized ||
		preview.KDEStatusPersisted ||
		preview.RuntimeDiagnosticsPersisted ||
		preview.ConsumerEnablementAuthorized ||
		preview.KDEConsumerEnabled ||
		preview.RuntimeConsumerEnabled ||
		preview.LookupRouteEnabled ||
		preview.OpaqueLookupEnabled ||
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
		t.Fatalf("unsafe KDE-safe status persistence authorization gates: %#v", preview)
	}
	if err := validateNoBackendTerms(preview, "KDE-safe status persistence authorization test"); err != nil {
		t.Fatalf("validateNoBackendTerms returned error: %v", err)
	}
}

func TestProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusPersistenceAuthorizationAuditPreviewFailsClosedWithoutSources(t *testing.T) {
	root := t.TempDir()
	if err := os.WriteFile(filepath.Join(root, "VERSION"), []byte(currentProjectVersion(t)+"\n"), 0o600); err != nil {
		t.Fatalf("WriteFile VERSION returned error: %v", err)
	}
	preview, err := NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusPersistenceAuthorizationAuditPreview(root)
	if err != nil {
		t.Fatalf("NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusPersistenceAuthorizationAuditPreview returned error: %v", err)
	}
	if !preview.PersistenceAuthorizationAuditRequired ||
		!preview.PersistenceAuthorizationModeled ||
		preview.StatusFanOutAuditConsumed ||
		preview.KDESafeStatusPersistenceGuidanceConsumed ||
		preview.CompatibilityCenterPersistenceAuthorizationModeled ||
		preview.RuntimeDiagnosticsPersistenceAuthorizationModeled ||
		preview.PersistenceAuthorizationBoundaryReady ||
		preview.StatusFanOutReady ||
		preview.KDESafeStatusOnly ||
		preview.OpaqueResultIDSupported ||
		preview.StatusPersistenceAuthorized ||
		preview.KDEStatusPersisted ||
		preview.RuntimeDiagnosticsPersisted ||
		preview.RawResultExposed ||
		preview.AuthorizationItemCount != 10 ||
		preview.RequiredAuthorizationItemCount != 10 ||
		preview.ReadyAuthorizationItemCount != 0 ||
		preview.MissingAuthorizationItemCount != 10 ||
		preview.CompatibilityCenterAuthorizationItemCount != 5 ||
		preview.RuntimeDiagnosticsAuthorizationItemCount != 5 ||
		preview.Counts.Total != 10 ||
		preview.Counts.Passed != 4 ||
		preview.Counts.Blocked != 6 ||
		preview.AuditDecision != "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-status-persistence-authorization-audit-blocked" {
		t.Fatalf("missing KDE-safe status persistence authorization sources must fail closed: %#v", preview)
	}
	for _, item := range preview.AuthorizationItems {
		if item.EvidencePresent ||
			item.StatusFanOutConsumed ||
			item.PersistenceAuthorizationModeled ||
			item.KDESafeStatusOnly ||
			item.StatusPersistenceAuthorized ||
			item.KDEStatusPersisted ||
			item.RuntimeDiagnosticsPersisted ||
			item.UserVisible ||
			item.RawResultExposed ||
			item.AuthorizationStatus != "missing-kde-safe-status-persistence-authorization-evidence" {
			t.Fatalf("missing KDE-safe status persistence authorization item must remain closed: %#v", item)
		}
	}
}
