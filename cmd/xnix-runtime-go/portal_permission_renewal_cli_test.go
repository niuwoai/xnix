package main

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestPortalPermissionRenewalPreviewCLI(t *testing.T) {
	registryPath, app := writeTestRepairGroupRegistry(t)
	stateRoot := t.TempDir()
	fileOpen := recordPortalPermissionForRenewal(t, stateRoot, app, "file-open", "complete")
	_ = fileOpen
	recordPortalPermissionForRenewal(t, stateRoot, app, "uri-open", "complete")
	recordPortalPermissionForRenewal(t, stateRoot, app, "print", "pending")
	recordPortalPermissionForRenewal(t, stateRoot, app, "clipboard", "denied")
	recordPortalPermissionForRenewal(t, stateRoot, app, "screenshot", "revoked")

	var output bytes.Buffer
	if err := run([]string{
		"portal-permission-renewal-preview",
		"--registry", registryPath,
		"--app", app,
		"--state-root", stateRoot,
		"--expiring", "uri-open",
	}, &output); err != nil {
		t.Fatalf("portal permission renewal returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["schema_version"] != "xnix.runtime.portal_permission_renewal.v1" ||
		payload["request_type"] != "portal-permission-renewal-preview" ||
		payload["runtime_method"] != "GetPortalPermissionRenewal" ||
		payload["read_method"] != "GetPortalPermissionRenewalPreview" ||
		payload["real_portal_call_enabled"] != false ||
		payload["permission_grant_enabled"] != false ||
		payload["permission_revoke_enabled"] != false ||
		payload["receipt_write_enabled"] != false ||
		payload["settings_persisted"] != false ||
		payload["execution_approved"] != false ||
		payload["host_root_modified"] != false {
		t.Fatalf("unexpected renewal payload: %#v", payload)
	}
	counts := payload["counts"].(map[string]any)
	if counts["current"] != float64(2) ||
		counts["expiring_soon"] != float64(1) ||
		counts["needs_review"] != float64(1) ||
		counts["denied"] != float64(1) ||
		counts["revoked"] != float64(1) ||
		counts["blocked_by_policy"] != float64(2) {
		t.Fatalf("unexpected renewal counts: %#v", counts)
	}
	if strings.Contains(output.String(), stateRoot) || strings.Contains(output.String(), registryPath) {
		t.Fatalf("renewal output exposed local paths: %s", output.String())
	}
	assertPortalPermissionRenewalCLISafe(t, output.String())
}

func TestPortalPermissionRenewalPreviewCLIMissingStateRootDoesNotCreate(t *testing.T) {
	registryPath, app := writeTestRepairGroupRegistry(t)
	missingRoot := filepath.Join(t.TempDir(), "missing-state")
	var output bytes.Buffer
	if err := run([]string{
		"portal-permission-renewal-preview",
		"--registry", registryPath,
		"--app", app,
		"--state-root", missingRoot,
	}, &output); err != nil {
		t.Fatalf("portal permission renewal returned error: %v", err)
	}
	if _, err := os.Stat(missingRoot); !os.IsNotExist(err) {
		t.Fatalf("read-only renewal preview should not create missing state root: %v", err)
	}
	assertPortalPermissionRenewalCLISafe(t, output.String())
}

func TestPortalPermissionRenewalPreviewCLIMalformedLedgerIsSurfaced(t *testing.T) {
	registryPath, app := writeTestRepairGroupRegistry(t)
	stateRoot := t.TempDir()
	requests := filepath.Join(stateRoot, "portal-requests")
	if err := os.MkdirAll(requests, 0o700); err != nil {
		t.Fatalf("MkdirAll: %v", err)
	}
	if err := os.WriteFile(filepath.Join(requests, "bad-receipt.json"), []byte("{not-json"), 0o600); err != nil {
		t.Fatalf("WriteFile malformed receipt: %v", err)
	}
	var output bytes.Buffer
	if err := run([]string{
		"portal-permission-renewal-preview",
		"--registry", registryPath,
		"--app", app,
		"--state-root", stateRoot,
	}, &output); err != nil {
		t.Fatalf("portal permission renewal returned error: %v", err)
	}
	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	counts := payload["counts"].(map[string]any)
	if counts["malformed_receipts"] != float64(1) {
		t.Fatalf("renewal preview must surface malformed ledger records: %#v", counts)
	}
	assertPortalPermissionRenewalCLISafe(t, output.String())
}

func TestPortalPermissionRenewalPreviewCLIRequiresRecipeSource(t *testing.T) {
	registryPath, app := writeTestRepairGroupRegistry(t)
	cases := [][]string{
		{"portal-permission-renewal-preview"},
		{"portal-permission-renewal-preview", "--registry", registryPath},
		{"portal-permission-renewal-preview", "--registry", registryPath, "--app", app, "--recipe", registryPath},
		{"portal-permission-renewal-preview", "--registry", registryPath, "--app", app, "extra"},
	}
	for _, args := range cases {
		var output bytes.Buffer
		if err := run(args, &output); err == nil {
			t.Fatalf("expected portal-permission-renewal-preview to reject args: %#v", args)
		}
	}
}

func recordPortalPermissionForRenewal(t *testing.T, stateRoot, app, operation, final string) string {
	t.Helper()
	var createOutput bytes.Buffer
	if err := run([]string{
		"portal-request-record",
		"--state-root", stateRoot,
		"--action", "create",
		"--app", app,
		"--operation", operation,
	}, &createOutput); err != nil {
		t.Fatalf("create Portal permission %s: %v", operation, err)
	}
	var created map[string]any
	if err := json.Unmarshal(createOutput.Bytes(), &created); err != nil {
		t.Fatalf("Unmarshal create: %v", err)
	}
	request := created["request"].(map[string]any)
	handle := request["handle_token"].(string)
	switch final {
	case "pending":
		return handle
	case "complete":
		resolvePortalPermissionForRenewal(t, stateRoot, handle, "granted")
		completePortalPermissionForRenewal(t, stateRoot, handle)
	case "denied":
		resolvePortalPermissionForRenewal(t, stateRoot, handle, "denied")
	case "revoked":
		var output bytes.Buffer
		if err := run([]string{"portal-request-record", "--state-root", stateRoot, "--action", "cancel", "--handle-token", handle}, &output); err != nil {
			t.Fatalf("cancel Portal permission %s: %v", operation, err)
		}
	}
	return handle
}

func resolvePortalPermissionForRenewal(t *testing.T, stateRoot, handle, outcome string) {
	t.Helper()
	var output bytes.Buffer
	if err := run([]string{"portal-request-record", "--state-root", stateRoot, "--action", "resolve", "--handle-token", handle, "--outcome", outcome}, &output); err != nil {
		t.Fatalf("resolve Portal permission %s: %v", handle, err)
	}
}

func completePortalPermissionForRenewal(t *testing.T, stateRoot, handle string) {
	t.Helper()
	var output bytes.Buffer
	if err := run([]string{"portal-request-record", "--state-root", stateRoot, "--action", "complete", "--handle-token", handle}, &output); err != nil {
		t.Fatalf("complete Portal permission %s: %v", handle, err)
	}
}

func assertPortalPermissionRenewalCLISafe(t *testing.T, text string) {
	t.Helper()
	lower := strings.ToLower(text)
	for _, forbidden := range []string{"prefix", ".exe", "program files", "qemu-system", "proton", "wine ", "wine/", ".wine", "/home", "/users", "/private", "file://"} {
		if strings.Contains(lower, forbidden) {
			t.Fatalf("portal permission renewal CLI exposed forbidden term %q: %s", forbidden, text)
		}
	}
}
