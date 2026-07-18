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

func TestProductionHumanAuthorizationReceiptConsolidationPreviewCommand(t *testing.T) {
	root := projectRootForRuntimeServiceBindingCommandTest(t)
	var output bytes.Buffer
	if err := run([]string{"production-human-authorization-receipt-consolidation-preview", "--root", root}, &output); err != nil {
		t.Fatalf("production-human-authorization-receipt-consolidation-preview returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["version"] != currentProjectVersion(t) ||
		payload["schema_version"] != "xnix.runtime.production_human_authorization_receipt_consolidation.v1" ||
		payload["request_type"] != "production-human-authorization-receipt-consolidation-preview" ||
		payload["consolidation_type"] != "owner-managed-opaque-human-authorization-receipt-boundary" ||
		payload["consolidation_decision"] != "production-human-authorization-receipt-consolidation-ready-authorization-disabled" ||
		payload["receipt_schema"] != "xnix.runtime.production_dbus_human_authorization_receipt.v1" ||
		payload["opaque_receipt_id"] != "production-dbus-human-authorization-receipt-id" {
		t.Fatalf("unexpected production human authorization receipt consolidation command payload: %s", output.String())
	}
	if payload["receipt_required"] != true ||
		payload["receipt_present"] != false ||
		payload["receipt_accepted"] != false ||
		payload["receipt_boundary_consolidated"] != true ||
		payload["owner_managed_opaque_receipt_lookup_ready"] != true ||
		payload["caller_state_root_required"] != false ||
		payload["explicit_operator_action_required"] != true ||
		payload["authorization_grant_ready"] != false ||
		payload["authorization_accepted"] != false ||
		payload["production_readiness"] != false ||
		payload["production_ownership_ready"] != false {
		t.Fatalf("unexpected production human authorization receipt consolidation decision: %s", output.String())
	}
	if payload["gate_count"] != float64(7) ||
		payload["required_gate_count"] != float64(7) ||
		payload["consumed_gate_count"] != float64(7) ||
		payload["missing_gate_count"] != float64(0) ||
		payload["authorization_accepted_gate_count"] != float64(0) ||
		payload["production_ready_gate_count"] != float64(0) {
		t.Fatalf("unexpected production human authorization receipt gate counts: %s", output.String())
	}
	gates := payload["gates"].([]any)
	if len(gates) != 7 {
		t.Fatalf("unexpected production human authorization receipt gate list: %s", output.String())
	}
	for _, item := range gates {
		gate := item.(map[string]any)
		if gate["evidence_present"] != true ||
			gate["receipt_boundary_ready"] != true ||
			gate["human_authorization_required"] != true ||
			gate["authorization_receipt_accepted"] != false ||
			gate["production_readiness"] != false ||
			gate["production_ownership_ready"] != false ||
			gate["runtime_owned"] != true ||
			gate["go_runtime_backed"] != true ||
			gate["kde_policy_owner"] != false ||
			gate["review_only"] != true ||
			gate["side_effects_disabled"] != true ||
			gate["write_methods_enabled"] != false ||
			gate["runtime_writes_enabled"] != false ||
			gate["backend_launch_enabled"] != false ||
			gate["host_root_modified"] != false ||
			gate["internal_details_exposed"] != false {
			t.Fatalf("unsafe production human authorization receipt gate: %s", output.String())
		}
	}
	counts := payload["counts"].(map[string]any)
	if counts["total"] != float64(8) ||
		counts["passed"] != float64(8) ||
		counts["pending"] != float64(0) ||
		counts["blocked"] != float64(0) {
		t.Fatalf("unexpected production human authorization receipt consolidation checks: %s", output.String())
	}
	if strings.Contains(output.String(), root) {
		t.Fatalf("production human authorization receipt consolidation output must not expose project root path: %s", output.String())
	}
	for _, key := range []string{"system_service_started", "session_bus_claimed", "production_bus_claimed", "production_owner_enabled", "production_activation_ready", "write_methods_enabled", "runtime_writes_enabled", "receipt_writer_enabled", "receipt_persistence_enabled", "receipt_lookup_writes_enabled", "desktop_files_written", "mimeapps_written", "shell_configuration_written", "settings_persisted", "notification_sent", "notification_delivery_enabled", "portal_request_created", "request_objects_created", "adapter_invocation_enabled", "backend_launch_enabled", "backend_process_started", "support_bundle_exported", "support_case_created", "snapshot_restore_executed", "state_cleanup_executed", "file_content_read", "file_paths_exposed", "network_required", "host_root_modified", "privileged_container_required", "state_root_path_exposed", "raw_command_exposed", "raw_executable_exposed", "backend_details_exposed"} {
		if payload[key] != false {
			t.Fatalf("unsafe production human authorization receipt gate %s must remain false: %s", key, output.String())
		}
	}
}

func TestProductionHumanAuthorizationReceiptConsolidationPreviewCommandRejectsPositionalArgs(t *testing.T) {
	var output bytes.Buffer
	if err := run([]string{"production-human-authorization-receipt-consolidation-preview", "extra"}, &output); err == nil {
		t.Fatalf("production-human-authorization-receipt-consolidation-preview must reject positional arguments")
	}
}

func TestProductionAuthorizationConsumptionAuditPreviewCommand(t *testing.T) {
	root := projectRootForRuntimeServiceBindingCommandTest(t)
	var output bytes.Buffer
	if err := run([]string{"production-authorization-consumption-audit-preview", "--root", root}, &output); err != nil {
		t.Fatalf("production-authorization-consumption-audit-preview returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["version"] != currentProjectVersion(t) ||
		payload["schema_version"] != "xnix.runtime.production_authorization_consumption_audit.v1" ||
		payload["request_type"] != "production-authorization-consumption-audit-preview" ||
		payload["audit_type"] != "production-gate-consolidated-authorization-consumption-audit" ||
		payload["audit_decision"] != "production-authorization-consumption-audit-ready-authorization-disabled" ||
		payload["receipt_schema"] != "xnix.runtime.production_dbus_human_authorization_receipt.v1" ||
		payload["opaque_receipt_id"] != "production-dbus-human-authorization-receipt-id" {
		t.Fatalf("unexpected production authorization consumption audit command payload: %s", output.String())
	}
	if payload["receipt_required"] != true ||
		payload["receipt_present"] != false ||
		payload["receipt_accepted"] != false ||
		payload["receipt_boundary_consolidated"] != true ||
		payload["consolidation_preview_consumed"] != true ||
		payload["owner_managed_opaque_boundary_ready"] != true ||
		payload["caller_state_root_required"] != false ||
		payload["authorization_accepted"] != false ||
		payload["production_readiness"] != false ||
		payload["production_ownership_ready"] != false {
		t.Fatalf("unexpected production authorization consumption audit decision: %s", output.String())
	}
	if payload["consumer_count"] != float64(6) ||
		payload["required_consumer_count"] != float64(6) ||
		payload["consumed_consumer_count"] != float64(6) ||
		payload["missing_consumer_count"] != float64(0) ||
		payload["authorization_accepted_consumer_count"] != float64(0) ||
		payload["production_ready_consumer_count"] != float64(0) ||
		payload["side_effect_consumer_count"] != float64(0) {
		t.Fatalf("unexpected production authorization consumption counts: %s", output.String())
	}
	consumers := payload["consumers"].([]any)
	if len(consumers) != 6 {
		t.Fatalf("unexpected production authorization consumer list: %s", output.String())
	}
	for _, item := range consumers {
		consumer := item.(map[string]any)
		if consumer["consumes_consolidated_boundary"] != true ||
			consumer["receipt_boundary_ready"] != true ||
			consumer["receipt_accepted"] != false ||
			consumer["authorization_accepted"] != false ||
			consumer["production_readiness"] != false ||
			consumer["production_ownership_ready"] != false ||
			consumer["runtime_owned"] != true ||
			consumer["go_runtime_backed"] != true ||
			consumer["kde_policy_owner"] != false ||
			consumer["review_only"] != true ||
			consumer["write_methods_enabled"] != false ||
			consumer["runtime_writes_enabled"] != false ||
			consumer["desktop_side_effects_enabled"] != false ||
			consumer["support_side_effects_enabled"] != false ||
			consumer["backend_launch_enabled"] != false ||
			consumer["host_root_modified"] != false ||
			consumer["internal_details_exposed"] != false ||
			consumer["audit_status"] != "consumed-authorization-disabled" {
			t.Fatalf("unsafe production authorization consumer: %s", output.String())
		}
	}
	counts := payload["counts"].(map[string]any)
	if counts["total"] != float64(8) ||
		counts["passed"] != float64(8) ||
		counts["pending"] != float64(0) ||
		counts["blocked"] != float64(0) {
		t.Fatalf("unexpected production authorization consumption audit checks: %s", output.String())
	}
	if strings.Contains(output.String(), root) {
		t.Fatalf("production authorization consumption audit output must not expose project root path: %s", output.String())
	}
	for _, key := range []string{"system_service_started", "session_bus_claimed", "production_bus_claimed", "production_owner_enabled", "production_activation_ready", "write_methods_enabled", "runtime_writes_enabled", "receipt_writer_enabled", "receipt_persistence_enabled", "receipt_lookup_writes_enabled", "desktop_files_written", "mimeapps_written", "shell_configuration_written", "settings_persisted", "notification_sent", "notification_delivery_enabled", "portal_request_created", "request_objects_created", "adapter_invocation_enabled", "backend_launch_enabled", "backend_process_started", "support_bundle_exported", "support_case_created", "snapshot_restore_executed", "state_cleanup_executed", "file_content_read", "file_paths_exposed", "network_required", "host_root_modified", "privileged_container_required", "state_root_path_exposed", "raw_command_exposed", "raw_executable_exposed", "backend_details_exposed"} {
		if payload[key] != false {
			t.Fatalf("unsafe production authorization audit gate %s must remain false: %s", key, output.String())
		}
	}
}

func TestProductionAuthorizationConsumptionAuditPreviewCommandRejectsPositionalArgs(t *testing.T) {
	var output bytes.Buffer
	if err := run([]string{"production-authorization-consumption-audit-preview", "extra"}, &output); err == nil {
		t.Fatalf("production-authorization-consumption-audit-preview must reject positional arguments")
	}
}

func TestProductionReceiptAcceptancePropagationPreflightPreviewCommand(t *testing.T) {
	root := projectRootForRuntimeServiceBindingCommandTest(t)
	var output bytes.Buffer
	if err := run([]string{"production-receipt-acceptance-propagation-preflight-preview", "--root", root}, &output); err != nil {
		t.Fatalf("production-receipt-acceptance-propagation-preflight-preview returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["version"] != currentProjectVersion(t) ||
		payload["schema_version"] != "xnix.runtime.production_receipt_acceptance_propagation_preflight.v1" ||
		payload["request_type"] != "production-receipt-acceptance-propagation-preflight-preview" ||
		payload["preflight_type"] != "future-authorization-receipt-acceptance-propagation-preflight" ||
		payload["preflight_decision"] != "production-receipt-acceptance-propagation-ready-acceptance-disabled" ||
		payload["receipt_schema"] != "xnix.runtime.production_dbus_human_authorization_receipt.v1" ||
		payload["opaque_receipt_id"] != "production-dbus-human-authorization-receipt-id" {
		t.Fatalf("unexpected production receipt acceptance propagation command payload: %s", output.String())
	}
	if payload["receipt_required"] != true ||
		payload["receipt_present"] != false ||
		payload["receipt_accepted"] != false ||
		payload["future_acceptance_modeled"] != true ||
		payload["acceptance_simulation_only"] != true ||
		payload["consumption_audit_consumed"] != true ||
		payload["owner_managed_opaque_boundary_ready"] != true ||
		payload["caller_state_root_required"] != false ||
		payload["authorization_accepted"] != false ||
		payload["propagation_preflight_ready"] != true ||
		payload["production_readiness"] != false ||
		payload["production_ownership_ready"] != false {
		t.Fatalf("unexpected production receipt acceptance propagation decision: %s", output.String())
	}
	if payload["target_count"] != float64(6) ||
		payload["required_target_count"] != float64(6) ||
		payload["propagation_ready_target_count"] != float64(6) ||
		payload["missing_target_count"] != float64(0) ||
		payload["acceptance_enabled_target_count"] != float64(0) ||
		payload["production_ready_target_count"] != float64(0) ||
		payload["side_effect_target_count"] != float64(0) {
		t.Fatalf("unexpected production receipt acceptance propagation counts: %s", output.String())
	}
	targets := payload["targets"].([]any)
	if len(targets) != 6 {
		t.Fatalf("unexpected production receipt acceptance target list: %s", output.String())
	}
	for _, item := range targets {
		target := item.(map[string]any)
		if target["consumes_audit_boundary"] != true ||
			target["propagates_future_acceptance"] != true ||
			target["receipt_boundary_ready"] != true ||
			target["future_acceptance_modeled"] != true ||
			target["receipt_accepted"] != false ||
			target["authorization_accepted"] != false ||
			target["production_readiness"] != false ||
			target["production_ownership_ready"] != false ||
			target["runtime_owned"] != true ||
			target["go_runtime_backed"] != true ||
			target["kde_policy_owner"] != false ||
			target["review_only"] != true ||
			target["write_methods_enabled"] != false ||
			target["runtime_writes_enabled"] != false ||
			target["desktop_side_effects_enabled"] != false ||
			target["support_side_effects_enabled"] != false ||
			target["backend_launch_enabled"] != false ||
			target["host_root_modified"] != false ||
			target["internal_details_exposed"] != false ||
			target["propagation_status"] != "future-acceptance-modeled-side-effects-disabled" {
			t.Fatalf("unsafe production receipt acceptance propagation target: %s", output.String())
		}
	}
	counts := payload["counts"].(map[string]any)
	if counts["total"] != float64(8) ||
		counts["passed"] != float64(8) ||
		counts["pending"] != float64(0) ||
		counts["blocked"] != float64(0) {
		t.Fatalf("unexpected production receipt acceptance propagation checks: %s", output.String())
	}
	if strings.Contains(output.String(), root) {
		t.Fatalf("production receipt acceptance propagation output must not expose project root path: %s", output.String())
	}
	for _, key := range []string{"system_service_started", "session_bus_claimed", "production_bus_claimed", "production_owner_enabled", "production_activation_ready", "write_methods_enabled", "runtime_writes_enabled", "receipt_writer_enabled", "receipt_persistence_enabled", "receipt_lookup_writes_enabled", "desktop_files_written", "mimeapps_written", "shell_configuration_written", "settings_persisted", "notification_sent", "notification_delivery_enabled", "portal_request_created", "request_objects_created", "adapter_invocation_enabled", "backend_launch_enabled", "backend_process_started", "support_bundle_exported", "support_case_created", "snapshot_restore_executed", "state_cleanup_executed", "file_content_read", "file_paths_exposed", "network_required", "host_root_modified", "privileged_container_required", "state_root_path_exposed", "raw_command_exposed", "raw_executable_exposed", "backend_details_exposed"} {
		if payload[key] != false {
			t.Fatalf("unsafe production receipt acceptance propagation gate %s must remain false: %s", key, output.String())
		}
	}
}

func TestProductionReceiptAcceptancePropagationPreflightPreviewCommandRejectsPositionalArgs(t *testing.T) {
	var output bytes.Buffer
	if err := run([]string{"production-receipt-acceptance-propagation-preflight-preview", "extra"}, &output); err == nil {
		t.Fatalf("production-receipt-acceptance-propagation-preflight-preview must reject positional arguments")
	}
}

func TestProductionReceiptWriterAuthorizationReviewPreviewCommand(t *testing.T) {
	root := projectRootForRuntimeServiceBindingCommandTest(t)
	var output bytes.Buffer
	if err := run([]string{"production-receipt-writer-authorization-review-preview", "--root", root}, &output); err != nil {
		t.Fatalf("production-receipt-writer-authorization-review-preview returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["version"] != currentProjectVersion(t) ||
		payload["schema_version"] != "xnix.runtime.production_receipt_writer_authorization_review.v1" ||
		payload["request_type"] != "production-receipt-writer-authorization-review-preview" ||
		payload["review_type"] != "receipt-writer-operator-authorization-boundary-review" ||
		payload["review_decision"] != "production-receipt-writer-authorization-review-ready-writes-disabled" ||
		payload["receipt_schema"] != "xnix.runtime.production_dbus_human_authorization_receipt.v1" ||
		payload["opaque_receipt_id"] != "production-dbus-human-authorization-receipt-id" {
		t.Fatalf("unexpected production receipt writer authorization review payload: %s", output.String())
	}
	if payload["receipt_required"] != true ||
		payload["receipt_present"] != false ||
		payload["receipt_accepted"] != false ||
		payload["writer_authorization_required"] != true ||
		payload["writer_authorization_modeled"] != true ||
		payload["operator_action_required"] != true ||
		payload["acceptance_propagation_consumed"] != true ||
		payload["owner_managed_opaque_boundary_ready"] != true ||
		payload["caller_state_root_required"] != false ||
		payload["authorization_accepted"] != false ||
		payload["writer_review_ready"] != true ||
		payload["production_readiness"] != false ||
		payload["production_ownership_ready"] != false {
		t.Fatalf("unexpected production receipt writer authorization decision: %s", output.String())
	}
	if payload["review_item_count"] != float64(5) ||
		payload["required_review_item_count"] != float64(5) ||
		payload["ready_review_item_count"] != float64(5) ||
		payload["missing_review_item_count"] != float64(0) ||
		payload["write_enabled_review_item_count"] != float64(0) ||
		payload["acceptance_enabled_review_item_count"] != float64(0) ||
		payload["side_effect_review_item_count"] != float64(0) {
		t.Fatalf("unexpected production receipt writer authorization counts: %s", output.String())
	}
	items := payload["review_items"].([]any)
	if len(items) != 5 {
		t.Fatalf("unexpected production receipt writer authorization item list: %s", output.String())
	}
	for _, item := range items {
		review := item.(map[string]any)
		if review["evidence_present"] != true ||
			review["writer_authorization_modeled"] != true ||
			review["receipt_writer_enabled"] != false ||
			review["receipt_persistence_enabled"] != false ||
			review["receipt_lookup_writes_enabled"] != false ||
			review["receipt_replay_enabled"] != false ||
			review["receipt_accepted"] != false ||
			review["authorization_accepted"] != false ||
			review["production_readiness"] != false ||
			review["production_ownership_ready"] != false ||
			review["runtime_owned"] != true ||
			review["go_runtime_backed"] != true ||
			review["kde_policy_owner"] != false ||
			review["review_only"] != true ||
			review["side_effects_disabled"] != true ||
			review["host_root_modified"] != false ||
			review["internal_details_exposed"] != false ||
			review["review_status"] != "writer-authorization-reviewed-writes-disabled" {
			t.Fatalf("unsafe production receipt writer authorization review item: %s", output.String())
		}
	}
	counts := payload["counts"].(map[string]any)
	if counts["total"] != float64(8) ||
		counts["passed"] != float64(8) ||
		counts["pending"] != float64(0) ||
		counts["blocked"] != float64(0) {
		t.Fatalf("unexpected production receipt writer authorization checks: %s", output.String())
	}
	if strings.Contains(output.String(), root) {
		t.Fatalf("production receipt writer authorization output must not expose project root path: %s", output.String())
	}
	for _, key := range []string{"system_service_started", "session_bus_claimed", "production_bus_claimed", "production_owner_enabled", "production_activation_ready", "write_methods_enabled", "runtime_writes_enabled", "receipt_writer_enabled", "receipt_persistence_enabled", "receipt_lookup_writes_enabled", "receipt_replay_enabled", "desktop_files_written", "mimeapps_written", "shell_configuration_written", "settings_persisted", "notification_sent", "notification_delivery_enabled", "portal_request_created", "request_objects_created", "adapter_invocation_enabled", "backend_launch_enabled", "backend_process_started", "support_bundle_exported", "support_case_created", "snapshot_restore_executed", "state_cleanup_executed", "file_content_read", "file_paths_exposed", "network_required", "host_root_modified", "privileged_container_required", "state_root_path_exposed", "raw_command_exposed", "raw_executable_exposed", "backend_details_exposed"} {
		if payload[key] != false {
			t.Fatalf("unsafe production receipt writer authorization gate %s must remain false: %s", key, output.String())
		}
	}
}

func TestProductionReceiptWriterAuthorizationReviewPreviewCommandRejectsPositionalArgs(t *testing.T) {
	var output bytes.Buffer
	if err := run([]string{"production-receipt-writer-authorization-review-preview", "extra"}, &output); err == nil {
		t.Fatalf("production-receipt-writer-authorization-review-preview must reject positional arguments")
	}
}

func TestProductionReceiptPersistenceThreatReviewPreviewCommand(t *testing.T) {
	root := projectRootForRuntimeServiceBindingCommandTest(t)
	var output bytes.Buffer
	if err := run([]string{"production-receipt-persistence-threat-review-preview", "--root", root}, &output); err != nil {
		t.Fatalf("production-receipt-persistence-threat-review-preview returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["version"] != currentProjectVersion(t) ||
		payload["schema_version"] != "xnix.runtime.production_receipt_persistence_threat_review.v1" ||
		payload["request_type"] != "production-receipt-persistence-threat-review-preview" ||
		payload["review_type"] != "receipt-persistence-expiry-revocation-replay-threat-review" ||
		payload["review_decision"] != "production-receipt-persistence-threat-review-ready-persistence-disabled" ||
		payload["receipt_schema"] != "xnix.runtime.production_dbus_human_authorization_receipt.v1" ||
		payload["opaque_receipt_id"] != "production-dbus-human-authorization-receipt-id" {
		t.Fatalf("unexpected production receipt persistence threat review payload: %s", output.String())
	}
	if payload["receipt_required"] != true ||
		payload["receipt_present"] != false ||
		payload["receipt_accepted"] != false ||
		payload["persistence_threat_review_required"] != true ||
		payload["persistence_threat_review_modeled"] != true ||
		payload["writer_authorization_review_consumed"] != true ||
		payload["owner_managed_opaque_boundary_ready"] != true ||
		payload["caller_state_root_required"] != false ||
		payload["authorization_accepted"] != false ||
		payload["persistence_threat_review_ready"] != true ||
		payload["production_readiness"] != false ||
		payload["production_ownership_ready"] != false {
		t.Fatalf("unexpected production receipt persistence threat decision: %s", output.String())
	}
	if payload["threat_item_count"] != float64(5) ||
		payload["required_threat_item_count"] != float64(5) ||
		payload["ready_threat_item_count"] != float64(5) ||
		payload["missing_threat_item_count"] != float64(0) ||
		payload["persistence_enabled_threat_count"] != float64(0) ||
		payload["replay_enabled_threat_count"] != float64(0) ||
		payload["acceptance_enabled_threat_count"] != float64(0) ||
		payload["side_effect_threat_count"] != float64(0) {
		t.Fatalf("unexpected production receipt persistence threat counts: %s", output.String())
	}
	items := payload["threat_items"].([]any)
	if len(items) != 5 {
		t.Fatalf("unexpected production receipt persistence threat item list: %s", output.String())
	}
	for _, item := range items {
		threat := item.(map[string]any)
		if threat["evidence_present"] != true ||
			threat["threat_modeled"] != true ||
			threat["receipt_persistence_enabled"] != false ||
			threat["receipt_lookup_writes_enabled"] != false ||
			threat["receipt_replay_enabled"] != false ||
			threat["receipt_expiry_write_enabled"] != false ||
			threat["receipt_revocation_write_enabled"] != false ||
			threat["receipt_accepted"] != false ||
			threat["authorization_accepted"] != false ||
			threat["production_readiness"] != false ||
			threat["production_ownership_ready"] != false ||
			threat["runtime_owned"] != true ||
			threat["go_runtime_backed"] != true ||
			threat["kde_policy_owner"] != false ||
			threat["review_only"] != true ||
			threat["side_effects_disabled"] != true ||
			threat["host_root_modified"] != false ||
			threat["internal_details_exposed"] != false ||
			threat["threat_status"] != "threat-modeled-persistence-disabled" {
			t.Fatalf("unsafe production receipt persistence threat item: %s", output.String())
		}
	}
	counts := payload["counts"].(map[string]any)
	if counts["total"] != float64(8) ||
		counts["passed"] != float64(8) ||
		counts["pending"] != float64(0) ||
		counts["blocked"] != float64(0) {
		t.Fatalf("unexpected production receipt persistence threat checks: %s", output.String())
	}
	if strings.Contains(output.String(), root) {
		t.Fatalf("production receipt persistence threat output must not expose project root path: %s", output.String())
	}
	for _, key := range []string{"system_service_started", "session_bus_claimed", "production_bus_claimed", "production_owner_enabled", "production_activation_ready", "write_methods_enabled", "runtime_writes_enabled", "receipt_writer_enabled", "receipt_persistence_enabled", "receipt_lookup_writes_enabled", "receipt_replay_enabled", "receipt_expiry_write_enabled", "receipt_revocation_write_enabled", "desktop_files_written", "mimeapps_written", "shell_configuration_written", "settings_persisted", "notification_sent", "notification_delivery_enabled", "portal_request_created", "request_objects_created", "adapter_invocation_enabled", "backend_launch_enabled", "backend_process_started", "support_bundle_exported", "support_case_created", "snapshot_restore_executed", "state_cleanup_executed", "file_content_read", "file_paths_exposed", "network_required", "host_root_modified", "privileged_container_required", "state_root_path_exposed", "raw_command_exposed", "raw_executable_exposed", "backend_details_exposed"} {
		if payload[key] != false {
			t.Fatalf("unsafe production receipt persistence threat gate %s must remain false: %s", key, output.String())
		}
	}
}

func TestProductionReceiptPersistenceThreatReviewPreviewCommandRejectsPositionalArgs(t *testing.T) {
	var output bytes.Buffer
	if err := run([]string{"production-receipt-persistence-threat-review-preview", "extra"}, &output); err == nil {
		t.Fatalf("production-receipt-persistence-threat-review-preview must reject positional arguments")
	}
}

func TestProductionReceiptRevocationVisibilityAuditPreviewCommand(t *testing.T) {
	root := projectRootForRuntimeServiceBindingCommandTest(t)
	var output bytes.Buffer
	if err := run([]string{"production-receipt-revocation-visibility-audit-preview", "--root", root}, &output); err != nil {
		t.Fatalf("production-receipt-revocation-visibility-audit-preview returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["version"] != currentProjectVersion(t) ||
		payload["schema_version"] != "xnix.runtime.production_receipt_revocation_visibility_audit.v1" ||
		payload["request_type"] != "production-receipt-revocation-visibility-audit-preview" ||
		payload["audit_type"] != "receipt-expiry-revocation-production-kde-visibility-audit" ||
		payload["audit_decision"] != "production-receipt-revocation-visibility-audit-ready-visibility-only" ||
		payload["receipt_schema"] != "xnix.runtime.production_dbus_human_authorization_receipt.v1" ||
		payload["opaque_receipt_id"] != "production-dbus-human-authorization-receipt-id" {
		t.Fatalf("unexpected production receipt revocation visibility payload: %s", output.String())
	}
	if payload["visibility_audit_required"] != true ||
		payload["visibility_audit_modeled"] != true ||
		payload["persistence_threat_review_consumed"] != true ||
		payload["desktop_side_effect_review_consumed"] != true ||
		payload["production_gate_visibility_modeled"] != true ||
		payload["kde_status_visibility_modeled"] != true ||
		payload["caller_state_root_required"] != false ||
		payload["receipt_present"] != false ||
		payload["receipt_accepted"] != false ||
		payload["authorization_accepted"] != false ||
		payload["revocation_visibility_ready"] != true ||
		payload["production_readiness"] != false ||
		payload["production_ownership_ready"] != false {
		t.Fatalf("unexpected production receipt revocation visibility decision: %s", output.String())
	}
	if payload["visibility_item_count"] != float64(5) ||
		payload["required_visibility_item_count"] != float64(5) ||
		payload["ready_visibility_item_count"] != float64(5) ||
		payload["missing_visibility_item_count"] != float64(0) ||
		payload["revocation_write_enabled_item_count"] != float64(0) ||
		payload["expiry_write_enabled_item_count"] != float64(0) ||
		payload["notification_enabled_item_count"] != float64(0) ||
		payload["side_effect_item_count"] != float64(0) {
		t.Fatalf("unexpected production receipt revocation visibility counts: %s", output.String())
	}
	items := payload["visibility_items"].([]any)
	if len(items) != 5 {
		t.Fatalf("unexpected production receipt revocation visibility item list: %s", output.String())
	}
	for _, item := range items {
		visibility := item.(map[string]any)
		if visibility["evidence_present"] != true ||
			visibility["visibility_modeled"] != true ||
			visibility["user_visible"] != true ||
			visibility["production_gate_visible"] != true ||
			visibility["kde_status_visible"] != true ||
			visibility["review_only"] != true ||
			visibility["runtime_owned"] != true ||
			visibility["go_runtime_backed"] != true ||
			visibility["kde_policy_owner"] != false ||
			visibility["receipt_revocation_write_enabled"] != false ||
			visibility["receipt_expiry_write_enabled"] != false ||
			visibility["receipt_persistence_enabled"] != false ||
			visibility["receipt_lookup_writes_enabled"] != false ||
			visibility["receipt_replay_enabled"] != false ||
			visibility["receipt_accepted"] != false ||
			visibility["authorization_accepted"] != false ||
			visibility["notification_sent"] != false ||
			visibility["notification_delivery_enabled"] != false ||
			visibility["desktop_files_written"] != false ||
			visibility["settings_persisted"] != false ||
			visibility["production_readiness"] != false ||
			visibility["production_ownership_ready"] != false ||
			visibility["side_effects_disabled"] != true ||
			visibility["host_root_modified"] != false ||
			visibility["internal_details_exposed"] != false ||
			visibility["visibility_status"] != "visibility-modeled-writes-disabled" {
			t.Fatalf("unsafe production receipt revocation visibility item: %s", output.String())
		}
	}
	counts := payload["counts"].(map[string]any)
	if counts["total"] != float64(9) ||
		counts["passed"] != float64(9) ||
		counts["pending"] != float64(0) ||
		counts["blocked"] != float64(0) {
		t.Fatalf("unexpected production receipt revocation visibility checks: %s", output.String())
	}
	if strings.Contains(output.String(), root) {
		t.Fatalf("production receipt revocation visibility output must not expose project root path: %s", output.String())
	}
	for _, key := range []string{"system_service_started", "session_bus_claimed", "production_bus_claimed", "production_owner_enabled", "production_activation_ready", "write_methods_enabled", "runtime_writes_enabled", "receipt_writer_enabled", "receipt_persistence_enabled", "receipt_lookup_writes_enabled", "receipt_replay_enabled", "receipt_expiry_write_enabled", "receipt_revocation_write_enabled", "receipt_revocation_visibility_persisted", "desktop_files_written", "mimeapps_written", "shell_configuration_written", "settings_persisted", "krunner_index_persisted", "task_manager_entry_active", "kwin_rule_applied", "live_tray_bridge_enabled", "tray_bridge_persisted", "notification_sent", "notification_delivery_enabled", "compatibility_center_persisted", "portal_request_created", "request_objects_created", "adapter_invocation_enabled", "backend_launch_enabled", "backend_process_started", "support_bundle_exported", "support_case_created", "snapshot_restore_executed", "state_cleanup_executed", "file_content_read", "file_paths_exposed", "network_required", "host_root_modified", "privileged_container_required", "state_root_path_exposed", "raw_command_exposed", "raw_executable_exposed", "backend_details_exposed"} {
		if payload[key] != false {
			t.Fatalf("unsafe production receipt revocation visibility gate %s must remain false: %s", key, output.String())
		}
	}
}

func TestProductionReceiptRevocationVisibilityAuditPreviewCommandRejectsPositionalArgs(t *testing.T) {
	var output bytes.Buffer
	if err := run([]string{"production-receipt-revocation-visibility-audit-preview", "extra"}, &output); err == nil {
		t.Fatalf("production-receipt-revocation-visibility-audit-preview must reject positional arguments")
	}
}

func TestProductionReceiptNotificationDeliveryGateAuditPreviewCommand(t *testing.T) {
	root := projectRootForRuntimeServiceBindingCommandTest(t)
	var output bytes.Buffer
	if err := run([]string{"production-receipt-notification-delivery-gate-audit-preview", "--root", root}, &output); err != nil {
		t.Fatalf("production-receipt-notification-delivery-gate-audit-preview returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["version"] != currentProjectVersion(t) ||
		payload["schema_version"] != "xnix.runtime.production_receipt_notification_delivery_gate_audit.v1" ||
		payload["request_type"] != "production-receipt-notification-delivery-gate-audit-preview" ||
		payload["audit_type"] != "receipt-expiry-revocation-notification-delivery-gate-audit" ||
		payload["audit_decision"] != "production-receipt-notification-delivery-gate-audit-ready-delivery-disabled" ||
		payload["receipt_schema"] != "xnix.runtime.production_dbus_human_authorization_receipt.v1" ||
		payload["opaque_receipt_id"] != "production-dbus-human-authorization-receipt-id" {
		t.Fatalf("unexpected production receipt notification delivery payload: %s", output.String())
	}
	if payload["notification_gate_required"] != true ||
		payload["notification_gate_modeled"] != true ||
		payload["revocation_visibility_audit_consumed"] != true ||
		payload["desktop_side_effect_review_consumed"] != true ||
		payload["operator_notification_approval_required"] != true ||
		payload["operator_notification_approval_present"] != false ||
		payload["notification_delivery_gate_ready"] != true ||
		payload["caller_state_root_required"] != false ||
		payload["receipt_present"] != false ||
		payload["receipt_accepted"] != false ||
		payload["authorization_accepted"] != false ||
		payload["production_readiness"] != false ||
		payload["production_ownership_ready"] != false {
		t.Fatalf("unexpected production receipt notification delivery decision: %s", output.String())
	}
	if payload["gate_item_count"] != float64(5) ||
		payload["required_gate_item_count"] != float64(5) ||
		payload["ready_gate_item_count"] != float64(5) ||
		payload["missing_gate_item_count"] != float64(0) ||
		payload["delivery_enabled_item_count"] != float64(0) ||
		payload["notification_sent_item_count"] != float64(0) ||
		payload["request_created_item_count"] != float64(0) ||
		payload["side_effect_item_count"] != float64(0) {
		t.Fatalf("unexpected production receipt notification delivery counts: %s", output.String())
	}
	items := payload["gate_items"].([]any)
	if len(items) != 5 {
		t.Fatalf("unexpected production receipt notification delivery item list: %s", output.String())
	}
	for _, item := range items {
		gate := item.(map[string]any)
		if gate["evidence_present"] != true ||
			gate["gate_modeled"] != true ||
			gate["user_visible"] != true ||
			gate["review_only"] != true ||
			gate["runtime_owned"] != true ||
			gate["go_runtime_backed"] != true ||
			gate["kde_policy_owner"] != false ||
			gate["operator_approval_required"] != true ||
			gate["operator_approval_present"] != false ||
			gate["notification_sent"] != false ||
			gate["notification_delivery_enabled"] != false ||
			gate["notification_action_enabled"] != false ||
			gate["portal_request_created"] != false ||
			gate["request_objects_created"] != false ||
			gate["receipt_revocation_write_enabled"] != false ||
			gate["receipt_expiry_write_enabled"] != false ||
			gate["receipt_persistence_enabled"] != false ||
			gate["receipt_lookup_writes_enabled"] != false ||
			gate["receipt_accepted"] != false ||
			gate["authorization_accepted"] != false ||
			gate["production_readiness"] != false ||
			gate["production_ownership_ready"] != false ||
			gate["side_effects_disabled"] != true ||
			gate["host_root_modified"] != false ||
			gate["internal_details_exposed"] != false ||
			gate["gate_status"] != "notification-gate-modeled-delivery-disabled" {
			t.Fatalf("unsafe production receipt notification delivery item: %s", output.String())
		}
	}
	counts := payload["counts"].(map[string]any)
	if counts["total"] != float64(9) ||
		counts["passed"] != float64(9) ||
		counts["pending"] != float64(0) ||
		counts["blocked"] != float64(0) {
		t.Fatalf("unexpected production receipt notification delivery checks: %s", output.String())
	}
	if strings.Contains(output.String(), root) {
		t.Fatalf("production receipt notification delivery output must not expose project root path: %s", output.String())
	}
	for _, key := range []string{"system_service_started", "session_bus_claimed", "production_bus_claimed", "production_owner_enabled", "production_activation_ready", "write_methods_enabled", "runtime_writes_enabled", "receipt_writer_enabled", "receipt_persistence_enabled", "receipt_lookup_writes_enabled", "receipt_replay_enabled", "receipt_expiry_write_enabled", "receipt_revocation_write_enabled", "receipt_revocation_visibility_persisted", "notification_delivery_gate_persisted", "desktop_files_written", "mimeapps_written", "shell_configuration_written", "settings_persisted", "krunner_index_persisted", "task_manager_entry_active", "kwin_rule_applied", "live_tray_bridge_enabled", "tray_bridge_persisted", "notification_sent", "notification_delivery_enabled", "notification_action_enabled", "compatibility_center_persisted", "portal_request_created", "request_objects_created", "adapter_invocation_enabled", "backend_launch_enabled", "backend_process_started", "support_bundle_exported", "support_case_created", "snapshot_restore_executed", "state_cleanup_executed", "file_content_read", "file_paths_exposed", "network_required", "host_root_modified", "privileged_container_required", "state_root_path_exposed", "raw_command_exposed", "raw_executable_exposed", "backend_details_exposed"} {
		if payload[key] != false {
			t.Fatalf("unsafe production receipt notification delivery gate %s must remain false: %s", key, output.String())
		}
	}
}

func TestProductionReceiptNotificationDeliveryGateAuditPreviewCommandRejectsPositionalArgs(t *testing.T) {
	var output bytes.Buffer
	if err := run([]string{"production-receipt-notification-delivery-gate-audit-preview", "extra"}, &output); err == nil {
		t.Fatalf("production-receipt-notification-delivery-gate-audit-preview must reject positional arguments")
	}
}

func TestProductionReceiptNotificationActionSafetyAuditPreviewCommand(t *testing.T) {
	root := projectRootForRuntimeServiceBindingCommandTest(t)
	var output bytes.Buffer
	if err := run([]string{"production-receipt-notification-action-safety-audit-preview", "--root", root}, &output); err != nil {
		t.Fatalf("production-receipt-notification-action-safety-audit-preview returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["version"] != currentProjectVersion(t) ||
		payload["schema_version"] != "xnix.runtime.production_receipt_notification_action_safety_audit.v1" ||
		payload["request_type"] != "production-receipt-notification-action-safety-audit-preview" ||
		payload["audit_type"] != "receipt-notification-action-request-safety-audit" ||
		payload["audit_decision"] != "production-receipt-notification-action-safety-audit-ready-actions-disabled" ||
		payload["receipt_schema"] != "xnix.runtime.production_dbus_human_authorization_receipt.v1" ||
		payload["opaque_receipt_id"] != "production-dbus-human-authorization-receipt-id" {
		t.Fatalf("unexpected production receipt notification action payload: %s", output.String())
	}
	if payload["action_safety_audit_required"] != true ||
		payload["action_safety_audit_modeled"] != true ||
		payload["notification_delivery_gate_consumed"] != true ||
		payload["desktop_side_effect_review_consumed"] != true ||
		payload["operator_action_approval_required"] != true ||
		payload["operator_action_approval_present"] != false ||
		payload["notification_action_safety_ready"] != true ||
		payload["caller_state_root_required"] != false ||
		payload["receipt_present"] != false ||
		payload["receipt_accepted"] != false ||
		payload["authorization_accepted"] != false ||
		payload["production_readiness"] != false ||
		payload["production_ownership_ready"] != false {
		t.Fatalf("unexpected production receipt notification action decision: %s", output.String())
	}
	if payload["action_item_count"] != float64(5) ||
		payload["required_action_item_count"] != float64(5) ||
		payload["ready_action_item_count"] != float64(5) ||
		payload["missing_action_item_count"] != float64(0) ||
		payload["enabled_action_item_count"] != float64(0) ||
		payload["request_created_item_count"] != float64(0) ||
		payload["navigation_enabled_item_count"] != float64(0) ||
		payload["side_effect_item_count"] != float64(0) {
		t.Fatalf("unexpected production receipt notification action counts: %s", output.String())
	}
	items := payload["action_items"].([]any)
	if len(items) != 5 {
		t.Fatalf("unexpected production receipt notification action item list: %s", output.String())
	}
	for _, item := range items {
		action := item.(map[string]any)
		if action["evidence_present"] != true ||
			action["action_safety_modeled"] != true ||
			action["user_visible"] != true ||
			action["review_only"] != true ||
			action["runtime_owned"] != true ||
			action["go_runtime_backed"] != true ||
			action["kde_policy_owner"] != false ||
			action["operator_approval_required"] != true ||
			action["operator_approval_present"] != false ||
			action["action_enabled"] != false ||
			action["navigation_enabled"] != false ||
			action["portal_request_created"] != false ||
			action["request_objects_created"] != false ||
			action["notification_sent"] != false ||
			action["notification_delivery_enabled"] != false ||
			action["notification_action_enabled"] != false ||
			action["receipt_accepted"] != false ||
			action["authorization_accepted"] != false ||
			action["receipt_revocation_write_enabled"] != false ||
			action["receipt_expiry_write_enabled"] != false ||
			action["receipt_persistence_enabled"] != false ||
			action["receipt_lookup_writes_enabled"] != false ||
			action["compatibility_center_opened"] != false ||
			action["compatibility_center_persisted"] != false ||
			action["support_bundle_exported"] != false ||
			action["support_case_created"] != false ||
			action["production_readiness"] != false ||
			action["production_ownership_ready"] != false ||
			action["side_effects_disabled"] != true ||
			action["host_root_modified"] != false ||
			action["internal_details_exposed"] != false ||
			action["action_status"] != "notification-action-modeled-actions-disabled" {
			t.Fatalf("unsafe production receipt notification action item: %s", output.String())
		}
	}
	counts := payload["counts"].(map[string]any)
	if counts["total"] != float64(9) ||
		counts["passed"] != float64(9) ||
		counts["pending"] != float64(0) ||
		counts["blocked"] != float64(0) {
		t.Fatalf("unexpected production receipt notification action checks: %s", output.String())
	}
	if strings.Contains(output.String(), root) {
		t.Fatalf("production receipt notification action output must not expose project root path: %s", output.String())
	}
	for _, key := range []string{"system_service_started", "session_bus_claimed", "production_bus_claimed", "production_owner_enabled", "production_activation_ready", "write_methods_enabled", "runtime_writes_enabled", "receipt_writer_enabled", "receipt_persistence_enabled", "receipt_lookup_writes_enabled", "receipt_replay_enabled", "receipt_expiry_write_enabled", "receipt_revocation_write_enabled", "notification_sent", "notification_delivery_enabled", "notification_action_enabled", "notification_action_safety_persisted", "review_action_enabled", "renew_action_enabled", "open_compatibility_center_enabled", "dismiss_action_enabled", "support_info_action_enabled", "compatibility_center_opened", "compatibility_center_persisted", "portal_request_created", "request_objects_created", "desktop_files_written", "mimeapps_written", "shell_configuration_written", "settings_persisted", "krunner_index_persisted", "task_manager_entry_active", "kwin_rule_applied", "live_tray_bridge_enabled", "tray_bridge_persisted", "adapter_invocation_enabled", "backend_launch_enabled", "backend_process_started", "support_bundle_exported", "support_case_created", "snapshot_restore_executed", "state_cleanup_executed", "file_content_read", "file_paths_exposed", "network_required", "host_root_modified", "privileged_container_required", "state_root_path_exposed", "raw_command_exposed", "raw_executable_exposed", "backend_details_exposed"} {
		if payload[key] != false {
			t.Fatalf("unsafe production receipt notification action gate %s must remain false: %s", key, output.String())
		}
	}
}

func TestProductionReceiptNotificationActionSafetyAuditPreviewCommandRejectsPositionalArgs(t *testing.T) {
	var output bytes.Buffer
	if err := run([]string{"production-receipt-notification-action-safety-audit-preview", "extra"}, &output); err == nil {
		t.Fatalf("production-receipt-notification-action-safety-audit-preview must reject positional arguments")
	}
}

func TestProductionReceiptNotificationActionRequestObjectAuditPreviewCommand(t *testing.T) {
	root := projectRootForRuntimeServiceBindingCommandTest(t)
	var output bytes.Buffer
	if err := run([]string{"production-receipt-notification-action-request-object-audit-preview", "--root", root}, &output); err != nil {
		t.Fatalf("production-receipt-notification-action-request-object-audit-preview returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["version"] != currentProjectVersion(t) ||
		payload["schema_version"] != "xnix.runtime.production_receipt_notification_action_request_object_audit.v1" ||
		payload["request_type"] != "production-receipt-notification-action-request-object-audit-preview" ||
		payload["audit_type"] != "receipt-notification-action-runtime-request-object-audit" ||
		payload["audit_decision"] != "production-receipt-notification-action-request-object-audit-ready-requests-disabled" ||
		payload["receipt_schema"] != "xnix.runtime.production_dbus_human_authorization_receipt.v1" ||
		payload["opaque_receipt_id"] != "production-dbus-human-authorization-receipt-id" {
		t.Fatalf("unexpected production receipt notification action request-object payload: %s", output.String())
	}
	if payload["request_object_audit_required"] != true ||
		payload["request_object_audit_modeled"] != true ||
		payload["action_safety_audit_consumed"] != true ||
		payload["dispatch_request_boundary_consumed"] != true ||
		payload["operator_action_approval_required"] != true ||
		payload["operator_action_approval_present"] != false ||
		payload["request_object_boundary_ready"] != true ||
		payload["caller_state_root_required"] != false ||
		payload["receipt_present"] != false ||
		payload["receipt_accepted"] != false ||
		payload["authorization_accepted"] != false ||
		payload["production_readiness"] != false ||
		payload["production_ownership_ready"] != false {
		t.Fatalf("unexpected production receipt notification action request-object decision: %s", output.String())
	}
	if payload["request_object_count"] != float64(5) ||
		payload["required_request_object_count"] != float64(5) ||
		payload["ready_request_object_count"] != float64(5) ||
		payload["missing_request_object_count"] != float64(0) ||
		payload["created_request_object_count"] != float64(0) ||
		payload["dispatched_request_object_count"] != float64(0) ||
		payload["portal_request_created_count"] != float64(0) ||
		payload["navigation_requested_count"] != float64(0) ||
		payload["side_effect_request_object_count"] != float64(0) {
		t.Fatalf("unexpected production receipt notification action request-object counts: %s", output.String())
	}
	objects := payload["request_objects"].([]any)
	if len(objects) != 5 {
		t.Fatalf("unexpected production receipt notification action request-object list: %s", output.String())
	}
	for _, object := range objects {
		request := object.(map[string]any)
		if request["evidence_present"] != true ||
			request["request_object_modeled"] != true ||
			request["user_visible"] != true ||
			request["review_only"] != true ||
			request["runtime_owned"] != true ||
			request["go_runtime_backed"] != true ||
			request["kde_policy_owner"] != false ||
			request["operator_approval_required"] != true ||
			request["operator_approval_present"] != false ||
			request["request_object_created"] != false ||
			request["request_object_dispatched"] != false ||
			request["request_object_persisted"] != false ||
			request["portal_request_created"] != false ||
			request["navigation_requested"] != false ||
			request["notification_action_enabled"] != false ||
			request["action_enabled"] != false ||
			request["receipt_accepted"] != false ||
			request["authorization_accepted"] != false ||
			request["receipt_writer_enabled"] != false ||
			request["receipt_persistence_enabled"] != false ||
			request["compatibility_center_opened"] != false ||
			request["compatibility_center_persisted"] != false ||
			request["support_bundle_exported"] != false ||
			request["support_case_created"] != false ||
			request["production_readiness"] != false ||
			request["production_ownership_ready"] != false ||
			request["side_effects_disabled"] != true ||
			request["host_root_modified"] != false ||
			request["internal_details_exposed"] != false ||
			request["request_object_status"] != "request-object-modeled-creation-disabled" {
			t.Fatalf("unsafe production receipt notification action request-object item: %s", output.String())
		}
	}
	counts := payload["counts"].(map[string]any)
	if counts["total"] != float64(9) ||
		counts["passed"] != float64(9) ||
		counts["pending"] != float64(0) ||
		counts["blocked"] != float64(0) {
		t.Fatalf("unexpected production receipt notification action request-object checks: %s", output.String())
	}
	if strings.Contains(output.String(), root) {
		t.Fatalf("production receipt notification action request-object output must not expose project root path: %s", output.String())
	}
	for _, key := range []string{"system_service_started", "session_bus_claimed", "production_bus_claimed", "production_owner_enabled", "production_activation_ready", "write_methods_enabled", "runtime_writes_enabled", "request_object_creation_enabled", "request_object_dispatch_enabled", "request_object_persistence_enabled", "portal_request_created", "request_objects_created", "request_objects_dispatched", "notification_sent", "notification_delivery_enabled", "notification_action_enabled", "review_action_enabled", "renew_action_enabled", "open_compatibility_center_enabled", "dismiss_action_enabled", "support_info_action_enabled", "compatibility_center_opened", "compatibility_center_persisted", "receipt_writer_enabled", "receipt_persistence_enabled", "receipt_lookup_writes_enabled", "receipt_replay_enabled", "receipt_expiry_write_enabled", "receipt_revocation_write_enabled", "desktop_files_written", "mimeapps_written", "shell_configuration_written", "settings_persisted", "krunner_index_persisted", "task_manager_entry_active", "kwin_rule_applied", "live_tray_bridge_enabled", "tray_bridge_persisted", "adapter_invocation_enabled", "backend_launch_enabled", "backend_process_started", "support_bundle_exported", "support_case_created", "snapshot_restore_executed", "state_cleanup_executed", "file_content_read", "file_paths_exposed", "network_required", "host_root_modified", "privileged_container_required", "state_root_path_exposed", "raw_command_exposed", "raw_executable_exposed", "backend_details_exposed"} {
		if payload[key] != false {
			t.Fatalf("unsafe production receipt notification action request-object gate %s must remain false: %s", key, output.String())
		}
	}
}

func TestProductionReceiptNotificationActionRequestObjectAuditPreviewCommandRejectsPositionalArgs(t *testing.T) {
	var output bytes.Buffer
	if err := run([]string{"production-receipt-notification-action-request-object-audit-preview", "extra"}, &output); err == nil {
		t.Fatalf("production-receipt-notification-action-request-object-audit-preview must reject positional arguments")
	}
}

func TestProductionReceiptNotificationActionDispatchAuthorizationAuditPreviewCommand(t *testing.T) {
	root := projectRootForRuntimeServiceBindingCommandTest(t)
	var output bytes.Buffer
	if err := run([]string{"production-receipt-notification-action-dispatch-authorization-audit-preview", "--root", root}, &output); err != nil {
		t.Fatalf("production-receipt-notification-action-dispatch-authorization-audit-preview returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["version"] != currentProjectVersion(t) ||
		payload["schema_version"] != "xnix.runtime.production_receipt_notification_action_dispatch_authorization_audit.v1" ||
		payload["request_type"] != "production-receipt-notification-action-dispatch-authorization-audit-preview" ||
		payload["audit_type"] != "receipt-notification-action-request-dispatch-authorization-audit" ||
		payload["audit_decision"] != "production-receipt-notification-action-dispatch-authorization-audit-ready-dispatch-disabled" ||
		payload["receipt_schema"] != "xnix.runtime.production_dbus_human_authorization_receipt.v1" ||
		payload["opaque_receipt_id"] != "production-dbus-human-authorization-receipt-id" {
		t.Fatalf("unexpected production receipt notification action dispatch authorization payload: %s", output.String())
	}
	if payload["dispatch_authorization_audit_required"] != true ||
		payload["dispatch_authorization_audit_modeled"] != true ||
		payload["request_object_audit_consumed"] != true ||
		payload["receipt_authorization_boundary_consumed"] != true ||
		payload["operator_dispatch_approval_required"] != true ||
		payload["operator_dispatch_approval_present"] != false ||
		payload["dispatch_authorization_ready"] != true ||
		payload["dispatch_authorization_granted"] != false ||
		payload["caller_state_root_required"] != false ||
		payload["receipt_present"] != false ||
		payload["receipt_accepted"] != false ||
		payload["authorization_accepted"] != false ||
		payload["production_readiness"] != false ||
		payload["production_ownership_ready"] != false {
		t.Fatalf("unexpected production receipt notification action dispatch authorization decision: %s", output.String())
	}
	if payload["authorization_item_count"] != float64(5) ||
		payload["required_authorization_item_count"] != float64(5) ||
		payload["ready_authorization_item_count"] != float64(5) ||
		payload["missing_authorization_item_count"] != float64(0) ||
		payload["granted_authorization_item_count"] != float64(0) ||
		payload["dispatched_authorization_item_count"] != float64(0) ||
		payload["created_request_object_count"] != float64(0) ||
		payload["portal_request_created_count"] != float64(0) ||
		payload["side_effect_authorization_item_count"] != float64(0) {
		t.Fatalf("unexpected production receipt notification action dispatch authorization counts: %s", output.String())
	}
	items := payload["authorization_items"].([]any)
	if len(items) != 5 {
		t.Fatalf("unexpected production receipt notification action dispatch authorization item list: %s", output.String())
	}
	for _, item := range items {
		authorization := item.(map[string]any)
		if authorization["evidence_present"] != true ||
			authorization["dispatch_authorization_modeled"] != true ||
			authorization["user_visible"] != true ||
			authorization["review_only"] != true ||
			authorization["runtime_owned"] != true ||
			authorization["go_runtime_backed"] != true ||
			authorization["kde_policy_owner"] != false ||
			authorization["operator_approval_required"] != true ||
			authorization["operator_approval_present"] != false ||
			authorization["dispatch_authorization_granted"] != false ||
			authorization["request_object_created"] != false ||
			authorization["request_object_dispatched"] != false ||
			authorization["request_object_persisted"] != false ||
			authorization["portal_request_created"] != false ||
			authorization["navigation_requested"] != false ||
			authorization["notification_action_enabled"] != false ||
			authorization["action_enabled"] != false ||
			authorization["receipt_accepted"] != false ||
			authorization["authorization_accepted"] != false ||
			authorization["receipt_writer_enabled"] != false ||
			authorization["receipt_persistence_enabled"] != false ||
			authorization["compatibility_center_opened"] != false ||
			authorization["compatibility_center_persisted"] != false ||
			authorization["support_bundle_exported"] != false ||
			authorization["support_case_created"] != false ||
			authorization["production_readiness"] != false ||
			authorization["production_ownership_ready"] != false ||
			authorization["side_effects_disabled"] != true ||
			authorization["host_root_modified"] != false ||
			authorization["internal_details_exposed"] != false ||
			authorization["dispatch_authorization_status"] != "dispatch-authorization-modeled-dispatch-disabled" {
			t.Fatalf("unsafe production receipt notification action dispatch authorization item: %s", output.String())
		}
	}
	counts := payload["counts"].(map[string]any)
	if counts["total"] != float64(9) ||
		counts["passed"] != float64(9) ||
		counts["pending"] != float64(0) ||
		counts["blocked"] != float64(0) {
		t.Fatalf("unexpected production receipt notification action dispatch authorization checks: %s", output.String())
	}
	if strings.Contains(output.String(), root) {
		t.Fatalf("production receipt notification action dispatch authorization output must not expose project root path: %s", output.String())
	}
	for _, key := range []string{"system_service_started", "session_bus_claimed", "production_bus_claimed", "production_owner_enabled", "production_activation_ready", "write_methods_enabled", "runtime_writes_enabled", "request_object_creation_enabled", "request_object_dispatch_enabled", "request_object_persistence_enabled", "dispatch_authorization_persisted", "portal_request_created", "request_objects_created", "request_objects_dispatched", "notification_sent", "notification_delivery_enabled", "notification_action_enabled", "review_action_enabled", "renew_action_enabled", "open_compatibility_center_enabled", "dismiss_action_enabled", "support_info_action_enabled", "compatibility_center_opened", "compatibility_center_persisted", "receipt_writer_enabled", "receipt_persistence_enabled", "receipt_lookup_writes_enabled", "receipt_replay_enabled", "receipt_expiry_write_enabled", "receipt_revocation_write_enabled", "desktop_files_written", "mimeapps_written", "shell_configuration_written", "settings_persisted", "krunner_index_persisted", "task_manager_entry_active", "kwin_rule_applied", "live_tray_bridge_enabled", "tray_bridge_persisted", "adapter_invocation_enabled", "backend_launch_enabled", "backend_process_started", "support_bundle_exported", "support_case_created", "snapshot_restore_executed", "state_cleanup_executed", "file_content_read", "file_paths_exposed", "network_required", "host_root_modified", "privileged_container_required", "state_root_path_exposed", "raw_command_exposed", "raw_executable_exposed", "backend_details_exposed"} {
		if payload[key] != false {
			t.Fatalf("unsafe production receipt notification action dispatch authorization gate %s must remain false: %s", key, output.String())
		}
	}
}

func TestProductionReceiptNotificationActionDispatchAuthorizationAuditPreviewCommandRejectsPositionalArgs(t *testing.T) {
	var output bytes.Buffer
	if err := run([]string{"production-receipt-notification-action-dispatch-authorization-audit-preview", "extra"}, &output); err == nil {
		t.Fatalf("production-receipt-notification-action-dispatch-authorization-audit-preview must reject positional arguments")
	}
}

func TestProductionReceiptNotificationActionDispatchDryRunAuditPreviewCommand(t *testing.T) {
	root := projectRootForRuntimeServiceBindingCommandTest(t)
	var output bytes.Buffer
	if err := run([]string{"production-receipt-notification-action-dispatch-dry-run-audit-preview", "--root", root}, &output); err != nil {
		t.Fatalf("production-receipt-notification-action-dispatch-dry-run-audit-preview returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["version"] != currentProjectVersion(t) ||
		payload["schema_version"] != "xnix.runtime.production_receipt_notification_action_dispatch_dry_run_audit.v1" ||
		payload["request_type"] != "production-receipt-notification-action-dispatch-dry-run-audit-preview" ||
		payload["audit_type"] != "receipt-notification-action-request-dispatch-dry-run-audit" ||
		payload["audit_decision"] != "production-receipt-notification-action-dispatch-dry-run-audit-ready-dry-run-only" ||
		payload["receipt_schema"] != "xnix.runtime.production_dbus_human_authorization_receipt.v1" ||
		payload["opaque_receipt_id"] != "production-dbus-human-authorization-receipt-id" {
		t.Fatalf("unexpected production receipt notification action dispatch dry-run payload: %s", output.String())
	}
	if payload["dispatch_dry_run_audit_required"] != true ||
		payload["dispatch_dry_run_audit_modeled"] != true ||
		payload["dispatch_authorization_audit_consumed"] != true ||
		payload["dispatch_dry_run_guidance_consumed"] != true ||
		payload["dispatch_dry_run_plan_ready"] != true ||
		payload["dispatch_dry_run_executed"] != false ||
		payload["dispatch_authorization_granted"] != false ||
		payload["operator_dispatch_approval_required"] != true ||
		payload["operator_dispatch_approval_present"] != false ||
		payload["caller_state_root_required"] != false ||
		payload["receipt_present"] != false ||
		payload["receipt_accepted"] != false ||
		payload["authorization_accepted"] != false ||
		payload["production_readiness"] != false ||
		payload["production_ownership_ready"] != false {
		t.Fatalf("unexpected production receipt notification action dispatch dry-run decision: %s", output.String())
	}
	if payload["dry_run_item_count"] != float64(5) ||
		payload["required_dry_run_item_count"] != float64(5) ||
		payload["ready_dry_run_item_count"] != float64(5) ||
		payload["missing_dry_run_item_count"] != float64(0) ||
		payload["executed_dry_run_item_count"] != float64(0) ||
		payload["dispatched_dry_run_item_count"] != float64(0) ||
		payload["created_request_object_count"] != float64(0) ||
		payload["portal_request_created_count"] != float64(0) ||
		payload["side_effect_dry_run_item_count"] != float64(0) {
		t.Fatalf("unexpected production receipt notification action dispatch dry-run counts: %s", output.String())
	}
	items := payload["dry_run_items"].([]any)
	if len(items) != 5 {
		t.Fatalf("unexpected production receipt notification action dispatch dry-run item list: %s", output.String())
	}
	for _, item := range items {
		dryRun := item.(map[string]any)
		if dryRun["evidence_present"] != true ||
			dryRun["dispatch_dry_run_modeled"] != true ||
			dryRun["user_visible"] != true ||
			dryRun["review_only"] != true ||
			dryRun["runtime_owned"] != true ||
			dryRun["go_runtime_backed"] != true ||
			dryRun["kde_policy_owner"] != false ||
			dryRun["operator_approval_required"] != true ||
			dryRun["operator_approval_present"] != false ||
			dryRun["dispatch_authorization_granted"] != false ||
			dryRun["dispatch_dry_run_executed"] != false ||
			dryRun["request_object_created"] != false ||
			dryRun["request_object_dispatched"] != false ||
			dryRun["request_object_persisted"] != false ||
			dryRun["portal_request_created"] != false ||
			dryRun["navigation_requested"] != false ||
			dryRun["notification_action_enabled"] != false ||
			dryRun["action_enabled"] != false ||
			dryRun["receipt_accepted"] != false ||
			dryRun["authorization_accepted"] != false ||
			dryRun["receipt_writer_enabled"] != false ||
			dryRun["receipt_persistence_enabled"] != false ||
			dryRun["compatibility_center_opened"] != false ||
			dryRun["compatibility_center_persisted"] != false ||
			dryRun["support_bundle_exported"] != false ||
			dryRun["support_case_created"] != false ||
			dryRun["production_readiness"] != false ||
			dryRun["production_ownership_ready"] != false ||
			dryRun["side_effects_disabled"] != true ||
			dryRun["host_root_modified"] != false ||
			dryRun["internal_details_exposed"] != false ||
			dryRun["dry_run_status"] != "dispatch-dry-run-modeled-execution-disabled" {
			t.Fatalf("unsafe production receipt notification action dispatch dry-run item: %s", output.String())
		}
	}
	counts := payload["counts"].(map[string]any)
	if counts["total"] != float64(9) ||
		counts["passed"] != float64(9) ||
		counts["pending"] != float64(0) ||
		counts["blocked"] != float64(0) {
		t.Fatalf("unexpected production receipt notification action dispatch dry-run checks: %s", output.String())
	}
	if strings.Contains(output.String(), root) {
		t.Fatalf("production receipt notification action dispatch dry-run output must not expose project root path: %s", output.String())
	}
	for _, key := range []string{"system_service_started", "session_bus_claimed", "production_bus_claimed", "production_owner_enabled", "production_activation_ready", "write_methods_enabled", "runtime_writes_enabled", "request_object_creation_enabled", "request_object_dispatch_enabled", "request_object_persistence_enabled", "dispatch_authorization_persisted", "dispatch_dry_run_execution_enabled", "dry_run_result_persisted", "portal_request_created", "request_objects_created", "request_objects_dispatched", "notification_sent", "notification_delivery_enabled", "notification_action_enabled", "review_action_enabled", "renew_action_enabled", "open_compatibility_center_enabled", "dismiss_action_enabled", "support_info_action_enabled", "compatibility_center_opened", "compatibility_center_persisted", "receipt_writer_enabled", "receipt_persistence_enabled", "receipt_lookup_writes_enabled", "receipt_replay_enabled", "receipt_expiry_write_enabled", "receipt_revocation_write_enabled", "desktop_files_written", "mimeapps_written", "shell_configuration_written", "settings_persisted", "krunner_index_persisted", "task_manager_entry_active", "kwin_rule_applied", "live_tray_bridge_enabled", "tray_bridge_persisted", "adapter_invocation_enabled", "backend_launch_enabled", "backend_process_started", "support_bundle_exported", "support_case_created", "snapshot_restore_executed", "state_cleanup_executed", "file_content_read", "file_paths_exposed", "network_required", "host_root_modified", "privileged_container_required", "state_root_path_exposed", "raw_command_exposed", "raw_executable_exposed", "backend_details_exposed"} {
		if payload[key] != false {
			t.Fatalf("unsafe production receipt notification action dispatch dry-run gate %s must remain false: %s", key, output.String())
		}
	}
}

func TestProductionReceiptNotificationActionDispatchDryRunAuditPreviewCommandRejectsPositionalArgs(t *testing.T) {
	var output bytes.Buffer
	if err := run([]string{"production-receipt-notification-action-dispatch-dry-run-audit-preview", "extra"}, &output); err == nil {
		t.Fatalf("production-receipt-notification-action-dispatch-dry-run-audit-preview must reject positional arguments")
	}
}

func TestProductionReceiptNotificationActionDryRunResultVisibilityAuditPreviewCommand(t *testing.T) {
	root := projectRootForRuntimeServiceBindingCommandTest(t)
	var output bytes.Buffer
	if err := run([]string{"production-receipt-notification-action-dry-run-result-visibility-audit-preview", "--root", root}, &output); err != nil {
		t.Fatalf("production-receipt-notification-action-dry-run-result-visibility-audit-preview returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["version"] != currentProjectVersion(t) ||
		payload["schema_version"] != "xnix.runtime.production_receipt_notification_action_dry_run_result_visibility_audit.v1" ||
		payload["request_type"] != "production-receipt-notification-action-dry-run-result-visibility-audit-preview" ||
		payload["audit_type"] != "receipt-notification-action-dry-run-result-visibility-audit" ||
		payload["audit_decision"] != "production-receipt-notification-action-dry-run-result-visibility-audit-ready-visibility-only" ||
		payload["receipt_schema"] != "xnix.runtime.production_dbus_human_authorization_receipt.v1" ||
		payload["opaque_receipt_id"] != "production-dbus-human-authorization-receipt-id" {
		t.Fatalf("unexpected production receipt notification action dry-run result visibility payload: %s", output.String())
	}
	if payload["result_visibility_audit_required"] != true ||
		payload["result_visibility_audit_modeled"] != true ||
		payload["dispatch_dry_run_audit_consumed"] != true ||
		payload["result_visibility_guidance_consumed"] != true ||
		payload["result_visibility_plan_ready"] != true ||
		payload["redacted_kde_visibility_modeled"] != true ||
		payload["runtime_diagnostics_visibility_modeled"] != true ||
		payload["dispatch_dry_run_executed"] != false ||
		payload["dry_run_result_persisted"] != false ||
		payload["dispatch_authorization_granted"] != false ||
		payload["operator_dispatch_approval_required"] != true ||
		payload["operator_dispatch_approval_present"] != false ||
		payload["caller_state_root_required"] != false ||
		payload["receipt_present"] != false ||
		payload["receipt_accepted"] != false ||
		payload["authorization_accepted"] != false ||
		payload["production_readiness"] != false ||
		payload["production_ownership_ready"] != false {
		t.Fatalf("unexpected production receipt notification action dry-run result visibility decision: %s", output.String())
	}
	if payload["visibility_item_count"] != float64(5) ||
		payload["required_visibility_item_count"] != float64(5) ||
		payload["ready_visibility_item_count"] != float64(5) ||
		payload["missing_visibility_item_count"] != float64(0) ||
		payload["persisted_visibility_item_count"] != float64(0) ||
		payload["executed_dry_run_item_count"] != float64(0) ||
		payload["dispatched_dry_run_item_count"] != float64(0) ||
		payload["created_request_object_count"] != float64(0) ||
		payload["portal_request_created_count"] != float64(0) ||
		payload["side_effect_visibility_item_count"] != float64(0) {
		t.Fatalf("unexpected production receipt notification action dry-run result visibility counts: %s", output.String())
	}
	items := payload["visibility_items"].([]any)
	if len(items) != 5 {
		t.Fatalf("unexpected production receipt notification action dry-run result visibility item list: %s", output.String())
	}
	for _, item := range items {
		visibility := item.(map[string]any)
		if visibility["evidence_present"] != true ||
			visibility["result_visibility_modeled"] != true ||
			visibility["redacted_for_kde"] != true ||
			visibility["runtime_diagnostics_modeled"] != true ||
			visibility["user_visible"] != true ||
			visibility["review_only"] != true ||
			visibility["runtime_owned"] != true ||
			visibility["go_runtime_backed"] != true ||
			visibility["kde_policy_owner"] != false ||
			visibility["operator_approval_required"] != true ||
			visibility["operator_approval_present"] != false ||
			visibility["dispatch_authorization_granted"] != false ||
			visibility["dispatch_dry_run_executed"] != false ||
			visibility["dry_run_result_persisted"] != false ||
			visibility["result_visibility_persisted"] != false ||
			visibility["request_object_created"] != false ||
			visibility["request_object_dispatched"] != false ||
			visibility["request_object_persisted"] != false ||
			visibility["portal_request_created"] != false ||
			visibility["navigation_requested"] != false ||
			visibility["notification_action_enabled"] != false ||
			visibility["action_enabled"] != false ||
			visibility["receipt_accepted"] != false ||
			visibility["authorization_accepted"] != false ||
			visibility["receipt_writer_enabled"] != false ||
			visibility["receipt_persistence_enabled"] != false ||
			visibility["compatibility_center_opened"] != false ||
			visibility["compatibility_center_persisted"] != false ||
			visibility["runtime_diagnostics_persisted"] != false ||
			visibility["support_bundle_exported"] != false ||
			visibility["support_case_created"] != false ||
			visibility["production_readiness"] != false ||
			visibility["production_ownership_ready"] != false ||
			visibility["side_effects_disabled"] != true ||
			visibility["host_root_modified"] != false ||
			visibility["internal_details_exposed"] != false ||
			visibility["visibility_status"] != "dry-run-result-visibility-modeled-persistence-disabled" {
			t.Fatalf("unsafe production receipt notification action dry-run result visibility item: %s", output.String())
		}
	}
	counts := payload["counts"].(map[string]any)
	if counts["total"] != float64(9) ||
		counts["passed"] != float64(9) ||
		counts["pending"] != float64(0) ||
		counts["blocked"] != float64(0) {
		t.Fatalf("unexpected production receipt notification action dry-run result visibility checks: %s", output.String())
	}
	if strings.Contains(output.String(), root) {
		t.Fatalf("production receipt notification action dry-run result visibility output must not expose project root path: %s", output.String())
	}
	for _, key := range []string{"system_service_started", "session_bus_claimed", "production_bus_claimed", "production_owner_enabled", "production_activation_ready", "write_methods_enabled", "runtime_writes_enabled", "request_object_creation_enabled", "request_object_dispatch_enabled", "request_object_persistence_enabled", "dispatch_authorization_persisted", "dispatch_dry_run_execution_enabled", "dry_run_result_persistence_enabled", "result_visibility_persistence_enabled", "portal_request_created", "request_objects_created", "request_objects_dispatched", "notification_sent", "notification_delivery_enabled", "notification_action_enabled", "review_action_enabled", "renew_action_enabled", "open_compatibility_center_enabled", "dismiss_action_enabled", "support_info_action_enabled", "compatibility_center_opened", "compatibility_center_persisted", "runtime_diagnostics_persisted", "receipt_writer_enabled", "receipt_persistence_enabled", "receipt_lookup_writes_enabled", "receipt_replay_enabled", "receipt_expiry_write_enabled", "receipt_revocation_write_enabled", "desktop_files_written", "mimeapps_written", "shell_configuration_written", "settings_persisted", "krunner_index_persisted", "task_manager_entry_active", "kwin_rule_applied", "live_tray_bridge_enabled", "tray_bridge_persisted", "adapter_invocation_enabled", "backend_launch_enabled", "backend_process_started", "support_bundle_exported", "support_case_created", "snapshot_restore_executed", "state_cleanup_executed", "file_content_read", "file_paths_exposed", "network_required", "host_root_modified", "privileged_container_required", "state_root_path_exposed", "raw_command_exposed", "raw_executable_exposed", "backend_details_exposed"} {
		if payload[key] != false {
			t.Fatalf("unsafe production receipt notification action dry-run result visibility gate %s must remain false: %s", key, output.String())
		}
	}
}

func TestProductionReceiptNotificationActionDryRunResultVisibilityAuditPreviewCommandRejectsPositionalArgs(t *testing.T) {
	var output bytes.Buffer
	if err := run([]string{"production-receipt-notification-action-dry-run-result-visibility-audit-preview", "extra"}, &output); err == nil {
		t.Fatalf("production-receipt-notification-action-dry-run-result-visibility-audit-preview must reject positional arguments")
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

func TestProductionDesktopSideEffectReviewPreviewCommand(t *testing.T) {
	root := projectRootForRuntimeServiceBindingCommandTest(t)
	var output bytes.Buffer
	if err := run([]string{"production-desktop-side-effect-review-preview", "--root", root}, &output); err != nil {
		t.Fatalf("production-desktop-side-effect-review-preview returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["version"] != currentProjectVersion(t) ||
		payload["schema_version"] != "xnix.runtime.production_desktop_side_effect_review.v1" ||
		payload["request_type"] != "production-desktop-side-effect-review-preview" ||
		payload["review_type"] != "kde-production-desktop-side-effect-review" ||
		payload["review_decision"] != "production-desktop-side-effect-review-ready-side-effects-disabled" {
		t.Fatalf("unexpected production desktop side-effect review command payload: %s", output.String())
	}
	if payload["surface_count"] != float64(7) ||
		payload["required_surface_count"] != float64(7) ||
		payload["reviewed_surface_count"] != float64(7) ||
		payload["active_surface_count"] != float64(0) ||
		payload["side_effect_surface_count"] != float64(0) ||
		payload["production_ownership_ready"] != false ||
		payload["official_desktop_only"] != true ||
		payload["kde_policy_owner"] != false {
		t.Fatalf("unexpected production desktop side-effect review counts: %s", output.String())
	}
	surfaces := payload["surfaces"].([]any)
	if len(surfaces) != 7 {
		t.Fatalf("unexpected production desktop side-effect surface count: %s", output.String())
	}
	for _, item := range surfaces {
		surface := item.(map[string]any)
		if surface["required_before_production"] != true ||
			surface["evidence_present"] != true ||
			surface["runtime_owned"] != true ||
			surface["go_runtime_backed"] != true ||
			surface["kde_policy_owner"] != false ||
			surface["review_only"] != true ||
			surface["active"] != false ||
			surface["side_effects_enabled"] != false ||
			surface["desktop_files_written"] != false ||
			surface["mimeapps_written"] != false ||
			surface["shell_configuration_written"] != false ||
			surface["settings_persisted"] != false ||
			surface["krunner_index_persisted"] != false ||
			surface["task_manager_entry_active"] != false ||
			surface["kwin_rule_applied"] != false ||
			surface["live_tray_bridge_enabled"] != false ||
			surface["notification_sent"] != false ||
			surface["notification_delivery_enabled"] != false ||
			surface["compatibility_center_persisted"] != false ||
			surface["portal_request_created"] != false ||
			surface["request_objects_created"] != false ||
			surface["backend_launch_enabled"] != false ||
			surface["host_root_modified"] != false ||
			surface["backend_details_exposed"] != false ||
			surface["review_status"] != "reviewed" {
			t.Fatalf("unsafe desktop side-effect surface: %s", output.String())
		}
	}
	counts := payload["counts"].(map[string]any)
	if counts["total"] != float64(8) ||
		counts["passed"] != float64(8) ||
		counts["pending"] != float64(0) ||
		counts["blocked"] != float64(0) {
		t.Fatalf("unexpected production desktop side-effect review checks: %s", output.String())
	}
	if strings.Contains(output.String(), root) {
		t.Fatalf("production desktop side-effect review output must not expose project root path: %s", output.String())
	}
	for _, key := range []string{"plasma_fork_required", "plasma_source_modified", "system_service_started", "session_bus_claimed", "production_bus_claimed", "production_owner_enabled", "production_activation_ready", "write_methods_enabled", "runtime_writes_enabled", "desktop_files_written", "mimeapps_written", "shell_configuration_written", "settings_persisted", "krunner_index_persisted", "task_manager_entry_active", "kwin_rule_applied", "live_tray_bridge_enabled", "tray_bridge_persisted", "notification_sent", "notification_delivery_enabled", "compatibility_center_persisted", "portal_request_created", "request_objects_created", "adapter_invocation_enabled", "backend_launch_enabled", "backend_process_started", "file_content_read", "file_paths_exposed", "network_required", "host_root_modified", "privileged_container_required", "state_root_path_exposed", "raw_command_exposed", "raw_executable_exposed", "backend_details_exposed"} {
		if payload[key] != false {
			t.Fatalf("unsafe desktop side-effect gate %s must remain false: %s", key, output.String())
		}
	}
}

func TestProductionDesktopSideEffectReviewPreviewCommandRejectsPositionalArgs(t *testing.T) {
	var output bytes.Buffer
	if err := run([]string{"production-desktop-side-effect-review-preview", "extra"}, &output); err == nil {
		t.Fatalf("production-desktop-side-effect-review-preview must reject positional arguments")
	}
}
