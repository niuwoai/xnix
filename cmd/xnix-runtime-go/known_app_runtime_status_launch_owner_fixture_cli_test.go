package main

import (
	"bytes"
	"encoding/json"
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
