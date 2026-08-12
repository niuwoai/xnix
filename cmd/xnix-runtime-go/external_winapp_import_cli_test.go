package main

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"xnix.local/xnix/internal/runtime/appidentity"
)

func TestExternalWinAppImportRecordCommandPersistsRuntimeManagedExecutable(t *testing.T) {
	tempDir := t.TempDir()
	stateRoot := filepath.Join(tempDir, "state")
	executablePath := filepath.Join(tempDir, "ExternalTool.exe")
	executable := []byte{'M', 'Z', 0x90, 0x00, 'x', 'n', 'i', 'x'}
	if err := os.WriteFile(executablePath, executable, 0o600); err != nil {
		t.Fatalf("WriteFile executable returned error: %v", err)
	}

	var output bytes.Buffer
	if err := run([]string{
		"external-winapp-import-record",
		"--state-root", stateRoot,
		"--executable", executablePath,
		"--app-id", "org.xnix.external.tool",
		"--display-name", "External Tool",
	}, &output); err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v\n%s", err, output.String())
	}
	if payload["version"] != currentProjectVersion(t) ||
		payload["schema_version"] != "xnix.runtime.external_winapp_import_record.v1" ||
		payload["request_type"] != "external-winapp-import-record" ||
		payload["record_type"] != "external-windows-app-import-record" ||
		payload["application_id"] != "org.xnix.external.tool" ||
		payload["display_name"] != "External Tool" ||
		payload["app_version"] != currentProjectVersion(t) ||
		payload["executable_name"] != "ExternalTool.exe" ||
		payload["artifact_size_bytes"] != float64(len(executable)) ||
		payload["windows_executable_validated"] != true ||
		payload["artifact_copied"] != true ||
		payload["import_recorded"] != true ||
		payload["desktop_page_renderable"] != true ||
		payload["runtime_owned"] != true ||
		payload["go_runtime_backed"] != true ||
		payload["kde_policy_owner"] != false ||
		payload["launch_enabled"] != false ||
		payload["backend_launch_enabled"] != false ||
		payload["state_root_path_exposed"] != false ||
		payload["raw_executable_path_exposed"] != false ||
		payload["backend_details_exposed"] != false ||
		payload["host_root_modified"] != false ||
		payload["docker_socket_mounted"] != false ||
		payload["broad_host_mount_required"] != false {
		t.Fatalf("unexpected external import output: %#v", payload)
	}
	if strings.Contains(output.String(), stateRoot) ||
		strings.Contains(output.String(), executablePath) ||
		strings.Contains(output.String(), "docker run") ||
		strings.Contains(output.String(), "/var/run/docker.sock") {
		t.Fatalf("external import output exposed unsafe details: %s", output.String())
	}
	artifactRelativePath := payload["artifact_relative_path"].(string)
	artifactBytes, err := os.ReadFile(filepath.Join(stateRoot, filepath.FromSlash(artifactRelativePath)))
	if err != nil {
		t.Fatalf("ReadFile imported artifact returned error: %v", err)
	}
	if string(artifactBytes) != string(executable) {
		t.Fatalf("imported artifact mismatch: %q != %q", string(artifactBytes), string(executable))
	}
	recordPath := filepath.Join(stateRoot, filepath.FromSlash(payload["record_relative_path"].(string)))

	var pageOutput bytes.Buffer
	if err := run([]string{
		"kde-center-page-preview",
		"--external-app-import-record", recordPath,
		"--decision", "approved",
	}, &pageOutput); err != nil {
		t.Fatalf("KDE center import record page returned error: %v", err)
	}
	var page map[string]any
	if err := json.Unmarshal(pageOutput.Bytes(), &page); err != nil {
		t.Fatalf("Unmarshal page output returned error: %v\n%s", err, pageOutput.String())
	}
	if page["application_id"] != "org.xnix.external.tool" ||
		page["application_name"] != "External Tool" ||
		page["icon"] != "application-x-executable" ||
		page["desktop_file"] != "xnix-org.xnix.external.tool.desktop" ||
		page["launch_enabled"] != false ||
		page["backend_process_started"] != false ||
		page["backend_details_exposed"] != false ||
		page["host_root_modified"] != false {
		t.Fatalf("unexpected KDE page from external import record: %#v", page)
	}
	if strings.Contains(pageOutput.String(), stateRoot) ||
		strings.Contains(pageOutput.String(), executablePath) ||
		strings.Contains(pageOutput.String(), "docker run") ||
		strings.Contains(pageOutput.String(), "/var/run/docker.sock") {
		t.Fatalf("external import KDE page exposed unsafe details: %s", pageOutput.String())
	}
}

func TestExternalWinAppImportRecordCommandRejectsNonPEExecutable(t *testing.T) {
	tempDir := t.TempDir()
	executablePath := filepath.Join(tempDir, "ExternalTool.exe")
	if err := os.WriteFile(executablePath, []byte("not a PE"), 0o600); err != nil {
		t.Fatalf("WriteFile executable returned error: %v", err)
	}

	var output bytes.Buffer
	err := run([]string{
		"external-winapp-import-record",
		"--state-root", filepath.Join(tempDir, "state"),
		"--executable", executablePath,
		"--app-id", "org.xnix.external.tool",
		"--display-name", "External Tool",
	}, &output)
	if err == nil || !strings.Contains(err.Error(), "MZ executable") {
		t.Fatalf("expected non-PE import to fail closed, got err=%v output=%s", err, output.String())
	}
}

func TestExternalWinAppRunCommandRunsImportedExecutable(t *testing.T) {
	tempDir := t.TempDir()
	stateRoot := filepath.Join(tempDir, "state")
	executablePath := filepath.Join(tempDir, "ExternalGui.exe")
	if err := os.WriteFile(executablePath, []byte{'M', 'Z', 0x90, 0x00, 'x', 'n', 'i', 'x'}, 0o600); err != nil {
		t.Fatalf("WriteFile executable returned error: %v", err)
	}
	var importOutput bytes.Buffer
	if err := run([]string{
		"external-winapp-import-record",
		"--state-root", stateRoot,
		"--executable", executablePath,
		"--app-id", "org.xnix.external.gui",
		"--display-name", "External GUI",
	}, &importOutput); err != nil {
		t.Fatalf("import run returned error: %v", err)
	}
	var importPayload map[string]any
	if err := json.Unmarshal(importOutput.Bytes(), &importPayload); err != nil {
		t.Fatalf("Unmarshal import output returned error: %v", err)
	}
	recordPath := filepath.Join(stateRoot, filepath.FromSlash(importPayload["record_relative_path"].(string)))

	dockerLog := filepath.Join(tempDir, "fake-docker.log")
	dockerPath := filepath.Join(tempDir, "fake-docker")
	dockerBody := "#!/bin/sh\n" +
		"printf '%s\\n' \"$@\" >> \"" + dockerLog + "\"\n" +
		"if test \"$1 $2\" = 'image inspect'; then printf 'linux/amd64\\n'; exit 0; fi\n" +
		"if test \"$1\" = 'create'; then printf 'fake-x-gui-container\\n'; exit 0; fi\n" +
		"if test \"$1\" = 'cp'; then exit 0; fi\n" +
		"if test \"$1 $2\" = 'start -a'; then " +
		"printf 'XNIX_X_GUI_XSERVER_STARTED=true\\n'\n" +
		"printf 'XNIX_X_GUI_WINE_BOOTSTRAP_ATTEMPTED=true\\n'\n" +
		"printf 'XNIX_X_GUI_FILE_ARGS_PASSED=0\\n'\n" +
		"printf 'XNIX_X_GUI_FILE_ARGS_WINEPATH_TRANSLATED=0\\n'\n" +
		"printf '0x700001 \"External GUI\": (\"ExternalGui.exe\" \"ExternalGui.exe\") 320x160+20+20 +20+20\\n'\n" +
		"printf 'XNIX_X_GUI_WINDOW_OBSERVED=true\\n'; exit 0; fi\n" +
		"if test \"$1\" = 'rm'; then exit 0; fi\n" +
		"exit 2\n"
	if err := os.WriteFile(dockerPath, []byte(dockerBody), 0o700); err != nil {
		t.Fatalf("WriteFile docker returned error: %v", err)
	}
	outputPath := filepath.Join(tempDir, "external-run.json")
	var output bytes.Buffer
	if err := run([]string{
		"windows-external-app-run",
		"--external-app-import-record", recordPath,
		"--window-match", "External GUI",
		"--image", "local/wine-x-gui:test",
		"--platform", "linux/amd64",
		"--docker", dockerPath,
		"--timeout", "5s",
		"--output", outputPath,
	}, &output); err != nil {
		t.Fatalf("external app run returned error: %v", err)
	}
	written, err := os.ReadFile(outputPath)
	if err != nil {
		t.Fatalf("ReadFile output returned error: %v", err)
	}
	if string(written) != output.String() {
		t.Fatalf("written output must match stdout\nstdout=%s\nwritten=%s", output.String(), string(written))
	}
	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal run output returned error: %v\n%s", err, output.String())
	}
	if payload["schema_version"] != "xnix.runtime.external_winapp_run.v1" ||
		payload["request_type"] != "windows-external-app-run" ||
		payload["status"] != "passed" ||
		payload["application_id"] != "org.xnix.external.gui" ||
		payload["display_name"] != "External GUI" ||
		payload["app_version"] != currentProjectVersion(t) ||
		payload["executable_name"] != "ExternalGui.exe" ||
		payload["external_app_import_record_consumed"] != true ||
		payload["imported_artifact_digest_verified"] != true ||
		payload["imported_artifact_sha256"] != importPayload["artifact_sha256"] ||
		payload["runtime_run_requested"] != true ||
		payload["runtime_run_executed"] != true ||
		payload["execution_started"] != true ||
		payload["backend_process_started"] != true ||
		payload["container_runtime_used"] != true ||
		payload["container_network_mode"] != "none" ||
		payload["container_host_mount_count"] != float64(0) ||
		payload["x_window_observed"] != true ||
		payload["window_observed"] != true ||
		payload["desktop_launch_enabled"] != false ||
		payload["action_execution_enabled"] != false ||
		payload["backend_details_exposed"] != false ||
		payload["raw_import_record_path_exposed"] != false ||
		payload["raw_state_root_path_exposed"] != false ||
		payload["raw_executable_path_exposed"] != false ||
		payload["host_root_modified"] != false ||
		payload["docker_socket_mounted"] != false {
		t.Fatalf("unexpected external app run payload: %#v", payload)
	}
	runtimePayload := payload["runtime_payload"].(map[string]any)
	if runtimePayload["external_app_import_record_consumed"] != true ||
		runtimePayload["imported_artifact_digest_verified"] != true ||
		runtimePayload["imported_artifact_sha256"] != importPayload["artifact_sha256"] {
		t.Fatalf("unexpected nested runtime payload: %#v", runtimePayload)
	}
	if strings.Contains(output.String(), recordPath) ||
		strings.Contains(output.String(), stateRoot) ||
		strings.Contains(output.String(), executablePath) ||
		strings.Contains(output.String(), dockerPath) ||
		strings.Contains(output.String(), "docker run") ||
		strings.Contains(output.String(), "/var/run/docker.sock") ||
		strings.Contains(output.String(), "--network host") ||
		strings.Contains(output.String(), "--privileged") {
		t.Fatalf("external app run output exposed unsafe details: %s", output.String())
	}
}

func TestExternalWinAppRunCommandRunsImportedExecutableByHandle(t *testing.T) {
	tempDir := t.TempDir()
	stateRoot := filepath.Join(tempDir, "state")
	executablePath := filepath.Join(tempDir, "ExternalGui.exe")
	if err := os.WriteFile(executablePath, []byte{'M', 'Z', 0x90, 0x00, 'x', 'n', 'i', 'x'}, 0o600); err != nil {
		t.Fatalf("WriteFile executable returned error: %v", err)
	}
	var importOutput bytes.Buffer
	if err := run([]string{
		"external-winapp-import-record",
		"--state-root", stateRoot,
		"--executable", executablePath,
		"--app-id", "org.xnix.external.gui",
		"--display-name", "External GUI",
	}, &importOutput); err != nil {
		t.Fatalf("import run returned error: %v", err)
	}
	var importPayload map[string]any
	if err := json.Unmarshal(importOutput.Bytes(), &importPayload); err != nil {
		t.Fatalf("Unmarshal import output returned error: %v", err)
	}
	documentPath := filepath.Join(tempDir, "report.docx")
	if err := os.WriteFile(documentPath, []byte("external app file-open fixture"), 0o600); err != nil {
		t.Fatalf("WriteFile document returned error: %v", err)
	}
	documentURI := "file://" + filepath.ToSlash(documentPath)

	dockerLog := filepath.Join(tempDir, "fake-docker.log")
	dockerPath := filepath.Join(tempDir, "fake-docker")
	dockerBody := "#!/bin/sh\n" +
		"printf '%s\\n' \"$@\" >> \"" + dockerLog + "\"\n" +
		"if test \"$1 $2\" = 'image inspect'; then printf 'linux/amd64\\n'; exit 0; fi\n" +
		"if test \"$1\" = 'create'; then printf 'fake-x-gui-container\\n'; exit 0; fi\n" +
		"if test \"$1\" = 'cp'; then exit 0; fi\n" +
		"if test \"$1 $2\" = 'start -a'; then " +
		"printf 'XNIX_X_GUI_XSERVER_STARTED=true\\n'\n" +
		"printf 'XNIX_X_GUI_WINE_BOOTSTRAP_ATTEMPTED=true\\n'\n" +
		"printf 'XNIX_X_GUI_FILE_ARGS_PASSED=1\\n'\n" +
		"printf 'XNIX_X_GUI_FILE_ARGS_WINEPATH_TRANSLATED=1\\n'\n" +
		"printf '0x700001 \"External GUI\": (\"ExternalGui.exe\" \"ExternalGui.exe\") 320x160+20+20 +20+20\\n'\n" +
		"printf 'XNIX_X_GUI_WINDOW_OBSERVED=true\\n'; exit 0; fi\n" +
		"if test \"$1\" = 'rm'; then exit 0; fi\n" +
		"exit 2\n"
	if err := os.WriteFile(dockerPath, []byte(dockerBody), 0o700); err != nil {
		t.Fatalf("WriteFile docker returned error: %v", err)
	}
	var output bytes.Buffer
	if err := run([]string{
		"windows-external-app-run",
		"--state-root", stateRoot,
		"--external-app-handle", "org.xnix.external.gui",
		"--window-match", "External GUI",
		"--image", "local/wine-x-gui:test",
		"--platform", "linux/amd64",
		"--docker", dockerPath,
		"--timeout", "5s",
		documentURI,
	}, &output); err != nil {
		t.Fatalf("external app handle run returned error: %v", err)
	}
	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal run output returned error: %v\n%s", err, output.String())
	}
	if payload["schema_version"] != "xnix.runtime.external_winapp_run.v1" ||
		payload["request_type"] != "windows-external-app-run" ||
		payload["status"] != "passed" ||
		payload["application_id"] != "org.xnix.external.gui" ||
		payload["external_app_import_record_consumed"] != true ||
		payload["external_app_handle_consumed"] != true ||
		payload["external_app_handle"] != "org.xnix.external.gui" ||
		payload["external_desktop_argument_count"] != float64(1) ||
		payload["external_file_uri_arguments_accepted"] != true ||
		payload["external_file_open_requested"] != true ||
		payload["external_file_bridge_copy_enabled"] != true ||
		payload["external_file_bridge_copied_count"] != float64(1) ||
		payload["external_file_bridge_arguments_passed"] != true ||
		payload["external_file_bridge_argument_observed_count"] != float64(1) ||
		payload["external_file_bridge_winepath_translated"] != true ||
		payload["external_file_bridge_winepath_translated_count"] != float64(1) ||
		payload["external_file_bridge_ready"] != true ||
		payload["external_file_bridge_mount_enabled"] != false ||
		payload["raw_file_uri_arguments_exposed"] != false ||
		payload["imported_artifact_digest_verified"] != true ||
		payload["imported_artifact_sha256"] != importPayload["artifact_sha256"] ||
		payload["raw_external_app_handle_path_exposed"] != false ||
		payload["raw_import_record_path_exposed"] != false ||
		payload["raw_state_root_path_exposed"] != false ||
		payload["raw_executable_path_exposed"] != false ||
		payload["container_network_mode"] != "none" ||
		payload["container_host_mount_count"] != float64(0) ||
		payload["x_window_observed"] != true ||
		payload["window_observed"] != true {
		t.Fatalf("unexpected external app handle run payload: %#v", payload)
	}
	dockerInvocation, err := os.ReadFile(dockerLog)
	if err != nil {
		t.Fatalf("ReadFile docker log returned error: %v", err)
	}
	if !strings.Contains(string(dockerInvocation), documentPath) ||
		!strings.Contains(string(dockerInvocation), "/file-1-report.docx") ||
		!strings.Contains(string(dockerInvocation), "XNIX_GUI_FILE_ARGS=/file-1-report.docx") {
		t.Fatalf("fake Docker did not receive copied file-open flow: %s", string(dockerInvocation))
	}
	if strings.Contains(output.String(), stateRoot) ||
		strings.Contains(output.String(), executablePath) ||
		strings.Contains(output.String(), dockerPath) ||
		strings.Contains(output.String(), documentPath) ||
		strings.Contains(output.String(), documentURI) ||
		strings.Contains(output.String(), "report.docx") ||
		strings.Contains(output.String(), "docker run") ||
		strings.Contains(output.String(), "/var/run/docker.sock") ||
		strings.Contains(output.String(), "--network host") ||
		strings.Contains(output.String(), "--privileged") {
		t.Fatalf("external app handle run output exposed unsafe details: %s", output.String())
	}
}

func TestExternalWinAppImportAndStageCommandImportsAndStagesDesktopActivation(t *testing.T) {
	tempDir := t.TempDir()
	stateRoot := filepath.Join(tempDir, "state")
	stagingRoot := filepath.Join(tempDir, "stage")
	executablePath := filepath.Join(tempDir, "ExternalGui.exe")
	if err := os.WriteFile(executablePath, []byte{'M', 'Z', 0x90, 0x00, 'x', 'n', 'i', 'x'}, 0o600); err != nil {
		t.Fatalf("WriteFile executable returned error: %v", err)
	}
	launcherSource := filepath.Join(tempDir, "xnix-compat-launch")
	launcherContent := []byte("#!/bin/sh\nprintf 'XNIX_EXTERNAL_STAGE_LAUNCHER_OK\\n'\n")
	if err := os.WriteFile(launcherSource, launcherContent, 0o755); err != nil {
		t.Fatalf("WriteFile launcher returned error: %v", err)
	}
	outputPath := filepath.Join(tempDir, "import-and-stage.json")
	var output bytes.Buffer
	if err := run([]string{
		"external-winapp-import-and-stage",
		"--state-root", stateRoot,
		"--executable", executablePath,
		"--app-id", "org.xnix.external.gui",
		"--display-name", "External GUI",
		"--mode", "development",
		"--staging-root", stagingRoot,
		"--managed-launcher-bin", launcherSource,
		"--output", outputPath,
	}, &output); err != nil {
		t.Fatalf("external app import-and-stage returned error: %v", err)
	}
	written, err := os.ReadFile(outputPath)
	if err != nil {
		t.Fatalf("ReadFile output returned error: %v", err)
	}
	if string(written) != output.String() {
		t.Fatalf("written output must match stdout\nstdout=%s\nwritten=%s", output.String(), string(written))
	}
	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal import-and-stage output returned error: %v\n%s", err, output.String())
	}
	if payload["schema_version"] != "xnix.runtime.desktop_activation_stage.v1" ||
		payload["request_type"] != "desktop-activation-stage" ||
		payload["application_id"] != "org.xnix.external.gui" ||
		payload["display_name"] != "External GUI" ||
		payload["desktop_file"] != "xnix-org.xnix.external.gui.desktop" ||
		payload["external_app_handle"] != "org.xnix.external.gui" ||
		payload["desktop_exec_uses_external_app_handle"] != true ||
		payload["external_app_desktop_handle_ready"] != true ||
		payload["desktop_exec_uses_raw_import_record"] != false ||
		payload["desktop_exec_uses_state_root"] != false ||
		payload["written_file_count"] != float64(6) ||
		payload["launch_enabled"] != false ||
		payload["backend_launch_enabled"] != false ||
		payload["execution_started"] != false ||
		payload["host_root_modified"] != false ||
		payload["staging_root_path_exposed"] != false {
		t.Fatalf("unexpected external import-and-stage payload: %#v", payload)
	}
	if !containsAnyString(payload["written_file_ids"].([]any), "managed-launcher-executable") {
		t.Fatalf("written_file_ids must include managed-launcher-executable: %#v", payload["written_file_ids"])
	}
	recordPath, err := appidentity.ExternalWinAppImportRecordPathFromHandle(stateRoot, "org.xnix.external.gui")
	if err != nil {
		t.Fatalf("ExternalWinAppImportRecordPathFromHandle returned error: %v", err)
	}
	if _, err := os.Stat(recordPath); err != nil {
		t.Fatalf("one-shot import-and-stage must persist the import record: %v", err)
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
	launcherPath := filepath.Join(stagingRoot, "usr/local/bin/xnix-compat-launch")
	launcherData, err := os.ReadFile(launcherPath)
	if err != nil {
		t.Fatalf("managed launcher executable was not staged: %v", err)
	}
	if !bytes.Equal(launcherData, launcherContent) {
		t.Fatalf("unexpected managed launcher executable:\n%s", launcherData)
	}
	for _, stagedPath := range []string{
		"usr/share/xnix/compatibility/manifests/org.xnix.external.gui.json",
		"usr/share/xnix/compatibility/activation-receipts/org.xnix.external.gui.json",
		"usr/share/xnix/compatibility/launcher-artifacts/xnix-compat-launch.json",
	} {
		if _, err := os.Stat(filepath.Join(stagingRoot, stagedPath)); err != nil {
			t.Fatalf("expected staged desktop activation artifact %s: %v", stagedPath, err)
		}
	}
	if strings.Contains(output.String(), stateRoot) ||
		strings.Contains(output.String(), stagingRoot) ||
		strings.Contains(output.String(), executablePath) ||
		strings.Contains(output.String(), launcherSource) ||
		strings.Contains(output.String(), recordPath) ||
		strings.Contains(output.String(), "/var/run/docker.sock") ||
		strings.Contains(output.String(), "--network host") ||
		strings.Contains(output.String(), "--privileged") {
		t.Fatalf("external import-and-stage output exposed unsafe details: %s", output.String())
	}
}

func TestExternalWinAppImportStageAndLaunchCommandInvokesStagedLauncher(t *testing.T) {
	tempDir := t.TempDir()
	stateRoot := filepath.Join(tempDir, "state")
	stagingRoot := filepath.Join(tempDir, "stage")
	executablePath := filepath.Join(tempDir, "ExternalGui.exe")
	if err := os.WriteFile(executablePath, []byte{'M', 'Z', 0x90, 0x00, 'x', 'n', 'i', 'x'}, 0o600); err != nil {
		t.Fatalf("WriteFile executable returned error: %v", err)
	}
	launcherLog := filepath.Join(tempDir, "launcher.log")
	launcherSource := filepath.Join(tempDir, "xnix-compat-launch-source")
	launcherBody := "#!/bin/sh\n" +
		"{\n" +
		"printf 'args=%s\\n' \"$*\"\n" +
		"printf 'state=%s\\n' \"$XNIX_EXTERNAL_APP_STATE_ROOT\"\n" +
		"printf 'image=%s\\n' \"$XNIX_WINE_IMAGE\"\n" +
		"printf 'platform=%s\\n' \"$XNIX_CONTAINER_PLATFORM\"\n" +
		"printf 'docker=%s\\n' \"$XNIX_DOCKER_BIN\"\n" +
		"printf 'packet=%s\\n' \"$XNIX_EXTERNAL_APP_DESKTOP_LAUNCH_PACKET_OUTPUT\"\n" +
		"} > \"" + launcherLog + "\"\n" +
		"if test -n \"$XNIX_EXTERNAL_APP_DESKTOP_LAUNCH_PACKET_OUTPUT\"; then\n" +
		"mkdir -p \"$(dirname \"$XNIX_EXTERNAL_APP_DESKTOP_LAUNCH_PACKET_OUTPUT\")\"\n" +
		"printf '%s\\n' '{\"schema_version\":\"xnix.runtime.desktop_external_winapp_launch_packet.v1\",\"request_type\":\"desktop-external-winapp-launch-packet-preview\",\"status\":\"passed\"}' > \"$XNIX_EXTERNAL_APP_DESKTOP_LAUNCH_PACKET_OUTPUT\"\n" +
		"fi\n" +
		"printf '%s\\n' '{\"version\":\"0.2.640-test\",\"schema_version\":\"xnix.runtime.external_winapp_run.v1\",\"request_type\":\"windows-external-app-run\",\"status\":\"passed\",\"application_id\":\"org.xnix.external.gui\",\"display_name\":\"External GUI\",\"app_version\":\"0.2.640-test\",\"executable_name\":\"ExternalGui.exe\",\"external_app_import_record_consumed\":true,\"external_app_handle_consumed\":true,\"external_app_handle\":\"org.xnix.external.gui\",\"external_desktop_argument_count\":1,\"external_file_uri_arguments_accepted\":true,\"external_file_open_requested\":true,\"external_file_bridge_ready\":true,\"imported_artifact_digest_verified\":true,\"runtime_run_requested\":true,\"runtime_run_executed\":true,\"execution_started\":true,\"backend_process_started\":true,\"container_runtime_used\":true,\"container_network_mode\":\"none\",\"container_host_mount_count\":0,\"x_window_observed\":true,\"window_observed\":true,\"runtime_owned\":true,\"go_runtime_backed\":true,\"kde_policy_owner\":false,\"desktop_launch_enabled\":false,\"action_execution_enabled\":false,\"backend_details_exposed\":false,\"raw_import_record_path_exposed\":false,\"raw_external_app_handle_path_exposed\":false,\"raw_state_root_path_exposed\":false,\"raw_executable_path_exposed\":false,\"host_root_modified\":false,\"privileged_container_required\":false,\"host_networking_required\":false,\"docker_socket_mounted\":false,\"broad_host_mount_required\":false}'\n"
	if err := os.WriteFile(launcherSource, []byte(launcherBody), 0o755); err != nil {
		t.Fatalf("WriteFile launcher returned error: %v", err)
	}
	documentPath := filepath.Join(tempDir, "report.docx")
	if err := os.WriteFile(documentPath, []byte("external staged launcher file-open fixture"), 0o600); err != nil {
		t.Fatalf("WriteFile document returned error: %v", err)
	}
	documentURI := "file://" + filepath.ToSlash(documentPath)
	packetOutput := filepath.Join(tempDir, "sidecars", "desktop-launch-packet.json")
	outputPath := filepath.Join(tempDir, "import-stage-launch.json")
	var output bytes.Buffer
	if err := run([]string{
		"external-winapp-import-stage-and-launch",
		"--state-root", stateRoot,
		"--executable", executablePath,
		"--app-id", "org.xnix.external.gui",
		"--display-name", "External GUI",
		"--mode", "development",
		"--staging-root", stagingRoot,
		"--managed-launcher-bin", launcherSource,
		"--desktop-launch-packet-output", packetOutput,
		"--image", "local/wine-x-gui:test",
		"--platform", "linux/amd64",
		"--docker", filepath.Join(tempDir, "fake-docker"),
		"--timeout", "5s",
		"--output", outputPath,
		documentURI,
	}, &output); err != nil {
		t.Fatalf("external app import-stage-and-launch returned error: %v", err)
	}
	written, err := os.ReadFile(outputPath)
	if err != nil {
		t.Fatalf("ReadFile output returned error: %v", err)
	}
	if string(written) != output.String() {
		t.Fatalf("written output must match stdout\nstdout=%s\nwritten=%s", output.String(), string(written))
	}
	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal import-stage-and-launch output returned error: %v\n%s", err, output.String())
	}
	if payload["schema_version"] != "xnix.runtime.external_winapp_import_stage_launch.v1" ||
		payload["request_type"] != "external-winapp-import-stage-and-launch" ||
		payload["status"] != "passed" ||
		payload["application_id"] != "org.xnix.external.gui" ||
		payload["external_app_handle"] != "org.xnix.external.gui" ||
		payload["import_recorded"] != true ||
		payload["desktop_activation_staged"] != true ||
		payload["staged_launcher_invoked"] != true ||
		payload["staged_launcher_from_activation_root"] != true ||
		payload["managed_launcher_executable_staged"] != true ||
		payload["desktop_exec_uses_external_app_handle"] != true ||
		payload["external_app_desktop_handle_ready"] != true ||
		payload["desktop_launch_packet_requested"] != true ||
		payload["desktop_launch_packet_written"] != true ||
		payload["launcher_request_type"] != "windows-external-app-run" ||
		payload["launcher_status"] != "passed" ||
		payload["external_app_import_record_consumed"] != true ||
		payload["external_app_handle_consumed"] != true ||
		payload["external_desktop_argument_count"] != float64(1) ||
		payload["external_file_uri_arguments_accepted"] != true ||
		payload["external_file_bridge_ready"] != true ||
		payload["imported_artifact_digest_verified"] != true ||
		payload["runtime_launch_executed"] != true ||
		payload["window_observed"] != true ||
		payload["x_window_observed"] != true ||
		payload["launch_enabled"] != false ||
		payload["backend_launch_enabled"] != false ||
		payload["execution_started"] != true ||
		payload["host_root_modified"] != false ||
		payload["docker_socket_mounted"] != false ||
		payload["raw_launcher_path_exposed"] != false ||
		payload["raw_launcher_output_exposed"] != false {
		t.Fatalf("unexpected external import-stage-and-launch payload: %#v", payload)
	}
	recordPath, err := appidentity.ExternalWinAppImportRecordPathFromHandle(stateRoot, "org.xnix.external.gui")
	if err != nil {
		t.Fatalf("ExternalWinAppImportRecordPathFromHandle returned error: %v", err)
	}
	if _, err := os.Stat(recordPath); err != nil {
		t.Fatalf("one-shot import-stage-and-launch must persist the import record: %v", err)
	}
	stagedLauncher := filepath.Join(stagingRoot, "usr/local/bin/xnix-compat-launch")
	if _, err := os.Stat(stagedLauncher); err != nil {
		t.Fatalf("one-shot import-stage-and-launch must stage the managed launcher: %v", err)
	}
	if _, err := os.Stat(packetOutput); err != nil {
		t.Fatalf("one-shot import-stage-and-launch must allow the launcher to write the desktop launch packet: %v", err)
	}
	launcherInvocation, err := os.ReadFile(launcherLog)
	if err != nil {
		t.Fatalf("ReadFile launcher log returned error: %v", err)
	}
	for _, expected := range []string{
		"--external-app-handle org.xnix.external.gui --timeout 5s " + documentURI,
		"state=" + stateRoot,
		"image=local/wine-x-gui:test",
		"platform=linux/amd64",
		"docker=" + filepath.Join(tempDir, "fake-docker"),
		"packet=" + packetOutput,
	} {
		if !strings.Contains(string(launcherInvocation), expected) {
			t.Fatalf("staged launcher invocation missing %q:\n%s", expected, string(launcherInvocation))
		}
	}
	for _, forbidden := range []string{stateRoot, stagingRoot, executablePath, launcherSource, stagedLauncher, packetOutput, documentPath, documentURI, filepath.Join(tempDir, "fake-docker")} {
		if strings.Contains(output.String(), forbidden) {
			t.Fatalf("external import-stage-and-launch output exposed forbidden path %q: %s", forbidden, output.String())
		}
	}
}

func TestExternalWinAppImportAndRunCommandImportsAndRunsFileOpen(t *testing.T) {
	tempDir := t.TempDir()
	stateRoot := filepath.Join(tempDir, "state")
	executablePath := filepath.Join(tempDir, "ExternalGui.exe")
	if err := os.WriteFile(executablePath, []byte{'M', 'Z', 0x90, 0x00, 'x', 'n', 'i', 'x'}, 0o600); err != nil {
		t.Fatalf("WriteFile executable returned error: %v", err)
	}
	documentPath := filepath.Join(tempDir, "sample-document.txt")
	if err := os.WriteFile(documentPath, []byte("external app one-shot file-open fixture"), 0o600); err != nil {
		t.Fatalf("WriteFile document returned error: %v", err)
	}
	documentURI := "file://" + filepath.ToSlash(documentPath)
	dockerLog := filepath.Join(tempDir, "fake-docker.log")
	dockerPath := filepath.Join(tempDir, "fake-docker")
	dockerBody := "#!/bin/sh\n" +
		"printf '%s\\n' \"$@\" >> \"" + dockerLog + "\"\n" +
		"if test \"$1 $2\" = 'image inspect'; then printf 'linux/amd64\\n'; exit 0; fi\n" +
		"if test \"$1\" = 'create'; then printf 'fake-x-gui-container\\n'; exit 0; fi\n" +
		"if test \"$1\" = 'cp'; then exit 0; fi\n" +
		"if test \"$1 $2\" = 'start -a'; then " +
		"printf 'XNIX_X_GUI_XSERVER_STARTED=true\\n'\n" +
		"printf 'XNIX_X_GUI_WINE_BOOTSTRAP_ATTEMPTED=true\\n'\n" +
		"printf 'XNIX_X_GUI_FILE_ARGS_PASSED=1\\n'\n" +
		"printf 'XNIX_X_GUI_FILE_ARGS_WINEPATH_TRANSLATED=1\\n'\n" +
		"printf '0x700001 \"sample-document.txt - Notepad\": (\"notepad.exe\" \"notepad.exe\") 320x160+20+20 +20+20\\n'\n" +
		"printf 'XNIX_X_GUI_WINDOW_OBSERVED=true\\n'; exit 0; fi\n" +
		"if test \"$1\" = 'rm'; then exit 0; fi\n" +
		"exit 2\n"
	if err := os.WriteFile(dockerPath, []byte(dockerBody), 0o700); err != nil {
		t.Fatalf("WriteFile docker returned error: %v", err)
	}
	outputPath := filepath.Join(tempDir, "import-and-run.json")
	var output bytes.Buffer
	if err := run([]string{
		"external-winapp-import-and-run",
		"--state-root", stateRoot,
		"--executable", executablePath,
		"--app-id", "org.xnix.external.gui",
		"--display-name", "External GUI",
		"--image", "local/wine-x-gui:test",
		"--platform", "linux/amd64",
		"--docker", dockerPath,
		"--timeout", "5s",
		"--output", outputPath,
		documentURI,
	}, &output); err != nil {
		t.Fatalf("external app import-and-run returned error: %v", err)
	}
	written, err := os.ReadFile(outputPath)
	if err != nil {
		t.Fatalf("ReadFile output returned error: %v", err)
	}
	if string(written) != output.String() {
		t.Fatalf("written output must match stdout\nstdout=%s\nwritten=%s", output.String(), string(written))
	}
	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal import-and-run output returned error: %v\n%s", err, output.String())
	}
	if payload["schema_version"] != "xnix.runtime.external_winapp_run.v1" ||
		payload["request_type"] != "windows-external-app-run" ||
		payload["status"] != "passed" ||
		payload["application_id"] != "org.xnix.external.gui" ||
		payload["external_app_import_record_consumed"] != true ||
		payload["external_app_handle_consumed"] != true ||
		payload["external_app_handle"] != "org.xnix.external.gui" ||
		payload["external_desktop_argument_count"] != float64(1) ||
		payload["external_file_bridge_ready"] != true ||
		payload["external_file_bridge_winepath_translated"] != true ||
		payload["imported_artifact_digest_verified"] != true ||
		payload["runtime_run_executed"] != true ||
		payload["container_network_mode"] != "none" ||
		payload["container_host_mount_count"] != float64(0) ||
		payload["x_window_observed"] != true ||
		payload["window_observed"] != true ||
		payload["host_root_modified"] != false ||
		payload["docker_socket_mounted"] != false {
		t.Fatalf("unexpected external import-and-run payload: %#v", payload)
	}
	runtimePayload := payload["runtime_payload"].(map[string]any)
	if runtimePayload["window_match"] != "sample-document.txt" ||
		runtimePayload["file_bridge_copied_count"] != float64(1) ||
		runtimePayload["file_bridge_winepath_translated_count"] != float64(1) {
		t.Fatalf("unexpected import-and-run runtime payload: %#v", runtimePayload)
	}
	recordPath, err := appidentity.ExternalWinAppImportRecordPathFromHandle(stateRoot, "org.xnix.external.gui")
	if err != nil {
		t.Fatalf("ExternalWinAppImportRecordPathFromHandle returned error: %v", err)
	}
	if _, err := os.Stat(recordPath); err != nil {
		t.Fatalf("one-shot import-and-run must persist the import record: %v", err)
	}
	dockerInvocation, err := os.ReadFile(dockerLog)
	if err != nil {
		t.Fatalf("ReadFile docker log returned error: %v", err)
	}
	if !strings.Contains(string(dockerInvocation), "ExternalGui.exe") ||
		!strings.Contains(string(dockerInvocation), documentPath) ||
		!strings.Contains(string(dockerInvocation), "/file-1-sample-document.txt") ||
		!strings.Contains(string(dockerInvocation), "XNIX_WINDOW_MATCH=sample-document.txt") {
		t.Fatalf("fake Docker did not receive the one-shot import-and-run flow: %s", string(dockerInvocation))
	}
	if strings.Contains(output.String(), stateRoot) ||
		strings.Contains(output.String(), executablePath) ||
		strings.Contains(output.String(), dockerPath) ||
		strings.Contains(output.String(), documentPath) ||
		strings.Contains(output.String(), documentURI) ||
		strings.Contains(output.String(), "docker run") ||
		strings.Contains(output.String(), "/var/run/docker.sock") ||
		strings.Contains(output.String(), "--network host") ||
		strings.Contains(output.String(), "--privileged") {
		t.Fatalf("external import-and-run output exposed unsafe details: %s", output.String())
	}
}
