package main

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"
)

func TestRuntimeOwnerReadinessPreviewCommandRendersGoReadModel(t *testing.T) {
	var output bytes.Buffer
	err := run([]string{"runtime-owner-readiness-preview", "--root", projectRootForRuntimeServiceBindingCommandTest(t)}, &output)
	if err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["version"] != "0.2.232" ||
		payload["schema_version"] != "xnix.runtime.owner_readiness.v1" ||
		payload["request_type"] != "runtime-owner-readiness-preview" ||
		payload["readiness_type"] != "runtime-owner-readiness" ||
		payload["source"] != "runtime-service-binding-preview+runtime-live-owner-gate-preview+runtime-owner-process-preview+runtime-owner-smoke-plan-preview+runtime-method-parity-manifest-preview+runtime-owner-route-manifest-preview+runtime-owner-recipe-trust-preview" ||
		payload["runtime_method"] != "GetRuntimeOwnerReadiness" ||
		payload["read_method"] != "GetRuntimeOwnerReadinessPreview" {
		t.Fatalf("unexpected Runtime owner readiness CLI schema: %#v", payload)
	}
	if payload["bus_name"] != "org.xnix.Compatibility1" ||
		payload["object_path"] != "/org/xnix/Compatibility1" ||
		payload["interface"] != "org.xnix.Compatibility1" {
		t.Fatalf("unexpected Runtime D-Bus identity: %#v", payload)
	}
	serviceBinding := payload["service_binding"].(map[string]any)
	if serviceBinding["request_type"] != "runtime-service-binding-preview" ||
		serviceBinding["production_status"] != "pending-live-owner" ||
		serviceBinding["activation_binding_ready"] != true ||
		serviceBinding["live_dbus_owner_ready"] != false ||
		serviceBinding["smoke_adapter_available"] != true {
		t.Fatalf("unexpected service binding summary: %#v", serviceBinding)
	}
	liveOwnerGate := payload["live_owner_gate"].(map[string]any)
	if liveOwnerGate["request_type"] != "runtime-live-owner-gate-preview" ||
		liveOwnerGate["gate_type"] != "runtime-live-owner-gate" ||
		liveOwnerGate["activation_binding_ready"] != true ||
		liveOwnerGate["live_dbus_owner_ready"] != false ||
		liveOwnerGate["production_owner_enabled"] != false ||
		liveOwnerGate["owner_transition_ready"] != false ||
		liveOwnerGate["smoke_adapter_is_production_owner"] != false {
		t.Fatalf("unexpected live owner gate summary: %#v", liveOwnerGate)
	}
	ownerProcess := payload["owner_process"].(map[string]any)
	if ownerProcess["request_type"] != "runtime-owner-process-preview" ||
		ownerProcess["process_type"] != "runtime-owner-process" ||
		ownerProcess["current_owner_language"] != "ruby-wrapper" ||
		ownerProcess["target_owner_language"] != "go" ||
		ownerProcess["service_activation_ready"] != true ||
		ownerProcess["packaged_entrypoint_ready"] != true ||
		ownerProcess["go_owner_process_ready"] != true ||
		ownerProcess["production_owner_process_ready"] != false {
		t.Fatalf("unexpected owner process summary: %#v", ownerProcess)
	}
	ownerSmokePlan := payload["owner_smoke_plan"].(map[string]any)
	if ownerSmokePlan["request_type"] != "runtime-owner-smoke-plan-preview" ||
		ownerSmokePlan["plan_type"] != "runtime-owner-smoke-plan" ||
		ownerSmokePlan["smoke_state"] != "planned" ||
		ownerSmokePlan["smoke_environment"] != "restricted-session" ||
		ownerSmokePlan["activation_binding_ready"] != true ||
		ownerSmokePlan["pending_step_count"] != float64(6) {
		t.Fatalf("unexpected owner smoke plan summary: %#v", ownerSmokePlan)
	}
	methodParity := payload["method_parity_manifest"].(map[string]any)
	if methodParity["request_type"] != "runtime-method-parity-manifest-preview" ||
		methodParity["manifest_type"] != "runtime-method-parity-manifest" ||
		methodParity["read_only_method_parity_ready"] != true ||
		methodParity["method_count"] != float64(57) ||
		methodParity["write_methods_supported"] != false ||
		methodParity["write_method_dispatch_enabled"] != false {
		t.Fatalf("unexpected method parity summary: %#v", methodParity)
	}
	ownerRouteManifest := payload["owner_route_manifest"].(map[string]any)
	if ownerRouteManifest["request_type"] != "runtime-owner-route-manifest-preview" ||
		ownerRouteManifest["manifest_type"] != "runtime-owner-route-manifest" ||
		ownerRouteManifest["route_count"] != float64(57) ||
		ownerRouteManifest["go_route_count"] != float64(57) ||
		ownerRouteManifest["c_core_route_count"] != float64(0) ||
		ownerRouteManifest["ruby_legacy_route_count"] != float64(0) ||
		ownerRouteManifest["go_owner_route_coverage_ready"] != true ||
		ownerRouteManifest["c_core_adapter_required"] != false ||
		ownerRouteManifest["legacy_runtime_routes_present"] != false ||
		ownerRouteManifest["production_owner_routes_ready"] != false {
		t.Fatalf("unexpected owner route manifest summary: %#v", ownerRouteManifest)
	}
	recipeTrust := payload["recipe_trust"].(map[string]any)
	if recipeTrust["request_type"] != "runtime-owner-recipe-trust-preview" ||
		recipeTrust["trust_type"] != "runtime-owner-recipe-trust" ||
		recipeTrust["registry_name"] != "xnix-local-development" ||
		recipeTrust["recipe_count"] != float64(1) ||
		recipeTrust["digest_verified"] != true ||
		recipeTrust["signed_recipe_validation"] != false ||
		recipeTrust["development_registry"] != true ||
		recipeTrust["unsigned_recipes_present"] != false ||
		recipeTrust["production_recipe_trust_ready"] != false {
		t.Fatalf("unexpected recipe trust summary: %#v", recipeTrust)
	}
	checks := payload["readiness_checks"].([]any)
	checkIDs := payload["check_ids"].([]any)
	expectedIDs := []string{
		"activation-binding",
		"read-only-method-parity",
		"owner-smoke-plan",
		"write-method-gate",
		"kde-ownership-boundary",
		"host-safety-boundary",
		"long-running-runtime-owner",
		"read-only-owner-routes",
		"production-bus-claim",
		"production-recipe-trust",
	}
	expectedStatuses := []string{"pass", "pass", "pass", "pass", "pass", "pass", "pending", "pending", "pending", "pending"}
	if len(checks) != len(expectedIDs) || len(checkIDs) != len(expectedIDs) {
		t.Fatalf("unexpected readiness checks: %#v ids=%#v", checks, checkIDs)
	}
	for index, id := range expectedIDs {
		check := checks[index].(map[string]any)
		if check["id"] != id || checkIDs[index] != id || check["status"] != expectedStatuses[index] {
			t.Fatalf("unexpected check at %d: %#v ids=%#v", index, checks, checkIDs)
		}
	}
	counts := payload["counts"].(map[string]any)
	if counts["total"] != float64(10) ||
		counts["passed"] != float64(6) ||
		counts["pending"] != float64(4) ||
		counts["blocked"] != float64(0) {
		t.Fatalf("unexpected readiness counts: %#v", counts)
	}
	if payload["runtime_owned"] != true ||
		payload["go_runtime_backed"] != true ||
		payload["kde_policy_owner"] != false ||
		payload["kde_may_claim_runtime_ownership"] != false ||
		payload["activation_binding_ready"] != true ||
		payload["read_only_method_parity_ready"] != true ||
		payload["owner_smoke_planned"] != true ||
		payload["live_dbus_owner_ready"] != false ||
		payload["production_owner_enabled"] != false ||
		payload["owner_transition_ready"] != false ||
		payload["production_recipe_trust_ready"] != false ||
		payload["production_owner_routes_ready"] != false ||
		payload["network_required"] != false ||
		payload["host_root_modified"] != false ||
		payload["privileged_container_required"] != false ||
		payload["system_service_started"] != false ||
		payload["production_bus_claimed"] != false ||
		payload["write_methods_enabled"] != false ||
		payload["backend_details_exposed"] != false {
		t.Fatalf("unexpected owner readiness safety flags: %#v", payload)
	}
	blockedActions := payload["blocked_actions"].([]any)
	if len(blockedActions) != 7 ||
		blockedActions[0] != "start production Runtime owner from readiness preview" ||
		blockedActions[1] != "claim production D-Bus name from readiness preview" {
		t.Fatalf("unexpected readiness blocked actions: %#v", blockedActions)
	}
	if strings.Contains(strings.ToLower(output.String()), "prefix") ||
		strings.Contains(strings.ToLower(output.String()), ".exe") ||
		strings.Contains(strings.ToLower(output.String()), "program files") {
		t.Fatalf("Runtime owner readiness CLI exposed backend terms: %s", output.String())
	}
}
