package appidentity

import (
	"os"
	"path/filepath"
	"testing"
)

func TestRuntimeOwnerProcessPreviewReportsCurrentWrapperAndGoTarget(t *testing.T) {
	preview, err := NewRuntimeOwnerProcessPreview(projectRootForRuntimeServiceBindingTest(t))
	if err != nil {
		t.Fatalf("NewRuntimeOwnerProcessPreview returned error: %v", err)
	}

	if preview.Version != "0.2.191" ||
		preview.SchemaVersion != "xnix.runtime.owner_process.v1" ||
		preview.RequestType != "runtime-owner-process-preview" ||
		preview.ProcessType != "runtime-owner-process" ||
		preview.Source != "runtime-service-binding-preview+libexec-wrapper+go-owner-target" ||
		preview.RuntimeMethod != "GetRuntimeOwnerProcess" ||
		preview.ReadMethod != "GetRuntimeOwnerProcessPreview" {
		t.Fatalf("unexpected Runtime owner process schema: %#v", preview)
	}
	if preview.BusName != "org.xnix.Compatibility1" ||
		preview.ObjectPath != "/org/xnix/Compatibility1" ||
		preview.Interface != "org.xnix.Compatibility1" {
		t.Fatalf("unexpected Runtime D-Bus identity: %#v", preview)
	}
	if preview.CurrentOwnerEntrypoint != "libexec/xnix/compatd" ||
		preview.CurrentOwnerLanguage != "ruby-wrapper" ||
		preview.TargetOwnerLanguage != "go" {
		t.Fatalf("unexpected owner process language metadata: %#v", preview)
	}
	expectedIDs := []string{"service-activation", "packaged-entrypoint", "go-owner-target", "production-owner-loop", "host-safety-boundary"}
	expectedStatuses := []string{"pass", "pass", "pending", "pending", "pass"}
	if len(preview.Checks) != len(expectedIDs) || len(preview.CheckIDs) != len(expectedIDs) {
		t.Fatalf("unexpected checks: %#v ids=%#v", preview.Checks, preview.CheckIDs)
	}
	for index, id := range expectedIDs {
		if preview.Checks[index].ID != id ||
			preview.CheckIDs[index] != id ||
			preview.Checks[index].Status != expectedStatuses[index] {
			t.Fatalf("unexpected owner process check at %d: %#v ids=%#v", index, preview.Checks, preview.CheckIDs)
		}
	}
	if preview.Counts.Total != 5 ||
		preview.Counts.Passed != 3 ||
		preview.Counts.Pending != 2 ||
		preview.Counts.Blocked != 0 {
		t.Fatalf("unexpected owner process counts: %#v", preview.Counts)
	}
	if !preview.RuntimeOwned ||
		!preview.GoRuntimeBacked ||
		preview.KDEPolicyOwner ||
		preview.KDEMayClaimRuntimeOwnership ||
		!preview.ServiceActivationReady ||
		!preview.PackagedEntrypointReady ||
		preview.GoOwnerProcessReady ||
		preview.ProductionOwnerProcessReady ||
		preview.SystemServiceStarted ||
		preview.ProductionBusClaimed ||
		preview.NetworkRequired ||
		preview.HostRootModified ||
		preview.PrivilegedContainerRequired ||
		preview.BackendDetailsExposed {
		t.Fatalf("unexpected owner process safety flags: %#v", preview)
	}
	if len(preview.BlockedActions) != 6 ||
		preview.BlockedActions[0] != "start production Runtime owner from process preview" ||
		len(preview.NextRequirements) != 4 ||
		preview.NextRequirements[0] != "Implement a Go long-running Runtime owner process." {
		t.Fatalf("unexpected blocked actions or next requirements: actions=%#v next=%#v", preview.BlockedActions, preview.NextRequirements)
	}
	if preview.DesktopSafeSummary != "Runtime owner process activation is aligned, but the production Go owner loop remains pending." {
		t.Fatalf("unexpected owner process summary: %q", preview.DesktopSafeSummary)
	}
	if err := validateNoBackendTerms(preview, "Runtime owner process preview test"); err != nil {
		t.Fatalf("validateNoBackendTerms returned error: %v", err)
	}
}

func TestRuntimeOwnerProcessPreviewBlocksMissingActivationAndEntrypoint(t *testing.T) {
	root := t.TempDir()
	if err := os.WriteFile(filepath.Join(root, "VERSION"), []byte("0.2.191\n"), 0o600); err != nil {
		t.Fatalf("WriteFile VERSION returned error: %v", err)
	}

	preview, err := NewRuntimeOwnerProcessPreview(root)
	if err != nil {
		t.Fatalf("NewRuntimeOwnerProcessPreview returned error: %v", err)
	}

	if preview.ServiceActivationReady ||
		preview.PackagedEntrypointReady ||
		preview.GoOwnerProcessReady ||
		preview.ProductionOwnerProcessReady ||
		preview.SystemServiceStarted ||
		preview.ProductionBusClaimed ||
		preview.HostRootModified ||
		preview.BackendDetailsExposed {
		t.Fatalf("missing activation and entrypoint must keep owner process closed: %#v", preview)
	}
	if preview.Checks[0].ID != "service-activation" ||
		preview.Checks[0].Status != "blocked" ||
		preview.Checks[1].ID != "packaged-entrypoint" ||
		preview.Checks[1].Status != "blocked" {
		t.Fatalf("missing owner process inputs must block checks: %#v", preview.Checks)
	}
	if preview.Counts.Total != 5 ||
		preview.Counts.Passed != 1 ||
		preview.Counts.Pending != 2 ||
		preview.Counts.Blocked != 2 {
		t.Fatalf("unexpected missing owner process counts: %#v", preview.Counts)
	}
	if preview.DesktopSafeSummary != "Runtime owner process activation is blocked by service or entrypoint defects." {
		t.Fatalf("unexpected missing owner process summary: %q", preview.DesktopSafeSummary)
	}
}
