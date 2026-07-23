package main

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"
)

func TestProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptStorageRecordWriterCallAuthorizationReceiptAcceptedReceiptGateCallAuthorizationReceiptAcceptedReceiptGateCallReceiptConsumerEnablementReceiptStorageRecordWriterAuthorizationReceiptPreviewCommand(t *testing.T) {
	root := projectRootForRuntimeServiceBindingCommandTest(t)
	var output bytes.Buffer
	command := "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-receipt-accepted-receipt-gate-call-authorization-receipt-accepted-receipt-gate-call-receipt-consumer-enablement-receipt-storage-record-writer-authorization-receipt-preview"
	if err := run([]string{command, "--root", root}, &output); err != nil {
		t.Fatalf("%s returned error: %v", command, err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["version"] != currentProjectVersion(t) ||
		payload["schema_version"] != "xnix.runtime.production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_kde_safe_redacted_status_closed_record_writer_storage_root_authorization_receipt_dry_run_lookup_result_consumer_projection_route_enablement_lookup_route_dispatch_dry_run_result_notification_action_request_object_dispatch_dry_run_result_persistence_receipt_consumer_enablement_receipt_storage_record_writer_call_authorization_receipt_accepted_receipt_gate_call_authorization_receipt_accepted_receipt_gate_call_receipt_consumer_enablement_receipt_storage_record_writer_authorization_receipt.v1" ||
		payload["request_type"] != command ||
		payload["preview_decision"] != "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-receipt-accepted-receipt-gate-call-authorization-receipt-accepted-receipt-gate-call-receipt-consumer-enablement-receipt-storage-record-writer-authorization-receipt-ready-receipt-disabled" {
		t.Fatalf("unexpected accepted receipt gate call receipt consumer enablement receipt storage record writer authorization receipt payload: %s", output.String())
	}
	for _, key := range []string{
		"current_mainline_consumed",
		"route_enablement_lookup_route_dispatch_dry_run_result_notification_action_request_object_dispatch_dry_run_result_persistence_receipt_consumer_enablement_receipt_storage_record_writer_call_authorization_receipt_accepted_receipt_gate_call_authorization_receipt_accepted_receipt_gate_call_receipt_consumer_enablement_receipt_storage_record_writer_authorization_consumed",
		"route_enablement_lookup_route_dispatch_dry_run_result_notification_action_request_object_dispatch_dry_run_result_persistence_receipt_consumer_enablement_receipt_storage_record_writer_call_authorization_receipt_accepted_receipt_gate_call_authorization_receipt_accepted_receipt_gate_call_receipt_consumer_enablement_receipt_storage_record_writer_authorization_receipt_storage_record_writer_authorization_evidence_consumed",
		"lookup_route_dispatch_dry_run_result_notification_action_request_object_dispatch_dry_run_result_persistence_receipt_consumer_enablement_receipt_storage_record_writer_call_authorization_receipt_accepted_receipt_gate_call_authorization_receipt_accepted_receipt_gate_call_receipt_consumer_enablement_receipt_storage_record_writer_authorization_receipt_accepted_receipt_gate_call_persistence_authorization_evidence_ready",
		"result_persistence_receipt_consumer_enablement_receipt_storage_record_writer_call_authorization_receipt_accepted_receipt_gate_call_authorization_receipt_accepted_receipt_gate_call_receipt_consumer_enablement_receipt_storage_record_writer_authorization_receipt_accepted_receipt_gate_call_persistence_authorization_evidence_ready",
		"lookup_route_dispatch_dry_run_result_notification_action_request_object_dispatch_dry_run_result_persistence_receipt_consumer_enablement_receipt_storage_record_writer_call_authorization_receipt_accepted_receipt_gate_call_authorization_receipt_accepted_receipt_gate_call_receipt_consumer_enablement_receipt_storage_record_writer_authorization_ready",
		"result_persistence_receipt_consumer_enablement_receipt_storage_record_writer_call_authorization_receipt_accepted_receipt_gate_call_authorization_receipt_accepted_receipt_gate_call_receipt_consumer_enablement_receipt_storage_record_writer_authorization_ready",
		"lookup_route_dispatch_dry_run_result_notification_action_request_object_dispatch_dry_run_result_persistence_receipt_consumer_enablement_receipt_storage_record_writer_call_authorization_receipt_accepted_receipt_gate_call_authorization_receipt_accepted_receipt_gate_call_receipt_consumer_enablement_receipt_storage_record_writer_authorization_receipt_ready",
		"receipt_storage_record_writer_authorization_receipt_boundary_modeled",
		"kde_safe_redacted_result_only",
		"raw_result_hidden",
	} {
		if payload[key] != true {
			t.Fatalf("expected %s to be true in accepted receipt gate call receipt consumer enablement receipt storage record writer authorization receipt payload: %s", key, output.String())
		}
	}
	if payload["notification_action_request_dispatch_dry_run_result_persistence_receipt_consumer_enablement_receipt_storage_record_writer_call_authorization_receipt_accepted_receipt_gate_call_authorization_receipt_accepted_receipt_gate_call_receipt_consumer_enablement_receipt_storage_record_writer_authorization_receipt_item_count"] != float64(4) ||
		payload["ready_notification_action_request_dispatch_dry_run_result_persistence_receipt_consumer_enablement_receipt_storage_record_writer_call_authorization_receipt_accepted_receipt_gate_call_authorization_receipt_accepted_receipt_gate_call_receipt_consumer_enablement_receipt_storage_record_writer_authorization_receipt_item_count"] != float64(4) ||
		payload["missing_notification_action_request_dispatch_dry_run_result_persistence_receipt_consumer_enablement_receipt_storage_record_writer_call_authorization_receipt_accepted_receipt_gate_call_authorization_receipt_accepted_receipt_gate_call_receipt_consumer_enablement_receipt_storage_record_writer_authorization_receipt_item_count"] != float64(0) ||
		payload["persistence_receipt_consumer_enablement_receipt_storage_record_writer_call_authorization_receipt_accepted_receipt_gate_call_authorization_receipt_accepted_receipt_gate_call_receipt_consumer_enablement_receipt_storage_record_writer_authorization_consumed_item_count"] != float64(4) ||
		payload["persistence_receipt_consumer_enablement_receipt_storage_record_writer_call_authorization_receipt_accepted_receipt_gate_call_authorization_receipt_accepted_receipt_gate_call_receipt_consumer_enablement_receipt_storage_record_writer_authorization_receipt_storage_record_writer_authorization_evidence_consumed_item_count"] != float64(4) ||
		payload["callable_result_persistence_receipt_consumer_enablement_receipt_storage_record_writer_authorization_receipt_item_count"] != float64(0) ||
		payload["granted_result_persistence_receipt_consumer_enablement_receipt_storage_record_writer_authorization_receipt_item_count"] != float64(0) ||
		payload["persisted_result_persistence_receipt_consumer_enablement_receipt_storage_record_writer_authorization_receipt_item_count"] != float64(0) ||
		payload["written_result_persistence_receipt_consumer_enablement_receipt_storage_record_item_count"] != float64(0) ||
		payload["side_effect_notification_action_request_dispatch_dry_run_result_persistence_receipt_consumer_enablement_receipt_storage_record_writer_call_authorization_receipt_accepted_receipt_gate_call_authorization_receipt_accepted_receipt_gate_call_receipt_consumer_enablement_receipt_storage_record_writer_authorization_receipt_item_count"] != float64(0) {
		t.Fatalf("unexpected accepted receipt gate call receipt consumer enablement receipt storage record writer authorization receipt counts: %s", output.String())
	}
	if !strings.Contains(output.String(), "storage-record-writer-call-authorization-receipt-accepted-receipt-gate-call-authorization-receipt-accepted-receipt-gate-call-receipt-consumer-enablement-receipt-storage-record-writer-authorization-receipt-install-failure") ||
		!strings.Contains(output.String(), "storage-record-writer-authorization-consumed") ||
		!strings.Contains(output.String(), "storage-record-writer-authorization-receipt-storage-record-writer-authorization-evidence-consumed") ||
		!strings.Contains(output.String(), "storage-record-writer-authorization-receipt-required-but-not-present") ||
		!strings.Contains(output.String(), "storage-record-writer-authorization-receipt-writes-visibility-persistence-and-dispatch-disabled") {
		t.Fatalf("accepted receipt gate call receipt consumer enablement receipt storage record writer authorization receipt payload must include item and check evidence: %s", output.String())
	}
	for _, key := range []string{"receipt_consumer_enablement_receipt_storage_record_writer_authorization_receipt_callable", "receipt_consumer_enablement_receipt_storage_record_writer_authorization_receipt_enabled", "receipt_consumer_enablement_receipt_storage_record_writer_authorization_receipt_granted", "receipt_consumer_enablement_receipt_storage_record_writer_authorization_receipt_present", "receipt_consumer_enablement_receipt_storage_record_writer_authorization_receipt_written", "receipt_consumer_enablement_receipt_storage_record_writer_authorization_receipt_persisted", "receipt_consumer_enablement_receipt_storage_record_writer_authorization_callable", "receipt_consumer_enablement_receipt_storage_record_writer_authorized", "receipt_consumer_enablement_receipt_storage_record_writer_authorization_granted", "receipt_consumer_enablement_receipt_storage_record_written", "receipt_consumer_enablement_receipt_storage_persistence_gate_passed", "receipt_consumer_enablement_receipt_storage_persisted", "receipt_consumer_enablement_receipt_persisted", "receipt_consumer_enablement_receipt_present", "receipt_consumer_enablement_receipt_write_enabled", "receipt_consumer_enablement_enabled", "consumer_enabled", "kde_consumer_enabled", "runtime_consumer_enabled", "receipt_consumption_enabled", "receipt_consumed", "receipt_write_enabled", "dry_run_result_persistence_enabled", "request_object_dispatch_enabled", "notification_action_enabled", "notification_sent", "user_visible", "storage_write_enabled", "production_ownership_ready", "production_bus_claimed", "write_methods_enabled", "runtime_writes_enabled", "backend_launch_enabled", "state_root_path_exposed", "file_paths_exposed", "file_content_read", "host_root_modified"} {
		if payload[key] != false {
			t.Fatalf("expected %s to remain false in accepted receipt gate call receipt consumer enablement receipt storage record writer authorization receipt payload: %s", key, output.String())
		}
	}
}

func TestProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptStorageRecordWriterCallAuthorizationReceiptAcceptedReceiptGateCallAuthorizationReceiptAcceptedReceiptGateCallReceiptConsumerEnablementReceiptStorageRecordWriterAuthorizationReceiptPreviewCommandRejectsPositionalArgs(t *testing.T) {
	var output bytes.Buffer
	if err := run([]string{"production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-receipt-accepted-receipt-gate-call-authorization-receipt-accepted-receipt-gate-call-receipt-consumer-enablement-receipt-storage-record-writer-authorization-receipt-preview", "extra"}, &output); err == nil {
		t.Fatalf("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-receipt-accepted-receipt-gate-call-authorization-receipt-accepted-receipt-gate-call-receipt-consumer-enablement-receipt-storage-record-writer-authorization-receipt-preview must reject positional arguments")
	}
}

func TestProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptStorageRecordWriterCallAuthorizationReceiptAcceptedReceiptGateCallAuthorizationReceiptAcceptedReceiptGateCallReceiptConsumerEnablementReceiptStorageRecordWriterAuthorizationReceiptAcceptancePreviewCommand(t *testing.T) {
	root := projectRootForRuntimeServiceBindingCommandTest(t)
	var output bytes.Buffer
	command := "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-receipt-accepted-receipt-gate-call-authorization-receipt-accepted-receipt-gate-call-receipt-consumer-enablement-receipt-storage-record-writer-authorization-receipt-acceptance-preview"
	if err := run([]string{command, "--root", root}, &output); err != nil {
		t.Fatalf("%s returned error: %v", command, err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["version"] != currentProjectVersion(t) ||
		payload["schema_version"] != "xnix.runtime.production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_kde_safe_redacted_status_closed_record_writer_storage_root_authorization_receipt_dry_run_lookup_result_consumer_projection_route_enablement_lookup_route_dispatch_dry_run_result_notification_action_request_object_dispatch_dry_run_result_persistence_receipt_consumer_enablement_receipt_storage_record_writer_call_authorization_receipt_accepted_receipt_gate_call_authorization_receipt_accepted_receipt_gate_call_receipt_consumer_enablement_receipt_storage_record_writer_authorization_receipt_acceptance.v1" ||
		payload["request_type"] != command ||
		payload["preview_decision"] != "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-receipt-accepted-receipt-gate-call-authorization-receipt-accepted-receipt-gate-call-receipt-consumer-enablement-receipt-storage-record-writer-authorization-receipt-acceptance-ready-acceptance-disabled" {
		t.Fatalf("unexpected storage record writer authorization receipt acceptance payload: %s", output.String())
	}
	for _, key := range []string{
		"current_mainline_consumed",
		"route_enablement_lookup_route_dispatch_dry_run_result_notification_action_request_object_dispatch_dry_run_result_persistence_receipt_consumer_enablement_receipt_storage_record_writer_call_authorization_receipt_accepted_receipt_gate_call_authorization_receipt_accepted_receipt_gate_call_receipt_consumer_enablement_receipt_storage_record_writer_authorization_receipt_consumed",
		"route_enablement_lookup_route_dispatch_dry_run_result_notification_action_request_object_dispatch_dry_run_result_persistence_receipt_consumer_enablement_receipt_storage_record_writer_call_authorization_receipt_accepted_receipt_gate_call_authorization_receipt_accepted_receipt_gate_call_receipt_consumer_enablement_receipt_storage_record_writer_authorization_receipt_acceptance_authorization_receipt_evidence_consumed",
		"lookup_route_dispatch_dry_run_result_notification_action_request_object_dispatch_dry_run_result_persistence_receipt_consumer_enablement_receipt_storage_record_writer_call_authorization_receipt_accepted_receipt_gate_call_authorization_receipt_accepted_receipt_gate_call_receipt_consumer_enablement_receipt_storage_record_writer_authorization_receipt_acceptance_ready",
		"result_persistence_receipt_consumer_enablement_receipt_storage_record_writer_call_authorization_receipt_accepted_receipt_gate_call_authorization_receipt_accepted_receipt_gate_call_receipt_consumer_enablement_receipt_storage_record_writer_authorization_receipt_acceptance_ready",
		"receipt_storage_record_writer_authorization_receipt_acceptance_boundary_modeled",
		"kde_safe_redacted_result_only",
		"raw_result_hidden",
	} {
		if payload[key] != true {
			t.Fatalf("expected %s to be true in storage record writer authorization receipt acceptance payload: %s", key, output.String())
		}
	}
	if payload["notification_action_request_dispatch_dry_run_result_persistence_receipt_consumer_enablement_receipt_storage_record_writer_call_authorization_receipt_accepted_receipt_gate_call_authorization_receipt_accepted_receipt_gate_call_receipt_consumer_enablement_receipt_storage_record_writer_authorization_receipt_acceptance_item_count"] != float64(4) ||
		payload["ready_notification_action_request_dispatch_dry_run_result_persistence_receipt_consumer_enablement_receipt_storage_record_writer_call_authorization_receipt_accepted_receipt_gate_call_authorization_receipt_accepted_receipt_gate_call_receipt_consumer_enablement_receipt_storage_record_writer_authorization_receipt_acceptance_item_count"] != float64(4) ||
		payload["missing_notification_action_request_dispatch_dry_run_result_persistence_receipt_consumer_enablement_receipt_storage_record_writer_call_authorization_receipt_accepted_receipt_gate_call_authorization_receipt_accepted_receipt_gate_call_receipt_consumer_enablement_receipt_storage_record_writer_authorization_receipt_acceptance_item_count"] != float64(0) ||
		payload["callable_result_persistence_receipt_consumer_enablement_receipt_storage_record_writer_authorization_receipt_acceptance_item_count"] != float64(0) ||
		payload["granted_result_persistence_receipt_consumer_enablement_receipt_storage_record_writer_authorization_receipt_acceptance_item_count"] != float64(0) ||
		payload["persisted_result_persistence_receipt_consumer_enablement_receipt_storage_record_writer_authorization_receipt_acceptance_item_count"] != float64(0) ||
		payload["side_effect_notification_action_request_dispatch_dry_run_result_persistence_receipt_consumer_enablement_receipt_storage_record_writer_call_authorization_receipt_accepted_receipt_gate_call_authorization_receipt_accepted_receipt_gate_call_receipt_consumer_enablement_receipt_storage_record_writer_authorization_receipt_acceptance_item_count"] != float64(0) {
		t.Fatalf("unexpected storage record writer authorization receipt acceptance counts: %s", output.String())
	}
	if !strings.Contains(output.String(), "storage-record-writer-call-authorization-receipt-accepted-receipt-gate-call-authorization-receipt-accepted-receipt-gate-call-receipt-consumer-enablement-receipt-storage-record-writer-authorization-receipt-acceptance-install-failure") ||
		!strings.Contains(output.String(), "storage-record-writer-authorization-receipt-consumed") ||
		!strings.Contains(output.String(), "storage-record-writer-authorization-receipt-acceptance-authorization-receipt-evidence-consumed") ||
		!strings.Contains(output.String(), "storage-record-writer-authorization-receipt-acceptance-required-but-not-accepted") ||
		!strings.Contains(output.String(), "storage-record-writer-authorization-receipt-acceptance-writes-visibility-persistence-and-dispatch-disabled") {
		t.Fatalf("storage record writer authorization receipt acceptance payload must include item and check evidence: %s", output.String())
	}
	for _, key := range []string{"receipt_consumer_enablement_receipt_storage_record_writer_authorization_receipt_acceptance_callable", "receipt_consumer_enablement_receipt_storage_record_writer_authorization_receipt_acceptance_enabled", "receipt_consumer_enablement_receipt_storage_record_writer_authorization_receipt_acceptance_granted", "receipt_consumer_enablement_receipt_storage_record_writer_authorization_receipt_accepted", "receipt_consumer_enablement_receipt_storage_record_writer_authorization_receipt_callable", "receipt_consumer_enablement_receipt_storage_record_writer_authorization_receipt_present", "receipt_consumer_enablement_receipt_storage_record_writer_authorized", "receipt_consumer_enablement_receipt_storage_record_written", "receipt_consumer_enablement_receipt_storage_persistence_gate_passed", "receipt_consumer_enablement_receipt_persisted", "receipt_consumer_enablement_receipt_present", "receipt_consumer_enablement_enabled", "consumer_enabled", "kde_consumer_enabled", "runtime_consumer_enabled", "receipt_consumption_enabled", "receipt_consumed", "receipt_acceptance_enabled", "receipt_accepted", "receipt_write_enabled", "dry_run_result_persistence_enabled", "request_object_dispatch_enabled", "notification_action_enabled", "notification_sent", "user_visible", "storage_write_enabled", "production_ownership_ready", "production_bus_claimed", "write_methods_enabled", "runtime_writes_enabled", "backend_launch_enabled", "state_root_path_exposed", "file_paths_exposed", "file_content_read", "host_root_modified"} {
		if payload[key] != false {
			t.Fatalf("expected %s to remain false in storage record writer authorization receipt acceptance payload: %s", key, output.String())
		}
	}
}

func TestProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptStorageRecordWriterCallAuthorizationReceiptAcceptedReceiptGateCallAuthorizationReceiptAcceptedReceiptGateCallReceiptConsumerEnablementReceiptStorageRecordWriterAuthorizationReceiptAcceptancePreviewCommandRejectsPositionalArgs(t *testing.T) {
	var output bytes.Buffer
	if err := run([]string{"production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-receipt-accepted-receipt-gate-call-authorization-receipt-accepted-receipt-gate-call-receipt-consumer-enablement-receipt-storage-record-writer-authorization-receipt-acceptance-preview", "extra"}, &output); err == nil {
		t.Fatalf("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-receipt-accepted-receipt-gate-call-authorization-receipt-accepted-receipt-gate-call-receipt-consumer-enablement-receipt-storage-record-writer-authorization-receipt-acceptance-preview must reject positional arguments")
	}
}

func TestProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptStorageRecordWriterCallAuthorizationReceiptAcceptedReceiptGateCallAuthorizationReceiptAcceptedReceiptGateCallReceiptConsumerEnablementReceiptStorageRecordWriterAcceptedReceiptGatePreviewCommand(t *testing.T) {
	root := projectRootForRuntimeServiceBindingCommandTest(t)
	var output bytes.Buffer
	command := "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-receipt-accepted-receipt-gate-call-authorization-receipt-accepted-receipt-gate-call-receipt-consumer-enablement-receipt-storage-record-writer-accepted-receipt-gate-preview"
	if err := run([]string{command, "--root", root}, &output); err != nil {
		t.Fatalf("%s returned error: %v", command, err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["version"] != currentProjectVersion(t) ||
		payload["schema_version"] != "xnix.runtime.production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_kde_safe_redacted_status_closed_record_writer_storage_root_authorization_receipt_dry_run_lookup_result_consumer_projection_route_enablement_lookup_route_dispatch_dry_run_result_notification_action_request_object_dispatch_dry_run_result_persistence_receipt_consumer_enablement_receipt_storage_record_writer_call_authorization_receipt_accepted_receipt_gate_call_authorization_receipt_accepted_receipt_gate_call_receipt_consumer_enablement_receipt_storage_record_writer_accepted_receipt_gate.v1" ||
		payload["request_type"] != command ||
		payload["preview_decision"] != "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-receipt-accepted-receipt-gate-call-authorization-receipt-accepted-receipt-gate-call-receipt-consumer-enablement-receipt-storage-record-writer-accepted-receipt-gate-ready-gate-disabled" {
		t.Fatalf("unexpected storage record writer accepted receipt gate payload: %s", output.String())
	}
	for _, key := range []string{
		"current_mainline_consumed",
		"route_enablement_lookup_route_dispatch_dry_run_result_notification_action_request_object_dispatch_dry_run_result_persistence_receipt_consumer_enablement_receipt_storage_record_writer_call_authorization_receipt_accepted_receipt_gate_call_authorization_receipt_accepted_receipt_gate_call_receipt_consumer_enablement_receipt_storage_record_writer_authorization_receipt_acceptance_consumed",
		"route_enablement_lookup_route_dispatch_dry_run_result_notification_action_request_object_dispatch_dry_run_result_persistence_receipt_consumer_enablement_receipt_storage_record_writer_call_authorization_receipt_accepted_receipt_gate_call_authorization_receipt_accepted_receipt_gate_call_receipt_consumer_enablement_receipt_storage_record_writer_accepted_receipt_gate_authorization_receipt_acceptance_evidence_consumed",
		"lookup_route_dispatch_dry_run_result_notification_action_request_object_dispatch_dry_run_result_persistence_receipt_consumer_enablement_receipt_storage_record_writer_call_authorization_receipt_accepted_receipt_gate_call_authorization_receipt_accepted_receipt_gate_call_receipt_consumer_enablement_receipt_storage_record_writer_accepted_receipt_gate_ready",
		"result_persistence_receipt_consumer_enablement_receipt_storage_record_writer_call_authorization_receipt_accepted_receipt_gate_call_authorization_receipt_accepted_receipt_gate_call_receipt_consumer_enablement_receipt_storage_record_writer_accepted_receipt_gate_ready",
		"receipt_storage_record_writer_accepted_receipt_gate_boundary_modeled",
		"kde_safe_redacted_result_only",
		"raw_result_hidden",
	} {
		if payload[key] != true {
			t.Fatalf("expected %s to be true in storage record writer accepted receipt gate payload: %s", key, output.String())
		}
	}
	if payload["notification_action_request_dispatch_dry_run_result_persistence_receipt_consumer_enablement_receipt_storage_record_writer_call_authorization_receipt_accepted_receipt_gate_call_authorization_receipt_accepted_receipt_gate_call_receipt_consumer_enablement_receipt_storage_record_writer_accepted_receipt_gate_item_count"] != float64(4) ||
		payload["ready_notification_action_request_dispatch_dry_run_result_persistence_receipt_consumer_enablement_receipt_storage_record_writer_call_authorization_receipt_accepted_receipt_gate_call_authorization_receipt_accepted_receipt_gate_call_receipt_consumer_enablement_receipt_storage_record_writer_accepted_receipt_gate_item_count"] != float64(4) ||
		payload["missing_notification_action_request_dispatch_dry_run_result_persistence_receipt_consumer_enablement_receipt_storage_record_writer_call_authorization_receipt_accepted_receipt_gate_call_authorization_receipt_accepted_receipt_gate_call_receipt_consumer_enablement_receipt_storage_record_writer_accepted_receipt_gate_item_count"] != float64(0) ||
		payload["persistence_receipt_consumer_enablement_receipt_storage_record_writer_call_authorization_receipt_accepted_receipt_gate_call_authorization_receipt_accepted_receipt_gate_call_receipt_consumer_enablement_receipt_storage_record_writer_accepted_receipt_gate_authorization_receipt_acceptance_evidence_consumed_item_count"] != float64(4) ||
		payload["callable_result_persistence_receipt_consumer_enablement_receipt_storage_record_writer_accepted_receipt_gate_item_count"] != float64(0) ||
		payload["accepted_result_persistence_receipt_consumer_enablement_receipt_storage_record_writer_accepted_receipt_gate_item_count"] != float64(0) ||
		payload["side_effect_notification_action_request_dispatch_dry_run_result_persistence_receipt_consumer_enablement_receipt_storage_record_writer_call_authorization_receipt_accepted_receipt_gate_call_authorization_receipt_accepted_receipt_gate_call_receipt_consumer_enablement_receipt_storage_record_writer_accepted_receipt_gate_item_count"] != float64(0) {
		t.Fatalf("unexpected storage record writer accepted receipt gate counts: %s", output.String())
	}
	if !strings.Contains(output.String(), "storage-record-writer-call-authorization-receipt-accepted-receipt-gate-call-authorization-receipt-accepted-receipt-gate-call-receipt-consumer-enablement-receipt-storage-record-writer-accepted-receipt-gate-install-failure") ||
		!strings.Contains(output.String(), "storage-record-writer-authorization-receipt-acceptance-consumed") ||
		!strings.Contains(output.String(), "storage-record-writer-accepted-receipt-gate-authorization-receipt-acceptance-evidence-consumed") ||
		!strings.Contains(output.String(), "storage-record-writer-accepted-receipt-gate-required-but-not-accepted") ||
		!strings.Contains(output.String(), "storage-record-writer-accepted-receipt-gate-writes-visibility-persistence-and-dispatch-disabled") {
		t.Fatalf("storage record writer accepted receipt gate payload must include item and check evidence: %s", output.String())
	}
	for _, key := range []string{"receipt_consumer_enablement_receipt_storage_record_writer_accepted_receipt_gate_callable", "receipt_consumer_enablement_receipt_storage_record_writer_accepted_receipt_gate_enabled", "receipt_consumer_enablement_receipt_storage_record_writer_accepted_receipt_gate_granted", "receipt_consumer_enablement_receipt_storage_record_writer_accepted_receipt_gate_accepted", "receipt_consumer_enablement_receipt_storage_record_writer_authorization_receipt_acceptance_callable", "receipt_consumer_enablement_receipt_storage_record_writer_authorization_receipt_accepted", "receipt_consumer_enablement_receipt_storage_record_written", "receipt_consumer_enablement_receipt_storage_persistence_gate_passed", "receipt_consumer_enablement_receipt_persisted", "receipt_consumer_enablement_receipt_present", "receipt_consumer_enablement_enabled", "consumer_enabled", "kde_consumer_enabled", "runtime_consumer_enabled", "receipt_consumption_enabled", "receipt_consumed", "receipt_accepted", "receipt_write_enabled", "dry_run_result_persistence_enabled", "request_object_dispatch_enabled", "notification_action_enabled", "notification_sent", "user_visible", "storage_write_enabled", "production_ownership_ready", "production_bus_claimed", "write_methods_enabled", "runtime_writes_enabled", "backend_launch_enabled", "state_root_path_exposed", "file_paths_exposed", "file_content_read", "host_root_modified"} {
		if payload[key] != false {
			t.Fatalf("expected %s to remain false in storage record writer accepted receipt gate payload: %s", key, output.String())
		}
	}
}

func TestProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptStorageRecordWriterCallAuthorizationReceiptAcceptedReceiptGateCallAuthorizationReceiptAcceptedReceiptGateCallReceiptConsumerEnablementReceiptStorageRecordWriterAcceptedReceiptGatePreviewCommandRejectsPositionalArgs(t *testing.T) {
	var output bytes.Buffer
	if err := run([]string{"production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-receipt-accepted-receipt-gate-call-authorization-receipt-accepted-receipt-gate-call-receipt-consumer-enablement-receipt-storage-record-writer-accepted-receipt-gate-preview", "extra"}, &output); err == nil {
		t.Fatalf("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-receipt-accepted-receipt-gate-call-authorization-receipt-accepted-receipt-gate-call-receipt-consumer-enablement-receipt-storage-record-writer-accepted-receipt-gate-preview must reject positional arguments")
	}
}
