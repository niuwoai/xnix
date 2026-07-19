package owner

type ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterGrantAuditPreview struct {
	Version                           string                                                                                                                  `json:"version"`
	SchemaVersion                     string                                                                                                                  `json:"schema_version"`
	RequestType                       string                                                                                                                  `json:"request_type"`
	AuditType                         string                                                                                                                  `json:"audit_type"`
	Source                            string                                                                                                                  `json:"source"`
	AuditDecision                     string                                                                                                                  `json:"audit_decision"`
	CurrentMainlineConsumed           bool                                                                                                                    `json:"current_mainline_consumed"`
	AcceptedReceiptGateAuditConsumed  bool                                                                                                                    `json:"accepted_receipt_gate_audit_consumed"`
	AcceptedReceiptGateReady          bool                                                                                                                    `json:"accepted_receipt_gate_ready"`
	WriterGrantRequired               bool                                                                                                                    `json:"writer_grant_required"`
	WriterGrantModeled                bool                                                                                                                    `json:"writer_grant_modeled"`
	WriterGrantReady                  bool                                                                                                                    `json:"writer_grant_ready"`
	WriterGrantBoundaryReady          bool                                                                                                                    `json:"writer_grant_boundary_ready"`
	OpaqueWriterAuthorizationReceipt  bool                                                                                                                    `json:"opaque_writer_authorization_receipt"`
	KDESafeRedactedStatusOnly         bool                                                                                                                    `json:"kde_safe_redacted_status_only"`
	CompatibilityCenterGrantModeled   bool                                                                                                                    `json:"compatibility_center_grant_modeled"`
	RuntimeDiagnosticsGrantModeled    bool                                                                                                                    `json:"runtime_diagnostics_grant_modeled"`
	ReceiptPresent                    bool                                                                                                                    `json:"receipt_present"`
	ReceiptAccepted                   bool                                                                                                                    `json:"receipt_accepted"`
	AuthorizationAccepted             bool                                                                                                                    `json:"authorization_accepted"`
	WriterAuthorizationGranted        bool                                                                                                                    `json:"writer_authorization_granted"`
	StatusWriterEnabled               bool                                                                                                                    `json:"status_writer_enabled"`
	StatusPersistenceWriteEnabled     bool                                                                                                                    `json:"status_persistence_write_enabled"`
	KDEStatusWriteEnabled             bool                                                                                                                    `json:"kde_status_write_enabled"`
	RuntimeDiagnosticsWriteEnabled    bool                                                                                                                    `json:"runtime_diagnostics_write_enabled"`
	GrantItemCount                    int                                                                                                                     `json:"grant_item_count"`
	RequiredGrantItemCount            int                                                                                                                     `json:"required_grant_item_count"`
	ReadyGrantItemCount               int                                                                                                                     `json:"ready_grant_item_count"`
	MissingGrantItemCount             int                                                                                                                     `json:"missing_grant_item_count"`
	AcceptedReceiptItemCount          int                                                                                                                     `json:"accepted_receipt_item_count"`
	GrantedWriterItemCount            int                                                                                                                     `json:"granted_writer_item_count"`
	EnabledWriterItemCount            int                                                                                                                     `json:"enabled_writer_item_count"`
	PersistedWriterItemCount          int                                                                                                                     `json:"persisted_writer_item_count"`
	RawExposedGrantItemCount          int                                                                                                                     `json:"raw_exposed_grant_item_count"`
	SideEffectGrantItemCount          int                                                                                                                     `json:"side_effect_grant_item_count"`
	CompatibilityCenterGrantItemCount int                                                                                                                     `json:"compatibility_center_grant_item_count"`
	RuntimeDiagnosticsGrantItemCount  int                                                                                                                     `json:"runtime_diagnostics_grant_item_count"`
	GrantItems                        []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterGrantAuditItem      `json:"grant_items"`
	GrantItemIDs                      []string                                                                                                                `json:"grant_item_ids"`
	Checks                            []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterGrantAuditCheck     `json:"checks"`
	CheckIDs                          []string                                                                                                                `json:"check_ids"`
	Counts                            ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterGrantAuditCheckCounts `json:"counts"`
	ConsumerConsumptionAuthorized     bool                                                                                                                    `json:"consumer_consumption_authorized"`
	ConsumerEnablementAuthorized      bool                                                                                                                    `json:"consumer_enablement_authorized"`
	KDEConsumerEnabled                bool                                                                                                                    `json:"kde_consumer_enabled"`
	RuntimeConsumerEnabled            bool                                                                                                                    `json:"runtime_consumer_enabled"`
	LookupRouteAuthorized             bool                                                                                                                    `json:"lookup_route_authorized"`
	LookupRouteEnabled                bool                                                                                                                    `json:"lookup_route_enabled"`
	OpaqueLookupEnabled               bool                                                                                                                    `json:"opaque_lookup_enabled"`
	RedactedSummaryPersisted          bool                                                                                                                    `json:"redacted_summary_persisted"`
	KDEStatusPersisted                bool                                                                                                                    `json:"kde_status_persisted"`
	RuntimeDiagnosticsPersisted       bool                                                                                                                    `json:"runtime_diagnostics_persisted"`
	DryRunResultPersisted             bool                                                                                                                    `json:"dry_run_result_persisted"`
	RawResultExposed                  bool                                                                                                                    `json:"raw_result_exposed"`
	DispatchDryRunExecuted            bool                                                                                                                    `json:"dispatch_dry_run_executed"`
	RequestObjectCreationEnabled      bool                                                                                                                    `json:"request_object_creation_enabled"`
	RequestObjectDispatchEnabled      bool                                                                                                                    `json:"request_object_dispatch_enabled"`
	PortalRequestCreated              bool                                                                                                                    `json:"portal_request_created"`
	NotificationActionEnabled         bool                                                                                                                    `json:"notification_action_enabled"`
	CompatibilityCenterOpened         bool                                                                                                                    `json:"compatibility_center_opened"`
	SupportBundleExported             bool                                                                                                                    `json:"support_bundle_exported"`
	SupportCaseCreated                bool                                                                                                                    `json:"support_case_created"`
	RuntimeOwned                      bool                                                                                                                    `json:"runtime_owned"`
	GoRuntimeBacked                   bool                                                                                                                    `json:"go_runtime_backed"`
	KDEPolicyOwner                    bool                                                                                                                    `json:"kde_policy_owner"`
	ProductionReadiness               bool                                                                                                                    `json:"production_readiness"`
	ProductionOwnershipReady          bool                                                                                                                    `json:"production_ownership_ready"`
	SystemServiceStarted              bool                                                                                                                    `json:"system_service_started"`
	SessionBusClaimed                 bool                                                                                                                    `json:"session_bus_claimed"`
	ProductionBusClaimed              bool                                                                                                                    `json:"production_bus_claimed"`
	ProductionOwnerEnabled            bool                                                                                                                    `json:"production_owner_enabled"`
	WriteMethodsEnabled               bool                                                                                                                    `json:"write_methods_enabled"`
	RuntimeWritesEnabled              bool                                                                                                                    `json:"runtime_writes_enabled"`
	DesktopFilesWritten               bool                                                                                                                    `json:"desktop_files_written"`
	SettingsPersisted                 bool                                                                                                                    `json:"settings_persisted"`
	AdapterInvocationEnabled          bool                                                                                                                    `json:"adapter_invocation_enabled"`
	BackendLaunchEnabled              bool                                                                                                                    `json:"backend_launch_enabled"`
	BackendProcessStarted             bool                                                                                                                    `json:"backend_process_started"`
	SnapshotRestoreExecuted           bool                                                                                                                    `json:"snapshot_restore_executed"`
	StateCleanupExecuted              bool                                                                                                                    `json:"state_cleanup_executed"`
	NetworkRequired                   bool                                                                                                                    `json:"network_required"`
	HostRootModified                  bool                                                                                                                    `json:"host_root_modified"`
	PrivilegedContainerRequired       bool                                                                                                                    `json:"privileged_container_required"`
	CallerStateRootRequired           bool                                                                                                                    `json:"caller_state_root_required"`
	StateRootPathExposed              bool                                                                                                                    `json:"state_root_path_exposed"`
	FilePathsExposed                  bool                                                                                                                    `json:"file_paths_exposed"`
	FileContentRead                   bool                                                                                                                    `json:"file_content_read"`
	RawCommandExposed                 bool                                                                                                                    `json:"raw_command_exposed"`
	RawExecutableExposed              bool                                                                                                                    `json:"raw_executable_exposed"`
	BackendDetailsExposed             bool                                                                                                                    `json:"backend_details_exposed"`
	BlockedActions                    []string                                                                                                                `json:"blocked_actions"`
	NextRequirements                  []string                                                                                                                `json:"next_requirements"`
	DesktopSafeSummary                string                                                                                                                  `json:"desktop_safe_summary"`
}

type ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterGrantAuditItem struct {
	ID                               string `json:"id"`
	ActionKind                       string `json:"action_kind"`
	SurfaceKind                      string `json:"surface_kind"`
	StatusConsumerKind               string `json:"status_consumer_kind"`
	OpaqueReceiptID                  string `json:"opaque_receipt_id"`
	EvidencePresent                  bool   `json:"evidence_present"`
	CurrentMainlineConsumed          bool   `json:"current_mainline_consumed"`
	AcceptedReceiptGateAuditConsumed bool   `json:"accepted_receipt_gate_audit_consumed"`
	WriterGrantModeled               bool   `json:"writer_grant_modeled"`
	WriterGrantBoundaryReady         bool   `json:"writer_grant_boundary_ready"`
	OpaqueWriterAuthorizationReceipt bool   `json:"opaque_writer_authorization_receipt"`
	KDESafeRedactedStatusOnly        bool   `json:"kde_safe_redacted_status_only"`
	ReceiptPresent                   bool   `json:"receipt_present"`
	ReceiptAccepted                  bool   `json:"receipt_accepted"`
	AuthorizationAccepted            bool   `json:"authorization_accepted"`
	WriterAuthorizationGranted       bool   `json:"writer_authorization_granted"`
	StatusWriterEnabled              bool   `json:"status_writer_enabled"`
	StatusPersistenceWriteEnabled    bool   `json:"status_persistence_write_enabled"`
	KDEStatusWriteEnabled            bool   `json:"kde_status_write_enabled"`
	RuntimeDiagnosticsWriteEnabled   bool   `json:"runtime_diagnostics_write_enabled"`
	RedactedSummaryPersisted         bool   `json:"redacted_summary_persisted"`
	KDEStatusPersisted               bool   `json:"kde_status_persisted"`
	RuntimeDiagnosticsPersisted      bool   `json:"runtime_diagnostics_persisted"`
	RawResultExposed                 bool   `json:"raw_result_exposed"`
	UserVisible                      bool   `json:"user_visible"`
	ReviewOnly                       bool   `json:"review_only"`
	RuntimeOwned                     bool   `json:"runtime_owned"`
	GoRuntimeBacked                  bool   `json:"go_runtime_backed"`
	KDEPolicyOwner                   bool   `json:"kde_policy_owner"`
	SideEffectsDisabled              bool   `json:"side_effects_disabled"`
	HostRootModified                 bool   `json:"host_root_modified"`
	InternalDetailsExposed           bool   `json:"internal_details_exposed"`
	WriterGrantStatus                string `json:"writer_grant_status"`
	NextRequirement                  string `json:"next_requirement"`
}

type ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterGrantAuditCheck struct {
	ID      string `json:"id"`
	Status  string `json:"status"`
	Summary string `json:"summary"`
}

type ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterGrantAuditCheckCounts struct {
	Total   int `json:"total"`
	Passed  int `json:"passed"`
	Pending int `json:"pending"`
	Blocked int `json:"blocked"`
}

func NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterGrantAuditPreview(root string) (ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterGrantAuditPreview, error) {
	version, err := readProductionDBusGateReviewVersion(root)
	if err != nil {
		return ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterGrantAuditPreview{}, err
	}
	sources := productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterGrantSources(root)
	mainlineReady := productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterGrantMainlineReady(sources.CurrentMainline)
	acceptedGateReady := productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterGrantAcceptedGateReady(sources.AcceptedReceiptGateAudit)
	items := productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterGrantItems(mainlineReady, acceptedGateReady)
	readyItemCount := productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterGrantReadyCount(items)
	grantReady := mainlineReady && acceptedGateReady && readyItemCount == len(items)
	preview := ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterGrantAuditPreview{
		Version:                           version,
		SchemaVersion:                     "xnix.runtime.production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_kde_safe_redacted_status_writer_grant_audit.v1",
		RequestType:                       "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-writer-grant-audit-preview",
		AuditType:                         "receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-writer-grant-audit",
		Source:                            "xnix-current-mainline+redacted-status-writer-accepted-receipt-gate",
		AuditDecision:                     "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-writer-grant-audit-blocked",
		CurrentMainlineConsumed:           mainlineReady,
		AcceptedReceiptGateAuditConsumed:  acceptedGateReady,
		AcceptedReceiptGateReady:          acceptedGateReady,
		WriterGrantRequired:               true,
		WriterGrantModeled:                true,
		WriterGrantReady:                  grantReady,
		WriterGrantBoundaryReady:          grantReady,
		OpaqueWriterAuthorizationReceipt:  acceptedGateReady,
		KDESafeRedactedStatusOnly:         grantReady,
		CompatibilityCenterGrantModeled:   grantReady,
		RuntimeDiagnosticsGrantModeled:    grantReady,
		ReceiptPresent:                    false,
		ReceiptAccepted:                   false,
		AuthorizationAccepted:             false,
		WriterAuthorizationGranted:        false,
		StatusWriterEnabled:               false,
		StatusPersistenceWriteEnabled:     false,
		KDEStatusWriteEnabled:             false,
		RuntimeDiagnosticsWriteEnabled:    false,
		GrantItemCount:                    len(items),
		RequiredGrantItemCount:            len(items),
		ReadyGrantItemCount:               readyItemCount,
		MissingGrantItemCount:             len(items) - readyItemCount,
		AcceptedReceiptItemCount:          0,
		GrantedWriterItemCount:            0,
		EnabledWriterItemCount:            0,
		PersistedWriterItemCount:          0,
		RawExposedGrantItemCount:          0,
		SideEffectGrantItemCount:          0,
		CompatibilityCenterGrantItemCount: productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterGrantSurfaceCount(items, "compatibility-center"),
		RuntimeDiagnosticsGrantItemCount:  productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterGrantSurfaceCount(items, "runtime-diagnostics"),
		GrantItems:                        items,
		GrantItemIDs:                      productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterGrantItemIDs(items),
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
			"treat this grant audit as permission to enable redacted status writers or write status records",
			"accept or consume real writer authorization receipts",
			"persist redacted status summaries, KDE status, Runtime diagnostics, dry-run results, or lookup records",
			"create request objects, dispatch actions, create Portal requests, send notifications, open Compatibility Center, export support bundles, or create support cases",
			"claim production ownership, start services, enable Runtime writes, launch compatibility engines, expose raw results or paths, or mutate host root",
		},
		NextRequirements: []string{
			"Implement a separate status writer enablement audit before any writer grant can enable a redacted status writer.",
			"Keep status writers and status persistence writes disabled until writer grants are explicitly consumed by a writer enablement boundary.",
			"Keep consumer enablement, lookup routes, dry-run execution, result persistence, request creation, dispatch, Portal requests, notifications, Compatibility Center navigation, and support writes disabled.",
			"Keep production ownership, Runtime writes, desktop side effects, engine launch, raw result exposure, and host mutation disabled.",
		},
		DesktopSafeSummary: "The KDE-safe redacted status writer grant audit consumes the Xnix current mainline and the accepted receipt gate audit, models the authorization boundary required before accepted opaque writer receipts can grant writer authorization, and keeps writer grants, status writers, status writes, desktop side effects, unsafe data exposure, and host mutation disabled.",
	}
	checks := productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterGrantChecks(preview)
	preview.Checks = checks
	preview.CheckIDs = productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterGrantCheckIDs(checks)
	preview.Counts = countProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterGrantChecks(checks)
	if preview.Counts.Blocked == 0 {
		preview.AuditDecision = "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-writer-grant-audit-ready-writer-disabled"
	}
	if err := validateNoBackendTerms(preview, "production receipt notification action dry-run result lookup consumer enablement kde safe redacted status writer grant audit preview"); err != nil {
		return ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterGrantAuditPreview{}, err
	}
	return preview, nil
}

type productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterGrantSourceSet struct {
	CurrentMainline          string
	AcceptedReceiptGateAudit string
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterGrantSources(root string) productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterGrantSourceSet {
	return productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterGrantSourceSet{
		CurrentMainline:          productionAuthorizationReadSources(root, []string{"docs/xnix-current-mainline.md"}),
		AcceptedReceiptGateAudit: productionAuthorizationReadSources(root, []string{"internal/runtime/owner/production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_kde_safe_redacted_status_writer_accepted_receipt_gate_audit.go", "internal/runtime/owner/production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_kde_safe_redacted_status_writer_accepted_receipt_gate_audit_test.go"}),
	}
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterGrantMainlineReady(source string) bool {
	return productionAuthorizationHasAll(source, []string{"Xnix Current Mainline", "redacted status writer grant audit", "redacted status writer accepted receipt gate audit", "No Runtime write methods", "No KDE configuration writes"})
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterGrantAcceptedGateReady(source string) bool {
	return productionAuthorizationHasAll(source, []string{"production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-writer-accepted-receipt-gate-audit-preview", "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-writer-accepted-receipt-gate-audit-ready-writer-disabled", "AcceptedReceiptGateReady", "AcceptedReceiptGateBoundaryReady", "OpaqueWriterAuthorizationReceipt", "WriterAuthorizationGranted", "StatusWriterEnabled", "CompatibilityCenterGateModeled", "RuntimeDiagnosticsGateModeled"})
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterGrantItems(mainlineReady bool, acceptedGateReady bool) []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterGrantAuditItem {
	return []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterGrantAuditItem{
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterGrantItem("review-receipt-dry-run-result-lookup-consumer-enablement-kde-safe-compatibility-center-redacted-status-writer-grant", "review", "compatibility-center", "review-result-compatibility-center-status", "review-compatibility-center-redacted-status-writer-grant-id", mainlineReady, acceptedGateReady),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterGrantItem("review-receipt-dry-run-result-lookup-consumer-enablement-kde-safe-runtime-diagnostics-redacted-status-writer-grant", "review", "runtime-diagnostics", "review-result-runtime-diagnostics-status", "review-runtime-diagnostics-redacted-status-writer-grant-id", mainlineReady, acceptedGateReady),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterGrantItem("renew-receipt-dry-run-result-lookup-consumer-enablement-kde-safe-compatibility-center-redacted-status-writer-grant", "renew", "compatibility-center", "renewal-result-compatibility-center-status", "renewal-compatibility-center-redacted-status-writer-grant-id", mainlineReady, acceptedGateReady),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterGrantItem("renew-receipt-dry-run-result-lookup-consumer-enablement-kde-safe-runtime-diagnostics-redacted-status-writer-grant", "renew", "runtime-diagnostics", "renewal-result-runtime-diagnostics-status", "renewal-runtime-diagnostics-redacted-status-writer-grant-id", mainlineReady, acceptedGateReady),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterGrantItem("open-compatibility-center-dry-run-result-lookup-consumer-enablement-kde-safe-compatibility-center-redacted-status-writer-grant", "open-compatibility-center", "compatibility-center", "navigation-result-compatibility-center-status", "navigation-compatibility-center-redacted-status-writer-grant-id", mainlineReady, acceptedGateReady),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterGrantItem("open-compatibility-center-dry-run-result-lookup-consumer-enablement-kde-safe-runtime-diagnostics-redacted-status-writer-grant", "open-compatibility-center", "runtime-diagnostics", "navigation-result-runtime-diagnostics-status", "navigation-runtime-diagnostics-redacted-status-writer-grant-id", mainlineReady, acceptedGateReady),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterGrantItem("dismiss-receipt-dry-run-result-lookup-consumer-enablement-kde-safe-compatibility-center-redacted-status-writer-grant", "dismiss", "compatibility-center", "dismissal-result-compatibility-center-status", "dismissal-compatibility-center-redacted-status-writer-grant-id", mainlineReady, acceptedGateReady),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterGrantItem("dismiss-receipt-dry-run-result-lookup-consumer-enablement-kde-safe-runtime-diagnostics-redacted-status-writer-grant", "dismiss", "runtime-diagnostics", "dismissal-result-runtime-diagnostics-status", "dismissal-runtime-diagnostics-redacted-status-writer-grant-id", mainlineReady, acceptedGateReady),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterGrantItem("support-info-dry-run-result-lookup-consumer-enablement-kde-safe-compatibility-center-redacted-status-writer-grant", "support-info", "compatibility-center", "support-info-result-compatibility-center-status", "support-info-compatibility-center-redacted-status-writer-grant-id", mainlineReady, acceptedGateReady),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterGrantItem("support-info-dry-run-result-lookup-consumer-enablement-kde-safe-runtime-diagnostics-redacted-status-writer-grant", "support-info", "runtime-diagnostics", "support-info-result-runtime-diagnostics-status", "support-info-runtime-diagnostics-redacted-status-writer-grant-id", mainlineReady, acceptedGateReady),
	}
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterGrantItem(id string, actionKind string, surfaceKind string, statusConsumerKind string, opaqueReceiptID string, mainlineReady bool, acceptedGateReady bool) ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterGrantAuditItem {
	ready := mainlineReady && acceptedGateReady
	status := "missing-redacted-status-writer-grant-evidence"
	if ready {
		status = "redacted-status-writer-grant-modeled-writer-disabled"
	}
	return ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterGrantAuditItem{
		ID:                               id,
		ActionKind:                       actionKind,
		SurfaceKind:                      surfaceKind,
		StatusConsumerKind:               statusConsumerKind,
		OpaqueReceiptID:                  opaqueReceiptID,
		EvidencePresent:                  ready,
		CurrentMainlineConsumed:          mainlineReady,
		AcceptedReceiptGateAuditConsumed: acceptedGateReady,
		WriterGrantModeled:               ready,
		WriterGrantBoundaryReady:         ready,
		OpaqueWriterAuthorizationReceipt: acceptedGateReady,
		KDESafeRedactedStatusOnly:        ready,
		ReceiptPresent:                   false,
		ReceiptAccepted:                  false,
		AuthorizationAccepted:            false,
		WriterAuthorizationGranted:       false,
		StatusWriterEnabled:              false,
		StatusPersistenceWriteEnabled:    false,
		KDEStatusWriteEnabled:            false,
		RuntimeDiagnosticsWriteEnabled:   false,
		RedactedSummaryPersisted:         false,
		KDEStatusPersisted:               false,
		RuntimeDiagnosticsPersisted:      false,
		RawResultExposed:                 false,
		UserVisible:                      ready,
		ReviewOnly:                       true,
		RuntimeOwned:                     true,
		GoRuntimeBacked:                  true,
		KDEPolicyOwner:                   false,
		SideEffectsDisabled:              true,
		HostRootModified:                 false,
		InternalDetailsExposed:           false,
		WriterGrantStatus:                status,
		NextRequirement:                  "require separate status writer enablement audit before enabling this redacted status writer",
	}
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterGrantChecks(preview ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterGrantAuditPreview) []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterGrantAuditCheck {
	return []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterGrantAuditCheck{
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterGrantCheck("current-mainline-consumed", productionAuthorizationPassBlocked(preview.CurrentMainlineConsumed), "The writer grant audit consumes the Xnix current mainline."),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterGrantCheck("accepted-receipt-gate-audit-consumed", productionAuthorizationPassBlocked(preview.AcceptedReceiptGateAuditConsumed && preview.AcceptedReceiptGateReady), "The writer grant audit consumes the accepted receipt gate audit."),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterGrantCheck("writer-grant-boundary-modeled", productionAuthorizationPassBlocked(preview.WriterGrantRequired && preview.WriterGrantModeled && preview.WriterGrantReady && preview.WriterGrantBoundaryReady), "The redacted status writer grant boundary is modeled."),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterGrantCheck("compatibility-center-and-runtime-grants-modeled", productionAuthorizationPassBlocked(preview.CompatibilityCenterGrantModeled && preview.RuntimeDiagnosticsGrantModeled && preview.CompatibilityCenterGrantItemCount == 5 && preview.RuntimeDiagnosticsGrantItemCount == 5), "Compatibility Center and Runtime diagnostics writer grant candidates are both modeled."),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterGrantCheck("ten-grant-items-ready-writer-disabled", productionAuthorizationPassBlocked(preview.GrantItemCount == 10 && preview.RequiredGrantItemCount == 10 && preview.ReadyGrantItemCount == 10 && preview.MissingGrantItemCount == 0 && preview.GrantedWriterItemCount == 0 && preview.EnabledWriterItemCount == 0 && preview.PersistedWriterItemCount == 0), "All redacted status writer grant candidates are present while writers remain disabled."),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterGrantCheck("writer-grants-writers-and-persistence-disabled", productionAuthorizationPassBlocked(!preview.ReceiptPresent && !preview.ReceiptAccepted && !preview.AuthorizationAccepted && !preview.WriterAuthorizationGranted && !preview.StatusWriterEnabled && !preview.StatusPersistenceWriteEnabled && !preview.KDEStatusWriteEnabled && !preview.RuntimeDiagnosticsWriteEnabled && !preview.RedactedSummaryPersisted && !preview.KDEStatusPersisted && !preview.RuntimeDiagnosticsPersisted && !preview.DryRunResultPersisted), "Writer authorization grants, status writers, and status persistence remain disabled."),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterGrantCheck("consumer-lookup-request-and-support-disabled", productionAuthorizationPassBlocked(!preview.ConsumerConsumptionAuthorized && !preview.ConsumerEnablementAuthorized && !preview.KDEConsumerEnabled && !preview.RuntimeConsumerEnabled && !preview.LookupRouteAuthorized && !preview.LookupRouteEnabled && !preview.OpaqueLookupEnabled && !preview.RequestObjectCreationEnabled && !preview.RequestObjectDispatchEnabled && !preview.PortalRequestCreated && !preview.NotificationActionEnabled && !preview.CompatibilityCenterOpened && !preview.SupportBundleExported && !preview.SupportCaseCreated), "Consumers, lookup, request objects, desktop actions, and support writes remain disabled."),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterGrantCheck("production-and-host-boundary-closed", productionAuthorizationPassBlocked(!preview.SystemServiceStarted && !preview.SessionBusClaimed && !preview.ProductionBusClaimed && !preview.ProductionOwnerEnabled && !preview.ProductionReadiness && !preview.ProductionOwnershipReady && !preview.WriteMethodsEnabled && !preview.RuntimeWritesEnabled && !preview.AdapterInvocationEnabled && !preview.BackendLaunchEnabled && !preview.BackendProcessStarted && !preview.NetworkRequired && !preview.HostRootModified && !preview.PrivilegedContainerRequired && !preview.CallerStateRootRequired && !preview.StateRootPathExposed && !preview.FilePathsExposed && !preview.FileContentRead && !preview.RawResultExposed && !preview.RawCommandExposed && !preview.RawExecutableExposed && !preview.BackendDetailsExposed && preview.RawExposedGrantItemCount == 0 && preview.SideEffectGrantItemCount == 0 && productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterGrantItemsKeepClosed(preview.GrantItems)), "Production ownership, writes, unsafe data exposure, network, privilege, internal details, and host mutation remain disabled."),
	}
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterGrantCheck(id string, status string, summary string) ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterGrantAuditCheck {
	return ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterGrantAuditCheck{ID: id, Status: status, Summary: summary}
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterGrantReadyCount(items []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterGrantAuditItem) int {
	count := 0
	for _, item := range items {
		if item.EvidencePresent && item.WriterGrantBoundaryReady && !item.WriterAuthorizationGranted && !item.StatusWriterEnabled && !item.HostRootModified && !item.InternalDetailsExposed {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterGrantSurfaceCount(items []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterGrantAuditItem, surfaceKind string) int {
	count := 0
	for _, item := range items {
		if item.SurfaceKind == surfaceKind {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterGrantItemsKeepClosed(items []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterGrantAuditItem) bool {
	for _, item := range items {
		if item.ReceiptPresent || item.ReceiptAccepted || item.AuthorizationAccepted || item.WriterAuthorizationGranted || item.StatusWriterEnabled || item.StatusPersistenceWriteEnabled || item.KDEStatusWriteEnabled || item.RuntimeDiagnosticsWriteEnabled || item.RedactedSummaryPersisted || item.KDEStatusPersisted || item.RuntimeDiagnosticsPersisted || item.RawResultExposed || item.KDEPolicyOwner || !item.SideEffectsDisabled || item.HostRootModified || item.InternalDetailsExposed {
			return false
		}
	}
	return true
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterGrantItemIDs(items []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterGrantAuditItem) []string {
	ids := make([]string, 0, len(items))
	for _, item := range items {
		ids = append(ids, item.ID)
	}
	return ids
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterGrantCheckIDs(checks []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterGrantAuditCheck) []string {
	ids := make([]string, 0, len(checks))
	for _, check := range checks {
		ids = append(ids, check.ID)
	}
	return ids
}

func countProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterGrantChecks(checks []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterGrantAuditCheck) ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterGrantAuditCheckCounts {
	counts := ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterGrantAuditCheckCounts{Total: len(checks)}
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
