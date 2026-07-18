package owner

type ProductionReceiptNotificationActionDispatchDryRunAuditPreview struct {
	Version                            string                                                        `json:"version"`
	SchemaVersion                      string                                                        `json:"schema_version"`
	RequestType                        string                                                        `json:"request_type"`
	AuditType                          string                                                        `json:"audit_type"`
	Source                             string                                                        `json:"source"`
	AuditDecision                      string                                                        `json:"audit_decision"`
	ReceiptSchema                      string                                                        `json:"receipt_schema"`
	OpaqueReceiptID                    string                                                        `json:"opaque_receipt_id"`
	DispatchDryRunAuditRequired        bool                                                          `json:"dispatch_dry_run_audit_required"`
	DispatchDryRunAuditModeled         bool                                                          `json:"dispatch_dry_run_audit_modeled"`
	DispatchAuthorizationAuditConsumed bool                                                          `json:"dispatch_authorization_audit_consumed"`
	DispatchDryRunGuidanceConsumed     bool                                                          `json:"dispatch_dry_run_guidance_consumed"`
	DispatchDryRunPlanReady            bool                                                          `json:"dispatch_dry_run_plan_ready"`
	DispatchDryRunExecuted             bool                                                          `json:"dispatch_dry_run_executed"`
	DispatchAuthorizationGranted       bool                                                          `json:"dispatch_authorization_granted"`
	OperatorDispatchApprovalRequired   bool                                                          `json:"operator_dispatch_approval_required"`
	OperatorDispatchApprovalPresent    bool                                                          `json:"operator_dispatch_approval_present"`
	CallerStateRootRequired            bool                                                          `json:"caller_state_root_required"`
	ReceiptPresent                     bool                                                          `json:"receipt_present"`
	ReceiptAccepted                    bool                                                          `json:"receipt_accepted"`
	AuthorizationAccepted              bool                                                          `json:"authorization_accepted"`
	ProductionReadiness                bool                                                          `json:"production_readiness"`
	ProductionOwnershipReady           bool                                                          `json:"production_ownership_ready"`
	DryRunItemCount                    int                                                           `json:"dry_run_item_count"`
	RequiredDryRunItemCount            int                                                           `json:"required_dry_run_item_count"`
	ReadyDryRunItemCount               int                                                           `json:"ready_dry_run_item_count"`
	MissingDryRunItemCount             int                                                           `json:"missing_dry_run_item_count"`
	ExecutedDryRunItemCount            int                                                           `json:"executed_dry_run_item_count"`
	DispatchedDryRunItemCount          int                                                           `json:"dispatched_dry_run_item_count"`
	CreatedRequestObjectCount          int                                                           `json:"created_request_object_count"`
	PortalRequestCreatedCount          int                                                           `json:"portal_request_created_count"`
	SideEffectDryRunItemCount          int                                                           `json:"side_effect_dry_run_item_count"`
	DryRunItems                        []ProductionReceiptNotificationActionDispatchDryRunAuditItem  `json:"dry_run_items"`
	DryRunItemIDs                      []string                                                      `json:"dry_run_item_ids"`
	RequiredBeforeDryRun               []string                                                      `json:"required_before_dry_run"`
	Checks                             []ProductionReceiptNotificationActionDispatchDryRunAuditCheck `json:"checks"`
	CheckIDs                           []string                                                      `json:"check_ids"`
	Counts                             ProductionReceiptNotificationActionDispatchDryRunAuditCounts  `json:"counts"`
	RuntimeOwned                       bool                                                          `json:"runtime_owned"`
	GoRuntimeBacked                    bool                                                          `json:"go_runtime_backed"`
	KDEPolicyOwner                     bool                                                          `json:"kde_policy_owner"`
	OfficialDesktopOnly                bool                                                          `json:"official_desktop_only"`
	PlasmaForkRequired                 bool                                                          `json:"plasma_fork_required"`
	PlasmaSourceModified               bool                                                          `json:"plasma_source_modified"`
	SystemServiceStarted               bool                                                          `json:"system_service_started"`
	SessionBusClaimed                  bool                                                          `json:"session_bus_claimed"`
	ProductionBusClaimed               bool                                                          `json:"production_bus_claimed"`
	ProductionOwnerEnabled             bool                                                          `json:"production_owner_enabled"`
	ProductionActivationReady          bool                                                          `json:"production_activation_ready"`
	WriteMethodsEnabled                bool                                                          `json:"write_methods_enabled"`
	RuntimeWritesEnabled               bool                                                          `json:"runtime_writes_enabled"`
	RequestObjectCreationEnabled       bool                                                          `json:"request_object_creation_enabled"`
	RequestObjectDispatchEnabled       bool                                                          `json:"request_object_dispatch_enabled"`
	RequestObjectPersistenceEnabled    bool                                                          `json:"request_object_persistence_enabled"`
	DispatchAuthorizationPersisted     bool                                                          `json:"dispatch_authorization_persisted"`
	DispatchDryRunExecutionEnabled     bool                                                          `json:"dispatch_dry_run_execution_enabled"`
	DryRunResultPersisted              bool                                                          `json:"dry_run_result_persisted"`
	PortalRequestCreated               bool                                                          `json:"portal_request_created"`
	RequestObjectsCreated              bool                                                          `json:"request_objects_created"`
	RequestObjectsDispatched           bool                                                          `json:"request_objects_dispatched"`
	NotificationSent                   bool                                                          `json:"notification_sent"`
	NotificationDeliveryEnabled        bool                                                          `json:"notification_delivery_enabled"`
	NotificationActionEnabled          bool                                                          `json:"notification_action_enabled"`
	ReviewActionEnabled                bool                                                          `json:"review_action_enabled"`
	RenewActionEnabled                 bool                                                          `json:"renew_action_enabled"`
	OpenCompatibilityCenterEnabled     bool                                                          `json:"open_compatibility_center_enabled"`
	DismissActionEnabled               bool                                                          `json:"dismiss_action_enabled"`
	SupportInfoActionEnabled           bool                                                          `json:"support_info_action_enabled"`
	CompatibilityCenterOpened          bool                                                          `json:"compatibility_center_opened"`
	CompatibilityCenterPersisted       bool                                                          `json:"compatibility_center_persisted"`
	ReceiptWriterEnabled               bool                                                          `json:"receipt_writer_enabled"`
	ReceiptPersistenceEnabled          bool                                                          `json:"receipt_persistence_enabled"`
	ReceiptLookupWritesEnabled         bool                                                          `json:"receipt_lookup_writes_enabled"`
	ReceiptReplayEnabled               bool                                                          `json:"receipt_replay_enabled"`
	ReceiptExpiryWriteEnabled          bool                                                          `json:"receipt_expiry_write_enabled"`
	ReceiptRevocationWriteEnabled      bool                                                          `json:"receipt_revocation_write_enabled"`
	DesktopFilesWritten                bool                                                          `json:"desktop_files_written"`
	MIMEAppsWritten                    bool                                                          `json:"mimeapps_written"`
	ShellConfigurationWritten          bool                                                          `json:"shell_configuration_written"`
	SettingsPersisted                  bool                                                          `json:"settings_persisted"`
	KRunnerIndexPersisted              bool                                                          `json:"krunner_index_persisted"`
	TaskManagerEntryActive             bool                                                          `json:"task_manager_entry_active"`
	KWinRuleApplied                    bool                                                          `json:"kwin_rule_applied"`
	LiveTrayBridgeEnabled              bool                                                          `json:"live_tray_bridge_enabled"`
	TrayBridgePersisted                bool                                                          `json:"tray_bridge_persisted"`
	AdapterInvocationEnabled           bool                                                          `json:"adapter_invocation_enabled"`
	BackendLaunchEnabled               bool                                                          `json:"backend_launch_enabled"`
	BackendProcessStarted              bool                                                          `json:"backend_process_started"`
	SupportBundleExported              bool                                                          `json:"support_bundle_exported"`
	SupportCaseCreated                 bool                                                          `json:"support_case_created"`
	SnapshotRestoreExecuted            bool                                                          `json:"snapshot_restore_executed"`
	StateCleanupExecuted               bool                                                          `json:"state_cleanup_executed"`
	FileContentRead                    bool                                                          `json:"file_content_read"`
	FilePathsExposed                   bool                                                          `json:"file_paths_exposed"`
	NetworkRequired                    bool                                                          `json:"network_required"`
	HostRootModified                   bool                                                          `json:"host_root_modified"`
	PrivilegedContainerRequired        bool                                                          `json:"privileged_container_required"`
	StateRootPathExposed               bool                                                          `json:"state_root_path_exposed"`
	RawCommandExposed                  bool                                                          `json:"raw_command_exposed"`
	RawExecutableExposed               bool                                                          `json:"raw_executable_exposed"`
	BackendDetailsExposed              bool                                                          `json:"backend_details_exposed"`
	BlockedActions                     []string                                                      `json:"blocked_actions"`
	NextRequirements                   []string                                                      `json:"next_requirements"`
	DesktopSafeSummary                 string                                                        `json:"desktop_safe_summary"`
}

type ProductionReceiptNotificationActionDispatchDryRunAuditItem struct {
	ID                           string `json:"id"`
	ActionKind                   string `json:"action_kind"`
	RequestObjectKind            string `json:"request_object_kind"`
	DispatchAuthorizationKind    string `json:"dispatch_authorization_kind"`
	DryRunKind                   string `json:"dry_run_kind"`
	RequiredEvidence             string `json:"required_evidence"`
	EvidencePresent              bool   `json:"evidence_present"`
	DispatchDryRunModeled        bool   `json:"dispatch_dry_run_modeled"`
	UserVisible                  bool   `json:"user_visible"`
	ReviewOnly                   bool   `json:"review_only"`
	RuntimeOwned                 bool   `json:"runtime_owned"`
	GoRuntimeBacked              bool   `json:"go_runtime_backed"`
	KDEPolicyOwner               bool   `json:"kde_policy_owner"`
	OperatorApprovalRequired     bool   `json:"operator_approval_required"`
	OperatorApprovalPresent      bool   `json:"operator_approval_present"`
	DispatchAuthorizationGranted bool   `json:"dispatch_authorization_granted"`
	DispatchDryRunExecuted       bool   `json:"dispatch_dry_run_executed"`
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
	SupportBundleExported        bool   `json:"support_bundle_exported"`
	SupportCaseCreated           bool   `json:"support_case_created"`
	ProductionReadiness          bool   `json:"production_readiness"`
	ProductionOwnershipReady     bool   `json:"production_ownership_ready"`
	SideEffectsDisabled          bool   `json:"side_effects_disabled"`
	HostRootModified             bool   `json:"host_root_modified"`
	InternalDetailsExposed       bool   `json:"internal_details_exposed"`
	DryRunStatus                 string `json:"dry_run_status"`
	NextRequirement              string `json:"next_requirement"`
}

type ProductionReceiptNotificationActionDispatchDryRunAuditCheck struct {
	ID      string `json:"id"`
	Status  string `json:"status"`
	Summary string `json:"summary"`
}

type ProductionReceiptNotificationActionDispatchDryRunAuditCounts struct {
	Total   int `json:"total"`
	Passed  int `json:"passed"`
	Pending int `json:"pending"`
	Blocked int `json:"blocked"`
}

func NewProductionReceiptNotificationActionDispatchDryRunAuditPreview(root string) (ProductionReceiptNotificationActionDispatchDryRunAuditPreview, error) {
	version, err := readProductionDBusGateReviewVersion(root)
	if err != nil {
		return ProductionReceiptNotificationActionDispatchDryRunAuditPreview{}, err
	}
	sources := productionReceiptNotificationActionDispatchDryRunAuditSources(root)
	items := productionReceiptNotificationActionDispatchDryRunItems(sources)
	authorizationReady := productionReceiptNotificationActionDispatchDryRunAuthorizationReady(sources.DispatchAuthorizationAudit)
	guidanceReady := productionReceiptNotificationActionDispatchDryRunGuidanceReady(sources.DispatchSheet)
	planReady := authorizationReady && guidanceReady && productionReceiptNotificationActionDispatchDryRunItemsReady(items)
	preview := ProductionReceiptNotificationActionDispatchDryRunAuditPreview{
		Version:                            version,
		SchemaVersion:                      "xnix.runtime.production_receipt_notification_action_dispatch_dry_run_audit.v1",
		RequestType:                        "production-receipt-notification-action-dispatch-dry-run-audit-preview",
		AuditType:                          "receipt-notification-action-request-dispatch-dry-run-audit",
		Source:                             "production-receipt-notification-action-dispatch-authorization-audit-preview+claude-code-current-dispatch-picks",
		AuditDecision:                      "production-receipt-notification-action-dispatch-dry-run-audit-blocked",
		ReceiptSchema:                      "xnix.runtime.production_dbus_human_authorization_receipt.v1",
		OpaqueReceiptID:                    ProductionDBusHumanAuthorizationReceiptID,
		DispatchDryRunAuditRequired:        true,
		DispatchDryRunAuditModeled:         true,
		DispatchAuthorizationAuditConsumed: authorizationReady,
		DispatchDryRunGuidanceConsumed:     guidanceReady,
		DispatchDryRunPlanReady:            planReady,
		DispatchDryRunExecuted:             false,
		DispatchAuthorizationGranted:       false,
		OperatorDispatchApprovalRequired:   true,
		OperatorDispatchApprovalPresent:    false,
		CallerStateRootRequired:            false,
		ReceiptPresent:                     false,
		ReceiptAccepted:                    false,
		AuthorizationAccepted:              false,
		ProductionReadiness:                false,
		ProductionOwnershipReady:           false,
		DryRunItemCount:                    len(items),
		RequiredDryRunItemCount:            5,
		ReadyDryRunItemCount:               productionReceiptNotificationActionDispatchDryRunReadyCount(items),
		MissingDryRunItemCount:             productionReceiptNotificationActionDispatchDryRunMissingCount(items),
		DryRunItems:                        items,
		DryRunItemIDs:                      productionReceiptNotificationActionDispatchDryRunItemIDs(items),
		RequiredBeforeDryRun: []string{
			"production-receipt-notification-action-dispatch-authorization-audit-preview",
			"explicit operator dispatch approval",
			"accepted opaque authorization receipt",
			"separate dry-run executor implementation",
			"redacted dry-run result visibility audit",
		},
		RuntimeOwned:                    true,
		GoRuntimeBacked:                 true,
		KDEPolicyOwner:                  false,
		OfficialDesktopOnly:             true,
		PlasmaForkRequired:              false,
		PlasmaSourceModified:            false,
		SystemServiceStarted:            false,
		SessionBusClaimed:               false,
		ProductionBusClaimed:            false,
		ProductionOwnerEnabled:          false,
		ProductionActivationReady:       false,
		WriteMethodsEnabled:             false,
		RuntimeWritesEnabled:            false,
		RequestObjectCreationEnabled:    false,
		RequestObjectDispatchEnabled:    false,
		RequestObjectPersistenceEnabled: false,
		DispatchAuthorizationPersisted:  false,
		DispatchDryRunExecutionEnabled:  false,
		DryRunResultPersisted:           false,
		PortalRequestCreated:            false,
		RequestObjectsCreated:           false,
		RequestObjectsDispatched:        false,
		NotificationSent:                false,
		NotificationDeliveryEnabled:     false,
		NotificationActionEnabled:       false,
		ReviewActionEnabled:             false,
		RenewActionEnabled:              false,
		OpenCompatibilityCenterEnabled:  false,
		DismissActionEnabled:            false,
		SupportInfoActionEnabled:        false,
		CompatibilityCenterOpened:       false,
		CompatibilityCenterPersisted:    false,
		ReceiptWriterEnabled:            false,
		ReceiptPersistenceEnabled:       false,
		ReceiptLookupWritesEnabled:      false,
		ReceiptReplayEnabled:            false,
		ReceiptExpiryWriteEnabled:       false,
		ReceiptRevocationWriteEnabled:   false,
		DesktopFilesWritten:             false,
		MIMEAppsWritten:                 false,
		ShellConfigurationWritten:       false,
		SettingsPersisted:               false,
		KRunnerIndexPersisted:           false,
		TaskManagerEntryActive:          false,
		KWinRuleApplied:                 false,
		LiveTrayBridgeEnabled:           false,
		TrayBridgePersisted:             false,
		AdapterInvocationEnabled:        false,
		BackendLaunchEnabled:            false,
		BackendProcessStarted:           false,
		SupportBundleExported:           false,
		SupportCaseCreated:              false,
		SnapshotRestoreExecuted:         false,
		StateCleanupExecuted:            false,
		FileContentRead:                 false,
		FilePathsExposed:                false,
		NetworkRequired:                 false,
		HostRootModified:                false,
		PrivilegedContainerRequired:     false,
		StateRootPathExposed:            false,
		RawCommandExposed:               false,
		RawExecutableExposed:            false,
		BackendDetailsExposed:           false,
		BlockedActions: []string{
			"treat dispatch dry-run audit as permission to execute dispatch dry runs",
			"create, persist, or dispatch notification action request objects from this audit",
			"create Portal requests, navigate KDE surfaces, write receipts, persist dry-run results, or emit notifications from this audit",
			"claim production D-Bus ownership, start services, launch compatibility engines, expose paths, or mutate host root",
		},
		NextRequirements: []string{
			"Add a redacted dry-run result visibility audit before any dry-run execution can be exposed to KDE.",
			"Require explicit operator dispatch approval and an accepted opaque authorization receipt before dry-run execution.",
			"Keep dry-run execution and result persistence disabled until a separate executor implementation is reviewed.",
			"Keep KDE surfaces read-only until dry-run results are redacted and Runtime-owned.",
		},
		DesktopSafeSummary: "The production receipt notification action dispatch dry-run audit models the dry-run boundary for future review, renew, open Compatibility Center, dismiss, and support-info dispatch actions, but it executes no dry runs, grants no dispatch authorization, creates no request objects, sends no Portal requests, opens no surfaces, writes no receipts or dry-run results, starts no services, launches no engines, and mutates no host state.",
	}
	preview.ExecutedDryRunItemCount = productionReceiptNotificationActionDispatchDryRunExecutedCount(items)
	preview.DispatchedDryRunItemCount = productionReceiptNotificationActionDispatchDryRunDispatchedCount(items)
	preview.CreatedRequestObjectCount = productionReceiptNotificationActionDispatchDryRunCreatedCount(items)
	preview.PortalRequestCreatedCount = productionReceiptNotificationActionDispatchDryRunPortalCount(items)
	preview.SideEffectDryRunItemCount = productionReceiptNotificationActionDispatchDryRunSideEffectCount(items)
	checks := productionReceiptNotificationActionDispatchDryRunAuditChecks(preview)
	preview.Checks = checks
	preview.CheckIDs = productionReceiptNotificationActionDispatchDryRunCheckIDs(checks)
	preview.Counts = countProductionReceiptNotificationActionDispatchDryRunChecks(checks)
	if preview.Counts.Blocked == 0 {
		preview.AuditDecision = "production-receipt-notification-action-dispatch-dry-run-audit-ready-dry-run-only"
	}
	if err := validateNoBackendTerms(preview, "production receipt notification action dispatch dry-run audit preview"); err != nil {
		return ProductionReceiptNotificationActionDispatchDryRunAuditPreview{}, err
	}
	return preview, nil
}

type productionReceiptNotificationActionDispatchDryRunAuditSourceSet struct {
	DispatchAuthorizationAudit string
	DispatchSheet              string
}

func productionReceiptNotificationActionDispatchDryRunAuditSources(root string) productionReceiptNotificationActionDispatchDryRunAuditSourceSet {
	return productionReceiptNotificationActionDispatchDryRunAuditSourceSet{
		DispatchAuthorizationAudit: productionAuthorizationReadSources(root, []string{"internal/runtime/owner/production_receipt_notification_action_dispatch_authorization_audit.go"}),
		DispatchSheet:              productionAuthorizationReadSources(root, []string{"docs/claude-code-current-dispatch-picks.md"}),
	}
}

func productionReceiptNotificationActionDispatchDryRunItems(sources productionReceiptNotificationActionDispatchDryRunAuditSourceSet) []ProductionReceiptNotificationActionDispatchDryRunAuditItem {
	combined := sources.DispatchAuthorizationAudit + sources.DispatchSheet
	return []ProductionReceiptNotificationActionDispatchDryRunAuditItem{
		productionReceiptNotificationActionDispatchDryRunItem("review-receipt-dispatch-dry-run", "review", "receipt-review-request", "review-request-dispatch-authorization", "review-request-dispatch-dry-run", "future review request dispatch dry-run requires authorization evidence", combined, []string{"review-receipt-dispatch-authorization", "receipt-review-request", "dispatch dry-run"}, "define review dispatch dry-run result visibility before execution"),
		productionReceiptNotificationActionDispatchDryRunItem("renew-receipt-dispatch-dry-run", "renew", "receipt-renewal-request", "renewal-request-dispatch-authorization", "renewal-request-dispatch-dry-run", "future renewal request dispatch dry-run requires authorization evidence", combined, []string{"renew-receipt-dispatch-authorization", "receipt-renewal-request", "dispatch dry-run"}, "define renewal dispatch dry-run result visibility before execution"),
		productionReceiptNotificationActionDispatchDryRunItem("open-compatibility-center-dispatch-dry-run", "open-compatibility-center", "compatibility-center-navigation-request", "navigation-request-dispatch-authorization", "navigation-request-dispatch-dry-run", "future Compatibility Center navigation dry-run requires authorization evidence", combined, []string{"open-compatibility-center-dispatch-authorization", "compatibility-center-navigation-request", "dispatch dry-run"}, "define navigation dispatch dry-run result visibility before execution"),
		productionReceiptNotificationActionDispatchDryRunItem("dismiss-receipt-dispatch-dry-run", "dismiss", "receipt-dismissal-request", "dismissal-request-dispatch-authorization", "dismissal-request-dispatch-dry-run", "future dismissal request dispatch dry-run requires authorization evidence", combined, []string{"dismiss-receipt-dispatch-authorization", "receipt-dismissal-request", "dispatch dry-run"}, "define dismissal dispatch dry-run result visibility before execution"),
		productionReceiptNotificationActionDispatchDryRunItem("support-info-dispatch-dry-run", "support-info", "support-info-request", "support-info-request-dispatch-authorization", "support-info-request-dispatch-dry-run", "future support-info request dispatch dry-run requires authorization evidence", combined, []string{"support-info-dispatch-authorization", "support-info-request", "dispatch dry-run"}, "define support-info dispatch dry-run result visibility before execution"),
	}
}

func productionReceiptNotificationActionDispatchDryRunItem(id string, actionKind string, requestKind string, authorizationKind string, dryRunKind string, evidence string, source string, tokens []string, nextRequirement string) ProductionReceiptNotificationActionDispatchDryRunAuditItem {
	ready := productionAuthorizationHasAll(source, tokens)
	status := "missing-dispatch-dry-run-evidence"
	if ready {
		status = "dispatch-dry-run-modeled-execution-disabled"
	}
	return ProductionReceiptNotificationActionDispatchDryRunAuditItem{
		ID:                           id,
		ActionKind:                   actionKind,
		RequestObjectKind:            requestKind,
		DispatchAuthorizationKind:    authorizationKind,
		DryRunKind:                   dryRunKind,
		RequiredEvidence:             evidence,
		EvidencePresent:              ready,
		DispatchDryRunModeled:        ready,
		UserVisible:                  ready,
		ReviewOnly:                   true,
		RuntimeOwned:                 true,
		GoRuntimeBacked:              true,
		KDEPolicyOwner:               false,
		OperatorApprovalRequired:     true,
		OperatorApprovalPresent:      false,
		DispatchAuthorizationGranted: false,
		DispatchDryRunExecuted:       false,
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
		SupportBundleExported:        false,
		SupportCaseCreated:           false,
		ProductionReadiness:          false,
		ProductionOwnershipReady:     false,
		SideEffectsDisabled:          true,
		HostRootModified:             false,
		InternalDetailsExposed:       false,
		DryRunStatus:                 status,
		NextRequirement:              nextRequirement,
	}
}

func productionReceiptNotificationActionDispatchDryRunAuditChecks(preview ProductionReceiptNotificationActionDispatchDryRunAuditPreview) []ProductionReceiptNotificationActionDispatchDryRunAuditCheck {
	return []ProductionReceiptNotificationActionDispatchDryRunAuditCheck{
		productionReceiptNotificationActionDispatchDryRunCheck("dispatch-authorization-audit-consumed", productionAuthorizationPassBlocked(preview.DispatchAuthorizationAuditConsumed), "The dry-run audit consumes dispatch authorization audit evidence."),
		productionReceiptNotificationActionDispatchDryRunCheck("dispatch-dry-run-guidance-consumed", productionAuthorizationPassBlocked(preview.DispatchDryRunGuidanceConsumed), "The dry-run audit consumes current dispatch guidance for dry-run-only boundaries."),
		productionReceiptNotificationActionDispatchDryRunCheck("dispatch-dry-run-modeled-only", productionAuthorizationPassBlocked(preview.DispatchDryRunAuditRequired && preview.DispatchDryRunAuditModeled && preview.DispatchDryRunPlanReady && preview.OperatorDispatchApprovalRequired && !preview.OperatorDispatchApprovalPresent && !preview.DispatchAuthorizationGranted && !preview.DispatchDryRunExecuted && !preview.ReceiptPresent && !preview.ReceiptAccepted && !preview.AuthorizationAccepted), "Dispatch dry-run is modeled without approval, authorization grant, execution, accepted receipt, or persistence."),
		productionReceiptNotificationActionDispatchDryRunCheck("five-dispatch-dry-run-items-present", productionAuthorizationPassBlocked(preview.DryRunItemCount == 5 && preview.RequiredDryRunItemCount == 5 && preview.MissingDryRunItemCount == 0), "Review, renew, open Compatibility Center, dismiss, and support-info dispatch dry-run items are modeled."),
		productionReceiptNotificationActionDispatchDryRunCheck("dispatch-dry-run-items-ready-execution-disabled", productionAuthorizationPassBlocked(preview.ReadyDryRunItemCount == 5 && productionReceiptNotificationActionDispatchDryRunItemsReady(preview.DryRunItems)), "Every dispatch dry-run item is ready while execution, grant, request creation, dispatch, Portal requests, and navigation remain disabled."),
		productionReceiptNotificationActionDispatchDryRunCheck("dry-run-execution-disabled", productionAuthorizationPassBlocked(!preview.DispatchDryRunExecutionEnabled && !preview.DispatchDryRunExecuted && !preview.DryRunResultPersisted && preview.ExecutedDryRunItemCount == 0 && preview.DispatchedDryRunItemCount == 0), "Dry-run execution, per-item execution, request dispatch, and dry-run result persistence remain disabled."),
		productionReceiptNotificationActionDispatchDryRunCheck("request-creation-dispatch-actions-and-portal-disabled", productionAuthorizationPassBlocked(!preview.RequestObjectCreationEnabled && !preview.RequestObjectDispatchEnabled && !preview.RequestObjectPersistenceEnabled && !preview.RequestObjectsCreated && !preview.RequestObjectsDispatched && !preview.PortalRequestCreated && !preview.NotificationActionEnabled && !preview.ReviewActionEnabled && !preview.RenewActionEnabled && !preview.OpenCompatibilityCenterEnabled && !preview.DismissActionEnabled && !preview.SupportInfoActionEnabled && preview.CreatedRequestObjectCount == 0 && preview.PortalRequestCreatedCount == 0), "Request creation, dispatch, persistence, notification actions, and Portal requests remain disabled."),
		productionReceiptNotificationActionDispatchDryRunCheck("receipt-notification-navigation-and-support-writes-disabled", productionAuthorizationPassBlocked(!preview.NotificationSent && !preview.NotificationDeliveryEnabled && !preview.CompatibilityCenterOpened && !preview.CompatibilityCenterPersisted && !preview.ReceiptWriterEnabled && !preview.ReceiptPersistenceEnabled && !preview.ReceiptLookupWritesEnabled && !preview.ReceiptReplayEnabled && !preview.ReceiptExpiryWriteEnabled && !preview.ReceiptRevocationWriteEnabled && !preview.SupportBundleExported && !preview.SupportCaseCreated), "Notification, navigation, receipt, and support write side effects remain disabled."),
		productionReceiptNotificationActionDispatchDryRunCheck("production-and-host-boundary-closed", productionAuthorizationPassBlocked(!preview.SystemServiceStarted && !preview.SessionBusClaimed && !preview.ProductionBusClaimed && !preview.ProductionOwnerEnabled && !preview.ProductionActivationReady && !preview.ProductionReadiness && !preview.ProductionOwnershipReady && !preview.AdapterInvocationEnabled && !preview.BackendLaunchEnabled && !preview.BackendProcessStarted && !preview.FileContentRead && !preview.FilePathsExposed && !preview.NetworkRequired && !preview.HostRootModified && !preview.PrivilegedContainerRequired && !preview.StateRootPathExposed && !preview.RawCommandExposed && !preview.RawExecutableExposed && !preview.BackendDetailsExposed && productionReceiptNotificationActionDispatchDryRunItemsKeepHostClosed(preview.DryRunItems)), "Production ownership, backend launch, unsafe data exposure, network, privilege, path exposure, internal details, and host mutation remain disabled."),
	}
}

func productionReceiptNotificationActionDispatchDryRunCheck(id string, status string, summary string) ProductionReceiptNotificationActionDispatchDryRunAuditCheck {
	return ProductionReceiptNotificationActionDispatchDryRunAuditCheck{ID: id, Status: status, Summary: summary}
}

func productionReceiptNotificationActionDispatchDryRunAuthorizationReady(source string) bool {
	return productionAuthorizationHasAll(source, []string{
		"production-receipt-notification-action-dispatch-authorization-audit-preview",
		"receipt-notification-action-request-dispatch-authorization-audit",
		"production-receipt-notification-action-dispatch-authorization-audit-ready-dispatch-disabled",
		"review-receipt-dispatch-authorization",
		"renew-receipt-dispatch-authorization",
		"open-compatibility-center-dispatch-authorization",
		"dismiss-receipt-dispatch-authorization",
		"support-info-dispatch-authorization",
		"DispatchAuthorizationGranted",
	})
}

func productionReceiptNotificationActionDispatchDryRunGuidanceReady(source string) bool {
	return productionAuthorizationHasAll(source, []string{
		"production receipt notification action dispatch dry-run audit",
		"dispatch dry-run",
		"real dispatch",
		"request creation",
	})
}

func productionReceiptNotificationActionDispatchDryRunReadyCount(items []ProductionReceiptNotificationActionDispatchDryRunAuditItem) int {
	count := 0
	for _, item := range items {
		if item.EvidencePresent {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionDispatchDryRunMissingCount(items []ProductionReceiptNotificationActionDispatchDryRunAuditItem) int {
	count := 0
	for _, item := range items {
		if !item.EvidencePresent {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionDispatchDryRunExecutedCount(items []ProductionReceiptNotificationActionDispatchDryRunAuditItem) int {
	count := 0
	for _, item := range items {
		if item.DispatchDryRunExecuted {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionDispatchDryRunDispatchedCount(items []ProductionReceiptNotificationActionDispatchDryRunAuditItem) int {
	count := 0
	for _, item := range items {
		if item.RequestObjectDispatched {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionDispatchDryRunCreatedCount(items []ProductionReceiptNotificationActionDispatchDryRunAuditItem) int {
	count := 0
	for _, item := range items {
		if item.RequestObjectCreated {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionDispatchDryRunPortalCount(items []ProductionReceiptNotificationActionDispatchDryRunAuditItem) int {
	count := 0
	for _, item := range items {
		if item.PortalRequestCreated {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionDispatchDryRunSideEffectCount(items []ProductionReceiptNotificationActionDispatchDryRunAuditItem) int {
	count := 0
	for _, item := range items {
		if item.DispatchAuthorizationGranted || item.DispatchDryRunExecuted || item.RequestObjectCreated || item.RequestObjectDispatched || item.RequestObjectPersisted || item.PortalRequestCreated || item.NavigationRequested || item.NotificationActionEnabled || item.ActionEnabled || item.ReceiptAccepted || item.AuthorizationAccepted || item.ReceiptWriterEnabled || item.ReceiptPersistenceEnabled || item.CompatibilityCenterOpened || item.CompatibilityCenterPersisted || item.SupportBundleExported || item.SupportCaseCreated || item.HostRootModified || item.InternalDetailsExposed {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionDispatchDryRunItemIDs(items []ProductionReceiptNotificationActionDispatchDryRunAuditItem) []string {
	ids := make([]string, 0, len(items))
	for _, item := range items {
		ids = append(ids, item.ID)
	}
	return ids
}

func productionReceiptNotificationActionDispatchDryRunCheckIDs(checks []ProductionReceiptNotificationActionDispatchDryRunAuditCheck) []string {
	ids := make([]string, 0, len(checks))
	for _, check := range checks {
		ids = append(ids, check.ID)
	}
	return ids
}

func countProductionReceiptNotificationActionDispatchDryRunChecks(checks []ProductionReceiptNotificationActionDispatchDryRunAuditCheck) ProductionReceiptNotificationActionDispatchDryRunAuditCounts {
	counts := ProductionReceiptNotificationActionDispatchDryRunAuditCounts{Total: len(checks)}
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

func productionReceiptNotificationActionDispatchDryRunItemsReady(items []ProductionReceiptNotificationActionDispatchDryRunAuditItem) bool {
	if len(items) != 5 {
		return false
	}
	for _, item := range items {
		if !item.EvidencePresent || !item.DispatchDryRunModeled || !item.UserVisible || item.DryRunStatus != "dispatch-dry-run-modeled-execution-disabled" {
			return false
		}
		if item.OperatorApprovalPresent || item.DispatchAuthorizationGranted || item.DispatchDryRunExecuted || item.RequestObjectCreated || item.RequestObjectDispatched || item.RequestObjectPersisted || item.PortalRequestCreated || item.NavigationRequested || item.NotificationActionEnabled || item.ActionEnabled || item.ReceiptAccepted || item.AuthorizationAccepted || item.ReceiptWriterEnabled || item.ReceiptPersistenceEnabled || item.CompatibilityCenterOpened || item.CompatibilityCenterPersisted || item.SupportBundleExported || item.SupportCaseCreated || item.ProductionReadiness || item.ProductionOwnershipReady || item.KDEPolicyOwner || !item.ReviewOnly || !item.SideEffectsDisabled {
			return false
		}
	}
	return true
}

func productionReceiptNotificationActionDispatchDryRunItemsKeepHostClosed(items []ProductionReceiptNotificationActionDispatchDryRunAuditItem) bool {
	for _, item := range items {
		if item.HostRootModified || item.InternalDetailsExposed {
			return false
		}
	}
	return true
}
