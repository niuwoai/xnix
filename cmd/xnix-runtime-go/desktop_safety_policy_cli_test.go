package main

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"
)

func TestDesktopSafetyPolicyPreviewCommand(t *testing.T) {
	var output bytes.Buffer
	if err := run([]string{"desktop-safety-policy-preview"}, &output); err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["schema_version"] != "xnix.runtime.desktop_safety_policy.v1" ||
		payload["request_type"] != "desktop-safety-policy-preview" ||
		payload["policy_type"] != "kde-first-user-facing-safety-policy" ||
		payload["runtime_method"] != "GetKDEIntegrationStatus" ||
		payload["entrypoint_count"] != float64(7) ||
		payload["runtime_owned"] != true ||
		payload["go_runtime_backed"] != true ||
		payload["kde_policy_owner"] != false ||
		payload["backend_terminology_hidden"] != true ||
		payload["write_methods_enabled"] != false ||
		payload["backend_launch_enabled"] != false ||
		payload["execution_enabled"] != false ||
		payload["real_portal_transport_enabled"] != false ||
		payload["ai_provider_call_enabled"] != false ||
		payload["host_root_modified"] != false ||
		payload["network_required"] != false ||
		payload["privileged_container_required"] != false {
		t.Fatalf("unexpected desktop safety policy payload: %#v", payload)
	}
	forbidden := stringSliceFromAny(t, payload["forbidden_user_terms"])
	for _, term := range []string{"prefix", "bottle", "wine", "proton", ".exe", "docker.sock"} {
		if !containsTestString(forbidden, term) {
			t.Fatalf("policy payload missing forbidden user term %q: %#v", term, forbidden)
		}
	}
	settings := stringSliceFromAny(t, payload["settings_field_ids"])
	if strings.Join(settings, ",") != "mode,preference,documents,downloads,camera,network,snapshots" {
		t.Fatalf("unexpected settings fields: %#v", settings)
	}
}

func TestDesktopSafetyPolicyPreviewCommandRejectsArguments(t *testing.T) {
	var output bytes.Buffer
	err := run([]string{"desktop-safety-policy-preview", "--app", "org.example"}, &output)
	if err == nil || !strings.Contains(err.Error(), "does not accept arguments") {
		t.Fatalf("expected argument rejection, got %v", err)
	}
}

func stringSliceFromAny(t *testing.T, value any) []string {
	t.Helper()
	values, ok := value.([]any)
	if !ok {
		t.Fatalf("expected JSON array, got %#v", value)
	}
	result := make([]string, 0, len(values))
	for _, item := range values {
		text, ok := item.(string)
		if !ok {
			t.Fatalf("expected string item, got %#v", item)
		}
		result = append(result, text)
	}
	return result
}

func containsTestString(values []string, candidate string) bool {
	for _, value := range values {
		if value == candidate {
			return true
		}
	}
	return false
}
