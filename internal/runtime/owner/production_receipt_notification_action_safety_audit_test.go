package owner

import (
	"os"
	"path/filepath"
	"testing"
)

func TestProductionReceiptNotificationActionSafetyAuditPreviewModelsActions(t *testing.T) {
	preview, err := NewProductionReceiptNotificationActionSafetyAuditPreview(projectRootForOwnerTest(t))
	if err != nil {
		t.Fatalf("NewProductionReceiptNotificationActionSafetyAuditPreview returned error: %v", err)
	}

	if preview.Version != currentProjectVersion(t) ||
		preview.SchemaVersion != "xnix.runtime.production_receipt_notification_action_safety_audit.v1" ||
		preview.RequestType != "production-receipt-notification-action-safety-audit-preview" ||
		preview.AuditType != "receipt-notification-action-request-safety-audit" ||
		preview.AuditDecision != "production-receipt-notification-action-safety-audit-ready-actions-disabled" ||
		preview.ReceiptSchema != "xnix.runtime.production_dbus_human_authorization_receipt.v1" ||
		preview.OpaqueReceiptID != ProductionDBusHumanAuthorizationReceiptID {
		t.Fatalf("unexpected production receipt notification action schema: %#v", preview)
	}
	if !preview.ActionSafetyAuditRequired ||
		!preview.ActionSafetyAuditModeled ||
		!preview.NotificationDeliveryGateConsumed ||
		!preview.DesktopSideEffectReviewConsumed ||
		!preview.OperatorActionApprovalRequired ||
		preview.OperatorActionApprovalPresent ||
		!preview.NotificationActionSafetyReady ||
		preview.CallerStateRootRequired ||
		preview.ReceiptPresent ||
		preview.ReceiptAccepted ||
		preview.AuthorizationAccepted ||
		preview.ProductionReadiness ||
		preview.ProductionOwnershipReady {
		t.Fatalf("unexpected production receipt notification action decision: %#v", preview)
	}
	if preview.ActionItemCount != 5 ||
		preview.RequiredActionItemCount != 5 ||
		preview.ReadyActionItemCount != 5 ||
		preview.MissingActionItemCount != 0 ||
		preview.EnabledActionItemCount != 0 ||
		preview.RequestCreatedItemCount != 0 ||
		preview.NavigationEnabledItemCount != 0 ||
		preview.SideEffectItemCount != 0 ||
		!sameStrings(preview.ActionItemIDs, []string{"review-receipt-action", "renew-receipt-action", "open-compatibility-center-action", "dismiss-receipt-action", "support-info-action"}) {
		t.Fatalf("unexpected production receipt notification action inventory: %#v", preview)
	}
	for _, item := range preview.ActionItems {
		if !item.EvidencePresent ||
			!item.ActionSafetyModeled ||
			!item.UserVisible ||
			!item.ReviewOnly ||
			!item.RuntimeOwned ||
			!item.GoRuntimeBacked ||
			item.KDEPolicyOwner ||
			!item.OperatorApprovalRequired ||
			item.OperatorApprovalPresent ||
			item.ActionEnabled ||
			item.NavigationEnabled ||
			item.PortalRequestCreated ||
			item.RequestObjectsCreated ||
			item.NotificationSent ||
			item.NotificationDeliveryEnabled ||
			item.NotificationActionEnabled ||
			item.ReceiptAccepted ||
			item.AuthorizationAccepted ||
			item.ReceiptRevocationWriteEnabled ||
			item.ReceiptExpiryWriteEnabled ||
			item.ReceiptPersistenceEnabled ||
			item.ReceiptLookupWritesEnabled ||
			item.CompatibilityCenterOpened ||
			item.CompatibilityCenterPersisted ||
			item.SupportBundleExported ||
			item.SupportCaseCreated ||
			item.ProductionReadiness ||
			item.ProductionOwnershipReady ||
			!item.SideEffectsDisabled ||
			item.HostRootModified ||
			item.InternalDetailsExposed ||
			item.ActionStatus != "notification-action-modeled-actions-disabled" {
			t.Fatalf("unsafe or missing production receipt notification action item: %#v", item)
		}
	}
	if preview.Counts.Total != 9 ||
		preview.Counts.Passed != 9 ||
		preview.Counts.Pending != 0 ||
		preview.Counts.Blocked != 0 ||
		!sameStrings(preview.CheckIDs, []string{"notification-delivery-gate-consumed", "desktop-side-effect-review-consumed", "action-safety-modeled-only", "five-action-items-present", "action-items-ready-actions-disabled", "notification-actions-disabled", "requests-and-navigation-disabled", "receipt-and-notification-writes-disabled", "production-and-host-boundary-closed"}) {
		t.Fatalf("unexpected production receipt notification action checks: counts=%#v ids=%#v", preview.Counts, preview.CheckIDs)
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
		preview.ReceiptWriterEnabled ||
		preview.ReceiptPersistenceEnabled ||
		preview.ReceiptLookupWritesEnabled ||
		preview.ReceiptReplayEnabled ||
		preview.ReceiptExpiryWriteEnabled ||
		preview.ReceiptRevocationWriteEnabled ||
		preview.NotificationSent ||
		preview.NotificationDeliveryEnabled ||
		preview.NotificationActionEnabled ||
		preview.NotificationActionSafetyPersisted ||
		preview.ReviewActionEnabled ||
		preview.RenewActionEnabled ||
		preview.OpenCompatibilityCenterEnabled ||
		preview.DismissActionEnabled ||
		preview.SupportInfoActionEnabled ||
		preview.CompatibilityCenterOpened ||
		preview.CompatibilityCenterPersisted ||
		preview.PortalRequestCreated ||
		preview.RequestObjectsCreated ||
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
		t.Fatalf("unsafe production receipt notification action gates: %#v", preview)
	}
	if err := validateNoBackendTerms(preview, "production receipt notification action test"); err != nil {
		t.Fatalf("validateNoBackendTerms returned error: %v", err)
	}
}

func TestProductionReceiptNotificationActionSafetyAuditPreviewFailsClosedWithoutSources(t *testing.T) {
	root := t.TempDir()
	if err := os.WriteFile(filepath.Join(root, "VERSION"), []byte(currentProjectVersion(t)+"\n"), 0o600); err != nil {
		t.Fatalf("WriteFile VERSION returned error: %v", err)
	}
	preview, err := NewProductionReceiptNotificationActionSafetyAuditPreview(root)
	if err != nil {
		t.Fatalf("NewProductionReceiptNotificationActionSafetyAuditPreview returned error: %v", err)
	}
	if preview.NotificationDeliveryGateConsumed ||
		preview.DesktopSideEffectReviewConsumed ||
		!preview.ActionSafetyAuditRequired ||
		!preview.ActionSafetyAuditModeled ||
		!preview.OperatorActionApprovalRequired ||
		preview.OperatorActionApprovalPresent ||
		preview.NotificationActionSafetyReady ||
		preview.ReceiptPresent ||
		preview.ReceiptAccepted ||
		preview.AuthorizationAccepted ||
		preview.ProductionReadiness ||
		preview.ProductionOwnershipReady ||
		preview.ActionItemCount != 5 ||
		preview.RequiredActionItemCount != 5 ||
		preview.ReadyActionItemCount != 0 ||
		preview.MissingActionItemCount != 5 ||
		preview.Counts.Total != 9 ||
		preview.Counts.Passed != 4 ||
		preview.Counts.Blocked != 5 ||
		preview.AuditDecision != "production-receipt-notification-action-safety-audit-blocked" {
		t.Fatalf("missing notification action sources must fail closed: %#v", preview)
	}
	for _, item := range preview.ActionItems {
		if item.EvidencePresent ||
			item.ActionSafetyModeled ||
			item.UserVisible ||
			item.OperatorApprovalPresent ||
			item.ActionEnabled ||
			item.NavigationEnabled ||
			item.PortalRequestCreated ||
			item.RequestObjectsCreated ||
			item.NotificationSent ||
			item.NotificationDeliveryEnabled ||
			item.NotificationActionEnabled ||
			item.ReceiptAccepted ||
			item.AuthorizationAccepted ||
			item.CompatibilityCenterOpened ||
			item.CompatibilityCenterPersisted ||
			item.SupportBundleExported ||
			item.SupportCaseCreated ||
			item.ProductionReadiness ||
			item.ProductionOwnershipReady ||
			item.ActionStatus != "missing-notification-action-evidence" {
			t.Fatalf("missing source notification action item must remain closed: %#v", item)
		}
	}
}
