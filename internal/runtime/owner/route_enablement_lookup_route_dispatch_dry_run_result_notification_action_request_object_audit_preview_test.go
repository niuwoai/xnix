package owner

import (
	"os"
	"path/filepath"
	"testing"
)

func TestProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectAuditPreviewModelsRequestObjectBoundary(t *testing.T) {
	if routeEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectAuditRequest != "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-audit-preview" {
		t.Fatalf("unexpected notification action request-object request constant: %s", routeEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectAuditRequest)
	}
	preview, err := NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectAuditPreview(projectRootForOwnerTest(t))
	if err != nil {
		t.Fatalf("NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectAuditPreview returned error: %v", err)
	}
	if preview["version"] != currentProjectVersion(t) ||
		preview["schema_version"] != "xnix.runtime.production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_kde_safe_redacted_status_closed_record_writer_storage_root_authorization_receipt_dry_run_lookup_result_consumer_projection_route_enablement_lookup_route_dispatch_dry_run_result_notification_action_request_object_audit.v1" ||
		preview["request_type"] != routeEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectAuditRequest ||
		preview["audit_type"] != "receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-audit" ||
		preview["audit_decision"] != "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-audit-ready-request-objects-disabled" {
		t.Fatalf("unexpected notification action request-object schema: %#v", preview)
	}
	for _, key := range []string{
		"current_mainline_consumed",
		"route_enablement_lookup_route_dispatch_dry_run_result_notification_action_enablement_audit_consumed",
		"route_enablement_lookup_route_dispatch_dry_run_result_notification_action_enablement_audit_ready",
		"route_enablement_lookup_route_dispatch_dry_run_result_notification_delivery_grant_audit_consumed",
		"route_enablement_lookup_route_dispatch_dry_run_result_notification_delivery_grant_audit_ready",
		"route_enablement_lookup_route_dispatch_dry_run_result_notification_delivery_execution_authorization_audit_consumed",
		"route_enablement_lookup_route_dispatch_dry_run_result_notification_delivery_execution_authorization_audit_ready",
		"lookup_route_dispatch_dry_run_result_notification_action_request_object_audit_required",
		"lookup_route_dispatch_dry_run_result_notification_action_request_object_modeled",
		"lookup_route_dispatch_dry_run_result_notification_action_request_object_ready",
		"route_enablement_lookup_route_dispatch_dry_run_result_notification_action_request_object_ready",
		"notification_action_request_object_review_only",
		"notification_action_enablement_review_only_consumed",
		"notification_delivery_grant_review_only_consumed",
		"kde_notification_action_request_object_modeled",
		"notification_center_action_request_object_modeled",
		"install_failure_notification_action_request_object_modeled",
		"repair_suggestion_notification_action_request_object_modeled",
		"environment_switch_notification_action_request_object_modeled",
		"approval_info_notification_action_request_object_modeled",
		"notification_action_required",
		"notification_action_ready",
		"notification_action_enablement_required",
		"notification_action_enablement_ready",
		"action_card_enablement_required",
		"action_card_enablement_ready",
		"request_object_required",
		"request_object_modeled",
		"request_object_ready",
		"notification_action_request_object_required",
		"notification_action_request_object_modeled",
		"notification_action_request_object_ready",
		"request_object_creation_required",
		"request_object_creation_ready",
		"kde_safe_redacted_result_only",
		"raw_result_hidden",
		"runtime_owned",
		"go_runtime_backed",
	} {
		if preview[key] != true {
			t.Fatalf("expected %s to be true in notification action request-object preview: %#v", key, preview)
		}
	}
	for _, key := range []string{
		"request_object_dispatch_authorization_ready",
		"request_object_creation_enabled",
		"request_object_dispatch_enabled",
		"request_object_persistence_enabled",
		"request_object_created",
		"request_objects_created",
		"request_objects_dispatched",
		"portal_request_created",
		"notification_action_enabled",
		"action_card_enabled",
		"notification_delivery_enabled",
		"notification_sent",
		"notification_center_event_triggered",
		"receipt_present",
		"receipt_accepted",
		"receipt_consumed",
		"consumer_authorization_granted",
		"consumer_enabled",
		"route_enablement_accepted",
		"lookup_route_enabled",
		"lookup_route_dispatch_authorized",
		"lookup_route_dispatch_callable",
		"storage_write_enabled",
		"status_persistence_write_enabled",
		"redacted_summary_persisted",
		"kde_status_persisted",
		"runtime_diagnostics_persisted",
		"dry_run_result_persisted",
		"raw_result_exposed",
		"dispatch_dry_run_executed",
		"compatibility_center_opened",
		"support_bundle_exported",
		"support_case_created",
		"production_readiness",
		"production_ownership_ready",
		"system_service_started",
		"session_bus_claimed",
		"production_bus_claimed",
		"write_methods_enabled",
		"runtime_writes_enabled",
		"desktop_files_written",
		"kde_configuration_written",
		"portal_call_executed",
		"adapter_invocation_enabled",
		"backend_launch_enabled",
		"backend_process_started",
		"network_required",
		"host_root_modified",
		"privileged_container_required",
		"state_root_path_exposed",
		"file_paths_exposed",
		"file_content_read",
		"raw_command_exposed",
		"raw_executable_exposed",
		"backend_details_exposed",
		"kde_policy_owner",
	} {
		if preview[key] != false {
			t.Fatalf("expected %s to remain false in notification action request-object preview: %#v", key, preview)
		}
	}
	if preview["notification_action_request_object_item_count"] != 4 ||
		preview["required_notification_action_request_object_item_count"] != 4 ||
		preview["ready_notification_action_request_object_item_count"] != 4 ||
		preview["missing_notification_action_request_object_item_count"] != 0 ||
		preview["action_enablement_audit_consumed_item_count"] != 4 ||
		preview["delivery_grant_audit_consumed_item_count"] != 4 ||
		preview["notification_action_request_object_modeled_item_count"] != 4 ||
		preview["request_object_creation_ready_item_count"] != 4 ||
		preview["created_request_object_item_count"] != 0 ||
		preview["dispatched_request_object_item_count"] != 0 ||
		preview["persisted_request_object_item_count"] != 0 ||
		preview["side_effect_notification_action_request_object_item_count"] != 0 {
		t.Fatalf("unexpected notification action request-object counts: %#v", preview)
	}
	if !sameStrings(preview["notification_action_request_object_item_ids"].([]string), []string{
		"route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-install-failure",
		"route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-repair-suggestion",
		"route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-environment-switch",
		"route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-approval-info",
	}) {
		t.Fatalf("unexpected notification action request-object item ids: %#v", preview["notification_action_request_object_item_ids"])
	}
	for _, item := range preview["notification_action_request_object_items"].([]map[string]any) {
		if item["evidence_present"] != true ||
			item["notification_action_enablement_audit_consumed"] != true ||
			item["notification_delivery_grant_audit_consumed"] != true ||
			item["request_object_required"] != true ||
			item["request_object_ready"] != true ||
			item["notification_action_request_object_required"] != true ||
			item["notification_action_request_object_ready"] != true ||
			item["request_object_creation_required"] != true ||
			item["request_object_creation_ready"] != true ||
			item["request_object_creation_enabled"] != false ||
			item["request_object_dispatch_enabled"] != false ||
			item["request_object_persistence_enabled"] != false ||
			item["request_object_created"] != false ||
			item["request_object_dispatched"] != false ||
			item["request_object_persisted"] != false ||
			item["notification_action_enabled"] != false ||
			item["action_card_enabled"] != false ||
			item["notification_sent"] != false ||
			item["notification_center_event_triggered"] != false ||
			item["request_object_review_only"] != true ||
			item["kde_safe_redacted_result_only"] != true ||
			item["raw_result_hidden"] != true ||
			item["side_effects_disabled"] != true ||
			item["host_root_modified"] != false ||
			item["notification_action_request_object_status"] != "redacted-status-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-ready-request-objects-disabled" {
			t.Fatalf("unsafe notification action request-object item: %#v", item)
		}
	}
	counts := preview["counts"].(map[string]int)
	if counts["total"] != 12 || counts["passed"] != 12 || counts["blocked"] != 0 || counts["pending"] != 0 {
		t.Fatalf("unexpected notification action request-object checks: %#v", counts)
	}
	if err := validateNoBackendTerms(preview, "KDE-safe notification action request-object test"); err != nil {
		t.Fatalf("validateNoBackendTerms returned error: %v", err)
	}
}

func TestProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectAuditPreviewFailsClosedWithoutSources(t *testing.T) {
	root, err := os.MkdirTemp("", "xnix-action-request-*")
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
	if err := os.WriteFile(filepath.Join(root, "docs", "xnix-current-mainline.md"), []byte("missing notification action request-object evidence"), 0o644); err != nil {
		t.Fatalf("WriteFile mainline returned error: %v", err)
	}

	preview, err := NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectAuditPreview(root)
	if err != nil {
		t.Fatalf("NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectAuditPreview returned error: %v", err)
	}
	if preview["audit_decision"] != "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-audit-blocked" ||
		preview["current_mainline_consumed"] != false ||
		preview["route_enablement_lookup_route_dispatch_dry_run_result_notification_action_enablement_audit_consumed"] != false ||
		preview["lookup_route_dispatch_dry_run_result_notification_action_request_object_ready"] != false ||
		preview["ready_notification_action_request_object_item_count"] != 0 ||
		preview["missing_notification_action_request_object_item_count"] != 4 {
		t.Fatalf("expected notification action request-object audit to fail closed without evidence: %#v", preview)
	}
	for _, item := range preview["notification_action_request_object_items"].([]map[string]any) {
		if item["evidence_present"] != false ||
			item["request_object_ready"] != false ||
			item["request_object_creation_enabled"] != false ||
			item["request_object_dispatch_enabled"] != false ||
			item["request_object_created"] != false ||
			item["notification_action_enabled"] != false ||
			item["notification_sent"] != false ||
			item["host_root_modified"] != false {
			t.Fatalf("unsafe fail-closed notification action request-object item: %#v", item)
		}
	}
}
