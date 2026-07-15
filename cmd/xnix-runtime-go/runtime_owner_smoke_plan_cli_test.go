package main

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"
)

func TestRuntimeOwnerSmokePlanPreviewCommandRendersGoReadModel(t *testing.T) {
	var output bytes.Buffer
	err := run([]string{"runtime-owner-smoke-plan-preview", "--root", projectRootForRuntimeServiceBindingCommandTest(t)}, &output)
	if err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["version"] != "0.2.222" ||
		payload["schema_version"] != "xnix.runtime.owner_smoke_plan.v1" ||
		payload["request_type"] != "runtime-owner-smoke-plan-preview" ||
		payload["plan_type"] != "runtime-owner-smoke-plan" ||
		payload["source"] != "runtime-live-owner-gate-preview+runtime-service-binding-preview" ||
		payload["runtime_method"] != "GetRuntimeOwnerSmokePlan" ||
		payload["read_method"] != "GetRuntimeOwnerSmokePlanPreview" {
		t.Fatalf("unexpected Runtime owner smoke plan CLI schema: %#v", payload)
	}
	if payload["bus_name"] != "org.xnix.Compatibility1" ||
		payload["object_path"] != "/org/xnix/Compatibility1" ||
		payload["interface"] != "org.xnix.Compatibility1" {
		t.Fatalf("unexpected Runtime D-Bus identity: %#v", payload)
	}
	liveOwnerGate := payload["live_owner_gate"].(map[string]any)
	if liveOwnerGate["request_type"] != "runtime-live-owner-gate-preview" ||
		liveOwnerGate["gate_type"] != "runtime-live-owner-gate" ||
		liveOwnerGate["activation_binding_ready"] != true ||
		liveOwnerGate["live_dbus_owner_ready"] != false ||
		liveOwnerGate["production_owner_enabled"] != false ||
		liveOwnerGate["owner_transition_ready"] != false {
		t.Fatalf("unexpected live owner gate summary: %#v", liveOwnerGate)
	}
	steps := payload["steps"].([]any)
	stepIDs := payload["step_ids"].([]any)
	expectedIDs := []string{
		"validate-activation-files",
		"start-packaged-runtime-owner",
		"assert-stable-bus-name",
		"check-read-only-method-parity",
		"reject-write-methods",
		"verify-non-production-smoke-adapter-boundary",
		"report-kde-safe-summary",
	}
	expectedStatuses := []string{"pass", "pending", "pending", "pending", "pending", "pending", "pending"}
	if len(steps) != len(expectedIDs) || len(stepIDs) != len(expectedIDs) {
		t.Fatalf("unexpected owner smoke steps: %#v ids=%#v", steps, stepIDs)
	}
	for index, id := range expectedIDs {
		step := steps[index].(map[string]any)
		if step["id"] != id || stepIDs[index] != id || step["status"] != expectedStatuses[index] {
			t.Fatalf("unexpected step at %d: %#v ids=%#v", index, steps, stepIDs)
		}
	}
	counts := payload["counts"].(map[string]any)
	if counts["total"] != float64(7) ||
		counts["passed"] != float64(1) ||
		counts["pending"] != float64(6) ||
		counts["blocked"] != float64(0) ||
		payload["pending_step_count"] != float64(6) {
		t.Fatalf("unexpected owner smoke counts: counts=%#v pending=%#v", counts, payload["pending_step_count"])
	}
	if payload["runtime_owned"] != true ||
		payload["go_runtime_backed"] != true ||
		payload["kde_policy_owner"] != false ||
		payload["activation_binding_ready"] != true ||
		payload["live_dbus_owner_ready"] != false ||
		payload["production_owner_enabled"] != false ||
		payload["owner_transition_ready"] != false ||
		payload["smoke_state"] != "planned" ||
		payload["smoke_environment"] != "restricted-session" ||
		payload["network_required"] != false ||
		payload["host_root_modified"] != false ||
		payload["privileged_container_required"] != false ||
		payload["system_service_started"] != false ||
		payload["production_bus_claimed"] != false ||
		payload["backend_details_exposed"] != false {
		t.Fatalf("unexpected owner smoke safety flags: %#v", payload)
	}
	blockedActions := payload["blocked_actions"].([]any)
	if len(blockedActions) != 5 ||
		blockedActions[0] != "Do not start a host system service from the smoke plan." ||
		blockedActions[1] != "Do not claim the production Runtime bus name from the smoke adapter." {
		t.Fatalf("unexpected owner smoke blocked actions: %#v", blockedActions)
	}
	if strings.Contains(strings.ToLower(output.String()), "prefix") ||
		strings.Contains(strings.ToLower(output.String()), ".exe") ||
		strings.Contains(strings.ToLower(output.String()), "program files") {
		t.Fatalf("Runtime owner smoke plan CLI exposed backend terms: %s", output.String())
	}
}
