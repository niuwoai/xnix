package main

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestQ4SampleNotepadAcceptancePreviewCommandConsumesWrapperEvidence(t *testing.T) {
	tempDir := t.TempDir()
	smokePath := filepath.Join(tempDir, "q4-sample-notepad.json")
	if err := os.WriteFile(smokePath, []byte(q4SampleNotepadAcceptanceCLIFixture(currentProjectVersion(t))), 0o600); err != nil {
		t.Fatalf("WriteFile smoke returned error: %v", err)
	}

	var output bytes.Buffer
	if err := run([]string{
		"q4-sample-notepad-acceptance-preview",
		"--q4-sample-notepad-smoke", smokePath,
	}, &output); err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["schema_version"] != "xnix.runtime.q4_sample_notepad_acceptance.v1" ||
		payload["request_type"] != "q4-sample-notepad-acceptance-preview" ||
		payload["acceptance_type"] != "q4-sample-notepad-real-windows-app-acceptance" ||
		payload["smoke_report_consumed"] != true ||
		payload["smoke_report_path_exposed"] != false ||
		payload["output_path_exposed"] != false ||
		payload["delegated_command_exposed"] != false ||
		payload["remote_host_exposed"] != false ||
		payload["app_id"] != "org.xnix.sample.notepad" ||
		payload["display_name"] != "Sample Notepad" ||
		payload["host_compilation_avoided"] != true ||
		payload["real_run_acceptance_ready"] != true ||
		payload["window_match_observed"] != true ||
		payload["owner_file_open_entrypoint_invoked"] != true ||
		payload["acceptance_ready"] != true {
		t.Fatalf("unexpected q4 Sample Notepad acceptance CLI payload: %#v", payload)
	}
	lower := strings.ToLower(output.String())
	for _, forbidden := range []string{strings.ToLower(smokePath), "root@q4", "/tmp/xnix-", "remote_wine_guest_gui_smoke", "qemu-system", "notepad.exe", "sample-document.txt"} {
		if strings.Contains(lower, forbidden) {
			t.Fatalf("q4 Sample Notepad acceptance CLI exposed forbidden term %q: %s", forbidden, output.String())
		}
	}
}

func TestQ4SampleNotepadAcceptancePreviewCommandRequiresSmokeReport(t *testing.T) {
	var output bytes.Buffer
	if err := run([]string{"q4-sample-notepad-acceptance-preview"}, &output); err == nil {
		t.Fatalf("q4-sample-notepad-acceptance-preview must require --q4-sample-notepad-smoke")
	}
	if err := run([]string{"q4-sample-notepad-acceptance-preview", "--q4-sample-notepad-smoke", "missing.json", "extra"}, &output); err == nil {
		t.Fatalf("q4-sample-notepad-acceptance-preview must reject positional arguments")
	}
}

func q4SampleNotepadAcceptanceCLIFixture(version string) string {
	return strings.ReplaceAll(`{
  "schema_version": "xnix.scripts.q4_sample_notepad_smoke.v1",
  "request_type": "q4-sample-notepad-smoke",
  "version": "VERSION_PLACEHOLDER",
  "status": "passed",
  "execute": true,
  "remote_host": "root@q4",
  "delegated_script": "scripts/remote_wine_guest_gui_smoke.rb",
  "delegated_command": ["ruby", "scripts/remote_wine_guest_gui_smoke.rb", "--execute"],
  "output_path": "/tmp/xnix-q4-sample-notepad.json",
  "app_id": "org.xnix.sample.notepad",
  "display_name": "Sample Notepad",
  "sample_file_argument": "sample-document.txt",
  "window_match": "sample-document.txt",
  "launch_mode": "owner-controlled-launch",
  "file_open_entrypoint_requested": true,
  "host_compilation_avoided": true,
  "host_root_modified": false,
  "privileged_container_required": false,
  "host_networking_required": false,
  "docker_socket_mounted": false,
  "broad_host_mount_required": false,
  "delegated_execute_result_schema": "xnix.scripts.remote_wine_guest_gui_smoke.execute_result.v1",
  "delegated_execute_result_output_written": true,
  "real_run_receipt_summary_ready": true,
  "real_run_receipt_summary_file_open_verified": true,
  "real_run_acceptance_output_written": true,
  "real_run_acceptance_ready": true,
  "real_run_acceptance_center_projection_consumed": true,
  "real_run_acceptance_kde_page_projection_consumed": true,
  "window_match_observed": true,
  "owner_file_open_entrypoint_invoked": true,
  "runtime_evidence_owner_file_open_entrypoint_invoked": true,
  "kde_page_known_app_owner_file_open_entrypoint_count": 1
}`, "VERSION_PLACEHOLDER", version)
}
