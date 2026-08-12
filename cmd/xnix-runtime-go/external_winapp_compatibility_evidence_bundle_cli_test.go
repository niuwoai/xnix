package main

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestExternalWinAppCompatibilityEvidenceBundlePreviewCommandConsumesDesktopArtifacts(t *testing.T) {
	tempDir := t.TempDir()
	oneShotPath := filepath.Join(tempDir, "one-shot.json")
	desktopPacketPath := filepath.Join(tempDir, "desktop-launch-packet.json")
	runtimePacketPath := filepath.Join(tempDir, "runtime-gui-evidence-packet.json")
	kdePagePath := filepath.Join(tempDir, "kde-page.json")
	outputPath := filepath.Join(tempDir, "bundle", "external-winapp-compatibility-evidence-bundle.json")
	writeTextFile(t, oneShotPath, externalWinAppBundleOneShotFixture())
	writeTextFile(t, desktopPacketPath, externalWinAppBundleDesktopLaunchPacketFixture())
	writeTextFile(t, runtimePacketPath, externalWinAppBundleRuntimePacketFixture())
	writeTextFile(t, kdePagePath, externalWinAppBundleKDEPageFixture())

	var output bytes.Buffer
	if err := run([]string{
		"external-winapp-compatibility-evidence-bundle-preview",
		"--one-shot-result", oneShotPath,
		"--desktop-launch-packet", desktopPacketPath,
		"--runtime-gui-evidence-packet", runtimePacketPath,
		"--kde-page", kdePagePath,
		"--output", outputPath,
	}, &output); err != nil {
		t.Fatalf("run returned error: %v", err)
	}
	written, err := os.ReadFile(outputPath)
	if err != nil {
		t.Fatalf("ReadFile output returned error: %v", err)
	}
	if string(written) != output.String() {
		t.Fatalf("written bundle must match stdout\nstdout=%s\nwritten=%s", output.String(), string(written))
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["schema_version"] != "xnix.runtime.external_winapp_compatibility_evidence_bundle.v1" ||
		payload["request_type"] != "external-winapp-compatibility-evidence-bundle-preview" ||
		payload["bundle_type"] != "runtime-owned-external-windows-app-compatibility-evidence" ||
		payload["runtime_method"] != "PreviewExternalWinAppCompatibilityEvidenceBundle" ||
		payload["read_method"] != "GetExternalWinAppCompatibilityEvidenceBundle" ||
		payload["desktop"] != "KDE Plasma" ||
		payload["application_id"] != "org.xnix.external.desktop-notepad-file-argument" ||
		payload["display_name"] != "External Desktop Notepad File Argument" ||
		payload["evidence_artifact_count"] != float64(4) ||
		payload["one_shot_runtime_launch_verified"] != true ||
		payload["desktop_launch_packet_verified"] != true ||
		payload["runtime_gui_evidence_packet_verified"] != true ||
		payload["kde_external_app_page_verified"] != true ||
		payload["real_windows_app_run_verified"] != true ||
		payload["external_app_import_record_consumed"] != true ||
		payload["external_app_handle_consumed"] != true ||
		payload["imported_artifact_digest_verified"] != true ||
		payload["external_file_bridge_ready"] != true ||
		payload["external_file_open_requested"] != true ||
		payload["external_desktop_argument_count"] != float64(1) ||
		payload["window_observed"] != true ||
		payload["x_window_observed"] != true ||
		payload["container_runtime_used"] != true ||
		payload["container_network_mode"] != "none" ||
		payload["container_host_mount_count"] != float64(0) ||
		payload["known_app_gui_evidence_count"] != float64(1) ||
		payload["known_app_gui_evidence_verified_count"] != float64(1) ||
		payload["kde_page_known_app_gui_evidence_count"] != float64(1) ||
		payload["runtime_owned"] != true ||
		payload["go_runtime_backed"] != true ||
		payload["kde_policy_owner"] != false ||
		payload["safe_for_kde"] != true ||
		payload["safe_for_ai_diagnostics"] != true ||
		payload["desktop_launch_enabled"] != false ||
		payload["backend_launch_enabled"] != false ||
		payload["backend_process_started"] != true ||
		payload["backend_details_exposed"] != false ||
		payload["raw_paths_exposed"] != false ||
		payload["raw_launcher_output_exposed"] != false ||
		payload["host_root_modified"] != false ||
		payload["docker_socket_mounted"] != false ||
		payload["broad_host_mount_required"] != false {
		t.Fatalf("unexpected external Windows app compatibility evidence bundle: %#v", payload)
	}
	artifacts := payload["evidence_artifacts"].([]any)
	if len(artifacts) != 4 {
		t.Fatalf("unexpected evidence artifacts: %#v", artifacts)
	}
	if strings.Contains(output.String(), oneShotPath) ||
		strings.Contains(output.String(), desktopPacketPath) ||
		strings.Contains(output.String(), runtimePacketPath) ||
		strings.Contains(output.String(), kdePagePath) ||
		strings.Contains(output.String(), "docker run") ||
		strings.Contains(output.String(), "/var/run/docker.sock") {
		t.Fatalf("compatibility evidence bundle exposed unsafe details: %s", output.String())
	}
}

func writeTextFile(t *testing.T, path string, content string) {
	t.Helper()
	if err := os.MkdirAll(filepath.Dir(path), 0o700); err != nil {
		t.Fatalf("MkdirAll returned error: %v", err)
	}
	if err := os.WriteFile(path, []byte(content), 0o600); err != nil {
		t.Fatalf("WriteFile returned error: %v", err)
	}
}

func externalWinAppBundleOneShotFixture() string {
	return `{
  "version": "0.2.640-test",
  "schema_version": "xnix.runtime.external_winapp_import_stage_launch.v1",
  "request_type": "external-winapp-import-stage-and-launch",
  "status": "passed",
  "application_id": "org.xnix.external.desktop-notepad-file-argument",
  "display_name": "External Desktop Notepad File Argument",
  "external_app_handle": "org.xnix.external.desktop-notepad-file-argument",
  "launcher_request_type": "windows-external-app-run",
  "launcher_status": "passed",
  "import_recorded": true,
  "desktop_activation_staged": true,
  "staged_launcher_invoked": true,
  "managed_launcher_executable_staged": true,
  "desktop_exec_uses_external_app_handle": true,
  "external_app_desktop_handle_ready": true,
  "desktop_launch_packet_written": true,
  "external_app_import_record_consumed": true,
  "external_app_handle_consumed": true,
  "external_file_bridge_ready": true,
  "imported_artifact_digest_verified": true,
  "runtime_launch_executed": true,
  "window_observed": true,
  "x_window_observed": true,
  "runtime_owned": true,
  "go_runtime_backed": true,
  "kde_policy_owner": false,
  "launch_enabled": false,
  "backend_launch_enabled": false,
  "host_root_modified": false,
  "docker_socket_mounted": false,
  "broad_host_mount_required": false,
  "raw_import_record_path_exposed": false,
  "raw_state_root_path_exposed": false,
  "raw_executable_path_exposed": false,
  "raw_launcher_path_exposed": false,
  "raw_launcher_output_exposed": false,
  "privileged_container_required": false,
  "host_networking_required": false
}`
}

func externalWinAppBundleDesktopLaunchPacketFixture() string {
	return `{
  "version": "0.2.640-test",
  "schema_version": "xnix.runtime.desktop_external_winapp_launch_packet.v1",
  "request_type": "desktop-external-winapp-launch-packet-preview",
  "status": "passed",
  "desktop": "KDE Plasma",
  "application_id": "org.xnix.external.desktop-notepad-file-argument",
  "display_name": "External Desktop Notepad File Argument",
  "app_version": "0.2.640-test",
  "external_app_handle": "org.xnix.external.desktop-notepad-file-argument",
  "desktop_exec_uses_external_app_handle": true,
  "external_app_desktop_handle_ready": true,
  "desktop_exec_uses_raw_import_record": false,
  "desktop_exec_uses_state_root": false,
  "external_app_run_record_consumed": true,
  "external_app_import_record_consumed": true,
  "external_app_handle_consumed": true,
  "external_desktop_argument_count": 1,
  "external_file_uri_arguments_accepted": true,
  "external_file_open_requested": true,
  "external_file_bridge_ready": true,
  "imported_artifact_digest_verified": true,
  "runtime_launch_executed": true,
  "backend_process_started": true,
  "window_observed": true,
  "x_window_observed": true,
  "container_runtime_used": true,
  "container_network_mode": "none",
  "container_host_mount_count": 0,
  "runtime_owned": true,
  "go_runtime_backed": true,
  "kde_policy_owner": false,
  "desktop_launch_packet_ready": true,
  "safe_for_kde": true,
  "backend_details_exposed": false,
  "raw_import_record_path_exposed": false,
  "raw_external_app_handle_path_exposed": false,
  "raw_state_root_path_exposed": false,
  "raw_executable_path_exposed": false,
  "host_root_modified": false,
  "privileged_container_required": false,
  "host_networking_required": false,
  "docker_socket_mounted": false,
  "broad_host_mount_required": false
}`
}

func externalWinAppBundleRuntimePacketFixture() string {
	return `{
  "version": "0.2.640-test",
  "schema_version": "xnix.runtime.real_winapp_gui_evidence_packet.v1",
  "request_type": "real-winapp-gui-evidence-packet-preview",
  "report_status": "passed",
  "report_consumed": true,
  "app_id": "org.xnix.external.desktop-notepad-file-argument",
  "display_name": "External Desktop Notepad File Argument",
  "app_version": "0.2.640-test",
  "known_app_gui_evidence_count": 1,
  "known_app_gui_evidence_verified_count": 1,
  "x_window_observed": true,
  "window_observed": true,
  "container_runtime_used": true,
  "container_network_mode": "none",
  "container_host_mount_count": 0,
  "external_app_run_record_consumed": true,
  "external_app_handle_consumed": true,
  "external_app_import_record_consumed": true,
  "imported_artifact_digest_verified": true,
  "runtime_owned": true,
  "go_runtime_backed": true,
  "kde_policy_owner": false,
  "desktop_launch_enabled": false,
  "backend_launch_enabled": false,
  "action_execution_enabled": false,
  "backend_details_exposed": false,
  "raw_output_exposed": false,
  "host_root_modified": false,
  "privileged_container_required": false,
  "host_networking_required": false,
  "docker_socket_mounted": false,
  "broad_host_mount_required": false
}`
}

func externalWinAppBundleKDEPageFixture() string {
	return `{
  "schema_version": "xnix.runtime.kde_center_page.v1",
  "request_type": "kde-center-page-preview",
  "desktop": "KDE Plasma",
  "application_id": "org.xnix.external.desktop-notepad-file-argument",
  "application_name": "External Desktop Notepad File Argument",
  "known_app_gui_evidence_count": 1,
  "runtime_owned": true,
  "go_runtime_backed": true,
  "kde_policy_owner": false,
  "launch_enabled": false,
  "backend_process_started": false,
  "host_root_modified": false,
  "backend_details_exposed": false
}`
}
