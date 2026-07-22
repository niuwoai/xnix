package owner

import (
	"os"
	"path/filepath"
	"testing"
)

func TestProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptStorageRecordContractPreviewModelsStorageRecordContractBoundary(t *testing.T) {
	preview, err := NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptStorageRecordContractPreview(projectRootForOwnerTest(t))
	if err != nil {
		t.Fatalf("NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptStorageRecordContractPreview returned error: %v", err)
	}
	if preview["version"] != currentProjectVersion(t) ||
		preview["schema_version"] != "xnix.runtime.production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_kde_safe_redacted_status_closed_record_writer_storage_root_authorization_receipt_dry_run_lookup_result_consumer_projection_route_enablement_lookup_route_dispatch_dry_run_result_notification_action_request_object_dispatch_dry_run_result_persistence_receipt_consumer_enablement_receipt_storage_record_contract.v1" ||
		preview["request_type"] != routeEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptStorageRecordContractRequest ||
		preview["preview_decision"] != "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-contract-ready-writes-disabled" {
		t.Fatalf("unexpected notification action request-object dispatch dry-run result persistence receipt consumer enablement receipt storage record contract schema: %#v", preview)
	}
	for _, key := range []string{
		"current_mainline_consumed",
		"route_enablement_lookup_route_dispatch_dry_run_result_notification_action_request_object_dispatch_dry_run_result_persistence_receipt_consumer_enablement_receipt_storage_persistence_gate_audit_consumed",
		"route_enablement_lookup_route_dispatch_dry_run_result_notification_action_request_object_dispatch_dry_run_result_persistence_receipt_consumer_enablement_receipt_closed_persistence_implementation_consumed",
		"route_enablement_lookup_route_dispatch_dry_run_result_notification_action_request_object_dispatch_dry_run_result_persistence_receipt_consumer_enablement_receipt_persistence_authorization_consumed",
		"route_enablement_lookup_route_dispatch_dry_run_result_notification_action_request_object_dispatch_dry_run_result_persistence_receipt_consumer_enablement_receipt_closed_implementation_consumed",
		"route_enablement_lookup_route_dispatch_dry_run_result_notification_action_request_object_dispatch_dry_run_result_persistence_receipt_consumer_enablement_receipt_enablement_audit_consumed",
		"route_enablement_lookup_route_dispatch_dry_run_result_notification_action_request_object_dispatch_dry_run_result_persistence_receipt_consumer_enablement_receipt_grant_audit_consumed",
		"lookup_route_dispatch_dry_run_result_notification_action_request_object_dispatch_dry_run_result_persistence_receipt_consumer_enablement_receipt_storage_persistence_gate_ready",
		"result_persistence_receipt_consumer_enablement_receipt_closed_persistence_implementation_accepted_receipt_gate_call_persistence_authorization_evidence_ready",
		"lookup_route_dispatch_dry_run_result_notification_action_request_object_dispatch_dry_run_result_persistence_receipt_consumer_enablement_receipt_storage_record_contract_required",
		"lookup_route_dispatch_dry_run_result_notification_action_request_object_dispatch_dry_run_result_persistence_receipt_consumer_enablement_receipt_storage_record_contract_modeled",
		"lookup_route_dispatch_dry_run_result_notification_action_request_object_dispatch_dry_run_result_persistence_receipt_consumer_enablement_receipt_storage_record_contract_ready",
		"receipt_storage_record_contract_boundary_modeled",
		"receipt_record_identity_boundary_modeled",
		"receipt_record_redaction_boundary_modeled",
		"receipt_record_retention_boundary_modeled",
		"receipt_storage_boundary_modeled",
		"receipt_retention_boundary_modeled",
		"notification_action_request_object_dispatch_dry_run_result_persistence_receipt_consumer_enablement_receipt_storage_record_contract_review_only",
		"result_persistence_receipt_consumer_enablement_receipt_storage_persistence_gate_ready",
		"result_persistence_receipt_consumer_enablement_receipt_storage_record_contract_required",
		"result_persistence_receipt_consumer_enablement_receipt_storage_record_contract_modeled",
		"result_persistence_receipt_consumer_enablement_receipt_storage_record_contract_ready",
		"kde_safe_redacted_result_only",
		"raw_result_hidden",
		"runtime_owned",
		"go_runtime_backed",
	} {
		if preview[key] != true {
			t.Fatalf("expected %s to be true in notification action request-object dispatch dry-run result persistence receipt consumer enablement receipt storage record contract preview: %#v", key, preview)
		}
	}
	for _, key := range []string{
		"receipt_consumer_enablement_receipt_storage_persistence_gate_passed",
		"receipt_consumer_enablement_receipt_storage_record_contract_enabled",
		"receipt_consumer_enablement_receipt_storage_record_contract_writable",
		"receipt_consumer_enablement_receipt_storage_record_written",
		"receipt_consumer_enablement_receipt_storage_persistence_enabled",
		"receipt_consumer_enablement_receipt_storage_persisted",
		"receipt_consumer_enablement_receipt_closed_persistence_implementation_enabled",
		"receipt_consumer_enablement_receipt_closed_persistence_implemented",
		"receipt_consumer_enablement_receipt_persist_authorized",
		"receipt_consumer_enablement_receipt_persistence_enabled",
		"receipt_consumer_enablement_receipt_persisted",
		"receipt_consumer_enablement_receipt_present",
		"receipt_consumer_enablement_receipt_written",
		"receipt_consumer_enablement_receipt_write_enabled",
		"receipt_consumer_enablement_receipt_callable",
		"receipt_consumer_enablement_receipt_implementation_enabled",
		"receipt_consumer_enablement_receipt_enablement_enabled",
		"receipt_consumer_enablement_receipt_enabled",
		"receipt_consumer_enablement_receipt_granted",
		"receipt_consumer_enablement_receipt_accepted",
		"receipt_consumer_enablement_receipt_consumed",
		"receipt_consumer_authorization_enabled",
		"consumer_authorization_granted",
		"receipt_consumer_enablement_enabled",
		"consumer_enabled",
		"kde_consumer_enabled",
		"runtime_consumer_enabled",
		"receipt_consumption_enabled",
		"receipt_consumed",
		"receipt_acceptance_enabled",
		"receipt_accepted",
		"receipt_write_enabled",
		"result_visibility_persistence_enabled",
		"dry_run_result_persistence_enabled",
		"dry_run_result_persisted",
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
			t.Fatalf("expected %s to remain false in notification action request-object dispatch dry-run result persistence receipt consumer enablement receipt storage record contract preview: %#v", key, preview)
		}
	}
	if preview["notification_action_request_dispatch_dry_run_result_persistence_receipt_consumer_enablement_receipt_storage_record_contract_item_count"] != 4 ||
		preview["ready_notification_action_request_dispatch_dry_run_result_persistence_receipt_consumer_enablement_receipt_storage_record_contract_item_count"] != 4 ||
		preview["missing_notification_action_request_dispatch_dry_run_result_persistence_receipt_consumer_enablement_receipt_storage_record_contract_item_count"] != 0 ||
		preview["persistence_receipt_consumer_enablement_receipt_storage_persistence_gate_consumed_item_count"] != 4 ||
		preview["persistence_receipt_consumer_enablement_receipt_storage_record_contract_modeled_item_count"] != 4 ||
		preview["writable_result_persistence_receipt_consumer_enablement_receipt_storage_record_item_count"] != 0 ||
		preview["written_result_persistence_receipt_consumer_enablement_receipt_storage_record_item_count"] != 0 ||
		preview["persisted_result_persistence_receipt_consumer_enablement_receipt_item_count"] != 0 ||
		preview["side_effect_notification_action_request_dispatch_dry_run_result_persistence_receipt_consumer_enablement_receipt_storage_record_contract_item_count"] != 0 {
		t.Fatalf("unexpected notification action request-object dispatch dry-run result persistence receipt consumer enablement receipt storage record contract counts: %#v", preview)
	}
	for _, item := range preview["notification_action_request_dispatch_dry_run_result_persistence_receipt_consumer_enablement_receipt_storage_record_contract_items"].([]map[string]any) {
		if item["evidence_present"] != true ||
			item["persistence_receipt_consumer_enablement_receipt_storage_persistence_gate_consumed"] != true ||
			item["persistence_receipt_consumer_enablement_receipt_closed_persistence_implementation_accepted_receipt_gate_call_persistence_authorization_evidence_ready"] != true ||
			item["result_persistence_receipt_consumer_enablement_receipt_storage_record_contract_ready"] != true ||
			item["receipt_storage_record_contract_boundary_modeled"] != true ||
			item["receipt_record_identity_boundary_modeled"] != true ||
			item["receipt_record_redaction_boundary_modeled"] != true ||
			item["receipt_consumer_enablement_receipt_storage_record_contract_enabled"] != false ||
			item["receipt_consumer_enablement_receipt_storage_record_contract_writable"] != false ||
			item["receipt_consumer_enablement_receipt_storage_record_written"] != false ||
			item["receipt_consumer_enablement_receipt_storage_persistence_gate_passed"] != false ||
			item["receipt_consumer_enablement_receipt_storage_persisted"] != false ||
			item["receipt_consumer_enablement_receipt_persisted"] != false ||
			item["receipt_consumer_enablement_receipt_present"] != false ||
			item["receipt_consumer_enablement_receipt_write_enabled"] != false ||
			item["receipt_consumer_enablement_receipt_callable"] != false ||
			item["receipt_consumer_enablement_enabled"] != false ||
			item["consumer_enabled"] != false ||
			item["dry_run_result_persistence_enabled"] != false ||
			item["dispatch_dry_run_execution_enabled"] != false ||
			item["request_object_creation_enabled"] != false ||
			item["notification_action_enabled"] != false ||
			item["storage_write_enabled"] != false ||
			item["side_effects_disabled"] != true ||
			item["state_root_path_exposed"] != false ||
			item["host_root_modified"] != false ||
			item["notification_action_request_object_dispatch_dry_run_result_persistence_receipt_consumer_enablement_receipt_storage_record_contract_status"] != "redacted-status-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-contract-ready-writes-disabled" {
			t.Fatalf("unsafe notification action request-object dispatch dry-run result persistence receipt consumer enablement receipt storage record contract item: %#v", item)
		}
	}
	counts := preview["counts"].(map[string]int)
	if counts["total"] != 8 || counts["passed"] != 8 || counts["blocked"] != 0 || counts["pending"] != 0 {
		t.Fatalf("unexpected notification action request-object dispatch dry-run result persistence receipt consumer enablement receipt storage record contract checks: %#v", counts)
	}
	if err := validateNoBackendTerms(preview, "KDE-safe notification action request-object dispatch dry-run result persistence receipt consumer enablement receipt storage record contract test"); err != nil {
		t.Fatalf("validateNoBackendTerms returned error: %v", err)
	}
}

func TestProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptStorageRecordContractPreviewFailsClosedWithoutSources(t *testing.T) {
	root, err := os.MkdirTemp("", "xnix-action-request-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-contract-*")
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
	if err := os.WriteFile(filepath.Join(root, "docs", "xnix-current-mainline.md"), []byte("missing notification action request-object dispatch dry-run result persistence receipt consumer enablement receipt storage record contract evidence"), 0o644); err != nil {
		t.Fatalf("WriteFile mainline returned error: %v", err)
	}

	preview, err := NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptStorageRecordContractPreview(root)
	if err != nil {
		t.Fatalf("NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptStorageRecordContractPreview returned error: %v", err)
	}
	if preview["preview_decision"] != "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-contract-blocked" ||
		preview["current_mainline_consumed"] != false ||
		preview["route_enablement_lookup_route_dispatch_dry_run_result_notification_action_request_object_dispatch_dry_run_result_persistence_receipt_consumer_enablement_receipt_storage_persistence_gate_audit_consumed"] != false ||
		preview["lookup_route_dispatch_dry_run_result_notification_action_request_object_dispatch_dry_run_result_persistence_receipt_consumer_enablement_receipt_storage_record_contract_ready"] != false ||
		preview["receipt_consumer_enablement_receipt_storage_record_contract_writable"] != false ||
		preview["receipt_consumer_enablement_receipt_storage_record_written"] != false ||
		preview["ready_notification_action_request_dispatch_dry_run_result_persistence_receipt_consumer_enablement_receipt_storage_record_contract_item_count"] != 0 ||
		preview["missing_notification_action_request_dispatch_dry_run_result_persistence_receipt_consumer_enablement_receipt_storage_record_contract_item_count"] != 4 {
		t.Fatalf("expected notification action request-object dispatch dry-run result persistence receipt consumer enablement receipt storage record contract preview to fail closed without evidence: %#v", preview)
	}
}
