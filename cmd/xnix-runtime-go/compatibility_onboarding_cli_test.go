package main

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"
)

func TestCompatibilityOnboardingChecklistPreviewCommand(t *testing.T) {
	registryPath, app := writeTestRepairGroupRegistry(t)
	var output bytes.Buffer
	err := run([]string{
		"compatibility-onboarding-checklist-preview",
		"--registry", registryPath,
		"--app", app,
		"--root", "../..",
	}, &output)
	if err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["schema_version"] != "xnix.runtime.compatibility_onboarding_checklist.v1" ||
		payload["request_type"] != "compatibility-onboarding-checklist-preview" ||
		payload["checklist_type"] != "first-run-compatibility-onboarding" ||
		payload["runtime_method"] != "GetCompatibilityOnboardingChecklist" ||
		payload["read_method"] != "GetCompatibilityOnboardingChecklistPreview" ||
		payload["desktop"] != "KDE Plasma" ||
		payload["section_count"] != float64(9) ||
		payload["ready"] != false {
		t.Fatalf("unexpected onboarding checklist payload: %#v", payload)
	}
	application := payload["application"].(map[string]any)
	if application["id"] != app ||
		application["name"] == "" ||
		application["desktop_file"] == "" ||
		application["runtime_mode"] != "automatic" {
		t.Fatalf("unexpected onboarding application: %#v", application)
	}
	states := payload["states"].(map[string]any)
	for _, key := range []string{"ready", "needs_review", "missing_evidence", "blocked", "not_yet_implemented"} {
		if states[key].(float64) == 0 {
			t.Fatalf("onboarding checklist did not expose state %q: %#v", key, states)
		}
	}
	if payload["runtime_owned"] != true ||
		payload["go_runtime_backed"] != true ||
		payload["kde_policy_owner"] != false ||
		payload["runtime_write_methods_enabled"] != false ||
		payload["request_objects_created"] != false ||
		payload["permission_grants_created"] != false ||
		payload["artifact_staged"] != false ||
		payload["settings_persisted"] != false ||
		payload["backend_process_started"] != false ||
		payload["launch_enabled"] != false ||
		payload["network_required"] != false ||
		payload["host_package_manager_invoked"] != false ||
		payload["host_root_modified"] != false ||
		payload["privileged_container_required"] != false ||
		payload["ai_provider_called"] != false ||
		payload["state_root_path_exposed"] != false ||
		payload["raw_executable_exposed"] != false ||
		payload["raw_command_exposed"] != false ||
		payload["file_content_read"] != false ||
		payload["backend_details_exposed"] != false {
		t.Fatalf("onboarding checklist enabled unsafe behavior: %#v", payload)
	}
	sections := payload["sections"].([]any)
	if len(sections) != 9 {
		t.Fatalf("unexpected section count: %#v", payload)
	}
	assertOnboardingCLISection(t, sections, "runtime-owner-readiness", "needs-review")
	assertOnboardingCLISection(t, sections, "recipe-trust", "needs-review")
	assertOnboardingCLISection(t, sections, "artifact-staging", "missing-evidence")
	assertOnboardingCLISection(t, sections, "portal-review", "needs-review")
	assertOnboardingCLISection(t, sections, "diagnostics-privacy", "ready")
	assertOnboardingCLISection(t, sections, "kde-entry-points", "ready")
	assertOnboardingCLISection(t, sections, "production-activation", "not-yet-implemented")
	if !strings.Contains(output.String(), "runtime-owner-readiness-preview") ||
		!strings.Contains(output.String(), "application-readiness-preview") ||
		!strings.Contains(output.String(), "portal-access-policy-preview") ||
		!strings.Contains(output.String(), "snapshot-plan-preview") ||
		!strings.Contains(output.String(), "ai-diagnostic-input-preview") ||
		!strings.Contains(output.String(), "desktop-safety-policy-preview") {
		t.Fatalf("onboarding checklist missed expected read-model links: %s", output.String())
	}
	assertKDEShellPayloadSafe(t, output.String())
}

func TestCompatibilityOnboardingChecklistPreviewCommandRejectsMalformedInputs(t *testing.T) {
	registryPath, app := writeTestRepairGroupRegistry(t)
	var output bytes.Buffer
	if err := run([]string{"compatibility-onboarding-checklist-preview", "--registry", registryPath, "--root", "../.."}, &output); err == nil {
		t.Fatalf("command accepted registry input without --app")
	}
	output.Reset()
	if err := run([]string{"compatibility-onboarding-checklist-preview", "--registry", registryPath, "--app", app, "--root", "../..", "--portal-operation", "full-disk"}, &output); err == nil {
		t.Fatalf("command accepted unsupported Portal operation")
	}
	output.Reset()
	if err := run([]string{"compatibility-onboarding-checklist-preview", "--registry", registryPath, "--app", app, "--root", "../..", "extra"}, &output); err == nil {
		t.Fatalf("command accepted positional arguments")
	}
}

func assertOnboardingCLISection(t *testing.T, sections []any, id string, state string) {
	t.Helper()
	for _, value := range sections {
		section := value.(map[string]any)
		if section["id"] == id {
			if section["state"] != state ||
				section["runtime_owned"] != true ||
				section["go_runtime_backed"] != true ||
				section["kde_policy_owner"] != false ||
				section["side_effects_enabled"] != false ||
				section["host_root_modified"] != false ||
				section["backend_details_exposed"] != false {
				t.Fatalf("unexpected onboarding CLI section %q: %#v", id, section)
			}
			return
		}
	}
	t.Fatalf("missing onboarding CLI section %q", id)
}
