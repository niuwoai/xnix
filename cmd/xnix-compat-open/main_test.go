package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestCompatOpenBuildsRuntimeFileOpenPreview(t *testing.T) {
	registryPath := writeCompatOpenRegistry(t)

	var output bytes.Buffer
	err := run([]string{
		"--registry", registryPath,
		"--app", "org.example.notes",
		"file:///home/test/Documents/report.txt",
	}, &output)
	if err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["schema_version"] != "xnix.runtime.file_open.v1" ||
		payload["request_type"] != "file-open-preview" ||
		payload["source"] != "dolphin-service-menu" ||
		payload["application_id"] != "org.example.notes" ||
		payload["display_name"] != "Example Notes" ||
		payload["runtime_method"] != "Launch" ||
		payload["portal_required"] != true ||
		payload["portal_interface"] != "org.freedesktop.portal.FileChooser" ||
		payload["portal_method"] != "OpenFile" ||
		payload["file_count"] != float64(1) ||
		payload["selected_extension"] != ".txt" ||
		payload["selection_mode"] != "explicit-application" ||
		payload["runtime_owned"] != true ||
		payload["backend_launch_enabled"] != false ||
		payload["direct_host_file_access"] != false ||
		payload["host_root_modified"] != false ||
		payload["backend_details_exposed"] != false {
		t.Fatalf("unexpected file-open payload: %#v", payload)
	}
	action := payload["action"].(map[string]any)
	argv := action["argv"].([]any)
	if action["type"] != "runtime-file-open" ||
		argv[0] != "xnix-compat-open" ||
		argv[1] != "--app" ||
		argv[2] != "org.example.notes" ||
		argv[3] != "%U" {
		t.Fatalf("unexpected action payload: %#v", action)
	}
	assertCompatOpenSafe(t, output.String())
}

func TestCompatOpenSelectsRecipeByExtension(t *testing.T) {
	registryPath := writeCompatOpenRegistry(t)

	var output bytes.Buffer
	err := run([]string{
		"--registry", registryPath,
		"file:///home/test/Documents/ledger.abc",
	}, &output)
	if err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["application_id"] != "org.example.ledger" ||
		payload["selected_extension"] != ".abc" ||
		payload["selection_mode"] != "extension-match" {
		t.Fatalf("unexpected selected recipe payload: %#v", payload)
	}
}

func TestCompatOpenAcceptsRecipeDirAlias(t *testing.T) {
	registryPath := writeCompatOpenRegistry(t)

	var output bytes.Buffer
	err := run([]string{
		"--recipe-dir", filepath.Dir(registryPath),
		"file:///home/test/Documents/report.txt",
	}, &output)
	if err != nil {
		t.Fatalf("run returned error: %v", err)
	}
	if !strings.Contains(output.String(), `"application_id": "org.example.notes"`) {
		t.Fatalf("recipe-dir alias did not load registry: %s", output.String())
	}
}

func TestCompatOpenExecuteDelegatesToManagedLauncher(t *testing.T) {
	registryPath := writeCompatOpenRegistry(t)
	documentPath := filepath.Join(t.TempDir(), "report with spaces.txt")
	if err := os.WriteFile(documentPath, []byte("file-open execution fixture"), 0o600); err != nil {
		t.Fatalf("WriteFile document returned error: %v", err)
	}
	launcherLog := filepath.Join(t.TempDir(), "launcher-argv.log")
	launcherPath := writeFakeCompatOpenLauncher(t, launcherLog)

	var output bytes.Buffer
	err := run([]string{
		"--registry", registryPath,
		"--execute",
		"--launcher-bin", launcherPath,
		"--guest-boundary", "managed-known-app-guest-smoke",
		"--state-root", "/tmp/xnix-state",
		"--receipt-id", "known-app-launch-receipt-org-example-notes",
		"--review-receipt-id", "known-app-session-gated-launch-review-receipt-org-example-notes",
		"--session-id", "controlled-execution-session-org-example-notes",
		"--window-match", "report with spaces.txt",
		"--timeout", "5s",
		"file://" + filepath.ToSlash(documentPath),
	}, &output)
	if err != nil {
		t.Fatalf("run returned error: %v", err)
	}
	if !strings.Contains(output.String(), `"schema_version":"fake.launcher.v1"`) ||
		!strings.Contains(output.String(), `"status":"passed"`) {
		t.Fatalf("unexpected managed launcher output: %s", output.String())
	}
	launcherArgs, err := os.ReadFile(launcherLog)
	if err != nil {
		t.Fatalf("ReadFile launcher log returned error: %v", err)
	}
	argvText := string(launcherArgs)
	for _, token := range []string{
		"--app\norg.example.notes\n",
		"--guest-boundary\nmanaged-known-app-guest-smoke\n",
		"--state-root\n/tmp/xnix-state\n",
		"--receipt-id\nknown-app-launch-receipt-org-example-notes\n",
		"--review-receipt-id\nknown-app-session-gated-launch-review-receipt-org-example-notes\n",
		"--session-id\ncontrolled-execution-session-org-example-notes\n",
		"--window-match\nreport with spaces.txt\n",
		"--timeout\n5s\n",
		"--file-argument\n" + documentPath + "\n",
	} {
		if !strings.Contains(argvText, token) {
			t.Fatalf("managed launcher argv missing %q in %s", token, argvText)
		}
	}
}

func TestCompatOpenExecuteAcceptsOwnerFileArgumentAlias(t *testing.T) {
	registryPath := writeCompatOpenRegistry(t)
	documentPath := filepath.Join(t.TempDir(), "owner-report.txt")
	if err := os.WriteFile(documentPath, []byte("owner file-open execution fixture"), 0o600); err != nil {
		t.Fatalf("WriteFile document returned error: %v", err)
	}
	launcherLog := filepath.Join(t.TempDir(), "launcher-argv.log")
	launcherPath := writeFakeCompatOpenLauncher(t, launcherLog)
	t.Setenv("XNIX_COMPAT_OPEN_REGISTRY", registryPath)
	t.Setenv("XNIX_COMPAT_LAUNCH", launcherPath)
	t.Setenv("XNIX_COMPAT_OPEN_EXECUTE", "1")

	var output bytes.Buffer
	err := run([]string{
		"--app", "org.example.notes",
		"--file-argument", documentPath,
	}, &output)
	if err != nil {
		t.Fatalf("run returned error: %v", err)
	}
	if !strings.Contains(output.String(), `"schema_version":"fake.launcher.v1"`) {
		t.Fatalf("unexpected managed launcher output: %s", output.String())
	}
	launcherArgs, err := os.ReadFile(launcherLog)
	if err != nil {
		t.Fatalf("ReadFile launcher log returned error: %v", err)
	}
	if strings.Count(string(launcherArgs), "--file-argument\n") != 1 ||
		!strings.Contains(string(launcherArgs), "--file-argument\n"+documentPath+"\n") {
		t.Fatalf("owner file-argument alias did not pass exactly one file to launcher: %s", string(launcherArgs))
	}
}

func TestCompatOpenExecuteUsesRuntimeOwnerEnvironmentDefaults(t *testing.T) {
	registryPath := writeCompatOpenRegistry(t)
	documentPath := filepath.Join(t.TempDir(), "owner-env-report.txt")
	if err := os.WriteFile(documentPath, []byte("owner environment file-open execution fixture"), 0o600); err != nil {
		t.Fatalf("WriteFile document returned error: %v", err)
	}
	launcherLog := filepath.Join(t.TempDir(), "launcher-argv.log")
	launcherPath := writeFakeCompatOpenLauncher(t, launcherLog)
	t.Setenv("XNIX_COMPAT_OPEN_REGISTRY", registryPath)
	t.Setenv("XNIX_COMPAT_OPEN_EXECUTE", "1")
	t.Setenv("XNIX_COMPAT_LAUNCH", launcherPath)
	t.Setenv("XNIX_COMPAT_OPEN_LAUNCHER_REGISTRY", "/runtime/recipes/registry.json")
	t.Setenv("XNIX_COMPAT_OPEN_CACHE_ROOT", "/runtime/cache")
	t.Setenv("XNIX_COMPAT_OPEN_GUEST_BOUNDARY", "managed-known-app-guest-smoke")
	t.Setenv("XNIX_COMPAT_OPEN_STATE_ROOT", "/runtime/state")
	t.Setenv("XNIX_COMPAT_OPEN_RECEIPT_ID", "known-app-launch-receipt-org-example-notes")
	t.Setenv("XNIX_COMPAT_OPEN_REVIEW_RECEIPT_ID", "known-app-session-gated-launch-review-receipt-org-example-notes")
	t.Setenv("XNIX_COMPAT_OPEN_SESSION_ID", "controlled-execution-session-org-example-notes")
	t.Setenv("XNIX_COMPAT_OPEN_GUEST_HOST", "127.0.0.1")
	t.Setenv("XNIX_COMPAT_OPEN_GUEST_PORT", "40229")
	t.Setenv("XNIX_COMPAT_OPEN_GUEST_USER", "root")
	t.Setenv("XNIX_COMPAT_OPEN_GUEST_KEY", "/runtime/keys/qemu")
	t.Setenv("XNIX_COMPAT_OPEN_GUEST_REMOTE_DIR", "/tmp/xnix-owner-open")
	t.Setenv("XNIX_COMPAT_OPEN_GUEST_SSH", "ssh")
	t.Setenv("XNIX_COMPAT_OPEN_GUEST_SCP", "scp")
	t.Setenv("XNIX_COMPAT_OPEN_GUEST_XWININFO", "xwininfo")
	t.Setenv("XNIX_COMPAT_OPEN_GUEST_DISPLAY", "10.0.2.2:106")
	t.Setenv("XNIX_COMPAT_OPEN_HOST_DISPLAY", ":106")
	t.Setenv("XNIX_COMPAT_OPEN_WINDOW_MATCH", "owner-env-report.txt")
	t.Setenv("XNIX_COMPAT_OPEN_TIMEOUT", "7s")
	t.Setenv("XNIX_COMPAT_OPEN_GUI_WAIT", "3s")

	var output bytes.Buffer
	err := run([]string{
		"--app", "org.example.notes",
		"file://" + filepath.ToSlash(documentPath),
	}, &output)
	if err != nil {
		t.Fatalf("run returned error: %v", err)
	}
	if !strings.Contains(output.String(), `"schema_version":"fake.launcher.v1"`) {
		t.Fatalf("unexpected managed launcher output: %s", output.String())
	}
	launcherArgs, err := os.ReadFile(launcherLog)
	if err != nil {
		t.Fatalf("ReadFile launcher log returned error: %v", err)
	}
	argvText := string(launcherArgs)
	for _, token := range []string{
		"--app\norg.example.notes\n",
		"--registry\n/runtime/recipes/registry.json\n",
		"--cache-root\n/runtime/cache\n",
		"--guest-boundary\nmanaged-known-app-guest-smoke\n",
		"--state-root\n/runtime/state\n",
		"--receipt-id\nknown-app-launch-receipt-org-example-notes\n",
		"--review-receipt-id\nknown-app-session-gated-launch-review-receipt-org-example-notes\n",
		"--session-id\ncontrolled-execution-session-org-example-notes\n",
		"--host\n127.0.0.1\n",
		"--port\n40229\n",
		"--user\nroot\n",
		"--key\n/runtime/keys/qemu\n",
		"--remote-dir\n/tmp/xnix-owner-open\n",
		"--ssh\nssh\n",
		"--scp\nscp\n",
		"--xwininfo\nxwininfo\n",
		"--guest-display\n10.0.2.2:106\n",
		"--host-display\n:106\n",
		"--window-match\nowner-env-report.txt\n",
		"--timeout\n7s\n",
		"--gui-wait\n3s\n",
		"--file-argument\n" + documentPath + "\n",
	} {
		if !strings.Contains(argvText, token) {
			t.Fatalf("managed launcher argv missing %q in %s", token, argvText)
		}
	}
}

func TestCompatOpenExecuteFlagsOverrideRuntimeOwnerEnvironmentDefaults(t *testing.T) {
	registryPath := writeCompatOpenRegistry(t)
	documentPath := filepath.Join(t.TempDir(), "owner-env-override.txt")
	if err := os.WriteFile(documentPath, []byte("owner environment override fixture"), 0o600); err != nil {
		t.Fatalf("WriteFile document returned error: %v", err)
	}
	launcherLog := filepath.Join(t.TempDir(), "launcher-argv.log")
	launcherPath := writeFakeCompatOpenLauncher(t, launcherLog)
	t.Setenv("XNIX_COMPAT_OPEN_REGISTRY", registryPath)
	t.Setenv("XNIX_COMPAT_OPEN_EXECUTE", "1")
	t.Setenv("XNIX_COMPAT_LAUNCH", launcherPath)
	t.Setenv("XNIX_COMPAT_OPEN_WINDOW_MATCH", "environment-window")
	t.Setenv("XNIX_COMPAT_OPEN_TIMEOUT", "7s")

	var output bytes.Buffer
	err := run([]string{
		"--app", "org.example.notes",
		"--window-match", "flag-window",
		"--timeout", "9s",
		"file://" + filepath.ToSlash(documentPath),
	}, &output)
	if err != nil {
		t.Fatalf("run returned error: %v", err)
	}
	launcherArgs, err := os.ReadFile(launcherLog)
	if err != nil {
		t.Fatalf("ReadFile launcher log returned error: %v", err)
	}
	argvText := string(launcherArgs)
	if !strings.Contains(argvText, "--window-match\nflag-window\n") ||
		!strings.Contains(argvText, "--timeout\n9s\n") ||
		strings.Contains(argvText, "environment-window") ||
		strings.Contains(argvText, "--timeout\n7s\n") {
		t.Fatalf("flags did not override environment defaults: %s", argvText)
	}
}

func TestCompatOpenExecuteRejectsRemoteFileURIHosts(t *testing.T) {
	registryPath := writeCompatOpenRegistry(t)
	launcherPath := writeFakeCompatOpenLauncher(t, filepath.Join(t.TempDir(), "launcher-argv.log"))

	var output bytes.Buffer
	err := run([]string{
		"--registry", registryPath,
		"--execute",
		"--launcher-bin", launcherPath,
		"file://remote-host/home/test/Documents/report.txt",
	}, &output)
	if err == nil || !strings.Contains(err.Error(), "file URI host must be empty or localhost") {
		t.Fatalf("expected remote file URI host rejection, got %v", err)
	}
}

func TestCompatOpenRejectsUnsafeRequests(t *testing.T) {
	registryPath := writeCompatOpenRegistry(t)
	tests := []struct {
		name string
		args []string
		want string
	}{
		{
			name: "missing file uri",
			args: []string{"--registry", registryPath},
			want: "requires at least one file URI",
		},
		{
			name: "non file uri",
			args: []string{"--registry", registryPath, "https://example.invalid/report.txt"},
			want: "only file URIs are accepted",
		},
		{
			name: "unknown app",
			args: []string{"--registry", registryPath, "--app", "org.example.missing", "file:///home/test/Documents/report.txt"},
			want: "unknown application",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var output bytes.Buffer
			err := run(test.args, &output)
			if err == nil || !strings.Contains(err.Error(), test.want) {
				t.Fatalf("expected %q rejection, got %v", test.want, err)
			}
		})
	}
}

func writeCompatOpenRegistry(t *testing.T) string {
	t.Helper()
	root := t.TempDir()
	ledgerData := []byte(`{"id":"org.example.ledger","name":"Example Ledger","icon":"office-chart-area","mode":"automatic","supported_extensions":[".abc"]}`)
	notesData := []byte(`{"id":"org.example.notes","name":"Example Notes","icon":"accessories-text-editor","mode":"automatic","supported_extensions":[".txt"]}`)
	writeRecipe := func(name string, data []byte) string {
		t.Helper()
		if err := os.WriteFile(filepath.Join(root, name), data, 0o600); err != nil {
			t.Fatalf("WriteFile %s returned error: %v", name, err)
		}
		sum := sha256.Sum256(data)
		return hex.EncodeToString(sum[:])
	}
	ledgerDigest := writeRecipe("org.example.ledger.json", ledgerData)
	notesDigest := writeRecipe("org.example.notes.json", notesData)
	registryData := []byte(`{"schema_version":1,"registry_name":"test-registry","recipes":[{"id":"org.example.ledger","path":"org.example.ledger.json","sha256":"` + ledgerDigest + `","signature_status":"development-only"},{"id":"org.example.notes","path":"org.example.notes.json","sha256":"` + notesDigest + `","signature_status":"development-only"}]}`)
	registryPath := filepath.Join(root, "registry.json")
	if err := os.WriteFile(registryPath, registryData, 0o600); err != nil {
		t.Fatalf("WriteFile registry returned error: %v", err)
	}
	return registryPath
}

func writeFakeCompatOpenLauncher(t *testing.T, logPath string) string {
	t.Helper()
	launcherPath := filepath.Join(t.TempDir(), "xnix-compat-launch")
	script := "#!/bin/sh\n: > \"$XNIX_TEST_LAUNCHER_LOG\"\nfor arg in \"$@\"; do printf '%s\\n' \"$arg\" >> \"$XNIX_TEST_LAUNCHER_LOG\"; done\nprintf '{\"schema_version\":\"fake.launcher.v1\",\"status\":\"passed\"}\\n'\n"
	if err := os.WriteFile(launcherPath, []byte(script), 0o700); err != nil {
		t.Fatalf("WriteFile fake launcher returned error: %v", err)
	}
	t.Setenv("XNIX_TEST_LAUNCHER_LOG", logPath)
	return launcherPath
}

func assertCompatOpenSafe(t *testing.T, text string) {
	t.Helper()
	for _, forbidden := range []string{
		"backend_process_started",
		"docker_socket_mounted\": true",
		"host_networking_required\": true",
		"broad_host_mount_required\": true",
		"host_root_modified\": true",
		"direct_host_file_access\": true",
		"backend_details_exposed\": true",
	} {
		if strings.Contains(text, forbidden) {
			t.Fatalf("file-open command exposed unsafe term %q in %s", forbidden, text)
		}
	}
}
