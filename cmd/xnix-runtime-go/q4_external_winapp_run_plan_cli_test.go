package main

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"
)

func TestQ4ExternalWinAppRunPlanPreviewCommandRendersQ4OnlyPlan(t *testing.T) {
	var output bytes.Buffer
	if err := run([]string{
		"q4-external-winapp-run-plan-preview",
		"--project-root", projectRootForRuntimeServiceBindingCommandTest(t),
		"--executable", "/tmp/xnix-upload/notepadpp.exe",
		"--app-id", "org.xnix.external.notepadplusplus",
		"--display-name", "Notepad++",
		"--window-match", "Notepad++",
		"--output", "/tmp/xnix-output/run.json",
		"--markdown-output", "/tmp/xnix-output/run.md",
		"--execute",
	}, &output); err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["version"] != currentProjectVersion(t) ||
		payload["schema_version"] != "xnix.runtime.q4_external_winapp_run_plan.v1" ||
		payload["request_type"] != "q4-external-winapp-run-plan-preview" ||
		payload["source"] != "runtime-q4-external-winapp-run-plan+upload-staged-smoke" ||
		payload["runtime_method"] != "PlanQ4ExternalWinAppRun" ||
		payload["read_method"] != "GetQ4ExternalWinAppRunPlanPreview" ||
		payload["app_id"] != "org.xnix.external.notepadplusplus" ||
		payload["display_name"] != "Notepad++" ||
		payload["local_executable_path_class"] != "scoped-temporary" ||
		payload["remote_materials_root_class"] != "q4-home-scoped" ||
		payload["local_host_role"] != "scoped-upload-and-operator-plan-only" ||
		payload["accepted_application_detail_state_required"] != "runtime-accepted-real-app-run" {
		t.Fatalf("unexpected q4 external Windows app run plan CLI payload: %#v", payload)
	}
	if payload["q4_materialization_required"] != true ||
		payload["q4_build_preferred"] != true ||
		payload["q4_execution_required"] != true ||
		payload["host_compilation_required"] != false ||
		payload["host_compilation_avoided"] != true ||
		payload["local_executable_mz_header_required"] != true ||
		payload["local_executable_sha256_required"] != true ||
		payload["remote_upload_required"] != true ||
		payload["real_file_open_evidence_required"] != true ||
		payload["windows_process_window_observation_required"] != true ||
		payload["runtime_owned"] != true ||
		payload["go_runtime_backed"] != true ||
		payload["kde_policy_owner"] != false ||
		payload["full_smoke_required"] != false ||
		payload["host_root_modified"] != false ||
		payload["privileged_container_required"] != false ||
		payload["host_networking_required"] != false ||
		payload["docker_socket_mounted"] != false ||
		payload["broad_host_mount_required"] != false ||
		payload["plan_ready"] != true {
		t.Fatalf("unexpected q4 external Windows app run plan CLI safety flags: %#v", payload)
	}
	command := payload["delegated_command"].([]any)
	if command[0] != "ruby" ||
		command[1] != "scripts/q4_external_winapp_run.rb" ||
		!sliceContainsAny(command, "--execute") ||
		!sliceContainsAny(command, "/tmp/xnix-upload/notepadpp.exe") {
		t.Fatalf("unexpected delegated command: %#v", command)
	}
	if strings.Contains(payload["desktop_safe_summary"].(string), "/tmp/xnix-upload") ||
		strings.Contains(payload["desktop_safe_summary"].(string), "root@q4") {
		t.Fatalf("desktop-safe summary exposed operator details: %s", payload["desktop_safe_summary"])
	}
}

func TestQ4ExternalWinAppRunPlanPreviewCommandRejectsMissingExecutable(t *testing.T) {
	var output bytes.Buffer
	err := run([]string{"q4-external-winapp-run-plan-preview", "--window-match", "Uploaded App"}, &output)
	if err == nil || !strings.Contains(err.Error(), "requires --executable") {
		t.Fatalf("expected missing executable rejection, got %v", err)
	}
}

func TestQ4ExternalWinAppRunPlanPreviewCommandRejectsArguments(t *testing.T) {
	var output bytes.Buffer
	err := run([]string{
		"q4-external-winapp-run-plan-preview",
		"--executable", "/tmp/xnix-upload/app.exe",
		"--window-match", "Uploaded App",
		"extra",
	}, &output)
	if err == nil || !strings.Contains(err.Error(), "does not accept positional arguments") {
		t.Fatalf("expected positional argument rejection, got %v", err)
	}
}

func sliceContainsAny(values []any, expected string) bool {
	for _, value := range values {
		if value == expected {
			return true
		}
	}
	return false
}
