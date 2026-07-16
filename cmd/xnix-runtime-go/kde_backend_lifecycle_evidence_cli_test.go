package main

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"
)

func TestKDEBackendLifecycleEvidenceRecordCommandJoinsInventory(t *testing.T) {
	stateRoot := t.TempDir()
	var output bytes.Buffer
	err := run([]string{
		"kde-backend-lifecycle-evidence-record",
		"--registry", "../../runtime/recipes/registry.json",
		"--app", "org.xnix.sample.notepad",
		"--state-root", stateRoot,
		"--mode", "test-only",
	}, &output)
	if err != nil {
		t.Fatalf("kde-backend-lifecycle-evidence-record returned error: %v", err)
	}
	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("decode command output: %v", err)
	}
	manager := payload["backend_manager"].(map[string]any)
	lifecycle := payload["lifecycle"].(map[string]any)
	execution := payload["execution"].(map[string]any)
	if manager["inventory_read_back"] != true || manager["managed_backend_count"] != float64(3) || manager["all_backends_planned"] != true || manager["all_backend_processes_stopped"] != true ||
		lifecycle["state"] != "ready" || execution["state"] != "blocked" || payload["backend_state_joined"] != true || payload["core_receipt_count"] != float64(8) || payload["all_checks_passed"] != true {
		t.Fatalf("unexpected backend lifecycle command payload: %s", output.String())
	}
	if strings.Contains(output.String(), stateRoot) {
		t.Fatalf("command output must not expose state-root path: %s", output.String())
	}
	for _, key := range []string{"backend_kinds_exposed_to_kde", "backend_install_enabled", "backend_download_enabled", "backend_launch_enabled", "backend_process_started", "vm_process_started", "raw_command_exposed", "profile_path_exposed", "real_portal_call_enabled", "snapshot_restore_enabled", "diagnostic_execution_enabled", "ai_provider_call_enabled", "repair_execution_enabled", "execution_approved", "launch_enabled", "execution_started", "production_bus_ownership", "network_required", "host_root_modified", "privileged_container_required", "backend_details_exposed", "secrets_exposed"} {
		if payload[key] != false {
			t.Fatalf("unsafe gate %s must remain false: %s", key, output.String())
		}
	}
}

func TestKDEBackendLifecycleEvidenceRecordCommandRequiresTestOnlyBoundary(t *testing.T) {
	base := []string{"kde-backend-lifecycle-evidence-record", "--registry", "../../runtime/recipes/registry.json", "--app", "org.xnix.sample.notepad"}
	for _, args := range [][]string{
		base,
		append(append([]string{}, base...), "--state-root", t.TempDir()),
		append(append([]string{}, base...), "--state-root", t.TempDir(), "--mode", "production"),
		append(append([]string{}, base...), "--state-root", t.TempDir(), "--mode", "test-only", "extra"),
	} {
		if err := run(args, &bytes.Buffer{}); err == nil {
			t.Fatalf("expected command to reject unsafe arguments: %v", args)
		}
	}
}
