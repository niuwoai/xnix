package main

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"
)

func TestLocalGoCompilePolicyPreviewCommand(t *testing.T) {
	root := projectRootForRuntimeServiceBindingCommandTest(t)
	var output bytes.Buffer
	if err := run([]string{"local-go-compile-policy-preview", "--root", root}, &output); err != nil {
		t.Fatalf("local-go-compile-policy-preview returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["version"] != currentProjectVersion(t) ||
		payload["schema_version"] != "xnix.runtime.local_go_compile_policy.v1" ||
		payload["request_type"] != "local-go-compile-policy-preview" ||
		payload["policy_type"] != "q4-first-go-compilation-policy" ||
		payload["policy_decision"] != "local-go-compilation-policy-ready-q4-first" {
		t.Fatalf("unexpected local Go compile policy payload: %s", output.String())
	}
	if payload["default_remote_host"] != "root@q4" ||
		payload["remote_build_command"] != "ruby scripts/remote_go_build.rb --execute" ||
		payload["remote_test_command"] != "ruby scripts/remote_go_test.rb --execute" ||
		payload["local_override_environment"] != "XNIX_ALLOW_LOCAL_GO_COMPILE=1" {
		t.Fatalf("unexpected local Go compile policy commands: %s", output.String())
	}
	for _, key := range []string{"q4_compilation_required_by_default", "remote_go_build_runner_present", "remote_go_test_runner_present", "remote_go_caches_pinned_to_q4", "targeted_small_version_testing", "twentieth_version_full_gate_required", "protected_claude_package_excluded", "host_compilation_avoided_observable"} {
		if payload[key] != true {
			t.Fatalf("local Go compile policy %s must be true: %s", key, output.String())
		}
	}
	for _, key := range []string{"local_go_compilation_allowed_by_default", "local_go_run_allowed_by_default", "q4_host_root_modified", "host_root_modified", "privileged_container_required", "host_networking_required", "docker_socket_mounted", "broad_host_mount_required"} {
		if payload[key] != false {
			t.Fatalf("local Go compile policy %s must be false: %s", key, output.String())
		}
	}
	counts := payload["counts"].(map[string]any)
	if counts["total"] != float64(8) || counts["passed"] != float64(8) || counts["blocked"] != float64(0) {
		t.Fatalf("unexpected local Go compile policy checks: %s", output.String())
	}
	if strings.Contains(output.String(), root) {
		t.Fatalf("local Go compile policy output must not expose project root path: %s", output.String())
	}
}

func TestLocalGoCompilePolicyPreviewCommandRejectsPositionalArgs(t *testing.T) {
	var output bytes.Buffer
	if err := run([]string{"local-go-compile-policy-preview", "extra"}, &output); err == nil {
		t.Fatalf("local-go-compile-policy-preview must reject positional arguments")
	}
}
