package appidentity

import (
	"strings"
	"testing"
)

func TestPreviewGUISmokeEvidenceConsumesPassedMessageBoxReport(t *testing.T) {
	preview, err := PreviewGUISmokeEvidenceJSON([]byte(guiSmokeEvidenceFixture(true)), GUISmokeEvidencePreviewRequest{
		AppID:       "org.xnix.fixture.messagebox",
		DisplayName: "Xnix MessageBox",
		AppVersion:  "fixture-version",
	})
	if err != nil {
		t.Fatalf("PreviewGUISmokeEvidenceJSON returned error: %v", err)
	}
	if preview.SchemaVersion != GUISmokeEvidencePreviewSchemaVersion ||
		preview.RequestType != GUISmokeEvidencePreviewRequestType ||
		preview.Source != "wine-guest-gui-smoke+runtime-evidence-consumer" ||
		preview.RuntimeMethod != "PreviewGUISmokeEvidence" ||
		preview.ReadMethod != "GetGUISmokeEvidence" ||
		preview.ReportStatus != "passed" ||
		!preview.ReportConsumed ||
		preview.ReportPathExposed ||
		preview.AppID != "org.xnix.fixture.messagebox" ||
		preview.DisplayName != "Xnix MessageBox" ||
		preview.AppVersion != "fixture-version" ||
		preview.GUIAppName != "xnix-messagebox-smoke.exe" ||
		!preview.LocalGUIExecutableConfigured ||
		!preview.ExecutableCopied ||
		!preview.WinebootInvoked ||
		!preview.XWindowObserved ||
		!preview.WindowObserved ||
		preview.XWindowChildCount != 14 ||
		preview.XWinInfoBytes != 1228 ||
		preview.GuestStderrBytes != 0 ||
		preview.GuestGraphicsDriverErrorObserved ||
		!preview.CompatibilityCenterProjectionReady ||
		!preview.KDECenterProjectionReady ||
		!preview.RuntimeOwned ||
		!preview.GoRuntimeBacked ||
		preview.KDEPolicyOwner ||
		preview.DesktopLaunchEnabled ||
		preview.BackendLaunchEnabled ||
		preview.ActionExecutionEnabled ||
		preview.BackendDetailsExposed ||
		preview.RawOutputExposed ||
		preview.RemotePathExposed ||
		preview.HostRootModified ||
		preview.PrivilegedContainerRequired ||
		preview.HostNetworkingRequired ||
		preview.DockerSocketMounted ||
		preview.BroadHostMountRequired {
		t.Fatalf("unexpected GUI smoke evidence preview: %#v", preview)
	}
	evidence := preview.KnownAppSmokeEvidence
	if evidence.AppID != "org.xnix.fixture.messagebox" ||
		evidence.DisplayName != "Xnix MessageBox" ||
		evidence.EvidenceKind != "known-application-gui-smoke" ||
		evidence.EvidenceSource != "wine-guest-gui-smoke" ||
		evidence.SmokeStatus != "passed" ||
		!evidence.XWindowObserved ||
		!evidence.WindowObserved ||
		evidence.CompatibilityState != "real-gui-qemu-wine-verified" ||
		evidence.CenterCardState != "validated-real-gui-runtime-run" ||
		evidence.PrimaryActionID != "review-known-app-gui-evidence" ||
		evidence.PrimaryActionKind != "review" ||
		!evidence.PrimaryActionEnabled ||
		evidence.MarkerObserved ||
		evidence.ChecksumVerified ||
		!evidence.ExecutionEvidenceRecorded ||
		!evidence.RuntimeDispatchVerified ||
		!evidence.LaunchAuthorizationRequired ||
		evidence.DirectLaunchEnabled ||
		evidence.DesktopLaunchEnabled ||
		evidence.ActionExecutionEnabled ||
		evidence.BackendLaunchEnabled ||
		evidence.HostRootModified ||
		evidence.BackendDetailsExposed ||
		evidence.RawArtifactPathExposed {
		t.Fatalf("unexpected GUI known app smoke evidence: %#v", evidence)
	}
	normalized, err := normalizeKnownAppSmokeEvidenceItem(evidence)
	if err != nil {
		t.Fatalf("normalizeKnownAppSmokeEvidenceItem returned error: %v", err)
	}
	if normalized.CompatibilityState != "real-gui-qemu-wine-verified" ||
		normalized.CenterCardState != "validated-real-gui-runtime-run" ||
		normalized.EvidenceKind != "known-application-gui-smoke" ||
		!normalized.XWindowObserved ||
		!normalized.WindowObserved ||
		!normalized.ExecutionEvidenceRecorded ||
		!normalized.RuntimeDispatchVerified {
		t.Fatalf("unexpected normalized GUI evidence: %#v", normalized)
	}
}

func TestPreviewGUISmokeEvidenceConsumesOwnerControlledMessageBoxReport(t *testing.T) {
	preview, err := PreviewGUISmokeEvidenceJSON([]byte(ownerControlledGUISmokeEvidenceFixture()), GUISmokeEvidencePreviewRequest{
		AppID:       "org.xnix.apps.messagebox",
		DisplayName: "Xnix MessageBox",
		AppVersion:  currentProjectVersion(t),
	})
	if err != nil {
		t.Fatalf("PreviewGUISmokeEvidenceJSON returned error: %v", err)
	}
	if !preview.OwnerControlledLaunchRequested ||
		!preview.OwnerExternalGUIAppRequested ||
		preview.OwnerExternalGUIAppDelivery != "owner-managed-copy" ||
		!preview.OwnerSeedEvidenceProjected ||
		!preview.OwnerServiceCallReady ||
		!preview.OwnerManagedLauncherInvoked ||
		!preview.OwnerDelegatedManagedArtifactCopied ||
		!preview.OwnerDelegatedSmokePassed ||
		preview.OwnerDelegatedEvidenceSource != "wine-guest-gui-smoke" ||
		!preview.OwnerDelegatedExecutionStarted ||
		!preview.OwnerDelegatedWindowObserved ||
		preview.OwnerDelegatedFileArgumentCount != 1 ||
		preview.OwnerDelegatedFileArgumentCopiedCount != 1 ||
		!preview.OwnerDelegatedFileArgumentsPassed ||
		!preview.OwnerDelegatedFileArgumentWinepathTranslated ||
		preview.OwnerDelegatedFileArgumentWinepathTranslatedCount != 1 ||
		preview.OwnerDelegatedRawFileArgumentPathExposed ||
		preview.OwnerDelegatedWindowMatch != "messagebox-document.txt" ||
		!preview.OwnerDelegatedWindowMatchObserved ||
		!preview.FileOpenEntrypointRequested ||
		!preview.OwnerFileOpenEntrypointInvoked ||
		!preview.XWindowObserved ||
		!preview.WindowObserved ||
		!preview.OwnerControlledLaunchVerified ||
		!preview.OwnerManagedCopyVerified ||
		!preview.OwnerFileOpenVerified ||
		!preview.OwnerEvidenceHandoffReady ||
		preview.OwnerEvidenceRelativePath != "runtime/kde-runtime-status-launch-evidence/owner-messagebox.json" ||
		preview.RemotePathExposed ||
		preview.BackendDetailsExposed ||
		preview.HostRootModified {
		t.Fatalf("unexpected owner-controlled GUI smoke evidence preview: %#v", preview)
	}
	evidence := preview.KnownAppSmokeEvidence
	if evidence.AppID != "org.xnix.apps.messagebox" ||
		evidence.CompatibilityState != "owner-controlled-gui-qemu-wine-verified" ||
		evidence.CenterCardState != "validated-owner-controlled-gui-runtime-run" ||
		evidence.PrimaryActionID != KnownAppKDERuntimeStatusLaunchAction ||
		evidence.PrimaryActionLabel != "Show Runtime-controlled launch" ||
		evidence.PrimaryActionKind != "runtime-status" ||
		!evidence.XWindowObserved ||
		!evidence.WindowObserved ||
		!evidence.ExecutionEvidenceRecorded ||
		!evidence.StagedLauncherVerified ||
		!evidence.OwnerControlledRuntimeLaunchVerified ||
		!evidence.OwnerManagedCopyVerified ||
		!evidence.OwnerFileOpenVerified ||
		!evidence.OwnerFileOpenEntrypointInvoked ||
		evidence.OwnerDelegatedFileArgumentCount != 1 ||
		evidence.OwnerDelegatedFileArgumentCopiedCount != 1 ||
		!evidence.OwnerDelegatedFileArgumentsPassed ||
		!evidence.OwnerDelegatedFileArgumentWinepathTranslated ||
		evidence.OwnerDelegatedFileArgumentWinepathTranslatedCount != 1 ||
		evidence.OwnerDelegatedRawFileArgumentPathExposed ||
		evidence.OwnerDelegatedWindowMatch != "messagebox-document.txt" ||
		!evidence.OwnerDelegatedWindowMatchObserved ||
		!evidence.OwnerServiceCallReady ||
		!evidence.OwnerEvidenceHandoffReady ||
		evidence.OwnerEvidenceRelativePath != "runtime/kde-runtime-status-launch-evidence/owner-messagebox.json" ||
		!evidence.RuntimeDispatchVerified ||
		evidence.RawArtifactPathExposed ||
		!strings.Contains(evidence.Summary, "file-open entrypoint launch") ||
		!strings.Contains(evidence.Summary, "handoff ready") {
		t.Fatalf("unexpected owner-controlled GUI known app evidence: %#v", evidence)
	}
}

func TestPreviewGUISmokeEvidenceConsumesContainerXGUIReport(t *testing.T) {
	preview, err := PreviewGUISmokeEvidenceJSON([]byte(containerXGUISmokeEvidenceFixture()), GUISmokeEvidencePreviewRequest{
		AppID:       "org.xnix.sample.notepad",
		DisplayName: "Notepad",
		AppVersion:  "container-local",
	})
	if err != nil {
		t.Fatalf("PreviewGUISmokeEvidenceJSON returned error: %v", err)
	}
	if preview.Source != "winapp-smoke-container-x-gui+runtime-evidence-consumer" ||
		preview.ReportStatus != "passed" ||
		!preview.ReportConsumed ||
		!preview.RecipeBacked ||
		preview.RecipeAppID != "org.xnix.sample.notepad" ||
		preview.GUIAppName != "notepad.exe" ||
		!preview.WinebootInvoked ||
		!preview.XWindowObserved ||
		!preview.WindowObserved ||
		preview.XWindowChildCount != 1 ||
		preview.XWinInfoBytes == 0 ||
		!preview.CompatibilityCenterProjectionReady ||
		!preview.KDECenterProjectionReady ||
		preview.DesktopLaunchEnabled ||
		preview.BackendLaunchEnabled ||
		preview.ActionExecutionEnabled ||
		preview.BackendDetailsExposed ||
		preview.RawOutputExposed ||
		preview.HostRootModified ||
		preview.PrivilegedContainerRequired ||
		preview.HostNetworkingRequired ||
		preview.DockerSocketMounted ||
		preview.BroadHostMountRequired {
		t.Fatalf("unexpected container X GUI smoke evidence preview: %#v", preview)
	}
	evidence := preview.KnownAppSmokeEvidence
	if evidence.AppID != "org.xnix.sample.notepad" ||
		evidence.DisplayName != "Notepad" ||
		evidence.EvidenceKind != "known-application-gui-smoke" ||
		evidence.EvidenceSource != "winapp-smoke-container-x-gui" ||
		!evidence.RecipeBacked ||
		evidence.RecipeAppID != "org.xnix.sample.notepad" ||
		!evidence.XWindowObserved ||
		!evidence.WindowObserved ||
		evidence.CompatibilityState != "real-gui-container-wine-verified" ||
		evidence.CenterCardState != "validated-real-gui-container-run" ||
		!evidence.ExecutionEvidenceRecorded ||
		!evidence.RuntimeDispatchVerified ||
		evidence.DesktopLaunchEnabled ||
		evidence.BackendLaunchEnabled ||
		evidence.BackendDetailsExposed ||
		!strings.Contains(evidence.Summary, "isolated container GUI") {
		t.Fatalf("unexpected container X GUI known app smoke evidence: %#v", evidence)
	}
	normalized, err := normalizeKnownAppSmokeEvidenceItem(evidence)
	if err != nil {
		t.Fatalf("normalizeKnownAppSmokeEvidenceItem returned error: %v", err)
	}
	if normalized.EvidenceSource != "winapp-smoke-container-x-gui" ||
		!normalized.RecipeBacked ||
		normalized.RecipeAppID != "org.xnix.sample.notepad" ||
		!normalized.XWindowObserved ||
		!normalized.WindowObserved ||
		normalized.CompatibilityState != "real-gui-container-wine-verified" ||
		normalized.CenterCardState != "validated-real-gui-container-run" ||
		!normalized.ExecutionEvidenceRecorded ||
		!normalized.RuntimeDispatchVerified {
		t.Fatalf("unexpected normalized container X GUI evidence: %#v", normalized)
	}
}

func TestPreviewGUISmokeEvidenceRejectsExecutedReportWithoutWindow(t *testing.T) {
	_, err := PreviewGUISmokeEvidenceJSON([]byte(guiSmokeEvidenceFixture(false)), GUISmokeEvidencePreviewRequest{})
	if err == nil || !strings.Contains(err.Error(), "observed X window") {
		t.Fatalf("expected observed X window error, got %v", err)
	}
}

func guiSmokeEvidenceFixture(windowObserved bool) string {
	windowObservedValue := "true"
	if !windowObserved {
		windowObservedValue = "false"
	}
	return `{
  "version": "version-under-test",
  "schema_version": "xnix.scripts.wine_guest_gui_smoke.v1",
  "request_type": "wine-guest-gui-smoke",
  "status": "passed",
  "execute": true,
  "backend": "qemu-guest-wine-x11",
  "gui_app_name": "xnix-messagebox-smoke.exe",
  "local_gui_executable_configured": true,
  "runtime_go_owned_gui_smoke": true,
  "host_root_modified": false,
  "privileged_container_required": false,
  "host_networking_required": false,
  "docker_socket_mounted": false,
  "broad_host_mount_required": false,
  "wineboot_invoked": true,
  "x_window_observed": ` + windowObservedValue + `,
  "x_window_child_count": 14,
  "runtime_payload_schema_version": "xnix.runtime.windows_app_guest_wine_gui_smoke.v1",
  "executable_copied": true,
  "xwininfo_bytes": 1228,
  "guest_stderr_bytes": 0,
  "guest_graphics_driver_error_observed": false,
  "kde_safe_output_summary": "Wine GUI window observed; child_windows=14 xwininfo_bytes=1228 guest_stderr_bytes=0"
}`
}

func containerXGUISmokeEvidenceFixture() string {
	return `{
  "version": "version-under-test",
  "schema_version": "xnix.runtime.winapp_smoke_report.v1",
  "report_type": "winapp-smoke",
  "format": "json",
  "backend": "container-x-gui",
  "status": "passed",
  "executable_source": "container-builtin-gui-app",
  "user_executable_supplied": false,
  "fixture_built": false,
  "success_mode": "startup-window",
  "container_smoke_invoked": true,
  "container_x_gui_smoke_invoked": true,
  "smoke_invoked": true,
  "container_image_available": true,
  "container_recipe_backed": true,
  "container_application_id": "org.xnix.sample.notepad",
  "container_display_name": "Sample Notepad",
  "container_app_version": "container-local",
  "container_gui_app": "notepad.exe",
  "container_window_match": "notepad.exe",
  "x_server_started": true,
  "x_window_observed": true,
  "startup_window_observed": true,
  "x_window_evidence_summary": "0x800001 \"Untitled - Notepad\": (\"notepad.exe\" \"notepad.exe\") 721x519+4+23 +4+23",
  "wine_bootstrap_attempted": true,
  "wine_bootstrap_succeeded": true,
  "host_root_modified": false,
  "privileged_container_required": false,
  "host_networking_required": false,
  "docker_socket_mounted": false,
  "broad_host_mount_required": false,
  "container_payload": {
    "schema_version": "xnix.runtime.windows_app_container_x_gui_smoke.v1",
    "request_type": "windows-app-container-x-gui-smoke",
    "status": "passed",
    "application_id": "org.xnix.sample.notepad",
    "display_name": "Sample Notepad",
    "app_version": "container-local",
    "recipe_backed": true,
    "network_mode": "none",
    "x_server_started": true,
    "wine_bootstrap_attempted": true,
    "image_available": true,
    "x_window_observed": true,
    "window_evidence_summary": "0x800001 \"Untitled - Notepad\": (\"notepad.exe\" \"notepad.exe\") 721x519+4+23 +4+23",
    "host_root_modified": false,
    "privileged_container_required": false,
    "host_networking_required": false,
    "docker_socket_mounted": false,
    "broad_host_mount_required": false,
    "host_mount_count": 0
  }
}`
}

func ownerControlledGUISmokeEvidenceFixture() string {
	return `{
  "version": "version-under-test",
  "schema_version": "xnix.scripts.wine_guest_gui_smoke.v1",
  "request_type": "wine-guest-gui-smoke",
  "status": "passed",
  "execute": true,
  "backend": "qemu-guest-wine-x11",
  "gui_app_name": "xnix-messagebox-smoke.exe",
  "local_gui_executable_configured": true,
  "runtime_go_owned_gui_smoke": true,
  "owner_controlled_launch_requested": true,
  "owner_external_gui_app_requested": true,
  "owner_external_gui_app_delivery": "owner-managed-copy",
  "owner_seed_evidence_projected": true,
  "owner_service_call_ready": true,
  "owner_managed_launcher_invoked": true,
  "owner_delegated_managed_artifact_copied": true,
  "owner_delegated_smoke_passed": true,
  "owner_delegated_evidence_source": "wine-guest-gui-smoke",
  "owner_delegated_execution_started": true,
  "owner_delegated_controlled_session_window_observed": true,
  "owner_delegated_file_argument_count": 1,
  "owner_delegated_file_argument_copied_count": 1,
  "owner_delegated_file_arguments_passed": true,
  "owner_delegated_file_argument_winepath_translated": true,
  "owner_delegated_file_argument_winepath_translated_count": 1,
  "owner_delegated_raw_file_argument_path_exposed": false,
  "owner_delegated_window_match": "messagebox-document.txt",
  "owner_delegated_window_match_observed": true,
  "file_open_entrypoint_requested": true,
  "owner_file_open_entrypoint_invoked": true,
  "owner_delegated_host_root_modified": false,
  "owner_delegated_docker_socket_mounted": false,
  "owner_delegated_broad_host_mount_required": false,
  "owner_delegated_raw_command_exposed": false,
  "owner_delegated_backend_details_exposed": false,
  "owner_evidence_handoff_ready": true,
  "owner_evidence_relative_path": "runtime/kde-runtime-status-launch-evidence/owner-messagebox.json",
  "host_root_modified": false,
  "privileged_container_required": false,
  "host_networking_required": false,
  "docker_socket_mounted": false,
  "broad_host_mount_required": false,
  "wineboot_invoked": true,
  "x_window_observed": true,
  "x_window_child_count": 12,
  "runtime_payload_schema_version": "xnix.runtime.windows_app_guest_wine_gui_smoke.v1",
  "executable_copied": true,
  "xwininfo_bytes": 1040,
  "guest_stderr_bytes": 667,
  "guest_graphics_driver_error_observed": false,
  "kde_safe_output_summary": "Wine GUI window observed; child_windows=12 xwininfo_bytes=1040 guest_stderr_bytes=667"
}`
}
