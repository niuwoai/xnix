package main

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"
)

func TestBackendAdapterContractOwnerRouteAuditPreviewCommand(t *testing.T) {
	root := projectRootForRuntimeServiceBindingCommandTest(t)
	var output bytes.Buffer
	if err := run([]string{"backend-adapter-contract-owner-route-audit-preview", "--root", root}, &output); err != nil {
		t.Fatalf("backend-adapter-contract-owner-route-audit-preview returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["version"] != currentProjectVersion(t) ||
		payload["schema_version"] != "xnix.runtime.backend_adapter_contract_owner_route_audit.v1" ||
		payload["request_type"] != "backend-adapter-contract-owner-route-audit-preview" ||
		payload["runtime_method"] != "GetBackendAdapterContractOwnerRouteAudit" ||
		payload["read_method"] != "GetBackendAdapterContractOwnerRouteAuditPreview" ||
		payload["subject_command"] != "backend-adapter-contract-preview" ||
		payload["proposed_owner_method"] != "GetBackendAdapterProfileAudit" ||
		payload["route_decision"] != "redacted-profile-route-ready" ||
		payload["current_route_status"] != "redacted-profile-route-ready-full-contract-fixture-local" {
		t.Fatalf("unexpected adapter contract owner-route audit command payload: %s", output.String())
	}
	if payload["cli_command_registered"] != true ||
		payload["go_read_model_present"] != true ||
		payload["fixture_matrix_consumes_contract"] != true ||
		payload["owner_dispatch_route_present"] != false ||
		payload["production_dbus_method_present"] != false ||
		payload["full_contract_contains_adapter_ids"] != true ||
		payload["kde_facing_projection_present"] != true ||
		payload["redacted_profile_route_present"] != true ||
		payload["requires_caller_root"] != true ||
		payload["adapter_invocation_enabled"] != false ||
		payload["backend_launch_enabled"] != false ||
		payload["command_materialized"] != false ||
		payload["owner_local_route_candidate_ready"] != true ||
		payload["production_dbus_exposure_ready"] != false {
		t.Fatalf("unexpected adapter contract owner-route audit command decision: %s", output.String())
	}
	counts := payload["counts"].(map[string]any)
	if counts["total"] != float64(9) ||
		counts["passed"] != float64(8) ||
		counts["pending"] != float64(1) ||
		counts["blocked"] != float64(0) {
		t.Fatalf("unexpected adapter contract owner-route audit counts: %s", output.String())
	}
	if strings.Contains(output.String(), root) {
		t.Fatalf("audit command output must not expose project root path: %s", output.String())
	}
	for _, key := range []string{"system_service_started", "session_bus_claimed", "production_bus_claimed", "write_methods_enabled", "runtime_writes_enabled", "network_required", "host_root_modified", "privileged_container_required", "state_root_path_exposed", "raw_command_exposed", "raw_executable_exposed", "backend_details_exposed"} {
		if payload[key] != false {
			t.Fatalf("unsafe gate %s must remain false: %s", key, output.String())
		}
	}
	for _, forbidden := range []string{"wine", "proton", "windows-vm"} {
		if strings.Contains(strings.ToLower(output.String()), forbidden) {
			t.Fatalf("audit command must not expose internal adapter id %q: %s", forbidden, output.String())
		}
	}
}

func TestBackendAdapterContractOwnerRouteAuditPreviewCommandRejectsPositionalArgs(t *testing.T) {
	var output bytes.Buffer
	if err := run([]string{"backend-adapter-contract-owner-route-audit-preview", "extra"}, &output); err == nil {
		t.Fatalf("backend-adapter-contract-owner-route-audit-preview must reject positional arguments")
	}
}
