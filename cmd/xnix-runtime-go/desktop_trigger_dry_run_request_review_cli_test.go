package main

import (
	"bytes"
	"encoding/json"
	"path/filepath"
	"strings"
	"testing"
)

func TestDesktopTriggerDryRunRequestReviewPreviewCommand(t *testing.T) {
	stateRoot := t.TempDir()
	record := recordKDEControlledLaunchActionCLIEvidence(t, stateRoot)

	var output bytes.Buffer
	err := run([]string{
		"desktop-trigger-dry-run-request-review-preview",
		"--state-root", stateRoot,
		"--desktop-entry-file", filepath.Join(projectRootForRuntimeServiceBindingCommandTest(t), "kde/actions/xnix-runtime-status-controlled-launch.desktop"),
		"--evidence-relative-path", record.EvidenceRelativePath,
		"--expected-evidence-sha256", record.EvidenceSHA256,
		"--full-checkpoint-promoted",
	}, &output)
	if err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("dry-run review output must be JSON: %v\n%s", err, output.String())
	}
	if payload["schema_version"] != "xnix.runtime.desktop_trigger_dry_run_request_review.v1" ||
		payload["request_type"] != "desktop-trigger-dry-run-request-review-preview" ||
		payload["review_state"] != "accepted-review" ||
		payload["envelope_guard_state"] != "accepted-for-review" ||
		payload["action_surface_state"] != "safe" ||
		payload["managed_launcher_acceptance_state"] != "accepted" ||
		payload["full_checkpoint_state"] != "ready" ||
		payload["runtime_status_evidence_state"] != "ready" ||
		payload["evidence_digest_verified"] != true ||
		payload["expected_digest_matched"] != true ||
		payload["route_id"] != "kde-dbus-runtime-status-action" ||
		payload["method_id"] != "ShowRuntimeControlledLaunch" ||
		payload["action_id"] != "xnix.runtime-status.controlled-launch" ||
		payload["caller_role"] != "kde-presentation-shell" ||
		payload["kde_action_id"] != "xnix.runtime-status.controlled-launch" ||
		payload["public_dbus_method"] != "org.xnix.Compatibility1.ShowRuntimeControlledLaunch" ||
		payload["forwarded_argument_kind"] != "evidence-relative-path" ||
		payload["envelope_accepted_for_review"] != true ||
		payload["action_surface_safe_for_human_smoke"] != true ||
		payload["managed_launcher_accepted"] != true ||
		payload["dry_run_review_only"] != true ||
		payload["request_object_written"] != false ||
		payload["permission_grant_created"] != false ||
		payload["runtime_owned"] != true ||
		payload["go_runtime_backed"] != true ||
		payload["kde_presentation_only"] != true ||
		payload["kde_forwards_only_evidence_handle"] != true ||
		payload["state_root_path_exposed"] != false ||
		payload["raw_launcher_output_exposed"] != false ||
		payload["backend_details_exposed"] != false ||
		payload["service_call_dispatch_enabled"] != false ||
		payload["service_call_dispatched"] != false ||
		payload["dbus_called"] != false ||
		payload["desktop_launch_enabled"] != false ||
		payload["backend_launch_enabled"] != false ||
		payload["execution_started"] != false ||
		payload["runtime_state_written"] != false ||
		payload["kde_configuration_written"] != false ||
		payload["network_required"] != false ||
		payload["host_root_modified"] != false ||
		payload["production_authorization_accepted"] != false {
		t.Fatalf("unexpected dry-run review payload: %#v", payload)
	}
	if strings.Contains(output.String(), stateRoot) ||
		strings.Contains(output.String(), "kde/actions/xnix-runtime-status-controlled-launch.desktop") ||
		strings.Contains(output.String(), "owner_service_call_args") ||
		strings.Contains(output.String(), "owner_service_cli_args") ||
		strings.Contains(output.String(), " --service-call ") ||
		strings.Contains(output.String(), ".exe") {
		t.Fatalf("dry-run review output exposed unsafe details: %s", output.String())
	}
}

func TestDesktopTriggerDryRunRequestReviewPreviewCommandRequiresInputs(t *testing.T) {
	var output bytes.Buffer
	err := run([]string{"desktop-trigger-dry-run-request-review-preview"}, &output)
	if err == nil || !strings.Contains(err.Error(), "requires --state-root") {
		t.Fatalf("missing state root must be rejected, got: %v", err)
	}

	err = run([]string{
		"desktop-trigger-dry-run-request-review-preview",
		"--state-root", t.TempDir(),
	}, &output)
	if err == nil || !strings.Contains(err.Error(), "requires --desktop-entry-file") {
		t.Fatalf("missing desktop entry must be rejected, got: %v", err)
	}
}

func TestDesktopTriggerDryRunRequestReviewPreviewCommandBlocksOwnerOnlyArgs(t *testing.T) {
	stateRoot := t.TempDir()
	record := recordKDEControlledLaunchActionCLIEvidence(t, stateRoot)

	var output bytes.Buffer
	err := run([]string{
		"desktop-trigger-dry-run-request-review-preview",
		"--state-root", stateRoot,
		"--desktop-entry-file", filepath.Join(projectRootForRuntimeServiceBindingCommandTest(t), "kde/actions/xnix-runtime-status-controlled-launch.desktop"),
		"--evidence-relative-path", record.EvidenceRelativePath,
		"--expected-evidence-sha256", record.EvidenceSHA256,
		"--full-checkpoint-promoted",
		"--kde-state-root", "/private/runtime",
		"--kde-backend-command", "private engine command",
	}, &output)
	if err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("dry-run review output must be JSON: %v\n%s", err, output.String())
	}
	if payload["review_state"] != "blocked-unsafe-envelope" ||
		payload["envelope_guard_state"] != "blocked-owner-args" ||
		payload["service_call_dispatched"] != false ||
		payload["runtime_state_written"] != false ||
		payload["host_root_modified"] != false {
		t.Fatalf("unexpected owner-only args dry-run review payload: %#v", payload)
	}
	for _, unsafe := range []string{"/private/runtime", "private engine command", stateRoot} {
		if strings.Contains(output.String(), unsafe) {
			t.Fatalf("dry-run review output leaked unsafe value %q: %s", unsafe, output.String())
		}
	}
}
