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

func TestRunPlanPreviewCommandRendersGoRunPlan(t *testing.T) {
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
	err := run([]string{"run-plan-preview", "--registry", registryPath, "--app", "org.example.ledger"}, &output)
	if err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["schema_version"] != "xnix.runtime.run_plan.v1" ||
		payload["request_type"] != "run-plan-preview" ||
		payload["plan_type"] != "compatibility-run" ||
		payload["source"] != "registry+go-runtime-run-plan" ||
		payload["desktop"] != "KDE Plasma" ||
		payload["runtime_method"] != "GetRunPlan" ||
		payload["read_method"] != "GetRunPlanPreview" {
		t.Fatalf("unexpected run plan CLI schema: %#v", payload)
	}
	application := payload["application"].(map[string]any)
	if application["id"] != "org.example.ledger" ||
		application["name"] != "Example Ledger" ||
		application["desktop_file"] != "xnix-org.example.ledger.desktop" ||
		application["runtime_mode"] != "automatic" {
		t.Fatalf("unexpected run plan application: %#v", application)
	}
	execution := payload["execution"].(map[string]any)
	engine := execution["selected_engine"].(map[string]any)
	profile := execution["selected_profile"].(map[string]any)
	binding := execution["backend_binding"].(map[string]any)
	if execution["strategy"] != "automatic-managed" ||
		execution["selection_source"] != "recipe" ||
		engine["id"] != "automatic" ||
		engine["ready"] != false ||
		engine["launch_enabled"] != false ||
		profile["id"] != "local-compatibility" ||
		profile["ready"] != false ||
		binding["ready"] != false ||
		binding["launch_enabled"] != false ||
		execution["backend_details_exposed"] != false ||
		execution["launch_enabled"] != false ||
		execution["execution_request_created"] != false ||
		execution["execution_started"] != false {
		t.Fatalf("unexpected run plan execution: %#v", execution)
	}
	preflight := payload["preflight"].(map[string]any)
	if preflight["portal_policy_required"] != true ||
		preflight["snapshot_before_risky_change"] != true ||
		preflight["diagnostics_required"] != true ||
		preflight["backend_binding_required"] != true ||
		preflight["runtime_write_gate_required"] != true {
		t.Fatalf("unexpected run plan preflight: %#v", preflight)
	}
	if payload["runtime_owned"] != true ||
		payload["go_runtime_backed"] != true ||
		payload["kde_policy_owner"] != false ||
		payload["user_visible"] != true ||
		payload["launch_enabled"] != false ||
		payload["execution_request_created"] != false ||
		payload["execution_started"] != false ||
		payload["backend_binding_ready"] != false ||
		payload["request_object_created"] != false ||
		payload["permission_granted"] != false ||
		payload["host_root_modified"] != false ||
		payload["network_required"] != false ||
		payload["backend_details_exposed"] != false ||
		payload["raw_command_exposed"] != false {
		t.Fatalf("unexpected run plan safety flags: %#v", payload)
	}

	serialized := strings.ToLower(output.String())
	for _, forbidden := range []string{"prefix", ".exe", "program files", "qemu-system", "proton", "wine ", "wine/", ".wine", "virtual machine"} {
		if strings.Contains(serialized, forbidden) {
			t.Fatalf("run plan CLI exposes forbidden term %q: %s", forbidden, output.String())
		}
	}
}
