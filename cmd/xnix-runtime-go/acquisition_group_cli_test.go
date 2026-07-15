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

func writeAcquisitionGroupRegistry(t *testing.T) (string, string) {
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

func TestPackageSourcePreviewCommandRendersRuntimeOwnedSource(t *testing.T) {
	registryPath, app := writeAcquisitionGroupRegistry(t)

	var output bytes.Buffer
	if err := run([]string{"package-source-preview", "--registry", registryPath, "--app", app}, &output); err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["schema_version"] != "xnix.runtime.package_source.v1" ||
		payload["request_type"] != "package-source-preview" ||
		payload["source_type"] != "compatibility-package-source" ||
		payload["runtime_method"] != "GetCompatibilityPackageSource" ||
		payload["go_runtime_backed"] != true ||
		payload["package_source_ready"] != false ||
		payload["install_enabled"] != false ||
		payload["host_root_modified"] != false ||
		payload["backend_details_exposed"] != false {
		t.Fatalf("unexpected package source payload: %#v", payload)
	}
	if payload["selected_strategy"] != "automatic-managed" {
		t.Fatalf("unexpected selected strategy: %#v", payload["selected_strategy"])
	}
	channels := payload["source_channel_ids"].([]any)
	if len(channels) != 3 || channels[0] != "os-managed-compatibility-packages" {
		t.Fatalf("unexpected source channels: %#v", channels)
	}
}

func TestAcquisitionPreflightPreviewCommandRendersChecks(t *testing.T) {
	registryPath, app := writeAcquisitionGroupRegistry(t)

	var output bytes.Buffer
	if err := run([]string{"acquisition-preflight-preview", "--registry", registryPath, "--app", app}, &output); err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["schema_version"] != "xnix.runtime.acquisition_preflight.v1" ||
		payload["runtime_method"] != "GetCompatibilityAcquisitionPreflight" ||
		payload["acquisition_ready"] != false ||
		payload["download_enabled"] != false ||
		payload["network_request_created"] != false ||
		payload["backend_details_exposed"] != false {
		t.Fatalf("unexpected acquisition preflight payload: %#v", payload)
	}
	checks := payload["check_ids"].([]any)
	if len(checks) != 5 || checks[0] != "package-source-ready" {
		t.Fatalf("unexpected acquisition checks: %#v", checks)
	}
}

func TestArtifactManifestPreviewCommandRendersGroups(t *testing.T) {
	registryPath, app := writeAcquisitionGroupRegistry(t)

	var output bytes.Buffer
	if err := run([]string{"artifact-manifest-preview", "--registry", registryPath, "--app", app}, &output); err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["schema_version"] != "xnix.runtime.artifact_manifest.v1" ||
		payload["manifest_type"] != "compatibility-artifact-manifest" ||
		payload["runtime_method"] != "GetCompatibilityArtifactManifest" ||
		payload["manifest_ready"] != false ||
		payload["signature_verified"] != false ||
		payload["host_root_modified"] != false ||
		payload["backend_details_exposed"] != false {
		t.Fatalf("unexpected artifact manifest payload: %#v", payload)
	}
	groups := payload["artifact_group_ids"].([]any)
	if len(groups) != 3 || groups[0] != "runtime-launch-metadata" {
		t.Fatalf("unexpected artifact groups: %#v", groups)
	}
}
