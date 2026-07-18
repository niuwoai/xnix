package main

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"
)

func TestKDETestLaunchMaterializationRecordCommandMaterializesPlanOnly(t *testing.T) {
	stateRoot := t.TempDir()
	var output bytes.Buffer
	err := run([]string{"kde-test-launch-materialization-record", "--registry", "../../runtime/recipes/registry.json", "--app", "org.xnix.sample.notepad", "--state-root", stateRoot, "--mode", "test-only", "--authorize", "authorize-restricted-test-preparation"}, &output)
	if err != nil {
		t.Fatalf("kde-test-launch-materialization-record returned error: %v", err)
	}
	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("decode command output: %v", err)
	}
	materialization := payload["materialization"].(map[string]any)
	if payload["schema_version"] != "xnix.runtime.kde_test_launch_materialization.v1" ||
		payload["record_type"] != "kde-test-launch-materialization-record" ||
		payload["mode"] != "test-only" ||
		payload["materialization_boundary_joined"] != true ||
		payload["plan_materialized"] != true ||
		payload["test_only"] != true ||
		payload["core_receipt_count"] != float64(11) ||
		payload["all_checks_passed"] != true ||
		materialization["status"] != "blocked-plan-materialized" ||
		materialization["materialization_scope"] != "test-only-review-plan" ||
		materialization["materialized_artifact_count"] != float64(4) ||
		materialization["blocked_by_count"] != float64(2) ||
		materialization["plan_materialized"] != true ||
		materialization["test_only"] != true {
		t.Fatalf("unexpected test launch materialization command payload: %s", output.String())
	}
	if strings.Contains(output.String(), stateRoot) {
		t.Fatalf("command output must not expose state-root path: %s", output.String())
	}
	for _, key := range []string{"product_image_ready", "production_trust_satisfied", "runtime_write_gate_enabled", "launch_preflight_passed", "launch_authorized", "execution_approved", "process_start_authorized", "command_materialized", "executable_path_resolved", "backend_selected_for_launch", "backend_launch_enabled", "backend_process_started", "production_bus_ownership", "network_required", "host_root_modified", "privileged_container_required", "raw_command_exposed", "raw_executable_exposed", "backend_details_exposed"} {
		if payload[key] != false {
			t.Fatalf("unsafe gate %s must remain false: %s", key, output.String())
		}
	}
	for _, key := range []string{"command_materialized", "executable_path_resolved", "backend_selected_for_launch", "backend_launch_enabled"} {
		if materialization[key] != false {
			t.Fatalf("unsafe materialization gate %s must remain false: %s", key, output.String())
		}
	}
}

func TestKDETestLaunchMaterializationRecordCommandRequiresAuthorization(t *testing.T) {
	base := []string{"kde-test-launch-materialization-record", "--registry", "../../runtime/recipes/registry.json", "--app", "org.xnix.sample.notepad", "--state-root", t.TempDir()}
	for _, args := range [][]string{base, append(append([]string{}, base...), "--mode", "test-only"), append(append([]string{}, base...), "--mode", "test-only", "--authorize", "yes")} {
		if err := run(args, &bytes.Buffer{}); err == nil {
			t.Fatalf("expected command to require exact authorization: %v", args)
		}
	}
}
