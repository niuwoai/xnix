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
	if payload["version"] != currentProjectVersion(t) ||
		payload["schema_version"] != "xnix.runtime.owner_candidate.v1" ||
		payload["request_type"] != "runtime-owner-candidate" ||
		payload["owner_type"] != "go-runtime-owner-candidate" ||
		payload["mode"] != "smoke-owner" {
		t.Fatalf("unexpected owner candidate schema: %#v", payload)
	}
	if payload["bus_name"] != "org.xnix.Compatibility1" ||
		payload["read_only_serve_ready"] != true ||
		payload["route_count"] != float64(61) ||
		payload["go_route_count"] != float64(61) ||
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
	if payload["version"] != currentProjectVersion(t) ||
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

func TestRuntimeOwnerCommandRendersWindowsCompatibilityOwnerLocalReadDispatch(t *testing.T) {
	var output bytes.Buffer
	if err := run([]string{"--root", "../..", "--dispatch-read", "GetWindowsCompatibilityWorkstreamsPreview"}, &output); err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["schema_version"] != "xnix.runtime.owner_read_dispatch.v1" ||
		payload["method"] != "GetWindowsCompatibilityWorkstreamsPreview" ||
		payload["route_source"] != "go-owner-local-preview" ||
		payload["go_command"] != "windows-compatibility-workstreams-preview" ||
		payload["route_status"] != "owner-local-preview-ready" ||
		payload["read_only_dispatch"] != true ||
		payload["write_methods_enabled"] != false ||
		payload["session_bus_claimed"] != false ||
		payload["production_bus_claimed"] != false ||
		payload["host_root_modified"] != false ||
		payload["backend_details_exposed"] != false {
		t.Fatalf("unexpected Windows compatibility read dispatch response: %#v", payload)
	}
	nested := payload["payload"].(map[string]any)
	if nested["request_type"] != "windows-compatibility-workstreams-preview" ||
		nested["read_method"] != "GetWindowsCompatibilityWorkstreamsPreview" ||
		nested["official_desktop"] != "KDE Plasma" ||
		nested["runtime_owned"] != true ||
		nested["go_runtime_backed"] != true ||
		nested["backend_launch_enabled"] != false ||
		nested["host_root_modified"] != false {
		t.Fatalf("unexpected Windows compatibility nested payload: %#v", nested)
	}
}

func TestRuntimeOwnerCommandRendersNotificationDigestOwnerLocalReadDispatch(t *testing.T) {
	var output bytes.Buffer
	if err := run([]string{
		"--root", "../..",
		"--dispatch-read", "GetKDENotificationDigestPreview",
		"org.xnix.sample.notepad",
		"permission-attention:permission-review",
	}, &output); err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["method"] != "GetKDENotificationDigestPreview" ||
		payload["route_source"] != "go-owner-local-preview" ||
		payload["go_command"] != "kde-notification-digest-preview" ||
		payload["route_status"] != "owner-local-preview-ready" ||
		payload["read_only_dispatch"] != true ||
		payload["write_methods_enabled"] != false ||
		payload["production_bus_claimed"] != false ||
		payload["host_root_modified"] != false {
		t.Fatalf("unexpected notification digest read dispatch response: %#v", payload)
	}
	nested := payload["payload"].(map[string]any)
	if nested["request_type"] != "kde-notification-digest-preview" ||
		nested["read_method"] != "GetKDENotificationDigestPreview" ||
		nested["notifications_sent"] != false ||
		nested["backend_launch_enabled"] != false ||
		nested["host_root_modified"] != false {
		t.Fatalf("unexpected notification digest nested payload: %#v", nested)
	}
}

func TestRuntimeOwnerCommandRendersSignedRecipeOwnerLocalReadDispatch(t *testing.T) {
	var output bytes.Buffer
	if err := run([]string{
		"--root", "../..",
		"--service-call", "GetSignedRecipeVerificationPreview",
		"org.xnix.sample.notepad",
	}, &output); err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["method"] != "GetSignedRecipeVerificationPreview" ||
		payload["call_type"] != "read-dispatch" ||
		payload["read_only_dispatch"] != true ||
		payload["write_methods_enabled"] != false ||
		payload["production_bus_claimed"] != false ||
		payload["network_required"] != false ||
		payload["host_root_modified"] != false {
		t.Fatalf("unexpected signed recipe service response: %#v", payload)
	}
	dispatch := payload["payload"].(map[string]any)
	nested := dispatch["payload"].(map[string]any)
	if dispatch["go_command"] != "signed-recipe-verifier-preview" ||
		nested["verification_state"] != "production-signature-required" ||
		nested["recipe_digest_verified"] != true ||
		nested["signature_verified"] != false ||
		nested["signature_material_exposed"] != false {
		t.Fatalf("unexpected signed recipe nested payload: %#v", nested)
	}
}

func TestRuntimeOwnerCommandRendersServiceCall(t *testing.T) {
	var output bytes.Buffer
	if err := run([]string{"--root", "../..", "--mode", "smoke-owner", "--service-call", "GetRuntimeWriteGate", "Launch"}, &output); err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["version"] != currentProjectVersion(t) ||
		payload["schema_version"] != "xnix.runtime.owner_service_call.v1" ||
		payload["request_type"] != "runtime-owner-service-call" ||
		payload["service_type"] != "go-runtime-owner-in-process-service" ||
		payload["method"] != "GetRuntimeWriteGate" ||
		payload["call_type"] != "read-dispatch" ||
		payload["read_only_dispatch"] != true ||
		payload["write_method"] != false ||
		payload["dispatch_ready"] != true ||
		payload["in_process_service_ready"] != true ||
		payload["event_loop_started"] != false ||
		payload["session_bus_claimed"] != false ||
		payload["production_bus_claimed"] != false ||
		payload["host_root_modified"] != false {
		t.Fatalf("unexpected service call response: %#v", payload)
	}
	nested := payload["payload"].(map[string]any)
	if nested["schema_version"] != "xnix.runtime.owner_read_dispatch.v1" ||
		nested["method"] != "GetRuntimeWriteGate" {
		t.Fatalf("unexpected nested service call payload: %#v", nested)
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
		if payload["version"] != currentProjectVersion(t) ||
			payload["schema_version"] != "xnix.runtime.owner_lifecycle_event.v1" ||
			payload["request_type"] != "runtime-owner-lifecycle-event" ||
			payload["event_type"] != wantTypes[index] ||
			payload["sequence"] != float64(index+1) ||
			payload["mode"] != "smoke-owner" ||
			payload["route_count"] != float64(61) ||
			payload["go_route_count"] != float64(61) ||
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

func TestRuntimeOwnerCommandRendersSmokeBatchJSONL(t *testing.T) {
	var output bytes.Buffer
	if err := run([]string{"--root", "../..", "--smoke-batch"}, &output); err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	lines := strings.Split(strings.TrimSpace(output.String()), "\n")
	if len(lines) != 68 {
		t.Fatalf("smoke batch line count = %d, want 68", len(lines))
	}
	readCount := 0
	writeCount := 0
	for index, line := range lines {
		var payload map[string]any
		if err := json.Unmarshal([]byte(line), &payload); err != nil {
			t.Fatalf("Unmarshal line %d returned error: %v", index, err)
		}
		if payload["version"] != currentProjectVersion(t) ||
			payload["schema_version"] != "xnix.runtime.owner_smoke_batch.v1" ||
			payload["request_type"] != "runtime-owner-smoke-batch-record" ||
			payload["batch_type"] != "restricted-session-owner-call-batch" ||
			payload["sequence"] != float64(index+1) ||
			payload["read_dispatch_method_count"] != float64(64) ||
			payload["write_method_count"] != float64(4) ||
			payload["runtime_owned"] != true ||
			payload["go_runtime_backed"] != true ||
			payload["event_loop_started"] != false ||
			payload["session_bus_claimed"] != false ||
			payload["production_bus_claimed"] != false ||
			payload["host_root_modified"] != false ||
			payload["backend_details_exposed"] != false {
			t.Fatalf("unexpected smoke batch record %d: %#v", index, payload)
		}
		switch payload["record_type"] {
		case "read-dispatch":
			readCount++
			if payload["read_only_dispatch"] != true ||
				payload["write_method"] != false ||
				payload["dispatch_ready"] != true {
				t.Fatalf("unexpected read smoke batch record %d: %#v", index, payload)
			}
			nested := payload["payload"].(map[string]any)
			if nested["request_type"] != "runtime-owner-service-call" ||
				nested["call_type"] != "read-dispatch" ||
				nested["read_only_dispatch"] != true {
				t.Fatalf("unexpected read smoke batch payload %d: %#v", index, nested)
			}
		case "write-denial":
			writeCount++
			if payload["read_only_dispatch"] != false ||
				payload["write_method"] != true ||
				payload["dispatch_ready"] != true ||
				payload["error_name"] != "org.xnix.Compatibility1.Error.WriteMethodDisabled" {
				t.Fatalf("unexpected write smoke batch record %d: %#v", index, payload)
			}
			nested := payload["payload"].(map[string]any)
			if nested["request_type"] != "runtime-owner-service-call" ||
				nested["call_type"] != "write-denial" ||
				nested["write_method"] != true {
				t.Fatalf("unexpected write smoke batch payload %d: %#v", index, nested)
			}
		default:
			t.Fatalf("unexpected smoke batch record type at %d: %#v", index, payload)
		}
	}
	if readCount != 64 || writeCount != 4 {
		t.Fatalf("smoke batch counts read=%d write=%d, want 64/4", readCount, writeCount)
	}
}

func TestRuntimeOwnerCommandRendersSessionBusSmokeJSONL(t *testing.T) {
	var output bytes.Buffer
	if err := run([]string{"--root", "../..", "--session-bus-smoke"}, &output); err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	lines := strings.Split(strings.TrimSpace(output.String()), "\n")
	if len(lines) != 73 {
		t.Fatalf("session bus smoke line count = %d, want 73", len(lines))
	}
	readCount := 0
	writeCount := 0
	unsupportedCount := 0
	for index, line := range lines {
		var payload map[string]any
		if err := json.Unmarshal([]byte(line), &payload); err != nil {
			t.Fatalf("Unmarshal line %d returned error: %v", index, err)
		}
		if payload["version"] != currentProjectVersion(t) ||
			payload["schema_version"] != "xnix.runtime.owner_session_bus_smoke.v1" ||
			payload["request_type"] != "runtime-owner-session-bus-smoke-step" ||
			payload["transcript_type"] != "restricted-private-session-bus-owner-smoke" ||
			payload["sequence"] != float64(index+1) ||
			payload["read_dispatch_method_count"] != float64(64) ||
			payload["write_method_count"] != float64(4) ||
			payload["runtime_owned"] != true ||
			payload["go_runtime_backed"] != true ||
			payload["kde_policy_owner"] != false ||
			payload["private_session_bus"] != true ||
			payload["event_loop_started"] != true ||
			payload["session_bus_claimed"] != true ||
			payload["production_bus_claimed"] != false ||
			payload["system_service_started"] != false ||
			payload["write_methods_enabled"] != false ||
			payload["host_root_modified"] != false ||
			payload["backend_details_exposed"] != false {
			t.Fatalf("unexpected session bus smoke payload at %d: %#v", index, payload)
		}
		switch payload["step_type"] {
		case "read-dispatch":
			readCount++
			if payload["read_only_dispatch"] != true || payload["write_method"] != false {
				t.Fatalf("unexpected read step at %d: %#v", index, payload)
			}
		case "write-denial":
			writeCount++
			if payload["read_only_dispatch"] != false ||
				payload["write_method"] != true ||
				payload["error_name"] != "org.xnix.Compatibility1.Error.WriteMethodDisabled" {
				t.Fatalf("unexpected write step at %d: %#v", index, payload)
			}
		case "reject-unsupported-read":
			unsupportedCount++
			if payload["unsupported_read"] != true ||
				payload["error_name"] != "org.xnix.Compatibility1.Error.UnsupportedMethod" {
				t.Fatalf("unexpected unsupported-read step at %d: %#v", index, payload)
			}
		}
	}
	if readCount != 64 || writeCount != 4 || unsupportedCount != 1 {
		t.Fatalf("session bus smoke counts read=%d write=%d unsupported=%d, want 64/4/1", readCount, writeCount, unsupportedCount)
	}
}

func TestRuntimeOwnerCommandRejectsMixedOperations(t *testing.T) {
	var output bytes.Buffer
	err := run([]string{"--deny-write", "Launch", "--lifecycle-log"}, &output)
	if err == nil || !strings.Contains(err.Error(), "accepts only one") {
		t.Fatalf("run must reject mixed lifecycle/write operations, got %v", err)
	}
}
