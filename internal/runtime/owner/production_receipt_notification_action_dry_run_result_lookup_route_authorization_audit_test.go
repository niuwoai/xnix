package owner

import (
	"os"
	"path/filepath"
	"testing"
)

func TestProductionReceiptNotificationActionDryRunResultLookupRouteAuthorizationAuditPreviewModelsRoutes(t *testing.T) {
	preview, err := NewProductionReceiptNotificationActionDryRunResultLookupRouteAuthorizationAuditPreview(projectRootForOwnerTest(t))
	if err != nil {
		t.Fatalf("NewProductionReceiptNotificationActionDryRunResultLookupRouteAuthorizationAuditPreview returned error: %v", err)
	}

	if preview.Version != currentProjectVersion(t) ||
		preview.SchemaVersion != "xnix.runtime.production_receipt_notification_action_dry_run_result_lookup_route_authorization_audit.v1" ||
		preview.RequestType != "production-receipt-notification-action-dry-run-result-lookup-route-authorization-audit-preview" ||
		preview.AuditType != "receipt-notification-action-dry-run-result-lookup-route-authorization-audit" ||
		preview.AuditDecision != "production-receipt-notification-action-dry-run-result-lookup-route-authorization-audit-ready-route-disabled" ||
		preview.ReceiptSchema != "xnix.runtime.production_dbus_human_authorization_receipt.v1" ||
		preview.OpaqueReceiptID != ProductionDBusHumanAuthorizationReceiptID {
		t.Fatalf("unexpected lookup route authorization schema: %#v", preview)
	}
	if !preview.LookupRouteAuthorizationAuditRequired ||
		!preview.LookupRouteAuthorizationAuditModeled ||
		!preview.OpaqueLookupAuditConsumed ||
		!preview.LookupRouteGuidanceConsumed ||
		!preview.LookupRouteAuthorizationReady ||
		!preview.OwnerLocalLookupRouteModeled ||
		!preview.OpaqueResultIDSupported ||
		preview.KDEConsumptionAuthorized ||
		preview.RuntimeConsumptionAuthorized ||
		preview.LookupRouteAuthorized ||
		preview.LookupRouteEnabled ||
		preview.LookupRoutePersisted ||
		preview.OpaqueLookupEnabled ||
		preview.OpaqueLookupPersisted ||
		preview.CallerStateRootRequired ||
		preview.StateRootPathExposed ||
		preview.FilePathsExposed ||
		preview.FileContentRead ||
		preview.ResultPersistenceAuthorized ||
		preview.DispatchDryRunExecuted ||
		preview.DryRunResultPersisted ||
		preview.ResultVisibilityPersisted ||
		preview.RuntimeDiagnosticsPersisted ||
		preview.OperatorRouteApprovalPresent ||
		preview.ProductionReadiness ||
		preview.ProductionOwnershipReady {
		t.Fatalf("unexpected lookup route authorization decision: %#v", preview)
	}
	if preview.RouteItemCount != 5 ||
		preview.RequiredRouteItemCount != 5 ||
		preview.ReadyRouteItemCount != 5 ||
		preview.MissingRouteItemCount != 0 ||
		preview.AuthorizedRouteItemCount != 0 ||
		preview.EnabledRouteItemCount != 0 ||
		preview.PersistedRouteItemCount != 0 ||
		preview.LookupEnabledItemCount != 0 ||
		preview.PersistedLookupItemCount != 0 ||
		preview.PathExposedRouteItemCount != 0 ||
		preview.PersistedResultItemCount != 0 ||
		preview.ExecutedDryRunItemCount != 0 ||
		preview.SideEffectRouteItemCount != 0 ||
		!sameStrings(preview.RouteItemIDs, []string{"review-receipt-dry-run-result-lookup-route-authorization", "renew-receipt-dry-run-result-lookup-route-authorization", "open-compatibility-center-dry-run-result-lookup-route-authorization", "dismiss-receipt-dry-run-result-lookup-route-authorization", "support-info-dry-run-result-lookup-route-authorization"}) {
		t.Fatalf("unexpected lookup route authorization inventory: %#v", preview)
	}
	for _, item := range preview.RouteItems {
		if !item.EvidencePresent ||
			!item.RouteAuthorizationModeled ||
			!item.OwnerLocalRouteModeled ||
			!item.OpaqueLookupConsumed ||
			!item.OpaqueResultIDSupported ||
			item.KDEConsumptionAuthorized ||
			item.RuntimeConsumptionAuthorized ||
			item.LookupRouteAuthorized ||
			item.LookupRouteEnabled ||
			item.LookupRoutePersisted ||
			item.LookupEnabled ||
			item.LookupPersisted ||
			item.CallerStateRootRequired ||
			item.StateRootPathExposed ||
			item.FilePathsExposed ||
			item.FileContentRead ||
			!item.UserVisible ||
			!item.ReviewOnly ||
			!item.RuntimeOwned ||
			!item.GoRuntimeBacked ||
			item.KDEPolicyOwner ||
			!item.OperatorApprovalRequired ||
			item.OperatorApprovalPresent ||
			item.ResultPersistenceAuthorized ||
			item.DispatchDryRunExecuted ||
			item.DryRunResultPersisted ||
			item.ResultVisibilityPersisted ||
			item.RuntimeDiagnosticsPersisted ||
			item.RequestObjectCreated ||
			item.RequestObjectDispatched ||
			item.PortalRequestCreated ||
			item.ReceiptAccepted ||
			item.AuthorizationAccepted ||
			item.CompatibilityCenterOpened ||
			item.SupportBundleExported ||
			item.SupportCaseCreated ||
			item.ProductionReadiness ||
			item.ProductionOwnershipReady ||
			!item.SideEffectsDisabled ||
			item.HostRootModified ||
			item.InternalDetailsExposed ||
			item.RouteStatus != "lookup-route-authorization-modeled-route-disabled" {
			t.Fatalf("unsafe lookup route authorization item: %#v", item)
		}
	}
	if preview.Counts.Total != 9 ||
		preview.Counts.Passed != 9 ||
		preview.Counts.Pending != 0 ||
		preview.Counts.Blocked != 0 ||
		!sameStrings(preview.CheckIDs, []string{"opaque-lookup-audit-consumed", "lookup-route-guidance-consumed", "lookup-route-authorization-modeled-only", "five-lookup-route-authorization-items-present", "route-authorization-items-ready-route-disabled", "route-lookup-persistence-result-persistence-and-execution-disabled", "request-creation-dispatch-actions-and-portal-disabled", "receipt-notification-navigation-and-support-writes-disabled", "production-and-host-boundary-closed"}) {
		t.Fatalf("unexpected lookup route authorization checks: counts=%#v ids=%#v", preview.Counts, preview.CheckIDs)
	}
	if preview.LookupRouteAuthorized ||
		preview.LookupRouteEnabled ||
		preview.LookupRoutePersisted ||
		preview.OpaqueLookupEnabled ||
		preview.OpaqueLookupPersisted ||
		preview.RetentionPolicyPersistenceEnabled ||
		preview.DispatchDryRunExecutionEnabled ||
		preview.DryRunResultPersistenceEnabled ||
		preview.ResultVisibilityPersistenceEnabled ||
		preview.RuntimeDiagnosticsPersisted ||
		preview.RequestObjectCreationEnabled ||
		preview.RequestObjectDispatchEnabled ||
		preview.PortalRequestCreated ||
		preview.NotificationSent ||
		preview.CompatibilityCenterOpened ||
		preview.ReceiptWriterEnabled ||
		preview.BackendLaunchEnabled ||
		preview.HostRootModified ||
		preview.BackendDetailsExposed {
		t.Fatalf("unsafe lookup route authorization gates: %#v", preview)
	}
	if err := validateNoBackendTerms(preview, "lookup route authorization test"); err != nil {
		t.Fatalf("validateNoBackendTerms returned error: %v", err)
	}
}

func TestProductionReceiptNotificationActionDryRunResultLookupRouteAuthorizationAuditPreviewFailsClosedWithoutSources(t *testing.T) {
	root := t.TempDir()
	if err := os.WriteFile(filepath.Join(root, "VERSION"), []byte(currentProjectVersion(t)+"\n"), 0o600); err != nil {
		t.Fatalf("WriteFile VERSION returned error: %v", err)
	}
	preview, err := NewProductionReceiptNotificationActionDryRunResultLookupRouteAuthorizationAuditPreview(root)
	if err != nil {
		t.Fatalf("NewProductionReceiptNotificationActionDryRunResultLookupRouteAuthorizationAuditPreview returned error: %v", err)
	}
	if preview.OpaqueLookupAuditConsumed ||
		preview.LookupRouteGuidanceConsumed ||
		!preview.LookupRouteAuthorizationAuditRequired ||
		!preview.LookupRouteAuthorizationAuditModeled ||
		preview.LookupRouteAuthorizationReady ||
		preview.OwnerLocalLookupRouteModeled ||
		preview.OpaqueResultIDSupported ||
		preview.KDEConsumptionAuthorized ||
		preview.RuntimeConsumptionAuthorized ||
		preview.LookupRouteAuthorized ||
		preview.LookupRouteEnabled ||
		preview.LookupRoutePersisted ||
		preview.CallerStateRootRequired ||
		preview.StateRootPathExposed ||
		preview.FilePathsExposed ||
		preview.FileContentRead ||
		preview.RouteItemCount != 5 ||
		preview.RequiredRouteItemCount != 5 ||
		preview.ReadyRouteItemCount != 0 ||
		preview.MissingRouteItemCount != 5 ||
		preview.Counts.Total != 9 ||
		preview.Counts.Passed != 4 ||
		preview.Counts.Blocked != 5 ||
		preview.AuditDecision != "production-receipt-notification-action-dry-run-result-lookup-route-authorization-audit-blocked" {
		t.Fatalf("missing lookup route authorization sources must fail closed: %#v", preview)
	}
	for _, item := range preview.RouteItems {
		if item.EvidencePresent ||
			item.RouteAuthorizationModeled ||
			item.OwnerLocalRouteModeled ||
			item.OpaqueLookupConsumed ||
			item.OpaqueResultIDSupported ||
			item.KDEConsumptionAuthorized ||
			item.RuntimeConsumptionAuthorized ||
			item.LookupRouteAuthorized ||
			item.LookupRouteEnabled ||
			item.LookupRoutePersisted ||
			item.CallerStateRootRequired ||
			item.StateRootPathExposed ||
			item.FilePathsExposed ||
			item.FileContentRead ||
			item.UserVisible ||
			item.RouteStatus != "missing-lookup-route-authorization-evidence" {
			t.Fatalf("missing lookup route authorization item must remain closed: %#v", item)
		}
	}
}
