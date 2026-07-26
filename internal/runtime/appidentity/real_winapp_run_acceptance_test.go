package appidentity

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestPreviewRealWinAppRunAcceptanceConsumesReceiptBackedExecuteResult(t *testing.T) {
	acceptance, err := PreviewRealWinAppRunAcceptanceJSON([]byte(realWinAppRunAcceptanceExecuteResultFixture(currentProjectVersion(t))))
	if err != nil {
		t.Fatalf("PreviewRealWinAppRunAcceptanceJSON returned error: %v", err)
	}
	if acceptance.SchemaVersion != RealWinAppRunAcceptanceSchemaVersion ||
		acceptance.RequestType != RealWinAppRunAcceptanceRequestType ||
		acceptance.AcceptanceType != "real-windows-app-run-acceptance" ||
		acceptance.ExecuteResultConsumed != true ||
		acceptance.ExecuteResultPathExposed != false ||
		acceptance.ReceiptSummaryPathExposed != false ||
		acceptance.CenterProjectionPathExposed != false ||
		acceptance.KDEPageProjectionPathExposed != false ||
		acceptance.RemoteHostExposed != false ||
		acceptance.AppID != "org.xnix.sample.notepad" ||
		acceptance.DisplayName != "Sample Notepad" ||
		acceptance.AppVersion != currentProjectVersion(t) ||
		acceptance.LaunchMode != "owner-controlled-launch" ||
		acceptance.RunPassed != true ||
		acceptance.RealExecutionObserved != true ||
		acceptance.WindowObserved != true ||
		acceptance.WindowMatchObserved != true ||
		acceptance.OwnerControlledFileOpenVerified != true ||
		acceptance.OwnerFileOpenEntrypointInvoked != true ||
		acceptance.RuntimeEvidenceConsumed != true ||
		acceptance.ReceiptSummaryGenerated != true ||
		acceptance.ReceiptSummaryReady != true ||
		acceptance.ReceiptSummaryFileOpenVerified != true ||
		acceptance.CompatibilityCenterProjectionConsumed != true ||
		acceptance.CompatibilityCenterKnownAppSmokeEvidenceCount != 1 ||
		acceptance.KDEPageProjectionConsumed != true ||
		acceptance.KDEPageKnownAppGUIEvidenceCount != 1 ||
		acceptance.KDEPageOwnerFileOpenVerifiedCount != 1 ||
		acceptance.KDEPageOwnerFileOpenEntrypointCount != 1 ||
		acceptance.HostRootModified != false ||
		acceptance.PrivilegedContainerRequired != false ||
		acceptance.HostNetworkingRequired != false ||
		acceptance.DockerSocketMounted != false ||
		acceptance.BroadHostMountRequired != false ||
		acceptance.BackendDetailsExposed != false ||
		acceptance.RawLauncherOutputExposed != false ||
		acceptance.RawWindowEvidenceExposed != false ||
		acceptance.ExecutableNameExposed != false ||
		acceptance.AcceptanceReady != true {
		t.Fatalf("unexpected real run acceptance: %#v", acceptance)
	}
	encoded, err := json.Marshal(acceptance)
	if err != nil {
		t.Fatalf("Marshal returned error: %v", err)
	}
	text := strings.ToLower(string(encoded))
	for _, forbidden := range []string{"/home/xnix-", "root@q4", "notepad.exe", "qemu-system", "/var/run/docker.sock", "sample-document.txt - notepad"} {
		if strings.Contains(text, forbidden) {
			t.Fatalf("real run acceptance exposed forbidden term %q: %s", forbidden, text)
		}
	}
}

func TestPreviewRealWinAppRunAcceptanceRejectsIncompleteReceiptProjection(t *testing.T) {
	fixture := realWinAppRunAcceptanceExecuteResultFixture(currentProjectVersion(t))
	missingReceipt := strings.Replace(fixture, `"real_run_receipt_summary_ready": true`, `"real_run_receipt_summary_ready": false`, 1)
	if _, err := PreviewRealWinAppRunAcceptanceJSON([]byte(missingReceipt)); err == nil {
		t.Fatalf("acceptance must reject missing receipt readiness")
	}
	missingKDE := strings.Replace(fixture, `"real_run_receipt_summary_kde_page_owner_file_open_entrypoint_count": 1`, `"real_run_receipt_summary_kde_page_owner_file_open_entrypoint_count": 0`, 1)
	if _, err := PreviewRealWinAppRunAcceptanceJSON([]byte(missingKDE)); err == nil {
		t.Fatalf("acceptance must reject missing KDE receipt projection evidence")
	}
	unsafe := strings.Replace(fixture, `"host_root_modified": false`, `"host_root_modified": true`, 1)
	if _, err := PreviewRealWinAppRunAcceptanceJSON([]byte(unsafe)); err == nil {
		t.Fatalf("acceptance must reject unsafe host mutation evidence")
	}
}

func realWinAppRunAcceptanceExecuteResultFixture(version string) string {
	return strings.ReplaceAll(`{
  "schema_version": "xnix.scripts.remote_wine_guest_gui_smoke.execute_result.v1",
  "request_type": "remote-wine-guest-gui-smoke",
  "version": "VERSION_PLACEHOLDER",
  "status": "passed",
  "execute": true,
  "remote_host": "root@q4",
  "remote_runtime_bin": "/home/xnix-build-cache/bin/xnix-runtime-go",
  "remote_file_open_bin": "/home/xnix-build-cache/bin/xnix-compat-open",
  "execute_result_output": "/home/xnix-run-materials/state/wine-gui-execute-result-VERSION_PLACEHOLDER.json",
  "execute_result_output_written": true,
  "real_run_receipt_summary_preview_planned": true,
  "real_run_receipt_summary_output": "/home/xnix-run-materials/state/wine-gui-real-run-receipt-summary-VERSION_PLACEHOLDER.json",
  "real_run_receipt_summary_output_written": true,
  "real_run_receipt_summary_ready": true,
  "real_run_receipt_summary_file_open_verified": true,
  "real_run_receipt_summary_center_output": "/home/xnix-run-materials/state/wine-gui-real-run-receipt-center-VERSION_PLACEHOLDER.json",
  "real_run_receipt_summary_center_output_written": true,
  "real_run_receipt_summary_center_known_app_smoke_evidence_count": 1,
  "real_run_receipt_summary_kde_page_output": "/home/xnix-run-materials/state/wine-gui-real-run-receipt-kde-page-VERSION_PLACEHOLDER.json",
  "real_run_receipt_summary_kde_page_output_written": true,
  "real_run_receipt_summary_kde_page_known_app_gui_evidence_count": 1,
  "real_run_receipt_summary_kde_page_owner_file_open_verified_count": 1,
  "real_run_receipt_summary_kde_page_owner_file_open_entrypoint_count": 1,
  "report_output": "/home/xnix-run-materials/state/wine-gui-smoke-VERSION_PLACEHOLDER.json",
  "evidence_output": "/home/xnix-run-materials/state/wine-gui-evidence-VERSION_PLACEHOLDER.json",
  "kde_page_output": "/home/xnix-run-materials/state/wine-gui-kde-page-VERSION_PLACEHOLDER.json",
  "kde_action_output": "/home/xnix-run-materials/state/wine-gui-kde-action-VERSION_PLACEHOLDER.json",
  "evidence_output_written": true,
  "kde_page_output_written": true,
  "kde_action_output_written": true,
  "smoke_status": "passed",
  "backend": "qemu-guest-wine-x11",
  "launch_mode": "owner-controlled-launch",
  "file_open_entrypoint_requested": true,
  "file_open_static_entrypoint": "xnix-compat-open %U",
  "known_app_id": "org.xnix.sample.notepad",
  "known_app_name": "Sample Notepad",
  "known_app_version": "VERSION_PLACEHOLDER",
  "gui_app_name": "notepad.exe",
  "runtime_go_owned_gui_smoke": true,
  "qemu_started": true,
  "guest_ssh_ready": true,
  "file_argument_count": 1,
  "file_argument_copied_count": 1,
  "file_arguments_passed": true,
  "file_argument_winepath_translated": true,
  "file_argument_winepath_translated_count": 1,
  "raw_file_argument_path_exposed": false,
  "window_match": "sample-document.txt",
  "window_match_observed": true,
  "window_evidence_summary": "0xa00003 \"file-1-sample-document.txt - Notepad\": (\"notepad.exe\" \"notepad.exe\")",
  "x_window_observed": true,
  "runtime_evidence_report_consumed": true,
  "runtime_evidence_window_observed": true,
  "runtime_evidence_owner_file_open_verified": true,
  "runtime_evidence_owner_file_open_entrypoint_invoked": true,
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
