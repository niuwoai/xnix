package main

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"
)

func TestRuntimeWriteGatePreviewCommandRendersGoGate(t *testing.T) {
	var output bytes.Buffer
	err := run([]string{"runtime-write-gate-preview", "--root", projectRootForRuntimeServiceBindingCommandTest(t), "--method", "Launch"}, &output)
	if err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["version"] != currentProjectVersion(t) ||
		payload["schema_version"] != "xnix.runtime.write_gate.v1" ||
		payload["request_type"] != "runtime-write-gate-preview" ||
		payload["gate_type"] != "runtime-write-gate" ||
		payload["source"] != "go-runtime-write-gate+runtime-service-activation-preflight-preview+production-dbus-gate-review-preview+production-dbus-human-authorization-preflight-preview+production-human-authorization-receipt-consolidation-preview" ||
		payload["runtime_method"] != "GetRuntimeWriteGate" ||
		payload["read_method"] != "GetRuntimeWriteGatePreview" ||
		payload["method_name"] != "Launch" {
		t.Fatalf("unexpected Runtime write gate CLI schema: %#v", payload)
	}
	if payload["runtime_owned"] != true ||
		payload["go_runtime_backed"] != true ||
		payload["kde_policy_owner"] != false ||
		payload["gate_decision"] != "blocked-until-production-backend" ||
		payload["production_gate_decision"] != "production-gates-consumed-write-gate-disabled" ||
		payload["write_method_enabled"] != false ||
		payload["dispatch_enabled"] != false ||
		payload["request_object_created"] != false ||
		payload["execution_started"] != false ||
		payload["network_required"] != false ||
		payload["host_root_modified"] != false ||
		payload["privileged_container_required"] != false ||
		payload["backend_details_exposed"] != false {
		t.Fatalf("unexpected Runtime write gate CLI safety flags: %#v", payload)
	}
	productionGate := payload["production_gate"].(map[string]any)
	if productionGate["request_type"] != "runtime-service-activation-preflight-preview" ||
		productionGate["preflight_type"] != "production-runtime-service-activation-preflight" ||
		productionGate["preflight_decision"] != "restricted-owner-smoke-ready" ||
		productionGate["service_activation_preflight_ready"] != true ||
		productionGate["production_dbus_gate_ready"] != true ||
		productionGate["human_authorization_preflight_ready"] != true ||
		productionGate["human_authorization_required"] != true ||
		productionGate["human_authorization_granted"] != false ||
		productionGate["authorization_receipt_accepted"] != false ||
		productionGate["production_activation_ready"] != false ||
		productionGate["restricted_smoke_ready"] != true ||
		productionGate["system_service_started"] != false ||
		productionGate["production_bus_claimed"] != false ||
		productionGate["write_methods_enabled"] != false ||
		productionGate["backend_launch_enabled"] != false ||
		productionGate["network_required"] != false ||
		productionGate["host_root_modified"] != false ||
		productionGate["privileged_container_required"] != false ||
		productionGate["backend_details_exposed"] != false {
		t.Fatalf("unexpected production gate summary: %#v", productionGate)
	}
	if payload["denial_error_name"] != "org.xnix.Compatibility1.Error.WriteMethodDisabled" {
		t.Fatalf("unexpected Runtime write gate denial: %#v", payload)
	}
	if len(payload["required_gates"].([]any)) != 10 ||
		len(payload["supported_write_methods"].([]any)) != 4 {
		t.Fatalf("unexpected Runtime write gate requirements: %#v", payload)
	}
	counts := payload["counts"].(map[string]any)
	if counts["total"] != float64(10) ||
		counts["passed"] != float64(3) ||
		counts["pending"] != float64(7) ||
		counts["blocked"] != float64(0) {
		t.Fatalf("unexpected Runtime write gate counts: %#v", counts)
	}

	serialized := strings.ToLower(output.String())
	for _, forbidden := range []string{"prefix", ".exe", "program files", "qemu-system", "proton", "wine "} {
		if strings.Contains(serialized, forbidden) {
			t.Fatalf("Runtime write gate CLI exposed forbidden term %q: %s", forbidden, output.String())
		}
	}
}

func TestRuntimeWriteGatePreviewCommandRejectsMissingMethod(t *testing.T) {
	var output bytes.Buffer
	if err := run([]string{"runtime-write-gate-preview", "--root", projectRootForRuntimeServiceBindingCommandTest(t)}, &output); err == nil {
		t.Fatal("runtime-write-gate-preview accepted a missing method")
	}
}
