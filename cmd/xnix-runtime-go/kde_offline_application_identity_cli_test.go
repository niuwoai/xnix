package main

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"
)

func TestKDEOfflineApplicationIdentityPreviewCommandUsesSampleFixture(t *testing.T) {
	var output bytes.Buffer
	err := run([]string{
		"kde-offline-application-identity-preview",
		"--registry", "../../runtime/recipes/registry.json",
		"--app", "org.xnix.sample.notepad",
	}, &output)
	if err != nil {
		t.Fatalf("run: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal: %v", err)
	}
	if payload["schema_version"] != "xnix.runtime.kde_offline_application_identity.v1" ||
		payload["application_id"] != "org.xnix.sample.notepad" ||
		payload["display_name"] != "Sample Notepad" ||
		payload["surface_count"] != float64(7) ||
		payload["recipe_digest_verified"] != true ||
		payload["cross_surface_identity_consistent"] != true ||
		payload["launch_enabled"] != false ||
		payload["live_tray_bridge_enabled"] != false ||
		payload["notification_sent"] != false ||
		payload["notification_delivery_enabled"] != false ||
		payload["execution_started"] != false ||
		payload["host_root_modified"] != false {
		t.Fatalf("unexpected command payload: %#v", payload)
	}
	assertKDEOfflineApplicationIdentityCLISafe(t, output.String())
}

func TestKDEOfflineApplicationIdentityPreviewCommandRejectsIncompleteArguments(t *testing.T) {
	for _, args := range [][]string{
		{"kde-offline-application-identity-preview"},
		{"kde-offline-application-identity-preview", "--registry", "../../runtime/recipes/registry.json"},
		{"kde-offline-application-identity-preview", "--registry", "../../runtime/recipes/registry.json", "--app", "org.xnix.sample.notepad", "extra"},
	} {
		if err := run(args, &bytes.Buffer{}); err == nil {
			t.Fatalf("accepted malformed args: %#v", args)
		}
	}
}

func assertKDEOfflineApplicationIdentityCLISafe(t *testing.T, text string) {
	t.Helper()
	serialized := strings.ToLower(text)
	for _, forbidden := range []string{"prefix", ".exe", "program files", "qemu-system", "proton", "wine ", "wine/", ".wine", "/home", "/users", "/private", "/var", "/opt", "/tmp", "file://"} {
		if strings.Contains(serialized, forbidden) {
			t.Fatalf("offline KDE identity CLI exposed forbidden term %q: %s", forbidden, text)
		}
	}
}
