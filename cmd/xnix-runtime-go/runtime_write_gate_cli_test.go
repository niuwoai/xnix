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
	if payload["version"] != "0.2.197" ||
		payload["schema_version"] != "xnix.runtime.write_gate.v1" ||
		payload["request_type"] != "runtime-write-gate-preview" ||
		payload["gate_type"] != "runtime-write-gate" ||
		payload["source"] != "go-runtime-write-gate" ||
		payload["runtime_method"] != "GetRuntimeWriteGate" ||
		payload["read_method"] != "GetRuntimeWriteGatePreview" ||
		payload["method_name"] != "Launch" {
		t.Fatalf("unexpected Runtime write gate CLI schema: %#v", payload)
	}
	if payload["runtime_owned"] != true ||
		payload["go_runtime_backed"] != true ||
		payload["kde_policy_owner"] != false ||
		payload["gate_decision"] != "blocked-until-production-backend" ||
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
	if payload["denial_error_name"] != "org.xnix.Compatibility1.Error.WriteMethodDisabled" {
		t.Fatalf("unexpected Runtime write gate denial: %#v", payload)
	}
	if len(payload["required_gates"].([]any)) != 6 ||
		len(payload["supported_write_methods"].([]any)) != 4 {
		t.Fatalf("unexpected Runtime write gate requirements: %#v", payload)
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
