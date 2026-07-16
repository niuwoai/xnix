package appidentity

import (
	"os"
	"path/filepath"
	"testing"
)

func TestRuntimeRouteConvergencePreviewClassifiesReadRoutes(t *testing.T) {
	preview, err := NewRuntimeRouteConvergencePreview(projectRootForRuntimeServiceBindingTest(t))
	if err != nil {
		t.Fatalf("NewRuntimeRouteConvergencePreview returned error: %v", err)
	}

	if preview.Version != "0.2.303" ||
		preview.SchemaVersion != "xnix.runtime.route_convergence.v1" ||
		preview.RequestType != "runtime-route-convergence-preview" ||
		preview.ReportType != "runtime-route-convergence" ||
		preview.Source != "runtime-owner-route-manifest-preview+runtime-method-parity-manifest-preview" ||
		preview.RuntimeMethod != "GetRuntimeOwnerRouteManifest" ||
		preview.ReadMethod != "GetRuntimeRouteConvergencePreview" {
		t.Fatalf("unexpected route convergence schema: %#v", preview)
	}
	if preview.ClassificationCounts.Total != 61 ||
		preview.ClassificationCounts.GoProductLogic != 61 ||
		preview.ClassificationCounts.CPolicyBridge != 0 ||
		preview.ClassificationCounts.RubySmokeBridge != 0 ||
		preview.ClassificationCounts.FixtureOnly != 0 ||
		preview.ClassificationCounts.ContractOnly != 0 ||
		preview.ClassificationCounts.Deprecated != 0 ||
		preview.ClassificationCounts.Unsupported != 0 ||
		preview.ClassificationCounts.Unclassified != 0 {
		t.Fatalf("unexpected classification counts: %#v", preview.ClassificationCounts)
	}
	if !preview.AllRoutesClassified || !preview.NativeGoCoverageReady || preview.ProductionOwnerReady {
		t.Fatalf("unexpected convergence readiness flags: %#v", preview)
	}
	if !preview.RuntimeOwned ||
		!preview.GoRuntimeBacked ||
		preview.KDEPolicyOwner ||
		preview.WriteMethodsEnabled ||
		preview.NetworkRequired ||
		preview.HostRootModified ||
		preview.PrivilegedContainerRequired ||
		preview.BackendLaunchEnabled ||
		preview.BackendDetailsExposed {
		t.Fatalf("unexpected convergence safety flags: %#v", preview)
	}
	if len(preview.Routes) != 61 || len(preview.RouteMethodNames) != 61 {
		t.Fatalf("unexpected route count: routes=%d names=%d", len(preview.Routes), len(preview.RouteMethodNames))
	}
	if len(preview.UnclassifiedRoutes) != 0 {
		t.Fatalf("expected no unclassified routes: %#v", preview.UnclassifiedRoutes)
	}
	if len(preview.MigrationGroups) == 0 || len(preview.NextMigrationOrder) == 0 {
		t.Fatalf("expected migration guidance: groups=%#v order=%#v", preview.MigrationGroups, preview.NextMigrationOrder)
	}

	for _, route := range preview.Routes {
		if route.Method == "" ||
			route.Domain == "" ||
			route.Classification != "go-product-logic" ||
			route.MigrationRisk != "low" ||
			route.MigrationPriority != 4 ||
			len(route.BlockedReasons) != 0 ||
			!route.KDEVisibleReadOnly ||
			!route.SafeForProductionGate {
			t.Fatalf("unexpected converged route: %#v", route)
		}
	}

	expectedChecks := []string{
		"owner-route-manifest-ready",
		"all-routes-classified",
		"native-go-route-coverage",
		"write-gate-disabled",
		"host-safety-boundary",
	}
	if len(preview.Checks) != len(expectedChecks) || len(preview.CheckIDs) != len(expectedChecks) {
		t.Fatalf("unexpected convergence checks: %#v ids=%#v", preview.Checks, preview.CheckIDs)
	}
	for index, id := range expectedChecks {
		if preview.Checks[index].ID != id || preview.CheckIDs[index] != id || preview.Checks[index].Status != "pass" {
			t.Fatalf("unexpected convergence check at %d: %#v ids=%#v", index, preview.Checks, preview.CheckIDs)
		}
	}
}

func TestRuntimeRouteConvergencePreviewFailsClosedForMissingRouteEvidence(t *testing.T) {
	root := t.TempDir()
	if err := os.WriteFile(filepath.Join(root, "VERSION"), []byte("0.2.303\n"), 0o600); err != nil {
		t.Fatalf("WriteFile VERSION returned error: %v", err)
	}

	preview, err := NewRuntimeRouteConvergencePreview(root)
	if err != nil {
		t.Fatalf("NewRuntimeRouteConvergencePreview returned error: %v", err)
	}

	if preview.AllRoutesClassified != true ||
		preview.NativeGoCoverageReady ||
		preview.ProductionOwnerReady ||
		preview.WriteMethodsEnabled ||
		preview.HostRootModified ||
		preview.BackendLaunchEnabled {
		t.Fatalf("missing evidence must fail closed without enabling unsafe gates: %#v", preview)
	}
	if preview.ClassificationCounts.Total != 61 ||
		preview.ClassificationCounts.ContractOnly != 61 ||
		preview.ClassificationCounts.GoProductLogic != 0 ||
		preview.ClassificationCounts.Unclassified != 0 {
		t.Fatalf("missing evidence must become contract-only route evidence: %#v", preview.ClassificationCounts)
	}
	if len(preview.NextMigrationOrder) == 0 {
		t.Fatalf("missing evidence must produce migration order: %#v", preview)
	}
	if preview.Checks[0].ID != "owner-route-manifest-ready" || preview.Checks[0].Status != "blocked" {
		t.Fatalf("missing owner route manifest evidence must block convergence: %#v", preview.Checks)
	}
	if preview.Checks[2].ID != "native-go-route-coverage" || preview.Checks[2].Status != "pending" {
		t.Fatalf("missing native Go coverage must remain pending: %#v", preview.Checks)
	}
}
