package owner

type ProductionReceiptNotificationActionRequestObjectAuditPreview struct {
	Version                         string                                                       `json:"version"`
	SchemaVersion                   string                                                       `json:"schema_version"`
	RequestType                     string                                                       `json:"request_type"`
	AuditType                       string                                                       `json:"audit_type"`
	Source                          string                                                       `json:"source"`
	AuditDecision                   string                                                       `json:"audit_decision"`
	ReceiptSchema                   string                                                       `json:"receipt_schema"`
	OpaqueReceiptID                 string                                                       `json:"opaque_receipt_id"`
	RequestObjectAuditRequired      bool                                                         `json:"request_object_audit_required"`
	RequestObjectAuditModeled       bool                                                         `json:"request_object_audit_modeled"`
	ActionSafetyAuditConsumed       bool                                                         `json:"action_safety_audit_consumed"`
	DispatchRequestBoundaryConsumed bool                                                         `json:"dispatch_request_boundary_consumed"`
	OperatorActionApprovalRequired  bool                                                         `json:"operator_action_approval_required"`
	OperatorActionApprovalPresent   bool                                                         `json:"operator_action_approval_present"`
	RequestObjectBoundaryReady      bool                                                         `json:"request_object_boundary_ready"`
	CallerStateRootRequired         bool                                                         `json:"caller_state_root_required"`
	ReceiptPresent                  bool                                                         `json:"receipt_present"`
	ReceiptAccepted                 bool                                                         `json:"receipt_accepted"`
	AuthorizationAccepted           bool                                                         `json:"authorization_accepted"`
	ProductionReadiness             bool                                                         `json:"production_readiness"`
	ProductionOwnershipReady        bool                                                         `json:"production_ownership_ready"`
	RequestObjectCount              int                                                          `json:"request_object_count"`
	RequiredRequestObjectCount      int                                                          `json:"required_request_object_count"`
	ReadyRequestObjectCount         int                                                          `json:"ready_request_object_count"`
	MissingRequestObjectCount       int                                                          `json:"missing_request_object_count"`
	CreatedRequestObjectCount       int                                                          `json:"created_request_object_count"`
	DispatchedRequestObjectCount    int                                                          `json:"dispatched_request_object_count"`
	PortalRequestCreatedCount       int                                                          `json:"portal_request_created_count"`
	NavigationRequestedCount        int                                                          `json:"navigation_requested_count"`
	SideEffectRequestObjectCount    int                                                          `json:"side_effect_request_object_count"`
	RequestObjects                  []ProductionReceiptNotificationActionRequestObjectAuditItem  `json:"request_objects"`
	RequestObjectIDs                []string                                                     `json:"request_object_ids"`
	RequiredBeforeRequestCreation   []string                                                     `json:"required_before_request_creation"`
	Checks                          []ProductionReceiptNotificationActionRequestObjectAuditCheck `json:"checks"`
	CheckIDs                        []string                                                     `json:"check_ids"`
	Counts                          ProductionReceiptNotificationActionRequestObjectAuditCounts  `json:"counts"`
	RuntimeOwned                    bool                                                         `json:"runtime_owned"`
	GoRuntimeBacked                 bool                                                         `json:"go_runtime_backed"`
	KDEPolicyOwner                  bool                                                         `json:"kde_policy_owner"`
	OfficialDesktopOnly             bool                                                         `json:"official_desktop_only"`
	PlasmaForkRequired              bool                                                         `json:"plasma_fork_required"`
	PlasmaSourceModified            bool                                                         `json:"plasma_source_modified"`
	SystemServiceStarted            bool                                                         `json:"system_service_started"`
	SessionBusClaimed               bool                                                         `json:"session_bus_claimed"`
	ProductionBusClaimed            bool                                                         `json:"production_bus_claimed"`
	ProductionOwnerEnabled          bool                                                         `json:"production_owner_enabled"`
	ProductionActivationReady       bool                                                         `json:"production_activation_ready"`
	WriteMethodsEnabled             bool                                                         `json:"write_methods_enabled"`
	RuntimeWritesEnabled            bool                                                         `json:"runtime_writes_enabled"`
	RequestObjectCreationEnabled    bool                                                         `json:"request_object_creation_enabled"`
	RequestObjectDispatchEnabled    bool                                                         `json:"request_object_dispatch_enabled"`
	RequestObjectPersistenceEnabled bool                                                         `json:"request_object_persistence_enabled"`
	PortalRequestCreated            bool                                                         `json:"portal_request_created"`
	RequestObjectsCreated           bool                                                         `json:"request_objects_created"`
	RequestObjectsDispatched        bool                                                         `json:"request_objects_dispatched"`
	NotificationSent                bool                                                         `json:"notification_sent"`
	NotificationDeliveryEnabled     bool                                                         `json:"notification_delivery_enabled"`
	NotificationActionEnabled       bool                                                         `json:"notification_action_enabled"`
	ReviewActionEnabled             bool                                                         `json:"review_action_enabled"`
	RenewActionEnabled              bool                                                         `json:"renew_action_enabled"`
	OpenCompatibilityCenterEnabled  bool                                                         `json:"open_compatibility_center_enabled"`
	DismissActionEnabled            bool                                                         `json:"dismiss_action_enabled"`
	SupportInfoActionEnabled        bool                                                         `json:"support_info_action_enabled"`
	CompatibilityCenterOpened       bool                                                         `json:"compatibility_center_opened"`
	CompatibilityCenterPersisted    bool                                                         `json:"compatibility_center_persisted"`
	ReceiptWriterEnabled            bool                                                         `json:"receipt_writer_enabled"`
	ReceiptPersistenceEnabled       bool                                                         `json:"receipt_persistence_enabled"`
	ReceiptLookupWritesEnabled      bool                                                         `json:"receipt_lookup_writes_enabled"`
	ReceiptReplayEnabled            bool                                                         `json:"receipt_replay_enabled"`
	ReceiptExpiryWriteEnabled       bool                                                         `json:"receipt_expiry_write_enabled"`
	ReceiptRevocationWriteEnabled   bool                                                         `json:"receipt_revocation_write_enabled"`
	DesktopFilesWritten             bool                                                         `json:"desktop_files_written"`
	MIMEAppsWritten                 bool                                                         `json:"mimeapps_written"`
	ShellConfigurationWritten       bool                                                         `json:"shell_configuration_written"`
	SettingsPersisted               bool                                                         `json:"settings_persisted"`
	KRunnerIndexPersisted           bool                                                         `json:"krunner_index_persisted"`
	TaskManagerEntryActive          bool                                                         `json:"task_manager_entry_active"`
	KWinRuleApplied                 bool                                                         `json:"kwin_rule_applied"`
	LiveTrayBridgeEnabled           bool                                                         `json:"live_tray_bridge_enabled"`
	TrayBridgePersisted             bool                                                         `json:"tray_bridge_persisted"`
	AdapterInvocationEnabled        bool                                                         `json:"adapter_invocation_enabled"`
	BackendLaunchEnabled            bool                                                         `json:"backend_launch_enabled"`
	BackendProcessStarted           bool                                                         `json:"backend_process_started"`
	SupportBundleExported           bool                                                         `json:"support_bundle_exported"`
	SupportCaseCreated              bool                                                         `json:"support_case_created"`
	SnapshotRestoreExecuted         bool                                                         `json:"snapshot_restore_executed"`
	StateCleanupExecuted            bool                                                         `json:"state_cleanup_executed"`
	FileContentRead                 bool                                                         `json:"file_content_read"`
	FilePathsExposed                bool                                                         `json:"file_paths_exposed"`
	NetworkRequired                 bool                                                         `json:"network_required"`
	HostRootModified                bool                                                         `json:"host_root_modified"`
	PrivilegedContainerRequired     bool                                                         `json:"privileged_container_required"`
	StateRootPathExposed            bool                                                         `json:"state_root_path_exposed"`
	RawCommandExposed               bool                                                         `json:"raw_command_exposed"`
	RawExecutableExposed            bool                                                         `json:"raw_executable_exposed"`
	BackendDetailsExposed           bool                                                         `json:"backend_details_exposed"`
	BlockedActions                  []string                                                     `json:"blocked_actions"`
	NextRequirements                []string                                                     `json:"next_requirements"`
	DesktopSafeSummary              string                                                       `json:"desktop_safe_summary"`
}

type ProductionReceiptNotificationActionRequestObjectAuditItem struct {
	ID                           string `json:"id"`
	ActionKind                   string `json:"action_kind"`
	RequestObjectKind            string `json:"request_object_kind"`
	RequiredEvidence             string `json:"required_evidence"`
	EvidencePresent              bool   `json:"evidence_present"`
	RequestObjectModeled         bool   `json:"request_object_modeled"`
	UserVisible                  bool   `json:"user_visible"`
	ReviewOnly                   bool   `json:"review_only"`
	RuntimeOwned                 bool   `json:"runtime_owned"`
	GoRuntimeBacked              bool   `json:"go_runtime_backed"`
	KDEPolicyOwner               bool   `json:"kde_policy_owner"`
	OperatorApprovalRequired     bool   `json:"operator_approval_required"`
	OperatorApprovalPresent      bool   `json:"operator_approval_present"`
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
	RequestObjectStatus          string `json:"request_object_status"`
	NextRequirement              string `json:"next_requirement"`
}

type ProductionReceiptNotificationActionRequestObjectAuditCheck struct {
	ID      string `json:"id"`
	Status  string `json:"status"`
	Summary string `json:"summary"`
}

type ProductionReceiptNotificationActionRequestObjectAuditCounts struct {
	Total   int `json:"total"`
	Passed  int `json:"passed"`
	Pending int `json:"pending"`
	Blocked int `json:"blocked"`
}

func NewProductionReceiptNotificationActionRequestObjectAuditPreview(root string) (ProductionReceiptNotificationActionRequestObjectAuditPreview, error) {
	version, err := readProductionDBusGateReviewVersion(root)
	if err != nil {
		return ProductionReceiptNotificationActionRequestObjectAuditPreview{}, err
	}
	sources := productionReceiptNotificationActionRequestObjectAuditSources(root)
	objects := productionReceiptNotificationActionRequestObjects(sources)
	actionSafetyReady := productionReceiptNotificationActionRequestObjectActionSafetyReady(sources.ActionSafetyAudit)
	dispatchBoundaryReady := productionReceiptNotificationActionRequestObjectDispatchBoundaryReady(sources.DispatchSheet + sources.ImplementationBrief)
	boundaryReady := actionSafetyReady && dispatchBoundaryReady && productionReceiptNotificationActionRequestObjectsReady(objects)
	preview := ProductionReceiptNotificationActionRequestObjectAuditPreview{
		Version:                         version,
		SchemaVersion:                   "xnix.runtime.production_receipt_notification_action_request_object_audit.v1",
		RequestType:                     "production-receipt-notification-action-request-object-audit-preview",
		AuditType:                       "receipt-notification-action-runtime-request-object-audit",
		Source:                          "production-receipt-notification-action-safety-audit-preview+claude-code-current-dispatch-picks+windows-app-compatibility-implementation-brief",
		AuditDecision:                   "production-receipt-notification-action-request-object-audit-blocked",
		ReceiptSchema:                   "xnix.runtime.production_dbus_human_authorization_receipt.v1",
		OpaqueReceiptID:                 ProductionDBusHumanAuthorizationReceiptID,
		RequestObjectAuditRequired:      true,
		RequestObjectAuditModeled:       true,
		ActionSafetyAuditConsumed:       actionSafetyReady,
		DispatchRequestBoundaryConsumed: dispatchBoundaryReady,
		OperatorActionApprovalRequired:  true,
		OperatorActionApprovalPresent:   false,
		RequestObjectBoundaryReady:      boundaryReady,
		CallerStateRootRequired:         false,
		ReceiptPresent:                  false,
		ReceiptAccepted:                 false,
		AuthorizationAccepted:           false,
		ProductionReadiness:             false,
		ProductionOwnershipReady:        false,
		RequestObjectCount:              len(objects),
		RequiredRequestObjectCount:      5,
		ReadyRequestObjectCount:         productionReceiptNotificationActionRequestObjectReadyCount(objects),
		MissingRequestObjectCount:       productionReceiptNotificationActionRequestObjectMissingCount(objects),
		RequestObjects:                  objects,
		RequestObjectIDs:                productionReceiptNotificationActionRequestObjectIDs(objects),
		RequiredBeforeRequestCreation: []string{
			"production-receipt-notification-action-safety-audit-preview",
			"explicit operator action approval",
			"separate request-object creation implementation",
			"separate request-object dispatch authorization",
			"separate KDE navigation and support route implementation",
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
			"treat request-object audit as permission to create, persist, or dispatch Runtime request objects",
			"enable review, renew, open Compatibility Center, dismiss, or support-info actions from this audit",
			"create Portal requests, navigate KDE surfaces, send notifications, or write receipt state from this audit",
			"claim production D-Bus ownership, start services, launch compatibility engines, expose paths, or mutate host root",
		},
		NextRequirements: []string{
			"Require explicit operator action approval before any notification action can create a request object.",
			"Add a separate request-object creation implementation before review or renewal actions can exist.",
			"Add separate dispatch authorization before request objects can leave this read-only model.",
			"Add separate KDE navigation and support routes before request objects can open surfaces or support flows.",
		},
		DesktopSafeSummary: "The production receipt notification action request-object audit maps review, renew, open Compatibility Center, dismiss, and support-info actions to future Runtime request-object kinds, but it creates no request objects, dispatches no requests, opens no surfaces, sends no notifications, writes no receipts, starts no services, launches no engines, and mutates no host state.",
	}
	preview.CreatedRequestObjectCount = productionReceiptNotificationActionRequestObjectCreatedCount(objects)
	preview.DispatchedRequestObjectCount = productionReceiptNotificationActionRequestObjectDispatchedCount(objects)
	preview.PortalRequestCreatedCount = productionReceiptNotificationActionRequestObjectPortalCount(objects)
	preview.NavigationRequestedCount = productionReceiptNotificationActionRequestObjectNavigationCount(objects)
	preview.SideEffectRequestObjectCount = productionReceiptNotificationActionRequestObjectSideEffectCount(objects)
	checks := productionReceiptNotificationActionRequestObjectAuditChecks(preview)
	preview.Checks = checks
	preview.CheckIDs = productionReceiptNotificationActionRequestObjectCheckIDs(checks)
	preview.Counts = countProductionReceiptNotificationActionRequestObjectChecks(checks)
	if preview.Counts.Blocked == 0 {
		preview.AuditDecision = "production-receipt-notification-action-request-object-audit-ready-requests-disabled"
	}
	if err := validateNoBackendTerms(preview, "production receipt notification action request object audit preview"); err != nil {
		return ProductionReceiptNotificationActionRequestObjectAuditPreview{}, err
	}
	return preview, nil
}

type productionReceiptNotificationActionRequestObjectAuditSourceSet struct {
	ActionSafetyAudit   string
	DispatchSheet       string
	ImplementationBrief string
}

func productionReceiptNotificationActionRequestObjectAuditSources(root string) productionReceiptNotificationActionRequestObjectAuditSourceSet {
	return productionReceiptNotificationActionRequestObjectAuditSourceSet{
		ActionSafetyAudit:   productionAuthorizationReadSources(root, []string{"internal/runtime/owner/production_receipt_notification_action_safety_audit.go"}),
		DispatchSheet:       productionAuthorizationReadSources(root, []string{"docs/claude-code-current-dispatch-picks.md"}),
		ImplementationBrief: productionAuthorizationReadSources(root, []string{"docs/windows-app-compatibility-implementation-brief.md"}),
	}
}

func productionReceiptNotificationActionRequestObjects(sources productionReceiptNotificationActionRequestObjectAuditSourceSet) []ProductionReceiptNotificationActionRequestObjectAuditItem {
	combined := sources.ActionSafetyAudit + sources.DispatchSheet + sources.ImplementationBrief
	return []ProductionReceiptNotificationActionRequestObjectAuditItem{
		productionReceiptNotificationActionRequestObject("review-receipt-request-object", "review", "receipt-review-request", "future review action requires a Runtime request object", combined, []string{"review", "request-object", "RequestObjectsCreated"}, "define review request-object creation before action enablement"),
		productionReceiptNotificationActionRequestObject("renew-receipt-request-object", "renew", "receipt-renewal-request", "future renew action requires a Runtime request object", combined, []string{"renew", "request-object", "operator"}, "define renewal request-object creation before action enablement"),
		productionReceiptNotificationActionRequestObject("open-compatibility-center-request-object", "open-compatibility-center", "compatibility-center-navigation-request", "future open Compatibility Center action requires a navigation request object", combined, []string{"navigation", "Compatibility Center", "RequestObjectsCreated"}, "define navigation request-object creation before action enablement"),
		productionReceiptNotificationActionRequestObject("dismiss-receipt-request-object", "dismiss", "receipt-dismissal-request", "future dismiss action requires a persistence-safe request object", combined, []string{"dismiss", "request creation", "NotificationActionEnabled"}, "define dismissal request-object creation before action enablement"),
		productionReceiptNotificationActionRequestObject("support-info-request-object", "support-info", "support-info-request", "future support info action requires a support request object", combined, []string{"support requests", "support", "SupportCaseCreated"}, "define support request-object creation before action enablement"),
	}
}

func productionReceiptNotificationActionRequestObject(id string, actionKind string, requestKind string, evidence string, source string, tokens []string, nextRequirement string) ProductionReceiptNotificationActionRequestObjectAuditItem {
	ready := productionAuthorizationHasAll(source, tokens)
	status := "missing-request-object-evidence"
	if ready {
		status = "request-object-modeled-creation-disabled"
	}
	return ProductionReceiptNotificationActionRequestObjectAuditItem{
		ID:                           id,
		ActionKind:                   actionKind,
		RequestObjectKind:            requestKind,
		RequiredEvidence:             evidence,
		EvidencePresent:              ready,
		RequestObjectModeled:         ready,
		UserVisible:                  ready,
		ReviewOnly:                   true,
		RuntimeOwned:                 true,
		GoRuntimeBacked:              true,
		KDEPolicyOwner:               false,
		OperatorApprovalRequired:     true,
		OperatorApprovalPresent:      false,
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
		RequestObjectStatus:          status,
		NextRequirement:              nextRequirement,
	}
}

func productionReceiptNotificationActionRequestObjectAuditChecks(preview ProductionReceiptNotificationActionRequestObjectAuditPreview) []ProductionReceiptNotificationActionRequestObjectAuditCheck {
	return []ProductionReceiptNotificationActionRequestObjectAuditCheck{
		productionReceiptNotificationActionRequestObjectCheck("action-safety-audit-consumed", productionAuthorizationPassBlocked(preview.ActionSafetyAuditConsumed), "The request-object audit consumes notification action safety evidence."),
		productionReceiptNotificationActionRequestObjectCheck("dispatch-request-boundary-consumed", productionAuthorizationPassBlocked(preview.DispatchRequestBoundaryConsumed), "The request-object audit consumes dispatch guidance for request-object boundaries."),
		productionReceiptNotificationActionRequestObjectCheck("request-object-modeled-only", productionAuthorizationPassBlocked(preview.RequestObjectAuditRequired && preview.RequestObjectAuditModeled && preview.RequestObjectBoundaryReady && preview.OperatorActionApprovalRequired && !preview.OperatorActionApprovalPresent && !preview.ReceiptPresent && !preview.ReceiptAccepted && !preview.AuthorizationAccepted), "Request objects are modeled without creation, dispatch, approval, or persistence."),
		productionReceiptNotificationActionRequestObjectCheck("five-request-objects-present", productionAuthorizationPassBlocked(preview.RequestObjectCount == 5 && preview.RequiredRequestObjectCount == 5 && preview.MissingRequestObjectCount == 0), "Review, renew, open Compatibility Center, dismiss, and support-info request objects are modeled."),
		productionReceiptNotificationActionRequestObjectCheck("request-objects-ready-creation-disabled", productionAuthorizationPassBlocked(preview.ReadyRequestObjectCount == 5 && productionReceiptNotificationActionRequestObjectsReady(preview.RequestObjects)), "Every request object is ready while creation, dispatch, Portal requests, and navigation remain disabled."),
		productionReceiptNotificationActionRequestObjectCheck("request-creation-disabled", productionAuthorizationPassBlocked(!preview.RequestObjectCreationEnabled && !preview.RequestObjectDispatchEnabled && !preview.RequestObjectPersistenceEnabled && !preview.RequestObjectsCreated && !preview.RequestObjectsDispatched && preview.CreatedRequestObjectCount == 0 && preview.DispatchedRequestObjectCount == 0), "Runtime request-object creation, dispatch, and persistence remain disabled."),
		productionReceiptNotificationActionRequestObjectCheck("actions-navigation-and-portal-disabled", productionAuthorizationPassBlocked(!preview.NotificationActionEnabled && !preview.ReviewActionEnabled && !preview.RenewActionEnabled && !preview.OpenCompatibilityCenterEnabled && !preview.DismissActionEnabled && !preview.SupportInfoActionEnabled && !preview.PortalRequestCreated && !preview.CompatibilityCenterOpened && !preview.CompatibilityCenterPersisted && preview.PortalRequestCreatedCount == 0 && preview.NavigationRequestedCount == 0), "Notification actions, Portal requests, Compatibility Center navigation, and persistence remain disabled."),
		productionReceiptNotificationActionRequestObjectCheck("receipt-notification-and-support-writes-disabled", productionAuthorizationPassBlocked(!preview.NotificationSent && !preview.NotificationDeliveryEnabled && !preview.ReceiptWriterEnabled && !preview.ReceiptPersistenceEnabled && !preview.ReceiptLookupWritesEnabled && !preview.ReceiptReplayEnabled && !preview.ReceiptExpiryWriteEnabled && !preview.ReceiptRevocationWriteEnabled && !preview.SupportBundleExported && !preview.SupportCaseCreated), "Notification, receipt, and support write side effects remain disabled."),
		productionReceiptNotificationActionRequestObjectCheck("production-and-host-boundary-closed", productionAuthorizationPassBlocked(!preview.SystemServiceStarted && !preview.SessionBusClaimed && !preview.ProductionBusClaimed && !preview.ProductionOwnerEnabled && !preview.ProductionActivationReady && !preview.ProductionReadiness && !preview.ProductionOwnershipReady && !preview.AdapterInvocationEnabled && !preview.BackendLaunchEnabled && !preview.BackendProcessStarted && !preview.FileContentRead && !preview.FilePathsExposed && !preview.NetworkRequired && !preview.HostRootModified && !preview.PrivilegedContainerRequired && !preview.StateRootPathExposed && !preview.RawCommandExposed && !preview.RawExecutableExposed && !preview.BackendDetailsExposed && productionReceiptNotificationActionRequestObjectsKeepHostClosed(preview.RequestObjects)), "Production ownership, backend launch, unsafe data exposure, network, privilege, path exposure, internal details, and host mutation remain disabled."),
	}
}

func productionReceiptNotificationActionRequestObjectCheck(id string, status string, summary string) ProductionReceiptNotificationActionRequestObjectAuditCheck {
	return ProductionReceiptNotificationActionRequestObjectAuditCheck{ID: id, Status: status, Summary: summary}
}

func productionReceiptNotificationActionRequestObjectActionSafetyReady(source string) bool {
	return productionAuthorizationHasAll(source, []string{
		"production-receipt-notification-action-safety-audit-preview",
		"receipt-notification-action-request-safety-audit",
		"production-receipt-notification-action-safety-audit-ready-actions-disabled",
		"review-receipt-action",
		"renew-receipt-action",
		"open-compatibility-center-action",
		"dismiss-receipt-action",
		"support-info-action",
		"RequestObjectsCreated",
	})
}

func productionReceiptNotificationActionRequestObjectDispatchBoundaryReady(source string) bool {
	return productionAuthorizationHasAll(source, []string{
		"production receipt notification action request-object audit",
		"request-object",
		"request creation",
		"review",
		"renew",
		"navigation",
		"support requests",
	})
}

func productionReceiptNotificationActionRequestObjectReadyCount(objects []ProductionReceiptNotificationActionRequestObjectAuditItem) int {
	count := 0
	for _, object := range objects {
		if object.EvidencePresent {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionRequestObjectMissingCount(objects []ProductionReceiptNotificationActionRequestObjectAuditItem) int {
	count := 0
	for _, object := range objects {
		if !object.EvidencePresent {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionRequestObjectCreatedCount(objects []ProductionReceiptNotificationActionRequestObjectAuditItem) int {
	count := 0
	for _, object := range objects {
		if object.RequestObjectCreated {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionRequestObjectDispatchedCount(objects []ProductionReceiptNotificationActionRequestObjectAuditItem) int {
	count := 0
	for _, object := range objects {
		if object.RequestObjectDispatched {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionRequestObjectPortalCount(objects []ProductionReceiptNotificationActionRequestObjectAuditItem) int {
	count := 0
	for _, object := range objects {
		if object.PortalRequestCreated {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionRequestObjectNavigationCount(objects []ProductionReceiptNotificationActionRequestObjectAuditItem) int {
	count := 0
	for _, object := range objects {
		if object.NavigationRequested || object.CompatibilityCenterOpened {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionRequestObjectSideEffectCount(objects []ProductionReceiptNotificationActionRequestObjectAuditItem) int {
	count := 0
	for _, object := range objects {
		if object.RequestObjectCreated || object.RequestObjectDispatched || object.RequestObjectPersisted || object.PortalRequestCreated || object.NavigationRequested || object.NotificationActionEnabled || object.ActionEnabled || object.ReceiptAccepted || object.AuthorizationAccepted || object.ReceiptWriterEnabled || object.ReceiptPersistenceEnabled || object.CompatibilityCenterOpened || object.CompatibilityCenterPersisted || object.SupportBundleExported || object.SupportCaseCreated || object.HostRootModified || object.InternalDetailsExposed {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionRequestObjectIDs(objects []ProductionReceiptNotificationActionRequestObjectAuditItem) []string {
	ids := make([]string, 0, len(objects))
	for _, object := range objects {
		ids = append(ids, object.ID)
	}
	return ids
}

func productionReceiptNotificationActionRequestObjectCheckIDs(checks []ProductionReceiptNotificationActionRequestObjectAuditCheck) []string {
	ids := make([]string, 0, len(checks))
	for _, check := range checks {
		ids = append(ids, check.ID)
	}
	return ids
}

func countProductionReceiptNotificationActionRequestObjectChecks(checks []ProductionReceiptNotificationActionRequestObjectAuditCheck) ProductionReceiptNotificationActionRequestObjectAuditCounts {
	counts := ProductionReceiptNotificationActionRequestObjectAuditCounts{Total: len(checks)}
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

func productionReceiptNotificationActionRequestObjectsReady(objects []ProductionReceiptNotificationActionRequestObjectAuditItem) bool {
	if len(objects) != 5 {
		return false
	}
	for _, object := range objects {
		if !object.EvidencePresent || !object.RequestObjectModeled || !object.UserVisible || object.RequestObjectStatus != "request-object-modeled-creation-disabled" {
			return false
		}
		if object.OperatorApprovalPresent || object.RequestObjectCreated || object.RequestObjectDispatched || object.RequestObjectPersisted || object.PortalRequestCreated || object.NavigationRequested || object.NotificationActionEnabled || object.ActionEnabled || object.ReceiptAccepted || object.AuthorizationAccepted || object.ReceiptWriterEnabled || object.ReceiptPersistenceEnabled || object.CompatibilityCenterOpened || object.CompatibilityCenterPersisted || object.SupportBundleExported || object.SupportCaseCreated || object.ProductionReadiness || object.ProductionOwnershipReady || object.KDEPolicyOwner || !object.ReviewOnly || !object.SideEffectsDisabled {
			return false
		}
	}
	return true
}

func productionReceiptNotificationActionRequestObjectsKeepHostClosed(objects []ProductionReceiptNotificationActionRequestObjectAuditItem) bool {
	for _, object := range objects {
		if object.HostRootModified || object.InternalDetailsExposed {
			return false
		}
	}
	return true
}
