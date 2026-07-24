package main

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestKnownAppRuntimeStatusLaunchOwnerFixtureRecordCommandBlocksWithoutArtifact(t *testing.T) {
	stateRoot := t.TempDir()
	cacheRoot := t.TempDir()
	var output bytes.Buffer
	err := run([]string{
		"known-app-runtime-status-launch-owner-fixture-record",
		"--app", "7zr",
		"--state-root", stateRoot,
		"--cache-root", cacheRoot,
	}, &output)
	if err != nil {
		t.Fatalf("run returned error: %v", err)
	}
	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("fixture output must be JSON: %v\n%s", err, output.String())
	}
	if payload["schema_version"] != "xnix.runtime.known_app_runtime_status_launch_owner_fixture.v1" ||
		payload["request_type"] != "known-app-runtime-status-launch-owner-fixture-record" ||
		payload["fixture_state"] != "blocked" ||
		payload["fixture_ready"] != false ||
		payload["launch_authorization_receipt_id"] != "known-app-launch-authorization-7zr-26.02" ||
		payload["launch_authorization_receipt_recorded"] != true ||
		payload["desktop_trigger_ready"] != false ||
		payload["owner_service_call_ready"] != false ||
		payload["desktop_evidence_handle_forwarded"] != false ||
		payload["runtime_owner_service_supplies_inputs"] != false ||
		payload["kde_forwards_only_evidence_handle"] != true ||
		payload["desktop_receipt_fields_reconstructed"] != false ||
		payload["desktop_kde_state_root_access"] != false {
		t.Fatalf("unexpected blocked fixture payload: %#v", payload)
	}
	for _, key := range []string{
		"state_root_path_exposed",
		"evidence_path_exposed",
		"managed_launcher_path_exposed",
		"raw_launcher_output_exposed",
		"backend_details_exposed",
		"host_root_modified",
		"docker_socket_mounted",
		"broad_host_mount_required",
		"desktop_launch_enabled",
		"backend_launch_enabled",
		"execution_started",
		"backend_process_started",
	} {
		if payload[key] != false {
			t.Fatalf("blocked fixture must keep %s=false: %#v", key, payload)
		}
	}
	if strings.Contains(output.String(), stateRoot) || strings.Contains(output.String(), cacheRoot) {
		t.Fatalf("fixture output exposed local paths: %s", output.String())
	}
}

func TestKnownAppRuntimeStatusLaunchOwnerFixtureRecordCommandRejectsMissingStateRoot(t *testing.T) {
	var output bytes.Buffer
	err := run([]string{"known-app-runtime-status-launch-owner-fixture-record"}, &output)
	if err == nil || !strings.Contains(err.Error(), "requires --state-root") {
		t.Fatalf("missing state root must be rejected, got: %v", err)
	}
}

func TestKnownAppRuntimeStatusLaunchOwnerFixtureRecordCommandConsumesGUISmokeEvidence(t *testing.T) {
	tempDir := t.TempDir()
	version := currentProjectVersion(t)
	reportPath := filepath.Join(tempDir, "mines-gui-smoke.json")
	if err := os.WriteFile(reportPath, []byte(guiSmokeEvidenceCLIFixture()), 0o600); err != nil {
		t.Fatalf("WriteFile GUI smoke report returned error: %v", err)
	}

	var evidence bytes.Buffer
	err := run([]string{
		"gui-smoke-evidence-preview",
		"--gui-smoke-report", reportPath,
		"--app-id", "org.xnix.apps.mines",
		"--display-name", "Mines",
		"--app-version", version,
	}, &evidence)
	if err != nil {
		t.Fatalf("gui-smoke-evidence-preview returned error: %v", err)
	}
	evidencePath := filepath.Join(tempDir, "mines-gui-evidence.json")
	if err := os.WriteFile(evidencePath, evidence.Bytes(), 0o600); err != nil {
		t.Fatalf("WriteFile GUI smoke evidence returned error: %v", err)
	}

	stateRoot := filepath.Join(tempDir, "state")
	if err := os.Mkdir(stateRoot, 0o700); err != nil {
		t.Fatalf("Mkdir state root returned error: %v", err)
	}
	var output bytes.Buffer
	err = run([]string{
		"known-app-runtime-status-launch-owner-fixture-record",
		"--state-root", stateRoot,
		"--cache-root", filepath.Join(tempDir, "cache"),
		"--gui-smoke-evidence-file", evidencePath,
	}, &output)
	if err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("fixture output must be JSON: %v\n%s", err, output.String())
	}
	if payload["fixture_ready"] != true ||
		payload["fixture_state"] != "ready" ||
		payload["app_id"] != "org.xnix.apps.mines" ||
		payload["display_name"] != "Mines" ||
		payload["desktop_trigger_ready"] != true ||
		payload["owner_service_call_ready"] != true ||
		payload["desktop_callable_runtime_method"] != "ShowRuntimeControlledLaunch" ||
		payload["desktop_dbus_method"] != "org.xnix.Compatibility1.ShowRuntimeControlledLaunch" {
		t.Fatalf("unexpected GUI-backed fixture payload: %#v", payload)
	}
	if strings.Contains(output.String(), stateRoot) ||
		strings.Contains(output.String(), evidencePath) ||
		strings.Contains(output.String(), filepath.Join(tempDir, "cache")) {
		t.Fatalf("GUI-backed fixture output exposed local paths: %s", output.String())
	}
}
