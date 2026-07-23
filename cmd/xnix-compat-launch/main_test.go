package main

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"
)

func TestCompatLaunchUsesKnownAppLaunchBridge(t *testing.T) {
	tempDir := t.TempDir()

	var output bytes.Buffer
	err := run([]string{
		"--app", "7zr",
		"--cache-root", tempDir,
	}, &output)
	if err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["schema_version"] != "xnix.runtime.known_windows_app_launch_bridge.v1" ||
		payload["request_type"] != "windows-known-app-launch-bridge-preview" ||
		payload["source"] != "windows-known-app-dispatch-preview" ||
		payload["status"] != "bridge-blocked" ||
		payload["managed_launcher"] != "xnix-compat-launch --app 7zr" ||
		payload["launcher_argv_accepted"] != true ||
		payload["request_id"] != "known-app-launch-request-7zr" ||
		payload["dispatch_id"] != "known-app-dispatch-7zr" ||
		payload["runtime_method"] != "BridgeKnownLauncherToDispatchSmoke" ||
		payload["dispatch_smoke_request_type"] != "windows-known-app-dispatch-smoke" ||
		payload["dispatch_smoke_request_materialized"] != false ||
		payload["app_id"] != "7zr" ||
		payload["dispatch_gate"] != "managed-known-app-guest-smoke" ||
		payload["guest_boundary_required"] != true ||
		payload["guest_boundary_supplied"] != false ||
		payload["smoke_harness_required"] != true ||
		payload["cache_status"] != "missing" ||
		payload["artifact_verified"] != false ||
		payload["bridge_preview_created"] != true ||
		payload["dispatch_ready"] != false ||
		payload["preparation_required"] != true ||
		payload["runtime_owned_bridge"] != true ||
		payload["kde_presentation_only"] != true ||
		payload["dry_run"] != true ||
		payload["dispatch_started"] != false ||
		payload["execution_started"] != false ||
		payload["backend_process_started"] != false ||
		payload["host_root_modified"] != false ||
		payload["host_networking_required"] != false ||
		payload["docker_socket_mounted"] != false ||
		payload["broad_host_mount_required"] != false ||
		payload["raw_host_path_exposed"] != false ||
		payload["raw_executable_path_exposed"] != false ||
		payload["raw_command_exposed"] != false ||
		payload["backend_details_exposed"] != false {
		t.Fatalf("unexpected launch bridge payload: %#v", payload)
	}
	assertCompatLaunchCLISafe(t, output.String(), tempDir)
}

func TestCompatLaunchRequiresApp(t *testing.T) {
	var output bytes.Buffer
	err := run(nil, &output)
	if err == nil || !strings.Contains(err.Error(), "--app is required") {
		t.Fatalf("expected missing app rejection, got %v", err)
	}
}

func TestCompatLaunchRejectsUnknownApp(t *testing.T) {
	var output bytes.Buffer
	err := run([]string{"--app", "missing-app"}, &output)
	if err == nil || !strings.Contains(err.Error(), "unknown known Windows app") {
		t.Fatalf("expected unknown app rejection, got %v", err)
	}
}

func assertCompatLaunchCLISafe(t *testing.T, text string, hostPath string) {
	t.Helper()
	serialized := strings.ToLower(text)
	for _, forbidden := range []string{".exe", "wine", "qemu", strings.ToLower(hostPath)} {
		if strings.Contains(serialized, forbidden) {
			t.Fatalf("compat launch CLI exposed forbidden term %q: %s", forbidden, text)
		}
	}
}
