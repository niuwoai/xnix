package main

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"
)

func TestKDENotificationDigestPreviewCLI(t *testing.T) {
	registryPath, app := writeTestRepairGroupRegistry(t)
	args := []string{"kde-notification-digest-preview", "--registry", registryPath, "--app", app,
		"--event", "needs-review:approval-required", "--event", "blocked-action:execution-blocked",
		"--event", "permission-attention:permission-review", "--event", "diagnostic-issue:diagnostic-failed",
		"--event", "snapshot-warning:snapshot-missing", "--event", "readiness-change:readiness-pending",
		"--event", "needs-review:approval-required"}
	var output bytes.Buffer
	if err := run(args, &output); err != nil {
		t.Fatalf("run digest preview: %v", err)
	}
	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("parse output: %v", err)
	}
	if payload["schema_version"] != "xnix.runtime.kde_notification_digest.v1" ||
		payload["request_type"] != "kde-notification-digest-preview" || payload["review_only"] != true || payload["digest_ready"] != true {
		t.Fatalf("unexpected payload identity: %+v", payload)
	}
	counts := payload["counts"].(map[string]any)
	if counts["input_event_count"] != float64(7) || counts["digest_entry_count"] != float64(6) || counts["duplicate_event_count"] != float64(1) {
		t.Fatalf("unexpected counts: %+v", counts)
	}
	for _, key := range []string{"notifications_sent", "live_tray_bridge_enabled", "request_objects_created", "permission_grant_enabled", "backend_launch_enabled", "backend_process_started", "state_root_path_exposed", "raw_command_exposed", "backend_details_exposed", "host_root_modified"} {
		if payload[key] != false {
			t.Fatalf("digest preview must keep %q disabled: %+v", key, payload)
		}
	}
	for _, forbidden := range []string{".exe", "program files", "qemu-system", "proton", "wine ", ".wine", "/home", "/users", "/private", "file://", "token=", "secret"} {
		if strings.Contains(strings.ToLower(output.String()), forbidden) {
			t.Fatalf("digest output exposed forbidden text %q", forbidden)
		}
	}
}

func TestKDENotificationDigestPreviewCLIRejectsInvalidArguments(t *testing.T) {
	registryPath, app := writeTestRepairGroupRegistry(t)
	for _, args := range [][]string{
		{"kde-notification-digest-preview"},
		{"kde-notification-digest-preview", "--registry", registryPath, "--app", app},
		{"kde-notification-digest-preview", "--registry", registryPath, "--app", app, "--event", "bad"},
		{"kde-notification-digest-preview", "--registry", registryPath, "--app", app, "--event", "blocked-action:unknown"},
	} {
		var output bytes.Buffer
		if err := run(args, &output); err == nil {
			t.Fatalf("expected arguments to fail: %+v", args)
		}
	}
}
