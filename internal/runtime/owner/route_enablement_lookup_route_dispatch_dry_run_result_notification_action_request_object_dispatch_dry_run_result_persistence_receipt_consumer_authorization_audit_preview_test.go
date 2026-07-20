package owner

import (
	"os"
	"path/filepath"
	"testing"
)

func TestProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerAuthorizationAuditPreviewModelsAuthorizationBoundary(t *testing.T) {
	if routeEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerAuthorizationAuditRequest != "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-authorization-audit-preview" {
		t.Fatalf("unexpected notification action request-object dispatch dry-run result persistence receipt consumer authorization request constant: %s", routeEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerAuthorizationAuditRequest)
	}
	preview, err := NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerAuthorizationAuditPreview(projectRootForOwnerTest(t))
	if err != nil {
		t.Fatalf("NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerAuthorizationAuditPreview returned error: %v", err)
	}
	if preview["version"] != currentProjectVersion(t) ||
		preview["schema_version"] != "xnix.runtime.production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_kde_safe_redacted_status_closed_record_writer_storage_root_authorization_receipt_dry_run_lookup_result_consumer_projection_route_enablement_lookup_route_dispatch_dry_run_result_notification_action_request_object_dispatch_dry_run_result_persistence_receipt_consumer_authorization_audit.v1" ||
		preview["request_type"] != routeEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerAuthorizationAuditRequest ||
		preview["audit_type"] != "receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-authorization-audit" ||
		preview["audit_decision"] != "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-authorization-audit-ready-authorization-disabled" {
		t.Fatalf("unexpected notification action request-object dispatch dry-run result persistence receipt consumer authorization schema: %#v", preview)
	}
	for _, key := range []string{
		"current_mainline_consumed",
		"route_enablement_lookup_route_dispatch_dry_run_result_notification_action_request_object_dispatch_dry_run_result_persistence_receipt_consumption_gate_audit_consumed",
		"route_enablement_lookup_route_dispatch_dry_run_result_notification_action_request_object_dispatch_dry_run_result_persistence_receipt_consumption_gate_audit_ready",
		"route_enablement_lookup_route_dispatch_dry_run_result_notification_action_request_object_dispatch_dry_run_result_persistence_receipt_acceptance_audit_ready",
		"route_enablement_lookup_route_dispatch_dry_run_result_notification_action_request_object_dispatch_dry_run_result_persistence_receipt_audit_ready",
		"route_enablement_lookup_route_dispatch_dry_run_result_notification_action_request_object_dispatch_dry_run_result_persistence_authorization_audit_ready",
		"route_enablement_lookup_route_dispatch_dry_run_result_notification_action_request_object_dispatch_dry_run_result_visibility_audit_ready",
		"route_enablement_lookup_route_dispatch_dry_run_result_notification_action_request_object_dispatch_dry_run_audit_ready",
		"lookup_route_dispatch_dry_run_result_notification_action_request_object_dispatch_dry_run_result_persistence_receipt_consumer_authorization_audit_required",
		"lookup_route_dispatch_dry_run_result_notification_action_request_object_dispatch_dry_run_result_persistence_receipt_consumer_authorization_modeled",
		"lookup_route_dispatch_dry_run_result_notification_action_request_object_dispatch_dry_run_result_persistence_receipt_consumer_authorization_ready",
		"route_enablement_lookup_route_dispatch_dry_run_result_notification_action_request_object_dispatch_dry_run_result_persistence_receipt_consumer_authorization_ready",
		"notification_action_request_object_dispatch_dry_run_result_persistence_receipt_consumer_authorization_review_only",
		"notification_action_request_object_dispatch_dry_run_result_persistence_receipt_consumption_gate_review_only_consumed",
		"result_persistence_authorization_ready",
		"result_persistence_receipt_ready",
		"result_persistence_receipt_acceptance_ready",
		"result_persistence_receipt_consumption_required",
		"result_persistence_receipt_consumption_ready",
		"result_persistence_receipt_consumer_authorization_required",
		"result_persistence_receipt_consumer_authorization_modeled",
		"result_persistence_receipt_consumer_authorization_ready",
		"kde_safe_redacted_result_only",
		"raw_result_hidden",
		"runtime_owned",
		"go_runtime_backed",
	} {
		if preview[key] != true {
			t.Fatalf("expected %s to be true in notification action request-object dispatch dry-run result persistence receipt consumer authorization preview: %#v", key, preview)
		}
	}
	for _, key := range []string{
		"result_persistence_authorized",
		"operator_persistence_approval_present",
		"caller_state_root_required",
		"result_persistence_receipt_present",
		"result_persistence_receipt_accepted",
		"result_persistence_receipt_consumed",
		"receipt_consumer_authorization_enabled",
		"receipt_consumer_authorized",
		"consumer_authorization_granted",
		"kde_consumer_authorized",
		"runtime_consumer_authorized",
		"receipt_consumption_gate_enabled",
		"receipt_consumption_enabled",
		"receipt_consumed",
		"receipt_acceptance_enabled",
		"receipt_accepted",
		"receipt_write_enabled",
		"receipt_persistence_enabled",
		"receipt_present",
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
		"status_persistence_write_enabled",
		"redacted_summary_persisted",
		"kde_status_persisted",
		"runtime_diagnostics_persisted",
		"raw_result_exposed",
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
			t.Fatalf("expected %s to remain false in notification action request-object dispatch dry-run result persistence receipt consumer authorization preview: %#v", key, preview)
		}
	}
	if preview["notification_action_request_dispatch_dry_run_result_persistence_receipt_consumer_authorization_item_count"] != 4 ||
		preview["ready_notification_action_request_dispatch_dry_run_result_persistence_receipt_consumer_authorization_item_count"] != 4 ||
		preview["missing_notification_action_request_dispatch_dry_run_result_persistence_receipt_consumer_authorization_item_count"] != 0 ||
		preview["persistence_receipt_consumption_gate_audit_consumed_item_count"] != 4 ||
		preview["authorized_result_persistence_receipt_consumer_item_count"] != 0 ||
		preview["consumed_result_persistence_receipt_item_count"] != 0 ||
		preview["accepted_result_persistence_receipt_item_count"] != 0 ||
		preview["present_result_persistence_receipt_item_count"] != 0 ||
		preview["persisted_dry_run_result_item_count"] != 0 ||
		preview["side_effect_notification_action_request_dispatch_dry_run_result_persistence_receipt_consumer_authorization_item_count"] != 0 {
		t.Fatalf("unexpected notification action request-object dispatch dry-run result persistence receipt consumer authorization counts: %#v", preview)
	}
	if !sameStrings(preview["notification_action_request_dispatch_dry_run_result_persistence_receipt_consumer_authorization_item_ids"].([]string), []string{
		"route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-authorization-install-failure",
		"route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-authorization-repair-suggestion",
		"route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-authorization-environment-switch",
		"route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-authorization-approval-info",
	}) {
		t.Fatalf("unexpected notification action request-object dispatch dry-run result persistence receipt consumer authorization item ids: %#v", preview["notification_action_request_dispatch_dry_run_result_persistence_receipt_consumer_authorization_item_ids"])
	}
	for _, item := range preview["notification_action_request_dispatch_dry_run_result_persistence_receipt_consumer_authorization_items"].([]map[string]any) {
		if item["evidence_present"] != true ||
			item["notification_action_request_object_dispatch_dry_run_result_persistence_receipt_consumption_gate_audit_consumed"] != true ||
			item["result_persistence_receipt_ready"] != true ||
			item["result_persistence_receipt_present"] != false ||
			item["result_persistence_receipt_consumed"] != false ||
			item["result_persistence_receipt_consumer_authorization_required"] != true ||
			item["result_persistence_receipt_consumer_authorization_modeled"] != true ||
			item["result_persistence_receipt_consumer_authorization_ready"] != true ||
			item["receipt_consumer_authorization_enabled"] != false ||
			item["receipt_consumer_authorized"] != false ||
			item["consumer_authorization_granted"] != false ||
			item["kde_consumer_authorized"] != false ||
			item["runtime_consumer_authorized"] != false ||
			item["receipt_consumption_enabled"] != false ||
			item["receipt_consumed"] != false ||
			item["receipt_acceptance_enabled"] != false ||
			item["receipt_accepted"] != false ||
			item["receipt_write_enabled"] != false ||
			item["dry_run_result_persistence_enabled"] != false ||
			item["dispatch_dry_run_execution_enabled"] != false ||
			item["request_object_creation_enabled"] != false ||
			item["notification_action_enabled"] != false ||
			item["notification_sent"] != false ||
			item["side_effects_disabled"] != true ||
			item["host_root_modified"] != false ||
			item["notification_action_request_object_dispatch_dry_run_result_persistence_receipt_consumer_authorization_status"] != "redacted-status-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-authorization-ready-authorization-disabled" {
			t.Fatalf("unsafe notification action request-object dispatch dry-run result persistence receipt consumer authorization item: %#v", item)
		}
	}
	counts := preview["counts"].(map[string]int)
	if counts["total"] != 10 || counts["passed"] != 10 || counts["blocked"] != 0 || counts["pending"] != 0 {
		t.Fatalf("unexpected notification action request-object dispatch dry-run result persistence receipt consumer authorization checks: %#v", counts)
	}
	if err := validateNoBackendTerms(preview, "KDE-safe notification action request-object dispatch dry-run result persistence receipt consumer authorization test"); err != nil {
		t.Fatalf("validateNoBackendTerms returned error: %v", err)
	}
}

func TestProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerAuthorizationAuditPreviewFailsClosedWithoutSources(t *testing.T) {
	root, err := os.MkdirTemp("", "xnix-action-request-dispatch-dry-run-result-persistence-receipt-consumer-authorization-*")
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
	if err := os.WriteFile(filepath.Join(root, "docs", "xnix-current-mainline.md"), []byte("missing notification action request-object dispatch dry-run result persistence receipt consumer authorization evidence"), 0o644); err != nil {
		t.Fatalf("WriteFile mainline returned error: %v", err)
	}

	preview, err := NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerAuthorizationAuditPreview(root)
	if err != nil {
		t.Fatalf("NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerAuthorizationAuditPreview returned error: %v", err)
	}
	if preview["audit_decision"] != "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-authorization-audit-blocked" ||
		preview["current_mainline_consumed"] != false ||
		preview["route_enablement_lookup_route_dispatch_dry_run_result_notification_action_request_object_dispatch_dry_run_result_persistence_receipt_consumption_gate_audit_consumed"] != false ||
		preview["lookup_route_dispatch_dry_run_result_notification_action_request_object_dispatch_dry_run_result_persistence_receipt_consumer_authorization_ready"] != false ||
		preview["receipt_consumer_authorization_enabled"] != false ||
		preview["consumer_authorization_granted"] != false ||
		preview["ready_notification_action_request_dispatch_dry_run_result_persistence_receipt_consumer_authorization_item_count"] != 0 ||
		preview["missing_notification_action_request_dispatch_dry_run_result_persistence_receipt_consumer_authorization_item_count"] != 4 {
		t.Fatalf("expected notification action request-object dispatch dry-run result persistence receipt consumer authorization audit to fail closed without evidence: %#v", preview)
	}
	for _, item := range preview["notification_action_request_dispatch_dry_run_result_persistence_receipt_consumer_authorization_items"].([]map[string]any) {
		if item["evidence_present"] != false ||
			item["result_persistence_receipt_consumer_authorization_ready"] != false ||
			item["receipt_consumer_authorization_enabled"] != false ||
			item["receipt_consumer_authorized"] != false ||
			item["consumer_authorization_granted"] != false ||
			item["receipt_consumption_enabled"] != false ||
			item["receipt_consumed"] != false ||
			item["receipt_write_enabled"] != false ||
			item["dry_run_result_persistence_enabled"] != false ||
			item["dispatch_dry_run_execution_enabled"] != false ||
			item["request_object_creation_enabled"] != false ||
			item["notification_action_enabled"] != false ||
			item["notification_sent"] != false ||
			item["host_root_modified"] != false ||
			item["notification_action_request_object_dispatch_dry_run_result_persistence_receipt_consumer_authorization_status"] != "missing-redacted-status-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-authorization-evidence" {
			t.Fatalf("unsafe fail-closed notification action request-object dispatch dry-run result persistence receipt consumer authorization item: %#v", item)
		}
	}
}
