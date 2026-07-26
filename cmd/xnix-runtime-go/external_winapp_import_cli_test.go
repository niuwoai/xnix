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
