package main

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"
)

func TestKDEIntegrationStatusPreviewCommand(t *testing.T) {
	var output bytes.Buffer
	if err := run([]string{"kde-integration-status-preview"}, &output); err != nil {
		t.Fatalf("run returned error: %v", err)
	}
	payload := decodeKDEShellPayload(t, output.Bytes())
	if payload["schema_version"] != "xnix.runtime.kde_integration_status.v1" ||
		payload["request_type"] != "kde-integration-status-preview" ||
		payload["runtime_method"] != "GetKDEIntegrationStatus" ||
		payload["read_method"] != "GetKDEIntegrationStatusPreview" ||
		payload["entry_point_count"] != float64(7) ||
		payload["go_runtime_backed"] != true ||
		payload["kde_policy_owner"] != false ||
		payload["host_root_modified"] != false ||
		payload["backend_details_exposed"] != false {
		t.Fatalf("unexpected KDE integration status payload: %#v", payload)
	}
	assertKDEShellPayloadSafe(t, output.String())
}

func TestKDEShellIntegrationPreviewCommand(t *testing.T) {
	var output bytes.Buffer
	if err := run([]string{"kde-shell-integration-preview"}, &output); err != nil {
		t.Fatalf("run returned error: %v", err)
	}
	payload := decodeKDEShellPayload(t, output.Bytes())
	if payload["schema_version"] != "xnix.runtime.kde_shell_integration.v1" ||
		payload["request_type"] != "kde-shell-integration-preview" ||
		payload["runtime_method"] != "GetKDEShellIntegrationPlan" ||
		payload["read_method"] != "GetKDEShellIntegrationPlanPreview" ||
		payload["component_count"] != float64(9) ||
		payload["go_runtime_backed"] != true ||
		payload["kde_policy_owner"] != false ||
		payload["plasma_fork_required"] != false ||
		payload["shell_configuration_written"] != false ||
		payload["host_root_modified"] != false {
		t.Fatalf("unexpected KDE shell integration payload: %#v", payload)
	}
	assertKDEShellPayloadSafe(t, output.String())
}

func TestKDEApplicationSurfacePreviewCommand(t *testing.T) {
	registryPath, app := writeTestRepairGroupRegistry(t)
	var output bytes.Buffer
	if err := run([]string{"kde-application-surface-preview", "--registry", registryPath, "--app", app}, &output); err != nil {
		t.Fatalf("run returned error: %v", err)
	}
	payload := decodeKDEShellPayload(t, output.Bytes())
	if payload["schema_version"] != "xnix.runtime.kde_application_surface.v1" ||
		payload["request_type"] != "kde-application-surface-preview" ||
		payload["runtime_method"] != "GetKDEApplicationSurfacePlan" ||
		payload["read_method"] != "GetKDEApplicationSurfacePlanPreview" ||
		payload["entry_point_count"] != float64(7) ||
		payload["normal_linux_application_surface"] != true ||
		payload["go_runtime_backed"] != true ||
		payload["launch_enabled"] != false ||
		payload["host_root_modified"] != false ||
		payload["backend_details_exposed"] != false {
		t.Fatalf("unexpected KDE application surface payload: %#v", payload)
	}
	application := payload["application"].(map[string]any)
	if application["id"] != app ||
		application["desktop_file"] == "" {
		t.Fatalf("unexpected application payload: %#v", application)
	}
	assertKDEShellPayloadSafe(t, output.String())
}

func TestKDEApplicationSurfacePreviewCommandConsumesActivationRoot(t *testing.T) {
	registryPath, app := writeTestRepairGroupRegistry(t)
	stageRoot := t.TempDir()

	var stageOutput bytes.Buffer
	if err := run([]string{"desktop-activation-stage", "--registry", registryPath, "--app", app, "--mode", "development", "--staging-root", stageRoot}, &stageOutput); err != nil {
		t.Fatalf("stage run returned error: %v", err)
	}

	var output bytes.Buffer
	if err := run([]string{"kde-application-surface-preview", "--registry", registryPath, "--app", app, "--activation-root", stageRoot}, &output); err != nil {
		t.Fatalf("run returned error: %v", err)
	}
	payload := decodeKDEShellPayload(t, output.Bytes())
	if payload["source"] != "go-runtime-kde-application-surface+desktop-activation-receipt" ||
		payload["activation_receipt_root"] != true ||
		payload["activation_receipt_backed"] != true ||
		payload["activation_receipt_path"] != "usr/share/xnix/compatibility/activation-receipts/org.example.ledger.json" {
		t.Fatalf("KDE application surface did not consume activation receipt: %#v", payload)
	}
	if payload["launch_enabled"] != false ||
		payload["backend_process_started"] != false ||
		payload["desktop_files_written"] != false ||
		payload["mimeapps_written"] != false ||
		payload["host_root_modified"] != false ||
		payload["backend_command_exposed"] != false ||
		payload["raw_windows_executable_exposed"] != false ||
		payload["backend_details_exposed"] != false {
		t.Fatalf("receipt-backed KDE application surface must remain gated: %#v", payload)
	}
	if strings.Contains(output.String(), stageRoot) {
		t.Fatalf("KDE application surface exposed activation root: %s", output.String())
	}
	assertKDEShellPayloadSafe(t, output.String())
}

func decodeKDEShellPayload(t *testing.T, data []byte) map[string]any {
	t.Helper()
	var payload map[string]any
	if err := json.Unmarshal(data, &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	return payload
}

func assertKDEShellPayloadSafe(t *testing.T, text string) {
	t.Helper()
	serialized := strings.ToLower(text)
	for _, forbidden := range []string{"prefix", ".exe", "program files", "qemu-system", "proton", "wine ", "virtual machine"} {
		if strings.Contains(serialized, forbidden) {
			t.Fatalf("KDE shell CLI exposed forbidden term %q: %s", forbidden, text)
		}
	}
}
