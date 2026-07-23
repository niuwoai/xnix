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
		"--known-app-post-review-dispatch-consumed",
		"--known-app-post-review-dispatch-state", "created-after-session-gated-review",
		"--known-app-session-gated-review-receipt-id", "known-app-session-gated-launch-review-7zr-26.02-known-app-controlled-execution-session-7zr-26.02",
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
		payload["known_app_post_review_dispatch_consumed_count"] != float64(1) ||
		payload["backend_launch_enabled"] != false ||
		payload["host_root_modified"] != false {
		t.Fatalf("unexpected Compatibility Center launch gate summary: %#v", payload)
	}
	evidenceItems := payload["known_app_smoke_evidence"].([]any)
	evidence := evidenceItems[0].(map[string]any)
	if evidence["center_card_state"] != "validated-post-review-dispatch" ||
		evidence["primary_action_id"] != "show-runtime-controlled-launch" ||
		evidence["primary_action_kind"] != "runtime-status" ||
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
		evidence["post_review_dispatch_consumed"] != true ||
		evidence["post_review_dispatch_state"] != "created-after-session-gated-review" ||
		evidence["session_gated_review_receipt_id"] != "known-app-session-gated-launch-review-7zr-26.02-known-app-controlled-execution-session-7zr-26.02" ||
		evidence["direct_launch_enabled"] != false ||
		evidence["desktop_launch_enabled"] != false ||
		evidence["backend_launch_enabled"] != false {
		t.Fatalf("unexpected Compatibility Center launch gate evidence: %#v", evidence)
	}
}

func TestCompatibilityCenterPreviewCommandAcceptsRuntimeProjectedKnownAppEvidenceJSON(t *testing.T) {
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
		"--known-app-evidence-json", knownAppRuntimeStatusLaunchProjectionFixture(),
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
		payload["known_app_post_review_dispatch_consumed_count"] != float64(1) ||
		payload["backend_launch_enabled"] != false ||
		payload["host_root_modified"] != false {
		t.Fatalf("unexpected Runtime-projected Compatibility Center summary: %#v", payload)
	}
	evidenceItems := payload["known_app_smoke_evidence"].([]any)
	evidence := evidenceItems[0].(map[string]any)
	if evidence["evidence_source"] != "staged-launcher-dispatch-smoke" ||
		evidence["center_card_state"] != "validated-post-review-dispatch" ||
		evidence["primary_action_id"] != "show-runtime-controlled-launch" ||
		evidence["launch_authorization_receipt_state"] != "recorded" ||
		evidence["launch_gate_state"] != "controlled-dispatch-ready" ||
		evidence["launch_gate_consumed"] != true ||
		evidence["controlled_dispatch_ready"] != true ||
		evidence["launcher_session_gate_consumed"] != true ||
		evidence["post_review_dispatch_consumed"] != true ||
		evidence["direct_launch_enabled"] != false ||
		evidence["desktop_launch_enabled"] != false ||
		evidence["backend_launch_enabled"] != false {
		t.Fatalf("unexpected Runtime-projected Compatibility Center evidence: %#v", evidence)
	}
}

func knownAppRuntimeStatusLaunchProjectionFixture() string {
	return `{
  "projection_type": "known-app-kde-runtime-status-launch-delegated-evidence",
  "runtime_method": "ProjectKnownAppKDERuntimeStatusLaunchDelegatedEvidence",
  "request_type": "windows-known-app-dispatch-smoke",
  "status": "passed",
  "app_id": "7zr",
  "display_name": "7-Zip Console",
  "app_version": "26.02",
  "guest_boundary": "managed-known-app-guest-smoke",
  "runtime_owned_dispatch": true,
  "artifact_verified": true,
  "marker_observed": true,
  "session_gated_controlled_dispatch_consumed": true,
  "session_gated_controlled_dispatch_state": "created-after-session-gated-review",
  "session_gated_review_receipt_id": "known-app-session-gated-launch-review-7zr-26.02-known-app-controlled-execution-session-7zr-26.02",
  "launch_authorization_receipt_id": "known-app-launch-authorization-7zr-26.02",
  "launch_authorization_receipt_state": "recorded",
  "launch_gate_state": "controlled-dispatch-ready",
  "launch_gate_consumed": true,
  "launch_gate_receipt_accepted": true,
  "launch_gate_guest_boundary_accepted": true,
  "controlled_dispatch_ready": true,
  "controlled_execution_session_consumed": true,
  "controlled_execution_session_id": "known-app-controlled-execution-session-7zr-26.02",
  "controlled_session_digest_verified": true,
  "controlled_session_relative_path": "execution-ledger/sessions/known-app-controlled-execution-session-7zr-26.02.json",
  "runtime_owner_consumable_session": true,
  "kde_read_model_consumable_session": true,
  "controlled_session_live_state_observed": false,
  "controlled_session_registered": false,
  "controlled_session_window_observed": false,
  "controlled_session_host_root_modified": false,
  "controlled_session_backend_process_start": false,
  "host_root_modified": false,
  "docker_socket_mounted": false,
  "broad_host_mount_required": false,
  "state_root_path_exposed": false,
  "managed_launcher_path_exposed": false,
  "raw_launcher_output_exposed": false,
  "backend_details_exposed": false,
  "compatibility_center_projection_ready": true,
  "kde_center_projection_ready": true,
  "desktop_safe_summary": "7-Zip Console delegated Runtime launcher evidence is projected for Compatibility Center and KDE Center consumption without exposing launcher output or state-root paths."
}`
}
