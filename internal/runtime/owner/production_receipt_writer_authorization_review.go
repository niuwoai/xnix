package owner

type ProductionReceiptWriterAuthorizationReviewPreview struct {
	Version                            string                                            `json:"version"`
	SchemaVersion                      string                                            `json:"schema_version"`
	RequestType                        string                                            `json:"request_type"`
	ReviewType                         string                                            `json:"review_type"`
	Source                             string                                            `json:"source"`
	ReviewDecision                     string                                            `json:"review_decision"`
	ReceiptSchema                      string                                            `json:"receipt_schema"`
	OpaqueReceiptID                    string                                            `json:"opaque_receipt_id"`
	ReceiptRequired                    bool                                              `json:"receipt_required"`
	ReceiptPresent                     bool                                              `json:"receipt_present"`
	ReceiptAccepted                    bool                                              `json:"receipt_accepted"`
	WriterAuthorizationRequired        bool                                              `json:"writer_authorization_required"`
	WriterAuthorizationModeled         bool                                              `json:"writer_authorization_modeled"`
	OperatorActionRequired             bool                                              `json:"operator_action_required"`
	AcceptancePropagationConsumed      bool                                              `json:"acceptance_propagation_consumed"`
	OwnerManagedOpaqueBoundaryReady    bool                                              `json:"owner_managed_opaque_boundary_ready"`
	CallerStateRootRequired            bool                                              `json:"caller_state_root_required"`
	AuthorizationAccepted              bool                                              `json:"authorization_accepted"`
	WriterReviewReady                  bool                                              `json:"writer_review_ready"`
	ProductionReadiness                bool                                              `json:"production_readiness"`
	ProductionOwnershipReady           bool                                              `json:"production_ownership_ready"`
	ReviewItemCount                    int                                               `json:"review_item_count"`
	RequiredReviewItemCount            int                                               `json:"required_review_item_count"`
	ReadyReviewItemCount               int                                               `json:"ready_review_item_count"`
	MissingReviewItemCount             int                                               `json:"missing_review_item_count"`
	WriteEnabledReviewItemCount        int                                               `json:"write_enabled_review_item_count"`
	AcceptanceEnabledReviewItemCount   int                                               `json:"acceptance_enabled_review_item_count"`
	SideEffectReviewItemCount          int                                               `json:"side_effect_review_item_count"`
	ReviewItems                        []ProductionReceiptWriterAuthorizationReviewItem  `json:"review_items"`
	ReviewItemIDs                      []string                                          `json:"review_item_ids"`
	RequiredBeforeReceiptWriterEnabled []string                                          `json:"required_before_receipt_writer_enabled"`
	Checks                             []ProductionReceiptWriterAuthorizationReviewCheck `json:"checks"`
	CheckIDs                           []string                                          `json:"check_ids"`
	Counts                             ProductionReceiptWriterAuthorizationReviewCounts  `json:"counts"`
	RuntimeOwned                       bool                                              `json:"runtime_owned"`
	GoRuntimeBacked                    bool                                              `json:"go_runtime_backed"`
	KDEPolicyOwner                     bool                                              `json:"kde_policy_owner"`
	SystemServiceStarted               bool                                              `json:"system_service_started"`
	SessionBusClaimed                  bool                                              `json:"session_bus_claimed"`
	ProductionBusClaimed               bool                                              `json:"production_bus_claimed"`
	ProductionOwnerEnabled             bool                                              `json:"production_owner_enabled"`
	ProductionActivationReady          bool                                              `json:"production_activation_ready"`
	WriteMethodsEnabled                bool                                              `json:"write_methods_enabled"`
	RuntimeWritesEnabled               bool                                              `json:"runtime_writes_enabled"`
	ReceiptWriterEnabled               bool                                              `json:"receipt_writer_enabled"`
	ReceiptPersistenceEnabled          bool                                              `json:"receipt_persistence_enabled"`
	ReceiptLookupWritesEnabled         bool                                              `json:"receipt_lookup_writes_enabled"`
	ReceiptReplayEnabled               bool                                              `json:"receipt_replay_enabled"`
	DesktopFilesWritten                bool                                              `json:"desktop_files_written"`
	MIMEAppsWritten                    bool                                              `json:"mimeapps_written"`
	ShellConfigurationWritten          bool                                              `json:"shell_configuration_written"`
	SettingsPersisted                  bool                                              `json:"settings_persisted"`
	NotificationSent                   bool                                              `json:"notification_sent"`
	NotificationDeliveryEnabled        bool                                              `json:"notification_delivery_enabled"`
	PortalRequestCreated               bool                                              `json:"portal_request_created"`
	RequestObjectsCreated              bool                                              `json:"request_objects_created"`
	AdapterInvocationEnabled           bool                                              `json:"adapter_invocation_enabled"`
	BackendLaunchEnabled               bool                                              `json:"backend_launch_enabled"`
	BackendProcessStarted              bool                                              `json:"backend_process_started"`
	SupportBundleExported              bool                                              `json:"support_bundle_exported"`
	SupportCaseCreated                 bool                                              `json:"support_case_created"`
	SnapshotRestoreExecuted            bool                                              `json:"snapshot_restore_executed"`
	StateCleanupExecuted               bool                                              `json:"state_cleanup_executed"`
	FileContentRead                    bool                                              `json:"file_content_read"`
	FilePathsExposed                   bool                                              `json:"file_paths_exposed"`
	NetworkRequired                    bool                                              `json:"network_required"`
	HostRootModified                   bool                                              `json:"host_root_modified"`
	PrivilegedContainerRequired        bool                                              `json:"privileged_container_required"`
	StateRootPathExposed               bool                                              `json:"state_root_path_exposed"`
	RawCommandExposed                  bool                                              `json:"raw_command_exposed"`
	RawExecutableExposed               bool                                              `json:"raw_executable_exposed"`
	BackendDetailsExposed              bool                                              `json:"backend_details_exposed"`
	BlockedActions                     []string                                          `json:"blocked_actions"`
	NextRequirements                   []string                                          `json:"next_requirements"`
	DesktopSafeSummary                 string                                            `json:"desktop_safe_summary"`
}

type ProductionReceiptWriterAuthorizationReviewItem struct {
	ID                         string `json:"id"`
	ReviewArea                 string `json:"review_area"`
	RequiredEvidence           string `json:"required_evidence"`
	EvidencePresent            bool   `json:"evidence_present"`
	WriterAuthorizationModeled bool   `json:"writer_authorization_modeled"`
	ReceiptWriterEnabled       bool   `json:"receipt_writer_enabled"`
	ReceiptPersistenceEnabled  bool   `json:"receipt_persistence_enabled"`
	ReceiptLookupWritesEnabled bool   `json:"receipt_lookup_writes_enabled"`
	ReceiptReplayEnabled       bool   `json:"receipt_replay_enabled"`
	ReceiptAccepted            bool   `json:"receipt_accepted"`
	AuthorizationAccepted      bool   `json:"authorization_accepted"`
	ProductionReadiness        bool   `json:"production_readiness"`
	ProductionOwnershipReady   bool   `json:"production_ownership_ready"`
	RuntimeOwned               bool   `json:"runtime_owned"`
	GoRuntimeBacked            bool   `json:"go_runtime_backed"`
	KDEPolicyOwner             bool   `json:"kde_policy_owner"`
	ReviewOnly                 bool   `json:"review_only"`
	SideEffectsDisabled        bool   `json:"side_effects_disabled"`
	HostRootModified           bool   `json:"host_root_modified"`
	InternalDetailsExposed     bool   `json:"internal_details_exposed"`
	ReviewStatus               string `json:"review_status"`
	NextRequirement            string `json:"next_requirement"`
}

type ProductionReceiptWriterAuthorizationReviewCheck struct {
	ID      string `json:"id"`
	Status  string `json:"status"`
	Summary string `json:"summary"`
}

type ProductionReceiptWriterAuthorizationReviewCounts struct {
	Total   int `json:"total"`
	Passed  int `json:"passed"`
	Pending int `json:"pending"`
	Blocked int `json:"blocked"`
}

func NewProductionReceiptWriterAuthorizationReviewPreview(root string) (ProductionReceiptWriterAuthorizationReviewPreview, error) {
	version, err := readProductionDBusGateReviewVersion(root)
	if err != nil {
		return ProductionReceiptWriterAuthorizationReviewPreview{}, err
	}
	sources := productionReceiptWriterAuthorizationReviewSources(root)
	items := productionReceiptWriterAuthorizationReviewItems(sources)
	acceptancePropagationReady := productionReceiptWriterAcceptancePropagationReady(sources.AcceptancePropagation)
	preview := ProductionReceiptWriterAuthorizationReviewPreview{
		Version:                          version,
		SchemaVersion:                    "xnix.runtime.production_receipt_writer_authorization_review.v1",
		RequestType:                      "production-receipt-writer-authorization-review-preview",
		ReviewType:                       "receipt-writer-operator-authorization-boundary-review",
		Source:                           "production-receipt-acceptance-propagation-preflight-preview+production-human-authorization-receipt-consolidation-preview+production-authorization-consumption-audit-preview",
		ReviewDecision:                   "production-receipt-writer-authorization-review-blocked",
		ReceiptSchema:                    "xnix.runtime.production_dbus_human_authorization_receipt.v1",
		OpaqueReceiptID:                  ProductionDBusHumanAuthorizationReceiptID,
		ReceiptRequired:                  true,
		ReceiptPresent:                   false,
		ReceiptAccepted:                  false,
		WriterAuthorizationRequired:      true,
		WriterAuthorizationModeled:       true,
		OperatorActionRequired:           true,
		AcceptancePropagationConsumed:    acceptancePropagationReady,
		OwnerManagedOpaqueBoundaryReady:  acceptancePropagationReady,
		CallerStateRootRequired:          false,
		AuthorizationAccepted:            false,
		WriterReviewReady:                acceptancePropagationReady && productionReceiptWriterAllReviewItemsReady(items),
		ProductionReadiness:              false,
		ProductionOwnershipReady:         false,
		ReviewItemCount:                  len(items),
		RequiredReviewItemCount:          5,
		ReadyReviewItemCount:             productionReceiptWriterReadyReviewItemCount(items),
		MissingReviewItemCount:           productionReceiptWriterMissingReviewItemCount(items),
		WriteEnabledReviewItemCount:      0,
		AcceptanceEnabledReviewItemCount: 0,
		SideEffectReviewItemCount:        0,
		ReviewItems:                      items,
		ReviewItemIDs:                    productionReceiptWriterReviewItemIDs(items),
		RequiredBeforeReceiptWriterEnabled: []string{
			"production-receipt-acceptance-propagation-preflight-preview",
			"production-human-authorization-receipt-consolidation-preview",
			"production-authorization-consumption-audit-preview",
			"separate operator authorization action outside this review",
			"separate persistence and replay threat review outside this review",
			"separate production service ownership proof outside this review",
		},
		RuntimeOwned:                true,
		GoRuntimeBacked:             true,
		KDEPolicyOwner:              false,
		SystemServiceStarted:        false,
		SessionBusClaimed:           false,
		ProductionBusClaimed:        false,
		ProductionOwnerEnabled:      false,
		ProductionActivationReady:   false,
		WriteMethodsEnabled:         false,
		RuntimeWritesEnabled:        false,
		ReceiptWriterEnabled:        false,
		ReceiptPersistenceEnabled:   false,
		ReceiptLookupWritesEnabled:  false,
		ReceiptReplayEnabled:        false,
		DesktopFilesWritten:         false,
		MIMEAppsWritten:             false,
		ShellConfigurationWritten:   false,
		SettingsPersisted:           false,
		NotificationSent:            false,
		NotificationDeliveryEnabled: false,
		PortalRequestCreated:        false,
		RequestObjectsCreated:       false,
		AdapterInvocationEnabled:    false,
		BackendLaunchEnabled:        false,
		BackendProcessStarted:       false,
		SupportBundleExported:       false,
		SupportCaseCreated:          false,
		SnapshotRestoreExecuted:     false,
		StateCleanupExecuted:        false,
		FileContentRead:             false,
		FilePathsExposed:            false,
		NetworkRequired:             false,
		HostRootModified:            false,
		PrivilegedContainerRequired: false,
		StateRootPathExposed:        false,
		RawCommandExposed:           false,
		RawExecutableExposed:        false,
		BackendDetailsExposed:       false,
		BlockedActions: []string{
			"treat writer authorization review as permission to write receipts",
			"write, persist, accept, replay, or look up authorization receipts from this review",
			"claim production D-Bus ownership, start services, or enable Runtime writes from this review",
			"enable desktop writes, notifications, Portal requests, request objects, adapter invocation, or compatibility engine launch",
			"export support bundles, create support cases, restore snapshots, clean state, read file contents, or expose paths",
			"require network, require privileged containers, expose raw commands, expose internal engine details, or mutate host root",
		},
		NextRequirements: []string{
			"Keep receipt writer enablement outside preview commands until a separate operator action exists.",
			"Require persistence and replay threat review before any accepted receipt can be stored.",
			"Keep service start, bus claim, write dispatch, desktop side effects, support side effects, restore, cleanup, launch, and host mutation disabled.",
			"Add an explicit writer implementation plan only after this review remains read-only and fail-closed.",
		},
		DesktopSafeSummary: "The production receipt writer authorization review identifies the operator and persistence boundary required before a future authorization receipt writer can exist, but it does not write receipts, accept authorization, claim a bus, start services, enable writes, emit desktop side effects, launch engines, or mutate host state.",
	}
	checks := productionReceiptWriterAuthorizationReviewChecks(preview)
	preview.Checks = checks
	preview.CheckIDs = productionReceiptWriterCheckIDs(checks)
	preview.Counts = countProductionReceiptWriterChecks(checks)
	if preview.Counts.Blocked == 0 {
		preview.ReviewDecision = "production-receipt-writer-authorization-review-ready-writes-disabled"
	}
	if err := validateNoBackendTerms(preview, "production receipt writer authorization review preview"); err != nil {
		return ProductionReceiptWriterAuthorizationReviewPreview{}, err
	}
	return preview, nil
}

type productionReceiptWriterAuthorizationReviewSourceSet struct {
	AcceptancePropagation string
	Consolidation         string
	ConsumptionAudit      string
}

func productionReceiptWriterAuthorizationReviewSources(root string) productionReceiptWriterAuthorizationReviewSourceSet {
	return productionReceiptWriterAuthorizationReviewSourceSet{
		AcceptancePropagation: productionAuthorizationReadSources(root, []string{"internal/runtime/owner/production_receipt_acceptance_propagation_preflight.go"}),
		Consolidation:         productionAuthorizationReadSources(root, []string{"internal/runtime/owner/production_human_authorization_receipt_consolidation.go"}),
		ConsumptionAudit:      productionAuthorizationReadSources(root, []string{"internal/runtime/owner/production_authorization_consumption_audit.go"}),
	}
}

func productionReceiptWriterAuthorizationReviewItems(sources productionReceiptWriterAuthorizationReviewSourceSet) []ProductionReceiptWriterAuthorizationReviewItem {
	return []ProductionReceiptWriterAuthorizationReviewItem{
		productionReceiptWriterReviewItem("operator-action-boundary", "operator-authorization", "explicit operator action and writer authorization requirement", sources.AcceptancePropagation, []string{"separate operator action authorizing receipt acceptance", "separate receipt writer and persistence review", "AcceptanceSimulationOnly"}, "define operator action without enabling receipt writes"),
		productionReceiptWriterReviewItem("opaque-boundary-consolidated", "receipt-boundary", "owner-managed opaque receipt id and schema", sources.Consolidation, []string{"owner-managed-opaque-human-authorization-receipt-boundary", "OwnerManagedOpaqueReceiptLookupReady", "ReceiptWriterEnabled"}, "keep opaque receipt boundary read-only"),
		productionReceiptWriterReviewItem("production-gates-consume-boundary", "gate-consumption", "all production gates consume the consolidated boundary", sources.ConsumptionAudit, []string{"consumers-use-consolidated-boundary", "authorization-not-accepted", "desktop-and-support-side-effects-disabled"}, "keep all production gates consuming the same boundary"),
		productionReceiptWriterReviewItem("persistence-disabled", "receipt-persistence", "receipt persistence and replay stay disabled", sources.Consolidation+sources.AcceptancePropagation, []string{"ReceiptPersistenceEnabled", "ReceiptLookupWritesEnabled", "ReceiptWriterEnabled"}, "add a separate persistence threat review before storing any accepted receipt"),
		productionReceiptWriterReviewItem("host-boundary-closed", "host-boundary", "network, privilege, paths, and host mutation stay closed", sources.ConsumptionAudit+sources.AcceptancePropagation, []string{"host-boundary-closed", "HostRootModified", "PrivilegedContainerRequired", "StateRootPathExposed"}, "keep writer review independent of host mutation and caller paths"),
	}
}

func productionReceiptWriterReviewItem(id string, area string, evidence string, source string, tokens []string, nextRequirement string) ProductionReceiptWriterAuthorizationReviewItem {
	ready := productionAuthorizationHasAll(source, tokens)
	status := "missing-review-evidence"
	if ready {
		status = "writer-authorization-reviewed-writes-disabled"
	}
	return ProductionReceiptWriterAuthorizationReviewItem{
		ID:                         id,
		ReviewArea:                 area,
		RequiredEvidence:           evidence,
		EvidencePresent:            ready,
		WriterAuthorizationModeled: ready,
		ReceiptWriterEnabled:       false,
		ReceiptPersistenceEnabled:  false,
		ReceiptLookupWritesEnabled: false,
		ReceiptReplayEnabled:       false,
		ReceiptAccepted:            false,
		AuthorizationAccepted:      false,
		ProductionReadiness:        false,
		ProductionOwnershipReady:   false,
		RuntimeOwned:               true,
		GoRuntimeBacked:            true,
		KDEPolicyOwner:             false,
		ReviewOnly:                 true,
		SideEffectsDisabled:        true,
		HostRootModified:           false,
		InternalDetailsExposed:     false,
		ReviewStatus:               status,
		NextRequirement:            nextRequirement,
	}
}

func productionReceiptWriterAuthorizationReviewChecks(preview ProductionReceiptWriterAuthorizationReviewPreview) []ProductionReceiptWriterAuthorizationReviewCheck {
	return []ProductionReceiptWriterAuthorizationReviewCheck{
		productionReceiptWriterCheck("acceptance-propagation-consumed", productionAuthorizationPassBlocked(preview.AcceptancePropagationConsumed && preview.OwnerManagedOpaqueBoundaryReady && preview.ReceiptSchema == "xnix.runtime.production_dbus_human_authorization_receipt.v1" && preview.OpaqueReceiptID == ProductionDBusHumanAuthorizationReceiptID), "The review consumes the future acceptance propagation preflight and opaque receipt boundary."),
		productionReceiptWriterCheck("writer-authorization-modeled-only", productionAuthorizationPassBlocked(preview.WriterAuthorizationRequired && preview.WriterAuthorizationModeled && preview.OperatorActionRequired && !preview.ReceiptPresent && !preview.ReceiptAccepted && !preview.AuthorizationAccepted), "The review models the writer authorization boundary without accepting a receipt."),
		productionReceiptWriterCheck("five-review-items-present", productionAuthorizationPassBlocked(preview.ReviewItemCount == 5 && preview.RequiredReviewItemCount == 5 && preview.MissingReviewItemCount == 0), "The review tracks operator action, opaque boundary, gate consumption, persistence, and host boundary items."),
		productionReceiptWriterCheck("review-items-ready-writes-disabled", productionAuthorizationPassBlocked(preview.ReadyReviewItemCount == 5 && productionReceiptWriterAllReviewItemsReady(preview.ReviewItems)), "Every review item is ready while receipt writes and acceptance remain disabled."),
		productionReceiptWriterCheck("receipt-writes-disabled", productionAuthorizationPassBlocked(!preview.ReceiptWriterEnabled && !preview.ReceiptPersistenceEnabled && !preview.ReceiptLookupWritesEnabled && !preview.ReceiptReplayEnabled && preview.WriteEnabledReviewItemCount == 0 && preview.AcceptanceEnabledReviewItemCount == 0), "Receipt writer, persistence, lookup writes, replay, and acceptance remain disabled."),
		productionReceiptWriterCheck("production-ownership-disabled", productionAuthorizationPassBlocked(!preview.SystemServiceStarted && !preview.SessionBusClaimed && !preview.ProductionBusClaimed && !preview.ProductionOwnerEnabled && !preview.ProductionActivationReady && !preview.ProductionReadiness && !preview.ProductionOwnershipReady), "Service start, session bus claim, production bus claim, and production ownership remain disabled."),
		productionReceiptWriterCheck("runtime-side-effects-disabled", productionAuthorizationPassBlocked(!preview.WriteMethodsEnabled && !preview.RuntimeWritesEnabled && !preview.DesktopFilesWritten && !preview.MIMEAppsWritten && !preview.ShellConfigurationWritten && !preview.SettingsPersisted && !preview.NotificationSent && !preview.NotificationDeliveryEnabled && !preview.PortalRequestCreated && !preview.RequestObjectsCreated && !preview.AdapterInvocationEnabled && !preview.BackendLaunchEnabled && !preview.BackendProcessStarted && !preview.SupportBundleExported && !preview.SupportCaseCreated && !preview.SnapshotRestoreExecuted && !preview.StateCleanupExecuted && preview.SideEffectReviewItemCount == 0), "Runtime writes, desktop, Portal, support, restore, cleanup, adapter invocation, and launch side effects remain disabled."),
		productionReceiptWriterCheck("host-boundary-closed", productionAuthorizationPassBlocked(!preview.FileContentRead && !preview.FilePathsExposed && !preview.NetworkRequired && !preview.HostRootModified && !preview.PrivilegedContainerRequired && !preview.StateRootPathExposed && !preview.RawCommandExposed && !preview.RawExecutableExposed && !preview.BackendDetailsExposed && productionReceiptWriterItemsKeepHostClosed(preview.ReviewItems)), "Unsafe data exposure, network, privilege, internal detail exposure, and host mutation gates remain closed."),
	}
}

func productionReceiptWriterCheck(id string, status string, summary string) ProductionReceiptWriterAuthorizationReviewCheck {
	return ProductionReceiptWriterAuthorizationReviewCheck{ID: id, Status: status, Summary: summary}
}

func productionReceiptWriterCheckIDs(checks []ProductionReceiptWriterAuthorizationReviewCheck) []string {
	ids := make([]string, 0, len(checks))
	for _, check := range checks {
		ids = append(ids, check.ID)
	}
	return ids
}

func countProductionReceiptWriterChecks(checks []ProductionReceiptWriterAuthorizationReviewCheck) ProductionReceiptWriterAuthorizationReviewCounts {
	counts := ProductionReceiptWriterAuthorizationReviewCounts{Total: len(checks)}
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

func productionReceiptWriterAcceptancePropagationReady(source string) bool {
	return productionAuthorizationHasAll(source, []string{
		"production-receipt-acceptance-propagation-preflight-preview",
		"future-authorization-receipt-acceptance-propagation-preflight",
		"production-receipt-acceptance-propagation-ready-acceptance-disabled",
		"future-acceptance-modeled-only",
		"acceptance-and-production-disabled",
		"write-and-launch-disabled",
		"desktop-support-and-host-boundary-closed",
	})
}

func productionReceiptWriterReadyReviewItemCount(items []ProductionReceiptWriterAuthorizationReviewItem) int {
	count := 0
	for _, item := range items {
		if item.EvidencePresent {
			count++
		}
	}
	return count
}

func productionReceiptWriterMissingReviewItemCount(items []ProductionReceiptWriterAuthorizationReviewItem) int {
	count := 0
	for _, item := range items {
		if !item.EvidencePresent {
			count++
		}
	}
	return count
}

func productionReceiptWriterReviewItemIDs(items []ProductionReceiptWriterAuthorizationReviewItem) []string {
	ids := make([]string, 0, len(items))
	for _, item := range items {
		ids = append(ids, item.ID)
	}
	return ids
}

func productionReceiptWriterAllReviewItemsReady(items []ProductionReceiptWriterAuthorizationReviewItem) bool {
	if len(items) != 5 {
		return false
	}
	for _, item := range items {
		if !item.EvidencePresent || !item.WriterAuthorizationModeled || item.ReviewStatus != "writer-authorization-reviewed-writes-disabled" {
			return false
		}
		if item.ReceiptWriterEnabled || item.ReceiptPersistenceEnabled || item.ReceiptLookupWritesEnabled || item.ReceiptReplayEnabled || item.ReceiptAccepted || item.AuthorizationAccepted || item.ProductionReadiness || item.ProductionOwnershipReady || item.KDEPolicyOwner || !item.ReviewOnly || !item.SideEffectsDisabled {
			return false
		}
	}
	return true
}

func productionReceiptWriterItemsKeepHostClosed(items []ProductionReceiptWriterAuthorizationReviewItem) bool {
	for _, item := range items {
		if item.HostRootModified || item.InternalDetailsExposed {
			return false
		}
	}
	return true
}
