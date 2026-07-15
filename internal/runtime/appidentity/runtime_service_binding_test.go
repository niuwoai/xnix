package appidentity

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestRuntimeServiceBindingPreviewReportsActivationReadiness(t *testing.T) {
	preview, err := NewRuntimeServiceBindingPreview(projectRootForRuntimeServiceBindingTest(t))
	if err != nil {
		t.Fatalf("NewRuntimeServiceBindingPreview returned error: %v", err)
	}

	if preview.Version != "0.2.191" ||
		preview.SchemaVersion != "xnix.runtime.service_binding.v1" ||
		preview.RequestType != "runtime-service-binding-preview" ||
		preview.BindingType != "runtime-service-binding" ||
		preview.Source != "activation-files+dbus-contract+runtime-owner-gate" ||
		preview.RuntimeMethod != "GetRuntimeServiceBinding" ||
		preview.ReadMethod != "GetRuntimeServiceBindingPreview" ||
		preview.ProductionStatus != "pending-live-owner" {
		t.Fatalf("unexpected Runtime service binding schema: %#v", preview)
	}
	if preview.BusName != "org.xnix.Compatibility1" ||
		preview.ObjectPath != "/org/xnix/Compatibility1" ||
		preview.Interface != "org.xnix.Compatibility1" {
		t.Fatalf("unexpected Runtime D-Bus identity: %#v", preview)
	}
	if preview.Activation.DBusServiceFile != "runtime/dbus/org.xnix.Compatibility1.service" ||
		preview.Activation.SystemdUnit != "runtime/systemd/xnix-compatd.service" ||
		preview.Activation.LibexecWrapper != "libexec/xnix/compatd" ||
		preview.Activation.DBusContract != "runtime/dbus/org.xnix.Compatibility1.xml" ||
		preview.Activation.PackagedWrapper != "/usr/libexec/xnix/compatd" {
		t.Fatalf("unexpected activation files: %#v", preview.Activation)
	}
	expectedIDs := []string{
		"dbus-service-activation",
		"systemd-service-hardening",
		"libexec-wrapper",
		"dbus-contract",
		"live-dbus-owner",
	}
	if len(preview.Checks) != len(expectedIDs) || len(preview.CheckIDs) != len(expectedIDs) {
		t.Fatalf("unexpected check count: %#v", preview.Checks)
	}
	for index, id := range expectedIDs {
		if preview.Checks[index].ID != id || preview.CheckIDs[index] != id {
			t.Fatalf("unexpected check at %d: %#v ids=%#v", index, preview.Checks, preview.CheckIDs)
		}
	}
	for index, check := range preview.Checks {
		expectedStatus := "pass"
		if index == len(preview.Checks)-1 {
			expectedStatus = "pending"
		}
		if check.Status != expectedStatus {
			t.Fatalf("unexpected check status at %d: %#v", index, check)
		}
	}
	if preview.Counts.Total != 5 ||
		preview.Counts.Passed != 4 ||
		preview.Counts.Pending != 1 ||
		preview.Counts.Blocked != 0 {
		t.Fatalf("unexpected check counts: %#v", preview.Counts)
	}
	if !preview.RuntimeOwned ||
		!preview.GoRuntimeBacked ||
		preview.KDEPolicyOwner ||
		!preview.ActivationBindingReady ||
		preview.LiveDBusOwnerReady ||
		!preview.SmokeAdapterAvailable ||
		preview.ServiceStarted ||
		preview.ProductionBusClaimed ||
		preview.KDEMayClaimRuntimeOwnership ||
		preview.NetworkRequired ||
		preview.HostRootModified ||
		preview.PrivilegedContainerRequired ||
		preview.BackendDetailsExposed {
		t.Fatalf("unexpected Runtime service binding safety flags: %#v", preview)
	}
	if len(preview.BlockedActions) != 4 ||
		preview.BlockedActions[0] != "start Runtime service from preview" ||
		preview.BlockedActions[1] != "claim production D-Bus owner from preview" ||
		preview.BlockedActions[2] != "let KDE claim Runtime ownership" ||
		preview.BlockedActions[3] != "mutate host root during service binding planning" {
		t.Fatalf("unexpected blocked actions: %#v", preview.BlockedActions)
	}
	if preview.DesktopSafeSummary != "Runtime service activation files are aligned; live D-Bus ownership remains pending." {
		t.Fatalf("unexpected desktop-safe summary: %q", preview.DesktopSafeSummary)
	}
	if err := validateNoBackendTerms(preview, "Runtime service binding preview test"); err != nil {
		t.Fatalf("validateNoBackendTerms returned error: %v", err)
	}
}

func TestRuntimeServiceBindingPreviewBlocksMissingActivationFiles(t *testing.T) {
	root := t.TempDir()
	if err := os.WriteFile(filepath.Join(root, "VERSION"), []byte("0.2.191\n"), 0o600); err != nil {
		t.Fatalf("WriteFile VERSION returned error: %v", err)
	}

	preview, err := NewRuntimeServiceBindingPreview(root)
	if err != nil {
		t.Fatalf("NewRuntimeServiceBindingPreview returned error: %v", err)
	}

	if preview.ActivationBindingReady ||
		preview.LiveDBusOwnerReady ||
		preview.SmokeAdapterAvailable ||
		preview.ServiceStarted ||
		preview.ProductionBusClaimed ||
		preview.KDEMayClaimRuntimeOwnership ||
		preview.HostRootModified ||
		preview.BackendDetailsExposed {
		t.Fatalf("missing activation files must keep Runtime service ownership closed: %#v", preview)
	}
	if preview.Counts.Total != 5 ||
		preview.Counts.Passed != 0 ||
		preview.Counts.Pending != 1 ||
		preview.Counts.Blocked != 4 {
		t.Fatalf("unexpected missing-file counts: %#v", preview.Counts)
	}
	for index, check := range preview.Checks {
		expectedStatus := "blocked"
		if check.ID == "live-dbus-owner" {
			expectedStatus = "pending"
		}
		if check.Status != expectedStatus {
			t.Fatalf("unexpected missing-file check at %d: %#v", index, check)
		}
	}
	if preview.DesktopSafeSummary != "Runtime service activation files need repair before production D-Bus ownership can be enabled." {
		t.Fatalf("unexpected missing-file summary: %q", preview.DesktopSafeSummary)
	}
}

func projectRootForRuntimeServiceBindingTest(t *testing.T) string {
	t.Helper()
	workingDirectory, err := os.Getwd()
	if err != nil {
		t.Fatalf("Getwd returned error: %v", err)
	}
	for {
		if _, err := os.Stat(filepath.Join(workingDirectory, "VERSION")); err == nil {
			return workingDirectory
		}
		parent := filepath.Dir(workingDirectory)
		if parent == workingDirectory || strings.TrimSpace(parent) == "" {
			t.Fatalf("could not find project root from %s", workingDirectory)
		}
		workingDirectory = parent
	}
}
