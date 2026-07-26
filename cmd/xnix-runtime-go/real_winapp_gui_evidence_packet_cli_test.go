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

	var compatibilityOutput bytes.Buffer
	if err := run([]string{"compatibility-center-preview", "--registry", "../../runtime/recipes/registry.json", "--known-app-evidence-file", outputPath}, &compatibilityOutput); err != nil {
		t.Fatalf("compatibility center consumption returned error: %v", err)
	}
	var compatibilityPayload map[string]any
	if err := json.Unmarshal(compatibilityOutput.Bytes(), &compatibilityPayload); err != nil {
		t.Fatalf("Unmarshal compatibility output returned error: %v", err)
	}
	if compatibilityPayload["known_app_smoke_evidence_count"] != float64(1) ||
		compatibilityPayload["known_app_smoke_passed_count"] != float64(1) ||
		compatibilityPayload["known_app_launch_authorization_required_count"] != float64(1) ||
		compatibilityPayload["backend_launch_enabled"] != false ||
		compatibilityPayload["backend_details_exposed"] != false ||
		compatibilityPayload["host_root_modified"] != false {
		t.Fatalf("unexpected compatibility payload from real GUI packet: %#v", compatibilityPayload)
	}

	var kdeOutput bytes.Buffer
	if err := run([]string{"kde-center-page-preview", "--registry", "../../runtime/recipes/registry.json", "--app", "org.xnix.sample.notepad", "--decision", "approved", "--known-app-evidence-file", outputPath}, &kdeOutput); err != nil {
		t.Fatalf("KDE center consumption returned error: %v", err)
	}
	var kdePayload map[string]any
	if err := json.Unmarshal(kdeOutput.Bytes(), &kdePayload); err != nil {
		t.Fatalf("Unmarshal KDE output returned error: %v", err)
	}
	if !strings.Contains(kdePayload["source"].(string), "known-app-gui-smoke-evidence") ||
		kdePayload["known_app_gui_evidence_count"] != float64(1) ||
		kdePayload["launch_enabled"] != false ||
		kdePayload["backend_details_exposed"] != false ||
		kdePayload["host_root_modified"] != false {
		t.Fatalf("unexpected KDE payload from real GUI packet: %#v", kdePayload)
	}
	cards := kdePayload["known_app_gui_evidence_cards"].([]any)
	if len(cards) != 1 {
		t.Fatalf("unexpected KDE real GUI packet cards: %#v", cards)
	}
	card := cards[0].(map[string]any)
	if card["app_id"] != "org.xnix.sample.notepad" ||
		card["evidence_source"] != "winapp-smoke-container-x-gui" ||
		card["recipe_backed"] != true ||
		card["recipe_app_id"] != "org.xnix.sample.notepad" ||
		card["compatibility_state"] != "real-gui-container-wine-verified" ||
		card["center_card_state"] != "validated-real-gui-container-run" ||
		card["execution_evidence_recorded"] != true ||
		card["runtime_dispatch_verified"] != true ||
		card["desktop_launch_enabled"] != false ||
		card["backend_launch_enabled"] != false ||
		card["backend_details_exposed"] != false ||
		card["host_root_modified"] != false {
		t.Fatalf("unexpected KDE card from real GUI packet: %#v", card)
	}
}

func TestRealWinAppGUIEvidencePacketPreviewCommandConsumesRawContainerRuntimePayload(t *testing.T) {
	tempDir := t.TempDir()
	reportPath := filepath.Join(tempDir, "staged-launcher-delegated-notepad.json")
	if err := os.WriteFile(reportPath, []byte(rawContainerXGUIRuntimePayloadCLIFixture()), 0o600); err != nil {
		t.Fatalf("WriteFile report returned error: %v", err)
	}

	var output bytes.Buffer
	if err := run([]string{
		"real-winapp-gui-evidence-packet-preview",
		"--gui-smoke-report", reportPath,
		"--app-id", "org.xnix.sample.notepad",
		"--display-name", "Sample Notepad",
		"--app-version", "0.2.640-test",
	}, &output); err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["schema_version"] != "xnix.runtime.real_winapp_gui_evidence_packet.v1" ||
		payload["request_type"] != "real-winapp-gui-evidence-packet-preview" ||
		payload["version"] != "0.2.640-test" ||
		payload["app_id"] != "org.xnix.sample.notepad" ||
		payload["display_name"] != "Sample Notepad" ||
		payload["gui_app_name"] != "notepad.exe" ||
		payload["evidence_source"] != "winapp-smoke-container-x-gui" ||
		payload["recipe_backed"] != true ||
		payload["recipe_app_id"] != "org.xnix.sample.notepad" ||
		payload["known_app_gui_evidence_verified_count"] != float64(1) ||
		payload["container_runtime_used"] != true ||
		payload["container_network_mode"] != "none" ||
		payload["container_host_mount_count"] != float64(0) ||
		payload["x_window_observed"] != true ||
		payload["desktop_launch_enabled"] != false ||
		payload["backend_launch_enabled"] != false ||
		payload["backend_details_exposed"] != false ||
		payload["host_root_modified"] != false ||
		payload["docker_socket_mounted"] != false {
		t.Fatalf("unexpected real GUI packet from raw Runtime payload: %#v", payload)
	}
	if strings.Contains(output.String(), reportPath) || strings.Contains(output.String(), "docker run") || strings.Contains(output.String(), "/var/run/docker.sock") {
		t.Fatalf("raw Runtime payload packet exposed unsafe details: %s", output.String())
	}
	evidence := payload["known_app_smoke_evidence"].(map[string]any)
	if evidence["staged_launcher_verified"] != true ||
		evidence["launch_authorization_receipt_id"] != "known-app-launch-authorization-org.xnix.sample.notepad-0.2.640-test" ||
		evidence["launch_gate_state"] != "controlled-dispatch-ready" ||
		evidence["launch_gate_consumed"] != true ||
		evidence["launch_gate_receipt_accepted"] != true ||
		evidence["launch_gate_guest_boundary_accepted"] != true ||
		evidence["controlled_dispatch_ready"] != true ||
		evidence["controlled_execution_session_id"] != "known-app-controlled-execution-session-org.xnix.sample.notepad-0.2.640-test" ||
		evidence["launcher_session_gate_consumed"] != true ||
		evidence["launcher_session_digest_verified"] != true ||
		evidence["launcher_session_relative_path"] != "execution-ledger/sessions/known-app-controlled-execution-session-org.xnix.sample.notepad-0.2.640-test.json" ||
		evidence["launcher_session_runtime_owner_consumable"] != true ||
		evidence["launcher_session_kde_read_model_consumable"] != true ||
		evidence["post_review_dispatch_consumed"] != true ||
		evidence["post_review_dispatch_state"] != "created-after-session-gated-review" ||
		evidence["session_gated_review_receipt_id"] != "known-app-session-gated-launch-review-org.xnix.sample.notepad-0.2.640-test-known-app-controlled-execution-session-org.xnix.sample.notepad-0.2.640-test" {
		t.Fatalf("raw Runtime payload must surface staged launcher session evidence: %#v", evidence)
	}
}

func TestRealWinAppGUIEvidencePacketPreviewCommandConsumesRawExternalExecutableRuntimePayload(t *testing.T) {
	tempDir := t.TempDir()
	reportPath := filepath.Join(tempDir, "external-notepad-container-gui.json")
	outputPath := filepath.Join(tempDir, "packet", "external-notepad-real-gui-evidence.json")
	if err := os.WriteFile(reportPath, []byte(rawExternalExecutableContainerXGUIRuntimePayloadCLIFixture()), 0o600); err != nil {
		t.Fatalf("WriteFile report returned error: %v", err)
	}

	var output bytes.Buffer
	if err := run([]string{
		"real-winapp-gui-evidence-packet-preview",
		"--gui-smoke-report", reportPath,
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
		payload["app_id"] != "org.xnix.external.notepad-file" ||
		payload["display_name"] != "External Notepad File" ||
		payload["app_version"] != "0.2.640-test" ||
		payload["recipe_backed"] != false ||
		payload["executable_name"] != "notepad.exe" ||
		payload["local_executable_copied"] != true ||
		payload["known_app_gui_evidence_verified_count"] != float64(1) ||
		payload["container_runtime_used"] != true ||
		payload["container_network_mode"] != "none" ||
		payload["container_host_mount_count"] != float64(0) ||
		payload["x_window_observed"] != true ||
		payload["desktop_launch_enabled"] != false ||
		payload["backend_launch_enabled"] != false ||
		payload["backend_details_exposed"] != false ||
		payload["host_root_modified"] != false ||
		payload["docker_socket_mounted"] != false {
		t.Fatalf("unexpected external executable real GUI packet: %#v", payload)
	}
	if strings.Contains(output.String(), reportPath) ||
		strings.Contains(output.String(), outputPath) ||
		strings.Contains(output.String(), "docker run") ||
		strings.Contains(output.String(), "/var/run/docker.sock") {
		t.Fatalf("external executable real GUI packet exposed unsafe details: %s", output.String())
	}

	var kdeOutput bytes.Buffer
	if err := run([]string{"kde-center-page-preview", "--registry", "../../runtime/recipes/registry.json", "--app", "org.xnix.sample.notepad", "--decision", "approved", "--known-app-evidence-file", outputPath}, &kdeOutput); err != nil {
		t.Fatalf("KDE center consumption returned error: %v", err)
	}
	var kdePayload map[string]any
	if err := json.Unmarshal(kdeOutput.Bytes(), &kdePayload); err != nil {
		t.Fatalf("Unmarshal KDE output returned error: %v", err)
	}
	if kdePayload["known_app_gui_evidence_count"] != float64(1) ||
		kdePayload["launch_enabled"] != false ||
		kdePayload["backend_details_exposed"] != false ||
		kdePayload["host_root_modified"] != false {
		t.Fatalf("unexpected KDE payload from external executable packet: %#v", kdePayload)
	}
	cards := kdePayload["known_app_gui_evidence_cards"].([]any)
	if len(cards) != 1 {
		t.Fatalf("unexpected KDE external executable packet cards: %#v", cards)
	}
	card := cards[0].(map[string]any)
	if card["app_id"] != "org.xnix.external.notepad-file" ||
		card["display_name"] != "External Notepad File" ||
		card["evidence_source"] != "winapp-smoke-container-x-gui" ||
		card["recipe_backed"] != false ||
		card["compatibility_state"] != "real-gui-container-wine-verified" ||
		card["center_card_state"] != "validated-real-gui-container-run" ||
		card["execution_evidence_recorded"] != true ||
		card["runtime_dispatch_verified"] != true ||
		card["desktop_launch_enabled"] != false ||
		card["backend_launch_enabled"] != false ||
		card["backend_details_exposed"] != false ||
		card["host_root_modified"] != false {
		t.Fatalf("unexpected KDE card from external executable packet: %#v", card)
	}
}

func rawContainerXGUIRuntimePayloadCLIFixture() string {
	return `{
  "schema_version": "xnix.runtime.windows_app_container_x_gui_smoke.v1",
  "request_type": "windows-app-container-x-gui-smoke",
  "status": "passed",
  "application_id": "org.xnix.sample.notepad",
  "display_name": "Sample Notepad",
  "app_version": "0.2.640-test",
  "recipe_backed": true,
  "application_name": "notepad.exe",
  "window_match": "notepad.exe",
  "container_image": "xnix-wine-smoke:local",
  "container_platform": "linux/arm64",
  "network_mode": "none",
  "x_server_started": true,
  "wine_bootstrap_attempted": true,
  "image_available": true,
  "x_window_observed": true,
  "window_evidence_summary": "0x600001 \"Untitled - Notepad\": (\"notepad.exe\" \"notepad.exe\")",
  "evidence_source": "winapp-smoke-container-x-gui",
  "dispatch_started": true,
  "execution_started": true,
  "smoke_passed": true,
  "runtime_owned_dispatch": true,
  "session_gated_controlled_dispatch_consumed": true,
  "session_gated_controlled_dispatch_state": "created-after-session-gated-review",
  "session_gated_review_receipt_id": "known-app-session-gated-launch-review-org.xnix.sample.notepad-0.2.640-test-known-app-controlled-execution-session-org.xnix.sample.notepad-0.2.640-test",
  "launch_authorization_receipt_id": "known-app-launch-authorization-org.xnix.sample.notepad-0.2.640-test",
  "controlled_execution_session_consumed": true,
  "controlled_execution_session_id": "known-app-controlled-execution-session-org.xnix.sample.notepad-0.2.640-test",
  "controlled_session_digest_verified": true,
  "controlled_session_relative_path": "execution-ledger/sessions/known-app-controlled-execution-session-org.xnix.sample.notepad-0.2.640-test.json",
  "runtime_owner_consumable_session": true,
  "kde_read_model_consumable_session": true,
  "controlled_session_live_state_observed": true,
  "controlled_session_registered": true,
  "controlled_session_window_observed": true,
  "controlled_session_host_root_modified": false,
  "controlled_session_container_process_start": true,
  "raw_command_exposed": false,
  "backend_details_exposed": false,
  "host_root_modified": false,
  "privileged_container_required": false,
  "host_networking_required": false,
  "docker_socket_mounted": false,
  "broad_host_mount_required": false,
  "host_mount_count": 0
}`
}

func rawExternalExecutableContainerXGUIRuntimePayloadCLIFixture() string {
	return `{
  "schema_version": "xnix.runtime.windows_app_container_x_gui_smoke.v1",
  "request_type": "windows-app-container-x-gui-smoke",
  "status": "passed",
  "application_id": "org.xnix.external.notepad-file",
  "display_name": "External Notepad File",
  "app_version": "0.2.640-test",
  "recipe_backed": false,
  "executable_name": "notepad.exe",
  "local_executable_copied": true,
  "application_name": "/notepad.exe",
  "window_match": "notepad.exe",
  "container_image": "xnix-wine-smoke:local",
  "container_platform": "linux/arm64",
  "network_mode": "none",
  "x_server_started": true,
  "wine_bootstrap_attempted": true,
  "image_available": true,
  "x_window_observed": true,
  "window_evidence_summary": "0xa00001 \"Untitled - Notepad\": (\"notepad.exe\" \"notepad.exe\") 721x519+4+23 +4+23",
  "host_root_modified": false,
  "privileged_container_required": false,
  "host_networking_required": false,
  "docker_socket_mounted": false,
  "broad_host_mount_required": false,
  "host_mount_count": 0
}`
}
