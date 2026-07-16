package main

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"
)

func stageDeactivationReceipt(t *testing.T, registryPath, app, stagingRoot string) {
	t.Helper()
	var output bytes.Buffer
	if err := run([]string{"desktop-activation-stage", "--registry", registryPath, "--app", app, "--mode", "development", "--staging-root", stagingRoot}, &output); err != nil {
		t.Fatalf("stage activation receipt: %v", err)
	}
}

func decodeDeactivation(t *testing.T, output *bytes.Buffer) map[string]any {
	t.Helper()
	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("parse output: %v\n%s", err, output.String())
	}
	return payload
}

func TestDesktopDeactivationDryRunPreviewCLIMissingReceipt(t *testing.T) {
	registryPath, app := writeTestRepairGroupRegistry(t)
	var output bytes.Buffer
	if err := run([]string{"desktop-deactivation-dry-run-preview", "--registry", registryPath, "--app", app}, &output); err != nil {
		t.Fatalf("run deactivation preview: %v", err)
	}
	payload := decodeDeactivation(t, &output)
	if payload["schema_version"] != "xnix.runtime.desktop_deactivation_dry_run.v1" ||
		payload["overall_state"] != "blocked" ||
		payload["receipt_state"] != "missing" {
		t.Fatalf("missing-receipt deactivation should be blocked: %+v", payload)
	}
	for _, key := range []string{"file_deletion_enabled", "mime_defaults_written", "kde_cache_refreshed", "receipts_rewritten", "session_terminated", "target_path_exposed", "host_root_modified"} {
		if payload[key] != false {
			t.Fatalf("preview must keep %q disabled: %+v", key, payload)
		}
	}
}

func TestDesktopDeactivationDryRunPreviewCLIReceiptBackedRemovable(t *testing.T) {
	registryPath, app := writeTestRepairGroupRegistry(t)
	stagingRoot := t.TempDir()
	stageDeactivationReceipt(t, registryPath, app, stagingRoot)

	var output bytes.Buffer
	if err := run([]string{"desktop-deactivation-dry-run-preview", "--registry", registryPath, "--app", app, "--activation-root", stagingRoot}, &output); err != nil {
		t.Fatalf("run deactivation preview: %v", err)
	}
	payload := decodeDeactivation(t, &output)
	if payload["overall_state"] != "removable" || payload["receipt_state"] != "receipt-backed" {
		t.Fatalf("staged receipt should make surfaces removable: %+v", payload["overall_state"])
	}
	if strings.Contains(output.String(), stagingRoot) {
		t.Fatalf("CLI output must not expose the staging-root path: %s", output.String())
	}
}

func TestDesktopDeactivationDryRunPreviewCLIActiveSessionAndSharedMIMEBlock(t *testing.T) {
	registryPath, app := writeTestRepairGroupRegistry(t)
	stagingRoot := t.TempDir()
	stageDeactivationReceipt(t, registryPath, app, stagingRoot)

	var sessionOut bytes.Buffer
	if err := run([]string{"desktop-deactivation-dry-run-preview", "--registry", registryPath, "--app", app, "--activation-root", stagingRoot, "--active-session"}, &sessionOut); err != nil {
		t.Fatalf("run active-session preview: %v", err)
	}
	if !strings.Contains(sessionOut.String(), "active-session-evidence") || !strings.Contains(sessionOut.String(), "\"blocked\"") {
		t.Fatalf("active session should block removal: %s", sessionOut.String())
	}

	var mimeOut bytes.Buffer
	if err := run([]string{"desktop-deactivation-dry-run-preview", "--registry", registryPath, "--app", app, "--activation-root", stagingRoot, "--shared-mime", "application/x-xnix-abc"}, &mimeOut); err != nil {
		t.Fatalf("run shared-mime preview: %v", err)
	}
	if !strings.Contains(mimeOut.String(), "shared-mime-association") {
		t.Fatalf("shared MIME association should block MIME removal: %s", mimeOut.String())
	}
}

func TestDesktopDeactivationDryRunPreviewCLIRequiresSource(t *testing.T) {
	registryPath, app := writeTestRepairGroupRegistry(t)
	tests := [][]string{
		{"desktop-deactivation-dry-run-preview"},
		{"desktop-deactivation-dry-run-preview", "--registry", registryPath},
		{"desktop-deactivation-dry-run-preview", "--registry", registryPath, "--app", app, "extra"},
	}
	for _, args := range tests {
		var output bytes.Buffer
		if err := run(args, &output); err == nil {
			t.Fatalf("expected CLI args to fail: %+v", args)
		}
	}
}
