package owner

import (
	"os"
	"path/filepath"
	"testing"
)

func TestProductionReceiptNotificationActionDispatchDryRunAuditPreviewModelsDryRun(t *testing.T) {
	preview, err := NewProductionReceiptNotificationActionDispatchDryRunAuditPreview(projectRootForOwnerTest(t))
	if err != nil {
		t.Fatalf("NewProductionReceiptNotificationActionDispatchDryRunAuditPreview returned error: %v", err)
	}

	if preview.Version != currentProjectVersion(t) ||
		preview.SchemaVersion != "xnix.runtime.production_receipt_notification_action_dispatch_dry_run_audit.v1" ||
		preview.RequestType != "production-receipt-notification-action-dispatch-dry-run-audit-preview" ||
		preview.AuditType != "receipt-notification-action-request-dispatch-dry-run-audit" ||
		preview.AuditDecision != "production-receipt-notification-action-dispatch-dry-run-audit-ready-dry-run-only" ||
		preview.ReceiptSchema != "xnix.runtime.production_dbus_human_authorization_receipt.v1" ||
		preview.OpaqueReceiptID != ProductionDBusHumanAuthorizationReceiptID {
		t.Fatalf("unexpected production receipt notification action dispatch dry-run schema: %#v", preview)
	}
	if !preview.DispatchDryRunAuditRequired ||
		!preview.DispatchDryRunAuditModeled ||
		!preview.DispatchAuthorizationAuditConsumed ||
		!preview.DispatchDryRunGuidanceConsumed ||
		!preview.DispatchDryRunPlanReady ||
		preview.DispatchDryRunExecuted ||
		preview.DispatchAuthorizationGranted ||
		!preview.OperatorDispatchApprovalRequired ||
		preview.OperatorDispatchApprovalPresent ||
		preview.CallerStateRootRequired ||
		preview.ReceiptPresent ||
		preview.ReceiptAccepted ||
		preview.AuthorizationAccepted ||
		preview.ProductionReadiness ||
		preview.ProductionOwnershipReady {
		t.Fatalf("unexpected production receipt notification action dispatch dry-run decision: %#v", preview)
	}
	if preview.DryRunItemCount != 5 ||
		preview.RequiredDryRunItemCount != 5 ||
		preview.ReadyDryRunItemCount != 5 ||
		preview.MissingDryRunItemCount != 0 ||
		preview.ExecutedDryRunItemCount != 0 ||
		preview.DispatchedDryRunItemCount != 0 ||
		preview.CreatedRequestObjectCount != 0 ||
		preview.PortalRequestCreatedCount != 0 ||
		preview.SideEffectDryRunItemCount != 0 ||
		!sameStrings(preview.DryRunItemIDs, []string{"review-receipt-dispatch-dry-run", "renew-receipt-dispatch-dry-run", "open-compatibility-center-dispatch-dry-run", "dismiss-receipt-dispatch-dry-run", "support-info-dispatch-dry-run"}) {
		t.Fatalf("unexpected production receipt notification action dispatch dry-run inventory: %#v", preview)
	}
	for _, item := range preview.DryRunItems {
		if !item.EvidencePresent ||
			!item.DispatchDryRunModeled ||
			!item.UserVisible ||
			!item.ReviewOnly ||
			!item.RuntimeOwned ||
			!item.GoRuntimeBacked ||
			item.KDEPolicyOwner ||
			!item.OperatorApprovalRequired ||
			item.OperatorApprovalPresent ||
			item.DispatchAuthorizationGranted ||
			item.DispatchDryRunExecuted ||
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
			item.SupportBundleExported ||
			item.SupportCaseCreated ||
			item.ProductionReadiness ||
			item.ProductionOwnershipReady ||
			!item.SideEffectsDisabled ||
			item.HostRootModified ||
			item.InternalDetailsExposed ||
			item.DryRunStatus != "dispatch-dry-run-modeled-execution-disabled" {
			t.Fatalf("unsafe or missing production receipt notification action dispatch dry-run item: %#v", item)
		}
	}
	if preview.Counts.Total != 9 ||
		preview.Counts.Passed != 9 ||
		preview.Counts.Pending != 0 ||
		preview.Counts.Blocked != 0 ||
		!sameStrings(preview.CheckIDs, []string{"dispatch-authorization-audit-consumed", "dispatch-dry-run-guidance-consumed", "dispatch-dry-run-modeled-only", "five-dispatch-dry-run-items-present", "dispatch-dry-run-items-ready-execution-disabled", "dry-run-execution-disabled", "request-creation-dispatch-actions-and-portal-disabled", "receipt-notification-navigation-and-support-writes-disabled", "production-and-host-boundary-closed"}) {
		t.Fatalf("unexpected production receipt notification action dispatch dry-run checks: counts=%#v ids=%#v", preview.Counts, preview.CheckIDs)
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
		preview.DryRunResultPersisted ||
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
		t.Fatalf("unsafe production receipt notification action dispatch dry-run gates: %#v", preview)
	}
	if err := validateNoBackendTerms(preview, "production receipt notification action dispatch dry-run test"); err != nil {
		t.Fatalf("validateNoBackendTerms returned error: %v", err)
	}
}

func TestProductionReceiptNotificationActionDispatchDryRunAuditPreviewFailsClosedWithoutSources(t *testing.T) {
	root := t.TempDir()
	if err := os.WriteFile(filepath.Join(root, "VERSION"), []byte(currentProjectVersion(t)+"\n"), 0o600); err != nil {
		t.Fatalf("WriteFile VERSION returned error: %v", err)
	}
	preview, err := NewProductionReceiptNotificationActionDispatchDryRunAuditPreview(root)
	if err != nil {
		t.Fatalf("NewProductionReceiptNotificationActionDispatchDryRunAuditPreview returned error: %v", err)
	}
	if preview.DispatchAuthorizationAuditConsumed ||
		preview.DispatchDryRunGuidanceConsumed ||
		!preview.DispatchDryRunAuditRequired ||
		!preview.DispatchDryRunAuditModeled ||
		preview.DispatchDryRunPlanReady ||
		preview.DispatchDryRunExecuted ||
		preview.DispatchAuthorizationGranted ||
		!preview.OperatorDispatchApprovalRequired ||
		preview.OperatorDispatchApprovalPresent ||
		preview.ReceiptPresent ||
		preview.ReceiptAccepted ||
		preview.AuthorizationAccepted ||
		preview.ProductionReadiness ||
		preview.ProductionOwnershipReady ||
		preview.DryRunItemCount != 5 ||
		preview.RequiredDryRunItemCount != 5 ||
		preview.ReadyDryRunItemCount != 0 ||
		preview.MissingDryRunItemCount != 5 ||
		preview.Counts.Total != 9 ||
		preview.Counts.Passed != 4 ||
		preview.Counts.Blocked != 5 ||
		preview.AuditDecision != "production-receipt-notification-action-dispatch-dry-run-audit-blocked" {
		t.Fatalf("missing dispatch dry-run sources must fail closed: %#v", preview)
	}
	for _, item := range preview.DryRunItems {
		if item.EvidencePresent ||
			item.DispatchDryRunModeled ||
			item.UserVisible ||
			item.OperatorApprovalPresent ||
			item.DispatchAuthorizationGranted ||
			item.DispatchDryRunExecuted ||
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
			item.SupportBundleExported ||
			item.SupportCaseCreated ||
			item.ProductionReadiness ||
			item.ProductionOwnershipReady ||
			item.DryRunStatus != "missing-dispatch-dry-run-evidence" {
			t.Fatalf("missing source dispatch dry-run item must remain closed: %#v", item)
		}
	}
}
