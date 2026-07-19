package owner

import (
	"os"
	"path/filepath"
	"testing"
)

func TestProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptAuditPreviewModelsReceiptBoundary(t *testing.T) {
	if routeEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptAuditRequest != "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-audit-preview" {
		t.Fatalf("unexpected notification action request-object dispatch dry-run result persistence receipt request constant: %s", routeEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptAuditRequest)
	}
	preview, err := NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptAuditPreview(projectRootForOwnerTest(t))
	if err != nil {
		t.Fatalf("NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptAuditPreview returned error: %v", err)
	}
	if preview["version"] != currentProjectVersion(t) ||
		preview["schema_version"] != "xnix.runtime.production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_kde_safe_redacted_status_closed_record_writer_storage_root_authorization_receipt_dry_run_lookup_result_consumer_projection_route_enablement_lookup_route_dispatch_dry_run_result_notification_action_request_object_dispatch_dry_run_result_persistence_receipt_audit.v1" ||
		preview["request_type"] != routeEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptAuditRequest ||
		preview["audit_type"] != "receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-audit" ||
		preview["audit_decision"] != "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-audit-ready-receipts-disabled" {
		t.Fatalf("unexpected notification action request-object dispatch dry-run result persistence receipt schema: %#v", preview)
	}
	for _, key := range []string{
		"current_mainline_consumed",
		"route_enablement_lookup_route_dispatch_dry_run_result_notification_action_request_object_dispatch_dry_run_result_persistence_authorization_audit_consumed",
		"route_enablement_lookup_route_dispatch_dry_run_result_notification_action_request_object_dispatch_dry_run_result_persistence_authorization_audit_ready",
		"route_enablement_lookup_route_dispatch_dry_run_result_notification_action_request_object_dispatch_dry_run_result_visibility_audit_consumed",
		"route_enablement_lookup_route_dispatch_dry_run_result_notification_action_request_object_dispatch_dry_run_audit_consumed",
		"route_enablement_lookup_route_dispatch_dry_run_result_notification_action_request_object_dispatch_authorization_audit_consumed",
		"lookup_route_dispatch_dry_run_result_notification_action_request_object_dispatch_dry_run_result_persistence_receipt_audit_required",
		"lookup_route_dispatch_dry_run_result_notification_action_request_object_dispatch_dry_run_result_persistence_receipt_modeled",
		"lookup_route_dispatch_dry_run_result_notification_action_request_object_dispatch_dry_run_result_persistence_receipt_ready",
		"route_enablement_lookup_route_dispatch_dry_run_result_notification_action_request_object_dispatch_dry_run_result_persistence_receipt_ready",
		"notification_action_request_object_dispatch_dry_run_result_persistence_receipt_review_only",
		"notification_action_request_object_dispatch_dry_run_result_persistence_authorization_review_only_consumed",
		"result_persistence_authorization_ready",
		"result_persistence_receipt_required",
		"result_persistence_receipt_modeled",
		"result_persistence_receipt_ready",
		"kde_safe_redacted_result_only",
		"raw_result_hidden",
		"runtime_owned",
		"go_runtime_backed",
	} {
		if preview[key] != true {
			t.Fatalf("expected %s to be true in notification action request-object dispatch dry-run result persistence receipt preview: %#v", key, preview)
		}
	}
	for _, key := range []string{
		"result_persistence_authorized",
		"operator_persistence_approval_present",
		"result_persistence_receipt_present",
		"result_persistence_receipt_accepted",
		"result_persistence_receipt_consumed",
		"receipt_write_enabled",
		"receipt_persistence_enabled",
		"receipt_present",
		"receipt_accepted",
		"receipt_consumed",
		"authorization_accepted",
		"user_visible",
		"result_visibility_persistence_enabled",
		"result_visibility_persisted",
		"dry_run_result_persistence_enabled",
		"dry_run_result_persisted",
		"dispatch_dry_run_execution_enabled",
		"dispatch_dry_run_executed",
		"dispatch_authorization_granted",
		"dispatch_authorization_persisted",
		"request_object_creation_enabled",
		"request_object_dispatch_enabled",
		"request_object_persistence_enabled",
		"request_object_created",
		"request_objects_created",
		"request_object_dispatched",
		"request_objects_dispatched",
		"request_object_persisted",
		"request_objects_persisted",
		"portal_request_created",
		"notification_action_enabled",
		"action_card_enabled",
		"notification_delivery_enabled",
		"notification_sent",
		"notification_center_event_triggered",
		"storage_write_enabled",
		"runtime_writes_enabled",
		"production_ownership_ready",
		"production_bus_claimed",
		"backend_launch_enabled",
		"host_root_modified",
		"state_root_path_exposed",
		"file_paths_exposed",
		"file_content_read",
		"raw_command_exposed",
		"raw_executable_exposed",
		"backend_details_exposed",
		"kde_policy_owner",
	} {
		if preview[key] != false {
			t.Fatalf("expected %s to remain false in notification action request-object dispatch dry-run result persistence receipt preview: %#v", key, preview)
		}
	}
	if preview["notification_action_request_dispatch_dry_run_result_persistence_receipt_item_count"] != 4 ||
		preview["ready_notification_action_request_dispatch_dry_run_result_persistence_receipt_item_count"] != 4 ||
		preview["missing_notification_action_request_dispatch_dry_run_result_persistence_receipt_item_count"] != 0 ||
		preview["persistence_authorization_audit_consumed_item_count"] != 4 ||
		preview["present_result_persistence_receipt_item_count"] != 0 ||
		preview["accepted_result_persistence_receipt_item_count"] != 0 ||
		preview["persisted_dry_run_result_item_count"] != 0 ||
		preview["side_effect_notification_action_request_dispatch_dry_run_result_persistence_receipt_item_count"] != 0 {
		t.Fatalf("unexpected notification action request-object dispatch dry-run result persistence receipt counts: %#v", preview)
	}
	if !sameStrings(preview["notification_action_request_dispatch_dry_run_result_persistence_receipt_item_ids"].([]string), []string{
		"route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-install-failure",
		"route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-repair-suggestion",
		"route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-environment-switch",
		"route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-approval-info",
	}) {
		t.Fatalf("unexpected notification action request-object dispatch dry-run result persistence receipt item ids: %#v", preview["notification_action_request_dispatch_dry_run_result_persistence_receipt_item_ids"])
	}
	for _, item := range preview["notification_action_request_dispatch_dry_run_result_persistence_receipt_items"].([]map[string]any) {
		if item["evidence_present"] != true ||
			item["notification_action_request_object_dispatch_dry_run_result_persistence_authorization_audit_consumed"] != true ||
			item["result_persistence_receipt_required"] != true ||
			item["result_persistence_receipt_modeled"] != true ||
			item["result_persistence_receipt_ready"] != true ||
			item["result_persistence_receipt_present"] != false ||
			item["result_persistence_receipt_accepted"] != false ||
			item["receipt_write_enabled"] != false ||
			item["receipt_persistence_enabled"] != false ||
			item["result_persistence_authorized"] != false ||
			item["dry_run_result_persistence_enabled"] != false ||
			item["dry_run_result_persisted"] != false ||
			item["user_visible"] != false ||
			item["dispatch_dry_run_execution_enabled"] != false ||
			item["dispatch_dry_run_executed"] != false ||
			item["request_object_creation_enabled"] != false ||
			item["request_object_dispatch_enabled"] != false ||
			item["notification_action_enabled"] != false ||
			item["notification_sent"] != false ||
			item["kde_safe_redacted_result_only"] != true ||
			item["raw_result_hidden"] != true ||
			item["side_effects_disabled"] != true ||
			item["host_root_modified"] != false ||
			item["notification_action_request_object_dispatch_dry_run_result_persistence_receipt_status"] != "redacted-status-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-ready-receipts-disabled" {
			t.Fatalf("unsafe notification action request-object dispatch dry-run result persistence receipt item: %#v", item)
		}
	}
	counts := preview["counts"].(map[string]int)
	if counts["total"] != 10 || counts["passed"] != 10 || counts["blocked"] != 0 || counts["pending"] != 0 {
		t.Fatalf("unexpected notification action request-object dispatch dry-run result persistence receipt checks: %#v", counts)
	}
	if err := validateNoBackendTerms(preview, "KDE-safe notification action request-object dispatch dry-run result persistence receipt test"); err != nil {
		t.Fatalf("validateNoBackendTerms returned error: %v", err)
	}
}

func TestProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptAuditPreviewFailsClosedWithoutSources(t *testing.T) {
	root, err := os.MkdirTemp("", "xnix-action-request-dispatch-dry-run-result-persistence-receipt-*")
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
	if err := os.WriteFile(filepath.Join(root, "docs", "xnix-current-mainline.md"), []byte("missing notification action request-object dispatch dry-run result persistence receipt evidence"), 0o644); err != nil {
		t.Fatalf("WriteFile mainline returned error: %v", err)
	}

	preview, err := NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptAuditPreview(root)
	if err != nil {
		t.Fatalf("NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptAuditPreview returned error: %v", err)
	}
	if preview["audit_decision"] != "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-audit-blocked" ||
		preview["current_mainline_consumed"] != false ||
		preview["route_enablement_lookup_route_dispatch_dry_run_result_notification_action_request_object_dispatch_dry_run_result_persistence_authorization_audit_consumed"] != false ||
		preview["lookup_route_dispatch_dry_run_result_notification_action_request_object_dispatch_dry_run_result_persistence_receipt_ready"] != false ||
		preview["result_persistence_receipt_present"] != false ||
		preview["receipt_write_enabled"] != false ||
		preview["ready_notification_action_request_dispatch_dry_run_result_persistence_receipt_item_count"] != 0 ||
		preview["missing_notification_action_request_dispatch_dry_run_result_persistence_receipt_item_count"] != 4 {
		t.Fatalf("expected notification action request-object dispatch dry-run result persistence receipt audit to fail closed without evidence: %#v", preview)
	}
	for _, item := range preview["notification_action_request_dispatch_dry_run_result_persistence_receipt_items"].([]map[string]any) {
		if item["evidence_present"] != false ||
			item["result_persistence_receipt_ready"] != false ||
			item["result_persistence_receipt_present"] != false ||
			item["result_persistence_receipt_accepted"] != false ||
			item["receipt_write_enabled"] != false ||
			item["dry_run_result_persistence_enabled"] != false ||
			item["dry_run_result_persisted"] != false ||
			item["dispatch_dry_run_execution_enabled"] != false ||
			item["request_object_creation_enabled"] != false ||
			item["notification_action_enabled"] != false ||
			item["notification_sent"] != false ||
			item["host_root_modified"] != false ||
			item["notification_action_request_object_dispatch_dry_run_result_persistence_receipt_status"] != "missing-redacted-status-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-evidence" {
			t.Fatalf("unsafe fail-closed notification action request-object dispatch dry-run result persistence receipt item: %#v", item)
		}
	}
}
