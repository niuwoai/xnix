package main

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"xnix.local/xnix/internal/runtime/appidentity"
)

func TestKDECenterPagePreviewCommandConsumesSessionRoot(t *testing.T) {
	registryPath, app := writeTestRepairGroupRegistry(t)
	sessionRoot, requestID := writeExecutionSessionRecordWithCLI(t, app)

	var output bytes.Buffer
	if err := run([]string{"kde-center-page-preview", "--registry", registryPath, "--app", app, "--decision", "approved", "--session-root", sessionRoot, "--session-request-id", requestID}, &output); err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["schema_version"] != "xnix.runtime.kde_center_page.v1" ||
		payload["request_type"] != "kde-center-page-preview" ||
		payload["runtime_method"] != "GetKDECenterPage" ||
		payload["read_method"] != "GetKDECenterPagePreview" ||
		payload["application_id"] != app {
		t.Fatalf("unexpected center page schema: %#v", payload)
	}
	if payload["source"] != "compatibility-center-preview+backend-selection-preview+desktop-activation-status-preview+execution-readiness-preview+application-readiness-preview+launch-intent-preview+window-identity-preview+file-association-plan+tray-status-preview+notification-preview+kde-action-card-deck-preview+kde-action-dependency-graph-preview+settings-preview+execution-session-record" {
		t.Fatalf("unexpected center page source: %#v", payload)
	}

	readiness := payload["application_readiness_evidence"].(map[string]any)
	if readiness["schema_version"] != "xnix.runtime.application_readiness.v1" ||
		readiness["request_type"] != "application-readiness-preview" ||
		readiness["graph_type"] != "runtime-application-readiness-evidence-graph" ||
		readiness["runtime_method"] != "GetApplicationReadiness" ||
		readiness["node_count"] != float64(7) ||
		readiness["ready"] != false ||
		readiness["launch_allowed"] != false ||
		readiness["launch_enabled"] != false ||
		readiness["execution_request_created"] != false ||
		readiness["execution_started"] != false ||
		readiness["backend_launch_enabled"] != false ||
		readiness["backend_process_started"] != false ||
		readiness["real_portal_transport_enabled"] != false ||
		readiness["request_object_created"] != false ||
		readiness["permission_granted"] != false ||
		readiness["snapshot_created"] != false ||
		readiness["restore_executed"] != false ||
		readiness["host_root_modified"] != false ||
		readiness["network_required"] != false ||
		readiness["privileged_container_required"] != false ||
		readiness["state_root_path_exposed"] != false ||
		readiness["backend_details_exposed"] != false ||
		readiness["raw_command_exposed"] != false ||
		readiness["raw_executable_exposed"] != false {
		t.Fatalf("unexpected application readiness evidence: %#v", readiness)
	}

	window := payload["window_identity_snapshot"].(map[string]any)
	if window["execution_session_root"] != true ||
		window["execution_session_backed"] != true ||
		window["execution_session_path"] != "execution-ledger/sessions/"+requestID+".json" ||
		window["task_manager_session_state"] != "blocked" ||
		window["kwin_session_state"] != "blocked" ||
		window["task_manager_entry_active"] != false ||
		window["kwin_rule_applied"] != false ||
		window["host_root_modified"] != false ||
		window["backend_details_exposed"] != false {
		t.Fatalf("unexpected session-backed window snapshot: %#v", window)
	}

	tray := payload["tray_status_snapshot"].(map[string]any)
	if tray["compatibility_state"] != "waiting-for-runtime-gates" ||
		tray["compatibility_label"] != "Runtime gates required" ||
		tray["execution_session_root"] != true ||
		tray["execution_session_backed"] != true ||
		tray["execution_session_path"] != "execution-ledger/sessions/"+requestID+".json" ||
		tray["execution_session_state"] != "blocked" ||
		tray["live_backend_bridge_enabled"] != false ||
		tray["host_root_modified"] != false ||
		tray["backend_details_exposed"] != false {
		t.Fatalf("unexpected session-backed tray snapshot: %#v", tray)
	}
	if payload["launch_enabled"] != false ||
		payload["execution_started"] != false ||
		payload["backend_process_started"] != false ||
		payload["request_objects_created"] != false ||
		payload["host_root_modified"] != false ||
		payload["backend_details_exposed"] != false {
		t.Fatalf("session-backed center page must remain gated: %#v", payload)
	}
	if strings.Contains(output.String(), sessionRoot) {
		t.Fatalf("center page exposed session root: %s", output.String())
	}
	assertWindowIdentityPayloadSafe(t, output.String())
}

func TestKDECenterPagePreviewCommandSurfacesKnownAppLauncherSessionGateCards(t *testing.T) {
	registryPath, app := writeTestRepairGroupRegistry(t)
	sessionID := appidentity.KnownAppControlledExecutionSessionID("7zr", "26.02")

	var output bytes.Buffer
	if err := run([]string{
		"kde-center-page-preview",
		"--registry", registryPath,
		"--app", app,
		"--decision", "approved",
		"--known-app-smoke-app", "7zr",
		"--known-app-smoke-name", "7-Zip Console",
		"--known-app-smoke-version", "26.02",
		"--known-app-smoke-source", "staged-launcher-dispatch-smoke",
		"--known-app-smoke-status", "passed",
		"--known-app-smoke-marker-observed",
		"--known-app-smoke-checksum-verified",
		"--known-app-launch-authorization-receipt-state", "recorded",
		"--known-app-launch-authorization-receipt-id", "known-app-launch-authorization-7zr-26.02",
		"--known-app-launch-gate-consumed",
		"--known-app-launch-gate-receipt-accepted",
		"--known-app-launch-gate-guest-boundary-accepted",
		"--known-app-controlled-dispatch-ready",
		"--known-app-controlled-execution-session-id", sessionID,
		"--known-app-launcher-session-gate-consumed",
		"--known-app-launcher-session-digest-verified",
		"--known-app-launcher-session-relative-path", "execution-ledger/sessions/" + sessionID + ".json",
		"--known-app-launcher-session-runtime-owner-consumable",
		"--known-app-launcher-session-kde-read-model-consumable",
		"--known-app-post-review-dispatch-consumed",
		"--known-app-post-review-dispatch-state", "created-after-session-gated-review",
		"--known-app-session-gated-review-receipt-id", "known-app-session-gated-launch-review-7zr-26.02-" + sessionID,
	}, &output); err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["source"] != "compatibility-center-preview+backend-selection-preview+desktop-activation-status-preview+execution-readiness-preview+application-readiness-preview+launch-intent-preview+window-identity-preview+file-association-plan+tray-status-preview+notification-preview+kde-action-card-deck-preview+kde-action-dependency-graph-preview+settings-preview+known-app-session-gate-evidence" ||
		payload["known_app_session_gate_evidence_count"] != float64(1) ||
		payload["known_app_launcher_session_gate_consumed_count"] != float64(1) ||
		payload["known_app_post_review_dispatch_consumed_count"] != float64(1) ||
		payload["launch_enabled"] != false ||
		payload["execution_started"] != false ||
		payload["backend_process_started"] != false ||
		payload["host_root_modified"] != false ||
		payload["backend_details_exposed"] != false {
		t.Fatalf("unexpected known app session gate page payload: %#v", payload)
	}
	cards := payload["known_app_session_gate_cards"].([]any)
	if len(cards) != 1 {
		t.Fatalf("unexpected known app session gate cards: %#v", cards)
	}
	card := cards[0].(map[string]any)
	if card["app_id"] != "7zr" ||
		card["center_card_state"] != "validated-post-review-dispatch" ||
		card["controlled_execution_session_id"] != sessionID ||
		card["launch_authorization_receipt_id"] != "known-app-launch-authorization-7zr-26.02" ||
		card["launcher_session_gate_consumed"] != true ||
		card["launcher_session_digest_verified"] != true ||
		card["launcher_session_relative_path"] != "execution-ledger/sessions/"+sessionID+".json" ||
		card["runtime_owner_consumable_session"] != true ||
		card["kde_read_model_consumable_session"] != true ||
		card["post_review_dispatch_consumed"] != true ||
		card["post_review_dispatch_state"] != "created-after-session-gated-review" ||
		card["session_gated_review_receipt_id"] != "known-app-session-gated-launch-review-7zr-26.02-"+sessionID ||
		card["primary_action_id"] != "show-runtime-controlled-launch" ||
		card["primary_action_kind"] != "runtime-status" ||
		card["runtime_status_launch_request_type"] != "known-app-kde-runtime-status-launch-request-preview" ||
		card["runtime_status_launch_runtime_method"] != "PreviewKnownAppKDERuntimeStatusLaunchRequest" ||
		card["runtime_status_launch_read_method"] != "GetKnownAppKDERuntimeStatusLaunchRequest" ||
		card["runtime_status_launch_required_id_count"] != float64(3) ||
		card["runtime_status_launch_collected_id_count"] != float64(3) ||
		card["runtime_status_launch_request_ready"] != true ||
		card["runtime_status_launch_state_root_required"] != true ||
		card["runtime_status_launch_state_root_owned_by_runtime"] != true ||
		card["review_route_request_type"] != "known-app-session-gated-launch-review-preview" ||
		card["review_route_runtime_method"] != "PreviewKnownAppSessionGatedLaunchReview" ||
		card["review_route_read_method"] != "GetKnownAppSessionGatedLaunchReview" ||
		card["read_before_write_required"] != true ||
		card["runtime_receipt_required"] != true ||
		card["desktop_launch_enabled"] != false ||
		card["backend_launch_enabled"] != false ||
		card["host_root_modified"] != false ||
		card["backend_details_exposed"] != false {
		t.Fatalf("unexpected known app session gate card: %#v", card)
	}
	managedArgv := card["runtime_status_launch_managed_launcher_argv"].([]any)
	if len(managedArgv) != 11 ||
		managedArgv[0] != "xnix-compat-launch" ||
		managedArgv[1] != "--app" ||
		managedArgv[2] != "7zr" ||
		managedArgv[3] != "--guest-boundary" ||
		managedArgv[4] != "managed-known-app-guest-smoke" ||
		managedArgv[5] != "--receipt-id" ||
		managedArgv[6] != "known-app-launch-authorization-7zr-26.02" ||
		managedArgv[7] != "--review-receipt-id" ||
		managedArgv[8] != "known-app-session-gated-launch-review-7zr-26.02-"+sessionID ||
		managedArgv[9] != "--session-id" ||
		managedArgv[10] != sessionID {
		t.Fatalf("unexpected Runtime-status launch managed argv: %#v", managedArgv)
	}
	assertWindowIdentityPayloadSafe(t, output.String())
}

func TestKDECenterPagePreviewCommandAcceptsRuntimeProjectedKnownAppEvidenceFile(t *testing.T) {
	registryPath, app := writeTestRepairGroupRegistry(t)
	evidencePath := filepath.Join(t.TempDir(), "runtime-status-launch-evidence.json")
	if err := os.WriteFile(evidencePath, []byte(knownAppRuntimeStatusLaunchProjectionFixture()), 0o600); err != nil {
		t.Fatalf("WriteFile evidence returned error: %v", err)
	}

	var output bytes.Buffer
	if err := run([]string{
		"kde-center-page-preview",
		"--registry", registryPath,
		"--app", app,
		"--decision", "approved",
		"--known-app-evidence-file", evidencePath,
	}, &output); err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["known_app_session_gate_evidence_count"] != float64(1) ||
		payload["known_app_launcher_session_gate_consumed_count"] != float64(1) ||
		payload["known_app_post_review_dispatch_consumed_count"] != float64(1) ||
		payload["launch_enabled"] != false ||
		payload["execution_started"] != false ||
		payload["backend_process_started"] != false ||
		payload["host_root_modified"] != false ||
		payload["backend_details_exposed"] != false {
		t.Fatalf("unexpected Runtime-projected KDE Center page payload: %#v", payload)
	}
	cards := payload["known_app_session_gate_cards"].([]any)
	card := cards[0].(map[string]any)
	if card["app_id"] != "7zr" ||
		card["center_card_state"] != "validated-post-review-dispatch" ||
		card["controlled_execution_session_id"] != "known-app-controlled-execution-session-7zr-26.02" ||
		card["launch_authorization_receipt_id"] != "known-app-launch-authorization-7zr-26.02" ||
		card["launcher_session_gate_consumed"] != true ||
		card["launcher_session_digest_verified"] != true ||
		card["runtime_owner_consumable_session"] != true ||
		card["kde_read_model_consumable_session"] != true ||
		card["post_review_dispatch_consumed"] != true ||
		card["post_review_dispatch_state"] != "created-after-session-gated-review" ||
		card["primary_action_id"] != "show-runtime-controlled-launch" ||
		card["primary_action_kind"] != "runtime-status" ||
		card["desktop_launch_enabled"] != false ||
		card["backend_launch_enabled"] != false ||
		card["host_root_modified"] != false {
		t.Fatalf("unexpected Runtime-projected KDE Center page card: %#v", card)
	}
	assertWindowIdentityPayloadSafe(t, output.String())
}
