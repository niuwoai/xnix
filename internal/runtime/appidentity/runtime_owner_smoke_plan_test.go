package appidentity

import (
	"os"
	"path/filepath"
	"testing"
)

func TestRuntimeOwnerSmokePlanPreviewKeepsProductionSmokePlanned(t *testing.T) {
	preview, err := NewRuntimeOwnerSmokePlanPreview(projectRootForRuntimeServiceBindingTest(t))
	if err != nil {
		t.Fatalf("NewRuntimeOwnerSmokePlanPreview returned error: %v", err)
	}

	if preview.Version != "0.2.209" ||
		preview.SchemaVersion != "xnix.runtime.owner_smoke_plan.v1" ||
		preview.RequestType != "runtime-owner-smoke-plan-preview" ||
		preview.PlanType != "runtime-owner-smoke-plan" ||
		preview.Source != "runtime-live-owner-gate-preview+runtime-service-binding-preview" ||
		preview.RuntimeMethod != "GetRuntimeOwnerSmokePlan" ||
		preview.ReadMethod != "GetRuntimeOwnerSmokePlanPreview" {
		t.Fatalf("unexpected Runtime owner smoke plan schema: %#v", preview)
	}
	if preview.BusName != "org.xnix.Compatibility1" ||
		preview.ObjectPath != "/org/xnix/Compatibility1" ||
		preview.Interface != "org.xnix.Compatibility1" {
		t.Fatalf("unexpected Runtime D-Bus identity: %#v", preview)
	}
	if preview.LiveOwnerGate.RequestType != "runtime-live-owner-gate-preview" ||
		preview.LiveOwnerGate.GateType != "runtime-live-owner-gate" ||
		!preview.LiveOwnerGate.ActivationBindingReady ||
		preview.LiveOwnerGate.LiveDBusOwnerReady ||
		preview.LiveOwnerGate.ProductionOwnerEnabled ||
		preview.LiveOwnerGate.OwnerTransitionReady ||
		preview.LiveOwnerGate.Counts.Passed != 1 ||
		preview.LiveOwnerGate.Counts.Pending != 4 {
		t.Fatalf("unexpected live owner gate summary: %#v", preview.LiveOwnerGate)
	}
	expectedIDs := []string{
		"validate-activation-files",
		"start-packaged-runtime-owner",
		"assert-stable-bus-name",
		"check-read-only-method-parity",
		"reject-write-methods",
		"verify-non-production-smoke-adapter-boundary",
		"report-kde-safe-summary",
	}
	expectedStatuses := []string{"pass", "pending", "pending", "pending", "pending", "pending", "pending"}
	if len(preview.Steps) != len(expectedIDs) || len(preview.StepIDs) != len(expectedIDs) {
		t.Fatalf("unexpected owner smoke steps: %#v ids=%#v", preview.Steps, preview.StepIDs)
	}
	for index, id := range expectedIDs {
		if preview.Steps[index].ID != id ||
			preview.StepIDs[index] != id ||
			preview.Steps[index].Status != expectedStatuses[index] {
			t.Fatalf("unexpected step at %d: %#v ids=%#v", index, preview.Steps, preview.StepIDs)
		}
	}
	if preview.Counts.Total != 7 ||
		preview.Counts.Passed != 1 ||
		preview.Counts.Pending != 6 ||
		preview.Counts.Blocked != 0 ||
		preview.PendingStepCount != 6 {
		t.Fatalf("unexpected owner smoke counts: %#v pending=%d", preview.Counts, preview.PendingStepCount)
	}
	if !preview.RuntimeOwned ||
		!preview.GoRuntimeBacked ||
		preview.KDEPolicyOwner ||
		!preview.ActivationBindingReady ||
		preview.LiveDBusOwnerReady ||
		preview.ProductionOwnerEnabled ||
		preview.OwnerTransitionReady ||
		preview.SmokeState != "planned" ||
		preview.SmokeEnvironment != "restricted-session" ||
		preview.NetworkRequired ||
		preview.HostRootModified ||
		preview.PrivilegedContainerRequired ||
		preview.SystemServiceStarted ||
		preview.ProductionBusClaimed ||
		preview.BackendDetailsExposed {
		t.Fatalf("unexpected owner smoke safety flags: %#v", preview)
	}
	if len(preview.BlockedActions) != 5 ||
		preview.BlockedActions[0] != "Do not start a host system service from the smoke plan." ||
		preview.BlockedActions[1] != "Do not claim the production Runtime bus name from the smoke adapter." {
		t.Fatalf("unexpected blocked actions: %#v", preview.BlockedActions)
	}
	if preview.DesktopSafeSummary != "Runtime owner smoke is planned; production bus ownership remains disabled until all gates pass." {
		t.Fatalf("unexpected desktop-safe summary: %q", preview.DesktopSafeSummary)
	}
	if err := validateNoBackendTerms(preview, "Runtime owner smoke plan preview test"); err != nil {
		t.Fatalf("validateNoBackendTerms returned error: %v", err)
	}
}

func TestRuntimeOwnerSmokePlanPreviewBlocksWhenActivationIsMissing(t *testing.T) {
	root := t.TempDir()
	if err := os.WriteFile(filepath.Join(root, "VERSION"), []byte("0.2.209\n"), 0o600); err != nil {
		t.Fatalf("WriteFile VERSION returned error: %v", err)
	}

	preview, err := NewRuntimeOwnerSmokePlanPreview(root)
	if err != nil {
		t.Fatalf("NewRuntimeOwnerSmokePlanPreview returned error: %v", err)
	}

	if preview.ActivationBindingReady ||
		preview.LiveDBusOwnerReady ||
		preview.ProductionOwnerEnabled ||
		preview.OwnerTransitionReady ||
		preview.SystemServiceStarted ||
		preview.ProductionBusClaimed ||
		preview.HostRootModified ||
		preview.BackendDetailsExposed {
		t.Fatalf("missing activation must keep owner smoke closed: %#v", preview)
	}
	if preview.LiveOwnerGate.Counts.Blocked != 1 ||
		preview.LiveOwnerGate.Counts.Pending != 4 {
		t.Fatalf("unexpected missing activation live owner gate counts: %#v", preview.LiveOwnerGate.Counts)
	}
	if preview.Steps[0].ID != "validate-activation-files" || preview.Steps[0].Status != "blocked" {
		t.Fatalf("missing activation must block activation validation step: %#v", preview.Steps)
	}
	if preview.Counts.Total != 7 ||
		preview.Counts.Passed != 0 ||
		preview.Counts.Pending != 6 ||
		preview.Counts.Blocked != 1 ||
		preview.PendingStepCount != 6 {
		t.Fatalf("unexpected missing activation owner smoke counts: %#v pending=%d", preview.Counts, preview.PendingStepCount)
	}
	if preview.DesktopSafeSummary != "Runtime activation files must be repaired before owner smoke can run." {
		t.Fatalf("unexpected missing activation summary: %q", preview.DesktopSafeSummary)
	}
}
