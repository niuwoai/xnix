package main

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"
)

func TestProductionDBusGateReviewPreviewCommand(t *testing.T) {
	root := projectRootForRuntimeServiceBindingCommandTest(t)
	var output bytes.Buffer
	if err := run([]string{"production-dbus-gate-review-preview", "--root", root}, &output); err != nil {
		t.Fatalf("production-dbus-gate-review-preview returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["version"] != currentProjectVersion(t) ||
		payload["schema_version"] != "xnix.runtime.production_dbus_gate_review.v1" ||
		payload["request_type"] != "production-dbus-gate-review-preview" ||
		payload["gate_type"] != "owner-local-smoke-covered-production-dbus-gate-review" ||
		payload["gate_decision"] != "production-dbus-gate-review-ready" ||
		payload["current_gate_status"] != "owner-local-smoke-covered-production-dbus-disabled" {
		t.Fatalf("unexpected production D-Bus gate review command payload: %s", output.String())
	}
	if payload["route_count"] != float64(3) ||
		payload["smoke_covered_route_count"] != float64(3) ||
		payload["production_readiness"] != false ||
		payload["human_authorization_required"] != true ||
		payload["human_authorization_preflight_ready"] != true ||
		payload["human_authorization_granted"] != false ||
		payload["authorization_receipt_accepted"] != false ||
		payload["production_owner_enabled"] != false ||
		payload["production_activation_ready"] != false {
		t.Fatalf("unexpected production D-Bus gate review command decision: %s", output.String())
	}
	routes := payload["routes"].([]any)
	if len(routes) != 3 {
		t.Fatalf("unexpected production D-Bus gate route count: %s", output.String())
	}
	for _, item := range routes {
		route := item.(map[string]any)
		if route["smoke_coverage_ready"] != true ||
			route["owner_local_route_ready"] != true ||
			route["production_dbus_method_present"] != false ||
			route["production_dbus_exposure_ready"] != false ||
			route["write_methods_enabled"] != false ||
			route["runtime_writes_enabled"] != false ||
			route["backend_launch_enabled"] != false ||
			route["host_root_modified"] != false ||
			route["internal_details_exposed"] != false {
			t.Fatalf("unsafe production D-Bus gate route: %s", output.String())
		}
	}
	counts := payload["counts"].(map[string]any)
	if counts["total"] != float64(9) ||
		counts["passed"] != float64(9) ||
		counts["pending"] != float64(0) ||
		counts["blocked"] != float64(0) {
		t.Fatalf("unexpected production D-Bus gate review counts: %s", output.String())
	}
	if strings.Contains(output.String(), root) {
		t.Fatalf("production D-Bus gate review output must not expose project root path: %s", output.String())
	}
	for _, key := range []string{"system_service_started", "session_bus_claimed", "production_bus_claimed", "write_methods_enabled", "runtime_writes_enabled", "adapter_invocation_enabled", "backend_launch_enabled", "backend_process_started", "support_bundle_exported", "support_case_created", "notification_sent", "network_required", "host_root_modified", "privileged_container_required", "state_root_path_exposed", "raw_command_exposed", "raw_executable_exposed", "backend_details_exposed"} {
		if payload[key] != false {
			t.Fatalf("unsafe gate %s must remain false: %s", key, output.String())
		}
	}
}

func TestProductionDBusGateReviewPreviewCommandRejectsPositionalArgs(t *testing.T) {
	var output bytes.Buffer
	if err := run([]string{"production-dbus-gate-review-preview", "extra"}, &output); err == nil {
		t.Fatalf("production-dbus-gate-review-preview must reject positional arguments")
	}
}

func TestProductionDBusHumanAuthorizationPreflightPreviewCommand(t *testing.T) {
	root := projectRootForRuntimeServiceBindingCommandTest(t)
	var output bytes.Buffer
	if err := run([]string{"production-dbus-human-authorization-preflight-preview", "--root", root}, &output); err != nil {
		t.Fatalf("production-dbus-human-authorization-preflight-preview returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["version"] != currentProjectVersion(t) ||
		payload["schema_version"] != "xnix.runtime.production_dbus_human_authorization_preflight.v1" ||
		payload["request_type"] != "production-dbus-human-authorization-preflight-preview" ||
		payload["preflight_type"] != "read-only-production-dbus-human-authorization-preflight" {
		t.Fatalf("unexpected production D-Bus human authorization preflight command payload: %s", output.String())
	}
	if payload["gate_review_present"] != true ||
		payload["route_inventory_present"] != true ||
		payload["explicit_human_authorization_required"] != true ||
		payload["authorization_receipt_required"] != true ||
		payload["authorization_receipt_present"] != false ||
		payload["authorization_grant_ready"] != false ||
		payload["authorization_accepted"] != false ||
		payload["preflight_ready"] != true ||
		payload["production_readiness"] != false ||
		payload["production_dbus_gate_review_required"] != true {
		t.Fatalf("unexpected production D-Bus human authorization preflight decision: %s", output.String())
	}
	counts := payload["counts"].(map[string]any)
	if counts["total"] != float64(7) ||
		counts["passed"] != float64(7) ||
		counts["pending"] != float64(0) ||
		counts["blocked"] != float64(0) {
		t.Fatalf("unexpected production D-Bus human authorization preflight counts: %s", output.String())
	}
	if strings.Contains(output.String(), root) {
		t.Fatalf("production D-Bus human authorization preflight output must not expose project root path: %s", output.String())
	}
	for _, key := range []string{"system_service_started", "session_bus_claimed", "production_bus_claimed", "production_owner_enabled", "production_activation_ready", "write_methods_enabled", "runtime_writes_enabled", "adapter_invocation_enabled", "backend_launch_enabled", "backend_process_started", "support_bundle_exported", "support_case_created", "notification_sent", "network_required", "host_root_modified", "privileged_container_required", "state_root_path_exposed", "raw_command_exposed", "raw_executable_exposed", "backend_details_exposed"} {
		if payload[key] != false {
			t.Fatalf("unsafe gate %s must remain false: %s", key, output.String())
		}
	}
}

func TestProductionDBusHumanAuthorizationPreflightPreviewCommandRejectsPositionalArgs(t *testing.T) {
	var output bytes.Buffer
	if err := run([]string{"production-dbus-human-authorization-preflight-preview", "extra"}, &output); err == nil {
		t.Fatalf("production-dbus-human-authorization-preflight-preview must reject positional arguments")
	}
}

func TestProductionDBusMethodReviewPreviewCommand(t *testing.T) {
	root := projectRootForRuntimeServiceBindingCommandTest(t)
	var output bytes.Buffer
	if err := run([]string{"production-dbus-method-review-preview", "--root", root}, &output); err != nil {
		t.Fatalf("production-dbus-method-review-preview returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["version"] != currentProjectVersion(t) ||
		payload["schema_version"] != "xnix.runtime.production_dbus_method_review.v1" ||
		payload["request_type"] != "production-dbus-method-review-preview" ||
		payload["review_type"] != "route-by-route-production-dbus-method-review" ||
		payload["review_decision"] != "production-dbus-method-review-ready-production-exposure-disabled" {
		t.Fatalf("unexpected production D-Bus method review command payload: %s", output.String())
	}
	if payload["read_only_contract_method_count"] != float64(61) ||
		payload["owner_local_candidate_count"] != float64(3) ||
		payload["write_method_count"] != float64(4) ||
		payload["reviewed_method_count"] != float64(68) ||
		payload["production_exposure_ready_count"] != float64(0) ||
		payload["new_production_method_request_count"] != float64(0) {
		t.Fatalf("unexpected production D-Bus method review counts: %s", output.String())
	}
	readOnlyMethods := payload["read_only_methods"].([]any)
	ownerLocalCandidates := payload["owner_local_candidates"].([]any)
	writeMethods := payload["write_methods"].([]any)
	if len(readOnlyMethods) != 61 || len(ownerLocalCandidates) != 3 || len(writeMethods) != 4 {
		t.Fatalf("unexpected production D-Bus method review route lists: %s", output.String())
	}
	for _, item := range readOnlyMethods {
		route := item.(map[string]any)
		if route["route_class"] != "dbus-read-only-contract" ||
			route["current_exposure"] != "read-only-contract-method-production-owner-disabled" ||
			route["future_exposure_decision"] != "reviewed-read-only-contract-production-owner-disabled" ||
			route["new_production_method_requested"] != false ||
			route["production_exposure_ready"] != false ||
			route["write_methods_enabled"] != false ||
			route["runtime_writes_enabled"] != false ||
			route["backend_launch_enabled"] != false ||
			route["host_root_modified"] != false ||
			route["internal_details_exposed"] != false {
			t.Fatalf("unsafe read-only method review route: %s", output.String())
		}
	}
	for _, item := range ownerLocalCandidates {
		route := item.(map[string]any)
		if route["route_class"] != "owner-local-candidate" ||
			route["current_exposure"] != "owner-local-only" ||
			route["future_exposure_decision"] != "owner-local-only-no-production-dbus-method" ||
			route["new_production_method_requested"] != false ||
			route["production_exposure_ready"] != false ||
			route["write_methods_enabled"] != false ||
			route["runtime_writes_enabled"] != false ||
			route["backend_launch_enabled"] != false ||
			route["host_root_modified"] != false ||
			route["internal_details_exposed"] != false {
			t.Fatalf("unsafe owner-local candidate review route: %s", output.String())
		}
	}
	for _, item := range writeMethods {
		route := item.(map[string]any)
		if route["route_class"] != "reserved-write-method" ||
			route["current_exposure"] != "write-method-disabled" ||
			route["future_exposure_decision"] != "write-method-disabled-no-production-dispatch" ||
			route["new_production_method_requested"] != false ||
			route["production_exposure_ready"] != false ||
			route["write_methods_enabled"] != false ||
			route["runtime_writes_enabled"] != false ||
			route["backend_launch_enabled"] != false ||
			route["host_root_modified"] != false ||
			route["internal_details_exposed"] != false {
			t.Fatalf("unsafe write method review route: %s", output.String())
		}
	}
	counts := payload["counts"].(map[string]any)
	if counts["total"] != float64(7) ||
		counts["passed"] != float64(7) ||
		counts["pending"] != float64(0) ||
		counts["blocked"] != float64(0) {
		t.Fatalf("unexpected production D-Bus method review checks: %s", output.String())
	}
	if strings.Contains(output.String(), root) {
		t.Fatalf("production D-Bus method review output must not expose project root path: %s", output.String())
	}
	for _, key := range []string{"system_service_started", "session_bus_claimed", "production_bus_claimed", "production_owner_enabled", "production_activation_ready", "write_methods_enabled", "runtime_writes_enabled", "adapter_invocation_enabled", "backend_launch_enabled", "backend_process_started", "notification_sent", "network_required", "host_root_modified", "privileged_container_required", "state_root_path_exposed", "raw_command_exposed", "raw_executable_exposed", "backend_details_exposed"} {
		if payload[key] != false {
			t.Fatalf("unsafe method review gate %s must remain false: %s", key, output.String())
		}
	}
}

func TestProductionDBusMethodReviewPreviewCommandRejectsPositionalArgs(t *testing.T) {
	var output bytes.Buffer
	if err := run([]string{"production-dbus-method-review-preview", "extra"}, &output); err == nil {
		t.Fatalf("production-dbus-method-review-preview must reject positional arguments")
	}
}

func TestProductionRollbackDiagnosticsReviewPreviewCommand(t *testing.T) {
	root := projectRootForRuntimeServiceBindingCommandTest(t)
	var output bytes.Buffer
	if err := run([]string{"production-rollback-diagnostics-review-preview", "--root", root}, &output); err != nil {
		t.Fatalf("production-rollback-diagnostics-review-preview returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["version"] != currentProjectVersion(t) ||
		payload["schema_version"] != "xnix.runtime.production_rollback_diagnostics_review.v1" ||
		payload["request_type"] != "production-rollback-diagnostics-review-preview" ||
		payload["review_type"] != "production-dbus-rollback-diagnostics-review" ||
		payload["review_decision"] != "production-rollback-diagnostics-review-ready-side-effects-disabled" {
		t.Fatalf("unexpected production rollback diagnostics review command payload: %s", output.String())
	}
	if payload["review_item_count"] != float64(8) ||
		payload["rollback_control_count"] != float64(6) ||
		payload["diagnostics_control_count"] != float64(4) ||
		payload["ready_control_count"] != float64(8) ||
		payload["side_effect_control_count"] != float64(0) ||
		payload["production_ownership_ready"] != false {
		t.Fatalf("unexpected production rollback diagnostics review counts: %s", output.String())
	}
	items := payload["items"].([]any)
	if len(items) != 8 {
		t.Fatalf("unexpected rollback diagnostics review item count: %s", output.String())
	}
	for _, itemValue := range items {
		item := itemValue.(map[string]any)
		if item["required_before_production"] != true ||
			item["evidence_present"] != true ||
			item["review_only"] != true ||
			item["side_effects_enabled"] != false ||
			item["restore_executed"] != false ||
			item["cleanup_executed"] != false ||
			item["support_bundle_exported"] != false ||
			item["support_case_created"] != false ||
			item["notification_sent"] != false ||
			item["file_content_read"] != false ||
			item["file_paths_exposed"] != false ||
			item["state_root_path_exposed"] != false ||
			item["raw_command_exposed"] != false ||
			item["raw_executable_exposed"] != false ||
			item["backend_details_exposed"] != false ||
			item["host_root_modified"] != false ||
			item["production_bus_claimed"] != false ||
			item["write_methods_enabled"] != false ||
			item["runtime_writes_enabled"] != false ||
			item["backend_launch_enabled"] != false ||
			item["review_status"] != "reviewed" {
			t.Fatalf("unsafe rollback diagnostics review item: %s", output.String())
		}
	}
	counts := payload["counts"].(map[string]any)
	if counts["total"] != float64(8) ||
		counts["passed"] != float64(8) ||
		counts["pending"] != float64(0) ||
		counts["blocked"] != float64(0) {
		t.Fatalf("unexpected production rollback diagnostics review checks: %s", output.String())
	}
	if strings.Contains(output.String(), root) {
		t.Fatalf("production rollback diagnostics review output must not expose project root path: %s", output.String())
	}
	for _, key := range []string{"system_service_started", "session_bus_claimed", "production_bus_claimed", "production_owner_enabled", "production_activation_ready", "write_methods_enabled", "runtime_writes_enabled", "adapter_invocation_enabled", "backend_launch_enabled", "backend_process_started", "support_bundle_exported", "support_case_created", "notification_sent", "snapshot_restore_executed", "state_cleanup_executed", "file_content_read", "file_paths_exposed", "network_required", "host_root_modified", "privileged_container_required", "state_root_path_exposed", "raw_command_exposed", "raw_executable_exposed", "backend_details_exposed"} {
		if payload[key] != false {
			t.Fatalf("unsafe rollback diagnostics gate %s must remain false: %s", key, output.String())
		}
	}
}

func TestProductionRollbackDiagnosticsReviewPreviewCommandRejectsPositionalArgs(t *testing.T) {
	var output bytes.Buffer
	if err := run([]string{"production-rollback-diagnostics-review-preview", "extra"}, &output); err == nil {
		t.Fatalf("production-rollback-diagnostics-review-preview must reject positional arguments")
	}
}
