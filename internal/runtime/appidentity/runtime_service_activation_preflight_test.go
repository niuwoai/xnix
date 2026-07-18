package appidentity

import (
	"os"
	"path/filepath"
	"testing"
)

func TestRuntimeServiceActivationPreflightPreviewGatesProductionActivation(t *testing.T) {
	preview, err := NewRuntimeServiceActivationPreflightPreview(projectRootForRuntimeServiceBindingTest(t))
	if err != nil {
		t.Fatalf("NewRuntimeServiceActivationPreflightPreview returned error: %v", err)
	}

	if preview.Version != currentProjectVersion(t) ||
		preview.SchemaVersion != "xnix.runtime.service_activation_preflight.v1" ||
		preview.RequestType != "runtime-service-activation-preflight-preview" ||
		preview.PreflightType != "production-runtime-service-activation-preflight" ||
		preview.Source != "runtime-service-binding-preview+runtime-owner-readiness-preview+runtime-owner-smoke-plan-preview" ||
		preview.RuntimeMethod != "GetRuntimeServiceActivationPreflight" ||
		preview.ReadMethod != "GetRuntimeServiceActivationPreflightPreview" {
		t.Fatalf("unexpected Runtime service activation preflight schema: %#v", preview)
	}
	if preview.BusName != "org.xnix.Compatibility1" ||
		preview.ObjectPath != "/org/xnix/Compatibility1" ||
		preview.Interface != "org.xnix.Compatibility1" {
		t.Fatalf("unexpected Runtime service activation D-Bus identity: %#v", preview)
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
	if preview.OwnerReadiness.RequestType != "runtime-owner-readiness-preview" ||
		preview.OwnerReadiness.ReadinessType != "runtime-owner-readiness" ||
		!preview.OwnerReadiness.ActivationBindingReady ||
		!preview.OwnerReadiness.ReadOnlyMethodParityReady ||
		!preview.OwnerReadiness.OwnerSmokePlanned ||
		preview.OwnerReadiness.ProductionOwnerRoutesReady ||
		preview.OwnerReadiness.ProductionRecipeTrustReady ||
		preview.OwnerReadiness.LiveDBusOwnerReady ||
		preview.OwnerReadiness.ProductionOwnerEnabled ||
		preview.OwnerReadiness.OwnerTransitionReady ||
		preview.OwnerReadiness.SystemServiceStarted ||
		preview.OwnerReadiness.ProductionBusClaimed ||
		preview.OwnerReadiness.WriteMethodsEnabled ||
		preview.OwnerReadiness.Counts.Passed != 6 ||
		preview.OwnerReadiness.Counts.Pending != 4 {
		t.Fatalf("unexpected owner readiness summary: %#v", preview.OwnerReadiness)
	}
	if preview.OwnerSmokePlan.RequestType != "runtime-owner-smoke-plan-preview" ||
		preview.OwnerSmokePlan.PlanType != "runtime-owner-smoke-plan" ||
		preview.OwnerSmokePlan.SmokeState != "planned" ||
		preview.OwnerSmokePlan.SmokeEnvironment != "restricted-session" ||
		!preview.OwnerSmokePlan.ActivationBindingReady ||
		preview.OwnerSmokePlan.PendingStepCount != 6 ||
		preview.OwnerSmokePlan.Counts.Passed != 1 ||
		preview.OwnerSmokePlan.Counts.Pending != 6 {
		t.Fatalf("unexpected owner smoke summary: %#v", preview.OwnerSmokePlan)
	}

	expectedIDs := []string{
		"activation-binding",
		"read-only-method-parity",
		"owner-smoke-plan",
		"write-method-gate",
		"kde-ownership-boundary",
		"host-safety-boundary",
		"long-running-runtime-owner",
		"restricted-owner-smoke",
		"production-bus-claim",
		"production-recipe-trust",
	}
	expectedStatuses := []string{"pass", "pass", "pass", "pass", "pass", "pass", "pending", "pending", "pending", "pending"}
	if len(preview.PreflightChecks) != len(expectedIDs) || len(preview.CheckIDs) != len(expectedIDs) {
		t.Fatalf("unexpected preflight checks: %#v ids=%#v", preview.PreflightChecks, preview.CheckIDs)
	}
	for index, id := range expectedIDs {
		if preview.PreflightChecks[index].ID != id ||
			preview.CheckIDs[index] != id ||
			preview.PreflightChecks[index].Status != expectedStatuses[index] {
			t.Fatalf("unexpected preflight check at %d: %#v ids=%#v", index, preview.PreflightChecks, preview.CheckIDs)
		}
	}
	if preview.Counts.Total != 10 ||
		preview.Counts.Passed != 6 ||
		preview.Counts.Pending != 4 ||
		preview.Counts.Blocked != 0 {
		t.Fatalf("unexpected preflight counts: %#v", preview.Counts)
	}
	if preview.PreflightDecision != "restricted-owner-smoke-ready" ||
		preview.ProductionActivationReady ||
		!preview.RestrictedSmokeReady ||
		!preview.HumanAuthorizationRequired ||
		!preview.RuntimeOwned ||
		!preview.GoRuntimeBacked ||
		preview.KDEPolicyOwner ||
		preview.KDEMayClaimRuntimeOwnership ||
		preview.SystemServiceStarted ||
		preview.ProductionBusClaimed ||
		preview.WriteMethodsEnabled ||
		preview.BackendLaunchEnabled ||
		preview.NetworkRequired ||
		preview.HostRootModified ||
		preview.PrivilegedContainerRequired ||
		preview.BackendDetailsExposed {
		t.Fatalf("unexpected Runtime service activation safety flags: %#v", preview)
	}
	if len(preview.BlockedReasons) != 4 ||
		preview.BlockedReasons[0] != "Live production Runtime ownership is not proven." ||
		len(preview.BlockedActions) != 6 ||
		preview.BlockedActions[0] != "start production Runtime service from preflight" ||
		len(preview.NextRequirements) != 4 {
		t.Fatalf("unexpected preflight blocked metadata: reasons=%#v actions=%#v next=%#v", preview.BlockedReasons, preview.BlockedActions, preview.NextRequirements)
	}
	if preview.DesktopSafeSummary != "Runtime service activation is ready for restricted owner smoke; production activation, bus claim, and service start remain disabled." {
		t.Fatalf("unexpected desktop-safe summary: %q", preview.DesktopSafeSummary)
	}
	if err := validateNoBackendTerms(preview, "Runtime service activation preflight test"); err != nil {
		t.Fatalf("validateNoBackendTerms returned error: %v", err)
	}
}

func TestRuntimeServiceActivationPreflightPreviewBlocksMissingActivationSources(t *testing.T) {
	root := t.TempDir()
	if err := os.WriteFile(filepath.Join(root, "VERSION"), []byte(currentProjectVersion(t)+"\n"), 0o600); err != nil {
		t.Fatalf("WriteFile VERSION returned error: %v", err)
	}

	preview, err := NewRuntimeServiceActivationPreflightPreview(root)
	if err != nil {
		t.Fatalf("NewRuntimeServiceActivationPreflightPreview returned error: %v", err)
	}

	if preview.PreflightDecision != "production-activation-blocked" ||
		preview.ProductionActivationReady ||
		preview.RestrictedSmokeReady ||
		preview.SystemServiceStarted ||
		preview.ProductionBusClaimed ||
		preview.WriteMethodsEnabled ||
		preview.BackendLaunchEnabled ||
		preview.HostRootModified ||
		preview.BackendDetailsExposed {
		t.Fatalf("missing sources must keep activation closed: %#v", preview)
	}
	if preview.PreflightChecks[0].ID != "activation-binding" ||
		preview.PreflightChecks[0].Status != "blocked" ||
		preview.PreflightChecks[1].ID != "read-only-method-parity" ||
		preview.PreflightChecks[1].Status != "blocked" {
		t.Fatalf("missing sources must block activation and method parity: %#v", preview.PreflightChecks)
	}
	if preview.Counts.Total != 10 ||
		preview.Counts.Passed != 4 ||
		preview.Counts.Pending != 4 ||
		preview.Counts.Blocked != 2 {
		t.Fatalf("unexpected missing-source preflight counts: %#v", preview.Counts)
	}
	if preview.DesktopSafeSummary != "Production Runtime service activation is blocked until activation, method parity, and safety defects are repaired." {
		t.Fatalf("unexpected missing-source summary: %q", preview.DesktopSafeSummary)
	}
}
