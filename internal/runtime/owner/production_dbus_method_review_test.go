package owner

import (
	"os"
	"path/filepath"
	"testing"
)

func TestProductionDBusMethodReviewPreviewInventoriesRoutes(t *testing.T) {
	preview, err := NewProductionDBusMethodReviewPreview(projectRootForOwnerTest(t))
	if err != nil {
		t.Fatalf("NewProductionDBusMethodReviewPreview returned error: %v", err)
	}

	if preview.Version != currentProjectVersion(t) ||
		preview.SchemaVersion != "xnix.runtime.production_dbus_method_review.v1" ||
		preview.RequestType != "production-dbus-method-review-preview" ||
		preview.ReviewType != "route-by-route-production-dbus-method-review" ||
		preview.Source != "runtime-method-parity-manifest-preview+runtime-owner-route-manifest-preview+production-dbus-gate-review-preview+runtime-write-gate-preview+production-human-authorization-receipt-consolidation-preview" ||
		preview.ReviewDecision != "production-dbus-method-review-ready-production-exposure-disabled" {
		t.Fatalf("unexpected production D-Bus method review schema: %#v", preview)
	}
	if preview.ReadOnlyContractMethodCount != 61 ||
		preview.OwnerLocalCandidateCount != 3 ||
		preview.WriteMethodCount != 4 ||
		preview.ReviewedMethodCount != 68 ||
		preview.ProductionExposureReadyCount != 0 ||
		preview.NewProductionMethodRequestCount != 0 {
		t.Fatalf("unexpected production D-Bus method review counts: %#v", preview)
	}
	for _, route := range preview.ReadOnlyMethods {
		if route.RouteClass != "dbus-read-only-contract" ||
			route.CurrentExposure != "read-only-contract-method-production-owner-disabled" ||
			route.FutureExposureDecision != "reviewed-read-only-contract-production-owner-disabled" ||
			!route.ContractMethodPresent ||
			route.OwnerLocalCandidate ||
			!route.GoRouteReady ||
			!route.SmokeCoverageReady ||
			!route.ReadOnly ||
			route.WriteMethod ||
			route.NewProductionMethodRequested ||
			route.ProductionExposureReady ||
			route.ProductionOwnerEnabled ||
			route.WriteMethodsEnabled ||
			route.RuntimeWritesEnabled ||
			route.BackendLaunchEnabled ||
			route.HostRootModified ||
			route.InternalDetailsExposed ||
			route.ReviewStatus != "reviewed" {
			t.Fatalf("unexpected read-only method route: %#v", route)
		}
	}
	for _, route := range preview.OwnerLocalCandidates {
		if route.RouteClass != "owner-local-candidate" ||
			route.CurrentExposure != "owner-local-only" ||
			route.FutureExposureDecision != "owner-local-only-no-production-dbus-method" ||
			route.ContractMethodPresent ||
			!route.OwnerLocalCandidate ||
			!route.GoRouteReady ||
			!route.SmokeCoverageReady ||
			!route.ReadOnly ||
			route.WriteMethod ||
			route.NewProductionMethodRequested ||
			route.ProductionExposureReady ||
			route.WriteMethodsEnabled ||
			route.RuntimeWritesEnabled ||
			route.BackendLaunchEnabled ||
			route.HostRootModified ||
			route.InternalDetailsExposed ||
			route.ReviewStatus != "reviewed" {
			t.Fatalf("unexpected owner-local candidate route: %#v", route)
		}
	}
	for _, route := range preview.WriteMethods {
		if route.RouteClass != "reserved-write-method" ||
			route.CurrentExposure != "write-method-disabled" ||
			route.FutureExposureDecision != "write-method-disabled-no-production-dispatch" ||
			!route.ContractMethodPresent ||
			route.OwnerLocalCandidate ||
			route.GoRouteReady ||
			route.SmokeCoverageReady ||
			route.ReadOnly ||
			!route.WriteMethod ||
			route.NewProductionMethodRequested ||
			route.ProductionExposureReady ||
			route.WriteMethodsEnabled ||
			route.RuntimeWritesEnabled ||
			route.BackendLaunchEnabled ||
			route.HostRootModified ||
			route.InternalDetailsExposed ||
			route.ReviewStatus != "reviewed-disabled" {
			t.Fatalf("unexpected write method route: %#v", route)
		}
	}
	if preview.Counts.Total != 7 ||
		preview.Counts.Passed != 7 ||
		preview.Counts.Pending != 0 ||
		preview.Counts.Blocked != 0 ||
		!sameStrings(preview.CheckIDs, []string{"read-only-contract-methods-reviewed", "owner-local-candidates-reviewed", "write-methods-reviewed-disabled", "no-new-production-methods-requested", "production-exposure-disabled", "route-manifest-consumed", "unsafe-gates-closed"}) {
		t.Fatalf("unexpected production D-Bus method review checks: %#v ids=%#v", preview.Counts, preview.CheckIDs)
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
		preview.NotificationSent ||
		preview.NetworkRequired ||
		preview.HostRootModified ||
		preview.PrivilegedContainerRequired ||
		preview.StateRootPathExposed ||
		preview.RawCommandExposed ||
		preview.RawExecutableExposed ||
		preview.BackendDetailsExposed {
		t.Fatalf("unsafe production D-Bus method review flag: %#v", preview)
	}
	if err := validateNoBackendTerms(preview, "production D-Bus method review test"); err != nil {
		t.Fatalf("validateNoBackendTerms returned error: %v", err)
	}
}

func TestProductionDBusMethodReviewPreviewFailsClosedWithoutSources(t *testing.T) {
	root := t.TempDir()
	if err := os.WriteFile(filepath.Join(root, "VERSION"), []byte(currentProjectVersion(t)+"\n"), 0o600); err != nil {
		t.Fatalf("WriteFile VERSION returned error: %v", err)
	}

	preview, err := NewProductionDBusMethodReviewPreview(root)
	if err != nil {
		t.Fatalf("NewProductionDBusMethodReviewPreview returned error: %v", err)
	}

	if preview.ReviewDecision != "production-dbus-method-review-blocked" ||
		preview.ProductionBusClaimed ||
		preview.WriteMethodsEnabled ||
		preview.RuntimeWritesEnabled ||
		preview.BackendLaunchEnabled ||
		preview.HostRootModified ||
		preview.BackendDetailsExposed {
		t.Fatalf("missing sources must keep production D-Bus method review closed: %#v", preview)
	}
	if preview.ReadOnlyContractMethodCount != 61 ||
		preview.OwnerLocalCandidateCount != 3 ||
		preview.WriteMethodCount != 4 ||
		preview.ReviewedMethodCount != 68 {
		t.Fatalf("missing sources should preserve review shape: %#v", preview)
	}
	if preview.ReadOnlyMethods[0].ReviewStatus != "blocked" ||
		preview.Checks[0].ID != "read-only-contract-methods-reviewed" ||
		preview.Checks[0].Status != "blocked" ||
		preview.Checks[1].ID != "owner-local-candidates-reviewed" ||
		preview.Checks[1].Status != "blocked" ||
		preview.Checks[5].ID != "route-manifest-consumed" ||
		preview.Checks[5].Status != "blocked" ||
		preview.Counts.Total != 7 ||
		preview.Counts.Passed != 4 ||
		preview.Counts.Pending != 0 ||
		preview.Counts.Blocked != 3 {
		t.Fatalf("unexpected missing-source production D-Bus method review checks: %#v ids=%#v", preview.Checks, preview.CheckIDs)
	}
}
