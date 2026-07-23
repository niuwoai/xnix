package owner

import (
	"os"
	"path/filepath"
	"testing"
)

func TestProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptStorageRecordWriterCallAuthorizationReceiptAcceptedReceiptGateCallAuthorizationReceiptAcceptedReceiptGateCallReceiptConsumerEnablementReceiptStorageRecordWriterAuthorizationReceiptAcceptancePreviewModelsWriterAuthorizationReceiptAcceptanceBoundary(t *testing.T) {
	preview, err := NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptStorageRecordWriterCallAuthorizationReceiptAcceptedReceiptGateCallAuthorizationReceiptAcceptedReceiptGateCallReceiptConsumerEnablementReceiptStorageRecordWriterAuthorizationReceiptAcceptancePreview(projectRootForOwnerTest(t))
	if err != nil {
		t.Fatalf("NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptStorageRecordWriterCallAuthorizationReceiptAcceptedReceiptGateCallAuthorizationReceiptAcceptedReceiptGateCallReceiptConsumerEnablementReceiptStorageRecordWriterAuthorizationReceiptAcceptancePreview returned error: %v", err)
	}
	if preview["version"] != currentProjectVersion(t) ||
		preview["schema_version"] != "xnix.runtime.production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_kde_safe_redacted_status_closed_record_writer_storage_root_authorization_receipt_dry_run_lookup_result_consumer_projection_route_enablement_lookup_route_dispatch_dry_run_result_notification_action_request_object_dispatch_dry_run_result_persistence_receipt_consumer_enablement_receipt_storage_record_writer_call_authorization_receipt_accepted_receipt_gate_call_authorization_receipt_accepted_receipt_gate_call_receipt_consumer_enablement_receipt_storage_record_writer_authorization_receipt_acceptance.v1" ||
		preview["request_type"] != routeEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptStorageRecordWriterCallAuthorizationReceiptAcceptedReceiptGateCallAuthorizationReceiptAcceptedReceiptGateCallReceiptConsumerEnablementReceiptStorageRecordWriterAuthorizationReceiptAcceptanceRequest ||
		preview["preview_decision"] != "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-receipt-accepted-receipt-gate-call-authorization-receipt-accepted-receipt-gate-call-receipt-consumer-enablement-receipt-storage-record-writer-authorization-receipt-acceptance-ready-acceptance-disabled" {
		t.Fatalf("unexpected storage record writer authorization receipt acceptance schema: %#v", preview)
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
		"runtime_owned",
		"go_runtime_backed",
	} {
		if preview[key] != true {
			t.Fatalf("expected %s to be true in storage record writer authorization receipt acceptance preview: %#v", key, preview)
		}
	}
	for _, key := range []string{
		"receipt_consumer_enablement_receipt_storage_record_writer_authorization_receipt_acceptance_callable",
		"receipt_consumer_enablement_receipt_storage_record_writer_authorization_receipt_acceptance_enabled",
		"receipt_consumer_enablement_receipt_storage_record_writer_authorization_receipt_acceptance_granted",
		"receipt_consumer_enablement_receipt_storage_record_writer_authorization_receipt_accepted",
		"receipt_consumer_enablement_receipt_storage_record_writer_authorization_receipt_callable",
		"receipt_consumer_enablement_receipt_storage_record_writer_authorization_receipt_present",
		"receipt_consumer_enablement_receipt_storage_record_writer_authorization_receipt_written",
		"receipt_consumer_enablement_receipt_storage_record_writer_authorization_receipt_persisted",
		"receipt_consumer_enablement_receipt_storage_record_writer_authorized",
		"receipt_consumer_enablement_receipt_storage_record_writer_authorization_granted",
		"receipt_consumer_enablement_receipt_storage_record_written",
		"receipt_consumer_enablement_receipt_storage_persistence_gate_passed",
		"receipt_consumer_enablement_receipt_persisted",
		"receipt_consumer_enablement_receipt_present",
		"receipt_consumer_enablement_receipt_write_enabled",
		"receipt_consumer_enablement_enabled",
		"consumer_enabled",
		"kde_consumer_enabled",
		"runtime_consumer_enabled",
		"receipt_consumption_enabled",
		"receipt_consumed",
		"receipt_acceptance_enabled",
		"receipt_accepted",
		"receipt_write_enabled",
		"dry_run_result_persistence_enabled",
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
			t.Fatalf("expected %s to remain false in storage record writer authorization receipt acceptance preview: %#v", key, preview)
		}
	}
	if preview["notification_action_request_dispatch_dry_run_result_persistence_receipt_consumer_enablement_receipt_storage_record_writer_call_authorization_receipt_accepted_receipt_gate_call_authorization_receipt_accepted_receipt_gate_call_receipt_consumer_enablement_receipt_storage_record_writer_authorization_receipt_acceptance_item_count"] != 4 ||
		preview["ready_notification_action_request_dispatch_dry_run_result_persistence_receipt_consumer_enablement_receipt_storage_record_writer_call_authorization_receipt_accepted_receipt_gate_call_authorization_receipt_accepted_receipt_gate_call_receipt_consumer_enablement_receipt_storage_record_writer_authorization_receipt_acceptance_item_count"] != 4 ||
		preview["missing_notification_action_request_dispatch_dry_run_result_persistence_receipt_consumer_enablement_receipt_storage_record_writer_call_authorization_receipt_accepted_receipt_gate_call_authorization_receipt_accepted_receipt_gate_call_receipt_consumer_enablement_receipt_storage_record_writer_authorization_receipt_acceptance_item_count"] != 0 ||
		preview["callable_result_persistence_receipt_consumer_enablement_receipt_storage_record_writer_authorization_receipt_acceptance_item_count"] != 0 ||
		preview["granted_result_persistence_receipt_consumer_enablement_receipt_storage_record_writer_authorization_receipt_acceptance_item_count"] != 0 ||
		preview["persisted_result_persistence_receipt_consumer_enablement_receipt_storage_record_writer_authorization_receipt_acceptance_item_count"] != 0 ||
		preview["written_result_persistence_receipt_consumer_enablement_receipt_storage_record_item_count"] != 0 {
		t.Fatalf("unexpected storage record writer authorization receipt acceptance counts: %#v", preview)
	}
	items := preview["notification_action_request_dispatch_dry_run_result_persistence_receipt_consumer_enablement_receipt_storage_record_writer_call_authorization_receipt_accepted_receipt_gate_call_authorization_receipt_accepted_receipt_gate_call_receipt_consumer_enablement_receipt_storage_record_writer_authorization_receipt_acceptance_items"].([]map[string]any)
	if len(items) != 4 {
		t.Fatalf("expected four storage record writer authorization receipt acceptance items: %#v", preview)
	}
	for _, item := range items {
		if item["evidence_present"] != true ||
			item["receipt_storage_record_writer_authorization_receipt_acceptance_boundary_modeled"] != true ||
			item["receipt_consumer_enablement_receipt_storage_record_writer_authorization_receipt_acceptance_callable"] != false ||
			item["receipt_consumer_enablement_receipt_storage_record_writer_authorization_receipt_accepted"] != false ||
			item["receipt_consumer_enablement_receipt_storage_record_written"] != false ||
			item["notification_action_enabled"] != false ||
			item["side_effects_disabled"] != true ||
			item["host_root_modified"] != false {
			t.Fatalf("unsafe storage record writer authorization receipt acceptance item: %#v", item)
		}
	}
	counts := preview["counts"].(map[string]int)
	if counts["total"] != 8 || counts["passed"] != 8 || counts["blocked"] != 0 || counts["pending"] != 0 {
		t.Fatalf("unexpected storage record writer authorization receipt acceptance checks: %#v", counts)
	}
	if err := validateNoBackendTerms(preview, "KDE-safe storage record writer authorization receipt acceptance test"); err != nil {
		t.Fatalf("validateNoBackendTerms returned error: %v", err)
	}
}

func TestProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptStorageRecordWriterCallAuthorizationReceiptAcceptedReceiptGateCallAuthorizationReceiptAcceptedReceiptGateCallReceiptConsumerEnablementReceiptStorageRecordWriterAuthorizationReceiptAcceptancePreviewFailsClosedWithoutMainlineEvidence(t *testing.T) {
	root, err := os.MkdirTemp("", "xnix-storage-record-writer-authorization-receipt-acceptance-*")
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
	if err := os.WriteFile(filepath.Join(root, "docs", "xnix-current-mainline.md"), []byte("missing storage record writer authorization receipt acceptance evidence"), 0o644); err != nil {
		t.Fatalf("WriteFile mainline returned error: %v", err)
	}

	preview, err := NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptStorageRecordWriterCallAuthorizationReceiptAcceptedReceiptGateCallAuthorizationReceiptAcceptedReceiptGateCallReceiptConsumerEnablementReceiptStorageRecordWriterAuthorizationReceiptAcceptancePreview(root)
	if err != nil {
		t.Fatalf("NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptStorageRecordWriterCallAuthorizationReceiptAcceptedReceiptGateCallAuthorizationReceiptAcceptedReceiptGateCallReceiptConsumerEnablementReceiptStorageRecordWriterAuthorizationReceiptAcceptancePreview returned error: %v", err)
	}
	if preview["preview_decision"] != "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-receipt-accepted-receipt-gate-call-authorization-receipt-accepted-receipt-gate-call-receipt-consumer-enablement-receipt-storage-record-writer-authorization-receipt-acceptance-blocked" ||
		preview["current_mainline_consumed"] != false ||
		preview["route_enablement_lookup_route_dispatch_dry_run_result_notification_action_request_object_dispatch_dry_run_result_persistence_receipt_consumer_enablement_receipt_storage_record_writer_call_authorization_receipt_accepted_receipt_gate_call_authorization_receipt_accepted_receipt_gate_call_receipt_consumer_enablement_receipt_storage_record_writer_authorization_receipt_consumed"] != false ||
		preview["route_enablement_lookup_route_dispatch_dry_run_result_notification_action_request_object_dispatch_dry_run_result_persistence_receipt_consumer_enablement_receipt_storage_record_writer_call_authorization_receipt_accepted_receipt_gate_call_authorization_receipt_accepted_receipt_gate_call_receipt_consumer_enablement_receipt_storage_record_writer_authorization_receipt_acceptance_authorization_receipt_evidence_consumed"] != false ||
		preview["ready_notification_action_request_dispatch_dry_run_result_persistence_receipt_consumer_enablement_receipt_storage_record_writer_call_authorization_receipt_accepted_receipt_gate_call_authorization_receipt_accepted_receipt_gate_call_receipt_consumer_enablement_receipt_storage_record_writer_authorization_receipt_acceptance_item_count"] != 0 ||
		preview["receipt_consumer_enablement_receipt_storage_record_writer_authorization_receipt_acceptance_callable"] != false ||
		preview["receipt_consumer_enablement_receipt_storage_record_writer_authorization_receipt_accepted"] != false {
		t.Fatalf("expected storage record writer authorization receipt acceptance preview to fail closed without evidence: %#v", preview)
	}
}
