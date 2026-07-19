package owner

type ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAcceptedReceiptGateAuditPreview struct {
	Version                                           string                                                                                                                                `json:"version"`
	SchemaVersion                                     string                                                                                                                                `json:"schema_version"`
	RequestType                                       string                                                                                                                                `json:"request_type"`
	AuditType                                         string                                                                                                                                `json:"audit_type"`
	Source                                            string                                                                                                                                `json:"source"`
	AuditDecision                                     string                                                                                                                                `json:"audit_decision"`
	CurrentMainlineConsumed                           bool                                                                                                                                  `json:"current_mainline_consumed"`
	WriterAuthorizationReceiptAcceptanceAuditConsumed bool                                                                                                                                  `json:"writer_authorization_receipt_acceptance_audit_consumed"`
	WriterAuthorizationReceiptAcceptanceReady         bool                                                                                                                                  `json:"writer_authorization_receipt_acceptance_ready"`
	AcceptedReceiptGateRequired                       bool                                                                                                                                  `json:"accepted_receipt_gate_required"`
	AcceptedReceiptGateModeled                        bool                                                                                                                                  `json:"accepted_receipt_gate_modeled"`
	AcceptedReceiptGateReady                          bool                                                                                                                                  `json:"accepted_receipt_gate_ready"`
	AcceptedReceiptGateBoundaryReady                  bool                                                                                                                                  `json:"accepted_receipt_gate_boundary_ready"`
	OpaqueWriterAuthorizationReceipt                  bool                                                                                                                                  `json:"opaque_writer_authorization_receipt"`
	KDESafeRedactedStatusOnly                         bool                                                                                                                                  `json:"kde_safe_redacted_status_only"`
	CompatibilityCenterGateModeled                    bool                                                                                                                                  `json:"compatibility_center_gate_modeled"`
	RuntimeDiagnosticsGateModeled                     bool                                                                                                                                  `json:"runtime_diagnostics_gate_modeled"`
	ReceiptPresent                                    bool                                                                                                                                  `json:"receipt_present"`
	ReceiptAccepted                                   bool                                                                                                                                  `json:"receipt_accepted"`
	AuthorizationAccepted                             bool                                                                                                                                  `json:"authorization_accepted"`
	WriterAuthorizationGranted                        bool                                                                                                                                  `json:"writer_authorization_granted"`
	StatusWriterEnabled                               bool                                                                                                                                  `json:"status_writer_enabled"`
	StatusPersistenceWriteEnabled                     bool                                                                                                                                  `json:"status_persistence_write_enabled"`
	KDEStatusWriteEnabled                             bool                                                                                                                                  `json:"kde_status_write_enabled"`
	RuntimeDiagnosticsWriteEnabled                    bool                                                                                                                                  `json:"runtime_diagnostics_write_enabled"`
	GateItemCount                                     int                                                                                                                                   `json:"gate_item_count"`
	RequiredGateItemCount                             int                                                                                                                                   `json:"required_gate_item_count"`
	ReadyGateItemCount                                int                                                                                                                                   `json:"ready_gate_item_count"`
	MissingGateItemCount                              int                                                                                                                                   `json:"missing_gate_item_count"`
	AcceptedReceiptItemCount                          int                                                                                                                                   `json:"accepted_receipt_item_count"`
	GrantedWriterItemCount                            int                                                                                                                                   `json:"granted_writer_item_count"`
	EnabledWriterItemCount                            int                                                                                                                                   `json:"enabled_writer_item_count"`
	PersistedWriterItemCount                          int                                                                                                                                   `json:"persisted_writer_item_count"`
	RawExposedGateItemCount                           int                                                                                                                                   `json:"raw_exposed_gate_item_count"`
	SideEffectGateItemCount                           int                                                                                                                                   `json:"side_effect_gate_item_count"`
	CompatibilityCenterGateItemCount                  int                                                                                                                                   `json:"compatibility_center_gate_item_count"`
	RuntimeDiagnosticsGateItemCount                   int                                                                                                                                   `json:"runtime_diagnostics_gate_item_count"`
	GateItems                                         []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAcceptedReceiptGateAuditItem      `json:"gate_items"`
	GateItemIDs                                       []string                                                                                                                              `json:"gate_item_ids"`
	Checks                                            []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAcceptedReceiptGateAuditCheck     `json:"checks"`
	CheckIDs                                          []string                                                                                                                              `json:"check_ids"`
	Counts                                            ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAcceptedReceiptGateAuditCheckCounts `json:"counts"`
	ConsumerConsumptionAuthorized                     bool                                                                                                                                  `json:"consumer_consumption_authorized"`
	ConsumerEnablementAuthorized                      bool                                                                                                                                  `json:"consumer_enablement_authorized"`
	KDEConsumerEnabled                                bool                                                                                                                                  `json:"kde_consumer_enabled"`
	RuntimeConsumerEnabled                            bool                                                                                                                                  `json:"runtime_consumer_enabled"`
	LookupRouteAuthorized                             bool                                                                                                                                  `json:"lookup_route_authorized"`
	LookupRouteEnabled                                bool                                                                                                                                  `json:"lookup_route_enabled"`
	OpaqueLookupEnabled                               bool                                                                                                                                  `json:"opaque_lookup_enabled"`
	RedactedSummaryPersisted                          bool                                                                                                                                  `json:"redacted_summary_persisted"`
	KDEStatusPersisted                                bool                                                                                                                                  `json:"kde_status_persisted"`
	RuntimeDiagnosticsPersisted                       bool                                                                                                                                  `json:"runtime_diagnostics_persisted"`
	DryRunResultPersisted                             bool                                                                                                                                  `json:"dry_run_result_persisted"`
	RawResultExposed                                  bool                                                                                                                                  `json:"raw_result_exposed"`
	DispatchDryRunExecuted                            bool                                                                                                                                  `json:"dispatch_dry_run_executed"`
	RequestObjectCreationEnabled                      bool                                                                                                                                  `json:"request_object_creation_enabled"`
	RequestObjectDispatchEnabled                      bool                                                                                                                                  `json:"request_object_dispatch_enabled"`
	PortalRequestCreated                              bool                                                                                                                                  `json:"portal_request_created"`
	NotificationActionEnabled                         bool                                                                                                                                  `json:"notification_action_enabled"`
	CompatibilityCenterOpened                         bool                                                                                                                                  `json:"compatibility_center_opened"`
	SupportBundleExported                             bool                                                                                                                                  `json:"support_bundle_exported"`
	SupportCaseCreated                                bool                                                                                                                                  `json:"support_case_created"`
	RuntimeOwned                                      bool                                                                                                                                  `json:"runtime_owned"`
	GoRuntimeBacked                                   bool                                                                                                                                  `json:"go_runtime_backed"`
	KDEPolicyOwner                                    bool                                                                                                                                  `json:"kde_policy_owner"`
	ProductionReadiness                               bool                                                                                                                                  `json:"production_readiness"`
	ProductionOwnershipReady                          bool                                                                                                                                  `json:"production_ownership_ready"`
	SystemServiceStarted                              bool                                                                                                                                  `json:"system_service_started"`
	SessionBusClaimed                                 bool                                                                                                                                  `json:"session_bus_claimed"`
	ProductionBusClaimed                              bool                                                                                                                                  `json:"production_bus_claimed"`
	ProductionOwnerEnabled                            bool                                                                                                                                  `json:"production_owner_enabled"`
	WriteMethodsEnabled                               bool                                                                                                                                  `json:"write_methods_enabled"`
	RuntimeWritesEnabled                              bool                                                                                                                                  `json:"runtime_writes_enabled"`
	DesktopFilesWritten                               bool                                                                                                                                  `json:"desktop_files_written"`
	SettingsPersisted                                 bool                                                                                                                                  `json:"settings_persisted"`
	AdapterInvocationEnabled                          bool                                                                                                                                  `json:"adapter_invocation_enabled"`
	BackendLaunchEnabled                              bool                                                                                                                                  `json:"backend_launch_enabled"`
	BackendProcessStarted                             bool                                                                                                                                  `json:"backend_process_started"`
	SnapshotRestoreExecuted                           bool                                                                                                                                  `json:"snapshot_restore_executed"`
	StateCleanupExecuted                              bool                                                                                                                                  `json:"state_cleanup_executed"`
	NetworkRequired                                   bool                                                                                                                                  `json:"network_required"`
	HostRootModified                                  bool                                                                                                                                  `json:"host_root_modified"`
	PrivilegedContainerRequired                       bool                                                                                                                                  `json:"privileged_container_required"`
	CallerStateRootRequired                           bool                                                                                                                                  `json:"caller_state_root_required"`
	StateRootPathExposed                              bool                                                                                                                                  `json:"state_root_path_exposed"`
	FilePathsExposed                                  bool                                                                                                                                  `json:"file_paths_exposed"`
	FileContentRead                                   bool                                                                                                                                  `json:"file_content_read"`
	RawCommandExposed                                 bool                                                                                                                                  `json:"raw_command_exposed"`
	RawExecutableExposed                              bool                                                                                                                                  `json:"raw_executable_exposed"`
	BackendDetailsExposed                             bool                                                                                                                                  `json:"backend_details_exposed"`
	BlockedActions                                    []string                                                                                                                              `json:"blocked_actions"`
	NextRequirements                                  []string                                                                                                                              `json:"next_requirements"`
	DesktopSafeSummary                                string                                                                                                                                `json:"desktop_safe_summary"`
}

type ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAcceptedReceiptGateAuditItem struct {
	ID                                           string `json:"id"`
	ActionKind                                   string `json:"action_kind"`
	SurfaceKind                                  string `json:"surface_kind"`
	StatusConsumerKind                           string `json:"status_consumer_kind"`
	OpaqueReceiptID                              string `json:"opaque_receipt_id"`
	EvidencePresent                              bool   `json:"evidence_present"`
	CurrentMainlineConsumed                      bool   `json:"current_mainline_consumed"`
	WriterAuthorizationReceiptAcceptanceConsumed bool   `json:"writer_authorization_receipt_acceptance_consumed"`
	AcceptedReceiptGateModeled                   bool   `json:"accepted_receipt_gate_modeled"`
	AcceptedReceiptGateBoundaryReady             bool   `json:"accepted_receipt_gate_boundary_ready"`
	OpaqueWriterAuthorizationReceipt             bool   `json:"opaque_writer_authorization_receipt"`
	KDESafeRedactedStatusOnly                    bool   `json:"kde_safe_redacted_status_only"`
	ReceiptPresent                               bool   `json:"receipt_present"`
	ReceiptAccepted                              bool   `json:"receipt_accepted"`
	AuthorizationAccepted                        bool   `json:"authorization_accepted"`
	WriterAuthorizationGranted                   bool   `json:"writer_authorization_granted"`
	StatusWriterEnabled                          bool   `json:"status_writer_enabled"`
	StatusPersistenceWriteEnabled                bool   `json:"status_persistence_write_enabled"`
	KDEStatusWriteEnabled                        bool   `json:"kde_status_write_enabled"`
	RuntimeDiagnosticsWriteEnabled               bool   `json:"runtime_diagnostics_write_enabled"`
	RedactedSummaryPersisted                     bool   `json:"redacted_summary_persisted"`
	KDEStatusPersisted                           bool   `json:"kde_status_persisted"`
	RuntimeDiagnosticsPersisted                  bool   `json:"runtime_diagnostics_persisted"`
	RawResultExposed                             bool   `json:"raw_result_exposed"`
	UserVisible                                  bool   `json:"user_visible"`
	ReviewOnly                                   bool   `json:"review_only"`
	RuntimeOwned                                 bool   `json:"runtime_owned"`
	GoRuntimeBacked                              bool   `json:"go_runtime_backed"`
	KDEPolicyOwner                               bool   `json:"kde_policy_owner"`
	SideEffectsDisabled                          bool   `json:"side_effects_disabled"`
	HostRootModified                             bool   `json:"host_root_modified"`
	InternalDetailsExposed                       bool   `json:"internal_details_exposed"`
	AcceptedReceiptGateStatus                    string `json:"accepted_receipt_gate_status"`
	NextRequirement                              string `json:"next_requirement"`
}

type ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAcceptedReceiptGateAuditCheck struct {
	ID      string `json:"id"`
	Status  string `json:"status"`
	Summary string `json:"summary"`
}

type ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAcceptedReceiptGateAuditCheckCounts struct {
	Total   int `json:"total"`
	Passed  int `json:"passed"`
	Pending int `json:"pending"`
	Blocked int `json:"blocked"`
}

func NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAcceptedReceiptGateAuditPreview(root string) (ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAcceptedReceiptGateAuditPreview, error) {
	version, err := readProductionDBusGateReviewVersion(root)
	if err != nil {
		return ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAcceptedReceiptGateAuditPreview{}, err
	}
	sources := productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAcceptedReceiptGateSources(root)
	mainlineReady := productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAcceptedReceiptGateMainlineReady(sources.CurrentMainline)
	acceptanceReady := productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAcceptedReceiptGateAcceptanceReady(sources.AcceptanceAudit)
	items := productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAcceptedReceiptGateItems(mainlineReady, acceptanceReady)
	readyItemCount := productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAcceptedReceiptGateReadyCount(items)
	gateReady := mainlineReady && acceptanceReady && readyItemCount == len(items)
	preview := ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAcceptedReceiptGateAuditPreview{
		Version:                 version,
		SchemaVersion:           "xnix.runtime.production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_kde_safe_redacted_status_writer_accepted_receipt_gate_audit.v1",
		RequestType:             "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-writer-accepted-receipt-gate-audit-preview",
		AuditType:               "receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-writer-accepted-receipt-gate-audit",
		Source:                  "xnix-current-mainline+redacted-status-writer-authorization-receipt-acceptance",
		AuditDecision:           "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-writer-accepted-receipt-gate-audit-blocked",
		CurrentMainlineConsumed: mainlineReady,
		WriterAuthorizationReceiptAcceptanceAuditConsumed: acceptanceReady,
		WriterAuthorizationReceiptAcceptanceReady:         acceptanceReady,
		AcceptedReceiptGateRequired:                       true,
		AcceptedReceiptGateModeled:                        true,
		AcceptedReceiptGateReady:                          gateReady,
		AcceptedReceiptGateBoundaryReady:                  gateReady,
		OpaqueWriterAuthorizationReceipt:                  acceptanceReady,
		KDESafeRedactedStatusOnly:                         gateReady,
		CompatibilityCenterGateModeled:                    gateReady,
		RuntimeDiagnosticsGateModeled:                     gateReady,
		ReceiptPresent:                                    false,
		ReceiptAccepted:                                   false,
		AuthorizationAccepted:                             false,
		WriterAuthorizationGranted:                        false,
		StatusWriterEnabled:                               false,
		StatusPersistenceWriteEnabled:                     false,
		KDEStatusWriteEnabled:                             false,
		RuntimeDiagnosticsWriteEnabled:                    false,
		GateItemCount:                                     len(items),
		RequiredGateItemCount:                             len(items),
		ReadyGateItemCount:                                readyItemCount,
		MissingGateItemCount:                              len(items) - readyItemCount,
		AcceptedReceiptItemCount:                          0,
		GrantedWriterItemCount:                            0,
		EnabledWriterItemCount:                            0,
		PersistedWriterItemCount:                          0,
		RawExposedGateItemCount:                           0,
		SideEffectGateItemCount:                           0,
		CompatibilityCenterGateItemCount:                  productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAcceptedReceiptGateSurfaceCount(items, "compatibility-center"),
		RuntimeDiagnosticsGateItemCount:                   productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAcceptedReceiptGateSurfaceCount(items, "runtime-diagnostics"),
		GateItems:                                         items,
		GateItemIDs:                                       productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAcceptedReceiptGateItemIDs(items),
		ConsumerConsumptionAuthorized:                     false,
		ConsumerEnablementAuthorized:                      false,
		KDEConsumerEnabled:                                false,
		RuntimeConsumerEnabled:                            false,
		LookupRouteAuthorized:                             false,
		LookupRouteEnabled:                                false,
		OpaqueLookupEnabled:                               false,
		RedactedSummaryPersisted:                          false,
		KDEStatusPersisted:                                false,
		RuntimeDiagnosticsPersisted:                       false,
		DryRunResultPersisted:                             false,
		RawResultExposed:                                  false,
		DispatchDryRunExecuted:                            false,
		RequestObjectCreationEnabled:                      false,
		RequestObjectDispatchEnabled:                      false,
		PortalRequestCreated:                              false,
		NotificationActionEnabled:                         false,
		CompatibilityCenterOpened:                         false,
		SupportBundleExported:                             false,
		SupportCaseCreated:                                false,
		RuntimeOwned:                                      true,
		GoRuntimeBacked:                                   true,
		KDEPolicyOwner:                                    false,
		ProductionReadiness:                               false,
		ProductionOwnershipReady:                          false,
		SystemServiceStarted:                              false,
		SessionBusClaimed:                                 false,
		ProductionBusClaimed:                              false,
		ProductionOwnerEnabled:                            false,
		WriteMethodsEnabled:                               false,
		RuntimeWritesEnabled:                              false,
		DesktopFilesWritten:                               false,
		SettingsPersisted:                                 false,
		AdapterInvocationEnabled:                          false,
		BackendLaunchEnabled:                              false,
		BackendProcessStarted:                             false,
		SnapshotRestoreExecuted:                           false,
		StateCleanupExecuted:                              false,
		NetworkRequired:                                   false,
		HostRootModified:                                  false,
		PrivilegedContainerRequired:                       false,
		CallerStateRootRequired:                           false,
		StateRootPathExposed:                              false,
		FilePathsExposed:                                  false,
		FileContentRead:                                   false,
		RawCommandExposed:                                 false,
		RawExecutableExposed:                              false,
		BackendDetailsExposed:                             false,
		BlockedActions: []string{
			"treat this gate audit as permission to grant writer authorization or enable redacted status writers",
			"accept or consume real writer authorization receipts",
			"persist redacted status summaries, KDE status, Runtime diagnostics, dry-run results, or lookup records",
			"create request objects, dispatch actions, create Portal requests, send notifications, open Compatibility Center, export support bundles, or create support cases",
			"claim production ownership, start services, enable Runtime writes, launch compatibility engines, expose raw results or paths, or mutate host root",
		},
		NextRequirements: []string{
			"Implement a separate writer grant audit before any accepted receipt can grant redacted status writer authorization.",
			"Keep status writers and status persistence writes disabled until writer grants are explicitly modeled and consumed.",
			"Keep consumer enablement, lookup routes, dry-run execution, result persistence, request creation, dispatch, Portal requests, notifications, Compatibility Center navigation, and support writes disabled.",
			"Keep production ownership, Runtime writes, desktop side effects, engine launch, raw result exposure, and host mutation disabled.",
		},
		DesktopSafeSummary: "The KDE-safe redacted status writer accepted receipt gate audit consumes the Xnix current mainline and the writer authorization receipt acceptance audit, models the gate required before accepted opaque writer receipts can grant writer authorization, and keeps receipt acceptance, writer grants, status writes, desktop side effects, unsafe data exposure, and host mutation disabled.",
	}
	checks := productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAcceptedReceiptGateChecks(preview)
	preview.Checks = checks
	preview.CheckIDs = productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAcceptedReceiptGateCheckIDs(checks)
	preview.Counts = countProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAcceptedReceiptGateChecks(checks)
	if preview.Counts.Blocked == 0 {
		preview.AuditDecision = "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-writer-accepted-receipt-gate-audit-ready-writer-disabled"
	}
	if err := validateNoBackendTerms(preview, "production receipt notification action dry-run result lookup consumer enablement kde safe redacted status writer accepted receipt gate audit preview"); err != nil {
		return ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAcceptedReceiptGateAuditPreview{}, err
	}
	return preview, nil
}

type productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAcceptedReceiptGateSourceSet struct {
	CurrentMainline string
	AcceptanceAudit string
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAcceptedReceiptGateSources(root string) productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAcceptedReceiptGateSourceSet {
	return productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAcceptedReceiptGateSourceSet{
		CurrentMainline: productionAuthorizationReadSources(root, []string{"docs/xnix-current-mainline.md"}),
		AcceptanceAudit: productionAuthorizationReadSources(root, []string{"internal/runtime/owner/production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_kde_safe_redacted_status_writer_authorization_receipt_acceptance_audit.go", "internal/runtime/owner/production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_kde_safe_redacted_status_writer_authorization_receipt_acceptance_audit_test.go"}),
	}
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAcceptedReceiptGateMainlineReady(source string) bool {
	return productionAuthorizationHasAll(source, []string{"Xnix Current Mainline", "redacted status writer accepted receipt gate audit", "redacted status writer authorization receipt acceptance audit", "No Runtime write methods", "No KDE configuration writes"})
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAcceptedReceiptGateAcceptanceReady(source string) bool {
	return productionAuthorizationHasAll(source, []string{"production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-writer-authorization-receipt-acceptance-audit-preview", "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-writer-authorization-receipt-acceptance-audit-ready-acceptance-disabled", "AcceptanceAuthorizationReady", "ReceiptAcceptanceBoundaryReady", "ReceiptAccepted", "WriterAuthorizationGranted", "StatusWriterEnabled", "CompatibilityCenterAcceptanceModeled", "RuntimeDiagnosticsAcceptanceModeled"})
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAcceptedReceiptGateItems(mainlineReady bool, acceptanceReady bool) []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAcceptedReceiptGateAuditItem {
	return []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAcceptedReceiptGateAuditItem{
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAcceptedReceiptGateItem("review-receipt-dry-run-result-lookup-consumer-enablement-kde-safe-compatibility-center-redacted-status-writer-accepted-receipt-gate", "review", "compatibility-center", "review-result-compatibility-center-status", "review-compatibility-center-redacted-status-writer-accepted-receipt-gate-id", mainlineReady, acceptanceReady),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAcceptedReceiptGateItem("review-receipt-dry-run-result-lookup-consumer-enablement-kde-safe-runtime-diagnostics-redacted-status-writer-accepted-receipt-gate", "review", "runtime-diagnostics", "review-result-runtime-diagnostics-status", "review-runtime-diagnostics-redacted-status-writer-accepted-receipt-gate-id", mainlineReady, acceptanceReady),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAcceptedReceiptGateItem("renew-receipt-dry-run-result-lookup-consumer-enablement-kde-safe-compatibility-center-redacted-status-writer-accepted-receipt-gate", "renew", "compatibility-center", "renewal-result-compatibility-center-status", "renewal-compatibility-center-redacted-status-writer-accepted-receipt-gate-id", mainlineReady, acceptanceReady),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAcceptedReceiptGateItem("renew-receipt-dry-run-result-lookup-consumer-enablement-kde-safe-runtime-diagnostics-redacted-status-writer-accepted-receipt-gate", "renew", "runtime-diagnostics", "renewal-result-runtime-diagnostics-status", "renewal-runtime-diagnostics-redacted-status-writer-accepted-receipt-gate-id", mainlineReady, acceptanceReady),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAcceptedReceiptGateItem("open-compatibility-center-dry-run-result-lookup-consumer-enablement-kde-safe-compatibility-center-redacted-status-writer-accepted-receipt-gate", "open-compatibility-center", "compatibility-center", "navigation-result-compatibility-center-status", "navigation-compatibility-center-redacted-status-writer-accepted-receipt-gate-id", mainlineReady, acceptanceReady),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAcceptedReceiptGateItem("open-compatibility-center-dry-run-result-lookup-consumer-enablement-kde-safe-runtime-diagnostics-redacted-status-writer-accepted-receipt-gate", "open-compatibility-center", "runtime-diagnostics", "navigation-result-runtime-diagnostics-status", "navigation-runtime-diagnostics-redacted-status-writer-accepted-receipt-gate-id", mainlineReady, acceptanceReady),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAcceptedReceiptGateItem("dismiss-receipt-dry-run-result-lookup-consumer-enablement-kde-safe-compatibility-center-redacted-status-writer-accepted-receipt-gate", "dismiss", "compatibility-center", "dismissal-result-compatibility-center-status", "dismissal-compatibility-center-redacted-status-writer-accepted-receipt-gate-id", mainlineReady, acceptanceReady),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAcceptedReceiptGateItem("dismiss-receipt-dry-run-result-lookup-consumer-enablement-kde-safe-runtime-diagnostics-redacted-status-writer-accepted-receipt-gate", "dismiss", "runtime-diagnostics", "dismissal-result-runtime-diagnostics-status", "dismissal-runtime-diagnostics-redacted-status-writer-accepted-receipt-gate-id", mainlineReady, acceptanceReady),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAcceptedReceiptGateItem("support-info-dry-run-result-lookup-consumer-enablement-kde-safe-compatibility-center-redacted-status-writer-accepted-receipt-gate", "support-info", "compatibility-center", "support-info-result-compatibility-center-status", "support-info-compatibility-center-redacted-status-writer-accepted-receipt-gate-id", mainlineReady, acceptanceReady),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAcceptedReceiptGateItem("support-info-dry-run-result-lookup-consumer-enablement-kde-safe-runtime-diagnostics-redacted-status-writer-accepted-receipt-gate", "support-info", "runtime-diagnostics", "support-info-result-runtime-diagnostics-status", "support-info-runtime-diagnostics-redacted-status-writer-accepted-receipt-gate-id", mainlineReady, acceptanceReady),
	}
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAcceptedReceiptGateItem(id string, actionKind string, surfaceKind string, statusConsumerKind string, opaqueReceiptID string, mainlineReady bool, acceptanceReady bool) ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAcceptedReceiptGateAuditItem {
	ready := mainlineReady && acceptanceReady
	status := "missing-redacted-status-writer-accepted-receipt-gate-evidence"
	if ready {
		status = "redacted-status-writer-accepted-receipt-gate-modeled-writer-disabled"
	}
	return ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAcceptedReceiptGateAuditItem{
		ID:                      id,
		ActionKind:              actionKind,
		SurfaceKind:             surfaceKind,
		StatusConsumerKind:      statusConsumerKind,
		OpaqueReceiptID:         opaqueReceiptID,
		EvidencePresent:         ready,
		CurrentMainlineConsumed: mainlineReady,
		WriterAuthorizationReceiptAcceptanceConsumed: acceptanceReady,
		AcceptedReceiptGateModeled:                   ready,
		AcceptedReceiptGateBoundaryReady:             ready,
		OpaqueWriterAuthorizationReceipt:             ready,
		KDESafeRedactedStatusOnly:                    ready,
		ReceiptPresent:                               false,
		ReceiptAccepted:                              false,
		AuthorizationAccepted:                        false,
		WriterAuthorizationGranted:                   false,
		StatusWriterEnabled:                          false,
		StatusPersistenceWriteEnabled:                false,
		KDEStatusWriteEnabled:                        false,
		RuntimeDiagnosticsWriteEnabled:               false,
		RedactedSummaryPersisted:                     false,
		KDEStatusPersisted:                           false,
		RuntimeDiagnosticsPersisted:                  false,
		RawResultExposed:                             false,
		UserVisible:                                  ready,
		ReviewOnly:                                   true,
		RuntimeOwned:                                 true,
		GoRuntimeBacked:                              true,
		KDEPolicyOwner:                               false,
		SideEffectsDisabled:                          true,
		HostRootModified:                             false,
		InternalDetailsExposed:                       false,
		AcceptedReceiptGateStatus:                    status,
		NextRequirement:                              "require separate writer grant audit before enabling this redacted status writer",
	}
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAcceptedReceiptGateChecks(preview ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAcceptedReceiptGateAuditPreview) []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAcceptedReceiptGateAuditCheck {
	return []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAcceptedReceiptGateAuditCheck{
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAcceptedReceiptGateCheck("current-mainline-consumed", productionAuthorizationPassBlocked(preview.CurrentMainlineConsumed), "The accepted receipt gate audit consumes the Xnix current mainline."),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAcceptedReceiptGateCheck("writer-authorization-receipt-acceptance-audit-consumed", productionAuthorizationPassBlocked(preview.WriterAuthorizationReceiptAcceptanceAuditConsumed && preview.WriterAuthorizationReceiptAcceptanceReady), "The accepted receipt gate audit consumes the writer authorization receipt acceptance audit."),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAcceptedReceiptGateCheck("accepted-receipt-gate-modeled", productionAuthorizationPassBlocked(preview.AcceptedReceiptGateRequired && preview.AcceptedReceiptGateModeled && preview.AcceptedReceiptGateReady && preview.AcceptedReceiptGateBoundaryReady && preview.OpaqueWriterAuthorizationReceipt), "The accepted receipt gate boundary is modeled."),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAcceptedReceiptGateCheck("compatibility-center-and-runtime-gates-modeled", productionAuthorizationPassBlocked(preview.CompatibilityCenterGateModeled && preview.RuntimeDiagnosticsGateModeled && preview.CompatibilityCenterGateItemCount == 5 && preview.RuntimeDiagnosticsGateItemCount == 5), "Compatibility Center and Runtime diagnostics accepted receipt gates are both modeled."),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAcceptedReceiptGateCheck("ten-gate-items-ready-writer-disabled", productionAuthorizationPassBlocked(preview.GateItemCount == 10 && preview.RequiredGateItemCount == 10 && preview.ReadyGateItemCount == 10 && preview.MissingGateItemCount == 0 && preview.AcceptedReceiptItemCount == 0 && preview.GrantedWriterItemCount == 0 && preview.EnabledWriterItemCount == 0), "All accepted receipt gate candidates are present while writer grants remain disabled."),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAcceptedReceiptGateCheck("receipt-writers-and-persistence-disabled", productionAuthorizationPassBlocked(!preview.ReceiptPresent && !preview.ReceiptAccepted && !preview.AuthorizationAccepted && !preview.WriterAuthorizationGranted && !preview.StatusWriterEnabled && !preview.StatusPersistenceWriteEnabled && !preview.KDEStatusWriteEnabled && !preview.RuntimeDiagnosticsWriteEnabled && !preview.RedactedSummaryPersisted && !preview.KDEStatusPersisted && !preview.RuntimeDiagnosticsPersisted && !preview.DryRunResultPersisted), "Receipt acceptance, writer authorization grants, status writers, and status persistence remain disabled."),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAcceptedReceiptGateCheck("consumer-lookup-request-and-support-disabled", productionAuthorizationPassBlocked(!preview.ConsumerConsumptionAuthorized && !preview.ConsumerEnablementAuthorized && !preview.KDEConsumerEnabled && !preview.RuntimeConsumerEnabled && !preview.LookupRouteAuthorized && !preview.LookupRouteEnabled && !preview.OpaqueLookupEnabled && !preview.RequestObjectCreationEnabled && !preview.RequestObjectDispatchEnabled && !preview.PortalRequestCreated && !preview.NotificationActionEnabled && !preview.CompatibilityCenterOpened && !preview.SupportBundleExported && !preview.SupportCaseCreated), "Consumers, lookup, request objects, desktop actions, and support writes remain disabled."),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAcceptedReceiptGateCheck("production-and-host-boundary-closed", productionAuthorizationPassBlocked(!preview.SystemServiceStarted && !preview.SessionBusClaimed && !preview.ProductionBusClaimed && !preview.ProductionOwnerEnabled && !preview.ProductionReadiness && !preview.ProductionOwnershipReady && !preview.WriteMethodsEnabled && !preview.RuntimeWritesEnabled && !preview.AdapterInvocationEnabled && !preview.BackendLaunchEnabled && !preview.BackendProcessStarted && !preview.NetworkRequired && !preview.HostRootModified && !preview.PrivilegedContainerRequired && !preview.CallerStateRootRequired && !preview.StateRootPathExposed && !preview.FilePathsExposed && !preview.FileContentRead && !preview.RawResultExposed && !preview.RawCommandExposed && !preview.RawExecutableExposed && !preview.BackendDetailsExposed && preview.RawExposedGateItemCount == 0 && preview.SideEffectGateItemCount == 0 && productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAcceptedReceiptGateItemsKeepClosed(preview.GateItems)), "Production ownership, writes, unsafe data exposure, network, privilege, internal details, and host mutation remain disabled."),
	}
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAcceptedReceiptGateCheck(id string, status string, summary string) ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAcceptedReceiptGateAuditCheck {
	return ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAcceptedReceiptGateAuditCheck{ID: id, Status: status, Summary: summary}
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAcceptedReceiptGateReadyCount(items []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAcceptedReceiptGateAuditItem) int {
	count := 0
	for _, item := range items {
		if item.EvidencePresent {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAcceptedReceiptGateSurfaceCount(items []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAcceptedReceiptGateAuditItem, surfaceKind string) int {
	count := 0
	for _, item := range items {
		if item.SurfaceKind == surfaceKind {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAcceptedReceiptGateItemsKeepClosed(items []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAcceptedReceiptGateAuditItem) bool {
	for _, item := range items {
		if item.ReceiptPresent || item.ReceiptAccepted || item.AuthorizationAccepted || item.WriterAuthorizationGranted || item.StatusWriterEnabled || item.StatusPersistenceWriteEnabled || item.KDEStatusWriteEnabled || item.RuntimeDiagnosticsWriteEnabled || item.RedactedSummaryPersisted || item.KDEStatusPersisted || item.RuntimeDiagnosticsPersisted || item.RawResultExposed || !item.SideEffectsDisabled || item.HostRootModified || item.InternalDetailsExposed {
			return false
		}
	}
	return true
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAcceptedReceiptGateItemIDs(items []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAcceptedReceiptGateAuditItem) []string {
	ids := make([]string, 0, len(items))
	for _, item := range items {
		ids = append(ids, item.ID)
	}
	return ids
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAcceptedReceiptGateCheckIDs(checks []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAcceptedReceiptGateAuditCheck) []string {
	ids := make([]string, 0, len(checks))
	for _, check := range checks {
		ids = append(ids, check.ID)
	}
	return ids
}

func countProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAcceptedReceiptGateChecks(checks []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAcceptedReceiptGateAuditCheck) ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAcceptedReceiptGateAuditCheckCounts {
	counts := ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAcceptedReceiptGateAuditCheckCounts{Total: len(checks)}
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
