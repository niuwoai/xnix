package appidentity

import (
	"os"
	"path/filepath"
	"testing"
)

func TestRuntimeWriteGatePreviewKeepsWriteMethodsDisabled(t *testing.T) {
	root := t.TempDir()
	if err := os.WriteFile(filepath.Join(root, "VERSION"), []byte("0.2.303\n"), 0o600); err != nil {
		t.Fatalf("WriteFile VERSION returned error: %v", err)
	}

	preview, err := NewRuntimeWriteGatePreview(root, "Launch")
	if err != nil {
		t.Fatalf("NewRuntimeWriteGatePreview returned error: %v", err)
	}

	if preview.Version != "0.2.303" ||
		preview.SchemaVersion != "xnix.runtime.write_gate.v1" ||
		preview.RequestType != "runtime-write-gate-preview" ||
		preview.GateType != "runtime-write-gate" ||
		preview.Source != "go-runtime-write-gate" ||
		preview.RuntimeMethod != "GetRuntimeWriteGate" ||
		preview.ReadMethod != "GetRuntimeWriteGatePreview" ||
		preview.MethodName != "Launch" {
		t.Fatalf("unexpected Runtime write gate schema: %#v", preview)
	}
	if !preview.RuntimeOwned ||
		!preview.GoRuntimeBacked ||
		preview.KDEPolicyOwner ||
		preview.GateDecision != "blocked-until-production-backend" ||
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
	if preview.DenialErrorName != runtimeWriteGateErrorName ||
		!sameStrings(preview.SupportedWriteMethods, runtimeMethodParityWriteMethods) ||
		!sameStrings(preview.RequiredGateIDs, runtimeWriteGateRequiredGateIDs) ||
		len(preview.RequiredGates) != len(runtimeWriteGateRequiredGateIDs) {
		t.Fatalf("unexpected Runtime write gate requirements: %#v", preview)
	}
	if preview.DesktopSafeSummary != "Launch is visible to desktop integrations but remains blocked until production Runtime gates pass." {
		t.Fatalf("unexpected Runtime write gate summary: %q", preview.DesktopSafeSummary)
	}
	if err := validateNoBackendTerms(preview, "Runtime write gate preview test"); err != nil {
		t.Fatalf("validateNoBackendTerms returned error: %v", err)
	}
}

func TestRuntimeWriteGatePreviewRejectsUnknownMethod(t *testing.T) {
	root := t.TempDir()
	if err := os.WriteFile(filepath.Join(root, "VERSION"), []byte("0.2.303\n"), 0o600); err != nil {
		t.Fatalf("WriteFile VERSION returned error: %v", err)
	}

	if _, err := NewRuntimeWriteGatePreview(root, "DeleteEverything"); err == nil {
		t.Fatal("NewRuntimeWriteGatePreview accepted an unknown write method")
	}
}
