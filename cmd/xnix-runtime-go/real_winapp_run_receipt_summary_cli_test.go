package main

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestRealWinAppRunReceiptSummaryPreviewCommandConsumesRemoteQ4Smoke(t *testing.T) {
	tempDir := t.TempDir()
	reportPath := filepath.Join(tempDir, "q4-remote-smoke.json")
	if err := os.WriteFile(reportPath, []byte(realWinAppRunReceiptSummaryCLIFixture(currentProjectVersion(t))), 0o600); err != nil {
		t.Fatalf("WriteFile report returned error: %v", err)
	}

	var output bytes.Buffer
	if err := run([]string{
		"real-winapp-run-receipt-summary-preview",
		"--remote-smoke-report", reportPath,
	}, &output); err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["schema_version"] != "xnix.runtime.real_winapp_run_receipt_summary.v1" ||
		payload["request_type"] != "real-winapp-run-receipt-summary-preview" ||
		payload["receipt_type"] != "real-windows-app-run-receipt-summary" ||
		payload["report_consumed"] != true ||
		payload["report_path_exposed"] != false ||
		payload["remote_host_exposed"] != false ||
		payload["app_id"] != "org.xnix.sample.notepad" ||
		payload["display_name"] != "Sample Notepad" ||
		payload["gui_app_name"] != "known-gui-app" ||
		payload["execution_host_class"] != "q4-remote-validation-host" ||
		payload["backend_class"] != "managed-guest-gui" ||
		payload["run_passed"] != true ||
		payload["real_execution_observed"] != true ||
		payload["window_observed"] != true ||
		payload["file_open_verified"] != true ||
		payload["document_content_marker_observation_required"] != true ||
		payload["document_content_marker_observed"] != true ||
		payload["owner_controlled_launch_verified"] != true ||
		payload["owner_file_open_entrypoint_invoked"] != true ||
		payload["runtime_evidence_consumed"] != true ||
		payload["kde_page_evidence_consumed"] != true ||
		payload["kde_action_evidence_consumed"] != true ||
		payload["receipt_ready"] != true ||
		payload["host_root_modified"] != false ||
		payload["privileged_container_required"] != false ||
		payload["host_networking_required"] != false ||
		payload["docker_socket_mounted"] != false ||
		payload["broad_host_mount_required"] != false ||
		payload["backend_details_exposed"] != false ||
		payload["raw_launcher_output_exposed"] != false {
		t.Fatalf("unexpected real run receipt summary CLI payload: %#v", payload)
	}
	lower := strings.ToLower(output.String())
	for _, forbidden := range []string{strings.ToLower(reportPath), "root@q4", "/home/xnix-", "notepad.exe", "qemu-system", "/var/run/docker.sock"} {
		if strings.Contains(lower, forbidden) {
			t.Fatalf("real run receipt summary CLI exposed forbidden term %q: %s", forbidden, output.String())
		}
	}
}

func TestRealWinAppRunReceiptSummaryFeedsCompatibilityAndKDECenterPreviews(t *testing.T) {
	tempDir := t.TempDir()
	reportPath := filepath.Join(tempDir, "q4-remote-smoke.json")
	if err := os.WriteFile(reportPath, []byte(realWinAppRunReceiptSummaryCLIFixture(currentProjectVersion(t))), 0o600); err != nil {
		t.Fatalf("WriteFile report returned error: %v", err)
	}

	var summaryOutput bytes.Buffer
	if err := run([]string{
		"real-winapp-run-receipt-summary-preview",
		"--remote-smoke-report", reportPath,
	}, &summaryOutput); err != nil {
		t.Fatalf("run summary returned error: %v", err)
	}
	summaryPath := filepath.Join(tempDir, "real-run-receipt-summary.json")
	if err := os.WriteFile(summaryPath, summaryOutput.Bytes(), 0o600); err != nil {
		t.Fatalf("WriteFile summary returned error: %v", err)
	}

	var centerOutput bytes.Buffer
	if err := run([]string{
		"compatibility-center-preview",
		"--registry", "../../runtime/recipes/registry.json",
		"--known-app-evidence-file", summaryPath,
	}, &centerOutput); err != nil {
		t.Fatalf("run compatibility center returned error: %v", err)
	}
	var centerPayload map[string]any
	if err := json.Unmarshal(centerOutput.Bytes(), &centerPayload); err != nil {
		t.Fatalf("Unmarshal center returned error: %v", err)
	}
	if centerPayload["known_app_smoke_evidence_count"] != float64(1) ||
		centerPayload["known_app_smoke_passed_count"] != float64(1) ||
		centerPayload["backend_launch_enabled"] != false ||
		centerPayload["host_root_modified"] != false ||
		centerPayload["backend_details_exposed"] != false {
		t.Fatalf("unexpected real run receipt center payload: %#v", centerPayload)
	}
	centerEvidence := centerPayload["known_app_smoke_evidence"].([]any)[0].(map[string]any)
	if centerEvidence["app_id"] != "org.xnix.sample.notepad" ||
		centerEvidence["display_name"] != "Sample Notepad" ||
		centerEvidence["evidence_kind"] != "known-application-gui-smoke" ||
		centerEvidence["evidence_source"] != "wine-guest-gui-smoke" ||
		centerEvidence["compatibility_state"] != "owner-controlled-gui-qemu-wine-verified" ||
		centerEvidence["center_card_state"] != "validated-owner-controlled-gui-runtime-run" ||
		centerEvidence["owner_file_open_verified"] != true ||
		centerEvidence["owner_file_open_entrypoint_invoked"] != true ||
		centerEvidence["owner_delegated_file_argument_count"] != float64(1) ||
		centerEvidence["owner_delegated_file_argument_copied_count"] != float64(1) ||
		centerEvidence["owner_delegated_file_arguments_passed"] != true ||
		centerEvidence["owner_delegated_file_argument_winepath_translated"] != true ||
		centerEvidence["owner_delegated_raw_file_argument_path_exposed"] != false ||
		centerEvidence["runtime_owned"] != true ||
		centerEvidence["desktop_launch_enabled"] != false ||
		centerEvidence["backend_launch_enabled"] != false ||
		centerEvidence["host_root_modified"] != false ||
		centerEvidence["backend_details_exposed"] != false {
		t.Fatalf("unexpected real run receipt center evidence: %#v", centerEvidence)
	}

	var kdeOutput bytes.Buffer
	if err := run([]string{
		"kde-center-page-preview",
		"--registry", "../../runtime/recipes/registry.json",
		"--app", "org.xnix.sample.notepad",
		"--decision", "approved",
		"--known-app-evidence-file", summaryPath,
	}, &kdeOutput); err != nil {
		t.Fatalf("run KDE center returned error: %v", err)
	}
	var kdePayload map[string]any
	if err := json.Unmarshal(kdeOutput.Bytes(), &kdePayload); err != nil {
		t.Fatalf("Unmarshal KDE returned error: %v", err)
	}
	if kdePayload["known_app_gui_evidence_count"] != float64(1) ||
		kdePayload["known_app_owner_controlled_gui_evidence_count"] != float64(1) ||
		kdePayload["known_app_owner_file_open_verified_count"] != float64(1) ||
		kdePayload["known_app_owner_file_open_entrypoint_count"] != float64(1) ||
		kdePayload["launch_enabled"] != false ||
		kdePayload["execution_started"] != false ||
		kdePayload["backend_process_started"] != false ||
		kdePayload["host_root_modified"] != false ||
		kdePayload["backend_details_exposed"] != false {
		t.Fatalf("unexpected real run receipt KDE payload: %#v", kdePayload)
	}
	kdeCard := kdePayload["known_app_gui_evidence_cards"].([]any)[0].(map[string]any)
	if kdeCard["app_id"] != "org.xnix.sample.notepad" ||
		kdeCard["owner_file_open_verified"] != true ||
		kdeCard["owner_file_open_entrypoint_invoked"] != true ||
		kdeCard["owner_delegated_file_argument_count"] != float64(1) ||
		kdeCard["owner_delegated_file_argument_copied_count"] != float64(1) ||
		kdeCard["owner_delegated_file_arguments_passed"] != true ||
		kdeCard["owner_delegated_file_argument_winepath_translated"] != true ||
		kdeCard["owner_delegated_raw_file_argument_path_exposed"] != false ||
		kdeCard["desktop_launch_enabled"] != false ||
		kdeCard["backend_launch_enabled"] != false ||
		kdeCard["host_root_modified"] != false ||
		kdeCard["backend_details_exposed"] != false {
		t.Fatalf("unexpected real run receipt KDE card: %#v", kdeCard)
	}
	for _, output := range []string{centerOutput.String(), kdeOutput.String()} {
		lower := strings.ToLower(output)
		for _, forbidden := range []string{strings.ToLower(reportPath), strings.ToLower(summaryPath), "root@q4", "/home/xnix-", "notepad.exe", "qemu-system", "/var/run/docker.sock"} {
			if strings.Contains(lower, forbidden) {
				t.Fatalf("real run receipt center output exposed forbidden term %q: %s", forbidden, output)
			}
		}
	}
}

func TestRealWinAppRunReceiptSummaryPreviewCommandRequiresReport(t *testing.T) {
	var output bytes.Buffer
	if err := run([]string{"real-winapp-run-receipt-summary-preview"}, &output); err == nil {
		t.Fatalf("real-winapp-run-receipt-summary-preview must require --remote-smoke-report")
	}
	if err := run([]string{"real-winapp-run-receipt-summary-preview", "--remote-smoke-report", "missing.json", "extra"}, &output); err == nil {
		t.Fatalf("real-winapp-run-receipt-summary-preview must reject positional arguments")
	}
}

func realWinAppRunReceiptSummaryCLIFixture(version string) string {
	return strings.ReplaceAll(`{
  "schema_version": "xnix.scripts.remote_wine_guest_gui_smoke.execute_result.v1",
  "request_type": "remote-wine-guest-gui-smoke",
  "version": "VERSION_PLACEHOLDER",
  "status": "passed",
  "remote_host": "root@q4",
  "smoke_status": "passed",
  "launch_mode": "owner-controlled-launch",
  "known_app_id": "org.xnix.sample.notepad",
  "known_app_name": "Sample Notepad",
  "known_app_version": "VERSION_PLACEHOLDER",
  "gui_app_name": "notepad.exe",
  "file_argument_count": 1,
  "file_argument_copied_count": 1,
  "file_arguments_passed": true,
  "file_argument_winepath_translated": true,
  "file_argument_winepath_translated_count": 1,
  "raw_file_argument_path_exposed": false,
  "window_match": "sample-document.txt",
  "window_match_observed": true,
  "document_content_marker_observation_required": true,
  "document_content_marker_observed": true,
  "x_window_observed": true,
  "runtime_evidence_report_consumed": true,
  "runtime_evidence_window_observed": true,
  "runtime_evidence_owner_file_open_verified": true,
  "runtime_evidence_owner_file_open_entrypoint_invoked": true,
  "kde_page_output_written": true,
  "kde_page_known_app_gui_evidence_count": 1,
  "kde_action_output_written": true,
  "kde_action_owner_file_open_verified": true,
  "kde_action_owner_file_open_entrypoint_invoked": true,
  "owner_controlled_launch_requested": true,
  "owner_evidence_handoff_ready": true,
  "owner_managed_launcher_invoked": true,
  "owner_file_open_entrypoint_invoked": true,
  "owner_delegated_smoke_passed": true,
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
