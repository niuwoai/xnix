package appidentity

import (
	"strings"
	"testing"
)

func TestPreviewQ4StagedExternalWinAppAcceptanceConsumesPassedSmoke(t *testing.T) {
	acceptance, err := PreviewQ4StagedExternalWinAppAcceptanceJSON([]byte(q4StagedExternalWinAppAcceptanceFixture(currentProjectVersion(t))))
	if err != nil {
		t.Fatalf("PreviewQ4StagedExternalWinAppAcceptanceJSON returned error: %v", err)
	}
	if acceptance.SchemaVersion != Q4StagedExternalWinAppAcceptanceSchemaVersion ||
		acceptance.RequestType != Q4StagedExternalWinAppAcceptanceRequestType ||
		acceptance.AcceptanceType != "staged-external-winapp-desktop-file-open-acceptance" ||
		!acceptance.SmokeReportConsumed ||
		acceptance.SmokeReportPathExposed ||
		acceptance.OutputPathExposed ||
		acceptance.DelegatedCommandExposed ||
		acceptance.RemoteHostExposed ||
		acceptance.AppID != "org.xnix.external.desktop-notepad-file-argument" ||
		acceptance.DisplayName != "External Desktop Notepad File Argument" ||
		!acceptance.HostCompilationAvoided ||
		!acceptance.DesktopExecUsesExternalAppHandle ||
		acceptance.DesktopArgumentMode != "file-uri" ||
		!acceptance.DesktopFileOpenLane ||
		!acceptance.ExternalFileOpenRequested ||
		!acceptance.ExternalFileBridgeReady ||
		!acceptance.WindowsProcessFileArgumentWindowObserved ||
		!acceptance.ApplicationDetailRealWindowsAppRunVerified ||
		acceptance.KDEPageFromApplicationDetailHeaderBadge != "Verified real app run" ||
		acceptance.KDEPageFromApplicationDetailSummaryCompatibilityState != "real-app-run-verified" ||
		acceptance.KDEPageFromApplicationDetailSummaryBackendLaunchEnabled ||
		acceptance.BackendDetailsExposed ||
		acceptance.RawPathExposed ||
		!acceptance.AcceptanceReady {
		t.Fatalf("unexpected q4 staged external Windows app acceptance: %#v", acceptance)
	}
}

func TestPreviewQ4StagedExternalWinAppAcceptanceRejectsIncompleteEvidence(t *testing.T) {
	fixture := q4StagedExternalWinAppAcceptanceFixture(currentProjectVersion(t))
	for name, mutated := range map[string]string{
		"missing verified KDE summary": strings.Replace(fixture, `"kde_page_from_application_detail_summary_compatibility_state": "real-app-run-verified"`, `"kde_page_from_application_detail_summary_compatibility_state": "registered"`, 1),
		"missing file bridge":          strings.Replace(fixture, `"external_file_bridge_ready": true`, `"external_file_bridge_ready": false`, 1),
		"unsafe host root":             strings.Replace(fixture, `"host_root_modified": false`, `"host_root_modified": true`, 1),
		"missing artifacts":            strings.Replace(fixture, `"artifact_fetch_count": 10`, `"artifact_fetch_count": 9`, 1),
	} {
		if _, err := PreviewQ4StagedExternalWinAppAcceptanceJSON([]byte(mutated)); err == nil {
			t.Fatalf("%s: expected rejection", name)
		}
	}
}

func TestPreviewQ4StagedExternalWinAppAcceptanceConsumesGUIOnlySmoke(t *testing.T) {
	fixture := q4StagedExternalWinAppAcceptanceGUIOnlyFixture(currentProjectVersion(t))
	acceptance, err := PreviewQ4StagedExternalWinAppAcceptanceJSON([]byte(fixture))
	if err != nil {
		t.Fatalf("PreviewQ4StagedExternalWinAppAcceptanceJSON returned error: %v", err)
	}
	if acceptance.AcceptanceType != "staged-external-winapp-desktop-gui-only-acceptance" ||
		acceptance.DesktopArgumentMode != "none" ||
		acceptance.DesktopFileOpenLane ||
		acceptance.ExternalFileOpenRequested ||
		acceptance.ExternalFileBridgeReady ||
		acceptance.WindowsProcessFileArgumentWindowObserved ||
		acceptance.ApplicationDetailFileOpenVerified ||
		!acceptance.WindowObserved ||
		!acceptance.XWindowObserved ||
		!acceptance.AcceptanceReady {
		t.Fatalf("unexpected q4 staged external Windows app GUI-only acceptance: %#v", acceptance)
	}

	missingWindow := strings.Replace(fixture, `"window_observed": true`, `"window_observed": false`, 1)
	if _, err := PreviewQ4StagedExternalWinAppAcceptanceJSON([]byte(missingWindow)); err == nil {
		t.Fatalf("GUI-only acceptance must reject missing window evidence")
	}
}

func q4StagedExternalWinAppAcceptanceGUIOnlyFixture(version string) string {
	replacer := strings.NewReplacer(
		`"app_id": "org.xnix.external.desktop-notepad-file-argument"`, `"app_id": "org.xnix.external.putty"`,
		`"display_name": "External Desktop Notepad File Argument"`, `"display_name": "PuTTY"`,
		`"desktop_argument_mode": "file-uri"`, `"desktop_argument_mode": "none"`,
		`"desktop_file_open_lane": true`, `"desktop_file_open_lane": false`,
		`"external_file_open_requested": true`, `"external_file_open_requested": false`,
		`"external_file_bridge_ready": true`, `"external_file_bridge_ready": false`,
		`"external_file_bridge_arguments_passed": true`, `"external_file_bridge_arguments_passed": false`,
		`"external_file_bridge_winepath_translated": true`, `"external_file_bridge_winepath_translated": false`,
		`"windows_process_file_argument_window_observed": true`, `"windows_process_file_argument_window_observed": false`,
		`"application_detail_file_open_verified": true`, `"application_detail_file_open_verified": false`,
	)
	return replacer.Replace(q4StagedExternalWinAppAcceptanceFixture(version))
}

func q4StagedExternalWinAppAcceptanceFixture(version string) string {
	return strings.ReplaceAll(`{
  "schema_version": "xnix.scripts.q4_staged_desktop_external_winapp_smoke.v1",
  "request_type": "q4-staged-desktop-external-winapp-smoke",
  "version": "VERSION_PLACEHOLDER",
  "status": "passed",
  "execute": true,
  "remote_host": "root@q4",
  "delegated_command": ["ruby", "scripts/staged_desktop_external_winapp_smoke.rb", "--execute"],
  "q4_compile_required": true,
  "host_compilation_avoided": true,
  "remote_build_completed": true,
  "delegated_status": "passed",
  "app_id": "org.xnix.external.desktop-notepad-file-argument",
  "display_name": "External Desktop Notepad File Argument",
  "desktop_argument_mode": "file-uri",
  "desktop_file_open_lane": true,
  "desktop_exec_uses_external_app_handle": true,
  "external_app_desktop_handle_ready": true,
  "desktop_exec_invocation_exact": true,
  "external_file_open_requested": true,
  "external_file_bridge_ready": true,
  "external_file_bridge_arguments_passed": true,
  "external_file_bridge_winepath_translated": true,
  "windows_process_file_argument_window_observed": true,
  "external_file_bridge_mount_enabled": false,
  "raw_file_uri_arguments_exposed": false,
  "external_app_handle_consumed": true,
  "imported_artifact_digest_verified": true,
  "window_observed": true,
  "x_window_observed": true,
  "container_network_mode": "none",
  "container_host_mount_count": 0,
  "one_shot_runtime_launch_executed": true,
  "one_shot_window_observed": true,
  "one_shot_host_root_modified": false,
  "one_shot_docker_socket_mounted": false,
  "one_shot_raw_paths_exposed": false,
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
  "kde_page_from_application_detail_header_backend_details_exposed": false,
  "kde_page_from_application_detail_summary_compatibility_state": "real-app-run-verified",
  "kde_page_from_application_detail_summary_compatibility_label": "Verified real app run",
  "kde_page_from_application_detail_summary_diagnostics_state": "real-app-run-verified",
  "kde_page_from_application_detail_summary_backend_launch_enabled": false,
  "kde_page_from_application_detail_summary_action_execution_enabled": false,
  "kde_page_from_application_detail_summary_settings_persistence_enabled": false,
  "kde_page_from_application_detail_summary_backend_details_exposed": false,
  "kde_page_from_application_detail_summary_host_root_modified": false,
  "artifact_fetch_count": 10,
  "host_root_modified": false,
  "privileged_container_required": false,
  "host_networking_required": false,
  "docker_socket_mounted": false,
  "broad_host_mount_required": false
}`, "VERSION_PLACEHOLDER", version)
}
