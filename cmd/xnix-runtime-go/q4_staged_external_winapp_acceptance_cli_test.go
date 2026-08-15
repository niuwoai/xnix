package main

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestQ4StagedExternalWinAppAcceptancePreviewCommandConsumesWrapperEvidence(t *testing.T) {
	tempDir := t.TempDir()
	smokePath := filepath.Join(tempDir, "q4-staged-external-winapp.json")
	outputPath := filepath.Join(tempDir, "acceptance.json")
	if err := os.WriteFile(smokePath, []byte(q4StagedExternalWinAppAcceptanceCLIFixture(currentProjectVersion(t))), 0o600); err != nil {
		t.Fatalf("WriteFile smoke returned error: %v", err)
	}

	var output bytes.Buffer
	if err := run([]string{
		"q4-staged-external-winapp-acceptance-preview",
		"--q4-staged-external-winapp-smoke", smokePath,
		"--output", outputPath,
	}, &output); err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["schema_version"] != "xnix.runtime.q4_staged_external_winapp_acceptance.v1" ||
		payload["request_type"] != "q4-staged-external-winapp-acceptance-preview" ||
		payload["acceptance_type"] != "staged-external-winapp-desktop-file-open-acceptance" ||
		payload["smoke_report_consumed"] != true ||
		payload["smoke_report_path_exposed"] != false ||
		payload["delegated_command_exposed"] != false ||
		payload["remote_host_exposed"] != false ||
		payload["app_id"] != "org.xnix.external.desktop-notepad-file-argument" ||
		payload["host_compilation_avoided"] != true ||
		payload["desktop_exec_uses_external_app_handle"] != true ||
		payload["windows_process_file_argument_window_observed"] != true ||
		payload["kde_page_from_application_detail_header_badge"] != "Verified real app run" ||
		payload["kde_page_from_application_detail_summary_compatibility_state"] != "real-app-run-verified" ||
		payload["acceptance_ready"] != true {
		t.Fatalf("unexpected q4 staged external Windows app acceptance CLI payload: %#v", payload)
	}
	written, err := os.ReadFile(outputPath)
	if err != nil {
		t.Fatalf("ReadFile output returned error: %v", err)
	}
	var writtenPayload map[string]any
	if err := json.Unmarshal(written, &writtenPayload); err != nil {
		t.Fatalf("Unmarshal written output returned error: %v", err)
	}
	if writtenPayload["schema_version"] != payload["schema_version"] ||
		writtenPayload["acceptance_ready"] != true {
		t.Fatalf("unexpected written q4 staged external Windows app acceptance payload: %#v", writtenPayload)
	}
	lower := strings.ToLower(output.String())
	for _, forbidden := range []string{strings.ToLower(smokePath), "root@q4", "/tmp/xnix-", "staged_desktop_external_winapp_smoke", "notepad.exe"} {
		if strings.Contains(lower, forbidden) {
			t.Fatalf("q4 staged external Windows app acceptance CLI exposed forbidden term %q: %s", forbidden, output.String())
		}
	}
}

func TestQ4StagedExternalWinAppAcceptancePreviewCommandRequiresSmokeReport(t *testing.T) {
	var output bytes.Buffer
	if err := run([]string{"q4-staged-external-winapp-acceptance-preview"}, &output); err == nil {
		t.Fatalf("q4-staged-external-winapp-acceptance-preview must require --q4-staged-external-winapp-smoke")
	}
	if err := run([]string{"q4-staged-external-winapp-acceptance-preview", "--q4-staged-external-winapp-smoke", "missing.json", "extra"}, &output); err == nil {
		t.Fatalf("q4-staged-external-winapp-acceptance-preview must reject positional arguments")
	}
}

func q4StagedExternalWinAppAcceptanceCLIFixture(version string) string {
	return strings.ReplaceAll(`{
  "schema_version": "xnix.scripts.q4_staged_desktop_external_winapp_smoke.v1",
  "request_type": "q4-staged-desktop-external-winapp-smoke",
  "version": "VERSION_PLACEHOLDER",
  "status": "passed",
  "execute": true,
  "remote_host": "root@q4",
  "delegated_command": ["ruby", "scripts/staged_desktop_external_winapp_smoke.rb", "--execute"],
  "q4_compile_required": true,
  "host_compilation_avoided": true,
  "remote_build_completed": true,
  "delegated_status": "passed",
  "app_id": "org.xnix.external.desktop-notepad-file-argument",
  "display_name": "External Desktop Notepad File Argument",
  "desktop_argument_mode": "file-uri",
  "desktop_file_open_lane": true,
  "desktop_exec_uses_external_app_handle": true,
  "external_app_desktop_handle_ready": true,
  "desktop_exec_invocation_exact": true,
  "external_file_open_requested": true,
  "external_file_bridge_ready": true,
  "external_file_bridge_arguments_passed": true,
  "external_file_bridge_winepath_translated": true,
  "windows_process_file_argument_window_observed": true,
  "external_file_bridge_mount_enabled": false,
  "raw_file_uri_arguments_exposed": false,
  "external_app_handle_consumed": true,
  "imported_artifact_digest_verified": true,
  "window_observed": true,
  "x_window_observed": true,
  "container_network_mode": "none",
  "container_host_mount_count": 0,
  "one_shot_runtime_launch_executed": true,
  "one_shot_window_observed": true,
  "one_shot_host_root_modified": false,
  "one_shot_docker_socket_mounted": false,
  "one_shot_raw_paths_exposed": false,
  "runtime_packet_external_app_run_record_consumed": true,
  "kde_page_card_external_app_run_record_consumed": true,
  "compatibility_evidence_bundle_generated": true,
  "compatibility_evidence_bundle_real_windows_app_run_verified": true,
  "application_detail_generated": true,
  "application_detail_real_windows_app_run_verified": true,
  "application_detail_file_open_verified": true,
  "kde_page_from_application_detail_generated": true,
  "kde_page_from_application_detail_consumed": true,
  "kde_page_from_application_detail_header_badge": "Verified real app run",
  "kde_page_from_application_detail_header_badge_tone": "success",
  "kde_page_from_application_detail_header_backend_details_exposed": false,
  "kde_page_from_application_detail_summary_compatibility_state": "real-app-run-verified",
  "kde_page_from_application_detail_summary_compatibility_label": "Verified real app run",
  "kde_page_from_application_detail_summary_diagnostics_state": "real-app-run-verified",
  "kde_page_from_application_detail_summary_backend_launch_enabled": false,
  "kde_page_from_application_detail_summary_action_execution_enabled": false,
  "kde_page_from_application_detail_summary_settings_persistence_enabled": false,
  "kde_page_from_application_detail_summary_backend_details_exposed": false,
  "kde_page_from_application_detail_summary_host_root_modified": false,
  "artifact_fetch_count": 10,
  "host_root_modified": false,
  "privileged_container_required": false,
  "host_networking_required": false,
  "docker_socket_mounted": false,
  "broad_host_mount_required": false
}`, "VERSION_PLACEHOLDER", version)
}
