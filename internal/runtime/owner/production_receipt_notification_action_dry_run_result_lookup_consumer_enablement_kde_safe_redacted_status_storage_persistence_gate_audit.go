package owner

type ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusStoragePersistenceGateAuditPreview struct {
	Version                           string                                                                                                                         `json:"version"`
	SchemaVersion                     string                                                                                                                         `json:"schema_version"`
	RequestType                       string                                                                                                                         `json:"request_type"`
	PreviewType                       string                                                                                                                         `json:"preview_type"`
	Source                            string                                                                                                                         `json:"source"`
	PreviewDecision                   string                                                                                                                         `json:"preview_decision"`
	CurrentMainlineConsumed           bool                                                                                                                           `json:"current_mainline_consumed"`
	ClosedPersistenceWriterConsumed   bool                                                                                                                           `json:"closed_persistence_writer_consumed"`
	ClosedPersistenceWriterReady      bool                                                                                                                           `json:"closed_persistence_writer_ready"`
	StoragePersistenceGateRequired    bool                                                                                                                           `json:"storage_persistence_gate_required"`
	StoragePersistenceGateModeled     bool                                                                                                                           `json:"storage_persistence_gate_modeled"`
	StoragePersistenceGateReady       bool                                                                                                                           `json:"storage_persistence_gate_ready"`
	GateInputBoundaryModeled          bool                                                                                                                           `json:"gate_input_boundary_modeled"`
	GateOutputBoundaryModeled         bool                                                                                                                           `json:"gate_output_boundary_modeled"`
	KDESafeRedactedStatusOnly         bool                                                                                                                           `json:"kde_safe_redacted_status_only"`
	CompatibilityCenterGateModeled    bool                                                                                                                           `json:"compatibility_center_gate_modeled"`
	RuntimeDiagnosticsGateModeled     bool                                                                                                                           `json:"runtime_diagnostics_gate_modeled"`
	WriterCallable                    bool                                                                                                                           `json:"writer_callable"`
	ClosedPersistenceWriterEnabled    bool                                                                                                                           `json:"closed_persistence_writer_enabled"`
	StorageGatePassed                 bool                                                                                                                           `json:"storage_gate_passed"`
	StoragePersistenceAuthorized      bool                                                                                                                           `json:"storage_persistence_authorized"`
	WriterAuthorizationGranted        bool                                                                                                                           `json:"writer_authorization_granted"`
	StatusWriterEnabled               bool                                                                                                                           `json:"status_writer_enabled"`
	StatusPersistenceAuthorized       bool                                                                                                                           `json:"status_persistence_authorized"`
	StatusPersistenceWriteEnabled     bool                                                                                                                           `json:"status_persistence_write_enabled"`
	KDEStatusWriteEnabled             bool                                                                                                                           `json:"kde_status_write_enabled"`
	RuntimeDiagnosticsWriteEnabled    bool                                                                                                                           `json:"runtime_diagnostics_write_enabled"`
	StorageWriteEnabled               bool                                                                                                                           `json:"storage_write_enabled"`
	GateItemCount                     int                                                                                                                            `json:"gate_item_count"`
	RequiredGateItemCount             int                                                                                                                            `json:"required_gate_item_count"`
	ReadyGateItemCount                int                                                                                                                            `json:"ready_gate_item_count"`
	MissingGateItemCount              int                                                                                                                            `json:"missing_gate_item_count"`
	PassedStorageGateItemCount        int                                                                                                                            `json:"passed_storage_gate_item_count"`
	AuthorizedStorageItemCount        int                                                                                                                            `json:"authorized_storage_item_count"`
	CallableWriterItemCount           int                                                                                                                            `json:"callable_writer_item_count"`
	EnabledPersistenceWriterItemCount int                                                                                                                            `json:"enabled_persistence_writer_item_count"`
	PersistedWriterItemCount          int                                                                                                                            `json:"persisted_writer_item_count"`
	RawExposedGateItemCount           int                                                                                                                            `json:"raw_exposed_gate_item_count"`
	SideEffectGateItemCount           int                                                                                                                            `json:"side_effect_gate_item_count"`
	CompatibilityCenterItemCount      int                                                                                                                            `json:"compatibility_center_item_count"`
	RuntimeDiagnosticsItemCount       int                                                                                                                            `json:"runtime_diagnostics_item_count"`
	GateItems                         []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusStoragePersistenceGateAuditItem  `json:"gate_items"`
	GateItemIDs                       []string                                                                                                                       `json:"gate_item_ids"`
	Checks                            []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusStoragePersistenceGateAuditCheck `json:"checks"`
	CheckIDs                          []string                                                                                                                       `json:"check_ids"`
	Counts                            ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusStoragePersistenceGateAuditCounts  `json:"counts"`
	ConsumerConsumptionAuthorized     bool                                                                                                                           `json:"consumer_consumption_authorized"`
	ConsumerEnablementAuthorized      bool                                                                                                                           `json:"consumer_enablement_authorized"`
	KDEConsumerEnabled                bool                                                                                                                           `json:"kde_consumer_enabled"`
	RuntimeConsumerEnabled            bool                                                                                                                           `json:"runtime_consumer_enabled"`
	LookupRouteAuthorized             bool                                                                                                                           `json:"lookup_route_authorized"`
	LookupRouteEnabled                bool                                                                                                                           `json:"lookup_route_enabled"`
	OpaqueLookupEnabled               bool                                                                                                                           `json:"opaque_lookup_enabled"`
	RedactedSummaryPersisted          bool                                                                                                                           `json:"redacted_summary_persisted"`
	KDEStatusPersisted                bool                                                                                                                           `json:"kde_status_persisted"`
	RuntimeDiagnosticsPersisted       bool                                                                                                                           `json:"runtime_diagnostics_persisted"`
	DryRunResultPersisted             bool                                                                                                                           `json:"dry_run_result_persisted"`
	RawResultExposed                  bool                                                                                                                           `json:"raw_result_exposed"`
	DispatchDryRunExecuted            bool                                                                                                                           `json:"dispatch_dry_run_executed"`
	RequestObjectCreationEnabled      bool                                                                                                                           `json:"request_object_creation_enabled"`
	RequestObjectDispatchEnabled      bool                                                                                                                           `json:"request_object_dispatch_enabled"`
	PortalRequestCreated              bool                                                                                                                           `json:"portal_request_created"`
	NotificationActionEnabled         bool                                                                                                                           `json:"notification_action_enabled"`
	CompatibilityCenterOpened         bool                                                                                                                           `json:"compatibility_center_opened"`
	SupportBundleExported             bool                                                                                                                           `json:"support_bundle_exported"`
	SupportCaseCreated                bool                                                                                                                           `json:"support_case_created"`
	RuntimeOwned                      bool                                                                                                                           `json:"runtime_owned"`
	GoRuntimeBacked                   bool                                                                                                                           `json:"go_runtime_backed"`
	KDEPolicyOwner                    bool                                                                                                                           `json:"kde_policy_owner"`
	ProductionReadiness               bool                                                                                                                           `json:"production_readiness"`
	ProductionOwnershipReady          bool                                                                                                                           `json:"production_ownership_ready"`
	SystemServiceStarted              bool                                                                                                                           `json:"system_service_started"`
	SessionBusClaimed                 bool                                                                                                                           `json:"session_bus_claimed"`
	ProductionBusClaimed              bool                                                                                                                           `json:"production_bus_claimed"`
	ProductionOwnerEnabled            bool                                                                                                                           `json:"production_owner_enabled"`
	WriteMethodsEnabled               bool                                                                                                                           `json:"write_methods_enabled"`
	RuntimeWritesEnabled              bool                                                                                                                           `json:"runtime_writes_enabled"`
	DesktopFilesWritten               bool                                                                                                                           `json:"desktop_files_written"`
	SettingsPersisted                 bool                                                                                                                           `json:"settings_persisted"`
	AdapterInvocationEnabled          bool                                                                                                                           `json:"adapter_invocation_enabled"`
	BackendLaunchEnabled              bool                                                                                                                           `json:"backend_launch_enabled"`
	BackendProcessStarted             bool                                                                                                                           `json:"backend_process_started"`
	SnapshotRestoreExecuted           bool                                                                                                                           `json:"snapshot_restore_executed"`
	StateCleanupExecuted              bool                                                                                                                           `json:"state_cleanup_executed"`
	NetworkRequired                   bool                                                                                                                           `json:"network_required"`
	HostRootModified                  bool                                                                                                                           `json:"host_root_modified"`
	PrivilegedContainerRequired       bool                                                                                                                           `json:"privileged_container_required"`
	CallerStateRootRequired           bool                                                                                                                           `json:"caller_state_root_required"`
	StateRootPathExposed              bool                                                                                                                           `json:"state_root_path_exposed"`
	FilePathsExposed                  bool                                                                                                                           `json:"file_paths_exposed"`
	FileContentRead                   bool                                                                                                                           `json:"file_content_read"`
	RawCommandExposed                 bool                                                                                                                           `json:"raw_command_exposed"`
	RawExecutableExposed              bool                                                                                                                           `json:"raw_executable_exposed"`
	BackendDetailsExposed             bool                                                                                                                           `json:"backend_details_exposed"`
	BlockedActions                    []string                                                                                                                       `json:"blocked_actions"`
	NextRequirements                  []string                                                                                                                       `json:"next_requirements"`
	DesktopSafeSummary                string                                                                                                                         `json:"desktop_safe_summary"`
}

type ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusStoragePersistenceGateAuditItem struct {
	ID                              string `json:"id"`
	ActionKind                      string `json:"action_kind"`
	SurfaceKind                     string `json:"surface_kind"`
	StatusConsumerKind              string `json:"status_consumer_kind"`
	GateScope                       string `json:"gate_scope"`
	EvidencePresent                 bool   `json:"evidence_present"`
	CurrentMainlineConsumed         bool   `json:"current_mainline_consumed"`
	ClosedPersistenceWriterConsumed bool   `json:"closed_persistence_writer_consumed"`
	ClosedPersistenceWriterReady    bool   `json:"closed_persistence_writer_ready"`
	StoragePersistenceGateModeled   bool   `json:"storage_persistence_gate_modeled"`
	GateInputBoundaryModeled        bool   `json:"gate_input_boundary_modeled"`
	GateOutputBoundaryModeled       bool   `json:"gate_output_boundary_modeled"`
	KDESafeRedactedStatusOnly       bool   `json:"kde_safe_redacted_status_only"`
	WriterCallable                  bool   `json:"writer_callable"`
	ClosedPersistenceWriterEnabled  bool   `json:"closed_persistence_writer_enabled"`
	StorageGatePassed               bool   `json:"storage_gate_passed"`
	StoragePersistenceAuthorized    bool   `json:"storage_persistence_authorized"`
	StatusPersistenceAuthorized     bool   `json:"status_persistence_authorized"`
	StatusPersistenceWriteEnabled   bool   `json:"status_persistence_write_enabled"`
	StatusWriterEnabled             bool   `json:"status_writer_enabled"`
	KDEStatusWriteEnabled           bool   `json:"kde_status_write_enabled"`
	RuntimeDiagnosticsWriteEnabled  bool   `json:"runtime_diagnostics_write_enabled"`
	StorageWriteEnabled             bool   `json:"storage_write_enabled"`
	RedactedSummaryPersisted        bool   `json:"redacted_summary_persisted"`
	KDEStatusPersisted              bool   `json:"kde_status_persisted"`
	RuntimeDiagnosticsPersisted     bool   `json:"runtime_diagnostics_persisted"`
	RawResultExposed                bool   `json:"raw_result_exposed"`
	UserVisible                     bool   `json:"user_visible"`
	ReviewOnly                      bool   `json:"review_only"`
	RuntimeOwned                    bool   `json:"runtime_owned"`
	GoRuntimeBacked                 bool   `json:"go_runtime_backed"`
	KDEPolicyOwner                  bool   `json:"kde_policy_owner"`
	SideEffectsDisabled             bool   `json:"side_effects_disabled"`
	HostRootModified                bool   `json:"host_root_modified"`
	InternalDetailsExposed          bool   `json:"internal_details_exposed"`
	StoragePersistenceGateStatus    string `json:"storage_persistence_gate_status"`
	NextRequirement                 string `json:"next_requirement"`
}

type ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusStoragePersistenceGateAuditCheck struct {
	ID      string `json:"id"`
	Status  string `json:"status"`
	Summary string `json:"summary"`
}

type ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusStoragePersistenceGateAuditCounts struct {
	Total   int `json:"total"`
	Passed  int `json:"passed"`
	Pending int `json:"pending"`
	Blocked int `json:"blocked"`
}

func NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusStoragePersistenceGateAuditPreview(root string) (ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusStoragePersistenceGateAuditPreview, error) {
	version, err := readProductionDBusGateReviewVersion(root)
	if err != nil {
		return ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusStoragePersistenceGateAuditPreview{}, err
	}
	sources := productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusStoragePersistenceGateAuditSources(root)
	mainlineReady := productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusStoragePersistenceGateAuditMainlineReady(sources.CurrentMainline)
	closedWriterReady := productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusStoragePersistenceGateAuditClosedWriterReady(sources.ClosedPersistenceWriter)
	items := productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusStoragePersistenceGateAuditItems(mainlineReady, closedWriterReady)
	readyCount := productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusStoragePersistenceGateAuditReadyCount(items)
	preview := ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusStoragePersistenceGateAuditPreview{
		Version:                           version,
		SchemaVersion:                     "xnix.runtime.production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_kde_safe_redacted_status_storage_persistence_gate_audit.v1",
		RequestType:                       "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-storage-persistence-gate-audit-preview",
		PreviewType:                       "receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-storage-persistence-gate-audit",
		Source:                            "xnix-current-mainline+redacted-status-closed-persistence-writer-preview",
		PreviewDecision:                   "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-storage-persistence-gate-audit-blocked",
		CurrentMainlineConsumed:           mainlineReady,
		ClosedPersistenceWriterConsumed:   closedWriterReady,
		ClosedPersistenceWriterReady:      closedWriterReady,
		StoragePersistenceGateRequired:    true,
		StoragePersistenceGateModeled:     mainlineReady && closedWriterReady,
		StoragePersistenceGateReady:       mainlineReady && closedWriterReady,
		GateInputBoundaryModeled:          mainlineReady && closedWriterReady,
		GateOutputBoundaryModeled:         mainlineReady && closedWriterReady,
		KDESafeRedactedStatusOnly:         mainlineReady && closedWriterReady,
		CompatibilityCenterGateModeled:    mainlineReady && closedWriterReady,
		RuntimeDiagnosticsGateModeled:     mainlineReady && closedWriterReady,
		WriterCallable:                    false,
		ClosedPersistenceWriterEnabled:    false,
		StorageGatePassed:                 false,
		StoragePersistenceAuthorized:      false,
		WriterAuthorizationGranted:        false,
		StatusWriterEnabled:               false,
		StatusPersistenceAuthorized:       false,
		StatusPersistenceWriteEnabled:     false,
		KDEStatusWriteEnabled:             false,
		RuntimeDiagnosticsWriteEnabled:    false,
		StorageWriteEnabled:               false,
		GateItemCount:                     len(items),
		RequiredGateItemCount:             len(items),
		ReadyGateItemCount:                readyCount,
		MissingGateItemCount:              len(items) - readyCount,
		PassedStorageGateItemCount:        0,
		AuthorizedStorageItemCount:        0,
		CallableWriterItemCount:           0,
		EnabledPersistenceWriterItemCount: 0,
		PersistedWriterItemCount:          0,
		RawExposedGateItemCount:           0,
		SideEffectGateItemCount:           0,
		CompatibilityCenterItemCount:      productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusStoragePersistenceGateAuditSurfaceCount(items, "compatibility-center"),
		RuntimeDiagnosticsItemCount:       productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusStoragePersistenceGateAuditSurfaceCount(items, "runtime-diagnostics"),
		GateItems:                         items,
		GateItemIDs:                       productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusStoragePersistenceGateAuditItemIDs(items),
		ConsumerConsumptionAuthorized:     false,
		ConsumerEnablementAuthorized:      false,
		KDEConsumerEnabled:                false,
		RuntimeConsumerEnabled:            false,
		LookupRouteAuthorized:             false,
		LookupRouteEnabled:                false,
		OpaqueLookupEnabled:               false,
		RedactedSummaryPersisted:          false,
		KDEStatusPersisted:                false,
		RuntimeDiagnosticsPersisted:       false,
		DryRunResultPersisted:             false,
		RawResultExposed:                  false,
		DispatchDryRunExecuted:            false,
		RequestObjectCreationEnabled:      false,
		RequestObjectDispatchEnabled:      false,
		PortalRequestCreated:              false,
		NotificationActionEnabled:         false,
		CompatibilityCenterOpened:         false,
		SupportBundleExported:             false,
		SupportCaseCreated:                false,
		RuntimeOwned:                      true,
		GoRuntimeBacked:                   true,
		KDEPolicyOwner:                    false,
		ProductionReadiness:               false,
		ProductionOwnershipReady:          false,
		SystemServiceStarted:              false,
		SessionBusClaimed:                 false,
		ProductionBusClaimed:              false,
		ProductionOwnerEnabled:            false,
		WriteMethodsEnabled:               false,
		RuntimeWritesEnabled:              false,
		DesktopFilesWritten:               false,
		SettingsPersisted:                 false,
		AdapterInvocationEnabled:          false,
		BackendLaunchEnabled:              false,
		BackendProcessStarted:             false,
		SnapshotRestoreExecuted:           false,
		StateCleanupExecuted:              false,
		NetworkRequired:                   false,
		HostRootModified:                  false,
		PrivilegedContainerRequired:       false,
		CallerStateRootRequired:           false,
		StateRootPathExposed:              false,
		FilePathsExposed:                  false,
		FileContentRead:                   false,
		RawCommandExposed:                 false,
		RawExecutableExposed:              false,
		BackendDetailsExposed:             false,
		BlockedActions: []string{
			"do not pass storage persistence gates",
			"do not call closed persistence writers",
			"do not persist redacted status summaries",
			"do not enable KDE or Runtime diagnostics consumers",
			"do not create lookup, dispatch, request, Portal, notification, navigation, or support side effects",
			"do not claim production D-Bus ownership, launch engines, expose paths, or mutate the host",
		},
		NextRequirements: []string{
			"add explicit durable storage policy evidence before any storage gate can pass",
			"add a redacted status persistence record contract before any closed writer can write",
			"keep the storage persistence gate modeled but not passed until explicit production authorization exists",
		},
		DesktopSafeSummary: "The KDE-safe redacted status storage persistence gate audit consumes the current mainline and closed persistence writer preview, models the gate that must pass before any redacted status storage writer can write, and keeps writer calls, persistence writes, consumers, desktop actions, unsafe data exposure, and host mutation disabled.",
	}
	checks := productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusStoragePersistenceGateAuditChecks(preview)
	preview.Checks = checks
	preview.CheckIDs = productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusStoragePersistenceGateAuditCheckIDs(checks)
	preview.Counts = countProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusStoragePersistenceGateAuditChecks(checks)
	if preview.Counts.Blocked == 0 {
		preview.PreviewDecision = "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-storage-persistence-gate-audit-ready-storage-disabled"
	}
	if err := validateNoBackendTerms(preview, "production receipt notification action dry-run result lookup consumer enablement kde safe redacted status storage persistence gate audit"); err != nil {
		return ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusStoragePersistenceGateAuditPreview{}, err
	}
	return preview, nil
}

type productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusStoragePersistenceGateAuditSourceSet struct {
	CurrentMainline         string
	ClosedPersistenceWriter string
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusStoragePersistenceGateAuditSources(root string) productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusStoragePersistenceGateAuditSourceSet {
	return productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusStoragePersistenceGateAuditSourceSet{
		CurrentMainline:         productionAuthorizationReadSources(root, []string{"docs/xnix-current-mainline.md"}),
		ClosedPersistenceWriter: productionAuthorizationReadSources(root, []string{"internal/runtime/owner/production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_kde_safe_redacted_status_closed_persistence_writer_preview.go", "internal/runtime/owner/production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_kde_safe_redacted_status_closed_persistence_writer_preview_test.go"}),
	}
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusStoragePersistenceGateAuditMainlineReady(source string) bool {
	return productionAuthorizationHasAll(source, []string{"Xnix Current Mainline", "redacted status storage persistence gate audit", "redacted status closed persistence writer preview", "No Runtime write methods", "No KDE configuration writes"})
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusStoragePersistenceGateAuditClosedWriterReady(source string) bool {
	return productionAuthorizationHasAll(source, []string{"production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-persistence-writer-preview", "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-persistence-writer-ready-storage-disabled", "ClosedPersistenceWriterReady", "StorageWriterShapeModeled", "StorageWriteEnabled"})
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusStoragePersistenceGateAuditItems(mainlineReady bool, closedWriterReady bool) []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusStoragePersistenceGateAuditItem {
	return []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusStoragePersistenceGateAuditItem{
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusStoragePersistenceGateAuditItem("review-receipt-dry-run-result-lookup-consumer-enablement-kde-safe-compatibility-center-redacted-status-storage-persistence-gate", "review", "compatibility-center", "review-result-compatibility-center-status", "compatibility-center-redacted-status-storage-gate", mainlineReady, closedWriterReady),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusStoragePersistenceGateAuditItem("review-receipt-dry-run-result-lookup-consumer-enablement-kde-safe-runtime-diagnostics-redacted-status-storage-persistence-gate", "review", "runtime-diagnostics", "review-result-runtime-diagnostics-status", "runtime-diagnostics-redacted-status-storage-gate", mainlineReady, closedWriterReady),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusStoragePersistenceGateAuditItem("renew-receipt-dry-run-result-lookup-consumer-enablement-kde-safe-compatibility-center-redacted-status-storage-persistence-gate", "renew", "compatibility-center", "renewal-result-compatibility-center-status", "compatibility-center-redacted-status-storage-gate", mainlineReady, closedWriterReady),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusStoragePersistenceGateAuditItem("renew-receipt-dry-run-result-lookup-consumer-enablement-kde-safe-runtime-diagnostics-redacted-status-storage-persistence-gate", "renew", "runtime-diagnostics", "renewal-result-runtime-diagnostics-status", "runtime-diagnostics-redacted-status-storage-gate", mainlineReady, closedWriterReady),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusStoragePersistenceGateAuditItem("open-compatibility-center-dry-run-result-lookup-consumer-enablement-kde-safe-compatibility-center-redacted-status-storage-persistence-gate", "open-compatibility-center", "compatibility-center", "navigation-result-compatibility-center-status", "compatibility-center-redacted-status-storage-gate", mainlineReady, closedWriterReady),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusStoragePersistenceGateAuditItem("open-compatibility-center-dry-run-result-lookup-consumer-enablement-kde-safe-runtime-diagnostics-redacted-status-storage-persistence-gate", "open-compatibility-center", "runtime-diagnostics", "navigation-result-runtime-diagnostics-status", "runtime-diagnostics-redacted-status-storage-gate", mainlineReady, closedWriterReady),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusStoragePersistenceGateAuditItem("dismiss-receipt-dry-run-result-lookup-consumer-enablement-kde-safe-compatibility-center-redacted-status-storage-persistence-gate", "dismiss", "compatibility-center", "dismissal-result-compatibility-center-status", "compatibility-center-redacted-status-storage-gate", mainlineReady, closedWriterReady),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusStoragePersistenceGateAuditItem("dismiss-receipt-dry-run-result-lookup-consumer-enablement-kde-safe-runtime-diagnostics-redacted-status-storage-persistence-gate", "dismiss", "runtime-diagnostics", "dismissal-result-runtime-diagnostics-status", "runtime-diagnostics-redacted-status-storage-gate", mainlineReady, closedWriterReady),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusStoragePersistenceGateAuditItem("support-info-dry-run-result-lookup-consumer-enablement-kde-safe-compatibility-center-redacted-status-storage-persistence-gate", "support-info", "compatibility-center", "support-info-result-compatibility-center-status", "compatibility-center-redacted-status-storage-gate", mainlineReady, closedWriterReady),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusStoragePersistenceGateAuditItem("support-info-dry-run-result-lookup-consumer-enablement-kde-safe-runtime-diagnostics-redacted-status-storage-persistence-gate", "support-info", "runtime-diagnostics", "support-info-result-runtime-diagnostics-status", "runtime-diagnostics-redacted-status-storage-gate", mainlineReady, closedWriterReady),
	}
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusStoragePersistenceGateAuditItem(id string, actionKind string, surfaceKind string, statusConsumerKind string, gateScope string, mainlineReady bool, closedWriterReady bool) ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusStoragePersistenceGateAuditItem {
	ready := mainlineReady && closedWriterReady
	status := "missing-redacted-status-storage-persistence-gate-evidence"
	if ready {
		status = "redacted-status-storage-persistence-gate-modeled-storage-disabled"
	}
	return ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusStoragePersistenceGateAuditItem{
		ID:                              id,
		ActionKind:                      actionKind,
		SurfaceKind:                     surfaceKind,
		StatusConsumerKind:              statusConsumerKind,
		GateScope:                       gateScope,
		EvidencePresent:                 ready,
		CurrentMainlineConsumed:         mainlineReady,
		ClosedPersistenceWriterConsumed: closedWriterReady,
		ClosedPersistenceWriterReady:    closedWriterReady,
		StoragePersistenceGateModeled:   ready,
		GateInputBoundaryModeled:        ready,
		GateOutputBoundaryModeled:       ready,
		KDESafeRedactedStatusOnly:       ready,
		WriterCallable:                  false,
		ClosedPersistenceWriterEnabled:  false,
		StorageGatePassed:               false,
		StoragePersistenceAuthorized:    false,
		StatusPersistenceAuthorized:     false,
		StatusPersistenceWriteEnabled:   false,
		StatusWriterEnabled:             false,
		KDEStatusWriteEnabled:           false,
		RuntimeDiagnosticsWriteEnabled:  false,
		StorageWriteEnabled:             false,
		RedactedSummaryPersisted:        false,
		KDEStatusPersisted:              false,
		RuntimeDiagnosticsPersisted:     false,
		RawResultExposed:                false,
		UserVisible:                     ready,
		ReviewOnly:                      true,
		RuntimeOwned:                    true,
		GoRuntimeBacked:                 true,
		KDEPolicyOwner:                  false,
		SideEffectsDisabled:             true,
		HostRootModified:                false,
		InternalDetailsExposed:          false,
		StoragePersistenceGateStatus:    status,
		NextRequirement:                 "require durable storage policy evidence before this gate can pass",
	}
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusStoragePersistenceGateAuditChecks(preview ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusStoragePersistenceGateAuditPreview) []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusStoragePersistenceGateAuditCheck {
	return []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusStoragePersistenceGateAuditCheck{
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusStoragePersistenceGateAuditCheck("current-mainline-consumed", productionAuthorizationPassBlocked(preview.CurrentMainlineConsumed), "The storage persistence gate audit consumes the Xnix current mainline."),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusStoragePersistenceGateAuditCheck("closed-persistence-writer-consumed", productionAuthorizationPassBlocked(preview.ClosedPersistenceWriterConsumed && preview.ClosedPersistenceWriterReady), "The storage persistence gate audit consumes the closed persistence writer preview."),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusStoragePersistenceGateAuditCheck("storage-persistence-gate-modeled", productionAuthorizationPassBlocked(preview.StoragePersistenceGateRequired && preview.StoragePersistenceGateModeled && preview.StoragePersistenceGateReady && preview.GateInputBoundaryModeled && preview.GateOutputBoundaryModeled), "The redacted status storage persistence gate is modeled."),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusStoragePersistenceGateAuditCheck("compatibility-center-and-runtime-storage-gates-modeled", productionAuthorizationPassBlocked(preview.CompatibilityCenterGateModeled && preview.RuntimeDiagnosticsGateModeled && preview.CompatibilityCenterItemCount == 5 && preview.RuntimeDiagnosticsItemCount == 5), "Compatibility Center and Runtime diagnostics storage gate candidates are both modeled."),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusStoragePersistenceGateAuditCheck("ten-gate-items-ready-storage-disabled", productionAuthorizationPassBlocked(preview.GateItemCount == 10 && preview.RequiredGateItemCount == 10 && preview.ReadyGateItemCount == 10 && preview.MissingGateItemCount == 0 && preview.PassedStorageGateItemCount == 0 && preview.AuthorizedStorageItemCount == 0 && preview.CallableWriterItemCount == 0 && preview.EnabledPersistenceWriterItemCount == 0 && preview.PersistedWriterItemCount == 0), "All storage persistence gate candidates are present while storage remains disabled."),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusStoragePersistenceGateAuditCheck("storage-persistence-writes-disabled", productionAuthorizationPassBlocked(!preview.WriterCallable && !preview.ClosedPersistenceWriterEnabled && !preview.StorageGatePassed && !preview.StoragePersistenceAuthorized && !preview.WriterAuthorizationGranted && !preview.StatusWriterEnabled && !preview.StatusPersistenceAuthorized && !preview.StatusPersistenceWriteEnabled && !preview.KDEStatusWriteEnabled && !preview.RuntimeDiagnosticsWriteEnabled && !preview.StorageWriteEnabled && !preview.RedactedSummaryPersisted && !preview.KDEStatusPersisted && !preview.RuntimeDiagnosticsPersisted && !preview.DryRunResultPersisted), "Storage gates, writer calls, status writes, and redacted status persistence remain disabled."),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusStoragePersistenceGateAuditCheck("consumer-lookup-request-and-support-disabled", productionAuthorizationPassBlocked(!preview.ConsumerConsumptionAuthorized && !preview.ConsumerEnablementAuthorized && !preview.KDEConsumerEnabled && !preview.RuntimeConsumerEnabled && !preview.LookupRouteAuthorized && !preview.LookupRouteEnabled && !preview.OpaqueLookupEnabled && !preview.RequestObjectCreationEnabled && !preview.RequestObjectDispatchEnabled && !preview.PortalRequestCreated && !preview.NotificationActionEnabled && !preview.CompatibilityCenterOpened && !preview.SupportBundleExported && !preview.SupportCaseCreated), "Consumers, lookup, request objects, desktop actions, and support writes remain disabled."),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusStoragePersistenceGateAuditCheck("production-and-host-boundary-closed", productionAuthorizationPassBlocked(!preview.SystemServiceStarted && !preview.SessionBusClaimed && !preview.ProductionBusClaimed && !preview.ProductionOwnerEnabled && !preview.ProductionReadiness && !preview.ProductionOwnershipReady && !preview.WriteMethodsEnabled && !preview.RuntimeWritesEnabled && !preview.AdapterInvocationEnabled && !preview.BackendLaunchEnabled && !preview.BackendProcessStarted && !preview.NetworkRequired && !preview.HostRootModified && !preview.PrivilegedContainerRequired && !preview.CallerStateRootRequired && !preview.StateRootPathExposed && !preview.FilePathsExposed && !preview.FileContentRead && !preview.RawResultExposed && !preview.RawCommandExposed && !preview.RawExecutableExposed && !preview.BackendDetailsExposed && preview.RawExposedGateItemCount == 0 && preview.SideEffectGateItemCount == 0 && productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusStoragePersistenceGateAuditItemsKeepClosed(preview.GateItems)), "Production ownership, writes, unsafe data exposure, network, privilege, internal details, and host mutation remain disabled."),
	}
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusStoragePersistenceGateAuditCheck(id string, status string, summary string) ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusStoragePersistenceGateAuditCheck {
	return ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusStoragePersistenceGateAuditCheck{ID: id, Status: status, Summary: summary}
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusStoragePersistenceGateAuditReadyCount(items []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusStoragePersistenceGateAuditItem) int {
	count := 0
	for _, item := range items {
		if item.EvidencePresent && item.StoragePersistenceGateModeled && item.GateInputBoundaryModeled && item.GateOutputBoundaryModeled && !item.StorageGatePassed && !item.StatusPersistenceWriteEnabled && !item.StorageWriteEnabled && !item.HostRootModified && !item.InternalDetailsExposed {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusStoragePersistenceGateAuditSurfaceCount(items []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusStoragePersistenceGateAuditItem, surfaceKind string) int {
	count := 0
	for _, item := range items {
		if item.SurfaceKind == surfaceKind {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusStoragePersistenceGateAuditItemsKeepClosed(items []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusStoragePersistenceGateAuditItem) bool {
	for _, item := range items {
		if item.WriterCallable || item.ClosedPersistenceWriterEnabled || item.StorageGatePassed || item.StoragePersistenceAuthorized || item.StatusPersistenceAuthorized || item.StatusPersistenceWriteEnabled || item.StatusWriterEnabled || item.KDEStatusWriteEnabled || item.RuntimeDiagnosticsWriteEnabled || item.StorageWriteEnabled || item.RedactedSummaryPersisted || item.KDEStatusPersisted || item.RuntimeDiagnosticsPersisted || item.RawResultExposed || item.KDEPolicyOwner || !item.SideEffectsDisabled || item.HostRootModified || item.InternalDetailsExposed {
			return false
		}
	}
	return true
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusStoragePersistenceGateAuditItemIDs(items []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusStoragePersistenceGateAuditItem) []string {
	ids := make([]string, 0, len(items))
	for _, item := range items {
		ids = append(ids, item.ID)
	}
	return ids
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusStoragePersistenceGateAuditCheckIDs(checks []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusStoragePersistenceGateAuditCheck) []string {
	ids := make([]string, 0, len(checks))
	for _, check := range checks {
		ids = append(ids, check.ID)
	}
	return ids
}

func countProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusStoragePersistenceGateAuditChecks(checks []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusStoragePersistenceGateAuditCheck) ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusStoragePersistenceGateAuditCounts {
	counts := ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusStoragePersistenceGateAuditCounts{Total: len(checks)}
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
