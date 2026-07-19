package owner

type ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementAcceptedReceiptGateAuditPreview struct {
	Version                                       string                                                                                                                                                                                                                              `json:"version"`
	SchemaVersion                                 string                                                                                                                                                                                                                              `json:"schema_version"`
	RequestType                                   string                                                                                                                                                                                                                              `json:"request_type"`
	AuditType                                     string                                                                                                                                                                                                                              `json:"audit_type"`
	Source                                        string                                                                                                                                                                                                                              `json:"source"`
	AuditDecision                                 string                                                                                                                                                                                                                              `json:"audit_decision"`
	CurrentMainlineConsumed                       bool                                                                                                                                                                                                                                `json:"current_mainline_consumed"`
	RouteEnablementReceiptAcceptanceAuditConsumed bool                                                                                                                                                                                                                                `json:"route_enablement_receipt_acceptance_audit_consumed"`
	RouteEnablementReceiptAcceptanceReady         bool                                                                                                                                                                                                                                `json:"route_enablement_receipt_acceptance_ready"`
	AcceptedReceiptGateRequired                   bool                                                                                                                                                                                                                                `json:"accepted_receipt_gate_required"`
	AcceptedReceiptGateModeled                    bool                                                                                                                                                                                                                                `json:"accepted_receipt_gate_modeled"`
	AcceptedReceiptGateReady                      bool                                                                                                                                                                                                                                `json:"accepted_receipt_gate_ready"`
	RouteEnablementAcceptedReceiptBoundaryReady   bool                                                                                                                                                                                                                                `json:"route_enablement_accepted_receipt_boundary_ready"`
	OpaqueRouteEnablementReceipt                  bool                                                                                                                                                                                                                                `json:"opaque_route_enablement_receipt"`
	KDESafeRedactedStatusOnly                     bool                                                                                                                                                                                                                                `json:"kde_safe_redacted_status_only"`
	CompatibilityCenterAcceptedReceiptModeled     bool                                                                                                                                                                                                                                `json:"compatibility_center_accepted_receipt_modeled"`
	RuntimeDiagnosticsAcceptedReceiptModeled      bool                                                                                                                                                                                                                                `json:"runtime_diagnostics_accepted_receipt_modeled"`
	ReceiptPresent                                bool                                                                                                                                                                                                                                `json:"receipt_present"`
	ReceiptAccepted                               bool                                                                                                                                                                                                                                `json:"receipt_accepted"`
	ReceiptConsumed                               bool                                                                                                                                                                                                                                `json:"receipt_consumed"`
	AcceptanceAuthorized                          bool                                                                                                                                                                                                                                `json:"acceptance_authorized"`
	RouteEnablementAccepted                       bool                                                                                                                                                                                                                                `json:"route_enablement_accepted"`
	LookupRouteEnablementGranted                  bool                                                                                                                                                                                                                                `json:"lookup_route_enablement_granted"`
	LookupRouteEnabled                            bool                                                                                                                                                                                                                                `json:"lookup_route_enabled"`
	StorageRootPolicyGrantAuthorized              bool                                                                                                                                                                                                                                `json:"storage_root_policy_grant_authorized"`
	RecordWriterCallAuthorized                    bool                                                                                                                                                                                                                                `json:"record_writer_call_authorized"`
	WriterCallable                                bool                                                                                                                                                                                                                                `json:"writer_callable"`
	StorageRootResolved                           bool                                                                                                                                                                                                                                `json:"storage_root_resolved"`
	StorageWriteEnabled                           bool                                                                                                                                                                                                                                `json:"storage_write_enabled"`
	StatusPersistenceWriteEnabled                 bool                                                                                                                                                                                                                                `json:"status_persistence_write_enabled"`
	GateItemCount                                 int                                                                                                                                                                                                                                 `json:"gate_item_count"`
	RequiredGateItemCount                         int                                                                                                                                                                                                                                 `json:"required_gate_item_count"`
	ReadyGateItemCount                            int                                                                                                                                                                                                                                 `json:"ready_gate_item_count"`
	MissingGateItemCount                          int                                                                                                                                                                                                                                 `json:"missing_gate_item_count"`
	AcceptedReceiptItemCount                      int                                                                                                                                                                                                                                 `json:"accepted_receipt_item_count"`
	EnabledRouteItemCount                         int                                                                                                                                                                                                                                 `json:"enabled_route_item_count"`
	PersistedGateItemCount                        int                                                                                                                                                                                                                                 `json:"persisted_gate_item_count"`
	RawExposedGateItemCount                       int                                                                                                                                                                                                                                 `json:"raw_exposed_gate_item_count"`
	SideEffectGateItemCount                       int                                                                                                                                                                                                                                 `json:"side_effect_gate_item_count"`
	CompatibilityCenterGateItemCount              int                                                                                                                                                                                                                                 `json:"compatibility_center_gate_item_count"`
	RuntimeDiagnosticsGateItemCount               int                                                                                                                                                                                                                                 `json:"runtime_diagnostics_gate_item_count"`
	GateItems                                     []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementAcceptedReceiptGateAuditItem      `json:"gate_items"`
	GateItemIDs                                   []string                                                                                                                                                                                                                            `json:"gate_item_ids"`
	Checks                                        []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementAcceptedReceiptGateAuditCheck     `json:"checks"`
	CheckIDs                                      []string                                                                                                                                                                                                                            `json:"check_ids"`
	Counts                                        ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementAcceptedReceiptGateAuditCheckCounts `json:"counts"`
	ConsumerDryRunLookupResultBoundaryAuthorized  bool                                                                                                                                                                                                                                `json:"consumer_dry_run_lookup_result_boundary_authorized"`
	ConsumerEnablementAuthorized                  bool                                                                                                                                                                                                                                `json:"consumer_enablement_authorized"`
	KDEConsumerEnabled                            bool                                                                                                                                                                                                                                `json:"kde_consumer_enabled"`
	RuntimeConsumerEnabled                        bool                                                                                                                                                                                                                                `json:"runtime_consumer_enabled"`
	LookupRouteAuthorized                         bool                                                                                                                                                                                                                                `json:"lookup_route_authorized"`
	LookupRouteEnablementAuthorized               bool                                                                                                                                                                                                                                `json:"lookup_route_enablement_authorized"`
	OpaqueLookupEnabled                           bool                                                                                                                                                                                                                                `json:"opaque_lookup_enabled"`
	RedactedSummaryPersisted                      bool                                                                                                                                                                                                                                `json:"redacted_summary_persisted"`
	KDEStatusPersisted                            bool                                                                                                                                                                                                                                `json:"kde_status_persisted"`
	RuntimeDiagnosticsPersisted                   bool                                                                                                                                                                                                                                `json:"runtime_diagnostics_persisted"`
	DryRunResultPersisted                         bool                                                                                                                                                                                                                                `json:"dry_run_result_persisted"`
	RawResultExposed                              bool                                                                                                                                                                                                                                `json:"raw_result_exposed"`
	DispatchDryRunExecuted                        bool                                                                                                                                                                                                                                `json:"dispatch_dry_run_executed"`
	RequestObjectCreationEnabled                  bool                                                                                                                                                                                                                                `json:"request_object_creation_enabled"`
	RequestObjectDispatchEnabled                  bool                                                                                                                                                                                                                                `json:"request_object_dispatch_enabled"`
	PortalRequestCreated                          bool                                                                                                                                                                                                                                `json:"portal_request_created"`
	NotificationActionEnabled                     bool                                                                                                                                                                                                                                `json:"notification_action_enabled"`
	CompatibilityCenterOpened                     bool                                                                                                                                                                                                                                `json:"compatibility_center_opened"`
	SupportBundleExported                         bool                                                                                                                                                                                                                                `json:"support_bundle_exported"`
	SupportCaseCreated                            bool                                                                                                                                                                                                                                `json:"support_case_created"`
	RuntimeOwned                                  bool                                                                                                                                                                                                                                `json:"runtime_owned"`
	GoRuntimeBacked                               bool                                                                                                                                                                                                                                `json:"go_runtime_backed"`
	KDEPolicyOwner                                bool                                                                                                                                                                                                                                `json:"kde_policy_owner"`
	ProductionReadiness                           bool                                                                                                                                                                                                                                `json:"production_readiness"`
	ProductionOwnershipReady                      bool                                                                                                                                                                                                                                `json:"production_ownership_ready"`
	SystemServiceStarted                          bool                                                                                                                                                                                                                                `json:"system_service_started"`
	SessionBusClaimed                             bool                                                                                                                                                                                                                                `json:"session_bus_claimed"`
	ProductionBusClaimed                          bool                                                                                                                                                                                                                                `json:"production_bus_claimed"`
	WriteMethodsEnabled                           bool                                                                                                                                                                                                                                `json:"write_methods_enabled"`
	RuntimeWritesEnabled                          bool                                                                                                                                                                                                                                `json:"runtime_writes_enabled"`
	DesktopFilesWritten                           bool                                                                                                                                                                                                                                `json:"desktop_files_written"`
	KDEConfigurationWritten                       bool                                                                                                                                                                                                                                `json:"kde_configuration_written"`
	PortalCallExecuted                            bool                                                                                                                                                                                                                                `json:"portal_call_executed"`
	AdapterInvocationEnabled                      bool                                                                                                                                                                                                                                `json:"adapter_invocation_enabled"`
	BackendLaunchEnabled                          bool                                                                                                                                                                                                                                `json:"backend_launch_enabled"`
	BackendProcessStarted                         bool                                                                                                                                                                                                                                `json:"backend_process_started"`
	NetworkRequired                               bool                                                                                                                                                                                                                                `json:"network_required"`
	HostRootModified                              bool                                                                                                                                                                                                                                `json:"host_root_modified"`
	PrivilegedContainerRequired                   bool                                                                                                                                                                                                                                `json:"privileged_container_required"`
	StateRootPathExposed                          bool                                                                                                                                                                                                                                `json:"state_root_path_exposed"`
	FilePathsExposed                              bool                                                                                                                                                                                                                                `json:"file_paths_exposed"`
	FileContentRead                               bool                                                                                                                                                                                                                                `json:"file_content_read"`
	RawCommandExposed                             bool                                                                                                                                                                                                                                `json:"raw_command_exposed"`
	RawExecutableExposed                          bool                                                                                                                                                                                                                                `json:"raw_executable_exposed"`
	BackendDetailsExposed                         bool                                                                                                                                                                                                                                `json:"backend_details_exposed"`
	BlockedActions                                []string                                                                                                                                                                                                                            `json:"blocked_actions"`
	NextRequirements                              []string                                                                                                                                                                                                                            `json:"next_requirements"`
	DesktopSafeSummary                            string                                                                                                                                                                                                                              `json:"desktop_safe_summary"`
}

type ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementAcceptedReceiptGateAuditItem struct {
	ID                                            string `json:"id"`
	ActionKind                                    string `json:"action_kind"`
	SurfaceKind                                   string `json:"surface_kind"`
	StatusConsumerKind                            string `json:"status_consumer_kind"`
	OpaqueReceiptID                               string `json:"opaque_receipt_id"`
	EvidencePresent                               bool   `json:"evidence_present"`
	CurrentMainlineConsumed                       bool   `json:"current_mainline_consumed"`
	RouteEnablementReceiptAcceptanceAuditConsumed bool   `json:"route_enablement_receipt_acceptance_audit_consumed"`
	RouteEnablementReceiptAcceptanceReady         bool   `json:"route_enablement_receipt_acceptance_ready"`
	AcceptedReceiptGateModeled                    bool   `json:"accepted_receipt_gate_modeled"`
	RouteEnablementAcceptedReceiptBoundaryReady   bool   `json:"route_enablement_accepted_receipt_boundary_ready"`
	OpaqueRouteEnablementReceipt                  bool   `json:"opaque_route_enablement_receipt"`
	KDESafeRedactedStatusOnly                     bool   `json:"kde_safe_redacted_status_only"`
	ReceiptPresent                                bool   `json:"receipt_present"`
	ReceiptAccepted                               bool   `json:"receipt_accepted"`
	ReceiptConsumed                               bool   `json:"receipt_consumed"`
	AcceptanceAuthorized                          bool   `json:"acceptance_authorized"`
	RouteEnablementAccepted                       bool   `json:"route_enablement_accepted"`
	LookupRouteEnablementGranted                  bool   `json:"lookup_route_enablement_granted"`
	LookupRouteEnabled                            bool   `json:"lookup_route_enabled"`
	StorageWriteEnabled                           bool   `json:"storage_write_enabled"`
	RedactedSummaryPersisted                      bool   `json:"redacted_summary_persisted"`
	KDEStatusPersisted                            bool   `json:"kde_status_persisted"`
	RuntimeDiagnosticsPersisted                   bool   `json:"runtime_diagnostics_persisted"`
	RawResultExposed                              bool   `json:"raw_result_exposed"`
	UserVisible                                   bool   `json:"user_visible"`
	ReviewOnly                                    bool   `json:"review_only"`
	RuntimeOwned                                  bool   `json:"runtime_owned"`
	GoRuntimeBacked                               bool   `json:"go_runtime_backed"`
	KDEPolicyOwner                                bool   `json:"kde_policy_owner"`
	SideEffectsDisabled                           bool   `json:"side_effects_disabled"`
	HostRootModified                              bool   `json:"host_root_modified"`
	InternalDetailsExposed                        bool   `json:"internal_details_exposed"`
	AcceptedReceiptGateStatus                     string `json:"accepted_receipt_gate_status"`
	NextRequirement                               string `json:"next_requirement"`
}

type ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementAcceptedReceiptGateAuditCheck struct {
	ID      string `json:"id"`
	Status  string `json:"status"`
	Summary string `json:"summary"`
}

type ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementAcceptedReceiptGateAuditCheckCounts struct {
	Total   int `json:"total"`
	Passed  int `json:"passed"`
	Pending int `json:"pending"`
	Blocked int `json:"blocked"`
}

type productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementAcceptedReceiptGateAuditSourceSet struct {
	CurrentMainline                string
	RouteEnablementAcceptanceAudit string
}

func NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementAcceptedReceiptGateAuditPreview(root string) (ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementAcceptedReceiptGateAuditPreview, error) {
	version, err := readProductionDBusGateReviewVersion(root)
	if err != nil {
		return ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementAcceptedReceiptGateAuditPreview{}, err
	}
	sources := productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementAcceptedReceiptGateAuditSources(root)
	mainlineReady := productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementAcceptedReceiptGateAuditMainlineReady(sources.CurrentMainline)
	acceptanceReady := productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementAcceptedReceiptGateAuditAcceptanceReady(sources.RouteEnablementAcceptanceAudit)
	items := productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementAcceptedReceiptGateAuditItems(mainlineReady, acceptanceReady)
	readyItemCount := productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementAcceptedReceiptGateAuditReadyCount(items)
	ready := mainlineReady && acceptanceReady && readyItemCount == len(items)
	preview := ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementAcceptedReceiptGateAuditPreview{
		Version:                 version,
		SchemaVersion:           "xnix.runtime.production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_kde_safe_redacted_status_closed_record_writer_storage_root_authorization_receipt_dry_run_lookup_result_consumer_projection_route_enablement_accepted_receipt_gate_audit.v1",
		RequestType:             "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-accepted-receipt-gate-audit-preview",
		AuditType:               "receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-accepted-receipt-gate-audit",
		Source:                  "xnix-current-mainline+redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-receipt-acceptance-audit",
		AuditDecision:           "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-accepted-receipt-gate-audit-blocked",
		CurrentMainlineConsumed: mainlineReady,
		RouteEnablementReceiptAcceptanceAuditConsumed: acceptanceReady,
		RouteEnablementReceiptAcceptanceReady:         acceptanceReady,
		AcceptedReceiptGateRequired:                   true,
		AcceptedReceiptGateModeled:                    true,
		AcceptedReceiptGateReady:                      ready,
		RouteEnablementAcceptedReceiptBoundaryReady:   ready,
		OpaqueRouteEnablementReceipt:                  acceptanceReady,
		KDESafeRedactedStatusOnly:                     ready,
		CompatibilityCenterAcceptedReceiptModeled:     ready,
		RuntimeDiagnosticsAcceptedReceiptModeled:      ready,
		GateItemCount:                                 len(items),
		RequiredGateItemCount:                         len(items),
		ReadyGateItemCount:                            readyItemCount,
		MissingGateItemCount:                          len(items) - readyItemCount,
		CompatibilityCenterGateItemCount:              productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementAcceptedReceiptGateAuditSurfaceCount(items, "compatibility-center"),
		RuntimeDiagnosticsGateItemCount:               productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementAcceptedReceiptGateAuditSurfaceCount(items, "runtime-diagnostics"),
		GateItems:                                     items,
		GateItemIDs:                                   productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementAcceptedReceiptGateAuditItemIDs(items),
		RuntimeOwned:                                  true,
		GoRuntimeBacked:                               true,
		KDEPolicyOwner:                                false,
		BlockedActions: []string{
			"receipt-present",
			"receipt-accept",
			"receipt-consume",
			"route-enablement-accept",
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
			"Add a route enablement lookup route grant audit before any lookup route can be enabled.",
			"Keep accepted receipt gates modeled, redacted, and disabled until production ownership is explicitly authorized.",
		},
		DesktopSafeSummary: "KDE-safe redacted route enablement accepted receipt gates are audited for Compatibility Center and Runtime diagnostics while receipt acceptance, lookup routes, consumers, dispatch, persistence, production, backend, and host side effects remain disabled.",
	}
	if ready {
		preview.AuditDecision = "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-accepted-receipt-gate-audit-ready-gate-disabled-routes-disabled"
	}
	preview.Checks = productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementAcceptedReceiptGateAuditChecks(preview)
	preview.CheckIDs = productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementAcceptedReceiptGateAuditCheckIDs(preview.Checks)
	preview.Counts = productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementAcceptedReceiptGateAuditCountsFor(preview.Checks)
	if err := validateNoBackendTerms(preview, "KDE-safe redacted status dry-run lookup result consumer projection route enablement accepted receipt gate audit preview"); err != nil {
		return ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementAcceptedReceiptGateAuditPreview{}, err
	}
	return preview, nil
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementAcceptedReceiptGateAuditSources(root string) productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementAcceptedReceiptGateAuditSourceSet {
	return productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementAcceptedReceiptGateAuditSourceSet{
		CurrentMainline: productionAuthorizationReadSources(root, []string{"docs/xnix-current-mainline.md"}),
		RouteEnablementAcceptanceAudit: productionAuthorizationReadSources(root, []string{
			"internal/runtime/owner/route_enablement_receipt_acceptance_audit_preview.go",
			"internal/runtime/owner/route_enablement_receipt_acceptance_audit_preview_test.go",
		}),
	}
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementAcceptedReceiptGateAuditMainlineReady(source string) bool {
	return productionAuthorizationHasAll(source, []string{
		"route enablement accepted receipt gate audit preview",
		"route enablement receipt acceptance audit preview",
		"without accepting or consuming receipts",
	})
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementAcceptedReceiptGateAuditAcceptanceReady(source string) bool {
	return productionAuthorizationHasAll(source, []string{
		"xnix.runtime.production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_kde_safe_redacted_status_closed_record_writer_storage_root_authorization_receipt_dry_run_lookup_result_consumer_projection_route_enablement_receipt_acceptance_audit.v1",
		"production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-receipt-acceptance-audit-ready-acceptance-disabled-routes-disabled",
		"AcceptanceAuthorizationReady",
		"RouteEnablementAcceptanceBoundaryReady",
	})
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementAcceptedReceiptGateAuditItems(mainlineReady, acceptanceReady bool) []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementAcceptedReceiptGateAuditItem {
	actionKinds := []string{"review", "renew", "open-compatibility-center", "dismiss", "support-info"}
	items := make([]ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementAcceptedReceiptGateAuditItem, 0, len(actionKinds)*2)
	for _, actionKind := range actionKinds {
		for _, surfaceKind := range []string{"compatibility-center", "runtime-diagnostics"} {
			ready := mainlineReady && acceptanceReady
			status := "missing-redacted-status-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-accepted-receipt-gate-evidence"
			if ready {
				status = "redacted-status-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-accepted-receipt-gate-modeled-gate-disabled-routes-disabled"
			}
			items = append(items, ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementAcceptedReceiptGateAuditItem{
				ID:                      actionKind + "-route-enablement-accepted-receipt-gate-" + surfaceKind,
				ActionKind:              actionKind,
				SurfaceKind:             surfaceKind,
				StatusConsumerKind:      "kde-safe-redacted-route-enablement-accepted-receipt-gate",
				OpaqueReceiptID:         "opaque-route-enablement-accepted-receipt-" + surfaceKind + "-" + actionKind,
				EvidencePresent:         ready,
				CurrentMainlineConsumed: mainlineReady,
				RouteEnablementReceiptAcceptanceAuditConsumed: acceptanceReady,
				RouteEnablementReceiptAcceptanceReady:         acceptanceReady,
				AcceptedReceiptGateModeled:                    ready,
				RouteEnablementAcceptedReceiptBoundaryReady:   ready,
				OpaqueRouteEnablementReceipt:                  acceptanceReady,
				KDESafeRedactedStatusOnly:                     ready,
				UserVisible:                                   ready,
				ReviewOnly:                                    ready,
				RuntimeOwned:                                  true,
				GoRuntimeBacked:                               true,
				KDEPolicyOwner:                                false,
				SideEffectsDisabled:                           true,
				AcceptedReceiptGateStatus:                     status,
				NextRequirement:                               "Require a separate lookup route grant audit before route enablement can be granted.",
			})
		}
	}
	return items
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementAcceptedReceiptGateAuditReadyCount(items []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementAcceptedReceiptGateAuditItem) int {
	count := 0
	for _, item := range items {
		if item.EvidencePresent && item.AcceptedReceiptGateModeled && item.RouteEnablementAcceptedReceiptBoundaryReady && item.KDESafeRedactedStatusOnly && !item.ReceiptPresent && !item.ReceiptAccepted && !item.ReceiptConsumed && !item.AcceptanceAuthorized && !item.RouteEnablementAccepted && !item.LookupRouteEnablementGranted && !item.LookupRouteEnabled && !item.StorageWriteEnabled && !item.RawResultExposed && !item.HostRootModified && !item.InternalDetailsExposed {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementAcceptedReceiptGateAuditSurfaceCount(items []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementAcceptedReceiptGateAuditItem, surfaceKind string) int {
	count := 0
	for _, item := range items {
		if item.SurfaceKind == surfaceKind {
			count++
		}
	}
	return count
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementAcceptedReceiptGateAuditItemIDs(items []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementAcceptedReceiptGateAuditItem) []string {
	ids := make([]string, 0, len(items))
	for _, item := range items {
		ids = append(ids, item.ID)
	}
	return ids
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementAcceptedReceiptGateAuditChecks(preview ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementAcceptedReceiptGateAuditPreview) []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementAcceptedReceiptGateAuditCheck {
	return []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementAcceptedReceiptGateAuditCheck{
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementAcceptedReceiptGateAuditCheck("current-mainline-consumed", preview.CurrentMainlineConsumed, "Xnix current mainline names the result consumer projection route enablement accepted receipt gate audit preview."),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementAcceptedReceiptGateAuditCheck("route-enablement-receipt-acceptance-audit-consumed", preview.RouteEnablementReceiptAcceptanceAuditConsumed && preview.RouteEnablementReceiptAcceptanceReady, "The accepted receipt gate audit consumes the ready route enablement receipt acceptance audit."),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementAcceptedReceiptGateAuditCheck("accepted-receipt-gate-modeled", preview.AcceptedReceiptGateRequired && preview.AcceptedReceiptGateModeled, "Accepted receipt gate is modeled without accepting a receipt."),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementAcceptedReceiptGateAuditCheck("compatibility-center-and-runtime-accepted-receipt-gates-modeled", preview.CompatibilityCenterAcceptedReceiptModeled && preview.RuntimeDiagnosticsAcceptedReceiptModeled, "Compatibility Center and Runtime diagnostics accepted receipt gate boundaries are modeled."),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementAcceptedReceiptGateAuditCheck("ten-gate-items-ready-gate-disabled", preview.GateItemCount == 10 && preview.ReadyGateItemCount == 10 && preview.MissingGateItemCount == 0, "Ten KDE-safe accepted receipt gate items are ready while route gates remain disabled."),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementAcceptedReceiptGateAuditCheck("accepted-receipt-routes-and-persistence-disabled", !preview.ReceiptPresent && !preview.ReceiptAccepted && !preview.ReceiptConsumed && !preview.AcceptanceAuthorized && !preview.RouteEnablementAccepted && !preview.LookupRouteEnablementGranted && !preview.LookupRouteEnabled && !preview.StorageWriteEnabled && !preview.StatusPersistenceWriteEnabled && !preview.RawResultExposed, "Accepted receipt gates, route enablement grants, lookup routes, storage, persistence, and raw result exposure remain disabled."),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementAcceptedReceiptGateAuditCheck("consumer-routes-dispatch-and-support-disabled", !preview.ConsumerEnablementAuthorized && !preview.KDEConsumerEnabled && !preview.RuntimeConsumerEnabled && !preview.LookupRouteEnabled && !preview.DispatchDryRunExecuted && !preview.SupportBundleExported && !preview.SupportCaseCreated, "Consumer enablement, lookup routes, dispatch, and support side effects remain disabled."),
		productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementAcceptedReceiptGateAuditCheck("production-and-host-boundary-closed", !preview.ProductionBusClaimed && !preview.WriteMethodsEnabled && !preview.RuntimeWritesEnabled && !preview.BackendLaunchEnabled && !preview.NetworkRequired && !preview.HostRootModified && !preview.PrivilegedContainerRequired, "Production, backend, network, privileged container, and host mutation gates remain closed."),
	}
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementAcceptedReceiptGateAuditCheck(id string, ok bool, summary string) ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementAcceptedReceiptGateAuditCheck {
	status := "blocked"
	if ok {
		status = "pass"
	}
	return ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementAcceptedReceiptGateAuditCheck{ID: id, Status: status, Summary: summary}
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementAcceptedReceiptGateAuditCheckIDs(checks []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementAcceptedReceiptGateAuditCheck) []string {
	ids := make([]string, 0, len(checks))
	for _, check := range checks {
		ids = append(ids, check.ID)
	}
	return ids
}

func productionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementAcceptedReceiptGateAuditCountsFor(checks []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementAcceptedReceiptGateAuditCheck) ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementAcceptedReceiptGateAuditCheckCounts {
	counts := ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementAcceptedReceiptGateAuditCheckCounts{Total: len(checks)}
	for _, check := range checks {
		if check.Status == "pass" {
			counts.Passed++
		} else {
			counts.Blocked++
		}
	}
	return counts
}
