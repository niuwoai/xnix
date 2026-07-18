package main

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"
)

func TestBackendAdapterRedactedProfileAuditPreviewCommand(t *testing.T) {
	root := projectRootForRuntimeServiceBindingCommandTest(t)
	var output bytes.Buffer
	if err := run([]string{"backend-adapter-redacted-profile-audit-preview", "--root", root}, &output); err != nil {
		t.Fatalf("backend-adapter-redacted-profile-audit-preview returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["version"] != currentProjectVersion(t) ||
		payload["schema_version"] != "xnix.runtime.backend_adapter_redacted_profile_audit.v1" ||
		payload["request_type"] != "backend-adapter-redacted-profile-audit-preview" ||
		payload["audit_type"] != "owner-local-redacted-adapter-profile-audit" ||
		payload["runtime_method"] != "GetBackendAdapterProfileAudit" ||
		payload["read_method"] != "GetBackendAdapterProfileAuditPreview" ||
		payload["subject_request_type"] != "backend-adapter-contract-preview" ||
		payload["subject_read_method"] != "GetBackendAdapterContractPreview" ||
		payload["owner_route_method"] != "GetBackendAdapterProfileAudit" ||
		payload["route_decision"] != "redacted-profile-route-ready" {
		t.Fatalf("unexpected redacted adapter profile audit payload: %s", output.String())
	}
	if payload["full_contract_consumed"] != true ||
		payload["kde_facing_projection_consumed"] != true ||
		payload["internal_adapter_ids_redacted"] != true ||
		payload["internal_profile_paths_redacted"] != true ||
		payload["owner_local_route_candidate_ready"] != true ||
		payload["full_contract_fixture_local"] != true ||
		payload["production_dbus_exposure_ready"] != false ||
		payload["caller_state_root_required"] != false ||
		payload["cli_project_root_only"] != true {
		t.Fatalf("unexpected redacted adapter profile audit decision: %s", output.String())
	}
	counts := payload["counts"].(map[string]any)
	if counts["total"] != float64(7) ||
		counts["passed"] != float64(7) ||
		counts["blocked"] != float64(0) ||
		counts["redacted_profiles"] != float64(3) ||
		counts["user_visible_profiles"] != float64(3) ||
		counts["enabled_invocations"] != float64(0) ||
		counts["enabled_launches"] != float64(0) ||
		counts["exposed_backend_details"] != float64(0) {
		t.Fatalf("unexpected redacted adapter profile audit counts: %s", output.String())
	}
	for _, key := range []string{"system_service_started", "session_bus_claimed", "production_bus_claimed", "write_methods_enabled", "runtime_writes_enabled", "adapter_invocation_enabled", "backend_install_enabled", "backend_download_enabled", "backend_launch_enabled", "backend_process_started", "command_materialized", "executable_path_resolved", "network_required", "host_root_modified", "privileged_container_required", "state_root_path_exposed", "raw_command_exposed", "raw_executable_exposed", "backend_details_exposed"} {
		if payload[key] != false {
			t.Fatalf("unsafe gate %s must remain false: %s", key, output.String())
		}
	}
	if strings.Contains(output.String(), root) {
		t.Fatalf("redacted adapter profile audit command output must not expose project root path: %s", output.String())
	}
	for _, forbidden := range []string{"wine", "proton", "windows-vm", "prefix", ".exe", "program files", "qemu-system", ".wine", "/tmp", "/users", "/private", "file://"} {
		if strings.Contains(strings.ToLower(output.String()), forbidden) {
			t.Fatalf("redacted adapter profile audit command must not expose forbidden term %q: %s", forbidden, output.String())
		}
	}
}

func TestBackendAdapterRedactedProfileAuditPreviewCommandRejectsPositionalArgs(t *testing.T) {
	var output bytes.Buffer
	if err := run([]string{"backend-adapter-redacted-profile-audit-preview", "extra"}, &output); err == nil {
		t.Fatalf("backend-adapter-redacted-profile-audit-preview must reject positional arguments")
	}
}
