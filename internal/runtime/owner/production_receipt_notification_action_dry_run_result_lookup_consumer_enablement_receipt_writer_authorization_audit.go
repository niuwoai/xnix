package owner

type ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptWriterAuthorizationAuditPreview struct {
	Version                               string                                                                                                        `json:"version"`
	SchemaVersion                         string                                                                                                        `json:"schema_version"`
	RequestType                           string                                                                                                        `json:"request_type"`
	AuditType                             string                                                                                                        `json:"audit_type"`
	Source                                string                                                                                                        `json:"source"`
	AuditDecision                         string                                                                                                        `json:"audit_decision"`
	ReceiptSchema                         string                                                                                                        `json:"receipt_schema"`
	OpaqueReceiptID                       string                                                                                                        `json:"opaque_receipt_id"`
	WriterAuthorizationAuditRequired      bool                                                                                                          `json:"writer_authorization_audit_required"`
	WriterAuthorizationAuditModeled       bool                                                                                                          `json:"writer_authorization_audit_modeled"`
	AuthorizationReceiptAuditConsumed     bool                                                                                                          `json:"authorization_receipt_audit_consumed"`
	BaseWriterAuthorizationReviewConsumed bool                                                                                                          `json:"base_writer_authorization_review_consumed"`
	WriterAuthorizationGuidanceConsumed   bool                                                                                                          `json:"writer_authorization_guidance_consumed"`
	WriterAuthorizationBoundaryReady      bool                                                                                                          `json:"writer_authorization_boundary_ready"`
	AuthorizationReceiptBoundaryReady     bool                                                                                                          `json:"authorization_receipt_boundary_ready"`
	ConsumerEnablementReceiptWriterReady  bool                                                                                                          `json:"consumer_enablement_receipt_writer_ready"`
	ReceiptWriterAuthorizationReady       bool                                                                                                          `json:"receipt_writer_authorization_ready"`
	ReceiptPresent                        bool                                                                                                          `json:"receipt_present"`
	ReceiptAccepted                       bool                                                                                                          `json:"receipt_accepted"`
	AuthorizationAccepted                 bool                                                                                                          `json:"authorization_accepted"`
	ReceiptWriterEnabled                  bool                                                                                                          `json:"receipt_writer_enabled"`
	ReceiptPersistenceEnabled             bool                                                                                                          `json:"receipt_persistence_enabled"`
	ReceiptLookupWritesEnabled            bool                                                                                                          `json:"receipt_lookup_writes_enabled"`
	ReceiptReplayEnabled                  bool                                                                                                          `json:"receipt_replay_enabled"`
	ConsumerConsumptionAuthorized         bool                                                                                                          `json:"consumer_consumption_authorized"`
	ConsumerEnablementAuthorized          bool                                                                                                          `json:"consumer_enablement_authorized"`
	KDEConsumerEnabled                    bool                                                                                                          `json:"kde_consumer_enabled"`
	RuntimeConsumerEnabled                bool                                                                                                          `json:"runtime_consumer_enabled"`
	LookupRouteAuthorized                 bool                                                                                                          `json:"lookup_route_authorized"`
	LookupRouteEnabled                    bool                                                                                                          `json:"lookup_route_enabled"`
	OpaqueLookupEnabled                   bool                                                                                                          `json:"opaque_lookup_enabled"`
	RedactedSummaryPersisted              bool                                                                                                          `json:"redacted_summary_persisted"`
	RawResultExposed                      bool                                                                                                          `json:"raw_result_exposed"`
	RuntimeDiagnosticsPersisted           bool                                                                                                          `json:"runtime_diagnostics_persisted"`
	DryRunResultPersisted                 bool                                                                                                          `json:"dry_run_result_persisted"`
	DispatchDryRunExecuted                bool                                                                                                          `json:"dispatch_dry_run_executed"`
	RequestObjectCreationEnabled          bool                                                                                                          `json:"request_object_creation_enabled"`
	RequestObjectDispatchEnabled          bool                                                                                                          `json:"request_object_dispatch_enabled"`
	PortalRequestCreated                  bool                                                                                                          `json:"portal_request_created"`
	NotificationActionEnabled             bool                                                                                                          `json:"notification_action_enabled"`
	CompatibilityCenterOpened             bool                                                                                                          `json:"compatibility_center_opened"`
	SupportBundleExported                 bool                                                                                                          `json:"support_bundle_exported"`
	SupportCaseCreated                    bool                                                                                                          `json:"support_case_created"`
	ProductionReadiness                   bool                                                                                                          `json:"production_readiness"`
	ProductionOwnershipReady              bool                                                                                                          `json:"production_ownership_ready"`
	AuthorizationItemCount                int                                                                                                           `json:"authorization_item_count"`
	RequiredAuthorizationItemCount        int                                                                                                           `json:"required_authorization_item_count"`
	ReadyAuthorizationItemCount           int                                                                                                           `json:"ready_authorization_item_count"`
	MissingAuthorizationItemCount         int                                                                                                           `json:"missing_authorization_item_count"`
	WriteEnabledAuthorizationItemCount    int                                                                                                           `json:"write_enabled_authorization_item_count"`
	PersistedAuthorizationItemCount       int                                                                                                           `json:"persisted_authorization_item_count"`
	AcceptedAuthorizationItemCount        int                                                                                                           `json:"accepted_authorization_item_count"`
	EnabledConsumerItemCount              int                                                                                                           `json:"enabled_consumer_item_count"`
	RawExposedAuthorizationItemCount      int                                                                                                           `json:"raw_exposed_authorization_item_count"`
	SideEffectAuthorizationItemCount      int                                                                                                           `json:"side_effect_authorization_item_count"`
	AuthorizationItems                    []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptWriterAuthorizationAuditItem  `json:"authorization_items"`
	AuthorizationItemIDs                  []string                                                                                                      `json:"authorization_item_ids"`
	RequiredBeforeReceiptWrite            []string                                                                                                      `json:"required_before_receipt_write"`
	Checks                                []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptWriterAuthorizationAuditCheck `json:"checks"`
	CheckIDs                              []string                                                                                                      `json:"check_ids"`
	Counts                                ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptWriterAuthorizationAuditCounts  `json:"counts"`
	RuntimeOwned                          bool                                                                                                          `json:"runtime_owned"`
	GoRuntimeBacked                       bool                                                                                                          `json:"go_runtime_backed"`
	KDEPolicyOwner                        bool                                                                                                          `json:"kde_policy_owner"`
	OfficialDesktopOnly                   bool                                                                                                          `json:"official_desktop_only"`
	PlasmaForkRequired                    bool                                                                                                          `json:"plasma_fork_required"`
	PlasmaSourceModified                  bool                                                                                                          `json:"plasma_source_modified"`
	SystemServiceStarted                  bool                                                                                                          `json:"system_service_started"`
	SessionBusClaimed                     bool                                                                                                          `json:"session_bus_claimed"`
	ProductionBusClaimed                  bool                                                                                                          `json:"production_bus_claimed"`
	ProductionOwnerEnabled                bool                                                                                                          `json:"production_owner_enabled"`
	ProductionActivationReady             bool                                                                                                          `json:"production_activation_ready"`
	WriteMethodsEnabled                   bool                                                                                                          `json:"write_methods_enabled"`
	RuntimeWritesEnabled                  bool                                                                                                          `json:"runtime_writes_enabled"`
	RequestObjectsCreated                 bool                                                                                                          `json:"request_objects_created"`
	RequestObjectsDispatched              bool                                                                                                          `json:"request_objects_dispatched"`
	NotificationSent                      bool                                                                                                          `json:"notification_sent"`
	NotificationDeliveryEnabled           bool                                                                                                          `json:"notification_delivery_enabled"`
	CompatibilityCenterPersisted          bool                                                                                                          `json:"compatibility_center_persisted"`
	DesktopFilesWritten                   bool                                                                                                          `json:"desktop_files_written"`
	MIMEAppsWritten                       bool                                                                                                          `json:"mimeapps_written"`
	ShellConfigurationWritten             bool                                                                                                          `json:"shell_configuration_written"`
	SettingsPersisted                     bool                                                                                                          `json:"settings_persisted"`
	AdapterInvocationEnabled              bool                                                                                                          `json:"adapter_invocation_enabled"`
	BackendLaunchEnabled                  bool                                                                                                          `json:"backend_launch_enabled"`
	BackendProcessStarted                 bool                                                                                                          `json:"backend_process_started"`
	SnapshotRestoreExecuted               bool                                                                                                          `json:"snapshot_restore_executed"`
	StateCleanupExecuted                  bool                                                                                                          `json:"state_cleanup_executed"`
	NetworkRequired                       bool                                                                                                          `json:"network_required"`
	HostRootModified                      bool                                                                                                          `json:"host_root_modified"`
	PrivilegedContainerRequired           bool                                                                                                          `json:"privileged_container_required"`
	CallerStateRootRequired               bool                                                                                                          `json:"caller_state_root_required"`
	StateRootPathExposed                  bool                                                                                                          `json:"state_root_path_exposed"`
	FilePathsExposed                      bool                                                                                                          `json:"file_paths_exposed"`
	FileContentRead                       bool                                                                                                          `json:"file_content_read"`
	RawCommandExposed                     bool                                                                                                          `json:"raw_command_exposed"`
	RawExecutableExposed                  bool                                                                                                          `json:"raw_executable_exposed"`
	BackendDetailsExposed                 bool                                                                                                          `json:"backend_details_exposed"`
	BlockedActions                        []string                                                                                                      `json:"blocked_actions"`
	NextRequirements                      []string                                                                                                      `json:"next_requirements"`
	DesktopSafeSummary                    string                                                                                                        `json:"desktop_safe_summary"`
}

type ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptWriterAuthorizationAuditItem struct {
	ID                                string `json:"id"`
	ActionKind                        string `json:"action_kind"`
	AuthorizationArea                 string `json:"authorization_area"`
	RequiredEvidence                  string `json:"required_evidence"`
	EvidencePresent                   bool   `json:"evidence_present"`
	AuthorizationReceiptAuditConsumed bool   `json:"authorization_receipt_audit_consumed"`
	WriterAuthorizationModeled        bool   `json:"writer_authorization_modeled"`
	ReceiptWriterAuthorizationReady   bool   `json:"receipt_writer_authorization_ready"`
	ReceiptPresent                    bool   `json:"receipt_present"`
	ReceiptAccepted                   bool   `json:"receipt_accepted"`
	AuthorizationAccepted             bool   `json:"authorization_accepted"`
	ReceiptWriterEnabled              bool   `json:"receipt_writer_enabled"`
	ReceiptPersistenceEnabled         bool   `json:"receipt_persistence_enabled"`
	ReceiptLookupWritesEnabled        bool   `json:"receipt_lookup_writes_enabled"`
	ReceiptReplayEnabled              bool   `json:"receipt_replay_enabled"`
	ConsumerConsumptionAuthorized     bool   `json:"consumer_consumption_authorized"`
	ConsumerEnablementAuthorized      bool   `json:"consumer_enablement_authorized"`
	KDEConsumerEnabled                bool   `json:"kde_consumer_enabled"`
	RuntimeConsumerEnabled            bool   `json:"runtime_consumer_enabled"`
	LookupRouteEnabled                bool   `json:"lookup_route_enabled"`
	OpaqueLookupEnabled               bool   `json:"opaque_lookup_enabled"`
	RawResultExposed                  bool   `json:"raw_result_exposed"`
	UserVisible                       bool   `json:"user_visible"`
	ReviewOnly                        bool   `json:"review_only"`
	RuntimeOwned                      bool   `json:"runtime_owned"`
	GoRuntimeBacked                   bool   `json:"go_runtime_backed"`
	KDEPolicyOwner                    bool   `json:"kde_policy_owner"`
	CallerStateRootRequired           bool   `json:"caller_state_root_required"`
	StateRootPathExposed              bool   `json:"state_root_path_exposed"`
	FilePathsExposed                  bool   `json:"file_paths_exposed"`
	FileContentRead                   bool   `json:"file_content_read"`
	RequestObjectCreated              bool   `json:"request_object_created"`
	RequestObjectDispatched           bool   `json:"request_object_dispatched"`
	PortalRequestCreated              bool   `json:"portal_request_created"`
	NotificationActionEnabled         bool   `json:"notification_action_enabled"`
	CompatibilityCenterOpened         bool   `json:"compatibility_center_opened"`
	SupportBundleExported             bool   `json:"support_bundle_exported"`
	SupportCaseCreated                bool   `json:"support_case_created"`
	ProductionReadiness               bool   `json:"production_readiness"`
	ProductionOwnershipReady          bool   `json:"production_ownership_ready"`
	SideEffectsDisabled               bool   `json:"side_effects_disabled"`
	HostRootModified                  bool   `json:"host_root_modified"`
	InternalDetailsExposed            bool   `json:"internal_details_exposed"`
	AuthorizationStatus               string `json:"authorization_status"`
	NextRequirement                   string `json:"next_requirement"`
}

type ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptWriterAuthorizationAuditCheck struct {
	ID      string `json:"id"`
	Status  string `json:"status"`
	Summary string `json:"summary"`
}

type ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptWriterAuthorizationAuditCounts struct {
	Total   int `json:"total"`
	Passed  int `json:"passed"`
	Pending int `json:"pending"`
	Blocked int `json:"blocked"`
}

func NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptWriterAuthorizationAuditPreview(root string) (ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptWriterAuthorizationAuditPreview, error) {
	version, err := readProductionDBusGateReviewVersion(root)
	if err != nil {
		return ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptWriterAuthorizationAuditPreview{}, err
	}
	sources := productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptWriterAuthorizationAuditSources(root)
	items := productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptWriterAuthorizationItems(sources)
	authorizationReceiptReady := productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptWriterAuthorizationReceiptReady(sources.AuthorizationReceiptAudit)
	baseWriterReady := productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptWriterAuthorizationBaseWriterReady(sources.BaseWriterReview)
	guidanceReady := productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptWriterAuthorizationGuidanceReady(sources.DispatchSheet)
	boundaryReady := authorizationReceiptReady && baseWriterReady && guidanceReady && productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptWriterAuthorizationItemsReady(items)
	preview := ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptWriterAuthorizationAuditPreview{
		Version:                               version,
		SchemaVersion:                         "xnix.runtime.production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_receipt_writer_authorization_audit.v1",
		RequestType:                           "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-receipt-writer-authorization-audit-preview",
		AuditType:                             "receipt-notification-action-dry-run-result-lookup-consumer-enablement-receipt-writer-authorization-audit",
		Source:                                "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-authorization-receipt-audit-preview+production-receipt-writer-authorization-review-preview+claude-code-current-dispatch-picks",
		AuditDecision:                         "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-receipt-writer-authorization-audit-blocked",
		ReceiptSchema:                         "xnix.runtime.production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_authorization_receipt.v1",
		OpaqueReceiptID:                       ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementAuthorizationReceiptID,
		WriterAuthorizationAuditRequired:      true,
		WriterAuthorizationAuditModeled:       true,
		AuthorizationReceiptAuditConsumed:     authorizationReceiptReady,
		BaseWriterAuthorizationReviewConsumed: baseWriterReady,
		WriterAuthorizationGuidanceConsumed:   guidanceReady,
		WriterAuthorizationBoundaryReady:      boundaryReady,
		AuthorizationReceiptBoundaryReady:     authorizationReceiptReady,
		ConsumerEnablementReceiptWriterReady:  boundaryReady,
		ReceiptWriterAuthorizationReady:       boundaryReady,
		ReceiptPresent:                        false,
		ReceiptAccepted:                       false,
		AuthorizationAccepted:                 false,
		ReceiptWriterEnabled:                  false,
		ReceiptPersistenceEnabled:             false,
		ReceiptLookupWritesEnabled:            false,
		ReceiptReplayEnabled:                  false,
		ConsumerConsumptionAuthorized:         false,
		ConsumerEnablementAuthorized:          false,
		KDEConsumerEnabled:                    false,
		RuntimeConsumerEnabled:                false,
		LookupRouteAuthorized:                 false,
		LookupRouteEnabled:                    false,
		OpaqueLookupEnabled:                   false,
		RedactedSummaryPersisted:              false,
		RawResultExposed:                      false,
		RuntimeDiagnosticsPersisted:           false,
		DryRunResultPersisted:                 false,
		DispatchDryRunExecuted:                false,
		RequestObjectCreationEnabled:          false,
		RequestObjectDispatchEnabled:          false,
		PortalRequestCreated:                  false,
		NotificationActionEnabled:             false,
		CompatibilityCenterOpened:             false,
		SupportBundleExported:                 false,
		SupportCaseCreated:                    false,
		ProductionReadiness:                   false,
		ProductionOwnershipReady:              false,
		AuthorizationItemCount:                len(items),
		RequiredAuthorizationItemCount:        5,
		ReadyAuthorizationItemCount:           productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptWriterAuthorizationReadyCount(items),
		MissingAuthorizationItemCount:         productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptWriterAuthorizationMissingCount(items),
		AuthorizationItems:                    items,
		AuthorizationItemIDs:                  productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptWriterAuthorizationItemIDs(items),
		RequiredBeforeReceiptWrite: []string{
			"consumer enablement authorization receipt boundary remains ready",
			"base production receipt writer authorization review remains ready",
			"operator action for consumer enablement receipt writing is reviewed separately",
			"receipt persistence, replay, expiry, and revocation are reviewed separately",
			"receipt writes remain disabled until an accepted writer authorization is consumed by a separate gate",
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
			"treat this audit as permission to write consumer enablement authorization receipts",
			"write, persist, accept, replay, revoke, expire, or look up consumer enablement authorization receipts",
			"authorize or enable KDE or Runtime consumers from this audit",
			"create request objects, dispatch actions, create Portal requests, send notifications, open Compatibility Center, export support bundles, or create support cases",
			"claim production ownership, start services, enable Runtime writes, launch compatibility engines, expose raw results or paths, or mutate host root",
		},
		NextRequirements: []string{
			"Implement a dedicated receipt writer authorization action before any consumer enablement authorization receipt can be written.",
			"Keep persistence, replay, expiry, revocation, and lookup writes disabled until a separate persistence gate exists.",
			"Keep consumer authorization and consumer enablement disabled until an accepted receipt is consumed by a separate gate.",
			"Keep production ownership, Runtime writes, desktop side effects, engine launch, raw result exposure, and host mutation disabled.",
		},
		DesktopSafeSummary: "The production receipt notification action dry-run result lookup consumer enablement receipt writer authorization audit models the writer authorization boundary required before a future consumer enablement authorization receipt can be written, but it writes no receipt, accepts no authorization, enables no consumer or lookup route, persists no state, exposes no raw result data, launches no engine, and mutates no host state.",
	}
	preview.WriteEnabledAuthorizationItemCount = productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptWriterAuthorizationWriteEnabledCount(items)
	preview.PersistedAuthorizationItemCount = productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptWriterAuthorizationPersistedCount(items)
	preview.AcceptedAuthorizationItemCount = productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptWriterAuthorizationAcceptedCount(items)
	preview.EnabledConsumerItemCount = productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptWriterAuthorizationEnabledConsumerCount(items)
	preview.RawExposedAuthorizationItemCount = productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptWriterAuthorizationRawExposedCount(items)
	preview.SideEffectAuthorizationItemCount = productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptWriterAuthorizationSideEffectCount(items)
	checks := productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptWriterAuthorizationChecks(preview)
	preview.Checks = checks
	preview.CheckIDs = productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptWriterAuthorizationCheckIDs(checks)
	preview.Counts = countProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptWriterAuthorizationChecks(checks)
	if preview.Counts.Blocked == 0 {
		preview.AuditDecision = "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-receipt-writer-authorization-audit-ready-writes-disabled"
	}
	if err := validateNoBackendTerms(preview, "production receipt notification action dry-run result lookup consumer enablement receipt writer authorization audit preview"); err != nil {
		return ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptWriterAuthorizationAuditPreview{}, err
	}
	return preview, nil
}

type productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptWriterAuthorizationSourceSet struct {
	AuthorizationReceiptAudit string
	BaseWriterReview          string
	DispatchSheet             string
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptWriterAuthorizationAuditSources(root string) productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptWriterAuthorizationSourceSet {
	return productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptWriterAuthorizationSourceSet{
		AuthorizationReceiptAudit: productionAuthorizationReadSources(root, []string{"internal/runtime/owner/production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_authorization_receipt_audit.go"}),
		BaseWriterReview:          productionAuthorizationReadSources(root, []string{"internal/runtime/owner/production_receipt_writer_authorization_review.go"}),
		DispatchSheet:             productionAuthorizationReadSources(root, []string{"docs/claude-code-current-dispatch-picks.md"}),
	}
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptWriterAuthorizationItems(sources productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptWriterAuthorizationSourceSet) []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptWriterAuthorizationAuditItem {
	combined := sources.AuthorizationReceiptAudit + sources.BaseWriterReview + sources.DispatchSheet
	return []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptWriterAuthorizationAuditItem{
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptWriterAuthorizationItem("review-receipt-dry-run-result-lookup-consumer-enablement-writer-authorization", "review", "review-consumer-enablement-receipt-writer-authorization", combined, []string{"review-receipt-dry-run-result-lookup-consumer-enablement-authorization-receipt", "production-receipt-writer-authorization-review-preview", "receipt writer authorization audit"}, "review receipt writing requires a dedicated writer authorization gate"),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptWriterAuthorizationItem("renew-receipt-dry-run-result-lookup-consumer-enablement-writer-authorization", "renew", "renewal-consumer-enablement-receipt-writer-authorization", combined, []string{"renew-receipt-dry-run-result-lookup-consumer-enablement-authorization-receipt", "production-receipt-writer-authorization-review-preview", "receipt writer authorization audit"}, "renewal receipt writing requires a dedicated writer authorization gate"),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptWriterAuthorizationItem("open-compatibility-center-dry-run-result-lookup-consumer-enablement-writer-authorization", "open-compatibility-center", "navigation-consumer-enablement-receipt-writer-authorization", combined, []string{"open-compatibility-center-dry-run-result-lookup-consumer-enablement-authorization-receipt", "production-receipt-writer-authorization-review-preview", "receipt writer authorization audit"}, "navigation receipt writing requires a dedicated writer authorization gate"),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptWriterAuthorizationItem("dismiss-receipt-dry-run-result-lookup-consumer-enablement-writer-authorization", "dismiss", "dismissal-consumer-enablement-receipt-writer-authorization", combined, []string{"dismiss-receipt-dry-run-result-lookup-consumer-enablement-authorization-receipt", "production-receipt-writer-authorization-review-preview", "receipt writer authorization audit"}, "dismissal receipt writing requires a dedicated writer authorization gate"),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptWriterAuthorizationItem("support-info-dry-run-result-lookup-consumer-enablement-writer-authorization", "support-info", "support-info-consumer-enablement-receipt-writer-authorization", combined, []string{"support-info-dry-run-result-lookup-consumer-enablement-authorization-receipt", "production-receipt-writer-authorization-review-preview", "receipt writer authorization audit"}, "support-info receipt writing requires a dedicated writer authorization gate"),
	}
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptWriterAuthorizationItem(id string, actionKind string, authorizationArea string, source string, tokens []string, nextRequirement string) ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptWriterAuthorizationAuditItem {
	ready := productionAuthorizationHasAll(source, tokens)
	status := "missing-consumer-enablement-receipt-writer-authorization-evidence"
	if ready {
		status = "consumer-enablement-receipt-writer-authorization-modeled-writes-disabled"
	}
	return ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptWriterAuthorizationAuditItem{
		ID:                                id,
		ActionKind:                        actionKind,
		AuthorizationArea:                 authorizationArea,
		RequiredEvidence:                  "future consumer enablement authorization receipt writes require a separate writer authorization boundary",
		EvidencePresent:                   ready,
		AuthorizationReceiptAuditConsumed: ready,
		WriterAuthorizationModeled:        ready,
		ReceiptWriterAuthorizationReady:   ready,
		ReceiptPresent:                    false,
		ReceiptAccepted:                   false,
		AuthorizationAccepted:             false,
		ReceiptWriterEnabled:              false,
		ReceiptPersistenceEnabled:         false,
		ReceiptLookupWritesEnabled:        false,
		ReceiptReplayEnabled:              false,
		ConsumerConsumptionAuthorized:     false,
		ConsumerEnablementAuthorized:      false,
		KDEConsumerEnabled:                false,
		RuntimeConsumerEnabled:            false,
		LookupRouteEnabled:                false,
		OpaqueLookupEnabled:               false,
		RawResultExposed:                  false,
		UserVisible:                       ready,
		ReviewOnly:                        true,
		RuntimeOwned:                      true,
		GoRuntimeBacked:                   true,
		KDEPolicyOwner:                    false,
		CallerStateRootRequired:           false,
		StateRootPathExposed:              false,
		FilePathsExposed:                  false,
		FileContentRead:                   false,
		RequestObjectCreated:              false,
		RequestObjectDispatched:           false,
		PortalRequestCreated:              false,
		NotificationActionEnabled:         false,
		CompatibilityCenterOpened:         false,
		SupportBundleExported:             false,
		SupportCaseCreated:                false,
		ProductionReadiness:               false,
		ProductionOwnershipReady:          false,
		SideEffectsDisabled:               true,
		HostRootModified:                  false,
		InternalDetailsExposed:            false,
		AuthorizationStatus:               status,
		NextRequirement:                   nextRequirement,
	}
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptWriterAuthorizationChecks(preview ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptWriterAuthorizationAuditPreview) []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptWriterAuthorizationAuditCheck {
	return []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptWriterAuthorizationAuditCheck{
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptWriterAuthorizationCheck("authorization-receipt-audit-consumed", productionAuthorizationPassBlocked(preview.AuthorizationReceiptAuditConsumed), "The writer authorization audit consumes the consumer enablement authorization receipt boundary."),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptWriterAuthorizationCheck("base-writer-authorization-review-consumed", productionAuthorizationPassBlocked(preview.BaseWriterAuthorizationReviewConsumed), "The audit consumes the base production receipt writer authorization review."),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptWriterAuthorizationCheck("writer-authorization-guidance-consumed", productionAuthorizationPassBlocked(preview.WriterAuthorizationGuidanceConsumed), "The audit consumes current dispatch guidance for the next writer authorization boundary."),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptWriterAuthorizationCheck("writer-authorization-boundary-modeled-only", productionAuthorizationPassBlocked(preview.WriterAuthorizationAuditRequired && preview.WriterAuthorizationAuditModeled && preview.WriterAuthorizationBoundaryReady && preview.ConsumerEnablementReceiptWriterReady && preview.ReceiptWriterAuthorizationReady && !preview.ReceiptPresent && !preview.ReceiptAccepted && !preview.AuthorizationAccepted && !preview.ReceiptWriterEnabled), "The writer authorization boundary is modeled without writing or accepting a receipt."),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptWriterAuthorizationCheck("five-writer-authorization-items-present", productionAuthorizationPassBlocked(preview.AuthorizationItemCount == 5 && preview.RequiredAuthorizationItemCount == 5 && preview.MissingAuthorizationItemCount == 0), "Review, renew, open Compatibility Center, dismiss, and support-info writer authorization items are modeled."),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptWriterAuthorizationCheck("writer-authorization-items-ready-writes-disabled", productionAuthorizationPassBlocked(preview.ReadyAuthorizationItemCount == 5 && productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptWriterAuthorizationItemsReady(preview.AuthorizationItems)), "Every writer authorization item is ready while writes remain disabled."),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptWriterAuthorizationCheck("receipt-writes-persistence-and-replay-disabled", productionAuthorizationPassBlocked(!preview.ReceiptWriterEnabled && !preview.ReceiptPersistenceEnabled && !preview.ReceiptLookupWritesEnabled && !preview.ReceiptReplayEnabled && preview.WriteEnabledAuthorizationItemCount == 0 && preview.PersistedAuthorizationItemCount == 0 && preview.AcceptedAuthorizationItemCount == 0), "Receipt writing, persistence, lookup writes, replay, and acceptance remain disabled."),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptWriterAuthorizationCheck("consumer-lookup-request-and-support-disabled", productionAuthorizationPassBlocked(!preview.ConsumerConsumptionAuthorized && !preview.ConsumerEnablementAuthorized && !preview.KDEConsumerEnabled && !preview.RuntimeConsumerEnabled && !preview.LookupRouteAuthorized && !preview.LookupRouteEnabled && !preview.OpaqueLookupEnabled && !preview.RedactedSummaryPersisted && !preview.DryRunResultPersisted && !preview.DispatchDryRunExecuted && !preview.RequestObjectCreationEnabled && !preview.RequestObjectDispatchEnabled && !preview.RequestObjectsCreated && !preview.RequestObjectsDispatched && !preview.PortalRequestCreated && !preview.NotificationActionEnabled && !preview.CompatibilityCenterOpened && !preview.SupportBundleExported && !preview.SupportCaseCreated && preview.EnabledConsumerItemCount == 0), "Consumer authorization, lookup enablement, request creation, dispatch, notifications, navigation, and support writes remain disabled."),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptWriterAuthorizationCheck("production-and-host-boundary-closed", productionAuthorizationPassBlocked(!preview.SystemServiceStarted && !preview.SessionBusClaimed && !preview.ProductionBusClaimed && !preview.ProductionOwnerEnabled && !preview.ProductionActivationReady && !preview.ProductionReadiness && !preview.ProductionOwnershipReady && !preview.WriteMethodsEnabled && !preview.RuntimeWritesEnabled && !preview.AdapterInvocationEnabled && !preview.BackendLaunchEnabled && !preview.BackendProcessStarted && !preview.NetworkRequired && !preview.HostRootModified && !preview.PrivilegedContainerRequired && !preview.StateRootPathExposed && !preview.FilePathsExposed && !preview.FileContentRead && !preview.RawResultExposed && !preview.RawCommandExposed && !preview.RawExecutableExposed && !preview.BackendDetailsExposed && preview.RawExposedAuthorizationItemCount == 0 && preview.SideEffectAuthorizationItemCount == 0), "Production ownership, writes, engine launch, unsafe data exposure, network, privilege, internal details, and host mutation remain disabled."),
	}
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptWriterAuthorizationCheck(id string, status string, summary string) ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptWriterAuthorizationAuditCheck {
	return ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptWriterAuthorizationAuditCheck{ID: id, Status: status, Summary: summary}
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptWriterAuthorizationReceiptReady(source string) bool {
	return productionAuthorizationHasAll(source, []string{
		"production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-authorization-receipt-audit-preview",
		"production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-authorization-receipt-audit-ready-receipt-disabled",
		"dry-run-result-lookup-consumer-enablement-authorization-receipt-id",
		"ConsumerAuthorizationReady",
		"ReceiptWriterEnabled",
	})
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptWriterAuthorizationBaseWriterReady(source string) bool {
	return productionAuthorizationHasAll(source, []string{
		"production-receipt-writer-authorization-review-preview",
		"production-receipt-writer-authorization-review-ready-writes-disabled",
		"WriterAuthorizationRequired",
		"WriterAuthorizationModeled",
		"ReceiptWriterEnabled",
	})
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptWriterAuthorizationGuidanceReady(source string) bool {
	return productionAuthorizationHasAll(source, []string{"receipt writer authorization audit", "consumer enablement authorization receipt", "receipt writes", "side effects"})
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptWriterAuthorizationReadyCount(items []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptWriterAuthorizationAuditItem) int {
	count := 0
	for _, item := range items {
		if item.EvidencePresent {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptWriterAuthorizationMissingCount(items []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptWriterAuthorizationAuditItem) int {
	count := 0
	for _, item := range items {
		if !item.EvidencePresent {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptWriterAuthorizationWriteEnabledCount(items []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptWriterAuthorizationAuditItem) int {
	count := 0
	for _, item := range items {
		if item.ReceiptWriterEnabled || item.ReceiptLookupWritesEnabled {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptWriterAuthorizationPersistedCount(items []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptWriterAuthorizationAuditItem) int {
	count := 0
	for _, item := range items {
		if item.ReceiptPersistenceEnabled || item.ReceiptReplayEnabled {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptWriterAuthorizationAcceptedCount(items []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptWriterAuthorizationAuditItem) int {
	count := 0
	for _, item := range items {
		if item.ReceiptPresent || item.ReceiptAccepted || item.AuthorizationAccepted {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptWriterAuthorizationEnabledConsumerCount(items []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptWriterAuthorizationAuditItem) int {
	count := 0
	for _, item := range items {
		if item.ConsumerConsumptionAuthorized || item.ConsumerEnablementAuthorized || item.KDEConsumerEnabled || item.RuntimeConsumerEnabled || item.LookupRouteEnabled || item.OpaqueLookupEnabled {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptWriterAuthorizationRawExposedCount(items []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptWriterAuthorizationAuditItem) int {
	count := 0
	for _, item := range items {
		if item.RawResultExposed || item.CallerStateRootRequired || item.StateRootPathExposed || item.FilePathsExposed || item.FileContentRead || item.InternalDetailsExposed {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptWriterAuthorizationSideEffectCount(items []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptWriterAuthorizationAuditItem) int {
	count := 0
	for _, item := range items {
		if item.ReceiptPresent || item.ReceiptAccepted || item.AuthorizationAccepted || item.ReceiptWriterEnabled || item.ReceiptPersistenceEnabled || item.ReceiptLookupWritesEnabled || item.ReceiptReplayEnabled || item.ConsumerConsumptionAuthorized || item.ConsumerEnablementAuthorized || item.KDEConsumerEnabled || item.RuntimeConsumerEnabled || item.LookupRouteEnabled || item.OpaqueLookupEnabled || item.RawResultExposed || item.RequestObjectCreated || item.RequestObjectDispatched || item.PortalRequestCreated || item.NotificationActionEnabled || item.CompatibilityCenterOpened || item.SupportBundleExported || item.SupportCaseCreated || item.HostRootModified || item.InternalDetailsExposed {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptWriterAuthorizationItemIDs(items []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptWriterAuthorizationAuditItem) []string {
	ids := make([]string, 0, len(items))
	for _, item := range items {
		ids = append(ids, item.ID)
	}
	return ids
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptWriterAuthorizationCheckIDs(checks []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptWriterAuthorizationAuditCheck) []string {
	ids := make([]string, 0, len(checks))
	for _, check := range checks {
		ids = append(ids, check.ID)
	}
	return ids
}

func countProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptWriterAuthorizationChecks(checks []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptWriterAuthorizationAuditCheck) ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptWriterAuthorizationAuditCounts {
	counts := ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptWriterAuthorizationAuditCounts{Total: len(checks)}
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

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptWriterAuthorizationItemsReady(items []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptWriterAuthorizationAuditItem) bool {
	if len(items) != 5 {
		return false
	}
	for _, item := range items {
		if !item.EvidencePresent || !item.AuthorizationReceiptAuditConsumed || !item.WriterAuthorizationModeled || !item.ReceiptWriterAuthorizationReady || item.AuthorizationStatus != "consumer-enablement-receipt-writer-authorization-modeled-writes-disabled" {
			return false
		}
		if item.ReceiptPresent || item.ReceiptAccepted || item.AuthorizationAccepted || item.ReceiptWriterEnabled || item.ReceiptPersistenceEnabled || item.ReceiptLookupWritesEnabled || item.ReceiptReplayEnabled || item.ConsumerConsumptionAuthorized || item.ConsumerEnablementAuthorized || item.KDEConsumerEnabled || item.RuntimeConsumerEnabled || item.LookupRouteEnabled || item.OpaqueLookupEnabled || item.RawResultExposed || item.CallerStateRootRequired || item.StateRootPathExposed || item.FilePathsExposed || item.FileContentRead || item.RequestObjectCreated || item.RequestObjectDispatched || item.PortalRequestCreated || item.NotificationActionEnabled || item.CompatibilityCenterOpened || item.SupportBundleExported || item.SupportCaseCreated || item.ProductionReadiness || item.ProductionOwnershipReady || item.KDEPolicyOwner || !item.ReviewOnly || !item.SideEffectsDisabled || item.HostRootModified || item.InternalDetailsExposed {
			return false
		}
	}
	return true
}
