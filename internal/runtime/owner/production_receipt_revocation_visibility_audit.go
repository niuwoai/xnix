package owner

type ProductionReceiptRevocationVisibilityAuditPreview struct {
	Version                              string                                            `json:"version"`
	SchemaVersion                        string                                            `json:"schema_version"`
	RequestType                          string                                            `json:"request_type"`
	AuditType                            string                                            `json:"audit_type"`
	Source                               string                                            `json:"source"`
	AuditDecision                        string                                            `json:"audit_decision"`
	ReceiptSchema                        string                                            `json:"receipt_schema"`
	OpaqueReceiptID                      string                                            `json:"opaque_receipt_id"`
	VisibilityAuditRequired              bool                                              `json:"visibility_audit_required"`
	VisibilityAuditModeled               bool                                              `json:"visibility_audit_modeled"`
	PersistenceThreatReviewConsumed      bool                                              `json:"persistence_threat_review_consumed"`
	DesktopSideEffectReviewConsumed      bool                                              `json:"desktop_side_effect_review_consumed"`
	ProductionGateVisibilityModeled      bool                                              `json:"production_gate_visibility_modeled"`
	KDEStatusVisibilityModeled           bool                                              `json:"kde_status_visibility_modeled"`
	CallerStateRootRequired              bool                                              `json:"caller_state_root_required"`
	ReceiptPresent                       bool                                              `json:"receipt_present"`
	ReceiptAccepted                      bool                                              `json:"receipt_accepted"`
	AuthorizationAccepted                bool                                              `json:"authorization_accepted"`
	RevocationVisibilityReady            bool                                              `json:"revocation_visibility_ready"`
	ProductionReadiness                  bool                                              `json:"production_readiness"`
	ProductionOwnershipReady             bool                                              `json:"production_ownership_ready"`
	VisibilityItemCount                  int                                               `json:"visibility_item_count"`
	RequiredVisibilityItemCount          int                                               `json:"required_visibility_item_count"`
	ReadyVisibilityItemCount             int                                               `json:"ready_visibility_item_count"`
	MissingVisibilityItemCount           int                                               `json:"missing_visibility_item_count"`
	RevocationWriteEnabledItemCount      int                                               `json:"revocation_write_enabled_item_count"`
	ExpiryWriteEnabledItemCount          int                                               `json:"expiry_write_enabled_item_count"`
	NotificationEnabledItemCount         int                                               `json:"notification_enabled_item_count"`
	SideEffectItemCount                  int                                               `json:"side_effect_item_count"`
	VisibilityItems                      []ProductionReceiptRevocationVisibilityAuditItem  `json:"visibility_items"`
	VisibilityItemIDs                    []string                                          `json:"visibility_item_ids"`
	RequiredBeforeRevocationVisibility   []string                                          `json:"required_before_revocation_visibility"`
	Checks                               []ProductionReceiptRevocationVisibilityAuditCheck `json:"checks"`
	CheckIDs                             []string                                          `json:"check_ids"`
	Counts                               ProductionReceiptRevocationVisibilityAuditCounts  `json:"counts"`
	RuntimeOwned                         bool                                              `json:"runtime_owned"`
	GoRuntimeBacked                      bool                                              `json:"go_runtime_backed"`
	KDEPolicyOwner                       bool                                              `json:"kde_policy_owner"`
	OfficialDesktopOnly                  bool                                              `json:"official_desktop_only"`
	PlasmaForkRequired                   bool                                              `json:"plasma_fork_required"`
	PlasmaSourceModified                 bool                                              `json:"plasma_source_modified"`
	SystemServiceStarted                 bool                                              `json:"system_service_started"`
	SessionBusClaimed                    bool                                              `json:"session_bus_claimed"`
	ProductionBusClaimed                 bool                                              `json:"production_bus_claimed"`
	ProductionOwnerEnabled               bool                                              `json:"production_owner_enabled"`
	ProductionActivationReady            bool                                              `json:"production_activation_ready"`
	WriteMethodsEnabled                  bool                                              `json:"write_methods_enabled"`
	RuntimeWritesEnabled                 bool                                              `json:"runtime_writes_enabled"`
	ReceiptWriterEnabled                 bool                                              `json:"receipt_writer_enabled"`
	ReceiptPersistenceEnabled            bool                                              `json:"receipt_persistence_enabled"`
	ReceiptLookupWritesEnabled           bool                                              `json:"receipt_lookup_writes_enabled"`
	ReceiptReplayEnabled                 bool                                              `json:"receipt_replay_enabled"`
	ReceiptExpiryWriteEnabled            bool                                              `json:"receipt_expiry_write_enabled"`
	ReceiptRevocationWriteEnabled        bool                                              `json:"receipt_revocation_write_enabled"`
	ReceiptRevocationVisibilityPersisted bool                                              `json:"receipt_revocation_visibility_persisted"`
	DesktopFilesWritten                  bool                                              `json:"desktop_files_written"`
	MIMEAppsWritten                      bool                                              `json:"mimeapps_written"`
	ShellConfigurationWritten            bool                                              `json:"shell_configuration_written"`
	SettingsPersisted                    bool                                              `json:"settings_persisted"`
	KRunnerIndexPersisted                bool                                              `json:"krunner_index_persisted"`
	TaskManagerEntryActive               bool                                              `json:"task_manager_entry_active"`
	KWinRuleApplied                      bool                                              `json:"kwin_rule_applied"`
	LiveTrayBridgeEnabled                bool                                              `json:"live_tray_bridge_enabled"`
	TrayBridgePersisted                  bool                                              `json:"tray_bridge_persisted"`
	NotificationSent                     bool                                              `json:"notification_sent"`
	NotificationDeliveryEnabled          bool                                              `json:"notification_delivery_enabled"`
	CompatibilityCenterPersisted         bool                                              `json:"compatibility_center_persisted"`
	PortalRequestCreated                 bool                                              `json:"portal_request_created"`
	RequestObjectsCreated                bool                                              `json:"request_objects_created"`
	AdapterInvocationEnabled             bool                                              `json:"adapter_invocation_enabled"`
	BackendLaunchEnabled                 bool                                              `json:"backend_launch_enabled"`
	BackendProcessStarted                bool                                              `json:"backend_process_started"`
	SupportBundleExported                bool                                              `json:"support_bundle_exported"`
	SupportCaseCreated                   bool                                              `json:"support_case_created"`
	SnapshotRestoreExecuted              bool                                              `json:"snapshot_restore_executed"`
	StateCleanupExecuted                 bool                                              `json:"state_cleanup_executed"`
	FileContentRead                      bool                                              `json:"file_content_read"`
	FilePathsExposed                     bool                                              `json:"file_paths_exposed"`
	NetworkRequired                      bool                                              `json:"network_required"`
	HostRootModified                     bool                                              `json:"host_root_modified"`
	PrivilegedContainerRequired          bool                                              `json:"privileged_container_required"`
	StateRootPathExposed                 bool                                              `json:"state_root_path_exposed"`
	RawCommandExposed                    bool                                              `json:"raw_command_exposed"`
	RawExecutableExposed                 bool                                              `json:"raw_executable_exposed"`
	BackendDetailsExposed                bool                                              `json:"backend_details_exposed"`
	BlockedActions                       []string                                          `json:"blocked_actions"`
	NextRequirements                     []string                                          `json:"next_requirements"`
	DesktopSafeSummary                   string                                            `json:"desktop_safe_summary"`
}

type ProductionReceiptRevocationVisibilityAuditItem struct {
	ID                            string `json:"id"`
	Audience                      string `json:"audience"`
	StatusKind                    string `json:"status_kind"`
	RequiredEvidence              string `json:"required_evidence"`
	EvidencePresent               bool   `json:"evidence_present"`
	VisibilityModeled             bool   `json:"visibility_modeled"`
	UserVisible                   bool   `json:"user_visible"`
	ProductionGateVisible         bool   `json:"production_gate_visible"`
	KDEStatusVisible              bool   `json:"kde_status_visible"`
	ReviewOnly                    bool   `json:"review_only"`
	RuntimeOwned                  bool   `json:"runtime_owned"`
	GoRuntimeBacked               bool   `json:"go_runtime_backed"`
	KDEPolicyOwner                bool   `json:"kde_policy_owner"`
	ReceiptRevocationWriteEnabled bool   `json:"receipt_revocation_write_enabled"`
	ReceiptExpiryWriteEnabled     bool   `json:"receipt_expiry_write_enabled"`
	ReceiptPersistenceEnabled     bool   `json:"receipt_persistence_enabled"`
	ReceiptLookupWritesEnabled    bool   `json:"receipt_lookup_writes_enabled"`
	ReceiptReplayEnabled          bool   `json:"receipt_replay_enabled"`
	ReceiptAccepted               bool   `json:"receipt_accepted"`
	AuthorizationAccepted         bool   `json:"authorization_accepted"`
	NotificationSent              bool   `json:"notification_sent"`
	NotificationDeliveryEnabled   bool   `json:"notification_delivery_enabled"`
	DesktopFilesWritten           bool   `json:"desktop_files_written"`
	SettingsPersisted             bool   `json:"settings_persisted"`
	ProductionReadiness           bool   `json:"production_readiness"`
	ProductionOwnershipReady      bool   `json:"production_ownership_ready"`
	SideEffectsDisabled           bool   `json:"side_effects_disabled"`
	HostRootModified              bool   `json:"host_root_modified"`
	InternalDetailsExposed        bool   `json:"internal_details_exposed"`
	VisibilityStatus              string `json:"visibility_status"`
	NextRequirement               string `json:"next_requirement"`
}

type ProductionReceiptRevocationVisibilityAuditCheck struct {
	ID      string `json:"id"`
	Status  string `json:"status"`
	Summary string `json:"summary"`
}

type ProductionReceiptRevocationVisibilityAuditCounts struct {
	Total   int `json:"total"`
	Passed  int `json:"passed"`
	Pending int `json:"pending"`
	Blocked int `json:"blocked"`
}

func NewProductionReceiptRevocationVisibilityAuditPreview(root string) (ProductionReceiptRevocationVisibilityAuditPreview, error) {
	version, err := readProductionDBusGateReviewVersion(root)
	if err != nil {
		return ProductionReceiptRevocationVisibilityAuditPreview{}, err
	}
	sources := productionReceiptRevocationVisibilityAuditSources(root)
	items := productionReceiptRevocationVisibilityAuditItems(sources)
	persistenceReady := productionReceiptRevocationVisibilityPersistenceReady(sources.PersistenceThreatReview)
	desktopReady := productionReceiptRevocationVisibilityDesktopReady(sources.DesktopSideEffectReview)
	visibilityReady := persistenceReady && desktopReady && productionReceiptRevocationVisibilityItemsReady(items)
	preview := ProductionReceiptRevocationVisibilityAuditPreview{
		Version:                         version,
		SchemaVersion:                   "xnix.runtime.production_receipt_revocation_visibility_audit.v1",
		RequestType:                     "production-receipt-revocation-visibility-audit-preview",
		AuditType:                       "receipt-expiry-revocation-production-kde-visibility-audit",
		Source:                          "production-receipt-persistence-threat-review-preview+production-desktop-side-effect-review-preview+claude-code-current-dispatch-picks",
		AuditDecision:                   "production-receipt-revocation-visibility-audit-blocked",
		ReceiptSchema:                   "xnix.runtime.production_dbus_human_authorization_receipt.v1",
		OpaqueReceiptID:                 ProductionDBusHumanAuthorizationReceiptID,
		VisibilityAuditRequired:         true,
		VisibilityAuditModeled:          true,
		PersistenceThreatReviewConsumed: persistenceReady,
		DesktopSideEffectReviewConsumed: desktopReady,
		ProductionGateVisibilityModeled: visibilityReady,
		KDEStatusVisibilityModeled:      visibilityReady,
		CallerStateRootRequired:         false,
		ReceiptPresent:                  false,
		ReceiptAccepted:                 false,
		AuthorizationAccepted:           false,
		RevocationVisibilityReady:       visibilityReady,
		ProductionReadiness:             false,
		ProductionOwnershipReady:        false,
		VisibilityItemCount:             len(items),
		RequiredVisibilityItemCount:     5,
		ReadyVisibilityItemCount:        productionReceiptRevocationVisibilityReadyItemCount(items),
		MissingVisibilityItemCount:      productionReceiptRevocationVisibilityMissingItemCount(items),
		VisibilityItems:                 items,
		VisibilityItemIDs:               productionReceiptRevocationVisibilityItemIDs(items),
		RequiredBeforeRevocationVisibility: []string{
			"production-receipt-persistence-threat-review-preview",
			"production-desktop-side-effect-review-preview",
			"explicit redacted status text for production gates",
			"explicit redacted status text for KDE Compatibility Center views",
			"separate operator-approved revocation implementation before any write",
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
			"treat revocation visibility as permission to revoke, expire, accept, persist, replay, or look up receipts",
			"write receipt state, write revocation visibility state, persist Compatibility Center state, or persist settings from this audit",
			"send notifications, create Portal requests, create Runtime request objects, apply KWin rules, or activate live tray bridges",
			"claim production D-Bus ownership, enable Runtime writes, start services, or launch compatibility engines",
			"read file contents, expose paths, expose state-root paths, expose raw commands, expose internal engine details, or mutate host root",
		},
		NextRequirements: []string{
			"Keep revocation and expiry visible only as redacted Runtime-owned review states until receipt storage and policy are implemented separately.",
			"Require a separate operator-approved revocation implementation before any receipt revocation write exists.",
			"Require a separate notification delivery review before any expired or revoked receipt can notify users.",
			"Keep KDE as shell and presentation only; Runtime remains the owner of compatibility policy and production gates.",
		},
		DesktopSafeSummary: "The production receipt revocation visibility audit models how current, expiring, expired, revoked, and missing-review receipt states would appear in production gates and KDE-safe status views, but it writes no receipts, sends no notifications, starts no services, launches no engines, and mutates no host state.",
	}
	preview.RevocationWriteEnabledItemCount = productionReceiptRevocationVisibilityRevocationWriteCount(items)
	preview.ExpiryWriteEnabledItemCount = productionReceiptRevocationVisibilityExpiryWriteCount(items)
	preview.NotificationEnabledItemCount = productionReceiptRevocationVisibilityNotificationCount(items)
	preview.SideEffectItemCount = productionReceiptRevocationVisibilitySideEffectCount(items)
	checks := productionReceiptRevocationVisibilityAuditChecks(preview)
	preview.Checks = checks
	preview.CheckIDs = productionReceiptRevocationVisibilityCheckIDs(checks)
	preview.Counts = countProductionReceiptRevocationVisibilityChecks(checks)
	if preview.Counts.Blocked == 0 {
		preview.AuditDecision = "production-receipt-revocation-visibility-audit-ready-visibility-only"
	}
	if err := validateNoBackendTerms(preview, "production receipt revocation visibility audit preview"); err != nil {
		return ProductionReceiptRevocationVisibilityAuditPreview{}, err
	}
	return preview, nil
}

type productionReceiptRevocationVisibilityAuditSourceSet struct {
	PersistenceThreatReview string
	DesktopSideEffectReview string
	DispatchSheet           string
}

func productionReceiptRevocationVisibilityAuditSources(root string) productionReceiptRevocationVisibilityAuditSourceSet {
	return productionReceiptRevocationVisibilityAuditSourceSet{
		PersistenceThreatReview: productionAuthorizationReadSources(root, []string{"internal/runtime/owner/production_receipt_persistence_threat_review.go"}),
		DesktopSideEffectReview: productionAuthorizationReadSources(root, []string{"internal/runtime/owner/production_desktop_side_effect_review.go"}),
		DispatchSheet:           productionAuthorizationReadSources(root, []string{"docs/claude-code-current-dispatch-picks.md"}),
	}
}

func productionReceiptRevocationVisibilityAuditItems(sources productionReceiptRevocationVisibilityAuditSourceSet) []ProductionReceiptRevocationVisibilityAuditItem {
	combined := sources.PersistenceThreatReview + sources.DesktopSideEffectReview + sources.DispatchSheet
	return []ProductionReceiptRevocationVisibilityAuditItem{
		productionReceiptRevocationVisibilityItem("receipt-current-visible", "production-gate+kde-status", "current", "future current receipt state stays visible without accepting a receipt", combined, []string{"ReceiptRequired", "ReceiptPresent", "ProductionReadiness"}, "define redacted current-state text before any accepted receipt is displayed"),
		productionReceiptRevocationVisibilityItem("receipt-expiring-visible", "production-gate+kde-status", "expiring", "future expiring receipt state stays visible without expiry writes", combined, []string{"expiry", "ReceiptExpiryWriteEnabled", "KDE-safe status views"}, "define redacted expiring-state text before any receipt lifetime is displayed"),
		productionReceiptRevocationVisibilityItem("receipt-expired-visible", "production-gate+kde-status", "expired", "future expired receipt state blocks production without expiry writes", combined, []string{"expiry", "production gates", "ReceiptExpiryWriteEnabled"}, "define production gate copy for expired receipts before any expiry policy is enforced"),
		productionReceiptRevocationVisibilityItem("receipt-revoked-visible", "production-gate+kde-status", "revoked", "future revoked receipt state blocks production without revocation writes", combined, []string{"revocation", "ReceiptRevocationWriteEnabled", "revocation writes"}, "define production gate copy for revoked receipts before any revocation policy is enforced"),
		productionReceiptRevocationVisibilityItem("receipt-missing-review-visible", "production-gate+kde-status", "missing-review", "future missing receipt review state stays visible without notifications", combined, []string{"missing-threat-evidence", "host-boundary-closed", "notifications"}, "define redacted missing-review guidance before any notification delivery is enabled"),
	}
}

func productionReceiptRevocationVisibilityItem(id string, audience string, statusKind string, evidence string, source string, tokens []string, nextRequirement string) ProductionReceiptRevocationVisibilityAuditItem {
	ready := productionAuthorizationHasAll(source, tokens)
	status := "missing-visibility-evidence"
	if ready {
		status = "visibility-modeled-writes-disabled"
	}
	return ProductionReceiptRevocationVisibilityAuditItem{
		ID:                            id,
		Audience:                      audience,
		StatusKind:                    statusKind,
		RequiredEvidence:              evidence,
		EvidencePresent:               ready,
		VisibilityModeled:             ready,
		UserVisible:                   ready,
		ProductionGateVisible:         ready,
		KDEStatusVisible:              ready,
		ReviewOnly:                    true,
		RuntimeOwned:                  true,
		GoRuntimeBacked:               true,
		KDEPolicyOwner:                false,
		ReceiptRevocationWriteEnabled: false,
		ReceiptExpiryWriteEnabled:     false,
		ReceiptPersistenceEnabled:     false,
		ReceiptLookupWritesEnabled:    false,
		ReceiptReplayEnabled:          false,
		ReceiptAccepted:               false,
		AuthorizationAccepted:         false,
		NotificationSent:              false,
		NotificationDeliveryEnabled:   false,
		DesktopFilesWritten:           false,
		SettingsPersisted:             false,
		ProductionReadiness:           false,
		ProductionOwnershipReady:      false,
		SideEffectsDisabled:           true,
		HostRootModified:              false,
		InternalDetailsExposed:        false,
		VisibilityStatus:              status,
		NextRequirement:               nextRequirement,
	}
}

func productionReceiptRevocationVisibilityAuditChecks(preview ProductionReceiptRevocationVisibilityAuditPreview) []ProductionReceiptRevocationVisibilityAuditCheck {
	return []ProductionReceiptRevocationVisibilityAuditCheck{
		productionReceiptRevocationVisibilityCheck("persistence-threat-review-consumed", productionAuthorizationPassBlocked(preview.PersistenceThreatReviewConsumed), "The visibility audit consumes the receipt persistence threat review."),
		productionReceiptRevocationVisibilityCheck("desktop-side-effect-review-consumed", productionAuthorizationPassBlocked(preview.DesktopSideEffectReviewConsumed), "The visibility audit consumes the KDE desktop side-effect review before exposing status views."),
		productionReceiptRevocationVisibilityCheck("revocation-visibility-modeled-only", productionAuthorizationPassBlocked(preview.VisibilityAuditRequired && preview.VisibilityAuditModeled && preview.RevocationVisibilityReady && preview.ProductionGateVisibilityModeled && preview.KDEStatusVisibilityModeled && !preview.ReceiptPresent && !preview.ReceiptAccepted && !preview.AuthorizationAccepted), "Revocation visibility is modeled without accepting or persisting a receipt."),
		productionReceiptRevocationVisibilityCheck("five-visibility-items-present", productionAuthorizationPassBlocked(preview.VisibilityItemCount == 5 && preview.RequiredVisibilityItemCount == 5 && preview.MissingVisibilityItemCount == 0), "Current, expiring, expired, revoked, and missing-review states are modeled."),
		productionReceiptRevocationVisibilityCheck("visibility-items-ready-writes-disabled", productionAuthorizationPassBlocked(preview.ReadyVisibilityItemCount == 5 && productionReceiptRevocationVisibilityItemsReady(preview.VisibilityItems)), "Every visibility item is ready while receipt writes, notifications, and side effects remain disabled."),
		productionReceiptRevocationVisibilityCheck("revocation-expiry-writes-disabled", productionAuthorizationPassBlocked(!preview.ReceiptRevocationWriteEnabled && !preview.ReceiptExpiryWriteEnabled && !preview.ReceiptPersistenceEnabled && !preview.ReceiptLookupWritesEnabled && !preview.ReceiptReplayEnabled && preview.RevocationWriteEnabledItemCount == 0 && preview.ExpiryWriteEnabledItemCount == 0), "Revocation, expiry, persistence, lookup, and replay writes remain disabled."),
		productionReceiptRevocationVisibilityCheck("notifications-and-desktop-side-effects-disabled", productionAuthorizationPassBlocked(!preview.NotificationSent && !preview.NotificationDeliveryEnabled && !preview.DesktopFilesWritten && !preview.SettingsPersisted && !preview.CompatibilityCenterPersisted && !preview.PortalRequestCreated && !preview.RequestObjectsCreated && preview.NotificationEnabledItemCount == 0 && preview.SideEffectItemCount == 0), "Notifications, desktop writes, settings persistence, Compatibility Center persistence, Portal requests, and Runtime request objects remain disabled."),
		productionReceiptRevocationVisibilityCheck("production-ownership-disabled", productionAuthorizationPassBlocked(!preview.SystemServiceStarted && !preview.SessionBusClaimed && !preview.ProductionBusClaimed && !preview.ProductionOwnerEnabled && !preview.ProductionActivationReady && !preview.ProductionReadiness && !preview.ProductionOwnershipReady), "Production service start, bus claims, ownership, activation, and readiness remain disabled."),
		productionReceiptRevocationVisibilityCheck("host-boundary-closed", productionAuthorizationPassBlocked(!preview.FileContentRead && !preview.FilePathsExposed && !preview.NetworkRequired && !preview.HostRootModified && !preview.PrivilegedContainerRequired && !preview.StateRootPathExposed && !preview.RawCommandExposed && !preview.RawExecutableExposed && !preview.BackendDetailsExposed && productionReceiptRevocationVisibilityItemsKeepHostClosed(preview.VisibilityItems)), "Unsafe data exposure, network, privilege, path exposure, internal detail exposure, and host mutation gates remain closed."),
	}
}

func productionReceiptRevocationVisibilityCheck(id string, status string, summary string) ProductionReceiptRevocationVisibilityAuditCheck {
	return ProductionReceiptRevocationVisibilityAuditCheck{ID: id, Status: status, Summary: summary}
}

func productionReceiptRevocationVisibilityPersistenceReady(source string) bool {
	return productionAuthorizationHasAll(source, []string{
		"production-receipt-persistence-threat-review-preview",
		"receipt-persistence-expiry-revocation-replay-threat-review",
		"production-receipt-persistence-threat-review-ready-persistence-disabled",
		"expiry-policy",
		"revocation-policy",
		"ReceiptExpiryWriteEnabled",
		"ReceiptRevocationWriteEnabled",
		"host-boundary-closed",
	})
}

func productionReceiptRevocationVisibilityDesktopReady(source string) bool {
	return productionAuthorizationHasAll(source, []string{
		"production-desktop-side-effect-review-preview",
		"kde-production-desktop-side-effect-review",
		"production-desktop-side-effect-review-ready-side-effects-disabled",
		"compatibility-center",
		"notifications-and-requests-disabled",
		"NotificationDeliveryEnabled",
	})
}

func productionReceiptRevocationVisibilityReadyItemCount(items []ProductionReceiptRevocationVisibilityAuditItem) int {
	count := 0
	for _, item := range items {
		if item.EvidencePresent {
			count++
		}
	}
	return count
}

func productionReceiptRevocationVisibilityMissingItemCount(items []ProductionReceiptRevocationVisibilityAuditItem) int {
	count := 0
	for _, item := range items {
		if !item.EvidencePresent {
			count++
		}
	}
	return count
}

func productionReceiptRevocationVisibilityRevocationWriteCount(items []ProductionReceiptRevocationVisibilityAuditItem) int {
	count := 0
	for _, item := range items {
		if item.ReceiptRevocationWriteEnabled {
			count++
		}
	}
	return count
}

func productionReceiptRevocationVisibilityExpiryWriteCount(items []ProductionReceiptRevocationVisibilityAuditItem) int {
	count := 0
	for _, item := range items {
		if item.ReceiptExpiryWriteEnabled {
			count++
		}
	}
	return count
}

func productionReceiptRevocationVisibilityNotificationCount(items []ProductionReceiptRevocationVisibilityAuditItem) int {
	count := 0
	for _, item := range items {
		if item.NotificationSent || item.NotificationDeliveryEnabled {
			count++
		}
	}
	return count
}

func productionReceiptRevocationVisibilitySideEffectCount(items []ProductionReceiptRevocationVisibilityAuditItem) int {
	count := 0
	for _, item := range items {
		if item.DesktopFilesWritten || item.SettingsPersisted || item.HostRootModified || item.InternalDetailsExposed {
			count++
		}
	}
	return count
}

func productionReceiptRevocationVisibilityItemIDs(items []ProductionReceiptRevocationVisibilityAuditItem) []string {
	ids := make([]string, 0, len(items))
	for _, item := range items {
		ids = append(ids, item.ID)
	}
	return ids
}

func productionReceiptRevocationVisibilityCheckIDs(checks []ProductionReceiptRevocationVisibilityAuditCheck) []string {
	ids := make([]string, 0, len(checks))
	for _, check := range checks {
		ids = append(ids, check.ID)
	}
	return ids
}

func countProductionReceiptRevocationVisibilityChecks(checks []ProductionReceiptRevocationVisibilityAuditCheck) ProductionReceiptRevocationVisibilityAuditCounts {
	counts := ProductionReceiptRevocationVisibilityAuditCounts{Total: len(checks)}
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

func productionReceiptRevocationVisibilityItemsReady(items []ProductionReceiptRevocationVisibilityAuditItem) bool {
	if len(items) != 5 {
		return false
	}
	for _, item := range items {
		if !item.EvidencePresent || !item.VisibilityModeled || !item.UserVisible || !item.ProductionGateVisible || !item.KDEStatusVisible || item.VisibilityStatus != "visibility-modeled-writes-disabled" {
			return false
		}
		if item.ReceiptRevocationWriteEnabled || item.ReceiptExpiryWriteEnabled || item.ReceiptPersistenceEnabled || item.ReceiptLookupWritesEnabled || item.ReceiptReplayEnabled || item.ReceiptAccepted || item.AuthorizationAccepted || item.NotificationSent || item.NotificationDeliveryEnabled || item.DesktopFilesWritten || item.SettingsPersisted || item.ProductionReadiness || item.ProductionOwnershipReady || item.KDEPolicyOwner || !item.ReviewOnly || !item.SideEffectsDisabled {
			return false
		}
	}
	return true
}

func productionReceiptRevocationVisibilityItemsKeepHostClosed(items []ProductionReceiptRevocationVisibilityAuditItem) bool {
	for _, item := range items {
		if item.HostRootModified || item.InternalDetailsExposed {
			return false
		}
	}
	return true
}
