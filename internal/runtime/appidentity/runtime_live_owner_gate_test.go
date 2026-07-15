package appidentity

import (
	"os"
	"path/filepath"
	"testing"
)

func TestRuntimeLiveOwnerGatePreviewKeepsProductionOwnershipGated(t *testing.T) {
	preview, err := NewRuntimeLiveOwnerGatePreview(projectRootForRuntimeServiceBindingTest(t))
	if err != nil {
		t.Fatalf("NewRuntimeLiveOwnerGatePreview returned error: %v", err)
	}

	if preview.Version != "0.2.228" ||
		preview.SchemaVersion != "xnix.runtime.live_owner_gate.v1" ||
		preview.RequestType != "runtime-live-owner-gate-preview" ||
		preview.GateType != "runtime-live-owner-gate" ||
		preview.Source != "runtime-service-binding-preview+owner-transition-gates" ||
		preview.RuntimeMethod != "GetRuntimeLiveOwnerGate" ||
		preview.ReadMethod != "GetRuntimeLiveOwnerGatePreview" {
		t.Fatalf("unexpected Runtime live owner gate schema: %#v", preview)
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
	expectedIDs := []string{
		"activation-binding",
		"long-running-runtime-owner",
		"bus-name-acquisition",
		"read-only-method-parity",
		"production-recipe-trust",
	}
	expectedStatuses := []string{"pass", "pending", "pending", "pending", "pending"}
	if len(preview.RequiredGates) != len(expectedIDs) || len(preview.GateIDs) != len(expectedIDs) {
		t.Fatalf("unexpected required gates: %#v ids=%#v", preview.RequiredGates, preview.GateIDs)
	}
	for index, id := range expectedIDs {
		if preview.RequiredGates[index].ID != id ||
			preview.GateIDs[index] != id ||
			preview.RequiredGates[index].Status != expectedStatuses[index] {
			t.Fatalf("unexpected gate at %d: %#v ids=%#v", index, preview.RequiredGates, preview.GateIDs)
		}
	}
	if preview.Counts.Total != 5 ||
		preview.Counts.Passed != 1 ||
		preview.Counts.Pending != 4 ||
		preview.Counts.Blocked != 0 {
		t.Fatalf("unexpected live owner gate counts: %#v", preview.Counts)
	}
	if !preview.RuntimeOwned ||
		!preview.GoRuntimeBacked ||
		preview.KDEPolicyOwner ||
		!preview.ActivationBindingReady ||
		preview.LiveDBusOwnerReady ||
		preview.ProductionOwnerEnabled ||
		preview.OwnerTransitionReady ||
		!preview.SmokeAdapterAvailable ||
		preview.SmokeAdapterIsProduction ||
		preview.KDEMayClaimRuntimeOwnership ||
		preview.NetworkRequired ||
		preview.HostRootModified ||
		preview.PrivilegedContainerRequired ||
		preview.BackendDetailsExposed {
		t.Fatalf("unexpected live owner gate safety flags: %#v", preview)
	}
	if len(preview.BlockedReasons) != 5 ||
		len(preview.BlockedActions) != 6 ||
		preview.BlockedActions[0] != "start production Runtime owner from preview" ||
		preview.BlockedActions[3] != "let KDE claim Runtime ownership" {
		t.Fatalf("unexpected live owner gate blockers: reasons=%#v actions=%#v", preview.BlockedReasons, preview.BlockedActions)
	}
	if preview.DesktopSafeSummary != "Runtime activation files are aligned, but production D-Bus ownership remains gated." {
		t.Fatalf("unexpected desktop-safe summary: %q", preview.DesktopSafeSummary)
	}
	if err := validateNoBackendTerms(preview, "Runtime live owner gate preview test"); err != nil {
		t.Fatalf("validateNoBackendTerms returned error: %v", err)
	}
}

func TestRuntimeLiveOwnerGatePreviewBlocksWhenActivationIsMissing(t *testing.T) {
	root := t.TempDir()
	if err := os.WriteFile(filepath.Join(root, "VERSION"), []byte("0.2.228\n"), 0o600); err != nil {
		t.Fatalf("WriteFile VERSION returned error: %v", err)
	}

	preview, err := NewRuntimeLiveOwnerGatePreview(root)
	if err != nil {
		t.Fatalf("NewRuntimeLiveOwnerGatePreview returned error: %v", err)
	}

	if preview.ActivationBindingReady ||
		preview.LiveDBusOwnerReady ||
		preview.ProductionOwnerEnabled ||
		preview.OwnerTransitionReady ||
		preview.SmokeAdapterAvailable ||
		preview.SmokeAdapterIsProduction ||
		preview.KDEMayClaimRuntimeOwnership ||
		preview.HostRootModified ||
		preview.BackendDetailsExposed {
		t.Fatalf("missing activation must keep live owner transition closed: %#v", preview)
	}
	if preview.ServiceBinding.Counts.Blocked != 4 ||
		preview.ServiceBinding.Counts.Pending != 1 {
		t.Fatalf("unexpected missing activation service binding counts: %#v", preview.ServiceBinding.Counts)
	}
	if preview.RequiredGates[0].ID != "activation-binding" || preview.RequiredGates[0].Status != "blocked" {
		t.Fatalf("missing activation must block activation-binding gate: %#v", preview.RequiredGates)
	}
	if preview.Counts.Total != 5 ||
		preview.Counts.Passed != 0 ||
		preview.Counts.Pending != 4 ||
		preview.Counts.Blocked != 1 {
		t.Fatalf("unexpected missing activation live owner gate counts: %#v", preview.Counts)
	}
	if preview.DesktopSafeSummary != "Runtime activation files are blocked; production D-Bus ownership cannot be enabled." {
		t.Fatalf("unexpected missing activation summary: %q", preview.DesktopSafeSummary)
	}
}
