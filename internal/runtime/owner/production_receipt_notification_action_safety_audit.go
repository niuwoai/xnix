package owner

type ProductionReceiptNotificationActionSafetyAuditPreview struct {
	Version                           string                                                `json:"version"`
	SchemaVersion                     string                                                `json:"schema_version"`
	RequestType                       string                                                `json:"request_type"`
	AuditType                         string                                                `json:"audit_type"`
	Source                            string                                                `json:"source"`
	AuditDecision                     string                                                `json:"audit_decision"`
	ReceiptSchema                     string                                                `json:"receipt_schema"`
	OpaqueReceiptID                   string                                                `json:"opaque_receipt_id"`
	ActionSafetyAuditRequired         bool                                                  `json:"action_safety_audit_required"`
	ActionSafetyAuditModeled          bool                                                  `json:"action_safety_audit_modeled"`
	NotificationDeliveryGateConsumed  bool                                                  `json:"notification_delivery_gate_consumed"`
	DesktopSideEffectReviewConsumed   bool                                                  `json:"desktop_side_effect_review_consumed"`
	OperatorActionApprovalRequired    bool                                                  `json:"operator_action_approval_required"`
	OperatorActionApprovalPresent     bool                                                  `json:"operator_action_approval_present"`
	NotificationActionSafetyReady     bool                                                  `json:"notification_action_safety_ready"`
	CallerStateRootRequired           bool                                                  `json:"caller_state_root_required"`
	ReceiptPresent                    bool                                                  `json:"receipt_present"`
	ReceiptAccepted                   bool                                                  `json:"receipt_accepted"`
	AuthorizationAccepted             bool                                                  `json:"authorization_accepted"`
	ProductionReadiness               bool                                                  `json:"production_readiness"`
	ProductionOwnershipReady          bool                                                  `json:"production_ownership_ready"`
	ActionItemCount                   int                                                   `json:"action_item_count"`
	RequiredActionItemCount           int                                                   `json:"required_action_item_count"`
	ReadyActionItemCount              int                                                   `json:"ready_action_item_count"`
	MissingActionItemCount            int                                                   `json:"missing_action_item_count"`
	EnabledActionItemCount            int                                                   `json:"enabled_action_item_count"`
	RequestCreatedItemCount           int                                                   `json:"request_created_item_count"`
	NavigationEnabledItemCount        int                                                   `json:"navigation_enabled_item_count"`
	SideEffectItemCount               int                                                   `json:"side_effect_item_count"`
	ActionItems                       []ProductionReceiptNotificationActionSafetyAuditItem  `json:"action_items"`
	ActionItemIDs                     []string                                              `json:"action_item_ids"`
	RequiredBeforeNotificationActions []string                                              `json:"required_before_notification_actions"`
	Checks                            []ProductionReceiptNotificationActionSafetyAuditCheck `json:"checks"`
	CheckIDs                          []string                                              `json:"check_ids"`
	Counts                            ProductionReceiptNotificationActionSafetyAuditCounts  `json:"counts"`
	RuntimeOwned                      bool                                                  `json:"runtime_owned"`
	GoRuntimeBacked                   bool                                                  `json:"go_runtime_backed"`
	KDEPolicyOwner                    bool                                                  `json:"kde_policy_owner"`
	OfficialDesktopOnly               bool                                                  `json:"official_desktop_only"`
	PlasmaForkRequired                bool                                                  `json:"plasma_fork_required"`
	PlasmaSourceModified              bool                                                  `json:"plasma_source_modified"`
	SystemServiceStarted              bool                                                  `json:"system_service_started"`
	SessionBusClaimed                 bool                                                  `json:"session_bus_claimed"`
	ProductionBusClaimed              bool                                                  `json:"production_bus_claimed"`
	ProductionOwnerEnabled            bool                                                  `json:"production_owner_enabled"`
	ProductionActivationReady         bool                                                  `json:"production_activation_ready"`
	WriteMethodsEnabled               bool                                                  `json:"write_methods_enabled"`
	RuntimeWritesEnabled              bool                                                  `json:"runtime_writes_enabled"`
	ReceiptWriterEnabled              bool                                                  `json:"receipt_writer_enabled"`
	ReceiptPersistenceEnabled         bool                                                  `json:"receipt_persistence_enabled"`
	ReceiptLookupWritesEnabled        bool                                                  `json:"receipt_lookup_writes_enabled"`
	ReceiptReplayEnabled              bool                                                  `json:"receipt_replay_enabled"`
	ReceiptExpiryWriteEnabled         bool                                                  `json:"receipt_expiry_write_enabled"`
	ReceiptRevocationWriteEnabled     bool                                                  `json:"receipt_revocation_write_enabled"`
	NotificationSent                  bool                                                  `json:"notification_sent"`
	NotificationDeliveryEnabled       bool                                                  `json:"notification_delivery_enabled"`
	NotificationActionEnabled         bool                                                  `json:"notification_action_enabled"`
	NotificationActionSafetyPersisted bool                                                  `json:"notification_action_safety_persisted"`
	ReviewActionEnabled               bool                                                  `json:"review_action_enabled"`
	RenewActionEnabled                bool                                                  `json:"renew_action_enabled"`
	OpenCompatibilityCenterEnabled    bool                                                  `json:"open_compatibility_center_enabled"`
	DismissActionEnabled              bool                                                  `json:"dismiss_action_enabled"`
	SupportInfoActionEnabled          bool                                                  `json:"support_info_action_enabled"`
	CompatibilityCenterOpened         bool                                                  `json:"compatibility_center_opened"`
	CompatibilityCenterPersisted      bool                                                  `json:"compatibility_center_persisted"`
	PortalRequestCreated              bool                                                  `json:"portal_request_created"`
	RequestObjectsCreated             bool                                                  `json:"request_objects_created"`
	DesktopFilesWritten               bool                                                  `json:"desktop_files_written"`
	MIMEAppsWritten                   bool                                                  `json:"mimeapps_written"`
	ShellConfigurationWritten         bool                                                  `json:"shell_configuration_written"`
	SettingsPersisted                 bool                                                  `json:"settings_persisted"`
	KRunnerIndexPersisted             bool                                                  `json:"krunner_index_persisted"`
	TaskManagerEntryActive            bool                                                  `json:"task_manager_entry_active"`
	KWinRuleApplied                   bool                                                  `json:"kwin_rule_applied"`
	LiveTrayBridgeEnabled             bool                                                  `json:"live_tray_bridge_enabled"`
	TrayBridgePersisted               bool                                                  `json:"tray_bridge_persisted"`
	AdapterInvocationEnabled          bool                                                  `json:"adapter_invocation_enabled"`
	BackendLaunchEnabled              bool                                                  `json:"backend_launch_enabled"`
	BackendProcessStarted             bool                                                  `json:"backend_process_started"`
	SupportBundleExported             bool                                                  `json:"support_bundle_exported"`
	SupportCaseCreated                bool                                                  `json:"support_case_created"`
	SnapshotRestoreExecuted           bool                                                  `json:"snapshot_restore_executed"`
	StateCleanupExecuted              bool                                                  `json:"state_cleanup_executed"`
	FileContentRead                   bool                                                  `json:"file_content_read"`
	FilePathsExposed                  bool                                                  `json:"file_paths_exposed"`
	NetworkRequired                   bool                                                  `json:"network_required"`
	HostRootModified                  bool                                                  `json:"host_root_modified"`
	PrivilegedContainerRequired       bool                                                  `json:"privileged_container_required"`
	StateRootPathExposed              bool                                                  `json:"state_root_path_exposed"`
	RawCommandExposed                 bool                                                  `json:"raw_command_exposed"`
	RawExecutableExposed              bool                                                  `json:"raw_executable_exposed"`
	BackendDetailsExposed             bool                                                  `json:"backend_details_exposed"`
	BlockedActions                    []string                                              `json:"blocked_actions"`
	NextRequirements                  []string                                              `json:"next_requirements"`
	DesktopSafeSummary                string                                                `json:"desktop_safe_summary"`
}

type ProductionReceiptNotificationActionSafetyAuditItem struct {
	ID                            string `json:"id"`
	ActionKind                    string `json:"action_kind"`
	RequiredEvidence              string `json:"required_evidence"`
	EvidencePresent               bool   `json:"evidence_present"`
	ActionSafetyModeled           bool   `json:"action_safety_modeled"`
	UserVisible                   bool   `json:"user_visible"`
	ReviewOnly                    bool   `json:"review_only"`
	RuntimeOwned                  bool   `json:"runtime_owned"`
	GoRuntimeBacked               bool   `json:"go_runtime_backed"`
	KDEPolicyOwner                bool   `json:"kde_policy_owner"`
	OperatorApprovalRequired      bool   `json:"operator_approval_required"`
	OperatorApprovalPresent       bool   `json:"operator_approval_present"`
	ActionEnabled                 bool   `json:"action_enabled"`
	NavigationEnabled             bool   `json:"navigation_enabled"`
	PortalRequestCreated          bool   `json:"portal_request_created"`
	RequestObjectsCreated         bool   `json:"request_objects_created"`
	NotificationSent              bool   `json:"notification_sent"`
	NotificationDeliveryEnabled   bool   `json:"notification_delivery_enabled"`
	NotificationActionEnabled     bool   `json:"notification_action_enabled"`
	ReceiptAccepted               bool   `json:"receipt_accepted"`
	AuthorizationAccepted         bool   `json:"authorization_accepted"`
	ReceiptRevocationWriteEnabled bool   `json:"receipt_revocation_write_enabled"`
	ReceiptExpiryWriteEnabled     bool   `json:"receipt_expiry_write_enabled"`
	ReceiptPersistenceEnabled     bool   `json:"receipt_persistence_enabled"`
	ReceiptLookupWritesEnabled    bool   `json:"receipt_lookup_writes_enabled"`
	CompatibilityCenterOpened     bool   `json:"compatibility_center_opened"`
	CompatibilityCenterPersisted  bool   `json:"compatibility_center_persisted"`
	SupportBundleExported         bool   `json:"support_bundle_exported"`
	SupportCaseCreated            bool   `json:"support_case_created"`
	ProductionReadiness           bool   `json:"production_readiness"`
	ProductionOwnershipReady      bool   `json:"production_ownership_ready"`
	SideEffectsDisabled           bool   `json:"side_effects_disabled"`
	HostRootModified              bool   `json:"host_root_modified"`
	InternalDetailsExposed        bool   `json:"internal_details_exposed"`
	ActionStatus                  string `json:"action_status"`
	NextRequirement               string `json:"next_requirement"`
}

type ProductionReceiptNotificationActionSafetyAuditCheck struct {
	ID      string `json:"id"`
	Status  string `json:"status"`
	Summary string `json:"summary"`
}

type ProductionReceiptNotificationActionSafetyAuditCounts struct {
	Total   int `json:"total"`
	Passed  int `json:"passed"`
	Pending int `json:"pending"`
	Blocked int `json:"blocked"`
}

func NewProductionReceiptNotificationActionSafetyAuditPreview(root string) (ProductionReceiptNotificationActionSafetyAuditPreview, error) {
	version, err := readProductionDBusGateReviewVersion(root)
	if err != nil {
		return ProductionReceiptNotificationActionSafetyAuditPreview{}, err
	}
	sources := productionReceiptNotificationActionSafetyAuditSources(root)
	items := productionReceiptNotificationActionSafetyAuditItems(sources)
	notificationGateReady := productionReceiptNotificationActionDeliveryGateReady(sources.NotificationDeliveryGateAudit)
	desktopReady := productionReceiptNotificationActionDesktopReady(sources.DesktopSideEffectReview)
	actionReady := notificationGateReady && desktopReady && productionReceiptNotificationActionItemsReady(items)
	preview := ProductionReceiptNotificationActionSafetyAuditPreview{
		Version:                          version,
		SchemaVersion:                    "xnix.runtime.production_receipt_notification_action_safety_audit.v1",
		RequestType:                      "production-receipt-notification-action-safety-audit-preview",
		AuditType:                        "receipt-notification-action-request-safety-audit",
		Source:                           "production-receipt-notification-delivery-gate-audit-preview+production-desktop-side-effect-review-preview+claude-code-current-dispatch-picks",
		AuditDecision:                    "production-receipt-notification-action-safety-audit-blocked",
		ReceiptSchema:                    "xnix.runtime.production_dbus_human_authorization_receipt.v1",
		OpaqueReceiptID:                  ProductionDBusHumanAuthorizationReceiptID,
		ActionSafetyAuditRequired:        true,
		ActionSafetyAuditModeled:         true,
		NotificationDeliveryGateConsumed: notificationGateReady,
		DesktopSideEffectReviewConsumed:  desktopReady,
		OperatorActionApprovalRequired:   true,
		OperatorActionApprovalPresent:    false,
		NotificationActionSafetyReady:    actionReady,
		CallerStateRootRequired:          false,
		ReceiptPresent:                   false,
		ReceiptAccepted:                  false,
		AuthorizationAccepted:            false,
		ProductionReadiness:              false,
		ProductionOwnershipReady:         false,
		ActionItemCount:                  len(items),
		RequiredActionItemCount:          5,
		ReadyActionItemCount:             productionReceiptNotificationActionReadyItemCount(items),
		MissingActionItemCount:           productionReceiptNotificationActionMissingItemCount(items),
		ActionItems:                      items,
		ActionItemIDs:                    productionReceiptNotificationActionItemIDs(items),
		RequiredBeforeNotificationActions: []string{
			"production-receipt-notification-delivery-gate-audit-preview",
			"production-desktop-side-effect-review-preview",
			"explicit operator notification action approval",
			"separate Runtime request-object implementation",
			"separate KDE Compatibility Center navigation implementation",
		},
		RuntimeOwned:                   true,
		GoRuntimeBacked:                true,
		KDEPolicyOwner:                 false,
		OfficialDesktopOnly:            true,
		PlasmaForkRequired:             false,
		PlasmaSourceModified:           false,
		SystemServiceStarted:           false,
		SessionBusClaimed:              false,
		ProductionBusClaimed:           false,
		ProductionOwnerEnabled:         false,
		ProductionActivationReady:      false,
		WriteMethodsEnabled:            false,
		RuntimeWritesEnabled:           false,
		ReceiptWriterEnabled:           false,
		ReceiptPersistenceEnabled:      false,
		ReceiptLookupWritesEnabled:     false,
		ReceiptReplayEnabled:           false,
		ReceiptExpiryWriteEnabled:      false,
		ReceiptRevocationWriteEnabled:  false,
		NotificationSent:               false,
		NotificationDeliveryEnabled:    false,
		NotificationActionEnabled:      false,
		ReviewActionEnabled:            false,
		RenewActionEnabled:             false,
		OpenCompatibilityCenterEnabled: false,
		DismissActionEnabled:           false,
		SupportInfoActionEnabled:       false,
		CompatibilityCenterOpened:      false,
		CompatibilityCenterPersisted:   false,
		PortalRequestCreated:           false,
		RequestObjectsCreated:          false,
		DesktopFilesWritten:            false,
		MIMEAppsWritten:                false,
		ShellConfigurationWritten:      false,
		SettingsPersisted:              false,
		KRunnerIndexPersisted:          false,
		TaskManagerEntryActive:         false,
		KWinRuleApplied:                false,
		LiveTrayBridgeEnabled:          false,
		TrayBridgePersisted:            false,
		AdapterInvocationEnabled:       false,
		BackendLaunchEnabled:           false,
		BackendProcessStarted:          false,
		SupportBundleExported:          false,
		SupportCaseCreated:             false,
		SnapshotRestoreExecuted:        false,
		StateCleanupExecuted:           false,
		FileContentRead:                false,
		FilePathsExposed:               false,
		NetworkRequired:                false,
		HostRootModified:               false,
		PrivilegedContainerRequired:    false,
		StateRootPathExposed:           false,
		RawCommandExposed:              false,
		RawExecutableExposed:           false,
		BackendDetailsExposed:          false,
		BlockedActions: []string{
			"treat notification action safety audit as permission to enable notification actions",
			"open Compatibility Center, create Portal requests, create Runtime request objects, or navigate from this audit",
			"enable review, renew, dismiss, support-info, or Compatibility Center actions from this audit",
			"send notifications, persist action safety state, write receipts, accept receipts, expire receipts, revoke receipts, or enable Runtime writes",
			"claim production D-Bus ownership, start services, launch compatibility engines, expose paths, or mutate host root",
		},
		NextRequirements: []string{
			"Require explicit operator action approval before notification buttons can create requests or navigate.",
			"Keep notification delivery and notification action enablement separate; both remain disabled in this audit.",
			"Require a separate Runtime request-object implementation before renew or review actions can exist.",
			"Require a separate KDE navigation implementation before notification actions can open Compatibility Center surfaces.",
		},
		DesktopSafeSummary: "The production receipt notification action safety audit models review, renew, open Compatibility Center, dismiss, and support-info notification actions, but it enables no actions, creates no requests, opens no surfaces, sends no notifications, writes no receipts, starts no services, launches no engines, and mutates no host state.",
	}
	preview.EnabledActionItemCount = productionReceiptNotificationActionEnabledCount(items)
	preview.RequestCreatedItemCount = productionReceiptNotificationActionRequestCount(items)
	preview.NavigationEnabledItemCount = productionReceiptNotificationActionNavigationCount(items)
	preview.SideEffectItemCount = productionReceiptNotificationActionSideEffectCount(items)
	checks := productionReceiptNotificationActionSafetyAuditChecks(preview)
	preview.Checks = checks
	preview.CheckIDs = productionReceiptNotificationActionCheckIDs(checks)
	preview.Counts = countProductionReceiptNotificationActionChecks(checks)
	if preview.Counts.Blocked == 0 {
		preview.AuditDecision = "production-receipt-notification-action-safety-audit-ready-actions-disabled"
	}
	if err := validateNoBackendTerms(preview, "production receipt notification action safety audit preview"); err != nil {
		return ProductionReceiptNotificationActionSafetyAuditPreview{}, err
	}
	return preview, nil
}

type productionReceiptNotificationActionSafetyAuditSourceSet struct {
	NotificationDeliveryGateAudit string
	DesktopSideEffectReview       string
	DispatchSheet                 string
}

func productionReceiptNotificationActionSafetyAuditSources(root string) productionReceiptNotificationActionSafetyAuditSourceSet {
	return productionReceiptNotificationActionSafetyAuditSourceSet{
		NotificationDeliveryGateAudit: productionAuthorizationReadSources(root, []string{"internal/runtime/owner/production_receipt_notification_delivery_gate_audit.go"}),
		DesktopSideEffectReview:       productionAuthorizationReadSources(root, []string{"internal/runtime/owner/production_desktop_side_effect_review.go"}),
		DispatchSheet:                 productionAuthorizationReadSources(root, []string{"docs/claude-code-current-dispatch-picks.md"}),
	}
}

func productionReceiptNotificationActionSafetyAuditItems(sources productionReceiptNotificationActionSafetyAuditSourceSet) []ProductionReceiptNotificationActionSafetyAuditItem {
	combined := sources.NotificationDeliveryGateAudit + sources.DesktopSideEffectReview + sources.DispatchSheet
	return []ProductionReceiptNotificationActionSafetyAuditItem{
		productionReceiptNotificationActionItem("review-receipt-action", "review", "future review action requires request-object safety", combined, []string{"review", "Runtime request object", "NotificationActionEnabled"}, "define request-object review flow before action enablement"),
		productionReceiptNotificationActionItem("renew-receipt-action", "renew", "future renew action requires explicit operator approval", combined, []string{"renew", "operator", "NotificationActionEnabled"}, "define renewal request safety before action enablement"),
		productionReceiptNotificationActionItem("open-compatibility-center-action", "open-compatibility-center", "future open Compatibility Center action requires navigation safety", combined, []string{"AI Compatibility Center", "navigation", "CompatibilityCenterPersisted"}, "define safe navigation target before action enablement"),
		productionReceiptNotificationActionItem("dismiss-receipt-action", "dismiss", "future dismiss action requires persistence safety", combined, []string{"notification action safety", "actions", "NotificationActionEnabled"}, "define dismiss persistence boundary before action enablement"),
		productionReceiptNotificationActionItem("support-info-action", "support-info", "future support info action requires support side-effect safety", combined, []string{"support requests", "support", "SupportCaseCreated"}, "define support-info read-only route before action enablement"),
	}
}

func productionReceiptNotificationActionItem(id string, kind string, evidence string, source string, tokens []string, nextRequirement string) ProductionReceiptNotificationActionSafetyAuditItem {
	ready := productionAuthorizationHasAll(source, tokens)
	status := "missing-notification-action-evidence"
	if ready {
		status = "notification-action-modeled-actions-disabled"
	}
	return ProductionReceiptNotificationActionSafetyAuditItem{
		ID:                            id,
		ActionKind:                    kind,
		RequiredEvidence:              evidence,
		EvidencePresent:               ready,
		ActionSafetyModeled:           ready,
		UserVisible:                   ready,
		ReviewOnly:                    true,
		RuntimeOwned:                  true,
		GoRuntimeBacked:               true,
		KDEPolicyOwner:                false,
		OperatorApprovalRequired:      true,
		OperatorApprovalPresent:       false,
		ActionEnabled:                 false,
		NavigationEnabled:             false,
		PortalRequestCreated:          false,
		RequestObjectsCreated:         false,
		NotificationSent:              false,
		NotificationDeliveryEnabled:   false,
		NotificationActionEnabled:     false,
		ReceiptAccepted:               false,
		AuthorizationAccepted:         false,
		ReceiptRevocationWriteEnabled: false,
		ReceiptExpiryWriteEnabled:     false,
		ReceiptPersistenceEnabled:     false,
		ReceiptLookupWritesEnabled:    false,
		CompatibilityCenterOpened:     false,
		CompatibilityCenterPersisted:  false,
		SupportBundleExported:         false,
		SupportCaseCreated:            false,
		ProductionReadiness:           false,
		ProductionOwnershipReady:      false,
		SideEffectsDisabled:           true,
		HostRootModified:              false,
		InternalDetailsExposed:        false,
		ActionStatus:                  status,
		NextRequirement:               nextRequirement,
	}
}

func productionReceiptNotificationActionSafetyAuditChecks(preview ProductionReceiptNotificationActionSafetyAuditPreview) []ProductionReceiptNotificationActionSafetyAuditCheck {
	return []ProductionReceiptNotificationActionSafetyAuditCheck{
		productionReceiptNotificationActionCheck("notification-delivery-gate-consumed", productionAuthorizationPassBlocked(preview.NotificationDeliveryGateConsumed), "The action safety audit consumes notification delivery gate evidence."),
		productionReceiptNotificationActionCheck("desktop-side-effect-review-consumed", productionAuthorizationPassBlocked(preview.DesktopSideEffectReviewConsumed), "The action safety audit consumes KDE desktop side-effect review evidence."),
		productionReceiptNotificationActionCheck("action-safety-modeled-only", productionAuthorizationPassBlocked(preview.ActionSafetyAuditRequired && preview.ActionSafetyAuditModeled && preview.NotificationActionSafetyReady && preview.OperatorActionApprovalRequired && !preview.OperatorActionApprovalPresent && !preview.ReceiptPresent && !preview.ReceiptAccepted && !preview.AuthorizationAccepted), "Notification action safety is modeled without approval, request creation, or navigation."),
		productionReceiptNotificationActionCheck("five-action-items-present", productionAuthorizationPassBlocked(preview.ActionItemCount == 5 && preview.RequiredActionItemCount == 5 && preview.MissingActionItemCount == 0), "Review, renew, open Compatibility Center, dismiss, and support-info actions are modeled."),
		productionReceiptNotificationActionCheck("action-items-ready-actions-disabled", productionAuthorizationPassBlocked(preview.ReadyActionItemCount == 5 && productionReceiptNotificationActionItemsReady(preview.ActionItems)), "Every action item is ready while actions, navigation, requests, and writes remain disabled."),
		productionReceiptNotificationActionCheck("notification-actions-disabled", productionAuthorizationPassBlocked(!preview.NotificationActionEnabled && !preview.ReviewActionEnabled && !preview.RenewActionEnabled && !preview.OpenCompatibilityCenterEnabled && !preview.DismissActionEnabled && !preview.SupportInfoActionEnabled && preview.EnabledActionItemCount == 0), "Notification actions and all per-action enablement gates remain disabled."),
		productionReceiptNotificationActionCheck("requests-and-navigation-disabled", productionAuthorizationPassBlocked(!preview.PortalRequestCreated && !preview.RequestObjectsCreated && !preview.CompatibilityCenterOpened && !preview.CompatibilityCenterPersisted && preview.RequestCreatedItemCount == 0 && preview.NavigationEnabledItemCount == 0), "Portal requests, Runtime request objects, Compatibility Center navigation, and persistence remain disabled."),
		productionReceiptNotificationActionCheck("receipt-and-notification-writes-disabled", productionAuthorizationPassBlocked(!preview.NotificationSent && !preview.NotificationDeliveryEnabled && !preview.ReceiptWriterEnabled && !preview.ReceiptPersistenceEnabled && !preview.ReceiptLookupWritesEnabled && !preview.ReceiptReplayEnabled && !preview.ReceiptExpiryWriteEnabled && !preview.ReceiptRevocationWriteEnabled), "Notification sending, delivery, receipt writes, lookup, replay, expiry, and revocation writes remain disabled."),
		productionReceiptNotificationActionCheck("production-and-host-boundary-closed", productionAuthorizationPassBlocked(!preview.SystemServiceStarted && !preview.SessionBusClaimed && !preview.ProductionBusClaimed && !preview.ProductionOwnerEnabled && !preview.ProductionActivationReady && !preview.ProductionReadiness && !preview.ProductionOwnershipReady && !preview.FileContentRead && !preview.FilePathsExposed && !preview.NetworkRequired && !preview.HostRootModified && !preview.PrivilegedContainerRequired && !preview.StateRootPathExposed && !preview.RawCommandExposed && !preview.RawExecutableExposed && !preview.BackendDetailsExposed && productionReceiptNotificationActionItemsKeepHostClosed(preview.ActionItems)), "Production ownership, unsafe data exposure, network, privilege, path exposure, internal details, and host mutation remain disabled."),
	}
}

func productionReceiptNotificationActionCheck(id string, status string, summary string) ProductionReceiptNotificationActionSafetyAuditCheck {
	return ProductionReceiptNotificationActionSafetyAuditCheck{ID: id, Status: status, Summary: summary}
}

func productionReceiptNotificationActionDeliveryGateReady(source string) bool {
	return productionAuthorizationHasAll(source, []string{
		"production-receipt-notification-delivery-gate-audit-preview",
		"receipt-expiry-revocation-notification-delivery-gate-audit",
		"production-receipt-notification-delivery-gate-audit-ready-delivery-disabled",
		"notification-delivery-disabled",
		"requests-and-persistence-disabled",
		"NotificationActionEnabled",
	})
}

func productionReceiptNotificationActionDesktopReady(source string) bool {
	return productionAuthorizationHasAll(source, []string{
		"production-desktop-side-effect-review-preview",
		"kde-production-desktop-side-effect-review",
		"production-desktop-side-effect-review-ready-side-effects-disabled",
		"notifications-and-requests-disabled",
		"CompatibilityCenterPersisted",
	})
}

func productionReceiptNotificationActionReadyItemCount(items []ProductionReceiptNotificationActionSafetyAuditItem) int {
	count := 0
	for _, item := range items {
		if item.EvidencePresent {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionMissingItemCount(items []ProductionReceiptNotificationActionSafetyAuditItem) int {
	count := 0
	for _, item := range items {
		if !item.EvidencePresent {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionEnabledCount(items []ProductionReceiptNotificationActionSafetyAuditItem) int {
	count := 0
	for _, item := range items {
		if item.ActionEnabled || item.NotificationActionEnabled {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionRequestCount(items []ProductionReceiptNotificationActionSafetyAuditItem) int {
	count := 0
	for _, item := range items {
		if item.PortalRequestCreated || item.RequestObjectsCreated {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionNavigationCount(items []ProductionReceiptNotificationActionSafetyAuditItem) int {
	count := 0
	for _, item := range items {
		if item.NavigationEnabled || item.CompatibilityCenterOpened {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionSideEffectCount(items []ProductionReceiptNotificationActionSafetyAuditItem) int {
	count := 0
	for _, item := range items {
		if item.ActionEnabled || item.NavigationEnabled || item.PortalRequestCreated || item.RequestObjectsCreated || item.NotificationSent || item.NotificationDeliveryEnabled || item.NotificationActionEnabled || item.CompatibilityCenterOpened || item.CompatibilityCenterPersisted || item.SupportBundleExported || item.SupportCaseCreated || item.HostRootModified || item.InternalDetailsExposed {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionItemIDs(items []ProductionReceiptNotificationActionSafetyAuditItem) []string {
	ids := make([]string, 0, len(items))
	for _, item := range items {
		ids = append(ids, item.ID)
	}
	return ids
}

func productionReceiptNotificationActionCheckIDs(checks []ProductionReceiptNotificationActionSafetyAuditCheck) []string {
	ids := make([]string, 0, len(checks))
	for _, check := range checks {
		ids = append(ids, check.ID)
	}
	return ids
}

func countProductionReceiptNotificationActionChecks(checks []ProductionReceiptNotificationActionSafetyAuditCheck) ProductionReceiptNotificationActionSafetyAuditCounts {
	counts := ProductionReceiptNotificationActionSafetyAuditCounts{Total: len(checks)}
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

func productionReceiptNotificationActionItemsReady(items []ProductionReceiptNotificationActionSafetyAuditItem) bool {
	if len(items) != 5 {
		return false
	}
	for _, item := range items {
		if !item.EvidencePresent || !item.ActionSafetyModeled || !item.UserVisible || item.ActionStatus != "notification-action-modeled-actions-disabled" {
			return false
		}
		if item.OperatorApprovalPresent || item.ActionEnabled || item.NavigationEnabled || item.PortalRequestCreated || item.RequestObjectsCreated || item.NotificationSent || item.NotificationDeliveryEnabled || item.NotificationActionEnabled || item.ReceiptAccepted || item.AuthorizationAccepted || item.ReceiptRevocationWriteEnabled || item.ReceiptExpiryWriteEnabled || item.ReceiptPersistenceEnabled || item.ReceiptLookupWritesEnabled || item.CompatibilityCenterOpened || item.CompatibilityCenterPersisted || item.SupportBundleExported || item.SupportCaseCreated || item.ProductionReadiness || item.ProductionOwnershipReady || item.KDEPolicyOwner || !item.ReviewOnly || !item.SideEffectsDisabled {
			return false
		}
	}
	return true
}

func productionReceiptNotificationActionItemsKeepHostClosed(items []ProductionReceiptNotificationActionSafetyAuditItem) bool {
	for _, item := range items {
		if item.HostRootModified || item.InternalDetailsExposed {
			return false
		}
	}
	return true
}
