package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func TestRecipeConflictAuditPreviewCLI(t *testing.T) {
	root := t.TempDir()
	body := `{"id":"org.example.clean","name":"Clean","version":"1.2.0","icon":"x","mode":"automatic","supported_extensions":[".abc"]}`
	if err := os.WriteFile(filepath.Join(root, "clean.json"), []byte(body), 0o600); err != nil {
		t.Fatalf("write recipe: %v", err)
	}
	sum := sha256.Sum256([]byte(body))
	registry := fmt.Sprintf(`{"schema_version":1,"registry_name":"clean","recipes":[{"id":"org.example.clean","path":"clean.json","sha256":%q,"signature_status":"signed"}]}`, hex.EncodeToString(sum[:]))
	registryPath := filepath.Join(root, "registry.json")
	if err := os.WriteFile(registryPath, []byte(registry), 0o600); err != nil {
		t.Fatalf("write registry: %v", err)
	}

	var output bytes.Buffer
	if err := run([]string{"recipe-conflict-audit-preview", "--registry", registryPath}, &output); err != nil {
		t.Fatalf("run recipe conflict audit: %v", err)
	}
	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("parse output: %v\n%s", err, output.String())
	}
	if payload["schema_version"] != "xnix.runtime.recipe_conflict_audit.v1" ||
		payload["request_type"] != "recipe-conflict-audit-preview" ||
		payload["overall_state"] != "clean" {
		t.Fatalf("unexpected payload identity: %+v", payload)
	}
	for _, key := range []string{"recipe_writes_enabled", "registry_migration_enabled", "artifact_staging_enabled", "network_required", "package_manager_invoked", "backend_launch_enabled", "host_root_modified"} {
		if payload[key] != false {
			t.Fatalf("preview must keep %q disabled: %+v", key, payload)
		}
	}
}

func TestRecipeConflictAuditPreviewCLIRequiresRegistry(t *testing.T) {
	tests := [][]string{
		{"recipe-conflict-audit-preview"},
		{"recipe-conflict-audit-preview", "--registry", filepath.Join(t.TempDir(), "missing.json")},
	}
	for _, args := range tests {
		var output bytes.Buffer
		if err := run(args, &output); err == nil {
			t.Fatalf("expected CLI args to fail: %+v", args)
		}
	}
}
