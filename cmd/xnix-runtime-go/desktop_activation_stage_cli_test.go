package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestDesktopActivationStageCommandWritesOnlyInsideStagingRoot(t *testing.T) {
	registryPath := writeStageCommandRegistry(t)
	stagingRoot := t.TempDir()

	var output bytes.Buffer
	err := run([]string{
		"desktop-activation-stage",
		"--registry", registryPath,
		"--app", "org.example.ledger",
		"--mode", "development",
		"--staging-root", stagingRoot,
	}, &output)
	if err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["schema_version"] != "xnix.runtime.desktop_activation_stage.v1" ||
		payload["request_type"] != "desktop-activation-stage" ||
		payload["stage_type"] != "kde-desktop-activation-test-root-stage" ||
		payload["desktop"] != "KDE Plasma" ||
		payload["application_id"] != "org.example.ledger" ||
		payload["desktop_file"] != "xnix-org.example.ledger.desktop" ||
		payload["install_mode"] != "development" ||
		payload["preflight_decision"] != "development-staging-ready" {
		t.Fatalf("unexpected desktop activation stage payload: %#v", payload)
	}
	if payload["runtime_owned"] != true ||
		payload["go_runtime_backed"] != true ||
		payload["kde_policy_owner"] != false ||
		payload["staging_root_required"] != true ||
		payload["staging_root_path_exposed"] != false ||
		payload["host_root_allowed"] != false ||
		payload["file_writes_performed"] != true ||
		payload["desktop_files_written"] != true ||
		payload["mimeapps_written"] != true ||
		payload["manifest_written"] != true ||
		payload["receipt_written"] != true ||
		payload["rollback_receipt_written"] != true ||
		payload["launch_enabled"] != false ||
		payload["backend_launch_enabled"] != false ||
		payload["execution_started"] != false ||
		payload["host_root_modified"] != false ||
		payload["network_required"] != false ||
		payload["privileged_container_required"] != false ||
		payload["backend_details_exposed"] != false {
		t.Fatalf("unexpected desktop activation stage safety flags: %#v", payload)
	}
	if strings.Contains(strings.ToLower(output.String()), strings.ToLower(stagingRoot)) {
		t.Fatalf("desktop activation stage output exposed staging root: %s", output.String())
	}
	writtenFiles := payload["written_files"].([]any)
	if len(writtenFiles) != 6 || payload["written_file_count"] != float64(6) {
		t.Fatalf("unexpected written file count: %#v", payload)
	}
	desktopEntryPath := filepath.Join(stagingRoot, "usr/share/applications/xnix-org.example.ledger.desktop")
	desktopEntry, err := os.ReadFile(desktopEntryPath)
	if err != nil {
		t.Fatalf("desktop entry was not staged: %v", err)
	}
	if !bytes.Contains(desktopEntry, []byte("Exec=xnix-compat-launch --app org.example.ledger %U\n")) {
		t.Fatalf("unexpected desktop entry:\n%s", desktopEntry)
	}
	launcherArtifactPath := filepath.Join(stagingRoot, "usr/share/xnix/compatibility/launcher-artifacts/xnix-compat-launch.json")
	launcherArtifact, err := os.ReadFile(launcherArtifactPath)
	if err != nil {
		t.Fatalf("managed launcher artifact was not staged: %v", err)
	}
	if !bytes.Contains(launcherArtifact, []byte(`"command": "xnix-compat-launch"`)) ||
		!bytes.Contains(launcherArtifact, []byte(`"source_package": "cmd/xnix-compat-launch"`)) ||
		!bytes.Contains(launcherArtifact, []byte(`"runtime_method": "PreviewKnownPortableLaunchBridge"`)) ||
		!bytes.Contains(launcherArtifact, []byte(`"binary_copied": false`)) {
		t.Fatalf("unexpected managed launcher artifact:\n%s", launcherArtifact)
	}
	receiptPath := filepath.Join(stagingRoot, "usr/share/xnix/compatibility/activation-receipts/org.example.ledger.json")
	receipt, err := os.ReadFile(receiptPath)
	if err != nil {
		t.Fatalf("receipt was not staged: %v", err)
	}
	if !bytes.Contains(receipt, []byte(`"requires_matching_sha256": true`)) ||
		!bytes.Contains(receipt, []byte(`"host_root_modified": false`)) ||
		!bytes.Contains(receipt, []byte(`"id": "managed-launcher-artifact"`)) {
		t.Fatalf("unexpected receipt:\n%s", receipt)
	}
}

func TestDesktopActivationStageCommandRejectsMissingRoot(t *testing.T) {
	registryPath := writeStageCommandRegistry(t)

	var output bytes.Buffer
	err := run([]string{
		"desktop-activation-stage",
		"--registry", registryPath,
		"--app", "org.example.ledger",
		"--mode", "development",
	}, &output)
	if err == nil || !strings.Contains(err.Error(), "desktop-activation-stage requires --staging-root") {
		t.Fatalf("desktop activation stage must require staging root, got %v", err)
	}
}

func writeStageCommandRegistry(t *testing.T) string {
	t.Helper()
	root := t.TempDir()
	recipeData := []byte(`{"id":"org.example.ledger","name":"Example Ledger","icon":"office-chart-area","mode":"automatic","supported_extensions":[".abc",".xls"]}`)
	sum := sha256.Sum256(recipeData)
	digest := hex.EncodeToString(sum[:])
	if err := os.WriteFile(filepath.Join(root, "org.example.ledger.json"), recipeData, 0o600); err != nil {
		t.Fatalf("WriteFile recipe returned error: %v", err)
	}
	registryPath := filepath.Join(root, "registry.json")
	registryData := []byte(`{"schema_version":1,"registry_name":"test-registry","recipes":[{"id":"org.example.ledger","path":"org.example.ledger.json","sha256":"` + digest + `","signature_status":"development-only"}]}`)
	if err := os.WriteFile(registryPath, registryData, 0o600); err != nil {
		t.Fatalf("WriteFile registry returned error: %v", err)
	}
	return registryPath
}
