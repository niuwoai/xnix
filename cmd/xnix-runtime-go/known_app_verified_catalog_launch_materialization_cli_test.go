package main

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestKnownAppVerifiedCatalogLaunchMaterializationRecordCommandConsumesHandoff(t *testing.T) {
	stateRoot := t.TempDir()
	acceptancePath := writeKnownAppVerifiedCatalogRunAcceptanceCLIFile(t, "7zr")
	var handoffOutput bytes.Buffer
	if err := run([]string{
		"known-app-verified-catalog-launch-handoff-record",
		"--state-root", stateRoot,
		"--known-app-verified-catalog-run-acceptance", acceptancePath,
	}, &handoffOutput); err != nil {
		t.Fatalf("known-app-verified-catalog-launch-handoff-record returned error: %v", err)
	}
	var handoff map[string]any
	if err := json.Unmarshal(handoffOutput.Bytes(), &handoff); err != nil {
		t.Fatalf("handoff output must be JSON: %v\n%s", err, handoffOutput.String())
	}
	relativePath, ok := handoff["handoff_relative_path"].(string)
	if !ok || relativePath == "" {
		t.Fatalf("handoff output must expose a relative handoff handle: %#v", handoff)
	}
	var output bytes.Buffer
	err := run([]string{
		"known-app-verified-catalog-launch-materialization-record",
		"--state-root", stateRoot,
		"--handoff-relative-path", relativePath,
		"--cache-root", filepath.Join(t.TempDir(), "missing-cache"),
	}, &output)
	if err != nil {
		t.Fatalf("known-app-verified-catalog-launch-materialization-record returned error: %v", err)
	}
	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("materialization output must be JSON: %v\n%s", err, output.String())
	}
	if payload["version"] != currentProjectVersion(t) ||
		payload["schema_version"] != "xnix.runtime.known_app_verified_catalog_launch_materialization.v1" ||
		payload["request_type"] != "known-app-verified-catalog-launch-materialization-record" ||
		payload["app_id"] != "7zr" ||
		payload["handoff_consumed"] != true ||
		payload["handoff_relative_path"] != relativePath ||
		payload["handoff_digest_verified"] != true ||
		payload["acceptance_ready"] != true ||
		payload["run_plan_matched"] != true ||
		payload["launch_authorization_receipt_recorded"] != true ||
		payload["controlled_session_record_state"] != "blocked" ||
		payload["materialization_state"] != "blocked-managed-artifact-required" ||
		payload["materialization_ready"] != false ||
		payload["runtime_owner_materialized"] != false ||
		payload["owner_materialization_required"] != true ||
		payload["runtime_owner_service_supplies_inputs"] != true ||
		payload["dispatch_runner_required"] != false ||
		payload["execution_started"] != false ||
		payload["backend_process_started"] != false {
		t.Fatalf("unexpected materialization payload: %#v", payload)
	}
	if strings.Contains(output.String(), stateRoot) || strings.Contains(output.String(), acceptancePath) {
		t.Fatalf("materialization output exposed local paths: %s", output.String())
	}
	if _, err := os.Stat(filepath.Join(stateRoot, filepath.FromSlash(relativePath))); err != nil {
		t.Fatalf("handoff payload should remain readable under state root: %v", err)
	}
	for _, forbidden := range []string{"prefix", ".exe", "program files", "qemu-system", "proton", "wine ", "wine/", ".wine"} {
		if strings.Contains(strings.ToLower(output.String()), forbidden) {
			t.Fatalf("materialization output exposed backend term %q: %s", forbidden, output.String())
		}
	}
}

func TestKnownAppVerifiedCatalogLaunchMaterializationRecordCommandRejectsMissingInputs(t *testing.T) {
	var output bytes.Buffer
	err := run([]string{"known-app-verified-catalog-launch-materialization-record"}, &output)
	if err == nil || !strings.Contains(err.Error(), "requires --state-root") {
		t.Fatalf("missing state root must be rejected, got: %v", err)
	}
	err = run([]string{"known-app-verified-catalog-launch-materialization-record", "--state-root", t.TempDir()}, &output)
	if err == nil || !strings.Contains(err.Error(), "requires --handoff-relative-path") {
		t.Fatalf("missing handoff path must be rejected, got: %v", err)
	}
}

func TestKnownAppVerifiedCatalogDispatchRequestRecordCommandBlocksUntilMaterializationReady(t *testing.T) {
	stateRoot := t.TempDir()
	acceptancePath := writeKnownAppVerifiedCatalogRunAcceptanceCLIFile(t, "7zr")
	var handoffOutput bytes.Buffer
	if err := run([]string{
		"known-app-verified-catalog-launch-handoff-record",
		"--state-root", stateRoot,
		"--known-app-verified-catalog-run-acceptance", acceptancePath,
	}, &handoffOutput); err != nil {
		t.Fatalf("known-app-verified-catalog-launch-handoff-record returned error: %v", err)
	}
	var handoff map[string]any
	if err := json.Unmarshal(handoffOutput.Bytes(), &handoff); err != nil {
		t.Fatalf("handoff output must be JSON: %v\n%s", err, handoffOutput.String())
	}
	relativePath, ok := handoff["handoff_relative_path"].(string)
	if !ok || relativePath == "" {
		t.Fatalf("handoff output must expose a relative handoff handle: %#v", handoff)
	}
	var output bytes.Buffer
	err := run([]string{
		"known-app-verified-catalog-dispatch-request-record",
		"--state-root", stateRoot,
		"--handoff-relative-path", relativePath,
		"--cache-root", filepath.Join(t.TempDir(), "missing-cache"),
	}, &output)
	if err != nil {
		t.Fatalf("known-app-verified-catalog-dispatch-request-record returned error: %v", err)
	}
	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("dispatch request output must be JSON: %v\n%s", err, output.String())
	}
	if payload["version"] != currentProjectVersion(t) ||
		payload["schema_version"] != "xnix.runtime.known_app_verified_catalog_dispatch_request.v1" ||
		payload["request_type"] != "known-app-verified-catalog-dispatch-request-record" ||
		payload["source"] != "known-app-verified-catalog-launch-materialization-record+dispatch-runner-request" ||
		payload["app_id"] != "7zr" ||
		payload["handoff_consumed"] != true ||
		payload["handoff_relative_path"] != relativePath ||
		payload["materialization_state"] != "blocked-managed-artifact-required" ||
		payload["materialization_ready"] != false ||
		payload["dispatch_request_record_state"] != "blocked-materialization-required" ||
		payload["dispatch_request_written"] != false ||
		payload["dispatch_runner_request_type"] != "windows-known-app-dispatch-smoke" ||
		payload["dispatch_runner_name"] != "xnix-compat-launch" ||
		payload["dispatch_runner_argument_values_exposed"] != false ||
		payload["runtime_owner_dispatch_inputs_ready"] != false ||
		payload["dispatch_allowed"] != false ||
		payload["dispatch_started"] != false ||
		payload["execution_started"] != false ||
		payload["backend_process_started"] != false ||
		payload["next_owner_action"] != "materialize-launch-session" {
		t.Fatalf("unexpected blocked dispatch request payload: %#v", payload)
	}
	if strings.Contains(output.String(), stateRoot) || strings.Contains(output.String(), acceptancePath) {
		t.Fatalf("dispatch request output exposed owner paths: %s", output.String())
	}
	if _, err := os.Stat(filepath.Join(stateRoot, "runtime", "known-app-verified-catalog-dispatch-requests")); !os.IsNotExist(err) {
		t.Fatalf("blocked dispatch request must not create request directory: %v", err)
	}
}

func TestKnownAppVerifiedCatalogDispatchRequestRecordCommandRejectsMissingInputs(t *testing.T) {
	var output bytes.Buffer
	err := run([]string{"known-app-verified-catalog-dispatch-request-record"}, &output)
	if err == nil || !strings.Contains(err.Error(), "requires --state-root") {
		t.Fatalf("missing state root must be rejected, got: %v", err)
	}
	err = run([]string{"known-app-verified-catalog-dispatch-request-record", "--state-root", t.TempDir()}, &output)
	if err == nil || !strings.Contains(err.Error(), "requires --handoff-relative-path") {
		t.Fatalf("missing handoff path must be rejected, got: %v", err)
	}
}

func TestKnownAppVerifiedCatalogDispatchRunnerExecutionCommandInvokesManagedLauncher(t *testing.T) {
	stateRoot := t.TempDir()
	cacheRoot := filepath.Join(t.TempDir(), "known-cache")
	sessionID := "known-app-controlled-execution-session-7zr-26.02"
	requestRelativePath := "runtime/known-app-verified-catalog-dispatch-requests/known-app-verified-catalog-dispatch-request-7zr-26.02.json"
	requestPath := filepath.Join(stateRoot, filepath.FromSlash(requestRelativePath))
	if err := os.MkdirAll(filepath.Dir(requestPath), 0o700); err != nil {
		t.Fatalf("MkdirAll request dir returned error: %v", err)
	}
	requestJSON := `{
  "schema_version": "xnix.runtime.known_app_verified_catalog_dispatch_request.v1",
  "request_type": "known-app-verified-catalog-dispatch-request-record",
  "app_id": "7zr",
  "display_name": "7-Zip standalone console executable",
  "app_version": "26.02",
  "launcher_name": "xnix-compat-launch",
  "runner_request_type": "windows-known-app-dispatch-smoke",
  "runner_argv": [
    "xnix-compat-launch",
    "--app", "7zr",
    "--cache-root", "` + cacheRoot + `",
    "--guest-boundary", "managed-known-app-guest-smoke",
    "--state-root", "` + stateRoot + `",
    "--receipt-id", "known-app-launch-authorization-7zr-26.02",
    "--review-receipt-id", "known-app-session-gated-launch-review-7zr-26.02-known-app-controlled-execution-session-7zr-26.02",
    "--session-id", "` + sessionID + `"
  ],
  "launch_authorization_receipt_id": "known-app-launch-authorization-7zr-26.02",
  "session_gated_review_receipt_id": "known-app-session-gated-launch-review-7zr-26.02-known-app-controlled-execution-session-7zr-26.02",
  "controlled_execution_session_id": "` + sessionID + `",
  "controlled_session_relative_path": "execution-ledger/sessions/known-app-controlled-execution-session-7zr-26.02.json",
  "guest_boundary": "managed-known-app-guest-smoke",
  "recorded_at_utc": "2026-07-27T07:08:09Z"
}
`
	if err := os.WriteFile(requestPath, []byte(requestJSON), 0o600); err != nil {
		t.Fatalf("WriteFile request returned error: %v", err)
	}
	argsLog := filepath.Join(t.TempDir(), "launcher-args.log")
	fakeLauncher := writeVerifiedCatalogFakeDispatchLauncher(t, argsLog)
	t.Setenv("XNIX_FAKE_DISPATCH_ARGS_LOG", argsLog)

	var output bytes.Buffer
	err := run([]string{
		"known-app-verified-catalog-dispatch-runner-execution",
		"--state-root", stateRoot,
		"--dispatch-request-relative-path", requestRelativePath,
		"--launcher", fakeLauncher,
		"--host", "127.0.0.1",
		"--port", "2222",
		"--user", "root",
		"--key", filepath.Join(t.TempDir(), "id_ed25519"),
		"--remote-dir", "/tmp/xnix-known-winapp-smoke",
		"--owner-timeout", "5s",
	}, &output)
	if err != nil {
		t.Fatalf("known-app-verified-catalog-dispatch-runner-execution returned error: %v\n%s", err, output.String())
	}
	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("execution output must be JSON: %v\n%s", err, output.String())
	}
	if payload["version"] != currentProjectVersion(t) ||
		payload["schema_version"] != "xnix.runtime.known_app_verified_catalog_dispatch_execution.v1" ||
		payload["request_type"] != "known-app-verified-catalog-dispatch-runner-execution" ||
		payload["source"] != "known-app-verified-catalog-dispatch-request-record+runner-consumption" ||
		payload["app_id"] != "7zr" ||
		payload["dispatch_request_consumed"] != true ||
		payload["dispatch_request_relative_path"] != requestRelativePath ||
		payload["dispatch_request_digest_verified"] != true ||
		payload["dispatch_runner_request_type"] != "windows-known-app-dispatch-smoke" ||
		payload["dispatch_runner_name"] != filepath.Base(fakeLauncher) ||
		payload["dispatch_runner_invoked"] != true ||
		payload["dispatch_runner_exit_code"] != float64(0) ||
		payload["dispatch_runner_output_json_observed"] != true ||
		payload["dispatch_runner_argument_values_exposed"] != false ||
		payload["runtime_owner_transport_supplied"] != true ||
		payload["guest_start_requested"] != false ||
		payload["guest_start_mode"] != "external" ||
		payload["delegated_request_type"] != "windows-known-app-dispatch-smoke" ||
		payload["delegated_evidence_source"] != "wine-guest-gui-smoke" ||
		payload["delegated_status"] != "passed" ||
		payload["delegated_cache_status"] != "guest-builtin-gui" ||
		payload["delegated_runtime_owned_dispatch"] != true ||
		payload["delegated_smoke_passed"] != true ||
		payload["delegated_execution_started"] != true ||
		payload["delegated_backend_process_started"] != false ||
		payload["delegated_managed_guest_runner_invoked"] != true ||
		payload["delegated_managed_guest_reachable"] != true ||
		payload["delegated_managed_guest_runtime_ready"] != true ||
		payload["delegated_file_argument_count"] != float64(1) ||
		payload["delegated_file_argument_copied_count"] != float64(1) ||
		payload["delegated_file_arguments_passed"] != true ||
		payload["delegated_file_argument_winepath_translated"] != true ||
		payload["delegated_file_argument_winepath_translated_count"] != float64(1) ||
		payload["delegated_window_match_observed"] != true ||
		payload["delegated_window_evidence_observed"] != true ||
		payload["delegated_gui_evidence_ready"] != true ||
		payload["delegated_controlled_execution_session_consumed"] != true ||
		payload["delegated_controlled_session_digest_verified"] != true ||
		payload["delegated_controlled_session_window_observed"] != true ||
		payload["raw_host_path_exposed"] != false ||
		payload["raw_executable_path_exposed"] != false ||
		payload["raw_file_argument_path_exposed"] != false ||
		payload["window_evidence_summary_exposed"] != false ||
		payload["raw_launcher_output_exposed"] != false ||
		payload["backend_details_exposed"] != false {
		t.Fatalf("unexpected dispatch runner execution payload: %#v", payload)
	}
	for _, forbidden := range []string{stateRoot, cacheRoot, fakeLauncher, "id_ed25519", "/tmp/xnix-known-winapp-smoke"} {
		if strings.Contains(output.String(), forbidden) {
			t.Fatalf("execution output exposed owner-only value %q: %s", forbidden, output.String())
		}
	}
	argsBytes, err := os.ReadFile(argsLog)
	if err != nil {
		t.Fatalf("ReadFile args log returned error: %v", err)
	}
	argsText := string(argsBytes)
	for _, expected := range []string{"--state-root", stateRoot, "--cache-root", cacheRoot, "--receipt-id", "known-app-launch-authorization-7zr-26.02", "--review-receipt-id", "--session-id", sessionID, "--host", "127.0.0.1", "--port", "2222"} {
		if !strings.Contains(argsText, expected) {
			t.Fatalf("managed launcher did not receive expected private argument %q: %s", expected, argsText)
		}
	}
}

func TestKnownAppVerifiedCatalogDispatchRunnerExecutionCommandRejectsMissingInputs(t *testing.T) {
	var output bytes.Buffer
	err := run([]string{"known-app-verified-catalog-dispatch-runner-execution"}, &output)
	if err == nil || !strings.Contains(err.Error(), "requires --state-root") {
		t.Fatalf("missing state root must be rejected, got: %v", err)
	}
	err = run([]string{"known-app-verified-catalog-dispatch-runner-execution", "--state-root", t.TempDir()}, &output)
	if err == nil || !strings.Contains(err.Error(), "requires --dispatch-request-relative-path") {
		t.Fatalf("missing dispatch request path must be rejected, got: %v", err)
	}
}

func writeVerifiedCatalogFakeDispatchLauncher(t *testing.T, argsLog string) string {
	t.Helper()
	launcher := filepath.Join(t.TempDir(), "xnix-compat-launch")
	script := "#!/bin/sh\nprintf '%s\\n' \"$@\" > \"$XNIX_FAKE_DISPATCH_ARGS_LOG\"\nprintf '%s\\n' '{\"request_type\":\"windows-known-app-dispatch-smoke\",\"evidence_source\":\"wine-guest-gui-smoke\",\"status\":\"passed\",\"guest_boundary\":\"managed-known-app-guest-smoke\",\"cache_status\":\"guest-builtin-gui\",\"runtime_owned_dispatch\":true,\"artifact_verified\":true,\"managed_artifact_copied\":true,\"marker_observed\":true,\"smoke_passed\":true,\"execution_started\":true,\"backend_process_started\":false,\"managed_guest_runner_invoked\":true,\"managed_guest_reachable\":true,\"managed_guest_runtime_ready\":true,\"file_argument_count\":1,\"file_argument_copied_count\":1,\"file_arguments_passed\":true,\"file_argument_winepath_translated\":true,\"file_argument_winepath_translated_count\":1,\"window_match_observed\":true,\"controlled_execution_session_consumed\":true,\"controlled_session_digest_verified\":true,\"controlled_session_window_observed\":true,\"host_root_modified\":false,\"privileged_container_required\":false,\"host_networking_required\":false,\"docker_socket_mounted\":false,\"broad_host_mount_required\":false,\"raw_host_path_exposed\":false,\"raw_executable_path_exposed\":false,\"raw_file_argument_path_exposed\":false,\"raw_command_exposed\":false,\"backend_details_exposed\":false}'\n"
	if err := os.WriteFile(launcher, []byte(script), 0o700); err != nil {
		t.Fatalf("WriteFile fake launcher returned error: %v", err)
	}
	return launcher
}
