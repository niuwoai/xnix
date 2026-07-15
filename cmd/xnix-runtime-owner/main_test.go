package main

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"
)

func TestRuntimeOwnerCommandRendersSmokeCandidate(t *testing.T) {
	var output bytes.Buffer
	if err := run([]string{"--root", "../..", "--mode", "smoke-owner"}, &output); err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["version"] != "0.2.218" ||
		payload["schema_version"] != "xnix.runtime.owner_candidate.v1" ||
		payload["request_type"] != "runtime-owner-candidate" ||
		payload["owner_type"] != "go-runtime-owner-candidate" ||
		payload["mode"] != "smoke-owner" {
		t.Fatalf("unexpected owner candidate schema: %#v", payload)
	}
	if payload["bus_name"] != "org.xnix.Compatibility1" ||
		payload["read_only_serve_ready"] != true ||
		payload["route_count"] != float64(57) ||
		payload["go_route_count"] != float64(57) ||
		payload["c_core_route_count"] != float64(0) ||
		payload["ruby_legacy_route_count"] != float64(0) ||
		payload["write_method_count"] != float64(4) {
		t.Fatalf("unexpected owner candidate readiness: %#v", payload)
	}
	if payload["runtime_owned"] != true ||
		payload["go_runtime_backed"] != true ||
		payload["kde_policy_owner"] != false ||
		payload["smoke_owner_mode"] != true ||
		payload["production_owner_mode"] != false ||
		payload["event_loop_started"] != false ||
		payload["session_bus_claimed"] != false ||
		payload["production_bus_claimed"] != false ||
		payload["write_methods_enabled"] != false ||
		payload["host_root_modified"] != false ||
		payload["backend_details_exposed"] != false {
		t.Fatalf("unexpected owner candidate safety flags: %#v", payload)
	}
}

func TestRuntimeOwnerCommandRendersDisabledWrite(t *testing.T) {
	var output bytes.Buffer
	if err := run([]string{"--deny-write", "Launch"}, &output); err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["method"] != "Launch" ||
		payload["error_name"] != "org.xnix.Compatibility1.Error.WriteMethodDisabled" ||
		payload["dispatch_enabled"] != false ||
		payload["request_created"] != false {
		t.Fatalf("unexpected disabled write response: %#v", payload)
	}
}

func TestRuntimeOwnerCommandRendersReadDispatch(t *testing.T) {
	var output bytes.Buffer
	if err := run([]string{"--root", "../..", "--dispatch-read", "GetRuntimeWriteGate", "Launch"}, &output); err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["version"] != "0.2.218" ||
		payload["schema_version"] != "xnix.runtime.owner_read_dispatch.v1" ||
		payload["request_type"] != "runtime-owner-read-dispatch" ||
		payload["dispatch_type"] != "go-owner-read-dispatch" ||
		payload["method"] != "GetRuntimeWriteGate" ||
		payload["go_command"] != "runtime-write-gate-preview" ||
		payload["read_only_dispatch"] != true ||
		payload["event_loop_started"] != false ||
		payload["session_bus_claimed"] != false ||
		payload["production_bus_claimed"] != false ||
		payload["host_root_modified"] != false {
		t.Fatalf("unexpected read dispatch response: %#v", payload)
	}
	nested := payload["payload"].(map[string]any)
	if nested["request_type"] != "runtime-write-gate-preview" ||
		nested["method_name"] != "Launch" ||
		nested["dispatch_enabled"] != false {
		t.Fatalf("unexpected read dispatch payload: %#v", nested)
	}
}

func TestRuntimeOwnerCommandRendersLifecycleJSONL(t *testing.T) {
	var output bytes.Buffer
	if err := run([]string{"--root", "../..", "--mode", "smoke-owner", "--lifecycle-log"}, &output); err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	lines := strings.Split(strings.TrimSpace(output.String()), "\n")
	if len(lines) != 4 {
		t.Fatalf("lifecycle line count = %d, want 4: %q", len(lines), output.String())
	}
	wantTypes := []string{"startup", "route-table", "readiness", "shutdown"}
	for index, line := range lines {
		var payload map[string]any
		if err := json.Unmarshal([]byte(line), &payload); err != nil {
			t.Fatalf("Unmarshal line %d returned error: %v", index, err)
		}
		if payload["version"] != "0.2.218" ||
			payload["schema_version"] != "xnix.runtime.owner_lifecycle_event.v1" ||
			payload["request_type"] != "runtime-owner-lifecycle-event" ||
			payload["event_type"] != wantTypes[index] ||
			payload["sequence"] != float64(index+1) ||
			payload["mode"] != "smoke-owner" ||
			payload["route_count"] != float64(57) ||
			payload["go_route_count"] != float64(57) ||
			payload["write_method_count"] != float64(4) {
			t.Fatalf("unexpected lifecycle event %d: %#v", index, payload)
		}
		if payload["event_loop_started"] != false ||
			payload["session_bus_claimed"] != false ||
			payload["production_bus_claimed"] != false ||
			payload["write_methods_enabled"] != false ||
			payload["host_root_modified"] != false ||
			payload["backend_details_exposed"] != false {
			t.Fatalf("unexpected lifecycle safety flags at %d: %#v", index, payload)
		}
	}
}

func TestRuntimeOwnerCommandRejectsMixedOperations(t *testing.T) {
	var output bytes.Buffer
	err := run([]string{"--deny-write", "Launch", "--lifecycle-log"}, &output)
	if err == nil || !strings.Contains(err.Error(), "accepts only one") {
		t.Fatalf("run must reject mixed lifecycle/write operations, got %v", err)
	}
}
