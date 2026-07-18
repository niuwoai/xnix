package main

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"
)

func TestKDETestLaunchMaterializationOwnerRouteAuditPreviewCommand(t *testing.T) {
	root := projectRootForRuntimeServiceBindingCommandTest(t)
	var output bytes.Buffer
	if err := run([]string{"kde-test-launch-materialization-owner-route-audit-preview", "--root", root}, &output); err != nil {
		t.Fatalf("kde-test-launch-materialization-owner-route-audit-preview returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["version"] != currentProjectVersion(t) ||
		payload["schema_version"] != "xnix.runtime.kde_test_launch_materialization_owner_route_audit.v1" ||
		payload["request_type"] != "kde-test-launch-materialization-owner-route-audit-preview" ||
		payload["runtime_method"] != "GetKDETestLaunchMaterializationOwnerRouteAudit" ||
		payload["read_method"] != "GetKDETestLaunchMaterializationOwnerRouteAuditPreview" ||
		payload["subject_command"] != "kde-test-launch-materialization-fanout-preview" ||
		payload["proposed_owner_method"] != "GetKDETestLaunchMaterializationFanOut" ||
		payload["route_decision"] != "consume-ready-opaque-lookup-missing" ||
		payload["current_route_status"] != "read-only-consume-ready-owner-route-blocked" {
		t.Fatalf("unexpected materialization owner-route audit command payload: %s", output.String())
	}
	if payload["cli_command_registered"] != true ||
		payload["go_read_model_present"] != true ||
		payload["owner_dispatch_route_present"] != false ||
		payload["production_dbus_method_present"] != false ||
		payload["requires_caller_registry_path"] != true ||
		payload["requires_caller_application_id"] != true ||
		payload["requires_caller_state_root"] != true ||
		payload["requires_explicit_authorization"] != true ||
		payload["materialization_writes_state_root"] != true ||
		payload["read_only_consume_command_registered"] != true ||
		payload["read_only_receipt_consumption_ready"] != true ||
		payload["read_only_consume_requires_caller_state_root"] != true ||
		payload["owner_managed_opaque_receipt_lookup_ready"] != false ||
		payload["opaque_materialization_receipt_id_supported"] != false ||
		payload["fan_out_writes_enabled"] != false ||
		payload["owner_local_route_candidate_ready"] != false ||
		payload["production_dbus_exposure_ready"] != false {
		t.Fatalf("unexpected materialization owner-route audit command decision: %s", output.String())
	}
	counts := payload["counts"].(map[string]any)
	if counts["total"] != float64(10) ||
		counts["passed"] != float64(8) ||
		counts["pending"] != float64(2) ||
		counts["blocked"] != float64(0) {
		t.Fatalf("unexpected materialization owner-route audit counts: %s", output.String())
	}
	if strings.Contains(output.String(), root) {
		t.Fatalf("audit command output must not expose project root path: %s", output.String())
	}
	for _, key := range []string{"system_service_started", "session_bus_claimed", "production_bus_claimed", "write_methods_enabled", "runtime_writes_enabled", "backend_launch_enabled", "backend_process_started", "network_required", "host_root_modified", "privileged_container_required", "state_root_path_exposed", "raw_command_exposed", "raw_executable_exposed", "backend_details_exposed"} {
		if payload[key] != false {
			t.Fatalf("unsafe gate %s must remain false: %s", key, output.String())
		}
	}
}

func TestKDETestLaunchMaterializationOwnerRouteAuditPreviewCommandRejectsPositionalArgs(t *testing.T) {
	var output bytes.Buffer
	if err := run([]string{"kde-test-launch-materialization-owner-route-audit-preview", "extra"}, &output); err == nil {
		t.Fatalf("kde-test-launch-materialization-owner-route-audit-preview must reject positional arguments")
	}
}
