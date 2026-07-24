package main

import (
	"bytes"
	"encoding/json"
	"path/filepath"
	"strings"
	"testing"
)

func TestDesktopTriggerRequestPreflightPreviewCommandReady(t *testing.T) {
	stateRoot := t.TempDir()
	record := recordKDEControlledLaunchActionCLIEvidence(t, stateRoot)

	var output bytes.Buffer
	err := run([]string{
		"desktop-trigger-request-preflight-preview",
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
		t.Fatalf("desktop-trigger request preflight output must be JSON: %v\n%s", err, output.String())
	}
	if payload["schema_version"] != "xnix.runtime.desktop_trigger_request_preflight.v1" ||
		payload["request_type"] != "desktop-trigger-request-preflight-preview" ||
		payload["preflight_state"] != "ready-for-operator-request" ||
		payload["materialization_state"] != "ready-for-human-authorized-service-call" ||
		payload["dry_run_review_state"] != "accepted-review" ||
		payload["owner_trigger_state"] != "ready" ||
		payload["full_checkpoint_state"] != "ready" ||
		payload["runtime_status_evidence_state"] != "ready" ||
		payload["evidence_digest_verified"] != true ||
		payload["expected_digest_matched"] != true ||
		payload["desktop_dbus_method"] != "org.xnix.Compatibility1.ShowRuntimeControlledLaunch" ||
		payload["owner_service_boundary"] != "go-runtime-owner-in-process-service" ||
		payload["owner_service_method"] != "ShowRuntimeControlledLaunch" ||
		payload["owner_service_call_shape_verified"] != true ||
		payload["owner_service_call_ready"] != true ||
		payload["operator_request_ready"] != true ||
		payload["formal_promotion_required"] != true ||
		payload["formal_promotion_observed"] != true ||
		payload["formal_release_ready"] != false ||
		payload["human_authorization_required"] != true ||
		payload["runtime_owner_service_supplies_inputs"] != true ||
		payload["desktop_evidence_handle_forwarded"] != true ||
		payload["runtime_owned"] != true ||
		payload["go_runtime_backed"] != true ||
		payload["kde_presentation_only"] != true ||
		payload["kde_forwards_only_evidence_handle"] != true ||
		payload["kde_receives_materialized_owner_args"] != false ||
		payload["state_root_path_exposed"] != false ||
		payload["raw_launcher_output_exposed"] != false ||
		payload["backend_details_exposed"] != false ||
		payload["request_object_written"] != false ||
		payload["permission_grant_created"] != false ||
		payload["service_call_dispatch_enabled"] != false ||
		payload["service_call_dispatched"] != false ||
		payload["dbus_called"] != false ||
		payload["desktop_launch_enabled"] != false ||
		payload["backend_launch_enabled"] != false ||
		payload["execution_started"] != false ||
		payload["runtime_state_written"] != false ||
		payload["kde_configuration_written"] != false ||
		payload["network_required"] != false ||
		payload["host_root_modified"] != false {
		t.Fatalf("unexpected desktop-trigger request preflight payload: %#v", payload)
	}
	if _, ok := payload["owner_service_call_args"]; ok {
		t.Fatalf("preflight must not emit owner service call args: %#v", payload)
	}
	if _, ok := payload["owner_service_cli_args"]; ok {
		t.Fatalf("preflight must not emit owner service CLI args: %#v", payload)
	}
	if strings.Contains(output.String(), stateRoot) ||
		strings.Contains(output.String(), "kde/actions/xnix-runtime-status-controlled-launch.desktop") ||
		strings.Contains(output.String(), "owner_service_call_args") ||
		strings.Contains(output.String(), "owner_service_cli_args") ||
		strings.Contains(output.String(), " --service-call ") ||
		strings.Contains(output.String(), "XNIX_RUNTIME_OWNER_") ||
		strings.Contains(output.String(), " --state-root ") ||
		strings.Contains(output.String(), ".exe") {
		t.Fatalf("desktop-trigger request preflight output exposed unsafe details: %s", output.String())
	}
}

func TestDesktopTriggerRequestPreflightPreviewCommandBlocksMissingPromotion(t *testing.T) {
	stateRoot := t.TempDir()
	record := recordKDEControlledLaunchActionCLIEvidence(t, stateRoot)

	var output bytes.Buffer
	err := run([]string{
		"desktop-trigger-request-preflight-preview",
		"--state-root", stateRoot,
		"--desktop-entry-file", filepath.Join(projectRootForRuntimeServiceBindingCommandTest(t), "kde/actions/xnix-runtime-status-controlled-launch.desktop"),
		"--evidence-relative-path", record.EvidenceRelativePath,
		"--expected-evidence-sha256", record.EvidenceSHA256,
	}, &output)
	if err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("desktop-trigger request preflight output must be JSON: %v\n%s", err, output.String())
	}
	if payload["preflight_state"] != "blocked-missing-promotion" ||
		payload["materialization_state"] != "blocked-missing-full-checkpoint" ||
		payload["full_checkpoint_state"] != "needs-full-checkpoint" ||
		payload["operator_request_ready"] != false ||
		payload["owner_service_call_ready"] != false ||
		payload["owner_service_call_shape_verified"] != false ||
		payload["formal_promotion_observed"] != false ||
		payload["formal_release_ready"] != false ||
		payload["service_call_dispatched"] != false ||
		payload["host_root_modified"] != false {
		t.Fatalf("unexpected missing-promotion preflight payload: %#v", payload)
	}
	if strings.Contains(output.String(), stateRoot) ||
		strings.Contains(output.String(), "owner_service_call_args") ||
		strings.Contains(output.String(), " --service-call ") ||
		strings.Contains(output.String(), ".exe") {
		t.Fatalf("desktop-trigger request preflight output exposed unsafe details: %s", output.String())
	}
}

func TestDesktopTriggerRequestPreflightPreviewCommandRequiresInputs(t *testing.T) {
	var output bytes.Buffer
	err := run([]string{"desktop-trigger-request-preflight-preview"}, &output)
	if err == nil || !strings.Contains(err.Error(), "requires --state-root") {
		t.Fatalf("missing state root must be rejected, got: %v", err)
	}

	err = run([]string{
		"desktop-trigger-request-preflight-preview",
		"--state-root", t.TempDir(),
	}, &output)
	if err == nil || !strings.Contains(err.Error(), "requires --desktop-entry-file") {
		t.Fatalf("missing desktop entry must be rejected, got: %v", err)
	}
}

func TestDesktopTriggerRequestPreflightPreviewCommandBlocksOwnerOnlyArgs(t *testing.T) {
	stateRoot := t.TempDir()
	record := recordKDEControlledLaunchActionCLIEvidence(t, stateRoot)

	var output bytes.Buffer
	err := run([]string{
		"desktop-trigger-request-preflight-preview",
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
		t.Fatalf("desktop-trigger request preflight output must be JSON: %v\n%s", err, output.String())
	}
	if payload["preflight_state"] != "blocked-unsafe-envelope" ||
		payload["materialization_state"] != "blocked-unsafe-envelope" ||
		payload["owner_service_call_ready"] != false ||
		payload["owner_service_call_shape_verified"] != false ||
		payload["operator_request_ready"] != false ||
		payload["service_call_dispatched"] != false ||
		payload["host_root_modified"] != false {
		t.Fatalf("unexpected owner-only preflight payload: %#v", payload)
	}
	for _, unsafe := range []string{stateRoot, "/private/runtime", "private engine command", "owner_service_call_args", " --service-call ", ".exe"} {
		if strings.Contains(output.String(), unsafe) {
			t.Fatalf("desktop-trigger request preflight leaked unsafe value %q: %s", unsafe, output.String())
		}
	}
}
