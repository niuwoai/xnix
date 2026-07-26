package main

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestKnownExistingWinAppAcceptancePreviewCommandConsumes7zrRun(t *testing.T) {
	tempDir := t.TempDir()
	runPath := filepath.Join(tempDir, "known-run.json")
	if err := os.WriteFile(runPath, []byte(knownExistingWinAppAcceptanceCLIFixture(currentProjectVersion(t))), 0o600); err != nil {
		t.Fatalf("WriteFile run report returned error: %v", err)
	}

	var output bytes.Buffer
	if err := run([]string{
		"known-existing-winapp-acceptance-preview",
		"--known-winapp-run", runPath,
	}, &output); err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["schema_version"] != "xnix.runtime.known_existing_winapp_acceptance.v1" ||
		payload["request_type"] != "known-existing-winapp-acceptance-preview" ||
		payload["acceptance_type"] != "known-existing-windows-app-real-run-acceptance" ||
		payload["run_report_consumed"] != true ||
		payload["run_report_path_exposed"] != false ||
		payload["remote_host_exposed"] != false ||
		payload["guest_endpoint_exposed"] != false ||
		payload["raw_path_exposed"] != false ||
		payload["raw_output_exposed"] != false ||
		payload["app_id"] != "7zr" ||
		payload["display_name"] != "7-Zip standalone console executable" ||
		payload["existing_windows_app"] != true ||
		payload["known_portable_catalog_backed"] != true ||
		payload["checksum_verified"] != true ||
		payload["marker_observed"] != true ||
		payload["isolated_guest_execution_observed"] != true ||
		payload["compatibility_engine_execution_observed"] != true ||
		payload["loopback_only_networking"] != true ||
		payload["acceptance_ready"] != true {
		t.Fatalf("unexpected known existing Windows app acceptance CLI payload: %#v", payload)
	}
	lower := strings.ToLower(output.String())
	for _, forbidden := range []string{strings.ToLower(runPath), "root@q4", "/home/xnix-", "42281", "guest_port", "qemu-system", "7zr.exe", "wine ", ".wine"} {
		if strings.Contains(lower, forbidden) {
			t.Fatalf("known existing Windows app acceptance CLI exposed forbidden term %q: %s", forbidden, output.String())
		}
	}
}

func TestKnownExistingWinAppAcceptancePreviewCommandRequiresRunReport(t *testing.T) {
	var output bytes.Buffer
	if err := run([]string{"known-existing-winapp-acceptance-preview"}, &output); err == nil {
		t.Fatalf("known-existing-winapp-acceptance-preview must require --known-winapp-run")
	}
	if err := run([]string{"known-existing-winapp-acceptance-preview", "--known-winapp-run", "missing.json", "extra"}, &output); err == nil {
		t.Fatalf("known-existing-winapp-acceptance-preview must reject positional arguments")
	}
}

func knownExistingWinAppAcceptanceCLIFixture(version string) string {
	return strings.ReplaceAll(`{
  "schema_version": "xnix.runtime.known_windows_app_run.v1",
  "request_type": "windows-known-app-run",
  "version": "VERSION_PLACEHOLDER",
  "status": "passed",
  "app_id": "7zr",
  "display_name": "7-Zip standalone console executable",
  "app_version": "26.02",
  "architecture": "windows-x86",
  "executable_name": "7zr.exe",
  "backend_ready": true,
  "launch_attempted": true,
  "runner_available": true,
  "checksum_verified": true,
  "guest_start_attempted": true,
  "guest_started": true,
  "guest_start_mode": "go-qemu",
  "qemu_serial_log_written": true,
  "executable_copied": true,
  "marker_observed": true,
  "raw_output_redacted": true,
  "guest_payload": {
    "schema_version": "xnix.runtime.known_windows_app_guest_wine_smoke.v1",
    "request_type": "windows-known-app-guest-wine-smoke",
    "status": "passed",
    "checksum_verified": true,
    "guest": {
      "schema_version": "xnix.runtime.windows_app_guest_wine_smoke.v1",
      "status": "passed",
      "guest_reachable": true,
      "wine_available": true,
      "executable_copied": true,
      "marker_observed": true,
      "exit_code": 0,
      "stdout_bytes": 2902,
      "stderr_bytes": 170,
      "stdout_line_count": 66,
      "stderr_line_count": 2,
      "raw_output_included": false,
      "raw_output_redacted": true,
      "loopback_only_networking": true,
      "host_root_modified": false,
      "privileged_container_required": false,
      "host_networking_required": false,
      "docker_socket_mounted": false,
      "broad_host_mount_required": false,
      "raw_host_path_exposed": false
    }
  },
  "loopback_only_networking": true,
  "network_required": false,
  "host_root_modified": false,
  "privileged_container_required": false,
  "host_networking_required": false,
  "docker_socket_mounted": false,
  "broad_host_mount_required": false,
  "docker_executed": false,
  "colima_executed": false,
  "network_checks_run": false,
  "package_manager_invoked": false,
  "raw_host_path_exposed": false,
  "raw_executable_path_exposed": false,
  "raw_profile_path_exposed": false,
  "raw_state_root_path_exposed": false,
  "raw_runtime_argv_exposed": false,
  "raw_runner_path_exposed": false,
  "raw_qemu_path_exposed": false
}`, "VERSION_PLACEHOLDER", version)
}
