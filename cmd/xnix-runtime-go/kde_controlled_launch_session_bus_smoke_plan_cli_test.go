package main

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"
)

func TestKDEControlledLaunchSessionBusSmokePlanPreviewCommandLinksActionToRestrictedSmoke(t *testing.T) {
	stateRoot := t.TempDir()
	record := recordKDEControlledLaunchActionCLIEvidence(t, stateRoot)

	var output bytes.Buffer
	err := run([]string{
		"kde-controlled-launch-session-bus-smoke-plan-preview",
		"--state-root", stateRoot,
		"--evidence-relative-path", record.EvidenceRelativePath,
	}, &output)
	if err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("KDE controlled launch session-bus smoke plan output must be JSON: %v\n%s", err, output.String())
	}
	if payload["schema_version"] != "xnix.runtime.kde_controlled_launch_session_bus_smoke_plan.v1" ||
		payload["request_type"] != "kde-controlled-launch-session-bus-smoke-plan-preview" ||
		payload["kde_action_id"] != "xnix.runtime-status.controlled-launch" ||
		payload["kde_action_preview_request_type"] != "kde-controlled-launch-action-preview" ||
		payload["public_dbus_method"] != "org.xnix.Compatibility1.ShowRuntimeControlledLaunch" ||
		payload["restricted_session_bus_plan_ready"] != true ||
		payload["dbus_controlled_launch_fixture_plan_ready"] != true ||
		payload["private_session_bus_required"] != true ||
		payload["dbus_session_bus_address_required"] != true ||
		payload["kde_forwards_only_evidence_handle"] != true ||
		payload["runtime_owned"] != true ||
		payload["go_runtime_backed"] != true ||
		payload["owner_service_args_exposed_to_kde"] != false ||
		payload["desktop_kde_state_root_access"] != false ||
		payload["desktop_receipt_fields_reconstructed"] != false ||
		payload["state_root_path_exposed"] != false ||
		payload["raw_launcher_output_exposed"] != false ||
		payload["host_root_modified"] != false ||
		payload["docker_socket_mounted"] != false ||
		payload["privileged_container_required"] != false ||
		payload["execution_started"] != false ||
		payload["smoke_executed_by_preview"] != false {
		t.Fatalf("unexpected KDE controlled launch session-bus smoke plan payload: %#v", payload)
	}

	outerCommand, ok := payload["outer_private_session_bus_command"].([]any)
	if !ok || len(outerCommand) != 4 ||
		outerCommand[0] != "dbus-run-session" ||
		outerCommand[1] != "--" ||
		outerCommand[2] != "ruby" ||
		outerCommand[3] != "scripts/runtime_status_owner_service_session_bus_smoke.rb" {
		t.Fatalf("unexpected private session bus command: %#v", payload)
	}
	containerCommand, ok := payload["container_smoke_command"].([]any)
	if !ok || len(containerCommand) != 3 ||
		containerCommand[0] != "ruby" ||
		containerCommand[1] != "scripts/container.rb" ||
		containerCommand[2] != "runtime-status-owner-service-session-bus-smoke" {
		t.Fatalf("unexpected container smoke command: %#v", payload)
	}
	dbusFixtureCommand, ok := payload["dbus_controlled_launch_fixture_command"].([]any)
	if !ok || len(dbusFixtureCommand) != 2 ||
		dbusFixtureCommand[0] != "ruby" ||
		dbusFixtureCommand[1] != "scripts/dbus_controlled_launch_owner_fixture_smoke.rb" {
		t.Fatalf("unexpected D-Bus fixture command: %#v", payload)
	}
	dbusFixtureContainerCommand, ok := payload["dbus_controlled_launch_fixture_container_command"].([]any)
	if !ok || len(dbusFixtureContainerCommand) != 3 ||
		dbusFixtureContainerCommand[0] != "ruby" ||
		dbusFixtureContainerCommand[1] != "scripts/container.rb" ||
		dbusFixtureContainerCommand[2] != "dbus-controlled-launch-owner-fixture-smoke" {
		t.Fatalf("unexpected D-Bus fixture container command: %#v", payload)
	}
	forwardedArgs, ok := payload["kde_forwarded_arguments"].([]any)
	if !ok || len(forwardedArgs) != 1 || forwardedArgs[0] != record.EvidenceRelativePath {
		t.Fatalf("KDE smoke plan must forward only the evidence handle: %#v", payload)
	}
	if strings.Contains(output.String(), "owner_service_call_args") ||
		strings.Contains(output.String(), stateRoot) {
		t.Fatalf("KDE smoke plan output exposed Runtime-owned details: %s", output.String())
	}
}

func TestKDEControlledLaunchSessionBusSmokePlanPreviewCommandRejectsMissingStateRoot(t *testing.T) {
	var output bytes.Buffer
	err := run([]string{"kde-controlled-launch-session-bus-smoke-plan-preview"}, &output)
	if err == nil || !strings.Contains(err.Error(), "requires --state-root") {
		t.Fatalf("missing state root must be rejected, got: %v", err)
	}
}
