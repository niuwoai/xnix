package owner

import (
	"os"
	"path/filepath"
	"testing"
)

func TestProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionEnablementAuditPreviewModelsActionEnablementBoundary(t *testing.T) {
	if routeEnablementLookupRouteDispatchDryRunResultNotificationActionEnablementAuditRequest != "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-enablement-audit-preview" {
		t.Fatalf("unexpected notification action enablement request constant: %s", routeEnablementLookupRouteDispatchDryRunResultNotificationActionEnablementAuditRequest)
	}
	preview, err := NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionEnablementAuditPreview(projectRootForOwnerTest(t))
	if err != nil {
		t.Fatalf("NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionEnablementAuditPreview returned error: %v", err)
	}
	if preview["version"] != currentProjectVersion(t) ||
		preview["schema_version"] != "xnix.runtime.production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_kde_safe_redacted_status_closed_record_writer_storage_root_authorization_receipt_dry_run_lookup_result_consumer_projection_route_enablement_lookup_route_dispatch_dry_run_result_notification_action_enablement_audit.v1" ||
		preview["request_type"] != routeEnablementLookupRouteDispatchDryRunResultNotificationActionEnablementAuditRequest ||
		preview["audit_type"] != "receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-enablement-audit" ||
		preview["audit_decision"] != "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-enablement-audit-ready-actions-disabled" {
		t.Fatalf("unexpected notification action enablement schema: %#v", preview)
	}
	for _, key := range []string{
		"current_mainline_consumed",
		"route_enablement_lookup_route_dispatch_dry_run_result_notification_delivery_grant_audit_consumed",
		"route_enablement_lookup_route_dispatch_dry_run_result_notification_delivery_grant_audit_ready",
		"route_enablement_lookup_route_dispatch_dry_run_result_notification_delivery_execution_authorization_audit_consumed",
		"route_enablement_lookup_route_dispatch_dry_run_result_notification_delivery_execution_authorization_audit_ready",
		"route_enablement_lookup_route_dispatch_dry_run_result_notification_delivery_consent_collection_receipt_consumer_enablement_gate_audit_consumed",
		"route_enablement_lookup_route_dispatch_dry_run_result_notification_delivery_consent_collection_receipt_consumer_enablement_gate_audit_ready",
		"route_enablement_lookup_route_dispatch_dry_run_result_notification_delivery_consent_collection_receipt_consumer_authorization_audit_consumed",
		"route_enablement_lookup_route_dispatch_dry_run_result_notification_delivery_consent_collection_receipt_consumer_authorization_audit_ready",
		"route_enablement_lookup_route_dispatch_dry_run_result_notification_delivery_consent_collection_receipt_consumption_gate_audit_consumed",
		"route_enablement_lookup_route_dispatch_dry_run_result_notification_delivery_consent_collection_receipt_consumption_gate_audit_ready",
		"route_enablement_lookup_route_dispatch_dry_run_result_notification_delivery_consent_collection_receipt_acceptance_audit_consumed",
		"route_enablement_lookup_route_dispatch_dry_run_result_notification_delivery_consent_collection_receipt_acceptance_audit_ready",
		"lookup_route_dispatch_dry_run_result_notification_action_enablement_audit_required",
		"lookup_route_dispatch_dry_run_result_notification_action_enablement_modeled",
		"lookup_route_dispatch_dry_run_result_notification_action_enablement_ready",
		"route_enablement_lookup_route_dispatch_dry_run_result_notification_action_enablement_ready",
		"notification_action_enablement_review_only",
		"notification_delivery_grant_review_only_consumed",
		"notification_delivery_execution_authorization_review_only_consumed",
		"kde_notification_action_enablement_modeled",
		"notification_center_action_enablement_modeled",
		"install_failure_notification_action_enablement_modeled",
		"repair_suggestion_notification_action_enablement_modeled",
		"environment_switch_notification_action_enablement_modeled",
		"approval_info_notification_action_enablement_modeled",
		"delivery_grant_required",
		"delivery_grant_modeled",
		"delivery_grant_ready",
		"grant_issuance_required",
		"grant_issuance_ready",
		"notification_action_required",
		"notification_action_modeled",
		"notification_action_ready",
		"notification_action_enablement_required",
		"notification_action_enablement_modeled",
		"notification_action_enablement_ready",
		"action_card_enablement_required",
		"action_card_enablement_modeled",
		"action_card_enablement_ready",
		"notification_delivery_disabled",
		"notification_action_disabled",
		"action_cards_remain_disabled",
		"kde_safe_redacted_result_only",
		"raw_result_hidden",
		"runtime_owned",
		"go_runtime_backed",
	} {
		if preview[key] != true {
			t.Fatalf("expected %s to be true in notification action enablement preview: %#v", key, preview)
		}
	}
	for _, key := range []string{
		"delivery_execution_authorized",
		"notification_delivery_execution_authorized",
		"notification_delivery_execution_enabled",
		"delivery_grant_issued",
		"notification_delivery_grant_issued",
		"notification_delivery_authorization_granted",
		"grant_issuance_enabled",
		"notification_action_enablement_enabled",
		"action_card_enablement_enabled",
		"explicit_user_consent_collected",
		"explicit_user_consent_persisted",
		"consent_receipt_created",
		"consent_receipt_persisted",
		"consent_receipt_accepted",
		"consent_receipt_consumed",
		"receipt_present",
		"receipt_accepted",
		"receipt_consumed",
		"consumer_authorization_granted",
		"kde_consumer_enabled",
		"runtime_consumer_enabled",
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
		"request_object_creation_enabled",
		"request_object_dispatch_enabled",
		"portal_request_created",
		"notification_delivery_enabled",
		"notification_sent",
		"notification_center_event_triggered",
		"notification_action_enabled",
		"action_card_enabled",
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
			t.Fatalf("expected %s to remain false in notification action enablement preview: %#v", key, preview)
		}
	}
	if preview["notification_action_enablement_item_count"] != 4 ||
		preview["required_notification_action_enablement_item_count"] != 4 ||
		preview["ready_notification_action_enablement_item_count"] != 4 ||
		preview["missing_notification_action_enablement_item_count"] != 0 ||
		preview["delivery_grant_audit_consumed_item_count"] != 4 ||
		preview["execution_authorization_audit_consumed_item_count"] != 4 ||
		preview["notification_action_enablement_modeled_item_count"] != 4 ||
		preview["action_card_enablement_modeled_item_count"] != 4 ||
		preview["enabled_notification_action_item_count"] != 0 ||
		preview["enabled_action_card_item_count"] != 0 ||
		preview["issued_notification_delivery_grant_item_count"] != 0 ||
		preview["delivered_notification_item_count"] != 0 ||
		preview["side_effect_notification_action_enablement_item_count"] != 0 {
		t.Fatalf("unexpected notification action enablement counts: %#v", preview)
	}
	if !sameStrings(preview["notification_action_enablement_item_ids"].([]string), []string{
		"route-enablement-lookup-route-dispatch-dry-run-result-notification-action-enablement-install-failure",
		"route-enablement-lookup-route-dispatch-dry-run-result-notification-action-enablement-repair-suggestion",
		"route-enablement-lookup-route-dispatch-dry-run-result-notification-action-enablement-environment-switch",
		"route-enablement-lookup-route-dispatch-dry-run-result-notification-action-enablement-approval-info",
	}) {
		t.Fatalf("unexpected notification action enablement item ids: %#v", preview["notification_action_enablement_item_ids"])
	}
	for _, item := range preview["notification_action_enablement_items"].([]map[string]any) {
		if item["evidence_present"] != true ||
			item["notification_delivery_grant_audit_consumed"] != true ||
			item["notification_delivery_execution_authorization_audit_consumed"] != true ||
			item["notification_action_required"] != true ||
			item["notification_action_ready"] != true ||
			item["notification_action_enablement_required"] != true ||
			item["notification_action_enablement_ready"] != true ||
			item["action_card_enablement_required"] != true ||
			item["action_card_enablement_ready"] != true ||
			item["notification_action_enablement_enabled"] != false ||
			item["action_card_enablement_enabled"] != false ||
			item["notification_action_enabled"] != false ||
			item["action_card_enabled"] != false ||
			item["notification_delivery_grant_issued"] != false ||
			item["notification_delivery_enabled"] != false ||
			item["notification_sent"] != false ||
			item["notification_center_event_triggered"] != false ||
			item["delivery_review_only"] != true ||
			item["kde_safe_redacted_result_only"] != true ||
			item["raw_result_hidden"] != true ||
			item["side_effects_disabled"] != true ||
			item["host_root_modified"] != false ||
			item["notification_action_enablement_status"] != "redacted-status-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-enablement-ready-actions-disabled" {
			t.Fatalf("unsafe notification action enablement item: %#v", item)
		}
	}
	counts := preview["counts"].(map[string]int)
	if counts["total"] != 11 || counts["passed"] != 11 || counts["blocked"] != 0 || counts["pending"] != 0 {
		t.Fatalf("unexpected notification action enablement checks: %#v", counts)
	}
	if !sameStrings(preview["check_ids"].([]string), []string{
		"current-mainline-consumed",
		"route-enablement-lookup-route-dispatch-dry-run-result-notification-delivery-grant-audit-consumed",
		"route-enablement-lookup-route-dispatch-dry-run-result-notification-delivery-execution-authorization-audit-consumed",
		"route-enablement-lookup-route-dispatch-dry-run-result-notification-delivery-consent-collection-receipt-consumer-enablement-gate-audit-consumed",
		"route-enablement-lookup-route-dispatch-dry-run-result-notification-delivery-consent-collection-receipt-consumer-authorization-audit-consumed",
		"route-enablement-lookup-route-dispatch-dry-run-result-notification-delivery-consent-collection-receipt-consumption-gate-audit-consumed",
		"route-enablement-lookup-route-dispatch-dry-run-result-notification-delivery-consent-collection-receipt-acceptance-audit-consumed",
		"lookup-route-dispatch-dry-run-result-notification-action-enablement-modeled",
		"four-kde-notification-action-enablement-events-review-only",
		"notification-actions-action-cards-delivery-and-events-disabled",
		"production-and-host-boundary-closed",
	}) {
		t.Fatalf("unexpected notification action enablement check ids: %#v", preview["check_ids"])
	}
	if err := validateNoBackendTerms(preview, "KDE-safe notification action enablement test"); err != nil {
		t.Fatalf("validateNoBackendTerms returned error: %v", err)
	}
}

func TestProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionEnablementAuditPreviewFailsClosedWithoutSources(t *testing.T) {
	root, err := os.MkdirTemp("", "xnix-action-enablement-*")
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
	if err := os.WriteFile(filepath.Join(root, "docs", "xnix-current-mainline.md"), []byte("missing notification action enablement evidence"), 0o644); err != nil {
		t.Fatalf("WriteFile mainline returned error: %v", err)
	}

	preview, err := NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionEnablementAuditPreview(root)
	if err != nil {
		t.Fatalf("NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionEnablementAuditPreview returned error: %v", err)
	}
	if preview["audit_decision"] != "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-enablement-audit-blocked" ||
		preview["current_mainline_consumed"] != false ||
		preview["route_enablement_lookup_route_dispatch_dry_run_result_notification_delivery_grant_audit_consumed"] != false ||
		preview["lookup_route_dispatch_dry_run_result_notification_action_enablement_ready"] != false ||
		preview["ready_notification_action_enablement_item_count"] != 0 ||
		preview["missing_notification_action_enablement_item_count"] != 4 {
		t.Fatalf("expected notification action enablement audit to fail closed without evidence: %#v", preview)
	}
	for _, item := range preview["notification_action_enablement_items"].([]map[string]any) {
		if item["evidence_present"] != false ||
			item["notification_action_enablement_ready"] != false ||
			item["notification_action_enablement_enabled"] != false ||
			item["action_card_enablement_enabled"] != false ||
			item["notification_action_enabled"] != false ||
			item["action_card_enabled"] != false ||
			item["notification_delivery_grant_issued"] != false ||
			item["notification_sent"] != false ||
			item["host_root_modified"] != false {
			t.Fatalf("unsafe fail-closed notification action enablement item: %#v", item)
		}
	}
}
