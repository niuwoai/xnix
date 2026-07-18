package appidentity

import (
	"os"
	"path/filepath"
	"testing"
)

func TestRuntimeWriteGatePreviewKeepsWriteMethodsDisabled(t *testing.T) {
	preview, err := NewRuntimeWriteGatePreview(projectRootForRuntimeServiceBindingTest(t), "Launch")
	if err != nil {
		t.Fatalf("NewRuntimeWriteGatePreview returned error: %v", err)
	}

	if preview.Version != currentProjectVersion(t) ||
		preview.SchemaVersion != "xnix.runtime.write_gate.v1" ||
		preview.RequestType != "runtime-write-gate-preview" ||
		preview.GateType != "runtime-write-gate" ||
		preview.Source != "go-runtime-write-gate+runtime-service-activation-preflight-preview+production-dbus-gate-review-preview+production-dbus-human-authorization-preflight-preview" ||
		preview.RuntimeMethod != "GetRuntimeWriteGate" ||
		preview.ReadMethod != "GetRuntimeWriteGatePreview" ||
		preview.MethodName != "Launch" {
		t.Fatalf("unexpected Runtime write gate schema: %#v", preview)
	}
	if !preview.RuntimeOwned ||
		!preview.GoRuntimeBacked ||
		preview.KDEPolicyOwner ||
		preview.GateDecision != "blocked-until-production-backend" ||
		preview.ProductionGateDecision != "production-gates-consumed-write-gate-disabled" ||
		preview.WriteMethodEnabled ||
		preview.DispatchEnabled ||
		preview.RequestObjectCreated ||
		preview.ExecutionStarted ||
		preview.NetworkRequired ||
		preview.HostRootModified ||
		preview.PrivilegedContainerRequired ||
		preview.BackendDetailsExposed {
		t.Fatalf("unexpected Runtime write gate safety flags: %#v", preview)
	}
	if preview.ProductionGate.RequestType != "runtime-service-activation-preflight-preview" ||
		preview.ProductionGate.PreflightType != "production-runtime-service-activation-preflight" ||
		preview.ProductionGate.PreflightDecision != "restricted-owner-smoke-ready" ||
		!preview.ProductionGate.ServiceActivationPreflightReady ||
		!preview.ProductionGate.ProductionDBusGateReady ||
		!preview.ProductionGate.HumanAuthorizationPreflightReady ||
		!preview.ProductionGate.HumanAuthorizationRequired ||
		preview.ProductionGate.HumanAuthorizationGranted ||
		preview.ProductionGate.AuthorizationReceiptAccepted ||
		preview.ProductionGate.ProductionActivationReady ||
		!preview.ProductionGate.RestrictedSmokeReady ||
		preview.ProductionGate.SystemServiceStarted ||
		preview.ProductionGate.ProductionBusClaimed ||
		preview.ProductionGate.WriteMethodsEnabled ||
		preview.ProductionGate.BackendLaunchEnabled ||
		preview.ProductionGate.NetworkRequired ||
		preview.ProductionGate.HostRootModified ||
		preview.ProductionGate.PrivilegedContainerRequired ||
		preview.ProductionGate.BackendDetailsExposed {
		t.Fatalf("unexpected production gate summary: %#v", preview.ProductionGate)
	}
	if preview.DenialErrorName != runtimeWriteGateErrorName ||
		!sameStrings(preview.SupportedWriteMethods, runtimeMethodParityWriteMethods) ||
		!sameStrings(preview.RequiredGateIDs, runtimeWriteGateRequiredGateIDs) ||
		len(preview.RequiredGates) != len(runtimeWriteGateRequiredGateIDs) {
		t.Fatalf("unexpected Runtime write gate requirements: %#v", preview)
	}
	expectedStatuses := []string{"pass", "pass", "pending", "pending", "pending", "pending", "pending", "pending", "pending"}
	for index, status := range expectedStatuses {
		if preview.RequiredGates[index].ID != runtimeWriteGateRequiredGateIDs[index] ||
			preview.RequiredGates[index].Status != status {
			t.Fatalf("unexpected Runtime write gate check at %d: %#v", index, preview.RequiredGates)
		}
	}
	if preview.Counts.Total != 9 ||
		preview.Counts.Passed != 2 ||
		preview.Counts.Pending != 7 ||
		preview.Counts.Blocked != 0 {
		t.Fatalf("unexpected Runtime write gate counts: %#v", preview.Counts)
	}
	if preview.DesktopSafeSummary != "Launch is visible to desktop integrations but remains blocked until production Runtime gates pass." {
		t.Fatalf("unexpected Runtime write gate summary: %q", preview.DesktopSafeSummary)
	}
	if err := validateNoBackendTerms(preview, "Runtime write gate preview test"); err != nil {
		t.Fatalf("validateNoBackendTerms returned error: %v", err)
	}
}

func TestRuntimeWriteGatePreviewBlocksMissingProductionGateSources(t *testing.T) {
	root := t.TempDir()
	if err := os.WriteFile(filepath.Join(root, "VERSION"), []byte(currentProjectVersion(t)+"\n"), 0o600); err != nil {
		t.Fatalf("WriteFile VERSION returned error: %v", err)
	}

	preview, err := NewRuntimeWriteGatePreview(root, "Launch")
	if err != nil {
		t.Fatalf("NewRuntimeWriteGatePreview returned error: %v", err)
	}

	if preview.ProductionGateDecision != "production-gate-consumption-blocked" ||
		preview.ProductionGate.ServiceActivationPreflightReady ||
		preview.ProductionGate.ProductionDBusGateReady ||
		preview.ProductionGate.HumanAuthorizationPreflightReady ||
		preview.ProductionGate.AuthorizationReceiptAccepted ||
		preview.WriteMethodEnabled ||
		preview.DispatchEnabled ||
		preview.RequestObjectCreated ||
		preview.ExecutionStarted ||
		preview.HostRootModified ||
		preview.BackendDetailsExposed {
		t.Fatalf("missing sources must keep write gate closed: %#v", preview)
	}
	if preview.RequiredGates[0].ID != "production-dbus-gate-review" ||
		preview.RequiredGates[0].Status != "blocked" ||
		preview.RequiredGates[1].ID != "production-service-activation-preflight" ||
		preview.RequiredGates[1].Status != "blocked" ||
		preview.RequiredGates[2].ID != "human-authorization-receipt" ||
		preview.RequiredGates[2].Status != "pending" {
		t.Fatalf("missing sources must block production gate consumption: %#v", preview.RequiredGates)
	}
	if preview.Counts.Total != 9 ||
		preview.Counts.Passed != 0 ||
		preview.Counts.Pending != 7 ||
		preview.Counts.Blocked != 2 {
		t.Fatalf("unexpected missing-source Runtime write gate counts: %#v", preview.Counts)
	}
}

func TestRuntimeWriteGatePreviewRejectsUnknownMethod(t *testing.T) {
	root := t.TempDir()
	if err := os.WriteFile(filepath.Join(root, "VERSION"), []byte(currentProjectVersion(t)+"\n"), 0o600); err != nil {
		t.Fatalf("WriteFile VERSION returned error: %v", err)
	}

	if _, err := NewRuntimeWriteGatePreview(root, "DeleteEverything"); err == nil {
		t.Fatal("NewRuntimeWriteGatePreview accepted an unknown write method")
	}
}
