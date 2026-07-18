package owner

import (
	"os"
	"path/filepath"
	"testing"
)

func TestProductionDesktopSideEffectReviewPreviewInventoriesKDESurfaces(t *testing.T) {
	preview, err := NewProductionDesktopSideEffectReviewPreview(projectRootForOwnerTest(t))
	if err != nil {
		t.Fatalf("NewProductionDesktopSideEffectReviewPreview returned error: %v", err)
	}

	if preview.Version != currentProjectVersion(t) ||
		preview.SchemaVersion != "xnix.runtime.production_desktop_side_effect_review.v1" ||
		preview.RequestType != "production-desktop-side-effect-review-preview" ||
		preview.ReviewType != "kde-production-desktop-side-effect-review" ||
		preview.ReviewDecision != "production-desktop-side-effect-review-ready-side-effects-disabled" {
		t.Fatalf("unexpected production desktop side-effect review schema: %#v", preview)
	}
	if preview.SurfaceCount != 7 ||
		preview.RequiredSurfaceCount != 7 ||
		preview.ReviewedSurfaceCount != 7 ||
		preview.ActiveSurfaceCount != 0 ||
		preview.SideEffectSurfaceCount != 0 ||
		preview.ProductionOwnershipReady {
		t.Fatalf("unexpected production desktop side-effect review counts: %#v", preview)
	}
	if !sameStrings(preview.SurfaceIDs, []string{"launcher", "task-manager", "file-manager", "system-tray", "notifications", "compatibility-center", "unified-settings"}) {
		t.Fatalf("unexpected desktop side-effect surface ids: %#v", preview.SurfaceIDs)
	}
	for _, surface := range preview.Surfaces {
		if !surface.RequiredBeforeProduction ||
			!surface.EvidencePresent ||
			!surface.RuntimeOwned ||
			!surface.GoRuntimeBacked ||
			surface.KDEPolicyOwner ||
			!surface.UserVisible ||
			!surface.ReviewOnly ||
			surface.Active ||
			surface.SideEffectsEnabled ||
			surface.DesktopFilesWritten ||
			surface.MIMEAppsWritten ||
			surface.ShellConfigurationWritten ||
			surface.SettingsPersisted ||
			surface.KRunnerIndexPersisted ||
			surface.TaskManagerEntryActive ||
			surface.KWinRuleApplied ||
			surface.LiveTrayBridgeEnabled ||
			surface.NotificationSent ||
			surface.NotificationDeliveryEnabled ||
			surface.CompatibilityCenterPersisted ||
			surface.PortalRequestCreated ||
			surface.RequestObjectsCreated ||
			surface.BackendLaunchEnabled ||
			surface.HostRootModified ||
			surface.BackendDetailsExposed ||
			surface.ReviewStatus != "reviewed" {
			t.Fatalf("unsafe desktop side-effect surface: %#v", surface)
		}
	}
	if preview.Counts.Total != 8 ||
		preview.Counts.Passed != 8 ||
		preview.Counts.Pending != 0 ||
		preview.Counts.Blocked != 0 ||
		!sameStrings(preview.CheckIDs, []string{"production-gate-consumed", "rollback-diagnostics-consumed", "seven-kde-surfaces-reviewed", "runtime-policy-owner", "desktop-writes-disabled", "live-shell-activation-disabled", "notifications-and-requests-disabled", "host-boundary-closed"}) {
		t.Fatalf("unexpected desktop side-effect review checks: %#v ids=%#v", preview.Counts, preview.CheckIDs)
	}
	if preview.PlasmaForkRequired ||
		preview.PlasmaSourceModified ||
		preview.SystemServiceStarted ||
		preview.SessionBusClaimed ||
		preview.ProductionBusClaimed ||
		preview.ProductionOwnerEnabled ||
		preview.ProductionActivationReady ||
		preview.WriteMethodsEnabled ||
		preview.RuntimeWritesEnabled ||
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
		preview.FileContentRead ||
		preview.FilePathsExposed ||
		preview.NetworkRequired ||
		preview.HostRootModified ||
		preview.PrivilegedContainerRequired ||
		preview.StateRootPathExposed ||
		preview.RawCommandExposed ||
		preview.RawExecutableExposed ||
		preview.BackendDetailsExposed {
		t.Fatalf("unsafe production desktop side-effect flag: %#v", preview)
	}
	if err := validateNoBackendTerms(preview, "production desktop side-effect review test"); err != nil {
		t.Fatalf("validateNoBackendTerms returned error: %v", err)
	}
}

func TestProductionDesktopSideEffectReviewPreviewFailsClosedWithoutSources(t *testing.T) {
	root := t.TempDir()
	if err := os.WriteFile(filepath.Join(root, "VERSION"), []byte(currentProjectVersion(t)+"\n"), 0o600); err != nil {
		t.Fatalf("WriteFile VERSION returned error: %v", err)
	}

	preview, err := NewProductionDesktopSideEffectReviewPreview(root)
	if err != nil {
		t.Fatalf("NewProductionDesktopSideEffectReviewPreview returned error: %v", err)
	}

	if preview.ReviewDecision != "production-desktop-side-effect-review-blocked" ||
		preview.ProductionOwnershipReady ||
		preview.ProductionBusClaimed ||
		preview.DesktopFilesWritten ||
		preview.SettingsPersisted ||
		preview.NotificationSent ||
		preview.LiveTrayBridgeEnabled ||
		preview.HostRootModified ||
		preview.BackendDetailsExposed {
		t.Fatalf("missing sources must keep desktop side-effect review closed: %#v", preview)
	}
	if preview.SurfaceCount != 7 ||
		preview.RequiredSurfaceCount != 7 ||
		preview.ReviewedSurfaceCount != 0 ||
		preview.ActiveSurfaceCount != 0 ||
		preview.SideEffectSurfaceCount != 0 {
		t.Fatalf("missing sources should preserve review shape: %#v", preview)
	}
	for _, surface := range preview.Surfaces {
		if surface.EvidencePresent || surface.ReviewStatus != "blocked" {
			t.Fatalf("missing-source surface must fail closed: %#v", surface)
		}
	}
	if preview.Checks[0].ID != "production-gate-consumed" ||
		preview.Checks[0].Status != "blocked" ||
		preview.Checks[1].ID != "rollback-diagnostics-consumed" ||
		preview.Checks[1].Status != "blocked" ||
		preview.Checks[2].ID != "seven-kde-surfaces-reviewed" ||
		preview.Checks[2].Status != "blocked" ||
		preview.Counts.Total != 8 ||
		preview.Counts.Passed != 5 ||
		preview.Counts.Pending != 0 ||
		preview.Counts.Blocked != 3 {
		t.Fatalf("unexpected missing-source desktop side-effect review checks: %#v ids=%#v", preview.Checks, preview.CheckIDs)
	}
}
