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
	evidence := packet.KnownAppSmokeEvidence
	if !evidence.StagedLauncherVerified ||
		evidence.LaunchAuthorizationReceiptID != "known-app-launch-authorization-org.xnix.sample.notepad-0.2.640-test" ||
		evidence.LaunchGateState != "controlled-dispatch-ready" ||
		!evidence.LaunchGateConsumed ||
		!evidence.LaunchGateReceiptAccepted ||
		!evidence.LaunchGateGuestBoundaryAccepted ||
		!evidence.ControlledDispatchReady ||
		evidence.ControlledExecutionSessionID != "known-app-controlled-execution-session-org.xnix.sample.notepad-0.2.640-test" ||
		!evidence.LauncherSessionGateConsumed ||
		!evidence.LauncherSessionDigestVerified ||
		evidence.LauncherSessionRelativePath != "execution-ledger/sessions/known-app-controlled-execution-session-org.xnix.sample.notepad-0.2.640-test.json" ||
		!evidence.LauncherSessionRuntimeOwnerConsumable ||
		!evidence.LauncherSessionKDEReadModelConsumable ||
		!evidence.PostReviewDispatchConsumed ||
		evidence.PostReviewDispatchState != "created-after-session-gated-review" ||
		evidence.SessionGatedReviewReceiptID != "known-app-session-gated-launch-review-org.xnix.sample.notepad-0.2.640-test-known-app-controlled-execution-session-org.xnix.sample.notepad-0.2.640-test" {
		t.Fatalf("raw Runtime payload must surface staged launcher session evidence: %#v", evidence)
	}
}

func TestPreviewRealWinAppGUIEvidencePacketConsumesRawExternalExecutableRuntimePayload(t *testing.T) {
	packet, err := PreviewRealWinAppGUIEvidencePacketJSON([]byte(rawExternalExecutableContainerXGUIRuntimePayloadFixture()), RealWinAppGUIEvidencePacketRequest{})
	if err != nil {
		t.Fatalf("PreviewRealWinAppGUIEvidencePacketJSON returned error: %v", err)
	}
	if packet.SchemaVersion != RealWinAppGUIEvidencePacketSchemaVersion ||
		packet.RequestType != RealWinAppGUIEvidencePacketRequestType ||
		packet.Version != "0.2.640-test" ||
		packet.AppID != "org.xnix.external.notepad-file" ||
		packet.DisplayName != "External Notepad File" ||
		packet.AppVersion != "0.2.640-test" ||
		packet.GUIAppName != "/notepad.exe" ||
		packet.EvidenceSource != GUISmokeEvidenceSourceContainerXGUI ||
		packet.RecipeBacked ||
		packet.RecipeAppID != "" ||
		packet.ExecutableName != "notepad.exe" ||
		!packet.LocalExecutableCopied ||
		!packet.ExternalAppImportRecordConsumed ||
		!packet.ImportedArtifactDigestVerified ||
		packet.ImportedArtifactSHA256 != "0d6f23e63c59bc99171659b6b1268010f5b37ee52adc8b9c79984dc8d9d7b208" ||
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
		t.Fatalf("unexpected external executable real GUI packet: %#v", packet)
	}
	evidence := packet.KnownAppSmokeEvidence
	if evidence.AppID != "org.xnix.external.notepad-file" ||
		evidence.DisplayName != "External Notepad File" ||
		evidence.AppVersion != "0.2.640-test" ||
		evidence.RecipeBacked ||
		evidence.RecipeAppID != "" ||
		evidence.EvidenceSource != GUISmokeEvidenceSourceContainerXGUI ||
		!evidence.ExternalAppImportRecordConsumed ||
		!evidence.ImportedArtifactDigestVerified ||
		evidence.ImportedArtifactSHA256 != packet.ImportedArtifactSHA256 ||
		evidence.CompatibilityState != "real-gui-container-wine-verified" ||
		evidence.CenterCardState != "validated-real-gui-container-run" ||
		!evidence.ExecutionEvidenceRecorded ||
		!evidence.RuntimeDispatchVerified ||
		evidence.DesktopLaunchEnabled ||
		evidence.BackendLaunchEnabled ||
		evidence.BackendDetailsExposed ||
		evidence.HostRootModified {
		t.Fatalf("unexpected external executable packet evidence: %#v", evidence)
	}
}

func TestExternalAppRecipeFromRealWinAppGUIEvidencePacket(t *testing.T) {
	packet, err := PreviewRealWinAppGUIEvidencePacketJSON([]byte(rawExternalExecutableContainerXGUIRuntimePayloadFixture()), RealWinAppGUIEvidencePacketRequest{})
	if err != nil {
		t.Fatalf("PreviewRealWinAppGUIEvidencePacketJSON returned error: %v", err)
	}
	payload, err := json.Marshal(packet)
	if err != nil {
		t.Fatalf("Marshal returned error: %v", err)
	}

	recipe, provenance, evidence, err := ExternalAppRecipeFromRealWinAppGUIEvidencePacket(payload)
	if err != nil {
		t.Fatalf("ExternalAppRecipeFromRealWinAppGUIEvidencePacket returned error: %v", err)
	}
	if recipe.ID != "org.xnix.external.notepad-file" ||
		recipe.Name != "External Notepad File" ||
		recipe.Version != "0.2.640-test" ||
		recipe.Icon != "application-x-executable" ||
		recipe.Mode != "automatic" ||
		len(recipe.SupportedExtensions) != 0 ||
		provenance.Source != "real-gui-evidence-packet" ||
		provenance.RegistryName != "runtime-observed-external-app" ||
		!provenance.DigestVerified ||
		provenance.SignatureStatus != "runtime-import-record-digest-verified" ||
		evidence.AppID != recipe.ID ||
		evidence.DisplayName != recipe.Name ||
		evidence.AppVersion != recipe.Version ||
		evidence.RecipeBacked ||
		evidence.RecipeAppID != "" ||
		!evidence.ExecutionEvidenceRecorded ||
		!evidence.RuntimeDispatchVerified ||
		evidence.DesktopLaunchEnabled ||
		evidence.BackendLaunchEnabled ||
		evidence.HostRootModified {
		t.Fatalf("unexpected external app recipe projection: recipe=%#v provenance=%#v evidence=%#v", recipe, provenance, evidence)
	}

	page, err := NewKDECenterPagePreviewWithOptions(recipe, provenance, "approved", nil, KDECenterPageOptions{
		KnownAppSmokeEvidence: []KnownAppSmokeEvidenceSummary{evidence},
	})
	if err != nil {
		t.Fatalf("NewKDECenterPagePreviewWithOptions returned error: %v", err)
	}
	if page.ApplicationID != "org.xnix.external.notepad-file" ||
		page.ApplicationName != "External Notepad File" ||
		page.Icon != "application-x-executable" ||
		page.KnownAppGUIEvidenceCount != 1 ||
		len(page.KnownAppGUIEvidenceCards) != 1 ||
		page.KnownAppGUIEvidenceCards[0].AppID != page.ApplicationID ||
		page.KnownAppGUIEvidenceCards[0].RecipeBacked ||
		!page.KnownAppGUIEvidenceCards[0].ExternalAppImportRecordConsumed ||
		!page.KnownAppGUIEvidenceCards[0].ImportedArtifactDigestVerified ||
		page.KnownAppGUIEvidenceCards[0].ImportedArtifactSHA256 != evidence.ImportedArtifactSHA256 ||
		page.LaunchEnabled ||
		page.BackendProcessStarted ||
		page.BackendDetailsExposed ||
		page.HostRootModified {
		t.Fatalf("unexpected external app KDE page: %#v", page)
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
  "evidence_source": "winapp-smoke-container-x-gui",
  "dispatch_started": true,
  "execution_started": true,
  "smoke_passed": true,
  "runtime_owned_dispatch": true,
  "session_gated_controlled_dispatch_consumed": true,
  "session_gated_controlled_dispatch_state": "created-after-session-gated-review",
  "session_gated_review_receipt_id": "known-app-session-gated-launch-review-org.xnix.sample.notepad-0.2.640-test-known-app-controlled-execution-session-org.xnix.sample.notepad-0.2.640-test",
  "launch_authorization_receipt_id": "known-app-launch-authorization-org.xnix.sample.notepad-0.2.640-test",
  "controlled_execution_session_consumed": true,
  "controlled_execution_session_id": "known-app-controlled-execution-session-org.xnix.sample.notepad-0.2.640-test",
  "controlled_session_digest_verified": true,
  "controlled_session_relative_path": "execution-ledger/sessions/known-app-controlled-execution-session-org.xnix.sample.notepad-0.2.640-test.json",
  "runtime_owner_consumable_session": true,
  "kde_read_model_consumable_session": true,
  "controlled_session_live_state_observed": true,
  "controlled_session_registered": true,
  "controlled_session_window_observed": true,
  "controlled_session_host_root_modified": false,
  "controlled_session_container_process_start": true,
  "raw_command_exposed": false,
  "backend_details_exposed": false,
  "host_root_modified": false,
  "privileged_container_required": false,
  "host_networking_required": false,
  "docker_socket_mounted": false,
  "broad_host_mount_required": false,
  "host_mount_count": 0
}`
}

func rawExternalExecutableContainerXGUIRuntimePayloadFixture() string {
	return `{
  "schema_version": "xnix.runtime.windows_app_container_x_gui_smoke.v1",
  "request_type": "windows-app-container-x-gui-smoke",
  "status": "passed",
  "application_id": "org.xnix.external.notepad-file",
  "display_name": "External Notepad File",
  "app_version": "0.2.640-test",
  "recipe_backed": false,
  "executable_name": "notepad.exe",
  "local_executable_copied": true,
  "external_app_import_record_consumed": true,
  "imported_artifact_digest_verified": true,
  "imported_artifact_sha256": "0d6f23e63c59bc99171659b6b1268010f5b37ee52adc8b9c79984dc8d9d7b208",
  "application_name": "/notepad.exe",
  "window_match": "notepad.exe",
  "container_image": "xnix-wine-smoke:local",
  "container_platform": "linux/arm64",
  "network_mode": "none",
  "x_server_started": true,
  "wine_bootstrap_attempted": true,
  "image_available": true,
  "x_window_observed": true,
  "window_evidence_summary": "0xa00001 \"Untitled - Notepad\": (\"notepad.exe\" \"notepad.exe\") 721x519+4+23 +4+23",
  "host_root_modified": false,
  "privileged_container_required": false,
  "host_networking_required": false,
  "docker_socket_mounted": false,
  "broad_host_mount_required": false,
  "host_mount_count": 0
}`
}
