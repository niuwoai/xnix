package owner

type ProductionReceiptNotificationDeliveryGateAuditPreview struct {
	Version                              string                                                `json:"version"`
	SchemaVersion                        string                                                `json:"schema_version"`
	RequestType                          string                                                `json:"request_type"`
	AuditType                            string                                                `json:"audit_type"`
	Source                               string                                                `json:"source"`
	AuditDecision                        string                                                `json:"audit_decision"`
	ReceiptSchema                        string                                                `json:"receipt_schema"`
	OpaqueReceiptID                      string                                                `json:"opaque_receipt_id"`
	NotificationGateRequired             bool                                                  `json:"notification_gate_required"`
	NotificationGateModeled              bool                                                  `json:"notification_gate_modeled"`
	RevocationVisibilityAuditConsumed    bool                                                  `json:"revocation_visibility_audit_consumed"`
	DesktopSideEffectReviewConsumed      bool                                                  `json:"desktop_side_effect_review_consumed"`
	OperatorNotificationApprovalRequired bool                                                  `json:"operator_notification_approval_required"`
	OperatorNotificationApprovalPresent  bool                                                  `json:"operator_notification_approval_present"`
	NotificationDeliveryGateReady        bool                                                  `json:"notification_delivery_gate_ready"`
	CallerStateRootRequired              bool                                                  `json:"caller_state_root_required"`
	ReceiptPresent                       bool                                                  `json:"receipt_present"`
	ReceiptAccepted                      bool                                                  `json:"receipt_accepted"`
	AuthorizationAccepted                bool                                                  `json:"authorization_accepted"`
	ProductionReadiness                  bool                                                  `json:"production_readiness"`
	ProductionOwnershipReady             bool                                                  `json:"production_ownership_ready"`
	GateItemCount                        int                                                   `json:"gate_item_count"`
	RequiredGateItemCount                int                                                   `json:"required_gate_item_count"`
	ReadyGateItemCount                   int                                                   `json:"ready_gate_item_count"`
	MissingGateItemCount                 int                                                   `json:"missing_gate_item_count"`
	DeliveryEnabledItemCount             int                                                   `json:"delivery_enabled_item_count"`
	NotificationSentItemCount            int                                                   `json:"notification_sent_item_count"`
	RequestCreatedItemCount              int                                                   `json:"request_created_item_count"`
	SideEffectItemCount                  int                                                   `json:"side_effect_item_count"`
	GateItems                            []ProductionReceiptNotificationDeliveryGateAuditItem  `json:"gate_items"`
	GateItemIDs                          []string                                              `json:"gate_item_ids"`
	RequiredBeforeNotificationDelivery   []string                                              `json:"required_before_notification_delivery"`
	Checks                               []ProductionReceiptNotificationDeliveryGateAuditCheck `json:"checks"`
	CheckIDs                             []string                                              `json:"check_ids"`
	Counts                               ProductionReceiptNotificationDeliveryGateAuditCounts  `json:"counts"`
	RuntimeOwned                         bool                                                  `json:"runtime_owned"`
	GoRuntimeBacked                      bool                                                  `json:"go_runtime_backed"`
	KDEPolicyOwner                       bool                                                  `json:"kde_policy_owner"`
	OfficialDesktopOnly                  bool                                                  `json:"official_desktop_only"`
	PlasmaForkRequired                   bool                                                  `json:"plasma_fork_required"`
	PlasmaSourceModified                 bool                                                  `json:"plasma_source_modified"`
	SystemServiceStarted                 bool                                                  `json:"system_service_started"`
	SessionBusClaimed                    bool                                                  `json:"session_bus_claimed"`
	ProductionBusClaimed                 bool                                                  `json:"production_bus_claimed"`
	ProductionOwnerEnabled               bool                                                  `json:"production_owner_enabled"`
	ProductionActivationReady            bool                                                  `json:"production_activation_ready"`
	WriteMethodsEnabled                  bool                                                  `json:"write_methods_enabled"`
	RuntimeWritesEnabled                 bool                                                  `json:"runtime_writes_enabled"`
	ReceiptWriterEnabled                 bool                                                  `json:"receipt_writer_enabled"`
	ReceiptPersistenceEnabled            bool                                                  `json:"receipt_persistence_enabled"`
	ReceiptLookupWritesEnabled           bool                                                  `json:"receipt_lookup_writes_enabled"`
	ReceiptReplayEnabled                 bool                                                  `json:"receipt_replay_enabled"`
	ReceiptExpiryWriteEnabled            bool                                                  `json:"receipt_expiry_write_enabled"`
	ReceiptRevocationWriteEnabled        bool                                                  `json:"receipt_revocation_write_enabled"`
	ReceiptRevocationVisibilityPersisted bool                                                  `json:"receipt_revocation_visibility_persisted"`
	NotificationDeliveryGatePersisted    bool                                                  `json:"notification_delivery_gate_persisted"`
	DesktopFilesWritten                  bool                                                  `json:"desktop_files_written"`
	MIMEAppsWritten                      bool                                                  `json:"mimeapps_written"`
	ShellConfigurationWritten            bool                                                  `json:"shell_configuration_written"`
	SettingsPersisted                    bool                                                  `json:"settings_persisted"`
	KRunnerIndexPersisted                bool                                                  `json:"krunner_index_persisted"`
	TaskManagerEntryActive               bool                                                  `json:"task_manager_entry_active"`
	KWinRuleApplied                      bool                                                  `json:"kwin_rule_applied"`
	LiveTrayBridgeEnabled                bool                                                  `json:"live_tray_bridge_enabled"`
	TrayBridgePersisted                  bool                                                  `json:"tray_bridge_persisted"`
	NotificationSent                     bool                                                  `json:"notification_sent"`
	NotificationDeliveryEnabled          bool                                                  `json:"notification_delivery_enabled"`
	NotificationActionEnabled            bool                                                  `json:"notification_action_enabled"`
	CompatibilityCenterPersisted         bool                                                  `json:"compatibility_center_persisted"`
	PortalRequestCreated                 bool                                                  `json:"portal_request_created"`
	RequestObjectsCreated                bool                                                  `json:"request_objects_created"`
	AdapterInvocationEnabled             bool                                                  `json:"adapter_invocation_enabled"`
	BackendLaunchEnabled                 bool                                                  `json:"backend_launch_enabled"`
	BackendProcessStarted                bool                                                  `json:"backend_process_started"`
	SupportBundleExported                bool                                                  `json:"support_bundle_exported"`
	SupportCaseCreated                   bool                                                  `json:"support_case_created"`
	SnapshotRestoreExecuted              bool                                                  `json:"snapshot_restore_executed"`
	StateCleanupExecuted                 bool                                                  `json:"state_cleanup_executed"`
	FileContentRead                      bool                                                  `json:"file_content_read"`
	FilePathsExposed                     bool                                                  `json:"file_paths_exposed"`
	NetworkRequired                      bool                                                  `json:"network_required"`
	HostRootModified                     bool                                                  `json:"host_root_modified"`
	PrivilegedContainerRequired          bool                                                  `json:"privileged_container_required"`
	StateRootPathExposed                 bool                                                  `json:"state_root_path_exposed"`
	RawCommandExposed                    bool                                                  `json:"raw_command_exposed"`
	RawExecutableExposed                 bool                                                  `json:"raw_executable_exposed"`
	BackendDetailsExposed                bool                                                  `json:"backend_details_exposed"`
	BlockedActions                       []string                                              `json:"blocked_actions"`
	NextRequirements                     []string                                              `json:"next_requirements"`
	DesktopSafeSummary                   string                                                `json:"desktop_safe_summary"`
}

type ProductionReceiptNotificationDeliveryGateAuditItem struct {
	ID                            string `json:"id"`
	NotificationCase              string `json:"notification_case"`
	RequiredEvidence              string `json:"required_evidence"`
	EvidencePresent               bool   `json:"evidence_present"`
	GateModeled                   bool   `json:"gate_modeled"`
	UserVisible                   bool   `json:"user_visible"`
	ReviewOnly                    bool   `json:"review_only"`
	RuntimeOwned                  bool   `json:"runtime_owned"`
	GoRuntimeBacked               bool   `json:"go_runtime_backed"`
	KDEPolicyOwner                bool   `json:"kde_policy_owner"`
	OperatorApprovalRequired      bool   `json:"operator_approval_required"`
	OperatorApprovalPresent       bool   `json:"operator_approval_present"`
	NotificationSent              bool   `json:"notification_sent"`
	NotificationDeliveryEnabled   bool   `json:"notification_delivery_enabled"`
	NotificationActionEnabled     bool   `json:"notification_action_enabled"`
	PortalRequestCreated          bool   `json:"portal_request_created"`
	RequestObjectsCreated         bool   `json:"request_objects_created"`
	ReceiptRevocationWriteEnabled bool   `json:"receipt_revocation_write_enabled"`
	ReceiptExpiryWriteEnabled     bool   `json:"receipt_expiry_write_enabled"`
	ReceiptPersistenceEnabled     bool   `json:"receipt_persistence_enabled"`
	ReceiptLookupWritesEnabled    bool   `json:"receipt_lookup_writes_enabled"`
	ReceiptAccepted               bool   `json:"receipt_accepted"`
	AuthorizationAccepted         bool   `json:"authorization_accepted"`
	ProductionReadiness           bool   `json:"production_readiness"`
	ProductionOwnershipReady      bool   `json:"production_ownership_ready"`
	SideEffectsDisabled           bool   `json:"side_effects_disabled"`
	HostRootModified              bool   `json:"host_root_modified"`
	InternalDetailsExposed        bool   `json:"internal_details_exposed"`
	GateStatus                    string `json:"gate_status"`
	NextRequirement               string `json:"next_requirement"`
}

type ProductionReceiptNotificationDeliveryGateAuditCheck struct {
	ID      string `json:"id"`
	Status  string `json:"status"`
	Summary string `json:"summary"`
}

type ProductionReceiptNotificationDeliveryGateAuditCounts struct {
	Total   int `json:"total"`
	Passed  int `json:"passed"`
	Pending int `json:"pending"`
	Blocked int `json:"blocked"`
}

func NewProductionReceiptNotificationDeliveryGateAuditPreview(root string) (ProductionReceiptNotificationDeliveryGateAuditPreview, error) {
	version, err := readProductionDBusGateReviewVersion(root)
	if err != nil {
		return ProductionReceiptNotificationDeliveryGateAuditPreview{}, err
	}
	sources := productionReceiptNotificationDeliveryGateAuditSources(root)
	items := productionReceiptNotificationDeliveryGateAuditItems(sources)
	visibilityReady := productionReceiptNotificationDeliveryVisibilityReady(sources.RevocationVisibilityAudit)
	desktopReady := productionReceiptNotificationDeliveryDesktopReady(sources.DesktopSideEffectReview)
	gateReady := visibilityReady && desktopReady && productionReceiptNotificationDeliveryItemsReady(items)
	preview := ProductionReceiptNotificationDeliveryGateAuditPreview{
		Version:                              version,
		SchemaVersion:                        "xnix.runtime.production_receipt_notification_delivery_gate_audit.v1",
		RequestType:                          "production-receipt-notification-delivery-gate-audit-preview",
		AuditType:                            "receipt-expiry-revocation-notification-delivery-gate-audit",
		Source:                               "production-receipt-revocation-visibility-audit-preview+production-desktop-side-effect-review-preview+claude-code-current-dispatch-picks",
		AuditDecision:                        "production-receipt-notification-delivery-gate-audit-blocked",
		ReceiptSchema:                        "xnix.runtime.production_dbus_human_authorization_receipt.v1",
		OpaqueReceiptID:                      ProductionDBusHumanAuthorizationReceiptID,
		NotificationGateRequired:             true,
		NotificationGateModeled:              true,
		RevocationVisibilityAuditConsumed:    visibilityReady,
		DesktopSideEffectReviewConsumed:      desktopReady,
		OperatorNotificationApprovalRequired: true,
		OperatorNotificationApprovalPresent:  false,
		NotificationDeliveryGateReady:        gateReady,
		CallerStateRootRequired:              false,
		ReceiptPresent:                       false,
		ReceiptAccepted:                      false,
		AuthorizationAccepted:                false,
		ProductionReadiness:                  false,
		ProductionOwnershipReady:             false,
		GateItemCount:                        len(items),
		RequiredGateItemCount:                5,
		ReadyGateItemCount:                   productionReceiptNotificationDeliveryReadyItemCount(items),
		MissingGateItemCount:                 productionReceiptNotificationDeliveryMissingItemCount(items),
		GateItems:                            items,
		GateItemIDs:                          productionReceiptNotificationDeliveryItemIDs(items),
		RequiredBeforeNotificationDelivery: []string{
			"production-receipt-revocation-visibility-audit-preview",
			"production-desktop-side-effect-review-preview",
			"explicit operator notification delivery approval",
			"separate desktop notification side-effect implementation",
			"separate notification action and Portal request review",
		},
		RuntimeOwned:                  true,
		GoRuntimeBacked:               true,
		KDEPolicyOwner:                false,
		OfficialDesktopOnly:           true,
		PlasmaForkRequired:            false,
		PlasmaSourceModified:          false,
		SystemServiceStarted:          false,
		SessionBusClaimed:             false,
		ProductionBusClaimed:          false,
		ProductionOwnerEnabled:        false,
		ProductionActivationReady:     false,
		WriteMethodsEnabled:           false,
		RuntimeWritesEnabled:          false,
		ReceiptWriterEnabled:          false,
		ReceiptPersistenceEnabled:     false,
		ReceiptLookupWritesEnabled:    false,
		ReceiptReplayEnabled:          false,
		ReceiptExpiryWriteEnabled:     false,
		ReceiptRevocationWriteEnabled: false,
		DesktopFilesWritten:           false,
		MIMEAppsWritten:               false,
		ShellConfigurationWritten:     false,
		SettingsPersisted:             false,
		KRunnerIndexPersisted:         false,
		TaskManagerEntryActive:        false,
		KWinRuleApplied:               false,
		LiveTrayBridgeEnabled:         false,
		TrayBridgePersisted:           false,
		NotificationSent:              false,
		NotificationDeliveryEnabled:   false,
		NotificationActionEnabled:     false,
		CompatibilityCenterPersisted:  false,
		PortalRequestCreated:          false,
		RequestObjectsCreated:         false,
		AdapterInvocationEnabled:      false,
		BackendLaunchEnabled:          false,
		BackendProcessStarted:         false,
		SupportBundleExported:         false,
		SupportCaseCreated:            false,
		SnapshotRestoreExecuted:       false,
		StateCleanupExecuted:          false,
		FileContentRead:               false,
		FilePathsExposed:              false,
		NetworkRequired:               false,
		HostRootModified:              false,
		PrivilegedContainerRequired:   false,
		StateRootPathExposed:          false,
		RawCommandExposed:             false,
		RawExecutableExposed:          false,
		BackendDetailsExposed:         false,
		BlockedActions: []string{
			"treat notification delivery gate audit as permission to send notifications",
			"persist notification delivery gates, Compatibility Center state, receipt state, settings, desktop files, or MIME defaults from this audit",
			"create notification actions, Portal requests, Runtime request objects, support cases, support bundles, restore snapshots, or clean state",
			"write, accept, persist, replay, expire, revoke, or look up authorization receipts from this audit",
			"claim production D-Bus ownership, enable Runtime writes, start services, launch compatibility engines, expose paths, or mutate host root",
		},
		NextRequirements: []string{
			"Require explicit operator notification delivery approval before expired or revoked receipt states can notify users.",
			"Keep notification delivery separate from revocation visibility; visibility remains read-only and delivery remains disabled.",
			"Require a separate desktop notification side-effect implementation before any KDE notification transport is used.",
			"Keep KDE as shell and presentation only; Runtime remains the owner of compatibility policy and notification gate decisions.",
		},
		DesktopSafeSummary: "The production receipt notification delivery gate audit models the separate gate required before expiring, expired, revoked, and missing-review receipt states can notify users, but it sends no notifications, creates no actions or requests, writes no receipts, starts no services, launches no engines, and mutates no host state.",
	}
	preview.DeliveryEnabledItemCount = productionReceiptNotificationDeliveryEnabledCount(items)
	preview.NotificationSentItemCount = productionReceiptNotificationSentCount(items)
	preview.RequestCreatedItemCount = productionReceiptNotificationRequestCreatedCount(items)
	preview.SideEffectItemCount = productionReceiptNotificationSideEffectCount(items)
	checks := productionReceiptNotificationDeliveryGateAuditChecks(preview)
	preview.Checks = checks
	preview.CheckIDs = productionReceiptNotificationDeliveryCheckIDs(checks)
	preview.Counts = countProductionReceiptNotificationDeliveryChecks(checks)
	if preview.Counts.Blocked == 0 {
		preview.AuditDecision = "production-receipt-notification-delivery-gate-audit-ready-delivery-disabled"
	}
	if err := validateNoBackendTerms(preview, "production receipt notification delivery gate audit preview"); err != nil {
		return ProductionReceiptNotificationDeliveryGateAuditPreview{}, err
	}
	return preview, nil
}

type productionReceiptNotificationDeliveryGateAuditSourceSet struct {
	RevocationVisibilityAudit string
	DesktopSideEffectReview   string
	DispatchSheet             string
}

func productionReceiptNotificationDeliveryGateAuditSources(root string) productionReceiptNotificationDeliveryGateAuditSourceSet {
	return productionReceiptNotificationDeliveryGateAuditSourceSet{
		RevocationVisibilityAudit: productionAuthorizationReadSources(root, []string{"internal/runtime/owner/production_receipt_revocation_visibility_audit.go"}),
		DesktopSideEffectReview:   productionAuthorizationReadSources(root, []string{"internal/runtime/owner/production_desktop_side_effect_review.go"}),
		DispatchSheet:             productionAuthorizationReadSources(root, []string{"docs/claude-code-current-dispatch-picks.md"}),
	}
}

func productionReceiptNotificationDeliveryGateAuditItems(sources productionReceiptNotificationDeliveryGateAuditSourceSet) []ProductionReceiptNotificationDeliveryGateAuditItem {
	combined := sources.RevocationVisibilityAudit + sources.DesktopSideEffectReview + sources.DispatchSheet
	return []ProductionReceiptNotificationDeliveryGateAuditItem{
		productionReceiptNotificationDeliveryItem("expiring-receipt-warning-gate", "expiring-receipt-warning", "future expiring receipt warning requires explicit notification delivery approval", combined, []string{"expiring", "notification delivery", "NotificationDeliveryEnabled"}, "define review-only warning copy before delivery can be approved"),
		productionReceiptNotificationDeliveryItem("expired-receipt-blocker-gate", "expired-receipt-blocker", "future expired receipt blocker requires desktop side-effect approval", combined, []string{"expired", "production gates", "notifications-and-desktop-side-effects-disabled"}, "define expired-state notification policy before delivery can be approved"),
		productionReceiptNotificationDeliveryItem("revoked-receipt-blocker-gate", "revoked-receipt-blocker", "future revoked receipt blocker requires revocation visibility without revocation writes", combined, []string{"revoked", "revocation writes", "NotificationDeliveryEnabled"}, "define revoked-state notification policy before delivery can be approved"),
		productionReceiptNotificationDeliveryItem("missing-review-reminder-gate", "missing-review-reminder", "future missing-review reminder requires operator approval", combined, []string{"missing-review", "operator", "notification delivery gate"}, "define reminder cadence and redaction before delivery can be approved"),
		productionReceiptNotificationDeliveryItem("operator-approval-gate", "operator-notification-approval", "future notification delivery requires separate operator approval", combined, []string{"operator", "notification action safety", "notification delivery"}, "record explicit operator approval before any delivery transport is enabled"),
	}
}

func productionReceiptNotificationDeliveryItem(id string, notificationCase string, evidence string, source string, tokens []string, nextRequirement string) ProductionReceiptNotificationDeliveryGateAuditItem {
	ready := productionAuthorizationHasAll(source, tokens)
	status := "missing-notification-gate-evidence"
	if ready {
		status = "notification-gate-modeled-delivery-disabled"
	}
	return ProductionReceiptNotificationDeliveryGateAuditItem{
		ID:                            id,
		NotificationCase:              notificationCase,
		RequiredEvidence:              evidence,
		EvidencePresent:               ready,
		GateModeled:                   ready,
		UserVisible:                   ready,
		ReviewOnly:                    true,
		RuntimeOwned:                  true,
		GoRuntimeBacked:               true,
		KDEPolicyOwner:                false,
		OperatorApprovalRequired:      true,
		OperatorApprovalPresent:       false,
		NotificationSent:              false,
		NotificationDeliveryEnabled:   false,
		NotificationActionEnabled:     false,
		PortalRequestCreated:          false,
		RequestObjectsCreated:         false,
		ReceiptRevocationWriteEnabled: false,
		ReceiptExpiryWriteEnabled:     false,
		ReceiptPersistenceEnabled:     false,
		ReceiptLookupWritesEnabled:    false,
		ReceiptAccepted:               false,
		AuthorizationAccepted:         false,
		ProductionReadiness:           false,
		ProductionOwnershipReady:      false,
		SideEffectsDisabled:           true,
		HostRootModified:              false,
		InternalDetailsExposed:        false,
		GateStatus:                    status,
		NextRequirement:               nextRequirement,
	}
}

func productionReceiptNotificationDeliveryGateAuditChecks(preview ProductionReceiptNotificationDeliveryGateAuditPreview) []ProductionReceiptNotificationDeliveryGateAuditCheck {
	return []ProductionReceiptNotificationDeliveryGateAuditCheck{
		productionReceiptNotificationDeliveryCheck("revocation-visibility-audit-consumed", productionAuthorizationPassBlocked(preview.RevocationVisibilityAuditConsumed), "The notification gate consumes the revocation visibility audit."),
		productionReceiptNotificationDeliveryCheck("desktop-side-effect-review-consumed", productionAuthorizationPassBlocked(preview.DesktopSideEffectReviewConsumed), "The notification gate consumes KDE desktop side-effect review evidence."),
		productionReceiptNotificationDeliveryCheck("notification-gate-modeled-only", productionAuthorizationPassBlocked(preview.NotificationGateRequired && preview.NotificationGateModeled && preview.NotificationDeliveryGateReady && preview.OperatorNotificationApprovalRequired && !preview.OperatorNotificationApprovalPresent && !preview.ReceiptPresent && !preview.ReceiptAccepted && !preview.AuthorizationAccepted), "Notification delivery is modeled without operator approval, receipt acceptance, or delivery."),
		productionReceiptNotificationDeliveryCheck("five-notification-gate-items-present", productionAuthorizationPassBlocked(preview.GateItemCount == 5 && preview.RequiredGateItemCount == 5 && preview.MissingGateItemCount == 0), "Expiring, expired, revoked, missing-review, and operator approval notification gates are modeled."),
		productionReceiptNotificationDeliveryCheck("gate-items-ready-delivery-disabled", productionAuthorizationPassBlocked(preview.ReadyGateItemCount == 5 && productionReceiptNotificationDeliveryItemsReady(preview.GateItems)), "Every notification gate item is ready while delivery, actions, requests, and writes remain disabled."),
		productionReceiptNotificationDeliveryCheck("notification-delivery-disabled", productionAuthorizationPassBlocked(!preview.NotificationSent && !preview.NotificationDeliveryEnabled && !preview.NotificationActionEnabled && preview.DeliveryEnabledItemCount == 0 && preview.NotificationSentItemCount == 0), "Notification sending, delivery, and actions remain disabled."),
		productionReceiptNotificationDeliveryCheck("requests-and-persistence-disabled", productionAuthorizationPassBlocked(!preview.NotificationDeliveryGatePersisted && !preview.ReceiptRevocationVisibilityPersisted && !preview.CompatibilityCenterPersisted && !preview.PortalRequestCreated && !preview.RequestObjectsCreated && preview.RequestCreatedItemCount == 0), "Gate persistence, Compatibility Center persistence, Portal requests, and Runtime request objects remain disabled."),
		productionReceiptNotificationDeliveryCheck("receipt-writes-disabled", productionAuthorizationPassBlocked(!preview.ReceiptWriterEnabled && !preview.ReceiptPersistenceEnabled && !preview.ReceiptLookupWritesEnabled && !preview.ReceiptReplayEnabled && !preview.ReceiptExpiryWriteEnabled && !preview.ReceiptRevocationWriteEnabled), "Receipt writes, persistence, lookup, replay, expiry writes, and revocation writes remain disabled."),
		productionReceiptNotificationDeliveryCheck("production-and-host-boundary-closed", productionAuthorizationPassBlocked(!preview.SystemServiceStarted && !preview.SessionBusClaimed && !preview.ProductionBusClaimed && !preview.ProductionOwnerEnabled && !preview.ProductionActivationReady && !preview.ProductionReadiness && !preview.ProductionOwnershipReady && !preview.FileContentRead && !preview.FilePathsExposed && !preview.NetworkRequired && !preview.HostRootModified && !preview.PrivilegedContainerRequired && !preview.StateRootPathExposed && !preview.RawCommandExposed && !preview.RawExecutableExposed && !preview.BackendDetailsExposed && productionReceiptNotificationDeliveryItemsKeepHostClosed(preview.GateItems)), "Production ownership, unsafe data exposure, network, privilege, path exposure, internal details, and host mutation remain disabled."),
	}
}

func productionReceiptNotificationDeliveryCheck(id string, status string, summary string) ProductionReceiptNotificationDeliveryGateAuditCheck {
	return ProductionReceiptNotificationDeliveryGateAuditCheck{ID: id, Status: status, Summary: summary}
}

func productionReceiptNotificationDeliveryVisibilityReady(source string) bool {
	return productionAuthorizationHasAll(source, []string{
		"production-receipt-revocation-visibility-audit-preview",
		"receipt-expiry-revocation-production-kde-visibility-audit",
		"production-receipt-revocation-visibility-audit-ready-visibility-only",
		"receipt-expiring-visible",
		"receipt-expired-visible",
		"receipt-revoked-visible",
		"receipt-missing-review-visible",
		"notifications-and-desktop-side-effects-disabled",
	})
}

func productionReceiptNotificationDeliveryDesktopReady(source string) bool {
	return productionAuthorizationHasAll(source, []string{
		"production-desktop-side-effect-review-preview",
		"kde-production-desktop-side-effect-review",
		"production-desktop-side-effect-review-ready-side-effects-disabled",
		"notifications",
		"NotificationSent",
		"NotificationDeliveryEnabled",
	})
}

func productionReceiptNotificationDeliveryReadyItemCount(items []ProductionReceiptNotificationDeliveryGateAuditItem) int {
	count := 0
	for _, item := range items {
		if item.EvidencePresent {
			count++
		}
	}
	return count
}

func productionReceiptNotificationDeliveryMissingItemCount(items []ProductionReceiptNotificationDeliveryGateAuditItem) int {
	count := 0
	for _, item := range items {
		if !item.EvidencePresent {
			count++
		}
	}
	return count
}

func productionReceiptNotificationDeliveryEnabledCount(items []ProductionReceiptNotificationDeliveryGateAuditItem) int {
	count := 0
	for _, item := range items {
		if item.NotificationDeliveryEnabled || item.NotificationActionEnabled {
			count++
		}
	}
	return count
}

func productionReceiptNotificationSentCount(items []ProductionReceiptNotificationDeliveryGateAuditItem) int {
	count := 0
	for _, item := range items {
		if item.NotificationSent {
			count++
		}
	}
	return count
}

func productionReceiptNotificationRequestCreatedCount(items []ProductionReceiptNotificationDeliveryGateAuditItem) int {
	count := 0
	for _, item := range items {
		if item.PortalRequestCreated || item.RequestObjectsCreated {
			count++
		}
	}
	return count
}

func productionReceiptNotificationSideEffectCount(items []ProductionReceiptNotificationDeliveryGateAuditItem) int {
	count := 0
	for _, item := range items {
		if item.NotificationSent || item.NotificationDeliveryEnabled || item.NotificationActionEnabled || item.PortalRequestCreated || item.RequestObjectsCreated || item.HostRootModified || item.InternalDetailsExposed {
			count++
		}
	}
	return count
}

func productionReceiptNotificationDeliveryItemIDs(items []ProductionReceiptNotificationDeliveryGateAuditItem) []string {
	ids := make([]string, 0, len(items))
	for _, item := range items {
		ids = append(ids, item.ID)
	}
	return ids
}

func productionReceiptNotificationDeliveryCheckIDs(checks []ProductionReceiptNotificationDeliveryGateAuditCheck) []string {
	ids := make([]string, 0, len(checks))
	for _, check := range checks {
		ids = append(ids, check.ID)
	}
	return ids
}

func countProductionReceiptNotificationDeliveryChecks(checks []ProductionReceiptNotificationDeliveryGateAuditCheck) ProductionReceiptNotificationDeliveryGateAuditCounts {
	counts := ProductionReceiptNotificationDeliveryGateAuditCounts{Total: len(checks)}
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

func productionReceiptNotificationDeliveryItemsReady(items []ProductionReceiptNotificationDeliveryGateAuditItem) bool {
	if len(items) != 5 {
		return false
	}
	for _, item := range items {
		if !item.EvidencePresent || !item.GateModeled || !item.UserVisible || item.GateStatus != "notification-gate-modeled-delivery-disabled" {
			return false
		}
		if item.OperatorApprovalPresent || item.NotificationSent || item.NotificationDeliveryEnabled || item.NotificationActionEnabled || item.PortalRequestCreated || item.RequestObjectsCreated || item.ReceiptRevocationWriteEnabled || item.ReceiptExpiryWriteEnabled || item.ReceiptPersistenceEnabled || item.ReceiptLookupWritesEnabled || item.ReceiptAccepted || item.AuthorizationAccepted || item.ProductionReadiness || item.ProductionOwnershipReady || item.KDEPolicyOwner || !item.ReviewOnly || !item.SideEffectsDisabled {
			return false
		}
	}
	return true
}

func productionReceiptNotificationDeliveryItemsKeepHostClosed(items []ProductionReceiptNotificationDeliveryGateAuditItem) bool {
	for _, item := range items {
		if item.HostRootModified || item.InternalDetailsExposed {
			return false
		}
	}
	return true
}
