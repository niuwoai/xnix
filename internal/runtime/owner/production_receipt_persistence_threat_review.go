package owner

type ProductionReceiptPersistenceThreatReviewPreview struct {
	Version                           string                                          `json:"version"`
	SchemaVersion                     string                                          `json:"schema_version"`
	RequestType                       string                                          `json:"request_type"`
	ReviewType                        string                                          `json:"review_type"`
	Source                            string                                          `json:"source"`
	ReviewDecision                    string                                          `json:"review_decision"`
	ReceiptSchema                     string                                          `json:"receipt_schema"`
	OpaqueReceiptID                   string                                          `json:"opaque_receipt_id"`
	ReceiptRequired                   bool                                            `json:"receipt_required"`
	ReceiptPresent                    bool                                            `json:"receipt_present"`
	ReceiptAccepted                   bool                                            `json:"receipt_accepted"`
	PersistenceThreatReviewRequired   bool                                            `json:"persistence_threat_review_required"`
	PersistenceThreatReviewModeled    bool                                            `json:"persistence_threat_review_modeled"`
	WriterAuthorizationReviewConsumed bool                                            `json:"writer_authorization_review_consumed"`
	OwnerManagedOpaqueBoundaryReady   bool                                            `json:"owner_managed_opaque_boundary_ready"`
	CallerStateRootRequired           bool                                            `json:"caller_state_root_required"`
	AuthorizationAccepted             bool                                            `json:"authorization_accepted"`
	PersistenceThreatReviewReady      bool                                            `json:"persistence_threat_review_ready"`
	ProductionReadiness               bool                                            `json:"production_readiness"`
	ProductionOwnershipReady          bool                                            `json:"production_ownership_ready"`
	ThreatItemCount                   int                                             `json:"threat_item_count"`
	RequiredThreatItemCount           int                                             `json:"required_threat_item_count"`
	ReadyThreatItemCount              int                                             `json:"ready_threat_item_count"`
	MissingThreatItemCount            int                                             `json:"missing_threat_item_count"`
	PersistenceEnabledThreatCount     int                                             `json:"persistence_enabled_threat_count"`
	ReplayEnabledThreatCount          int                                             `json:"replay_enabled_threat_count"`
	AcceptanceEnabledThreatCount      int                                             `json:"acceptance_enabled_threat_count"`
	SideEffectThreatCount             int                                             `json:"side_effect_threat_count"`
	ThreatItems                       []ProductionReceiptPersistenceThreatReviewItem  `json:"threat_items"`
	ThreatItemIDs                     []string                                        `json:"threat_item_ids"`
	RequiredBeforeReceiptPersistence  []string                                        `json:"required_before_receipt_persistence"`
	Checks                            []ProductionReceiptPersistenceThreatReviewCheck `json:"checks"`
	CheckIDs                          []string                                        `json:"check_ids"`
	Counts                            ProductionReceiptPersistenceThreatReviewCounts  `json:"counts"`
	RuntimeOwned                      bool                                            `json:"runtime_owned"`
	GoRuntimeBacked                   bool                                            `json:"go_runtime_backed"`
	KDEPolicyOwner                    bool                                            `json:"kde_policy_owner"`
	SystemServiceStarted              bool                                            `json:"system_service_started"`
	SessionBusClaimed                 bool                                            `json:"session_bus_claimed"`
	ProductionBusClaimed              bool                                            `json:"production_bus_claimed"`
	ProductionOwnerEnabled            bool                                            `json:"production_owner_enabled"`
	ProductionActivationReady         bool                                            `json:"production_activation_ready"`
	WriteMethodsEnabled               bool                                            `json:"write_methods_enabled"`
	RuntimeWritesEnabled              bool                                            `json:"runtime_writes_enabled"`
	ReceiptWriterEnabled              bool                                            `json:"receipt_writer_enabled"`
	ReceiptPersistenceEnabled         bool                                            `json:"receipt_persistence_enabled"`
	ReceiptLookupWritesEnabled        bool                                            `json:"receipt_lookup_writes_enabled"`
	ReceiptReplayEnabled              bool                                            `json:"receipt_replay_enabled"`
	ReceiptExpiryWriteEnabled         bool                                            `json:"receipt_expiry_write_enabled"`
	ReceiptRevocationWriteEnabled     bool                                            `json:"receipt_revocation_write_enabled"`
	DesktopFilesWritten               bool                                            `json:"desktop_files_written"`
	MIMEAppsWritten                   bool                                            `json:"mimeapps_written"`
	ShellConfigurationWritten         bool                                            `json:"shell_configuration_written"`
	SettingsPersisted                 bool                                            `json:"settings_persisted"`
	NotificationSent                  bool                                            `json:"notification_sent"`
	NotificationDeliveryEnabled       bool                                            `json:"notification_delivery_enabled"`
	PortalRequestCreated              bool                                            `json:"portal_request_created"`
	RequestObjectsCreated             bool                                            `json:"request_objects_created"`
	AdapterInvocationEnabled          bool                                            `json:"adapter_invocation_enabled"`
	BackendLaunchEnabled              bool                                            `json:"backend_launch_enabled"`
	BackendProcessStarted             bool                                            `json:"backend_process_started"`
	SupportBundleExported             bool                                            `json:"support_bundle_exported"`
	SupportCaseCreated                bool                                            `json:"support_case_created"`
	SnapshotRestoreExecuted           bool                                            `json:"snapshot_restore_executed"`
	StateCleanupExecuted              bool                                            `json:"state_cleanup_executed"`
	FileContentRead                   bool                                            `json:"file_content_read"`
	FilePathsExposed                  bool                                            `json:"file_paths_exposed"`
	NetworkRequired                   bool                                            `json:"network_required"`
	HostRootModified                  bool                                            `json:"host_root_modified"`
	PrivilegedContainerRequired       bool                                            `json:"privileged_container_required"`
	StateRootPathExposed              bool                                            `json:"state_root_path_exposed"`
	RawCommandExposed                 bool                                            `json:"raw_command_exposed"`
	RawExecutableExposed              bool                                            `json:"raw_executable_exposed"`
	BackendDetailsExposed             bool                                            `json:"backend_details_exposed"`
	BlockedActions                    []string                                        `json:"blocked_actions"`
	NextRequirements                  []string                                        `json:"next_requirements"`
	DesktopSafeSummary                string                                          `json:"desktop_safe_summary"`
}

type ProductionReceiptPersistenceThreatReviewItem struct {
	ID                            string `json:"id"`
	ThreatArea                    string `json:"threat_area"`
	RequiredEvidence              string `json:"required_evidence"`
	EvidencePresent               bool   `json:"evidence_present"`
	ThreatModeled                 bool   `json:"threat_modeled"`
	ReceiptPersistenceEnabled     bool   `json:"receipt_persistence_enabled"`
	ReceiptLookupWritesEnabled    bool   `json:"receipt_lookup_writes_enabled"`
	ReceiptReplayEnabled          bool   `json:"receipt_replay_enabled"`
	ReceiptExpiryWriteEnabled     bool   `json:"receipt_expiry_write_enabled"`
	ReceiptRevocationWriteEnabled bool   `json:"receipt_revocation_write_enabled"`
	ReceiptAccepted               bool   `json:"receipt_accepted"`
	AuthorizationAccepted         bool   `json:"authorization_accepted"`
	ProductionReadiness           bool   `json:"production_readiness"`
	ProductionOwnershipReady      bool   `json:"production_ownership_ready"`
	RuntimeOwned                  bool   `json:"runtime_owned"`
	GoRuntimeBacked               bool   `json:"go_runtime_backed"`
	KDEPolicyOwner                bool   `json:"kde_policy_owner"`
	ReviewOnly                    bool   `json:"review_only"`
	SideEffectsDisabled           bool   `json:"side_effects_disabled"`
	HostRootModified              bool   `json:"host_root_modified"`
	InternalDetailsExposed        bool   `json:"internal_details_exposed"`
	ThreatStatus                  string `json:"threat_status"`
	NextRequirement               string `json:"next_requirement"`
}

type ProductionReceiptPersistenceThreatReviewCheck struct {
	ID      string `json:"id"`
	Status  string `json:"status"`
	Summary string `json:"summary"`
}

type ProductionReceiptPersistenceThreatReviewCounts struct {
	Total   int `json:"total"`
	Passed  int `json:"passed"`
	Pending int `json:"pending"`
	Blocked int `json:"blocked"`
}

func NewProductionReceiptPersistenceThreatReviewPreview(root string) (ProductionReceiptPersistenceThreatReviewPreview, error) {
	version, err := readProductionDBusGateReviewVersion(root)
	if err != nil {
		return ProductionReceiptPersistenceThreatReviewPreview{}, err
	}
	sources := productionReceiptPersistenceThreatReviewSources(root)
	items := productionReceiptPersistenceThreatReviewItems(sources)
	writerReviewReady := productionReceiptPersistenceWriterReviewReady(sources.WriterAuthorizationReview)
	preview := ProductionReceiptPersistenceThreatReviewPreview{
		Version:                           version,
		SchemaVersion:                     "xnix.runtime.production_receipt_persistence_threat_review.v1",
		RequestType:                       "production-receipt-persistence-threat-review-preview",
		ReviewType:                        "receipt-persistence-expiry-revocation-replay-threat-review",
		Source:                            "production-receipt-writer-authorization-review-preview+production-receipt-acceptance-propagation-preflight-preview",
		ReviewDecision:                    "production-receipt-persistence-threat-review-blocked",
		ReceiptSchema:                     "xnix.runtime.production_dbus_human_authorization_receipt.v1",
		OpaqueReceiptID:                   ProductionDBusHumanAuthorizationReceiptID,
		ReceiptRequired:                   true,
		ReceiptPresent:                    false,
		ReceiptAccepted:                   false,
		PersistenceThreatReviewRequired:   true,
		PersistenceThreatReviewModeled:    true,
		WriterAuthorizationReviewConsumed: writerReviewReady,
		OwnerManagedOpaqueBoundaryReady:   writerReviewReady,
		CallerStateRootRequired:           false,
		AuthorizationAccepted:             false,
		PersistenceThreatReviewReady:      writerReviewReady && productionReceiptPersistenceAllThreatItemsReady(items),
		ProductionReadiness:               false,
		ProductionOwnershipReady:          false,
		ThreatItemCount:                   len(items),
		RequiredThreatItemCount:           5,
		ReadyThreatItemCount:              productionReceiptPersistenceReadyThreatItemCount(items),
		MissingThreatItemCount:            productionReceiptPersistenceMissingThreatItemCount(items),
		PersistenceEnabledThreatCount:     0,
		ReplayEnabledThreatCount:          0,
		AcceptanceEnabledThreatCount:      0,
		SideEffectThreatCount:             0,
		ThreatItems:                       items,
		ThreatItemIDs:                     productionReceiptPersistenceThreatItemIDs(items),
		RequiredBeforeReceiptPersistence: []string{
			"production-receipt-writer-authorization-review-preview",
			"production-receipt-acceptance-propagation-preflight-preview",
			"separate encrypted receipt storage design outside this review",
			"separate expiry and revocation policy outside this review",
			"separate replay protection design outside this review",
			"separate production service ownership proof outside this review",
		},
		RuntimeOwned:                  true,
		GoRuntimeBacked:               true,
		KDEPolicyOwner:                false,
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
		NotificationSent:              false,
		NotificationDeliveryEnabled:   false,
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
			"treat persistence threat review as permission to store receipts",
			"write, persist, accept, replay, expire, revoke, or look up authorization receipts from this review",
			"claim production D-Bus ownership, start services, or enable Runtime writes from this review",
			"enable desktop writes, notifications, Portal requests, request objects, adapter invocation, or compatibility engine launch",
			"export support bundles, create support cases, restore snapshots, clean state, read file contents, or expose paths",
			"require network, require privileged containers, expose raw commands, expose internal engine details, or mutate host root",
		},
		NextRequirements: []string{
			"Keep receipt persistence outside preview commands until encrypted storage, expiry, revocation, and replay protection are separately reviewed.",
			"Require explicit operator authorization and production service ownership proof before any accepted receipt can be persisted.",
			"Keep service start, bus claim, write dispatch, desktop side effects, support side effects, restore, cleanup, launch, and host mutation disabled.",
			"Add a separate persistence implementation plan only after this threat review remains read-only and fail-closed.",
		},
		DesktopSafeSummary: "The production receipt persistence threat review models storage, expiry, revocation, replay protection, and audit visibility requirements for future authorization receipts, but it does not persist receipts, accept authorization, claim a bus, start services, enable writes, emit desktop side effects, launch engines, or mutate host state.",
	}
	checks := productionReceiptPersistenceThreatReviewChecks(preview)
	preview.Checks = checks
	preview.CheckIDs = productionReceiptPersistenceCheckIDs(checks)
	preview.Counts = countProductionReceiptPersistenceChecks(checks)
	if preview.Counts.Blocked == 0 {
		preview.ReviewDecision = "production-receipt-persistence-threat-review-ready-persistence-disabled"
	}
	if err := validateNoBackendTerms(preview, "production receipt persistence threat review preview"); err != nil {
		return ProductionReceiptPersistenceThreatReviewPreview{}, err
	}
	return preview, nil
}

type productionReceiptPersistenceThreatReviewSourceSet struct {
	WriterAuthorizationReview string
	AcceptancePropagation     string
	DispatchSheet             string
}

func productionReceiptPersistenceThreatReviewSources(root string) productionReceiptPersistenceThreatReviewSourceSet {
	return productionReceiptPersistenceThreatReviewSourceSet{
		WriterAuthorizationReview: productionAuthorizationReadSources(root, []string{"internal/runtime/owner/production_receipt_writer_authorization_review.go"}),
		AcceptancePropagation:     productionAuthorizationReadSources(root, []string{"internal/runtime/owner/production_receipt_acceptance_propagation_preflight.go"}),
		DispatchSheet:             productionAuthorizationReadSources(root, []string{"docs/claude-code-current-dispatch-picks.md"}),
	}
}

func productionReceiptPersistenceThreatReviewItems(sources productionReceiptPersistenceThreatReviewSourceSet) []ProductionReceiptPersistenceThreatReviewItem {
	combined := sources.WriterAuthorizationReview + sources.AcceptancePropagation + sources.DispatchSheet
	return []ProductionReceiptPersistenceThreatReviewItem{
		productionReceiptPersistenceThreatItem("storage-confidentiality", "storage", "future encrypted receipt storage remains design-only", combined, []string{"stored", "ReceiptPersistenceEnabled", "FilePathsExposed"}, "define encrypted storage before any receipt file can exist"),
		productionReceiptPersistenceThreatItem("expiry-policy", "expiry", "future receipt expiry stays disabled until policy exists", combined, []string{"expiry", "ReceiptPersistenceEnabled", "WriterAuthorizationRequired"}, "define expiry policy before any receipt lifetime can be stored"),
		productionReceiptPersistenceThreatItem("revocation-policy", "revocation", "future receipt revocation stays disabled until policy exists", combined, []string{"revocation", "ReceiptLookupWritesEnabled", "OperatorActionRequired"}, "define revocation policy before any receipt lookup write can exist"),
		productionReceiptPersistenceThreatItem("replay-protection", "replay", "future receipt replay protection stays disabled until design exists", combined, []string{"replay", "ReceiptReplayEnabled", "AcceptanceEnabledReviewItemCount"}, "define replay protection before any receipt can be accepted or replayed"),
		productionReceiptPersistenceThreatItem("audit-visibility", "audit", "future receipt audit visibility stays redacted and host-safe", combined, []string{"host-boundary-closed", "StateRootPathExposed", "BackendDetailsExposed"}, "define redacted audit visibility before any receipt audit trail is persisted"),
	}
}

func productionReceiptPersistenceThreatItem(id string, area string, evidence string, source string, tokens []string, nextRequirement string) ProductionReceiptPersistenceThreatReviewItem {
	ready := productionAuthorizationHasAll(source, tokens)
	status := "missing-threat-evidence"
	if ready {
		status = "threat-modeled-persistence-disabled"
	}
	return ProductionReceiptPersistenceThreatReviewItem{
		ID:                            id,
		ThreatArea:                    area,
		RequiredEvidence:              evidence,
		EvidencePresent:               ready,
		ThreatModeled:                 ready,
		ReceiptPersistenceEnabled:     false,
		ReceiptLookupWritesEnabled:    false,
		ReceiptReplayEnabled:          false,
		ReceiptExpiryWriteEnabled:     false,
		ReceiptRevocationWriteEnabled: false,
		ReceiptAccepted:               false,
		AuthorizationAccepted:         false,
		ProductionReadiness:           false,
		ProductionOwnershipReady:      false,
		RuntimeOwned:                  true,
		GoRuntimeBacked:               true,
		KDEPolicyOwner:                false,
		ReviewOnly:                    true,
		SideEffectsDisabled:           true,
		HostRootModified:              false,
		InternalDetailsExposed:        false,
		ThreatStatus:                  status,
		NextRequirement:               nextRequirement,
	}
}

func productionReceiptPersistenceThreatReviewChecks(preview ProductionReceiptPersistenceThreatReviewPreview) []ProductionReceiptPersistenceThreatReviewCheck {
	return []ProductionReceiptPersistenceThreatReviewCheck{
		productionReceiptPersistenceCheck("writer-authorization-review-consumed", productionAuthorizationPassBlocked(preview.WriterAuthorizationReviewConsumed && preview.OwnerManagedOpaqueBoundaryReady && preview.ReceiptSchema == "xnix.runtime.production_dbus_human_authorization_receipt.v1" && preview.OpaqueReceiptID == ProductionDBusHumanAuthorizationReceiptID), "The threat review consumes the writer authorization review and opaque receipt boundary."),
		productionReceiptPersistenceCheck("persistence-threat-modeled-only", productionAuthorizationPassBlocked(preview.PersistenceThreatReviewRequired && preview.PersistenceThreatReviewModeled && !preview.ReceiptPresent && !preview.ReceiptAccepted && !preview.AuthorizationAccepted), "The review models persistence threats without accepting or persisting a receipt."),
		productionReceiptPersistenceCheck("five-threat-items-present", productionAuthorizationPassBlocked(preview.ThreatItemCount == 5 && preview.RequiredThreatItemCount == 5 && preview.MissingThreatItemCount == 0), "The review tracks storage, expiry, revocation, replay, and audit visibility threats."),
		productionReceiptPersistenceCheck("threat-items-ready-persistence-disabled", productionAuthorizationPassBlocked(preview.ReadyThreatItemCount == 5 && productionReceiptPersistenceAllThreatItemsReady(preview.ThreatItems)), "Every threat item is ready while persistence, lookup writes, replay, expiry, revocation, and acceptance remain disabled."),
		productionReceiptPersistenceCheck("receipt-persistence-disabled", productionAuthorizationPassBlocked(!preview.ReceiptWriterEnabled && !preview.ReceiptPersistenceEnabled && !preview.ReceiptLookupWritesEnabled && !preview.ReceiptReplayEnabled && !preview.ReceiptExpiryWriteEnabled && !preview.ReceiptRevocationWriteEnabled && preview.PersistenceEnabledThreatCount == 0 && preview.ReplayEnabledThreatCount == 0 && preview.AcceptanceEnabledThreatCount == 0), "Receipt writer, persistence, lookup writes, replay, expiry, revocation, and acceptance remain disabled."),
		productionReceiptPersistenceCheck("production-ownership-disabled", productionAuthorizationPassBlocked(!preview.SystemServiceStarted && !preview.SessionBusClaimed && !preview.ProductionBusClaimed && !preview.ProductionOwnerEnabled && !preview.ProductionActivationReady && !preview.ProductionReadiness && !preview.ProductionOwnershipReady), "Service start, session bus claim, production bus claim, and production ownership remain disabled."),
		productionReceiptPersistenceCheck("runtime-side-effects-disabled", productionAuthorizationPassBlocked(!preview.WriteMethodsEnabled && !preview.RuntimeWritesEnabled && !preview.DesktopFilesWritten && !preview.MIMEAppsWritten && !preview.ShellConfigurationWritten && !preview.SettingsPersisted && !preview.NotificationSent && !preview.NotificationDeliveryEnabled && !preview.PortalRequestCreated && !preview.RequestObjectsCreated && !preview.AdapterInvocationEnabled && !preview.BackendLaunchEnabled && !preview.BackendProcessStarted && !preview.SupportBundleExported && !preview.SupportCaseCreated && !preview.SnapshotRestoreExecuted && !preview.StateCleanupExecuted && preview.SideEffectThreatCount == 0), "Runtime writes, desktop, Portal, support, restore, cleanup, adapter invocation, and launch side effects remain disabled."),
		productionReceiptPersistenceCheck("host-boundary-closed", productionAuthorizationPassBlocked(!preview.FileContentRead && !preview.FilePathsExposed && !preview.NetworkRequired && !preview.HostRootModified && !preview.PrivilegedContainerRequired && !preview.StateRootPathExposed && !preview.RawCommandExposed && !preview.RawExecutableExposed && !preview.BackendDetailsExposed && productionReceiptPersistenceItemsKeepHostClosed(preview.ThreatItems)), "Unsafe data exposure, network, privilege, internal detail exposure, and host mutation gates remain closed."),
	}
}

func productionReceiptPersistenceCheck(id string, status string, summary string) ProductionReceiptPersistenceThreatReviewCheck {
	return ProductionReceiptPersistenceThreatReviewCheck{ID: id, Status: status, Summary: summary}
}

func productionReceiptPersistenceCheckIDs(checks []ProductionReceiptPersistenceThreatReviewCheck) []string {
	ids := make([]string, 0, len(checks))
	for _, check := range checks {
		ids = append(ids, check.ID)
	}
	return ids
}

func countProductionReceiptPersistenceChecks(checks []ProductionReceiptPersistenceThreatReviewCheck) ProductionReceiptPersistenceThreatReviewCounts {
	counts := ProductionReceiptPersistenceThreatReviewCounts{Total: len(checks)}
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

func productionReceiptPersistenceWriterReviewReady(source string) bool {
	return productionAuthorizationHasAll(source, []string{
		"production-receipt-writer-authorization-review-preview",
		"receipt-writer-operator-authorization-boundary-review",
		"production-receipt-writer-authorization-review-ready-writes-disabled",
		"writer-authorization-modeled-only",
		"receipt-writes-disabled",
		"runtime-side-effects-disabled",
		"host-boundary-closed",
	})
}

func productionReceiptPersistenceReadyThreatItemCount(items []ProductionReceiptPersistenceThreatReviewItem) int {
	count := 0
	for _, item := range items {
		if item.EvidencePresent {
			count++
		}
	}
	return count
}

func productionReceiptPersistenceMissingThreatItemCount(items []ProductionReceiptPersistenceThreatReviewItem) int {
	count := 0
	for _, item := range items {
		if !item.EvidencePresent {
			count++
		}
	}
	return count
}

func productionReceiptPersistenceThreatItemIDs(items []ProductionReceiptPersistenceThreatReviewItem) []string {
	ids := make([]string, 0, len(items))
	for _, item := range items {
		ids = append(ids, item.ID)
	}
	return ids
}

func productionReceiptPersistenceAllThreatItemsReady(items []ProductionReceiptPersistenceThreatReviewItem) bool {
	if len(items) != 5 {
		return false
	}
	for _, item := range items {
		if !item.EvidencePresent || !item.ThreatModeled || item.ThreatStatus != "threat-modeled-persistence-disabled" {
			return false
		}
		if item.ReceiptPersistenceEnabled || item.ReceiptLookupWritesEnabled || item.ReceiptReplayEnabled || item.ReceiptExpiryWriteEnabled || item.ReceiptRevocationWriteEnabled || item.ReceiptAccepted || item.AuthorizationAccepted || item.ProductionReadiness || item.ProductionOwnershipReady || item.KDEPolicyOwner || !item.ReviewOnly || !item.SideEffectsDisabled {
			return false
		}
	}
	return true
}

func productionReceiptPersistenceItemsKeepHostClosed(items []ProductionReceiptPersistenceThreatReviewItem) bool {
	for _, item := range items {
		if item.HostRootModified || item.InternalDetailsExposed {
			return false
		}
	}
	return true
}
