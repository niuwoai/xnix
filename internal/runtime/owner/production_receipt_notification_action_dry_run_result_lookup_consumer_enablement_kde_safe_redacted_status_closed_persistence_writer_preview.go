package owner

type ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedPersistenceWriterPreview struct {
	Version                                string                                                                                                                     `json:"version"`
	SchemaVersion                          string                                                                                                                     `json:"schema_version"`
	RequestType                            string                                                                                                                     `json:"request_type"`
	PreviewType                            string                                                                                                                     `json:"preview_type"`
	Source                                 string                                                                                                                     `json:"source"`
	PreviewDecision                        string                                                                                                                     `json:"preview_decision"`
	CurrentMainlineConsumed                bool                                                                                                                       `json:"current_mainline_consumed"`
	WriterPersistenceAuthorizationConsumed bool                                                                                                                       `json:"writer_persistence_authorization_consumed"`
	WriterPersistenceAuthorizationReady    bool                                                                                                                       `json:"writer_persistence_authorization_ready"`
	ClosedPersistenceWriterRequired        bool                                                                                                                       `json:"closed_persistence_writer_required"`
	ClosedPersistenceWriterModeled         bool                                                                                                                       `json:"closed_persistence_writer_modeled"`
	ClosedPersistenceWriterReady           bool                                                                                                                       `json:"closed_persistence_writer_ready"`
	StorageWriterShapeModeled              bool                                                                                                                       `json:"storage_writer_shape_modeled"`
	StorageInputBoundaryModeled            bool                                                                                                                       `json:"storage_input_boundary_modeled"`
	StorageOutputBoundaryModeled           bool                                                                                                                       `json:"storage_output_boundary_modeled"`
	KDESafeRedactedStatusOnly              bool                                                                                                                       `json:"kde_safe_redacted_status_only"`
	CompatibilityCenterStorageModeled      bool                                                                                                                       `json:"compatibility_center_storage_modeled"`
	RuntimeDiagnosticsStorageModeled       bool                                                                                                                       `json:"runtime_diagnostics_storage_modeled"`
	WriterCallable                         bool                                                                                                                       `json:"writer_callable"`
	ClosedPersistenceWriterEnabled         bool                                                                                                                       `json:"closed_persistence_writer_enabled"`
	WriterAuthorizationGranted             bool                                                                                                                       `json:"writer_authorization_granted"`
	StatusWriterEnabled                    bool                                                                                                                       `json:"status_writer_enabled"`
	StatusPersistenceAuthorized            bool                                                                                                                       `json:"status_persistence_authorized"`
	StatusPersistenceWriteEnabled          bool                                                                                                                       `json:"status_persistence_write_enabled"`
	KDEStatusWriteEnabled                  bool                                                                                                                       `json:"kde_status_write_enabled"`
	RuntimeDiagnosticsWriteEnabled         bool                                                                                                                       `json:"runtime_diagnostics_write_enabled"`
	StorageWriteEnabled                    bool                                                                                                                       `json:"storage_write_enabled"`
	WriterItemCount                        int                                                                                                                        `json:"writer_item_count"`
	RequiredWriterItemCount                int                                                                                                                        `json:"required_writer_item_count"`
	ReadyWriterItemCount                   int                                                                                                                        `json:"ready_writer_item_count"`
	MissingWriterItemCount                 int                                                                                                                        `json:"missing_writer_item_count"`
	CallableWriterItemCount                int                                                                                                                        `json:"callable_writer_item_count"`
	EnabledPersistenceWriterItemCount      int                                                                                                                        `json:"enabled_persistence_writer_item_count"`
	PersistedWriterItemCount               int                                                                                                                        `json:"persisted_writer_item_count"`
	RawExposedWriterItemCount              int                                                                                                                        `json:"raw_exposed_writer_item_count"`
	SideEffectWriterItemCount              int                                                                                                                        `json:"side_effect_writer_item_count"`
	CompatibilityCenterItemCount           int                                                                                                                        `json:"compatibility_center_item_count"`
	RuntimeDiagnosticsItemCount            int                                                                                                                        `json:"runtime_diagnostics_item_count"`
	WriterItems                            []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedPersistenceWriterItem  `json:"writer_items"`
	WriterItemIDs                          []string                                                                                                                   `json:"writer_item_ids"`
	Checks                                 []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedPersistenceWriterCheck `json:"checks"`
	CheckIDs                               []string                                                                                                                   `json:"check_ids"`
	Counts                                 ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedPersistenceWriterCounts  `json:"counts"`
	ConsumerConsumptionAuthorized          bool                                                                                                                       `json:"consumer_consumption_authorized"`
	ConsumerEnablementAuthorized           bool                                                                                                                       `json:"consumer_enablement_authorized"`
	KDEConsumerEnabled                     bool                                                                                                                       `json:"kde_consumer_enabled"`
	RuntimeConsumerEnabled                 bool                                                                                                                       `json:"runtime_consumer_enabled"`
	LookupRouteAuthorized                  bool                                                                                                                       `json:"lookup_route_authorized"`
	LookupRouteEnabled                     bool                                                                                                                       `json:"lookup_route_enabled"`
	OpaqueLookupEnabled                    bool                                                                                                                       `json:"opaque_lookup_enabled"`
	RedactedSummaryPersisted               bool                                                                                                                       `json:"redacted_summary_persisted"`
	KDEStatusPersisted                     bool                                                                                                                       `json:"kde_status_persisted"`
	RuntimeDiagnosticsPersisted            bool                                                                                                                       `json:"runtime_diagnostics_persisted"`
	DryRunResultPersisted                  bool                                                                                                                       `json:"dry_run_result_persisted"`
	RawResultExposed                       bool                                                                                                                       `json:"raw_result_exposed"`
	DispatchDryRunExecuted                 bool                                                                                                                       `json:"dispatch_dry_run_executed"`
	RequestObjectCreationEnabled           bool                                                                                                                       `json:"request_object_creation_enabled"`
	RequestObjectDispatchEnabled           bool                                                                                                                       `json:"request_object_dispatch_enabled"`
	PortalRequestCreated                   bool                                                                                                                       `json:"portal_request_created"`
	NotificationActionEnabled              bool                                                                                                                       `json:"notification_action_enabled"`
	CompatibilityCenterOpened              bool                                                                                                                       `json:"compatibility_center_opened"`
	SupportBundleExported                  bool                                                                                                                       `json:"support_bundle_exported"`
	SupportCaseCreated                     bool                                                                                                                       `json:"support_case_created"`
	RuntimeOwned                           bool                                                                                                                       `json:"runtime_owned"`
	GoRuntimeBacked                        bool                                                                                                                       `json:"go_runtime_backed"`
	KDEPolicyOwner                         bool                                                                                                                       `json:"kde_policy_owner"`
	ProductionReadiness                    bool                                                                                                                       `json:"production_readiness"`
	ProductionOwnershipReady               bool                                                                                                                       `json:"production_ownership_ready"`
	SystemServiceStarted                   bool                                                                                                                       `json:"system_service_started"`
	SessionBusClaimed                      bool                                                                                                                       `json:"session_bus_claimed"`
	ProductionBusClaimed                   bool                                                                                                                       `json:"production_bus_claimed"`
	ProductionOwnerEnabled                 bool                                                                                                                       `json:"production_owner_enabled"`
	WriteMethodsEnabled                    bool                                                                                                                       `json:"write_methods_enabled"`
	RuntimeWritesEnabled                   bool                                                                                                                       `json:"runtime_writes_enabled"`
	DesktopFilesWritten                    bool                                                                                                                       `json:"desktop_files_written"`
	SettingsPersisted                      bool                                                                                                                       `json:"settings_persisted"`
	AdapterInvocationEnabled               bool                                                                                                                       `json:"adapter_invocation_enabled"`
	BackendLaunchEnabled                   bool                                                                                                                       `json:"backend_launch_enabled"`
	BackendProcessStarted                  bool                                                                                                                       `json:"backend_process_started"`
	SnapshotRestoreExecuted                bool                                                                                                                       `json:"snapshot_restore_executed"`
	StateCleanupExecuted                   bool                                                                                                                       `json:"state_cleanup_executed"`
	NetworkRequired                        bool                                                                                                                       `json:"network_required"`
	HostRootModified                       bool                                                                                                                       `json:"host_root_modified"`
	PrivilegedContainerRequired            bool                                                                                                                       `json:"privileged_container_required"`
	CallerStateRootRequired                bool                                                                                                                       `json:"caller_state_root_required"`
	StateRootPathExposed                   bool                                                                                                                       `json:"state_root_path_exposed"`
	FilePathsExposed                       bool                                                                                                                       `json:"file_paths_exposed"`
	FileContentRead                        bool                                                                                                                       `json:"file_content_read"`
	RawCommandExposed                      bool                                                                                                                       `json:"raw_command_exposed"`
	RawExecutableExposed                   bool                                                                                                                       `json:"raw_executable_exposed"`
	BackendDetailsExposed                  bool                                                                                                                       `json:"backend_details_exposed"`
	BlockedActions                         []string                                                                                                                   `json:"blocked_actions"`
	NextRequirements                       []string                                                                                                                   `json:"next_requirements"`
	DesktopSafeSummary                     string                                                                                                                     `json:"desktop_safe_summary"`
}

type ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedPersistenceWriterItem struct {
	ID                                  string `json:"id"`
	ActionKind                          string `json:"action_kind"`
	SurfaceKind                         string `json:"surface_kind"`
	StatusConsumerKind                  string `json:"status_consumer_kind"`
	StorageWriterScope                  string `json:"storage_writer_scope"`
	EvidencePresent                     bool   `json:"evidence_present"`
	CurrentMainlineConsumed             bool   `json:"current_mainline_consumed"`
	WriterPersistenceAuthorizationReady bool   `json:"writer_persistence_authorization_ready"`
	ClosedPersistenceWriterModeled      bool   `json:"closed_persistence_writer_modeled"`
	StorageWriterShapeModeled           bool   `json:"storage_writer_shape_modeled"`
	StorageInputBoundaryModeled         bool   `json:"storage_input_boundary_modeled"`
	StorageOutputBoundaryModeled        bool   `json:"storage_output_boundary_modeled"`
	KDESafeRedactedStatusOnly           bool   `json:"kde_safe_redacted_status_only"`
	WriterCallable                      bool   `json:"writer_callable"`
	ClosedPersistenceWriterEnabled      bool   `json:"closed_persistence_writer_enabled"`
	StatusPersistenceAuthorized         bool   `json:"status_persistence_authorized"`
	StatusPersistenceWriteEnabled       bool   `json:"status_persistence_write_enabled"`
	StatusWriterEnabled                 bool   `json:"status_writer_enabled"`
	KDEStatusWriteEnabled               bool   `json:"kde_status_write_enabled"`
	RuntimeDiagnosticsWriteEnabled      bool   `json:"runtime_diagnostics_write_enabled"`
	StorageWriteEnabled                 bool   `json:"storage_write_enabled"`
	RedactedSummaryPersisted            bool   `json:"redacted_summary_persisted"`
	KDEStatusPersisted                  bool   `json:"kde_status_persisted"`
	RuntimeDiagnosticsPersisted         bool   `json:"runtime_diagnostics_persisted"`
	RawResultExposed                    bool   `json:"raw_result_exposed"`
	UserVisible                         bool   `json:"user_visible"`
	ReviewOnly                          bool   `json:"review_only"`
	RuntimeOwned                        bool   `json:"runtime_owned"`
	GoRuntimeBacked                     bool   `json:"go_runtime_backed"`
	KDEPolicyOwner                      bool   `json:"kde_policy_owner"`
	SideEffectsDisabled                 bool   `json:"side_effects_disabled"`
	HostRootModified                    bool   `json:"host_root_modified"`
	InternalDetailsExposed              bool   `json:"internal_details_exposed"`
	ClosedPersistenceWriterStatus       string `json:"closed_persistence_writer_status"`
	NextRequirement                     string `json:"next_requirement"`
}

type ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedPersistenceWriterCheck struct {
	ID      string `json:"id"`
	Status  string `json:"status"`
	Summary string `json:"summary"`
}

type ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedPersistenceWriterCounts struct {
	Total   int `json:"total"`
	Passed  int `json:"passed"`
	Pending int `json:"pending"`
	Blocked int `json:"blocked"`
}

func NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedPersistenceWriterPreview(root string) (ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedPersistenceWriterPreview, error) {
	version, err := readProductionDBusGateReviewVersion(root)
	if err != nil {
		return ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedPersistenceWriterPreview{}, err
	}
	sources := productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedPersistenceWriterSources(root)
	mainlineReady := productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedPersistenceWriterMainlineReady(sources.CurrentMainline)
	authorizationReady := productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedPersistenceWriterAuthorizationReady(sources.WriterPersistenceAuthorization)
	items := productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedPersistenceWriterItems(mainlineReady, authorizationReady)
	readyCount := productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedPersistenceWriterReadyCount(items)
	preview := ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedPersistenceWriterPreview{
		Version:                                version,
		SchemaVersion:                          "xnix.runtime.production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_kde_safe_redacted_status_closed_persistence_writer.v1",
		RequestType:                            "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-persistence-writer-preview",
		PreviewType:                            "receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-persistence-writer",
		Source:                                 "xnix-current-mainline+redacted-status-writer-persistence-authorization-preview",
		PreviewDecision:                        "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-persistence-writer-blocked",
		CurrentMainlineConsumed:                mainlineReady,
		WriterPersistenceAuthorizationConsumed: authorizationReady,
		WriterPersistenceAuthorizationReady:    authorizationReady,
		ClosedPersistenceWriterRequired:        true,
		ClosedPersistenceWriterModeled:         mainlineReady && authorizationReady,
		ClosedPersistenceWriterReady:           mainlineReady && authorizationReady,
		StorageWriterShapeModeled:              mainlineReady && authorizationReady,
		StorageInputBoundaryModeled:            mainlineReady && authorizationReady,
		StorageOutputBoundaryModeled:           mainlineReady && authorizationReady,
		KDESafeRedactedStatusOnly:              mainlineReady && authorizationReady,
		CompatibilityCenterStorageModeled:      mainlineReady && authorizationReady,
		RuntimeDiagnosticsStorageModeled:       mainlineReady && authorizationReady,
		WriterCallable:                         false,
		ClosedPersistenceWriterEnabled:         false,
		WriterAuthorizationGranted:             false,
		StatusWriterEnabled:                    false,
		StatusPersistenceAuthorized:            false,
		StatusPersistenceWriteEnabled:          false,
		KDEStatusWriteEnabled:                  false,
		RuntimeDiagnosticsWriteEnabled:         false,
		StorageWriteEnabled:                    false,
		WriterItemCount:                        len(items),
		RequiredWriterItemCount:                len(items),
		ReadyWriterItemCount:                   readyCount,
		MissingWriterItemCount:                 len(items) - readyCount,
		CallableWriterItemCount:                0,
		EnabledPersistenceWriterItemCount:      0,
		PersistedWriterItemCount:               0,
		RawExposedWriterItemCount:              0,
		SideEffectWriterItemCount:              0,
		CompatibilityCenterItemCount:           productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedPersistenceWriterSurfaceCount(items, "compatibility-center"),
		RuntimeDiagnosticsItemCount:            productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedPersistenceWriterSurfaceCount(items, "runtime-diagnostics"),
		WriterItems:                            items,
		WriterItemIDs:                          productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedPersistenceWriterItemIDs(items),
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
			"do not call closed persistence writers",
			"do not persist redacted status summaries",
			"do not enable KDE or Runtime diagnostics consumers",
			"do not create lookup, dispatch, request, Portal, notification, navigation, or support side effects",
			"do not claim production D-Bus ownership, launch engines, expose paths, or mutate the host",
		},
		NextRequirements: []string{
			"add a storage persistence gate before any redacted status can be written",
			"add a consumer read model before KDE surfaces can consume persisted status",
			"keep the closed persistence writer modeled but disabled until explicit production authorization exists",
		},
		DesktopSafeSummary: "The KDE-safe redacted status closed persistence writer preview consumes the current mainline and writer persistence authorization preview, models the redacted status storage writer shape, and keeps writer calls, persistence writes, consumers, desktop actions, unsafe data exposure, and host mutation disabled.",
	}
	checks := productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedPersistenceWriterChecks(preview)
	preview.Checks = checks
	preview.CheckIDs = productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedPersistenceWriterCheckIDs(checks)
	preview.Counts = countProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedPersistenceWriterChecks(checks)
	if preview.Counts.Blocked == 0 {
		preview.PreviewDecision = "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-persistence-writer-ready-storage-disabled"
	}
	if err := validateNoBackendTerms(preview, "production receipt notification action dry-run result lookup consumer enablement kde safe redacted status closed persistence writer preview"); err != nil {
		return ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedPersistenceWriterPreview{}, err
	}
	return preview, nil
}

type productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedPersistenceWriterSourceSet struct {
	CurrentMainline                string
	WriterPersistenceAuthorization string
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedPersistenceWriterSources(root string) productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedPersistenceWriterSourceSet {
	return productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedPersistenceWriterSourceSet{
		CurrentMainline:                productionAuthorizationReadSources(root, []string{"docs/xnix-current-mainline.md"}),
		WriterPersistenceAuthorization: productionAuthorizationReadSources(root, []string{"internal/runtime/owner/production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_kde_safe_redacted_status_writer_persistence_authorization_preview.go", "internal/runtime/owner/production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_kde_safe_redacted_status_writer_persistence_authorization_preview_test.go"}),
	}
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedPersistenceWriterMainlineReady(source string) bool {
	return productionAuthorizationHasAll(source, []string{"Xnix Current Mainline", "redacted status closed persistence writer preview", "redacted status writer persistence authorization preview", "No Runtime write methods", "No KDE configuration writes"})
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedPersistenceWriterAuthorizationReady(source string) bool {
	return productionAuthorizationHasAll(source, []string{"production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-writer-persistence-authorization-preview", "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-writer-persistence-authorization-ready-persistence-disabled", "PersistenceAuthorizationReady", "WriterPersistenceBoundaryReady", "StatusPersistenceAuthorized", "StatusPersistenceWriteEnabled"})
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedPersistenceWriterItems(mainlineReady bool, authorizationReady bool) []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedPersistenceWriterItem {
	return []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedPersistenceWriterItem{
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedPersistenceWriterItem("review-receipt-dry-run-result-lookup-consumer-enablement-kde-safe-compatibility-center-redacted-status-closed-persistence-writer", "review", "compatibility-center", "review-result-compatibility-center-status", "compatibility-center-redacted-status-storage", mainlineReady, authorizationReady),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedPersistenceWriterItem("review-receipt-dry-run-result-lookup-consumer-enablement-kde-safe-runtime-diagnostics-redacted-status-closed-persistence-writer", "review", "runtime-diagnostics", "review-result-runtime-diagnostics-status", "runtime-diagnostics-redacted-status-storage", mainlineReady, authorizationReady),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedPersistenceWriterItem("renew-receipt-dry-run-result-lookup-consumer-enablement-kde-safe-compatibility-center-redacted-status-closed-persistence-writer", "renew", "compatibility-center", "renewal-result-compatibility-center-status", "compatibility-center-redacted-status-storage", mainlineReady, authorizationReady),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedPersistenceWriterItem("renew-receipt-dry-run-result-lookup-consumer-enablement-kde-safe-runtime-diagnostics-redacted-status-closed-persistence-writer", "renew", "runtime-diagnostics", "renewal-result-runtime-diagnostics-status", "runtime-diagnostics-redacted-status-storage", mainlineReady, authorizationReady),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedPersistenceWriterItem("open-compatibility-center-dry-run-result-lookup-consumer-enablement-kde-safe-compatibility-center-redacted-status-closed-persistence-writer", "open-compatibility-center", "compatibility-center", "navigation-result-compatibility-center-status", "compatibility-center-redacted-status-storage", mainlineReady, authorizationReady),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedPersistenceWriterItem("open-compatibility-center-dry-run-result-lookup-consumer-enablement-kde-safe-runtime-diagnostics-redacted-status-closed-persistence-writer", "open-compatibility-center", "runtime-diagnostics", "navigation-result-runtime-diagnostics-status", "runtime-diagnostics-redacted-status-storage", mainlineReady, authorizationReady),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedPersistenceWriterItem("dismiss-receipt-dry-run-result-lookup-consumer-enablement-kde-safe-compatibility-center-redacted-status-closed-persistence-writer", "dismiss", "compatibility-center", "dismissal-result-compatibility-center-status", "compatibility-center-redacted-status-storage", mainlineReady, authorizationReady),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedPersistenceWriterItem("dismiss-receipt-dry-run-result-lookup-consumer-enablement-kde-safe-runtime-diagnostics-redacted-status-closed-persistence-writer", "dismiss", "runtime-diagnostics", "dismissal-result-runtime-diagnostics-status", "runtime-diagnostics-redacted-status-storage", mainlineReady, authorizationReady),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedPersistenceWriterItem("support-info-dry-run-result-lookup-consumer-enablement-kde-safe-compatibility-center-redacted-status-closed-persistence-writer", "support-info", "compatibility-center", "support-info-result-compatibility-center-status", "compatibility-center-redacted-status-storage", mainlineReady, authorizationReady),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedPersistenceWriterItem("support-info-dry-run-result-lookup-consumer-enablement-kde-safe-runtime-diagnostics-redacted-status-closed-persistence-writer", "support-info", "runtime-diagnostics", "support-info-result-runtime-diagnostics-status", "runtime-diagnostics-redacted-status-storage", mainlineReady, authorizationReady),
	}
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedPersistenceWriterItem(id string, actionKind string, surfaceKind string, statusConsumerKind string, storageWriterScope string, mainlineReady bool, authorizationReady bool) ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedPersistenceWriterItem {
	ready := mainlineReady && authorizationReady
	status := "missing-redacted-status-closed-persistence-writer-evidence"
	if ready {
		status = "redacted-status-closed-persistence-writer-modeled-storage-disabled"
	}
	return ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedPersistenceWriterItem{
		ID:                                  id,
		ActionKind:                          actionKind,
		SurfaceKind:                         surfaceKind,
		StatusConsumerKind:                  statusConsumerKind,
		StorageWriterScope:                  storageWriterScope,
		EvidencePresent:                     ready,
		CurrentMainlineConsumed:             mainlineReady,
		WriterPersistenceAuthorizationReady: authorizationReady,
		ClosedPersistenceWriterModeled:      ready,
		StorageWriterShapeModeled:           ready,
		StorageInputBoundaryModeled:         ready,
		StorageOutputBoundaryModeled:        ready,
		KDESafeRedactedStatusOnly:           ready,
		WriterCallable:                      false,
		ClosedPersistenceWriterEnabled:      false,
		StatusPersistenceAuthorized:         false,
		StatusPersistenceWriteEnabled:       false,
		StatusWriterEnabled:                 false,
		KDEStatusWriteEnabled:               false,
		RuntimeDiagnosticsWriteEnabled:      false,
		StorageWriteEnabled:                 false,
		RedactedSummaryPersisted:            false,
		KDEStatusPersisted:                  false,
		RuntimeDiagnosticsPersisted:         false,
		RawResultExposed:                    false,
		UserVisible:                         ready,
		ReviewOnly:                          true,
		RuntimeOwned:                        true,
		GoRuntimeBacked:                     true,
		KDEPolicyOwner:                      false,
		SideEffectsDisabled:                 true,
		HostRootModified:                    false,
		InternalDetailsExposed:              false,
		ClosedPersistenceWriterStatus:       status,
		NextRequirement:                     "require storage persistence gate before the closed persistence writer can write redacted status",
	}
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedPersistenceWriterChecks(preview ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedPersistenceWriterPreview) []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedPersistenceWriterCheck {
	return []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedPersistenceWriterCheck{
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedPersistenceWriterCheck("current-mainline-consumed", productionAuthorizationPassBlocked(preview.CurrentMainlineConsumed), "The closed persistence writer preview consumes the Xnix current mainline."),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedPersistenceWriterCheck("writer-persistence-authorization-consumed", productionAuthorizationPassBlocked(preview.WriterPersistenceAuthorizationConsumed && preview.WriterPersistenceAuthorizationReady), "The closed persistence writer preview consumes the writer persistence authorization preview."),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedPersistenceWriterCheck("closed-persistence-writer-shape-modeled", productionAuthorizationPassBlocked(preview.ClosedPersistenceWriterRequired && preview.ClosedPersistenceWriterModeled && preview.ClosedPersistenceWriterReady && preview.StorageWriterShapeModeled && preview.StorageInputBoundaryModeled && preview.StorageOutputBoundaryModeled), "The closed redacted status storage writer shape is modeled."),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedPersistenceWriterCheck("compatibility-center-and-runtime-storage-writers-modeled", productionAuthorizationPassBlocked(preview.CompatibilityCenterStorageModeled && preview.RuntimeDiagnosticsStorageModeled && preview.CompatibilityCenterItemCount == 5 && preview.RuntimeDiagnosticsItemCount == 5), "Compatibility Center and Runtime diagnostics storage writer candidates are both modeled."),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedPersistenceWriterCheck("ten-writer-items-ready-storage-disabled", productionAuthorizationPassBlocked(preview.WriterItemCount == 10 && preview.RequiredWriterItemCount == 10 && preview.ReadyWriterItemCount == 10 && preview.MissingWriterItemCount == 0 && preview.CallableWriterItemCount == 0 && preview.EnabledPersistenceWriterItemCount == 0 && preview.PersistedWriterItemCount == 0), "All closed persistence writer candidates are present while storage remains disabled."),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedPersistenceWriterCheck("closed-persistence-writes-disabled", productionAuthorizationPassBlocked(!preview.WriterCallable && !preview.ClosedPersistenceWriterEnabled && !preview.WriterAuthorizationGranted && !preview.StatusWriterEnabled && !preview.StatusPersistenceAuthorized && !preview.StatusPersistenceWriteEnabled && !preview.KDEStatusWriteEnabled && !preview.RuntimeDiagnosticsWriteEnabled && !preview.StorageWriteEnabled && !preview.RedactedSummaryPersisted && !preview.KDEStatusPersisted && !preview.RuntimeDiagnosticsPersisted && !preview.DryRunResultPersisted), "Writer calls, status writes, and redacted status persistence remain disabled."),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedPersistenceWriterCheck("consumer-lookup-request-and-support-disabled", productionAuthorizationPassBlocked(!preview.ConsumerConsumptionAuthorized && !preview.ConsumerEnablementAuthorized && !preview.KDEConsumerEnabled && !preview.RuntimeConsumerEnabled && !preview.LookupRouteAuthorized && !preview.LookupRouteEnabled && !preview.OpaqueLookupEnabled && !preview.RequestObjectCreationEnabled && !preview.RequestObjectDispatchEnabled && !preview.PortalRequestCreated && !preview.NotificationActionEnabled && !preview.CompatibilityCenterOpened && !preview.SupportBundleExported && !preview.SupportCaseCreated), "Consumers, lookup, request objects, desktop actions, and support writes remain disabled."),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedPersistenceWriterCheck("production-and-host-boundary-closed", productionAuthorizationPassBlocked(!preview.SystemServiceStarted && !preview.SessionBusClaimed && !preview.ProductionBusClaimed && !preview.ProductionOwnerEnabled && !preview.ProductionReadiness && !preview.ProductionOwnershipReady && !preview.WriteMethodsEnabled && !preview.RuntimeWritesEnabled && !preview.AdapterInvocationEnabled && !preview.BackendLaunchEnabled && !preview.BackendProcessStarted && !preview.NetworkRequired && !preview.HostRootModified && !preview.PrivilegedContainerRequired && !preview.CallerStateRootRequired && !preview.StateRootPathExposed && !preview.FilePathsExposed && !preview.FileContentRead && !preview.RawResultExposed && !preview.RawCommandExposed && !preview.RawExecutableExposed && !preview.BackendDetailsExposed && preview.RawExposedWriterItemCount == 0 && preview.SideEffectWriterItemCount == 0 && productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedPersistenceWriterItemsKeepClosed(preview.WriterItems)), "Production ownership, writes, unsafe data exposure, network, privilege, internal details, and host mutation remain disabled."),
	}
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedPersistenceWriterCheck(id string, status string, summary string) ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedPersistenceWriterCheck {
	return ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedPersistenceWriterCheck{ID: id, Status: status, Summary: summary}
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedPersistenceWriterReadyCount(items []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedPersistenceWriterItem) int {
	count := 0
	for _, item := range items {
		if item.EvidencePresent && item.ClosedPersistenceWriterModeled && item.StorageWriterShapeModeled && !item.StatusPersistenceWriteEnabled && !item.StorageWriteEnabled && !item.HostRootModified && !item.InternalDetailsExposed {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedPersistenceWriterSurfaceCount(items []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedPersistenceWriterItem, surfaceKind string) int {
	count := 0
	for _, item := range items {
		if item.SurfaceKind == surfaceKind {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedPersistenceWriterItemsKeepClosed(items []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedPersistenceWriterItem) bool {
	for _, item := range items {
		if item.WriterCallable || item.ClosedPersistenceWriterEnabled || item.StatusPersistenceAuthorized || item.StatusPersistenceWriteEnabled || item.StatusWriterEnabled || item.KDEStatusWriteEnabled || item.RuntimeDiagnosticsWriteEnabled || item.StorageWriteEnabled || item.RedactedSummaryPersisted || item.KDEStatusPersisted || item.RuntimeDiagnosticsPersisted || item.RawResultExposed || item.KDEPolicyOwner || !item.SideEffectsDisabled || item.HostRootModified || item.InternalDetailsExposed {
			return false
		}
	}
	return true
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedPersistenceWriterItemIDs(items []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedPersistenceWriterItem) []string {
	ids := make([]string, 0, len(items))
	for _, item := range items {
		ids = append(ids, item.ID)
	}
	return ids
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedPersistenceWriterCheckIDs(checks []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedPersistenceWriterCheck) []string {
	ids := make([]string, 0, len(checks))
	for _, check := range checks {
		ids = append(ids, check.ID)
	}
	return ids
}

func countProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedPersistenceWriterChecks(checks []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedPersistenceWriterCheck) ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedPersistenceWriterCounts {
	counts := ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedPersistenceWriterCounts{Total: len(checks)}
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
