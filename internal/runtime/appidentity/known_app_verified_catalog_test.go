package appidentity

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestPreviewKnownAppVerifiedCatalogConsumesMatrixEvidence(t *testing.T) {
	content := knownAppVerifiedCatalogMatrixEvidenceFixture(t)

	preview, err := PreviewKnownAppVerifiedCatalogJSON(content)
	if err != nil {
		t.Fatalf("PreviewKnownAppVerifiedCatalogJSON returned error: %v", err)
	}
	if preview.SchemaVersion != KnownAppVerifiedCatalogPreviewSchemaVersion ||
		preview.RequestType != KnownAppVerifiedCatalogPreviewRequestType ||
		preview.Source != "known-app-matrix-evidence+runtime-verified-catalog" ||
		preview.Desktop != "KDE Plasma" ||
		preview.RuntimeMethod != "ListKnownVerifiedApplications" ||
		preview.ReadMethod != "ListKnownVerifiedApplicationsPreview" ||
		preview.VerificationSource != KnownAppMatrixEvidencePreviewRequestType ||
		!preview.MatrixEvidenceConsumed ||
		preview.MatrixStatus != "passed" {
		t.Fatalf("unexpected verified catalog schema: %#v", preview)
	}
	if preview.ApplicationCount != 2 ||
		preview.VerifiedApplicationCount != 2 ||
		preview.QEMUExecutedCount != 2 ||
		preview.WineExecutedCount != 2 ||
		preview.ChecksumVerifiedCount != 2 ||
		preview.MarkerObservedCount != 2 ||
		preview.RawOutputRedactedCount != 2 {
		t.Fatalf("unexpected verified catalog counts: %#v", preview)
	}
	if !preview.CompatibilityCenterProjectionReady ||
		!preview.KDECenterProjectionReady ||
		!preview.RuntimeOwned ||
		!preview.GoRuntimeBacked ||
		preview.KDEPolicyOwner ||
		!preview.UserVisible ||
		!preview.StandardDesktopEntries ||
		!preview.ReviewOnly ||
		!preview.OperatorReviewRequired ||
		preview.LaunchEnabled ||
		preview.ExecutionStarted ||
		preview.BackendLaunchEnabled ||
		preview.DesktopFilesWritten ||
		preview.HostRootModified ||
		preview.BackendDetailsExposed ||
		preview.RawOutputExposed ||
		preview.RemotePathExposed {
		t.Fatalf("unexpected verified catalog safety flags: %#v", preview)
	}
	if strings.Join(preview.ApplicationIDs, ",") != "7zr,busybox-w32" {
		t.Fatalf("unexpected application ids: %#v", preview.ApplicationIDs)
	}
	busybox := preview.Applications[1]
	if busybox.AppID != "busybox-w32" ||
		busybox.DisplayName != "BusyBox-w32 standalone console executable" ||
		busybox.VerificationState != "verified-real-q4-matrix-run" ||
		busybox.DesktopCatalogState != "visible-review-only" ||
		busybox.LauncherSurface != "xnix-compat-launch" ||
		strings.Join(busybox.LaunchRequestCommand, " ") != "xnix-compat-launch --app busybox-w32" ||
		busybox.PrimaryActionKind != "review" ||
		busybox.DirectLaunchEnabled ||
		!busybox.OperatorReviewRequired ||
		!busybox.QEMUExecuted ||
		!busybox.WineExecuted ||
		!busybox.ChecksumVerified ||
		!busybox.MarkerObserved ||
		!busybox.RawOutputRedacted ||
		!busybox.SerialLogEvidence ||
		busybox.BackendLaunchEnabled ||
		busybox.BackendDetailsExposed ||
		busybox.RawOutputExposed ||
		busybox.RemotePathExposed ||
		busybox.HostRootModified {
		t.Fatalf("unexpected BusyBox verified catalog entry: %#v", busybox)
	}
}

func TestPreviewKnownAppVerifiedCatalogConsumesGUIEvidencePacket(t *testing.T) {
	matrixContent := knownAppVerifiedCatalogMatrixEvidenceFixture(t)
	packet, err := PreviewRealWinAppGUIEvidencePacketJSON([]byte(containerXGUISmokeEvidenceFixture()), RealWinAppGUIEvidencePacketRequest{
		AppID:       "org.xnix.sample.notepad",
		DisplayName: "Sample Notepad",
		AppVersion:  "container-local",
	})
	if err != nil {
		t.Fatalf("PreviewRealWinAppGUIEvidencePacketJSON returned error: %v", err)
	}
	packetContent, err := json.Marshal(packet)
	if err != nil {
		t.Fatalf("Marshal packet returned error: %v", err)
	}

	preview, err := PreviewKnownAppVerifiedCatalogWithGUIEvidenceJSON(matrixContent, packetContent)
	if err != nil {
		t.Fatalf("PreviewKnownAppVerifiedCatalogWithGUIEvidenceJSON returned error: %v", err)
	}
	if preview.Source != "known-app-matrix-evidence+real-winapp-gui-evidence-packet+runtime-verified-catalog" ||
		!preview.MatrixEvidenceConsumed ||
		!preview.GUIEvidencePacketConsumed ||
		preview.ApplicationCount != 3 ||
		preview.VerifiedApplicationCount != 3 ||
		preview.GUIVerifiedApplicationCount != 1 ||
		preview.GUIWindowObservedCount != 1 ||
		strings.Join(preview.ApplicationIDs, ",") != "7zr,busybox-w32,org.xnix.sample.notepad" ||
		strings.Join(preview.GUIEvidenceApplicationIDs, ",") != "org.xnix.sample.notepad" {
		t.Fatalf("unexpected GUI-augmented verified catalog: %#v", preview)
	}
	notepad := preview.Applications[2]
	if notepad.AppID != "org.xnix.sample.notepad" ||
		notepad.DisplayName != "Sample Notepad" ||
		notepad.VerificationState != "verified-real-gui-q4-run" ||
		notepad.CompatibilityState != "real-gui-container-wine-verified" ||
		notepad.DesktopCatalogState != "visible-review-only" ||
		notepad.EvidenceSource != GUISmokeEvidenceSourceContainerXGUI ||
		!notepad.GUIEvidence ||
		!notepad.WindowObserved ||
		notepad.FileOpenVerified ||
		strings.Join(notepad.LaunchRequestCommand, " ") != "xnix-compat-launch --app org.xnix.sample.notepad" ||
		notepad.PrimaryActionKind != "review" ||
		notepad.DirectLaunchEnabled ||
		!notepad.OperatorReviewRequired ||
		notepad.KDEPolicyOwner ||
		notepad.BackendLaunchEnabled ||
		notepad.BackendDetailsExposed ||
		notepad.RawOutputExposed ||
		notepad.RemotePathExposed ||
		notepad.HostRootModified {
		t.Fatalf("unexpected GUI catalog entry: %#v", notepad)
	}
}

func TestPreviewKnownAppVerifiedCatalogRejectsIncompleteMatrixEvidence(t *testing.T) {
	content := knownAppVerifiedCatalogMatrixEvidenceFixture(t)
	var evidence KnownAppMatrixEvidencePreview
	if err := json.Unmarshal(content, &evidence); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	evidence.Apps = evidence.Apps[:1]
	evidence.AppCount = 1
	malformed, err := json.Marshal(evidence)
	if err != nil {
		t.Fatalf("Marshal returned error: %v", err)
	}

	_, err = PreviewKnownAppVerifiedCatalogJSON(malformed)
	if err == nil || !strings.Contains(err.Error(), "busybox-w32") {
		t.Fatalf("expected missing BusyBox evidence error, got %v", err)
	}
}

func TestPreviewKnownAppVerifiedCatalogRunPlanSelectsGUIRunnableApp(t *testing.T) {
	catalog, err := PreviewKnownAppVerifiedCatalogWithGUIEvidenceJSON(knownAppVerifiedCatalogMatrixEvidenceFixture(t), []byte(knownAppVerifiedCatalogGUIEvidencePacketFixture()))
	if err != nil {
		t.Fatalf("PreviewKnownAppVerifiedCatalogWithGUIEvidenceJSON returned error: %v", err)
	}
	content, err := json.Marshal(catalog)
	if err != nil {
		t.Fatalf("Marshal catalog returned error: %v", err)
	}

	preview, err := PreviewKnownAppVerifiedCatalogRunPlanJSON(content, "org.xnix.apps.messagebox")
	if err != nil {
		t.Fatalf("PreviewKnownAppVerifiedCatalogRunPlanJSON returned error: %v", err)
	}
	if preview.AppID != "org.xnix.apps.messagebox" ||
		preview.VerificationState != "verified-real-gui-q4-run" ||
		preview.CompatibilityState != "owner-controlled-gui-qemu-wine-verified" ||
		strings.Join(preview.LaunchRequestCommand, " ") != "xnix-compat-launch --app org.xnix.apps.messagebox" ||
		strings.Join(preview.RemoteSmokeCommand, " ") != "ruby scripts/q4_messagebox_smoke.rb --execute --owner-file-open" ||
		preview.RemoteSmokeRequestType != "q4-messagebox-smoke" ||
		!preview.RuntimeOwnedActionReady ||
		preview.DesktopCallableActionID != "review-known-app-gui-evidence" ||
		preview.DesktopCallableRoute != "runtime-owner://known-app-verified-catalog/run-plan" ||
		preview.DesktopCallableRuntimeMethod != "PlanKnownVerifiedApplicationRun" ||
		preview.DesktopCallableExecutionType != "review-only-q4-gui-smoke" ||
		strings.Join(preview.DesktopForwardedArguments, " ") != "org.xnix.apps.messagebox" ||
		!preview.DesktopForwardsOnlyAppID ||
		preview.DesktopReceiptFieldsReconstructed ||
		preview.DesktopKDEStateRootAccess ||
		preview.DesktopOwnerInputsExposed ||
		!preview.GUIEvidenceRequired ||
		!preview.GUIEvidenceConsumed ||
		!preview.WindowObservationRequired ||
		!preview.OwnerFileOpenRequired ||
		!preview.OwnerFileOpenVerified ||
		!preview.Q4ExecutionRequired ||
		!preview.Q4ExecutionPlanned ||
		preview.Q4ExecutionStarted ||
		!preview.HostCompilationAvoided ||
		preview.BackendLaunchEnabled ||
		preview.RawOutputExposed ||
		preview.RemotePathExposed {
		t.Fatalf("unexpected GUI run plan: %#v", preview)
	}
}

func TestPreviewKnownAppVerifiedCatalogRunPlanSelectsQ4RunnableApp(t *testing.T) {
	catalog, err := PreviewKnownAppVerifiedCatalogJSON(knownAppVerifiedCatalogMatrixEvidenceFixture(t))
	if err != nil {
		t.Fatalf("PreviewKnownAppVerifiedCatalogJSON returned error: %v", err)
	}
	content, err := json.Marshal(catalog)
	if err != nil {
		t.Fatalf("Marshal catalog returned error: %v", err)
	}

	preview, err := PreviewKnownAppVerifiedCatalogRunPlanJSON(content, "busybox-w32")
	if err != nil {
		t.Fatalf("PreviewKnownAppVerifiedCatalogRunPlanJSON returned error: %v", err)
	}
	if preview.SchemaVersion != KnownAppVerifiedCatalogRunPlanSchemaVersion ||
		preview.RequestType != KnownAppVerifiedCatalogRunPlanRequestType ||
		preview.Source != "known-app-verified-catalog+q4-run-plan" ||
		preview.RuntimeMethod != "PlanKnownVerifiedApplicationRun" ||
		preview.ReadMethod != "GetKnownVerifiedApplicationRunPlan" ||
		!preview.VerifiedCatalogConsumed ||
		preview.RequestedAppID != "busybox-w32" ||
		preview.AppID != "busybox-w32" ||
		preview.VerificationState != "verified-real-q4-matrix-run" ||
		preview.DesktopCatalogState != "visible-review-only" {
		t.Fatalf("unexpected run plan schema: %#v", preview)
	}
	if strings.Join(preview.LaunchRequestCommand, " ") != "xnix-compat-launch --app busybox-w32" ||
		strings.Join(preview.RemoteSmokeCommand, " ") != "ruby scripts/remote_known_winapp_guest_wine_smoke.rb --execute --app busybox-w32" ||
		preview.RemoteSmokeRequestType != "remote-known-winapp-guest-wine-smoke" {
		t.Fatalf("unexpected run commands: %#v", preview)
	}
	if !preview.RuntimeOwnedActionReady ||
		preview.DesktopCallableActionID != "review-known-app-matrix-evidence" ||
		preview.DesktopCallableRoute != "runtime-owner://known-app-verified-catalog/run-plan" ||
		preview.DesktopCallableRuntimeMethod != "PlanKnownVerifiedApplicationRun" ||
		preview.DesktopCallableExecutionType != "review-only-q4-known-app-smoke" ||
		strings.Join(preview.DesktopForwardedArguments, " ") != "busybox-w32" ||
		!preview.DesktopForwardsOnlyAppID ||
		preview.DesktopReceiptFieldsReconstructed ||
		preview.DesktopKDEStateRootAccess ||
		preview.DesktopOwnerInputsExposed ||
		preview.GUIEvidenceRequired ||
		preview.GUIEvidenceConsumed ||
		preview.WindowObservationRequired ||
		preview.OwnerFileOpenRequired ||
		preview.OwnerFileOpenVerified {
		t.Fatalf("unexpected desktop action handoff fields: %#v", preview)
	}
	if !preview.Q4ExecutionRequired ||
		!preview.Q4ExecutionPlanned ||
		preview.Q4ExecutionStarted ||
		!preview.ReviewOnly ||
		!preview.OperatorReviewRequired ||
		!preview.RuntimeOwned ||
		!preview.GoRuntimeBacked ||
		preview.KDEPolicyOwner ||
		preview.DirectLaunchEnabled ||
		preview.LaunchEnabled ||
		preview.ExecutionStarted ||
		preview.BackendLaunchEnabled ||
		preview.DesktopFilesWritten ||
		preview.HostRootModified ||
		preview.BackendDetailsExposed ||
		preview.RawOutputExposed ||
		preview.RemotePathExposed ||
		preview.PrivilegedContainerRequired ||
		preview.HostNetworkingRequired ||
		preview.DockerSocketMounted ||
		preview.BroadHostMountRequired ||
		preview.HostCompilationRequired ||
		!preview.HostCompilationAvoided ||
		!preview.TargetedRemoteVerificationReady {
		t.Fatalf("unexpected run plan safety flags: %#v", preview)
	}
}

func TestPreviewKnownAppVerifiedCatalogRunPlanRejectsUnknownApp(t *testing.T) {
	catalog, err := PreviewKnownAppVerifiedCatalogJSON(knownAppVerifiedCatalogMatrixEvidenceFixture(t))
	if err != nil {
		t.Fatalf("PreviewKnownAppVerifiedCatalogJSON returned error: %v", err)
	}
	content, err := json.Marshal(catalog)
	if err != nil {
		t.Fatalf("Marshal catalog returned error: %v", err)
	}

	_, err = PreviewKnownAppVerifiedCatalogRunPlanJSON(content, "missing-app")
	if err == nil || !strings.Contains(err.Error(), "missing-app") {
		t.Fatalf("expected unknown app error, got %v", err)
	}
}

func TestPreviewKnownAppVerifiedCatalogRunAcceptanceConsumesMatchedRun(t *testing.T) {
	catalog, err := PreviewKnownAppVerifiedCatalogJSON(knownAppVerifiedCatalogMatrixEvidenceFixture(t))
	if err != nil {
		t.Fatalf("PreviewKnownAppVerifiedCatalogJSON returned error: %v", err)
	}
	catalogContent, err := json.Marshal(catalog)
	if err != nil {
		t.Fatalf("Marshal catalog returned error: %v", err)
	}
	runPlan, err := PreviewKnownAppVerifiedCatalogRunPlanJSON(catalogContent, "7zr")
	if err != nil {
		t.Fatalf("PreviewKnownAppVerifiedCatalogRunPlanJSON returned error: %v", err)
	}
	runPlanContent, err := json.Marshal(runPlan)
	if err != nil {
		t.Fatalf("Marshal run plan returned error: %v", err)
	}

	preview, err := PreviewKnownAppVerifiedCatalogRunAcceptanceJSON(runPlanContent, []byte(knownExistingWinAppAcceptanceFixture(currentProjectVersion(t), "7zr")))
	if err != nil {
		t.Fatalf("PreviewKnownAppVerifiedCatalogRunAcceptanceJSON returned error: %v", err)
	}
	if preview.SchemaVersion != KnownAppVerifiedCatalogRunAcceptanceSchemaVersion ||
		preview.RequestType != KnownAppVerifiedCatalogRunAcceptanceRequestType ||
		preview.AcceptanceType != "verified-catalog-app-q4-real-run-acceptance" ||
		!preview.RunPlanConsumed ||
		!preview.RunReportConsumed ||
		preview.RunPlanPathExposed ||
		preview.RunReportPathExposed ||
		preview.RemoteHostExposed ||
		preview.RawPathExposed ||
		preview.RawOutputExposed ||
		preview.RuntimeArgvExposed ||
		preview.RunnerPathExposed ||
		preview.RequestedAppID != "7zr" ||
		preview.AppID != "7zr" ||
		preview.VerificationState != "verified-real-q4-matrix-run" ||
		preview.DesktopCatalogState != "visible-review-only" ||
		!preview.RunPlanMatched {
		t.Fatalf("unexpected run acceptance schema: %#v", preview)
	}
	if !preview.ExistingWindowsApp ||
		!preview.KnownPortableCatalogBacked ||
		!preview.LaunchAttempted ||
		!preview.ChecksumVerified ||
		!preview.MarkerObserved ||
		!preview.RuntimeStartedIsolatedGuest ||
		!preview.IsolatedGuestExecutionObserved ||
		!preview.CompatibilityEngineExecutionObserved ||
		!preview.LoopbackOnlyNetworking ||
		!preview.SerialLogPersisted ||
		!preview.OutputRedacted ||
		!preview.Q4ExecutionObserved ||
		!preview.HostCompilationAvoided ||
		preview.NetworkRequired ||
		preview.HostRootModified ||
		preview.PrivilegedContainerRequired ||
		preview.HostNetworkingRequired ||
		preview.DockerSocketMounted ||
		preview.BroadHostMountRequired ||
		preview.DockerExecuted ||
		preview.ColimaExecuted ||
		preview.NetworkChecksRun ||
		preview.PackageManagerInvoked ||
		!preview.AcceptanceReady {
		t.Fatalf("unexpected run acceptance safety flags: %#v", preview)
	}
}

func TestPreviewKnownAppVerifiedCatalogRunAcceptanceRejectsMismatchedRun(t *testing.T) {
	catalog, err := PreviewKnownAppVerifiedCatalogJSON(knownAppVerifiedCatalogMatrixEvidenceFixture(t))
	if err != nil {
		t.Fatalf("PreviewKnownAppVerifiedCatalogJSON returned error: %v", err)
	}
	catalogContent, err := json.Marshal(catalog)
	if err != nil {
		t.Fatalf("Marshal catalog returned error: %v", err)
	}
	runPlan, err := PreviewKnownAppVerifiedCatalogRunPlanJSON(catalogContent, "busybox-w32")
	if err != nil {
		t.Fatalf("PreviewKnownAppVerifiedCatalogRunPlanJSON returned error: %v", err)
	}
	runPlanContent, err := json.Marshal(runPlan)
	if err != nil {
		t.Fatalf("Marshal run plan returned error: %v", err)
	}

	_, err = PreviewKnownAppVerifiedCatalogRunAcceptanceJSON(runPlanContent, []byte(knownExistingWinAppAcceptanceFixture(currentProjectVersion(t), "7zr")))
	if err == nil || !strings.Contains(err.Error(), "app ids to match") {
		t.Fatalf("expected mismatched app error, got %v", err)
	}
}

func knownAppVerifiedCatalogGUIEvidencePacketFixture() string {
	return `{
  "version": "0.2.640-test",
  "schema_version": "xnix.runtime.real_winapp_gui_evidence_packet.v1",
  "request_type": "real-winapp-gui-evidence-packet-preview",
  "packet_type": "real-windows-app-gui-evidence",
  "source": "wine-guest-gui-smoke+runtime-evidence-consumer+real-winapp-desktop-packet",
  "runtime_method": "PreviewRealWinAppGUIEvidencePacket",
  "read_method": "GetRealWinAppGUIEvidencePacket",
  "report_status": "passed",
  "report_consumed": true,
  "report_path_exposed": false,
  "app_id": "org.xnix.apps.messagebox",
  "display_name": "Xnix MessageBox",
  "app_version": "0.2.640-test",
  "gui_app_name": "xnix-messagebox-smoke.exe",
  "evidence_source": "wine-guest-gui-smoke",
  "compatibility_state": "owner-controlled-gui-qemu-wine-verified",
  "center_card_state": "validated-owner-controlled-gui-runtime-run",
  "known_app_gui_evidence_count": 1,
  "known_app_gui_evidence_verified_count": 1,
  "known_app_smoke_evidence": {
    "app_id": "org.xnix.apps.messagebox",
    "display_name": "Xnix MessageBox",
    "app_version": "0.2.640-test",
    "evidence_kind": "known-application-gui-smoke",
    "evidence_source": "wine-guest-gui-smoke",
    "smoke_status": "passed",
    "x_window_observed": true,
    "window_observed": true,
    "compatibility_state": "owner-controlled-gui-qemu-wine-verified",
    "center_card_state": "validated-owner-controlled-gui-runtime-run",
    "owner_file_open_verified": true,
    "owner_file_open_entrypoint_invoked": true,
    "owner_delegated_file_argument_count": 1,
    "owner_delegated_file_argument_copied_count": 1,
    "owner_delegated_file_arguments_passed": true,
    "owner_delegated_file_argument_winepath_translated": true,
    "owner_delegated_file_argument_winepath_translated_count": 1,
    "owner_delegated_raw_file_argument_path_exposed": false,
    "owner_delegated_window_match_observed": true,
    "marker_observed": true,
    "checksum_verified": false,
    "execution_evidence_recorded": true,
    "runtime_dispatch_verified": true,
    "launch_authorization_required": true,
    "desktop_launch_enabled": false,
    "runtime_owned": true,
    "kde_policy_owner": false,
    "action_execution_enabled": false,
    "backend_launch_enabled": false,
    "host_root_modified": false,
    "backend_details_exposed": false,
    "raw_artifact_path_exposed": false,
    "summary": "Xnix MessageBox passed owner-controlled GUI evidence."
  },
  "wineboot_invoked": true,
  "x_window_observed": true,
  "window_observed": true,
  "x_window_child_count": 1,
  "compatibility_center_projection_ready": true,
  "kde_center_projection_ready": true,
  "container_runtime_used": false,
  "container_host_mount_count": 0,
  "runtime_owned": true,
  "go_runtime_backed": true,
  "kde_policy_owner": false,
  "desktop_launch_enabled": false,
  "backend_launch_enabled": false,
  "action_execution_enabled": false,
  "backend_details_exposed": false,
  "raw_output_exposed": false,
  "host_root_modified": false,
  "privileged_container_required": false,
  "host_networking_required": false,
  "docker_socket_mounted": false,
  "broad_host_mount_required": false,
  "desktop_safe_summary": "Xnix MessageBox produced owner-controlled GUI evidence."
}`
}

func TestKnownAppSmokeEvidenceFromKnownAppVerifiedCatalogRunAcceptance(t *testing.T) {
	catalog, err := PreviewKnownAppVerifiedCatalogJSON(knownAppVerifiedCatalogMatrixEvidenceFixture(t))
	if err != nil {
		t.Fatalf("PreviewKnownAppVerifiedCatalogJSON returned error: %v", err)
	}
	catalogContent, err := json.Marshal(catalog)
	if err != nil {
		t.Fatalf("Marshal catalog returned error: %v", err)
	}
	runPlan, err := PreviewKnownAppVerifiedCatalogRunPlanJSON(catalogContent, "7zr")
	if err != nil {
		t.Fatalf("PreviewKnownAppVerifiedCatalogRunPlanJSON returned error: %v", err)
	}
	runPlanContent, err := json.Marshal(runPlan)
	if err != nil {
		t.Fatalf("Marshal run plan returned error: %v", err)
	}
	acceptance, err := PreviewKnownAppVerifiedCatalogRunAcceptanceJSON(runPlanContent, []byte(knownExistingWinAppAcceptanceFixture(currentProjectVersion(t), "7zr")))
	if err != nil {
		t.Fatalf("PreviewKnownAppVerifiedCatalogRunAcceptanceJSON returned error: %v", err)
	}
	acceptanceContent, err := json.Marshal(acceptance)
	if err != nil {
		t.Fatalf("Marshal acceptance returned error: %v", err)
	}

	evidence, err := KnownAppSmokeEvidenceFromKnownAppVerifiedCatalogRunAcceptance(acceptanceContent)
	if err != nil {
		t.Fatalf("KnownAppSmokeEvidenceFromKnownAppVerifiedCatalogRunAcceptance returned error: %v", err)
	}
	if evidence.AppID != "7zr" ||
		evidence.EvidenceKind != "known-application-verified-catalog-run-acceptance" ||
		evidence.EvidenceSource != "verified-catalog-app-q4-real-run-acceptance" ||
		evidence.CompatibilityState != "verified-catalog-real-q4-run-accepted" ||
		evidence.CenterCardState != "validated-verified-catalog-real-runtime-run" ||
		evidence.PrimaryActionID != "review-known-app-verified-catalog-run-acceptance" ||
		evidence.PrimaryActionKind != "review" ||
		evidence.SmokeStatus != "passed" ||
		!evidence.MarkerObserved ||
		!evidence.ChecksumVerified ||
		!evidence.ExecutionEvidenceRecorded ||
		!evidence.RuntimeDispatchVerified ||
		!evidence.LaunchAuthorizationRequired ||
		evidence.DesktopLaunchEnabled ||
		evidence.BackendLaunchEnabled ||
		evidence.HostRootModified ||
		evidence.BackendDetailsExposed ||
		evidence.RawArtifactPathExposed {
		t.Fatalf("unexpected accepted catalog run evidence: %#v", evidence)
	}
}

func TestKnownAppSmokeEvidenceFromKnownAppVerifiedCatalogRunAcceptanceRejectsUnsafePayload(t *testing.T) {
	catalog, err := PreviewKnownAppVerifiedCatalogJSON(knownAppVerifiedCatalogMatrixEvidenceFixture(t))
	if err != nil {
		t.Fatalf("PreviewKnownAppVerifiedCatalogJSON returned error: %v", err)
	}
	catalogContent, err := json.Marshal(catalog)
	if err != nil {
		t.Fatalf("Marshal catalog returned error: %v", err)
	}
	runPlan, err := PreviewKnownAppVerifiedCatalogRunPlanJSON(catalogContent, "7zr")
	if err != nil {
		t.Fatalf("PreviewKnownAppVerifiedCatalogRunPlanJSON returned error: %v", err)
	}
	runPlanContent, err := json.Marshal(runPlan)
	if err != nil {
		t.Fatalf("Marshal run plan returned error: %v", err)
	}
	acceptance, err := PreviewKnownAppVerifiedCatalogRunAcceptanceJSON(runPlanContent, []byte(knownExistingWinAppAcceptanceFixture(currentProjectVersion(t), "7zr")))
	if err != nil {
		t.Fatalf("PreviewKnownAppVerifiedCatalogRunAcceptanceJSON returned error: %v", err)
	}
	acceptance.RawOutputExposed = true
	acceptanceContent, err := json.Marshal(acceptance)
	if err != nil {
		t.Fatalf("Marshal acceptance returned error: %v", err)
	}

	_, err = KnownAppSmokeEvidenceFromKnownAppVerifiedCatalogRunAcceptance(acceptanceContent)
	if err == nil || !strings.Contains(err.Error(), "unsafe") {
		t.Fatalf("expected unsafe acceptance error, got %v", err)
	}
}

func knownAppVerifiedCatalogMatrixEvidenceFixture(t *testing.T) []byte {
	t.Helper()
	apps := []KnownAppMatrixEvidenceApp{
		knownAppVerifiedCatalogMatrixEvidenceApp("7zr", "7-Zip standalone console executable", "26.02"),
		knownAppVerifiedCatalogMatrixEvidenceApp("busybox-w32", "BusyBox-w32 standalone console executable", "current-2026-07-24"),
	}
	evidence := KnownAppMatrixEvidencePreview{
		SchemaVersion:                      KnownAppMatrixEvidencePreviewSchemaVersion,
		RequestType:                        KnownAppMatrixEvidencePreviewRequestType,
		Source:                             "remote-known-winapp-matrix-smoke+runtime-evidence-consumer",
		RuntimeMethod:                      "PreviewKnownAppMatrixEvidence",
		ReadMethod:                         "GetKnownAppMatrixEvidence",
		MatrixStatus:                       "passed",
		MatrixReportConsumed:               true,
		MatrixReportPathExposed:            false,
		MatrixReportOutputWritten:          true,
		AppCount:                           2,
		PassedCount:                        2,
		FailedCount:                        0,
		EvidenceCount:                      2,
		PassedEvidenceCount:                2,
		FailedEvidenceCount:                0,
		QEMUExecutedCount:                  2,
		WineExecutedCount:                  2,
		MarkerObservedCount:                2,
		ChecksumVerifiedCount:              2,
		RawOutputRedactedCount:             2,
		SerialLogEvidenceCount:             2,
		GuestStartedCount:                  2,
		GuestPortAutoCount:                 2,
		CompatibilityCenterProjectionReady: true,
		KDECenterProjectionReady:           true,
		Apps:                               apps,
		RuntimeOwned:                       true,
		GoRuntimeBacked:                    true,
		KDEPolicyOwner:                     false,
		DesktopLaunchEnabled:               false,
		BackendLaunchEnabled:               false,
		ActionExecutionEnabled:             false,
		BackendDetailsExposed:              false,
		RawOutputExposed:                   false,
		RemotePathExposed:                  false,
		HostRootModified:                   false,
		PrivilegedContainerRequired:        false,
		HostNetworkingRequired:             false,
		DockerSocketMounted:                false,
		BroadHostMountRequired:             false,
	}
	content, err := json.Marshal(evidence)
	if err != nil {
		t.Fatalf("Marshal returned error: %v", err)
	}
	return content
}

func knownAppVerifiedCatalogMatrixEvidenceApp(appID string, displayName string, appVersion string) KnownAppMatrixEvidenceApp {
	return KnownAppMatrixEvidenceApp{
		AppID:                 appID,
		DisplayName:           displayName,
		AppVersion:            appVersion,
		SmokeStatus:           "passed",
		CompatibilityState:    "real-qemu-wine-verified",
		MarkerObserved:        true,
		ChecksumVerified:      true,
		QEMUExecuted:          true,
		WineExecuted:          true,
		GuestStarted:          true,
		GuestPortAuto:         true,
		RawOutputRedacted:     true,
		SerialLogEvidence:     true,
		ReportEvidence:        true,
		RuntimeOwned:          true,
		GoRuntimeBacked:       true,
		KDEPolicyOwner:        false,
		DesktopLaunchEnabled:  false,
		BackendLaunchEnabled:  false,
		BackendDetailsExposed: false,
		RawOutputExposed:      false,
		RemotePathExposed:     false,
		HostRootModified:      false,
	}
}
