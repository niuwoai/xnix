package main

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"
)

func TestCurrentMainlineOwnershipAuditPreviewCommand(t *testing.T) {
	root := projectRootForRuntimeServiceBindingCommandTest(t)
	var output bytes.Buffer
	if err := run([]string{"current-mainline-ownership-audit-preview", "--root", root}, &output); err != nil {
		t.Fatalf("current-mainline-ownership-audit-preview returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["version"] != currentProjectVersion(t) ||
		payload["schema_version"] != "xnix.runtime.current_mainline_ownership_audit.v1" ||
		payload["request_type"] != "current-mainline-ownership-audit-preview" ||
		payload["audit_type"] != "current-mainline-ownership-audit" ||
		payload["audit_decision"] != "current-mainline-ownership-audit-ready-autonomous-mainline" {
		t.Fatalf("unexpected current mainline ownership payload: %s", output.String())
	}
	if payload["current_mainline_document_present"] != true ||
		payload["codex_owned_mainline"] != true ||
		payload["external_agent_dependency_required"] != false ||
		payload["historical_external_agent_documents_allowed"] != true ||
		payload["kde_flagship_only"] != true ||
		payload["gnome_first_release_target"] != false ||
		payload["xfce_first_release_target"] != false ||
		payload["go_runtime_owned"] != true ||
		payload["ruby_test_harness_only"] != true ||
		payload["c_low_level_only"] != true ||
		payload["kde_plugins_presentation_only"] != true ||
		payload["runtime_owns_compatibility_decisions"] != true ||
		payload["seven_kde_entrypoints_present"] != true ||
		payload["safe_next_task_present"] != true ||
		payload["full_gate_requires_authorization"] != true {
		t.Fatalf("unexpected current mainline ownership decision: %s", output.String())
	}
	counts := payload["counts"].(map[string]any)
	if counts["total"] != float64(8) || counts["passed"] != float64(8) || counts["blocked"] != float64(0) {
		t.Fatalf("unexpected current mainline ownership checks: %s", output.String())
	}
	if strings.Contains(output.String(), root) {
		t.Fatalf("current mainline ownership output must not expose project root path: %s", output.String())
	}
	for _, key := range []string{"docker_executed", "qemu_executed", "network_fetch_enabled", "package_manager_invoked", "production_bus_claimed", "write_methods_enabled", "runtime_writes_enabled", "kde_configuration_written", "portal_call_executed", "backend_launch_enabled", "backend_process_started", "host_root_modified", "privileged_container_required", "state_root_path_exposed", "file_paths_exposed", "file_content_read", "raw_command_exposed", "raw_executable_exposed", "backend_details_exposed"} {
		if payload[key] != false {
			t.Fatalf("unsafe current mainline ownership %s must remain false: %s", key, output.String())
		}
	}
}

func TestCurrentMainlineOwnershipAuditPreviewCommandRejectsPositionalArgs(t *testing.T) {
	var output bytes.Buffer
	if err := run([]string{"current-mainline-ownership-audit-preview", "extra"}, &output); err == nil {
		t.Fatalf("current-mainline-ownership-audit-preview must reject positional arguments")
	}
}
