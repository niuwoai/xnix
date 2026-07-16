package main

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"
)

func TestTaskManagerIdentityPreviewCommand(t *testing.T) {
	registryPath, app := writeTestRepairGroupRegistry(t)
	var output bytes.Buffer
	if err := run([]string{"task-manager-identity-preview", "--registry", registryPath, "--app", app}, &output); err != nil {
		t.Fatalf("run returned error: %v", err)
	}
	payload := decodeWindowIdentityPayload(t, output.Bytes())
	if payload["schema_version"] != "xnix.runtime.task_manager_identity.v1" ||
		payload["request_type"] != "task-manager-identity-preview" ||
		payload["runtime_method"] != "GetTaskManagerIdentityPlan" ||
		payload["read_method"] != "GetTaskManagerIdentityPlanPreview" ||
		payload["application_id"] != app ||
		payload["pinning_allowed"] != true ||
		payload["restore_allowed"] != true ||
		payload["task_manager_entry_active"] != false ||
		payload["execution_started"] != false ||
		payload["host_root_modified"] != false ||
		payload["backend_details_exposed"] != false {
		t.Fatalf("unexpected task manager identity payload: %#v", payload)
	}
	assertWindowIdentityPayloadSafe(t, output.String())
}

func TestTaskManagerIdentityPreviewCommandConsumesActivationRoot(t *testing.T) {
	registryPath, app := writeTestRepairGroupRegistry(t)
	stageRoot := t.TempDir()

	var stageOutput bytes.Buffer
	if err := run([]string{"desktop-activation-stage", "--registry", registryPath, "--app", app, "--mode", "development", "--staging-root", stageRoot}, &stageOutput); err != nil {
		t.Fatalf("stage run returned error: %v", err)
	}

	var output bytes.Buffer
	if err := run([]string{"task-manager-identity-preview", "--registry", registryPath, "--app", app, "--activation-root", stageRoot}, &output); err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	payload := decodeWindowIdentityPayload(t, output.Bytes())
	if payload["source"] != "window-identity-preview+desktop-activation-receipt" ||
		payload["activation_receipt_root"] != true ||
		payload["activation_receipt_backed"] != true ||
		payload["activation_receipt_path"] != "usr/share/xnix/compatibility/activation-receipts/org.example.ledger.json" {
		t.Fatalf("task manager identity did not consume activation receipt: %#v", payload)
	}
	if payload["task_manager_entry_active"] != false ||
		payload["window_observation_started"] != false ||
		payload["launch_enabled"] != false ||
		payload["execution_started"] != false ||
		payload["host_root_modified"] != false ||
		payload["backend_details_exposed"] != false {
		t.Fatalf("receipt-backed task manager identity must remain gated: %#v", payload)
	}
	if strings.Contains(output.String(), stageRoot) {
		t.Fatalf("task manager identity exposed activation root: %s", output.String())
	}
	assertWindowIdentityPayloadSafe(t, output.String())
}

func TestTaskManagerIdentityPreviewCommandConsumesSessionRoot(t *testing.T) {
	registryPath, app := writeTestRepairGroupRegistry(t)
	sessionRoot, requestID := writeExecutionSessionRecordWithCLI(t, app)

	var output bytes.Buffer
	if err := run([]string{"task-manager-identity-preview", "--registry", registryPath, "--app", app, "--session-root", sessionRoot, "--session-request-id", requestID}, &output); err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	payload := decodeWindowIdentityPayload(t, output.Bytes())
	if payload["source"] != "window-identity-preview+execution-session-record" ||
		payload["execution_session_root"] != true ||
		payload["execution_session_backed"] != true ||
		payload["execution_session_path"] != "execution-ledger/sessions/"+requestID+".json" ||
		payload["execution_session_state"] != "blocked" {
		t.Fatalf("task manager identity did not consume session record: %#v", payload)
	}
	if payload["task_manager_entry_active"] != false ||
		payload["window_observation_started"] != false ||
		payload["launch_enabled"] != false ||
		payload["execution_started"] != false ||
		payload["host_root_modified"] != false ||
		payload["backend_details_exposed"] != false {
		t.Fatalf("session-backed task manager identity must remain gated: %#v", payload)
	}
	if strings.Contains(output.String(), sessionRoot) {
		t.Fatalf("task manager identity exposed session root: %s", output.String())
	}
	assertWindowIdentityPayloadSafe(t, output.String())
}

func TestKWinWindowRulePreviewCommand(t *testing.T) {
	registryPath, app := writeTestRepairGroupRegistry(t)
	var output bytes.Buffer
	if err := run([]string{"kwin-window-rule-preview", "--registry", registryPath, "--app", app}, &output); err != nil {
		t.Fatalf("run returned error: %v", err)
	}
	payload := decodeWindowIdentityPayload(t, output.Bytes())
	if payload["schema_version"] != "xnix.runtime.kwin_window_rule.v1" ||
		payload["request_type"] != "kwin-window-rule-preview" ||
		payload["runtime_method"] != "GetKWinWindowRulePlan" ||
		payload["read_method"] != "GetKWinWindowRulePlanPreview" ||
		payload["application_id"] != app ||
		payload["window_manager_policy_only"] != true ||
		payload["runtime_owns_backend_policy"] != true ||
		payload["kwin_rule_applied"] != false ||
		payload["execution_started"] != false ||
		payload["host_root_modified"] != false ||
		payload["backend_details_exposed"] != false {
		t.Fatalf("unexpected KWin window rule payload: %#v", payload)
	}
	match := payload["match"].(map[string]any)
	set := payload["set"].(map[string]any)
	if match["resource_name"] != app ||
		set["application_id"] != app ||
		set["task_manager_grouping_key"] != app ||
		set["show_in_switcher"] != true {
		t.Fatalf("unexpected KWin match/set payload: match=%#v set=%#v", match, set)
	}
	assertWindowIdentityPayloadSafe(t, output.String())
}

func TestKWinWindowRulePreviewCommandConsumesActivationRoot(t *testing.T) {
	registryPath, app := writeTestRepairGroupRegistry(t)
	stageRoot := t.TempDir()

	var stageOutput bytes.Buffer
	if err := run([]string{"desktop-activation-stage", "--registry", registryPath, "--app", app, "--mode", "development", "--staging-root", stageRoot}, &stageOutput); err != nil {
		t.Fatalf("stage run returned error: %v", err)
	}

	var output bytes.Buffer
	if err := run([]string{"kwin-window-rule-preview", "--registry", registryPath, "--app", app, "--activation-root", stageRoot}, &output); err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	payload := decodeWindowIdentityPayload(t, output.Bytes())
	if payload["source"] != "window-identity-preview+desktop-activation-receipt" ||
		payload["activation_receipt_root"] != true ||
		payload["activation_receipt_backed"] != true ||
		payload["activation_receipt_path"] != "usr/share/xnix/compatibility/activation-receipts/org.example.ledger.json" {
		t.Fatalf("KWin window rule did not consume activation receipt: %#v", payload)
	}
	if payload["kwin_rule_applied"] != false ||
		payload["task_manager_entry_active"] != false ||
		payload["window_observation_started"] != false ||
		payload["launch_enabled"] != false ||
		payload["execution_started"] != false ||
		payload["host_root_modified"] != false ||
		payload["backend_details_exposed"] != false {
		t.Fatalf("receipt-backed KWin window rule must remain gated: %#v", payload)
	}
	if strings.Contains(output.String(), stageRoot) {
		t.Fatalf("KWin window rule exposed activation root: %s", output.String())
	}
	assertWindowIdentityPayloadSafe(t, output.String())
}

func TestKWinWindowRulePreviewCommandConsumesSessionRoot(t *testing.T) {
	registryPath, app := writeTestRepairGroupRegistry(t)
	sessionRoot, requestID := writeExecutionSessionRecordWithCLI(t, app)

	var output bytes.Buffer
	if err := run([]string{"kwin-window-rule-preview", "--registry", registryPath, "--app", app, "--session-root", sessionRoot, "--session-request-id", requestID}, &output); err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	payload := decodeWindowIdentityPayload(t, output.Bytes())
	if payload["source"] != "window-identity-preview+execution-session-record" ||
		payload["execution_session_root"] != true ||
		payload["execution_session_backed"] != true ||
		payload["execution_session_path"] != "execution-ledger/sessions/"+requestID+".json" ||
		payload["execution_session_state"] != "blocked" {
		t.Fatalf("KWin rule did not consume session record: %#v", payload)
	}
	if payload["kwin_rule_applied"] != false ||
		payload["task_manager_entry_active"] != false ||
		payload["window_observation_started"] != false ||
		payload["launch_enabled"] != false ||
		payload["execution_started"] != false ||
		payload["host_root_modified"] != false ||
		payload["backend_details_exposed"] != false {
		t.Fatalf("session-backed KWin rule must remain gated: %#v", payload)
	}
	if strings.Contains(output.String(), sessionRoot) {
		t.Fatalf("KWin rule exposed session root: %s", output.String())
	}
	assertWindowIdentityPayloadSafe(t, output.String())
}

func decodeWindowIdentityPayload(t *testing.T, data []byte) map[string]any {
	t.Helper()
	var payload map[string]any
	if err := json.Unmarshal(data, &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	return payload
}

func writeExecutionSessionRecordWithCLI(t *testing.T, app string) (string, string) {
	t.Helper()
	root := t.TempDir()
	requestID := "xnix-exec-" + strings.ReplaceAll(app, ".", "-") + "-1"
	var output bytes.Buffer
	if err := run([]string{"execution-ledger-record", "--state-root", root, "--app", app}, &output); err != nil {
		t.Fatalf("execution-ledger-record returned error: %v", err)
	}
	output.Reset()
	if err := run([]string{"execution-session-record", "--state-root", root, "--request-id", requestID}, &output); err != nil {
		t.Fatalf("execution-session-record returned error: %v", err)
	}
	return root, requestID
}

func assertWindowIdentityPayloadSafe(t *testing.T, text string) {
	t.Helper()
	serialized := strings.ToLower(text)
	for _, forbidden := range []string{"prefix", ".exe", "program files", "qemu-system", "proton", "wine ", "virtual machine"} {
		if strings.Contains(serialized, forbidden) {
			t.Fatalf("window identity CLI exposed forbidden term %q: %s", forbidden, text)
		}
	}
}
