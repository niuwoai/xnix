package main

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
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

func TestPortalRequestRecordCommandPersistsPermissionFlow(t *testing.T) {
	stateRoot := t.TempDir()
	var createOutput bytes.Buffer
	if err := run([]string{"portal-request-record", "--state-root", stateRoot, "--app", "org.example.ledger", "--operation", "file-open", "--action", "create", "--reason", "Open a spreadsheet"}, &createOutput); err != nil {
		t.Fatalf("create run returned error: %v", err)
	}

	var created map[string]any
	if err := json.Unmarshal(createOutput.Bytes(), &created); err != nil {
		t.Fatalf("Unmarshal created returned error: %v", err)
	}
	if created["schema_version"] != "xnix.runtime.portal_request_record.v1" ||
		created["record_type"] != "portal-permission-request-record" ||
		created["source"] != "go-runtime-state-root-portal-broker" ||
		created["action"] != "create" ||
		created["relative_path"] != "portal-requests/xnix_org_example_ledger_file_open_1.json" ||
		created["state_root_path_exposed"] != false ||
		created["real_portal_call_enabled"] != false ||
		created["request_object_created"] != true ||
		created["permission_granted"] != false ||
		created["execution_approved"] != false ||
		created["host_permission_changed"] != false ||
		created["host_root_modified"] != false ||
		created["backend_details_exposed"] != false {
		t.Fatalf("unexpected created Portal request record: %#v", created)
	}
	request := created["request"].(map[string]any)
	handleToken := request["handle_token"].(string)
	if request["state"] != "pending-user-mediation" ||
		request["permission_state"] != "pending" ||
		request["direct_access_allowed"] != false ||
		request["host_permission_changed"] != false ||
		request["backend_details_exposed"] != false {
		t.Fatalf("unexpected created Portal request: %#v", request)
	}
	if _, err := os.Stat(filepath.Join(stateRoot, "portal-requests", "xnix_org_example_ledger_file_open_1.json")); err != nil {
		t.Fatalf("portal-request-record must persist under state root: %v", err)
	}

	var resolveOutput bytes.Buffer
	if err := run([]string{"portal-request-record", "--state-root", stateRoot, "--action", "resolve", "--handle-token", handleToken, "--outcome", "granted"}, &resolveOutput); err != nil {
		t.Fatalf("resolve run returned error: %v", err)
	}
	var resolved map[string]any
	if err := json.Unmarshal(resolveOutput.Bytes(), &resolved); err != nil {
		t.Fatalf("Unmarshal resolved returned error: %v", err)
	}
	resolvedRequest := resolved["request"].(map[string]any)
	if resolved["action"] != "resolve" ||
		resolved["permission_granted"] != true ||
		resolved["execution_approved"] != false ||
		resolvedRequest["state"] != "granted" ||
		resolvedRequest["permission_state"] != "granted" {
		t.Fatalf("unexpected resolved Portal request record: %#v", resolved)
	}

	var completeOutput bytes.Buffer
	if err := run([]string{"portal-request-record", "--state-root", stateRoot, "--action", "complete", "--handle-token", handleToken}, &completeOutput); err != nil {
		t.Fatalf("complete run returned error: %v", err)
	}
	var completed map[string]any
	if err := json.Unmarshal(completeOutput.Bytes(), &completed); err != nil {
		t.Fatalf("Unmarshal completed returned error: %v", err)
	}
	completedRequest := completed["request"].(map[string]any)
	if completed["action"] != "complete" ||
		completed["permission_granted"] != true ||
		completed["execution_approved"] != false ||
		completedRequest["state"] != "completed" ||
		completedRequest["permission_state"] != "granted" {
		t.Fatalf("unexpected completed Portal request record: %#v", completed)
	}
}

func TestPortalRequestRecordCommandRejectsIncompleteRequests(t *testing.T) {
	stateRoot := t.TempDir()
	var missingStateRoot bytes.Buffer
	if err := run([]string{"portal-request-record", "--action", "create", "--app", "org.example.ledger", "--operation", "file-open"}, &missingStateRoot); err == nil {
		t.Fatalf("portal-request-record must require --state-root")
	}
	var missingAction bytes.Buffer
	if err := run([]string{"portal-request-record", "--state-root", stateRoot}, &missingAction); err == nil {
		t.Fatalf("portal-request-record must require --action")
	}
	var missingOperation bytes.Buffer
	if err := run([]string{"portal-request-record", "--state-root", stateRoot, "--action", "create", "--app", "org.example.ledger"}, &missingOperation); err == nil {
		t.Fatalf("portal-request-record create must require --operation")
	}
	var missingHandle bytes.Buffer
	if err := run([]string{"portal-request-record", "--state-root", stateRoot, "--action", "inspect"}, &missingHandle); err == nil {
		t.Fatalf("portal-request-record non-create action must require --handle-token")
	}
	var missingOutcome bytes.Buffer
	if err := run([]string{"portal-request-record", "--state-root", stateRoot, "--action", "resolve", "--handle-token", "xnix_org_example_ledger_file_open_1"}, &missingOutcome); err == nil {
		t.Fatalf("portal-request-record resolve must require --outcome")
	}
}
