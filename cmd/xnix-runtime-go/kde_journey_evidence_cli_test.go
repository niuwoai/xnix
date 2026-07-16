package main

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"
)

func TestKDEJourneyEvidencePreviewCommand(t *testing.T) {
	registryPath, app := writeTestRepairGroupRegistry(t)
	var output bytes.Buffer
	err := run([]string{
		"kde-journey-evidence-preview",
		"--registry", registryPath,
		"--app", app,
		"--decision", "approved",
		"--runtime-root", "../..",
		"file:///home/test/Documents/book.xls",
	}, &output)
	if err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["schema_version"] != "xnix.runtime.kde_journey_evidence.v1" ||
		payload["request_type"] != "kde-journey-evidence-preview" ||
		payload["journey_type"] != "kde-seven-entrypoint-runtime-evidence" ||
		payload["runtime_method"] != "GetKDEJourneyEvidence" ||
		payload["read_method"] != "GetKDEJourneyEvidencePreview" ||
		payload["application_id"] != app ||
		payload["entry_point_count"] != float64(7) ||
		payload["cross_linked_read_model_count"] != float64(10) ||
		payload["shared_readiness_status"] != "not-ready" ||
		payload["ready"] != false ||
		payload["missing_evidence_count"] != float64(35) ||
		payload["blocked_action_count"] != float64(7) ||
		payload["journey_evidence_created"] != true ||
		payload["journey_evidence_persisted"] != false ||
		payload["runtime_write_methods_enabled"] != false ||
		payload["request_objects_created"] != false ||
		payload["permission_grant_created"] != false ||
		payload["settings_persisted"] != false ||
		payload["launch_allowed"] != false ||
		payload["launch_enabled"] != false ||
		payload["execution_started"] != false ||
		payload["backend_process_started"] != false ||
		payload["kwin_rule_applied"] != false ||
		payload["tray_bridge_activated"] != false ||
		payload["notification_sent"] != false ||
		payload["host_root_modified"] != false ||
		payload["state_root_path_exposed"] != false ||
		payload["raw_executable_exposed"] != false ||
		payload["raw_command_exposed"] != false ||
		payload["file_content_read"] != false ||
		payload["backend_details_exposed"] != false {
		t.Fatalf("unexpected journey evidence payload: %#v", payload)
	}
	agreement := payload["agreement"].(map[string]any)
	if agreement["app_id_consistent"] != true ||
		agreement["display_name_consistent"] != true ||
		agreement["desktop_file_consistent"] != true ||
		agreement["readiness_state_consistent"] != true ||
		agreement["disabled_action_state_shared"] != true ||
		agreement["unsafe_side_effects_disabled"] != true {
		t.Fatalf("unexpected journey agreement: %#v", agreement)
	}
	entryPoints := payload["entry_points"].([]any)
	if len(entryPoints) != 7 {
		t.Fatalf("unexpected entry point count: %#v", payload)
	}
	if !strings.Contains(output.String(), "launcher") ||
		!strings.Contains(output.String(), "task-manager") ||
		!strings.Contains(output.String(), "file-manager") ||
		!strings.Contains(output.String(), "system-tray") ||
		!strings.Contains(output.String(), "notification-center") ||
		!strings.Contains(output.String(), "ai-compatibility-center") ||
		!strings.Contains(output.String(), "unified-settings") ||
		!strings.Contains(output.String(), "application-readiness-preview") ||
		!strings.Contains(output.String(), "kde-action-dependency-graph-preview") ||
		!strings.Contains(output.String(), "kde-center-page-preview") {
		t.Fatalf("journey evidence did not include expected cross-links: %s", output.String())
	}
	assertKDEShellPayloadSafe(t, output.String())
}

func TestKDEJourneyEvidencePreviewCommandRejectsMalformedInputs(t *testing.T) {
	registryPath, app := writeTestRepairGroupRegistry(t)
	var output bytes.Buffer
	if err := run([]string{"kde-journey-evidence-preview", "--registry", registryPath, "--app", app, "--runtime-root", "../.."}, &output); err == nil {
		t.Fatalf("command accepted missing decision")
	}
	output.Reset()
	if err := run([]string{"kde-journey-evidence-preview", "--registry", registryPath, "--app", app, "--decision", "approved", "--runtime-root", "../..", "https://example.invalid/book.xls"}, &output); err == nil {
		t.Fatalf("command accepted a non-file URI")
	}
}
