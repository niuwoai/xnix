package owner

type ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusPersistenceWriteModelAuditPreview struct {
	Version                                string                                                                                                                            `json:"version"`
	SchemaVersion                          string                                                                                                                            `json:"schema_version"`
	RequestType                            string                                                                                                                            `json:"request_type"`
	AuditType                              string                                                                                                                            `json:"audit_type"`
	Source                                 string                                                                                                                            `json:"source"`
	AuditDecision                          string                                                                                                                            `json:"audit_decision"`
	WriteModelAuditRequired                bool                                                                                                                              `json:"write_model_audit_required"`
	WriteModelModeled                      bool                                                                                                                              `json:"write_model_modeled"`
	PersistenceAuthorizationAuditConsumed  bool                                                                                                                              `json:"persistence_authorization_audit_consumed"`
	RedactedWriteModelGuidanceConsumed     bool                                                                                                                              `json:"redacted_write_model_guidance_consumed"`
	CompatibilityCenterWriteModelModeled   bool                                                                                                                              `json:"compatibility_center_write_model_modeled"`
	RuntimeDiagnosticsWriteModelModeled    bool                                                                                                                              `json:"runtime_diagnostics_write_model_modeled"`
	WriteModelBoundaryReady                bool                                                                                                                              `json:"write_model_boundary_ready"`
	PersistenceAuthorizationBoundaryReady  bool                                                                                                                              `json:"persistence_authorization_boundary_ready"`
	StatusFanOutReady                      bool                                                                                                                              `json:"status_fanout_ready"`
	KDESafeStatusOnly                      bool                                                                                                                              `json:"kde_safe_status_only"`
	OpaqueResultIDSupported                bool                                                                                                                              `json:"opaque_result_id_supported"`
	RedactedSummaryShapeModeled            bool                                                                                                                              `json:"redacted_summary_shape_modeled"`
	StatusPersistenceAuthorized            bool                                                                                                                              `json:"status_persistence_authorized"`
	StatusPersistenceWriteEnabled          bool                                                                                                                              `json:"status_persistence_write_enabled"`
	KDEStatusWriteEnabled                  bool                                                                                                                              `json:"kde_status_write_enabled"`
	RuntimeDiagnosticsWriteEnabled         bool                                                                                                                              `json:"runtime_diagnostics_write_enabled"`
	CompatibilityCenterWriteModelItemCount int                                                                                                                               `json:"compatibility_center_write_model_item_count"`
	RuntimeDiagnosticsWriteModelItemCount  int                                                                                                                               `json:"runtime_diagnostics_write_model_item_count"`
	WriteModelItemCount                    int                                                                                                                               `json:"write_model_item_count"`
	RequiredWriteModelItemCount            int                                                                                                                               `json:"required_write_model_item_count"`
	ReadyWriteModelItemCount               int                                                                                                                               `json:"ready_write_model_item_count"`
	MissingWriteModelItemCount             int                                                                                                                               `json:"missing_write_model_item_count"`
	WriteEnabledItemCount                  int                                                                                                                               `json:"write_enabled_item_count"`
	PersistedStatusItemCount               int                                                                                                                               `json:"persisted_status_item_count"`
	ConsumerEnabledWriteModelItemCount     int                                                                                                                               `json:"consumer_enabled_write_model_item_count"`
	RawExposedWriteModelItemCount          int                                                                                                                               `json:"raw_exposed_write_model_item_count"`
	SideEffectWriteModelItemCount          int                                                                                                                               `json:"side_effect_write_model_item_count"`
	WriteModelItems                        []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusPersistenceWriteModelAuditItem      `json:"write_model_items"`
	WriteModelItemIDs                      []string                                                                                                                          `json:"write_model_item_ids"`
	Checks                                 []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusPersistenceWriteModelAuditCheck     `json:"checks"`
	CheckIDs                               []string                                                                                                                          `json:"check_ids"`
	Counts                                 ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusPersistenceWriteModelAuditCheckCounts `json:"counts"`
	ConsumerConsumptionAuthorized          bool                                                                                                                              `json:"consumer_consumption_authorized"`
	ConsumerEnablementAuthorized           bool                                                                                                                              `json:"consumer_enablement_authorized"`
	KDEConsumerEnabled                     bool                                                                                                                              `json:"kde_consumer_enabled"`
	RuntimeConsumerEnabled                 bool                                                                                                                              `json:"runtime_consumer_enabled"`
	LookupRouteAuthorized                  bool                                                                                                                              `json:"lookup_route_authorized"`
	LookupRouteEnabled                     bool                                                                                                                              `json:"lookup_route_enabled"`
	OpaqueLookupEnabled                    bool                                                                                                                              `json:"opaque_lookup_enabled"`
	RedactedSummaryPersisted               bool                                                                                                                              `json:"redacted_summary_persisted"`
	KDEStatusPersisted                     bool                                                                                                                              `json:"kde_status_persisted"`
	RuntimeDiagnosticsPersisted            bool                                                                                                                              `json:"runtime_diagnostics_persisted"`
	DryRunResultPersisted                  bool                                                                                                                              `json:"dry_run_result_persisted"`
	RawResultExposed                       bool                                                                                                                              `json:"raw_result_exposed"`
	DispatchDryRunExecuted                 bool                                                                                                                              `json:"dispatch_dry_run_executed"`
	RequestObjectCreationEnabled           bool                                                                                                                              `json:"request_object_creation_enabled"`
	RequestObjectDispatchEnabled           bool                                                                                                                              `json:"request_object_dispatch_enabled"`
	PortalRequestCreated                   bool                                                                                                                              `json:"portal_request_created"`
	NotificationActionEnabled              bool                                                                                                                              `json:"notification_action_enabled"`
	CompatibilityCenterOpened              bool                                                                                                                              `json:"compatibility_center_opened"`
	SupportBundleExported                  bool                                                                                                                              `json:"support_bundle_exported"`
	SupportCaseCreated                     bool                                                                                                                              `json:"support_case_created"`
	RuntimeOwned                           bool                                                                                                                              `json:"runtime_owned"`
	GoRuntimeBacked                        bool                                                                                                                              `json:"go_runtime_backed"`
	KDEPolicyOwner                         bool                                                                                                                              `json:"kde_policy_owner"`
	OfficialDesktopOnly                    bool                                                                                                                              `json:"official_desktop_only"`
	ProductionReadiness                    bool                                                                                                                              `json:"production_readiness"`
	ProductionOwnershipReady               bool                                                                                                                              `json:"production_ownership_ready"`
	SystemServiceStarted                   bool                                                                                                                              `json:"system_service_started"`
	SessionBusClaimed                      bool                                                                                                                              `json:"session_bus_claimed"`
	ProductionBusClaimed                   bool                                                                                                                              `json:"production_bus_claimed"`
	ProductionOwnerEnabled                 bool                                                                                                                              `json:"production_owner_enabled"`
	WriteMethodsEnabled                    bool                                                                                                                              `json:"write_methods_enabled"`
	RuntimeWritesEnabled                   bool                                                                                                                              `json:"runtime_writes_enabled"`
	DesktopFilesWritten                    bool                                                                                                                              `json:"desktop_files_written"`
	SettingsPersisted                      bool                                                                                                                              `json:"settings_persisted"`
	AdapterInvocationEnabled               bool                                                                                                                              `json:"adapter_invocation_enabled"`
	BackendLaunchEnabled                   bool                                                                                                                              `json:"backend_launch_enabled"`
	BackendProcessStarted                  bool                                                                                                                              `json:"backend_process_started"`
	SnapshotRestoreExecuted                bool                                                                                                                              `json:"snapshot_restore_executed"`
	StateCleanupExecuted                   bool                                                                                                                              `json:"state_cleanup_executed"`
	NetworkRequired                        bool                                                                                                                              `json:"network_required"`
	HostRootModified                       bool                                                                                                                              `json:"host_root_modified"`
	PrivilegedContainerRequired            bool                                                                                                                              `json:"privileged_container_required"`
	CallerStateRootRequired                bool                                                                                                                              `json:"caller_state_root_required"`
	StateRootPathExposed                   bool                                                                                                                              `json:"state_root_path_exposed"`
	FilePathsExposed                       bool                                                                                                                              `json:"file_paths_exposed"`
	FileContentRead                        bool                                                                                                                              `json:"file_content_read"`
	RawCommandExposed                      bool                                                                                                                              `json:"raw_command_exposed"`
	RawExecutableExposed                   bool                                                                                                                              `json:"raw_executable_exposed"`
	BackendDetailsExposed                  bool                                                                                                                              `json:"backend_details_exposed"`
	BlockedActions                         []string                                                                                                                          `json:"blocked_actions"`
	NextRequirements                       []string                                                                                                                          `json:"next_requirements"`
	DesktopSafeSummary                     string                                                                                                                            `json:"desktop_safe_summary"`
}

type ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusPersistenceWriteModelAuditItem struct {
	ID                                     string   `json:"id"`
	ActionKind                             string   `json:"action_kind"`
	SurfaceKind                            string   `json:"surface_kind"`
	StatusConsumerKind                     string   `json:"status_consumer_kind"`
	OpaqueResultID                         string   `json:"opaque_result_id"`
	RecordKey                              string   `json:"record_key"`
	RedactedFields                         []string `json:"redacted_fields"`
	EvidencePresent                        bool     `json:"evidence_present"`
	PersistenceAuthorizationConsumed       bool     `json:"persistence_authorization_consumed"`
	WriteModelModeled                      bool     `json:"write_model_modeled"`
	RedactedSummaryShapeModeled            bool     `json:"redacted_summary_shape_modeled"`
	KDESafeStatusOnly                      bool     `json:"kde_safe_status_only"`
	CompatibilityCenterWriteModelCandidate bool     `json:"compatibility_center_write_model_candidate"`
	RuntimeDiagnosticsWriteModelCandidate  bool     `json:"runtime_diagnostics_write_model_candidate"`
	StatusPersistenceAuthorized            bool     `json:"status_persistence_authorized"`
	StatusPersistenceWriteEnabled          bool     `json:"status_persistence_write_enabled"`
	KDEStatusWriteEnabled                  bool     `json:"kde_status_write_enabled"`
	RuntimeDiagnosticsWriteEnabled         bool     `json:"runtime_diagnostics_write_enabled"`
	RedactedSummaryPersisted               bool     `json:"redacted_summary_persisted"`
	KDEStatusPersisted                     bool     `json:"kde_status_persisted"`
	RuntimeDiagnosticsPersisted            bool     `json:"runtime_diagnostics_persisted"`
	DryRunResultPersisted                  bool     `json:"dry_run_result_persisted"`
	ConsumerConsumptionAuthorized          bool     `json:"consumer_consumption_authorized"`
	ConsumerEnablementAuthorized           bool     `json:"consumer_enablement_authorized"`
	KDEConsumerEnabled                     bool     `json:"kde_consumer_enabled"`
	RuntimeConsumerEnabled                 bool     `json:"runtime_consumer_enabled"`
	LookupRouteEnabled                     bool     `json:"lookup_route_enabled"`
	OpaqueLookupEnabled                    bool     `json:"opaque_lookup_enabled"`
	RawResultExposed                       bool     `json:"raw_result_exposed"`
	DispatchDryRunExecuted                 bool     `json:"dispatch_dry_run_executed"`
	UserVisible                            bool     `json:"user_visible"`
	ReviewOnly                             bool     `json:"review_only"`
	RuntimeOwned                           bool     `json:"runtime_owned"`
	GoRuntimeBacked                        bool     `json:"go_runtime_backed"`
	KDEPolicyOwner                         bool     `json:"kde_policy_owner"`
	RequestObjectCreated                   bool     `json:"request_object_created"`
	RequestObjectDispatched                bool     `json:"request_object_dispatched"`
	PortalRequestCreated                   bool     `json:"portal_request_created"`
	NotificationActionEnabled              bool     `json:"notification_action_enabled"`
	CompatibilityCenterOpened              bool     `json:"compatibility_center_opened"`
	SupportBundleExported                  bool     `json:"support_bundle_exported"`
	SupportCaseCreated                     bool     `json:"support_case_created"`
	CallerStateRootRequired                bool     `json:"caller_state_root_required"`
	StateRootPathExposed                   bool     `json:"state_root_path_exposed"`
	FilePathsExposed                       bool     `json:"file_paths_exposed"`
	FileContentRead                        bool     `json:"file_content_read"`
	SideEffectsDisabled                    bool     `json:"side_effects_disabled"`
	HostRootModified                       bool     `json:"host_root_modified"`
	InternalDetailsExposed                 bool     `json:"internal_details_exposed"`
	WriteModelStatus                       string   `json:"write_model_status"`
	NextRequirement                        string   `json:"next_requirement"`
}

type ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusPersistenceWriteModelAuditCheck struct {
	ID      string `json:"id"`
	Status  string `json:"status"`
	Summary string `json:"summary"`
}

type ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusPersistenceWriteModelAuditCheckCounts struct {
	Total   int `json:"total"`
	Passed  int `json:"passed"`
	Pending int `json:"pending"`
	Blocked int `json:"blocked"`
}

func NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusPersistenceWriteModelAuditPreview(root string) (ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusPersistenceWriteModelAuditPreview, error) {
	version, err := readProductionDBusGateReviewVersion(root)
	if err != nil {
		return ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusPersistenceWriteModelAuditPreview{}, err
	}
	sources := productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusPersistenceWriteModelAuditSources(root)
	items := productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusPersistenceWriteModelItems(sources)
	authorizationReady := productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusPersistenceWriteModelAuthorizationReady(sources.PersistenceAuthorizationAudit)
	guidanceReady := productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusPersistenceWriteModelGuidanceReady(sources.DispatchSheet)
	compatibilityCenterReady := productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusPersistenceWriteModelCompatibilityCenterCount(items) == 5 && productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusPersistenceWriteModelReadySurfaceCount(items, "compatibility-center") == 5
	runtimeDiagnosticsReady := productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusPersistenceWriteModelRuntimeDiagnosticsCount(items) == 5 && productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusPersistenceWriteModelReadySurfaceCount(items, "runtime-diagnostics") == 5
	boundaryReady := authorizationReady && guidanceReady && compatibilityCenterReady && runtimeDiagnosticsReady && productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusPersistenceWriteModelItemsReady(items)
	preview := ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusPersistenceWriteModelAuditPreview{
		Version:                                version,
		SchemaVersion:                          "xnix.runtime.production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_kde_safe_redacted_status_persistence_write_model_audit.v1",
		RequestType:                            "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-persistence-write-model-audit-preview",
		AuditType:                              "receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-persistence-write-model-audit",
		Source:                                 "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-status-persistence-authorization-audit-preview+claude-code-current-dispatch-picks",
		AuditDecision:                          "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-persistence-write-model-audit-blocked",
		WriteModelAuditRequired:                true,
		WriteModelModeled:                      true,
		PersistenceAuthorizationAuditConsumed:  authorizationReady,
		RedactedWriteModelGuidanceConsumed:     guidanceReady,
		CompatibilityCenterWriteModelModeled:   compatibilityCenterReady,
		RuntimeDiagnosticsWriteModelModeled:    runtimeDiagnosticsReady,
		WriteModelBoundaryReady:                boundaryReady,
		PersistenceAuthorizationBoundaryReady:  authorizationReady,
		StatusFanOutReady:                      authorizationReady,
		KDESafeStatusOnly:                      authorizationReady,
		OpaqueResultIDSupported:                authorizationReady,
		RedactedSummaryShapeModeled:            boundaryReady,
		StatusPersistenceAuthorized:            false,
		StatusPersistenceWriteEnabled:          false,
		KDEStatusWriteEnabled:                  false,
		RuntimeDiagnosticsWriteEnabled:         false,
		CompatibilityCenterWriteModelItemCount: productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusPersistenceWriteModelCompatibilityCenterCount(items),
		RuntimeDiagnosticsWriteModelItemCount:  productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusPersistenceWriteModelRuntimeDiagnosticsCount(items),
		WriteModelItemCount:                    len(items),
		RequiredWriteModelItemCount:            10,
		ReadyWriteModelItemCount:               productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusPersistenceWriteModelReadyCount(items),
		MissingWriteModelItemCount:             productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusPersistenceWriteModelMissingCount(items),
		WriteEnabledItemCount:                  productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusPersistenceWriteModelWriteEnabledCount(items),
		PersistedStatusItemCount:               productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusPersistenceWriteModelPersistedCount(items),
		ConsumerEnabledWriteModelItemCount:     productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusPersistenceWriteModelConsumerEnabledCount(items),
		RawExposedWriteModelItemCount:          productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusPersistenceWriteModelRawExposedCount(items),
		SideEffectWriteModelItemCount:          productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusPersistenceWriteModelSideEffectCount(items),
		WriteModelItems:                        items,
		WriteModelItemIDs:                      productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusPersistenceWriteModelItemIDs(items),
		ConsumerConsumptionAuthorized:          false,
		ConsumerEnablementAuthorized:           false,
		KDEConsumerEnabled:                     false,
		RuntimeConsumerEnabled:                 false,
		LookupRouteAuthorized:                  false,
		LookupRouteEnabled:                     false,
		OpaqueLookupEnabled:                    false,
		RedactedSummaryPersisted:               false,
		KDEStatusPersisted:                     false,
		RuntimeDiagnosticsPersisted:            false,
		DryRunResultPersisted:                  false,
		RawResultExposed:                       false,
		DispatchDryRunExecuted:                 false,
		RequestObjectCreationEnabled:           false,
		RequestObjectDispatchEnabled:           false,
		PortalRequestCreated:                   false,
		NotificationActionEnabled:              false,
		CompatibilityCenterOpened:              false,
		SupportBundleExported:                  false,
		SupportCaseCreated:                     false,
		RuntimeOwned:                           true,
		GoRuntimeBacked:                        true,
		KDEPolicyOwner:                         false,
		OfficialDesktopOnly:                    true,
		ProductionReadiness:                    false,
		ProductionOwnershipReady:               false,
		SystemServiceStarted:                   false,
		SessionBusClaimed:                      false,
		ProductionBusClaimed:                   false,
		ProductionOwnerEnabled:                 false,
		WriteMethodsEnabled:                    false,
		RuntimeWritesEnabled:                   false,
		DesktopFilesWritten:                    false,
		SettingsPersisted:                      false,
		AdapterInvocationEnabled:               false,
		BackendLaunchEnabled:                   false,
		BackendProcessStarted:                  false,
		SnapshotRestoreExecuted:                false,
		StateCleanupExecuted:                   false,
		NetworkRequired:                        false,
		HostRootModified:                       false,
		PrivilegedContainerRequired:            false,
		CallerStateRootRequired:                false,
		StateRootPathExposed:                   false,
		FilePathsExposed:                       false,
		FileContentRead:                        false,
		RawCommandExposed:                      false,
		RawExecutableExposed:                   false,
		BackendDetailsExposed:                  false,
		BlockedActions: []string{
			"treat the redacted write model as permission to write status records",
			"enable consumers, lookup routes, opaque lookup, dry-run execution, request creation, notifications, support writes, or production ownership",
			"persist raw result data, state-root paths, host paths, file contents, backend details, raw commands, or raw executables",
		},
		NextRequirements: []string{
			"Implement an explicit redacted status writer behind a separate authorization gate before durable storage is allowed.",
			"Add retention, revocation, and consumer-read policy checks before any Compatibility Center or Runtime diagnostics records are stored.",
			"Keep write-model evidence separate from persistence authorization, consumer enablement, lookup route enablement, and execution.",
		},
		DesktopSafeSummary: "The production receipt notification action dry-run result lookup consumer enablement KDE-safe redacted status persistence write-model audit defines the future redacted status record shape for Compatibility Center and Runtime diagnostics while writing no status records, enabling no consumers or lookup routes, executing no dry runs, exposing no raw result data, launching no engines, and mutating no host state.",
	}
	checks := productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusPersistenceWriteModelAuditChecks(preview)
	preview.Checks = checks
	preview.CheckIDs = productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusPersistenceWriteModelCheckIDs(checks)
	preview.Counts = countProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusPersistenceWriteModelChecks(checks)
	if preview.Counts.Blocked == 0 {
		preview.AuditDecision = "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-persistence-write-model-audit-ready-writes-disabled"
	}
	if err := validateNoBackendTerms(preview, "production receipt notification action dry-run result lookup consumer enablement kde safe redacted status persistence write model audit preview"); err != nil {
		return ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusPersistenceWriteModelAuditPreview{}, err
	}
	return preview, nil
}

type productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusPersistenceWriteModelAuditSourceSet struct {
	PersistenceAuthorizationAudit string
	DispatchSheet                 string
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusPersistenceWriteModelAuditSources(root string) productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusPersistenceWriteModelAuditSourceSet {
	return productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusPersistenceWriteModelAuditSourceSet{
		PersistenceAuthorizationAudit: productionAuthorizationReadSources(root, []string{"internal/runtime/owner/production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_kde_safe_status_persistence_authorization_audit.go"}),
		DispatchSheet:                 productionAuthorizationReadSources(root, []string{"docs/claude-code-current-dispatch-picks.md"}),
	}
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusPersistenceWriteModelItems(sources productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusPersistenceWriteModelAuditSourceSet) []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusPersistenceWriteModelAuditItem {
	combined := sources.PersistenceAuthorizationAudit + sources.DispatchSheet
	return []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusPersistenceWriteModelAuditItem{
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusPersistenceWriteModelItem("review-receipt-dry-run-result-lookup-consumer-enablement-kde-safe-compatibility-center-redacted-status-write-model", "review", "compatibility-center", "review-result-compatibility-center-status", "review-dry-run-result-opaque-id", combined, []string{"review-receipt-dry-run-result-lookup-consumer-enablement-kde-safe-compatibility-center-status-persistence-authorization", "redacted status persistence write-model audit", "Compatibility Center", "without writing it"}, "implement gated Compatibility Center status writer separately"),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusPersistenceWriteModelItem("review-receipt-dry-run-result-lookup-consumer-enablement-kde-safe-runtime-diagnostics-redacted-status-write-model", "review", "runtime-diagnostics", "review-result-runtime-diagnostics-status", "review-dry-run-result-opaque-id", combined, []string{"review-receipt-dry-run-result-lookup-consumer-enablement-kde-safe-runtime-diagnostics-status-persistence-authorization", "redacted status persistence write-model audit", "Runtime diagnostics", "without writing it"}, "implement gated Runtime diagnostics status writer separately"),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusPersistenceWriteModelItem("renew-receipt-dry-run-result-lookup-consumer-enablement-kde-safe-compatibility-center-redacted-status-write-model", "renew", "compatibility-center", "renewal-result-compatibility-center-status", "renewal-dry-run-result-opaque-id", combined, []string{"renew-receipt-dry-run-result-lookup-consumer-enablement-kde-safe-compatibility-center-status-persistence-authorization", "redacted status persistence write-model audit", "Compatibility Center", "without writing it"}, "implement gated Compatibility Center status writer separately"),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusPersistenceWriteModelItem("renew-receipt-dry-run-result-lookup-consumer-enablement-kde-safe-runtime-diagnostics-redacted-status-write-model", "renew", "runtime-diagnostics", "renewal-result-runtime-diagnostics-status", "renewal-dry-run-result-opaque-id", combined, []string{"renew-receipt-dry-run-result-lookup-consumer-enablement-kde-safe-runtime-diagnostics-status-persistence-authorization", "redacted status persistence write-model audit", "Runtime diagnostics", "without writing it"}, "implement gated Runtime diagnostics status writer separately"),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusPersistenceWriteModelItem("open-compatibility-center-dry-run-result-lookup-consumer-enablement-kde-safe-compatibility-center-redacted-status-write-model", "open-compatibility-center", "compatibility-center", "navigation-result-compatibility-center-status", "navigation-dry-run-result-opaque-id", combined, []string{"open-compatibility-center-dry-run-result-lookup-consumer-enablement-kde-safe-compatibility-center-status-persistence-authorization", "redacted status persistence write-model audit", "Compatibility Center", "without writing it"}, "implement gated Compatibility Center status writer separately"),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusPersistenceWriteModelItem("open-compatibility-center-dry-run-result-lookup-consumer-enablement-kde-safe-runtime-diagnostics-redacted-status-write-model", "open-compatibility-center", "runtime-diagnostics", "navigation-result-runtime-diagnostics-status", "navigation-dry-run-result-opaque-id", combined, []string{"open-compatibility-center-dry-run-result-lookup-consumer-enablement-kde-safe-runtime-diagnostics-status-persistence-authorization", "redacted status persistence write-model audit", "Runtime diagnostics", "without writing it"}, "implement gated Runtime diagnostics status writer separately"),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusPersistenceWriteModelItem("dismiss-receipt-dry-run-result-lookup-consumer-enablement-kde-safe-compatibility-center-redacted-status-write-model", "dismiss", "compatibility-center", "dismissal-result-compatibility-center-status", "dismissal-dry-run-result-opaque-id", combined, []string{"dismiss-receipt-dry-run-result-lookup-consumer-enablement-kde-safe-compatibility-center-status-persistence-authorization", "redacted status persistence write-model audit", "Compatibility Center", "without writing it"}, "implement gated Compatibility Center status writer separately"),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusPersistenceWriteModelItem("dismiss-receipt-dry-run-result-lookup-consumer-enablement-kde-safe-runtime-diagnostics-redacted-status-write-model", "dismiss", "runtime-diagnostics", "dismissal-result-runtime-diagnostics-status", "dismissal-dry-run-result-opaque-id", combined, []string{"dismiss-receipt-dry-run-result-lookup-consumer-enablement-kde-safe-runtime-diagnostics-status-persistence-authorization", "redacted status persistence write-model audit", "Runtime diagnostics", "without writing it"}, "implement gated Runtime diagnostics status writer separately"),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusPersistenceWriteModelItem("support-info-dry-run-result-lookup-consumer-enablement-kde-safe-compatibility-center-redacted-status-write-model", "support-info", "compatibility-center", "support-info-result-compatibility-center-status", "support-info-dry-run-result-opaque-id", combined, []string{"support-info-dry-run-result-lookup-consumer-enablement-kde-safe-compatibility-center-status-persistence-authorization", "redacted status persistence write-model audit", "Compatibility Center", "without writing it"}, "implement gated Compatibility Center status writer separately"),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusPersistenceWriteModelItem("support-info-dry-run-result-lookup-consumer-enablement-kde-safe-runtime-diagnostics-redacted-status-write-model", "support-info", "runtime-diagnostics", "support-info-result-runtime-diagnostics-status", "support-info-dry-run-result-opaque-id", combined, []string{"support-info-dry-run-result-lookup-consumer-enablement-kde-safe-runtime-diagnostics-status-persistence-authorization", "redacted status persistence write-model audit", "Runtime diagnostics", "without writing it"}, "implement gated Runtime diagnostics status writer separately"),
	}
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusPersistenceWriteModelItem(id string, actionKind string, surfaceKind string, statusConsumerKind string, opaqueResultID string, source string, tokens []string, nextRequirement string) ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusPersistenceWriteModelAuditItem {
	ready := productionAuthorizationHasAll(source, tokens)
	status := "missing-kde-safe-redacted-status-write-model-evidence"
	if ready {
		status = "kde-safe-redacted-status-write-model-modeled-writes-disabled"
	}
	return ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusPersistenceWriteModelAuditItem{
		ID:                                     id,
		ActionKind:                             actionKind,
		SurfaceKind:                            surfaceKind,
		StatusConsumerKind:                     statusConsumerKind,
		OpaqueResultID:                         opaqueResultID,
		RecordKey:                              surfaceKind + ":" + actionKind + ":opaque-result-status",
		RedactedFields:                         []string{"action_kind", "surface_kind", "status_consumer_kind", "opaque_result_id", "compatibility_status", "safe_summary", "next_requirement"},
		EvidencePresent:                        ready,
		PersistenceAuthorizationConsumed:       ready,
		WriteModelModeled:                      ready,
		RedactedSummaryShapeModeled:            ready,
		KDESafeStatusOnly:                      ready,
		CompatibilityCenterWriteModelCandidate: surfaceKind == "compatibility-center",
		RuntimeDiagnosticsWriteModelCandidate:  surfaceKind == "runtime-diagnostics",
		StatusPersistenceAuthorized:            false,
		StatusPersistenceWriteEnabled:          false,
		KDEStatusWriteEnabled:                  false,
		RuntimeDiagnosticsWriteEnabled:         false,
		RedactedSummaryPersisted:               false,
		KDEStatusPersisted:                     false,
		RuntimeDiagnosticsPersisted:            false,
		DryRunResultPersisted:                  false,
		ConsumerConsumptionAuthorized:          false,
		ConsumerEnablementAuthorized:           false,
		KDEConsumerEnabled:                     false,
		RuntimeConsumerEnabled:                 false,
		LookupRouteEnabled:                     false,
		OpaqueLookupEnabled:                    false,
		RawResultExposed:                       false,
		DispatchDryRunExecuted:                 false,
		UserVisible:                            ready,
		ReviewOnly:                             true,
		RuntimeOwned:                           true,
		GoRuntimeBacked:                        true,
		KDEPolicyOwner:                         false,
		RequestObjectCreated:                   false,
		RequestObjectDispatched:                false,
		PortalRequestCreated:                   false,
		NotificationActionEnabled:              false,
		CompatibilityCenterOpened:              false,
		SupportBundleExported:                  false,
		SupportCaseCreated:                     false,
		CallerStateRootRequired:                false,
		StateRootPathExposed:                   false,
		FilePathsExposed:                       false,
		FileContentRead:                        false,
		SideEffectsDisabled:                    true,
		HostRootModified:                       false,
		InternalDetailsExposed:                 false,
		WriteModelStatus:                       status,
		NextRequirement:                        nextRequirement,
	}
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusPersistenceWriteModelAuditChecks(preview ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusPersistenceWriteModelAuditPreview) []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusPersistenceWriteModelAuditCheck {
	return []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusPersistenceWriteModelAuditCheck{
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusPersistenceWriteModelCheck("kde-safe-redacted-status-write-model-audit-required", productionAuthorizationPassBlocked(preview.WriteModelAuditRequired && preview.WriteModelModeled), "The KDE-safe redacted status write-model audit is present and modeled."),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusPersistenceWriteModelCheck("persistence-authorization-audit-consumed", productionAuthorizationPassBlocked(preview.PersistenceAuthorizationAuditConsumed && preview.PersistenceAuthorizationBoundaryReady), "The write model consumes the persistence authorization boundary."),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusPersistenceWriteModelCheck("redacted-write-model-guidance-consumed", productionAuthorizationPassBlocked(preview.RedactedWriteModelGuidanceConsumed), "The write model consumes current dispatch guidance."),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusPersistenceWriteModelCheck("compatibility-center-write-model-modeled", productionAuthorizationPassBlocked(preview.CompatibilityCenterWriteModelModeled && preview.CompatibilityCenterWriteModelItemCount == 5), "Compatibility Center redacted status records are modeled without writes."),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusPersistenceWriteModelCheck("runtime-diagnostics-write-model-modeled", productionAuthorizationPassBlocked(preview.RuntimeDiagnosticsWriteModelModeled && preview.RuntimeDiagnosticsWriteModelItemCount == 5), "Runtime diagnostics redacted status records are modeled without writes."),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusPersistenceWriteModelCheck("ten-kde-safe-redacted-status-write-model-items-present", productionAuthorizationPassBlocked(preview.WriteModelItemCount == 10 && preview.RequiredWriteModelItemCount == 10 && preview.ReadyWriteModelItemCount == 10 && preview.MissingWriteModelItemCount == 0), "Five notification actions have both Compatibility Center and Runtime diagnostics redacted write-model records."),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusPersistenceWriteModelCheck("redacted-shape-modeled-only", productionAuthorizationPassBlocked(preview.WriteModelBoundaryReady && preview.RedactedSummaryShapeModeled && preview.KDESafeStatusOnly && preview.OpaqueResultIDSupported && !preview.StatusPersistenceAuthorized && !preview.StatusPersistenceWriteEnabled), "The redacted shape is modeled without authorizing or enabling writes."),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusPersistenceWriteModelCheck("writes-consumers-and-lookup-disabled", productionAuthorizationPassBlocked(!preview.KDEStatusPersisted && !preview.RuntimeDiagnosticsPersisted && !preview.RedactedSummaryPersisted && !preview.DryRunResultPersisted && !preview.LookupRouteAuthorized && !preview.LookupRouteEnabled && !preview.OpaqueLookupEnabled && !preview.ConsumerConsumptionAuthorized && !preview.ConsumerEnablementAuthorized && !preview.KDEConsumerEnabled && !preview.RuntimeConsumerEnabled && preview.WriteEnabledItemCount == 0 && preview.PersistedStatusItemCount == 0 && preview.ConsumerEnabledWriteModelItemCount == 0), "Status writes, consumers, lookup, and result persistence remain disabled."),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusPersistenceWriteModelCheck("request-notification-navigation-and-support-disabled", productionAuthorizationPassBlocked(!preview.RequestObjectCreationEnabled && !preview.RequestObjectDispatchEnabled && !preview.PortalRequestCreated && !preview.NotificationActionEnabled && !preview.CompatibilityCenterOpened && !preview.SupportBundleExported && !preview.SupportCaseCreated), "Request creation, Portal requests, notifications, navigation, and support writes remain disabled."),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusPersistenceWriteModelCheck("production-and-host-boundary-closed", productionAuthorizationPassBlocked(!preview.SystemServiceStarted && !preview.SessionBusClaimed && !preview.ProductionBusClaimed && !preview.ProductionOwnerEnabled && !preview.ProductionReadiness && !preview.ProductionOwnershipReady && !preview.WriteMethodsEnabled && !preview.RuntimeWritesEnabled && !preview.AdapterInvocationEnabled && !preview.BackendLaunchEnabled && !preview.BackendProcessStarted && !preview.NetworkRequired && !preview.HostRootModified && !preview.PrivilegedContainerRequired && !preview.CallerStateRootRequired && !preview.StateRootPathExposed && !preview.FilePathsExposed && !preview.FileContentRead && !preview.RawResultExposed && !preview.RawCommandExposed && !preview.RawExecutableExposed && !preview.BackendDetailsExposed && preview.RawExposedWriteModelItemCount == 0 && preview.SideEffectWriteModelItemCount == 0 && productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusPersistenceWriteModelItemsKeepHostClosed(preview.WriteModelItems)), "Production ownership, writes, engine launch, unsafe data exposure, network, privilege, internal details, and host mutation remain disabled."),
	}
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusPersistenceWriteModelCheck(id string, status string, summary string) ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusPersistenceWriteModelAuditCheck {
	return ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusPersistenceWriteModelAuditCheck{ID: id, Status: status, Summary: summary}
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusPersistenceWriteModelAuthorizationReady(source string) bool {
	return productionAuthorizationHasAll(source, []string{
		"production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-status-persistence-authorization-audit-preview",
		"production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-status-persistence-authorization-audit-ready-persistence-disabled",
		"PersistenceAuthorizationBoundaryReady",
		"StatusFanOutReady",
		"KDESafeStatusOnly",
		"OpaqueResultIDSupported",
		"StatusPersistenceAuthorized",
		"KDEStatusPersistenceAuthorized",
		"RuntimeDiagnosticsPersistenceAuthorized",
		"RedactedSummaryPersisted",
		"KDEStatusPersisted",
		"RuntimeDiagnosticsPersisted",
	})
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusPersistenceWriteModelGuidanceReady(source string) bool {
	return productionAuthorizationHasAll(source, []string{"redacted status persistence write-model audit", "Compatibility Center", "Runtime diagnostics", "without writing it"})
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusPersistenceWriteModelReadyCount(items []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusPersistenceWriteModelAuditItem) int {
	count := 0
	for _, item := range items {
		if item.EvidencePresent {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusPersistenceWriteModelMissingCount(items []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusPersistenceWriteModelAuditItem) int {
	return len(items) - productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusPersistenceWriteModelReadyCount(items)
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusPersistenceWriteModelCompatibilityCenterCount(items []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusPersistenceWriteModelAuditItem) int {
	count := 0
	for _, item := range items {
		if item.CompatibilityCenterWriteModelCandidate {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusPersistenceWriteModelRuntimeDiagnosticsCount(items []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusPersistenceWriteModelAuditItem) int {
	count := 0
	for _, item := range items {
		if item.RuntimeDiagnosticsWriteModelCandidate {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusPersistenceWriteModelReadySurfaceCount(items []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusPersistenceWriteModelAuditItem, surfaceKind string) int {
	count := 0
	for _, item := range items {
		if item.SurfaceKind == surfaceKind && item.EvidencePresent && item.PersistenceAuthorizationConsumed && item.WriteModelModeled && item.RedactedSummaryShapeModeled {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusPersistenceWriteModelWriteEnabledCount(items []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusPersistenceWriteModelAuditItem) int {
	count := 0
	for _, item := range items {
		if item.StatusPersistenceWriteEnabled || item.KDEStatusWriteEnabled || item.RuntimeDiagnosticsWriteEnabled {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusPersistenceWriteModelPersistedCount(items []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusPersistenceWriteModelAuditItem) int {
	count := 0
	for _, item := range items {
		if item.RedactedSummaryPersisted || item.KDEStatusPersisted || item.RuntimeDiagnosticsPersisted || item.DryRunResultPersisted {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusPersistenceWriteModelConsumerEnabledCount(items []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusPersistenceWriteModelAuditItem) int {
	count := 0
	for _, item := range items {
		if item.ConsumerConsumptionAuthorized || item.ConsumerEnablementAuthorized || item.KDEConsumerEnabled || item.RuntimeConsumerEnabled {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusPersistenceWriteModelRawExposedCount(items []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusPersistenceWriteModelAuditItem) int {
	count := 0
	for _, item := range items {
		if item.RawResultExposed || item.InternalDetailsExposed || item.StateRootPathExposed || item.FilePathsExposed || item.FileContentRead {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusPersistenceWriteModelSideEffectCount(items []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusPersistenceWriteModelAuditItem) int {
	count := 0
	for _, item := range items {
		if item.DispatchDryRunExecuted || item.RequestObjectCreated || item.RequestObjectDispatched || item.PortalRequestCreated || item.NotificationActionEnabled || item.CompatibilityCenterOpened || item.SupportBundleExported || item.SupportCaseCreated || item.HostRootModified || !item.SideEffectsDisabled {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusPersistenceWriteModelItemsReady(items []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusPersistenceWriteModelAuditItem) bool {
	if len(items) != 10 {
		return false
	}
	for _, item := range items {
		if !item.EvidencePresent || !item.PersistenceAuthorizationConsumed || !item.WriteModelModeled || !item.RedactedSummaryShapeModeled || !item.KDESafeStatusOnly || item.StatusPersistenceWriteEnabled || item.RedactedSummaryPersisted || item.KDEStatusPersisted || item.RuntimeDiagnosticsPersisted || item.KDEConsumerEnabled || item.RuntimeConsumerEnabled || item.RawResultExposed || !item.SideEffectsDisabled {
			return false
		}
	}
	return true
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusPersistenceWriteModelItemsKeepHostClosed(items []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusPersistenceWriteModelAuditItem) bool {
	for _, item := range items {
		if item.CallerStateRootRequired || item.StateRootPathExposed || item.FilePathsExposed || item.FileContentRead || item.HostRootModified || item.InternalDetailsExposed {
			return false
		}
	}
	return true
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusPersistenceWriteModelItemIDs(items []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusPersistenceWriteModelAuditItem) []string {
	ids := make([]string, 0, len(items))
	for _, item := range items {
		ids = append(ids, item.ID)
	}
	return ids
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusPersistenceWriteModelCheckIDs(checks []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusPersistenceWriteModelAuditCheck) []string {
	ids := make([]string, 0, len(checks))
	for _, check := range checks {
		ids = append(ids, check.ID)
	}
	return ids
}

func countProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusPersistenceWriteModelChecks(checks []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusPersistenceWriteModelAuditCheck) ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusPersistenceWriteModelAuditCheckCounts {
	counts := ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusPersistenceWriteModelAuditCheckCounts{Total: len(checks)}
	for _, check := range checks {
		switch check.Status {
		case "pass":
			counts.Passed++
		case "pending":
			counts.Pending++
		default:
			counts.Blocked++
		}
	}
	return counts
}
