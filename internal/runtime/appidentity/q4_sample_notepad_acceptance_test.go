package appidentity

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestPreviewQ4SampleNotepadAcceptanceConsumesWrapperEvidence(t *testing.T) {
	acceptance, err := PreviewQ4SampleNotepadAcceptanceJSON([]byte(q4SampleNotepadAcceptanceFixture(currentProjectVersion(t))))
	if err != nil {
		t.Fatalf("PreviewQ4SampleNotepadAcceptanceJSON returned error: %v", err)
	}
	if acceptance.SchemaVersion != Q4SampleNotepadAcceptanceSchemaVersion ||
		acceptance.RequestType != Q4SampleNotepadAcceptanceRequestType ||
		acceptance.AcceptanceType != "q4-sample-notepad-real-windows-app-acceptance" ||
		acceptance.SmokeReportConsumed != true ||
		acceptance.SmokeReportPathExposed != false ||
		acceptance.OutputPathExposed != false ||
		acceptance.DelegatedCommandExposed != false ||
		acceptance.RemoteHostExposed != false ||
		acceptance.AppID != "org.xnix.sample.notepad" ||
		acceptance.DisplayName != "Sample Notepad" ||
		acceptance.LaunchMode != "owner-controlled-launch" ||
		acceptance.FileOpenEntrypointRequested != true ||
		acceptance.HostCompilationAvoided != true ||
		acceptance.DelegatedExecuteResultConsumed != true ||
		acceptance.RealRunReceiptSummaryReady != true ||
		acceptance.RealRunAcceptanceReady != true ||
		acceptance.WindowMatchObserved != true ||
		acceptance.OwnerFileOpenEntrypointInvoked != true ||
		acceptance.RuntimeEvidenceOwnerFileOpenEntrypointInvoked != true ||
		acceptance.KDEPageOwnerFileOpenEntrypointCount != 1 ||
		acceptance.HostRootModified != false ||
		acceptance.PrivilegedContainerRequired != false ||
		acceptance.HostNetworkingRequired != false ||
		acceptance.DockerSocketMounted != false ||
		acceptance.BroadHostMountRequired != false ||
		acceptance.BackendDetailsExposed != false ||
		acceptance.RawWindowEvidenceExposed != false ||
		acceptance.ExecutableNameExposed != false ||
		acceptance.AcceptanceReady != true {
		t.Fatalf("unexpected q4 Sample Notepad acceptance: %#v", acceptance)
	}
	encoded, err := json.Marshal(acceptance)
	if err != nil {
		t.Fatalf("Marshal returned error: %v", err)
	}
	lower := strings.ToLower(string(encoded))
	for _, forbidden := range []string{"root@q4", "/tmp/xnix-", "remote_wine_guest_gui_smoke", "qemu-system", "notepad.exe", "sample-document.txt"} {
		if strings.Contains(lower, forbidden) {
			t.Fatalf("q4 Sample Notepad acceptance exposed forbidden term %q: %s", forbidden, lower)
		}
	}
}

func TestPreviewQ4SampleNotepadAcceptanceRejectsIncompleteEvidence(t *testing.T) {
	fixture := q4SampleNotepadAcceptanceFixture(currentProjectVersion(t))
	missingAcceptance := strings.Replace(fixture, `"real_run_acceptance_ready": true`, `"real_run_acceptance_ready": false`, 1)
	if _, err := PreviewQ4SampleNotepadAcceptanceJSON([]byte(missingAcceptance)); err == nil {
		t.Fatalf("q4 Sample Notepad acceptance must reject missing Go real-run acceptance")
	}
	missingWindow := strings.Replace(fixture, `"window_match_observed": true`, `"window_match_observed": false`, 1)
	if _, err := PreviewQ4SampleNotepadAcceptanceJSON([]byte(missingWindow)); err == nil {
		t.Fatalf("q4 Sample Notepad acceptance must reject missing window evidence")
	}
	unsafe := strings.Replace(fixture, `"host_networking_required": false`, `"host_networking_required": true`, 1)
	if _, err := PreviewQ4SampleNotepadAcceptanceJSON([]byte(unsafe)); err == nil {
		t.Fatalf("q4 Sample Notepad acceptance must reject unsafe host networking evidence")
	}
}

func q4SampleNotepadAcceptanceFixture(version string) string {
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
  "source_sync_planned": true,
  "app_id": "org.xnix.sample.notepad",
  "display_name": "Sample Notepad",
  "sample_file_argument": "sample-document.txt",
  "window_match": "sample-document.txt",
  "launch_mode": "owner-controlled-launch",
  "file_open_entrypoint_requested": true,
  "real_run_acceptance_required": true,
  "q4_compile_required": true,
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
