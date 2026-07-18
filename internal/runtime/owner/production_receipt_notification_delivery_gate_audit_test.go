package owner

import (
	"os"
	"path/filepath"
	"testing"
)

func TestProductionReceiptNotificationDeliveryGateAuditPreviewModelsGate(t *testing.T) {
	preview, err := NewProductionReceiptNotificationDeliveryGateAuditPreview(projectRootForOwnerTest(t))
	if err != nil {
		t.Fatalf("NewProductionReceiptNotificationDeliveryGateAuditPreview returned error: %v", err)
	}

	if preview.Version != currentProjectVersion(t) ||
		preview.SchemaVersion != "xnix.runtime.production_receipt_notification_delivery_gate_audit.v1" ||
		preview.RequestType != "production-receipt-notification-delivery-gate-audit-preview" ||
		preview.AuditType != "receipt-expiry-revocation-notification-delivery-gate-audit" ||
		preview.AuditDecision != "production-receipt-notification-delivery-gate-audit-ready-delivery-disabled" ||
		preview.ReceiptSchema != "xnix.runtime.production_dbus_human_authorization_receipt.v1" ||
		preview.OpaqueReceiptID != ProductionDBusHumanAuthorizationReceiptID {
		t.Fatalf("unexpected production receipt notification delivery schema: %#v", preview)
	}
	if !preview.NotificationGateRequired ||
		!preview.NotificationGateModeled ||
		!preview.RevocationVisibilityAuditConsumed ||
		!preview.DesktopSideEffectReviewConsumed ||
		!preview.OperatorNotificationApprovalRequired ||
		preview.OperatorNotificationApprovalPresent ||
		!preview.NotificationDeliveryGateReady ||
		preview.CallerStateRootRequired ||
		preview.ReceiptPresent ||
		preview.ReceiptAccepted ||
		preview.AuthorizationAccepted ||
		preview.ProductionReadiness ||
		preview.ProductionOwnershipReady {
		t.Fatalf("unexpected production receipt notification delivery decision: %#v", preview)
	}
	if preview.GateItemCount != 5 ||
		preview.RequiredGateItemCount != 5 ||
		preview.ReadyGateItemCount != 5 ||
		preview.MissingGateItemCount != 0 ||
		preview.DeliveryEnabledItemCount != 0 ||
		preview.NotificationSentItemCount != 0 ||
		preview.RequestCreatedItemCount != 0 ||
		preview.SideEffectItemCount != 0 ||
		!sameStrings(preview.GateItemIDs, []string{"expiring-receipt-warning-gate", "expired-receipt-blocker-gate", "revoked-receipt-blocker-gate", "missing-review-reminder-gate", "operator-approval-gate"}) {
		t.Fatalf("unexpected production receipt notification delivery inventory: %#v", preview)
	}
	for _, item := range preview.GateItems {
		if !item.EvidencePresent ||
			!item.GateModeled ||
			!item.UserVisible ||
			!item.ReviewOnly ||
			!item.RuntimeOwned ||
			!item.GoRuntimeBacked ||
			item.KDEPolicyOwner ||
			!item.OperatorApprovalRequired ||
			item.OperatorApprovalPresent ||
			item.NotificationSent ||
			item.NotificationDeliveryEnabled ||
			item.NotificationActionEnabled ||
			item.PortalRequestCreated ||
			item.RequestObjectsCreated ||
			item.ReceiptRevocationWriteEnabled ||
			item.ReceiptExpiryWriteEnabled ||
			item.ReceiptPersistenceEnabled ||
			item.ReceiptLookupWritesEnabled ||
			item.ReceiptAccepted ||
			item.AuthorizationAccepted ||
			item.ProductionReadiness ||
			item.ProductionOwnershipReady ||
			!item.SideEffectsDisabled ||
			item.HostRootModified ||
			item.InternalDetailsExposed ||
			item.GateStatus != "notification-gate-modeled-delivery-disabled" {
			t.Fatalf("unsafe or missing production receipt notification delivery item: %#v", item)
		}
	}
	if preview.Counts.Total != 9 ||
		preview.Counts.Passed != 9 ||
		preview.Counts.Pending != 0 ||
		preview.Counts.Blocked != 0 ||
		!sameStrings(preview.CheckIDs, []string{"revocation-visibility-audit-consumed", "desktop-side-effect-review-consumed", "notification-gate-modeled-only", "five-notification-gate-items-present", "gate-items-ready-delivery-disabled", "notification-delivery-disabled", "requests-and-persistence-disabled", "receipt-writes-disabled", "production-and-host-boundary-closed"}) {
		t.Fatalf("unexpected production receipt notification delivery checks: counts=%#v ids=%#v", preview.Counts, preview.CheckIDs)
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
		preview.ReceiptRevocationVisibilityPersisted ||
		preview.NotificationDeliveryGatePersisted ||
		preview.DesktopFilesWritten ||
		preview.MIMEAppsWritten ||
		preview.ShellConfigurationWritten ||
		preview.SettingsPersisted ||
		preview.KRunnerIndexPersisted ||
		preview.TaskManagerEntryActive ||
		preview.KWinRuleApplied ||
		preview.LiveTrayBridgeEnabled ||
		preview.TrayBridgePersisted ||
		preview.NotificationSent ||
		preview.NotificationDeliveryEnabled ||
		preview.NotificationActionEnabled ||
		preview.CompatibilityCenterPersisted ||
		preview.PortalRequestCreated ||
		preview.RequestObjectsCreated ||
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
		t.Fatalf("unsafe production receipt notification delivery gates: %#v", preview)
	}
	if err := validateNoBackendTerms(preview, "production receipt notification delivery test"); err != nil {
		t.Fatalf("validateNoBackendTerms returned error: %v", err)
	}
}

func TestProductionReceiptNotificationDeliveryGateAuditPreviewFailsClosedWithoutSources(t *testing.T) {
	root := t.TempDir()
	if err := os.WriteFile(filepath.Join(root, "VERSION"), []byte(currentProjectVersion(t)+"\n"), 0o600); err != nil {
		t.Fatalf("WriteFile VERSION returned error: %v", err)
	}
	preview, err := NewProductionReceiptNotificationDeliveryGateAuditPreview(root)
	if err != nil {
		t.Fatalf("NewProductionReceiptNotificationDeliveryGateAuditPreview returned error: %v", err)
	}
	if preview.RevocationVisibilityAuditConsumed ||
		preview.DesktopSideEffectReviewConsumed ||
		!preview.NotificationGateRequired ||
		!preview.NotificationGateModeled ||
		!preview.OperatorNotificationApprovalRequired ||
		preview.OperatorNotificationApprovalPresent ||
		preview.NotificationDeliveryGateReady ||
		preview.ReceiptPresent ||
		preview.ReceiptAccepted ||
		preview.AuthorizationAccepted ||
		preview.ProductionReadiness ||
		preview.ProductionOwnershipReady ||
		preview.GateItemCount != 5 ||
		preview.RequiredGateItemCount != 5 ||
		preview.ReadyGateItemCount != 0 ||
		preview.MissingGateItemCount != 5 ||
		preview.Counts.Total != 9 ||
		preview.Counts.Passed != 4 ||
		preview.Counts.Blocked != 5 ||
		preview.AuditDecision != "production-receipt-notification-delivery-gate-audit-blocked" {
		t.Fatalf("missing notification delivery sources must fail closed: %#v", preview)
	}
	for _, item := range preview.GateItems {
		if item.EvidencePresent ||
			item.GateModeled ||
			item.UserVisible ||
			item.OperatorApprovalPresent ||
			item.NotificationSent ||
			item.NotificationDeliveryEnabled ||
			item.NotificationActionEnabled ||
			item.PortalRequestCreated ||
			item.RequestObjectsCreated ||
			item.ReceiptRevocationWriteEnabled ||
			item.ReceiptExpiryWriteEnabled ||
			item.ReceiptPersistenceEnabled ||
			item.ReceiptLookupWritesEnabled ||
			item.ReceiptAccepted ||
			item.AuthorizationAccepted ||
			item.ProductionReadiness ||
			item.ProductionOwnershipReady ||
			item.GateStatus != "missing-notification-gate-evidence" {
			t.Fatalf("missing source notification item must remain closed: %#v", item)
		}
	}
}
