package owner

type ProductionReceiptNotificationActionDryRunResultLookupRouteAuthorizationAuditPreview struct {
	Version                               string                                                                              `json:"version"`
	SchemaVersion                         string                                                                              `json:"schema_version"`
	RequestType                           string                                                                              `json:"request_type"`
	AuditType                             string                                                                              `json:"audit_type"`
	Source                                string                                                                              `json:"source"`
	AuditDecision                         string                                                                              `json:"audit_decision"`
	ReceiptSchema                         string                                                                              `json:"receipt_schema"`
	OpaqueReceiptID                       string                                                                              `json:"opaque_receipt_id"`
	LookupRouteAuthorizationAuditRequired bool                                                                                `json:"lookup_route_authorization_audit_required"`
	LookupRouteAuthorizationAuditModeled  bool                                                                                `json:"lookup_route_authorization_audit_modeled"`
	OpaqueLookupAuditConsumed             bool                                                                                `json:"opaque_lookup_audit_consumed"`
	LookupRouteGuidanceConsumed           bool                                                                                `json:"lookup_route_guidance_consumed"`
	LookupRouteAuthorizationReady         bool                                                                                `json:"lookup_route_authorization_ready"`
	OwnerLocalLookupRouteModeled          bool                                                                                `json:"owner_local_lookup_route_modeled"`
	KDEConsumptionAuthorized              bool                                                                                `json:"kde_consumption_authorized"`
	RuntimeConsumptionAuthorized          bool                                                                                `json:"runtime_consumption_authorized"`
	LookupRouteAuthorized                 bool                                                                                `json:"lookup_route_authorized"`
	LookupRouteEnabled                    bool                                                                                `json:"lookup_route_enabled"`
	LookupRoutePersisted                  bool                                                                                `json:"lookup_route_persisted"`
	OpaqueLookupEnabled                   bool                                                                                `json:"opaque_lookup_enabled"`
	OpaqueLookupPersisted                 bool                                                                                `json:"opaque_lookup_persisted"`
	OpaqueResultIDSupported               bool                                                                                `json:"opaque_result_id_supported"`
	CallerStateRootRequired               bool                                                                                `json:"caller_state_root_required"`
	StateRootPathExposed                  bool                                                                                `json:"state_root_path_exposed"`
	FilePathsExposed                      bool                                                                                `json:"file_paths_exposed"`
	FileContentRead                       bool                                                                                `json:"file_content_read"`
	ResultPersistenceAuthorized           bool                                                                                `json:"result_persistence_authorized"`
	RetentionEnforcementEnabled           bool                                                                                `json:"retention_enforcement_enabled"`
	RedactionEnforcementEnabled           bool                                                                                `json:"redaction_enforcement_enabled"`
	DispatchDryRunExecuted                bool                                                                                `json:"dispatch_dry_run_executed"`
	DryRunResultPersisted                 bool                                                                                `json:"dry_run_result_persisted"`
	ResultVisibilityPersisted             bool                                                                                `json:"result_visibility_persisted"`
	RuntimeDiagnosticsPersisted           bool                                                                                `json:"runtime_diagnostics_persisted"`
	DispatchAuthorizationGranted          bool                                                                                `json:"dispatch_authorization_granted"`
	OperatorRouteApprovalRequired         bool                                                                                `json:"operator_route_approval_required"`
	OperatorRouteApprovalPresent          bool                                                                                `json:"operator_route_approval_present"`
	ReceiptPresent                        bool                                                                                `json:"receipt_present"`
	ReceiptAccepted                       bool                                                                                `json:"receipt_accepted"`
	AuthorizationAccepted                 bool                                                                                `json:"authorization_accepted"`
	ProductionReadiness                   bool                                                                                `json:"production_readiness"`
	ProductionOwnershipReady              bool                                                                                `json:"production_ownership_ready"`
	RouteItemCount                        int                                                                                 `json:"route_item_count"`
	RequiredRouteItemCount                int                                                                                 `json:"required_route_item_count"`
	ReadyRouteItemCount                   int                                                                                 `json:"ready_route_item_count"`
	MissingRouteItemCount                 int                                                                                 `json:"missing_route_item_count"`
	AuthorizedRouteItemCount              int                                                                                 `json:"authorized_route_item_count"`
	EnabledRouteItemCount                 int                                                                                 `json:"enabled_route_item_count"`
	PersistedRouteItemCount               int                                                                                 `json:"persisted_route_item_count"`
	LookupEnabledItemCount                int                                                                                 `json:"lookup_enabled_item_count"`
	PersistedLookupItemCount              int                                                                                 `json:"persisted_lookup_item_count"`
	PathExposedRouteItemCount             int                                                                                 `json:"path_exposed_route_item_count"`
	PersistedResultItemCount              int                                                                                 `json:"persisted_result_item_count"`
	ExecutedDryRunItemCount               int                                                                                 `json:"executed_dry_run_item_count"`
	CreatedRequestObjectCount             int                                                                                 `json:"created_request_object_count"`
	PortalRequestCreatedCount             int                                                                                 `json:"portal_request_created_count"`
	SideEffectRouteItemCount              int                                                                                 `json:"side_effect_route_item_count"`
	RouteItems                            []ProductionReceiptNotificationActionDryRunResultLookupRouteAuthorizationAuditItem  `json:"route_items"`
	RouteItemIDs                          []string                                                                            `json:"route_item_ids"`
	RequiredBeforeRouteEnablement         []string                                                                            `json:"required_before_route_enablement"`
	Checks                                []ProductionReceiptNotificationActionDryRunResultLookupRouteAuthorizationAuditCheck `json:"checks"`
	CheckIDs                              []string                                                                            `json:"check_ids"`
	Counts                                ProductionReceiptNotificationActionDryRunResultLookupRouteAuthorizationAuditCounts  `json:"counts"`
	RuntimeOwned                          bool                                                                                `json:"runtime_owned"`
	GoRuntimeBacked                       bool                                                                                `json:"go_runtime_backed"`
	KDEPolicyOwner                        bool                                                                                `json:"kde_policy_owner"`
	OfficialDesktopOnly                   bool                                                                                `json:"official_desktop_only"`
	PlasmaForkRequired                    bool                                                                                `json:"plasma_fork_required"`
	PlasmaSourceModified                  bool                                                                                `json:"plasma_source_modified"`
	SystemServiceStarted                  bool                                                                                `json:"system_service_started"`
	SessionBusClaimed                     bool                                                                                `json:"session_bus_claimed"`
	ProductionBusClaimed                  bool                                                                                `json:"production_bus_claimed"`
	ProductionOwnerEnabled                bool                                                                                `json:"production_owner_enabled"`
	ProductionActivationReady             bool                                                                                `json:"production_activation_ready"`
	WriteMethodsEnabled                   bool                                                                                `json:"write_methods_enabled"`
	RuntimeWritesEnabled                  bool                                                                                `json:"runtime_writes_enabled"`
	RequestObjectCreationEnabled          bool                                                                                `json:"request_object_creation_enabled"`
	RequestObjectDispatchEnabled          bool                                                                                `json:"request_object_dispatch_enabled"`
	RequestObjectPersistenceEnabled       bool                                                                                `json:"request_object_persistence_enabled"`
	DispatchAuthorizationPersisted        bool                                                                                `json:"dispatch_authorization_persisted"`
	DispatchDryRunExecutionEnabled        bool                                                                                `json:"dispatch_dry_run_execution_enabled"`
	DryRunResultPersistenceEnabled        bool                                                                                `json:"dry_run_result_persistence_enabled"`
	ResultVisibilityPersistenceEnabled    bool                                                                                `json:"result_visibility_persistence_enabled"`
	RetentionPolicyPersistenceEnabled     bool                                                                                `json:"retention_policy_persistence_enabled"`
	PortalRequestCreated                  bool                                                                                `json:"portal_request_created"`
	RequestObjectsCreated                 bool                                                                                `json:"request_objects_created"`
	RequestObjectsDispatched              bool                                                                                `json:"request_objects_dispatched"`
	NotificationSent                      bool                                                                                `json:"notification_sent"`
	NotificationDeliveryEnabled           bool                                                                                `json:"notification_delivery_enabled"`
	NotificationActionEnabled             bool                                                                                `json:"notification_action_enabled"`
	ReviewActionEnabled                   bool                                                                                `json:"review_action_enabled"`
	RenewActionEnabled                    bool                                                                                `json:"renew_action_enabled"`
	OpenCompatibilityCenterEnabled        bool                                                                                `json:"open_compatibility_center_enabled"`
	DismissActionEnabled                  bool                                                                                `json:"dismiss_action_enabled"`
	SupportInfoActionEnabled              bool                                                                                `json:"support_info_action_enabled"`
	CompatibilityCenterOpened             bool                                                                                `json:"compatibility_center_opened"`
	CompatibilityCenterPersisted          bool                                                                                `json:"compatibility_center_persisted"`
	ReceiptWriterEnabled                  bool                                                                                `json:"receipt_writer_enabled"`
	ReceiptPersistenceEnabled             bool                                                                                `json:"receipt_persistence_enabled"`
	ReceiptLookupWritesEnabled            bool                                                                                `json:"receipt_lookup_writes_enabled"`
	ReceiptReplayEnabled                  bool                                                                                `json:"receipt_replay_enabled"`
	ReceiptExpiryWriteEnabled             bool                                                                                `json:"receipt_expiry_write_enabled"`
	ReceiptRevocationWriteEnabled         bool                                                                                `json:"receipt_revocation_write_enabled"`
	DesktopFilesWritten                   bool                                                                                `json:"desktop_files_written"`
	MIMEAppsWritten                       bool                                                                                `json:"mimeapps_written"`
	ShellConfigurationWritten             bool                                                                                `json:"shell_configuration_written"`
	SettingsPersisted                     bool                                                                                `json:"settings_persisted"`
	KRunnerIndexPersisted                 bool                                                                                `json:"krunner_index_persisted"`
	TaskManagerEntryActive                bool                                                                                `json:"task_manager_entry_active"`
	KWinRuleApplied                       bool                                                                                `json:"kwin_rule_applied"`
	LiveTrayBridgeEnabled                 bool                                                                                `json:"live_tray_bridge_enabled"`
	TrayBridgePersisted                   bool                                                                                `json:"tray_bridge_persisted"`
	AdapterInvocationEnabled              bool                                                                                `json:"adapter_invocation_enabled"`
	BackendLaunchEnabled                  bool                                                                                `json:"backend_launch_enabled"`
	BackendProcessStarted                 bool                                                                                `json:"backend_process_started"`
	SupportBundleExported                 bool                                                                                `json:"support_bundle_exported"`
	SupportCaseCreated                    bool                                                                                `json:"support_case_created"`
	SnapshotRestoreExecuted               bool                                                                                `json:"snapshot_restore_executed"`
	StateCleanupExecuted                  bool                                                                                `json:"state_cleanup_executed"`
	NetworkRequired                       bool                                                                                `json:"network_required"`
	HostRootModified                      bool                                                                                `json:"host_root_modified"`
	PrivilegedContainerRequired           bool                                                                                `json:"privileged_container_required"`
	RawCommandExposed                     bool                                                                                `json:"raw_command_exposed"`
	RawExecutableExposed                  bool                                                                                `json:"raw_executable_exposed"`
	BackendDetailsExposed                 bool                                                                                `json:"backend_details_exposed"`
	BlockedActions                        []string                                                                            `json:"blocked_actions"`
	NextRequirements                      []string                                                                            `json:"next_requirements"`
	DesktopSafeSummary                    string                                                                              `json:"desktop_safe_summary"`
}

type ProductionReceiptNotificationActionDryRunResultLookupRouteAuthorizationAuditItem struct {
	ID                           string `json:"id"`
	ActionKind                   string `json:"action_kind"`
	RouteKind                    string `json:"route_kind"`
	SurfaceKind                  string `json:"surface_kind"`
	OpaqueLookupKind             string `json:"opaque_lookup_kind"`
	OpaqueResultID               string `json:"opaque_result_id"`
	RequiredEvidence             string `json:"required_evidence"`
	EvidencePresent              bool   `json:"evidence_present"`
	RouteAuthorizationModeled    bool   `json:"route_authorization_modeled"`
	OwnerLocalRouteModeled       bool   `json:"owner_local_route_modeled"`
	OpaqueLookupConsumed         bool   `json:"opaque_lookup_consumed"`
	OpaqueResultIDSupported      bool   `json:"opaque_result_id_supported"`
	KDEConsumptionAuthorized     bool   `json:"kde_consumption_authorized"`
	RuntimeConsumptionAuthorized bool   `json:"runtime_consumption_authorized"`
	LookupRouteAuthorized        bool   `json:"lookup_route_authorized"`
	LookupRouteEnabled           bool   `json:"lookup_route_enabled"`
	LookupRoutePersisted         bool   `json:"lookup_route_persisted"`
	LookupEnabled                bool   `json:"lookup_enabled"`
	LookupPersisted              bool   `json:"lookup_persisted"`
	CallerStateRootRequired      bool   `json:"caller_state_root_required"`
	StateRootPathExposed         bool   `json:"state_root_path_exposed"`
	FilePathsExposed             bool   `json:"file_paths_exposed"`
	FileContentRead              bool   `json:"file_content_read"`
	UserVisible                  bool   `json:"user_visible"`
	ReviewOnly                   bool   `json:"review_only"`
	RuntimeOwned                 bool   `json:"runtime_owned"`
	GoRuntimeBacked              bool   `json:"go_runtime_backed"`
	KDEPolicyOwner               bool   `json:"kde_policy_owner"`
	OperatorApprovalRequired     bool   `json:"operator_approval_required"`
	OperatorApprovalPresent      bool   `json:"operator_approval_present"`
	ResultPersistenceAuthorized  bool   `json:"result_persistence_authorized"`
	RetentionEnforcementEnabled  bool   `json:"retention_enforcement_enabled"`
	RedactionEnforcementEnabled  bool   `json:"redaction_enforcement_enabled"`
	DispatchDryRunExecuted       bool   `json:"dispatch_dry_run_executed"`
	DryRunResultPersisted        bool   `json:"dry_run_result_persisted"`
	ResultVisibilityPersisted    bool   `json:"result_visibility_persisted"`
	RuntimeDiagnosticsPersisted  bool   `json:"runtime_diagnostics_persisted"`
	RequestObjectCreated         bool   `json:"request_object_created"`
	RequestObjectDispatched      bool   `json:"request_object_dispatched"`
	PortalRequestCreated         bool   `json:"portal_request_created"`
	ReceiptAccepted              bool   `json:"receipt_accepted"`
	AuthorizationAccepted        bool   `json:"authorization_accepted"`
	CompatibilityCenterOpened    bool   `json:"compatibility_center_opened"`
	SupportBundleExported        bool   `json:"support_bundle_exported"`
	SupportCaseCreated           bool   `json:"support_case_created"`
	ProductionReadiness          bool   `json:"production_readiness"`
	ProductionOwnershipReady     bool   `json:"production_ownership_ready"`
	SideEffectsDisabled          bool   `json:"side_effects_disabled"`
	HostRootModified             bool   `json:"host_root_modified"`
	InternalDetailsExposed       bool   `json:"internal_details_exposed"`
	RouteStatus                  string `json:"route_status"`
	NextRequirement              string `json:"next_requirement"`
}

type ProductionReceiptNotificationActionDryRunResultLookupRouteAuthorizationAuditCheck struct {
	ID      string `json:"id"`
	Status  string `json:"status"`
	Summary string `json:"summary"`
}

type ProductionReceiptNotificationActionDryRunResultLookupRouteAuthorizationAuditCounts struct {
	Total   int `json:"total"`
	Passed  int `json:"passed"`
	Pending int `json:"pending"`
	Blocked int `json:"blocked"`
}

func NewProductionReceiptNotificationActionDryRunResultLookupRouteAuthorizationAuditPreview(root string) (ProductionReceiptNotificationActionDryRunResultLookupRouteAuthorizationAuditPreview, error) {
	version, err := readProductionDBusGateReviewVersion(root)
	if err != nil {
		return ProductionReceiptNotificationActionDryRunResultLookupRouteAuthorizationAuditPreview{}, err
	}
	sources := productionReceiptNotificationActionDryRunResultLookupRouteAuthorizationAuditSources(root)
	items := productionReceiptNotificationActionDryRunResultLookupRouteAuthorizationItems(sources)
	opaqueLookupReady := productionReceiptNotificationActionDryRunResultLookupRouteAuthorizationOpaqueLookupReady(sources.OpaqueLookupAudit)
	guidanceReady := productionReceiptNotificationActionDryRunResultLookupRouteAuthorizationGuidanceReady(sources.DispatchSheet)
	routeReady := opaqueLookupReady && guidanceReady && productionReceiptNotificationActionDryRunResultLookupRouteAuthorizationItemsReady(items)
	preview := ProductionReceiptNotificationActionDryRunResultLookupRouteAuthorizationAuditPreview{
		Version:                               version,
		SchemaVersion:                         "xnix.runtime.production_receipt_notification_action_dry_run_result_lookup_route_authorization_audit.v1",
		RequestType:                           "production-receipt-notification-action-dry-run-result-lookup-route-authorization-audit-preview",
		AuditType:                             "receipt-notification-action-dry-run-result-lookup-route-authorization-audit",
		Source:                                "production-receipt-notification-action-dry-run-result-opaque-lookup-audit-preview+claude-code-current-dispatch-picks",
		AuditDecision:                         "production-receipt-notification-action-dry-run-result-lookup-route-authorization-audit-blocked",
		ReceiptSchema:                         "xnix.runtime.production_dbus_human_authorization_receipt.v1",
		OpaqueReceiptID:                       ProductionDBusHumanAuthorizationReceiptID,
		LookupRouteAuthorizationAuditRequired: true,
		LookupRouteAuthorizationAuditModeled:  true,
		OpaqueLookupAuditConsumed:             opaqueLookupReady,
		LookupRouteGuidanceConsumed:           guidanceReady,
		LookupRouteAuthorizationReady:         routeReady,
		OwnerLocalLookupRouteModeled:          routeReady,
		KDEConsumptionAuthorized:              false,
		RuntimeConsumptionAuthorized:          false,
		LookupRouteAuthorized:                 false,
		LookupRouteEnabled:                    false,
		LookupRoutePersisted:                  false,
		OpaqueLookupEnabled:                   false,
		OpaqueLookupPersisted:                 false,
		OpaqueResultIDSupported:               routeReady,
		CallerStateRootRequired:               false,
		StateRootPathExposed:                  false,
		FilePathsExposed:                      false,
		FileContentRead:                       false,
		ResultPersistenceAuthorized:           false,
		RetentionEnforcementEnabled:           false,
		RedactionEnforcementEnabled:           false,
		DispatchDryRunExecuted:                false,
		DryRunResultPersisted:                 false,
		ResultVisibilityPersisted:             false,
		RuntimeDiagnosticsPersisted:           false,
		DispatchAuthorizationGranted:          false,
		OperatorRouteApprovalRequired:         true,
		OperatorRouteApprovalPresent:          false,
		ReceiptPresent:                        false,
		ReceiptAccepted:                       false,
		AuthorizationAccepted:                 false,
		ProductionReadiness:                   false,
		ProductionOwnershipReady:              false,
		RouteItemCount:                        len(items),
		RequiredRouteItemCount:                5,
		ReadyRouteItemCount:                   productionReceiptNotificationActionDryRunResultLookupRouteAuthorizationReadyCount(items),
		MissingRouteItemCount:                 productionReceiptNotificationActionDryRunResultLookupRouteAuthorizationMissingCount(items),
		RouteItems:                            items,
		RouteItemIDs:                          productionReceiptNotificationActionDryRunResultLookupRouteAuthorizationItemIDs(items),
		RequiredBeforeRouteEnablement: []string{
			"production-receipt-notification-action-dry-run-result-opaque-lookup-audit-preview",
			"separate route authorization receipt accepted by the Runtime owner",
			"separate owner-local lookup route implementation",
			"separate lookup persistence implementation",
			"separate KDE and Runtime consumer redaction review",
		},
		RuntimeOwned:                       true,
		GoRuntimeBacked:                    true,
		KDEPolicyOwner:                     false,
		OfficialDesktopOnly:                true,
		PlasmaForkRequired:                 false,
		PlasmaSourceModified:               false,
		SystemServiceStarted:               false,
		SessionBusClaimed:                  false,
		ProductionBusClaimed:               false,
		ProductionOwnerEnabled:             false,
		ProductionActivationReady:          false,
		WriteMethodsEnabled:                false,
		RuntimeWritesEnabled:               false,
		RequestObjectCreationEnabled:       false,
		RequestObjectDispatchEnabled:       false,
		RequestObjectPersistenceEnabled:    false,
		DispatchAuthorizationPersisted:     false,
		DispatchDryRunExecutionEnabled:     false,
		DryRunResultPersistenceEnabled:     false,
		ResultVisibilityPersistenceEnabled: false,
		RetentionPolicyPersistenceEnabled:  false,
		PortalRequestCreated:               false,
		RequestObjectsCreated:              false,
		RequestObjectsDispatched:           false,
		NotificationSent:                   false,
		NotificationDeliveryEnabled:        false,
		NotificationActionEnabled:          false,
		ReviewActionEnabled:                false,
		RenewActionEnabled:                 false,
		OpenCompatibilityCenterEnabled:     false,
		DismissActionEnabled:               false,
		SupportInfoActionEnabled:           false,
		CompatibilityCenterOpened:          false,
		CompatibilityCenterPersisted:       false,
		ReceiptWriterEnabled:               false,
		ReceiptPersistenceEnabled:          false,
		ReceiptLookupWritesEnabled:         false,
		ReceiptReplayEnabled:               false,
		ReceiptExpiryWriteEnabled:          false,
		ReceiptRevocationWriteEnabled:      false,
		DesktopFilesWritten:                false,
		MIMEAppsWritten:                    false,
		ShellConfigurationWritten:          false,
		SettingsPersisted:                  false,
		KRunnerIndexPersisted:              false,
		TaskManagerEntryActive:             false,
		KWinRuleApplied:                    false,
		LiveTrayBridgeEnabled:              false,
		TrayBridgePersisted:                false,
		AdapterInvocationEnabled:           false,
		BackendLaunchEnabled:               false,
		BackendProcessStarted:              false,
		SupportBundleExported:              false,
		SupportCaseCreated:                 false,
		SnapshotRestoreExecuted:            false,
		StateCleanupExecuted:               false,
		NetworkRequired:                    false,
		HostRootModified:                   false,
		PrivilegedContainerRequired:        false,
		RawCommandExposed:                  false,
		RawExecutableExposed:               false,
		BackendDetailsExposed:              false,
		BlockedActions: []string{
			"treat lookup route authorization audit as permission to expose a Runtime read route",
			"allow KDE or Runtime consumers to read opaque dry-run result summaries from this audit",
			"persist route authorization, lookup indexes, dry-run results, result visibility, or Runtime diagnostics from this audit",
			"accept caller-provided Runtime state-root paths, expose host paths, or read file contents from this audit",
			"execute dispatch dry runs, create Portal requests, write receipts, claim production ownership, launch compatibility engines, or mutate host root",
		},
		NextRequirements: []string{
			"Add an accepted route authorization receipt before enabling any owner-local result lookup route.",
			"Implement the owner-local lookup route separately while keeping caller state-root paths hidden.",
			"Review KDE and Runtime consumers separately before they can consume opaque result identifiers.",
			"Keep lookup persistence, result storage, diagnostics persistence, dry-run execution, and notification actions disabled until separate writers are reviewed.",
		},
		DesktopSafeSummary: "The production receipt notification action dry-run result lookup route authorization audit models the authorization boundary for future owner-local consumption of opaque dry-run result identifiers, but it authorizes no consumers, enables no lookup route, persists no route state, persists no results, executes no dry runs, creates no request objects, starts no services, launches no engines, exposes no paths, and mutates no host state.",
	}
	preview.AuthorizedRouteItemCount = productionReceiptNotificationActionDryRunResultLookupRouteAuthorizationAuthorizedCount(items)
	preview.EnabledRouteItemCount = productionReceiptNotificationActionDryRunResultLookupRouteAuthorizationEnabledCount(items)
	preview.PersistedRouteItemCount = productionReceiptNotificationActionDryRunResultLookupRouteAuthorizationPersistedRouteCount(items)
	preview.LookupEnabledItemCount = productionReceiptNotificationActionDryRunResultLookupRouteAuthorizationLookupEnabledCount(items)
	preview.PersistedLookupItemCount = productionReceiptNotificationActionDryRunResultLookupRouteAuthorizationLookupPersistedCount(items)
	preview.PathExposedRouteItemCount = productionReceiptNotificationActionDryRunResultLookupRouteAuthorizationPathExposedCount(items)
	preview.PersistedResultItemCount = productionReceiptNotificationActionDryRunResultLookupRouteAuthorizationPersistedResultCount(items)
	preview.ExecutedDryRunItemCount = productionReceiptNotificationActionDryRunResultLookupRouteAuthorizationExecutedCount(items)
	preview.CreatedRequestObjectCount = productionReceiptNotificationActionDryRunResultLookupRouteAuthorizationCreatedCount(items)
	preview.PortalRequestCreatedCount = productionReceiptNotificationActionDryRunResultLookupRouteAuthorizationPortalCount(items)
	preview.SideEffectRouteItemCount = productionReceiptNotificationActionDryRunResultLookupRouteAuthorizationSideEffectCount(items)
	checks := productionReceiptNotificationActionDryRunResultLookupRouteAuthorizationAuditChecks(preview)
	preview.Checks = checks
	preview.CheckIDs = productionReceiptNotificationActionDryRunResultLookupRouteAuthorizationCheckIDs(checks)
	preview.Counts = countProductionReceiptNotificationActionDryRunResultLookupRouteAuthorizationChecks(checks)
	if preview.Counts.Blocked == 0 {
		preview.AuditDecision = "production-receipt-notification-action-dry-run-result-lookup-route-authorization-audit-ready-route-disabled"
	}
	if err := validateNoBackendTerms(preview, "production receipt notification action dry-run result lookup route authorization audit preview"); err != nil {
		return ProductionReceiptNotificationActionDryRunResultLookupRouteAuthorizationAuditPreview{}, err
	}
	return preview, nil
}

type productionReceiptNotificationActionDryRunResultLookupRouteAuthorizationAuditSourceSet struct {
	OpaqueLookupAudit string
	DispatchSheet     string
}

func productionReceiptNotificationActionDryRunResultLookupRouteAuthorizationAuditSources(root string) productionReceiptNotificationActionDryRunResultLookupRouteAuthorizationAuditSourceSet {
	return productionReceiptNotificationActionDryRunResultLookupRouteAuthorizationAuditSourceSet{
		OpaqueLookupAudit: productionAuthorizationReadSources(root, []string{"internal/runtime/owner/production_receipt_notification_action_dry_run_result_opaque_lookup_audit.go"}),
		DispatchSheet:     productionAuthorizationReadSources(root, []string{"docs/claude-code-current-dispatch-picks.md"}),
	}
}

func productionReceiptNotificationActionDryRunResultLookupRouteAuthorizationItems(sources productionReceiptNotificationActionDryRunResultLookupRouteAuthorizationAuditSourceSet) []ProductionReceiptNotificationActionDryRunResultLookupRouteAuthorizationAuditItem {
	combined := sources.OpaqueLookupAudit + sources.DispatchSheet
	return []ProductionReceiptNotificationActionDryRunResultLookupRouteAuthorizationAuditItem{
		productionReceiptNotificationActionDryRunResultLookupRouteAuthorizationItem("review-receipt-dry-run-result-lookup-route-authorization", "review", "review-result-owner-local-lookup-route", "kde-review-and-runtime-diagnostics", "review-result-owner-managed-opaque-lookup", "review-dry-run-result-opaque-id", combined, []string{"review-receipt-dry-run-result-opaque-lookup", "lookup route authorization"}, "accept a route authorization receipt before enabling review result lookup"),
		productionReceiptNotificationActionDryRunResultLookupRouteAuthorizationItem("renew-receipt-dry-run-result-lookup-route-authorization", "renew", "renewal-result-owner-local-lookup-route", "kde-renewal-and-runtime-diagnostics", "renewal-result-owner-managed-opaque-lookup", "renewal-dry-run-result-opaque-id", combined, []string{"renew-receipt-dry-run-result-opaque-lookup", "lookup route authorization"}, "accept a route authorization receipt before enabling renewal result lookup"),
		productionReceiptNotificationActionDryRunResultLookupRouteAuthorizationItem("open-compatibility-center-dry-run-result-lookup-route-authorization", "open-compatibility-center", "navigation-result-owner-local-lookup-route", "kde-navigation-and-runtime-diagnostics", "navigation-result-owner-managed-opaque-lookup", "navigation-dry-run-result-opaque-id", combined, []string{"open-compatibility-center-dry-run-result-opaque-lookup", "lookup route authorization"}, "accept a route authorization receipt before enabling navigation result lookup"),
		productionReceiptNotificationActionDryRunResultLookupRouteAuthorizationItem("dismiss-receipt-dry-run-result-lookup-route-authorization", "dismiss", "dismissal-result-owner-local-lookup-route", "kde-dismissal-and-runtime-diagnostics", "dismissal-result-owner-managed-opaque-lookup", "dismissal-dry-run-result-opaque-id", combined, []string{"dismiss-receipt-dry-run-result-opaque-lookup", "lookup route authorization"}, "accept a route authorization receipt before enabling dismissal result lookup"),
		productionReceiptNotificationActionDryRunResultLookupRouteAuthorizationItem("support-info-dry-run-result-lookup-route-authorization", "support-info", "support-info-result-owner-local-lookup-route", "kde-support-and-runtime-diagnostics", "support-info-result-owner-managed-opaque-lookup", "support-info-dry-run-result-opaque-id", combined, []string{"support-info-dry-run-result-opaque-lookup", "lookup route authorization"}, "accept a route authorization receipt before enabling support-info result lookup"),
	}
}

func productionReceiptNotificationActionDryRunResultLookupRouteAuthorizationItem(id string, actionKind string, routeKind string, surfaceKind string, lookupKind string, opaqueResultID string, source string, tokens []string, nextRequirement string) ProductionReceiptNotificationActionDryRunResultLookupRouteAuthorizationAuditItem {
	ready := productionAuthorizationHasAll(source, tokens)
	status := "missing-lookup-route-authorization-evidence"
	if ready {
		status = "lookup-route-authorization-modeled-route-disabled"
	}
	return ProductionReceiptNotificationActionDryRunResultLookupRouteAuthorizationAuditItem{
		ID:                           id,
		ActionKind:                   actionKind,
		RouteKind:                    routeKind,
		SurfaceKind:                  surfaceKind,
		OpaqueLookupKind:             lookupKind,
		OpaqueResultID:               opaqueResultID,
		RequiredEvidence:             "future dry-run result consumers require owner-local lookup route authorization",
		EvidencePresent:              ready,
		RouteAuthorizationModeled:    ready,
		OwnerLocalRouteModeled:       ready,
		OpaqueLookupConsumed:         ready,
		OpaqueResultIDSupported:      ready,
		KDEConsumptionAuthorized:     false,
		RuntimeConsumptionAuthorized: false,
		LookupRouteAuthorized:        false,
		LookupRouteEnabled:           false,
		LookupRoutePersisted:         false,
		LookupEnabled:                false,
		LookupPersisted:              false,
		CallerStateRootRequired:      false,
		StateRootPathExposed:         false,
		FilePathsExposed:             false,
		FileContentRead:              false,
		UserVisible:                  ready,
		ReviewOnly:                   true,
		RuntimeOwned:                 true,
		GoRuntimeBacked:              true,
		KDEPolicyOwner:               false,
		OperatorApprovalRequired:     true,
		OperatorApprovalPresent:      false,
		ResultPersistenceAuthorized:  false,
		RetentionEnforcementEnabled:  false,
		RedactionEnforcementEnabled:  false,
		DispatchDryRunExecuted:       false,
		DryRunResultPersisted:        false,
		ResultVisibilityPersisted:    false,
		RuntimeDiagnosticsPersisted:  false,
		RequestObjectCreated:         false,
		RequestObjectDispatched:      false,
		PortalRequestCreated:         false,
		ReceiptAccepted:              false,
		AuthorizationAccepted:        false,
		CompatibilityCenterOpened:    false,
		SupportBundleExported:        false,
		SupportCaseCreated:           false,
		ProductionReadiness:          false,
		ProductionOwnershipReady:     false,
		SideEffectsDisabled:          true,
		HostRootModified:             false,
		InternalDetailsExposed:       false,
		RouteStatus:                  status,
		NextRequirement:              nextRequirement,
	}
}

func productionReceiptNotificationActionDryRunResultLookupRouteAuthorizationAuditChecks(preview ProductionReceiptNotificationActionDryRunResultLookupRouteAuthorizationAuditPreview) []ProductionReceiptNotificationActionDryRunResultLookupRouteAuthorizationAuditCheck {
	return []ProductionReceiptNotificationActionDryRunResultLookupRouteAuthorizationAuditCheck{
		productionReceiptNotificationActionDryRunResultLookupRouteAuthorizationCheck("opaque-lookup-audit-consumed", productionAuthorizationPassBlocked(preview.OpaqueLookupAuditConsumed), "The lookup route authorization audit consumes owner-managed opaque lookup evidence."),
		productionReceiptNotificationActionDryRunResultLookupRouteAuthorizationCheck("lookup-route-guidance-consumed", productionAuthorizationPassBlocked(preview.LookupRouteGuidanceConsumed), "The lookup route authorization audit consumes current dispatch guidance."),
		productionReceiptNotificationActionDryRunResultLookupRouteAuthorizationCheck("lookup-route-authorization-modeled-only", productionAuthorizationPassBlocked(preview.LookupRouteAuthorizationAuditRequired && preview.LookupRouteAuthorizationAuditModeled && preview.LookupRouteAuthorizationReady && preview.OwnerLocalLookupRouteModeled && preview.OpaqueResultIDSupported && !preview.LookupRouteAuthorized && !preview.LookupRouteEnabled && !preview.LookupRoutePersisted && !preview.KDEConsumptionAuthorized && !preview.RuntimeConsumptionAuthorized), "Route authorization is modeled without authorizing consumers, enabling routes, or persisting route state."),
		productionReceiptNotificationActionDryRunResultLookupRouteAuthorizationCheck("five-lookup-route-authorization-items-present", productionAuthorizationPassBlocked(preview.RouteItemCount == 5 && preview.RequiredRouteItemCount == 5 && preview.MissingRouteItemCount == 0), "Review, renew, open Compatibility Center, dismiss, and support-info route authorization boundaries are modeled."),
		productionReceiptNotificationActionDryRunResultLookupRouteAuthorizationCheck("route-authorization-items-ready-route-disabled", productionAuthorizationPassBlocked(preview.ReadyRouteItemCount == 5 && productionReceiptNotificationActionDryRunResultLookupRouteAuthorizationItemsReady(preview.RouteItems)), "Every route authorization item is ready while route authorization, route enablement, and lookup remain disabled."),
		productionReceiptNotificationActionDryRunResultLookupRouteAuthorizationCheck("route-lookup-persistence-result-persistence-and-execution-disabled", productionAuthorizationPassBlocked(!preview.LookupRouteAuthorized && !preview.LookupRouteEnabled && !preview.LookupRoutePersisted && !preview.OpaqueLookupEnabled && !preview.OpaqueLookupPersisted && !preview.RetentionPolicyPersistenceEnabled && !preview.DispatchDryRunExecutionEnabled && !preview.DryRunResultPersistenceEnabled && !preview.ResultVisibilityPersistenceEnabled && !preview.RuntimeDiagnosticsPersisted && preview.AuthorizedRouteItemCount == 0 && preview.EnabledRouteItemCount == 0 && preview.PersistedRouteItemCount == 0 && preview.LookupEnabledItemCount == 0 && preview.PersistedLookupItemCount == 0 && preview.PersistedResultItemCount == 0 && preview.ExecutedDryRunItemCount == 0), "Route authorization, route persistence, lookup, result persistence, diagnostics persistence, policy persistence, and dry-run execution remain disabled."),
		productionReceiptNotificationActionDryRunResultLookupRouteAuthorizationCheck("request-creation-dispatch-actions-and-portal-disabled", productionAuthorizationPassBlocked(!preview.RequestObjectCreationEnabled && !preview.RequestObjectDispatchEnabled && !preview.RequestObjectPersistenceEnabled && !preview.RequestObjectsCreated && !preview.RequestObjectsDispatched && !preview.PortalRequestCreated && !preview.NotificationActionEnabled && preview.CreatedRequestObjectCount == 0 && preview.PortalRequestCreatedCount == 0), "Request creation, dispatch, persistence, notification actions, and Portal requests remain disabled."),
		productionReceiptNotificationActionDryRunResultLookupRouteAuthorizationCheck("receipt-notification-navigation-and-support-writes-disabled", productionAuthorizationPassBlocked(!preview.NotificationSent && !preview.NotificationDeliveryEnabled && !preview.CompatibilityCenterOpened && !preview.CompatibilityCenterPersisted && !preview.ReceiptWriterEnabled && !preview.ReceiptPersistenceEnabled && !preview.SupportBundleExported && !preview.SupportCaseCreated), "Notification, navigation, receipt, Compatibility Center, and support writes remain disabled."),
		productionReceiptNotificationActionDryRunResultLookupRouteAuthorizationCheck("production-and-host-boundary-closed", productionAuthorizationPassBlocked(!preview.SystemServiceStarted && !preview.SessionBusClaimed && !preview.ProductionBusClaimed && !preview.ProductionOwnerEnabled && !preview.ProductionActivationReady && !preview.ProductionReadiness && !preview.ProductionOwnershipReady && !preview.AdapterInvocationEnabled && !preview.BackendLaunchEnabled && !preview.BackendProcessStarted && !preview.FileContentRead && !preview.FilePathsExposed && !preview.NetworkRequired && !preview.HostRootModified && !preview.PrivilegedContainerRequired && !preview.StateRootPathExposed && !preview.RawCommandExposed && !preview.RawExecutableExposed && !preview.BackendDetailsExposed && preview.PathExposedRouteItemCount == 0 && productionReceiptNotificationActionDryRunResultLookupRouteAuthorizationItemsKeepHostClosed(preview.RouteItems)), "Production ownership, engine launch, unsafe data exposure, network, privilege, path exposure, internal details, and host mutation remain disabled."),
	}
}

func productionReceiptNotificationActionDryRunResultLookupRouteAuthorizationCheck(id string, status string, summary string) ProductionReceiptNotificationActionDryRunResultLookupRouteAuthorizationAuditCheck {
	return ProductionReceiptNotificationActionDryRunResultLookupRouteAuthorizationAuditCheck{ID: id, Status: status, Summary: summary}
}

func productionReceiptNotificationActionDryRunResultLookupRouteAuthorizationOpaqueLookupReady(source string) bool {
	return productionAuthorizationHasAll(source, []string{
		"production-receipt-notification-action-dry-run-result-opaque-lookup-audit-preview",
		"production-receipt-notification-action-dry-run-result-opaque-lookup-audit-ready-lookup-disabled",
		"review-receipt-dry-run-result-opaque-lookup",
		"renew-receipt-dry-run-result-opaque-lookup",
		"open-compatibility-center-dry-run-result-opaque-lookup",
		"dismiss-receipt-dry-run-result-opaque-lookup",
		"support-info-dry-run-result-opaque-lookup",
		"OpaqueLookupEnabled",
	})
}

func productionReceiptNotificationActionDryRunResultLookupRouteAuthorizationGuidanceReady(source string) bool {
	return productionAuthorizationHasAll(source, []string{"lookup route authorization", "opaque result identifiers", "side effects"})
}

func productionReceiptNotificationActionDryRunResultLookupRouteAuthorizationReadyCount(items []ProductionReceiptNotificationActionDryRunResultLookupRouteAuthorizationAuditItem) int {
	count := 0
	for _, item := range items {
		if item.EvidencePresent {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionDryRunResultLookupRouteAuthorizationMissingCount(items []ProductionReceiptNotificationActionDryRunResultLookupRouteAuthorizationAuditItem) int {
	count := 0
	for _, item := range items {
		if !item.EvidencePresent {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionDryRunResultLookupRouteAuthorizationAuthorizedCount(items []ProductionReceiptNotificationActionDryRunResultLookupRouteAuthorizationAuditItem) int {
	count := 0
	for _, item := range items {
		if item.LookupRouteAuthorized || item.KDEConsumptionAuthorized || item.RuntimeConsumptionAuthorized {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionDryRunResultLookupRouteAuthorizationEnabledCount(items []ProductionReceiptNotificationActionDryRunResultLookupRouteAuthorizationAuditItem) int {
	count := 0
	for _, item := range items {
		if item.LookupRouteEnabled {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionDryRunResultLookupRouteAuthorizationPersistedRouteCount(items []ProductionReceiptNotificationActionDryRunResultLookupRouteAuthorizationAuditItem) int {
	count := 0
	for _, item := range items {
		if item.LookupRoutePersisted {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionDryRunResultLookupRouteAuthorizationLookupEnabledCount(items []ProductionReceiptNotificationActionDryRunResultLookupRouteAuthorizationAuditItem) int {
	count := 0
	for _, item := range items {
		if item.LookupEnabled {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionDryRunResultLookupRouteAuthorizationLookupPersistedCount(items []ProductionReceiptNotificationActionDryRunResultLookupRouteAuthorizationAuditItem) int {
	count := 0
	for _, item := range items {
		if item.LookupPersisted {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionDryRunResultLookupRouteAuthorizationPathExposedCount(items []ProductionReceiptNotificationActionDryRunResultLookupRouteAuthorizationAuditItem) int {
	count := 0
	for _, item := range items {
		if item.CallerStateRootRequired || item.StateRootPathExposed || item.FilePathsExposed || item.FileContentRead {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionDryRunResultLookupRouteAuthorizationPersistedResultCount(items []ProductionReceiptNotificationActionDryRunResultLookupRouteAuthorizationAuditItem) int {
	count := 0
	for _, item := range items {
		if item.DryRunResultPersisted || item.ResultVisibilityPersisted || item.RuntimeDiagnosticsPersisted {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionDryRunResultLookupRouteAuthorizationExecutedCount(items []ProductionReceiptNotificationActionDryRunResultLookupRouteAuthorizationAuditItem) int {
	count := 0
	for _, item := range items {
		if item.DispatchDryRunExecuted {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionDryRunResultLookupRouteAuthorizationCreatedCount(items []ProductionReceiptNotificationActionDryRunResultLookupRouteAuthorizationAuditItem) int {
	count := 0
	for _, item := range items {
		if item.RequestObjectCreated {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionDryRunResultLookupRouteAuthorizationPortalCount(items []ProductionReceiptNotificationActionDryRunResultLookupRouteAuthorizationAuditItem) int {
	count := 0
	for _, item := range items {
		if item.PortalRequestCreated {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionDryRunResultLookupRouteAuthorizationSideEffectCount(items []ProductionReceiptNotificationActionDryRunResultLookupRouteAuthorizationAuditItem) int {
	count := 0
	for _, item := range items {
		if item.KDEConsumptionAuthorized || item.RuntimeConsumptionAuthorized || item.LookupRouteAuthorized || item.LookupRouteEnabled || item.LookupRoutePersisted || item.LookupEnabled || item.LookupPersisted || item.ResultPersistenceAuthorized || item.RetentionEnforcementEnabled || item.RedactionEnforcementEnabled || item.DispatchDryRunExecuted || item.DryRunResultPersisted || item.ResultVisibilityPersisted || item.RuntimeDiagnosticsPersisted || item.RequestObjectCreated || item.RequestObjectDispatched || item.PortalRequestCreated || item.ReceiptAccepted || item.AuthorizationAccepted || item.CompatibilityCenterOpened || item.SupportBundleExported || item.SupportCaseCreated || item.HostRootModified || item.InternalDetailsExposed {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionDryRunResultLookupRouteAuthorizationItemIDs(items []ProductionReceiptNotificationActionDryRunResultLookupRouteAuthorizationAuditItem) []string {
	ids := make([]string, 0, len(items))
	for _, item := range items {
		ids = append(ids, item.ID)
	}
	return ids
}

func productionReceiptNotificationActionDryRunResultLookupRouteAuthorizationCheckIDs(checks []ProductionReceiptNotificationActionDryRunResultLookupRouteAuthorizationAuditCheck) []string {
	ids := make([]string, 0, len(checks))
	for _, check := range checks {
		ids = append(ids, check.ID)
	}
	return ids
}

func countProductionReceiptNotificationActionDryRunResultLookupRouteAuthorizationChecks(checks []ProductionReceiptNotificationActionDryRunResultLookupRouteAuthorizationAuditCheck) ProductionReceiptNotificationActionDryRunResultLookupRouteAuthorizationAuditCounts {
	counts := ProductionReceiptNotificationActionDryRunResultLookupRouteAuthorizationAuditCounts{Total: len(checks)}
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

func productionReceiptNotificationActionDryRunResultLookupRouteAuthorizationItemsReady(items []ProductionReceiptNotificationActionDryRunResultLookupRouteAuthorizationAuditItem) bool {
	if len(items) != 5 {
		return false
	}
	for _, item := range items {
		if !item.EvidencePresent || !item.RouteAuthorizationModeled || !item.OwnerLocalRouteModeled || !item.OpaqueLookupConsumed || !item.OpaqueResultIDSupported || item.RouteStatus != "lookup-route-authorization-modeled-route-disabled" {
			return false
		}
		if item.KDEConsumptionAuthorized || item.RuntimeConsumptionAuthorized || item.LookupRouteAuthorized || item.LookupRouteEnabled || item.LookupRoutePersisted || item.LookupEnabled || item.LookupPersisted || item.CallerStateRootRequired || item.StateRootPathExposed || item.FilePathsExposed || item.FileContentRead || item.OperatorApprovalPresent || item.ResultPersistenceAuthorized || item.RetentionEnforcementEnabled || item.RedactionEnforcementEnabled || item.DispatchDryRunExecuted || item.DryRunResultPersisted || item.ResultVisibilityPersisted || item.RuntimeDiagnosticsPersisted || item.RequestObjectCreated || item.RequestObjectDispatched || item.PortalRequestCreated || item.ReceiptAccepted || item.AuthorizationAccepted || item.CompatibilityCenterOpened || item.SupportBundleExported || item.SupportCaseCreated || item.ProductionReadiness || item.ProductionOwnershipReady || item.KDEPolicyOwner || !item.ReviewOnly || !item.SideEffectsDisabled {
			return false
		}
	}
	return true
}

func productionReceiptNotificationActionDryRunResultLookupRouteAuthorizationItemsKeepHostClosed(items []ProductionReceiptNotificationActionDryRunResultLookupRouteAuthorizationAuditItem) bool {
	for _, item := range items {
		if item.CallerStateRootRequired || item.StateRootPathExposed || item.FilePathsExposed || item.FileContentRead || item.HostRootModified || item.InternalDetailsExposed {
			return false
		}
	}
	return true
}
