package owner

type ProductionReceiptNotificationActionDryRunResultPersistenceAuthorizationAuditPreview struct {
	Version                                  string                                                                              `json:"version"`
	SchemaVersion                            string                                                                              `json:"schema_version"`
	RequestType                              string                                                                              `json:"request_type"`
	AuditType                                string                                                                              `json:"audit_type"`
	Source                                   string                                                                              `json:"source"`
	AuditDecision                            string                                                                              `json:"audit_decision"`
	ReceiptSchema                            string                                                                              `json:"receipt_schema"`
	OpaqueReceiptID                          string                                                                              `json:"opaque_receipt_id"`
	PersistenceAuthorizationAuditRequired    bool                                                                                `json:"persistence_authorization_audit_required"`
	PersistenceAuthorizationAuditModeled     bool                                                                                `json:"persistence_authorization_audit_modeled"`
	ResultVisibilityAuditConsumed            bool                                                                                `json:"result_visibility_audit_consumed"`
	PersistenceAuthorizationGuidanceConsumed bool                                                                                `json:"persistence_authorization_guidance_consumed"`
	PersistenceAuthorizationReady            bool                                                                                `json:"persistence_authorization_ready"`
	ResultPersistenceAuthorized              bool                                                                                `json:"result_persistence_authorized"`
	DispatchDryRunExecuted                   bool                                                                                `json:"dispatch_dry_run_executed"`
	DryRunResultPersisted                    bool                                                                                `json:"dry_run_result_persisted"`
	ResultVisibilityPersisted                bool                                                                                `json:"result_visibility_persisted"`
	DispatchAuthorizationGranted             bool                                                                                `json:"dispatch_authorization_granted"`
	OperatorPersistenceApprovalRequired      bool                                                                                `json:"operator_persistence_approval_required"`
	OperatorPersistenceApprovalPresent       bool                                                                                `json:"operator_persistence_approval_present"`
	CallerStateRootRequired                  bool                                                                                `json:"caller_state_root_required"`
	ReceiptPresent                           bool                                                                                `json:"receipt_present"`
	ReceiptAccepted                          bool                                                                                `json:"receipt_accepted"`
	AuthorizationAccepted                    bool                                                                                `json:"authorization_accepted"`
	ProductionReadiness                      bool                                                                                `json:"production_readiness"`
	ProductionOwnershipReady                 bool                                                                                `json:"production_ownership_ready"`
	AuthorizationItemCount                   int                                                                                 `json:"authorization_item_count"`
	RequiredAuthorizationItemCount           int                                                                                 `json:"required_authorization_item_count"`
	ReadyAuthorizationItemCount              int                                                                                 `json:"ready_authorization_item_count"`
	MissingAuthorizationItemCount            int                                                                                 `json:"missing_authorization_item_count"`
	GrantedPersistenceItemCount              int                                                                                 `json:"granted_persistence_item_count"`
	PersistedResultItemCount                 int                                                                                 `json:"persisted_result_item_count"`
	ExecutedDryRunItemCount                  int                                                                                 `json:"executed_dry_run_item_count"`
	CreatedRequestObjectCount                int                                                                                 `json:"created_request_object_count"`
	PortalRequestCreatedCount                int                                                                                 `json:"portal_request_created_count"`
	SideEffectAuthorizationItemCount         int                                                                                 `json:"side_effect_authorization_item_count"`
	AuthorizationItems                       []ProductionReceiptNotificationActionDryRunResultPersistenceAuthorizationAuditItem  `json:"authorization_items"`
	AuthorizationItemIDs                     []string                                                                            `json:"authorization_item_ids"`
	RequiredBeforePersistence                []string                                                                            `json:"required_before_persistence"`
	Checks                                   []ProductionReceiptNotificationActionDryRunResultPersistenceAuthorizationAuditCheck `json:"checks"`
	CheckIDs                                 []string                                                                            `json:"check_ids"`
	Counts                                   ProductionReceiptNotificationActionDryRunResultPersistenceAuthorizationAuditCounts  `json:"counts"`
	RuntimeOwned                             bool                                                                                `json:"runtime_owned"`
	GoRuntimeBacked                          bool                                                                                `json:"go_runtime_backed"`
	KDEPolicyOwner                           bool                                                                                `json:"kde_policy_owner"`
	OfficialDesktopOnly                      bool                                                                                `json:"official_desktop_only"`
	PlasmaForkRequired                       bool                                                                                `json:"plasma_fork_required"`
	PlasmaSourceModified                     bool                                                                                `json:"plasma_source_modified"`
	SystemServiceStarted                     bool                                                                                `json:"system_service_started"`
	SessionBusClaimed                        bool                                                                                `json:"session_bus_claimed"`
	ProductionBusClaimed                     bool                                                                                `json:"production_bus_claimed"`
	ProductionOwnerEnabled                   bool                                                                                `json:"production_owner_enabled"`
	ProductionActivationReady                bool                                                                                `json:"production_activation_ready"`
	WriteMethodsEnabled                      bool                                                                                `json:"write_methods_enabled"`
	RuntimeWritesEnabled                     bool                                                                                `json:"runtime_writes_enabled"`
	RequestObjectCreationEnabled             bool                                                                                `json:"request_object_creation_enabled"`
	RequestObjectDispatchEnabled             bool                                                                                `json:"request_object_dispatch_enabled"`
	RequestObjectPersistenceEnabled          bool                                                                                `json:"request_object_persistence_enabled"`
	DispatchAuthorizationPersisted           bool                                                                                `json:"dispatch_authorization_persisted"`
	DispatchDryRunExecutionEnabled           bool                                                                                `json:"dispatch_dry_run_execution_enabled"`
	DryRunResultPersistenceEnabled           bool                                                                                `json:"dry_run_result_persistence_enabled"`
	ResultVisibilityPersistenceEnabled       bool                                                                                `json:"result_visibility_persistence_enabled"`
	PortalRequestCreated                     bool                                                                                `json:"portal_request_created"`
	RequestObjectsCreated                    bool                                                                                `json:"request_objects_created"`
	RequestObjectsDispatched                 bool                                                                                `json:"request_objects_dispatched"`
	NotificationSent                         bool                                                                                `json:"notification_sent"`
	NotificationDeliveryEnabled              bool                                                                                `json:"notification_delivery_enabled"`
	NotificationActionEnabled                bool                                                                                `json:"notification_action_enabled"`
	ReviewActionEnabled                      bool                                                                                `json:"review_action_enabled"`
	RenewActionEnabled                       bool                                                                                `json:"renew_action_enabled"`
	OpenCompatibilityCenterEnabled           bool                                                                                `json:"open_compatibility_center_enabled"`
	DismissActionEnabled                     bool                                                                                `json:"dismiss_action_enabled"`
	SupportInfoActionEnabled                 bool                                                                                `json:"support_info_action_enabled"`
	CompatibilityCenterOpened                bool                                                                                `json:"compatibility_center_opened"`
	CompatibilityCenterPersisted             bool                                                                                `json:"compatibility_center_persisted"`
	RuntimeDiagnosticsPersisted              bool                                                                                `json:"runtime_diagnostics_persisted"`
	ReceiptWriterEnabled                     bool                                                                                `json:"receipt_writer_enabled"`
	ReceiptPersistenceEnabled                bool                                                                                `json:"receipt_persistence_enabled"`
	ReceiptLookupWritesEnabled               bool                                                                                `json:"receipt_lookup_writes_enabled"`
	ReceiptReplayEnabled                     bool                                                                                `json:"receipt_replay_enabled"`
	ReceiptExpiryWriteEnabled                bool                                                                                `json:"receipt_expiry_write_enabled"`
	ReceiptRevocationWriteEnabled            bool                                                                                `json:"receipt_revocation_write_enabled"`
	DesktopFilesWritten                      bool                                                                                `json:"desktop_files_written"`
	MIMEAppsWritten                          bool                                                                                `json:"mimeapps_written"`
	ShellConfigurationWritten                bool                                                                                `json:"shell_configuration_written"`
	SettingsPersisted                        bool                                                                                `json:"settings_persisted"`
	KRunnerIndexPersisted                    bool                                                                                `json:"krunner_index_persisted"`
	TaskManagerEntryActive                   bool                                                                                `json:"task_manager_entry_active"`
	KWinRuleApplied                          bool                                                                                `json:"kwin_rule_applied"`
	LiveTrayBridgeEnabled                    bool                                                                                `json:"live_tray_bridge_enabled"`
	TrayBridgePersisted                      bool                                                                                `json:"tray_bridge_persisted"`
	AdapterInvocationEnabled                 bool                                                                                `json:"adapter_invocation_enabled"`
	BackendLaunchEnabled                     bool                                                                                `json:"backend_launch_enabled"`
	BackendProcessStarted                    bool                                                                                `json:"backend_process_started"`
	SupportBundleExported                    bool                                                                                `json:"support_bundle_exported"`
	SupportCaseCreated                       bool                                                                                `json:"support_case_created"`
	SnapshotRestoreExecuted                  bool                                                                                `json:"snapshot_restore_executed"`
	StateCleanupExecuted                     bool                                                                                `json:"state_cleanup_executed"`
	FileContentRead                          bool                                                                                `json:"file_content_read"`
	FilePathsExposed                         bool                                                                                `json:"file_paths_exposed"`
	NetworkRequired                          bool                                                                                `json:"network_required"`
	HostRootModified                         bool                                                                                `json:"host_root_modified"`
	PrivilegedContainerRequired              bool                                                                                `json:"privileged_container_required"`
	StateRootPathExposed                     bool                                                                                `json:"state_root_path_exposed"`
	RawCommandExposed                        bool                                                                                `json:"raw_command_exposed"`
	RawExecutableExposed                     bool                                                                                `json:"raw_executable_exposed"`
	BackendDetailsExposed                    bool                                                                                `json:"backend_details_exposed"`
	BlockedActions                           []string                                                                            `json:"blocked_actions"`
	NextRequirements                         []string                                                                            `json:"next_requirements"`
	DesktopSafeSummary                       string                                                                              `json:"desktop_safe_summary"`
}

type ProductionReceiptNotificationActionDryRunResultPersistenceAuthorizationAuditItem struct {
	ID                              string `json:"id"`
	ActionKind                      string `json:"action_kind"`
	ResultVisibilityKind            string `json:"result_visibility_kind"`
	PersistenceAuthorizationKind    string `json:"persistence_authorization_kind"`
	PersistenceTargetKind           string `json:"persistence_target_kind"`
	RequiredEvidence                string `json:"required_evidence"`
	EvidencePresent                 bool   `json:"evidence_present"`
	PersistenceAuthorizationModeled bool   `json:"persistence_authorization_modeled"`
	RedactedForKDE                  bool   `json:"redacted_for_kde"`
	RuntimeDiagnosticsModeled       bool   `json:"runtime_diagnostics_modeled"`
	UserVisible                     bool   `json:"user_visible"`
	ReviewOnly                      bool   `json:"review_only"`
	RuntimeOwned                    bool   `json:"runtime_owned"`
	GoRuntimeBacked                 bool   `json:"go_runtime_backed"`
	KDEPolicyOwner                  bool   `json:"kde_policy_owner"`
	OperatorApprovalRequired        bool   `json:"operator_approval_required"`
	OperatorApprovalPresent         bool   `json:"operator_approval_present"`
	ResultPersistenceAuthorized     bool   `json:"result_persistence_authorized"`
	DispatchAuthorizationGranted    bool   `json:"dispatch_authorization_granted"`
	DispatchDryRunExecuted          bool   `json:"dispatch_dry_run_executed"`
	DryRunResultPersisted           bool   `json:"dry_run_result_persisted"`
	ResultVisibilityPersisted       bool   `json:"result_visibility_persisted"`
	RequestObjectCreated            bool   `json:"request_object_created"`
	RequestObjectDispatched         bool   `json:"request_object_dispatched"`
	RequestObjectPersisted          bool   `json:"request_object_persisted"`
	PortalRequestCreated            bool   `json:"portal_request_created"`
	NavigationRequested             bool   `json:"navigation_requested"`
	NotificationActionEnabled       bool   `json:"notification_action_enabled"`
	ActionEnabled                   bool   `json:"action_enabled"`
	ReceiptAccepted                 bool   `json:"receipt_accepted"`
	AuthorizationAccepted           bool   `json:"authorization_accepted"`
	ReceiptWriterEnabled            bool   `json:"receipt_writer_enabled"`
	ReceiptPersistenceEnabled       bool   `json:"receipt_persistence_enabled"`
	CompatibilityCenterOpened       bool   `json:"compatibility_center_opened"`
	CompatibilityCenterPersisted    bool   `json:"compatibility_center_persisted"`
	RuntimeDiagnosticsPersisted     bool   `json:"runtime_diagnostics_persisted"`
	SupportBundleExported           bool   `json:"support_bundle_exported"`
	SupportCaseCreated              bool   `json:"support_case_created"`
	ProductionReadiness             bool   `json:"production_readiness"`
	ProductionOwnershipReady        bool   `json:"production_ownership_ready"`
	SideEffectsDisabled             bool   `json:"side_effects_disabled"`
	HostRootModified                bool   `json:"host_root_modified"`
	InternalDetailsExposed          bool   `json:"internal_details_exposed"`
	PersistenceAuthorizationStatus  string `json:"persistence_authorization_status"`
	NextRequirement                 string `json:"next_requirement"`
}

type ProductionReceiptNotificationActionDryRunResultPersistenceAuthorizationAuditCheck struct {
	ID      string `json:"id"`
	Status  string `json:"status"`
	Summary string `json:"summary"`
}

type ProductionReceiptNotificationActionDryRunResultPersistenceAuthorizationAuditCounts struct {
	Total   int `json:"total"`
	Passed  int `json:"passed"`
	Pending int `json:"pending"`
	Blocked int `json:"blocked"`
}

func NewProductionReceiptNotificationActionDryRunResultPersistenceAuthorizationAuditPreview(root string) (ProductionReceiptNotificationActionDryRunResultPersistenceAuthorizationAuditPreview, error) {
	version, err := readProductionDBusGateReviewVersion(root)
	if err != nil {
		return ProductionReceiptNotificationActionDryRunResultPersistenceAuthorizationAuditPreview{}, err
	}
	sources := productionReceiptNotificationActionDryRunResultPersistenceAuthorizationAuditSources(root)
	items := productionReceiptNotificationActionDryRunResultPersistenceAuthorizationItems(sources)
	visibilityReady := productionReceiptNotificationActionDryRunResultPersistenceAuthorizationVisibilityReady(sources.ResultVisibilityAudit)
	guidanceReady := productionReceiptNotificationActionDryRunResultPersistenceAuthorizationGuidanceReady(sources.DispatchSheet)
	authorizationReady := visibilityReady && guidanceReady && productionReceiptNotificationActionDryRunResultPersistenceAuthorizationItemsReady(items)
	preview := ProductionReceiptNotificationActionDryRunResultPersistenceAuthorizationAuditPreview{
		Version:                                  version,
		SchemaVersion:                            "xnix.runtime.production_receipt_notification_action_dry_run_result_persistence_authorization_audit.v1",
		RequestType:                              "production-receipt-notification-action-dry-run-result-persistence-authorization-audit-preview",
		AuditType:                                "receipt-notification-action-dry-run-result-persistence-authorization-audit",
		Source:                                   "production-receipt-notification-action-dry-run-result-visibility-audit-preview+claude-code-current-dispatch-picks",
		AuditDecision:                            "production-receipt-notification-action-dry-run-result-persistence-authorization-audit-blocked",
		ReceiptSchema:                            "xnix.runtime.production_dbus_human_authorization_receipt.v1",
		OpaqueReceiptID:                          ProductionDBusHumanAuthorizationReceiptID,
		PersistenceAuthorizationAuditRequired:    true,
		PersistenceAuthorizationAuditModeled:     true,
		ResultVisibilityAuditConsumed:            visibilityReady,
		PersistenceAuthorizationGuidanceConsumed: guidanceReady,
		PersistenceAuthorizationReady:            authorizationReady,
		ResultPersistenceAuthorized:              false,
		DispatchDryRunExecuted:                   false,
		DryRunResultPersisted:                    false,
		ResultVisibilityPersisted:                false,
		DispatchAuthorizationGranted:             false,
		OperatorPersistenceApprovalRequired:      true,
		OperatorPersistenceApprovalPresent:       false,
		CallerStateRootRequired:                  false,
		ReceiptPresent:                           false,
		ReceiptAccepted:                          false,
		AuthorizationAccepted:                    false,
		ProductionReadiness:                      false,
		ProductionOwnershipReady:                 false,
		AuthorizationItemCount:                   len(items),
		RequiredAuthorizationItemCount:           5,
		ReadyAuthorizationItemCount:              productionReceiptNotificationActionDryRunResultPersistenceAuthorizationReadyCount(items),
		MissingAuthorizationItemCount:            productionReceiptNotificationActionDryRunResultPersistenceAuthorizationMissingCount(items),
		AuthorizationItems:                       items,
		AuthorizationItemIDs:                     productionReceiptNotificationActionDryRunResultPersistenceAuthorizationItemIDs(items),
		RequiredBeforePersistence: []string{
			"production-receipt-notification-action-dry-run-result-visibility-audit-preview",
			"explicit operator persistence approval",
			"accepted opaque authorization receipt",
			"separate persistence writer implementation",
			"separate retention and redaction policy",
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
			"treat persistence authorization audit as permission to persist dry-run results",
			"execute dispatch dry runs, persist result visibility, or persist Runtime diagnostics from this audit",
			"create, persist, or dispatch notification action request objects from this audit",
			"create Portal requests, navigate KDE surfaces, write receipts, emit notifications, or export support data from this audit",
			"claim production D-Bus ownership, start services, launch compatibility engines, expose paths, or mutate host root",
		},
		NextRequirements: []string{
			"Define retention and redaction policy before any dry-run result summary can be written.",
			"Add a separate persistence writer implementation after authorization and retention are reviewed.",
			"Keep KDE and Runtime diagnostics read-only until a persisted result summary has an owner-managed opaque lookup.",
			"Require explicit operator persistence approval and an accepted opaque authorization receipt before persistence can be granted.",
		},
		DesktopSafeSummary: "The production receipt notification action dry-run result persistence authorization audit models the authorization boundary for future storage of redacted dry-run result summaries, but it grants no persistence authorization, executes no dry runs, persists no results, creates no request objects, sends no Portal requests, opens no surfaces, writes no receipts, starts no services, launches no engines, and mutates no host state.",
	}
	preview.GrantedPersistenceItemCount = productionReceiptNotificationActionDryRunResultPersistenceAuthorizationGrantedCount(items)
	preview.PersistedResultItemCount = productionReceiptNotificationActionDryRunResultPersistenceAuthorizationPersistedCount(items)
	preview.ExecutedDryRunItemCount = productionReceiptNotificationActionDryRunResultPersistenceAuthorizationExecutedCount(items)
	preview.CreatedRequestObjectCount = productionReceiptNotificationActionDryRunResultPersistenceAuthorizationCreatedCount(items)
	preview.PortalRequestCreatedCount = productionReceiptNotificationActionDryRunResultPersistenceAuthorizationPortalCount(items)
	preview.SideEffectAuthorizationItemCount = productionReceiptNotificationActionDryRunResultPersistenceAuthorizationSideEffectCount(items)
	checks := productionReceiptNotificationActionDryRunResultPersistenceAuthorizationAuditChecks(preview)
	preview.Checks = checks
	preview.CheckIDs = productionReceiptNotificationActionDryRunResultPersistenceAuthorizationCheckIDs(checks)
	preview.Counts = countProductionReceiptNotificationActionDryRunResultPersistenceAuthorizationChecks(checks)
	if preview.Counts.Blocked == 0 {
		preview.AuditDecision = "production-receipt-notification-action-dry-run-result-persistence-authorization-audit-ready-persistence-disabled"
	}
	if err := validateNoBackendTerms(preview, "production receipt notification action dry-run result persistence authorization audit preview"); err != nil {
		return ProductionReceiptNotificationActionDryRunResultPersistenceAuthorizationAuditPreview{}, err
	}
	return preview, nil
}

type productionReceiptNotificationActionDryRunResultPersistenceAuthorizationAuditSourceSet struct {
	ResultVisibilityAudit string
	DispatchSheet         string
}

func productionReceiptNotificationActionDryRunResultPersistenceAuthorizationAuditSources(root string) productionReceiptNotificationActionDryRunResultPersistenceAuthorizationAuditSourceSet {
	return productionReceiptNotificationActionDryRunResultPersistenceAuthorizationAuditSourceSet{
		ResultVisibilityAudit: productionAuthorizationReadSources(root, []string{"internal/runtime/owner/production_receipt_notification_action_dry_run_result_visibility_audit.go"}),
		DispatchSheet:         productionAuthorizationReadSources(root, []string{"docs/claude-code-current-dispatch-picks.md"}),
	}
}

func productionReceiptNotificationActionDryRunResultPersistenceAuthorizationItems(sources productionReceiptNotificationActionDryRunResultPersistenceAuthorizationAuditSourceSet) []ProductionReceiptNotificationActionDryRunResultPersistenceAuthorizationAuditItem {
	combined := sources.ResultVisibilityAudit + sources.DispatchSheet
	return []ProductionReceiptNotificationActionDryRunResultPersistenceAuthorizationAuditItem{
		productionReceiptNotificationActionDryRunResultPersistenceAuthorizationItem("review-receipt-dry-run-result-persistence-authorization", "review", "review-dry-run-result-visibility", "review-result-persistence-authorization", "redacted-review-result-summary", "future review dry-run result persistence requires authorization", combined, []string{"review-receipt-dry-run-result-visibility", "dry-run result visibility", "persistence authorization"}, "define review retention and redaction policy before persistence"),
		productionReceiptNotificationActionDryRunResultPersistenceAuthorizationItem("renew-receipt-dry-run-result-persistence-authorization", "renew", "renewal-dry-run-result-visibility", "renewal-result-persistence-authorization", "redacted-renewal-result-summary", "future renewal dry-run result persistence requires authorization", combined, []string{"renew-receipt-dry-run-result-visibility", "dry-run result visibility", "persistence authorization"}, "define renewal retention and redaction policy before persistence"),
		productionReceiptNotificationActionDryRunResultPersistenceAuthorizationItem("open-compatibility-center-dry-run-result-persistence-authorization", "open-compatibility-center", "navigation-dry-run-result-visibility", "navigation-result-persistence-authorization", "redacted-navigation-result-summary", "future navigation dry-run result persistence requires authorization", combined, []string{"open-compatibility-center-dry-run-result-visibility", "dry-run result visibility", "persistence authorization"}, "define navigation retention and redaction policy before persistence"),
		productionReceiptNotificationActionDryRunResultPersistenceAuthorizationItem("dismiss-receipt-dry-run-result-persistence-authorization", "dismiss", "dismissal-dry-run-result-visibility", "dismissal-result-persistence-authorization", "redacted-dismissal-result-summary", "future dismissal dry-run result persistence requires authorization", combined, []string{"dismiss-receipt-dry-run-result-visibility", "dry-run result visibility", "persistence authorization"}, "define dismissal retention and redaction policy before persistence"),
		productionReceiptNotificationActionDryRunResultPersistenceAuthorizationItem("support-info-dry-run-result-persistence-authorization", "support-info", "support-info-dry-run-result-visibility", "support-info-result-persistence-authorization", "redacted-support-info-result-summary", "future support-info dry-run result persistence requires authorization", combined, []string{"support-info-dry-run-result-visibility", "dry-run result visibility", "persistence authorization"}, "define support-info retention and redaction policy before persistence"),
	}
}

func productionReceiptNotificationActionDryRunResultPersistenceAuthorizationItem(id string, actionKind string, visibilityKind string, authorizationKind string, targetKind string, evidence string, source string, tokens []string, nextRequirement string) ProductionReceiptNotificationActionDryRunResultPersistenceAuthorizationAuditItem {
	ready := productionAuthorizationHasAll(source, tokens)
	status := "missing-dry-run-result-persistence-authorization-evidence"
	if ready {
		status = "dry-run-result-persistence-authorization-modeled-persistence-disabled"
	}
	return ProductionReceiptNotificationActionDryRunResultPersistenceAuthorizationAuditItem{
		ID:                              id,
		ActionKind:                      actionKind,
		ResultVisibilityKind:            visibilityKind,
		PersistenceAuthorizationKind:    authorizationKind,
		PersistenceTargetKind:           targetKind,
		RequiredEvidence:                evidence,
		EvidencePresent:                 ready,
		PersistenceAuthorizationModeled: ready,
		RedactedForKDE:                  ready,
		RuntimeDiagnosticsModeled:       ready,
		UserVisible:                     ready,
		ReviewOnly:                      true,
		RuntimeOwned:                    true,
		GoRuntimeBacked:                 true,
		KDEPolicyOwner:                  false,
		OperatorApprovalRequired:        true,
		OperatorApprovalPresent:         false,
		ResultPersistenceAuthorized:     false,
		DispatchAuthorizationGranted:    false,
		DispatchDryRunExecuted:          false,
		DryRunResultPersisted:           false,
		ResultVisibilityPersisted:       false,
		RequestObjectCreated:            false,
		RequestObjectDispatched:         false,
		RequestObjectPersisted:          false,
		PortalRequestCreated:            false,
		NavigationRequested:             false,
		NotificationActionEnabled:       false,
		ActionEnabled:                   false,
		ReceiptAccepted:                 false,
		AuthorizationAccepted:           false,
		ReceiptWriterEnabled:            false,
		ReceiptPersistenceEnabled:       false,
		CompatibilityCenterOpened:       false,
		CompatibilityCenterPersisted:    false,
		RuntimeDiagnosticsPersisted:     false,
		SupportBundleExported:           false,
		SupportCaseCreated:              false,
		ProductionReadiness:             false,
		ProductionOwnershipReady:        false,
		SideEffectsDisabled:             true,
		HostRootModified:                false,
		InternalDetailsExposed:          false,
		PersistenceAuthorizationStatus:  status,
		NextRequirement:                 nextRequirement,
	}
}

func productionReceiptNotificationActionDryRunResultPersistenceAuthorizationAuditChecks(preview ProductionReceiptNotificationActionDryRunResultPersistenceAuthorizationAuditPreview) []ProductionReceiptNotificationActionDryRunResultPersistenceAuthorizationAuditCheck {
	return []ProductionReceiptNotificationActionDryRunResultPersistenceAuthorizationAuditCheck{
		productionReceiptNotificationActionDryRunResultPersistenceAuthorizationCheck("result-visibility-audit-consumed", productionAuthorizationPassBlocked(preview.ResultVisibilityAuditConsumed), "The persistence authorization audit consumes dry-run result visibility evidence."),
		productionReceiptNotificationActionDryRunResultPersistenceAuthorizationCheck("persistence-authorization-guidance-consumed", productionAuthorizationPassBlocked(preview.PersistenceAuthorizationGuidanceConsumed), "The persistence authorization audit consumes current dispatch guidance for persistence authorization."),
		productionReceiptNotificationActionDryRunResultPersistenceAuthorizationCheck("persistence-authorization-modeled-only", productionAuthorizationPassBlocked(preview.PersistenceAuthorizationAuditRequired && preview.PersistenceAuthorizationAuditModeled && preview.PersistenceAuthorizationReady && preview.OperatorPersistenceApprovalRequired && !preview.OperatorPersistenceApprovalPresent && !preview.ResultPersistenceAuthorized && !preview.DispatchDryRunExecuted && !preview.DryRunResultPersisted && !preview.ResultVisibilityPersisted && !preview.ReceiptPresent && !preview.ReceiptAccepted && !preview.AuthorizationAccepted), "Persistence authorization is modeled without approval, grant, dry-run execution, result persistence, visibility persistence, accepted receipt, or production readiness."),
		productionReceiptNotificationActionDryRunResultPersistenceAuthorizationCheck("five-persistence-authorizations-present", productionAuthorizationPassBlocked(preview.AuthorizationItemCount == 5 && preview.RequiredAuthorizationItemCount == 5 && preview.MissingAuthorizationItemCount == 0), "Review, renew, open Compatibility Center, dismiss, and support-info persistence authorizations are modeled."),
		productionReceiptNotificationActionDryRunResultPersistenceAuthorizationCheck("persistence-authorizations-ready-persistence-disabled", productionAuthorizationPassBlocked(preview.ReadyAuthorizationItemCount == 5 && productionReceiptNotificationActionDryRunResultPersistenceAuthorizationItemsReady(preview.AuthorizationItems)), "Every persistence authorization item is ready while persistence, execution, request creation, dispatch, Portal requests, and navigation remain disabled."),
		productionReceiptNotificationActionDryRunResultPersistenceAuthorizationCheck("dry-run-execution-result-and-visibility-persistence-disabled", productionAuthorizationPassBlocked(!preview.DispatchDryRunExecutionEnabled && !preview.DispatchDryRunExecuted && !preview.DryRunResultPersisted && !preview.ResultVisibilityPersisted && !preview.DryRunResultPersistenceEnabled && !preview.ResultVisibilityPersistenceEnabled && !preview.RuntimeDiagnosticsPersisted && preview.GrantedPersistenceItemCount == 0 && preview.PersistedResultItemCount == 0 && preview.ExecutedDryRunItemCount == 0), "Dry-run execution, persistence grants, result persistence, visibility persistence, Runtime diagnostics persistence, and per-item execution remain disabled."),
		productionReceiptNotificationActionDryRunResultPersistenceAuthorizationCheck("request-creation-dispatch-actions-and-portal-disabled", productionAuthorizationPassBlocked(!preview.RequestObjectCreationEnabled && !preview.RequestObjectDispatchEnabled && !preview.RequestObjectPersistenceEnabled && !preview.RequestObjectsCreated && !preview.RequestObjectsDispatched && !preview.PortalRequestCreated && !preview.NotificationActionEnabled && !preview.ReviewActionEnabled && !preview.RenewActionEnabled && !preview.OpenCompatibilityCenterEnabled && !preview.DismissActionEnabled && !preview.SupportInfoActionEnabled && preview.CreatedRequestObjectCount == 0 && preview.PortalRequestCreatedCount == 0), "Request creation, dispatch, persistence, notification actions, and Portal requests remain disabled."),
		productionReceiptNotificationActionDryRunResultPersistenceAuthorizationCheck("receipt-notification-navigation-and-support-writes-disabled", productionAuthorizationPassBlocked(!preview.NotificationSent && !preview.NotificationDeliveryEnabled && !preview.CompatibilityCenterOpened && !preview.CompatibilityCenterPersisted && !preview.ReceiptWriterEnabled && !preview.ReceiptPersistenceEnabled && !preview.ReceiptLookupWritesEnabled && !preview.ReceiptReplayEnabled && !preview.ReceiptExpiryWriteEnabled && !preview.ReceiptRevocationWriteEnabled && !preview.SupportBundleExported && !preview.SupportCaseCreated), "Notification, navigation, receipt, Compatibility Center, and support write side effects remain disabled."),
		productionReceiptNotificationActionDryRunResultPersistenceAuthorizationCheck("production-and-host-boundary-closed", productionAuthorizationPassBlocked(!preview.SystemServiceStarted && !preview.SessionBusClaimed && !preview.ProductionBusClaimed && !preview.ProductionOwnerEnabled && !preview.ProductionActivationReady && !preview.ProductionReadiness && !preview.ProductionOwnershipReady && !preview.AdapterInvocationEnabled && !preview.BackendLaunchEnabled && !preview.BackendProcessStarted && !preview.FileContentRead && !preview.FilePathsExposed && !preview.NetworkRequired && !preview.HostRootModified && !preview.PrivilegedContainerRequired && !preview.StateRootPathExposed && !preview.RawCommandExposed && !preview.RawExecutableExposed && !preview.BackendDetailsExposed && productionReceiptNotificationActionDryRunResultPersistenceAuthorizationItemsKeepHostClosed(preview.AuthorizationItems)), "Production ownership, backend launch, unsafe data exposure, network, privilege, path exposure, internal details, and host mutation remain disabled."),
	}
}

func productionReceiptNotificationActionDryRunResultPersistenceAuthorizationCheck(id string, status string, summary string) ProductionReceiptNotificationActionDryRunResultPersistenceAuthorizationAuditCheck {
	return ProductionReceiptNotificationActionDryRunResultPersistenceAuthorizationAuditCheck{ID: id, Status: status, Summary: summary}
}

func productionReceiptNotificationActionDryRunResultPersistenceAuthorizationVisibilityReady(source string) bool {
	return productionAuthorizationHasAll(source, []string{
		"production-receipt-notification-action-dry-run-result-visibility-audit-preview",
		"receipt-notification-action-dry-run-result-visibility-audit",
		"production-receipt-notification-action-dry-run-result-visibility-audit-ready-visibility-only",
		"review-receipt-dry-run-result-visibility",
		"renew-receipt-dry-run-result-visibility",
		"open-compatibility-center-dry-run-result-visibility",
		"dismiss-receipt-dry-run-result-visibility",
		"support-info-dry-run-result-visibility",
		"DryRunResultPersisted",
		"ResultVisibilityPersisted",
	})
}

func productionReceiptNotificationActionDryRunResultPersistenceAuthorizationGuidanceReady(source string) bool {
	return productionAuthorizationHasAll(source, []string{
		"dry-run result persistence authorization",
		"persistence",
		"dry-run execution",
		"side effects",
	})
}

func productionReceiptNotificationActionDryRunResultPersistenceAuthorizationReadyCount(items []ProductionReceiptNotificationActionDryRunResultPersistenceAuthorizationAuditItem) int {
	count := 0
	for _, item := range items {
		if item.EvidencePresent {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionDryRunResultPersistenceAuthorizationMissingCount(items []ProductionReceiptNotificationActionDryRunResultPersistenceAuthorizationAuditItem) int {
	count := 0
	for _, item := range items {
		if !item.EvidencePresent {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionDryRunResultPersistenceAuthorizationGrantedCount(items []ProductionReceiptNotificationActionDryRunResultPersistenceAuthorizationAuditItem) int {
	count := 0
	for _, item := range items {
		if item.ResultPersistenceAuthorized {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionDryRunResultPersistenceAuthorizationPersistedCount(items []ProductionReceiptNotificationActionDryRunResultPersistenceAuthorizationAuditItem) int {
	count := 0
	for _, item := range items {
		if item.DryRunResultPersisted || item.ResultVisibilityPersisted || item.RuntimeDiagnosticsPersisted {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionDryRunResultPersistenceAuthorizationExecutedCount(items []ProductionReceiptNotificationActionDryRunResultPersistenceAuthorizationAuditItem) int {
	count := 0
	for _, item := range items {
		if item.DispatchDryRunExecuted {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionDryRunResultPersistenceAuthorizationCreatedCount(items []ProductionReceiptNotificationActionDryRunResultPersistenceAuthorizationAuditItem) int {
	count := 0
	for _, item := range items {
		if item.RequestObjectCreated {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionDryRunResultPersistenceAuthorizationPortalCount(items []ProductionReceiptNotificationActionDryRunResultPersistenceAuthorizationAuditItem) int {
	count := 0
	for _, item := range items {
		if item.PortalRequestCreated {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionDryRunResultPersistenceAuthorizationSideEffectCount(items []ProductionReceiptNotificationActionDryRunResultPersistenceAuthorizationAuditItem) int {
	count := 0
	for _, item := range items {
		if item.ResultPersistenceAuthorized || item.DispatchAuthorizationGranted || item.DispatchDryRunExecuted || item.DryRunResultPersisted || item.ResultVisibilityPersisted || item.RuntimeDiagnosticsPersisted || item.RequestObjectCreated || item.RequestObjectDispatched || item.RequestObjectPersisted || item.PortalRequestCreated || item.NavigationRequested || item.NotificationActionEnabled || item.ActionEnabled || item.ReceiptAccepted || item.AuthorizationAccepted || item.ReceiptWriterEnabled || item.ReceiptPersistenceEnabled || item.CompatibilityCenterOpened || item.CompatibilityCenterPersisted || item.SupportBundleExported || item.SupportCaseCreated || item.HostRootModified || item.InternalDetailsExposed {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionDryRunResultPersistenceAuthorizationItemIDs(items []ProductionReceiptNotificationActionDryRunResultPersistenceAuthorizationAuditItem) []string {
	ids := make([]string, 0, len(items))
	for _, item := range items {
		ids = append(ids, item.ID)
	}
	return ids
}

func productionReceiptNotificationActionDryRunResultPersistenceAuthorizationCheckIDs(checks []ProductionReceiptNotificationActionDryRunResultPersistenceAuthorizationAuditCheck) []string {
	ids := make([]string, 0, len(checks))
	for _, check := range checks {
		ids = append(ids, check.ID)
	}
	return ids
}

func countProductionReceiptNotificationActionDryRunResultPersistenceAuthorizationChecks(checks []ProductionReceiptNotificationActionDryRunResultPersistenceAuthorizationAuditCheck) ProductionReceiptNotificationActionDryRunResultPersistenceAuthorizationAuditCounts {
	counts := ProductionReceiptNotificationActionDryRunResultPersistenceAuthorizationAuditCounts{Total: len(checks)}
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

func productionReceiptNotificationActionDryRunResultPersistenceAuthorizationItemsReady(items []ProductionReceiptNotificationActionDryRunResultPersistenceAuthorizationAuditItem) bool {
	if len(items) != 5 {
		return false
	}
	for _, item := range items {
		if !item.EvidencePresent || !item.PersistenceAuthorizationModeled || !item.RedactedForKDE || !item.RuntimeDiagnosticsModeled || !item.UserVisible || item.PersistenceAuthorizationStatus != "dry-run-result-persistence-authorization-modeled-persistence-disabled" {
			return false
		}
		if item.OperatorApprovalPresent || item.ResultPersistenceAuthorized || item.DispatchAuthorizationGranted || item.DispatchDryRunExecuted || item.DryRunResultPersisted || item.ResultVisibilityPersisted || item.RuntimeDiagnosticsPersisted || item.RequestObjectCreated || item.RequestObjectDispatched || item.RequestObjectPersisted || item.PortalRequestCreated || item.NavigationRequested || item.NotificationActionEnabled || item.ActionEnabled || item.ReceiptAccepted || item.AuthorizationAccepted || item.ReceiptWriterEnabled || item.ReceiptPersistenceEnabled || item.CompatibilityCenterOpened || item.CompatibilityCenterPersisted || item.SupportBundleExported || item.SupportCaseCreated || item.ProductionReadiness || item.ProductionOwnershipReady || item.KDEPolicyOwner || !item.ReviewOnly || !item.SideEffectsDisabled {
			return false
		}
	}
	return true
}

func productionReceiptNotificationActionDryRunResultPersistenceAuthorizationItemsKeepHostClosed(items []ProductionReceiptNotificationActionDryRunResultPersistenceAuthorizationAuditItem) bool {
	for _, item := range items {
		if item.HostRootModified || item.InternalDetailsExposed {
			return false
		}
	}
	return true
}
