package appidentity

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"
)

func TestPreviewQ4WinAppAcceptanceConsumesGenericKnownAppEvidence(t *testing.T) {
	acceptance, err := PreviewQ4WinAppAcceptanceJSON([]byte(q4WinAppAcceptanceFixture(currentProjectVersion(t), "org.xnix.apps.mines", false, true)))
	if err != nil {
		t.Fatalf("PreviewQ4WinAppAcceptanceJSON returned error: %v", err)
	}
	if acceptance.SchemaVersion != Q4WinAppAcceptanceSchemaVersion ||
		acceptance.RequestType != Q4WinAppAcceptanceRequestType ||
		acceptance.AcceptanceType != "generic-q4-windows-app-real-run-acceptance" ||
		acceptance.SmokeReportConsumed != true ||
		acceptance.SmokeReportPathExposed != false ||
		acceptance.OutputPathExposed != false ||
		acceptance.DelegatedCommandExposed != false ||
		acceptance.RemoteHostExposed != false ||
		acceptance.AppID != "org.xnix.apps.mines" ||
		acceptance.DisplayName != "Mines" ||
		acceptance.KnownAppID != "org.xnix.apps.mines" ||
		acceptance.KnownAppSelected != true ||
		acceptance.RemoteExecutableConfigured != false ||
		acceptance.RemoteExecutablePathExposed != false ||
		acceptance.RemoteFileArgumentConfigured != false ||
		acceptance.RemoteFileArgumentPathExposed != false ||
		acceptance.SampleFileArgumentConfigured != true ||
		acceptance.LaunchMode != "owner-controlled-launch" ||
		acceptance.FileOpenEntrypointRequested != true ||
		acceptance.RealRunAcceptanceRequired != true ||
		acceptance.Q4CompileRequired != true ||
		acceptance.HostCompilationAvoided != true ||
		acceptance.DelegatedExecuteResultConsumed != true ||
		acceptance.RemoteBuildCompleted != true ||
		acceptance.EvidenceOutputWritten != true ||
		acceptance.KDEPageOutputWritten != true ||
		acceptance.KDEActionOutputWritten != true ||
		acceptance.WindowObserved != true ||
		acceptance.WindowMatchObserved != true ||
		acceptance.DocumentContentMarkerObservationRequired != true ||
		acceptance.DocumentContentMarkerObserved != true ||
		acceptance.RealRunReceiptSummaryReady != true ||
		acceptance.RealRunReceiptSummaryFileOpenVerified != true ||
		acceptance.RealRunAcceptanceReady != true ||
		acceptance.OwnerFileOpenEntrypointInvoked != true ||
		acceptance.RuntimeEvidenceOwnerFileOpenEntrypointInvoked != true ||
		acceptance.KDEPageOwnerFileOpenEntrypointCount != 1 ||
		acceptance.AcceptanceReady != true {
		t.Fatalf("unexpected q4 Windows app acceptance: %#v", acceptance)
	}
	encoded, err := json.Marshal(acceptance)
	if err != nil {
		t.Fatalf("Marshal returned error: %v", err)
	}
	lower := strings.ToLower(string(encoded))
	for _, forbidden := range []string{"root@q4", "/tmp/xnix-", "remote_wine_guest_gui_smoke", "qemu-system", "winemine.exe", "sample-document.txt"} {
		if strings.Contains(lower, forbidden) {
			t.Fatalf("q4 Windows app acceptance exposed forbidden term %q: %s", forbidden, lower)
		}
	}
}

func TestPreviewQ4WinAppAcceptanceConsumesRemoteExecutableEvidence(t *testing.T) {
	acceptance, err := PreviewQ4WinAppAcceptanceJSON([]byte(q4WinAppAcceptanceFixture(currentProjectVersion(t), "", true, false)))
	if err != nil {
		t.Fatalf("PreviewQ4WinAppAcceptanceJSON returned error: %v", err)
	}
	if acceptance.KnownAppSelected ||
		acceptance.KnownAppID != "" ||
		acceptance.RemoteExecutableConfigured != true ||
		acceptance.RealRunAcceptanceRequired ||
		acceptance.LaunchMode != "direct" ||
		acceptance.WindowMatchObserved != true ||
		acceptance.AcceptanceReady != true {
		t.Fatalf("unexpected remote executable acceptance: %#v", acceptance)
	}
}

func TestPreviewQ4WinAppAcceptanceRejectsIncompleteEvidence(t *testing.T) {
	fixture := q4WinAppAcceptanceFixture(currentProjectVersion(t), "org.xnix.apps.mines", false, true)
	missingMatch := strings.Replace(fixture, `"window_match_observed": true`, `"window_match_observed": false`, 1)
	if _, err := PreviewQ4WinAppAcceptanceJSON([]byte(missingMatch)); err == nil {
		t.Fatalf("q4 Windows app acceptance must reject missing window-match evidence")
	}
	unsafe := strings.Replace(fixture, `"docker_socket_mounted": false`, `"docker_socket_mounted": true`, 1)
	if _, err := PreviewQ4WinAppAcceptanceJSON([]byte(unsafe)); err == nil {
		t.Fatalf("q4 Windows app acceptance must reject unsafe Docker socket evidence")
	}
	leakedPath := strings.Replace(fixture, `"remote_executable_path_exposed": false`, `"remote_executable_path_exposed": true`, 1)
	if _, err := PreviewQ4WinAppAcceptanceJSON([]byte(leakedPath)); err == nil {
		t.Fatalf("q4 Windows app acceptance must reject raw executable path exposure")
	}
	missingRealAcceptance := strings.Replace(fixture, `"real_run_acceptance_ready": true`, `"real_run_acceptance_ready": false`, 1)
	if _, err := PreviewQ4WinAppAcceptanceJSON([]byte(missingRealAcceptance)); err == nil {
		t.Fatalf("q4 Windows app acceptance must reject missing required real-run acceptance")
	}
	missingDocumentMarker := strings.Replace(fixture, `"document_content_marker_observed": true`, `"document_content_marker_observed": false`, 1)
	if _, err := PreviewQ4WinAppAcceptanceJSON([]byte(missingDocumentMarker)); err == nil {
		t.Fatalf("q4 Windows app acceptance must reject missing required document content marker evidence")
	}
}

func q4WinAppAcceptanceFixture(version string, knownAppID string, remoteExecutableConfigured bool, requireRealRunAcceptance bool) string {
	appID := "org.example.demo-tool"
	displayName := "Demo Tool"
	launchMode := "direct"
	sampleFileArgument := ""
	fileOpenEntrypointRequested := false
	realRunAcceptanceOutputWritten := false
	realRunAcceptanceReady := false
	realRunAcceptanceCenterProjectionConsumed := false
	realRunAcceptanceKDEPageProjectionConsumed := false
	realRunReceiptSummaryReady := false
	realRunReceiptSummaryFileOpenVerified := false
	documentContentMarkerObservationRequired := false
	documentContentMarkerObserved := false
	ownerFileOpenEntrypointInvoked := false
	runtimeEvidenceOwnerFileOpenEntrypointInvoked := false
	kdePageOwnerFileOpenEntrypointCount := 0
	if knownAppID != "" {
		appID = knownAppID
		displayName = "Mines"
		launchMode = "owner-controlled-launch"
		sampleFileArgument = "sample-document.txt"
		fileOpenEntrypointRequested = true
		realRunAcceptanceOutputWritten = true
		realRunAcceptanceReady = true
		realRunAcceptanceCenterProjectionConsumed = true
		realRunAcceptanceKDEPageProjectionConsumed = true
		realRunReceiptSummaryReady = true
		realRunReceiptSummaryFileOpenVerified = true
		documentContentMarkerObservationRequired = true
		documentContentMarkerObserved = true
		ownerFileOpenEntrypointInvoked = true
		runtimeEvidenceOwnerFileOpenEntrypointInvoked = true
		kdePageOwnerFileOpenEntrypointCount = 1
	}
	return fmt.Sprintf(`{
  "schema_version": "xnix.scripts.q4_winapp_smoke.v1",
  "request_type": "q4-winapp-smoke",
  "version": %q,
  "status": "passed",
  "execute": true,
  "remote_host": "root@q4",
  "delegated_script": "scripts/remote_wine_guest_gui_smoke.rb",
  "delegated_command": ["ruby", "scripts/remote_wine_guest_gui_smoke.rb", "--execute"],
  "output_path": "/tmp/xnix-q4-winapp.json",
  "source_sync_planned": true,
  "known_app_id": %q,
  "remote_executable_configured": %t,
  "remote_executable_path_exposed": false,
  "app_id": %q,
  "display_name": %q,
  "remote_file_argument_configured": false,
  "remote_file_argument_path_exposed": false,
  "sample_file_argument": %q,
  "window_match": "winemine.exe",
  "launch_mode": %q,
  "file_open_entrypoint_requested": %t,
  "real_run_acceptance_required": %t,
  "q4_compile_required": true,
  "host_compilation_avoided": true,
  "host_root_modified": false,
  "privileged_container_required": false,
  "host_networking_required": false,
  "docker_socket_mounted": false,
  "broad_host_mount_required": false,
  "delegated_execute_result_schema": "xnix.scripts.remote_wine_guest_gui_smoke.execute_result.v1",
  "delegated_execute_result_output_written": true,
  "remote_build_completed": true,
  "evidence_output_written": true,
  "kde_page_output_written": true,
  "kde_action_output_written": true,
  "window_observed": true,
  "window_match_observed": true,
  "document_content_marker_observation_required": %t,
  "document_content_marker_observed": %t,
  "real_run_receipt_summary_ready": %t,
  "real_run_receipt_summary_file_open_verified": %t,
  "real_run_acceptance_output_written": %t,
  "real_run_acceptance_ready": %t,
  "real_run_acceptance_center_projection_consumed": %t,
  "real_run_acceptance_kde_page_projection_consumed": %t,
  "owner_file_open_entrypoint_invoked": %t,
  "runtime_evidence_owner_file_open_entrypoint_invoked": %t,
  "kde_page_known_app_owner_file_open_entrypoint_count": %d
}`, version, knownAppID, remoteExecutableConfigured, appID, displayName, sampleFileArgument, launchMode, fileOpenEntrypointRequested, requireRealRunAcceptance, documentContentMarkerObservationRequired, documentContentMarkerObserved, realRunReceiptSummaryReady, realRunReceiptSummaryFileOpenVerified, realRunAcceptanceOutputWritten, realRunAcceptanceReady, realRunAcceptanceCenterProjectionConsumed, realRunAcceptanceKDEPageProjectionConsumed, ownerFileOpenEntrypointInvoked, runtimeEvidenceOwnerFileOpenEntrypointInvoked, kdePageOwnerFileOpenEntrypointCount)
}
