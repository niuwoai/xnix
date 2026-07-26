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

	"xnix.local/xnix/internal/runtime/appidentity"
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

func TestDesktopActivationStageCommandStagesExternalImportedAppHandleLauncher(t *testing.T) {
	tempDir := t.TempDir()
	stateRoot := filepath.Join(tempDir, "state")
	stagingRoot := filepath.Join(tempDir, "stage")
	executablePath := filepath.Join(tempDir, "ExternalGui.exe")
	if err := os.WriteFile(executablePath, []byte{'M', 'Z', 0x90, 0x00, 'x', 'n', 'i', 'x'}, 0o600); err != nil {
		t.Fatalf("WriteFile executable returned error: %v", err)
	}
	record, err := appidentity.RecordExternalWinAppImport(appidentity.ExternalWinAppImportRequest{
		Version:        "0.2.640-test",
		StateRoot:      stateRoot,
		ExecutablePath: executablePath,
		AppID:          "org.xnix.external.gui",
		DisplayName:    "External GUI",
		AppVersion:     "0.2.640-test",
	})
	if err != nil {
		t.Fatalf("RecordExternalWinAppImport returned error: %v", err)
	}
	recordPath := filepath.Join(stateRoot, filepath.FromSlash(record.RecordRelativePath))

	var output bytes.Buffer
	err = run([]string{
		"desktop-activation-stage",
		"--external-app-import-record", recordPath,
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
	if payload["application_id"] != "org.xnix.external.gui" ||
		payload["display_name"] != "External GUI" ||
		payload["desktop_file"] != "xnix-org.xnix.external.gui.desktop" ||
		payload["external_app_handle"] != "org.xnix.external.gui" ||
		payload["desktop_exec_uses_external_app_handle"] != true ||
		payload["external_app_desktop_handle_ready"] != true ||
		payload["desktop_exec_uses_raw_import_record"] != false ||
		payload["desktop_exec_uses_state_root"] != false ||
		payload["written_file_count"] != float64(5) ||
		payload["mimeapps_written"] != false ||
		payload["launch_enabled"] != false ||
		payload["backend_launch_enabled"] != false ||
		payload["execution_started"] != false ||
		payload["host_root_modified"] != false ||
		payload["staging_root_path_exposed"] != false {
		t.Fatalf("unexpected external imported app stage payload: %#v", payload)
	}
	desktopEntryPath := filepath.Join(stagingRoot, "usr/share/applications/xnix-org.xnix.external.gui.desktop")
	desktopEntry, err := os.ReadFile(desktopEntryPath)
	if err != nil {
		t.Fatalf("external imported app desktop entry was not staged: %v", err)
	}
	if !bytes.Contains(desktopEntry, []byte("Exec=xnix-compat-launch --external-app-handle org.xnix.external.gui %U\n")) ||
		bytes.Contains(desktopEntry, []byte("MimeType=")) {
		t.Fatalf("unexpected external imported app desktop entry:\n%s", desktopEntry)
	}
	for _, forbidden := range []string{recordPath, stateRoot, stagingRoot, executablePath, "ExternalGui.exe", "wine ", "docker", "qemu-system", "--state-root", "--external-app-import-record"} {
		if strings.Contains(strings.ToLower(output.String()), strings.ToLower(forbidden)) ||
			strings.Contains(strings.ToLower(string(desktopEntry)), strings.ToLower(forbidden)) {
			t.Fatalf("external imported app staging exposed forbidden term %q\noutput=%s\nentry=%s", forbidden, output.String(), string(desktopEntry))
		}
	}
}

func TestDesktopActivationStageCommandStagesCanonicalLauncherOnlyGUIApp(t *testing.T) {
	stagingRoot := t.TempDir()

	var output bytes.Buffer
	err := run([]string{
		"desktop-activation-stage",
		"--registry", "../../runtime/recipes/registry.json",
		"--app", "org.xnix.apps.mines",
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
	if payload["application_id"] != "org.xnix.apps.mines" ||
		payload["display_name"] != "Mines" ||
		payload["desktop_file"] != "xnix-org.xnix.apps.mines.desktop" ||
		payload["written_file_count"] != float64(5) ||
		payload["mimeapps_written"] != false ||
		payload["launch_enabled"] != false ||
		payload["backend_launch_enabled"] != false ||
		payload["execution_started"] != false ||
		payload["host_root_modified"] != false {
		t.Fatalf("unexpected launcher-only desktop activation stage payload: %#v", payload)
	}
	writtenFileIDs := payload["written_file_ids"].([]any)
	if containsAnyString(writtenFileIDs, "mimeapps-list") {
		t.Fatalf("launcher-only stage must not write mimeapps-list: %#v", writtenFileIDs)
	}

	desktopEntryPath := filepath.Join(stagingRoot, "usr/share/applications/xnix-org.xnix.apps.mines.desktop")
	desktopEntry, err := os.ReadFile(desktopEntryPath)
	if err != nil {
		t.Fatalf("desktop entry was not staged: %v", err)
	}
	if !bytes.Contains(desktopEntry, []byte("Name=Mines\n")) ||
		!bytes.Contains(desktopEntry, []byte("Icon=applications-games\n")) ||
		!bytes.Contains(desktopEntry, []byte("Exec=xnix-compat-launch --app org.xnix.apps.mines %U\n")) ||
		!bytes.Contains(desktopEntry, []byte("X-Xnix-ApplicationId=org.xnix.apps.mines\n")) ||
		bytes.Contains(desktopEntry, []byte("MimeType=")) ||
		bytes.Contains(desktopEntry, []byte("winemine.exe")) ||
		bytes.Contains(desktopEntry, []byte("qemu-system")) ||
		bytes.Contains(desktopEntry, []byte("wine ")) {
		t.Fatalf("unexpected launcher-only desktop entry:\n%s", desktopEntry)
	}
	if _, err := os.Stat(filepath.Join(stagingRoot, "usr/share/applications/mimeapps.list")); !os.IsNotExist(err) {
		t.Fatalf("launcher-only stage must not create mimeapps.list, stat error: %v", err)
	}
}

func TestDesktopActivationStageCommandStagesRecipeBackedContainerGUIApp(t *testing.T) {
	stagingRoot := t.TempDir()

	var output bytes.Buffer
	err := run([]string{
		"desktop-activation-stage",
		"--registry", "../../runtime/recipes/registry.json",
		"--app", "org.xnix.sample.notepad",
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
	if payload["application_id"] != "org.xnix.sample.notepad" ||
		payload["display_name"] != "Sample Notepad" ||
		payload["desktop_file"] != "xnix-org.xnix.sample.notepad.desktop" ||
		payload["written_file_count"] != float64(8) ||
		payload["launch_enabled"] != false ||
		payload["backend_launch_enabled"] != false ||
		payload["execution_started"] != false ||
		payload["host_root_modified"] != false {
		t.Fatalf("unexpected recipe-backed container GUI stage payload: %#v", payload)
	}
	writtenFileIDs := payload["written_file_ids"].([]any)
	for _, id := range []string{"desktop-entry", "mimeapps-list", "recipe-registry", "application-recipe", "managed-launcher-artifact"} {
		if !containsAnyString(writtenFileIDs, id) {
			t.Fatalf("recipe-backed stage missing file id %q: %#v", id, writtenFileIDs)
		}
	}

	desktopEntryPath := filepath.Join(stagingRoot, "usr/share/applications/xnix-org.xnix.sample.notepad.desktop")
	desktopEntry, err := os.ReadFile(desktopEntryPath)
	if err != nil {
		t.Fatalf("desktop entry was not staged: %v", err)
	}
	if !bytes.Contains(desktopEntry, []byte("Exec=xnix-compat-launch --app org.xnix.sample.notepad --registry /usr/share/xnix/compatibility/recipes/registry.json %U\n")) ||
		bytes.Contains(desktopEntry, []byte("notepad.exe")) ||
		bytes.Contains(desktopEntry, []byte("docker")) ||
		bytes.Contains(desktopEntry, []byte("wine ")) {
		t.Fatalf("unexpected recipe-backed desktop entry:\n%s", desktopEntry)
	}

	registryPath := filepath.Join(stagingRoot, "usr/share/xnix/compatibility/recipes/registry.json")
	registryData, err := os.ReadFile(registryPath)
	if err != nil {
		t.Fatalf("recipe registry was not staged: %v", err)
	}
	if !bytes.Contains(registryData, []byte(`"id": "org.xnix.sample.notepad"`)) ||
		!bytes.Contains(registryData, []byte(`"path": "org.xnix.sample.notepad.json"`)) {
		t.Fatalf("unexpected staged recipe registry:\n%s", registryData)
	}
	recipePath := filepath.Join(stagingRoot, "usr/share/xnix/compatibility/recipes/org.xnix.sample.notepad.json")
	recipeData, err := os.ReadFile(recipePath)
	if err != nil {
		t.Fatalf("application recipe was not staged: %v", err)
	}
	if !bytes.Contains(recipeData, []byte(`"app": "notepad.exe"`)) ||
		!bytes.Contains(recipeData, []byte(`"window_match": "notepad.exe"`)) {
		t.Fatalf("staged recipe must preserve Runtime container GUI hints:\n%s", recipeData)
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

func TestDesktopActivationStageCommandCanCopyManagedLauncherBinary(t *testing.T) {
	registryPath := writeStageCommandRegistry(t)
	stagingRoot := t.TempDir()
	launcherSource := filepath.Join(t.TempDir(), "xnix-compat-launch")
	launcherContent := []byte("#!/bin/sh\nprintf 'XNIX_MANAGED_LAUNCHER_OK\\n'\n")
	if err := os.WriteFile(launcherSource, launcherContent, 0o755); err != nil {
		t.Fatalf("WriteFile launcher source returned error: %v", err)
	}

	var output bytes.Buffer
	err := run([]string{
		"desktop-activation-stage",
		"--registry", registryPath,
		"--app", "org.example.ledger",
		"--mode", "development",
		"--staging-root", stagingRoot,
		"--managed-launcher-bin", launcherSource,
	}, &output)
	if err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["written_file_count"] != float64(7) {
		t.Fatalf("unexpected written file count: %#v", payload)
	}
	writtenFileIDs := payload["written_file_ids"].([]any)
	if !containsAnyString(writtenFileIDs, "managed-launcher-executable") {
		t.Fatalf("written_file_ids must include managed-launcher-executable: %#v", writtenFileIDs)
	}
	launcherPath := filepath.Join(stagingRoot, "usr/local/bin/xnix-compat-launch")
	launcherData, err := os.ReadFile(launcherPath)
	if err != nil {
		t.Fatalf("managed launcher executable was not staged: %v", err)
	}
	if !bytes.Equal(launcherData, launcherContent) {
		t.Fatalf("unexpected managed launcher executable:\n%s", launcherData)
	}
	info, err := os.Stat(launcherPath)
	if err != nil {
		t.Fatalf("Stat launcher returned error: %v", err)
	}
	if mode := info.Mode() & 0o777; mode != 0o755 {
		t.Fatalf("managed launcher executable mode = %04o, want 0755", mode)
	}
	artifact, err := os.ReadFile(filepath.Join(stagingRoot, "usr/share/xnix/compatibility/launcher-artifacts/xnix-compat-launch.json"))
	if err != nil {
		t.Fatalf("managed launcher artifact was not staged: %v", err)
	}
	if !bytes.Contains(artifact, []byte(`"binary_copied": true`)) ||
		!bytes.Contains(artifact, []byte(`"executable_staged": true`)) ||
		!bytes.Contains(artifact, []byte(`"staged_executable": "usr/local/bin/xnix-compat-launch"`)) {
		t.Fatalf("managed launcher artifact did not record executable staging:\n%s", artifact)
	}
	if strings.Contains(strings.ToLower(output.String()), strings.ToLower(stagingRoot)) ||
		strings.Contains(strings.ToLower(output.String()), strings.ToLower(launcherSource)) {
		t.Fatalf("desktop activation stage output exposed local paths: %s", output.String())
	}
}

func containsAnyString(values []any, expected string) bool {
	for _, value := range values {
		if value == expected {
			return true
		}
	}
	return false
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
