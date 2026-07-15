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

func TestDiagnosticRunHistoryCommandSummarizesSafeStateRootRecords(t *testing.T) {
	fixturePath, stateRoot := writeDiagnosticRecordFixture(t)
	if err := run([]string{"diagnostic-run-record", "--state-root", stateRoot, "--app", "org.example.ledger", "--run-id", "smoke-001", "--fixture", fixturePath}, &bytes.Buffer{}); err != nil {
		t.Fatalf("record first run: %v", err)
	}
	if err := run([]string{"diagnostic-run-record", "--state-root", stateRoot, "--app", "org.example.ledger", "--run-id", "smoke-002", "--fixture", fixturePath}, &bytes.Buffer{}); err != nil {
		t.Fatalf("record second run: %v", err)
	}

	var output bytes.Buffer
	err := run([]string{"diagnostic-run-history", "--state-root", stateRoot, "--app", "org.example.ledger"}, &output)
	if err != nil {
		t.Fatalf("history returned error: %v", err)
	}
	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["schema_version"] != "xnix.runtime.diagnostic_run_history.v1" ||
		payload["record_type"] != "diagnostic-run-history" ||
		payload["source"] != "go-runtime-state-root-diagnostic-run-history" ||
		payload["application_id"] != "org.example.ledger" ||
		payload["runtime_owned"] != true ||
		payload["go_runtime_backed"] != true ||
		payload["kde_policy_owner"] != false ||
		payload["state_root_path_exposed"] != false ||
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
		t.Fatalf("unexpected diagnostic history payload: %#v", payload)
	}
	counts := payload["counts"].(map[string]any)
	if counts["total"] != float64(2) || counts["failed"] != float64(2) {
		t.Fatalf("unexpected diagnostic history counts: %#v", counts)
	}
	records := payload["records"].([]any)
	if len(records) != 2 {
		t.Fatalf("unexpected history record count: %#v", records)
	}
	latest := payload["latest"].(map[string]any)
	if latest["run_id"] != "smoke-002" ||
		latest["repair_issue"] != "engine-binding-pending" ||
		latest["snapshot_required"] != true ||
		latest["backend_started"] != false ||
		latest["ai_provider_called"] != false ||
		latest["repair_executed"] != false ||
		latest["host_root_modified"] != false ||
		latest["backend_details_exposed"] != false ||
		latest["file_contents_included"] != false {
		t.Fatalf("unexpected latest diagnostic history record: %#v", latest)
	}
	if strings.Contains(output.String(), stateRoot) || strings.Contains(output.String(), fixturePath) {
		t.Fatalf("diagnostic history output must not expose local roots: %s", output.String())
	}
}

func TestDiagnosticRunHistoryCommandRequiresStateRoot(t *testing.T) {
	var output bytes.Buffer
	if err := run([]string{"diagnostic-run-history", "--app", "org.example.ledger"}, &output); err == nil {
		t.Fatalf("diagnostic-run-history must require --state-root")
	}
}

func TestDiagnosticHistoryPreviewCommandRendersCompatibilityCenterProjection(t *testing.T) {
	fixturePath, stateRoot := writeDiagnosticRecordFixture(t)
	if err := run([]string{"diagnostic-run-record", "--state-root", stateRoot, "--app", "org.example.ledger", "--run-id", "smoke-001", "--fixture", fixturePath}, &bytes.Buffer{}); err != nil {
		t.Fatalf("record run: %v", err)
	}

	var output bytes.Buffer
	err := run([]string{"diagnostic-history-preview", "--state-root", stateRoot, "--app", "org.example.ledger"}, &output)
	if err != nil {
		t.Fatalf("preview returned error: %v", err)
	}
	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["schema_version"] != "xnix.runtime.diagnostic_history_preview.v1" ||
		payload["request_type"] != "diagnostic-history-preview" ||
		payload["source"] != "go-runtime-state-root-diagnostic-run-history+kde-read-model" ||
		payload["desktop"] != "KDE Plasma" ||
		payload["runtime_method"] != "GetDiagnostics" ||
		payload["read_method"] != "GetDiagnosticHistoryPreview" ||
		payload["application_id"] != "org.example.ledger" ||
		payload["runtime_owned"] != true ||
		payload["go_runtime_backed"] != true ||
		payload["kde_policy_owner"] != false ||
		payload["user_visible"] != true ||
		payload["compatibility_center_card"] != true ||
		payload["safe_for_ai_diagnostics"] != true ||
		payload["state_root_path_exposed"] != false ||
		payload["file_content_read"] != false ||
		payload["file_paths_exposed"] != false ||
		payload["ai_provider_call_enabled"] != false ||
		payload["request_object_created"] != false ||
		payload["permission_granted"] != false ||
		payload["launch_enabled"] != false ||
		payload["execution_started"] != false ||
		payload["repair_execution_enabled"] != false ||
		payload["settings_persisted"] != false ||
		payload["host_root_modified"] != false ||
		payload["network_required"] != false ||
		payload["privileged_container_required"] != false ||
		payload["backend_details_exposed"] != false {
		t.Fatalf("unexpected diagnostic history preview payload: %#v", payload)
	}
	center := payload["compatibility_center"].(map[string]any)
	if center["history_state"] != "needs-review" ||
		center["latest_run_id"] != "smoke-001" ||
		center["latest_overall"] != "fail" ||
		center["latest_repair_issue"] != "engine-binding-pending" ||
		center["total_run_count"] != float64(1) ||
		center["failing_run_count"] != float64(1) ||
		center["action_execution_enabled"] != false ||
		center["repair_execution_enabled"] != false ||
		center["backend_launch_enabled"] != false {
		t.Fatalf("unexpected Compatibility Center projection: %#v", center)
	}
	if strings.Contains(output.String(), stateRoot) || strings.Contains(output.String(), fixturePath) {
		t.Fatalf("diagnostic history preview must not expose local roots: %s", output.String())
	}
}
