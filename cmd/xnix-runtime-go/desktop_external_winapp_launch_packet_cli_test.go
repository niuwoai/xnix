package main

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestDesktopExternalWinAppLaunchPacketPreviewCommandConsumesStagedHandleRun(t *testing.T) {
	tempDir := t.TempDir()
	stateRoot := filepath.Join(tempDir, "state")
	stageRoot := filepath.Join(tempDir, "stage")
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

	var stageOutput bytes.Buffer
	if err := run([]string{
		"desktop-activation-stage",
		"--external-app-import-record", recordPath,
		"--mode", "development",
		"--staging-root", stageRoot,
	}, &stageOutput); err != nil {
		t.Fatalf("desktop activation stage returned error: %v", err)
	}

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
	runRecordPath := filepath.Join(tempDir, "handle-run.json")
	var runOutput bytes.Buffer
	if err := run([]string{
		"windows-external-app-run",
		"--state-root", stateRoot,
		"--external-app-handle", "org.xnix.external.gui",
		"--window-match", "External GUI",
		"--image", "local/wine-x-gui:test",
		"--platform", "linux/amd64",
		"--docker", dockerPath,
		"--timeout", "5s",
		"--output", runRecordPath,
	}, &runOutput); err != nil {
		t.Fatalf("external app handle run returned error: %v", err)
	}

	outputPath := filepath.Join(tempDir, "desktop-launch-packet.json")
	var output bytes.Buffer
	if err := run([]string{
		"desktop-external-winapp-launch-packet-preview",
		"--external-app-import-record", recordPath,
		"--activation-root", stageRoot,
		"--run-record", runRecordPath,
		"--mode", "development",
		"--output", outputPath,
	}, &output); err != nil {
		t.Fatalf("desktop external launch packet returned error: %v", err)
	}
	written, err := os.ReadFile(outputPath)
	if err != nil {
		t.Fatalf("ReadFile packet output returned error: %v", err)
	}
	if string(written) != output.String() {
		t.Fatalf("written packet must match stdout\nstdout=%s\nwritten=%s", output.String(), string(written))
	}
	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal packet returned error: %v\n%s", err, output.String())
	}
	if payload["schema_version"] != "xnix.runtime.desktop_external_winapp_launch_packet.v1" ||
		payload["request_type"] != "desktop-external-winapp-launch-packet-preview" ||
		payload["packet_type"] != "kde-desktop-external-winapp-launch-evidence" ||
		payload["status"] != "passed" ||
		payload["application_id"] != "org.xnix.external.gui" ||
		payload["display_name"] != "External GUI" ||
		payload["external_app_handle"] != "org.xnix.external.gui" ||
		payload["activation_receipt_backed"] != true ||
		payload["activation_receipt_safe_for_kde"] != true ||
		payload["desktop_exec_uses_external_app_handle"] != true ||
		payload["external_app_desktop_handle_ready"] != true ||
		payload["desktop_exec_uses_raw_import_record"] != false ||
		payload["desktop_exec_uses_state_root"] != false ||
		payload["run_record_consumed"] != true ||
		payload["external_app_run_record_consumed"] != true ||
		payload["external_app_import_record_consumed"] != true ||
		payload["external_app_handle_consumed"] != true ||
		payload["imported_artifact_digest_verified"] != true ||
		payload["runtime_launch_executed"] != true ||
		payload["window_observed"] != true ||
		payload["x_window_observed"] != true ||
		payload["container_network_mode"] != "none" ||
		payload["container_host_mount_count"] != float64(0) ||
		payload["runtime_launch_authority"] != true ||
		payload["kde_launch_authority"] != false ||
		payload["desktop_launch_packet_ready"] != true ||
		payload["safe_for_kde"] != true ||
		payload["backend_details_exposed"] != false ||
		payload["raw_import_record_path_exposed"] != false ||
		payload["raw_state_root_path_exposed"] != false ||
		payload["raw_executable_path_exposed"] != false ||
		payload["host_root_modified"] != false ||
		payload["docker_socket_mounted"] != false ||
		payload["broad_host_mount_required"] != false {
		t.Fatalf("unexpected desktop external launch packet: %#v", payload)
	}
	if reasons := payload["unsafe_reason_ids"].([]any); len(reasons) != 0 {
		t.Fatalf("unsafe_reason_ids = %#v, want empty", reasons)
	}
	for _, forbidden := range []string{".exe", recordPath, stateRoot, stageRoot, runRecordPath, executablePath, dockerPath, "docker run", "/var/run/docker.sock", "--network host", "--privileged"} {
		if strings.Contains(strings.ToLower(output.String()), strings.ToLower(forbidden)) {
			t.Fatalf("desktop external launch packet exposed forbidden term %q: %s", forbidden, output.String())
		}
	}
}
