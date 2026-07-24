package main

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"
)

func TestManagedLauncherAcceptanceReportPreviewCommand(t *testing.T) {
	stateRoot := t.TempDir()
	record := recordKDEControlledLaunchActionCLIEvidence(t, stateRoot)

	var output bytes.Buffer
	err := run([]string{
		"managed-launcher-acceptance-report-preview",
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
		t.Fatalf("acceptance output must be JSON: %v\n%s", err, output.String())
	}
	if payload["schema_version"] != "xnix.runtime.managed_launcher_acceptance_report.v1" ||
		payload["request_type"] != "managed-launcher-acceptance-report-preview" ||
		payload["acceptance_state"] != "accepted" ||
		payload["known_app_identity_state"] != "ready" ||
		payload["artifact_digest_state"] != "ready" ||
		payload["launch_authorization_state"] != "ready" ||
		payload["session_gated_review_state"] != "ready" ||
		payload["controlled_execution_session_state"] != "ready" ||
		payload["managed_launcher_request_state"] != "ready" ||
		payload["guest_smoke_boundary_state"] != "ready" ||
		payload["runtime_status_evidence_state"] != "ready" ||
		payload["kde_safe_projection_state"] != "ready" ||
		payload["full_checkpoint_state"] != "ready" ||
		payload["desktop_trigger_readiness_state"] != "ready" ||
		payload["compatibility_center_projection_ready"] != true ||
		payload["kde_center_projection_ready"] != true ||
		payload["managed_launcher_argv_ready"] != true ||
		payload["known_app_acceptance_ready"] != true ||
		payload["formal_release_ready"] != true ||
		payload["state_root_path_exposed"] != false ||
		payload["raw_launcher_output_exposed"] != false ||
		payload["backend_details_exposed"] != false ||
		payload["desktop_launch_enabled"] != false ||
		payload["backend_launch_enabled"] != false ||
		payload["execution_started"] != false ||
		payload["runtime_state_written"] != false ||
		payload["kde_configuration_written"] != false ||
		payload["dbus_called"] != false ||
		payload["network_required"] != false ||
		payload["host_root_modified"] != false {
		t.Fatalf("unexpected acceptance payload: %#v", payload)
	}
	if strings.Contains(output.String(), stateRoot) ||
		strings.Contains(output.String(), "owner_service_call_args") ||
		strings.Contains(output.String(), "owner_service_cli_args") ||
		strings.Contains(output.String(), " --service-call ") ||
		strings.Contains(output.String(), ".exe") {
		t.Fatalf("acceptance output exposed unsafe details: %s", output.String())
	}
}

func TestManagedLauncherAcceptanceReportPreviewCommandRejectsMissingStateRoot(t *testing.T) {
	var output bytes.Buffer
	err := run([]string{"managed-launcher-acceptance-report-preview"}, &output)
	if err == nil || !strings.Contains(err.Error(), "requires --state-root") {
		t.Fatalf("missing state root must be rejected, got: %v", err)
	}
}

func TestManagedLauncherAcceptanceReportPreviewCommandClassifiesMissingEvidence(t *testing.T) {
	var output bytes.Buffer
	err := run([]string{
		"managed-launcher-acceptance-report-preview",
		"--state-root", t.TempDir(),
	}, &output)
	if err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("acceptance output must be JSON: %v\n%s", err, output.String())
	}
	if payload["acceptance_state"] != "missing-evidence" ||
		payload["runtime_status_evidence_state"] != "missing-evidence" ||
		payload["formal_release_ready"] != false {
		t.Fatalf("unexpected missing evidence payload: %#v", payload)
	}
}
