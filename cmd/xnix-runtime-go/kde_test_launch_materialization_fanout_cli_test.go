package main

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"
)

func TestKDETestLaunchMaterializationFanOutPreviewCommandCoversKDESurfaces(t *testing.T) {
	stateRoot := t.TempDir()
	var output bytes.Buffer
	err := run([]string{"kde-test-launch-materialization-fanout-preview", "--registry", "../../runtime/recipes/registry.json", "--app", "org.xnix.sample.notepad", "--state-root", stateRoot, "--mode", "test-only", "--authorize", "authorize-restricted-test-preparation"}, &output)
	if err != nil {
		t.Fatalf("kde-test-launch-materialization-fanout-preview returned error: %v", err)
	}
	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("decode command output: %v", err)
	}
	if payload["schema_version"] != "xnix.runtime.kde_test_launch_materialization_fanout.v1" ||
		payload["request_type"] != "kde-test-launch-materialization-fanout-preview" ||
		payload["mode"] != "test-only" ||
		payload["materialization_status"] != "blocked-plan-materialized" ||
		payload["materialization_scope"] != "test-only-review-plan" ||
		payload["surface_count"] != float64(4) ||
		payload["all_checks_passed"] != true ||
		payload["materialization_receipt_consumed"] != true ||
		payload["execution_session_fan_out_consumed"] != true ||
		payload["fan_out_writes_enabled"] != false ||
		payload["notification_sent"] != false {
		t.Fatalf("unexpected test launch materialization fan-out command payload: %s", output.String())
	}
	if strings.Contains(output.String(), stateRoot) {
		t.Fatalf("command output must not expose state-root path: %s", output.String())
	}
	surfaces := payload["surfaces"].([]any)
	seen := map[string]bool{}
	for _, entry := range surfaces {
		surface := entry.(map[string]any)
		seen[surface["id"].(string)] = true
		for _, key := range []string{"mutates_runtime", "starts_program", "delivers_notification", "activates_task_manager_entry", "enables_tray_bridge", "enables_center_actions", "exposes_state_root_path", "exposes_raw_command", "exposes_raw_executable", "exposes_backend_details"} {
			if surface[key] != false {
				t.Fatalf("surface gate %s must remain false: %s", key, output.String())
			}
		}
	}
	for _, id := range []string{"compatibility-center", "task-manager", "tray", "notification"} {
		if !seen[id] {
			t.Fatalf("missing KDE fan-out surface %s: %s", id, output.String())
		}
	}
	for _, key := range []string{"product_image_ready", "production_trust_satisfied", "runtime_write_gate_enabled", "launch_preflight_passed", "launch_authorized", "execution_approved", "process_start_authorized", "command_materialized", "executable_path_resolved", "backend_selected_for_launch", "backend_launch_enabled", "backend_process_started", "task_manager_entry_active", "live_tray_bridge_enabled", "notification_delivery_enabled", "compatibility_center_actions_enabled", "request_objects_created", "runtime_writes_enabled", "production_bus_ownership", "network_required", "host_root_modified", "privileged_container_required", "raw_command_exposed", "raw_executable_exposed", "backend_details_exposed"} {
		if payload[key] != false {
			t.Fatalf("unsafe gate %s must remain false: %s", key, output.String())
		}
	}
}

func TestKDETestLaunchMaterializationFanOutPreviewCommandRequiresAuthorization(t *testing.T) {
	base := []string{"kde-test-launch-materialization-fanout-preview", "--registry", "../../runtime/recipes/registry.json", "--app", "org.xnix.sample.notepad", "--state-root", t.TempDir()}
	for _, args := range [][]string{base, append(append([]string{}, base...), "--mode", "test-only"), append(append([]string{}, base...), "--mode", "test-only", "--authorize", "yes")} {
		if err := run(args, &bytes.Buffer{}); err == nil {
			t.Fatalf("expected command to require exact authorization: %v", args)
		}
	}
}
