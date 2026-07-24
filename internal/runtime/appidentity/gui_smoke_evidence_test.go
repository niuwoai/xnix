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
