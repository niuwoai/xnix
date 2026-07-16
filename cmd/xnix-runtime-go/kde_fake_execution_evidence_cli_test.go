package main

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"
)

func TestKDEFakeExecutionEvidenceRecordCommandWritesControlledEvidence(t *testing.T) {
	stateRoot := t.TempDir()
	var output bytes.Buffer
	err := run([]string{
		"kde-fake-execution-evidence-record",
		"--registry", "../../runtime/recipes/registry.json",
		"--app", "org.xnix.sample.notepad",
		"--state-root", stateRoot,
		"--mode", "test-only",
	}, &output)
	if err != nil {
		t.Fatalf("kde-fake-execution-evidence-record returned error: %v", err)
	}
	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("decode command output: %v", err)
	}
	application := payload["application"].(map[string]any)
	lifecycle := payload["lifecycle"].(map[string]any)
	executionRecord := payload["execution"].(map[string]any)
	if application["id"] != "org.xnix.sample.notepad" ||
		lifecycle["state"] != "staged" ||
		executionRecord["state"] != "blocked" ||
		payload["mode"] != "test-only" ||
		payload["state_root_record_count"] != float64(3) ||
		payload["all_checks_passed"] != true ||
		payload["fake_execution_recorded"] != true ||
		payload["state_root_writes_enabled"] != true {
		t.Fatalf("unexpected command payload: %s", output.String())
	}
	if strings.Contains(output.String(), stateRoot) {
		t.Fatalf("command output must not expose state-root path: %s", output.String())
	}
	for _, key := range []string{"launch_allowed", "launch_enabled", "execution_started", "backend_process_started", "real_portal_call_enabled", "production_bus_ownership", "host_root_modified", "privileged_container_required", "backend_details_exposed"} {
		if payload[key] != false {
			t.Fatalf("unsafe gate %s must remain false: %s", key, output.String())
		}
	}
}

func TestKDEFakeExecutionEvidenceRecordCommandRequiresExplicitTestBoundary(t *testing.T) {
	base := []string{"kde-fake-execution-evidence-record", "--registry", "../../runtime/recipes/registry.json", "--app", "org.xnix.sample.notepad"}
	cases := [][]string{
		base,
		append(append([]string{}, base...), "--state-root", t.TempDir()),
		append(append([]string{}, base...), "--state-root", t.TempDir(), "--mode", "production"),
		append(append([]string{}, base...), "--state-root", t.TempDir(), "--mode", "test-only", "extra"),
	}
	for _, args := range cases {
		if err := run(args, &bytes.Buffer{}); err == nil {
			t.Fatalf("expected command to reject unsafe arguments: %v", args)
		}
	}
}
