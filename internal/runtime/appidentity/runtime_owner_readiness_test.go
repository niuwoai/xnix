package appidentity

import (
	"os"
	"path/filepath"
	"testing"
)

func TestRuntimeOwnerReadinessPreviewAggregatesOwnerGates(t *testing.T) {
	preview, err := NewRuntimeOwnerReadinessPreview(projectRootForRuntimeServiceBindingTest(t))
	if err != nil {
		t.Fatalf("NewRuntimeOwnerReadinessPreview returned error: %v", err)
	}

	if preview.Version != "0.2.189" ||
		preview.SchemaVersion != "xnix.runtime.owner_readiness.v1" ||
		preview.RequestType != "runtime-owner-readiness-preview" ||
		preview.ReadinessType != "runtime-owner-readiness" ||
		preview.Source != "runtime-service-binding-preview+runtime-live-owner-gate-preview+runtime-owner-process-preview+runtime-owner-smoke-plan-preview+runtime-method-parity-manifest-preview+runtime-owner-route-manifest-preview+runtime-owner-recipe-trust-preview" ||
		preview.RuntimeMethod != "GetRuntimeOwnerReadiness" ||
		preview.ReadMethod != "GetRuntimeOwnerReadinessPreview" {
		t.Fatalf("unexpected Runtime owner readiness schema: %#v", preview)
	}
	if preview.BusName != "org.xnix.Compatibility1" ||
		preview.ObjectPath != "/org/xnix/Compatibility1" ||
		preview.Interface != "org.xnix.Compatibility1" {
		t.Fatalf("unexpected Runtime D-Bus identity: %#v", preview)
	}
	if preview.ServiceBinding.RequestType != "runtime-service-binding-preview" ||
		preview.ServiceBinding.ProductionStatus != "pending-live-owner" ||
		!preview.ServiceBinding.ActivationBindingReady ||
		preview.ServiceBinding.LiveDBusOwnerReady ||
		!preview.ServiceBinding.SmokeAdapterAvailable ||
		preview.ServiceBinding.Counts.Passed != 4 ||
		preview.ServiceBinding.Counts.Pending != 1 {
		t.Fatalf("unexpected service binding summary: %#v", preview.ServiceBinding)
	}
	if preview.LiveOwnerGate.RequestType != "runtime-live-owner-gate-preview" ||
		preview.LiveOwnerGate.GateType != "runtime-live-owner-gate" ||
		!preview.LiveOwnerGate.ActivationBindingReady ||
		preview.LiveOwnerGate.LiveDBusOwnerReady ||
		preview.LiveOwnerGate.ProductionOwnerEnabled ||
		preview.LiveOwnerGate.OwnerTransitionReady ||
		preview.LiveOwnerGate.SmokeAdapterIsProduction ||
		preview.LiveOwnerGate.Counts.Passed != 1 ||
		preview.LiveOwnerGate.Counts.Pending != 4 {
		t.Fatalf("unexpected live owner gate summary: %#v", preview.LiveOwnerGate)
	}
	if preview.OwnerProcess.RequestType != "runtime-owner-process-preview" ||
		preview.OwnerProcess.ProcessType != "runtime-owner-process" ||
		preview.OwnerProcess.CurrentOwnerLanguage != "ruby-wrapper" ||
		preview.OwnerProcess.TargetOwnerLanguage != "go" ||
		!preview.OwnerProcess.ServiceActivationReady ||
		!preview.OwnerProcess.PackagedEntrypointReady ||
		preview.OwnerProcess.GoOwnerProcessReady ||
		preview.OwnerProcess.ProductionOwnerProcessReady ||
		preview.OwnerProcess.Counts.Passed != 3 ||
		preview.OwnerProcess.Counts.Pending != 2 {
		t.Fatalf("unexpected owner process summary: %#v", preview.OwnerProcess)
	}
	if preview.OwnerSmokePlan.RequestType != "runtime-owner-smoke-plan-preview" ||
		preview.OwnerSmokePlan.PlanType != "runtime-owner-smoke-plan" ||
		preview.OwnerSmokePlan.SmokeState != "planned" ||
		preview.OwnerSmokePlan.SmokeEnvironment != "restricted-session" ||
		!preview.OwnerSmokePlan.ActivationBindingReady ||
		preview.OwnerSmokePlan.PendingStepCount != 6 ||
		preview.OwnerSmokePlan.Counts.Passed != 1 ||
		preview.OwnerSmokePlan.Counts.Pending != 6 {
		t.Fatalf("unexpected owner smoke plan summary: %#v", preview.OwnerSmokePlan)
	}
	if preview.MethodParityManifest.RequestType != "runtime-method-parity-manifest-preview" ||
		preview.MethodParityManifest.ManifestType != "runtime-method-parity-manifest" ||
		!preview.MethodParityManifest.ReadOnlyMethodParityReady ||
		preview.MethodParityManifest.MethodCount != 57 ||
		preview.MethodParityManifest.WriteMethodsSupported ||
		preview.MethodParityManifest.WriteMethodDispatchEnabled ||
		preview.MethodParityManifest.Counts.Passed != 5 {
		t.Fatalf("unexpected method parity summary: %#v", preview.MethodParityManifest)
	}
	if preview.OwnerRouteManifest.RequestType != "runtime-owner-route-manifest-preview" ||
		preview.OwnerRouteManifest.ManifestType != "runtime-owner-route-manifest" ||
		preview.OwnerRouteManifest.RouteCount != 57 ||
		preview.OwnerRouteManifest.GoRouteCount != len(runtimeOwnerRouteGoCommands) ||
		preview.OwnerRouteManifest.CCoreRouteCount != len(runtimeOwnerRouteCCoreCommands) ||
		preview.OwnerRouteManifest.RubyLegacyRouteCount != 0 ||
		preview.OwnerRouteManifest.GoOwnerRouteCoverageReady ||
		!preview.OwnerRouteManifest.CCoreAdapterRequired ||
		preview.OwnerRouteManifest.LegacyRuntimeRoutesPresent ||
		preview.OwnerRouteManifest.ProductionOwnerRoutesReady ||
		preview.OwnerRouteManifest.Counts.Passed != 4 ||
		preview.OwnerRouteManifest.Counts.Pending != 2 {
		t.Fatalf("unexpected owner route manifest summary: %#v", preview.OwnerRouteManifest)
	}
	if preview.RecipeTrust.RequestType != "runtime-owner-recipe-trust-preview" ||
		preview.RecipeTrust.TrustType != "runtime-owner-recipe-trust" ||
		preview.RecipeTrust.RegistryName != "xnix-local-development" ||
		preview.RecipeTrust.RecipeCount != 1 ||
		!preview.RecipeTrust.DigestVerified ||
		preview.RecipeTrust.SignedRecipeValidation ||
		!preview.RecipeTrust.DevelopmentRegistry ||
		preview.RecipeTrust.UnsignedRecipesPresent ||
		preview.RecipeTrust.ProductionRecipeTrustReady ||
		preview.RecipeTrust.Counts.Passed != 3 ||
		preview.RecipeTrust.Counts.Pending != 2 {
		t.Fatalf("unexpected recipe trust summary: %#v", preview.RecipeTrust)
	}

	expectedIDs := []string{
		"activation-binding",
		"read-only-method-parity",
		"owner-smoke-plan",
		"write-method-gate",
		"kde-ownership-boundary",
		"host-safety-boundary",
		"long-running-runtime-owner",
		"read-only-owner-routes",
		"production-bus-claim",
		"production-recipe-trust",
	}
	expectedStatuses := []string{"pass", "pass", "pass", "pass", "pass", "pass", "pending", "pending", "pending", "pending"}
	if len(preview.ReadinessChecks) != len(expectedIDs) || len(preview.CheckIDs) != len(expectedIDs) {
		t.Fatalf("unexpected readiness checks: %#v ids=%#v", preview.ReadinessChecks, preview.CheckIDs)
	}
	for index, id := range expectedIDs {
		if preview.ReadinessChecks[index].ID != id ||
			preview.CheckIDs[index] != id ||
			preview.ReadinessChecks[index].Status != expectedStatuses[index] {
			t.Fatalf("unexpected readiness check at %d: %#v ids=%#v", index, preview.ReadinessChecks, preview.CheckIDs)
		}
	}
	if preview.Counts.Total != 10 ||
		preview.Counts.Passed != 6 ||
		preview.Counts.Pending != 4 ||
		preview.Counts.Blocked != 0 {
		t.Fatalf("unexpected readiness counts: %#v", preview.Counts)
	}
	if !preview.RuntimeOwned ||
		!preview.GoRuntimeBacked ||
		preview.KDEPolicyOwner ||
		preview.KDEMayClaimRuntimeOwnership ||
		!preview.ActivationBindingReady ||
		!preview.ReadOnlyMethodParityReady ||
		!preview.OwnerSmokePlanned ||
		preview.LiveDBusOwnerReady ||
		preview.ProductionOwnerEnabled ||
		preview.OwnerTransitionReady ||
		preview.ProductionRecipeTrustReady ||
		preview.ProductionOwnerRoutesReady ||
		preview.NetworkRequired ||
		preview.HostRootModified ||
		preview.PrivilegedContainerRequired ||
		preview.SystemServiceStarted ||
		preview.ProductionBusClaimed ||
		preview.WriteMethodsEnabled ||
		preview.BackendDetailsExposed {
		t.Fatalf("unexpected owner readiness safety flags: %#v", preview)
	}
	if len(preview.BlockedReasons) != 6 ||
		preview.BlockedReasons[0] != "Live production Runtime ownership is still pending." ||
		len(preview.BlockedActions) != 7 ||
		preview.BlockedActions[0] != "start production Runtime owner from readiness preview" {
		t.Fatalf("unexpected blocked readiness metadata: reasons=%#v actions=%#v", preview.BlockedReasons, preview.BlockedActions)
	}
	if preview.DesktopSafeSummary != "Runtime owner readiness has aligned activation and method parity; C owner adapter migration, live production ownership, bus claim, and production-signed recipe trust remain pending." {
		t.Fatalf("unexpected desktop-safe summary: %q", preview.DesktopSafeSummary)
	}
	if err := validateNoBackendTerms(preview, "Runtime owner readiness preview test"); err != nil {
		t.Fatalf("validateNoBackendTerms returned error: %v", err)
	}
}

func TestRuntimeOwnerReadinessPreviewBlocksMissingActivationAndParity(t *testing.T) {
	root := t.TempDir()
	if err := os.WriteFile(filepath.Join(root, "VERSION"), []byte("0.2.189\n"), 0o600); err != nil {
		t.Fatalf("WriteFile VERSION returned error: %v", err)
	}

	preview, err := NewRuntimeOwnerReadinessPreview(root)
	if err != nil {
		t.Fatalf("NewRuntimeOwnerReadinessPreview returned error: %v", err)
	}

	if preview.ActivationBindingReady ||
		preview.ReadOnlyMethodParityReady ||
		!preview.OwnerSmokePlanned ||
		preview.LiveDBusOwnerReady ||
		preview.ProductionOwnerEnabled ||
		preview.OwnerTransitionReady ||
		preview.ProductionRecipeTrustReady ||
		preview.ProductionOwnerRoutesReady ||
		preview.SystemServiceStarted ||
		preview.ProductionBusClaimed ||
		preview.HostRootModified ||
		preview.BackendDetailsExposed {
		t.Fatalf("missing activation and parity must keep owner readiness closed: %#v", preview)
	}
	if preview.ReadinessChecks[0].ID != "activation-binding" ||
		preview.ReadinessChecks[0].Status != "blocked" ||
		preview.ReadinessChecks[1].ID != "read-only-method-parity" ||
		preview.ReadinessChecks[1].Status != "blocked" {
		t.Fatalf("missing sources must block activation and method parity: %#v", preview.ReadinessChecks)
	}
	if preview.Counts.Total != 10 ||
		preview.Counts.Passed != 4 ||
		preview.Counts.Pending != 1 ||
		preview.Counts.Blocked != 5 {
		t.Fatalf("unexpected missing-source readiness counts: %#v", preview.Counts)
	}
	if preview.ServiceBinding.Counts.Blocked != 4 ||
		preview.OwnerProcess.Counts.Blocked != 2 ||
		preview.MethodParityManifest.Counts.Blocked != 5 ||
		preview.OwnerRouteManifest.Counts.Blocked != 1 ||
		preview.RecipeTrust.Counts.Blocked != 3 {
		t.Fatalf("unexpected missing-source child counts: service=%#v process=%#v method=%#v routes=%#v trust=%#v", preview.ServiceBinding.Counts, preview.OwnerProcess.Counts, preview.MethodParityManifest.Counts, preview.OwnerRouteManifest.Counts, preview.RecipeTrust.Counts)
	}
	if preview.DesktopSafeSummary != "Runtime owner readiness is blocked by activation or method parity defects." {
		t.Fatalf("unexpected missing-source summary: %q", preview.DesktopSafeSummary)
	}
}
