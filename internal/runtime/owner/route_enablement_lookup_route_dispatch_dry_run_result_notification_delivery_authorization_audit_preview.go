package owner

type ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationDeliveryAuthorizationAuditPreview struct {
	Version                                                                        string                                                                                                                                                                                                                                                                           `json:"version"`
	SchemaVersion                                                                  string                                                                                                                                                                                                                                                                           `json:"schema_version"`
	RequestType                                                                    string                                                                                                                                                                                                                                                                           `json:"request_type"`
	AuditType                                                                      string                                                                                                                                                                                                                                                                           `json:"audit_type"`
	Source                                                                         string                                                                                                                                                                                                                                                                           `json:"source"`
	AuditDecision                                                                  string                                                                                                                                                                                                                                                                           `json:"audit_decision"`
	CurrentMainlineConsumed                                                        bool                                                                                                                                                                                                                                                                             `json:"current_mainline_consumed"`
	RouteEnablementLookupRouteDispatchDryRunResultNotificationPresentationConsumed bool                                                                                                                                                                                                                                                                             `json:"route_enablement_lookup_route_dispatch_dry_run_result_notification_presentation_consumed"`
	RouteEnablementLookupRouteDispatchDryRunResultNotificationPresentationReady    bool                                                                                                                                                                                                                                                                             `json:"route_enablement_lookup_route_dispatch_dry_run_result_notification_presentation_ready"`
	LookupRouteDispatchDryRunResultNotificationDeliveryAuthorizationRequired       bool                                                                                                                                                                                                                                                                             `json:"lookup_route_dispatch_dry_run_result_notification_delivery_authorization_required"`
	LookupRouteDispatchDryRunResultNotificationDeliveryAuthorizationModeled        bool                                                                                                                                                                                                                                                                             `json:"lookup_route_dispatch_dry_run_result_notification_delivery_authorization_modeled"`
	LookupRouteDispatchDryRunResultNotificationDeliveryAuthorizationReady          bool                                                                                                                                                                                                                                                                             `json:"lookup_route_dispatch_dry_run_result_notification_delivery_authorization_ready"`
	RouteEnablementLookupRouteDispatchDryRunResultNotificationDeliveryReady        bool                                                                                                                                                                                                                                                                             `json:"route_enablement_lookup_route_dispatch_dry_run_result_notification_delivery_ready"`
	NotificationPresentationReviewOnlyConsumed                                     bool                                                                                                                                                                                                                                                                             `json:"notification_presentation_review_only_consumed"`
	KDENotificationDeliveryAuthorizationModeled                                    bool                                                                                                                                                                                                                                                                             `json:"kde_notification_delivery_authorization_modeled"`
	NotificationCenterDeliveryAuthorizationModeled                                 bool                                                                                                                                                                                                                                                                             `json:"notification_center_delivery_authorization_modeled"`
	InstallFailureDeliveryAuthorizationModeled                                     bool                                                                                                                                                                                                                                                                             `json:"install_failure_delivery_authorization_modeled"`
	RepairSuggestionDeliveryAuthorizationModeled                                   bool                                                                                                                                                                                                                                                                             `json:"repair_suggestion_delivery_authorization_modeled"`
	EnvironmentSwitchDeliveryAuthorizationModeled                                  bool                                                                                                                                                                                                                                                                             `json:"environment_switch_delivery_authorization_modeled"`
	ApprovalInfoDeliveryAuthorizationModeled                                       bool                                                                                                                                                                                                                                                                             `json:"approval_info_delivery_authorization_modeled"`
	NotificationDeliveryAuthorizationReviewOnly                                    bool                                                                                                                                                                                                                                                                             `json:"notification_delivery_authorization_review_only"`
	NotificationDeliveryAuthorizationGranted                                       bool                                                                                                                                                                                                                                                                             `json:"notification_delivery_authorization_granted"`
	NotificationDeliveryDisabled                                                   bool                                                                                                                                                                                                                                                                             `json:"notification_delivery_disabled"`
	NotificationActionDisabled                                                     bool                                                                                                                                                                                                                                                                             `json:"notification_action_disabled"`
	ActionCardsRemainDisabled                                                      bool                                                                                                                                                                                                                                                                             `json:"action_cards_remain_disabled"`
	ExplicitUserConsentRequired                                                    bool                                                                                                                                                                                                                                                                             `json:"explicit_user_consent_required"`
	ExplicitUserConsentCollected                                                   bool                                                                                                                                                                                                                                                                             `json:"explicit_user_consent_collected"`
	KDESafeRedactedResultOnly                                                      bool                                                                                                                                                                                                                                                                             `json:"kde_safe_redacted_result_only"`
	RawResultHidden                                                                bool                                                                                                                                                                                                                                                                             `json:"raw_result_hidden"`
	NotificationDeliveryAuthorizationItemCount                                     int                                                                                                                                                                                                                                                                              `json:"notification_delivery_authorization_item_count"`
	RequiredNotificationDeliveryAuthorizationItemCount                             int                                                                                                                                                                                                                                                                              `json:"required_notification_delivery_authorization_item_count"`
	ReadyNotificationDeliveryAuthorizationItemCount                                int                                                                                                                                                                                                                                                                              `json:"ready_notification_delivery_authorization_item_count"`
	MissingNotificationDeliveryAuthorizationItemCount                              int                                                                                                                                                                                                                                                                              `json:"missing_notification_delivery_authorization_item_count"`
	AuthorizedNotificationDeliveryItemCount                                        int                                                                                                                                                                                                                                                                              `json:"authorized_notification_delivery_item_count"`
	DeliveredNotificationItemCount                                                 int                                                                                                                                                                                                                                                                              `json:"delivered_notification_item_count"`
	NotificationActionEnabledItemCount                                             int                                                                                                                                                                                                                                                                              `json:"notification_action_enabled_item_count"`
	ActionCardEnabledItemCount                                                     int                                                                                                                                                                                                                                                                              `json:"action_card_enabled_item_count"`
	SideEffectNotificationDeliveryAuthorizationItemCount                           int                                                                                                                                                                                                                                                                              `json:"side_effect_notification_delivery_authorization_item_count"`
	NotificationDeliveryAuthorizationItems                                         []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationDeliveryAuthorizationAuditItem      `json:"notification_delivery_authorization_items"`
	NotificationDeliveryAuthorizationItemIDs                                       []string                                                                                                                                                                                                                                                                         `json:"notification_delivery_authorization_item_ids"`
	Checks                                                                         []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationDeliveryAuthorizationAuditCheck     `json:"checks"`
	CheckIDs                                                                       []string                                                                                                                                                                                                                                                                         `json:"check_ids"`
	Counts                                                                         ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationDeliveryAuthorizationAuditCheckCounts `json:"counts"`
	ReceiptPresent                                                                 bool                                                                                                                                                                                                                                                                             `json:"receipt_present"`
	ReceiptAccepted                                                                bool                                                                                                                                                                                                                                                                             `json:"receipt_accepted"`
	ReceiptConsumed                                                                bool                                                                                                                                                                                                                                                                             `json:"receipt_consumed"`
	RouteEnablementAccepted                                                        bool                                                                                                                                                                                                                                                                             `json:"route_enablement_accepted"`
	LookupRouteEnabled                                                             bool                                                                                                                                                                                                                                                                             `json:"lookup_route_enabled"`
	LookupRouteDispatchAuthorized                                                  bool                                                                                                                                                                                                                                                                             `json:"lookup_route_dispatch_authorized"`
	LookupRouteDispatchCallable                                                    bool                                                                                                                                                                                                                                                                             `json:"lookup_route_dispatch_callable"`
	StorageWriteEnabled                                                            bool                                                                                                                                                                                                                                                                             `json:"storage_write_enabled"`
	StatusPersistenceWriteEnabled                                                  bool                                                                                                                                                                                                                                                                             `json:"status_persistence_write_enabled"`
	RedactedSummaryPersisted                                                       bool                                                                                                                                                                                                                                                                             `json:"redacted_summary_persisted"`
	KDEStatusPersisted                                                             bool                                                                                                                                                                                                                                                                             `json:"kde_status_persisted"`
	RuntimeDiagnosticsPersisted                                                    bool                                                                                                                                                                                                                                                                             `json:"runtime_diagnostics_persisted"`
	DryRunResultPersisted                                                          bool                                                                                                                                                                                                                                                                             `json:"dry_run_result_persisted"`
	RawResultExposed                                                               bool                                                                                                                                                                                                                                                                             `json:"raw_result_exposed"`
	DispatchDryRunExecuted                                                         bool                                                                                                                                                                                                                                                                             `json:"dispatch_dry_run_executed"`
	RequestObjectCreationEnabled                                                   bool                                                                                                                                                                                                                                                                             `json:"request_object_creation_enabled"`
	RequestObjectDispatchEnabled                                                   bool                                                                                                                                                                                                                                                                             `json:"request_object_dispatch_enabled"`
	PortalRequestCreated                                                           bool                                                                                                                                                                                                                                                                             `json:"portal_request_created"`
	NotificationDeliveryEnabled                                                    bool                                                                                                                                                                                                                                                                             `json:"notification_delivery_enabled"`
	NotificationSent                                                               bool                                                                                                                                                                                                                                                                             `json:"notification_sent"`
	NotificationCenterEventTriggered                                               bool                                                                                                                                                                                                                                                                             `json:"notification_center_event_triggered"`
	NotificationActionEnabled                                                      bool                                                                                                                                                                                                                                                                             `json:"notification_action_enabled"`
	ActionCardEnabled                                                              bool                                                                                                                                                                                                                                                                             `json:"action_card_enabled"`
	CompatibilityCenterOpened                                                      bool                                                                                                                                                                                                                                                                             `json:"compatibility_center_opened"`
	SupportBundleExported                                                          bool                                                                                                                                                                                                                                                                             `json:"support_bundle_exported"`
	SupportCaseCreated                                                             bool                                                                                                                                                                                                                                                                             `json:"support_case_created"`
	RuntimeOwned                                                                   bool                                                                                                                                                                                                                                                                             `json:"runtime_owned"`
	GoRuntimeBacked                                                                bool                                                                                                                                                                                                                                                                             `json:"go_runtime_backed"`
	KDEPolicyOwner                                                                 bool                                                                                                                                                                                                                                                                             `json:"kde_policy_owner"`
	ProductionReadiness                                                            bool                                                                                                                                                                                                                                                                             `json:"production_readiness"`
	ProductionOwnershipReady                                                       bool                                                                                                                                                                                                                                                                             `json:"production_ownership_ready"`
	SystemServiceStarted                                                           bool                                                                                                                                                                                                                                                                             `json:"system_service_started"`
	SessionBusClaimed                                                              bool                                                                                                                                                                                                                                                                             `json:"session_bus_claimed"`
	ProductionBusClaimed                                                           bool                                                                                                                                                                                                                                                                             `json:"production_bus_claimed"`
	WriteMethodsEnabled                                                            bool                                                                                                                                                                                                                                                                             `json:"write_methods_enabled"`
	RuntimeWritesEnabled                                                           bool                                                                                                                                                                                                                                                                             `json:"runtime_writes_enabled"`
	DesktopFilesWritten                                                            bool                                                                                                                                                                                                                                                                             `json:"desktop_files_written"`
	KDEConfigurationWritten                                                        bool                                                                                                                                                                                                                                                                             `json:"kde_configuration_written"`
	PortalCallExecuted                                                             bool                                                                                                                                                                                                                                                                             `json:"portal_call_executed"`
	AdapterInvocationEnabled                                                       bool                                                                                                                                                                                                                                                                             `json:"adapter_invocation_enabled"`
	BackendLaunchEnabled                                                           bool                                                                                                                                                                                                                                                                             `json:"backend_launch_enabled"`
	BackendProcessStarted                                                          bool                                                                                                                                                                                                                                                                             `json:"backend_process_started"`
	NetworkRequired                                                                bool                                                                                                                                                                                                                                                                             `json:"network_required"`
	HostRootModified                                                               bool                                                                                                                                                                                                                                                                             `json:"host_root_modified"`
	PrivilegedContainerRequired                                                    bool                                                                                                                                                                                                                                                                             `json:"privileged_container_required"`
	StateRootPathExposed                                                           bool                                                                                                                                                                                                                                                                             `json:"state_root_path_exposed"`
	FilePathsExposed                                                               bool                                                                                                                                                                                                                                                                             `json:"file_paths_exposed"`
	FileContentRead                                                                bool                                                                                                                                                                                                                                                                             `json:"file_content_read"`
	RawCommandExposed                                                              bool                                                                                                                                                                                                                                                                             `json:"raw_command_exposed"`
	RawExecutableExposed                                                           bool                                                                                                                                                                                                                                                                             `json:"raw_executable_exposed"`
	BackendDetailsExposed                                                          bool                                                                                                                                                                                                                                                                             `json:"backend_details_exposed"`
	BlockedActions                                                                 []string                                                                                                                                                                                                                                                                         `json:"blocked_actions"`
	NextRequirements                                                               []string                                                                                                                                                                                                                                                                         `json:"next_requirements"`
	DesktopSafeSummary                                                             string                                                                                                                                                                                                                                                                           `json:"desktop_safe_summary"`
}

type ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationDeliveryAuthorizationAuditItem struct {
	ID                               string `json:"id"`
	NotificationKind                 string `json:"notification_kind"`
	DeliveryKind                     string `json:"delivery_kind"`
	EvidencePresent                  bool   `json:"evidence_present"`
	CurrentMainlineConsumed          bool   `json:"current_mainline_consumed"`
	NotificationPresentationConsumed bool   `json:"notification_presentation_consumed"`
	NotificationPresentationReady    bool   `json:"notification_presentation_ready"`
	DeliveryAuthorizationRequired    bool   `json:"delivery_authorization_required"`
	DeliveryAuthorizationModeled     bool   `json:"delivery_authorization_modeled"`
	DeliveryAuthorizationReady       bool   `json:"delivery_authorization_ready"`
	DeliveryAuthorizationGranted     bool   `json:"delivery_authorization_granted"`
	ExplicitUserConsentRequired      bool   `json:"explicit_user_consent_required"`
	ExplicitUserConsentCollected     bool   `json:"explicit_user_consent_collected"`
	DeliveryReviewOnly               bool   `json:"delivery_review_only"`
	KDESafeRedactedResultOnly        bool   `json:"kde_safe_redacted_result_only"`
	RawResultHidden                  bool   `json:"raw_result_hidden"`
	NotificationDeliveryEnabled      bool   `json:"notification_delivery_enabled"`
	NotificationSent                 bool   `json:"notification_sent"`
	NotificationCenterEventTriggered bool   `json:"notification_center_event_triggered"`
	NotificationActionEnabled        bool   `json:"notification_action_enabled"`
	ActionCardEnabled                bool   `json:"action_card_enabled"`
	SideEffectsDisabled              bool   `json:"side_effects_disabled"`
	RuntimeOwned                     bool   `json:"runtime_owned"`
	GoRuntimeBacked                  bool   `json:"go_runtime_backed"`
	KDEPolicyOwner                   bool   `json:"kde_policy_owner"`
	HostRootModified                 bool   `json:"host_root_modified"`
	InternalDetailsExposed           bool   `json:"internal_details_exposed"`
	DeliveryAuthorizationStatus      string `json:"delivery_authorization_status"`
	NextRequirement                  string `json:"next_requirement"`
}

type ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationDeliveryAuthorizationAuditCheck struct {
	ID      string `json:"id"`
	Status  string `json:"status"`
	Summary string `json:"summary"`
}

type ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationDeliveryAuthorizationAuditCheckCounts struct {
	Total   int `json:"total"`
	Passed  int `json:"passed"`
	Pending int `json:"pending"`
	Blocked int `json:"blocked"`
}

type routeEnablementLookupRouteDispatchDryRunResultNotificationDeliveryAuthorizationAuditSourceSet struct {
	CurrentMainline                                                        string
	RouteEnablementLookupRouteDispatchDryRunResultNotificationPresentation string
}

func NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationDeliveryAuthorizationAuditPreview(root string) (ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationDeliveryAuthorizationAuditPreview, error) {
	version, err := readProductionDBusGateReviewVersion(root)
	if err != nil {
		return ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationDeliveryAuthorizationAuditPreview{}, err
	}
	sources := routeEnablementLookupRouteDispatchDryRunResultNotificationDeliveryAuthorizationAuditSources(root)
	mainlineReady := routeEnablementLookupRouteDispatchDryRunResultNotificationDeliveryAuthorizationAuditMainlineReady(sources.CurrentMainline)
	presentationReady := routeEnablementLookupRouteDispatchDryRunResultNotificationDeliveryAuthorizationAuditPresentationReady(sources.RouteEnablementLookupRouteDispatchDryRunResultNotificationPresentation)
	items := routeEnablementLookupRouteDispatchDryRunResultNotificationDeliveryAuthorizationAuditItems(mainlineReady, presentationReady)
	readyItemCount := routeEnablementLookupRouteDispatchDryRunResultNotificationDeliveryAuthorizationAuditReadyCount(items)
	ready := mainlineReady && presentationReady && readyItemCount == len(items)
	preview := ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationDeliveryAuthorizationAuditPreview{
		Version:                 version,
		SchemaVersion:           "xnix.runtime.production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_kde_safe_redacted_status_closed_record_writer_storage_root_authorization_receipt_dry_run_lookup_result_consumer_projection_route_enablement_lookup_route_dispatch_dry_run_result_notification_delivery_authorization_audit.v1",
		RequestType:             "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-delivery-authorization-audit-preview",
		AuditType:               "receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-delivery-authorization-audit",
		Source:                  "docs/xnix-current-mainline.md+route-enable-lookup-route-dispatch-dry-run-result-notification-presentation-audit",
		AuditDecision:           "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-delivery-authorization-audit-blocked",
		CurrentMainlineConsumed: mainlineReady,
		RouteEnablementLookupRouteDispatchDryRunResultNotificationPresentationConsumed: presentationReady,
		RouteEnablementLookupRouteDispatchDryRunResultNotificationPresentationReady:    presentationReady,
		LookupRouteDispatchDryRunResultNotificationDeliveryAuthorizationRequired:       true,
		LookupRouteDispatchDryRunResultNotificationDeliveryAuthorizationModeled:        true,
		LookupRouteDispatchDryRunResultNotificationDeliveryAuthorizationReady:          ready,
		RouteEnablementLookupRouteDispatchDryRunResultNotificationDeliveryReady:        ready,
		NotificationPresentationReviewOnlyConsumed:                                     presentationReady,
		KDENotificationDeliveryAuthorizationModeled:                                    ready,
		NotificationCenterDeliveryAuthorizationModeled:                                 ready,
		InstallFailureDeliveryAuthorizationModeled:                                     ready,
		RepairSuggestionDeliveryAuthorizationModeled:                                   ready,
		EnvironmentSwitchDeliveryAuthorizationModeled:                                  ready,
		ApprovalInfoDeliveryAuthorizationModeled:                                       ready,
		NotificationDeliveryAuthorizationReviewOnly:                                    true,
		NotificationDeliveryAuthorizationGranted:                                       false,
		NotificationDeliveryDisabled:                                                   true,
		NotificationActionDisabled:                                                     true,
		ActionCardsRemainDisabled:                                                      true,
		ExplicitUserConsentRequired:                                                    true,
		ExplicitUserConsentCollected:                                                   false,
		KDESafeRedactedResultOnly:                                                      presentationReady,
		RawResultHidden:                                                                presentationReady,
		NotificationDeliveryAuthorizationItemCount:                                     len(items),
		RequiredNotificationDeliveryAuthorizationItemCount:                             len(items),
		ReadyNotificationDeliveryAuthorizationItemCount:                                readyItemCount,
		MissingNotificationDeliveryAuthorizationItemCount:                              len(items) - readyItemCount,
		NotificationDeliveryAuthorizationItems:                                         items,
		NotificationDeliveryAuthorizationItemIDs:                                       routeEnablementLookupRouteDispatchDryRunResultNotificationDeliveryAuthorizationAuditItemIDs(items),
		RuntimeOwned:                                                                   true,
		GoRuntimeBacked:                                                                true,
		KDEPolicyOwner:                                                                 false,
		BlockedActions: []string{
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
		NextRequirements: []string{
			"Add a separate notification delivery grant audit before KDE Notification Center entries can be sent.",
			"Keep delivery authorization review-only until explicit user consent collection and production ownership are separately authorized.",
		},
		DesktopSafeSummary: "KDE Notification Center delivery authorization for redacted lookup route dispatch dry-run result presentation is modeled for install failure, repair suggestion, environment switch, and approval information events while delivery grants, notification sending, notification actions, action cards, raw results, receipts, lookup routes, dispatch, persistence, production, backend, and host side effects remain disabled.",
	}
	if ready {
		preview.AuditDecision = "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-delivery-authorization-audit-ready-delivery-grant-disabled"
	}
	preview.Checks = routeEnablementLookupRouteDispatchDryRunResultNotificationDeliveryAuthorizationAuditChecks(preview)
	preview.CheckIDs = routeEnablementLookupRouteDispatchDryRunResultNotificationDeliveryAuthorizationAuditCheckIDs(preview.Checks)
	preview.Counts = routeEnablementLookupRouteDispatchDryRunResultNotificationDeliveryAuthorizationAuditCountsFor(preview.Checks)
	if err := validateNoBackendTerms(preview, "KDE-safe redacted status dry-run lookup result consumer projection route enablement lookup route dispatch dry-run result notification delivery authorization audit preview"); err != nil {
		return ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationDeliveryAuthorizationAuditPreview{}, err
	}
	return preview, nil
}

func routeEnablementLookupRouteDispatchDryRunResultNotificationDeliveryAuthorizationAuditSources(root string) routeEnablementLookupRouteDispatchDryRunResultNotificationDeliveryAuthorizationAuditSourceSet {
	return routeEnablementLookupRouteDispatchDryRunResultNotificationDeliveryAuthorizationAuditSourceSet{
		CurrentMainline: productionAuthorizationReadSources(root, []string{"docs/xnix-current-mainline.md"}),
		RouteEnablementLookupRouteDispatchDryRunResultNotificationPresentation: productionAuthorizationReadSources(root, []string{
			"internal/runtime/owner/route_enablement_lookup_route_dispatch_dry_run_result_notification_presentation_audit_preview.go",
			"internal/runtime/owner/route_enablement_lookup_route_dispatch_dry_run_result_notification_presentation_audit_preview_test.go",
		}),
	}
}

func routeEnablementLookupRouteDispatchDryRunResultNotificationDeliveryAuthorizationAuditMainlineReady(source string) bool {
	return productionAuthorizationHasAll(source, []string{
		"route enablement lookup route dispatch dry-run result notification delivery authorization audit preview",
		"storage-root authorization receipt dry-run lookup result consumer projection route enablement lookup route dispatch dry-run result notification presentation audit preview",
		"without sending notifications",
	})
}

func routeEnablementLookupRouteDispatchDryRunResultNotificationDeliveryAuthorizationAuditPresentationReady(source string) bool {
	return productionAuthorizationHasAll(source, []string{
		"xnix.runtime.production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_kde_safe_redacted_status_closed_record_writer_storage_root_authorization_receipt_dry_run_lookup_result_consumer_projection_route_enablement_lookup_route_dispatch_dry_run_result_notification_presentation_audit.v1",
		"production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-presentation-audit-ready-notification-review-only-delivery-disabled",
		"LookupRouteDispatchDryRunResultNotificationPresentationReady",
		"RouteEnablementLookupRouteDispatchDryRunResultNotificationPresentationReady",
	})
}

func routeEnablementLookupRouteDispatchDryRunResultNotificationDeliveryAuthorizationAuditItems(mainlineReady, presentationReady bool) []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationDeliveryAuthorizationAuditItem {
	kinds := []struct {
		id   string
		kind string
	}{
		{id: "install-failure", kind: "notification-center-install-failure-delivery-authorization"},
		{id: "repair-suggestion", kind: "notification-center-repair-suggestion-delivery-authorization"},
		{id: "environment-switch", kind: "notification-center-environment-switch-delivery-authorization"},
		{id: "approval-info", kind: "notification-center-approval-info-delivery-authorization"},
	}
	items := make([]ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationDeliveryAuthorizationAuditItem, 0, len(kinds))
	for _, notification := range kinds {
		ready := mainlineReady && presentationReady
		status := "missing-redacted-status-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-delivery-authorization-evidence"
		if ready {
			status = "redacted-status-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-delivery-authorization-ready-delivery-grant-disabled"
		}
		items = append(items, ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationDeliveryAuthorizationAuditItem{
			ID:                               "route-enablement-lookup-route-dispatch-dry-run-result-notification-delivery-authorization-" + notification.id,
			NotificationKind:                 notification.id,
			DeliveryKind:                     notification.kind,
			EvidencePresent:                  ready,
			CurrentMainlineConsumed:          mainlineReady,
			NotificationPresentationConsumed: presentationReady,
			NotificationPresentationReady:    presentationReady,
			DeliveryAuthorizationRequired:    true,
			DeliveryAuthorizationModeled:     ready,
			DeliveryAuthorizationReady:       ready,
			DeliveryAuthorizationGranted:     false,
			ExplicitUserConsentRequired:      true,
			ExplicitUserConsentCollected:     false,
			DeliveryReviewOnly:               true,
			KDESafeRedactedResultOnly:        presentationReady,
			RawResultHidden:                  presentationReady,
			SideEffectsDisabled:              true,
			RuntimeOwned:                     true,
			GoRuntimeBacked:                  true,
			DeliveryAuthorizationStatus:      status,
			NextRequirement:                  "Require a separate notification delivery grant audit before this authorized shape can send a KDE Notification Center event.",
		})
	}
	return items
}

func routeEnablementLookupRouteDispatchDryRunResultNotificationDeliveryAuthorizationAuditReadyCount(items []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationDeliveryAuthorizationAuditItem) int {
	count := 0
	for _, item := range items {
		if item.EvidencePresent && item.DeliveryAuthorizationRequired && item.DeliveryAuthorizationModeled && item.DeliveryAuthorizationReady && !item.DeliveryAuthorizationGranted && item.ExplicitUserConsentRequired && !item.ExplicitUserConsentCollected && item.DeliveryReviewOnly && item.KDESafeRedactedResultOnly && item.RawResultHidden && item.SideEffectsDisabled && !item.NotificationDeliveryEnabled && !item.NotificationSent && !item.NotificationCenterEventTriggered && !item.NotificationActionEnabled && !item.ActionCardEnabled && !item.HostRootModified {
			count++
		}
	}
	return count
}

func routeEnablementLookupRouteDispatchDryRunResultNotificationDeliveryAuthorizationAuditItemIDs(items []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationDeliveryAuthorizationAuditItem) []string {
	ids := make([]string, 0, len(items))
	for _, item := range items {
		ids = append(ids, item.ID)
	}
	return ids
}

func routeEnablementLookupRouteDispatchDryRunResultNotificationDeliveryAuthorizationAuditChecks(preview ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationDeliveryAuthorizationAuditPreview) []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationDeliveryAuthorizationAuditCheck {
	return []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationDeliveryAuthorizationAuditCheck{
		routeEnablementLookupRouteDispatchDryRunResultNotificationDeliveryAuthorizationAuditCheck("current-mainline-consumed", preview.CurrentMainlineConsumed, "Current mainline names the notification delivery authorization continuation."),
		routeEnablementLookupRouteDispatchDryRunResultNotificationDeliveryAuthorizationAuditCheck("route-enablement-lookup-route-dispatch-dry-run-result-notification-presentation-consumed", preview.RouteEnablementLookupRouteDispatchDryRunResultNotificationPresentationConsumed && preview.RouteEnablementLookupRouteDispatchDryRunResultNotificationPresentationReady, "Notification presentation audit is consumed as predecessor evidence."),
		routeEnablementLookupRouteDispatchDryRunResultNotificationDeliveryAuthorizationAuditCheck("lookup-route-dispatch-dry-run-result-notification-delivery-authorization-modeled", preview.LookupRouteDispatchDryRunResultNotificationDeliveryAuthorizationModeled && preview.LookupRouteDispatchDryRunResultNotificationDeliveryAuthorizationReady, "Notification delivery authorization boundary is modeled."),
		routeEnablementLookupRouteDispatchDryRunResultNotificationDeliveryAuthorizationAuditCheck("four-kde-notification-delivery-events-review-only", preview.NotificationDeliveryAuthorizationItemCount == 4 && preview.ReadyNotificationDeliveryAuthorizationItemCount == 4 && preview.AuthorizedNotificationDeliveryItemCount == 0 && preview.DeliveredNotificationItemCount == 0, "Four KDE Notification Center delivery event classes are authorization-modeled and remain review-only."),
		routeEnablementLookupRouteDispatchDryRunResultNotificationDeliveryAuthorizationAuditCheck("delivery-grants-sending-and-events-disabled", preview.NotificationDeliveryDisabled && !preview.NotificationDeliveryAuthorizationGranted && !preview.NotificationDeliveryEnabled && !preview.NotificationSent && !preview.NotificationCenterEventTriggered, "Delivery grants, notification sending, and Notification Center events remain disabled."),
		routeEnablementLookupRouteDispatchDryRunResultNotificationDeliveryAuthorizationAuditCheck("notification-actions-action-cards-and-raw-results-disabled", preview.NotificationActionDisabled && preview.ActionCardsRemainDisabled && !preview.NotificationActionEnabled && !preview.ActionCardEnabled && preview.RawResultHidden && !preview.RawResultExposed && !preview.DryRunResultPersisted && !preview.StorageWriteEnabled, "Notification actions, action cards, raw result exposure, and persistence remain disabled."),
		routeEnablementLookupRouteDispatchDryRunResultNotificationDeliveryAuthorizationAuditCheck("production-and-host-boundary-closed", !preview.ProductionOwnershipReady && !preview.ProductionBusClaimed && !preview.WriteMethodsEnabled && !preview.RuntimeWritesEnabled && !preview.BackendLaunchEnabled && !preview.HostRootModified, "Production and host boundaries remain closed."),
	}
}

func routeEnablementLookupRouteDispatchDryRunResultNotificationDeliveryAuthorizationAuditCheck(id string, passed bool, summary string) ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationDeliveryAuthorizationAuditCheck {
	status := "blocked"
	if passed {
		status = "pass"
	}
	return ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationDeliveryAuthorizationAuditCheck{ID: id, Status: status, Summary: summary}
}

func routeEnablementLookupRouteDispatchDryRunResultNotificationDeliveryAuthorizationAuditCheckIDs(checks []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationDeliveryAuthorizationAuditCheck) []string {
	ids := make([]string, 0, len(checks))
	for _, check := range checks {
		ids = append(ids, check.ID)
	}
	return ids
}

func routeEnablementLookupRouteDispatchDryRunResultNotificationDeliveryAuthorizationAuditCountsFor(checks []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationDeliveryAuthorizationAuditCheck) ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationDeliveryAuthorizationAuditCheckCounts {
	counts := ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationDeliveryAuthorizationAuditCheckCounts{Total: len(checks)}
	for _, check := range checks {
		if check.Status == "pass" {
			counts.Passed++
		} else {
			counts.Blocked++
		}
	}
	return counts
}
