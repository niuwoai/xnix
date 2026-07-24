package main

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"
)

func TestOwnerServiceLaunchEnvelopeGuardPreviewCommandAcceptsEvidenceOnlyEnvelope(t *testing.T) {
	stateRoot := t.TempDir()
	record := recordKDEControlledLaunchActionCLIEvidence(t, stateRoot)

	var output bytes.Buffer
	err := run([]string{
		"owner-service-launch-envelope-guard-preview",
		"--state-root", stateRoot,
		"--route-id", "kde-dbus-runtime-status-action",
		"--method-id", "ShowRuntimeControlledLaunch",
		"--evidence-handle-kind", "evidence-relative-path",
		"--evidence-handle", record.EvidenceRelativePath,
		"--expected-evidence-sha256", record.EvidenceSHA256,
		"--action-id", "xnix.runtime-status.controlled-launch",
		"--caller-role", "kde-presentation-shell",
		"--request-freshness", "fresh",
		"--replay-marker", "fresh",
	}, &output)
	if err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("guard output must be JSON: %v\n%s", err, output.String())
	}
	if payload["schema_version"] != "xnix.runtime.owner_service_launch_envelope_guard.v1" ||
		payload["request_type"] != "owner-service-launch-envelope-guard-preview" ||
		payload["guard_state"] != "accepted-for-review" ||
		payload["accepted_for_review"] != true ||
		payload["route_matched"] != true ||
		payload["method_matched"] != true ||
		payload["action_matched"] != true ||
		payload["caller_role_accepted"] != true ||
		payload["evidence_handle_shape_accepted"] != true ||
		payload["evidence_digest_verified"] != true ||
		payload["evidence_digest_matched"] != true ||
		payload["owner_only_arguments_present"] != false ||
		payload["kde_can_only_forward_evidence_handle"] != true ||
		payload["service_call_dispatched"] != false ||
		payload["dbus_ownership_enabled"] != false ||
		payload["desktop_launch_enabled"] != false ||
		payload["backend_launch_enabled"] != false ||
		payload["execution_started"] != false ||
		payload["runtime_state_written"] != false ||
		payload["host_root_modified"] != false {
		t.Fatalf("unexpected guard payload: %#v", payload)
	}
	if strings.Contains(output.String(), stateRoot) ||
		strings.Contains(output.String(), "owner_service_call_args") ||
		strings.Contains(output.String(), " --service-call ") {
		t.Fatalf("guard output exposed Runtime-owned details: %s", output.String())
	}
}

func TestOwnerServiceLaunchEnvelopeGuardPreviewCommandBlocksOwnerOnlyArgsWithoutLeakingValues(t *testing.T) {
	stateRoot := t.TempDir()
	record := recordKDEControlledLaunchActionCLIEvidence(t, stateRoot)

	var output bytes.Buffer
	err := run([]string{
		"owner-service-launch-envelope-guard-preview",
		"--state-root", stateRoot,
		"--route-id", "kde-dbus-runtime-status-action",
		"--method-id", "ShowRuntimeControlledLaunch",
		"--evidence-handle-kind", "evidence-relative-path",
		"--evidence-handle", record.EvidenceRelativePath,
		"--expected-evidence-sha256", record.EvidenceSHA256,
		"--action-id", "xnix.runtime-status.controlled-launch",
		"--caller-role", "kde-presentation-shell",
		"--kde-state-root", "/private/runtime",
		"--kde-cache-root", "/private/cache",
		"--kde-launcher-path", "/private/launcher",
		"--kde-timeout", "5m",
		"--kde-raw-executable-path", "C:/private/app",
		"--kde-backend-command", "private engine command",
		"--kde-receipt-id", "receipt-1",
		"--kde-session-id", "session-1",
		"--kde-dispatch-id", "dispatch-1",
	}, &output)
	if err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("guard output must be JSON: %v\n%s", err, output.String())
	}
	if payload["guard_state"] != "blocked-owner-args" ||
		payload["accepted_for_review"] != false ||
		payload["owner_only_arguments_present"] != true ||
		payload["owner_only_argument_count"] != float64(9) ||
		payload["kde_state_root_rejected"] != true ||
		payload["kde_cache_root_rejected"] != true ||
		payload["kde_launcher_path_rejected"] != true ||
		payload["kde_timeout_value_rejected"] != true ||
		payload["kde_raw_executable_path_rejected"] != true ||
		payload["kde_backend_command_rejected"] != true ||
		payload["kde_receipt_fields_rejected"] != true ||
		payload["kde_session_fields_rejected"] != true ||
		payload["kde_dispatch_fields_rejected"] != true ||
		payload["service_call_dispatched"] != false ||
		payload["runtime_state_written"] != false {
		t.Fatalf("unexpected owner-args guard payload: %#v", payload)
	}
	for _, unsafe := range []string{stateRoot, "/private/runtime", "/private/cache", "/private/launcher", "C:/private/app", "private engine command", "receipt-1", "session-1", "dispatch-1"} {
		if strings.Contains(output.String(), unsafe) {
			t.Fatalf("guard output leaked owner-only value %q: %s", unsafe, output.String())
		}
	}
}

func TestOwnerServiceLaunchEnvelopeGuardPreviewCommandRejectsMissingStateRoot(t *testing.T) {
	var output bytes.Buffer
	err := run([]string{"owner-service-launch-envelope-guard-preview"}, &output)
	if err == nil || !strings.Contains(err.Error(), "requires --state-root") {
		t.Fatalf("missing state root must be rejected, got: %v", err)
	}
}
