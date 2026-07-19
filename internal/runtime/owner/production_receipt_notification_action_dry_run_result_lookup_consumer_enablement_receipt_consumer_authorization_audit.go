package owner

type ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerAuthorizationAuditPreview struct {
	Version                               string                                                                                                     `json:"version"`
	SchemaVersion                         string                                                                                                     `json:"schema_version"`
	RequestType                           string                                                                                                     `json:"request_type"`
	AuditType                             string                                                                                                     `json:"audit_type"`
	Source                                string                                                                                                     `json:"source"`
	AuditDecision                         string                                                                                                     `json:"audit_decision"`
	ReceiptSchema                         string                                                                                                     `json:"receipt_schema"`
	OpaqueReceiptID                       string                                                                                                     `json:"opaque_receipt_id"`
	ConsumerAuthorizationAuditRequired    bool                                                                                                       `json:"consumer_authorization_audit_required"`
	ConsumerAuthorizationAuditModeled     bool                                                                                                       `json:"consumer_authorization_audit_modeled"`
	ConsumptionGateAuditConsumed          bool                                                                                                       `json:"consumption_gate_audit_consumed"`
	ConsumerRedactionAuditConsumed        bool                                                                                                       `json:"consumer_redaction_audit_consumed"`
	LookupRouteAuthorizationAuditConsumed bool                                                                                                       `json:"lookup_route_authorization_audit_consumed"`
	ConsumerAuthorizationGuidanceConsumed bool                                                                                                       `json:"consumer_authorization_guidance_consumed"`
	ConsumerAuthorizationBoundaryReady    bool                                                                                                       `json:"consumer_authorization_boundary_ready"`
	ReceiptConsumptionGateReady           bool                                                                                                       `json:"receipt_consumption_gate_ready"`
	ConsumerRedactionBoundaryReady        bool                                                                                                       `json:"consumer_redaction_boundary_ready"`
	LookupRouteAuthorizationBoundaryReady bool                                                                                                       `json:"lookup_route_authorization_boundary_ready"`
	ConsumerAuthorizationReady            bool                                                                                                       `json:"consumer_authorization_ready"`
	ReceiptPresent                        bool                                                                                                       `json:"receipt_present"`
	ReceiptPersisted                      bool                                                                                                       `json:"receipt_persisted"`
	ReceiptAccepted                       bool                                                                                                       `json:"receipt_accepted"`
	ReceiptConsumed                       bool                                                                                                       `json:"receipt_consumed"`
	AuthorizationAccepted                 bool                                                                                                       `json:"authorization_accepted"`
	ConsumerAuthorizationGranted          bool                                                                                                       `json:"consumer_authorization_granted"`
	ConsumerConsumptionAuthorized         bool                                                                                                       `json:"consumer_consumption_authorized"`
	ConsumerEnablementAuthorized          bool                                                                                                       `json:"consumer_enablement_authorized"`
	KDEConsumerAuthorized                 bool                                                                                                       `json:"kde_consumer_authorized"`
	RuntimeConsumerAuthorized             bool                                                                                                       `json:"runtime_consumer_authorized"`
	KDEConsumerEnabled                    bool                                                                                                       `json:"kde_consumer_enabled"`
	RuntimeConsumerEnabled                bool                                                                                                       `json:"runtime_consumer_enabled"`
	LookupRouteAuthorized                 bool                                                                                                       `json:"lookup_route_authorized"`
	LookupRouteEnabled                    bool                                                                                                       `json:"lookup_route_enabled"`
	OpaqueLookupEnabled                   bool                                                                                                       `json:"opaque_lookup_enabled"`
	RedactedSummaryPersisted              bool                                                                                                       `json:"redacted_summary_persisted"`
	RawResultExposed                      bool                                                                                                       `json:"raw_result_exposed"`
	RuntimeDiagnosticsPersisted           bool                                                                                                       `json:"runtime_diagnostics_persisted"`
	DryRunResultPersisted                 bool                                                                                                       `json:"dry_run_result_persisted"`
	DispatchDryRunExecuted                bool                                                                                                       `json:"dispatch_dry_run_executed"`
	RequestObjectCreationEnabled          bool                                                                                                       `json:"request_object_creation_enabled"`
	RequestObjectDispatchEnabled          bool                                                                                                       `json:"request_object_dispatch_enabled"`
	PortalRequestCreated                  bool                                                                                                       `json:"portal_request_created"`
	NotificationActionEnabled             bool                                                                                                       `json:"notification_action_enabled"`
	CompatibilityCenterOpened             bool                                                                                                       `json:"compatibility_center_opened"`
	SupportBundleExported                 bool                                                                                                       `json:"support_bundle_exported"`
	SupportCaseCreated                    bool                                                                                                       `json:"support_case_created"`
	ProductionReadiness                   bool                                                                                                       `json:"production_readiness"`
	ProductionOwnershipReady              bool                                                                                                       `json:"production_ownership_ready"`
	AuthorizationItemCount                int                                                                                                        `json:"authorization_item_count"`
	RequiredAuthorizationItemCount        int                                                                                                        `json:"required_authorization_item_count"`
	ReadyAuthorizationItemCount           int                                                                                                        `json:"ready_authorization_item_count"`
	MissingAuthorizationItemCount         int                                                                                                        `json:"missing_authorization_item_count"`
	GrantedAuthorizationItemCount         int                                                                                                        `json:"granted_authorization_item_count"`
	AcceptedReceiptItemCount              int                                                                                                        `json:"accepted_receipt_item_count"`
	ConsumedReceiptItemCount              int                                                                                                        `json:"consumed_receipt_item_count"`
	AuthorizedConsumerItemCount           int                                                                                                        `json:"authorized_consumer_item_count"`
	EnabledConsumerItemCount              int                                                                                                        `json:"enabled_consumer_item_count"`
	LookupEnabledAuthorizationItemCount   int                                                                                                        `json:"lookup_enabled_authorization_item_count"`
	RawExposedAuthorizationItemCount      int                                                                                                        `json:"raw_exposed_authorization_item_count"`
	SideEffectAuthorizationItemCount      int                                                                                                        `json:"side_effect_authorization_item_count"`
	AuthorizationItems                    []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerAuthorizationItem  `json:"authorization_items"`
	AuthorizationItemIDs                  []string                                                                                                   `json:"authorization_item_ids"`
	RequiredBeforeConsumerAuthorization   []string                                                                                                   `json:"required_before_consumer_authorization"`
	Checks                                []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerAuthorizationCheck `json:"checks"`
	CheckIDs                              []string                                                                                                   `json:"check_ids"`
	Counts                                ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerAuthorizationCounts  `json:"counts"`
	RuntimeOwned                          bool                                                                                                       `json:"runtime_owned"`
	GoRuntimeBacked                       bool                                                                                                       `json:"go_runtime_backed"`
	KDEPolicyOwner                        bool                                                                                                       `json:"kde_policy_owner"`
	OfficialDesktopOnly                   bool                                                                                                       `json:"official_desktop_only"`
	PlasmaForkRequired                    bool                                                                                                       `json:"plasma_fork_required"`
	PlasmaSourceModified                  bool                                                                                                       `json:"plasma_source_modified"`
	SystemServiceStarted                  bool                                                                                                       `json:"system_service_started"`
	SessionBusClaimed                     bool                                                                                                       `json:"session_bus_claimed"`
	ProductionBusClaimed                  bool                                                                                                       `json:"production_bus_claimed"`
	ProductionOwnerEnabled                bool                                                                                                       `json:"production_owner_enabled"`
	ProductionActivationReady             bool                                                                                                       `json:"production_activation_ready"`
	WriteMethodsEnabled                   bool                                                                                                       `json:"write_methods_enabled"`
	RuntimeWritesEnabled                  bool                                                                                                       `json:"runtime_writes_enabled"`
	RequestObjectsCreated                 bool                                                                                                       `json:"request_objects_created"`
	RequestObjectsDispatched              bool                                                                                                       `json:"request_objects_dispatched"`
	NotificationSent                      bool                                                                                                       `json:"notification_sent"`
	NotificationDeliveryEnabled           bool                                                                                                       `json:"notification_delivery_enabled"`
	CompatibilityCenterPersisted          bool                                                                                                       `json:"compatibility_center_persisted"`
	DesktopFilesWritten                   bool                                                                                                       `json:"desktop_files_written"`
	MIMEAppsWritten                       bool                                                                                                       `json:"mimeapps_written"`
	ShellConfigurationWritten             bool                                                                                                       `json:"shell_configuration_written"`
	SettingsPersisted                     bool                                                                                                       `json:"settings_persisted"`
	AdapterInvocationEnabled              bool                                                                                                       `json:"adapter_invocation_enabled"`
	BackendLaunchEnabled                  bool                                                                                                       `json:"backend_launch_enabled"`
	BackendProcessStarted                 bool                                                                                                       `json:"backend_process_started"`
	SnapshotRestoreExecuted               bool                                                                                                       `json:"snapshot_restore_executed"`
	StateCleanupExecuted                  bool                                                                                                       `json:"state_cleanup_executed"`
	NetworkRequired                       bool                                                                                                       `json:"network_required"`
	HostRootModified                      bool                                                                                                       `json:"host_root_modified"`
	PrivilegedContainerRequired           bool                                                                                                       `json:"privileged_container_required"`
	CallerStateRootRequired               bool                                                                                                       `json:"caller_state_root_required"`
	StateRootPathExposed                  bool                                                                                                       `json:"state_root_path_exposed"`
	FilePathsExposed                      bool                                                                                                       `json:"file_paths_exposed"`
	FileContentRead                       bool                                                                                                       `json:"file_content_read"`
	RawCommandExposed                     bool                                                                                                       `json:"raw_command_exposed"`
	RawExecutableExposed                  bool                                                                                                       `json:"raw_executable_exposed"`
	BackendDetailsExposed                 bool                                                                                                       `json:"backend_details_exposed"`
	BlockedActions                        []string                                                                                                   `json:"blocked_actions"`
	NextRequirements                      []string                                                                                                   `json:"next_requirements"`
	DesktopSafeSummary                    string                                                                                                     `json:"desktop_safe_summary"`
}

type ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerAuthorizationItem struct {
	ID                                    string `json:"id"`
	ActionKind                            string `json:"action_kind"`
	AuthorizationArea                     string `json:"authorization_area"`
	RequiredEvidence                      string `json:"required_evidence"`
	EvidencePresent                       bool   `json:"evidence_present"`
	ConsumptionGateAuditConsumed          bool   `json:"consumption_gate_audit_consumed"`
	ConsumerRedactionAuditConsumed        bool   `json:"consumer_redaction_audit_consumed"`
	LookupRouteAuthorizationAuditConsumed bool   `json:"lookup_route_authorization_audit_consumed"`
	ConsumerAuthorizationModeled          bool   `json:"consumer_authorization_modeled"`
	ConsumerAuthorizationReady            bool   `json:"consumer_authorization_ready"`
	ReceiptPresent                        bool   `json:"receipt_present"`
	ReceiptAccepted                       bool   `json:"receipt_accepted"`
	ReceiptConsumed                       bool   `json:"receipt_consumed"`
	ConsumerAuthorizationGranted          bool   `json:"consumer_authorization_granted"`
	ConsumerConsumptionAuthorized         bool   `json:"consumer_consumption_authorized"`
	ConsumerEnablementAuthorized          bool   `json:"consumer_enablement_authorized"`
	KDEConsumerAuthorized                 bool   `json:"kde_consumer_authorized"`
	RuntimeConsumerAuthorized             bool   `json:"runtime_consumer_authorized"`
	KDEConsumerEnabled                    bool   `json:"kde_consumer_enabled"`
	RuntimeConsumerEnabled                bool   `json:"runtime_consumer_enabled"`
	LookupRouteEnabled                    bool   `json:"lookup_route_enabled"`
	OpaqueLookupEnabled                   bool   `json:"opaque_lookup_enabled"`
	RawResultExposed                      bool   `json:"raw_result_exposed"`
	UserVisible                           bool   `json:"user_visible"`
	ReviewOnly                            bool   `json:"review_only"`
	RuntimeOwned                          bool   `json:"runtime_owned"`
	GoRuntimeBacked                       bool   `json:"go_runtime_backed"`
	KDEPolicyOwner                        bool   `json:"kde_policy_owner"`
	CallerStateRootRequired               bool   `json:"caller_state_root_required"`
	StateRootPathExposed                  bool   `json:"state_root_path_exposed"`
	FilePathsExposed                      bool   `json:"file_paths_exposed"`
	FileContentRead                       bool   `json:"file_content_read"`
	RequestObjectCreated                  bool   `json:"request_object_created"`
	RequestObjectDispatched               bool   `json:"request_object_dispatched"`
	PortalRequestCreated                  bool   `json:"portal_request_created"`
	NotificationActionEnabled             bool   `json:"notification_action_enabled"`
	CompatibilityCenterOpened             bool   `json:"compatibility_center_opened"`
	SupportBundleExported                 bool   `json:"support_bundle_exported"`
	SupportCaseCreated                    bool   `json:"support_case_created"`
	ProductionReadiness                   bool   `json:"production_readiness"`
	ProductionOwnershipReady              bool   `json:"production_ownership_ready"`
	SideEffectsDisabled                   bool   `json:"side_effects_disabled"`
	HostRootModified                      bool   `json:"host_root_modified"`
	InternalDetailsExposed                bool   `json:"internal_details_exposed"`
	AuthorizationStatus                   string `json:"authorization_status"`
	NextRequirement                       string `json:"next_requirement"`
}

type ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerAuthorizationCheck struct {
	ID      string `json:"id"`
	Status  string `json:"status"`
	Summary string `json:"summary"`
}

type ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerAuthorizationCounts struct {
	Total   int `json:"total"`
	Passed  int `json:"passed"`
	Pending int `json:"pending"`
	Blocked int `json:"blocked"`
}

func NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerAuthorizationAuditPreview(root string) (ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerAuthorizationAuditPreview, error) {
	version, err := readProductionDBusGateReviewVersion(root)
	if err != nil {
		return ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerAuthorizationAuditPreview{}, err
	}
	sources := productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerAuthorizationSources(root)
	items := productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerAuthorizationItems(sources)
	consumptionReady := productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerAuthorizationConsumptionReady(sources.ConsumptionGateAudit)
	redactionReady := productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerAuthorizationRedactionReady(sources.ConsumerRedactionAudit)
	routeReady := productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerAuthorizationRouteReady(sources.LookupRouteAuthorizationAudit)
	guidanceReady := productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerAuthorizationGuidanceReady(sources.DispatchSheet)
	boundaryReady := consumptionReady && redactionReady && routeReady && guidanceReady && productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerAuthorizationItemsReady(items)
	preview := ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerAuthorizationAuditPreview{
		Version:                               version,
		SchemaVersion:                         "xnix.runtime.production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_receipt_consumer_authorization_audit.v1",
		RequestType:                           "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-receipt-consumer-authorization-audit-preview",
		AuditType:                             "receipt-notification-action-dry-run-result-lookup-consumer-enablement-receipt-consumer-authorization-audit",
		Source:                                "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-receipt-consumption-gate-audit-preview+production-receipt-notification-action-dry-run-result-lookup-consumer-redaction-audit-preview+production-receipt-notification-action-dry-run-result-lookup-route-authorization-audit-preview+claude-code-current-dispatch-picks",
		AuditDecision:                         "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-receipt-consumer-authorization-audit-blocked",
		ReceiptSchema:                         "xnix.runtime.production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_authorization_receipt.v1",
		OpaqueReceiptID:                       ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementAuthorizationReceiptID,
		ConsumerAuthorizationAuditRequired:    true,
		ConsumerAuthorizationAuditModeled:     true,
		ConsumptionGateAuditConsumed:          consumptionReady,
		ConsumerRedactionAuditConsumed:        redactionReady,
		LookupRouteAuthorizationAuditConsumed: routeReady,
		ConsumerAuthorizationGuidanceConsumed: guidanceReady,
		ConsumerAuthorizationBoundaryReady:    boundaryReady,
		ReceiptConsumptionGateReady:           consumptionReady,
		ConsumerRedactionBoundaryReady:        redactionReady,
		LookupRouteAuthorizationBoundaryReady: routeReady,
		ConsumerAuthorizationReady:            boundaryReady,
		AuthorizationItemCount:                len(items),
		RequiredAuthorizationItemCount:        5,
		ReadyAuthorizationItemCount:           productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerAuthorizationReadyCount(items),
		MissingAuthorizationItemCount:         productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerAuthorizationMissingCount(items),
		AuthorizationItems:                    items,
		AuthorizationItemIDs:                  productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerAuthorizationItemIDs(items),
		RequiredBeforeConsumerAuthorization: []string{
			"receipt consumption gate remains ready but does not consume a receipt",
			"KDE and Runtime consumer redaction remains modeled",
			"lookup route authorization remains modeled before route enablement",
			"consumer authorization grants remain disabled until a future explicit authorization action exists",
			"consumer enablement remains disabled until a separate enablement action consumes authorization",
		},
		RuntimeOwned:                     true,
		GoRuntimeBacked:                  true,
		OfficialDesktopOnly:              true,
		SideEffectAuthorizationItemCount: 0,
		BlockedActions: []string{
			"treat this audit as permission to authorize or enable KDE or Runtime consumers",
			"consume, accept, persist, replay, revoke, expire, or look up consumer enablement authorization receipts",
			"grant consumer authorization, enable lookup routes, persist redacted summaries, execute dry-runs, create request objects, dispatch actions, or send notifications",
			"claim production ownership, start services, enable Runtime writes, launch compatibility engines, expose raw results or paths, or mutate host root",
		},
		NextRequirements: []string{
			"Implement a separate consumer authorization action before any KDE or Runtime consumer can be authorized.",
			"Implement a separate consumer enablement gate after authorization grants are represented.",
			"Keep route enablement, lookup enablement, dry-run execution, result persistence, request creation, and dispatch disabled.",
			"Keep production ownership, Runtime writes, desktop side effects, engine launch, raw result exposure, and host mutation disabled.",
		},
		DesktopSafeSummary: "The production receipt notification action dry-run result lookup consumer enablement receipt consumer authorization audit models the explicit authorization boundary after the receipt consumption gate, but grants no authorization, enables no consumer or lookup route, consumes no receipt, persists no state, exposes no raw result data, launches no engine, and mutates no host state.",
	}
	preview.GrantedAuthorizationItemCount = productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerAuthorizationGrantedCount(items)
	preview.AcceptedReceiptItemCount = productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerAuthorizationAcceptedCount(items)
	preview.ConsumedReceiptItemCount = productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerAuthorizationConsumedCount(items)
	preview.AuthorizedConsumerItemCount = productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerAuthorizationAuthorizedCount(items)
	preview.EnabledConsumerItemCount = productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerAuthorizationEnabledCount(items)
	preview.LookupEnabledAuthorizationItemCount = productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerAuthorizationLookupEnabledCount(items)
	preview.RawExposedAuthorizationItemCount = productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerAuthorizationRawExposedCount(items)
	preview.SideEffectAuthorizationItemCount = productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerAuthorizationSideEffectCount(items)
	checks := productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerAuthorizationChecks(preview)
	preview.Checks = checks
	preview.CheckIDs = productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerAuthorizationCheckIDs(checks)
	preview.Counts = countProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerAuthorizationChecks(checks)
	if preview.Counts.Blocked == 0 {
		preview.AuditDecision = "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-receipt-consumer-authorization-audit-ready-authorization-disabled"
	}
	if err := validateNoBackendTerms(preview, "production receipt notification action dry-run result lookup consumer enablement receipt consumer authorization audit preview"); err != nil {
		return ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerAuthorizationAuditPreview{}, err
	}
	return preview, nil
}

type productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerAuthorizationSourceSet struct {
	ConsumptionGateAudit          string
	ConsumerRedactionAudit        string
	LookupRouteAuthorizationAudit string
	DispatchSheet                 string
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerAuthorizationSources(root string) productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerAuthorizationSourceSet {
	return productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerAuthorizationSourceSet{
		ConsumptionGateAudit:          productionAuthorizationReadSources(root, []string{"internal/runtime/owner/production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_receipt_consumption_gate_audit.go"}),
		ConsumerRedactionAudit:        productionAuthorizationReadSources(root, []string{"internal/runtime/owner/production_receipt_notification_action_dry_run_result_lookup_consumer_redaction_audit.go"}),
		LookupRouteAuthorizationAudit: productionAuthorizationReadSources(root, []string{"internal/runtime/owner/production_receipt_notification_action_dry_run_result_lookup_route_authorization_audit.go"}),
		DispatchSheet:                 productionAuthorizationReadSources(root, []string{"docs/claude-code-current-dispatch-picks.md"}),
	}
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerAuthorizationItems(sources productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerAuthorizationSourceSet) []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerAuthorizationItem {
	combined := sources.ConsumptionGateAudit + sources.ConsumerRedactionAudit + sources.LookupRouteAuthorizationAudit + sources.DispatchSheet
	return []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerAuthorizationItem{
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerAuthorizationItem("review-receipt-dry-run-result-lookup-consumer-enablement-consumer-authorization", "review", combined, []string{"review-receipt-dry-run-result-lookup-consumer-enablement-consumption-gate", "review-receipt-dry-run-result-lookup-consumer-redaction", "review-receipt-dry-run-result-lookup-route-authorization", "consumer authorization audit"}),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerAuthorizationItem("renew-receipt-dry-run-result-lookup-consumer-enablement-consumer-authorization", "renew", combined, []string{"renew-receipt-dry-run-result-lookup-consumer-enablement-consumption-gate", "renew-receipt-dry-run-result-lookup-consumer-redaction", "renew-receipt-dry-run-result-lookup-route-authorization", "consumer authorization audit"}),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerAuthorizationItem("open-compatibility-center-dry-run-result-lookup-consumer-enablement-consumer-authorization", "open-compatibility-center", combined, []string{"open-compatibility-center-dry-run-result-lookup-consumer-enablement-consumption-gate", "open-compatibility-center-dry-run-result-lookup-consumer-redaction", "open-compatibility-center-dry-run-result-lookup-route-authorization", "consumer authorization audit"}),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerAuthorizationItem("dismiss-receipt-dry-run-result-lookup-consumer-enablement-consumer-authorization", "dismiss", combined, []string{"dismiss-receipt-dry-run-result-lookup-consumer-enablement-consumption-gate", "dismiss-receipt-dry-run-result-lookup-consumer-redaction", "dismiss-receipt-dry-run-result-lookup-route-authorization", "consumer authorization audit"}),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerAuthorizationItem("support-info-dry-run-result-lookup-consumer-enablement-consumer-authorization", "support-info", combined, []string{"support-info-dry-run-result-lookup-consumer-enablement-consumption-gate", "support-info-dry-run-result-lookup-consumer-redaction", "support-info-dry-run-result-lookup-route-authorization", "consumer authorization audit"}),
	}
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerAuthorizationItem(id string, actionKind string, source string, tokens []string) ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerAuthorizationItem {
	ready := productionAuthorizationHasAll(source, tokens)
	status := "missing-consumer-enablement-receipt-consumer-authorization-evidence"
	if ready {
		status = "consumer-enablement-receipt-consumer-authorization-modeled-authorization-disabled"
	}
	return ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerAuthorizationItem{
		ID:                                    id,
		ActionKind:                            actionKind,
		AuthorizationArea:                     actionKind + "-consumer-authorization",
		RequiredEvidence:                      "future consumer authorization requires receipt consumption gate, redaction, and lookup route authorization evidence",
		EvidencePresent:                       ready,
		ConsumptionGateAuditConsumed:          ready,
		ConsumerRedactionAuditConsumed:        ready,
		LookupRouteAuthorizationAuditConsumed: ready,
		ConsumerAuthorizationModeled:          ready,
		ConsumerAuthorizationReady:            ready,
		UserVisible:                           ready,
		ReviewOnly:                            true,
		RuntimeOwned:                          true,
		GoRuntimeBacked:                       true,
		SideEffectsDisabled:                   true,
		AuthorizationStatus:                   status,
		NextRequirement:                       "grant consumer authorization only through a separate future authorization action",
	}
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerAuthorizationChecks(preview ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerAuthorizationAuditPreview) []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerAuthorizationCheck {
	return []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerAuthorizationCheck{
		{"consumption-gate-audit-consumed", productionAuthorizationPassBlocked(preview.ConsumptionGateAuditConsumed), "The consumer authorization audit consumes the receipt consumption gate."},
		{"consumer-redaction-audit-consumed", productionAuthorizationPassBlocked(preview.ConsumerRedactionAuditConsumed), "The consumer authorization audit consumes consumer redaction evidence."},
		{"lookup-route-authorization-audit-consumed", productionAuthorizationPassBlocked(preview.LookupRouteAuthorizationAuditConsumed), "The consumer authorization audit consumes lookup route authorization evidence."},
		{"consumer-authorization-guidance-consumed", productionAuthorizationPassBlocked(preview.ConsumerAuthorizationGuidanceConsumed), "The consumer authorization audit consumes dispatch guidance."},
		{"consumer-authorization-boundary-modeled-only", productionAuthorizationPassBlocked(preview.ConsumerAuthorizationAuditRequired && preview.ConsumerAuthorizationAuditModeled && preview.ConsumerAuthorizationBoundaryReady && preview.ConsumerAuthorizationReady && !preview.ConsumerAuthorizationGranted && !preview.ConsumerConsumptionAuthorized && !preview.ConsumerEnablementAuthorized), "The authorization boundary is modeled without granting authorization."},
		{"five-consumer-authorization-items-present", productionAuthorizationPassBlocked(preview.AuthorizationItemCount == 5 && preview.RequiredAuthorizationItemCount == 5 && preview.MissingAuthorizationItemCount == 0), "Five consumer authorization items are modeled."},
		{"consumer-authorization-items-ready-authorization-disabled", productionAuthorizationPassBlocked(preview.ReadyAuthorizationItemCount == 5 && productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerAuthorizationItemsReady(preview.AuthorizationItems)), "Every authorization item is ready while grants remain disabled."},
		{"receipt-consumption-and-consumer-authorization-disabled", productionAuthorizationPassBlocked(!preview.ReceiptAccepted && !preview.ReceiptConsumed && !preview.AuthorizationAccepted && !preview.ConsumerAuthorizationGranted && !preview.ConsumerConsumptionAuthorized && !preview.ConsumerEnablementAuthorized && preview.GrantedAuthorizationItemCount == 0 && preview.ConsumedReceiptItemCount == 0 && preview.AuthorizedConsumerItemCount == 0), "Receipt consumption and consumer authorization remain disabled."},
		{"consumer-lookup-request-and-support-disabled", productionAuthorizationPassBlocked(!preview.KDEConsumerAuthorized && !preview.RuntimeConsumerAuthorized && !preview.KDEConsumerEnabled && !preview.RuntimeConsumerEnabled && !preview.LookupRouteAuthorized && !preview.LookupRouteEnabled && !preview.OpaqueLookupEnabled && !preview.RedactedSummaryPersisted && !preview.DryRunResultPersisted && !preview.DispatchDryRunExecuted && !preview.RequestObjectCreationEnabled && !preview.RequestObjectDispatchEnabled && !preview.PortalRequestCreated && !preview.NotificationActionEnabled && !preview.CompatibilityCenterOpened && !preview.SupportBundleExported && !preview.SupportCaseCreated && preview.EnabledConsumerItemCount == 0 && preview.LookupEnabledAuthorizationItemCount == 0), "Consumers, lookup, request, notification, navigation, and support side effects remain disabled."},
		{"production-and-host-boundary-closed", productionAuthorizationPassBlocked(!preview.SystemServiceStarted && !preview.SessionBusClaimed && !preview.ProductionBusClaimed && !preview.ProductionOwnerEnabled && !preview.ProductionActivationReady && !preview.ProductionReadiness && !preview.ProductionOwnershipReady && !preview.WriteMethodsEnabled && !preview.RuntimeWritesEnabled && !preview.AdapterInvocationEnabled && !preview.BackendLaunchEnabled && !preview.BackendProcessStarted && !preview.NetworkRequired && !preview.HostRootModified && !preview.PrivilegedContainerRequired && !preview.StateRootPathExposed && !preview.FilePathsExposed && !preview.FileContentRead && !preview.RawResultExposed && !preview.RawCommandExposed && !preview.RawExecutableExposed && !preview.BackendDetailsExposed && preview.RawExposedAuthorizationItemCount == 0 && preview.SideEffectAuthorizationItemCount == 0), "Production ownership, writes, engine launch, unsafe data exposure, and host mutation remain disabled."},
	}
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerAuthorizationConsumptionReady(source string) bool {
	return productionAuthorizationHasAll(source, []string{"production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-receipt-consumption-gate-audit-preview", "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-receipt-consumption-gate-audit-ready-consumption-disabled", "ReceiptConsumptionGateReady", "ReceiptConsumed"})
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerAuthorizationRedactionReady(source string) bool {
	return productionAuthorizationHasAll(source, []string{"production-receipt-notification-action-dry-run-result-lookup-consumer-redaction-audit-preview", "production-receipt-notification-action-dry-run-result-lookup-consumer-redaction-audit-ready-consumers-disabled", "ConsumerRedactionReady"})
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerAuthorizationRouteReady(source string) bool {
	return productionAuthorizationHasAll(source, []string{"production-receipt-notification-action-dry-run-result-lookup-route-authorization-audit-preview", "production-receipt-notification-action-dry-run-result-lookup-route-authorization-audit-ready-route-disabled", "LookupRouteAuthorizationReady"})
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerAuthorizationGuidanceReady(source string) bool {
	return productionAuthorizationHasAll(source, []string{"consumer authorization audit", "consumed consumer enablement authorization receipt", "enable KDE or Runtime consumers", "side effects"})
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerAuthorizationReadyCount(items []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerAuthorizationItem) int {
	count := 0
	for _, item := range items {
		if item.EvidencePresent {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerAuthorizationMissingCount(items []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerAuthorizationItem) int {
	count := 0
	for _, item := range items {
		if !item.EvidencePresent {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerAuthorizationGrantedCount(items []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerAuthorizationItem) int {
	count := 0
	for _, item := range items {
		if item.ConsumerAuthorizationGranted {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerAuthorizationAcceptedCount(items []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerAuthorizationItem) int {
	count := 0
	for _, item := range items {
		if item.ReceiptAccepted {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerAuthorizationConsumedCount(items []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerAuthorizationItem) int {
	count := 0
	for _, item := range items {
		if item.ReceiptConsumed {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerAuthorizationAuthorizedCount(items []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerAuthorizationItem) int {
	count := 0
	for _, item := range items {
		if item.ConsumerConsumptionAuthorized || item.ConsumerEnablementAuthorized || item.KDEConsumerAuthorized || item.RuntimeConsumerAuthorized {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerAuthorizationEnabledCount(items []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerAuthorizationItem) int {
	count := 0
	for _, item := range items {
		if item.KDEConsumerEnabled || item.RuntimeConsumerEnabled {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerAuthorizationLookupEnabledCount(items []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerAuthorizationItem) int {
	count := 0
	for _, item := range items {
		if item.LookupRouteEnabled || item.OpaqueLookupEnabled {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerAuthorizationRawExposedCount(items []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerAuthorizationItem) int {
	count := 0
	for _, item := range items {
		if item.RawResultExposed || item.StateRootPathExposed || item.FilePathsExposed || item.FileContentRead || item.InternalDetailsExposed {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerAuthorizationSideEffectCount(items []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerAuthorizationItem) int {
	count := 0
	for _, item := range items {
		if !item.SideEffectsDisabled || item.HostRootModified || item.RequestObjectCreated || item.RequestObjectDispatched || item.PortalRequestCreated || item.NotificationActionEnabled || item.CompatibilityCenterOpened || item.SupportBundleExported || item.SupportCaseCreated {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerAuthorizationItemIDs(items []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerAuthorizationItem) []string {
	ids := make([]string, 0, len(items))
	for _, item := range items {
		ids = append(ids, item.ID)
	}
	return ids
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerAuthorizationItemsReady(items []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerAuthorizationItem) bool {
	if len(items) != 5 {
		return false
	}
	for _, item := range items {
		if !item.EvidencePresent || !item.ConsumptionGateAuditConsumed || !item.ConsumerRedactionAuditConsumed || !item.LookupRouteAuthorizationAuditConsumed || !item.ConsumerAuthorizationModeled || !item.ConsumerAuthorizationReady || item.AuthorizationStatus != "consumer-enablement-receipt-consumer-authorization-modeled-authorization-disabled" {
			return false
		}
		if item.ReceiptAccepted || item.ReceiptConsumed || item.ConsumerAuthorizationGranted || item.ConsumerConsumptionAuthorized || item.ConsumerEnablementAuthorized || item.KDEConsumerAuthorized || item.RuntimeConsumerAuthorized || item.KDEConsumerEnabled || item.RuntimeConsumerEnabled || item.LookupRouteEnabled || item.OpaqueLookupEnabled || item.RawResultExposed || !item.ReviewOnly || !item.SideEffectsDisabled || item.HostRootModified || item.InternalDetailsExposed {
			return false
		}
	}
	return true
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerAuthorizationCheckIDs(checks []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerAuthorizationCheck) []string {
	ids := make([]string, 0, len(checks))
	for _, check := range checks {
		ids = append(ids, check.ID)
	}
	return ids
}

func countProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerAuthorizationChecks(checks []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerAuthorizationCheck) ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerAuthorizationCounts {
	counts := ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerAuthorizationCounts{Total: len(checks)}
	for _, check := range checks {
		switch check.Status {
		case "pass":
			counts.Passed++
		case "pending":
			counts.Pending++
		case "blocked":
			counts.Blocked++
		}
	}
	return counts
}
