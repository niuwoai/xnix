package main

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestQ4WinAppAcceptancePreviewCommandConsumesGenericWrapperEvidence(t *testing.T) {
	tempDir := t.TempDir()
	smokePath := filepath.Join(tempDir, "q4-winapp.json")
	if err := os.WriteFile(smokePath, []byte(q4WinAppAcceptanceCLIFixture(currentProjectVersion(t))), 0o600); err != nil {
		t.Fatalf("WriteFile smoke returned error: %v", err)
	}

	var output bytes.Buffer
	if err := run([]string{
		"q4-winapp-acceptance-preview",
		"--q4-winapp-smoke", smokePath,
	}, &output); err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["schema_version"] != "xnix.runtime.q4_winapp_acceptance.v1" ||
		payload["request_type"] != "q4-winapp-acceptance-preview" ||
		payload["acceptance_type"] != "generic-q4-windows-app-real-run-acceptance" ||
		payload["smoke_report_consumed"] != true ||
		payload["smoke_report_path_exposed"] != false ||
		payload["output_path_exposed"] != false ||
		payload["delegated_command_exposed"] != false ||
		payload["remote_host_exposed"] != false ||
		payload["app_id"] != "org.xnix.apps.mines" ||
		payload["display_name"] != "Mines" ||
		payload["known_app_selected"] != true ||
		payload["remote_executable_configured"] != false ||
		payload["host_compilation_avoided"] != true ||
		payload["window_match_observed"] != true ||
		payload["document_content_marker_observation_required"] != true ||
		payload["document_content_marker_observed"] != true ||
		payload["acceptance_ready"] != true {
		t.Fatalf("unexpected q4 Windows app acceptance CLI payload: %#v", payload)
	}
	lower := strings.ToLower(output.String())
	for _, forbidden := range []string{strings.ToLower(smokePath), "root@q4", "/tmp/xnix-", "remote_wine_guest_gui_smoke", "qemu-system", "winemine.exe", "sample-document.txt"} {
		if strings.Contains(lower, forbidden) {
			t.Fatalf("q4 Windows app acceptance CLI exposed forbidden term %q: %s", forbidden, output.String())
		}
	}
}

func TestQ4WinAppAcceptancePreviewCommandRequiresSmokeReport(t *testing.T) {
	var output bytes.Buffer
	if err := run([]string{"q4-winapp-acceptance-preview"}, &output); err == nil {
		t.Fatalf("q4-winapp-acceptance-preview must require --q4-winapp-smoke")
	}
	if err := run([]string{"q4-winapp-acceptance-preview", "--q4-winapp-smoke", "missing.json", "extra"}, &output); err == nil {
		t.Fatalf("q4-winapp-acceptance-preview must reject positional arguments")
	}
}

func q4WinAppAcceptanceCLIFixture(version string) string {
	return strings.ReplaceAll(`{
  "schema_version": "xnix.scripts.q4_winapp_smoke.v1",
  "request_type": "q4-winapp-smoke",
  "version": "VERSION_PLACEHOLDER",
  "status": "passed",
  "execute": true,
  "remote_host": "root@q4",
  "delegated_script": "scripts/remote_wine_guest_gui_smoke.rb",
  "delegated_command": ["ruby", "scripts/remote_wine_guest_gui_smoke.rb", "--execute"],
  "output_path": "/tmp/xnix-q4-winapp.json",
  "source_sync_planned": true,
  "known_app_id": "org.xnix.apps.mines",
  "remote_executable_configured": false,
  "remote_executable_path_exposed": false,
  "app_id": "org.xnix.apps.mines",
  "display_name": "Mines",
  "remote_file_argument_configured": false,
  "remote_file_argument_path_exposed": false,
  "sample_file_argument": "sample-document.txt",
  "window_match": "winemine.exe",
  "launch_mode": "owner-controlled-launch",
  "file_open_entrypoint_requested": true,
  "real_run_acceptance_required": true,
  "q4_compile_required": true,
  "host_compilation_avoided": true,
  "host_root_modified": false,
  "privileged_container_required": false,
  "host_networking_required": false,
  "docker_socket_mounted": false,
  "broad_host_mount_required": false,
  "delegated_execute_result_schema": "xnix.scripts.remote_wine_guest_gui_smoke.execute_result.v1",
  "delegated_execute_result_output_written": true,
  "remote_build_completed": true,
  "evidence_output_written": true,
  "kde_page_output_written": true,
  "kde_action_output_written": true,
  "window_observed": true,
  "window_match_observed": true,
  "document_content_marker_observation_required": true,
  "document_content_marker_observed": true,
  "real_run_receipt_summary_ready": true,
  "real_run_receipt_summary_file_open_verified": true,
  "real_run_acceptance_output_written": true,
  "real_run_acceptance_ready": true,
  "real_run_acceptance_center_projection_consumed": true,
  "real_run_acceptance_kde_page_projection_consumed": true,
  "owner_file_open_entrypoint_invoked": true,
  "runtime_evidence_owner_file_open_entrypoint_invoked": true,
  "kde_page_known_app_owner_file_open_entrypoint_count": 1
}`, "VERSION_PLACEHOLDER", version)
}
