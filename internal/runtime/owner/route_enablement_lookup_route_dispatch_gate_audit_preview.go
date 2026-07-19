package owner

type ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchGateAuditPreview struct {
	Version                                             string                                                                                                                                                                                                                                  `json:"version"`
	SchemaVersion                                       string                                                                                                                                                                                                                                  `json:"schema_version"`
	RequestType                                         string                                                                                                                                                                                                                                  `json:"request_type"`
	AuditType                                           string                                                                                                                                                                                                                                  `json:"audit_type"`
	Source                                              string                                                                                                                                                                                                                                  `json:"source"`
	AuditDecision                                       string                                                                                                                                                                                                                                  `json:"audit_decision"`
	CurrentMainlineConsumed                             bool                                                                                                                                                                                                                                    `json:"current_mainline_consumed"`
	RouteEnablementLookupRouteEnablementConsumed        bool                                                                                                                                                                                                                                    `json:"route_enablement_lookup_route_enablement_consumed"`
	RouteEnablementLookupRouteEnablementReady           bool                                                                                                                                                                                                                                    `json:"route_enablement_lookup_route_enablement_ready"`
	LookupRouteDispatchGateRequired                     bool                                                                                                                                                                                                                                    `json:"lookup_route_dispatch_gate_required"`
	LookupRouteDispatchGateModeled                      bool                                                                                                                                                                                                                                    `json:"lookup_route_dispatch_gate_modeled"`
	LookupRouteDispatchGateReady                        bool                                                                                                                                                                                                                                    `json:"lookup_route_dispatch_gate_ready"`
	RouteEnablementLookupRouteDispatchGateBoundaryReady bool                                                                                                                                                                                                                                    `json:"route_enablement_lookup_route_dispatch_gate_boundary_ready"`
	OpaqueRouteEnablementReceipt                        bool                                                                                                                                                                                                                                    `json:"opaque_route_enablement_receipt"`
	KDESafeRedactedStatusOnly                           bool                                                                                                                                                                                                                                    `json:"kde_safe_redacted_status_only"`
	CompatibilityCenterLookupRouteDispatchGateModeled   bool                                                                                                                                                                                                                                    `json:"compatibility_center_lookup_route_dispatch_gate_modeled"`
	RuntimeDiagnosticsLookupRouteDispatchGateModeled    bool                                                                                                                                                                                                                                    `json:"runtime_diagnostics_lookup_route_dispatch_gate_modeled"`
	ReceiptPresent                                      bool                                                                                                                                                                                                                                    `json:"receipt_present"`
	ReceiptAccepted                                     bool                                                                                                                                                                                                                                    `json:"receipt_accepted"`
	ReceiptConsumed                                     bool                                                                                                                                                                                                                                    `json:"receipt_consumed"`
	AcceptanceAuthorized                                bool                                                                                                                                                                                                                                    `json:"acceptance_authorized"`
	RouteEnablementAccepted                             bool                                                                                                                                                                                                                                    `json:"route_enablement_accepted"`
	LookupRouteEnablementGranted                        bool                                                                                                                                                                                                                                    `json:"lookup_route_enablement_granted"`
	LookupRouteGrantAuthorized                          bool                                                                                                                                                                                                                                    `json:"lookup_route_grant_authorized"`
	LookupRouteEnabled                                  bool                                                                                                                                                                                                                                    `json:"lookup_route_enabled"`
	LookupRouteDispatchAuthorized                       bool                                                                                                                                                                                                                                    `json:"lookup_route_dispatch_authorized"`
	LookupRouteDispatchGatePassed                       bool                                                                                                                                                                                                                                    `json:"lookup_route_dispatch_gate_passed"`
	LookupRouteDispatchCallable                         bool                                                                                                                                                                                                                                    `json:"lookup_route_dispatch_callable"`
	StorageRootPolicyGrantAuthorized                    bool                                                                                                                                                                                                                                    `json:"storage_root_policy_grant_authorized"`
	RecordWriterCallAuthorized                          bool                                                                                                                                                                                                                                    `json:"record_writer_call_authorized"`
	WriterCallable                                      bool                                                                                                                                                                                                                                    `json:"writer_callable"`
	StorageRootResolved                                 bool                                                                                                                                                                                                                                    `json:"storage_root_resolved"`
	StorageWriteEnabled                                 bool                                                                                                                                                                                                                                    `json:"storage_write_enabled"`
	StatusPersistenceWriteEnabled                       bool                                                                                                                                                                                                                                    `json:"status_persistence_write_enabled"`
	DispatchGateItemCount                               int                                                                                                                                                                                                                                     `json:"dispatch_gate_item_count"`
	RequiredDispatchGateItemCount                       int                                                                                                                                                                                                                                     `json:"required_dispatch_gate_item_count"`
	ReadyDispatchGateItemCount                          int                                                                                                                                                                                                                                     `json:"ready_dispatch_gate_item_count"`
	MissingDispatchGateItemCount                        int                                                                                                                                                                                                                                     `json:"missing_dispatch_gate_item_count"`
	PassedDispatchGateItemCount                         int                                                                                                                                                                                                                                     `json:"passed_dispatch_gate_item_count"`
	EnabledRouteItemCount                               int                                                                                                                                                                                                                                     `json:"enabled_route_item_count"`
	PersistedDispatchGateItemCount                      int                                                                                                                                                                                                                                     `json:"persisted_dispatch_gate_item_count"`
	RawExposedDispatchGateItemCount                     int                                                                                                                                                                                                                                     `json:"raw_exposed_dispatch_gate_item_count"`
	SideEffectDispatchGateItemCount                     int                                                                                                                                                                                                                                     `json:"side_effect_dispatch_gate_item_count"`
	CompatibilityCenterDispatchGateItemCount            int                                                                                                                                                                                                                                     `json:"compatibility_center_dispatch_gate_item_count"`
	RuntimeDiagnosticsDispatchGateItemCount             int                                                                                                                                                                                                                                     `json:"runtime_diagnostics_dispatch_gate_item_count"`
	DispatchGateItems                                   []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchGateAuditItem      `json:"dispatch_gate_items"`
	DispatchGateItemIDs                                 []string                                                                                                                                                                                                                                `json:"dispatch_gate_item_ids"`
	Checks                                              []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchGateAuditCheck     `json:"checks"`
	CheckIDs                                            []string                                                                                                                                                                                                                                `json:"check_ids"`
	Counts                                              ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchGateAuditCheckCounts `json:"counts"`
	ConsumerDryRunLookupResultBoundaryAuthorized        bool                                                                                                                                                                                                                                    `json:"consumer_dry_run_lookup_result_boundary_authorized"`
	ConsumerEnablementAuthorized                        bool                                                                                                                                                                                                                                    `json:"consumer_enablement_authorized"`
	KDEConsumerEnabled                                  bool                                                                                                                                                                                                                                    `json:"kde_consumer_enabled"`
	RuntimeConsumerEnabled                              bool                                                                                                                                                                                                                                    `json:"runtime_consumer_enabled"`
	LookupRouteAuthorized                               bool                                                                                                                                                                                                                                    `json:"lookup_route_authorized"`
	LookupRouteEnablementAuthorized                     bool                                                                                                                                                                                                                                    `json:"lookup_route_enablement_authorized"`
	OpaqueLookupEnabled                                 bool                                                                                                                                                                                                                                    `json:"opaque_lookup_enabled"`
	RedactedSummaryPersisted                            bool                                                                                                                                                                                                                                    `json:"redacted_summary_persisted"`
	KDEStatusPersisted                                  bool                                                                                                                                                                                                                                    `json:"kde_status_persisted"`
	RuntimeDiagnosticsPersisted                         bool                                                                                                                                                                                                                                    `json:"runtime_diagnostics_persisted"`
	DryRunResultPersisted                               bool                                                                                                                                                                                                                                    `json:"dry_run_result_persisted"`
	RawResultExposed                                    bool                                                                                                                                                                                                                                    `json:"raw_result_exposed"`
	DispatchDryRunExecuted                              bool                                                                                                                                                                                                                                    `json:"dispatch_dry_run_executed"`
	RequestObjectCreationEnabled                        bool                                                                                                                                                                                                                                    `json:"request_object_creation_enabled"`
	RequestObjectDispatchEnabled                        bool                                                                                                                                                                                                                                    `json:"request_object_dispatch_enabled"`
	PortalRequestCreated                                bool                                                                                                                                                                                                                                    `json:"portal_request_created"`
	NotificationActionEnabled                           bool                                                                                                                                                                                                                                    `json:"notification_action_enabled"`
	CompatibilityCenterOpened                           bool                                                                                                                                                                                                                                    `json:"compatibility_center_opened"`
	SupportBundleExported                               bool                                                                                                                                                                                                                                    `json:"support_bundle_exported"`
	SupportCaseCreated                                  bool                                                                                                                                                                                                                                    `json:"support_case_created"`
	RuntimeOwned                                        bool                                                                                                                                                                                                                                    `json:"runtime_owned"`
	GoRuntimeBacked                                     bool                                                                                                                                                                                                                                    `json:"go_runtime_backed"`
	KDEPolicyOwner                                      bool                                                                                                                                                                                                                                    `json:"kde_policy_owner"`
	ProductionReadiness                                 bool                                                                                                                                                                                                                                    `json:"production_readiness"`
	ProductionOwnershipReady                            bool                                                                                                                                                                                                                                    `json:"production_ownership_ready"`
	SystemServiceStarted                                bool                                                                                                                                                                                                                                    `json:"system_service_started"`
	SessionBusClaimed                                   bool                                                                                                                                                                                                                                    `json:"session_bus_claimed"`
	ProductionBusClaimed                                bool                                                                                                                                                                                                                                    `json:"production_bus_claimed"`
	WriteMethodsEnabled                                 bool                                                                                                                                                                                                                                    `json:"write_methods_enabled"`
	RuntimeWritesEnabled                                bool                                                                                                                                                                                                                                    `json:"runtime_writes_enabled"`
	DesktopFilesWritten                                 bool                                                                                                                                                                                                                                    `json:"desktop_files_written"`
	KDEConfigurationWritten                             bool                                                                                                                                                                                                                                    `json:"kde_configuration_written"`
	PortalCallExecuted                                  bool                                                                                                                                                                                                                                    `json:"portal_call_executed"`
	AdapterInvocationEnabled                            bool                                                                                                                                                                                                                                    `json:"adapter_invocation_enabled"`
	BackendLaunchEnabled                                bool                                                                                                                                                                                                                                    `json:"backend_launch_enabled"`
	BackendProcessStarted                               bool                                                                                                                                                                                                                                    `json:"backend_process_started"`
	NetworkRequired                                     bool                                                                                                                                                                                                                                    `json:"network_required"`
	HostRootModified                                    bool                                                                                                                                                                                                                                    `json:"host_root_modified"`
	PrivilegedContainerRequired                         bool                                                                                                                                                                                                                                    `json:"privileged_container_required"`
	StateRootPathExposed                                bool                                                                                                                                                                                                                                    `json:"state_root_path_exposed"`
	FilePathsExposed                                    bool                                                                                                                                                                                                                                    `json:"file_paths_exposed"`
	FileContentRead                                     bool                                                                                                                                                                                                                                    `json:"file_content_read"`
	RawCommandExposed                                   bool                                                                                                                                                                                                                                    `json:"raw_command_exposed"`
	RawExecutableExposed                                bool                                                                                                                                                                                                                                    `json:"raw_executable_exposed"`
	BackendDetailsExposed                               bool                                                                                                                                                                                                                                    `json:"backend_details_exposed"`
	BlockedActions                                      []string                                                                                                                                                                                                                                `json:"blocked_actions"`
	NextRequirements                                    []string                                                                                                                                                                                                                                `json:"next_requirements"`
	DesktopSafeSummary                                  string                                                                                                                                                                                                                                  `json:"desktop_safe_summary"`
}

type ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchGateAuditItem struct {
	ID                                                  string `json:"id"`
	ActionKind                                          string `json:"action_kind"`
	SurfaceKind                                         string `json:"surface_kind"`
	StatusConsumerKind                                  string `json:"status_consumer_kind"`
	OpaqueReceiptID                                     string `json:"opaque_receipt_id"`
	EvidencePresent                                     bool   `json:"evidence_present"`
	CurrentMainlineConsumed                             bool   `json:"current_mainline_consumed"`
	RouteEnablementLookupRouteEnablementConsumed        bool   `json:"route_enablement_lookup_route_enablement_consumed"`
	RouteEnablementLookupRouteEnablementReady           bool   `json:"route_enablement_lookup_route_enablement_ready"`
	LookupRouteDispatchGateModeled                      bool   `json:"lookup_route_dispatch_gate_modeled"`
	RouteEnablementLookupRouteDispatchGateBoundaryReady bool   `json:"route_enablement_lookup_route_dispatch_gate_boundary_ready"`
	OpaqueRouteEnablementReceipt                        bool   `json:"opaque_route_enablement_receipt"`
	KDESafeRedactedStatusOnly                           bool   `json:"kde_safe_redacted_status_only"`
	ReceiptPresent                                      bool   `json:"receipt_present"`
	ReceiptAccepted                                     bool   `json:"receipt_accepted"`
	ReceiptConsumed                                     bool   `json:"receipt_consumed"`
	AcceptanceAuthorized                                bool   `json:"acceptance_authorized"`
	RouteEnablementAccepted                             bool   `json:"route_enablement_accepted"`
	LookupRouteEnablementGranted                        bool   `json:"lookup_route_enablement_granted"`
	LookupRouteGrantAuthorized                          bool   `json:"lookup_route_grant_authorized"`
	LookupRouteEnabled                                  bool   `json:"lookup_route_enabled"`
	LookupRouteDispatchAuthorized                       bool   `json:"lookup_route_dispatch_authorized"`
	LookupRouteDispatchGatePassed                       bool   `json:"lookup_route_dispatch_gate_passed"`
	LookupRouteDispatchCallable                         bool   `json:"lookup_route_dispatch_callable"`
	StorageWriteEnabled                                 bool   `json:"storage_write_enabled"`
	RedactedSummaryPersisted                            bool   `json:"redacted_summary_persisted"`
	KDEStatusPersisted                                  bool   `json:"kde_status_persisted"`
	RuntimeDiagnosticsPersisted                         bool   `json:"runtime_diagnostics_persisted"`
	RawResultExposed                                    bool   `json:"raw_result_exposed"`
	UserVisible                                         bool   `json:"user_visible"`
	ReviewOnly                                          bool   `json:"review_only"`
	RuntimeOwned                                        bool   `json:"runtime_owned"`
	GoRuntimeBacked                                     bool   `json:"go_runtime_backed"`
	KDEPolicyOwner                                      bool   `json:"kde_policy_owner"`
	SideEffectsDisabled                                 bool   `json:"side_effects_disabled"`
	HostRootModified                                    bool   `json:"host_root_modified"`
	InternalDetailsExposed                              bool   `json:"internal_details_exposed"`
	LookupRouteDispatchGateStatus                       string `json:"lookup_route_dispatch_gate_status"`
	NextRequirement                                     string `json:"next_requirement"`
}

type ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchGateAuditCheck struct {
	ID      string `json:"id"`
	Status  string `json:"status"`
	Summary string `json:"summary"`
}

type ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchGateAuditCheckCounts struct {
	Total   int `json:"total"`
	Passed  int `json:"passed"`
	Pending int `json:"pending"`
	Blocked int `json:"blocked"`
}

type routeEnablementLookupRouteDispatchGateAuditSourceSet struct {
	CurrentMainline                           string
	RouteEnablementLookupRouteEnablementAudit string
}

func NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchGateAuditPreview(root string) (ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchGateAuditPreview, error) {
	version, err := readProductionDBusGateReviewVersion(root)
	if err != nil {
		return ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchGateAuditPreview{}, err
	}
	sources := routeEnablementLookupRouteDispatchGateAuditSources(root)
	mainlineReady := routeEnablementLookupRouteDispatchGateAuditMainlineReady(sources.CurrentMainline)
	enablementReady := routeEnablementLookupRouteDispatchGateAuditEnablementReady(sources.RouteEnablementLookupRouteEnablementAudit)
	items := routeEnablementLookupRouteDispatchGateAuditItems(mainlineReady, enablementReady)
	readyItemCount := routeEnablementLookupRouteDispatchGateAuditReadyCount(items)
	ready := mainlineReady && enablementReady && readyItemCount == len(items)
	preview := ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchGateAuditPreview{
		Version:                 version,
		SchemaVersion:           "xnix.runtime.production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_kde_safe_redacted_status_closed_record_writer_storage_root_authorization_receipt_dry_run_lookup_result_consumer_projection_route_enablement_lookup_route_dispatch_gate_audit.v1",
		RequestType:             "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-gate-audit-preview",
		AuditType:               "receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-gate-audit",
		Source:                  "docs/xnix-current-mainline.md+route-enable-lookup-route-enablement-audit",
		AuditDecision:           "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-gate-audit-blocked",
		CurrentMainlineConsumed: mainlineReady,
		RouteEnablementLookupRouteEnablementConsumed:        enablementReady,
		RouteEnablementLookupRouteEnablementReady:           enablementReady,
		LookupRouteDispatchGateRequired:                     true,
		LookupRouteDispatchGateModeled:                      true,
		LookupRouteDispatchGateReady:                        ready,
		RouteEnablementLookupRouteDispatchGateBoundaryReady: ready,
		OpaqueRouteEnablementReceipt:                        enablementReady,
		KDESafeRedactedStatusOnly:                           ready,
		CompatibilityCenterLookupRouteDispatchGateModeled:   ready,
		RuntimeDiagnosticsLookupRouteDispatchGateModeled:    ready,
		DispatchGateItemCount:                               len(items),
		RequiredDispatchGateItemCount:                       len(items),
		ReadyDispatchGateItemCount:                          readyItemCount,
		MissingDispatchGateItemCount:                        len(items) - readyItemCount,
		CompatibilityCenterDispatchGateItemCount:            routeEnablementLookupRouteDispatchGateAuditSurfaceCount(items, "compatibility-center"),
		RuntimeDiagnosticsDispatchGateItemCount:             routeEnablementLookupRouteDispatchGateAuditSurfaceCount(items, "runtime-diagnostics"),
		DispatchGateItems:                                   items,
		DispatchGateItemIDs:                                 routeEnablementLookupRouteDispatchGateAuditItemIDs(items),
		RuntimeOwned:                                        true,
		GoRuntimeBacked:                                     true,
		KDEPolicyOwner:                                      false,
		BlockedActions: []string{
			"receipt-present",
			"receipt-accept",
			"receipt-consume",
			"lookup-route-dispatch-gate",
			"lookup-route-enable",
			"lookup-route-dispatch",
			"consumer-enable",
			"dry-run-dispatch",
			"result-persistence",
			"raw-result-exposure",
			"production-ownership",
			"backend-launch",
			"host-mutation",
		},
		NextRequirements: []string{
			"Add a route enablement lookup route dispatch authorization audit before any modeled dispatch gate can execute a dry-run.",
			"Keep lookup route dispatch gate modeled, redacted, and disabled until production ownership is explicitly authorized.",
		},
		DesktopSafeSummary: "KDE-safe redacted route enablement lookup route dispatch gate is audited for Compatibility Center and Runtime diagnostics while receipts, route grants, lookup routes, consumers, dispatch, persistence, production, backend, and host side effects remain disabled.",
	}
	if ready {
		preview.AuditDecision = "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-gate-audit-ready-route-disabled-dispatch-disabled"
	}
	preview.Checks = routeEnablementLookupRouteDispatchGateAuditChecks(preview)
	preview.CheckIDs = routeEnablementLookupRouteDispatchGateAuditCheckIDs(preview.Checks)
	preview.Counts = routeEnablementLookupRouteDispatchGateAuditCountsFor(preview.Checks)
	if err := validateNoBackendTerms(preview, "KDE-safe redacted status dry-run lookup result consumer projection route enablement lookup route dispatch gate audit preview"); err != nil {
		return ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchGateAuditPreview{}, err
	}
	return preview, nil
}

func routeEnablementLookupRouteDispatchGateAuditSources(root string) routeEnablementLookupRouteDispatchGateAuditSourceSet {
	return routeEnablementLookupRouteDispatchGateAuditSourceSet{
		CurrentMainline: productionAuthorizationReadSources(root, []string{"docs/xnix-current-mainline.md"}),
		RouteEnablementLookupRouteEnablementAudit: productionAuthorizationReadSources(root, []string{
			"internal/runtime/owner/route_enablement_lookup_route_enablement_audit_preview.go",
			"internal/runtime/owner/route_enablement_lookup_route_enablement_audit_preview_test.go",
		}),
	}
}

func routeEnablementLookupRouteDispatchGateAuditMainlineReady(source string) bool {
	return productionAuthorizationHasAll(source, []string{
		"route enablement lookup route dispatch gate audit preview",
		"route enablement lookup route enablement audit preview",
		"without accepting or consuming receipts",
	})
}

func routeEnablementLookupRouteDispatchGateAuditEnablementReady(source string) bool {
	return productionAuthorizationHasAll(source, []string{
		"xnix.runtime.production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_kde_safe_redacted_status_closed_record_writer_storage_root_authorization_receipt_dry_run_lookup_result_consumer_projection_route_enablement_lookup_route_enablement_audit.v1",
		"production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-enablement-audit-ready-route-disabled-dispatch-disabled",
		"LookupRouteEnablementReady",
		"RouteEnablementLookupRouteEnablementBoundaryReady",
	})
}

func routeEnablementLookupRouteDispatchGateAuditItems(mainlineReady, enablementReady bool) []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchGateAuditItem {
	actionKinds := []string{"review", "renew", "open-compatibility-center", "dismiss", "support-info"}
	items := make([]ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchGateAuditItem, 0, len(actionKinds)*2)
	for _, actionKind := range actionKinds {
		for _, surfaceKind := range []string{"compatibility-center", "runtime-diagnostics"} {
			ready := mainlineReady && enablementReady
			status := "missing-redacted-status-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-gate-evidence"
			if ready {
				status = "redacted-status-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-gate-modeled-dispatch-disabled-route-disabled"
			}
			items = append(items, ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchGateAuditItem{
				ID:                      actionKind + "-route-enablement-lookup-route-dispatch-gate-" + surfaceKind,
				ActionKind:              actionKind,
				SurfaceKind:             surfaceKind,
				StatusConsumerKind:      "kde-safe-redacted-route-enablement-lookup-route-dispatch-gate",
				OpaqueReceiptID:         "opaque-route-enablement-lookup-route-dispatch-gate-" + surfaceKind + "-" + actionKind,
				EvidencePresent:         ready,
				CurrentMainlineConsumed: mainlineReady,
				RouteEnablementLookupRouteEnablementConsumed:        enablementReady,
				RouteEnablementLookupRouteEnablementReady:           enablementReady,
				LookupRouteDispatchGateModeled:                      ready,
				RouteEnablementLookupRouteDispatchGateBoundaryReady: ready,
				OpaqueRouteEnablementReceipt:                        enablementReady,
				KDESafeRedactedStatusOnly:                           ready,
				UserVisible:                                         ready,
				ReviewOnly:                                          ready,
				RuntimeOwned:                                        true,
				GoRuntimeBacked:                                     true,
				KDEPolicyOwner:                                      false,
				SideEffectsDisabled:                                 true,
				LookupRouteDispatchGateStatus:                       status,
				NextRequirement:                                     "Require a separate lookup route dispatch authorization audit before the modeled dispatch gate can execute any dry-run.",
			})
		}
	}
	return items
}

func routeEnablementLookupRouteDispatchGateAuditReadyCount(items []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchGateAuditItem) int {
	count := 0
	for _, item := range items {
		if item.EvidencePresent && item.LookupRouteDispatchGateModeled && item.RouteEnablementLookupRouteDispatchGateBoundaryReady && item.KDESafeRedactedStatusOnly && !item.ReceiptPresent && !item.ReceiptAccepted && !item.ReceiptConsumed && !item.AcceptanceAuthorized && !item.RouteEnablementAccepted && !item.LookupRouteEnablementGranted && !item.LookupRouteGrantAuthorized && !item.LookupRouteEnabled && !item.LookupRouteDispatchAuthorized && !item.LookupRouteDispatchGatePassed && !item.LookupRouteDispatchCallable && !item.StorageWriteEnabled && !item.RawResultExposed && !item.HostRootModified && !item.InternalDetailsExposed {
			count++
		}
	}
	return count
}

func routeEnablementLookupRouteDispatchGateAuditSurfaceCount(items []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchGateAuditItem, surfaceKind string) int {
	count := 0
	for _, item := range items {
		if item.SurfaceKind == surfaceKind {
			count++
		}
	}
	return count
}

func routeEnablementLookupRouteDispatchGateAuditItemIDs(items []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchGateAuditItem) []string {
	ids := make([]string, 0, len(items))
	for _, item := range items {
		ids = append(ids, item.ID)
	}
	return ids
}

func routeEnablementLookupRouteDispatchGateAuditChecks(preview ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchGateAuditPreview) []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchGateAuditCheck {
	return []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchGateAuditCheck{
		routeEnablementLookupRouteDispatchGateAuditCheck("current-mainline-consumed", preview.CurrentMainlineConsumed, "Xnix current mainline names the result consumer projection route enablement lookup route dispatch gate audit preview."),
		routeEnablementLookupRouteDispatchGateAuditCheck("route-enablement-lookup-route-enablement-consumed", preview.RouteEnablementLookupRouteEnablementConsumed && preview.RouteEnablementLookupRouteEnablementReady, "The lookup route dispatch gate audit consumes the ready route enablement lookup route enablement audit."),
		routeEnablementLookupRouteDispatchGateAuditCheck("lookup-route-dispatch-gate-modeled", preview.LookupRouteDispatchGateRequired && preview.LookupRouteDispatchGateModeled, "Lookup route dispatch gate is modeled without enabling a route."),
		routeEnablementLookupRouteDispatchGateAuditCheck("compatibility-center-and-runtime-lookup-route-dispatch-gates-modeled", preview.CompatibilityCenterLookupRouteDispatchGateModeled && preview.RuntimeDiagnosticsLookupRouteDispatchGateModeled, "Compatibility Center and Runtime diagnostics lookup route dispatch gate boundaries are modeled."),
		routeEnablementLookupRouteDispatchGateAuditCheck("ten-dispatch-gate-items-ready-dispatch-disabled", preview.DispatchGateItemCount == 10 && preview.ReadyDispatchGateItemCount == 10 && preview.MissingDispatchGateItemCount == 0, "Ten KDE-safe lookup route dispatch gate items are ready while route and dispatch gates remain disabled."),
		routeEnablementLookupRouteDispatchGateAuditCheck("lookup-route-dispatch-gate-and-persistence-disabled", !preview.ReceiptPresent && !preview.ReceiptAccepted && !preview.ReceiptConsumed && !preview.AcceptanceAuthorized && !preview.RouteEnablementAccepted && !preview.LookupRouteEnablementGranted && !preview.LookupRouteGrantAuthorized && !preview.LookupRouteEnabled && !preview.LookupRouteDispatchAuthorized && !preview.LookupRouteDispatchGatePassed && !preview.LookupRouteDispatchCallable && !preview.StorageWriteEnabled && !preview.StatusPersistenceWriteEnabled && !preview.RawResultExposed, "Lookup route dispatch gate, lookup grants, lookup routes, storage, persistence, and raw result exposure remain disabled."),
		routeEnablementLookupRouteDispatchGateAuditCheck("consumer-routes-dispatch-and-support-disabled", !preview.ConsumerEnablementAuthorized && !preview.KDEConsumerEnabled && !preview.RuntimeConsumerEnabled && !preview.LookupRouteEnabled && !preview.LookupRouteDispatchAuthorized && !preview.DispatchDryRunExecuted && !preview.SupportBundleExported && !preview.SupportCaseCreated, "Consumer enablement, lookup routes, dispatch, and support side effects remain disabled."),
		routeEnablementLookupRouteDispatchGateAuditCheck("production-and-host-boundary-closed", !preview.ProductionBusClaimed && !preview.WriteMethodsEnabled && !preview.RuntimeWritesEnabled && !preview.BackendLaunchEnabled && !preview.NetworkRequired && !preview.HostRootModified && !preview.PrivilegedContainerRequired, "Production, backend, network, privileged container, and host mutation gates remain closed."),
	}
}

func routeEnablementLookupRouteDispatchGateAuditCheck(id string, ok bool, summary string) ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchGateAuditCheck {
	status := "blocked"
	if ok {
		status = "pass"
	}
	return ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchGateAuditCheck{ID: id, Status: status, Summary: summary}
}

func routeEnablementLookupRouteDispatchGateAuditCheckIDs(checks []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchGateAuditCheck) []string {
	ids := make([]string, 0, len(checks))
	for _, check := range checks {
		ids = append(ids, check.ID)
	}
	return ids
}

func routeEnablementLookupRouteDispatchGateAuditCountsFor(checks []ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchGateAuditCheck) ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchGateAuditCheckCounts {
	counts := ProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchGateAuditCheckCounts{Total: len(checks)}
	for _, check := range checks {
		if check.Status == "pass" {
			counts.Passed++
		} else {
			counts.Blocked++
		}
	}
	return counts
}
