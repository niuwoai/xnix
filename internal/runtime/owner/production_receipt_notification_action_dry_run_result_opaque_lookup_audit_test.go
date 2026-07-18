package owner

import (
	"os"
	"path/filepath"
	"testing"
)

func TestProductionReceiptNotificationActionDryRunResultOpaqueLookupAuditPreviewModelsLookup(t *testing.T) {
	preview, err := NewProductionReceiptNotificationActionDryRunResultOpaqueLookupAuditPreview(projectRootForOwnerTest(t))
	if err != nil {
		t.Fatalf("NewProductionReceiptNotificationActionDryRunResultOpaqueLookupAuditPreview returned error: %v", err)
	}

	if preview.Version != currentProjectVersion(t) ||
		preview.SchemaVersion != "xnix.runtime.production_receipt_notification_action_dry_run_result_opaque_lookup_audit.v1" ||
		preview.RequestType != "production-receipt-notification-action-dry-run-result-opaque-lookup-audit-preview" ||
		preview.AuditType != "receipt-notification-action-dry-run-result-opaque-lookup-audit" ||
		preview.AuditDecision != "production-receipt-notification-action-dry-run-result-opaque-lookup-audit-ready-lookup-disabled" ||
		preview.ReceiptSchema != "xnix.runtime.production_dbus_human_authorization_receipt.v1" ||
		preview.OpaqueReceiptID != ProductionDBusHumanAuthorizationReceiptID {
		t.Fatalf("unexpected opaque lookup schema: %#v", preview)
	}
	if !preview.OpaqueLookupAuditRequired ||
		!preview.OpaqueLookupAuditModeled ||
		!preview.RetentionRedactionPolicyAuditConsumed ||
		!preview.OpaqueLookupGuidanceConsumed ||
		!preview.OpaqueLookupBoundaryReady ||
		!preview.OwnerManagedOpaqueLookupModeled ||
		!preview.OpaqueResultIDSupported ||
		preview.OpaqueLookupEnabled ||
		preview.OpaqueLookupPersisted ||
		preview.CallerStateRootRequired ||
		preview.StateRootPathExposed ||
		preview.FilePathsExposed ||
		preview.FileContentRead ||
		preview.ResultPersistenceAuthorized ||
		preview.RetentionEnforcementEnabled ||
		preview.RedactionEnforcementEnabled ||
		preview.DispatchDryRunExecuted ||
		preview.DryRunResultPersisted ||
		preview.ResultVisibilityPersisted ||
		preview.RuntimeDiagnosticsPersisted ||
		preview.OperatorLookupApprovalPresent ||
		preview.ProductionReadiness ||
		preview.ProductionOwnershipReady {
		t.Fatalf("unexpected opaque lookup decision: %#v", preview)
	}
	if preview.LookupItemCount != 5 ||
		preview.RequiredLookupItemCount != 5 ||
		preview.ReadyLookupItemCount != 5 ||
		preview.MissingLookupItemCount != 0 ||
		preview.EnabledLookupItemCount != 0 ||
		preview.PersistedLookupItemCount != 0 ||
		preview.PathExposedLookupItemCount != 0 ||
		preview.PersistedResultItemCount != 0 ||
		preview.ExecutedDryRunItemCount != 0 ||
		preview.CreatedRequestObjectCount != 0 ||
		preview.PortalRequestCreatedCount != 0 ||
		preview.SideEffectLookupItemCount != 0 ||
		!sameStrings(preview.LookupItemIDs, []string{"review-receipt-dry-run-result-opaque-lookup", "renew-receipt-dry-run-result-opaque-lookup", "open-compatibility-center-dry-run-result-opaque-lookup", "dismiss-receipt-dry-run-result-opaque-lookup", "support-info-dry-run-result-opaque-lookup"}) {
		t.Fatalf("unexpected opaque lookup inventory: %#v", preview)
	}
	for _, item := range preview.LookupItems {
		if !item.EvidencePresent ||
			!item.OpaqueLookupModeled ||
			!item.OwnerManagedLookup ||
			!item.OpaqueResultIDSupported ||
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
			item.LookupEnabled ||
			item.LookupPersisted ||
			item.ResultPersistenceAuthorized ||
			item.RetentionEnforcementEnabled ||
			item.RedactionEnforcementEnabled ||
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
			item.CompatibilityCenterPersisted ||
			item.SupportBundleExported ||
			item.SupportCaseCreated ||
			item.ProductionReadiness ||
			item.ProductionOwnershipReady ||
			!item.SideEffectsDisabled ||
			item.HostRootModified ||
			item.InternalDetailsExposed ||
			item.LookupStatus != "opaque-lookup-modeled-lookup-disabled" {
			t.Fatalf("unsafe opaque lookup item: %#v", item)
		}
	}
	if preview.Counts.Total != 9 ||
		preview.Counts.Passed != 9 ||
		preview.Counts.Pending != 0 ||
		preview.Counts.Blocked != 0 ||
		!sameStrings(preview.CheckIDs, []string{"retention-redaction-policy-audit-consumed", "opaque-lookup-guidance-consumed", "opaque-lookup-modeled-only", "five-opaque-lookup-items-present", "opaque-lookup-items-ready-lookup-disabled", "lookup-persistence-result-persistence-and-execution-disabled", "request-creation-dispatch-actions-and-portal-disabled", "receipt-notification-navigation-and-support-writes-disabled", "production-and-host-boundary-closed"}) {
		t.Fatalf("unexpected opaque lookup checks: counts=%#v ids=%#v", preview.Counts, preview.CheckIDs)
	}
	if preview.OpaqueLookupEnabled ||
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
		t.Fatalf("unsafe opaque lookup gates: %#v", preview)
	}
	if err := validateNoBackendTerms(preview, "opaque lookup test"); err != nil {
		t.Fatalf("validateNoBackendTerms returned error: %v", err)
	}
}

func TestProductionReceiptNotificationActionDryRunResultOpaqueLookupAuditPreviewFailsClosedWithoutSources(t *testing.T) {
	root := t.TempDir()
	if err := os.WriteFile(filepath.Join(root, "VERSION"), []byte(currentProjectVersion(t)+"\n"), 0o600); err != nil {
		t.Fatalf("WriteFile VERSION returned error: %v", err)
	}
	preview, err := NewProductionReceiptNotificationActionDryRunResultOpaqueLookupAuditPreview(root)
	if err != nil {
		t.Fatalf("NewProductionReceiptNotificationActionDryRunResultOpaqueLookupAuditPreview returned error: %v", err)
	}
	if preview.RetentionRedactionPolicyAuditConsumed ||
		preview.OpaqueLookupGuidanceConsumed ||
		!preview.OpaqueLookupAuditRequired ||
		!preview.OpaqueLookupAuditModeled ||
		preview.OpaqueLookupBoundaryReady ||
		preview.OwnerManagedOpaqueLookupModeled ||
		preview.OpaqueResultIDSupported ||
		preview.OpaqueLookupEnabled ||
		preview.OpaqueLookupPersisted ||
		preview.CallerStateRootRequired ||
		preview.StateRootPathExposed ||
		preview.FilePathsExposed ||
		preview.FileContentRead ||
		preview.LookupItemCount != 5 ||
		preview.RequiredLookupItemCount != 5 ||
		preview.ReadyLookupItemCount != 0 ||
		preview.MissingLookupItemCount != 5 ||
		preview.Counts.Total != 9 ||
		preview.Counts.Passed != 4 ||
		preview.Counts.Blocked != 5 ||
		preview.AuditDecision != "production-receipt-notification-action-dry-run-result-opaque-lookup-audit-blocked" {
		t.Fatalf("missing opaque lookup sources must fail closed: %#v", preview)
	}
	for _, item := range preview.LookupItems {
		if item.EvidencePresent ||
			item.OpaqueLookupModeled ||
			item.OwnerManagedLookup ||
			item.OpaqueResultIDSupported ||
			item.CallerStateRootRequired ||
			item.StateRootPathExposed ||
			item.FilePathsExposed ||
			item.FileContentRead ||
			item.UserVisible ||
			item.LookupEnabled ||
			item.LookupPersisted ||
			item.ResultPersistenceAuthorized ||
			item.RetentionEnforcementEnabled ||
			item.RedactionEnforcementEnabled ||
			item.DispatchDryRunExecuted ||
			item.DryRunResultPersisted ||
			item.ResultVisibilityPersisted ||
			item.RuntimeDiagnosticsPersisted ||
			item.LookupStatus != "missing-opaque-lookup-evidence" {
			t.Fatalf("missing opaque lookup item must remain closed: %#v", item)
		}
	}
}
