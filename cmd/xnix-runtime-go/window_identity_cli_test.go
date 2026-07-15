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

func decodeWindowIdentityPayload(t *testing.T, data []byte) map[string]any {
	t.Helper()
	var payload map[string]any
	if err := json.Unmarshal(data, &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	return payload
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
