package main

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestExternalWinAppApplicationDetailPreviewCommandConsumesCompatibilityBundle(t *testing.T) {
	tempDir := t.TempDir()
	bundlePath := filepath.Join(tempDir, "compatibility-evidence-bundle.json")
	outputPath := filepath.Join(tempDir, "detail", "external-winapp-detail.json")
	writeTextFile(t, bundlePath, externalWinAppApplicationDetailBundleFixture())

	var output bytes.Buffer
	if err := run([]string{
		"external-winapp-application-detail-preview",
		"--compatibility-evidence-bundle", bundlePath,
		"--output", outputPath,
	}, &output); err != nil {
		t.Fatalf("run returned error: %v", err)
	}
	written, err := os.ReadFile(outputPath)
	if err != nil {
		t.Fatalf("ReadFile output returned error: %v", err)
	}
	if string(written) != output.String() {
		t.Fatalf("written detail must match stdout\nstdout=%s\nwritten=%s", output.String(), string(written))
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["schema_version"] != "xnix.runtime.external_winapp_application_detail.v1" ||
		payload["request_type"] != "external-winapp-application-detail-preview" ||
		payload["detail_type"] != "runtime-owned-existing-windows-app-detail" ||
		payload["runtime_method"] != "PreviewExternalWinAppApplicationDetail" ||
		payload["read_method"] != "GetExternalWinAppApplicationDetail" ||
		payload["desktop"] != "KDE Plasma" ||
		payload["application_id"] != "org.xnix.external.desktop-notepad-file-argument" ||
		payload["display_name"] != "External Desktop Notepad File Argument" ||
		payload["compatibility_state"] != "real-app-run-verified" ||
		payload["compatibility_label"] != "Verified real app run" ||
		payload["primary_status_tone"] != "success" ||
		payload["existing_windows_app_verified"] != true ||
		payload["real_windows_app_run_verified"] != true ||
		payload["file_open_verified"] != true ||
		payload["runtime_gui_evidence_verified"] != true ||
		payload["desktop_evidence_verified"] != true ||
		payload["kde_page_evidence_verified"] != true ||
		payload["evidence_bundle_consumed"] != true ||
		payload["evidence_artifact_count"] != float64(4) ||
		payload["ai_compatibility_runtime_owner"] != true ||
		payload["runtime_owned"] != true ||
		payload["go_runtime_backed"] != true ||
		payload["kde_policy_owner"] != false ||
		payload["safe_for_kde"] != true ||
		payload["safe_for_ai_diagnostics"] != true ||
		payload["launch_enabled"] != false ||
		payload["backend_launch_enabled"] != false ||
		payload["action_execution_enabled"] != false ||
		payload["backend_process_started"] != false ||
		payload["backend_process_evidence_recorded"] != true ||
		payload["backend_details_exposed"] != false ||
		payload["raw_paths_exposed"] != false ||
		payload["raw_launcher_output_exposed"] != false ||
		payload["host_root_modified"] != false ||
		payload["docker_socket_mounted"] != false ||
		payload["broad_host_mount_required"] != false {
		t.Fatalf("unexpected external Windows app detail payload: %#v", payload)
	}
	signals := payload["evidence_signals"].([]any)
	if len(signals) != 4 {
		t.Fatalf("unexpected evidence signals: %#v", signals)
	}
	cards := payload["review_cards"].([]any)
	if len(cards) != 3 {
		t.Fatalf("unexpected review cards: %#v", cards)
	}
	primary := payload["primary_action"].(map[string]any)
	if primary["id"] != "review-compatibility-evidence" ||
		primary["enabled"] != true ||
		primary["runtime_owned"] != true ||
		primary["kde_policy_owner"] != false ||
		primary["backend_details_exposed"] != false {
		t.Fatalf("unexpected primary action: %#v", primary)
	}
	if strings.Contains(output.String(), bundlePath) ||
		strings.Contains(output.String(), "docker run") ||
		strings.Contains(output.String(), "/var/run/docker.sock") ||
		strings.Contains(output.String(), ".exe") ||
		strings.Contains(strings.ToLower(output.String()), "wine ") {
		t.Fatalf("external Windows app detail exposed unsafe details: %s", output.String())
	}

	var kdePageOutput bytes.Buffer
	kdePageOutputPath := filepath.Join(tempDir, "kde", "external-winapp-page-from-detail.json")
	if err := run([]string{
		"kde-center-page-preview",
		"--external-app-application-detail", outputPath,
		"--decision", "approved",
		"--output", kdePageOutputPath,
	}, &kdePageOutput); err != nil {
		t.Fatalf("kde-center-page-preview returned error: %v", err)
	}
	writtenKDEPage, err := os.ReadFile(kdePageOutputPath)
	if err != nil {
		t.Fatalf("ReadFile KDE page output returned error: %v", err)
	}
	if string(writtenKDEPage) != kdePageOutput.String() {
		t.Fatalf("written KDE page must match stdout\nstdout=%s\nwritten=%s", kdePageOutput.String(), string(writtenKDEPage))
	}
	var kdePage map[string]any
	if err := json.Unmarshal(kdePageOutput.Bytes(), &kdePage); err != nil {
		t.Fatalf("Unmarshal KDE page output returned error: %v", err)
	}
	if kdePage["request_type"] != "kde-center-page-preview" ||
		kdePage["application_id"] != "org.xnix.external.desktop-notepad-file-argument" ||
		kdePage["application_name"] != "External Desktop Notepad File Argument" ||
		kdePage["external_winapp_application_detail_consumed"] != true ||
		kdePage["external_winapp_application_detail_count"] != float64(1) ||
		kdePage["runtime_owned"] != true ||
		kdePage["go_runtime_backed"] != true ||
		kdePage["kde_policy_owner"] != false ||
		kdePage["safe_for_ai_diagnostics"] != true ||
		kdePage["launch_enabled"] != false ||
		kdePage["backend_process_started"] != false ||
		kdePage["backend_details_exposed"] != false ||
		kdePage["host_root_modified"] != false {
		t.Fatalf("unexpected KDE page from external app detail: %#v", kdePage)
	}
	header := kdePage["header"].(map[string]any)
	if header["badge"] != "Verified real app run" ||
		header["badge_tone"] != "success" ||
		header["backend_details_exposed"] != false {
		t.Fatalf("unexpected KDE page header from external app detail: %#v", header)
	}
	applicationSummary := kdePage["application_summary"].(map[string]any)
	if applicationSummary["compatibility_state"] != "real-app-run-verified" ||
		applicationSummary["compatibility_label"] != "Verified real app run" ||
		applicationSummary["diagnostics_state"] != "real-app-run-verified" ||
		applicationSummary["backend_launch_enabled"] != false ||
		applicationSummary["action_execution_enabled"] != false ||
		applicationSummary["settings_persistence_enabled"] != false ||
		applicationSummary["backend_details_exposed"] != false ||
		applicationSummary["host_root_modified"] != false {
		t.Fatalf("unexpected KDE application summary from external app detail: %#v", applicationSummary)
	}
	detailCards := kdePage["external_winapp_application_detail_cards"].([]any)
	if len(detailCards) != 1 {
		t.Fatalf("unexpected KDE external app detail cards: %#v", detailCards)
	}
	detailCard := detailCards[0].(map[string]any)
	if detailCard["app_id"] != kdePage["application_id"] ||
		detailCard["real_windows_app_run_verified"] != true ||
		detailCard["file_open_verified"] != true ||
		detailCard["runtime_gui_evidence_verified"] != true ||
		detailCard["desktop_evidence_verified"] != true ||
		detailCard["kde_page_evidence_verified"] != true ||
		detailCard["evidence_signal_count"] != float64(4) ||
		detailCard["primary_action_id"] != "review-compatibility-evidence" ||
		detailCard["primary_action_enabled"] != true ||
		detailCard["runtime_owned"] != true ||
		detailCard["go_runtime_backed"] != true ||
		detailCard["kde_policy_owner"] != false ||
		detailCard["safe_for_kde"] != true ||
		detailCard["safe_for_ai_diagnostics"] != true ||
		detailCard["launch_enabled"] != false ||
		detailCard["backend_launch_enabled"] != false ||
		detailCard["backend_details_exposed"] != false ||
		detailCard["raw_paths_exposed"] != false ||
		detailCard["host_root_modified"] != false {
		t.Fatalf("unexpected KDE external app detail card: %#v", detailCard)
	}
	if strings.Contains(kdePageOutput.String(), outputPath) ||
		strings.Contains(kdePageOutput.String(), kdePageOutputPath) ||
		strings.Contains(kdePageOutput.String(), bundlePath) ||
		strings.Contains(kdePageOutput.String(), "docker run") ||
		strings.Contains(kdePageOutput.String(), "/var/run/docker.sock") ||
		strings.Contains(kdePageOutput.String(), ".exe") ||
		strings.Contains(strings.ToLower(kdePageOutput.String()), "wine ") {
		t.Fatalf("KDE page from external Windows app detail exposed unsafe details: %s", kdePageOutput.String())
	}
}

func externalWinAppApplicationDetailBundleFixture() string {
	return `{
  "version": "0.2.640-test",
  "schema_version": "xnix.runtime.external_winapp_compatibility_evidence_bundle.v1",
  "request_type": "external-winapp-compatibility-evidence-bundle-preview",
  "bundle_type": "runtime-owned-external-windows-app-compatibility-evidence",
  "desktop": "KDE Plasma",
  "application_id": "org.xnix.external.desktop-notepad-file-argument",
  "display_name": "External Desktop Notepad File Argument",
  "app_version": "0.2.640-test",
  "evidence_artifact_count": 4,
  "one_shot_runtime_launch_verified": true,
  "desktop_launch_packet_verified": true,
  "runtime_gui_evidence_packet_verified": true,
  "kde_external_app_page_verified": true,
  "real_windows_app_run_verified": true,
  "external_app_import_record_consumed": true,
  "external_app_handle_consumed": true,
  "imported_artifact_digest_verified": true,
  "external_file_bridge_ready": true,
  "external_file_open_requested": true,
  "external_desktop_argument_count": 1,
  "window_observed": true,
  "x_window_observed": true,
  "container_runtime_used": true,
  "container_network_mode": "none",
  "container_host_mount_count": 0,
  "known_app_gui_evidence_count": 1,
  "known_app_gui_evidence_verified_count": 1,
  "kde_page_known_app_gui_evidence_count": 1,
  "runtime_owned": true,
  "go_runtime_backed": true,
  "kde_policy_owner": false,
  "safe_for_kde": true,
  "safe_for_ai_diagnostics": true,
  "desktop_launch_enabled": false,
  "backend_launch_enabled": false,
  "backend_process_started": true,
  "action_execution_enabled": false,
  "backend_details_exposed": false,
  "raw_paths_exposed": false,
  "raw_launcher_output_exposed": false,
  "host_root_modified": false,
  "privileged_container_required": false,
  "host_networking_required": false,
  "docker_socket_mounted": false,
  "broad_host_mount_required": false,
  "desktop_safe_summary": "External Windows app compatibility evidence is Runtime-owned and verified through a real isolated GUI run."
}`
}
