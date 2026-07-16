package main

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"
)

func TestKDEActionDependencyGraphPreviewCommand(t *testing.T) {
	registryPath, app := writeTestRepairGroupRegistry(t)
	var output bytes.Buffer
	err := run([]string{
		"kde-action-dependency-graph-preview",
		"--registry", registryPath,
		"--app", app,
		"--decision", "approved",
		"file:///home/test/Documents/book.xls",
	}, &output)
	if err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["schema_version"] != "xnix.runtime.kde_action_dependency_graph.v1" ||
		payload["request_type"] != "kde-action-dependency-graph-preview" ||
		payload["graph_type"] != "compatibility-center-action-dependency-graph" ||
		payload["runtime_method"] != "GetKDEActionDependencyGraph" ||
		payload["read_method"] != "GetKDEActionDependencyGraphPreview" ||
		payload["application_id"] != app ||
		payload["action_node_count"] != float64(7) ||
		payload["evidence_node_count"] != float64(35) ||
		payload["gate_node_count"] != float64(7) ||
		payload["missing_evidence_count"] != float64(35) ||
		payload["blocked_action_count"] != float64(7) ||
		payload["dependency_graph_created"] != true ||
		payload["dependency_graph_persisted"] != false ||
		payload["runtime_launch_approval"] != false ||
		payload["launch_allowed"] != false ||
		payload["execution_started"] != false ||
		payload["backend_process_started"] != false ||
		payload["request_objects_created"] != false ||
		payload["permission_grant_created"] != false ||
		payload["settings_persisted"] != false ||
		payload["host_root_modified"] != false ||
		payload["state_root_path_exposed"] != false ||
		payload["raw_command_exposed"] != false ||
		payload["file_content_read"] != false ||
		payload["backend_details_exposed"] != false {
		t.Fatalf("unexpected dependency graph payload: %#v", payload)
	}
	validation := payload["receipt_validation"].(map[string]any)
	if validation["rejects_mismatched_app_id"] != true ||
		validation["rejects_malformed_operation_id"] != true ||
		validation["rejects_path_escape_evidence"] != true ||
		validation["rejects_unsafe_side_effects"] != true ||
		validation["receipt_recorded"] != false ||
		validation["request_objects_created"] != false ||
		validation["permission_grant_created"] != false ||
		validation["settings_persisted"] != false ||
		validation["execution_started"] != false ||
		validation["host_root_modified"] != false {
		t.Fatalf("unexpected receipt validation: %#v", validation)
	}
	nodes := payload["nodes"].([]any)
	if len(nodes) != 49 {
		t.Fatalf("unexpected node count: %#v", payload)
	}
	if !strings.Contains(output.String(), "portal-file-access-receipt") ||
		!strings.Contains(output.String(), "runtime-write-gate-preview") ||
		!strings.Contains(output.String(), "review-file-manager-action") {
		t.Fatalf("dependency graph did not expose expected evidence ids: %s", output.String())
	}
	assertKDEShellPayloadSafe(t, output.String())
}

func TestKDEActionDependencyGraphPreviewCommandRejectsMalformedInputs(t *testing.T) {
	registryPath, app := writeTestRepairGroupRegistry(t)
	var output bytes.Buffer
	if err := run([]string{"kde-action-dependency-graph-preview", "--registry", registryPath, "--app", app}, &output); err == nil {
		t.Fatalf("command accepted missing decision")
	}
	output.Reset()
	if err := run([]string{"kde-action-dependency-graph-preview", "--registry", registryPath, "--app", app, "--decision", "approved", "https://example.invalid/book.xls"}, &output); err == nil {
		t.Fatalf("command accepted a non-file URI")
	}
}
