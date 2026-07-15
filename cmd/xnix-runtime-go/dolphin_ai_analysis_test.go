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

func TestDolphinAIAnalysisPreviewCommandRendersPrivacySafePlan(t *testing.T) {
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

	var output bytes.Buffer
	err := run([]string{"dolphin-ai-analysis-preview", "--registry", registryPath, "file:///home/test/Documents/book.xls", "file:///home/test/Documents/tax.xls"}, &output)
	if err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["schema_version"] != "xnix.runtime.dolphin_ai_analysis.v1" ||
		payload["request_type"] != "dolphin-ai-analysis-preview" ||
		payload["source"] != "dolphin-ai-action" ||
		payload["analysis_surface"] != "Dolphin" {
		t.Fatalf("unexpected Dolphin AI schema: %#v", payload)
	}
	if payload["application_id"] != "org.example.ledger" ||
		payload["desktop_file"] != "xnix-org.example.ledger.desktop" ||
		payload["runtime_method"] != "GetAIDiagnosticInput" ||
		payload["analysis_task"] != "compatibility-file-review" {
		t.Fatalf("unexpected Dolphin AI identity: %#v", payload)
	}
	if payload["file_count"] != float64(2) ||
		payload["selected_extension"] != ".xls" ||
		payload["selection_mode"] != "extension-match" ||
		payload["selected_file_disclosure"] != "count-and-extension-only" {
		t.Fatalf("unexpected Dolphin AI selection: %#v", payload)
	}
	if payload["portal_required"] != true ||
		payload["portal_interface"] != "org.freedesktop.portal.FileChooser" ||
		payload["portal_method"] != "OpenFile" {
		t.Fatalf("unexpected Dolphin AI Portal metadata: %#v", payload)
	}
	if payload["runtime_owned"] != true ||
		payload["go_runtime_backed"] != true ||
		payload["kde_policy_owner"] != false ||
		payload["safe_for_ai_diagnostics"] != true ||
		payload["user_review_required"] != true ||
		payload["ai_provider_call_enabled"] != false ||
		payload["network_required"] != false ||
		payload["file_content_read"] != false ||
		payload["file_paths_exposed"] != false ||
		payload["request_object_created"] != false ||
		payload["permission_granted"] != false ||
		payload["backend_launch_enabled"] != false ||
		payload["host_root_modified"] != false ||
		payload["backend_details_exposed"] != false {
		t.Fatalf("unexpected Dolphin AI safety flags: %#v", payload)
	}
	rendered := strings.ToLower(output.String())
	for _, forbidden := range []string{"file://", "/home/test", "documents/book.xls", "documents/tax.xls"} {
		if strings.Contains(rendered, forbidden) {
			t.Fatalf("Dolphin AI analysis command exposed forbidden file detail %q: %s", forbidden, rendered)
		}
	}
}
