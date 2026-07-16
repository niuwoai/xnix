package main

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"
)

func TestKDERestrictedLaunchAuthorizationRecordCommandRequiresExplicitDirective(t *testing.T) {
	stateRoot := t.TempDir()
	var output bytes.Buffer
	err := run([]string{
		"kde-restricted-launch-authorization-record",
		"--registry", "../../runtime/recipes/registry.json",
		"--app", "org.xnix.sample.notepad",
		"--state-root", stateRoot,
		"--mode", "test-only",
		"--authorize", "authorize-restricted-test-preparation",
	}, &output)
	if err != nil {
		t.Fatalf("kde-restricted-launch-authorization-record returned error: %v", err)
	}
	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("decode command output: %v", err)
	}
	authorization := payload["authorization"].(map[string]any)
	execution := payload["execution"].(map[string]any)
	if authorization["state"] != "authorized-preparation-only" || authorization["preparation_authorized"] != true || authorization["receipt_read_back"] != true || execution["state"] != "blocked" || payload["authorization_boundary_joined"] != true || payload["core_receipt_count"] != float64(9) || payload["all_checks_passed"] != true {
		t.Fatalf("unexpected restricted authorization command payload: %s", output.String())
	}
	if strings.Contains(output.String(), stateRoot) {
		t.Fatalf("command output must not expose state-root path: %s", output.String())
	}
	for _, key := range []string{"production_trust_satisfied", "runtime_write_gate_enabled", "artifact_acquisition_enabled", "backend_install_enabled", "backend_launch_enabled", "backend_process_started", "real_portal_call_enabled", "execution_approved", "launch_authorized", "launch_allowed", "launch_enabled", "execution_started", "process_start_authorized", "production_bus_ownership", "network_required", "host_root_modified", "privileged_container_required", "raw_command_exposed", "backend_details_exposed"} {
		if payload[key] != false {
			t.Fatalf("unsafe gate %s must remain false: %s", key, output.String())
		}
	}
}

func TestKDERestrictedLaunchAuthorizationRecordCommandRejectsImplicitAuthorization(t *testing.T) {
	base := []string{"kde-restricted-launch-authorization-record", "--registry", "../../runtime/recipes/registry.json", "--app", "org.xnix.sample.notepad", "--state-root", t.TempDir()}
	for _, args := range [][]string{
		base,
		append(append([]string{}, base...), "--mode", "test-only"),
		append(append([]string{}, base...), "--mode", "production", "--authorize", "authorize-restricted-test-preparation"),
		append(append([]string{}, base...), "--mode", "test-only", "--authorize", "yes"),
		append(append([]string{}, base...), "--mode", "test-only", "--authorize", "authorize-restricted-test-preparation", "extra"),
	} {
		if err := run(args, &bytes.Buffer{}); err == nil {
			t.Fatalf("expected command to reject implicit or unsafe authorization: %v", args)
		}
	}
}
