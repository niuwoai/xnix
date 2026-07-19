package owner

type ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusPersistenceAuthorizationAuditPreview struct {
	Version                                            string                                                                                                                       `json:"version"`
	SchemaVersion                                      string                                                                                                                       `json:"schema_version"`
	RequestType                                        string                                                                                                                       `json:"request_type"`
	AuditType                                          string                                                                                                                       `json:"audit_type"`
	Source                                             string                                                                                                                       `json:"source"`
	AuditDecision                                      string                                                                                                                       `json:"audit_decision"`
	PersistenceAuthorizationAuditRequired              bool                                                                                                                         `json:"persistence_authorization_audit_required"`
	PersistenceAuthorizationModeled                    bool                                                                                                                         `json:"persistence_authorization_modeled"`
	StatusFanOutAuditConsumed                          bool                                                                                                                         `json:"status_fanout_audit_consumed"`
	KDESafeStatusPersistenceGuidanceConsumed           bool                                                                                                                         `json:"kde_safe_status_persistence_guidance_consumed"`
	CompatibilityCenterPersistenceAuthorizationModeled bool                                                                                                                         `json:"compatibility_center_persistence_authorization_modeled"`
	RuntimeDiagnosticsPersistenceAuthorizationModeled  bool                                                                                                                         `json:"runtime_diagnostics_persistence_authorization_modeled"`
	PersistenceAuthorizationBoundaryReady              bool                                                                                                                         `json:"persistence_authorization_boundary_ready"`
	StatusFanOutReady                                  bool                                                                                                                         `json:"status_fanout_ready"`
	ConsumerEnablementGateClosed                       bool                                                                                                                         `json:"consumer_enablement_gate_closed"`
	KDESafeStatusOnly                                  bool                                                                                                                         `json:"kde_safe_status_only"`
	OpaqueResultIDSupported                            bool                                                                                                                         `json:"opaque_result_id_supported"`
	StatusPersistenceAuthorized                        bool                                                                                                                         `json:"status_persistence_authorized"`
	KDEStatusPersistenceAuthorized                     bool                                                                                                                         `json:"kde_status_persistence_authorized"`
	RuntimeDiagnosticsPersistenceAuthorized            bool                                                                                                                         `json:"runtime_diagnostics_persistence_authorized"`
	CompatibilityCenterAuthorizationItemCount          int                                                                                                                          `json:"compatibility_center_authorization_item_count"`
	RuntimeDiagnosticsAuthorizationItemCount           int                                                                                                                          `json:"runtime_diagnostics_authorization_item_count"`
	AuthorizationItemCount                             int                                                                                                                          `json:"authorization_item_count"`
	RequiredAuthorizationItemCount                     int                                                                                                                          `json:"required_authorization_item_count"`
	ReadyAuthorizationItemCount                        int                                                                                                                          `json:"ready_authorization_item_count"`
	MissingAuthorizationItemCount                      int                                                                                                                          `json:"missing_authorization_item_count"`
	AuthorizedPersistenceItemCount                     int                                                                                                                          `json:"authorized_persistence_item_count"`
	PersistedStatusItemCount                           int                                                                                                                          `json:"persisted_status_item_count"`
	ConsumerEnabledAuthorizationItemCount              int                                                                                                                          `json:"consumer_enabled_authorization_item_count"`
	RawExposedAuthorizationItemCount                   int                                                                                                                          `json:"raw_exposed_authorization_item_count"`
	SideEffectAuthorizationItemCount                   int                                                                                                                          `json:"side_effect_authorization_item_count"`
	AuthorizationItems                                 []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusPersistenceAuthorizationAuditItem      `json:"authorization_items"`
	AuthorizationItemIDs                               []string                                                                                                                     `json:"authorization_item_ids"`
	Checks                                             []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusPersistenceAuthorizationAuditCheck     `json:"checks"`
	CheckIDs                                           []string                                                                                                                     `json:"check_ids"`
	Counts                                             ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusPersistenceAuthorizationAuditCheckCounts `json:"counts"`
	ConsumerConsumptionAuthorized                      bool                                                                                                                         `json:"consumer_consumption_authorized"`
	ConsumerEnablementAuthorized                       bool                                                                                                                         `json:"consumer_enablement_authorized"`
	KDEConsumerEnabled                                 bool                                                                                                                         `json:"kde_consumer_enabled"`
	RuntimeConsumerEnabled                             bool                                                                                                                         `json:"runtime_consumer_enabled"`
	LookupRouteAuthorized                              bool                                                                                                                         `json:"lookup_route_authorized"`
	LookupRouteEnabled                                 bool                                                                                                                         `json:"lookup_route_enabled"`
	LookupRoutePersisted                               bool                                                                                                                         `json:"lookup_route_persisted"`
	OpaqueLookupEnabled                                bool                                                                                                                         `json:"opaque_lookup_enabled"`
	OpaqueLookupPersisted                              bool                                                                                                                         `json:"opaque_lookup_persisted"`
	RedactedSummaryPersisted                           bool                                                                                                                         `json:"redacted_summary_persisted"`
	KDEStatusPersisted                                 bool                                                                                                                         `json:"kde_status_persisted"`
	RuntimeDiagnosticsPersisted                        bool                                                                                                                         `json:"runtime_diagnostics_persisted"`
	DryRunResultPersisted                              bool                                                                                                                         `json:"dry_run_result_persisted"`
	RawResultExposed                                   bool                                                                                                                         `json:"raw_result_exposed"`
	DispatchDryRunExecuted                             bool                                                                                                                         `json:"dispatch_dry_run_executed"`
	RequestObjectCreationEnabled                       bool                                                                                                                         `json:"request_object_creation_enabled"`
	RequestObjectDispatchEnabled                       bool                                                                                                                         `json:"request_object_dispatch_enabled"`
	PortalRequestCreated                               bool                                                                                                                         `json:"portal_request_created"`
	NotificationActionEnabled                          bool                                                                                                                         `json:"notification_action_enabled"`
	CompatibilityCenterOpened                          bool                                                                                                                         `json:"compatibility_center_opened"`
	SupportBundleExported                              bool                                                                                                                         `json:"support_bundle_exported"`
	SupportCaseCreated                                 bool                                                                                                                         `json:"support_case_created"`
	RuntimeOwned                                       bool                                                                                                                         `json:"runtime_owned"`
	GoRuntimeBacked                                    bool                                                                                                                         `json:"go_runtime_backed"`
	KDEPolicyOwner                                     bool                                                                                                                         `json:"kde_policy_owner"`
	OfficialDesktopOnly                                bool                                                                                                                         `json:"official_desktop_only"`
	PlasmaForkRequired                                 bool                                                                                                                         `json:"plasma_fork_required"`
	PlasmaSourceModified                               bool                                                                                                                         `json:"plasma_source_modified"`
	ProductionReadiness                                bool                                                                                                                         `json:"production_readiness"`
	ProductionOwnershipReady                           bool                                                                                                                         `json:"production_ownership_ready"`
	SystemServiceStarted                               bool                                                                                                                         `json:"system_service_started"`
	SessionBusClaimed                                  bool                                                                                                                         `json:"session_bus_claimed"`
	ProductionBusClaimed                               bool                                                                                                                         `json:"production_bus_claimed"`
	ProductionOwnerEnabled                             bool                                                                                                                         `json:"production_owner_enabled"`
	ProductionActivationReady                          bool                                                                                                                         `json:"production_activation_ready"`
	WriteMethodsEnabled                                bool                                                                                                                         `json:"write_methods_enabled"`
	RuntimeWritesEnabled                               bool                                                                                                                         `json:"runtime_writes_enabled"`
	RequestObjectsCreated                              bool                                                                                                                         `json:"request_objects_created"`
	RequestObjectsDispatched                           bool                                                                                                                         `json:"request_objects_dispatched"`
	NotificationSent                                   bool                                                                                                                         `json:"notification_sent"`
	NotificationDeliveryEnabled                        bool                                                                                                                         `json:"notification_delivery_enabled"`
	DesktopFilesWritten                                bool                                                                                                                         `json:"desktop_files_written"`
	MIMEAppsWritten                                    bool                                                                                                                         `json:"mimeapps_written"`
	ShellConfigurationWritten                          bool                                                                                                                         `json:"shell_configuration_written"`
	SettingsPersisted                                  bool                                                                                                                         `json:"settings_persisted"`
	AdapterInvocationEnabled                           bool                                                                                                                         `json:"adapter_invocation_enabled"`
	BackendLaunchEnabled                               bool                                                                                                                         `json:"backend_launch_enabled"`
	BackendProcessStarted                              bool                                                                                                                         `json:"backend_process_started"`
	SnapshotRestoreExecuted                            bool                                                                                                                         `json:"snapshot_restore_executed"`
	StateCleanupExecuted                               bool                                                                                                                         `json:"state_cleanup_executed"`
	NetworkRequired                                    bool                                                                                                                         `json:"network_required"`
	HostRootModified                                   bool                                                                                                                         `json:"host_root_modified"`
	PrivilegedContainerRequired                        bool                                                                                                                         `json:"privileged_container_required"`
	CallerStateRootRequired                            bool                                                                                                                         `json:"caller_state_root_required"`
	StateRootPathExposed                               bool                                                                                                                         `json:"state_root_path_exposed"`
	FilePathsExposed                                   bool                                                                                                                         `json:"file_paths_exposed"`
	FileContentRead                                    bool                                                                                                                         `json:"file_content_read"`
	RawCommandExposed                                  bool                                                                                                                         `json:"raw_command_exposed"`
	RawExecutableExposed                               bool                                                                                                                         `json:"raw_executable_exposed"`
	BackendDetailsExposed                              bool                                                                                                                         `json:"backend_details_exposed"`
	BlockedActions                                     []string                                                                                                                     `json:"blocked_actions"`
	NextRequirements                                   []string                                                                                                                     `json:"next_requirements"`
	DesktopSafeSummary                                 string                                                                                                                       `json:"desktop_safe_summary"`
}

type ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusPersistenceAuthorizationAuditItem struct {
	ID                                      string `json:"id"`
	ActionKind                              string `json:"action_kind"`
	SurfaceKind                             string `json:"surface_kind"`
	StatusConsumerKind                      string `json:"status_consumer_kind"`
	OpaqueResultID                          string `json:"opaque_result_id"`
	RequiredEvidence                        string `json:"required_evidence"`
	EvidencePresent                         bool   `json:"evidence_present"`
	StatusFanOutConsumed                    bool   `json:"status_fanout_consumed"`
	PersistenceAuthorizationModeled         bool   `json:"persistence_authorization_modeled"`
	KDESafeStatusOnly                       bool   `json:"kde_safe_status_only"`
	CompatibilityCenterPersistenceCandidate bool   `json:"compatibility_center_persistence_candidate"`
	RuntimeDiagnosticsPersistenceCandidate  bool   `json:"runtime_diagnostics_persistence_candidate"`
	StatusPersistenceAuthorized             bool   `json:"status_persistence_authorized"`
	KDEStatusPersistenceAuthorized          bool   `json:"kde_status_persistence_authorized"`
	RuntimeDiagnosticsPersistenceAuthorized bool   `json:"runtime_diagnostics_persistence_authorized"`
	ConsumerConsumptionAuthorized           bool   `json:"consumer_consumption_authorized"`
	ConsumerEnablementAuthorized            bool   `json:"consumer_enablement_authorized"`
	KDEConsumerEnabled                      bool   `json:"kde_consumer_enabled"`
	RuntimeConsumerEnabled                  bool   `json:"runtime_consumer_enabled"`
	LookupRouteEnabled                      bool   `json:"lookup_route_enabled"`
	OpaqueLookupEnabled                     bool   `json:"opaque_lookup_enabled"`
	RedactedSummaryPersisted                bool   `json:"redacted_summary_persisted"`
	KDEStatusPersisted                      bool   `json:"kde_status_persisted"`
	RuntimeDiagnosticsPersisted             bool   `json:"runtime_diagnostics_persisted"`
	DryRunResultPersisted                   bool   `json:"dry_run_result_persisted"`
	RawResultExposed                        bool   `json:"raw_result_exposed"`
	DispatchDryRunExecuted                  bool   `json:"dispatch_dry_run_executed"`
	UserVisible                             bool   `json:"user_visible"`
	ReviewOnly                              bool   `json:"review_only"`
	RuntimeOwned                            bool   `json:"runtime_owned"`
	GoRuntimeBacked                         bool   `json:"go_runtime_backed"`
	KDEPolicyOwner                          bool   `json:"kde_policy_owner"`
	RequestObjectCreated                    bool   `json:"request_object_created"`
	RequestObjectDispatched                 bool   `json:"request_object_dispatched"`
	PortalRequestCreated                    bool   `json:"portal_request_created"`
	NotificationActionEnabled               bool   `json:"notification_action_enabled"`
	CompatibilityCenterOpened               bool   `json:"compatibility_center_opened"`
	SupportBundleExported                   bool   `json:"support_bundle_exported"`
	SupportCaseCreated                      bool   `json:"support_case_created"`
	CallerStateRootRequired                 bool   `json:"caller_state_root_required"`
	StateRootPathExposed                    bool   `json:"state_root_path_exposed"`
	FilePathsExposed                        bool   `json:"file_paths_exposed"`
	FileContentRead                         bool   `json:"file_content_read"`
	SideEffectsDisabled                     bool   `json:"side_effects_disabled"`
	HostRootModified                        bool   `json:"host_root_modified"`
	InternalDetailsExposed                  bool   `json:"internal_details_exposed"`
	AuthorizationStatus                     string `json:"authorization_status"`
	NextRequirement                         string `json:"next_requirement"`
}

type ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusPersistenceAuthorizationAuditCheck struct {
	ID      string `json:"id"`
	Status  string `json:"status"`
	Summary string `json:"summary"`
}

type ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusPersistenceAuthorizationAuditCheckCounts struct {
	Total   int `json:"total"`
	Passed  int `json:"passed"`
	Pending int `json:"pending"`
	Blocked int `json:"blocked"`
}

func NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusPersistenceAuthorizationAuditPreview(root string) (ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusPersistenceAuthorizationAuditPreview, error) {
	version, err := readProductionDBusGateReviewVersion(root)
	if err != nil {
		return ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusPersistenceAuthorizationAuditPreview{}, err
	}
	sources := productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusPersistenceAuthorizationAuditSources(root)
	items := productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusPersistenceAuthorizationItems(sources)
	fanOutReady := productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusPersistenceAuthorizationFanOutReady(sources.StatusFanOutAudit)
	guidanceReady := productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusPersistenceAuthorizationGuidanceReady(sources.DispatchSheet)
	compatibilityCenterReady := productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusPersistenceAuthorizationCompatibilityCenterCount(items) == 5 && productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusPersistenceAuthorizationReadySurfaceCount(items, "compatibility-center") == 5
	runtimeDiagnosticsReady := productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusPersistenceAuthorizationRuntimeDiagnosticsCount(items) == 5 && productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusPersistenceAuthorizationReadySurfaceCount(items, "runtime-diagnostics") == 5
	boundaryReady := fanOutReady && guidanceReady && compatibilityCenterReady && runtimeDiagnosticsReady && productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusPersistenceAuthorizationItemsReady(items)
	preview := ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusPersistenceAuthorizationAuditPreview{
		Version:                                  version,
		SchemaVersion:                            "xnix.runtime.production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_kde_safe_status_persistence_authorization_audit.v1",
		RequestType:                              "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-status-persistence-authorization-audit-preview",
		AuditType:                                "receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-status-persistence-authorization-audit",
		Source:                                   "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-status-fanout-audit-preview+claude-code-current-dispatch-picks",
		AuditDecision:                            "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-status-persistence-authorization-audit-blocked",
		PersistenceAuthorizationAuditRequired:    true,
		PersistenceAuthorizationModeled:          true,
		StatusFanOutAuditConsumed:                fanOutReady,
		KDESafeStatusPersistenceGuidanceConsumed: guidanceReady,
		CompatibilityCenterPersistenceAuthorizationModeled: compatibilityCenterReady,
		RuntimeDiagnosticsPersistenceAuthorizationModeled:  runtimeDiagnosticsReady,
		PersistenceAuthorizationBoundaryReady:              boundaryReady,
		StatusFanOutReady:                                  fanOutReady,
		ConsumerEnablementGateClosed:                       fanOutReady,
		KDESafeStatusOnly:                                  fanOutReady,
		OpaqueResultIDSupported:                            fanOutReady,
		StatusPersistenceAuthorized:                        false,
		KDEStatusPersistenceAuthorized:                     false,
		RuntimeDiagnosticsPersistenceAuthorized:            false,
		CompatibilityCenterAuthorizationItemCount:          productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusPersistenceAuthorizationCompatibilityCenterCount(items),
		RuntimeDiagnosticsAuthorizationItemCount:           productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusPersistenceAuthorizationRuntimeDiagnosticsCount(items),
		AuthorizationItemCount:                             len(items),
		RequiredAuthorizationItemCount:                     10,
		ReadyAuthorizationItemCount:                        productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusPersistenceAuthorizationReadyCount(items),
		MissingAuthorizationItemCount:                      productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusPersistenceAuthorizationMissingCount(items),
		AuthorizedPersistenceItemCount:                     productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusPersistenceAuthorizationAuthorizedCount(items),
		PersistedStatusItemCount:                           productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusPersistenceAuthorizationPersistedCount(items),
		ConsumerEnabledAuthorizationItemCount:              productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusPersistenceAuthorizationConsumerEnabledCount(items),
		RawExposedAuthorizationItemCount:                   productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusPersistenceAuthorizationRawExposedCount(items),
		SideEffectAuthorizationItemCount:                   productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusPersistenceAuthorizationSideEffectCount(items),
		AuthorizationItems:                                 items,
		AuthorizationItemIDs:                               productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusPersistenceAuthorizationItemIDs(items),
		ConsumerConsumptionAuthorized:                      false,
		ConsumerEnablementAuthorized:                       false,
		KDEConsumerEnabled:                                 false,
		RuntimeConsumerEnabled:                             false,
		LookupRouteAuthorized:                              false,
		LookupRouteEnabled:                                 false,
		LookupRoutePersisted:                               false,
		OpaqueLookupEnabled:                                false,
		OpaqueLookupPersisted:                              false,
		RedactedSummaryPersisted:                           false,
		KDEStatusPersisted:                                 false,
		RuntimeDiagnosticsPersisted:                        false,
		DryRunResultPersisted:                              false,
		RawResultExposed:                                   false,
		DispatchDryRunExecuted:                             false,
		RequestObjectCreationEnabled:                       false,
		RequestObjectDispatchEnabled:                       false,
		PortalRequestCreated:                               false,
		NotificationActionEnabled:                          false,
		CompatibilityCenterOpened:                          false,
		SupportBundleExported:                              false,
		SupportCaseCreated:                                 false,
		RuntimeOwned:                                       true,
		GoRuntimeBacked:                                    true,
		KDEPolicyOwner:                                     false,
		OfficialDesktopOnly:                                true,
		PlasmaForkRequired:                                 false,
		PlasmaSourceModified:                               false,
		ProductionReadiness:                                false,
		ProductionOwnershipReady:                           false,
		SystemServiceStarted:                               false,
		SessionBusClaimed:                                  false,
		ProductionBusClaimed:                               false,
		ProductionOwnerEnabled:                             false,
		ProductionActivationReady:                          false,
		WriteMethodsEnabled:                                false,
		RuntimeWritesEnabled:                               false,
		RequestObjectsCreated:                              false,
		RequestObjectsDispatched:                           false,
		NotificationSent:                                   false,
		NotificationDeliveryEnabled:                        false,
		DesktopFilesWritten:                                false,
		MIMEAppsWritten:                                    false,
		ShellConfigurationWritten:                          false,
		SettingsPersisted:                                  false,
		AdapterInvocationEnabled:                           false,
		BackendLaunchEnabled:                               false,
		BackendProcessStarted:                              false,
		SnapshotRestoreExecuted:                            false,
		StateCleanupExecuted:                               false,
		NetworkRequired:                                    false,
		HostRootModified:                                   false,
		PrivilegedContainerRequired:                        false,
		CallerStateRootRequired:                            false,
		StateRootPathExposed:                               false,
		FilePathsExposed:                                   false,
		FileContentRead:                                    false,
		RawCommandExposed:                                  false,
		RawExecutableExposed:                               false,
		BackendDetailsExposed:                              false,
		BlockedActions: []string{
			"treat status persistence authorization as permission to persist KDE or Runtime diagnostics status summaries",
			"enable consumers, lookup routes, opaque lookup, dry-run execution, request creation, notifications, support writes, or production ownership",
			"expose raw result data, state-root paths, host paths, file contents, backend details, raw commands, or raw executables",
		},
		NextRequirements: []string{
			"Implement the separate redacted status persistence write model before any durable status storage is allowed.",
			"Implement an explicit retention and revocation policy for persisted KDE-safe status summaries.",
			"Keep status persistence authorization separate from consumer enablement authorization.",
		},
		DesktopSafeSummary: "The production receipt notification action dry-run result lookup consumer enablement KDE-safe status persistence authorization audit models the authorization boundary required before status summaries can be stored durably, but it grants no persistence authorization, writes no status, enables no consumers or lookup routes, executes no dry runs, exposes no raw result data, launches no engines, and mutates no host state.",
	}
	checks := productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusPersistenceAuthorizationAuditChecks(preview)
	preview.Checks = checks
	preview.CheckIDs = productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusPersistenceAuthorizationCheckIDs(checks)
	preview.Counts = countProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusPersistenceAuthorizationChecks(checks)
	if preview.Counts.Blocked == 0 {
		preview.AuditDecision = "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-status-persistence-authorization-audit-ready-persistence-disabled"
	}
	if err := validateNoBackendTerms(preview, "production receipt notification action dry-run result lookup consumer enablement kde safe status persistence authorization audit preview"); err != nil {
		return ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusPersistenceAuthorizationAuditPreview{}, err
	}
	return preview, nil
}

type productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusPersistenceAuthorizationAuditSourceSet struct {
	StatusFanOutAudit string
	DispatchSheet     string
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusPersistenceAuthorizationAuditSources(root string) productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusPersistenceAuthorizationAuditSourceSet {
	return productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusPersistenceAuthorizationAuditSourceSet{
		StatusFanOutAudit: productionAuthorizationReadSources(root, []string{"internal/runtime/owner/production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_kde_safe_status_fanout_audit.go"}),
		DispatchSheet:     productionAuthorizationReadSources(root, []string{"docs/claude-code-current-dispatch-picks.md"}),
	}
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusPersistenceAuthorizationItems(sources productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusPersistenceAuthorizationAuditSourceSet) []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusPersistenceAuthorizationAuditItem {
	combined := sources.StatusFanOutAudit + sources.DispatchSheet
	return []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusPersistenceAuthorizationAuditItem{
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusPersistenceAuthorizationItem("review-receipt-dry-run-result-lookup-consumer-enablement-kde-safe-compatibility-center-status-persistence-authorization", "review", "compatibility-center", "review-result-compatibility-center-status", "review-dry-run-result-opaque-id", combined, []string{"review-receipt-dry-run-result-lookup-consumer-enablement-kde-safe-compatibility-center-status-fanout", "KDE-safe status persistence authorization audit", "Compatibility Center", "stored durably"}, "implement redacted Compatibility Center status storage separately"),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusPersistenceAuthorizationItem("review-receipt-dry-run-result-lookup-consumer-enablement-kde-safe-runtime-diagnostics-status-persistence-authorization", "review", "runtime-diagnostics", "review-result-runtime-diagnostics-status", "review-dry-run-result-opaque-id", combined, []string{"review-receipt-dry-run-result-lookup-consumer-enablement-kde-safe-runtime-diagnostics-status-fanout", "KDE-safe status persistence authorization audit", "Runtime diagnostics", "stored durably"}, "implement redacted Runtime diagnostics status storage separately"),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusPersistenceAuthorizationItem("renew-receipt-dry-run-result-lookup-consumer-enablement-kde-safe-compatibility-center-status-persistence-authorization", "renew", "compatibility-center", "renewal-result-compatibility-center-status", "renewal-dry-run-result-opaque-id", combined, []string{"renew-receipt-dry-run-result-lookup-consumer-enablement-kde-safe-compatibility-center-status-fanout", "KDE-safe status persistence authorization audit", "Compatibility Center", "stored durably"}, "implement redacted Compatibility Center status storage separately"),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusPersistenceAuthorizationItem("renew-receipt-dry-run-result-lookup-consumer-enablement-kde-safe-runtime-diagnostics-status-persistence-authorization", "renew", "runtime-diagnostics", "renewal-result-runtime-diagnostics-status", "renewal-dry-run-result-opaque-id", combined, []string{"renew-receipt-dry-run-result-lookup-consumer-enablement-kde-safe-runtime-diagnostics-status-fanout", "KDE-safe status persistence authorization audit", "Runtime diagnostics", "stored durably"}, "implement redacted Runtime diagnostics status storage separately"),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusPersistenceAuthorizationItem("open-compatibility-center-dry-run-result-lookup-consumer-enablement-kde-safe-compatibility-center-status-persistence-authorization", "open-compatibility-center", "compatibility-center", "navigation-result-compatibility-center-status", "navigation-dry-run-result-opaque-id", combined, []string{"open-compatibility-center-dry-run-result-lookup-consumer-enablement-kde-safe-compatibility-center-status-fanout", "KDE-safe status persistence authorization audit", "Compatibility Center", "stored durably"}, "implement redacted Compatibility Center status storage separately"),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusPersistenceAuthorizationItem("open-compatibility-center-dry-run-result-lookup-consumer-enablement-kde-safe-runtime-diagnostics-status-persistence-authorization", "open-compatibility-center", "runtime-diagnostics", "navigation-result-runtime-diagnostics-status", "navigation-dry-run-result-opaque-id", combined, []string{"open-compatibility-center-dry-run-result-lookup-consumer-enablement-kde-safe-runtime-diagnostics-status-fanout", "KDE-safe status persistence authorization audit", "Runtime diagnostics", "stored durably"}, "implement redacted Runtime diagnostics status storage separately"),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusPersistenceAuthorizationItem("dismiss-receipt-dry-run-result-lookup-consumer-enablement-kde-safe-compatibility-center-status-persistence-authorization", "dismiss", "compatibility-center", "dismissal-result-compatibility-center-status", "dismissal-dry-run-result-opaque-id", combined, []string{"dismiss-receipt-dry-run-result-lookup-consumer-enablement-kde-safe-compatibility-center-status-fanout", "KDE-safe status persistence authorization audit", "Compatibility Center", "stored durably"}, "implement redacted Compatibility Center status storage separately"),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusPersistenceAuthorizationItem("dismiss-receipt-dry-run-result-lookup-consumer-enablement-kde-safe-runtime-diagnostics-status-persistence-authorization", "dismiss", "runtime-diagnostics", "dismissal-result-runtime-diagnostics-status", "dismissal-dry-run-result-opaque-id", combined, []string{"dismiss-receipt-dry-run-result-lookup-consumer-enablement-kde-safe-runtime-diagnostics-status-fanout", "KDE-safe status persistence authorization audit", "Runtime diagnostics", "stored durably"}, "implement redacted Runtime diagnostics status storage separately"),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusPersistenceAuthorizationItem("support-info-dry-run-result-lookup-consumer-enablement-kde-safe-compatibility-center-status-persistence-authorization", "support-info", "compatibility-center", "support-info-result-compatibility-center-status", "support-info-dry-run-result-opaque-id", combined, []string{"support-info-dry-run-result-lookup-consumer-enablement-kde-safe-compatibility-center-status-fanout", "KDE-safe status persistence authorization audit", "Compatibility Center", "stored durably"}, "implement redacted Compatibility Center status storage separately"),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusPersistenceAuthorizationItem("support-info-dry-run-result-lookup-consumer-enablement-kde-safe-runtime-diagnostics-status-persistence-authorization", "support-info", "runtime-diagnostics", "support-info-result-runtime-diagnostics-status", "support-info-dry-run-result-opaque-id", combined, []string{"support-info-dry-run-result-lookup-consumer-enablement-kde-safe-runtime-diagnostics-status-fanout", "KDE-safe status persistence authorization audit", "Runtime diagnostics", "stored durably"}, "implement redacted Runtime diagnostics status storage separately"),
	}
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusPersistenceAuthorizationItem(id string, actionKind string, surfaceKind string, statusConsumerKind string, opaqueResultID string, source string, tokens []string, nextRequirement string) ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusPersistenceAuthorizationAuditItem {
	ready := productionAuthorizationHasAll(source, tokens)
	status := "missing-kde-safe-status-persistence-authorization-evidence"
	if ready {
		status = "kde-safe-status-persistence-authorization-modeled-persistence-disabled"
	}
	return ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusPersistenceAuthorizationAuditItem{
		ID:                                      id,
		ActionKind:                              actionKind,
		SurfaceKind:                             surfaceKind,
		StatusConsumerKind:                      statusConsumerKind,
		OpaqueResultID:                          opaqueResultID,
		RequiredEvidence:                        "KDE-safe status persistence requires status-only fan-out and an explicit persistence authorization boundary before durable storage",
		EvidencePresent:                         ready,
		StatusFanOutConsumed:                    ready,
		PersistenceAuthorizationModeled:         ready,
		KDESafeStatusOnly:                       ready,
		CompatibilityCenterPersistenceCandidate: ready && surfaceKind == "compatibility-center",
		RuntimeDiagnosticsPersistenceCandidate:  ready && surfaceKind == "runtime-diagnostics",
		StatusPersistenceAuthorized:             false,
		KDEStatusPersistenceAuthorized:          false,
		RuntimeDiagnosticsPersistenceAuthorized: false,
		ConsumerConsumptionAuthorized:           false,
		ConsumerEnablementAuthorized:            false,
		KDEConsumerEnabled:                      false,
		RuntimeConsumerEnabled:                  false,
		LookupRouteEnabled:                      false,
		OpaqueLookupEnabled:                     false,
		RedactedSummaryPersisted:                false,
		KDEStatusPersisted:                      false,
		RuntimeDiagnosticsPersisted:             false,
		DryRunResultPersisted:                   false,
		RawResultExposed:                        false,
		DispatchDryRunExecuted:                  false,
		UserVisible:                             ready,
		ReviewOnly:                              true,
		RuntimeOwned:                            true,
		GoRuntimeBacked:                         true,
		KDEPolicyOwner:                          false,
		RequestObjectCreated:                    false,
		RequestObjectDispatched:                 false,
		PortalRequestCreated:                    false,
		NotificationActionEnabled:               false,
		CompatibilityCenterOpened:               false,
		SupportBundleExported:                   false,
		SupportCaseCreated:                      false,
		CallerStateRootRequired:                 false,
		StateRootPathExposed:                    false,
		FilePathsExposed:                        false,
		FileContentRead:                         false,
		SideEffectsDisabled:                     true,
		HostRootModified:                        false,
		InternalDetailsExposed:                  false,
		AuthorizationStatus:                     status,
		NextRequirement:                         nextRequirement,
	}
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusPersistenceAuthorizationAuditChecks(preview ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusPersistenceAuthorizationAuditPreview) []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusPersistenceAuthorizationAuditCheck {
	return []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusPersistenceAuthorizationAuditCheck{
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusPersistenceAuthorizationCheck("kde-safe-status-persistence-authorization-audit-required", productionAuthorizationPassBlocked(preview.PersistenceAuthorizationAuditRequired && preview.PersistenceAuthorizationModeled), "The KDE-safe status persistence authorization audit is present and modeled."),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusPersistenceAuthorizationCheck("status-fanout-audit-consumed", productionAuthorizationPassBlocked(preview.StatusFanOutAuditConsumed && preview.StatusFanOutReady), "The persistence authorization audit consumes the KDE-safe status fan-out audit."),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusPersistenceAuthorizationCheck("kde-safe-status-persistence-guidance-consumed", productionAuthorizationPassBlocked(preview.KDESafeStatusPersistenceGuidanceConsumed), "The persistence authorization audit consumes current dispatch guidance."),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusPersistenceAuthorizationCheck("compatibility-center-persistence-authorization-modeled", productionAuthorizationPassBlocked(preview.CompatibilityCenterPersistenceAuthorizationModeled && preview.CompatibilityCenterAuthorizationItemCount == 5), "Compatibility Center status persistence candidates are modeled without writes."),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusPersistenceAuthorizationCheck("runtime-diagnostics-persistence-authorization-modeled", productionAuthorizationPassBlocked(preview.RuntimeDiagnosticsPersistenceAuthorizationModeled && preview.RuntimeDiagnosticsAuthorizationItemCount == 5), "Runtime diagnostics status persistence candidates are modeled without writes."),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusPersistenceAuthorizationCheck("ten-kde-safe-status-persistence-authorization-items-present", productionAuthorizationPassBlocked(preview.AuthorizationItemCount == 10 && preview.RequiredAuthorizationItemCount == 10 && preview.ReadyAuthorizationItemCount == 10 && preview.MissingAuthorizationItemCount == 0), "Five notification actions have both Compatibility Center and Runtime diagnostics persistence authorization candidates."),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusPersistenceAuthorizationCheck("persistence-authorization-boundary-modeled-only", productionAuthorizationPassBlocked(preview.PersistenceAuthorizationBoundaryReady && preview.ConsumerEnablementGateClosed && preview.KDESafeStatusOnly && preview.OpaqueResultIDSupported && !preview.StatusPersistenceAuthorized && !preview.KDEStatusPersistenceAuthorized && !preview.RuntimeDiagnosticsPersistenceAuthorized), "The boundary is modeled without granting persistence authorization."),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusPersistenceAuthorizationCheck("persistence-writes-consumers-and-lookup-disabled", productionAuthorizationPassBlocked(!preview.KDEStatusPersisted && !preview.RuntimeDiagnosticsPersisted && !preview.RedactedSummaryPersisted && !preview.DryRunResultPersisted && !preview.LookupRouteAuthorized && !preview.LookupRouteEnabled && !preview.OpaqueLookupEnabled && !preview.ConsumerConsumptionAuthorized && !preview.ConsumerEnablementAuthorized && !preview.KDEConsumerEnabled && !preview.RuntimeConsumerEnabled && preview.AuthorizedPersistenceItemCount == 0 && preview.PersistedStatusItemCount == 0 && preview.ConsumerEnabledAuthorizationItemCount == 0), "Status writes, consumers, lookup, and result persistence remain disabled."),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusPersistenceAuthorizationCheck("request-notification-navigation-and-support-disabled", productionAuthorizationPassBlocked(!preview.RequestObjectCreationEnabled && !preview.RequestObjectDispatchEnabled && !preview.RequestObjectsCreated && !preview.RequestObjectsDispatched && !preview.PortalRequestCreated && !preview.NotificationActionEnabled && !preview.NotificationSent && !preview.NotificationDeliveryEnabled && !preview.CompatibilityCenterOpened && !preview.SupportBundleExported && !preview.SupportCaseCreated), "Request creation, dispatch, Portal requests, notifications, navigation, and support writes remain disabled."),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusPersistenceAuthorizationCheck("production-and-host-boundary-closed", productionAuthorizationPassBlocked(!preview.SystemServiceStarted && !preview.SessionBusClaimed && !preview.ProductionBusClaimed && !preview.ProductionOwnerEnabled && !preview.ProductionActivationReady && !preview.ProductionReadiness && !preview.ProductionOwnershipReady && !preview.WriteMethodsEnabled && !preview.RuntimeWritesEnabled && !preview.AdapterInvocationEnabled && !preview.BackendLaunchEnabled && !preview.BackendProcessStarted && !preview.NetworkRequired && !preview.HostRootModified && !preview.PrivilegedContainerRequired && !preview.CallerStateRootRequired && !preview.StateRootPathExposed && !preview.FilePathsExposed && !preview.FileContentRead && !preview.RawResultExposed && !preview.RawCommandExposed && !preview.RawExecutableExposed && !preview.BackendDetailsExposed && preview.RawExposedAuthorizationItemCount == 0 && preview.SideEffectAuthorizationItemCount == 0 && productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusPersistenceAuthorizationItemsKeepHostClosed(preview.AuthorizationItems)), "Production ownership, writes, engine launch, unsafe data exposure, network, privilege, internal details, and host mutation remain disabled."),
	}
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusPersistenceAuthorizationCheck(id string, status string, summary string) ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusPersistenceAuthorizationAuditCheck {
	return ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusPersistenceAuthorizationAuditCheck{ID: id, Status: status, Summary: summary}
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusPersistenceAuthorizationFanOutReady(source string) bool {
	return productionAuthorizationHasAll(source, []string{
		"production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-status-fanout-audit-preview",
		"production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-status-fanout-audit-ready-status-only-consumers-disabled",
		"StatusFanOutReady",
		"ConsumerEnablementGateClosed",
		"KDESafeStatus",
		"CompatibilityCenterStatusModeled",
		"RuntimeDiagnosticsStatusModeled",
		"KDEStatusPersisted",
		"RuntimeDiagnosticsPersisted",
	})
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusPersistenceAuthorizationGuidanceReady(source string) bool {
	return productionAuthorizationHasAll(source, []string{"KDE-safe status persistence authorization audit", "Compatibility Center", "Runtime diagnostics", "stored durably"})
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusPersistenceAuthorizationReadyCount(items []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusPersistenceAuthorizationAuditItem) int {
	count := 0
	for _, item := range items {
		if item.EvidencePresent {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusPersistenceAuthorizationMissingCount(items []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusPersistenceAuthorizationAuditItem) int {
	count := 0
	for _, item := range items {
		if !item.EvidencePresent {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusPersistenceAuthorizationCompatibilityCenterCount(items []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusPersistenceAuthorizationAuditItem) int {
	count := 0
	for _, item := range items {
		if item.SurfaceKind == "compatibility-center" {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusPersistenceAuthorizationRuntimeDiagnosticsCount(items []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusPersistenceAuthorizationAuditItem) int {
	count := 0
	for _, item := range items {
		if item.SurfaceKind == "runtime-diagnostics" {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusPersistenceAuthorizationReadySurfaceCount(items []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusPersistenceAuthorizationAuditItem, surfaceKind string) int {
	count := 0
	for _, item := range items {
		if item.SurfaceKind == surfaceKind && item.EvidencePresent {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusPersistenceAuthorizationAuthorizedCount(items []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusPersistenceAuthorizationAuditItem) int {
	count := 0
	for _, item := range items {
		if item.StatusPersistenceAuthorized || item.KDEStatusPersistenceAuthorized || item.RuntimeDiagnosticsPersistenceAuthorized {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusPersistenceAuthorizationPersistedCount(items []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusPersistenceAuthorizationAuditItem) int {
	count := 0
	for _, item := range items {
		if item.RedactedSummaryPersisted || item.KDEStatusPersisted || item.RuntimeDiagnosticsPersisted || item.DryRunResultPersisted {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusPersistenceAuthorizationConsumerEnabledCount(items []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusPersistenceAuthorizationAuditItem) int {
	count := 0
	for _, item := range items {
		if item.ConsumerConsumptionAuthorized || item.ConsumerEnablementAuthorized || item.KDEConsumerEnabled || item.RuntimeConsumerEnabled || item.LookupRouteEnabled || item.OpaqueLookupEnabled {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusPersistenceAuthorizationRawExposedCount(items []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusPersistenceAuthorizationAuditItem) int {
	count := 0
	for _, item := range items {
		if item.RawResultExposed || item.StateRootPathExposed || item.FilePathsExposed || item.FileContentRead || item.InternalDetailsExposed {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusPersistenceAuthorizationSideEffectCount(items []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusPersistenceAuthorizationAuditItem) int {
	count := 0
	for _, item := range items {
		if item.RequestObjectCreated || item.RequestObjectDispatched || item.PortalRequestCreated || item.NotificationActionEnabled || item.CompatibilityCenterOpened || item.SupportBundleExported || item.SupportCaseCreated || item.HostRootModified {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusPersistenceAuthorizationItemsReady(items []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusPersistenceAuthorizationAuditItem) bool {
	if len(items) != 10 {
		return false
	}
	for _, item := range items {
		if !item.EvidencePresent || !item.StatusFanOutConsumed || !item.PersistenceAuthorizationModeled || !item.KDESafeStatusOnly || item.StatusPersistenceAuthorized || item.KDEStatusPersisted || item.RuntimeDiagnosticsPersisted || item.KDEConsumerEnabled || item.RuntimeConsumerEnabled || item.RawResultExposed || !item.SideEffectsDisabled {
			return false
		}
	}
	return true
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusPersistenceAuthorizationItemsKeepHostClosed(items []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusPersistenceAuthorizationAuditItem) bool {
	for _, item := range items {
		if item.HostRootModified || item.InternalDetailsExposed || item.StateRootPathExposed || item.FilePathsExposed || item.FileContentRead || item.RequestObjectCreated || item.RequestObjectDispatched || item.PortalRequestCreated || item.NotificationActionEnabled || item.CompatibilityCenterOpened || item.SupportBundleExported || item.SupportCaseCreated {
			return false
		}
	}
	return true
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusPersistenceAuthorizationItemIDs(items []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusPersistenceAuthorizationAuditItem) []string {
	ids := make([]string, 0, len(items))
	for _, item := range items {
		ids = append(ids, item.ID)
	}
	return ids
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusPersistenceAuthorizationCheckIDs(checks []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusPersistenceAuthorizationAuditCheck) []string {
	ids := make([]string, 0, len(checks))
	for _, check := range checks {
		ids = append(ids, check.ID)
	}
	return ids
}

func countProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusPersistenceAuthorizationChecks(checks []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusPersistenceAuthorizationAuditCheck) ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusPersistenceAuthorizationAuditCheckCounts {
	counts := ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusPersistenceAuthorizationAuditCheckCounts{Total: len(checks)}
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
