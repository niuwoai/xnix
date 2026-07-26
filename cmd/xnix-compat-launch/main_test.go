package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"xnix.local/xnix/internal/runtime/activation"
	"xnix.local/xnix/internal/runtime/appidentity"
	"xnix.local/xnix/internal/runtime/execution"
	"xnix.local/xnix/internal/runtime/winapp"
)

func TestCompatLaunchUsesKnownAppLaunchBridge(t *testing.T) {
	tempDir := t.TempDir()

	var output bytes.Buffer
	err := run([]string{
		"--app", "7zr",
		"--cache-root", tempDir,
	}, &output)
	if err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["schema_version"] != "xnix.runtime.known_windows_app_launch_bridge.v1" ||
		payload["request_type"] != "windows-known-app-launch-bridge-preview" ||
		payload["source"] != "windows-known-app-dispatch-preview" ||
		payload["status"] != "bridge-blocked" ||
		payload["managed_launcher"] != "xnix-compat-launch --app 7zr" ||
		payload["launcher_argv_accepted"] != true ||
		payload["request_id"] != "known-app-launch-request-7zr" ||
		payload["dispatch_id"] != "known-app-dispatch-7zr" ||
		payload["runtime_method"] != "BridgeKnownLauncherToDispatchSmoke" ||
		payload["dispatch_smoke_request_type"] != "windows-known-app-dispatch-smoke" ||
		payload["dispatch_smoke_request_materialized"] != false ||
		payload["app_id"] != "7zr" ||
		payload["dispatch_gate"] != "managed-known-app-guest-smoke" ||
		payload["guest_boundary_required"] != true ||
		payload["guest_boundary_supplied"] != false ||
		payload["smoke_harness_required"] != true ||
		payload["cache_status"] != "missing" ||
		payload["artifact_verified"] != false ||
		payload["bridge_preview_created"] != true ||
		payload["dispatch_ready"] != false ||
		payload["preparation_required"] != true ||
		payload["runtime_owned_bridge"] != true ||
		payload["kde_presentation_only"] != true ||
		payload["dry_run"] != true ||
		payload["dispatch_started"] != false ||
		payload["execution_started"] != false ||
		payload["backend_process_started"] != false ||
		payload["host_root_modified"] != false ||
		payload["host_networking_required"] != false ||
		payload["docker_socket_mounted"] != false ||
		payload["broad_host_mount_required"] != false ||
		payload["raw_host_path_exposed"] != false ||
		payload["raw_executable_path_exposed"] != false ||
		payload["raw_command_exposed"] != false ||
		payload["backend_details_exposed"] != false {
		t.Fatalf("unexpected launch bridge payload: %#v", payload)
	}
	assertCompatLaunchCLISafe(t, output.String(), tempDir)
}

func TestCompatLaunchRequiresApp(t *testing.T) {
	var output bytes.Buffer
	err := run(nil, &output)
	if err == nil || !strings.Contains(err.Error(), "--app is required") {
		t.Fatalf("expected missing app rejection, got %v", err)
	}
}

func TestCompatLaunchRunsExternalImportedAppThroughRuntimeRun(t *testing.T) {
	tempDir := t.TempDir()
	executablePath := filepath.Join(tempDir, "ExternalGui.exe")
	if err := os.WriteFile(executablePath, []byte{'M', 'Z', 0x90, 0x00, 'x', 'n', 'i', 'x'}, 0o600); err != nil {
		t.Fatalf("WriteFile executable returned error: %v", err)
	}
	stateRoot := filepath.Join(tempDir, "state")
	record, err := appidentity.RecordExternalWinAppImport(appidentity.ExternalWinAppImportRequest{
		Version:        "0.2.640-test",
		StateRoot:      stateRoot,
		ExecutablePath: executablePath,
		AppID:          "org.xnix.external.gui",
		DisplayName:    "External GUI",
		AppVersion:     "0.2.640-test",
	})
	if err != nil {
		t.Fatalf("RecordExternalWinAppImport returned error: %v", err)
	}
	recordPath := filepath.Join(stateRoot, filepath.FromSlash(record.RecordRelativePath))
	dockerPath, dockerLog := writeExternalImportedFakeDocker(t, tempDir)

	var output bytes.Buffer
	err = run([]string{
		"--external-app-import-record", recordPath,
		"--image", "local/wine-x-gui:test",
		"--platform", "linux/amd64",
		"--docker", dockerPath,
		"--timeout", "5s",
	}, &output)
	if err != nil {
		t.Fatalf("run returned error: %v", err)
	}
	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["schema_version"] != appidentity.ExternalWinAppRunSchemaVersion ||
		payload["request_type"] != appidentity.ExternalWinAppRunRequestType ||
		payload["status"] != "passed" ||
		payload["application_id"] != "org.xnix.external.gui" ||
		payload["display_name"] != "External GUI" ||
		payload["app_version"] != "0.2.640-test" ||
		payload["executable_name"] != "ExternalGui.exe" ||
		payload["external_app_import_record_consumed"] != true ||
		payload["imported_artifact_digest_verified"] != true ||
		payload["imported_artifact_sha256"] != record.ArtifactSHA256 ||
		payload["runtime_run_requested"] != true ||
		payload["runtime_run_executed"] != true ||
		payload["execution_started"] != true ||
		payload["backend_process_started"] != true ||
		payload["container_runtime_used"] != true ||
		payload["container_network_mode"] != "none" ||
		payload["container_host_mount_count"] != float64(0) ||
		payload["x_window_observed"] != true ||
		payload["window_observed"] != true ||
		payload["desktop_launch_enabled"] != false ||
		payload["action_execution_enabled"] != false ||
		payload["backend_details_exposed"] != false ||
		payload["raw_import_record_path_exposed"] != false ||
		payload["raw_state_root_path_exposed"] != false ||
		payload["raw_executable_path_exposed"] != false ||
		payload["host_root_modified"] != false ||
		payload["docker_socket_mounted"] != false {
		t.Fatalf("unexpected external imported launcher payload: %#v", payload)
	}
	runtimePayload := payload["runtime_payload"].(map[string]any)
	if runtimePayload["application_name"] != "/ExternalGui.exe" ||
		runtimePayload["local_executable_copied"] != true ||
		runtimePayload["external_app_import_record_consumed"] != true ||
		runtimePayload["imported_artifact_digest_verified"] != true {
		t.Fatalf("unexpected nested external imported runtime payload: %#v", runtimePayload)
	}
	dockerInvocation, err := os.ReadFile(dockerLog)
	if err != nil {
		t.Fatalf("ReadFile docker log returned error: %v", err)
	}
	if !strings.Contains(string(dockerInvocation), "cp") ||
		!strings.Contains(string(dockerInvocation), "ExternalGui.exe") {
		t.Fatalf("fake Docker did not receive copied imported executable flow: %s", string(dockerInvocation))
	}
	assertCompatLaunchContainerGUIDispatchSafe(t, output.String(), recordPath, stateRoot, executablePath, dockerPath)
}

func TestCompatLaunchRunsExternalImportedAppHandleThroughRuntimeRun(t *testing.T) {
	tempDir := t.TempDir()
	executablePath := filepath.Join(tempDir, "ExternalGui.exe")
	if err := os.WriteFile(executablePath, []byte{'M', 'Z', 0x90, 0x00, 'x', 'n', 'i', 'x'}, 0o600); err != nil {
		t.Fatalf("WriteFile executable returned error: %v", err)
	}
	stateRoot := filepath.Join(tempDir, "state")
	record, err := appidentity.RecordExternalWinAppImport(appidentity.ExternalWinAppImportRequest{
		Version:        "0.2.640-test",
		StateRoot:      stateRoot,
		ExecutablePath: executablePath,
		AppID:          "org.xnix.external.gui",
		DisplayName:    "External GUI",
		AppVersion:     "0.2.640-test",
	})
	if err != nil {
		t.Fatalf("RecordExternalWinAppImport returned error: %v", err)
	}
	recipe, provenance, err := appidentity.ExternalAppRecipeFromImportRecord(record)
	if err != nil {
		t.Fatalf("ExternalAppRecipeFromImportRecord returned error: %v", err)
	}
	plan, err := appidentity.NewPlanWithProvenance(recipe, provenance)
	if err != nil {
		t.Fatalf("NewPlanWithProvenance returned error: %v", err)
	}
	stageRoot := filepath.Join(tempDir, "stage")
	if _, err := activation.Stage(activation.StageRequest{
		Root: stageRoot,
		Mode: "development",
		Plan: plan,
	}); err != nil {
		t.Fatalf("Stage returned error: %v", err)
	}
	dockerPath, dockerLog := writeExternalImportedFakeDocker(t, tempDir)
	launchPacketOutput := filepath.Join(tempDir, "sidecars", "external-launch-packet.json")
	t.Setenv(appidentity.ExternalWinAppStateRootEnv, stateRoot)
	t.Setenv(compatLaunchContainerImageEnv, "local/wine-x-gui:test")
	t.Setenv(compatLaunchContainerPlatformEnv, "linux/amd64")
	t.Setenv(compatLaunchDockerEnv, dockerPath)
	t.Setenv(compatLaunchTimeoutEnv, "5s")
	t.Setenv(externalDesktopActivationRootEnv, stageRoot)
	t.Setenv(externalDesktopLaunchPacketOutputEnv, launchPacketOutput)
	t.Setenv(externalDesktopLaunchPacketModeEnv, "development")

	var output bytes.Buffer
	err = run([]string{
		"--external-app-handle", "org.xnix.external.gui",
		"file:///home/alice/Documents/report.docx",
	}, &output)
	if err != nil {
		t.Fatalf("run returned error: %v", err)
	}
	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["schema_version"] != appidentity.ExternalWinAppRunSchemaVersion ||
		payload["request_type"] != appidentity.ExternalWinAppRunRequestType ||
		payload["status"] != "passed" ||
		payload["application_id"] != "org.xnix.external.gui" ||
		payload["external_app_import_record_consumed"] != true ||
		payload["external_app_handle_consumed"] != true ||
		payload["external_app_handle"] != "org.xnix.external.gui" ||
		payload["external_desktop_argument_count"] != float64(1) ||
		payload["external_file_uri_arguments_accepted"] != true ||
		payload["external_file_open_requested"] != true ||
		payload["external_file_bridge_mount_enabled"] != false ||
		payload["raw_file_uri_arguments_exposed"] != false ||
		payload["imported_artifact_digest_verified"] != true ||
		payload["imported_artifact_sha256"] != record.ArtifactSHA256 ||
		payload["raw_external_app_handle_path_exposed"] != false ||
		payload["raw_import_record_path_exposed"] != false ||
		payload["raw_state_root_path_exposed"] != false ||
		payload["raw_executable_path_exposed"] != false ||
		payload["container_network_mode"] != "none" ||
		payload["container_host_mount_count"] != float64(0) ||
		payload["x_window_observed"] != true ||
		payload["window_observed"] != true {
		t.Fatalf("unexpected external handle launcher payload: %#v", payload)
	}
	dockerInvocation, err := os.ReadFile(dockerLog)
	if err != nil {
		t.Fatalf("ReadFile docker log returned error: %v", err)
	}
	if !strings.Contains(string(dockerInvocation), "cp") ||
		!strings.Contains(string(dockerInvocation), "ExternalGui.exe") {
		t.Fatalf("fake Docker did not receive copied imported executable flow: %s", string(dockerInvocation))
	}
	packetBytes, err := os.ReadFile(launchPacketOutput)
	if err != nil {
		t.Fatalf("ReadFile launch packet sidecar returned error: %v", err)
	}
	var launchPacket map[string]any
	if err := json.Unmarshal(packetBytes, &launchPacket); err != nil {
		t.Fatalf("Unmarshal launch packet returned error: %v\n%s", err, string(packetBytes))
	}
	if launchPacket["schema_version"] != appidentity.DesktopExternalWinAppLaunchPacketSchemaVersion ||
		launchPacket["request_type"] != appidentity.DesktopExternalWinAppLaunchPacketRequestType ||
		launchPacket["status"] != "passed" ||
		launchPacket["application_id"] != "org.xnix.external.gui" ||
		launchPacket["external_app_handle"] != "org.xnix.external.gui" ||
		launchPacket["external_desktop_argument_count"] != float64(1) ||
		launchPacket["external_file_uri_arguments_accepted"] != true ||
		launchPacket["external_file_open_requested"] != true ||
		launchPacket["external_file_bridge_mount_enabled"] != false ||
		launchPacket["raw_file_uri_arguments_exposed"] != false ||
		launchPacket["activation_receipt_backed"] != true ||
		launchPacket["activation_receipt_safe_for_kde"] != true ||
		launchPacket["desktop_exec_uses_external_app_handle"] != true ||
		launchPacket["external_app_desktop_handle_ready"] != true ||
		launchPacket["run_record_consumed"] != true ||
		launchPacket["external_app_handle_consumed"] != true ||
		launchPacket["imported_artifact_digest_verified"] != true ||
		launchPacket["runtime_launch_executed"] != true ||
		launchPacket["window_observed"] != true ||
		launchPacket["x_window_observed"] != true ||
		launchPacket["desktop_launch_packet_ready"] != true ||
		launchPacket["safe_for_kde"] != true ||
		launchPacket["runtime_launch_authority"] != true ||
		launchPacket["kde_launch_authority"] != false ||
		launchPacket["raw_import_record_path_exposed"] != false ||
		launchPacket["raw_state_root_path_exposed"] != false ||
		launchPacket["raw_executable_path_exposed"] != false ||
		launchPacket["host_root_modified"] != false {
		t.Fatalf("unexpected launch packet sidecar: %#v", launchPacket)
	}
	assertCompatLaunchContainerGUIDispatchSafe(t, string(packetBytes), stageRoot, stateRoot, executablePath, dockerPath)
	if strings.Contains(output.String(), "report.docx") || strings.Contains(string(packetBytes), "report.docx") {
		t.Fatalf("external launcher exposed raw file URI argument\nstdout=%s\npacket=%s", output.String(), string(packetBytes))
	}
	assertCompatLaunchContainerGUIDispatchSafe(t, output.String(), stateRoot, executablePath, dockerPath)
}

func TestCompatLaunchExternalImportedAppRejectsUnsafeDesktopArgument(t *testing.T) {
	tempDir := t.TempDir()
	executablePath := filepath.Join(tempDir, "ExternalGui.exe")
	if err := os.WriteFile(executablePath, []byte{'M', 'Z', 0x90, 0x00, 'x', 'n', 'i', 'x'}, 0o600); err != nil {
		t.Fatalf("WriteFile executable returned error: %v", err)
	}
	stateRoot := filepath.Join(tempDir, "state")
	if _, err := appidentity.RecordExternalWinAppImport(appidentity.ExternalWinAppImportRequest{
		Version:        "0.2.640-test",
		StateRoot:      stateRoot,
		ExecutablePath: executablePath,
		AppID:          "org.xnix.external.gui",
		DisplayName:    "External GUI",
		AppVersion:     "0.2.640-test",
	}); err != nil {
		t.Fatalf("RecordExternalWinAppImport returned error: %v", err)
	}
	t.Setenv(appidentity.ExternalWinAppStateRootEnv, stateRoot)

	var output bytes.Buffer
	err := run([]string{
		"--external-app-handle", "org.xnix.external.gui",
		"http://example.test/report.docx",
	}, &output)
	if err == nil || !strings.Contains(err.Error(), "file URI") {
		t.Fatalf("expected unsafe desktop argument rejection, got err=%v output=%s", err, output.String())
	}
}

func TestCompatLaunchRunsExternalImportedAppHandleThroughRuntimeDefaultStateRoot(t *testing.T) {
	tempDir := t.TempDir()
	executablePath := filepath.Join(tempDir, "ExternalGui.exe")
	if err := os.WriteFile(executablePath, []byte{'M', 'Z', 0x90, 0x00, 'x', 'n', 'i', 'x'}, 0o600); err != nil {
		t.Fatalf("WriteFile executable returned error: %v", err)
	}
	stateRoot := filepath.Join(tempDir, "state")
	record, err := appidentity.RecordExternalWinAppImport(appidentity.ExternalWinAppImportRequest{
		Version:        "0.2.640-test",
		StateRoot:      stateRoot,
		ExecutablePath: executablePath,
		AppID:          "org.xnix.external.gui",
		DisplayName:    "External GUI",
		AppVersion:     "0.2.640-test",
	})
	if err != nil {
		t.Fatalf("RecordExternalWinAppImport returned error: %v", err)
	}
	t.Setenv(appidentity.ExternalWinAppStateRootEnv, stateRoot)
	dockerPath, _ := writeExternalImportedFakeDocker(t, tempDir)

	var output bytes.Buffer
	err = run([]string{
		"--external-app-handle", "org.xnix.external.gui",
		"--image", "local/wine-x-gui:test",
		"--platform", "linux/amd64",
		"--docker", dockerPath,
		"--timeout", "5s",
	}, &output)
	if err != nil {
		t.Fatalf("run returned error: %v", err)
	}
	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["status"] != "passed" ||
		payload["application_id"] != "org.xnix.external.gui" ||
		payload["external_app_handle_consumed"] != true ||
		payload["external_app_handle"] != "org.xnix.external.gui" ||
		payload["imported_artifact_sha256"] != record.ArtifactSHA256 ||
		payload["raw_external_app_handle_path_exposed"] != false ||
		payload["raw_state_root_path_exposed"] != false ||
		payload["x_window_observed"] != true ||
		payload["window_observed"] != true {
		t.Fatalf("unexpected external handle default state root payload: %#v", payload)
	}
	assertCompatLaunchContainerGUIDispatchSafe(t, output.String(), stateRoot, executablePath, dockerPath)
}

func TestCompatLaunchExternalImportedAppRejectsKnownAppCombination(t *testing.T) {
	var output bytes.Buffer
	err := run([]string{"--app", "7zr", "--external-app-import-record", "record.json"}, &output)
	if err == nil || !strings.Contains(err.Error(), "--external-app-import-record or --external-app-handle cannot be combined with --app") {
		t.Fatalf("expected external import record/app combination rejection, got %v", err)
	}
}

func TestCompatLaunchRequiresReceiptForDispatchBoundary(t *testing.T) {
	var output bytes.Buffer
	err := run([]string{
		"--app", "7zr",
		"--guest-boundary", "managed-known-app-guest-smoke",
	}, &output)
	if err == nil || !strings.Contains(err.Error(), "--state-root, --receipt-id, and --review-receipt-id are required") {
		t.Fatalf("expected missing receipt rejection, got %v", err)
	}
}

func TestCompatLaunchConsumesSessionGatedControlledDispatchRequestBeforeDispatch(t *testing.T) {
	cacheRoot := t.TempDir()
	stateRoot := t.TempDir()
	sessionID, _ := recordLauncherSessionGateFixture(t, stateRoot)
	reviewReceipt, err := appidentity.RecordKnownAppSessionGatedLaunchReviewReceipt(appidentity.KnownAppSessionGatedLaunchReviewReceiptRequest{
		AppID:     "7zr",
		StateRoot: stateRoot,
		SessionID: sessionID,
		ActionID:  appidentity.KnownAppSessionGatedLaunchReviewAction,
		Decision:  "approved",
	})
	if err != nil {
		t.Fatalf("RecordKnownAppSessionGatedLaunchReviewReceipt returned error: %v", err)
	}
	receipt, err := appidentity.RecordKnownAppLaunchAuthorizationReceipt(appidentity.KnownAppLaunchAuthorizationReceiptRequest{
		AppID:     "7zr",
		StateRoot: stateRoot,
		Authorize: appidentity.KnownAppLaunchAuthorizationReceiptAction,
	})
	if err != nil {
		t.Fatalf("RecordKnownAppLaunchAuthorizationReceipt returned error: %v", err)
	}

	var output bytes.Buffer
	err = run([]string{
		"--app", "7zr",
		"--cache-root", cacheRoot,
		"--guest-boundary", "managed-known-app-guest-smoke",
		"--state-root", stateRoot,
		"--receipt-id", receipt.ReceiptID,
		"--review-receipt-id", reviewReceipt.ReceiptID,
		"--session-id", sessionID,
	}, &output)
	if err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["schema_version"] != appidentity.KnownAppSessionGatedControlledDispatchSchemaVersion ||
		payload["request_type"] != appidentity.KnownAppSessionGatedControlledDispatchRequestType ||
		payload["source"] != appidentity.KnownAppSessionGatedLaunchReviewGateRequestType+"+"+appidentity.KnownAppControlledDispatchRequestType ||
		payload["execution_session_id"] != sessionID ||
		payload["review_receipt_id"] != reviewReceipt.ReceiptID ||
		payload["review_receipt_consumed"] != true ||
		payload["review_receipt_accepted"] != true ||
		payload["review_gate_ready"] != true ||
		payload["dispatch_state_advance_ready"] != true ||
		payload["launch_authorization_receipt_id"] != receipt.ReceiptID ||
		payload["launch_gate_receipt_accepted"] != true ||
		payload["launch_gate_guest_boundary_accepted"] != true ||
		payload["launch_gate_state"] != "dispatch-preparation-required" ||
		payload["controlled_dispatch_request_created"] != false ||
		payload["controlled_dispatch_request_state"] != "blocked" ||
		payload["request_objects_created"] != false ||
		payload["dispatch_started"] != false ||
		payload["execution_started"] != false ||
		payload["direct_launch_enabled"] != false ||
		payload["desktop_launch_enabled"] != false ||
		payload["backend_launch_enabled"] != false ||
		payload["backend_process_started"] != false ||
		payload["host_root_modified"] != false ||
		payload["review_receipt_path_exposed"] != false ||
		payload["receipt_path_exposed"] != false ||
		payload["state_root_path_exposed"] != false {
		t.Fatalf("unexpected session-gated controlled dispatch request payload: %#v", payload)
	}
	assertCompatLaunchCLISafe(t, output.String(), cacheRoot)
	assertCompatLaunchCLISafe(t, output.String(), stateRoot)
}

func TestCompatLaunchConsumesDigestVerifiedSessionGate(t *testing.T) {
	stateRoot := t.TempDir()
	sessionID, sessionRelativePath := recordLauncherSessionGateFixture(t, stateRoot)

	preview, err := consumeControlledExecutionSessionForLaunch("7zr", stateRoot, "")
	if err != nil {
		t.Fatalf("consumeControlledExecutionSessionForLaunch returned error: %v", err)
	}
	if preview.ExecutionSessionID != sessionID ||
		!preview.RecordConsumed ||
		!preview.LedgerRecordConsumed ||
		!preview.SessionRecordConsumed ||
		!preview.SessionDigestVerified ||
		preview.SessionRelativePath != sessionRelativePath ||
		!preview.RuntimeOwnerConsumable ||
		!preview.KDEReadModelConsumable ||
		preview.LiveStateObserved ||
		preview.SessionRegistered ||
		preview.WindowObserved ||
		preview.HostRootModified ||
		preview.BackendProcessStarted {
		t.Fatalf("unexpected launcher session gate preview: %#v", preview)
	}
}

func TestCompatLaunchRunsMinesThroughGuestGUIDispatchSmoke(t *testing.T) {
	app, err := winapp.LookupKnownPortableApp("org.xnix.apps.mines")
	if err != nil {
		t.Fatalf("LookupKnownPortableApp returned error: %v", err)
	}
	stateRoot := t.TempDir()
	sessionID, _ := recordLauncherSessionGateFixtureForApp(t, stateRoot, app.ID, app.Version)
	reviewReceipt, err := appidentity.RecordKnownAppSessionGatedLaunchReviewReceipt(appidentity.KnownAppSessionGatedLaunchReviewReceiptRequest{
		AppID:     app.ID,
		StateRoot: stateRoot,
		SessionID: sessionID,
		ActionID:  appidentity.KnownAppSessionGatedLaunchReviewAction,
		Decision:  "approved",
	})
	if err != nil {
		t.Fatalf("RecordKnownAppSessionGatedLaunchReviewReceipt returned error: %v", err)
	}
	receipt, err := appidentity.RecordKnownAppLaunchAuthorizationReceipt(appidentity.KnownAppLaunchAuthorizationReceiptRequest{
		AppID:     app.ID,
		StateRoot: stateRoot,
		Authorize: appidentity.KnownAppLaunchAuthorizationReceiptAction,
	})
	if err != nil {
		t.Fatalf("RecordKnownAppLaunchAuthorizationReceipt returned error: %v", err)
	}
	sshPath, xwininfoPath := writeGuestGUIFakeTools(t)

	var output bytes.Buffer
	err = run([]string{
		"--app", app.ID,
		"--cache-root", t.TempDir(),
		"--guest-boundary", "managed-known-app-guest-smoke",
		"--state-root", stateRoot,
		"--receipt-id", receipt.ReceiptID,
		"--review-receipt-id", reviewReceipt.ReceiptID,
		"--session-id", sessionID,
		"--ssh", sshPath,
		"--xwininfo", xwininfoPath,
		"--gui-app", "/tmp/xnix-wine-guest-gui-smoke/owner-messagebox.exe",
		"--guest-display", "10.0.2.2:127",
		"--host-display", ":127",
		"--timeout", "5s",
		"--gui-wait", (1 * time.Millisecond).String(),
	}, &output)
	if err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["schema_version"] != "xnix.runtime.known_windows_app_dispatch_smoke.v1" ||
		payload["request_type"] != "windows-known-app-dispatch-smoke" ||
		payload["evidence_source"] != "wine-guest-gui-smoke" ||
		payload["status"] != "passed" ||
		payload["app_id"] != app.ID ||
		payload["cache_status"] != "guest-builtin-gui" ||
		payload["artifact_verified"] != false ||
		payload["marker_observed"] != false ||
		payload["smoke_passed"] != true ||
		payload["execution_started"] != true ||
		payload["managed_guest_runner_invoked"] != true ||
		payload["managed_guest_reachable"] != true ||
		payload["managed_guest_runtime_ready"] != true ||
		payload["session_gated_controlled_dispatch_consumed"] != true ||
		payload["session_gated_controlled_dispatch_state"] != "created-after-session-gated-review" ||
		payload["controlled_execution_session_consumed"] != true ||
		payload["controlled_execution_session_id"] != sessionID ||
		payload["controlled_session_window_observed"] != true ||
		payload["controlled_session_host_root_modified"] != false ||
		payload["controlled_session_backend_process_start"] != false ||
		payload["host_root_modified"] != false ||
		payload["backend_process_started"] != false ||
		payload["raw_command_exposed"] != false ||
		payload["backend_details_exposed"] != false {
		t.Fatalf("unexpected Mines GUI dispatch payload: %#v", payload)
	}
	assertCompatLaunchGUIDispatchSafe(t, output.String(), stateRoot, sshPath, xwininfoPath)
}

func TestCompatLaunchCopiesOwnerSuppliedGUIExecutable(t *testing.T) {
	app, err := winapp.LookupKnownPortableApp("org.xnix.apps.messagebox")
	if err != nil {
		t.Fatalf("LookupKnownPortableApp returned error: %v", err)
	}
	stateRoot := t.TempDir()
	sessionID, _ := recordLauncherSessionGateFixtureForApp(t, stateRoot, app.ID, app.Version)
	reviewReceipt, err := appidentity.RecordKnownAppSessionGatedLaunchReviewReceipt(appidentity.KnownAppSessionGatedLaunchReviewReceiptRequest{
		AppID:     app.ID,
		StateRoot: stateRoot,
		SessionID: sessionID,
		ActionID:  appidentity.KnownAppSessionGatedLaunchReviewAction,
		Decision:  "approved",
	})
	if err != nil {
		t.Fatalf("RecordKnownAppSessionGatedLaunchReviewReceipt returned error: %v", err)
	}
	receipt, err := appidentity.RecordKnownAppLaunchAuthorizationReceipt(appidentity.KnownAppLaunchAuthorizationReceiptRequest{
		AppID:     app.ID,
		StateRoot: stateRoot,
		Authorize: appidentity.KnownAppLaunchAuthorizationReceiptAction,
	})
	if err != nil {
		t.Fatalf("RecordKnownAppLaunchAuthorizationReceipt returned error: %v", err)
	}
	sshPath, xwininfoPath := writeGuestGUIFakeTools(t)
	scpPath, scpLogPath := writeGuestGUIFakeSCPTool(t)
	executablePath := filepath.Join(t.TempDir(), "owner-messagebox.exe")
	if err := os.WriteFile(executablePath, []byte("MZ"), 0o600); err != nil {
		t.Fatalf("WriteFile executable returned error: %v", err)
	}

	var output bytes.Buffer
	err = run([]string{
		"--app", app.ID,
		"--cache-root", t.TempDir(),
		"--guest-boundary", "managed-known-app-guest-smoke",
		"--state-root", stateRoot,
		"--receipt-id", receipt.ReceiptID,
		"--review-receipt-id", reviewReceipt.ReceiptID,
		"--session-id", sessionID,
		"--ssh", sshPath,
		"--scp", scpPath,
		"--xwininfo", xwininfoPath,
		"--executable", executablePath,
		"--guest-display", "10.0.2.2:127",
		"--host-display", ":127",
		"--timeout", "5s",
		"--gui-wait", (1 * time.Millisecond).String(),
	}, &output)
	if err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["status"] != "passed" ||
		payload["app_id"] != app.ID ||
		payload["evidence_source"] != "wine-guest-gui-smoke" ||
		payload["managed_artifact_copied"] != true ||
		payload["smoke_passed"] != true ||
		payload["execution_started"] != true ||
		payload["controlled_session_window_observed"] != true ||
		payload["host_root_modified"] != false ||
		payload["raw_command_exposed"] != false ||
		payload["backend_details_exposed"] != false {
		t.Fatalf("unexpected owner-supplied executable GUI payload: %#v", payload)
	}
	scpLog, err := os.ReadFile(scpLogPath)
	if err != nil {
		t.Fatalf("ReadFile scp log returned error: %v", err)
	}
	if !strings.Contains(string(scpLog), executablePath) ||
		!strings.Contains(string(scpLog), "/tmp/xnix-known-winapp-smoke/owner-messagebox.exe") {
		t.Fatalf("fake scp did not copy the owner-supplied executable into the guest work dir: %s", string(scpLog))
	}
	assertCompatLaunchGUIDispatchSafe(t, output.String(), stateRoot, sshPath, scpPath, xwininfoPath, executablePath)
}

func TestCompatLaunchRunsRecipeBackedNotepadThroughContainerXGUI(t *testing.T) {
	app, err := winapp.LookupKnownPortableApp("org.xnix.sample.notepad")
	if err != nil {
		t.Fatalf("LookupKnownPortableApp returned error: %v", err)
	}
	if !app.RecipeBackedContainerGUI {
		t.Fatalf("sample Notepad must be marked as recipe-backed container GUI")
	}
	stateRoot := t.TempDir()
	sessionID, _ := recordLauncherSessionGateFixtureForApp(t, stateRoot, app.ID, app.Version)
	reviewReceipt, err := appidentity.RecordKnownAppSessionGatedLaunchReviewReceipt(appidentity.KnownAppSessionGatedLaunchReviewReceiptRequest{
		AppID:     app.ID,
		StateRoot: stateRoot,
		SessionID: sessionID,
		ActionID:  appidentity.KnownAppSessionGatedLaunchReviewAction,
		Decision:  "approved",
	})
	if err != nil {
		t.Fatalf("RecordKnownAppSessionGatedLaunchReviewReceipt returned error: %v", err)
	}
	receipt, err := appidentity.RecordKnownAppLaunchAuthorizationReceipt(appidentity.KnownAppLaunchAuthorizationReceiptRequest{
		AppID:     app.ID,
		StateRoot: stateRoot,
		Authorize: appidentity.KnownAppLaunchAuthorizationReceiptAction,
	})
	if err != nil {
		t.Fatalf("RecordKnownAppLaunchAuthorizationReceipt returned error: %v", err)
	}

	tempDir := t.TempDir()
	registryPath, recipePath := writeNotepadRecipeRegistryFixture(t, tempDir, app.Version)
	dockerPath, dockerLog := writeRecipeBackedFakeDocker(t, tempDir)

	var output bytes.Buffer
	err = run([]string{
		"--app", app.ID,
		"--cache-root", t.TempDir(),
		"--guest-boundary", winapp.KnownDispatchGuestBoundary,
		"--state-root", stateRoot,
		"--receipt-id", receipt.ReceiptID,
		"--review-receipt-id", reviewReceipt.ReceiptID,
		"--session-id", sessionID,
		"--registry", registryPath,
		"--image", "local/wine-x-gui:test",
		"--platform", "linux/amd64",
		"--docker", dockerPath,
		"--timeout", "5s",
	}, &output)
	if err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["schema_version"] != "xnix.runtime.windows_app_container_x_gui_smoke.v1" ||
		payload["request_type"] != "windows-app-container-x-gui-smoke" ||
		payload["status"] != "passed" ||
		payload["evidence_source"] != "winapp-smoke-container-x-gui" ||
		payload["application_id"] != app.ID ||
		payload["display_name"] != "Sample Notepad" ||
		payload["app_version"] != app.Version ||
		payload["recipe_backed"] != true ||
		payload["application_name"] != "notepad.exe" ||
		payload["window_match"] != "notepad.exe" ||
		payload["x_window_observed"] != true ||
		payload["dispatch_started"] != true ||
		payload["execution_started"] != true ||
		payload["smoke_passed"] != true ||
		payload["runtime_owned_dispatch"] != true ||
		payload["session_gated_controlled_dispatch_consumed"] != true ||
		payload["session_gated_controlled_dispatch_state"] != "created-after-session-gated-review" ||
		payload["controlled_execution_session_consumed"] != true ||
		payload["controlled_execution_session_id"] != sessionID ||
		payload["controlled_session_window_observed"] != true ||
		payload["controlled_session_host_root_modified"] != false ||
		payload["host_root_modified"] != false ||
		payload["network_mode"] != "none" ||
		payload["host_mount_count"] != float64(0) ||
		payload["docker_socket_mounted"] != false ||
		payload["host_networking_required"] != false ||
		payload["broad_host_mount_required"] != false ||
		payload["raw_command_exposed"] != false ||
		payload["backend_details_exposed"] != false {
		t.Fatalf("unexpected recipe-backed Notepad launcher payload: %#v", payload)
	}
	dockerInvocation, err := os.ReadFile(dockerLog)
	if err != nil {
		t.Fatalf("ReadFile docker log returned error: %v", err)
	}
	if !strings.Contains(string(dockerInvocation), "XNIX_GUI_APP=notepad.exe") ||
		!strings.Contains(string(dockerInvocation), "XNIX_WINDOW_MATCH=notepad.exe") {
		t.Fatalf("docker invocation did not receive recipe GUI hints: %s", string(dockerInvocation))
	}
	assertCompatLaunchContainerGUIDispatchSafe(t, output.String(), stateRoot, registryPath, recipePath, dockerPath)
}

func TestResolveStagedRegistryPathMapsPackagedRegistryUnderStagingRoot(t *testing.T) {
	stagingRoot := t.TempDir()
	packagedRegistryPath := packagedRecipeRegistryDir + "/registry.json"
	stagedRegistryDir := filepath.Join(stagingRoot, strings.TrimPrefix(packagedRecipeRegistryDir, "/"))
	if err := os.MkdirAll(stagedRegistryDir, 0o700); err != nil {
		t.Fatalf("MkdirAll staged registry dir returned error: %v", err)
	}
	stagedRegistryPath := filepath.Join(stagedRegistryDir, "registry.json")
	if err := os.WriteFile(stagedRegistryPath, []byte(`{"schema_version":1,"registry_name":"staged","recipes":[]}`), 0o600); err != nil {
		t.Fatalf("WriteFile staged registry returned error: %v", err)
	}

	t.Setenv(stagingRootEnv, stagingRoot)
	got, err := resolveStagedRegistryPath(packagedRegistryPath)
	if err != nil {
		t.Fatalf("resolveStagedRegistryPath returned error: %v", err)
	}
	if got != stagedRegistryPath {
		t.Fatalf("expected staged registry path %q, got %q", stagedRegistryPath, got)
	}

	otherPath := "/opt/xnix/registry.json"
	got, err = resolveStagedRegistryPath(otherPath)
	if err != nil {
		t.Fatalf("resolveStagedRegistryPath non-packaged path returned error: %v", err)
	}
	if got != otherPath {
		t.Fatalf("expected non-packaged path to remain unchanged, got %q", got)
	}
}

func TestCompatLaunchRunsRecipeBackedNotepadThroughStagedPackagedRegistry(t *testing.T) {
	app, err := winapp.LookupKnownPortableApp("org.xnix.sample.notepad")
	if err != nil {
		t.Fatalf("LookupKnownPortableApp returned error: %v", err)
	}
	stateRoot := t.TempDir()
	sessionID, _ := recordLauncherSessionGateFixtureForApp(t, stateRoot, app.ID, app.Version)
	reviewReceipt, err := appidentity.RecordKnownAppSessionGatedLaunchReviewReceipt(appidentity.KnownAppSessionGatedLaunchReviewReceiptRequest{
		AppID:     app.ID,
		StateRoot: stateRoot,
		SessionID: sessionID,
		ActionID:  appidentity.KnownAppSessionGatedLaunchReviewAction,
		Decision:  "approved",
	})
	if err != nil {
		t.Fatalf("RecordKnownAppSessionGatedLaunchReviewReceipt returned error: %v", err)
	}
	receipt, err := appidentity.RecordKnownAppLaunchAuthorizationReceipt(appidentity.KnownAppLaunchAuthorizationReceiptRequest{
		AppID:     app.ID,
		StateRoot: stateRoot,
		Authorize: appidentity.KnownAppLaunchAuthorizationReceiptAction,
	})
	if err != nil {
		t.Fatalf("RecordKnownAppLaunchAuthorizationReceipt returned error: %v", err)
	}

	stagingRoot := t.TempDir()
	stagedRegistryDir := filepath.Join(stagingRoot, strings.TrimPrefix(packagedRecipeRegistryDir, "/"))
	if err := os.MkdirAll(stagedRegistryDir, 0o700); err != nil {
		t.Fatalf("MkdirAll staged registry dir returned error: %v", err)
	}
	stagedRegistryPath, stagedRecipePath := writeNotepadRecipeRegistryFixture(t, stagedRegistryDir, app.Version)
	dockerPath, dockerLog := writeRecipeBackedFakeDocker(t, t.TempDir())
	t.Setenv(stagingRootEnv, stagingRoot)

	var output bytes.Buffer
	err = run([]string{
		"--app", app.ID,
		"--cache-root", t.TempDir(),
		"--guest-boundary", winapp.KnownDispatchGuestBoundary,
		"--state-root", stateRoot,
		"--receipt-id", receipt.ReceiptID,
		"--review-receipt-id", reviewReceipt.ReceiptID,
		"--session-id", sessionID,
		"--registry", packagedRecipeRegistryDir + "/registry.json",
		"--image", "local/wine-x-gui:test",
		"--platform", "linux/amd64",
		"--docker", dockerPath,
		"--timeout", "5s",
	}, &output)
	if err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["status"] != "passed" ||
		payload["application_id"] != app.ID ||
		payload["app_version"] != app.Version ||
		payload["recipe_backed"] != true ||
		payload["smoke_passed"] != true ||
		payload["controlled_execution_session_id"] != sessionID ||
		payload["network_mode"] != "none" ||
		payload["host_mount_count"] != float64(0) ||
		payload["host_root_modified"] != false {
		t.Fatalf("unexpected staged packaged Notepad launcher payload: %#v", payload)
	}
	dockerInvocation, err := os.ReadFile(dockerLog)
	if err != nil {
		t.Fatalf("ReadFile docker log returned error: %v", err)
	}
	if !strings.Contains(string(dockerInvocation), "XNIX_GUI_APP=notepad.exe") ||
		!strings.Contains(string(dockerInvocation), "XNIX_WINDOW_MATCH=notepad.exe") {
		t.Fatalf("docker invocation did not receive staged recipe GUI hints: %s", string(dockerInvocation))
	}
	assertCompatLaunchContainerGUIDispatchSafe(t, output.String(), stateRoot, stagingRoot, stagedRegistryPath, stagedRecipePath, dockerPath)
}

func writeNotepadRecipeRegistryFixture(t *testing.T, registryDir string, version string) (string, string) {
	t.Helper()
	recipeData := []byte(`{
  "id": "org.xnix.sample.notepad",
  "name": "Sample Notepad",
  "version": "` + version + `",
  "icon": "accessories-text-editor",
  "mode": "automatic",
  "supported_extensions": [".txt"],
  "container_gui_smoke": {
    "app": "notepad.exe",
    "window_match": "notepad.exe"
  }
}`)
	recipePath := filepath.Join(registryDir, "org.xnix.sample.notepad.json")
	if err := os.WriteFile(recipePath, recipeData, 0o600); err != nil {
		t.Fatalf("WriteFile recipe returned error: %v", err)
	}
	digest := sha256.Sum256(recipeData)
	registryPath := filepath.Join(registryDir, "registry.json")
	registryData := []byte(fmt.Sprintf(`{"schema_version":1,"registry_name":"test-registry","recipes":[{"id":"org.xnix.sample.notepad","path":"org.xnix.sample.notepad.json","sha256":"%x","signature_status":"development-only"}]}`, digest[:]))
	if err := os.WriteFile(registryPath, registryData, 0o600); err != nil {
		t.Fatalf("WriteFile registry returned error: %v", err)
	}
	return registryPath, recipePath
}

func writeRecipeBackedFakeDocker(t *testing.T, tempDir string) (string, string) {
	t.Helper()
	dockerLog := filepath.Join(tempDir, "fake-docker.log")
	dockerPath := filepath.Join(tempDir, "fake-docker")
	dockerBody := "#!/bin/sh\n" +
		"printf '%s\\n' \"$@\" >> \"" + dockerLog + "\"\n" +
		"if test \"$1 $2\" = 'image inspect'; then exit 0; fi\n" +
		"printf 'XNIX_X_GUI_XSERVER_STARTED=true\\n'\n" +
		"printf 'XNIX_X_GUI_WINE_BOOTSTRAP_ATTEMPTED=true\\n'\n" +
		"printf '0x600001 \"Untitled - Notepad\": (\"notepad.exe\" \"notepad.exe\") 721x519+4+23 +4+23\\n'\n" +
		"printf 'XNIX_X_GUI_WINDOW_OBSERVED=true\\n'\n"
	if err := os.WriteFile(dockerPath, []byte(dockerBody), 0o700); err != nil {
		t.Fatalf("WriteFile docker returned error: %v", err)
	}
	return dockerPath, dockerLog
}

func writeExternalImportedFakeDocker(t *testing.T, tempDir string) (string, string) {
	t.Helper()
	dockerLog := filepath.Join(tempDir, "fake-external-docker.log")
	dockerPath := filepath.Join(tempDir, "fake-external-docker")
	dockerBody := "#!/bin/sh\n" +
		"printf '%s\\n' \"$@\" >> \"" + dockerLog + "\"\n" +
		"if test \"$1 $2\" = 'image inspect'; then printf 'linux/amd64\\n'; exit 0; fi\n" +
		"if test \"$1\" = 'create'; then printf 'fake-external-x-gui-container\\n'; exit 0; fi\n" +
		"if test \"$1\" = 'cp'; then exit 0; fi\n" +
		"if test \"$1 $2\" = 'start -a'; then " +
		"printf 'XNIX_X_GUI_XSERVER_STARTED=true\\n'\n" +
		"printf 'XNIX_X_GUI_WINE_BOOTSTRAP_ATTEMPTED=true\\n'\n" +
		"printf '0x700001 \"External GUI\": (\"ExternalGui.exe\" \"ExternalGui.exe\") 320x160+20+20 +20+20\\n'\n" +
		"printf 'XNIX_X_GUI_WINDOW_OBSERVED=true\\n'; exit 0; fi\n" +
		"if test \"$1\" = 'rm'; then exit 0; fi\n" +
		"exit 2\n"
	if err := os.WriteFile(dockerPath, []byte(dockerBody), 0o700); err != nil {
		t.Fatalf("WriteFile docker returned error: %v", err)
	}
	return dockerPath, dockerLog
}

func recordLauncherSessionGateFixture(t *testing.T, stateRoot string) (string, string) {
	t.Helper()
	return recordLauncherSessionGateFixtureForApp(t, stateRoot, "7zr", "26.02")
}

func recordLauncherSessionGateFixtureForApp(t *testing.T, stateRoot string, appID string, appVersion string) (string, string) {
	t.Helper()
	sessionID := appidentity.KnownAppControlledExecutionSessionID(appID, appVersion)
	ledger, err := execution.NewLedger(stateRoot)
	if err != nil {
		t.Fatalf("NewLedger returned error: %v", err)
	}
	if _, err := ledger.Record(execution.Transaction{
		RequestID:      sessionID,
		ApplicationID:  appID,
		Profile:        "known-app-managed-guest",
		State:          execution.StateBlocked,
		ReviewDecision: execution.DecisionApproved,
		Gates: []execution.Gate{
			{ID: "controlled-execution-session-handoff", Status: execution.GatePass, Reason: "Runtime execution session handoff ready"},
			{ID: "runtime-write-gate", Status: execution.GateBlocked, Reason: "dispatch runner must consume the recorded session before execution"},
		},
		BlockedReasons:   []string{"runtime-write-gate: dispatch runner must consume the recorded session before execution"},
		LaunchAllowed:    false,
		LaunchEnabled:    false,
		BackendStarted:   false,
		HostRootModified: false,
		NetworkRequired:  false,
		Summary:          "Runtime recorded a known application execution session handoff for launcher consumption.",
	}); err != nil {
		t.Fatalf("Record returned error: %v", err)
	}
	session, err := ledger.RecordSession(sessionID)
	if err != nil {
		t.Fatalf("RecordSession returned error: %v", err)
	}
	return sessionID, session.RelativePath
}

func writeGuestGUIFakeTools(t *testing.T) (string, string) {
	t.Helper()
	tempDir := t.TempDir()
	sshPath := filepath.Join(tempDir, "fake-ssh")
	sshBody := "#!/bin/sh\n" +
		"case \"$*\" in\n" +
		"  *' true') exit 0 ;;\n" +
		"  *'command -v wine'*) exit 0 ;;\n" +
		"  *'winex11'*) exit 0 ;;\n" +
		"  *'mkdir -p'*) exit 0 ;;\n" +
		"  *'wineboot --init'*) printf 'boot initialized\\n' >&2; exit 0 ;;\n" +
		"  *'wine '*'winemine.exe'*) exit 0 ;;\n" +
		"  *'wine '*'owner-messagebox.exe'*) exit 0 ;;\n" +
		"  *'cat '*'stderr.txt'*) printf ''; exit 0 ;;\n" +
		"  *'wineserver -k'*) exit 0 ;;\n" +
		"esac\n" +
		"exit 2\n"
	if err := os.WriteFile(sshPath, []byte(sshBody), 0o700); err != nil {
		t.Fatalf("WriteFile ssh returned error: %v", err)
	}
	xwininfoPath := filepath.Join(tempDir, "fake-xwininfo")
	xwininfoBody := "#!/bin/sh\n" +
		"printf 'xwininfo: Window id: 0x3a7 (the root window)\\n'\n" +
		"printf '  0x200001 \"WineMine\": ()  320x240+0+0  +0+0\\n'\n"
	if err := os.WriteFile(xwininfoPath, []byte(xwininfoBody), 0o700); err != nil {
		t.Fatalf("WriteFile xwininfo returned error: %v", err)
	}
	return sshPath, xwininfoPath
}

func writeGuestGUIFakeSCPTool(t *testing.T) (string, string) {
	t.Helper()
	tempDir := t.TempDir()
	scpPath := filepath.Join(tempDir, "fake-scp")
	logPath := filepath.Join(tempDir, "scp-args.txt")
	body := "#!/bin/sh\n" +
		"printf '%s\\n' \"$@\" > \"" + logPath + "\"\n" +
		"exit 0\n"
	if err := os.WriteFile(scpPath, []byte(body), 0o700); err != nil {
		t.Fatalf("WriteFile scp returned error: %v", err)
	}
	return scpPath, logPath
}

func TestCompatLaunchSessionGateRejectsMissingSessionWithoutPathLeak(t *testing.T) {
	stateRoot := t.TempDir()

	_, err := consumeControlledExecutionSessionForLaunch("7zr", stateRoot, "")
	if err == nil {
		t.Fatal("expected missing controlled execution session record to be rejected")
	}
	if strings.Contains(strings.ToLower(err.Error()), strings.ToLower(stateRoot)) {
		t.Fatalf("session gate error exposed state root path: %v", err)
	}
	if !strings.Contains(err.Error(), "controlled execution session gate rejected dispatch") ||
		!strings.Contains(err.Error(), "digest-verified ledger record") {
		t.Fatalf("unexpected session gate error: %v", err)
	}
}

func TestCompatLaunchRejectsUnknownApp(t *testing.T) {
	var output bytes.Buffer
	err := run([]string{"--app", "missing-app"}, &output)
	if err == nil || !strings.Contains(err.Error(), "unknown known Windows app") {
		t.Fatalf("expected unknown app rejection, got %v", err)
	}
}

func assertCompatLaunchCLISafe(t *testing.T, text string, hostPath string) {
	t.Helper()
	serialized := strings.ToLower(text)
	for _, forbidden := range []string{".exe", "wine", "qemu", strings.ToLower(hostPath)} {
		if strings.Contains(serialized, forbidden) {
			t.Fatalf("compat launch CLI exposed forbidden term %q: %s", forbidden, text)
		}
	}
}

func assertCompatLaunchGUIDispatchSafe(t *testing.T, text string, hostPaths ...string) {
	t.Helper()
	serialized := strings.ToLower(text)
	for _, forbidden := range []string{".exe", "qemu-system", "program files", "/usr/lib/wine", "wine ", ".wine"} {
		if strings.Contains(serialized, forbidden) {
			t.Fatalf("compat launch GUI dispatch exposed forbidden term %q: %s", forbidden, text)
		}
	}
	for _, hostPath := range hostPaths {
		if strings.Contains(serialized, strings.ToLower(hostPath)) {
			t.Fatalf("compat launch GUI dispatch exposed host path %q: %s", hostPath, text)
		}
	}
}

func assertCompatLaunchContainerGUIDispatchSafe(t *testing.T, text string, hostPaths ...string) {
	t.Helper()
	serialized := strings.ToLower(text)
	for _, forbidden := range []string{"docker.sock", "--privileged", "--network host", "qemu-system", "program files", "/usr/lib/wine", ".wine"} {
		if strings.Contains(serialized, forbidden) {
			t.Fatalf("compat launch container GUI dispatch exposed forbidden term %q: %s", forbidden, text)
		}
	}
	for _, hostPath := range hostPaths {
		if strings.Contains(serialized, strings.ToLower(hostPath)) {
			t.Fatalf("compat launch container GUI dispatch exposed host path %q: %s", hostPath, text)
		}
	}
}
