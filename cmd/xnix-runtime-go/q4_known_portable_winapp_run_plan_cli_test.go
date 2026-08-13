package main

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"
)

func TestQ4KnownPortableWinAppRunPlanPreviewCommandRendersCatalogBackedQ4Plan(t *testing.T) {
	var output bytes.Buffer
	if err := run([]string{
		"q4-known-portable-winapp-run-plan-preview",
		"--app", "org.xnix.external.notepadplusplus",
		"--remote", "root@q4",
		"--remote-materials-root", "/home/xnix-run-materials",
		"--remote-source-root", "/home/xnix-build/xnix-known-portable",
		"--remote-build-root", "/home/xnix-build-cache",
		"--output", "/tmp/xnix-output/known-portable-run.json",
		"--markdown-output", "/tmp/xnix-output/known-portable-run.md",
		"--execute",
	}, &output); err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["version"] != currentProjectVersion(t) ||
		payload["schema_version"] != "xnix.runtime.q4_known_portable_winapp_run_plan.v1" ||
		payload["request_type"] != "q4-known-portable-winapp-run-plan-preview" ||
		payload["source"] != "runtime-q4-known-portable-winapp-run-plan+catalog-backed-runner" ||
		payload["runtime_method"] != "PlanQ4KnownPortableWinAppRun" ||
		payload["read_method"] != "GetQ4KnownPortableWinAppRunPlanPreview" ||
		payload["app_id"] != "org.xnix.external.notepadplusplus" ||
		payload["display_name"] != "Notepad++ Portable" ||
		payload["app_version"] != "8.9.7" ||
		payload["catalog_artifact_kind"] != "portable-zip-bundle" ||
		payload["download_artifact_name"] != "npp.8.9.7.portable.zip" ||
		payload["executable_relative_path"] != "notepad++.exe" ||
		payload["launch_source_request_type"] != "windows-known-app-bundle-stage-and-launch" ||
		payload["local_host_role"] != "operator-plan-and-artifact-review-only" ||
		payload["accepted_application_detail_state_needed"] != "runtime-accepted-real-app-run" {
		t.Fatalf("unexpected known portable run plan payload: %#v", payload)
	}
	supportedIDs := payload["supported_known_portable_app_ids"].([]any)
	if len(supportedIDs) != 2 ||
		!sliceContainsAny(supportedIDs, "org.xnix.external.notepadplusplus") ||
		!sliceContainsAny(supportedIDs, "org.xnix.external.putty") {
		t.Fatalf("unexpected supported q4 known portable app ids: %#v", supportedIDs)
	}
	if payload["known_catalog_app"] != true ||
		payload["portable_directory_external_app"] != true ||
		payload["single_file_external_app"] != false ||
		payload["official_download_required"] != true ||
		payload["pinned_checksum_required"] != true ||
		payload["known_bundle_import_required"] != true ||
		payload["known_bundle_stage_launch_required"] != true ||
		payload["runtime_acceptance_required"] != true ||
		payload["remote_host_configured"] != true ||
		payload["remote_host_exposed_to_desktop"] != false ||
		payload["remote_paths_exposed_to_desktop"] != false ||
		payload["q4_download_required"] != true ||
		payload["q4_extract_required"] != true ||
		payload["q4_compile_required"] != true ||
		payload["q4_execution_required"] != true ||
		payload["q4_build_preferred"] != true ||
		payload["host_compilation_required"] != false ||
		payload["host_compilation_avoided"] != true ||
		payload["host_download_avoided"] != true ||
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
		t.Fatalf("unexpected known portable run plan safety flags: %#v", payload)
	}
	command := payload["delegated_command"].([]any)
	if command[0] != "ruby" ||
		command[1] != "scripts/q4_known_portable_winapp_run.rb" ||
		!sliceContainsAny(command, "--execute") ||
		!sliceContainsAny(command, "org.xnix.external.notepadplusplus") {
		t.Fatalf("unexpected delegated command: %#v", command)
	}
	if strings.Contains(payload["desktop_safe_summary"].(string), "root@q4") ||
		strings.Contains(payload["desktop_safe_summary"].(string), "/home/xnix-") {
		t.Fatalf("desktop-safe summary exposed q4 details: %s", payload["desktop_safe_summary"])
	}
}

func TestQ4KnownPortableWinAppRunPlanPreviewCommandRendersPuttySingleExecutableLane(t *testing.T) {
	var output bytes.Buffer
	if err := run([]string{
		"q4-known-portable-winapp-run-plan-preview",
		"--app", "org.xnix.external.putty",
		"--remote", "root@q4",
		"--remote-materials-root", "/home/xnix-run-materials",
		"--remote-source-root", "/home/xnix-build/xnix-known-portable",
		"--remote-build-root", "/home/xnix-build-cache",
		"--output", "/tmp/xnix-output/putty-known-run.json",
		"--markdown-output", "/tmp/xnix-output/putty-known-run.md",
		"--execute",
	}, &output); err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["app_id"] != "org.xnix.external.putty" ||
		payload["display_name"] != "PuTTY" ||
		payload["app_version"] != "0.84" ||
		payload["catalog_artifact_kind"] != "single-executable" ||
		payload["download_artifact_name"] != "putty.exe" ||
		payload["executable_relative_path"] != "putty.exe" ||
		payload["launch_source_request_type"] != "external-winapp-import-stage-and-launch" ||
		payload["delegated_script"] != "scripts/q4_putty_external_winapp_smoke.rb" ||
		payload["delegated_request_type"] != "q4-putty-external-winapp-smoke" ||
		payload["delegated_acceptance_request_type"] != "q4-staged-external-winapp-acceptance-preview" {
		t.Fatalf("unexpected PuTTY known portable run plan payload: %#v", payload)
	}
	if payload["known_catalog_app"] != true ||
		payload["portable_directory_external_app"] != false ||
		payload["single_file_external_app"] != true ||
		payload["official_download_required"] != true ||
		payload["pinned_checksum_required"] != true ||
		payload["known_bundle_import_required"] != false ||
		payload["known_bundle_stage_launch_required"] != false ||
		payload["runtime_acceptance_required"] != true ||
		payload["q4_download_required"] != true ||
		payload["q4_extract_required"] != false ||
		payload["q4_compile_required"] != true ||
		payload["q4_execution_required"] != true ||
		payload["host_compilation_avoided"] != true ||
		payload["host_download_avoided"] != true ||
		payload["host_root_modified"] != false ||
		payload["privileged_container_required"] != false ||
		payload["host_networking_required"] != false ||
		payload["docker_socket_mounted"] != false ||
		payload["broad_host_mount_required"] != false ||
		payload["plan_ready"] != true {
		t.Fatalf("unexpected PuTTY known portable run plan safety flags: %#v", payload)
	}
	command := payload["delegated_command"].([]any)
	if command[0] != "ruby" ||
		command[1] != "scripts/q4_putty_external_winapp_smoke.rb" ||
		!sliceContainsAny(command, "--execute") ||
		!sliceContainsAny(command, "/tmp/xnix-output/putty-known-run.json") {
		t.Fatalf("unexpected PuTTY delegated command: %#v", command)
	}
	if strings.Contains(payload["desktop_safe_summary"].(string), "root@q4") ||
		strings.Contains(payload["desktop_safe_summary"].(string), "/home/xnix-") {
		t.Fatalf("desktop-safe summary exposed q4 details: %s", payload["desktop_safe_summary"])
	}
}

func TestQ4KnownPortableWinAppRunPlanPreviewCommandRejectsUnsupportedApp(t *testing.T) {
	var output bytes.Buffer
	err := run([]string{
		"q4-known-portable-winapp-run-plan-preview",
		"--app", "org.example.unsupported",
	}, &output)
	if err == nil || !strings.Contains(err.Error(), "unknown known Windows app") {
		t.Fatalf("expected unsupported app rejection, got %v", err)
	}
}

func TestQ4KnownPortableWinAppRunPlanPreviewCommandRejectsKnownNonBundleApp(t *testing.T) {
	var output bytes.Buffer
	err := run([]string{
		"q4-known-portable-winapp-run-plan-preview",
		"--app", "7zr",
	}, &output)
	if err == nil || !strings.Contains(err.Error(), "requires a Runtime catalog q4-runnable GUI app") {
		t.Fatalf("expected known non-bundle app rejection, got %v", err)
	}
}

func TestQ4KnownPortableWinAppRunPlanPreviewCommandRejectsArguments(t *testing.T) {
	var output bytes.Buffer
	err := run([]string{
		"q4-known-portable-winapp-run-plan-preview",
		"--app", "org.xnix.external.notepadplusplus",
		"extra",
	}, &output)
	if err == nil || !strings.Contains(err.Error(), "does not accept positional arguments") {
		t.Fatalf("expected positional argument rejection, got %v", err)
	}
}
