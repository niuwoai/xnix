package owner

type ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptAcceptanceAuthorizationAuditPreview struct {
	Version                                    string                                                                                                            `json:"version"`
	SchemaVersion                              string                                                                                                            `json:"schema_version"`
	RequestType                                string                                                                                                            `json:"request_type"`
	AuditType                                  string                                                                                                            `json:"audit_type"`
	Source                                     string                                                                                                            `json:"source"`
	AuditDecision                              string                                                                                                            `json:"audit_decision"`
	ReceiptSchema                              string                                                                                                            `json:"receipt_schema"`
	OpaqueReceiptID                            string                                                                                                            `json:"opaque_receipt_id"`
	AcceptanceAuthorizationAuditRequired       bool                                                                                                              `json:"acceptance_authorization_audit_required"`
	AcceptanceAuthorizationAuditModeled        bool                                                                                                              `json:"acceptance_authorization_audit_modeled"`
	PersistenceAuthorizationAuditConsumed      bool                                                                                                              `json:"persistence_authorization_audit_consumed"`
	BaseAcceptancePropagationPreflightConsumed bool                                                                                                              `json:"base_acceptance_propagation_preflight_consumed"`
	AcceptanceAuthorizationGuidanceConsumed    bool                                                                                                              `json:"acceptance_authorization_guidance_consumed"`
	AcceptanceAuthorizationBoundaryReady       bool                                                                                                              `json:"acceptance_authorization_boundary_ready"`
	ReceiptPersistenceAuthorizationReady       bool                                                                                                              `json:"receipt_persistence_authorization_ready"`
	ReceiptAcceptancePropagationReady          bool                                                                                                              `json:"receipt_acceptance_propagation_ready"`
	ConsumerEnablementReceiptAcceptanceReady   bool                                                                                                              `json:"consumer_enablement_receipt_acceptance_ready"`
	ReceiptAcceptanceAuthorizationReady        bool                                                                                                              `json:"receipt_acceptance_authorization_ready"`
	ReceiptPresent                             bool                                                                                                              `json:"receipt_present"`
	ReceiptPersisted                           bool                                                                                                              `json:"receipt_persisted"`
	ReceiptAccepted                            bool                                                                                                              `json:"receipt_accepted"`
	AuthorizationAccepted                      bool                                                                                                              `json:"authorization_accepted"`
	ReceiptWriterEnabled                       bool                                                                                                              `json:"receipt_writer_enabled"`
	ReceiptPersistenceEnabled                  bool                                                                                                              `json:"receipt_persistence_enabled"`
	ReceiptLookupWritesEnabled                 bool                                                                                                              `json:"receipt_lookup_writes_enabled"`
	ReceiptReplayEnabled                       bool                                                                                                              `json:"receipt_replay_enabled"`
	ReceiptExpiryWriteEnabled                  bool                                                                                                              `json:"receipt_expiry_write_enabled"`
	ReceiptRevocationWriteEnabled              bool                                                                                                              `json:"receipt_revocation_write_enabled"`
	ConsumerConsumptionAuthorized              bool                                                                                                              `json:"consumer_consumption_authorized"`
	ConsumerEnablementAuthorized               bool                                                                                                              `json:"consumer_enablement_authorized"`
	KDEConsumerEnabled                         bool                                                                                                              `json:"kde_consumer_enabled"`
	RuntimeConsumerEnabled                     bool                                                                                                              `json:"runtime_consumer_enabled"`
	LookupRouteAuthorized                      bool                                                                                                              `json:"lookup_route_authorized"`
	LookupRouteEnabled                         bool                                                                                                              `json:"lookup_route_enabled"`
	OpaqueLookupEnabled                        bool                                                                                                              `json:"opaque_lookup_enabled"`
	RedactedSummaryPersisted                   bool                                                                                                              `json:"redacted_summary_persisted"`
	RawResultExposed                           bool                                                                                                              `json:"raw_result_exposed"`
	RuntimeDiagnosticsPersisted                bool                                                                                                              `json:"runtime_diagnostics_persisted"`
	DryRunResultPersisted                      bool                                                                                                              `json:"dry_run_result_persisted"`
	DispatchDryRunExecuted                     bool                                                                                                              `json:"dispatch_dry_run_executed"`
	RequestObjectCreationEnabled               bool                                                                                                              `json:"request_object_creation_enabled"`
	RequestObjectDispatchEnabled               bool                                                                                                              `json:"request_object_dispatch_enabled"`
	PortalRequestCreated                       bool                                                                                                              `json:"portal_request_created"`
	NotificationActionEnabled                  bool                                                                                                              `json:"notification_action_enabled"`
	CompatibilityCenterOpened                  bool                                                                                                              `json:"compatibility_center_opened"`
	SupportBundleExported                      bool                                                                                                              `json:"support_bundle_exported"`
	SupportCaseCreated                         bool                                                                                                              `json:"support_case_created"`
	ProductionReadiness                        bool                                                                                                              `json:"production_readiness"`
	ProductionOwnershipReady                   bool                                                                                                              `json:"production_ownership_ready"`
	AuthorizationItemCount                     int                                                                                                               `json:"authorization_item_count"`
	RequiredAuthorizationItemCount             int                                                                                                               `json:"required_authorization_item_count"`
	ReadyAuthorizationItemCount                int                                                                                                               `json:"ready_authorization_item_count"`
	MissingAuthorizationItemCount              int                                                                                                               `json:"missing_authorization_item_count"`
	AcceptedAuthorizationItemCount             int                                                                                                               `json:"accepted_authorization_item_count"`
	PersistedAuthorizationItemCount            int                                                                                                               `json:"persisted_authorization_item_count"`
	ConsumerAuthorizedItemCount                int                                                                                                               `json:"consumer_authorized_item_count"`
	EnabledConsumerItemCount                   int                                                                                                               `json:"enabled_consumer_item_count"`
	LookupEnabledAuthorizationItemCount        int                                                                                                               `json:"lookup_enabled_authorization_item_count"`
	RawExposedAuthorizationItemCount           int                                                                                                               `json:"raw_exposed_authorization_item_count"`
	SideEffectAuthorizationItemCount           int                                                                                                               `json:"side_effect_authorization_item_count"`
	AuthorizationItems                         []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptAcceptanceAuthorizationAuditItem  `json:"authorization_items"`
	AuthorizationItemIDs                       []string                                                                                                          `json:"authorization_item_ids"`
	RequiredBeforeReceiptAcceptance            []string                                                                                                          `json:"required_before_receipt_acceptance"`
	Checks                                     []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptAcceptanceAuthorizationAuditCheck `json:"checks"`
	CheckIDs                                   []string                                                                                                          `json:"check_ids"`
	Counts                                     ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptAcceptanceAuthorizationAuditCounts  `json:"counts"`
	RuntimeOwned                               bool                                                                                                              `json:"runtime_owned"`
	GoRuntimeBacked                            bool                                                                                                              `json:"go_runtime_backed"`
	KDEPolicyOwner                             bool                                                                                                              `json:"kde_policy_owner"`
	OfficialDesktopOnly                        bool                                                                                                              `json:"official_desktop_only"`
	PlasmaForkRequired                         bool                                                                                                              `json:"plasma_fork_required"`
	PlasmaSourceModified                       bool                                                                                                              `json:"plasma_source_modified"`
	SystemServiceStarted                       bool                                                                                                              `json:"system_service_started"`
	SessionBusClaimed                          bool                                                                                                              `json:"session_bus_claimed"`
	ProductionBusClaimed                       bool                                                                                                              `json:"production_bus_claimed"`
	ProductionOwnerEnabled                     bool                                                                                                              `json:"production_owner_enabled"`
	ProductionActivationReady                  bool                                                                                                              `json:"production_activation_ready"`
	WriteMethodsEnabled                        bool                                                                                                              `json:"write_methods_enabled"`
	RuntimeWritesEnabled                       bool                                                                                                              `json:"runtime_writes_enabled"`
	RequestObjectsCreated                      bool                                                                                                              `json:"request_objects_created"`
	RequestObjectsDispatched                   bool                                                                                                              `json:"request_objects_dispatched"`
	NotificationSent                           bool                                                                                                              `json:"notification_sent"`
	NotificationDeliveryEnabled                bool                                                                                                              `json:"notification_delivery_enabled"`
	CompatibilityCenterPersisted               bool                                                                                                              `json:"compatibility_center_persisted"`
	DesktopFilesWritten                        bool                                                                                                              `json:"desktop_files_written"`
	MIMEAppsWritten                            bool                                                                                                              `json:"mimeapps_written"`
	ShellConfigurationWritten                  bool                                                                                                              `json:"shell_configuration_written"`
	SettingsPersisted                          bool                                                                                                              `json:"settings_persisted"`
	AdapterInvocationEnabled                   bool                                                                                                              `json:"adapter_invocation_enabled"`
	BackendLaunchEnabled                       bool                                                                                                              `json:"backend_launch_enabled"`
	BackendProcessStarted                      bool                                                                                                              `json:"backend_process_started"`
	SnapshotRestoreExecuted                    bool                                                                                                              `json:"snapshot_restore_executed"`
	StateCleanupExecuted                       bool                                                                                                              `json:"state_cleanup_executed"`
	NetworkRequired                            bool                                                                                                              `json:"network_required"`
	HostRootModified                           bool                                                                                                              `json:"host_root_modified"`
	PrivilegedContainerRequired                bool                                                                                                              `json:"privileged_container_required"`
	CallerStateRootRequired                    bool                                                                                                              `json:"caller_state_root_required"`
	StateRootPathExposed                       bool                                                                                                              `json:"state_root_path_exposed"`
	FilePathsExposed                           bool                                                                                                              `json:"file_paths_exposed"`
	FileContentRead                            bool                                                                                                              `json:"file_content_read"`
	RawCommandExposed                          bool                                                                                                              `json:"raw_command_exposed"`
	RawExecutableExposed                       bool                                                                                                              `json:"raw_executable_exposed"`
	BackendDetailsExposed                      bool                                                                                                              `json:"backend_details_exposed"`
	BlockedActions                             []string                                                                                                          `json:"blocked_actions"`
	NextRequirements                           []string                                                                                                          `json:"next_requirements"`
	DesktopSafeSummary                         string                                                                                                            `json:"desktop_safe_summary"`
}

type ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptAcceptanceAuthorizationAuditItem struct {
	ID                                     string `json:"id"`
	ActionKind                             string `json:"action_kind"`
	AuthorizationArea                      string `json:"authorization_area"`
	RequiredEvidence                       string `json:"required_evidence"`
	EvidencePresent                        bool   `json:"evidence_present"`
	PersistenceAuthorizationAuditConsumed  bool   `json:"persistence_authorization_audit_consumed"`
	AcceptancePropagationPreflightConsumed bool   `json:"acceptance_propagation_preflight_consumed"`
	AcceptanceAuthorizationModeled         bool   `json:"acceptance_authorization_modeled"`
	ReceiptAcceptanceAuthorizationReady    bool   `json:"receipt_acceptance_authorization_ready"`
	ReceiptPresent                         bool   `json:"receipt_present"`
	ReceiptPersisted                       bool   `json:"receipt_persisted"`
	ReceiptAccepted                        bool   `json:"receipt_accepted"`
	AuthorizationAccepted                  bool   `json:"authorization_accepted"`
	ReceiptWriterEnabled                   bool   `json:"receipt_writer_enabled"`
	ReceiptPersistenceEnabled              bool   `json:"receipt_persistence_enabled"`
	ReceiptLookupWritesEnabled             bool   `json:"receipt_lookup_writes_enabled"`
	ReceiptReplayEnabled                   bool   `json:"receipt_replay_enabled"`
	ReceiptExpiryWriteEnabled              bool   `json:"receipt_expiry_write_enabled"`
	ReceiptRevocationWriteEnabled          bool   `json:"receipt_revocation_write_enabled"`
	ConsumerConsumptionAuthorized          bool   `json:"consumer_consumption_authorized"`
	ConsumerEnablementAuthorized           bool   `json:"consumer_enablement_authorized"`
	KDEConsumerEnabled                     bool   `json:"kde_consumer_enabled"`
	RuntimeConsumerEnabled                 bool   `json:"runtime_consumer_enabled"`
	LookupRouteEnabled                     bool   `json:"lookup_route_enabled"`
	OpaqueLookupEnabled                    bool   `json:"opaque_lookup_enabled"`
	RawResultExposed                       bool   `json:"raw_result_exposed"`
	UserVisible                            bool   `json:"user_visible"`
	ReviewOnly                             bool   `json:"review_only"`
	RuntimeOwned                           bool   `json:"runtime_owned"`
	GoRuntimeBacked                        bool   `json:"go_runtime_backed"`
	KDEPolicyOwner                         bool   `json:"kde_policy_owner"`
	CallerStateRootRequired                bool   `json:"caller_state_root_required"`
	StateRootPathExposed                   bool   `json:"state_root_path_exposed"`
	FilePathsExposed                       bool   `json:"file_paths_exposed"`
	FileContentRead                        bool   `json:"file_content_read"`
	RequestObjectCreated                   bool   `json:"request_object_created"`
	RequestObjectDispatched                bool   `json:"request_object_dispatched"`
	PortalRequestCreated                   bool   `json:"portal_request_created"`
	NotificationActionEnabled              bool   `json:"notification_action_enabled"`
	CompatibilityCenterOpened              bool   `json:"compatibility_center_opened"`
	SupportBundleExported                  bool   `json:"support_bundle_exported"`
	SupportCaseCreated                     bool   `json:"support_case_created"`
	ProductionReadiness                    bool   `json:"production_readiness"`
	ProductionOwnershipReady               bool   `json:"production_ownership_ready"`
	SideEffectsDisabled                    bool   `json:"side_effects_disabled"`
	HostRootModified                       bool   `json:"host_root_modified"`
	InternalDetailsExposed                 bool   `json:"internal_details_exposed"`
	AuthorizationStatus                    string `json:"authorization_status"`
	NextRequirement                        string `json:"next_requirement"`
}

type ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptAcceptanceAuthorizationAuditCheck struct {
	ID      string `json:"id"`
	Status  string `json:"status"`
	Summary string `json:"summary"`
}

type ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptAcceptanceAuthorizationAuditCounts struct {
	Total   int `json:"total"`
	Passed  int `json:"passed"`
	Pending int `json:"pending"`
	Blocked int `json:"blocked"`
}

func NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptAcceptanceAuthorizationAuditPreview(root string) (ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptAcceptanceAuthorizationAuditPreview, error) {
	version, err := readProductionDBusGateReviewVersion(root)
	if err != nil {
		return ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptAcceptanceAuthorizationAuditPreview{}, err
	}
	sources := productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptAcceptanceAuthorizationAuditSources(root)
	items := productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptAcceptanceAuthorizationItems(sources)
	persistenceReady := productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptAcceptanceAuthorizationPersistenceReady(sources.PersistenceAuthorizationAudit)
	propagationReady := productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptAcceptanceAuthorizationPropagationReady(sources.BaseAcceptancePropagationPreflight)
	guidanceReady := productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptAcceptanceAuthorizationGuidanceReady(sources.DispatchSheet)
	boundaryReady := persistenceReady && propagationReady && guidanceReady && productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptAcceptanceAuthorizationItemsReady(items)
	preview := ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptAcceptanceAuthorizationAuditPreview{
		Version:                               version,
		SchemaVersion:                         "xnix.runtime.production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_receipt_acceptance_authorization_audit.v1",
		RequestType:                           "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-receipt-acceptance-authorization-audit-preview",
		AuditType:                             "receipt-notification-action-dry-run-result-lookup-consumer-enablement-receipt-acceptance-authorization-audit",
		Source:                                "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-receipt-persistence-authorization-audit-preview+production-receipt-acceptance-propagation-preflight-preview+claude-code-current-dispatch-picks",
		AuditDecision:                         "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-receipt-acceptance-authorization-audit-blocked",
		ReceiptSchema:                         "xnix.runtime.production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_authorization_receipt.v1",
		OpaqueReceiptID:                       ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementAuthorizationReceiptID,
		AcceptanceAuthorizationAuditRequired:  true,
		AcceptanceAuthorizationAuditModeled:   true,
		PersistenceAuthorizationAuditConsumed: persistenceReady,
		BaseAcceptancePropagationPreflightConsumed: propagationReady,
		AcceptanceAuthorizationGuidanceConsumed:    guidanceReady,
		AcceptanceAuthorizationBoundaryReady:       boundaryReady,
		ReceiptPersistenceAuthorizationReady:       persistenceReady,
		ReceiptAcceptancePropagationReady:          propagationReady,
		ConsumerEnablementReceiptAcceptanceReady:   boundaryReady,
		ReceiptAcceptanceAuthorizationReady:        boundaryReady,
		ReceiptPresent:                             false,
		ReceiptPersisted:                           false,
		ReceiptAccepted:                            false,
		AuthorizationAccepted:                      false,
		ReceiptWriterEnabled:                       false,
		ReceiptPersistenceEnabled:                  false,
		ReceiptLookupWritesEnabled:                 false,
		ReceiptReplayEnabled:                       false,
		ReceiptExpiryWriteEnabled:                  false,
		ReceiptRevocationWriteEnabled:              false,
		ConsumerConsumptionAuthorized:              false,
		ConsumerEnablementAuthorized:               false,
		KDEConsumerEnabled:                         false,
		RuntimeConsumerEnabled:                     false,
		LookupRouteAuthorized:                      false,
		LookupRouteEnabled:                         false,
		OpaqueLookupEnabled:                        false,
		RedactedSummaryPersisted:                   false,
		RawResultExposed:                           false,
		RuntimeDiagnosticsPersisted:                false,
		DryRunResultPersisted:                      false,
		DispatchDryRunExecuted:                     false,
		RequestObjectCreationEnabled:               false,
		RequestObjectDispatchEnabled:               false,
		PortalRequestCreated:                       false,
		NotificationActionEnabled:                  false,
		CompatibilityCenterOpened:                  false,
		SupportBundleExported:                      false,
		SupportCaseCreated:                         false,
		ProductionReadiness:                        false,
		ProductionOwnershipReady:                   false,
		AuthorizationItemCount:                     len(items),
		RequiredAuthorizationItemCount:             5,
		ReadyAuthorizationItemCount:                productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptAcceptanceAuthorizationReadyCount(items),
		MissingAuthorizationItemCount:              productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptAcceptanceAuthorizationMissingCount(items),
		AuthorizationItems:                         items,
		AuthorizationItemIDs:                       productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptAcceptanceAuthorizationItemIDs(items),
		RequiredBeforeReceiptAcceptance: []string{
			"consumer enablement receipt persistence authorization boundary remains ready",
			"base production receipt acceptance propagation preflight remains ready",
			"operator action for consumer enablement receipt acceptance is reviewed separately",
			"receipt acceptance remains disabled until an accepted authorization is consumed by a separate gate",
			"consumer enablement remains disabled until a separate consumer gate consumes an accepted receipt",
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
			"treat this audit as permission to accept or consume consumer enablement authorization receipts",
			"write, persist, accept, replay, revoke, expire, or look up consumer enablement authorization receipts",
			"authorize or enable KDE or Runtime consumers from this acceptance audit",
			"create request objects, dispatch actions, create Portal requests, send notifications, open Compatibility Center, export support bundles, or create support cases",
			"claim production ownership, start services, enable Runtime writes, launch compatibility engines, expose raw results or paths, or mutate host root",
		},
		NextRequirements: []string{
			"Implement a dedicated receipt acceptance authorization action before any consumer enablement authorization receipt can be accepted.",
			"Keep consumer authorization and consumer enablement disabled until an accepted receipt is consumed by a separate gate.",
			"Keep route enablement, lookup enablement, dry-run execution, result persistence, request creation, and dispatch disabled.",
			"Keep production ownership, Runtime writes, desktop side effects, engine launch, raw result exposure, and host mutation disabled.",
		},
		DesktopSafeSummary: "The production receipt notification action dry-run result lookup consumer enablement receipt acceptance authorization audit models the acceptance boundary required before a future consumer enablement authorization receipt can be consumed, but it accepts no receipt, enables no consumer or lookup route, persists no state, exposes no raw result data, launches no engine, and mutates no host state.",
	}
	preview.AcceptedAuthorizationItemCount = productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptAcceptanceAuthorizationAcceptedCount(items)
	preview.PersistedAuthorizationItemCount = productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptAcceptanceAuthorizationPersistedCount(items)
	preview.ConsumerAuthorizedItemCount = productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptAcceptanceAuthorizationConsumerAuthorizedCount(items)
	preview.EnabledConsumerItemCount = productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptAcceptanceAuthorizationEnabledConsumerCount(items)
	preview.LookupEnabledAuthorizationItemCount = productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptAcceptanceAuthorizationLookupEnabledCount(items)
	preview.RawExposedAuthorizationItemCount = productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptAcceptanceAuthorizationRawExposedCount(items)
	preview.SideEffectAuthorizationItemCount = productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptAcceptanceAuthorizationSideEffectCount(items)
	checks := productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptAcceptanceAuthorizationChecks(preview)
	preview.Checks = checks
	preview.CheckIDs = productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptAcceptanceAuthorizationCheckIDs(checks)
	preview.Counts = countProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptAcceptanceAuthorizationChecks(checks)
	if preview.Counts.Blocked == 0 {
		preview.AuditDecision = "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-receipt-acceptance-authorization-audit-ready-acceptance-disabled"
	}
	if err := validateNoBackendTerms(preview, "production receipt notification action dry-run result lookup consumer enablement receipt acceptance authorization audit preview"); err != nil {
		return ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptAcceptanceAuthorizationAuditPreview{}, err
	}
	return preview, nil
}

type productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptAcceptanceAuthorizationSourceSet struct {
	PersistenceAuthorizationAudit      string
	BaseAcceptancePropagationPreflight string
	DispatchSheet                      string
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptAcceptanceAuthorizationAuditSources(root string) productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptAcceptanceAuthorizationSourceSet {
	return productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptAcceptanceAuthorizationSourceSet{
		PersistenceAuthorizationAudit:      productionAuthorizationReadSources(root, []string{"internal/runtime/owner/production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_receipt_persistence_authorization_audit.go"}),
		BaseAcceptancePropagationPreflight: productionAuthorizationReadSources(root, []string{"internal/runtime/owner/production_receipt_acceptance_propagation_preflight.go"}),
		DispatchSheet:                      productionAuthorizationReadSources(root, []string{"docs/claude-code-current-dispatch-picks.md"}),
	}
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptAcceptanceAuthorizationItems(sources productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptAcceptanceAuthorizationSourceSet) []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptAcceptanceAuthorizationAuditItem {
	combined := sources.PersistenceAuthorizationAudit + sources.BaseAcceptancePropagationPreflight + sources.DispatchSheet
	return []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptAcceptanceAuthorizationAuditItem{
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptAcceptanceAuthorizationItem("review-receipt-dry-run-result-lookup-consumer-enablement-acceptance-authorization", "review", "review-consumer-enablement-receipt-acceptance-authorization", combined, []string{"review-receipt-dry-run-result-lookup-consumer-enablement-persistence-authorization", "production-receipt-acceptance-propagation-preflight-preview", "receipt acceptance authorization audit"}, "review receipt acceptance requires a dedicated acceptance authorization gate"),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptAcceptanceAuthorizationItem("renew-receipt-dry-run-result-lookup-consumer-enablement-acceptance-authorization", "renew", "renewal-consumer-enablement-receipt-acceptance-authorization", combined, []string{"renew-receipt-dry-run-result-lookup-consumer-enablement-persistence-authorization", "production-receipt-acceptance-propagation-preflight-preview", "receipt acceptance authorization audit"}, "renewal receipt acceptance requires a dedicated acceptance authorization gate"),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptAcceptanceAuthorizationItem("open-compatibility-center-dry-run-result-lookup-consumer-enablement-acceptance-authorization", "open-compatibility-center", "navigation-consumer-enablement-receipt-acceptance-authorization", combined, []string{"open-compatibility-center-dry-run-result-lookup-consumer-enablement-persistence-authorization", "production-receipt-acceptance-propagation-preflight-preview", "receipt acceptance authorization audit"}, "navigation receipt acceptance requires a dedicated acceptance authorization gate"),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptAcceptanceAuthorizationItem("dismiss-receipt-dry-run-result-lookup-consumer-enablement-acceptance-authorization", "dismiss", "dismissal-consumer-enablement-receipt-acceptance-authorization", combined, []string{"dismiss-receipt-dry-run-result-lookup-consumer-enablement-persistence-authorization", "production-receipt-acceptance-propagation-preflight-preview", "receipt acceptance authorization audit"}, "dismissal receipt acceptance requires a dedicated acceptance authorization gate"),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptAcceptanceAuthorizationItem("support-info-dry-run-result-lookup-consumer-enablement-acceptance-authorization", "support-info", "support-info-consumer-enablement-receipt-acceptance-authorization", combined, []string{"support-info-dry-run-result-lookup-consumer-enablement-persistence-authorization", "production-receipt-acceptance-propagation-preflight-preview", "receipt acceptance authorization audit"}, "support-info receipt acceptance requires a dedicated acceptance authorization gate"),
	}
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptAcceptanceAuthorizationItem(id string, actionKind string, authorizationArea string, source string, tokens []string, nextRequirement string) ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptAcceptanceAuthorizationAuditItem {
	ready := productionAuthorizationHasAll(source, tokens)
	status := "missing-consumer-enablement-receipt-acceptance-authorization-evidence"
	if ready {
		status = "consumer-enablement-receipt-acceptance-authorization-modeled-acceptance-disabled"
	}
	return ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptAcceptanceAuthorizationAuditItem{
		ID:                                     id,
		ActionKind:                             actionKind,
		AuthorizationArea:                      authorizationArea,
		RequiredEvidence:                       "future consumer enablement authorization receipt acceptance requires a separate acceptance authorization boundary",
		EvidencePresent:                        ready,
		PersistenceAuthorizationAuditConsumed:  ready,
		AcceptancePropagationPreflightConsumed: ready,
		AcceptanceAuthorizationModeled:         ready,
		ReceiptAcceptanceAuthorizationReady:    ready,
		ReceiptPresent:                         false,
		ReceiptPersisted:                       false,
		ReceiptAccepted:                        false,
		AuthorizationAccepted:                  false,
		ReceiptWriterEnabled:                   false,
		ReceiptPersistenceEnabled:              false,
		ReceiptLookupWritesEnabled:             false,
		ReceiptReplayEnabled:                   false,
		ReceiptExpiryWriteEnabled:              false,
		ReceiptRevocationWriteEnabled:          false,
		ConsumerConsumptionAuthorized:          false,
		ConsumerEnablementAuthorized:           false,
		KDEConsumerEnabled:                     false,
		RuntimeConsumerEnabled:                 false,
		LookupRouteEnabled:                     false,
		OpaqueLookupEnabled:                    false,
		RawResultExposed:                       false,
		UserVisible:                            ready,
		ReviewOnly:                             true,
		RuntimeOwned:                           true,
		GoRuntimeBacked:                        true,
		KDEPolicyOwner:                         false,
		CallerStateRootRequired:                false,
		StateRootPathExposed:                   false,
		FilePathsExposed:                       false,
		FileContentRead:                        false,
		RequestObjectCreated:                   false,
		RequestObjectDispatched:                false,
		PortalRequestCreated:                   false,
		NotificationActionEnabled:              false,
		CompatibilityCenterOpened:              false,
		SupportBundleExported:                  false,
		SupportCaseCreated:                     false,
		ProductionReadiness:                    false,
		ProductionOwnershipReady:               false,
		SideEffectsDisabled:                    true,
		HostRootModified:                       false,
		InternalDetailsExposed:                 false,
		AuthorizationStatus:                    status,
		NextRequirement:                        nextRequirement,
	}
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptAcceptanceAuthorizationChecks(preview ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptAcceptanceAuthorizationAuditPreview) []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptAcceptanceAuthorizationAuditCheck {
	return []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptAcceptanceAuthorizationAuditCheck{
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptAcceptanceAuthorizationCheck("persistence-authorization-audit-consumed", productionAuthorizationPassBlocked(preview.PersistenceAuthorizationAuditConsumed), "The acceptance authorization audit consumes the consumer enablement receipt persistence authorization boundary."),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptAcceptanceAuthorizationCheck("base-acceptance-propagation-preflight-consumed", productionAuthorizationPassBlocked(preview.BaseAcceptancePropagationPreflightConsumed), "The audit consumes the base production receipt acceptance propagation preflight."),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptAcceptanceAuthorizationCheck("acceptance-authorization-guidance-consumed", productionAuthorizationPassBlocked(preview.AcceptanceAuthorizationGuidanceConsumed), "The audit consumes current dispatch guidance for the next acceptance authorization boundary."),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptAcceptanceAuthorizationCheck("acceptance-authorization-boundary-modeled-only", productionAuthorizationPassBlocked(preview.AcceptanceAuthorizationAuditRequired && preview.AcceptanceAuthorizationAuditModeled && preview.AcceptanceAuthorizationBoundaryReady && preview.ConsumerEnablementReceiptAcceptanceReady && preview.ReceiptAcceptanceAuthorizationReady && !preview.ReceiptPresent && !preview.ReceiptPersisted && !preview.ReceiptAccepted && !preview.AuthorizationAccepted), "The acceptance authorization boundary is modeled without accepting or consuming a receipt."),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptAcceptanceAuthorizationCheck("five-acceptance-authorization-items-present", productionAuthorizationPassBlocked(preview.AuthorizationItemCount == 5 && preview.RequiredAuthorizationItemCount == 5 && preview.MissingAuthorizationItemCount == 0), "Review, renew, open Compatibility Center, dismiss, and support-info acceptance authorization items are modeled."),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptAcceptanceAuthorizationCheck("acceptance-authorization-items-ready-acceptance-disabled", productionAuthorizationPassBlocked(preview.ReadyAuthorizationItemCount == 5 && productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptAcceptanceAuthorizationItemsReady(preview.AuthorizationItems)), "Every acceptance authorization item is ready while acceptance remains disabled."),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptAcceptanceAuthorizationCheck("receipt-acceptance-and-consumption-disabled", productionAuthorizationPassBlocked(!preview.ReceiptAccepted && !preview.AuthorizationAccepted && !preview.ReceiptPersisted && !preview.ReceiptWriterEnabled && !preview.ReceiptPersistenceEnabled && !preview.ReceiptLookupWritesEnabled && !preview.ReceiptReplayEnabled && !preview.ReceiptExpiryWriteEnabled && !preview.ReceiptRevocationWriteEnabled && preview.AcceptedAuthorizationItemCount == 0 && preview.PersistedAuthorizationItemCount == 0), "Receipt acceptance, persistence, lookup writes, replay, expiry, revocation, and consumption remain disabled."),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptAcceptanceAuthorizationCheck("consumer-lookup-request-and-support-disabled", productionAuthorizationPassBlocked(!preview.ConsumerConsumptionAuthorized && !preview.ConsumerEnablementAuthorized && !preview.KDEConsumerEnabled && !preview.RuntimeConsumerEnabled && !preview.LookupRouteAuthorized && !preview.LookupRouteEnabled && !preview.OpaqueLookupEnabled && !preview.RedactedSummaryPersisted && !preview.DryRunResultPersisted && !preview.DispatchDryRunExecuted && !preview.RequestObjectCreationEnabled && !preview.RequestObjectDispatchEnabled && !preview.RequestObjectsCreated && !preview.RequestObjectsDispatched && !preview.PortalRequestCreated && !preview.NotificationActionEnabled && !preview.CompatibilityCenterOpened && !preview.SupportBundleExported && !preview.SupportCaseCreated && preview.ConsumerAuthorizedItemCount == 0 && preview.EnabledConsumerItemCount == 0 && preview.LookupEnabledAuthorizationItemCount == 0), "Consumer authorization, lookup enablement, request creation, dispatch, notifications, navigation, and support writes remain disabled."),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptAcceptanceAuthorizationCheck("production-and-host-boundary-closed", productionAuthorizationPassBlocked(!preview.SystemServiceStarted && !preview.SessionBusClaimed && !preview.ProductionBusClaimed && !preview.ProductionOwnerEnabled && !preview.ProductionActivationReady && !preview.ProductionReadiness && !preview.ProductionOwnershipReady && !preview.WriteMethodsEnabled && !preview.RuntimeWritesEnabled && !preview.AdapterInvocationEnabled && !preview.BackendLaunchEnabled && !preview.BackendProcessStarted && !preview.NetworkRequired && !preview.HostRootModified && !preview.PrivilegedContainerRequired && !preview.StateRootPathExposed && !preview.FilePathsExposed && !preview.FileContentRead && !preview.RawResultExposed && !preview.RawCommandExposed && !preview.RawExecutableExposed && !preview.BackendDetailsExposed && preview.RawExposedAuthorizationItemCount == 0 && preview.SideEffectAuthorizationItemCount == 0), "Production ownership, writes, engine launch, unsafe data exposure, network, privilege, internal details, and host mutation remain disabled."),
	}
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptAcceptanceAuthorizationCheck(id string, status string, summary string) ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptAcceptanceAuthorizationAuditCheck {
	return ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptAcceptanceAuthorizationAuditCheck{ID: id, Status: status, Summary: summary}
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptAcceptanceAuthorizationPersistenceReady(source string) bool {
	return productionAuthorizationHasAll(source, []string{
		"production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-receipt-persistence-authorization-audit-preview",
		"production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-receipt-persistence-authorization-audit-ready-persistence-disabled",
		"ReceiptPersistenceAuthorizationReady",
		"ReceiptAccepted",
	})
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptAcceptanceAuthorizationPropagationReady(source string) bool {
	return productionAuthorizationHasAll(source, []string{
		"production-receipt-acceptance-propagation-preflight-preview",
		"production-receipt-acceptance-propagation-ready-acceptance-disabled",
		"FutureAcceptanceModeled",
		"AcceptanceSimulationOnly",
	})
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptAcceptanceAuthorizationGuidanceReady(source string) bool {
	return productionAuthorizationHasAll(source, []string{"receipt acceptance authorization audit", "consumer enablement authorization receipt", "authorize a consumer", "side effects"})
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptAcceptanceAuthorizationReadyCount(items []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptAcceptanceAuthorizationAuditItem) int {
	count := 0
	for _, item := range items {
		if item.EvidencePresent {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptAcceptanceAuthorizationMissingCount(items []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptAcceptanceAuthorizationAuditItem) int {
	count := 0
	for _, item := range items {
		if !item.EvidencePresent {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptAcceptanceAuthorizationAcceptedCount(items []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptAcceptanceAuthorizationAuditItem) int {
	count := 0
	for _, item := range items {
		if item.ReceiptAccepted || item.AuthorizationAccepted {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptAcceptanceAuthorizationPersistedCount(items []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptAcceptanceAuthorizationAuditItem) int {
	count := 0
	for _, item := range items {
		if item.ReceiptPersisted || item.ReceiptPersistenceEnabled || item.ReceiptLookupWritesEnabled {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptAcceptanceAuthorizationConsumerAuthorizedCount(items []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptAcceptanceAuthorizationAuditItem) int {
	count := 0
	for _, item := range items {
		if item.ConsumerConsumptionAuthorized || item.ConsumerEnablementAuthorized {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptAcceptanceAuthorizationEnabledConsumerCount(items []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptAcceptanceAuthorizationAuditItem) int {
	count := 0
	for _, item := range items {
		if item.KDEConsumerEnabled || item.RuntimeConsumerEnabled {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptAcceptanceAuthorizationLookupEnabledCount(items []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptAcceptanceAuthorizationAuditItem) int {
	count := 0
	for _, item := range items {
		if item.LookupRouteEnabled || item.OpaqueLookupEnabled {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptAcceptanceAuthorizationRawExposedCount(items []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptAcceptanceAuthorizationAuditItem) int {
	count := 0
	for _, item := range items {
		if item.RawResultExposed || item.StateRootPathExposed || item.FilePathsExposed || item.FileContentRead || item.InternalDetailsExposed {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptAcceptanceAuthorizationSideEffectCount(items []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptAcceptanceAuthorizationAuditItem) int {
	count := 0
	for _, item := range items {
		if !item.SideEffectsDisabled || item.HostRootModified || item.RequestObjectCreated || item.RequestObjectDispatched || item.PortalRequestCreated || item.NotificationActionEnabled || item.CompatibilityCenterOpened || item.SupportBundleExported || item.SupportCaseCreated {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptAcceptanceAuthorizationItemIDs(items []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptAcceptanceAuthorizationAuditItem) []string {
	ids := make([]string, 0, len(items))
	for _, item := range items {
		ids = append(ids, item.ID)
	}
	return ids
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptAcceptanceAuthorizationItemsReady(items []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptAcceptanceAuthorizationAuditItem) bool {
	if len(items) != 5 {
		return false
	}
	for _, item := range items {
		if !item.EvidencePresent || !item.PersistenceAuthorizationAuditConsumed || !item.AcceptancePropagationPreflightConsumed || !item.AcceptanceAuthorizationModeled || !item.ReceiptAcceptanceAuthorizationReady || item.AuthorizationStatus != "consumer-enablement-receipt-acceptance-authorization-modeled-acceptance-disabled" {
			return false
		}
		if item.ReceiptPersisted || item.ReceiptWriterEnabled || item.ReceiptPersistenceEnabled || item.ReceiptLookupWritesEnabled || item.ReceiptReplayEnabled || item.ReceiptExpiryWriteEnabled || item.ReceiptRevocationWriteEnabled || item.ReceiptAccepted || item.AuthorizationAccepted || item.ConsumerConsumptionAuthorized || item.ConsumerEnablementAuthorized || item.KDEConsumerEnabled || item.RuntimeConsumerEnabled || item.LookupRouteEnabled || item.OpaqueLookupEnabled || item.RawResultExposed || item.ProductionReadiness || item.ProductionOwnershipReady || item.KDEPolicyOwner || !item.ReviewOnly || !item.SideEffectsDisabled || item.HostRootModified || item.InternalDetailsExposed {
			return false
		}
	}
	return true
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptAcceptanceAuthorizationCheckIDs(checks []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptAcceptanceAuthorizationAuditCheck) []string {
	ids := make([]string, 0, len(checks))
	for _, check := range checks {
		ids = append(ids, check.ID)
	}
	return ids
}

func countProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptAcceptanceAuthorizationChecks(checks []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptAcceptanceAuthorizationAuditCheck) ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptAcceptanceAuthorizationAuditCounts {
	counts := ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptAcceptanceAuthorizationAuditCounts{Total: len(checks)}
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
