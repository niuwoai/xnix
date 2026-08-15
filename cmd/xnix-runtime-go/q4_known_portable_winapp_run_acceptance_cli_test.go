package main

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestQ4KnownPortableWinAppRunAcceptancePreviewCommandConsumesOperatorRun(t *testing.T) {
	tempDir := t.TempDir()
	runPath := filepath.Join(tempDir, "q4-known-portable-winapp-run.json")
	outputPath := filepath.Join(tempDir, "acceptance", "q4-known-portable-winapp-run-acceptance.json")
	if err := os.WriteFile(runPath, []byte(q4KnownPortableWinAppRunAcceptanceCLIFixture(currentProjectVersion(t))), 0o644); err != nil {
		t.Fatalf("WriteFile returned error: %v", err)
	}

	var output bytes.Buffer
	if err := run([]string{
		"q4-known-portable-winapp-run-acceptance-preview",
		"--q4-known-portable-winapp-run", runPath,
		"--output", outputPath,
	}, &output); err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if _, err := os.Stat(outputPath); err != nil {
		t.Fatalf("expected acceptance output file: %v", err)
	}
	if payload["version"] != currentProjectVersion(t) ||
		payload["schema_version"] != "xnix.runtime.q4_known_portable_winapp_run_acceptance.v1" ||
		payload["request_type"] != "q4-known-portable-winapp-run-acceptance-preview" ||
		payload["acceptance_type"] != "q4-known-portable-real-winapp-operator-run-acceptance" ||
		payload["app_id"] != "org.xnix.external.notepadplusplus" ||
		payload["display_name"] != "Notepad++ Portable" ||
		payload["accepted_application_detail_state"] != "runtime-accepted-real-app-run" ||
		payload["kde_accepted_page_state"] != "runtime-accepted-real-app-run" {
		t.Fatalf("unexpected q4 known portable Windows app operator acceptance: %#v", payload)
	}
	if payload["run_report_consumed"] != true ||
		payload["run_report_path_exposed"] != false ||
		payload["remote_host_exposed"] != false ||
		payload["delegated_command_exposed"] != false ||
		payload["raw_path_exposed"] != false ||
		payload["known_catalog_app"] != true ||
		payload["portable_directory_external_app"] != true ||
		payload["official_archive_checksum_verified"] != true ||
		payload["known_portable_bundle_imported"] != true ||
		payload["known_portable_bundle_acceptance_ready"] != true ||
		payload["runtime_accepted_chain_verified"] != true ||
		payload["operator_run_ready"] != true ||
		payload["q4_download_required"] != true ||
		payload["q4_extract_required"] != true ||
		payload["q4_compile_required"] != true ||
		payload["q4_execution_required"] != true ||
		payload["host_compilation_avoided"] != true ||
		payload["host_download_avoided"] != true ||
		payload["host_root_modified"] != false ||
		payload["privileged_container_required"] != false ||
		payload["host_networking_required"] != false ||
		payload["docker_socket_mounted"] != false ||
		payload["broad_host_mount_required"] != false ||
		payload["backend_details_exposed"] != false ||
		payload["acceptance_ready"] != true {
		t.Fatalf("unexpected safety flags: %#v", payload)
	}
	for _, forbidden := range []string{runPath, outputPath, "root@q4", "/home/xnix-", "/Users/rocky", "docker run", "/var/run/docker.sock", " notepad++.exe "} {
		if strings.Contains(output.String(), forbidden) {
			t.Fatalf("q4 known portable run acceptance exposed forbidden term %q: %s", forbidden, output.String())
		}
	}
}

func TestQ4KnownPortableWinAppRunAcceptancePreviewCommandConsumesPuttyOperatorRun(t *testing.T) {
	tempDir := t.TempDir()
	runPath := filepath.Join(tempDir, "q4-known-portable-putty-run.json")
	if err := os.WriteFile(runPath, []byte(q4KnownPortableWinAppRunAcceptancePuttyFixture(currentProjectVersion(t))), 0o644); err != nil {
		t.Fatalf("WriteFile returned error: %v", err)
	}

	var output bytes.Buffer
	if err := run([]string{
		"q4-known-portable-winapp-run-acceptance-preview",
		"--q4-known-portable-winapp-run", runPath,
	}, &output); err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["app_id"] != "org.xnix.external.putty" ||
		payload["display_name"] != "PuTTY" ||
		payload["portable_directory_external_app"] != false ||
		payload["single_file_external_app"] != true ||
		payload["official_archive_checksum_verified"] != false ||
		payload["official_executable_checksum_verified"] != true ||
		payload["known_portable_bundle_imported"] != false ||
		payload["staged_external_winapp_acceptance_ready"] != true ||
		payload["external_file_open_requested"] != false ||
		payload["external_file_bridge_ready"] != false ||
		payload["windows_process_file_argument_window_observed"] != false ||
		payload["window_observed"] != true ||
		payload["x_window_observed"] != true ||
		payload["q4_extract_required"] != false ||
		payload["runtime_accepted_chain_verified"] != true ||
		payload["acceptance_ready"] != true {
		t.Fatalf("unexpected PuTTY q4 known portable Windows app operator acceptance: %#v", payload)
	}
	for _, forbidden := range []string{runPath, "root@q4", "/home/xnix-", "/Users/rocky", "docker run", "/var/run/docker.sock", " putty.exe "} {
		if strings.Contains(output.String(), forbidden) {
			t.Fatalf("q4 known portable PuTTY acceptance exposed forbidden term %q: %s", forbidden, output.String())
		}
	}
}

func TestQ4KnownPortableWinAppRunAcceptancePreviewCommandRequiresPassedRun(t *testing.T) {
	tempDir := t.TempDir()
	runPath := filepath.Join(tempDir, "q4-known-portable-winapp-run.json")
	if err := os.WriteFile(runPath, []byte(strings.Replace(q4KnownPortableWinAppRunAcceptanceCLIFixture(currentProjectVersion(t)), `"status": "passed"`, `"status": "failed"`, 1)), 0o644); err != nil {
		t.Fatalf("WriteFile returned error: %v", err)
	}
	var output bytes.Buffer
	if err := run([]string{"q4-known-portable-winapp-run-acceptance-preview", "--q4-known-portable-winapp-run", runPath}, &output); err == nil {
		t.Fatalf("q4 known portable run acceptance must reject failed runs")
	}
	if err := run([]string{"q4-known-portable-winapp-run-acceptance-preview"}, &output); err == nil {
		t.Fatalf("q4 known portable run acceptance must require a run report")
	}
}

func q4KnownPortableWinAppRunAcceptancePuttyFixture(version string) string {
	return `{
  "schema_version": "xnix.scripts.q4_known_portable_winapp_run.v1",
  "request_type": "q4-known-portable-winapp-run",
  "version": "` + version + `",
  "status": "passed",
  "execute": true,
  "app_id": "org.xnix.external.putty",
  "display_name": "PuTTY",
  "known_catalog_app": true,
  "portable_directory_external_app": false,
  "single_file_external_app": true,
  "official_download_required": true,
  "pinned_checksum_required": true,
  "remote_host": "root@q4",
  "remote_materials_root": "/home/xnix-run-materials",
  "remote_paths_exposed_only_for_operator": true,
  "delegated_command": ["ruby", "scripts/q4_putty_external_winapp_smoke.rb", "--execute"],
  "runtime_owned": true,
  "go_runtime_backed": true,
  "kde_policy_owner": false,
  "q4_download_required": true,
  "q4_extract_required": false,
  "q4_compile_required": true,
  "q4_execution_required": true,
  "host_compilation_avoided": true,
  "host_download_avoided": true,
  "full_smoke_required": false,
  "host_root_modified": false,
  "privileged_container_required": false,
  "host_networking_required": false,
  "docker_socket_mounted": false,
  "broad_host_mount_required": false,
  "delegated_status": "passed",
  "delegated_artifact_fetch_count": 13,
  "official_archive_checksum_verified": false,
  "official_executable_checksum_verified": true,
  "known_portable_bundle_imported": false,
  "staged_external_winapp_acceptance_request_type": "q4-staged-external-winapp-acceptance-preview",
  "staged_external_winapp_acceptance_ready": true,
  "external_file_open_requested": false,
  "external_file_bridge_ready": false,
  "windows_process_file_argument_window_observed": false,
  "window_observed": true,
  "x_window_observed": true,
  "runtime_accepted_chain_verified": true,
  "accepted_application_detail_state": "runtime-accepted-real-app-run",
  "kde_accepted_page_state": "runtime-accepted-real-app-run",
  "operator_run_ready": true
}`
}

func q4KnownPortableWinAppRunAcceptanceCLIFixture(version string) string {
	return `{
  "schema_version": "xnix.scripts.q4_known_portable_winapp_run.v1",
  "request_type": "q4-known-portable-winapp-run",
  "version": "` + version + `",
  "status": "passed",
  "execute": true,
  "app_id": "org.xnix.external.notepadplusplus",
  "display_name": "Notepad++ Portable",
  "known_catalog_app": true,
  "portable_directory_external_app": true,
  "official_download_required": true,
  "pinned_checksum_required": true,
  "remote_host": "root@q4",
  "remote_materials_root": "/home/xnix-run-materials",
  "remote_source_root": "/home/xnix-build/xnix-q4-known-portable-winapp-run-test",
  "remote_build_root": "/home/xnix-build-cache",
  "remote_paths_exposed_only_for_operator": true,
  "delegated_command": ["ruby", "scripts/q4_notepadpp_portable_winapp_smoke.rb", "--execute"],
  "runtime_owned": true,
  "go_runtime_backed": true,
  "kde_policy_owner": false,
  "q4_download_required": true,
  "q4_extract_required": true,
  "q4_compile_required": true,
  "q4_execution_required": true,
  "host_compilation_avoided": true,
  "host_download_avoided": true,
  "full_smoke_required": false,
  "host_root_modified": false,
  "privileged_container_required": false,
  "host_networking_required": false,
  "docker_socket_mounted": false,
  "broad_host_mount_required": false,
  "delegated_status": "passed",
  "delegated_artifact_fetch_count": 13,
  "official_archive_checksum_verified": true,
  "known_portable_bundle_imported": true,
  "known_portable_bundle_stage_launch_status": "passed",
  "known_portable_bundle_acceptance_request_type": "q4-known-portable-bundle-winapp-acceptance-preview",
  "known_portable_bundle_acceptance_ready": true,
  "runtime_accepted_chain_verified": true,
  "accepted_application_detail_state": "runtime-accepted-real-app-run",
  "kde_accepted_page_state": "runtime-accepted-real-app-run",
  "operator_run_ready": true
}`
}
