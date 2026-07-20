package owner

import (
	"os"
	"path/filepath"
	"testing"
)

func TestProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptStorageRecordWriterCallAuthorizationReceiptAcceptedReceiptGatePreviewModelsAcceptedReceiptGateBoundary(t *testing.T) {
	preview, err := NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptStorageRecordWriterCallAuthorizationReceiptAcceptedReceiptGatePreview(projectRootForOwnerTest(t))
	if err != nil {
		t.Fatalf("NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptStorageRecordWriterCallAuthorizationReceiptAcceptedReceiptGatePreview returned error: %v", err)
	}
	if preview["version"] != currentProjectVersion(t) ||
		preview["schema_version"] != "xnix.runtime.production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_kde_safe_redacted_status_closed_record_writer_storage_root_authorization_receipt_dry_run_lookup_result_consumer_projection_route_enablement_lookup_route_dispatch_dry_run_result_notification_action_request_object_dispatch_dry_run_result_persistence_receipt_consumer_enablement_receipt_storage_record_writer_call_authorization_receipt_accepted_receipt_gate.v1" ||
		preview["request_type"] != routeEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptStorageRecordWriterCallAuthorizationReceiptAcceptedReceiptGateRequest ||
		preview["preview_decision"] != "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-receipt-accepted-receipt-gate-ready-gate-disabled" {
		t.Fatalf("unexpected storage record writer call authorization receipt accepted receipt gate schema: %#v", preview)
	}
	for _, key := range []string{
		"current_mainline_consumed",
		"route_enablement_lookup_route_dispatch_dry_run_result_notification_action_request_object_dispatch_dry_run_result_persistence_receipt_consumer_enablement_receipt_storage_record_writer_call_authorization_receipt_acceptance_consumed",
		"lookup_route_dispatch_dry_run_result_notification_action_request_object_dispatch_dry_run_result_persistence_receipt_consumer_enablement_receipt_storage_record_writer_call_authorization_receipt_acceptance_ready",
		"lookup_route_dispatch_dry_run_result_notification_action_request_object_dispatch_dry_run_result_persistence_receipt_consumer_enablement_receipt_storage_record_writer_call_authorization_receipt_accepted_receipt_gate_required",
		"lookup_route_dispatch_dry_run_result_notification_action_request_object_dispatch_dry_run_result_persistence_receipt_consumer_enablement_receipt_storage_record_writer_call_authorization_receipt_accepted_receipt_gate_modeled",
		"lookup_route_dispatch_dry_run_result_notification_action_request_object_dispatch_dry_run_result_persistence_receipt_consumer_enablement_receipt_storage_record_writer_call_authorization_receipt_accepted_receipt_gate_ready",
		"receipt_storage_record_writer_call_authorization_receipt_accepted_receipt_gate_boundary_modeled",
		"notification_action_request_object_dispatch_dry_run_result_persistence_receipt_consumer_enablement_receipt_storage_record_writer_call_authorization_receipt_accepted_receipt_gate_review_only",
		"result_persistence_receipt_consumer_enablement_receipt_storage_record_writer_call_authorization_receipt_acceptance_ready",
		"result_persistence_receipt_consumer_enablement_receipt_storage_record_writer_call_authorization_receipt_accepted_receipt_gate_required",
		"result_persistence_receipt_consumer_enablement_receipt_storage_record_writer_call_authorization_receipt_accepted_receipt_gate_modeled",
		"result_persistence_receipt_consumer_enablement_receipt_storage_record_writer_call_authorization_receipt_accepted_receipt_gate_ready",
		"runtime_owned",
		"go_runtime_backed",
	} {
		if preview[key] != true {
			t.Fatalf("expected %s to be true in storage record writer call authorization receipt accepted receipt gate preview: %#v", key, preview)
		}
	}
	for _, key := range []string{
		"receipt_consumer_enablement_receipt_storage_record_writer_call_authorization_receipt_accepted_receipt_gate_callable",
		"receipt_consumer_enablement_receipt_storage_record_writer_call_authorization_receipt_accepted_receipt_gate_enabled",
		"receipt_consumer_enablement_receipt_storage_record_writer_call_authorization_receipt_accepted_receipt_gate_granted",
		"receipt_consumer_enablement_receipt_storage_record_writer_call_authorization_receipt_acceptance_callable",
		"receipt_consumer_enablement_receipt_storage_record_writer_call_authorization_receipt_acceptance_enabled",
		"receipt_consumer_enablement_receipt_storage_record_writer_call_authorization_receipt_acceptance_granted",
		"receipt_consumer_enablement_receipt_storage_record_writer_authorization_receipt_accepted",
		"receipt_consumer_enablement_receipt_storage_record_writer_authorized",
		"receipt_consumer_enablement_receipt_storage_record_written",
		"receipt_consumer_enablement_receipt_storage_persistence_gate_passed",
		"receipt_consumer_enablement_receipt_persisted",
		"receipt_consumer_enablement_receipt_write_enabled",
		"receipt_consumer_enablement_enabled",
		"consumer_enabled",
		"kde_consumer_enabled",
		"runtime_consumer_enabled",
		"receipt_consumption_enabled",
		"receipt_consumed",
		"receipt_accepted",
		"receipt_write_enabled",
		"dry_run_result_persistence_enabled",
		"dispatch_dry_run_execution_enabled",
		"request_object_creation_enabled",
		"request_object_dispatch_enabled",
		"notification_action_enabled",
		"notification_sent",
		"user_visible",
		"storage_write_enabled",
		"production_ownership_ready",
		"production_bus_claimed",
		"write_methods_enabled",
		"runtime_writes_enabled",
		"backend_launch_enabled",
		"host_root_modified",
		"state_root_path_exposed",
		"file_paths_exposed",
		"file_content_read",
	} {
		if preview[key] != false {
			t.Fatalf("expected %s to remain false in storage record writer call authorization receipt accepted receipt gate preview: %#v", key, preview)
		}
	}
	if preview["notification_action_request_dispatch_dry_run_result_persistence_receipt_consumer_enablement_receipt_storage_record_writer_call_authorization_receipt_accepted_receipt_gate_item_count"] != 4 ||
		preview["ready_notification_action_request_dispatch_dry_run_result_persistence_receipt_consumer_enablement_receipt_storage_record_writer_call_authorization_receipt_accepted_receipt_gate_item_count"] != 4 ||
		preview["missing_notification_action_request_dispatch_dry_run_result_persistence_receipt_consumer_enablement_receipt_storage_record_writer_call_authorization_receipt_accepted_receipt_gate_item_count"] != 0 ||
		preview["persistence_receipt_consumer_enablement_receipt_storage_record_writer_call_authorization_receipt_accepted_receipt_gate_modeled_item_count"] != 4 ||
		preview["side_effect_notification_action_request_dispatch_dry_run_result_persistence_receipt_consumer_enablement_receipt_storage_record_writer_call_authorization_receipt_accepted_receipt_gate_item_count"] != 0 {
		t.Fatalf("unexpected storage record writer call authorization receipt accepted receipt gate counts: %#v", preview)
	}
	for _, item := range preview["notification_action_request_dispatch_dry_run_result_persistence_receipt_consumer_enablement_receipt_storage_record_writer_call_authorization_receipt_accepted_receipt_gate_items"].([]map[string]any) {
		if item["persistence_receipt_consumer_enablement_receipt_storage_record_writer_call_authorization_receipt_acceptance_ready"] != true ||
			item["result_persistence_receipt_consumer_enablement_receipt_storage_record_writer_call_authorization_receipt_accepted_receipt_gate_ready"] != true ||
			item["receipt_storage_record_writer_call_authorization_receipt_accepted_receipt_gate_boundary_modeled"] != true ||
			item["receipt_consumer_enablement_receipt_storage_record_writer_call_authorization_receipt_accepted_receipt_gate_callable"] != false ||
			item["receipt_consumer_enablement_receipt_storage_record_writer_call_authorization_receipt_accepted_receipt_gate_enabled"] != false ||
			item["receipt_consumer_enablement_receipt_storage_record_writer_call_authorization_receipt_accepted_receipt_gate_granted"] != false ||
			item["receipt_consumer_enablement_receipt_storage_record_writer_authorization_receipt_accepted"] != false ||
			item["receipt_consumer_enablement_receipt_storage_record_written"] != false ||
			item["receipt_consumer_enablement_receipt_storage_persistence_gate_passed"] != false ||
			item["receipt_consumer_enablement_enabled"] != false ||
			item["consumer_enabled"] != false ||
			item["receipt_consumption_enabled"] != false ||
			item["dry_run_result_persistence_enabled"] != false ||
			item["dispatch_dry_run_execution_enabled"] != false ||
			item["request_object_creation_enabled"] != false ||
			item["notification_action_enabled"] != false ||
			item["storage_write_enabled"] != false ||
			item["side_effects_disabled"] != true ||
			item["state_root_path_exposed"] != false ||
			item["file_paths_exposed"] != false ||
			item["file_content_read"] != false ||
			item["host_root_modified"] != false ||
			item["notification_action_request_object_dispatch_dry_run_result_persistence_receipt_consumer_enablement_receipt_storage_record_writer_call_authorization_receipt_accepted_receipt_gate_status"] != "redacted-status-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-receipt-accepted-receipt-gate-ready-gate-disabled" {
			t.Fatalf("unsafe storage record writer call authorization receipt accepted receipt gate item: %#v", item)
		}
	}
	counts := preview["counts"].(map[string]int)
	if counts["total"] != 7 || counts["passed"] != 7 || counts["blocked"] != 0 || counts["pending"] != 0 {
		t.Fatalf("unexpected storage record writer call authorization receipt accepted receipt gate checks: %#v", counts)
	}
	if err := validateNoBackendTerms(preview, "KDE-safe storage record writer call authorization receipt accepted receipt gate test"); err != nil {
		t.Fatalf("validateNoBackendTerms returned error: %v", err)
	}
}

func TestProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptStorageRecordWriterCallAuthorizationReceiptAcceptedReceiptGatePreviewFailsClosedWithoutMainlineEvidence(t *testing.T) {
	root, err := os.MkdirTemp("", "xnix-action-request-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-call-authorization-receipt-accepted-receipt-gate-*")
	if err != nil {
		t.Fatalf("MkdirTemp returned error: %v", err)
	}
	t.Cleanup(func() { _ = os.RemoveAll(root) })
	if err := os.WriteFile(filepath.Join(root, "VERSION"), []byte(currentProjectVersion(t)+"\n"), 0o644); err != nil {
		t.Fatalf("WriteFile VERSION returned error: %v", err)
	}
	if err := os.MkdirAll(filepath.Join(root, "docs"), 0o755); err != nil {
		t.Fatalf("MkdirAll docs returned error: %v", err)
	}
	if err := os.WriteFile(filepath.Join(root, "docs", "xnix-current-mainline.md"), []byte("missing call authorization receipt accepted receipt gate evidence"), 0o644); err != nil {
		t.Fatalf("WriteFile mainline returned error: %v", err)
	}

	preview, err := NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptStorageRecordWriterCallAuthorizationReceiptAcceptedReceiptGatePreview(root)
	if err != nil {
		t.Fatalf("NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptStorageRecordWriterCallAuthorizationReceiptAcceptedReceiptGatePreview returned error: %v", err)
	}
	if preview["preview_decision"] != "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-receipt-accepted-receipt-gate-blocked" ||
		preview["current_mainline_consumed"] != false ||
		preview["route_enablement_lookup_route_dispatch_dry_run_result_notification_action_request_object_dispatch_dry_run_result_persistence_receipt_consumer_enablement_receipt_storage_record_writer_call_authorization_receipt_acceptance_consumed"] != false ||
		preview["lookup_route_dispatch_dry_run_result_notification_action_request_object_dispatch_dry_run_result_persistence_receipt_consumer_enablement_receipt_storage_record_writer_call_authorization_receipt_accepted_receipt_gate_ready"] != false ||
		preview["receipt_consumer_enablement_receipt_storage_record_writer_call_authorization_receipt_accepted_receipt_gate_callable"] != false ||
		preview["receipt_consumer_enablement_receipt_storage_record_writer_call_authorization_receipt_accepted_receipt_gate_enabled"] != false ||
		preview["receipt_consumer_enablement_receipt_storage_record_writer_call_authorization_receipt_accepted_receipt_gate_granted"] != false ||
		preview["receipt_consumer_enablement_receipt_storage_record_writer_authorization_receipt_accepted"] != false ||
		preview["ready_notification_action_request_dispatch_dry_run_result_persistence_receipt_consumer_enablement_receipt_storage_record_writer_call_authorization_receipt_accepted_receipt_gate_item_count"] != 0 {
		t.Fatalf("expected storage record writer call authorization receipt accepted receipt gate preview to fail closed without evidence: %#v", preview)
	}
}
