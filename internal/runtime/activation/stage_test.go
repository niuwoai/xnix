package activation

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"xnix.local/xnix/internal/runtime/appidentity"
)

func TestStageWritesDesktopActivationArtifactsInsideStagingRoot(t *testing.T) {
	root := t.TempDir()
	plan := testPlan(t)

	result, err := Stage(StageRequest{
		Root: root,
		Mode: "development",
		Plan: plan,
	})
	if err != nil {
		t.Fatalf("Stage returned error: %v", err)
	}

	if result.SchemaVersion != "xnix.runtime.desktop_activation_stage.v1" ||
		result.RequestType != "desktop-activation-stage" ||
		result.StageType != "kde-desktop-activation-test-root-stage" ||
		result.Desktop != "KDE Plasma" ||
		result.ApplicationID != "org.example.ledger" ||
		result.DisplayName != "Example Ledger" ||
		result.DesktopFile != "xnix-org.example.ledger.desktop" ||
		result.InstallMode != "development" ||
		result.PreflightDecision != "development-staging-ready" {
		t.Fatalf("unexpected stage identity: %#v", result)
	}
	if got, want := result.WrittenFileIDs, []string{"desktop-activation-receipt", "desktop-entry", "desktop-integration-manifest", "dolphin-service-menu", "managed-launcher-artifact", "mimeapps-list"}; !sameStrings(got, want) {
		t.Fatalf("WrittenFileIDs = %#v, want %#v", got, want)
	}
	if result.WrittenFileCount != 6 || len(result.WrittenFiles) != 6 {
		t.Fatalf("unexpected written file count: %#v", result)
	}
	if !result.RuntimeOwned || !result.GoRuntimeBacked || result.KDEPolicyOwner ||
		!result.StagingRootRequired || result.StagingRootPathExposed ||
		result.HostRootAllowed || !result.FileWritesPerformed ||
		!result.DesktopFilesWritten || !result.MIMEAppsWritten ||
		!result.ManifestWritten || !result.ReceiptWritten ||
		!result.RollbackReceiptWritten {
		t.Fatalf("unexpected staging flags: %#v", result)
	}
	if result.SettingsPersisted || result.NotificationsSent ||
		result.TaskManagerEntryActive || result.KWinRuleApplied ||
		result.LiveTrayBridgeEnabled || result.LaunchEnabled ||
		result.BackendLaunchEnabled || result.ExecutionStarted ||
		result.HostRootModified || result.NetworkRequired ||
		result.PrivilegedContainerRequired || result.BackendDetailsExposed {
		t.Fatalf("desktop activation staging opened an unsafe gate: %#v", result)
	}

	expectedFiles := map[string]string{
		"desktop-entry":                "usr/share/applications/xnix-org.example.ledger.desktop",
		"dolphin-service-menu":         "usr/share/kio/servicemenus/xnix-open-with-compatibility.desktop",
		"mimeapps-list":                "usr/share/applications/mimeapps.list",
		"desktop-integration-manifest": "usr/share/xnix/compatibility/manifests/org.example.ledger.json",
		"desktop-activation-receipt":   "usr/share/xnix/compatibility/activation-receipts/org.example.ledger.json",
		"managed-launcher-artifact":    "usr/share/xnix/compatibility/launcher-artifacts/xnix-compat-launch.json",
	}
	for _, file := range result.WrittenFiles {
		if file.RelativePath != expectedFiles[file.ID] {
			t.Fatalf("unexpected relative path for %s: %q", file.ID, file.RelativePath)
		}
		if !file.Written || file.HostRootModified || file.BackendDetailsExposed {
			t.Fatalf("staged file has unsafe flags: %#v", file)
		}
		data, err := os.ReadFile(filepath.Join(root, file.RelativePath))
		if err != nil {
			t.Fatalf("staged file %s was not written: %v", file.RelativePath, err)
		}
		if sha256Hex(string(data)) != file.SHA256 {
			t.Fatalf("staged file %s digest mismatch", file.RelativePath)
		}
	}

	desktopEntry := readStageFile(t, root, expectedFiles["desktop-entry"])
	if !strings.Contains(desktopEntry, "Name=Example Ledger\n") ||
		!strings.Contains(desktopEntry, "Exec=xnix-compat-launch --app org.example.ledger %U\n") ||
		!strings.Contains(desktopEntry, "MimeType=application/x-xnix-abc;application/x-xnix-xls;\n") {
		t.Fatalf("unexpected desktop entry content: %q", desktopEntry)
	}
	serviceMenu := readStageFile(t, root, expectedFiles["dolphin-service-menu"])
	if !strings.Contains(serviceMenu, "Name=Open with Xnix Compatibility\n") ||
		!strings.Contains(serviceMenu, "Exec=xnix-compat-open %U\n") {
		t.Fatalf("unexpected Dolphin service menu content: %q", serviceMenu)
	}
	launcherArtifact := readStageFile(t, root, expectedFiles["managed-launcher-artifact"])
	if !strings.Contains(launcherArtifact, "\"command\": \"xnix-compat-launch\"") ||
		!strings.Contains(launcherArtifact, "\"source_package\": \"cmd/xnix-compat-launch\"") ||
		!strings.Contains(launcherArtifact, "\"build_output\": \"usr/local/bin/xnix-compat-launch\"") ||
		!strings.Contains(launcherArtifact, "\"runtime_method\": \"PreviewKnownPortableLaunchBridge\"") ||
		!strings.Contains(launcherArtifact, "\"dispatch_gate\": \"managed-known-app-guest-smoke\"") ||
		!strings.Contains(launcherArtifact, "\"binary_copied\": false") ||
		!strings.Contains(launcherArtifact, "\"host_root_modified\": false") {
		t.Fatalf("unexpected managed launcher artifact content: %q", launcherArtifact)
	}
	receipt := readStageFile(t, root, expectedFiles["desktop-activation-receipt"])
	if !strings.Contains(receipt, "\"requires_matching_sha256\": true") ||
		!strings.Contains(receipt, "\"host_root_modified\": false") ||
		!strings.Contains(receipt, "\"id\": \"managed-launcher-artifact\"") {
		t.Fatalf("unexpected receipt content: %q", receipt)
	}

	encoded, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("Marshal returned error: %v", err)
	}
	text := strings.ToLower(string(encoded))
	if strings.Contains(text, strings.ToLower(root)) {
		t.Fatalf("stage result exposes staging root path: %s", text)
	}
	for _, forbidden := range []string{"prefix", ".exe", "program files", "qemu-system", "proton", "wine ", "virtual machine"} {
		if strings.Contains(text, forbidden) {
			t.Fatalf("stage result exposes forbidden term %q: %s", forbidden, text)
		}
	}
}

func TestStageRejectsConflictsAndBlockedProductionMode(t *testing.T) {
	root := t.TempDir()
	plan := testPlan(t)
	conflictPath := filepath.Join(root, "usr/share/applications/xnix-org.example.ledger.desktop")
	if err := os.MkdirAll(filepath.Dir(conflictPath), 0o755); err != nil {
		t.Fatalf("MkdirAll returned error: %v", err)
	}
	if err := os.WriteFile(conflictPath, []byte("existing"), 0o644); err != nil {
		t.Fatalf("WriteFile returned error: %v", err)
	}
	if _, err := Stage(StageRequest{Root: root, Mode: "development", Plan: plan}); err == nil ||
		!strings.Contains(err.Error(), "refusing to overwrite existing staged file") {
		t.Fatalf("Stage must reject existing files, got %v", err)
	}

	if _, err := Stage(StageRequest{Root: t.TempDir(), Mode: "production", Plan: plan}); err == nil ||
		!strings.Contains(err.Error(), "desktop activation staging is blocked") {
		t.Fatalf("Stage must reject blocked production activation, got %v", err)
	}
	if _, err := Stage(StageRequest{Root: string(filepath.Separator), Mode: "development", Plan: plan}); err == nil ||
		!strings.Contains(err.Error(), "refusing to stage desktop activation into filesystem root") {
		t.Fatalf("Stage must reject filesystem root staging, got %v", err)
	}
}

func testPlan(t *testing.T) appidentity.Plan {
	t.Helper()
	plan, err := appidentity.NewPlanWithProvenance(appidentity.Recipe{
		ID:                  "org.example.ledger",
		Name:                "Example Ledger",
		Icon:                "office-chart-area",
		Mode:                "automatic",
		SupportedExtensions: []string{".xls", ".abc"},
	}, appidentity.Provenance{
		Source:          "registry",
		RegistryName:    "test-registry",
		DigestVerified:  true,
		SignatureStatus: "development-only",
	})
	if err != nil {
		t.Fatalf("NewPlanWithProvenance returned error: %v", err)
	}
	return plan
}

func readStageFile(t *testing.T, root string, relativePath string) string {
	t.Helper()
	data, err := os.ReadFile(filepath.Join(root, relativePath))
	if err != nil {
		t.Fatalf("ReadFile(%s) returned error: %v", relativePath, err)
	}
	return string(data)
}

func sameStrings(got []string, want []string) bool {
	if len(got) != len(want) {
		return false
	}
	for index := range got {
		if got[index] != want[index] {
			return false
		}
	}
	return true
}
