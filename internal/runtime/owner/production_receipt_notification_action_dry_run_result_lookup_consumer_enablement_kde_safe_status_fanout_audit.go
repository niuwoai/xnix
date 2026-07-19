package owner

type ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusFanOutAuditPreview struct {
	Version                               string                                                                                                     `json:"version"`
	SchemaVersion                         string                                                                                                     `json:"schema_version"`
	RequestType                           string                                                                                                     `json:"request_type"`
	AuditType                             string                                                                                                     `json:"audit_type"`
	Source                                string                                                                                                     `json:"source"`
	AuditDecision                         string                                                                                                     `json:"audit_decision"`
	StatusFanOutAuditRequired             bool                                                                                                       `json:"status_fanout_audit_required"`
	StatusFanOutModeled                   bool                                                                                                       `json:"status_fanout_modeled"`
	ConsumerEnablementGateAuditConsumed   bool                                                                                                       `json:"consumer_enablement_gate_audit_consumed"`
	KDESafeStatusGuidanceConsumed         bool                                                                                                       `json:"kde_safe_status_guidance_consumed"`
	CompatibilityCenterStatusModeled      bool                                                                                                       `json:"compatibility_center_status_modeled"`
	RuntimeDiagnosticsStatusModeled       bool                                                                                                       `json:"runtime_diagnostics_status_modeled"`
	StatusFanOutReady                     bool                                                                                                       `json:"status_fanout_ready"`
	ConsumerEnablementGateClosed          bool                                                                                                       `json:"consumer_enablement_gate_closed"`
	ConsumerAuthorizationPrerequisiteSeen bool                                                                                                       `json:"consumer_authorization_prerequisite_seen"`
	RouteAuthorizationPrerequisiteSeen    bool                                                                                                       `json:"route_authorization_prerequisite_seen"`
	ConsumerRedactionPrerequisiteSeen     bool                                                                                                       `json:"consumer_redaction_prerequisite_seen"`
	OpaqueResultIDSupported               bool                                                                                                       `json:"opaque_result_id_supported"`
	CompatibilityCenterStatusItemCount    int                                                                                                        `json:"compatibility_center_status_item_count"`
	RuntimeDiagnosticsStatusItemCount     int                                                                                                        `json:"runtime_diagnostics_status_item_count"`
	StatusItemCount                       int                                                                                                        `json:"status_item_count"`
	RequiredStatusItemCount               int                                                                                                        `json:"required_status_item_count"`
	ReadyStatusItemCount                  int                                                                                                        `json:"ready_status_item_count"`
	MissingStatusItemCount                int                                                                                                        `json:"missing_status_item_count"`
	ConsumerEnabledStatusItemCount        int                                                                                                        `json:"consumer_enabled_status_item_count"`
	PersistedStatusItemCount              int                                                                                                        `json:"persisted_status_item_count"`
	RawExposedStatusItemCount             int                                                                                                        `json:"raw_exposed_status_item_count"`
	SideEffectStatusItemCount             int                                                                                                        `json:"side_effect_status_item_count"`
	StatusItems                           []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusFanOutAuditItem      `json:"status_items"`
	StatusItemIDs                         []string                                                                                                   `json:"status_item_ids"`
	Checks                                []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusFanOutAuditCheck     `json:"checks"`
	CheckIDs                              []string                                                                                                   `json:"check_ids"`
	Counts                                ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusFanOutAuditCheckCounts `json:"counts"`
	ConsumerConsumptionAuthorized         bool                                                                                                       `json:"consumer_consumption_authorized"`
	ConsumerEnablementAuthorized          bool                                                                                                       `json:"consumer_enablement_authorized"`
	KDEConsumerEnabled                    bool                                                                                                       `json:"kde_consumer_enabled"`
	RuntimeConsumerEnabled                bool                                                                                                       `json:"runtime_consumer_enabled"`
	LookupRouteAuthorized                 bool                                                                                                       `json:"lookup_route_authorized"`
	LookupRouteEnabled                    bool                                                                                                       `json:"lookup_route_enabled"`
	LookupRoutePersisted                  bool                                                                                                       `json:"lookup_route_persisted"`
	OpaqueLookupEnabled                   bool                                                                                                       `json:"opaque_lookup_enabled"`
	OpaqueLookupPersisted                 bool                                                                                                       `json:"opaque_lookup_persisted"`
	RedactedSummaryPersisted              bool                                                                                                       `json:"redacted_summary_persisted"`
	KDEStatusPersisted                    bool                                                                                                       `json:"kde_status_persisted"`
	RuntimeDiagnosticsPersisted           bool                                                                                                       `json:"runtime_diagnostics_persisted"`
	DryRunResultPersisted                 bool                                                                                                       `json:"dry_run_result_persisted"`
	RawResultExposed                      bool                                                                                                       `json:"raw_result_exposed"`
	DispatchDryRunExecuted                bool                                                                                                       `json:"dispatch_dry_run_executed"`
	RequestObjectCreationEnabled          bool                                                                                                       `json:"request_object_creation_enabled"`
	RequestObjectDispatchEnabled          bool                                                                                                       `json:"request_object_dispatch_enabled"`
	PortalRequestCreated                  bool                                                                                                       `json:"portal_request_created"`
	NotificationActionEnabled             bool                                                                                                       `json:"notification_action_enabled"`
	CompatibilityCenterOpened             bool                                                                                                       `json:"compatibility_center_opened"`
	SupportBundleExported                 bool                                                                                                       `json:"support_bundle_exported"`
	SupportCaseCreated                    bool                                                                                                       `json:"support_case_created"`
	RuntimeOwned                          bool                                                                                                       `json:"runtime_owned"`
	GoRuntimeBacked                       bool                                                                                                       `json:"go_runtime_backed"`
	KDEPolicyOwner                        bool                                                                                                       `json:"kde_policy_owner"`
	OfficialDesktopOnly                   bool                                                                                                       `json:"official_desktop_only"`
	PlasmaForkRequired                    bool                                                                                                       `json:"plasma_fork_required"`
	PlasmaSourceModified                  bool                                                                                                       `json:"plasma_source_modified"`
	ProductionReadiness                   bool                                                                                                       `json:"production_readiness"`
	ProductionOwnershipReady              bool                                                                                                       `json:"production_ownership_ready"`
	SystemServiceStarted                  bool                                                                                                       `json:"system_service_started"`
	SessionBusClaimed                     bool                                                                                                       `json:"session_bus_claimed"`
	ProductionBusClaimed                  bool                                                                                                       `json:"production_bus_claimed"`
	ProductionOwnerEnabled                bool                                                                                                       `json:"production_owner_enabled"`
	ProductionActivationReady             bool                                                                                                       `json:"production_activation_ready"`
	WriteMethodsEnabled                   bool                                                                                                       `json:"write_methods_enabled"`
	RuntimeWritesEnabled                  bool                                                                                                       `json:"runtime_writes_enabled"`
	RequestObjectsCreated                 bool                                                                                                       `json:"request_objects_created"`
	RequestObjectsDispatched              bool                                                                                                       `json:"request_objects_dispatched"`
	NotificationSent                      bool                                                                                                       `json:"notification_sent"`
	NotificationDeliveryEnabled           bool                                                                                                       `json:"notification_delivery_enabled"`
	DesktopFilesWritten                   bool                                                                                                       `json:"desktop_files_written"`
	MIMEAppsWritten                       bool                                                                                                       `json:"mimeapps_written"`
	ShellConfigurationWritten             bool                                                                                                       `json:"shell_configuration_written"`
	SettingsPersisted                     bool                                                                                                       `json:"settings_persisted"`
	AdapterInvocationEnabled              bool                                                                                                       `json:"adapter_invocation_enabled"`
	BackendLaunchEnabled                  bool                                                                                                       `json:"backend_launch_enabled"`
	BackendProcessStarted                 bool                                                                                                       `json:"backend_process_started"`
	SnapshotRestoreExecuted               bool                                                                                                       `json:"snapshot_restore_executed"`
	StateCleanupExecuted                  bool                                                                                                       `json:"state_cleanup_executed"`
	NetworkRequired                       bool                                                                                                       `json:"network_required"`
	HostRootModified                      bool                                                                                                       `json:"host_root_modified"`
	PrivilegedContainerRequired           bool                                                                                                       `json:"privileged_container_required"`
	CallerStateRootRequired               bool                                                                                                       `json:"caller_state_root_required"`
	StateRootPathExposed                  bool                                                                                                       `json:"state_root_path_exposed"`
	FilePathsExposed                      bool                                                                                                       `json:"file_paths_exposed"`
	FileContentRead                       bool                                                                                                       `json:"file_content_read"`
	RawCommandExposed                     bool                                                                                                       `json:"raw_command_exposed"`
	RawExecutableExposed                  bool                                                                                                       `json:"raw_executable_exposed"`
	BackendDetailsExposed                 bool                                                                                                       `json:"backend_details_exposed"`
	BlockedActions                        []string                                                                                                   `json:"blocked_actions"`
	NextRequirements                      []string                                                                                                   `json:"next_requirements"`
	DesktopSafeSummary                    string                                                                                                     `json:"desktop_safe_summary"`
}

type ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusFanOutAuditItem struct {
	ID                                  string `json:"id"`
	ActionKind                          string `json:"action_kind"`
	SurfaceKind                         string `json:"surface_kind"`
	StatusConsumerKind                  string `json:"status_consumer_kind"`
	OpaqueResultID                      string `json:"opaque_result_id"`
	RequiredEvidence                    string `json:"required_evidence"`
	EvidencePresent                     bool   `json:"evidence_present"`
	KDESafeStatus                       bool   `json:"kde_safe_status"`
	CompatibilityCenterStatus           bool   `json:"compatibility_center_status"`
	RuntimeDiagnosticsStatus            bool   `json:"runtime_diagnostics_status"`
	ConsumerEnablementGateClosed        bool   `json:"consumer_enablement_gate_closed"`
	ConsumerEnablementGateAuditConsumed bool   `json:"consumer_enablement_gate_audit_consumed"`
	OpaqueResultIDSupported             bool   `json:"opaque_result_id_supported"`
	UserVisible                         bool   `json:"user_visible"`
	ReviewOnly                          bool   `json:"review_only"`
	RuntimeOwned                        bool   `json:"runtime_owned"`
	GoRuntimeBacked                     bool   `json:"go_runtime_backed"`
	KDEPolicyOwner                      bool   `json:"kde_policy_owner"`
	ConsumerConsumptionAuthorized       bool   `json:"consumer_consumption_authorized"`
	ConsumerEnablementAuthorized        bool   `json:"consumer_enablement_authorized"`
	KDEConsumerEnabled                  bool   `json:"kde_consumer_enabled"`
	RuntimeConsumerEnabled              bool   `json:"runtime_consumer_enabled"`
	LookupRouteEnabled                  bool   `json:"lookup_route_enabled"`
	OpaqueLookupEnabled                 bool   `json:"opaque_lookup_enabled"`
	RedactedSummaryPersisted            bool   `json:"redacted_summary_persisted"`
	KDEStatusPersisted                  bool   `json:"kde_status_persisted"`
	RuntimeDiagnosticsPersisted         bool   `json:"runtime_diagnostics_persisted"`
	DryRunResultPersisted               bool   `json:"dry_run_result_persisted"`
	RawResultExposed                    bool   `json:"raw_result_exposed"`
	DispatchDryRunExecuted              bool   `json:"dispatch_dry_run_executed"`
	RequestObjectCreated                bool   `json:"request_object_created"`
	RequestObjectDispatched             bool   `json:"request_object_dispatched"`
	PortalRequestCreated                bool   `json:"portal_request_created"`
	NotificationActionEnabled           bool   `json:"notification_action_enabled"`
	CompatibilityCenterOpened           bool   `json:"compatibility_center_opened"`
	SupportBundleExported               bool   `json:"support_bundle_exported"`
	SupportCaseCreated                  bool   `json:"support_case_created"`
	CallerStateRootRequired             bool   `json:"caller_state_root_required"`
	StateRootPathExposed                bool   `json:"state_root_path_exposed"`
	FilePathsExposed                    bool   `json:"file_paths_exposed"`
	FileContentRead                     bool   `json:"file_content_read"`
	SideEffectsDisabled                 bool   `json:"side_effects_disabled"`
	HostRootModified                    bool   `json:"host_root_modified"`
	InternalDetailsExposed              bool   `json:"internal_details_exposed"`
	Status                              string `json:"status"`
	NextRequirement                     string `json:"next_requirement"`
}

type ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusFanOutAuditCheck struct {
	ID      string `json:"id"`
	Status  string `json:"status"`
	Summary string `json:"summary"`
}

type ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusFanOutAuditCheckCounts struct {
	Total   int `json:"total"`
	Passed  int `json:"passed"`
	Pending int `json:"pending"`
	Blocked int `json:"blocked"`
}

func NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusFanOutAuditPreview(root string) (ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusFanOutAuditPreview, error) {
	version, err := readProductionDBusGateReviewVersion(root)
	if err != nil {
		return ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusFanOutAuditPreview{}, err
	}
	sources := productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusFanOutAuditSources(root)
	items := productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusFanOutItems(sources)
	gateReady := productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusFanOutGateReady(sources.ConsumerEnablementGateAudit)
	guidanceReady := productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusFanOutGuidanceReady(sources.DispatchSheet)
	compatibilityCenterReady := productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusFanOutCompatibilityCenterCount(items) == 5 && productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusFanOutReadySurfaceCount(items, "compatibility-center") == 5
	runtimeDiagnosticsReady := productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusFanOutRuntimeDiagnosticsCount(items) == 5 && productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusFanOutReadySurfaceCount(items, "runtime-diagnostics") == 5
	fanOutReady := gateReady && guidanceReady && compatibilityCenterReady && runtimeDiagnosticsReady && productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusFanOutItemsReady(items)
	preview := ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusFanOutAuditPreview{
		Version:                               version,
		SchemaVersion:                         "xnix.runtime.production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_kde_safe_status_fanout_audit.v1",
		RequestType:                           "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-status-fanout-audit-preview",
		AuditType:                             "receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-status-fanout-audit",
		Source:                                "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-receipt-consumer-enablement-gate-audit-preview+claude-code-current-dispatch-picks",
		AuditDecision:                         "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-status-fanout-audit-blocked",
		StatusFanOutAuditRequired:             true,
		StatusFanOutModeled:                   true,
		ConsumerEnablementGateAuditConsumed:   gateReady,
		KDESafeStatusGuidanceConsumed:         guidanceReady,
		CompatibilityCenterStatusModeled:      compatibilityCenterReady,
		RuntimeDiagnosticsStatusModeled:       runtimeDiagnosticsReady,
		StatusFanOutReady:                     fanOutReady,
		ConsumerEnablementGateClosed:          gateReady,
		ConsumerAuthorizationPrerequisiteSeen: gateReady,
		RouteAuthorizationPrerequisiteSeen:    gateReady,
		ConsumerRedactionPrerequisiteSeen:     gateReady,
		OpaqueResultIDSupported:               gateReady,
		CompatibilityCenterStatusItemCount:    productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusFanOutCompatibilityCenterCount(items),
		RuntimeDiagnosticsStatusItemCount:     productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusFanOutRuntimeDiagnosticsCount(items),
		StatusItemCount:                       len(items),
		RequiredStatusItemCount:               10,
		ReadyStatusItemCount:                  productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusFanOutReadyCount(items),
		MissingStatusItemCount:                productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusFanOutMissingCount(items),
		ConsumerEnabledStatusItemCount:        productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusFanOutConsumerEnabledCount(items),
		PersistedStatusItemCount:              productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusFanOutPersistedCount(items),
		RawExposedStatusItemCount:             productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusFanOutRawExposedCount(items),
		SideEffectStatusItemCount:             productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusFanOutSideEffectCount(items),
		StatusItems:                           items,
		StatusItemIDs:                         productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusFanOutItemIDs(items),
		ConsumerConsumptionAuthorized:         false,
		ConsumerEnablementAuthorized:          false,
		KDEConsumerEnabled:                    false,
		RuntimeConsumerEnabled:                false,
		LookupRouteAuthorized:                 false,
		LookupRouteEnabled:                    false,
		LookupRoutePersisted:                  false,
		OpaqueLookupEnabled:                   false,
		OpaqueLookupPersisted:                 false,
		RedactedSummaryPersisted:              false,
		KDEStatusPersisted:                    false,
		RuntimeDiagnosticsPersisted:           false,
		DryRunResultPersisted:                 false,
		RawResultExposed:                      false,
		DispatchDryRunExecuted:                false,
		RequestObjectCreationEnabled:          false,
		RequestObjectDispatchEnabled:          false,
		PortalRequestCreated:                  false,
		NotificationActionEnabled:             false,
		CompatibilityCenterOpened:             false,
		SupportBundleExported:                 false,
		SupportCaseCreated:                    false,
		RuntimeOwned:                          true,
		GoRuntimeBacked:                       true,
		KDEPolicyOwner:                        false,
		OfficialDesktopOnly:                   true,
		PlasmaForkRequired:                    false,
		PlasmaSourceModified:                  false,
		ProductionReadiness:                   false,
		ProductionOwnershipReady:              false,
		SystemServiceStarted:                  false,
		SessionBusClaimed:                     false,
		ProductionBusClaimed:                  false,
		ProductionOwnerEnabled:                false,
		ProductionActivationReady:             false,
		WriteMethodsEnabled:                   false,
		RuntimeWritesEnabled:                  false,
		RequestObjectsCreated:                 false,
		RequestObjectsDispatched:              false,
		NotificationSent:                      false,
		NotificationDeliveryEnabled:           false,
		DesktopFilesWritten:                   false,
		MIMEAppsWritten:                       false,
		ShellConfigurationWritten:             false,
		SettingsPersisted:                     false,
		AdapterInvocationEnabled:              false,
		BackendLaunchEnabled:                  false,
		BackendProcessStarted:                 false,
		SnapshotRestoreExecuted:               false,
		StateCleanupExecuted:                  false,
		NetworkRequired:                       false,
		HostRootModified:                      false,
		PrivilegedContainerRequired:           false,
		CallerStateRootRequired:               false,
		StateRootPathExposed:                  false,
		FilePathsExposed:                      false,
		FileContentRead:                       false,
		RawCommandExposed:                     false,
		RawExecutableExposed:                  false,
		BackendDetailsExposed:                 false,
		BlockedActions: []string{
			"treat KDE-safe status fan-out as permission to enable KDE or Runtime consumers",
			"persist status summaries, persist diagnostics, enable lookup routes, enable opaque lookup, or expose raw dry-run result data",
			"create request objects, dispatch actions, open Compatibility Center, send notifications, write support bundles, start services, launch engines, or mutate host root",
		},
		NextRequirements: []string{
			"Implement redacted summary persistence separately before any durable status storage is allowed.",
			"Keep Compatibility Center and Runtime diagnostics status fan-out read-only until consumer enablement is explicitly authorized.",
			"Keep all status payloads KDE-safe: no raw result data, no state-root paths, no host paths, no raw commands, and no engine details.",
		},
		DesktopSafeSummary: "The production receipt notification action dry-run result lookup consumer enablement KDE-safe status fan-out audit projects the closed consumer enablement gate to Compatibility Center and Runtime diagnostics status summaries without enabling consumers, lookup routes, persistence, execution, production ownership, engine launch, unsafe data exposure, or host mutation.",
	}
	checks := productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusFanOutAuditChecks(preview)
	preview.Checks = checks
	preview.CheckIDs = productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusFanOutCheckIDs(checks)
	preview.Counts = countProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusFanOutChecks(checks)
	if preview.Counts.Blocked == 0 {
		preview.AuditDecision = "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-status-fanout-audit-ready-status-only-consumers-disabled"
	}
	if err := validateNoBackendTerms(preview, "production receipt notification action dry-run result lookup consumer enablement kde safe status fanout audit preview"); err != nil {
		return ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusFanOutAuditPreview{}, err
	}
	return preview, nil
}

type productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusFanOutAuditSourceSet struct {
	ConsumerEnablementGateAudit string
	DispatchSheet               string
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusFanOutAuditSources(root string) productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusFanOutAuditSourceSet {
	return productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusFanOutAuditSourceSet{
		ConsumerEnablementGateAudit: productionAuthorizationReadSources(root, []string{"internal/runtime/owner/production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_receipt_consumer_enablement_gate_audit.go"}),
		DispatchSheet:               productionAuthorizationReadSources(root, []string{"docs/claude-code-current-dispatch-picks.md"}),
	}
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusFanOutItems(sources productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusFanOutAuditSourceSet) []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusFanOutAuditItem {
	combined := sources.ConsumerEnablementGateAudit + sources.DispatchSheet
	return []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusFanOutAuditItem{
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusFanOutItem("review-receipt-dry-run-result-lookup-consumer-enablement-kde-safe-compatibility-center-status-fanout", "review", "compatibility-center", "review-result-compatibility-center-status", "review-dry-run-result-opaque-id", combined, []string{"review-receipt-dry-run-result-lookup-consumer-enablement-receipt-consumer-enablement-gate", "KDE-safe status fan-out audit", "Compatibility Center", "without enabling consumers"}, "model durable redacted status storage before Compatibility Center status can persist"),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusFanOutItem("review-receipt-dry-run-result-lookup-consumer-enablement-kde-safe-runtime-diagnostics-status-fanout", "review", "runtime-diagnostics", "review-result-runtime-diagnostics-status", "review-dry-run-result-opaque-id", combined, []string{"review-receipt-dry-run-result-lookup-consumer-enablement-receipt-consumer-enablement-gate", "KDE-safe status fan-out audit", "Runtime diagnostics", "without enabling consumers"}, "model durable redacted diagnostics storage before Runtime diagnostics status can persist"),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusFanOutItem("renew-receipt-dry-run-result-lookup-consumer-enablement-kde-safe-compatibility-center-status-fanout", "renew", "compatibility-center", "renewal-result-compatibility-center-status", "renewal-dry-run-result-opaque-id", combined, []string{"renew-receipt-dry-run-result-lookup-consumer-enablement-receipt-consumer-enablement-gate", "KDE-safe status fan-out audit", "Compatibility Center", "without enabling consumers"}, "model durable redacted status storage before Compatibility Center status can persist"),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusFanOutItem("renew-receipt-dry-run-result-lookup-consumer-enablement-kde-safe-runtime-diagnostics-status-fanout", "renew", "runtime-diagnostics", "renewal-result-runtime-diagnostics-status", "renewal-dry-run-result-opaque-id", combined, []string{"renew-receipt-dry-run-result-lookup-consumer-enablement-receipt-consumer-enablement-gate", "KDE-safe status fan-out audit", "Runtime diagnostics", "without enabling consumers"}, "model durable redacted diagnostics storage before Runtime diagnostics status can persist"),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusFanOutItem("open-compatibility-center-dry-run-result-lookup-consumer-enablement-kde-safe-compatibility-center-status-fanout", "open-compatibility-center", "compatibility-center", "navigation-result-compatibility-center-status", "navigation-dry-run-result-opaque-id", combined, []string{"open-compatibility-center-dry-run-result-lookup-consumer-enablement-receipt-consumer-enablement-gate", "KDE-safe status fan-out audit", "Compatibility Center", "without enabling consumers"}, "model durable redacted status storage before Compatibility Center status can persist"),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusFanOutItem("open-compatibility-center-dry-run-result-lookup-consumer-enablement-kde-safe-runtime-diagnostics-status-fanout", "open-compatibility-center", "runtime-diagnostics", "navigation-result-runtime-diagnostics-status", "navigation-dry-run-result-opaque-id", combined, []string{"open-compatibility-center-dry-run-result-lookup-consumer-enablement-receipt-consumer-enablement-gate", "KDE-safe status fan-out audit", "Runtime diagnostics", "without enabling consumers"}, "model durable redacted diagnostics storage before Runtime diagnostics status can persist"),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusFanOutItem("dismiss-receipt-dry-run-result-lookup-consumer-enablement-kde-safe-compatibility-center-status-fanout", "dismiss", "compatibility-center", "dismissal-result-compatibility-center-status", "dismissal-dry-run-result-opaque-id", combined, []string{"dismiss-receipt-dry-run-result-lookup-consumer-enablement-receipt-consumer-enablement-gate", "KDE-safe status fan-out audit", "Compatibility Center", "without enabling consumers"}, "model durable redacted status storage before Compatibility Center status can persist"),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusFanOutItem("dismiss-receipt-dry-run-result-lookup-consumer-enablement-kde-safe-runtime-diagnostics-status-fanout", "dismiss", "runtime-diagnostics", "dismissal-result-runtime-diagnostics-status", "dismissal-dry-run-result-opaque-id", combined, []string{"dismiss-receipt-dry-run-result-lookup-consumer-enablement-receipt-consumer-enablement-gate", "KDE-safe status fan-out audit", "Runtime diagnostics", "without enabling consumers"}, "model durable redacted diagnostics storage before Runtime diagnostics status can persist"),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusFanOutItem("support-info-dry-run-result-lookup-consumer-enablement-kde-safe-compatibility-center-status-fanout", "support-info", "compatibility-center", "support-info-result-compatibility-center-status", "support-info-dry-run-result-opaque-id", combined, []string{"support-info-dry-run-result-lookup-consumer-enablement-receipt-consumer-enablement-gate", "KDE-safe status fan-out audit", "Compatibility Center", "without enabling consumers"}, "model durable redacted status storage before Compatibility Center status can persist"),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusFanOutItem("support-info-dry-run-result-lookup-consumer-enablement-kde-safe-runtime-diagnostics-status-fanout", "support-info", "runtime-diagnostics", "support-info-result-runtime-diagnostics-status", "support-info-dry-run-result-opaque-id", combined, []string{"support-info-dry-run-result-lookup-consumer-enablement-receipt-consumer-enablement-gate", "KDE-safe status fan-out audit", "Runtime diagnostics", "without enabling consumers"}, "model durable redacted diagnostics storage before Runtime diagnostics status can persist"),
	}
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusFanOutItem(id string, actionKind string, surfaceKind string, statusConsumerKind string, opaqueResultID string, source string, tokens []string, nextRequirement string) ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusFanOutAuditItem {
	ready := productionAuthorizationHasAll(source, tokens)
	status := "missing-kde-safe-status-fanout-evidence"
	if ready {
		status = "kde-safe-status-fanout-modeled-consumers-disabled"
	}
	return ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusFanOutAuditItem{
		ID:                                  id,
		ActionKind:                          actionKind,
		SurfaceKind:                         surfaceKind,
		StatusConsumerKind:                  statusConsumerKind,
		OpaqueResultID:                      opaqueResultID,
		RequiredEvidence:                    "KDE-safe status fan-out requires a closed consumer enablement gate and explicit dispatch guidance before status-only summaries can be shown",
		EvidencePresent:                     ready,
		KDESafeStatus:                       ready,
		CompatibilityCenterStatus:           ready && surfaceKind == "compatibility-center",
		RuntimeDiagnosticsStatus:            ready && surfaceKind == "runtime-diagnostics",
		ConsumerEnablementGateClosed:        ready,
		ConsumerEnablementGateAuditConsumed: ready,
		OpaqueResultIDSupported:             ready,
		UserVisible:                         ready,
		ReviewOnly:                          true,
		RuntimeOwned:                        true,
		GoRuntimeBacked:                     true,
		KDEPolicyOwner:                      false,
		ConsumerConsumptionAuthorized:       false,
		ConsumerEnablementAuthorized:        false,
		KDEConsumerEnabled:                  false,
		RuntimeConsumerEnabled:              false,
		LookupRouteEnabled:                  false,
		OpaqueLookupEnabled:                 false,
		RedactedSummaryPersisted:            false,
		KDEStatusPersisted:                  false,
		RuntimeDiagnosticsPersisted:         false,
		DryRunResultPersisted:               false,
		RawResultExposed:                    false,
		DispatchDryRunExecuted:              false,
		RequestObjectCreated:                false,
		RequestObjectDispatched:             false,
		PortalRequestCreated:                false,
		NotificationActionEnabled:           false,
		CompatibilityCenterOpened:           false,
		SupportBundleExported:               false,
		SupportCaseCreated:                  false,
		CallerStateRootRequired:             false,
		StateRootPathExposed:                false,
		FilePathsExposed:                    false,
		FileContentRead:                     false,
		SideEffectsDisabled:                 true,
		HostRootModified:                    false,
		InternalDetailsExposed:              false,
		Status:                              status,
		NextRequirement:                     nextRequirement,
	}
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusFanOutAuditChecks(preview ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusFanOutAuditPreview) []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusFanOutAuditCheck {
	return []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusFanOutAuditCheck{
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusFanOutCheck("kde-safe-status-fanout-audit-required", productionAuthorizationPassBlocked(preview.StatusFanOutAuditRequired && preview.StatusFanOutModeled), "The KDE-safe status fan-out audit is present and modeled."),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusFanOutCheck("consumer-enablement-gate-audit-consumed", productionAuthorizationPassBlocked(preview.ConsumerEnablementGateAuditConsumed && preview.ConsumerEnablementGateClosed), "The status fan-out consumes the closed consumer enablement gate."),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusFanOutCheck("kde-safe-status-guidance-consumed", productionAuthorizationPassBlocked(preview.KDESafeStatusGuidanceConsumed), "The status fan-out consumes current dispatch guidance."),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusFanOutCheck("compatibility-center-status-modeled", productionAuthorizationPassBlocked(preview.CompatibilityCenterStatusModeled && preview.CompatibilityCenterStatusItemCount == 5), "Compatibility Center receives status-only fan-out rows."),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusFanOutCheck("runtime-diagnostics-status-modeled", productionAuthorizationPassBlocked(preview.RuntimeDiagnosticsStatusModeled && preview.RuntimeDiagnosticsStatusItemCount == 5), "Runtime diagnostics receives status-only fan-out rows."),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusFanOutCheck("ten-kde-safe-status-fanout-items-present", productionAuthorizationPassBlocked(preview.StatusItemCount == 10 && preview.RequiredStatusItemCount == 10 && preview.ReadyStatusItemCount == 10 && preview.MissingStatusItemCount == 0), "Five notification actions are projected to both Compatibility Center and Runtime diagnostics status summaries."),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusFanOutCheck("status-fanout-modeled-only", productionAuthorizationPassBlocked(preview.StatusFanOutReady && preview.ConsumerAuthorizationPrerequisiteSeen && preview.RouteAuthorizationPrerequisiteSeen && preview.ConsumerRedactionPrerequisiteSeen && preview.OpaqueResultIDSupported && !preview.ConsumerConsumptionAuthorized && !preview.ConsumerEnablementAuthorized && !preview.KDEConsumerEnabled && !preview.RuntimeConsumerEnabled), "The fan-out shows closed status without authorizing or enabling consumers."),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusFanOutCheck("consumer-lookup-and-persistence-disabled", productionAuthorizationPassBlocked(!preview.LookupRouteAuthorized && !preview.LookupRouteEnabled && !preview.LookupRoutePersisted && !preview.OpaqueLookupEnabled && !preview.OpaqueLookupPersisted && !preview.RedactedSummaryPersisted && !preview.KDEStatusPersisted && !preview.RuntimeDiagnosticsPersisted && !preview.DryRunResultPersisted && !preview.RawResultExposed && !preview.DispatchDryRunExecuted && preview.ConsumerEnabledStatusItemCount == 0 && preview.PersistedStatusItemCount == 0 && preview.RawExposedStatusItemCount == 0), "Consumers, lookup, persistence, raw result exposure, and execution remain disabled."),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusFanOutCheck("request-notification-navigation-and-support-disabled", productionAuthorizationPassBlocked(!preview.RequestObjectCreationEnabled && !preview.RequestObjectDispatchEnabled && !preview.RequestObjectsCreated && !preview.RequestObjectsDispatched && !preview.PortalRequestCreated && !preview.NotificationActionEnabled && !preview.NotificationSent && !preview.NotificationDeliveryEnabled && !preview.CompatibilityCenterOpened && !preview.SupportBundleExported && !preview.SupportCaseCreated), "Request creation, dispatch, Portal requests, notifications, navigation, and support writes remain disabled."),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusFanOutCheck("production-and-host-boundary-closed", productionAuthorizationPassBlocked(!preview.SystemServiceStarted && !preview.SessionBusClaimed && !preview.ProductionBusClaimed && !preview.ProductionOwnerEnabled && !preview.ProductionActivationReady && !preview.ProductionReadiness && !preview.ProductionOwnershipReady && !preview.WriteMethodsEnabled && !preview.RuntimeWritesEnabled && !preview.AdapterInvocationEnabled && !preview.BackendLaunchEnabled && !preview.BackendProcessStarted && !preview.NetworkRequired && !preview.HostRootModified && !preview.PrivilegedContainerRequired && !preview.CallerStateRootRequired && !preview.StateRootPathExposed && !preview.FilePathsExposed && !preview.FileContentRead && !preview.RawCommandExposed && !preview.RawExecutableExposed && !preview.BackendDetailsExposed && preview.SideEffectStatusItemCount == 0 && productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusFanOutItemsKeepHostClosed(preview.StatusItems)), "Production ownership, writes, engine launch, unsafe data exposure, network, privilege, internal details, and host mutation remain disabled."),
	}
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusFanOutCheck(id string, status string, summary string) ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusFanOutAuditCheck {
	return ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusFanOutAuditCheck{ID: id, Status: status, Summary: summary}
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusFanOutGateReady(source string) bool {
	return productionAuthorizationHasAll(source, []string{
		"production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-receipt-consumer-enablement-gate-audit-preview",
		"production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-receipt-consumer-enablement-gate-audit-ready-consumers-disabled",
		"ConsumerEnablementGateReady",
		"ConsumerAuthorizationPrerequisiteModeled",
		"RouteAuthorizationPrerequisiteModeled",
		"ConsumerRedactionPrerequisiteModeled",
		"KDEConsumerEnabled",
		"RuntimeConsumerEnabled",
	})
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusFanOutGuidanceReady(source string) bool {
	return productionAuthorizationHasAll(source, []string{"KDE-safe status fan-out audit", "Compatibility Center", "Runtime diagnostics", "without enabling consumers"})
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusFanOutReadyCount(items []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusFanOutAuditItem) int {
	count := 0
	for _, item := range items {
		if item.EvidencePresent {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusFanOutMissingCount(items []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusFanOutAuditItem) int {
	count := 0
	for _, item := range items {
		if !item.EvidencePresent {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusFanOutCompatibilityCenterCount(items []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusFanOutAuditItem) int {
	count := 0
	for _, item := range items {
		if item.SurfaceKind == "compatibility-center" {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusFanOutRuntimeDiagnosticsCount(items []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusFanOutAuditItem) int {
	count := 0
	for _, item := range items {
		if item.SurfaceKind == "runtime-diagnostics" {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusFanOutReadySurfaceCount(items []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusFanOutAuditItem, surfaceKind string) int {
	count := 0
	for _, item := range items {
		if item.SurfaceKind == surfaceKind && item.EvidencePresent {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusFanOutConsumerEnabledCount(items []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusFanOutAuditItem) int {
	count := 0
	for _, item := range items {
		if item.ConsumerConsumptionAuthorized || item.ConsumerEnablementAuthorized || item.KDEConsumerEnabled || item.RuntimeConsumerEnabled || item.LookupRouteEnabled || item.OpaqueLookupEnabled {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusFanOutPersistedCount(items []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusFanOutAuditItem) int {
	count := 0
	for _, item := range items {
		if item.RedactedSummaryPersisted || item.KDEStatusPersisted || item.RuntimeDiagnosticsPersisted || item.DryRunResultPersisted {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusFanOutRawExposedCount(items []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusFanOutAuditItem) int {
	count := 0
	for _, item := range items {
		if item.RawResultExposed || item.StateRootPathExposed || item.FilePathsExposed || item.FileContentRead || item.InternalDetailsExposed {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusFanOutSideEffectCount(items []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusFanOutAuditItem) int {
	count := 0
	for _, item := range items {
		if item.RequestObjectCreated || item.RequestObjectDispatched || item.PortalRequestCreated || item.NotificationActionEnabled || item.CompatibilityCenterOpened || item.SupportBundleExported || item.SupportCaseCreated || item.HostRootModified {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusFanOutItemsReady(items []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusFanOutAuditItem) bool {
	if len(items) != 10 {
		return false
	}
	for _, item := range items {
		if !item.EvidencePresent || !item.KDESafeStatus || !item.ConsumerEnablementGateClosed || item.ConsumerConsumptionAuthorized || item.KDEConsumerEnabled || item.RuntimeConsumerEnabled || item.KDEStatusPersisted || item.RuntimeDiagnosticsPersisted || item.RawResultExposed || !item.SideEffectsDisabled {
			return false
		}
	}
	return true
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusFanOutItemsKeepHostClosed(items []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusFanOutAuditItem) bool {
	for _, item := range items {
		if item.HostRootModified || item.InternalDetailsExposed || item.StateRootPathExposed || item.FilePathsExposed || item.FileContentRead || item.RequestObjectCreated || item.RequestObjectDispatched || item.PortalRequestCreated || item.NotificationActionEnabled || item.CompatibilityCenterOpened || item.SupportBundleExported || item.SupportCaseCreated {
			return false
		}
	}
	return true
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusFanOutItemIDs(items []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusFanOutAuditItem) []string {
	ids := make([]string, 0, len(items))
	for _, item := range items {
		ids = append(ids, item.ID)
	}
	return ids
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusFanOutCheckIDs(checks []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusFanOutAuditCheck) []string {
	ids := make([]string, 0, len(checks))
	for _, check := range checks {
		ids = append(ids, check.ID)
	}
	return ids
}

func countProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusFanOutChecks(checks []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusFanOutAuditCheck) ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusFanOutAuditCheckCounts {
	counts := ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusFanOutAuditCheckCounts{Total: len(checks)}
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
