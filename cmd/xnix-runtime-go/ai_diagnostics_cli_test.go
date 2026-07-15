package main

import (
	"bytes"
	"encoding/json"
	"testing"
)

func TestAIDiagnosticInputPreviewCommandRendersSafeInput(t *testing.T) {
	registryPath, app := writeTestRepairGroupRegistry(t)

	var output bytes.Buffer
	if err := run([]string{"ai-diagnostic-input-preview", "--registry", registryPath, "--app", app, "--issue", "engine-binding-pending", "--test-type", "preflight"}, &output); err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["schema_version"] != "xnix.runtime.ai_diagnostic_input.v1" ||
		payload["request_type"] != "ai-diagnostic-input-preview" ||
		payload["runtime_method"] != "GetAIDiagnosticInput" ||
		payload["read_method"] != "GetAIDiagnosticInputPreview" ||
		payload["go_runtime_backed"] != true ||
		payload["ai_provider_called"] != false ||
		payload["ai_provider_call_enabled"] != false ||
		payload["network_required"] != false ||
		payload["file_content_read"] != false ||
		payload["file_paths_exposed"] != false ||
		payload["request_object_created"] != false ||
		payload["host_root_modified"] != false ||
		payload["backend_details_exposed"] != false {
		t.Fatalf("unexpected AI diagnostic input payload: %#v", payload)
	}
	sections := payload["context_sections"].([]any)
	if len(sections) != 4 {
		t.Fatalf("unexpected context sections: %#v", sections)
	}
}

func TestAIDiagnosticRecommendationPreviewCommandRendersReviewOnlyActions(t *testing.T) {
	registryPath, app := writeTestRepairGroupRegistry(t)

	var output bytes.Buffer
	if err := run([]string{"ai-diagnostic-recommendation-preview", "--registry", registryPath, "--app", app, "--test-type", "smoke"}, &output); err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["schema_version"] != "xnix.runtime.ai_diagnostic_recommendation.v1" ||
		payload["request_type"] != "ai-diagnostic-recommendation-preview" ||
		payload["runtime_method"] != "GetAIDiagnosticRecommendation" ||
		payload["auto_execution_allowed"] != false ||
		payload["repair_execution_requested"] != false ||
		payload["repair_executed"] != false ||
		payload["backend_details_exposed"] != false {
		t.Fatalf("unexpected AI diagnostic recommendation payload: %#v", payload)
	}
	recommendations := payload["recommendations"].([]any)
	approvals := payload["approval_required_actions"].([]any)
	if len(recommendations) != 3 || len(approvals) != 1 {
		t.Fatalf("unexpected AI diagnostic recommendation actions: recommendations=%#v approvals=%#v", recommendations, approvals)
	}
}

func TestAIRepairApprovalGatePreviewCommandBlocksRepairExecution(t *testing.T) {
	registryPath, app := writeTestRepairGroupRegistry(t)

	var output bytes.Buffer
	if err := run([]string{"ai-repair-approval-gate-preview", "--registry", registryPath, "--app", app}, &output); err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["schema_version"] != "xnix.runtime.ai_repair_approval_gate.v1" ||
		payload["request_type"] != "ai-repair-approval-gate-preview" ||
		payload["runtime_method"] != "GetAIRepairApprovalGate" ||
		payload["gate_decision"] != "blocked-until-approval" ||
		payload["repair_execution_requested"] != false ||
		payload["repair_executed"] != false ||
		payload["auto_execution_allowed"] != false ||
		payload["host_root_modified"] != false ||
		payload["backend_details_exposed"] != false {
		t.Fatalf("unexpected AI repair approval gate payload: %#v", payload)
	}
	requiredGates := payload["required_gates"].([]any)
	if len(requiredGates) != 3 {
		t.Fatalf("unexpected required gates: %#v", requiredGates)
	}
}
