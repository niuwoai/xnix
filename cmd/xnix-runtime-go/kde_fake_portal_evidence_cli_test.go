package main

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"
)

func TestKDEFakePortalEvidenceRecordCommandJoinsPortalReceipt(t *testing.T) {
	stateRoot := t.TempDir()
	var output bytes.Buffer
	err := run([]string{
		"kde-fake-portal-evidence-record",
		"--registry", "../../runtime/recipes/registry.json",
		"--app", "org.xnix.sample.notepad",
		"--state-root", stateRoot,
		"--mode", "test-only",
	}, &output)
	if err != nil {
		t.Fatalf("kde-fake-portal-evidence-record returned error: %v", err)
	}
	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("decode command output: %v", err)
	}
	portalEvidence := payload["portal"].(map[string]any)
	executionEvidence := payload["execution"].(map[string]any)
	if portalEvidence["completed_state"] != "completed" ||
		portalEvidence["permission_state"] != "granted" ||
		executionEvidence["state"] != "blocked" ||
		payload["portal_gate_changed_only"] != true ||
		payload["portal_evidence_recorded"] != true ||
		payload["state_root_record_count"] != float64(4) ||
		payload["all_checks_passed"] != true {
		t.Fatalf("unexpected fake Portal command payload: %s", output.String())
	}
	if strings.Contains(output.String(), stateRoot) {
		t.Fatalf("command output must not expose state-root path: %s", output.String())
	}
	for _, key := range []string{"real_portal_call_enabled", "host_permission_changed", "execution_approved", "launch_enabled", "execution_started", "backend_process_started", "production_bus_ownership", "host_root_modified", "privileged_container_required", "backend_details_exposed"} {
		if payload[key] != false {
			t.Fatalf("unsafe gate %s must remain false: %s", key, output.String())
		}
	}
}

func TestKDEFakePortalEvidenceRecordCommandRequiresTestOnlyBoundary(t *testing.T) {
	base := []string{"kde-fake-portal-evidence-record", "--registry", "../../runtime/recipes/registry.json", "--app", "org.xnix.sample.notepad"}
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
