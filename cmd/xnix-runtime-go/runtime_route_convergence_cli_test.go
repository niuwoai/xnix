package main

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"
)

func TestRuntimeRouteConvergencePreviewCommandRendersMigrationEvidence(t *testing.T) {
	var output bytes.Buffer
	err := run([]string{"runtime-route-convergence-preview", "--root", projectRootForRuntimeServiceBindingCommandTest(t)}, &output)
	if err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["version"] != "0.2.300" ||
		payload["schema_version"] != "xnix.runtime.route_convergence.v1" ||
		payload["request_type"] != "runtime-route-convergence-preview" ||
		payload["report_type"] != "runtime-route-convergence" ||
		payload["source"] != "runtime-owner-route-manifest-preview+runtime-method-parity-manifest-preview" ||
		payload["runtime_method"] != "GetRuntimeOwnerRouteManifest" ||
		payload["read_method"] != "GetRuntimeRouteConvergencePreview" {
		t.Fatalf("unexpected Runtime route convergence CLI schema: %#v", payload)
	}

	counts := payload["classification_counts"].(map[string]any)
	if counts["total"] != float64(61) ||
		counts["go_product_logic"] != float64(61) ||
		counts["c_policy_bridge"] != float64(0) ||
		counts["ruby_smoke_bridge"] != float64(0) ||
		counts["fixture_only"] != float64(0) ||
		counts["contract_only"] != float64(0) ||
		counts["deprecated"] != float64(0) ||
		counts["unsupported"] != float64(0) ||
		counts["unclassified"] != float64(0) {
		t.Fatalf("unexpected route convergence counts: %#v", counts)
	}
	if payload["all_routes_classified"] != true ||
		payload["native_go_coverage_ready"] != true ||
		payload["production_owner_ready"] != false ||
		payload["runtime_owned"] != true ||
		payload["go_runtime_backed"] != true ||
		payload["kde_policy_owner"] != false ||
		payload["write_methods_enabled"] != false ||
		payload["network_required"] != false ||
		payload["host_root_modified"] != false ||
		payload["privileged_container_required"] != false ||
		payload["backend_launch_enabled"] != false ||
		payload["backend_details_exposed"] != false {
		t.Fatalf("unexpected route convergence safety flags: %#v", payload)
	}

	groups := payload["migration_groups"].([]any)
	if len(groups) == 0 {
		t.Fatalf("expected migration groups: %#v", payload)
	}
	checks := payload["checks"].([]any)
	checkIDs := payload["check_ids"].([]any)
	expectedIDs := []string{
		"owner-route-manifest-ready",
		"all-routes-classified",
		"native-go-route-coverage",
		"write-gate-disabled",
		"host-safety-boundary",
	}
	if len(checks) != len(expectedIDs) || len(checkIDs) != len(expectedIDs) {
		t.Fatalf("unexpected checks: %#v ids=%#v", checks, checkIDs)
	}
	for index, id := range expectedIDs {
		check := checks[index].(map[string]any)
		if check["id"] != id || checkIDs[index] != id || check["status"] != "pass" {
			t.Fatalf("unexpected check at %d: %#v ids=%#v", index, checks, checkIDs)
		}
	}

	serialized := strings.ToLower(output.String())
	for _, forbidden := range []string{"prefix", ".exe", "program files", "qemu-system", "proton", "wine "} {
		if strings.Contains(serialized, forbidden) {
			t.Fatalf("Runtime route convergence CLI exposed forbidden term %q: %s", forbidden, output.String())
		}
	}
}
