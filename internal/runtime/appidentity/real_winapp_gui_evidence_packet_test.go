package appidentity

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestPreviewRealWinAppGUIEvidencePacketConsumesContainerNotepadReport(t *testing.T) {
	packet, err := PreviewRealWinAppGUIEvidencePacketJSON([]byte(containerXGUISmokeEvidenceFixture()), RealWinAppGUIEvidencePacketRequest{
		AppID:       "org.xnix.sample.notepad",
		DisplayName: "Sample Notepad",
		AppVersion:  "container-local",
	})
	if err != nil {
		t.Fatalf("PreviewRealWinAppGUIEvidencePacketJSON returned error: %v", err)
	}
	if packet.SchemaVersion != RealWinAppGUIEvidencePacketSchemaVersion ||
		packet.RequestType != RealWinAppGUIEvidencePacketRequestType ||
		packet.PacketType != "real-windows-app-gui-evidence" ||
		packet.RuntimeMethod != "PreviewRealWinAppGUIEvidencePacket" ||
		packet.ReadMethod != "GetRealWinAppGUIEvidencePacket" ||
		packet.ReportStatus != "passed" ||
		!packet.ReportConsumed ||
		packet.ReportPathExposed ||
		packet.AppID != "org.xnix.sample.notepad" ||
		packet.DisplayName != "Sample Notepad" ||
		packet.AppVersion != "container-local" ||
		packet.GUIAppName != "notepad.exe" ||
		packet.EvidenceSource != GUISmokeEvidenceSourceContainerXGUI ||
		!packet.RecipeBacked ||
		packet.RecipeAppID != "org.xnix.sample.notepad" ||
		packet.CompatibilityState != "real-gui-container-wine-verified" ||
		packet.CenterCardState != "validated-real-gui-container-run" ||
		packet.KnownAppGUIEvidenceCount != 1 ||
		packet.KnownAppGUIEvidenceVerifiedCount != 1 ||
		!packet.WinebootInvoked ||
		!packet.XWindowObserved ||
		packet.XWindowChildCount != 1 ||
		!packet.CompatibilityCenterProjectionReady ||
		!packet.KDECenterProjectionReady ||
		!packet.ContainerRuntimeUsed ||
		packet.ContainerNetworkMode != "none" ||
		packet.ContainerHostMountCount != 0 ||
		!packet.RuntimeOwned ||
		!packet.GoRuntimeBacked ||
		packet.KDEPolicyOwner ||
		packet.DesktopLaunchEnabled ||
		packet.BackendLaunchEnabled ||
		packet.ActionExecutionEnabled ||
		packet.BackendDetailsExposed ||
		packet.RawOutputExposed ||
		packet.HostRootModified ||
		packet.PrivilegedContainerRequired ||
		packet.HostNetworkingRequired ||
		packet.DockerSocketMounted ||
		packet.BroadHostMountRequired {
		t.Fatalf("unexpected real Windows app GUI evidence packet: %#v", packet)
	}
	if packet.KnownAppSmokeEvidence.EvidenceKind != "known-application-gui-smoke" ||
		packet.KnownAppSmokeEvidence.EvidenceSource != GUISmokeEvidenceSourceContainerXGUI ||
		!packet.KnownAppSmokeEvidence.ExecutionEvidenceRecorded ||
		!packet.KnownAppSmokeEvidence.RuntimeDispatchVerified ||
		packet.KnownAppSmokeEvidence.DesktopLaunchEnabled ||
		packet.KnownAppSmokeEvidence.BackendLaunchEnabled ||
		packet.KnownAppSmokeEvidence.BackendDetailsExposed ||
		packet.KnownAppSmokeEvidence.HostRootModified {
		t.Fatalf("unexpected packet evidence item: %#v", packet.KnownAppSmokeEvidence)
	}
	if !strings.Contains(packet.Source, "real-winapp-desktop-packet") ||
		!strings.Contains(packet.DesktopSafeSummary, "container X GUI smoke evidence observed a real isolated GUI window") {
		t.Fatalf("unexpected packet source or summary: %#v", packet)
	}
}

func TestKnownAppSmokeEvidenceFromRealWinAppGUIEvidencePacket(t *testing.T) {
	packet, err := PreviewRealWinAppGUIEvidencePacketJSON([]byte(containerXGUISmokeEvidenceFixture()), RealWinAppGUIEvidencePacketRequest{
		AppID:       "org.xnix.sample.notepad",
		DisplayName: "Sample Notepad",
		AppVersion:  "container-local",
	})
	if err != nil {
		t.Fatalf("PreviewRealWinAppGUIEvidencePacketJSON returned error: %v", err)
	}
	payload, err := json.Marshal(packet)
	if err != nil {
		t.Fatalf("Marshal returned error: %v", err)
	}

	evidence, err := KnownAppSmokeEvidenceFromRealWinAppGUIEvidencePacket(payload)
	if err != nil {
		t.Fatalf("KnownAppSmokeEvidenceFromRealWinAppGUIEvidencePacket returned error: %v", err)
	}
	if evidence.AppID != "org.xnix.sample.notepad" ||
		evidence.DisplayName != "Sample Notepad" ||
		evidence.AppVersion != "container-local" ||
		evidence.EvidenceKind != "known-application-gui-smoke" ||
		evidence.EvidenceSource != GUISmokeEvidenceSourceContainerXGUI ||
		!evidence.RecipeBacked ||
		evidence.RecipeAppID != "org.xnix.sample.notepad" ||
		evidence.CompatibilityState != "real-gui-container-wine-verified" ||
		evidence.CenterCardState != "validated-real-gui-container-run" ||
		evidence.SmokeStatus != "passed" ||
		!evidence.ExecutionEvidenceRecorded ||
		!evidence.RuntimeDispatchVerified ||
		evidence.DesktopLaunchEnabled ||
		evidence.BackendLaunchEnabled ||
		evidence.BackendDetailsExposed ||
		evidence.HostRootModified {
		t.Fatalf("unexpected evidence from real GUI packet: %#v", evidence)
	}
}

func TestPreviewRealWinAppGUIEvidencePacketConsumesRawContainerRuntimePayload(t *testing.T) {
	packet, err := PreviewRealWinAppGUIEvidencePacketJSON([]byte(rawContainerXGUIRuntimePayloadFixture()), RealWinAppGUIEvidencePacketRequest{
		AppID:       "org.xnix.sample.notepad",
		DisplayName: "Sample Notepad",
		AppVersion:  "0.2.640-test",
	})
	if err != nil {
		t.Fatalf("PreviewRealWinAppGUIEvidencePacketJSON returned error: %v", err)
	}
	if packet.SchemaVersion != RealWinAppGUIEvidencePacketSchemaVersion ||
		packet.RequestType != RealWinAppGUIEvidencePacketRequestType ||
		packet.Version != "0.2.640-test" ||
		packet.AppID != "org.xnix.sample.notepad" ||
		packet.DisplayName != "Sample Notepad" ||
		packet.AppVersion != "0.2.640-test" ||
		packet.GUIAppName != "notepad.exe" ||
		packet.EvidenceSource != GUISmokeEvidenceSourceContainerXGUI ||
		!packet.RecipeBacked ||
		packet.RecipeAppID != "org.xnix.sample.notepad" ||
		packet.CompatibilityState != "real-gui-container-wine-verified" ||
		packet.CenterCardState != "validated-real-gui-container-run" ||
		packet.KnownAppGUIEvidenceCount != 1 ||
		packet.KnownAppGUIEvidenceVerifiedCount != 1 ||
		!packet.ContainerRuntimeUsed ||
		packet.ContainerNetworkMode != "none" ||
		packet.ContainerHostMountCount != 0 ||
		!packet.XWindowObserved ||
		packet.XWindowChildCount != 1 ||
		packet.DesktopLaunchEnabled ||
		packet.BackendLaunchEnabled ||
		packet.BackendDetailsExposed ||
		packet.HostRootModified ||
		packet.DockerSocketMounted {
		t.Fatalf("unexpected real Windows app GUI evidence packet from raw Runtime payload: %#v", packet)
	}
}

func rawContainerXGUIRuntimePayloadFixture() string {
	return `{
  "schema_version": "xnix.runtime.windows_app_container_x_gui_smoke.v1",
  "request_type": "windows-app-container-x-gui-smoke",
  "status": "passed",
  "application_id": "org.xnix.sample.notepad",
  "display_name": "Sample Notepad",
  "app_version": "0.2.640-test",
  "recipe_backed": true,
  "application_name": "notepad.exe",
  "window_match": "notepad.exe",
  "container_image": "xnix-wine-smoke:local",
  "container_platform": "linux/arm64",
  "network_mode": "none",
  "x_server_started": true,
  "wine_bootstrap_attempted": true,
  "image_available": true,
  "x_window_observed": true,
  "window_evidence_summary": "0x600001 \"Untitled - Notepad\": (\"notepad.exe\" \"notepad.exe\")",
  "host_root_modified": false,
  "privileged_container_required": false,
  "host_networking_required": false,
  "docker_socket_mounted": false,
  "broad_host_mount_required": false,
  "host_mount_count": 0
}`
}
