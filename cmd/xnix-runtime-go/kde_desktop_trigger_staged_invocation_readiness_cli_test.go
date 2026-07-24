package main

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"
)

func TestDesktopTriggerStagedInvocationReadinessPreviewCommand(t *testing.T) {
	stateRoot := t.TempDir()
	record := recordKDEControlledLaunchActionCLIEvidence(t, stateRoot)

	var output bytes.Buffer
	err := run([]string{
		"desktop-trigger-staged-invocation-readiness-preview",
		"--state-root", stateRoot,
		"--evidence-relative-path", record.EvidenceRelativePath,
		"--expected-evidence-sha256", record.EvidenceSHA256,
		"--full-checkpoint-promoted",
	}, &output)
	if err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("readiness output must be JSON: %v\n%s", err, output.String())
	}
	if payload["schema_version"] != "xnix.runtime.desktop_trigger_staged_invocation_readiness.v1" ||
		payload["request_type"] != "desktop-trigger-staged-invocation-readiness-preview" ||
		payload["readiness_state"] != "ready" ||
		payload["release_gate_state"] != "ready" ||
		payload["kde_action_state"] != "ready" ||
		payload["public_dbus_route_state"] != "ready" ||
		payload["owner_service_trigger_state"] != "ready" ||
		payload["runtime_status_evidence_state"] != "ready" ||
		payload["managed_launcher_request_state"] != "ready" ||
		payload["known_app_artifact_state"] != "ready" ||
		payload["guest_smoke_boundary_state"] != "ready" ||
		payload["public_dbus_method"] != "org.xnix.Compatibility1.ShowRuntimeControlledLaunch" ||
		payload["kde_forwards_only_evidence_handle"] != true ||
		payload["managed_launcher_argv_ready"] != true ||
		payload["formal_release_ready"] != true ||
		payload["desktop_launch_enabled"] != false ||
		payload["backend_launch_enabled"] != false ||
		payload["execution_started"] != false ||
		payload["runtime_state_written"] != false ||
		payload["kde_configuration_written"] != false ||
		payload["dbus_called"] != false ||
		payload["smoke_executed_by_preview"] != false ||
		payload["network_fetch_required"] != false ||
		payload["host_root_modified"] != false {
		t.Fatalf("unexpected readiness payload: %#v", payload)
	}
	if strings.Contains(output.String(), stateRoot) ||
		strings.Contains(output.String(), "owner_service_call_args") ||
		strings.Contains(output.String(), "owner_service_cli_args") ||
		strings.Contains(output.String(), " --service-call ") ||
		strings.Contains(output.String(), ".exe") {
		t.Fatalf("readiness output exposed unsafe details: %s", output.String())
	}
}

func TestDesktopTriggerStagedInvocationReadinessPreviewCommandRejectsMissingStateRoot(t *testing.T) {
	var output bytes.Buffer
	err := run([]string{"desktop-trigger-staged-invocation-readiness-preview"}, &output)
	if err == nil || !strings.Contains(err.Error(), "requires --state-root") {
		t.Fatalf("missing state root must be rejected, got: %v", err)
	}
}

func TestDesktopTriggerStagedInvocationReadinessPreviewCommandClassifiesMissingEvidence(t *testing.T) {
	var output bytes.Buffer
	err := run([]string{
		"desktop-trigger-staged-invocation-readiness-preview",
		"--state-root", t.TempDir(),
	}, &output)
	if err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("readiness output must be JSON: %v\n%s", err, output.String())
	}
	if payload["readiness_state"] != "missing-evidence" ||
		payload["runtime_status_evidence_state"] != "missing-evidence" ||
		payload["formal_release_ready"] != false {
		t.Fatalf("unexpected missing evidence payload: %#v", payload)
	}
}
