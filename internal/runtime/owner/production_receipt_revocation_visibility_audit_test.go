package owner

import (
	"os"
	"path/filepath"
	"testing"
)

func TestProductionReceiptRevocationVisibilityAuditPreviewModelsVisibility(t *testing.T) {
	preview, err := NewProductionReceiptRevocationVisibilityAuditPreview(projectRootForOwnerTest(t))
	if err != nil {
		t.Fatalf("NewProductionReceiptRevocationVisibilityAuditPreview returned error: %v", err)
	}

	if preview.Version != currentProjectVersion(t) ||
		preview.SchemaVersion != "xnix.runtime.production_receipt_revocation_visibility_audit.v1" ||
		preview.RequestType != "production-receipt-revocation-visibility-audit-preview" ||
		preview.AuditType != "receipt-expiry-revocation-production-kde-visibility-audit" ||
		preview.AuditDecision != "production-receipt-revocation-visibility-audit-ready-visibility-only" ||
		preview.ReceiptSchema != "xnix.runtime.production_dbus_human_authorization_receipt.v1" ||
		preview.OpaqueReceiptID != ProductionDBusHumanAuthorizationReceiptID {
		t.Fatalf("unexpected production receipt revocation visibility schema: %#v", preview)
	}
	if !preview.VisibilityAuditRequired ||
		!preview.VisibilityAuditModeled ||
		!preview.PersistenceThreatReviewConsumed ||
		!preview.DesktopSideEffectReviewConsumed ||
		!preview.ProductionGateVisibilityModeled ||
		!preview.KDEStatusVisibilityModeled ||
		preview.CallerStateRootRequired ||
		preview.ReceiptPresent ||
		preview.ReceiptAccepted ||
		preview.AuthorizationAccepted ||
		!preview.RevocationVisibilityReady ||
		preview.ProductionReadiness ||
		preview.ProductionOwnershipReady {
		t.Fatalf("unexpected production receipt revocation visibility decision: %#v", preview)
	}
	if preview.VisibilityItemCount != 5 ||
		preview.RequiredVisibilityItemCount != 5 ||
		preview.ReadyVisibilityItemCount != 5 ||
		preview.MissingVisibilityItemCount != 0 ||
		preview.RevocationWriteEnabledItemCount != 0 ||
		preview.ExpiryWriteEnabledItemCount != 0 ||
		preview.NotificationEnabledItemCount != 0 ||
		preview.SideEffectItemCount != 0 ||
		!sameStrings(preview.VisibilityItemIDs, []string{"receipt-current-visible", "receipt-expiring-visible", "receipt-expired-visible", "receipt-revoked-visible", "receipt-missing-review-visible"}) {
		t.Fatalf("unexpected production receipt revocation visibility inventory: %#v", preview)
	}
	for _, item := range preview.VisibilityItems {
		if !item.EvidencePresent ||
			!item.VisibilityModeled ||
			!item.UserVisible ||
			!item.ProductionGateVisible ||
			!item.KDEStatusVisible ||
			!item.ReviewOnly ||
			!item.RuntimeOwned ||
			!item.GoRuntimeBacked ||
			item.KDEPolicyOwner ||
			item.ReceiptRevocationWriteEnabled ||
			item.ReceiptExpiryWriteEnabled ||
			item.ReceiptPersistenceEnabled ||
			item.ReceiptLookupWritesEnabled ||
			item.ReceiptReplayEnabled ||
			item.ReceiptAccepted ||
			item.AuthorizationAccepted ||
			item.NotificationSent ||
			item.NotificationDeliveryEnabled ||
			item.DesktopFilesWritten ||
			item.SettingsPersisted ||
			item.ProductionReadiness ||
			item.ProductionOwnershipReady ||
			!item.SideEffectsDisabled ||
			item.HostRootModified ||
			item.InternalDetailsExposed ||
			item.VisibilityStatus != "visibility-modeled-writes-disabled" {
			t.Fatalf("unsafe or missing production receipt revocation visibility item: %#v", item)
		}
	}
	if preview.Counts.Total != 9 ||
		preview.Counts.Passed != 9 ||
		preview.Counts.Pending != 0 ||
		preview.Counts.Blocked != 0 ||
		!sameStrings(preview.CheckIDs, []string{"persistence-threat-review-consumed", "desktop-side-effect-review-consumed", "revocation-visibility-modeled-only", "five-visibility-items-present", "visibility-items-ready-writes-disabled", "revocation-expiry-writes-disabled", "notifications-and-desktop-side-effects-disabled", "production-ownership-disabled", "host-boundary-closed"}) {
		t.Fatalf("unexpected production receipt revocation visibility checks: counts=%#v ids=%#v", preview.Counts, preview.CheckIDs)
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
		t.Fatalf("unsafe production receipt revocation visibility gates: %#v", preview)
	}
	if err := validateNoBackendTerms(preview, "production receipt revocation visibility test"); err != nil {
		t.Fatalf("validateNoBackendTerms returned error: %v", err)
	}
}

func TestProductionReceiptRevocationVisibilityAuditPreviewFailsClosedWithoutSources(t *testing.T) {
	root := t.TempDir()
	if err := os.WriteFile(filepath.Join(root, "VERSION"), []byte(currentProjectVersion(t)+"\n"), 0o600); err != nil {
		t.Fatalf("WriteFile VERSION returned error: %v", err)
	}
	preview, err := NewProductionReceiptRevocationVisibilityAuditPreview(root)
	if err != nil {
		t.Fatalf("NewProductionReceiptRevocationVisibilityAuditPreview returned error: %v", err)
	}
	if preview.PersistenceThreatReviewConsumed ||
		preview.DesktopSideEffectReviewConsumed ||
		preview.ProductionGateVisibilityModeled ||
		preview.KDEStatusVisibilityModeled ||
		!preview.VisibilityAuditRequired ||
		!preview.VisibilityAuditModeled ||
		preview.ReceiptPresent ||
		preview.ReceiptAccepted ||
		preview.AuthorizationAccepted ||
		preview.RevocationVisibilityReady ||
		preview.ProductionReadiness ||
		preview.ProductionOwnershipReady ||
		preview.VisibilityItemCount != 5 ||
		preview.RequiredVisibilityItemCount != 5 ||
		preview.ReadyVisibilityItemCount != 0 ||
		preview.MissingVisibilityItemCount != 5 ||
		preview.Counts.Total != 9 ||
		preview.Counts.Passed != 4 ||
		preview.Counts.Blocked != 5 ||
		preview.AuditDecision != "production-receipt-revocation-visibility-audit-blocked" {
		t.Fatalf("missing revocation visibility sources must fail closed: %#v", preview)
	}
	for _, item := range preview.VisibilityItems {
		if item.EvidencePresent ||
			item.VisibilityModeled ||
			item.UserVisible ||
			item.ProductionGateVisible ||
			item.KDEStatusVisible ||
			item.ReceiptRevocationWriteEnabled ||
			item.ReceiptExpiryWriteEnabled ||
			item.ReceiptPersistenceEnabled ||
			item.ReceiptLookupWritesEnabled ||
			item.ReceiptReplayEnabled ||
			item.ReceiptAccepted ||
			item.AuthorizationAccepted ||
			item.NotificationSent ||
			item.NotificationDeliveryEnabled ||
			item.ProductionReadiness ||
			item.ProductionOwnershipReady ||
			item.VisibilityStatus != "missing-visibility-evidence" {
			t.Fatalf("missing source visibility item must remain closed: %#v", item)
		}
	}
}
