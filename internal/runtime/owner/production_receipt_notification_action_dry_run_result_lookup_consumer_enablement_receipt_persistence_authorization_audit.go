package owner

type ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptPersistenceAuthorizationAuditPreview struct {
	Version                                      string                                                                                                             `json:"version"`
	SchemaVersion                                string                                                                                                             `json:"schema_version"`
	RequestType                                  string                                                                                                             `json:"request_type"`
	AuditType                                    string                                                                                                             `json:"audit_type"`
	Source                                       string                                                                                                             `json:"source"`
	AuditDecision                                string                                                                                                             `json:"audit_decision"`
	ReceiptSchema                                string                                                                                                             `json:"receipt_schema"`
	OpaqueReceiptID                              string                                                                                                             `json:"opaque_receipt_id"`
	PersistenceAuthorizationAuditRequired        bool                                                                                                               `json:"persistence_authorization_audit_required"`
	PersistenceAuthorizationAuditModeled         bool                                                                                                               `json:"persistence_authorization_audit_modeled"`
	WriterAuthorizationAuditConsumed             bool                                                                                                               `json:"writer_authorization_audit_consumed"`
	BasePersistenceThreatReviewConsumed          bool                                                                                                               `json:"base_persistence_threat_review_consumed"`
	PersistenceAuthorizationGuidanceConsumed     bool                                                                                                               `json:"persistence_authorization_guidance_consumed"`
	PersistenceAuthorizationBoundaryReady        bool                                                                                                               `json:"persistence_authorization_boundary_ready"`
	ReceiptWriterAuthorizationReady              bool                                                                                                               `json:"receipt_writer_authorization_ready"`
	ReceiptPersistenceThreatBoundaryReady        bool                                                                                                               `json:"receipt_persistence_threat_boundary_ready"`
	ConsumerEnablementReceiptPersistenceReady    bool                                                                                                               `json:"consumer_enablement_receipt_persistence_ready"`
	ReceiptPersistenceAuthorizationReady         bool                                                                                                               `json:"receipt_persistence_authorization_ready"`
	ReceiptPresent                               bool                                                                                                               `json:"receipt_present"`
	ReceiptAccepted                              bool                                                                                                               `json:"receipt_accepted"`
	AuthorizationAccepted                        bool                                                                                                               `json:"authorization_accepted"`
	ReceiptWriterEnabled                         bool                                                                                                               `json:"receipt_writer_enabled"`
	ReceiptPersistenceEnabled                    bool                                                                                                               `json:"receipt_persistence_enabled"`
	ReceiptLookupWritesEnabled                   bool                                                                                                               `json:"receipt_lookup_writes_enabled"`
	ReceiptReplayEnabled                         bool                                                                                                               `json:"receipt_replay_enabled"`
	ReceiptExpiryWriteEnabled                    bool                                                                                                               `json:"receipt_expiry_write_enabled"`
	ReceiptRevocationWriteEnabled                bool                                                                                                               `json:"receipt_revocation_write_enabled"`
	ConsumerConsumptionAuthorized                bool                                                                                                               `json:"consumer_consumption_authorized"`
	ConsumerEnablementAuthorized                 bool                                                                                                               `json:"consumer_enablement_authorized"`
	KDEConsumerEnabled                           bool                                                                                                               `json:"kde_consumer_enabled"`
	RuntimeConsumerEnabled                       bool                                                                                                               `json:"runtime_consumer_enabled"`
	LookupRouteAuthorized                        bool                                                                                                               `json:"lookup_route_authorized"`
	LookupRouteEnabled                           bool                                                                                                               `json:"lookup_route_enabled"`
	OpaqueLookupEnabled                          bool                                                                                                               `json:"opaque_lookup_enabled"`
	RedactedSummaryPersisted                     bool                                                                                                               `json:"redacted_summary_persisted"`
	RawResultExposed                             bool                                                                                                               `json:"raw_result_exposed"`
	RuntimeDiagnosticsPersisted                  bool                                                                                                               `json:"runtime_diagnostics_persisted"`
	DryRunResultPersisted                        bool                                                                                                               `json:"dry_run_result_persisted"`
	DispatchDryRunExecuted                       bool                                                                                                               `json:"dispatch_dry_run_executed"`
	RequestObjectCreationEnabled                 bool                                                                                                               `json:"request_object_creation_enabled"`
	RequestObjectDispatchEnabled                 bool                                                                                                               `json:"request_object_dispatch_enabled"`
	PortalRequestCreated                         bool                                                                                                               `json:"portal_request_created"`
	NotificationActionEnabled                    bool                                                                                                               `json:"notification_action_enabled"`
	CompatibilityCenterOpened                    bool                                                                                                               `json:"compatibility_center_opened"`
	SupportBundleExported                        bool                                                                                                               `json:"support_bundle_exported"`
	SupportCaseCreated                           bool                                                                                                               `json:"support_case_created"`
	ProductionReadiness                          bool                                                                                                               `json:"production_readiness"`
	ProductionOwnershipReady                     bool                                                                                                               `json:"production_ownership_ready"`
	AuthorizationItemCount                       int                                                                                                                `json:"authorization_item_count"`
	RequiredAuthorizationItemCount               int                                                                                                                `json:"required_authorization_item_count"`
	ReadyAuthorizationItemCount                  int                                                                                                                `json:"ready_authorization_item_count"`
	MissingAuthorizationItemCount                int                                                                                                                `json:"missing_authorization_item_count"`
	PersistenceEnabledAuthorizationItemCount     int                                                                                                                `json:"persistence_enabled_authorization_item_count"`
	ReplayEnabledAuthorizationItemCount          int                                                                                                                `json:"replay_enabled_authorization_item_count"`
	ExpiryWriteEnabledAuthorizationItemCount     int                                                                                                                `json:"expiry_write_enabled_authorization_item_count"`
	RevocationWriteEnabledAuthorizationItemCount int                                                                                                                `json:"revocation_write_enabled_authorization_item_count"`
	AcceptedAuthorizationItemCount               int                                                                                                                `json:"accepted_authorization_item_count"`
	EnabledConsumerItemCount                     int                                                                                                                `json:"enabled_consumer_item_count"`
	RawExposedAuthorizationItemCount             int                                                                                                                `json:"raw_exposed_authorization_item_count"`
	SideEffectAuthorizationItemCount             int                                                                                                                `json:"side_effect_authorization_item_count"`
	AuthorizationItems                           []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptPersistenceAuthorizationAuditItem  `json:"authorization_items"`
	AuthorizationItemIDs                         []string                                                                                                           `json:"authorization_item_ids"`
	RequiredBeforeReceiptPersistence             []string                                                                                                           `json:"required_before_receipt_persistence"`
	Checks                                       []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptPersistenceAuthorizationAuditCheck `json:"checks"`
	CheckIDs                                     []string                                                                                                           `json:"check_ids"`
	Counts                                       ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptPersistenceAuthorizationAuditCounts  `json:"counts"`
	RuntimeOwned                                 bool                                                                                                               `json:"runtime_owned"`
	GoRuntimeBacked                              bool                                                                                                               `json:"go_runtime_backed"`
	KDEPolicyOwner                               bool                                                                                                               `json:"kde_policy_owner"`
	OfficialDesktopOnly                          bool                                                                                                               `json:"official_desktop_only"`
	PlasmaForkRequired                           bool                                                                                                               `json:"plasma_fork_required"`
	PlasmaSourceModified                         bool                                                                                                               `json:"plasma_source_modified"`
	SystemServiceStarted                         bool                                                                                                               `json:"system_service_started"`
	SessionBusClaimed                            bool                                                                                                               `json:"session_bus_claimed"`
	ProductionBusClaimed                         bool                                                                                                               `json:"production_bus_claimed"`
	ProductionOwnerEnabled                       bool                                                                                                               `json:"production_owner_enabled"`
	ProductionActivationReady                    bool                                                                                                               `json:"production_activation_ready"`
	WriteMethodsEnabled                          bool                                                                                                               `json:"write_methods_enabled"`
	RuntimeWritesEnabled                         bool                                                                                                               `json:"runtime_writes_enabled"`
	RequestObjectsCreated                        bool                                                                                                               `json:"request_objects_created"`
	RequestObjectsDispatched                     bool                                                                                                               `json:"request_objects_dispatched"`
	NotificationSent                             bool                                                                                                               `json:"notification_sent"`
	NotificationDeliveryEnabled                  bool                                                                                                               `json:"notification_delivery_enabled"`
	CompatibilityCenterPersisted                 bool                                                                                                               `json:"compatibility_center_persisted"`
	DesktopFilesWritten                          bool                                                                                                               `json:"desktop_files_written"`
	MIMEAppsWritten                              bool                                                                                                               `json:"mimeapps_written"`
	ShellConfigurationWritten                    bool                                                                                                               `json:"shell_configuration_written"`
	SettingsPersisted                            bool                                                                                                               `json:"settings_persisted"`
	AdapterInvocationEnabled                     bool                                                                                                               `json:"adapter_invocation_enabled"`
	BackendLaunchEnabled                         bool                                                                                                               `json:"backend_launch_enabled"`
	BackendProcessStarted                        bool                                                                                                               `json:"backend_process_started"`
	SnapshotRestoreExecuted                      bool                                                                                                               `json:"snapshot_restore_executed"`
	StateCleanupExecuted                         bool                                                                                                               `json:"state_cleanup_executed"`
	NetworkRequired                              bool                                                                                                               `json:"network_required"`
	HostRootModified                             bool                                                                                                               `json:"host_root_modified"`
	PrivilegedContainerRequired                  bool                                                                                                               `json:"privileged_container_required"`
	CallerStateRootRequired                      bool                                                                                                               `json:"caller_state_root_required"`
	StateRootPathExposed                         bool                                                                                                               `json:"state_root_path_exposed"`
	FilePathsExposed                             bool                                                                                                               `json:"file_paths_exposed"`
	FileContentRead                              bool                                                                                                               `json:"file_content_read"`
	RawCommandExposed                            bool                                                                                                               `json:"raw_command_exposed"`
	RawExecutableExposed                         bool                                                                                                               `json:"raw_executable_exposed"`
	BackendDetailsExposed                        bool                                                                                                               `json:"backend_details_exposed"`
	BlockedActions                               []string                                                                                                           `json:"blocked_actions"`
	NextRequirements                             []string                                                                                                           `json:"next_requirements"`
	DesktopSafeSummary                           string                                                                                                             `json:"desktop_safe_summary"`
}

type ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptPersistenceAuthorizationAuditItem struct {
	ID                                   string `json:"id"`
	ActionKind                           string `json:"action_kind"`
	AuthorizationArea                    string `json:"authorization_area"`
	RequiredEvidence                     string `json:"required_evidence"`
	EvidencePresent                      bool   `json:"evidence_present"`
	WriterAuthorizationAuditConsumed     bool   `json:"writer_authorization_audit_consumed"`
	PersistenceThreatReviewConsumed      bool   `json:"persistence_threat_review_consumed"`
	PersistenceAuthorizationModeled      bool   `json:"persistence_authorization_modeled"`
	ReceiptPersistenceAuthorizationReady bool   `json:"receipt_persistence_authorization_ready"`
	ReceiptPresent                       bool   `json:"receipt_present"`
	ReceiptAccepted                      bool   `json:"receipt_accepted"`
	AuthorizationAccepted                bool   `json:"authorization_accepted"`
	ReceiptWriterEnabled                 bool   `json:"receipt_writer_enabled"`
	ReceiptPersistenceEnabled            bool   `json:"receipt_persistence_enabled"`
	ReceiptLookupWritesEnabled           bool   `json:"receipt_lookup_writes_enabled"`
	ReceiptReplayEnabled                 bool   `json:"receipt_replay_enabled"`
	ReceiptExpiryWriteEnabled            bool   `json:"receipt_expiry_write_enabled"`
	ReceiptRevocationWriteEnabled        bool   `json:"receipt_revocation_write_enabled"`
	ConsumerConsumptionAuthorized        bool   `json:"consumer_consumption_authorized"`
	ConsumerEnablementAuthorized         bool   `json:"consumer_enablement_authorized"`
	KDEConsumerEnabled                   bool   `json:"kde_consumer_enabled"`
	RuntimeConsumerEnabled               bool   `json:"runtime_consumer_enabled"`
	LookupRouteEnabled                   bool   `json:"lookup_route_enabled"`
	OpaqueLookupEnabled                  bool   `json:"opaque_lookup_enabled"`
	RawResultExposed                     bool   `json:"raw_result_exposed"`
	UserVisible                          bool   `json:"user_visible"`
	ReviewOnly                           bool   `json:"review_only"`
	RuntimeOwned                         bool   `json:"runtime_owned"`
	GoRuntimeBacked                      bool   `json:"go_runtime_backed"`
	KDEPolicyOwner                       bool   `json:"kde_policy_owner"`
	CallerStateRootRequired              bool   `json:"caller_state_root_required"`
	StateRootPathExposed                 bool   `json:"state_root_path_exposed"`
	FilePathsExposed                     bool   `json:"file_paths_exposed"`
	FileContentRead                      bool   `json:"file_content_read"`
	RequestObjectCreated                 bool   `json:"request_object_created"`
	RequestObjectDispatched              bool   `json:"request_object_dispatched"`
	PortalRequestCreated                 bool   `json:"portal_request_created"`
	NotificationActionEnabled            bool   `json:"notification_action_enabled"`
	CompatibilityCenterOpened            bool   `json:"compatibility_center_opened"`
	SupportBundleExported                bool   `json:"support_bundle_exported"`
	SupportCaseCreated                   bool   `json:"support_case_created"`
	ProductionReadiness                  bool   `json:"production_readiness"`
	ProductionOwnershipReady             bool   `json:"production_ownership_ready"`
	SideEffectsDisabled                  bool   `json:"side_effects_disabled"`
	HostRootModified                     bool   `json:"host_root_modified"`
	InternalDetailsExposed               bool   `json:"internal_details_exposed"`
	AuthorizationStatus                  string `json:"authorization_status"`
	NextRequirement                      string `json:"next_requirement"`
}

type ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptPersistenceAuthorizationAuditCheck struct {
	ID      string `json:"id"`
	Status  string `json:"status"`
	Summary string `json:"summary"`
}

type ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptPersistenceAuthorizationAuditCounts struct {
	Total   int `json:"total"`
	Passed  int `json:"passed"`
	Pending int `json:"pending"`
	Blocked int `json:"blocked"`
}

func NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptPersistenceAuthorizationAuditPreview(root string) (ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptPersistenceAuthorizationAuditPreview, error) {
	version, err := readProductionDBusGateReviewVersion(root)
	if err != nil {
		return ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptPersistenceAuthorizationAuditPreview{}, err
	}
	sources := productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptPersistenceAuthorizationAuditSources(root)
	items := productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptPersistenceAuthorizationItems(sources)
	writerReady := productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptPersistenceAuthorizationWriterReady(sources.WriterAuthorizationAudit)
	threatReady := productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptPersistenceAuthorizationThreatReady(sources.BasePersistenceThreatReview)
	guidanceReady := productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptPersistenceAuthorizationGuidanceReady(sources.DispatchSheet)
	boundaryReady := writerReady && threatReady && guidanceReady && productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptPersistenceAuthorizationItemsReady(items)
	preview := ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptPersistenceAuthorizationAuditPreview{
		Version:                                   version,
		SchemaVersion:                             "xnix.runtime.production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_receipt_persistence_authorization_audit.v1",
		RequestType:                               "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-receipt-persistence-authorization-audit-preview",
		AuditType:                                 "receipt-notification-action-dry-run-result-lookup-consumer-enablement-receipt-persistence-authorization-audit",
		Source:                                    "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-receipt-writer-authorization-audit-preview+production-receipt-persistence-threat-review-preview+claude-code-current-dispatch-picks",
		AuditDecision:                             "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-receipt-persistence-authorization-audit-blocked",
		ReceiptSchema:                             "xnix.runtime.production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_authorization_receipt.v1",
		OpaqueReceiptID:                           ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementAuthorizationReceiptID,
		PersistenceAuthorizationAuditRequired:     true,
		PersistenceAuthorizationAuditModeled:      true,
		WriterAuthorizationAuditConsumed:          writerReady,
		BasePersistenceThreatReviewConsumed:       threatReady,
		PersistenceAuthorizationGuidanceConsumed:  guidanceReady,
		PersistenceAuthorizationBoundaryReady:     boundaryReady,
		ReceiptWriterAuthorizationReady:           writerReady,
		ReceiptPersistenceThreatBoundaryReady:     threatReady,
		ConsumerEnablementReceiptPersistenceReady: boundaryReady,
		ReceiptPersistenceAuthorizationReady:      boundaryReady,
		ReceiptPresent:                            false,
		ReceiptAccepted:                           false,
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
		OpaqueLookupEnabled:                       false,
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
		AuthorizationItemCount:                    len(items),
		RequiredAuthorizationItemCount:            5,
		ReadyAuthorizationItemCount:               productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptPersistenceAuthorizationReadyCount(items),
		MissingAuthorizationItemCount:             productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptPersistenceAuthorizationMissingCount(items),
		AuthorizationItems:                        items,
		AuthorizationItemIDs:                      productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptPersistenceAuthorizationItemIDs(items),
		RequiredBeforeReceiptPersistence: []string{
			"consumer enablement receipt writer authorization boundary remains ready",
			"base production receipt persistence threat review remains ready",
			"operator action for consumer enablement receipt persistence is reviewed separately",
			"receipt storage, lookup writes, replay, expiry, and revocation remain disabled until accepted persistence authorization exists",
			"receipt acceptance remains disabled until a separate gate consumes an accepted persistence authorization receipt",
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
			"treat this audit as permission to persist or accept consumer enablement authorization receipts",
			"write, persist, accept, replay, revoke, expire, or look up consumer enablement authorization receipts",
			"authorize or enable KDE or Runtime consumers from this persistence audit",
			"create request objects, dispatch actions, create Portal requests, send notifications, open Compatibility Center, export support bundles, or create support cases",
			"claim production ownership, start services, enable Runtime writes, launch compatibility engines, expose raw results or paths, or mutate host root",
		},
		NextRequirements: []string{
			"Implement a dedicated receipt persistence authorization action before any consumer enablement authorization receipt can be persisted.",
			"Keep replay, expiry, revocation, and lookup writes disabled until a separate persistence implementation gate exists.",
			"Keep receipt acceptance disabled until an accepted persistence authorization is consumed by a separate gate.",
			"Keep consumer authorization and consumer enablement disabled until an accepted receipt is consumed by a separate gate.",
			"Keep production ownership, Runtime writes, desktop side effects, engine launch, raw result exposure, and host mutation disabled.",
		},
		DesktopSafeSummary: "The production receipt notification action dry-run result lookup consumer enablement receipt persistence authorization audit models the persistence boundary required before a future consumer enablement authorization receipt can be persisted or accepted, but it writes no receipt, persists no state, accepts no authorization, enables no consumer or lookup route, exposes no raw result data, launches no engine, and mutates no host state.",
	}
	preview.PersistenceEnabledAuthorizationItemCount = productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptPersistenceAuthorizationPersistenceEnabledCount(items)
	preview.ReplayEnabledAuthorizationItemCount = productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptPersistenceAuthorizationReplayEnabledCount(items)
	preview.ExpiryWriteEnabledAuthorizationItemCount = productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptPersistenceAuthorizationExpiryWriteEnabledCount(items)
	preview.RevocationWriteEnabledAuthorizationItemCount = productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptPersistenceAuthorizationRevocationWriteEnabledCount(items)
	preview.AcceptedAuthorizationItemCount = productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptPersistenceAuthorizationAcceptedCount(items)
	preview.EnabledConsumerItemCount = productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptPersistenceAuthorizationEnabledConsumerCount(items)
	preview.RawExposedAuthorizationItemCount = productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptPersistenceAuthorizationRawExposedCount(items)
	preview.SideEffectAuthorizationItemCount = productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptPersistenceAuthorizationSideEffectCount(items)
	checks := productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptPersistenceAuthorizationChecks(preview)
	preview.Checks = checks
	preview.CheckIDs = productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptPersistenceAuthorizationCheckIDs(checks)
	preview.Counts = countProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptPersistenceAuthorizationChecks(checks)
	if preview.Counts.Blocked == 0 {
		preview.AuditDecision = "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-receipt-persistence-authorization-audit-ready-persistence-disabled"
	}
	if err := validateNoBackendTerms(preview, "production receipt notification action dry-run result lookup consumer enablement receipt persistence authorization audit preview"); err != nil {
		return ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptPersistenceAuthorizationAuditPreview{}, err
	}
	return preview, nil
}

type productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptPersistenceAuthorizationSourceSet struct {
	WriterAuthorizationAudit    string
	BasePersistenceThreatReview string
	DispatchSheet               string
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptPersistenceAuthorizationAuditSources(root string) productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptPersistenceAuthorizationSourceSet {
	return productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptPersistenceAuthorizationSourceSet{
		WriterAuthorizationAudit:    productionAuthorizationReadSources(root, []string{"internal/runtime/owner/production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_receipt_writer_authorization_audit.go"}),
		BasePersistenceThreatReview: productionAuthorizationReadSources(root, []string{"internal/runtime/owner/production_receipt_persistence_threat_review.go"}),
		DispatchSheet:               productionAuthorizationReadSources(root, []string{"docs/claude-code-current-dispatch-picks.md"}),
	}
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptPersistenceAuthorizationItems(sources productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptPersistenceAuthorizationSourceSet) []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptPersistenceAuthorizationAuditItem {
	combined := sources.WriterAuthorizationAudit + sources.BasePersistenceThreatReview + sources.DispatchSheet
	return []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptPersistenceAuthorizationAuditItem{
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptPersistenceAuthorizationItem("review-receipt-dry-run-result-lookup-consumer-enablement-persistence-authorization", "review", "review-consumer-enablement-receipt-persistence-authorization", combined, []string{"review-receipt-dry-run-result-lookup-consumer-enablement-writer-authorization", "production-receipt-persistence-threat-review-preview", "receipt persistence authorization audit"}, "review receipt persistence requires a dedicated persistence authorization gate"),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptPersistenceAuthorizationItem("renew-receipt-dry-run-result-lookup-consumer-enablement-persistence-authorization", "renew", "renewal-consumer-enablement-receipt-persistence-authorization", combined, []string{"renew-receipt-dry-run-result-lookup-consumer-enablement-writer-authorization", "production-receipt-persistence-threat-review-preview", "receipt persistence authorization audit"}, "renewal receipt persistence requires a dedicated persistence authorization gate"),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptPersistenceAuthorizationItem("open-compatibility-center-dry-run-result-lookup-consumer-enablement-persistence-authorization", "open-compatibility-center", "navigation-consumer-enablement-receipt-persistence-authorization", combined, []string{"open-compatibility-center-dry-run-result-lookup-consumer-enablement-writer-authorization", "production-receipt-persistence-threat-review-preview", "receipt persistence authorization audit"}, "navigation receipt persistence requires a dedicated persistence authorization gate"),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptPersistenceAuthorizationItem("dismiss-receipt-dry-run-result-lookup-consumer-enablement-persistence-authorization", "dismiss", "dismissal-consumer-enablement-receipt-persistence-authorization", combined, []string{"dismiss-receipt-dry-run-result-lookup-consumer-enablement-writer-authorization", "production-receipt-persistence-threat-review-preview", "receipt persistence authorization audit"}, "dismissal receipt persistence requires a dedicated persistence authorization gate"),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptPersistenceAuthorizationItem("support-info-dry-run-result-lookup-consumer-enablement-persistence-authorization", "support-info", "support-info-consumer-enablement-receipt-persistence-authorization", combined, []string{"support-info-dry-run-result-lookup-consumer-enablement-writer-authorization", "production-receipt-persistence-threat-review-preview", "receipt persistence authorization audit"}, "support-info receipt persistence requires a dedicated persistence authorization gate"),
	}
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptPersistenceAuthorizationItem(id string, actionKind string, authorizationArea string, source string, tokens []string, nextRequirement string) ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptPersistenceAuthorizationAuditItem {
	ready := productionAuthorizationHasAll(source, tokens)
	status := "missing-consumer-enablement-receipt-persistence-authorization-evidence"
	if ready {
		status = "consumer-enablement-receipt-persistence-authorization-modeled-persistence-disabled"
	}
	return ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptPersistenceAuthorizationAuditItem{
		ID:                                   id,
		ActionKind:                           actionKind,
		AuthorizationArea:                    authorizationArea,
		RequiredEvidence:                     "future consumer enablement authorization receipt persistence requires a separate persistence authorization boundary",
		EvidencePresent:                      ready,
		WriterAuthorizationAuditConsumed:     ready,
		PersistenceThreatReviewConsumed:      ready,
		PersistenceAuthorizationModeled:      ready,
		ReceiptPersistenceAuthorizationReady: ready,
		ReceiptPresent:                       false,
		ReceiptAccepted:                      false,
		AuthorizationAccepted:                false,
		ReceiptWriterEnabled:                 false,
		ReceiptPersistenceEnabled:            false,
		ReceiptLookupWritesEnabled:           false,
		ReceiptReplayEnabled:                 false,
		ReceiptExpiryWriteEnabled:            false,
		ReceiptRevocationWriteEnabled:        false,
		ConsumerConsumptionAuthorized:        false,
		ConsumerEnablementAuthorized:         false,
		KDEConsumerEnabled:                   false,
		RuntimeConsumerEnabled:               false,
		LookupRouteEnabled:                   false,
		OpaqueLookupEnabled:                  false,
		RawResultExposed:                     false,
		UserVisible:                          ready,
		ReviewOnly:                           true,
		RuntimeOwned:                         true,
		GoRuntimeBacked:                      true,
		KDEPolicyOwner:                       false,
		CallerStateRootRequired:              false,
		StateRootPathExposed:                 false,
		FilePathsExposed:                     false,
		FileContentRead:                      false,
		RequestObjectCreated:                 false,
		RequestObjectDispatched:              false,
		PortalRequestCreated:                 false,
		NotificationActionEnabled:            false,
		CompatibilityCenterOpened:            false,
		SupportBundleExported:                false,
		SupportCaseCreated:                   false,
		ProductionReadiness:                  false,
		ProductionOwnershipReady:             false,
		SideEffectsDisabled:                  true,
		HostRootModified:                     false,
		InternalDetailsExposed:               false,
		AuthorizationStatus:                  status,
		NextRequirement:                      nextRequirement,
	}
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptPersistenceAuthorizationChecks(preview ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptPersistenceAuthorizationAuditPreview) []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptPersistenceAuthorizationAuditCheck {
	return []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptPersistenceAuthorizationAuditCheck{
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptPersistenceAuthorizationCheck("writer-authorization-audit-consumed", productionAuthorizationPassBlocked(preview.WriterAuthorizationAuditConsumed), "The persistence authorization audit consumes the consumer enablement receipt writer authorization boundary."),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptPersistenceAuthorizationCheck("base-persistence-threat-review-consumed", productionAuthorizationPassBlocked(preview.BasePersistenceThreatReviewConsumed), "The audit consumes the base production receipt persistence threat review."),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptPersistenceAuthorizationCheck("persistence-authorization-guidance-consumed", productionAuthorizationPassBlocked(preview.PersistenceAuthorizationGuidanceConsumed), "The audit consumes current dispatch guidance for the next persistence authorization boundary."),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptPersistenceAuthorizationCheck("persistence-authorization-boundary-modeled-only", productionAuthorizationPassBlocked(preview.PersistenceAuthorizationAuditRequired && preview.PersistenceAuthorizationAuditModeled && preview.PersistenceAuthorizationBoundaryReady && preview.ConsumerEnablementReceiptPersistenceReady && preview.ReceiptPersistenceAuthorizationReady && !preview.ReceiptPresent && !preview.ReceiptAccepted && !preview.AuthorizationAccepted && !preview.ReceiptPersistenceEnabled), "The persistence authorization boundary is modeled without persisting or accepting a receipt."),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptPersistenceAuthorizationCheck("five-persistence-authorization-items-present", productionAuthorizationPassBlocked(preview.AuthorizationItemCount == 5 && preview.RequiredAuthorizationItemCount == 5 && preview.MissingAuthorizationItemCount == 0), "Review, renew, open Compatibility Center, dismiss, and support-info persistence authorization items are modeled."),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptPersistenceAuthorizationCheck("persistence-authorization-items-ready-persistence-disabled", productionAuthorizationPassBlocked(preview.ReadyAuthorizationItemCount == 5 && productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptPersistenceAuthorizationItemsReady(preview.AuthorizationItems)), "Every persistence authorization item is ready while persistence remains disabled."),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptPersistenceAuthorizationCheck("receipt-persistence-replay-expiry-revocation-disabled", productionAuthorizationPassBlocked(!preview.ReceiptWriterEnabled && !preview.ReceiptPersistenceEnabled && !preview.ReceiptLookupWritesEnabled && !preview.ReceiptReplayEnabled && !preview.ReceiptExpiryWriteEnabled && !preview.ReceiptRevocationWriteEnabled && !preview.ReceiptAccepted && preview.PersistenceEnabledAuthorizationItemCount == 0 && preview.ReplayEnabledAuthorizationItemCount == 0 && preview.ExpiryWriteEnabledAuthorizationItemCount == 0 && preview.RevocationWriteEnabledAuthorizationItemCount == 0 && preview.AcceptedAuthorizationItemCount == 0), "Receipt writing, persistence, lookup writes, replay, expiry, revocation, and acceptance remain disabled."),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptPersistenceAuthorizationCheck("consumer-lookup-request-and-support-disabled", productionAuthorizationPassBlocked(!preview.ConsumerConsumptionAuthorized && !preview.ConsumerEnablementAuthorized && !preview.KDEConsumerEnabled && !preview.RuntimeConsumerEnabled && !preview.LookupRouteAuthorized && !preview.LookupRouteEnabled && !preview.OpaqueLookupEnabled && !preview.RedactedSummaryPersisted && !preview.DryRunResultPersisted && !preview.DispatchDryRunExecuted && !preview.RequestObjectCreationEnabled && !preview.RequestObjectDispatchEnabled && !preview.RequestObjectsCreated && !preview.RequestObjectsDispatched && !preview.PortalRequestCreated && !preview.NotificationActionEnabled && !preview.CompatibilityCenterOpened && !preview.SupportBundleExported && !preview.SupportCaseCreated && preview.EnabledConsumerItemCount == 0), "Consumer authorization, lookup enablement, request creation, dispatch, notifications, navigation, and support writes remain disabled."),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptPersistenceAuthorizationCheck("production-and-host-boundary-closed", productionAuthorizationPassBlocked(!preview.SystemServiceStarted && !preview.SessionBusClaimed && !preview.ProductionBusClaimed && !preview.ProductionOwnerEnabled && !preview.ProductionActivationReady && !preview.ProductionReadiness && !preview.ProductionOwnershipReady && !preview.WriteMethodsEnabled && !preview.RuntimeWritesEnabled && !preview.AdapterInvocationEnabled && !preview.BackendLaunchEnabled && !preview.BackendProcessStarted && !preview.NetworkRequired && !preview.HostRootModified && !preview.PrivilegedContainerRequired && !preview.StateRootPathExposed && !preview.FilePathsExposed && !preview.FileContentRead && !preview.RawResultExposed && !preview.RawCommandExposed && !preview.RawExecutableExposed && !preview.BackendDetailsExposed && preview.RawExposedAuthorizationItemCount == 0 && preview.SideEffectAuthorizationItemCount == 0), "Production ownership, writes, engine launch, unsafe data exposure, network, privilege, internal details, and host mutation remain disabled."),
	}
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptPersistenceAuthorizationCheck(id string, status string, summary string) ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptPersistenceAuthorizationAuditCheck {
	return ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptPersistenceAuthorizationAuditCheck{ID: id, Status: status, Summary: summary}
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptPersistenceAuthorizationWriterReady(source string) bool {
	return productionAuthorizationHasAll(source, []string{
		"production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-receipt-writer-authorization-audit-preview",
		"production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-receipt-writer-authorization-audit-ready-writes-disabled",
		"ReceiptWriterAuthorizationReady",
		"ReceiptPersistenceEnabled",
	})
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptPersistenceAuthorizationThreatReady(source string) bool {
	return productionAuthorizationHasAll(source, []string{
		"production-receipt-persistence-threat-review-preview",
		"production-receipt-persistence-threat-review-ready-persistence-disabled",
		"PersistenceThreatReviewReady",
		"ReceiptPersistenceEnabled",
		"ReceiptReplayEnabled",
	})
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptPersistenceAuthorizationGuidanceReady(source string) bool {
	return productionAuthorizationHasAll(source, []string{"receipt persistence authorization audit", "consumer enablement authorization receipt", "persisted or accepted", "side effects"})
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptPersistenceAuthorizationReadyCount(items []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptPersistenceAuthorizationAuditItem) int {
	count := 0
	for _, item := range items {
		if item.EvidencePresent {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptPersistenceAuthorizationMissingCount(items []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptPersistenceAuthorizationAuditItem) int {
	count := 0
	for _, item := range items {
		if !item.EvidencePresent {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptPersistenceAuthorizationPersistenceEnabledCount(items []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptPersistenceAuthorizationAuditItem) int {
	count := 0
	for _, item := range items {
		if item.ReceiptPersistenceEnabled || item.ReceiptLookupWritesEnabled {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptPersistenceAuthorizationReplayEnabledCount(items []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptPersistenceAuthorizationAuditItem) int {
	count := 0
	for _, item := range items {
		if item.ReceiptReplayEnabled {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptPersistenceAuthorizationExpiryWriteEnabledCount(items []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptPersistenceAuthorizationAuditItem) int {
	count := 0
	for _, item := range items {
		if item.ReceiptExpiryWriteEnabled {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptPersistenceAuthorizationRevocationWriteEnabledCount(items []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptPersistenceAuthorizationAuditItem) int {
	count := 0
	for _, item := range items {
		if item.ReceiptRevocationWriteEnabled {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptPersistenceAuthorizationAcceptedCount(items []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptPersistenceAuthorizationAuditItem) int {
	count := 0
	for _, item := range items {
		if item.ReceiptAccepted || item.AuthorizationAccepted {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptPersistenceAuthorizationEnabledConsumerCount(items []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptPersistenceAuthorizationAuditItem) int {
	count := 0
	for _, item := range items {
		if item.ConsumerConsumptionAuthorized || item.ConsumerEnablementAuthorized || item.KDEConsumerEnabled || item.RuntimeConsumerEnabled {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptPersistenceAuthorizationRawExposedCount(items []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptPersistenceAuthorizationAuditItem) int {
	count := 0
	for _, item := range items {
		if item.RawResultExposed || item.StateRootPathExposed || item.FilePathsExposed || item.FileContentRead || item.InternalDetailsExposed {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptPersistenceAuthorizationSideEffectCount(items []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptPersistenceAuthorizationAuditItem) int {
	count := 0
	for _, item := range items {
		if !item.SideEffectsDisabled || item.HostRootModified || item.RequestObjectCreated || item.RequestObjectDispatched || item.PortalRequestCreated || item.NotificationActionEnabled || item.CompatibilityCenterOpened || item.SupportBundleExported || item.SupportCaseCreated {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptPersistenceAuthorizationItemIDs(items []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptPersistenceAuthorizationAuditItem) []string {
	ids := make([]string, 0, len(items))
	for _, item := range items {
		ids = append(ids, item.ID)
	}
	return ids
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptPersistenceAuthorizationItemsReady(items []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptPersistenceAuthorizationAuditItem) bool {
	if len(items) != 5 {
		return false
	}
	for _, item := range items {
		if !item.EvidencePresent || !item.WriterAuthorizationAuditConsumed || !item.PersistenceThreatReviewConsumed || !item.PersistenceAuthorizationModeled || !item.ReceiptPersistenceAuthorizationReady || item.AuthorizationStatus != "consumer-enablement-receipt-persistence-authorization-modeled-persistence-disabled" {
			return false
		}
		if item.ReceiptWriterEnabled || item.ReceiptPersistenceEnabled || item.ReceiptLookupWritesEnabled || item.ReceiptReplayEnabled || item.ReceiptExpiryWriteEnabled || item.ReceiptRevocationWriteEnabled || item.ReceiptAccepted || item.AuthorizationAccepted || item.ConsumerEnablementAuthorized || item.KDEConsumerEnabled || item.RuntimeConsumerEnabled || item.LookupRouteEnabled || item.OpaqueLookupEnabled || item.RawResultExposed || item.ProductionReadiness || item.ProductionOwnershipReady || item.KDEPolicyOwner || !item.ReviewOnly || !item.SideEffectsDisabled || item.HostRootModified || item.InternalDetailsExposed {
			return false
		}
	}
	return true
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptPersistenceAuthorizationCheckIDs(checks []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptPersistenceAuthorizationAuditCheck) []string {
	ids := make([]string, 0, len(checks))
	for _, check := range checks {
		ids = append(ids, check.ID)
	}
	return ids
}

func countProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptPersistenceAuthorizationChecks(checks []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptPersistenceAuthorizationAuditCheck) ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptPersistenceAuthorizationAuditCounts {
	counts := ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptPersistenceAuthorizationAuditCounts{Total: len(checks)}
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
