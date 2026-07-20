package owner

import (
	"os"
	"path/filepath"
	"testing"
)

func TestProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptAuditPreviewModelsEnablementReceiptBoundary(t *testing.T) {
	if routeEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptAuditRequest != "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-audit-preview" {
		t.Fatalf("unexpected notification action request-object dispatch dry-run result persistence receipt consumer enablement receipt request constant: %s", routeEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptAuditRequest)
	}
	preview, err := NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptAuditPreview(projectRootForOwnerTest(t))
	if err != nil {
		t.Fatalf("NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptAuditPreview returned error: %v", err)
	}
	if preview["version"] != currentProjectVersion(t) ||
		preview["schema_version"] != "xnix.runtime.production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_kde_safe_redacted_status_closed_record_writer_storage_root_authorization_receipt_dry_run_lookup_result_consumer_projection_route_enablement_lookup_route_dispatch_dry_run_result_notification_action_request_object_dispatch_dry_run_result_persistence_receipt_consumer_enablement_receipt_audit.v1" ||
		preview["request_type"] != routeEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptAuditRequest ||
		preview["audit_decision"] != "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-audit-ready-receipts-disabled" {
		t.Fatalf("unexpected notification action request-object dispatch dry-run result persistence receipt consumer enablement receipt schema: %#v", preview)
	}
	for _, key := range []string{
		"current_mainline_consumed",
		"route_enablement_lookup_route_dispatch_dry_run_result_notification_action_request_object_dispatch_dry_run_result_persistence_receipt_consumer_enablement_gate_audit_consumed",
		"route_enablement_lookup_route_dispatch_dry_run_result_notification_action_request_object_dispatch_dry_run_result_persistence_receipt_consumer_enablement_gate_audit_ready",
		"route_enablement_lookup_route_dispatch_dry_run_result_notification_action_request_object_dispatch_dry_run_result_persistence_receipt_consumer_authorization_audit_ready",
		"lookup_route_dispatch_dry_run_result_notification_action_request_object_dispatch_dry_run_result_persistence_receipt_consumer_enablement_receipt_audit_required",
		"lookup_route_dispatch_dry_run_result_notification_action_request_object_dispatch_dry_run_result_persistence_receipt_consumer_enablement_receipt_modeled",
		"lookup_route_dispatch_dry_run_result_notification_action_request_object_dispatch_dry_run_result_persistence_receipt_consumer_enablement_receipt_ready",
		"notification_action_request_object_dispatch_dry_run_result_persistence_receipt_consumer_enablement_receipt_review_only",
		"notification_action_request_object_dispatch_dry_run_result_persistence_receipt_consumer_enablement_gate_review_only_consumed",
		"result_persistence_receipt_consumer_enablement_ready",
		"result_persistence_receipt_consumer_enablement_receipt_required",
		"result_persistence_receipt_consumer_enablement_receipt_modeled",
		"result_persistence_receipt_consumer_enablement_receipt_ready",
		"kde_safe_redacted_result_only",
		"raw_result_hidden",
		"runtime_owned",
		"go_runtime_backed",
	} {
		if preview[key] != true {
			t.Fatalf("expected %s to be true in notification action request-object dispatch dry-run result persistence receipt consumer enablement receipt preview: %#v", key, preview)
		}
	}
	for _, key := range []string{
		"receipt_consumer_enablement_receipt_present",
		"receipt_consumer_enablement_receipt_written",
		"receipt_consumer_enablement_receipt_accepted",
		"receipt_consumer_enablement_receipt_consumed",
		"receipt_consumer_enablement_receipt_write_enabled",
		"receipt_consumer_authorization_enabled",
		"receipt_consumer_authorized",
		"consumer_authorization_granted",
		"receipt_consumer_enablement_enabled",
		"receipt_consumer_enabled",
		"consumer_enabled",
		"kde_consumer_enabled",
		"runtime_consumer_enabled",
		"receipt_consumption_enabled",
		"receipt_consumed",
		"receipt_acceptance_enabled",
		"receipt_accepted",
		"receipt_write_enabled",
		"dry_run_result_persistence_enabled",
		"dry_run_result_persisted",
		"dispatch_dry_run_execution_enabled",
		"dispatch_dry_run_executed",
		"request_object_creation_enabled",
		"request_object_dispatch_enabled",
		"notification_action_enabled",
		"notification_sent",
		"user_visible",
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
			t.Fatalf("expected %s to remain false in notification action request-object dispatch dry-run result persistence receipt consumer enablement receipt preview: %#v", key, preview)
		}
	}
	if preview["notification_action_request_dispatch_dry_run_result_persistence_receipt_consumer_enablement_receipt_item_count"] != 4 ||
		preview["ready_notification_action_request_dispatch_dry_run_result_persistence_receipt_consumer_enablement_receipt_item_count"] != 4 ||
		preview["missing_notification_action_request_dispatch_dry_run_result_persistence_receipt_consumer_enablement_receipt_item_count"] != 0 ||
		preview["persistence_receipt_consumer_enablement_gate_audit_consumed_item_count"] != 4 ||
		preview["persistence_receipt_consumer_enablement_receipt_modeled_item_count"] != 4 ||
		preview["present_result_persistence_receipt_consumer_enablement_receipt_item_count"] != 0 ||
		preview["accepted_result_persistence_receipt_consumer_enablement_receipt_item_count"] != 0 ||
		preview["consumed_result_persistence_receipt_consumer_enablement_receipt_item_count"] != 0 ||
		preview["side_effect_notification_action_request_dispatch_dry_run_result_persistence_receipt_consumer_enablement_receipt_item_count"] != 0 {
		t.Fatalf("unexpected notification action request-object dispatch dry-run result persistence receipt consumer enablement receipt counts: %#v", preview)
	}
	if !sameStrings(preview["notification_action_request_dispatch_dry_run_result_persistence_receipt_consumer_enablement_receipt_item_ids"].([]string), []string{
		"route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-install-failure",
		"route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-repair-suggestion",
		"route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-environment-switch",
		"route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-approval-info",
	}) {
		t.Fatalf("unexpected notification action request-object dispatch dry-run result persistence receipt consumer enablement receipt item ids: %#v", preview["notification_action_request_dispatch_dry_run_result_persistence_receipt_consumer_enablement_receipt_item_ids"])
	}
	for _, item := range preview["notification_action_request_dispatch_dry_run_result_persistence_receipt_consumer_enablement_receipt_items"].([]map[string]any) {
		if item["evidence_present"] != true ||
			item["persistence_receipt_consumer_enablement_gate_audit_consumed"] != true ||
			item["result_persistence_receipt_consumer_enablement_receipt_ready"] != true ||
			item["receipt_consumer_enablement_receipt_present"] != false ||
			item["receipt_consumer_enablement_receipt_written"] != false ||
			item["receipt_consumer_enablement_receipt_write_enabled"] != false ||
			item["receipt_consumer_enablement_enabled"] != false ||
			item["consumer_enabled"] != false ||
			item["kde_consumer_enabled"] != false ||
			item["runtime_consumer_enabled"] != false ||
			item["receipt_consumed"] != false ||
			item["dry_run_result_persistence_enabled"] != false ||
			item["dispatch_dry_run_execution_enabled"] != false ||
			item["request_object_creation_enabled"] != false ||
			item["notification_action_enabled"] != false ||
			item["notification_sent"] != false ||
			item["side_effects_disabled"] != true ||
			item["host_root_modified"] != false ||
			item["notification_action_request_object_dispatch_dry_run_result_persistence_receipt_consumer_enablement_receipt_status"] != "redacted-status-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-ready-receipts-disabled" {
			t.Fatalf("unsafe notification action request-object dispatch dry-run result persistence receipt consumer enablement receipt item: %#v", item)
		}
	}
	counts := preview["counts"].(map[string]int)
	if counts["total"] != 7 || counts["passed"] != 7 || counts["blocked"] != 0 || counts["pending"] != 0 {
		t.Fatalf("unexpected notification action request-object dispatch dry-run result persistence receipt consumer enablement receipt checks: %#v", counts)
	}
	if err := validateNoBackendTerms(preview, "KDE-safe notification action request-object dispatch dry-run result persistence receipt consumer enablement receipt test"); err != nil {
		t.Fatalf("validateNoBackendTerms returned error: %v", err)
	}
}

func TestProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptAuditPreviewFailsClosedWithoutSources(t *testing.T) {
	root, err := os.MkdirTemp("", "xnix-action-request-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-*")
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
	if err := os.WriteFile(filepath.Join(root, "docs", "xnix-current-mainline.md"), []byte("missing notification action request-object dispatch dry-run result persistence receipt consumer enablement receipt evidence"), 0o644); err != nil {
		t.Fatalf("WriteFile mainline returned error: %v", err)
	}

	preview, err := NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptAuditPreview(root)
	if err != nil {
		t.Fatalf("NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptAuditPreview returned error: %v", err)
	}
	if preview["audit_decision"] != "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-audit-blocked" ||
		preview["current_mainline_consumed"] != false ||
		preview["route_enablement_lookup_route_dispatch_dry_run_result_notification_action_request_object_dispatch_dry_run_result_persistence_receipt_consumer_enablement_gate_audit_consumed"] != false ||
		preview["lookup_route_dispatch_dry_run_result_notification_action_request_object_dispatch_dry_run_result_persistence_receipt_consumer_enablement_receipt_ready"] != false ||
		preview["receipt_consumer_enablement_receipt_present"] != false ||
		preview["receipt_consumer_enablement_receipt_write_enabled"] != false ||
		preview["ready_notification_action_request_dispatch_dry_run_result_persistence_receipt_consumer_enablement_receipt_item_count"] != 0 ||
		preview["missing_notification_action_request_dispatch_dry_run_result_persistence_receipt_consumer_enablement_receipt_item_count"] != 4 {
		t.Fatalf("expected notification action request-object dispatch dry-run result persistence receipt consumer enablement receipt audit to fail closed without evidence: %#v", preview)
	}
	for _, item := range preview["notification_action_request_dispatch_dry_run_result_persistence_receipt_consumer_enablement_receipt_items"].([]map[string]any) {
		if item["evidence_present"] != false ||
			item["result_persistence_receipt_consumer_enablement_receipt_ready"] != false ||
			item["receipt_consumer_enablement_receipt_present"] != false ||
			item["receipt_consumer_enablement_receipt_write_enabled"] != false ||
			item["consumer_enabled"] != false ||
			item["kde_consumer_enabled"] != false ||
			item["runtime_consumer_enabled"] != false ||
			item["receipt_consumed"] != false ||
			item["dry_run_result_persistence_enabled"] != false ||
			item["dispatch_dry_run_execution_enabled"] != false ||
			item["request_object_creation_enabled"] != false ||
			item["notification_action_enabled"] != false ||
			item["notification_sent"] != false ||
			item["host_root_modified"] != false ||
			item["notification_action_request_object_dispatch_dry_run_result_persistence_receipt_consumer_enablement_receipt_status"] != "missing-redacted-status-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-evidence" {
			t.Fatalf("unsafe fail-closed notification action request-object dispatch dry-run result persistence receipt consumer enablement receipt item: %#v", item)
		}
	}
}
