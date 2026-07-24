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
