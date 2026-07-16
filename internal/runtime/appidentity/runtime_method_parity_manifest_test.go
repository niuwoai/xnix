package appidentity

import (
	"os"
	"path/filepath"
	"testing"
)

func TestRuntimeMethodParityManifestPreviewChecksProjectSources(t *testing.T) {
	preview, err := NewRuntimeMethodParityManifestPreview(projectRootForRuntimeServiceBindingTest(t))
	if err != nil {
		t.Fatalf("NewRuntimeMethodParityManifestPreview returned error: %v", err)
	}

	if preview.Version != "0.2.246" ||
		preview.SchemaVersion != "xnix.runtime.method_parity_manifest.v1" ||
		preview.RequestType != "runtime-method-parity-manifest-preview" ||
		preview.ManifestType != "runtime-method-parity-manifest" ||
		preview.Source != "dbus-contract+runtime-dispatch+dbus-client+smoke-adapter+session-smoke" ||
		preview.RuntimeMethod != "GetRuntimeMethodParityManifest" ||
		preview.ReadMethod != "GetRuntimeMethodParityManifestPreview" {
		t.Fatalf("unexpected Runtime method parity schema: %#v", preview)
	}
	if preview.BusName != "org.xnix.Compatibility1" ||
		preview.ObjectPath != "/org/xnix/Compatibility1" ||
		preview.Interface != "org.xnix.Compatibility1" {
		t.Fatalf("unexpected Runtime D-Bus identity: %#v", preview)
	}
	if preview.MethodCount != 61 || len(preview.ReadOnlyMethods) != 61 {
		t.Fatalf("unexpected read-only method count: %d len=%d", preview.MethodCount, len(preview.ReadOnlyMethods))
	}
	for _, method := range []string{
		"ListApplications",
		"GetDesktopActivationManifest",
		"GetRuntimeServiceBinding",
		"GetRuntimeLiveOwnerGate",
		"GetRuntimeOwnerSmokePlan",
		"GetRuntimeMethodParityManifest",
		"GetRuntimeWriteGate",
		"GetKDECenterPageSectionDetail",
	} {
		if !containsString(preview.ReadOnlyMethods, method) {
			t.Fatalf("read-only methods missing %s: %#v", method, preview.ReadOnlyMethods)
		}
	}
	expectedCheckIDs := []string{"dbus-contract", "runtime-dispatch", "dbus-client", "smoke-adapter", "session-smoke"}
	if len(preview.ParityChecks) != len(expectedCheckIDs) || len(preview.CheckIDs) != len(expectedCheckIDs) {
		t.Fatalf("unexpected parity check count: %#v ids=%#v", preview.ParityChecks, preview.CheckIDs)
	}
	for index, id := range expectedCheckIDs {
		check := preview.ParityChecks[index]
		if check.ID != id ||
			preview.CheckIDs[index] != id ||
			check.Status != "pass" ||
			check.MethodCount != 61 ||
			len(check.MissingMethods) != 0 {
			t.Fatalf("unexpected parity check at %d: %#v ids=%#v", index, preview.ParityChecks, preview.CheckIDs)
		}
	}
	if preview.Counts.Total != 5 ||
		preview.Counts.Passed != 5 ||
		preview.Counts.Blocked != 0 ||
		preview.Counts.Pending != 0 ||
		!preview.ReadOnlyMethodParityReady {
		t.Fatalf("unexpected parity counts: %#v ready=%v", preview.Counts, preview.ReadOnlyMethodParityReady)
	}
	if len(preview.WriteMethods) != 4 ||
		preview.WriteMethods[0] != "InstallRecipe" ||
		preview.WriteMethods[1] != "Launch" ||
		preview.WriteMethods[2] != "CreateSnapshot" ||
		preview.WriteMethods[3] != "RestoreSnapshot" ||
		preview.WriteMethodsSupported ||
		preview.WriteMethodDispatchEnabled {
		t.Fatalf("unexpected write-method boundary: %#v supported=%v dispatch=%v", preview.WriteMethods, preview.WriteMethodsSupported, preview.WriteMethodDispatchEnabled)
	}
	if !preview.RuntimeOwned ||
		!preview.GoRuntimeBacked ||
		preview.KDEPolicyOwner ||
		preview.NetworkRequired ||
		preview.HostRootModified ||
		preview.PrivilegedContainerRequired ||
		preview.BackendDetailsExposed {
		t.Fatalf("unexpected method parity safety flags: %#v", preview)
	}
	if preview.DesktopSafeSummary != "Runtime read-only D-Bus method parity is ready for owner smoke." {
		t.Fatalf("unexpected desktop-safe summary: %q", preview.DesktopSafeSummary)
	}
	if err := validateNoBackendTerms(preview, "Runtime method parity manifest preview test"); err != nil {
		t.Fatalf("validateNoBackendTerms returned error: %v", err)
	}
}

func TestRuntimeMethodParityManifestPreviewBlocksMissingSources(t *testing.T) {
	root := t.TempDir()
	if err := os.WriteFile(filepath.Join(root, "VERSION"), []byte("0.2.246\n"), 0o600); err != nil {
		t.Fatalf("WriteFile VERSION returned error: %v", err)
	}

	preview, err := NewRuntimeMethodParityManifestPreview(root)
	if err != nil {
		t.Fatalf("NewRuntimeMethodParityManifestPreview returned error: %v", err)
	}

	if preview.ReadOnlyMethodParityReady {
		t.Fatalf("missing sources must not be parity-ready: %#v", preview)
	}
	if preview.Counts.Total != 5 ||
		preview.Counts.Passed != 0 ||
		preview.Counts.Blocked != 5 ||
		preview.Counts.Pending != 0 {
		t.Fatalf("unexpected missing-source counts: %#v", preview.Counts)
	}
	for index, check := range preview.ParityChecks {
		if check.Status != "blocked" ||
			check.MethodCount != 0 ||
			len(check.MissingMethods) != 61 {
			t.Fatalf("unexpected missing-source check at %d: %#v", index, check)
		}
	}
	if preview.WriteMethodsSupported ||
		preview.WriteMethodDispatchEnabled ||
		preview.HostRootModified ||
		preview.BackendDetailsExposed {
		t.Fatalf("missing sources must keep method parity read-only and safe: %#v", preview)
	}
	if preview.DesktopSafeSummary != "Runtime read-only D-Bus method parity must be repaired before production owner smoke." {
		t.Fatalf("unexpected missing-source summary: %q", preview.DesktopSafeSummary)
	}
}
