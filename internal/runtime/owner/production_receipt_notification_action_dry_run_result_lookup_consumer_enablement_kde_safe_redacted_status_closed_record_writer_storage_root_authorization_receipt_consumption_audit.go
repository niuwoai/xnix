package owner

type ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptConsumptionAuditPreview struct {
	Version                                 string                                                                                                                                                          `json:"version"`
	SchemaVersion                           string                                                                                                                                                          `json:"schema_version"`
	RequestType                             string                                                                                                                                                          `json:"request_type"`
	PreviewType                             string                                                                                                                                                          `json:"preview_type"`
	Source                                  string                                                                                                                                                          `json:"source"`
	PreviewDecision                         string                                                                                                                                                          `json:"preview_decision"`
	CurrentMainlineConsumed                 bool                                                                                                                                                            `json:"current_mainline_consumed"`
	StorageRootPolicyConsumed               bool                                                                                                                                                            `json:"storage_root_policy_consumed"`
	StorageRootPolicyReady                  bool                                                                                                                                                            `json:"storage_root_policy_ready"`
	AuthorizationReceiptConsumptionRequired bool                                                                                                                                                            `json:"authorization_receipt_consumption_required"`
	AuthorizationReceiptConsumptionModeled  bool                                                                                                                                                            `json:"authorization_receipt_consumption_modeled"`
	AuthorizationReceiptConsumptionReady    bool                                                                                                                                                            `json:"authorization_receipt_consumption_ready"`
	AcceptedReceiptBoundaryModeled          bool                                                                                                                                                            `json:"accepted_receipt_boundary_modeled"`
	ReceiptScopeBoundaryModeled             bool                                                                                                                                                            `json:"receipt_scope_boundary_modeled"`
	StorageRootPolicyGrantBoundaryModeled   bool                                                                                                                                                            `json:"storage_root_policy_grant_boundary_modeled"`
	RecordWriterCallGrantBoundaryModeled    bool                                                                                                                                                            `json:"record_writer_call_grant_boundary_modeled"`
	KDESafeRedactedStatusOnly               bool                                                                                                                                                            `json:"kde_safe_redacted_status_only"`
	CompatibilityCenterConsumptionModeled   bool                                                                                                                                                            `json:"compatibility_center_consumption_modeled"`
	RuntimeDiagnosticsConsumptionModeled    bool                                                                                                                                                            `json:"runtime_diagnostics_consumption_modeled"`
	ReceiptPresent                          bool                                                                                                                                                            `json:"receipt_present"`
	ReceiptAccepted                         bool                                                                                                                                                            `json:"receipt_accepted"`
	ReceiptConsumed                         bool                                                                                                                                                            `json:"receipt_consumed"`
	StorageRootPolicyGrantAuthorized        bool                                                                                                                                                            `json:"storage_root_policy_grant_authorized"`
	RecordWriterCallAuthorized              bool                                                                                                                                                            `json:"record_writer_call_authorized"`
	WriterCallable                          bool                                                                                                                                                            `json:"writer_callable"`
	ClosedRecordWriterEnabled               bool                                                                                                                                                            `json:"closed_record_writer_enabled"`
	StorageRootResolved                     bool                                                                                                                                                            `json:"storage_root_resolved"`
	StorageRootCreated                      bool                                                                                                                                                            `json:"storage_root_created"`
	StorageRootMounted                      bool                                                                                                                                                            `json:"storage_root_mounted"`
	StorageRootOwnershipGranted             bool                                                                                                                                                            `json:"storage_root_ownership_granted"`
	StorageRootNamespaceGranted             bool                                                                                                                                                            `json:"storage_root_namespace_granted"`
	StorageRetentionEnforced                bool                                                                                                                                                            `json:"storage_retention_enforced"`
	StorageRedactionEnforced                bool                                                                                                                                                            `json:"storage_redaction_enforced"`
	StorageGatePassed                       bool                                                                                                                                                            `json:"storage_gate_passed"`
	StoragePersistenceAuthorized            bool                                                                                                                                                            `json:"storage_persistence_authorized"`
	DurableRecordWriteEnabled               bool                                                                                                                                                            `json:"durable_record_write_enabled"`
	WriterAuthorizationGranted              bool                                                                                                                                                            `json:"writer_authorization_granted"`
	StatusWriterEnabled                     bool                                                                                                                                                            `json:"status_writer_enabled"`
	StatusPersistenceAuthorized             bool                                                                                                                                                            `json:"status_persistence_authorized"`
	StatusPersistenceWriteEnabled           bool                                                                                                                                                            `json:"status_persistence_write_enabled"`
	KDEStatusWriteEnabled                   bool                                                                                                                                                            `json:"kde_status_write_enabled"`
	RuntimeDiagnosticsWriteEnabled          bool                                                                                                                                                            `json:"runtime_diagnostics_write_enabled"`
	StorageWriteEnabled                     bool                                                                                                                                                            `json:"storage_write_enabled"`
	ConsumptionItemCount                    int                                                                                                                                                             `json:"consumption_item_count"`
	RequiredConsumptionItemCount            int                                                                                                                                                             `json:"required_consumption_item_count"`
	ReadyConsumptionItemCount               int                                                                                                                                                             `json:"ready_consumption_item_count"`
	MissingConsumptionItemCount             int                                                                                                                                                             `json:"missing_consumption_item_count"`
	AcceptedReceiptItemCount                int                                                                                                                                                             `json:"accepted_receipt_item_count"`
	ConsumedReceiptItemCount                int                                                                                                                                                             `json:"consumed_receipt_item_count"`
	AuthorizedStorageRootGrantItemCount     int                                                                                                                                                             `json:"authorized_storage_root_grant_item_count"`
	AuthorizedRecordWriterCallItemCount     int                                                                                                                                                             `json:"authorized_record_writer_call_item_count"`
	CallableWriterItemCount                 int                                                                                                                                                             `json:"callable_writer_item_count"`
	EnabledRecordWriterItemCount            int                                                                                                                                                             `json:"enabled_record_writer_item_count"`
	DurableWrittenRecordItemCount           int                                                                                                                                                             `json:"durable_written_record_item_count"`
	RawExposedConsumptionItemCount          int                                                                                                                                                             `json:"raw_exposed_consumption_item_count"`
	SideEffectConsumptionItemCount          int                                                                                                                                                             `json:"side_effect_consumption_item_count"`
	CompatibilityCenterItemCount            int                                                                                                                                                             `json:"compatibility_center_item_count"`
	RuntimeDiagnosticsItemCount             int                                                                                                                                                             `json:"runtime_diagnostics_item_count"`
	ConsumptionItems                        []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptConsumptionItem  `json:"consumption_items"`
	ConsumptionItemIDs                      []string                                                                                                                                                        `json:"consumption_item_ids"`
	Checks                                  []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptConsumptionCheck `json:"checks"`
	CheckIDs                                []string                                                                                                                                                        `json:"check_ids"`
	Counts                                  ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptConsumptionCounts  `json:"counts"`
	ConsumerConsumptionAuthorized           bool                                                                                                                                                            `json:"consumer_consumption_authorized"`
	ConsumerEnablementAuthorized            bool                                                                                                                                                            `json:"consumer_enablement_authorized"`
	KDEConsumerEnabled                      bool                                                                                                                                                            `json:"kde_consumer_enabled"`
	RuntimeConsumerEnabled                  bool                                                                                                                                                            `json:"runtime_consumer_enabled"`
	LookupRouteAuthorized                   bool                                                                                                                                                            `json:"lookup_route_authorized"`
	LookupRouteEnabled                      bool                                                                                                                                                            `json:"lookup_route_enabled"`
	OpaqueLookupEnabled                     bool                                                                                                                                                            `json:"opaque_lookup_enabled"`
	RedactedSummaryPersisted                bool                                                                                                                                                            `json:"redacted_summary_persisted"`
	KDEStatusPersisted                      bool                                                                                                                                                            `json:"kde_status_persisted"`
	RuntimeDiagnosticsPersisted             bool                                                                                                                                                            `json:"runtime_diagnostics_persisted"`
	DryRunResultPersisted                   bool                                                                                                                                                            `json:"dry_run_result_persisted"`
	RawResultExposed                        bool                                                                                                                                                            `json:"raw_result_exposed"`
	DispatchDryRunExecuted                  bool                                                                                                                                                            `json:"dispatch_dry_run_executed"`
	RequestObjectCreationEnabled            bool                                                                                                                                                            `json:"request_object_creation_enabled"`
	RequestObjectDispatchEnabled            bool                                                                                                                                                            `json:"request_object_dispatch_enabled"`
	PortalRequestCreated                    bool                                                                                                                                                            `json:"portal_request_created"`
	NotificationActionEnabled               bool                                                                                                                                                            `json:"notification_action_enabled"`
	CompatibilityCenterOpened               bool                                                                                                                                                            `json:"compatibility_center_opened"`
	SupportBundleExported                   bool                                                                                                                                                            `json:"support_bundle_exported"`
	SupportCaseCreated                      bool                                                                                                                                                            `json:"support_case_created"`
	RuntimeOwned                            bool                                                                                                                                                            `json:"runtime_owned"`
	GoRuntimeBacked                         bool                                                                                                                                                            `json:"go_runtime_backed"`
	KDEPolicyOwner                          bool                                                                                                                                                            `json:"kde_policy_owner"`
	ProductionReadiness                     bool                                                                                                                                                            `json:"production_readiness"`
	ProductionOwnershipReady                bool                                                                                                                                                            `json:"production_ownership_ready"`
	SystemServiceStarted                    bool                                                                                                                                                            `json:"system_service_started"`
	SessionBusClaimed                       bool                                                                                                                                                            `json:"session_bus_claimed"`
	ProductionBusClaimed                    bool                                                                                                                                                            `json:"production_bus_claimed"`
	ProductionOwnerEnabled                  bool                                                                                                                                                            `json:"production_owner_enabled"`
	WriteMethodsEnabled                     bool                                                                                                                                                            `json:"write_methods_enabled"`
	RuntimeWritesEnabled                    bool                                                                                                                                                            `json:"runtime_writes_enabled"`
	DesktopFilesWritten                     bool                                                                                                                                                            `json:"desktop_files_written"`
	SettingsPersisted                       bool                                                                                                                                                            `json:"settings_persisted"`
	AdapterInvocationEnabled                bool                                                                                                                                                            `json:"adapter_invocation_enabled"`
	BackendLaunchEnabled                    bool                                                                                                                                                            `json:"backend_launch_enabled"`
	BackendProcessStarted                   bool                                                                                                                                                            `json:"backend_process_started"`
	SnapshotRestoreExecuted                 bool                                                                                                                                                            `json:"snapshot_restore_executed"`
	StateCleanupExecuted                    bool                                                                                                                                                            `json:"state_cleanup_executed"`
	NetworkRequired                         bool                                                                                                                                                            `json:"network_required"`
	HostRootModified                        bool                                                                                                                                                            `json:"host_root_modified"`
	PrivilegedContainerRequired             bool                                                                                                                                                            `json:"privileged_container_required"`
	CallerStateRootRequired                 bool                                                                                                                                                            `json:"caller_state_root_required"`
	StateRootPathExposed                    bool                                                                                                                                                            `json:"state_root_path_exposed"`
	FilePathsExposed                        bool                                                                                                                                                            `json:"file_paths_exposed"`
	FileContentRead                         bool                                                                                                                                                            `json:"file_content_read"`
	RawCommandExposed                       bool                                                                                                                                                            `json:"raw_command_exposed"`
	RawExecutableExposed                    bool                                                                                                                                                            `json:"raw_executable_exposed"`
	BackendDetailsExposed                   bool                                                                                                                                                            `json:"backend_details_exposed"`
	BlockedActions                          []string                                                                                                                                                        `json:"blocked_actions"`
	NextRequirements                        []string                                                                                                                                                        `json:"next_requirements"`
	DesktopSafeSummary                      string                                                                                                                                                          `json:"desktop_safe_summary"`
}

type ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptConsumptionItem struct {
	ID                                     string `json:"id"`
	ActionKind                             string `json:"action_kind"`
	SurfaceKind                            string `json:"surface_kind"`
	StatusConsumerKind                     string `json:"status_consumer_kind"`
	ReceiptScope                           string `json:"receipt_scope"`
	StorageRootPolicyScope                 string `json:"storage_root_policy_scope"`
	EvidencePresent                        bool   `json:"evidence_present"`
	CurrentMainlineConsumed                bool   `json:"current_mainline_consumed"`
	StorageRootPolicyConsumed              bool   `json:"storage_root_policy_consumed"`
	StorageRootPolicyReady                 bool   `json:"storage_root_policy_ready"`
	AuthorizationReceiptConsumptionModeled bool   `json:"authorization_receipt_consumption_modeled"`
	AcceptedReceiptBoundaryModeled         bool   `json:"accepted_receipt_boundary_modeled"`
	ReceiptScopeBoundaryModeled            bool   `json:"receipt_scope_boundary_modeled"`
	StorageRootPolicyGrantBoundaryModeled  bool   `json:"storage_root_policy_grant_boundary_modeled"`
	RecordWriterCallGrantBoundaryModeled   bool   `json:"record_writer_call_grant_boundary_modeled"`
	KDESafeRedactedStatusOnly              bool   `json:"kde_safe_redacted_status_only"`
	ReceiptPresent                         bool   `json:"receipt_present"`
	ReceiptAccepted                        bool   `json:"receipt_accepted"`
	ReceiptConsumed                        bool   `json:"receipt_consumed"`
	StorageRootPolicyGrantAuthorized       bool   `json:"storage_root_policy_grant_authorized"`
	RecordWriterCallAuthorized             bool   `json:"record_writer_call_authorized"`
	WriterCallable                         bool   `json:"writer_callable"`
	ClosedRecordWriterEnabled              bool   `json:"closed_record_writer_enabled"`
	StorageRootResolved                    bool   `json:"storage_root_resolved"`
	StorageRootOwnershipGranted            bool   `json:"storage_root_ownership_granted"`
	StorageRootNamespaceGranted            bool   `json:"storage_root_namespace_granted"`
	StorageRetentionEnforced               bool   `json:"storage_retention_enforced"`
	StorageRedactionEnforced               bool   `json:"storage_redaction_enforced"`
	StorageGatePassed                      bool   `json:"storage_gate_passed"`
	StoragePersistenceAuthorized           bool   `json:"storage_persistence_authorized"`
	DurableRecordWriteEnabled              bool   `json:"durable_record_write_enabled"`
	StatusPersistenceAuthorized            bool   `json:"status_persistence_authorized"`
	StatusPersistenceWriteEnabled          bool   `json:"status_persistence_write_enabled"`
	StatusWriterEnabled                    bool   `json:"status_writer_enabled"`
	KDEStatusWriteEnabled                  bool   `json:"kde_status_write_enabled"`
	RuntimeDiagnosticsWriteEnabled         bool   `json:"runtime_diagnostics_write_enabled"`
	StorageWriteEnabled                    bool   `json:"storage_write_enabled"`
	RedactedSummaryPersisted               bool   `json:"redacted_summary_persisted"`
	KDEStatusPersisted                     bool   `json:"kde_status_persisted"`
	RuntimeDiagnosticsPersisted            bool   `json:"runtime_diagnostics_persisted"`
	RawResultExposed                       bool   `json:"raw_result_exposed"`
	UserVisible                            bool   `json:"user_visible"`
	ReviewOnly                             bool   `json:"review_only"`
	RuntimeOwned                           bool   `json:"runtime_owned"`
	GoRuntimeBacked                        bool   `json:"go_runtime_backed"`
	KDEPolicyOwner                         bool   `json:"kde_policy_owner"`
	SideEffectsDisabled                    bool   `json:"side_effects_disabled"`
	HostRootModified                       bool   `json:"host_root_modified"`
	InternalDetailsExposed                 bool   `json:"internal_details_exposed"`
	AuthorizationReceiptConsumptionStatus  string `json:"authorization_receipt_consumption_status"`
	NextRequirement                        string `json:"next_requirement"`
}

type ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptConsumptionCheck struct {
	ID      string `json:"id"`
	Status  string `json:"status"`
	Summary string `json:"summary"`
}

type ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptConsumptionCounts struct {
	Total   int `json:"total"`
	Passed  int `json:"passed"`
	Pending int `json:"pending"`
	Blocked int `json:"blocked"`
}

func NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptConsumptionAuditPreview(root string) (ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptConsumptionAuditPreview, error) {
	version, err := readProductionDBusGateReviewVersion(root)
	if err != nil {
		return ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptConsumptionAuditPreview{}, err
	}
	sources := productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptConsumptionSources(root)
	mainlineReady := productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptConsumptionMainlineReady(sources.CurrentMainline)
	policyReady := productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptConsumptionPolicyReady(sources.StorageRootPolicy)
	items := productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptConsumptionItems(mainlineReady, policyReady)
	readyCount := productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptConsumptionReadyCount(items)
	preview := ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptConsumptionAuditPreview{
		Version:                                 version,
		SchemaVersion:                           "xnix.runtime.production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_kde_safe_redacted_status_closed_record_writer_storage_root_authorization_receipt_consumption_audit.v1",
		RequestType:                             "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-consumption-audit-preview",
		PreviewType:                             "receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-consumption-audit",
		Source:                                  "xnix-current-mainline+redacted-status-closed-record-writer-storage-root-policy-audit",
		PreviewDecision:                         "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-consumption-audit-blocked",
		CurrentMainlineConsumed:                 mainlineReady,
		StorageRootPolicyConsumed:               policyReady,
		StorageRootPolicyReady:                  policyReady,
		AuthorizationReceiptConsumptionRequired: true,
		AuthorizationReceiptConsumptionModeled:  mainlineReady && policyReady,
		AuthorizationReceiptConsumptionReady:    mainlineReady && policyReady,
		AcceptedReceiptBoundaryModeled:          mainlineReady && policyReady,
		ReceiptScopeBoundaryModeled:             mainlineReady && policyReady,
		StorageRootPolicyGrantBoundaryModeled:   mainlineReady && policyReady,
		RecordWriterCallGrantBoundaryModeled:    mainlineReady && policyReady,
		KDESafeRedactedStatusOnly:               mainlineReady && policyReady,
		CompatibilityCenterConsumptionModeled:   mainlineReady && policyReady,
		RuntimeDiagnosticsConsumptionModeled:    mainlineReady && policyReady,
		ReceiptPresent:                          false,
		ReceiptAccepted:                         false,
		ReceiptConsumed:                         false,
		StorageRootPolicyGrantAuthorized:        false,
		RecordWriterCallAuthorized:              false,
		WriterCallable:                          false,
		ClosedRecordWriterEnabled:               false,
		StorageRootResolved:                     false,
		StorageRootCreated:                      false,
		StorageRootMounted:                      false,
		StorageRootOwnershipGranted:             false,
		StorageRootNamespaceGranted:             false,
		StorageRetentionEnforced:                false,
		StorageRedactionEnforced:                false,
		StorageGatePassed:                       false,
		StoragePersistenceAuthorized:            false,
		DurableRecordWriteEnabled:               false,
		WriterAuthorizationGranted:              false,
		StatusWriterEnabled:                     false,
		StatusPersistenceAuthorized:             false,
		StatusPersistenceWriteEnabled:           false,
		KDEStatusWriteEnabled:                   false,
		RuntimeDiagnosticsWriteEnabled:          false,
		StorageWriteEnabled:                     false,
		ConsumptionItemCount:                    len(items),
		RequiredConsumptionItemCount:            len(items),
		ReadyConsumptionItemCount:               readyCount,
		MissingConsumptionItemCount:             len(items) - readyCount,
		AcceptedReceiptItemCount:                0,
		ConsumedReceiptItemCount:                0,
		AuthorizedStorageRootGrantItemCount:     0,
		AuthorizedRecordWriterCallItemCount:     0,
		CallableWriterItemCount:                 0,
		EnabledRecordWriterItemCount:            0,
		DurableWrittenRecordItemCount:           0,
		RawExposedConsumptionItemCount:          0,
		SideEffectConsumptionItemCount:          0,
		CompatibilityCenterItemCount:            productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptConsumptionSurfaceCount(items, "compatibility-center"),
		RuntimeDiagnosticsItemCount:             productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptConsumptionSurfaceCount(items, "runtime-diagnostics"),
		ConsumptionItems:                        items,
		ConsumptionItemIDs:                      productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptConsumptionItemIDs(items),
		ConsumerConsumptionAuthorized:           false,
		ConsumerEnablementAuthorized:            false,
		KDEConsumerEnabled:                      false,
		RuntimeConsumerEnabled:                  false,
		LookupRouteAuthorized:                   false,
		LookupRouteEnabled:                      false,
		OpaqueLookupEnabled:                     false,
		RedactedSummaryPersisted:                false,
		KDEStatusPersisted:                      false,
		RuntimeDiagnosticsPersisted:             false,
		DryRunResultPersisted:                   false,
		RawResultExposed:                        false,
		DispatchDryRunExecuted:                  false,
		RequestObjectCreationEnabled:            false,
		RequestObjectDispatchEnabled:            false,
		PortalRequestCreated:                    false,
		NotificationActionEnabled:               false,
		CompatibilityCenterOpened:               false,
		SupportBundleExported:                   false,
		SupportCaseCreated:                      false,
		RuntimeOwned:                            true,
		GoRuntimeBacked:                         true,
		KDEPolicyOwner:                          false,
		ProductionReadiness:                     false,
		ProductionOwnershipReady:                false,
		SystemServiceStarted:                    false,
		SessionBusClaimed:                       false,
		ProductionBusClaimed:                    false,
		ProductionOwnerEnabled:                  false,
		WriteMethodsEnabled:                     false,
		RuntimeWritesEnabled:                    false,
		DesktopFilesWritten:                     false,
		SettingsPersisted:                       false,
		AdapterInvocationEnabled:                false,
		BackendLaunchEnabled:                    false,
		BackendProcessStarted:                   false,
		SnapshotRestoreExecuted:                 false,
		StateCleanupExecuted:                    false,
		NetworkRequired:                         false,
		HostRootModified:                        false,
		PrivilegedContainerRequired:             false,
		CallerStateRootRequired:                 false,
		StateRootPathExposed:                    false,
		FilePathsExposed:                        false,
		FileContentRead:                         false,
		RawCommandExposed:                       false,
		RawExecutableExposed:                    false,
		BackendDetailsExposed:                   false,
		BlockedActions: []string{
			"do not accept or consume storage-root authorization receipts",
			"do not authorize storage-root policy grants or record writer calls",
			"do not resolve, create, mount, or grant storage roots",
			"do not call closed record writers or write durable records",
			"do not enable KDE or Runtime diagnostics consumers",
			"do not create lookup, dispatch, request, Portal, notification, navigation, or support side effects",
			"do not claim production D-Bus ownership, launch engines, expose paths, or mutate the host",
		},
		NextRequirements: []string{
			"add dry-run authorization receipt lookup evidence before any receipt can be consumed",
			"add storage-root grant dry-run evidence before policy grants can be authorized",
			"keep authorization receipt consumption modeled but unaccepted until explicit production authorization exists",
		},
		DesktopSafeSummary: "The KDE-safe storage-root authorization receipt consumption audit consumes the current mainline and storage-root policy audit, models accepted receipt and writer-call grant boundaries, and keeps receipt consumption, storage-root grants, writer calls, persistence writes, consumers, desktop actions, unsafe data exposure, and host mutation disabled.",
	}
	checks := productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptConsumptionChecks(preview)
	preview.Checks = checks
	preview.CheckIDs = productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptConsumptionCheckIDs(checks)
	preview.Counts = countProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptConsumptionChecks(checks)
	if preview.Counts.Blocked == 0 {
		preview.PreviewDecision = "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-consumption-audit-ready-writes-disabled"
	}
	if err := validateNoBackendTerms(preview, "production receipt notification action dry-run result lookup consumer enablement kde safe redacted status closed record writer storage-root authorization receipt consumption audit preview"); err != nil {
		return ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptConsumptionAuditPreview{}, err
	}
	return preview, nil
}

type productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptConsumptionSourceSet struct {
	CurrentMainline   string
	StorageRootPolicy string
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptConsumptionSources(root string) productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptConsumptionSourceSet {
	return productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptConsumptionSourceSet{
		CurrentMainline:   productionAuthorizationReadSources(root, []string{"docs/xnix-current-mainline.md"}),
		StorageRootPolicy: productionAuthorizationReadSources(root, []string{"internal/runtime/owner/production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_kde_safe_redacted_status_closed_record_writer_storage_root_policy_audit.go", "internal/runtime/owner/production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_kde_safe_redacted_status_closed_record_writer_storage_root_policy_audit_test.go"}),
	}
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptConsumptionMainlineReady(source string) bool {
	return productionAuthorizationHasAll(source, []string{"Xnix Current Mainline", "redacted status closed record writer storage-root policy audit preview", "redacted status closed record writer storage-root authorization receipt consumption audit preview", "No Runtime write methods", "No KDE configuration writes"})
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptConsumptionPolicyReady(source string) bool {
	return productionAuthorizationHasAll(source, []string{"production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-policy-audit-preview", "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-policy-audit-ready-writes-disabled", "StorageRootPolicyReady", "StorageRootOwnershipModeled", "StorageRootResolved", "StorageWriteEnabled"})
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptConsumptionItems(mainlineReady bool, policyReady bool) []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptConsumptionItem {
	return []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptConsumptionItem{
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptConsumptionItem("review-receipt-dry-run-result-lookup-consumer-enablement-kde-safe-compatibility-center-storage-root-authorization-consumption", "review", "compatibility-center", "review-result-compatibility-center-status", "compatibility-center-storage-root-authorization-receipt", "compatibility-center-redacted-status-storage-root", mainlineReady, policyReady),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptConsumptionItem("review-receipt-dry-run-result-lookup-consumer-enablement-kde-safe-runtime-diagnostics-storage-root-authorization-consumption", "review", "runtime-diagnostics", "review-result-runtime-diagnostics-status", "runtime-diagnostics-storage-root-authorization-receipt", "runtime-diagnostics-redacted-status-storage-root", mainlineReady, policyReady),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptConsumptionItem("renew-receipt-dry-run-result-lookup-consumer-enablement-kde-safe-compatibility-center-storage-root-authorization-consumption", "renew", "compatibility-center", "renewal-result-compatibility-center-status", "compatibility-center-storage-root-authorization-receipt", "compatibility-center-redacted-status-storage-root", mainlineReady, policyReady),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptConsumptionItem("renew-receipt-dry-run-result-lookup-consumer-enablement-kde-safe-runtime-diagnostics-storage-root-authorization-consumption", "renew", "runtime-diagnostics", "renewal-result-runtime-diagnostics-status", "runtime-diagnostics-storage-root-authorization-receipt", "runtime-diagnostics-redacted-status-storage-root", mainlineReady, policyReady),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptConsumptionItem("open-compatibility-center-dry-run-result-lookup-consumer-enablement-kde-safe-compatibility-center-storage-root-authorization-consumption", "open-compatibility-center", "compatibility-center", "navigation-result-compatibility-center-status", "compatibility-center-storage-root-authorization-receipt", "compatibility-center-redacted-status-storage-root", mainlineReady, policyReady),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptConsumptionItem("open-compatibility-center-dry-run-result-lookup-consumer-enablement-kde-safe-runtime-diagnostics-storage-root-authorization-consumption", "open-compatibility-center", "runtime-diagnostics", "navigation-result-runtime-diagnostics-status", "runtime-diagnostics-storage-root-authorization-receipt", "runtime-diagnostics-redacted-status-storage-root", mainlineReady, policyReady),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptConsumptionItem("dismiss-receipt-dry-run-result-lookup-consumer-enablement-kde-safe-compatibility-center-storage-root-authorization-consumption", "dismiss", "compatibility-center", "dismissal-result-compatibility-center-status", "compatibility-center-storage-root-authorization-receipt", "compatibility-center-redacted-status-storage-root", mainlineReady, policyReady),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptConsumptionItem("dismiss-receipt-dry-run-result-lookup-consumer-enablement-kde-safe-runtime-diagnostics-storage-root-authorization-consumption", "dismiss", "runtime-diagnostics", "dismissal-result-runtime-diagnostics-status", "runtime-diagnostics-storage-root-authorization-receipt", "runtime-diagnostics-redacted-status-storage-root", mainlineReady, policyReady),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptConsumptionItem("support-info-dry-run-result-lookup-consumer-enablement-kde-safe-compatibility-center-storage-root-authorization-consumption", "support-info", "compatibility-center", "support-info-result-compatibility-center-status", "compatibility-center-storage-root-authorization-receipt", "compatibility-center-redacted-status-storage-root", mainlineReady, policyReady),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptConsumptionItem("support-info-dry-run-result-lookup-consumer-enablement-kde-safe-runtime-diagnostics-storage-root-authorization-consumption", "support-info", "runtime-diagnostics", "support-info-result-runtime-diagnostics-status", "runtime-diagnostics-storage-root-authorization-receipt", "runtime-diagnostics-redacted-status-storage-root", mainlineReady, policyReady),
	}
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptConsumptionItem(id string, actionKind string, surfaceKind string, statusConsumerKind string, receiptScope string, storageRootPolicyScope string, mainlineReady bool, policyReady bool) ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptConsumptionItem {
	ready := mainlineReady && policyReady
	status := "missing-redacted-status-storage-root-authorization-receipt-consumption-evidence"
	if ready {
		status = "redacted-status-storage-root-authorization-receipt-consumption-modeled-writes-disabled"
	}
	return ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptConsumptionItem{
		ID:                                     id,
		ActionKind:                             actionKind,
		SurfaceKind:                            surfaceKind,
		StatusConsumerKind:                     statusConsumerKind,
		ReceiptScope:                           receiptScope,
		StorageRootPolicyScope:                 storageRootPolicyScope,
		EvidencePresent:                        ready,
		CurrentMainlineConsumed:                mainlineReady,
		StorageRootPolicyConsumed:              policyReady,
		StorageRootPolicyReady:                 policyReady,
		AuthorizationReceiptConsumptionModeled: ready,
		AcceptedReceiptBoundaryModeled:         ready,
		ReceiptScopeBoundaryModeled:            ready,
		StorageRootPolicyGrantBoundaryModeled:  ready,
		RecordWriterCallGrantBoundaryModeled:   ready,
		KDESafeRedactedStatusOnly:              ready,
		ReceiptPresent:                         false,
		ReceiptAccepted:                        false,
		ReceiptConsumed:                        false,
		StorageRootPolicyGrantAuthorized:       false,
		RecordWriterCallAuthorized:             false,
		WriterCallable:                         false,
		ClosedRecordWriterEnabled:              false,
		StorageRootResolved:                    false,
		StorageRootOwnershipGranted:            false,
		StorageRootNamespaceGranted:            false,
		StorageRetentionEnforced:               false,
		StorageRedactionEnforced:               false,
		StorageGatePassed:                      false,
		StoragePersistenceAuthorized:           false,
		DurableRecordWriteEnabled:              false,
		StatusPersistenceAuthorized:            false,
		StatusPersistenceWriteEnabled:          false,
		StatusWriterEnabled:                    false,
		KDEStatusWriteEnabled:                  false,
		RuntimeDiagnosticsWriteEnabled:         false,
		StorageWriteEnabled:                    false,
		RedactedSummaryPersisted:               false,
		KDEStatusPersisted:                     false,
		RuntimeDiagnosticsPersisted:            false,
		RawResultExposed:                       false,
		UserVisible:                            ready,
		ReviewOnly:                             true,
		RuntimeOwned:                           true,
		GoRuntimeBacked:                        true,
		KDEPolicyOwner:                         false,
		SideEffectsDisabled:                    true,
		HostRootModified:                       false,
		InternalDetailsExposed:                 false,
		AuthorizationReceiptConsumptionStatus:  status,
		NextRequirement:                        "require dry-run receipt lookup and explicit production authorization before storage-root policy grants can authorize writer calls",
	}
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptConsumptionChecks(preview ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptConsumptionAuditPreview) []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptConsumptionCheck {
	return []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptConsumptionCheck{
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptConsumptionCheck("current-mainline-consumed", productionAuthorizationPassBlocked(preview.CurrentMainlineConsumed), "The storage-root authorization receipt consumption audit consumes the Xnix current mainline."),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptConsumptionCheck("storage-root-policy-consumed", productionAuthorizationPassBlocked(preview.StorageRootPolicyConsumed && preview.StorageRootPolicyReady), "The storage-root authorization receipt consumption audit consumes the storage-root policy audit."),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptConsumptionCheck("authorization-receipt-consumption-modeled", productionAuthorizationPassBlocked(preview.AuthorizationReceiptConsumptionRequired && preview.AuthorizationReceiptConsumptionModeled && preview.AuthorizationReceiptConsumptionReady && preview.AcceptedReceiptBoundaryModeled && preview.ReceiptScopeBoundaryModeled && preview.StorageRootPolicyGrantBoundaryModeled && preview.RecordWriterCallGrantBoundaryModeled), "Accepted receipt, scope, storage-root policy grant, and writer-call grant boundaries are modeled."),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptConsumptionCheck("compatibility-center-and-runtime-consumption-modeled", productionAuthorizationPassBlocked(preview.CompatibilityCenterConsumptionModeled && preview.RuntimeDiagnosticsConsumptionModeled && preview.CompatibilityCenterItemCount == 5 && preview.RuntimeDiagnosticsItemCount == 5), "Compatibility Center and Runtime diagnostics authorization consumption candidates are both modeled."),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptConsumptionCheck("ten-consumption-items-ready-writes-disabled", productionAuthorizationPassBlocked(preview.ConsumptionItemCount == 10 && preview.RequiredConsumptionItemCount == 10 && preview.ReadyConsumptionItemCount == 10 && preview.MissingConsumptionItemCount == 0 && preview.AcceptedReceiptItemCount == 0 && preview.ConsumedReceiptItemCount == 0 && preview.AuthorizedStorageRootGrantItemCount == 0 && preview.AuthorizedRecordWriterCallItemCount == 0 && preview.CallableWriterItemCount == 0 && preview.EnabledRecordWriterItemCount == 0 && preview.DurableWrittenRecordItemCount == 0), "All authorization receipt consumption candidates are present while receipts, grants, writer calls, and writes remain disabled."),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptConsumptionCheck("authorization-consumption-and-writes-disabled", productionAuthorizationPassBlocked(!preview.ReceiptPresent && !preview.ReceiptAccepted && !preview.ReceiptConsumed && !preview.StorageRootPolicyGrantAuthorized && !preview.RecordWriterCallAuthorized && !preview.WriterCallable && !preview.ClosedRecordWriterEnabled && !preview.StorageRootResolved && !preview.StorageRootCreated && !preview.StorageRootMounted && !preview.StorageRootOwnershipGranted && !preview.StorageRootNamespaceGranted && !preview.StorageRetentionEnforced && !preview.StorageRedactionEnforced && !preview.StorageGatePassed && !preview.StoragePersistenceAuthorized && !preview.DurableRecordWriteEnabled && !preview.WriterAuthorizationGranted && !preview.StatusWriterEnabled && !preview.StatusPersistenceAuthorized && !preview.StatusPersistenceWriteEnabled && !preview.KDEStatusWriteEnabled && !preview.RuntimeDiagnosticsWriteEnabled && !preview.StorageWriteEnabled && !preview.RedactedSummaryPersisted && !preview.KDEStatusPersisted && !preview.RuntimeDiagnosticsPersisted && !preview.DryRunResultPersisted), "Receipt consumption, grants, storage-root resolution, writer calls, durable record writes, status writes, and persistence remain disabled."),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptConsumptionCheck("consumer-lookup-request-and-support-disabled", productionAuthorizationPassBlocked(!preview.ConsumerConsumptionAuthorized && !preview.ConsumerEnablementAuthorized && !preview.KDEConsumerEnabled && !preview.RuntimeConsumerEnabled && !preview.LookupRouteAuthorized && !preview.LookupRouteEnabled && !preview.OpaqueLookupEnabled && !preview.RequestObjectCreationEnabled && !preview.RequestObjectDispatchEnabled && !preview.PortalRequestCreated && !preview.NotificationActionEnabled && !preview.CompatibilityCenterOpened && !preview.SupportBundleExported && !preview.SupportCaseCreated), "Consumers, lookup, request objects, desktop actions, and support writes remain disabled."),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptConsumptionCheck("production-and-host-boundary-closed", productionAuthorizationPassBlocked(!preview.SystemServiceStarted && !preview.SessionBusClaimed && !preview.ProductionBusClaimed && !preview.ProductionOwnerEnabled && !preview.ProductionReadiness && !preview.ProductionOwnershipReady && !preview.WriteMethodsEnabled && !preview.RuntimeWritesEnabled && !preview.AdapterInvocationEnabled && !preview.BackendLaunchEnabled && !preview.BackendProcessStarted && !preview.NetworkRequired && !preview.HostRootModified && !preview.PrivilegedContainerRequired && !preview.CallerStateRootRequired && !preview.StateRootPathExposed && !preview.FilePathsExposed && !preview.FileContentRead && !preview.RawResultExposed && !preview.RawCommandExposed && !preview.RawExecutableExposed && !preview.BackendDetailsExposed && preview.RawExposedConsumptionItemCount == 0 && preview.SideEffectConsumptionItemCount == 0 && productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptConsumptionItemsKeepClosed(preview.ConsumptionItems)), "Production ownership, writes, unsafe data exposure, network, privilege, internal details, and host mutation remain disabled."),
	}
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptConsumptionCheck(id string, status string, summary string) ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptConsumptionCheck {
	return ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptConsumptionCheck{ID: id, Status: status, Summary: summary}
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptConsumptionReadyCount(items []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptConsumptionItem) int {
	count := 0
	for _, item := range items {
		if item.EvidencePresent && item.AuthorizationReceiptConsumptionModeled && item.AcceptedReceiptBoundaryModeled && item.ReceiptScopeBoundaryModeled && item.StorageRootPolicyGrantBoundaryModeled && item.RecordWriterCallGrantBoundaryModeled && !item.ReceiptAccepted && !item.ReceiptConsumed && !item.StorageRootPolicyGrantAuthorized && !item.RecordWriterCallAuthorized && !item.WriterCallable && !item.StorageWriteEnabled && !item.HostRootModified && !item.InternalDetailsExposed {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptConsumptionSurfaceCount(items []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptConsumptionItem, surfaceKind string) int {
	count := 0
	for _, item := range items {
		if item.SurfaceKind == surfaceKind {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptConsumptionItemsKeepClosed(items []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptConsumptionItem) bool {
	for _, item := range items {
		if item.ReceiptPresent || item.ReceiptAccepted || item.ReceiptConsumed || item.StorageRootPolicyGrantAuthorized || item.RecordWriterCallAuthorized || item.WriterCallable || item.ClosedRecordWriterEnabled || item.StorageRootResolved || item.StorageRootOwnershipGranted || item.StorageRootNamespaceGranted || item.StorageRetentionEnforced || item.StorageRedactionEnforced || item.StorageGatePassed || item.StoragePersistenceAuthorized || item.DurableRecordWriteEnabled || item.StatusPersistenceAuthorized || item.StatusPersistenceWriteEnabled || item.StatusWriterEnabled || item.KDEStatusWriteEnabled || item.RuntimeDiagnosticsWriteEnabled || item.StorageWriteEnabled || item.RedactedSummaryPersisted || item.KDEStatusPersisted || item.RuntimeDiagnosticsPersisted || item.RawResultExposed || item.KDEPolicyOwner || !item.SideEffectsDisabled || item.HostRootModified || item.InternalDetailsExposed {
			return false
		}
	}
	return true
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptConsumptionItemIDs(items []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptConsumptionItem) []string {
	ids := make([]string, 0, len(items))
	for _, item := range items {
		ids = append(ids, item.ID)
	}
	return ids
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptConsumptionCheckIDs(checks []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptConsumptionCheck) []string {
	ids := make([]string, 0, len(checks))
	for _, check := range checks {
		ids = append(ids, check.ID)
	}
	return ids
}

func countProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptConsumptionChecks(checks []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptConsumptionCheck) ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptConsumptionCounts {
	counts := ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptConsumptionCounts{Total: len(checks)}
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
