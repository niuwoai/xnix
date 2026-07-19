package owner

type ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAuthorizationReceiptAuditPreview struct {
	Version                              string                                                                                                                                 `json:"version"`
	SchemaVersion                        string                                                                                                                                 `json:"schema_version"`
	RequestType                          string                                                                                                                                 `json:"request_type"`
	AuditType                            string                                                                                                                                 `json:"audit_type"`
	Source                               string                                                                                                                                 `json:"source"`
	AuditDecision                        string                                                                                                                                 `json:"audit_decision"`
	CurrentMainlineConsumed              bool                                                                                                                                   `json:"current_mainline_consumed"`
	WriterAuthorizationGateAuditConsumed bool                                                                                                                                   `json:"writer_authorization_gate_audit_consumed"`
	WriterAuthorizationGateReady         bool                                                                                                                                   `json:"writer_authorization_gate_ready"`
	WriterAuthorizationReceiptRequired   bool                                                                                                                                   `json:"writer_authorization_receipt_required"`
	WriterAuthorizationReceiptModeled    bool                                                                                                                                   `json:"writer_authorization_receipt_modeled"`
	WriterAuthorizationReceiptReady      bool                                                                                                                                   `json:"writer_authorization_receipt_ready"`
	OpaqueWriterAuthorizationReceipt     bool                                                                                                                                   `json:"opaque_writer_authorization_receipt"`
	KDESafeRedactedStatusOnly            bool                                                                                                                                   `json:"kde_safe_redacted_status_only"`
	CompatibilityCenterReceiptModeled    bool                                                                                                                                   `json:"compatibility_center_receipt_modeled"`
	RuntimeDiagnosticsReceiptModeled     bool                                                                                                                                   `json:"runtime_diagnostics_receipt_modeled"`
	ReceiptPresent                       bool                                                                                                                                   `json:"receipt_present"`
	ReceiptAccepted                      bool                                                                                                                                   `json:"receipt_accepted"`
	AuthorizationAccepted                bool                                                                                                                                   `json:"authorization_accepted"`
	WriterAuthorizationGranted           bool                                                                                                                                   `json:"writer_authorization_granted"`
	StatusWriterEnabled                  bool                                                                                                                                   `json:"status_writer_enabled"`
	StatusPersistenceWriteEnabled        bool                                                                                                                                   `json:"status_persistence_write_enabled"`
	KDEStatusWriteEnabled                bool                                                                                                                                   `json:"kde_status_write_enabled"`
	RuntimeDiagnosticsWriteEnabled       bool                                                                                                                                   `json:"runtime_diagnostics_write_enabled"`
	ReceiptItemCount                     int                                                                                                                                    `json:"receipt_item_count"`
	RequiredReceiptItemCount             int                                                                                                                                    `json:"required_receipt_item_count"`
	ReadyReceiptItemCount                int                                                                                                                                    `json:"ready_receipt_item_count"`
	MissingReceiptItemCount              int                                                                                                                                    `json:"missing_receipt_item_count"`
	AcceptedReceiptItemCount             int                                                                                                                                    `json:"accepted_receipt_item_count"`
	GrantedWriterItemCount               int                                                                                                                                    `json:"granted_writer_item_count"`
	EnabledWriterItemCount               int                                                                                                                                    `json:"enabled_writer_item_count"`
	PersistedWriterItemCount             int                                                                                                                                    `json:"persisted_writer_item_count"`
	RawExposedReceiptItemCount           int                                                                                                                                    `json:"raw_exposed_receipt_item_count"`
	SideEffectReceiptItemCount           int                                                                                                                                    `json:"side_effect_receipt_item_count"`
	CompatibilityCenterReceiptItemCount  int                                                                                                                                    `json:"compatibility_center_receipt_item_count"`
	RuntimeDiagnosticsReceiptItemCount   int                                                                                                                                    `json:"runtime_diagnostics_receipt_item_count"`
	ReceiptItems                         []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAuthorizationReceiptAuditItem      `json:"receipt_items"`
	ReceiptItemIDs                       []string                                                                                                                               `json:"receipt_item_ids"`
	Checks                               []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAuthorizationReceiptAuditCheck     `json:"checks"`
	CheckIDs                             []string                                                                                                                               `json:"check_ids"`
	Counts                               ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAuthorizationReceiptAuditCheckCounts `json:"counts"`
	ConsumerConsumptionAuthorized        bool                                                                                                                                   `json:"consumer_consumption_authorized"`
	ConsumerEnablementAuthorized         bool                                                                                                                                   `json:"consumer_enablement_authorized"`
	KDEConsumerEnabled                   bool                                                                                                                                   `json:"kde_consumer_enabled"`
	RuntimeConsumerEnabled               bool                                                                                                                                   `json:"runtime_consumer_enabled"`
	LookupRouteAuthorized                bool                                                                                                                                   `json:"lookup_route_authorized"`
	LookupRouteEnabled                   bool                                                                                                                                   `json:"lookup_route_enabled"`
	OpaqueLookupEnabled                  bool                                                                                                                                   `json:"opaque_lookup_enabled"`
	RedactedSummaryPersisted             bool                                                                                                                                   `json:"redacted_summary_persisted"`
	KDEStatusPersisted                   bool                                                                                                                                   `json:"kde_status_persisted"`
	RuntimeDiagnosticsPersisted          bool                                                                                                                                   `json:"runtime_diagnostics_persisted"`
	DryRunResultPersisted                bool                                                                                                                                   `json:"dry_run_result_persisted"`
	RawResultExposed                     bool                                                                                                                                   `json:"raw_result_exposed"`
	DispatchDryRunExecuted               bool                                                                                                                                   `json:"dispatch_dry_run_executed"`
	RequestObjectCreationEnabled         bool                                                                                                                                   `json:"request_object_creation_enabled"`
	RequestObjectDispatchEnabled         bool                                                                                                                                   `json:"request_object_dispatch_enabled"`
	PortalRequestCreated                 bool                                                                                                                                   `json:"portal_request_created"`
	NotificationActionEnabled            bool                                                                                                                                   `json:"notification_action_enabled"`
	CompatibilityCenterOpened            bool                                                                                                                                   `json:"compatibility_center_opened"`
	SupportBundleExported                bool                                                                                                                                   `json:"support_bundle_exported"`
	SupportCaseCreated                   bool                                                                                                                                   `json:"support_case_created"`
	RuntimeOwned                         bool                                                                                                                                   `json:"runtime_owned"`
	GoRuntimeBacked                      bool                                                                                                                                   `json:"go_runtime_backed"`
	KDEPolicyOwner                       bool                                                                                                                                   `json:"kde_policy_owner"`
	ProductionReadiness                  bool                                                                                                                                   `json:"production_readiness"`
	ProductionOwnershipReady             bool                                                                                                                                   `json:"production_ownership_ready"`
	SystemServiceStarted                 bool                                                                                                                                   `json:"system_service_started"`
	SessionBusClaimed                    bool                                                                                                                                   `json:"session_bus_claimed"`
	ProductionBusClaimed                 bool                                                                                                                                   `json:"production_bus_claimed"`
	ProductionOwnerEnabled               bool                                                                                                                                   `json:"production_owner_enabled"`
	WriteMethodsEnabled                  bool                                                                                                                                   `json:"write_methods_enabled"`
	RuntimeWritesEnabled                 bool                                                                                                                                   `json:"runtime_writes_enabled"`
	DesktopFilesWritten                  bool                                                                                                                                   `json:"desktop_files_written"`
	SettingsPersisted                    bool                                                                                                                                   `json:"settings_persisted"`
	AdapterInvocationEnabled             bool                                                                                                                                   `json:"adapter_invocation_enabled"`
	BackendLaunchEnabled                 bool                                                                                                                                   `json:"backend_launch_enabled"`
	BackendProcessStarted                bool                                                                                                                                   `json:"backend_process_started"`
	SnapshotRestoreExecuted              bool                                                                                                                                   `json:"snapshot_restore_executed"`
	StateCleanupExecuted                 bool                                                                                                                                   `json:"state_cleanup_executed"`
	NetworkRequired                      bool                                                                                                                                   `json:"network_required"`
	HostRootModified                     bool                                                                                                                                   `json:"host_root_modified"`
	PrivilegedContainerRequired          bool                                                                                                                                   `json:"privileged_container_required"`
	CallerStateRootRequired              bool                                                                                                                                   `json:"caller_state_root_required"`
	StateRootPathExposed                 bool                                                                                                                                   `json:"state_root_path_exposed"`
	FilePathsExposed                     bool                                                                                                                                   `json:"file_paths_exposed"`
	FileContentRead                      bool                                                                                                                                   `json:"file_content_read"`
	RawCommandExposed                    bool                                                                                                                                   `json:"raw_command_exposed"`
	RawExecutableExposed                 bool                                                                                                                                   `json:"raw_executable_exposed"`
	BackendDetailsExposed                bool                                                                                                                                   `json:"backend_details_exposed"`
	BlockedActions                       []string                                                                                                                               `json:"blocked_actions"`
	NextRequirements                     []string                                                                                                                               `json:"next_requirements"`
	DesktopSafeSummary                   string                                                                                                                                 `json:"desktop_safe_summary"`
}

type ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAuthorizationReceiptAuditItem struct {
	ID                                   string `json:"id"`
	ActionKind                           string `json:"action_kind"`
	SurfaceKind                          string `json:"surface_kind"`
	StatusConsumerKind                   string `json:"status_consumer_kind"`
	OpaqueReceiptID                      string `json:"opaque_receipt_id"`
	EvidencePresent                      bool   `json:"evidence_present"`
	CurrentMainlineConsumed              bool   `json:"current_mainline_consumed"`
	WriterAuthorizationGateAuditConsumed bool   `json:"writer_authorization_gate_audit_consumed"`
	WriterAuthorizationReceiptModeled    bool   `json:"writer_authorization_receipt_modeled"`
	OpaqueWriterAuthorizationReceipt     bool   `json:"opaque_writer_authorization_receipt"`
	KDESafeRedactedStatusOnly            bool   `json:"kde_safe_redacted_status_only"`
	ReceiptPresent                       bool   `json:"receipt_present"`
	ReceiptAccepted                      bool   `json:"receipt_accepted"`
	AuthorizationAccepted                bool   `json:"authorization_accepted"`
	WriterAuthorizationGranted           bool   `json:"writer_authorization_granted"`
	StatusWriterEnabled                  bool   `json:"status_writer_enabled"`
	StatusPersistenceWriteEnabled        bool   `json:"status_persistence_write_enabled"`
	KDEStatusWriteEnabled                bool   `json:"kde_status_write_enabled"`
	RuntimeDiagnosticsWriteEnabled       bool   `json:"runtime_diagnostics_write_enabled"`
	RedactedSummaryPersisted             bool   `json:"redacted_summary_persisted"`
	KDEStatusPersisted                   bool   `json:"kde_status_persisted"`
	RuntimeDiagnosticsPersisted          bool   `json:"runtime_diagnostics_persisted"`
	RawResultExposed                     bool   `json:"raw_result_exposed"`
	UserVisible                          bool   `json:"user_visible"`
	ReviewOnly                           bool   `json:"review_only"`
	RuntimeOwned                         bool   `json:"runtime_owned"`
	GoRuntimeBacked                      bool   `json:"go_runtime_backed"`
	KDEPolicyOwner                       bool   `json:"kde_policy_owner"`
	SideEffectsDisabled                  bool   `json:"side_effects_disabled"`
	HostRootModified                     bool   `json:"host_root_modified"`
	InternalDetailsExposed               bool   `json:"internal_details_exposed"`
	WriterAuthorizationReceiptStatus     string `json:"writer_authorization_receipt_status"`
	NextRequirement                      string `json:"next_requirement"`
}

type ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAuthorizationReceiptAuditCheck struct {
	ID      string `json:"id"`
	Status  string `json:"status"`
	Summary string `json:"summary"`
}

type ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAuthorizationReceiptAuditCheckCounts struct {
	Total   int `json:"total"`
	Passed  int `json:"passed"`
	Pending int `json:"pending"`
	Blocked int `json:"blocked"`
}

func NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAuthorizationReceiptAuditPreview(root string) (ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAuthorizationReceiptAuditPreview, error) {
	version, err := readProductionDBusGateReviewVersion(root)
	if err != nil {
		return ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAuthorizationReceiptAuditPreview{}, err
	}
	sources := productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAuthorizationReceiptSources(root)
	currentMainlineReady := productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAuthorizationReceiptMainlineReady(sources.CurrentMainline)
	gateReady := productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAuthorizationReceiptGateReady(sources.GateAudit)
	items := productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAuthorizationReceiptItems(currentMainlineReady, gateReady)
	readyCount := productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAuthorizationReceiptReadyCount(items)
	compatibilityCenterCount := productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAuthorizationReceiptSurfaceCount(items, "compatibility-center")
	runtimeDiagnosticsCount := productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAuthorizationReceiptSurfaceCount(items, "runtime-diagnostics")
	ready := currentMainlineReady && gateReady && readyCount == 10 && compatibilityCenterCount == 5 && runtimeDiagnosticsCount == 5
	preview := ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAuthorizationReceiptAuditPreview{
		Version:                              version,
		SchemaVersion:                        "xnix.runtime.production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_kde_safe_redacted_status_writer_authorization_receipt_audit.v1",
		RequestType:                          "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-writer-authorization-receipt-audit-preview",
		AuditType:                            "receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-writer-authorization-receipt-audit",
		Source:                               "xnix-current-mainline+redacted-status-writer-authorization-gate",
		AuditDecision:                        "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-writer-authorization-receipt-audit-blocked",
		CurrentMainlineConsumed:              currentMainlineReady,
		WriterAuthorizationGateAuditConsumed: gateReady,
		WriterAuthorizationGateReady:         gateReady,
		WriterAuthorizationReceiptRequired:   true,
		WriterAuthorizationReceiptModeled:    ready,
		WriterAuthorizationReceiptReady:      ready,
		OpaqueWriterAuthorizationReceipt:     ready,
		KDESafeRedactedStatusOnly:            ready,
		CompatibilityCenterReceiptModeled:    compatibilityCenterCount == 5 && readyCount == 10,
		RuntimeDiagnosticsReceiptModeled:     runtimeDiagnosticsCount == 5 && readyCount == 10,
		ReceiptPresent:                       false,
		ReceiptAccepted:                      false,
		AuthorizationAccepted:                false,
		WriterAuthorizationGranted:           false,
		StatusWriterEnabled:                  false,
		StatusPersistenceWriteEnabled:        false,
		KDEStatusWriteEnabled:                false,
		RuntimeDiagnosticsWriteEnabled:       false,
		ReceiptItemCount:                     len(items),
		RequiredReceiptItemCount:             10,
		ReadyReceiptItemCount:                readyCount,
		MissingReceiptItemCount:              len(items) - readyCount,
		AcceptedReceiptItemCount:             0,
		GrantedWriterItemCount:               0,
		EnabledWriterItemCount:               0,
		PersistedWriterItemCount:             0,
		RawExposedReceiptItemCount:           0,
		SideEffectReceiptItemCount:           0,
		CompatibilityCenterReceiptItemCount:  compatibilityCenterCount,
		RuntimeDiagnosticsReceiptItemCount:   runtimeDiagnosticsCount,
		ReceiptItems:                         items,
		ReceiptItemIDs:                       productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAuthorizationReceiptItemIDs(items),
		ConsumerConsumptionAuthorized:        false,
		ConsumerEnablementAuthorized:         false,
		KDEConsumerEnabled:                   false,
		RuntimeConsumerEnabled:               false,
		LookupRouteAuthorized:                false,
		LookupRouteEnabled:                   false,
		OpaqueLookupEnabled:                  false,
		RedactedSummaryPersisted:             false,
		KDEStatusPersisted:                   false,
		RuntimeDiagnosticsPersisted:          false,
		DryRunResultPersisted:                false,
		RawResultExposed:                     false,
		DispatchDryRunExecuted:               false,
		RequestObjectCreationEnabled:         false,
		RequestObjectDispatchEnabled:         false,
		PortalRequestCreated:                 false,
		NotificationActionEnabled:            false,
		CompatibilityCenterOpened:            false,
		SupportBundleExported:                false,
		SupportCaseCreated:                   false,
		RuntimeOwned:                         true,
		GoRuntimeBacked:                      true,
		KDEPolicyOwner:                       false,
		ProductionReadiness:                  false,
		ProductionOwnershipReady:             false,
		SystemServiceStarted:                 false,
		SessionBusClaimed:                    false,
		ProductionBusClaimed:                 false,
		ProductionOwnerEnabled:               false,
		WriteMethodsEnabled:                  false,
		RuntimeWritesEnabled:                 false,
		DesktopFilesWritten:                  false,
		SettingsPersisted:                    false,
		AdapterInvocationEnabled:             false,
		BackendLaunchEnabled:                 false,
		BackendProcessStarted:                false,
		SnapshotRestoreExecuted:              false,
		StateCleanupExecuted:                 false,
		NetworkRequired:                      false,
		HostRootModified:                     false,
		PrivilegedContainerRequired:          false,
		CallerStateRootRequired:              false,
		StateRootPathExposed:                 false,
		FilePathsExposed:                     false,
		FileContentRead:                      false,
		RawCommandExposed:                    false,
		RawExecutableExposed:                 false,
		BackendDetailsExposed:                false,
		BlockedActions: []string{
			"accept redacted status writer authorization receipts",
			"grant writer authorization from modeled receipt records",
			"enable status writers, persist status records, open desktop surfaces, create Portal requests, export support data, or mutate host state",
		},
		NextRequirements: []string{
			"Implement a separate receipt acceptance authorization audit before modeled writer receipts can be accepted.",
			"Keep receipt acceptance separate from status writer enablement.",
			"Keep KDE-facing output status-only and avoid exposing project roots, file paths, raw commands, or implementation internals.",
		},
		DesktopSafeSummary: "The KDE-safe redacted status writer authorization receipt audit consumes the Xnix current mainline and the writer authorization gate, models opaque receipt records required before future writer authorization can be accepted, and keeps receipt acceptance, writer authorization grants, status writes, desktop side effects, unsafe data exposure, and host mutation disabled.",
	}
	checks := productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAuthorizationReceiptChecks(preview)
	preview.Checks = checks
	preview.CheckIDs = productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAuthorizationReceiptCheckIDs(checks)
	preview.Counts = countProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAuthorizationReceiptChecks(checks)
	if preview.Counts.Blocked == 0 {
		preview.AuditDecision = "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-writer-authorization-receipt-audit-ready-receipt-disabled"
	}
	if err := validateNoBackendTerms(preview, "production receipt notification action dry-run result lookup consumer enablement kde safe redacted status writer authorization receipt audit preview"); err != nil {
		return ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAuthorizationReceiptAuditPreview{}, err
	}
	return preview, nil
}

type productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAuthorizationReceiptSourceSet struct {
	CurrentMainline string
	GateAudit       string
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAuthorizationReceiptSources(root string) productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAuthorizationReceiptSourceSet {
	return productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAuthorizationReceiptSourceSet{
		CurrentMainline: productionAuthorizationReadSources(root, []string{"docs/xnix-current-mainline.md"}),
		GateAudit:       productionAuthorizationReadSources(root, []string{"internal/runtime/owner/production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_kde_safe_redacted_status_writer_authorization_gate_audit.go", "internal/runtime/owner/production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_kde_safe_redacted_status_writer_authorization_gate_audit_test.go"}),
	}
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAuthorizationReceiptMainlineReady(source string) bool {
	return productionAuthorizationHasAll(source, []string{"Xnix Current Mainline", "redacted status writer authorization receipt audit", "redacted status writer authorization gate audit", "No Runtime write methods", "No KDE configuration writes"})
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAuthorizationReceiptGateReady(source string) bool {
	return productionAuthorizationHasAll(source, []string{"production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-writer-authorization-gate-audit-preview", "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-writer-authorization-gate-audit-ready-writer-disabled", "WriterAuthorizationBoundaryReady", "WriterAuthorizationGranted", "StatusWriterEnabled", "CompatibilityCenterWriterModeled", "RuntimeDiagnosticsWriterModeled"})
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAuthorizationReceiptItems(mainlineReady bool, gateReady bool) []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAuthorizationReceiptAuditItem {
	return []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAuthorizationReceiptAuditItem{
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAuthorizationReceiptItem("review-receipt-dry-run-result-lookup-consumer-enablement-kde-safe-compatibility-center-redacted-status-writer-authorization-receipt", "review", "compatibility-center", "review-result-compatibility-center-status", "review-compatibility-center-redacted-status-writer-authorization-receipt-id", mainlineReady, gateReady),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAuthorizationReceiptItem("review-receipt-dry-run-result-lookup-consumer-enablement-kde-safe-runtime-diagnostics-redacted-status-writer-authorization-receipt", "review", "runtime-diagnostics", "review-result-runtime-diagnostics-status", "review-runtime-diagnostics-redacted-status-writer-authorization-receipt-id", mainlineReady, gateReady),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAuthorizationReceiptItem("renew-receipt-dry-run-result-lookup-consumer-enablement-kde-safe-compatibility-center-redacted-status-writer-authorization-receipt", "renew", "compatibility-center", "renewal-result-compatibility-center-status", "renewal-compatibility-center-redacted-status-writer-authorization-receipt-id", mainlineReady, gateReady),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAuthorizationReceiptItem("renew-receipt-dry-run-result-lookup-consumer-enablement-kde-safe-runtime-diagnostics-redacted-status-writer-authorization-receipt", "renew", "runtime-diagnostics", "renewal-result-runtime-diagnostics-status", "renewal-runtime-diagnostics-redacted-status-writer-authorization-receipt-id", mainlineReady, gateReady),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAuthorizationReceiptItem("open-compatibility-center-dry-run-result-lookup-consumer-enablement-kde-safe-compatibility-center-redacted-status-writer-authorization-receipt", "open-compatibility-center", "compatibility-center", "navigation-result-compatibility-center-status", "navigation-compatibility-center-redacted-status-writer-authorization-receipt-id", mainlineReady, gateReady),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAuthorizationReceiptItem("open-compatibility-center-dry-run-result-lookup-consumer-enablement-kde-safe-runtime-diagnostics-redacted-status-writer-authorization-receipt", "open-compatibility-center", "runtime-diagnostics", "navigation-result-runtime-diagnostics-status", "navigation-runtime-diagnostics-redacted-status-writer-authorization-receipt-id", mainlineReady, gateReady),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAuthorizationReceiptItem("dismiss-receipt-dry-run-result-lookup-consumer-enablement-kde-safe-compatibility-center-redacted-status-writer-authorization-receipt", "dismiss", "compatibility-center", "dismissal-result-compatibility-center-status", "dismissal-compatibility-center-redacted-status-writer-authorization-receipt-id", mainlineReady, gateReady),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAuthorizationReceiptItem("dismiss-receipt-dry-run-result-lookup-consumer-enablement-kde-safe-runtime-diagnostics-redacted-status-writer-authorization-receipt", "dismiss", "runtime-diagnostics", "dismissal-result-runtime-diagnostics-status", "dismissal-runtime-diagnostics-redacted-status-writer-authorization-receipt-id", mainlineReady, gateReady),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAuthorizationReceiptItem("support-info-dry-run-result-lookup-consumer-enablement-kde-safe-compatibility-center-redacted-status-writer-authorization-receipt", "support-info", "compatibility-center", "support-info-result-compatibility-center-status", "support-info-compatibility-center-redacted-status-writer-authorization-receipt-id", mainlineReady, gateReady),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAuthorizationReceiptItem("support-info-dry-run-result-lookup-consumer-enablement-kde-safe-runtime-diagnostics-redacted-status-writer-authorization-receipt", "support-info", "runtime-diagnostics", "support-info-result-runtime-diagnostics-status", "support-info-runtime-diagnostics-redacted-status-writer-authorization-receipt-id", mainlineReady, gateReady),
	}
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAuthorizationReceiptItem(id string, actionKind string, surfaceKind string, statusConsumerKind string, opaqueReceiptID string, mainlineReady bool, gateReady bool) ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAuthorizationReceiptAuditItem {
	ready := mainlineReady && gateReady
	status := "missing-redacted-status-writer-authorization-receipt-evidence"
	if ready {
		status = "redacted-status-writer-authorization-receipt-modeled-acceptance-disabled"
	}
	return ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAuthorizationReceiptAuditItem{
		ID:                                   id,
		ActionKind:                           actionKind,
		SurfaceKind:                          surfaceKind,
		StatusConsumerKind:                   statusConsumerKind,
		OpaqueReceiptID:                      opaqueReceiptID,
		EvidencePresent:                      ready,
		CurrentMainlineConsumed:              mainlineReady,
		WriterAuthorizationGateAuditConsumed: gateReady,
		WriterAuthorizationReceiptModeled:    ready,
		OpaqueWriterAuthorizationReceipt:     ready,
		KDESafeRedactedStatusOnly:            ready,
		ReceiptPresent:                       false,
		ReceiptAccepted:                      false,
		AuthorizationAccepted:                false,
		WriterAuthorizationGranted:           false,
		StatusWriterEnabled:                  false,
		StatusPersistenceWriteEnabled:        false,
		KDEStatusWriteEnabled:                false,
		RuntimeDiagnosticsWriteEnabled:       false,
		RedactedSummaryPersisted:             false,
		KDEStatusPersisted:                   false,
		RuntimeDiagnosticsPersisted:          false,
		RawResultExposed:                     false,
		UserVisible:                          ready,
		ReviewOnly:                           true,
		RuntimeOwned:                         true,
		GoRuntimeBacked:                      true,
		KDEPolicyOwner:                       false,
		SideEffectsDisabled:                  true,
		HostRootModified:                     false,
		InternalDetailsExposed:               false,
		WriterAuthorizationReceiptStatus:     status,
		NextRequirement:                      "require separate receipt acceptance authorization before accepting this writer authorization receipt",
	}
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAuthorizationReceiptChecks(preview ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAuthorizationReceiptAuditPreview) []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAuthorizationReceiptAuditCheck {
	return []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAuthorizationReceiptAuditCheck{
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAuthorizationReceiptCheck("current-mainline-consumed", productionAuthorizationPassBlocked(preview.CurrentMainlineConsumed), "The receipt audit consumes the Xnix current mainline."),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAuthorizationReceiptCheck("writer-authorization-gate-consumed", productionAuthorizationPassBlocked(preview.WriterAuthorizationGateAuditConsumed && preview.WriterAuthorizationGateReady), "The receipt audit consumes the writer authorization gate."),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAuthorizationReceiptCheck("writer-authorization-receipt-modeled", productionAuthorizationPassBlocked(preview.WriterAuthorizationReceiptRequired && preview.WriterAuthorizationReceiptModeled && preview.WriterAuthorizationReceiptReady && preview.OpaqueWriterAuthorizationReceipt), "The opaque writer authorization receipt boundary is modeled."),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAuthorizationReceiptCheck("compatibility-center-and-runtime-receipts-modeled", productionAuthorizationPassBlocked(preview.CompatibilityCenterReceiptModeled && preview.RuntimeDiagnosticsReceiptModeled && preview.CompatibilityCenterReceiptItemCount == 5 && preview.RuntimeDiagnosticsReceiptItemCount == 5), "Compatibility Center and Runtime diagnostics writer receipt candidates are both modeled."),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAuthorizationReceiptCheck("ten-receipt-items-ready-acceptance-disabled", productionAuthorizationPassBlocked(preview.ReceiptItemCount == 10 && preview.RequiredReceiptItemCount == 10 && preview.ReadyReceiptItemCount == 10 && preview.MissingReceiptItemCount == 0 && preview.AcceptedReceiptItemCount == 0 && preview.GrantedWriterItemCount == 0 && preview.EnabledWriterItemCount == 0), "All redacted status writer authorization receipt candidates are present while receipt acceptance and writers remain disabled."),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAuthorizationReceiptCheck("receipt-acceptance-and-writers-disabled", productionAuthorizationPassBlocked(!preview.ReceiptPresent && !preview.ReceiptAccepted && !preview.AuthorizationAccepted && !preview.WriterAuthorizationGranted && !preview.StatusWriterEnabled && !preview.StatusPersistenceWriteEnabled && !preview.KDEStatusWriteEnabled && !preview.RuntimeDiagnosticsWriteEnabled && !preview.RedactedSummaryPersisted && !preview.KDEStatusPersisted && !preview.RuntimeDiagnosticsPersisted && !preview.DryRunResultPersisted), "Receipt acceptance, writer authorization grants, and status persistence remain disabled."),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAuthorizationReceiptCheck("consumer-lookup-request-and-support-disabled", productionAuthorizationPassBlocked(!preview.ConsumerConsumptionAuthorized && !preview.ConsumerEnablementAuthorized && !preview.KDEConsumerEnabled && !preview.RuntimeConsumerEnabled && !preview.LookupRouteAuthorized && !preview.LookupRouteEnabled && !preview.OpaqueLookupEnabled && !preview.RequestObjectCreationEnabled && !preview.RequestObjectDispatchEnabled && !preview.PortalRequestCreated && !preview.NotificationActionEnabled && !preview.CompatibilityCenterOpened && !preview.SupportBundleExported && !preview.SupportCaseCreated), "Consumers, lookup, request objects, desktop actions, and support writes remain disabled."),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAuthorizationReceiptCheck("production-and-host-boundary-closed", productionAuthorizationPassBlocked(!preview.SystemServiceStarted && !preview.SessionBusClaimed && !preview.ProductionBusClaimed && !preview.ProductionOwnerEnabled && !preview.ProductionReadiness && !preview.ProductionOwnershipReady && !preview.WriteMethodsEnabled && !preview.RuntimeWritesEnabled && !preview.AdapterInvocationEnabled && !preview.BackendLaunchEnabled && !preview.BackendProcessStarted && !preview.NetworkRequired && !preview.HostRootModified && !preview.PrivilegedContainerRequired && !preview.CallerStateRootRequired && !preview.StateRootPathExposed && !preview.FilePathsExposed && !preview.FileContentRead && !preview.RawResultExposed && !preview.RawCommandExposed && !preview.RawExecutableExposed && !preview.BackendDetailsExposed && preview.RawExposedReceiptItemCount == 0 && preview.SideEffectReceiptItemCount == 0 && productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAuthorizationReceiptItemsKeepClosed(preview.ReceiptItems)), "Production ownership, writes, unsafe data exposure, network, privilege, internal details, and host mutation remain disabled."),
	}
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAuthorizationReceiptCheck(id string, status string, summary string) ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAuthorizationReceiptAuditCheck {
	return ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAuthorizationReceiptAuditCheck{ID: id, Status: status, Summary: summary}
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAuthorizationReceiptReadyCount(items []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAuthorizationReceiptAuditItem) int {
	count := 0
	for _, item := range items {
		if item.EvidencePresent {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAuthorizationReceiptSurfaceCount(items []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAuthorizationReceiptAuditItem, surfaceKind string) int {
	count := 0
	for _, item := range items {
		if item.SurfaceKind == surfaceKind {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAuthorizationReceiptItemsKeepClosed(items []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAuthorizationReceiptAuditItem) bool {
	for _, item := range items {
		if item.ReceiptPresent || item.ReceiptAccepted || item.AuthorizationAccepted || item.WriterAuthorizationGranted || item.StatusWriterEnabled || item.StatusPersistenceWriteEnabled || item.KDEStatusWriteEnabled || item.RuntimeDiagnosticsWriteEnabled || item.RedactedSummaryPersisted || item.KDEStatusPersisted || item.RuntimeDiagnosticsPersisted || item.RawResultExposed || !item.SideEffectsDisabled || item.HostRootModified || item.InternalDetailsExposed {
			return false
		}
	}
	return true
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAuthorizationReceiptItemIDs(items []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAuthorizationReceiptAuditItem) []string {
	ids := make([]string, 0, len(items))
	for _, item := range items {
		ids = append(ids, item.ID)
	}
	return ids
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAuthorizationReceiptCheckIDs(checks []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAuthorizationReceiptAuditCheck) []string {
	ids := make([]string, 0, len(checks))
	for _, check := range checks {
		ids = append(ids, check.ID)
	}
	return ids
}

func countProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAuthorizationReceiptChecks(checks []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAuthorizationReceiptAuditCheck) ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAuthorizationReceiptAuditCheckCounts {
	counts := ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAuthorizationReceiptAuditCheckCounts{Total: len(checks)}
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
