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
		!preview.OwnerControlledLaunchVerified ||
		!preview.OwnerManagedCopyVerified ||
		preview.RemotePathExposed ||
		preview.BackendDetailsExposed ||
		preview.HostRootModified {
		t.Fatalf("unexpected owner-controlled GUI smoke evidence preview: %#v", preview)
	}
	evidence := preview.KnownAppSmokeEvidence
	if evidence.AppID != "org.xnix.apps.messagebox" ||
		evidence.CompatibilityState != "owner-controlled-gui-qemu-wine-verified" ||
		evidence.CenterCardState != "validated-owner-controlled-gui-runtime-run" ||
		!evidence.ExecutionEvidenceRecorded ||
		!evidence.StagedLauncherVerified ||
		!evidence.RuntimeDispatchVerified ||
		evidence.RawArtifactPathExposed ||
		!strings.Contains(evidence.Summary, "managed launcher copied") {
		t.Fatalf("unexpected owner-controlled GUI known app evidence: %#v", evidence)
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
  "owner_delegated_host_root_modified": false,
  "owner_delegated_docker_socket_mounted": false,
  "owner_delegated_broad_host_mount_required": false,
  "owner_delegated_raw_command_exposed": false,
  "owner_delegated_backend_details_exposed": false,
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
