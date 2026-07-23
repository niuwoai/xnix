package main

import (
	"bytes"
	"encoding/json"
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
