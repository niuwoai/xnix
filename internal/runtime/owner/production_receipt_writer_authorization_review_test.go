package owner

import (
	"os"
	"path/filepath"
	"testing"
)

func TestProductionReceiptWriterAuthorizationReviewPreviewModelsWriterBoundary(t *testing.T) {
	preview, err := NewProductionReceiptWriterAuthorizationReviewPreview(projectRootForOwnerTest(t))
	if err != nil {
		t.Fatalf("NewProductionReceiptWriterAuthorizationReviewPreview returned error: %v", err)
	}

	if preview.Version != currentProjectVersion(t) ||
		preview.SchemaVersion != "xnix.runtime.production_receipt_writer_authorization_review.v1" ||
		preview.RequestType != "production-receipt-writer-authorization-review-preview" ||
		preview.ReviewType != "receipt-writer-operator-authorization-boundary-review" ||
		preview.ReviewDecision != "production-receipt-writer-authorization-review-ready-writes-disabled" ||
		preview.ReceiptSchema != "xnix.runtime.production_dbus_human_authorization_receipt.v1" ||
		preview.OpaqueReceiptID != ProductionDBusHumanAuthorizationReceiptID {
		t.Fatalf("unexpected production receipt writer authorization review schema: %#v", preview)
	}
	if !preview.ReceiptRequired ||
		preview.ReceiptPresent ||
		preview.ReceiptAccepted ||
		!preview.WriterAuthorizationRequired ||
		!preview.WriterAuthorizationModeled ||
		!preview.OperatorActionRequired ||
		!preview.AcceptancePropagationConsumed ||
		!preview.OwnerManagedOpaqueBoundaryReady ||
		preview.CallerStateRootRequired ||
		preview.AuthorizationAccepted ||
		!preview.WriterReviewReady ||
		preview.ProductionReadiness ||
		preview.ProductionOwnershipReady {
		t.Fatalf("unexpected production receipt writer authorization decision: %#v", preview)
	}
	if preview.ReviewItemCount != 5 ||
		preview.RequiredReviewItemCount != 5 ||
		preview.ReadyReviewItemCount != 5 ||
		preview.MissingReviewItemCount != 0 ||
		preview.WriteEnabledReviewItemCount != 0 ||
		preview.AcceptanceEnabledReviewItemCount != 0 ||
		preview.SideEffectReviewItemCount != 0 ||
		!sameStrings(preview.ReviewItemIDs, []string{"operator-action-boundary", "opaque-boundary-consolidated", "production-gates-consume-boundary", "persistence-disabled", "host-boundary-closed"}) {
		t.Fatalf("unexpected production receipt writer review inventory: counts=%#v ids=%#v", preview.ReviewItemCount, preview.ReviewItemIDs)
	}
	for _, item := range preview.ReviewItems {
		if !item.EvidencePresent ||
			!item.WriterAuthorizationModeled ||
			item.ReceiptWriterEnabled ||
			item.ReceiptPersistenceEnabled ||
			item.ReceiptLookupWritesEnabled ||
			item.ReceiptReplayEnabled ||
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
			item.ReviewStatus != "writer-authorization-reviewed-writes-disabled" {
			t.Fatalf("unsafe or missing production receipt writer review item: %#v", item)
		}
	}
	if preview.Counts.Total != 8 ||
		preview.Counts.Passed != 8 ||
		preview.Counts.Pending != 0 ||
		preview.Counts.Blocked != 0 ||
		!sameStrings(preview.CheckIDs, []string{"acceptance-propagation-consumed", "writer-authorization-modeled-only", "five-review-items-present", "review-items-ready-writes-disabled", "receipt-writes-disabled", "production-ownership-disabled", "runtime-side-effects-disabled", "host-boundary-closed"}) {
		t.Fatalf("unexpected production receipt writer authorization checks: counts=%#v ids=%#v", preview.Counts, preview.CheckIDs)
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
		t.Fatalf("unsafe production receipt writer authorization review gates: %#v", preview)
	}
	if err := validateNoBackendTerms(preview, "production receipt writer authorization review test"); err != nil {
		t.Fatalf("validateNoBackendTerms returned error: %v", err)
	}
}

func TestProductionReceiptWriterAuthorizationReviewPreviewFailsClosedWithoutSources(t *testing.T) {
	root := t.TempDir()
	if err := os.WriteFile(filepath.Join(root, "VERSION"), []byte(currentProjectVersion(t)+"\n"), 0o600); err != nil {
		t.Fatalf("WriteFile VERSION returned error: %v", err)
	}
	preview, err := NewProductionReceiptWriterAuthorizationReviewPreview(root)
	if err != nil {
		t.Fatalf("NewProductionReceiptWriterAuthorizationReviewPreview returned error: %v", err)
	}
	if preview.AcceptancePropagationConsumed ||
		preview.OwnerManagedOpaqueBoundaryReady ||
		!preview.WriterAuthorizationRequired ||
		!preview.WriterAuthorizationModeled ||
		!preview.OperatorActionRequired ||
		preview.ReceiptPresent ||
		preview.ReceiptAccepted ||
		preview.AuthorizationAccepted ||
		preview.WriterReviewReady ||
		preview.ProductionReadiness ||
		preview.ProductionOwnershipReady ||
		preview.ReviewItemCount != 5 ||
		preview.RequiredReviewItemCount != 5 ||
		preview.ReadyReviewItemCount != 0 ||
		preview.MissingReviewItemCount != 5 ||
		preview.Counts.Total != 8 ||
		preview.Counts.Passed != 5 ||
		preview.Counts.Blocked != 3 ||
		preview.ReviewDecision != "production-receipt-writer-authorization-review-blocked" {
		t.Fatalf("missing writer authorization sources must fail closed: %#v", preview)
	}
	for _, item := range preview.ReviewItems {
		if item.EvidencePresent ||
			item.WriterAuthorizationModeled ||
			item.ReceiptWriterEnabled ||
			item.ReceiptPersistenceEnabled ||
			item.ReceiptLookupWritesEnabled ||
			item.ReceiptReplayEnabled ||
			item.ReceiptAccepted ||
			item.AuthorizationAccepted ||
			item.ProductionReadiness ||
			item.ProductionOwnershipReady ||
			item.ReviewStatus != "missing-review-evidence" {
			t.Fatalf("missing source review item must remain closed: %#v", item)
		}
	}
}
