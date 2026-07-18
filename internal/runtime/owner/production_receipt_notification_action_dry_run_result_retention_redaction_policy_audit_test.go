package owner

import (
	"os"
	"path/filepath"
	"testing"
)

func TestProductionReceiptNotificationActionDryRunResultRetentionRedactionPolicyAuditPreviewModelsPolicy(t *testing.T) {
	preview, err := NewProductionReceiptNotificationActionDryRunResultRetentionRedactionPolicyAuditPreview(projectRootForOwnerTest(t))
	if err != nil {
		t.Fatalf("NewProductionReceiptNotificationActionDryRunResultRetentionRedactionPolicyAuditPreview returned error: %v", err)
	}

	if preview.Version != currentProjectVersion(t) ||
		preview.SchemaVersion != "xnix.runtime.production_receipt_notification_action_dry_run_result_retention_redaction_policy_audit.v1" ||
		preview.RequestType != "production-receipt-notification-action-dry-run-result-retention-redaction-policy-audit-preview" ||
		preview.AuditType != "receipt-notification-action-dry-run-result-retention-redaction-policy-audit" ||
		preview.AuditDecision != "production-receipt-notification-action-dry-run-result-retention-redaction-policy-audit-ready-policy-only" ||
		preview.ReceiptSchema != "xnix.runtime.production_dbus_human_authorization_receipt.v1" ||
		preview.OpaqueReceiptID != ProductionDBusHumanAuthorizationReceiptID {
		t.Fatalf("unexpected retention redaction policy schema: %#v", preview)
	}
	if !preview.RetentionRedactionPolicyAuditRequired ||
		!preview.RetentionRedactionPolicyAuditModeled ||
		!preview.PersistenceAuthorizationAuditConsumed ||
		!preview.RetentionRedactionGuidanceConsumed ||
		!preview.RetentionRedactionPolicyReady ||
		!preview.RetentionWindowModeled ||
		!preview.RedactionRulesModeled ||
		preview.ResultPersistenceAuthorized ||
		preview.RetentionEnforcementEnabled ||
		preview.RedactionEnforcementEnabled ||
		preview.DispatchDryRunExecuted ||
		preview.DryRunResultPersisted ||
		preview.ResultVisibilityPersisted ||
		preview.OperatorPersistenceApprovalPresent ||
		preview.ProductionReadiness ||
		preview.ProductionOwnershipReady {
		t.Fatalf("unexpected retention redaction policy decision: %#v", preview)
	}
	if preview.PolicyItemCount != 5 ||
		preview.RequiredPolicyItemCount != 5 ||
		preview.ReadyPolicyItemCount != 5 ||
		preview.MissingPolicyItemCount != 0 ||
		preview.RetentionEnabledItemCount != 0 ||
		preview.RedactionWriteItemCount != 0 ||
		preview.PersistedResultItemCount != 0 ||
		preview.ExecutedDryRunItemCount != 0 ||
		preview.CreatedRequestObjectCount != 0 ||
		preview.PortalRequestCreatedCount != 0 ||
		preview.SideEffectPolicyItemCount != 0 ||
		!sameStrings(preview.PolicyItemIDs, []string{"review-receipt-dry-run-result-retention-redaction-policy", "renew-receipt-dry-run-result-retention-redaction-policy", "open-compatibility-center-dry-run-result-retention-redaction-policy", "dismiss-receipt-dry-run-result-retention-redaction-policy", "support-info-dry-run-result-retention-redaction-policy"}) {
		t.Fatalf("unexpected retention redaction policy inventory: %#v", preview)
	}
	for _, item := range preview.PolicyItems {
		if !item.EvidencePresent ||
			!item.RetentionPolicyModeled ||
			!item.RedactionPolicyModeled ||
			!item.RedactedForKDE ||
			!item.RuntimeDiagnosticsModeled ||
			!item.UserVisible ||
			!item.ReviewOnly ||
			!item.RuntimeOwned ||
			!item.GoRuntimeBacked ||
			item.KDEPolicyOwner ||
			!item.OperatorApprovalRequired ||
			item.OperatorApprovalPresent ||
			item.ResultPersistenceAuthorized ||
			item.RetentionEnforcementEnabled ||
			item.RedactionEnforcementEnabled ||
			item.DispatchDryRunExecuted ||
			item.DryRunResultPersisted ||
			item.ResultVisibilityPersisted ||
			item.RequestObjectCreated ||
			item.RequestObjectDispatched ||
			item.PortalRequestCreated ||
			item.ReceiptAccepted ||
			item.AuthorizationAccepted ||
			item.CompatibilityCenterOpened ||
			item.CompatibilityCenterPersisted ||
			item.RuntimeDiagnosticsPersisted ||
			item.SupportBundleExported ||
			item.SupportCaseCreated ||
			item.ProductionReadiness ||
			item.ProductionOwnershipReady ||
			!item.SideEffectsDisabled ||
			item.HostRootModified ||
			item.InternalDetailsExposed ||
			item.PolicyStatus != "retention-redaction-policy-modeled-enforcement-disabled" {
			t.Fatalf("unsafe retention redaction policy item: %#v", item)
		}
	}
	if preview.Counts.Total != 9 ||
		preview.Counts.Passed != 9 ||
		preview.Counts.Pending != 0 ||
		preview.Counts.Blocked != 0 ||
		!sameStrings(preview.CheckIDs, []string{"persistence-authorization-audit-consumed", "retention-redaction-guidance-consumed", "retention-redaction-policy-modeled-only", "five-retention-redaction-policy-items-present", "policy-items-ready-enforcement-disabled", "enforcement-persistence-and-execution-disabled", "request-creation-dispatch-actions-and-portal-disabled", "receipt-notification-navigation-and-support-writes-disabled", "production-and-host-boundary-closed"}) {
		t.Fatalf("unexpected retention redaction policy checks: counts=%#v ids=%#v", preview.Counts, preview.CheckIDs)
	}
	if preview.RetentionPolicyPersistenceEnabled ||
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
		t.Fatalf("unsafe retention redaction policy gates: %#v", preview)
	}
	if err := validateNoBackendTerms(preview, "retention redaction policy test"); err != nil {
		t.Fatalf("validateNoBackendTerms returned error: %v", err)
	}
}

func TestProductionReceiptNotificationActionDryRunResultRetentionRedactionPolicyAuditPreviewFailsClosedWithoutSources(t *testing.T) {
	root := t.TempDir()
	if err := os.WriteFile(filepath.Join(root, "VERSION"), []byte(currentProjectVersion(t)+"\n"), 0o600); err != nil {
		t.Fatalf("WriteFile VERSION returned error: %v", err)
	}
	preview, err := NewProductionReceiptNotificationActionDryRunResultRetentionRedactionPolicyAuditPreview(root)
	if err != nil {
		t.Fatalf("NewProductionReceiptNotificationActionDryRunResultRetentionRedactionPolicyAuditPreview returned error: %v", err)
	}
	if preview.PersistenceAuthorizationAuditConsumed ||
		preview.RetentionRedactionGuidanceConsumed ||
		!preview.RetentionRedactionPolicyAuditRequired ||
		!preview.RetentionRedactionPolicyAuditModeled ||
		preview.RetentionRedactionPolicyReady ||
		preview.RetentionWindowModeled ||
		preview.RedactionRulesModeled ||
		preview.ResultPersistenceAuthorized ||
		preview.RetentionEnforcementEnabled ||
		preview.RedactionEnforcementEnabled ||
		preview.PolicyItemCount != 5 ||
		preview.RequiredPolicyItemCount != 5 ||
		preview.ReadyPolicyItemCount != 0 ||
		preview.MissingPolicyItemCount != 5 ||
		preview.Counts.Total != 9 ||
		preview.Counts.Passed != 4 ||
		preview.Counts.Blocked != 5 ||
		preview.AuditDecision != "production-receipt-notification-action-dry-run-result-retention-redaction-policy-audit-blocked" {
		t.Fatalf("missing retention redaction policy sources must fail closed: %#v", preview)
	}
	for _, item := range preview.PolicyItems {
		if item.EvidencePresent ||
			item.RetentionPolicyModeled ||
			item.RedactionPolicyModeled ||
			item.RedactedForKDE ||
			item.RuntimeDiagnosticsModeled ||
			item.UserVisible ||
			item.ResultPersistenceAuthorized ||
			item.RetentionEnforcementEnabled ||
			item.RedactionEnforcementEnabled ||
			item.DispatchDryRunExecuted ||
			item.DryRunResultPersisted ||
			item.ResultVisibilityPersisted ||
			item.PolicyStatus != "missing-retention-redaction-policy-evidence" {
			t.Fatalf("missing retention redaction policy item must remain closed: %#v", item)
		}
	}
}
