package main

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"
)

func TestCompatibilityInstallPreviewCommandRendersDevelopmentPlan(t *testing.T) {
	registryPath, app := writeAcquisitionGroupRegistry(t)

	var output bytes.Buffer
	if err := run([]string{"compatibility-install-preview", "--registry", registryPath, "--app", app, "--mode", "development"}, &output); err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	payload := decodeInstallPreviewPayload(t, output.Bytes())
	if payload["schema_version"] != "xnix.runtime.compatibility_install_plan.v1" ||
		payload["request_type"] != "compatibility-install-preview" ||
		payload["plan_type"] != "compatibility-install-plan" ||
		payload["runtime_method"] != "GetCompatibilityInstallPlan" ||
		payload["read_method"] != "GetCompatibilityInstallPlanPreview" ||
		payload["environment"] != "development" ||
		payload["go_runtime_backed"] != true ||
		payload["install_ready"] != false ||
		payload["desktop_activation_ready"] != false ||
		payload["download_enabled"] != false ||
		payload["install_enabled"] != false ||
		payload["host_root_modified"] != false ||
		payload["backend_details_exposed"] != false {
		t.Fatalf("unexpected compatibility install payload: %#v", payload)
	}
	readiness := payload["readiness"].(map[string]any)
	if readiness["recipe_install_allowed"] != true ||
		readiness["recipe_install_decision"] != "allow" ||
		readiness["artifact_manifest_ready"] != false ||
		readiness["state_root_allocated"] != false {
		t.Fatalf("unexpected compatibility install readiness: %#v", readiness)
	}
	phases := payload["phase_ids"].([]any)
	if len(phases) != 7 || phases[0] != "resolve-artifact-manifest" {
		t.Fatalf("unexpected compatibility install phases: %#v", phases)
	}
	assertInstallPreviewPayloadSafe(t, output.String())
}

func TestCompatibilityInstallPreviewCommandBlocksProductionDevelopmentRecipe(t *testing.T) {
	registryPath, app := writeAcquisitionGroupRegistry(t)

	var output bytes.Buffer
	if err := run([]string{"compatibility-install-preview", "--registry", registryPath, "--app", app, "--mode", "production"}, &output); err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	payload := decodeInstallPreviewPayload(t, output.Bytes())
	readiness := payload["readiness"].(map[string]any)
	installGate := payload["install_gate"].(map[string]any)
	if payload["environment"] != "production" ||
		readiness["recipe_install_allowed"] != false ||
		readiness["recipe_install_decision"] != "block" ||
		installGate["decision"] != "block" {
		t.Fatalf("unexpected production compatibility install payload: %#v", payload)
	}
	assertInstallPreviewPayloadSafe(t, output.String())
}

func decodeInstallPreviewPayload(t *testing.T, data []byte) map[string]any {
	t.Helper()
	var payload map[string]any
	if err := json.Unmarshal(data, &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	return payload
}

func assertInstallPreviewPayloadSafe(t *testing.T, text string) {
	t.Helper()
	serialized := strings.ToLower(text)
	for _, forbidden := range []string{"prefix", ".exe", "program files", "qemu-system", "proton", "wine ", "virtual machine"} {
		if strings.Contains(serialized, forbidden) {
			t.Fatalf("compatibility install CLI exposed forbidden term %q: %s", forbidden, text)
		}
	}
}
