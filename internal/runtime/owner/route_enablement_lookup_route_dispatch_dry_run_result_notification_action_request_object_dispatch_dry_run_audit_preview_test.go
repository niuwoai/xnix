package owner

import (
	"os"
	"path/filepath"
	"testing"
)

func TestProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunAuditPreviewModelsDryRunBoundary(t *testing.T) {
	if routeEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunAuditRequest != "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-audit-preview" {
		t.Fatalf("unexpected notification action request-object dispatch dry-run request constant: %s", routeEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunAuditRequest)
	}
	preview, err := NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunAuditPreview(projectRootForOwnerTest(t))
	if err != nil {
		t.Fatalf("NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunAuditPreview returned error: %v", err)
	}
	if preview["version"] != currentProjectVersion(t) ||
		preview["schema_version"] != "xnix.runtime.production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_kde_safe_redacted_status_closed_record_writer_storage_root_authorization_receipt_dry_run_lookup_result_consumer_projection_route_enablement_lookup_route_dispatch_dry_run_result_notification_action_request_object_dispatch_dry_run_audit.v1" ||
		preview["request_type"] != routeEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunAuditRequest ||
		preview["audit_type"] != "receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-audit" ||
		preview["audit_decision"] != "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-audit-ready-execution-disabled" {
		t.Fatalf("unexpected notification action request-object dispatch dry-run schema: %#v", preview)
	}
	for _, key := range []string{
		"current_mainline_consumed",
		"route_enablement_lookup_route_dispatch_dry_run_result_notification_action_request_object_dispatch_authorization_audit_consumed",
		"route_enablement_lookup_route_dispatch_dry_run_result_notification_action_request_object_dispatch_authorization_audit_ready",
		"route_enablement_lookup_route_dispatch_dry_run_result_notification_action_request_object_audit_consumed",
		"route_enablement_lookup_route_dispatch_dry_run_result_notification_action_request_object_audit_ready",
		"route_enablement_lookup_route_dispatch_dry_run_result_notification_action_enablement_audit_consumed",
		"route_enablement_lookup_route_dispatch_dry_run_result_notification_delivery_grant_audit_consumed",
		"route_enablement_lookup_route_dispatch_dry_run_result_notification_delivery_execution_authorization_audit_consumed",
		"lookup_route_dispatch_dry_run_result_notification_action_request_object_dispatch_dry_run_audit_required",
		"lookup_route_dispatch_dry_run_result_notification_action_request_object_dispatch_dry_run_modeled",
		"lookup_route_dispatch_dry_run_result_notification_action_request_object_dispatch_dry_run_ready",
		"route_enablement_lookup_route_dispatch_dry_run_result_notification_action_request_object_dispatch_dry_run_ready",
		"notification_action_request_object_dispatch_dry_run_review_only",
		"notification_action_request_object_dispatch_authorization_review_only_consumed",
		"notification_action_request_object_review_only_consumed",
		"kde_notification_action_request_object_dispatch_dry_run_modeled",
		"notification_center_action_request_object_dispatch_dry_run_modeled",
		"install_failure_notification_action_request_object_dispatch_dry_run_modeled",
		"repair_suggestion_notification_action_request_object_dispatch_dry_run_modeled",
		"environment_switch_notification_action_request_object_dispatch_dry_run_modeled",
		"approval_info_notification_action_request_object_dispatch_dry_run_modeled",
		"request_object_dispatch_authorization_required",
		"request_object_dispatch_authorization_modeled",
		"request_object_dispatch_authorization_ready",
		"request_object_dispatch_dry_run_required",
		"request_object_dispatch_dry_run_modeled",
		"request_object_dispatch_dry_run_ready",
		"dispatch_dry_run_required",
		"dispatch_dry_run_modeled",
		"dispatch_dry_run_plan_ready",
		"kde_safe_redacted_result_only",
		"raw_result_hidden",
		"runtime_owned",
		"go_runtime_backed",
	} {
		if preview[key] != true {
			t.Fatalf("expected %s to be true in notification action request-object dispatch dry-run preview: %#v", key, preview)
		}
	}
	for _, key := range []string{
		"notification_action_request_dispatch_authorized",
		"dispatch_authorization_granted",
		"dispatch_dry_run_execution_enabled",
		"dispatch_dry_run_executed",
		"dry_run_result_persisted",
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
			t.Fatalf("expected %s to remain false in notification action request-object dispatch dry-run preview: %#v", key, preview)
		}
	}
	if preview["notification_action_request_dispatch_dry_run_item_count"] != 4 ||
		preview["required_notification_action_request_dispatch_dry_run_item_count"] != 4 ||
		preview["ready_notification_action_request_dispatch_dry_run_item_count"] != 4 ||
		preview["missing_notification_action_request_dispatch_dry_run_item_count"] != 0 ||
		preview["notification_action_request_dispatch_dry_run_modeled_item_count"] != 4 ||
		preview["dispatch_authorization_audit_consumed_item_count"] != 4 ||
		preview["request_object_audit_consumed_item_count"] != 4 ||
		preview["executed_notification_action_request_dispatch_dry_run_item_count"] != 0 ||
		preview["dispatched_request_object_item_count"] != 0 ||
		preview["created_request_object_item_count"] != 0 ||
		preview["persisted_request_object_item_count"] != 0 ||
		preview["persisted_dry_run_result_item_count"] != 0 ||
		preview["side_effect_notification_action_request_dispatch_dry_run_item_count"] != 0 {
		t.Fatalf("unexpected notification action request-object dispatch dry-run counts: %#v", preview)
	}
	if !sameStrings(preview["notification_action_request_dispatch_dry_run_item_ids"].([]string), []string{
		"route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-install-failure",
		"route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-repair-suggestion",
		"route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-environment-switch",
		"route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-approval-info",
	}) {
		t.Fatalf("unexpected notification action request-object dispatch dry-run item ids: %#v", preview["notification_action_request_dispatch_dry_run_item_ids"])
	}
	for _, item := range preview["notification_action_request_dispatch_dry_run_items"].([]map[string]any) {
		if item["evidence_present"] != true ||
			item["notification_action_request_object_dispatch_authorization_audit_consumed"] != true ||
			item["request_object_dispatch_dry_run_required"] != true ||
			item["request_object_dispatch_dry_run_modeled"] != true ||
			item["request_object_dispatch_dry_run_ready"] != true ||
			item["dispatch_dry_run_required"] != true ||
			item["dispatch_dry_run_modeled"] != true ||
			item["dispatch_dry_run_plan_ready"] != true ||
			item["dispatch_dry_run_execution_enabled"] != false ||
			item["dispatch_dry_run_executed"] != false ||
			item["dry_run_result_persisted"] != false ||
			item["dispatch_authorization_granted"] != false ||
			item["request_object_creation_enabled"] != false ||
			item["request_object_dispatch_enabled"] != false ||
			item["request_object_created"] != false ||
			item["request_object_dispatched"] != false ||
			item["notification_action_enabled"] != false ||
			item["notification_sent"] != false ||
			item["request_object_dispatch_dry_run_review_only"] != true ||
			item["side_effects_disabled"] != true ||
			item["host_root_modified"] != false ||
			item["notification_action_request_object_dispatch_dry_run_status"] != "redacted-status-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-ready-execution-disabled" {
			t.Fatalf("unsafe notification action request-object dispatch dry-run item: %#v", item)
		}
	}
	counts := preview["counts"].(map[string]int)
	if counts["total"] != 11 || counts["passed"] != 11 || counts["blocked"] != 0 || counts["pending"] != 0 {
		t.Fatalf("unexpected notification action request-object dispatch dry-run checks: %#v", counts)
	}
	if err := validateNoBackendTerms(preview, "KDE-safe notification action request-object dispatch dry-run test"); err != nil {
		t.Fatalf("validateNoBackendTerms returned error: %v", err)
	}
}

func TestProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunAuditPreviewFailsClosedWithoutSources(t *testing.T) {
	root, err := os.MkdirTemp("", "xnix-action-request-dispatch-dry-run-*")
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
	if err := os.WriteFile(filepath.Join(root, "docs", "xnix-current-mainline.md"), []byte("missing notification action request-object dispatch dry-run evidence"), 0o644); err != nil {
		t.Fatalf("WriteFile mainline returned error: %v", err)
	}

	preview, err := NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunAuditPreview(root)
	if err != nil {
		t.Fatalf("NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunAuditPreview returned error: %v", err)
	}
	if preview["audit_decision"] != "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-audit-blocked" ||
		preview["current_mainline_consumed"] != false ||
		preview["route_enablement_lookup_route_dispatch_dry_run_result_notification_action_request_object_dispatch_authorization_audit_consumed"] != false ||
		preview["lookup_route_dispatch_dry_run_result_notification_action_request_object_dispatch_dry_run_ready"] != false ||
		preview["dispatch_dry_run_execution_enabled"] != false ||
		preview["dispatch_dry_run_executed"] != false ||
		preview["ready_notification_action_request_dispatch_dry_run_item_count"] != 0 ||
		preview["missing_notification_action_request_dispatch_dry_run_item_count"] != 4 {
		t.Fatalf("expected notification action request-object dispatch dry-run audit to fail closed without evidence: %#v", preview)
	}
	for _, item := range preview["notification_action_request_dispatch_dry_run_items"].([]map[string]any) {
		if item["evidence_present"] != false ||
			item["request_object_dispatch_dry_run_ready"] != false ||
			item["dispatch_dry_run_execution_enabled"] != false ||
			item["dispatch_dry_run_executed"] != false ||
			item["request_object_creation_enabled"] != false ||
			item["request_object_dispatch_enabled"] != false ||
			item["notification_action_enabled"] != false ||
			item["notification_sent"] != false ||
			item["host_root_modified"] != false ||
			item["notification_action_request_object_dispatch_dry_run_status"] != "missing-redacted-status-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-evidence" {
			t.Fatalf("unsafe fail-closed notification action request-object dispatch dry-run item: %#v", item)
		}
	}
}
