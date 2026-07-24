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
