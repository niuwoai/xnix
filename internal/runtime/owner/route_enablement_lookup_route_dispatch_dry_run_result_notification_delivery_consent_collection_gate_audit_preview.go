package owner

const routeEnablementLookupRouteDispatchDryRunResultNotificationDeliveryConsentCollectionGateAuditRequest = "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-delivery-consent-collection-gate-audit-preview"

type ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationDeliveryConsentCollectionGateAuditPreview map[string]any

func NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationDeliveryConsentCollectionGateAuditPreview(root string) (ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationDeliveryConsentCollectionGateAuditPreview, error) {
	version, err := readProductionDBusGateReviewVersion(root)
	if err != nil {
		return nil, err
	}

	sources := routeEnablementLookupRouteDispatchDryRunResultNotificationDeliveryConsentCollectionGateAuditSources(root)
	mainlineReady := routeEnablementLookupRouteDispatchDryRunResultNotificationDeliveryConsentCollectionGateAuditMainlineReady(sources.CurrentMainline)
	explicitConsentReady := routeEnablementLookupRouteDispatchDryRunResultNotificationDeliveryConsentCollectionGateAuditExplicitConsentReady(sources.RouteEnablementLookupRouteDispatchDryRunResultNotificationDeliveryExplicitConsent)
	items := routeEnablementLookupRouteDispatchDryRunResultNotificationDeliveryConsentCollectionGateAuditItems(mainlineReady, explicitConsentReady)
	readyItemCount := routeEnablementLookupRouteDispatchDryRunResultNotificationDeliveryConsentCollectionGateAuditReadyCount(items)
	ready := mainlineReady && explicitConsentReady && readyItemCount == len(items)
	decision := "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-delivery-consent-collection-gate-audit-blocked"
	if ready {
		decision = "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-delivery-consent-collection-gate-audit-ready-collection-disabled"
	}

	preview := ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationDeliveryConsentCollectionGateAuditPreview{
		"version":                   version,
		"schema_version":            "xnix.runtime.production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_kde_safe_redacted_status_closed_record_writer_storage_root_authorization_receipt_dry_run_lookup_result_consumer_projection_route_enablement_lookup_route_dispatch_dry_run_result_notification_delivery_consent_collection_gate_audit.v1",
		"request_type":              routeEnablementLookupRouteDispatchDryRunResultNotificationDeliveryConsentCollectionGateAuditRequest,
		"audit_type":                "receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-delivery-consent-collection-gate-audit",
		"source":                    "docs/xnix-current-mainline.md+route-enable-lookup-route-dispatch-dry-run-result-notification-delivery-explicit-consent-audit",
		"audit_decision":            decision,
		"current_mainline_consumed": mainlineReady,
		"route_enablement_lookup_route_dispatch_dry_run_result_notification_delivery_explicit_consent_consumed":     explicitConsentReady,
		"route_enablement_lookup_route_dispatch_dry_run_result_notification_delivery_explicit_consent_ready":        explicitConsentReady,
		"lookup_route_dispatch_dry_run_result_notification_delivery_consent_collection_gate_audit_required":         true,
		"lookup_route_dispatch_dry_run_result_notification_delivery_consent_collection_gate_modeled":                true,
		"lookup_route_dispatch_dry_run_result_notification_delivery_consent_collection_gate_ready":                  ready,
		"route_enablement_lookup_route_dispatch_dry_run_result_notification_delivery_consent_collection_gate_ready": ready,
		"notification_delivery_explicit_consent_review_only_consumed":                                               explicitConsentReady,
		"kde_notification_delivery_consent_collection_gate_modeled":                                                 ready,
		"notification_center_delivery_consent_collection_gate_modeled":                                              ready,
		"install_failure_consent_collection_gate_modeled":                                                           ready,
		"repair_suggestion_consent_collection_gate_modeled":                                                         ready,
		"environment_switch_consent_collection_gate_modeled":                                                        ready,
		"approval_info_consent_collection_gate_modeled":                                                             ready,
		"notification_delivery_consent_collection_gate_review_only":                                                 true,
		"explicit_user_consent_required":                                                                            true,
		"explicit_user_consent_prepared":                                                                            ready,
		"explicit_user_consent_collection_gate_modeled":                                                             ready,
		"explicit_user_consent_collection_gate_ready":                                                               ready,
		"explicit_user_consent_collection_allowed":                                                                  false,
		"explicit_user_consent_collection_enabled":                                                                  false,
		"explicit_user_consent_collected":                                                                           false,
		"explicit_user_consent_persisted":                                                                           false,
		"consent_receipt_created":                                                                                   false,
		"consent_receipt_persisted":                                                                                 false,
		"notification_delivery_grant_issued":                                                                        false,
		"notification_delivery_authorization_granted":                                                               false,
		"notification_delivery_disabled":                                                                            true,
		"notification_action_disabled":                                                                              true,
		"action_cards_remain_disabled":                                                                              true,
		"kde_safe_redacted_result_only":                                                                             explicitConsentReady,
		"raw_result_hidden":                                                                                         explicitConsentReady,
		"notification_delivery_consent_collection_gate_item_count":                                                  len(items),
		"required_notification_delivery_consent_collection_gate_item_count":                                         len(items),
		"ready_notification_delivery_consent_collection_gate_item_count":                                            readyItemCount,
		"missing_notification_delivery_consent_collection_gate_item_count":                                          len(items) - readyItemCount,
		"prepared_explicit_consent_item_count":                                                                      readyItemCount,
		"collection_gate_modeled_item_count":                                                                        readyItemCount,
		"collected_explicit_consent_item_count":                                                                     0,
		"persisted_explicit_consent_item_count":                                                                     0,
		"consent_receipt_created_item_count":                                                                        0,
		"issued_notification_delivery_grant_item_count":                                                             0,
		"authorized_notification_delivery_item_count":                                                               0,
		"delivered_notification_item_count":                                                                         0,
		"notification_action_enabled_item_count":                                                                    0,
		"action_card_enabled_item_count":                                                                            0,
		"side_effect_notification_delivery_consent_collection_gate_item_count":                                      0,
		"notification_delivery_consent_collection_gate_items":                                                       items,
		"notification_delivery_consent_collection_gate_item_ids":                                                    routeEnablementLookupRouteDispatchDryRunResultNotificationDeliveryConsentCollectionGateAuditItemIDs(items),
		"runtime_owned":                       true,
		"go_runtime_backed":                   true,
		"kde_policy_owner":                    false,
		"receipt_present":                     false,
		"receipt_accepted":                    false,
		"receipt_consumed":                    false,
		"route_enablement_accepted":           false,
		"lookup_route_enabled":                false,
		"lookup_route_dispatch_authorized":    false,
		"lookup_route_dispatch_callable":      false,
		"storage_write_enabled":               false,
		"status_persistence_write_enabled":    false,
		"redacted_summary_persisted":          false,
		"kde_status_persisted":                false,
		"runtime_diagnostics_persisted":       false,
		"dry_run_result_persisted":            false,
		"raw_result_exposed":                  false,
		"dispatch_dry_run_executed":           false,
		"request_object_creation_enabled":     false,
		"request_object_dispatch_enabled":     false,
		"portal_request_created":              false,
		"notification_delivery_enabled":       false,
		"notification_sent":                   false,
		"notification_center_event_triggered": false,
		"notification_action_enabled":         false,
		"action_card_enabled":                 false,
		"compatibility_center_opened":         false,
		"support_bundle_exported":             false,
		"support_case_created":                false,
		"production_readiness":                false,
		"production_ownership_ready":          false,
		"system_service_started":              false,
		"session_bus_claimed":                 false,
		"production_bus_claimed":              false,
		"write_methods_enabled":               false,
		"runtime_writes_enabled":              false,
		"desktop_files_written":               false,
		"kde_configuration_written":           false,
		"portal_call_executed":                false,
		"adapter_invocation_enabled":          false,
		"backend_launch_enabled":              false,
		"backend_process_started":             false,
		"network_required":                    false,
		"host_root_modified":                  false,
		"privileged_container_required":       false,
		"state_root_path_exposed":             false,
		"file_paths_exposed":                  false,
		"file_content_read":                   false,
		"raw_command_exposed":                 false,
		"raw_executable_exposed":              false,
		"backend_details_exposed":             false,
		"blocked_actions": []string{
			"explicit-consent-collection-allow",
			"explicit-consent-collection-enable",
			"explicit-consent-collection",
			"explicit-consent-persistence",
			"consent-receipt-create",
			"consent-receipt-persist",
			"notification-delivery-grant-issuance",
			"delivery-authorization-grant",
			"notification-delivery",
			"notification-center-event",
			"notification-action",
			"action-card-enable",
			"receipt-present",
			"receipt-accept",
			"receipt-consume",
			"lookup-route-enable",
			"lookup-route-dispatch",
			"dry-run-dispatch",
			"result-persistence",
			"raw-result-exposure",
			"production-ownership",
			"backend-launch",
			"host-mutation",
		},
		"next_requirements": []string{
			"Add a separate consent collection receipt preview before user consent can be collected or persisted.",
			"Keep delivery grant issuance disabled until consent collection receipts, delivery execution, production ownership, and host boundaries are separately authorized.",
		},
		"desktop_safe_summary": "KDE Notification Center consent collection gates are modeled for redacted lookup route dispatch dry-run result install failure, repair suggestion, environment switch, and approval information events while consent collection, consent persistence, consent receipts, grant issuance, delivery, events, actions, action cards, raw results, receipts, lookup routes, dispatch, persistence, production, backend, and host side effects remain disabled.",
	}
	preview["checks"] = routeEnablementLookupRouteDispatchDryRunResultNotificationDeliveryConsentCollectionGateAuditChecks(preview)
	preview["check_ids"] = routeEnablementLookupRouteDispatchDryRunResultNotificationDeliveryConsentCollectionGateAuditCheckIDs(preview["checks"].([]map[string]any))
	preview["counts"] = routeEnablementLookupRouteDispatchDryRunResultNotificationDeliveryConsentCollectionGateAuditCountsFor(preview["checks"].([]map[string]any))
	if err := validateNoBackendTerms(preview, "KDE-safe redacted status dry-run lookup result consumer projection route enablement lookup route dispatch dry-run result notification delivery consent collection gate audit preview"); err != nil {
		return nil, err
	}
	return preview, nil
}

type routeEnablementLookupRouteDispatchDryRunResultNotificationDeliveryConsentCollectionGateAuditSourceSet struct {
	CurrentMainline                                                                   string
	RouteEnablementLookupRouteDispatchDryRunResultNotificationDeliveryExplicitConsent string
}

func routeEnablementLookupRouteDispatchDryRunResultNotificationDeliveryConsentCollectionGateAuditSources(root string) routeEnablementLookupRouteDispatchDryRunResultNotificationDeliveryConsentCollectionGateAuditSourceSet {
	return routeEnablementLookupRouteDispatchDryRunResultNotificationDeliveryConsentCollectionGateAuditSourceSet{
		CurrentMainline: productionAuthorizationReadSources(root, []string{"docs/xnix-current-mainline.md"}),
		RouteEnablementLookupRouteDispatchDryRunResultNotificationDeliveryExplicitConsent: productionAuthorizationReadSources(root, []string{
			"internal/runtime/owner/route_enablement_lookup_route_dispatch_dry_run_result_notification_delivery_explicit_consent_audit_preview.go",
			"internal/runtime/owner/route_enablement_lookup_route_dispatch_dry_run_result_notification_delivery_explicit_consent_audit_preview_test.go",
		}),
	}
}

func routeEnablementLookupRouteDispatchDryRunResultNotificationDeliveryConsentCollectionGateAuditMainlineReady(source string) bool {
	return productionAuthorizationHasAll(source, []string{
		"route enablement lookup route dispatch dry-run result notification delivery consent collection gate audit preview",
		"route enablement lookup route dispatch dry-run result notification delivery explicit consent audit preview",
		"without sending notifications",
		"collecting consent",
	})
}

func routeEnablementLookupRouteDispatchDryRunResultNotificationDeliveryConsentCollectionGateAuditExplicitConsentReady(source string) bool {
	return productionAuthorizationHasAll(source, []string{
		"xnix.runtime.production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_kde_safe_redacted_status_closed_record_writer_storage_root_authorization_receipt_dry_run_lookup_result_consumer_projection_route_enablement_lookup_route_dispatch_dry_run_result_notification_delivery_explicit_consent_audit.v1",
		"production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-delivery-explicit-consent-audit-ready-consent-collection-disabled",
		"explicit_user_consent_prepared",
		"explicit_user_consent_collection_enabled",
		"notification_delivery_explicit_consent_review_only",
	})
}

func routeEnablementLookupRouteDispatchDryRunResultNotificationDeliveryConsentCollectionGateAuditItems(mainlineReady, explicitConsentReady bool) []map[string]any {
	kinds := []struct {
		id   string
		kind string
	}{
		{id: "install-failure", kind: "notification-center-install-failure-consent-collection-gate"},
		{id: "repair-suggestion", kind: "notification-center-repair-suggestion-consent-collection-gate"},
		{id: "environment-switch", kind: "notification-center-environment-switch-consent-collection-gate"},
		{id: "approval-info", kind: "notification-center-approval-info-consent-collection-gate"},
	}
	items := make([]map[string]any, 0, len(kinds))
	for _, notification := range kinds {
		ready := mainlineReady && explicitConsentReady
		status := "missing-redacted-status-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-delivery-consent-collection-gate-evidence"
		if ready {
			status = "redacted-status-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-delivery-consent-collection-gate-ready-collection-disabled"
		}
		items = append(items, map[string]any{
			"id":                        "route-enablement-lookup-route-dispatch-dry-run-result-notification-delivery-consent-collection-gate-" + notification.id,
			"notification_kind":         notification.id,
			"gate_kind":                 notification.kind,
			"evidence_present":          ready,
			"current_mainline_consumed": mainlineReady,
			"notification_delivery_explicit_consent_consumed": explicitConsentReady,
			"notification_delivery_explicit_consent_ready":    explicitConsentReady,
			"explicit_consent_required":                       true,
			"explicit_consent_prepared":                       ready,
			"consent_collection_gate_modeled":                 ready,
			"consent_collection_gate_ready":                   ready,
			"consent_collection_allowed":                      false,
			"consent_collection_enabled":                      false,
			"explicit_user_consent_collected":                 false,
			"explicit_user_consent_persisted":                 false,
			"consent_receipt_created":                         false,
			"consent_receipt_persisted":                       false,
			"delivery_grant_issued":                           false,
			"delivery_authorization_granted":                  false,
			"delivery_review_only":                            true,
			"kde_safe_redacted_result_only":                   explicitConsentReady,
			"raw_result_hidden":                               explicitConsentReady,
			"notification_delivery_enabled":                   false,
			"notification_sent":                               false,
			"notification_center_event_triggered":             false,
			"notification_action_enabled":                     false,
			"action_card_enabled":                             false,
			"side_effects_disabled":                           true,
			"runtime_owned":                                   true,
			"go_runtime_backed":                               true,
			"kde_policy_owner":                                false,
			"host_root_modified":                              false,
			"internal_details_exposed":                        false,
			"consent_collection_gate_status":                  status,
			"next_requirement":                                "Require a separate consent collection receipt preview before this gate can collect or persist user consent.",
		})
	}
	return items
}

func routeEnablementLookupRouteDispatchDryRunResultNotificationDeliveryConsentCollectionGateAuditReadyCount(items []map[string]any) int {
	count := 0
	for _, item := range items {
		if item["evidence_present"] == true && item["explicit_consent_required"] == true && item["explicit_consent_prepared"] == true && item["consent_collection_gate_modeled"] == true && item["consent_collection_gate_ready"] == true && item["consent_collection_allowed"] == false && item["consent_collection_enabled"] == false && item["explicit_user_consent_collected"] == false && item["explicit_user_consent_persisted"] == false && item["consent_receipt_created"] == false && item["consent_receipt_persisted"] == false && item["delivery_grant_issued"] == false && item["delivery_authorization_granted"] == false && item["delivery_review_only"] == true && item["kde_safe_redacted_result_only"] == true && item["raw_result_hidden"] == true && item["side_effects_disabled"] == true && item["notification_delivery_enabled"] == false && item["notification_sent"] == false && item["notification_center_event_triggered"] == false && item["notification_action_enabled"] == false && item["action_card_enabled"] == false && item["host_root_modified"] == false {
			count++
		}
	}
	return count
}

func routeEnablementLookupRouteDispatchDryRunResultNotificationDeliveryConsentCollectionGateAuditItemIDs(items []map[string]any) []string {
	ids := make([]string, 0, len(items))
	for _, item := range items {
		ids = append(ids, item["id"].(string))
	}
	return ids
}

func routeEnablementLookupRouteDispatchDryRunResultNotificationDeliveryConsentCollectionGateAuditChecks(preview ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationDeliveryConsentCollectionGateAuditPreview) []map[string]any {
	return []map[string]any{
		routeEnablementLookupRouteDispatchDryRunResultNotificationDeliveryConsentCollectionGateAuditCheck("current-mainline-consumed", preview["current_mainline_consumed"] == true, "Current mainline names the notification delivery consent collection gate continuation."),
		routeEnablementLookupRouteDispatchDryRunResultNotificationDeliveryConsentCollectionGateAuditCheck("route-enablement-lookup-route-dispatch-dry-run-result-notification-delivery-explicit-consent-consumed", preview["route_enablement_lookup_route_dispatch_dry_run_result_notification_delivery_explicit_consent_consumed"] == true && preview["route_enablement_lookup_route_dispatch_dry_run_result_notification_delivery_explicit_consent_ready"] == true, "Notification delivery explicit consent audit is consumed as predecessor evidence."),
		routeEnablementLookupRouteDispatchDryRunResultNotificationDeliveryConsentCollectionGateAuditCheck("lookup-route-dispatch-dry-run-result-notification-delivery-consent-collection-gate-modeled", preview["lookup_route_dispatch_dry_run_result_notification_delivery_consent_collection_gate_modeled"] == true && preview["lookup_route_dispatch_dry_run_result_notification_delivery_consent_collection_gate_ready"] == true, "Notification delivery consent collection gate boundary is modeled."),
		routeEnablementLookupRouteDispatchDryRunResultNotificationDeliveryConsentCollectionGateAuditCheck("four-kde-notification-delivery-events-consent-collection-gate-review-only", preview["notification_delivery_consent_collection_gate_item_count"] == 4 && preview["ready_notification_delivery_consent_collection_gate_item_count"] == 4 && preview["collected_explicit_consent_item_count"] == 0 && preview["persisted_explicit_consent_item_count"] == 0 && preview["delivered_notification_item_count"] == 0, "Four KDE Notification Center delivery event classes are consent-collection-gate-modeled and remain review-only."),
		routeEnablementLookupRouteDispatchDryRunResultNotificationDeliveryConsentCollectionGateAuditCheck("consent-collection-persistence-grant-issuance-sending-and-events-disabled", preview["explicit_user_consent_collection_allowed"] == false && preview["explicit_user_consent_collection_enabled"] == false && preview["explicit_user_consent_collected"] == false && preview["explicit_user_consent_persisted"] == false && preview["consent_receipt_created"] == false && preview["consent_receipt_persisted"] == false && preview["notification_delivery_grant_issued"] == false && preview["notification_delivery_authorization_granted"] == false && preview["notification_delivery_enabled"] == false && preview["notification_sent"] == false && preview["notification_center_event_triggered"] == false, "Consent collection, persistence, receipt creation, grant issuance, notification sending, and Notification Center events remain disabled."),
		routeEnablementLookupRouteDispatchDryRunResultNotificationDeliveryConsentCollectionGateAuditCheck("notification-actions-action-cards-and-raw-results-disabled", preview["notification_action_disabled"] == true && preview["action_cards_remain_disabled"] == true && preview["notification_action_enabled"] == false && preview["action_card_enabled"] == false && preview["raw_result_hidden"] == true && preview["raw_result_exposed"] == false && preview["dry_run_result_persisted"] == false && preview["storage_write_enabled"] == false, "Notification actions, action cards, raw result exposure, and persistence remain disabled."),
		routeEnablementLookupRouteDispatchDryRunResultNotificationDeliveryConsentCollectionGateAuditCheck("production-and-host-boundary-closed", preview["production_ownership_ready"] == false && preview["production_bus_claimed"] == false && preview["write_methods_enabled"] == false && preview["runtime_writes_enabled"] == false && preview["backend_launch_enabled"] == false && preview["host_root_modified"] == false, "Production and host boundaries remain closed."),
	}
}

func routeEnablementLookupRouteDispatchDryRunResultNotificationDeliveryConsentCollectionGateAuditCheck(id string, passed bool, summary string) map[string]any {
	status := "blocked"
	if passed {
		status = "pass"
	}
	return map[string]any{"id": id, "status": status, "summary": summary}
}

func routeEnablementLookupRouteDispatchDryRunResultNotificationDeliveryConsentCollectionGateAuditCheckIDs(checks []map[string]any) []string {
	ids := make([]string, 0, len(checks))
	for _, check := range checks {
		ids = append(ids, check["id"].(string))
	}
	return ids
}

func routeEnablementLookupRouteDispatchDryRunResultNotificationDeliveryConsentCollectionGateAuditCountsFor(checks []map[string]any) map[string]int {
	counts := map[string]int{"total": len(checks), "passed": 0, "pending": 0, "blocked": 0}
	for _, check := range checks {
		if check["status"] == "pass" {
			counts["passed"]++
		} else {
			counts["blocked"]++
		}
	}
	return counts
}
