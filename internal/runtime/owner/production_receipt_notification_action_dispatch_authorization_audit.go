package owner

type ProductionReceiptNotificationActionDispatchAuthorizationAuditPreview struct {
	Version                              string                                                               `json:"version"`
	SchemaVersion                        string                                                               `json:"schema_version"`
	RequestType                          string                                                               `json:"request_type"`
	AuditType                            string                                                               `json:"audit_type"`
	Source                               string                                                               `json:"source"`
	AuditDecision                        string                                                               `json:"audit_decision"`
	ReceiptSchema                        string                                                               `json:"receipt_schema"`
	OpaqueReceiptID                      string                                                               `json:"opaque_receipt_id"`
	DispatchAuthorizationAuditRequired   bool                                                                 `json:"dispatch_authorization_audit_required"`
	DispatchAuthorizationAuditModeled    bool                                                                 `json:"dispatch_authorization_audit_modeled"`
	RequestObjectAuditConsumed           bool                                                                 `json:"request_object_audit_consumed"`
	ReceiptAuthorizationBoundaryConsumed bool                                                                 `json:"receipt_authorization_boundary_consumed"`
	OperatorDispatchApprovalRequired     bool                                                                 `json:"operator_dispatch_approval_required"`
	OperatorDispatchApprovalPresent      bool                                                                 `json:"operator_dispatch_approval_present"`
	DispatchAuthorizationReady           bool                                                                 `json:"dispatch_authorization_ready"`
	DispatchAuthorizationGranted         bool                                                                 `json:"dispatch_authorization_granted"`
	CallerStateRootRequired              bool                                                                 `json:"caller_state_root_required"`
	ReceiptPresent                       bool                                                                 `json:"receipt_present"`
	ReceiptAccepted                      bool                                                                 `json:"receipt_accepted"`
	AuthorizationAccepted                bool                                                                 `json:"authorization_accepted"`
	ProductionReadiness                  bool                                                                 `json:"production_readiness"`
	ProductionOwnershipReady             bool                                                                 `json:"production_ownership_ready"`
	AuthorizationItemCount               int                                                                  `json:"authorization_item_count"`
	RequiredAuthorizationItemCount       int                                                                  `json:"required_authorization_item_count"`
	ReadyAuthorizationItemCount          int                                                                  `json:"ready_authorization_item_count"`
	MissingAuthorizationItemCount        int                                                                  `json:"missing_authorization_item_count"`
	GrantedAuthorizationItemCount        int                                                                  `json:"granted_authorization_item_count"`
	DispatchedAuthorizationItemCount     int                                                                  `json:"dispatched_authorization_item_count"`
	CreatedRequestObjectCount            int                                                                  `json:"created_request_object_count"`
	PortalRequestCreatedCount            int                                                                  `json:"portal_request_created_count"`
	SideEffectAuthorizationItemCount     int                                                                  `json:"side_effect_authorization_item_count"`
	AuthorizationItems                   []ProductionReceiptNotificationActionDispatchAuthorizationAuditItem  `json:"authorization_items"`
	AuthorizationItemIDs                 []string                                                             `json:"authorization_item_ids"`
	RequiredBeforeDispatch               []string                                                             `json:"required_before_dispatch"`
	Checks                               []ProductionReceiptNotificationActionDispatchAuthorizationAuditCheck `json:"checks"`
	CheckIDs                             []string                                                             `json:"check_ids"`
	Counts                               ProductionReceiptNotificationActionDispatchAuthorizationAuditCounts  `json:"counts"`
	RuntimeOwned                         bool                                                                 `json:"runtime_owned"`
	GoRuntimeBacked                      bool                                                                 `json:"go_runtime_backed"`
	KDEPolicyOwner                       bool                                                                 `json:"kde_policy_owner"`
	OfficialDesktopOnly                  bool                                                                 `json:"official_desktop_only"`
	PlasmaForkRequired                   bool                                                                 `json:"plasma_fork_required"`
	PlasmaSourceModified                 bool                                                                 `json:"plasma_source_modified"`
	SystemServiceStarted                 bool                                                                 `json:"system_service_started"`
	SessionBusClaimed                    bool                                                                 `json:"session_bus_claimed"`
	ProductionBusClaimed                 bool                                                                 `json:"production_bus_claimed"`
	ProductionOwnerEnabled               bool                                                                 `json:"production_owner_enabled"`
	ProductionActivationReady            bool                                                                 `json:"production_activation_ready"`
	WriteMethodsEnabled                  bool                                                                 `json:"write_methods_enabled"`
	RuntimeWritesEnabled                 bool                                                                 `json:"runtime_writes_enabled"`
	RequestObjectCreationEnabled         bool                                                                 `json:"request_object_creation_enabled"`
	RequestObjectDispatchEnabled         bool                                                                 `json:"request_object_dispatch_enabled"`
	RequestObjectPersistenceEnabled      bool                                                                 `json:"request_object_persistence_enabled"`
	DispatchAuthorizationPersisted       bool                                                                 `json:"dispatch_authorization_persisted"`
	PortalRequestCreated                 bool                                                                 `json:"portal_request_created"`
	RequestObjectsCreated                bool                                                                 `json:"request_objects_created"`
	RequestObjectsDispatched             bool                                                                 `json:"request_objects_dispatched"`
	NotificationSent                     bool                                                                 `json:"notification_sent"`
	NotificationDeliveryEnabled          bool                                                                 `json:"notification_delivery_enabled"`
	NotificationActionEnabled            bool                                                                 `json:"notification_action_enabled"`
	ReviewActionEnabled                  bool                                                                 `json:"review_action_enabled"`
	RenewActionEnabled                   bool                                                                 `json:"renew_action_enabled"`
	OpenCompatibilityCenterEnabled       bool                                                                 `json:"open_compatibility_center_enabled"`
	DismissActionEnabled                 bool                                                                 `json:"dismiss_action_enabled"`
	SupportInfoActionEnabled             bool                                                                 `json:"support_info_action_enabled"`
	CompatibilityCenterOpened            bool                                                                 `json:"compatibility_center_opened"`
	CompatibilityCenterPersisted         bool                                                                 `json:"compatibility_center_persisted"`
	ReceiptWriterEnabled                 bool                                                                 `json:"receipt_writer_enabled"`
	ReceiptPersistenceEnabled            bool                                                                 `json:"receipt_persistence_enabled"`
	ReceiptLookupWritesEnabled           bool                                                                 `json:"receipt_lookup_writes_enabled"`
	ReceiptReplayEnabled                 bool                                                                 `json:"receipt_replay_enabled"`
	ReceiptExpiryWriteEnabled            bool                                                                 `json:"receipt_expiry_write_enabled"`
	ReceiptRevocationWriteEnabled        bool                                                                 `json:"receipt_revocation_write_enabled"`
	DesktopFilesWritten                  bool                                                                 `json:"desktop_files_written"`
	MIMEAppsWritten                      bool                                                                 `json:"mimeapps_written"`
	ShellConfigurationWritten            bool                                                                 `json:"shell_configuration_written"`
	SettingsPersisted                    bool                                                                 `json:"settings_persisted"`
	KRunnerIndexPersisted                bool                                                                 `json:"krunner_index_persisted"`
	TaskManagerEntryActive               bool                                                                 `json:"task_manager_entry_active"`
	KWinRuleApplied                      bool                                                                 `json:"kwin_rule_applied"`
	LiveTrayBridgeEnabled                bool                                                                 `json:"live_tray_bridge_enabled"`
	TrayBridgePersisted                  bool                                                                 `json:"tray_bridge_persisted"`
	AdapterInvocationEnabled             bool                                                                 `json:"adapter_invocation_enabled"`
	BackendLaunchEnabled                 bool                                                                 `json:"backend_launch_enabled"`
	BackendProcessStarted                bool                                                                 `json:"backend_process_started"`
	SupportBundleExported                bool                                                                 `json:"support_bundle_exported"`
	SupportCaseCreated                   bool                                                                 `json:"support_case_created"`
	SnapshotRestoreExecuted              bool                                                                 `json:"snapshot_restore_executed"`
	StateCleanupExecuted                 bool                                                                 `json:"state_cleanup_executed"`
	FileContentRead                      bool                                                                 `json:"file_content_read"`
	FilePathsExposed                     bool                                                                 `json:"file_paths_exposed"`
	NetworkRequired                      bool                                                                 `json:"network_required"`
	HostRootModified                     bool                                                                 `json:"host_root_modified"`
	PrivilegedContainerRequired          bool                                                                 `json:"privileged_container_required"`
	StateRootPathExposed                 bool                                                                 `json:"state_root_path_exposed"`
	RawCommandExposed                    bool                                                                 `json:"raw_command_exposed"`
	RawExecutableExposed                 bool                                                                 `json:"raw_executable_exposed"`
	BackendDetailsExposed                bool                                                                 `json:"backend_details_exposed"`
	BlockedActions                       []string                                                             `json:"blocked_actions"`
	NextRequirements                     []string                                                             `json:"next_requirements"`
	DesktopSafeSummary                   string                                                               `json:"desktop_safe_summary"`
}

type ProductionReceiptNotificationActionDispatchAuthorizationAuditItem struct {
	ID                           string `json:"id"`
	ActionKind                   string `json:"action_kind"`
	RequestObjectKind            string `json:"request_object_kind"`
	DispatchAuthorizationKind    string `json:"dispatch_authorization_kind"`
	RequiredEvidence             string `json:"required_evidence"`
	EvidencePresent              bool   `json:"evidence_present"`
	DispatchAuthorizationModeled bool   `json:"dispatch_authorization_modeled"`
	UserVisible                  bool   `json:"user_visible"`
	ReviewOnly                   bool   `json:"review_only"`
	RuntimeOwned                 bool   `json:"runtime_owned"`
	GoRuntimeBacked              bool   `json:"go_runtime_backed"`
	KDEPolicyOwner               bool   `json:"kde_policy_owner"`
	OperatorApprovalRequired     bool   `json:"operator_approval_required"`
	OperatorApprovalPresent      bool   `json:"operator_approval_present"`
	DispatchAuthorizationGranted bool   `json:"dispatch_authorization_granted"`
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
	DispatchAuthorizationStatus  string `json:"dispatch_authorization_status"`
	NextRequirement              string `json:"next_requirement"`
}

type ProductionReceiptNotificationActionDispatchAuthorizationAuditCheck struct {
	ID      string `json:"id"`
	Status  string `json:"status"`
	Summary string `json:"summary"`
}

type ProductionReceiptNotificationActionDispatchAuthorizationAuditCounts struct {
	Total   int `json:"total"`
	Passed  int `json:"passed"`
	Pending int `json:"pending"`
	Blocked int `json:"blocked"`
}

func NewProductionReceiptNotificationActionDispatchAuthorizationAuditPreview(root string) (ProductionReceiptNotificationActionDispatchAuthorizationAuditPreview, error) {
	version, err := readProductionDBusGateReviewVersion(root)
	if err != nil {
		return ProductionReceiptNotificationActionDispatchAuthorizationAuditPreview{}, err
	}
	sources := productionReceiptNotificationActionDispatchAuthorizationAuditSources(root)
	items := productionReceiptNotificationActionDispatchAuthorizationItems(sources)
	requestObjectReady := productionReceiptNotificationActionDispatchAuthorizationRequestObjectReady(sources.RequestObjectAudit)
	receiptBoundaryReady := productionReceiptNotificationActionDispatchAuthorizationReceiptBoundaryReady(sources.AuthorizationConsumptionAudit + sources.DispatchSheet)
	dispatchReady := requestObjectReady && receiptBoundaryReady && productionReceiptNotificationActionDispatchAuthorizationItemsReady(items)
	preview := ProductionReceiptNotificationActionDispatchAuthorizationAuditPreview{
		Version:                              version,
		SchemaVersion:                        "xnix.runtime.production_receipt_notification_action_dispatch_authorization_audit.v1",
		RequestType:                          "production-receipt-notification-action-dispatch-authorization-audit-preview",
		AuditType:                            "receipt-notification-action-request-dispatch-authorization-audit",
		Source:                               "production-receipt-notification-action-request-object-audit-preview+production-authorization-consumption-audit-preview+claude-code-current-dispatch-picks",
		AuditDecision:                        "production-receipt-notification-action-dispatch-authorization-audit-blocked",
		ReceiptSchema:                        "xnix.runtime.production_dbus_human_authorization_receipt.v1",
		OpaqueReceiptID:                      ProductionDBusHumanAuthorizationReceiptID,
		DispatchAuthorizationAuditRequired:   true,
		DispatchAuthorizationAuditModeled:    true,
		RequestObjectAuditConsumed:           requestObjectReady,
		ReceiptAuthorizationBoundaryConsumed: receiptBoundaryReady,
		OperatorDispatchApprovalRequired:     true,
		OperatorDispatchApprovalPresent:      false,
		DispatchAuthorizationReady:           dispatchReady,
		DispatchAuthorizationGranted:         false,
		CallerStateRootRequired:              false,
		ReceiptPresent:                       false,
		ReceiptAccepted:                      false,
		AuthorizationAccepted:                false,
		ProductionReadiness:                  false,
		ProductionOwnershipReady:             false,
		AuthorizationItemCount:               len(items),
		RequiredAuthorizationItemCount:       5,
		ReadyAuthorizationItemCount:          productionReceiptNotificationActionDispatchAuthorizationReadyCount(items),
		MissingAuthorizationItemCount:        productionReceiptNotificationActionDispatchAuthorizationMissingCount(items),
		AuthorizationItems:                   items,
		AuthorizationItemIDs:                 productionReceiptNotificationActionDispatchAuthorizationItemIDs(items),
		RequiredBeforeDispatch: []string{
			"production-receipt-notification-action-request-object-audit-preview",
			"production-authorization-consumption-audit-preview",
			"explicit operator dispatch approval",
			"accepted opaque authorization receipt",
			"separate dispatch implementation",
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
			"treat dispatch authorization audit as permission to dispatch Runtime request objects",
			"create, persist, or dispatch notification action request objects from this audit",
			"enable notification actions, Portal requests, KDE navigation, support cases, or receipt writes from this audit",
			"claim production D-Bus ownership, start services, launch compatibility engines, expose paths, or mutate host root",
		},
		NextRequirements: []string{
			"Require explicit operator dispatch approval before any request object can be dispatched.",
			"Require an accepted opaque authorization receipt before dispatch can be granted.",
			"Add a separate dispatch implementation after this audit models the authorization boundary.",
			"Keep KDE surfaces read-only until dispatch produces redacted user-safe results.",
		},
		DesktopSafeSummary: "The production receipt notification action dispatch authorization audit models the authorization boundary for dispatching review, renew, open Compatibility Center, dismiss, and support-info request objects, but it grants no authorization, dispatches no requests, opens no surfaces, sends no notifications, writes no receipts, starts no services, launches no engines, and mutates no host state.",
	}
	preview.GrantedAuthorizationItemCount = productionReceiptNotificationActionDispatchAuthorizationGrantedCount(items)
	preview.DispatchedAuthorizationItemCount = productionReceiptNotificationActionDispatchAuthorizationDispatchedCount(items)
	preview.CreatedRequestObjectCount = productionReceiptNotificationActionDispatchAuthorizationCreatedCount(items)
	preview.PortalRequestCreatedCount = productionReceiptNotificationActionDispatchAuthorizationPortalCount(items)
	preview.SideEffectAuthorizationItemCount = productionReceiptNotificationActionDispatchAuthorizationSideEffectCount(items)
	checks := productionReceiptNotificationActionDispatchAuthorizationAuditChecks(preview)
	preview.Checks = checks
	preview.CheckIDs = productionReceiptNotificationActionDispatchAuthorizationCheckIDs(checks)
	preview.Counts = countProductionReceiptNotificationActionDispatchAuthorizationChecks(checks)
	if preview.Counts.Blocked == 0 {
		preview.AuditDecision = "production-receipt-notification-action-dispatch-authorization-audit-ready-dispatch-disabled"
	}
	if err := validateNoBackendTerms(preview, "production receipt notification action dispatch authorization audit preview"); err != nil {
		return ProductionReceiptNotificationActionDispatchAuthorizationAuditPreview{}, err
	}
	return preview, nil
}

type productionReceiptNotificationActionDispatchAuthorizationAuditSourceSet struct {
	RequestObjectAudit            string
	AuthorizationConsumptionAudit string
	DispatchSheet                 string
}

func productionReceiptNotificationActionDispatchAuthorizationAuditSources(root string) productionReceiptNotificationActionDispatchAuthorizationAuditSourceSet {
	return productionReceiptNotificationActionDispatchAuthorizationAuditSourceSet{
		RequestObjectAudit:            productionAuthorizationReadSources(root, []string{"internal/runtime/owner/production_receipt_notification_action_request_object_audit.go"}),
		AuthorizationConsumptionAudit: productionAuthorizationReadSources(root, []string{"internal/runtime/owner/production_authorization_consumption_audit.go"}),
		DispatchSheet:                 productionAuthorizationReadSources(root, []string{"docs/claude-code-current-dispatch-picks.md"}),
	}
}

func productionReceiptNotificationActionDispatchAuthorizationItems(sources productionReceiptNotificationActionDispatchAuthorizationAuditSourceSet) []ProductionReceiptNotificationActionDispatchAuthorizationAuditItem {
	combined := sources.RequestObjectAudit + sources.AuthorizationConsumptionAudit + sources.DispatchSheet
	return []ProductionReceiptNotificationActionDispatchAuthorizationAuditItem{
		productionReceiptNotificationActionDispatchAuthorizationItem("review-receipt-dispatch-authorization", "review", "receipt-review-request", "review-request-dispatch-authorization", "future review request dispatch requires authorization", combined, []string{"review-receipt-request-object", "receipt-review-request", "dispatch authorization"}, "define review dispatch authorization before dispatch enablement"),
		productionReceiptNotificationActionDispatchAuthorizationItem("renew-receipt-dispatch-authorization", "renew", "receipt-renewal-request", "renewal-request-dispatch-authorization", "future renewal request dispatch requires authorization", combined, []string{"renew-receipt-request-object", "receipt-renewal-request", "operator"}, "define renewal dispatch authorization before dispatch enablement"),
		productionReceiptNotificationActionDispatchAuthorizationItem("open-compatibility-center-dispatch-authorization", "open-compatibility-center", "compatibility-center-navigation-request", "navigation-request-dispatch-authorization", "future Compatibility Center navigation dispatch requires authorization", combined, []string{"open-compatibility-center-request-object", "compatibility-center-navigation-request", "navigation"}, "define navigation dispatch authorization before dispatch enablement"),
		productionReceiptNotificationActionDispatchAuthorizationItem("dismiss-receipt-dispatch-authorization", "dismiss", "receipt-dismissal-request", "dismissal-request-dispatch-authorization", "future dismissal request dispatch requires authorization", combined, []string{"dismiss-receipt-request-object", "receipt-dismissal-request", "request creation"}, "define dismissal dispatch authorization before dispatch enablement"),
		productionReceiptNotificationActionDispatchAuthorizationItem("support-info-dispatch-authorization", "support-info", "support-info-request", "support-info-request-dispatch-authorization", "future support-info request dispatch requires authorization", combined, []string{"support-info-request-object", "support-info-request", "support requests"}, "define support-info dispatch authorization before dispatch enablement"),
	}
}

func productionReceiptNotificationActionDispatchAuthorizationItem(id string, actionKind string, requestKind string, authorizationKind string, evidence string, source string, tokens []string, nextRequirement string) ProductionReceiptNotificationActionDispatchAuthorizationAuditItem {
	ready := productionAuthorizationHasAll(source, tokens)
	status := "missing-dispatch-authorization-evidence"
	if ready {
		status = "dispatch-authorization-modeled-dispatch-disabled"
	}
	return ProductionReceiptNotificationActionDispatchAuthorizationAuditItem{
		ID:                           id,
		ActionKind:                   actionKind,
		RequestObjectKind:            requestKind,
		DispatchAuthorizationKind:    authorizationKind,
		RequiredEvidence:             evidence,
		EvidencePresent:              ready,
		DispatchAuthorizationModeled: ready,
		UserVisible:                  ready,
		ReviewOnly:                   true,
		RuntimeOwned:                 true,
		GoRuntimeBacked:              true,
		KDEPolicyOwner:               false,
		OperatorApprovalRequired:     true,
		OperatorApprovalPresent:      false,
		DispatchAuthorizationGranted: false,
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
		DispatchAuthorizationStatus:  status,
		NextRequirement:              nextRequirement,
	}
}

func productionReceiptNotificationActionDispatchAuthorizationAuditChecks(preview ProductionReceiptNotificationActionDispatchAuthorizationAuditPreview) []ProductionReceiptNotificationActionDispatchAuthorizationAuditCheck {
	return []ProductionReceiptNotificationActionDispatchAuthorizationAuditCheck{
		productionReceiptNotificationActionDispatchAuthorizationCheck("request-object-audit-consumed", productionAuthorizationPassBlocked(preview.RequestObjectAuditConsumed), "The dispatch authorization audit consumes notification action request-object evidence."),
		productionReceiptNotificationActionDispatchAuthorizationCheck("receipt-authorization-boundary-consumed", productionAuthorizationPassBlocked(preview.ReceiptAuthorizationBoundaryConsumed), "The dispatch authorization audit consumes the consolidated receipt authorization boundary."),
		productionReceiptNotificationActionDispatchAuthorizationCheck("dispatch-authorization-modeled-only", productionAuthorizationPassBlocked(preview.DispatchAuthorizationAuditRequired && preview.DispatchAuthorizationAuditModeled && preview.DispatchAuthorizationReady && preview.OperatorDispatchApprovalRequired && !preview.OperatorDispatchApprovalPresent && !preview.DispatchAuthorizationGranted && !preview.ReceiptPresent && !preview.ReceiptAccepted && !preview.AuthorizationAccepted), "Dispatch authorization is modeled without approval, accepted receipt, dispatch, or persistence."),
		productionReceiptNotificationActionDispatchAuthorizationCheck("five-dispatch-authorizations-present", productionAuthorizationPassBlocked(preview.AuthorizationItemCount == 5 && preview.RequiredAuthorizationItemCount == 5 && preview.MissingAuthorizationItemCount == 0), "Review, renew, open Compatibility Center, dismiss, and support-info dispatch authorizations are modeled."),
		productionReceiptNotificationActionDispatchAuthorizationCheck("dispatch-authorizations-ready-dispatch-disabled", productionAuthorizationPassBlocked(preview.ReadyAuthorizationItemCount == 5 && productionReceiptNotificationActionDispatchAuthorizationItemsReady(preview.AuthorizationItems)), "Every dispatch authorization item is ready while grant, creation, dispatch, Portal requests, and navigation remain disabled."),
		productionReceiptNotificationActionDispatchAuthorizationCheck("dispatch-disabled", productionAuthorizationPassBlocked(!preview.DispatchAuthorizationGranted && !preview.RequestObjectDispatchEnabled && !preview.DispatchAuthorizationPersisted && !preview.RequestObjectsDispatched && preview.GrantedAuthorizationItemCount == 0 && preview.DispatchedAuthorizationItemCount == 0), "Dispatch grant, dispatch persistence, request-object dispatch, and per-item dispatch remain disabled."),
		productionReceiptNotificationActionDispatchAuthorizationCheck("request-creation-actions-and-portal-disabled", productionAuthorizationPassBlocked(!preview.RequestObjectCreationEnabled && !preview.RequestObjectPersistenceEnabled && !preview.RequestObjectsCreated && !preview.PortalRequestCreated && !preview.NotificationActionEnabled && !preview.ReviewActionEnabled && !preview.RenewActionEnabled && !preview.OpenCompatibilityCenterEnabled && !preview.DismissActionEnabled && !preview.SupportInfoActionEnabled && preview.CreatedRequestObjectCount == 0 && preview.PortalRequestCreatedCount == 0), "Request creation, persistence, notification actions, and Portal requests remain disabled."),
		productionReceiptNotificationActionDispatchAuthorizationCheck("receipt-notification-navigation-and-support-writes-disabled", productionAuthorizationPassBlocked(!preview.NotificationSent && !preview.NotificationDeliveryEnabled && !preview.CompatibilityCenterOpened && !preview.CompatibilityCenterPersisted && !preview.ReceiptWriterEnabled && !preview.ReceiptPersistenceEnabled && !preview.ReceiptLookupWritesEnabled && !preview.ReceiptReplayEnabled && !preview.ReceiptExpiryWriteEnabled && !preview.ReceiptRevocationWriteEnabled && !preview.SupportBundleExported && !preview.SupportCaseCreated), "Notification, navigation, receipt, and support write side effects remain disabled."),
		productionReceiptNotificationActionDispatchAuthorizationCheck("production-and-host-boundary-closed", productionAuthorizationPassBlocked(!preview.SystemServiceStarted && !preview.SessionBusClaimed && !preview.ProductionBusClaimed && !preview.ProductionOwnerEnabled && !preview.ProductionActivationReady && !preview.ProductionReadiness && !preview.ProductionOwnershipReady && !preview.AdapterInvocationEnabled && !preview.BackendLaunchEnabled && !preview.BackendProcessStarted && !preview.FileContentRead && !preview.FilePathsExposed && !preview.NetworkRequired && !preview.HostRootModified && !preview.PrivilegedContainerRequired && !preview.StateRootPathExposed && !preview.RawCommandExposed && !preview.RawExecutableExposed && !preview.BackendDetailsExposed && productionReceiptNotificationActionDispatchAuthorizationItemsKeepHostClosed(preview.AuthorizationItems)), "Production ownership, backend launch, unsafe data exposure, network, privilege, path exposure, internal details, and host mutation remain disabled."),
	}
}

func productionReceiptNotificationActionDispatchAuthorizationCheck(id string, status string, summary string) ProductionReceiptNotificationActionDispatchAuthorizationAuditCheck {
	return ProductionReceiptNotificationActionDispatchAuthorizationAuditCheck{ID: id, Status: status, Summary: summary}
}

func productionReceiptNotificationActionDispatchAuthorizationRequestObjectReady(source string) bool {
	return productionAuthorizationHasAll(source, []string{
		"production-receipt-notification-action-request-object-audit-preview",
		"receipt-notification-action-runtime-request-object-audit",
		"production-receipt-notification-action-request-object-audit-ready-requests-disabled",
		"review-receipt-request-object",
		"renew-receipt-request-object",
		"open-compatibility-center-request-object",
		"dismiss-receipt-request-object",
		"support-info-request-object",
		"RequestObjectsDispatched",
	})
}

func productionReceiptNotificationActionDispatchAuthorizationReceiptBoundaryReady(source string) bool {
	return productionAuthorizationHasAll(source, []string{
		"production-authorization-consumption-audit-preview",
		"production-authorization-consumption-audit-ready-authorization-disabled",
		"AuthorizationAccepted",
		"ReceiptWriterEnabled",
		"notification action dispatch authorization",
	})
}

func productionReceiptNotificationActionDispatchAuthorizationReadyCount(items []ProductionReceiptNotificationActionDispatchAuthorizationAuditItem) int {
	count := 0
	for _, item := range items {
		if item.EvidencePresent {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionDispatchAuthorizationMissingCount(items []ProductionReceiptNotificationActionDispatchAuthorizationAuditItem) int {
	count := 0
	for _, item := range items {
		if !item.EvidencePresent {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionDispatchAuthorizationGrantedCount(items []ProductionReceiptNotificationActionDispatchAuthorizationAuditItem) int {
	count := 0
	for _, item := range items {
		if item.DispatchAuthorizationGranted {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionDispatchAuthorizationDispatchedCount(items []ProductionReceiptNotificationActionDispatchAuthorizationAuditItem) int {
	count := 0
	for _, item := range items {
		if item.RequestObjectDispatched {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionDispatchAuthorizationCreatedCount(items []ProductionReceiptNotificationActionDispatchAuthorizationAuditItem) int {
	count := 0
	for _, item := range items {
		if item.RequestObjectCreated {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionDispatchAuthorizationPortalCount(items []ProductionReceiptNotificationActionDispatchAuthorizationAuditItem) int {
	count := 0
	for _, item := range items {
		if item.PortalRequestCreated {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionDispatchAuthorizationSideEffectCount(items []ProductionReceiptNotificationActionDispatchAuthorizationAuditItem) int {
	count := 0
	for _, item := range items {
		if item.DispatchAuthorizationGranted || item.RequestObjectCreated || item.RequestObjectDispatched || item.RequestObjectPersisted || item.PortalRequestCreated || item.NavigationRequested || item.NotificationActionEnabled || item.ActionEnabled || item.ReceiptAccepted || item.AuthorizationAccepted || item.ReceiptWriterEnabled || item.ReceiptPersistenceEnabled || item.CompatibilityCenterOpened || item.CompatibilityCenterPersisted || item.SupportBundleExported || item.SupportCaseCreated || item.HostRootModified || item.InternalDetailsExposed {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionDispatchAuthorizationItemIDs(items []ProductionReceiptNotificationActionDispatchAuthorizationAuditItem) []string {
	ids := make([]string, 0, len(items))
	for _, item := range items {
		ids = append(ids, item.ID)
	}
	return ids
}

func productionReceiptNotificationActionDispatchAuthorizationCheckIDs(checks []ProductionReceiptNotificationActionDispatchAuthorizationAuditCheck) []string {
	ids := make([]string, 0, len(checks))
	for _, check := range checks {
		ids = append(ids, check.ID)
	}
	return ids
}

func countProductionReceiptNotificationActionDispatchAuthorizationChecks(checks []ProductionReceiptNotificationActionDispatchAuthorizationAuditCheck) ProductionReceiptNotificationActionDispatchAuthorizationAuditCounts {
	counts := ProductionReceiptNotificationActionDispatchAuthorizationAuditCounts{Total: len(checks)}
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

func productionReceiptNotificationActionDispatchAuthorizationItemsReady(items []ProductionReceiptNotificationActionDispatchAuthorizationAuditItem) bool {
	if len(items) != 5 {
		return false
	}
	for _, item := range items {
		if !item.EvidencePresent || !item.DispatchAuthorizationModeled || !item.UserVisible || item.DispatchAuthorizationStatus != "dispatch-authorization-modeled-dispatch-disabled" {
			return false
		}
		if item.OperatorApprovalPresent || item.DispatchAuthorizationGranted || item.RequestObjectCreated || item.RequestObjectDispatched || item.RequestObjectPersisted || item.PortalRequestCreated || item.NavigationRequested || item.NotificationActionEnabled || item.ActionEnabled || item.ReceiptAccepted || item.AuthorizationAccepted || item.ReceiptWriterEnabled || item.ReceiptPersistenceEnabled || item.CompatibilityCenterOpened || item.CompatibilityCenterPersisted || item.SupportBundleExported || item.SupportCaseCreated || item.ProductionReadiness || item.ProductionOwnershipReady || item.KDEPolicyOwner || !item.ReviewOnly || !item.SideEffectsDisabled {
			return false
		}
	}
	return true
}

func productionReceiptNotificationActionDispatchAuthorizationItemsKeepHostClosed(items []ProductionReceiptNotificationActionDispatchAuthorizationAuditItem) bool {
	for _, item := range items {
		if item.HostRootModified || item.InternalDetailsExposed {
			return false
		}
	}
	return true
}
