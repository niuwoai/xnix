package owner

type ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteAuthorizationEvidenceGatePreview struct {
	Version                                                 string                                                                                                                                                                                                                 `json:"version"`
	SchemaVersion                                           string                                                                                                                                                                                                                 `json:"schema_version"`
	RequestType                                             string                                                                                                                                                                                                                 `json:"request_type"`
	GateType                                                string                                                                                                                                                                                                                 `json:"gate_type"`
	Source                                                  string                                                                                                                                                                                                                 `json:"source"`
	AuditDecision                                           string                                                                                                                                                                                                                 `json:"audit_decision"`
	CurrentMainlineConsumed                                 bool                                                                                                                                                                                                                   `json:"current_mainline_consumed"`
	ResultConsumerProjectionEvidenceAuditConsumed           bool                                                                                                                                                                                                                   `json:"result_consumer_projection_evidence_audit_consumed"`
	ResultConsumerProjectionEvidenceAuditReady              bool                                                                                                                                                                                                                   `json:"result_consumer_projection_evidence_audit_ready"`
	RouteAuthorizationEvidenceGateRequired                  bool                                                                                                                                                                                                                   `json:"route_authorization_evidence_gate_required"`
	ResultConsumerProjectionRouteAuthorizationEvidenceGated bool                                                                                                                                                                                                                   `json:"result_consumer_projection_route_authorization_evidence_gated"`
	RouteAuthorizationEvidenceGateReady                     bool                                                                                                                                                                                                                   `json:"route_authorization_evidence_gate_ready"`
	CompatibilityCenterRouteAuthorizationEvidenceGated      bool                                                                                                                                                                                                                   `json:"compatibility_center_route_authorization_evidence_gated"`
	RuntimeDiagnosticsRouteAuthorizationEvidenceGated       bool                                                                                                                                                                                                                   `json:"runtime_diagnostics_route_authorization_evidence_gated"`
	RedactedProjectionRouteAuthorizationEvidenceGated       bool                                                                                                                                                                                                                   `json:"redacted_projection_route_authorization_evidence_gated"`
	UserVisibleRouteAuthorizationEvidenceGated              bool                                                                                                                                                                                                                   `json:"user_visible_route_authorization_evidence_gated"`
	FailureProjectionRouteAuthorizationEvidenceGated        bool                                                                                                                                                                                                                   `json:"failure_projection_route_authorization_evidence_gated"`
	KDESafeRedactedStatusOnly                               bool                                                                                                                                                                                                                   `json:"kde_safe_redacted_status_only"`
	ReceiptPresent                                          bool                                                                                                                                                                                                                   `json:"receipt_present"`
	ReceiptAccepted                                         bool                                                                                                                                                                                                                   `json:"receipt_accepted"`
	ReceiptConsumed                                         bool                                                                                                                                                                                                                   `json:"receipt_consumed"`
	StorageRootPolicyGrantAuthorized                        bool                                                                                                                                                                                                                   `json:"storage_root_policy_grant_authorized"`
	RecordWriterCallAuthorized                              bool                                                                                                                                                                                                                   `json:"record_writer_call_authorized"`
	WriterCallable                                          bool                                                                                                                                                                                                                   `json:"writer_callable"`
	StorageRootResolved                                     bool                                                                                                                                                                                                                   `json:"storage_root_resolved"`
	StorageWriteEnabled                                     bool                                                                                                                                                                                                                   `json:"storage_write_enabled"`
	StatusPersistenceWriteEnabled                           bool                                                                                                                                                                                                                   `json:"status_persistence_write_enabled"`
	EvidenceItemCount                                       int                                                                                                                                                                                                                    `json:"evidence_item_count"`
	ReadyEvidenceItemCount                                  int                                                                                                                                                                                                                    `json:"ready_evidence_item_count"`
	MissingEvidenceItemCount                                int                                                                                                                                                                                                                    `json:"missing_evidence_item_count"`
	CompatibilityCenterEvidenceItemCount                    int                                                                                                                                                                                                                    `json:"compatibility_center_evidence_item_count"`
	RuntimeDiagnosticsEvidenceItemCount                     int                                                                                                                                                                                                                    `json:"runtime_diagnostics_evidence_item_count"`
	RawExposedEvidenceItemCount                             int                                                                                                                                                                                                                    `json:"raw_exposed_evidence_item_count"`
	SideEffectEvidenceItemCount                             int                                                                                                                                                                                                                    `json:"side_effect_evidence_item_count"`
	EvidenceItems                                           []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteAuthorizationEvidenceGateItem  `json:"evidence_items"`
	EvidenceItemIDs                                         []string                                                                                                                                                                                                               `json:"evidence_item_ids"`
	Checks                                                  []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteAuthorizationEvidenceGateCheck `json:"checks"`
	CheckIDs                                                []string                                                                                                                                                                                                               `json:"check_ids"`
	Counts                                                  ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteAuthorizationEvidenceGateCounts  `json:"counts"`
	ConsumerDryRunLookupResultBoundaryAuthorized            bool                                                                                                                                                                                                                   `json:"consumer_dry_run_lookup_result_boundary_authorized"`
	ConsumerEnablementAuthorized                            bool                                                                                                                                                                                                                   `json:"consumer_enablement_authorized"`
	KDEConsumerEnabled                                      bool                                                                                                                                                                                                                   `json:"kde_consumer_enabled"`
	RuntimeConsumerEnabled                                  bool                                                                                                                                                                                                                   `json:"runtime_consumer_enabled"`
	LookupRouteAuthorized                                   bool                                                                                                                                                                                                                   `json:"lookup_route_authorized"`
	LookupRouteEnabled                                      bool                                                                                                                                                                                                                   `json:"lookup_route_enabled"`
	OpaqueLookupEnabled                                     bool                                                                                                                                                                                                                   `json:"opaque_lookup_enabled"`
	RedactedSummaryPersisted                                bool                                                                                                                                                                                                                   `json:"redacted_summary_persisted"`
	KDEStatusPersisted                                      bool                                                                                                                                                                                                                   `json:"kde_status_persisted"`
	RuntimeDiagnosticsPersisted                             bool                                                                                                                                                                                                                   `json:"runtime_diagnostics_persisted"`
	DryRunResultPersisted                                   bool                                                                                                                                                                                                                   `json:"dry_run_result_persisted"`
	RawResultExposed                                        bool                                                                                                                                                                                                                   `json:"raw_result_exposed"`
	DispatchDryRunExecuted                                  bool                                                                                                                                                                                                                   `json:"dispatch_dry_run_executed"`
	RequestObjectCreationEnabled                            bool                                                                                                                                                                                                                   `json:"request_object_creation_enabled"`
	RequestObjectDispatchEnabled                            bool                                                                                                                                                                                                                   `json:"request_object_dispatch_enabled"`
	PortalRequestCreated                                    bool                                                                                                                                                                                                                   `json:"portal_request_created"`
	NotificationActionEnabled                               bool                                                                                                                                                                                                                   `json:"notification_action_enabled"`
	CompatibilityCenterOpened                               bool                                                                                                                                                                                                                   `json:"compatibility_center_opened"`
	SupportBundleExported                                   bool                                                                                                                                                                                                                   `json:"support_bundle_exported"`
	SupportCaseCreated                                      bool                                                                                                                                                                                                                   `json:"support_case_created"`
	RuntimeOwned                                            bool                                                                                                                                                                                                                   `json:"runtime_owned"`
	GoRuntimeBacked                                         bool                                                                                                                                                                                                                   `json:"go_runtime_backed"`
	KDEPolicyOwner                                          bool                                                                                                                                                                                                                   `json:"kde_policy_owner"`
	ProductionReadiness                                     bool                                                                                                                                                                                                                   `json:"production_readiness"`
	ProductionOwnershipReady                                bool                                                                                                                                                                                                                   `json:"production_ownership_ready"`
	SystemServiceStarted                                    bool                                                                                                                                                                                                                   `json:"system_service_started"`
	SessionBusClaimed                                       bool                                                                                                                                                                                                                   `json:"session_bus_claimed"`
	ProductionBusClaimed                                    bool                                                                                                                                                                                                                   `json:"production_bus_claimed"`
	ProductionOwnerEnabled                                  bool                                                                                                                                                                                                                   `json:"production_owner_enabled"`
	WriteMethodsEnabled                                     bool                                                                                                                                                                                                                   `json:"write_methods_enabled"`
	RuntimeWritesEnabled                                    bool                                                                                                                                                                                                                   `json:"runtime_writes_enabled"`
	DesktopFilesWritten                                     bool                                                                                                                                                                                                                   `json:"desktop_files_written"`
	SettingsPersisted                                       bool                                                                                                                                                                                                                   `json:"settings_persisted"`
	AdapterInvocationEnabled                                bool                                                                                                                                                                                                                   `json:"adapter_invocation_enabled"`
	BackendLaunchEnabled                                    bool                                                                                                                                                                                                                   `json:"backend_launch_enabled"`
	BackendProcessStarted                                   bool                                                                                                                                                                                                                   `json:"backend_process_started"`
	NetworkRequired                                         bool                                                                                                                                                                                                                   `json:"network_required"`
	HostRootModified                                        bool                                                                                                                                                                                                                   `json:"host_root_modified"`
	PrivilegedContainerRequired                             bool                                                                                                                                                                                                                   `json:"privileged_container_required"`
	CallerStateRootRequired                                 bool                                                                                                                                                                                                                   `json:"caller_state_root_required"`
	StateRootPathExposed                                    bool                                                                                                                                                                                                                   `json:"state_root_path_exposed"`
	FilePathsExposed                                        bool                                                                                                                                                                                                                   `json:"file_paths_exposed"`
	FileContentRead                                         bool                                                                                                                                                                                                                   `json:"file_content_read"`
	RawCommandExposed                                       bool                                                                                                                                                                                                                   `json:"raw_command_exposed"`
	RawExecutableExposed                                    bool                                                                                                                                                                                                                   `json:"raw_executable_exposed"`
	BackendDetailsExposed                                   bool                                                                                                                                                                                                                   `json:"backend_details_exposed"`
	BlockedActions                                          []string                                                                                                                                                                                                               `json:"blocked_actions"`
	NextRequirements                                        []string                                                                                                                                                                                                               `json:"next_requirements"`
	DesktopSafeSummary                                      string                                                                                                                                                                                                                 `json:"desktop_safe_summary"`
}

type ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteAuthorizationEvidenceGateItem struct {
	ID                                                      string `json:"id"`
	ActionKind                                              string `json:"action_kind"`
	SurfaceKind                                             string `json:"surface_kind"`
	EvidenceKind                                            string `json:"evidence_kind"`
	EvidencePresent                                         bool   `json:"evidence_present"`
	CurrentMainlineConsumed                                 bool   `json:"current_mainline_consumed"`
	ProjectionEvidenceAuditConsumed                         bool   `json:"projection_evidence_audit_consumed"`
	ProjectionEvidenceAuditReady                            bool   `json:"projection_evidence_audit_ready"`
	ResultConsumerProjectionRouteAuthorizationEvidenceGated bool   `json:"result_consumer_projection_route_authorization_evidence_gated"`
	RedactedProjectionRouteAuthorizationEvidenceGated       bool   `json:"redacted_projection_route_authorization_evidence_gated"`
	UserVisibleRouteAuthorizationEvidenceGated              bool   `json:"user_visible_route_authorization_evidence_gated"`
	FailureProjectionRouteAuthorizationEvidenceGated        bool   `json:"failure_projection_route_authorization_evidence_gated"`
	KDESafeRedactedStatusOnly                               bool   `json:"kde_safe_redacted_status_only"`
	ReceiptPresent                                          bool   `json:"receipt_present"`
	ReceiptAccepted                                         bool   `json:"receipt_accepted"`
	ReceiptConsumed                                         bool   `json:"receipt_consumed"`
	StorageRootPolicyGrantAuthorized                        bool   `json:"storage_root_policy_grant_authorized"`
	RecordWriterCallAuthorized                              bool   `json:"record_writer_call_authorized"`
	WriterCallable                                          bool   `json:"writer_callable"`
	StorageWriteEnabled                                     bool   `json:"storage_write_enabled"`
	ConsumerEnablementAuthorized                            bool   `json:"consumer_enablement_authorized"`
	LookupRouteEnabled                                      bool   `json:"lookup_route_enabled"`
	RedactedSummaryPersisted                                bool   `json:"redacted_summary_persisted"`
	DryRunResultPersisted                                   bool   `json:"dry_run_result_persisted"`
	RawResultExposed                                        bool   `json:"raw_result_exposed"`
	UserVisible                                             bool   `json:"user_visible"`
	ReviewOnly                                              bool   `json:"review_only"`
	RuntimeOwned                                            bool   `json:"runtime_owned"`
	GoRuntimeBacked                                         bool   `json:"go_runtime_backed"`
	KDEPolicyOwner                                          bool   `json:"kde_policy_owner"`
	SideEffectsDisabled                                     bool   `json:"side_effects_disabled"`
	HostRootModified                                        bool   `json:"host_root_modified"`
	InternalDetailsExposed                                  bool   `json:"internal_details_exposed"`
	RouteAuthorizationGateStatus                            string `json:"route_authorization_gate_status"`
	NextRequirement                                         string `json:"next_requirement"`
}

type ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteAuthorizationEvidenceGateCheck struct {
	ID      string `json:"id"`
	Status  string `json:"status"`
	Summary string `json:"summary"`
}

type ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteAuthorizationEvidenceGateCounts struct {
	Total   int `json:"total"`
	Passed  int `json:"passed"`
	Pending int `json:"pending"`
	Blocked int `json:"blocked"`
}

func NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteAuthorizationEvidenceGatePreview(root string) (ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteAuthorizationEvidenceGatePreview, error) {
	version, err := readProductionDBusGateReviewVersion(root)
	if err != nil {
		return ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteAuthorizationEvidenceGatePreview{}, err
	}
	sources := productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteAuthorizationEvidenceGateSources(root)
	mainlineReady := productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteAuthorizationEvidenceGateMainlineReady(sources.CurrentMainline)
	projectionEvidenceAuditReady := productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteAuthorizationEvidenceGateProjectionEvidenceAuditReady(sources.ProjectionEvidenceAudit)
	items := productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteAuthorizationEvidenceGateItems(mainlineReady, projectionEvidenceAuditReady)
	readyCount := productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteAuthorizationEvidenceGateReadyCount(items)
	ready := mainlineReady && projectionEvidenceAuditReady
	preview := ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteAuthorizationEvidenceGatePreview{
		Version:                 version,
		SchemaVersion:           "xnix.runtime.production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_kde_safe_redacted_status_closed_record_writer_storage_root_authorization_receipt_dry_run_lookup_result_consumer_projection_route_authorization_evidence_gate.v1",
		RequestType:             "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-authorization-evidence-gate-preview",
		GateType:                "receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-authorization-evidence-gate",
		Source:                  "xnix-current-mainline+redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-evidence-audit",
		AuditDecision:           "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-authorization-evidence-gate-blocked",
		CurrentMainlineConsumed: mainlineReady,
		ResultConsumerProjectionEvidenceAuditConsumed:           projectionEvidenceAuditReady,
		ResultConsumerProjectionEvidenceAuditReady:              projectionEvidenceAuditReady,
		RouteAuthorizationEvidenceGateRequired:                  true,
		ResultConsumerProjectionRouteAuthorizationEvidenceGated: ready,
		RouteAuthorizationEvidenceGateReady:                     ready,
		CompatibilityCenterRouteAuthorizationEvidenceGated:      ready,
		RuntimeDiagnosticsRouteAuthorizationEvidenceGated:       ready,
		RedactedProjectionRouteAuthorizationEvidenceGated:       ready,
		UserVisibleRouteAuthorizationEvidenceGated:              ready,
		FailureProjectionRouteAuthorizationEvidenceGated:        ready,
		KDESafeRedactedStatusOnly:                               ready,
		EvidenceItemCount:                                       len(items),
		ReadyEvidenceItemCount:                                  readyCount,
		MissingEvidenceItemCount:                                len(items) - readyCount,
		CompatibilityCenterEvidenceItemCount:                    productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteAuthorizationEvidenceGateSurfaceCount(items, "compatibility-center"),
		RuntimeDiagnosticsEvidenceItemCount:                     productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteAuthorizationEvidenceGateSurfaceCount(items, "runtime-diagnostics"),
		RawExposedEvidenceItemCount:                             0,
		SideEffectEvidenceItemCount:                             0,
		EvidenceItems:                                           items,
		EvidenceItemIDs:                                         productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteAuthorizationEvidenceGateItemIDs(items),
		RuntimeOwned:                                            true,
		GoRuntimeBacked:                                         true,
		KDEPolicyOwner:                                          false,
		BlockedActions: []string{
			"receipt-consumption",
			"storage-root-policy-grant",
			"record-writer-call",
			"lookup-route-enable",
			"consumer-enable",
			"dry-run-dispatch",
			"result-persistence",
			"raw-result-exposure",
			"production-ownership",
			"backend-launch",
			"host-mutation",
		},
		NextRequirements: []string{
			"Add a route authorization evidence gate before lookup routes can be enabled.",
			"Keep route authorization evidence gate output redacted, review-only, and side-effect free.",
		},
		DesktopSafeSummary: "KDE-safe redacted dry-run lookup result consumer projection route authorization evidence is gated for Compatibility Center and Runtime diagnostics while the lookup route itself and every writer, dispatch, persistence, production, backend, and host side effect remain disabled.",
	}
	if ready {
		preview.AuditDecision = "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-authorization-evidence-gate-ready-routes-disabled"
	}
	preview.Checks = productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteAuthorizationEvidenceGateChecks(preview)
	preview.CheckIDs = productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteAuthorizationEvidenceGateCheckIDs(preview.Checks)
	preview.Counts = productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteAuthorizationEvidenceGateCountsFor(preview.Checks)
	if err := validateNoBackendTerms(preview, "KDE-safe redacted status dry-run lookup result consumer projection route authorization evidence gate preview"); err != nil {
		return ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteAuthorizationEvidenceGatePreview{}, err
	}
	return preview, nil
}

type productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteAuthorizationEvidenceGateSourceSet struct {
	CurrentMainline         string
	ProjectionEvidenceAudit string
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteAuthorizationEvidenceGateSources(root string) productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteAuthorizationEvidenceGateSourceSet {
	return productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteAuthorizationEvidenceGateSourceSet{
		CurrentMainline: productionAuthorizationReadSources(root, []string{"docs/xnix-current-mainline.md"}),
		ProjectionEvidenceAudit: productionAuthorizationReadSources(root, []string{
			"internal/runtime/owner/production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_kde_safe_redacted_status_closed_record_writer_storage_root_authorization_receipt_dry_run_lookup_result_consumer_projection_evidence_audit_preview.go",
			"internal/runtime/owner/production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_kde_safe_redacted_status_closed_record_writer_storage_root_authorization_receipt_dry_run_lookup_result_consumer_projection_evidence_audit_preview_test.go",
		}),
	}
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteAuthorizationEvidenceGateMainlineReady(source string) bool {
	return productionAuthorizationHasAll(source, []string{
		"result consumer projection route authorization evidence gate preview",
		"storage-root authorization receipt dry-run lookup result consumer projection evidence audit preview",
		"without accepting or consuming receipts",
	})
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteAuthorizationEvidenceGateProjectionEvidenceAuditReady(source string) bool {
	return productionAuthorizationHasAll(source, []string{
		"xnix.runtime.production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_kde_safe_redacted_status_closed_record_writer_storage_root_authorization_receipt_dry_run_lookup_result_consumer_projection_evidence_audit.v1",
		"production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-evidence-audit-ready-writes-disabled",
		"ResultConsumerProjectionEvidenceAudited",
		"CompatibilityCenterEvidenceAudited",
		"RuntimeDiagnosticsEvidenceAudited",
	})
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteAuthorizationEvidenceGateItems(mainlineReady, projectionEvidenceAuditReady bool) []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteAuthorizationEvidenceGateItem {
	actionKinds := []string{"review", "renew", "open-compatibility-center", "dismiss", "support-info"}
	items := make([]ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteAuthorizationEvidenceGateItem, 0, len(actionKinds)*2)
	for _, actionKind := range actionKinds {
		for _, surfaceKind := range []string{"compatibility-center", "runtime-diagnostics"} {
			ready := mainlineReady && projectionEvidenceAuditReady
			status := "missing-redacted-status-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-authorization-evidence-gate"
			if ready {
				status = "redacted-status-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-authorization-evidence-gated-routes-disabled"
			}
			items = append(items, ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteAuthorizationEvidenceGateItem{
				ID:                              "result-consumer-projection-evidence-" + surfaceKind + "-" + actionKind,
				ActionKind:                      actionKind,
				SurfaceKind:                     surfaceKind,
				EvidenceKind:                    "kde-safe-redacted-lookup-result-projection-route-authorization-evidence",
				EvidencePresent:                 ready,
				CurrentMainlineConsumed:         mainlineReady,
				ProjectionEvidenceAuditConsumed: projectionEvidenceAuditReady,
				ProjectionEvidenceAuditReady:    projectionEvidenceAuditReady,
				ResultConsumerProjectionRouteAuthorizationEvidenceGated: ready,
				RedactedProjectionRouteAuthorizationEvidenceGated:       ready,
				UserVisibleRouteAuthorizationEvidenceGated:              ready,
				FailureProjectionRouteAuthorizationEvidenceGated:        ready,
				KDESafeRedactedStatusOnly:                               ready,
				UserVisible:                                             ready,
				ReviewOnly:                                              ready,
				RuntimeOwned:                                            true,
				GoRuntimeBacked:                                         true,
				KDEPolicyOwner:                                          false,
				SideEffectsDisabled:                                     true,
				RouteAuthorizationGateStatus:                            status,
				NextRequirement:                                         "Require a separate route enablement gate before enabling lookup routes, consumers, persistence, dry-run dispatch, or production side effects.",
			})
		}
	}
	return items
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteAuthorizationEvidenceGateReadyCount(items []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteAuthorizationEvidenceGateItem) int {
	count := 0
	for _, item := range items {
		if item.EvidencePresent && item.ResultConsumerProjectionRouteAuthorizationEvidenceGated && item.RedactedProjectionRouteAuthorizationEvidenceGated && item.UserVisibleRouteAuthorizationEvidenceGated && item.FailureProjectionRouteAuthorizationEvidenceGated && item.KDESafeRedactedStatusOnly && !item.RawResultExposed && !item.StorageWriteEnabled && !item.HostRootModified && !item.InternalDetailsExposed {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteAuthorizationEvidenceGateSurfaceCount(items []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteAuthorizationEvidenceGateItem, surfaceKind string) int {
	count := 0
	for _, item := range items {
		if item.SurfaceKind == surfaceKind {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteAuthorizationEvidenceGateItemIDs(items []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteAuthorizationEvidenceGateItem) []string {
	ids := make([]string, 0, len(items))
	for _, item := range items {
		ids = append(ids, item.ID)
	}
	return ids
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteAuthorizationEvidenceGateChecks(preview ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteAuthorizationEvidenceGatePreview) []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteAuthorizationEvidenceGateCheck {
	return []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteAuthorizationEvidenceGateCheck{
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteAuthorizationEvidenceGateCheck("current-mainline-consumed", preview.CurrentMainlineConsumed, "Xnix current mainline names the result consumer projection route authorization evidence gate preview."),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteAuthorizationEvidenceGateCheck("result-consumer-projection-evidence-audit-consumed", preview.ResultConsumerProjectionEvidenceAuditConsumed, "The gate consumes the dry-run lookup result consumer projection evidence audit."),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteAuthorizationEvidenceGateCheck("result-consumer-projection-route-authorization-evidence-gated", preview.ResultConsumerProjectionRouteAuthorizationEvidenceGated, "Projection route authorization identity, user-visible redaction, and failure evidence are gated."),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteAuthorizationEvidenceGateCheck("compatibility-center-and-runtime-route-authorization-evidence-gated", preview.CompatibilityCenterRouteAuthorizationEvidenceGated && preview.RuntimeDiagnosticsRouteAuthorizationEvidenceGated, "Compatibility Center and Runtime diagnostics route authorization evidence boundaries are gated."),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteAuthorizationEvidenceGateCheck("ten-route-authorization-evidence-items-ready-routes-disabled", preview.EvidenceItemCount == 10 && preview.ReadyEvidenceItemCount == 10 && preview.MissingEvidenceItemCount == 0, "Ten KDE-safe route authorization evidence items are ready while lookup routes remain disabled."),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteAuthorizationEvidenceGateCheck("route-authorization-evidence-and-writes-disabled", preview.StorageWriteEnabled == false && preview.StatusPersistenceWriteEnabled == false && preview.DryRunResultPersisted == false && preview.RawResultExposed == false, "Route authorization evidence gate does not enable storage, status persistence, dry-run result persistence, or raw result exposure."),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteAuthorizationEvidenceGateCheck("consumer-routes-dispatch-and-support-disabled", preview.ConsumerEnablementAuthorized == false && preview.KDEConsumerEnabled == false && preview.RuntimeConsumerEnabled == false && preview.LookupRouteEnabled == false && preview.DispatchDryRunExecuted == false && preview.SupportBundleExported == false && preview.SupportCaseCreated == false, "Consumer enablement, lookup routes, dispatch, and support side effects remain disabled."),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteAuthorizationEvidenceGateCheck("production-and-host-boundary-closed", preview.ProductionBusClaimed == false && preview.WriteMethodsEnabled == false && preview.RuntimeWritesEnabled == false && preview.BackendLaunchEnabled == false && preview.NetworkRequired == false && preview.HostRootModified == false && preview.PrivilegedContainerRequired == false, "Production, backend, network, privileged container, and host mutation gates remain closed."),
	}
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteAuthorizationEvidenceGateCheck(id string, ok bool, summary string) ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteAuthorizationEvidenceGateCheck {
	status := "blocked"
	if ok {
		status = "pass"
	}
	return ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteAuthorizationEvidenceGateCheck{ID: id, Status: status, Summary: summary}
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteAuthorizationEvidenceGateCheckIDs(checks []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteAuthorizationEvidenceGateCheck) []string {
	ids := make([]string, 0, len(checks))
	for _, check := range checks {
		ids = append(ids, check.ID)
	}
	return ids
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteAuthorizationEvidenceGateCountsFor(checks []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteAuthorizationEvidenceGateCheck) ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteAuthorizationEvidenceGateCounts {
	counts := ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteAuthorizationEvidenceGateCounts{Total: len(checks)}
	for _, check := range checks {
		if check.Status == "pass" {
			counts.Passed++
		} else {
			counts.Blocked++
		}
	}
	return counts
}
