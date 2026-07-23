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
