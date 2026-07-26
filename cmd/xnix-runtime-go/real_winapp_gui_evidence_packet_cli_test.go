package main

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestRealWinAppGUIEvidencePacketPreviewCommandConsumesContainerNotepadReport(t *testing.T) {
	tempDir := t.TempDir()
	reportPath := filepath.Join(tempDir, "winapp-container-x-gui-notepad.json")
	outputPath := filepath.Join(tempDir, "packet", "notepad-real-gui-evidence.json")
	if err := os.WriteFile(reportPath, []byte(containerXGUISmokeEvidenceCLIFixture()), 0o600); err != nil {
		t.Fatalf("WriteFile report returned error: %v", err)
	}

	var output bytes.Buffer
	if err := run([]string{
		"real-winapp-gui-evidence-packet-preview",
		"--gui-smoke-report", reportPath,
		"--app-id", "org.xnix.sample.notepad",
		"--display-name", "Sample Notepad",
		"--app-version", "container-local",
		"--output", outputPath,
	}, &output); err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	written, err := os.ReadFile(outputPath)
	if err != nil {
		t.Fatalf("ReadFile output returned error: %v", err)
	}
	if string(written) != output.String() {
		t.Fatalf("written packet must match stdout\nstdout=%s\nwritten=%s", output.String(), string(written))
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["schema_version"] != "xnix.runtime.real_winapp_gui_evidence_packet.v1" ||
		payload["request_type"] != "real-winapp-gui-evidence-packet-preview" ||
		payload["packet_type"] != "real-windows-app-gui-evidence" ||
		payload["runtime_method"] != "PreviewRealWinAppGUIEvidencePacket" ||
		payload["read_method"] != "GetRealWinAppGUIEvidencePacket" ||
		payload["report_status"] != "passed" ||
		payload["report_consumed"] != true ||
		payload["report_path_exposed"] != false ||
		payload["app_id"] != "org.xnix.sample.notepad" ||
		payload["display_name"] != "Sample Notepad" ||
		payload["gui_app_name"] != "notepad.exe" ||
		payload["evidence_source"] != "winapp-smoke-container-x-gui" ||
		payload["recipe_backed"] != true ||
		payload["recipe_app_id"] != "org.xnix.sample.notepad" ||
		payload["compatibility_state"] != "real-gui-container-wine-verified" ||
		payload["center_card_state"] != "validated-real-gui-container-run" ||
		payload["known_app_gui_evidence_count"] != float64(1) ||
		payload["known_app_gui_evidence_verified_count"] != float64(1) ||
		payload["x_window_observed"] != true ||
		payload["compatibility_center_projection_ready"] != true ||
		payload["kde_center_projection_ready"] != true ||
		payload["container_runtime_used"] != true ||
		payload["container_network_mode"] != "none" ||
		payload["container_host_mount_count"] != float64(0) ||
		payload["desktop_launch_enabled"] != false ||
		payload["backend_launch_enabled"] != false ||
		payload["backend_details_exposed"] != false ||
		payload["host_root_modified"] != false ||
		payload["privileged_container_required"] != false ||
		payload["host_networking_required"] != false ||
		payload["docker_socket_mounted"] != false ||
		payload["broad_host_mount_required"] != false {
		t.Fatalf("unexpected real GUI evidence packet payload: %#v", payload)
	}
	if strings.Contains(output.String(), reportPath) || strings.Contains(output.String(), "docker run") || strings.Contains(output.String(), "/var/run/docker.sock") {
		t.Fatalf("real GUI evidence packet exposed unsafe details: %s", output.String())
	}
	evidence := payload["known_app_smoke_evidence"].(map[string]any)
	if evidence["evidence_kind"] != "known-application-gui-smoke" ||
		evidence["evidence_source"] != "winapp-smoke-container-x-gui" ||
		evidence["execution_evidence_recorded"] != true ||
		evidence["runtime_dispatch_verified"] != true ||
		evidence["desktop_launch_enabled"] != false ||
		evidence["backend_launch_enabled"] != false ||
		evidence["backend_details_exposed"] != false ||
		evidence["host_root_modified"] != false {
		t.Fatalf("unexpected nested evidence item: %#v", evidence)
	}
}
