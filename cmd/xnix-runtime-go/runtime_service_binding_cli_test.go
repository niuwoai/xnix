package main

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestRuntimeServiceBindingPreviewCommandRendersGoReadModel(t *testing.T) {
	var output bytes.Buffer
	err := run([]string{"runtime-service-binding-preview", "--root", projectRootForRuntimeServiceBindingCommandTest(t)}, &output)
	if err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["version"] != currentProjectVersion(t) ||
		payload["schema_version"] != "xnix.runtime.service_binding.v1" ||
		payload["request_type"] != "runtime-service-binding-preview" ||
		payload["binding_type"] != "runtime-service-binding" ||
		payload["source"] != "activation-files+dbus-contract+runtime-owner-gate" ||
		payload["runtime_method"] != "GetRuntimeServiceBinding" ||
		payload["read_method"] != "GetRuntimeServiceBindingPreview" ||
		payload["production_status"] != "pending-live-owner" {
		t.Fatalf("unexpected Runtime service binding CLI schema: %#v", payload)
	}
	if payload["bus_name"] != "org.xnix.Compatibility1" ||
		payload["object_path"] != "/org/xnix/Compatibility1" ||
		payload["interface"] != "org.xnix.Compatibility1" {
		t.Fatalf("unexpected Runtime D-Bus identity: %#v", payload)
	}
	activation := payload["activation"].(map[string]any)
	if activation["dbus_service_file"] != "runtime/dbus/org.xnix.Compatibility1.service" ||
		activation["systemd_unit"] != "runtime/systemd/xnix-compatd.service" ||
		activation["libexec_wrapper"] != "libexec/xnix/compatd" ||
		activation["dbus_contract"] != "runtime/dbus/org.xnix.Compatibility1.xml" ||
		activation["packaged_wrapper"] != "/usr/libexec/xnix/compatd" {
		t.Fatalf("unexpected activation payload: %#v", activation)
	}
	checks := payload["checks"].([]any)
	checkIDs := payload["check_ids"].([]any)
	expectedIDs := []string{
		"dbus-service-activation",
		"systemd-service-hardening",
		"libexec-wrapper",
		"dbus-contract",
		"live-dbus-owner",
	}
	if len(checks) != len(expectedIDs) || len(checkIDs) != len(expectedIDs) {
		t.Fatalf("unexpected Runtime service binding checks: %#v ids=%#v", checks, checkIDs)
	}
	for index, id := range expectedIDs {
		check := checks[index].(map[string]any)
		if check["id"] != id || checkIDs[index] != id {
			t.Fatalf("unexpected check at %d: %#v ids=%#v", index, checks, checkIDs)
		}
		expectedStatus := "pass"
		if id == "live-dbus-owner" {
			expectedStatus = "pending"
		}
		if check["status"] != expectedStatus {
			t.Fatalf("unexpected status for %s: %#v", id, check)
		}
	}
	counts := payload["counts"].(map[string]any)
	if counts["total"] != float64(5) ||
		counts["passed"] != float64(4) ||
		counts["pending"] != float64(1) ||
		counts["blocked"] != float64(0) {
		t.Fatalf("unexpected Runtime service binding counts: %#v", counts)
	}
	if payload["runtime_owned"] != true ||
		payload["go_runtime_backed"] != true ||
		payload["kde_policy_owner"] != false ||
		payload["activation_binding_ready"] != true ||
		payload["live_dbus_owner_ready"] != false ||
		payload["smoke_adapter_available"] != true ||
		payload["service_started"] != false ||
		payload["production_bus_claimed"] != false ||
		payload["kde_may_claim_runtime_ownership"] != false ||
		payload["network_required"] != false ||
		payload["host_root_modified"] != false ||
		payload["privileged_container_required"] != false ||
		payload["backend_details_exposed"] != false {
		t.Fatalf("unexpected Runtime service binding safety flags: %#v", payload)
	}
	blockedActions := payload["blocked_actions"].([]any)
	if len(blockedActions) != 4 ||
		blockedActions[0] != "start Runtime service from preview" ||
		blockedActions[1] != "claim production D-Bus owner from preview" ||
		blockedActions[2] != "let KDE claim Runtime ownership" ||
		blockedActions[3] != "mutate host root during service binding planning" {
		t.Fatalf("unexpected blocked actions: %#v", blockedActions)
	}
	if strings.Contains(strings.ToLower(output.String()), "prefix") ||
		strings.Contains(strings.ToLower(output.String()), ".exe") ||
		strings.Contains(strings.ToLower(output.String()), "program files") {
		t.Fatalf("Runtime service binding CLI exposed backend terms: %s", output.String())
	}
}

func TestRuntimeServiceActivationPreflightPreviewCommandRendersGoReadModel(t *testing.T) {
	var output bytes.Buffer
	err := run([]string{"runtime-service-activation-preflight-preview", "--root", projectRootForRuntimeServiceBindingCommandTest(t)}, &output)
	if err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["version"] != currentProjectVersion(t) ||
		payload["schema_version"] != "xnix.runtime.service_activation_preflight.v1" ||
		payload["request_type"] != "runtime-service-activation-preflight-preview" ||
		payload["preflight_type"] != "production-runtime-service-activation-preflight" ||
		payload["source"] != "runtime-service-binding-preview+runtime-owner-readiness-preview+runtime-owner-smoke-plan-preview+production-dbus-gate-review-preview+production-dbus-human-authorization-preflight-preview+production-human-authorization-receipt-consolidation-preview" ||
		payload["runtime_method"] != "GetRuntimeServiceActivationPreflight" ||
		payload["read_method"] != "GetRuntimeServiceActivationPreflightPreview" {
		t.Fatalf("unexpected Runtime service activation preflight CLI schema: %#v", payload)
	}
	if payload["bus_name"] != "org.xnix.Compatibility1" ||
		payload["object_path"] != "/org/xnix/Compatibility1" ||
		payload["interface"] != "org.xnix.Compatibility1" {
		t.Fatalf("unexpected Runtime service activation D-Bus identity: %#v", payload)
	}

	binding := payload["service_binding"].(map[string]any)
	if binding["request_type"] != "runtime-service-binding-preview" ||
		binding["production_status"] != "pending-live-owner" ||
		binding["activation_binding_ready"] != true ||
		binding["live_dbus_owner_ready"] != false ||
		binding["smoke_adapter_available"] != true {
		t.Fatalf("unexpected service binding summary: %#v", binding)
	}
	readiness := payload["owner_readiness"].(map[string]any)
	if readiness["request_type"] != "runtime-owner-readiness-preview" ||
		readiness["readiness_type"] != "runtime-owner-readiness" ||
		readiness["activation_binding_ready"] != true ||
		readiness["read_only_method_parity_ready"] != true ||
		readiness["owner_smoke_planned"] != true ||
		readiness["live_dbus_owner_ready"] != false ||
		readiness["production_owner_enabled"] != false ||
		readiness["owner_transition_ready"] != false ||
		readiness["production_recipe_trust_ready"] != false ||
		readiness["system_service_started"] != false ||
		readiness["production_bus_claimed"] != false ||
		readiness["write_methods_enabled"] != false {
		t.Fatalf("unexpected owner readiness summary: %#v", readiness)
	}
	smoke := payload["owner_smoke_plan"].(map[string]any)
	if smoke["request_type"] != "runtime-owner-smoke-plan-preview" ||
		smoke["plan_type"] != "runtime-owner-smoke-plan" ||
		smoke["smoke_state"] != "planned" ||
		smoke["smoke_environment"] != "restricted-session" ||
		smoke["pending_step_count"] != float64(6) {
		t.Fatalf("unexpected owner smoke plan summary: %#v", smoke)
	}
	productionDBusGate := payload["production_dbus_gate"].(map[string]any)
	if productionDBusGate["request_type"] != "production-dbus-gate-review-preview" ||
		productionDBusGate["gate_type"] != "owner-local-smoke-covered-production-dbus-gate-review" ||
		productionDBusGate["gate_decision"] != "production-dbus-gate-review-consumed-activation-still-blocked" ||
		productionDBusGate["gate_review_source_present"] != true ||
		productionDBusGate["human_authorization_preflight_present"] != true ||
		productionDBusGate["gate_review_ready"] != true ||
		productionDBusGate["human_authorization_preflight_ready"] != true ||
		productionDBusGate["human_authorization_required"] != true ||
		productionDBusGate["human_authorization_granted"] != false ||
		productionDBusGate["authorization_receipt_accepted"] != false ||
		productionDBusGate["production_readiness"] != false ||
		productionDBusGate["production_owner_enabled"] != false ||
		productionDBusGate["production_activation_ready"] != false ||
		productionDBusGate["system_service_started"] != false ||
		productionDBusGate["session_bus_claimed"] != false ||
		productionDBusGate["production_bus_claimed"] != false ||
		productionDBusGate["write_methods_enabled"] != false ||
		productionDBusGate["runtime_writes_enabled"] != false ||
		productionDBusGate["backend_launch_enabled"] != false ||
		productionDBusGate["host_root_modified"] != false {
		t.Fatalf("unexpected production D-Bus gate summary: %#v", productionDBusGate)
	}

	checks := payload["preflight_checks"].([]any)
	checkIDs := payload["check_ids"].([]any)
	expectedIDs := []string{
		"activation-binding",
		"read-only-method-parity",
		"owner-smoke-plan",
		"production-dbus-gate-review",
		"human-authorization-preflight",
		"human-authorization-receipt",
		"write-method-gate",
		"kde-ownership-boundary",
		"host-safety-boundary",
		"long-running-runtime-owner",
		"restricted-owner-smoke",
		"production-bus-claim",
		"production-recipe-trust",
	}
	expectedStatuses := []string{"pass", "pass", "pass", "pass", "pass", "pending", "pass", "pass", "pass", "pending", "pending", "pending", "pending"}
	if len(checks) != len(expectedIDs) || len(checkIDs) != len(expectedIDs) {
		t.Fatalf("unexpected Runtime service activation preflight checks: %#v ids=%#v", checks, checkIDs)
	}
	for index, id := range expectedIDs {
		check := checks[index].(map[string]any)
		if check["id"] != id || checkIDs[index] != id || check["status"] != expectedStatuses[index] {
			t.Fatalf("unexpected preflight check at %d: %#v ids=%#v", index, checks, checkIDs)
		}
	}
	counts := payload["counts"].(map[string]any)
	if counts["total"] != float64(13) ||
		counts["passed"] != float64(8) ||
		counts["pending"] != float64(5) ||
		counts["blocked"] != float64(0) {
		t.Fatalf("unexpected Runtime service activation preflight counts: %#v", counts)
	}
	if payload["preflight_decision"] != "restricted-owner-smoke-ready" ||
		payload["production_activation_ready"] != false ||
		payload["restricted_smoke_ready"] != true ||
		payload["production_dbus_gate_ready"] != true ||
		payload["human_authorization_preflight_ready"] != true ||
		payload["human_authorization_required"] != true ||
		payload["human_authorization_granted"] != false ||
		payload["authorization_receipt_accepted"] != false ||
		payload["runtime_owned"] != true ||
		payload["go_runtime_backed"] != true ||
		payload["kde_policy_owner"] != false ||
		payload["kde_may_claim_runtime_ownership"] != false ||
		payload["system_service_started"] != false ||
		payload["production_bus_claimed"] != false ||
		payload["write_methods_enabled"] != false ||
		payload["backend_launch_enabled"] != false ||
		payload["network_required"] != false ||
		payload["host_root_modified"] != false ||
		payload["privileged_container_required"] != false ||
		payload["backend_details_exposed"] != false {
		t.Fatalf("unexpected Runtime service activation preflight safety flags: %#v", payload)
	}
	if strings.Contains(strings.ToLower(output.String()), "prefix") ||
		strings.Contains(strings.ToLower(output.String()), ".exe") ||
		strings.Contains(strings.ToLower(output.String()), "program files") {
		t.Fatalf("Runtime service activation preflight CLI exposed backend terms: %s", output.String())
	}
}

func projectRootForRuntimeServiceBindingCommandTest(t *testing.T) string {
	t.Helper()
	workingDirectory, err := os.Getwd()
	if err != nil {
		t.Fatalf("Getwd returned error: %v", err)
	}
	for {
		if _, err := os.Stat(filepath.Join(workingDirectory, "VERSION")); err == nil {
			return workingDirectory
		}
		parent := filepath.Dir(workingDirectory)
		if parent == workingDirectory || strings.TrimSpace(parent) == "" {
			t.Fatalf("could not find project root from %s", workingDirectory)
		}
		workingDirectory = parent
	}
}
