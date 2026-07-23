package main

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"

	"xnix.local/xnix/internal/runtime/appidentity"
	"xnix.local/xnix/internal/runtime/execution"
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
	if err == nil || !strings.Contains(err.Error(), "--state-root and --receipt-id are required") {
		t.Fatalf("expected missing receipt rejection, got %v", err)
	}
}

func TestCompatLaunchConsumesControlledDispatchRequestBeforeDispatch(t *testing.T) {
	cacheRoot := t.TempDir()
	stateRoot := t.TempDir()
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
	}, &output)
	if err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["schema_version"] != appidentity.KnownAppControlledDispatchSchemaVersion ||
		payload["request_type"] != appidentity.KnownAppControlledDispatchRequestType ||
		payload["receipt_accepted"] != true ||
		payload["guest_boundary_accepted"] != true ||
		payload["launch_gate_state"] != "dispatch-preparation-required" ||
		payload["controlled_dispatch_request_created"] != false ||
		payload["dispatch_started"] != false ||
		payload["execution_started"] != false ||
		payload["direct_launch_enabled"] != false ||
		payload["desktop_launch_enabled"] != false ||
		payload["backend_launch_enabled"] != false ||
		payload["backend_process_started"] != false ||
		payload["host_root_modified"] != false ||
		payload["receipt_path_exposed"] != false ||
		payload["state_root_path_exposed"] != false {
		t.Fatalf("unexpected controlled dispatch request payload: %#v", payload)
	}
	assertCompatLaunchCLISafe(t, output.String(), cacheRoot)
	assertCompatLaunchCLISafe(t, output.String(), stateRoot)
}

func TestCompatLaunchConsumesDigestVerifiedSessionGate(t *testing.T) {
	stateRoot := t.TempDir()
	sessionID := appidentity.KnownAppControlledExecutionSessionID("7zr", "26.02")
	ledger, err := execution.NewLedger(stateRoot)
	if err != nil {
		t.Fatalf("NewLedger returned error: %v", err)
	}
	if _, err := ledger.Record(execution.Transaction{
		RequestID:      sessionID,
		ApplicationID:  "7zr",
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

	preview, err := consumeControlledExecutionSessionForLaunch("7zr", stateRoot, "")
	if err != nil {
		t.Fatalf("consumeControlledExecutionSessionForLaunch returned error: %v", err)
	}
	if preview.ExecutionSessionID != sessionID ||
		!preview.RecordConsumed ||
		!preview.LedgerRecordConsumed ||
		!preview.SessionRecordConsumed ||
		!preview.SessionDigestVerified ||
		preview.SessionRelativePath != session.RelativePath ||
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
