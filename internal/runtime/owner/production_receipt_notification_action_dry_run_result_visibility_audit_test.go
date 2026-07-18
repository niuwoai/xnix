package owner

import (
	"os"
	"path/filepath"
	"testing"
)

func TestProductionReceiptNotificationActionDryRunResultVisibilityAuditPreviewModelsVisibility(t *testing.T) {
	preview, err := NewProductionReceiptNotificationActionDryRunResultVisibilityAuditPreview(projectRootForOwnerTest(t))
	if err != nil {
		t.Fatalf("NewProductionReceiptNotificationActionDryRunResultVisibilityAuditPreview returned error: %v", err)
	}

	if preview.Version != currentProjectVersion(t) ||
		preview.SchemaVersion != "xnix.runtime.production_receipt_notification_action_dry_run_result_visibility_audit.v1" ||
		preview.RequestType != "production-receipt-notification-action-dry-run-result-visibility-audit-preview" ||
		preview.AuditType != "receipt-notification-action-dry-run-result-visibility-audit" ||
		preview.AuditDecision != "production-receipt-notification-action-dry-run-result-visibility-audit-ready-visibility-only" ||
		preview.ReceiptSchema != "xnix.runtime.production_dbus_human_authorization_receipt.v1" ||
		preview.OpaqueReceiptID != ProductionDBusHumanAuthorizationReceiptID {
		t.Fatalf("unexpected production receipt notification action dry-run result visibility schema: %#v", preview)
	}
	if !preview.ResultVisibilityAuditRequired ||
		!preview.ResultVisibilityAuditModeled ||
		!preview.DispatchDryRunAuditConsumed ||
		!preview.ResultVisibilityGuidanceConsumed ||
		!preview.ResultVisibilityPlanReady ||
		!preview.RedactedKDEVisibilityModeled ||
		!preview.RuntimeDiagnosticsVisibilityModeled ||
		preview.DispatchDryRunExecuted ||
		preview.DryRunResultPersisted ||
		preview.DispatchAuthorizationGranted ||
		!preview.OperatorDispatchApprovalRequired ||
		preview.OperatorDispatchApprovalPresent ||
		preview.CallerStateRootRequired ||
		preview.ReceiptPresent ||
		preview.ReceiptAccepted ||
		preview.AuthorizationAccepted ||
		preview.ProductionReadiness ||
		preview.ProductionOwnershipReady {
		t.Fatalf("unexpected production receipt notification action dry-run result visibility decision: %#v", preview)
	}
	if preview.VisibilityItemCount != 5 ||
		preview.RequiredVisibilityItemCount != 5 ||
		preview.ReadyVisibilityItemCount != 5 ||
		preview.MissingVisibilityItemCount != 0 ||
		preview.PersistedVisibilityItemCount != 0 ||
		preview.ExecutedDryRunItemCount != 0 ||
		preview.DispatchedDryRunItemCount != 0 ||
		preview.CreatedRequestObjectCount != 0 ||
		preview.PortalRequestCreatedCount != 0 ||
		preview.SideEffectVisibilityItemCount != 0 ||
		!sameStrings(preview.VisibilityItemIDs, []string{"review-receipt-dry-run-result-visibility", "renew-receipt-dry-run-result-visibility", "open-compatibility-center-dry-run-result-visibility", "dismiss-receipt-dry-run-result-visibility", "support-info-dry-run-result-visibility"}) {
		t.Fatalf("unexpected production receipt notification action dry-run result visibility inventory: %#v", preview)
	}
	for _, item := range preview.VisibilityItems {
		if !item.EvidencePresent ||
			!item.ResultVisibilityModeled ||
			!item.RedactedForKDE ||
			!item.RuntimeDiagnosticsModeled ||
			!item.UserVisible ||
			!item.ReviewOnly ||
			!item.RuntimeOwned ||
			!item.GoRuntimeBacked ||
			item.KDEPolicyOwner ||
			!item.OperatorApprovalRequired ||
			item.OperatorApprovalPresent ||
			item.DispatchAuthorizationGranted ||
			item.DispatchDryRunExecuted ||
			item.DryRunResultPersisted ||
			item.ResultVisibilityPersisted ||
			item.RequestObjectCreated ||
			item.RequestObjectDispatched ||
			item.RequestObjectPersisted ||
			item.PortalRequestCreated ||
			item.NavigationRequested ||
			item.NotificationActionEnabled ||
			item.ActionEnabled ||
			item.ReceiptAccepted ||
			item.AuthorizationAccepted ||
			item.ReceiptWriterEnabled ||
			item.ReceiptPersistenceEnabled ||
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
			item.VisibilityStatus != "dry-run-result-visibility-modeled-persistence-disabled" {
			t.Fatalf("unsafe or missing production receipt notification action dry-run result visibility item: %#v", item)
		}
	}
	if preview.Counts.Total != 9 ||
		preview.Counts.Passed != 9 ||
		preview.Counts.Pending != 0 ||
		preview.Counts.Blocked != 0 ||
		!sameStrings(preview.CheckIDs, []string{"dispatch-dry-run-audit-consumed", "dry-run-result-visibility-guidance-consumed", "dry-run-result-visibility-modeled-only", "five-dry-run-result-visibility-items-present", "visibility-items-ready-results-redacted", "dry-run-execution-and-result-persistence-disabled", "request-creation-dispatch-actions-and-portal-disabled", "receipt-notification-navigation-and-support-writes-disabled", "production-and-host-boundary-closed"}) {
		t.Fatalf("unexpected production receipt notification action dry-run result visibility checks: counts=%#v ids=%#v", preview.Counts, preview.CheckIDs)
	}
	if !preview.RuntimeOwned ||
		!preview.GoRuntimeBacked ||
		preview.KDEPolicyOwner ||
		!preview.OfficialDesktopOnly ||
		preview.PlasmaForkRequired ||
		preview.PlasmaSourceModified ||
		preview.SystemServiceStarted ||
		preview.SessionBusClaimed ||
		preview.ProductionBusClaimed ||
		preview.ProductionOwnerEnabled ||
		preview.ProductionActivationReady ||
		preview.WriteMethodsEnabled ||
		preview.RuntimeWritesEnabled ||
		preview.RequestObjectCreationEnabled ||
		preview.RequestObjectDispatchEnabled ||
		preview.RequestObjectPersistenceEnabled ||
		preview.DispatchAuthorizationPersisted ||
		preview.DispatchDryRunExecutionEnabled ||
		preview.DryRunResultPersistenceEnabled ||
		preview.ResultVisibilityPersistenceEnabled ||
		preview.PortalRequestCreated ||
		preview.RequestObjectsCreated ||
		preview.RequestObjectsDispatched ||
		preview.NotificationSent ||
		preview.NotificationDeliveryEnabled ||
		preview.NotificationActionEnabled ||
		preview.ReviewActionEnabled ||
		preview.RenewActionEnabled ||
		preview.OpenCompatibilityCenterEnabled ||
		preview.DismissActionEnabled ||
		preview.SupportInfoActionEnabled ||
		preview.CompatibilityCenterOpened ||
		preview.CompatibilityCenterPersisted ||
		preview.RuntimeDiagnosticsPersisted ||
		preview.ReceiptWriterEnabled ||
		preview.ReceiptPersistenceEnabled ||
		preview.ReceiptLookupWritesEnabled ||
		preview.ReceiptReplayEnabled ||
		preview.ReceiptExpiryWriteEnabled ||
		preview.ReceiptRevocationWriteEnabled ||
		preview.DesktopFilesWritten ||
		preview.MIMEAppsWritten ||
		preview.ShellConfigurationWritten ||
		preview.SettingsPersisted ||
		preview.KRunnerIndexPersisted ||
		preview.TaskManagerEntryActive ||
		preview.KWinRuleApplied ||
		preview.LiveTrayBridgeEnabled ||
		preview.TrayBridgePersisted ||
		preview.AdapterInvocationEnabled ||
		preview.BackendLaunchEnabled ||
		preview.BackendProcessStarted ||
		preview.SupportBundleExported ||
		preview.SupportCaseCreated ||
		preview.SnapshotRestoreExecuted ||
		preview.StateCleanupExecuted ||
		preview.FileContentRead ||
		preview.FilePathsExposed ||
		preview.NetworkRequired ||
		preview.HostRootModified ||
		preview.PrivilegedContainerRequired ||
		preview.StateRootPathExposed ||
		preview.RawCommandExposed ||
		preview.RawExecutableExposed ||
		preview.BackendDetailsExposed {
		t.Fatalf("unsafe production receipt notification action dry-run result visibility gates: %#v", preview)
	}
	if err := validateNoBackendTerms(preview, "production receipt notification action dry-run result visibility test"); err != nil {
		t.Fatalf("validateNoBackendTerms returned error: %v", err)
	}
}

func TestProductionReceiptNotificationActionDryRunResultVisibilityAuditPreviewFailsClosedWithoutSources(t *testing.T) {
	root := t.TempDir()
	if err := os.WriteFile(filepath.Join(root, "VERSION"), []byte(currentProjectVersion(t)+"\n"), 0o600); err != nil {
		t.Fatalf("WriteFile VERSION returned error: %v", err)
	}
	preview, err := NewProductionReceiptNotificationActionDryRunResultVisibilityAuditPreview(root)
	if err != nil {
		t.Fatalf("NewProductionReceiptNotificationActionDryRunResultVisibilityAuditPreview returned error: %v", err)
	}
	if preview.DispatchDryRunAuditConsumed ||
		preview.ResultVisibilityGuidanceConsumed ||
		!preview.ResultVisibilityAuditRequired ||
		!preview.ResultVisibilityAuditModeled ||
		preview.ResultVisibilityPlanReady ||
		preview.RedactedKDEVisibilityModeled ||
		preview.RuntimeDiagnosticsVisibilityModeled ||
		preview.DispatchDryRunExecuted ||
		preview.DryRunResultPersisted ||
		preview.DispatchAuthorizationGranted ||
		!preview.OperatorDispatchApprovalRequired ||
		preview.OperatorDispatchApprovalPresent ||
		preview.ReceiptPresent ||
		preview.ReceiptAccepted ||
		preview.AuthorizationAccepted ||
		preview.ProductionReadiness ||
		preview.ProductionOwnershipReady ||
		preview.VisibilityItemCount != 5 ||
		preview.RequiredVisibilityItemCount != 5 ||
		preview.ReadyVisibilityItemCount != 0 ||
		preview.MissingVisibilityItemCount != 5 ||
		preview.Counts.Total != 9 ||
		preview.Counts.Passed != 4 ||
		preview.Counts.Blocked != 5 ||
		preview.AuditDecision != "production-receipt-notification-action-dry-run-result-visibility-audit-blocked" {
		t.Fatalf("missing dry-run result visibility sources must fail closed: %#v", preview)
	}
	for _, item := range preview.VisibilityItems {
		if item.EvidencePresent ||
			item.ResultVisibilityModeled ||
			item.RedactedForKDE ||
			item.RuntimeDiagnosticsModeled ||
			item.UserVisible ||
			item.OperatorApprovalPresent ||
			item.DispatchAuthorizationGranted ||
			item.DispatchDryRunExecuted ||
			item.DryRunResultPersisted ||
			item.ResultVisibilityPersisted ||
			item.RequestObjectCreated ||
			item.RequestObjectDispatched ||
			item.RequestObjectPersisted ||
			item.PortalRequestCreated ||
			item.NavigationRequested ||
			item.NotificationActionEnabled ||
			item.ActionEnabled ||
			item.ReceiptAccepted ||
			item.AuthorizationAccepted ||
			item.CompatibilityCenterOpened ||
			item.CompatibilityCenterPersisted ||
			item.RuntimeDiagnosticsPersisted ||
			item.SupportBundleExported ||
			item.SupportCaseCreated ||
			item.ProductionReadiness ||
			item.ProductionOwnershipReady ||
			item.VisibilityStatus != "missing-dry-run-result-visibility-evidence" {
			t.Fatalf("missing source dry-run result visibility item must remain closed: %#v", item)
		}
	}
}
