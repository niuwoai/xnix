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

func TestMultiApplicationInstallQueuePreviewCommandRendersRegistryQueue(t *testing.T) {
	registryPath, apps := writeMultiApplicationInstallQueueRegistry(t)

	var output bytes.Buffer
	err := run([]string{"multi-application-install-queue-preview", "--registry", registryPath, "--mode", "development"}, &output)
	if err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["schema_version"] != "xnix.runtime.multi_application_install_queue.v1" ||
		payload["request_type"] != "multi-application-install-queue-preview" ||
		payload["queue_type"] != "review-only-multi-application-install-queue" ||
		payload["runtime_method"] != "GetMultiApplicationInstallQueue" ||
		payload["read_method"] != "GetMultiApplicationInstallQueuePreview" ||
		payload["environment"] != "development" ||
		payload["application_count"] != float64(2) ||
		payload["queue_status"] != "missing-evidence" ||
		payload["runtime_owned"] != true ||
		payload["go_runtime_backed"] != true ||
		payload["kde_policy_owner"] != false ||
		payload["user_visible"] != true ||
		payload["review_only"] != true ||
		payload["queue_persisted"] != false ||
		payload["request_objects_created"] != false ||
		payload["artifacts_staged"] != false ||
		payload["artifacts_downloaded"] != false ||
		payload["network_request_created"] != false ||
		payload["host_package_manager_invoked"] != false ||
		payload["desktop_activation_started"] != false ||
		payload["install_started"] != false ||
		payload["backend_process_started"] != false ||
		payload["launch_enabled"] != false ||
		payload["execution_started"] != false ||
		payload["host_root_modified"] != false ||
		payload["state_root_path_exposed"] != false ||
		payload["raw_command_exposed"] != false ||
		payload["backend_details_exposed"] != false {
		t.Fatalf("unexpected multi-application install queue payload: %#v", payload)
	}
	ids := payload["application_ids"].([]any)
	items := payload["applications"].([]any)
	counts := payload["counts"].(map[string]any)
	if len(ids) != 2 || ids[0] != apps[0] || ids[1] != apps[1] || len(items) != 2 ||
		counts["total"] != float64(2) ||
		counts["missing_evidence"] != float64(2) {
		t.Fatalf("unexpected queue apps: ids=%#v items=%#v counts=%#v", ids, items, counts)
	}
	first := items[0].(map[string]any)
	if first["application_id"] != apps[0] ||
		first["queue_state"] != "missing-evidence" ||
		first["request_object_created"] != false ||
		first["artifact_staged"] != false ||
		first["backend_process_started"] != false ||
		first["host_root_modified"] != false ||
		first["raw_command_exposed"] != false ||
		first["backend_details_exposed"] != false {
		t.Fatalf("unexpected first queue item: %#v", first)
	}
	if strings.Contains(output.String(), registryPath) {
		t.Fatalf("multi-application install queue must not expose local registry path: %s", output.String())
	}
	assertMultiApplicationInstallQueueCLISafe(t, output.String())
}

func TestMultiApplicationInstallQueuePreviewCommandFiltersApps(t *testing.T) {
	registryPath, apps := writeMultiApplicationInstallQueueRegistry(t)

	var output bytes.Buffer
	err := run([]string{
		"multi-application-install-queue-preview",
		"--registry", registryPath,
		"--app", apps[1],
		"--mode", "development",
	}, &output)
	if err != nil {
		t.Fatalf("run returned error: %v", err)
	}
	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	ids := payload["application_ids"].([]any)
	if payload["application_count"] != float64(1) || len(ids) != 1 || ids[0] != apps[1] {
		t.Fatalf("unexpected filtered app ids: %#v", payload)
	}
}

func TestMultiApplicationInstallQueuePreviewCommandRejectsMissingRegistry(t *testing.T) {
	var output bytes.Buffer
	if err := run([]string{"multi-application-install-queue-preview"}, &output); err == nil {
		t.Fatalf("multi-application-install-queue-preview must require --registry")
	}
}

func writeMultiApplicationInstallQueueRegistry(t *testing.T) (string, []string) {
	t.Helper()
	root := t.TempDir()
	recipes := []struct {
		id        string
		name      string
		icon      string
		mode      string
		extension string
		signature string
		file      string
	}{
		{"org.example.ledger", "Example Ledger", "office-chart-area", "automatic", ".abc", "development-only", "org.example.ledger.json"},
		{"org.example.viewer", "Example Viewer", "accessories-text-editor", "wine", ".xyz", "development-only", "org.example.viewer.json"},
	}
	type registryEntry struct {
		ID              string `json:"id"`
		Path            string `json:"path"`
		SHA256          string `json:"sha256"`
		SignatureStatus string `json:"signature_status"`
	}
	var entries []registryEntry
	var ids []string
	for _, recipe := range recipes {
		data := []byte(`{"id":"` + recipe.id + `","name":"` + recipe.name + `","icon":"` + recipe.icon + `","mode":"` + recipe.mode + `","supported_extensions":["` + recipe.extension + `"]}`)
		sum := sha256.Sum256(data)
		if err := os.WriteFile(filepath.Join(root, recipe.file), data, 0o600); err != nil {
			t.Fatalf("WriteFile recipe returned error: %v", err)
		}
		entries = append(entries, registryEntry{
			ID:              recipe.id,
			Path:            recipe.file,
			SHA256:          hex.EncodeToString(sum[:]),
			SignatureStatus: recipe.signature,
		})
		ids = append(ids, recipe.id)
	}
	registry := struct {
		SchemaVersion int             `json:"schema_version"`
		RegistryName  string          `json:"registry_name"`
		Recipes       []registryEntry `json:"recipes"`
	}{
		SchemaVersion: 1,
		RegistryName:  "test-registry",
		Recipes:       entries,
	}
	data, err := json.Marshal(registry)
	if err != nil {
		t.Fatalf("Marshal registry returned error: %v", err)
	}
	registryPath := filepath.Join(root, "registry.json")
	if err := os.WriteFile(registryPath, data, 0o600); err != nil {
		t.Fatalf("WriteFile registry returned error: %v", err)
	}
	return registryPath, ids
}

func assertMultiApplicationInstallQueueCLISafe(t *testing.T, text string) {
	t.Helper()
	serialized := strings.ToLower(text)
	for _, forbidden := range []string{"prefix", ".exe", "program files", "qemu-system", "proton", "wine ", "wine/", ".wine", "virtual machine", "/home", "/users", "/private", "/var", "/opt", "/tmp", "file://"} {
		if strings.Contains(serialized, forbidden) {
			t.Fatalf("multi-application install queue CLI exposed forbidden term %q: %s", forbidden, text)
		}
	}
}
