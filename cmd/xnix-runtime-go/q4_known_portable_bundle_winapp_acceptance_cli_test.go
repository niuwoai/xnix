package main

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestQ4KnownPortableBundleWinAppAcceptancePreviewCommandConsumesNotepadPPReport(t *testing.T) {
	tempDir := t.TempDir()
	reportPath := filepath.Join(tempDir, "q4-notepadpp-portable-winapp-smoke.json")
	outputPath := filepath.Join(tempDir, "acceptance", "q4-known-portable-bundle-winapp-acceptance.json")
	if err := os.WriteFile(reportPath, []byte(q4KnownPortableBundleWinAppAcceptanceCLIFixture(currentProjectVersion(t))), 0o600); err != nil {
		t.Fatalf("WriteFile report returned error: %v", err)
	}

	var output bytes.Buffer
	if err := run([]string{
		"q4-known-portable-bundle-winapp-acceptance-preview",
		"--q4-notepadpp-portable-winapp-smoke", reportPath,
		"--output", outputPath,
	}, &output); err != nil {
		t.Fatalf("run returned error: %v", err)
	}
	written, err := os.ReadFile(outputPath)
	if err != nil {
		t.Fatalf("ReadFile output returned error: %v", err)
	}
	if string(written) != output.String() {
		t.Fatalf("written acceptance must match stdout\nstdout=%s\nwritten=%s", output.String(), string(written))
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["schema_version"] != "xnix.runtime.q4_known_portable_bundle_winapp_acceptance.v1" ||
		payload["request_type"] != "q4-known-portable-bundle-winapp-acceptance-preview" ||
		payload["acceptance_type"] != "q4-known-portable-bundle-real-winapp-acceptance" ||
		payload["smoke_report_consumed"] != true ||
		payload["smoke_report_path_exposed"] != false ||
		payload["remote_host_exposed"] != false ||
		payload["delegated_command_exposed"] != false ||
		payload["raw_path_exposed"] != false ||
		payload["app_id"] != "org.xnix.external.notepadplusplus" ||
		payload["display_name"] != "Notepad++ Portable" ||
		payload["real_third_party_windows_app"] != true ||
		payload["portable_directory_app"] != true ||
		payload["host_compilation_avoided"] != true ||
		payload["host_download_avoided"] != true ||
		payload["official_archive_checksum_verified"] != true ||
		payload["known_bundle_import_verified"] != true ||
		payload["known_bundle_stage_launch_verified"] != true ||
		payload["known_bundle_gui_evidence_packet_verified"] != true ||
		payload["known_bundle_kde_page_verified"] != true ||
		payload["known_bundle_compatibility_bundle_verified"] != true ||
		payload["known_bundle_application_detail_verified"] != true ||
		payload["known_bundle_kde_page_from_detail_verified"] != true ||
		payload["launch_source_request_type"] != "windows-known-app-bundle-stage-and-launch" ||
		payload["known_portable_bundle_stage_launch_consumed"] != true ||
		payload["runtime_accepted_chain_verified"] != true ||
		payload["runtime_accepted_application_detail_state"] != "runtime-accepted-real-app-run" ||
		payload["runtime_accepted_kde_page_state"] != "runtime-accepted-real-app-run" ||
		payload["acceptance_ready"] != true ||
		payload["host_root_modified"] != false ||
		payload["docker_socket_mounted"] != false ||
		payload["broad_host_mount_required"] != false ||
		payload["backend_details_exposed"] != false {
		t.Fatalf("unexpected q4 known portable bundle Windows app acceptance: %#v", payload)
	}
	lower := strings.ToLower(output.String())
	for _, forbidden := range []string{strings.ToLower(reportPath), strings.ToLower(outputPath), "root@q4", "/home/xnix-", "notepad++.exe", "docker run", "/var/run/docker.sock", " wine "} {
		if strings.Contains(lower, forbidden) {
			t.Fatalf("q4 known portable bundle acceptance exposed forbidden term %q: %s", forbidden, output.String())
		}
	}
}

func TestQ4KnownPortableBundleWinAppAcceptancePreviewCommandRequiresPassedReport(t *testing.T) {
	var output bytes.Buffer
	if err := run([]string{"q4-known-portable-bundle-winapp-acceptance-preview"}, &output); err == nil {
		t.Fatalf("q4-known-portable-bundle-winapp-acceptance-preview must require a smoke report")
	}
	if err := run([]string{"q4-known-portable-bundle-winapp-acceptance-preview", "--q4-notepadpp-portable-winapp-smoke", "missing.json", "extra"}, &output); err == nil {
		t.Fatalf("q4-known-portable-bundle-winapp-acceptance-preview must reject positional arguments")
	}
}

func q4KnownPortableBundleWinAppAcceptanceCLIFixture(version string) string {
	return strings.ReplaceAll(`{
  "schema_version": "xnix.scripts.q4_notepadpp_portable_winapp_smoke.v1",
  "request_type": "q4-notepadpp-portable-winapp-smoke",
  "version": "VERSION_PLACEHOLDER",
  "status": "passed",
  "execute": true,
  "source_kind": "official-notepad-plus-plus-github-release",
  "app_id": "org.xnix.external.notepadplusplus",
  "display_name": "Notepad++ Portable",
  "real_third_party_windows_app": true,
  "single_file_windows_app": false,
  "portable_directory_bundle_import_required": true,
  "q4_compile_required": true,
  "q4_download_required": true,
  "q4_extract_required": true,
  "host_compilation_avoided": true,
  "host_download_avoided": true,
  "notepadpp_sha256_verified": true,
  "notepadpp_extracted": true,
  "known_portable_bundle_import_status": "passed",
  "known_portable_bundle_import_recorded": true,
  "known_portable_bundle_archive_verified": true,
  "known_portable_bundle_checksum_verified": true,
  "known_portable_bundle_stage_launch_status": "passed",
  "known_portable_bundle_stage_launch_record_first": true,
  "known_portable_bundle_stage_launch_existing_import_record_consumed": true,
  "known_portable_bundle_stage_launch_application_workspace_copied": true,
  "known_portable_bundle_stage_launch_windows_process_file_argument_window_observed": true,
  "known_portable_bundle_gui_evidence_packet_request_type": "real-winapp-gui-evidence-packet-preview",
  "known_portable_bundle_gui_evidence_packet_external_app_run_record_consumed": true,
  "known_portable_bundle_gui_evidence_packet_imported_artifact_digest_verified": true,
  "known_portable_bundle_gui_evidence_packet_artifact_fetched": true,
  "known_portable_bundle_kde_page_request_type": "kde-center-page-preview",
  "known_portable_bundle_kde_page_artifact_fetched": true,
  "known_portable_bundle_compatibility_bundle_request_type": "external-winapp-compatibility-evidence-bundle-preview",
  "known_portable_bundle_compatibility_bundle_launch_source_request_type": "windows-known-app-bundle-stage-and-launch",
  "known_portable_bundle_compatibility_bundle_stage_launch_consumed": true,
  "known_portable_bundle_compatibility_bundle_runtime_gui_evidence_packet_verified": true,
  "known_portable_bundle_compatibility_bundle_kde_external_app_page_verified": true,
  "known_portable_bundle_compatibility_bundle_real_windows_app_run_verified": true,
  "known_portable_bundle_compatibility_bundle_artifact_fetched": true,
  "known_portable_bundle_application_detail_request_type": "external-winapp-application-detail-preview",
  "known_portable_bundle_application_detail_launch_source_request_type": "windows-known-app-bundle-stage-and-launch",
  "known_portable_bundle_application_detail_stage_launch_consumed": true,
  "known_portable_bundle_application_detail_real_windows_app_run_verified": true,
  "known_portable_bundle_application_detail_file_open_verified": true,
  "known_portable_bundle_application_detail_artifact_fetched": true,
  "known_portable_bundle_kde_page_from_detail_request_type": "kde-center-page-preview",
  "known_portable_bundle_kde_page_from_detail_launch_source_request_type": "windows-known-app-bundle-stage-and-launch",
  "known_portable_bundle_kde_page_from_detail_stage_launch_consumed": true,
  "known_portable_bundle_kde_page_from_detail_artifact_fetched": true,
  "acceptance_ready": true,
  "accepted_application_detail_state": "runtime-accepted-real-app-run",
  "kde_accepted_page_state": "runtime-accepted-real-app-run",
  "artifact_fetch_count": 13,
  "host_root_modified": false,
  "privileged_container_required": false,
  "host_networking_required": false,
  "docker_socket_mounted": false,
  "broad_host_mount_required": false
}`, "VERSION_PLACEHOLDER", version)
}
