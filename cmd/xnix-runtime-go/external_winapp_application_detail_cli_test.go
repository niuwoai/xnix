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

func TestExternalWinAppApplicationDetailPreviewCommandPreservesKnownPortableBundleProvenance(t *testing.T) {
	tempDir := t.TempDir()
	bundlePath := filepath.Join(tempDir, "known-portable-bundle-compatibility-evidence-bundle.json")
	outputPath := filepath.Join(tempDir, "detail", "known-portable-bundle-detail.json")
	writeTextFile(t, bundlePath, knownPortableExternalWinAppApplicationDetailBundleFixture())

	var output bytes.Buffer
	if err := run([]string{
		"external-winapp-application-detail-preview",
		"--compatibility-evidence-bundle", bundlePath,
		"--output", outputPath,
	}, &output); err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["schema_version"] != "xnix.runtime.external_winapp_application_detail.v1" ||
		payload["request_type"] != "external-winapp-application-detail-preview" ||
		payload["launch_source_request_type"] != "windows-known-app-bundle-stage-and-launch" ||
		payload["known_portable_bundle_stage_launch_consumed"] != true ||
		payload["known_portable_bundle_stage_launch_verified"] != true ||
		payload["compatibility_state"] != "real-app-run-verified" ||
		payload["real_windows_app_run_verified"] != true ||
		payload["file_open_verified"] != true ||
		payload["runtime_gui_evidence_verified"] != true ||
		payload["desktop_evidence_verified"] != true ||
		payload["kde_page_evidence_verified"] != true ||
		payload["safe_for_kde"] != true ||
		payload["runtime_owned"] != true ||
		payload["go_runtime_backed"] != true ||
		payload["backend_details_exposed"] != false ||
		payload["raw_paths_exposed"] != false ||
		payload["host_root_modified"] != false {
		t.Fatalf("unexpected known portable bundle external Windows app detail payload: %#v", payload)
	}
	signals := payload["evidence_signals"].([]any)
	if len(signals) != 5 {
		t.Fatalf("known portable bundle detail must expose an extra signal: %#v", signals)
	}
	cards := payload["review_cards"].([]any)
	if len(cards) != 4 {
		t.Fatalf("known portable bundle detail must expose an extra review card: %#v", cards)
	}
	foundSignal := false
	for _, item := range signals {
		signal := item.(map[string]any)
		if signal["id"] == "known-portable-bundle-stage-launch" && signal["verified"] == true {
			foundSignal = true
		}
	}
	if !foundSignal {
		t.Fatalf("known portable bundle detail did not preserve launch signal: %#v", signals)
	}

	var kdePageOutput bytes.Buffer
	if err := run([]string{
		"kde-center-page-preview",
		"--external-app-application-detail", outputPath,
		"--decision", "approved",
	}, &kdePageOutput); err != nil {
		t.Fatalf("kde-center-page-preview returned error: %v", err)
	}
	var kdePage map[string]any
	if err := json.Unmarshal(kdePageOutput.Bytes(), &kdePage); err != nil {
		t.Fatalf("Unmarshal KDE page output returned error: %v", err)
	}
	detailCards := kdePage["external_winapp_application_detail_cards"].([]any)
	if len(detailCards) != 1 {
		t.Fatalf("unexpected KDE external app detail cards: %#v", detailCards)
	}
	detailCard := detailCards[0].(map[string]any)
	if detailCard["launch_source_request_type"] != "windows-known-app-bundle-stage-and-launch" ||
		detailCard["known_portable_bundle_stage_launch_consumed"] != true ||
		detailCard["known_portable_bundle_stage_launch_verified"] != true ||
		detailCard["evidence_signal_count"] != float64(5) ||
		detailCard["evidence_artifact_count"] != float64(4) ||
		detailCard["runtime_owned"] != true ||
		detailCard["go_runtime_backed"] != true ||
		detailCard["kde_policy_owner"] != false ||
		detailCard["safe_for_kde"] != true ||
		detailCard["backend_details_exposed"] != false ||
		detailCard["raw_paths_exposed"] != false ||
		detailCard["host_root_modified"] != false {
		t.Fatalf("unexpected KDE known portable bundle external app detail card: %#v", detailCard)
	}
	for _, rendered := range []string{output.String(), kdePageOutput.String()} {
		if strings.Contains(rendered, bundlePath) ||
			strings.Contains(rendered, "docker run") ||
			strings.Contains(rendered, "/var/run/docker.sock") ||
			strings.Contains(rendered, ".exe") ||
			strings.Contains(strings.ToLower(rendered), "wine ") {
			t.Fatalf("known portable bundle external app detail exposed unsafe details: %s", rendered)
		}
	}
}

func TestExternalWinAppApplicationDetailPreviewCommandConsumesQ4StagedAcceptance(t *testing.T) {
	tempDir := t.TempDir()
	bundlePath := filepath.Join(tempDir, "compatibility-evidence-bundle.json")
	acceptancePath := filepath.Join(tempDir, "q4-staged-external-winapp-acceptance.json")
	outputPath := filepath.Join(tempDir, "detail", "external-winapp-detail.json")
	writeTextFile(t, bundlePath, externalWinAppApplicationDetailBundleFixture())
	writeTextFile(t, acceptancePath, externalWinAppApplicationDetailQ4StagedAcceptanceFixture(currentProjectVersion(t)))

	var output bytes.Buffer
	if err := run([]string{
		"external-winapp-application-detail-preview",
		"--compatibility-evidence-bundle", bundlePath,
		"--q4-staged-external-winapp-acceptance", acceptancePath,
		"--output", outputPath,
	}, &output); err != nil {
		t.Fatalf("run returned error: %v", err)
	}
	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["source"] != "external-winapp-compatibility-evidence-bundle+q4-staged-external-winapp-acceptance" ||
		payload["compatibility_state"] != "runtime-accepted-real-app-run" ||
		payload["compatibility_label"] != "Runtime accepted real app run" ||
		payload["q4_staged_external_winapp_acceptance_consumed"] != true ||
		payload["q4_staged_external_winapp_acceptance_ready"] != true ||
		payload["q4_staged_external_winapp_acceptance_type"] != "staged-external-winapp-desktop-real-run-acceptance" ||
		payload["go_owned_staged_external_winapp_acceptance_verified"] != true ||
		payload["evidence_artifact_count"] != float64(5) {
		t.Fatalf("unexpected acceptance-backed external Windows app detail payload: %#v", payload)
	}
	if len(payload["evidence_signals"].([]any)) != 5 ||
		len(payload["review_cards"].([]any)) != 4 {
		t.Fatalf("acceptance-backed detail must expose an extra signal and review card: %#v", payload)
	}

	var kdePageOutput bytes.Buffer
	if err := run([]string{
		"kde-center-page-preview",
		"--external-app-application-detail", outputPath,
		"--decision", "approved",
	}, &kdePageOutput); err != nil {
		t.Fatalf("kde-center-page-preview returned error: %v", err)
	}
	var kdePage map[string]any
	if err := json.Unmarshal(kdePageOutput.Bytes(), &kdePage); err != nil {
		t.Fatalf("Unmarshal KDE page output returned error: %v", err)
	}
	detailCards := kdePage["external_winapp_application_detail_cards"].([]any)
	if len(detailCards) != 1 {
		t.Fatalf("unexpected KDE external app detail cards: %#v", detailCards)
	}
	detailCard := detailCards[0].(map[string]any)
	if detailCard["compatibility_state"] != "runtime-accepted-real-app-run" ||
		detailCard["compatibility_label"] != "Runtime accepted real app run" ||
		detailCard["q4_staged_external_winapp_acceptance_consumed"] != true ||
		detailCard["q4_staged_external_winapp_acceptance_ready"] != true ||
		detailCard["go_owned_staged_external_winapp_acceptance_verified"] != true ||
		detailCard["evidence_signal_count"] != float64(5) ||
		detailCard["backend_details_exposed"] != false ||
		detailCard["raw_paths_exposed"] != false ||
		detailCard["host_root_modified"] != false {
		t.Fatalf("unexpected KDE acceptance-backed external app detail card: %#v", detailCard)
	}
	header := kdePage["header"].(map[string]any)
	applicationSummary := kdePage["application_summary"].(map[string]any)
	if header["badge"] != "Runtime accepted real app run" ||
		header["badge_tone"] != "success" ||
		applicationSummary["compatibility_state"] != "runtime-accepted-real-app-run" ||
		applicationSummary["compatibility_label"] != "Runtime accepted real app run" ||
		applicationSummary["backend_details_exposed"] != false {
		t.Fatalf("unexpected KDE acceptance-backed top-level state: header=%#v summary=%#v", header, applicationSummary)
	}
	for _, rendered := range []string{output.String(), kdePageOutput.String()} {
		if strings.Contains(rendered, acceptancePath) ||
			strings.Contains(rendered, "root@q4") ||
			strings.Contains(rendered, "/tmp/xnix-") ||
			strings.Contains(rendered, "staged_desktop_external_winapp_smoke") ||
			strings.Contains(rendered, "notepad.exe") {
			t.Fatalf("acceptance-backed external app detail exposed unsafe details: %s", rendered)
		}
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

func knownPortableExternalWinAppApplicationDetailBundleFixture() string {
	return strings.Replace(externalWinAppApplicationDetailBundleFixture(), `"desktop": "KDE Plasma",`, `"source": "windows-known-app-bundle-stage-and-launch+desktop-launch-packet+real-winapp-gui-evidence-packet+kde-center-page",
  "launch_source_request_type": "windows-known-app-bundle-stage-and-launch",
  "desktop": "KDE Plasma",
  "known_portable_bundle_stage_launch_consumed": true,
  "known_portable_bundle_stage_launch_verified": true,`, 1)
}

func externalWinAppApplicationDetailQ4StagedAcceptanceFixture(version string) string {
	return strings.ReplaceAll(`{
  "version": "VERSION_PLACEHOLDER",
  "schema_version": "xnix.runtime.q4_staged_external_winapp_acceptance.v1",
  "request_type": "q4-staged-external-winapp-acceptance-preview",
  "source": "q4-staged-desktop-external-winapp-smoke+go-runtime-acceptance",
  "runtime_method": "PreviewQ4StagedExternalWinAppAcceptance",
  "read_method": "GetQ4StagedExternalWinAppAcceptance",
  "acceptance_type": "staged-external-winapp-desktop-real-run-acceptance",
  "smoke_report_consumed": true,
  "smoke_report_path_exposed": false,
  "output_path_exposed": false,
  "delegated_command_exposed": false,
  "remote_host_exposed": false,
  "app_id": "org.xnix.external.desktop-notepad-file-argument",
  "display_name": "External Desktop Notepad File Argument",
  "q4_compile_required": true,
  "host_compilation_avoided": true,
  "remote_build_completed": true,
  "delegated_smoke_passed": true,
  "desktop_exec_uses_external_app_handle": true,
  "desktop_exec_invocation_exact": true,
  "external_app_desktop_handle_ready": true,
  "external_app_handle_consumed": true,
  "imported_artifact_digest_verified": true,
  "external_file_bridge_ready": true,
  "external_file_bridge_arguments_passed": true,
  "external_file_bridge_winepath_translated": true,
  "windows_process_file_argument_window_observed": true,
  "window_observed": true,
  "x_window_observed": true,
  "one_shot_runtime_launch_executed": true,
  "one_shot_window_observed": true,
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
  "kde_page_from_application_detail_summary_compatibility_state": "real-app-run-verified",
  "kde_page_from_application_detail_summary_diagnostics_state": "real-app-run-verified",
  "kde_page_from_application_detail_summary_backend_launch_enabled": false,
  "kde_page_from_application_detail_summary_action_execution_enabled": false,
  "kde_page_from_application_detail_summary_settings_persistence_enabled": false,
  "artifact_fetch_count": 11,
  "host_root_modified": false,
  "privileged_container_required": false,
  "host_networking_required": false,
  "docker_socket_mounted": false,
  "broad_host_mount_required": false,
  "backend_details_exposed": false,
  "raw_path_exposed": false,
  "acceptance_ready": true,
  "desktop_safe_summary": "A q4 staged external Windows GUI app completed the handle-only desktop launch, file-open, Runtime detail, and KDE detail acceptance lane."
}`, "VERSION_PLACEHOLDER", version)
}
