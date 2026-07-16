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

func TestApplicationUpgradeImpactPreviewCommandRendersCandidateRegistryImpact(t *testing.T) {
	currentRegistry := writeApplicationUpgradeImpactRegistry(t, applicationUpgradeImpactRecipeFixture{
		ID:        "org.example.ledger",
		Name:      "Example Ledger",
		Version:   "1.0.0",
		Icon:      "office-chart-area",
		Mode:      "automatic",
		Extension: ".abc",
		Signature: "signed",
	}, false)
	candidateRegistry := writeApplicationUpgradeImpactRegistry(t, applicationUpgradeImpactRecipeFixture{
		ID:        "org.example.ledger",
		Name:      "Example Ledger",
		Version:   "1.1.0",
		Icon:      "office-chart-area",
		Mode:      "automatic",
		Extension: ".abc",
		Signature: "signed",
	}, false)

	var output bytes.Buffer
	err := run([]string{
		"application-upgrade-impact-preview",
		"--registry", currentRegistry,
		"--candidate-registry", candidateRegistry,
		"--app", "org.example.ledger",
		"--mode", "development",
	}, &output)
	if err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["schema_version"] != "xnix.runtime.application_upgrade_impact.v1" ||
		payload["request_type"] != "application-upgrade-impact-preview" ||
		payload["impact_type"] != "review-only-application-upgrade-impact" ||
		payload["runtime_method"] != "GetApplicationUpgradeImpact" ||
		payload["read_method"] != "GetApplicationUpgradeImpactPreview" ||
		payload["environment"] != "development" ||
		payload["application_id"] != "org.example.ledger" ||
		payload["candidate_application_id"] != "org.example.ledger" ||
		payload["current_version"] != "1.0.0" ||
		payload["candidate_version"] != "1.1.0" ||
		payload["candidate_relation"] != "candidate-newer" ||
		payload["section_count"] != float64(8) ||
		payload["overall_state"] != "missing-evidence" ||
		payload["ready_for_review"] != true ||
		payload["runtime_owned"] != true ||
		payload["go_runtime_backed"] != true ||
		payload["kde_policy_owner"] != false ||
		payload["review_only"] != true ||
		payload["recipe_written"] != false ||
		payload["registry_migrated"] != false ||
		payload["request_objects_created"] != false ||
		payload["artifacts_staged"] != false ||
		payload["artifacts_downloaded"] != false ||
		payload["network_request_created"] != false ||
		payload["host_package_manager_invoked"] != false ||
		payload["desktop_activation_started"] != false ||
		payload["desktop_files_written"] != false ||
		payload["backend_process_started"] != false ||
		payload["launch_enabled"] != false ||
		payload["execution_started"] != false ||
		payload["host_root_modified"] != false ||
		payload["state_root_path_exposed"] != false ||
		payload["raw_command_exposed"] != false ||
		payload["backend_details_exposed"] != false {
		t.Fatalf("unexpected application upgrade impact payload: %#v", payload)
	}
	sections := payload["sections"].([]any)
	first := sections[0].(map[string]any)
	counts := payload["counts"].(map[string]any)
	if first["id"] != "recipe-metadata" ||
		first["state"] != "changed" ||
		first["changed"] != true ||
		counts["missing_evidence"] == float64(0) ||
		counts["changed"] == float64(0) {
		t.Fatalf("unexpected upgrade impact sections: first=%#v counts=%#v", first, counts)
	}
	if strings.Contains(output.String(), currentRegistry) || strings.Contains(output.String(), candidateRegistry) {
		t.Fatalf("application upgrade impact must not expose local registry paths: %s", output.String())
	}
	assertApplicationUpgradeImpactCLISafe(t, output.String())
}

func TestApplicationUpgradeImpactPreviewCommandSupportsCandidateAppOverride(t *testing.T) {
	currentRegistry := writeApplicationUpgradeImpactRegistry(t, applicationUpgradeImpactRecipeFixture{
		ID:        "org.example.ledger",
		Name:      "Example Ledger",
		Version:   "1.0.0",
		Icon:      "office-chart-area",
		Mode:      "automatic",
		Extension: ".abc",
		Signature: "signed",
	}, false)
	candidateRegistry := writeApplicationUpgradeImpactRegistry(t, applicationUpgradeImpactRecipeFixture{
		ID:        "org.example.viewer",
		Name:      "Example Viewer",
		Version:   "1.1.0",
		Icon:      "accessories-text-editor",
		Mode:      "automatic",
		Extension: ".xyz",
		Signature: "signed",
	}, false)

	var output bytes.Buffer
	err := run([]string{
		"application-upgrade-impact-preview",
		"--registry", currentRegistry,
		"--candidate-registry", candidateRegistry,
		"--app", "org.example.ledger",
		"--candidate-app", "org.example.viewer",
		"--mode", "development",
	}, &output)
	if err != nil {
		t.Fatalf("run returned error: %v", err)
	}
	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["overall_state"] != "blocked" || payload["candidate_application_id"] != "org.example.viewer" {
		t.Fatalf("unexpected candidate-app override payload: %#v", payload)
	}
}

func TestApplicationUpgradeImpactPreviewCommandRejectsMissingInputs(t *testing.T) {
	var output bytes.Buffer
	if err := run([]string{"application-upgrade-impact-preview"}, &output); err == nil {
		t.Fatalf("application-upgrade-impact-preview must require registry inputs")
	}
}

func TestApplicationUpgradeImpactPreviewCommandRejectsDigestDrift(t *testing.T) {
	currentRegistry := writeApplicationUpgradeImpactRegistry(t, applicationUpgradeImpactRecipeFixture{
		ID:        "org.example.ledger",
		Name:      "Example Ledger",
		Version:   "1.0.0",
		Icon:      "office-chart-area",
		Mode:      "automatic",
		Extension: ".abc",
		Signature: "signed",
	}, false)
	candidateRegistry := writeApplicationUpgradeImpactRegistry(t, applicationUpgradeImpactRecipeFixture{
		ID:        "org.example.ledger",
		Name:      "Example Ledger",
		Version:   "1.1.0",
		Icon:      "office-chart-area",
		Mode:      "automatic",
		Extension: ".abc",
		Signature: "signed",
	}, true)

	var output bytes.Buffer
	err := run([]string{
		"application-upgrade-impact-preview",
		"--registry", currentRegistry,
		"--candidate-registry", candidateRegistry,
		"--app", "org.example.ledger",
	}, &output)
	if err == nil {
		t.Fatalf("application-upgrade-impact-preview accepted candidate digest drift")
	}
	if !strings.Contains(err.Error(), "recipe digest mismatch") {
		t.Fatalf("expected digest mismatch error, got %v", err)
	}
}

type applicationUpgradeImpactRecipeFixture struct {
	ID        string
	Name      string
	Version   string
	Icon      string
	Mode      string
	Extension string
	Signature string
}

func writeApplicationUpgradeImpactRegistry(t *testing.T, fixture applicationUpgradeImpactRecipeFixture, drift bool) string {
	t.Helper()
	root := t.TempDir()
	recipe := struct {
		ID                  string   `json:"id"`
		Name                string   `json:"name"`
		Version             string   `json:"version"`
		Icon                string   `json:"icon"`
		Mode                string   `json:"mode"`
		SupportedExtensions []string `json:"supported_extensions"`
	}{
		ID:                  fixture.ID,
		Name:                fixture.Name,
		Version:             fixture.Version,
		Icon:                fixture.Icon,
		Mode:                fixture.Mode,
		SupportedExtensions: []string{fixture.Extension},
	}
	recipeData, err := json.Marshal(recipe)
	if err != nil {
		t.Fatalf("Marshal recipe returned error: %v", err)
	}
	recipeFile := fixture.ID + ".json"
	if err := os.WriteFile(filepath.Join(root, recipeFile), recipeData, 0o600); err != nil {
		t.Fatalf("WriteFile recipe returned error: %v", err)
	}
	sumSource := recipeData
	if drift {
		sumSource = []byte("{}")
	}
	sum := sha256.Sum256(sumSource)
	registry := struct {
		SchemaVersion int    `json:"schema_version"`
		RegistryName  string `json:"registry_name"`
		Recipes       []struct {
			ID              string `json:"id"`
			Path            string `json:"path"`
			SHA256          string `json:"sha256"`
			SignatureStatus string `json:"signature_status"`
		} `json:"recipes"`
	}{
		SchemaVersion: 1,
		RegistryName:  "test-registry",
		Recipes: []struct {
			ID              string `json:"id"`
			Path            string `json:"path"`
			SHA256          string `json:"sha256"`
			SignatureStatus string `json:"signature_status"`
		}{
			{
				ID:              fixture.ID,
				Path:            recipeFile,
				SHA256:          hex.EncodeToString(sum[:]),
				SignatureStatus: fixture.Signature,
			},
		},
	}
	registryData, err := json.Marshal(registry)
	if err != nil {
		t.Fatalf("Marshal registry returned error: %v", err)
	}
	registryPath := filepath.Join(root, "registry.json")
	if err := os.WriteFile(registryPath, registryData, 0o600); err != nil {
		t.Fatalf("WriteFile registry returned error: %v", err)
	}
	return registryPath
}

func assertApplicationUpgradeImpactCLISafe(t *testing.T, text string) {
	t.Helper()
	serialized := strings.ToLower(text)
	for _, forbidden := range []string{"prefix", ".exe", "program files", "qemu-system", "proton", "wine ", "wine/", ".wine", "virtual machine", "/home", "/users", "/private", "/var", "/opt", "/tmp", "file://"} {
		if strings.Contains(serialized, forbidden) {
			t.Fatalf("application upgrade impact CLI exposed forbidden term %q: %s", forbidden, text)
		}
	}
}
