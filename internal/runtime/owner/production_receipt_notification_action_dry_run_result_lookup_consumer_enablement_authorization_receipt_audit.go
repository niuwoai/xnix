package owner

const ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementAuthorizationReceiptID = "dry-run-result-lookup-consumer-enablement-authorization-receipt-id"

type ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementAuthorizationReceiptAuditPreview struct {
	Version                              string                                                                                                  `json:"version"`
	SchemaVersion                        string                                                                                                  `json:"schema_version"`
	RequestType                          string                                                                                                  `json:"request_type"`
	AuditType                            string                                                                                                  `json:"audit_type"`
	Source                               string                                                                                                  `json:"source"`
	AuditDecision                        string                                                                                                  `json:"audit_decision"`
	ReceiptSchema                        string                                                                                                  `json:"receipt_schema"`
	OpaqueReceiptID                      string                                                                                                  `json:"opaque_receipt_id"`
	AuthorizationReceiptAuditRequired    bool                                                                                                    `json:"authorization_receipt_audit_required"`
	AuthorizationReceiptAuditModeled     bool                                                                                                    `json:"authorization_receipt_audit_modeled"`
	ConsumerEnablementGateAuditConsumed  bool                                                                                                    `json:"consumer_enablement_gate_audit_consumed"`
	AuthorizationReceiptGuidanceConsumed bool                                                                                                    `json:"authorization_receipt_guidance_consumed"`
	AuthorizationReceiptBoundaryReady    bool                                                                                                    `json:"authorization_receipt_boundary_ready"`
	ConsumerEnablementGateReady          bool                                                                                                    `json:"consumer_enablement_gate_ready"`
	OpaqueAuthorizationReceiptModeled    bool                                                                                                    `json:"opaque_authorization_receipt_modeled"`
	ConsumerAuthorizationReady           bool                                                                                                    `json:"consumer_authorization_ready"`
	ReceiptPresent                       bool                                                                                                    `json:"receipt_present"`
	ReceiptAccepted                      bool                                                                                                    `json:"receipt_accepted"`
	AuthorizationAccepted                bool                                                                                                    `json:"authorization_accepted"`
	ConsumerConsumptionAuthorized        bool                                                                                                    `json:"consumer_consumption_authorized"`
	ConsumerEnablementAuthorized         bool                                                                                                    `json:"consumer_enablement_authorized"`
	KDEConsumerEnabled                   bool                                                                                                    `json:"kde_consumer_enabled"`
	RuntimeConsumerEnabled               bool                                                                                                    `json:"runtime_consumer_enabled"`
	LookupRouteAuthorized                bool                                                                                                    `json:"lookup_route_authorized"`
	LookupRouteEnabled                   bool                                                                                                    `json:"lookup_route_enabled"`
	LookupRoutePersisted                 bool                                                                                                    `json:"lookup_route_persisted"`
	OpaqueLookupEnabled                  bool                                                                                                    `json:"opaque_lookup_enabled"`
	OpaqueLookupPersisted                bool                                                                                                    `json:"opaque_lookup_persisted"`
	RedactedSummaryPersisted             bool                                                                                                    `json:"redacted_summary_persisted"`
	RawResultExposed                     bool                                                                                                    `json:"raw_result_exposed"`
	RuntimeDiagnosticsPersisted          bool                                                                                                    `json:"runtime_diagnostics_persisted"`
	DryRunResultPersisted                bool                                                                                                    `json:"dry_run_result_persisted"`
	DispatchDryRunExecuted               bool                                                                                                    `json:"dispatch_dry_run_executed"`
	RequestObjectCreationEnabled         bool                                                                                                    `json:"request_object_creation_enabled"`
	RequestObjectDispatchEnabled         bool                                                                                                    `json:"request_object_dispatch_enabled"`
	PortalRequestCreated                 bool                                                                                                    `json:"portal_request_created"`
	NotificationActionEnabled            bool                                                                                                    `json:"notification_action_enabled"`
	CompatibilityCenterOpened            bool                                                                                                    `json:"compatibility_center_opened"`
	SupportBundleExported                bool                                                                                                    `json:"support_bundle_exported"`
	SupportCaseCreated                   bool                                                                                                    `json:"support_case_created"`
	ProductionReadiness                  bool                                                                                                    `json:"production_readiness"`
	ProductionOwnershipReady             bool                                                                                                    `json:"production_ownership_ready"`
	ReceiptItemCount                     int                                                                                                     `json:"receipt_item_count"`
	RequiredReceiptItemCount             int                                                                                                     `json:"required_receipt_item_count"`
	ReadyReceiptItemCount                int                                                                                                     `json:"ready_receipt_item_count"`
	MissingReceiptItemCount              int                                                                                                     `json:"missing_receipt_item_count"`
	AcceptedReceiptItemCount             int                                                                                                     `json:"accepted_receipt_item_count"`
	AuthorizedConsumerItemCount          int                                                                                                     `json:"authorized_consumer_item_count"`
	EnabledConsumerItemCount             int                                                                                                     `json:"enabled_consumer_item_count"`
	PersistedReceiptItemCount            int                                                                                                     `json:"persisted_receipt_item_count"`
	RawExposedReceiptItemCount           int                                                                                                     `json:"raw_exposed_receipt_item_count"`
	SideEffectReceiptItemCount           int                                                                                                     `json:"side_effect_receipt_item_count"`
	ReceiptItems                         []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementAuthorizationReceiptAuditItem  `json:"receipt_items"`
	ReceiptItemIDs                       []string                                                                                                `json:"receipt_item_ids"`
	RequiredBeforeReceiptAcceptance      []string                                                                                                `json:"required_before_receipt_acceptance"`
	Checks                               []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementAuthorizationReceiptAuditCheck `json:"checks"`
	CheckIDs                             []string                                                                                                `json:"check_ids"`
	Counts                               ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementAuthorizationReceiptAuditCounts  `json:"counts"`
	RuntimeOwned                         bool                                                                                                    `json:"runtime_owned"`
	GoRuntimeBacked                      bool                                                                                                    `json:"go_runtime_backed"`
	KDEPolicyOwner                       bool                                                                                                    `json:"kde_policy_owner"`
	OfficialDesktopOnly                  bool                                                                                                    `json:"official_desktop_only"`
	PlasmaForkRequired                   bool                                                                                                    `json:"plasma_fork_required"`
	PlasmaSourceModified                 bool                                                                                                    `json:"plasma_source_modified"`
	SystemServiceStarted                 bool                                                                                                    `json:"system_service_started"`
	SessionBusClaimed                    bool                                                                                                    `json:"session_bus_claimed"`
	ProductionBusClaimed                 bool                                                                                                    `json:"production_bus_claimed"`
	ProductionOwnerEnabled               bool                                                                                                    `json:"production_owner_enabled"`
	ProductionActivationReady            bool                                                                                                    `json:"production_activation_ready"`
	WriteMethodsEnabled                  bool                                                                                                    `json:"write_methods_enabled"`
	RuntimeWritesEnabled                 bool                                                                                                    `json:"runtime_writes_enabled"`
	RequestObjectsCreated                bool                                                                                                    `json:"request_objects_created"`
	RequestObjectsDispatched             bool                                                                                                    `json:"request_objects_dispatched"`
	NotificationSent                     bool                                                                                                    `json:"notification_sent"`
	NotificationDeliveryEnabled          bool                                                                                                    `json:"notification_delivery_enabled"`
	CompatibilityCenterPersisted         bool                                                                                                    `json:"compatibility_center_persisted"`
	ReceiptWriterEnabled                 bool                                                                                                    `json:"receipt_writer_enabled"`
	ReceiptPersistenceEnabled            bool                                                                                                    `json:"receipt_persistence_enabled"`
	DesktopFilesWritten                  bool                                                                                                    `json:"desktop_files_written"`
	MIMEAppsWritten                      bool                                                                                                    `json:"mimeapps_written"`
	ShellConfigurationWritten            bool                                                                                                    `json:"shell_configuration_written"`
	SettingsPersisted                    bool                                                                                                    `json:"settings_persisted"`
	AdapterInvocationEnabled             bool                                                                                                    `json:"adapter_invocation_enabled"`
	BackendLaunchEnabled                 bool                                                                                                    `json:"backend_launch_enabled"`
	BackendProcessStarted                bool                                                                                                    `json:"backend_process_started"`
	SnapshotRestoreExecuted              bool                                                                                                    `json:"snapshot_restore_executed"`
	StateCleanupExecuted                 bool                                                                                                    `json:"state_cleanup_executed"`
	NetworkRequired                      bool                                                                                                    `json:"network_required"`
	HostRootModified                     bool                                                                                                    `json:"host_root_modified"`
	PrivilegedContainerRequired          bool                                                                                                    `json:"privileged_container_required"`
	CallerStateRootRequired              bool                                                                                                    `json:"caller_state_root_required"`
	StateRootPathExposed                 bool                                                                                                    `json:"state_root_path_exposed"`
	FilePathsExposed                     bool                                                                                                    `json:"file_paths_exposed"`
	FileContentRead                      bool                                                                                                    `json:"file_content_read"`
	RawCommandExposed                    bool                                                                                                    `json:"raw_command_exposed"`
	RawExecutableExposed                 bool                                                                                                    `json:"raw_executable_exposed"`
	BackendDetailsExposed                bool                                                                                                    `json:"backend_details_exposed"`
	BlockedActions                       []string                                                                                                `json:"blocked_actions"`
	NextRequirements                     []string                                                                                                `json:"next_requirements"`
	DesktopSafeSummary                   string                                                                                                  `json:"desktop_safe_summary"`
}

type ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementAuthorizationReceiptAuditItem struct {
	ID                                  string `json:"id"`
	ActionKind                          string `json:"action_kind"`
	SurfaceKind                         string `json:"surface_kind"`
	ReceiptKind                         string `json:"receipt_kind"`
	OpaqueReceiptID                     string `json:"opaque_receipt_id"`
	RequiredEvidence                    string `json:"required_evidence"`
	EvidencePresent                     bool   `json:"evidence_present"`
	ConsumerEnablementGateConsumed      bool   `json:"consumer_enablement_gate_consumed"`
	AuthorizationReceiptBoundaryModeled bool   `json:"authorization_receipt_boundary_modeled"`
	OpaqueAuthorizationReceiptModeled   bool   `json:"opaque_authorization_receipt_modeled"`
	ReceiptPresent                      bool   `json:"receipt_present"`
	ReceiptAccepted                     bool   `json:"receipt_accepted"`
	AuthorizationAccepted               bool   `json:"authorization_accepted"`
	ConsumerConsumptionAuthorized       bool   `json:"consumer_consumption_authorized"`
	ConsumerEnablementAuthorized        bool   `json:"consumer_enablement_authorized"`
	KDEConsumerEnabled                  bool   `json:"kde_consumer_enabled"`
	RuntimeConsumerEnabled              bool   `json:"runtime_consumer_enabled"`
	LookupRouteAuthorized               bool   `json:"lookup_route_authorized"`
	LookupRouteEnabled                  bool   `json:"lookup_route_enabled"`
	LookupRoutePersisted                bool   `json:"lookup_route_persisted"`
	OpaqueLookupEnabled                 bool   `json:"opaque_lookup_enabled"`
	OpaqueLookupPersisted               bool   `json:"opaque_lookup_persisted"`
	RedactedSummaryPersisted            bool   `json:"redacted_summary_persisted"`
	RawResultExposed                    bool   `json:"raw_result_exposed"`
	RuntimeDiagnosticsPersisted         bool   `json:"runtime_diagnostics_persisted"`
	DryRunResultPersisted               bool   `json:"dry_run_result_persisted"`
	DispatchDryRunExecuted              bool   `json:"dispatch_dry_run_executed"`
	UserVisible                         bool   `json:"user_visible"`
	ReviewOnly                          bool   `json:"review_only"`
	RuntimeOwned                        bool   `json:"runtime_owned"`
	GoRuntimeBacked                     bool   `json:"go_runtime_backed"`
	KDEPolicyOwner                      bool   `json:"kde_policy_owner"`
	CallerStateRootRequired             bool   `json:"caller_state_root_required"`
	StateRootPathExposed                bool   `json:"state_root_path_exposed"`
	FilePathsExposed                    bool   `json:"file_paths_exposed"`
	FileContentRead                     bool   `json:"file_content_read"`
	RequestObjectCreated                bool   `json:"request_object_created"`
	RequestObjectDispatched             bool   `json:"request_object_dispatched"`
	PortalRequestCreated                bool   `json:"portal_request_created"`
	NotificationActionEnabled           bool   `json:"notification_action_enabled"`
	CompatibilityCenterOpened           bool   `json:"compatibility_center_opened"`
	SupportBundleExported               bool   `json:"support_bundle_exported"`
	SupportCaseCreated                  bool   `json:"support_case_created"`
	ProductionReadiness                 bool   `json:"production_readiness"`
	ProductionOwnershipReady            bool   `json:"production_ownership_ready"`
	SideEffectsDisabled                 bool   `json:"side_effects_disabled"`
	HostRootModified                    bool   `json:"host_root_modified"`
	InternalDetailsExposed              bool   `json:"internal_details_exposed"`
	ReceiptStatus                       string `json:"receipt_status"`
	NextRequirement                     string `json:"next_requirement"`
}

type ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementAuthorizationReceiptAuditCheck struct {
	ID      string `json:"id"`
	Status  string `json:"status"`
	Summary string `json:"summary"`
}

type ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementAuthorizationReceiptAuditCounts struct {
	Total   int `json:"total"`
	Passed  int `json:"passed"`
	Pending int `json:"pending"`
	Blocked int `json:"blocked"`
}

func NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementAuthorizationReceiptAuditPreview(root string) (ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementAuthorizationReceiptAuditPreview, error) {
	version, err := readProductionDBusGateReviewVersion(root)
	if err != nil {
		return ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementAuthorizationReceiptAuditPreview{}, err
	}
	sources := productionReceiptNotificationActionDryRunResultLookupConsumerEnablementAuthorizationReceiptAuditSources(root)
	items := productionReceiptNotificationActionDryRunResultLookupConsumerEnablementAuthorizationReceiptItems(sources)
	gateReady := productionReceiptNotificationActionDryRunResultLookupConsumerEnablementAuthorizationReceiptGateReady(sources.ConsumerEnablementGateAudit)
	guidanceReady := productionReceiptNotificationActionDryRunResultLookupConsumerEnablementAuthorizationReceiptGuidanceReady(sources.DispatchSheet)
	boundaryReady := gateReady && guidanceReady && productionReceiptNotificationActionDryRunResultLookupConsumerEnablementAuthorizationReceiptItemsReady(items)
	preview := ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementAuthorizationReceiptAuditPreview{
		Version:                              version,
		SchemaVersion:                        "xnix.runtime.production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_authorization_receipt_audit.v1",
		RequestType:                          "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-authorization-receipt-audit-preview",
		AuditType:                            "receipt-notification-action-dry-run-result-lookup-consumer-enablement-authorization-receipt-audit",
		Source:                               "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-gate-audit-preview+claude-code-current-dispatch-picks",
		AuditDecision:                        "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-authorization-receipt-audit-blocked",
		ReceiptSchema:                        "xnix.runtime.production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_authorization_receipt.v1",
		OpaqueReceiptID:                      ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementAuthorizationReceiptID,
		AuthorizationReceiptAuditRequired:    true,
		AuthorizationReceiptAuditModeled:     true,
		ConsumerEnablementGateAuditConsumed:  gateReady,
		AuthorizationReceiptGuidanceConsumed: guidanceReady,
		AuthorizationReceiptBoundaryReady:    boundaryReady,
		ConsumerEnablementGateReady:          gateReady,
		OpaqueAuthorizationReceiptModeled:    boundaryReady,
		ConsumerAuthorizationReady:           boundaryReady,
		ReceiptPresent:                       false,
		ReceiptAccepted:                      false,
		AuthorizationAccepted:                false,
		ConsumerConsumptionAuthorized:        false,
		ConsumerEnablementAuthorized:         false,
		KDEConsumerEnabled:                   false,
		RuntimeConsumerEnabled:               false,
		LookupRouteAuthorized:                false,
		LookupRouteEnabled:                   false,
		LookupRoutePersisted:                 false,
		OpaqueLookupEnabled:                  false,
		OpaqueLookupPersisted:                false,
		RedactedSummaryPersisted:             false,
		RawResultExposed:                     false,
		RuntimeDiagnosticsPersisted:          false,
		DryRunResultPersisted:                false,
		DispatchDryRunExecuted:               false,
		RequestObjectCreationEnabled:         false,
		RequestObjectDispatchEnabled:         false,
		PortalRequestCreated:                 false,
		NotificationActionEnabled:            false,
		CompatibilityCenterOpened:            false,
		SupportBundleExported:                false,
		SupportCaseCreated:                   false,
		ProductionReadiness:                  false,
		ProductionOwnershipReady:             false,
		ReceiptItemCount:                     len(items),
		RequiredReceiptItemCount:             5,
		ReadyReceiptItemCount:                productionReceiptNotificationActionDryRunResultLookupConsumerEnablementAuthorizationReceiptReadyCount(items),
		MissingReceiptItemCount:              productionReceiptNotificationActionDryRunResultLookupConsumerEnablementAuthorizationReceiptMissingCount(items),
		ReceiptItems:                         items,
		ReceiptItemIDs:                       productionReceiptNotificationActionDryRunResultLookupConsumerEnablementAuthorizationReceiptItemIDs(items),
		RequiredBeforeReceiptAcceptance: []string{
			"accepted operator authorization for consumer enablement",
			"consumer enablement gate evidence still ready",
			"receipt writer authorization reviewed separately",
			"receipt persistence and replay reviewed separately",
			"consumer enablement remains disabled until an accepted receipt is consumed by a separate gate",
		},
		RuntimeOwned:                 true,
		GoRuntimeBacked:              true,
		KDEPolicyOwner:               false,
		OfficialDesktopOnly:          true,
		PlasmaForkRequired:           false,
		PlasmaSourceModified:         false,
		SystemServiceStarted:         false,
		SessionBusClaimed:            false,
		ProductionBusClaimed:         false,
		ProductionOwnerEnabled:       false,
		ProductionActivationReady:    false,
		WriteMethodsEnabled:          false,
		RuntimeWritesEnabled:         false,
		RequestObjectsCreated:        false,
		RequestObjectsDispatched:     false,
		NotificationSent:             false,
		NotificationDeliveryEnabled:  false,
		CompatibilityCenterPersisted: false,
		ReceiptWriterEnabled:         false,
		ReceiptPersistenceEnabled:    false,
		DesktopFilesWritten:          false,
		MIMEAppsWritten:              false,
		ShellConfigurationWritten:    false,
		SettingsPersisted:            false,
		AdapterInvocationEnabled:     false,
		BackendLaunchEnabled:         false,
		BackendProcessStarted:        false,
		SnapshotRestoreExecuted:      false,
		StateCleanupExecuted:         false,
		NetworkRequired:              false,
		HostRootModified:             false,
		PrivilegedContainerRequired:  false,
		CallerStateRootRequired:      false,
		StateRootPathExposed:         false,
		FilePathsExposed:             false,
		FileContentRead:              false,
		RawCommandExposed:            false,
		RawExecutableExposed:         false,
		BackendDetailsExposed:        false,
		BlockedActions: []string{
			"treat this authorization receipt audit as an accepted receipt",
			"authorize or enable KDE or Runtime consumers from this audit",
			"grant lookup route authorization, enable lookup routes, enable opaque lookup, persist lookup state, persist receipts, or persist redacted summaries",
			"expose raw dry-run result data, state-root paths, host paths, file contents, backend details, raw commands, or raw executables",
			"create request objects, dispatch actions, create Portal requests, send notifications, claim production ownership, launch compatibility engines, or mutate host root",
		},
		NextRequirements: []string{
			"Implement receipt writer authorization before any consumer enablement receipt can be written.",
			"Implement receipt persistence, replay protection, expiry, and revocation before any receipt can be accepted.",
			"Implement a separate acceptance gate that consumes a real accepted receipt while keeping consumer enablement disabled until reviewed.",
			"Keep KDE and Runtime consumers disabled until receipt acceptance and consumer enablement are both reviewed.",
		},
		DesktopSafeSummary: "The production receipt notification action dry-run result lookup consumer enablement authorization receipt audit models the opaque receipt boundary required before any consumer enablement gate can authorize KDE or Runtime consumers, but it has no receipt present, accepts no receipt, grants no authorization, enables no consumer or lookup route, persists no state, exposes no raw result data, starts no services, launches no engines, and mutates no host state.",
	}
	preview.AcceptedReceiptItemCount = productionReceiptNotificationActionDryRunResultLookupConsumerEnablementAuthorizationReceiptAcceptedCount(items)
	preview.AuthorizedConsumerItemCount = productionReceiptNotificationActionDryRunResultLookupConsumerEnablementAuthorizationReceiptAuthorizedCount(items)
	preview.EnabledConsumerItemCount = productionReceiptNotificationActionDryRunResultLookupConsumerEnablementAuthorizationReceiptEnabledCount(items)
	preview.PersistedReceiptItemCount = productionReceiptNotificationActionDryRunResultLookupConsumerEnablementAuthorizationReceiptPersistedCount(items)
	preview.RawExposedReceiptItemCount = productionReceiptNotificationActionDryRunResultLookupConsumerEnablementAuthorizationReceiptRawExposedCount(items)
	preview.SideEffectReceiptItemCount = productionReceiptNotificationActionDryRunResultLookupConsumerEnablementAuthorizationReceiptSideEffectCount(items)
	checks := productionReceiptNotificationActionDryRunResultLookupConsumerEnablementAuthorizationReceiptAuditChecks(preview)
	preview.Checks = checks
	preview.CheckIDs = productionReceiptNotificationActionDryRunResultLookupConsumerEnablementAuthorizationReceiptCheckIDs(checks)
	preview.Counts = countProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementAuthorizationReceiptChecks(checks)
	if preview.Counts.Blocked == 0 {
		preview.AuditDecision = "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-authorization-receipt-audit-ready-receipt-disabled"
	}
	if err := validateNoBackendTerms(preview, "production receipt notification action dry-run result lookup consumer enablement authorization receipt audit preview"); err != nil {
		return ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementAuthorizationReceiptAuditPreview{}, err
	}
	return preview, nil
}

type productionReceiptNotificationActionDryRunResultLookupConsumerEnablementAuthorizationReceiptAuditSourceSet struct {
	ConsumerEnablementGateAudit string
	DispatchSheet               string
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementAuthorizationReceiptAuditSources(root string) productionReceiptNotificationActionDryRunResultLookupConsumerEnablementAuthorizationReceiptAuditSourceSet {
	return productionReceiptNotificationActionDryRunResultLookupConsumerEnablementAuthorizationReceiptAuditSourceSet{
		ConsumerEnablementGateAudit: productionAuthorizationReadSources(root, []string{"internal/runtime/owner/production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_gate_audit.go"}),
		DispatchSheet:               productionAuthorizationReadSources(root, []string{"docs/claude-code-current-dispatch-picks.md"}),
	}
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementAuthorizationReceiptItems(sources productionReceiptNotificationActionDryRunResultLookupConsumerEnablementAuthorizationReceiptAuditSourceSet) []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementAuthorizationReceiptAuditItem {
	combined := sources.ConsumerEnablementGateAudit + sources.DispatchSheet
	return []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementAuthorizationReceiptAuditItem{
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementAuthorizationReceiptItem("review-receipt-dry-run-result-lookup-consumer-enablement-authorization-receipt", "review", "kde-review-and-runtime-diagnostics", "review-consumer-enablement-authorization-receipt", combined, []string{"review-receipt-dry-run-result-lookup-consumer-enablement-gate", "consumer enablement authorization receipt"}, "review consumer enablement requires a separately accepted receipt"),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementAuthorizationReceiptItem("renew-receipt-dry-run-result-lookup-consumer-enablement-authorization-receipt", "renew", "kde-renewal-and-runtime-diagnostics", "renewal-consumer-enablement-authorization-receipt", combined, []string{"renew-receipt-dry-run-result-lookup-consumer-enablement-gate", "consumer enablement authorization receipt"}, "renewal consumer enablement requires a separately accepted receipt"),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementAuthorizationReceiptItem("open-compatibility-center-dry-run-result-lookup-consumer-enablement-authorization-receipt", "open-compatibility-center", "kde-navigation-and-runtime-diagnostics", "navigation-consumer-enablement-authorization-receipt", combined, []string{"open-compatibility-center-dry-run-result-lookup-consumer-enablement-gate", "consumer enablement authorization receipt"}, "navigation consumer enablement requires a separately accepted receipt"),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementAuthorizationReceiptItem("dismiss-receipt-dry-run-result-lookup-consumer-enablement-authorization-receipt", "dismiss", "kde-dismissal-and-runtime-diagnostics", "dismissal-consumer-enablement-authorization-receipt", combined, []string{"dismiss-receipt-dry-run-result-lookup-consumer-enablement-gate", "consumer enablement authorization receipt"}, "dismissal consumer enablement requires a separately accepted receipt"),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementAuthorizationReceiptItem("support-info-dry-run-result-lookup-consumer-enablement-authorization-receipt", "support-info", "kde-support-and-runtime-diagnostics", "support-info-consumer-enablement-authorization-receipt", combined, []string{"support-info-dry-run-result-lookup-consumer-enablement-gate", "consumer enablement authorization receipt"}, "support-info consumer enablement requires a separately accepted receipt"),
	}
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementAuthorizationReceiptItem(id string, actionKind string, surfaceKind string, receiptKind string, source string, tokens []string, nextRequirement string) ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementAuthorizationReceiptAuditItem {
	ready := productionAuthorizationHasAll(source, tokens)
	status := "missing-consumer-enablement-authorization-receipt-evidence"
	if ready {
		status = "consumer-enablement-authorization-receipt-modeled-receipt-disabled"
	}
	return ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementAuthorizationReceiptAuditItem{
		ID:                                  id,
		ActionKind:                          actionKind,
		SurfaceKind:                         surfaceKind,
		ReceiptKind:                         receiptKind,
		OpaqueReceiptID:                     ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementAuthorizationReceiptID,
		RequiredEvidence:                    "future consumer enablement requires a separate opaque authorization receipt boundary",
		EvidencePresent:                     ready,
		ConsumerEnablementGateConsumed:      ready,
		AuthorizationReceiptBoundaryModeled: ready,
		OpaqueAuthorizationReceiptModeled:   ready,
		ReceiptPresent:                      false,
		ReceiptAccepted:                     false,
		AuthorizationAccepted:               false,
		ConsumerConsumptionAuthorized:       false,
		ConsumerEnablementAuthorized:        false,
		KDEConsumerEnabled:                  false,
		RuntimeConsumerEnabled:              false,
		LookupRouteAuthorized:               false,
		LookupRouteEnabled:                  false,
		LookupRoutePersisted:                false,
		OpaqueLookupEnabled:                 false,
		OpaqueLookupPersisted:               false,
		RedactedSummaryPersisted:            false,
		RawResultExposed:                    false,
		RuntimeDiagnosticsPersisted:         false,
		DryRunResultPersisted:               false,
		DispatchDryRunExecuted:              false,
		UserVisible:                         ready,
		ReviewOnly:                          true,
		RuntimeOwned:                        true,
		GoRuntimeBacked:                     true,
		KDEPolicyOwner:                      false,
		CallerStateRootRequired:             false,
		StateRootPathExposed:                false,
		FilePathsExposed:                    false,
		FileContentRead:                     false,
		RequestObjectCreated:                false,
		RequestObjectDispatched:             false,
		PortalRequestCreated:                false,
		NotificationActionEnabled:           false,
		CompatibilityCenterOpened:           false,
		SupportBundleExported:               false,
		SupportCaseCreated:                  false,
		ProductionReadiness:                 false,
		ProductionOwnershipReady:            false,
		SideEffectsDisabled:                 true,
		HostRootModified:                    false,
		InternalDetailsExposed:              false,
		ReceiptStatus:                       status,
		NextRequirement:                     nextRequirement,
	}
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementAuthorizationReceiptAuditChecks(preview ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementAuthorizationReceiptAuditPreview) []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementAuthorizationReceiptAuditCheck {
	return []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementAuthorizationReceiptAuditCheck{
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementAuthorizationReceiptCheck("consumer-enablement-gate-audit-consumed", productionAuthorizationPassBlocked(preview.ConsumerEnablementGateAuditConsumed), "The authorization receipt audit consumes the fail-closed consumer enablement gate."),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementAuthorizationReceiptCheck("authorization-receipt-guidance-consumed", productionAuthorizationPassBlocked(preview.AuthorizationReceiptGuidanceConsumed), "The authorization receipt audit consumes current dispatch guidance."),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementAuthorizationReceiptCheck("authorization-receipt-boundary-modeled-only", productionAuthorizationPassBlocked(preview.AuthorizationReceiptAuditRequired && preview.AuthorizationReceiptAuditModeled && preview.AuthorizationReceiptBoundaryReady && preview.ConsumerEnablementGateReady && preview.OpaqueAuthorizationReceiptModeled && preview.ConsumerAuthorizationReady && !preview.ReceiptPresent && !preview.ReceiptAccepted && !preview.AuthorizationAccepted && !preview.ConsumerEnablementAuthorized), "The opaque authorization receipt boundary is modeled without accepting a receipt or granting consumer authorization."),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementAuthorizationReceiptCheck("five-authorization-receipt-items-present", productionAuthorizationPassBlocked(preview.ReceiptItemCount == 5 && preview.RequiredReceiptItemCount == 5 && preview.MissingReceiptItemCount == 0), "Review, renew, open Compatibility Center, dismiss, and support-info receipt boundaries are modeled."),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementAuthorizationReceiptCheck("authorization-receipt-items-ready-receipt-disabled", productionAuthorizationPassBlocked(preview.ReadyReceiptItemCount == 5 && productionReceiptNotificationActionDryRunResultLookupConsumerEnablementAuthorizationReceiptItemsReady(preview.ReceiptItems)), "Every receipt boundary is ready while receipts and consumer authorization remain disabled."),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementAuthorizationReceiptCheck("receipt-acceptance-consumer-and-lookup-disabled", productionAuthorizationPassBlocked(!preview.ReceiptPresent && !preview.ReceiptAccepted && !preview.AuthorizationAccepted && !preview.ConsumerConsumptionAuthorized && !preview.ConsumerEnablementAuthorized && !preview.KDEConsumerEnabled && !preview.RuntimeConsumerEnabled && !preview.LookupRouteAuthorized && !preview.LookupRouteEnabled && !preview.OpaqueLookupEnabled && !preview.RedactedSummaryPersisted && !preview.DryRunResultPersisted && preview.AcceptedReceiptItemCount == 0 && preview.AuthorizedConsumerItemCount == 0 && preview.EnabledConsumerItemCount == 0 && preview.PersistedReceiptItemCount == 0), "Receipt acceptance, consumer authorization, lookup enablement, redacted summaries, and result persistence remain disabled."),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementAuthorizationReceiptCheck("request-notification-navigation-and-support-disabled", productionAuthorizationPassBlocked(!preview.RequestObjectCreationEnabled && !preview.RequestObjectDispatchEnabled && !preview.RequestObjectsCreated && !preview.RequestObjectsDispatched && !preview.PortalRequestCreated && !preview.NotificationActionEnabled && !preview.NotificationSent && !preview.CompatibilityCenterOpened && !preview.SupportBundleExported && !preview.SupportCaseCreated), "Request creation, dispatch, Portal requests, notifications, navigation, and support writes remain disabled."),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementAuthorizationReceiptCheck("raw-result-and-path-exposure-disabled", productionAuthorizationPassBlocked(!preview.RawResultExposed && !preview.CallerStateRootRequired && !preview.StateRootPathExposed && !preview.FilePathsExposed && !preview.FileContentRead && preview.RawExposedReceiptItemCount == 0), "Raw results, state-root paths, host paths, and file contents remain hidden."),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementAuthorizationReceiptCheck("production-and-host-boundary-closed", productionAuthorizationPassBlocked(!preview.SystemServiceStarted && !preview.SessionBusClaimed && !preview.ProductionBusClaimed && !preview.ProductionOwnerEnabled && !preview.ProductionActivationReady && !preview.ProductionReadiness && !preview.ProductionOwnershipReady && !preview.WriteMethodsEnabled && !preview.RuntimeWritesEnabled && !preview.AdapterInvocationEnabled && !preview.BackendLaunchEnabled && !preview.BackendProcessStarted && !preview.NetworkRequired && !preview.HostRootModified && !preview.PrivilegedContainerRequired && !preview.RawCommandExposed && !preview.RawExecutableExposed && !preview.BackendDetailsExposed && preview.SideEffectReceiptItemCount == 0 && productionReceiptNotificationActionDryRunResultLookupConsumerEnablementAuthorizationReceiptItemsKeepHostClosed(preview.ReceiptItems)), "Production ownership, writes, engine launch, unsafe data exposure, network, privilege, internal details, and host mutation remain disabled."),
	}
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementAuthorizationReceiptCheck(id string, status string, summary string) ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementAuthorizationReceiptAuditCheck {
	return ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementAuthorizationReceiptAuditCheck{ID: id, Status: status, Summary: summary}
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementAuthorizationReceiptGateReady(source string) bool {
	return productionAuthorizationHasAll(source, []string{
		"production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-gate-audit-preview",
		"production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-gate-audit-ready-consumers-disabled",
		"review-receipt-dry-run-result-lookup-consumer-enablement-gate",
		"renew-receipt-dry-run-result-lookup-consumer-enablement-gate",
		"open-compatibility-center-dry-run-result-lookup-consumer-enablement-gate",
		"dismiss-receipt-dry-run-result-lookup-consumer-enablement-gate",
		"support-info-dry-run-result-lookup-consumer-enablement-gate",
		"ConsumerEnablementAuthorized",
	})
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementAuthorizationReceiptGuidanceReady(source string) bool {
	return productionAuthorizationHasAll(source, []string{"consumer enablement authorization receipt audit", "opaque result identifiers", "receipt acceptance", "consumer authorization", "side effects"})
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementAuthorizationReceiptReadyCount(items []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementAuthorizationReceiptAuditItem) int {
	count := 0
	for _, item := range items {
		if item.EvidencePresent {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementAuthorizationReceiptMissingCount(items []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementAuthorizationReceiptAuditItem) int {
	count := 0
	for _, item := range items {
		if !item.EvidencePresent {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementAuthorizationReceiptAcceptedCount(items []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementAuthorizationReceiptAuditItem) int {
	count := 0
	for _, item := range items {
		if item.ReceiptPresent || item.ReceiptAccepted || item.AuthorizationAccepted {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementAuthorizationReceiptAuthorizedCount(items []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementAuthorizationReceiptAuditItem) int {
	count := 0
	for _, item := range items {
		if item.ConsumerConsumptionAuthorized || item.ConsumerEnablementAuthorized || item.LookupRouteAuthorized {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementAuthorizationReceiptEnabledCount(items []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementAuthorizationReceiptAuditItem) int {
	count := 0
	for _, item := range items {
		if item.KDEConsumerEnabled || item.RuntimeConsumerEnabled || item.LookupRouteEnabled || item.OpaqueLookupEnabled {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementAuthorizationReceiptPersistedCount(items []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementAuthorizationReceiptAuditItem) int {
	count := 0
	for _, item := range items {
		if item.LookupRoutePersisted || item.OpaqueLookupPersisted || item.RedactedSummaryPersisted || item.RuntimeDiagnosticsPersisted || item.DryRunResultPersisted {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementAuthorizationReceiptRawExposedCount(items []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementAuthorizationReceiptAuditItem) int {
	count := 0
	for _, item := range items {
		if item.RawResultExposed || item.CallerStateRootRequired || item.StateRootPathExposed || item.FilePathsExposed || item.FileContentRead || item.InternalDetailsExposed {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementAuthorizationReceiptSideEffectCount(items []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementAuthorizationReceiptAuditItem) int {
	count := 0
	for _, item := range items {
		if item.ReceiptPresent || item.ReceiptAccepted || item.AuthorizationAccepted || item.ConsumerConsumptionAuthorized || item.ConsumerEnablementAuthorized || item.KDEConsumerEnabled || item.RuntimeConsumerEnabled || item.LookupRouteAuthorized || item.LookupRouteEnabled || item.LookupRoutePersisted || item.OpaqueLookupEnabled || item.OpaqueLookupPersisted || item.RedactedSummaryPersisted || item.RawResultExposed || item.RuntimeDiagnosticsPersisted || item.DryRunResultPersisted || item.DispatchDryRunExecuted || item.RequestObjectCreated || item.RequestObjectDispatched || item.PortalRequestCreated || item.NotificationActionEnabled || item.CompatibilityCenterOpened || item.SupportBundleExported || item.SupportCaseCreated || item.HostRootModified || item.InternalDetailsExposed {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementAuthorizationReceiptItemIDs(items []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementAuthorizationReceiptAuditItem) []string {
	ids := make([]string, 0, len(items))
	for _, item := range items {
		ids = append(ids, item.ID)
	}
	return ids
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementAuthorizationReceiptCheckIDs(checks []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementAuthorizationReceiptAuditCheck) []string {
	ids := make([]string, 0, len(checks))
	for _, check := range checks {
		ids = append(ids, check.ID)
	}
	return ids
}

func countProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementAuthorizationReceiptChecks(checks []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementAuthorizationReceiptAuditCheck) ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementAuthorizationReceiptAuditCounts {
	counts := ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementAuthorizationReceiptAuditCounts{Total: len(checks)}
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

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementAuthorizationReceiptItemsReady(items []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementAuthorizationReceiptAuditItem) bool {
	if len(items) != 5 {
		return false
	}
	for _, item := range items {
		if !item.EvidencePresent || !item.ConsumerEnablementGateConsumed || !item.AuthorizationReceiptBoundaryModeled || !item.OpaqueAuthorizationReceiptModeled || item.ReceiptStatus != "consumer-enablement-authorization-receipt-modeled-receipt-disabled" {
			return false
		}
		if item.ReceiptPresent || item.ReceiptAccepted || item.AuthorizationAccepted || item.ConsumerConsumptionAuthorized || item.ConsumerEnablementAuthorized || item.KDEConsumerEnabled || item.RuntimeConsumerEnabled || item.LookupRouteAuthorized || item.LookupRouteEnabled || item.LookupRoutePersisted || item.OpaqueLookupEnabled || item.OpaqueLookupPersisted || item.RedactedSummaryPersisted || item.RawResultExposed || item.RuntimeDiagnosticsPersisted || item.DryRunResultPersisted || item.DispatchDryRunExecuted || item.CallerStateRootRequired || item.StateRootPathExposed || item.FilePathsExposed || item.FileContentRead || item.RequestObjectCreated || item.RequestObjectDispatched || item.PortalRequestCreated || item.NotificationActionEnabled || item.CompatibilityCenterOpened || item.SupportBundleExported || item.SupportCaseCreated || item.ProductionReadiness || item.ProductionOwnershipReady || item.KDEPolicyOwner || !item.ReviewOnly || !item.SideEffectsDisabled {
			return false
		}
	}
	return true
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementAuthorizationReceiptItemsKeepHostClosed(items []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementAuthorizationReceiptAuditItem) bool {
	for _, item := range items {
		if item.CallerStateRootRequired || item.StateRootPathExposed || item.FilePathsExposed || item.FileContentRead || item.HostRootModified || item.InternalDetailsExposed {
			return false
		}
	}
	return true
}
