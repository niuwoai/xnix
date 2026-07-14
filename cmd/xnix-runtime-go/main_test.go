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

func TestDesktopIdentityPlanCommandLoadsRegistryApplication(t *testing.T) {
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

	var output bytes.Buffer
	err := run([]string{"desktop-identity-plan", "--registry", registryPath, "--app", "org.example.ledger"}, &output)
	if err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["application_id"] != "org.example.ledger" {
		t.Fatalf("application_id = %#v", payload["application_id"])
	}
	if payload["recipe_source"] != "registry" || payload["registry_name"] != "test-registry" || payload["recipe_digest_verified"] != true {
		t.Fatalf("unexpected registry provenance: %#v", payload)
	}
}

func TestDesktopEntryPreviewCommandRendersManagedLauncher(t *testing.T) {
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

	var output bytes.Buffer
	err := run([]string{"desktop-entry-preview", "--registry", registryPath, "--app", "org.example.ledger"}, &output)
	if err != nil {
		t.Fatalf("run returned error: %v", err)
	}
	entry := output.String()
	required := []string{
		"[Desktop Entry]\n",
		"Name=Example Ledger\n",
		"Exec=xnix-compat-launch --app org.example.ledger %U\n",
		"MimeType=application/x-xnix-abc;\n",
	}
	for _, fragment := range required {
		if !bytes.Contains(output.Bytes(), []byte(fragment)) {
			t.Fatalf("desktop entry missing %q in:\n%s", fragment, entry)
		}
	}
}

func TestMIMEAppsPreviewCommandRendersAssociations(t *testing.T) {
	root := t.TempDir()
	recipeData := []byte(`{"id":"org.example.ledger","name":"Example Ledger","icon":"office-chart-area","mode":"automatic","supported_extensions":[".abc",".log"]}`)
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

	var output bytes.Buffer
	err := run([]string{"mimeapps-preview", "--registry", registryPath, "--app", "org.example.ledger"}, &output)
	if err != nil {
		t.Fatalf("run returned error: %v", err)
	}
	mimeapps := output.String()
	required := []string{
		"[Default Applications]\n",
		"application/x-xnix-abc=xnix-org.example.ledger.desktop\n",
		"application/x-xnix-log=xnix-org.example.ledger.desktop\n",
		"[Added Associations]\n",
		"application/x-xnix-abc=xnix-org.example.ledger.desktop;\n",
		"application/x-xnix-log=xnix-org.example.ledger.desktop;\n",
	}
	for _, fragment := range required {
		if !bytes.Contains(output.Bytes(), []byte(fragment)) {
			t.Fatalf("MIME apps preview missing %q in:\n%s", fragment, mimeapps)
		}
	}
}

func TestWindowIdentityPreviewCommandRendersKDEWindowIdentity(t *testing.T) {
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

	var output bytes.Buffer
	err := run([]string{"window-identity-preview", "--registry", registryPath, "--app", "org.example.ledger"}, &output)
	if err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["schema_version"] != "xnix.runtime.window_identity.v1" {
		t.Fatalf("schema_version = %#v", payload["schema_version"])
	}
	if payload["desktop_file"] != "xnix-org.example.ledger.desktop" || payload["launcher_url"] != "applications:xnix-org.example.ledger.desktop" {
		t.Fatalf("unexpected desktop identity: %#v", payload)
	}
	taskManager := payload["task_manager"].(map[string]any)
	if taskManager["grouping_key"] != "org.example.ledger" || taskManager["pinning_allowed"] != true ||
		taskManager["restore_allowed"] != true || taskManager["skip_taskbar"] != false || taskManager["show_in_switcher"] != true {
		t.Fatalf("unexpected task manager hints: %#v", taskManager)
	}
	kwin := payload["kwin"].(map[string]any)
	if kwin["script_role"] != "identity-and-layout" || kwin["desktop_file"] != "xnix-org.example.ledger.desktop" ||
		kwin["window_manager_policy_only"] != true || kwin["runtime_owns_backend_policy"] != true {
		t.Fatalf("unexpected KWin hints: %#v", kwin)
	}
}
