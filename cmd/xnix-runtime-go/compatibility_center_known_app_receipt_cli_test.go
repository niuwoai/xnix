package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

func TestCompatibilityCenterPreviewCommandAcceptsKnownAppLaunchAuthorizationReceiptState(t *testing.T) {
	root := t.TempDir()
	recipeData := []byte(`{"id":"org.example.ledger","name":"Example Ledger","icon":"office-chart-area","mode":"automatic","supported_extensions":[".abc"]}`)
	sum := sha256.Sum256(recipeData)
	digest := hex.EncodeToString(sum[:])
	if err := os.WriteFile(filepath.Join(root, "org.example.ledger.json"), recipeData, 0o600); err != nil {
		t.Fatalf("WriteFile recipe returned error: %v", err)
	}
	registryPath := filepath.Join(root, "registry.json")
	registryData := []byte(`{"schema_version":1,"registry_name":"test-registry","recipes":[{"id":"org.example.ledger","path":"org.example.ledger.json","sha256":"` + digest + `","signature_status":"development-only"}]}`)
	if err := os.WriteFile(registryPath, registryData, 0o600); err != nil {
		t.Fatalf("WriteFile registry returned error: %v", err)
	}

	var output bytes.Buffer
	err := run([]string{
		"compatibility-center-preview",
		"--registry", registryPath,
		"--known-app-smoke-app", "7zr",
		"--known-app-smoke-name", "7-Zip Console",
		"--known-app-smoke-version", "26.02",
		"--known-app-smoke-source", "staged-launcher-dispatch-smoke",
		"--known-app-smoke-status", "passed",
		"--known-app-smoke-marker-observed",
		"--known-app-smoke-checksum-verified",
		"--known-app-launch-authorization-receipt-state", "recorded",
		"--known-app-launch-authorization-receipt-id", "known-app-launch-authorization-7zr-26.02",
	}, &output)
	if err != nil {
		t.Fatalf("run returned error: %v", err)
	}
	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["known_app_launch_authorization_recorded_count"] != float64(1) ||
		payload["backend_launch_enabled"] != false ||
		payload["host_root_modified"] != false ||
		payload["backend_details_exposed"] != false {
		t.Fatalf("unexpected Compatibility Center receipt summary: %#v", payload)
	}
	evidenceItems := payload["known_app_smoke_evidence"].([]any)
	evidence := evidenceItems[0].(map[string]any)
	if evidence["center_card_state"] != "validated-launch-authorization-recorded" ||
		evidence["launch_authorization_state"] != "recorded" ||
		evidence["primary_action_id"] != "run-through-launch-gate" ||
		evidence["primary_action_kind"] != "launch-gate-review" ||
		evidence["launch_authorization_receipt_state"] != "recorded" ||
		evidence["launch_authorization_receipt_id"] != "known-app-launch-authorization-7zr-26.02" ||
		evidence["launch_gate_state"] != "receipt-recorded-launch-still-gated" ||
		evidence["direct_launch_enabled"] != false ||
		evidence["desktop_launch_enabled"] != false ||
		evidence["backend_launch_enabled"] != false {
		t.Fatalf("unexpected Compatibility Center receipt evidence: %#v", evidence)
	}
}

func TestCompatibilityCenterPreviewCommandAcceptsKnownAppLaunchGateState(t *testing.T) {
	root := t.TempDir()
	recipeData := []byte(`{"id":"org.example.ledger","name":"Example Ledger","icon":"office-chart-area","mode":"automatic","supported_extensions":[".abc"]}`)
	sum := sha256.Sum256(recipeData)
	digest := hex.EncodeToString(sum[:])
	if err := os.WriteFile(filepath.Join(root, "org.example.ledger.json"), recipeData, 0o600); err != nil {
		t.Fatalf("WriteFile recipe returned error: %v", err)
	}
	registryPath := filepath.Join(root, "registry.json")
	registryData := []byte(`{"schema_version":1,"registry_name":"test-registry","recipes":[{"id":"org.example.ledger","path":"org.example.ledger.json","sha256":"` + digest + `","signature_status":"development-only"}]}`)
	if err := os.WriteFile(registryPath, registryData, 0o600); err != nil {
		t.Fatalf("WriteFile registry returned error: %v", err)
	}

	var output bytes.Buffer
	err := run([]string{
		"compatibility-center-preview",
		"--registry", registryPath,
		"--known-app-smoke-app", "7zr",
		"--known-app-smoke-name", "7-Zip Console",
		"--known-app-smoke-version", "26.02",
		"--known-app-smoke-source", "staged-launcher-dispatch-smoke",
		"--known-app-smoke-status", "passed",
		"--known-app-smoke-marker-observed",
		"--known-app-smoke-checksum-verified",
		"--known-app-launch-authorization-receipt-state", "recorded",
		"--known-app-launch-authorization-receipt-id", "known-app-launch-authorization-7zr-26.02",
		"--known-app-launch-gate-state", "controlled-dispatch-ready",
		"--known-app-launch-gate-consumed",
		"--known-app-launch-gate-receipt-accepted",
		"--known-app-launch-gate-guest-boundary-accepted",
		"--known-app-controlled-dispatch-ready",
		"--known-app-controlled-execution-session-id", "known-app-controlled-execution-session-7zr-26.02",
		"--known-app-launcher-session-gate-consumed",
		"--known-app-launcher-session-digest-verified",
		"--known-app-launcher-session-relative-path", "execution-ledger/sessions/known-app-controlled-execution-session-7zr-26.02.json",
		"--known-app-launcher-session-runtime-owner-consumable",
		"--known-app-launcher-session-kde-read-model-consumable",
	}, &output)
	if err != nil {
		t.Fatalf("run returned error: %v", err)
	}
	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["known_app_launch_gate_consumed_count"] != float64(1) ||
		payload["known_app_controlled_dispatch_ready_count"] != float64(1) ||
		payload["known_app_launcher_session_gate_consumed_count"] != float64(1) ||
		payload["backend_launch_enabled"] != false ||
		payload["host_root_modified"] != false {
		t.Fatalf("unexpected Compatibility Center launch gate summary: %#v", payload)
	}
	evidenceItems := payload["known_app_smoke_evidence"].([]any)
	evidence := evidenceItems[0].(map[string]any)
	if evidence["center_card_state"] != "validated-session-gated-dispatch" ||
		evidence["primary_action_id"] != "review-session-gated-dispatch" ||
		evidence["primary_action_kind"] != "session-gate-review" ||
		evidence["launch_gate_state"] != "controlled-dispatch-ready" ||
		evidence["launch_gate_consumed"] != true ||
		evidence["launch_gate_receipt_accepted"] != true ||
		evidence["launch_gate_guest_boundary_accepted"] != true ||
		evidence["controlled_dispatch_ready"] != true ||
		evidence["controlled_execution_session_id"] != "known-app-controlled-execution-session-7zr-26.02" ||
		evidence["launcher_session_gate_consumed"] != true ||
		evidence["launcher_session_digest_verified"] != true ||
		evidence["launcher_session_relative_path"] != "execution-ledger/sessions/known-app-controlled-execution-session-7zr-26.02.json" ||
		evidence["launcher_session_runtime_owner_consumable"] != true ||
		evidence["launcher_session_kde_read_model_consumable"] != true ||
		evidence["direct_launch_enabled"] != false ||
		evidence["desktop_launch_enabled"] != false ||
		evidence["backend_launch_enabled"] != false {
		t.Fatalf("unexpected Compatibility Center launch gate evidence: %#v", evidence)
	}
}
