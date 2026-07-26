package main

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"
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
