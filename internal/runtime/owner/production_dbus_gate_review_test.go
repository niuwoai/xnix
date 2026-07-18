package owner

import (
	"os"
	"path/filepath"
	"testing"
)

func TestProductionDBusGateReviewPreviewInventoriesSmokeCoveredRoutes(t *testing.T) {
	preview, err := NewProductionDBusGateReviewPreview(projectRootForOwnerTest(t))
	if err != nil {
		t.Fatalf("NewProductionDBusGateReviewPreview returned error: %v", err)
	}

	if preview.Version != currentProjectVersion(t) ||
		preview.SchemaVersion != "xnix.runtime.production_dbus_gate_review.v1" ||
		preview.RequestType != "production-dbus-gate-review-preview" ||
		preview.GateType != "owner-local-smoke-covered-production-dbus-gate-review" ||
		preview.GateDecision != "production-dbus-gate-review-ready" ||
		preview.CurrentGateStatus != "owner-local-smoke-covered-production-dbus-disabled" ||
		preview.RouteCount != 3 ||
		preview.SmokeCoveredRouteCount != 3 ||
		preview.ProductionReadiness ||
		!preview.HumanAuthorizationRequired {
		t.Fatalf("unexpected production D-Bus gate review schema: %#v", preview)
	}
	if !sameStrings(preview.RouteIDs, []string{"materialization-fanout-owner-route", "restricted-smoke-fanout-owner-route", "redacted-adapter-profile-owner-route"}) {
		t.Fatalf("unexpected production D-Bus gate routes: %#v", preview.RouteIDs)
	}
	for _, route := range preview.Routes {
		if !route.SmokeCoverageReady ||
			!route.OwnerLocalRouteReady ||
			route.ProductionDBusMethodPresent ||
			route.ProductionDBusExposureReady ||
			route.WriteMethodsEnabled ||
			route.RuntimeWritesEnabled ||
			route.BackendLaunchEnabled ||
			route.HostRootModified ||
			route.InternalDetailsExposed {
			t.Fatalf("unsafe or uncovered production D-Bus gate route: %#v", route)
		}
	}
	if preview.Counts.Total != 8 ||
		preview.Counts.Passed != 8 ||
		preview.Counts.Pending != 0 ||
		preview.Counts.Blocked != 0 ||
		!sameStrings(preview.CheckIDs, []string{"tracked-routes-present", "smoke-coverage-present", "production-dbus-disabled", "route-production-exposure-disabled", "writes-and-launch-disabled", "desktop-side-effects-disabled", "host-boundary-closed", "human-authorization-required"}) {
		t.Fatalf("unexpected production D-Bus gate review checks: counts=%#v ids=%#v", preview.Counts, preview.CheckIDs)
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
		preview.AdapterInvocationEnabled ||
		preview.BackendLaunchEnabled ||
		preview.BackendProcessStarted ||
		preview.SupportBundleExported ||
		preview.SupportCaseCreated ||
		preview.NotificationSent ||
		preview.NetworkRequired ||
		preview.HostRootModified ||
		preview.PrivilegedContainerRequired ||
		preview.StateRootPathExposed ||
		preview.RawCommandExposed ||
		preview.RawExecutableExposed ||
		preview.BackendDetailsExposed {
		t.Fatalf("unsafe production D-Bus gate review gates: %#v", preview)
	}
	if err := validateNoBackendTerms(preview, "production D-Bus gate review test"); err != nil {
		t.Fatalf("validateNoBackendTerms returned error: %v", err)
	}
}

func TestProductionDBusGateReviewPreviewFailsClosedWithoutSources(t *testing.T) {
	root := t.TempDir()
	if err := os.WriteFile(filepath.Join(root, "VERSION"), []byte(currentProjectVersion(t)+"\n"), 0o600); err != nil {
		t.Fatalf("WriteFile VERSION returned error: %v", err)
	}
	preview, err := NewProductionDBusGateReviewPreview(root)
	if err != nil {
		t.Fatalf("NewProductionDBusGateReviewPreview returned error: %v", err)
	}
	if preview.GateDecision != "production-dbus-gate-review-blocked" ||
		preview.ProductionReadiness ||
		!preview.HumanAuthorizationRequired ||
		preview.SmokeCoveredRouteCount != 0 ||
		preview.Counts.Pending != 1 ||
		preview.Counts.Blocked != 0 {
		t.Fatalf("missing sources must fail closed without production readiness: %#v", preview)
	}
}

func projectRootForOwnerTest(t *testing.T) string {
	t.Helper()
	workingDirectory, err := os.Getwd()
	if err != nil {
		t.Fatalf("Getwd returned error: %v", err)
	}
	for {
		if _, err := os.Stat(filepath.Join(workingDirectory, "VERSION")); err == nil {
			return workingDirectory
		}
		parent := filepath.Dir(workingDirectory)
		if parent == workingDirectory {
			t.Fatalf("could not find project root from %s", workingDirectory)
		}
		workingDirectory = parent
	}
}
