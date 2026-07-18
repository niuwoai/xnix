package owner

type ProductionReceiptNotificationActionDryRunResultLookupConsumerRedactionAuditPreview struct {
	Version                               string                                                                             `json:"version"`
	SchemaVersion                         string                                                                             `json:"schema_version"`
	RequestType                           string                                                                             `json:"request_type"`
	AuditType                             string                                                                             `json:"audit_type"`
	Source                                string                                                                             `json:"source"`
	AuditDecision                         string                                                                             `json:"audit_decision"`
	ConsumerRedactionAuditRequired        bool                                                                               `json:"consumer_redaction_audit_required"`
	ConsumerRedactionAuditModeled         bool                                                                               `json:"consumer_redaction_audit_modeled"`
	LookupRouteAuthorizationAuditConsumed bool                                                                               `json:"lookup_route_authorization_audit_consumed"`
	ConsumerRedactionGuidanceConsumed     bool                                                                               `json:"consumer_redaction_guidance_consumed"`
	ConsumerRedactionReady                bool                                                                               `json:"consumer_redaction_ready"`
	KDEConsumerRedactionModeled           bool                                                                               `json:"kde_consumer_redaction_modeled"`
	RuntimeConsumerRedactionModeled       bool                                                                               `json:"runtime_consumer_redaction_modeled"`
	OpaqueResultIDSupported               bool                                                                               `json:"opaque_result_id_supported"`
	ConsumerConsumptionAuthorized         bool                                                                               `json:"consumer_consumption_authorized"`
	KDEConsumerEnabled                    bool                                                                               `json:"kde_consumer_enabled"`
	RuntimeConsumerEnabled                bool                                                                               `json:"runtime_consumer_enabled"`
	LookupRouteAuthorized                 bool                                                                               `json:"lookup_route_authorized"`
	LookupRouteEnabled                    bool                                                                               `json:"lookup_route_enabled"`
	LookupRoutePersisted                  bool                                                                               `json:"lookup_route_persisted"`
	OpaqueLookupEnabled                   bool                                                                               `json:"opaque_lookup_enabled"`
	OpaqueLookupPersisted                 bool                                                                               `json:"opaque_lookup_persisted"`
	RedactedSummaryPersisted              bool                                                                               `json:"redacted_summary_persisted"`
	RawResultExposed                      bool                                                                               `json:"raw_result_exposed"`
	RuntimeDiagnosticsPersisted           bool                                                                               `json:"runtime_diagnostics_persisted"`
	DryRunResultPersisted                 bool                                                                               `json:"dry_run_result_persisted"`
	DispatchDryRunExecuted                bool                                                                               `json:"dispatch_dry_run_executed"`
	RequestObjectCreationEnabled          bool                                                                               `json:"request_object_creation_enabled"`
	RequestObjectDispatchEnabled          bool                                                                               `json:"request_object_dispatch_enabled"`
	PortalRequestCreated                  bool                                                                               `json:"portal_request_created"`
	NotificationActionEnabled             bool                                                                               `json:"notification_action_enabled"`
	CompatibilityCenterOpened             bool                                                                               `json:"compatibility_center_opened"`
	SupportBundleExported                 bool                                                                               `json:"support_bundle_exported"`
	SupportCaseCreated                    bool                                                                               `json:"support_case_created"`
	ProductionReadiness                   bool                                                                               `json:"production_readiness"`
	ProductionOwnershipReady              bool                                                                               `json:"production_ownership_ready"`
	ConsumerItemCount                     int                                                                                `json:"consumer_item_count"`
	RequiredConsumerItemCount             int                                                                                `json:"required_consumer_item_count"`
	ReadyConsumerItemCount                int                                                                                `json:"ready_consumer_item_count"`
	MissingConsumerItemCount              int                                                                                `json:"missing_consumer_item_count"`
	RedactedConsumerItemCount             int                                                                                `json:"redacted_consumer_item_count"`
	EnabledConsumerItemCount              int                                                                                `json:"enabled_consumer_item_count"`
	RawExposedConsumerItemCount           int                                                                                `json:"raw_exposed_consumer_item_count"`
	PersistedConsumerItemCount            int                                                                                `json:"persisted_consumer_item_count"`
	SideEffectConsumerItemCount           int                                                                                `json:"side_effect_consumer_item_count"`
	ConsumerItems                         []ProductionReceiptNotificationActionDryRunResultLookupConsumerRedactionAuditItem  `json:"consumer_items"`
	ConsumerItemIDs                       []string                                                                           `json:"consumer_item_ids"`
	RequiredBeforeConsumerEnablement      []string                                                                           `json:"required_before_consumer_enablement"`
	Checks                                []ProductionReceiptNotificationActionDryRunResultLookupConsumerRedactionAuditCheck `json:"checks"`
	CheckIDs                              []string                                                                           `json:"check_ids"`
	Counts                                ProductionReceiptNotificationActionDryRunResultLookupConsumerRedactionAuditCounts  `json:"counts"`
	RuntimeOwned                          bool                                                                               `json:"runtime_owned"`
	GoRuntimeBacked                       bool                                                                               `json:"go_runtime_backed"`
	KDEPolicyOwner                        bool                                                                               `json:"kde_policy_owner"`
	OfficialDesktopOnly                   bool                                                                               `json:"official_desktop_only"`
	PlasmaForkRequired                    bool                                                                               `json:"plasma_fork_required"`
	PlasmaSourceModified                  bool                                                                               `json:"plasma_source_modified"`
	SystemServiceStarted                  bool                                                                               `json:"system_service_started"`
	SessionBusClaimed                     bool                                                                               `json:"session_bus_claimed"`
	ProductionBusClaimed                  bool                                                                               `json:"production_bus_claimed"`
	ProductionOwnerEnabled                bool                                                                               `json:"production_owner_enabled"`
	ProductionActivationReady             bool                                                                               `json:"production_activation_ready"`
	WriteMethodsEnabled                   bool                                                                               `json:"write_methods_enabled"`
	RuntimeWritesEnabled                  bool                                                                               `json:"runtime_writes_enabled"`
	RequestObjectsCreated                 bool                                                                               `json:"request_objects_created"`
	RequestObjectsDispatched              bool                                                                               `json:"request_objects_dispatched"`
	NotificationSent                      bool                                                                               `json:"notification_sent"`
	NotificationDeliveryEnabled           bool                                                                               `json:"notification_delivery_enabled"`
	CompatibilityCenterPersisted          bool                                                                               `json:"compatibility_center_persisted"`
	ReceiptWriterEnabled                  bool                                                                               `json:"receipt_writer_enabled"`
	ReceiptPersistenceEnabled             bool                                                                               `json:"receipt_persistence_enabled"`
	DesktopFilesWritten                   bool                                                                               `json:"desktop_files_written"`
	MIMEAppsWritten                       bool                                                                               `json:"mimeapps_written"`
	ShellConfigurationWritten             bool                                                                               `json:"shell_configuration_written"`
	SettingsPersisted                     bool                                                                               `json:"settings_persisted"`
	AdapterInvocationEnabled              bool                                                                               `json:"adapter_invocation_enabled"`
	BackendLaunchEnabled                  bool                                                                               `json:"backend_launch_enabled"`
	BackendProcessStarted                 bool                                                                               `json:"backend_process_started"`
	SnapshotRestoreExecuted               bool                                                                               `json:"snapshot_restore_executed"`
	StateCleanupExecuted                  bool                                                                               `json:"state_cleanup_executed"`
	NetworkRequired                       bool                                                                               `json:"network_required"`
	HostRootModified                      bool                                                                               `json:"host_root_modified"`
	PrivilegedContainerRequired           bool                                                                               `json:"privileged_container_required"`
	CallerStateRootRequired               bool                                                                               `json:"caller_state_root_required"`
	StateRootPathExposed                  bool                                                                               `json:"state_root_path_exposed"`
	FilePathsExposed                      bool                                                                               `json:"file_paths_exposed"`
	FileContentRead                       bool                                                                               `json:"file_content_read"`
	RawCommandExposed                     bool                                                                               `json:"raw_command_exposed"`
	RawExecutableExposed                  bool                                                                               `json:"raw_executable_exposed"`
	BackendDetailsExposed                 bool                                                                               `json:"backend_details_exposed"`
	BlockedActions                        []string                                                                           `json:"blocked_actions"`
	NextRequirements                      []string                                                                           `json:"next_requirements"`
	DesktopSafeSummary                    string                                                                             `json:"desktop_safe_summary"`
}

type ProductionReceiptNotificationActionDryRunResultLookupConsumerRedactionAuditItem struct {
	ID                              string `json:"id"`
	ActionKind                      string `json:"action_kind"`
	SurfaceKind                     string `json:"surface_kind"`
	ConsumerKind                    string `json:"consumer_kind"`
	OpaqueResultID                  string `json:"opaque_result_id"`
	RequiredEvidence                string `json:"required_evidence"`
	EvidencePresent                 bool   `json:"evidence_present"`
	ConsumerRedactionModeled        bool   `json:"consumer_redaction_modeled"`
	KDEConsumerRedactionModeled     bool   `json:"kde_consumer_redaction_modeled"`
	RuntimeConsumerRedactionModeled bool   `json:"runtime_consumer_redaction_modeled"`
	OpaqueResultIDSupported         bool   `json:"opaque_result_id_supported"`
	ConsumerConsumptionAuthorized   bool   `json:"consumer_consumption_authorized"`
	KDEConsumerEnabled              bool   `json:"kde_consumer_enabled"`
	RuntimeConsumerEnabled          bool   `json:"runtime_consumer_enabled"`
	LookupRouteAuthorized           bool   `json:"lookup_route_authorized"`
	LookupRouteEnabled              bool   `json:"lookup_route_enabled"`
	LookupRoutePersisted            bool   `json:"lookup_route_persisted"`
	OpaqueLookupEnabled             bool   `json:"opaque_lookup_enabled"`
	OpaqueLookupPersisted           bool   `json:"opaque_lookup_persisted"`
	RedactedSummaryPersisted        bool   `json:"redacted_summary_persisted"`
	RawResultExposed                bool   `json:"raw_result_exposed"`
	RuntimeDiagnosticsPersisted     bool   `json:"runtime_diagnostics_persisted"`
	DryRunResultPersisted           bool   `json:"dry_run_result_persisted"`
	DispatchDryRunExecuted          bool   `json:"dispatch_dry_run_executed"`
	UserVisible                     bool   `json:"user_visible"`
	ReviewOnly                      bool   `json:"review_only"`
	RuntimeOwned                    bool   `json:"runtime_owned"`
	GoRuntimeBacked                 bool   `json:"go_runtime_backed"`
	KDEPolicyOwner                  bool   `json:"kde_policy_owner"`
	CallerStateRootRequired         bool   `json:"caller_state_root_required"`
	StateRootPathExposed            bool   `json:"state_root_path_exposed"`
	FilePathsExposed                bool   `json:"file_paths_exposed"`
	FileContentRead                 bool   `json:"file_content_read"`
	RequestObjectCreated            bool   `json:"request_object_created"`
	RequestObjectDispatched         bool   `json:"request_object_dispatched"`
	PortalRequestCreated            bool   `json:"portal_request_created"`
	NotificationActionEnabled       bool   `json:"notification_action_enabled"`
	CompatibilityCenterOpened       bool   `json:"compatibility_center_opened"`
	SupportBundleExported           bool   `json:"support_bundle_exported"`
	SupportCaseCreated              bool   `json:"support_case_created"`
	ProductionReadiness             bool   `json:"production_readiness"`
	ProductionOwnershipReady        bool   `json:"production_ownership_ready"`
	SideEffectsDisabled             bool   `json:"side_effects_disabled"`
	HostRootModified                bool   `json:"host_root_modified"`
	InternalDetailsExposed          bool   `json:"internal_details_exposed"`
	ConsumerStatus                  string `json:"consumer_status"`
	NextRequirement                 string `json:"next_requirement"`
}

type ProductionReceiptNotificationActionDryRunResultLookupConsumerRedactionAuditCheck struct {
	ID      string `json:"id"`
	Status  string `json:"status"`
	Summary string `json:"summary"`
}

type ProductionReceiptNotificationActionDryRunResultLookupConsumerRedactionAuditCounts struct {
	Total   int `json:"total"`
	Passed  int `json:"passed"`
	Pending int `json:"pending"`
	Blocked int `json:"blocked"`
}

func NewProductionReceiptNotificationActionDryRunResultLookupConsumerRedactionAuditPreview(root string) (ProductionReceiptNotificationActionDryRunResultLookupConsumerRedactionAuditPreview, error) {
	version, err := readProductionDBusGateReviewVersion(root)
	if err != nil {
		return ProductionReceiptNotificationActionDryRunResultLookupConsumerRedactionAuditPreview{}, err
	}
	sources := productionReceiptNotificationActionDryRunResultLookupConsumerRedactionAuditSources(root)
	items := productionReceiptNotificationActionDryRunResultLookupConsumerRedactionItems(sources)
	authorizationReady := productionReceiptNotificationActionDryRunResultLookupConsumerRedactionAuthorizationReady(sources.LookupRouteAuthorizationAudit)
	guidanceReady := productionReceiptNotificationActionDryRunResultLookupConsumerRedactionGuidanceReady(sources.DispatchSheet)
	consumerReady := authorizationReady && guidanceReady && productionReceiptNotificationActionDryRunResultLookupConsumerRedactionItemsReady(items)
	preview := ProductionReceiptNotificationActionDryRunResultLookupConsumerRedactionAuditPreview{
		Version:                               version,
		SchemaVersion:                         "xnix.runtime.production_receipt_notification_action_dry_run_result_lookup_consumer_redaction_audit.v1",
		RequestType:                           "production-receipt-notification-action-dry-run-result-lookup-consumer-redaction-audit-preview",
		AuditType:                             "receipt-notification-action-dry-run-result-lookup-consumer-redaction-audit",
		Source:                                "production-receipt-notification-action-dry-run-result-lookup-route-authorization-audit-preview+claude-code-current-dispatch-picks",
		AuditDecision:                         "production-receipt-notification-action-dry-run-result-lookup-consumer-redaction-audit-blocked",
		ConsumerRedactionAuditRequired:        true,
		ConsumerRedactionAuditModeled:         true,
		LookupRouteAuthorizationAuditConsumed: authorizationReady,
		ConsumerRedactionGuidanceConsumed:     guidanceReady,
		ConsumerRedactionReady:                consumerReady,
		KDEConsumerRedactionModeled:           consumerReady,
		RuntimeConsumerRedactionModeled:       consumerReady,
		OpaqueResultIDSupported:               consumerReady,
		ConsumerConsumptionAuthorized:         false,
		KDEConsumerEnabled:                    false,
		RuntimeConsumerEnabled:                false,
		LookupRouteAuthorized:                 false,
		LookupRouteEnabled:                    false,
		LookupRoutePersisted:                  false,
		OpaqueLookupEnabled:                   false,
		OpaqueLookupPersisted:                 false,
		RedactedSummaryPersisted:              false,
		RawResultExposed:                      false,
		RuntimeDiagnosticsPersisted:           false,
		DryRunResultPersisted:                 false,
		DispatchDryRunExecuted:                false,
		RequestObjectCreationEnabled:          false,
		RequestObjectDispatchEnabled:          false,
		PortalRequestCreated:                  false,
		NotificationActionEnabled:             false,
		CompatibilityCenterOpened:             false,
		SupportBundleExported:                 false,
		SupportCaseCreated:                    false,
		ProductionReadiness:                   false,
		ProductionOwnershipReady:              false,
		ConsumerItemCount:                     len(items),
		RequiredConsumerItemCount:             5,
		ReadyConsumerItemCount:                productionReceiptNotificationActionDryRunResultLookupConsumerRedactionReadyCount(items),
		MissingConsumerItemCount:              productionReceiptNotificationActionDryRunResultLookupConsumerRedactionMissingCount(items),
		RedactedConsumerItemCount:             productionReceiptNotificationActionDryRunResultLookupConsumerRedactionRedactedCount(items),
		ConsumerItems:                         items,
		ConsumerItemIDs:                       productionReceiptNotificationActionDryRunResultLookupConsumerRedactionItemIDs(items),
		RequiredBeforeConsumerEnablement: []string{
			"accepted lookup route authorization receipt",
			"separate owner-local lookup route implementation",
			"separate consumer enablement review",
			"separate lookup and redacted summary persistence review",
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
			"treat consumer redaction modeling as permission to let KDE or Runtime consume opaque dry-run result identifiers",
			"enable lookup routes, lookup persistence, redacted summary persistence, Runtime diagnostics persistence, or dry-run execution",
			"expose raw dry-run result data, file contents, state-root paths, host paths, backend details, raw commands, or raw executables",
			"create request objects, dispatch actions, create Portal requests, write receipts, claim production ownership, launch compatibility engines, or mutate host root",
		},
		NextRequirements: []string{
			"Add an accepted lookup route authorization receipt before enabling any consumer read path.",
			"Implement the owner-local lookup route separately while preserving opaque identifiers and caller path hiding.",
			"Review lookup persistence and redacted summary storage separately before any consumer data is persisted.",
			"Keep KDE and Runtime consumers disabled until a separate enablement review proves redaction enforcement.",
		},
		DesktopSafeSummary: "The production receipt notification action dry-run result lookup consumer redaction audit models how KDE and Runtime consumers should see only redacted summaries behind opaque identifiers, but it authorizes no consumers, enables no lookup route, persists no summaries, executes no dry runs, exposes no raw result data, starts no services, launches no engines, and mutates no host state.",
	}
	preview.EnabledConsumerItemCount = productionReceiptNotificationActionDryRunResultLookupConsumerRedactionEnabledCount(items)
	preview.RawExposedConsumerItemCount = productionReceiptNotificationActionDryRunResultLookupConsumerRedactionRawExposedCount(items)
	preview.PersistedConsumerItemCount = productionReceiptNotificationActionDryRunResultLookupConsumerRedactionPersistedCount(items)
	preview.SideEffectConsumerItemCount = productionReceiptNotificationActionDryRunResultLookupConsumerRedactionSideEffectCount(items)
	checks := productionReceiptNotificationActionDryRunResultLookupConsumerRedactionAuditChecks(preview)
	preview.Checks = checks
	preview.CheckIDs = productionReceiptNotificationActionDryRunResultLookupConsumerRedactionCheckIDs(checks)
	preview.Counts = countProductionReceiptNotificationActionDryRunResultLookupConsumerRedactionChecks(checks)
	if preview.Counts.Blocked == 0 {
		preview.AuditDecision = "production-receipt-notification-action-dry-run-result-lookup-consumer-redaction-audit-ready-consumers-disabled"
	}
	if err := validateNoBackendTerms(preview, "production receipt notification action dry-run result lookup consumer redaction audit preview"); err != nil {
		return ProductionReceiptNotificationActionDryRunResultLookupConsumerRedactionAuditPreview{}, err
	}
	return preview, nil
}

type productionReceiptNotificationActionDryRunResultLookupConsumerRedactionAuditSourceSet struct {
	LookupRouteAuthorizationAudit string
	DispatchSheet                 string
}

func productionReceiptNotificationActionDryRunResultLookupConsumerRedactionAuditSources(root string) productionReceiptNotificationActionDryRunResultLookupConsumerRedactionAuditSourceSet {
	return productionReceiptNotificationActionDryRunResultLookupConsumerRedactionAuditSourceSet{
		LookupRouteAuthorizationAudit: productionAuthorizationReadSources(root, []string{"internal/runtime/owner/production_receipt_notification_action_dry_run_result_lookup_route_authorization_audit.go"}),
		DispatchSheet:                 productionAuthorizationReadSources(root, []string{"docs/claude-code-current-dispatch-picks.md"}),
	}
}

func productionReceiptNotificationActionDryRunResultLookupConsumerRedactionItems(sources productionReceiptNotificationActionDryRunResultLookupConsumerRedactionAuditSourceSet) []ProductionReceiptNotificationActionDryRunResultLookupConsumerRedactionAuditItem {
	combined := sources.LookupRouteAuthorizationAudit + sources.DispatchSheet
	return []ProductionReceiptNotificationActionDryRunResultLookupConsumerRedactionAuditItem{
		productionReceiptNotificationActionDryRunResultLookupConsumerRedactionItem("review-receipt-dry-run-result-lookup-consumer-redaction", "review", "kde-review-and-runtime-diagnostics", "review-result-redacted-consumer", "review-dry-run-result-opaque-id", combined, []string{"review-receipt-dry-run-result-lookup-route-authorization", "consumer redaction"}, "prove redaction enforcement before enabling review result consumers"),
		productionReceiptNotificationActionDryRunResultLookupConsumerRedactionItem("renew-receipt-dry-run-result-lookup-consumer-redaction", "renew", "kde-renewal-and-runtime-diagnostics", "renewal-result-redacted-consumer", "renewal-dry-run-result-opaque-id", combined, []string{"renew-receipt-dry-run-result-lookup-route-authorization", "consumer redaction"}, "prove redaction enforcement before enabling renewal result consumers"),
		productionReceiptNotificationActionDryRunResultLookupConsumerRedactionItem("open-compatibility-center-dry-run-result-lookup-consumer-redaction", "open-compatibility-center", "kde-navigation-and-runtime-diagnostics", "navigation-result-redacted-consumer", "navigation-dry-run-result-opaque-id", combined, []string{"open-compatibility-center-dry-run-result-lookup-route-authorization", "consumer redaction"}, "prove redaction enforcement before enabling navigation result consumers"),
		productionReceiptNotificationActionDryRunResultLookupConsumerRedactionItem("dismiss-receipt-dry-run-result-lookup-consumer-redaction", "dismiss", "kde-dismissal-and-runtime-diagnostics", "dismissal-result-redacted-consumer", "dismissal-dry-run-result-opaque-id", combined, []string{"dismiss-receipt-dry-run-result-lookup-route-authorization", "consumer redaction"}, "prove redaction enforcement before enabling dismissal result consumers"),
		productionReceiptNotificationActionDryRunResultLookupConsumerRedactionItem("support-info-dry-run-result-lookup-consumer-redaction", "support-info", "kde-support-and-runtime-diagnostics", "support-info-result-redacted-consumer", "support-info-dry-run-result-opaque-id", combined, []string{"support-info-dry-run-result-lookup-route-authorization", "consumer redaction"}, "prove redaction enforcement before enabling support-info result consumers"),
	}
}

func productionReceiptNotificationActionDryRunResultLookupConsumerRedactionItem(id string, actionKind string, surfaceKind string, consumerKind string, opaqueResultID string, source string, tokens []string, nextRequirement string) ProductionReceiptNotificationActionDryRunResultLookupConsumerRedactionAuditItem {
	ready := productionAuthorizationHasAll(source, tokens)
	status := "missing-lookup-consumer-redaction-evidence"
	if ready {
		status = "lookup-consumer-redaction-modeled-consumers-disabled"
	}
	return ProductionReceiptNotificationActionDryRunResultLookupConsumerRedactionAuditItem{
		ID:                              id,
		ActionKind:                      actionKind,
		SurfaceKind:                     surfaceKind,
		ConsumerKind:                    consumerKind,
		OpaqueResultID:                  opaqueResultID,
		RequiredEvidence:                "future dry-run result consumers require KDE and Runtime redaction boundaries",
		EvidencePresent:                 ready,
		ConsumerRedactionModeled:        ready,
		KDEConsumerRedactionModeled:     ready,
		RuntimeConsumerRedactionModeled: ready,
		OpaqueResultIDSupported:         ready,
		ConsumerConsumptionAuthorized:   false,
		KDEConsumerEnabled:              false,
		RuntimeConsumerEnabled:          false,
		LookupRouteAuthorized:           false,
		LookupRouteEnabled:              false,
		LookupRoutePersisted:            false,
		OpaqueLookupEnabled:             false,
		OpaqueLookupPersisted:           false,
		RedactedSummaryPersisted:        false,
		RawResultExposed:                false,
		RuntimeDiagnosticsPersisted:     false,
		DryRunResultPersisted:           false,
		DispatchDryRunExecuted:          false,
		UserVisible:                     ready,
		ReviewOnly:                      true,
		RuntimeOwned:                    true,
		GoRuntimeBacked:                 true,
		KDEPolicyOwner:                  false,
		CallerStateRootRequired:         false,
		StateRootPathExposed:            false,
		FilePathsExposed:                false,
		FileContentRead:                 false,
		RequestObjectCreated:            false,
		RequestObjectDispatched:         false,
		PortalRequestCreated:            false,
		NotificationActionEnabled:       false,
		CompatibilityCenterOpened:       false,
		SupportBundleExported:           false,
		SupportCaseCreated:              false,
		ProductionReadiness:             false,
		ProductionOwnershipReady:        false,
		SideEffectsDisabled:             true,
		HostRootModified:                false,
		InternalDetailsExposed:          false,
		ConsumerStatus:                  status,
		NextRequirement:                 nextRequirement,
	}
}

func productionReceiptNotificationActionDryRunResultLookupConsumerRedactionAuditChecks(preview ProductionReceiptNotificationActionDryRunResultLookupConsumerRedactionAuditPreview) []ProductionReceiptNotificationActionDryRunResultLookupConsumerRedactionAuditCheck {
	return []ProductionReceiptNotificationActionDryRunResultLookupConsumerRedactionAuditCheck{
		productionReceiptNotificationActionDryRunResultLookupConsumerRedactionCheck("lookup-route-authorization-audit-consumed", productionAuthorizationPassBlocked(preview.LookupRouteAuthorizationAuditConsumed), "The consumer redaction audit consumes lookup route authorization evidence."),
		productionReceiptNotificationActionDryRunResultLookupConsumerRedactionCheck("consumer-redaction-guidance-consumed", productionAuthorizationPassBlocked(preview.ConsumerRedactionGuidanceConsumed), "The consumer redaction audit consumes current dispatch guidance."),
		productionReceiptNotificationActionDryRunResultLookupConsumerRedactionCheck("consumer-redaction-modeled-only", productionAuthorizationPassBlocked(preview.ConsumerRedactionAuditRequired && preview.ConsumerRedactionAuditModeled && preview.ConsumerRedactionReady && preview.KDEConsumerRedactionModeled && preview.RuntimeConsumerRedactionModeled && preview.OpaqueResultIDSupported && !preview.ConsumerConsumptionAuthorized && !preview.KDEConsumerEnabled && !preview.RuntimeConsumerEnabled), "KDE and Runtime consumer redaction is modeled without authorizing or enabling consumers."),
		productionReceiptNotificationActionDryRunResultLookupConsumerRedactionCheck("five-consumer-redaction-items-present", productionAuthorizationPassBlocked(preview.ConsumerItemCount == 5 && preview.RequiredConsumerItemCount == 5 && preview.MissingConsumerItemCount == 0), "Review, renew, open Compatibility Center, dismiss, and support-info consumer redaction boundaries are modeled."),
		productionReceiptNotificationActionDryRunResultLookupConsumerRedactionCheck("consumer-redaction-items-ready-consumers-disabled", productionAuthorizationPassBlocked(preview.ReadyConsumerItemCount == 5 && preview.RedactedConsumerItemCount == 5 && productionReceiptNotificationActionDryRunResultLookupConsumerRedactionItemsReady(preview.ConsumerItems)), "Every consumer redaction item is ready while consumers, lookup routes, and persistence remain disabled."),
		productionReceiptNotificationActionDryRunResultLookupConsumerRedactionCheck("lookup-persistence-result-persistence-and-execution-disabled", productionAuthorizationPassBlocked(!preview.LookupRouteAuthorized && !preview.LookupRouteEnabled && !preview.LookupRoutePersisted && !preview.OpaqueLookupEnabled && !preview.OpaqueLookupPersisted && !preview.RedactedSummaryPersisted && !preview.RuntimeDiagnosticsPersisted && !preview.DryRunResultPersisted && !preview.DispatchDryRunExecuted && preview.EnabledConsumerItemCount == 0 && preview.PersistedConsumerItemCount == 0), "Lookup routes, lookup persistence, redacted summaries, diagnostics persistence, result persistence, and dry-run execution remain disabled."),
		productionReceiptNotificationActionDryRunResultLookupConsumerRedactionCheck("raw-result-and-path-exposure-disabled", productionAuthorizationPassBlocked(!preview.RawResultExposed && !preview.CallerStateRootRequired && !preview.StateRootPathExposed && !preview.FilePathsExposed && !preview.FileContentRead && preview.RawExposedConsumerItemCount == 0), "Raw results, state-root paths, host paths, and file contents remain hidden."),
		productionReceiptNotificationActionDryRunResultLookupConsumerRedactionCheck("request-dispatch-notification-and-support-disabled", productionAuthorizationPassBlocked(!preview.RequestObjectCreationEnabled && !preview.RequestObjectDispatchEnabled && !preview.RequestObjectsCreated && !preview.RequestObjectsDispatched && !preview.PortalRequestCreated && !preview.NotificationActionEnabled && !preview.NotificationSent && !preview.CompatibilityCenterOpened && !preview.SupportBundleExported && !preview.SupportCaseCreated), "Request creation, dispatch, Portal requests, notifications, navigation, and support writes remain disabled."),
		productionReceiptNotificationActionDryRunResultLookupConsumerRedactionCheck("production-and-host-boundary-closed", productionAuthorizationPassBlocked(!preview.SystemServiceStarted && !preview.SessionBusClaimed && !preview.ProductionBusClaimed && !preview.ProductionOwnerEnabled && !preview.ProductionActivationReady && !preview.ProductionReadiness && !preview.ProductionOwnershipReady && !preview.WriteMethodsEnabled && !preview.RuntimeWritesEnabled && !preview.AdapterInvocationEnabled && !preview.BackendLaunchEnabled && !preview.BackendProcessStarted && !preview.NetworkRequired && !preview.HostRootModified && !preview.PrivilegedContainerRequired && !preview.RawCommandExposed && !preview.RawExecutableExposed && !preview.BackendDetailsExposed && preview.SideEffectConsumerItemCount == 0 && productionReceiptNotificationActionDryRunResultLookupConsumerRedactionItemsKeepHostClosed(preview.ConsumerItems)), "Production ownership, writes, engine launch, unsafe data exposure, network, privilege, internal details, and host mutation remain disabled."),
	}
}

func productionReceiptNotificationActionDryRunResultLookupConsumerRedactionCheck(id string, status string, summary string) ProductionReceiptNotificationActionDryRunResultLookupConsumerRedactionAuditCheck {
	return ProductionReceiptNotificationActionDryRunResultLookupConsumerRedactionAuditCheck{ID: id, Status: status, Summary: summary}
}

func productionReceiptNotificationActionDryRunResultLookupConsumerRedactionAuthorizationReady(source string) bool {
	return productionAuthorizationHasAll(source, []string{
		"production-receipt-notification-action-dry-run-result-lookup-route-authorization-audit-preview",
		"production-receipt-notification-action-dry-run-result-lookup-route-authorization-audit-ready-route-disabled",
		"review-receipt-dry-run-result-lookup-route-authorization",
		"renew-receipt-dry-run-result-lookup-route-authorization",
		"open-compatibility-center-dry-run-result-lookup-route-authorization",
		"dismiss-receipt-dry-run-result-lookup-route-authorization",
		"support-info-dry-run-result-lookup-route-authorization",
		"LookupRouteAuthorized",
	})
}

func productionReceiptNotificationActionDryRunResultLookupConsumerRedactionGuidanceReady(source string) bool {
	return productionAuthorizationHasAll(source, []string{"lookup consumer redaction audit", "opaque result identifiers", "side effects"})
}

func productionReceiptNotificationActionDryRunResultLookupConsumerRedactionReadyCount(items []ProductionReceiptNotificationActionDryRunResultLookupConsumerRedactionAuditItem) int {
	count := 0
	for _, item := range items {
		if item.EvidencePresent {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionDryRunResultLookupConsumerRedactionMissingCount(items []ProductionReceiptNotificationActionDryRunResultLookupConsumerRedactionAuditItem) int {
	count := 0
	for _, item := range items {
		if !item.EvidencePresent {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionDryRunResultLookupConsumerRedactionRedactedCount(items []ProductionReceiptNotificationActionDryRunResultLookupConsumerRedactionAuditItem) int {
	count := 0
	for _, item := range items {
		if item.ConsumerRedactionModeled && item.KDEConsumerRedactionModeled && item.RuntimeConsumerRedactionModeled && !item.RawResultExposed {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionDryRunResultLookupConsumerRedactionEnabledCount(items []ProductionReceiptNotificationActionDryRunResultLookupConsumerRedactionAuditItem) int {
	count := 0
	for _, item := range items {
		if item.ConsumerConsumptionAuthorized || item.KDEConsumerEnabled || item.RuntimeConsumerEnabled || item.LookupRouteEnabled {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionDryRunResultLookupConsumerRedactionRawExposedCount(items []ProductionReceiptNotificationActionDryRunResultLookupConsumerRedactionAuditItem) int {
	count := 0
	for _, item := range items {
		if item.RawResultExposed || item.CallerStateRootRequired || item.StateRootPathExposed || item.FilePathsExposed || item.FileContentRead || item.InternalDetailsExposed {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionDryRunResultLookupConsumerRedactionPersistedCount(items []ProductionReceiptNotificationActionDryRunResultLookupConsumerRedactionAuditItem) int {
	count := 0
	for _, item := range items {
		if item.RedactedSummaryPersisted || item.RuntimeDiagnosticsPersisted || item.DryRunResultPersisted || item.LookupRoutePersisted || item.OpaqueLookupPersisted {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionDryRunResultLookupConsumerRedactionSideEffectCount(items []ProductionReceiptNotificationActionDryRunResultLookupConsumerRedactionAuditItem) int {
	count := 0
	for _, item := range items {
		if item.ConsumerConsumptionAuthorized || item.KDEConsumerEnabled || item.RuntimeConsumerEnabled || item.LookupRouteAuthorized || item.LookupRouteEnabled || item.LookupRoutePersisted || item.OpaqueLookupEnabled || item.OpaqueLookupPersisted || item.RedactedSummaryPersisted || item.RawResultExposed || item.RuntimeDiagnosticsPersisted || item.DryRunResultPersisted || item.DispatchDryRunExecuted || item.RequestObjectCreated || item.RequestObjectDispatched || item.PortalRequestCreated || item.NotificationActionEnabled || item.CompatibilityCenterOpened || item.SupportBundleExported || item.SupportCaseCreated || item.HostRootModified || item.InternalDetailsExposed {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionDryRunResultLookupConsumerRedactionItemIDs(items []ProductionReceiptNotificationActionDryRunResultLookupConsumerRedactionAuditItem) []string {
	ids := make([]string, 0, len(items))
	for _, item := range items {
		ids = append(ids, item.ID)
	}
	return ids
}

func productionReceiptNotificationActionDryRunResultLookupConsumerRedactionCheckIDs(checks []ProductionReceiptNotificationActionDryRunResultLookupConsumerRedactionAuditCheck) []string {
	ids := make([]string, 0, len(checks))
	for _, check := range checks {
		ids = append(ids, check.ID)
	}
	return ids
}

func countProductionReceiptNotificationActionDryRunResultLookupConsumerRedactionChecks(checks []ProductionReceiptNotificationActionDryRunResultLookupConsumerRedactionAuditCheck) ProductionReceiptNotificationActionDryRunResultLookupConsumerRedactionAuditCounts {
	counts := ProductionReceiptNotificationActionDryRunResultLookupConsumerRedactionAuditCounts{Total: len(checks)}
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

func productionReceiptNotificationActionDryRunResultLookupConsumerRedactionItemsReady(items []ProductionReceiptNotificationActionDryRunResultLookupConsumerRedactionAuditItem) bool {
	if len(items) != 5 {
		return false
	}
	for _, item := range items {
		if !item.EvidencePresent || !item.ConsumerRedactionModeled || !item.KDEConsumerRedactionModeled || !item.RuntimeConsumerRedactionModeled || !item.OpaqueResultIDSupported || item.ConsumerStatus != "lookup-consumer-redaction-modeled-consumers-disabled" {
			return false
		}
		if item.ConsumerConsumptionAuthorized || item.KDEConsumerEnabled || item.RuntimeConsumerEnabled || item.LookupRouteAuthorized || item.LookupRouteEnabled || item.LookupRoutePersisted || item.OpaqueLookupEnabled || item.OpaqueLookupPersisted || item.RedactedSummaryPersisted || item.RawResultExposed || item.RuntimeDiagnosticsPersisted || item.DryRunResultPersisted || item.DispatchDryRunExecuted || item.CallerStateRootRequired || item.StateRootPathExposed || item.FilePathsExposed || item.FileContentRead || item.RequestObjectCreated || item.RequestObjectDispatched || item.PortalRequestCreated || item.NotificationActionEnabled || item.CompatibilityCenterOpened || item.SupportBundleExported || item.SupportCaseCreated || item.ProductionReadiness || item.ProductionOwnershipReady || item.KDEPolicyOwner || !item.ReviewOnly || !item.SideEffectsDisabled {
			return false
		}
	}
	return true
}

func productionReceiptNotificationActionDryRunResultLookupConsumerRedactionItemsKeepHostClosed(items []ProductionReceiptNotificationActionDryRunResultLookupConsumerRedactionAuditItem) bool {
	for _, item := range items {
		if item.CallerStateRootRequired || item.StateRootPathExposed || item.FilePathsExposed || item.FileContentRead || item.HostRootModified || item.InternalDetailsExposed {
			return false
		}
	}
	return true
}
