package main

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"
)

func TestKDERestrictedLaunchPreflightRecordCommandIsFailClosed(t *testing.T) {
	stateRoot := t.TempDir()
	var output bytes.Buffer
	err := run([]string{"kde-restricted-launch-preflight-record", "--registry", "../../runtime/recipes/registry.json", "--app", "org.xnix.sample.notepad", "--state-root", stateRoot, "--mode", "test-only", "--authorize", "authorize-restricted-test-preparation"}, &output)
	if err != nil {
		t.Fatalf("kde-restricted-launch-preflight-record returned error: %v", err)
	}
	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("decode command output: %v", err)
	}
	preflight := payload["preflight"].(map[string]any)
	if preflight["status"] != "blocked" || preflight["blocker_count"] != float64(2) || preflight["ready_for_packet_assembly"] != true || preflight["product_image_ready"] != false || preflight["launch_preflight_passed"] != false || payload["preflight_boundary_joined"] != true || payload["core_receipt_count"] != float64(10) || payload["all_checks_passed"] != true {
		t.Fatalf("unexpected restricted preflight command payload: %s", output.String())
	}
	if strings.Contains(output.String(), stateRoot) {
		t.Fatalf("command output must not expose state-root path: %s", output.String())
	}
	for _, key := range []string{"product_image_ready", "production_trust_satisfied", "runtime_write_gate_enabled", "launch_preflight_passed", "launch_authorized", "execution_approved", "process_start_authorized", "command_materialized", "executable_path_resolved", "backend_selected_for_launch", "backend_launch_enabled", "backend_process_started", "production_bus_ownership", "network_required", "host_root_modified", "privileged_container_required", "raw_command_exposed", "backend_details_exposed"} {
		if payload[key] != false {
			t.Fatalf("unsafe gate %s must remain false: %s", key, output.String())
		}
	}
}

func TestKDERestrictedLaunchPreflightRecordCommandRequiresAuthorization(t *testing.T) {
	base := []string{"kde-restricted-launch-preflight-record", "--registry", "../../runtime/recipes/registry.json", "--app", "org.xnix.sample.notepad", "--state-root", t.TempDir()}
	for _, args := range [][]string{base, append(append([]string{}, base...), "--mode", "test-only"), append(append([]string{}, base...), "--mode", "test-only", "--authorize", "yes")} {
		if err := run(args, &bytes.Buffer{}); err == nil {
			t.Fatalf("expected command to require exact authorization: %v", args)
		}
	}
}
