package owner

import (
	"os"
	"path/filepath"
	"testing"
)

func TestProductionReceiptNotificationActionRequestObjectAuditPreviewModelsRequests(t *testing.T) {
	preview, err := NewProductionReceiptNotificationActionRequestObjectAuditPreview(projectRootForOwnerTest(t))
	if err != nil {
		t.Fatalf("NewProductionReceiptNotificationActionRequestObjectAuditPreview returned error: %v", err)
	}

	if preview.Version != currentProjectVersion(t) ||
		preview.SchemaVersion != "xnix.runtime.production_receipt_notification_action_request_object_audit.v1" ||
		preview.RequestType != "production-receipt-notification-action-request-object-audit-preview" ||
		preview.AuditType != "receipt-notification-action-runtime-request-object-audit" ||
		preview.AuditDecision != "production-receipt-notification-action-request-object-audit-ready-requests-disabled" ||
		preview.ReceiptSchema != "xnix.runtime.production_dbus_human_authorization_receipt.v1" ||
		preview.OpaqueReceiptID != ProductionDBusHumanAuthorizationReceiptID {
		t.Fatalf("unexpected production receipt notification action request-object schema: %#v", preview)
	}
	if !preview.RequestObjectAuditRequired ||
		!preview.RequestObjectAuditModeled ||
		!preview.ActionSafetyAuditConsumed ||
		!preview.DispatchRequestBoundaryConsumed ||
		!preview.OperatorActionApprovalRequired ||
		preview.OperatorActionApprovalPresent ||
		!preview.RequestObjectBoundaryReady ||
		preview.CallerStateRootRequired ||
		preview.ReceiptPresent ||
		preview.ReceiptAccepted ||
		preview.AuthorizationAccepted ||
		preview.ProductionReadiness ||
		preview.ProductionOwnershipReady {
		t.Fatalf("unexpected production receipt notification action request-object decision: %#v", preview)
	}
	if preview.RequestObjectCount != 5 ||
		preview.RequiredRequestObjectCount != 5 ||
		preview.ReadyRequestObjectCount != 5 ||
		preview.MissingRequestObjectCount != 0 ||
		preview.CreatedRequestObjectCount != 0 ||
		preview.DispatchedRequestObjectCount != 0 ||
		preview.PortalRequestCreatedCount != 0 ||
		preview.NavigationRequestedCount != 0 ||
		preview.SideEffectRequestObjectCount != 0 ||
		!sameStrings(preview.RequestObjectIDs, []string{"review-receipt-request-object", "renew-receipt-request-object", "open-compatibility-center-request-object", "dismiss-receipt-request-object", "support-info-request-object"}) {
		t.Fatalf("unexpected production receipt notification action request-object inventory: %#v", preview)
	}
	for _, object := range preview.RequestObjects {
		if !object.EvidencePresent ||
			!object.RequestObjectModeled ||
			!object.UserVisible ||
			!object.ReviewOnly ||
			!object.RuntimeOwned ||
			!object.GoRuntimeBacked ||
			object.KDEPolicyOwner ||
			!object.OperatorApprovalRequired ||
			object.OperatorApprovalPresent ||
			object.RequestObjectCreated ||
			object.RequestObjectDispatched ||
			object.RequestObjectPersisted ||
			object.PortalRequestCreated ||
			object.NavigationRequested ||
			object.NotificationActionEnabled ||
			object.ActionEnabled ||
			object.ReceiptAccepted ||
			object.AuthorizationAccepted ||
			object.ReceiptWriterEnabled ||
			object.ReceiptPersistenceEnabled ||
			object.CompatibilityCenterOpened ||
			object.CompatibilityCenterPersisted ||
			object.SupportBundleExported ||
			object.SupportCaseCreated ||
			object.ProductionReadiness ||
			object.ProductionOwnershipReady ||
			!object.SideEffectsDisabled ||
			object.HostRootModified ||
			object.InternalDetailsExposed ||
			object.RequestObjectStatus != "request-object-modeled-creation-disabled" {
			t.Fatalf("unsafe or missing production receipt notification action request-object item: %#v", object)
		}
	}
	if preview.Counts.Total != 9 ||
		preview.Counts.Passed != 9 ||
		preview.Counts.Pending != 0 ||
		preview.Counts.Blocked != 0 ||
		!sameStrings(preview.CheckIDs, []string{"action-safety-audit-consumed", "dispatch-request-boundary-consumed", "request-object-modeled-only", "five-request-objects-present", "request-objects-ready-creation-disabled", "request-creation-disabled", "actions-navigation-and-portal-disabled", "receipt-notification-and-support-writes-disabled", "production-and-host-boundary-closed"}) {
		t.Fatalf("unexpected production receipt notification action request-object checks: counts=%#v ids=%#v", preview.Counts, preview.CheckIDs)
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
		t.Fatalf("unsafe production receipt notification action request-object gates: %#v", preview)
	}
	if err := validateNoBackendTerms(preview, "production receipt notification action request-object test"); err != nil {
		t.Fatalf("validateNoBackendTerms returned error: %v", err)
	}
}

func TestProductionReceiptNotificationActionRequestObjectAuditPreviewFailsClosedWithoutSources(t *testing.T) {
	root := t.TempDir()
	if err := os.WriteFile(filepath.Join(root, "VERSION"), []byte(currentProjectVersion(t)+"\n"), 0o600); err != nil {
		t.Fatalf("WriteFile VERSION returned error: %v", err)
	}
	preview, err := NewProductionReceiptNotificationActionRequestObjectAuditPreview(root)
	if err != nil {
		t.Fatalf("NewProductionReceiptNotificationActionRequestObjectAuditPreview returned error: %v", err)
	}
	if preview.ActionSafetyAuditConsumed ||
		preview.DispatchRequestBoundaryConsumed ||
		!preview.RequestObjectAuditRequired ||
		!preview.RequestObjectAuditModeled ||
		!preview.OperatorActionApprovalRequired ||
		preview.OperatorActionApprovalPresent ||
		preview.RequestObjectBoundaryReady ||
		preview.ReceiptPresent ||
		preview.ReceiptAccepted ||
		preview.AuthorizationAccepted ||
		preview.ProductionReadiness ||
		preview.ProductionOwnershipReady ||
		preview.RequestObjectCount != 5 ||
		preview.RequiredRequestObjectCount != 5 ||
		preview.ReadyRequestObjectCount != 0 ||
		preview.MissingRequestObjectCount != 5 ||
		preview.Counts.Total != 9 ||
		preview.Counts.Passed != 4 ||
		preview.Counts.Blocked != 5 ||
		preview.AuditDecision != "production-receipt-notification-action-request-object-audit-blocked" {
		t.Fatalf("missing request-object sources must fail closed: %#v", preview)
	}
	for _, object := range preview.RequestObjects {
		if object.EvidencePresent ||
			object.RequestObjectModeled ||
			object.UserVisible ||
			object.OperatorApprovalPresent ||
			object.RequestObjectCreated ||
			object.RequestObjectDispatched ||
			object.RequestObjectPersisted ||
			object.PortalRequestCreated ||
			object.NavigationRequested ||
			object.NotificationActionEnabled ||
			object.ActionEnabled ||
			object.ReceiptAccepted ||
			object.AuthorizationAccepted ||
			object.CompatibilityCenterOpened ||
			object.CompatibilityCenterPersisted ||
			object.SupportBundleExported ||
			object.SupportCaseCreated ||
			object.ProductionReadiness ||
			object.ProductionOwnershipReady ||
			object.RequestObjectStatus != "missing-request-object-evidence" {
			t.Fatalf("missing source request-object item must remain closed: %#v", object)
		}
	}
}
