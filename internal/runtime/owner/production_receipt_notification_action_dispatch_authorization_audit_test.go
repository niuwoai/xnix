package owner

import (
	"os"
	"path/filepath"
	"testing"
)

func TestProductionReceiptNotificationActionDispatchAuthorizationAuditPreviewModelsDispatch(t *testing.T) {
	preview, err := NewProductionReceiptNotificationActionDispatchAuthorizationAuditPreview(projectRootForOwnerTest(t))
	if err != nil {
		t.Fatalf("NewProductionReceiptNotificationActionDispatchAuthorizationAuditPreview returned error: %v", err)
	}

	if preview.Version != currentProjectVersion(t) ||
		preview.SchemaVersion != "xnix.runtime.production_receipt_notification_action_dispatch_authorization_audit.v1" ||
		preview.RequestType != "production-receipt-notification-action-dispatch-authorization-audit-preview" ||
		preview.AuditType != "receipt-notification-action-request-dispatch-authorization-audit" ||
		preview.AuditDecision != "production-receipt-notification-action-dispatch-authorization-audit-ready-dispatch-disabled" ||
		preview.ReceiptSchema != "xnix.runtime.production_dbus_human_authorization_receipt.v1" ||
		preview.OpaqueReceiptID != ProductionDBusHumanAuthorizationReceiptID {
		t.Fatalf("unexpected production receipt notification action dispatch authorization schema: %#v", preview)
	}
	if !preview.DispatchAuthorizationAuditRequired ||
		!preview.DispatchAuthorizationAuditModeled ||
		!preview.RequestObjectAuditConsumed ||
		!preview.ReceiptAuthorizationBoundaryConsumed ||
		!preview.OperatorDispatchApprovalRequired ||
		preview.OperatorDispatchApprovalPresent ||
		!preview.DispatchAuthorizationReady ||
		preview.DispatchAuthorizationGranted ||
		preview.CallerStateRootRequired ||
		preview.ReceiptPresent ||
		preview.ReceiptAccepted ||
		preview.AuthorizationAccepted ||
		preview.ProductionReadiness ||
		preview.ProductionOwnershipReady {
		t.Fatalf("unexpected production receipt notification action dispatch authorization decision: %#v", preview)
	}
	if preview.AuthorizationItemCount != 5 ||
		preview.RequiredAuthorizationItemCount != 5 ||
		preview.ReadyAuthorizationItemCount != 5 ||
		preview.MissingAuthorizationItemCount != 0 ||
		preview.GrantedAuthorizationItemCount != 0 ||
		preview.DispatchedAuthorizationItemCount != 0 ||
		preview.CreatedRequestObjectCount != 0 ||
		preview.PortalRequestCreatedCount != 0 ||
		preview.SideEffectAuthorizationItemCount != 0 ||
		!sameStrings(preview.AuthorizationItemIDs, []string{"review-receipt-dispatch-authorization", "renew-receipt-dispatch-authorization", "open-compatibility-center-dispatch-authorization", "dismiss-receipt-dispatch-authorization", "support-info-dispatch-authorization"}) {
		t.Fatalf("unexpected production receipt notification action dispatch authorization inventory: %#v", preview)
	}
	for _, item := range preview.AuthorizationItems {
		if !item.EvidencePresent ||
			!item.DispatchAuthorizationModeled ||
			!item.UserVisible ||
			!item.ReviewOnly ||
			!item.RuntimeOwned ||
			!item.GoRuntimeBacked ||
			item.KDEPolicyOwner ||
			!item.OperatorApprovalRequired ||
			item.OperatorApprovalPresent ||
			item.DispatchAuthorizationGranted ||
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
			item.DispatchAuthorizationStatus != "dispatch-authorization-modeled-dispatch-disabled" {
			t.Fatalf("unsafe or missing production receipt notification action dispatch authorization item: %#v", item)
		}
	}
	if preview.Counts.Total != 9 ||
		preview.Counts.Passed != 9 ||
		preview.Counts.Pending != 0 ||
		preview.Counts.Blocked != 0 ||
		!sameStrings(preview.CheckIDs, []string{"request-object-audit-consumed", "receipt-authorization-boundary-consumed", "dispatch-authorization-modeled-only", "five-dispatch-authorizations-present", "dispatch-authorizations-ready-dispatch-disabled", "dispatch-disabled", "request-creation-actions-and-portal-disabled", "receipt-notification-navigation-and-support-writes-disabled", "production-and-host-boundary-closed"}) {
		t.Fatalf("unexpected production receipt notification action dispatch authorization checks: counts=%#v ids=%#v", preview.Counts, preview.CheckIDs)
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
		t.Fatalf("unsafe production receipt notification action dispatch authorization gates: %#v", preview)
	}
	if err := validateNoBackendTerms(preview, "production receipt notification action dispatch authorization test"); err != nil {
		t.Fatalf("validateNoBackendTerms returned error: %v", err)
	}
}

func TestProductionReceiptNotificationActionDispatchAuthorizationAuditPreviewFailsClosedWithoutSources(t *testing.T) {
	root := t.TempDir()
	if err := os.WriteFile(filepath.Join(root, "VERSION"), []byte(currentProjectVersion(t)+"\n"), 0o600); err != nil {
		t.Fatalf("WriteFile VERSION returned error: %v", err)
	}
	preview, err := NewProductionReceiptNotificationActionDispatchAuthorizationAuditPreview(root)
	if err != nil {
		t.Fatalf("NewProductionReceiptNotificationActionDispatchAuthorizationAuditPreview returned error: %v", err)
	}
	if preview.RequestObjectAuditConsumed ||
		preview.ReceiptAuthorizationBoundaryConsumed ||
		!preview.DispatchAuthorizationAuditRequired ||
		!preview.DispatchAuthorizationAuditModeled ||
		!preview.OperatorDispatchApprovalRequired ||
		preview.OperatorDispatchApprovalPresent ||
		preview.DispatchAuthorizationReady ||
		preview.DispatchAuthorizationGranted ||
		preview.ReceiptPresent ||
		preview.ReceiptAccepted ||
		preview.AuthorizationAccepted ||
		preview.ProductionReadiness ||
		preview.ProductionOwnershipReady ||
		preview.AuthorizationItemCount != 5 ||
		preview.RequiredAuthorizationItemCount != 5 ||
		preview.ReadyAuthorizationItemCount != 0 ||
		preview.MissingAuthorizationItemCount != 5 ||
		preview.Counts.Total != 9 ||
		preview.Counts.Passed != 4 ||
		preview.Counts.Blocked != 5 ||
		preview.AuditDecision != "production-receipt-notification-action-dispatch-authorization-audit-blocked" {
		t.Fatalf("missing dispatch authorization sources must fail closed: %#v", preview)
	}
	for _, item := range preview.AuthorizationItems {
		if item.EvidencePresent ||
			item.DispatchAuthorizationModeled ||
			item.UserVisible ||
			item.OperatorApprovalPresent ||
			item.DispatchAuthorizationGranted ||
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
			item.DispatchAuthorizationStatus != "missing-dispatch-authorization-evidence" {
			t.Fatalf("missing source dispatch authorization item must remain closed: %#v", item)
		}
	}
}
