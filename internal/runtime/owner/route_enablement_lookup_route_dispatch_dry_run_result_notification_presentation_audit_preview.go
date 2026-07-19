package owner

type ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationPresentationAuditPreview struct {
	Version                                                                     string                                                                                                                                                                                                                                                                  `json:"version"`
	SchemaVersion                                                               string                                                                                                                                                                                                                                                                  `json:"schema_version"`
	RequestType                                                                 string                                                                                                                                                                                                                                                                  `json:"request_type"`
	AuditType                                                                   string                                                                                                                                                                                                                                                                  `json:"audit_type"`
	Source                                                                      string                                                                                                                                                                                                                                                                  `json:"source"`
	AuditDecision                                                               string                                                                                                                                                                                                                                                                  `json:"audit_decision"`
	CurrentMainlineConsumed                                                     bool                                                                                                                                                                                                                                                                    `json:"current_mainline_consumed"`
	RouteEnablementLookupRouteDispatchDryRunResultPresentationConsentConsumed   bool                                                                                                                                                                                                                                                                    `json:"route_enablement_lookup_route_dispatch_dry_run_result_presentation_consent_consumed"`
	RouteEnablementLookupRouteDispatchDryRunResultPresentationConsentReady      bool                                                                                                                                                                                                                                                                    `json:"route_enablement_lookup_route_dispatch_dry_run_result_presentation_consent_ready"`
	LookupRouteDispatchDryRunResultNotificationPresentationRequired             bool                                                                                                                                                                                                                                                                    `json:"lookup_route_dispatch_dry_run_result_notification_presentation_required"`
	LookupRouteDispatchDryRunResultNotificationPresentationModeled              bool                                                                                                                                                                                                                                                                    `json:"lookup_route_dispatch_dry_run_result_notification_presentation_modeled"`
	LookupRouteDispatchDryRunResultNotificationPresentationReady                bool                                                                                                                                                                                                                                                                    `json:"lookup_route_dispatch_dry_run_result_notification_presentation_ready"`
	RouteEnablementLookupRouteDispatchDryRunResultNotificationPresentationReady bool                                                                                                                                                                                                                                                                    `json:"route_enablement_lookup_route_dispatch_dry_run_result_notification_presentation_ready"`
	ConsentModeledRedactedPresentationConsumed                                  bool                                                                                                                                                                                                                                                                    `json:"consent_modeled_redacted_presentation_consumed"`
	KDENotificationPresentationModeled                                          bool                                                                                                                                                                                                                                                                    `json:"kde_notification_presentation_modeled"`
	NotificationCenterSurfaceModeled                                            bool                                                                                                                                                                                                                                                                    `json:"notification_center_surface_modeled"`
	InstallFailureNotificationModeled                                           bool                                                                                                                                                                                                                                                                    `json:"install_failure_notification_modeled"`
	RepairSuggestionNotificationModeled                                         bool                                                                                                                                                                                                                                                                    `json:"repair_suggestion_notification_modeled"`
	EnvironmentSwitchNotificationModeled                                        bool                                                                                                                                                                                                                                                                    `json:"environment_switch_notification_modeled"`
	ApprovalInfoNotificationModeled                                             bool                                                                                                                                                                                                                                                                    `json:"approval_info_notification_modeled"`
	NotificationPresentationReviewOnly                                          bool                                                                                                                                                                                                                                                                    `json:"notification_presentation_review_only"`
	NotificationDeliveryDisabled                                                bool                                                                                                                                                                                                                                                                    `json:"notification_delivery_disabled"`
	NotificationActionDisabled                                                  bool                                                                                                                                                                                                                                                                    `json:"notification_action_disabled"`
	ActionCardsRemainDisabled                                                   bool                                                                                                                                                                                                                                                                    `json:"action_cards_remain_disabled"`
	ExplicitUserConsentRequired                                                 bool                                                                                                                                                                                                                                                                    `json:"explicit_user_consent_required"`
	ExplicitUserConsentCollected                                                bool                                                                                                                                                                                                                                                                    `json:"explicit_user_consent_collected"`
	KDESafeRedactedResultOnly                                                   bool                                                                                                                                                                                                                                                                    `json:"kde_safe_redacted_result_only"`
	RawResultHidden                                                             bool                                                                                                                                                                                                                                                                    `json:"raw_result_hidden"`
	NotificationPresentationItemCount                                           int                                                                                                                                                                                                                                                                     `json:"notification_presentation_item_count"`
	RequiredNotificationPresentationItemCount                                   int                                                                                                                                                                                                                                                                     `json:"required_notification_presentation_item_count"`
	ReadyNotificationPresentationItemCount                                      int                                                                                                                                                                                                                                                                     `json:"ready_notification_presentation_item_count"`
	MissingNotificationPresentationItemCount                                    int                                                                                                                                                                                                                                                                     `json:"missing_notification_presentation_item_count"`
	DeliveredNotificationItemCount                                              int                                                                                                                                                                                                                                                                     `json:"delivered_notification_item_count"`
	NotificationActionEnabledItemCount                                          int                                                                                                                                                                                                                                                                     `json:"notification_action_enabled_item_count"`
	ActionCardEnabledItemCount                                                  int                                                                                                                                                                                                                                                                     `json:"action_card_enabled_item_count"`
	SideEffectNotificationPresentationItemCount                                 int                                                                                                                                                                                                                                                                     `json:"side_effect_notification_presentation_item_count"`
	NotificationPresentationItems                                               []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationPresentationAuditItem      `json:"notification_presentation_items"`
	NotificationPresentationItemIDs                                             []string                                                                                                                                                                                                                                                                `json:"notification_presentation_item_ids"`
	Checks                                                                      []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationPresentationAuditCheck     `json:"checks"`
	CheckIDs                                                                    []string                                                                                                                                                                                                                                                                `json:"check_ids"`
	Counts                                                                      ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationPresentationAuditCheckCounts `json:"counts"`
	ReceiptPresent                                                              bool                                                                                                                                                                                                                                                                    `json:"receipt_present"`
	ReceiptAccepted                                                             bool                                                                                                                                                                                                                                                                    `json:"receipt_accepted"`
	ReceiptConsumed                                                             bool                                                                                                                                                                                                                                                                    `json:"receipt_consumed"`
	RouteEnablementAccepted                                                     bool                                                                                                                                                                                                                                                                    `json:"route_enablement_accepted"`
	LookupRouteEnabled                                                          bool                                                                                                                                                                                                                                                                    `json:"lookup_route_enabled"`
	LookupRouteDispatchAuthorized                                               bool                                                                                                                                                                                                                                                                    `json:"lookup_route_dispatch_authorized"`
	LookupRouteDispatchCallable                                                 bool                                                                                                                                                                                                                                                                    `json:"lookup_route_dispatch_callable"`
	StorageWriteEnabled                                                         bool                                                                                                                                                                                                                                                                    `json:"storage_write_enabled"`
	StatusPersistenceWriteEnabled                                               bool                                                                                                                                                                                                                                                                    `json:"status_persistence_write_enabled"`
	RedactedSummaryPersisted                                                    bool                                                                                                                                                                                                                                                                    `json:"redacted_summary_persisted"`
	KDEStatusPersisted                                                          bool                                                                                                                                                                                                                                                                    `json:"kde_status_persisted"`
	RuntimeDiagnosticsPersisted                                                 bool                                                                                                                                                                                                                                                                    `json:"runtime_diagnostics_persisted"`
	DryRunResultPersisted                                                       bool                                                                                                                                                                                                                                                                    `json:"dry_run_result_persisted"`
	RawResultExposed                                                            bool                                                                                                                                                                                                                                                                    `json:"raw_result_exposed"`
	DispatchDryRunExecuted                                                      bool                                                                                                                                                                                                                                                                    `json:"dispatch_dry_run_executed"`
	RequestObjectCreationEnabled                                                bool                                                                                                                                                                                                                                                                    `json:"request_object_creation_enabled"`
	RequestObjectDispatchEnabled                                                bool                                                                                                                                                                                                                                                                    `json:"request_object_dispatch_enabled"`
	PortalRequestCreated                                                        bool                                                                                                                                                                                                                                                                    `json:"portal_request_created"`
	NotificationDeliveryEnabled                                                 bool                                                                                                                                                                                                                                                                    `json:"notification_delivery_enabled"`
	NotificationSent                                                            bool                                                                                                                                                                                                                                                                    `json:"notification_sent"`
	NotificationCenterEventTriggered                                            bool                                                                                                                                                                                                                                                                    `json:"notification_center_event_triggered"`
	NotificationActionEnabled                                                   bool                                                                                                                                                                                                                                                                    `json:"notification_action_enabled"`
	ActionCardEnabled                                                           bool                                                                                                                                                                                                                                                                    `json:"action_card_enabled"`
	CompatibilityCenterOpened                                                   bool                                                                                                                                                                                                                                                                    `json:"compatibility_center_opened"`
	SupportBundleExported                                                       bool                                                                                                                                                                                                                                                                    `json:"support_bundle_exported"`
	SupportCaseCreated                                                          bool                                                                                                                                                                                                                                                                    `json:"support_case_created"`
	RuntimeOwned                                                                bool                                                                                                                                                                                                                                                                    `json:"runtime_owned"`
	GoRuntimeBacked                                                             bool                                                                                                                                                                                                                                                                    `json:"go_runtime_backed"`
	KDEPolicyOwner                                                              bool                                                                                                                                                                                                                                                                    `json:"kde_policy_owner"`
	ProductionReadiness                                                         bool                                                                                                                                                                                                                                                                    `json:"production_readiness"`
	ProductionOwnershipReady                                                    bool                                                                                                                                                                                                                                                                    `json:"production_ownership_ready"`
	SystemServiceStarted                                                        bool                                                                                                                                                                                                                                                                    `json:"system_service_started"`
	SessionBusClaimed                                                           bool                                                                                                                                                                                                                                                                    `json:"session_bus_claimed"`
	ProductionBusClaimed                                                        bool                                                                                                                                                                                                                                                                    `json:"production_bus_claimed"`
	WriteMethodsEnabled                                                         bool                                                                                                                                                                                                                                                                    `json:"write_methods_enabled"`
	RuntimeWritesEnabled                                                        bool                                                                                                                                                                                                                                                                    `json:"runtime_writes_enabled"`
	DesktopFilesWritten                                                         bool                                                                                                                                                                                                                                                                    `json:"desktop_files_written"`
	KDEConfigurationWritten                                                     bool                                                                                                                                                                                                                                                                    `json:"kde_configuration_written"`
	PortalCallExecuted                                                          bool                                                                                                                                                                                                                                                                    `json:"portal_call_executed"`
	AdapterInvocationEnabled                                                    bool                                                                                                                                                                                                                                                                    `json:"adapter_invocation_enabled"`
	BackendLaunchEnabled                                                        bool                                                                                                                                                                                                                                                                    `json:"backend_launch_enabled"`
	BackendProcessStarted                                                       bool                                                                                                                                                                                                                                                                    `json:"backend_process_started"`
	NetworkRequired                                                             bool                                                                                                                                                                                                                                                                    `json:"network_required"`
	HostRootModified                                                            bool                                                                                                                                                                                                                                                                    `json:"host_root_modified"`
	PrivilegedContainerRequired                                                 bool                                                                                                                                                                                                                                                                    `json:"privileged_container_required"`
	StateRootPathExposed                                                        bool                                                                                                                                                                                                                                                                    `json:"state_root_path_exposed"`
	FilePathsExposed                                                            bool                                                                                                                                                                                                                                                                    `json:"file_paths_exposed"`
	FileContentRead                                                             bool                                                                                                                                                                                                                                                                    `json:"file_content_read"`
	RawCommandExposed                                                           bool                                                                                                                                                                                                                                                                    `json:"raw_command_exposed"`
	RawExecutableExposed                                                        bool                                                                                                                                                                                                                                                                    `json:"raw_executable_exposed"`
	BackendDetailsExposed                                                       bool                                                                                                                                                                                                                                                                    `json:"backend_details_exposed"`
	BlockedActions                                                              []string                                                                                                                                                                                                                                                                `json:"blocked_actions"`
	NextRequirements                                                            []string                                                                                                                                                                                                                                                                `json:"next_requirements"`
	DesktopSafeSummary                                                          string                                                                                                                                                                                                                                                                  `json:"desktop_safe_summary"`
}

type ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationPresentationAuditItem struct {
	ID                                                                        string `json:"id"`
	NotificationKind                                                          string `json:"notification_kind"`
	PresentationKind                                                          string `json:"presentation_kind"`
	EvidencePresent                                                           bool   `json:"evidence_present"`
	CurrentMainlineConsumed                                                   bool   `json:"current_mainline_consumed"`
	RouteEnablementLookupRouteDispatchDryRunResultPresentationConsentConsumed bool   `json:"route_enablement_lookup_route_dispatch_dry_run_result_presentation_consent_consumed"`
	RouteEnablementLookupRouteDispatchDryRunResultPresentationConsentReady    bool   `json:"route_enablement_lookup_route_dispatch_dry_run_result_presentation_consent_ready"`
	NotificationPresentationRequired                                          bool   `json:"notification_presentation_required"`
	NotificationPresentationModeled                                           bool   `json:"notification_presentation_modeled"`
	NotificationPresentationReady                                             bool   `json:"notification_presentation_ready"`
	ExplicitUserConsentRequired                                               bool   `json:"explicit_user_consent_required"`
	ExplicitUserConsentCollected                                              bool   `json:"explicit_user_consent_collected"`
	NotificationReviewOnly                                                    bool   `json:"notification_review_only"`
	KDESafeRedactedResultOnly                                                 bool   `json:"kde_safe_redacted_result_only"`
	RawResultHidden                                                           bool   `json:"raw_result_hidden"`
	NotificationDeliveryEnabled                                               bool   `json:"notification_delivery_enabled"`
	NotificationSent                                                          bool   `json:"notification_sent"`
	NotificationActionEnabled                                                 bool   `json:"notification_action_enabled"`
	ActionCardEnabled                                                         bool   `json:"action_card_enabled"`
	SideEffectsDisabled                                                       bool   `json:"side_effects_disabled"`
	RuntimeOwned                                                              bool   `json:"runtime_owned"`
	GoRuntimeBacked                                                           bool   `json:"go_runtime_backed"`
	KDEPolicyOwner                                                            bool   `json:"kde_policy_owner"`
	HostRootModified                                                          bool   `json:"host_root_modified"`
	InternalDetailsExposed                                                    bool   `json:"internal_details_exposed"`
	NotificationPresentationStatus                                            string `json:"notification_presentation_status"`
	NextRequirement                                                           string `json:"next_requirement"`
}

type ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationPresentationAuditCheck struct {
	ID      string `json:"id"`
	Status  string `json:"status"`
	Summary string `json:"summary"`
}

type ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationPresentationAuditCheckCounts struct {
	Total   int `json:"total"`
	Passed  int `json:"passed"`
	Pending int `json:"pending"`
	Blocked int `json:"blocked"`
}

type routeEnablementLookupRouteDispatchDryRunResultNotificationPresentationAuditSourceSet struct {
	CurrentMainline                                                   string
	RouteEnablementLookupRouteDispatchDryRunResultPresentationConsent string
}

func NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationPresentationAuditPreview(root string) (ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationPresentationAuditPreview, error) {
	version, err := readProductionDBusGateReviewVersion(root)
	if err != nil {
		return ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationPresentationAuditPreview{}, err
	}
	sources := routeEnablementLookupRouteDispatchDryRunResultNotificationPresentationAuditSources(root)
	mainlineReady := routeEnablementLookupRouteDispatchDryRunResultNotificationPresentationAuditMainlineReady(sources.CurrentMainline)
	consentReady := routeEnablementLookupRouteDispatchDryRunResultNotificationPresentationAuditConsentReady(sources.RouteEnablementLookupRouteDispatchDryRunResultPresentationConsent)
	items := routeEnablementLookupRouteDispatchDryRunResultNotificationPresentationAuditItems(mainlineReady, consentReady)
	readyItemCount := routeEnablementLookupRouteDispatchDryRunResultNotificationPresentationAuditReadyCount(items)
	ready := mainlineReady && consentReady && readyItemCount == len(items)
	preview := ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationPresentationAuditPreview{
		Version:                 version,
		SchemaVersion:           "xnix.runtime.production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_kde_safe_redacted_status_closed_record_writer_storage_root_authorization_receipt_dry_run_lookup_result_consumer_projection_route_enablement_lookup_route_dispatch_dry_run_result_notification_presentation_audit.v1",
		RequestType:             "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-presentation-audit-preview",
		AuditType:               "receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-presentation-audit",
		Source:                  "docs/xnix-current-mainline.md+route-enable-lookup-route-dispatch-dry-run-result-presentation-consent-audit",
		AuditDecision:           "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-presentation-audit-blocked",
		CurrentMainlineConsumed: mainlineReady,
		RouteEnablementLookupRouteDispatchDryRunResultPresentationConsentConsumed:   consentReady,
		RouteEnablementLookupRouteDispatchDryRunResultPresentationConsentReady:      consentReady,
		LookupRouteDispatchDryRunResultNotificationPresentationRequired:             true,
		LookupRouteDispatchDryRunResultNotificationPresentationModeled:              true,
		LookupRouteDispatchDryRunResultNotificationPresentationReady:                ready,
		RouteEnablementLookupRouteDispatchDryRunResultNotificationPresentationReady: ready,
		ConsentModeledRedactedPresentationConsumed:                                  consentReady,
		KDENotificationPresentationModeled:                                          ready,
		NotificationCenterSurfaceModeled:                                            ready,
		InstallFailureNotificationModeled:                                           ready,
		RepairSuggestionNotificationModeled:                                         ready,
		EnvironmentSwitchNotificationModeled:                                        ready,
		ApprovalInfoNotificationModeled:                                             ready,
		NotificationPresentationReviewOnly:                                          true,
		NotificationDeliveryDisabled:                                                true,
		NotificationActionDisabled:                                                  true,
		ActionCardsRemainDisabled:                                                   true,
		ExplicitUserConsentRequired:                                                 true,
		ExplicitUserConsentCollected:                                                false,
		KDESafeRedactedResultOnly:                                                   consentReady,
		RawResultHidden:                                                             consentReady,
		NotificationPresentationItemCount:                                           len(items),
		RequiredNotificationPresentationItemCount:                                   len(items),
		ReadyNotificationPresentationItemCount:                                      readyItemCount,
		MissingNotificationPresentationItemCount:                                    len(items) - readyItemCount,
		NotificationPresentationItems:                                               items,
		NotificationPresentationItemIDs:                                             routeEnablementLookupRouteDispatchDryRunResultNotificationPresentationAuditItemIDs(items),
		RuntimeOwned:                                                                true,
		GoRuntimeBacked:                                                             true,
		KDEPolicyOwner:                                                              false,
		BlockedActions: []string{
			"receipt-present",
			"receipt-accept",
			"receipt-consume",
			"lookup-route-enable",
			"lookup-route-dispatch",
			"dry-run-dispatch",
			"result-persistence",
			"raw-result-exposure",
			"notification-delivery",
			"notification-action",
			"notification-center-event",
			"action-card-enable",
			"production-ownership",
			"backend-launch",
			"host-mutation",
		},
		NextRequirements: []string{
			"Add a separate notification delivery authorization audit before Notification Center entries can be delivered.",
			"Keep notification presentation modeled and review-only until explicit user consent collection and production ownership are separately authorized.",
		},
		DesktopSafeSummary: "KDE Notification Center presentation for consent-modeled redacted lookup route dispatch dry-run results is modeled for install failure, repair suggestion, environment switch, and approval information events while notification delivery, notification actions, action cards, raw results, receipts, lookup routes, dispatch, persistence, production, backend, and host side effects remain disabled.",
	}
	if ready {
		preview.AuditDecision = "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-presentation-audit-ready-notification-review-only-delivery-disabled"
	}
	preview.Checks = routeEnablementLookupRouteDispatchDryRunResultNotificationPresentationAuditChecks(preview)
	preview.CheckIDs = routeEnablementLookupRouteDispatchDryRunResultNotificationPresentationAuditCheckIDs(preview.Checks)
	preview.Counts = routeEnablementLookupRouteDispatchDryRunResultNotificationPresentationAuditCountsFor(preview.Checks)
	if err := validateNoBackendTerms(preview, "KDE-safe redacted status dry-run lookup result consumer projection route enablement lookup route dispatch dry-run result notification presentation audit preview"); err != nil {
		return ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationPresentationAuditPreview{}, err
	}
	return preview, nil
}

func routeEnablementLookupRouteDispatchDryRunResultNotificationPresentationAuditSources(root string) routeEnablementLookupRouteDispatchDryRunResultNotificationPresentationAuditSourceSet {
	return routeEnablementLookupRouteDispatchDryRunResultNotificationPresentationAuditSourceSet{
		CurrentMainline: productionAuthorizationReadSources(root, []string{"docs/xnix-current-mainline.md"}),
		RouteEnablementLookupRouteDispatchDryRunResultPresentationConsent: productionAuthorizationReadSources(root, []string{
			"internal/runtime/owner/route_enablement_lookup_route_dispatch_dry_run_result_presentation_consent_audit_preview.go",
			"internal/runtime/owner/route_enablement_lookup_route_dispatch_dry_run_result_presentation_consent_audit_preview_test.go",
		}),
	}
}

func routeEnablementLookupRouteDispatchDryRunResultNotificationPresentationAuditMainlineReady(source string) bool {
	activeTaskReady := productionAuthorizationHasAll(source, []string{
		"storage-root authorization receipt dry-run lookup result consumer projection route enablement lookup route dispatch dry-run result notification presentation audit preview",
		"storage-root authorization receipt dry-run lookup result consumer projection route enablement lookup route dispatch dry-run result presentation consent audit preview",
		"accepting or consuming receipts",
	})
	continuityReady := productionAuthorizationHasAll(source, []string{
		"route enablement lookup route dispatch dry-run result notification presentation audit preview",
		"route enablement lookup route dispatch dry-run result notification action request-object dispatch authorization audit preview",
		"notification presentation audit preview",
	})
	return activeTaskReady || continuityReady
}

func routeEnablementLookupRouteDispatchDryRunResultNotificationPresentationAuditConsentReady(source string) bool {
	return productionAuthorizationHasAll(source, []string{
		"xnix.runtime.production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_kde_safe_redacted_status_closed_record_writer_storage_root_authorization_receipt_dry_run_lookup_result_consumer_projection_route_enablement_lookup_route_dispatch_dry_run_result_presentation_consent_audit.v1",
		"production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-presentation-consent-audit-ready-consent-review-only-notifications-disabled-action-cards-disabled",
		"LookupRouteDispatchDryRunResultPresentationConsentReady",
		"RouteEnablementLookupRouteDispatchDryRunResultPresentationConsentReady",
	})
}

func routeEnablementLookupRouteDispatchDryRunResultNotificationPresentationAuditItems(mainlineReady, consentReady bool) []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationPresentationAuditItem {
	kinds := []struct {
		id   string
		kind string
	}{
		{id: "install-failure", kind: "notification-center-install-failure"},
		{id: "repair-suggestion", kind: "notification-center-repair-suggestion"},
		{id: "environment-switch", kind: "notification-center-environment-switch"},
		{id: "approval-info", kind: "notification-center-approval-info"},
	}
	items := make([]ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationPresentationAuditItem, 0, len(kinds))
	for _, notification := range kinds {
		ready := mainlineReady && consentReady
		status := "missing-redacted-status-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-presentation-evidence"
		if ready {
			status = "redacted-status-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-presentation-ready-review-only-delivery-disabled"
		}
		items = append(items, ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationPresentationAuditItem{
			ID:                      "route-enablement-lookup-route-dispatch-dry-run-result-notification-presentation-" + notification.id,
			NotificationKind:        notification.id,
			PresentationKind:        notification.kind,
			EvidencePresent:         ready,
			CurrentMainlineConsumed: mainlineReady,
			RouteEnablementLookupRouteDispatchDryRunResultPresentationConsentConsumed: consentReady,
			RouteEnablementLookupRouteDispatchDryRunResultPresentationConsentReady:    consentReady,
			NotificationPresentationRequired:                                          true,
			NotificationPresentationModeled:                                           ready,
			NotificationPresentationReady:                                             ready,
			ExplicitUserConsentRequired:                                               true,
			ExplicitUserConsentCollected:                                              false,
			NotificationReviewOnly:                                                    true,
			KDESafeRedactedResultOnly:                                                 consentReady,
			RawResultHidden:                                                           consentReady,
			SideEffectsDisabled:                                                       true,
			RuntimeOwned:                                                              true,
			GoRuntimeBacked:                                                           true,
			NotificationPresentationStatus:                                            status,
			NextRequirement:                                                           "Require a separate notification delivery authorization audit before this modeled event can be delivered to KDE Notification Center.",
		})
	}
	return items
}

func routeEnablementLookupRouteDispatchDryRunResultNotificationPresentationAuditReadyCount(items []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationPresentationAuditItem) int {
	count := 0
	for _, item := range items {
		if item.EvidencePresent && item.NotificationPresentationRequired && item.NotificationPresentationModeled && item.NotificationPresentationReady && item.ExplicitUserConsentRequired && !item.ExplicitUserConsentCollected && item.NotificationReviewOnly && item.KDESafeRedactedResultOnly && item.RawResultHidden && item.SideEffectsDisabled && !item.NotificationDeliveryEnabled && !item.NotificationSent && !item.NotificationActionEnabled && !item.ActionCardEnabled && !item.HostRootModified {
			count++
		}
	}
	return count
}

func routeEnablementLookupRouteDispatchDryRunResultNotificationPresentationAuditItemIDs(items []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationPresentationAuditItem) []string {
	ids := make([]string, 0, len(items))
	for _, item := range items {
		ids = append(ids, item.ID)
	}
	return ids
}

func routeEnablementLookupRouteDispatchDryRunResultNotificationPresentationAuditChecks(preview ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationPresentationAuditPreview) []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationPresentationAuditCheck {
	return []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationPresentationAuditCheck{
		routeEnablementLookupRouteDispatchDryRunResultNotificationPresentationAuditCheck("current-mainline-consumed", preview.CurrentMainlineConsumed, "Current mainline names the notification presentation continuation."),
		routeEnablementLookupRouteDispatchDryRunResultNotificationPresentationAuditCheck("route-enablement-lookup-route-dispatch-dry-run-result-presentation-consent-consumed", preview.RouteEnablementLookupRouteDispatchDryRunResultPresentationConsentConsumed && preview.RouteEnablementLookupRouteDispatchDryRunResultPresentationConsentReady, "Presentation consent audit is consumed as predecessor evidence."),
		routeEnablementLookupRouteDispatchDryRunResultNotificationPresentationAuditCheck("lookup-route-dispatch-dry-run-result-notification-presentation-modeled", preview.LookupRouteDispatchDryRunResultNotificationPresentationModeled && preview.LookupRouteDispatchDryRunResultNotificationPresentationReady, "Notification presentation boundary is modeled."),
		routeEnablementLookupRouteDispatchDryRunResultNotificationPresentationAuditCheck("four-kde-notification-events-review-only", preview.NotificationPresentationItemCount == 4 && preview.ReadyNotificationPresentationItemCount == 4 && preview.DeliveredNotificationItemCount == 0, "Four KDE Notification Center event classes are modeled and remain review-only."),
		routeEnablementLookupRouteDispatchDryRunResultNotificationPresentationAuditCheck("notification-delivery-actions-and-action-cards-disabled", preview.NotificationDeliveryDisabled && preview.NotificationActionDisabled && preview.ActionCardsRemainDisabled && !preview.NotificationDeliveryEnabled && !preview.NotificationSent && !preview.NotificationCenterEventTriggered && !preview.NotificationActionEnabled && !preview.ActionCardEnabled, "Notification delivery, notification actions, and action cards remain disabled."),
		routeEnablementLookupRouteDispatchDryRunResultNotificationPresentationAuditCheck("raw-result-exposure-and-persistence-disabled", preview.RawResultHidden && !preview.RawResultExposed && !preview.DryRunResultPersisted && !preview.RedactedSummaryPersisted && !preview.StorageWriteEnabled, "Raw result exposure and persistence remain disabled."),
		routeEnablementLookupRouteDispatchDryRunResultNotificationPresentationAuditCheck("production-and-host-boundary-closed", !preview.ProductionOwnershipReady && !preview.ProductionBusClaimed && !preview.WriteMethodsEnabled && !preview.RuntimeWritesEnabled && !preview.BackendLaunchEnabled && !preview.HostRootModified, "Production and host boundaries remain closed."),
	}
}

func routeEnablementLookupRouteDispatchDryRunResultNotificationPresentationAuditCheck(id string, passed bool, summary string) ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationPresentationAuditCheck {
	status := "blocked"
	if passed {
		status = "pass"
	}
	return ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationPresentationAuditCheck{ID: id, Status: status, Summary: summary}
}

func routeEnablementLookupRouteDispatchDryRunResultNotificationPresentationAuditCheckIDs(checks []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationPresentationAuditCheck) []string {
	ids := make([]string, 0, len(checks))
	for _, check := range checks {
		ids = append(ids, check.ID)
	}
	return ids
}

func routeEnablementLookupRouteDispatchDryRunResultNotificationPresentationAuditCountsFor(checks []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationPresentationAuditCheck) ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationPresentationAuditCheckCounts {
	counts := ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationPresentationAuditCheckCounts{Total: len(checks)}
	for _, check := range checks {
		if check.Status == "pass" {
			counts.Passed++
		} else {
			counts.Blocked++
		}
	}
	return counts
}
