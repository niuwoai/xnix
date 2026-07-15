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

func writeBackendGroupRegistry(t *testing.T) (string, string) {
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

func TestBackendCapabilityMatrixPreviewCommandRendersMatrix(t *testing.T) {
	var output bytes.Buffer
	if err := run([]string{"backend-capability-matrix-preview"}, &output); err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["schema_version"] != "xnix.runtime.backend_capability_matrix.v1" ||
		payload["runtime_method"] != "GetBackendCapabilityMatrix" ||
		payload["go_runtime_backed"] != true ||
		payload["selection_enabled"] != false ||
		payload["backend_launch_enabled"] != false ||
		payload["backend_details_exposed"] != false {
		t.Fatalf("unexpected backend capability matrix payload: %#v", payload)
	}
	if payload["profile_count"] != float64(2) || payload["capability_count"] != float64(7) {
		t.Fatalf("unexpected matrix counts: %#v", payload)
	}

	// The command must reject positional arguments.
	var rejected bytes.Buffer
	if err := run([]string{"backend-capability-matrix-preview", "extra"}, &rejected); err == nil {
		t.Fatalf("backend-capability-matrix-preview must reject arguments")
	}
}

func TestBackendLifecyclePreviewCommandRendersStages(t *testing.T) {
	registryPath, app := writeBackendGroupRegistry(t)

	var output bytes.Buffer
	if err := run([]string{"backend-lifecycle-preview", "--registry", registryPath, "--app", app}, &output); err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["schema_version"] != "xnix.runtime.backend_lifecycle.v1" ||
		payload["runtime_method"] != "GetBackendLifecycle" ||
		payload["lifecycle_state"] != "blocked" ||
		payload["overall_status"] != "not-ready" ||
		payload["launch_enabled"] != false ||
		payload["host_root_modified"] != false ||
		payload["backend_details_exposed"] != false {
		t.Fatalf("unexpected backend lifecycle payload: %#v", payload)
	}
	stages := payload["stage_ids"].([]any)
	if len(stages) != 5 || stages[0] != "recipe-loaded" {
		t.Fatalf("unexpected backend lifecycle stages: %#v", stages)
	}
}
