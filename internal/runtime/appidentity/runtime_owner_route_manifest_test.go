package appidentity

import (
	"os"
	"path/filepath"
	"testing"
)

func TestRuntimeOwnerRouteManifestPreviewReportsGoAndLegacyRoutes(t *testing.T) {
	preview, err := NewRuntimeOwnerRouteManifestPreview(projectRootForRuntimeServiceBindingTest(t))
	if err != nil {
		t.Fatalf("NewRuntimeOwnerRouteManifestPreview returned error: %v", err)
	}

	if preview.Version != "0.2.190" ||
		preview.SchemaVersion != "xnix.runtime.owner_route_manifest.v1" ||
		preview.RequestType != "runtime-owner-route-manifest-preview" ||
		preview.ManifestType != "runtime-owner-route-manifest" ||
		preview.Source != "runtime-method-parity-manifest-preview+go-runtime-cli+c-runtime-core+runtime-dispatch" ||
		preview.RuntimeMethod != "GetRuntimeOwnerRouteManifest" ||
		preview.ReadMethod != "GetRuntimeOwnerRouteManifestPreview" {
		t.Fatalf("unexpected Runtime owner route manifest schema: %#v", preview)
	}
	if preview.BusName != "org.xnix.Compatibility1" ||
		preview.ObjectPath != "/org/xnix/Compatibility1" ||
		preview.Interface != "org.xnix.Compatibility1" {
		t.Fatalf("unexpected Runtime D-Bus identity: %#v", preview)
	}
	if preview.RouteCounts.Total != len(runtimeMethodParityReadOnlyMethods) ||
		preview.RouteCounts.GoRouted != len(runtimeOwnerRouteGoCommands) ||
		preview.RouteCounts.CCoreBacked != len(runtimeOwnerRouteCCoreCommands) ||
		preview.RouteCounts.RubyLegacy != 0 ||
		preview.RouteCounts.Ready != len(runtimeOwnerRouteGoCommands) ||
		preview.RouteCounts.Pending != len(runtimeOwnerRouteCCoreCommands) ||
		preview.RouteCounts.Blocked != 0 {
		t.Fatalf("unexpected route counts: %#v", preview.RouteCounts)
	}
	if !preview.MethodParityReady ||
		preview.GoOwnerRouteCoverageReady ||
		!preview.CCoreAdapterRequired ||
		preview.LegacyRuntimeRoutesPresent ||
		preview.ProductionOwnerRoutesReady ||
		!preview.RuntimeOwned ||
		!preview.GoRuntimeBacked ||
		preview.KDEPolicyOwner ||
		preview.KDEMayClaimRuntimeOwnership ||
		preview.SystemServiceStarted ||
		preview.ProductionBusClaimed ||
		preview.WriteMethodsEnabled ||
		preview.NetworkRequired ||
		preview.HostRootModified ||
		preview.PrivilegedContainerRequired ||
		preview.BackendDetailsExposed {
		t.Fatalf("unexpected route manifest safety flags: %#v", preview)
	}
	expectedIDs := []string{"method-parity", "go-route-coverage", "c-core-adapter-boundary", "ruby-legacy-dispatch", "write-route-gate", "host-safety-boundary"}
	expectedStatuses := []string{"pass", "pending", "pending", "pass", "pass", "pass"}
	if len(preview.Checks) != len(expectedIDs) || len(preview.CheckIDs) != len(expectedIDs) {
		t.Fatalf("unexpected checks: %#v ids=%#v", preview.Checks, preview.CheckIDs)
	}
	for index, id := range expectedIDs {
		if preview.Checks[index].ID != id ||
			preview.CheckIDs[index] != id ||
			preview.Checks[index].Status != expectedStatuses[index] {
			t.Fatalf("unexpected route manifest check at %d: %#v ids=%#v", index, preview.Checks, preview.CheckIDs)
		}
	}
	if preview.Counts.Total != 6 ||
		preview.Counts.Passed != 4 ||
		preview.Counts.Pending != 2 ||
		preview.Counts.Blocked != 0 {
		t.Fatalf("unexpected check counts: %#v", preview.Counts)
	}

	assertRuntimeOwnerRoute(t, preview, "GetRuntimeServiceBinding", "go-runtime-cli", "runtime-service-binding-preview", "go-preview-ready")
	assertRuntimeOwnerRoute(t, preview, "ListApplications", "go-runtime-cli", "applications-preview", "go-preview-ready")
	assertRuntimeOwnerRoute(t, preview, "GetApplication", "go-runtime-cli", "application-preview", "go-preview-ready")
	assertRuntimeOwnerRoute(t, preview, "GetRunPlan", "c-runtime-core", "compatibility-run-plan", "c-adapter-pending")
	assertRuntimeOwnerRoute(t, preview, "GetDiagnostics", "go-runtime-cli", "diagnostics-preview", "go-preview-ready")

	if len(preview.BlockedActions) != 6 ||
		preview.BlockedActions[0] != "start production Runtime owner from route manifest preview" ||
		len(preview.NextRequirements) != 4 ||
		preview.NextRequirements[0] != "Complete native Go owner handlers or owner adapters for every non-Go read-only route." {
		t.Fatalf("unexpected blocked actions or next requirements: actions=%#v next=%#v", preview.BlockedActions, preview.NextRequirements)
	}
	if preview.DesktopSafeSummary != "Runtime owner routes have Go coverage for migrated application catalog, diagnostics, and current Go previews, with C adapter migration still pending." {
		t.Fatalf("unexpected route manifest summary: %q", preview.DesktopSafeSummary)
	}
	if err := validateNoBackendTerms(preview, "Runtime owner route manifest preview test"); err != nil {
		t.Fatalf("validateNoBackendTerms returned error: %v", err)
	}
}

func TestRuntimeOwnerRouteManifestPreviewBlocksMissingRouteSources(t *testing.T) {
	root := t.TempDir()
	if err := os.WriteFile(filepath.Join(root, "VERSION"), []byte("0.2.190\n"), 0o600); err != nil {
		t.Fatalf("WriteFile VERSION returned error: %v", err)
	}

	preview, err := NewRuntimeOwnerRouteManifestPreview(root)
	if err != nil {
		t.Fatalf("NewRuntimeOwnerRouteManifestPreview returned error: %v", err)
	}

	if preview.MethodParityReady ||
		preview.GoOwnerRouteCoverageReady ||
		preview.ProductionOwnerRoutesReady ||
		preview.SystemServiceStarted ||
		preview.ProductionBusClaimed ||
		preview.WriteMethodsEnabled ||
		preview.HostRootModified ||
		preview.BackendDetailsExposed {
		t.Fatalf("missing route sources must keep owner routes closed: %#v", preview)
	}
	if preview.RouteCounts.Total != len(runtimeMethodParityReadOnlyMethods) ||
		preview.RouteCounts.Ready != 0 ||
		preview.RouteCounts.Pending != 0 ||
		preview.RouteCounts.Blocked != len(runtimeMethodParityReadOnlyMethods) {
		t.Fatalf("unexpected missing route counts: %#v", preview.RouteCounts)
	}
	if preview.Checks[0].ID != "method-parity" ||
		preview.Checks[0].Status != "blocked" ||
		preview.Counts.Blocked != 1 {
		t.Fatalf("missing method parity evidence must block route checks: checks=%#v counts=%#v", preview.Checks, preview.Counts)
	}
	if preview.DesktopSafeSummary != "Runtime owner routes are blocked by missing read-only route evidence." {
		t.Fatalf("unexpected missing route manifest summary: %q", preview.DesktopSafeSummary)
	}
}

func assertRuntimeOwnerRoute(t *testing.T, preview RuntimeOwnerRouteManifestPreview, method string, source string, command string, status string) {
	t.Helper()
	for _, route := range preview.Routes {
		if route.Method != method {
			continue
		}
		if route.CurrentSource != source || route.RouteStatus != status {
			t.Fatalf("unexpected route for %s: %#v", method, route)
		}
		if route.GoCommand != "" && route.GoCommand != command {
			t.Fatalf("unexpected Go command for %s: %#v", method, route)
		}
		if route.CCoreCommand != "" && route.CCoreCommand != command {
			t.Fatalf("unexpected C command for %s: %#v", method, route)
		}
		if route.LegacyDispatchMethod != "" && route.LegacyDispatchMethod != command {
			t.Fatalf("unexpected legacy method for %s: %#v", method, route)
		}
		return
	}
	t.Fatalf("missing route for %s: %#v", method, preview.Routes)
}
