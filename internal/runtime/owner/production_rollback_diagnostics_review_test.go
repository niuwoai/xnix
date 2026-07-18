package owner

import (
	"os"
	"path/filepath"
	"testing"
)

func TestProductionRollbackDiagnosticsReviewPreviewConsumesSafetySurfaces(t *testing.T) {
	preview, err := NewProductionRollbackDiagnosticsReviewPreview(projectRootForOwnerTest(t))
	if err != nil {
		t.Fatalf("NewProductionRollbackDiagnosticsReviewPreview returned error: %v", err)
	}

	if preview.Version != currentProjectVersion(t) ||
		preview.SchemaVersion != "xnix.runtime.production_rollback_diagnostics_review.v1" ||
		preview.RequestType != "production-rollback-diagnostics-review-preview" ||
		preview.ReviewType != "production-dbus-rollback-diagnostics-review" ||
		preview.ReviewDecision != "production-rollback-diagnostics-review-ready-side-effects-disabled" {
		t.Fatalf("unexpected production rollback diagnostics review schema: %#v", preview)
	}
	if preview.ReviewItemCount != 8 ||
		preview.RollbackControlCount != 6 ||
		preview.DiagnosticsControlCount != 4 ||
		preview.ReadyControlCount != 8 ||
		preview.SideEffectControlCount != 0 ||
		preview.ProductionOwnershipReady {
		t.Fatalf("unexpected rollback diagnostics review counts: %#v", preview)
	}
	if !sameStrings(preview.ItemIDs, []string{
		"production-bus-rollback-boundary",
		"service-activation-fail-closed-boundary",
		"method-exposure-freeze-boundary",
		"runtime-write-freeze-boundary",
		"support-bundle-redaction-boundary",
		"support-case-timeline-boundary",
		"snapshot-restore-candidate-boundary",
		"state-root-retention-dry-run-boundary",
	}) {
		t.Fatalf("unexpected rollback diagnostics review item ids: %#v", preview.ItemIDs)
	}
	for _, item := range preview.Items {
		if !item.RequiredBeforeProduction ||
			!item.EvidencePresent ||
			!item.UserVisible ||
			!item.ReviewOnly ||
			item.SideEffectsEnabled ||
			item.RestoreExecuted ||
			item.CleanupExecuted ||
			item.SupportBundleExported ||
			item.SupportCaseCreated ||
			item.NotificationSent ||
			item.FileContentRead ||
			item.FilePathsExposed ||
			item.StateRootPathExposed ||
			item.RawCommandExposed ||
			item.RawExecutableExposed ||
			item.BackendDetailsExposed ||
			item.HostRootModified ||
			item.ProductionBusClaimed ||
			item.WriteMethodsEnabled ||
			item.RuntimeWritesEnabled ||
			item.BackendLaunchEnabled ||
			item.ReviewStatus != "reviewed" {
			t.Fatalf("unsafe rollback diagnostics review item: %#v", item)
		}
	}
	if preview.Counts.Total != 8 ||
		preview.Counts.Passed != 8 ||
		preview.Counts.Pending != 0 ||
		preview.Counts.Blocked != 0 ||
		!sameStrings(preview.CheckIDs, []string{"production-gate-consumed", "method-review-consumed", "rollback-controls-reviewed", "diagnostics-controls-reviewed", "support-side-effects-disabled", "restore-and-cleanup-disabled", "unsafe-data-hidden", "host-boundary-closed"}) {
		t.Fatalf("unexpected rollback diagnostics review checks: %#v ids=%#v", preview.Counts, preview.CheckIDs)
	}
	if preview.SystemServiceStarted ||
		preview.SessionBusClaimed ||
		preview.ProductionBusClaimed ||
		preview.ProductionOwnerEnabled ||
		preview.ProductionActivationReady ||
		preview.WriteMethodsEnabled ||
		preview.RuntimeWritesEnabled ||
		preview.AdapterInvocationEnabled ||
		preview.BackendLaunchEnabled ||
		preview.BackendProcessStarted ||
		preview.SupportBundleExported ||
		preview.SupportCaseCreated ||
		preview.NotificationSent ||
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
		t.Fatalf("unsafe rollback diagnostics review flag: %#v", preview)
	}
	if err := validateNoBackendTerms(preview, "production rollback diagnostics review test"); err != nil {
		t.Fatalf("validateNoBackendTerms returned error: %v", err)
	}
}

func TestProductionRollbackDiagnosticsReviewPreviewFailsClosedWithoutSources(t *testing.T) {
	root := t.TempDir()
	if err := os.WriteFile(filepath.Join(root, "VERSION"), []byte(currentProjectVersion(t)+"\n"), 0o600); err != nil {
		t.Fatalf("WriteFile VERSION returned error: %v", err)
	}

	preview, err := NewProductionRollbackDiagnosticsReviewPreview(root)
	if err != nil {
		t.Fatalf("NewProductionRollbackDiagnosticsReviewPreview returned error: %v", err)
	}

	if preview.ReviewDecision != "production-rollback-diagnostics-review-blocked" ||
		preview.ProductionOwnershipReady ||
		preview.ProductionBusClaimed ||
		preview.SupportBundleExported ||
		preview.SupportCaseCreated ||
		preview.SnapshotRestoreExecuted ||
		preview.StateCleanupExecuted ||
		preview.HostRootModified ||
		preview.BackendDetailsExposed {
		t.Fatalf("missing sources must keep rollback diagnostics review closed: %#v", preview)
	}
	if preview.ReviewItemCount != 8 ||
		preview.RollbackControlCount != 6 ||
		preview.DiagnosticsControlCount != 4 ||
		preview.ReadyControlCount != 0 ||
		preview.SideEffectControlCount != 0 {
		t.Fatalf("missing sources should preserve review shape: %#v", preview)
	}
	for _, item := range preview.Items {
		if item.EvidencePresent || item.ReviewStatus != "blocked" {
			t.Fatalf("missing-source item must fail closed: %#v", item)
		}
	}
	if preview.Checks[0].ID != "production-gate-consumed" ||
		preview.Checks[0].Status != "blocked" ||
		preview.Checks[1].ID != "method-review-consumed" ||
		preview.Checks[1].Status != "blocked" ||
		preview.Checks[2].ID != "rollback-controls-reviewed" ||
		preview.Checks[2].Status != "blocked" ||
		preview.Checks[3].ID != "diagnostics-controls-reviewed" ||
		preview.Checks[3].Status != "blocked" ||
		preview.Counts.Total != 8 ||
		preview.Counts.Passed != 4 ||
		preview.Counts.Pending != 0 ||
		preview.Counts.Blocked != 4 {
		t.Fatalf("unexpected missing-source rollback diagnostics review checks: %#v ids=%#v", preview.Checks, preview.CheckIDs)
	}
}
