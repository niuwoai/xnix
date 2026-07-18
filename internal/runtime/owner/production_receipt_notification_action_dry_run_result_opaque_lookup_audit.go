package owner

type ProductionReceiptNotificationActionDryRunResultOpaqueLookupAuditPreview struct {
	Version                               string                                                                  `json:"version"`
	SchemaVersion                         string                                                                  `json:"schema_version"`
	RequestType                           string                                                                  `json:"request_type"`
	AuditType                             string                                                                  `json:"audit_type"`
	Source                                string                                                                  `json:"source"`
	AuditDecision                         string                                                                  `json:"audit_decision"`
	ReceiptSchema                         string                                                                  `json:"receipt_schema"`
	OpaqueReceiptID                       string                                                                  `json:"opaque_receipt_id"`
	OpaqueLookupAuditRequired             bool                                                                    `json:"opaque_lookup_audit_required"`
	OpaqueLookupAuditModeled              bool                                                                    `json:"opaque_lookup_audit_modeled"`
	RetentionRedactionPolicyAuditConsumed bool                                                                    `json:"retention_redaction_policy_audit_consumed"`
	OpaqueLookupGuidanceConsumed          bool                                                                    `json:"opaque_lookup_guidance_consumed"`
	OpaqueLookupBoundaryReady             bool                                                                    `json:"opaque_lookup_boundary_ready"`
	OwnerManagedOpaqueLookupModeled       bool                                                                    `json:"owner_managed_opaque_lookup_modeled"`
	OpaqueLookupEnabled                   bool                                                                    `json:"opaque_lookup_enabled"`
	OpaqueLookupPersisted                 bool                                                                    `json:"opaque_lookup_persisted"`
	OpaqueResultIDSupported               bool                                                                    `json:"opaque_result_id_supported"`
	CallerStateRootRequired               bool                                                                    `json:"caller_state_root_required"`
	StateRootPathExposed                  bool                                                                    `json:"state_root_path_exposed"`
	FilePathsExposed                      bool                                                                    `json:"file_paths_exposed"`
	FileContentRead                       bool                                                                    `json:"file_content_read"`
	ResultPersistenceAuthorized           bool                                                                    `json:"result_persistence_authorized"`
	RetentionEnforcementEnabled           bool                                                                    `json:"retention_enforcement_enabled"`
	RedactionEnforcementEnabled           bool                                                                    `json:"redaction_enforcement_enabled"`
	DispatchDryRunExecuted                bool                                                                    `json:"dispatch_dry_run_executed"`
	DryRunResultPersisted                 bool                                                                    `json:"dry_run_result_persisted"`
	ResultVisibilityPersisted             bool                                                                    `json:"result_visibility_persisted"`
	RuntimeDiagnosticsPersisted           bool                                                                    `json:"runtime_diagnostics_persisted"`
	DispatchAuthorizationGranted          bool                                                                    `json:"dispatch_authorization_granted"`
	OperatorLookupApprovalRequired        bool                                                                    `json:"operator_lookup_approval_required"`
	OperatorLookupApprovalPresent         bool                                                                    `json:"operator_lookup_approval_present"`
	ReceiptPresent                        bool                                                                    `json:"receipt_present"`
	ReceiptAccepted                       bool                                                                    `json:"receipt_accepted"`
	AuthorizationAccepted                 bool                                                                    `json:"authorization_accepted"`
	ProductionReadiness                   bool                                                                    `json:"production_readiness"`
	ProductionOwnershipReady              bool                                                                    `json:"production_ownership_ready"`
	LookupItemCount                       int                                                                     `json:"lookup_item_count"`
	RequiredLookupItemCount               int                                                                     `json:"required_lookup_item_count"`
	ReadyLookupItemCount                  int                                                                     `json:"ready_lookup_item_count"`
	MissingLookupItemCount                int                                                                     `json:"missing_lookup_item_count"`
	EnabledLookupItemCount                int                                                                     `json:"enabled_lookup_item_count"`
	PersistedLookupItemCount              int                                                                     `json:"persisted_lookup_item_count"`
	PathExposedLookupItemCount            int                                                                     `json:"path_exposed_lookup_item_count"`
	PersistedResultItemCount              int                                                                     `json:"persisted_result_item_count"`
	ExecutedDryRunItemCount               int                                                                     `json:"executed_dry_run_item_count"`
	CreatedRequestObjectCount             int                                                                     `json:"created_request_object_count"`
	PortalRequestCreatedCount             int                                                                     `json:"portal_request_created_count"`
	SideEffectLookupItemCount             int                                                                     `json:"side_effect_lookup_item_count"`
	LookupItems                           []ProductionReceiptNotificationActionDryRunResultOpaqueLookupAuditItem  `json:"lookup_items"`
	LookupItemIDs                         []string                                                                `json:"lookup_item_ids"`
	RequiredBeforeLookupEnablement        []string                                                                `json:"required_before_lookup_enablement"`
	Checks                                []ProductionReceiptNotificationActionDryRunResultOpaqueLookupAuditCheck `json:"checks"`
	CheckIDs                              []string                                                                `json:"check_ids"`
	Counts                                ProductionReceiptNotificationActionDryRunResultOpaqueLookupAuditCounts  `json:"counts"`
	RuntimeOwned                          bool                                                                    `json:"runtime_owned"`
	GoRuntimeBacked                       bool                                                                    `json:"go_runtime_backed"`
	KDEPolicyOwner                        bool                                                                    `json:"kde_policy_owner"`
	OfficialDesktopOnly                   bool                                                                    `json:"official_desktop_only"`
	PlasmaForkRequired                    bool                                                                    `json:"plasma_fork_required"`
	PlasmaSourceModified                  bool                                                                    `json:"plasma_source_modified"`
	SystemServiceStarted                  bool                                                                    `json:"system_service_started"`
	SessionBusClaimed                     bool                                                                    `json:"session_bus_claimed"`
	ProductionBusClaimed                  bool                                                                    `json:"production_bus_claimed"`
	ProductionOwnerEnabled                bool                                                                    `json:"production_owner_enabled"`
	ProductionActivationReady             bool                                                                    `json:"production_activation_ready"`
	WriteMethodsEnabled                   bool                                                                    `json:"write_methods_enabled"`
	RuntimeWritesEnabled                  bool                                                                    `json:"runtime_writes_enabled"`
	RequestObjectCreationEnabled          bool                                                                    `json:"request_object_creation_enabled"`
	RequestObjectDispatchEnabled          bool                                                                    `json:"request_object_dispatch_enabled"`
	RequestObjectPersistenceEnabled       bool                                                                    `json:"request_object_persistence_enabled"`
	DispatchAuthorizationPersisted        bool                                                                    `json:"dispatch_authorization_persisted"`
	DispatchDryRunExecutionEnabled        bool                                                                    `json:"dispatch_dry_run_execution_enabled"`
	DryRunResultPersistenceEnabled        bool                                                                    `json:"dry_run_result_persistence_enabled"`
	ResultVisibilityPersistenceEnabled    bool                                                                    `json:"result_visibility_persistence_enabled"`
	RetentionPolicyPersistenceEnabled     bool                                                                    `json:"retention_policy_persistence_enabled"`
	PortalRequestCreated                  bool                                                                    `json:"portal_request_created"`
	RequestObjectsCreated                 bool                                                                    `json:"request_objects_created"`
	RequestObjectsDispatched              bool                                                                    `json:"request_objects_dispatched"`
	NotificationSent                      bool                                                                    `json:"notification_sent"`
	NotificationDeliveryEnabled           bool                                                                    `json:"notification_delivery_enabled"`
	NotificationActionEnabled             bool                                                                    `json:"notification_action_enabled"`
	ReviewActionEnabled                   bool                                                                    `json:"review_action_enabled"`
	RenewActionEnabled                    bool                                                                    `json:"renew_action_enabled"`
	OpenCompatibilityCenterEnabled        bool                                                                    `json:"open_compatibility_center_enabled"`
	DismissActionEnabled                  bool                                                                    `json:"dismiss_action_enabled"`
	SupportInfoActionEnabled              bool                                                                    `json:"support_info_action_enabled"`
	CompatibilityCenterOpened             bool                                                                    `json:"compatibility_center_opened"`
	CompatibilityCenterPersisted          bool                                                                    `json:"compatibility_center_persisted"`
	ReceiptWriterEnabled                  bool                                                                    `json:"receipt_writer_enabled"`
	ReceiptPersistenceEnabled             bool                                                                    `json:"receipt_persistence_enabled"`
	ReceiptLookupWritesEnabled            bool                                                                    `json:"receipt_lookup_writes_enabled"`
	ReceiptReplayEnabled                  bool                                                                    `json:"receipt_replay_enabled"`
	ReceiptExpiryWriteEnabled             bool                                                                    `json:"receipt_expiry_write_enabled"`
	ReceiptRevocationWriteEnabled         bool                                                                    `json:"receipt_revocation_write_enabled"`
	DesktopFilesWritten                   bool                                                                    `json:"desktop_files_written"`
	MIMEAppsWritten                       bool                                                                    `json:"mimeapps_written"`
	ShellConfigurationWritten             bool                                                                    `json:"shell_configuration_written"`
	SettingsPersisted                     bool                                                                    `json:"settings_persisted"`
	KRunnerIndexPersisted                 bool                                                                    `json:"krunner_index_persisted"`
	TaskManagerEntryActive                bool                                                                    `json:"task_manager_entry_active"`
	KWinRuleApplied                       bool                                                                    `json:"kwin_rule_applied"`
	LiveTrayBridgeEnabled                 bool                                                                    `json:"live_tray_bridge_enabled"`
	TrayBridgePersisted                   bool                                                                    `json:"tray_bridge_persisted"`
	AdapterInvocationEnabled              bool                                                                    `json:"adapter_invocation_enabled"`
	BackendLaunchEnabled                  bool                                                                    `json:"backend_launch_enabled"`
	BackendProcessStarted                 bool                                                                    `json:"backend_process_started"`
	SupportBundleExported                 bool                                                                    `json:"support_bundle_exported"`
	SupportCaseCreated                    bool                                                                    `json:"support_case_created"`
	SnapshotRestoreExecuted               bool                                                                    `json:"snapshot_restore_executed"`
	StateCleanupExecuted                  bool                                                                    `json:"state_cleanup_executed"`
	NetworkRequired                       bool                                                                    `json:"network_required"`
	HostRootModified                      bool                                                                    `json:"host_root_modified"`
	PrivilegedContainerRequired           bool                                                                    `json:"privileged_container_required"`
	RawCommandExposed                     bool                                                                    `json:"raw_command_exposed"`
	RawExecutableExposed                  bool                                                                    `json:"raw_executable_exposed"`
	BackendDetailsExposed                 bool                                                                    `json:"backend_details_exposed"`
	BlockedActions                        []string                                                                `json:"blocked_actions"`
	NextRequirements                      []string                                                                `json:"next_requirements"`
	DesktopSafeSummary                    string                                                                  `json:"desktop_safe_summary"`
}

type ProductionReceiptNotificationActionDryRunResultOpaqueLookupAuditItem struct {
	ID                           string `json:"id"`
	ActionKind                   string `json:"action_kind"`
	RetentionPolicyKind          string `json:"retention_policy_kind"`
	OpaqueLookupKind             string `json:"opaque_lookup_kind"`
	OpaqueResultID               string `json:"opaque_result_id"`
	RequiredEvidence             string `json:"required_evidence"`
	EvidencePresent              bool   `json:"evidence_present"`
	OpaqueLookupModeled          bool   `json:"opaque_lookup_modeled"`
	OwnerManagedLookup           bool   `json:"owner_managed_lookup"`
	OpaqueResultIDSupported      bool   `json:"opaque_result_id_supported"`
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
	LookupEnabled                bool   `json:"lookup_enabled"`
	LookupPersisted              bool   `json:"lookup_persisted"`
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
	CompatibilityCenterPersisted bool   `json:"compatibility_center_persisted"`
	SupportBundleExported        bool   `json:"support_bundle_exported"`
	SupportCaseCreated           bool   `json:"support_case_created"`
	ProductionReadiness          bool   `json:"production_readiness"`
	ProductionOwnershipReady     bool   `json:"production_ownership_ready"`
	SideEffectsDisabled          bool   `json:"side_effects_disabled"`
	HostRootModified             bool   `json:"host_root_modified"`
	InternalDetailsExposed       bool   `json:"internal_details_exposed"`
	LookupStatus                 string `json:"lookup_status"`
	NextRequirement              string `json:"next_requirement"`
}

type ProductionReceiptNotificationActionDryRunResultOpaqueLookupAuditCheck struct {
	ID      string `json:"id"`
	Status  string `json:"status"`
	Summary string `json:"summary"`
}

type ProductionReceiptNotificationActionDryRunResultOpaqueLookupAuditCounts struct {
	Total   int `json:"total"`
	Passed  int `json:"passed"`
	Pending int `json:"pending"`
	Blocked int `json:"blocked"`
}

func NewProductionReceiptNotificationActionDryRunResultOpaqueLookupAuditPreview(root string) (ProductionReceiptNotificationActionDryRunResultOpaqueLookupAuditPreview, error) {
	version, err := readProductionDBusGateReviewVersion(root)
	if err != nil {
		return ProductionReceiptNotificationActionDryRunResultOpaqueLookupAuditPreview{}, err
	}
	sources := productionReceiptNotificationActionDryRunResultOpaqueLookupAuditSources(root)
	items := productionReceiptNotificationActionDryRunResultOpaqueLookupItems(sources)
	retentionReady := productionReceiptNotificationActionDryRunResultOpaqueLookupRetentionReady(sources.RetentionRedactionPolicyAudit)
	guidanceReady := productionReceiptNotificationActionDryRunResultOpaqueLookupGuidanceReady(sources.DispatchSheet)
	lookupReady := retentionReady && guidanceReady && productionReceiptNotificationActionDryRunResultOpaqueLookupItemsReady(items)
	preview := ProductionReceiptNotificationActionDryRunResultOpaqueLookupAuditPreview{
		Version:                               version,
		SchemaVersion:                         "xnix.runtime.production_receipt_notification_action_dry_run_result_opaque_lookup_audit.v1",
		RequestType:                           "production-receipt-notification-action-dry-run-result-opaque-lookup-audit-preview",
		AuditType:                             "receipt-notification-action-dry-run-result-opaque-lookup-audit",
		Source:                                "production-receipt-notification-action-dry-run-result-retention-redaction-policy-audit-preview+claude-code-current-dispatch-picks",
		AuditDecision:                         "production-receipt-notification-action-dry-run-result-opaque-lookup-audit-blocked",
		ReceiptSchema:                         "xnix.runtime.production_dbus_human_authorization_receipt.v1",
		OpaqueReceiptID:                       ProductionDBusHumanAuthorizationReceiptID,
		OpaqueLookupAuditRequired:             true,
		OpaqueLookupAuditModeled:              true,
		RetentionRedactionPolicyAuditConsumed: retentionReady,
		OpaqueLookupGuidanceConsumed:          guidanceReady,
		OpaqueLookupBoundaryReady:             lookupReady,
		OwnerManagedOpaqueLookupModeled:       lookupReady,
		OpaqueLookupEnabled:                   false,
		OpaqueLookupPersisted:                 false,
		OpaqueResultIDSupported:               lookupReady,
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
		OperatorLookupApprovalRequired:        true,
		OperatorLookupApprovalPresent:         false,
		ReceiptPresent:                        false,
		ReceiptAccepted:                       false,
		AuthorizationAccepted:                 false,
		ProductionReadiness:                   false,
		ProductionOwnershipReady:              false,
		LookupItemCount:                       len(items),
		RequiredLookupItemCount:               5,
		ReadyLookupItemCount:                  productionReceiptNotificationActionDryRunResultOpaqueLookupReadyCount(items),
		MissingLookupItemCount:                productionReceiptNotificationActionDryRunResultOpaqueLookupMissingCount(items),
		LookupItems:                           items,
		LookupItemIDs:                         productionReceiptNotificationActionDryRunResultOpaqueLookupItemIDs(items),
		RequiredBeforeLookupEnablement: []string{
			"production-receipt-notification-action-dry-run-result-retention-redaction-policy-audit-preview",
			"accepted opaque result persistence authorization receipt",
			"separate owner-managed lookup route authorization",
			"separate lookup persistence implementation",
			"separate result writer implementation",
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
			"treat opaque lookup audit as permission to enable lookups",
			"persist lookup indexes, dry-run results, result visibility, or Runtime diagnostics from this audit",
			"accept caller-provided Runtime state-root paths, expose host paths, or read file contents from this audit",
			"execute dispatch dry runs, create Portal requests, write receipts, or export support data from this audit",
			"claim production D-Bus ownership, start services, launch compatibility engines, or mutate host root",
		},
		NextRequirements: []string{
			"Add route authorization for owner-managed opaque dry-run result lookup before any KDE or Runtime surface can consume it.",
			"Implement lookup persistence only after route authorization and result writer boundaries are reviewed.",
			"Keep dry-run execution, result storage, Runtime diagnostics persistence, and notification actions disabled until separate writers are reviewed.",
			"Keep callers away from Runtime state-root paths and raw file locations.",
		},
		DesktopSafeSummary: "The production receipt notification action dry-run result opaque lookup audit models owner-managed opaque identifiers for future stored dry-run result summaries, but it enables no lookup route, persists no lookup index, persists no results, executes no dry runs, creates no request objects, starts no services, launches no engines, exposes no paths, and mutates no host state.",
	}
	preview.EnabledLookupItemCount = productionReceiptNotificationActionDryRunResultOpaqueLookupEnabledCount(items)
	preview.PersistedLookupItemCount = productionReceiptNotificationActionDryRunResultOpaqueLookupPersistedCount(items)
	preview.PathExposedLookupItemCount = productionReceiptNotificationActionDryRunResultOpaqueLookupPathExposedCount(items)
	preview.PersistedResultItemCount = productionReceiptNotificationActionDryRunResultOpaqueLookupPersistedResultCount(items)
	preview.ExecutedDryRunItemCount = productionReceiptNotificationActionDryRunResultOpaqueLookupExecutedCount(items)
	preview.CreatedRequestObjectCount = productionReceiptNotificationActionDryRunResultOpaqueLookupCreatedCount(items)
	preview.PortalRequestCreatedCount = productionReceiptNotificationActionDryRunResultOpaqueLookupPortalCount(items)
	preview.SideEffectLookupItemCount = productionReceiptNotificationActionDryRunResultOpaqueLookupSideEffectCount(items)
	checks := productionReceiptNotificationActionDryRunResultOpaqueLookupAuditChecks(preview)
	preview.Checks = checks
	preview.CheckIDs = productionReceiptNotificationActionDryRunResultOpaqueLookupCheckIDs(checks)
	preview.Counts = countProductionReceiptNotificationActionDryRunResultOpaqueLookupChecks(checks)
	if preview.Counts.Blocked == 0 {
		preview.AuditDecision = "production-receipt-notification-action-dry-run-result-opaque-lookup-audit-ready-lookup-disabled"
	}
	if err := validateNoBackendTerms(preview, "production receipt notification action dry-run result opaque lookup audit preview"); err != nil {
		return ProductionReceiptNotificationActionDryRunResultOpaqueLookupAuditPreview{}, err
	}
	return preview, nil
}

type productionReceiptNotificationActionDryRunResultOpaqueLookupAuditSourceSet struct {
	RetentionRedactionPolicyAudit string
	DispatchSheet                 string
}

func productionReceiptNotificationActionDryRunResultOpaqueLookupAuditSources(root string) productionReceiptNotificationActionDryRunResultOpaqueLookupAuditSourceSet {
	return productionReceiptNotificationActionDryRunResultOpaqueLookupAuditSourceSet{
		RetentionRedactionPolicyAudit: productionAuthorizationReadSources(root, []string{"internal/runtime/owner/production_receipt_notification_action_dry_run_result_retention_redaction_policy_audit.go"}),
		DispatchSheet:                 productionAuthorizationReadSources(root, []string{"docs/claude-code-current-dispatch-picks.md"}),
	}
}

func productionReceiptNotificationActionDryRunResultOpaqueLookupItems(sources productionReceiptNotificationActionDryRunResultOpaqueLookupAuditSourceSet) []ProductionReceiptNotificationActionDryRunResultOpaqueLookupAuditItem {
	combined := sources.RetentionRedactionPolicyAudit + sources.DispatchSheet
	return []ProductionReceiptNotificationActionDryRunResultOpaqueLookupAuditItem{
		productionReceiptNotificationActionDryRunResultOpaqueLookupItem("review-receipt-dry-run-result-opaque-lookup", "review", "review-result-retention-policy", "review-result-owner-managed-opaque-lookup", "review-dry-run-result-opaque-id", combined, []string{"review-receipt-dry-run-result-retention-redaction-policy", "owner-managed opaque lookup"}, "authorize an owner-local read route before enabling review result lookup"),
		productionReceiptNotificationActionDryRunResultOpaqueLookupItem("renew-receipt-dry-run-result-opaque-lookup", "renew", "renewal-result-retention-policy", "renewal-result-owner-managed-opaque-lookup", "renewal-dry-run-result-opaque-id", combined, []string{"renew-receipt-dry-run-result-retention-redaction-policy", "owner-managed opaque lookup"}, "authorize an owner-local read route before enabling renewal result lookup"),
		productionReceiptNotificationActionDryRunResultOpaqueLookupItem("open-compatibility-center-dry-run-result-opaque-lookup", "open-compatibility-center", "navigation-result-retention-policy", "navigation-result-owner-managed-opaque-lookup", "navigation-dry-run-result-opaque-id", combined, []string{"open-compatibility-center-dry-run-result-retention-redaction-policy", "owner-managed opaque lookup"}, "authorize an owner-local read route before enabling navigation result lookup"),
		productionReceiptNotificationActionDryRunResultOpaqueLookupItem("dismiss-receipt-dry-run-result-opaque-lookup", "dismiss", "dismissal-result-retention-policy", "dismissal-result-owner-managed-opaque-lookup", "dismissal-dry-run-result-opaque-id", combined, []string{"dismiss-receipt-dry-run-result-retention-redaction-policy", "owner-managed opaque lookup"}, "authorize an owner-local read route before enabling dismissal result lookup"),
		productionReceiptNotificationActionDryRunResultOpaqueLookupItem("support-info-dry-run-result-opaque-lookup", "support-info", "support-info-result-retention-policy", "support-info-result-owner-managed-opaque-lookup", "support-info-dry-run-result-opaque-id", combined, []string{"support-info-dry-run-result-retention-redaction-policy", "owner-managed opaque lookup"}, "authorize an owner-local read route before enabling support-info result lookup"),
	}
}

func productionReceiptNotificationActionDryRunResultOpaqueLookupItem(id string, actionKind string, retentionKind string, lookupKind string, opaqueResultID string, source string, tokens []string, nextRequirement string) ProductionReceiptNotificationActionDryRunResultOpaqueLookupAuditItem {
	ready := productionAuthorizationHasAll(source, tokens)
	status := "missing-opaque-lookup-evidence"
	if ready {
		status = "opaque-lookup-modeled-lookup-disabled"
	}
	return ProductionReceiptNotificationActionDryRunResultOpaqueLookupAuditItem{
		ID:                           id,
		ActionKind:                   actionKind,
		RetentionPolicyKind:          retentionKind,
		OpaqueLookupKind:             lookupKind,
		OpaqueResultID:               opaqueResultID,
		RequiredEvidence:             "future stored dry-run result summary requires an owner-managed opaque lookup boundary",
		EvidencePresent:              ready,
		OpaqueLookupModeled:          ready,
		OwnerManagedLookup:           ready,
		OpaqueResultIDSupported:      ready,
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
		LookupEnabled:                false,
		LookupPersisted:              false,
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
		CompatibilityCenterPersisted: false,
		SupportBundleExported:        false,
		SupportCaseCreated:           false,
		ProductionReadiness:          false,
		ProductionOwnershipReady:     false,
		SideEffectsDisabled:          true,
		HostRootModified:             false,
		InternalDetailsExposed:       false,
		LookupStatus:                 status,
		NextRequirement:              nextRequirement,
	}
}

func productionReceiptNotificationActionDryRunResultOpaqueLookupAuditChecks(preview ProductionReceiptNotificationActionDryRunResultOpaqueLookupAuditPreview) []ProductionReceiptNotificationActionDryRunResultOpaqueLookupAuditCheck {
	return []ProductionReceiptNotificationActionDryRunResultOpaqueLookupAuditCheck{
		productionReceiptNotificationActionDryRunResultOpaqueLookupCheck("retention-redaction-policy-audit-consumed", productionAuthorizationPassBlocked(preview.RetentionRedactionPolicyAuditConsumed), "The opaque lookup audit consumes retention and redaction policy evidence."),
		productionReceiptNotificationActionDryRunResultOpaqueLookupCheck("opaque-lookup-guidance-consumed", productionAuthorizationPassBlocked(preview.OpaqueLookupGuidanceConsumed), "The opaque lookup audit consumes current dispatch guidance."),
		productionReceiptNotificationActionDryRunResultOpaqueLookupCheck("opaque-lookup-modeled-only", productionAuthorizationPassBlocked(preview.OpaqueLookupAuditRequired && preview.OpaqueLookupAuditModeled && preview.OpaqueLookupBoundaryReady && preview.OwnerManagedOpaqueLookupModeled && preview.OpaqueResultIDSupported && !preview.OpaqueLookupEnabled && !preview.OpaqueLookupPersisted), "Owner-managed opaque lookup is modeled without enabling lookup or persistence."),
		productionReceiptNotificationActionDryRunResultOpaqueLookupCheck("five-opaque-lookup-items-present", productionAuthorizationPassBlocked(preview.LookupItemCount == 5 && preview.RequiredLookupItemCount == 5 && preview.MissingLookupItemCount == 0), "Review, renew, open Compatibility Center, dismiss, and support-info lookup boundaries are modeled."),
		productionReceiptNotificationActionDryRunResultOpaqueLookupCheck("opaque-lookup-items-ready-lookup-disabled", productionAuthorizationPassBlocked(preview.ReadyLookupItemCount == 5 && productionReceiptNotificationActionDryRunResultOpaqueLookupItemsReady(preview.LookupItems)), "Every lookup item is ready while lookup routes and persistence remain disabled."),
		productionReceiptNotificationActionDryRunResultOpaqueLookupCheck("lookup-persistence-result-persistence-and-execution-disabled", productionAuthorizationPassBlocked(!preview.OpaqueLookupEnabled && !preview.OpaqueLookupPersisted && !preview.RetentionPolicyPersistenceEnabled && !preview.DispatchDryRunExecutionEnabled && !preview.DryRunResultPersistenceEnabled && !preview.ResultVisibilityPersistenceEnabled && !preview.RuntimeDiagnosticsPersisted && preview.EnabledLookupItemCount == 0 && preview.PersistedLookupItemCount == 0 && preview.PersistedResultItemCount == 0 && preview.ExecutedDryRunItemCount == 0), "Lookup persistence, result persistence, diagnostics persistence, policy persistence, and dry-run execution remain disabled."),
		productionReceiptNotificationActionDryRunResultOpaqueLookupCheck("request-creation-dispatch-actions-and-portal-disabled", productionAuthorizationPassBlocked(!preview.RequestObjectCreationEnabled && !preview.RequestObjectDispatchEnabled && !preview.RequestObjectPersistenceEnabled && !preview.RequestObjectsCreated && !preview.RequestObjectsDispatched && !preview.PortalRequestCreated && !preview.NotificationActionEnabled && preview.CreatedRequestObjectCount == 0 && preview.PortalRequestCreatedCount == 0), "Request creation, dispatch, persistence, notification actions, and Portal requests remain disabled."),
		productionReceiptNotificationActionDryRunResultOpaqueLookupCheck("receipt-notification-navigation-and-support-writes-disabled", productionAuthorizationPassBlocked(!preview.NotificationSent && !preview.NotificationDeliveryEnabled && !preview.CompatibilityCenterOpened && !preview.CompatibilityCenterPersisted && !preview.ReceiptWriterEnabled && !preview.ReceiptPersistenceEnabled && !preview.SupportBundleExported && !preview.SupportCaseCreated), "Notification, navigation, receipt, Compatibility Center, and support writes remain disabled."),
		productionReceiptNotificationActionDryRunResultOpaqueLookupCheck("production-and-host-boundary-closed", productionAuthorizationPassBlocked(!preview.SystemServiceStarted && !preview.SessionBusClaimed && !preview.ProductionBusClaimed && !preview.ProductionOwnerEnabled && !preview.ProductionActivationReady && !preview.ProductionReadiness && !preview.ProductionOwnershipReady && !preview.AdapterInvocationEnabled && !preview.BackendLaunchEnabled && !preview.BackendProcessStarted && !preview.FileContentRead && !preview.FilePathsExposed && !preview.NetworkRequired && !preview.HostRootModified && !preview.PrivilegedContainerRequired && !preview.StateRootPathExposed && !preview.RawCommandExposed && !preview.RawExecutableExposed && !preview.BackendDetailsExposed && preview.PathExposedLookupItemCount == 0 && productionReceiptNotificationActionDryRunResultOpaqueLookupItemsKeepHostClosed(preview.LookupItems)), "Production ownership, engine launch, unsafe data exposure, network, privilege, path exposure, internal details, and host mutation remain disabled."),
	}
}

func productionReceiptNotificationActionDryRunResultOpaqueLookupCheck(id string, status string, summary string) ProductionReceiptNotificationActionDryRunResultOpaqueLookupAuditCheck {
	return ProductionReceiptNotificationActionDryRunResultOpaqueLookupAuditCheck{ID: id, Status: status, Summary: summary}
}

func productionReceiptNotificationActionDryRunResultOpaqueLookupRetentionReady(source string) bool {
	return productionAuthorizationHasAll(source, []string{
		"production-receipt-notification-action-dry-run-result-retention-redaction-policy-audit-preview",
		"production-receipt-notification-action-dry-run-result-retention-redaction-policy-audit-ready-policy-only",
		"review-receipt-dry-run-result-retention-redaction-policy",
		"renew-receipt-dry-run-result-retention-redaction-policy",
		"open-compatibility-center-dry-run-result-retention-redaction-policy",
		"dismiss-receipt-dry-run-result-retention-redaction-policy",
		"support-info-dry-run-result-retention-redaction-policy",
		"RetentionPolicyPersistenceEnabled",
	})
}

func productionReceiptNotificationActionDryRunResultOpaqueLookupGuidanceReady(source string) bool {
	return productionAuthorizationHasAll(source, []string{"owner-managed opaque lookup", "opaque result identifiers", "side effects"})
}

func productionReceiptNotificationActionDryRunResultOpaqueLookupReadyCount(items []ProductionReceiptNotificationActionDryRunResultOpaqueLookupAuditItem) int {
	count := 0
	for _, item := range items {
		if item.EvidencePresent {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionDryRunResultOpaqueLookupMissingCount(items []ProductionReceiptNotificationActionDryRunResultOpaqueLookupAuditItem) int {
	count := 0
	for _, item := range items {
		if !item.EvidencePresent {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionDryRunResultOpaqueLookupEnabledCount(items []ProductionReceiptNotificationActionDryRunResultOpaqueLookupAuditItem) int {
	count := 0
	for _, item := range items {
		if item.LookupEnabled {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionDryRunResultOpaqueLookupPersistedCount(items []ProductionReceiptNotificationActionDryRunResultOpaqueLookupAuditItem) int {
	count := 0
	for _, item := range items {
		if item.LookupPersisted {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionDryRunResultOpaqueLookupPathExposedCount(items []ProductionReceiptNotificationActionDryRunResultOpaqueLookupAuditItem) int {
	count := 0
	for _, item := range items {
		if item.CallerStateRootRequired || item.StateRootPathExposed || item.FilePathsExposed || item.FileContentRead {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionDryRunResultOpaqueLookupPersistedResultCount(items []ProductionReceiptNotificationActionDryRunResultOpaqueLookupAuditItem) int {
	count := 0
	for _, item := range items {
		if item.DryRunResultPersisted || item.ResultVisibilityPersisted || item.RuntimeDiagnosticsPersisted {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionDryRunResultOpaqueLookupExecutedCount(items []ProductionReceiptNotificationActionDryRunResultOpaqueLookupAuditItem) int {
	count := 0
	for _, item := range items {
		if item.DispatchDryRunExecuted {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionDryRunResultOpaqueLookupCreatedCount(items []ProductionReceiptNotificationActionDryRunResultOpaqueLookupAuditItem) int {
	count := 0
	for _, item := range items {
		if item.RequestObjectCreated {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionDryRunResultOpaqueLookupPortalCount(items []ProductionReceiptNotificationActionDryRunResultOpaqueLookupAuditItem) int {
	count := 0
	for _, item := range items {
		if item.PortalRequestCreated {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionDryRunResultOpaqueLookupSideEffectCount(items []ProductionReceiptNotificationActionDryRunResultOpaqueLookupAuditItem) int {
	count := 0
	for _, item := range items {
		if item.LookupEnabled || item.LookupPersisted || item.ResultPersistenceAuthorized || item.RetentionEnforcementEnabled || item.RedactionEnforcementEnabled || item.DispatchDryRunExecuted || item.DryRunResultPersisted || item.ResultVisibilityPersisted || item.RuntimeDiagnosticsPersisted || item.RequestObjectCreated || item.RequestObjectDispatched || item.PortalRequestCreated || item.ReceiptAccepted || item.AuthorizationAccepted || item.CompatibilityCenterOpened || item.CompatibilityCenterPersisted || item.SupportBundleExported || item.SupportCaseCreated || item.HostRootModified || item.InternalDetailsExposed {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionDryRunResultOpaqueLookupItemIDs(items []ProductionReceiptNotificationActionDryRunResultOpaqueLookupAuditItem) []string {
	ids := make([]string, 0, len(items))
	for _, item := range items {
		ids = append(ids, item.ID)
	}
	return ids
}

func productionReceiptNotificationActionDryRunResultOpaqueLookupCheckIDs(checks []ProductionReceiptNotificationActionDryRunResultOpaqueLookupAuditCheck) []string {
	ids := make([]string, 0, len(checks))
	for _, check := range checks {
		ids = append(ids, check.ID)
	}
	return ids
}

func countProductionReceiptNotificationActionDryRunResultOpaqueLookupChecks(checks []ProductionReceiptNotificationActionDryRunResultOpaqueLookupAuditCheck) ProductionReceiptNotificationActionDryRunResultOpaqueLookupAuditCounts {
	counts := ProductionReceiptNotificationActionDryRunResultOpaqueLookupAuditCounts{Total: len(checks)}
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

func productionReceiptNotificationActionDryRunResultOpaqueLookupItemsReady(items []ProductionReceiptNotificationActionDryRunResultOpaqueLookupAuditItem) bool {
	if len(items) != 5 {
		return false
	}
	for _, item := range items {
		if !item.EvidencePresent || !item.OpaqueLookupModeled || !item.OwnerManagedLookup || !item.OpaqueResultIDSupported || item.LookupStatus != "opaque-lookup-modeled-lookup-disabled" {
			return false
		}
		if item.CallerStateRootRequired || item.StateRootPathExposed || item.FilePathsExposed || item.FileContentRead || item.LookupEnabled || item.LookupPersisted || item.OperatorApprovalPresent || item.ResultPersistenceAuthorized || item.RetentionEnforcementEnabled || item.RedactionEnforcementEnabled || item.DispatchDryRunExecuted || item.DryRunResultPersisted || item.ResultVisibilityPersisted || item.RuntimeDiagnosticsPersisted || item.RequestObjectCreated || item.RequestObjectDispatched || item.PortalRequestCreated || item.ReceiptAccepted || item.AuthorizationAccepted || item.CompatibilityCenterOpened || item.CompatibilityCenterPersisted || item.SupportBundleExported || item.SupportCaseCreated || item.ProductionReadiness || item.ProductionOwnershipReady || item.KDEPolicyOwner || !item.ReviewOnly || !item.SideEffectsDisabled {
			return false
		}
	}
	return true
}

func productionReceiptNotificationActionDryRunResultOpaqueLookupItemsKeepHostClosed(items []ProductionReceiptNotificationActionDryRunResultOpaqueLookupAuditItem) bool {
	for _, item := range items {
		if item.CallerStateRootRequired || item.StateRootPathExposed || item.FilePathsExposed || item.FileContentRead || item.HostRootModified || item.InternalDetailsExposed {
			return false
		}
	}
	return true
}
