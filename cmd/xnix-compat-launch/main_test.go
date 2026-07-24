package main

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

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
