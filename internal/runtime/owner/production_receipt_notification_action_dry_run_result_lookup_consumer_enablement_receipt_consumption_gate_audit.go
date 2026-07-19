package owner

type ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumptionGateAuditPreview struct {
	Version                                   string                                                                                                    `json:"version"`
	SchemaVersion                             string                                                                                                    `json:"schema_version"`
	RequestType                               string                                                                                                    `json:"request_type"`
	AuditType                                 string                                                                                                    `json:"audit_type"`
	Source                                    string                                                                                                    `json:"source"`
	AuditDecision                             string                                                                                                    `json:"audit_decision"`
	ReceiptSchema                             string                                                                                                    `json:"receipt_schema"`
	OpaqueReceiptID                           string                                                                                                    `json:"opaque_receipt_id"`
	ConsumptionGateRequired                   bool                                                                                                      `json:"consumption_gate_required"`
	ConsumptionGateModeled                    bool                                                                                                      `json:"consumption_gate_modeled"`
	AcceptanceAuthorizationAuditConsumed      bool                                                                                                      `json:"acceptance_authorization_audit_consumed"`
	ConsumerRedactionAuditConsumed            bool                                                                                                      `json:"consumer_redaction_audit_consumed"`
	LookupRouteAuthorizationAuditConsumed     bool                                                                                                      `json:"lookup_route_authorization_audit_consumed"`
	ConsumptionGateGuidanceConsumed           bool                                                                                                      `json:"consumption_gate_guidance_consumed"`
	ConsumptionGateBoundaryReady              bool                                                                                                      `json:"consumption_gate_boundary_ready"`
	ReceiptAcceptanceAuthorizationReady       bool                                                                                                      `json:"receipt_acceptance_authorization_ready"`
	ConsumerRedactionBoundaryReady            bool                                                                                                      `json:"consumer_redaction_boundary_ready"`
	LookupRouteAuthorizationBoundaryReady     bool                                                                                                      `json:"lookup_route_authorization_boundary_ready"`
	ConsumerEnablementReceiptConsumptionReady bool                                                                                                      `json:"consumer_enablement_receipt_consumption_ready"`
	ReceiptConsumptionGateReady               bool                                                                                                      `json:"receipt_consumption_gate_ready"`
	ReceiptPresent                            bool                                                                                                      `json:"receipt_present"`
	ReceiptPersisted                          bool                                                                                                      `json:"receipt_persisted"`
	ReceiptAccepted                           bool                                                                                                      `json:"receipt_accepted"`
	ReceiptConsumed                           bool                                                                                                      `json:"receipt_consumed"`
	AuthorizationAccepted                     bool                                                                                                      `json:"authorization_accepted"`
	ReceiptWriterEnabled                      bool                                                                                                      `json:"receipt_writer_enabled"`
	ReceiptPersistenceEnabled                 bool                                                                                                      `json:"receipt_persistence_enabled"`
	ReceiptLookupWritesEnabled                bool                                                                                                      `json:"receipt_lookup_writes_enabled"`
	ReceiptReplayEnabled                      bool                                                                                                      `json:"receipt_replay_enabled"`
	ReceiptExpiryWriteEnabled                 bool                                                                                                      `json:"receipt_expiry_write_enabled"`
	ReceiptRevocationWriteEnabled             bool                                                                                                      `json:"receipt_revocation_write_enabled"`
	ConsumerConsumptionAuthorized             bool                                                                                                      `json:"consumer_consumption_authorized"`
	ConsumerEnablementAuthorized              bool                                                                                                      `json:"consumer_enablement_authorized"`
	KDEConsumerEnabled                        bool                                                                                                      `json:"kde_consumer_enabled"`
	RuntimeConsumerEnabled                    bool                                                                                                      `json:"runtime_consumer_enabled"`
	LookupRouteAuthorized                     bool                                                                                                      `json:"lookup_route_authorized"`
	LookupRouteEnabled                        bool                                                                                                      `json:"lookup_route_enabled"`
	LookupRoutePersisted                      bool                                                                                                      `json:"lookup_route_persisted"`
	OpaqueLookupEnabled                       bool                                                                                                      `json:"opaque_lookup_enabled"`
	OpaqueLookupPersisted                     bool                                                                                                      `json:"opaque_lookup_persisted"`
	RedactedSummaryPersisted                  bool                                                                                                      `json:"redacted_summary_persisted"`
	RawResultExposed                          bool                                                                                                      `json:"raw_result_exposed"`
	RuntimeDiagnosticsPersisted               bool                                                                                                      `json:"runtime_diagnostics_persisted"`
	DryRunResultPersisted                     bool                                                                                                      `json:"dry_run_result_persisted"`
	DispatchDryRunExecuted                    bool                                                                                                      `json:"dispatch_dry_run_executed"`
	RequestObjectCreationEnabled              bool                                                                                                      `json:"request_object_creation_enabled"`
	RequestObjectDispatchEnabled              bool                                                                                                      `json:"request_object_dispatch_enabled"`
	PortalRequestCreated                      bool                                                                                                      `json:"portal_request_created"`
	NotificationActionEnabled                 bool                                                                                                      `json:"notification_action_enabled"`
	CompatibilityCenterOpened                 bool                                                                                                      `json:"compatibility_center_opened"`
	SupportBundleExported                     bool                                                                                                      `json:"support_bundle_exported"`
	SupportCaseCreated                        bool                                                                                                      `json:"support_case_created"`
	ProductionReadiness                       bool                                                                                                      `json:"production_readiness"`
	ProductionOwnershipReady                  bool                                                                                                      `json:"production_ownership_ready"`
	GateItemCount                             int                                                                                                       `json:"gate_item_count"`
	RequiredGateItemCount                     int                                                                                                       `json:"required_gate_item_count"`
	ReadyGateItemCount                        int                                                                                                       `json:"ready_gate_item_count"`
	MissingGateItemCount                      int                                                                                                       `json:"missing_gate_item_count"`
	AcceptedReceiptItemCount                  int                                                                                                       `json:"accepted_receipt_item_count"`
	ConsumedReceiptItemCount                  int                                                                                                       `json:"consumed_receipt_item_count"`
	AuthorizedConsumerItemCount               int                                                                                                       `json:"authorized_consumer_item_count"`
	EnabledConsumerItemCount                  int                                                                                                       `json:"enabled_consumer_item_count"`
	LookupEnabledGateItemCount                int                                                                                                       `json:"lookup_enabled_gate_item_count"`
	RawExposedGateItemCount                   int                                                                                                       `json:"raw_exposed_gate_item_count"`
	SideEffectGateItemCount                   int                                                                                                       `json:"side_effect_gate_item_count"`
	GateItems                                 []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumptionGateAuditItem  `json:"gate_items"`
	GateItemIDs                               []string                                                                                                  `json:"gate_item_ids"`
	RequiredBeforeReceiptConsumption          []string                                                                                                  `json:"required_before_receipt_consumption"`
	Checks                                    []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumptionGateAuditCheck `json:"checks"`
	CheckIDs                                  []string                                                                                                  `json:"check_ids"`
	Counts                                    ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumptionGateAuditCounts  `json:"counts"`
	RuntimeOwned                              bool                                                                                                      `json:"runtime_owned"`
	GoRuntimeBacked                           bool                                                                                                      `json:"go_runtime_backed"`
	KDEPolicyOwner                            bool                                                                                                      `json:"kde_policy_owner"`
	OfficialDesktopOnly                       bool                                                                                                      `json:"official_desktop_only"`
	PlasmaForkRequired                        bool                                                                                                      `json:"plasma_fork_required"`
	PlasmaSourceModified                      bool                                                                                                      `json:"plasma_source_modified"`
	SystemServiceStarted                      bool                                                                                                      `json:"system_service_started"`
	SessionBusClaimed                         bool                                                                                                      `json:"session_bus_claimed"`
	ProductionBusClaimed                      bool                                                                                                      `json:"production_bus_claimed"`
	ProductionOwnerEnabled                    bool                                                                                                      `json:"production_owner_enabled"`
	ProductionActivationReady                 bool                                                                                                      `json:"production_activation_ready"`
	WriteMethodsEnabled                       bool                                                                                                      `json:"write_methods_enabled"`
	RuntimeWritesEnabled                      bool                                                                                                      `json:"runtime_writes_enabled"`
	RequestObjectsCreated                     bool                                                                                                      `json:"request_objects_created"`
	RequestObjectsDispatched                  bool                                                                                                      `json:"request_objects_dispatched"`
	NotificationSent                          bool                                                                                                      `json:"notification_sent"`
	NotificationDeliveryEnabled               bool                                                                                                      `json:"notification_delivery_enabled"`
	CompatibilityCenterPersisted              bool                                                                                                      `json:"compatibility_center_persisted"`
	DesktopFilesWritten                       bool                                                                                                      `json:"desktop_files_written"`
	MIMEAppsWritten                           bool                                                                                                      `json:"mimeapps_written"`
	ShellConfigurationWritten                 bool                                                                                                      `json:"shell_configuration_written"`
	SettingsPersisted                         bool                                                                                                      `json:"settings_persisted"`
	AdapterInvocationEnabled                  bool                                                                                                      `json:"adapter_invocation_enabled"`
	BackendLaunchEnabled                      bool                                                                                                      `json:"backend_launch_enabled"`
	BackendProcessStarted                     bool                                                                                                      `json:"backend_process_started"`
	SnapshotRestoreExecuted                   bool                                                                                                      `json:"snapshot_restore_executed"`
	StateCleanupExecuted                      bool                                                                                                      `json:"state_cleanup_executed"`
	NetworkRequired                           bool                                                                                                      `json:"network_required"`
	HostRootModified                          bool                                                                                                      `json:"host_root_modified"`
	PrivilegedContainerRequired               bool                                                                                                      `json:"privileged_container_required"`
	CallerStateRootRequired                   bool                                                                                                      `json:"caller_state_root_required"`
	StateRootPathExposed                      bool                                                                                                      `json:"state_root_path_exposed"`
	FilePathsExposed                          bool                                                                                                      `json:"file_paths_exposed"`
	FileContentRead                           bool                                                                                                      `json:"file_content_read"`
	RawCommandExposed                         bool                                                                                                      `json:"raw_command_exposed"`
	RawExecutableExposed                      bool                                                                                                      `json:"raw_executable_exposed"`
	BackendDetailsExposed                     bool                                                                                                      `json:"backend_details_exposed"`
	BlockedActions                            []string                                                                                                  `json:"blocked_actions"`
	NextRequirements                          []string                                                                                                  `json:"next_requirements"`
	DesktopSafeSummary                        string                                                                                                    `json:"desktop_safe_summary"`
}

type ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumptionGateAuditItem struct {
	ID                                    string `json:"id"`
	ActionKind                            string `json:"action_kind"`
	GateArea                              string `json:"gate_area"`
	RequiredEvidence                      string `json:"required_evidence"`
	EvidencePresent                       bool   `json:"evidence_present"`
	AcceptanceAuthorizationAuditConsumed  bool   `json:"acceptance_authorization_audit_consumed"`
	ConsumerRedactionAuditConsumed        bool   `json:"consumer_redaction_audit_consumed"`
	LookupRouteAuthorizationAuditConsumed bool   `json:"lookup_route_authorization_audit_consumed"`
	ConsumptionGateModeled                bool   `json:"consumption_gate_modeled"`
	ReceiptConsumptionGateReady           bool   `json:"receipt_consumption_gate_ready"`
	ReceiptPresent                        bool   `json:"receipt_present"`
	ReceiptPersisted                      bool   `json:"receipt_persisted"`
	ReceiptAccepted                       bool   `json:"receipt_accepted"`
	ReceiptConsumed                       bool   `json:"receipt_consumed"`
	AuthorizationAccepted                 bool   `json:"authorization_accepted"`
	ReceiptWriterEnabled                  bool   `json:"receipt_writer_enabled"`
	ReceiptPersistenceEnabled             bool   `json:"receipt_persistence_enabled"`
	ReceiptLookupWritesEnabled            bool   `json:"receipt_lookup_writes_enabled"`
	ReceiptReplayEnabled                  bool   `json:"receipt_replay_enabled"`
	ReceiptExpiryWriteEnabled             bool   `json:"receipt_expiry_write_enabled"`
	ReceiptRevocationWriteEnabled         bool   `json:"receipt_revocation_write_enabled"`
	ConsumerConsumptionAuthorized         bool   `json:"consumer_consumption_authorized"`
	ConsumerEnablementAuthorized          bool   `json:"consumer_enablement_authorized"`
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
	GateStatus                            string `json:"gate_status"`
	NextRequirement                       string `json:"next_requirement"`
}

type ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumptionGateAuditCheck struct {
	ID      string `json:"id"`
	Status  string `json:"status"`
	Summary string `json:"summary"`
}

type ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumptionGateAuditCounts struct {
	Total   int `json:"total"`
	Passed  int `json:"passed"`
	Pending int `json:"pending"`
	Blocked int `json:"blocked"`
}

func NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumptionGateAuditPreview(root string) (ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumptionGateAuditPreview, error) {
	version, err := readProductionDBusGateReviewVersion(root)
	if err != nil {
		return ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumptionGateAuditPreview{}, err
	}
	sources := productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumptionGateAuditSources(root)
	items := productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumptionGateItems(sources)
	acceptanceReady := productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumptionGateAcceptanceReady(sources.AcceptanceAuthorizationAudit)
	redactionReady := productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumptionGateRedactionReady(sources.ConsumerRedactionAudit)
	routeReady := productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumptionGateRouteReady(sources.LookupRouteAuthorizationAudit)
	guidanceReady := productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumptionGateGuidanceReady(sources.DispatchSheet)
	boundaryReady := acceptanceReady && redactionReady && routeReady && guidanceReady && productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumptionGateItemsReady(items)
	preview := ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumptionGateAuditPreview{
		Version:                                   version,
		SchemaVersion:                             "xnix.runtime.production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_receipt_consumption_gate_audit.v1",
		RequestType:                               "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-receipt-consumption-gate-audit-preview",
		AuditType:                                 "receipt-notification-action-dry-run-result-lookup-consumer-enablement-receipt-consumption-gate-audit",
		Source:                                    "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-receipt-acceptance-authorization-audit-preview+production-receipt-notification-action-dry-run-result-lookup-consumer-redaction-audit-preview+production-receipt-notification-action-dry-run-result-lookup-route-authorization-audit-preview+claude-code-current-dispatch-picks",
		AuditDecision:                             "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-receipt-consumption-gate-audit-blocked",
		ReceiptSchema:                             "xnix.runtime.production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_authorization_receipt.v1",
		OpaqueReceiptID:                           ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementAuthorizationReceiptID,
		ConsumptionGateRequired:                   true,
		ConsumptionGateModeled:                    true,
		AcceptanceAuthorizationAuditConsumed:      acceptanceReady,
		ConsumerRedactionAuditConsumed:            redactionReady,
		LookupRouteAuthorizationAuditConsumed:     routeReady,
		ConsumptionGateGuidanceConsumed:           guidanceReady,
		ConsumptionGateBoundaryReady:              boundaryReady,
		ReceiptAcceptanceAuthorizationReady:       acceptanceReady,
		ConsumerRedactionBoundaryReady:            redactionReady,
		LookupRouteAuthorizationBoundaryReady:     routeReady,
		ConsumerEnablementReceiptConsumptionReady: boundaryReady,
		ReceiptConsumptionGateReady:               boundaryReady,
		ReceiptPresent:                            false,
		ReceiptPersisted:                          false,
		ReceiptAccepted:                           false,
		ReceiptConsumed:                           false,
		AuthorizationAccepted:                     false,
		ReceiptWriterEnabled:                      false,
		ReceiptPersistenceEnabled:                 false,
		ReceiptLookupWritesEnabled:                false,
		ReceiptReplayEnabled:                      false,
		ReceiptExpiryWriteEnabled:                 false,
		ReceiptRevocationWriteEnabled:             false,
		ConsumerConsumptionAuthorized:             false,
		ConsumerEnablementAuthorized:              false,
		KDEConsumerEnabled:                        false,
		RuntimeConsumerEnabled:                    false,
		LookupRouteAuthorized:                     false,
		LookupRouteEnabled:                        false,
		LookupRoutePersisted:                      false,
		OpaqueLookupEnabled:                       false,
		OpaqueLookupPersisted:                     false,
		RedactedSummaryPersisted:                  false,
		RawResultExposed:                          false,
		RuntimeDiagnosticsPersisted:               false,
		DryRunResultPersisted:                     false,
		DispatchDryRunExecuted:                    false,
		RequestObjectCreationEnabled:              false,
		RequestObjectDispatchEnabled:              false,
		PortalRequestCreated:                      false,
		NotificationActionEnabled:                 false,
		CompatibilityCenterOpened:                 false,
		SupportBundleExported:                     false,
		SupportCaseCreated:                        false,
		ProductionReadiness:                       false,
		ProductionOwnershipReady:                  false,
		GateItemCount:                             len(items),
		RequiredGateItemCount:                     5,
		ReadyGateItemCount:                        productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumptionGateReadyCount(items),
		MissingGateItemCount:                      productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumptionGateMissingCount(items),
		GateItems:                                 items,
		GateItemIDs:                               productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumptionGateItemIDs(items),
		RequiredBeforeReceiptConsumption: []string{
			"accepted consumer enablement authorization receipt remains behind a separate write and persistence boundary",
			"receipt acceptance authorization remains ready but does not accept or consume a receipt",
			"KDE and Runtime consumer redaction remains modeled before consumer authorization",
			"lookup route authorization remains modeled before route enablement",
			"consumer authorization remains disabled until a future accepted receipt is consumed by an explicit gate",
		},
		RuntimeOwned:                 true,
		GoRuntimeBacked:              true,
		KDEPolicyOwner:               false,
		OfficialDesktopOnly:          true,
		PlasmaForkRequired:           false,
		PlasmaSourceModified:         false,
		SystemServiceStarted:         false,
		SessionBusClaimed:            false,
		ProductionBusClaimed:         false,
		ProductionOwnerEnabled:       false,
		ProductionActivationReady:    false,
		WriteMethodsEnabled:          false,
		RuntimeWritesEnabled:         false,
		RequestObjectsCreated:        false,
		RequestObjectsDispatched:     false,
		NotificationSent:             false,
		NotificationDeliveryEnabled:  false,
		CompatibilityCenterPersisted: false,
		DesktopFilesWritten:          false,
		MIMEAppsWritten:              false,
		ShellConfigurationWritten:    false,
		SettingsPersisted:            false,
		AdapterInvocationEnabled:     false,
		BackendLaunchEnabled:         false,
		BackendProcessStarted:        false,
		SnapshotRestoreExecuted:      false,
		StateCleanupExecuted:         false,
		NetworkRequired:              false,
		HostRootModified:             false,
		PrivilegedContainerRequired:  false,
		CallerStateRootRequired:      false,
		StateRootPathExposed:         false,
		FilePathsExposed:             false,
		FileContentRead:              false,
		RawCommandExposed:            false,
		RawExecutableExposed:         false,
		BackendDetailsExposed:        false,
		BlockedActions: []string{
			"treat this audit as permission to consume consumer enablement authorization receipts",
			"accept, consume, persist, replay, revoke, expire, or look up consumer enablement authorization receipts",
			"authorize or enable KDE or Runtime consumers from this consumption gate audit",
			"enable lookup routes, persist redacted summaries, execute dry-runs, create request objects, dispatch actions, or send notifications",
			"claim production ownership, start services, enable Runtime writes, launch compatibility engines, expose raw results or paths, or mutate host root",
		},
		NextRequirements: []string{
			"Implement a real accepted receipt materialization path only after human authorization, writer authorization, persistence authorization, acceptance authorization, and replay protection are all present.",
			"Implement a separate consumer authorization action before KDE or Runtime consumers can be enabled.",
			"Keep route enablement, lookup enablement, dry-run execution, result persistence, request creation, and dispatch disabled.",
			"Keep production ownership, Runtime writes, desktop side effects, engine launch, raw result exposure, and host mutation disabled.",
		},
		DesktopSafeSummary: "The production receipt notification action dry-run result lookup consumer enablement receipt consumption gate audit models the final fail-closed boundary before a future accepted consumer enablement authorization receipt can authorize KDE or Runtime consumers, but it consumes no receipt, enables no consumer or lookup route, persists no state, exposes no raw result data, launches no engine, and mutates no host state.",
	}
	preview.AcceptedReceiptItemCount = productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumptionGateAcceptedCount(items)
	preview.ConsumedReceiptItemCount = productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumptionGateConsumedCount(items)
	preview.AuthorizedConsumerItemCount = productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumptionGateAuthorizedCount(items)
	preview.EnabledConsumerItemCount = productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumptionGateEnabledCount(items)
	preview.LookupEnabledGateItemCount = productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumptionGateLookupEnabledCount(items)
	preview.RawExposedGateItemCount = productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumptionGateRawExposedCount(items)
	preview.SideEffectGateItemCount = productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumptionGateSideEffectCount(items)
	checks := productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumptionGateChecks(preview)
	preview.Checks = checks
	preview.CheckIDs = productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumptionGateCheckIDs(checks)
	preview.Counts = countProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumptionGateChecks(checks)
	if preview.Counts.Blocked == 0 {
		preview.AuditDecision = "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-receipt-consumption-gate-audit-ready-consumption-disabled"
	}
	if err := validateNoBackendTerms(preview, "production receipt notification action dry-run result lookup consumer enablement receipt consumption gate audit preview"); err != nil {
		return ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumptionGateAuditPreview{}, err
	}
	return preview, nil
}

type productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumptionGateSourceSet struct {
	AcceptanceAuthorizationAudit  string
	ConsumerRedactionAudit        string
	LookupRouteAuthorizationAudit string
	DispatchSheet                 string
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumptionGateAuditSources(root string) productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumptionGateSourceSet {
	return productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumptionGateSourceSet{
		AcceptanceAuthorizationAudit:  productionAuthorizationReadSources(root, []string{"internal/runtime/owner/production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_receipt_acceptance_authorization_audit.go"}),
		ConsumerRedactionAudit:        productionAuthorizationReadSources(root, []string{"internal/runtime/owner/production_receipt_notification_action_dry_run_result_lookup_consumer_redaction_audit.go"}),
		LookupRouteAuthorizationAudit: productionAuthorizationReadSources(root, []string{"internal/runtime/owner/production_receipt_notification_action_dry_run_result_lookup_route_authorization_audit.go"}),
		DispatchSheet:                 productionAuthorizationReadSources(root, []string{"docs/claude-code-current-dispatch-picks.md"}),
	}
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumptionGateItems(sources productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumptionGateSourceSet) []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumptionGateAuditItem {
	combined := sources.AcceptanceAuthorizationAudit + sources.ConsumerRedactionAudit + sources.LookupRouteAuthorizationAudit + sources.DispatchSheet
	return []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumptionGateAuditItem{
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumptionGateItem("review-receipt-dry-run-result-lookup-consumer-enablement-consumption-gate", "review", "review-consumer-enablement-receipt-consumption-gate", combined, []string{"review-receipt-dry-run-result-lookup-consumer-enablement-acceptance-authorization", "review-receipt-dry-run-result-lookup-consumer-redaction", "review-receipt-dry-run-result-lookup-route-authorization", "receipt consumption gate audit"}, "consume an accepted review receipt only after a separate consumer authorization gate exists"),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumptionGateItem("renew-receipt-dry-run-result-lookup-consumer-enablement-consumption-gate", "renew", "renewal-consumer-enablement-receipt-consumption-gate", combined, []string{"renew-receipt-dry-run-result-lookup-consumer-enablement-acceptance-authorization", "renew-receipt-dry-run-result-lookup-consumer-redaction", "renew-receipt-dry-run-result-lookup-route-authorization", "receipt consumption gate audit"}, "consume an accepted renewal receipt only after a separate consumer authorization gate exists"),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumptionGateItem("open-compatibility-center-dry-run-result-lookup-consumer-enablement-consumption-gate", "open-compatibility-center", "navigation-consumer-enablement-receipt-consumption-gate", combined, []string{"open-compatibility-center-dry-run-result-lookup-consumer-enablement-acceptance-authorization", "open-compatibility-center-dry-run-result-lookup-consumer-redaction", "open-compatibility-center-dry-run-result-lookup-route-authorization", "receipt consumption gate audit"}, "consume an accepted navigation receipt only after a separate consumer authorization gate exists"),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumptionGateItem("dismiss-receipt-dry-run-result-lookup-consumer-enablement-consumption-gate", "dismiss", "dismissal-consumer-enablement-receipt-consumption-gate", combined, []string{"dismiss-receipt-dry-run-result-lookup-consumer-enablement-acceptance-authorization", "dismiss-receipt-dry-run-result-lookup-consumer-redaction", "dismiss-receipt-dry-run-result-lookup-route-authorization", "receipt consumption gate audit"}, "consume an accepted dismissal receipt only after a separate consumer authorization gate exists"),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumptionGateItem("support-info-dry-run-result-lookup-consumer-enablement-consumption-gate", "support-info", "support-info-consumer-enablement-receipt-consumption-gate", combined, []string{"support-info-dry-run-result-lookup-consumer-enablement-acceptance-authorization", "support-info-dry-run-result-lookup-consumer-redaction", "support-info-dry-run-result-lookup-route-authorization", "receipt consumption gate audit"}, "consume an accepted support-info receipt only after a separate consumer authorization gate exists"),
	}
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumptionGateItem(id string, actionKind string, gateArea string, source string, tokens []string, nextRequirement string) ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumptionGateAuditItem {
	ready := productionAuthorizationHasAll(source, tokens)
	status := "missing-consumer-enablement-receipt-consumption-gate-evidence"
	if ready {
		status = "consumer-enablement-receipt-consumption-gate-modeled-consumption-disabled"
	}
	return ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumptionGateAuditItem{
		ID:                                    id,
		ActionKind:                            actionKind,
		GateArea:                              gateArea,
		RequiredEvidence:                      "future consumer enablement authorization receipt consumption requires accepted receipt, acceptance authorization, consumer redaction, and lookup route authorization evidence",
		EvidencePresent:                       ready,
		AcceptanceAuthorizationAuditConsumed:  ready,
		ConsumerRedactionAuditConsumed:        ready,
		LookupRouteAuthorizationAuditConsumed: ready,
		ConsumptionGateModeled:                ready,
		ReceiptConsumptionGateReady:           ready,
		ReviewOnly:                            true,
		RuntimeOwned:                          true,
		GoRuntimeBacked:                       true,
		SideEffectsDisabled:                   true,
		UserVisible:                           ready,
		GateStatus:                            status,
		NextRequirement:                       nextRequirement,
	}
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumptionGateChecks(preview ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumptionGateAuditPreview) []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumptionGateAuditCheck {
	return []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumptionGateAuditCheck{
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumptionGateCheck("acceptance-authorization-audit-consumed", productionAuthorizationPassBlocked(preview.AcceptanceAuthorizationAuditConsumed), "The consumption gate consumes receipt acceptance authorization evidence."),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumptionGateCheck("consumer-redaction-audit-consumed", productionAuthorizationPassBlocked(preview.ConsumerRedactionAuditConsumed), "The consumption gate consumes KDE and Runtime consumer redaction evidence."),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumptionGateCheck("lookup-route-authorization-audit-consumed", productionAuthorizationPassBlocked(preview.LookupRouteAuthorizationAuditConsumed), "The consumption gate consumes lookup route authorization evidence."),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumptionGateCheck("consumption-gate-guidance-consumed", productionAuthorizationPassBlocked(preview.ConsumptionGateGuidanceConsumed), "The consumption gate consumes current dispatch guidance."),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumptionGateCheck("consumption-gate-modeled-only", productionAuthorizationPassBlocked(preview.ConsumptionGateRequired && preview.ConsumptionGateModeled && preview.ConsumptionGateBoundaryReady && preview.ConsumerEnablementReceiptConsumptionReady && preview.ReceiptConsumptionGateReady && !preview.ReceiptAccepted && !preview.ReceiptConsumed && !preview.ConsumerConsumptionAuthorized && !preview.ConsumerEnablementAuthorized), "The receipt consumption gate is modeled without consuming a receipt or authorizing consumers."),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumptionGateCheck("five-consumption-gate-items-present", productionAuthorizationPassBlocked(preview.GateItemCount == 5 && preview.RequiredGateItemCount == 5 && preview.MissingGateItemCount == 0), "Review, renew, open Compatibility Center, dismiss, and support-info consumption gate items are modeled."),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumptionGateCheck("consumption-gate-items-ready-consumption-disabled", productionAuthorizationPassBlocked(preview.ReadyGateItemCount == 5 && productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumptionGateItemsReady(preview.GateItems)), "Every consumption gate item is ready while receipt consumption remains disabled."),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumptionGateCheck("receipt-acceptance-and-consumption-disabled", productionAuthorizationPassBlocked(!preview.ReceiptPresent && !preview.ReceiptPersisted && !preview.ReceiptAccepted && !preview.ReceiptConsumed && !preview.AuthorizationAccepted && !preview.ReceiptWriterEnabled && !preview.ReceiptPersistenceEnabled && !preview.ReceiptLookupWritesEnabled && !preview.ReceiptReplayEnabled && !preview.ReceiptExpiryWriteEnabled && !preview.ReceiptRevocationWriteEnabled && preview.AcceptedReceiptItemCount == 0 && preview.ConsumedReceiptItemCount == 0), "Receipt presence, acceptance, consumption, persistence, lookup writes, replay, expiry, and revocation remain disabled."),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumptionGateCheck("consumer-lookup-request-and-support-disabled", productionAuthorizationPassBlocked(!preview.ConsumerConsumptionAuthorized && !preview.ConsumerEnablementAuthorized && !preview.KDEConsumerEnabled && !preview.RuntimeConsumerEnabled && !preview.LookupRouteAuthorized && !preview.LookupRouteEnabled && !preview.LookupRoutePersisted && !preview.OpaqueLookupEnabled && !preview.OpaqueLookupPersisted && !preview.RedactedSummaryPersisted && !preview.DryRunResultPersisted && !preview.DispatchDryRunExecuted && !preview.RequestObjectCreationEnabled && !preview.RequestObjectDispatchEnabled && !preview.RequestObjectsCreated && !preview.RequestObjectsDispatched && !preview.PortalRequestCreated && !preview.NotificationActionEnabled && !preview.CompatibilityCenterOpened && !preview.SupportBundleExported && !preview.SupportCaseCreated && preview.AuthorizedConsumerItemCount == 0 && preview.EnabledConsumerItemCount == 0 && preview.LookupEnabledGateItemCount == 0), "Consumer authorization, lookup enablement, request creation, dispatch, notifications, navigation, and support writes remain disabled."),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumptionGateCheck("production-and-host-boundary-closed", productionAuthorizationPassBlocked(!preview.SystemServiceStarted && !preview.SessionBusClaimed && !preview.ProductionBusClaimed && !preview.ProductionOwnerEnabled && !preview.ProductionActivationReady && !preview.ProductionReadiness && !preview.ProductionOwnershipReady && !preview.WriteMethodsEnabled && !preview.RuntimeWritesEnabled && !preview.AdapterInvocationEnabled && !preview.BackendLaunchEnabled && !preview.BackendProcessStarted && !preview.NetworkRequired && !preview.HostRootModified && !preview.PrivilegedContainerRequired && !preview.StateRootPathExposed && !preview.FilePathsExposed && !preview.FileContentRead && !preview.RawResultExposed && !preview.RawCommandExposed && !preview.RawExecutableExposed && !preview.BackendDetailsExposed && preview.RawExposedGateItemCount == 0 && preview.SideEffectGateItemCount == 0), "Production ownership, writes, engine launch, unsafe data exposure, network, privilege, internal details, and host mutation remain disabled."),
	}
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumptionGateCheck(id string, status string, summary string) ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumptionGateAuditCheck {
	return ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumptionGateAuditCheck{ID: id, Status: status, Summary: summary}
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumptionGateAcceptanceReady(source string) bool {
	return productionAuthorizationHasAll(source, []string{"production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-receipt-acceptance-authorization-audit-preview", "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-receipt-acceptance-authorization-audit-ready-acceptance-disabled", "ReceiptAcceptanceAuthorizationReady", "ReceiptAccepted"})
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumptionGateRedactionReady(source string) bool {
	return productionAuthorizationHasAll(source, []string{"production-receipt-notification-action-dry-run-result-lookup-consumer-redaction-audit-preview", "production-receipt-notification-action-dry-run-result-lookup-consumer-redaction-audit-ready-consumers-disabled", "ConsumerRedactionReady", "KDEConsumerRedactionModeled", "RuntimeConsumerRedactionModeled"})
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumptionGateRouteReady(source string) bool {
	return productionAuthorizationHasAll(source, []string{"production-receipt-notification-action-dry-run-result-lookup-route-authorization-audit-preview", "production-receipt-notification-action-dry-run-result-lookup-route-authorization-audit-ready-route-disabled", "LookupRouteAuthorizationReady", "OwnerLocalLookupRouteModeled"})
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumptionGateGuidanceReady(source string) bool {
	return productionAuthorizationHasAll(source, []string{"receipt consumption gate audit", "accepted consumer enablement authorization receipt", "authorize KDE or Runtime consumers", "side effects"})
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumptionGateReadyCount(items []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumptionGateAuditItem) int {
	count := 0
	for _, item := range items {
		if item.EvidencePresent {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumptionGateMissingCount(items []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumptionGateAuditItem) int {
	count := 0
	for _, item := range items {
		if !item.EvidencePresent {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumptionGateAcceptedCount(items []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumptionGateAuditItem) int {
	count := 0
	for _, item := range items {
		if item.ReceiptAccepted || item.AuthorizationAccepted {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumptionGateConsumedCount(items []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumptionGateAuditItem) int {
	count := 0
	for _, item := range items {
		if item.ReceiptConsumed {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumptionGateAuthorizedCount(items []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumptionGateAuditItem) int {
	count := 0
	for _, item := range items {
		if item.ConsumerConsumptionAuthorized || item.ConsumerEnablementAuthorized {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumptionGateEnabledCount(items []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumptionGateAuditItem) int {
	count := 0
	for _, item := range items {
		if item.KDEConsumerEnabled || item.RuntimeConsumerEnabled {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumptionGateLookupEnabledCount(items []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumptionGateAuditItem) int {
	count := 0
	for _, item := range items {
		if item.LookupRouteEnabled || item.OpaqueLookupEnabled {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumptionGateRawExposedCount(items []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumptionGateAuditItem) int {
	count := 0
	for _, item := range items {
		if item.RawResultExposed || item.StateRootPathExposed || item.FilePathsExposed || item.FileContentRead || item.InternalDetailsExposed {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumptionGateSideEffectCount(items []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumptionGateAuditItem) int {
	count := 0
	for _, item := range items {
		if !item.SideEffectsDisabled || item.HostRootModified || item.RequestObjectCreated || item.RequestObjectDispatched || item.PortalRequestCreated || item.NotificationActionEnabled || item.CompatibilityCenterOpened || item.SupportBundleExported || item.SupportCaseCreated {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumptionGateItemIDs(items []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumptionGateAuditItem) []string {
	ids := make([]string, 0, len(items))
	for _, item := range items {
		ids = append(ids, item.ID)
	}
	return ids
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumptionGateItemsReady(items []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumptionGateAuditItem) bool {
	if len(items) != 5 {
		return false
	}
	for _, item := range items {
		if !item.EvidencePresent || !item.AcceptanceAuthorizationAuditConsumed || !item.ConsumerRedactionAuditConsumed || !item.LookupRouteAuthorizationAuditConsumed || !item.ConsumptionGateModeled || !item.ReceiptConsumptionGateReady || item.GateStatus != "consumer-enablement-receipt-consumption-gate-modeled-consumption-disabled" {
			return false
		}
		if item.ReceiptPresent || item.ReceiptPersisted || item.ReceiptAccepted || item.ReceiptConsumed || item.AuthorizationAccepted || item.ReceiptWriterEnabled || item.ReceiptPersistenceEnabled || item.ReceiptLookupWritesEnabled || item.ReceiptReplayEnabled || item.ReceiptExpiryWriteEnabled || item.ReceiptRevocationWriteEnabled || item.ConsumerConsumptionAuthorized || item.ConsumerEnablementAuthorized || item.KDEConsumerEnabled || item.RuntimeConsumerEnabled || item.LookupRouteEnabled || item.OpaqueLookupEnabled || item.RawResultExposed || !item.ReviewOnly || !item.SideEffectsDisabled || item.HostRootModified || item.InternalDetailsExposed {
			return false
		}
	}
	return true
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumptionGateCheckIDs(checks []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumptionGateAuditCheck) []string {
	ids := make([]string, 0, len(checks))
	for _, check := range checks {
		ids = append(ids, check.ID)
	}
	return ids
}

func countProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumptionGateChecks(checks []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumptionGateAuditCheck) ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumptionGateAuditCounts {
	counts := ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumptionGateAuditCounts{Total: len(checks)}
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
