package main

import (
	"bytes"
	"encoding/json"
	"testing"
)

func TestStateRootPreviewCommandRendersPlannedStorage(t *testing.T) {
	registryPath, app := writeTestRepairGroupRegistry(t)

	var output bytes.Buffer
	if err := run([]string{"state-root-preview", "--registry", registryPath, "--app", app}, &output); err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["schema_version"] != "xnix.runtime.state_root.v1" ||
		payload["request_type"] != "state-root-preview" ||
		payload["runtime_method"] != "GetApplicationStateRoot" ||
		payload["go_runtime_backed"] != true ||
		payload["directories_created"] != false ||
		payload["host_root_modified"] != false ||
		payload["user_documents_included"] != false ||
		payload["backend_details_exposed"] != false {
		t.Fatalf("unexpected state root payload: %#v", payload)
	}
	scopeIDs := payload["managed_scope_ids"].([]any)
	if len(scopeIDs) != 4 || scopeIDs[0] != "application-data" {
		t.Fatalf("unexpected state root scopes: %#v", scopeIDs)
	}
}

func TestSnapshotPlanPreviewCommandRendersClosedSnapshot(t *testing.T) {
	var output bytes.Buffer
	if err := run([]string{"snapshot-plan-preview", "--app", "org.example.ledger", "--reason", "before-engine-change"}, &output); err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["schema_version"] != "xnix.runtime.snapshot_plan.v1" ||
		payload["request_type"] != "snapshot-plan-preview" ||
		payload["runtime_method"] != "GetSnapshotPlan" ||
		payload["reason"] != "before-engine-change" ||
		payload["go_runtime_backed"] != true ||
		payload["snapshot_request_created"] != false ||
		payload["snapshot_created"] != false ||
		payload["restore_requested"] != false ||
		payload["restore_executed"] != false ||
		payload["host_root_modified"] != false ||
		payload["backend_details_exposed"] != false {
		t.Fatalf("unexpected snapshot payload: %#v", payload)
	}
}

func TestPortalAccessPolicyPreviewCommandRendersMediatedAccess(t *testing.T) {
	var output bytes.Buffer
	if err := run([]string{"portal-access-policy-preview", "--app", "org.example.ledger", "--operation", "camera"}, &output); err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["schema_version"] != "xnix.runtime.portal_access_policy.v1" ||
		payload["request_type"] != "portal-access-policy-preview" ||
		payload["runtime_method"] != "GetPortalAccessPolicy" ||
		payload["operation"] != "camera" ||
		payload["decision"] != "deny" ||
		payload["go_runtime_backed"] != true ||
		payload["portal_required"] != true ||
		payload["direct_access_allowed"] != false ||
		payload["request_object_created"] != false ||
		payload["permission_granted"] != false ||
		payload["host_permission_changed"] != false ||
		payload["host_root_modified"] != false ||
		payload["backend_details_exposed"] != false {
		t.Fatalf("unexpected Portal policy payload: %#v", payload)
	}
}
