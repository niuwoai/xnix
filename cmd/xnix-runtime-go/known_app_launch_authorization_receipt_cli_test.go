package main

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"

	"xnix.local/xnix/internal/runtime/appidentity"
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
