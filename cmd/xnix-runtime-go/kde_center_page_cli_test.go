package main

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"
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
	sessionID := "xnix-known-app-session-7zr-26-02"

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
		"--known-app-launch-authorization-receipt-id", "xnix-known-app-launch-authorization-7zr-26-02",
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
		card["center_card_state"] != "validated-session-gated-dispatch" ||
		card["controlled_execution_session_id"] != sessionID ||
		card["launcher_session_gate_consumed"] != true ||
		card["launcher_session_digest_verified"] != true ||
		card["launcher_session_relative_path"] != "execution-ledger/sessions/"+sessionID+".json" ||
		card["runtime_owner_consumable_session"] != true ||
		card["kde_read_model_consumable_session"] != true ||
		card["primary_action_id"] != "review-session-gated-dispatch" ||
		card["primary_action_kind"] != "session-gate-review" ||
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
	assertWindowIdentityPayloadSafe(t, output.String())
}
