package owner

import (
	"os"
	"path/filepath"
	"testing"
)

func TestProductionReceiptPersistenceThreatReviewPreviewModelsThreats(t *testing.T) {
	preview, err := NewProductionReceiptPersistenceThreatReviewPreview(projectRootForOwnerTest(t))
	if err != nil {
		t.Fatalf("NewProductionReceiptPersistenceThreatReviewPreview returned error: %v", err)
	}

	if preview.Version != currentProjectVersion(t) ||
		preview.SchemaVersion != "xnix.runtime.production_receipt_persistence_threat_review.v1" ||
		preview.RequestType != "production-receipt-persistence-threat-review-preview" ||
		preview.ReviewType != "receipt-persistence-expiry-revocation-replay-threat-review" ||
		preview.ReviewDecision != "production-receipt-persistence-threat-review-ready-persistence-disabled" ||
		preview.ReceiptSchema != "xnix.runtime.production_dbus_human_authorization_receipt.v1" ||
		preview.OpaqueReceiptID != ProductionDBusHumanAuthorizationReceiptID {
		t.Fatalf("unexpected production receipt persistence threat review schema: %#v", preview)
	}
	if !preview.ReceiptRequired ||
		preview.ReceiptPresent ||
		preview.ReceiptAccepted ||
		!preview.PersistenceThreatReviewRequired ||
		!preview.PersistenceThreatReviewModeled ||
		!preview.WriterAuthorizationReviewConsumed ||
		!preview.OwnerManagedOpaqueBoundaryReady ||
		preview.CallerStateRootRequired ||
		preview.AuthorizationAccepted ||
		!preview.PersistenceThreatReviewReady ||
		preview.ProductionReadiness ||
		preview.ProductionOwnershipReady {
		t.Fatalf("unexpected production receipt persistence threat decision: %#v", preview)
	}
	if preview.ThreatItemCount != 5 ||
		preview.RequiredThreatItemCount != 5 ||
		preview.ReadyThreatItemCount != 5 ||
		preview.MissingThreatItemCount != 0 ||
		preview.PersistenceEnabledThreatCount != 0 ||
		preview.ReplayEnabledThreatCount != 0 ||
		preview.AcceptanceEnabledThreatCount != 0 ||
		preview.SideEffectThreatCount != 0 ||
		!sameStrings(preview.ThreatItemIDs, []string{"storage-confidentiality", "expiry-policy", "revocation-policy", "replay-protection", "audit-visibility"}) {
		t.Fatalf("unexpected production receipt persistence threat inventory: counts=%#v ids=%#v", preview.ThreatItemCount, preview.ThreatItemIDs)
	}
	for _, item := range preview.ThreatItems {
		if !item.EvidencePresent ||
			!item.ThreatModeled ||
			item.ReceiptPersistenceEnabled ||
			item.ReceiptLookupWritesEnabled ||
			item.ReceiptReplayEnabled ||
			item.ReceiptExpiryWriteEnabled ||
			item.ReceiptRevocationWriteEnabled ||
			item.ReceiptAccepted ||
			item.AuthorizationAccepted ||
			item.ProductionReadiness ||
			item.ProductionOwnershipReady ||
			!item.RuntimeOwned ||
			!item.GoRuntimeBacked ||
			item.KDEPolicyOwner ||
			!item.ReviewOnly ||
			!item.SideEffectsDisabled ||
			item.HostRootModified ||
			item.InternalDetailsExposed ||
			item.ThreatStatus != "threat-modeled-persistence-disabled" {
			t.Fatalf("unsafe or missing production receipt persistence threat item: %#v", item)
		}
	}
	if preview.Counts.Total != 8 ||
		preview.Counts.Passed != 8 ||
		preview.Counts.Pending != 0 ||
		preview.Counts.Blocked != 0 ||
		!sameStrings(preview.CheckIDs, []string{"writer-authorization-review-consumed", "persistence-threat-modeled-only", "five-threat-items-present", "threat-items-ready-persistence-disabled", "receipt-persistence-disabled", "production-ownership-disabled", "runtime-side-effects-disabled", "host-boundary-closed"}) {
		t.Fatalf("unexpected production receipt persistence threat checks: counts=%#v ids=%#v", preview.Counts, preview.CheckIDs)
	}
	if !preview.RuntimeOwned ||
		!preview.GoRuntimeBacked ||
		preview.KDEPolicyOwner ||
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
		preview.DesktopFilesWritten ||
		preview.MIMEAppsWritten ||
		preview.ShellConfigurationWritten ||
		preview.SettingsPersisted ||
		preview.NotificationSent ||
		preview.NotificationDeliveryEnabled ||
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
		t.Fatalf("unsafe production receipt persistence threat review gates: %#v", preview)
	}
	if err := validateNoBackendTerms(preview, "production receipt persistence threat review test"); err != nil {
		t.Fatalf("validateNoBackendTerms returned error: %v", err)
	}
}

func TestProductionReceiptPersistenceThreatReviewPreviewFailsClosedWithoutSources(t *testing.T) {
	root := t.TempDir()
	if err := os.WriteFile(filepath.Join(root, "VERSION"), []byte(currentProjectVersion(t)+"\n"), 0o600); err != nil {
		t.Fatalf("WriteFile VERSION returned error: %v", err)
	}
	preview, err := NewProductionReceiptPersistenceThreatReviewPreview(root)
	if err != nil {
		t.Fatalf("NewProductionReceiptPersistenceThreatReviewPreview returned error: %v", err)
	}
	if preview.WriterAuthorizationReviewConsumed ||
		preview.OwnerManagedOpaqueBoundaryReady ||
		!preview.PersistenceThreatReviewRequired ||
		!preview.PersistenceThreatReviewModeled ||
		preview.ReceiptPresent ||
		preview.ReceiptAccepted ||
		preview.AuthorizationAccepted ||
		preview.PersistenceThreatReviewReady ||
		preview.ProductionReadiness ||
		preview.ProductionOwnershipReady ||
		preview.ThreatItemCount != 5 ||
		preview.RequiredThreatItemCount != 5 ||
		preview.ReadyThreatItemCount != 0 ||
		preview.MissingThreatItemCount != 5 ||
		preview.Counts.Total != 8 ||
		preview.Counts.Passed != 5 ||
		preview.Counts.Blocked != 3 ||
		preview.ReviewDecision != "production-receipt-persistence-threat-review-blocked" {
		t.Fatalf("missing persistence threat sources must fail closed: %#v", preview)
	}
	for _, item := range preview.ThreatItems {
		if item.EvidencePresent ||
			item.ThreatModeled ||
			item.ReceiptPersistenceEnabled ||
			item.ReceiptLookupWritesEnabled ||
			item.ReceiptReplayEnabled ||
			item.ReceiptExpiryWriteEnabled ||
			item.ReceiptRevocationWriteEnabled ||
			item.ReceiptAccepted ||
			item.AuthorizationAccepted ||
			item.ProductionReadiness ||
			item.ProductionOwnershipReady ||
			item.ThreatStatus != "missing-threat-evidence" {
			t.Fatalf("missing source threat item must remain closed: %#v", item)
		}
	}
}
