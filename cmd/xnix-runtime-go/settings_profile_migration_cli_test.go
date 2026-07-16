package main

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"
)

func TestSettingsProfileMigrationPreviewCLI(t *testing.T) {
	registryPath, app := writeTestRepairGroupRegistry(t)

	var output bytes.Buffer
	if err := run([]string{
		"settings-profile-migration-preview",
		"--registry", registryPath,
		"--app", app,
		"--from-schema", "xnix.runtime.settings.v0",
		"--to-schema", "xnix.runtime.settings.v1",
		"--blocked-setting", "devices.camera",
	}, &output); err != nil {
		t.Fatalf("settings profile migration returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["schema_version"] != "xnix.runtime.settings_profile_migration.v1" ||
		payload["request_type"] != "settings-profile-migration-preview" ||
		payload["runtime_method"] != "GetCompatibilitySettingsProfileMigration" ||
		payload["read_method"] != "GetCompatibilitySettingsProfileMigrationPreview" ||
		payload["schema_state"] != "old-schema" ||
		payload["settings_persisted"] != false ||
		payload["resource_grant_enabled"] != false ||
		payload["real_portal_call_enabled"] != false ||
		payload["backend_launch_enabled"] != false ||
		payload["host_root_modified"] != false ||
		payload["backend_details_exposed"] != false {
		t.Fatalf("unexpected settings migration payload: %#v", payload)
	}
	if payload["row_count"] != float64(8) {
		t.Fatalf("unexpected row count: %#v", payload)
	}
	counts := payload["counts"].(map[string]any)
	if counts["blocked"] != float64(1) || counts["review_required"] != float64(5) {
		t.Fatalf("unexpected counts: %#v", counts)
	}
	assertSettingsProfileMigrationCLISafe(t, output.String())
}

func TestSettingsProfileMigrationPreviewCLIRejectsUnsafeArguments(t *testing.T) {
	registryPath, app := writeTestRepairGroupRegistry(t)
	cases := [][]string{
		{"settings-profile-migration-preview"},
		{"settings-profile-migration-preview", "--registry", registryPath},
		{"settings-profile-migration-preview", "--registry", registryPath, "--app", app, "--recipe", registryPath},
		{"settings-profile-migration-preview", "--registry", registryPath, "--app", app, "--blocked-setting", "unknown.setting"},
		{"settings-profile-migration-preview", "--registry", registryPath, "--app", app, "extra"},
	}
	for _, args := range cases {
		var output bytes.Buffer
		if err := run(args, &output); err == nil {
			t.Fatalf("expected settings-profile-migration-preview to reject args: %#v", args)
		}
	}
}

func assertSettingsProfileMigrationCLISafe(t *testing.T, text string) {
	t.Helper()
	lower := strings.ToLower(text)
	for _, forbidden := range []string{"prefix", ".exe", "program files", "qemu-system", "proton", "wine ", "wine/", ".wine", "/home", "/users", "/private", "file://", "token=", "secret"} {
		if strings.Contains(lower, forbidden) {
			t.Fatalf("settings migration CLI exposed forbidden term %q: %s", forbidden, text)
		}
	}
}
