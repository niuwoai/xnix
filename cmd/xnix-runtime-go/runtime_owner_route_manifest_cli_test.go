package main

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"
)

func TestRuntimeOwnerRouteManifestPreviewCommandRendersGoReadModel(t *testing.T) {
	var output bytes.Buffer
	err := run([]string{"runtime-owner-route-manifest-preview", "--root", projectRootForRuntimeServiceBindingCommandTest(t)}, &output)
	if err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["version"] != "0.2.246" ||
		payload["schema_version"] != "xnix.runtime.owner_route_manifest.v1" ||
		payload["request_type"] != "runtime-owner-route-manifest-preview" ||
		payload["manifest_type"] != "runtime-owner-route-manifest" ||
		payload["source"] != "runtime-method-parity-manifest-preview+go-runtime-cli+c-runtime-core+runtime-dispatch" ||
		payload["runtime_method"] != "GetRuntimeOwnerRouteManifest" ||
		payload["read_method"] != "GetRuntimeOwnerRouteManifestPreview" {
		t.Fatalf("unexpected Runtime owner route manifest CLI schema: %#v", payload)
	}
	if payload["bus_name"] != "org.xnix.Compatibility1" ||
		payload["object_path"] != "/org/xnix/Compatibility1" ||
		payload["interface"] != "org.xnix.Compatibility1" {
		t.Fatalf("unexpected Runtime D-Bus identity: %#v", payload)
	}

	routeCounts := payload["route_counts"].(map[string]any)
	if routeCounts["total"] != float64(61) ||
		routeCounts["go_routed"] != float64(61) ||
		routeCounts["c_core_backed"] != float64(0) ||
		routeCounts["ruby_legacy"] != float64(0) ||
		routeCounts["ready"] != float64(61) ||
		routeCounts["pending"] != float64(0) ||
		routeCounts["blocked"] != float64(0) {
		t.Fatalf("unexpected route counts: %#v", routeCounts)
	}
	checks := payload["checks"].([]any)
	checkIDs := payload["check_ids"].([]any)
	expectedIDs := []string{"method-parity", "go-route-coverage", "c-core-adapter-boundary", "ruby-legacy-dispatch", "write-route-gate", "host-safety-boundary"}
	expectedStatuses := []string{"pass", "pass", "pass", "pass", "pass", "pass"}
	if len(checks) != len(expectedIDs) || len(checkIDs) != len(expectedIDs) {
		t.Fatalf("unexpected route checks: %#v ids=%#v", checks, checkIDs)
	}
	for index, id := range expectedIDs {
		check := checks[index].(map[string]any)
		if check["id"] != id || checkIDs[index] != id || check["status"] != expectedStatuses[index] {
			t.Fatalf("unexpected route check at %d: %#v ids=%#v", index, checks, checkIDs)
		}
	}
	if payload["method_parity_ready"] != true ||
		payload["go_owner_route_coverage_ready"] != true ||
		payload["c_core_adapter_required"] != false ||
		payload["legacy_runtime_routes_present"] != false ||
		payload["production_owner_routes_ready"] != false ||
		payload["runtime_owned"] != true ||
		payload["go_runtime_backed"] != true ||
		payload["kde_policy_owner"] != false ||
		payload["kde_may_claim_runtime_ownership"] != false ||
		payload["system_service_started"] != false ||
		payload["production_bus_claimed"] != false ||
		payload["write_methods_enabled"] != false ||
		payload["network_required"] != false ||
		payload["host_root_modified"] != false ||
		payload["privileged_container_required"] != false ||
		payload["backend_details_exposed"] != false {
		t.Fatalf("unexpected route manifest safety flags: %#v", payload)
	}

	serialized := strings.ToLower(output.String())
	for _, forbidden := range []string{"prefix", ".exe", "program files", "qemu-system", "proton", "wine "} {
		if strings.Contains(serialized, forbidden) {
			t.Fatalf("Runtime owner route manifest CLI exposed forbidden term %q: %s", forbidden, output.String())
		}
	}
}
