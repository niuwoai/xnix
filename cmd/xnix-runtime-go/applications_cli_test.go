package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestApplicationsPreviewCommandRendersGoCatalog(t *testing.T) {
	registryPath := writeApplicationPreviewRegistry(t)

	var output bytes.Buffer
	err := run([]string{"applications-preview", "--registry", registryPath}, &output)
	if err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["schema_version"] != "xnix.runtime.applications.v1" ||
		payload["request_type"] != "applications-preview" ||
		payload["catalog_type"] != "runtime-application-catalog" ||
		payload["source"] != "registry+go-runtime-desktop-identity" ||
		payload["desktop"] != "KDE Plasma" ||
		payload["runtime_method"] != "ListApplications" ||
		payload["read_method"] != "ListApplicationsPreview" {
		t.Fatalf("unexpected applications CLI schema: %#v", payload)
	}
	if payload["registry_name"] != "test-registry" ||
		payload["recipe_digest_verified"] != true ||
		payload["recipe_signature_status"] != "development-only" ||
		payload["application_count"] != float64(1) {
		t.Fatalf("unexpected applications summary: %#v", payload)
	}
	applications := payload["applications"].([]any)
	application := applications[0].(map[string]any)
	if application["id"] != "org.example.ledger" ||
		application["name"] != "Example Ledger" ||
		application["runtime_mode"] != "Automatic" ||
		application["desktop_file"] != "xnix-org.example.ledger.desktop" ||
		application["launcher_action"] != "runtime-launch" ||
		application["user_visible"] != true ||
		application["standard_desktop_entry"] != true ||
		application["accepts_file_uris"] != true ||
		application["backend_launch_enabled"] != false ||
		application["backend_details_exposed"] != false ||
		application["raw_windows_executable_exposed"] != false {
		t.Fatalf("unexpected applications entry: %#v", application)
	}
	if payload["runtime_owned"] != true ||
		payload["go_runtime_backed"] != true ||
		payload["kde_policy_owner"] != false ||
		payload["standard_desktop_entries"] != true ||
		payload["recipe_registry_verified"] != true ||
		payload["launch_enabled"] != false ||
		payload["desktop_files_written"] != false ||
		payload["host_root_modified"] != false ||
		payload["backend_details_exposed"] != false ||
		payload["raw_windows_executable_exposed"] != false {
		t.Fatalf("unexpected applications safety flags: %#v", payload)
	}
	assertNoForbiddenApplicationCLIOutput(t, output.String())
}

func TestApplicationPreviewCommandRendersGoApplication(t *testing.T) {
	registryPath := writeApplicationPreviewRegistry(t)

	var output bytes.Buffer
	err := run([]string{"application-preview", "--registry", registryPath, "--app", "org.example.ledger"}, &output)
	if err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["schema_version"] != "xnix.runtime.application.v1" ||
		payload["request_type"] != "application-preview" ||
		payload["catalog_type"] != "runtime-application" ||
		payload["source"] != "registry+go-runtime-desktop-identity" ||
		payload["desktop"] != "KDE Plasma" ||
		payload["runtime_method"] != "GetApplication" ||
		payload["read_method"] != "GetApplicationPreview" {
		t.Fatalf("unexpected application CLI schema: %#v", payload)
	}
	application := payload["application"].(map[string]any)
	if application["id"] != "org.example.ledger" ||
		application["name"] != "Example Ledger" ||
		application["desktop_file"] != "xnix-org.example.ledger.desktop" ||
		application["registry_name"] != "test-registry" ||
		application["recipe_digest_verified"] != true ||
		application["recipe_signature_status"] != "development-only" {
		t.Fatalf("unexpected application payload: %#v", application)
	}
	if payload["runtime_owned"] != true ||
		payload["go_runtime_backed"] != true ||
		payload["kde_policy_owner"] != false ||
		payload["standard_desktop_entry"] != true ||
		payload["recipe_registry_verified"] != true ||
		payload["launch_enabled"] != false ||
		payload["desktop_file_written"] != false ||
		payload["host_root_modified"] != false ||
		payload["backend_details_exposed"] != false ||
		payload["raw_windows_executable_exposed"] != false {
		t.Fatalf("unexpected application safety flags: %#v", payload)
	}
	assertNoForbiddenApplicationCLIOutput(t, output.String())
}

func writeApplicationPreviewRegistry(t *testing.T) string {
	t.Helper()
	root := t.TempDir()
	recipeData := []byte(`{"id":"org.example.ledger","name":"Example Ledger","icon":"office-chart-area","mode":"automatic","supported_extensions":[".abc",".xls"]}`)
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
	return registryPath
}

func assertNoForbiddenApplicationCLIOutput(t *testing.T, output string) {
	t.Helper()
	serialized := strings.ToLower(output)
	for _, forbidden := range []string{"prefix", ".exe", "program files", "qemu-system", "proton", "wine "} {
		if strings.Contains(serialized, forbidden) {
			t.Fatalf("application CLI exposes forbidden term %q: %s", forbidden, output)
		}
	}
}
