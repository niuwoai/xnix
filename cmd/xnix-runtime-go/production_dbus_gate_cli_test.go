package main

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"
)

func TestProductionDBusGateReviewPreviewCommand(t *testing.T) {
	root := projectRootForRuntimeServiceBindingCommandTest(t)
	var output bytes.Buffer
	if err := run([]string{"production-dbus-gate-review-preview", "--root", root}, &output); err != nil {
		t.Fatalf("production-dbus-gate-review-preview returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["version"] != currentProjectVersion(t) ||
		payload["schema_version"] != "xnix.runtime.production_dbus_gate_review.v1" ||
		payload["request_type"] != "production-dbus-gate-review-preview" ||
		payload["gate_type"] != "owner-local-smoke-covered-production-dbus-gate-review" ||
		payload["gate_decision"] != "production-dbus-gate-review-ready" ||
		payload["current_gate_status"] != "owner-local-smoke-covered-production-dbus-disabled" {
		t.Fatalf("unexpected production D-Bus gate review command payload: %s", output.String())
	}
	if payload["route_count"] != float64(3) ||
		payload["smoke_covered_route_count"] != float64(3) ||
		payload["production_readiness"] != false ||
		payload["human_authorization_required"] != true ||
		payload["production_owner_enabled"] != false ||
		payload["production_activation_ready"] != false {
		t.Fatalf("unexpected production D-Bus gate review command decision: %s", output.String())
	}
	routes := payload["routes"].([]any)
	if len(routes) != 3 {
		t.Fatalf("unexpected production D-Bus gate route count: %s", output.String())
	}
	for _, item := range routes {
		route := item.(map[string]any)
		if route["smoke_coverage_ready"] != true ||
			route["owner_local_route_ready"] != true ||
			route["production_dbus_method_present"] != false ||
			route["production_dbus_exposure_ready"] != false ||
			route["write_methods_enabled"] != false ||
			route["runtime_writes_enabled"] != false ||
			route["backend_launch_enabled"] != false ||
			route["host_root_modified"] != false ||
			route["internal_details_exposed"] != false {
			t.Fatalf("unsafe production D-Bus gate route: %s", output.String())
		}
	}
	counts := payload["counts"].(map[string]any)
	if counts["total"] != float64(8) ||
		counts["passed"] != float64(8) ||
		counts["pending"] != float64(0) ||
		counts["blocked"] != float64(0) {
		t.Fatalf("unexpected production D-Bus gate review counts: %s", output.String())
	}
	if strings.Contains(output.String(), root) {
		t.Fatalf("production D-Bus gate review output must not expose project root path: %s", output.String())
	}
	for _, key := range []string{"system_service_started", "session_bus_claimed", "production_bus_claimed", "write_methods_enabled", "runtime_writes_enabled", "adapter_invocation_enabled", "backend_launch_enabled", "backend_process_started", "support_bundle_exported", "support_case_created", "notification_sent", "network_required", "host_root_modified", "privileged_container_required", "state_root_path_exposed", "raw_command_exposed", "raw_executable_exposed", "backend_details_exposed"} {
		if payload[key] != false {
			t.Fatalf("unsafe gate %s must remain false: %s", key, output.String())
		}
	}
}

func TestProductionDBusGateReviewPreviewCommandRejectsPositionalArgs(t *testing.T) {
	var output bytes.Buffer
	if err := run([]string{"production-dbus-gate-review-preview", "extra"}, &output); err == nil {
		t.Fatalf("production-dbus-gate-review-preview must reject positional arguments")
	}
}
