package owner

import (
	"os"
	"path/filepath"
	"testing"
)

func TestProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultVisibilityAuditPreviewModelsVisibilityBoundary(t *testing.T) {
	if routeEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultVisibilityAuditRequest != "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-visibility-audit-preview" {
		t.Fatalf("unexpected notification action request-object dispatch dry-run result visibility request constant: %s", routeEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultVisibilityAuditRequest)
	}
	preview, err := NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultVisibilityAuditPreview(projectRootForOwnerTest(t))
	if err != nil {
		t.Fatalf("NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultVisibilityAuditPreview returned error: %v", err)
	}
	if preview["version"] != currentProjectVersion(t) ||
		preview["schema_version"] != "xnix.runtime.production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_kde_safe_redacted_status_closed_record_writer_storage_root_authorization_receipt_dry_run_lookup_result_consumer_projection_route_enablement_lookup_route_dispatch_dry_run_result_notification_action_request_object_dispatch_dry_run_result_visibility_audit.v1" ||
		preview["request_type"] != routeEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultVisibilityAuditRequest ||
		preview["audit_type"] != "receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-visibility-audit" ||
		preview["audit_decision"] != "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-visibility-audit-ready-visibility-persistence-disabled" {
		t.Fatalf("unexpected notification action request-object dispatch dry-run result visibility schema: %#v", preview)
	}
	for _, key := range []string{
		"current_mainline_consumed",
		"route_enablement_lookup_route_dispatch_dry_run_result_notification_action_request_object_dispatch_dry_run_audit_consumed",
		"route_enablement_lookup_route_dispatch_dry_run_result_notification_action_request_object_dispatch_dry_run_audit_ready",
		"route_enablement_lookup_route_dispatch_dry_run_result_notification_action_request_object_dispatch_authorization_audit_consumed",
		"route_enablement_lookup_route_dispatch_dry_run_result_notification_action_request_object_audit_consumed",
		"route_enablement_lookup_route_dispatch_dry_run_result_notification_action_enablement_audit_consumed",
		"route_enablement_lookup_route_dispatch_dry_run_result_notification_delivery_grant_audit_consumed",
		"route_enablement_lookup_route_dispatch_dry_run_result_notification_delivery_execution_authorization_audit_consumed",
		"lookup_route_dispatch_dry_run_result_notification_action_request_object_dispatch_dry_run_result_visibility_audit_required",
		"lookup_route_dispatch_dry_run_result_notification_action_request_object_dispatch_dry_run_result_visibility_modeled",
		"lookup_route_dispatch_dry_run_result_notification_action_request_object_dispatch_dry_run_result_visibility_ready",
		"route_enablement_lookup_route_dispatch_dry_run_result_notification_action_request_object_dispatch_dry_run_result_visibility_ready",
		"notification_action_request_object_dispatch_dry_run_result_visibility_review_only",
		"notification_action_request_object_dispatch_dry_run_review_only_consumed",
		"kde_notification_action_request_object_dispatch_dry_run_result_visibility_modeled",
		"notification_center_action_request_object_dispatch_dry_run_result_visibility_modeled",
		"install_failure_notification_action_request_object_dispatch_dry_run_result_visibility_modeled",
		"repair_suggestion_notification_action_request_object_dispatch_dry_run_result_visibility_modeled",
		"environment_switch_notification_action_request_object_dispatch_dry_run_result_visibility_modeled",
		"approval_info_notification_action_request_object_dispatch_dry_run_result_visibility_modeled",
		"redacted_kde_visibility_modeled",
		"runtime_diagnostics_visibility_modeled",
		"result_visibility_plan_ready",
		"result_visibility_audit_required",
		"result_visibility_audit_modeled",
		"result_visibility_ready",
		"kde_safe_redacted_result_only",
		"raw_result_hidden",
		"runtime_owned",
		"go_runtime_backed",
	} {
		if preview[key] != true {
			t.Fatalf("expected %s to be true in notification action request-object dispatch dry-run result visibility preview: %#v", key, preview)
		}
	}
	for _, key := range []string{
		"user_visible",
		"result_visibility_persistence_enabled",
		"result_visibility_persisted",
		"dry_run_result_persistence_enabled",
		"dry_run_result_persisted",
		"dispatch_dry_run_execution_enabled",
		"dispatch_dry_run_executed",
		"dispatch_authorization_granted",
		"request_object_creation_enabled",
		"request_object_dispatch_enabled",
		"request_object_persistence_enabled",
		"request_object_created",
		"request_objects_created",
		"request_object_dispatched",
		"request_objects_dispatched",
		"request_object_persisted",
		"dispatch_authorization_persisted",
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
		"explicit_user_consent_collected",
		"consent_receipt_created",
		"route_enablement_accepted",
		"lookup_route_enabled",
		"lookup_route_dispatch_authorized",
		"lookup_route_dispatch_callable",
		"storage_write_enabled",
		"status_persistence_write_enabled",
		"redacted_summary_persisted",
		"kde_status_persisted",
		"runtime_diagnostics_persisted",
		"raw_result_exposed",
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
			t.Fatalf("expected %s to remain false in notification action request-object dispatch dry-run result visibility preview: %#v", key, preview)
		}
	}
	if preview["notification_action_request_dispatch_dry_run_result_visibility_item_count"] != 4 ||
		preview["required_notification_action_request_dispatch_dry_run_result_visibility_item_count"] != 4 ||
		preview["ready_notification_action_request_dispatch_dry_run_result_visibility_item_count"] != 4 ||
		preview["missing_notification_action_request_dispatch_dry_run_result_visibility_item_count"] != 0 ||
		preview["notification_action_request_dispatch_dry_run_result_visibility_modeled_item_count"] != 4 ||
		preview["dispatch_dry_run_audit_consumed_item_count"] != 4 ||
		preview["dispatch_authorization_audit_consumed_item_count"] != 4 ||
		preview["persisted_result_visibility_item_count"] != 0 ||
		preview["persisted_dry_run_result_item_count"] != 0 ||
		preview["executed_notification_action_request_dispatch_dry_run_item_count"] != 0 ||
		preview["dispatched_request_object_item_count"] != 0 ||
		preview["created_request_object_item_count"] != 0 ||
		preview["side_effect_notification_action_request_dispatch_dry_run_result_visibility_item_count"] != 0 {
		t.Fatalf("unexpected notification action request-object dispatch dry-run result visibility counts: %#v", preview)
	}
	if !sameStrings(preview["notification_action_request_dispatch_dry_run_result_visibility_item_ids"].([]string), []string{
		"route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-visibility-install-failure",
		"route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-visibility-repair-suggestion",
		"route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-visibility-environment-switch",
		"route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-visibility-approval-info",
	}) {
		t.Fatalf("unexpected notification action request-object dispatch dry-run result visibility item ids: %#v", preview["notification_action_request_dispatch_dry_run_result_visibility_item_ids"])
	}
	for _, item := range preview["notification_action_request_dispatch_dry_run_result_visibility_items"].([]map[string]any) {
		if item["evidence_present"] != true ||
			item["notification_action_request_object_dispatch_dry_run_audit_consumed"] != true ||
			item["result_visibility_required"] != true ||
			item["result_visibility_modeled"] != true ||
			item["result_visibility_ready"] != true ||
			item["result_visibility_persistence_enabled"] != false ||
			item["result_visibility_persisted"] != false ||
			item["user_visible"] != false ||
			item["dispatch_dry_run_execution_enabled"] != false ||
			item["dispatch_dry_run_executed"] != false ||
			item["dry_run_result_persisted"] != false ||
			item["dispatch_authorization_granted"] != false ||
			item["request_object_creation_enabled"] != false ||
			item["request_object_dispatch_enabled"] != false ||
			item["notification_action_enabled"] != false ||
			item["notification_sent"] != false ||
			item["kde_safe_redacted_result_only"] != true ||
			item["raw_result_hidden"] != true ||
			item["side_effects_disabled"] != true ||
			item["host_root_modified"] != false ||
			item["notification_action_request_object_dispatch_dry_run_result_visibility_status"] != "redacted-status-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-visibility-ready-persistence-disabled" {
			t.Fatalf("unsafe notification action request-object dispatch dry-run result visibility item: %#v", item)
		}
	}
	counts := preview["counts"].(map[string]int)
	if counts["total"] != 10 || counts["passed"] != 10 || counts["blocked"] != 0 || counts["pending"] != 0 {
		t.Fatalf("unexpected notification action request-object dispatch dry-run result visibility checks: %#v", counts)
	}
	if err := validateNoBackendTerms(preview, "KDE-safe notification action request-object dispatch dry-run result visibility test"); err != nil {
		t.Fatalf("validateNoBackendTerms returned error: %v", err)
	}
}

func TestProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultVisibilityAuditPreviewFailsClosedWithoutSources(t *testing.T) {
	root, err := os.MkdirTemp("", "xnix-action-request-dispatch-dry-run-result-visibility-*")
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
	if err := os.WriteFile(filepath.Join(root, "docs", "xnix-current-mainline.md"), []byte("missing notification action request-object dispatch dry-run result visibility evidence"), 0o644); err != nil {
		t.Fatalf("WriteFile mainline returned error: %v", err)
	}

	preview, err := NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultVisibilityAuditPreview(root)
	if err != nil {
		t.Fatalf("NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultVisibilityAuditPreview returned error: %v", err)
	}
	if preview["audit_decision"] != "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-visibility-audit-blocked" ||
		preview["current_mainline_consumed"] != false ||
		preview["route_enablement_lookup_route_dispatch_dry_run_result_notification_action_request_object_dispatch_dry_run_audit_consumed"] != false ||
		preview["lookup_route_dispatch_dry_run_result_notification_action_request_object_dispatch_dry_run_result_visibility_ready"] != false ||
		preview["result_visibility_persistence_enabled"] != false ||
		preview["result_visibility_persisted"] != false ||
		preview["user_visible"] != false ||
		preview["dispatch_dry_run_execution_enabled"] != false ||
		preview["dispatch_dry_run_executed"] != false ||
		preview["ready_notification_action_request_dispatch_dry_run_result_visibility_item_count"] != 0 ||
		preview["missing_notification_action_request_dispatch_dry_run_result_visibility_item_count"] != 4 {
		t.Fatalf("expected notification action request-object dispatch dry-run result visibility audit to fail closed without evidence: %#v", preview)
	}
	for _, item := range preview["notification_action_request_dispatch_dry_run_result_visibility_items"].([]map[string]any) {
		if item["evidence_present"] != false ||
			item["result_visibility_ready"] != false ||
			item["result_visibility_persistence_enabled"] != false ||
			item["result_visibility_persisted"] != false ||
			item["user_visible"] != false ||
			item["dispatch_dry_run_execution_enabled"] != false ||
			item["dispatch_dry_run_executed"] != false ||
			item["request_object_creation_enabled"] != false ||
			item["request_object_dispatch_enabled"] != false ||
			item["notification_action_enabled"] != false ||
			item["notification_sent"] != false ||
			item["host_root_modified"] != false ||
			item["notification_action_request_object_dispatch_dry_run_result_visibility_status"] != "missing-redacted-status-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-visibility-evidence" {
			t.Fatalf("unsafe fail-closed notification action request-object dispatch dry-run result visibility item: %#v", item)
		}
	}
}
