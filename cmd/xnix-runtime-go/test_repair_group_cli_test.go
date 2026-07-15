package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

func writeTestRepairGroupRegistry(t *testing.T) (string, string) {
	t.Helper()
	root := t.TempDir()
	recipeData := []byte(`{"id":"org.example.ledger","name":"Example Ledger","icon":"office-chart-area","mode":"automatic","supported_extensions":[".abc"]}`)
	sum := sha256.Sum256(recipeData)
	digest := hex.EncodeToString(sum[:])
	if err := os.WriteFile(filepath.Join(root, "org.example.ledger.json"), recipeData, 0o600); err != nil {
		t.Fatalf("WriteFile recipe returned error: %v", err)
	}
	registryPath := filepath.Join(root, "registry.json")
	registryData := []byte(`{"schema_version":1,"registry_name":"test-registry","recipes":[{"id":"org.example.ledger","path":"org.example.ledger.json","sha256":"` + digest + `","signature_status":"development-only"}]}`)
	if err := os.WriteFile(registryPath, registryData, 0o600); err != nil {
		t.Fatalf("WriteFile registry returned error: %v", err)
	}
	return registryPath, "org.example.ledger"
}

func TestRepairPlanPreviewCommandRendersIssueRule(t *testing.T) {
	registryPath, app := writeTestRepairGroupRegistry(t)

	var output bytes.Buffer
	if err := run([]string{"repair-plan-preview", "--registry", registryPath, "--app", app, "--issue", "engine-binding-pending"}, &output); err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["schema_version"] != "xnix.runtime.repair_plan.v1" ||
		payload["runtime_method"] != "GetRepairPlan" ||
		payload["severity"] != "warning" ||
		payload["snapshot_required"] != true ||
		payload["backend_launch_enabled"] != false ||
		payload["host_root_modified"] != false ||
		payload["backend_details_exposed"] != false {
		t.Fatalf("unexpected repair plan payload: %#v", payload)
	}

	// The command must require --issue.
	var missing bytes.Buffer
	if err := run([]string{"repair-plan-preview", "--registry", registryPath, "--app", app}, &missing); err == nil {
		t.Fatalf("repair-plan-preview must require --issue")
	}
}

func TestTestPlanPreviewCommandRendersSteps(t *testing.T) {
	registryPath, app := writeTestRepairGroupRegistry(t)

	var output bytes.Buffer
	if err := run([]string{"test-plan-preview", "--registry", registryPath, "--app", app}, &output); err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["schema_version"] != "xnix.runtime.test_plan.v1" ||
		payload["runtime_method"] != "GetTestPlan" ||
		payload["test_type"] != "preflight" ||
		payload["blocked"] != false ||
		payload["test_executed"] != false ||
		payload["backend_details_exposed"] != false {
		t.Fatalf("unexpected test plan payload: %#v", payload)
	}
	steps := payload["step_ids"].([]any)
	if len(steps) != 4 || steps[0] != "recipe-validation" {
		t.Fatalf("unexpected test plan steps: %#v", steps)
	}
}

func TestTestResultPreviewCommandRendersCounts(t *testing.T) {
	registryPath, app := writeTestRepairGroupRegistry(t)

	var output bytes.Buffer
	if err := run([]string{"test-result-preview", "--registry", registryPath, "--app", app, "--test-type", "smoke"}, &output); err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["schema_version"] != "xnix.runtime.test_result.v1" ||
		payload["runtime_method"] != "GetTestResult" ||
		payload["test_type"] != "smoke" ||
		payload["overall_status"] != "pending" ||
		payload["test_executed"] != false ||
		payload["backend_details_exposed"] != false {
		t.Fatalf("unexpected test result payload: %#v", payload)
	}
	counts := payload["counts"].(map[string]any)
	if counts["total"] != float64(4) || counts["passed"] != float64(1) || counts["pending"] != float64(3) {
		t.Fatalf("unexpected test result counts: %#v", counts)
	}
}
