package main

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
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
		readiness["recipe_trust_diagnostics_ready"] != true ||
		readiness["artifact_manifest_ready"] != false ||
		readiness["state_root_allocated"] != false ||
		len(readiness["recipe_trust_blocking_reasons"].([]any)) != 0 {
		t.Fatalf("unexpected compatibility install readiness: %#v", readiness)
	}
	phases := payload["phase_ids"].([]any)
	if len(phases) != 8 || phases[0] != "resolve-artifact-manifest" || phases[2] != "consume-artifact-stage-receipt" {
		t.Fatalf("unexpected compatibility install phases: %#v", phases)
	}
	assertInstallPreviewPayloadSafe(t, output.String())
}

func TestCompatibilityInstallPreviewCommandConsumesArtifactReceipt(t *testing.T) {
	registryPath, app := writeAcquisitionGroupRegistry(t)
	manifestPath, fixtureRoot, cacheRoot := writeArtifactStageFixture(t)
	var stageOutput bytes.Buffer
	if err := run([]string{"artifact-stage-record", "--manifest", manifestPath, "--fixture-root", fixtureRoot, "--cache-root", cacheRoot}, &stageOutput); err != nil {
		t.Fatalf("artifact-stage-record returned error: %v", err)
	}
	receiptPath := filepath.Join(cacheRoot, "artifact-ledger", "receipts", "org.example.ledger.json")

	var output bytes.Buffer
	if err := run([]string{"compatibility-install-preview", "--registry", registryPath, "--app", app, "--mode", "development", "--artifact-receipt", receiptPath}, &output); err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	payload := decodeInstallPreviewPayload(t, output.Bytes())
	if payload["source"] != "artifact-manifest-preview+acquisition-preflight-preview+package-source-preview+state-root-preview+recipe-install-gate+artifact-stage-receipt" {
		t.Fatalf("install preview must expose artifact receipt source: %#v", payload)
	}
	readiness := payload["readiness"].(map[string]any)
	if readiness["artifact_stage_receipt_ready"] != true ||
		readiness["artifact_stage_digest_verified"] != true ||
		readiness["required_artifacts_staged"] != true ||
		len(readiness["artifact_stage_blocking_reasons"].([]any)) != 0 {
		t.Fatalf("install preview must consume valid artifact receipt: %#v", readiness)
	}
	receipt := payload["artifact_stage_receipt"].(map[string]any)
	if receipt["relative_path"] != "artifact-ledger/receipts/org.example.ledger.json" ||
		receipt["required_artifacts_staged"] != true ||
		receipt["staged_key_count"] != float64(2) {
		t.Fatalf("unexpected install artifact receipt payload: %#v", receipt)
	}
	assertInstallPreviewPayloadSafe(t, output.String())
	if strings.Contains(output.String(), cacheRoot) || strings.Contains(output.String(), fixtureRoot) {
		t.Fatalf("install preview must not expose artifact roots: %s", output.String())
	}
}

func TestCompatibilityInstallPreviewCommandFailsClosedForTamperedArtifactReceipt(t *testing.T) {
	registryPath, app := writeAcquisitionGroupRegistry(t)
	manifestPath, fixtureRoot, cacheRoot := writeArtifactStageFixture(t)
	var stageOutput bytes.Buffer
	if err := run([]string{"artifact-stage-record", "--manifest", manifestPath, "--fixture-root", fixtureRoot, "--cache-root", cacheRoot}, &stageOutput); err != nil {
		t.Fatalf("artifact-stage-record returned error: %v", err)
	}
	receiptPath := filepath.Join(cacheRoot, "artifact-ledger", "receipts", "org.example.ledger.json")
	data, err := os.ReadFile(receiptPath)
	if err != nil {
		t.Fatalf("ReadFile receipt: %v", err)
	}
	data = bytes.Replace(data, []byte(`"host_root_modified": false`), []byte(`"host_root_modified": true`), 1)
	if err := os.WriteFile(receiptPath, data, 0o600); err != nil {
		t.Fatalf("WriteFile receipt: %v", err)
	}

	var output bytes.Buffer
	if err := run([]string{"compatibility-install-preview", "--registry", registryPath, "--app", app, "--mode", "development", "--artifact-receipt", receiptPath}, &output); err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	payload := decodeInstallPreviewPayload(t, output.Bytes())
	readiness := payload["readiness"].(map[string]any)
	if readiness["artifact_stage_receipt_ready"] != false ||
		readiness["artifact_stage_digest_verified"] != false ||
		len(readiness["artifact_stage_blocking_reasons"].([]any)) == 0 {
		t.Fatalf("tampered artifact receipt must fail closed: %#v", readiness)
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
		readiness["recipe_trust_diagnostics_ready"] != true ||
		installGate["decision"] != "block" {
		t.Fatalf("unexpected production compatibility install payload: %#v", payload)
	}
	reasons := readiness["recipe_trust_blocking_reasons"].([]any)
	if len(reasons) != 2 ||
		reasons[0] != "production signed recipe validation is not enabled" ||
		reasons[1] != "registry contains development-only recipes" {
		t.Fatalf("unexpected production compatibility install blocking reasons: %#v", reasons)
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
