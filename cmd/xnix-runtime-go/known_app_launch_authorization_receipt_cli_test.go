package main

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"xnix.local/xnix/internal/runtime/appidentity"
	"xnix.local/xnix/internal/runtime/execution"
)

func TestKnownAppLaunchAuthorizationReceiptPreviewCommandWritesReceipt(t *testing.T) {
	stateRoot := t.TempDir()
	var output bytes.Buffer
	err := run([]string{
		"known-app-launch-authorization-receipt-preview",
		"--app", "7zr",
		"--state-root", stateRoot,
		"--authorize", appidentity.KnownAppLaunchAuthorizationReceiptAction,
	}, &output)
	if err != nil {
		t.Fatalf("run returned error: %v", err)
	}
	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["schema_version"] != appidentity.KnownAppLaunchAuthorizationReceiptSchemaVersion ||
		payload["request_type"] != appidentity.KnownAppLaunchAuthorizationReceiptRequestType ||
		payload["app_id"] != "7zr" ||
		payload["receipt_id"] != "known-app-launch-authorization-7zr-26.02" ||
		payload["receipt_state"] != "recorded" ||
		payload["receipt_written"] != true ||
		payload["receipt_path_exposed"] != false ||
		payload["state_root_path_exposed"] != false ||
		payload["launch_gate_ready"] != true ||
		payload["direct_launch_enabled"] != false ||
		payload["desktop_launch_enabled"] != false ||
		payload["backend_launch_enabled"] != false ||
		payload["backend_process_started"] != false {
		t.Fatalf("unexpected receipt preview payload: %#v", payload)
	}
	text := strings.ToLower(output.String())
	if strings.Contains(text, strings.ToLower(stateRoot)) {
		t.Fatalf("receipt preview exposed state root path: %s", text)
	}
}

func TestKnownAppKDERuntimeStatusLaunchRequestPreviewCommandCollectsOpaqueIDs(t *testing.T) {
	sessionID := appidentity.KnownAppControlledExecutionSessionID("7zr", "26.02")
	launchReceiptID := appidentity.KnownAppLaunchAuthorizationReceiptID("7zr", "26.02")
	reviewReceiptID := appidentity.KnownAppSessionGatedLaunchReviewReceiptID("7zr", "26.02", sessionID)
	var output bytes.Buffer
	err := run([]string{
		"known-app-kde-runtime-status-launch-request-preview",
		"--app", "7zr",
		"--launch-authorization-receipt-id", launchReceiptID,
		"--session-gated-review-receipt-id", reviewReceiptID,
		"--session-id", sessionID,
		"--center-card-state", "validated-post-review-dispatch",
		"--primary-action-id", appidentity.KnownAppKDERuntimeStatusLaunchAction,
		"--post-review-dispatch-state", "created-after-session-gated-review",
	}, &output)
	if err != nil {
		t.Fatalf("run returned error: %v", err)
	}
	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["schema_version"] != appidentity.KnownAppKDERuntimeStatusLaunchRequestSchemaVersion ||
		payload["request_type"] != appidentity.KnownAppKDERuntimeStatusLaunchRequestType ||
		payload["runtime_method"] != "PreviewKnownAppKDERuntimeStatusLaunchRequest" ||
		payload["read_method"] != "GetKnownAppKDERuntimeStatusLaunchRequest" ||
		payload["action_id"] != appidentity.KnownAppKDERuntimeStatusLaunchAction ||
		payload["action_kind"] != "runtime-status" ||
		payload["center_card_state"] != "validated-post-review-dispatch" ||
		payload["launch_authorization_receipt_id"] != launchReceiptID ||
		payload["session_gated_review_receipt_id"] != reviewReceiptID ||
		payload["controlled_execution_session_id"] != sessionID ||
		payload["managed_launcher_argv_ready"] != true ||
		payload["required_opaque_id_count"] != float64(3) ||
		payload["collected_opaque_id_count"] != float64(3) ||
		payload["launch_request_created"] != true ||
		payload["runtime_owned_request"] != true ||
		payload["runtime_owned_launch"] != true ||
		payload["runtime_owned_dispatch"] != true ||
		payload["kde_presentation_only"] != true ||
		payload["kde_action_forwarded"] != true ||
		payload["state_root_required"] != true ||
		payload["state_root_supplied_by_runtime"] != true ||
		payload["kde_state_root_access"] != false ||
		payload["direct_launch_enabled"] != false ||
		payload["desktop_launch_enabled"] != false ||
		payload["backend_launch_enabled"] != false ||
		payload["execution_started"] != false ||
		payload["backend_process_started"] != false ||
		payload["request_objects_created"] != false ||
		payload["permission_grant_created"] != false ||
		payload["state_root_path_exposed"] != false ||
		payload["receipt_path_exposed"] != false ||
		payload["session_path_exposed"] != false ||
		payload["raw_artifact_path_exposed"] != false ||
		payload["raw_command_exposed"] != false ||
		payload["backend_details_exposed"] != false ||
		payload["host_root_modified"] != false {
		t.Fatalf("unexpected KDE Runtime-status launch request payload: %#v", payload)
	}
	argv := payload["managed_launcher_argv"].([]any)
	if len(argv) != 11 ||
		argv[0] != "xnix-compat-launch" ||
		argv[1] != "--app" ||
		argv[2] != "7zr" ||
		argv[3] != "--guest-boundary" ||
		argv[4] != "managed-known-app-guest-smoke" ||
		argv[5] != "--receipt-id" ||
		argv[6] != launchReceiptID ||
		argv[7] != "--review-receipt-id" ||
		argv[8] != reviewReceiptID ||
		argv[9] != "--session-id" ||
		argv[10] != sessionID {
		t.Fatalf("unexpected managed launcher argv: %#v", argv)
	}
	assertWindowIdentityPayloadSafe(t, output.String())
}

func TestKnownAppKDERuntimeStatusLaunchExecutionCommandInvokesManagedLauncherWithRuntimeStateRoot(t *testing.T) {
	stateRoot := t.TempDir()
	sessionID := writeKnownAppRuntimeStatusLaunchExecutionCLIFixture(t, stateRoot)
	launchReceiptID := appidentity.KnownAppLaunchAuthorizationReceiptID("7zr", "26.02")
	reviewReceiptID := appidentity.KnownAppSessionGatedLaunchReviewReceiptID("7zr", "26.02", sessionID)
	fakeLauncher, fakeArgsPath := writeFakeRuntimeStatusManagedLauncher(t)
	var output bytes.Buffer
	err := run([]string{
		appidentity.KnownAppKDERuntimeStatusLaunchExecutionRequestType,
		"--app", "7zr",
		"--state-root", stateRoot,
		"--launch-authorization-receipt-id", launchReceiptID,
		"--session-gated-review-receipt-id", reviewReceiptID,
		"--session-id", sessionID,
		"--center-card-state", "validated-post-review-dispatch",
		"--primary-action-id", appidentity.KnownAppKDERuntimeStatusLaunchAction,
		"--post-review-dispatch-state", "created-after-session-gated-review",
		"--launcher", fakeLauncher,
		"--owner-timeout", "5s",
		"--timeout", "1s",
		"--arg", "--help",
	}, &output)
	if err != nil {
		t.Fatalf("run returned error: %v", err)
	}
	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["schema_version"] != appidentity.KnownAppKDERuntimeStatusLaunchExecutionSchemaVersion ||
		payload["request_type"] != appidentity.KnownAppKDERuntimeStatusLaunchExecutionRequestType ||
		payload["runtime_method"] != "PrepareKnownAppKDERuntimeStatusLaunchExecution" ||
		payload["execution_method"] != "RunKnownAppKDERuntimeStatusLaunchExecution" ||
		payload["request_preview_type"] != appidentity.KnownAppKDERuntimeStatusLaunchRequestType ||
		payload["launch_authorization_receipt_id"] != launchReceiptID ||
		payload["session_gated_review_receipt_id"] != reviewReceiptID ||
		payload["controlled_execution_session_id"] != sessionID ||
		payload["launch_receipt_revalidated"] != true ||
		payload["guest_boundary_revalidated"] != true ||
		payload["review_receipt_revalidated"] != true ||
		payload["controlled_session_revalidated"] != true ||
		payload["controlled_session_digest_verified"] != true ||
		payload["runtime_managed_launcher_argv_ready"] != true ||
		payload["state_root_required"] != true ||
		payload["state_root_accepted"] != true ||
		payload["state_root_injected_by_runtime"] != true ||
		payload["state_root_supplied_by_runtime"] != true ||
		payload["kde_state_root_access"] != false ||
		payload["managed_launcher_invocation_ready"] != true ||
		payload["existing_managed_launcher_path_used"] != true ||
		payload["managed_launcher_invoked"] != true ||
		payload["existing_managed_launcher_invoked"] != true ||
		payload["launcher_output_json_observed"] != true ||
		payload["delegated_request_type"] != "windows-known-app-dispatch-smoke" ||
		payload["delegated_status"] != "passed" ||
		payload["delegated_guest_boundary"] != "managed-known-app-guest-smoke" ||
		payload["delegated_runtime_owned_dispatch"] != true ||
		payload["delegated_artifact_verified"] != true ||
		payload["delegated_marker_observed"] != true ||
		payload["delegated_smoke_passed"] != true ||
		payload["delegated_execution_started"] != true ||
		payload["delegated_backend_process_started"] != false ||
		payload["delegated_session_gated_controlled_dispatch_consumed"] != true ||
		payload["delegated_session_gated_controlled_dispatch_state"] != "created-after-session-gated-review" ||
		payload["delegated_session_gated_review_receipt_id"] != reviewReceiptID ||
		payload["delegated_launch_authorization_receipt_id"] != launchReceiptID ||
		payload["delegated_controlled_execution_session_consumed"] != true ||
		payload["delegated_controlled_execution_session_id"] != sessionID ||
		payload["delegated_controlled_session_digest_verified"] != true ||
		payload["delegated_controlled_session_relative_path"] != "execution-ledger/sessions/"+sessionID+".json" ||
		payload["delegated_runtime_owner_consumable_session"] != true ||
		payload["delegated_kde_read_model_consumable_session"] != true ||
		payload["delegated_controlled_session_live_state_observed"] != false ||
		payload["delegated_controlled_session_registered"] != false ||
		payload["delegated_controlled_session_window_observed"] != false ||
		payload["delegated_controlled_session_host_root_modified"] != false ||
		payload["delegated_controlled_session_backend_process_start"] != false ||
		payload["delegated_host_root_modified"] != false ||
		payload["delegated_docker_socket_mounted"] != false ||
		payload["delegated_broad_host_mount_required"] != false ||
		payload["delegated_raw_command_exposed"] != false ||
		payload["delegated_backend_details_exposed"] != false ||
		payload["state_root_path_exposed"] != false ||
		payload["managed_launcher_path_exposed"] != false ||
		payload["raw_launcher_output_exposed"] != false ||
		payload["host_root_modified"] != false ||
		payload["docker_socket_mounted"] != false ||
		payload["broad_host_mount_required"] != false {
		t.Fatalf("unexpected Runtime-status launch execution payload: %#v", payload)
	}
	runtimeArgv := payload["runtime_managed_launcher_argv"].([]any)
	if len(runtimeArgv) != 13 ||
		runtimeArgv[0] != "xnix-compat-launch" ||
		runtimeArgv[5] != "--state-root" ||
		runtimeArgv[6] != "<runtime-owned-state-root>" {
		t.Fatalf("unexpected redacted Runtime managed launcher argv: %#v", runtimeArgv)
	}
	projection := payload["compatibility_center_known_app_evidence"].(map[string]any)
	if projection["projection_type"] != "known-app-kde-runtime-status-launch-delegated-evidence" ||
		projection["runtime_method"] != "ProjectKnownAppKDERuntimeStatusLaunchDelegatedEvidence" ||
		projection["request_type"] != "windows-known-app-dispatch-smoke" ||
		projection["status"] != "passed" ||
		projection["app_id"] != "7zr" ||
		projection["guest_boundary"] != "managed-known-app-guest-smoke" ||
		projection["runtime_owned_dispatch"] != true ||
		projection["artifact_verified"] != true ||
		projection["marker_observed"] != true ||
		projection["session_gated_controlled_dispatch_consumed"] != true ||
		projection["session_gated_controlled_dispatch_state"] != "created-after-session-gated-review" ||
		projection["session_gated_review_receipt_id"] != reviewReceiptID ||
		projection["launch_authorization_receipt_id"] != launchReceiptID ||
		projection["controlled_execution_session_consumed"] != true ||
		projection["controlled_execution_session_id"] != sessionID ||
		projection["controlled_session_digest_verified"] != true ||
		projection["controlled_session_relative_path"] != "execution-ledger/sessions/"+sessionID+".json" ||
		projection["runtime_owner_consumable_session"] != true ||
		projection["kde_read_model_consumable_session"] != true ||
		projection["state_root_path_exposed"] != false ||
		projection["managed_launcher_path_exposed"] != false ||
		projection["raw_launcher_output_exposed"] != false ||
		projection["backend_details_exposed"] != false ||
		projection["compatibility_center_projection_ready"] != true ||
		projection["kde_center_projection_ready"] != true {
		t.Fatalf("unexpected Compatibility Center known app evidence projection: %#v", projection)
	}
	argsData, err := os.ReadFile(fakeArgsPath)
	if err != nil {
		t.Fatalf("ReadFile fake launcher args returned error: %v", err)
	}
	argsText := string(argsData)
	for _, token := range []string{"--state-root\n" + stateRoot, "--receipt-id\n" + launchReceiptID, "--review-receipt-id\n" + reviewReceiptID, "--session-id\n" + sessionID, "--arg\n--help"} {
		if !strings.Contains(argsText, token) {
			t.Fatalf("fake launcher args missing %q: %s", token, argsText)
		}
	}
	text := strings.ToLower(output.String())
	for _, forbidden := range []string{strings.ToLower(stateRoot), strings.ToLower(fakeLauncher), ".exe", "program files", "qemu-system", "proton", "wine ", "/tmp"} {
		if strings.Contains(text, forbidden) {
			t.Fatalf("Runtime-status launch execution output exposes forbidden term %q: %s", forbidden, text)
		}
	}
}

func TestKnownAppSessionGatedLaunchReviewReceiptRecordCommandWritesReceipt(t *testing.T) {
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
		Summary:          "Runtime recorded a known application execution session handoff for later consumption.",
	}); err != nil {
		t.Fatalf("Record returned error: %v", err)
	}
	if _, err := ledger.RecordSession(sessionID); err != nil {
		t.Fatalf("RecordSession returned error: %v", err)
	}

	var output bytes.Buffer
	err = run([]string{
		"known-app-session-gated-launch-review-receipt-record",
		"--app", "7zr",
		"--state-root", stateRoot,
		"--session-id", sessionID,
		"--action", appidentity.KnownAppSessionGatedLaunchReviewAction,
		"--decision", "approved",
	}, &output)
	if err != nil {
		t.Fatalf("session-gated launch review receipt run returned error: %v", err)
	}
	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal session-gated launch review receipt returned error: %v", err)
	}
	receiptID := appidentity.KnownAppSessionGatedLaunchReviewReceiptID("7zr", "26.02", sessionID)
	receiptRelativePath := appidentity.KnownAppSessionGatedLaunchReviewReceiptRelativePath(receiptID)
	if payload["schema_version"] != appidentity.KnownAppSessionGatedLaunchReviewReceiptSchemaVersion ||
		payload["request_type"] != appidentity.KnownAppSessionGatedLaunchReviewReceiptRequestType ||
		payload["source"] != appidentity.KnownAppSessionGatedLaunchReviewRequestType+"+runtime-review-receipt-store" ||
		payload["runtime_method"] != "RecordKnownAppSessionGatedLaunchReviewReceipt" ||
		payload["read_method"] != "GetKnownAppSessionGatedLaunchReviewReceipt" ||
		payload["app_id"] != "7zr" ||
		payload["action_id"] != "review-session-gated-dispatch" ||
		payload["action_kind"] != "session-gate-review" ||
		payload["decision"] != "approved" ||
		payload["read_before_write_consumed"] != true ||
		payload["review_route_consumed"] != true ||
		payload["runtime_receipt_required"] != true ||
		payload["execution_session_id"] != sessionID ||
		payload["session_record_consumed"] != true ||
		payload["session_digest_verified"] != true ||
		payload["session_relative_path"] != "execution-ledger/sessions/"+sessionID+".json" ||
		payload["receipt_id"] != receiptID ||
		payload["receipt_relative_path"] != receiptRelativePath ||
		payload["receipt_state"] != "recorded-dispatch-still-gated" ||
		payload["review_receipt_recorded"] != true ||
		payload["runtime_owner_consumable"] != true ||
		payload["kde_read_model_consumable"] != true ||
		payload["runtime_owned"] != true ||
		payload["go_runtime_backed"] != true ||
		payload["kde_policy_owner"] != false ||
		payload["user_decision_allows_launch"] != false ||
		payload["runtime_launch_approval"] != false ||
		payload["launch_allowed"] != false ||
		payload["launch_enabled"] != false ||
		payload["desktop_launch_enabled"] != false ||
		payload["execution_started"] != false ||
		payload["backend_launch_enabled"] != false ||
		payload["backend_process_started"] != false ||
		payload["request_objects_created"] != false ||
		payload["permission_grant_created"] != false ||
		payload["state_root_path_exposed"] != false ||
		payload["session_path_exposed"] != false ||
		payload["receipt_path_exposed"] != false ||
		payload["raw_artifact_path_exposed"] != false ||
		payload["backend_details_exposed"] != false ||
		payload["host_root_modified"] != false {
		t.Fatalf("unexpected session-gated launch review receipt payload: %#v", payload)
	}
	if digest, ok := payload["receipt_sha256"].(string); !ok || len(digest) != 64 {
		t.Fatalf("session-gated launch review receipt must expose a digest, got %#v", payload["receipt_sha256"])
	}
	if _, err := os.Stat(filepath.Join(stateRoot, filepath.FromSlash(receiptRelativePath))); err != nil {
		t.Fatalf("expected persisted session-gated launch review receipt: %v", err)
	}
	text := strings.ToLower(output.String())
	if strings.Contains(text, strings.ToLower(stateRoot)) {
		t.Fatalf("session-gated launch review receipt exposed state root path: %s", text)
	}
}

func TestKnownAppSessionGatedLaunchReviewGatePreviewCommandAcceptsReceipt(t *testing.T) {
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
		Summary:          "Runtime recorded a known application execution session handoff for later consumption.",
	}); err != nil {
		t.Fatalf("Record returned error: %v", err)
	}
	if _, err := ledger.RecordSession(sessionID); err != nil {
		t.Fatalf("RecordSession returned error: %v", err)
	}
	receipt, err := appidentity.RecordKnownAppSessionGatedLaunchReviewReceipt(appidentity.KnownAppSessionGatedLaunchReviewReceiptRequest{
		AppID:     "7zr",
		StateRoot: stateRoot,
		SessionID: sessionID,
		ActionID:  appidentity.KnownAppSessionGatedLaunchReviewAction,
		Decision:  "approved",
	})
	if err != nil {
		t.Fatalf("RecordKnownAppSessionGatedLaunchReviewReceipt returned error: %v", err)
	}

	var output bytes.Buffer
	err = run([]string{
		"known-app-session-gated-launch-review-gate-preview",
		"--app", "7zr",
		"--state-root", stateRoot,
		"--session-id", sessionID,
		"--receipt-id", receipt.ReceiptID,
	}, &output)
	if err != nil {
		t.Fatalf("session-gated launch review gate run returned error: %v", err)
	}
	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal session-gated launch review gate returned error: %v", err)
	}
	if payload["schema_version"] != appidentity.KnownAppSessionGatedLaunchReviewGateSchemaVersion ||
		payload["request_type"] != appidentity.KnownAppSessionGatedLaunchReviewGateRequestType ||
		payload["source"] != appidentity.KnownAppSessionGatedLaunchReviewReceiptRequestType+"+runtime-review-receipt-gate" ||
		payload["runtime_method"] != "PreviewKnownAppSessionGatedLaunchReviewGate" ||
		payload["read_method"] != "GetKnownAppSessionGatedLaunchReviewGate" ||
		payload["app_id"] != "7zr" ||
		payload["action_id"] != "review-session-gated-dispatch" ||
		payload["action_kind"] != "session-gate-review" ||
		payload["execution_session_id"] != sessionID ||
		payload["session_record_consumed"] != true ||
		payload["session_digest_verified"] != true ||
		payload["session_relative_path"] != "execution-ledger/sessions/"+sessionID+".json" ||
		payload["read_before_write_revalidated"] != true ||
		payload["receipt_id"] != receipt.ReceiptID ||
		payload["receipt_relative_path"] != receipt.ReceiptRelativePath ||
		payload["receipt_sha256"] != receipt.ReceiptSHA256 ||
		payload["receipt_lookup_state"] != "accepted-receipt" ||
		payload["receipt_decision"] != "approved" ||
		payload["review_receipt_consumed"] != true ||
		payload["review_receipt_accepted"] != true ||
		payload["review_gate_state"] != "review-receipt-accepted-dispatch-still-gated" ||
		payload["review_gate_ready"] != true ||
		payload["controlled_dispatch_gate_ready"] != true ||
		payload["dispatch_state_advance_ready"] != true ||
		payload["dispatch_state_advanced"] != false ||
		payload["runtime_owner_consumable"] != true ||
		payload["kde_read_model_consumable"] != true ||
		payload["runtime_owned"] != true ||
		payload["go_runtime_backed"] != true ||
		payload["kde_policy_owner"] != false ||
		payload["user_decision_allows_launch"] != true ||
		payload["runtime_launch_approval"] != false ||
		payload["launch_allowed"] != false ||
		payload["launch_enabled"] != false ||
		payload["desktop_launch_enabled"] != false ||
		payload["execution_started"] != false ||
		payload["backend_launch_enabled"] != false ||
		payload["backend_process_started"] != false ||
		payload["controlled_dispatch_request_created"] != false ||
		payload["permission_grant_created"] != false ||
		payload["state_root_path_exposed"] != false ||
		payload["session_path_exposed"] != false ||
		payload["receipt_path_exposed"] != false ||
		payload["raw_artifact_path_exposed"] != false ||
		payload["backend_details_exposed"] != false ||
		payload["host_root_modified"] != false {
		t.Fatalf("unexpected session-gated launch review gate payload: %#v", payload)
	}
	text := strings.ToLower(output.String())
	if strings.Contains(text, strings.ToLower(stateRoot)) {
		t.Fatalf("session-gated launch review gate exposed state root path: %s", text)
	}
}

func TestKnownAppSessionGatedControlledDispatchRequestPreviewCommandRequiresReviewGate(t *testing.T) {
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
		Summary:          "Runtime recorded a known application execution session handoff for later consumption.",
	}); err != nil {
		t.Fatalf("Record returned error: %v", err)
	}
	if _, err := ledger.RecordSession(sessionID); err != nil {
		t.Fatalf("RecordSession returned error: %v", err)
	}
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
	launchReceipt, err := appidentity.RecordKnownAppLaunchAuthorizationReceipt(appidentity.KnownAppLaunchAuthorizationReceiptRequest{
		AppID:     "7zr",
		StateRoot: stateRoot,
		Authorize: appidentity.KnownAppLaunchAuthorizationReceiptAction,
	})
	if err != nil {
		t.Fatalf("RecordKnownAppLaunchAuthorizationReceipt returned error: %v", err)
	}

	var output bytes.Buffer
	err = run([]string{
		"known-app-session-gated-controlled-dispatch-request-preview",
		"--app", "7zr",
		"--state-root", stateRoot,
		"--session-id", sessionID,
		"--review-receipt-id", reviewReceipt.ReceiptID,
		"--launch-receipt-id", launchReceipt.ReceiptID,
		"--cache-root", t.TempDir(),
		"--guest-boundary", "managed-known-app-guest-smoke",
	}, &output)
	if err != nil {
		t.Fatalf("session-gated controlled dispatch request run returned error: %v", err)
	}
	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal session-gated controlled dispatch request returned error: %v", err)
	}
	if payload["schema_version"] != appidentity.KnownAppSessionGatedControlledDispatchSchemaVersion ||
		payload["request_type"] != appidentity.KnownAppSessionGatedControlledDispatchRequestType ||
		payload["source"] != appidentity.KnownAppSessionGatedLaunchReviewGateRequestType+"+"+appidentity.KnownAppControlledDispatchRequestType ||
		payload["runtime_method"] != "PreviewKnownAppSessionGatedControlledDispatchRequest" ||
		payload["read_method"] != "GetKnownAppSessionGatedControlledDispatchRequest" ||
		payload["app_id"] != "7zr" ||
		payload["execution_session_id"] != sessionID ||
		payload["session_record_consumed"] != true ||
		payload["session_digest_verified"] != true ||
		payload["session_relative_path"] != "execution-ledger/sessions/"+sessionID+".json" ||
		payload["review_receipt_id"] != reviewReceipt.ReceiptID ||
		payload["review_receipt_relative_path"] != reviewReceipt.ReceiptRelativePath ||
		payload["review_receipt_sha256"] != reviewReceipt.ReceiptSHA256 ||
		payload["review_receipt_consumed"] != true ||
		payload["review_receipt_accepted"] != true ||
		payload["review_gate_state"] != "review-receipt-accepted-dispatch-still-gated" ||
		payload["review_gate_ready"] != true ||
		payload["dispatch_state_advance_ready"] != true ||
		payload["launch_authorization_receipt_id"] != launchReceipt.ReceiptID ||
		payload["launch_gate_state"] != "dispatch-preparation-required" ||
		payload["launch_gate_receipt_accepted"] != true ||
		payload["launch_gate_guest_boundary_accepted"] != true ||
		payload["controlled_dispatch_gate_ready"] != false ||
		payload["controlled_dispatch_request_created"] != false ||
		payload["controlled_dispatch_request_state"] != "blocked" ||
		payload["runtime_owned_dispatch_request"] != false ||
		payload["artifact_verified"] != false ||
		payload["dispatch_ready"] != false ||
		payload["request_objects_created"] != false ||
		payload["dispatch_allowed"] != false ||
		payload["dispatch_started"] != false ||
		payload["execution_started"] != false ||
		payload["direct_launch_enabled"] != false ||
		payload["desktop_launch_enabled"] != false ||
		payload["backend_launch_enabled"] != false ||
		payload["backend_process_started"] != false ||
		payload["permission_grant_created"] != false ||
		payload["review_receipt_path_exposed"] != false ||
		payload["receipt_path_exposed"] != false ||
		payload["state_root_path_exposed"] != false ||
		payload["session_path_exposed"] != false ||
		payload["host_root_modified"] != false {
		t.Fatalf("unexpected session-gated controlled dispatch request payload: %#v", payload)
	}
	text := strings.ToLower(output.String())
	if strings.Contains(text, strings.ToLower(stateRoot)) {
		t.Fatalf("session-gated controlled dispatch request exposed state root path: %s", text)
	}
}

func TestKnownAppLaunchAuthorizationReceiptPreviewCommandRequiresDirective(t *testing.T) {
	var output bytes.Buffer
	err := run([]string{
		"known-app-launch-authorization-receipt-preview",
		"--app", "7zr",
		"--state-root", t.TempDir(),
	}, &output)
	if err == nil {
		t.Fatal("expected missing authorization directive to fail")
	}
}

func TestKnownAppLaunchGatePreviewCommandConsumesReceipt(t *testing.T) {
	stateRoot := t.TempDir()
	var receiptOutput bytes.Buffer
	err := run([]string{
		"known-app-launch-authorization-receipt-preview",
		"--app", "7zr",
		"--state-root", stateRoot,
		"--authorize", appidentity.KnownAppLaunchAuthorizationReceiptAction,
	}, &receiptOutput)
	if err != nil {
		t.Fatalf("receipt run returned error: %v", err)
	}
	var receiptPayload map[string]any
	if err := json.Unmarshal(receiptOutput.Bytes(), &receiptPayload); err != nil {
		t.Fatalf("Unmarshal receipt returned error: %v", err)
	}

	var gateOutput bytes.Buffer
	err = run([]string{
		"known-app-launch-gate-preview",
		"--app", "7zr",
		"--state-root", stateRoot,
		"--receipt-id", receiptPayload["receipt_id"].(string),
		"--cache-root", t.TempDir(),
		"--guest-boundary", "managed-known-app-guest-smoke",
	}, &gateOutput)
	if err != nil {
		t.Fatalf("gate run returned error: %v", err)
	}
	var gatePayload map[string]any
	if err := json.Unmarshal(gateOutput.Bytes(), &gatePayload); err != nil {
		t.Fatalf("Unmarshal gate returned error: %v", err)
	}
	if gatePayload["schema_version"] != appidentity.KnownAppLaunchGateSchemaVersion ||
		gatePayload["request_type"] != appidentity.KnownAppLaunchGateRequestType ||
		gatePayload["app_id"] != "7zr" ||
		gatePayload["receipt_lookup_state"] != "accepted-receipt" ||
		gatePayload["receipt_accepted"] != true ||
		gatePayload["receipt_path_exposed"] != false ||
		gatePayload["state_root_path_exposed"] != false ||
		gatePayload["guest_boundary_accepted"] != true ||
		gatePayload["launch_gate_state"] != "dispatch-preparation-required" ||
		gatePayload["launch_gate_blocked_reason"] != "managed artifact preparation is required before dispatch" ||
		gatePayload["controlled_dispatch_ready"] != false ||
		gatePayload["direct_launch_enabled"] != false ||
		gatePayload["desktop_launch_enabled"] != false ||
		gatePayload["backend_launch_enabled"] != false ||
		gatePayload["backend_process_started"] != false {
		t.Fatalf("unexpected gate preview payload: %#v", gatePayload)
	}
	text := strings.ToLower(gateOutput.String())
	if strings.Contains(text, strings.ToLower(stateRoot)) {
		t.Fatalf("gate preview exposed state root path: %s", text)
	}
}

func TestKnownAppLaunchGatePreviewCommandRequiresReceiptID(t *testing.T) {
	var output bytes.Buffer
	err := run([]string{
		"known-app-launch-gate-preview",
		"--app", "7zr",
		"--state-root", t.TempDir(),
	}, &output)
	if err == nil {
		t.Fatal("expected missing receipt id to fail")
	}
}

func TestKnownAppControlledDispatchRequestPreviewCommandBlocksUntilArtifactVerified(t *testing.T) {
	stateRoot := t.TempDir()
	var receiptOutput bytes.Buffer
	err := run([]string{
		"known-app-launch-authorization-receipt-preview",
		"--app", "7zr",
		"--state-root", stateRoot,
		"--authorize", appidentity.KnownAppLaunchAuthorizationReceiptAction,
	}, &receiptOutput)
	if err != nil {
		t.Fatalf("receipt run returned error: %v", err)
	}
	var receiptPayload map[string]any
	if err := json.Unmarshal(receiptOutput.Bytes(), &receiptPayload); err != nil {
		t.Fatalf("Unmarshal receipt returned error: %v", err)
	}

	var output bytes.Buffer
	err = run([]string{
		"known-app-controlled-dispatch-request-preview",
		"--app", "7zr",
		"--state-root", stateRoot,
		"--receipt-id", receiptPayload["receipt_id"].(string),
		"--cache-root", t.TempDir(),
		"--guest-boundary", "managed-known-app-guest-smoke",
	}, &output)
	if err != nil {
		t.Fatalf("controlled dispatch request run returned error: %v", err)
	}
	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal controlled dispatch request returned error: %v", err)
	}
	if payload["schema_version"] != appidentity.KnownAppControlledDispatchSchemaVersion ||
		payload["request_type"] != appidentity.KnownAppControlledDispatchRequestType ||
		payload["app_id"] != "7zr" ||
		payload["receipt_accepted"] != true ||
		payload["guest_boundary_accepted"] != true ||
		payload["launch_gate_state"] != "dispatch-preparation-required" ||
		payload["controlled_dispatch_request_state"] != "blocked" ||
		payload["controlled_dispatch_request_created"] != false ||
		payload["controlled_dispatch_ready"] != false ||
		payload["dispatch_started"] != false ||
		payload["execution_started"] != false ||
		payload["direct_launch_enabled"] != false ||
		payload["desktop_launch_enabled"] != false ||
		payload["backend_launch_enabled"] != false ||
		payload["backend_process_started"] != false ||
		payload["receipt_path_exposed"] != false ||
		payload["state_root_path_exposed"] != false {
		t.Fatalf("unexpected controlled dispatch request payload: %#v", payload)
	}
	text := strings.ToLower(output.String())
	if strings.Contains(text, strings.ToLower(stateRoot)) {
		t.Fatalf("controlled dispatch request preview exposed state root path: %s", text)
	}
}

func TestKnownAppControlledExecutionSessionPreviewCommandBlocksUntilDispatchRequestCreated(t *testing.T) {
	stateRoot := t.TempDir()
	var receiptOutput bytes.Buffer
	err := run([]string{
		"known-app-launch-authorization-receipt-preview",
		"--app", "7zr",
		"--state-root", stateRoot,
		"--authorize", appidentity.KnownAppLaunchAuthorizationReceiptAction,
	}, &receiptOutput)
	if err != nil {
		t.Fatalf("receipt run returned error: %v", err)
	}
	var receiptPayload map[string]any
	if err := json.Unmarshal(receiptOutput.Bytes(), &receiptPayload); err != nil {
		t.Fatalf("Unmarshal receipt returned error: %v", err)
	}

	var output bytes.Buffer
	err = run([]string{
		"known-app-controlled-execution-session-preview",
		"--app", "7zr",
		"--state-root", stateRoot,
		"--receipt-id", receiptPayload["receipt_id"].(string),
		"--cache-root", t.TempDir(),
		"--guest-boundary", "managed-known-app-guest-smoke",
	}, &output)
	if err != nil {
		t.Fatalf("controlled execution session run returned error: %v", err)
	}
	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal controlled execution session returned error: %v", err)
	}
	if payload["schema_version"] != appidentity.KnownAppControlledExecutionSessionSchemaVersion ||
		payload["request_type"] != appidentity.KnownAppControlledExecutionSessionRequestType ||
		payload["source"] != appidentity.KnownAppControlledDispatchRequestType ||
		payload["app_id"] != "7zr" ||
		payload["receipt_accepted"] != true ||
		payload["guest_boundary_accepted"] != true ||
		payload["launch_gate_state"] != "dispatch-preparation-required" ||
		payload["controlled_dispatch_request_state"] != "blocked" ||
		payload["controlled_dispatch_request_created"] != false ||
		payload["runtime_owned_execution_session"] != false ||
		payload["execution_session_request_type"] != "known-app-runtime-execution-session-handoff" ||
		payload["execution_session_state"] != "blocked" ||
		payload["execution_session_handoff_created"] != false ||
		payload["session_handoff_ready"] != false ||
		payload["session_registered"] != false ||
		payload["window_observed"] != false ||
		payload["dispatch_started"] != false ||
		payload["execution_started"] != false ||
		payload["direct_launch_enabled"] != false ||
		payload["desktop_launch_enabled"] != false ||
		payload["backend_launch_enabled"] != false ||
		payload["backend_process_started"] != false ||
		payload["receipt_path_exposed"] != false ||
		payload["state_root_path_exposed"] != false {
		t.Fatalf("unexpected controlled execution session payload: %#v", payload)
	}
	text := strings.ToLower(output.String())
	if strings.Contains(text, strings.ToLower(stateRoot)) {
		t.Fatalf("controlled execution session preview exposed state root path: %s", text)
	}
}

func TestKnownAppControlledExecutionSessionRecordCommandBlocksUntilHandoffReady(t *testing.T) {
	stateRoot := t.TempDir()
	var receiptOutput bytes.Buffer
	err := run([]string{
		"known-app-launch-authorization-receipt-preview",
		"--app", "7zr",
		"--state-root", stateRoot,
		"--authorize", appidentity.KnownAppLaunchAuthorizationReceiptAction,
	}, &receiptOutput)
	if err != nil {
		t.Fatalf("receipt run returned error: %v", err)
	}
	var receiptPayload map[string]any
	if err := json.Unmarshal(receiptOutput.Bytes(), &receiptPayload); err != nil {
		t.Fatalf("Unmarshal receipt returned error: %v", err)
	}

	var output bytes.Buffer
	err = run([]string{
		"known-app-controlled-execution-session-record",
		"--app", "7zr",
		"--state-root", stateRoot,
		"--receipt-id", receiptPayload["receipt_id"].(string),
		"--cache-root", t.TempDir(),
		"--guest-boundary", "managed-known-app-guest-smoke",
	}, &output)
	if err != nil {
		t.Fatalf("controlled execution session record run returned error: %v", err)
	}
	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal controlled execution session record returned error: %v", err)
	}
	if payload["schema_version"] != appidentity.KnownAppControlledExecutionSessionRecordSchemaVersion ||
		payload["request_type"] != appidentity.KnownAppControlledExecutionSessionRecordRequestType ||
		payload["source"] != appidentity.KnownAppControlledExecutionSessionRequestType+"+execution-ledger+execution-session-record" ||
		payload["app_id"] != "7zr" ||
		payload["receipt_accepted"] != true ||
		payload["guest_boundary_accepted"] != true ||
		payload["controlled_dispatch_request_created"] != false ||
		payload["execution_session_handoff_created"] != false ||
		payload["session_handoff_ready"] != false ||
		payload["record_state"] != "blocked" ||
		payload["ledger_record_written"] != false ||
		payload["session_record_written"] != false ||
		payload["session_state"] != "not-recorded" ||
		payload["compatibility_center_state"] != "not-recorded" ||
		payload["runtime_owned_execution_session"] != false ||
		payload["state_root_path_exposed"] != false ||
		payload["transaction_path_exposed"] != false ||
		payload["session_path_exposed"] != false ||
		payload["dispatch_started"] != false ||
		payload["execution_started"] != false ||
		payload["backend_process_started"] != false ||
		payload["host_root_modified"] != false {
		t.Fatalf("unexpected controlled execution session record payload: %#v", payload)
	}
	text := strings.ToLower(output.String())
	if strings.Contains(text, strings.ToLower(stateRoot)) {
		t.Fatalf("controlled execution session record exposed state root path: %s", text)
	}
}

func TestKnownAppControlledExecutionSessionConsumePreviewCommandReadsPersistedRecord(t *testing.T) {
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
		Summary:          "Runtime recorded a known application execution session handoff for later consumption.",
	}); err != nil {
		t.Fatalf("Record returned error: %v", err)
	}
	if _, err := ledger.RecordSession(sessionID); err != nil {
		t.Fatalf("RecordSession returned error: %v", err)
	}

	var output bytes.Buffer
	err = run([]string{
		"known-app-controlled-execution-session-consume-preview",
		"--app", "7zr",
		"--state-root", stateRoot,
	}, &output)
	if err != nil {
		t.Fatalf("controlled execution session consume run returned error: %v", err)
	}
	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal controlled execution session consume returned error: %v", err)
	}
	if payload["schema_version"] != appidentity.KnownAppControlledExecutionSessionConsumeSchemaVersion ||
		payload["request_type"] != appidentity.KnownAppControlledExecutionSessionConsumeRequestType ||
		payload["source"] != appidentity.KnownAppControlledExecutionSessionRecordRequestType+"+execution-session-fanout-evidence" ||
		payload["app_id"] != "7zr" ||
		payload["execution_session_id"] != sessionID ||
		payload["record_consumed"] != true ||
		payload["ledger_record_consumed"] != true ||
		payload["session_record_consumed"] != true ||
		payload["session_digest_verified"] != true ||
		payload["session_state"] != "blocked" ||
		payload["task_manager_state"] != "blocked" ||
		payload["kwin_state"] != "blocked" ||
		payload["tray_state"] != "blocked" ||
		payload["compatibility_center_state"] != "waiting-for-runtime-gates" ||
		payload["fan_out_request_type"] != "execution-session-fanout-evidence" ||
		payload["surface_count"] != float64(4) ||
		payload["runtime_owner_consumable"] != true ||
		payload["kde_read_model_consumable"] != true ||
		payload["safe_for_kde"] != true ||
		payload["state_root_path_exposed"] != false ||
		payload["transaction_path_exposed"] != false ||
		payload["session_path_exposed"] != false ||
		payload["live_state_observed"] != false ||
		payload["session_registered"] != false ||
		payload["window_observed"] != false ||
		payload["task_manager_entry_active"] != false ||
		payload["kwin_rule_applied"] != false ||
		payload["live_tray_bridge_enabled"] != false ||
		payload["dispatch_started"] != false ||
		payload["execution_started"] != false ||
		payload["direct_launch_enabled"] != false ||
		payload["desktop_launch_enabled"] != false ||
		payload["backend_launch_enabled"] != false ||
		payload["backend_process_started"] != false ||
		payload["host_root_modified"] != false {
		t.Fatalf("unexpected controlled execution session consume payload: %#v", payload)
	}
	text := strings.ToLower(output.String())
	if strings.Contains(text, strings.ToLower(stateRoot)) {
		t.Fatalf("controlled execution session consume exposed state root path: %s", text)
	}
}

func TestKnownAppSessionGatedLaunchReviewPreviewCommandConsumesSessionBeforeReview(t *testing.T) {
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
		Summary:          "Runtime recorded a known application execution session handoff for later consumption.",
	}); err != nil {
		t.Fatalf("Record returned error: %v", err)
	}
	if _, err := ledger.RecordSession(sessionID); err != nil {
		t.Fatalf("RecordSession returned error: %v", err)
	}

	var output bytes.Buffer
	err = run([]string{
		"known-app-session-gated-launch-review-preview",
		"--app", "7zr",
		"--state-root", stateRoot,
		"--session-id", sessionID,
		"--action", appidentity.KnownAppSessionGatedLaunchReviewAction,
		"--decision", "approved",
	}, &output)
	if err != nil {
		t.Fatalf("session-gated launch review run returned error: %v", err)
	}
	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal session-gated launch review returned error: %v", err)
	}
	if payload["schema_version"] != appidentity.KnownAppSessionGatedLaunchReviewSchemaVersion ||
		payload["request_type"] != appidentity.KnownAppSessionGatedLaunchReviewRequestType ||
		payload["source"] != appidentity.KnownAppControlledExecutionSessionConsumeRequestType+"+kde-center-page-session-gate-card" ||
		payload["runtime_method"] != "PreviewKnownAppSessionGatedLaunchReview" ||
		payload["read_method"] != "GetKnownAppSessionGatedLaunchReview" ||
		payload["app_id"] != "7zr" ||
		payload["action_id"] != "review-session-gated-dispatch" ||
		payload["action_kind"] != "session-gate-review" ||
		payload["decision"] != "approved" ||
		payload["review_route_created"] != true ||
		payload["review_route_request_type"] != appidentity.KnownAppSessionGatedLaunchReviewRequestType ||
		payload["read_before_write_required"] != true ||
		payload["runtime_receipt_required"] != true ||
		payload["execution_session_id"] != sessionID ||
		payload["session_record_consumed"] != true ||
		payload["session_digest_verified"] != true ||
		payload["session_relative_path"] != "execution-ledger/sessions/"+sessionID+".json" ||
		payload["runtime_owner_consumable"] != true ||
		payload["kde_read_model_consumable"] != true ||
		payload["safe_for_kde"] != true ||
		payload["runtime_owned"] != true ||
		payload["go_runtime_backed"] != true ||
		payload["kde_policy_owner"] != false ||
		payload["user_decision_allows_launch"] != false ||
		payload["runtime_launch_approval"] != false ||
		payload["launch_enabled"] != false ||
		payload["desktop_launch_enabled"] != false ||
		payload["execution_started"] != false ||
		payload["backend_launch_enabled"] != false ||
		payload["backend_process_started"] != false ||
		payload["request_objects_created"] != false ||
		payload["permission_grant_created"] != false ||
		payload["review_receipt_recorded"] != false ||
		payload["state_root_path_exposed"] != false ||
		payload["session_path_exposed"] != false ||
		payload["raw_artifact_path_exposed"] != false ||
		payload["backend_details_exposed"] != false ||
		payload["host_root_modified"] != false {
		t.Fatalf("unexpected session-gated launch review payload: %#v", payload)
	}
	text := strings.ToLower(output.String())
	if strings.Contains(text, strings.ToLower(stateRoot)) {
		t.Fatalf("session-gated launch review exposed state root path: %s", text)
	}
}

func TestKnownAppKDERuntimeStatusLaunchEvidenceRecordCommandPersistsProjectionFile(t *testing.T) {
	stateRoot := t.TempDir()
	evidenceFile := filepath.Join(t.TempDir(), "runtime-status-launch-evidence.json")
	if err := os.WriteFile(evidenceFile, []byte(knownAppRuntimeStatusLaunchProjectionFixture()), 0o600); err != nil {
		t.Fatalf("WriteFile evidence returned error: %v", err)
	}

	var output bytes.Buffer
	if err := run([]string{
		appidentity.KnownAppKDERuntimeStatusLaunchEvidenceRecordRequestType,
		"--state-root", stateRoot,
		"--evidence-file", evidenceFile,
	}, &output); err != nil {
		t.Fatalf("run returned error: %v", err)
	}
	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["schema_version"] != appidentity.KnownAppKDERuntimeStatusLaunchEvidenceRecordSchemaVersion ||
		payload["request_type"] != appidentity.KnownAppKDERuntimeStatusLaunchEvidenceRecordRequestType ||
		payload["runtime_method"] != "RecordKnownAppKDERuntimeStatusLaunchEvidence" ||
		payload["read_method"] != "GetKnownAppKDERuntimeStatusLaunchEvidence" ||
		payload["evidence_state"] != "persisted" ||
		payload["projection_type"] != "known-app-kde-runtime-status-launch-delegated-evidence" ||
		payload["compatibility_center_projection_ready"] != true ||
		payload["kde_center_projection_ready"] != true ||
		payload["known_app_smoke_evidence_ready"] != true ||
		payload["runtime_owned"] != true ||
		payload["runtime_owned_dispatch"] != true ||
		payload["go_runtime_backed"] != true ||
		payload["kde_policy_owner"] != false ||
		payload["state_root_path_exposed"] != false ||
		payload["evidence_path_exposed"] != false ||
		payload["managed_launcher_path_exposed"] != false ||
		payload["raw_launcher_output_exposed"] != false ||
		payload["backend_details_exposed"] != false ||
		payload["host_root_modified"] != false ||
		payload["docker_socket_mounted"] != false ||
		payload["broad_host_mount_required"] != false ||
		payload["desktop_launch_enabled"] != false ||
		payload["backend_launch_enabled"] != false ||
		payload["execution_started"] != false ||
		payload["backend_process_started"] != false {
		t.Fatalf("unexpected Runtime-status launch evidence record payload: %#v", payload)
	}
	relativePath := payload["evidence_relative_path"].(string)
	if filepath.IsAbs(relativePath) || strings.Contains(filepath.Clean(relativePath), "..") {
		t.Fatalf("unsafe evidence relative path: %q", relativePath)
	}
	if _, err := os.Stat(filepath.Join(stateRoot, filepath.FromSlash(relativePath))); err != nil {
		t.Fatalf("expected persisted evidence file: %v", err)
	}
	if strings.Contains(output.String(), stateRoot) || strings.Contains(output.String(), evidenceFile) {
		t.Fatalf("record output exposed local paths: %s", output.String())
	}
	assertWindowIdentityPayloadSafe(t, output.String())
}

func TestKnownAppKDERuntimeStatusLaunchEvidencePreviewCommandConsumesRecordedHandoff(t *testing.T) {
	stateRoot := t.TempDir()
	evidenceFile := filepath.Join(t.TempDir(), "runtime-status-launch-evidence.json")
	if err := os.WriteFile(evidenceFile, []byte(knownAppRuntimeStatusLaunchProjectionFixture()), 0o600); err != nil {
		t.Fatalf("WriteFile evidence returned error: %v", err)
	}
	var recordOutput bytes.Buffer
	if err := run([]string{
		appidentity.KnownAppKDERuntimeStatusLaunchEvidenceRecordRequestType,
		"--state-root", stateRoot,
		"--evidence-file", evidenceFile,
	}, &recordOutput); err != nil {
		t.Fatalf("record run returned error: %v", err)
	}
	var record map[string]any
	if err := json.Unmarshal(recordOutput.Bytes(), &record); err != nil {
		t.Fatalf("Unmarshal record returned error: %v", err)
	}

	var output bytes.Buffer
	if err := run([]string{
		appidentity.KnownAppKDERuntimeStatusLaunchEvidencePreviewRequestType,
		"--state-root", stateRoot,
		"--evidence-relative-path", record["evidence_relative_path"].(string),
	}, &output); err != nil {
		t.Fatalf("preview run returned error: %v", err)
	}
	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal preview returned error: %v", err)
	}
	if payload["schema_version"] != appidentity.KnownAppKDERuntimeStatusLaunchEvidencePreviewSchemaVersion ||
		payload["request_type"] != appidentity.KnownAppKDERuntimeStatusLaunchEvidencePreviewRequestType ||
		payload["runtime_method"] != "PreviewKnownAppKDERuntimeStatusLaunchEvidence" ||
		payload["read_method"] != "GetKnownAppKDERuntimeStatusLaunchEvidence" ||
		payload["evidence_read_state"] != "consumed" ||
		payload["evidence_handoff_consumed"] != true ||
		payload["evidence_relative_path"] != record["evidence_relative_path"] ||
		payload["evidence_sha256"] != record["evidence_sha256"] ||
		payload["evidence_digest_verified"] != true ||
		payload["known_app_smoke_evidence_ready"] != true ||
		payload["runtime_owned"] != true ||
		payload["runtime_owned_dispatch"] != true ||
		payload["go_runtime_backed"] != true ||
		payload["kde_policy_owner"] != false ||
		payload["state_root_path_exposed"] != false ||
		payload["evidence_path_exposed"] != false ||
		payload["raw_launcher_output_exposed"] != false ||
		payload["backend_details_exposed"] != false ||
		payload["host_root_modified"] != false ||
		payload["desktop_launch_enabled"] != false ||
		payload["backend_launch_enabled"] != false ||
		payload["execution_started"] != false {
		t.Fatalf("unexpected Runtime-status launch evidence preview payload: %#v", payload)
	}
	evidence := payload["known_app_smoke_evidence"].(map[string]any)
	if evidence["center_card_state"] != "validated-post-review-dispatch" ||
		evidence["primary_action_id"] != "show-runtime-controlled-launch" ||
		evidence["primary_action_kind"] != "runtime-status" ||
		evidence["desktop_launch_enabled"] != false ||
		evidence["backend_launch_enabled"] != false ||
		evidence["host_root_modified"] != false {
		t.Fatalf("unexpected handoff known app smoke evidence: %#v", evidence)
	}
	if strings.Contains(output.String(), stateRoot) || strings.Contains(output.String(), evidenceFile) {
		t.Fatalf("preview output exposed local paths: %s", output.String())
	}
	assertWindowIdentityPayloadSafe(t, output.String())
}

func TestKnownAppKDERuntimeStatusLaunchActionTriggerCommandAssemblesLaunchRequestFromHandoff(t *testing.T) {
	stateRoot := t.TempDir()
	evidenceFile := filepath.Join(t.TempDir(), "runtime-status-launch-evidence.json")
	if err := os.WriteFile(evidenceFile, []byte(knownAppRuntimeStatusLaunchProjectionFixture()), 0o600); err != nil {
		t.Fatalf("WriteFile evidence returned error: %v", err)
	}
	var recordOutput bytes.Buffer
	if err := run([]string{
		appidentity.KnownAppKDERuntimeStatusLaunchEvidenceRecordRequestType,
		"--state-root", stateRoot,
		"--evidence-file", evidenceFile,
	}, &recordOutput); err != nil {
		t.Fatalf("record run returned error: %v", err)
	}
	var record map[string]any
	if err := json.Unmarshal(recordOutput.Bytes(), &record); err != nil {
		t.Fatalf("Unmarshal record returned error: %v", err)
	}

	var output bytes.Buffer
	if err := run([]string{
		appidentity.KnownAppKDERuntimeStatusLaunchActionTriggerRequestType,
		"--state-root", stateRoot,
		"--evidence-relative-path", record["evidence_relative_path"].(string),
	}, &output); err != nil {
		t.Fatalf("trigger run returned error: %v", err)
	}
	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal trigger returned error: %v", err)
	}
	if payload["schema_version"] != appidentity.KnownAppKDERuntimeStatusLaunchActionTriggerSchemaVersion ||
		payload["request_type"] != appidentity.KnownAppKDERuntimeStatusLaunchActionTriggerRequestType ||
		payload["runtime_method"] != "PreviewKnownAppKDERuntimeStatusLaunchActionTrigger" ||
		payload["read_method"] != "GetKnownAppKDERuntimeStatusLaunchActionTrigger" ||
		payload["action_id"] != appidentity.KnownAppKDERuntimeStatusLaunchAction ||
		payload["action_kind"] != "runtime-status" ||
		payload["trigger_state"] != "runtime-launch-request-assembled" ||
		payload["evidence_read_state"] != "consumed" ||
		payload["evidence_handoff_consumed"] != true ||
		payload["evidence_relative_path"] != record["evidence_relative_path"] ||
		payload["evidence_sha256"] != record["evidence_sha256"] ||
		payload["evidence_digest_verified"] != true ||
		payload["launch_request_type"] != appidentity.KnownAppKDERuntimeStatusLaunchRequestType ||
		payload["launch_request_runtime_method"] != "PreviewKnownAppKDERuntimeStatusLaunchRequest" ||
		payload["launch_request_read_method"] != "GetKnownAppKDERuntimeStatusLaunchRequest" ||
		payload["launch_request_created"] != true ||
		payload["managed_launcher_argv_ready"] != true ||
		payload["required_opaque_id_count"] != float64(3) ||
		payload["collected_opaque_id_count"] != float64(3) ||
		payload["runtime_owned_trigger"] != true ||
		payload["runtime_owned_request"] != true ||
		payload["runtime_owned_launch"] != true ||
		payload["runtime_owned_dispatch"] != true ||
		payload["kde_presentation_only"] != true ||
		payload["kde_action_forwarded"] != true ||
		payload["state_root_required"] != true ||
		payload["state_root_supplied_by_runtime"] != true ||
		payload["kde_state_root_access"] != false ||
		payload["direct_launch_enabled"] != false ||
		payload["desktop_launch_enabled"] != false ||
		payload["backend_launch_enabled"] != false ||
		payload["execution_started"] != false ||
		payload["backend_process_started"] != false ||
		payload["request_objects_created"] != false ||
		payload["permission_grant_created"] != false ||
		payload["state_root_path_exposed"] != false ||
		payload["evidence_path_exposed"] != false ||
		payload["receipt_path_exposed"] != false ||
		payload["session_path_exposed"] != false ||
		payload["raw_artifact_path_exposed"] != false ||
		payload["raw_command_exposed"] != false ||
		payload["raw_launcher_output_exposed"] != false ||
		payload["backend_details_exposed"] != false ||
		payload["host_root_modified"] != false {
		t.Fatalf("unexpected Runtime-status launch action trigger payload: %#v", payload)
	}
	argv := payload["managed_launcher_argv"].([]any)
	if len(argv) != 11 ||
		argv[0] != "xnix-compat-launch" ||
		argv[2] != "7zr" ||
		argv[4] != "managed-known-app-guest-smoke" ||
		argv[6] != "known-app-launch-authorization-7zr-26.02" ||
		argv[8] != "known-app-session-gated-launch-review-7zr-26.02-known-app-controlled-execution-session-7zr-26.02" ||
		argv[10] != "known-app-controlled-execution-session-7zr-26.02" {
		t.Fatalf("unexpected trigger managed launcher argv: %#v", argv)
	}
	launchRequest := payload["launch_request"].(map[string]any)
	if launchRequest["request_type"] != appidentity.KnownAppKDERuntimeStatusLaunchRequestType ||
		launchRequest["action_id"] != appidentity.KnownAppKDERuntimeStatusLaunchAction ||
		launchRequest["launch_request_created"] != true ||
		launchRequest["state_root_supplied_by_runtime"] != true ||
		launchRequest["kde_state_root_access"] != false ||
		launchRequest["execution_started"] != false {
		t.Fatalf("unexpected nested launch request: %#v", launchRequest)
	}
	if strings.Contains(output.String(), stateRoot) || strings.Contains(output.String(), evidenceFile) {
		t.Fatalf("trigger output exposed local paths: %s", output.String())
	}
	assertWindowIdentityPayloadSafe(t, output.String())
}

func writeKnownAppRuntimeStatusLaunchExecutionCLIFixture(t *testing.T, stateRoot string) string {
	t.Helper()
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
		Summary:          "Runtime recorded a known application execution session handoff for later consumption.",
	}); err != nil {
		t.Fatalf("Record returned error: %v", err)
	}
	if _, err := ledger.RecordSession(sessionID); err != nil {
		t.Fatalf("RecordSession returned error: %v", err)
	}
	if _, err := appidentity.RecordKnownAppLaunchAuthorizationReceipt(appidentity.KnownAppLaunchAuthorizationReceiptRequest{
		AppID:     "7zr",
		StateRoot: stateRoot,
		Authorize: appidentity.KnownAppLaunchAuthorizationReceiptAction,
	}); err != nil {
		t.Fatalf("RecordKnownAppLaunchAuthorizationReceipt returned error: %v", err)
	}
	if _, err := appidentity.RecordKnownAppSessionGatedLaunchReviewReceipt(appidentity.KnownAppSessionGatedLaunchReviewReceiptRequest{
		AppID:     "7zr",
		StateRoot: stateRoot,
		SessionID: sessionID,
		ActionID:  appidentity.KnownAppSessionGatedLaunchReviewAction,
		Decision:  "approved",
	}); err != nil {
		t.Fatalf("RecordKnownAppSessionGatedLaunchReviewReceipt returned error: %v", err)
	}
	return sessionID
}

func writeFakeRuntimeStatusManagedLauncher(t *testing.T) (string, string) {
	t.Helper()
	dir := t.TempDir()
	launcherPath := filepath.Join(dir, "fake-xnix-compat-launch")
	argsPath := filepath.Join(dir, "launcher-args.txt")
	t.Setenv("XNIX_FAKE_LAUNCHER_ARGS_FILE", argsPath)
	script := "#!/bin/sh\nprintf '%s\n' \"$@\" > \"$XNIX_FAKE_LAUNCHER_ARGS_FILE\"\nprintf '%s\n' '{\"request_type\":\"windows-known-app-dispatch-smoke\",\"status\":\"passed\",\"guest_boundary\":\"managed-known-app-guest-smoke\",\"runtime_owned_dispatch\":true,\"artifact_verified\":true,\"marker_observed\":true,\"smoke_passed\":true,\"execution_started\":true,\"backend_process_started\":false,\"session_gated_controlled_dispatch_consumed\":true,\"session_gated_controlled_dispatch_state\":\"created-after-session-gated-review\",\"session_gated_review_receipt_id\":\"known-app-session-gated-launch-review-7zr-26.02-known-app-controlled-execution-session-7zr-26.02\",\"launch_authorization_receipt_id\":\"known-app-launch-authorization-7zr-26.02\",\"controlled_execution_session_consumed\":true,\"controlled_execution_session_id\":\"known-app-controlled-execution-session-7zr-26.02\",\"controlled_session_digest_verified\":true,\"controlled_session_relative_path\":\"execution-ledger/sessions/known-app-controlled-execution-session-7zr-26.02.json\",\"runtime_owner_consumable_session\":true,\"kde_read_model_consumable_session\":true,\"controlled_session_live_state_observed\":false,\"controlled_session_registered\":false,\"controlled_session_window_observed\":false,\"controlled_session_host_root_modified\":false,\"controlled_session_backend_process_start\":false,\"host_root_modified\":false,\"docker_socket_mounted\":false,\"broad_host_mount_required\":false,\"raw_command_exposed\":false,\"backend_details_exposed\":false}'\n"
	if err := os.WriteFile(launcherPath, []byte(script), 0o755); err != nil {
		t.Fatalf("WriteFile fake launcher returned error: %v", err)
	}
	return launcherPath, argsPath
}
