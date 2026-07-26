package main

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestRealWinAppRunAcceptancePreviewCommandConsumesReceiptBackedExecuteResult(t *testing.T) {
	tempDir := t.TempDir()
	executeResultPath := filepath.Join(tempDir, "remote-execute-result.json")
	if err := os.WriteFile(executeResultPath, []byte(realWinAppRunAcceptanceCLIFixture(currentProjectVersion(t))), 0o600); err != nil {
		t.Fatalf("WriteFile execute result returned error: %v", err)
	}

	var output bytes.Buffer
	if err := run([]string{
		"real-winapp-run-acceptance-preview",
		"--remote-execute-result", executeResultPath,
	}, &output); err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["schema_version"] != "xnix.runtime.real_winapp_run_acceptance.v1" ||
		payload["request_type"] != "real-winapp-run-acceptance-preview" ||
		payload["acceptance_type"] != "real-windows-app-run-acceptance" ||
		payload["execute_result_consumed"] != true ||
		payload["execute_result_path_exposed"] != false ||
		payload["receipt_summary_path_exposed"] != false ||
		payload["center_projection_path_exposed"] != false ||
		payload["kde_page_projection_path_exposed"] != false ||
		payload["remote_host_exposed"] != false ||
		payload["app_id"] != "org.xnix.sample.notepad" ||
		payload["display_name"] != "Sample Notepad" ||
		payload["run_passed"] != true ||
		payload["real_execution_observed"] != true ||
		payload["owner_controlled_file_open_verified"] != true ||
		payload["owner_file_open_entrypoint_invoked"] != true ||
		payload["document_content_marker_observation_required"] != true ||
		payload["document_content_marker_observed"] != true ||
		payload["receipt_summary_generated"] != true ||
		payload["receipt_summary_ready"] != true ||
		payload["compatibility_center_projection_consumed"] != true ||
		payload["compatibility_center_known_app_smoke_evidence_count"] != float64(1) ||
		payload["kde_page_projection_consumed"] != true ||
		payload["kde_page_known_app_gui_evidence_count"] != float64(1) ||
		payload["kde_page_owner_file_open_verified_count"] != float64(1) ||
		payload["kde_page_owner_file_open_entrypoint_count"] != float64(1) ||
		payload["acceptance_ready"] != true ||
		payload["backend_details_exposed"] != false ||
		payload["raw_window_evidence_exposed"] != false ||
		payload["executable_name_exposed"] != false {
		t.Fatalf("unexpected real run acceptance CLI payload: %#v", payload)
	}
	lower := strings.ToLower(output.String())
	for _, forbidden := range []string{strings.ToLower(executeResultPath), "root@q4", "/home/xnix-", "notepad.exe", "qemu-system", "/var/run/docker.sock", "sample-document.txt - notepad"} {
		if strings.Contains(lower, forbidden) {
			t.Fatalf("real run acceptance CLI exposed forbidden term %q: %s", forbidden, output.String())
		}
	}
}

func TestRealWinAppRunAcceptancePreviewCommandRequiresExecuteResult(t *testing.T) {
	var output bytes.Buffer
	if err := run([]string{"real-winapp-run-acceptance-preview"}, &output); err == nil {
		t.Fatalf("real-winapp-run-acceptance-preview must require --remote-execute-result")
	}
	if err := run([]string{"real-winapp-run-acceptance-preview", "--remote-execute-result", "missing.json", "extra"}, &output); err == nil {
		t.Fatalf("real-winapp-run-acceptance-preview must reject positional arguments")
	}
}

func realWinAppRunAcceptanceCLIFixture(version string) string {
	return strings.ReplaceAll(`{
  "schema_version": "xnix.scripts.remote_wine_guest_gui_smoke.execute_result.v1",
  "request_type": "remote-wine-guest-gui-smoke",
  "version": "VERSION_PLACEHOLDER",
  "status": "passed",
  "smoke_status": "passed",
  "remote_host": "root@q4",
  "execute_result_output_written": true,
  "real_run_receipt_summary_preview_planned": true,
  "real_run_receipt_summary_output_written": true,
  "real_run_receipt_summary_ready": true,
  "real_run_receipt_summary_file_open_verified": true,
  "real_run_receipt_summary_center_output_written": true,
  "real_run_receipt_summary_center_known_app_smoke_evidence_count": 1,
  "real_run_receipt_summary_kde_page_output_written": true,
  "real_run_receipt_summary_kde_page_known_app_gui_evidence_count": 1,
  "real_run_receipt_summary_kde_page_owner_file_open_verified_count": 1,
  "real_run_receipt_summary_kde_page_owner_file_open_entrypoint_count": 1,
  "report_output": "/home/xnix-run-materials/state/wine-gui-smoke-VERSION_PLACEHOLDER.json",
  "gui_app_name": "notepad.exe",
  "known_app_id": "org.xnix.sample.notepad",
  "known_app_name": "Sample Notepad",
  "known_app_version": "VERSION_PLACEHOLDER",
  "launch_mode": "owner-controlled-launch",
  "file_open_entrypoint_requested": true,
  "x_window_observed": true,
  "window_match_observed": true,
  "document_content_marker_observation_required": true,
  "document_content_marker_observed": true,
  "window_evidence_summary": "0xa00003 \"file-1-sample-document.txt - Notepad\": (\"notepad.exe\" \"notepad.exe\")",
  "runtime_evidence_report_consumed": true,
  "runtime_evidence_window_observed": true,
  "runtime_evidence_owner_file_open_verified": true,
  "kde_action_owner_file_open_verified": true,
  "owner_controlled_launch_requested": true,
  "owner_file_open_entrypoint_invoked": true,
  "owner_delegated_file_arguments_passed": true,
  "owner_delegated_file_argument_winepath_translated": true,
  "owner_delegated_raw_file_argument_path_exposed": false,
  "host_root_modified": false,
  "privileged_container_required": false,
  "host_networking_required": false,
  "docker_socket_mounted": false,
  "broad_host_mount_required": false
}`, "VERSION_PLACEHOLDER", version)
}
