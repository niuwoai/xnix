package owner

type ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerEnablementGateAuditPreview struct {
	Version                                  string                                                                                                           `json:"version"`
	SchemaVersion                            string                                                                                                           `json:"schema_version"`
	RequestType                              string                                                                                                           `json:"request_type"`
	AuditType                                string                                                                                                           `json:"audit_type"`
	Source                                   string                                                                                                           `json:"source"`
	AuditDecision                            string                                                                                                           `json:"audit_decision"`
	ConsumerEnablementGateRequired           bool                                                                                                             `json:"consumer_enablement_gate_required"`
	ConsumerEnablementGateModeled            bool                                                                                                             `json:"consumer_enablement_gate_modeled"`
	ConsumerAuthorizationAuditConsumed       bool                                                                                                             `json:"consumer_authorization_audit_consumed"`
	LookupRouteAuthorizationAuditConsumed    bool                                                                                                             `json:"lookup_route_authorization_audit_consumed"`
	ConsumerRedactionAuditConsumed           bool                                                                                                             `json:"consumer_redaction_audit_consumed"`
	ConsumerEnablementGuidanceConsumed       bool                                                                                                             `json:"consumer_enablement_guidance_consumed"`
	ConsumerEnablementGateReady              bool                                                                                                             `json:"consumer_enablement_gate_ready"`
	ConsumerAuthorizationPrerequisiteModeled bool                                                                                                             `json:"consumer_authorization_prerequisite_modeled"`
	RouteAuthorizationPrerequisiteModeled    bool                                                                                                             `json:"route_authorization_prerequisite_modeled"`
	ConsumerRedactionPrerequisiteModeled     bool                                                                                                             `json:"consumer_redaction_prerequisite_modeled"`
	OpaqueResultIDSupported                  bool                                                                                                             `json:"opaque_result_id_supported"`
	ConsumerConsumptionAuthorized            bool                                                                                                             `json:"consumer_consumption_authorized"`
	ConsumerEnablementAuthorized             bool                                                                                                             `json:"consumer_enablement_authorized"`
	KDEConsumerEnabled                       bool                                                                                                             `json:"kde_consumer_enabled"`
	RuntimeConsumerEnabled                   bool                                                                                                             `json:"runtime_consumer_enabled"`
	LookupRouteAuthorized                    bool                                                                                                             `json:"lookup_route_authorized"`
	LookupRouteEnabled                       bool                                                                                                             `json:"lookup_route_enabled"`
	LookupRoutePersisted                     bool                                                                                                             `json:"lookup_route_persisted"`
	OpaqueLookupEnabled                      bool                                                                                                             `json:"opaque_lookup_enabled"`
	OpaqueLookupPersisted                    bool                                                                                                             `json:"opaque_lookup_persisted"`
	RedactedSummaryPersisted                 bool                                                                                                             `json:"redacted_summary_persisted"`
	RawResultExposed                         bool                                                                                                             `json:"raw_result_exposed"`
	RuntimeDiagnosticsPersisted              bool                                                                                                             `json:"runtime_diagnostics_persisted"`
	DryRunResultPersisted                    bool                                                                                                             `json:"dry_run_result_persisted"`
	DispatchDryRunExecuted                   bool                                                                                                             `json:"dispatch_dry_run_executed"`
	RequestObjectCreationEnabled             bool                                                                                                             `json:"request_object_creation_enabled"`
	RequestObjectDispatchEnabled             bool                                                                                                             `json:"request_object_dispatch_enabled"`
	PortalRequestCreated                     bool                                                                                                             `json:"portal_request_created"`
	NotificationActionEnabled                bool                                                                                                             `json:"notification_action_enabled"`
	CompatibilityCenterOpened                bool                                                                                                             `json:"compatibility_center_opened"`
	SupportBundleExported                    bool                                                                                                             `json:"support_bundle_exported"`
	SupportCaseCreated                       bool                                                                                                             `json:"support_case_created"`
	ProductionReadiness                      bool                                                                                                             `json:"production_readiness"`
	ProductionOwnershipReady                 bool                                                                                                             `json:"production_ownership_ready"`
	GateItemCount                            int                                                                                                              `json:"gate_item_count"`
	RequiredGateItemCount                    int                                                                                                              `json:"required_gate_item_count"`
	ReadyGateItemCount                       int                                                                                                              `json:"ready_gate_item_count"`
	MissingGateItemCount                     int                                                                                                              `json:"missing_gate_item_count"`
	AuthorizedGateItemCount                  int                                                                                                              `json:"authorized_gate_item_count"`
	EnabledGateItemCount                     int                                                                                                              `json:"enabled_gate_item_count"`
	RedactionReadyGateItemCount              int                                                                                                              `json:"redaction_ready_gate_item_count"`
	RouteReadyGateItemCount                  int                                                                                                              `json:"route_ready_gate_item_count"`
	PersistedGateItemCount                   int                                                                                                              `json:"persisted_gate_item_count"`
	RawExposedGateItemCount                  int                                                                                                              `json:"raw_exposed_gate_item_count"`
	SideEffectGateItemCount                  int                                                                                                              `json:"side_effect_gate_item_count"`
	GateItems                                []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerEnablementGateAuditItem  `json:"gate_items"`
	GateItemIDs                              []string                                                                                                         `json:"gate_item_ids"`
	RequiredBeforeConsumerEnablement         []string                                                                                                         `json:"required_before_consumer_enablement"`
	Checks                                   []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerEnablementGateAuditCheck `json:"checks"`
	CheckIDs                                 []string                                                                                                         `json:"check_ids"`
	Counts                                   ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerEnablementGateAuditCounts  `json:"counts"`
	RuntimeOwned                             bool                                                                                                             `json:"runtime_owned"`
	GoRuntimeBacked                          bool                                                                                                             `json:"go_runtime_backed"`
	KDEPolicyOwner                           bool                                                                                                             `json:"kde_policy_owner"`
	OfficialDesktopOnly                      bool                                                                                                             `json:"official_desktop_only"`
	PlasmaForkRequired                       bool                                                                                                             `json:"plasma_fork_required"`
	PlasmaSourceModified                     bool                                                                                                             `json:"plasma_source_modified"`
	SystemServiceStarted                     bool                                                                                                             `json:"system_service_started"`
	SessionBusClaimed                        bool                                                                                                             `json:"session_bus_claimed"`
	ProductionBusClaimed                     bool                                                                                                             `json:"production_bus_claimed"`
	ProductionOwnerEnabled                   bool                                                                                                             `json:"production_owner_enabled"`
	ProductionActivationReady                bool                                                                                                             `json:"production_activation_ready"`
	WriteMethodsEnabled                      bool                                                                                                             `json:"write_methods_enabled"`
	RuntimeWritesEnabled                     bool                                                                                                             `json:"runtime_writes_enabled"`
	RequestObjectsCreated                    bool                                                                                                             `json:"request_objects_created"`
	RequestObjectsDispatched                 bool                                                                                                             `json:"request_objects_dispatched"`
	NotificationSent                         bool                                                                                                             `json:"notification_sent"`
	NotificationDeliveryEnabled              bool                                                                                                             `json:"notification_delivery_enabled"`
	CompatibilityCenterPersisted             bool                                                                                                             `json:"compatibility_center_persisted"`
	ReceiptWriterEnabled                     bool                                                                                                             `json:"receipt_writer_enabled"`
	ReceiptPersistenceEnabled                bool                                                                                                             `json:"receipt_persistence_enabled"`
	DesktopFilesWritten                      bool                                                                                                             `json:"desktop_files_written"`
	MIMEAppsWritten                          bool                                                                                                             `json:"mimeapps_written"`
	ShellConfigurationWritten                bool                                                                                                             `json:"shell_configuration_written"`
	SettingsPersisted                        bool                                                                                                             `json:"settings_persisted"`
	AdapterInvocationEnabled                 bool                                                                                                             `json:"adapter_invocation_enabled"`
	BackendLaunchEnabled                     bool                                                                                                             `json:"backend_launch_enabled"`
	BackendProcessStarted                    bool                                                                                                             `json:"backend_process_started"`
	SnapshotRestoreExecuted                  bool                                                                                                             `json:"snapshot_restore_executed"`
	StateCleanupExecuted                     bool                                                                                                             `json:"state_cleanup_executed"`
	NetworkRequired                          bool                                                                                                             `json:"network_required"`
	HostRootModified                         bool                                                                                                             `json:"host_root_modified"`
	PrivilegedContainerRequired              bool                                                                                                             `json:"privileged_container_required"`
	CallerStateRootRequired                  bool                                                                                                             `json:"caller_state_root_required"`
	StateRootPathExposed                     bool                                                                                                             `json:"state_root_path_exposed"`
	FilePathsExposed                         bool                                                                                                             `json:"file_paths_exposed"`
	FileContentRead                          bool                                                                                                             `json:"file_content_read"`
	RawCommandExposed                        bool                                                                                                             `json:"raw_command_exposed"`
	RawExecutableExposed                     bool                                                                                                             `json:"raw_executable_exposed"`
	BackendDetailsExposed                    bool                                                                                                             `json:"backend_details_exposed"`
	BlockedActions                           []string                                                                                                         `json:"blocked_actions"`
	NextRequirements                         []string                                                                                                         `json:"next_requirements"`
	DesktopSafeSummary                       string                                                                                                           `json:"desktop_safe_summary"`
}

type ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerEnablementGateAuditItem struct {
	ID                                       string `json:"id"`
	ActionKind                               string `json:"action_kind"`
	SurfaceKind                              string `json:"surface_kind"`
	ConsumerKind                             string `json:"consumer_kind"`
	OpaqueResultID                           string `json:"opaque_result_id"`
	RequiredEvidence                         string `json:"required_evidence"`
	EvidencePresent                          bool   `json:"evidence_present"`
	ConsumerAuthorizationPrerequisiteModeled bool   `json:"consumer_authorization_prerequisite_modeled"`
	RouteAuthorizationPrerequisiteModeled    bool   `json:"route_authorization_prerequisite_modeled"`
	ConsumerRedactionPrerequisiteModeled     bool   `json:"consumer_redaction_prerequisite_modeled"`
	ConsumerEnablementGateModeled            bool   `json:"consumer_enablement_gate_modeled"`
	OpaqueResultIDSupported                  bool   `json:"opaque_result_id_supported"`
	ConsumerConsumptionAuthorized            bool   `json:"consumer_consumption_authorized"`
	ConsumerEnablementAuthorized             bool   `json:"consumer_enablement_authorized"`
	KDEConsumerEnabled                       bool   `json:"kde_consumer_enabled"`
	RuntimeConsumerEnabled                   bool   `json:"runtime_consumer_enabled"`
	LookupRouteAuthorized                    bool   `json:"lookup_route_authorized"`
	LookupRouteEnabled                       bool   `json:"lookup_route_enabled"`
	LookupRoutePersisted                     bool   `json:"lookup_route_persisted"`
	OpaqueLookupEnabled                      bool   `json:"opaque_lookup_enabled"`
	OpaqueLookupPersisted                    bool   `json:"opaque_lookup_persisted"`
	RedactedSummaryPersisted                 bool   `json:"redacted_summary_persisted"`
	RawResultExposed                         bool   `json:"raw_result_exposed"`
	RuntimeDiagnosticsPersisted              bool   `json:"runtime_diagnostics_persisted"`
	DryRunResultPersisted                    bool   `json:"dry_run_result_persisted"`
	DispatchDryRunExecuted                   bool   `json:"dispatch_dry_run_executed"`
	UserVisible                              bool   `json:"user_visible"`
	ReviewOnly                               bool   `json:"review_only"`
	RuntimeOwned                             bool   `json:"runtime_owned"`
	GoRuntimeBacked                          bool   `json:"go_runtime_backed"`
	KDEPolicyOwner                           bool   `json:"kde_policy_owner"`
	CallerStateRootRequired                  bool   `json:"caller_state_root_required"`
	StateRootPathExposed                     bool   `json:"state_root_path_exposed"`
	FilePathsExposed                         bool   `json:"file_paths_exposed"`
	FileContentRead                          bool   `json:"file_content_read"`
	RequestObjectCreated                     bool   `json:"request_object_created"`
	RequestObjectDispatched                  bool   `json:"request_object_dispatched"`
	PortalRequestCreated                     bool   `json:"portal_request_created"`
	NotificationActionEnabled                bool   `json:"notification_action_enabled"`
	CompatibilityCenterOpened                bool   `json:"compatibility_center_opened"`
	SupportBundleExported                    bool   `json:"support_bundle_exported"`
	SupportCaseCreated                       bool   `json:"support_case_created"`
	ProductionReadiness                      bool   `json:"production_readiness"`
	ProductionOwnershipReady                 bool   `json:"production_ownership_ready"`
	SideEffectsDisabled                      bool   `json:"side_effects_disabled"`
	HostRootModified                         bool   `json:"host_root_modified"`
	InternalDetailsExposed                   bool   `json:"internal_details_exposed"`
	GateStatus                               string `json:"gate_status"`
	NextRequirement                          string `json:"next_requirement"`
}

type ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerEnablementGateAuditCheck struct {
	ID      string `json:"id"`
	Status  string `json:"status"`
	Summary string `json:"summary"`
}

type ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerEnablementGateAuditCounts struct {
	Total   int `json:"total"`
	Passed  int `json:"passed"`
	Pending int `json:"pending"`
	Blocked int `json:"blocked"`
}

func NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerEnablementGateAuditPreview(root string) (ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerEnablementGateAuditPreview, error) {
	version, err := readProductionDBusGateReviewVersion(root)
	if err != nil {
		return ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerEnablementGateAuditPreview{}, err
	}
	sources := productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerEnablementGateAuditSources(root)
	items := productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerEnablementGateItems(sources)
	authorizationReady := productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerEnablementGateAuthorizationReady(sources.ConsumerAuthorizationAudit)
	routeReady := productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerEnablementGateRouteReady(sources.LookupRouteAuthorizationAudit)
	redactionReady := productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerEnablementGateRedactionReady(sources.ConsumerRedactionAudit)
	guidanceReady := productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerEnablementGateGuidanceReady(sources.DispatchSheet)
	gateReady := authorizationReady && routeReady && redactionReady && guidanceReady && productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerEnablementGateItemsReady(items)
	preview := ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerEnablementGateAuditPreview{
		Version:                                  version,
		SchemaVersion:                            "xnix.runtime.production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_receipt_consumer_enablement_gate_audit.v1",
		RequestType:                              "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-receipt-consumer-enablement-gate-audit-preview",
		AuditType:                                "receipt-notification-action-dry-run-result-lookup-consumer-enablement-receipt-consumer-enablement-gate-audit",
		Source:                                   "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-receipt-consumer-authorization-audit-preview+production-receipt-notification-action-dry-run-result-lookup-route-authorization-audit-preview+production-receipt-notification-action-dry-run-result-lookup-consumer-redaction-audit-preview+claude-code-current-dispatch-picks",
		AuditDecision:                            "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-receipt-consumer-enablement-gate-audit-blocked",
		ConsumerEnablementGateRequired:           true,
		ConsumerEnablementGateModeled:            true,
		ConsumerAuthorizationAuditConsumed:       authorizationReady,
		LookupRouteAuthorizationAuditConsumed:    routeReady,
		ConsumerRedactionAuditConsumed:           redactionReady,
		ConsumerEnablementGuidanceConsumed:       guidanceReady,
		ConsumerEnablementGateReady:              gateReady,
		ConsumerAuthorizationPrerequisiteModeled: gateReady,
		RouteAuthorizationPrerequisiteModeled:    gateReady,
		ConsumerRedactionPrerequisiteModeled:     gateReady,
		OpaqueResultIDSupported:                  gateReady,
		ConsumerConsumptionAuthorized:            false,
		ConsumerEnablementAuthorized:             false,
		KDEConsumerEnabled:                       false,
		RuntimeConsumerEnabled:                   false,
		LookupRouteAuthorized:                    false,
		LookupRouteEnabled:                       false,
		LookupRoutePersisted:                     false,
		OpaqueLookupEnabled:                      false,
		OpaqueLookupPersisted:                    false,
		RedactedSummaryPersisted:                 false,
		RawResultExposed:                         false,
		RuntimeDiagnosticsPersisted:              false,
		DryRunResultPersisted:                    false,
		DispatchDryRunExecuted:                   false,
		RequestObjectCreationEnabled:             false,
		RequestObjectDispatchEnabled:             false,
		PortalRequestCreated:                     false,
		NotificationActionEnabled:                false,
		CompatibilityCenterOpened:                false,
		SupportBundleExported:                    false,
		SupportCaseCreated:                       false,
		ProductionReadiness:                      false,
		ProductionOwnershipReady:                 false,
		GateItemCount:                            len(items),
		RequiredGateItemCount:                    5,
		ReadyGateItemCount:                       productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerEnablementGateReadyCount(items),
		MissingGateItemCount:                     productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerEnablementGateMissingCount(items),
		RouteReadyGateItemCount:                  productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerEnablementGateRouteReadyCount(items),
		RedactionReadyGateItemCount:              productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerEnablementGateRedactionReadyCount(items),
		GateItems:                                items,
		GateItemIDs:                              productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerEnablementGateItemIDs(items),
		RequiredBeforeConsumerEnablement: []string{
			"modeled receipt consumer authorization audit",
			"accepted lookup route authorization receipt",
			"enforced KDE and Runtime consumer redaction",
			"separate owner-local lookup route implementation",
			"separate consumer enablement authorization receipt",
			"separate lookup and redacted summary persistence implementation",
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
		ReceiptWriterEnabled:         false,
		ReceiptPersistenceEnabled:    false,
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
			"treat this enablement gate audit as permission to enable KDE or Runtime consumers",
			"grant lookup route authorization, enable lookup routes, enable opaque lookup, persist lookup state, or persist redacted summaries",
			"expose raw dry-run result data, state-root paths, host paths, file contents, backend details, raw commands, or raw executables",
			"create request objects, dispatch actions, create Portal requests, write receipts, send notifications, claim production ownership, launch compatibility engines, or mutate host root",
		},
		NextRequirements: []string{
			"Implement the accepted consumer enablement authorization receipt separately.",
			"Implement the owner-local lookup route separately after route authorization is accepted.",
			"Implement redacted summary persistence separately after storage authorization is accepted.",
			"Keep KDE and Runtime consumers disabled until the enablement gate consumes real accepted receipts.",
		},
		DesktopSafeSummary: "The production receipt notification action dry-run result lookup consumer enablement receipt consumer enablement gate audit combines route authorization and consumer redaction evidence before future KDE or Runtime consumers can be enabled, but it grants no authorization, enables no consumer or lookup route, persists no summaries, executes no dry runs, exposes no raw result data, starts no services, launches no engines, and mutates no host state.",
	}
	preview.AuthorizedGateItemCount = productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerEnablementGateAuthorizedCount(items)
	preview.EnabledGateItemCount = productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerEnablementGateEnabledCount(items)
	preview.PersistedGateItemCount = productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerEnablementGatePersistedCount(items)
	preview.RawExposedGateItemCount = productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerEnablementGateRawExposedCount(items)
	preview.SideEffectGateItemCount = productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerEnablementGateSideEffectCount(items)
	checks := productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerEnablementGateAuditChecks(preview)
	preview.Checks = checks
	preview.CheckIDs = productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerEnablementGateCheckIDs(checks)
	preview.Counts = countProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerEnablementGateChecks(checks)
	if preview.Counts.Blocked == 0 {
		preview.AuditDecision = "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-receipt-consumer-enablement-gate-audit-ready-consumers-disabled"
	}
	if err := validateNoBackendTerms(preview, "production receipt notification action dry-run result lookup consumer enablement receipt consumer enablement gate audit preview"); err != nil {
		return ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerEnablementGateAuditPreview{}, err
	}
	return preview, nil
}

type productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerEnablementGateAuditSourceSet struct {
	ConsumerAuthorizationAudit    string
	LookupRouteAuthorizationAudit string
	ConsumerRedactionAudit        string
	DispatchSheet                 string
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerEnablementGateAuditSources(root string) productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerEnablementGateAuditSourceSet {
	return productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerEnablementGateAuditSourceSet{
		ConsumerAuthorizationAudit:    productionAuthorizationReadSources(root, []string{"internal/runtime/owner/production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_receipt_consumer_authorization_audit.go"}),
		LookupRouteAuthorizationAudit: productionAuthorizationReadSources(root, []string{"internal/runtime/owner/production_receipt_notification_action_dry_run_result_lookup_route_authorization_audit.go"}),
		ConsumerRedactionAudit:        productionAuthorizationReadSources(root, []string{"internal/runtime/owner/production_receipt_notification_action_dry_run_result_lookup_consumer_redaction_audit.go"}),
		DispatchSheet:                 productionAuthorizationReadSources(root, []string{"docs/claude-code-current-dispatch-picks.md"}),
	}
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerEnablementGateItems(sources productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerEnablementGateAuditSourceSet) []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerEnablementGateAuditItem {
	combined := sources.ConsumerAuthorizationAudit + sources.LookupRouteAuthorizationAudit + sources.ConsumerRedactionAudit + sources.DispatchSheet
	return []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerEnablementGateAuditItem{
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerEnablementGateItem("review-receipt-dry-run-result-lookup-consumer-enablement-receipt-consumer-enablement-gate", "review", "kde-review-and-runtime-diagnostics", "review-result-consumer-enablement-receipt-consumer-enablement-gate", "review-dry-run-result-opaque-id", combined, []string{"review-receipt-dry-run-result-lookup-consumer-enablement-consumer-authorization", "review-receipt-dry-run-result-lookup-route-authorization", "review-receipt-dry-run-result-lookup-consumer-redaction", "consumer enablement gate"}, "accept a consumer enablement receipt before enabling review result consumers"),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerEnablementGateItem("renew-receipt-dry-run-result-lookup-consumer-enablement-receipt-consumer-enablement-gate", "renew", "kde-renewal-and-runtime-diagnostics", "renewal-result-consumer-enablement-receipt-consumer-enablement-gate", "renewal-dry-run-result-opaque-id", combined, []string{"renew-receipt-dry-run-result-lookup-consumer-enablement-consumer-authorization", "renew-receipt-dry-run-result-lookup-route-authorization", "renew-receipt-dry-run-result-lookup-consumer-redaction", "consumer enablement gate"}, "accept a consumer enablement receipt before enabling renewal result consumers"),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerEnablementGateItem("open-compatibility-center-dry-run-result-lookup-consumer-enablement-receipt-consumer-enablement-gate", "open-compatibility-center", "kde-navigation-and-runtime-diagnostics", "navigation-result-consumer-enablement-receipt-consumer-enablement-gate", "navigation-dry-run-result-opaque-id", combined, []string{"open-compatibility-center-dry-run-result-lookup-consumer-enablement-consumer-authorization", "open-compatibility-center-dry-run-result-lookup-route-authorization", "open-compatibility-center-dry-run-result-lookup-consumer-redaction", "consumer enablement gate"}, "accept a consumer enablement receipt before enabling navigation result consumers"),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerEnablementGateItem("dismiss-receipt-dry-run-result-lookup-consumer-enablement-receipt-consumer-enablement-gate", "dismiss", "kde-dismissal-and-runtime-diagnostics", "dismissal-result-consumer-enablement-receipt-consumer-enablement-gate", "dismissal-dry-run-result-opaque-id", combined, []string{"dismiss-receipt-dry-run-result-lookup-consumer-enablement-consumer-authorization", "dismiss-receipt-dry-run-result-lookup-route-authorization", "dismiss-receipt-dry-run-result-lookup-consumer-redaction", "consumer enablement gate"}, "accept a consumer enablement receipt before enabling dismissal result consumers"),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerEnablementGateItem("support-info-dry-run-result-lookup-consumer-enablement-receipt-consumer-enablement-gate", "support-info", "kde-support-and-runtime-diagnostics", "support-info-result-consumer-enablement-receipt-consumer-enablement-gate", "support-info-dry-run-result-opaque-id", combined, []string{"support-info-dry-run-result-lookup-consumer-enablement-consumer-authorization", "support-info-dry-run-result-lookup-route-authorization", "support-info-dry-run-result-lookup-consumer-redaction", "consumer enablement gate"}, "accept a consumer enablement receipt before enabling support-info result consumers"),
	}
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerEnablementGateItem(id string, actionKind string, surfaceKind string, consumerKind string, opaqueResultID string, source string, tokens []string, nextRequirement string) ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerEnablementGateAuditItem {
	ready := productionAuthorizationHasAll(source, tokens)
	status := "missing-lookup-consumer-enablement-receipt-consumer-enablement-gate-evidence"
	if ready {
		status = "lookup-consumer-enablement-receipt-consumer-enablement-gate-modeled-consumers-disabled"
	}
	return ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerEnablementGateAuditItem{
		ID:                                       id,
		ActionKind:                               actionKind,
		SurfaceKind:                              surfaceKind,
		ConsumerKind:                             consumerKind,
		OpaqueResultID:                           opaqueResultID,
		RequiredEvidence:                         "future dry-run result consumers require consumer authorization, route authorization, redaction, and a separate fail-closed enablement gate",
		EvidencePresent:                          ready,
		ConsumerAuthorizationPrerequisiteModeled: ready,
		RouteAuthorizationPrerequisiteModeled:    ready,
		ConsumerRedactionPrerequisiteModeled:     ready,
		ConsumerEnablementGateModeled:            ready,
		OpaqueResultIDSupported:                  ready,
		ConsumerConsumptionAuthorized:            false,
		ConsumerEnablementAuthorized:             false,
		KDEConsumerEnabled:                       false,
		RuntimeConsumerEnabled:                   false,
		LookupRouteAuthorized:                    false,
		LookupRouteEnabled:                       false,
		LookupRoutePersisted:                     false,
		OpaqueLookupEnabled:                      false,
		OpaqueLookupPersisted:                    false,
		RedactedSummaryPersisted:                 false,
		RawResultExposed:                         false,
		RuntimeDiagnosticsPersisted:              false,
		DryRunResultPersisted:                    false,
		DispatchDryRunExecuted:                   false,
		UserVisible:                              ready,
		ReviewOnly:                               true,
		RuntimeOwned:                             true,
		GoRuntimeBacked:                          true,
		KDEPolicyOwner:                           false,
		CallerStateRootRequired:                  false,
		StateRootPathExposed:                     false,
		FilePathsExposed:                         false,
		FileContentRead:                          false,
		RequestObjectCreated:                     false,
		RequestObjectDispatched:                  false,
		PortalRequestCreated:                     false,
		NotificationActionEnabled:                false,
		CompatibilityCenterOpened:                false,
		SupportBundleExported:                    false,
		SupportCaseCreated:                       false,
		ProductionReadiness:                      false,
		ProductionOwnershipReady:                 false,
		SideEffectsDisabled:                      true,
		HostRootModified:                         false,
		InternalDetailsExposed:                   false,
		GateStatus:                               status,
		NextRequirement:                          nextRequirement,
	}
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerEnablementGateAuditChecks(preview ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerEnablementGateAuditPreview) []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerEnablementGateAuditCheck {
	return []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerEnablementGateAuditCheck{
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerEnablementGateCheck("consumer-enablement-receipt-consumer-enablement-gate-audit-required", productionAuthorizationPassBlocked(preview.ConsumerEnablementGateRequired && preview.ConsumerEnablementGateModeled), "The consumer enablement gate audit is present and modeled."),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerEnablementGateCheck("consumer-authorization-audit-consumed", productionAuthorizationPassBlocked(preview.ConsumerAuthorizationAuditConsumed), "The enablement gate consumes receipt consumer authorization evidence."),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerEnablementGateCheck("lookup-route-authorization-audit-consumed", productionAuthorizationPassBlocked(preview.LookupRouteAuthorizationAuditConsumed), "The enablement gate consumes lookup route authorization evidence."),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerEnablementGateCheck("consumer-redaction-audit-consumed", productionAuthorizationPassBlocked(preview.ConsumerRedactionAuditConsumed), "The enablement gate consumes consumer redaction evidence."),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerEnablementGateCheck("consumer-enablement-guidance-consumed", productionAuthorizationPassBlocked(preview.ConsumerEnablementGuidanceConsumed), "The enablement gate consumes current dispatch guidance."),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerEnablementGateCheck("consumer-enablement-receipt-consumer-enablement-gate-modeled-only", productionAuthorizationPassBlocked(preview.ConsumerEnablementGateReady && preview.ConsumerAuthorizationPrerequisiteModeled && preview.RouteAuthorizationPrerequisiteModeled && preview.ConsumerRedactionPrerequisiteModeled && preview.OpaqueResultIDSupported && !preview.ConsumerConsumptionAuthorized && !preview.ConsumerEnablementAuthorized && !preview.KDEConsumerEnabled && !preview.RuntimeConsumerEnabled), "The gate combines prerequisites without authorizing or enabling consumers."),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerEnablementGateCheck("five-consumer-enablement-receipt-consumer-enablement-gate-items-present", productionAuthorizationPassBlocked(preview.GateItemCount == 5 && preview.RequiredGateItemCount == 5 && preview.MissingGateItemCount == 0 && preview.RouteReadyGateItemCount == 5 && preview.RedactionReadyGateItemCount == 5), "Review, renew, open Compatibility Center, dismiss, and support-info enablement gates are modeled."),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerEnablementGateCheck("lookup-result-persistence-and-execution-disabled", productionAuthorizationPassBlocked(!preview.LookupRouteAuthorized && !preview.LookupRouteEnabled && !preview.LookupRoutePersisted && !preview.OpaqueLookupEnabled && !preview.OpaqueLookupPersisted && !preview.RedactedSummaryPersisted && !preview.RuntimeDiagnosticsPersisted && !preview.DryRunResultPersisted && !preview.DispatchDryRunExecuted && preview.AuthorizedGateItemCount == 0 && preview.EnabledGateItemCount == 0 && preview.PersistedGateItemCount == 0), "Route authorization, lookup, redacted summaries, diagnostics, result persistence, and dry-run execution remain disabled."),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerEnablementGateCheck("request-notification-navigation-and-support-disabled", productionAuthorizationPassBlocked(!preview.RequestObjectCreationEnabled && !preview.RequestObjectDispatchEnabled && !preview.RequestObjectsCreated && !preview.RequestObjectsDispatched && !preview.PortalRequestCreated && !preview.NotificationActionEnabled && !preview.NotificationSent && !preview.CompatibilityCenterOpened && !preview.SupportBundleExported && !preview.SupportCaseCreated), "Request creation, dispatch, Portal requests, notifications, navigation, and support writes remain disabled."),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerEnablementGateCheck("production-and-host-boundary-closed", productionAuthorizationPassBlocked(!preview.SystemServiceStarted && !preview.SessionBusClaimed && !preview.ProductionBusClaimed && !preview.ProductionOwnerEnabled && !preview.ProductionActivationReady && !preview.ProductionReadiness && !preview.ProductionOwnershipReady && !preview.WriteMethodsEnabled && !preview.RuntimeWritesEnabled && !preview.AdapterInvocationEnabled && !preview.BackendLaunchEnabled && !preview.BackendProcessStarted && !preview.NetworkRequired && !preview.HostRootModified && !preview.PrivilegedContainerRequired && !preview.CallerStateRootRequired && !preview.StateRootPathExposed && !preview.FilePathsExposed && !preview.FileContentRead && !preview.RawResultExposed && !preview.RawCommandExposed && !preview.RawExecutableExposed && !preview.BackendDetailsExposed && preview.RawExposedGateItemCount == 0 && preview.SideEffectGateItemCount == 0 && productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerEnablementGateItemsKeepHostClosed(preview.GateItems)), "Production ownership, writes, engine launch, unsafe data exposure, network, privilege, internal details, and host mutation remain disabled."),
	}
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerEnablementGateCheck(id string, status string, summary string) ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerEnablementGateAuditCheck {
	return ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerEnablementGateAuditCheck{ID: id, Status: status, Summary: summary}
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerEnablementGateAuthorizationReady(source string) bool {
	return productionAuthorizationHasAll(source, []string{
		"production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-receipt-consumer-authorization-audit-preview",
		"production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-receipt-consumer-authorization-audit-ready-authorization-disabled",
		"ConsumerAuthorizationReady",
		"ConsumerAuthorizationGranted",
		"KDEConsumerAuthorized",
		"RuntimeConsumerAuthorized",
	})
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerEnablementGateRouteReady(source string) bool {
	return productionAuthorizationHasAll(source, []string{
		"production-receipt-notification-action-dry-run-result-lookup-route-authorization-audit-preview",
		"production-receipt-notification-action-dry-run-result-lookup-route-authorization-audit-ready-route-disabled",
		"review-receipt-dry-run-result-lookup-route-authorization",
		"renew-receipt-dry-run-result-lookup-route-authorization",
		"open-compatibility-center-dry-run-result-lookup-route-authorization",
		"dismiss-receipt-dry-run-result-lookup-route-authorization",
		"support-info-dry-run-result-lookup-route-authorization",
		"LookupRouteAuthorized",
	})
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerEnablementGateRedactionReady(source string) bool {
	return productionAuthorizationHasAll(source, []string{
		"production-receipt-notification-action-dry-run-result-lookup-consumer-redaction-audit-preview",
		"production-receipt-notification-action-dry-run-result-lookup-consumer-redaction-audit-ready-consumers-disabled",
		"review-receipt-dry-run-result-lookup-consumer-redaction",
		"renew-receipt-dry-run-result-lookup-consumer-redaction",
		"open-compatibility-center-dry-run-result-lookup-consumer-redaction",
		"dismiss-receipt-dry-run-result-lookup-consumer-redaction",
		"support-info-dry-run-result-lookup-consumer-redaction",
		"ConsumerConsumptionAuthorized",
	})
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerEnablementGateGuidanceReady(source string) bool {
	return productionAuthorizationHasAll(source, []string{"consumer enablement gate audit", "receipt consumer authorization audit", "KDE or Runtime consumers can be enabled", "side effects"})
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerEnablementGateReadyCount(items []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerEnablementGateAuditItem) int {
	count := 0
	for _, item := range items {
		if item.EvidencePresent {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerEnablementGateMissingCount(items []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerEnablementGateAuditItem) int {
	count := 0
	for _, item := range items {
		if !item.EvidencePresent {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerEnablementGateRouteReadyCount(items []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerEnablementGateAuditItem) int {
	count := 0
	for _, item := range items {
		if item.RouteAuthorizationPrerequisiteModeled {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerEnablementGateRedactionReadyCount(items []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerEnablementGateAuditItem) int {
	count := 0
	for _, item := range items {
		if item.ConsumerRedactionPrerequisiteModeled {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerEnablementGateAuthorizedCount(items []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerEnablementGateAuditItem) int {
	count := 0
	for _, item := range items {
		if item.ConsumerConsumptionAuthorized || item.ConsumerEnablementAuthorized || item.LookupRouteAuthorized {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerEnablementGateEnabledCount(items []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerEnablementGateAuditItem) int {
	count := 0
	for _, item := range items {
		if item.KDEConsumerEnabled || item.RuntimeConsumerEnabled || item.LookupRouteEnabled || item.OpaqueLookupEnabled {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerEnablementGatePersistedCount(items []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerEnablementGateAuditItem) int {
	count := 0
	for _, item := range items {
		if item.LookupRoutePersisted || item.OpaqueLookupPersisted || item.RedactedSummaryPersisted || item.RuntimeDiagnosticsPersisted || item.DryRunResultPersisted {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerEnablementGateRawExposedCount(items []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerEnablementGateAuditItem) int {
	count := 0
	for _, item := range items {
		if item.RawResultExposed || item.CallerStateRootRequired || item.StateRootPathExposed || item.FilePathsExposed || item.FileContentRead || item.InternalDetailsExposed {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerEnablementGateSideEffectCount(items []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerEnablementGateAuditItem) int {
	count := 0
	for _, item := range items {
		if item.ConsumerConsumptionAuthorized || item.ConsumerEnablementAuthorized || item.KDEConsumerEnabled || item.RuntimeConsumerEnabled || item.LookupRouteAuthorized || item.LookupRouteEnabled || item.LookupRoutePersisted || item.OpaqueLookupEnabled || item.OpaqueLookupPersisted || item.RedactedSummaryPersisted || item.RawResultExposed || item.RuntimeDiagnosticsPersisted || item.DryRunResultPersisted || item.DispatchDryRunExecuted || item.RequestObjectCreated || item.RequestObjectDispatched || item.PortalRequestCreated || item.NotificationActionEnabled || item.CompatibilityCenterOpened || item.SupportBundleExported || item.SupportCaseCreated || item.HostRootModified || item.InternalDetailsExposed {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerEnablementGateItemIDs(items []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerEnablementGateAuditItem) []string {
	ids := make([]string, 0, len(items))
	for _, item := range items {
		ids = append(ids, item.ID)
	}
	return ids
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerEnablementGateCheckIDs(checks []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerEnablementGateAuditCheck) []string {
	ids := make([]string, 0, len(checks))
	for _, check := range checks {
		ids = append(ids, check.ID)
	}
	return ids
}

func countProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerEnablementGateChecks(checks []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerEnablementGateAuditCheck) ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerEnablementGateAuditCounts {
	counts := ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerEnablementGateAuditCounts{Total: len(checks)}
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

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerEnablementGateItemsReady(items []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerEnablementGateAuditItem) bool {
	if len(items) != 5 {
		return false
	}
	for _, item := range items {
		if !item.EvidencePresent || !item.ConsumerAuthorizationPrerequisiteModeled || !item.RouteAuthorizationPrerequisiteModeled || !item.ConsumerRedactionPrerequisiteModeled || !item.ConsumerEnablementGateModeled || !item.OpaqueResultIDSupported || item.GateStatus != "lookup-consumer-enablement-receipt-consumer-enablement-gate-modeled-consumers-disabled" {
			return false
		}
		if item.ConsumerConsumptionAuthorized || item.ConsumerEnablementAuthorized || item.KDEConsumerEnabled || item.RuntimeConsumerEnabled || item.LookupRouteAuthorized || item.LookupRouteEnabled || item.LookupRoutePersisted || item.OpaqueLookupEnabled || item.OpaqueLookupPersisted || item.RedactedSummaryPersisted || item.RawResultExposed || item.RuntimeDiagnosticsPersisted || item.DryRunResultPersisted || item.DispatchDryRunExecuted || item.CallerStateRootRequired || item.StateRootPathExposed || item.FilePathsExposed || item.FileContentRead || item.RequestObjectCreated || item.RequestObjectDispatched || item.PortalRequestCreated || item.NotificationActionEnabled || item.CompatibilityCenterOpened || item.SupportBundleExported || item.SupportCaseCreated || item.ProductionReadiness || item.ProductionOwnershipReady || item.KDEPolicyOwner || !item.ReviewOnly || !item.SideEffectsDisabled {
			return false
		}
	}
	return true
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerEnablementGateItemsKeepHostClosed(items []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerEnablementGateAuditItem) bool {
	for _, item := range items {
		if item.CallerStateRootRequired || item.StateRootPathExposed || item.FilePathsExposed || item.FileContentRead || item.HostRootModified || item.InternalDetailsExposed {
			return false
		}
	}
	return true
}
