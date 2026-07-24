package main

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestKnownAppMatrixEvidencePreviewCommandConsumesAggregateReport(t *testing.T) {
	reportPath := filepath.Join(t.TempDir(), "known-run-matrix.json")
	if err := os.WriteFile(reportPath, []byte(knownAppMatrixEvidenceCLIFixture()), 0o600); err != nil {
		t.Fatalf("WriteFile report returned error: %v", err)
	}

	var output bytes.Buffer
	if err := run([]string{"known-app-matrix-evidence-preview", "--matrix-report", reportPath}, &output); err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["schema_version"] != "xnix.runtime.known_app_matrix_evidence_preview.v1" ||
		payload["request_type"] != "known-app-matrix-evidence-preview" ||
		payload["matrix_status"] != "passed" ||
		payload["matrix_report_consumed"] != true ||
		payload["matrix_report_path_exposed"] != false ||
		payload["matrix_report_output_written"] != true ||
		payload["app_count"] != float64(2) ||
		payload["passed_count"] != float64(2) ||
		payload["failed_count"] != float64(0) ||
		payload["qemu_executed_count"] != float64(2) ||
		payload["wine_executed_count"] != float64(2) ||
		payload["checksum_verified_count"] != float64(2) ||
		payload["raw_output_redacted_count"] != float64(2) ||
		payload["compatibility_center_projection_ready"] != true ||
		payload["kde_center_projection_ready"] != true ||
		payload["runtime_owned"] != true ||
		payload["go_runtime_backed"] != true ||
		payload["kde_policy_owner"] != false ||
		payload["desktop_launch_enabled"] != false ||
		payload["backend_launch_enabled"] != false ||
		payload["backend_details_exposed"] != false ||
		payload["raw_output_exposed"] != false ||
		payload["remote_path_exposed"] != false ||
		payload["host_root_modified"] != false {
		t.Fatalf("unexpected matrix evidence payload: %#v", payload)
	}
	if strings.Contains(output.String(), "/home/xnix-run-materials") {
		t.Fatalf("matrix evidence preview exposed remote path: %s", output.String())
	}
	apps := payload["apps"].([]any)
	if len(apps) != 2 {
		t.Fatalf("unexpected apps: %#v", apps)
	}
	busybox := apps[1].(map[string]any)
	if busybox["app_id"] != "busybox-w32" ||
		busybox["display_name"] != "BusyBox-w32 standalone console executable" ||
		busybox["smoke_status"] != "passed" ||
		busybox["compatibility_state"] != "real-qemu-wine-verified" ||
		busybox["marker_observed"] != true ||
		busybox["checksum_verified"] != true ||
		busybox["qemu_executed"] != true ||
		busybox["wine_executed"] != true ||
		busybox["remote_path_exposed"] != false {
		t.Fatalf("unexpected BusyBox app evidence: %#v", busybox)
	}
}

func TestKnownAppMatrixEvidencePreviewCommandRequiresReport(t *testing.T) {
	var output bytes.Buffer
	err := run([]string{"known-app-matrix-evidence-preview"}, &output)
	if err == nil || !strings.Contains(err.Error(), "requires --matrix-report") {
		t.Fatalf("expected missing report error, got %v", err)
	}
}

func TestCompatibilityCenterPreviewCommandConsumesKnownAppMatrixReport(t *testing.T) {
	registryPath, _ := writeTestRepairGroupRegistry(t)
	reportPath := filepath.Join(t.TempDir(), "known-run-matrix.json")
	if err := os.WriteFile(reportPath, []byte(knownAppMatrixEvidenceCLIFixture()), 0o600); err != nil {
		t.Fatalf("WriteFile report returned error: %v", err)
	}

	var output bytes.Buffer
	if err := run([]string{"compatibility-center-preview", "--registry", registryPath, "--known-app-matrix-report", reportPath}, &output); err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["known_app_smoke_evidence_count"] != float64(2) ||
		payload["known_app_smoke_passed_count"] != float64(2) ||
		payload["known_app_staged_launcher_passed_count"] != float64(0) ||
		payload["known_app_launch_authorization_required_count"] != float64(2) ||
		payload["known_app_launch_authorization_recorded_count"] != float64(0) ||
		payload["known_app_launch_gate_consumed_count"] != float64(0) ||
		payload["known_app_controlled_dispatch_ready_count"] != float64(0) ||
		payload["backend_launch_enabled"] != false ||
		payload["host_root_modified"] != false ||
		payload["backend_details_exposed"] != false {
		t.Fatalf("unexpected compatibility center matrix payload: %#v", payload)
	}
	if strings.Contains(output.String(), "/home/xnix-run-materials") {
		t.Fatalf("compatibility center matrix projection exposed remote path: %s", output.String())
	}
	evidenceItems := payload["known_app_smoke_evidence"].([]any)
	if len(evidenceItems) != 2 {
		t.Fatalf("unexpected evidence items: %#v", evidenceItems)
	}
	first := evidenceItems[0].(map[string]any)
	if first["evidence_kind"] != "known-application-matrix-smoke" ||
		first["evidence_source"] != "remote-known-winapp-matrix-smoke" ||
		first["compatibility_state"] != "real-qemu-wine-verified" ||
		first["center_card_state"] != "validated-real-runtime-run" ||
		first["primary_action_id"] != "review-known-app-matrix-evidence" ||
		first["primary_action_kind"] != "review" ||
		first["marker_observed"] != true ||
		first["checksum_verified"] != true ||
		first["runtime_dispatch_verified"] != true ||
		first["desktop_launch_enabled"] != false ||
		first["backend_launch_enabled"] != false ||
		first["backend_details_exposed"] != false ||
		first["raw_artifact_path_exposed"] != false {
		t.Fatalf("unexpected matrix evidence item: %#v", first)
	}
}

func TestKDECenterPagePreviewCommandConsumesKnownAppMatrixReport(t *testing.T) {
	registryPath, app := writeTestRepairGroupRegistry(t)
	reportPath := filepath.Join(t.TempDir(), "known-run-matrix.json")
	if err := os.WriteFile(reportPath, []byte(knownAppMatrixEvidenceCLIFixture()), 0o600); err != nil {
		t.Fatalf("WriteFile report returned error: %v", err)
	}

	var output bytes.Buffer
	if err := run([]string{"kde-center-page-preview", "--registry", registryPath, "--app", app, "--decision", "approved", "--known-app-matrix-report", reportPath}, &output); err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if !strings.Contains(payload["source"].(string), "known-app-matrix-evidence") ||
		payload["known_app_session_gate_evidence_count"] != float64(0) ||
		payload["known_app_launcher_session_gate_consumed_count"] != float64(0) ||
		payload["known_app_post_review_dispatch_consumed_count"] != float64(0) ||
		payload["known_app_matrix_evidence_count"] != float64(2) ||
		payload["launch_enabled"] != false ||
		payload["execution_started"] != false ||
		payload["backend_process_started"] != false ||
		payload["request_objects_created"] != false ||
		payload["host_root_modified"] != false ||
		payload["backend_details_exposed"] != false {
		t.Fatalf("unexpected KDE center matrix payload: %#v", payload)
	}
	if strings.Contains(output.String(), "/home/xnix-run-materials") {
		t.Fatalf("KDE center matrix projection exposed remote path: %s", output.String())
	}
	cards := payload["known_app_matrix_evidence_cards"].([]any)
	if len(cards) != 2 {
		t.Fatalf("unexpected matrix cards: %#v", cards)
	}
	card := cards[1].(map[string]any)
	if card["app_id"] != "busybox-w32" ||
		card["evidence_kind"] != "known-application-matrix-smoke" ||
		card["evidence_source"] != "remote-known-winapp-matrix-smoke" ||
		card["smoke_status"] != "passed" ||
		card["compatibility_state"] != "real-qemu-wine-verified" ||
		card["center_card_state"] != "validated-real-runtime-run" ||
		card["primary_action_id"] != "review-known-app-matrix-evidence" ||
		card["primary_action_kind"] != "review" ||
		card["marker_observed"] != true ||
		card["checksum_verified"] != true ||
		card["execution_evidence_recorded"] != true ||
		card["runtime_dispatch_verified"] != true ||
		card["desktop_launch_enabled"] != false ||
		card["backend_launch_enabled"] != false ||
		card["host_root_modified"] != false ||
		card["backend_details_exposed"] != false ||
		card["raw_artifact_path_exposed"] != false {
		t.Fatalf("unexpected KDE matrix card: %#v", card)
	}
}

func TestGUISmokeEvidencePreviewCommandConsumesMessageBoxReport(t *testing.T) {
	reportPath := filepath.Join(t.TempDir(), "wine-gui-messagebox.json")
	if err := os.WriteFile(reportPath, []byte(guiSmokeEvidenceCLIFixture()), 0o600); err != nil {
		t.Fatalf("WriteFile report returned error: %v", err)
	}

	var output bytes.Buffer
	if err := run([]string{
		"gui-smoke-evidence-preview",
		"--gui-smoke-report", reportPath,
		"--app-id", "org.xnix.fixture.messagebox",
		"--display-name", "Xnix MessageBox",
		"--app-version", "fixture-version",
	}, &output); err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["schema_version"] != "xnix.runtime.gui_smoke_evidence_preview.v1" ||
		payload["request_type"] != "gui-smoke-evidence-preview" ||
		payload["report_status"] != "passed" ||
		payload["report_consumed"] != true ||
		payload["report_path_exposed"] != false ||
		payload["app_id"] != "org.xnix.fixture.messagebox" ||
		payload["display_name"] != "Xnix MessageBox" ||
		payload["gui_app_name"] != "xnix-messagebox-smoke.exe" ||
		payload["local_gui_executable_configured"] != true ||
		payload["executable_copied"] != true ||
		payload["x_window_observed"] != true ||
		payload["x_window_child_count"] != float64(14) ||
		payload["compatibility_center_projection_ready"] != true ||
		payload["kde_center_projection_ready"] != true ||
		payload["desktop_launch_enabled"] != false ||
		payload["backend_launch_enabled"] != false ||
		payload["backend_details_exposed"] != false ||
		payload["host_root_modified"] != false {
		t.Fatalf("unexpected GUI smoke evidence payload: %#v", payload)
	}
	if strings.Contains(output.String(), "/home/xnix-run-materials") || strings.Contains(output.String(), reportPath) {
		t.Fatalf("GUI smoke evidence preview exposed raw paths: %s", output.String())
	}
	evidence := payload["known_app_smoke_evidence"].(map[string]any)
	if evidence["evidence_kind"] != "known-application-gui-smoke" ||
		evidence["evidence_source"] != "wine-guest-gui-smoke" ||
		evidence["compatibility_state"] != "real-gui-qemu-wine-verified" ||
		evidence["center_card_state"] != "validated-real-gui-runtime-run" ||
		evidence["marker_observed"] != false ||
		evidence["checksum_verified"] != false ||
		evidence["execution_evidence_recorded"] != true ||
		evidence["runtime_dispatch_verified"] != true ||
		evidence["desktop_launch_enabled"] != false ||
		evidence["backend_details_exposed"] != false {
		t.Fatalf("unexpected GUI known app evidence: %#v", evidence)
	}
}

func TestGUISmokeEvidencePreviewCommandWritesProjectedEvidenceOutput(t *testing.T) {
	tempDir := t.TempDir()
	reportPath := filepath.Join(tempDir, "wine-gui-messagebox.json")
	outputPath := filepath.Join(tempDir, "evidence", "wine-gui-messagebox-evidence.json")
	if err := os.WriteFile(reportPath, []byte(guiSmokeEvidenceCLIFixture()), 0o600); err != nil {
		t.Fatalf("WriteFile report returned error: %v", err)
	}

	var output bytes.Buffer
	if err := run([]string{
		"gui-smoke-evidence-preview",
		"--gui-smoke-report", reportPath,
		"--app-id", "org.xnix.fixture.messagebox",
		"--display-name", "Xnix MessageBox",
		"--app-version", "fixture-version",
		"--output", outputPath,
	}, &output); err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	written, err := os.ReadFile(outputPath)
	if err != nil {
		t.Fatalf("ReadFile output returned error: %v", err)
	}
	if string(written) != output.String() {
		t.Fatalf("written GUI smoke evidence must match stdout\nstdout=%s\nwritten=%s", output.String(), string(written))
	}
	if strings.Contains(string(written), "/home/xnix-run-materials") || strings.Contains(string(written), reportPath) {
		t.Fatalf("written GUI smoke evidence exposed raw paths: %s", string(written))
	}

	registryPath, app := writeTestRepairGroupRegistry(t)
	var compatibilityOutput bytes.Buffer
	if err := run([]string{"compatibility-center-preview", "--registry", registryPath, "--known-app-evidence-file", outputPath}, &compatibilityOutput); err != nil {
		t.Fatalf("compatibility center consumption returned error: %v", err)
	}
	var compatibilityPayload map[string]any
	if err := json.Unmarshal(compatibilityOutput.Bytes(), &compatibilityPayload); err != nil {
		t.Fatalf("Unmarshal compatibility output returned error: %v", err)
	}
	if compatibilityPayload["known_app_smoke_evidence_count"] != float64(1) ||
		compatibilityPayload["known_app_smoke_passed_count"] != float64(1) ||
		compatibilityPayload["backend_details_exposed"] != false ||
		compatibilityPayload["host_root_modified"] != false {
		t.Fatalf("unexpected compatibility payload from GUI evidence file: %#v", compatibilityPayload)
	}

	var kdeOutput bytes.Buffer
	if err := run([]string{"kde-center-page-preview", "--registry", registryPath, "--app", app, "--decision", "approved", "--known-app-evidence-file", outputPath}, &kdeOutput); err != nil {
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
		t.Fatalf("unexpected KDE payload from GUI evidence file: %#v", kdePayload)
	}
}

func TestCompatibilityCenterPreviewCommandConsumesGUISmokeReport(t *testing.T) {
	registryPath, _ := writeTestRepairGroupRegistry(t)
	reportPath := filepath.Join(t.TempDir(), "wine-gui-messagebox.json")
	if err := os.WriteFile(reportPath, []byte(guiSmokeEvidenceCLIFixture()), 0o600); err != nil {
		t.Fatalf("WriteFile report returned error: %v", err)
	}

	var output bytes.Buffer
	if err := run([]string{"compatibility-center-preview", "--registry", registryPath, "--known-app-gui-smoke-report", reportPath}, &output); err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["known_app_smoke_evidence_count"] != float64(1) ||
		payload["known_app_smoke_passed_count"] != float64(1) ||
		payload["known_app_staged_launcher_passed_count"] != float64(0) ||
		payload["known_app_launch_authorization_required_count"] != float64(1) ||
		payload["backend_launch_enabled"] != false ||
		payload["host_root_modified"] != false ||
		payload["backend_details_exposed"] != false {
		t.Fatalf("unexpected compatibility center GUI payload: %#v", payload)
	}
	if strings.Contains(output.String(), "/home/xnix-run-materials") || strings.Contains(output.String(), reportPath) {
		t.Fatalf("compatibility center GUI projection exposed raw paths: %s", output.String())
	}
	evidenceItems := payload["known_app_smoke_evidence"].([]any)
	if len(evidenceItems) != 1 {
		t.Fatalf("unexpected GUI evidence items: %#v", evidenceItems)
	}
	evidence := evidenceItems[0].(map[string]any)
	if evidence["evidence_kind"] != "known-application-gui-smoke" ||
		evidence["evidence_source"] != "wine-guest-gui-smoke" ||
		evidence["compatibility_state"] != "real-gui-qemu-wine-verified" ||
		evidence["center_card_state"] != "validated-real-gui-runtime-run" ||
		evidence["marker_observed"] != false ||
		evidence["checksum_verified"] != false ||
		evidence["execution_evidence_recorded"] != true ||
		evidence["runtime_dispatch_verified"] != true ||
		evidence["desktop_launch_enabled"] != false ||
		evidence["backend_launch_enabled"] != false ||
		evidence["backend_details_exposed"] != false {
		t.Fatalf("unexpected compatibility center GUI evidence: %#v", evidence)
	}
}

func TestKDECenterPagePreviewCommandConsumesGUISmokeReport(t *testing.T) {
	registryPath, app := writeTestRepairGroupRegistry(t)
	reportPath := filepath.Join(t.TempDir(), "wine-gui-messagebox.json")
	if err := os.WriteFile(reportPath, []byte(guiSmokeEvidenceCLIFixture()), 0o600); err != nil {
		t.Fatalf("WriteFile report returned error: %v", err)
	}

	var output bytes.Buffer
	if err := run([]string{"kde-center-page-preview", "--registry", registryPath, "--app", app, "--decision", "approved", "--known-app-gui-smoke-report", reportPath}, &output); err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if !strings.Contains(payload["source"].(string), "known-app-gui-smoke-evidence") ||
		payload["known_app_matrix_evidence_count"] != float64(0) ||
		payload["known_app_gui_evidence_count"] != float64(1) ||
		payload["launch_enabled"] != false ||
		payload["execution_started"] != false ||
		payload["backend_process_started"] != false ||
		payload["request_objects_created"] != false ||
		payload["host_root_modified"] != false ||
		payload["backend_details_exposed"] != false {
		t.Fatalf("unexpected KDE center GUI payload: %#v", payload)
	}
	if strings.Contains(output.String(), "/home/xnix-run-materials") || strings.Contains(output.String(), reportPath) {
		t.Fatalf("KDE center GUI projection exposed raw path: %s", output.String())
	}
	cards := payload["known_app_gui_evidence_cards"].([]any)
	if len(cards) != 1 {
		t.Fatalf("unexpected GUI cards: %#v", cards)
	}
	card := cards[0].(map[string]any)
	if card["app_id"] != "org.xnix.fixture.messagebox" ||
		card["evidence_kind"] != "known-application-gui-smoke" ||
		card["evidence_source"] != "wine-guest-gui-smoke" ||
		card["smoke_status"] != "passed" ||
		card["compatibility_state"] != "real-gui-qemu-wine-verified" ||
		card["center_card_state"] != "validated-real-gui-runtime-run" ||
		card["primary_action_id"] != "review-known-app-gui-evidence" ||
		card["primary_action_kind"] != "review" ||
		card["marker_observed"] != false ||
		card["checksum_verified"] != false ||
		card["execution_evidence_recorded"] != true ||
		card["runtime_dispatch_verified"] != true ||
		card["desktop_launch_enabled"] != false ||
		card["backend_launch_enabled"] != false ||
		card["host_root_modified"] != false ||
		card["backend_details_exposed"] != false ||
		card["raw_artifact_path_exposed"] != false {
		t.Fatalf("unexpected KDE GUI card: %#v", card)
	}
}

func knownAppMatrixEvidenceCLIFixture() string {
	return `{
  "schema_version": "xnix.scripts.remote_known_windows_app_matrix_smoke.v1",
  "request_type": "remote-known-winapp-matrix-smoke",
  "status": "passed",
  "execute": true,
  "matrix_report_output": "/home/xnix-run-materials/state/known-run-matrix-version-under-test.json",
  "matrix_report_output_written": true,
  "app_count": 2,
  "backend": "guest-wine",
  "start_qemu": true,
  "guest_port": "auto",
  "redact_output": true,
  "apps": [
    {
      "app_id": "7zr",
      "report_output": "/home/xnix-run-materials/state/known-run-version-under-test-7zr.json",
      "serial_log_output": "/home/xnix-run-materials/state/qemu-serial-version-under-test-7zr.log",
      "status": "passed",
      "app_version": "26.02",
      "executable_name": "7zr.exe",
      "backend": "guest-wine",
      "guest_started": true,
      "guest_port_auto": true,
      "checksum_verified": true,
      "raw_output_redacted": true,
      "marker_observed": true,
      "qemu_serial_log_written": true,
      "qemu_executed": true,
      "wine_executed": true
    },
    {
      "app_id": "busybox-w32",
      "report_output": "/home/xnix-run-materials/state/known-run-version-under-test-busybox-w32.json",
      "serial_log_output": "/home/xnix-run-materials/state/qemu-serial-version-under-test-busybox-w32.log",
      "status": "passed",
      "app_version": "current-2026-07-24",
      "executable_name": "busybox.exe",
      "backend": "guest-wine",
      "guest_started": true,
      "guest_port_auto": true,
      "checksum_verified": true,
      "raw_output_redacted": true,
      "marker_observed": true,
      "qemu_serial_log_written": true,
      "qemu_executed": true,
      "wine_executed": true
    }
  ],
  "host_root_modified": false,
  "privileged_container_required": false,
  "host_networking_required": false,
  "docker_socket_mounted": false,
  "broad_host_mount_required": false,
  "passed_count": 2,
  "failed_count": 0
}`
}

func guiSmokeEvidenceCLIFixture() string {
	return `{
  "version": "version-under-test",
  "schema_version": "xnix.scripts.wine_guest_gui_smoke.v1",
  "request_type": "wine-guest-gui-smoke",
  "status": "passed",
  "execute": true,
  "backend": "qemu-guest-wine-x11",
  "gui_app_name": "xnix-messagebox-smoke.exe",
  "local_gui_executable_configured": true,
  "runtime_go_owned_gui_smoke": true,
  "state_root": "/home/xnix-run-materials/state/wine-gui-messagebox-version-under-test",
  "kernel_image": "/home/xnix-build/xnix-wine-i386-output-gui-version-under-test/images/bzImage",
  "host_root_modified": false,
  "privileged_container_required": false,
  "host_networking_required": false,
  "docker_socket_mounted": false,
  "broad_host_mount_required": false,
  "wineboot_invoked": true,
  "x_window_observed": true,
  "x_window_child_count": 14,
  "runtime_payload_schema_version": "xnix.runtime.windows_app_guest_wine_gui_smoke.v1",
  "executable_copied": true,
  "xwininfo_bytes": 1228,
  "guest_stderr_bytes": 0,
  "guest_graphics_driver_error_observed": false,
  "kde_safe_output_summary": "Wine GUI window observed; child_windows=14 xwininfo_bytes=1228 guest_stderr_bytes=0"
}`
}
