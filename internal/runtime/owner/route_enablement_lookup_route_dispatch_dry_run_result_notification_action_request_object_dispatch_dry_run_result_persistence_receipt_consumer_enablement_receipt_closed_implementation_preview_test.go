package owner

import (
	"os"
	"path/filepath"
	"testing"
)

func TestProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptClosedImplementationPreviewModelsClosedImplementationBoundary(t *testing.T) {
	preview, err := NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptClosedImplementationPreview(projectRootForOwnerTest(t))
	if err != nil {
		t.Fatalf("NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptClosedImplementationPreview returned error: %v", err)
	}
	if preview["version"] != currentProjectVersion(t) ||
		preview["schema_version"] != "xnix.runtime.production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_kde_safe_redacted_status_closed_record_writer_storage_root_authorization_receipt_dry_run_lookup_result_consumer_projection_route_enablement_lookup_route_dispatch_dry_run_result_notification_action_request_object_dispatch_dry_run_result_persistence_receipt_consumer_enablement_receipt_closed_implementation.v1" ||
		preview["request_type"] != routeEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptClosedImplementationRequest ||
		preview["preview_decision"] != "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-closed-implementation-ready-receipt-disabled" {
		t.Fatalf("unexpected notification action request-object dispatch dry-run result persistence receipt consumer enablement receipt closed implementation schema: %#v", preview)
	}
	for _, key := range []string{
		"current_mainline_consumed",
		"route_enablement_lookup_route_dispatch_dry_run_result_notification_action_request_object_dispatch_dry_run_result_persistence_receipt_consumer_enablement_receipt_enablement_audit_consumed",
		"route_enablement_lookup_route_dispatch_dry_run_result_notification_action_request_object_dispatch_dry_run_result_persistence_receipt_consumer_enablement_receipt_enablement_audit_ready",
		"lookup_route_dispatch_dry_run_result_notification_action_request_object_dispatch_dry_run_result_persistence_receipt_consumer_enablement_receipt_enablement_ready",
		"lookup_route_dispatch_dry_run_result_notification_action_request_object_dispatch_dry_run_result_persistence_receipt_consumer_enablement_receipt_closed_implementation_required",
		"lookup_route_dispatch_dry_run_result_notification_action_request_object_dispatch_dry_run_result_persistence_receipt_consumer_enablement_receipt_closed_implementation_modeled",
		"lookup_route_dispatch_dry_run_result_notification_action_request_object_dispatch_dry_run_result_persistence_receipt_consumer_enablement_receipt_closed_implementation_ready",
		"closed_receipt_shape_modeled",
		"receipt_input_boundary_modeled",
		"receipt_output_boundary_modeled",
		"notification_action_request_object_dispatch_dry_run_result_persistence_receipt_consumer_enablement_receipt_closed_implementation_review_only",
		"result_persistence_receipt_consumer_enablement_receipt_enablement_ready",
		"result_persistence_receipt_consumer_enablement_receipt_closed_implementation_required",
		"result_persistence_receipt_consumer_enablement_receipt_closed_implementation_modeled",
		"result_persistence_receipt_consumer_enablement_receipt_closed_implementation_ready",
		"kde_safe_redacted_result_only",
		"raw_result_hidden",
		"runtime_owned",
		"go_runtime_backed",
	} {
		if preview[key] != true {
			t.Fatalf("expected %s to be true in notification action request-object dispatch dry-run result persistence receipt consumer enablement receipt closed implementation preview: %#v", key, preview)
		}
	}
	for _, key := range []string{
		"receipt_consumer_enablement_receipt_callable",
		"receipt_consumer_enablement_receipt_implementation_enabled",
		"receipt_consumer_enablement_receipt_enablement_enabled",
		"receipt_consumer_enablement_receipt_enabled",
		"receipt_consumer_enablement_receipt_grant_enabled",
		"receipt_consumer_enablement_receipt_granted",
		"receipt_consumer_enablement_receipt_accepted_receipt_gate_enabled",
		"receipt_consumer_enablement_receipt_acceptance_gate_enabled",
		"receipt_consumer_enablement_receipt_persist_authorized",
		"receipt_consumer_enablement_receipt_persistence_enabled",
		"receipt_consumer_enablement_receipt_persisted",
		"receipt_consumer_enablement_receipt_present",
		"receipt_consumer_enablement_receipt_written",
		"receipt_consumer_enablement_receipt_write_enabled",
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
		"dispatch_dry_run_execution_enabled",
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
			t.Fatalf("expected %s to remain false in notification action request-object dispatch dry-run result persistence receipt consumer enablement receipt closed implementation preview: %#v", key, preview)
		}
	}
	if preview["notification_action_request_dispatch_dry_run_result_persistence_receipt_consumer_enablement_receipt_closed_implementation_item_count"] != 4 ||
		preview["ready_notification_action_request_dispatch_dry_run_result_persistence_receipt_consumer_enablement_receipt_closed_implementation_item_count"] != 4 ||
		preview["missing_notification_action_request_dispatch_dry_run_result_persistence_receipt_consumer_enablement_receipt_closed_implementation_item_count"] != 0 ||
		preview["persistence_receipt_consumer_enablement_receipt_enablement_audit_consumed_item_count"] != 4 ||
		preview["persistence_receipt_consumer_enablement_receipt_closed_implementation_modeled_item_count"] != 4 ||
		preview["callable_result_persistence_receipt_consumer_enablement_receipt_item_count"] != 0 ||
		preview["enabled_result_persistence_receipt_consumer_enablement_receipt_item_count"] != 0 ||
		preview["side_effect_notification_action_request_dispatch_dry_run_result_persistence_receipt_consumer_enablement_receipt_closed_implementation_item_count"] != 0 {
		t.Fatalf("unexpected notification action request-object dispatch dry-run result persistence receipt consumer enablement receipt closed implementation counts: %#v", preview)
	}
	for _, item := range preview["notification_action_request_dispatch_dry_run_result_persistence_receipt_consumer_enablement_receipt_closed_implementation_items"].([]map[string]any) {
		if item["evidence_present"] != true ||
			item["persistence_receipt_consumer_enablement_receipt_enablement_audit_consumed"] != true ||
			item["result_persistence_receipt_consumer_enablement_receipt_closed_implementation_ready"] != true ||
			item["closed_receipt_shape_modeled"] != true ||
			item["receipt_input_boundary_modeled"] != true ||
			item["receipt_output_boundary_modeled"] != true ||
			item["receipt_consumer_enablement_receipt_callable"] != false ||
			item["receipt_consumer_enablement_receipt_implementation_enabled"] != false ||
			item["receipt_consumer_enablement_receipt_enablement_enabled"] != false ||
			item["receipt_consumer_enablement_receipt_enabled"] != false ||
			item["receipt_consumer_enablement_receipt_granted"] != false ||
			item["receipt_consumer_enablement_receipt_accepted"] != false ||
			item["receipt_consumer_enablement_receipt_persisted"] != false ||
			item["receipt_consumer_enablement_enabled"] != false ||
			item["consumer_enabled"] != false ||
			item["dry_run_result_persistence_enabled"] != false ||
			item["dispatch_dry_run_execution_enabled"] != false ||
			item["request_object_creation_enabled"] != false ||
			item["notification_action_enabled"] != false ||
			item["notification_sent"] != false ||
			item["side_effects_disabled"] != true ||
			item["host_root_modified"] != false ||
			item["notification_action_request_object_dispatch_dry_run_result_persistence_receipt_consumer_enablement_receipt_closed_implementation_status"] != "redacted-status-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-closed-implementation-ready-receipt-disabled" {
			t.Fatalf("unsafe notification action request-object dispatch dry-run result persistence receipt consumer enablement receipt closed implementation item: %#v", item)
		}
	}
	counts := preview["counts"].(map[string]int)
	if counts["total"] != 7 || counts["passed"] != 7 || counts["blocked"] != 0 || counts["pending"] != 0 {
		t.Fatalf("unexpected notification action request-object dispatch dry-run result persistence receipt consumer enablement receipt closed implementation checks: %#v", counts)
	}
	if err := validateNoBackendTerms(preview, "KDE-safe notification action request-object dispatch dry-run result persistence receipt consumer enablement receipt closed implementation test"); err != nil {
		t.Fatalf("validateNoBackendTerms returned error: %v", err)
	}
}

func TestProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptClosedImplementationPreviewFailsClosedWithoutSources(t *testing.T) {
	root, err := os.MkdirTemp("", "xnix-action-request-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-closed-implementation-*")
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
	if err := os.WriteFile(filepath.Join(root, "docs", "xnix-current-mainline.md"), []byte("missing notification action request-object dispatch dry-run result persistence receipt consumer enablement receipt closed implementation evidence"), 0o644); err != nil {
		t.Fatalf("WriteFile mainline returned error: %v", err)
	}

	preview, err := NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptClosedImplementationPreview(root)
	if err != nil {
		t.Fatalf("NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptClosedImplementationPreview returned error: %v", err)
	}
	if preview["preview_decision"] != "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-closed-implementation-blocked" ||
		preview["current_mainline_consumed"] != false ||
		preview["route_enablement_lookup_route_dispatch_dry_run_result_notification_action_request_object_dispatch_dry_run_result_persistence_receipt_consumer_enablement_receipt_enablement_audit_consumed"] != false ||
		preview["lookup_route_dispatch_dry_run_result_notification_action_request_object_dispatch_dry_run_result_persistence_receipt_consumer_enablement_receipt_closed_implementation_ready"] != false ||
		preview["receipt_consumer_enablement_receipt_callable"] != false ||
		preview["receipt_consumer_enablement_receipt_implementation_enabled"] != false ||
		preview["ready_notification_action_request_dispatch_dry_run_result_persistence_receipt_consumer_enablement_receipt_closed_implementation_item_count"] != 0 ||
		preview["missing_notification_action_request_dispatch_dry_run_result_persistence_receipt_consumer_enablement_receipt_closed_implementation_item_count"] != 4 {
		t.Fatalf("expected notification action request-object dispatch dry-run result persistence receipt consumer enablement receipt closed implementation preview to fail closed without evidence: %#v", preview)
	}
}
