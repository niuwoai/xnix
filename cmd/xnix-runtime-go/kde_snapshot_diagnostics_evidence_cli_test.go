package main

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"
)

func TestKDESnapshotDiagnosticsEvidenceRecordCommandConvergesEvidence(t *testing.T) {
	stateRoot := t.TempDir()
	var output bytes.Buffer
	err := run([]string{
		"kde-snapshot-diagnostics-evidence-record",
		"--registry", "../../runtime/recipes/registry.json",
		"--app", "org.xnix.sample.notepad",
		"--state-root", stateRoot,
		"--mode", "test-only",
	}, &output)
	if err != nil {
		t.Fatalf("kde-snapshot-diagnostics-evidence-record returned error: %v", err)
	}
	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("decode command output: %v", err)
	}
	diagnostic := payload["diagnostic"].(map[string]any)
	snapshot := payload["snapshot"].(map[string]any)
	lifecycle := payload["lifecycle"].(map[string]any)
	execution := payload["execution"].(map[string]any)
	if diagnostic["overall"] != "pass" || diagnostic["redaction_verified"] != true ||
		snapshot["verified"] != true || snapshot["baseline_present"] != true ||
		lifecycle["state"] != "ready" || execution["state"] != "blocked" ||
		payload["core_receipt_count"] != float64(7) || payload["all_checks_passed"] != true {
		t.Fatalf("unexpected snapshot diagnostics command payload: %s", output.String())
	}
	if strings.Contains(output.String(), stateRoot) {
		t.Fatalf("command output must not expose state-root path: %s", output.String())
	}
	for _, key := range []string{"snapshot_restore_enabled", "snapshot_deletion_enabled", "diagnostic_execution_enabled", "ai_provider_call_enabled", "repair_execution_enabled", "real_portal_call_enabled", "host_permission_changed", "execution_approved", "launch_enabled", "execution_started", "backend_process_started", "production_bus_ownership", "host_root_modified", "privileged_container_required", "backend_details_exposed", "file_contents_exposed"} {
		if payload[key] != false {
			t.Fatalf("unsafe gate %s must remain false: %s", key, output.String())
		}
	}
}

func TestKDESnapshotDiagnosticsEvidenceRecordCommandRequiresTestOnlyBoundary(t *testing.T) {
	base := []string{"kde-snapshot-diagnostics-evidence-record", "--registry", "../../runtime/recipes/registry.json", "--app", "org.xnix.sample.notepad"}
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
