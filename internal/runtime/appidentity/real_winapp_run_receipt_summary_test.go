package appidentity

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestPreviewRealWinAppRunReceiptSummaryConsumesPassedQ4FileOpenSmoke(t *testing.T) {
	version := currentProjectVersion(t)
	summary, err := PreviewRealWinAppRunReceiptSummaryJSON([]byte(realWinAppRunReceiptSummaryRemoteSmokeFixture(version)))
	if err != nil {
		t.Fatalf("PreviewRealWinAppRunReceiptSummaryJSON returned error: %v", err)
	}
	if summary.SchemaVersion != RealWinAppRunReceiptSummarySchemaVersion ||
		summary.RequestType != RealWinAppRunReceiptSummaryRequestType ||
		summary.ReceiptType != "real-windows-app-run-receipt-summary" ||
		summary.ReportConsumed != true ||
		summary.ReportPathExposed != false ||
		summary.RemoteHostExposed != false ||
		summary.AppID != "org.xnix.sample.notepad" ||
		summary.DisplayName != "Sample Notepad" ||
		summary.AppVersion != version ||
		summary.GUIAppName != "known-gui-app" ||
		summary.ExecutionHostClass != "q4-remote-validation-host" ||
		summary.BackendClass != "managed-guest-gui" ||
		summary.LaunchMode != "owner-controlled-launch" ||
		summary.RunPassed != true ||
		summary.RealExecutionObserved != true ||
		summary.WindowObserved != true ||
		summary.WindowMatch != "sample-document.txt" ||
		summary.WindowMatchObserved != true ||
		summary.FileOpenVerified != true ||
		summary.FileArgumentCount != 1 ||
		summary.FileArgumentCopiedCount != 1 ||
		summary.FileArgumentsPassed != true ||
		summary.FileArgumentWinePathTranslated != true ||
		summary.FileArgumentWinePathTranslatedCount != 1 ||
		summary.RawFileArgumentPathExposed != false ||
		summary.OwnerControlledLaunchVerified != true ||
		summary.OwnerManagedLauncherInvoked != true ||
		summary.OwnerFileOpenEntrypointInvoked != true ||
		summary.OwnerDelegatedSmokePassed != true ||
		summary.RuntimeEvidenceConsumed != true ||
		summary.RuntimeEvidenceOwnerFileOpenVerified != true ||
		summary.RuntimeEvidenceOwnerFileOpenEntrypointInvoked != true ||
		summary.KDEPageEvidenceConsumed != true ||
		summary.KDEActionEvidenceConsumed != true ||
		summary.KDEActionOwnerFileOpenVerified != true ||
		summary.KDEActionOwnerFileOpenEntrypointInvoked != true ||
		summary.HostRootModified != false ||
		summary.PrivilegedContainerRequired != false ||
		summary.HostNetworkingRequired != false ||
		summary.DockerSocketMounted != false ||
		summary.BroadHostMountRequired != false ||
		summary.BackendDetailsExposed != false ||
		summary.RawLauncherOutputExposed != false ||
		summary.ReceiptReady != true {
		t.Fatalf("unexpected real run receipt summary: %#v", summary)
	}
	encoded, err := json.Marshal(summary)
	if err != nil {
		t.Fatalf("Marshal returned error: %v", err)
	}
	text := strings.ToLower(string(encoded))
	for _, forbidden := range []string{"/home/xnix-", "root@q4", "notepad.exe", "qemu-system", "/var/run/docker.sock", "sample-document.txt - notepad"} {
		if strings.Contains(text, forbidden) {
			t.Fatalf("real run receipt summary exposed forbidden term %q: %s", forbidden, text)
		}
	}
}

func TestPreviewRealWinAppRunReceiptSummaryRejectsIncompleteEvidence(t *testing.T) {
	fixture := realWinAppRunReceiptSummaryRemoteSmokeFixture(currentProjectVersion(t))
	incomplete := strings.Replace(fixture, `"kde_action_owner_file_open_verified": true`, `"kde_action_owner_file_open_verified": false`, 1)
	if _, err := PreviewRealWinAppRunReceiptSummaryJSON([]byte(incomplete)); err == nil {
		t.Fatalf("receipt summary must reject incomplete KDE action evidence")
	}
	unsafe := strings.Replace(fixture, `"host_root_modified": false`, `"host_root_modified": true`, 1)
	if _, err := PreviewRealWinAppRunReceiptSummaryJSON([]byte(unsafe)); err == nil {
		t.Fatalf("receipt summary must reject unsafe host mutation evidence")
	}
}

func TestPreviewRealWinAppRunReceiptSummaryConsumesExternalAppIdentity(t *testing.T) {
	version := currentProjectVersion(t)
	fixture := realWinAppRunReceiptSummaryRemoteSmokeFixture(version)
	fixture = strings.Replace(fixture, `"known_app_id": "org.xnix.sample.notepad"`, `"known_app_id": ""`, 1)
	fixture = strings.Replace(fixture, `"known_app_name": "Sample Notepad"`, `"known_app_name": ""`, 1)
	fixture = strings.Replace(fixture, `"known_app_version": "VERSION_PLACEHOLDER"`, `"known_app_version": ""`, 1)
	fixture = strings.Replace(fixture, `"runtime_evidence_app_id": "org.xnix.sample.notepad"`, `"runtime_evidence_app_id": "org.xnix.apps.messagebox"`, 1)
	fixture = strings.Replace(fixture, `"runtime_evidence_display_name": "Sample Notepad"`, `"runtime_evidence_display_name": "Xnix MessageBox"`, 1)

	summary, err := PreviewRealWinAppRunReceiptSummaryJSON([]byte(fixture))
	if err != nil {
		t.Fatalf("PreviewRealWinAppRunReceiptSummaryJSON returned error for external identity: %v", err)
	}
	if summary.AppID != "org.xnix.apps.messagebox" ||
		summary.DisplayName != "Xnix MessageBox" ||
		summary.AppVersion != version ||
		!summary.ReceiptReady {
		t.Fatalf("unexpected external app receipt summary: %#v", summary)
	}
}

func TestKnownAppSmokeEvidenceFromRealWinAppRunReceiptSummary(t *testing.T) {
	summary, err := PreviewRealWinAppRunReceiptSummaryJSON([]byte(realWinAppRunReceiptSummaryRemoteSmokeFixture(currentProjectVersion(t))))
	if err != nil {
		t.Fatalf("PreviewRealWinAppRunReceiptSummaryJSON returned error: %v", err)
	}
	encoded, err := json.Marshal(summary)
	if err != nil {
		t.Fatalf("Marshal returned error: %v", err)
	}

	evidence, err := KnownAppSmokeEvidenceFromRealWinAppRunReceiptSummary(encoded)
	if err != nil {
		t.Fatalf("KnownAppSmokeEvidenceFromRealWinAppRunReceiptSummary returned error: %v", err)
	}
	if evidence.AppID != "org.xnix.sample.notepad" ||
		evidence.DisplayName != "Sample Notepad" ||
		evidence.EvidenceKind != "known-application-gui-smoke" ||
		evidence.EvidenceSource != GUISmokeEvidenceSourceWineGuest ||
		evidence.SmokeStatus != "passed" ||
		evidence.CompatibilityState != "owner-controlled-gui-qemu-wine-verified" ||
		evidence.CenterCardState != "validated-owner-controlled-gui-runtime-run" ||
		evidence.PrimaryActionID != "review-known-app-gui-evidence" ||
		evidence.PrimaryActionKind != "review" ||
		!evidence.XWindowObserved ||
		!evidence.WindowObserved ||
		!evidence.ExecutionEvidenceRecorded ||
		!evidence.StagedLauncherVerified ||
		!evidence.RuntimeDispatchVerified ||
		!evidence.OwnerControlledRuntimeLaunchVerified ||
		!evidence.OwnerFileOpenVerified ||
		!evidence.OwnerFileOpenEntrypointInvoked ||
		evidence.OwnerDelegatedFileArgumentCount != 1 ||
		evidence.OwnerDelegatedFileArgumentCopiedCount != 1 ||
		!evidence.OwnerDelegatedFileArgumentsPassed ||
		!evidence.OwnerDelegatedFileArgumentWinepathTranslated ||
		evidence.OwnerDelegatedFileArgumentWinepathTranslatedCount != 1 ||
		evidence.OwnerDelegatedRawFileArgumentPathExposed ||
		evidence.OwnerDelegatedWindowMatch != "sample-document.txt" ||
		!evidence.OwnerDelegatedWindowMatchObserved ||
		evidence.DesktopLaunchEnabled ||
		evidence.BackendLaunchEnabled ||
		evidence.HostRootModified ||
		evidence.BackendDetailsExposed ||
		evidence.RawArtifactPathExposed {
		t.Fatalf("unexpected real run receipt known-app evidence: %#v", evidence)
	}
}

func realWinAppRunReceiptSummaryRemoteSmokeFixture(version string) string {
	return strings.ReplaceAll(`{
  "schema_version": "xnix.scripts.remote_wine_guest_gui_smoke.execute_result.v1",
  "request_type": "remote-wine-guest-gui-smoke",
  "version": "VERSION_PLACEHOLDER",
  "status": "passed",
  "execute": true,
  "remote_host": "root@q4",
  "smoke_status": "passed",
  "backend": "qemu-guest-wine-x11",
  "launch_mode": "owner-controlled-launch",
  "known_app_id": "org.xnix.sample.notepad",
  "known_app_name": "Sample Notepad",
  "known_app_version": "VERSION_PLACEHOLDER",
  "gui_app_name": "notepad.exe",
  "runtime_go_owned_gui_smoke": true,
  "qemu_started": true,
  "guest_ssh_ready": true,
  "wineboot_invoked": true,
  "guest_x11_driver_available": true,
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
  "runtime_evidence_app_id": "org.xnix.sample.notepad",
  "runtime_evidence_display_name": "Sample Notepad",
  "runtime_evidence_app_version": "VERSION_PLACEHOLDER",
  "runtime_evidence_window_observed": true,
  "runtime_evidence_owner_file_open_verified": true,
  "runtime_evidence_owner_file_open_entrypoint_invoked": true,
  "kde_page_output_written": true,
  "kde_page_known_app_gui_evidence_count": 1,
  "kde_page_known_app_owner_file_open_verified_count": 1,
  "kde_page_known_app_owner_file_open_entrypoint_count": 1,
  "kde_action_output_written": true,
  "kde_action_owner_file_open_verified": true,
  "kde_action_owner_file_open_entrypoint_invoked": true,
  "owner_controlled_launch_requested": true,
  "owner_evidence_handoff_ready": true,
  "owner_managed_launcher_invoked": true,
  "owner_file_open_entrypoint_invoked": true,
  "owner_delegated_smoke_passed": true,
  "owner_delegated_file_argument_count": 1,
  "owner_delegated_file_argument_copied_count": 1,
  "owner_delegated_file_arguments_passed": true,
  "owner_delegated_file_argument_winepath_translated": true,
  "owner_delegated_file_argument_winepath_translated_count": 1,
  "owner_delegated_raw_file_argument_path_exposed": false,
  "owner_delegated_window_match": "sample-document.txt",
  "owner_delegated_window_match_observed": true,
  "host_root_modified": false,
  "privileged_container_required": false,
  "host_networking_required": false,
  "docker_socket_mounted": false,
  "broad_host_mount_required": false
}`, "VERSION_PLACEHOLDER", version)
}
