package main

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"
)

func TestRuntimeOwnerProcessPreviewCommandRendersGoReadModel(t *testing.T) {
	var output bytes.Buffer
	err := run([]string{"runtime-owner-process-preview", "--root", projectRootForRuntimeServiceBindingCommandTest(t)}, &output)
	if err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["version"] != "0.2.211" ||
		payload["schema_version"] != "xnix.runtime.owner_process.v1" ||
		payload["request_type"] != "runtime-owner-process-preview" ||
		payload["process_type"] != "runtime-owner-process" ||
		payload["source"] != "runtime-service-binding-preview+libexec-wrapper+go-owner-target" ||
		payload["runtime_method"] != "GetRuntimeOwnerProcess" ||
		payload["read_method"] != "GetRuntimeOwnerProcessPreview" {
		t.Fatalf("unexpected Runtime owner process CLI schema: %#v", payload)
	}
	if payload["bus_name"] != "org.xnix.Compatibility1" ||
		payload["object_path"] != "/org/xnix/Compatibility1" ||
		payload["interface"] != "org.xnix.Compatibility1" {
		t.Fatalf("unexpected Runtime D-Bus identity: %#v", payload)
	}
	if payload["current_owner_entrypoint"] != "libexec/xnix/compatd" ||
		payload["current_owner_language"] != "ruby-wrapper" ||
		payload["target_owner_language"] != "go" {
		t.Fatalf("unexpected owner process language metadata: %#v", payload)
	}
	checks := payload["checks"].([]any)
	checkIDs := payload["check_ids"].([]any)
	expectedIDs := []string{"service-activation", "packaged-entrypoint", "go-owner-target", "production-owner-loop", "host-safety-boundary"}
	expectedStatuses := []string{"pass", "pass", "pass", "pending", "pass"}
	if len(checks) != len(expectedIDs) || len(checkIDs) != len(expectedIDs) {
		t.Fatalf("unexpected checks: %#v ids=%#v", checks, checkIDs)
	}
	for index, id := range expectedIDs {
		check := checks[index].(map[string]any)
		if check["id"] != id || checkIDs[index] != id || check["status"] != expectedStatuses[index] {
			t.Fatalf("unexpected check at %d: %#v ids=%#v", index, checks, checkIDs)
		}
	}
	counts := payload["counts"].(map[string]any)
	if counts["total"] != float64(5) ||
		counts["passed"] != float64(4) ||
		counts["pending"] != float64(1) ||
		counts["blocked"] != float64(0) {
		t.Fatalf("unexpected owner process counts: %#v", counts)
	}
	if payload["runtime_owned"] != true ||
		payload["go_runtime_backed"] != true ||
		payload["kde_policy_owner"] != false ||
		payload["kde_may_claim_runtime_ownership"] != false ||
		payload["service_activation_ready"] != true ||
		payload["packaged_entrypoint_ready"] != true ||
		payload["go_owner_process_ready"] != true ||
		payload["production_owner_process_ready"] != false ||
		payload["system_service_started"] != false ||
		payload["production_bus_claimed"] != false ||
		payload["network_required"] != false ||
		payload["host_root_modified"] != false ||
		payload["privileged_container_required"] != false ||
		payload["backend_details_exposed"] != false {
		t.Fatalf("unexpected owner process safety flags: %#v", payload)
	}
	if strings.Contains(strings.ToLower(output.String()), "prefix") ||
		strings.Contains(strings.ToLower(output.String()), ".exe") ||
		strings.Contains(strings.ToLower(output.String()), "program files") {
		t.Fatalf("Runtime owner process CLI exposed backend terms: %s", output.String())
	}
}
