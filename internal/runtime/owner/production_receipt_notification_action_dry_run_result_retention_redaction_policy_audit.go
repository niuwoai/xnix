package owner

type ProductionReceiptNotificationActionDryRunResultRetentionRedactionPolicyAuditPreview struct {
	Version                               string                                                                              `json:"version"`
	SchemaVersion                         string                                                                              `json:"schema_version"`
	RequestType                           string                                                                              `json:"request_type"`
	AuditType                             string                                                                              `json:"audit_type"`
	Source                                string                                                                              `json:"source"`
	AuditDecision                         string                                                                              `json:"audit_decision"`
	ReceiptSchema                         string                                                                              `json:"receipt_schema"`
	OpaqueReceiptID                       string                                                                              `json:"opaque_receipt_id"`
	RetentionRedactionPolicyAuditRequired bool                                                                                `json:"retention_redaction_policy_audit_required"`
	RetentionRedactionPolicyAuditModeled  bool                                                                                `json:"retention_redaction_policy_audit_modeled"`
	PersistenceAuthorizationAuditConsumed bool                                                                                `json:"persistence_authorization_audit_consumed"`
	RetentionRedactionGuidanceConsumed    bool                                                                                `json:"retention_redaction_guidance_consumed"`
	RetentionRedactionPolicyReady         bool                                                                                `json:"retention_redaction_policy_ready"`
	RetentionWindowModeled                bool                                                                                `json:"retention_window_modeled"`
	RedactionRulesModeled                 bool                                                                                `json:"redaction_rules_modeled"`
	ResultPersistenceAuthorized           bool                                                                                `json:"result_persistence_authorized"`
	RetentionEnforcementEnabled           bool                                                                                `json:"retention_enforcement_enabled"`
	RedactionEnforcementEnabled           bool                                                                                `json:"redaction_enforcement_enabled"`
	DispatchDryRunExecuted                bool                                                                                `json:"dispatch_dry_run_executed"`
	DryRunResultPersisted                 bool                                                                                `json:"dry_run_result_persisted"`
	ResultVisibilityPersisted             bool                                                                                `json:"result_visibility_persisted"`
	DispatchAuthorizationGranted          bool                                                                                `json:"dispatch_authorization_granted"`
	OperatorPersistenceApprovalRequired   bool                                                                                `json:"operator_persistence_approval_required"`
	OperatorPersistenceApprovalPresent    bool                                                                                `json:"operator_persistence_approval_present"`
	CallerStateRootRequired               bool                                                                                `json:"caller_state_root_required"`
	ReceiptPresent                        bool                                                                                `json:"receipt_present"`
	ReceiptAccepted                       bool                                                                                `json:"receipt_accepted"`
	AuthorizationAccepted                 bool                                                                                `json:"authorization_accepted"`
	ProductionReadiness                   bool                                                                                `json:"production_readiness"`
	ProductionOwnershipReady              bool                                                                                `json:"production_ownership_ready"`
	PolicyItemCount                       int                                                                                 `json:"policy_item_count"`
	RequiredPolicyItemCount               int                                                                                 `json:"required_policy_item_count"`
	ReadyPolicyItemCount                  int                                                                                 `json:"ready_policy_item_count"`
	MissingPolicyItemCount                int                                                                                 `json:"missing_policy_item_count"`
	RetentionEnabledItemCount             int                                                                                 `json:"retention_enabled_item_count"`
	RedactionWriteItemCount               int                                                                                 `json:"redaction_write_item_count"`
	PersistedResultItemCount              int                                                                                 `json:"persisted_result_item_count"`
	ExecutedDryRunItemCount               int                                                                                 `json:"executed_dry_run_item_count"`
	CreatedRequestObjectCount             int                                                                                 `json:"created_request_object_count"`
	PortalRequestCreatedCount             int                                                                                 `json:"portal_request_created_count"`
	SideEffectPolicyItemCount             int                                                                                 `json:"side_effect_policy_item_count"`
	PolicyItems                           []ProductionReceiptNotificationActionDryRunResultRetentionRedactionPolicyAuditItem  `json:"policy_items"`
	PolicyItemIDs                         []string                                                                            `json:"policy_item_ids"`
	RequiredBeforePolicyEnforcement       []string                                                                            `json:"required_before_policy_enforcement"`
	Checks                                []ProductionReceiptNotificationActionDryRunResultRetentionRedactionPolicyAuditCheck `json:"checks"`
	CheckIDs                              []string                                                                            `json:"check_ids"`
	Counts                                ProductionReceiptNotificationActionDryRunResultRetentionRedactionPolicyAuditCounts  `json:"counts"`
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
	RuntimeDiagnosticsPersisted           bool                                                                                `json:"runtime_diagnostics_persisted"`
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
	FileContentRead                       bool                                                                                `json:"file_content_read"`
	FilePathsExposed                      bool                                                                                `json:"file_paths_exposed"`
	NetworkRequired                       bool                                                                                `json:"network_required"`
	HostRootModified                      bool                                                                                `json:"host_root_modified"`
	PrivilegedContainerRequired           bool                                                                                `json:"privileged_container_required"`
	StateRootPathExposed                  bool                                                                                `json:"state_root_path_exposed"`
	RawCommandExposed                     bool                                                                                `json:"raw_command_exposed"`
	RawExecutableExposed                  bool                                                                                `json:"raw_executable_exposed"`
	BackendDetailsExposed                 bool                                                                                `json:"backend_details_exposed"`
	BlockedActions                        []string                                                                            `json:"blocked_actions"`
	NextRequirements                      []string                                                                            `json:"next_requirements"`
	DesktopSafeSummary                    string                                                                              `json:"desktop_safe_summary"`
}

type ProductionReceiptNotificationActionDryRunResultRetentionRedactionPolicyAuditItem struct {
	ID                           string `json:"id"`
	ActionKind                   string `json:"action_kind"`
	PersistenceAuthorizationKind string `json:"persistence_authorization_kind"`
	RetentionPolicyKind          string `json:"retention_policy_kind"`
	RedactionPolicyKind          string `json:"redaction_policy_kind"`
	RetentionWindow              string `json:"retention_window"`
	RequiredEvidence             string `json:"required_evidence"`
	EvidencePresent              bool   `json:"evidence_present"`
	RetentionPolicyModeled       bool   `json:"retention_policy_modeled"`
	RedactionPolicyModeled       bool   `json:"redaction_policy_modeled"`
	RedactedForKDE               bool   `json:"redacted_for_kde"`
	RuntimeDiagnosticsModeled    bool   `json:"runtime_diagnostics_modeled"`
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
	RequestObjectCreated         bool   `json:"request_object_created"`
	RequestObjectDispatched      bool   `json:"request_object_dispatched"`
	PortalRequestCreated         bool   `json:"portal_request_created"`
	ReceiptAccepted              bool   `json:"receipt_accepted"`
	AuthorizationAccepted        bool   `json:"authorization_accepted"`
	CompatibilityCenterOpened    bool   `json:"compatibility_center_opened"`
	CompatibilityCenterPersisted bool   `json:"compatibility_center_persisted"`
	RuntimeDiagnosticsPersisted  bool   `json:"runtime_diagnostics_persisted"`
	SupportBundleExported        bool   `json:"support_bundle_exported"`
	SupportCaseCreated           bool   `json:"support_case_created"`
	ProductionReadiness          bool   `json:"production_readiness"`
	ProductionOwnershipReady     bool   `json:"production_ownership_ready"`
	SideEffectsDisabled          bool   `json:"side_effects_disabled"`
	HostRootModified             bool   `json:"host_root_modified"`
	InternalDetailsExposed       bool   `json:"internal_details_exposed"`
	PolicyStatus                 string `json:"policy_status"`
	NextRequirement              string `json:"next_requirement"`
}

type ProductionReceiptNotificationActionDryRunResultRetentionRedactionPolicyAuditCheck struct {
	ID      string `json:"id"`
	Status  string `json:"status"`
	Summary string `json:"summary"`
}

type ProductionReceiptNotificationActionDryRunResultRetentionRedactionPolicyAuditCounts struct {
	Total   int `json:"total"`
	Passed  int `json:"passed"`
	Pending int `json:"pending"`
	Blocked int `json:"blocked"`
}

func NewProductionReceiptNotificationActionDryRunResultRetentionRedactionPolicyAuditPreview(root string) (ProductionReceiptNotificationActionDryRunResultRetentionRedactionPolicyAuditPreview, error) {
	version, err := readProductionDBusGateReviewVersion(root)
	if err != nil {
		return ProductionReceiptNotificationActionDryRunResultRetentionRedactionPolicyAuditPreview{}, err
	}
	sources := productionReceiptNotificationActionDryRunResultRetentionRedactionPolicyAuditSources(root)
	items := productionReceiptNotificationActionDryRunResultRetentionRedactionPolicyItems(sources)
	authorizationReady := productionReceiptNotificationActionDryRunResultRetentionRedactionPolicyAuthorizationReady(sources.PersistenceAuthorizationAudit)
	guidanceReady := productionReceiptNotificationActionDryRunResultRetentionRedactionPolicyGuidanceReady(sources.DispatchSheet)
	policyReady := authorizationReady && guidanceReady && productionReceiptNotificationActionDryRunResultRetentionRedactionPolicyItemsReady(items)
	preview := ProductionReceiptNotificationActionDryRunResultRetentionRedactionPolicyAuditPreview{
		Version:                               version,
		SchemaVersion:                         "xnix.runtime.production_receipt_notification_action_dry_run_result_retention_redaction_policy_audit.v1",
		RequestType:                           "production-receipt-notification-action-dry-run-result-retention-redaction-policy-audit-preview",
		AuditType:                             "receipt-notification-action-dry-run-result-retention-redaction-policy-audit",
		Source:                                "production-receipt-notification-action-dry-run-result-persistence-authorization-audit-preview+claude-code-current-dispatch-picks",
		AuditDecision:                         "production-receipt-notification-action-dry-run-result-retention-redaction-policy-audit-blocked",
		ReceiptSchema:                         "xnix.runtime.production_dbus_human_authorization_receipt.v1",
		OpaqueReceiptID:                       ProductionDBusHumanAuthorizationReceiptID,
		RetentionRedactionPolicyAuditRequired: true,
		RetentionRedactionPolicyAuditModeled:  true,
		PersistenceAuthorizationAuditConsumed: authorizationReady,
		RetentionRedactionGuidanceConsumed:    guidanceReady,
		RetentionRedactionPolicyReady:         policyReady,
		RetentionWindowModeled:                policyReady,
		RedactionRulesModeled:                 policyReady,
		ResultPersistenceAuthorized:           false,
		RetentionEnforcementEnabled:           false,
		RedactionEnforcementEnabled:           false,
		DispatchDryRunExecuted:                false,
		DryRunResultPersisted:                 false,
		ResultVisibilityPersisted:             false,
		DispatchAuthorizationGranted:          false,
		OperatorPersistenceApprovalRequired:   true,
		OperatorPersistenceApprovalPresent:    false,
		CallerStateRootRequired:               false,
		ReceiptPresent:                        false,
		ReceiptAccepted:                       false,
		AuthorizationAccepted:                 false,
		ProductionReadiness:                   false,
		ProductionOwnershipReady:              false,
		PolicyItemCount:                       len(items),
		RequiredPolicyItemCount:               5,
		ReadyPolicyItemCount:                  productionReceiptNotificationActionDryRunResultRetentionRedactionPolicyReadyCount(items),
		MissingPolicyItemCount:                productionReceiptNotificationActionDryRunResultRetentionRedactionPolicyMissingCount(items),
		PolicyItems:                           items,
		PolicyItemIDs:                         productionReceiptNotificationActionDryRunResultRetentionRedactionPolicyItemIDs(items),
		RequiredBeforePolicyEnforcement: []string{
			"production-receipt-notification-action-dry-run-result-persistence-authorization-audit-preview",
			"accepted opaque persistence authorization receipt",
			"separate retention enforcement implementation",
			"separate redaction enforcement implementation",
			"owner-managed opaque result lookup",
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
		RuntimeDiagnosticsPersisted:        false,
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
		FileContentRead:                    false,
		FilePathsExposed:                   false,
		NetworkRequired:                    false,
		HostRootModified:                   false,
		PrivilegedContainerRequired:        false,
		StateRootPathExposed:               false,
		RawCommandExposed:                  false,
		RawExecutableExposed:               false,
		BackendDetailsExposed:              false,
		BlockedActions: []string{
			"treat retention and redaction policy audit as permission to enforce retention or redaction",
			"persist policy records, dry-run results, result visibility, or Runtime diagnostics from this audit",
			"execute dispatch dry runs, create Portal requests, write receipts, or export support data from this audit",
			"claim production D-Bus ownership, start services, launch compatibility engines, expose paths, or mutate host root",
		},
		NextRequirements: []string{
			"Add owner-managed opaque dry-run result lookup before KDE can consume persisted summaries.",
			"Implement retention enforcement only after persistence authorization, policy review, and lookup boundaries exist.",
			"Implement redaction enforcement only after KDE-safe and diagnostics-safe vocabularies are reviewed.",
			"Keep policy storage, dry-run result storage, and Runtime diagnostics persistence disabled until a separate writer is reviewed.",
		},
		DesktopSafeSummary: "The production receipt notification action dry-run result retention redaction policy audit models retention windows and redaction rules for future stored dry-run result summaries, but it enforces no retention, writes no policies, persists no results, executes no dry runs, creates no request objects, starts no services, launches no engines, and mutates no host state.",
	}
	preview.RetentionEnabledItemCount = productionReceiptNotificationActionDryRunResultRetentionRedactionPolicyRetentionEnabledCount(items)
	preview.RedactionWriteItemCount = productionReceiptNotificationActionDryRunResultRetentionRedactionPolicyRedactionWriteCount(items)
	preview.PersistedResultItemCount = productionReceiptNotificationActionDryRunResultRetentionRedactionPolicyPersistedCount(items)
	preview.ExecutedDryRunItemCount = productionReceiptNotificationActionDryRunResultRetentionRedactionPolicyExecutedCount(items)
	preview.CreatedRequestObjectCount = productionReceiptNotificationActionDryRunResultRetentionRedactionPolicyCreatedCount(items)
	preview.PortalRequestCreatedCount = productionReceiptNotificationActionDryRunResultRetentionRedactionPolicyPortalCount(items)
	preview.SideEffectPolicyItemCount = productionReceiptNotificationActionDryRunResultRetentionRedactionPolicySideEffectCount(items)
	checks := productionReceiptNotificationActionDryRunResultRetentionRedactionPolicyAuditChecks(preview)
	preview.Checks = checks
	preview.CheckIDs = productionReceiptNotificationActionDryRunResultRetentionRedactionPolicyCheckIDs(checks)
	preview.Counts = countProductionReceiptNotificationActionDryRunResultRetentionRedactionPolicyChecks(checks)
	if preview.Counts.Blocked == 0 {
		preview.AuditDecision = "production-receipt-notification-action-dry-run-result-retention-redaction-policy-audit-ready-policy-only"
	}
	if err := validateNoBackendTerms(preview, "production receipt notification action dry-run result retention redaction policy audit preview"); err != nil {
		return ProductionReceiptNotificationActionDryRunResultRetentionRedactionPolicyAuditPreview{}, err
	}
	return preview, nil
}

type productionReceiptNotificationActionDryRunResultRetentionRedactionPolicyAuditSourceSet struct {
	PersistenceAuthorizationAudit string
	DispatchSheet                 string
}

func productionReceiptNotificationActionDryRunResultRetentionRedactionPolicyAuditSources(root string) productionReceiptNotificationActionDryRunResultRetentionRedactionPolicyAuditSourceSet {
	return productionReceiptNotificationActionDryRunResultRetentionRedactionPolicyAuditSourceSet{
		PersistenceAuthorizationAudit: productionAuthorizationReadSources(root, []string{"internal/runtime/owner/production_receipt_notification_action_dry_run_result_persistence_authorization_audit.go"}),
		DispatchSheet:                 productionAuthorizationReadSources(root, []string{"docs/claude-code-current-dispatch-picks.md"}),
	}
}

func productionReceiptNotificationActionDryRunResultRetentionRedactionPolicyItems(sources productionReceiptNotificationActionDryRunResultRetentionRedactionPolicyAuditSourceSet) []ProductionReceiptNotificationActionDryRunResultRetentionRedactionPolicyAuditItem {
	combined := sources.PersistenceAuthorizationAudit + sources.DispatchSheet
	return []ProductionReceiptNotificationActionDryRunResultRetentionRedactionPolicyAuditItem{
		productionReceiptNotificationActionDryRunResultRetentionRedactionPolicyItem("review-receipt-dry-run-result-retention-redaction-policy", "review", "review-result-persistence-authorization", "review-result-retention-policy", "review-result-redaction-policy", combined, []string{"review-receipt-dry-run-result-persistence-authorization", "retention", "redaction"}, "define owner-managed review result lookup before enforcement"),
		productionReceiptNotificationActionDryRunResultRetentionRedactionPolicyItem("renew-receipt-dry-run-result-retention-redaction-policy", "renew", "renewal-result-persistence-authorization", "renewal-result-retention-policy", "renewal-result-redaction-policy", combined, []string{"renew-receipt-dry-run-result-persistence-authorization", "retention", "redaction"}, "define owner-managed renewal result lookup before enforcement"),
		productionReceiptNotificationActionDryRunResultRetentionRedactionPolicyItem("open-compatibility-center-dry-run-result-retention-redaction-policy", "open-compatibility-center", "navigation-result-persistence-authorization", "navigation-result-retention-policy", "navigation-result-redaction-policy", combined, []string{"open-compatibility-center-dry-run-result-persistence-authorization", "retention", "redaction"}, "define owner-managed navigation result lookup before enforcement"),
		productionReceiptNotificationActionDryRunResultRetentionRedactionPolicyItem("dismiss-receipt-dry-run-result-retention-redaction-policy", "dismiss", "dismissal-result-persistence-authorization", "dismissal-result-retention-policy", "dismissal-result-redaction-policy", combined, []string{"dismiss-receipt-dry-run-result-persistence-authorization", "retention", "redaction"}, "define owner-managed dismissal result lookup before enforcement"),
		productionReceiptNotificationActionDryRunResultRetentionRedactionPolicyItem("support-info-dry-run-result-retention-redaction-policy", "support-info", "support-info-result-persistence-authorization", "support-info-result-retention-policy", "support-info-result-redaction-policy", combined, []string{"support-info-dry-run-result-persistence-authorization", "retention", "redaction"}, "define owner-managed support-info result lookup before enforcement"),
	}
}

func productionReceiptNotificationActionDryRunResultRetentionRedactionPolicyItem(id string, actionKind string, authorizationKind string, retentionKind string, redactionKind string, source string, tokens []string, nextRequirement string) ProductionReceiptNotificationActionDryRunResultRetentionRedactionPolicyAuditItem {
	ready := productionAuthorizationHasAll(source, tokens)
	status := "missing-retention-redaction-policy-evidence"
	if ready {
		status = "retention-redaction-policy-modeled-enforcement-disabled"
	}
	return ProductionReceiptNotificationActionDryRunResultRetentionRedactionPolicyAuditItem{
		ID:                           id,
		ActionKind:                   actionKind,
		PersistenceAuthorizationKind: authorizationKind,
		RetentionPolicyKind:          retentionKind,
		RedactionPolicyKind:          redactionKind,
		RetentionWindow:              "redacted-summary-short-retention-policy-not-enforced",
		RequiredEvidence:             "future stored dry-run result summary requires retention and redaction policy",
		EvidencePresent:              ready,
		RetentionPolicyModeled:       ready,
		RedactionPolicyModeled:       ready,
		RedactedForKDE:               ready,
		RuntimeDiagnosticsModeled:    ready,
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
		RequestObjectCreated:         false,
		RequestObjectDispatched:      false,
		PortalRequestCreated:         false,
		ReceiptAccepted:              false,
		AuthorizationAccepted:        false,
		CompatibilityCenterOpened:    false,
		CompatibilityCenterPersisted: false,
		RuntimeDiagnosticsPersisted:  false,
		SupportBundleExported:        false,
		SupportCaseCreated:           false,
		ProductionReadiness:          false,
		ProductionOwnershipReady:     false,
		SideEffectsDisabled:          true,
		HostRootModified:             false,
		InternalDetailsExposed:       false,
		PolicyStatus:                 status,
		NextRequirement:              nextRequirement,
	}
}

func productionReceiptNotificationActionDryRunResultRetentionRedactionPolicyAuditChecks(preview ProductionReceiptNotificationActionDryRunResultRetentionRedactionPolicyAuditPreview) []ProductionReceiptNotificationActionDryRunResultRetentionRedactionPolicyAuditCheck {
	return []ProductionReceiptNotificationActionDryRunResultRetentionRedactionPolicyAuditCheck{
		productionReceiptNotificationActionDryRunResultRetentionRedactionPolicyCheck("persistence-authorization-audit-consumed", productionAuthorizationPassBlocked(preview.PersistenceAuthorizationAuditConsumed), "The retention redaction policy audit consumes persistence authorization evidence."),
		productionReceiptNotificationActionDryRunResultRetentionRedactionPolicyCheck("retention-redaction-guidance-consumed", productionAuthorizationPassBlocked(preview.RetentionRedactionGuidanceConsumed), "The retention redaction policy audit consumes current dispatch guidance."),
		productionReceiptNotificationActionDryRunResultRetentionRedactionPolicyCheck("retention-redaction-policy-modeled-only", productionAuthorizationPassBlocked(preview.RetentionRedactionPolicyAuditRequired && preview.RetentionRedactionPolicyAuditModeled && preview.RetentionRedactionPolicyReady && preview.RetentionWindowModeled && preview.RedactionRulesModeled && !preview.ResultPersistenceAuthorized && !preview.RetentionEnforcementEnabled && !preview.RedactionEnforcementEnabled && !preview.DryRunResultPersisted && !preview.ResultVisibilityPersisted), "Retention and redaction policy is modeled without enforcement, persistence authorization, or writes."),
		productionReceiptNotificationActionDryRunResultRetentionRedactionPolicyCheck("five-retention-redaction-policy-items-present", productionAuthorizationPassBlocked(preview.PolicyItemCount == 5 && preview.RequiredPolicyItemCount == 5 && preview.MissingPolicyItemCount == 0), "Review, renew, open Compatibility Center, dismiss, and support-info policies are modeled."),
		productionReceiptNotificationActionDryRunResultRetentionRedactionPolicyCheck("policy-items-ready-enforcement-disabled", productionAuthorizationPassBlocked(preview.ReadyPolicyItemCount == 5 && productionReceiptNotificationActionDryRunResultRetentionRedactionPolicyItemsReady(preview.PolicyItems)), "Every policy item is ready while enforcement and persistence remain disabled."),
		productionReceiptNotificationActionDryRunResultRetentionRedactionPolicyCheck("enforcement-persistence-and-execution-disabled", productionAuthorizationPassBlocked(!preview.RetentionEnforcementEnabled && !preview.RedactionEnforcementEnabled && !preview.RetentionPolicyPersistenceEnabled && !preview.DispatchDryRunExecutionEnabled && !preview.DryRunResultPersistenceEnabled && !preview.ResultVisibilityPersistenceEnabled && preview.RetentionEnabledItemCount == 0 && preview.RedactionWriteItemCount == 0 && preview.PersistedResultItemCount == 0 && preview.ExecutedDryRunItemCount == 0), "Retention enforcement, redaction writes, policy persistence, result persistence, and dry-run execution remain disabled."),
		productionReceiptNotificationActionDryRunResultRetentionRedactionPolicyCheck("request-creation-dispatch-actions-and-portal-disabled", productionAuthorizationPassBlocked(!preview.RequestObjectCreationEnabled && !preview.RequestObjectDispatchEnabled && !preview.RequestObjectPersistenceEnabled && !preview.RequestObjectsCreated && !preview.RequestObjectsDispatched && !preview.PortalRequestCreated && !preview.NotificationActionEnabled && preview.CreatedRequestObjectCount == 0 && preview.PortalRequestCreatedCount == 0), "Request creation, dispatch, persistence, notification actions, and Portal requests remain disabled."),
		productionReceiptNotificationActionDryRunResultRetentionRedactionPolicyCheck("receipt-notification-navigation-and-support-writes-disabled", productionAuthorizationPassBlocked(!preview.NotificationSent && !preview.NotificationDeliveryEnabled && !preview.CompatibilityCenterOpened && !preview.CompatibilityCenterPersisted && !preview.ReceiptWriterEnabled && !preview.ReceiptPersistenceEnabled && !preview.SupportBundleExported && !preview.SupportCaseCreated), "Notification, navigation, receipt, Compatibility Center, and support writes remain disabled."),
		productionReceiptNotificationActionDryRunResultRetentionRedactionPolicyCheck("production-and-host-boundary-closed", productionAuthorizationPassBlocked(!preview.SystemServiceStarted && !preview.SessionBusClaimed && !preview.ProductionBusClaimed && !preview.ProductionOwnerEnabled && !preview.ProductionActivationReady && !preview.ProductionReadiness && !preview.ProductionOwnershipReady && !preview.AdapterInvocationEnabled && !preview.BackendLaunchEnabled && !preview.BackendProcessStarted && !preview.FileContentRead && !preview.FilePathsExposed && !preview.NetworkRequired && !preview.HostRootModified && !preview.PrivilegedContainerRequired && !preview.StateRootPathExposed && !preview.RawCommandExposed && !preview.RawExecutableExposed && !preview.BackendDetailsExposed && productionReceiptNotificationActionDryRunResultRetentionRedactionPolicyItemsKeepHostClosed(preview.PolicyItems)), "Production ownership, backend launch, unsafe data exposure, network, privilege, path exposure, internal details, and host mutation remain disabled."),
	}
}

func productionReceiptNotificationActionDryRunResultRetentionRedactionPolicyCheck(id string, status string, summary string) ProductionReceiptNotificationActionDryRunResultRetentionRedactionPolicyAuditCheck {
	return ProductionReceiptNotificationActionDryRunResultRetentionRedactionPolicyAuditCheck{ID: id, Status: status, Summary: summary}
}

func productionReceiptNotificationActionDryRunResultRetentionRedactionPolicyAuthorizationReady(source string) bool {
	return productionAuthorizationHasAll(source, []string{
		"production-receipt-notification-action-dry-run-result-persistence-authorization-audit-preview",
		"production-receipt-notification-action-dry-run-result-persistence-authorization-audit-ready-persistence-disabled",
		"review-receipt-dry-run-result-persistence-authorization",
		"renew-receipt-dry-run-result-persistence-authorization",
		"open-compatibility-center-dry-run-result-persistence-authorization",
		"dismiss-receipt-dry-run-result-persistence-authorization",
		"support-info-dry-run-result-persistence-authorization",
		"ResultPersistenceAuthorized",
	})
}

func productionReceiptNotificationActionDryRunResultRetentionRedactionPolicyGuidanceReady(source string) bool {
	return productionAuthorizationHasAll(source, []string{"retention", "redaction", "persistence", "side effects"})
}

func productionReceiptNotificationActionDryRunResultRetentionRedactionPolicyReadyCount(items []ProductionReceiptNotificationActionDryRunResultRetentionRedactionPolicyAuditItem) int {
	count := 0
	for _, item := range items {
		if item.EvidencePresent {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionDryRunResultRetentionRedactionPolicyMissingCount(items []ProductionReceiptNotificationActionDryRunResultRetentionRedactionPolicyAuditItem) int {
	count := 0
	for _, item := range items {
		if !item.EvidencePresent {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionDryRunResultRetentionRedactionPolicyRetentionEnabledCount(items []ProductionReceiptNotificationActionDryRunResultRetentionRedactionPolicyAuditItem) int {
	count := 0
	for _, item := range items {
		if item.RetentionEnforcementEnabled {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionDryRunResultRetentionRedactionPolicyRedactionWriteCount(items []ProductionReceiptNotificationActionDryRunResultRetentionRedactionPolicyAuditItem) int {
	count := 0
	for _, item := range items {
		if item.RedactionEnforcementEnabled {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionDryRunResultRetentionRedactionPolicyPersistedCount(items []ProductionReceiptNotificationActionDryRunResultRetentionRedactionPolicyAuditItem) int {
	count := 0
	for _, item := range items {
		if item.DryRunResultPersisted || item.ResultVisibilityPersisted || item.RuntimeDiagnosticsPersisted {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionDryRunResultRetentionRedactionPolicyExecutedCount(items []ProductionReceiptNotificationActionDryRunResultRetentionRedactionPolicyAuditItem) int {
	count := 0
	for _, item := range items {
		if item.DispatchDryRunExecuted {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionDryRunResultRetentionRedactionPolicyCreatedCount(items []ProductionReceiptNotificationActionDryRunResultRetentionRedactionPolicyAuditItem) int {
	count := 0
	for _, item := range items {
		if item.RequestObjectCreated {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionDryRunResultRetentionRedactionPolicyPortalCount(items []ProductionReceiptNotificationActionDryRunResultRetentionRedactionPolicyAuditItem) int {
	count := 0
	for _, item := range items {
		if item.PortalRequestCreated {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionDryRunResultRetentionRedactionPolicySideEffectCount(items []ProductionReceiptNotificationActionDryRunResultRetentionRedactionPolicyAuditItem) int {
	count := 0
	for _, item := range items {
		if item.ResultPersistenceAuthorized || item.RetentionEnforcementEnabled || item.RedactionEnforcementEnabled || item.DispatchDryRunExecuted || item.DryRunResultPersisted || item.ResultVisibilityPersisted || item.RuntimeDiagnosticsPersisted || item.RequestObjectCreated || item.RequestObjectDispatched || item.PortalRequestCreated || item.ReceiptAccepted || item.AuthorizationAccepted || item.CompatibilityCenterOpened || item.CompatibilityCenterPersisted || item.SupportBundleExported || item.SupportCaseCreated || item.HostRootModified || item.InternalDetailsExposed {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionDryRunResultRetentionRedactionPolicyItemIDs(items []ProductionReceiptNotificationActionDryRunResultRetentionRedactionPolicyAuditItem) []string {
	ids := make([]string, 0, len(items))
	for _, item := range items {
		ids = append(ids, item.ID)
	}
	return ids
}

func productionReceiptNotificationActionDryRunResultRetentionRedactionPolicyCheckIDs(checks []ProductionReceiptNotificationActionDryRunResultRetentionRedactionPolicyAuditCheck) []string {
	ids := make([]string, 0, len(checks))
	for _, check := range checks {
		ids = append(ids, check.ID)
	}
	return ids
}

func countProductionReceiptNotificationActionDryRunResultRetentionRedactionPolicyChecks(checks []ProductionReceiptNotificationActionDryRunResultRetentionRedactionPolicyAuditCheck) ProductionReceiptNotificationActionDryRunResultRetentionRedactionPolicyAuditCounts {
	counts := ProductionReceiptNotificationActionDryRunResultRetentionRedactionPolicyAuditCounts{Total: len(checks)}
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

func productionReceiptNotificationActionDryRunResultRetentionRedactionPolicyItemsReady(items []ProductionReceiptNotificationActionDryRunResultRetentionRedactionPolicyAuditItem) bool {
	if len(items) != 5 {
		return false
	}
	for _, item := range items {
		if !item.EvidencePresent || !item.RetentionPolicyModeled || !item.RedactionPolicyModeled || !item.RedactedForKDE || !item.RuntimeDiagnosticsModeled || item.PolicyStatus != "retention-redaction-policy-modeled-enforcement-disabled" {
			return false
		}
		if item.OperatorApprovalPresent || item.ResultPersistenceAuthorized || item.RetentionEnforcementEnabled || item.RedactionEnforcementEnabled || item.DispatchDryRunExecuted || item.DryRunResultPersisted || item.ResultVisibilityPersisted || item.RuntimeDiagnosticsPersisted || item.RequestObjectCreated || item.RequestObjectDispatched || item.PortalRequestCreated || item.ReceiptAccepted || item.AuthorizationAccepted || item.CompatibilityCenterOpened || item.CompatibilityCenterPersisted || item.SupportBundleExported || item.SupportCaseCreated || item.ProductionReadiness || item.ProductionOwnershipReady || item.KDEPolicyOwner || !item.ReviewOnly || !item.SideEffectsDisabled {
			return false
		}
	}
	return true
}

func productionReceiptNotificationActionDryRunResultRetentionRedactionPolicyItemsKeepHostClosed(items []ProductionReceiptNotificationActionDryRunResultRetentionRedactionPolicyAuditItem) bool {
	for _, item := range items {
		if item.HostRootModified || item.InternalDetailsExposed {
			return false
		}
	}
	return true
}
