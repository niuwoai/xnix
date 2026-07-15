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

func TestDiagnosticsPreviewCommandRendersGoDiagnostics(t *testing.T) {
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
	err := run([]string{"diagnostics-preview", "--registry", registryPath, "--app", "org.example.ledger"}, &output)
	if err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["schema_version"] != "xnix.runtime.diagnostics.v1" ||
		payload["request_type"] != "diagnostics-preview" ||
		payload["diagnostics_type"] != "runtime-diagnostics" ||
		payload["source"] != "registry+go-runtime-previews" ||
		payload["desktop"] != "KDE Plasma" ||
		payload["runtime_method"] != "GetDiagnostics" ||
		payload["read_method"] != "GetDiagnosticsPreview" {
		t.Fatalf("unexpected diagnostics CLI schema: %#v", payload)
	}
	application := payload["application"].(map[string]any)
	if application["id"] != "org.example.ledger" ||
		application["name"] != "Example Ledger" ||
		application["desktop_file"] != "xnix-org.example.ledger.desktop" ||
		application["registry_name"] != "test-registry" ||
		application["digest_verified"] != true ||
		application["signature_status"] != "development-only" {
		t.Fatalf("unexpected diagnostics application: %#v", application)
	}
	counts := payload["counts"].(map[string]any)
	if counts["total"] != float64(6) ||
		counts["passed"] != float64(3) ||
		counts["pending"] != float64(1) ||
		counts["blocked"] != float64(2) {
		t.Fatalf("unexpected diagnostics counts: %#v", counts)
	}
	execution := payload["execution_readiness"].(map[string]any)
	if execution["request_type"] != "execution-readiness-preview" ||
		execution["overall_status"] != "not-ready" ||
		execution["execution_state"] != "blocked" ||
		execution["launch_enabled"] != false ||
		execution["safe_for_ai_diagnostics"] != true ||
		execution["backend_details_exposed"] != false {
		t.Fatalf("unexpected execution diagnostics: %#v", execution)
	}
	actionQueue := payload["action_queue"].(map[string]any)
	if actionQueue["request_type"] != "kde-action-queue-preview" ||
		actionQueue["action_count"] != float64(7) ||
		actionQueue["pending_action_count"] != float64(7) ||
		actionQueue["action_queue_created"] != true ||
		actionQueue["action_queue_persisted"] != false ||
		actionQueue["execution_started"] != false {
		t.Fatalf("unexpected action queue diagnostics: %#v", actionQueue)
	}
	ai := payload["ai"].(map[string]any)
	if ai["runtime_method"] != "GetAIDiagnosticInput" ||
		ai["analysis_task"] != "compatibility-status-review" ||
		ai["safe_for_ai_diagnostics"] != true ||
		ai["ai_provider_call_enabled"] != false ||
		ai["network_required"] != false ||
		ai["file_content_read"] != false ||
		ai["file_paths_exposed"] != false {
		t.Fatalf("unexpected AI diagnostics: %#v", ai)
	}
	if payload["runtime_owned"] != true ||
		payload["go_runtime_backed"] != true ||
		payload["kde_policy_owner"] != false ||
		payload["safe_for_ai_diagnostics"] != true ||
		payload["ai_provider_call_enabled"] != false ||
		payload["file_content_read"] != false ||
		payload["file_paths_exposed"] != false ||
		payload["request_object_created"] != false ||
		payload["permission_granted"] != false ||
		payload["launch_enabled"] != false ||
		payload["execution_started"] != false ||
		payload["repair_execution_enabled"] != false ||
		payload["settings_persisted"] != false ||
		payload["host_root_modified"] != false ||
		payload["network_required"] != false ||
		payload["backend_details_exposed"] != false {
		t.Fatalf("unexpected diagnostics safety flags: %#v", payload)
	}
	serialized := strings.ToLower(output.String())
	for _, forbidden := range []string{"prefix", ".exe", "program files", "qemu-system", "proton", "wine "} {
		if strings.Contains(serialized, forbidden) {
			t.Fatalf("diagnostics CLI exposes forbidden term %q: %s", forbidden, output.String())
		}
	}
}
