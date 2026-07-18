package owner

type ProductionReceiptNotificationActionDryRunResultVisibilityAuditPreview struct {
	Version                             string                                                                `json:"version"`
	SchemaVersion                       string                                                                `json:"schema_version"`
	RequestType                         string                                                                `json:"request_type"`
	AuditType                           string                                                                `json:"audit_type"`
	Source                              string                                                                `json:"source"`
	AuditDecision                       string                                                                `json:"audit_decision"`
	ReceiptSchema                       string                                                                `json:"receipt_schema"`
	OpaqueReceiptID                     string                                                                `json:"opaque_receipt_id"`
	ResultVisibilityAuditRequired       bool                                                                  `json:"result_visibility_audit_required"`
	ResultVisibilityAuditModeled        bool                                                                  `json:"result_visibility_audit_modeled"`
	DispatchDryRunAuditConsumed         bool                                                                  `json:"dispatch_dry_run_audit_consumed"`
	ResultVisibilityGuidanceConsumed    bool                                                                  `json:"result_visibility_guidance_consumed"`
	ResultVisibilityPlanReady           bool                                                                  `json:"result_visibility_plan_ready"`
	RedactedKDEVisibilityModeled        bool                                                                  `json:"redacted_kde_visibility_modeled"`
	RuntimeDiagnosticsVisibilityModeled bool                                                                  `json:"runtime_diagnostics_visibility_modeled"`
	DispatchDryRunExecuted              bool                                                                  `json:"dispatch_dry_run_executed"`
	DryRunResultPersisted               bool                                                                  `json:"dry_run_result_persisted"`
	DispatchAuthorizationGranted        bool                                                                  `json:"dispatch_authorization_granted"`
	OperatorDispatchApprovalRequired    bool                                                                  `json:"operator_dispatch_approval_required"`
	OperatorDispatchApprovalPresent     bool                                                                  `json:"operator_dispatch_approval_present"`
	CallerStateRootRequired             bool                                                                  `json:"caller_state_root_required"`
	ReceiptPresent                      bool                                                                  `json:"receipt_present"`
	ReceiptAccepted                     bool                                                                  `json:"receipt_accepted"`
	AuthorizationAccepted               bool                                                                  `json:"authorization_accepted"`
	ProductionReadiness                 bool                                                                  `json:"production_readiness"`
	ProductionOwnershipReady            bool                                                                  `json:"production_ownership_ready"`
	VisibilityItemCount                 int                                                                   `json:"visibility_item_count"`
	RequiredVisibilityItemCount         int                                                                   `json:"required_visibility_item_count"`
	ReadyVisibilityItemCount            int                                                                   `json:"ready_visibility_item_count"`
	MissingVisibilityItemCount          int                                                                   `json:"missing_visibility_item_count"`
	PersistedVisibilityItemCount        int                                                                   `json:"persisted_visibility_item_count"`
	ExecutedDryRunItemCount             int                                                                   `json:"executed_dry_run_item_count"`
	DispatchedDryRunItemCount           int                                                                   `json:"dispatched_dry_run_item_count"`
	CreatedRequestObjectCount           int                                                                   `json:"created_request_object_count"`
	PortalRequestCreatedCount           int                                                                   `json:"portal_request_created_count"`
	SideEffectVisibilityItemCount       int                                                                   `json:"side_effect_visibility_item_count"`
	VisibilityItems                     []ProductionReceiptNotificationActionDryRunResultVisibilityAuditItem  `json:"visibility_items"`
	VisibilityItemIDs                   []string                                                              `json:"visibility_item_ids"`
	RequiredBeforeResultVisibility      []string                                                              `json:"required_before_result_visibility"`
	Checks                              []ProductionReceiptNotificationActionDryRunResultVisibilityAuditCheck `json:"checks"`
	CheckIDs                            []string                                                              `json:"check_ids"`
	Counts                              ProductionReceiptNotificationActionDryRunResultVisibilityAuditCounts  `json:"counts"`
	RuntimeOwned                        bool                                                                  `json:"runtime_owned"`
	GoRuntimeBacked                     bool                                                                  `json:"go_runtime_backed"`
	KDEPolicyOwner                      bool                                                                  `json:"kde_policy_owner"`
	OfficialDesktopOnly                 bool                                                                  `json:"official_desktop_only"`
	PlasmaForkRequired                  bool                                                                  `json:"plasma_fork_required"`
	PlasmaSourceModified                bool                                                                  `json:"plasma_source_modified"`
	SystemServiceStarted                bool                                                                  `json:"system_service_started"`
	SessionBusClaimed                   bool                                                                  `json:"session_bus_claimed"`
	ProductionBusClaimed                bool                                                                  `json:"production_bus_claimed"`
	ProductionOwnerEnabled              bool                                                                  `json:"production_owner_enabled"`
	ProductionActivationReady           bool                                                                  `json:"production_activation_ready"`
	WriteMethodsEnabled                 bool                                                                  `json:"write_methods_enabled"`
	RuntimeWritesEnabled                bool                                                                  `json:"runtime_writes_enabled"`
	RequestObjectCreationEnabled        bool                                                                  `json:"request_object_creation_enabled"`
	RequestObjectDispatchEnabled        bool                                                                  `json:"request_object_dispatch_enabled"`
	RequestObjectPersistenceEnabled     bool                                                                  `json:"request_object_persistence_enabled"`
	DispatchAuthorizationPersisted      bool                                                                  `json:"dispatch_authorization_persisted"`
	DispatchDryRunExecutionEnabled      bool                                                                  `json:"dispatch_dry_run_execution_enabled"`
	DryRunResultPersistenceEnabled      bool                                                                  `json:"dry_run_result_persistence_enabled"`
	ResultVisibilityPersistenceEnabled  bool                                                                  `json:"result_visibility_persistence_enabled"`
	PortalRequestCreated                bool                                                                  `json:"portal_request_created"`
	RequestObjectsCreated               bool                                                                  `json:"request_objects_created"`
	RequestObjectsDispatched            bool                                                                  `json:"request_objects_dispatched"`
	NotificationSent                    bool                                                                  `json:"notification_sent"`
	NotificationDeliveryEnabled         bool                                                                  `json:"notification_delivery_enabled"`
	NotificationActionEnabled           bool                                                                  `json:"notification_action_enabled"`
	ReviewActionEnabled                 bool                                                                  `json:"review_action_enabled"`
	RenewActionEnabled                  bool                                                                  `json:"renew_action_enabled"`
	OpenCompatibilityCenterEnabled      bool                                                                  `json:"open_compatibility_center_enabled"`
	DismissActionEnabled                bool                                                                  `json:"dismiss_action_enabled"`
	SupportInfoActionEnabled            bool                                                                  `json:"support_info_action_enabled"`
	CompatibilityCenterOpened           bool                                                                  `json:"compatibility_center_opened"`
	CompatibilityCenterPersisted        bool                                                                  `json:"compatibility_center_persisted"`
	RuntimeDiagnosticsPersisted         bool                                                                  `json:"runtime_diagnostics_persisted"`
	ReceiptWriterEnabled                bool                                                                  `json:"receipt_writer_enabled"`
	ReceiptPersistenceEnabled           bool                                                                  `json:"receipt_persistence_enabled"`
	ReceiptLookupWritesEnabled          bool                                                                  `json:"receipt_lookup_writes_enabled"`
	ReceiptReplayEnabled                bool                                                                  `json:"receipt_replay_enabled"`
	ReceiptExpiryWriteEnabled           bool                                                                  `json:"receipt_expiry_write_enabled"`
	ReceiptRevocationWriteEnabled       bool                                                                  `json:"receipt_revocation_write_enabled"`
	DesktopFilesWritten                 bool                                                                  `json:"desktop_files_written"`
	MIMEAppsWritten                     bool                                                                  `json:"mimeapps_written"`
	ShellConfigurationWritten           bool                                                                  `json:"shell_configuration_written"`
	SettingsPersisted                   bool                                                                  `json:"settings_persisted"`
	KRunnerIndexPersisted               bool                                                                  `json:"krunner_index_persisted"`
	TaskManagerEntryActive              bool                                                                  `json:"task_manager_entry_active"`
	KWinRuleApplied                     bool                                                                  `json:"kwin_rule_applied"`
	LiveTrayBridgeEnabled               bool                                                                  `json:"live_tray_bridge_enabled"`
	TrayBridgePersisted                 bool                                                                  `json:"tray_bridge_persisted"`
	AdapterInvocationEnabled            bool                                                                  `json:"adapter_invocation_enabled"`
	BackendLaunchEnabled                bool                                                                  `json:"backend_launch_enabled"`
	BackendProcessStarted               bool                                                                  `json:"backend_process_started"`
	SupportBundleExported               bool                                                                  `json:"support_bundle_exported"`
	SupportCaseCreated                  bool                                                                  `json:"support_case_created"`
	SnapshotRestoreExecuted             bool                                                                  `json:"snapshot_restore_executed"`
	StateCleanupExecuted                bool                                                                  `json:"state_cleanup_executed"`
	FileContentRead                     bool                                                                  `json:"file_content_read"`
	FilePathsExposed                    bool                                                                  `json:"file_paths_exposed"`
	NetworkRequired                     bool                                                                  `json:"network_required"`
	HostRootModified                    bool                                                                  `json:"host_root_modified"`
	PrivilegedContainerRequired         bool                                                                  `json:"privileged_container_required"`
	StateRootPathExposed                bool                                                                  `json:"state_root_path_exposed"`
	RawCommandExposed                   bool                                                                  `json:"raw_command_exposed"`
	RawExecutableExposed                bool                                                                  `json:"raw_executable_exposed"`
	BackendDetailsExposed               bool                                                                  `json:"backend_details_exposed"`
	BlockedActions                      []string                                                              `json:"blocked_actions"`
	NextRequirements                    []string                                                              `json:"next_requirements"`
	DesktopSafeSummary                  string                                                                `json:"desktop_safe_summary"`
}

type ProductionReceiptNotificationActionDryRunResultVisibilityAuditItem struct {
	ID                           string `json:"id"`
	ActionKind                   string `json:"action_kind"`
	RequestObjectKind            string `json:"request_object_kind"`
	DryRunKind                   string `json:"dry_run_kind"`
	ResultVisibilityKind         string `json:"result_visibility_kind"`
	KDEVisibilitySurface         string `json:"kde_visibility_surface"`
	RuntimeDiagnosticsSurface    string `json:"runtime_diagnostics_surface"`
	RequiredEvidence             string `json:"required_evidence"`
	EvidencePresent              bool   `json:"evidence_present"`
	ResultVisibilityModeled      bool   `json:"result_visibility_modeled"`
	RedactedForKDE               bool   `json:"redacted_for_kde"`
	RuntimeDiagnosticsModeled    bool   `json:"runtime_diagnostics_modeled"`
	UserVisible                  bool   `json:"user_visible"`
	ReviewOnly                   bool   `json:"review_only"`
	RuntimeOwned                 bool   `json:"runtime_owned"`
	GoRuntimeBacked              bool   `json:"go_runtime_backed"`
	KDEPolicyOwner               bool   `json:"kde_policy_owner"`
	OperatorApprovalRequired     bool   `json:"operator_approval_required"`
	OperatorApprovalPresent      bool   `json:"operator_approval_present"`
	DispatchAuthorizationGranted bool   `json:"dispatch_authorization_granted"`
	DispatchDryRunExecuted       bool   `json:"dispatch_dry_run_executed"`
	DryRunResultPersisted        bool   `json:"dry_run_result_persisted"`
	ResultVisibilityPersisted    bool   `json:"result_visibility_persisted"`
	RequestObjectCreated         bool   `json:"request_object_created"`
	RequestObjectDispatched      bool   `json:"request_object_dispatched"`
	RequestObjectPersisted       bool   `json:"request_object_persisted"`
	PortalRequestCreated         bool   `json:"portal_request_created"`
	NavigationRequested          bool   `json:"navigation_requested"`
	NotificationActionEnabled    bool   `json:"notification_action_enabled"`
	ActionEnabled                bool   `json:"action_enabled"`
	ReceiptAccepted              bool   `json:"receipt_accepted"`
	AuthorizationAccepted        bool   `json:"authorization_accepted"`
	ReceiptWriterEnabled         bool   `json:"receipt_writer_enabled"`
	ReceiptPersistenceEnabled    bool   `json:"receipt_persistence_enabled"`
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
	VisibilityStatus             string `json:"visibility_status"`
	NextRequirement              string `json:"next_requirement"`
}

type ProductionReceiptNotificationActionDryRunResultVisibilityAuditCheck struct {
	ID      string `json:"id"`
	Status  string `json:"status"`
	Summary string `json:"summary"`
}

type ProductionReceiptNotificationActionDryRunResultVisibilityAuditCounts struct {
	Total   int `json:"total"`
	Passed  int `json:"passed"`
	Pending int `json:"pending"`
	Blocked int `json:"blocked"`
}

func NewProductionReceiptNotificationActionDryRunResultVisibilityAuditPreview(root string) (ProductionReceiptNotificationActionDryRunResultVisibilityAuditPreview, error) {
	version, err := readProductionDBusGateReviewVersion(root)
	if err != nil {
		return ProductionReceiptNotificationActionDryRunResultVisibilityAuditPreview{}, err
	}
	sources := productionReceiptNotificationActionDryRunResultVisibilityAuditSources(root)
	items := productionReceiptNotificationActionDryRunResultVisibilityItems(sources)
	dryRunReady := productionReceiptNotificationActionDryRunResultVisibilityDryRunReady(sources.DispatchDryRunAudit)
	guidanceReady := productionReceiptNotificationActionDryRunResultVisibilityGuidanceReady(sources.DispatchSheet)
	planReady := dryRunReady && guidanceReady && productionReceiptNotificationActionDryRunResultVisibilityItemsReady(items)
	preview := ProductionReceiptNotificationActionDryRunResultVisibilityAuditPreview{
		Version:                             version,
		SchemaVersion:                       "xnix.runtime.production_receipt_notification_action_dry_run_result_visibility_audit.v1",
		RequestType:                         "production-receipt-notification-action-dry-run-result-visibility-audit-preview",
		AuditType:                           "receipt-notification-action-dry-run-result-visibility-audit",
		Source:                              "production-receipt-notification-action-dispatch-dry-run-audit-preview+claude-code-current-dispatch-picks",
		AuditDecision:                       "production-receipt-notification-action-dry-run-result-visibility-audit-blocked",
		ReceiptSchema:                       "xnix.runtime.production_dbus_human_authorization_receipt.v1",
		OpaqueReceiptID:                     ProductionDBusHumanAuthorizationReceiptID,
		ResultVisibilityAuditRequired:       true,
		ResultVisibilityAuditModeled:        true,
		DispatchDryRunAuditConsumed:         dryRunReady,
		ResultVisibilityGuidanceConsumed:    guidanceReady,
		ResultVisibilityPlanReady:           planReady,
		RedactedKDEVisibilityModeled:        planReady,
		RuntimeDiagnosticsVisibilityModeled: planReady,
		DispatchDryRunExecuted:              false,
		DryRunResultPersisted:               false,
		DispatchAuthorizationGranted:        false,
		OperatorDispatchApprovalRequired:    true,
		OperatorDispatchApprovalPresent:     false,
		CallerStateRootRequired:             false,
		ReceiptPresent:                      false,
		ReceiptAccepted:                     false,
		AuthorizationAccepted:               false,
		ProductionReadiness:                 false,
		ProductionOwnershipReady:            false,
		VisibilityItemCount:                 len(items),
		RequiredVisibilityItemCount:         5,
		ReadyVisibilityItemCount:            productionReceiptNotificationActionDryRunResultVisibilityReadyCount(items),
		MissingVisibilityItemCount:          productionReceiptNotificationActionDryRunResultVisibilityMissingCount(items),
		VisibilityItems:                     items,
		VisibilityItemIDs:                   productionReceiptNotificationActionDryRunResultVisibilityItemIDs(items),
		RequiredBeforeResultVisibility: []string{
			"production-receipt-notification-action-dispatch-dry-run-audit-preview",
			"redacted KDE-visible result vocabulary",
			"Runtime diagnostics result summary contract",
			"separate dry-run executor implementation",
			"separate result persistence authorization",
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
			"treat result visibility audit as permission to execute dispatch dry runs",
			"persist dry-run results or result visibility records from this audit",
			"create, persist, or dispatch notification action request objects from this audit",
			"create Portal requests, navigate KDE surfaces, write receipts, emit notifications, or export support data from this audit",
			"claim production D-Bus ownership, start services, launch compatibility engines, expose paths, or mutate host root",
		},
		NextRequirements: []string{
			"Define a separate dry-run result persistence authorization before any result can be stored.",
			"Define redacted KDE copy for each result state before KDE surfaces can render real dry-run output.",
			"Keep Runtime diagnostics visibility read-only until a separate dry-run executor produces redacted evidence.",
			"Require explicit operator dispatch approval and an accepted opaque authorization receipt before dry-run execution.",
		},
		DesktopSafeSummary: "The production receipt notification action dry-run result visibility audit models how future dry-run results for review, renew, open Compatibility Center, dismiss, and support-info actions would be redacted for KDE and summarized for Runtime diagnostics, but it executes no dry runs, persists no results, grants no dispatch authorization, creates no request objects, sends no Portal requests, opens no surfaces, writes no receipts, starts no services, launches no engines, and mutates no host state.",
	}
	preview.PersistedVisibilityItemCount = productionReceiptNotificationActionDryRunResultVisibilityPersistedCount(items)
	preview.ExecutedDryRunItemCount = productionReceiptNotificationActionDryRunResultVisibilityExecutedCount(items)
	preview.DispatchedDryRunItemCount = productionReceiptNotificationActionDryRunResultVisibilityDispatchedCount(items)
	preview.CreatedRequestObjectCount = productionReceiptNotificationActionDryRunResultVisibilityCreatedCount(items)
	preview.PortalRequestCreatedCount = productionReceiptNotificationActionDryRunResultVisibilityPortalCount(items)
	preview.SideEffectVisibilityItemCount = productionReceiptNotificationActionDryRunResultVisibilitySideEffectCount(items)
	checks := productionReceiptNotificationActionDryRunResultVisibilityAuditChecks(preview)
	preview.Checks = checks
	preview.CheckIDs = productionReceiptNotificationActionDryRunResultVisibilityCheckIDs(checks)
	preview.Counts = countProductionReceiptNotificationActionDryRunResultVisibilityChecks(checks)
	if preview.Counts.Blocked == 0 {
		preview.AuditDecision = "production-receipt-notification-action-dry-run-result-visibility-audit-ready-visibility-only"
	}
	if err := validateNoBackendTerms(preview, "production receipt notification action dry-run result visibility audit preview"); err != nil {
		return ProductionReceiptNotificationActionDryRunResultVisibilityAuditPreview{}, err
	}
	return preview, nil
}

type productionReceiptNotificationActionDryRunResultVisibilityAuditSourceSet struct {
	DispatchDryRunAudit string
	DispatchSheet       string
}

func productionReceiptNotificationActionDryRunResultVisibilityAuditSources(root string) productionReceiptNotificationActionDryRunResultVisibilityAuditSourceSet {
	return productionReceiptNotificationActionDryRunResultVisibilityAuditSourceSet{
		DispatchDryRunAudit: productionAuthorizationReadSources(root, []string{"internal/runtime/owner/production_receipt_notification_action_dispatch_dry_run_audit.go"}),
		DispatchSheet:       productionAuthorizationReadSources(root, []string{"docs/claude-code-current-dispatch-picks.md"}),
	}
}

func productionReceiptNotificationActionDryRunResultVisibilityItems(sources productionReceiptNotificationActionDryRunResultVisibilityAuditSourceSet) []ProductionReceiptNotificationActionDryRunResultVisibilityAuditItem {
	combined := sources.DispatchDryRunAudit + sources.DispatchSheet
	return []ProductionReceiptNotificationActionDryRunResultVisibilityAuditItem{
		productionReceiptNotificationActionDryRunResultVisibilityItem("review-receipt-dry-run-result-visibility", "review", "receipt-review-request", "review-request-dispatch-dry-run", "review-dry-run-result-visibility", "notification-center-review-card", "runtime-diagnostics-review-row", "future review dry-run result requires redacted visibility", combined, []string{"review-receipt-dispatch-dry-run", "dispatch dry-run", "dry-run result visibility"}, "define review result persistence authorization before real visibility"),
		productionReceiptNotificationActionDryRunResultVisibilityItem("renew-receipt-dry-run-result-visibility", "renew", "receipt-renewal-request", "renewal-request-dispatch-dry-run", "renewal-dry-run-result-visibility", "notification-center-renewal-card", "runtime-diagnostics-renewal-row", "future renewal dry-run result requires redacted visibility", combined, []string{"renew-receipt-dispatch-dry-run", "dispatch dry-run", "dry-run result visibility"}, "define renewal result persistence authorization before real visibility"),
		productionReceiptNotificationActionDryRunResultVisibilityItem("open-compatibility-center-dry-run-result-visibility", "open-compatibility-center", "compatibility-center-navigation-request", "navigation-request-dispatch-dry-run", "navigation-dry-run-result-visibility", "compatibility-center-navigation-card", "runtime-diagnostics-navigation-row", "future navigation dry-run result requires redacted visibility", combined, []string{"open-compatibility-center-dispatch-dry-run", "dispatch dry-run", "dry-run result visibility"}, "define navigation result persistence authorization before real visibility"),
		productionReceiptNotificationActionDryRunResultVisibilityItem("dismiss-receipt-dry-run-result-visibility", "dismiss", "receipt-dismissal-request", "dismissal-request-dispatch-dry-run", "dismissal-dry-run-result-visibility", "notification-center-dismissal-card", "runtime-diagnostics-dismissal-row", "future dismissal dry-run result requires redacted visibility", combined, []string{"dismiss-receipt-dispatch-dry-run", "dispatch dry-run", "dry-run result visibility"}, "define dismissal result persistence authorization before real visibility"),
		productionReceiptNotificationActionDryRunResultVisibilityItem("support-info-dry-run-result-visibility", "support-info", "support-info-request", "support-info-request-dispatch-dry-run", "support-info-dry-run-result-visibility", "compatibility-center-support-card", "runtime-diagnostics-support-row", "future support-info dry-run result requires redacted visibility", combined, []string{"support-info-dispatch-dry-run", "dispatch dry-run", "dry-run result visibility"}, "define support-info result persistence authorization before real visibility"),
	}
}

func productionReceiptNotificationActionDryRunResultVisibilityItem(id string, actionKind string, requestKind string, dryRunKind string, visibilityKind string, kdeSurface string, diagnosticsSurface string, evidence string, source string, tokens []string, nextRequirement string) ProductionReceiptNotificationActionDryRunResultVisibilityAuditItem {
	ready := productionAuthorizationHasAll(source, tokens)
	status := "missing-dry-run-result-visibility-evidence"
	if ready {
		status = "dry-run-result-visibility-modeled-persistence-disabled"
	}
	return ProductionReceiptNotificationActionDryRunResultVisibilityAuditItem{
		ID:                           id,
		ActionKind:                   actionKind,
		RequestObjectKind:            requestKind,
		DryRunKind:                   dryRunKind,
		ResultVisibilityKind:         visibilityKind,
		KDEVisibilitySurface:         kdeSurface,
		RuntimeDiagnosticsSurface:    diagnosticsSurface,
		RequiredEvidence:             evidence,
		EvidencePresent:              ready,
		ResultVisibilityModeled:      ready,
		RedactedForKDE:               ready,
		RuntimeDiagnosticsModeled:    ready,
		UserVisible:                  ready,
		ReviewOnly:                   true,
		RuntimeOwned:                 true,
		GoRuntimeBacked:              true,
		KDEPolicyOwner:               false,
		OperatorApprovalRequired:     true,
		OperatorApprovalPresent:      false,
		DispatchAuthorizationGranted: false,
		DispatchDryRunExecuted:       false,
		DryRunResultPersisted:        false,
		ResultVisibilityPersisted:    false,
		RequestObjectCreated:         false,
		RequestObjectDispatched:      false,
		RequestObjectPersisted:       false,
		PortalRequestCreated:         false,
		NavigationRequested:          false,
		NotificationActionEnabled:    false,
		ActionEnabled:                false,
		ReceiptAccepted:              false,
		AuthorizationAccepted:        false,
		ReceiptWriterEnabled:         false,
		ReceiptPersistenceEnabled:    false,
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
		VisibilityStatus:             status,
		NextRequirement:              nextRequirement,
	}
}

func productionReceiptNotificationActionDryRunResultVisibilityAuditChecks(preview ProductionReceiptNotificationActionDryRunResultVisibilityAuditPreview) []ProductionReceiptNotificationActionDryRunResultVisibilityAuditCheck {
	return []ProductionReceiptNotificationActionDryRunResultVisibilityAuditCheck{
		productionReceiptNotificationActionDryRunResultVisibilityCheck("dispatch-dry-run-audit-consumed", productionAuthorizationPassBlocked(preview.DispatchDryRunAuditConsumed), "The result visibility audit consumes dispatch dry-run boundary evidence."),
		productionReceiptNotificationActionDryRunResultVisibilityCheck("dry-run-result-visibility-guidance-consumed", productionAuthorizationPassBlocked(preview.ResultVisibilityGuidanceConsumed), "The result visibility audit consumes current dispatch guidance for redacted KDE and Runtime diagnostics visibility."),
		productionReceiptNotificationActionDryRunResultVisibilityCheck("dry-run-result-visibility-modeled-only", productionAuthorizationPassBlocked(preview.ResultVisibilityAuditRequired && preview.ResultVisibilityAuditModeled && preview.ResultVisibilityPlanReady && preview.RedactedKDEVisibilityModeled && preview.RuntimeDiagnosticsVisibilityModeled && preview.OperatorDispatchApprovalRequired && !preview.OperatorDispatchApprovalPresent && !preview.DispatchAuthorizationGranted && !preview.DispatchDryRunExecuted && !preview.DryRunResultPersisted && !preview.ReceiptPresent && !preview.ReceiptAccepted && !preview.AuthorizationAccepted), "Dry-run result visibility is modeled without approval, authorization grant, execution, result persistence, accepted receipt, or production readiness."),
		productionReceiptNotificationActionDryRunResultVisibilityCheck("five-dry-run-result-visibility-items-present", productionAuthorizationPassBlocked(preview.VisibilityItemCount == 5 && preview.RequiredVisibilityItemCount == 5 && preview.MissingVisibilityItemCount == 0), "Review, renew, open Compatibility Center, dismiss, and support-info result visibility items are modeled."),
		productionReceiptNotificationActionDryRunResultVisibilityCheck("visibility-items-ready-results-redacted", productionAuthorizationPassBlocked(preview.ReadyVisibilityItemCount == 5 && productionReceiptNotificationActionDryRunResultVisibilityItemsReady(preview.VisibilityItems)), "Every result visibility item is ready with redacted KDE and Runtime diagnostics surfaces while persistence and execution remain disabled."),
		productionReceiptNotificationActionDryRunResultVisibilityCheck("dry-run-execution-and-result-persistence-disabled", productionAuthorizationPassBlocked(!preview.DispatchDryRunExecutionEnabled && !preview.DispatchDryRunExecuted && !preview.DryRunResultPersisted && !preview.DryRunResultPersistenceEnabled && !preview.ResultVisibilityPersistenceEnabled && !preview.RuntimeDiagnosticsPersisted && preview.PersistedVisibilityItemCount == 0 && preview.ExecutedDryRunItemCount == 0 && preview.DispatchedDryRunItemCount == 0), "Dry-run execution, result persistence, visibility persistence, Runtime diagnostics persistence, per-item execution, and request dispatch remain disabled."),
		productionReceiptNotificationActionDryRunResultVisibilityCheck("request-creation-dispatch-actions-and-portal-disabled", productionAuthorizationPassBlocked(!preview.RequestObjectCreationEnabled && !preview.RequestObjectDispatchEnabled && !preview.RequestObjectPersistenceEnabled && !preview.RequestObjectsCreated && !preview.RequestObjectsDispatched && !preview.PortalRequestCreated && !preview.NotificationActionEnabled && !preview.ReviewActionEnabled && !preview.RenewActionEnabled && !preview.OpenCompatibilityCenterEnabled && !preview.DismissActionEnabled && !preview.SupportInfoActionEnabled && preview.CreatedRequestObjectCount == 0 && preview.PortalRequestCreatedCount == 0), "Request creation, dispatch, persistence, notification actions, and Portal requests remain disabled."),
		productionReceiptNotificationActionDryRunResultVisibilityCheck("receipt-notification-navigation-and-support-writes-disabled", productionAuthorizationPassBlocked(!preview.NotificationSent && !preview.NotificationDeliveryEnabled && !preview.CompatibilityCenterOpened && !preview.CompatibilityCenterPersisted && !preview.ReceiptWriterEnabled && !preview.ReceiptPersistenceEnabled && !preview.ReceiptLookupWritesEnabled && !preview.ReceiptReplayEnabled && !preview.ReceiptExpiryWriteEnabled && !preview.ReceiptRevocationWriteEnabled && !preview.SupportBundleExported && !preview.SupportCaseCreated), "Notification, navigation, receipt, Compatibility Center, and support write side effects remain disabled."),
		productionReceiptNotificationActionDryRunResultVisibilityCheck("production-and-host-boundary-closed", productionAuthorizationPassBlocked(!preview.SystemServiceStarted && !preview.SessionBusClaimed && !preview.ProductionBusClaimed && !preview.ProductionOwnerEnabled && !preview.ProductionActivationReady && !preview.ProductionReadiness && !preview.ProductionOwnershipReady && !preview.AdapterInvocationEnabled && !preview.BackendLaunchEnabled && !preview.BackendProcessStarted && !preview.FileContentRead && !preview.FilePathsExposed && !preview.NetworkRequired && !preview.HostRootModified && !preview.PrivilegedContainerRequired && !preview.StateRootPathExposed && !preview.RawCommandExposed && !preview.RawExecutableExposed && !preview.BackendDetailsExposed && productionReceiptNotificationActionDryRunResultVisibilityItemsKeepHostClosed(preview.VisibilityItems)), "Production ownership, backend launch, unsafe data exposure, network, privilege, path exposure, internal details, and host mutation remain disabled."),
	}
}

func productionReceiptNotificationActionDryRunResultVisibilityCheck(id string, status string, summary string) ProductionReceiptNotificationActionDryRunResultVisibilityAuditCheck {
	return ProductionReceiptNotificationActionDryRunResultVisibilityAuditCheck{ID: id, Status: status, Summary: summary}
}

func productionReceiptNotificationActionDryRunResultVisibilityDryRunReady(source string) bool {
	return productionAuthorizationHasAll(source, []string{
		"production-receipt-notification-action-dispatch-dry-run-audit-preview",
		"receipt-notification-action-request-dispatch-dry-run-audit",
		"production-receipt-notification-action-dispatch-dry-run-audit-ready-dry-run-only",
		"review-receipt-dispatch-dry-run",
		"renew-receipt-dispatch-dry-run",
		"open-compatibility-center-dispatch-dry-run",
		"dismiss-receipt-dispatch-dry-run",
		"support-info-dispatch-dry-run",
		"DispatchDryRunExecuted",
		"DryRunResultPersisted",
	})
}

func productionReceiptNotificationActionDryRunResultVisibilityGuidanceReady(source string) bool {
	return productionAuthorizationHasAll(source, []string{
		"dry-run result visibility",
		"KDE",
		"Runtime diagnostics",
		"result persistence",
	})
}

func productionReceiptNotificationActionDryRunResultVisibilityReadyCount(items []ProductionReceiptNotificationActionDryRunResultVisibilityAuditItem) int {
	count := 0
	for _, item := range items {
		if item.EvidencePresent {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionDryRunResultVisibilityMissingCount(items []ProductionReceiptNotificationActionDryRunResultVisibilityAuditItem) int {
	count := 0
	for _, item := range items {
		if !item.EvidencePresent {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionDryRunResultVisibilityPersistedCount(items []ProductionReceiptNotificationActionDryRunResultVisibilityAuditItem) int {
	count := 0
	for _, item := range items {
		if item.DryRunResultPersisted || item.ResultVisibilityPersisted || item.RuntimeDiagnosticsPersisted {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionDryRunResultVisibilityExecutedCount(items []ProductionReceiptNotificationActionDryRunResultVisibilityAuditItem) int {
	count := 0
	for _, item := range items {
		if item.DispatchDryRunExecuted {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionDryRunResultVisibilityDispatchedCount(items []ProductionReceiptNotificationActionDryRunResultVisibilityAuditItem) int {
	count := 0
	for _, item := range items {
		if item.RequestObjectDispatched {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionDryRunResultVisibilityCreatedCount(items []ProductionReceiptNotificationActionDryRunResultVisibilityAuditItem) int {
	count := 0
	for _, item := range items {
		if item.RequestObjectCreated {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionDryRunResultVisibilityPortalCount(items []ProductionReceiptNotificationActionDryRunResultVisibilityAuditItem) int {
	count := 0
	for _, item := range items {
		if item.PortalRequestCreated {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionDryRunResultVisibilitySideEffectCount(items []ProductionReceiptNotificationActionDryRunResultVisibilityAuditItem) int {
	count := 0
	for _, item := range items {
		if item.DispatchAuthorizationGranted || item.DispatchDryRunExecuted || item.DryRunResultPersisted || item.ResultVisibilityPersisted || item.RuntimeDiagnosticsPersisted || item.RequestObjectCreated || item.RequestObjectDispatched || item.RequestObjectPersisted || item.PortalRequestCreated || item.NavigationRequested || item.NotificationActionEnabled || item.ActionEnabled || item.ReceiptAccepted || item.AuthorizationAccepted || item.ReceiptWriterEnabled || item.ReceiptPersistenceEnabled || item.CompatibilityCenterOpened || item.CompatibilityCenterPersisted || item.SupportBundleExported || item.SupportCaseCreated || item.HostRootModified || item.InternalDetailsExposed {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionDryRunResultVisibilityItemIDs(items []ProductionReceiptNotificationActionDryRunResultVisibilityAuditItem) []string {
	ids := make([]string, 0, len(items))
	for _, item := range items {
		ids = append(ids, item.ID)
	}
	return ids
}

func productionReceiptNotificationActionDryRunResultVisibilityCheckIDs(checks []ProductionReceiptNotificationActionDryRunResultVisibilityAuditCheck) []string {
	ids := make([]string, 0, len(checks))
	for _, check := range checks {
		ids = append(ids, check.ID)
	}
	return ids
}

func countProductionReceiptNotificationActionDryRunResultVisibilityChecks(checks []ProductionReceiptNotificationActionDryRunResultVisibilityAuditCheck) ProductionReceiptNotificationActionDryRunResultVisibilityAuditCounts {
	counts := ProductionReceiptNotificationActionDryRunResultVisibilityAuditCounts{Total: len(checks)}
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

func productionReceiptNotificationActionDryRunResultVisibilityItemsReady(items []ProductionReceiptNotificationActionDryRunResultVisibilityAuditItem) bool {
	if len(items) != 5 {
		return false
	}
	for _, item := range items {
		if !item.EvidencePresent || !item.ResultVisibilityModeled || !item.RedactedForKDE || !item.RuntimeDiagnosticsModeled || !item.UserVisible || item.VisibilityStatus != "dry-run-result-visibility-modeled-persistence-disabled" {
			return false
		}
		if item.OperatorApprovalPresent || item.DispatchAuthorizationGranted || item.DispatchDryRunExecuted || item.DryRunResultPersisted || item.ResultVisibilityPersisted || item.RuntimeDiagnosticsPersisted || item.RequestObjectCreated || item.RequestObjectDispatched || item.RequestObjectPersisted || item.PortalRequestCreated || item.NavigationRequested || item.NotificationActionEnabled || item.ActionEnabled || item.ReceiptAccepted || item.AuthorizationAccepted || item.ReceiptWriterEnabled || item.ReceiptPersistenceEnabled || item.CompatibilityCenterOpened || item.CompatibilityCenterPersisted || item.SupportBundleExported || item.SupportCaseCreated || item.ProductionReadiness || item.ProductionOwnershipReady || item.KDEPolicyOwner || !item.ReviewOnly || !item.SideEffectsDisabled {
			return false
		}
	}
	return true
}

func productionReceiptNotificationActionDryRunResultVisibilityItemsKeepHostClosed(items []ProductionReceiptNotificationActionDryRunResultVisibilityAuditItem) bool {
	for _, item := range items {
		if item.HostRootModified || item.InternalDetailsExposed {
			return false
		}
	}
	return true
}
