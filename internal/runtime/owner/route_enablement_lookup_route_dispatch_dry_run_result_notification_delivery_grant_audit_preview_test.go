package owner

import (
	"os"
	"path/filepath"
	"testing"
)

func TestProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationDeliveryGrantAuditPreviewModelsGrantBoundary(t *testing.T) {
	if routeEnablementLookupRouteDispatchDryRunResultNotificationDeliveryGrantAuditRequest != "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-delivery-grant-audit-preview" {
		t.Fatalf("unexpected delivery grant request constant: %s", routeEnablementLookupRouteDispatchDryRunResultNotificationDeliveryGrantAuditRequest)
	}
	preview, err := NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationDeliveryGrantAuditPreview(projectRootForOwnerTest(t))
	if err != nil {
		t.Fatalf("NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationDeliveryGrantAuditPreview returned error: %v", err)
	}
	if preview["version"] != currentProjectVersion(t) ||
		preview["schema_version"] != "xnix.runtime.production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_kde_safe_redacted_status_closed_record_writer_storage_root_authorization_receipt_dry_run_lookup_result_consumer_projection_route_enablement_lookup_route_dispatch_dry_run_result_notification_delivery_grant_audit.v1" ||
		preview["request_type"] != routeEnablementLookupRouteDispatchDryRunResultNotificationDeliveryGrantAuditRequest ||
		preview["audit_type"] != "receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-delivery-grant-audit" ||
		preview["audit_decision"] != "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-delivery-grant-audit-ready-grant-issuance-disabled" {
		t.Fatalf("unexpected notification delivery grant schema: %#v", preview)
	}

	for _, key := range []string{
		"current_mainline_consumed",
		"route_enablement_lookup_route_dispatch_dry_run_result_notification_delivery_authorization_consumed",
		"route_enablement_lookup_route_dispatch_dry_run_result_notification_delivery_authorization_ready",
		"lookup_route_dispatch_dry_run_result_notification_delivery_grant_audit_required",
		"lookup_route_dispatch_dry_run_result_notification_delivery_grant_modeled",
		"lookup_route_dispatch_dry_run_result_notification_delivery_grant_ready",
		"route_enablement_lookup_route_dispatch_dry_run_result_notification_delivery_grant_ready",
		"notification_delivery_authorization_review_only_consumed",
		"kde_notification_delivery_grant_modeled",
		"notification_center_delivery_grant_modeled",
		"install_failure_delivery_grant_modeled",
		"repair_suggestion_delivery_grant_modeled",
		"environment_switch_delivery_grant_modeled",
		"approval_info_delivery_grant_modeled",
		"notification_delivery_grant_review_only",
		"notification_delivery_disabled",
		"notification_action_disabled",
		"action_cards_remain_disabled",
		"explicit_user_consent_required",
		"kde_safe_redacted_result_only",
		"raw_result_hidden",
		"runtime_owned",
		"go_runtime_backed",
	} {
		if preview[key] != true {
			t.Fatalf("expected %s to be true in delivery grant preview: %#v", key, preview)
		}
	}

	for _, key := range []string{
		"notification_delivery_grant_issued",
		"notification_delivery_authorization_granted",
		"explicit_user_consent_collected",
		"receipt_present",
		"receipt_accepted",
		"receipt_consumed",
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
			t.Fatalf("expected %s to remain false in delivery grant preview: %#v", key, preview)
		}
	}

	if preview["notification_delivery_grant_item_count"] != 4 ||
		preview["required_notification_delivery_grant_item_count"] != 4 ||
		preview["ready_notification_delivery_grant_item_count"] != 4 ||
		preview["missing_notification_delivery_grant_item_count"] != 0 ||
		preview["issued_notification_delivery_grant_item_count"] != 0 ||
		preview["authorized_notification_delivery_item_count"] != 0 ||
		preview["delivered_notification_item_count"] != 0 ||
		preview["notification_action_enabled_item_count"] != 0 ||
		preview["action_card_enabled_item_count"] != 0 ||
		preview["side_effect_notification_delivery_grant_item_count"] != 0 {
		t.Fatalf("unexpected delivery grant counts: %#v", preview)
	}
	if !sameStrings(preview["notification_delivery_grant_item_ids"].([]string), []string{
		"route-enablement-lookup-route-dispatch-dry-run-result-notification-delivery-grant-install-failure",
		"route-enablement-lookup-route-dispatch-dry-run-result-notification-delivery-grant-repair-suggestion",
		"route-enablement-lookup-route-dispatch-dry-run-result-notification-delivery-grant-environment-switch",
		"route-enablement-lookup-route-dispatch-dry-run-result-notification-delivery-grant-approval-info",
	}) {
		t.Fatalf("unexpected delivery grant item ids: %#v", preview["notification_delivery_grant_item_ids"])
	}
	for _, item := range preview["notification_delivery_grant_items"].([]map[string]any) {
		if item["evidence_present"] != true ||
			item["notification_delivery_authorization_consumed"] != true ||
			item["notification_delivery_authorization_ready"] != true ||
			item["delivery_grant_required"] != true ||
			item["delivery_grant_modeled"] != true ||
			item["delivery_grant_ready"] != true ||
			item["delivery_grant_issued"] != false ||
			item["delivery_authorization_granted"] != false ||
			item["explicit_user_consent_required"] != true ||
			item["explicit_user_consent_collected"] != false ||
			item["delivery_review_only"] != true ||
			item["kde_safe_redacted_result_only"] != true ||
			item["raw_result_hidden"] != true ||
			item["notification_delivery_enabled"] != false ||
			item["notification_sent"] != false ||
			item["notification_center_event_triggered"] != false ||
			item["notification_action_enabled"] != false ||
			item["action_card_enabled"] != false ||
			item["side_effects_disabled"] != true ||
			item["host_root_modified"] != false ||
			item["delivery_grant_status"] != "redacted-status-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-delivery-grant-ready-grant-issuance-disabled" {
			t.Fatalf("unsafe delivery grant item: %#v", item)
		}
	}

	counts := preview["counts"].(map[string]int)
	if counts["total"] != 7 || counts["passed"] != 7 || counts["blocked"] != 0 || counts["pending"] != 0 {
		t.Fatalf("unexpected delivery grant checks: %#v", counts)
	}
	if !sameStrings(preview["check_ids"].([]string), []string{
		"current-mainline-consumed",
		"route-enablement-lookup-route-dispatch-dry-run-result-notification-delivery-authorization-consumed",
		"lookup-route-dispatch-dry-run-result-notification-delivery-grant-modeled",
		"four-kde-notification-delivery-events-grant-review-only",
		"grant-issuance-sending-and-events-disabled",
		"notification-actions-action-cards-and-raw-results-disabled",
		"production-and-host-boundary-closed",
	}) {
		t.Fatalf("unexpected delivery grant check ids: %#v", preview["check_ids"])
	}
	if err := validateNoBackendTerms(preview, "KDE-safe notification delivery grant test"); err != nil {
		t.Fatalf("validateNoBackendTerms returned error: %v", err)
	}
}

func TestProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationDeliveryGrantAuditPreviewFailsClosedWithoutSources(t *testing.T) {
	root, err := os.MkdirTemp("", "xnix-grant-*")
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
	if err := os.WriteFile(filepath.Join(root, "docs", "xnix-current-mainline.md"), []byte("missing delivery grant evidence"), 0o644); err != nil {
		t.Fatalf("WriteFile mainline returned error: %v", err)
	}

	preview, err := NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationDeliveryGrantAuditPreview(root)
	if err != nil {
		t.Fatalf("NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationDeliveryGrantAuditPreview returned error: %v", err)
	}
	if preview["audit_decision"] != "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-delivery-grant-audit-blocked" ||
		preview["current_mainline_consumed"] != false ||
		preview["route_enablement_lookup_route_dispatch_dry_run_result_notification_delivery_authorization_consumed"] != false ||
		preview["lookup_route_dispatch_dry_run_result_notification_delivery_grant_ready"] != false ||
		preview["ready_notification_delivery_grant_item_count"] != 0 ||
		preview["missing_notification_delivery_grant_item_count"] != 4 ||
		preview["notification_delivery_grant_issued"] != false ||
		preview["notification_sent"] != false ||
		preview["host_root_modified"] != false {
		t.Fatalf("delivery grant preview must fail closed without sources: %#v", preview)
	}
	counts := preview["counts"].(map[string]int)
	if counts["total"] != 7 || counts["passed"] != 2 || counts["blocked"] != 5 {
		t.Fatalf("unexpected fail-closed delivery grant counts: %#v", counts)
	}
}
