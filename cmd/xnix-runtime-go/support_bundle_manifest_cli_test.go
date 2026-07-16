package main

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestSupportBundleManifestPreviewCommandRendersRedactedManifest(t *testing.T) {
	registryPath, app := writeTestRepairGroupRegistry(t)
	fixturePath, stateRoot := writeDiagnosticRecordFixture(t)
	runtimeRoot := writeSupportBundleCLIRuntimeRoot(t)
	if err := run([]string{"diagnostic-run-record", "--state-root", stateRoot, "--app", app, "--run-id", "smoke-001", "--fixture", fixturePath}, &bytes.Buffer{}); err != nil {
		t.Fatalf("record diagnostic run: %v", err)
	}

	var output bytes.Buffer
	err := run([]string{
		"support-bundle-manifest-preview",
		"--registry", registryPath,
		"--app", app,
		"--state-root", stateRoot,
		"--runtime-root", runtimeRoot,
		"--issue", "engine-binding-pending",
		"--test-type", "smoke",
	}, &output)
	if err != nil {
		t.Fatalf("support bundle manifest returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["version"] != "0.2.303" ||
		payload["schema_version"] != "xnix.runtime.support_bundle_manifest.v1" ||
		payload["request_type"] != "support-bundle-manifest-preview" ||
		payload["manifest_type"] != "redacted-offline-support-bundle-manifest" ||
		payload["runtime_method"] != "GetSupportBundleManifest" ||
		payload["read_method"] != "GetSupportBundleManifestPreview" ||
		payload["runtime_owned"] != true ||
		payload["go_runtime_backed"] != true ||
		payload["kde_policy_owner"] != false ||
		payload["user_visible"] != true ||
		payload["offline_only"] != true ||
		payload["archive_created"] != false ||
		payload["file_content_read"] != false ||
		payload["file_paths_exposed"] != false ||
		payload["ai_provider_called"] != false ||
		payload["ai_provider_call_enabled"] != false ||
		payload["auto_repair_requested"] != false ||
		payload["auto_repair_executed"] != false ||
		payload["backend_process_started"] != false ||
		payload["host_root_modified"] != false ||
		payload["state_root_path_exposed"] != false ||
		payload["raw_executable_exposed"] != false ||
		payload["raw_command_exposed"] != false ||
		payload["backend_details_exposed"] != false ||
		payload["network_required"] != false ||
		payload["privileged_container_required"] != false {
		t.Fatalf("unexpected support bundle manifest payload: %#v", payload)
	}
	sections := payload["sections"].([]any)
	runs := payload["diagnostic_run_summaries"].([]any)
	failing := payload["failing_signal_ids"].([]any)
	categories := payload["repair_recommendation_categories"].([]any)
	omitted := payload["omitted_evidence"].(map[string]any)
	if len(sections) != 7 ||
		len(runs) != 1 ||
		len(failing) != 1 ||
		failing[0] != "runtime-launch-binding" ||
		len(categories) != 3 ||
		omitted["file_contents"] == float64(0) ||
		omitted["host_paths"] == float64(0) ||
		omitted["environment_variables"] == float64(0) ||
		omitted["token_shaped_values"] == float64(0) ||
		omitted["command_shaped_values"] == float64(0) ||
		omitted["usernames"] == float64(0) {
		t.Fatalf("unexpected support bundle manifest evidence: sections=%#v runs=%#v failing=%#v categories=%#v omitted=%#v", sections, runs, failing, categories, omitted)
	}
	if strings.Contains(output.String(), stateRoot) ||
		strings.Contains(output.String(), fixturePath) ||
		strings.Contains(output.String(), runtimeRoot) ||
		strings.Contains(output.String(), registryPath) {
		t.Fatalf("support bundle manifest must not expose local paths: %s", output.String())
	}
	assertSupportBundleCLISafe(t, output.String())
}

func TestSupportBundleManifestPreviewCommandWorksWithoutStateRoot(t *testing.T) {
	registryPath, app := writeTestRepairGroupRegistry(t)
	runtimeRoot := writeSupportBundleCLIRuntimeRoot(t)
	missingStateRoot := filepath.Join(t.TempDir(), "missing-state-root")

	var output bytes.Buffer
	err := run([]string{
		"support-bundle-manifest-preview",
		"--registry", registryPath,
		"--app", app,
		"--runtime-root", runtimeRoot,
	}, &output)
	if err != nil {
		t.Fatalf("support bundle manifest returned error: %v", err)
	}
	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["diagnostic_run_count"] != float64(0) ||
		payload["archive_created"] != false ||
		payload["file_content_read"] != false ||
		payload["host_root_modified"] != false {
		t.Fatalf("unexpected empty support manifest payload: %#v", payload)
	}
	if _, err := os.Stat(missingStateRoot); !os.IsNotExist(err) {
		t.Fatalf("support bundle manifest without --state-root should not create local state root: %v", err)
	}
	assertSupportBundleCLISafe(t, output.String())
}

func TestSupportBundleManifestPreviewCommandReadOnlyStateRootDoesNotCreateMissingRoot(t *testing.T) {
	registryPath, app := writeTestRepairGroupRegistry(t)
	runtimeRoot := writeSupportBundleCLIRuntimeRoot(t)
	missingStateRoot := filepath.Join(t.TempDir(), "missing-state-root")

	var output bytes.Buffer
	err := run([]string{
		"support-bundle-manifest-preview",
		"--registry", registryPath,
		"--app", app,
		"--state-root", missingStateRoot,
		"--runtime-root", runtimeRoot,
	}, &output)
	if err != nil {
		t.Fatalf("support bundle manifest returned error: %v", err)
	}
	if _, err := os.Stat(missingStateRoot); !os.IsNotExist(err) {
		t.Fatalf("read-only support bundle manifest should not create missing state root: %v", err)
	}
	assertSupportBundleCLISafe(t, output.String())
}

func TestSupportBundleManifestPreviewCommandRequiresRecipeSource(t *testing.T) {
	runtimeRoot := writeSupportBundleCLIRuntimeRoot(t)
	var output bytes.Buffer
	if err := run([]string{"support-bundle-manifest-preview", "--runtime-root", runtimeRoot}, &output); err == nil {
		t.Fatalf("support-bundle-manifest-preview must require a recipe source")
	}
}

func writeSupportBundleCLIRuntimeRoot(t *testing.T) string {
	t.Helper()
	root := t.TempDir()
	if err := os.WriteFile(filepath.Join(root, "VERSION"), []byte("0.2.303\n"), 0o600); err != nil {
		t.Fatalf("WriteFile VERSION returned error: %v", err)
	}
	return root
}

func assertSupportBundleCLISafe(t *testing.T, text string) {
	t.Helper()
	serialized := strings.ToLower(text)
	for _, forbidden := range []string{"prefix", ".exe", "program files", "qemu-system", "proton", "wine ", "wine/", ".wine", "virtual machine", "/home", "/users", "/private", "/var", "/opt", "/tmp", "file://"} {
		if strings.Contains(serialized, forbidden) {
			t.Fatalf("support bundle CLI exposed forbidden term %q: %s", forbidden, text)
		}
	}
}
