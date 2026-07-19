package owner

import (
	"os"
	"path/filepath"
	"testing"
)

func TestProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationDeliveryConsentCollectionReceiptConsumerEnablementGateAuditPreviewModelsConsumerEnablementGateBoundary(t *testing.T) {
	if routeEnablementLookupRouteDispatchDryRunResultNotificationDeliveryConsentCollectionReceiptConsumerEnablementGateAuditRequest != "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-delivery-consent-collection-receipt-consumer-enablement-gate-audit-preview" {
		t.Fatalf("unexpected consent collection receipt consumer enablement gate request constant: %s", routeEnablementLookupRouteDispatchDryRunResultNotificationDeliveryConsentCollectionReceiptConsumerEnablementGateAuditRequest)
	}
	preview, err := NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationDeliveryConsentCollectionReceiptConsumerEnablementGateAuditPreview(projectRootForOwnerTest(t))
	if err != nil {
		t.Fatalf("NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationDeliveryConsentCollectionReceiptConsumerEnablementGateAuditPreview returned error: %v", err)
	}
	if preview["version"] != currentProjectVersion(t) ||
		preview["schema_version"] != "xnix.runtime.production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_kde_safe_redacted_status_closed_record_writer_storage_root_authorization_receipt_dry_run_lookup_result_consumer_projection_route_enablement_lookup_route_dispatch_dry_run_result_notification_delivery_consent_collection_receipt_consumer_enablement_gate_audit.v1" ||
		preview["request_type"] != routeEnablementLookupRouteDispatchDryRunResultNotificationDeliveryConsentCollectionReceiptConsumerEnablementGateAuditRequest ||
		preview["audit_type"] != "receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-delivery-consent-collection-receipt-consumer-enablement-gate-audit" ||
		preview["audit_decision"] != "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-delivery-consent-collection-receipt-consumer-enablement-gate-audit-ready-consumers-disabled" {
		t.Fatalf("unexpected notification delivery consent collection receipt consumer enablement gate schema: %#v", preview)
	}
	for _, key := range []string{
		"current_mainline_consumed",
		"route_enablement_lookup_route_dispatch_dry_run_result_notification_delivery_consent_collection_receipt_consumer_authorization_audit_consumed",
		"route_enablement_lookup_route_dispatch_dry_run_result_notification_delivery_consent_collection_receipt_consumer_authorization_audit_ready",
		"route_enablement_lookup_route_dispatch_dry_run_result_notification_delivery_consent_collection_receipt_consumption_gate_audit_consumed",
		"route_enablement_lookup_route_dispatch_dry_run_result_notification_delivery_consent_collection_receipt_consumption_gate_audit_ready",
		"route_enablement_lookup_route_dispatch_dry_run_result_notification_delivery_consent_collection_receipt_acceptance_audit_consumed",
		"route_enablement_lookup_route_dispatch_dry_run_result_notification_delivery_consent_collection_receipt_acceptance_audit_ready",
		"lookup_route_dispatch_dry_run_result_notification_delivery_consent_collection_receipt_consumer_enablement_gate_audit_required",
		"lookup_route_dispatch_dry_run_result_notification_delivery_consent_collection_receipt_consumer_enablement_gate_modeled",
		"lookup_route_dispatch_dry_run_result_notification_delivery_consent_collection_receipt_consumer_enablement_gate_ready",
		"route_enablement_lookup_route_dispatch_dry_run_result_notification_delivery_consent_collection_receipt_consumer_enablement_gate_ready",
		"notification_delivery_consent_collection_receipt_consumer_authorization_review_only_consumed",
		"notification_delivery_consent_collection_receipt_consumption_gate_review_only_consumed",
		"notification_delivery_consent_collection_receipt_acceptance_review_only_consumed",
		"kde_notification_delivery_consent_collection_receipt_consumer_enablement_gate_modeled",
		"notification_center_delivery_consent_collection_receipt_consumer_enablement_gate_modeled",
		"install_failure_consent_collection_receipt_consumer_enablement_gate_modeled",
		"repair_suggestion_consent_collection_receipt_consumer_enablement_gate_modeled",
		"environment_switch_consent_collection_receipt_consumer_enablement_gate_modeled",
		"approval_info_consent_collection_receipt_consumer_enablement_gate_modeled",
		"notification_delivery_consent_collection_receipt_consumer_enablement_gate_review_only",
		"consumer_enablement_gate_required",
		"consumer_enablement_gate_modeled",
		"consumer_enablement_gate_ready",
		"consumer_enablement_required",
		"consumer_enablement_modeled",
		"consumer_enablement_ready",
		"notification_delivery_disabled",
		"notification_action_disabled",
		"action_cards_remain_disabled",
		"kde_safe_redacted_result_only",
		"raw_result_hidden",
		"runtime_owned",
		"go_runtime_backed",
	} {
		if preview[key] != true {
			t.Fatalf("expected %s to be true in consent collection receipt consumer enablement gate preview: %#v", key, preview)
		}
	}
	for _, key := range []string{
		"consent_collection_receipt_creation_allowed",
		"consent_collection_receipt_creation_enabled",
		"consent_collection_receipt_acceptance_allowed",
		"consent_collection_receipt_acceptance_enabled",
		"consent_collection_receipt_consumption_allowed",
		"consent_collection_receipt_consumption_enabled",
		"consumer_authorization_granted",
		"kde_consumer_authorized",
		"runtime_consumer_authorized",
		"kde_consumer_enabled",
		"runtime_consumer_enabled",
		"consumer_enabled",
		"explicit_user_consent_collection_allowed",
		"explicit_user_consent_collection_enabled",
		"explicit_user_consent_collected",
		"explicit_user_consent_persisted",
		"consent_receipt_created",
		"consent_receipt_persisted",
		"consent_receipt_accepted",
		"consent_receipt_consumed",
		"notification_delivery_grant_issued",
		"notification_delivery_authorization_granted",
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
			t.Fatalf("expected %s to remain false in consent collection receipt consumer enablement gate preview: %#v", key, preview)
		}
	}
	if preview["notification_delivery_consent_collection_receipt_consumer_enablement_gate_item_count"] != 4 ||
		preview["required_notification_delivery_consent_collection_receipt_consumer_enablement_gate_item_count"] != 4 ||
		preview["ready_notification_delivery_consent_collection_receipt_consumer_enablement_gate_item_count"] != 4 ||
		preview["missing_notification_delivery_consent_collection_receipt_consumer_enablement_gate_item_count"] != 0 ||
		preview["receipt_consumer_authorization_audit_consumed_item_count"] != 4 ||
		preview["receipt_consumption_gate_audit_consumed_item_count"] != 4 ||
		preview["receipt_consumer_enablement_gate_modeled_item_count"] != 4 ||
		preview["consent_receipt_consumed_item_count"] != 0 ||
		preview["consumer_authorization_granted_item_count"] != 0 ||
		preview["authorized_kde_consumer_item_count"] != 0 ||
		preview["authorized_runtime_consumer_item_count"] != 0 ||
		preview["enabled_consumer_item_count"] != 0 ||
		preview["issued_notification_delivery_grant_item_count"] != 0 ||
		preview["side_effect_notification_delivery_consent_collection_receipt_consumer_enablement_gate_item_count"] != 0 {
		t.Fatalf("unexpected consent collection receipt consumer enablement gate counts: %#v", preview)
	}
	if !sameStrings(preview["notification_delivery_consent_collection_receipt_consumer_enablement_gate_item_ids"].([]string), []string{
		"route-enablement-lookup-route-dispatch-dry-run-result-notification-delivery-consent-collection-receipt-consumer-enablement-gate-install-failure",
		"route-enablement-lookup-route-dispatch-dry-run-result-notification-delivery-consent-collection-receipt-consumer-enablement-gate-repair-suggestion",
		"route-enablement-lookup-route-dispatch-dry-run-result-notification-delivery-consent-collection-receipt-consumer-enablement-gate-environment-switch",
		"route-enablement-lookup-route-dispatch-dry-run-result-notification-delivery-consent-collection-receipt-consumer-enablement-gate-approval-info",
	}) {
		t.Fatalf("unexpected consent collection receipt consumer enablement gate item ids: %#v", preview["notification_delivery_consent_collection_receipt_consumer_enablement_gate_item_ids"])
	}
	for _, item := range preview["notification_delivery_consent_collection_receipt_consumer_enablement_gate_items"].([]map[string]any) {
		if item["evidence_present"] != true ||
			item["notification_delivery_consent_collection_receipt_consumer_authorization_audit_consumed"] != true ||
			item["notification_delivery_consent_collection_receipt_consumption_gate_audit_consumed"] != true ||
			item["consumer_enablement_gate_required"] != true ||
			item["consumer_enablement_gate_modeled"] != true ||
			item["consumer_enablement_gate_ready"] != true ||
			item["consumer_enabled"] != false ||
			item["kde_consumer_enabled"] != false ||
			item["runtime_consumer_enabled"] != false ||
			item["consumer_authorization_granted"] != false ||
			item["consent_receipt_consumed"] != false ||
			item["delivery_grant_issued"] != false ||
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
			item["consent_collection_receipt_consumer_enablement_gate_status"] != "redacted-status-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-delivery-consent-collection-receipt-consumer-enablement-gate-ready-consumers-disabled" {
			t.Fatalf("unsafe consent collection receipt consumer enablement gate item: %#v", item)
		}
	}
	counts := preview["counts"].(map[string]int)
	if counts["total"] != 8 || counts["passed"] != 8 || counts["blocked"] != 0 || counts["pending"] != 0 {
		t.Fatalf("unexpected consent collection receipt consumer enablement gate checks: %#v", counts)
	}
	if !sameStrings(preview["check_ids"].([]string), []string{
		"current-mainline-consumed",
		"route-enablement-lookup-route-dispatch-dry-run-result-notification-delivery-consent-collection-receipt-consumer-authorization-audit-consumed",
		"route-enablement-lookup-route-dispatch-dry-run-result-notification-delivery-consent-collection-receipt-consumption-gate-audit-consumed",
		"lookup-route-dispatch-dry-run-result-notification-delivery-consent-collection-receipt-consumer-enablement-gate-modeled",
		"four-kde-notification-delivery-events-consent-collection-receipt-consumer-enablement-gate-review-only",
		"consumer-enablements-authorization-grants-consumption-delivery-and-events-disabled",
		"notification-actions-action-cards-and-raw-results-disabled",
		"production-and-host-boundary-closed",
	}) {
		t.Fatalf("unexpected consent collection receipt consumer enablement gate check ids: %#v", preview["check_ids"])
	}
	if err := validateNoBackendTerms(preview, "KDE-safe notification delivery consent collection receipt consumer enablement gate test"); err != nil {
		t.Fatalf("validateNoBackendTerms returned error: %v", err)
	}
}

func TestProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationDeliveryConsentCollectionReceiptConsumerEnablementGateAuditPreviewFailsClosedWithoutSources(t *testing.T) {
	root, err := os.MkdirTemp("", "xnix-receipt-consumer-enablement-gate-*")
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
	if err := os.WriteFile(filepath.Join(root, "docs", "xnix-current-mainline.md"), []byte("missing consent collection receipt consumer enablement gate evidence"), 0o644); err != nil {
		t.Fatalf("WriteFile mainline returned error: %v", err)
	}

	preview, err := NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationDeliveryConsentCollectionReceiptConsumerEnablementGateAuditPreview(root)
	if err != nil {
		t.Fatalf("NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationDeliveryConsentCollectionReceiptConsumerEnablementGateAuditPreview returned error: %v", err)
	}
	if preview["audit_decision"] != "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-delivery-consent-collection-receipt-consumer-enablement-gate-audit-blocked" ||
		preview["current_mainline_consumed"] != false ||
		preview["route_enablement_lookup_route_dispatch_dry_run_result_notification_delivery_consent_collection_receipt_consumer_authorization_audit_consumed"] != false ||
		preview["lookup_route_dispatch_dry_run_result_notification_delivery_consent_collection_receipt_consumer_enablement_gate_ready"] != false ||
		preview["ready_notification_delivery_consent_collection_receipt_consumer_enablement_gate_item_count"] != 0 ||
		preview["missing_notification_delivery_consent_collection_receipt_consumer_enablement_gate_item_count"] != 4 {
		t.Fatalf("expected consent collection receipt consumer enablement gate audit to fail closed without evidence: %#v", preview)
	}
	for _, item := range preview["notification_delivery_consent_collection_receipt_consumer_enablement_gate_items"].([]map[string]any) {
		if item["evidence_present"] != false ||
			item["consumer_enablement_gate_ready"] != false ||
			item["consumer_enabled"] != false ||
			item["kde_consumer_enabled"] != false ||
			item["runtime_consumer_enabled"] != false ||
			item["notification_sent"] != false ||
			item["host_root_modified"] != false {
			t.Fatalf("unsafe fail-closed consent collection receipt consumer enablement gate item: %#v", item)
		}
	}
}
