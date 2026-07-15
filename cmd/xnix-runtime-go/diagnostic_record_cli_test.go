package main

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func writeDiagnosticRecordFixture(t *testing.T) (string, string) {
	t.Helper()
	root := t.TempDir()
	fixturePath := filepath.Join(root, "fixture.json")
	fixture := []byte(`{
  "test_type": "smoke",
  "signals": [
    {
      "id": "recipe-validation",
      "category": "recipe",
      "outcome": "pass",
      "summary": "Recipe metadata is valid."
    },
    {
      "id": "runtime-launch-binding",
      "category": "backend",
      "outcome": "fail",
      "summary": "Launch binding is not ready."
    }
  ]
}`)
	if err := os.WriteFile(fixturePath, fixture, 0o600); err != nil {
		t.Fatalf("WriteFile fixture returned error: %v", err)
	}
	return fixturePath, filepath.Join(root, "state")
}

func TestDiagnosticRunRecordCommandWritesSafeStateRootReceipt(t *testing.T) {
	fixturePath, stateRoot := writeDiagnosticRecordFixture(t)

	var output bytes.Buffer
	err := run([]string{"diagnostic-run-record", "--state-root", stateRoot, "--app", "org.example.ledger", "--run-id", "smoke-001", "--fixture", fixturePath}, &output)
	if err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["schema_version"] != "xnix.runtime.diagnostic_run_record.v1" ||
		payload["record_type"] != "diagnostic-run-record" ||
		payload["source"] != "go-runtime-state-root-diagnostic-run-record" ||
		payload["application_id"] != "org.example.ledger" ||
		payload["run_id"] != "smoke-001" ||
		payload["relative_path"] != "diagnostics-ledger/runs/smoke-001.json" ||
		payload["runtime_owned"] != true ||
		payload["go_runtime_backed"] != true ||
		payload["kde_policy_owner"] != false ||
		payload["state_root_path_exposed"] != false ||
		payload["fixture_path_exposed"] != false ||
		payload["backend_started"] != false ||
		payload["ai_provider_called"] != false ||
		payload["real_ai_provider_enabled"] != false ||
		payload["auto_repair_allowed"] != false ||
		payload["repair_executed"] != false ||
		payload["host_root_modified"] != false ||
		payload["network_required"] != false ||
		payload["privileged_container_required"] != false ||
		payload["backend_details_exposed"] != false ||
		payload["file_contents_included"] != false {
		t.Fatalf("unexpected diagnostic run record payload: %#v", payload)
	}
	if payload["sha256"] == "" {
		t.Fatalf("diagnostic run record must include a digest")
	}
	result := payload["result"].(map[string]any)
	if result["overall"] != "fail" || result["backend_started"] != false || result["host_root_modified"] != false {
		t.Fatalf("unexpected diagnostic result: %#v", result)
	}
	input := payload["diagnostic_input"].(map[string]any)
	if input["network_required"] != false || input["file_contents_included"] != false {
		t.Fatalf("unexpected diagnostic input: %#v", input)
	}
	recommendation := payload["repair_recommendation"].(map[string]any)
	if recommendation["issue"] != "engine-binding-pending" ||
		recommendation["snapshot_required"] != true ||
		recommendation["auto_apply_allowed"] != false {
		t.Fatalf("unexpected repair recommendation: %#v", recommendation)
	}
	if strings.Contains(output.String(), stateRoot) || strings.Contains(output.String(), fixturePath) {
		t.Fatalf("diagnostic run output must not expose local roots: %s", output.String())
	}
	if _, err := os.Stat(filepath.Join(stateRoot, "diagnostics-ledger", "runs", "smoke-001.json")); err != nil {
		t.Fatalf("diagnostic record was not written under state root: %v", err)
	}
}

func TestDiagnosticRunRecordCommandRequiresInputs(t *testing.T) {
	fixturePath, stateRoot := writeDiagnosticRecordFixture(t)
	cases := [][]string{
		{"diagnostic-run-record", "--app", "org.example.ledger", "--run-id", "smoke-001", "--fixture", fixturePath},
		{"diagnostic-run-record", "--state-root", stateRoot, "--run-id", "smoke-001", "--fixture", fixturePath},
		{"diagnostic-run-record", "--state-root", stateRoot, "--app", "org.example.ledger", "--fixture", fixturePath},
		{"diagnostic-run-record", "--state-root", stateRoot, "--app", "org.example.ledger", "--run-id", "smoke-001"},
	}
	for _, args := range cases {
		var output bytes.Buffer
		if err := run(args, &output); err == nil {
			t.Fatalf("expected command to reject missing input: %#v", args)
		}
	}
}
