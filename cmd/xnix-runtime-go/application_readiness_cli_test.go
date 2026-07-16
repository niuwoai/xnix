package main

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"
)

func TestApplicationReadinessPreviewCommand(t *testing.T) {
	registryPath, app := writeTestRepairGroupRegistry(t)
	var output bytes.Buffer
	if err := run([]string{"application-readiness-preview", "--registry", registryPath, "--app", app, "--root", "../.."}, &output); err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["version"] != "0.2.300" ||
		payload["schema_version"] != "xnix.runtime.application_readiness.v1" ||
		payload["request_type"] != "application-readiness-preview" ||
		payload["graph_type"] != "runtime-application-readiness-evidence-graph" ||
		payload["runtime_method"] != "GetApplicationReadiness" ||
		payload["node_count"] != float64(7) ||
		payload["ready"] != false ||
		payload["runtime_owned"] != true ||
		payload["go_runtime_backed"] != true ||
		payload["kde_policy_owner"] != false ||
		payload["launch_enabled"] != false ||
		payload["write_methods_enabled"] != false ||
		payload["execution_started"] != false ||
		payload["backend_launch_enabled"] != false ||
		payload["real_portal_transport_enabled"] != false ||
		payload["request_object_created"] != false ||
		payload["snapshot_created"] != false ||
		payload["host_root_modified"] != false ||
		payload["network_required"] != false ||
		payload["privileged_container_required"] != false ||
		payload["state_root_path_exposed"] != false ||
		payload["backend_details_exposed"] != false {
		t.Fatalf("unexpected readiness payload: %#v", payload)
	}
	application := payload["application"].(map[string]any)
	if application["id"] != app {
		t.Fatalf("unexpected readiness application: %#v", application)
	}
	for _, node := range []string{"recipe-trust", "artifact-stage-receipt", "backend-lifecycle", "portal-review", "snapshot-baseline", "execution-readiness", "runtime-write-gate"} {
		if !strings.Contains(output.String(), node) {
			t.Fatalf("readiness output missing node %q: %s", node, output.String())
		}
	}
	assertApplicationReadinessCLISafe(t, output.String())
}

func TestApplicationReadinessPreviewCommandRejectsAmbiguousSource(t *testing.T) {
	registryPath, app := writeTestRepairGroupRegistry(t)
	var output bytes.Buffer
	err := run([]string{"application-readiness-preview", "--registry", registryPath, "--app", app, "--recipe", registryPath}, &output)
	if err == nil || !strings.Contains(err.Error(), "requires exactly one source") {
		t.Fatalf("expected source rejection, got %v", err)
	}
}

func assertApplicationReadinessCLISafe(t *testing.T, text string) {
	t.Helper()
	serialized := strings.ToLower(text)
	for _, forbidden := range []string{"prefix", ".exe", "program files", "qemu-system", "proton", "wine ", "wine/", ".wine", "virtual machine", "/users/", "/private/", "docker.sock"} {
		if strings.Contains(serialized, forbidden) {
			t.Fatalf("application readiness CLI exposed forbidden term %q: %s", forbidden, text)
		}
	}
}
