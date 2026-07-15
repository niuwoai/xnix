package main

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"
)

func TestRuntimeLiveOwnerGatePreviewCommandRendersGoReadModel(t *testing.T) {
	var output bytes.Buffer
	err := run([]string{"runtime-live-owner-gate-preview", "--root", projectRootForRuntimeServiceBindingCommandTest(t)}, &output)
	if err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["version"] != "0.2.183" ||
		payload["schema_version"] != "xnix.runtime.live_owner_gate.v1" ||
		payload["request_type"] != "runtime-live-owner-gate-preview" ||
		payload["gate_type"] != "runtime-live-owner-gate" ||
		payload["source"] != "runtime-service-binding-preview+owner-transition-gates" ||
		payload["runtime_method"] != "GetRuntimeLiveOwnerGate" ||
		payload["read_method"] != "GetRuntimeLiveOwnerGatePreview" {
		t.Fatalf("unexpected Runtime live owner gate CLI schema: %#v", payload)
	}
	if payload["bus_name"] != "org.xnix.Compatibility1" ||
		payload["object_path"] != "/org/xnix/Compatibility1" ||
		payload["interface"] != "org.xnix.Compatibility1" {
		t.Fatalf("unexpected Runtime D-Bus identity: %#v", payload)
	}
	serviceBinding := payload["service_binding"].(map[string]any)
	if serviceBinding["request_type"] != "runtime-service-binding-preview" ||
		serviceBinding["production_status"] != "pending-live-owner" ||
		serviceBinding["activation_binding_ready"] != true ||
		serviceBinding["live_dbus_owner_ready"] != false ||
		serviceBinding["smoke_adapter_available"] != true {
		t.Fatalf("unexpected service binding summary: %#v", serviceBinding)
	}
	serviceBindingCounts := serviceBinding["counts"].(map[string]any)
	if serviceBindingCounts["passed"] != float64(4) ||
		serviceBindingCounts["pending"] != float64(1) ||
		serviceBindingCounts["blocked"] != float64(0) {
		t.Fatalf("unexpected service binding counts: %#v", serviceBindingCounts)
	}
	gates := payload["required_gates"].([]any)
	gateIDs := payload["gate_ids"].([]any)
	expectedIDs := []string{
		"activation-binding",
		"long-running-runtime-owner",
		"bus-name-acquisition",
		"read-only-method-parity",
		"production-recipe-trust",
	}
	expectedStatuses := []string{"pass", "pending", "pending", "pending", "pending"}
	if len(gates) != len(expectedIDs) || len(gateIDs) != len(expectedIDs) {
		t.Fatalf("unexpected live owner gates: %#v ids=%#v", gates, gateIDs)
	}
	for index, id := range expectedIDs {
		gate := gates[index].(map[string]any)
		if gate["id"] != id || gateIDs[index] != id || gate["status"] != expectedStatuses[index] {
			t.Fatalf("unexpected gate at %d: %#v ids=%#v", index, gates, gateIDs)
		}
	}
	counts := payload["counts"].(map[string]any)
	if counts["total"] != float64(5) ||
		counts["passed"] != float64(1) ||
		counts["pending"] != float64(4) ||
		counts["blocked"] != float64(0) {
		t.Fatalf("unexpected live owner gate counts: %#v", counts)
	}
	if payload["runtime_owned"] != true ||
		payload["go_runtime_backed"] != true ||
		payload["kde_policy_owner"] != false ||
		payload["activation_binding_ready"] != true ||
		payload["live_dbus_owner_ready"] != false ||
		payload["production_owner_enabled"] != false ||
		payload["owner_transition_ready"] != false ||
		payload["smoke_adapter_available"] != true ||
		payload["smoke_adapter_is_production_owner"] != false ||
		payload["kde_may_claim_runtime_ownership"] != false ||
		payload["network_required"] != false ||
		payload["host_root_modified"] != false ||
		payload["privileged_container_required"] != false ||
		payload["backend_details_exposed"] != false {
		t.Fatalf("unexpected live owner gate safety flags: %#v", payload)
	}
	blockedReasons := payload["blocked_reasons"].([]any)
	blockedActions := payload["blocked_actions"].([]any)
	if len(blockedReasons) != 5 ||
		len(blockedActions) != 6 ||
		blockedActions[0] != "start production Runtime owner from preview" ||
		blockedActions[3] != "let KDE claim Runtime ownership" {
		t.Fatalf("unexpected live owner gate blockers: reasons=%#v actions=%#v", blockedReasons, blockedActions)
	}
	if strings.Contains(strings.ToLower(output.String()), "prefix") ||
		strings.Contains(strings.ToLower(output.String()), ".exe") ||
		strings.Contains(strings.ToLower(output.String()), "program files") {
		t.Fatalf("Runtime live owner gate CLI exposed backend terms: %s", output.String())
	}
}
